package instruction

import (
	"crypto/ecdsa"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Data struct {
	InstructionID             common.Hash    `json:"instructionId"`
	TeeID                     common.Address `json:"teeId"`
	RewardEpochID             *big.Int       `json:"rewardEpochId"`
	OPType                    common.Hash    `json:"opType"`
	OPCommand                 common.Hash    `json:"opCommand"`
	OriginalMessage           []byte         `json:"originalMessage"`
	AdditionalFixedMessage    []byte         `json:"additionalFixedMessage"`
	AdditionalVariableMessage []byte         `json:"additionalVariableMessage"`
}

func (d Data) HashForSigning() (common.Hash, error) {
	m, err := json.Marshal(d)
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(m), nil
}

// SignInstructionHash signs the hash of the tee instruction
//
// FOR REFERENCE
func SignInstructionHash(hash common.Hash, pk *ecdsa.PrivateKey) ([]byte, error) {
	hashToSign := accounts.TextHash(hash[:])
	return crypto.Sign(hashToSign, pk)
}

type Instruction struct {
	Challenge any    `json:"challenge"`
	Data      Data   `json:"data"`
	Signature []byte `json:"signature"`
}

func (i Instruction) RecoverSignersPubKey() ([]byte, error) {
	hash, err := i.Data.HashForSigning()
	if err != nil {
		return []byte{}, nil
	}
	signedHash := accounts.TextHash(hash[:])

	return crypto.Ecrecover(signedHash, i.Signature)
}
