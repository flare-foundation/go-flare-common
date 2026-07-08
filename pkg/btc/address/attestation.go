package address

import (
	"errors"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"golang.org/x/crypto/sha3"
)

// BtcAccountConfigured — single-version schema (no wallet epoch).
//
// There is exactly ONE schema version. BtcAccountConfigured carries no
// version/epoch selector (open-question #51, 2026-06-16, reverses the wallet-
// epoch design of #23 and the "epoch is the single version selector" framing
// of #30). The derivation path prefix, signature algorithm, and script
// descriptor are not on the request — the verifier derives them from
// (AccountIndex, network, Threshold, Xpubs) via DeriveV1ScriptParameters, a
// single fixed parameter set (BIP-87 path, double-SHA256-ECDSA, sorted
// wsh(multi) descriptor).
//
// Post-quantum migration is NOT an in-protocol versioning event. When BIP-361
// (consensus-level sunset of secp256k1) and BIP-360 (the P2MR address type a
// PQ wallet migrates to) land, TEE signing, the script type, and the whole
// key model change — so migration is a full new deployment: a NEW wallet
// (new walletId, new xpubs, new script type) is registered and funds are swept
// across, rather than bumping an epoch on the existing wallet. There is thus
// nothing for an on-wire version field to select; the migration replaces the
// deployment wholesale. See:
//   docs/shared/attestation-types/BtcAccountConfigured.md
//   docs/open-questions.md (#51 epoch removal; #30 request minimisation)
//   docs/future-explorations.md (PQ migration as full redeploy)
//   docs/shared/bip-tracking.md#bip-361
//
// Under open-question #30 (2026-06-04) the three previously-carried fields
// (derivationPathPrefix, sigAlgorithm, scriptDescriptor) were dropped from the
// wire: the verifier already rejected anything other than the canonical
// values, so the fields only restated information the derivation already
// encodes.

// SigAlgorithm is the bytes32 identifier the verifier emits alongside the
// attestation result and that KeyExistence proofs carry on the wire. It is
// keccak256 of the algorithm name, matching what Solidity would produce on
// the contract side.
type SigAlgorithm [32]byte

// String renders the algorithm hash in 0x-prefixed lowercase hex for logs.
func (s SigAlgorithm) String() string {
	const hexdigits = "0123456789abcdef"
	out := make([]byte, 2+2*len(s))
	out[0], out[1] = '0', 'x'
	for i, b := range s {
		out[2+2*i] = hexdigits[b>>4]
		out[2+2*i+1] = hexdigits[b&0x0f]
	}
	return string(out)
}

// Algorithm identifiers reserved by the spec.
//
// v1 verifiers MUST derive SigAlgDoubleSHA256ECDSA. The remaining values are
// name reservations for a FUTURE, SEPARATE post-quantum deployment (#51 —
// migration is a new wallet/new deploy, not an epoch bump on this one):
//   - SigAlgSecp256k1Schnorr  — future BIP-340 Schnorr / Taproot key-path
//   - SigAlgMLDSA44           — joint BIP-361 (sunset trigger) + BIP-360 (P2MR), CRYSTALS-Dilithium / ML-DSA
//   - SigAlgFalcon512         — joint BIP-361 (sunset trigger) + BIP-360 (P2MR), FALCON-512
//
// The two PQ values are reserved jointly: BIP-361 is the consensus change
// that forces migration off secp256k1; BIP-360 is the destination address
// type that carries the ML-DSA / FALCON keys. Renaming or dropping either
// reservation MUST update both bip-tracking entries together.
//
// SigAlgDoubleSHA256ECDSA was renamed from the chain-specific
// keccak256("doubleSHA256-secp256k1-ecdsa") form per open-question #24
// (2026-05-25); curve identity moved to a separate keyType field. Cross-repo
// coordination: docs/external-requirements.md (ER-01..03). The constant is
// still produced by the verifier and consumed by KeyExistence proofs and
// the FCC (keyType, signingAlgo) register even though it is no longer
// carried on the BtcAccountConfigured wire (open-question #30, 2026-06-04).
var (
	SigAlgDoubleSHA256ECDSA = keccak256("double-SHA256-ECDSA")
	SigAlgSecp256k1Schnorr  = keccak256("secp256k1-schnorr")
	SigAlgMLDSA44           = keccak256("ml-dsa-44")
	SigAlgFalcon512         = keccak256("falcon-512")
)

