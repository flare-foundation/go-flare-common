package address

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
)

// testMainnetXpub is the fixed BIP-32 mainnet WALLET-LEVEL (parent) xpub used
// across all attestation_test fixtures — depth 2 (m/87'/0'), since the account
// level is a NON-HARDENED child at accountIndex.
// Mainnet version byte ⇒ MainNetParams. Deterministically generated from seed
// 0x2a×32 → m/87'/0' (the same shape xpubAtDepth(t, params, 87, 0) produces).
const testMainnetXpub = "xpub6AA1xY86BDWPPrATRwypWcB7Z5Kxu2fhAdTLDTUXKbL7mMQ9NJXwnvsitFKg3bCMBComzwbdo3Je1zwAY1GMgfrXbtC2gPknszPQETGv2d1"

// validV1Xpubs returns a fixed 2-of-3 xpub set that ValidateV1 accepts on
// mainnet. The three keys are DISTINCT (deterministic seeds 1..3, m/87'/0') —
// ValidateV1 rejects duplicate signers, so a genuine k-of-n fixture must use
// independent keys. Tests that don't care about xpub identity reuse it; tests
// that probe network-identity or n-bound assertions construct their own.
func validV1Xpubs() []string {
	return []string{
		"xpub6BT5TaYRGEymDz7PN6eEwJHaZuR3cDbLaLZGAaGn42YLo5rhz76cUcTnJ9zexAewNRCgCG8bV8PmCzveQCnQGfn7sFyWt2P5WnzSo4S7RRa",
		"xpub6AizowepZHn9HVBcyFtePudx2ufnhKNNjQCZLnwuzjBrhXnmwKk7b2SAhcXw1ddfgrme9FyBZUspiMRJmX5bGxQLgoGpLtSTUoeyKY6xTCP",
		"xpub6BbwkRCpSqWb354VTWBzLUEdErvYucHhdDMJ7irvkiXCzR3EJTUH81aKn3JvU9f9ryFCS8JWQr9NNfSF8oB9JyYPB3ERx6DXjfujuUeWLib",
	}
}

// solidityReservedTagsCount mirrors MintingTagManager._reservedTagsCount in
// flare-smart-contracts-v2, and verifierAnchorCountCap mirrors the
// `1 <= anchors.length <= 32` bound the FDC verifier enforces (and that
// the WalletPayments diamond relies on). They are
// duplicated here ON PURPOSE: the reserved derivation-index layout is a
// protocol-wide constant, NOT a per-wallet attestation field, so its
// cross-repo agreement is not enforced on the wire. This test is the Go half
// of that guard — if someone edits the Go layout constants, this fails and
// forces the Solidity side to be revisited rather than letting user-deposit
// indices silently collide with the anchor / protocol-reserved bands.
const (
	solidityReservedTagsCount = 256
	verifierAnchorCountCap    = 32
)

// TestReservedLayoutInvariants pins the derivation-index reserved layout and
// its cross-repo mirror values.
func TestReservedLayoutInvariants(t *testing.T) {
	// The two externally-mirrored constants.
	if MaxReservedAddresses != solidityReservedTagsCount {
		t.Errorf("MaxReservedAddresses = %d, but MintingTagManager._reservedTagsCount is %d; "+
			"user deposits start at index MaxReservedAddresses, so these MUST agree",
			MaxReservedAddresses, solidityReservedTagsCount)
	}
	if MaxAnchors != verifierAnchorCountCap {
		t.Errorf("MaxAnchors = %d, but the FDC verifier / WalletPayments anchors.length cap is %d; "+
			"these MUST agree", MaxAnchors, verifierAnchorCountCap)
	}

	// Band ordering: anchors [0, MaxAnchors) | protocol-reserved
	// [MaxAnchors, MaxReservedAddresses) | user deposits [MaxReservedAddresses, ...).
	if 0 >= MaxAnchors || MaxAnchors >= MaxReservedAddresses {
		t.Fatalf("reserved bands out of order: MaxAnchors=%d MaxReservedAddresses=%d "+
			"(want 0 < MaxAnchors < MaxReservedAddresses)", MaxAnchors, MaxReservedAddresses)
	}

	// FirstProtocolReservedIndex is the first slot past the anchor band, and
	// must leave a non-empty protocol-reserved band before user deposits.
	if FirstProtocolReservedIndex != MaxAnchors {
		t.Errorf("FirstProtocolReservedIndex = %d, want MaxAnchors = %d",
			FirstProtocolReservedIndex, MaxAnchors)
	}
	if FirstProtocolReservedIndex >= MaxReservedAddresses {
		t.Errorf("FirstProtocolReservedIndex = %d must be < MaxReservedAddresses = %d "+
			"(protocol-reserved band would be empty)", FirstProtocolReservedIndex, MaxReservedAddresses)
	}

	// FAssets' default N must fit inside the anchor band.
	if 1 > ActiveAnchors || ActiveAnchors > MaxAnchors {
		t.Errorf("ActiveAnchors = %d out of [1, MaxAnchors=%d]", ActiveAnchors, MaxAnchors)
	}
}

