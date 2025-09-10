// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teeextensionregistry

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ITeeMachineRegistryTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type ITeeMachineRegistryTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
}

// TeeExtensionRegistryMetaData contains all meta data concerning the TeeExtensionRegistry contract.
var TeeExtensionRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CodeHashZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InstructionIdEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCodeHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInstructionsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPlatform\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPlatforms\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyInstructionsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OpTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeConstantsProviderNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"PlatformAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PlatformEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"name\":\"SystemOpTypeNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SystemOwnedExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"UnsupportedPlatform\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionEmpty\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"CodeHashPlatformDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"SupportedPlatformAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"name\":\"SupportedWalletProjectOpTypeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"name\":\"SupportedWalletProjectOpTypeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"teeExtensionStateVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"TeeExtensionContractsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"TeeExtensionRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structITeeMachineRegistry.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"platforms\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"governanceHash\",\"type\":\"bytes32\"}],\"name\":\"TeeVersionAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYSTEM_OP_TYPE_PREFIX\",\"outputs\":[{\"internalType\":\"bytes2\",\"name\":\"\",\"type\":\"bytes2\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"contractITeeWalletProjectOpTypeConstants[]\",\"name\":\"_opTypeConstantsProviders\",\"type\":\"address[]\"}],\"name\":\"addOrUpdateSupportedWalletProjectOpTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"}],\"name\":\"addSupportedPlatforms\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"addTeeVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"codeHashPlatformDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"disableCodeHashPlatform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extensionsCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"}],\"name\":\"getCodeHashInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getExtensionOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getSupportedWalletProjectOpTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_supportedOpTypes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSystemInstructionInitiators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getTeeExtensionInstructionsSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getTeeExtensionStateVerifier\",\"outputs\":[{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"}],\"name\":\"getTeeGovernanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"}],\"name\":\"getWalletProjectOpTypeConstantsProvider\",\"outputs\":[{\"internalType\":\"contractITeeWalletProjectOpTypeConstants\",\"name\":\"_opTypeConstantsProvider\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"isCodeHashPlatformSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"}],\"name\":\"isWalletProjectOpTypeSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"proposedExtensionOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"_teeExtensionStateVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_instructionInitiators\",\"type\":\"address[]\"}],\"name\":\"registerSystemInstructionInitiators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_opTypes\",\"type\":\"bytes32[]\"}],\"name\":\"removeSupportedWalletProjectOpTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"sendInstructions\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"_teeExtensionStateVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"setExtensionContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeFeeCalculator\",\"outputs\":[{\"internalType\":\"contractITeeFeeCalculator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeGovernance\",\"outputs\":[{\"internalType\":\"contractITeeGovernance\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeMachineRegistry\",\"outputs\":[{\"internalType\":\"contractITeeMachineRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_instructionInitiators\",\"type\":\"address[]\"}],\"name\":\"unregisterSystemInstructionInitiators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeeExtensionRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeExtensionRegistryMetaData.ABI instead.
var TeeExtensionRegistryABI = TeeExtensionRegistryMetaData.ABI

// TeeExtensionRegistry is an auto generated Go binding around an Ethereum contract.
type TeeExtensionRegistry struct {
	TeeExtensionRegistryCaller     // Read-only binding to the contract
	TeeExtensionRegistryTransactor // Write-only binding to the contract
	TeeExtensionRegistryFilterer   // Log filterer for contract events
}

// TeeExtensionRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeExtensionRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeExtensionRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeExtensionRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeExtensionRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeExtensionRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeExtensionRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeExtensionRegistrySession struct {
	Contract     *TeeExtensionRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TeeExtensionRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeExtensionRegistryCallerSession struct {
	Contract *TeeExtensionRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TeeExtensionRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeExtensionRegistryTransactorSession struct {
	Contract     *TeeExtensionRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TeeExtensionRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeExtensionRegistryRaw struct {
	Contract *TeeExtensionRegistry // Generic contract binding to access the raw methods on
}

// TeeExtensionRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeExtensionRegistryCallerRaw struct {
	Contract *TeeExtensionRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TeeExtensionRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeExtensionRegistryTransactorRaw struct {
	Contract *TeeExtensionRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeExtensionRegistry creates a new instance of TeeExtensionRegistry, bound to a specific deployed contract.
func NewTeeExtensionRegistry(address common.Address, backend bind.ContractBackend) (*TeeExtensionRegistry, error) {
	contract, err := bindTeeExtensionRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistry{TeeExtensionRegistryCaller: TeeExtensionRegistryCaller{contract: contract}, TeeExtensionRegistryTransactor: TeeExtensionRegistryTransactor{contract: contract}, TeeExtensionRegistryFilterer: TeeExtensionRegistryFilterer{contract: contract}}, nil
}

// NewTeeExtensionRegistryCaller creates a new read-only instance of TeeExtensionRegistry, bound to a specific deployed contract.
func NewTeeExtensionRegistryCaller(address common.Address, caller bind.ContractCaller) (*TeeExtensionRegistryCaller, error) {
	contract, err := bindTeeExtensionRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryCaller{contract: contract}, nil
}

// NewTeeExtensionRegistryTransactor creates a new write-only instance of TeeExtensionRegistry, bound to a specific deployed contract.
func NewTeeExtensionRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeExtensionRegistryTransactor, error) {
	contract, err := bindTeeExtensionRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryTransactor{contract: contract}, nil
}

// NewTeeExtensionRegistryFilterer creates a new log filterer instance of TeeExtensionRegistry, bound to a specific deployed contract.
func NewTeeExtensionRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeExtensionRegistryFilterer, error) {
	contract, err := bindTeeExtensionRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryFilterer{contract: contract}, nil
}

// bindTeeExtensionRegistry binds a generic wrapper to an already deployed contract.
func bindTeeExtensionRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeExtensionRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeExtensionRegistry *TeeExtensionRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeExtensionRegistry.Contract.TeeExtensionRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeExtensionRegistry *TeeExtensionRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.TeeExtensionRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeExtensionRegistry *TeeExtensionRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.TeeExtensionRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeExtensionRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.contract.Transact(opts, method, params...)
}

// SYSTEMOPTYPEPREFIX is a free data retrieval call binding the contract method 0x88b135a9.
//
// Solidity: function SYSTEM_OP_TYPE_PREFIX() view returns(bytes2)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) SYSTEMOPTYPEPREFIX(opts *bind.CallOpts) ([2]byte, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "SYSTEM_OP_TYPE_PREFIX")

	if err != nil {
		return *new([2]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2]byte)).(*[2]byte)

	return out0, err

}

// SYSTEMOPTYPEPREFIX is a free data retrieval call binding the contract method 0x88b135a9.
//
// Solidity: function SYSTEM_OP_TYPE_PREFIX() view returns(bytes2)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) SYSTEMOPTYPEPREFIX() ([2]byte, error) {
	return _TeeExtensionRegistry.Contract.SYSTEMOPTYPEPREFIX(&_TeeExtensionRegistry.CallOpts)
}

// SYSTEMOPTYPEPREFIX is a free data retrieval call binding the contract method 0x88b135a9.
//
// Solidity: function SYSTEM_OP_TYPE_PREFIX() view returns(bytes2)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) SYSTEMOPTYPEPREFIX() ([2]byte, error) {
	return _TeeExtensionRegistry.Contract.SYSTEMOPTYPEPREFIX(&_TeeExtensionRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeExtensionRegistry.Contract.UPGRADEINTERFACEVERSION(&_TeeExtensionRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeExtensionRegistry.Contract.UPGRADEINTERFACEVERSION(&_TeeExtensionRegistry.CallOpts)
}

// CodeHashPlatformDisabled is a free data retrieval call binding the contract method 0x010f6a8b.
//
// Solidity: function codeHashPlatformDisabled(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) CodeHashPlatformDisabled(opts *bind.CallOpts, _extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "codeHashPlatformDisabled", _extensionId, _codeHash, _platform)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CodeHashPlatformDisabled is a free data retrieval call binding the contract method 0x010f6a8b.
//
// Solidity: function codeHashPlatformDisabled(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) CodeHashPlatformDisabled(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.CodeHashPlatformDisabled(&_TeeExtensionRegistry.CallOpts, _extensionId, _codeHash, _platform)
}

// CodeHashPlatformDisabled is a free data retrieval call binding the contract method 0x010f6a8b.
//
// Solidity: function codeHashPlatformDisabled(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) CodeHashPlatformDisabled(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.CodeHashPlatformDisabled(&_TeeExtensionRegistry.CallOpts, _extensionId, _codeHash, _platform)
}

