// Package csp holds the canonical form of a Consensus-Sequenced Proposals
// envelope: the object a proposer publishes to the DAL, the data providers vote
// on, and the TEE machines sign against.
//
// Four independent components compute this hash — the proposer, the FDC2
// verifier, the relay client and the machine. If any two of them compute it
// differently the system does not fail loudly; it fails as a proposal that
// never reaches threshold, or as a machine refusing to sign something the chain
// already finalized. So the encoding lives here, once, and the consumers import
// it rather than reimplement it.
//
// # Why abi.encode
//
// A hash commitment is only sound if the encoding is INJECTIVE: no two distinct
// messages may produce the same bytes. abi.encode is injective for a fixed type
// signature by construction — its head/tail scheme length-prefixes every
// dynamic value and offsets every one of them. A hand-rolled layout can be
// injective too, but only because its author was careful, and that care has to
// be re-established by review every time a field is added. The classic failure
// is in the same family: abi.encodePacked("a","bc") and abi.encodePacked("ab","c")
// are the same bytes, so the hash cannot tell the two apart.
//
// # What is deliberately absent
//
// No Bitcoin address appears in an envelope in any form. bech32 is
// case-insensitive but its checksum is not, so mixed-case and lowercase forms
// hash differently, and the HRP differs by network — an ambiguity no encoding
// removes, because the disagreement is about the value rather than its
// serialization. Recipients live inside RawUnsignedTx as raw scriptPubKeys, and
// the account is named by its hash (below), never by its address string.
//
// # The wallet's shape travels with the proposal
//
// The signer set — the n parent xpubs and k — is carried in the envelope, so
// every consumer reads it from the one object the chain finalized. Provisioned
// separately, the verifier would read the registry, the machines a key binding
// and the facilitator its configuration, and the three copies would agree only
// as long as nothing made them disagree.
//
// Carrying the keys does not let a proposer choose them. The FDC2 verifier
// requires them to equal the registry's, so a proposal naming any other set
// fails the vote; once it is finalized, the package hash binds exactly the set
// the data providers checked. Inputs remain COORDINATES — every party still
// derives each script itself, from these keys — which is what keeps a proposer
// from supplying a script.
//
// There is no coin-type field. Each key is the wallet level m/87'/coin', and a
// BIP-32 key records the index of its last derivation step in its own bytes, so
// the coin is read out of the keys (CoinType). A separate field would be a
// second copy of one fact, and a consistency check to keep them equal.
//
// # The account in the chain's own terms
//
// The WalletPayments diamond keys a wallet by (registry, walletId) — a TEE
// wallet and a custodian wallet live in different registries, and a walletId is
// unique only within one — and an account by its hash,
// keccak256(abi.encode(sourceId, accountAddress)), which is also the topic of
// every account event. The envelope carries WalletRegistry and AccountHash, so
// the finalized hash binds the account exactly as the chain stores it, and a
// consumer that sees an event compares a topic rather than looking an account
// up.
//
// The escrow carries no timeout coordinates. The diamond fixes the escrow's
// timeout branch at the wallet's k-of-n at (EscrowTimeoutChain,
// EscrowTimeoutIndex), so a field for it could only ever hold that constant.
// The remaining terms take the contract's names (counterpartyPubKey,
// expiresAt).
package csp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/flare-foundation/go-flare-common/pkg/abicoder"
	"github.com/flare-foundation/go-flare-common/pkg/signing"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/tee"
)

// The escrow's timeout branch is always the wallet's own k-of-n at this
// derivation coordinate, as the diamond fixes it; an escrow input carries it.
const (
	EscrowTimeoutChain uint8  = 1
	EscrowTimeoutIndex uint32 = 0
)

// MaxSigners is OP_CHECKMULTISIG's consensus limit on n.
const MaxSigners = 20

// XpubLen is the BIP-32 serialization of an extended key: version(4) ‖
// depth(1) ‖ parent fingerprint(4) ‖ child number(4) ‖ chain code(32) ‖
// key(33). Carried raw rather than as base58: fixed width, no case or checksum
// ambiguity in the hash preimage.
const XpubLen = 78

// parentXpubDepth is the depth of the published wallet-level key m/87'/coin'.
// Only at this depth is a key's child number its coin type.
const parentXpubDepth = 2

// SLIP-44 coin types this package accepts, and the extended-public-key version
// bytes each one requires. Every Bitcoin test network shares coin type 1 and
// the tpub version bytes with it.
const (
	CoinTypeBitcoin   uint32 = 0
	CoinTypeTestnet   uint32 = 1
	hardenedKeyStart  uint32 = 0x80000000
	mainnetPubVersion uint32 = 0x0488b21e // xpub
	testnetPubVersion uint32 = 0x043587cf // tpub (testnet3, signet, regtest)
)

