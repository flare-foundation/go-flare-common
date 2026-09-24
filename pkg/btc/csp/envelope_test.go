package csp_test

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/btc/csp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testChainID = 31337

// signerXpubs returns n distinct wallet-level keys at m/87'/coinType', raw — the
// form an envelope carries.
func signerXpubs(n int, coinType uint32, params *chaincfg.Params) [][]byte {
	out := make([][]byte, n)
	for i := range n {
		seed := make([]byte, 32)
		for j := range seed {
			seed[j] = byte(i*13 + j + 1)
		}
		root, err := hdkeychain.NewMaster(seed, params)
		if err != nil {
			panic(err)
		}
		purpose, _ := root.Derive(hdkeychain.HardenedKeyStart + 87)
		coin, _ := purpose.Derive(hdkeychain.HardenedKeyStart + coinType)
		pub, _ := coin.Neuter()
		raw, err := csp.DecodeXpub(pub.String())
		if err != nil {
			panic(err)
		}
		out[i] = raw
	}
	return out
}

func sample() csp.Envelope {
	return csp.Envelope{
		Version:            csp.EnvelopeVersion,
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
		ParentXpubs:     signerXpubs(3, csp.CoinTypeTestnet, &chaincfg.RegressionNetParams),
		Threshold:       2,
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
		"threshold":          func(e *csp.Envelope) { e.Threshold = 3 },
		"parentXpubs.order":  func(e *csp.Envelope) { e.ParentXpubs[0], e.ParentXpubs[1] = e.ParentXpubs[1], e.ParentXpubs[0] },
		"parentXpubs.count":  func(e *csp.Envelope) { e.ParentXpubs = e.ParentXpubs[:2] },
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

		// Version 2 is the only one: a v1 package commits to a tuple that no
		// longer exists.
		"version 1": func(e *csp.Envelope) { e.Version = 1 },
		"version 3": func(e *csp.Envelope) { e.Version = 3 },

		// The wallet's shape.
		"no signers":        func(e *csp.Envelope) { e.ParentXpubs = nil },
		"threshold zero":    func(e *csp.Envelope) { e.Threshold = 0 },
		"threshold above n": func(e *csp.Envelope) { e.Threshold = 4 },
		"more than 20":      func(e *csp.Envelope) { e.ParentXpubs = signerXpubs(21, 1, &chaincfg.RegressionNetParams) },
		"duplicate signer":  func(e *csp.Envelope) { e.ParentXpubs[2] = e.ParentXpubs[0] },
		"short xpub":        func(e *csp.Envelope) { e.ParentXpubs[0] = e.ParentXpubs[0][:77] },
		"wrong depth": func(e *csp.Envelope) {
			e.ParentXpubs[1] = append([]byte{}, e.ParentXpubs[1]...)
			e.ParentXpubs[1][4] = 3
		},
		"not a public key": func(e *csp.Envelope) {
			e.ParentXpubs[1] = append([]byte{}, e.ParentXpubs[1]...)
			e.ParentXpubs[1][45] = 0x00
		},
	}
	for name, mutate := range cases {
		e := sample()
		mutate(&e)
		assert.Error(t, e.Validate(), "%s: Validate accepted it", name)
		_, err := e.Encode()
		assert.Error(t, err, "%s: Encode accepted it - Validate must gate Encode", name)
	}
}

// A base58 key round-trips through the raw form without change, and a private
// key is refused: it would put a secret into a published hash preimage.
func TestXpubEncoding(t *testing.T) {
	e := sample()
	for i, s := range e.XpubStrings() {
		raw, err := csp.DecodeXpub(s)
		require.NoError(t, err)
		assert.Equal(t, e.ParentXpubs[i], raw)
		assert.Equal(t, s, csp.EncodeXpub(raw))
	}

	root, err := hdkeychain.NewMaster(make([]byte, 32), &chaincfg.RegressionNetParams)
	require.NoError(t, err)
	_, err = csp.DecodeXpub(root.String())
	assert.ErrorContains(t, err, "PRIVATE")
}

