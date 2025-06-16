//go:generate  abigen --abi=wallet.abi --pkg=wallet --type=Wallet --out=autogen.go
package wallet

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

const OPType = "WALLET"

type OPCommand string

const (
	KeyGenerate            OPCommand = "KEY_GENERATE"
	KeyDelete              OPCommand = "KEY_DELETE"
	KeyDataProviderRestore OPCommand = "KEY_DATA_PROVIDER_RESTORE"
	SetPausingAddress      OPCommand = "SET_PAUSING_ADDRESSES"
	Resume                 OPCommand = "RESUME"
)

var opCommands = []OPCommand{
	KeyGenerate,
	KeyDelete,
	KeyDataProviderRestore,
	SetPausingAddress,
	Resume,
}

// i-th method correspond to a method in TeeWalletStruct interface whose
// input is the type of message emitted with i-th opCommands.
var methods = []string{
	"keyGenerateStruct",
	"keyDeleteStruct",
	"keyDataProviderRestoreStruct",
	"setPausingAddressesStruct",
	"resumeStruct",
}

var MessageArguments map[OPCommand]abi.Argument

var KeyExistenceStructArg abi.Argument

func init() {
	walletmanagerAbi, err := WalletMetaData.GetAbi()
	if err != nil {
		logger.Panicf("error getting registry abi: %v", err)
	}

	if len(methods) != len(opCommands) {
		logger.Panicf("methods, opCommands miss match")
	}

	MessageArguments = make(map[OPCommand]abi.Argument)
	for j := range opCommands {
		method, ok := walletmanagerAbi.Methods[methods[j]]
		if !ok {
			logger.Panicf("missing method %s", methods[j])
		}
		MessageArguments[opCommands[j]] = method.Inputs[0]
	}

	method, ok := walletmanagerAbi.Methods["keyExistenceStruct"]
	if !ok {
		logger.Panicf("missing method %s", "keyExistenceStruct")
	}

	KeyExistenceStructArg = method.Inputs[0]
}
