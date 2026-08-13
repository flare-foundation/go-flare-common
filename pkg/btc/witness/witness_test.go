package witness_test

import (
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/btc/address"
	"github.com/flare-foundation/go-flare-common/pkg/btc/csp"
	"github.com/flare-foundation/go-flare-common/pkg/btc/witness"
)

var params = &chaincfg.RegressionNetParams

// wallet is a 2-of-3 whose private keys the test holds, so it can play all
// three machines.
type wallet struct {
	masters []*hdkeychain.ExtendedKey
	parents []string
	k       int
}

// newWallet builds an n-of-k wallet whose seeds depend on the SALT as well as
// the key index, so two wallets in one test are actually different. Without
// that they are byte-identical, and a test about "a key from another wallet"
// silently becomes a test about the same key twice.
func newWallet(t *testing.T, n, k int, salt byte) wallet {
	t.Helper()
	w := wallet{k: k}
	for i := range n {
		seed := make([]byte, 32)
		for j := range seed {
			seed[j] = byte(i*7+j+1) ^ salt
		}
		m, err := hdkeychain.NewMaster(seed, params)
		require.NoError(t, err)
		pub, err := m.Neuter()
		require.NoError(t, err)
		w.masters = append(w.masters, m)
		w.parents = append(w.parents, pub.String())
	}
	return w
}

// signer derives machine i's private key for one address.
func (w wallet) signer(t *testing.T, machine int, accountIndex uint32, chain address.Chain, index uint32) *btcec.PrivateKey {
	t.Helper()
	acct, err := w.masters[machine].Derive(accountIndex)
	require.NoError(t, err)
	ch, err := acct.Derive(uint32(chain))
	require.NoError(t, err)
	leaf, err := ch.Derive(index)
	require.NoError(t, err)
	priv, err := leaf.ECPrivKey()
	require.NoError(t, err)
	return priv
}

// spendable builds a one-input, one-output transaction spending the wallet's
// address at (chain, index), and the envelope that describes it.
func spendable(t *testing.T, w wallet, valueSat int64) (csp.Envelope, []byte) {
	t.Helper()
	acct, err := address.DeriveAccountXpubs(w.parents, 0, params)
	require.NoError(t, err)
	_, script, _, err := address.Derive(acct, w.k, address.External, 0, params)
	require.NoError(t, err)

	prev := chainhash.Hash{0x11}
	tx := wire.NewMsgTx(3)
	tx.AddTxIn(wire.NewTxIn(wire.NewOutPoint(&prev, 0), nil, nil))
	tx.AddTxOut(wire.NewTxOut(valueSat-1000, []byte{0x51}))

	var buf []byte
	buf, err = serializeNoWitness(tx)
	require.NoError(t, err)

	return csp.Envelope{
		Version: 1, AccountIndex: 0, RawUnsignedTx: buf,
		Inputs: []csp.Input{{
			Txid: [32]byte(prev), Vout: 0, ValueSat: uint64(valueSat),
			Chain: uint8(address.External), Index: 0,
		}},
	}, script
}

func serializeNoWitness(tx *wire.MsgTx) ([]byte, error) {
	var b []byte
	w := &sliceWriter{&b}
	if err := tx.SerializeNoWitness(w); err != nil {
		return nil, err
	}
	return b, nil
}

type sliceWriter struct{ b *[]byte }

func (s *sliceWriter) Write(p []byte) (int, error) { *s.b = append(*s.b, p...); return len(p), nil }

// signAll produces machine i's signature for every input of the transaction.
func signAll(t *testing.T, w wallet, env csp.Envelope, script []byte, machine int) witness.Answer {
	t.Helper()
	tx, err := csp.Tx(env)
	require.NoError(t, err)

	pk, err := witness.P2WSH(script)
	require.NoError(t, err)
	prevouts := txscript.NewMultiPrevOutFetcher(nil)
	for i, in := range env.Inputs {
		prevouts.AddPrevOut(tx.TxIn[i].PreviousOutPoint, wire.NewTxOut(int64(in.ValueSat), pk))
	}
	hashes := txscript.NewTxSigHashes(tx, prevouts)

	ans := witness.Answer{Machine: machine}
	for i, in := range env.Inputs {
		sighash, herr := txscript.CalcWitnessSigHash(script, hashes, txscript.SigHashAll, tx, i, int64(in.ValueSat))
		require.NoError(t, herr)
		priv := w.signer(t, machine, env.AccountIndex, address.Chain(in.Chain), in.Index)
		sig := ecdsaSign(t, priv, sighash)
		ans.Signatures = append(ans.Signatures, witness.InputSignature{
			Index: i, Signature: sig, PubKey: priv.PubKey().SerializeCompressed(),
		})
	}
	return ans
}

func ecdsaSign(t *testing.T, priv *btcec.PrivateKey, hash []byte) []byte {
	t.Helper()
	sig, err := signDER(priv, hash)
	require.NoError(t, err)
	return sig
}