// There is no coin-type field: the coin is the child number every key records
// for its last derivation step, m/87'/coin'.
func TestCoinTypeIsReadFromTheKeys(t *testing.T) {
	e := sample()
	for _, x := range e.ParentXpubs {
		assert.Equal(t, uint32(0x80000000+csp.CoinTypeTestnet), binary.BigEndian.Uint32(x[9:13]))
	}
	ct, err := e.CoinType()
	require.NoError(t, err)
	assert.Equal(t, csp.CoinTypeTestnet, ct)
	p, err := e.KeyParams()
	require.NoError(t, err)
	for _, net := range []*chaincfg.Params{&chaincfg.TestNet3Params, &chaincfg.SigNetParams, &chaincfg.RegressionNetParams} {
		ct, err := csp.CoinTypeFor(net)
		require.NoError(t, err)
		assert.Equal(t, csp.CoinTypeTestnet, ct, net.Name)
		assert.Equal(t, p.HDPublicKeyID, net.HDPublicKeyID, "%s shares the tpub version", net.Name)
	}

	mainnet := sample()
	mainnet.ParentXpubs = signerXpubs(3, csp.CoinTypeBitcoin, &chaincfg.MainNetParams)
	require.NoError(t, mainnet.Validate())
	ct, err = mainnet.CoinType()
	require.NoError(t, err)
	assert.Equal(t, csp.CoinTypeBitcoin, ct)
	p, err = mainnet.KeyParams()
	require.NoError(t, err)
	assert.Equal(t, chaincfg.MainNetParams.HDPublicKeyID, p.HDPublicKeyID)
	ct, err = csp.CoinTypeFor(&chaincfg.MainNetParams)
	require.NoError(t, err)
	assert.Equal(t, csp.CoinTypeBitcoin, ct)
}

// withChildNumber returns copies of keys with their child number rewritten —
// the label a BIP-32 key carries for its last derivation step.
func withChildNumber(keys [][]byte, child uint32) [][]byte {
	out := make([][]byte, len(keys))
	for i, k := range keys {
		out[i] = append([]byte{}, k...)
		binary.BigEndian.PutUint32(out[i][9:13], child)
	}
	return out
}

// Each refusal is asserted by its reason, so a case cannot pass by tripping
// over some other check first.
func TestCoinTypeRefusals(t *testing.T) {
	cases := map[string]struct {
		mutate func(*csp.Envelope)
		reason string
	}{
		// Every key individually valid, but not for one network.
		"keys disagree on coin": {
			func(e *csp.Envelope) {
				e.ParentXpubs[1] = signerXpubs(1, csp.CoinTypeBitcoin, &chaincfg.MainNetParams)[0]
			},
			"disagree on the coin",
		},
		"unsupported coin in the child number": {
			func(e *csp.Envelope) { e.ParentXpubs = withChildNumber(e.ParentXpubs, 0x80000000+3) },
			"coin type 3",
		},
		// A tpub whose child number says m/87'/0': the child number names the
		// coin, and the version bytes must be that coin's.
		"version bytes do not match the child number's coin": {
			func(e *csp.Envelope) {
				e.ParentXpubs = signerXpubs(3, csp.CoinTypeBitcoin, &chaincfg.RegressionNetParams)
			},
			"do not match coin type 0",
		},
		// And the converse: an xpub at m/87'/1'.
		"mainnet version bytes on a coin-1 key": {
			func(e *csp.Envelope) { e.ParentXpubs = signerXpubs(3, csp.CoinTypeTestnet, &chaincfg.MainNetParams) },
			"do not match coin type 1",
		},
		// m/87'/1 is not the wallet level: the coin step is hardened.
		"child number not hardened": {
			func(e *csp.Envelope) { e.ParentXpubs = withChildNumber(e.ParentXpubs, 1) },
			"not hardened",
		},
	}
	for name, c := range cases {
		e := sample()
		c.mutate(&e)
		_, err := e.CoinType()
		assert.ErrorContains(t, err, c.reason, "%s: CoinType", name)
		assert.ErrorContains(t, e.Validate(), c.reason, "%s: Validate", name)
		_, err = e.KeyParams()
		assert.ErrorContains(t, err, c.reason, "%s: KeyParams", name)
		_, err = e.Encode()
		assert.Error(t, err, "%s: Encode accepted it - Validate must gate Encode", name)
	}
}

