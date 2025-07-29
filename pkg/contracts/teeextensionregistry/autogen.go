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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"CodeHashPlatformDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"teeExtensionStateVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"ExtensionContractsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ExtensionRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"name\":\"OpTypeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"name\":\"OpTypeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"PlatformAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structITeeMachineRegistry.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"platforms\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"governanceHash\",\"type\":\"bytes32\"}],\"name\":\"TeeVersionAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"codeHashPlatformDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"}],\"name\":\"getCodeHashInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getExtensionOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"}],\"name\":\"getOpTypeConstantsProvider\",\"outputs\":[{\"internalType\":\"contractITeeWalletProjectOpTypeConstants\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getSupportedOpTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_supportedOpTypes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getTeeExtensionInstructionsSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getTeeExtensionStateVerifier\",\"outputs\":[{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"}],\"name\":\"getTeeGovernanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"isCodeHashPlatformSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"}],\"name\":\"isOpTypeSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"_teeExtensionStateVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendInstructions\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"_teeExtensionStateVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"setExtensionContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetOpTypeConstantsProvider is a free data retrieval call binding the contract method 0x8ad751ab.
//
// Solidity: function getOpTypeConstantsProvider(uint256 _extensionId, bytes32 _opType) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetOpTypeConstantsProvider(opts *bind.CallOpts, _extensionId *big.Int, _opType [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getOpTypeConstantsProvider", _extensionId, _opType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOpTypeConstantsProvider is a free data retrieval call binding the contract method 0x8ad751ab.
//
// Solidity: function getOpTypeConstantsProvider(uint256 _extensionId, bytes32 _opType) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetOpTypeConstantsProvider(_extensionId *big.Int, _opType [32]byte) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetOpTypeConstantsProvider(&_TeeExtensionRegistry.CallOpts, _extensionId, _opType)
}

// GetOpTypeConstantsProvider is a free data retrieval call binding the contract method 0x8ad751ab.
//
// Solidity: function getOpTypeConstantsProvider(uint256 _extensionId, bytes32 _opType) view returns(address)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetOpTypeConstantsProvider(_extensionId *big.Int, _opType [32]byte) (common.Address, error) {
	return _TeeExtensionRegistry.Contract.GetOpTypeConstantsProvider(&_TeeExtensionRegistry.CallOpts, _extensionId, _opType)
}

// GetSupportedOpTypes is a free data retrieval call binding the contract method 0x2cf24f9e.
//
// Solidity: function getSupportedOpTypes(uint256 _extensionId) view returns(bytes32[] _supportedOpTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) GetSupportedOpTypes(opts *bind.CallOpts, _extensionId *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "getSupportedOpTypes", _extensionId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSupportedOpTypes is a free data retrieval call binding the contract method 0x2cf24f9e.
//
// Solidity: function getSupportedOpTypes(uint256 _extensionId) view returns(bytes32[] _supportedOpTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) GetSupportedOpTypes(_extensionId *big.Int) ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSupportedOpTypes(&_TeeExtensionRegistry.CallOpts, _extensionId)
}

// GetSupportedOpTypes is a free data retrieval call binding the contract method 0x2cf24f9e.
//
// Solidity: function getSupportedOpTypes(uint256 _extensionId) view returns(bytes32[] _supportedOpTypes)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) GetSupportedOpTypes(_extensionId *big.Int) ([][32]byte, error) {
	return _TeeExtensionRegistry.Contract.GetSupportedOpTypes(&_TeeExtensionRegistry.CallOpts, _extensionId)
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

// IsOpTypeSupported is a free data retrieval call binding the contract method 0xb4474d3e.
//
// Solidity: function isOpTypeSupported(uint256 _extensionId, bytes32 _opType) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCaller) IsOpTypeSupported(opts *bind.CallOpts, _extensionId *big.Int, _opType [32]byte) (bool, error) {
	var out []interface{}
	err := _TeeExtensionRegistry.contract.Call(opts, &out, "isOpTypeSupported", _extensionId, _opType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpTypeSupported is a free data retrieval call binding the contract method 0xb4474d3e.
//
// Solidity: function isOpTypeSupported(uint256 _extensionId, bytes32 _opType) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) IsOpTypeSupported(_extensionId *big.Int, _opType [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsOpTypeSupported(&_TeeExtensionRegistry.CallOpts, _extensionId, _opType)
}

