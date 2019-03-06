package spynode

import (
	"context"
	"io"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/tokenized/smart-contract/pkg/spynode/handlers"
	"github.com/tokenized/smart-contract/pkg/spynode/handlers/data"
	handlerstorage "github.com/tokenized/smart-contract/pkg/spynode/handlers/storage"
	"github.com/tokenized/smart-contract/pkg/spynode/logger"
	"github.com/tokenized/smart-contract/pkg/storage"
	"github.com/tokenized/smart-contract/pkg/wire"
)

const (
	MainNetBch wire.BitcoinNet = 0xe8f3e1e3
	TestNetBch wire.BitcoinNet = 0xf4f3e5f4
	RegTestBch wire.BitcoinNet = 0xfabfb5da

	SubSystem = "SpyNode" // For logger
)

// Node is the main object for spynode.
type Node struct {
	config         data.Config                        // Configuration
	state          *data.State                        // Non-persistent data
	store          storage.Storage                    // Persistent data
	peers          *handlerstorage.PeerRepository     // Peer data
	blocks         *handlerstorage.BlockRepository    // Block data
	txs            *handlerstorage.TxRepository       // Tx data
	txTracker      *data.TxTracker                    // Tracks tx requests to ensure all txs are received
	memPool        *data.MemPool                      // Tracks which txs have been received and checked
	handlers       map[string]handlers.CommandHandler // Handlers for messages from trusted node
	conn           net.Conn                           // Connection to trusted node
	outgoing       chan wire.Message                  // Channel for messages to send to trusted node
	outgoingOpen   bool
	listeners      []handlers.Listener  // Receive data and notifications about transactions
	txFilters      []handlers.TxFilter  // Determines if a tx should be seen by listeners
	untrustedNodes []*UntrustedNode     // Randomized peer connections to monitor for double spends
	addresses      map[string]time.Time // Recently used peer addresses
	needsRestart   bool
	stopping       bool
	stopped        bool
	outgoingLock   sync.Mutex
	untrustedMutex sync.Mutex
}

// NewNode creates a new node.
// See handlers/handlers.go for the listener interface definitions.
func NewNode(config data.Config, store storage.Storage) *Node {
	result := Node{
		config:         config,
		state:          data.NewState(),
		store:          store,
		peers:          handlerstorage.NewPeerRepository(store),
		blocks:         handlerstorage.NewBlockRepository(store),
		txs:            handlerstorage.NewTxRepository(store),
		txTracker:      data.NewTxTracker(),
		memPool:        data.NewMemPool(),
		outgoing:       make(chan wire.Message, 100),
		outgoingOpen:   true,
		listeners:      make([]handlers.Listener, 0),
		txFilters:      make([]handlers.TxFilter, 0),
		untrustedNodes: make([]*UntrustedNode, 0),
		addresses:      make(map[string]time.Time),
		needsRestart:   false,
		stopping:       false,
		stopped:        false,
	}
	return &result
}

func (node *Node) RegisterListener(listener handlers.Listener) {
	node.listeners = append(node.listeners, listener)
}

// AddTxFilter adds a tx filter.
// See handlers/filters.go for specification of a filter.
// If no tx filters, then all txs are sent to listeners.
// If any of the tx filters return true the tx will be sent to listeners.
func (node *Node) AddTxFilter(filter handlers.TxFilter) {
	node.txFilters = append(node.txFilters, filter)
}

