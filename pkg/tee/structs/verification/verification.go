//go:generate  abigen --abi=verification.abi --pkg=verification --type=Verification --out=autogen.go
package verification

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

type OPCommand string

const (
	TeeAttestation OPCommand = "TEE_ATTESTATION"
)

var opCommands = []OPCommand{
	TeeAttestation,
}

// i-th method correspond to a method in TeePaymentStruct interface whose
// input is the type of message emitted with i-th opCommands
var methods = []string{
	"teeAttestationStruct",
}

var MessageArguments map[OPCommand]abi.Argument

func init() {
	verificationAbi, err := VerificationMetaData.GetAbi()
	if err != nil {
		logger.Panicf("error getting verification abi: %v", err)
	}

	if len(methods) != len(opCommands) {
		logger.Panicf("methods, opCommands miss match")
	}

	MessageArguments = make(map[OPCommand]abi.Argument)
	for j := range opCommands {
		method, ok := verificationAbi.Methods[methods[j]]
		if !ok {
			logger.Panicf("missing method %s", methods[j])
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}
}
