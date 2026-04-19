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

// TeeExtensionRegistryMetaData contains all meta data concerning the TeeExtensionRegistry contract.
var TeeExtensionRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CodeHashPlatformAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CodeHashZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCodeHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInstructionsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPlatform\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyTypeEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPlatforms\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"NoSigningAlgos\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"PlatformAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PlatformEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"}],\"name\":\"SigningAlgoAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SigningAlgoEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SystemOwnedExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"UnsupportedPlatform\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"CodeHashPlatformDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTypes\",\"type\":\"bytes32[]\"}],\"name\":\"SupportedKeyTypesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTypes\",\"type\":\"bytes32[]\"}],\"name\":\"SupportedKeyTypesRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTypes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[][]\",\"name\":\"_signingAlgosByKeyType\",\"type\":\"bytes32[][]\"}],\"name\":\"SystemSupportedKeyTypesAndSigningAlgosAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"platforms\",\"type\":\"bytes32[]\"}],\"name\":\"SystemSupportedPlatformsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"teeExtensionStateVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"TeeExtensionContractsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"TeeExtensionRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"platforms\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"governanceHash\",\"type\":\"bytes32\"}],\"name\":\"TeeVersionAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_keyTypes\",\"type\":\"bytes32[]\"}],\"name\":\"addSupportedKeyTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_keyTypes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_signingAlgosByKeyType\",\"type\":\"bytes32[][]\"}],\"name\":\"addSystemSupportedKeyTypesAndSigningAlgos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"}],\"name\":\"addSystemSupportedPlatforms\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"addTeeVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"disableCodeHashPlatform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extensionsCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"}],\"name\":\"getCodeHashInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getExtensionOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getSupportedCodeHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_supportedCodeHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getSupportedKeyTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_supportedKeyTypes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSystemSupportedKeyTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSystemSupportedPlatforms\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"}],\"name\":\"getSystemSupportedSigningAlgos\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getTeeExtensionInstructionsSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getTeeExtensionStateVerifier\",\"outputs\":[{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"}],\"name\":\"getTeeGovernanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"isCodeHashPlatformDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"isCodeHashPlatformSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"}],\"name\":\"isKeyTypeSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_signingAlgo\",\"type\":\"bytes32\"}],\"name\":\"isSigningAlgoSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"_teeExtensionStateVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_keyTypes\",\"type\":\"bytes32[]\"}],\"name\":\"removeSupportedKeyTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"_teeExtensionStateVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"setExtensionContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetSupportedCodeHashes is a free data retrieval call binding the contract method 0x9ebb72b2.
//
// Solidity: function getSupportedCodeHashes(uint256 _extensionId) view returns(bytes32[] _supportedCodeHashes)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetSupportedCodeHashes(opts *bind.CallOpts, _extensionId *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getSupportedCodeHashes", _extensionId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSupportedCodeHashes is a free data retrieval call binding the contract method 0x9ebb72b2.
//
// Solidity: function getSupportedCodeHashes(uint256 _extensionId) view returns(bytes32[] _supportedCodeHashes)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetSupportedCodeHashes(_extensionId *big.Int) ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSupportedCodeHashes(&_TeeExtensionRegistry.CallOpts, _extensionId)
}

// GetSupportedCodeHashes is a free data retrieval call binding the contract method 0x9ebb72b2.
//
// Solidity: function getSupportedCodeHashes(uint256 _extensionId) view returns(bytes32[] _supportedCodeHashes)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetSupportedCodeHashes(_extensionId *big.Int) ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSupportedCodeHashes(&_TeeExtensionRegistry.CallOpts, _extensionId)
}

// GetSupportedKeyTypes is a free data retrieval call binding the contract method 0x8bab3177.
//
// Solidity: function getSupportedKeyTypes(uint256 _extensionId) view returns(bytes32[] _supportedKeyTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetSupportedKeyTypes(opts *bind.CallOpts, _extensionId *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getSupportedKeyTypes", _extensionId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSupportedKeyTypes is a free data retrieval call binding the contract method 0x8bab3177.
//
// Solidity: function getSupportedKeyTypes(uint256 _extensionId) view returns(bytes32[] _supportedKeyTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetSupportedKeyTypes(_extensionId *big.Int) ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSupportedKeyTypes(&_TeeExtensionRegistry.CallOpts, _extensionId)
}