// load loads the data for the node.
// Must be called after adding filter(s), but before Run()
func (node *Node) load(ctx context.Context) error {
	node.outgoingLock.Lock()
	defer node.outgoingLock.Unlock()

	ctx = logger.ContextWithLogSubSystem(ctx, SubSystem)
	if err := node.peers.Load(ctx); err != nil {
		return err
	}
	logger.Log(ctx, logger.Verbose, "Loaded %d peers", node.peers.Count())

	if err := node.blocks.Load(ctx); err != nil {
		return err
	}
	logger.Log(ctx, logger.Info, "Loaded blocks to height %d", node.blocks.LastHeight())
	startHeight, exists := node.blocks.Height(&node.config.StartHash)
	if exists {
		node.state.StartHeight = startHeight
		logger.Log(ctx, logger.Info, "Start block height %d", node.state.StartHeight)
	} else {
		logger.Log(ctx, logger.Info, "Start block not found yet")
	}

	if err := node.txs.Load(ctx); err != nil {
		return err
	}

	node.handlers = handlers.NewTrustedCommandHandlers(ctx, node.config, node.state, node.peers,
		node.blocks, node.txs, node.txTracker, node.memPool, node.listeners, node.txFilters, node)
	return nil
}

// Run runs the node.
// Doesn't stop until there is a failure or Stop() is called.
func (node *Node) Run(ctx context.Context) error {
	ctx = logger.ContextWithLogSubSystem(ctx, SubSystem)

	var err error = nil
	if err = node.load(ctx); err != nil {
		return err
	}

	initial := true
	for {
		if initial {
			logger.Log(ctx, logger.Verbose, "Connecting to %s", node.config.NodeAddress)
		} else {
			logger.Log(ctx, logger.Verbose, "Re-connecting to %s", node.config.NodeAddress)
			node.state.LogRestart()
		}
		if err = node.connect(ctx); err != nil {
			logger.Log(ctx, logger.Verbose, "Trusted connection failed to %s : %s", node.config.NodeAddress, err.Error())
			break
		}
		initial = false

		node.outgoingLock.Lock()
		if !node.outgoingOpen {
			// Recreate outgoing message channel
			node.outgoing = make(chan wire.Message, 100)
			node.outgoingOpen = true
		}

		// Queue version message to start handshake
		version := buildVersionMsg(node.config.UserAgent, int32(node.blocks.LastHeight()))
		node.outgoing <- version
		node.outgoingLock.Unlock()

		wg := sync.WaitGroup{}
		wg.Add(5)

		go func() {
			defer wg.Done()
			node.monitorIncoming(ctx)
			logger.Log(ctx, logger.Debug, "Monitor incoming finished")
		}()

		go func() {
			defer wg.Done()
			node.monitorRequestTimeouts(ctx)
			logger.Log(ctx, logger.Debug, "Monitor request timeouts finished")
		}()

		go func() {
			defer wg.Done()
			sendOutgoing(ctx, node.conn, node.outgoing)
			logger.Log(ctx, logger.Debug, "Send outgoing finished")
		}()

		go func() {
			defer wg.Done()
			node.checkTxDelays(ctx)
			logger.Log(ctx, logger.Debug, "Check tx delays finished")
		}()

		go func() {
			defer wg.Done()
			node.monitorUntrustedNodes(ctx)
			logger.Log(ctx, logger.Debug, "Monitor untrusted finished")
		}()

		// Block until goroutines finish as a result of Stop()
		wg.Wait()

		// Save block repository
		logger.Log(ctx, logger.Verbose, "Saving")
		node.blocks.Save(ctx)
		node.txs.Save(ctx)

		if !node.needsRestart {
			break
		}

		logger.Log(ctx, logger.Verbose, "Restarting")
		node.needsRestart = false
		node.stopping = false
		node.state.Reset()
	}

	logger.Log(ctx, logger.Verbose, "Stopped")
	node.stopped = true
	return err
}

// Stop closes the connection and causes Run() to return.
func (node *Node) Stop(ctx context.Context) error {
	ctx = logger.ContextWithLogSubSystem(ctx, SubSystem)
	err := node.requestStop(ctx)
	for !node.stopped {
		time.Sleep(100 * time.Millisecond)
	}
	return err
}

func (node *Node) requestStop(ctx context.Context) error {
	node.stopping = true

	if node.outgoingOpen {
		close(node.outgoing)
		node.outgoingOpen = false
	}

	return node.disconnect(ctx)
}

