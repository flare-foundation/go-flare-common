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
// serialization. The account is (walletID, accountIndex); recipients live
// inside RawUnsignedTx as raw scriptPubKeys.
//
// # What changed in version 2: the wallet's shape travels with the proposal
//
// The signer set — the n parent xpubs and k — is carried in the envelope rather
// than provisioned to each consumer separately. Before, three parties held
// three copies from three sources: the verifier read the registry, the machines
// held a KEY_BINDING an authorised sender wrote, and the facilitator was
// configured. They agreed only because nothing had made them disagree.
//
// Carrying the keys does not let a proposer choose them. The FDC2 verifier
// requires them to equal the registry's, so a proposal naming any other set
// fails the vote; once it is finalized, the package hash binds exactly the set
// the data providers checked. Inputs remain COORDINATES — every party still
// derives each script itself, from these keys — which is what keeps a proposer
// from supplying a script.
//
// The two fields are APPENDED after escrow; every v1 field keeps its v1
// position. That is not wire compatibility — adding a dynamic field moves every
// offset in the tuple head, so no v1 reader can read a v2 package, which is
// what the version bump says — but it keeps v1's checks applying to v1's
// fields unchanged, and lets a reviewer who knows v1 read two new fields rather
// than a re-laid-out struct.
//
// There is no coin-type field. Each key is the wallet level m/87'/coin', and a
// BIP-32 key records the index of its last derivation step in its own bytes, so
// the coin is read out of the keys (CoinType). A separate field would be a
// second copy of one fact, and a consistency check to keep them equal.
package csp

import (
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
	"github.com/flare-foundation/go-flare-common/pkg/signing"
)

// EnvelopeVersion is the only version this package encodes or accepts.
// Version 2 appended the wallet's shape (parentXpubs, threshold) after escrow;
// a version-1 package no longer decodes, because its hash commits to a tuple
// that no longer exists.
const EnvelopeVersion uint16 = 2

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
// the path that produced it.
type Input struct {
	// Txid is the outpoint's transaction id in INTERNAL byte order — the same
	// order the raw transaction uses, not the reversed display form.
	Txid     [32]byte
	Vout     uint32
	ValueSat uint64
	// Chain is 0 for external, 1 for internal (change), per BIP-32.
	Chain uint8
	Index uint32
	// IsEscrow marks an input whose witness script is the envelope's ESCROW
	// script rather than the plain k-of-n multisig at (Chain, Index).
	//
	// The coordinates still name the KEY each signer uses — the escrow's
	// timeout branch is the wallet's own k-of-n — so this changes which script
	// is committed to, not who may spend. Without it a signer would compute a
	// sighash over the wrong script and produce a signature no one can use.
	IsEscrow bool
}

// Escrow is the HTLC an ESCROW_CREATE pays into or an ESCROW_RECLAIM spends.
//
// Carried in the envelope, so it is covered by the proposal hash and therefore
// by everything the chain finalized. A signer rebuilds the script from these
// fields plus the envelope's signer set — never from a script the proposal
// supplies. The multisig half is therefore the wallet's own: the verifier
// required ParentXpubs to equal the registry's before the hash was finalized,
// and a machine signs only if its own key is among them, so a proposal naming
// keys the wallet does not hold is refused before any signature exists.
type Escrow struct {
	// PreimageHash is SHA256(preimage) for the hash path.
	PreimageHash [32]byte
	// CustodianPubKey is the compressed key that can take the hash path.
	CustodianPubKey []byte
	// Timeout is the UNIX timestamp after which the wallet may reclaim.
	Timeout uint64
	// Chain and Index locate the multisig keys of the TIMEOUT branch.
	Chain uint8
	Index uint32
}

// Envelope is the proposal object.
//
// RawUnsignedTx is the NON-WITNESS serialization, which is also the txid
// preimage — so committing to the envelope pins the txid directly rather than
// by derivation. This is a serialization choice, not a transaction-type choice:
// every transaction in this design is SegWit.
type Envelope struct {
	Version            uint16
	WalletID           [32]byte
	SourceID           [32]byte
	AccountIndex       uint32
	SequencePosition   uint64
	EligibleGeneration uint64
	InstructionType    uint8
	PaymentCount       uint32
	AnchorIndex        uint32
	Nonce              uint64
	Attempt            uint32
	RawUnsignedTx      []byte
	Inputs             []Input
	ProposerAddress    common.Address
	// Escrow is present only for the ESCROW instruction kinds. A zero value
	// means the proposal creates or spends no HTLC, which is every payment and
	// every consolidation.
	Escrow Escrow

	// Version 2 — appended after every v1 field.

	// ParentXpubs are the wallet's n published keys at m/87'/coin', raw
	// XpubLen bytes each, in REGISTRY order. Order does not change any script —
	// BIP-67 sorts the leaf keys — but it is committed, so it has one canonical
	// value. The keys also name the coin (CoinType).
	ParentXpubs [][]byte
	// Threshold is k in the k-of-n.
	Threshold uint32
}

