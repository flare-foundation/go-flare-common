//go:generate  abigen --abi=vrf.abi --pkg=vrf --type=TeeVRF --out=autogen.go
package vrf

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/contracts"
	"github.com/flare-foundation/go-flare-common/pkg/tee/op"
)

var opCommands = []op.Command{
	op.VRF,
}

// i-th method correspond to a method in TeeVrfStruct interface whose
// input is the type of message emitted with i-th opCommands.
var methods = []string{
	"vrfInstructionMessageStruct",
}

var MessageArguments map[op.Command]abi.Argument

func init() {
	vrfABI := contracts.MustABI(TeeVRFMetaData)

	if len(methods) != len(opCommands) {
		panic("methods, opCommands miss match")
	}

	MessageArguments = make(map[op.Command]abi.Argument)
	for j := range opCommands {
		MessageArguments[opCommands[j]] = contracts.MustMethod(vrfABI, methods[j]).Inputs[0]
	}
}