// layoutFixture populates EVERY field, escrow included, and gives each numeric
// field a value no other field has, so each word of the encoding can be traced
// to exactly one field and a swap of two fields shows.
func layoutFixture() csp.Envelope {
	var wallet [32]byte
	for i := range wallet {
		wallet[i] = 0x11
	}
	var source [32]byte
	copy(source[:], "BTC")
	custodian := append([]byte{0x02}, bytes.Repeat([]byte{0xcc}, 32)...)
	return csp.Envelope{
		Version:            csp.EnvelopeVersion,
		WalletID:           wallet,
		SourceID:           source,
		AccountIndex:       7,
		SequencePosition:   1001,
		EligibleGeneration: 17,
		InstructionType:    4,
		PaymentCount:       6,
		AnchorIndex:        5,
		Nonce:              42,
		Attempt:            1,
		RawUnsignedTx:      []byte{0x03, 0x00, 0x00, 0x00, 0x01, 0xde, 0xad},
		Inputs: []csp.Input{
			{Txid: [32]byte{0x21}, Vout: 2, ValueSat: 50_000, Chain: 0, Index: 5},
			{Txid: [32]byte{0x22}, Vout: 1, ValueSat: 40_000, Chain: 1, Index: 0, IsEscrow: true},
		},
		ProposerAddress: common.HexToAddress("0x00000000000000000000000000000000000A11CE"),
		Escrow: csp.Escrow{
			PreimageHash:    [32]byte{0x33},
			CustodianPubKey: custodian,
			Timeout:         1_900_000_000,
			Chain:           1,
			Index:           0,
		},
		ParentXpubs: signerXpubs(4, csp.CoinTypeTestnet, &chaincfg.RegressionNetParams),
		Threshold:   3,
	}
}

// Decode(Encode(e)) is e, every field — including the two version 2 appended.
func TestRoundTripIsExact(t *testing.T) {
	e := layoutFixture()
	b, err := e.Encode()
	require.NoError(t, err)
	got, err := csp.Decode(b)
	require.NoError(t, err)
	require.Equal(t, e, got)
}

// The v2 layout vector: layoutFixture's encoding, one 32-byte word per line.
// It was produced by Foundry's `cast abi-encode` — an ABI encoder independent
// of go-ethereum's — over the signature
//
//	f((uint16,bytes32,bytes32,uint32,uint64,uint64,uint8,uint32,uint32,uint64,uint32,bytes,
//	   (bytes32,uint32,uint64,uint8,uint32,bool)[],address,(bytes32,bytes,uint64,uint8,uint32),
//	   bytes[],uint32))
//
// and the proposalHash by `cast keccak` over abi.encode(CSP_PROPOSAL, 31337,
// keccak256(encoding)). Any other implementation of the envelope — a machine,
// a verifier, a contract — can check itself against these three.
const (
	layoutVectorFile        = "testdata/envelope_v2_layout.hex"
	layoutVectorKeccak      = "0x2eb01b6a91e7d1d6bf3e46f0f52138778aa5c04a74d35c75f31f7d734ce00ec3"
	layoutVectorProposalHex = "0x08e499c7da0d754427162924ce747f054313a8e0792a09787787f77f6333d9f3"
)

