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
package csp

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/signing"
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
// three fields plus its OWN provisioned keys — the multisig half never comes
// from the proposal, which is what keeps the no-signing-oracle property: a
// proposal naming keys the wallet does not hold cannot produce a spendable
// script, because those keys are not the ones used.
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
}

// envelopeArgs is the ABI type tuple. Field order matches the struct, and the
// tuple is fixed: changing it changes every proposalHash ever computed, so it
// is a breaking change to the whole system and not a refactor.
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
}

// Validate rejects envelopes that are structurally unusable. It is called by
// Encode, so no caller can hash something it would refuse.
func (e Envelope) Validate() error {
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
	return nil
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
