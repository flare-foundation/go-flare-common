//go:generate  abigen --abi=tee.abi --pkg=tee --type=Tee --out=autogen.go
package tee

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
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

	teeAbi, err := TeeMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("error getting tee abi: %v", err))
	}

	for j := range s {
		name := string(s[j]) + "Struct"

		method, ok := teeAbi.Methods[name]
		if !ok {
			panic(fmt.Sprintf("missing method %s", name))
		}

		StructArg[s[j]] = method.Inputs[0]
	}
}
