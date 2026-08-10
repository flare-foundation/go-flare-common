package csp_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/btc/csp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testChainID = 31337

func sample() csp.Envelope {
	return csp.Envelope{
		Version:            1,
		WalletID:           [32]byte([]byte("wallet-1\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")),
		SourceID:           [32]byte([]byte("BTC\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00")),
		AccountIndex:       0,
		SequencePosition:   7,
		EligibleGeneration: 3,
		InstructionType:    0,
		PaymentCount:       2,
		AnchorIndex:        1,
		Nonce:              42,
		Attempt:            0,
		RawUnsignedTx:      []byte{0x03, 0x00, 0x00, 0x00, 0x01, 0xde, 0xad},
		Inputs: []csp.Input{
			{Txid: [32]byte{0x11}, Vout: 0, ValueSat: 500_000_000, Chain: 0, Index: 0},
			{Txid: [32]byte{0x22}, Vout: 1, ValueSat: 1_000, Chain: 1, Index: 5},
		},
		ProposerAddress: common.HexToAddress("0xA11CE"),
	}
}

func TestRoundTrip(t *testing.T) {
	e := sample()
	b, err := e.Encode()
	require.NoError(t, err)
	got, err := csp.Decode(b)
	require.NoError(t, err)
	again, err := got.Encode()
	require.NoError(t, err)
	require.Equal(t, b, again, "round trip is not byte-stable")
	assert.Equal(t, e.Inputs, got.Inputs, "inputs did not survive the round trip")
}

// The injectivity property the encoding exists for. abi.encodePacked("a","bc")
// and abi.encodePacked("ab","c") are the same bytes; abi.encode never is,
// because every dynamic value is length-prefixed and offset.
func TestAdjacentDynamicFieldsCannotBeResplit(t *testing.T) {
	a := sample()
	a.RawUnsignedTx = []byte{0x01}
	a.Inputs = []csp.Input{{Txid: [32]byte{0x02, 0x03}, ValueSat: 1}}

	b := sample()
	b.RawUnsignedTx = []byte{0x01, 0x02}
	b.Inputs = []csp.Input{{Txid: [32]byte{0x03}, ValueSat: 1}}

	ea, err := a.Encode()
	require.NoError(t, err)
	eb, err := b.Encode()
	require.NoError(t, err)
	require.NotEqual(t, ea, eb, "two distinct envelopes encoded identically - the commitment is not injective")
}

// Every field must reach the hash. A field that silently does not is a field an
// adversary can vary freely while the proposal still verifies.
func TestEveryFieldChangesTheHash(t *testing.T) {
	base, err := sample().Hash(testChainID)
	require.NoError(t, err)

	mutations := map[string]func(*csp.Envelope){
		"version":            func(e *csp.Envelope) { e.Version = 2 },
		"walletId":           func(e *csp.Envelope) { e.WalletID[0] ^= 0xff },
		"sourceId":           func(e *csp.Envelope) { e.SourceID[0] ^= 0xff },
		"accountIndex":       func(e *csp.Envelope) { e.AccountIndex = 1 },
		"sequencePosition":   func(e *csp.Envelope) { e.SequencePosition++ },
		"eligibleGeneration": func(e *csp.Envelope) { e.EligibleGeneration++ },
		"instructionType":    func(e *csp.Envelope) { e.InstructionType = 1 },
		"paymentCount":       func(e *csp.Envelope) { e.PaymentCount++ },
		"anchorIndex":        func(e *csp.Envelope) { e.AnchorIndex++ },
		"nonce":              func(e *csp.Envelope) { e.Nonce++ },
		"attempt":            func(e *csp.Envelope) { e.Attempt++ },
		"rawUnsignedTx":      func(e *csp.Envelope) { e.RawUnsignedTx = append(e.RawUnsignedTx, 0x99) },
		"input.txid":         func(e *csp.Envelope) { e.Inputs[0].Txid[3] ^= 0xff },
		"input.vout":         func(e *csp.Envelope) { e.Inputs[0].Vout++ },
		"input.valueSat":     func(e *csp.Envelope) { e.Inputs[0].ValueSat++ },
		"input.chain":        func(e *csp.Envelope) { e.Inputs[0].Chain = 1 },
		"input.index":        func(e *csp.Envelope) { e.Inputs[0].Index++ },
		"input.count":        func(e *csp.Envelope) { e.Inputs = e.Inputs[:1] },
		"proposerAddress":    func(e *csp.Envelope) { e.ProposerAddress = common.HexToAddress("0xB0B") },
	}

	for name, mutate := range mutations {
		e := sample()
		mutate(&e)
		h, err := e.Hash(testChainID)
		require.NoError(t, err)
		assert.NotEqual(t, base, h, "mutating %s did not change the hash - it is not committed to", name)
	}
}

// The value BIP-143 signs over must be committed: a proposal whose declared
// value differs from the real one produces signatures over the wrong sighash.
func TestInputOrderIsCommitted(t *testing.T) {
	a := sample()
	b := sample()
	b.Inputs[0], b.Inputs[1] = b.Inputs[1], b.Inputs[0]

	ha, err := a.Hash(testChainID)
	require.NoError(t, err)
	hb, err := b.Hash(testChainID)
	require.NoError(t, err)
	require.NotEqual(t, ha, hb, "swapping two inputs left the hash unchanged - input order is not committed")
}

// A proposal for one network must not verify on another.
func TestChainIDSeparatesTheHash(t *testing.T) {
	e := sample()
	h1, err := e.Hash(1)
	require.NoError(t, err)
	h2, err := e.Hash(2)
	require.NoError(t, err)
	require.NotEqual(t, h1, h2, "the same envelope hashed identically on two chains")
}

func TestDecodeRejectsTrailingBytes(t *testing.T) {
	b, err := sample().Encode()
	require.NoError(t, err)
	_, err = csp.Decode(append(b, 0x00))
	require.Error(t, err, "trailing bytes were accepted - two byte strings would decode to one message")
}

func TestValidateRejectsUnusableEnvelopes(t *testing.T) {
	cases := map[string]func(*csp.Envelope){
		"no inputs":       func(e *csp.Envelope) { e.Inputs = nil },
		"empty tx":        func(e *csp.Envelope) { e.RawUnsignedTx = nil },
		"bad chain value": func(e *csp.Envelope) { e.Inputs[0].Chain = 2 },
	}
	for name, mutate := range cases {
		e := sample()
		mutate(&e)
		assert.Error(t, e.Validate(), "%s: Validate accepted it", name)
		_, err := e.Encode()
		assert.Error(t, err, "%s: Encode accepted it - Validate must gate Encode", name)
	}
}
