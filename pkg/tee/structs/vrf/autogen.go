// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf

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

// IVrfVrfInstructionMessage is an auto generated low-level Go binding around an user-defined struct.
type IVrfVrfInstructionMessage struct {
	WalletId [32]byte
	KeyId    uint64
	Nonce    []byte
}

// TeeVRFMetaData contains all meta data concerning the TeeVRF contract.
var TeeVRFMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"nonce\",\"type\":\"bytes\"}],\"internalType\":\"structIVrf.VrfInstructionMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"vrfInstructionMessageStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeVRFABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeVRFMetaData.ABI instead.
var TeeVRFABI = TeeVRFMetaData.ABI

// TeeVRF is an auto generated Go binding around an Ethereum contract.
type TeeVRF struct {
	TeeVRFCaller     // Read-only binding to the contract
	TeeVRFTransactor // Write-only binding to the contract
	TeeVRFFilterer   // Log filterer for contract events
}

// TeeVRFCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeVRFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVRFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeVRFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVRFFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeVRFFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVRFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeVRFSession struct {
	Contract     *TeeVRF           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeVRFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeVRFCallerSession struct {
	Contract *TeeVRFCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TeeVRFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeVRFTransactorSession struct {
	Contract     *TeeVRFTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeVRFRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeVRFRaw struct {
	Contract *TeeVRF // Generic contract binding to access the raw methods on
}

// TeeVRFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeVRFCallerRaw struct {
	Contract *TeeVRFCaller // Generic read-only contract binding to access the raw methods on
}

// TeeVRFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeVRFTransactorRaw struct {
	Contract *TeeVRFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeVRF creates a new instance of TeeVRF, bound to a specific deployed contract.
func NewTeeVRF(address common.Address, backend bind.ContractBackend) (*TeeVRF, error) {
	contract, err := bindTeeVRF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeVRF{TeeVRFCaller: TeeVRFCaller{contract: contract}, TeeVRFTransactor: TeeVRFTransactor{contract: contract}, TeeVRFFilterer: TeeVRFFilterer{contract: contract}}, nil
}

// NewTeeVRFCaller creates a new read-only instance of TeeVRF, bound to a specific deployed contract.
func NewTeeVRFCaller(address common.Address, caller bind.ContractCaller) (*TeeVRFCaller, error) {
	contract, err := bindTeeVRF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeVRFCaller{contract: contract}, nil
}

// NewTeeVRFTransactor creates a new write-only instance of TeeVRF, bound to a specific deployed contract.
func NewTeeVRFTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeVRFTransactor, error) {
	contract, err := bindTeeVRF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeVRFTransactor{contract: contract}, nil
}

// NewTeeVRFFilterer creates a new log filterer instance of TeeVRF, bound to a specific deployed contract.
func NewTeeVRFFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeVRFFilterer, error) {
	contract, err := bindTeeVRF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeVRFFilterer{contract: contract}, nil
}

// bindTeeVRF binds a generic wrapper to an already deployed contract.
func bindTeeVRF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeVRFMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeVRF *TeeVRFRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeVRF.Contract.TeeVRFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeVRF *TeeVRFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVRF.Contract.TeeVRFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeVRF *TeeVRFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeVRF.Contract.TeeVRFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeVRF *TeeVRFCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeVRF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeVRF *TeeVRFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVRF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeVRF *TeeVRFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeVRF.Contract.contract.Transact(opts, method, params...)
}

// VrfInstructionMessageStruct is a paid mutator transaction binding the contract method 0x9a320128.
//
// Solidity: function vrfInstructionMessageStruct((bytes32,uint64,bytes) ) returns()
func (_TeeVRF *TeeVRFTransactor) VrfInstructionMessageStruct(opts *bind.TransactOpts, arg0 IVrfVrfInstructionMessage) (*types.Transaction, error) {
	return _TeeVRF.contract.Transact(opts, "vrfInstructionMessageStruct", arg0)
}

// VrfInstructionMessageStruct is a paid mutator transaction binding the contract method 0x9a320128.
//
// Solidity: function vrfInstructionMessageStruct((bytes32,uint64,bytes) ) returns()
func (_TeeVRF *TeeVRFSession) VrfInstructionMessageStruct(arg0 IVrfVrfInstructionMessage) (*types.Transaction, error) {
	return _TeeVRF.Contract.VrfInstructionMessageStruct(&_TeeVRF.TransactOpts, arg0)
}

// VrfInstructionMessageStruct is a paid mutator transaction binding the contract method 0x9a320128.
//
// Solidity: function vrfInstructionMessageStruct((bytes32,uint64,bytes) ) returns()
func (_TeeVRF *TeeVRFTransactorSession) VrfInstructionMessageStruct(arg0 IVrfVrfInstructionMessage) (*types.Transaction, error) {
	return _TeeVRF.Contract.VrfInstructionMessageStruct(&_TeeVRF.TransactOpts, arg0)
}
