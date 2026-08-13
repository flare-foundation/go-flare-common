// Package witness assembles k machines' partial signatures into a spendable
// Bitcoin transaction.
//
// This is the facilitator's job, and it is deliberately KEYLESS: everything
// here is public — the envelope the chain finalized, the wallet's published
// xpubs, and signatures the machines already produced. Assembly authorises
// nothing, which is why anyone may do it and why withholding is the only thing
// a facilitator can do wrong.
//
// Bitcoin signs PER INPUT, not once per transaction as XRPL does, so "k of n
// have signed" is a claim about every input separately: a batch is assemblable
// only when each of its inputs has k verified signatures. That is the whole
// shape difference from the XRPL arm of the same bot.
package witness

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"slices"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"

	"github.com/flare-foundation/go-flare-common/pkg/btc/address"
	"github.com/flare-foundation/go-flare-common/pkg/btc/csp"
	"github.com/flare-foundation/go-flare-common/pkg/btc/htlc"
)

// ErrNotEnoughSignatures means some input has fewer than k verified signatures.
// Callers wait and retry rather than treating it as a failure: signatures
// arrive from independent machines and arriving late is ordinary.
var ErrNotEnoughSignatures = errors.New("not enough verified signatures")

// InputSignature is one machine's contribution for one input.
type InputSignature struct {
	// Index is the input's position in the transaction.
	Index int
	// Signature is DER, WITHOUT the trailing sighash-type byte: whether a
	// signature is SIGHASH_ALL is a property of how it is used, so the byte is
	// appended at assembly rather than carried.
	Signature []byte
	// PubKey is the compressed key that produced it, which is how the signature
	// is placed in the script's BIP-67 order.
	PubKey []byte
}

// Answer is one machine's reply for a whole batch.
type Answer struct {
	// Machine identifies the source, for error messages only. Assembly does not
	// trust it: a signature counts because it verifies, not because of who sent
	// it, so a machine claiming another's identity gains nothing.
	Machine    int
	Signatures []InputSignature
}

// Wallet is what assembly needs to know about the signer set. All of it is
// public and all of it comes from the chain.
type Wallet struct {
	// ParentXpubs are the wallet's published keys, at the wallet level.
	ParentXpubs []string
	// Threshold is k in the k-of-n.
	Threshold int
	Params    *chaincfg.Params
}

// Assemble builds the witness for every input and returns the signed
// transaction and its txid.
//
// Every signature is VERIFIED against the sighash before it is used. An
// unverified one would produce a witness the mempool rejects with a message
// about the script, which sends whoever reads it to entirely the wrong place.
func Assemble(env csp.Envelope, w Wallet, answers []Answer) (*wire.MsgTx, [32]byte, error) {
	var txid [32]byte

	tx, err := csp.Tx(env)
	if err != nil {
		return nil, txid, fmt.Errorf("decoding the proposal's transaction: %w", err)
	}
	if len(tx.TxIn) != len(env.Inputs) {
		return nil, txid, fmt.Errorf("transaction has %d inputs but the envelope describes %d",
			len(tx.TxIn), len(env.Inputs))
	}
	if w.Threshold < 1 || w.Threshold > len(w.ParentXpubs) {
		return nil, txid, fmt.Errorf("threshold %d is not a k of %d", w.Threshold, len(w.ParentXpubs))
	}

	accountXpubs, err := address.DeriveAccountXpubs(w.ParentXpubs, env.AccountIndex, w.Params)
	if err != nil {
		return nil, txid, fmt.Errorf("deriving account xpubs: %w", err)
	}

	scripts, sorted, err := inputScripts(env, w, accountXpubs)
	if err != nil {
		return nil, txid, err
	}

	// Prevouts for the BIP-143 sighash. Every input is one of this wallet's own
	// P2WSH outputs, so each scriptPubKey comes from its own derivation rather
	// than from a lookup.
	prevouts := txscript.NewMultiPrevOutFetcher(nil)
	for i, in := range env.Inputs {
		pk, perr := P2WSH(scripts[i])
		if perr != nil {
			return nil, txid, fmt.Errorf("scriptPubKey for input %d: %w", i, perr)
		}
		prevouts.AddPrevOut(tx.TxIn[i].PreviousOutPoint, wire.NewTxOut(int64(in.ValueSat), pk))
	}
	sigHashes := txscript.NewTxSigHashes(tx, prevouts)

	for i, in := range env.Inputs {
		sighash, herr := txscript.CalcWitnessSigHash(
			scripts[i], sigHashes, txscript.SigHashAll, tx, i, int64(in.ValueSat))
		if herr != nil {
			return nil, txid, fmt.Errorf("sighash for input %d: %w", i, herr)
		}

		bySlot, verr := verifiedSlots(i, sighash, sorted[i], answers)
		if verr != nil {
			return nil, txid, verr
		}
		if len(bySlot) < w.Threshold {
			return nil, txid, fmt.Errorf("%w: input %d has %d, needs %d",
				ErrNotEnoughSignatures, i, len(bySlot), w.Threshold)
		}

		tx.TxIn[i].Witness = stack(bySlot, w.Threshold, scripts[i], in.IsEscrow)
	}

	return tx, [32]byte(tx.TxHash()), nil
}

