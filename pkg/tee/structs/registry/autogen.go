// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// ITeeRegistryAvailabilityCheckRequest is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryAvailabilityCheckRequest struct {
	TeeId     common.Address
	Url       string
	CodeHash  [32]byte
	Platform  [32]byte
	Timestamp *big.Int
}

// ITeeRegistryAvailabilityCheckResponse is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryAvailabilityCheckResponse struct {
	TeeId     common.Address
	Url       string
	CodeHash  [32]byte
	Platform  [32]byte
	Timestamp *big.Int
	Status    uint8
}

// ITeeRegistryPauseForUpgrade is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryPauseForUpgrade struct {
	TeeId common.Address
}

// ITeeRegistryReplicateTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryReplicateTeeMachine struct {
	OldTeeMachine ITeeRegistryTeeMachineWithAttestationData
	NewTeeMachine ITeeRegistryTeeMachineWithAttestationData
}

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

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structITeeRegistry.AvailabilityCheckRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckRequestStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumITeeRegistry.AvailabilityStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structITeeRegistry.AvailabilityCheckResponse\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckResponseStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"internalType\":\"structITeeRegistry.PauseForUpgrade\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pauseForUpgradeStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"oldTeeMachine\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"newTeeMachine\",\"type\":\"tuple\"}],\"internalType\":\"structITeeRegistry.ReplicateTeeMachine\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"replicateTeeMachineStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structITeeRegistry.TeeMachine\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeMachineStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeMachineWithAttestationDataStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// AvailabilityCheckRequestStruct is a paid mutator transaction binding the contract method 0xa0f09cd4.
//
// Solidity: function availabilityCheckRequestStruct((address,string,bytes32,bytes32,uint256) ) returns()
func (_Registry *RegistryTransactor) AvailabilityCheckRequestStruct(opts *bind.TransactOpts, arg0 ITeeRegistryAvailabilityCheckRequest) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "availabilityCheckRequestStruct", arg0)
}

// AvailabilityCheckRequestStruct is a paid mutator transaction binding the contract method 0xa0f09cd4.
//
// Solidity: function availabilityCheckRequestStruct((address,string,bytes32,bytes32,uint256) ) returns()
func (_Registry *RegistrySession) AvailabilityCheckRequestStruct(arg0 ITeeRegistryAvailabilityCheckRequest) (*types.Transaction, error) {
	return _Registry.Contract.AvailabilityCheckRequestStruct(&_Registry.TransactOpts, arg0)
}

// AvailabilityCheckRequestStruct is a paid mutator transaction binding the contract method 0xa0f09cd4.
//
// Solidity: function availabilityCheckRequestStruct((address,string,bytes32,bytes32,uint256) ) returns()
func (_Registry *RegistryTransactorSession) AvailabilityCheckRequestStruct(arg0 ITeeRegistryAvailabilityCheckRequest) (*types.Transaction, error) {
	return _Registry.Contract.AvailabilityCheckRequestStruct(&_Registry.TransactOpts, arg0)
}

// AvailabilityCheckResponseStruct is a paid mutator transaction binding the contract method 0x8b727171.
//
// Solidity: function availabilityCheckResponseStruct((address,string,bytes32,bytes32,uint256,uint8) ) returns()
func (_Registry *RegistryTransactor) AvailabilityCheckResponseStruct(opts *bind.TransactOpts, arg0 ITeeRegistryAvailabilityCheckResponse) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "availabilityCheckResponseStruct", arg0)
}

// AvailabilityCheckResponseStruct is a paid mutator transaction binding the contract method 0x8b727171.
//
// Solidity: function availabilityCheckResponseStruct((address,string,bytes32,bytes32,uint256,uint8) ) returns()
func (_Registry *RegistrySession) AvailabilityCheckResponseStruct(arg0 ITeeRegistryAvailabilityCheckResponse) (*types.Transaction, error) {
	return _Registry.Contract.AvailabilityCheckResponseStruct(&_Registry.TransactOpts, arg0)
}

// AvailabilityCheckResponseStruct is a paid mutator transaction binding the contract method 0x8b727171.
//
// Solidity: function availabilityCheckResponseStruct((address,string,bytes32,bytes32,uint256,uint8) ) returns()
func (_Registry *RegistryTransactorSession) AvailabilityCheckResponseStruct(arg0 ITeeRegistryAvailabilityCheckResponse) (*types.Transaction, error) {
	return _Registry.Contract.AvailabilityCheckResponseStruct(&_Registry.TransactOpts, arg0)
}

// PauseForUpgradeStruct is a paid mutator transaction binding the contract method 0x880347fb.
//
// Solidity: function pauseForUpgradeStruct((address) ) returns()
func (_Registry *RegistryTransactor) PauseForUpgradeStruct(opts *bind.TransactOpts, arg0 ITeeRegistryPauseForUpgrade) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "pauseForUpgradeStruct", arg0)
}

