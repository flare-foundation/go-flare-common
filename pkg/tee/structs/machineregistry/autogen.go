// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package machineregistry

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

// ITeeMachineRegistryTeeMachineWithAttestationData is an auto generated low-level Go binding around an user-defined struct.
type ITeeMachineRegistryTeeMachineWithAttestationData struct {
	TeeId        common.Address
	InitialTeeId common.Address
	Url          string
	CodeHash     [32]byte
	Platform     [32]byte
}

// ITeeReplicationPauseForUpgrade is an auto generated low-level Go binding around an user-defined struct.
type ITeeReplicationPauseForUpgrade struct {
	TeeId        common.Address
	InitialTeeId common.Address
}

// ITeeReplicationReplicateTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type ITeeReplicationReplicateTeeMachine struct {
	OldTeeMachine ITeeMachineRegistryTeeMachineWithAttestationData
	NewTeeMachine ITeeMachineRegistryTeeMachineWithAttestationData
}

// MachineRegistryMetaData contains all meta data concerning the MachineRegistry contract.
var MachineRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"}],\"internalType\":\"structITeeReplication.PauseForUpgrade\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pauseForUpgradeStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeMachineRegistry.TeeMachineWithAttestationData\",\"name\":\"oldTeeMachine\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeMachineRegistry.TeeMachineWithAttestationData\",\"name\":\"newTeeMachine\",\"type\":\"tuple\"}],\"internalType\":\"structITeeReplication.ReplicateTeeMachine\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"replicateTeeMachineStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structITeeMachineRegistry.TeeMachine\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeMachineStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeMachineRegistry.TeeMachineWithAttestationData\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeMachineWithAttestationDataStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MachineRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MachineRegistryMetaData.ABI instead.
var MachineRegistryABI = MachineRegistryMetaData.ABI

// MachineRegistry is an auto generated Go binding around an Ethereum contract.
type MachineRegistry struct {
	MachineRegistryCaller     // Read-only binding to the contract
	MachineRegistryTransactor // Write-only binding to the contract
	MachineRegistryFilterer   // Log filterer for contract events
}

// MachineRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MachineRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MachineRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MachineRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MachineRegistrySession struct {
	Contract     *MachineRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MachineRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MachineRegistryCallerSession struct {
	Contract *MachineRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MachineRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MachineRegistryTransactorSession struct {
	Contract     *MachineRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MachineRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MachineRegistryRaw struct {
	Contract *MachineRegistry // Generic contract binding to access the raw methods on
}

// MachineRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MachineRegistryCallerRaw struct {
	Contract *MachineRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MachineRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MachineRegistryTransactorRaw struct {
	Contract *MachineRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMachineRegistry creates a new instance of MachineRegistry, bound to a specific deployed contract.
func NewMachineRegistry(address common.Address, backend bind.ContractBackend) (*MachineRegistry, error) {
	contract, err := bindMachineRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MachineRegistry{MachineRegistryCaller: MachineRegistryCaller{contract: contract}, MachineRegistryTransactor: MachineRegistryTransactor{contract: contract}, MachineRegistryFilterer: MachineRegistryFilterer{contract: contract}}, nil
}

// NewMachineRegistryCaller creates a new read-only instance of MachineRegistry, bound to a specific deployed contract.
func NewMachineRegistryCaller(address common.Address, caller bind.ContractCaller) (*MachineRegistryCaller, error) {
	contract, err := bindMachineRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MachineRegistryCaller{contract: contract}, nil
}

// NewMachineRegistryTransactor creates a new write-only instance of MachineRegistry, bound to a specific deployed contract.
func NewMachineRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MachineRegistryTransactor, error) {
	contract, err := bindMachineRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MachineRegistryTransactor{contract: contract}, nil
}

// NewMachineRegistryFilterer creates a new log filterer instance of MachineRegistry, bound to a specific deployed contract.
func NewMachineRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MachineRegistryFilterer, error) {
	contract, err := bindMachineRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MachineRegistryFilterer{contract: contract}, nil
}

// bindMachineRegistry binds a generic wrapper to an already deployed contract.
func bindMachineRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MachineRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MachineRegistry *MachineRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MachineRegistry.Contract.MachineRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MachineRegistry *MachineRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MachineRegistry.Contract.MachineRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MachineRegistry *MachineRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MachineRegistry.Contract.MachineRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MachineRegistry *MachineRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MachineRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MachineRegistry *MachineRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MachineRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MachineRegistry *MachineRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MachineRegistry.Contract.contract.Transact(opts, method, params...)
}

// PauseForUpgradeStruct is a paid mutator transaction binding the contract method 0xc1abcc81.
//
// Solidity: function pauseForUpgradeStruct((address,address) ) returns()
func (_MachineRegistry *MachineRegistryTransactor) PauseForUpgradeStruct(opts *bind.TransactOpts, arg0 ITeeReplicationPauseForUpgrade) (*types.Transaction, error) {
	return _MachineRegistry.contract.Transact(opts, "pauseForUpgradeStruct", arg0)
}

// PauseForUpgradeStruct is a paid mutator transaction binding the contract method 0xc1abcc81.
//
// Solidity: function pauseForUpgradeStruct((address,address) ) returns()
func (_MachineRegistry *MachineRegistrySession) PauseForUpgradeStruct(arg0 ITeeReplicationPauseForUpgrade) (*types.Transaction, error) {
	return _MachineRegistry.Contract.PauseForUpgradeStruct(&_MachineRegistry.TransactOpts, arg0)
}

