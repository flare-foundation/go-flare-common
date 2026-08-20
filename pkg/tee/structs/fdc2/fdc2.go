//go:generate  abigen --abi=fdc2.abi --pkg=fdc2 --type=Fdc2 --out=autogen.go
package fdc2

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/tee/op"
)

const (
	reqStruct   string = "RequestBodyStruct"
	resStruct   string = "ResponseBodyStruct"
	proofStruct string = "ProofStruct"
)

var opCommands = []op.Command{
	op.Prove,
}

// i-th method correspond to a method in TeeDataConnectorStruct interface whose
// input is the type of message emitted with i-th element of opCommands.
var methods = []string{
	"fdc2AttestationRequestStruct",
}

var MessageArguments map[op.Command]abi.Argument

type AttestationType string

const (
	AvailabilityCheck            AttestationType = "TeeAvailabilityCheck"
	PMWPaymentStatus             AttestationType = "PMWPaymentStatus"
	PMWMultisigAccountConfigured AttestationType = "PMWMultisigAccountConfigured"
	PMWFeeProof                  AttestationType = "PMWFeeProof"
	PMWMultisigUtxoConfigured    AttestationType = "PMWMultisigUtxoConfigured"
	// PMWUtxoProposalCheck is CSP's proposal predicate: the data providers
	// decide, before any signature exists, which of several competing
	// transactions fills a batch.
	PMWUtxoProposalCheck AttestationType = "PMWUtxoProposalCheck"
	// BtcDeposit proves that one OUTPUT of one confirmed Bitcoin transaction paid
	// the address a wallet derives at a reserved index. It is the deposit
	// counterpart of PMWPaymentStatus: that type proves a payment a
	// protocol-managed wallet MADE, this one proves a payment it RECEIVED, and
	// the receiving address is the identity — Bitcoin has no destination tag, so
	// the index the address derives from is what says who to credit.
	BtcDeposit AttestationType = "BtcDeposit"
	// BtcPayment proves that one named Bitcoin address paid another in one
	// confirmed transaction: the amount received, what left the payer, the
	// transaction's OP_RETURN reference, and the block it sits in.
	//
	// WHY IT IS NOT PMWPaymentStatus, WHICH LOOKS LIKE THE SAME FACT. That type
	// proves a payment a PROTOCOL-MANAGED WALLET made, and it proves it by
	// resolving the payment to its batch, recomputing the batch instruction id,
	// and matching the output group against the instruction the channel emitted.
	// Every step of that needs the payer to be a channel account. A Core Vault
	// is not one and cannot become one — a channel account's keys are generated
	// inside TEE machines (WalletKeyManagerFacet.addKey requires a machine in
	// PRODUCTION) and a vault's keys are held by human custodians, which is the
	// separate custody a vault exists to provide.
	//
	// So this type proves the weaker, chain-only fact, and the machines attest
	// it as OBSERVERS rather than as custodians: nothing about proving a Bitcoin
	// payment requires controlling the wallet that made it. The payer is named
	// in the REQUEST and checked against the inputs' prevouts, because a
	// transaction has many inputs and "who sent it" is otherwise not a question
	// Bitcoin answers.
	BtcPayment AttestationType = "BtcPayment"
)

var attestationTypes = []AttestationType{
	AvailabilityCheck,
	PMWPaymentStatus,
	PMWMultisigAccountConfigured,
	PMWFeeProof,
	PMWMultisigUtxoConfigured,
	PMWUtxoProposalCheck,
	BtcDeposit,
	BtcPayment,
}

// i-th method correspond to a method in TeeDataConnectorStruct interface whose
// input is the proof type of i-th attestation type.
var attestationTypeMethods = []string{
	"availabilityCheck",
	"pmwPaymentStatus",
	"pmwMultisigAccountConfigured",
	"pmwFeeProof",
	"pmwMultisigUtxoConfigured",
	"pmwUtxoProposalCheck",
	"btcDeposit",
	"btcPayment",
}

type AttestationArguments struct {
	Request  abi.Argument
	Response abi.Argument
	Proof    abi.Argument
}

var RequestHeaderArg abi.Argument
var ResponseHeaderArg abi.Argument
var AttestationRequestArg abi.Argument

var AttestationTypeArguments map[AttestationType]AttestationArguments

func init() {
	fdc2ABI, err := Fdc2MetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("error getting tee data connector abi: %v", err))
	}

	if len(methods) != len(opCommands) {
		panic("methods, opCommands miss match")
	}

	MessageArguments = make(map[op.Command]abi.Argument)
	for j := range opCommands {
		method, ok := fdc2ABI.Methods[methods[j]]
		if !ok {
			panic(fmt.Sprintf("missing method %s", methods[j]))
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}

	if len(attestationTypes) != len(attestationTypeMethods) {
		panic("attestationTypes, attestationTypeMethods miss match")
	}

	AttestationTypeArguments = make(map[AttestationType]AttestationArguments)

	for j := range attestationTypes {
		request := attestationTypeMethods[j] + reqStruct
		method, ok := fdc2ABI.Methods[request]
		if !ok {
			panic(fmt.Sprintf("missing method %s", request))
		}
		requestABI := method.Inputs[0]

		response := attestationTypeMethods[j] + resStruct
		method, ok = fdc2ABI.Methods[response]
		if !ok {
			panic(fmt.Sprintf("missing method %s", response))
		}
		responseABI := method.Inputs[0]

		proof := attestationTypeMethods[j] + proofStruct
		method, ok = fdc2ABI.Methods[proof]
		if !ok {
			panic(fmt.Sprintf("missing method %s", proof))
		}
		proofABI := method.Inputs[0]

		AttestationTypeArguments[attestationTypes[j]] = AttestationArguments{
			Request:  requestABI,
			Response: responseABI,
			Proof:    proofABI,
		}
	}

	method, ok := fdc2ABI.Methods["fdc2RequestHeaderStruct"]
	if !ok {
		panic("missing method fdc2RequestHeaderStruct")
	}
	RequestHeaderArg = method.Inputs[0]

	method, ok = fdc2ABI.Methods["fdc2ResponseHeaderStruct"]
	if !ok {
		panic("missing method fdc2ResponseHeaderStruct")
	}
	ResponseHeaderArg = method.Inputs[0]

	method, ok = fdc2ABI.Methods["fdc2AttestationRequestStruct"]
	if !ok {
		panic("missing method fdc2AttestationRequestStruct")
	}
	AttestationRequestArg = method.Inputs[0]
}
