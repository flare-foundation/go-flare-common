package csp

import (
	"bytes"
	"errors"
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/signing"
)

// # The confirmed-attempt proposal
//
// A reissue opens a new attempt of a settled instruction and resets the
// account's eligibility pointer to it; only that instruction's next settlement
// clears the reset. If the ORIGINAL attempt confirms on Bitcoin before the new
// one is attested, the new one respends an anchor outpoint that no longer
// exists and can never be attested, and every anchor chain of the account
// stops. Anyone can cause it: the original is public, and its keyless P2A
// output lets anyone fee-bump it into a block.
//
// A ConfirmedAttempt is the proposal that releases such a position. It
// proposes no transaction. It states that attempt ConfirmedAttempt of the
// instruction at SequencePosition — an attempt that settled before the reissue
// — is on Bitcoin as Txid. The verifiers check that fact in CspProposalCheck,
// the channel admits it only at a pending reset and for a settled attempt
// other than the open one, and its settlement dispatches nothing to sign: the
// reset clears and the pointer moves on.
//
// # Why a message of its own
//
// It is not an Envelope with fields left empty. Every consumer of an Envelope
// — the machines, the facilitator, the fee bumper — reads it as a transaction
// to sign or carry, and an Envelope that meant "nothing to sign" would be one
// field away from being signed. A separate encoding cannot be mistaken for
// one: Decode refuses it, and SplitPackage names it (ErrConfirmedAttempt).
//
// The two encodings are disjoint by their first word. An Envelope is a
// DYNAMIC tuple, so abi.encode opens it with the offset word 0x20; a
// ConfirmedAttempt is a STATIC tuple that opens with ConfirmedAttemptTag, a
// keccak256 output. No byte string is both, so the proposer signs both under
// the one CSPProposal domain (Hash) without a hash of one ever standing for
// the other.

// ConfirmedAttemptTag opens every ConfirmedAttempt encoding.
var ConfirmedAttemptTag = crypto.Keccak256Hash([]byte("CspConfirmedAttempt.v1"))

// ConfirmedAttemptScore is the score a confirmed-attempt proposal is attested
// with: the maximum. The channel ranks it above every ordinary proposal by a
// rule of its own, whatever the scores; the maximum keeps the events and an
// off-chain reader of them in agreement with that rule.
const ConfirmedAttemptScore uint64 = math.MaxUint64

// confirmedAttemptWords is the encoding's length in words: ten static fields.
const confirmedAttemptWords = 10

// ConfirmedAttemptLen is the byte length of an encoded ConfirmedAttempt.
const ConfirmedAttemptLen = confirmedAttemptWords * 32

// ErrConfirmedAttempt is what SplitPackage returns for the package of a
// confirmed-attempt proposal: well-formed, and not a transaction. A consumer
// that only handles transactions treats it as "nothing to sign or carry";
// one that handles both calls SplitConfirmedAttemptPackage.
var ErrConfirmedAttempt = errors.New("the package is a confirmed-attempt proposal, not a transaction envelope")

// ConfirmedAttempt is the confirmed-attempt proposal.
type ConfirmedAttempt struct {
	WalletID           [32]byte
	SourceID           [32]byte
	AccountIndex       uint32
	SequencePosition   uint64
	EligibleGeneration uint64
	// Attempt is the OPEN attempt the eligibility pointer names — the one the
	// reissue opened. The request names it too: the proposal fills that
	// decision point, the way any proposal does.
	Attempt uint32
	// ConfirmedAttempt is the earlier, settled attempt whose transaction is on
	// Bitcoin.
	ConfirmedAttempt uint32
	// Txid is that attempt's transaction id as the channel recorded it at the
	// attempt's settlement, in INTERNAL byte order.
	Txid            [32]byte
	ProposerAddress common.Address
}

var confirmedAttemptArgs = func() abi.Arguments {
	b32, _ := abi.NewType("bytes32", "", nil)
	u32, _ := abi.NewType("uint32", "", nil)
	u64, _ := abi.NewType("uint64", "", nil)
	addr, _ := abi.NewType("address", "", nil)
	return abi.Arguments{
		{Name: "tag", Type: b32},
		{Name: "walletId", Type: b32},
		{Name: "sourceId", Type: b32},
		{Name: "accountIndex", Type: u32},
		{Name: "sequencePosition", Type: u64},
		{Name: "eligibleGeneration", Type: u64},
		{Name: "attempt", Type: u32},
		{Name: "confirmedAttempt", Type: u32},
		{Name: "txid", Type: b32},
		{Name: "proposerAddress", Type: addr},
	}
}()

