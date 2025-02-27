//go:generate  abigen --abi=pkg/tee/structs/walletmanager/walletmanager.abi --pkg=walletmanager --type=Walletmanager --out=pkg/tee/structs/walletmanager/autogen.go
package walletmanager

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

const OPType = "WALLET"

type OPCommand string

const (
	KeyGeneration          OPCommand = "KEY_GENERATE"
	KeyDelete              OPCommand = "KEY_DELETE"
	KeyMachineBackup       OPCommand = "KEY_MACHINE_BACKUP"
	KeyMachineRestore      OPCommand = "KEY_MACHINE_RESTORE"
	KeyMachineBackupRemove OPCommand = "KEY_MACHINE_BACKUP_REMOVE"
	KeyCustodianBackup     OPCommand = "KEY_CUSTODIAN_BACKUP"
	KeyCustodianRestore    OPCommand = "KEY_CUSTODIAN_RESTORE"
)

var opCommands = []OPCommand{
	KeyGeneration,
	KeyDelete,
	KeyMachineBackup,
	KeyMachineRestore,
	KeyMachineBackupRemove,
	KeyCustodianBackup,
	KeyCustodianRestore,
}

// i-th method correspond to a method in TeeRegistryStruct interface whose
// input is the type of message emitted with i-th opCommands
var methods = []string{
	"keyGenerateStruct",
	"keyDeleteStruct",
	"keyMachineBackupStruct",
	"keyMachineRestoreStruct",
	"keyMachineBackupRemoveStruct",
	"keyCustodianBackupStruct",
	"keyCustodianRestoreStruct",
}

var MessageArguments map[OPCommand]abi.Argument

func init() {
	walletmanagerAbi, err := WalletmanagerMetaData.GetAbi()
	if err != nil {
		logger.Panicf("error getting registry abi: %v", err)
	}

	if len(methods) != len(opCommands) {
		logger.Panicf("methods, opCommands miss match")
	}

	for j := range opCommands {
		method, ok := walletmanagerAbi.Methods[methods[j]]
		if !ok {
			logger.Panicf("missing method %s", methods[j])
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}
}
