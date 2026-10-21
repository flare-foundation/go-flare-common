//go:generate  abigen --abi=instructions.abi --pkg=instructions --type=TeeInstructions --out=autogen.go
package instructions

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/contracts"
)

// TeeInstructionParamsStructArg is the abi.Argument for ITeeInstructions.TeeInstructionParams,
// used to ABI-encode/decode the parameters off-chain relays carry alongside an instruction
// submission.
var TeeInstructionParamsStructArg abi.Argument

func init() {
	instructionsABI := contracts.MustABI(TeeInstructionsMetaData)
	TeeInstructionParamsStructArg = contracts.MustMethod(instructionsABI, "teeInstructionParamsStruct").Inputs[0]
}