// ExtensionsCounter is a free data retrieval call binding the contract method 0xfad5902b.
//
// Solidity: function extensionsCounter() view returns(uint256)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) ExtensionsCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "extensionsCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtensionsCounter is a free data retrieval call binding the contract method 0xfad5902b.
//
// Solidity: function extensionsCounter() view returns(uint256)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) ExtensionsCounter() (*big.Int, error) {
	return _TeeExtensionRegistry.Contract.ExtensionsCounter(&_TeeExtensionRegistry.CallOpts)
}

// ExtensionsCounter is a free data retrieval call binding the contract method 0xfad5902b.
//
// Solidity: function extensionsCounter() view returns(uint256)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) ExtensionsCounter() (*big.Int, error) {
	return _TeeExtensionRegistry.Contract.ExtensionsCounter(&_TeeExtensionRegistry.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) FlareSystemsManager() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.FlareSystemsManager(&_TeeExtensionRegistry.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) FlareSystemsManager() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.FlareSystemsManager(&_TeeExtensionRegistry.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetAddressUpdater() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetAddressUpdater(&_TeeExtensionRegistry.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetAddressUpdater(&_TeeExtensionRegistry.CallOpts)
}

// GetCodeHashInfo is a free data retrieval call binding the contract method 0x62672305.
//
// Solidity: function getCodeHashInfo(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32 _governanceHash, string _version, bytes32[] _platforms)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetCodeHashInfo(opts *bind.CallOpts, _extensionId *big.Int, _codeHash [32]byte) (struct {
	GovernanceHash [32]byte
	Version        string
	Platforms      [][32]byte
}, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getCodeHashInfo", _extensionId, _codeHash)

	outstruct := new(struct {
		GovernanceHash [32]byte
		Version        string
		Platforms      [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GovernanceHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Version = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Platforms = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

// GetCodeHashInfo is a free data retrieval call binding the contract method 0x62672305.
//
// Solidity: function getCodeHashInfo(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32 _governanceHash, string _version, bytes32[] _platforms)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetCodeHashInfo(_extensionId *big.Int, _codeHash [32]byte) (struct {
	GovernanceHash [32]byte
	Version        string
	Platforms      [][32]byte
}, error) {
	return _TeeExtensionRegistry.Contract.GetCodeHashInfo(&_TeeExtensionRegistry.CallOpts, _extensionId, _codeHash)
}

// GetCodeHashInfo is a free data retrieval call binding the contract method 0x62672305.
//
// Solidity: function getCodeHashInfo(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32 _governanceHash, string _version, bytes32[] _platforms)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetCodeHashInfo(_extensionId *big.Int, _codeHash [32]byte) (struct {
	GovernanceHash [32]byte
	Version        string
	Platforms      [][32]byte
}, error) {
	return _TeeExtensionRegistry.Contract.GetCodeHashInfo(&_TeeExtensionRegistry.CallOpts, _extensionId, _codeHash)
}

// GetExtensionOwner is a free data retrieval call binding the contract method 0x5e46e380.
//
// Solidity: function getExtensionOwner(uint256 _extensionId) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetExtensionOwner(opts *bind.CallOpts, _extensionId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getExtensionOwner", _extensionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExtensionOwner is a free data retrieval call binding the contract method 0x5e46e380.
//
// Solidity: function getExtensionOwner(uint256 _extensionId) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetExtensionOwner(_extensionId *big.Int) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetExtensionOwner(&_TeeExtensionRegistry.CallOpts, _extensionId)
}

// GetExtensionOwner is a free data retrieval call binding the contract method 0x5e46e380.
//
// Solidity: function getExtensionOwner(uint256 _extensionId) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetExtensionOwner(_extensionId *big.Int) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetExtensionOwner(&_TeeExtensionRegistry.CallOpts, _extensionId)
}

// GetSupportedWalletProjectOpTypes is a free data retrieval call binding the contract method 0xd3543fb2.
//
// Solidity: function getSupportedWalletProjectOpTypes(uint256 _extensionId) view returns(bytes32[] _supportedOpTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetSupportedWalletProjectOpTypes(opts *bind.CallOpts, _extensionId *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getSupportedWalletProjectOpTypes", _extensionId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSupportedWalletProjectOpTypes is a free data retrieval call binding the contract method 0xd3543fb2.
//
// Solidity: function getSupportedWalletProjectOpTypes(uint256 _extensionId) view returns(bytes32[] _supportedOpTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetSupportedWalletProjectOpTypes(_extensionId *big.Int) ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSupportedWalletProjectOpTypes(&_TeeExtensionRegistry.CallOpts, _extensionId)
}

// GetSupportedWalletProjectOpTypes is a free data retrieval call binding the contract method 0xd3543fb2.
//
// Solidity: function getSupportedWalletProjectOpTypes(uint256 _extensionId) view returns(bytes32[] _supportedOpTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetSupportedWalletProjectOpTypes(_extensionId *big.Int) ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSupportedWalletProjectOpTypes(&_TeeExtensionRegistry.CallOpts, _extensionId)
}

// GetSystemInstructionInitiators is a free data retrieval call binding the contract method 0x4a4cbd8d.
//
// Solidity: function getSystemInstructionInitiators() view returns(address[])
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetSystemInstructionInitiators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getSystemInstructionInitiators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSystemInstructionInitiators is a free data retrieval call binding the contract method 0x4a4cbd8d.
//
// Solidity: function getSystemInstructionInitiators() view returns(address[])
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetSystemInstructionInitiators() ([]common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetSystemInstructionInitiators(&_TeeExtensionRegistry.CallOpts)
}

// GetSystemInstructionInitiators is a free data retrieval call binding the contract method 0x4a4cbd8d.
//
// Solidity: function getSystemInstructionInitiators() view returns(address[])
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetSystemInstructionInitiators() ([]common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetSystemInstructionInitiators(&_TeeExtensionRegistry.CallOpts)
}

// GetTeeExtensionInstructionsSender is a free data retrieval call binding the contract method 0x2c177358.
//
// Solidity: function getTeeExtensionInstructionsSender(uint256 _extensionId) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetTeeExtensionInstructionsSender(opts *bind.CallOpts, _extensionId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getTeeExtensionInstructionsSender", _extensionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeeExtensionInstructionsSender is a free data retrieval call binding the contract method 0x2c177358.
//
// Solidity: function getTeeExtensionInstructionsSender(uint256 _extensionId) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetTeeExtensionInstructionsSender(_extensionId *big.Int) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetTeeExtensionInstructionsSender(&_TeeExtensionRegistry.CallOpts, _extensionId)
}

// GetTeeExtensionInstructionsSender is a free data retrieval call binding the contract method 0x2c177358.
//
// Solidity: function getTeeExtensionInstructionsSender(uint256 _extensionId) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetTeeExtensionInstructionsSender(_extensionId *big.Int) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetTeeExtensionInstructionsSender(&_TeeExtensionRegistry.CallOpts, _extensionId)
}

// GetTeeExtensionStateVerifier is a free data retrieval call binding the contract method 0x709ebdcd.
//
// Solidity: function getTeeExtensionStateVerifier(uint256 _extensionId) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetTeeExtensionStateVerifier(opts *bind.CallOpts, _extensionId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getTeeExtensionStateVerifier", _extensionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeeExtensionStateVerifier is a free data retrieval call binding the contract method 0x709ebdcd.
//
// Solidity: function getTeeExtensionStateVerifier(uint256 _extensionId) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetTeeExtensionStateVerifier(_extensionId *big.Int) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetTeeExtensionStateVerifier(&_TeeExtensionRegistry.CallOpts, _extensionId)
}

// GetTeeExtensionStateVerifier is a free data retrieval call binding the contract method 0x709ebdcd.
//
// Solidity: function getTeeExtensionStateVerifier(uint256 _extensionId) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetTeeExtensionStateVerifier(_extensionId *big.Int) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetTeeExtensionStateVerifier(&_TeeExtensionRegistry.CallOpts, _extensionId)
}

// GetTeeGovernanceHash is a free data retrieval call binding the contract method 0x9f300b9b.
//
// Solidity: function getTeeGovernanceHash(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetTeeGovernanceHash(opts *bind.CallOpts, _extensionId *big.Int, _codeHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getTeeGovernanceHash", _extensionId, _codeHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTeeGovernanceHash is a free data retrieval call binding the contract method 0x9f300b9b.
//
// Solidity: function getTeeGovernanceHash(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetTeeGovernanceHash(_extensionId *big.Int, _codeHash [32]byte) ([32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetTeeGovernanceHash(&_TeeExtensionRegistry.CallOpts, _extensionId, _codeHash)
}

// GetTeeGovernanceHash is a free data retrieval call binding the contract method 0x9f300b9b.
//
// Solidity: function getTeeGovernanceHash(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetTeeGovernanceHash(_extensionId *big.Int, _codeHash [32]byte) ([32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetTeeGovernanceHash(&_TeeExtensionRegistry.CallOpts, _extensionId, _codeHash)
}

