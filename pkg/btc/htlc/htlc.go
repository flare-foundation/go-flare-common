package htlc

import (
	"bytes"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

// maxTimeout is BIP-65's effective upper bound for OP_CHECKLOCKTIMEVERIFY:
// CLTV compares against nLockTime as a script number, and script numbers are
// signed — anything above 2^31-1 wraps to negative and the OP fails. The
// proposer would never want this value (Bitcoin's chain reaches the
// timestamp 2^31-1 in 2038, and block height 2^31-1 is centuries away),
// but rejecting up front turns a confusing consensus reject at spend time
// into a clear validation error at script-build time.
const maxTimeout int64 = 0x7FFFFFFF

// minTimeout is the smallest UNIX timestamp accepted by BuildScript.
// BIP-65 interprets OP_CHECKLOCKTIMEVERIFY values < 500,000,000 as block
// heights, not timestamps. Core Vault escrows are always specified as
// absolute UNIX timestamps (mirroring XRPL's CancelAfter semantics), so
// a caller supplying a small integer likely made a time-vs-height mistake.
// Reject at build time to prevent a near-immediately-spendable escrow.
const minTimeout int64 = 500_000_000

// BuildScript constructs the witness script for a Core Vault escrow HTLC.
//
// Canonical miniscript ([BIP-379](https://github.com/bitcoin/bips/blob/master/bip-0379.mediawiki)) expression:
//
//	or_i(
//	    and_v(v:sha256(H), pk(custodian)),       // hash path
//	    and_v(v:after(T), multi(k, K1, ..., Kn)) // timeout path
//	)
//
// Compiled to Bitcoin Script (matches the bytes this function emits):
//
//	OP_IF
//	    OP_SIZE <32> OP_EQUALVERIFY     // preimage must be exactly 32 bytes
//	    OP_SHA256 <preimageHash> OP_EQUALVERIFY
//	    <custodianPubKey> OP_CHECKSIG
//	OP_ELSE
//	    <timeout> OP_CHECKLOCKTIMEVERIFY OP_VERIFY
//	    OP_k <sortedMultisigKeys...> OP_n OP_CHECKMULTISIG
//	OP_ENDIF
//
// The OP_SIZE preimage-length check (added in 2026-05 to align with the
// miniscript form) prevents non-32-byte preimage forgeries; the spec
// always uses 32-byte SHA256 preimages, so this is a tightening of
// invariants, not a behavior change.
//
// The miniscript form gives static-analysis guarantees (type-checking,
// satisfaction-weight bounds, no malleability). We emit the bytes by hand
// to avoid depending on an unaudited miniscript compiler in the critical
// path; equivalence to the miniscript-compiled output is asserted by
// TestHTLCMatchesMiniscript in this package.
//
// BIP-110/444 note: this script is consumed as a P2WSH **witness script**.
// BIP-110's `OP_IF`/`OP_NOTIF` ban applies only to Tapscript, so the
// OP_IF/OP_ELSE branching used here is exempt. If this HTLC is ever
// migrated to Tapscript, split the two spend paths into separate
// tapleaves instead of branching with OP_IF inside one leaf — see
// docs/shared/bip-tracking.md §"BIP-110 / BIP-444".
//
// preimageHash is SHA256(preimage), 32 bytes.
// custodianPubKey is the compressed public key of the custodian (33 bytes).
// timeout is the UNIX timestamp after which the timeout path becomes spendable.
// multisigKeys must be pre-sorted (BIP-67 lexicographic order), compressed (33 bytes each).
// threshold is k in the k-of-n multisig.
func BuildScript(
	preimageHash [32]byte,
	custodianPubKey []byte,
	timeout int64,
	multisigKeys [][]byte,
	threshold int,
) ([]byte, error) {
	if threshold < 1 || threshold > len(multisigKeys) {
		return nil, fmt.Errorf("threshold %d out of range for n=%d", threshold, len(multisigKeys))
	}
	if len(multisigKeys) > 20 {
		return nil, fmt.Errorf("n=%d exceeds OP_CHECKMULTISIG limit of 20", len(multisigKeys))
	}
	if timeout < minTimeout || timeout > maxTimeout {
		return nil, fmt.Errorf(
			"timeout %d out of range: CV escrows must use UNIX timestamps (≥ %d); "+
				"values < %d are interpreted as block heights by OP_CHECKLOCKTIMEVERIFY (BIP-65) "+
				"and would create a near-immediately-spendable escrow",
			timeout, minTimeout, minTimeout,
		)
	}
	// Pubkey validation: every key on either spend path must be a valid
	// compressed secp256k1 point. btcec.ParsePubKey catches:
	//   - wrong byte length (we still keep the explicit 33-byte length check
	//     for a clearer error in that common case),
	//   - bad leading byte (compressed keys must start with 0x02 or 0x03),
	//   - point not on the curve (a malformed x-coordinate or invalid y
	//     derivation).
	// Without this, BuildScript would happily emit a P2WSH whose
	// spending witness would fail at OP_CHECKSIG / OP_CHECKMULTISIG — a
	// silent unspendable-funds footgun discovered only when someone tries
	// to claim the HTLC.
	if len(custodianPubKey) != 33 {
		return nil, fmt.Errorf("custodian pubkey must be 33 bytes compressed, got %d", len(custodianPubKey))
	}
	if _, err := btcec.ParsePubKey(custodianPubKey); err != nil {
		return nil, fmt.Errorf("custodian pubkey invalid: %w", err)
	}
	for i, key := range multisigKeys {
		if len(key) != 33 {
			return nil, fmt.Errorf("multisigKeys[%d] must be 33 bytes compressed, got %d", i, len(key))
		}
		if _, err := btcec.ParsePubKey(key); err != nil {
			return nil, fmt.Errorf("multisigKeys[%d] invalid: %w", i, err)
		}
	}
	// Duplicate-key check: k-of-n with duplicates effectively reduces the
	// threshold (one signer's key, multiplied, can produce k "distinct"
	// signature slots). The spec assumes n distinct signers.
	for i := 0; i < len(multisigKeys); i++ {
		for j := i + 1; j < len(multisigKeys); j++ {
			if bytes.Equal(multisigKeys[i], multisigKeys[j]) {
				return nil, fmt.Errorf("multisigKeys contain duplicates at indices %d and %d", i, j)
			}
		}
	}

	builder := txscript.NewScriptBuilder()

	// IF branch: hash path (custodian emergency claim).
	// Canonical miniscript: and_v(v:sha256(H), pk(custodian))
	builder.AddOp(txscript.OP_IF)
	builder.AddOp(txscript.OP_SIZE) // miniscript v:sha256(H) starts with preimage-size check
	builder.AddInt64(32)            // 32-byte preimage
	builder.AddOp(txscript.OP_EQUALVERIFY)
	builder.AddOp(txscript.OP_SHA256)
	builder.AddData(preimageHash[:])
	builder.AddOp(txscript.OP_EQUALVERIFY)
	builder.AddData(custodianPubKey)
	builder.AddOp(txscript.OP_CHECKSIG)

	// ELSE branch: timeout path (k-of-n multisig after CLTV).
	// Canonical miniscript: and_v(v:after(T), multi(k, K1, ..., Kn))
	builder.AddOp(txscript.OP_ELSE)
	builder.AddInt64(timeout)
	builder.AddOp(txscript.OP_CHECKLOCKTIMEVERIFY)
	builder.AddOp(txscript.OP_VERIFY) // v:after(T) wraps after(T) with VERIFY, not DROP
	// Use AddInt64 so threshold > 16 (up to the n ≤ 20 cap above) emits a
	// minimal integer push. AddOp(OP_1-1+threshold) would silently emit
	// the wrong opcode (e.g., OP_NOTIF for 20) and break the script.
	builder.AddInt64(int64(threshold))
	for _, key := range multisigKeys {
		builder.AddData(key)
	}
	builder.AddInt64(int64(len(multisigKeys)))
	builder.AddOp(txscript.OP_CHECKMULTISIG)

	builder.AddOp(txscript.OP_ENDIF)

	return builder.Script()
}

// Address computes the P2WSH address for an HTLC witness script.
func Address(witnessScript []byte, params *chaincfg.Params) (btcutil.Address, error) {
	scriptHash := sha256.Sum256(witnessScript)
	return btcutil.NewAddressWitnessScriptHash(scriptHash[:], params)
}

// PreimageHash computes SHA256(preimage) for use in the HTLC script.
func PreimageHash(preimage [32]byte) [32]byte {
	return sha256.Sum256(preimage[:])
}

// HashPathWitness returns the canonical P2WSH witness stack for the
// HTLC hash path (IF branch — custodian emergency claim), per the normative
// encoding in docs/components/core-vault-btc.md §"Normative spend-path encodings":
//
//	[0] <custodianSig>  DER-encoded + sighash byte (caller appends 0x01 for SIGHASH_ALL)
//	[1] <preimage>      exactly 32 bytes
//	[2] {0x01}          MINIMALIF — selects the IF branch
//	[3] <witnessScript> full HTLC witness script
//
// sig must include the sighash byte (i.e. DER + 0x01).
// preimage must be exactly 32 bytes.
func HashPathWitness(sig []byte, preimage [32]byte, witnessScript []byte) [][]byte {
	return [][]byte{
		sig,
		preimage[:],
		{0x01}, // MINIMALIF: non-empty → IF branch
		witnessScript,
	}
}

// TimeoutPathWitness returns the canonical P2WSH witness stack for the
// HTLC timeout path (ELSE branch — vault reclaim after CLTV), per the normative
// encoding in docs/components/core-vault-btc.md §"Normative spend-path encodings":
//
//	[0]     <empty>         BIP-147 NULLDUMMY (OP_CHECKMULTISIG off-by-one)
//	[1..k]  <sig_i>         k DER-encoded signatures + sighash byte, BIP-67 order
//	[k+1]   <empty>         MINIMALIF — empty item selects the ELSE branch
//	[k+2]   <witnessScript> full HTLC witness script
//
// sigs must be in BIP-67 sorted-pubkey order and must each include the sighash byte.
// A {0x00} byte for the ELSE selector is non-minimal and causes relay rejection
// under the MINIMALIF standardness rule — the empty item is the correct encoding.
func TimeoutPathWitness(sigs [][]byte, witnessScript []byte) [][]byte {
	// [0] empty dummy (BIP-147 NULLDUMMY)
	witness := [][]byte{nil}
	// [1..k] signatures in BIP-67 sorted-pubkey order
	for _, sig := range sigs {
		witness = append(witness, sig)
	}
	// [k+1] empty ELSE selector (MINIMALIF)
	witness = append(witness, []byte{})
	// [k+2] witness script
	witness = append(witness, witnessScript)
	return witness
}
