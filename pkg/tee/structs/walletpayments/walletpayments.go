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
}

// The i-th method is the one in the payments-struct interface whose sole input
// is the message the account-based payment model emits with the i-th
// opCommand: a payment and its reissue carry a PaymentInstructionMessage, an
// XRPL escrow an XrpEscrowCreateMessage.
var methods = []string{
	"paymentInstructionMessageStruct",
	"paymentInstructionMessageStruct",
	"xrpEscrowCreateMessageStruct",
}

// MessageArguments are the messages of the account-based payment model, by the
// command they are emitted with.
var MessageArguments map[op.Command]abi.Argument

// CspMessageArgument is the CspInstructionMessage that CSP settlement emits
// with every command it dispatches. It is not in MessageArguments because the
// account-based model emits different messages with the same commands.
var CspMessageArgument abi.Argument

// BtcProposalCommitmentArgument is the BtcProposalCommitment a Bitcoin CSP
// proposal commits to, whose keccak256(abi.encode(...)) is the proposal's
// chainCommitmentHash.
var BtcProposalCommitmentArgument abi.Argument

// BtcConfirmedAttemptCommitmentArgument is the BtcConfirmedAttemptCommitment a
// confirmed-attempt proposal commits to, whose keccak256(abi.encode(...)) is
// the proposal's chainCommitmentHash.
var BtcConfirmedAttemptCommitmentArgument abi.Argument

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
		MessageArguments[opCommands[j]] = structArgument(paymentsABI, methods[j])
	}

	CspMessageArgument = structArgument(paymentsABI, "cspInstructionMessageStruct")
	BtcProposalCommitmentArgument = structArgument(paymentsABI, "btcProposalCommitmentStruct")
	BtcConfirmedAttemptCommitmentArgument = structArgument(paymentsABI, "btcConfirmedAttemptCommitmentStruct")
}

// structArgument returns the sole input of the struct ABI's method name.
func structArgument(a *abi.ABI, name string) abi.Argument {
	method, ok := a.Methods[name]
	if !ok {
		panic(fmt.Sprintf("missing method %s", name))
	}
	return method.Inputs[0]
}