// GetSupportedKeyTypes is a free data retrieval call binding the contract method 0x8bab3177.
//
// Solidity: function getSupportedKeyTypes(uint256 _extensionId) view returns(bytes32[] _supportedKeyTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetSupportedKeyTypes(_extensionId *big.Int) ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSupportedKeyTypes(&_TeeExtensionRegistry.CallOpts, _extensionId)
}

// GetSystemSupportedKeyTypes is a free data retrieval call binding the contract method 0x50a15fcc.
//
// Solidity: function getSystemSupportedKeyTypes() view returns(bytes32[])
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetSystemSupportedKeyTypes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getSystemSupportedKeyTypes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSystemSupportedKeyTypes is a free data retrieval call binding the contract method 0x50a15fcc.
//
// Solidity: function getSystemSupportedKeyTypes() view returns(bytes32[])
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetSystemSupportedKeyTypes() ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSystemSupportedKeyTypes(&_TeeExtensionRegistry.CallOpts)
}

// GetSystemSupportedKeyTypes is a free data retrieval call binding the contract method 0x50a15fcc.
//
// Solidity: function getSystemSupportedKeyTypes() view returns(bytes32[])
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetSystemSupportedKeyTypes() ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSystemSupportedKeyTypes(&_TeeExtensionRegistry.CallOpts)
}

// GetSystemSupportedPlatforms is a free data retrieval call binding the contract method 0xe38c8964.
//
// Solidity: function getSystemSupportedPlatforms() view returns(bytes32[])
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetSystemSupportedPlatforms(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getSystemSupportedPlatforms")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSystemSupportedPlatforms is a free data retrieval call binding the contract method 0xe38c8964.
//
// Solidity: function getSystemSupportedPlatforms() view returns(bytes32[])
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetSystemSupportedPlatforms() ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSystemSupportedPlatforms(&_TeeExtensionRegistry.CallOpts)
}

// GetSystemSupportedPlatforms is a free data retrieval call binding the contract method 0xe38c8964.
//
// Solidity: function getSystemSupportedPlatforms() view returns(bytes32[])
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetSystemSupportedPlatforms() ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSystemSupportedPlatforms(&_TeeExtensionRegistry.CallOpts)
}

// GetSystemSupportedSigningAlgos is a free data retrieval call binding the contract method 0xc6cba800.
//
// Solidity: function getSystemSupportedSigningAlgos(bytes32 _keyType) view returns(bytes32[])
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetSystemSupportedSigningAlgos(opts *bind.CallOpts, _keyType [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getSystemSupportedSigningAlgos", _keyType)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSystemSupportedSigningAlgos is a free data retrieval call binding the contract method 0xc6cba800.
//
// Solidity: function getSystemSupportedSigningAlgos(bytes32 _keyType) view returns(bytes32[])
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetSystemSupportedSigningAlgos(_keyType [32]byte) ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSystemSupportedSigningAlgos(&_TeeExtensionRegistry.CallOpts, _keyType)
}

// GetSystemSupportedSigningAlgos is a free data retrieval call binding the contract method 0xc6cba800.
//
// Solidity: function getSystemSupportedSigningAlgos(bytes32 _keyType) view returns(bytes32[])
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetSystemSupportedSigningAlgos(_keyType [32]byte) ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSystemSupportedSigningAlgos(&_TeeExtensionRegistry.CallOpts, _keyType)
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

// IsCodeHashPlatformDisabled is a free data retrieval call binding the contract method 0x2a19c98e.
//
// Solidity: function isCodeHashPlatformDisabled(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) IsCodeHashPlatformDisabled(opts *bind.CallOpts, _extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "isCodeHashPlatformDisabled", _extensionId, _codeHash, _platform)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCodeHashPlatformDisabled is a free data retrieval call binding the contract method 0x2a19c98e.
//
// Solidity: function isCodeHashPlatformDisabled(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) IsCodeHashPlatformDisabled(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsCodeHashPlatformDisabled(&_TeeExtensionRegistry.CallOpts, _extensionId, _codeHash, _platform)
}