// envelopeArgs is the ABI type tuple. Field order matches the struct, and the
// tuple is fixed: changing it changes every proposalHash ever computed, so it
// is a breaking change to the whole system and not a refactor.
// TestEnvelopeLayout pins it word by word.
var envelopeArgs abi.Arguments

func init() {
	tupleTy, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "version", Type: "uint16"},
		{Name: "walletId", Type: "bytes32"},
		{Name: "sourceId", Type: "bytes32"},
		{Name: "accountIndex", Type: "uint32"},
		{Name: "sequencePosition", Type: "uint64"},
		{Name: "eligibleGeneration", Type: "uint64"},
		{Name: "instructionType", Type: "uint8"},
		{Name: "paymentCount", Type: "uint32"},
		{Name: "anchorIndex", Type: "uint32"},
		{Name: "nonce", Type: "uint64"},
		{Name: "attempt", Type: "uint32"},
		{Name: "rawUnsignedTx", Type: "bytes"},
		{Name: "inputs", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "txid", Type: "bytes32"},
			{Name: "vout", Type: "uint32"},
			{Name: "valueSat", Type: "uint64"},
			{Name: "chain", Type: "uint8"},
			{Name: "index", Type: "uint32"},
			{Name: "isEscrow", Type: "bool"},
		}},
		{Name: "proposerAddress", Type: "address"},
		{Name: "escrow", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "preimageHash", Type: "bytes32"},
			{Name: "custodianPubKey", Type: "bytes"},
			{Name: "timeout", Type: "uint64"},
			{Name: "chain", Type: "uint8"},
			{Name: "index", Type: "uint32"},
		}},
		{Name: "parentXpubs", Type: "bytes[]"},
		{Name: "threshold", Type: "uint32"},
	})
	if err != nil {
		panic("building envelope ABI type: " + err.Error())
	}
	envelopeArgs = abi.Arguments{{Type: tupleTy}}
}

// abiInput mirrors Input with the field tags abi.Pack expects.
type abiInput struct {
	Txid     [32]byte `abi:"txid"`
	Vout     uint32   `abi:"vout"`
	ValueSat uint64   `abi:"valueSat"`
	Chain    uint8    `abi:"chain"`
	Index    uint32   `abi:"index"`
	IsEscrow bool     `abi:"isEscrow"`
}

// abiEscrow mirrors Escrow with the field tags abi.Pack expects.
type abiEscrow struct {
	PreimageHash    [32]byte `abi:"preimageHash"`
	CustodianPubKey []byte   `abi:"custodianPubKey"`
	Timeout         uint64   `abi:"timeout"`
	Chain           uint8    `abi:"chain"`
	Index           uint32   `abi:"index"`
}

type abiEnvelope struct {
	Version            uint16         `abi:"version"`
	WalletID           [32]byte       `abi:"walletId"`
	SourceID           [32]byte       `abi:"sourceId"`
	AccountIndex       uint32         `abi:"accountIndex"`
	SequencePosition   uint64         `abi:"sequencePosition"`
	EligibleGeneration uint64         `abi:"eligibleGeneration"`
	InstructionType    uint8          `abi:"instructionType"`
	PaymentCount       uint32         `abi:"paymentCount"`
	AnchorIndex        uint32         `abi:"anchorIndex"`
	Nonce              uint64         `abi:"nonce"`
	Attempt            uint32         `abi:"attempt"`
	RawUnsignedTx      []byte         `abi:"rawUnsignedTx"`
	Inputs             []abiInput     `abi:"inputs"`
	ProposerAddress    common.Address `abi:"proposerAddress"`
	Escrow             abiEscrow      `abi:"escrow"`
	ParentXpubs        [][]byte       `abi:"parentXpubs"`
	Threshold          uint32         `abi:"threshold"`
}