// PauseForUpgradeStruct is a paid mutator transaction binding the contract method 0x880347fb.
//
// Solidity: function pauseForUpgradeStruct((address) ) returns()
func (_Registry *RegistrySession) PauseForUpgradeStruct(arg0 ITeeRegistryPauseForUpgrade) (*types.Transaction, error) {
	return _Registry.Contract.PauseForUpgradeStruct(&_Registry.TransactOpts, arg0)
}

// PauseForUpgradeStruct is a paid mutator transaction binding the contract method 0x880347fb.
//
// Solidity: function pauseForUpgradeStruct((address) ) returns()
func (_Registry *RegistryTransactorSession) PauseForUpgradeStruct(arg0 ITeeRegistryPauseForUpgrade) (*types.Transaction, error) {
	return _Registry.Contract.PauseForUpgradeStruct(&_Registry.TransactOpts, arg0)
}

// ReplicateTeeMachineStruct is a paid mutator transaction binding the contract method 0xb6c61ed6.
//
// Solidity: function replicateTeeMachineStruct(((address,address,string,bytes32,bytes32),(address,address,string,bytes32,bytes32)) ) returns()
func (_Registry *RegistryTransactor) ReplicateTeeMachineStruct(opts *bind.TransactOpts, arg0 ITeeRegistryReplicateTeeMachine) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "replicateTeeMachineStruct", arg0)
}

// ReplicateTeeMachineStruct is a paid mutator transaction binding the contract method 0xb6c61ed6.
//
// Solidity: function replicateTeeMachineStruct(((address,address,string,bytes32,bytes32),(address,address,string,bytes32,bytes32)) ) returns()
func (_Registry *RegistrySession) ReplicateTeeMachineStruct(arg0 ITeeRegistryReplicateTeeMachine) (*types.Transaction, error) {
	return _Registry.Contract.ReplicateTeeMachineStruct(&_Registry.TransactOpts, arg0)
}

// ReplicateTeeMachineStruct is a paid mutator transaction binding the contract method 0xb6c61ed6.
//
// Solidity: function replicateTeeMachineStruct(((address,address,string,bytes32,bytes32),(address,address,string,bytes32,bytes32)) ) returns()
func (_Registry *RegistryTransactorSession) ReplicateTeeMachineStruct(arg0 ITeeRegistryReplicateTeeMachine) (*types.Transaction, error) {
	return _Registry.Contract.ReplicateTeeMachineStruct(&_Registry.TransactOpts, arg0)
}

// TeeMachineStruct is a paid mutator transaction binding the contract method 0xef60cfee.
//
// Solidity: function teeMachineStruct((address,address,string) ) returns()
func (_Registry *RegistryTransactor) TeeMachineStruct(opts *bind.TransactOpts, arg0 ITeeRegistryTeeMachine) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "teeMachineStruct", arg0)
}

// TeeMachineStruct is a paid mutator transaction binding the contract method 0xef60cfee.
//
// Solidity: function teeMachineStruct((address,address,string) ) returns()
func (_Registry *RegistrySession) TeeMachineStruct(arg0 ITeeRegistryTeeMachine) (*types.Transaction, error) {
	return _Registry.Contract.TeeMachineStruct(&_Registry.TransactOpts, arg0)
}

// TeeMachineStruct is a paid mutator transaction binding the contract method 0xef60cfee.
//
// Solidity: function teeMachineStruct((address,address,string) ) returns()
func (_Registry *RegistryTransactorSession) TeeMachineStruct(arg0 ITeeRegistryTeeMachine) (*types.Transaction, error) {
	return _Registry.Contract.TeeMachineStruct(&_Registry.TransactOpts, arg0)
}

// TeeMachineWithAttestationDataStruct is a paid mutator transaction binding the contract method 0xd77f63b9.
//
// Solidity: function teeMachineWithAttestationDataStruct((address,address,string,bytes32,bytes32) ) returns()
func (_Registry *RegistryTransactor) TeeMachineWithAttestationDataStruct(opts *bind.TransactOpts, arg0 ITeeRegistryTeeMachineWithAttestationData) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "teeMachineWithAttestationDataStruct", arg0)
}

// TeeMachineWithAttestationDataStruct is a paid mutator transaction binding the contract method 0xd77f63b9.
//
// Solidity: function teeMachineWithAttestationDataStruct((address,address,string,bytes32,bytes32) ) returns()
func (_Registry *RegistrySession) TeeMachineWithAttestationDataStruct(arg0 ITeeRegistryTeeMachineWithAttestationData) (*types.Transaction, error) {
	return _Registry.Contract.TeeMachineWithAttestationDataStruct(&_Registry.TransactOpts, arg0)
}

// TeeMachineWithAttestationDataStruct is a paid mutator transaction binding the contract method 0xd77f63b9.
//
// Solidity: function teeMachineWithAttestationDataStruct((address,address,string,bytes32,bytes32) ) returns()
func (_Registry *RegistryTransactorSession) TeeMachineWithAttestationDataStruct(arg0 ITeeRegistryTeeMachineWithAttestationData) (*types.Transaction, error) {
	return _Registry.Contract.TeeMachineWithAttestationDataStruct(&_Registry.TransactOpts, arg0)
}