// BroadcastTx broadcasts a tx to the network.
func (node *Node) BroadcastTx(ctx context.Context, tx *wire.MsgTx) error {
	ctx = logger.ContextWithLogSubSystem(ctx, SubSystem)
	logger.Log(ctx, logger.Info, "Broadcasting tx : %s", tx.TxHash())

	// Send to trusted node
	node.outgoingLock.Lock()
	if !node.outgoingOpen {
		node.outgoingLock.Unlock()
		return errors.New("Node inactive")
	}
	node.outgoing <- tx
	node.outgoingLock.Unlock()

	// Send to untrusted nodes
	node.untrustedMutex.Lock()
	for _, untrusted := range node.untrustedNodes {
		untrusted.BroadcastTx(ctx, tx)
	}
	node.untrustedMutex.Unlock()
	return nil
}

// HandleTx processes a tx through spynode as if it came from the network.
// Used to feed "response" txs directly back through spynode.
// TODO Figure out how to handle this if it gets called while a block is being processed and the
//   mempool or tx repo are locked.
func (node *Node) HandleTx(ctx context.Context, tx *wire.MsgTx) error {
	ctx = logger.ContextWithLogSubSystem(ctx, SubSystem)
	ctx = context.WithValue(ctx, handlers.DirectTxKey, true)
	logger.Log(ctx, logger.Info, "Directly handling tx : %s", tx.TxHash())
	err := handleMessage(ctx, node.handlers, tx, &node.outgoingLock, &node.outgoingOpen, node.outgoing)
	if err != nil {
		logger.Log(ctx, logger.Info, "Failed to directly handle tx : %s", err.Error())
	}
	return err
}

// ProcessBlock is called when a block is being processed.
// Implements handlers.BlockProcessor interface
// It is responsible for any cleanup as a result of a block.
func (node *Node) ProcessBlock(ctx context.Context, block *wire.MsgBlock) error {
	logger.Log(ctx, logger.Debug, "Cleaning up after block : %s", block.BlockHash())
	txids, err := block.TxHashes()
	if err != nil {
		return err
	}

	node.txTracker.Remove(ctx, txids)

	node.untrustedMutex.Lock()
	defer node.untrustedMutex.Unlock()

	for _, untrusted := range node.untrustedNodes {
		untrusted.ProcessBlock(ctx, txids)
	}

	return nil
}

func (node *Node) connect(ctx context.Context) error {
	node.disconnect(ctx)

	conn, err := net.Dial("tcp", node.config.NodeAddress)
	if err != nil {
		return err
	}

	node.conn = conn
	now := time.Now()
	node.state.ConnectedTime = &now
	return nil
}

func (node *Node) disconnect(ctx context.Context) error {
	if node.conn == nil {
		return nil
	}

	// close the connection, ignoring any errors
	_ = node.conn.Close()

	node.conn = nil
	return nil
}

// monitorIncoming processes incoming messages.
//
// This is a blocking function that will run forever, so it should be run
// in a goroutine.
func (node *Node) monitorIncoming(ctx context.Context) {
	for {
		if err := node.check(ctx); err != nil {
			logger.Log(ctx, logger.Warn, "Check failed : %v", err.Error())
			node.requestStop(ctx)
			break
		}

		// read new messages, blocking
		if node.conn == nil {
			break
		}
		msg, _, err := wire.ReadMessage(node.conn, wire.ProtocolVersion, MainNetBch)
		if err == io.EOF {
			// Happens when the connection is closed
			logger.Log(ctx, logger.Verbose, "Connection closed")
			node.restart(ctx)
			break
		}
		if err != nil {
			// Happens when the connection is closed
			logger.Log(ctx, logger.Warn, "Failed to read message : %v", err.Error())
			node.restart(ctx)
			break
		}

		if err := handleMessage(ctx, node.handlers, msg, &node.outgoingLock, &node.outgoingOpen, node.outgoing); err != nil {
			logger.Log(ctx, logger.Warn, "Failed to handle (%s) message : %s", msg.Command(), err.Error())
			node.requestStop(ctx)
			break
		}
	}
}