// IsOpTypeSupported is a free data retrieval call binding the contract method 0xb4474d3e.
//
// Solidity: function isOpTypeSupported(uint256 _extensionId, bytes32 _opType) view returns(bool)
func (_TeeExtensionRegistry *TeeExtensionRegistryCallerSession) IsOpTypeSupported(_extensionId *big.Int, _opType [32]byte) (bool, error) {
	return _TeeExtensionRegistry.Contract.IsOpTypeSupported(&_TeeExtensionRegistry.CallOpts, _extensionId, _opType)
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

// SendInstructions is a paid mutator transaction binding the contract method 0x2bf49fc1.
//
// Solidity: function sendInstructions(bytes32 _instructionId, uint256 _extensionId, address[] _teeIds, bytes32 _opType, bytes32 _opCommand, bytes _message) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactor) SendInstructions(opts *bind.TransactOpts, _instructionId [32]byte, _extensionId *big.Int, _teeIds []common.Address, _opType [32]byte, _opCommand [32]byte, _message []byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.contract.Transact(opts, "sendInstructions", _instructionId, _extensionId, _teeIds, _opType, _opCommand, _message)
}

// SendInstructions is a paid mutator transaction binding the contract method 0x2bf49fc1.
//
// Solidity: function sendInstructions(bytes32 _instructionId, uint256 _extensionId, address[] _teeIds, bytes32 _opType, bytes32 _opCommand, bytes _message) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistrySession) SendInstructions(_instructionId [32]byte, _extensionId *big.Int, _teeIds []common.Address, _opType [32]byte, _opCommand [32]byte, _message []byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.SendInstructions(&_TeeExtensionRegistry.TransactOpts, _instructionId, _extensionId, _teeIds, _opType, _opCommand, _message)
}

