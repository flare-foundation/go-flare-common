//go:generate  abigen --abi=payments.abi --pkg=payments --type=TeePayments --out=autogen.go
package payments

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/contracts"
	"github.com/flare-foundation/go-flare-common/pkg/tee/op"
)

var opCommands = []op.Command{
	op.Pay,
	op.Reissue,
}

// i-th method correspond to a method in TeePaymentStruct interface whose
// input is the type of message emitted with i-th opCommands.
var methods = []string{
	"paymentInstructionMessageStruct",
	"paymentInstructionMessageStruct",
}

var MessageArguments map[op.Command]abi.Argument

func init() {
	paymentsABI := contracts.MustABI(TeePaymentsMetaData)

	if len(methods) != len(opCommands) {
		panic("methods, opCommands miss match")
	}

	MessageArguments = make(map[op.Command]abi.Argument)
	for j := range opCommands {
		MessageArguments[opCommands[j]] = contracts.MustMethod(paymentsABI, methods[j]).Inputs[0]
	}
}
