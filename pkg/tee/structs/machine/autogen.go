// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package machine

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

// IMachineManagerTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
}

// IMachineManagerTeeMachineData is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachineData struct {
	ExtensionId  *big.Int
	InitialOwner common.Address
	CodeHash     [32]byte
	Platform     [32]byte
	PublicKey    PublicKey
}

// IMachineManagerTeeMachineWithAttestationData is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachineWithAttestationData struct {
	TeeId        common.Address
	InitialTeeId common.Address
	Url          string
	CodeHash     [32]byte
	Platform     [32]byte
}

// PublicKey is an auto generated low-level Go binding around an user-defined struct.
type PublicKey struct {
	X [32]byte
	Y [32]byte
}

// TeeMachineMetaData contains all meta data concerning the TeeMachine contract.
var TeeMachineMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"internalType\":\"structIMachineManager.TeeMachineData\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeMachineDataStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIMachineManager.TeeMachine\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeMachineStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIMachineManager.TeeMachineWithAttestationData\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeMachineWithAttestationDataStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeMachineABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeMachineMetaData.ABI instead.
var TeeMachineABI = TeeMachineMetaData.ABI

// TeeMachine is an auto generated Go binding around an Ethereum contract.
type TeeMachine struct {
	TeeMachineCaller     // Read-only binding to the contract
	TeeMachineTransactor // Write-only binding to the contract
	TeeMachineFilterer   // Log filterer for contract events
}

// TeeMachineCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeMachineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeMachineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeMachineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeMachineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeMachineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeMachineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeMachineSession struct {
	Contract     *TeeMachine       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeMachineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeMachineCallerSession struct {
	Contract *TeeMachineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TeeMachineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeMachineTransactorSession struct {
	Contract     *TeeMachineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TeeMachineRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeMachineRaw struct {
	Contract *TeeMachine // Generic contract binding to access the raw methods on
}

// TeeMachineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeMachineCallerRaw struct {
	Contract *TeeMachineCaller // Generic read-only contract binding to access the raw methods on
}

// TeeMachineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeMachineTransactorRaw struct {
	Contract *TeeMachineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeMachine creates a new instance of TeeMachine, bound to a specific deployed contract.
func NewTeeMachine(address common.Address, backend bind.ContractBackend) (*TeeMachine, error) {
	contract, err := bindTeeMachine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeMachine{TeeMachineCaller: TeeMachineCaller{contract: contract}, TeeMachineTransactor: TeeMachineTransactor{contract: contract}, TeeMachineFilterer: TeeMachineFilterer{contract: contract}}, nil
}

// NewTeeMachineCaller creates a new read-only instance of TeeMachine, bound to a specific deployed contract.
func NewTeeMachineCaller(address common.Address, caller bind.ContractCaller) (*TeeMachineCaller, error) {
	contract, err := bindTeeMachine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeMachineCaller{contract: contract}, nil
}

// NewTeeMachineTransactor creates a new write-only instance of TeeMachine, bound to a specific deployed contract.
func NewTeeMachineTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeMachineTransactor, error) {
	contract, err := bindTeeMachine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeMachineTransactor{contract: contract}, nil
}

// NewTeeMachineFilterer creates a new log filterer instance of TeeMachine, bound to a specific deployed contract.
func NewTeeMachineFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeMachineFilterer, error) {
	contract, err := bindTeeMachine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeMachineFilterer{contract: contract}, nil
}

// bindTeeMachine binds a generic wrapper to an already deployed contract.
func bindTeeMachine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeMachineMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeMachine *TeeMachineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeMachine.Contract.TeeMachineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeMachine *TeeMachineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeMachine.Contract.TeeMachineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeMachine *TeeMachineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeMachine.Contract.TeeMachineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeMachine *TeeMachineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeMachine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeMachine *TeeMachineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeMachine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeMachine *TeeMachineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeMachine.Contract.contract.Transact(opts, method, params...)
}

// TeeMachineDataStruct is a paid mutator transaction binding the contract method 0x6d3d9ce3.
//
// Solidity: function teeMachineDataStruct((uint256,address,bytes32,bytes32,(bytes32,bytes32)) ) returns()
func (_TeeMachine *TeeMachineTransactor) TeeMachineDataStruct(opts *bind.TransactOpts, arg0 IMachineManagerTeeMachineData) (*types.Transaction, error) {
	return _TeeMachine.contract.Transact(opts, "teeMachineDataStruct", arg0)
}