func TestValidateV1_HappyPath(t *testing.T) {
	b := &BtcAccountConfigured{
		AccountIndex: 0,
		Xpubs:        validV1Xpubs(),
		Threshold:    2,
		Anchors:      defaultAnchorSet(),
	}
	if err := b.ValidateV1(&chaincfg.MainNetParams); err != nil {
		t.Fatalf("ValidateV1: %v", err)
	}
}

func TestValidateV1_HappyPath_HigherAccountIndex(t *testing.T) {
	// AccountIndex is variable (one published parent xpub serves many accounts,
	// each a non-hardened child); pin a non-zero value to confirm the verifier
	// handles it. The path the verifier derives — m/87'/0'/7 (account NON-
	// hardened) — is implicit; the test asserts ValidateV1 accepts and
	// DeriveV1PathPrefix threads it.
	b := &BtcAccountConfigured{
		AccountIndex: 7,
		Xpubs:        validV1Xpubs(),
		Threshold:    2,
		Anchors:      defaultAnchorSet(),
	}
	if err := b.ValidateV1(&chaincfg.MainNetParams); err != nil {
		t.Fatalf("ValidateV1 with non-zero accountIndex: %v", err)
	}
	got, err := DeriveV1PathPrefix(&chaincfg.MainNetParams, b.AccountIndex)
	if err != nil {
		t.Fatalf("DeriveV1PathPrefix: %v", err)
	}
	if got != "m/87'/0'/7" {
		t.Errorf("derived prefix = %q, want m/87'/0'/7", got)
	}
}

// xpubAtDepth derives an xpub stopping after the given hardened levels, so
// callers can produce a key at an arbitrary BIP-32 depth (e.g. depth 2 for
// m/87'/0', one level short of the account level).
func xpubAtDepth(t *testing.T, params *chaincfg.Params, hardened ...uint32) string {
	t.Helper()
	k, err := hdkeychain.NewMaster(bytes.Repeat([]byte{0x2a}, 32), params)
	if err != nil {
		t.Fatalf("master: %v", err)
	}
	for _, h := range hardened {
		if k, err = k.Derive(hdkeychain.HardenedKeyStart + h); err != nil {
			t.Fatalf("derive %d': %v", h, err)
		}
	}
	xpub, err := k.Neuter()
	if err != nil {
		t.Fatalf("neuter: %v", err)
	}
	return xpub.String()
}

// generateTestParentXpubs produces n deterministic WALLET-LEVEL (parent) xpubs
// at m/87'/0' (BIP-32 depth 2) from per-key seeds, network-appropriate version
// byte. These are the shape the attestation now expects (the account level is a
// non-hardened child at accountIndex) — the
// depth-2 counterpart of derive_test.go's account-level generateTestXpubs.
func generateTestParentXpubs(t *testing.T, params *chaincfg.Params, n int) []string {
	t.Helper()
	coin, err := bip87CoinType(params)
	if err != nil {
		t.Fatalf("coin for %s: %v", params.Name, err)
	}
	out := make([]string, n)
	for i := range n {
		k, err := hdkeychain.NewMaster(bytes.Repeat([]byte{byte(i + 1)}, 32), params)
		if err != nil {
			t.Fatalf("master[%d]: %v", i, err)
		}
		for _, h := range []uint32{87, coin} { // m/87'/coin' — wallet (parent) level, the network's coin
			if k, err = k.Derive(hdkeychain.HardenedKeyStart + h); err != nil {
				t.Fatalf("derive %d': %v", h, err)
			}
		}
		xpub, err := k.Neuter()
		if err != nil {
			t.Fatalf("neuter[%d]: %v", i, err)
		}
		out[i] = xpub.String()
	}
	return out
}