// TestEnvelopeLayout pins the tuple: v1's fifteen fields in v1's order, then
// parentXpubs and threshold appended. A reorder, an insertion or a changed
// width fails here, naming the field, rather than as a proposal the other
// implementations never finalize.
func TestEnvelopeLayout(t *testing.T) {
	e := layoutFixture()
	enc, err := e.Encode()
	require.NoError(t, err)

	// abi.encode of one dynamic tuple opens with the tuple's offset; every
	// offset inside the tuple is relative to where the tuple starts.
	require.GreaterOrEqual(t, len(enc), 32)
	require.Equal(t, word(32), enc[:32], "the encoding does not open with the tuple offset 0x20")
	tuple := enc[32:]
	// Each helper takes the t of whichever (sub)test calls it: a require on
	// another test's t does not stop the caller.
	at := func(t *testing.T, off int) []byte {
		t.Helper()
		require.True(t, off >= 0 && off+32 <= len(tuple), "offset %d is outside the %d-byte tuple", off, len(tuple))
		return tuple[off : off+32]
	}
	head := func(t *testing.T, i int) []byte { t.Helper(); return at(t, 32*i) }
	tail := func(t *testing.T, i int) int {
		t.Helper()
		w := head(t, i)
		require.Equal(t, make([]byte, 24), w[:24], "head word %d is not an offset", i)
		return int(binary.BigEndian.Uint64(w[24:]))
	}

	type field struct {
		name  string
		check func(t *testing.T, w []byte, i int)
	}
	static := func(want []byte) func(*testing.T, []byte, int) {
		return func(t *testing.T, w []byte, _ int) { assert.Equal(t, want, w) }
	}
	fields := []field{
		{"version", static(word(2))},
		{"walletId", static(e.WalletID[:])},
		{"sourceId", static(e.SourceID[:])},
		{"accountIndex", static(word(7))},
		{"sequencePosition", static(word(1001))},
		{"eligibleGeneration", static(word(17))},
		{"instructionType", static(word(4))},
		{"paymentCount", static(word(6))},
		{"anchorIndex", static(word(5))},
		{"nonce", static(word(42))},
		{"attempt", static(word(1))},
		{"rawUnsignedTx", func(t *testing.T, _ []byte, i int) {
			off := tail(t, i)
			assert.Equal(t, word(uint64(len(e.RawUnsignedTx))), at(t, off), "length")
			assert.Equal(t, e.RawUnsignedTx, at(t, off+32)[:len(e.RawUnsignedTx)])
		}},
		{"inputs", func(t *testing.T, _ []byte, i int) {
			off := tail(t, i)
			assert.Equal(t, word(2), at(t, off), "count")
			// Each input is a static tuple of six words, encoded in place.
			assert.Equal(t, e.Inputs[0].Txid[:], at(t, off+32), "inputs[0].txid")
			assert.Equal(t, word(1), at(t, off+32+6*32+5*32), "inputs[1].isEscrow")
		}},
		{"proposerAddress", static(common.LeftPadBytes(e.ProposerAddress[:], 32))},
		{"escrow", func(t *testing.T, _ []byte, i int) {
			off := tail(t, i)
			assert.Equal(t, e.Escrow.PreimageHash[:], at(t, off), "escrow.preimageHash")
			assert.Equal(t, word(e.Escrow.Timeout), at(t, off+2*32), "escrow.timeout")
		}},
		{"parentXpubs", func(t *testing.T, _ []byte, i int) {
			off := tail(t, i)
			require.Equal(t, word(uint64(len(e.ParentXpubs))), at(t, off), "count")
			// Element offsets are relative to the array's first element.
			first := off + 32 + int(binary.BigEndian.Uint64(at(t, off+32)[24:]))
			assert.Equal(t, word(csp.XpubLen), at(t, first), "parentXpubs[0] length")
			assert.Equal(t, e.ParentXpubs[0][:32], at(t, first+32), "parentXpubs[0] bytes")
		}},
		{"threshold", static(word(3))},
	}

	for i, f := range fields {
		t.Run(fmt.Sprintf("%02d_%s", i+1, f.name), func(t *testing.T) { f.check(t, head(t, i), i) })
	}
	// The head is exactly one word per field: the first tail — rawUnsignedTx's,
	// since tails follow field order — begins where the head ends.
	assert.Equal(t, 32*len(fields), tail(t, 11), "the tuple head is not %d words: a field was added or dropped", len(fields))

	// And the whole encoding, byte for byte, against the independent vector.
	golden, err := os.ReadFile(layoutVectorFile)
	require.NoError(t, err)
	want, err := hex.DecodeString(strings.Join(strings.Fields(string(golden)), ""))
	require.NoError(t, err)
	require.Equal(t, want, enc, "the encoding moved off %s", layoutVectorFile)
	assert.Equal(t, layoutVectorKeccak, crypto.Keccak256Hash(enc).Hex())
	h, err := e.Hash(testChainID)
	require.NoError(t, err)
	assert.Equal(t, layoutVectorProposalHex, h.Hex())

	// The vector decodes back to the fixture.
	got, err := csp.Decode(want)
	require.NoError(t, err)
	require.Equal(t, e, got)
}

// word is v as one big-endian ABI word.
func word(v uint64) []byte {
	w := make([]byte, 32)
	binary.BigEndian.PutUint64(w[24:], v)
	return w
}
