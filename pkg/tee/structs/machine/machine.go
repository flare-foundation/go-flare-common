//go:generate  abigen --abi=machine.abi --pkg=machine --type=TeeMachine --out=autogen.go
package machine

import "github.com/ethereum/go-ethereum/accounts/abi"

var TeeMachineDataStructArg abi.Argument

func init() {
	machineABI, err := TeeMachineMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	method, ok := machineABI.Methods["teeMachineDataStruct"]
	if !ok {
		panic("missing method teeMachineDataStruct")
	}

	TeeMachineDataStructArg = method.Inputs[0]
}
