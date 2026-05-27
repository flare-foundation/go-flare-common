//go:generate  abigen --abi=machinepath.abi --pkg=machinepath --type=TeeMachinePath --out=autogen.go
package machinepath

import "github.com/ethereum/go-ethereum/accounts/abi"

// MachinePathStructArg is the abi.Argument for IMachinePathManager.MachinePath,
// used to ABI-decode the path payload that off-chain relays forward alongside a
// direct backup/restore instruction.
var MachinePathStructArg abi.Argument

func init() {
	machinePathABI, err := TeeMachinePathMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	method, ok := machinePathABI.Methods["machinePathStruct"]
	if !ok {
		panic("missing method machinePathStruct")
	}

	MachinePathStructArg = method.Inputs[0]
}
