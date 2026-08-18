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
//	[3…]  payment groups, in instruction order, then the change outputs (0, 1 or
//	      several — see ParseBatchWithChange)
//
// Payment-group grammar (parsed left-to-right from index 3): an OP_RETURN opens a
// new group (its payload is that payment's reference); the value outputs that
// follow attach to the current group until the next OP_RETURN. Trailing change
// outputs land in the last group unless the caller supplies the change script
// (ParseBatchWithChange) — the matcher (Match) filters a group's outputs by the
// expected recipient, and the grammar guarantees a recipient script is never the
// change script, so change is excluded either way. Keeping the change script
// optional keeps the parser usable by callers that cannot derive it.
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
	// Change holds the outputs paying the wallet's own change script, in
	// transaction order, when the caller supplied that script. Empty otherwise —
	// change then stays in the last group, which is what Match already tolerates.
	Change []Output
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
	return ParseBatchWithChange(outputs, nil)
}

// ParseBatchWithChange parses outputs the same way, but routes every output
// paying changeScript into Batch.Change instead of the current group.
//
// WHY A CALLER WOULD WANT THAT. Change is identified by its SCRIPT, never by its
// position or by how many outputs there are. A batch may emit SEVERAL change
// outputs — sized on the denomination ladder — so that the confirmed pool it
// leaves behind is wide enough for the next batch to fund itself without
// spending an unconfirmed coin (which BIP-431 forbids for a batch that must stay
// fee-bumpable). Change-blind parsing folds those outputs into the last group,
// where they are merely ignored by Match; that is fine for reading a payment,
// and not fine for judging a proposal, because "this group carries value that
// does not pay its recipient" is exactly the check that catches a batch paying
// someone else.
//
// changeScript nil parses exactly as before, so callers that do not know the
// wallet's change address are unaffected.
func ParseBatchWithChange(outputs []Output, changeScript []byte) (*Batch, error) {
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
		if len(changeScript) > 0 && bytes.Equal(out.PkScript, changeScript) {
			b.Change = append(b.Change, Output{PkScript: out.PkScript, Value: out.Value})
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
