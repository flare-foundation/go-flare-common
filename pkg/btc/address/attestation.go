package address

import (
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
)

// BtcAccountConfigured is the request body of the BtcAccountConfigured
// attestation.
//
// The request carries only its irreducible inputs. The derivation path,
// signature algorithm and script descriptor are not on it: there is one fixed
// parameter set (BIP-87 path with a non-hardened account step,
// double-SHA256-ECDSA, sorted wsh(multi) script), so carrying them would only
// restate what (AccountIndex, network, Threshold, Xpubs) already determine.
// There is no version or epoch selector either: a change of script type or key
// model (for example a post-quantum migration) is a new wallet with new keys,
// and funds are swept across to it.
//
// Anchors always carries the FULL anchor set. A wallet declares N parallel
// anchor chains at registration (1 <= N <= MaxAnchors); chain i's anchor UTXO
// lives at the P2WSH multisig derived at external index i. N may only grow:
// to add chains, the same (AccountIndex, Xpubs, Threshold) is re-attested
// with a longer Anchors that keeps the registered set as its prefix, and the
// contract holding the registered set checks that.
//
// ValidateV1 runs the stateless full-body checks, for both a first
// registration and a growth re-attestation.
type BtcAccountConfigured struct {
	AccountIndex uint32
	Xpubs        []string
	Threshold    int
	// Anchors is the wallet's FULL anchor set: N entries (1 <= N <= MaxAnchors),
	// one per chain. The array is ORDER-SIGNIFICANT — anchors[i] is chain i at
	// derivation index i; order must be preserved end-to-end.
	// ValidateAnchorSet enforces N's bounds and outpoint distinctness; the
	// verifier checks each anchor's on-chain value against MinAnchorValueSat.
	Anchors []AnchorBinding

	// There is no ChangeAddress field: the change destination is neither
	// attested nor stored on chain. It is a pure function of (Xpubs,
	// Threshold, AccountIndex): the P2WSH multisig at internal-chain leaf 0
	// (m/87'/coin'/account/1/0), i.e. Derive(DeriveAccountXpubs(Xpubs,
	// AccountIndex), Threshold, Internal, 0). Whoever builds or checks a batch
	// derives it on demand.
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
// re-attesting a prefix-preserving superset (see BtcAccountConfigured).
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
	// address — the dedicated recipient of Core Vault → PMW float top-ups.
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
	// is unchanged.
	PmwFloatIndex = MaxReservedAddresses

	// MinAnchorValueSat is the minimum on-chain value (satoshis) each anchor
	// UTXO must carry. It is NOT part of the attestation — the request names
	// only the outpoint; the verifier reads the value from the *fetched*
	// on-chain UTXO and rejects the set if any anchor is below this floor.
	// 10,000 sat is the anchor value the protocol funds; any amount >= this
	// is accepted.
	MinAnchorValueSat = 10_000
)

// AnchorBinding names one funded anchor UTXO by its outpoint. Its position
// in the Anchors array is its chain identity and derivation index: anchors[i]
// is chain i, whose anchor address is the P2WSH multisig derived at index i.
// The array is ORDER-SIGNIFICANT — position is the only thing that ties a
// UTXO to its chain, so any layer that (de)serializes the set MUST preserve
// order (a decode->re-encode round-trip must not reorder it). There are no
// anchorIndex or derivation-index fields; both equal the position. The anchor's
// value is not carried here either — the verifier reads it from the on-chain
// UTXO and checks it against MinAnchorValueSat.
//
// Txid is the genesis funding txid in INTERNAL byte order: the order its bytes
// take in a spending input's outpoint and in a chainhash.Hash, the reverse of
// what bitcoind and block explorers print. The BtcAccountConfigured request
// carries it in that order, and so do the anchors the contract stores and the
// registry returns, as every Bitcoin txid held as bytes32 is. Display order
// belongs only at a boundary — a bitcoind RPC argument such as gettxout's txid,
// or a log line — and is converted there, once. String renders it.
type AnchorBinding struct {
	Txid [32]byte
	Vout uint32
}

// String renders the outpoint as bitcoind and explorers print it,
// "<display-order txid>:<vout>", so a txid in an error or a log line can be
// looked up as written.
func (a AnchorBinding) String() string {
	return wire.OutPoint{Hash: a.Txid, Index: a.Vout}.String()
}