// TestValidateV1_RejectsAnotherCoin: a key at the right depth and under the
// network's version bytes is still refused when its child number is another
// coin. m/87'/0' re-serialized as a tpub is a MAINNET key, and on a test
// network it would put mainnet keys behind a test wallet.
func TestValidateV1_RejectsAnotherCoin(t *testing.T) {
	params := &chaincfg.RegressionNetParams
	keys := generateTestParentXpubs(t, params, 3)
	k, err := hdkeychain.NewMaster(bytes.Repeat([]byte{9}, 32), params)
	if err != nil {
		t.Fatal(err)
	}
	for _, h := range []uint32{87, 0} { // coin 0' under a tpub
		if k, err = k.Derive(hdkeychain.HardenedKeyStart + h); err != nil {
			t.Fatal(err)
		}
	}
	pub, err := k.Neuter()
	if err != nil {
		t.Fatal(err)
	}
	keys[1] = pub.String()
	b := &BtcAccountConfigured{AccountIndex: 0, Xpubs: keys, Threshold: 2, Anchors: defaultAnchorSet()}
	err = b.ValidateV1(params)
	if err == nil || !strings.Contains(err.Error(), "not coin 1'") {
		t.Fatalf("a coin-0' key passed on regtest: %v", err)
	}
}

// TestValidateXpubs_RejectsNonParentLevelDepth pins the per-key depth check
// on BtcAccountConfigured: every published xpub must be at the WALLET
// (parent) level (BIP-32 depth 2, m/87'/coin'). The account level is a
// non-hardened child at accountIndex, so a depth-3 (account-level) xpub is
// one level too deep and is rejected.
// Called directly (rather than via ValidateV1) so the depth assertion is
// isolated.
func TestValidateXpubs_RejectsNonParentLevelDepth(t *testing.T) {
	params := &chaincfg.MainNetParams

	// Depth-3 xpub: m/87'/0'/0' — the (old) account level, now one level too deep.
	// Single key with threshold 1 so the depth check (not a count/duplicate check)
	// is what rejects it.
	deep := xpubAtDepth(t, params, 87, 0, 0)
	err := validateXpubsAndThreshold([]string{deep}, 1, params)
	if err == nil {
		t.Fatal("expected rejection of a depth-3 xpub, got nil")
	}
	if !strings.Contains(err.Error(), "depth 3") || !strings.Contains(err.Error(), "depth 2") {
		t.Errorf("error %q should name the actual depth (3) and the expected wallet-level depth (2)", err)
	}

	// Distinct depth-2 xpubs: m/87'/0' — the wallet (parent) level — must pass.
	good := validV1Xpubs()
	if err := validateXpubsAndThreshold(good[:2], 2, params); err != nil {
		t.Fatalf("wallet-level (depth-2) parent xpubs should validate: %v", err)
	}
}

// TestValidateV1_RejectsXpubsThreshold pins the BIP-32 / multisig invariants
// that the attestation layer must enforce *before* any tx-construction code
// reaches the wallet. Tests the four reject paths in validateXpubsAndThreshold:
// empty Xpubs, n>20 (the OP_CHECKMULTISIG cap), threshold out of [1, n], malformed
// xpub, and a cross-network xpub (mainnet xpub bound to a testnet wallet).
func TestValidateV1_RejectsXpubsThreshold(t *testing.T) {
	good := validV1Xpubs() // 3 mainnet xpubs
	manyXpubs := make([]string, 21)
	for i := range manyXpubs {
		manyXpubs[i] = testMainnetXpub
	}

	cases := []struct {
		name      string
		xpubs     []string
		threshold int
		params    *chaincfg.Params
		want      string
	}{
		{"empty xpubs", nil, 1, &chaincfg.MainNetParams, "at least one key"},
		{"n=21 exceeds CHECKMULTISIG cap", manyXpubs, 11, &chaincfg.MainNetParams, "exceeds OP_CHECKMULTISIG limit"},
		{"threshold 0", good, 0, &chaincfg.MainNetParams, "threshold 0 out of range"},
		{"threshold > n", good, 4, &chaincfg.MainNetParams, "threshold 4 out of range"},
		{"malformed xpub", []string{"not-an-xpub", testMainnetXpub, testMainnetXpub}, 2, &chaincfg.MainNetParams, "xpubs[0]"},
		{"duplicate signer", []string{good[0], good[0], good[1]}, 2, &chaincfg.MainNetParams, "duplicates xpubs"},
		// Mainnet xpub bound to a testnet wallet — the version-byte check
		// must catch this even though every other field is well-formed.
		{"network mismatch", good, 2, &chaincfg.TestNet3Params, "does not match network"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b := &BtcAccountConfigured{
				AccountIndex: 0,
				Xpubs:        c.xpubs,
				Threshold:    c.threshold,
				Anchors:      defaultAnchorSet(),
			}
			err := b.ValidateV1(c.params)
			if err == nil {
				t.Fatalf("expected reject for %s, got nil", c.name)
			}
			if !strings.Contains(err.Error(), c.want) {
				t.Errorf("error %q does not contain %q", err, c.want)
			}
		})
	}
}

