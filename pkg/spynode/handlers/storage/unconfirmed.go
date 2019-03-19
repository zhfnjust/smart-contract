package storage

import (
	"context"
	"encoding/binary"
	"io"
	"time"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

const (
	unconfirmedTxSize = chainhash.HashSize + 10
)

var (
	TrueData  = []byte{0xff}
	FalseData = []byte{0x00}
)

// Mark an unconfirmed tx as unsafe
// Returns true if the tx was marked
func (repo *TxRepository) MarkUnsafe(ctx context.Context, txid chainhash.Hash) (bool, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	if tx, exists := repo.unconfirmed[txid]; exists {
		tx.unsafe = true
		return true, nil
	}
	return false, nil
}

// Returns all transactions not marked as unsafe or safe that have a "seen" time before the
//   specified time.
// Also marks all returned txs as safe
func (repo *TxRepository) GetNewSafe(ctx context.Context, beforeTime time.Time) ([]chainhash.Hash, error) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	result := make([]chainhash.Hash, 0)
	for hash, tx := range repo.unconfirmed {
		if !tx.safe && !tx.unsafe && tx.time.Before(beforeTime) {
			tx.safe = true
			result = append(result, hash)
		}
	}

	return result, nil
}

type unconfirmedTx struct { // Tx ID hash is key of map containing this struct
	time   time.Time // Time first seen
	unsafe bool      // Conflict seen
	safe   bool      // Safe notification sent
}

func newUnconfirmedTx() *unconfirmedTx {
	result := unconfirmedTx{
		time:   time.Now(),
		unsafe: false,
		safe:   false,
	}
	return &result
}

func (tx *unconfirmedTx) Write(out io.Writer, txid *chainhash.Hash) error {
	var err error

	// TxID
	_, err = out.Write(txid[:])
	if err != nil {
		return err
	}

	// Time
	err = binary.Write(out, binary.LittleEndian, int64(tx.time.UnixNano())/1e6) // Milliseconds
	if err != nil {
		return err
	}

	// Unsafe
	if tx.unsafe {
		_, err = out.Write(TrueData[:])
	} else {
		_, err = out.Write(FalseData[:])
	}
	if err != nil {
		return err
	}

	// Safe
	if tx.safe {
		_, err = out.Write(TrueData[:])
	} else {
		_, err = out.Write(FalseData[:])
	}
	if err != nil {
		return err
	}

	return nil
}

func readUnconfirmedTx(in io.Reader) (chainhash.Hash, *unconfirmedTx, error) {
	var txid chainhash.Hash
	var tx unconfirmedTx
	var err error

	_, err = in.Read(txid[:])
	if err != nil {
		return txid, &tx, err
	}

	// Time
	var milliseconds int64
	err = binary.Read(in, binary.LittleEndian, &milliseconds) // Milliseconds
	if err != nil {
		return txid, &tx, err
	}
	tx.time = time.Unix(0, milliseconds*1e6)

	// Unsafe
	value := []byte{0x00}
	_, err = in.Read(value[:])
	if err != nil {
		return txid, &tx, err
	}
	if value[0] == 0x00 {
		tx.unsafe = false
	} else {
		tx.unsafe = true
	}

	// Safe
	_, err = in.Read(value[:])
	if err != nil {
		return txid, &tx, err
	}
	if value[0] == 0x00 {
		tx.safe = false
	} else {
		tx.safe = true
	}

	return txid, &tx, nil
}