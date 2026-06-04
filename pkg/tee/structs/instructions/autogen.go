// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package instructions

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

// IInstructionsTeeInstructionParams is an auto generated low-level Go binding around an user-defined struct.
type IInstructionsTeeInstructionParams struct {
	OpType             [32]byte
	OpCommand          [32]byte
	Message            []byte
	Cosigners          []common.Address
	CosignersThreshold uint64
	ClaimBackAddress   common.Address
}

// TeeInstructionsMetaData contains all meta data concerning the TeeInstructions contract.
var TeeInstructionsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"}],\"internalType\":\"structIInstructions.TeeInstructionParams\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeInstructionParamsStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeInstructionsABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeInstructionsMetaData.ABI instead.
var TeeInstructionsABI = TeeInstructionsMetaData.ABI

// TeeInstructions is an auto generated Go binding around an Ethereum contract.
type TeeInstructions struct {
	TeeInstructionsCaller     // Read-only binding to the contract
	TeeInstructionsTransactor // Write-only binding to the contract
	TeeInstructionsFilterer   // Log filterer for contract events
}

// TeeInstructionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeInstructionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeInstructionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeInstructionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeInstructionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeInstructionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeInstructionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeInstructionsSession struct {
	Contract     *TeeInstructions  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeInstructionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeInstructionsCallerSession struct {
	Contract *TeeInstructionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TeeInstructionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeInstructionsTransactorSession struct {
	Contract     *TeeInstructionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TeeInstructionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeInstructionsRaw struct {
	Contract *TeeInstructions // Generic contract binding to access the raw methods on
}

// TeeInstructionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeInstructionsCallerRaw struct {
	Contract *TeeInstructionsCaller // Generic read-only contract binding to access the raw methods on
}

// TeeInstructionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeInstructionsTransactorRaw struct {
	Contract *TeeInstructionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeInstructions creates a new instance of TeeInstructions, bound to a specific deployed contract.
func NewTeeInstructions(address common.Address, backend bind.ContractBackend) (*TeeInstructions, error) {
	contract, err := bindTeeInstructions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeInstructions{TeeInstructionsCaller: TeeInstructionsCaller{contract: contract}, TeeInstructionsTransactor: TeeInstructionsTransactor{contract: contract}, TeeInstructionsFilterer: TeeInstructionsFilterer{contract: contract}}, nil
}

// NewTeeInstructionsCaller creates a new read-only instance of TeeInstructions, bound to a specific deployed contract.
func NewTeeInstructionsCaller(address common.Address, caller bind.ContractCaller) (*TeeInstructionsCaller, error) {
	contract, err := bindTeeInstructions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsCaller{contract: contract}, nil
}

// NewTeeInstructionsTransactor creates a new write-only instance of TeeInstructions, bound to a specific deployed contract.
func NewTeeInstructionsTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeInstructionsTransactor, error) {
	contract, err := bindTeeInstructions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsTransactor{contract: contract}, nil
}

// NewTeeInstructionsFilterer creates a new log filterer instance of TeeInstructions, bound to a specific deployed contract.
func NewTeeInstructionsFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeInstructionsFilterer, error) {
	contract, err := bindTeeInstructions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsFilterer{contract: contract}, nil
}

// bindTeeInstructions binds a generic wrapper to an already deployed contract.
func bindTeeInstructions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeInstructionsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeInstructions *TeeInstructionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeInstructions.Contract.TeeInstructionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeInstructions *TeeInstructionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeInstructions.Contract.TeeInstructionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeInstructions *TeeInstructionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeInstructions.Contract.TeeInstructionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeInstructions *TeeInstructionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeInstructions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeInstructions *TeeInstructionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeInstructions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeInstructions *TeeInstructionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeInstructions.Contract.contract.Transact(opts, method, params...)
}

// TeeInstructionParamsStruct is a paid mutator transaction binding the contract method 0x6cc4ff22.
//
// Solidity: function teeInstructionParamsStruct((bytes32,bytes32,bytes,address[],uint64,address) ) returns()
func (_TeeInstructions *TeeInstructionsTransactor) TeeInstructionParamsStruct(opts *bind.TransactOpts, arg0 IInstructionsTeeInstructionParams) (*types.Transaction, error) {
	return _TeeInstructions.contract.Transact(opts, "teeInstructionParamsStruct", arg0)
}

// TeeInstructionParamsStruct is a paid mutator transaction binding the contract method 0x6cc4ff22.
//
// Solidity: function teeInstructionParamsStruct((bytes32,bytes32,bytes,address[],uint64,address) ) returns()
func (_TeeInstructions *TeeInstructionsSession) TeeInstructionParamsStruct(arg0 IInstructionsTeeInstructionParams) (*types.Transaction, error) {
	return _TeeInstructions.Contract.TeeInstructionParamsStruct(&_TeeInstructions.TransactOpts, arg0)
}

// TeeInstructionParamsStruct is a paid mutator transaction binding the contract method 0x6cc4ff22.
//
// Solidity: function teeInstructionParamsStruct((bytes32,bytes32,bytes,address[],uint64,address) ) returns()
func (_TeeInstructions *TeeInstructionsTransactorSession) TeeInstructionParamsStruct(arg0 IInstructionsTeeInstructionParams) (*types.Transaction, error) {
	return _TeeInstructions.Contract.TeeInstructionParamsStruct(&_TeeInstructions.TransactOpts, arg0)
}