func keccak256(s string) SigAlgorithm {
	h := sha3.NewLegacyKeccak256()
	h.Write([]byte(s))
	var out SigAlgorithm
	copy(out[:], h.Sum(nil))
	return out
}

// Append-only anchor growth (open-question #33, 2026-06-08, partially relaxes
// #29's immutability rule). There is ONE request shape for both initial
// registration and growing N — no mode selector. A caller always submits a
// BtcAccountConfigured carrying the FULL anchor set. The consuming contract
// distinguishes the two cases from its own stored state:
//
//   - chainCountByWallet[walletId] == 0  → initial registration.
//   - chainCountByWallet[walletId]  > 0  → an extension; the new Anchors MUST
//     be a prefix-preserving superset of the registered set (existing chains
//     keep their outpoint and index), checked by ValidateAnchorGrowth. The
//     contract then bumps chainCountByWallet to len(Anchors), append-only,
//     capped at MaxAnchors.
//
// Growth therefore re-transmits the existing anchor outpoints (a handful of
// 36-byte {txid,vout} entries); the contract verifies the prefix matches. No
// "add-anchors" mode, no minimal delta body, no FromChainIdx/ToChainIdx on the
// wire — the same fields are sent every time (register-or-extend "upsert").

// BtcAccountConfigured is the request body of the BtcAccountConfigured
// attestation. Fields mirror docs/shared/attestation-types/BtcAccountConfigured.md.
//
// Open-question #30 (2026-06-04) trims the request to its irreducible inputs:
// the verifier computes derivationPathPrefix, sigAlgorithm, and
// scriptDescriptor from (AccountIndex, network, Threshold, Xpubs)
// via DeriveV1ScriptParameters — they are no longer carried on the wire.
//
// Open-question #33 (2026-06-08) makes N append-only growable via the same
// request shape — no mode selector. The Anchors field is REQUIRED and always
// carries the FULL set: a wallet declares an initial N parallel anchor chains
// at first registration (1 <= N <= MaxAnchors); each chain's anchor UTXO lives
// at the P2WSH multisig derived at index = chainIdx (chain i at index i).
// Under open-question #29 N was originally immutable; #33 relaxes that to
// grow-only — to add chains, re-attest with the SAME (AccountIndex,
// Xpubs, Threshold) and a longer Anchors that is a
// prefix-preserving superset of the registered set. The consuming contract
// checks the superset (ValidateAnchorGrowth) and bumps chainCountByWallet to
// len(Anchors), append-only, capped at MaxAnchors. FAssets registers at
// N = ActiveAnchors = 8.
//
// Under open-question #32 (2026-06-08, supersedes the off-chain selector of
// #31 for production), TeePayments.sol mirrors N at
// chainCountByWallet[walletId] (populated at addPMWMultisigAccount from
// anchors.length, bumped on the register-or-extend upsert under #33) and picks chainIdx for each
// batch via _selectChainForBatch — first-free-by-delay walk gated by the
// per-wallet anchorReuseDelayByWallet (default 720s ≈ 2 BTC blocks). Chain
// 0 carries steady-state traffic; chains 1..N-1 absorb load only when chain
// 0 is still inside its delay window. Each chain holds AT MOST ONE
// unconfirmed batch — fee escalation goes through the fee-bumping bot's
// P2A child RBF, not through a chained N+1 on the same anchor.
//
// ValidateV1 runs the full-body invariant checks; it is stateless and used
// for BOTH initial registration and a growth re-attestation. The stateful
// "is this a valid append-only extension of the stored set?" check is
// ValidateAnchorGrowth(prev, next), run by the consuming contract against its
// stored anchor set.
type BtcAccountConfigured struct {
	AccountIndex uint32
	Xpubs        []string
	Threshold    int
	// Anchors is the wallet's FULL anchor set: N entries (1 <= N <= MaxAnchors),
	// one per chain. To grow N, re-attest with a longer Anchors that is a
	// prefix-preserving superset of the registered set (open-question #33;
	// ValidateAnchorGrowth). The array is ORDER-SIGNIFICANT — anchors[i] is
	// chain i at derivation index i; order must be preserved end-to-end.
	// ValidateAnchorSet enforces N's bounds and outpoint distinctness; the
	// verifier checks each anchor's on-chain value against MinAnchorValueSat.
	Anchors []AnchorBinding

	// NOTE: there is no ChangeAddress field. The on-chain change-address pin
	// (open-question #36) was DECLINED — the change destination is neither
	// carried in the attestation nor stored on-chain. It is a pure function of
	// (Xpubs, Threshold): the P2WSH multisig at BIP-87 internal-chain leaf 0
	// (m/87'/coin'/account/1/0, account non-hardened), i.e. Internal-chain leaf
	// 0 from the account-level xpubs (DeriveAccountXpubs(Xpubs, AccountIndex)
	// then Derive(.., Internal, 0)). The proposer / TEE / relay derive it on
	// demand; the relay's Step 10 re-derives it and compares against the
	// proposal's change output.
}