// TeeMachineDataStruct is a paid mutator transaction binding the contract method 0x6d3d9ce3.
//
// Solidity: function teeMachineDataStruct((uint256,address,bytes32,bytes32,(bytes32,bytes32)) ) returns()
func (_TeeMachine *TeeMachineSession) TeeMachineDataStruct(arg0 IMachineManagerTeeMachineData) (*types.Transaction, error) {
	return _TeeMachine.Contract.TeeMachineDataStruct(&_TeeMachine.TransactOpts, arg0)
}

// TeeMachineDataStruct is a paid mutator transaction binding the contract method 0x6d3d9ce3.
//
// Solidity: function teeMachineDataStruct((uint256,address,bytes32,bytes32,(bytes32,bytes32)) ) returns()
func (_TeeMachine *TeeMachineTransactorSession) TeeMachineDataStruct(arg0 IMachineManagerTeeMachineData) (*types.Transaction, error) {
	return _TeeMachine.Contract.TeeMachineDataStruct(&_TeeMachine.TransactOpts, arg0)
}

// TeeMachineStruct is a paid mutator transaction binding the contract method 0xef60cfee.
//
// Solidity: function teeMachineStruct((address,address,string) ) returns()
func (_TeeMachine *TeeMachineTransactor) TeeMachineStruct(opts *bind.TransactOpts, arg0 IMachineManagerTeeMachine) (*types.Transaction, error) {
	return _TeeMachine.contract.Transact(opts, "teeMachineStruct", arg0)
}

// TeeMachineStruct is a paid mutator transaction binding the contract method 0xef60cfee.
//
// Solidity: function teeMachineStruct((address,address,string) ) returns()
func (_TeeMachine *TeeMachineSession) TeeMachineStruct(arg0 IMachineManagerTeeMachine) (*types.Transaction, error) {
	return _TeeMachine.Contract.TeeMachineStruct(&_TeeMachine.TransactOpts, arg0)
}

// TeeMachineStruct is a paid mutator transaction binding the contract method 0xef60cfee.
//
// Solidity: function teeMachineStruct((address,address,string) ) returns()
func (_TeeMachine *TeeMachineTransactorSession) TeeMachineStruct(arg0 IMachineManagerTeeMachine) (*types.Transaction, error) {
	return _TeeMachine.Contract.TeeMachineStruct(&_TeeMachine.TransactOpts, arg0)
}

// TeeMachineWithAttestationDataStruct is a paid mutator transaction binding the contract method 0xd77f63b9.
//
// Solidity: function teeMachineWithAttestationDataStruct((address,address,string,bytes32,bytes32) ) returns()
func (_TeeMachine *TeeMachineTransactor) TeeMachineWithAttestationDataStruct(opts *bind.TransactOpts, arg0 IMachineManagerTeeMachineWithAttestationData) (*types.Transaction, error) {
	return _TeeMachine.contract.Transact(opts, "teeMachineWithAttestationDataStruct", arg0)
}

// TeeMachineWithAttestationDataStruct is a paid mutator transaction binding the contract method 0xd77f63b9.
//
// Solidity: function teeMachineWithAttestationDataStruct((address,address,string,bytes32,bytes32) ) returns()
func (_TeeMachine *TeeMachineSession) TeeMachineWithAttestationDataStruct(arg0 IMachineManagerTeeMachineWithAttestationData) (*types.Transaction, error) {
	return _TeeMachine.Contract.TeeMachineWithAttestationDataStruct(&_TeeMachine.TransactOpts, arg0)
}

// TeeMachineWithAttestationDataStruct is a paid mutator transaction binding the contract method 0xd77f63b9.
//
// Solidity: function teeMachineWithAttestationDataStruct((address,address,string,bytes32,bytes32) ) returns()
func (_TeeMachine *TeeMachineTransactorSession) TeeMachineWithAttestationDataStruct(arg0 IMachineManagerTeeMachineWithAttestationData) (*types.Transaction, error) {
	return _TeeMachine.Contract.TeeMachineWithAttestationDataStruct(&_TeeMachine.TransactOpts, arg0)
}