// Validate rejects envelopes that are structurally unusable. It is called by
// Encode, so no caller can hash something it would refuse.
//
// After the version, v1's checks run unchanged on v1's fields; the signer set
// is checked last, as the part version 2 added.
func (e Envelope) Validate() error {
	if e.Version != EnvelopeVersion {
		return fmt.Errorf("envelope version %d, want %d", e.Version, EnvelopeVersion)
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
		}
	}
	if escrowInputs > 1 {
		return fmt.Errorf("%d escrow inputs; one instruction reclaims one escrow", escrowInputs)
	}
	// Terms without an output to apply them to, or an input with no terms to
	// rebuild its script from, are both unusable rather than merely odd.
	hasTerms := e.Escrow.Timeout != 0 || e.Escrow.PreimageHash != [32]byte{} || len(e.Escrow.CustodianPubKey) > 0
	if escrowInputs == 1 && !hasTerms {
		return errors.New("an escrow input needs the escrow terms to rebuild its witness script")
	}
	if hasTerms {
		if len(e.Escrow.CustodianPubKey) != 33 {
			return fmt.Errorf("escrow custodian key is %d bytes, want a 33-byte compressed key", len(e.Escrow.CustodianPubKey))
		}
		if e.Escrow.Chain > 1 {
			return fmt.Errorf("escrow multisig chain is %d, want 0 or 1", e.Escrow.Chain)
		}
	}
	return e.validateSigners()
}

// Encode returns the canonical bytes.
func (e Envelope) Encode() ([]byte, error) {
	if err := e.Validate(); err != nil {
		return nil, err
	}
	ins := make([]abiInput, len(e.Inputs))
	for i, in := range e.Inputs {
		ins[i] = abiInput(in)
	}
	packed, err := envelopeArgs.Pack(abiEnvelope{
		Version:            e.Version,
		WalletID:           e.WalletID,
		SourceID:           e.SourceID,
		AccountIndex:       e.AccountIndex,
		SequencePosition:   e.SequencePosition,
		EligibleGeneration: e.EligibleGeneration,
		InstructionType:    e.InstructionType,
		PaymentCount:       e.PaymentCount,
		AnchorIndex:        e.AnchorIndex,
		Nonce:              e.Nonce,
		Attempt:            e.Attempt,
		RawUnsignedTx:      e.RawUnsignedTx,
		Inputs:             ins,
		ProposerAddress:    e.ProposerAddress,
		Escrow:             abiEscrow(e.Escrow),
		ParentXpubs:        e.ParentXpubs,
		Threshold:          e.Threshold,
	})
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
// Trailing bytes are REJECTED rather than ignored: if two byte strings could
// decode to one message, the hash no longer identifies the message.
func Decode(b []byte) (Envelope, error) {
	vals, err := envelopeArgs.Unpack(b)
	if err != nil {
		return Envelope{}, fmt.Errorf("decoding envelope: %w", err)
	}
	if len(vals) != 1 {
		return Envelope{}, errors.New("envelope decoded to more than one value")
	}
	// The tuple arrives as an anonymous struct; ConvertType maps it onto the
	// named one field by field.
	raw, ok := abi.ConvertType(vals[0], new(abiEnvelope)).(*abiEnvelope)
	if !ok || raw == nil {
		return Envelope{}, errors.New("envelope did not convert to the expected shape")
	}
	out := Envelope{
		Version:            raw.Version,
		WalletID:           raw.WalletID,
		SourceID:           raw.SourceID,
		AccountIndex:       raw.AccountIndex,
		SequencePosition:   raw.SequencePosition,
		EligibleGeneration: raw.EligibleGeneration,
		InstructionType:    raw.InstructionType,
		PaymentCount:       raw.PaymentCount,
		AnchorIndex:        raw.AnchorIndex,
		Nonce:              raw.Nonce,
		Attempt:            raw.Attempt,
		RawUnsignedTx:      raw.RawUnsignedTx,
		ProposerAddress:    raw.ProposerAddress,
		Escrow:             Escrow(raw.Escrow),
		ParentXpubs:        raw.ParentXpubs,
		Threshold:          raw.Threshold,
	}
	out.Inputs = make([]Input, len(raw.Inputs))
	for i, in := range raw.Inputs {
		out.Inputs[i] = Input(in)
	}

	// Re-encode and compare: the cheapest exact check that the input carried no
	// trailing bytes and no non-canonical padding.
	round, err := out.Encode()
	if err != nil {
		return Envelope{}, err
	}
	if len(round) != len(b) {
		return Envelope{}, fmt.Errorf("envelope is not canonical: %d bytes in, %d re-encoded", len(b), len(round))
	}
	for i := range round {
		if round[i] != b[i] {
			return Envelope{}, errors.New("envelope is not canonical: byte mismatch on re-encode")
		}
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
	seen := make(map[string]struct{}, n)
	for i, x := range e.ParentXpubs {
		if x[45] != 0x02 && x[45] != 0x03 {
			return fmt.Errorf("parentXpubs[%d] does not carry a compressed public key", i)
		}
		if _, dup := seen[string(x)]; dup {
			return fmt.Errorf("parentXpubs[%d] repeats an earlier signer", i)
		}
		seen[string(x)] = struct{}{}
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
