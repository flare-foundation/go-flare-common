// Package btcaddr provides reusable Bitcoin address derivation and script
// construction utilities for the Flare BTC infrastructure.
//
// It implements the spec's P2WSH k-of-n multisig address derivation from
// docs/shared/wallet-and-address-derivation.md, combining BIP-32 child
// derivation, BIP-67 lex sorting, BIP-141 P2WSH witness script hashing,
// and BIP-173 Bech32 encoding.
//
// Derive is path-agnostic: it consumes account-level xpubs and the leaf
// chain/index, so the same code works for the spec's BIP-87 path
// (m/87'/coin'/account, recommended for multisig) and the legacy BIP-44 path.
//
// Account derivation is NON-HARDENED (2026-06-22, aligns with the deployed
// flare-smart-contracts-v2 UTXO model — IPMWMultisigUtxoConfigured /
// docs/specs/FCC/Payments.md). The TEEs publish ONE wallet-level (parent) xpub
// at m/87'/coin'; the account-level xpub is its non-hardened child at
// accountIndex (DeriveAccountXpubs), and every leaf address is then derived
// from that account-level xpub via Derive. A single published parent xpub
// therefore serves all accountIndex values with no per-account attestation.
// The hardened-account boundary was deliberately dropped: any child-key leak
// already forces a full TEE rotation, so the boundary bought no real
// protection (see docs/inputs/2026-06-22-non-hardened-account-derivation.md).
package address

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

// Chain selects between the external (receive) and internal (change)
// chains when deriving a leaf address from an account-level xpub. The
// convention is shared by BIP-44 and BIP-87: external = 0, internal = 1.
type Chain uint32

const (
	External Chain = 0
	Internal Chain = 1
)

// Derive produces the deterministic P2WSH k-of-n multisig address for the
// given chain and leaf index, derived from the supplied account-level xpubs.
//
// Each xpub must be at the account level — i.e. the non-hardened child of the
// published wallet-level (parent) xpub at accountIndex (m/87'/coin_type'/account,
// account NON-hardened; see DeriveAccountXpubs and the package doc). The older
// BIP-44 leaf shape is also accepted. Callers that hold the published parent
// xpubs derive the account-level xpubs with DeriveAccountXpubs first.
// The procedure applies CKDpub twice -- once for the chain (0 external,
// 1 internal), once for the leaf index -- then BIP-67-sorts the resulting
// compressed pubkeys, builds an OP_CHECKMULTISIG script, and encodes the
// SHA256 of the script as a P2WSH address (witness v0).
//
// Returns the address, the witness script (useful for inputs), and the
// sorted compressed pubkeys (useful for debugging or signing).
func Derive(
	xpubs []string,
	threshold int,
	chain Chain,
	leafIndex uint32,
	params *chaincfg.Params,
) (btcutil.Address, []byte, [][]byte, error) {
	if threshold < 1 || threshold > len(xpubs) {
		return nil, nil, nil, fmt.Errorf("threshold %d out of range for n=%d", threshold, len(xpubs))
	}
	if len(xpubs) > 20 {
		return nil, nil, nil, fmt.Errorf("n=%d exceeds OP_CHECKMULTISIG limit of 20", len(xpubs))
	}

	// Step 1: derive each leaf pubkey via two CKDpub steps.
	pubKeys := make([][]byte, len(xpubs))
	for i, xs := range xpubs {
		acct, err := hdkeychain.NewKeyFromString(xs)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("parse xpub[%d]: %w", i, err)
		}
		if !acct.IsForNet(params) {
			return nil, nil, nil, fmt.Errorf("xpub[%d] version byte does not match network %q", i, params.Name)
		}
		if acct.IsPrivate() {
			return nil, nil, nil, fmt.Errorf("xpub[%d] is a private extended key; public keys only", i)
		}
		chainKey, err := acct.Derive(uint32(chain))
		if err != nil {
			return nil, nil, nil, fmt.Errorf("derive chain[%d]: %w", i, err)
		}
		leaf, err := chainKey.Derive(leafIndex)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("derive leaf[%d]: %w", i, err)
		}
		pk, err := leaf.ECPubKey()
		if err != nil {
			return nil, nil, nil, fmt.Errorf("ec pubkey[%d]: %w", i, err)
		}
		pubKeys[i] = pk.SerializeCompressed()
	}

	// Reject duplicate derived pubkeys: identical keys collapse signer
	// independence (a single holder can satisfy multiple CHECKMULTISIG slots),
	// so a k-of-n with repeats is not a genuine k-of-n. This guards the script
	// directly, complementing the parent-xpub uniqueness check in ValidateV1.
	seen := make(map[string]struct{}, len(pubKeys))
	for i, pk := range pubKeys {
		if _, dup := seen[string(pk)]; dup {
			return nil, nil, nil, fmt.Errorf("duplicate public key among signers (xpub[%d]); each signer must be distinct", i)
		}
		seen[string(pk)] = struct{}{}
	}

	// Step 2: BIP-67 lex sort.
	sort.Slice(pubKeys, func(i, j int) bool {
		return bytes.Compare(pubKeys[i], pubKeys[j]) < 0
	})

	// Step 3: build OP_CHECKMULTISIG witness script.
	//   OP_k <sorted[0]> ... <sorted[n-1]> OP_n OP_CHECKMULTISIG
	//
	// AddInt64 picks OP_1..OP_16 for k/n in 1..16 and a minimal scriptNum
	// push for 17..20 — both canonical for CHECKMULTISIG. A naïve
	// byte(OP_1-1+x) encoding would emit OP_NOP (0x61) for x=17 and break
	// at n>16 (see open-question #14: n up to 20 is supported).
	builder := txscript.NewScriptBuilder()
	// OP_1..OP_16 cover threshold in 1..16; AddInt64 falls back to a minimal
	// integer push for values above 16 (up to the n ≤ 20 cap checked above).
	// AddOp(OP_1-1+threshold) for threshold > 16 silently emits the wrong
	// opcode (e.g., OP_NOTIF for 20) and produces an unspendable script.
	builder.AddInt64(int64(threshold))
	for _, pk := range pubKeys {
		builder.AddData(pk)
	}
	builder.AddInt64(int64(len(pubKeys)))
	builder.AddOp(txscript.OP_CHECKMULTISIG)
	witnessScript, err := builder.Script()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("build script: %w", err)
	}

	// Step 4: SHA256 of script, encoded as P2WSH (witness v0).
	scriptHash := sha256.Sum256(witnessScript)
	addr, err := btcutil.NewAddressWitnessScriptHash(scriptHash[:], params)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("encode address: %w", err)
	}
	return addr, witnessScript, pubKeys, nil
}

