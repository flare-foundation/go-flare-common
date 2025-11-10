//go:generate  abigen --abi=wallet.abi --pkg=wallet --type=Wallet --out=autogen.go
package wallet

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/tee/op"
)

var opCommands = []op.Command{
	op.KeyDataProviderRestore,
	// constants.KeyDataProviderRestoreTest, // TODO
	op.KeyDelete,
	op.KeyGenerate,
}

// i-th method correspond to a method in TeeWalletStruct interface whose
// input is the type of message emitted with i-th opCommands.
var methods = []string{
	"keyDataProviderRestoreStruct",
	"keyDeleteStruct",
	"keyGenerateStruct",
}

var MessageArguments map[op.Command]abi.Argument

var KeyExistenceStructArg abi.Argument
var BackupIdStructArg abi.Argument

func init() {
	walletmanagerAbi, err := WalletMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("error getting registry abi: %v", err))
	}

	if len(methods) != len(opCommands) {
		panic("methods, opCommands miss match")
	}

	MessageArguments = make(map[op.Command]abi.Argument)
	for j := range opCommands {
		method, ok := walletmanagerAbi.Methods[methods[j]]
		if !ok {
			panic(fmt.Sprintf("missing method %s", methods[j]))
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}

	method, ok := walletmanagerAbi.Methods["keyExistenceStruct"]
	if !ok {
		panic(fmt.Sprintf("missing method %s", "keyExistenceStruct"))
	}

	KeyExistenceStructArg = method.Inputs[0]

	method, ok = walletmanagerAbi.Methods["backupIdStruct"]
	if !ok {
		panic(fmt.Sprintf("missing method %s", "backupIdStruct"))
	}

	BackupIdStructArg = method.Inputs[0]
}