// Validate rejects a confirmed-attempt proposal that could never be admitted.
// It is called by Encode, so nothing can hash what it would refuse.
//
// The confirmed attempt precedes the open one: attempt numbers only rise, and
// the open attempt — the one a reissue opened last — is the highest. The open
// one is therefore at least 1.
func (c ConfirmedAttempt) Validate() error {
	if c.Attempt == 0 {
		return errors.New("attempt 0 is never open behind a reissue")
	}
	if c.ConfirmedAttempt >= c.Attempt {
		return fmt.Errorf("confirmed attempt %d does not precede the open attempt %d", c.ConfirmedAttempt, c.Attempt)
	}
	if c.Txid == ([32]byte{}) {
		return errors.New("txid is zero")
	}
	if c.ProposerAddress == (common.Address{}) {
		return errors.New("proposer address is zero")
	}
	return nil
}

// Encode returns the canonical bytes: abi.encode of the ten static fields,
// ConfirmedAttemptTag first.
func (c ConfirmedAttempt) Encode() ([]byte, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}
	packed, err := confirmedAttemptArgs.Pack(
		[32]byte(ConfirmedAttemptTag), c.WalletID, c.SourceID, c.AccountIndex, c.SequencePosition,
		c.EligibleGeneration, c.Attempt, c.ConfirmedAttempt, c.Txid, c.ProposerAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("encoding confirmed attempt: %w", err)
	}
	return packed, nil
}

// Hash is what the proposer signs: the CSPProposal domain, as for an Envelope,
// over keccak256 of the encoding. The encodings are disjoint (see the package
// note), so the shared domain cannot make one stand for the other.
func (c ConfirmedAttempt) Hash(chainID uint64) (common.Hash, error) {
	encoded, err := c.Encode()
	if err != nil {
		return common.Hash{}, err
	}
	h, err := signing.NewPayload(signing.CSPProposal, chainID, [32]byte(crypto.Keccak256Hash(encoded))).Hash()
	if err != nil {
		return common.Hash{}, fmt.Errorf("hashing confirmed attempt: %w", err)
	}
	return h, nil
}

// Package is the proposal as the DAL stores it: the encoding followed by the
// proposer's [R || S || V] signature over Hash. Its keccak256 is the
// packageHash the request names.
func (c ConfirmedAttempt) Package(signature []byte) ([]byte, error) {
	if len(signature) != ProposerSigLen {
		return nil, fmt.Errorf("signature must be %d bytes, got %d", ProposerSigLen, len(signature))
	}
	encoded, err := c.Encode()
	if err != nil {
		return nil, err
	}
	return append(encoded, signature...), nil
}

// Commitment is what the proposal commits to on the channel.
func (c ConfirmedAttempt) Commitment() BtcConfirmedAttemptCommitment {
	return BtcConfirmedAttemptCommitment{Attempt: c.ConfirmedAttempt, Txid: c.Txid}
}

// DecodeConfirmedAttempt parses canonical bytes. Like Decode it validates and
// refuses anything that does not re-encode to the same bytes: the input is
// whatever the open DAL served.
func DecodeConfirmedAttempt(b []byte) (ConfirmedAttempt, error) {
	if len(b) != ConfirmedAttemptLen {
		return ConfirmedAttempt{}, fmt.Errorf("a confirmed attempt is %d bytes, got %d", ConfirmedAttemptLen, len(b))
	}
	if !bytes.Equal(b[:32], ConfirmedAttemptTag[:]) {
		return ConfirmedAttempt{}, errors.New("the bytes do not open with the confirmed-attempt tag")
	}
	vals, err := confirmedAttemptArgs.Unpack(b)
	if err != nil {
		return ConfirmedAttempt{}, fmt.Errorf("decoding confirmed attempt: %w", err)
	}
	if len(vals) != confirmedAttemptWords {
		return ConfirmedAttempt{}, fmt.Errorf("confirmed attempt decoded to %d values", len(vals))
	}
	out := ConfirmedAttempt{}
	var ok [9]bool
	out.WalletID, ok[0] = vals[1].([32]byte)
	out.SourceID, ok[1] = vals[2].([32]byte)
	out.AccountIndex, ok[2] = vals[3].(uint32)
	out.SequencePosition, ok[3] = vals[4].(uint64)
	out.EligibleGeneration, ok[4] = vals[5].(uint64)
	out.Attempt, ok[5] = vals[6].(uint32)
	out.ConfirmedAttempt, ok[6] = vals[7].(uint32)
	out.Txid, ok[7] = vals[8].([32]byte)
	out.ProposerAddress, ok[8] = vals[9].(common.Address)
	for i, fine := range ok {
		if !fine {
			return ConfirmedAttempt{}, fmt.Errorf("confirmed attempt field %d has an unexpected type", i+1)
		}
	}
	// Re-encode and compare: rejects non-canonical padding in the narrow
	// integers and the address, which Unpack does not check.
	round, err := out.Encode()
	if err != nil {
		return ConfirmedAttempt{}, err
	}
	if !bytes.Equal(round, b) {
		return ConfirmedAttempt{}, errors.New("confirmed attempt is not canonical: byte mismatch on re-encode")
	}
	return out, nil
}