// GetWalletProjectOpTypeConstantsProvider is a free data retrieval call binding the contract method 0xca1c9601.
//
// Solidity: function getWalletProjectOpTypeConstantsProvider(uint256 _extensionId, bytes32 _opType) view returns(address _opTypeConstantsProvider)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetWalletProjectOpTypeConstantsProvider(opts *bind.CallOpts, _extensionId *big.Int, _opType [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getWalletProjectOpTypeConstantsProvider", _extensionId, _opType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWalletProjectOpTypeConstantsProvider is a free data retrieval call binding the contract method 0xca1c9601.
//
// Solidity: function getWalletProjectOpTypeConstantsProvider(uint256 _extensionId, bytes32 _opType) view returns(address _opTypeConstantsProvider)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetWalletProjectOpTypeConstantsProvider(_extensionId *big.Int, _opType [32]byte) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetWalletProjectOpTypeConstantsProvider(&_TeeExtensionRegistry.CallOpts, _extensionId, _opType)
}

// GetWalletProjectOpTypeConstantsProvider is a free data retrieval call binding the contract method 0xca1c9601.
//
// Solidity: function getWalletProjectOpTypeConstantsProvider(uint256 _extensionId, bytes32 _opType) view returns(address _opTypeConstantsProvider)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetWalletProjectOpTypeConstantsProvider(_extensionId *big.Int, _opType [32]byte) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetWalletProjectOpTypeConstantsProvider(&_TeeExtensionRegistry.CallOpts, _extensionId, _opType)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) Governance() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.Governance(&_TeeExtensionRegistry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) Governance() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.Governance(&_TeeExtensionRegistry.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GovernanceSettings() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GovernanceSettings(&_TeeExtensionRegistry.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GovernanceSettings(&_TeeExtensionRegistry.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) Implementation() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.Implementation(&_TeeExtensionRegistry.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) Implementation() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.Implementation(&_TeeExtensionRegistry.CallOpts)
}

// IsCodeHashPlatformSupported is a free data retrieval call binding the contract method 0xa1d925c6.
//
// Solidity: function isCodeHashPlatformSupported(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) IsCodeHashPlatformSupported(opts *bind.CallOpts, _extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "isCodeHashPlatformSupported", _extensionId, _codeHash, _platform)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCodeHashPlatformSupported is a free data retrieval call binding the contract method 0xa1d925c6.
//
// Solidity: function isCodeHashPlatformSupported(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) IsCodeHashPlatformSupported(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsCodeHashPlatformSupported(&_TeeExtensionRegistry.CallOpts, _extensionId, _codeHash, _platform)
}

// IsCodeHashPlatformSupported is a free data retrieval call binding the contract method 0xa1d925c6.
//
// Solidity: function isCodeHashPlatformSupported(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) IsCodeHashPlatformSupported(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsCodeHashPlatformSupported(&_TeeExtensionRegistry.CallOpts, _extensionId, _codeHash, _platform)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsExecutor(&_TeeExtensionRegistry.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsExecutor(&_TeeExtensionRegistry.CallOpts, _address)
}

// IsWalletProjectOpTypeSupported is a free data retrieval call binding the contract method 0xf2c33275.
//
// Solidity: function isWalletProjectOpTypeSupported(uint256 _extensionId, bytes32 _opType) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) IsWalletProjectOpTypeSupported(opts *bind.CallOpts, _extensionId *big.Int, _opType [32]byte) (bool, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "isWalletProjectOpTypeSupported", _extensionId, _opType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWalletProjectOpTypeSupported is a free data retrieval call binding the contract method 0xf2c33275.
//
// Solidity: function isWalletProjectOpTypeSupported(uint256 _extensionId, bytes32 _opType) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) IsWalletProjectOpTypeSupported(_extensionId *big.Int, _opType [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsWalletProjectOpTypeSupported(&_TeeExtensionRegistry.CallOpts, _extensionId, _opType)
}

// IsWalletProjectOpTypeSupported is a free data retrieval call binding the contract method 0xf2c33275.
//
// Solidity: function isWalletProjectOpTypeSupported(uint256 _extensionId, bytes32 _opType) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) IsWalletProjectOpTypeSupported(_extensionId *big.Int, _opType [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsWalletProjectOpTypeSupported(&_TeeExtensionRegistry.CallOpts, _extensionId, _opType)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) ProductionMode() (bool, error) {
	return _TeeExtensionRegistry.Contract.ProductionMode(&_TeeExtensionRegistry.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) ProductionMode() (bool, error) {
	return _TeeExtensionRegistry.Contract.ProductionMode(&_TeeExtensionRegistry.CallOpts)
}

// ProposedExtensionOwner is a free data retrieval call binding the contract method 0xd21d045e.
//
// Solidity: function proposedExtensionOwner(uint256 extensionId) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) ProposedExtensionOwner(opts *bind.CallOpts, extensionId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "proposedExtensionOwner", extensionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposedExtensionOwner is a free data retrieval call binding the contract method 0xd21d045e.
//
// Solidity: function proposedExtensionOwner(uint256 extensionId) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) ProposedExtensionOwner(extensionId *big.Int) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.ProposedExtensionOwner(&_TeeExtensionRegistry.CallOpts, extensionId)
}

// ProposedExtensionOwner is a free data retrieval call binding the contract method 0xd21d045e.
//
// Solidity: function proposedExtensionOwner(uint256 extensionId) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) ProposedExtensionOwner(extensionId *big.Int) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.ProposedExtensionOwner(&_TeeExtensionRegistry.CallOpts, extensionId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _TeeExtensionRegistry.Contract.ProxiableUUID(&_TeeExtensionRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeExtensionRegistry.Contract.ProxiableUUID(&_TeeExtensionRegistry.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) RewardManager() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.RewardManager(&_TeeExtensionRegistry.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) RewardManager() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.RewardManager(&_TeeExtensionRegistry.CallOpts)
}

// TeeFeeCalculator is a free data retrieval call binding the contract method 0xde4ae7a9.
//
// Solidity: function teeFeeCalculator() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) TeeFeeCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "teeFeeCalculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeFeeCalculator is a free data retrieval call binding the contract method 0xde4ae7a9.
//
// Solidity: function teeFeeCalculator() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) TeeFeeCalculator() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.TeeFeeCalculator(&_TeeExtensionRegistry.CallOpts)
}

// TeeFeeCalculator is a free data retrieval call binding the contract method 0xde4ae7a9.
//
// Solidity: function teeFeeCalculator() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) TeeFeeCalculator() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.TeeFeeCalculator(&_TeeExtensionRegistry.CallOpts)
}

// TeeGovernance is a free data retrieval call binding the contract method 0x36352c70.
//
// Solidity: function teeGovernance() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) TeeGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "teeGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeGovernance is a free data retrieval call binding the contract method 0x36352c70.
//
// Solidity: function teeGovernance() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) TeeGovernance() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.TeeGovernance(&_TeeExtensionRegistry.CallOpts)
}

// TeeGovernance is a free data retrieval call binding the contract method 0x36352c70.
//
// Solidity: function teeGovernance() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) TeeGovernance() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.TeeGovernance(&_TeeExtensionRegistry.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) TeeMachineRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "teeMachineRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) TeeMachineRegistry() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.TeeMachineRegistry(&_TeeExtensionRegistry.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) TeeMachineRegistry() (common.Address, error) {
	return _TeeExtensionRegistry.Contract.TeeMachineRegistry(&_TeeExtensionRegistry.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "timelockedCalls", selector)

	outstruct := new(struct {
		AllowedAfterTimestamp *big.Int
		EncodedCall           []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowedAfterTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EncodedCall = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeExtensionRegistry.Contract.TimelockedCalls(&_TeeExtensionRegistry.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeExtensionRegistry.Contract.TimelockedCalls(&_TeeExtensionRegistry.CallOpts, selector)
}

// AddOrUpdateSupportedWalletProjectOpTypes is a paid mutator transaction binding the contract method 0xffc7e814.
//
// Solidity: function addOrUpdateSupportedWalletProjectOpTypes(uint256 _extensionId, address[] _opTypeConstantsProviders) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) AddOrUpdateSupportedWalletProjectOpTypes(opts *bind.TransactOpts, _extensionId *big.Int, _opTypeConstantsProviders []common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "addOrUpdateSupportedWalletProjectOpTypes", _extensionId, _opTypeConstantsProviders)
}

// AddOrUpdateSupportedWalletProjectOpTypes is a paid mutator transaction binding the contract method 0xffc7e814.
//
// Solidity: function addOrUpdateSupportedWalletProjectOpTypes(uint256 _extensionId, address[] _opTypeConstantsProviders) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) AddOrUpdateSupportedWalletProjectOpTypes(_extensionId *big.Int, _opTypeConstantsProviders []common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.AddOrUpdateSupportedWalletProjectOpTypes(&_TeeExtensionRegistry.TransactOpts, _extensionId, _opTypeConstantsProviders)
}