// respell serializes an xpub carrying keyOf's public key and every other field
// of rest, with edit (if any) applied to the compressed key, chain code, parent
// fingerprint and child number.
func respell(t *testing.T, keyOf, rest string, edit func(pub, chainCode, parentFP []byte, child *uint32)) string {
	t.Helper()
	k, err := hdkeychain.NewKeyFromString(keyOf)
	if err != nil {
		t.Fatalf("parse %s: %v", keyOf, err)
	}
	pk, err := k.ECPubKey()
	if err != nil {
		t.Fatalf("pubkey of %s: %v", keyOf, err)
	}
	r, err := hdkeychain.NewKeyFromString(rest)
	if err != nil {
		t.Fatalf("parse %s: %v", rest, err)
	}
	pub := pk.SerializeCompressed()
	chainCode := r.ChainCode()
	parentFP := binary.BigEndian.AppendUint32(nil, r.ParentFingerprint())
	child := r.ChildIndex()
	if edit != nil {
		edit(pub, chainCode, parentFP, &child)
	}
	return hdkeychain.NewExtendedKey(r.Version(), pub, chainCode, parentFP, r.Depth(), child, false).String()
}

// TestValidateV1_SignerIsThePublicKey pins what the duplicate-signer rule
// compares: the public key up to sign, not the serialization. Each variant
// carries good[0]'s X coordinate under a chain code, parent fingerprint or
// parity byte that ValidateV1 does not constrain, and is accepted in good[0]'s
// place; beside good[0] it is the same key holder in a second slot, and is
// refused. The negated variants flip the parity byte: -P, whose private key is
// the negation of P's. The child number is not among them: ValidateV1 requires
// it to be the network's coin, so a respelled one is refused alone
// (TestValidateV1_RejectsAnotherCoin).
func TestValidateV1_SignerIsThePublicKey(t *testing.T) {
	good := validV1Xpubs()
	variants := []struct {
		name string
		edit func(pub, chainCode, parentFP []byte, child *uint32)
	}{
		{"different chain code", func(_, cc, _ []byte, _ *uint32) { cc[0] ^= 0xff }},
		{"different parent fingerprint", func(_, _, fp []byte, _ *uint32) { fp[0] ^= 0xff }},
		{"negated key", func(pub, _, _ []byte, _ *uint32) { pub[0] ^= 0x01 }},
		{"negated key, different chain code", func(pub, cc, _ []byte, _ *uint32) { pub[0] ^= 0x01; cc[0] ^= 0xff }},
	}
	for _, v := range variants {
		t.Run(v.name, func(t *testing.T) {
			same := respell(t, good[0], good[0], v.edit)
			if same == good[0] {
				t.Fatal("the variant serializes identically to good[0]; it tests nothing")
			}
			alone := &BtcAccountConfigured{
				Xpubs:     []string{same, good[1], good[2]},
				Threshold: 2,
				Anchors:   defaultAnchorSet(),
			}
			if err := alone.ValidateV1(&chaincfg.MainNetParams); err != nil {
				t.Fatalf("the variant is not a valid key by itself: %v", err)
			}
			twice := &BtcAccountConfigured{
				Xpubs:     []string{good[0], good[1], same},
				Threshold: 2,
				Anchors:   defaultAnchorSet(),
			}
			err := twice.ValidateV1(&chaincfg.MainNetParams)
			if err == nil {
				t.Fatal("expected reject: xpubs[2] carries xpubs[0]'s X coordinate")
			}
			if !strings.Contains(err.Error(), "xpubs[2] duplicates xpubs[0]") {
				t.Errorf("error %q does not name the repeated signer", err)
			}
		})
	}

	// The converse: every field but the X coordinate equal to good[0]'s, parity
	// byte included, is a different signer.
	k0, err := hdkeychain.NewKeyFromString(good[0])
	if err != nil {
		t.Fatalf("parse good[0]: %v", err)
	}
	pub0, err := k0.ECPubKey()
	if err != nil {
		t.Fatalf("pubkey of good[0]: %v", err)
	}
	parity := func(pub, _, _ []byte, _ *uint32) { pub[0] = pub0.SerializeCompressed()[0] }
	b := &BtcAccountConfigured{
		Xpubs:     []string{good[0], good[1], respell(t, good[2], good[0], parity)},
		Threshold: 2,
		Anchors:   defaultAnchorSet(),
	}
	if err := b.ValidateV1(&chaincfg.MainNetParams); err != nil {
		t.Fatalf("two xpubs differing only in the X coordinate were refused: %v", err)
	}
}