// MaxInputs bounds the input array so a malformed envelope cannot be used to
// force unbounded work in a verifier. A TRUC batch is capped at 10,000 vB,
// which admits far fewer inputs than this.
const MaxInputs = 1024

// Input carries what the transaction bytes cannot supply on their own: the
// spent value, because BIP-143 signs over it and it is not present in the
// input; and the derivation coordinates, because a scriptPubKey does not reveal
// the path that produced it. It is TeeStructs.CspInput, as generated:
//
//   - Txid is the outpoint's transaction id in INTERNAL byte order — the same
//     order the raw transaction uses, not the reversed display form.
//   - Chain is 0 for external, 1 for internal (change), per BIP-32.
//   - IsEscrow marks an input whose witness script is the envelope's ESCROW
//     script rather than the plain k-of-n multisig at (Chain, Index). The
//     coordinates still name the KEY each signer uses — the escrow's timeout
//     branch is the wallet's own k-of-n — so this changes which script is
//     committed to, not who may spend, and an escrow input's coordinates must
//     be (EscrowTimeoutChain, EscrowTimeoutIndex). Without the flag a signer
//     would compute a sighash over the wrong script and produce a signature no
//     one can use.
type Input = tee.TeeStructsCspInput

// Escrow is the HTLC an ESCROW_CREATE pays into or an ESCROW_RECLAIM spends:
// TeeStructs.CspEscrow, as generated, whose fields are the diamond's
// BtcEscrowTerms under the same names.
//
// Carried in the envelope, so it is covered by the proposal hash and therefore
// by everything the chain finalized. A signer rebuilds the script from these
// fields plus the envelope's signer set — never from a script the proposal
// supplies. The multisig half is therefore the wallet's own: the verifier
// required ParentXpubs to equal the registry's before the hash was finalized,
// and a machine signs only if its own key is among them, so a proposal naming
// keys the wallet does not hold is refused before any signature exists.
//
//   - PreimageHash is SHA256(preimage) for the hash path.
//   - CounterpartyPubKey is the compressed key that can take the hash path.
//   - ExpiresAt is the UNIX timestamp after which the wallet may reclaim: the
//     script's CLTV locktime.
//
// The timeout branch's keys are at (EscrowTimeoutChain, EscrowTimeoutIndex).
type Escrow = tee.TeeStructsCspEscrow

// Envelope is the proposal object: TeeStructs.CspEnvelope, as generated from
// the struct ABI. The layout is defined there and changes there first; this
// package keeps only the package framing, the hashing and the checks by hand.
//
// There is no version field. A decoder accepts only the canonical encoding of
// this tuple — decoded and encoded again, the bytes must be identical (Decode)
// — and that is what a version check was for: an appended field leaves a tail
// the re-encoding does not reproduce, and a changed or reordered field fails to
// decode or to re-encode.
//
//   - RawUnsignedTx is the NON-WITNESS serialization, which is also the txid
//     preimage — so committing to the envelope pins the txid directly rather
//     than by derivation. This is a serialization choice, not a
//     transaction-type choice: every transaction in this design is SegWit.
//   - Escrow is present only for the ESCROW instruction kinds. A zero value
//     means the proposal creates or spends no HTLC, which is every payment and
//     every consolidation.
//   - ParentXpubs are the wallet's n published keys at m/87'/coin', raw
//     XpubLen bytes each, in REGISTRY order. Order does not change any script —
//     BIP-67 sorts the leaf keys — but it is committed, so it has one canonical
//     value. The keys also name the coin (CoinType).
//   - Threshold is k in the k-of-n.
//   - WalletRegistry is the contract the wallet is registered in (the TEE
//     manager or the custodian wallet manager). A walletId is unique only
//     within its registry.
//   - AccountHash is keccak256(abi.encode(sourceId, accountAddress)): the key
//     the diamond stores the account under, and the topic of its events.
//
// The field order is the struct's, and it is fixed: changing it changes every
// proposalHash ever computed, so it is a breaking change to the whole system
// and not a refactor. TestEnvelopeLayout pins it word by word.
type Envelope tee.TeeStructsCspEnvelope

// envelopeArgument is the generated CspEnvelope's ABI type.
var envelopeArgument = tee.StructArg[tee.CspEnvelope]

