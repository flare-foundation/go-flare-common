//go:generate  abigen --abi=machineregistry.abi --pkg=machineregistry --type=MachineRegistry --out=autogen.go
package machineregistry

import "github.com/ethereum/go-ethereum/accounts/abi"

var TeeMachineDataStructArg abi.Argument

func init() {
	machineregistryABI, err := MachineRegistryMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	method, ok := machineregistryABI.Methods["teeMachineDataStruct"]
	if !ok {
		panic("missing method teeMachineDataStruct")
	}

	TeeMachineDataStructArg = method.Inputs[0]
}
