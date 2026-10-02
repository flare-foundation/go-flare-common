package csp_test

import (
	"bytes"
	"encoding/binary"
	"errors"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/btc/csp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func confirmedSample() csp.ConfirmedAttempt {
	return csp.ConfirmedAttempt{
		WalletID:           [32]byte{0x01, 0x02},
		SourceID:           [32]byte{'B', 'T', 'C'},
		AccountIndex:       3,
		SequencePosition:   7,
		EligibleGeneration: 11,
		Attempt:            2,
		ConfirmedAttempt:   0,
		Txid:               [32]byte{0xaa, 0xbb},
		ProposerAddress:    common.HexToAddress("0x00000000000000000000000000000000000A11CE"),
	}
}

func TestConfirmedAttemptRoundTrip(t *testing.T) {
	c := confirmedSample()
	b, err := c.Encode()
	require.NoError(t, err)
	require.Len(t, b, csp.ConfirmedAttemptLen)
	assert.Equal(t, csp.ConfirmedAttemptTag[:], b[:32], "the tag opens the encoding")
	got, err := csp.DecodeConfirmedAttempt(b)
	require.NoError(t, err)
	assert.Equal(t, c, got)
}

// The layout is ten static words in declaration order; pinned word by word so
// that a reordering is a test failure rather than a silent change of every
// package hash.
func TestConfirmedAttemptLayout(t *testing.T) {
	c := confirmedSample()
	b, err := c.Encode()
	require.NoError(t, err)
	w := func(i int) []byte { return b[i*32 : (i+1)*32] }
	assert.Equal(t, c.WalletID[:], w(1))
	assert.Equal(t, c.SourceID[:], w(2))
	assert.Equal(t, uint64(3), binary.BigEndian.Uint64(w(3)[24:]))
	assert.Equal(t, uint64(7), binary.BigEndian.Uint64(w(4)[24:]))
	assert.Equal(t, uint64(11), binary.BigEndian.Uint64(w(5)[24:]))
	assert.Equal(t, uint64(2), binary.BigEndian.Uint64(w(6)[24:]))
	assert.Equal(t, uint64(0), binary.BigEndian.Uint64(w(7)[24:]))
	assert.Equal(t, c.Txid[:], w(8))
	assert.Equal(t, c.ProposerAddress.Bytes(), w(9)[12:])
}

func TestConfirmedAttemptEveryFieldChangesTheHash(t *testing.T) {
	base, err := confirmedSample().Hash(testChainID)
	require.NoError(t, err)
	mutations := map[string]func(*csp.ConfirmedAttempt){
		"walletId":           func(c *csp.ConfirmedAttempt) { c.WalletID[31] ^= 1 },
		"sourceId":           func(c *csp.ConfirmedAttempt) { c.SourceID[31] ^= 1 },
		"accountIndex":       func(c *csp.ConfirmedAttempt) { c.AccountIndex++ },
		"sequencePosition":   func(c *csp.ConfirmedAttempt) { c.SequencePosition++ },
		"eligibleGeneration": func(c *csp.ConfirmedAttempt) { c.EligibleGeneration++ },
		"attempt":            func(c *csp.ConfirmedAttempt) { c.Attempt++ },
		"confirmedAttempt":   func(c *csp.ConfirmedAttempt) { c.ConfirmedAttempt++ },
		"txid":               func(c *csp.ConfirmedAttempt) { c.Txid[0] ^= 1 },
		"proposerAddress":    func(c *csp.ConfirmedAttempt) { c.ProposerAddress[0] ^= 1 },
	}
	for name, mutate := range mutations {
		c := confirmedSample()
		mutate(&c)
		h, err := c.Hash(testChainID)
		require.NoError(t, err, name)
		assert.NotEqual(t, base, h, "%s is not committed", name)
	}
	other, err := confirmedSample().Hash(testChainID + 1)
	require.NoError(t, err)
	assert.NotEqual(t, base, other, "the chain id separates the hash")
}

func TestConfirmedAttemptValidate(t *testing.T) {
	cases := map[string]func(*csp.ConfirmedAttempt){
		"attempt 0 is never open behind a reissue": func(c *csp.ConfirmedAttempt) { c.Attempt, c.ConfirmedAttempt = 0, 0 },
		"the open attempt itself":                  func(c *csp.ConfirmedAttempt) { c.ConfirmedAttempt = c.Attempt },
		"an attempt after the open one":            func(c *csp.ConfirmedAttempt) { c.ConfirmedAttempt = c.Attempt + 1 },
		"a zero txid":                              func(c *csp.ConfirmedAttempt) { c.Txid = [32]byte{} },
		"no proposer":                              func(c *csp.ConfirmedAttempt) { c.ProposerAddress = common.Address{} },
	}
	for name, mutate := range cases {
		c := confirmedSample()
		mutate(&c)
		_, err := c.Encode()
		assert.Error(t, err, name)
	}
}

// Anything the open DAL serves is attacker input: Decode re-encodes, so a
// value with garbage in a narrow field's padding, a wrong length or a missing
// tag is refused, and so is an invalid proposal.
func TestDecodeConfirmedAttemptRefusals(t *testing.T) {
	good, err := confirmedSample().Encode()
	require.NoError(t, err)
	mutate := func(f func([]byte)) []byte {
		b := bytes.Clone(good)
		f(b)
		return b
	}
	cases := map[string][]byte{
		"empty":                 nil,
		"short":                 good[:len(good)-1],
		"trailing byte":         append(bytes.Clone(good), 0),
		"no tag":                mutate(func(b []byte) { b[0] ^= 1 }),
		"accountIndex padding":  mutate(func(b []byte) { b[3*32] = 1 }),
		"attempt padding":       mutate(func(b []byte) { b[6*32+27] = 1 }),
		"address padding":       mutate(func(b []byte) { b[9*32] = 1 }),
		"confirmed is the open": mutate(func(b []byte) { b[7*32+31] = 2 }),
		"confirmed after open":  mutate(func(b []byte) { b[7*32+31] = 3 }),
		"open attempt zero":     mutate(func(b []byte) { b[6*32+31] = 0 }),
		"zero txid":             mutate(func(b []byte) { copy(b[8*32:9*32], make([]byte, 32)) }),
	}
	for name, b := range cases {
		_, err := csp.DecodeConfirmedAttempt(b)
		assert.Error(t, err, name)
	}
}

func signConfirmed(t *testing.T, c csp.ConfirmedAttempt) []byte {
	t.Helper()
	key, err := crypto.HexToECDSA("4c0883a69102937d6231471b5dbb6204fe5129617082792ae468d01a3f362318")
	require.NoError(t, err)
	c.ProposerAddress = crypto.PubkeyToAddress(key.PublicKey)
	h, err := c.Hash(testChainID)
	require.NoError(t, err)
	sig, err := crypto.Sign(accounts.TextHash(h[:]), key)
	require.NoError(t, err)
	pkg, err := c.Package(sig)
	require.NoError(t, err)
	return pkg
}

func TestConfirmedAttemptPackage(t *testing.T) {
	pkg := signConfirmed(t, confirmedSample())
	require.True(t, csp.IsConfirmedAttemptPackage(pkg))

	c, sig, err := csp.SplitConfirmedAttemptPackage(pkg)
	require.NoError(t, err)
	signer, err := c.Signer(sig, testChainID)
	require.NoError(t, err)
	assert.Equal(t, c.ProposerAddress, signer)

	_, err = c.Signer(sig, testChainID+1)
	assert.Error(t, err, "a signature on another chain names another signer")
	c.ProposerAddress[0] ^= 1
	_, err = c.Signer(sig, testChainID)
	assert.Error(t, err, "a package naming someone else as its proposer")
}

// The two encodings are disjoint: no transaction consumer can read a
// confirmed attempt as an envelope, and SplitPackage names it rather than
// reporting it malformed.
func TestConfirmedAttemptIsNoEnvelope(t *testing.T) {
	pkg := signConfirmed(t, confirmedSample())
	_, _, err := csp.SplitPackage(pkg)
	assert.True(t, errors.Is(err, csp.ErrConfirmedAttempt), "SplitPackage names a confirmed attempt: %v", err)
	_, err = csp.Decode(pkg[:csp.ConfirmedAttemptLen])
	assert.Error(t, err)

	env, err := sample().Encode()
	require.NoError(t, err)
	assert.False(t, csp.IsConfirmedAttemptPackage(append(env, make([]byte, csp.ProposerSigLen)...)))
	_, err = csp.DecodeConfirmedAttempt(env)
	assert.Error(t, err)
	_, _, err = csp.SplitConfirmedAttemptPackage(append(env, make([]byte, csp.ProposerSigLen)...))
	assert.Error(t, err)
}

// The commitment is abi.encode(uint32 attempt, bytes32 txid): two words, the
// CONTRACT's order, never a BtcProposalCommitment's five.
func TestConfirmedAttemptCommitmentPreimage(t *testing.T) {
	c := csp.BtcConfirmedAttemptCommitment{Attempt: 5, Txid: [32]byte{0xcc}}
	pre := make([]byte, 64)
	pre[31] = 5
	pre[32] = 0xcc
	assert.Equal(t, [32]byte(crypto.Keccak256Hash(pre)), c.CommitmentHash())
	assert.Equal(t, confirmedSample().Commitment(), csp.BtcConfirmedAttemptCommitment{Attempt: 0, Txid: [32]byte{0xaa, 0xbb}})
}