func init() {
	bytes32Ty, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		panic("building bytes32 type: " + err.Error())
	}
	stringTy, err := abi.NewType("string", "", nil)
	if err != nil {
		panic("building string type: " + err.Error())
	}
	accountKeyArgs = abi.Arguments{{Type: bytes32Ty}, {Type: stringTy}}
}

// accountKeyArgs is (bytes32 sourceId, string accountAddress), the tuple the
// diamond hashes into an account key.
var accountKeyArgs abi.Arguments

// Validate rejects envelopes that are structurally unusable. It is called by
// Encode, so no caller can hash something it would refuse.
//
// It checks the instruction, its inputs and escrow terms, then the signer set
// and finally the account.
func (e Envelope) Validate() error {
	// Positions start at 1 on the channel; 0 means "none" everywhere it is
	// read, so an envelope for position 0 fills no instruction.
	if e.SequencePosition == 0 {
		return errors.New("sequence position 0 names no instruction; positions start at 1")
	}
	if len(e.RawUnsignedTx) == 0 {
		return errors.New("rawUnsignedTx is empty")
	}
	if len(e.Inputs) == 0 {
		return errors.New("envelope has no inputs")
	}
	if len(e.Inputs) > MaxInputs {
		return fmt.Errorf("%d inputs exceeds the %d cap", len(e.Inputs), MaxInputs)
	}
	escrowInputs := 0
	for i, in := range e.Inputs {
		if in.Chain > 1 {
			return fmt.Errorf("input %d has chain %d, want 0 (external) or 1 (internal)", i, in.Chain)
		}
		if in.IsEscrow {
			escrowInputs++
			if i == 0 {
				return errors.New("input 0 is the anchor and can never be an escrow output")
			}
			if in.Chain != EscrowTimeoutChain || in.Index != EscrowTimeoutIndex {
				return fmt.Errorf("escrow input %d is at (%d, %d); an escrow's timeout branch is at (%d, %d)",
					i, in.Chain, in.Index, EscrowTimeoutChain, EscrowTimeoutIndex)
			}
		}
	}
	if escrowInputs > 1 {
		return fmt.Errorf("%d escrow inputs; one instruction reclaims one escrow", escrowInputs)
	}
	// Terms without an output to apply them to, or an input with no terms to
	// rebuild its script from, are both unusable rather than merely odd.
	hasTerms := e.Escrow.ExpiresAt != 0 || e.Escrow.PreimageHash != [32]byte{} || len(e.Escrow.CounterpartyPubKey) > 0
	if escrowInputs == 1 && !hasTerms {
		return errors.New("an escrow input needs the escrow terms to rebuild its witness script")
	}
	if hasTerms && len(e.Escrow.CounterpartyPubKey) != 33 {
		return fmt.Errorf("escrow counterparty key is %d bytes, want a 33-byte compressed key", len(e.Escrow.CounterpartyPubKey))
	}
	if err := e.validateSigners(); err != nil {
		return err
	}
	// The account is named in the chain's terms or not at all: a zero registry
	// or hash would leave the hash binding nothing a consumer can compare.
	if e.WalletRegistry == (common.Address{}) {
		return errors.New("envelope names no wallet registry")
	}
	if e.AccountHash == ([32]byte{}) {
		return errors.New("envelope names no account hash")
	}
	return nil
}

// AccountHashOf is the diamond's account key for (sourceId, accountAddress):
// keccak256(abi.encode(sourceId, accountAddress)), the value AccountHash must
// carry.
func AccountHashOf(sourceID [32]byte, accountAddress string) [32]byte {
	packed, err := accountKeyArgs.Pack(sourceID, accountAddress)
	if err != nil {
		// A bytes32 and a string always pack.
		panic("packing account key: " + err.Error())
	}
	return crypto.Keccak256Hash(packed)
}

// Encode returns the canonical bytes: abi.encode(CspEnvelope).
func (e Envelope) Encode() ([]byte, error) {
	if err := e.Validate(); err != nil {
		return nil, err
	}
	packed, err := abicoder.Encode(envelopeArgument, tee.TeeStructsCspEnvelope(e))
	if err != nil {
		return nil, fmt.Errorf("encoding envelope: %w", err)
	}
	return packed, nil
}

// Hash is the proposalHash: domain-separated with the house prefix scheme, so
// it can never collide with another hash in the system, and bound to chainID so
// a proposal from one network cannot be replayed onto another.
func (e Envelope) Hash(chainID uint64) (common.Hash, error) {
	encoded, err := e.Encode()
	if err != nil {
		return common.Hash{}, err
	}
	h, err := signing.NewPayload(signing.CSPProposal, chainID, [32]byte(crypto.Keccak256Hash(encoded))).Hash()
	if err != nil {
		return common.Hash{}, fmt.Errorf("hashing envelope: %w", err)
	}
	return h, nil
}

