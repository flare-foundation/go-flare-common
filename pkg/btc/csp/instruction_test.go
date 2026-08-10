package csp_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/btc/csp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// boundPair returns an envelope and the instruction that decided on it.
func boundPair(t *testing.T) (csp.Envelope, csp.Instruction) {
	t.Helper()
	e := sample()
	h, err := e.Hash(testChainID)
	require.NoError(t, err)
	return e, csp.Instruction{
		WalletID:         e.WalletID,
		SourceID:         e.SourceID,
		AccountIndex:     e.AccountIndex,
		SequencePosition: e.SequencePosition,
		Attempt:          e.Attempt,
		AnchorIndex:      e.AnchorIndex,
		Nonce:            e.Nonce,
		ProposalHash:     h,
		Txid:             common.HexToHash("0xfeed"),
		Proposer:         e.ProposerAddress,
	}
}

func TestInstructionRoundTrip(t *testing.T) {
	_, in := boundPair(t)
	b, err := csp.EncodeInstruction(in)
	require.NoError(t, err)
	got, err := csp.DecodeInstruction(b)
	require.NoError(t, err)
	assert.Equal(t, in, got)
}

func TestBindEnvelopeAcceptsTheDecidedEnvelope(t *testing.T) {
	e, in := boundPair(t)
	require.NoError(t, in.BindEnvelope(e, testChainID))
}

// The DAL is untrusted content-addressed storage. Without this check a machine
// would sign whatever bytes it was handed.
func TestBindEnvelopeRejectsASubstitutedEnvelope(t *testing.T) {
	_, in := boundPair(t)

	tampered := sample()
	tampered.RawUnsignedTx = append(tampered.RawUnsignedTx, 0xff)
	assert.Error(t, in.BindEnvelope(tampered, testChainID),
		"an envelope with different transaction bytes was accepted")

	redirected := sample()
	redirected.Inputs[0].ValueSat++
	assert.Error(t, in.BindEnvelope(redirected, testChainID),
		"an envelope with a different input value was accepted")
}

// A proposal decided on one chain must not bind on another.
func TestBindEnvelopeIsChainBound(t *testing.T) {
	e, in := boundPair(t)
	assert.Error(t, in.BindEnvelope(e, testChainID+1))
}

func TestBindEnvelopeRejectsIdentityMismatch(t *testing.T) {
	e, in := boundPair(t)
	in.Attempt++ // hash still matches the envelope; the identity no longer does
	assert.Error(t, in.BindEnvelope(e, testChainID))
}
