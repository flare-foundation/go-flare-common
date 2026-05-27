// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package machinepath

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

// IMachinePathManagerMachinePath is an auto generated low-level Go binding around an user-defined struct.
type IMachinePathManagerMachinePath struct {
	SourceTeeIds      []common.Address
	DestinationTeeIds []common.Address
}

// TeeMachinePathMetaData contains all meta data concerning the TeeMachinePath contract.
var TeeMachinePathMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"sourceTeeIds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"destinationTeeIds\",\"type\":\"address[]\"}],\"internalType\":\"structIMachinePathManager.MachinePath\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"machinePathStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeMachinePathABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeMachinePathMetaData.ABI instead.
var TeeMachinePathABI = TeeMachinePathMetaData.ABI

// TeeMachinePath is an auto generated Go binding around an Ethereum contract.
type TeeMachinePath struct {
	TeeMachinePathCaller     // Read-only binding to the contract
	TeeMachinePathTransactor // Write-only binding to the contract
	TeeMachinePathFilterer   // Log filterer for contract events
}

// TeeMachinePathCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeMachinePathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeMachinePathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeMachinePathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeMachinePathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeMachinePathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeMachinePathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeMachinePathSession struct {
	Contract     *TeeMachinePath   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeMachinePathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeMachinePathCallerSession struct {
	Contract *TeeMachinePathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TeeMachinePathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeMachinePathTransactorSession struct {
	Contract     *TeeMachinePathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TeeMachinePathRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeMachinePathRaw struct {
	Contract *TeeMachinePath // Generic contract binding to access the raw methods on
}

// TeeMachinePathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeMachinePathCallerRaw struct {
	Contract *TeeMachinePathCaller // Generic read-only contract binding to access the raw methods on
}

// TeeMachinePathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeMachinePathTransactorRaw struct {
	Contract *TeeMachinePathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeMachinePath creates a new instance of TeeMachinePath, bound to a specific deployed contract.
func NewTeeMachinePath(address common.Address, backend bind.ContractBackend) (*TeeMachinePath, error) {
	contract, err := bindTeeMachinePath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeMachinePath{TeeMachinePathCaller: TeeMachinePathCaller{contract: contract}, TeeMachinePathTransactor: TeeMachinePathTransactor{contract: contract}, TeeMachinePathFilterer: TeeMachinePathFilterer{contract: contract}}, nil
}

// NewTeeMachinePathCaller creates a new read-only instance of TeeMachinePath, bound to a specific deployed contract.
func NewTeeMachinePathCaller(address common.Address, caller bind.ContractCaller) (*TeeMachinePathCaller, error) {
	contract, err := bindTeeMachinePath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeMachinePathCaller{contract: contract}, nil
}

// NewTeeMachinePathTransactor creates a new write-only instance of TeeMachinePath, bound to a specific deployed contract.
func NewTeeMachinePathTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeMachinePathTransactor, error) {
	contract, err := bindTeeMachinePath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeMachinePathTransactor{contract: contract}, nil
}

// NewTeeMachinePathFilterer creates a new log filterer instance of TeeMachinePath, bound to a specific deployed contract.
func NewTeeMachinePathFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeMachinePathFilterer, error) {
	contract, err := bindTeeMachinePath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeMachinePathFilterer{contract: contract}, nil
}

// bindTeeMachinePath binds a generic wrapper to an already deployed contract.
func bindTeeMachinePath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeMachinePathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeMachinePath *TeeMachinePathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeMachinePath.Contract.TeeMachinePathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeMachinePath *TeeMachinePathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeMachinePath.Contract.TeeMachinePathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeMachinePath *TeeMachinePathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeMachinePath.Contract.TeeMachinePathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeMachinePath *TeeMachinePathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeMachinePath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeMachinePath *TeeMachinePathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeMachinePath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeMachinePath *TeeMachinePathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeMachinePath.Contract.contract.Transact(opts, method, params...)
}

// MachinePathStruct is a paid mutator transaction binding the contract method 0xe8e6b1c1.
//
// Solidity: function machinePathStruct((address[],address[]) ) returns()
func (_TeeMachinePath *TeeMachinePathTransactor) MachinePathStruct(opts *bind.TransactOpts, arg0 IMachinePathManagerMachinePath) (*types.Transaction, error) {
	return _TeeMachinePath.contract.Transact(opts, "machinePathStruct", arg0)
}

// MachinePathStruct is a paid mutator transaction binding the contract method 0xe8e6b1c1.
//
// Solidity: function machinePathStruct((address[],address[]) ) returns()
func (_TeeMachinePath *TeeMachinePathSession) MachinePathStruct(arg0 IMachinePathManagerMachinePath) (*types.Transaction, error) {
	return _TeeMachinePath.Contract.MachinePathStruct(&_TeeMachinePath.TransactOpts, arg0)
}

// MachinePathStruct is a paid mutator transaction binding the contract method 0xe8e6b1c1.
//
// Solidity: function machinePathStruct((address[],address[]) ) returns()
func (_TeeMachinePath *TeeMachinePathTransactorSession) MachinePathStruct(arg0 IMachinePathManagerMachinePath) (*types.Transaction, error) {
	return _TeeMachinePath.Contract.MachinePathStruct(&_TeeMachinePath.TransactOpts, arg0)
}
