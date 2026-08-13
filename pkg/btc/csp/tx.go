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

// opReturn is OP_RETURN. Spelled out rather than imported so this file keeps
// depending only on wire.
const opReturn = 0x6a

// PaymentGroups counts the payments the transaction actually makes.
//
// The batch grammar fixes outputs [0..2] as the anchor, the P2A and the nonce
// OP_RETURN. Everything after is a sequence of payment groups, each OPENED by a
// reference OP_RETURN, with at most one change output last. So the number of
// OP_RETURNs past index 2 is the number of payments — and change, never an
// OP_RETURN, cannot be miscounted as one.
//
// This is what binds the envelope's declared PaymentCount to the transaction.
// The count decides the range the chain marks settled, so without this check a
// proposal could claim more payments than it pays and the difference would be
// recorded as delivered while the money never moved.
func PaymentGroups(e Envelope) (uint32, error) {
	tx, err := Tx(e)
	if err != nil {
		return 0, err
	}
	if len(tx.TxOut) < 3 {
		return 0, fmt.Errorf("batch has %d outputs, fewer than the three the grammar fixes", len(tx.TxOut))
	}
	var n uint32
	for _, o := range tx.TxOut[3:] {
		if len(o.PkScript) > 0 && o.PkScript[0] == opReturn {
			n++
		}
	}
	return n, nil
}
