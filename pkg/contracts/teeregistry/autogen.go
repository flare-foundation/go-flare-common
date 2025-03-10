// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teeregistry

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

// ITeeRegistryTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryTeeMachine struct {
	TeeId common.Address
	Owner common.Address
	Url   string
}

// ITeeRegistryTeeMachineWithAttestationData is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryTeeMachineWithAttestationData struct {
	TeeId    common.Address
	Owner    common.Address
	Url      string
	CodeHash [32]byte
	Platform [32]byte
}

// TeeRegistryMetaData contains all meta data concerning the TeeRegistry contract.
var TeeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_backupTeeIds\",\"type\":\"address[]\"}],\"name\":\"arePlatformsCompatible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachine\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structITeeRegistry.TeeMachine\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineStatus\",\"outputs\":[{\"internalType\":\"enumITeeRegistry.TeeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineWithAttestationData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"}],\"name\":\"isOpTypeSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TeeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeRegistryMetaData.ABI instead.
var TeeRegistryABI = TeeRegistryMetaData.ABI

// TeeRegistry is an auto generated Go binding around an Ethereum contract.
type TeeRegistry struct {
	TeeRegistryCaller     // Read-only binding to the contract
	TeeRegistryTransactor // Write-only binding to the contract
	TeeRegistryFilterer   // Log filterer for contract events
}

// TeeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeRegistrySession struct {
	Contract     *TeeRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeRegistryCallerSession struct {
	Contract *TeeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TeeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeRegistryTransactorSession struct {
	Contract     *TeeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TeeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeRegistryRaw struct {
	Contract *TeeRegistry // Generic contract binding to access the raw methods on
}

// TeeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeRegistryCallerRaw struct {
	Contract *TeeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TeeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeRegistryTransactorRaw struct {
	Contract *TeeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeRegistry creates a new instance of TeeRegistry, bound to a specific deployed contract.
func NewTeeRegistry(address common.Address, backend bind.ContractBackend) (*TeeRegistry, error) {
	contract, err := bindTeeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeRegistry{TeeRegistryCaller: TeeRegistryCaller{contract: contract}, TeeRegistryTransactor: TeeRegistryTransactor{contract: contract}, TeeRegistryFilterer: TeeRegistryFilterer{contract: contract}}, nil
}

// NewTeeRegistryCaller creates a new read-only instance of TeeRegistry, bound to a specific deployed contract.
func NewTeeRegistryCaller(address common.Address, caller bind.ContractCaller) (*TeeRegistryCaller, error) {
	contract, err := bindTeeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryCaller{contract: contract}, nil
}

// NewTeeRegistryTransactor creates a new write-only instance of TeeRegistry, bound to a specific deployed contract.
func NewTeeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeRegistryTransactor, error) {
	contract, err := bindTeeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryTransactor{contract: contract}, nil
}

// NewTeeRegistryFilterer creates a new log filterer instance of TeeRegistry, bound to a specific deployed contract.
func NewTeeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeRegistryFilterer, error) {
	contract, err := bindTeeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryFilterer{contract: contract}, nil
}

// bindTeeRegistry binds a generic wrapper to an already deployed contract.
func bindTeeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeRegistry *TeeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeRegistry.Contract.TeeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeRegistry *TeeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeRegistry.Contract.TeeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeRegistry *TeeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeRegistry.Contract.TeeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeRegistry *TeeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeRegistry *TeeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeRegistry *TeeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeRegistry.Contract.contract.Transact(opts, method, params...)
}

// ArePlatformsCompatible is a free data retrieval call binding the contract method 0x2bcabed9.
//
// Solidity: function arePlatformsCompatible(address _teeId, address[] _backupTeeIds) view returns(bool)
func (_TeeRegistry *TeeRegistryCaller) ArePlatformsCompatible(opts *bind.CallOpts, _teeId common.Address, _backupTeeIds []common.Address) (bool, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "arePlatformsCompatible", _teeId, _backupTeeIds)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ArePlatformsCompatible is a free data retrieval call binding the contract method 0x2bcabed9.
//
// Solidity: function arePlatformsCompatible(address _teeId, address[] _backupTeeIds) view returns(bool)
func (_TeeRegistry *TeeRegistrySession) ArePlatformsCompatible(_teeId common.Address, _backupTeeIds []common.Address) (bool, error) {
	return _TeeRegistry.Contract.ArePlatformsCompatible(&_TeeRegistry.CallOpts, _teeId, _backupTeeIds)
}