// Anchor-chain layout constants.
//
//   - MaxAnchors caps the number of anchor chains a wallet may declare; it
//     also reserves derivation indices [0, MaxAnchors) for anchors (one
//     address per chain).
//   - MaxReservedAddresses is the first user-deposit derivation index:
//     [0, MaxAnchors) host anchors, [MaxAnchors, MaxReservedAddresses) are
//     reserved for future protocol features, and user deposits start at
//     MaxReservedAddresses. This is the FAssets value; MintingTagManager's
//     _reservedTagsCount mirrors it.
//   - ActiveAnchors is the FAssets default for N (chains funded at
//     registration). It is a per-protocol choice, not a hard rule —
//     ValidateAnchorSet enforces 1 <= N <= MaxAnchors and outpoint
//     distinctness.
//
// 32 chains ≈ 38k payments/h; the FAssets default of 8 ≈ 9,600 payments/h.
// N is set at registration and grows monotonically up to MaxAnchors by
// re-attesting a prefix-preserving superset (open-question #33, 2026-06-08,
// partially relaxes the immutability rule of open-question #29). See
// BtcAccountConfigured.Anchors and ValidateAnchorGrowth.
const (
	MaxAnchors           = 32
	MaxReservedAddresses = 256
	ActiveAnchors        = 8

	// FirstProtocolReservedIndex is the first derivation index in the
	// [MaxAnchors, MaxReservedAddresses) band (indices 32..255) reserved for
	// future protocol features. Non-anchor, non-user-deposit protocol
	// addresses live here — e.g. the Core Vault / HTLC k-of-n multisig
	// (timeout / refund) address — so they never collide with an anchor
	// chain's address ([0, MaxAnchors)) nor a user deposit
	// ([MaxReservedAddresses, ...)). It equals MaxAnchors by construction
	// (the band starts exactly where the anchor band ends); referencing this
	// name instead of a literal keeps the intent explicit and survives a
	// future change to the layout constants.
	FirstProtocolReservedIndex = MaxAnchors

	// PmwFloatIndex is the derivation index of the PMW redemption-float
	// address — the dedicated recipient of Core Vault → PMW float top-ups
	// (docs/workflows/btc-lifecycle.md §8, docs/components/core-vault-manager-btc.md).
	// It is the FIRST free index after the anchor band ([0, MaxAnchors)) and
	// the protocol-reserved band ([MaxAnchors, MaxReservedAddresses)), i.e. the
	// first reservation in the user-deposit band, taken by the protocol via the
	// DerivationIndexManager — subsequent reservations go to genuine user
	// deposits. It is protocol-owned (not a user deposit) and lives on the
	// external chain. Equals MaxReservedAddresses by construction (the deposit
	// band starts exactly where the protocol-reserved band ends); referencing
	// this name instead of a literal keeps the intent explicit and survives a
	// future change to the layout constants. "User deposits start at
	// MaxReservedAddresses" stays true — the float is merely the first such
	// reservation — so the _reservedTagsCount == MaxReservedAddresses invariant
	// (external-requirements ER-30) is unchanged.
	PmwFloatIndex = MaxReservedAddresses

	// MinAnchorValueSat is the minimum on-chain value (satoshis) each anchor
	// UTXO must carry. It is NOT part of the attestation — the request names
	// only the outpoint; the verifier reads the value from the *fetched*
	// on-chain UTXO and rejects the set if any anchor is below this floor.
	// 10,000 sat is the spec anchor value (open-question #6); any amount >=
	// this is accepted.
	MinAnchorValueSat = 10_000
)

