// Package instruction provides TEE instruction data types, hashing, and signing utilities.
package instruction

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/tee/op"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/tee"
)

// MaxMessageSize caps each of the variable-length byte fields on an Instruction's Data.
// On-chain events feed Data unchecked; without this cap a pathological message length would
// inflate hash cost and downstream allocation. 64 KiB is an order of magnitude above any
// legitimate payload observed in production.
const MaxMessageSize = 64 * 1024

// MaxCosigners bounds the cosigner set so threshold arithmetic stays sane.
// rippled multi-sig caps at 32 signers; cosigners here mirror that scale.
const MaxCosigners = 32

type Data struct {
	DataFixed
	AdditionalVariableMessage hexutil.Bytes `json:"additionalVariableMessage"`
}

// Validate enforces shape and policy invariants on an instruction parsed
// from an on-chain event. Audit finding M14: untrusted size and field
// content must not flow into hash/dispatch paths unchecked.
func (d *Data) Validate() error {
	if err := d.DataFixed.Validate(); err != nil {
		return err
	}
	if len(d.AdditionalVariableMessage) > MaxMessageSize {
		return fmt.Errorf("AdditionalVariableMessage %d bytes exceeds %d", len(d.AdditionalVariableMessage), MaxMessageSize)
	}
	return nil
}

// Validate enforces shape and policy invariants on the fixed portion of an
// instruction. See Data.Validate.
func (d *DataFixed) Validate() error {
	if len(d.OriginalMessage) > MaxMessageSize {
		return fmt.Errorf("OriginalMessage %d bytes exceeds %d", len(d.OriginalMessage), MaxMessageSize)
	}
	if len(d.AdditionalFixedMessage) > MaxMessageSize {
		return fmt.Errorf("AdditionalFixedMessage %d bytes exceeds %d", len(d.AdditionalFixedMessage), MaxMessageSize)
	}
	if len(d.Cosigners) > MaxCosigners {
		return fmt.Errorf("cosigner count %d exceeds max %d", len(d.Cosigners), MaxCosigners)
	}
	if d.CosignersThreshold == 0 {
		return errors.New("cosigner threshold must be non-zero")
	}
	if d.CosignersThreshold > uint64(len(d.Cosigners)) {
		return fmt.Errorf("cosigner threshold %d exceeds cosigner count %d", d.CosignersThreshold, len(d.Cosigners))
	}
	seen := make(map[common.Address]struct{}, len(d.Cosigners))
	for i, c := range d.Cosigners {
		if _, dup := seen[c]; dup {
			return fmt.Errorf("duplicate cosigner at index %d: %s", i, c.Hex())
		}
		seen[c] = struct{}{}
	}
	if !op.IsValidPair(d.OPType, d.OPCommand) {
		return fmt.Errorf("invalid op type/command pair: type=%s command=%s", d.OPType.Hex(), d.OPCommand.Hex())
	}
	return nil
}

type DataFixed struct {
	InstructionID          common.Hash      `json:"instructionId"`
	TeeID                  common.Address   `json:"teeId"`
	Timestamp              uint64           `json:"timestamp"`
	RewardEpochID          uint32           `json:"rewardEpochId"`
	OPType                 common.Hash      `json:"opType"`
	OPCommand              common.Hash      `json:"opCommand"`
	Cosigners              []common.Address `json:"cosigners"`
	CosignersThreshold     uint64           `json:"cosignersThreshold"`
	OriginalMessage        hexutil.Bytes    `json:"originalMessage"`
	AdditionalFixedMessage hexutil.Bytes    `json:"additionalFixedMessage"`
}

// HashFixed computes the hash of the DataFixed.
func (d *DataFixed) HashFixed() (common.Hash, error) {
	e, err := structs.Encode(tee.StructArg[tee.Instruction], d.prepareForEncoding())
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(e), nil
}

func (d *DataFixed) prepareForEncoding() tee.TeeStructsInstruction {
	return tee.TeeStructsInstruction{
		InstructionId:          d.InstructionID,
		TeeId:                  d.TeeID,
		Timestamp:              d.Timestamp,
		RewardEpochId:          d.RewardEpochID,
		OpType:                 d.OPType,
		OpCommand:              d.OPCommand,
		Cosigners:              d.Cosigners,
		CosignersThreshold:     d.CosignersThreshold,
		OriginalMessage:        d.OriginalMessage,
		AdditionalFixedMessage: d.AdditionalFixedMessage,
	}
}

// InitialVoteHash computes the initial vote hash for the voting on instruction data.
func (d *DataFixed) InitialVoteHash() (common.Hash, error) {
	ih, err := d.HashFixed()
	if err != nil {
		return common.Hash{}, err
	}

	s := tee.TeeStructsVoteSequenceInit{
		InstructionId:   d.InstructionID,
		InstructionHash: ih,
		RewardEpochId:   d.RewardEpochID,
		TeeId:           d.TeeID,
	}

	e, err := structs.Encode(tee.StructArg[tee.VoteSequenceInit], s)
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(e), nil
}

func NextVoteHash(hash common.Hash, sequence uint64, signature, additionalVariableMessage []byte, time uint64) (common.Hash, error) {
	avmh := crypto.Keccak256Hash(additionalVariableMessage)

	s := tee.TeeStructsVoteSequenceNext{
		VoteHash:                      hash,
		Sequence:                      sequence,
		Signature:                     signature,
		AdditionalVariableMessageHash: avmh,
		Timestamp:                     time,
	}

	e, err := structs.Encode(tee.StructArg[tee.VoteSequenceNext], s)
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(e), nil
}

// HashForSigning computes the hash of the Data d that is signed by the provider.
func (d Data) HashForSigning() (common.Hash, error) {
	fixed, err := d.HashFixed()
	if err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash(fixed[:], crypto.Keccak256(d.AdditionalVariableMessage)), nil
}

// SignInstructionHash signs the hash of the tee instruction.
//
// FOR REFERENCE!
func SignInstructionHash(hash common.Hash, pk *ecdsa.PrivateKey) ([]byte, error) {
	hashToSign := accounts.TextHash(hash[:])
	return crypto.Sign(hashToSign, pk)
}

type Instruction struct {
	Data      Data          `json:"data"`
	Signature hexutil.Bytes `json:"signature"`
}

// RecoverSignersPubKey recovers the signers public key from Data and Signature.
// Non-canonical (high-S) signatures are rejected to keep the (data, signature) pair unique.
func (i Instruction) RecoverSignersPubKey() (*ecdsa.PublicKey, error) {
	hash, err := i.Data.HashForSigning()
	if err != nil {
		return nil, err
	}
	signedHash := accounts.TextHash(hash[:])

	if len(i.Signature) != 65 {
		return nil, fmt.Errorf("signature must be 65 bytes, got %d", len(i.Signature))
	}
	r := new(big.Int).SetBytes(i.Signature[:32])
	s := new(big.Int).SetBytes(i.Signature[32:64])
	v := i.Signature[64]
	if !crypto.ValidateSignatureValues(v, r, s, true) {
		return nil, errors.New("non-canonical signature (high-S or zero scalar)")
	}

	return crypto.SigToPub(signedHash, i.Signature)
}