// TestValidateV1_RoundTripsToDerive proves that the attestation-layer
// xpub/threshold validation accepts exactly the inputs Derive() also
// accepts. A mismatch between the two would let the attestation pass an
// (Xpubs, k) pair that subsequently fails inside Derive — surfacing as a
// runtime error in tx construction rather than at attestation time.
func TestValidateV1_RoundTripsToDerive(t *testing.T) {
	params := &chaincfg.RegressionNetParams
	const accountIndex = 0
	parentXpubs := generateTestParentXpubs(t, params, 3)

	b := &BtcAccountConfigured{
		AccountIndex: accountIndex,
		Xpubs:        parentXpubs,
		Threshold:    2,
		Anchors:      defaultAnchorSet(),
	}
	if err := b.ValidateV1(params); err != nil {
		t.Fatalf("ValidateV1 with parent xpubs: %v", err)
	}
	// The same parent xpubs must round-trip through the non-hardened account
	// step (DeriveAccountXpubs) and then Derive() cleanly — proving the
	// attestation accepts exactly what the address-derivation path consumes.
	acctXpubs, err := DeriveAccountXpubs(parentXpubs, accountIndex, params)
	if err != nil {
		t.Fatalf("DeriveAccountXpubs: %v", err)
	}
	if _, _, _, err := Derive(acctXpubs, 2, External, 0, params); err != nil {
		t.Fatalf("Derive with derived account xpubs: %v", err)
	}
}

// TestValidateV1_RejectsAccountIndexInHardenedRange pins the BIP-32 bound:
// the bare integer in a hardened path segment must be < 2^31, otherwise the
// hardened encoding (n + 2^31) overflows uint32. Bound is the same one
// derive.go ValidateDerivationPath applies to leaf indices.
func TestValidateV1_RejectsAccountIndexInHardenedRange(t *testing.T) {
	// 2^31 is the smallest value already in the hardened range.
	const reservedHardened = uint32(1) << 31
	b := &BtcAccountConfigured{
		AccountIndex: reservedHardened,
		Xpubs:        validV1Xpubs(),
		Threshold:    2,
		Anchors:      defaultAnchorSet(),
	}
	err := b.ValidateV1(&chaincfg.MainNetParams)
	if err == nil {
		t.Fatal("expected reject for accountIndex >= 2^31")
	}
	if !strings.Contains(err.Error(), "hardened range") {
		t.Errorf("error %q does not flag hardened-range overflow", err)
	}
}

