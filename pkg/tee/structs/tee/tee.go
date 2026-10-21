//go:generate  abigen --abi=tee.abi --pkg=tee --type=Tee --out=autogen.go
package tee

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/contracts"
)

type Struct string

const (
	Instruction      Struct = "instruction"
	Attestation      Struct = "attestation"
	VoteSequenceInit Struct = "voteSequenceInit"
	VoteSequenceNext Struct = "voteSequenceNext"
	VoteReceipt      Struct = "voteReceipt"
)

var StructArg map[Struct]abi.Argument

func init() {
	StructArg = make(map[Struct]abi.Argument)

	s := []Struct{
		Instruction,
		Attestation,
		VoteSequenceInit,
		VoteSequenceNext,
		VoteReceipt,
	}

	teeAbi := contracts.MustABI(TeeMetaData)

	for j := range s {
		StructArg[s[j]] = contracts.MustMethod(teeAbi, string(s[j])+"Struct").Inputs[0]
	}
}
