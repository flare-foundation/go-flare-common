package instruction

import (
	"crypto/ecdsa"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Data struct {
	DataFixed
	AdditionalVariableMessage hexutil.Bytes `json:"additionalVariableMessage"`
}

type DataFixed struct {
	InstructionID          common.Hash    `json:"instructionId"`
	TeeID                  common.Address `json:"teeId"`
	Timestamp              uint32         `json:"timestamp"`
	RewardEpochID          *big.Int       `json:"rewardEpochId"`
	OPType                 common.Hash    `json:"opType"`
	OPCommand              common.Hash    `json:"opCommand"`
	OriginalMessage        hexutil.Bytes  `json:"originalMessage"`
	AdditionalFixedMessage hexutil.Bytes  `json:"additionalFixedMessage"`
}

// Hash computes the hash of the DataFixed d.
func (d DataFixed) HashFixed() (common.Hash, error) {
	m, err := json.Marshal(d)
	if err != nil {
		return common.Hash{}, err
	}

	return crypto.Keccak256Hash(m), nil
}

// HashForSigning computes the hash of the Data d that is signed by the provider.
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
	Challenge common.Hash   `json:"challenge"`
	Data      Data          `json:"data"`
	Signature hexutil.Bytes `json:"signature"`
}

// RecoverSignersPubKey recovers the signers public key from Data and Signature.
func (i Instruction) RecoverSignersPubKey() ([]byte, error) {
	hash, err := i.Data.HashForSigning()
	if err != nil {
		return []byte{}, nil
	}
	signedHash := accounts.TextHash(hash[:])

	return crypto.Ecrecover(signedHash, i.Signature)
}
