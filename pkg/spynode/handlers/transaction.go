package handlers

import (
	"context"

	"github.com/tokenized/smart-contract/pkg/logger"
	"github.com/tokenized/smart-contract/pkg/spynode/handlers/data"
	"github.com/tokenized/smart-contract/pkg/spynode/handlers/storage"
	"github.com/tokenized/smart-contract/pkg/wire"

	"github.com/pkg/errors"
)

// TXHandler exists to handle the tx command.
type TXHandler struct {
	ready     StateReady
	memPool   *data.MemPool
	txs       *storage.TxRepository
	listeners []Listener
	txFilters []TxFilter
}

// NewTXHandler returns a new TXHandler with the given Config.
func NewTXHandler(ready StateReady, memPool *data.MemPool, txs *storage.TxRepository, listeners []Listener, txFilters []TxFilter) *TXHandler {
	result := TXHandler{
		ready:     ready,
		memPool:   memPool,
		txs:       txs,
		listeners: listeners,
		txFilters: txFilters,
	}
	return &result
}

type TxKey int

var DirectTxKey TxKey = 1 // Used in context to flag when a tx is from the system

// Handle implements the handler interface for transaction handler.
func (handler *TXHandler) Handle(ctx context.Context, m wire.Message) ([]wire.Message, error) {
	msg, ok := m.(*wire.MsgTx)
	if !ok {
		return nil, errors.New("Could not assert as *wire.MsgTx")
	}

	hash := msg.TxHash()
	logger.Debug(ctx, "Received tx : %s", hash)

	// Only notify of transactions when in sync or they might be duplicated
	if !handler.ready.IsReady() && ctx.Value(DirectTxKey) == nil {
		return nil, nil
	}

	// The mempool is needed to track which transactions have been sent to listeners and to check
	//   for attempted double spends.
	conflicts, added := handler.memPool.AddTransaction(msg)

	if !added {
		return nil, nil // Already saw this tx
	}

	if len(conflicts) > 0 {
		logger.Warn(ctx, "Found %d conflicts with %s", len(conflicts), hash)
		// Notify of attempted double spend
		for _, conflict := range conflicts {
			marked, err := handler.txs.MarkUnsafe(ctx, *conflict)
			if err != nil {
				return nil, errors.Wrap(err, "Failed to check tx repo")
			}
			if marked { // Only send for txs that previously matched filters.
				for _, listener := range handler.listeners {
					if err := listener.HandleTxState(ctx, ListenerMsgTxStateUnsafe, *conflict); err != nil {
						return nil, err
					}
				}
			}
		}
	}

	// We have to succesfully add to tx repo because it is protected by a lock and will prevent
	//   processing the same tx twice at the same time.
	if added, err := handler.txs.Add(ctx, hash, -1); err != nil {
		return nil, errors.Wrap(err, "Failed to add to tx repo")
	} else if !added {
		return nil, nil // Already seen
	}

	if !matchesFilter(ctx, msg, handler.txFilters) {
		if _, err := handler.txs.Remove(ctx, hash, -1); err != nil {
			return nil, errors.Wrap(err, "Failed to remove from tx repo")
		}
		return nil, nil // Filter out
	}

	// Notify of new tx
	marked := false
	var mark bool
	var err error
	for _, listener := range handler.listeners {
		if mark, err = listener.HandleTx(ctx, msg); err != nil {
			return nil, err
		}
		if mark {
			marked = true
		}
	}

	if marked {
		// Notify of conflicting txs
		if len(conflicts) > 0 {
			handler.txs.MarkUnsafe(ctx, hash)
			for _, listener := range handler.listeners {
				if err := listener.HandleTxState(ctx, ListenerMsgTxStateUnsafe, hash); err != nil {
					return nil, err
				}
			}
		}
	} else {
		// Remove from tx repository
		if _, err := handler.txs.Remove(ctx, hash, -1); err != nil {
			return nil, errors.Wrap(err, "Failed to remove from tx repo")
		}
	}

	return nil, nil
}
