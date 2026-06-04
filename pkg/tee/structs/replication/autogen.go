// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package replication

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

// IMachineManagerTeeMachineWithAttestationData is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachineWithAttestationData struct {
	TeeId        common.Address
	InitialTeeId common.Address
	Url          string
	CodeHash     [32]byte
	Platform     [32]byte
}

// IReplicationPauseForUpgrade is an auto generated low-level Go binding around an user-defined struct.
type IReplicationPauseForUpgrade struct {
	TeeId        common.Address
	InitialTeeId common.Address
}

// IReplicationReplicateTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type IReplicationReplicateTeeMachine struct {
	OldTeeMachine        IMachineManagerTeeMachineWithAttestationData
	NewTeeMachine        IMachineManagerTeeMachineWithAttestationData
	MachinePathListNonce *big.Int
}

// TeeReplicationMetaData contains all meta data concerning the TeeReplication contract.
var TeeReplicationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"}],\"internalType\":\"structIReplication.PauseForUpgrade\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pauseForUpgradeStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIMachineManager.TeeMachineWithAttestationData\",\"name\":\"oldTeeMachine\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIMachineManager.TeeMachineWithAttestationData\",\"name\":\"newTeeMachine\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"machinePathListNonce\",\"type\":\"uint256\"}],\"internalType\":\"structIReplication.ReplicateTeeMachine\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"replicateTeeMachineStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeReplicationABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeReplicationMetaData.ABI instead.
var TeeReplicationABI = TeeReplicationMetaData.ABI

// TeeReplication is an auto generated Go binding around an Ethereum contract.
type TeeReplication struct {
	TeeReplicationCaller     // Read-only binding to the contract
	TeeReplicationTransactor // Write-only binding to the contract
	TeeReplicationFilterer   // Log filterer for contract events
}

// TeeReplicationCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeReplicationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeReplicationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeReplicationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeReplicationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeReplicationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeReplicationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeReplicationSession struct {
	Contract     *TeeReplication   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeReplicationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeReplicationCallerSession struct {
	Contract *TeeReplicationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TeeReplicationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeReplicationTransactorSession struct {
	Contract     *TeeReplicationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TeeReplicationRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeReplicationRaw struct {
	Contract *TeeReplication // Generic contract binding to access the raw methods on
}

// TeeReplicationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeReplicationCallerRaw struct {
	Contract *TeeReplicationCaller // Generic read-only contract binding to access the raw methods on
}

// TeeReplicationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeReplicationTransactorRaw struct {
	Contract *TeeReplicationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeReplication creates a new instance of TeeReplication, bound to a specific deployed contract.
func NewTeeReplication(address common.Address, backend bind.ContractBackend) (*TeeReplication, error) {
	contract, err := bindTeeReplication(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeReplication{TeeReplicationCaller: TeeReplicationCaller{contract: contract}, TeeReplicationTransactor: TeeReplicationTransactor{contract: contract}, TeeReplicationFilterer: TeeReplicationFilterer{contract: contract}}, nil
}

// NewTeeReplicationCaller creates a new read-only instance of TeeReplication, bound to a specific deployed contract.
func NewTeeReplicationCaller(address common.Address, caller bind.ContractCaller) (*TeeReplicationCaller, error) {
	contract, err := bindTeeReplication(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeReplicationCaller{contract: contract}, nil
}

// NewTeeReplicationTransactor creates a new write-only instance of TeeReplication, bound to a specific deployed contract.
func NewTeeReplicationTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeReplicationTransactor, error) {
	contract, err := bindTeeReplication(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeReplicationTransactor{contract: contract}, nil
}

// NewTeeReplicationFilterer creates a new log filterer instance of TeeReplication, bound to a specific deployed contract.
func NewTeeReplicationFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeReplicationFilterer, error) {
	contract, err := bindTeeReplication(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeReplicationFilterer{contract: contract}, nil
}

// bindTeeReplication binds a generic wrapper to an already deployed contract.
func bindTeeReplication(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeReplicationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeReplication *TeeReplicationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeReplication.Contract.TeeReplicationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeReplication *TeeReplicationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeReplication.Contract.TeeReplicationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeReplication *TeeReplicationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeReplication.Contract.TeeReplicationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeReplication *TeeReplicationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeReplication.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeReplication *TeeReplicationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeReplication.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeReplication *TeeReplicationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeReplication.Contract.contract.Transact(opts, method, params...)
}

// PauseForUpgradeStruct is a paid mutator transaction binding the contract method 0xc1abcc81.
//
// Solidity: function pauseForUpgradeStruct((address,address) ) returns()
func (_TeeReplication *TeeReplicationTransactor) PauseForUpgradeStruct(opts *bind.TransactOpts, arg0 IReplicationPauseForUpgrade) (*types.Transaction, error) {
	return _TeeReplication.contract.Transact(opts, "pauseForUpgradeStruct", arg0)
}

// PauseForUpgradeStruct is a paid mutator transaction binding the contract method 0xc1abcc81.
//
// Solidity: function pauseForUpgradeStruct((address,address) ) returns()
func (_TeeReplication *TeeReplicationSession) PauseForUpgradeStruct(arg0 IReplicationPauseForUpgrade) (*types.Transaction, error) {
	return _TeeReplication.Contract.PauseForUpgradeStruct(&_TeeReplication.TransactOpts, arg0)
}

// PauseForUpgradeStruct is a paid mutator transaction binding the contract method 0xc1abcc81.
//
// Solidity: function pauseForUpgradeStruct((address,address) ) returns()
func (_TeeReplication *TeeReplicationTransactorSession) PauseForUpgradeStruct(arg0 IReplicationPauseForUpgrade) (*types.Transaction, error) {
	return _TeeReplication.Contract.PauseForUpgradeStruct(&_TeeReplication.TransactOpts, arg0)
}

// ReplicateTeeMachineStruct is a paid mutator transaction binding the contract method 0x22890d2d.
//
// Solidity: function replicateTeeMachineStruct(((address,address,string,bytes32,bytes32),(address,address,string,bytes32,bytes32),uint256) ) returns()
func (_TeeReplication *TeeReplicationTransactor) ReplicateTeeMachineStruct(opts *bind.TransactOpts, arg0 IReplicationReplicateTeeMachine) (*types.Transaction, error) {
	return _TeeReplication.contract.Transact(opts, "replicateTeeMachineStruct", arg0)
}

// ReplicateTeeMachineStruct is a paid mutator transaction binding the contract method 0x22890d2d.
//
// Solidity: function replicateTeeMachineStruct(((address,address,string,bytes32,bytes32),(address,address,string,bytes32,bytes32),uint256) ) returns()
func (_TeeReplication *TeeReplicationSession) ReplicateTeeMachineStruct(arg0 IReplicationReplicateTeeMachine) (*types.Transaction, error) {
	return _TeeReplication.Contract.ReplicateTeeMachineStruct(&_TeeReplication.TransactOpts, arg0)
}

// ReplicateTeeMachineStruct is a paid mutator transaction binding the contract method 0x22890d2d.
//
// Solidity: function replicateTeeMachineStruct(((address,address,string,bytes32,bytes32),(address,address,string,bytes32,bytes32),uint256) ) returns()
func (_TeeReplication *TeeReplicationTransactorSession) ReplicateTeeMachineStruct(arg0 IReplicationReplicateTeeMachine) (*types.Transaction, error) {
	return _TeeReplication.Contract.ReplicateTeeMachineStruct(&_TeeReplication.TransactOpts, arg0)
}
