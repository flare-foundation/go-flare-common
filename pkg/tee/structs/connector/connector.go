//go:generate  abigen --abi=connector.abi --pkg=connector --type=Connector --out=autogen.go
package connector

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
)

var attestationTypes = []AttestationType{
	AvailabilityCheck,
	PMWPaymentStatus,
	PMWMultisigAccountConfigured,
}

// i-th method correspond to a method in TeeDataConnectorStruct interface whose
// input is the proof type of i-th attestation type.
var attestationTypeMethods = []string{
	"availabilityCheck",
	"pmwPaymentStatus",
	"pmwMultisigAccountConfigured",
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
	connectorABI, err := ConnectorMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("error getting tee data connector abi: %v", err))
	}

	if len(methods) != len(opCommands) {
		panic("methods, opCommands miss match")
	}

	MessageArguments = make(map[op.Command]abi.Argument)
	for j := range opCommands {
		method, ok := connectorABI.Methods[methods[j]]
		if !ok {
			panic(fmt.Sprintf("missing method %s", methods[j]))
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}

	AttestationTypeArguments = make(map[AttestationType]AttestationArguments)

	for j := range attestationTypes {
		request := attestationTypeMethods[j] + reqStruct
		method, ok := connectorABI.Methods[request]
		if !ok {
			panic(fmt.Sprintf("missing method %s", request))
		}
		requestABI := method.Inputs[0]

		response := attestationTypeMethods[j] + resStruct
		method, ok = connectorABI.Methods[response]
		if !ok {
			panic(fmt.Sprintf("missing method %s", response))
		}
		responseABI := method.Inputs[0]

		proof := attestationTypeMethods[j] + proofStruct
		method, ok = connectorABI.Methods[proof]
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

	method, ok := connectorABI.Methods["fdc2RequestHeaderStruct"]
	if !ok {
		panic("missing method fdc2RequestHeaderStruct")
	}
	RequestHeaderArg = method.Inputs[0]

	method, ok = connectorABI.Methods["fdc2ResponseHeaderStruct"]
	if !ok {
		panic("missing method fdc2ResponseHeaderStruct")
	}
	ResponseHeaderArg = method.Inputs[0]

	method, ok = connectorABI.Methods["fdc2AttestationRequestStruct"]
	if !ok {
		panic("missing method fdc2AttestationRequestStruct")
	}
	AttestationRequestArg = method.Inputs[0]
}
