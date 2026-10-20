//go:generate  abigen --abi=wallet.abi --pkg=wallet --type=TeeWallet --out=autogen.go
package wallet

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/contracts"
	"github.com/flare-foundation/go-flare-common/pkg/tee/op"
)

var opCommands = []op.Command{
	op.KeyDataProviderRestore,
	// constants.KeyDataProviderRestoreTest, // TODO
	op.KeyDelete,
	op.KeyDirectBackup,
	op.KeyDirectRestore,
	op.KeyGenerate,
}

// i-th method correspond to a method in TeeWalletStruct interface whose
// input is the type of message emitted with i-th opCommands.
var methods = []string{
	"keyDataProviderRestoreStruct",
	"keyDeleteStruct",
	"keyDirectBackupStruct",
	"keyDirectRestoreStruct",
	"keyGenerateStruct",
}

var MessageArguments map[op.Command]abi.Argument

var KeyExistenceStructArg abi.Argument
var BackupIdStructArg abi.Argument

func init() {
	walletABI := contracts.MustABI(TeeWalletMetaData)

	if len(methods) != len(opCommands) {
		panic("methods, opCommands miss match")
	}

	MessageArguments = make(map[op.Command]abi.Argument)
	for j := range opCommands {
		MessageArguments[opCommands[j]] = contracts.MustMethod(walletABI, methods[j]).Inputs[0]
	}

	KeyExistenceStructArg = contracts.MustMethod(walletABI, "keyExistenceStruct").Inputs[0]
	BackupIdStructArg = contracts.MustMethod(walletABI, "backupIdStruct").Inputs[0]
}
