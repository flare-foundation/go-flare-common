package csp

import (
	"bytes"
	"fmt"

	"github.com/btcsuite/btcd/wire"
)

// Tx parses the envelope's transaction.
//
// RawUnsignedTx is the NON-WITNESS serialization, so this is also the byte
// string the txid is computed over.
func Tx(e Envelope) (*wire.MsgTx, error) {
	var tx wire.MsgTx
	if err := tx.DeserializeNoWitness(bytes.NewReader(e.RawUnsignedTx)); err != nil {
		return nil, fmt.Errorf("decoding unsigned transaction: %w", err)
	}
	return &tx, nil
}

// TxidOf is the transaction id the envelope commits to.
//
// Because the envelope carries the non-witness form, the txid follows from the
// bytes directly rather than by re-deriving it from a witness-stripped copy —
// which is what makes "finalizing the hash pins the txid" exact rather than
// approximate.
func TxidOf(e Envelope) [32]byte {
	tx, err := Tx(e)
	if err != nil {
		return [32]byte{}
	}
	return [32]byte(tx.TxHash())
}

// OutputSum is the total value the transaction pays out.
//
// Subtracted from the input total it gives the fee, which is why both the
// verifier's score and the proposer's fee check go through here rather than
// each parsing outputs their own way.
func OutputSum(e Envelope) (uint64, error) {
	tx, err := Tx(e)
	if err != nil {
		return 0, err
	}
	var sum uint64
	for i, o := range tx.TxOut {
		if o.Value < 0 {
			return 0, fmt.Errorf("output %d has negative value", i)
		}
		sum += uint64(o.Value)
	}
	return sum, nil
}

// InputSum is the total value the envelope's inputs bring in.
func InputSum(e Envelope) uint64 {
	var sum uint64
	for _, i := range e.Inputs {
		sum += i.ValueSat
	}
	return sum
}