func TestDeriveV1PathPrefix(t *testing.T) {
	// The shared v1 path-prefix deriver — single source of truth for the
	// verifier and off-chain callers (e.g. a proposer validating its config).
	cases := []struct {
		name    string
		params  *chaincfg.Params
		account uint32
		want    string
	}{
		{"mainnet", &chaincfg.MainNetParams, 0, "m/87'/0'/0"},
		{"mainnet account 7", &chaincfg.MainNetParams, 7, "m/87'/0'/7"},
		{"signet", &chaincfg.SigNetParams, 0, "m/87'/1'/0"},
		{"testnet3", &chaincfg.TestNet3Params, 3, "m/87'/1'/3"},
		{"regtest", &chaincfg.RegressionNetParams, 0, "m/87'/1'/0"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := DeriveV1PathPrefix(c.params, c.account)
			if err != nil {
				t.Fatalf("DeriveV1PathPrefix(%s, %d): %v", c.name, c.account, err)
			}
			if got != c.want {
				t.Errorf("got %q, want %q", got, c.want)
			}
		})
	}
}

func TestDeriveV1PathPrefix_Rejects(t *testing.T) {
	// Hardened-range accountIndex and unsupported network are rejected.
	if _, err := DeriveV1PathPrefix(&chaincfg.MainNetParams, 1<<31); err == nil {
		t.Error("expected reject for accountIndex in the hardened range (>= 2^31)")
	}
	if _, err := DeriveV1PathPrefix(nil, 0); err == nil {
		t.Error("expected reject for nil params")
	}
}

// --- Anchor-set tests ---

// defaultAnchorSet returns a representative valid anchor set: N=ActiveAnchors
// (8) chains. The array is order-significant — anchors[i] is chain i at
// derivation index i; there are no anchorIndex, derivation-index, or value
// fields. Used as the happy-path Anchors value across the tests.
func defaultAnchorSet() []AnchorBinding {
	anchors := make([]AnchorBinding, ActiveAnchors)
	for i := range uint32(ActiveAnchors) {
		anchors[i] = AnchorBinding{
			Txid: [32]byte{0xaa, byte(i)},
			Vout: i,
		}
	}
	return anchors
}

func TestValidateV1_HappyPath_WithDefaultAnchorSet(t *testing.T) {
	b := &BtcAccountConfigured{
		AccountIndex: 0,
		Xpubs:        validV1Xpubs(),
		Threshold:    2,
		Anchors:      defaultAnchorSet(),
	}
	if err := b.ValidateV1(&chaincfg.MainNetParams); err != nil {
		t.Fatalf("ValidateV1 with default anchor set: %v", err)
	}
}

func TestValidateAnchorSet_RejectsEmpty(t *testing.T) {
	// Empty must be rejected — a wallet declares at least N=1 chain.
	err := ValidateAnchorSet(nil)
	if err == nil {
		t.Fatal("expected reject for empty anchor set")
	}
	if !strings.Contains(err.Error(), "MaxAnchors") {
		t.Errorf("error %q does not flag the N bound", err)
	}
}

func TestValidateV1_RejectsEmptyAnchors(t *testing.T) {
	// ValidateV1 must reject a request missing Anchors: every wallet declares
	// its N anchor chains at registration, so a legitimate attestation cannot
	// omit them.
	b := &BtcAccountConfigured{
		AccountIndex: 0,
		Xpubs:        validV1Xpubs(),
		Threshold:    2,
		// Anchors omitted on purpose
	}
	if err := b.ValidateV1(&chaincfg.MainNetParams); err == nil {
		t.Fatal("ValidateV1 must reject a request with no Anchors")
	}
}

func TestValidateAnchorSet_AcceptsArbitraryN(t *testing.T) {
	// N is per-wallet in [1, MaxAnchors] (chosen at registration, grow-only
	// after). anchors[i] is chain i by position; no anchorIndex field.
	for _, n := range []int{1, 2, ActiveAnchors, MaxAnchors} {
		anchors := make([]AnchorBinding, n)
		for i := range anchors {
			anchors[i] = AnchorBinding{Txid: [32]byte{0xaa, byte(i)}, Vout: uint32(i)}
		}
		if err := ValidateAnchorSet(anchors); err != nil {
			t.Errorf("N=%d should be valid, got %v", n, err)
		}
	}
}

func TestValidateAnchorSet_RejectsOverflow(t *testing.T) {
	anchors := make([]AnchorBinding, MaxAnchors+1)
	for i := range anchors {
		anchors[i] = AnchorBinding{Txid: [32]byte{0xaa}, Vout: uint32(i)}
	}
	err := ValidateAnchorSet(anchors)
	if err == nil {
		t.Fatal("expected reject for > MaxAnchors elements")
	}
	if !strings.Contains(err.Error(), "MaxAnchors") {
		t.Errorf("error %q does not cite MaxAnchors", err)
	}
}