// TestAssembledWitnessSatisfiesTheScript runs Bitcoin's own script interpreter
// over the result.
//
// The assertion that matters: not that assembly produced bytes, but that the
// consensus engine accepts them. A witness with the signatures in the wrong
// order, or missing CHECKMULTISIG's dummy item, is well-formed and still
// unspendable.
func TestAssembledWitnessSatisfiesTheScript(t *testing.T) {
	w := newWallet(t, 3, 2, 0x00)
	env, script := spendable(t, w, 100_000)

	// Two of three machines, which is exactly the threshold.
	answers := []witness.Answer{
		signAll(t, w, env, script, 0),
		signAll(t, w, env, script, 2),
	}

	tx, txid, err := witness.Assemble(env, witness.Wallet{
		ParentXpubs: w.parents, Threshold: 2, Params: params,
	}, answers)
	require.NoError(t, err)
	assert.NotEqual(t, [32]byte{}, txid)

	pk, err := witness.P2WSH(script)
	require.NoError(t, err)
	prevouts := txscript.NewMultiPrevOutFetcher(nil)
	prevouts.AddPrevOut(tx.TxIn[0].PreviousOutPoint, wire.NewTxOut(100_000, pk))

	vm, err := txscript.NewEngine(pk, tx, 0,
		txscript.StandardVerifyFlags, nil, txscript.NewTxSigHashes(tx, prevouts), 100_000, prevouts)
	require.NoError(t, err)
	require.NoError(t, vm.Execute(), "the script interpreter must accept the assembled witness")
}

// TestBelowThresholdIsNotAssembled: one signature of a 2-of-3 is not a batch
// that can be broadcast, and saying so is the facilitator's only real state.
func TestBelowThresholdIsNotAssembled(t *testing.T) {
	w := newWallet(t, 3, 2, 0x00)
	env, script := spendable(t, w, 100_000)

	_, _, err := witness.Assemble(env, witness.Wallet{
		ParentXpubs: w.parents, Threshold: 2, Params: params,
	}, []witness.Answer{signAll(t, w, env, script, 1)})

	require.ErrorIs(t, err, witness.ErrNotEnoughSignatures)
}

// TestExtraSignaturesAreTrimmed: three of three is more than CHECKMULTISIG
// wants, and a witness carrying all of them fails. Exactly k are used.
func TestExtraSignaturesAreTrimmed(t *testing.T) {
	w := newWallet(t, 3, 2, 0x00)
	env, script := spendable(t, w, 100_000)

	tx, _, err := witness.Assemble(env, witness.Wallet{
		ParentXpubs: w.parents, Threshold: 2, Params: params,
	}, []witness.Answer{
		signAll(t, w, env, script, 0),
		signAll(t, w, env, script, 1),
		signAll(t, w, env, script, 2),
	})
	require.NoError(t, err)

	// dummy + 2 signatures + script.
	assert.Len(t, tx.TxIn[0].Witness, 4)

	pk, err := witness.P2WSH(script)
	require.NoError(t, err)
	prevouts := txscript.NewMultiPrevOutFetcher(nil)
	prevouts.AddPrevOut(tx.TxIn[0].PreviousOutPoint, wire.NewTxOut(100_000, pk))
	vm, err := txscript.NewEngine(pk, tx, 0,
		txscript.StandardVerifyFlags, nil, txscript.NewTxSigHashes(tx, prevouts), 100_000, prevouts)
	require.NoError(t, err)
	require.NoError(t, vm.Execute())
}

// TestForeignSignatureIsRejected: a signature from a key outside the wallet's
// script has no slot to sit in, and assembly must say so rather than drop it
// silently — a dropped signature reads later as a machine that never answered.
func TestForeignSignatureIsRejected(t *testing.T) {
	w := newWallet(t, 3, 2, 0x00)
	other := newWallet(t, 3, 2, 0xA5)
	env, script := spendable(t, w, 100_000)

	answers := []witness.Answer{
		signAll(t, w, env, script, 0),
		signAll(t, other, env, script, 0),
	}
	_, _, err := witness.Assemble(env, witness.Wallet{
		ParentXpubs: w.parents, Threshold: 2, Params: params,
	}, answers)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not in the wallet's script")
}

// TestCorruptSignatureIsRejected: assembly verifies before it assembles, so a
// signature that does not match the sighash is refused here rather than by the
// mempool, whose message would be about the script and send the reader to
// entirely the wrong place.
func TestCorruptSignatureIsRejected(t *testing.T) {
	w := newWallet(t, 3, 2, 0x00)
	env, script := spendable(t, w, 100_000)

	bad := signAll(t, w, env, script, 1)
	// Flip a byte inside the DER body, keeping it parseable.
	bad.Signatures[0].Signature[len(bad.Signatures[0].Signature)-1] ^= 0x01

	_, _, err := witness.Assemble(env, witness.Wallet{
		ParentXpubs: w.parents, Threshold: 2, Params: params,
	}, []witness.Answer{signAll(t, w, env, script, 0), bad})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "does not verify")
}
