//go:generate  abigen --abi=verification.abi --pkg=verification --type=TeeVerification --out=autogen.go
package verification

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/contracts"
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
	verificationABI := contracts.MustABI(TeeVerificationMetaData)

	if len(methods) != len(opCommands) {
		panic("methods, opCommands miss match")
	}

	MessageArguments = make(map[op.Command]abi.Argument)
	for j := range opCommands {
		MessageArguments[opCommands[j]] = contracts.MustMethod(verificationABI, methods[j]).Inputs[0]
	}
}
