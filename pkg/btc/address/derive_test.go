package address

import (
	"bytes"
	"crypto/sha256"
	"sort"
	"testing"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

// generateTestXpubs produces n deterministic account-level xpubs at the
// non-hardened account m/87'/0'/0 (BIP-87, 2026-06-22 model) from fixed seeds:
// two hardened levels to the wallet-level parent m/87'/0', then the
// NON-hardened account child via DeriveAccountXpubs — the production shape.
// Reproducible across test runs.
func generateTestXpubs(t *testing.T, params *chaincfg.Params, n int) []string {
	t.Helper()
	parents := make([]string, n)
	for i := range n {
		seed := bytes.Repeat([]byte{byte(i + 1)}, 32)
		master, err := hdkeychain.NewMaster(seed, params)
		if err != nil {
			t.Fatalf("master[%d]: %v", i, err)
		}
		parent := master
		for _, hard := range []uint32{87, 0} { // m/87'/0' — wallet-level (parent), hardened
			parent, err = parent.Derive(hdkeychain.HardenedKeyStart + hard)
			if err != nil {
				t.Fatalf("derive %d: %v", hard, err)
			}
		}
		xpub, err := parent.Neuter()
		if err != nil {
			t.Fatalf("neuter[%d]: %v", i, err)
		}
		parents[i] = xpub.String()
	}
	acctXpubs, err := DeriveAccountXpubs(parents, 0, params) // account 0, NON-hardened
	if err != nil {
		t.Fatalf("DeriveAccountXpubs: %v", err)
	}
	return acctXpubs
}

// TestDeriveAccountXpubs_NonHardenedStep proves the account-level xpub is the
// NON-HARDENED CKDpub child of the published wallet-level (parent) xpub at
// accountIndex (2026-06-22 UTXO-model alignment), and that the resulting
// account xpub feeds Derive to a real address. Also pins the hardened-range
// rejection (accountIndex >= 2^31 cannot be derived from a public parent).
func TestDeriveAccountXpubs_NonHardenedStep(t *testing.T) {
	params := &chaincfg.RegressionNetParams
	const accountIndex = 7

	// Parent (wallet-level) xpub at m/87'/1' (regtest coin_type).
	master, err := hdkeychain.NewMaster(bytes.Repeat([]byte{0x07}, 32), params)
	if err != nil {
		t.Fatalf("master: %v", err)
	}
	parent := master
	for _, h := range []uint32{87, 1} {
		if parent, err = parent.Derive(hdkeychain.HardenedKeyStart + h); err != nil {
			t.Fatalf("derive %d': %v", h, err)
		}
	}
	parentXpub, err := parent.Neuter()
	if err != nil {
		t.Fatalf("neuter: %v", err)
	}

	got, err := DeriveAccountXpubs([]string{parentXpub.String()}, accountIndex, params)
	if err != nil {
		t.Fatalf("DeriveAccountXpubs: %v", err)
	}

	// Independently derive the expected account xpub: parent / accountIndex,
	// NON-hardened (no HardenedKeyStart offset).
	wantAcct, err := parentXpub.Derive(accountIndex)
	if err != nil {
		t.Fatalf("manual non-hardened account derive: %v", err)
	}
	if got[0] != wantAcct.String() {
		t.Errorf("account xpub = %q, want non-hardened child %q", got[0], wantAcct.String())
	}

	// The derived account xpub must feed Derive cleanly (real leaf address).
	if _, _, _, err := Derive(got, 1, External, 0, params); err != nil {
		t.Fatalf("Derive on derived account xpub: %v", err)
	}

	// accountIndex in the hardened range is unreachable from a public parent.
	if _, err := DeriveAccountXpubs([]string{parentXpub.String()}, 1<<31, params); err == nil {
		t.Error("expected reject for accountIndex >= 2^31 (hardened range)")
	}
}

func TestDerive_BasicProperties(t *testing.T) {
	params := &chaincfg.RegressionNetParams
	xpubs := generateTestXpubs(t, params, 3)

	// 2-of-3 multisig at external chain, index 0
	addr, witnessScript, pubkeys, err := Derive(xpubs, 2, External, 0, params)
	if err != nil {
		t.Fatalf("Derive: %v", err)
	}

	// Address should be non-nil and bech32
	if addr == nil {
		t.Fatal("expected non-nil address")
	}
	addrStr := addr.EncodeAddress()
	if len(addrStr) == 0 {
		t.Fatal("empty address string")
	}
	// Regtest bech32 addresses start with "bcrt1"
	if addrStr[:5] != "bcrt1" {
		t.Errorf("expected bcrt1 prefix, got %s", addrStr[:5])
	}

	// Witness script should be non-empty
	if len(witnessScript) == 0 {
		t.Fatal("empty witness script")
	}

	// Should have 3 sorted pubkeys, each 33 bytes
	if len(pubkeys) != 3 {
		t.Fatalf("expected 3 pubkeys, got %d", len(pubkeys))
	}
	for i, pk := range pubkeys {
		if len(pk) != 33 {
			t.Errorf("pubkey[%d] length = %d, want 33", i, len(pk))
		}
	}

	// Pubkeys should be BIP-67 sorted (lexicographic)
	for i := 0; i < len(pubkeys)-1; i++ {
		if bytes.Compare(pubkeys[i], pubkeys[i+1]) >= 0 {
			t.Errorf("pubkeys not sorted at index %d", i)
		}
	}

	// Verify the address matches SHA256(witnessScript)
	scriptHash := sha256.Sum256(witnessScript)
	reconstructed, err := btcutil.NewAddressWitnessScriptHash(scriptHash[:], params)
	if err != nil {
		t.Fatalf("reconstruct address: %v", err)
	}
	if reconstructed.EncodeAddress() != addrStr {
		t.Errorf("address mismatch: %s != %s", reconstructed.EncodeAddress(), addrStr)
	}
}

func TestDerive_DeterministicAndDifferentIndices(t *testing.T) {
	params := &chaincfg.RegressionNetParams
	xpubs := generateTestXpubs(t, params, 3)

	// Same inputs should produce same output
	addr1, _, _, err := Derive(xpubs, 2, External, 0, params)
	if err != nil {
		t.Fatalf("Derive 1: %v", err)
	}
	addr2, _, _, err := Derive(xpubs, 2, External, 0, params)
	if err != nil {
		t.Fatalf("Derive 2: %v", err)
	}
	if addr1.EncodeAddress() != addr2.EncodeAddress() {
		t.Error("same inputs produced different addresses")
	}

	// Different leaf index should produce different address
	addr3, _, _, err := Derive(xpubs, 2, External, 1, params)
	if err != nil {
		t.Fatalf("Derive 3: %v", err)
	}
	if addr1.EncodeAddress() == addr3.EncodeAddress() {
		t.Error("different leaf index produced same address")
	}

	// Different chain should produce different address
	addr4, _, _, err := Derive(xpubs, 2, Internal, 0, params)
	if err != nil {
		t.Fatalf("Derive 4: %v", err)
	}
	if addr1.EncodeAddress() == addr4.EncodeAddress() {
		t.Error("different chain produced same address")
	}
}

// generateLegacyBIP44Xpubs is the BIP-44 counterpart of generateTestXpubs:
// it derives at m/44'/0'/0' instead of m/87'/0'/0. Used to pin the back-compat
// claim in docs/shared/wallet-and-address-derivation.md that pre-BIP-87 wallets
// (with derivationPathPrefix = m/44'/0'/accountIndex') still work through the
// derive stack unchanged. The Derive function takes xpubs as bytes — it does
// not see the path — so this test guards against any future change that would
// implicitly bind Derive to a particular purpose level.
func generateLegacyBIP44Xpubs(t *testing.T, params *chaincfg.Params, n int) []string {
	t.Helper()
	out := make([]string, n)
	for i := range n {
		seed := bytes.Repeat([]byte{byte(i + 1)}, 32)
		master, err := hdkeychain.NewMaster(seed, params)
		if err != nil {
			t.Fatalf("master[%d]: %v", i, err)
		}
		acct := master
		for _, hard := range []uint32{44, 0, 0} { // BIP-44 legacy hierarchy
			acct, err = acct.Derive(hdkeychain.HardenedKeyStart + hard)
			if err != nil {
				t.Fatalf("derive %d: %v", hard, err)
			}
		}
		xpub, err := acct.Neuter()
		if err != nil {
			t.Fatalf("neuter[%d]: %v", i, err)
		}
		out[i] = xpub.String()
	}
	return out
}

func TestDerive_LegacyBIP44Compat(t *testing.T) {
	params := &chaincfg.RegressionNetParams

	bip87 := generateTestXpubs(t, params, 3)
	bip44 := generateLegacyBIP44Xpubs(t, params, 3)

	// BIP-44 and BIP-87 must produce different xpubs from the same seeds
	// (otherwise the migration would be a no-op).
	for i := range bip87 {
		if bip87[i] == bip44[i] {
			t.Fatalf("xpub[%d] is identical across BIP-44 and BIP-87 paths — derivation is not path-sensitive", i)
		}
	}

	// Both must derive to a valid bcrt1 P2WSH address through Derive.
	addr87, _, _, err := Derive(bip87, 2, External, 0, params)
	if err != nil {
		t.Fatalf("Derive BIP-87: %v", err)
	}
	addr44, _, _, err := Derive(bip44, 2, External, 0, params)
	if err != nil {
		t.Fatalf("Derive BIP-44: %v", err)
	}

	if addr87.EncodeAddress()[:5] != "bcrt1" {
		t.Errorf("BIP-87 address not bcrt1: %s", addr87.EncodeAddress())
	}
	if addr44.EncodeAddress()[:5] != "bcrt1" {
		t.Errorf("BIP-44 address not bcrt1: %s", addr44.EncodeAddress())
	}

	// The two address sets must be disjoint — different xpubs ⇒ different scripts.
	if addr87.EncodeAddress() == addr44.EncodeAddress() {
		t.Error("BIP-44 and BIP-87 derivations collapsed to the same address")
	}
}

func TestDerive_ValidationErrors(t *testing.T) {
	params := &chaincfg.RegressionNetParams
	xpubs := generateTestXpubs(t, params, 3)

	// threshold too low
	_, _, _, err := Derive(xpubs, 0, External, 0, params)
	if err == nil {
		t.Error("expected error for threshold 0")
	}

	// threshold too high
	_, _, _, err = Derive(xpubs, 4, External, 0, params)
	if err == nil {
		t.Error("expected error for threshold > n")
	}

	// invalid xpub string
	_, _, _, err = Derive([]string{"not-a-valid-xpub"}, 1, External, 0, params)
	if err == nil {
		t.Error("expected error for invalid xpub")
	}

	// n exceeds the OP_CHECKMULTISIG limit of 20.
	tooMany := generateTestXpubs(t, params, 21)
	if _, _, _, err = Derive(tooMany, 1, External, 0, params); err == nil {
		t.Error("expected error for n > 20")
	}

	// well-formed xpub for the wrong network fails the IsForNet check.
	mainnetXpubs := generateTestXpubs(t, &chaincfg.MainNetParams, 1)
	if _, _, _, err = Derive(mainnetXpubs, 1, External, 0, params); err == nil {
		t.Error("expected error for xpub whose network version does not match params")
	}
}

// TestDeriveAccountXpubs_Errors covers the reachable failure branches of
// DeriveAccountXpubs: an unparseable parent xpub and a well-formed parent xpub
// for the wrong network. (The hardened-range guard is covered in
// TestDeriveAccountXpubs_NonHardenedStep.)
func TestDeriveAccountXpubs_Errors(t *testing.T) {
	params := &chaincfg.RegressionNetParams

	if _, err := DeriveAccountXpubs([]string{"not-a-valid-xpub"}, 0, params); err == nil {
		t.Error("expected error for unparseable parent xpub")
	}

	mainnetXpubs := generateTestXpubs(t, &chaincfg.MainNetParams, 1)
	if _, err := DeriveAccountXpubs(mainnetXpubs, 0, params); err == nil {
		t.Error("expected error for parent xpub whose network version does not match params")
	}
}

// TestDerive_RejectsPrivateKey ensures a serialized private extended key
// (xprv/tprv) is rejected in the public-key inputs: IsForNet alone is true for
// private keys, so an explicit IsPrivate guard is required. Secret material must
// never be accepted where a public key is expected.
func TestDerive_RejectsPrivateKey(t *testing.T) {
	params := &chaincfg.RegressionNetParams
	master, err := hdkeychain.NewMaster(bytes.Repeat([]byte{0x09}, 32), params)
	if err != nil {
		t.Fatalf("master: %v", err)
	}
	// A depth-2 PRIVATE extended key (not neutered): m/87'/1'.
	priv := master
	for _, h := range []uint32{87, 1} {
		if priv, err = priv.Derive(hdkeychain.HardenedKeyStart + h); err != nil {
			t.Fatalf("derive %d': %v", h, err)
		}
	}
	privStr := priv.String() // xprv/tprv form

	if _, _, _, err := Derive([]string{privStr}, 1, External, 0, params); err == nil {
		t.Error("Derive: expected rejection of a private extended key")
	}
	if _, err := DeriveAccountXpubs([]string{privStr}, 0, params); err == nil {
		t.Error("DeriveAccountXpubs: expected rejection of a private extended key")
	}
}

// TestDerive_RejectsDuplicateKeys ensures a repeated signer key is rejected: a
// k-of-n with duplicates collapses signer independence (one holder can satisfy
// multiple CHECKMULTISIG slots), so it is not a genuine k-of-n multisig.
func TestDerive_RejectsDuplicateKeys(t *testing.T) {
	params := &chaincfg.RegressionNetParams
	xpubs := generateTestXpubs(t, params, 3)
	dup := []string{xpubs[0], xpubs[0], xpubs[1]} // repeat signer 0

	if _, _, _, err := Derive(dup, 2, External, 0, params); err == nil {
		t.Error("Derive: expected rejection of duplicate signer keys")
	}
}

// TestDerive_KnownAnswerAddress pins the P2WSH multisig address for a fixed
// 2-of-3 wallet (deterministic seeds) against an INDEPENDENT reconstruction and
// a hard-coded literal. The reconstruction assembles the witnessScript from raw
// bytes (no ScriptBuilder), so a logic bug in Derive — dropped BIP-67 sort,
// wrong chain/leaf index, HASH160 instead of SHA256, wrong opcode — diverges and
// is caught. The literal freezes behaviour against a shared-primitive change.
func TestDerive_KnownAnswerAddress(t *testing.T) {
	params := &chaincfg.MainNetParams
	parents := generateTestXpubs(t, params, 3)
	acct, err := DeriveAccountXpubs(parents, 0, params)
	if err != nil {
		t.Fatalf("DeriveAccountXpubs: %v", err)
	}

	addr, _, _, err := Derive(acct, 2, External, 0, params)
	if err != nil {
		t.Fatalf("Derive: %v", err)
	}
	got := addr.EncodeAddress()

	want := independentP2WSHAddress(t, acct, 2, External, 0, params)
	if got != want {
		t.Fatalf("Derive address %q != independent reconstruction %q", got, want)
	}

	const knownAnswer = "bc1qk2mze0hf2grhe2axtqv0d6pn4rxlrzlr3xtxm4vxy09h4uxqrpksswq0x7"
	if got != knownAnswer {
		t.Fatalf("address %q != pinned known answer %q (update the literal only after verifying the change is intended)", got, knownAnswer)
	}
}

// independentP2WSHAddress derives the sorted k-of-n P2WSH address the same way
// the spec defines it but WITHOUT reusing Derive's script-building path: it
// CKDpubs each account xpub to (chain, leaf), BIP-67 sorts the compressed
// pubkeys, hand-assembles `OP_k <pk...> OP_n OP_CHECKMULTISIG` as raw bytes,
// SHA256s it, and bech32-encodes the witness-v0 program.
func independentP2WSHAddress(t *testing.T, acctXpubs []string, k int, chain Chain, leaf uint32, params *chaincfg.Params) string {
	t.Helper()
	pubs := make([][]byte, len(acctXpubs))
	for i, xs := range acctXpubs {
		key, err := hdkeychain.NewKeyFromString(xs)
		if err != nil {
			t.Fatalf("parse[%d]: %v", i, err)
		}
		ck, err := key.Derive(uint32(chain))
		if err != nil {
			t.Fatalf("chain[%d]: %v", i, err)
		}
		lk, err := ck.Derive(leaf)
		if err != nil {
			t.Fatalf("leaf[%d]: %v", i, err)
		}
		pk, err := lk.ECPubKey()
		if err != nil {
			t.Fatalf("pubkey[%d]: %v", i, err)
		}
		pubs[i] = pk.SerializeCompressed()
	}
	sort.Slice(pubs, func(i, j int) bool { return bytes.Compare(pubs[i], pubs[j]) < 0 })

	n := len(pubs)
	if k < 1 || k > 16 || n < 1 || n > 16 {
		t.Fatalf("independent helper only supports 1..16 for k and n; got k=%d n=%d", k, n)
	}
	var script []byte
	script = append(script, byte(txscript.OP_1-1+k)) // OP_k (0x50+k for 1..16)
	for _, pk := range pubs {
		script = append(script, byte(len(pk))) // 0x21 push of 33 bytes
		script = append(script, pk...)
	}
	script = append(script, byte(txscript.OP_1-1+n)) // OP_n
	script = append(script, txscript.OP_CHECKMULTISIG)

	h := sha256.Sum256(script)
	addr, err := btcutil.NewAddressWitnessScriptHash(h[:], params)
	if err != nil {
		t.Fatalf("encode: %v", err)
	}
	return addr.EncodeAddress()
}

// TestDerive_LargeN exercises the n=17..20 range that triggers the
// AddInt64 minimal scriptNum push path (vs. OP_1..OP_16). A regression
// to byte(OP_1-1+x) would emit OP_NOP (0x61) for n=17, OP_VER (0x62)
// for n=18, etc., yielding an unspendable script. See open-question #14
// (n up to 20 supported) and the matching fix in htlc.go.
func TestDerive_LargeN(t *testing.T) {
	params := &chaincfg.RegressionNetParams

	// AddInt64 emits OP_k for k in 1..16 (single byte 0x50+k) and a minimal
	// scriptNum push for k in 17..20 (0x01 followed by the value byte).
	expectedHead := func(x int) []byte {
		if x <= 16 {
			return []byte{byte(txscript.OP_1 - 1 + x)}
		}
		return []byte{0x01, byte(x)}
	}

	for n := 17; n <= 20; n++ {
		xpubs := generateTestXpubs(t, params, n)
		for _, k := range []int{1, n / 2, n} {
			addr, witnessScript, _, err := Derive(xpubs, k, External, 0, params)
			if err != nil {
				t.Fatalf("Derive(n=%d, k=%d): %v", n, k, err)
			}
			if addr == nil || witnessScript == nil {
				t.Fatalf("Derive(n=%d, k=%d): nil address or script", n, k)
			}

			// Leading bytes: minimal push of threshold k.
			want := expectedHead(k)
			if !bytes.Equal(witnessScript[:len(want)], want) {
				t.Errorf("n=%d, k=%d: head bytes = % x, want % x", n, k, witnessScript[:len(want)], want)
			}

			// Trailing bytes: <push n> OP_CHECKMULTISIG.
			wantTail := append(expectedHead(n), txscript.OP_CHECKMULTISIG)
			tail := witnessScript[len(witnessScript)-len(wantTail):]
			if !bytes.Equal(tail, wantTail) {
				t.Errorf("n=%d, k=%d: tail bytes = % x, want % x", n, k, tail, wantTail)
			}

			// Sanity: tokenise the script and check that we see no OP_NOP /
			// OP_VER / OP_IF / OP_NOTIF (which the buggy byte(OP_1-1+x)
			// encoding would have injected for x in 17..20) and the
			// non-pubkey ops are exactly [<push k>, ..., <push n>,
			// OP_CHECKMULTISIG].
			tok := txscript.MakeScriptTokenizer(0, witnessScript)
			var ops []byte
			for tok.Next() {
				op := tok.Opcode()
				// Skip the pubkey pushes (data-driven, 33-byte payloads).
				if len(tok.Data()) == 33 {
					continue
				}
				ops = append(ops, op)
				// A bug-injected OP_NOP (0x61), OP_VER (0x62), OP_IF (0x63)
				// or OP_NOTIF (0x64) would fire here.
				if op == txscript.OP_NOP || op == txscript.OP_VER || op == txscript.OP_IF || op == txscript.OP_NOTIF {
					t.Errorf("n=%d, k=%d: unexpected control opcode 0x%02x in script — n>16 regression?", n, k, op)
				}
			}
			if err := tok.Err(); err != nil {
				t.Errorf("n=%d, k=%d: tokenizer: %v", n, k, err)
			}
			if len(ops) != 3 || ops[len(ops)-1] != txscript.OP_CHECKMULTISIG {
				t.Errorf("n=%d, k=%d: non-pubkey ops = % x, want [<k>, <n>, OP_CHECKMULTISIG]", n, k, ops)
			}
		}
	}
}

func TestDerive_RejectsNTooLarge(t *testing.T) {
	params := &chaincfg.RegressionNetParams
	xpubs := generateTestXpubs(t, params, 21)
	_, _, _, err := Derive(xpubs, 11, External, 0, params)
	if err == nil {
		t.Fatal("expected error for n=21 (exceeds OP_CHECKMULTISIG limit of 20)")
	}
}

func TestValidateDerivationPath_Accepts(t *testing.T) {
	cases := []struct {
		name   string
		prefix string
		path   string
	}{
		{"bip87-external-zero", "m/87'/0'/0", "m/87'/0'/0/0/0"},
		{"bip87-external-small", "m/87'/0'/0", "m/87'/0'/0/0/19"},
		{"bip87-internal-change", "m/87'/0'/0", "m/87'/0'/0/1/42"},
		{"bip44-legacy-external", "m/44'/0'/0'", "m/44'/0'/0'/0/7"},
		{"bip44-legacy-internal", "m/44'/0'/0'", "m/44'/0'/0'/1/0"},
		{"bip87-account-7", "m/87'/0'/7", "m/87'/0'/7/0/1234"},
		{"bip87-testnet-coin", "m/87'/1'/0", "m/87'/1'/0/0/0"},
		{"max-leaf-non-hardened", "m/87'/0'/0", "m/87'/0'/0/0/2147483647"},
		{"empty-prefix-skips", "", "anything goes when prefix is empty"},
		{"trailing-slash-prefix-trimmed", "m/87'/0'/0/", "m/87'/0'/0/0/0"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := ValidateDerivationPath(c.path, c.prefix); err != nil {
				t.Errorf("ValidateDerivationPath(%q, %q) = %v, want nil", c.path, c.prefix, err)
			}
		})
	}
}

