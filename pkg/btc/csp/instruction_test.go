package csp_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/btc/csp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// proposerKey signs the packages these tests build. The envelope names the
// matching address, because BindPackage now requires the two to agree.
var proposerKey, _ = crypto.HexToECDSA("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318")

// packageOf returns what a proposer publishes: the envelope followed by its
// signature over the envelope's own hash.
func packageOf(t *testing.T, e csp.Envelope, chainID uint64) []byte {
	t.Helper()
	encoded, err := e.Encode()
	require.NoError(t, err)
	h, err := e.Hash(chainID)
	require.NoError(t, err)
	sig, err := crypto.Sign(accounts.TextHash(h[:]), proposerKey)
	require.NoError(t, err)
	return append(append([]byte{}, encoded...), sig...)
}

// boundPair returns a published package and the instruction that decided on it.
func boundPair(t *testing.T) ([]byte, csp.Instruction) {
	t.Helper()
	e := sample()
	e.ProposerAddress = crypto.PubkeyToAddress(proposerKey.PublicKey)
	raw := packageOf(t, e, testChainID)
	h := crypto.Keccak256Hash(raw)
	return raw, csp.Instruction{
		TeeIDKeyIDPairs: []csp.TeeIDKeyIDPair{
			{TeeID: common.HexToAddress("0xA11CE"), KeyID: 1},
			{TeeID: common.HexToAddress("0xB0B"), KeyID: 2},
		},
		WalletID:         e.WalletID,
		SourceID:         e.SourceID,
		AccountIndex:     e.AccountIndex,
		SequencePosition: e.SequencePosition,
		Attempt:          e.Attempt,
		AnchorIndex:      e.AnchorIndex,
		Nonce:            e.Nonce,
		PackageHash:      h,
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

func TestBindPackageAcceptsTheDecidedPackage(t *testing.T) {
	raw, in := boundPair(t)
	_, err := in.BindPackage(raw, testChainID)
	require.NoError(t, err)
}

// A bare envelope is what the store held before proposals were committed to.
// It must no longer bind: the signature is what says the proposal is the work
// of the party being paid for it.
func TestBindPackageRejectsAnUnsignedProposal(t *testing.T) {
	e := sample()
	e.ProposerAddress = crypto.PubkeyToAddress(proposerKey.PublicKey)
	encoded, err := e.Encode()
	require.NoError(t, err)
	in := csp.Instruction{
		WalletID: e.WalletID, SourceID: e.SourceID, AccountIndex: e.AccountIndex,
		SequencePosition: e.SequencePosition, Attempt: e.Attempt,
		AnchorIndex: e.AnchorIndex, Nonce: e.Nonce,
		PackageHash: crypto.Keccak256Hash(encoded),
	}
	_, err = in.BindPackage(encoded, testChainID)
	assert.Error(t, err, "an unsigned proposal was accepted for signing")
}

// The theft this check prevents: republishing a rival's envelope under your own
// commitment. The envelope still names its true author, so the signature is
// what catches it.
func TestBindPackageRejectsAnotherPartysSignature(t *testing.T) {
	other, err := crypto.HexToECDSA("1111111111111111111111111111111111111111111111111111111111111111")
	require.NoError(t, err)
	e := sample()
	e.ProposerAddress = crypto.PubkeyToAddress(proposerKey.PublicKey)
	encoded, err := e.Encode()
	require.NoError(t, err)
	h, err := e.Hash(testChainID)
	require.NoError(t, err)
	sig, err := crypto.Sign(accounts.TextHash(h[:]), other)
	require.NoError(t, err)
	raw := append(append([]byte{}, encoded...), sig...)

	in := csp.Instruction{
		WalletID: e.WalletID, SourceID: e.SourceID, AccountIndex: e.AccountIndex,
		SequencePosition: e.SequencePosition, Attempt: e.Attempt,
		AnchorIndex: e.AnchorIndex, Nonce: e.Nonce,
		PackageHash: crypto.Keccak256Hash(raw),
	}
	_, err = in.BindPackage(raw, testChainID)
	assert.Error(t, err)
}

// The DAL is untrusted content-addressed storage. Without this check a machine
// would sign whatever bytes it was handed.
func TestBindPackageRejectsASubstitutedProposal(t *testing.T) {
	_, in := boundPair(t)

	tampered := sample()
	tampered.ProposerAddress = crypto.PubkeyToAddress(proposerKey.PublicKey)
	tampered.RawUnsignedTx = append(tampered.RawUnsignedTx, 0xff)
	_, err := in.BindPackage(packageOf(t, tampered, testChainID), testChainID)
	assert.Error(t, err, "a proposal with different transaction bytes was accepted")

	redirected := sample()
	redirected.ProposerAddress = crypto.PubkeyToAddress(proposerKey.PublicKey)
	redirected.Inputs[0].ValueSat++
	_, err = in.BindPackage(packageOf(t, redirected, testChainID), testChainID)
	assert.Error(t, err, "a proposal with a different input value was accepted")
}

// A proposal decided on one chain must not bind on another.
func TestBindPackageIsChainBound(t *testing.T) {
	raw, in := boundPair(t)
	_, err := in.BindPackage(raw, testChainID+1)
	assert.Error(t, err)
}

func TestBindPackageRejectsIdentityMismatch(t *testing.T) {
	raw, in := boundPair(t)
	in.Attempt++ // hash still matches the package; the identity no longer does
	_, err := in.BindPackage(raw, testChainID)
	assert.Error(t, err)
}

func TestKeysForSelectsThisMachine(t *testing.T) {
	me := common.HexToAddress("0xA11CE")
	other := common.HexToAddress("0xB0B")
	in := csp.Instruction{TeeIDKeyIDPairs: []csp.TeeIDKeyIDPair{
		{TeeID: other, KeyID: 1},
		{TeeID: me, KeyID: 7},
		{TeeID: other, KeyID: 2},
	}}
	assert.Equal(t, []uint64{7}, in.KeysFor(me))
	// A machine that is not a signer for this wallet gets nothing, so it
	// declines rather than hunting for a key it was never asked to use.
	assert.Empty(t, in.KeysFor(common.HexToAddress("0xDEAD")))
}
