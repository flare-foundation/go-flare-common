package signing

import (
	"fmt"
)

type Prefix [32]byte

var (
	TEEInstruction Prefix = mustStringBytes32("TEE_INSTRUCTION")

	// must be aligned with smart contracts
	FDC2               Prefix = mustStringBytes32("FDC2")
	TEEPausingAddress  Prefix = mustStringBytes32("TEE_PAUSING_ADDRESSES")
	TEEMachinePathList Prefix = mustStringBytes32("TEE_MACHINE_PATH_LIST")
)

func mustStringBytes32(s string) [32]byte {
	if len(s) > 32 {
		panic(fmt.Sprintf("string %s longer than 32", s))
	}
	x := [32]byte{}
	copy(x[:], s)

	return x
}
