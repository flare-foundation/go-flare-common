//go:generate  abigen --abi=fdc2.abi --pkg=fdc2 --type=Fdc2 --out=autogen.go
package fdc2

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/contracts"
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
)

var attestationTypes = []AttestationType{
	AvailabilityCheck,
	PMWPaymentStatus,
	PMWMultisigAccountConfigured,
	PMWFeeProof,
}

// i-th method correspond to a method in TeeDataConnectorStruct interface whose
// input is the proof type of i-th attestation type.
var attestationTypeMethods = []string{
	"availabilityCheck",
	"pmwPaymentStatus",
	"pmwMultisigAccountConfigured",
	"pmwFeeProof",
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
	fdc2ABI := contracts.MustABI(Fdc2MetaData)

	if len(methods) != len(opCommands) {
		panic("methods, opCommands miss match")
	}

	MessageArguments = make(map[op.Command]abi.Argument)
	for j := range opCommands {
		MessageArguments[opCommands[j]] = contracts.MustMethod(fdc2ABI, methods[j]).Inputs[0]
	}

	if len(attestationTypes) != len(attestationTypeMethods) {
		panic("attestationTypes, attestationTypeMethods miss match")
	}

	AttestationTypeArguments = make(map[AttestationType]AttestationArguments)

	for j := range attestationTypes {
		name := attestationTypeMethods[j]
		AttestationTypeArguments[attestationTypes[j]] = AttestationArguments{
			Request:  contracts.MustMethod(fdc2ABI, name+reqStruct).Inputs[0],
			Response: contracts.MustMethod(fdc2ABI, name+resStruct).Inputs[0],
			Proof:    contracts.MustMethod(fdc2ABI, name+proofStruct).Inputs[0],
		}
	}

	RequestHeaderArg = contracts.MustMethod(fdc2ABI, "fdc2RequestHeaderStruct").Inputs[0]
	ResponseHeaderArg = contracts.MustMethod(fdc2ABI, "fdc2ResponseHeaderStruct").Inputs[0]
	AttestationRequestArg = contracts.MustMethod(fdc2ABI, "fdc2AttestationRequestStruct").Inputs[0]
}
