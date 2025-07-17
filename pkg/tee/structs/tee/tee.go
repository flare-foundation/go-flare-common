//go:generate  abigen --abi=tee.abi --pkg=tee --type=Tee --out=autogen.go
package tee

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

type Struct string

const (
	Instruction      Struct = "teeInstruction"
	Attestation      Struct = "attestation"
	PMWState         Struct = "pmwState"
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
		PMWState,
		VoteSequenceInit,
		VoteSequenceNext,
		VoteReceipt,
	}

	teeAbi, err := TeeMetaData.GetAbi()
	if err != nil {
		logger.Panicf("error getting tee abi: %v", err)
	}

	for j := range s {
		name := string(s[j]) + "Struct"

		method, ok := teeAbi.Methods[name]
		if !ok {
			logger.Panicf("missing method %s", name)
		}

		StructArg[s[j]] = method.Inputs[0]
	}
}