// ValidateAnchorSet enforces the invariant on the Anchors set:
//
//   - 1 <= len(anchors) <= MaxAnchors (the wallet's N, chosen at
//     registration and only ever grown, by re-attesting a prefix-preserving
//     superset)
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
		return fmt.Errorf("attestation: anchors set has %d elements; a wallet declares N anchor chains with 1 <= N <= %d (MaxAnchors), append-only growable via re-attestation", n, MaxAnchors)
	}
	seen := make(map[AnchorBinding]int, n)
	for i, a := range anchors {
		if j, dup := seen[a]; dup {
			return fmt.Errorf("attestation: anchors[%d] and anchors[%d] reference the same outpoint %s; every anchor chain must be backed by a distinct UTXO", j, i, a)
		}
		seen[a] = i
	}
	return nil
}

// ValidateV1 enforces the v1 invariants on the request:
//
//   - Threshold must satisfy 1 ≤ Threshold ≤ len(Xpubs) ≤ 20 (the
//     OP_CHECKMULTISIG consensus cap).
//   - AccountIndex must be < 2^31 (BIP-32 non-hardened range; the account
//     level is a NON-HARDENED CKDpub child of the parent xpub, so the index
//     must be a real non-hardened index).
//   - Each Xpub must parse as a BIP-32 extended key whose version-byte
//     network matches params, and be published at depth 2 (wallet/parent level).
//     The mainnet/testnet boundary is enforced here so a mainnet xpub
//     cannot be accepted under a testnet wallet (or vice versa) at the
//     attestation layer.
//   - Anchors must satisfy ValidateAnchorSet (1 ≤ N ≤ MaxAnchors, distinct
//     outpoints; positional — anchors[i] is chain i).
//
// The derivation path and script are not on the request; the verifier derives
// them from (AccountIndex, network, Threshold, Xpubs) for its address and
// anchor-set checks.
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
	// No change-address validation: the attestation carries no change
	// address. The change destination is derived off-chain from (Xpubs,
	// Threshold) at internal-chain leaf 0.
	return nil
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
// (parent) xpub at accountIndex. This is the single source of truth for the v1
// path shape: off-chain components (e.g. a proposer validating its config)
// call it so their expected path cannot drift from what the verifier derives.
// accountIndex must be in the BIP-32 non-hardened range (< 2^31); params must
// be a supported network.
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
// NON-HARDENED child at accountIndex (IBtcAccountConfigured publishes
// wallet-level parent xpubs). Depth 2 is required so DeriveAccountXpubs
// produces a coherent derivation chain. Validating here catches mis-bound attestations before any
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
		return fmt.Errorf("attestation: n=%d exceeds OP_CHECKMULTISIG limit of 20", n)
	}
	if threshold < 1 || threshold > n {
		return fmt.Errorf("attestation: threshold %d out of range for n=%d (must satisfy 1 ≤ k ≤ n)", threshold, n)
	}
	coin, err := bip87CoinType(params)
	if err != nil {
		return err
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
		// The depth says the level; the child number says WHICH coin. A key at
		// m/87'/0' re-serialized as a tpub passes the version check on a test
		// network and is still a mainnet key: the coin level is hardened, so it is
		// the child number, not the version bytes, that separates one network's
		// keys from another's. The chain holds a wallet's key type to the
		// deployment's network, and the key type names this coin.
		if ci, want := key.ChildIndex(), hdkeychain.HardenedKeyStart+coin; ci != want {
			return fmt.Errorf("attestation: xpubs[%d] is published under child number %d, not coin %d' (%d): the wallet-level key on %s is m/87'/%d'", i, ci, coin, want, params.Name, coin)
		}
		// Reject duplicate signers, comparing the PUBLIC KEY UP TO SIGN — its X
		// coordinate — alone. A k-of-n is genuine only if its n signers are n
		// key holders, and two xpubs sharing an X coordinate are one holder
		// whatever else differs. With one public key, whoever holds the private
		// key derives the children of both, since the chain code is public, so
		// comparing the whole serialization (version, depth, fingerprint, child
		// number, chain code) would let one party fill two slots. A key and its
		// negation differ only in the parity byte, and their private keys are d
		// and -d, so whoever holds one holds the other. Derive's leaf check and
		// pkg/btc/htlc's catch neither, because both pairs derive different
		// children. One holder behind two unrelated keys is invisible to any
		// comparison; that is the registry's to prevent.
		pub, err := key.ECPubKey()
		if err != nil {
			return fmt.Errorf("attestation: xpubs[%d]: %w", i, err)
		}
		norm := string(pub.SerializeCompressed()[1:])
		if j, dup := seen[norm]; dup {
			return fmt.Errorf("attestation: xpubs[%d] duplicates xpubs[%d]; every signer must be a distinct key", i, j)
		}
		seen[norm] = i
	}
	return nil
}
