//go:generate  abigen --abi=instructions.abi --pkg=instructions --type=TeeInstructions --out=autogen.go
package instructions

import "github.com/ethereum/go-ethereum/accounts/abi"

// TeeInstructionParamsStructArg is the abi.Argument for ITeeInstructions.TeeInstructionParams,
// used to ABI-encode/decode the parameters off-chain relays carry alongside an instruction
// submission.
var TeeInstructionParamsStructArg abi.Argument

func init() {
	instructionsABI, err := TeeInstructionsMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	method, ok := instructionsABI.Methods["teeInstructionParamsStruct"]
	if !ok {
		panic("missing method teeInstructionParamsStruct")
	}

	TeeInstructionParamsStructArg = method.Inputs[0]
}
