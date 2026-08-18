package batch

import (
	"encoding/binary"
	"testing"

	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/stretchr/testify/require"
)

// --- fixture helpers -------------------------------------------------------

func nonceScript(t *testing.T, nonce uint64) []byte {
	t.Helper()
	payload := make([]byte, 0, noncePayloadLen)
	payload = append(payload, flrPrefix...)
	payload = binary.BigEndian.AppendUint64(payload, nonce)
	s, err := txscript.NullDataScript(payload)
	require.NoError(t, err)
	return s
}

func opReturn(t *testing.T, data []byte) []byte {
	t.Helper()
	s, err := txscript.NullDataScript(data)
	require.NoError(t, err)
	return s
}

// pkScript returns a distinct, arbitrary (but valid-length) output script.
func pkScript(tag byte) []byte {
	s := make([]byte, 34) // P2WSH-shaped: OP_0 <32 bytes>
	s[0] = txscript.OP_0
	s[1] = 0x20
	for i := 2; i < len(s); i++ {
		s[i] = tag
	}
	return s
}

func out(script []byte, v int64) Output { return Output{PkScript: script, Value: v} }

// prefix builds the fixed [anchor, P2A, nonce] output prefix.
func prefix(t *testing.T, nonce uint64) []Output {
	t.Helper()
	return []Output{
		out(pkScript(0x01), 1000),     // [0] anchor
		out(pkScript(0x02), 330),      // [1] P2A
		out(nonceScript(t, nonce), 0), // [2] nonce OP_RETURN
	}
}

var (
	refA = []byte{0xa0, 0xa1, 0xa2}
	refB = []byte{0xb0, 0xb1, 0xb2}
	refC = []byte{0xc0, 0xc1, 0xc2}
)

// --- tests -----------------------------------------------------------------

func TestParseBatch_HappyPath(t *testing.T) {
	recipA := pkScript(0xAA)
	recipC := pkScript(0xCC)
	change := pkScript(0xFF)

	outs := append(prefix(t, 5),
		out(opReturn(t, refA), 0), // group 0 opener
		out(recipA, 1000),         //   delivered
		out(opReturn(t, refB), 0), // group 1 opener (undeliverable, no value)
		out(opReturn(t, refC), 0), // group 2 opener (last)
		out(recipC, 500),          //   delivered
		out(change, 250),          //   trailing change -> lands in last group
	)

	b, err := ParseBatch(outs)
	require.NoError(t, err)
	require.Equal(t, uint64(5), b.Nonce)
	require.Len(t, b.Groups, 3)

	require.Equal(t, refA, b.Groups[0].Reference)
	require.Len(t, b.Groups[0].Outputs, 1)

	require.Equal(t, refB, b.Groups[1].Reference)
	require.Empty(t, b.Groups[1].Outputs) // K=0

	require.Equal(t, refC, b.Groups[2].Reference)
	require.Len(t, b.Groups[2].Outputs, 2) // recipient + trailing change
}

func TestParseBatch_ValueBeforeGroupRejected(t *testing.T) {
	outs := append(prefix(t, 1), out(pkScript(0xAA), 500))
	_, err := ParseBatch(outs)
	require.ErrorIs(t, err, ErrValueBeforeGroup)
}

func TestParseBatch_EmptyBatch(t *testing.T) {
	b, err := ParseBatch(prefix(t, 2))
	require.NoError(t, err)
	require.Equal(t, uint64(2), b.Nonce)
	require.Empty(t, b.Groups)
}

