//go:generate  abigen --abi=vrf.abi --pkg=vrf --type=Vrf --out=autogen.go
package vrf

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
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
	vrfAbi, err := VrfMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("error getting vrf abi: %v", err))
	}

	if len(methods) != len(opCommands) {
		panic("methods, opCommands miss match")
	}

	MessageArguments = make(map[op.Command]abi.Argument)
	for j := range opCommands {
		method, ok := vrfAbi.Methods[methods[j]]
		if !ok {
			panic(fmt.Sprintf("missing method %s", methods[j]))
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}
}