// Decode parses canonical bytes back into an Envelope.
//
// Only the canonical encoding is accepted: the bytes are decoded, encoded
// again, and must come back identical. This is the envelope's only format
// guard — there is no version. Trailing bytes, an appended word (the shape of
// an envelope with a field this layout does not have), non-minimal padding and
// aliased offsets all re-encode to something else and are refused. If two byte
// strings could decode to one message, the hash would no longer identify the
// message.
func Decode(b []byte) (Envelope, error) {
	raw, err := abicoder.Decode[tee.TeeStructsCspEnvelope](envelopeArgument, b)
	if err != nil {
		return Envelope{}, fmt.Errorf("decoding envelope: %w", err)
	}
	out := Envelope(raw)
	// Encode validates, so nothing this package would refuse to encode
	// decodes either. abicoder.Decode already re-encodes; the comparison is
	// repeated here so the guard does not rest on a helper's internals.
	round, err := out.Encode()
	if err != nil {
		return Envelope{}, err
	}
	if !bytes.Equal(round, b) {
		return Envelope{}, fmt.Errorf("envelope is not canonical: %d bytes in, %d re-encoded", len(b), len(round))
	}
	return out, nil
}

// validateSigners checks the wallet's shape is one a k-of-n P2WSH can have, and
// that every key is a PUBLIC wallet-level key of one agreed coin (CoinType).
func (e Envelope) validateSigners() error {
	n := len(e.ParentXpubs)
	if n == 0 {
		return errors.New("envelope names no signers")
	}
	if n > MaxSigners {
		return fmt.Errorf("%d signers exceeds the OP_CHECKMULTISIG limit of %d", n, MaxSigners)
	}
	if e.Threshold < 1 || e.Threshold > uint32(n) { // n <= MaxSigners, so the conversion is exact
		return fmt.Errorf("threshold %d is not a k of %d", e.Threshold, n)
	}
	if _, err := e.CoinType(); err != nil {
		return err
	}
	// A signer is its public key up to sign — the X coordinate, bytes 46–78 —
	// and not the whole serialization. A k-of-n is genuine only if its n
	// signers are n key holders, and two entries sharing an X coordinate are
	// one holder whatever else differs. With one point under a different chain
	// code, whoever holds the private key derives the children of both, since
	// the chain code is public. P and -P differ only in the parity byte, and
	// their private keys are d and -d, so whoever holds one holds the other.
	// Comparing all 78 bytes, or all 33 of the point, would let one party fill
	// two slots, and the leaf checks in pkg/btc/address and pkg/btc/htlc catch
	// neither, because both pairs derive different children. An X coordinate
	// has one encoding, so equal keys are equal bytes here. One holder behind
	// two unrelated keys is invisible to any comparison; that is the
	// registry's to prevent.
	seen := make(map[string]struct{}, n)
	for i, x := range e.ParentXpubs {
		if x[45] != 0x02 && x[45] != 0x03 {
			return fmt.Errorf("parentXpubs[%d] does not carry a compressed public key", i)
		}
		pub := string(x[46:])
		if _, dup := seen[pub]; dup {
			return fmt.Errorf("parentXpubs[%d] repeats an earlier signer", i)
		}
		seen[pub] = struct{}{}
	}
	return nil
}

// CoinType returns the SLIP-44 coin type the envelope's keys agree on.
//
// There is no field for it: a key at m/87'/coin' names its coin in its own
// child number (bytes 9–12), which the machine's BIP-32 library wrote when it
// generated the key. Every key must name the same hardened coin, at depth 2,
// with the version bytes of that coin — xpub for 0, tpub for 1. Any other coin
// is refused until something supports it.
//
// The child number is a label, not a proof: it does not enter child
// derivation, so a key with a rewritten one still derives. It is trustworthy
// here because the verifier requires the keys to equal the registry's — what
// the machines themselves published — and a machine signs only when its own
// key's bytes are among them.
//
// Consumers call this rather than reading bytes 9–12 themselves.
func (e Envelope) CoinType() (uint32, error) {
	if len(e.ParentXpubs) == 0 {
		return 0, errors.New("envelope names no signers, so no coin type")
	}
	var coin uint32
	for i, x := range e.ParentXpubs {
		c, err := walletKeyCoinType(x)
		if err != nil {
			return 0, fmt.Errorf("parentXpubs[%d]: %w", i, err)
		}
		if i == 0 {
			coin = c
			continue
		}
		if c != coin {
			return 0, fmt.Errorf("the keys disagree on the coin: parentXpubs[0] is m/87'/%d', parentXpubs[%d] is m/87'/%d'",
				coin, i, c)
		}
	}
	return coin, nil
}