// inputScripts is the witness script and the BIP-67 key order for each input.
func inputScripts(env csp.Envelope, w Wallet, accountXpubs []string) ([][]byte, [][][]byte, error) {
	scripts := make([][]byte, len(env.Inputs))
	sorted := make([][][]byte, len(env.Inputs))

	for i, in := range env.Inputs {
		if in.IsEscrow {
			// The escrow's TIMEOUT branch. Its quorum is this wallet's own
			// k-of-n at the escrow's coordinates, so the key ORDER — which
			// decides where each signature sits — comes from the same
			// derivation an ordinary input uses. Only the script differs.
			script, _, err := htlc.ForWallet(w.ParentXpubs, env.AccountIndex, w.Threshold, htlc.Terms{
				PreimageHash:    env.Escrow.PreimageHash,
				CustodianPubKey: env.Escrow.CustodianPubKey,
				Timeout:         int64(env.Escrow.Timeout),
				Chain:           address.Chain(env.Escrow.Chain),
				Index:           env.Escrow.Index,
			}, w.Params)
			if err != nil {
				return nil, nil, fmt.Errorf("rebuilding the escrow script for input %d: %w", i, err)
			}
			_, _, pubs, err := address.Derive(
				accountXpubs, w.Threshold, address.Chain(env.Escrow.Chain), env.Escrow.Index, w.Params)
			if err != nil {
				return nil, nil, fmt.Errorf("deriving the escrow's keys for input %d: %w", i, err)
			}
			scripts[i], sorted[i] = script, pubs
			continue
		}

		_, script, pubs, err := address.Derive(
			accountXpubs, w.Threshold, address.Chain(in.Chain), in.Index, w.Params)
		if err != nil {
			return nil, nil, fmt.Errorf("deriving input %d (chain %d, index %d): %w", i, in.Chain, in.Index, err)
		}
		scripts[i], sorted[i] = script, pubs
	}
	return scripts, sorted, nil
}

// verifiedSlots maps each signature to WHERE its key sits in the script.
//
// OP_CHECKMULTISIG walks signatures and public keys in lockstep, so a witness
// whose signatures are in any other order fails even though every signature in
// it is individually valid. The slot, not the arrival order, is what matters.
func verifiedSlots(i int, sighash []byte, keys [][]byte, answers []Answer) (map[int][]byte, error) {
	bySlot := map[int][]byte{}

	for _, m := range answers {
		for _, s := range m.Signatures {
			if s.Index != i {
				continue
			}

			slot := slices.IndexFunc(keys, func(k []byte) bool { return bytes.Equal(k, s.PubKey) })
			if slot < 0 {
				return nil, fmt.Errorf("machine %d input %d: signing key %x is not in the wallet's script",
					m.Machine, i, s.PubKey)
			}

			pub, err := btcec.ParsePubKey(s.PubKey)
			if err != nil {
				return nil, fmt.Errorf("machine %d input %d: unparseable pubkey: %w", m.Machine, i, err)
			}
			sig, err := ecdsa.ParseDERSignature(s.Signature)
			if err != nil {
				return nil, fmt.Errorf("machine %d input %d: unparseable signature: %w", m.Machine, i, err)
			}
			if !sig.Verify(sighash, pub) {
				return nil, fmt.Errorf("machine %d input %d: signature does not verify against the sighash",
					m.Machine, i)
			}

			// SigHashAll is a property of how the signature is USED, so the
			// machines return DER without it and it is appended here.
			bySlot[slot] = append(append([]byte{}, s.Signature...), byte(txscript.SigHashAll))
		}
	}
	return bySlot, nil
}

// stack builds one input's witness: the CHECKMULTISIG dummy, exactly k
// signatures in key order, an OP_ELSE selector for an escrow, then the script.
func stack(bySlot map[int][]byte, k int, script []byte, isEscrow bool) wire.TxWitness {
	slots := make([]int, 0, len(bySlot))
	for s := range bySlot {
		slots = append(slots, s)
	}
	slices.Sort(slots)
	slots = slots[:k] // exactly k: a witness with extras fails CHECKMULTISIG

	// The leading empty item is CHECKMULTISIG's off-by-one, which consensus
	// preserves.
	w := wire.TxWitness{nil}
	for _, s := range slots {
		w = append(w, bySlot[s])
	}
	if isEscrow {
		// The escrow script branches on OP_IF, so the witness must SELECT a
		// branch: an empty item is false and takes OP_ELSE, the timeout path.
		// Without it the interpreter runs the hash path, demands a preimage the
		// wallet does not have, and the spend fails with signatures that were
		// all perfectly valid.
		w = append(w, nil)
	}
	return append(w, script)
}

// P2WSH is the witness-v0 scriptPubKey committing to a witness script.
func P2WSH(witnessScript []byte) ([]byte, error) {
	h := sha256.Sum256(witnessScript)
	return txscript.NewScriptBuilder().AddOp(txscript.OP_0).AddData(h[:]).Script()
}

// Serialize is the wire form to hand to sendrawtransaction.
func Serialize(tx *wire.MsgTx) ([]byte, error) {
	var buf bytes.Buffer
	if err := tx.Serialize(&buf); err != nil {
		return nil, fmt.Errorf("serializing the signed transaction: %w", err)
	}
	return buf.Bytes(), nil
}