// SendInstructions is a paid mutator transaction binding the contract method 0x2bf49fc1.
//
// Solidity: function sendInstructions(bytes32 _instructionId, uint256 _extensionId, address[] _teeIds, bytes32 _opType, bytes32 _opCommand, bytes _message) payable returns()
func (_TeeExtensionRegistry *TeeExtensionRegistryTransactorSession) SendInstructions(_instructionId [32]byte, _extensionId *big.Int, _teeIds []common.Address, _opType [32]byte, _opCommand [32]byte, _message []byte) (*types.Transaction, error) {
	return _TeeExtensionRegistry.Contract.SendInstructions(&_TeeExtensionRegistry.TransactOpts, _instructionId, _extensionId, _teeIds, _opType, _opCommand, _message)
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

// TeeExtensionRegistryExtensionContractsSetIterator is returned from FilterExtensionContractsSet and is used to iterate over the raw logs and unpacked data for ExtensionContractsSet events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryExtensionContractsSetIterator struct {
	Event *TeeExtensionRegistryExtensionContractsSet // Event containing the contract specifics and raw log

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
func (it *TeeExtensionRegistryExtensionContractsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryExtensionContractsSet)
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
		it.Event = new(TeeExtensionRegistryExtensionContractsSet)
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
func (it *TeeExtensionRegistryExtensionContractsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryExtensionContractsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryExtensionContractsSet represents a ExtensionContractsSet event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryExtensionContractsSet struct {
	ExtensionId                    *big.Int
	TeeExtensionStateVerifier      common.Address
	TeeExtensionInstructionsSender common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterExtensionContractsSet is a free log retrieval operation binding the contract event 0xcce5f286edee14b3fcc5011e930a7044a1fcdb4c16259f8d131b6a32b1d87286.
//
// Solidity: event ExtensionContractsSet(uint256 indexed extensionId, address indexed teeExtensionStateVerifier, address indexed teeExtensionInstructionsSender)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterExtensionContractsSet(opts *bind.FilterOpts, extensionId []*big.Int, teeExtensionStateVerifier []common.Address, teeExtensionInstructionsSender []common.Address) (*TeeExtensionRegistryExtensionContractsSetIterator, error) {

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

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "ExtensionContractsSet", extensionIdRule, teeExtensionStateVerifierRule, teeExtensionInstructionsSenderRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryExtensionContractsSetIterator{contract: _TeeExtensionRegistry.contract, event: "ExtensionContractsSet", logs: logs, sub: sub}, nil
}

// WatchExtensionContractsSet is a free log subscription operation binding the contract event 0xcce5f286edee14b3fcc5011e930a7044a1fcdb4c16259f8d131b6a32b1d87286.
//
// Solidity: event ExtensionContractsSet(uint256 indexed extensionId, address indexed teeExtensionStateVerifier, address indexed teeExtensionInstructionsSender)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchExtensionContractsSet(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryExtensionContractsSet, extensionId []*big.Int, teeExtensionStateVerifier []common.Address, teeExtensionInstructionsSender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "ExtensionContractsSet", extensionIdRule, teeExtensionStateVerifierRule, teeExtensionInstructionsSenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryExtensionContractsSet)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "ExtensionContractsSet", log); err != nil {
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

// ParseExtensionContractsSet is a log parse operation binding the contract event 0xcce5f286edee14b3fcc5011e930a7044a1fcdb4c16259f8d131b6a32b1d87286.
//
// Solidity: event ExtensionContractsSet(uint256 indexed extensionId, address indexed teeExtensionStateVerifier, address indexed teeExtensionInstructionsSender)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseExtensionContractsSet(log types.Log) (*TeeExtensionRegistryExtensionContractsSet, error) {
	event := new(TeeExtensionRegistryExtensionContractsSet)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "ExtensionContractsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryExtensionRegisteredIterator is returned from FilterExtensionRegistered and is used to iterate over the raw logs and unpacked data for ExtensionRegistered events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryExtensionRegisteredIterator struct {
	Event *TeeExtensionRegistryExtensionRegistered // Event containing the contract specifics and raw log

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
func (it *TeeExtensionRegistryExtensionRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryExtensionRegistered)
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
		it.Event = new(TeeExtensionRegistryExtensionRegistered)
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
func (it *TeeExtensionRegistryExtensionRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryExtensionRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryExtensionRegistered represents a ExtensionRegistered event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryExtensionRegistered struct {
	ExtensionId *big.Int
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExtensionRegistered is a free log retrieval operation binding the contract event 0x040fcbb1e96b0cba5c9e4ff6e5e585a8c815ddd591dff3f19fb56fc12c9f3cb6.
//
// Solidity: event ExtensionRegistered(uint256 indexed extensionId, address indexed owner)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterExtensionRegistered(opts *bind.FilterOpts, extensionId []*big.Int, owner []common.Address) (*TeeExtensionRegistryExtensionRegisteredIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "ExtensionRegistered", extensionIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryExtensionRegisteredIterator{contract: _TeeExtensionRegistry.contract, event: "ExtensionRegistered", logs: logs, sub: sub}, nil
}

// WatchExtensionRegistered is a free log subscription operation binding the contract event 0x040fcbb1e96b0cba5c9e4ff6e5e585a8c815ddd591dff3f19fb56fc12c9f3cb6.
//
// Solidity: event ExtensionRegistered(uint256 indexed extensionId, address indexed owner)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchExtensionRegistered(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryExtensionRegistered, extensionId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "ExtensionRegistered", extensionIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryExtensionRegistered)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "ExtensionRegistered", log); err != nil {
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

// ParseExtensionRegistered is a log parse operation binding the contract event 0x040fcbb1e96b0cba5c9e4ff6e5e585a8c815ddd591dff3f19fb56fc12c9f3cb6.
//
// Solidity: event ExtensionRegistered(uint256 indexed extensionId, address indexed owner)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseExtensionRegistered(log types.Log) (*TeeExtensionRegistryExtensionRegistered, error) {
	event := new(TeeExtensionRegistryExtensionRegistered)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "ExtensionRegistered", log); err != nil {
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

// TeeExtensionRegistryOpTypeAddedIterator is returned from FilterOpTypeAdded and is used to iterate over the raw logs and unpacked data for OpTypeAdded events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryOpTypeAddedIterator struct {
	Event *TeeExtensionRegistryOpTypeAdded // Event containing the contract specifics and raw log

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
func (it *TeeExtensionRegistryOpTypeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryOpTypeAdded)
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
		it.Event = new(TeeExtensionRegistryOpTypeAdded)
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
func (it *TeeExtensionRegistryOpTypeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryOpTypeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryOpTypeAdded represents a OpTypeAdded event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryOpTypeAdded struct {
	ExtensionId *big.Int
	OpType      [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOpTypeAdded is a free log retrieval operation binding the contract event 0xf8022e8cd39ecc9c555fd13b9b97af40a880e914c24d08f234ae1114791b51f6.
//
// Solidity: event OpTypeAdded(uint256 indexed extensionId, bytes32 indexed opType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterOpTypeAdded(opts *bind.FilterOpts, extensionId []*big.Int, opType [][32]byte) (*TeeExtensionRegistryOpTypeAddedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var opTypeRule []interface{}
	for _, opTypeItem := range opType {
		opTypeRule = append(opTypeRule, opTypeItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "OpTypeAdded", extensionIdRule, opTypeRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryOpTypeAddedIterator{contract: _TeeExtensionRegistry.contract, event: "OpTypeAdded", logs: logs, sub: sub}, nil
}

// WatchOpTypeAdded is a free log subscription operation binding the contract event 0xf8022e8cd39ecc9c555fd13b9b97af40a880e914c24d08f234ae1114791b51f6.
//
// Solidity: event OpTypeAdded(uint256 indexed extensionId, bytes32 indexed opType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchOpTypeAdded(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryOpTypeAdded, extensionId []*big.Int, opType [][32]byte) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var opTypeRule []interface{}
	for _, opTypeItem := range opType {
		opTypeRule = append(opTypeRule, opTypeItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "OpTypeAdded", extensionIdRule, opTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryOpTypeAdded)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "OpTypeAdded", log); err != nil {
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

// ParseOpTypeAdded is a log parse operation binding the contract event 0xf8022e8cd39ecc9c555fd13b9b97af40a880e914c24d08f234ae1114791b51f6.
//
// Solidity: event OpTypeAdded(uint256 indexed extensionId, bytes32 indexed opType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseOpTypeAdded(log types.Log) (*TeeExtensionRegistryOpTypeAdded, error) {
	event := new(TeeExtensionRegistryOpTypeAdded)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "OpTypeAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryOpTypeRemovedIterator is returned from FilterOpTypeRemoved and is used to iterate over the raw logs and unpacked data for OpTypeRemoved events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryOpTypeRemovedIterator struct {
	Event *TeeExtensionRegistryOpTypeRemoved // Event containing the contract specifics and raw log

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
func (it *TeeExtensionRegistryOpTypeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryOpTypeRemoved)
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
		it.Event = new(TeeExtensionRegistryOpTypeRemoved)
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
func (it *TeeExtensionRegistryOpTypeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryOpTypeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryOpTypeRemoved represents a OpTypeRemoved event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryOpTypeRemoved struct {
	ExtensionId *big.Int
	OpType      [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOpTypeRemoved is a free log retrieval operation binding the contract event 0xc22efddd00dbabbf901448f58a5f8ad61df1c6cd2507b79a1ae5711c9f510761.
//
// Solidity: event OpTypeRemoved(uint256 indexed extensionId, bytes32 indexed opType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterOpTypeRemoved(opts *bind.FilterOpts, extensionId []*big.Int, opType [][32]byte) (*TeeExtensionRegistryOpTypeRemovedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var opTypeRule []interface{}
	for _, opTypeItem := range opType {
		opTypeRule = append(opTypeRule, opTypeItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "OpTypeRemoved", extensionIdRule, opTypeRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryOpTypeRemovedIterator{contract: _TeeExtensionRegistry.contract, event: "OpTypeRemoved", logs: logs, sub: sub}, nil
}

// WatchOpTypeRemoved is a free log subscription operation binding the contract event 0xc22efddd00dbabbf901448f58a5f8ad61df1c6cd2507b79a1ae5711c9f510761.
//
// Solidity: event OpTypeRemoved(uint256 indexed extensionId, bytes32 indexed opType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchOpTypeRemoved(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryOpTypeRemoved, extensionId []*big.Int, opType [][32]byte) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var opTypeRule []interface{}
	for _, opTypeItem := range opType {
		opTypeRule = append(opTypeRule, opTypeItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "OpTypeRemoved", extensionIdRule, opTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryOpTypeRemoved)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "OpTypeRemoved", log); err != nil {
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

// ParseOpTypeRemoved is a log parse operation binding the contract event 0xc22efddd00dbabbf901448f58a5f8ad61df1c6cd2507b79a1ae5711c9f510761.
//
// Solidity: event OpTypeRemoved(uint256 indexed extensionId, bytes32 indexed opType)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParseOpTypeRemoved(log types.Log) (*TeeExtensionRegistryOpTypeRemoved, error) {
	event := new(TeeExtensionRegistryOpTypeRemoved)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "OpTypeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeExtensionRegistryPlatformAddedIterator is returned from FilterPlatformAdded and is used to iterate over the raw logs and unpacked data for PlatformAdded events raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryPlatformAddedIterator struct {
	Event *TeeExtensionRegistryPlatformAdded // Event containing the contract specifics and raw log

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
func (it *TeeExtensionRegistryPlatformAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeExtensionRegistryPlatformAdded)
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
		it.Event = new(TeeExtensionRegistryPlatformAdded)
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
func (it *TeeExtensionRegistryPlatformAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeExtensionRegistryPlatformAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeExtensionRegistryPlatformAdded represents a PlatformAdded event raised by the TeeExtensionRegistry contract.
type TeeExtensionRegistryPlatformAdded struct {
	Platform [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPlatformAdded is a free log retrieval operation binding the contract event 0x1faeda337d831b32cfd4d60241aa89f3c5f1f10688cba6c8a145a6105d8bec31.
//
// Solidity: event PlatformAdded(bytes32 indexed platform)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) FilterPlatformAdded(opts *bind.FilterOpts, platform [][32]byte) (*TeeExtensionRegistryPlatformAddedIterator, error) {

	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.FilterLogs(opts, "PlatformAdded", platformRule)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionRegistryPlatformAddedIterator{contract: _TeeExtensionRegistry.contract, event: "PlatformAdded", logs: logs, sub: sub}, nil
}

// WatchPlatformAdded is a free log subscription operation binding the contract event 0x1faeda337d831b32cfd4d60241aa89f3c5f1f10688cba6c8a145a6105d8bec31.
//
// Solidity: event PlatformAdded(bytes32 indexed platform)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) WatchPlatformAdded(opts *bind.WatchOpts, sink chan<- *TeeExtensionRegistryPlatformAdded, platform [][32]byte) (event.Subscription, error) {

	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _TeeExtensionRegistry.contract.WatchLogs(opts, "PlatformAdded", platformRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeExtensionRegistryPlatformAdded)
				if err := _TeeExtensionRegistry.contract.UnpackLog(event, "PlatformAdded", log); err != nil {
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

// ParsePlatformAdded is a log parse operation binding the contract event 0x1faeda337d831b32cfd4d60241aa89f3c5f1f10688cba6c8a145a6105d8bec31.
//
// Solidity: event PlatformAdded(bytes32 indexed platform)
func (_TeeExtensionRegistry *TeeExtensionRegistryFilterer) ParsePlatformAdded(log types.Log) (*TeeExtensionRegistryPlatformAdded, error) {
	event := new(TeeExtensionRegistryPlatformAdded)
	if err := _TeeExtensionRegistry.contract.UnpackLog(event, "PlatformAdded", log); err != nil {
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