// AnchorBinding names one funded anchor UTXO by its outpoint. Its position
// in the Anchors array is its chain identity and derivation index: anchors[i]
// is chain i, whose anchor address is the P2WSH multisig derived at index i.
// The array is ORDER-SIGNIFICANT — position is the only thing that ties a
// UTXO to its chain, so any layer that (de)serializes the set MUST preserve
// order (a decode->re-encode round-trip must not reorder it). There are no
// chainIdx or derivation-index fields; both equal the position. The anchor's
// value is not carried here either — the verifier reads it from the on-chain
// UTXO and checks it against MinAnchorValueSat.
//
// Txid is the genesis funding txid in *display* (big-endian) byte order —
// same convention as the wire-format BtcAccountConfigured request.
type AnchorBinding struct {
	Txid [32]byte
	Vout uint32
}

// ValidateAnchorSet enforces the invariant on the Anchors set:
//
//   - 1 <= len(anchors) <= MaxAnchors (the wallet's N, chosen at
//     registration; immutable under open-question #29, partially relaxed to
//     append-only grow-only under open-question #33 — grow by re-attesting a
//     prefix-preserving superset, MaxAnchors cap preserved; see
//     ValidateAnchorGrowth)
//   - every outpoint (Txid, Vout) is distinct — each anchor chain must be
//     backed by its own physical UTXO
//
// The array is order-significant: anchors[i] is chain i at derivation index
// i. Order is a structural invariant of the request — layers that
// (de)serialize the set MUST preserve order (verified by a round-trip test at
// the binding boundary). Per-anchor value is NOT validated here (the request
// carries no value); the verifier reads each anchor UTXO from the chain and
// checks its value against MinAnchorValueSat when confirming the set.
//
// The distinctness check is defense-in-depth: the on-chain verification step
// already forces it implicitly, because anchors[i] must pay to the index-i
// address and the per-index addresses are pairwise distinct, so one UTXO can
// satisfy at most one position. Enforcing it here rejects a duplicate at the
// structural layer with a clear error, without relying on every verifier
// implementing the per-index scriptPubKey match.
//
// Returns nil on accept. Empty input is rejected (N >= 1).
func ValidateAnchorSet(anchors []AnchorBinding) error {
	n := len(anchors)
	if n < 1 || n > MaxAnchors {
		return fmt.Errorf("attestation: anchors set has %d elements; a wallet declares N anchor chains with 1 <= N <= %d (MaxAnchors), append-only growable via re-attestation (open-questions #29 / #33)", n, MaxAnchors)
	}
	seen := make(map[AnchorBinding]int, n)
	for i, a := range anchors {
		if j, dup := seen[a]; dup {
			return fmt.Errorf("attestation: anchors[%d] and anchors[%d] reference the same outpoint %x:%d; every anchor chain must be backed by a distinct UTXO", j, i, a.Txid, a.Vout)
		}
		seen[a] = i
	}
	return nil
}

