// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package upgrade

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
	OldTeeMachine IMachineManagerTeeMachineWithAttestationData
	NewTeeMachine IMachineManagerTeeMachineWithAttestationData
}

// IUpgradeManagerTeeNodeVersion is an auto generated low-level Go binding around an user-defined struct.
type IUpgradeManagerTeeNodeVersion struct {
	CodeHash [32]byte
	Platform [32]byte
}

// IUpgradeManagerTeeUpgradePath is an auto generated low-level Go binding around an user-defined struct.
type IUpgradeManagerTeeUpgradePath struct {
	SourceVersions []IUpgradeManagerTeeNodeVersion
	TargetVersions []IUpgradeManagerTeeNodeVersion
}

// TeeUpgradeMetaData contains all meta data concerning the TeeUpgrade contract.
var TeeUpgradeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"}],\"internalType\":\"structIReplication.PauseForUpgrade\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pauseForUpgradeStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIMachineManager.TeeMachineWithAttestationData\",\"name\":\"oldTeeMachine\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIMachineManager.TeeMachineWithAttestationData\",\"name\":\"newTeeMachine\",\"type\":\"tuple\"}],\"internalType\":\"structIReplication.ReplicateTeeMachine\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"replicateTeeMachineStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIUpgradeManager.TeeNodeVersion\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeNodeVersionStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIUpgradeManager.TeeNodeVersion[]\",\"name\":\"sourceVersions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIUpgradeManager.TeeNodeVersion[]\",\"name\":\"targetVersions\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUpgradeManager.TeeUpgradePath\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeUpgradePathStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeUpgradeABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeUpgradeMetaData.ABI instead.
var TeeUpgradeABI = TeeUpgradeMetaData.ABI

// TeeUpgrade is an auto generated Go binding around an Ethereum contract.
type TeeUpgrade struct {
	TeeUpgradeCaller     // Read-only binding to the contract
	TeeUpgradeTransactor // Write-only binding to the contract
	TeeUpgradeFilterer   // Log filterer for contract events
}

// TeeUpgradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeUpgradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeUpgradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeUpgradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeUpgradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeUpgradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeUpgradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeUpgradeSession struct {
	Contract     *TeeUpgrade       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeUpgradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeUpgradeCallerSession struct {
	Contract *TeeUpgradeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TeeUpgradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeUpgradeTransactorSession struct {
	Contract     *TeeUpgradeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TeeUpgradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeUpgradeRaw struct {
	Contract *TeeUpgrade // Generic contract binding to access the raw methods on
}

// TeeUpgradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeUpgradeCallerRaw struct {
	Contract *TeeUpgradeCaller // Generic read-only contract binding to access the raw methods on
}

// TeeUpgradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeUpgradeTransactorRaw struct {
	Contract *TeeUpgradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeUpgrade creates a new instance of TeeUpgrade, bound to a specific deployed contract.
func NewTeeUpgrade(address common.Address, backend bind.ContractBackend) (*TeeUpgrade, error) {
	contract, err := bindTeeUpgrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeUpgrade{TeeUpgradeCaller: TeeUpgradeCaller{contract: contract}, TeeUpgradeTransactor: TeeUpgradeTransactor{contract: contract}, TeeUpgradeFilterer: TeeUpgradeFilterer{contract: contract}}, nil
}

// NewTeeUpgradeCaller creates a new read-only instance of TeeUpgrade, bound to a specific deployed contract.
func NewTeeUpgradeCaller(address common.Address, caller bind.ContractCaller) (*TeeUpgradeCaller, error) {
	contract, err := bindTeeUpgrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeUpgradeCaller{contract: contract}, nil
}

// NewTeeUpgradeTransactor creates a new write-only instance of TeeUpgrade, bound to a specific deployed contract.
func NewTeeUpgradeTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeUpgradeTransactor, error) {
	contract, err := bindTeeUpgrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeUpgradeTransactor{contract: contract}, nil
}

// NewTeeUpgradeFilterer creates a new log filterer instance of TeeUpgrade, bound to a specific deployed contract.
func NewTeeUpgradeFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeUpgradeFilterer, error) {
	contract, err := bindTeeUpgrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeUpgradeFilterer{contract: contract}, nil
}

// bindTeeUpgrade binds a generic wrapper to an already deployed contract.
func bindTeeUpgrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeUpgradeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeUpgrade *TeeUpgradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeUpgrade.Contract.TeeUpgradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeUpgrade *TeeUpgradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeUpgrade.Contract.TeeUpgradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeUpgrade *TeeUpgradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeUpgrade.Contract.TeeUpgradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeUpgrade *TeeUpgradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeUpgrade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeUpgrade *TeeUpgradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeUpgrade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeUpgrade *TeeUpgradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeUpgrade.Contract.contract.Transact(opts, method, params...)
}

// PauseForUpgradeStruct is a paid mutator transaction binding the contract method 0xc1abcc81.
//
// Solidity: function pauseForUpgradeStruct((address,address) ) returns()
func (_TeeUpgrade *TeeUpgradeTransactor) PauseForUpgradeStruct(opts *bind.TransactOpts, arg0 IReplicationPauseForUpgrade) (*types.Transaction, error) {
	return _TeeUpgrade.contract.Transact(opts, "pauseForUpgradeStruct", arg0)
}

// PauseForUpgradeStruct is a paid mutator transaction binding the contract method 0xc1abcc81.
//
// Solidity: function pauseForUpgradeStruct((address,address) ) returns()
func (_TeeUpgrade *TeeUpgradeSession) PauseForUpgradeStruct(arg0 IReplicationPauseForUpgrade) (*types.Transaction, error) {
	return _TeeUpgrade.Contract.PauseForUpgradeStruct(&_TeeUpgrade.TransactOpts, arg0)
}

