package instruction

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs"
	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/tee"
)

type Data struct {
	DataFixed
	AdditionalVariableMessage hexutil.Bytes `json:"AdditionalVariableMessage"`
}

type DataFixed tee.TeeStructsInstruction

// Hash computes the hash of the DataFixed d.
func (d DataFixed) HashFixed() (common.Hash, error) {
	e, err := structs.Encode(tee.StructArg[tee.Instruction], d)
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(e), nil
}

func (d *DataFixed) InitialVoteHash() (common.Hash, error) {
	ih, err := d.HashFixed()
	if err != nil {
		return common.Hash{}, err
	}

	s := tee.TeeStructsVoteSequenceInit{
		InstructionId:   d.InstructionId,
		InstructionHash: ih,
		RewardEpochId:   d.RewardEpochId,
		TeeId:           d.TeeId,
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
func (i Instruction) RecoverSignersPubKey() ([]byte, error) {
	hash, err := i.Data.HashForSigning()
	if err != nil {
		return []byte{}, err
	}
	signedHash := accounts.TextHash(hash[:])

	return crypto.Ecrecover(signedHash, i.Signature)
}
