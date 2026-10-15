package csp

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/flare-foundation/go-flare-common/pkg/abicoder"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/walletpayments"
)

// Instruction is the CspInstructionMessage that CSP settlement dispatches
// through the FlareTeeManager diamond when a leader settles, carried in the TEE
// instruction's OriginalMessage. Its fields are the contract's, as generated
// from the struct ABI.
//
// It describes WHAT WAS DECIDED and not what is to be signed. The proposal
// package — the envelope and the proposer's signature — lives in the DAL under
// PackageHash, which binds those bytes to this decision (BindPackage), so an
// unbounded transaction stays off chain while the binding remains checkable by
// anyone holding both. ChainCommitmentHash commits to the winning proposal's
// chain-specific facts without naming them; a signer reads those facts from
// the envelope it signs. TeeIdKeyIdPairs says which machine holds which of the
// wallet's keys (KeysFor).
type Instruction walletpayments.ICspSettlementCspInstructionMessage

// EncodeInstruction ABI-encodes the instruction as the contract dispatches it,
// for tests and for anything simulating the channel.
func EncodeInstruction(i Instruction) ([]byte, error) {
	b, err := abicoder.Encode(walletpayments.CspMessageArgument, walletpayments.ICspSettlementCspInstructionMessage(i))
	if err != nil {
		return nil, fmt.Errorf("encoding instruction: %w", err)
	}
	return b, nil
}

// DecodeInstruction parses the dispatched payload.
func DecodeInstruction(b []byte) (Instruction, error) {
	m, err := abicoder.Decode[walletpayments.ICspSettlementCspInstructionMessage](walletpayments.CspMessageArgument, b)
	if err != nil {
		return Instruction{}, fmt.Errorf("decoding instruction: %w", err)
	}
	return Instruction(m), nil
}

// ProposerSigLen is [R || S || V]: fixed width, so a package splits without a
// length prefix.
const ProposerSigLen = 65

// SplitPackage separates stored proposal bytes into the envelope and the
// proposer's signature over it.
//
// The bare-envelope form is still decoded, without a signature, because some
// lookups are hints rather than authorities. Callers that must not accept an
// unsigned proposal check for a nil signature themselves.
//
// The package of a confirmed-attempt proposal is refused with
// ErrConfirmedAttempt: it is well-formed and carries no transaction, which a
// caller that signs, carries or accounts for transactions must not mistake for
// a malformed package (see SplitConfirmedAttemptPackage).
func SplitPackage(raw []byte) (Envelope, []byte, error) {
	if IsConfirmedAttemptPackage(raw) {
		return Envelope{}, nil, ErrConfirmedAttempt
	}
	if len(raw) > ProposerSigLen {
		if env, err := Decode(raw[:len(raw)-ProposerSigLen]); err == nil {
			return env, raw[len(raw)-ProposerSigLen:], nil
		}
	}
	env, err := Decode(raw)
	return env, nil, err
}

// BindPackage checks that proposal bytes fetched from the DAL are the ones this
// instruction decided on.
//
// This is the ONE check that makes untrusted content-addressed storage safe to
// read from: without it a machine would sign whatever bytes it was handed.
// Both the relay client and the machine run it — the relay client so it never
// signs an unbound proposal, the machine because it holds both halves and the
// check is free.
//
// It binds the PACKAGE, not the envelope. What the proposer committed to on
// chain, and what the channel finalized, is keccak(envelope ‖ proposerSig), so
// binding the envelope alone would leave the signature unchecked — and the
// signature is what says the proposal is the work of the party being paid for
// it.
func (i Instruction) BindPackage(raw []byte, chainID uint64) (Envelope, error) {
	if h := crypto.Keccak256Hash(raw); h != common.Hash(i.PackageHash) {
		return Envelope{}, fmt.Errorf("package hash %s does not match the finalized package hash %s",
			h, common.Hash(i.PackageHash))
	}
	e, sig, err := SplitPackage(raw)
	if err != nil {
		return Envelope{}, fmt.Errorf("decoding the proposal package: %w", err)
	}
	if sig == nil {
		return Envelope{}, errors.New("the stored proposal carries no proposer signature")
	}
	if err := i.bindIdentity(e, sig, chainID); err != nil {
		return Envelope{}, err
	}
	return e, nil
}

func (i Instruction) bindIdentity(e Envelope, sig []byte, chainID uint64) error {
	h, err := e.Hash(chainID)
	if err != nil {
		return err
	}
	signer, err := signatureSigner(h[:], sig)
	if err != nil {
		return fmt.Errorf("proposer signature does not recover: %w", err)
	}
	if signer != e.ProposerAddress {
		return fmt.Errorf("package signed by %s but names %s as its proposer", signer, e.ProposerAddress)
	}
	// The identity fields must agree too. A hash match already implies it, but
	// a mismatch here means the two were built from different intentions and is
	// worth reporting as itself rather than as an opaque hash difference.
	if e.WalletId != i.WalletId || e.AccountIndex != i.AccountIndex ||
		e.SequencePosition != i.SequencePosition || e.Attempt != i.Attempt {
		return errors.New("envelope identity does not match the instruction")
	}
	// The SOURCE is part of the account's identity — the chain keys an account
	// by (sourceId, account) — and it is the registry the verifier checked the
	// envelope's keys against. The instruction's comes from the chain, so it is
	// the one that wins. (It does not choose the network: the keys name their
	// own coin, Envelope.CoinType.)
	if e.SourceId != i.SourceId {
		return errors.New("envelope names a different source than the instruction")
	}
	return nil
}

// KeysFor returns the key ids this machine holds for the wallet.
//
// A machine that appears in no pair is not a signer for this wallet, which is a
// routing mistake rather than a failure: it should decline rather than search
// its storage for a key it was never asked to use.
func (i Instruction) KeysFor(teeID common.Address) []uint64 {
	out := make([]uint64, 0, 1)
	for _, p := range i.TeeIdKeyIdPairs {
		if p.TeeId == teeID {
			out = append(out, p.KeyId)
		}
	}
	return out
}

// signatureSigner recovers the EIP-191 signer of a hash.
//
// Kept here rather than imported from the TEE packages so that this library,
// which both the relay client and the machines depend on, does not acquire a
// dependency on either.
func signatureSigner(hash, signature []byte) (common.Address, error) {
	if len(signature) != ProposerSigLen {
		return common.Address{}, fmt.Errorf("signature must be %d bytes, got %d", ProposerSigLen, len(signature))
	}
	pub, err := crypto.SigToPub(accounts.TextHash(hash), signature)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*pub), nil
}
