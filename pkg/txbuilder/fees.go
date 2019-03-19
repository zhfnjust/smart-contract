package txbuilder

import (
	"github.com/tokenized/smart-contract/pkg/wire"
)

const (
	// P2PKH/P2SH input size 147
	//   Previous Transaction ID = 32 bytes
	//   Previous Transaction Output Index = 4 bytes
	//   script size = 2 bytes
	//   Signature push to stack = 75
	//       push size = 1 byte
	//       signature up to = 73 bytes
	//       signature hash type = 1 byte
	//   Public key push to stack = 34
	//       push size = 1 byte
	//       public key size = 33 bytes
	estimatedInputSize = 32 + 4 + 2 + 75 + 34

	// P2PKH/P2SH output size 33
	//   amount = 8 bytes
	//   script size = 1 byte
	//   Script (24 bytes) OP_DUP OP_HASH160 <PUB KEY/SCRIPT HASH (20 bytes)> OP_EQUALVERIFY
	//     OP_CHECKSIG
	// estimatedOutputSize = 8 + 25

	// BaseTxFee is the size of the tx not included in inputs and outputs.
	//   Version = 4 bytes
	//   LockTime = 4 bytes
	baseTxSize = 8
)

// The fee should be estimated before signing, then after signing the fee should be checked.
// If the fee is too low after signing, then the fee should be adjusted and the tx re-signed.

func (tx *Tx) Fee() uint64 {
	return tx.inputSum() - tx.outputSum(true)
}

// EstimatedSize returns the estimated size in bytes of the tx after signatures are added.
// It assumes all inputs are P2PKH.
func (tx *Tx) EstimatedSize() int {
	result := baseTxSize + wire.VarIntSerializeSize(uint64(len(tx.MsgTx.TxIn))) +
		wire.VarIntSerializeSize(uint64(len(tx.MsgTx.TxOut)))

	for _, input := range tx.MsgTx.TxIn {
		if len(input.SignatureScript) > 0 {
			result += input.SerializeSize()
		} else {
			result += estimatedInputSize
		}
	}

	for _, output := range tx.MsgTx.TxOut {
		result += output.SerializeSize()
	}

	return result
}

// inputSum returns the sum of the values of the inputs.
func (tx *Tx) inputSum() uint64 {
	inputValue := uint64(0)
	for _, input := range tx.Inputs {
		inputValue += input.Value
	}
	return inputValue
}

// outputSum returns the sum of the values of the outputs.
func (tx *Tx) outputSum(includeChange bool) uint64 {
	outputValue := uint64(0)
	for i, output := range tx.MsgTx.TxOut {
		if includeChange || !tx.Outputs[i].IsChange {
			outputValue += uint64(output.Value)
		}
	}
	return outputValue
}

// changeSum returns the sum of the values of the outputs.
func (tx *Tx) changeSum() uint64 {
	changeValue := uint64(0)
	for i, output := range tx.MsgTx.TxOut {
		if tx.Outputs[i].IsChange {
			changeValue += uint64(output.Value)
		}
	}
	return changeValue
}

// adjustFee
func (tx *Tx) adjustFee(amount int64) error {
	if amount == int64(0) {
		return nil
	}

	// Find change output
	changeOutputIndex := 0xffffffff
	for i, output := range tx.Outputs {
		if output.IsChange {
			changeOutputIndex = i
			break
		}
	}

	if amount > int64(0) {
		// Increase fee, transfer from change
		if changeOutputIndex == 0xffffffff {
			return InputValueInsufficientError // No existing change to move to fee
		}

		if tx.MsgTx.TxOut[changeOutputIndex].Value < amount {
			return InputValueInsufficientError // Not enough change to move to fee
		}

		// Decrease change, thereby increasing the fee
		tx.MsgTx.TxOut[changeOutputIndex].Value -= amount
	} else {
		// Decrease fee, transfer to change
		if changeOutputIndex == 0xffffffff {
			// Add a change output
			tx.AddP2PKHOutput(tx.ChangePKH, uint64(-amount), true, false)
			return nil
		}

		// Increase change, thereby decreasing the fee
		// (amount is negative so subracting it increases the change value)
		tx.MsgTx.TxOut[changeOutputIndex].Value -= amount
	}

	return nil
}