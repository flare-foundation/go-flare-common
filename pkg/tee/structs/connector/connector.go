//go:generate  abigen --abi=connector.abi --pkg=connector --type=Connector --out=autogen.go
package connector

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

type OPCommand string

const (
	Prove OPCommand = "PROVE"
)

var opCommands = []OPCommand{
	Prove,
}

// i-th method correspond to a method in TeeDataConnectorStruct interface whose
// input is the type of message emitted with i-th opCommands
var methods = []string{
	"ftdcProveStruct",
}

var MessageArguments map[OPCommand]abi.Argument

func init() {
	paymentAbi, err := ConnectorMetaData.GetAbi()
	if err != nil {
		logger.Panicf("error getting tee data connector abi: %v", err)
	}

	if len(methods) != len(opCommands) {
		logger.Panicf("methods, opCommands miss match")
	}

	MessageArguments = make(map[OPCommand]abi.Argument)
	for j := range opCommands {
		method, ok := paymentAbi.Methods[methods[j]]
		if !ok {
			logger.Panicf("missing method %s", methods[j])
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}
}