func (node *Node) restart(ctx context.Context) {
	if node.stopping {
		return
	}
	node.needsRestart = true
	node.requestStop(ctx)
}

// check checks the state of spynode and performs state related actions.
func (node *Node) check(ctx context.Context) error {
	node.outgoingLock.Lock()
	defer node.outgoingLock.Unlock()
	if !node.outgoingOpen {
		return nil
	}

	if !node.state.VersionReceived {
		return nil // Still performing handshake
	}

	if !node.state.HandshakeComplete {
		// Send header request to kick off sync
		msg, err := buildHeaderRequest(ctx, node.state.ProtocolVersion, node.blocks, 1, 50)
		if err != nil {
			return err
		}
		node.outgoing <- msg
		now := time.Now()
		node.state.HeadersRequested = &now
		node.state.HandshakeComplete = true
	}

	// Check sync
	if node.state.IsInSync {
		if !node.state.SentSendHeaders {
			// Send sendheaders message to get headers instead of block inventories.
			sendheaders := wire.NewMsgSendHeaders()
			node.outgoing <- sendheaders
			node.state.SentSendHeaders = true
		}

		if !node.state.AddressesRequested {
			addresses := wire.NewMsgGetAddr()
			node.outgoing <- addresses
			node.state.AddressesRequested = true
		}

		if !node.state.MemPoolRequested {
			// Send mempool request
			// This tells the peer to send inventory of all tx in their mempool.
			mempool := wire.NewMsgMemPool()
			node.outgoing <- mempool
			node.state.MemPoolRequested = true
		} else if !node.state.NotifiedSync {
			// TODO Add method to wait for mempool to sync
			for _, listener := range node.listeners {
				listener.HandleInSync(ctx)
			}
			node.state.NotifiedSync = true
		}

		responses, err := node.txTracker.Check(ctx, node.memPool)
		if err != nil {
			return err
		}
		// Queue messages to be sent in response
		for _, response := range responses {
			node.outgoing <- response
		}
	} else if node.state.HeadersRequested == nil && node.state.BlocksRequestedCount() < 5 {
		// Request more headers
		msg, err := buildHeaderRequest(ctx, node.state.ProtocolVersion, node.blocks, 1, 50)
		if err != nil {
			return err
		}
		node.outgoing <- msg
		now := time.Now()
		node.state.HeadersRequested = &now
	}

	return nil
}

// monitorRequestTimeouts monitors for request timeouts.
//
// This is a blocking function that will run forever, so it should be run
// in a goroutine.
func (node *Node) monitorRequestTimeouts(ctx context.Context) {
	for !node.stopping {
		sleepUntilStop(10, &node.stopping) // Only check every 10 seconds

		if err := node.state.CheckTimeouts(); err != nil {
			logger.Log(ctx, logger.Warn, err.Error())
			node.restart(ctx)
			break
		}
	}
}

// checkTxDelays monitors txs for when they have passed the safe tx delay without seeing a
//   conflicting tx.
//
// This is a blocking function that will run forever, so it should be run
// in a goroutine.
func (node *Node) checkTxDelays(ctx context.Context) error {
	for !node.stopping {
		time.Sleep(200 * time.Millisecond)

		if !node.state.IsInSync {
			continue
		}

		// Get newly safe txs
		cutoffTime := time.Now().Add(time.Millisecond * -time.Duration(node.config.SafeTxDelay))
		txids, err := node.txs.GetNewSafe(ctx, cutoffTime)
		if err != nil {
			logger.Log(ctx, logger.Warn, err.Error())
			node.restart(ctx)
			break
		}

		for _, txid := range txids {
			for _, listener := range node.listeners {
				err := listener.HandleTxState(ctx, handlers.ListenerMsgTxStateSafe, txid)
				if err != nil {
					logger.Log(ctx, logger.Warn, err.Error())
				}
			}
		}
	}
	return nil
}

