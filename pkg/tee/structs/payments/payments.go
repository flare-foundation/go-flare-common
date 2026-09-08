//go:generate  abigen --abi=payments.abi --pkg=payments --type=TeePayments --out=autogen.go
package payments

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/tee/op"
)

var opCommands = []op.Command{
	op.Pay,
	op.Reissue,
	op.EscrowCreate,
	op.EscrowReclaim,
	op.Consolidate,
}

// i-th method correspond to a method in TeePaymentStruct interface whose
// input is the type of message emitted with i-th opCommands.
//
// The UTXO channel emits ONE message shape for every instruction kind — settle
// encodes a `UtxoCspInstructionMessage` whatever the batch does — so the escrow
// and consolidation commands share the payment message's arguments rather than
// having any of their own.
var methods = []string{
	"paymentInstructionMessageStruct",
	"paymentInstructionMessageStruct",
	"paymentInstructionMessageStruct",
	"paymentInstructionMessageStruct",
	"paymentInstructionMessageStruct",
}

var MessageArguments map[op.Command]abi.Argument

func init() {
	paymentsABI, err := TeePaymentsMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("error getting payment abi: %v", err))
	}

	if len(methods) != len(opCommands) {
		panic("methods, opCommands miss match")
	}

	MessageArguments = make(map[op.Command]abi.Argument)
	for j := range opCommands {
		method, ok := paymentsABI.Methods[methods[j]]
		if !ok {
			panic(fmt.Sprintf("missing method %s", methods[j]))
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}
}