// IsCodeHashPlatformDisabled is a free data retrieval call binding the contract method 0x2a19c98e.
//
// Solidity: function isCodeHashPlatformDisabled(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) IsCodeHashPlatformDisabled(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsCodeHashPlatformDisabled(&_TeeExtensionRegistry.CallOpts, _extensionId, _codeHash, _platform)
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

// IsKeyTypeSupported is a free data retrieval call binding the contract method 0xf4ac4279.
//
// Solidity: function isKeyTypeSupported(uint256 _extensionId, bytes32 _keyType) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) IsKeyTypeSupported(opts *bind.CallOpts, _extensionId *big.Int, _keyType [32]byte) (bool, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "isKeyTypeSupported", _extensionId, _keyType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKeyTypeSupported is a free data retrieval call binding the contract method 0xf4ac4279.
//
// Solidity: function isKeyTypeSupported(uint256 _extensionId, bytes32 _keyType) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) IsKeyTypeSupported(_extensionId *big.Int, _keyType [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsKeyTypeSupported(&_TeeExtensionRegistry.CallOpts, _extensionId, _keyType)
}

// IsKeyTypeSupported is a free data retrieval call binding the contract method 0xf4ac4279.
//
// Solidity: function isKeyTypeSupported(uint256 _extensionId, bytes32 _keyType) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) IsKeyTypeSupported(_extensionId *big.Int, _keyType [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsKeyTypeSupported(&_TeeExtensionRegistry.CallOpts, _extensionId, _keyType)
}

// IsSigningAlgoSupported is a free data retrieval call binding the contract method 0x62487d98.
//
// Solidity: function isSigningAlgoSupported(bytes32 _keyType, bytes32 _signingAlgo) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) IsSigningAlgoSupported(opts *bind.CallOpts, _keyType [32]byte, _signingAlgo [32]byte) (bool, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "isSigningAlgoSupported", _keyType, _signingAlgo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigningAlgoSupported is a free data retrieval call binding the contract method 0x62487d98.
//
// Solidity: function isSigningAlgoSupported(bytes32 _keyType, bytes32 _signingAlgo) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) IsSigningAlgoSupported(_keyType [32]byte, _signingAlgo [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsSigningAlgoSupported(&_TeeExtensionRegistry.CallOpts, _keyType, _signingAlgo)
}

// IsSigningAlgoSupported is a free data retrieval call binding the contract method 0x62487d98.
//
// Solidity: function isSigningAlgoSupported(bytes32 _keyType, bytes32 _signingAlgo) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) IsSigningAlgoSupported(_keyType [32]byte, _signingAlgo [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsSigningAlgoSupported(&_TeeExtensionRegistry.CallOpts, _keyType, _signingAlgo)
}

// AddSupportedKeyTypes is a paid mutator transaction binding the contract method 0x174c5e26.
//
// Solidity: function addSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) AddSupportedKeyTypes(opts *bind.TransactOpts, _extensionId *big.Int, _keyTypes [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "addSupportedKeyTypes", _extensionId, _keyTypes)
}

// AddSupportedKeyTypes is a paid mutator transaction binding the contract method 0x174c5e26.
//
// Solidity: function addSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) AddSupportedKeyTypes(_extensionId *big.Int, _keyTypes [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.AddSupportedKeyTypes(&_TeeExtensionRegistry.TransactOpts, _extensionId, _keyTypes)
}

// AddSupportedKeyTypes is a paid mutator transaction binding the contract method 0x174c5e26.
//
// Solidity: function addSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) AddSupportedKeyTypes(_extensionId *big.Int, _keyTypes [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.AddSupportedKeyTypes(&_TeeExtensionRegistry.TransactOpts, _extensionId, _keyTypes)
}

