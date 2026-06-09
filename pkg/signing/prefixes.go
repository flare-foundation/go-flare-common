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
	TEEMachineRegister Prefix = mustStringBytes32("TEE_MACHINE_REGISTER")
	TEEKeyExistence    Prefix = mustStringBytes32("TEE_KEY_EXISTENCE")
	TEEActionResult    Prefix = mustStringBytes32("TEE_ACTION_RESULT")
	TEEVoteHash        Prefix = mustStringBytes32("TEE_VOTE_HASH")
	TEEKeyDirectBackup Prefix = mustStringBytes32("TEE_KEY_DIRECT_BACKUP")
	TEEWalletBackup    Prefix = mustStringBytes32("TEE_WALLET_BACKUP")
	PMWWalletBackup    Prefix = mustStringBytes32("PMW_WALLET_BACKUP")
	PMWKeySplit        Prefix = mustStringBytes32("PMW_KEY_SPLIT")
	ProxyActionResult  Prefix = mustStringBytes32("PROXY_ACTION_RESULT")
	ProxyTeeInfo       Prefix = mustStringBytes32("PROXY_TEE_INFO")
	ProxyVoteReceipt   Prefix = mustStringBytes32("PROXY_VOTE_RECEIPT")
)

func mustStringBytes32(s string) [32]byte {
	if len(s) > 32 {
		panic(fmt.Sprintf("string %s longer than 32", s))
	}
	x := [32]byte{}
	copy(x[:], s)

	return x
}