// ArePlatformsCompatible is a free data retrieval call binding the contract method 0x2bcabed9.
//
// Solidity: function arePlatformsCompatible(address _teeId, address[] _backupTeeIds) view returns(bool)
func (_TeeRegistry *TeeRegistryCallerSession) ArePlatformsCompatible(_teeId common.Address, _backupTeeIds []common.Address) (bool, error) {
	return _TeeRegistry.Contract.ArePlatformsCompatible(&_TeeRegistry.CallOpts, _teeId, _backupTeeIds)
}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_TeeRegistry *TeeRegistryCaller) GetTeeMachine(opts *bind.CallOpts, _teeId common.Address) (ITeeRegistryTeeMachine, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getTeeMachine", _teeId)

	if err != nil {
		return *new(ITeeRegistryTeeMachine), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeeRegistryTeeMachine)).(*ITeeRegistryTeeMachine)

	return out0, err

}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_TeeRegistry *TeeRegistrySession) GetTeeMachine(_teeId common.Address) (ITeeRegistryTeeMachine, error) {
	return _TeeRegistry.Contract.GetTeeMachine(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_TeeRegistry *TeeRegistryCallerSession) GetTeeMachine(_teeId common.Address) (ITeeRegistryTeeMachine, error) {
	return _TeeRegistry.Contract.GetTeeMachine(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_TeeRegistry *TeeRegistryCaller) GetTeeMachineStatus(opts *bind.CallOpts, _teeId common.Address) (uint8, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getTeeMachineStatus", _teeId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_TeeRegistry *TeeRegistrySession) GetTeeMachineStatus(_teeId common.Address) (uint8, error) {
	return _TeeRegistry.Contract.GetTeeMachineStatus(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_TeeRegistry *TeeRegistryCallerSession) GetTeeMachineStatus(_teeId common.Address) (uint8, error) {
	return _TeeRegistry.Contract.GetTeeMachineStatus(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_TeeRegistry *TeeRegistryCaller) GetTeeMachineWithAttestationData(opts *bind.CallOpts, _teeId common.Address) (ITeeRegistryTeeMachineWithAttestationData, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getTeeMachineWithAttestationData", _teeId)

	if err != nil {
		return *new(ITeeRegistryTeeMachineWithAttestationData), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeeRegistryTeeMachineWithAttestationData)).(*ITeeRegistryTeeMachineWithAttestationData)

	return out0, err

}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_TeeRegistry *TeeRegistrySession) GetTeeMachineWithAttestationData(_teeId common.Address) (ITeeRegistryTeeMachineWithAttestationData, error) {
	return _TeeRegistry.Contract.GetTeeMachineWithAttestationData(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_TeeRegistry *TeeRegistryCallerSession) GetTeeMachineWithAttestationData(_teeId common.Address) (ITeeRegistryTeeMachineWithAttestationData, error) {
	return _TeeRegistry.Contract.GetTeeMachineWithAttestationData(&_TeeRegistry.CallOpts, _teeId)
}

// IsOpTypeSupported is a free data retrieval call binding the contract method 0xf132d9c8.
//
// Solidity: function isOpTypeSupported(address _teeId, bytes32 _opType) view returns(bool)
func (_TeeRegistry *TeeRegistryCaller) IsOpTypeSupported(opts *bind.CallOpts, _teeId common.Address, _opType [32]byte) (bool, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "isOpTypeSupported", _teeId, _opType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpTypeSupported is a free data retrieval call binding the contract method 0xf132d9c8.
//
// Solidity: function isOpTypeSupported(address _teeId, bytes32 _opType) view returns(bool)
func (_TeeRegistry *TeeRegistrySession) IsOpTypeSupported(_teeId common.Address, _opType [32]byte) (bool, error) {
	return _TeeRegistry.Contract.IsOpTypeSupported(&_TeeRegistry.CallOpts, _teeId, _opType)
}

// IsOpTypeSupported is a free data retrieval call binding the contract method 0xf132d9c8.
//
// Solidity: function isOpTypeSupported(address _teeId, bytes32 _opType) view returns(bool)
func (_TeeRegistry *TeeRegistryCallerSession) IsOpTypeSupported(_teeId common.Address, _opType [32]byte) (bool, error) {
	return _TeeRegistry.Contract.IsOpTypeSupported(&_TeeRegistry.CallOpts, _teeId, _opType)
}
