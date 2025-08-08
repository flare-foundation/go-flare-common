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
	ABI: "[{\"inputs\":[],\"name\":\"CodeHashZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InstructionIdEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCodeHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInstructionsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPlatform\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPlatforms\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyInstructionsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OpTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeConstantsProviderNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"PlatformAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PlatformEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"name\":\"SystemOpTypeNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SystemOwnedExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"UnsupportedPlatform\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionEmpty\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"CodeHashPlatformDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"SupportedPlatformAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"name\":\"SupportedWalletProjectOpTypeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"name\":\"SupportedWalletProjectOpTypeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"teeExtensionStateVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"TeeExtensionContractsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"TeeExtensionRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structITeeMachineRegistry.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"platforms\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"governanceHash\",\"type\":\"bytes32\"}],\"name\":\"TeeVersionAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"contractITeeWalletProjectOpTypeConstants[]\",\"name\":\"_opTypeConstantsProviders\",\"type\":\"address[]\"}],\"name\":\"addOrUpdateSupportedWalletProjectOpTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"addTeeVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"codeHashPlatformDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"disableCodeHashPlatform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extensionsCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"}],\"name\":\"getCodeHashInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getExtensionOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getSupportedWalletProjectOpTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_supportedOpTypes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSystemInstructionInitiators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getTeeExtensionInstructionsSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getTeeExtensionStateVerifier\",\"outputs\":[{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"}],\"name\":\"getTeeGovernanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"}],\"name\":\"getWalletProjectOpTypeConstantsProvider\",\"outputs\":[{\"internalType\":\"contractITeeWalletProjectOpTypeConstants\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"isCodeHashPlatformSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"}],\"name\":\"isWalletProjectOpTypeSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"_teeExtensionStateVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_opTypes\",\"type\":\"bytes32[]\"}],\"name\":\"removeSupportedWalletProjectOpTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendInstructions\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"_teeExtensionStateVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"setExtensionContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
// Solidity: function getTeeGovernanceHash(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32 _governanceHash)
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
// Solidity: function getTeeGovernanceHash(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32 _governanceHash)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetTeeGovernanceHash(_extensionId *big.Int, _codeHash [32]byte) ([32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetTeeGovernanceHash(&_TeeExtensionRegistry.CallOpts, _extensionId, _codeHash)
}

// GetTeeGovernanceHash is a free data retrieval call binding the contract method 0x9f300b9b.
//
// Solidity: function getTeeGovernanceHash(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32 _governanceHash)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetTeeGovernanceHash(_extensionId *big.Int, _codeHash [32]byte) ([32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetTeeGovernanceHash(&_TeeExtensionRegistry.CallOpts, _extensionId, _codeHash)
}

// GetWalletProjectOpTypeConstantsProvider is a free data retrieval call binding the contract method 0xca1c9601.
//
// Solidity: function getWalletProjectOpTypeConstantsProvider(uint256 _extensionId, bytes32 _opType) view returns(address)
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
// Solidity: function getWalletProjectOpTypeConstantsProvider(uint256 _extensionId, bytes32 _opType) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetWalletProjectOpTypeConstantsProvider(_extensionId *big.Int, _opType [32]byte) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetWalletProjectOpTypeConstantsProvider(&_TeeExtensionRegistry.CallOpts, _extensionId, _opType)
}

// GetWalletProjectOpTypeConstantsProvider is a free data retrieval call binding the contract method 0xca1c9601.
//
// Solidity: function getWalletProjectOpTypeConstantsProvider(uint256 _extensionId, bytes32 _opType) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetWalletProjectOpTypeConstantsProvider(_extensionId *big.Int, _opType [32]byte) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetWalletProjectOpTypeConstantsProvider(&_TeeExtensionRegistry.CallOpts, _extensionId, _opType)
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
// Solidity: function register(address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) Register(opts *bind.TransactOpts, _teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "register", _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) Register(_teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.Register(&_TeeExtensionRegistry.TransactOpts, _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) Register(_teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.Register(&_TeeExtensionRegistry.TransactOpts, _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
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

// SendInstructions is a paid mutator transaction binding the contract method 0x86979aae.
//
// Solidity: function sendInstructions(bytes32 _instructionId, address[] _teeIds, bytes32 _opType, bytes32 _opCommand, bytes _message) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) SendInstructions(opts *bind.TransactOpts, _instructionId [32]byte, _teeIds []common.Address, _opType [32]byte, _opCommand [32]byte, _message []byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "sendInstructions", _instructionId, _teeIds, _opType, _opCommand, _message)
}

// SendInstructions is a paid mutator transaction binding the contract method 0x86979aae.
//
// Solidity: function sendInstructions(bytes32 _instructionId, address[] _teeIds, bytes32 _opType, bytes32 _opCommand, bytes _message) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) SendInstructions(_instructionId [32]byte, _teeIds []common.Address, _opType [32]byte, _opCommand [32]byte, _message []byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.SendInstructions(&_TeeExtensionRegistry.TransactOpts, _instructionId, _teeIds, _opType, _opCommand, _message)
}

// SendInstructions is a paid mutator transaction binding the contract method 0x86979aae.
//
// Solidity: function sendInstructions(bytes32 _instructionId, address[] _teeIds, bytes32 _opType, bytes32 _opCommand, bytes _message) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) SendInstructions(_instructionId [32]byte, _teeIds []common.Address, _opType [32]byte, _opCommand [32]byte, _message []byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.SendInstructions(&_TeeExtensionRegistry.TransactOpts, _instructionId, _teeIds, _opType, _opCommand, _message)
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
	InstructionId [32]byte
	RewardEpochId uint32
	TeeMachines   []ITeeMachineRegistryTeeMachine
	OpType        [32]byte
	OpCommand     [32]byte
	Message       []byte
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTeeInstructionsSent is a free log retrieval operation binding the contract event 0x0b155d08a134ef49412293aa17be27ead7cc04b35ada16085a36c4355fd6f40d.
//
// Solidity: event TeeInstructionsSent(bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, uint256 fee)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, instructionId [][32]byte, rewardEpochId []uint32) (*TeeExtensionRegistryTeeInstructionsSentIterator, error) {

	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "TeeInstructionsSent", instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryTeeInstructionsSentIterator{contract: _TeeExtensionRegistry.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0x0b155d08a134ef49412293aa17be27ead7cc04b35ada16085a36c4355fd6f40d.
//
// Solidity: event TeeInstructionsSent(bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, uint256 fee)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryTeeInstructionsSent, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "TeeInstructionsSent", instructionIdRule, rewardEpochIdRule)
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

// ParseTeeInstructionsSent is a log parse operation binding the contract event 0x0b155d08a134ef49412293aa17be27ead7cc04b35ada16085a36c4355fd6f40d.
//
// Solidity: event TeeInstructionsSent(bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, uint256 fee)
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
