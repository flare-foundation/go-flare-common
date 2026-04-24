//go:generate  abigen --abi=verification.abi --pkg=verification --type=TeeVerification --out=autogen.go
package verification

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/tee/op"
)

var opCommands = []op.Command{
	op.TEEAttestation,
}

// i-th method correspond to a method in TeePaymentStruct interface whose
// input is the type of message emitted with i-th opCommands.
var methods = []string{
	"teeAttestationStruct",
}

var MessageArguments map[op.Command]abi.Argument

func init() {
	verificationABI, err := TeeVerificationMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("error getting verification abi: %v", err))
	}

	if len(methods) != len(opCommands) {
		panic("methods, opCommands miss match")
	}

	MessageArguments = make(map[op.Command]abi.Argument)
	for j := range opCommands {
		method, ok := verificationABI.Methods[methods[j]]
		if !ok {
			panic(fmt.Sprintf("missing method %s", methods[j]))
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}
}
