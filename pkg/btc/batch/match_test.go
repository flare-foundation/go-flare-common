package batch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatch_Success(t *testing.T) {
	recip := pkScript(0xAA)
	g := Group{Reference: refA, Outputs: []Output{out(recip, 1000)}}
	st, recv, err := Match(g, Expected{Reference: refA, Recipient: recip, Amount: 1000})
	require.NoError(t, err)
	require.Equal(t, StatusSuccess, st)
	require.Equal(t, int64(1000), recv)
}

func TestMatch_SuccessMultiValue(t *testing.T) {
	recip := pkScript(0xAA)
	g := Group{Reference: refA, Outputs: []Output{out(recip, 700), out(recip, 300)}}
	st, recv, err := Match(g, Expected{Reference: refA, Recipient: recip, Amount: 1000})
	require.NoError(t, err)
	require.Equal(t, StatusSuccess, st)
	require.Equal(t, int64(1000), recv)
}

func TestMatch_SuccessIgnoresTrailingChange(t *testing.T) {
	recip := pkScript(0xAA)
	change := pkScript(0xFF)
	// Last group: recipient output + trailing change (different script) — change ignored.
	g := Group{Reference: refA, Outputs: []Output{out(recip, 1000), out(change, 250)}}
	st, recv, err := Match(g, Expected{Reference: refA, Recipient: recip, Amount: 1000})
	require.NoError(t, err)
	require.Equal(t, StatusSuccess, st)
	require.Equal(t, int64(1000), recv)
}

func TestMatch_UndeliverableNoOutputs(t *testing.T) {
	g := Group{Reference: refA} // K=0
	st, recv, err := Match(g, Expected{Reference: refA, Recipient: pkScript(0xAA), Amount: 1000})
	require.NoError(t, err)
	require.Equal(t, StatusUndeliverable, st)
	require.Equal(t, int64(0), recv)
}

func TestMatch_UndeliverableOnlyChange(t *testing.T) {
	// K=0 last group followed by change: no output pays the recipient -> undeliverable.
	g := Group{Reference: refA, Outputs: []Output{out(pkScript(0xFF), 250)}}
	st, _, err := Match(g, Expected{Reference: refA, Recipient: pkScript(0xAA), Amount: 1000})
	require.NoError(t, err)
	require.Equal(t, StatusUndeliverable, st)
}

func TestMatch_ReferenceMismatch(t *testing.T) {
	recip := pkScript(0xAA)
	g := Group{Reference: refB, Outputs: []Output{out(recip, 1000)}}
	_, _, err := Match(g, Expected{Reference: refA, Recipient: recip, Amount: 1000})
	require.ErrorIs(t, err, ErrReferenceMismatch)
}

func TestMatch_AmountMismatch(t *testing.T) {
	recip := pkScript(0xAA)
	g := Group{Reference: refA, Outputs: []Output{out(recip, 999)}}
	_, _, err := Match(g, Expected{Reference: refA, Recipient: recip, Amount: 1000})
	require.ErrorIs(t, err, ErrAmountMismatch)
}
