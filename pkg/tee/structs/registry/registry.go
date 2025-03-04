//go:generate  abigen --abi=pkg/tee/structs/registry/registry.abi --pkg=registry --type=Registry --out=pkg/tee/structs/registry/autogen.go
package registry

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

const OPType = "REG"

type OPCommand string

const (
	AvailabilityCheck OPCommand = "AVAILABILITY_CHECK"
	ToPauseForUpgrade OPCommand = "TO_PAUSE_FOR_UPGRADE"
	ReplicateFrom     OPCommand = "REPLICATE_FROM"
)

var opCommands = []OPCommand{
	AvailabilityCheck,
	ToPauseForUpgrade,
	ReplicateFrom,
}

// i-th method correspond to a method in TeeRegistryStruct interface whose
// input is the type of message emitted with i-th opCommands
var methods = []string{
	"availabilityCheckResponseStruct",
	"pauseForUpgradeStruct",
	"replicateTeeMachineStruct",
}

var MessageArguments map[OPCommand]abi.Argument

func init() {
	registryAbi, err := RegistryMetaData.GetAbi()
	if err != nil {
		logger.Panicf("error getting registry abi: %v", err)
	}

	if len(methods) != len(opCommands) {
		logger.Panicf("methods, opCommands miss match")
	}

	MessageArguments = make(map[OPCommand]abi.Argument)

	for j := range opCommands {
		method, ok := registryAbi.Methods[methods[j]]
		if !ok {
			logger.Panicf("missing method %s", methods[j])
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}
}
