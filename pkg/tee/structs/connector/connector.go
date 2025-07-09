//go:generate  abigen --abi=connector.abi --pkg=connector --type=Connector --out=autogen.go
package connector

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/flare-foundation/go-flare-common/pkg/tee/constants"
)

const (
	reqStruct   string = "RequestBodyStruct"
	resStruct   string = "ResponseBodyStruct"
	proofStruct string = "ProofStruct"
)

var opCommands = []constants.OPCommand{
	constants.Prove,
}

// i-th method correspond to a method in TeeDataConnectorStruct interface whose
// input is the type of message emitted with i-th element of opCommands.
var methods = []string{
	"ftdcAttestationRequestStruct",
}

var MessageArguments map[constants.OPCommand]abi.Argument

type AttestationType string

const (
	AvailabilityCheck AttestationType = "TeeAvailabilityCheck"
	PMWPaymentStatus  AttestationType = "PMWPaymentStatus"
)

var attestationTypes = []AttestationType{
	AvailabilityCheck,
	PMWPaymentStatus,
}

// i-th method correspond to a method in TeeDataConnectorStruct interface whose
// input is the proof type of i-th attestation type.
var attestationTypeMethods = []string{
	"availabilityCheck",
	"pmwPaymentStatus",
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
	connectorAbi, err := ConnectorMetaData.GetAbi()
	if err != nil {
		logger.Panicf("error getting tee data connector abi: %v", err)
	}

	if len(methods) != len(opCommands) {
		logger.Panicf("methods, opCommands miss match")
	}

	MessageArguments = make(map[constants.OPCommand]abi.Argument)
	for j := range opCommands {
		method, ok := connectorAbi.Methods[methods[j]]
		if !ok {
			logger.Panicf("missing method %s", methods[j])
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}

	AttestationTypeArguments = make(map[AttestationType]AttestationArguments)

	for j := range attestationTypes {
		request := attestationTypeMethods[j] + reqStruct
		method, ok := connectorAbi.Methods[request]
		if !ok {
			logger.Panicf("missing method %s", request)
		}
		requestAbi := method.Inputs[0]

		response := attestationTypeMethods[j] + resStruct
		method, ok = connectorAbi.Methods[response]
		if !ok {
			logger.Panicf("missing method %s", response)
		}
		responseAbi := method.Inputs[0]

		proof := attestationTypeMethods[j] + proofStruct
		method, ok = connectorAbi.Methods[proof]
		if !ok {
			logger.Panicf("missing method %s", proof)
		}
		proofAbi := method.Inputs[0]

		AttestationTypeArguments[attestationTypes[j]] = AttestationArguments{
			Request:  requestAbi,
			Response: responseAbi,
			Proof:    proofAbi,
		}
	}

	method, ok := connectorAbi.Methods["ftdcRequestHeaderStruct"]
	if !ok {
		logger.Panic("missing method ftdcRequestHeaderStruct")
	}
	RequestHeaderArg = method.Inputs[0]

	method, ok = connectorAbi.Methods["ftdcResponseHeaderStruct"]
	if !ok {
		logger.Panic("missing method ftdcResponseHeaderStruct")
	}
	ResponseHeaderArg = method.Inputs[0]

	method, ok = connectorAbi.Methods["ftdcAttestationRequestStruct"]
	if !ok {
		logger.Panic("missing method ftdcAttestationRequestStruct")
	}
	AttestationRequestArg = method.Inputs[0]
}