// ValidateAnchorGrowth enforces the append-only growth rule when a
// BtcAccountConfigured is re-attested for an already-registered wallet to ADD
// anchor chains (open-question #33). The same request shape is reused — the
// caller resends the FULL set — so the growth check is stateful and lives
// wherever the registered set is held: the consuming contract. This is the
// reference rule.
//
// next must be a PREFIX-PRESERVING SUPERSET of prev:
//
//   - next satisfies ValidateAnchorSet (1 <= len <= MaxAnchors, distinct);
//   - len(next) >= len(prev) — the set may only grow (or stay equal); and
//   - next[i] == prev[i] for every i < len(prev) — existing chains are
//     immutable (chain i keeps its outpoint and its index). No reindex,
//     removal, replacement, or reorder of a registered chain is allowed.
//
// len(next) == len(prev) with an equal prefix is an idempotent no-op (a replay
// of the current registration), accepted not rejected. A stale, shorter set is
// rejected — so growth is monotonic and replay-safe without a sequence number.
//
// prev may be empty/nil — the initial-registration case, which reduces to
// ValidateAnchorSet(next). The contract separately requires the re-attested
// Xpubs / Threshold to equal the stored registration.
//
// Returns nil on accept, a descriptive error on reject.
func ValidateAnchorGrowth(prev, next []AnchorBinding) error {
	if err := ValidateAnchorSet(next); err != nil {
		return err
	}
	if len(next) < len(prev) {
		return fmt.Errorf("attestation: anchor set may only grow: re-attested set has %d chains, fewer than the %d already registered (existing chains are append-only, open-question #33)", len(next), len(prev))
	}
	for i := range prev {
		if next[i] != prev[i] {
			return fmt.Errorf("attestation: anchor chain %d is immutable: re-attested outpoint %x:%d differs from the registered %x:%d (growth is append-only — no reindex/replace/remove of a registered chain, open-question #33)", i, next[i].Txid, next[i].Vout, prev[i].Txid, prev[i].Vout)
		}
	}
	return nil
}

// ValidateV1 enforces the v1 invariants on the request:
//
//   - Threshold must satisfy 1 ≤ Threshold ≤ len(Xpubs) ≤ 20 (OP_CHECKMULTISIG
//     consensus cap; see open-question #14 in docs/open-questions.md).
//   - AccountIndex must be < 2^31 (BIP-32 non-hardened range; the account
//     level is a NON-HARDENED CKDpub child of the parent xpub, so the index
//     must be a real non-hardened index — 2026-06-22 UTXO-model alignment).
//   - Each Xpub must parse as a BIP-32 extended key whose version-byte
//     network matches params, and be published at depth 2 (wallet/parent level).
//     The mainnet/testnet boundary is enforced here so a mainnet xpub
//     cannot be accepted under a testnet wallet (or vice versa) at the
//     attestation layer.
//   - Anchors must satisfy ValidateAnchorSet (1 ≤ N ≤ MaxAnchors, distinct
//     outpoints; positional — anchors[i] is chain i).
//
// The previously-carried fields DerivationPathPrefix, SigAlgorithm, and
// ScriptDescriptor were dropped under open-question #30 (2026-06-04). The
// verifier computes the corresponding canonical values from
// (AccountIndex, network, Threshold, Xpubs) via DeriveV1ScriptParameters and
// uses them for the on-chain address-derivation and anchor-set checks (step 3+
// of the verification procedure documented in
// docs/shared/attestation-types/BtcAccountConfigured.md). There is no
// version/epoch selector (open-question #51).
func (b *BtcAccountConfigured) ValidateV1(params *chaincfg.Params) error {
	if b.AccountIndex >= 1<<31 {
		return fmt.Errorf("attestation: accountIndex %d is in the BIP-32 hardened range (must be < 2^31; the account level is a non-hardened CKDpub child of the parent xpub)", b.AccountIndex)
	}
	if err := validateXpubsAndThreshold(b.Xpubs, b.Threshold, params); err != nil {
		return err
	}
	// Anchors is mandatory in v1: every wallet declares N funded anchor
	// chains at registration, so no legitimate attestation can omit them.
	// ValidateAnchorSet enforces 1 <= N <= MaxAnchors and distinct outpoints
	// (the array is order-significant, anchors[i] is chain i); the verifier
	// checks each anchor's on-chain value against MinAnchorValueSat.
	if err := ValidateAnchorSet(b.Anchors); err != nil {
		return err
	}
	// No change-address validation: the on-chain change-address pin
	// (open-question #36) was declined, so the attestation carries no change
	// address. The change destination is derived off-chain by the proposer /
	// relay from (Xpubs, Threshold) at internal-chain leaf 0 — there is
	// nothing to validate here.
	return nil
}

// DerivedV1Parameters bundles the canonical (derivationPathPrefix,
// sigAlgorithm, scriptDescriptor) triple a v1 verifier computes from a
// BtcAccountConfigured request. Carried in this Go type only — the wire
// format does not transport these values (open-question #30, 2026-06-04).
//
// scriptDescriptor is the full BIP-380 form with the 8-character checksum
// suffix; downstream tooling that wants to display or copy the descriptor
// into Sparrow / Specter / hardware coordinators uses this value verbatim.
type DerivedV1Parameters struct {
	DerivationPathPrefix string
	SigAlgorithm         SigAlgorithm
	ScriptDescriptor     string
}

