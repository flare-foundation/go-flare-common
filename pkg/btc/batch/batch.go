// Package batch parses a Flare PMW batch Bitcoin transaction into its batch nonce
// and ordered payment groups. The batch layout is a protocol invariant shared by
// the proposer (which builds batches), the relay client, and the FDC verifiers
// (which read them), so the grammar lives here as a single source of truth.
//
// Transaction output layout:
//
//	[0]   next anchor UTXO (chain continuation)
//	[1]   P2A fee-bump output (BIP-433)
//	[2]   nonce OP_RETURN: OP_RETURN <FLR\0><uint64 big-endian nonce>
//	[3…]  payment groups, in instruction order, then an optional change output
//
// Payment-group grammar (parsed left-to-right from index 3): an OP_RETURN opens a
// new group (its payload is that payment's reference); the value outputs that
// follow attach to the current group until the next OP_RETURN. The transaction's
// trailing change output, if any, lands in the last group's outputs — it is not
// separated here because the matcher (Match) filters a group's outputs by the
// expected recipient, and the grammar guarantees a recipient script is never the
// change script, so change is naturally excluded. This keeps the parser free of
// the wallet's change-address derivation.
//
// The parser operates on a neutral []Output slice so callers can feed it either a
// live transaction (via FromMsgTx) or reconstructed outputs (e.g. from an indexer
// database). It validates layout only; matching a group against the on-chain
// instruction is Match's / the caller's job.
package batch

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

const (
	// prefixOutputs is the fixed output prefix before payment groups:
	// [0] anchor, [1] P2A fee-bump, [2] nonce OP_RETURN.
	prefixOutputs = 3
	// nonceOutputIndex is the index of the nonce OP_RETURN.
	nonceOutputIndex = 2
	// noncePayloadLen is the nonce OP_RETURN payload: FLR\0 (4) + uint64 (8).
	noncePayloadLen = 12
)

// flrPrefix is the 4-byte "FLR\0" marker distinguishing the Flare nonce OP_RETURN
// from a payment-reference OP_RETURN.
var flrPrefix = []byte{0x46, 0x4C, 0x52, 0x00}

// Errors returned by the parser. All indicate a transaction that does not match
// the PMW batch layout.
var (
	ErrTooFewOutputs    = errors.New("btcbatch: transaction has fewer than the 3 prefix outputs")
	ErrNonceOpReturn    = errors.New("btcbatch: output[2] is not a valid Flare nonce OP_RETURN")
	ErrValueBeforeGroup = errors.New("btcbatch: value output before the first payment group")
)

// Output is one transaction output: its locking script and value in satoshis.
type Output struct {
	PkScript []byte
	Value    int64
}

// Group is one payment group: the opening reference and the value outputs that
// follow it (up to the next group). For the last group these may include the
// transaction's trailing change output; Match filters by recipient so change is
// excluded.
type Group struct {
	Reference []byte
	Outputs   []Output
}

// Batch is the parsed view of a PMW batch transaction.
type Batch struct {
	Nonce  uint64  // batch nonce from the index-2 OP_RETURN
	Groups []Group // payment groups in instruction order (output[3]…)
}

// FromMsgTx projects a wire transaction's outputs into the neutral []Output form
// the parser consumes. Convenience for callers holding a live transaction.
func FromMsgTx(tx *wire.MsgTx) []Output {
	outputs := make([]Output, len(tx.TxOut))
	for i, o := range tx.TxOut {
		outputs[i] = Output{PkScript: o.PkScript, Value: o.Value}
	}
	return outputs
}

// ParseNonce extracts the batch nonce from the index-2 OP_RETURN
// (OP_RETURN <FLR\0><uint64 big-endian>).
func ParseNonce(outputs []Output) (uint64, error) {
	if len(outputs) <= nonceOutputIndex {
		return 0, ErrTooFewOutputs
	}
	data, ok := opReturnData(outputs[nonceOutputIndex].PkScript)
	if !ok || len(data) != noncePayloadLen || !bytes.Equal(data[:len(flrPrefix)], flrPrefix) {
		return 0, ErrNonceOpReturn
	}
	return binary.BigEndian.Uint64(data[len(flrPrefix):]), nil
}

// ParseBatch parses outputs into the batch nonce and ordered payment groups. It
// validates layout only (not against any Flare instruction). A returned error
// means the outputs do not match the PMW batch layout.
func ParseBatch(outputs []Output) (*Batch, error) {
	nonce, err := ParseNonce(outputs)
	if err != nil {
		return nil, err
	}

	b := &Batch{Nonce: nonce}
	curIdx := -1 // index in b.Groups of the group currently taking value outputs
	for i := prefixOutputs; i < len(outputs); i++ {
		out := outputs[i]
		if ref, ok := opReturnData(out.PkScript); ok {
			b.Groups = append(b.Groups, Group{Reference: ref})
			curIdx = len(b.Groups) - 1
			continue
		}
		if curIdx < 0 {
			return nil, ErrValueBeforeGroup
		}
		b.Groups[curIdx].Outputs = append(b.Groups[curIdx].Outputs, Output{PkScript: out.PkScript, Value: out.Value})
	}
	return b, nil
}

// opReturnData returns the concatenated pushed data of an OP_RETURN script and
// true, or (nil, false) if the script is not an OP_RETURN. An OP_RETURN whose
// pushes cannot be parsed yields empty data (still reported as an OP_RETURN).
func opReturnData(pkScript []byte) ([]byte, bool) {
	if len(pkScript) == 0 || pkScript[0] != txscript.OP_RETURN {
		return nil, false
	}
	pushes, err := txscript.PushedData(pkScript)
	if err != nil {
		return nil, true
	}
	return bytes.Join(pushes, nil), true
}