// AddSystemSupportedKeyTypesAndSigningAlgos is a paid mutator transaction binding the contract method 0x38c77e46.
//
// Solidity: function addSystemSupportedKeyTypesAndSigningAlgos(bytes32[] _keyTypes, bytes32[][] _signingAlgosByKeyType) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) AddSystemSupportedKeyTypesAndSigningAlgos(opts *bind.TransactOpts, _keyTypes [][32]byte, _signingAlgosByKeyType [][][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "addSystemSupportedKeyTypesAndSigningAlgos", _keyTypes, _signingAlgosByKeyType)
}

// AddSystemSupportedKeyTypesAndSigningAlgos is a paid mutator transaction binding the contract method 0x38c77e46.
//
// Solidity: function addSystemSupportedKeyTypesAndSigningAlgos(bytes32[] _keyTypes, bytes32[][] _signingAlgosByKeyType) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) AddSystemSupportedKeyTypesAndSigningAlgos(_keyTypes [][32]byte, _signingAlgosByKeyType [][][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.AddSystemSupportedKeyTypesAndSigningAlgos(&_TeeExtensionRegistry.TransactOpts, _keyTypes, _signingAlgosByKeyType)
}

// AddSystemSupportedKeyTypesAndSigningAlgos is a paid mutator transaction binding the contract method 0x38c77e46.
//
// Solidity: function addSystemSupportedKeyTypesAndSigningAlgos(bytes32[] _keyTypes, bytes32[][] _signingAlgosByKeyType) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) AddSystemSupportedKeyTypesAndSigningAlgos(_keyTypes [][32]byte, _signingAlgosByKeyType [][][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.AddSystemSupportedKeyTypesAndSigningAlgos(&_TeeExtensionRegistry.TransactOpts, _keyTypes, _signingAlgosByKeyType)
}

// AddSystemSupportedPlatforms is a paid mutator transaction binding the contract method 0x0c14a66b.
//
// Solidity: function addSystemSupportedPlatforms(bytes32[] _platforms) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) AddSystemSupportedPlatforms(opts *bind.TransactOpts, _platforms [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "addSystemSupportedPlatforms", _platforms)
}

// AddSystemSupportedPlatforms is a paid mutator transaction binding the contract method 0x0c14a66b.
//
// Solidity: function addSystemSupportedPlatforms(bytes32[] _platforms) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) AddSystemSupportedPlatforms(_platforms [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.AddSystemSupportedPlatforms(&_TeeExtensionRegistry.TransactOpts, _platforms)
}

// AddSystemSupportedPlatforms is a paid mutator transaction binding the contract method 0x0c14a66b.
//
// Solidity: function addSystemSupportedPlatforms(bytes32[] _platforms) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) AddSystemSupportedPlatforms(_platforms [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.AddSystemSupportedPlatforms(&_TeeExtensionRegistry.TransactOpts, _platforms)
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

// RemoveSupportedKeyTypes is a paid mutator transaction binding the contract method 0xecb0233f.
//
// Solidity: function removeSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) RemoveSupportedKeyTypes(opts *bind.TransactOpts, _extensionId *big.Int, _keyTypes [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "removeSupportedKeyTypes", _extensionId, _keyTypes)
}

// RemoveSupportedKeyTypes is a paid mutator transaction binding the contract method 0xecb0233f.
//
// Solidity: function removeSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) RemoveSupportedKeyTypes(_extensionId *big.Int, _keyTypes [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.RemoveSupportedKeyTypes(&_TeeExtensionRegistry.TransactOpts, _extensionId, _keyTypes)
}

