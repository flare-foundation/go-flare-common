package csp

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Instruction is the payload UtxoInstructionChannel dispatches through the
// FlareTeeManager diamond when a leader settles, carried in the TEE
// instruction's OriginalMessage.
//
// It describes WHAT WAS DECIDED and not what is to be signed: the Bitcoin
// transaction lives in the DAL under Txid, and ProposalHash is what binds those
// bytes to this decision. Splitting it that way keeps an unbounded transaction
// off chain while still making the binding checkable by anyone holding both.
type Instruction struct {
	// TeeIDKeyIDPairs says which machine holds which key for this wallet. A
	// machine filters it for its own id to learn which key to sign with — the
	// same shape TeePayments carries, so the machine-side lookup is unchanged.
	TeeIDKeyIDPairs  []TeeIDKeyIDPair
	WalletID         [32]byte
	SourceID         [32]byte
	AccountIndex     uint32
	SequencePosition uint64
	Attempt          uint32
	AnchorIndex      uint32
	Nonce            uint64
	FromPaymentID    uint64
	ToPaymentID      uint64
	PackageHash      common.Hash
	Txid             common.Hash
	Proposer         common.Address
}

// TeeIDKeyIDPair binds a machine to one of its keys.
type TeeIDKeyIDPair struct {
	TeeID common.Address
	KeyID uint64
}

var instructionArgs abi.Arguments

func init() {
	ty, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "teeIdKeyIdPairs", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "teeId", Type: "address"},
			{Name: "keyId", Type: "uint64"},
		}},
		{Name: "walletId", Type: "bytes32"},
		{Name: "sourceId", Type: "bytes32"},
		{Name: "accountIndex", Type: "uint32"},
		{Name: "sequencePosition", Type: "uint64"},
		{Name: "attempt", Type: "uint32"},
		{Name: "anchorIndex", Type: "uint32"},
		{Name: "nonce", Type: "uint64"},
		{Name: "fromPaymentId", Type: "uint64"},
		{Name: "toPaymentId", Type: "uint64"},
		{Name: "packageHash", Type: "bytes32"},
		{Name: "txid", Type: "bytes32"},
		{Name: "proposer", Type: "address"},
	})
	if err != nil {
		panic("building instruction ABI type: " + err.Error())
	}
	instructionArgs = abi.Arguments{{Type: ty}}
}

type abiPair struct {
	TeeID common.Address `abi:"teeId"`
	KeyID uint64         `abi:"keyId"`
}

type abiInstruction struct {
	TeeIDKeyIDPairs  []abiPair      `abi:"teeIdKeyIdPairs"`
	WalletID         [32]byte       `abi:"walletId"`
	SourceID         [32]byte       `abi:"sourceId"`
	AccountIndex     uint32         `abi:"accountIndex"`
	SequencePosition uint64         `abi:"sequencePosition"`
	Attempt          uint32         `abi:"attempt"`
	AnchorIndex      uint32         `abi:"anchorIndex"`
	Nonce            uint64         `abi:"nonce"`
	FromPaymentID    uint64         `abi:"fromPaymentId"`
	ToPaymentID      uint64         `abi:"toPaymentId"`
	PackageHash      [32]byte       `abi:"packageHash"`
	Txid             [32]byte       `abi:"txid"`
	Proposer         common.Address `abi:"proposer"`
}

// EncodeInstruction is the mirror of DecodeInstruction, used by tests and by
// anything simulating the channel.
func EncodeInstruction(i Instruction) ([]byte, error) {
	pairs := make([]abiPair, len(i.TeeIDKeyIDPairs))
	for n, p := range i.TeeIDKeyIDPairs {
		pairs[n] = abiPair(p)
	}
	b, err := instructionArgs.Pack(abiInstruction{
		TeeIDKeyIDPairs:  pairs,
		WalletID:         i.WalletID,
		SourceID:         i.SourceID,
		AccountIndex:     i.AccountIndex,
		SequencePosition: i.SequencePosition,
		Attempt:          i.Attempt,
		AnchorIndex:      i.AnchorIndex,
		Nonce:            i.Nonce,
		FromPaymentID:    i.FromPaymentID,
		ToPaymentID:      i.ToPaymentID,
		PackageHash:      i.PackageHash,
		Txid:             i.Txid,
		Proposer:         i.Proposer,
	})
	if err != nil {
		return nil, fmt.Errorf("encoding instruction: %w", err)
	}
	return b, nil
}

// DecodeInstruction parses the dispatched payload.
func DecodeInstruction(b []byte) (Instruction, error) {
	vals, err := instructionArgs.Unpack(b)
	if err != nil {
		return Instruction{}, fmt.Errorf("decoding instruction: %w", err)
	}
	if len(vals) != 1 {
		return Instruction{}, errors.New("instruction decoded to more than one value")
	}
	raw, ok := abi.ConvertType(vals[0], new(abiInstruction)).(*abiInstruction)
	if !ok || raw == nil {
		return Instruction{}, errors.New("instruction did not convert to the expected shape")
	}
	out := Instruction{
		WalletID:         raw.WalletID,
		SourceID:         raw.SourceID,
		AccountIndex:     raw.AccountIndex,
		SequencePosition: raw.SequencePosition,
		Attempt:          raw.Attempt,
		AnchorIndex:      raw.AnchorIndex,
		Nonce:            raw.Nonce,
		FromPaymentID:    raw.FromPaymentID,
		ToPaymentID:      raw.ToPaymentID,
		PackageHash:      raw.PackageHash,
		Txid:             raw.Txid,
		Proposer:         raw.Proposer,
	}
	out.TeeIDKeyIDPairs = make([]TeeIDKeyIDPair, len(raw.TeeIDKeyIDPairs))
	for n, p := range raw.TeeIDKeyIDPairs {
		out.TeeIDKeyIDPairs[n] = TeeIDKeyIDPair(p)
	}
	return out, nil
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
func SplitPackage(raw []byte) (Envelope, []byte, error) {
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
	if crypto.Keccak256Hash(raw) != i.PackageHash {
		return Envelope{}, fmt.Errorf("package hash %s does not match the finalized package hash %s",
			crypto.Keccak256Hash(raw), i.PackageHash)
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
	if e.WalletID != i.WalletID || e.AccountIndex != i.AccountIndex ||
		e.SequencePosition != i.SequencePosition || e.Attempt != i.Attempt {
		return errors.New("envelope identity does not match the instruction")
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
	for _, p := range i.TeeIDKeyIDPairs {
		if p.TeeID == teeID {
			out = append(out, p.KeyID)
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