func TestParseNonce_Errors(t *testing.T) {
	t.Run("too few outputs", func(t *testing.T) {
		_, err := ParseNonce([]Output{out(pkScript(0x01), 0)})
		require.ErrorIs(t, err, ErrTooFewOutputs)
	})
	t.Run("output[2] not OP_RETURN", func(t *testing.T) {
		outs := []Output{out(pkScript(1), 0), out(pkScript(2), 0), out(pkScript(3), 0)}
		_, err := ParseNonce(outs)
		require.ErrorIs(t, err, ErrNonceOpReturn)
	})
	t.Run("wrong prefix", func(t *testing.T) {
		bad, err := txscript.NullDataScript(append([]byte{0x00, 0x00, 0x00, 0x00}, make([]byte, 8)...))
		require.NoError(t, err)
		outs := []Output{out(pkScript(1), 0), out(pkScript(2), 0), out(bad, 0)}
		_, err = ParseNonce(outs)
		require.ErrorIs(t, err, ErrNonceOpReturn)
	})
	t.Run("wrong length", func(t *testing.T) {
		short, err := txscript.NullDataScript(flrPrefix)
		require.NoError(t, err)
		outs := []Output{out(pkScript(1), 0), out(pkScript(2), 0), out(short, 0)}
		_, err = ParseNonce(outs)
		require.ErrorIs(t, err, ErrNonceOpReturn)
	})
	t.Run("valid nonce round-trips", func(t *testing.T) {
		n, err := ParseNonce(prefix(t, 0xDEADBEEF))
		require.NoError(t, err)
		require.Equal(t, uint64(0xDEADBEEF), n)
	})
}

func TestFromMsgTx(t *testing.T) {
	tx := wire.NewMsgTx(wire.TxVersion)
	tx.AddTxIn(wire.NewTxIn(&wire.OutPoint{}, nil, nil))
	tx.AddTxOut(wire.NewTxOut(1000, pkScript(0x01)))
	tx.AddTxOut(wire.NewTxOut(330, pkScript(0x02)))
	tx.AddTxOut(wire.NewTxOut(0, nonceScript(t, 9)))
	tx.AddTxOut(wire.NewTxOut(0, opReturn(t, refA)))
	tx.AddTxOut(wire.NewTxOut(1000, pkScript(0xAA)))

	b, err := ParseBatch(FromMsgTx(tx))
	require.NoError(t, err)
	require.Equal(t, uint64(9), b.Nonce)
	require.Len(t, b.Groups, 1)
	require.Equal(t, refA, b.Groups[0].Reference)
	require.Len(t, b.Groups[0].Outputs, 1)
}

func TestParseBatchWithChangeSeparatesChange(t *testing.T) {
	recipA := pkScript(0xAA)
	change := pkScript(0xFF)

	// Two change outputs, the shape a batch takes when it is replenishing the
	// confirmed pool it leaves for its successor.
	outs := append(prefix(t, 7),
		out(opReturn(t, refA), 0),
		out(recipA, 1000),
		out(change, 250),
		out(change, 400),
	)

	b, err := ParseBatchWithChange(outs, change)
	require.NoError(t, err)
	require.Len(t, b.Groups, 1)
	require.Len(t, b.Groups[0].Outputs, 1) // the recipient only
	require.Len(t, b.Change, 2)
	require.Equal(t, int64(250), b.Change[0].Value)
	require.Equal(t, int64(400), b.Change[1].Value)
}

// A K=0 group is the reference OP_RETURN standing alone. Change-blind parsing
// puts the trailing change inside it, and a caller that reads "this group holds
// outputs" as "it paid someone other than its recipient" then rejects an honest
// batch. Change-aware parsing is what keeps that check meaningful.
func TestParseBatchWithChangeLeavesAK0GroupEmpty(t *testing.T) {
	change := pkScript(0xFF)
	outs := append(prefix(t, 8),
		out(opReturn(t, refA), 0),
		out(change, 900),
	)

	blind, err := ParseBatch(outs)
	require.NoError(t, err)
	require.Len(t, blind.Groups[0].Outputs, 1) // change absorbed into the group

	aware, err := ParseBatchWithChange(outs, change)
	require.NoError(t, err)
	require.Empty(t, aware.Groups[0].Outputs) // genuinely K=0
	require.Len(t, aware.Change, 1)
}

func TestParseBatchWithChangeNilIsUnchanged(t *testing.T) {
	recipA := pkScript(0xAA)
	change := pkScript(0xFF)
	outs := append(prefix(t, 9),
		out(opReturn(t, refA), 0),
		out(recipA, 1000),
		out(change, 250),
	)

	withNil, err := ParseBatchWithChange(outs, nil)
	require.NoError(t, err)
	plain, err := ParseBatch(outs)
	require.NoError(t, err)
	require.Equal(t, plain, withNil)
	require.Empty(t, withNil.Change)
}