// PauseForUpgradeStruct is a paid mutator transaction binding the contract method 0xc1abcc81.
//
// Solidity: function pauseForUpgradeStruct((address,address) ) returns()
func (_MachineRegistry *MachineRegistryTransactorSession) PauseForUpgradeStruct(arg0 ITeeReplicationPauseForUpgrade) (*types.Transaction, error) {
	return _MachineRegistry.Contract.PauseForUpgradeStruct(&_MachineRegistry.TransactOpts, arg0)
}

// ReplicateTeeMachineStruct is a paid mutator transaction binding the contract method 0xb6c61ed6.
//
// Solidity: function replicateTeeMachineStruct(((address,address,string,bytes32,bytes32),(address,address,string,bytes32,bytes32)) ) returns()
func (_MachineRegistry *MachineRegistryTransactor) ReplicateTeeMachineStruct(opts *bind.TransactOpts, arg0 ITeeReplicationReplicateTeeMachine) (*types.Transaction, error) {
	return _MachineRegistry.contract.Transact(opts, "replicateTeeMachineStruct", arg0)
}

// ReplicateTeeMachineStruct is a paid mutator transaction binding the contract method 0xb6c61ed6.
//
// Solidity: function replicateTeeMachineStruct(((address,address,string,bytes32,bytes32),(address,address,string,bytes32,bytes32)) ) returns()
func (_MachineRegistry *MachineRegistrySession) ReplicateTeeMachineStruct(arg0 ITeeReplicationReplicateTeeMachine) (*types.Transaction, error) {
	return _MachineRegistry.Contract.ReplicateTeeMachineStruct(&_MachineRegistry.TransactOpts, arg0)
}

// ReplicateTeeMachineStruct is a paid mutator transaction binding the contract method 0xb6c61ed6.
//
// Solidity: function replicateTeeMachineStruct(((address,address,string,bytes32,bytes32),(address,address,string,bytes32,bytes32)) ) returns()
func (_MachineRegistry *MachineRegistryTransactorSession) ReplicateTeeMachineStruct(arg0 ITeeReplicationReplicateTeeMachine) (*types.Transaction, error) {
	return _MachineRegistry.Contract.ReplicateTeeMachineStruct(&_MachineRegistry.TransactOpts, arg0)
}

// TeeMachineStruct is a paid mutator transaction binding the contract method 0xef60cfee.
//
// Solidity: function teeMachineStruct((address,address,string) ) returns()
func (_MachineRegistry *MachineRegistryTransactor) TeeMachineStruct(opts *bind.TransactOpts, arg0 ITeeMachineRegistryTeeMachine) (*types.Transaction, error) {
	return _MachineRegistry.contract.Transact(opts, "teeMachineStruct", arg0)
}

// TeeMachineStruct is a paid mutator transaction binding the contract method 0xef60cfee.
//
// Solidity: function teeMachineStruct((address,address,string) ) returns()
func (_MachineRegistry *MachineRegistrySession) TeeMachineStruct(arg0 ITeeMachineRegistryTeeMachine) (*types.Transaction, error) {
	return _MachineRegistry.Contract.TeeMachineStruct(&_MachineRegistry.TransactOpts, arg0)
}

// TeeMachineStruct is a paid mutator transaction binding the contract method 0xef60cfee.
//
// Solidity: function teeMachineStruct((address,address,string) ) returns()
func (_MachineRegistry *MachineRegistryTransactorSession) TeeMachineStruct(arg0 ITeeMachineRegistryTeeMachine) (*types.Transaction, error) {
	return _MachineRegistry.Contract.TeeMachineStruct(&_MachineRegistry.TransactOpts, arg0)
}

// TeeMachineWithAttestationDataStruct is a paid mutator transaction binding the contract method 0xd77f63b9.
//
// Solidity: function teeMachineWithAttestationDataStruct((address,address,string,bytes32,bytes32) ) returns()
func (_MachineRegistry *MachineRegistryTransactor) TeeMachineWithAttestationDataStruct(opts *bind.TransactOpts, arg0 ITeeMachineRegistryTeeMachineWithAttestationData) (*types.Transaction, error) {
	return _MachineRegistry.contract.Transact(opts, "teeMachineWithAttestationDataStruct", arg0)
}

// TeeMachineWithAttestationDataStruct is a paid mutator transaction binding the contract method 0xd77f63b9.
//
// Solidity: function teeMachineWithAttestationDataStruct((address,address,string,bytes32,bytes32) ) returns()
func (_MachineRegistry *MachineRegistrySession) TeeMachineWithAttestationDataStruct(arg0 ITeeMachineRegistryTeeMachineWithAttestationData) (*types.Transaction, error) {
	return _MachineRegistry.Contract.TeeMachineWithAttestationDataStruct(&_MachineRegistry.TransactOpts, arg0)
}

// TeeMachineWithAttestationDataStruct is a paid mutator transaction binding the contract method 0xd77f63b9.
//
// Solidity: function teeMachineWithAttestationDataStruct((address,address,string,bytes32,bytes32) ) returns()
func (_MachineRegistry *MachineRegistryTransactorSession) TeeMachineWithAttestationDataStruct(arg0 ITeeMachineRegistryTeeMachineWithAttestationData) (*types.Transaction, error) {
	return _MachineRegistry.Contract.TeeMachineWithAttestationDataStruct(&_MachineRegistry.TransactOpts, arg0)
}