// IsConfirmedAttemptPackage reports whether stored proposal bytes are a
// SIGNED confirmed-attempt package by shape — its length and its tag. It does
// not validate; SplitConfirmedAttemptPackage does.
func IsConfirmedAttemptPackage(raw []byte) bool {
	return len(raw) == ConfirmedAttemptLen+ProposerSigLen && bytes.Equal(raw[:32], ConfirmedAttemptTag[:])
}

// SplitConfirmedAttemptPackage separates a confirmed-attempt package into the
// proposal and the proposer's signature. Unlike SplitPackage there is no
// unsigned form: a confirmed attempt is only ever read as a proposal under
// vote or one the channel settled, and both are signed.
func SplitConfirmedAttemptPackage(raw []byte) (ConfirmedAttempt, []byte, error) {
	if len(raw) != ConfirmedAttemptLen+ProposerSigLen {
		return ConfirmedAttempt{}, nil, fmt.Errorf("a confirmed-attempt package is %d bytes, got %d",
			ConfirmedAttemptLen+ProposerSigLen, len(raw))
	}
	c, err := DecodeConfirmedAttempt(raw[:ConfirmedAttemptLen])
	if err != nil {
		return ConfirmedAttempt{}, nil, err
	}
	return c, raw[ConfirmedAttemptLen:], nil
}

// Signer recovers who signed a confirmed attempt, and requires it to be the
// proposer the proposal names.
func (c ConfirmedAttempt) Signer(signature []byte, chainID uint64) (common.Address, error) {
	h, err := c.Hash(chainID)
	if err != nil {
		return common.Address{}, err
	}
	signer, err := signatureSigner(h[:], signature)
	if err != nil {
		return common.Address{}, fmt.Errorf("proposer signature does not recover: %w", err)
	}
	if signer != c.ProposerAddress {
		return common.Address{}, fmt.Errorf("confirmed attempt signed by %s but names %s as its proposer",
			signer, c.ProposerAddress)
	}
	return signer, nil
}

// BtcConfirmedAttemptCommitment is what a confirmed-attempt proposal commits to
// on the channel: IBtcAccounts.BtcConfirmedAttemptCommitment, the settled
// attempt and its txid. The proposer passes it in calldata to
// finalizeConfirmedAttempt, the verifier hashes it into the response, and the
// contract binds the two by CommitmentHash. Two static words, where a
// BtcProposalCommitment is five, so no encoding of one is an encoding of the
// other. FIELD ORDER IS THE CONTRACT'S.
type BtcConfirmedAttemptCommitment struct {
	Attempt uint32
	Txid    [32]byte
}

var confirmedCommitmentArgs = func() abi.Arguments {
	u32, _ := abi.NewType("uint32", "", nil)
	b32, _ := abi.NewType("bytes32", "", nil)
	return abi.Arguments{{Type: u32}, {Type: b32}}
}()

// CommitmentHash returns keccak256(abi.encode(commitment)), the value the
// attested response carries as chainCommitmentHash.
func (c BtcConfirmedAttemptCommitment) CommitmentHash() [32]byte {
	packed, err := confirmedCommitmentArgs.Pack(c.Attempt, c.Txid)
	if err != nil {
		// Both fields are static types built above; packing cannot fail.
		panic("csp: packing BtcConfirmedAttemptCommitment: " + err.Error())
	}
	return [32]byte(crypto.Keccak256Hash(packed))
}
