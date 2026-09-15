//go:generate  abigen --abi=walletpayments.abi --pkg=walletpayments --type=WalletPayments --out=autogen.go
package walletpayments

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

// The i-th method is the one in the payments-struct interface whose sole input
// is the message type emitted with the i-th opCommand.
//
// The CSP control plane emits ONE message shape for every instruction kind —
// settle encodes a `CspPaymentInstructionMessage` whatever the batch does — so
// escrow and consolidation share the payment message's arguments rather than
// having any of their own. This map is the ACCOUNT model's (`IPayments`); the
// CSP tuple is resolved separately where it is decoded, because the two planes
// no longer share a struct.
var methods = []string{
	"paymentInstructionMessageStruct",
	"paymentInstructionMessageStruct",
	"paymentInstructionMessageStruct",
	"paymentInstructionMessageStruct",
	"paymentInstructionMessageStruct",
}

var MessageArguments map[op.Command]abi.Argument

func init() {
	paymentsABI, err := WalletPaymentsMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("error getting the WalletPayments abi: %v", err))
	}

	if len(methods) != len(opCommands) {
		panic("methods and opCommands disagree in length")
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