// AddOrUpdateSupportedWalletProjectOpTypes is a paid mutator transaction binding the contract method 0xffc7e814.
//
// Solidity: function addOrUpdateSupportedWalletProjectOpTypes(uint256 _extensionId, address[] _opTypeConstantsProviders) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) AddOrUpdateSupportedWalletProjectOpTypes(_extensionId *big.Int, _opTypeConstantsProviders []common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.AddOrUpdateSupportedWalletProjectOpTypes(&_TeeExtensionRegistry.TransactOpts, _extensionId, _opTypeConstantsProviders)
}

// AddSupportedPlatforms is a paid mutator transaction binding the contract method 0xf08d008b.
//
// Solidity: function addSupportedPlatforms(bytes32[] _platforms) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) AddSupportedPlatforms(opts *bind.TransactOpts, _platforms [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "addSupportedPlatforms", _platforms)
}

// AddSupportedPlatforms is a paid mutator transaction binding the contract method 0xf08d008b.
//
// Solidity: function addSupportedPlatforms(bytes32[] _platforms) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) AddSupportedPlatforms(_platforms [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.AddSupportedPlatforms(&_TeeExtensionRegistry.TransactOpts, _platforms)
}

// AddSupportedPlatforms is a paid mutator transaction binding the contract method 0xf08d008b.
//
// Solidity: function addSupportedPlatforms(bytes32[] _platforms) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) AddSupportedPlatforms(_platforms [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.AddSupportedPlatforms(&_TeeExtensionRegistry.TransactOpts, _platforms)
}

// AddTeeVersion is a paid mutator transaction binding the contract method 0x65b771f7.
//
// Solidity: function addTeeVersion(uint256 _extensionId, string _version, bytes32 _codeHash, bytes32[] _platforms, bytes32 _governanceHash) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) AddTeeVersion(opts *bind.TransactOpts, _extensionId *big.Int, _version string, _codeHash [32]byte, _platforms [][32]byte, _governanceHash [32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "addTeeVersion", _extensionId, _version, _codeHash, _platforms, _governanceHash)
}

// AddTeeVersion is a paid mutator transaction binding the contract method 0x65b771f7.
//
// Solidity: function addTeeVersion(uint256 _extensionId, string _version, bytes32 _codeHash, bytes32[] _platforms, bytes32 _governanceHash) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) AddTeeVersion(_extensionId *big.Int, _version string, _codeHash [32]byte, _platforms [][32]byte, _governanceHash [32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.AddTeeVersion(&_TeeExtensionRegistry.TransactOpts, _extensionId, _version, _codeHash, _platforms, _governanceHash)
}

// AddTeeVersion is a paid mutator transaction binding the contract method 0x65b771f7.
//
// Solidity: function addTeeVersion(uint256 _extensionId, string _version, bytes32 _codeHash, bytes32[] _platforms, bytes32 _governanceHash) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) AddTeeVersion(_extensionId *big.Int, _version string, _codeHash [32]byte, _platforms [][32]byte, _governanceHash [32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.AddTeeVersion(&_TeeExtensionRegistry.TransactOpts, _extensionId, _version, _codeHash, _platforms, _governanceHash)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.CancelGovernanceCall(&_TeeExtensionRegistry.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.CancelGovernanceCall(&_TeeExtensionRegistry.TransactOpts, _selector)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x290fbb03.
//
// Solidity: function confirmOwnership(uint256 _extensionId) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) ConfirmOwnership(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "confirmOwnership", _extensionId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x290fbb03.
//
// Solidity: function confirmOwnership(uint256 _extensionId) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) ConfirmOwnership(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.ConfirmOwnership(&_TeeExtensionRegistry.TransactOpts, _extensionId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x290fbb03.
//
// Solidity: function confirmOwnership(uint256 _extensionId) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) ConfirmOwnership(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.ConfirmOwnership(&_TeeExtensionRegistry.TransactOpts, _extensionId)
}

// DisableCodeHashPlatform is a paid mutator transaction binding the contract method 0x96f54e59.
//
// Solidity: function disableCodeHashPlatform(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) DisableCodeHashPlatform(opts *bind.TransactOpts, _extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "disableCodeHashPlatform", _extensionId, _codeHash, _platform)
}

// DisableCodeHashPlatform is a paid mutator transaction binding the contract method 0x96f54e59.
//
// Solidity: function disableCodeHashPlatform(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) DisableCodeHashPlatform(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.DisableCodeHashPlatform(&_TeeExtensionRegistry.TransactOpts, _extensionId, _codeHash, _platform)
}

// DisableCodeHashPlatform is a paid mutator transaction binding the contract method 0x96f54e59.
//
// Solidity: function disableCodeHashPlatform(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) DisableCodeHashPlatform(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.DisableCodeHashPlatform(&_TeeExtensionRegistry.TransactOpts, _extensionId, _codeHash, _platform)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.ExecuteGovernanceCall(&_TeeExtensionRegistry.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.ExecuteGovernanceCall(&_TeeExtensionRegistry.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.Initialise(&_TeeExtensionRegistry.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.Initialise(&_TeeExtensionRegistry.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.Initialize(&_TeeExtensionRegistry.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.Initialize(&_TeeExtensionRegistry.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0xe947a09f.
//
// Solidity: function proposeNewOwner(uint256 _extensionId, address _newOwner) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) ProposeNewOwner(opts *bind.TransactOpts, _extensionId *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "proposeNewOwner", _extensionId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0xe947a09f.
//
// Solidity: function proposeNewOwner(uint256 _extensionId, address _newOwner) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) ProposeNewOwner(_extensionId *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.ProposeNewOwner(&_TeeExtensionRegistry.TransactOpts, _extensionId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0xe947a09f.
//
// Solidity: function proposeNewOwner(uint256 _extensionId, address _newOwner) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) ProposeNewOwner(_extensionId *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.ProposeNewOwner(&_TeeExtensionRegistry.TransactOpts, _extensionId, _newOwner)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns(uint256 _extensionId)
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) Register(opts *bind.TransactOpts, _teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "register", _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns(uint256 _extensionId)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) Register(_teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.Register(&_TeeExtensionRegistry.TransactOpts, _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns(uint256 _extensionId)
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) Register(_teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.Register(&_TeeExtensionRegistry.TransactOpts, _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// RegisterSystemInstructionInitiators is a paid mutator transaction binding the contract method 0x079568b6.
//
// Solidity: function registerSystemInstructionInitiators(address[] _instructionInitiators) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) RegisterSystemInstructionInitiators(opts *bind.TransactOpts, _instructionInitiators []common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "registerSystemInstructionInitiators", _instructionInitiators)
}

// RegisterSystemInstructionInitiators is a paid mutator transaction binding the contract method 0x079568b6.
//
// Solidity: function registerSystemInstructionInitiators(address[] _instructionInitiators) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) RegisterSystemInstructionInitiators(_instructionInitiators []common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.RegisterSystemInstructionInitiators(&_TeeExtensionRegistry.TransactOpts, _instructionInitiators)
}

// RegisterSystemInstructionInitiators is a paid mutator transaction binding the contract method 0x079568b6.
//
// Solidity: function registerSystemInstructionInitiators(address[] _instructionInitiators) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) RegisterSystemInstructionInitiators(_instructionInitiators []common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.RegisterSystemInstructionInitiators(&_TeeExtensionRegistry.TransactOpts, _instructionInitiators)
}

// RemoveSupportedWalletProjectOpTypes is a paid mutator transaction binding the contract method 0x03a0f960.
//
// Solidity: function removeSupportedWalletProjectOpTypes(uint256 _extensionId, bytes32[] _opTypes) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) RemoveSupportedWalletProjectOpTypes(opts *bind.TransactOpts, _extensionId *big.Int, _opTypes [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "removeSupportedWalletProjectOpTypes", _extensionId, _opTypes)
}

// RemoveSupportedWalletProjectOpTypes is a paid mutator transaction binding the contract method 0x03a0f960.
//
// Solidity: function removeSupportedWalletProjectOpTypes(uint256 _extensionId, bytes32[] _opTypes) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) RemoveSupportedWalletProjectOpTypes(_extensionId *big.Int, _opTypes [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.RemoveSupportedWalletProjectOpTypes(&_TeeExtensionRegistry.TransactOpts, _extensionId, _opTypes)
}