// PauseForUpgradeStruct is a paid mutator transaction binding the contract method 0xc1abcc81.
//
// Solidity: function pauseForUpgradeStruct((address,address) ) returns()
func (_TeeUpgrade *TeeUpgradeTransactorSession) PauseForUpgradeStruct(arg0 IReplicationPauseForUpgrade) (*types.Transaction, error) {
	return _TeeUpgrade.Contract.PauseForUpgradeStruct(&_TeeUpgrade.TransactOpts, arg0)
}

// ReplicateTeeMachineStruct is a paid mutator transaction binding the contract method 0xb6c61ed6.
//
// Solidity: function replicateTeeMachineStruct(((address,address,string,bytes32,bytes32),(address,address,string,bytes32,bytes32)) ) returns()
func (_TeeUpgrade *TeeUpgradeTransactor) ReplicateTeeMachineStruct(opts *bind.TransactOpts, arg0 IReplicationReplicateTeeMachine) (*types.Transaction, error) {
	return _TeeUpgrade.contract.Transact(opts, "replicateTeeMachineStruct", arg0)
}

// ReplicateTeeMachineStruct is a paid mutator transaction binding the contract method 0xb6c61ed6.
//
// Solidity: function replicateTeeMachineStruct(((address,address,string,bytes32,bytes32),(address,address,string,bytes32,bytes32)) ) returns()
func (_TeeUpgrade *TeeUpgradeSession) ReplicateTeeMachineStruct(arg0 IReplicationReplicateTeeMachine) (*types.Transaction, error) {
	return _TeeUpgrade.Contract.ReplicateTeeMachineStruct(&_TeeUpgrade.TransactOpts, arg0)
}

// ReplicateTeeMachineStruct is a paid mutator transaction binding the contract method 0xb6c61ed6.
//
// Solidity: function replicateTeeMachineStruct(((address,address,string,bytes32,bytes32),(address,address,string,bytes32,bytes32)) ) returns()
func (_TeeUpgrade *TeeUpgradeTransactorSession) ReplicateTeeMachineStruct(arg0 IReplicationReplicateTeeMachine) (*types.Transaction, error) {
	return _TeeUpgrade.Contract.ReplicateTeeMachineStruct(&_TeeUpgrade.TransactOpts, arg0)
}

// TeeNodeVersionStruct is a paid mutator transaction binding the contract method 0x609e259d.
//
// Solidity: function teeNodeVersionStruct((bytes32,bytes32) ) returns()
func (_TeeUpgrade *TeeUpgradeTransactor) TeeNodeVersionStruct(opts *bind.TransactOpts, arg0 IUpgradeManagerTeeNodeVersion) (*types.Transaction, error) {
	return _TeeUpgrade.contract.Transact(opts, "teeNodeVersionStruct", arg0)
}

// TeeNodeVersionStruct is a paid mutator transaction binding the contract method 0x609e259d.
//
// Solidity: function teeNodeVersionStruct((bytes32,bytes32) ) returns()
func (_TeeUpgrade *TeeUpgradeSession) TeeNodeVersionStruct(arg0 IUpgradeManagerTeeNodeVersion) (*types.Transaction, error) {
	return _TeeUpgrade.Contract.TeeNodeVersionStruct(&_TeeUpgrade.TransactOpts, arg0)
}

// TeeNodeVersionStruct is a paid mutator transaction binding the contract method 0x609e259d.
//
// Solidity: function teeNodeVersionStruct((bytes32,bytes32) ) returns()
func (_TeeUpgrade *TeeUpgradeTransactorSession) TeeNodeVersionStruct(arg0 IUpgradeManagerTeeNodeVersion) (*types.Transaction, error) {
	return _TeeUpgrade.Contract.TeeNodeVersionStruct(&_TeeUpgrade.TransactOpts, arg0)
}

// TeeUpgradePathStruct is a paid mutator transaction binding the contract method 0x470b8626.
//
// Solidity: function teeUpgradePathStruct(((bytes32,bytes32)[],(bytes32,bytes32)[]) ) returns()
func (_TeeUpgrade *TeeUpgradeTransactor) TeeUpgradePathStruct(opts *bind.TransactOpts, arg0 IUpgradeManagerTeeUpgradePath) (*types.Transaction, error) {
	return _TeeUpgrade.contract.Transact(opts, "teeUpgradePathStruct", arg0)
}

// TeeUpgradePathStruct is a paid mutator transaction binding the contract method 0x470b8626.
//
// Solidity: function teeUpgradePathStruct(((bytes32,bytes32)[],(bytes32,bytes32)[]) ) returns()
func (_TeeUpgrade *TeeUpgradeSession) TeeUpgradePathStruct(arg0 IUpgradeManagerTeeUpgradePath) (*types.Transaction, error) {
	return _TeeUpgrade.Contract.TeeUpgradePathStruct(&_TeeUpgrade.TransactOpts, arg0)
}

// TeeUpgradePathStruct is a paid mutator transaction binding the contract method 0x470b8626.
//
// Solidity: function teeUpgradePathStruct(((bytes32,bytes32)[],(bytes32,bytes32)[]) ) returns()
func (_TeeUpgrade *TeeUpgradeTransactorSession) TeeUpgradePathStruct(arg0 IUpgradeManagerTeeUpgradePath) (*types.Transaction, error) {
	return _TeeUpgrade.Contract.TeeUpgradePathStruct(&_TeeUpgrade.TransactOpts, arg0)
}