// monitorUntrustedNodes monitors untrusted nodes.
// Attempt to keep the specified number running.
// Watch for when they become inactive and replace them.
//
// This is a blocking function that will run forever, so it should be run
// in a goroutine.
func (node *Node) monitorUntrustedNodes(ctx context.Context) {
	wg := sync.WaitGroup{}
	for !node.stopping {
		node.untrustedMutex.Lock()

		if !node.state.IsInSync {
			node.untrustedMutex.Unlock()
			sleepUntilStop(5, &node.stopping)
			continue
		}

		// Check for inactive
		for {
			removed := false
			for i, untrusted := range node.untrustedNodes {
				if !untrusted.Active {
					// Remove
					node.untrustedNodes = append(node.untrustedNodes[:i], node.untrustedNodes[i+1:]...)
					removed = true
					break
				}
			}

			if !removed {
				break
			}
		}

		count := len(node.untrustedNodes)
		verifiedCount := 0
		for _, untrusted := range node.untrustedNodes {
			if untrusted.state.Verified {
				verifiedCount++
			}
		}

		node.untrustedMutex.Unlock()

		if verifiedCount < node.config.UntrustedCount {
			logger.Log(ctx, logger.Debug, "Untrusted connections : %d", verifiedCount)
		}

		if count < node.config.UntrustedCount/2 {
			// Try for peers with a good score
			for !node.stopping && count < node.config.UntrustedCount/2 {
				if node.addUntrustedNode(ctx, &wg, 5) {
					count++
				} else {
					break
				}
			}
		}

		// Try for peers with a non-negative score
		for !node.stopping && count < node.config.UntrustedCount {
			if node.addUntrustedNode(ctx, &wg, 0) {
				count++
			} else {
				break
			}
		}

		node.peers.Save(ctx)

		if node.stopping {
			break
		}

		sleepUntilStop(5, &node.stopping) // Only check every 5 seconds
	}

	// Stop all
	node.untrustedMutex.Lock()
	for _, untrusted := range node.untrustedNodes {
		untrusted.Stop(ctx)
	}
	node.untrustedMutex.Unlock()

	logger.Log(ctx, logger.Verbose, "Waiting for %d untrusted nodes to finish", len(node.untrustedNodes))
	wg.Wait()
}

// addUntrustedNode adds a new untrusted node.
// Returns true if a new node connection was attempted
func (node *Node) addUntrustedNode(ctx context.Context, wg *sync.WaitGroup, minScore int32) bool {
	// Get new address
	// Check we aren't already connected and haven't used it recently
	peers, err := node.peers.Get(ctx, minScore)
	if err != nil {
		return false
	}

	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	var address string
	for {
		if node.stopping || len(peers) == 0 {
			return false
		}

		if len(peers) == 1 {
			if node.checkAddress(ctx, peers[0].Address) {
				address = peers[0].Address
				break
			} else {
				return false
			}
		}

		// Pick one randomly
		random := seed.Intn(len(peers))
		if node.checkAddress(ctx, peers[random].Address) {
			address = peers[random].Address
			break
		}

		// Remove this address and try again
		peers = append(peers[:random], peers[random+1:]...)
	}

	// Attempt connection
	newNode := NewUntrustedNode(address, node.config, node.store, node.peers, node.blocks, node.txs, node.memPool, node.listeners, node.txFilters)
	node.untrustedMutex.Lock()
	node.untrustedNodes = append(node.untrustedNodes, newNode)
	node.untrustedMutex.Unlock()
	wg.Add(1)
	go func() {
		defer wg.Done()
		newNode.Run(ctx)
	}()
	return true
}

// checkAddress checks if an address was recently used.
func (node *Node) checkAddress(ctx context.Context, address string) bool {
	lastUsed, exists := node.addresses[address]
	if exists {
		if time.Now().Sub(lastUsed).Minutes() > 10 {
			// Address hasn't been used for a while
			node.addresses[address] = time.Now()
			return true
		}

		// Address was used recently
		return false
	}

	// Add address
	node.addresses[address] = time.Now()
	return true
}

func sleepUntilStop(seconds int, stop *bool) {
	for i := 0; i < seconds; i++ {
		if *stop {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
