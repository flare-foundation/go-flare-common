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

// ITeeVrfVrfInstructionMessage is an auto generated low-level Go binding around an user-defined struct.
type ITeeVrfVrfInstructionMessage struct {
	WalletId [32]byte
	KeyId    uint64
	Nonce    []byte
}

// VrfMetaData contains all meta data concerning the Vrf contract.
var VrfMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"nonce\",\"type\":\"bytes\"}],\"internalType\":\"structITeeVrf.VrfInstructionMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"vrfInstructionMessageStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VrfABI is the input ABI used to generate the binding from.
// Deprecated: Use VrfMetaData.ABI instead.
var VrfABI = VrfMetaData.ABI

// Vrf is an auto generated Go binding around an Ethereum contract.
type Vrf struct {
	VrfCaller     // Read-only binding to the contract
	VrfTransactor // Write-only binding to the contract
	VrfFilterer   // Log filterer for contract events
}

// VrfCaller is an auto generated read-only Go binding around an Ethereum contract.
type VrfCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VrfTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VrfTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VrfFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VrfFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VrfSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VrfSession struct {
	Contract     *Vrf              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VrfCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VrfCallerSession struct {
	Contract *VrfCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VrfTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VrfTransactorSession struct {
	Contract     *VrfTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VrfRaw is an auto generated low-level Go binding around an Ethereum contract.
type VrfRaw struct {
	Contract *Vrf // Generic contract binding to access the raw methods on
}

// VrfCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VrfCallerRaw struct {
	Contract *VrfCaller // Generic read-only contract binding to access the raw methods on
}

// VrfTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VrfTransactorRaw struct {
	Contract *VrfTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVrf creates a new instance of Vrf, bound to a specific deployed contract.
func NewVrf(address common.Address, backend bind.ContractBackend) (*Vrf, error) {
	contract, err := bindVrf(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vrf{VrfCaller: VrfCaller{contract: contract}, VrfTransactor: VrfTransactor{contract: contract}, VrfFilterer: VrfFilterer{contract: contract}}, nil
}

// NewVrfCaller creates a new read-only instance of Vrf, bound to a specific deployed contract.
func NewVrfCaller(address common.Address, caller bind.ContractCaller) (*VrfCaller, error) {
	contract, err := bindVrf(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VrfCaller{contract: contract}, nil
}

// NewVrfTransactor creates a new write-only instance of Vrf, bound to a specific deployed contract.
func NewVrfTransactor(address common.Address, transactor bind.ContractTransactor) (*VrfTransactor, error) {
	contract, err := bindVrf(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VrfTransactor{contract: contract}, nil
}

// NewVrfFilterer creates a new log filterer instance of Vrf, bound to a specific deployed contract.
func NewVrfFilterer(address common.Address, filterer bind.ContractFilterer) (*VrfFilterer, error) {
	contract, err := bindVrf(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VrfFilterer{contract: contract}, nil
}

// bindVrf binds a generic wrapper to an already deployed contract.
func bindVrf(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VrfMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vrf *VrfRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vrf.Contract.VrfCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vrf *VrfRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vrf.Contract.VrfTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vrf *VrfRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vrf.Contract.VrfTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vrf *VrfCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vrf.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vrf *VrfTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vrf.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vrf *VrfTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vrf.Contract.contract.Transact(opts, method, params...)
}

// VrfInstructionMessageStruct is a paid mutator transaction binding the contract method 0x9a320128.
//
// Solidity: function vrfInstructionMessageStruct((bytes32,uint64,bytes) ) returns()
func (_Vrf *VrfTransactor) VrfInstructionMessageStruct(opts *bind.TransactOpts, arg0 ITeeVrfVrfInstructionMessage) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "vrfInstructionMessageStruct", arg0)
}

// VrfInstructionMessageStruct is a paid mutator transaction binding the contract method 0x9a320128.
//
// Solidity: function vrfInstructionMessageStruct((bytes32,uint64,bytes) ) returns()
func (_Vrf *VrfSession) VrfInstructionMessageStruct(arg0 ITeeVrfVrfInstructionMessage) (*types.Transaction, error) {
	return _Vrf.Contract.VrfInstructionMessageStruct(&_Vrf.TransactOpts, arg0)
}

// VrfInstructionMessageStruct is a paid mutator transaction binding the contract method 0x9a320128.
//
// Solidity: function vrfInstructionMessageStruct((bytes32,uint64,bytes) ) returns()
func (_Vrf *VrfTransactorSession) VrfInstructionMessageStruct(arg0 ITeeVrfVrfInstructionMessage) (*types.Transaction, error) {
	return _Vrf.Contract.VrfInstructionMessageStruct(&_Vrf.TransactOpts, arg0)
}