// DeriveAccountXpubs maps the published wallet-level (parent) xpubs to their
// account-level xpubs by applying a single NON-HARDENED CKDpub step at
// accountIndex (2026-06-22, aligns with the deployed flare-smart-contracts-v2
// UTXO model — the account-level keys are the non-hardened children of the
// parent xpubs at accountIndex, per IPMWMultisigUtxoConfigured /
// docs/specs/FCC/Payments.md).
//
// Each input must be a wallet-level (parent) xpub, i.e. one level above the
// account (BIP-87 m/87'/coin'); the returned strings are the account-level
// xpubs (m/87'/coin'/accountIndex) that Derive consumes for leaf addresses.
// accountIndex must be in the BIP-32 non-hardened range (< 2^31) — the account
// level is no longer hardened, so it is derivable from the public parent xpub
// alone, which is exactly what lets one published parent serve every account.
func DeriveAccountXpubs(
	parentXpubs []string,
	accountIndex uint32,
	params *chaincfg.Params,
) ([]string, error) {
	if accountIndex >= 1<<31 {
		return nil, fmt.Errorf("accountIndex %d is in the BIP-32 hardened range (must be < 2^31 for non-hardened account derivation)", accountIndex)
	}
	out := make([]string, len(parentXpubs))
	for i, xs := range parentXpubs {
		parent, err := hdkeychain.NewKeyFromString(xs)
		if err != nil {
			return nil, fmt.Errorf("parse parent xpub[%d]: %w", i, err)
		}
		if !parent.IsForNet(params) {
			return nil, fmt.Errorf("parent xpub[%d] version byte does not match network %q", i, params.Name)
		}
		if parent.IsPrivate() {
			return nil, fmt.Errorf("parent xpub[%d] is a private extended key; public keys only", i)
		}
		acct, err := parent.Derive(accountIndex)
		if err != nil {
			return nil, fmt.Errorf("derive account xpub[%d] at index %d: %w", i, accountIndex, err)
		}
		out[i] = acct.String()
	}
	return out, nil
}

// ValidateDerivationPath verifies that path is `<prefix>/<chain>/<leafIndex>`
// where chain is "0" (external) or "1" (internal/change) and leafIndex is a
// canonical decimal in [0, 2^31). The prefix is the account-level shape
// validated separately at attestation time (see Attestation.ValidateV1); this
// helper only checks the per-UTXO leaf segment.
//
// If prefix is empty, ValidateDerivationPath returns nil — callers in legacy
// code paths or tests that do not yet have an attestation in hand can opt
// out cleanly. Production paths (proposer, relay client) should always pass
// the prefix from the wallet's BtcAccountConfigured binding.
func ValidateDerivationPath(path, prefix string) error {
	if prefix == "" {
		return nil
	}
	// Tolerate a caller passing a trailing slash on the prefix — the
	// attestation regex already rejects that shape upstream, but a defensive
	// trim makes this helper safe to feed from caller code that string-builds
	// the prefix without normalising it.
	prefix = strings.TrimRight(prefix, "/")
	if !strings.HasPrefix(path, prefix+"/") {
		return fmt.Errorf("derivationPath %q does not start with prefix %q", path, prefix)
	}
	leafPart := path[len(prefix)+1:]
	chain, leafIdx, found := strings.Cut(leafPart, "/")
	if !found {
		return fmt.Errorf("derivationPath %q missing leafIndex (expected %s/<0|1>/<leafIndex>)", path, prefix)
	}
	if strings.Contains(leafIdx, "/") {
		return fmt.Errorf("derivationPath %q has extra components after leafIndex", path)
	}
	if chain != "0" && chain != "1" {
		return fmt.Errorf("derivationPath %q has unexpected chain %q (expected 0 external or 1 internal)", path, chain)
	}
	if leafIdx == "" {
		return fmt.Errorf("derivationPath %q has empty leafIndex", path)
	}
	n, err := strconv.ParseUint(leafIdx, 10, 32)
	if err != nil {
		return fmt.Errorf("derivationPath %q has unparseable leafIndex %q: %v", path, leafIdx, err)
	}
	if leafIdx != strconv.FormatUint(n, 10) {
		return fmt.Errorf("derivationPath %q has non-canonical leafIndex %q (leading zero or other non-canonical decimal form)", path, leafIdx)
	}
	if n >= 1<<31 {
		return fmt.Errorf("derivationPath %q leafIndex %d is in the hardened range (BIP-32 reserves 2^31+ for hardened keys)", path, n)
	}
	return nil
}