// walletKeyCoinType reads the coin a wallet-level key names, and requires the
// key's version bytes to be that coin's.
func walletKeyCoinType(x []byte) (uint32, error) {
	if len(x) != XpubLen {
		return 0, fmt.Errorf("%d bytes, want %d", len(x), XpubLen)
	}
	if x[4] != parentXpubDepth {
		return 0, fmt.Errorf("at depth %d, want %d (m/87'/coin')", x[4], parentXpubDepth)
	}
	child := binary.BigEndian.Uint32(x[9:13])
	if child < hardenedKeyStart {
		return 0, fmt.Errorf("child number %d is not hardened; the wallet level is m/87'/coin'", child)
	}
	coin := child - hardenedKeyStart
	var want uint32
	switch coin {
	case CoinTypeBitcoin:
		want = mainnetPubVersion
	case CoinTypeTestnet:
		want = testnetPubVersion
	default:
		return 0, fmt.Errorf("coin type %d (child number %#x) is not supported", coin, child)
	}
	if v := binary.BigEndian.Uint32(x[0:4]); v != want {
		return 0, fmt.Errorf("version bytes %#08x do not match coin type %d, which needs %#08x", v, coin, want)
	}
	return coin, nil
}

// XpubStrings returns ParentXpubs in their base58check form, in the same
// order — the form pkg/btc/address derives from.
func (e Envelope) XpubStrings() []string {
	out := make([]string, len(e.ParentXpubs))
	for i, x := range e.ParentXpubs {
		out[i] = EncodeXpub(x)
	}
	return out
}

// EncodeXpub is base58check over a raw XpubLen serialization.
func EncodeXpub(raw []byte) string {
	sum := chainhash.DoubleHashB(raw)
	return base58.Encode(append(append([]byte{}, raw...), sum[:4]...))
}

// DecodeXpub parses a base58 extended PUBLIC key into its raw serialization,
// the form ParentXpubs carries. A private key is refused: it would put a
// secret into a hash preimage that is published.
func DecodeXpub(s string) ([]byte, error) {
	k, err := hdkeychain.NewKeyFromString(s)
	if err != nil {
		return nil, fmt.Errorf("parsing extended key: %w", err)
	}
	if k.IsPrivate() {
		return nil, errors.New("an extended PRIVATE key cannot be a signer entry")
	}
	decoded := base58.Decode(s)
	if len(decoded) != XpubLen+4 {
		return nil, fmt.Errorf("extended key decodes to %d bytes, want %d", len(decoded), XpubLen+4)
	}
	return decoded[:XpubLen], nil
}

// KeyParams returns chain parameters whose extended-key version bytes match
// the keys' coin (CoinType). It is for KEY handling — deriving, checking
// version bytes — and not for address encoding: every test network shares coin
// type 1 and tpub, but not the bech32 HRP, so anything that renders an address
// needs its own network, not this.
func (e Envelope) KeyParams() (*chaincfg.Params, error) {
	coin, err := e.CoinType()
	if err != nil {
		return nil, err
	}
	return KeyParamsForCoinType(coin)
}

// KeyParamsForCoinType is KeyParams without an envelope.
func KeyParamsForCoinType(coinType uint32) (*chaincfg.Params, error) {
	switch coinType {
	case CoinTypeBitcoin:
		return &chaincfg.MainNetParams, nil
	case CoinTypeTestnet:
		return &chaincfg.TestNet3Params, nil
	default:
		return nil, fmt.Errorf("coin type %d is not supported", coinType)
	}
}

// CoinTypeFor is the SLIP-44 coin type of a Bitcoin network: 0 on mainnet, 1 on
// every test network.
func CoinTypeFor(params *chaincfg.Params) (uint32, error) {
	switch params.Net {
	case chaincfg.MainNetParams.Net:
		return CoinTypeBitcoin, nil
	case chaincfg.TestNet3Params.Net, chaincfg.SigNetParams.Net, chaincfg.RegressionNetParams.Net:
		return CoinTypeTestnet, nil
	default:
		return 0, fmt.Errorf("network %q has no coin type here", params.Name)
	}
}
