package inspector

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
)

type Input struct {
	Address btcutil.Address
	Index   uint32
	Value   int64
	UTXO    UTXO
}

type Output struct {
	Address btcutil.Address
	Index   uint32
	Value   int64
	UTXO    UTXO
}

type UTXO struct {
	AddressID string         `db:"address_id" json:"address_id"`
	Hash      chainhash.Hash `db:"hash" json:"hash"`
	PkScript  []byte         `db:"pkscript" json:"pkscript"`
	Index     uint32         `db:"index" json:"index"`
	Value     int64          `db:"value" json:"value"`
}

func (u UTXO) ID() string {
	return fmt.Sprintf("%v:%v", u.Hash, u.Index)
}

func (u UTXO) PublicAddress(params *chaincfg.Params) (btcutil.Address, error) {
	if len(u.PkScript) != 25 &&
		u.PkScript[0] != txscript.OP_DUP &&
		u.PkScript[1] != txscript.OP_HASH160 {

		return nil, fmt.Errorf("Invalid pkScript %s : %v", u.Hash, u.PkScript)
	}

	return btcutil.NewAddressPubKeyHash(u.PkScript[3:23], params)
}

// UTXOs is a wrapper for a []UTXO.
type UTXOs []UTXO

// Value returns the total value of the set of UTXO's.
func (u UTXOs) Value() int64 {
	v := int64(0)

	for _, utxo := range u {
		v += utxo.Value
	}

	return v
}

// ForAddress returns UTXOs that match the given Address.
func (u UTXOs) ForAddress(address btcutil.Address) (UTXOs, error) {
	filtered := UTXOs{}

	s := address.String()

	for _, utxo := range u {
		publicAddress, err := utxo.PublicAddress(&chaincfg.MainNetParams)
		if err != nil {
			return nil, err
		}

		if publicAddress.String() != s {
			continue
		}

		filtered = append(filtered, utxo)
	}

	return filtered, nil
}