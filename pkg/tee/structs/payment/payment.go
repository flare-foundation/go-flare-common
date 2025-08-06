//go:generate  abigen --abi=payment.abi --pkg=payment --type=Payment --out=autogen.go
package payment

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/flare-foundation/go-flare-common/pkg/tee/constants"
)

var opCommands = []constants.OPCommand{
	constants.Pay,
	constants.Reissue,
}

// i-th method correspond to a method in TeePaymentStruct interface whose
// input is the type of message emitted with i-th opCommands.
var methods = []string{
	"paymentInstructionMessageStruct",
	"paymentInstructionMessageStruct",
}

var MessageArguments map[constants.OPCommand]abi.Argument

func init() {
	paymentAbi, err := PaymentMetaData.GetAbi()
	if err != nil {
		logger.Panicf("error getting payment abi: %v", err)
	}

	if len(methods) != len(opCommands) {
		logger.Panicf("methods, opCommands miss match")
	}

	MessageArguments = make(map[constants.OPCommand]abi.Argument)
	for j := range opCommands {
		method, ok := paymentAbi.Methods[methods[j]]
		if !ok {
			logger.Panicf("missing method %s", methods[j])
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}
}
