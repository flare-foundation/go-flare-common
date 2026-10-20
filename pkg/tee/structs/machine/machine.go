//go:generate  abigen --abi=machine.abi --pkg=machine --type=TeeMachine --out=autogen.go
package machine

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/contracts"
)

var TeeMachineDataStructArg abi.Argument

func init() {
	machineABI := contracts.MustABI(TeeMachineMetaData)
	TeeMachineDataStructArg = contracts.MustMethod(machineABI, "teeMachineDataStruct").Inputs[0]
}