func TestValidateDerivationPath_Rejects(t *testing.T) {
	const prefix = "m/87'/0'/0"
	cases := []struct {
		name string
		path string
		want string // substring expected in the error
	}{
		{"wrong-prefix", "m/44'/0'/0'/0/0", "does not start with prefix"},
		{"missing-leaf", "m/87'/0'/0/0", "missing leafIndex"},
		{"empty-leaf", "m/87'/0'/0/0/", "empty leafIndex"},
		{"chain-2", "m/87'/0'/0/2/0", "unexpected chain"},
		{"chain-hardened", "m/87'/0'/0/0'/0", "unexpected chain"},
		{"non-numeric-leaf", "m/87'/0'/0/0/abc", "unparseable leafIndex"},
		{"hardened-leaf", "m/87'/0'/0/0/0'", "unparseable leafIndex"},
		{"leading-zero-leaf", "m/87'/0'/0/0/01", "leading zero"},
		{"extra-components", "m/87'/0'/0/0/0/extra", "extra components"},
		{"leaf-in-hardened-range", "m/87'/0'/0/0/2147483648", "hardened range"},
		{"negative-leaf", "m/87'/0'/0/0/-1", "unparseable leafIndex"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := ValidateDerivationPath(c.path, prefix)
			if err == nil {
				t.Fatalf("ValidateDerivationPath(%q, %q) = nil, want error mentioning %q", c.path, prefix, c.want)
			}
			if c.want != "" && !bytes.Contains([]byte(err.Error()), []byte(c.want)) {
				t.Errorf("error %q does not contain %q", err.Error(), c.want)
			}
		})
	}
}