func TestValidateAnchorSet_RejectsDuplicateOutpoint(t *testing.T) {
	// Two positions backed by the same outpoint must be rejected: with
	// positional anchors that would mean two logical chains sharing one
	// physical UTXO. (The on-chain per-index scriptPubKey check would also
	// reject it, but the structural validator catches it earlier.)
	dup := AnchorBinding{Txid: [32]byte{0xaa, 0x01}, Vout: 7}
	anchors := []AnchorBinding{
		{Txid: [32]byte{0xaa, 0x00}, Vout: 0},
		dup,
		{Txid: [32]byte{0xaa, 0x02}, Vout: 2},
		dup, // same outpoint as anchors[1]
	}
	err := ValidateAnchorSet(anchors)
	if err == nil {
		t.Fatal("expected reject for duplicate anchor outpoint")
	}
	if !strings.Contains(err.Error(), "distinct UTXO") {
		t.Errorf("error %q does not flag the distinctness rule", err)
	}
}

// genesisTxidDisplay is the mainnet genesis coinbase txid as bitcoind and
// block explorers print it.
const genesisTxidDisplay = "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b"

// genesisTxidInternal is the same txid in internal byte order, taken from
// hashing the serialized transaction rather than from reversing the string, so
// the tests below do not share their conversion with the code they check.
func genesisTxidInternal() [32]byte {
	return [32]byte(chaincfg.MainNetParams.GenesisBlock.Transactions[0].TxHash())
}

func TestAnchorBinding_StringIsDisplayOrder(t *testing.T) {
	a := AnchorBinding{Txid: genesisTxidInternal(), Vout: 3}
	if got, want := a.String(), genesisTxidDisplay+":3"; got != want {
		t.Fatalf("String() = %q, want %q", got, want)
	}

	// The mistake a caller makes by hex-decoding an RPC txid without reversing
	// it: the binding then names a transaction that does not exist, and must
	// neither equal nor render as the real one.
	var display [32]byte
	if _, err := hex.Decode(display[:], []byte(genesisTxidDisplay)); err != nil {
		t.Fatal(err)
	}
	wrong := AnchorBinding{Txid: display, Vout: 3}
	if wrong == a || wrong.String() == a.String() {
		t.Fatalf("a display-order Txid reads as the internal one: %s", wrong)
	}
}

// TestAnchorErrors_NameDisplayTxid: Txid holds internal order, so printing its
// bytes as they lie would give a reader a txid no explorer or bitcoind knows.
// The anchor-set error must name the outpoint in display order.
func TestAnchorErrors_NameDisplayTxid(t *testing.T) {
	genesis := AnchorBinding{Txid: genesisTxidInternal(), Vout: 0}
	internalHex := hex.EncodeToString(genesis.Txid[:])

	check := func(t *testing.T, err error) {
		t.Helper()
		if err == nil {
			t.Fatal("expected a reject")
		}
		if !strings.Contains(err.Error(), genesisTxidDisplay+":0") {
			t.Errorf("error %q does not name the outpoint in display order", err)
		}
		if strings.Contains(err.Error(), internalHex) {
			t.Errorf("error %q names the txid in internal order", err)
		}
	}

	check(t, ValidateAnchorSet([]AnchorBinding{genesis, genesis}))
}

// Note: per-anchor value is no longer validated by ValidateAnchorSet — the
// request carries no value. The verifier reads each anchor UTXO from the
// chain and checks its value against MinAnchorValueSat (a verifier-side rule).

func TestValidateAnchorSet_AcceptsFullCapacity(t *testing.T) {
	// Pin a happy-path with the cap exactly hit (N == MaxAnchors).
	anchors := make([]AnchorBinding, MaxAnchors)
	for i := range anchors {
		anchors[i] = AnchorBinding{Txid: [32]byte{byte(i)}, Vout: uint32(i)}
	}
	if err := ValidateAnchorSet(anchors); err != nil {
		t.Fatalf("expected accept for N=MaxAnchors, got %v", err)
	}
}

// Note: the attestation carries no changeAddress and ValidateV1 does not
// validate one; the change destination is derived off-chain, so there is no
// change-address case here.
