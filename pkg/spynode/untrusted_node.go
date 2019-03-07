package spynode

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/pkg/errors"
	"github.com/tokenized/smart-contract/pkg/spynode/handlers"
	"github.com/tokenized/smart-contract/pkg/spynode/handlers/data"
	handlerstorage "github.com/tokenized/smart-contract/pkg/spynode/handlers/storage"
	"github.com/tokenized/smart-contract/pkg/spynode/logger"
	"github.com/tokenized/smart-contract/pkg/storage"
	"github.com/tokenized/smart-contract/pkg/wire"
)

type UntrustedNode struct {
	address    string
	config     data.Config
	state      *data.UntrustedState
	peers      *handlerstorage.PeerRepository
	blocks     *handlerstorage.BlockRepository
	txs        *handlerstorage.TxRepository
	txTracker  *data.TxTracker
	memPool    *data.MemPool
	handlers   map[string]handlers.CommandHandler
	connection net.Conn
	outgoing   chan wire.Message
	listeners  []handlers.Listener
	txFilters  []handlers.TxFilter
	stopping   bool
	Active     bool // Set to false when connection is closed
	lock       sync.Mutex
}

func NewUntrustedNode(address string, config data.Config, store storage.Storage, peers *handlerstorage.PeerRepository, blocks *handlerstorage.BlockRepository, txs *handlerstorage.TxRepository, memPool *data.MemPool, listeners []handlers.Listener, txFilters []handlers.TxFilter) *UntrustedNode {
	result := UntrustedNode{
		address:   address,
		config:    config,
		state:     data.NewUntrustedState(),
		peers:     peers,
		blocks:    blocks,
		txs:       txs,
		txTracker: data.NewTxTracker(),
		memPool:   memPool,
		outgoing:  make(chan wire.Message, 100),
		listeners: listeners,
		txFilters: txFilters,
		stopping:  false,
		Active:    false,
	}
	return &result
}

// Run the node
// Doesn't stop until there is a failure or Stop() is called.
func (node *UntrustedNode) Run(ctx context.Context) error {
	node.lock.Lock()
	if node.stopping {
		node.lock.Unlock()
		return nil
	}

	node.handlers = handlers.NewUntrustedCommandHandlers(ctx, node.state, node.peers, node.blocks, node.txs, node.txTracker, node.memPool, node.listeners, node.txFilters)

	if err := node.connect(); err != nil {
		node.lock.Unlock()
		node.peers.UpdateScore(ctx, node.address, -1)
		logger.Log(ctx, logger.Debug, "Connection failed to %s : %s", node.address, err.Error())
		return err
	}
	node.lock.Unlock()

	node.Active = true
	defer func() {
		node.Active = false
	}()

	// Queue version message to start handshake
	version := buildVersionMsg(node.config.UserAgent, int32(node.blocks.LastHeight()))
	node.outgoing <- version

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		node.monitorIncoming(ctx)
		logger.Log(ctx, logger.Debug, "Untrusted monitor incoming finished")
	}()

	go func() {
		defer wg.Done()
		node.monitorRequestTimeouts(ctx)
		logger.Log(ctx, logger.Debug, "Untrusted monitor request timeouts finished")
	}()

	go func() {
		defer wg.Done()
		node.sendOutgoing(ctx)
		logger.Log(ctx, logger.Debug, "Untrusted send outgoing finished")
	}()

	// Block until goroutines finish as a result of Stop()
	wg.Wait()
	return nil
}

func (node *UntrustedNode) Stop(ctx context.Context) error {
	node.lock.Lock()
	defer node.lock.Unlock()

	if node.stopping {
		return nil
	}

	node.stopping = true
	close(node.outgoing)
	if node.connection != nil {
		return node.connection.Close()
	}
	return nil
}

// Broadcast a tx to the peer
func (node *UntrustedNode) BroadcastTx(ctx context.Context, tx *wire.MsgTx) error {
	if !node.queueOutgoing(tx) {
		return errors.New("Node inactive")
	}
	return nil
}

// ProcessBlock is called when a block is being processed.
// It is responsible for any cleanup as a result of a block.
func (node *UntrustedNode) ProcessBlock(ctx context.Context, txids []chainhash.Hash) error {
	node.txTracker.Remove(ctx, txids)
	return nil
}

func (node *UntrustedNode) connect() error {
	conn, err := net.DialTimeout("tcp", node.address, 10*time.Second)
	if err != nil {
		return err
	}

	node.connection = conn
	now := time.Now()
	node.state.ConnectedTime = &now
	return nil
}