// DeriveV1ScriptParameters computes the v1 derivation triple — the single
// fixed parameter set (no version/epoch selector; open-question #51). The
// procedure is:
//
//  1. coin_type is 0' on mainnet, 1' on signet/testnet/regtest.
//  2. derivationPathPrefix = m/87'/<coin_type>'/<AccountIndex> (BIP-87; the
//     account level is NON-HARDENED — see DeriveV1PathPrefix).
//  3. sigAlgorithm = keccak256("double-SHA256-ECDSA").
//  4. scriptDescriptor = wsh(sortedmulti(Threshold,
//     xpub_1/<AccountIndex>/0/*, ..., xpub_n/<AccountIndex>/0/*)) plus the
//     8-character BIP-380 checksum suffix. The xpubs are the published
//     wallet-level (parent) keys, so the descriptor carries the non-hardened
//     <AccountIndex> step explicitly (2026-06-22, UTXO-model alignment).
//
// Requires that b has already cleared ValidateV1 (well-formed Threshold,
// Xpubs, network match); calling this on a malformed request is undefined.
// Returns an error if descriptor checksum computation fails (only possible on
// an invalid bech32 char in an xpub, which would have been caught by
// validateXpubsAndThreshold).
//
// Pre-spec wallets bound under BIP-44 (m/44'/...) are not accepted: the spec
// is still DRAFT and there are no live wallets to grandfather. (There is no
// version/epoch field to select a BIP-44 variant; a future scheme change is a
// separate deployment — open-question #51.)
func (b *BtcAccountConfigured) DeriveV1ScriptParameters(params *chaincfg.Params) (DerivedV1Parameters, error) {
	prefix, err := DeriveV1PathPrefix(params, b.AccountIndex)
	if err != nil {
		return DerivedV1Parameters{}, err
	}

	parts := make([]string, 0, len(b.Xpubs))
	for _, x := range b.Xpubs {
		// Parent (wallet-level) xpub → non-hardened account step at AccountIndex
		// → external chain (0) → leaf (*). The explicit <AccountIndex> step is
		// what makes one published parent xpub serve every account.
		parts = append(parts, fmt.Sprintf("%s/%d/0/*", x, b.AccountIndex))
	}
	expr := fmt.Sprintf("wsh(sortedmulti(%d,%s))", b.Threshold, strings.Join(parts, ","))
	csum, err := ComputeDescriptorChecksum(expr)
	if err != nil {
		return DerivedV1Parameters{}, fmt.Errorf("attestation: derive scriptDescriptor checksum: %w", err)
	}
	return DerivedV1Parameters{
		DerivationPathPrefix: prefix,
		SigAlgorithm:         SigAlgDoubleSHA256ECDSA,
		ScriptDescriptor:     expr + "#" + csum,
	}, nil
}

// bip87CoinType returns the BIP-87 coin_type segment for params: 0 on
// mainnet, 1 on signet / testnet3 / regtest. The verifier picks the
// network from the FDC2 sourceId header; this helper centralises the
// mapping so v1 derivation has one place to update if a new test network
// joins.
func bip87CoinType(params *chaincfg.Params) (uint32, error) {
	if params == nil {
		return 0, errors.New("attestation: params required for coin_type derivation")
	}
	switch params.Net {
	case chaincfg.MainNetParams.Net:
		return 0, nil
	case chaincfg.SigNetParams.Net,
		chaincfg.TestNet3Params.Net,
		chaincfg.RegressionNetParams.Net:
		return 1, nil
	default:
		return 0, fmt.Errorf("attestation: unsupported network %q for v1 coin_type derivation", params.Name)
	}
}

