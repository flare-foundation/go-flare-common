//go:generate  abigen --abi=machinepath.abi --pkg=machinepath --type=TeeMachinePath --out=autogen.go
package machinepath

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/contracts"
)

// MachinePathStructArg is the abi.Argument for IMachinePathManager.MachinePath,
// used to ABI-decode the path payload that off-chain relays forward alongside a
// direct backup/restore instruction.
var MachinePathStructArg abi.Argument

func init() {
	machinePathABI := contracts.MustABI(TeeMachinePathMetaData)
	MachinePathStructArg = contracts.MustMethod(machinePathABI, "machinePathStruct").Inputs[0]
}