// monitorIncoming monitors incoming messages.
//
// This is a blocking function that will run forever, so it should be run
// in a goroutine.
func (node *UntrustedNode) monitorIncoming(ctx context.Context) {
	for !node.stopping {
		if err := node.check(ctx); err != nil {
			logger.Log(ctx, logger.Debug, "Check failed : %s", err.Error())
			node.Stop(ctx)
			break
		}

		if node.stopping {
			break
		}

		// read new messages, blocking
		msg, _, err := wire.ReadMessage(node.connection, wire.ProtocolVersion, MainNetBch)
		if err == io.EOF {
			// Happens when the connection is closed
			logger.Log(ctx, logger.Debug, "Connection closed")
			node.Stop(ctx)
			break
		}
		if err != nil {
			// Happens when the connection is closed
			logger.Log(ctx, logger.Debug, "Failed to read message : %s", err.Error())
			node.Stop(ctx)
			break
		}

		if err := node.handleMessage(ctx, msg); err != nil {
			node.peers.UpdateScore(ctx, node.address, -1)
			logger.Log(ctx, logger.Debug, "Failed to handle (%s) message : %s", msg.Command(), err.Error())
			node.Stop(ctx)
			break
		}
	}
}

// Check state
func (node *UntrustedNode) check(ctx context.Context) error {
	if !node.state.VersionReceived {
		return nil // Still performing handshake
	}

	if !node.state.HandshakeComplete {
		// Send header request to verify chain
		headerRequest, err := buildHeaderRequest(ctx, node.state.ProtocolVersion, node.blocks, handlers.UntrustedHeaderDelta, 10)
		if err != nil {
			return err
		}
		if node.queueOutgoing(headerRequest) {
			now := time.Now()
			node.state.HeadersRequested = &now
			node.state.HandshakeComplete = true
		}
	}

	// Check sync
	if !node.state.Verified {
		return nil
	}

	if !node.state.ScoreUpdated {
		node.peers.UpdateScore(ctx, node.address, 5)
		node.state.ScoreUpdated = true
	}

	if !node.state.AddressesRequested {
		addresses := wire.NewMsgGetAddr()
		if node.queueOutgoing(addresses) {
			node.state.AddressesRequested = true
		}
	}

	if !node.state.MemPoolRequested {
		// Send mempool request
		// This tells the peer to send inventory of all tx in their mempool.
		mempool := wire.NewMsgMemPool()
		if node.queueOutgoing(mempool) {
			node.state.MemPoolRequested = true
		}
	}

	responses, err := node.txTracker.Check(ctx, node.memPool)
	if err != nil {
		return err
	}
	// Queue messages to be sent in response
	for _, response := range responses {
		if !node.queueOutgoing(response) {
			break
		}
	}

	return nil
}

// Monitor for request timeouts
func (node *UntrustedNode) monitorRequestTimeouts(ctx context.Context) {
	for !node.stopping {
		sleepUntilStop(10, &node.stopping) // Only check every 10 seconds
		if node.stopping {
			break
		}

		if err := node.state.CheckTimeouts(); err != nil {
			logger.Log(ctx, logger.Debug, "Timed out : %s", err.Error())
			node.peers.UpdateScore(ctx, node.address, -1)
			node.Stop(ctx)
			break
		}
	}
}

// sendOutgoing waits for and sends outgoing messages
//
// This is a blocking function that will run forever, so it should be run
// in a goroutine.
func (node *UntrustedNode) sendOutgoing(ctx context.Context) error {
	for !node.stopping {
		// Wait for outgoing message on channel
		msg, ok := <-node.outgoing

		if !ok || node.stopping {
			break
		}

		if err := sendAsync(ctx, node.connection, msg); err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to send %s : %v", msg.Command()))
		}
	}

	return nil
}

// handleMessage Processes an incoming message
func (node *UntrustedNode) handleMessage(ctx context.Context, msg wire.Message) error {
	if node.stopping {
		return nil
	}

	handler, ok := node.handlers[msg.Command()]
	if !ok {
		// no handler for this command
		return nil
	}

	responses, err := handler.Handle(ctx, msg)
	if err != nil {
		return err
	}

	// Queue messages to be sent in response
	for _, response := range responses {
		if !node.queueOutgoing(response) {
			break
		}
	}

	return nil
}

func (node *UntrustedNode) queueOutgoing(msg wire.Message) bool {
	node.lock.Lock()
	defer node.lock.Unlock()
	if node.stopping {
		return false
	}
	node.outgoing <- msg
	return true
}