// DeriveV1PathPrefix returns the v1 BIP-87 account-level derivation
// path prefix for a wallet on params at accountIndex:
// m/87'/<coin_type>'/<accountIndex> (coin_type 0 on mainnet, 1 on
// signet / testnet3 / regtest). The account level is NON-HARDENED (no trailing
// apostrophe) — it is a public CKDpub child of the published wallet-level
// (parent) xpub at accountIndex (2026-06-22, aligns with the deployed
// flare-smart-contracts-v2 UTXO model). This is the single source of truth for
// the v1 path shape: DeriveV1ScriptParameters uses it, and off-chain components
// (e.g. the proposer's startup config validation) call it so their expected
// path cannot drift from what the verifier derives. accountIndex must be in the
// BIP-32 non-hardened range (< 2^31) — and now genuinely IS a non-hardened
// derivation index, not merely a hardened path component; params must be a
// supported network.
func DeriveV1PathPrefix(params *chaincfg.Params, accountIndex uint32) (string, error) {
	if accountIndex >= 1<<31 {
		return "", fmt.Errorf("attestation: accountIndex %d is in the BIP-32 hardened range (must be < 2^31 for non-hardened account derivation)", accountIndex)
	}
	coinType, err := bip87CoinType(params)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("m/87'/%d'/%d", coinType, accountIndex), nil
}

// validateXpubsAndThreshold checks the multisig key set against the BIP-87
// limits enforced by Derive (n ≤ 20, 1 ≤ k ≤ n) and BIP-32 well-formedness
// of each xpub, including the version-byte network identity check that
// Derive also enforces. Each xpub must be published at the WALLET (parent)
// level, i.e. BIP-32 depth 2 (m/purpose'/coin'): the account level is a
// NON-HARDENED child at accountIndex (2026-06-22, aligns with the deployed
// flare-smart-contracts-v2 UTXO model — IPMWMultisigUtxoConfigured publishes
// wallet-level parent xpubs). The v1 derivation table requires depth 2 so
// DeriveV1ScriptParameters / DeriveAccountXpubs produce a coherent derivation
// chain. Validating here catches mis-bound attestations before any
// tx-construction code touches the wallet.
func validateXpubsAndThreshold(xpubs []string, threshold int, params *chaincfg.Params) error {
	if params == nil {
		return errors.New("attestation: params required for xpub network validation")
	}
	n := len(xpubs)
	if n == 0 {
		return errors.New("attestation: xpubs must contain at least one key")
	}
	if n > 20 {
		// Mirrored in scripts/spec-tests/boundaries/boundaries_test.go
		// (validateMultisigKeyCount). Keep the error string in sync.
		return fmt.Errorf("attestation: n=%d exceeds OP_CHECKMULTISIG limit of 20 (open-question #14)", n)
	}
	if threshold < 1 || threshold > n {
		return fmt.Errorf("attestation: threshold %d out of range for n=%d (must satisfy 1 ≤ k ≤ n)", threshold, n)
	}
	seen := make(map[string]int, n)
	for i, xs := range xpubs {
		key, err := hdkeychain.NewKeyFromString(xs)
		if err != nil {
			return fmt.Errorf("attestation: xpubs[%d]: %w", i, err)
		}
		if !key.IsForNet(params) {
			return fmt.Errorf("attestation: xpubs[%d] version byte does not match network %q", i, params.Name)
		}
		if key.IsPrivate() {
			return fmt.Errorf("attestation: xpubs[%d] is a private extended key; publicKeys must carry public (xpub/tpub) keys only", i)
		}
		if d := key.Depth(); d != 2 {
			return fmt.Errorf("attestation: xpubs[%d] is at BIP-32 depth %d, expected wallet-level (parent) depth 2 (m/purpose'/coin'); the account level is a non-hardened child at accountIndex", i, d)
		}
		// Reject duplicate signers. A repeated key collapses signer
		// independence — a k-of-n with duplicates can be satisfied by fewer
		// than k independent parties, so it is not a genuine k-of-n multisig.
		// Distinct parent xpubs yield distinct derived leaf pubkeys, so this
		// suffices; Derive additionally guards the derived set. Normalized via
		// canonical re-serialization so encoding variants cannot slip through.
		norm := key.String()
		if j, dup := seen[norm]; dup {
			return fmt.Errorf("attestation: xpubs[%d] duplicates xpubs[%d]; every signer must be a distinct key", i, j)
		}
		seen[norm] = i
	}
	return nil
}