// RemoveSupportedWalletProjectOpTypes is a paid mutator transaction binding the contract method 0x03a0f960.
//
// Solidity: function removeSupportedWalletProjectOpTypes(uint256 _extensionId, bytes32[] _opTypes) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) RemoveSupportedWalletProjectOpTypes(_extensionId *big.Int, _opTypes [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.RemoveSupportedWalletProjectOpTypes(&_TeeExtensionRegistry.TransactOpts, _extensionId, _opTypes)
}

// SendInstructions is a paid mutator transaction binding the contract method 0x9b517780.
//
// Solidity: function sendInstructions(bytes32 _instructionId, address[] _teeIds, bytes32 _opType, bytes32 _opCommand, bytes _message, address[] _cosigners, uint64 _cosignersThreshold) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) SendInstructions(opts *bind.TransactOpts, _instructionId [32]byte, _teeIds []common.Address, _opType [32]byte, _opCommand [32]byte, _message []byte, _cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "sendInstructions", _instructionId, _teeIds, _opType, _opCommand, _message, _cosigners, _cosignersThreshold)
}

// SendInstructions is a paid mutator transaction binding the contract method 0x9b517780.
//
// Solidity: function sendInstructions(bytes32 _instructionId, address[] _teeIds, bytes32 _opType, bytes32 _opCommand, bytes _message, address[] _cosigners, uint64 _cosignersThreshold) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) SendInstructions(_instructionId [32]byte, _teeIds []common.Address, _opType [32]byte, _opCommand [32]byte, _message []byte, _cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.SendInstructions(&_TeeExtensionRegistry.TransactOpts, _instructionId, _teeIds, _opType, _opCommand, _message, _cosigners, _cosignersThreshold)
}

// SendInstructions is a paid mutator transaction binding the contract method 0x9b517780.
//
// Solidity: function sendInstructions(bytes32 _instructionId, address[] _teeIds, bytes32 _opType, bytes32 _opCommand, bytes _message, address[] _cosigners, uint64 _cosignersThreshold) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) SendInstructions(_instructionId [32]byte, _teeIds []common.Address, _opType [32]byte, _opCommand [32]byte, _message []byte, _cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.SendInstructions(&_TeeExtensionRegistry.TransactOpts, _instructionId, _teeIds, _opType, _opCommand, _message, _cosigners, _cosignersThreshold)
}

// SetExtensionContracts is a paid mutator transaction binding the contract method 0x6df0108f.
//
// Solidity: function setExtensionContracts(uint256 _extensionId, address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) SetExtensionContracts(opts *bind.TransactOpts, _extensionId *big.Int, _teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "setExtensionContracts", _extensionId, _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// SetExtensionContracts is a paid mutator transaction binding the contract method 0x6df0108f.
//
// Solidity: function setExtensionContracts(uint256 _extensionId, address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) SetExtensionContracts(_extensionId *big.Int, _teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.SetExtensionContracts(&_TeeExtensionRegistry.TransactOpts, _extensionId, _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// SetExtensionContracts is a paid mutator transaction binding the contract method 0x6df0108f.
//
// Solidity: function setExtensionContracts(uint256 _extensionId, address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) SetExtensionContracts(_extensionId *big.Int, _teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.SetExtensionContracts(&_TeeExtensionRegistry.TransactOpts, _extensionId, _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.SwitchToProductionMode(&_TeeExtensionRegistry.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.SwitchToProductionMode(&_TeeExtensionRegistry.TransactOpts)
}

// UnregisterSystemInstructionInitiators is a paid mutator transaction binding the contract method 0xdf57ea21.
//
// Solidity: function unregisterSystemInstructionInitiators(address[] _instructionInitiators) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) UnregisterSystemInstructionInitiators(opts *bind.TransactOpts, _instructionInitiators []common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "unregisterSystemInstructionInitiators", _instructionInitiators)
}

// UnregisterSystemInstructionInitiators is a paid mutator transaction binding the contract method 0xdf57ea21.
//
// Solidity: function unregisterSystemInstructionInitiators(address[] _instructionInitiators) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) UnregisterSystemInstructionInitiators(_instructionInitiators []common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.UnregisterSystemInstructionInitiators(&_TeeExtensionRegistry.TransactOpts, _instructionInitiators)
}

// UnregisterSystemInstructionInitiators is a paid mutator transaction binding the contract method 0xdf57ea21.
//
// Solidity: function unregisterSystemInstructionInitiators(address[] _instructionInitiators) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) UnregisterSystemInstructionInitiators(_instructionInitiators []common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.UnregisterSystemInstructionInitiators(&_TeeExtensionRegistry.TransactOpts, _instructionInitiators)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.UpdateContractAddresses(&_TeeExtensionRegistry.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.UpdateContractAddresses(&_TeeExtensionRegistry.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.UpgradeToAndCall(&_TeeExtensionRegistry.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.UpgradeToAndCall(&_TeeExtensionRegistry.TransactOpts, _newImplementation, _data)
}

// TeeExtensionRegistryCodeHashPlatformDisabledIterator is returned from FilterCodeHashPlatformDisabled and is used to iterate over the raw logs and unpacked data for CodeHashPlatformDisabled events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryCodeHashPlatformDisabledIterator struct {
	Event *TeeExtensionRegistryCodeHashPlatformDisabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistryCodeHashPlatformDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryCodeHashPlatformDisabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistryCodeHashPlatformDisabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistryCodeHashPlatformDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryCodeHashPlatformDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryCodeHashPlatformDisabled represents a CodeHashPlatformDisabled event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryCodeHashPlatformDisabled struct {
	ExtensionId *big.Int
	CodeHash    [32]byte
	Platform    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCodeHashPlatformDisabled is a free log retrieval operation binding the contract event 0xb7e18c4d8b0a9a0d658b8fd94ba32fd58527494df08bfcf74010b54518fb6777.
//
// Solidity: event CodeHashPlatformDisabled(uint256 indexed extensionId, bytes32 indexed codeHash, bytes32 indexed platform)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterCodeHashPlatformDisabled(opts *bind.FilterOpts, extensionId []*big.Int, codeHash [][32]byte, platform [][32]byte) (*TeeExtensionRegistryCodeHashPlatformDisabledIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var codeHashRule []interface{}
	for _, codeHashItem := range codeHash {
		codeHashRule = append(codeHashRule, codeHashItem)
	}
	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "CodeHashPlatformDisabled", extensionIdRule, codeHashRule, platformRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryCodeHashPlatformDisabledIterator{contract: _TeeExtensionRegistry.contract, event: "CodeHashPlatformDisabled", logs: logs, sub: sub}, nil
}

// WatchCodeHashPlatformDisabled is a free log subscription operation binding the contract event 0xb7e18c4d8b0a9a0d658b8fd94ba32fd58527494df08bfcf74010b54518fb6777.
//
// Solidity: event CodeHashPlatformDisabled(uint256 indexed extensionId, bytes32 indexed codeHash, bytes32 indexed platform)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchCodeHashPlatformDisabled(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryCodeHashPlatformDisabled, extensionId []*big.Int, codeHash [][32]byte, platform [][32]byte) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var codeHashRule []interface{}
	for _, codeHashItem := range codeHash {
		codeHashRule = append(codeHashRule, codeHashItem)
	}
	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "CodeHashPlatformDisabled", extensionIdRule, codeHashRule, platformRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryCodeHashPlatformDisabled)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "CodeHashPlatformDisabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCodeHashPlatformDisabled is a log parse operation binding the contract event 0xb7e18c4d8b0a9a0d658b8fd94ba32fd58527494df08bfcf74010b54518fb6777.
