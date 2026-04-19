// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verification

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

// IMachineManagerFacetTeeMachineWithAttestationData is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerFacetTeeMachineWithAttestationData struct {
	TeeId        common.Address
	InitialTeeId common.Address
	Url          string
	CodeHash     [32]byte
	Platform     [32]byte
}

// ISystemStateVerifierFacetTeeSystemState is an auto generated low-level Go binding around an user-defined struct.
type ISystemStateVerifierFacetTeeSystemState struct {
	Status            uint8
	InitialTeeId      common.Address
	TeeGovernanceHash [32]byte
}

// IVerificationFacetTeeAttestation is an auto generated low-level Go binding around an user-defined struct.
type IVerificationFacetTeeAttestation struct {
	TeeMachine IMachineManagerFacetTeeMachineWithAttestationData
	Challenge  [32]byte
}

// VerificationMetaData contains all meta data concerning the Verification contract.
var VerificationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIMachineManagerFacet.TeeMachineWithAttestationData\",\"name\":\"teeMachine\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"internalType\":\"structIVerificationFacet.TeeAttestation\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeAttestationStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumISystemStateVerifierFacet.TeeMachineStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"teeGovernanceHash\",\"type\":\"bytes32\"}],\"internalType\":\"structISystemStateVerifierFacet.TeeSystemState\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeSystemStateStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VerificationABI is the input ABI used to generate the binding from.
// Deprecated: Use VerificationMetaData.ABI instead.
var VerificationABI = VerificationMetaData.ABI

// Verification is an auto generated Go binding around an Ethereum contract.
type Verification struct {
	VerificationCaller     // Read-only binding to the contract
	VerificationTransactor // Write-only binding to the contract
	VerificationFilterer   // Log filterer for contract events
}

// VerificationCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerificationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerificationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerificationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerificationSession struct {
	Contract     *Verification     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerificationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerificationCallerSession struct {
	Contract *VerificationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VerificationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerificationTransactorSession struct {
	Contract     *VerificationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VerificationRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerificationRaw struct {
	Contract *Verification // Generic contract binding to access the raw methods on
}

// VerificationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerificationCallerRaw struct {
	Contract *VerificationCaller // Generic read-only contract binding to access the raw methods on
}

// VerificationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerificationTransactorRaw struct {
	Contract *VerificationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerification creates a new instance of Verification, bound to a specific deployed contract.
func NewVerification(address common.Address, backend bind.ContractBackend) (*Verification, error) {
	contract, err := bindVerification(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verification{VerificationCaller: VerificationCaller{contract: contract}, VerificationTransactor: VerificationTransactor{contract: contract}, VerificationFilterer: VerificationFilterer{contract: contract}}, nil
}

// NewVerificationCaller creates a new read-only instance of Verification, bound to a specific deployed contract.
func NewVerificationCaller(address common.Address, caller bind.ContractCaller) (*VerificationCaller, error) {
	contract, err := bindVerification(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationCaller{contract: contract}, nil
}

// NewVerificationTransactor creates a new write-only instance of Verification, bound to a specific deployed contract.
func NewVerificationTransactor(address common.Address, transactor bind.ContractTransactor) (*VerificationTransactor, error) {
	contract, err := bindVerification(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationTransactor{contract: contract}, nil
}

// NewVerificationFilterer creates a new log filterer instance of Verification, bound to a specific deployed contract.
func NewVerificationFilterer(address common.Address, filterer bind.ContractFilterer) (*VerificationFilterer, error) {
	contract, err := bindVerification(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerificationFilterer{contract: contract}, nil
}

// bindVerification binds a generic wrapper to an already deployed contract.
func bindVerification(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerificationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verification *VerificationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verification.Contract.VerificationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verification *VerificationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verification.Contract.VerificationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verification *VerificationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verification.Contract.VerificationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verification *VerificationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verification.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verification *VerificationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verification.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verification *VerificationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verification.Contract.contract.Transact(opts, method, params...)
}

// TeeAttestationStruct is a paid mutator transaction binding the contract method 0x050d0fc8.
//
// Solidity: function teeAttestationStruct(((address,address,string,bytes32,bytes32),bytes32) ) returns()
func (_Verification *VerificationTransactor) TeeAttestationStruct(opts *bind.TransactOpts, arg0 IVerificationFacetTeeAttestation) (*types.Transaction, error) {
	return _Verification.contract.Transact(opts, "teeAttestationStruct", arg0)
}

// TeeAttestationStruct is a paid mutator transaction binding the contract method 0x050d0fc8.
//
// Solidity: function teeAttestationStruct(((address,address,string,bytes32,bytes32),bytes32) ) returns()
func (_Verification *VerificationSession) TeeAttestationStruct(arg0 IVerificationFacetTeeAttestation) (*types.Transaction, error) {
	return _Verification.Contract.TeeAttestationStruct(&_Verification.TransactOpts, arg0)
}

// TeeAttestationStruct is a paid mutator transaction binding the contract method 0x050d0fc8.
//
// Solidity: function teeAttestationStruct(((address,address,string,bytes32,bytes32),bytes32) ) returns()
func (_Verification *VerificationTransactorSession) TeeAttestationStruct(arg0 IVerificationFacetTeeAttestation) (*types.Transaction, error) {
	return _Verification.Contract.TeeAttestationStruct(&_Verification.TransactOpts, arg0)
}

// TeeSystemStateStruct is a paid mutator transaction binding the contract method 0xc7115680.
//
// Solidity: function teeSystemStateStruct((uint8,address,bytes32) ) returns()
func (_Verification *VerificationTransactor) TeeSystemStateStruct(opts *bind.TransactOpts, arg0 ISystemStateVerifierFacetTeeSystemState) (*types.Transaction, error) {
	return _Verification.contract.Transact(opts, "teeSystemStateStruct", arg0)
}

// TeeSystemStateStruct is a paid mutator transaction binding the contract method 0xc7115680.
//
// Solidity: function teeSystemStateStruct((uint8,address,bytes32) ) returns()
func (_Verification *VerificationSession) TeeSystemStateStruct(arg0 ISystemStateVerifierFacetTeeSystemState) (*types.Transaction, error) {
	return _Verification.Contract.TeeSystemStateStruct(&_Verification.TransactOpts, arg0)
}

// TeeSystemStateStruct is a paid mutator transaction binding the contract method 0xc7115680.
//
// Solidity: function teeSystemStateStruct((uint8,address,bytes32) ) returns()
func (_Verification *VerificationTransactorSession) TeeSystemStateStruct(arg0 ISystemStateVerifierFacetTeeSystemState) (*types.Transaction, error) {
	return _Verification.Contract.TeeSystemStateStruct(&_Verification.TransactOpts, arg0)
}