// RemoveSupportedKeyTypes is a paid mutator transaction binding the contract method 0xecb0233f.
//
// Solidity: function removeSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) RemoveSupportedKeyTypes(_extensionId *big.Int, _keyTypes [][32]byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.RemoveSupportedKeyTypes(&_TeeExtensionRegistry.TransactOpts, _extensionId, _keyTypes)
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
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeExtensionRegistryGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryGovernanceCallTimelockedIterator{contract: _TeeExtensionRegistry.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
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

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
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

// TeeExtensionRegistrySupportedKeyTypesAddedIterator is returned from FilterSupportedKeyTypesAdded and is used to iterate over the raw logs and unpacked data for SupportedKeyTypesAdded events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySupportedKeyTypesAddedIterator struct {
	Event *TeeExtensionRegistrySupportedKeyTypesAdded // Event containing the contract specifics and raw log

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
func (it *TeeExtensionRegistrySupportedKeyTypesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistrySupportedKeyTypesAdded)
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
		it.Event = new(TeeExtensionRegistrySupportedKeyTypesAdded)
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
func (it *TeeExtensionRegistrySupportedKeyTypesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistrySupportedKeyTypesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistrySupportedKeyTypesAdded represents a SupportedKeyTypesAdded event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySupportedKeyTypesAdded struct {
	ExtensionId *big.Int
	KeyTypes    [][32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSupportedKeyTypesAdded is a free log retrieval operation binding the contract event 0x77e5c9aa84ad0acef154e8ec75a77ac7031308cc906e405a1de6d51e104b596f.
//
// Solidity: event SupportedKeyTypesAdded(uint256 indexed extensionId, bytes32[] keyTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterSupportedKeyTypesAdded(opts *bind.FilterOpts, extensionId []*big.Int) (*TeeExtensionRegistrySupportedKeyTypesAddedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "SupportedKeyTypesAdded", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistrySupportedKeyTypesAddedIterator{contract: _TeeExtensionRegistry.contract, event: "SupportedKeyTypesAdded", logs: logs, sub: sub}, nil
}

// WatchSupportedKeyTypesAdded is a free log subscription operation binding the contract event 0x77e5c9aa84ad0acef154e8ec75a77ac7031308cc906e405a1de6d51e104b596f.
//
// Solidity: event SupportedKeyTypesAdded(uint256 indexed extensionId, bytes32[] keyTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchSupportedKeyTypesAdded(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistrySupportedKeyTypesAdded, extensionId []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "SupportedKeyTypesAdded", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistrySupportedKeyTypesAdded)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SupportedKeyTypesAdded", log); err != nil {
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

// ParseSupportedKeyTypesAdded is a log parse operation binding the contract event 0x77e5c9aa84ad0acef154e8ec75a77ac7031308cc906e405a1de6d51e104b596f.
//
// Solidity: event SupportedKeyTypesAdded(uint256 indexed extensionId, bytes32[] keyTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseSupportedKeyTypesAdded(log types.Log) (*TeeExtensionRegistrySupportedKeyTypesAdded, error) {
	event := new(TeeExtensionRegistrySupportedKeyTypesAdded)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SupportedKeyTypesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistrySupportedKeyTypesRemovedIterator is returned from FilterSupportedKeyTypesRemoved and is used to iterate over the raw logs and unpacked data for SupportedKeyTypesRemoved events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySupportedKeyTypesRemovedIterator struct {
	Event *TeeExtensionRegistrySupportedKeyTypesRemoved // Event containing the contract specifics and raw log

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
func (it *TeeExtensionRegistrySupportedKeyTypesRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistrySupportedKeyTypesRemoved)
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
		it.Event = new(TeeExtensionRegistrySupportedKeyTypesRemoved)
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
func (it *TeeExtensionRegistrySupportedKeyTypesRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistrySupportedKeyTypesRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistrySupportedKeyTypesRemoved represents a SupportedKeyTypesRemoved event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySupportedKeyTypesRemoved struct {
	ExtensionId *big.Int
	KeyTypes    [][32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSupportedKeyTypesRemoved is a free log retrieval operation binding the contract event 0x6d91cb2b5bca8a5b196fb4f70a87602c7ed5e0cfa271cf713959b8980b5591b7.
//
// Solidity: event SupportedKeyTypesRemoved(uint256 indexed extensionId, bytes32[] keyTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterSupportedKeyTypesRemoved(opts *bind.FilterOpts, extensionId []*big.Int) (*TeeExtensionRegistrySupportedKeyTypesRemovedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "SupportedKeyTypesRemoved", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistrySupportedKeyTypesRemovedIterator{contract: _TeeExtensionRegistry.contract, event: "SupportedKeyTypesRemoved", logs: logs, sub: sub}, nil
}

// WatchSupportedKeyTypesRemoved is a free log subscription operation binding the contract event 0x6d91cb2b5bca8a5b196fb4f70a87602c7ed5e0cfa271cf713959b8980b5591b7.
//
// Solidity: event SupportedKeyTypesRemoved(uint256 indexed extensionId, bytes32[] keyTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchSupportedKeyTypesRemoved(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistrySupportedKeyTypesRemoved, extensionId []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "SupportedKeyTypesRemoved", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistrySupportedKeyTypesRemoved)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SupportedKeyTypesRemoved", log); err != nil {
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

// ParseSupportedKeyTypesRemoved is a log parse operation binding the contract event 0x6d91cb2b5bca8a5b196fb4f70a87602c7ed5e0cfa271cf713959b8980b5591b7.
//
// Solidity: event SupportedKeyTypesRemoved(uint256 indexed extensionId, bytes32[] keyTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseSupportedKeyTypesRemoved(log types.Log) (*TeeExtensionRegistrySupportedKeyTypesRemoved, error) {
	event := new(TeeExtensionRegistrySupportedKeyTypesRemoved)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SupportedKeyTypesRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAddedIterator is returned from FilterSystemSupportedKeyTypesAndSigningAlgosAdded and is used to iterate over the raw logs and unpacked data for SystemSupportedKeyTypesAndSigningAlgosAdded events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAddedIterator struct {
	Event *TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAdded // Event containing the contract specifics and raw log

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
func (it *TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAdded)
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
		it.Event = new(TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAdded)
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
func (it *TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAdded represents a SystemSupportedKeyTypesAndSigningAlgosAdded event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAdded struct {
	KeyTypes              [][32]byte
	SigningAlgosByKeyType [][][32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSystemSupportedKeyTypesAndSigningAlgosAdded is a free log retrieval operation binding the contract event 0xda8ce9e8e45bfb16c746599b6ad42edeef963bcf2fe5d4c84438154e0251ab37.
//
// Solidity: event SystemSupportedKeyTypesAndSigningAlgosAdded(bytes32[] keyTypes, bytes32[][] _signingAlgosByKeyType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterSystemSupportedKeyTypesAndSigningAlgosAdded(opts *bind.FilterOpts) (*TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAddedIterator, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "SystemSupportedKeyTypesAndSigningAlgosAdded")
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAddedIterator{contract: _TeeExtensionRegistry.contract, event: "SystemSupportedKeyTypesAndSigningAlgosAdded", logs: logs, sub: sub}, nil
}

// WatchSystemSupportedKeyTypesAndSigningAlgosAdded is a free log subscription operation binding the contract event 0xda8ce9e8e45bfb16c746599b6ad42edeef963bcf2fe5d4c84438154e0251ab37.
//
// Solidity: event SystemSupportedKeyTypesAndSigningAlgosAdded(bytes32[] keyTypes, bytes32[][] _signingAlgosByKeyType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchSystemSupportedKeyTypesAndSigningAlgosAdded(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAdded) (event.Subscription, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "SystemSupportedKeyTypesAndSigningAlgosAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAdded)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SystemSupportedKeyTypesAndSigningAlgosAdded", log); err != nil {
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

// ParseSystemSupportedKeyTypesAndSigningAlgosAdded is a log parse operation binding the contract event 0xda8ce9e8e45bfb16c746599b6ad42edeef963bcf2fe5d4c84438154e0251ab37.
//
// Solidity: event SystemSupportedKeyTypesAndSigningAlgosAdded(bytes32[] keyTypes, bytes32[][] _signingAlgosByKeyType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseSystemSupportedKeyTypesAndSigningAlgosAdded(log types.Log) (*TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAdded, error) {
	event := new(TeeExtensionRegistrySystemSupportedKeyTypesAndSigningAlgosAdded)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SystemSupportedKeyTypesAndSigningAlgosAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistrySystemSupportedPlatformsAddedIterator is returned from FilterSystemSupportedPlatformsAdded and is used to iterate over the raw logs and unpacked data for SystemSupportedPlatformsAdded events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySystemSupportedPlatformsAddedIterator struct {
	Event *TeeExtensionRegistrySystemSupportedPlatformsAdded // Event containing the contract specifics and raw log

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
func (it *TeeExtensionRegistrySystemSupportedPlatformsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistrySystemSupportedPlatformsAdded)
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
		it.Event = new(TeeExtensionRegistrySystemSupportedPlatformsAdded)
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
func (it *TeeExtensionRegistrySystemSupportedPlatformsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistrySystemSupportedPlatformsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistrySystemSupportedPlatformsAdded represents a SystemSupportedPlatformsAdded event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistrySystemSupportedPlatformsAdded struct {
	Platforms [][32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSystemSupportedPlatformsAdded is a free log retrieval operation binding the contract event 0xef6456053b4f920a8c95c4d306ef8f280b497bc5c6950ffb8d16f732b7474c61.
//
// Solidity: event SystemSupportedPlatformsAdded(bytes32[] platforms)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterSystemSupportedPlatformsAdded(opts *bind.FilterOpts) (*TeeExtensionRegistrySystemSupportedPlatformsAddedIterator, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "SystemSupportedPlatformsAdded")
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistrySystemSupportedPlatformsAddedIterator{contract: _TeeExtensionRegistry.contract, event: "SystemSupportedPlatformsAdded", logs: logs, sub: sub}, nil
}

// WatchSystemSupportedPlatformsAdded is a free log subscription operation binding the contract event 0xef6456053b4f920a8c95c4d306ef8f280b497bc5c6950ffb8d16f732b7474c61.
//
// Solidity: event SystemSupportedPlatformsAdded(bytes32[] platforms)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchSystemSupportedPlatformsAdded(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistrySystemSupportedPlatformsAdded) (event.Subscription, error) {

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "SystemSupportedPlatformsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistrySystemSupportedPlatformsAdded)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SystemSupportedPlatformsAdded", log); err != nil {
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

// ParseSystemSupportedPlatformsAdded is a log parse operation binding the contract event 0xef6456053b4f920a8c95c4d306ef8f280b497bc5c6950ffb8d16f732b7474c61.
//
// Solidity: event SystemSupportedPlatformsAdded(bytes32[] platforms)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseSystemSupportedPlatformsAdded(log types.Log) (*TeeExtensionRegistrySystemSupportedPlatformsAdded, error) {
	event := new(TeeExtensionRegistrySystemSupportedPlatformsAdded)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "SystemSupportedPlatformsAdded", log); err != nil {
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
	Version        string
	CodeHash       [32]byte
	Platforms      [][32]byte
	GovernanceHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTeeVersionAdded is a free log retrieval operation binding the contract event 0x22e0033da6f945065ea0809d45dd72ddcd4cee1f5e8296c9bde27e59d99490b8.
//
// Solidity: event TeeVersionAdded(uint256 indexed extensionId, string version, bytes32 indexed codeHash, bytes32[] platforms, bytes32 governanceHash)
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

// WatchTeeVersionAdded is a free log subscription operation binding the contract event 0x22e0033da6f945065ea0809d45dd72ddcd4cee1f5e8296c9bde27e59d99490b8.
//
// Solidity: event TeeVersionAdded(uint256 indexed extensionId, string version, bytes32 indexed codeHash, bytes32[] platforms, bytes32 governanceHash)
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

// ParseTeeVersionAdded is a log parse operation binding the contract event 0x22e0033da6f945065ea0809d45dd72ddcd4cee1f5e8296c9bde27e59d99490b8.
//
// Solidity: event TeeVersionAdded(uint256 indexed extensionId, string version, bytes32 indexed codeHash, bytes32[] platforms, bytes32 governanceHash)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseTeeVersionAdded(log types.Log) (*TeeExtensionRegistryTeeVersionAdded, error) {
	event := new(TeeExtensionRegistryTeeVersionAdded)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "TeeVersionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