//
// Solidity: event CodeHashPlatformDisabled(uint256 indexed extensionId, bytes32 indexed codeHash, bytes32 indexed platform)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseCodeHashPlatformDisabled(log types.Log) (*TeeExtensionRegistryCodeHashPlatformDisabled, error) {
	event := new(TeeExtensionRegistryCodeHashPlatformDisabled)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "CodeHashPlatformDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryGovernanceCallTimelockedIterator struct {
	Event *TeeExtensionRegistryGovernanceCallTimelocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistryGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryGovernanceCallTimelocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistryGovernanceCallTimelocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistryGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeExtensionRegistryGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryGovernanceCallTimelockedIterator{contract: _TeeExtensionRegistry.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryGovernanceCallTimelocked)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeExtensionRegistryGovernanceCallTimelocked, error) {
	event := new(TeeExtensionRegistryGovernanceCallTimelocked)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryGovernanceInitialisedIterator struct {
	Event *TeeExtensionRegistryGovernanceInitialised // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistryGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryGovernanceInitialised)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistryGovernanceInitialised)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistryGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryGovernanceInitialised represents a GovernanceInitialised event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeExtensionRegistryGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryGovernanceInitialisedIterator{contract: _TeeExtensionRegistry.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryGovernanceInitialised)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceInitialised is a log parse operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseGovernanceInitialised(log types.Log) (*TeeExtensionRegistryGovernanceInitialised, error) {
	event := new(TeeExtensionRegistryGovernanceInitialised)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryGovernedProductionModeEnteredIterator struct {
	Event *TeeExtensionRegistryGovernedProductionModeEntered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistryGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryGovernedProductionModeEntered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistryGovernedProductionModeEntered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistryGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeeExtensionRegistryGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryGovernedProductionModeEnteredIterator{contract: _TeeExtensionRegistry.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryGovernedProductionModeEntered)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernedProductionModeEntered is a log parse operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeeExtensionRegistryGovernedProductionModeEntered, error) {
	event := new(TeeExtensionRegistryGovernedProductionModeEntered)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryNewOwnerConfirmedIterator is returned from FilterNewOwnerConfirmed and is used to iterate over the raw logs and unpacked data for NewOwnerConfirmed events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryNewOwnerConfirmedIterator struct {
	Event *TeeExtensionRegistryNewOwnerConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistryNewOwnerConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryNewOwnerConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistryNewOwnerConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistryNewOwnerConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryNewOwnerConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryNewOwnerConfirmed represents a NewOwnerConfirmed event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryNewOwnerConfirmed struct {
	ExtensionId *big.Int
	NewOwner    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerConfirmed is a free log retrieval operation binding the contract event 0x478435e37f9d5830b1c923af7693562bc1b09a86ac7b2f6acf40b57a3e316e90.
//
// Solidity: event NewOwnerConfirmed(uint256 indexed extensionId, address indexed newOwner)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterNewOwnerConfirmed(opts *bind.FilterOpts, extensionId []*big.Int, newOwner []common.Address) (*TeeExtensionRegistryNewOwnerConfirmedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "NewOwnerConfirmed", extensionIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryNewOwnerConfirmedIterator{contract: _TeeExtensionRegistry.contract, event: "NewOwnerConfirmed", logs: logs, sub: sub}, nil
}

// WatchNewOwnerConfirmed is a free log subscription operation binding the contract event 0x478435e37f9d5830b1c923af7693562bc1b09a86ac7b2f6acf40b57a3e316e90.
//
// Solidity: event NewOwnerConfirmed(uint256 indexed extensionId, address indexed newOwner)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchNewOwnerConfirmed(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryNewOwnerConfirmed, extensionId []*big.Int, newOwner []common.Address) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "NewOwnerConfirmed", extensionIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryNewOwnerConfirmed)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "NewOwnerConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewOwnerConfirmed is a log parse operation binding the contract event 0x478435e37f9d5830b1c923af7693562bc1b09a86ac7b2f6acf40b57a3e316e90.
//
// Solidity: event NewOwnerConfirmed(uint256 indexed extensionId, address indexed newOwner)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseNewOwnerConfirmed(log types.Log) (*TeeExtensionRegistryNewOwnerConfirmed, error) {
	event := new(TeeExtensionRegistryNewOwnerConfirmed)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "NewOwnerConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryNewOwnerProposedIterator is returned from FilterNewOwnerProposed and is used to iterate over the raw logs and unpacked data for NewOwnerProposed events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryNewOwnerProposedIterator struct {
	Event *TeeExtensionRegistryNewOwnerProposed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistryNewOwnerProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryNewOwnerProposed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistryNewOwnerProposed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistryNewOwnerProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryNewOwnerProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryNewOwnerProposed represents a NewOwnerProposed event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryNewOwnerProposed struct {
	ExtensionId *big.Int
	OldOwner    common.Address
	NewOwner    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerProposed is a free log retrieval operation binding the contract event 0x4bc12216fc8d0204041572c77b2c83fafdb0c3fd38ac4de2edb26e8462cbb63a.
//
// Solidity: event NewOwnerProposed(uint256 indexed extensionId, address indexed oldOwner, address indexed newOwner)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterNewOwnerProposed(opts *bind.FilterOpts, extensionId []*big.Int, oldOwner []common.Address, newOwner []common.Address) (*TeeExtensionRegistryNewOwnerProposedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "NewOwnerProposed", extensionIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryNewOwnerProposedIterator{contract: _TeeExtensionRegistry.contract, event: "NewOwnerProposed", logs: logs, sub: sub}, nil
}

// WatchNewOwnerProposed is a free log subscription operation binding the contract event 0x4bc12216fc8d0204041572c77b2c83fafdb0c3fd38ac4de2edb26e8462cbb63a.
//
// Solidity: event NewOwnerProposed(uint256 indexed extensionId, address indexed oldOwner, address indexed newOwner)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchNewOwnerProposed(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryNewOwnerProposed, extensionId []*big.Int, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "NewOwnerProposed", extensionIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryNewOwnerProposed)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewOwnerProposed is a log parse operation binding the contract event 0x4bc12216fc8d0204041572c77b2c83fafdb0c3fd38ac4de2edb26e8462cbb63a.
//
// Solidity: event NewOwnerProposed(uint256 indexed extensionId, address indexed oldOwner, address indexed newOwner)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseNewOwnerProposed(log types.Log) (*TeeExtensionRegistryNewOwnerProposed, error) {
	event := new(TeeExtensionRegistryNewOwnerProposed)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistrySupportedPlatformAddedIterator is returned from FilterSupportedPlatformAdded and is used to iterate over the raw logs and unpacked data for SupportedPlatformAdded events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySupportedPlatformAddedIterator struct {
	Event *TeeExtensionRegistrySupportedPlatformAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistrySupportedPlatformAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistrySupportedPlatformAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistrySupportedPlatformAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistrySupportedPlatformAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistrySupportedPlatformAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistrySupportedPlatformAdded represents a SupportedPlatformAdded event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySupportedPlatformAdded struct {
	Platform [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSupportedPlatformAdded is a free log retrieval operation binding the contract event 0xf28352c70129dc8c465c9fd7d85bec07b1454dc6870af70a4ab37a0b7ab7f369.
//
// Solidity: event SupportedPlatformAdded(bytes32 indexed platform)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterSupportedPlatformAdded(opts *bind.FilterOpts, platform [][32]byte) (*TeeExtensionRegistrySupportedPlatformAddedIterator, error) {

	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "SupportedPlatformAdded", platformRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistrySupportedPlatformAddedIterator{contract: _TeeExtensionRegistry.contract, event: "SupportedPlatformAdded", logs: logs, sub: sub}, nil
}

// WatchSupportedPlatformAdded is a free log subscription operation binding the contract event 0xf28352c70129dc8c465c9fd7d85bec07b1454dc6870af70a4ab37a0b7ab7f369.
//
// Solidity: event SupportedPlatformAdded(bytes32 indexed platform)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchSupportedPlatformAdded(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistrySupportedPlatformAdded, platform [][32]byte) (event.Subscription, error) {

	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "SupportedPlatformAdded", platformRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistrySupportedPlatformAdded)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SupportedPlatformAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSupportedPlatformAdded is a log parse operation binding the contract event 0xf28352c70129dc8c465c9fd7d85bec07b1454dc6870af70a4ab37a0b7ab7f369.
//
// Solidity: event SupportedPlatformAdded(bytes32 indexed platform)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseSupportedPlatformAdded(log types.Log) (*TeeExtensionRegistrySupportedPlatformAdded, error) {
	event := new(TeeExtensionRegistrySupportedPlatformAdded)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SupportedPlatformAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistrySupportedWalletProjectOpTypeAddedIterator is returned from FilterSupportedWalletProjectOpTypeAdded and is used to iterate over the raw logs and unpacked data for SupportedWalletProjectOpTypeAdded events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySupportedWalletProjectOpTypeAddedIterator struct {
	Event *TeeExtensionRegistrySupportedWalletProjectOpTypeAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistrySupportedWalletProjectOpTypeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistrySupportedWalletProjectOpTypeAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistrySupportedWalletProjectOpTypeAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistrySupportedWalletProjectOpTypeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistrySupportedWalletProjectOpTypeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistrySupportedWalletProjectOpTypeAdded represents a SupportedWalletProjectOpTypeAdded event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySupportedWalletProjectOpTypeAdded struct {
	ExtensionId *big.Int
	OpType      [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSupportedWalletProjectOpTypeAdded is a free log retrieval operation binding the contract event 0x14f7cd286eb660340c3f76200081ca65cd27b0a2065c549e801c510c18a3db46.
//
// Solidity: event SupportedWalletProjectOpTypeAdded(uint256 indexed extensionId, bytes32 indexed opType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterSupportedWalletProjectOpTypeAdded(opts *bind.FilterOpts, extensionId []*big.Int, opType [][32]byte) (*TeeExtensionRegistrySupportedWalletProjectOpTypeAddedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var opTypeRule []interface{}
	for _, opTypeItem := range opType {
		opTypeRule = append(opTypeRule, opTypeItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "SupportedWalletProjectOpTypeAdded", extensionIdRule, opTypeRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistrySupportedWalletProjectOpTypeAddedIterator{contract: _TeeExtensionRegistry.contract, event: "SupportedWalletProjectOpTypeAdded", logs: logs, sub: sub}, nil
}

// WatchSupportedWalletProjectOpTypeAdded is a free log subscription operation binding the contract event 0x14f7cd286eb660340c3f76200081ca65cd27b0a2065c549e801c510c18a3db46.
//
// Solidity: event SupportedWalletProjectOpTypeAdded(uint256 indexed extensionId, bytes32 indexed opType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchSupportedWalletProjectOpTypeAdded(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistrySupportedWalletProjectOpTypeAdded, extensionId []*big.Int, opType [][32]byte) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var opTypeRule []interface{}
	for _, opTypeItem := range opType {
		opTypeRule = append(opTypeRule, opTypeItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "SupportedWalletProjectOpTypeAdded", extensionIdRule, opTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistrySupportedWalletProjectOpTypeAdded)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SupportedWalletProjectOpTypeAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSupportedWalletProjectOpTypeAdded is a log parse operation binding the contract event 0x14f7cd286eb660340c3f76200081ca65cd27b0a2065c549e801c510c18a3db46.
//
// Solidity: event SupportedWalletProjectOpTypeAdded(uint256 indexed extensionId, bytes32 indexed opType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseSupportedWalletProjectOpTypeAdded(log types.Log) (*TeeExtensionRegistrySupportedWalletProjectOpTypeAdded, error) {
	event := new(TeeExtensionRegistrySupportedWalletProjectOpTypeAdded)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SupportedWalletProjectOpTypeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistrySupportedWalletProjectOpTypeRemovedIterator is returned from FilterSupportedWalletProjectOpTypeRemoved and is used to iterate over the raw logs and unpacked data for SupportedWalletProjectOpTypeRemoved events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySupportedWalletProjectOpTypeRemovedIterator struct {
	Event *TeeExtensionRegistrySupportedWalletProjectOpTypeRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistrySupportedWalletProjectOpTypeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistrySupportedWalletProjectOpTypeRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistrySupportedWalletProjectOpTypeRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistrySupportedWalletProjectOpTypeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistrySupportedWalletProjectOpTypeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistrySupportedWalletProjectOpTypeRemoved represents a SupportedWalletProjectOpTypeRemoved event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySupportedWalletProjectOpTypeRemoved struct {
	ExtensionId *big.Int
	OpType      [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSupportedWalletProjectOpTypeRemoved is a free log retrieval operation binding the contract event 0x728dce34af3917dbaebf0c52bafb1fbab2e6a982d4880c7fe6b8c87af50971c4.
//
// Solidity: event SupportedWalletProjectOpTypeRemoved(uint256 indexed extensionId, bytes32 indexed opType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterSupportedWalletProjectOpTypeRemoved(opts *bind.FilterOpts, extensionId []*big.Int, opType [][32]byte) (*TeeExtensionRegistrySupportedWalletProjectOpTypeRemovedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var opTypeRule []interface{}
	for _, opTypeItem := range opType {
		opTypeRule = append(opTypeRule, opTypeItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "SupportedWalletProjectOpTypeRemoved", extensionIdRule, opTypeRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistrySupportedWalletProjectOpTypeRemovedIterator{contract: _TeeExtensionRegistry.contract, event: "SupportedWalletProjectOpTypeRemoved", logs: logs, sub: sub}, nil
}

// WatchSupportedWalletProjectOpTypeRemoved is a free log subscription operation binding the contract event 0x728dce34af3917dbaebf0c52bafb1fbab2e6a982d4880c7fe6b8c87af50971c4.
//
// Solidity: event SupportedWalletProjectOpTypeRemoved(uint256 indexed extensionId, bytes32 indexed opType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchSupportedWalletProjectOpTypeRemoved(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistrySupportedWalletProjectOpTypeRemoved, extensionId []*big.Int, opType [][32]byte) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var opTypeRule []interface{}
	for _, opTypeItem := range opType {
		opTypeRule = append(opTypeRule, opTypeItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "SupportedWalletProjectOpTypeRemoved", extensionIdRule, opTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistrySupportedWalletProjectOpTypeRemoved)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SupportedWalletProjectOpTypeRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSupportedWalletProjectOpTypeRemoved is a log parse operation binding the contract event 0x728dce34af3917dbaebf0c52bafb1fbab2e6a982d4880c7fe6b8c87af50971c4.
//
// Solidity: event SupportedWalletProjectOpTypeRemoved(uint256 indexed extensionId, bytes32 indexed opType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseSupportedWalletProjectOpTypeRemoved(log types.Log) (*TeeExtensionRegistrySupportedWalletProjectOpTypeRemoved, error) {
	event := new(TeeExtensionRegistrySupportedWalletProjectOpTypeRemoved)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SupportedWalletProjectOpTypeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryTeeExtensionContractsSetIterator is returned from FilterTeeExtensionContractsSet and is used to iterate over the raw logs and unpacked data for TeeExtensionContractsSet events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryTeeExtensionContractsSetIterator struct {
	Event *TeeExtensionRegistryTeeExtensionContractsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistryTeeExtensionContractsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryTeeExtensionContractsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistryTeeExtensionContractsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistryTeeExtensionContractsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryTeeExtensionContractsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryTeeExtensionContractsSet represents a TeeExtensionContractsSet event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryTeeExtensionContractsSet struct {
	ExtensionId                    *big.Int
	TeeExtensionStateVerifier      common.Address
	TeeExtensionInstructionsSender common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterTeeExtensionContractsSet is a free log retrieval operation binding the contract event 0xb288ce795dac7e8dff05f2e01dbac436986c930c5a57e27deeb206aad89526fa.
//
// Solidity: event TeeExtensionContractsSet(uint256 indexed extensionId, address indexed teeExtensionStateVerifier, address indexed teeExtensionInstructionsSender)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterTeeExtensionContractsSet(opts *bind.FilterOpts, extensionId []*big.Int, teeExtensionStateVerifier []common.Address, teeExtensionInstructionsSender []common.Address) (*TeeExtensionRegistryTeeExtensionContractsSetIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var teeExtensionStateVerifierRule []interface{}
	for _, teeExtensionStateVerifierItem := range teeExtensionStateVerifier {
		teeExtensionStateVerifierRule = append(teeExtensionStateVerifierRule, teeExtensionStateVerifierItem)
	}
	var teeExtensionInstructionsSenderRule []interface{}
	for _, teeExtensionInstructionsSenderItem := range teeExtensionInstructionsSender {
		teeExtensionInstructionsSenderRule = append(teeExtensionInstructionsSenderRule, teeExtensionInstructionsSenderItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "TeeExtensionContractsSet", extensionIdRule, teeExtensionStateVerifierRule, teeExtensionInstructionsSenderRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryTeeExtensionContractsSetIterator{contract: _TeeExtensionRegistry.contract, event: "TeeExtensionContractsSet", logs: logs, sub: sub}, nil
}

// WatchTeeExtensionContractsSet is a free log subscription operation binding the contract event 0xb288ce795dac7e8dff05f2e01dbac436986c930c5a57e27deeb206aad89526fa.
//
// Solidity: event TeeExtensionContractsSet(uint256 indexed extensionId, address indexed teeExtensionStateVerifier, address indexed teeExtensionInstructionsSender)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchTeeExtensionContractsSet(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryTeeExtensionContractsSet, extensionId []*big.Int, teeExtensionStateVerifier []common.Address, teeExtensionInstructionsSender []common.Address) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var teeExtensionStateVerifierRule []interface{}
	for _, teeExtensionStateVerifierItem := range teeExtensionStateVerifier {
		teeExtensionStateVerifierRule = append(teeExtensionStateVerifierRule, teeExtensionStateVerifierItem)
	}
	var teeExtensionInstructionsSenderRule []interface{}
	for _, teeExtensionInstructionsSenderItem := range teeExtensionInstructionsSender {
		teeExtensionInstructionsSenderRule = append(teeExtensionInstructionsSenderRule, teeExtensionInstructionsSenderItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "TeeExtensionContractsSet", extensionIdRule, teeExtensionStateVerifierRule, teeExtensionInstructionsSenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryTeeExtensionContractsSet)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "TeeExtensionContractsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeeExtensionContractsSet is a log parse operation binding the contract event 0xb288ce795dac7e8dff05f2e01dbac436986c930c5a57e27deeb206aad89526fa.
//
// Solidity: event TeeExtensionContractsSet(uint256 indexed extensionId, address indexed teeExtensionStateVerifier, address indexed teeExtensionInstructionsSender)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseTeeExtensionContractsSet(log types.Log) (*TeeExtensionRegistryTeeExtensionContractsSet, error) {
	event := new(TeeExtensionRegistryTeeExtensionContractsSet)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "TeeExtensionContractsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryTeeExtensionRegisteredIterator is returned from FilterTeeExtensionRegistered and is used to iterate over the raw logs and unpacked data for TeeExtensionRegistered events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryTeeExtensionRegisteredIterator struct {
	Event *TeeExtensionRegistryTeeExtensionRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistryTeeExtensionRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryTeeExtensionRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistryTeeExtensionRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistryTeeExtensionRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryTeeExtensionRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryTeeExtensionRegistered represents a TeeExtensionRegistered event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryTeeExtensionRegistered struct {
	ExtensionId *big.Int
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTeeExtensionRegistered is a free log retrieval operation binding the contract event 0x8b9cae23662229e1f17f8b1d3a6670c1000f5cd2c26b1dadb976c46620da9c21.
//
// Solidity: event TeeExtensionRegistered(uint256 indexed extensionId, address indexed owner)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterTeeExtensionRegistered(opts *bind.FilterOpts, extensionId []*big.Int, owner []common.Address) (*TeeExtensionRegistryTeeExtensionRegisteredIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "TeeExtensionRegistered", extensionIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryTeeExtensionRegisteredIterator{contract: _TeeExtensionRegistry.contract, event: "TeeExtensionRegistered", logs: logs, sub: sub}, nil
}

// WatchTeeExtensionRegistered is a free log subscription operation binding the contract event 0x8b9cae23662229e1f17f8b1d3a6670c1000f5cd2c26b1dadb976c46620da9c21.
//
// Solidity: event TeeExtensionRegistered(uint256 indexed extensionId, address indexed owner)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchTeeExtensionRegistered(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryTeeExtensionRegistered, extensionId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "TeeExtensionRegistered", extensionIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryTeeExtensionRegistered)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "TeeExtensionRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeeExtensionRegistered is a log parse operation binding the contract event 0x8b9cae23662229e1f17f8b1d3a6670c1000f5cd2c26b1dadb976c46620da9c21.
//
// Solidity: event TeeExtensionRegistered(uint256 indexed extensionId, address indexed owner)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseTeeExtensionRegistered(log types.Log) (*TeeExtensionRegistryTeeExtensionRegistered, error) {
	event := new(TeeExtensionRegistryTeeExtensionRegistered)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "TeeExtensionRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryTeeInstructionsSentIterator struct {
	Event *TeeExtensionRegistryTeeInstructionsSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistryTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryTeeInstructionsSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistryTeeInstructionsSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistryTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryTeeInstructionsSent represents a TeeInstructionsSent event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryTeeInstructionsSent struct {
	ExtensionId        *big.Int
	InstructionId      [32]byte
	RewardEpochId      uint32
	TeeMachines        []ITeeMachineRegistryTeeMachine
	OpType             [32]byte
	OpCommand          [32]byte
	Message            []byte
	Cosigners          []common.Address
	CosignersThreshold uint64
	Fee                *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTeeInstructionsSent is a free log retrieval operation binding the contract event 0xd2b490c6cf441de1940e58ec5d773c37109f3543213cd6992247896744d8c03b.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, uint256 fee)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (*TeeExtensionRegistryTeeInstructionsSentIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryTeeInstructionsSentIterator{contract: _TeeExtensionRegistry.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0xd2b490c6cf441de1940e58ec5d773c37109f3543213cd6992247896744d8c03b.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, uint256 fee)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryTeeInstructionsSent, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryTeeInstructionsSent)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeeInstructionsSent is a log parse operation binding the contract event 0xd2b490c6cf441de1940e58ec5d773c37109f3543213cd6992247896744d8c03b.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, uint256 fee)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseTeeInstructionsSent(log types.Log) (*TeeExtensionRegistryTeeInstructionsSent, error) {
	event := new(TeeExtensionRegistryTeeInstructionsSent)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryTeeVersionAddedIterator is returned from FilterTeeVersionAdded and is used to iterate over the raw logs and unpacked data for TeeVersionAdded events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryTeeVersionAddedIterator struct {
	Event *TeeExtensionRegistryTeeVersionAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistryTeeVersionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryTeeVersionAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistryTeeVersionAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistryTeeVersionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryTeeVersionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryTeeVersionAdded represents a TeeVersionAdded event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryTeeVersionAdded struct {
	ExtensionId    *big.Int
	CodeHash       [32]byte
	Version        string
	Platforms      [][32]byte
	GovernanceHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTeeVersionAdded is a free log retrieval operation binding the contract event 0x7579546f6a19926a903415c470d64c40353a972a647c72fcd7418572accfe88b.
//
// Solidity: event TeeVersionAdded(uint256 indexed extensionId, bytes32 indexed codeHash, string version, bytes32[] platforms, bytes32 governanceHash)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterTeeVersionAdded(opts *bind.FilterOpts, extensionId []*big.Int, codeHash [][32]byte) (*TeeExtensionRegistryTeeVersionAddedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var codeHashRule []interface{}
	for _, codeHashItem := range codeHash {
		codeHashRule = append(codeHashRule, codeHashItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "TeeVersionAdded", extensionIdRule, codeHashRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryTeeVersionAddedIterator{contract: _TeeExtensionRegistry.contract, event: "TeeVersionAdded", logs: logs, sub: sub}, nil
}

// WatchTeeVersionAdded is a free log subscription operation binding the contract event 0x7579546f6a19926a903415c470d64c40353a972a647c72fcd7418572accfe88b.
//
// Solidity: event TeeVersionAdded(uint256 indexed extensionId, bytes32 indexed codeHash, string version, bytes32[] platforms, bytes32 governanceHash)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchTeeVersionAdded(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryTeeVersionAdded, extensionId []*big.Int, codeHash [][32]byte) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var codeHashRule []interface{}
	for _, codeHashItem := range codeHash {
		codeHashRule = append(codeHashRule, codeHashItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "TeeVersionAdded", extensionIdRule, codeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryTeeVersionAdded)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "TeeVersionAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeeVersionAdded is a log parse operation binding the contract event 0x7579546f6a19926a903415c470d64c40353a972a647c72fcd7418572accfe88b.
//
// Solidity: event TeeVersionAdded(uint256 indexed extensionId, bytes32 indexed codeHash, string version, bytes32[] platforms, bytes32 governanceHash)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseTeeVersionAdded(log types.Log) (*TeeExtensionRegistryTeeVersionAdded, error) {
	event := new(TeeExtensionRegistryTeeVersionAdded)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "TeeVersionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryTimelockedGovernanceCallCanceledIterator struct {
	Event *TeeExtensionRegistryTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistryTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryTimelockedGovernanceCallCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistryTimelockedGovernanceCallCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistryTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeeExtensionRegistryTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryTimelockedGovernanceCallCanceledIterator{contract: _TeeExtensionRegistry.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryTimelockedGovernanceCallCanceled)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimelockedGovernanceCallCanceled is a log parse operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeeExtensionRegistryTimelockedGovernanceCallCanceled, error) {
	event := new(TeeExtensionRegistryTimelockedGovernanceCallCanceled)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryTimelockedGovernanceCallExecutedIterator struct {
	Event *TeeExtensionRegistryTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistryTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryTimelockedGovernanceCallExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistryTimelockedGovernanceCallExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistryTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeeExtensionRegistryTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryTimelockedGovernanceCallExecutedIterator{contract: _TeeExtensionRegistry.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryTimelockedGovernanceCallExecuted)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimelockedGovernanceCallExecuted is a log parse operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeeExtensionRegistryTimelockedGovernanceCallExecuted, error) {
	event := new(TeeExtensionRegistryTimelockedGovernanceCallExecuted)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryUpgradedIterator struct {
	Event *TeeExtensionRegistryUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeExtensionRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeExtensionRegistryUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeExtensionRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryUpgraded represents a Upgraded event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeeExtensionRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryUpgradedIterator{contract: _TeeExtensionRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryUpgraded)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseUpgraded(log types.Log) (*TeeExtensionRegistryUpgraded, error) {
	event := new(TeeExtensionRegistryUpgraded)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
