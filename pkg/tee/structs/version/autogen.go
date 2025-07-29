// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package version

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

// ITeeVersionManagerTeeNodeVersion is an auto generated low-level Go binding around an user-defined struct.
type ITeeVersionManagerTeeNodeVersion struct {
	CodeHash [32]byte
	Platform [32]byte
}

// ITeeVersionManagerTeeUpgradePath is an auto generated low-level Go binding around an user-defined struct.
type ITeeVersionManagerTeeUpgradePath struct {
	SourceVersions []ITeeVersionManagerTeeNodeVersion
	TargetVersions []ITeeVersionManagerTeeNodeVersion
}

// VersionMetaData contains all meta data concerning the Version contract.
var VersionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeVersionManager.TeeNodeVersion\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeNodeVersionStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeVersionManager.TeeNodeVersion[]\",\"name\":\"sourceVersions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeVersionManager.TeeNodeVersion[]\",\"name\":\"targetVersions\",\"type\":\"tuple[]\"}],\"internalType\":\"structITeeVersionManager.TeeUpgradePath\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"teeUpgradePathStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VersionABI is the input ABI used to generate the binding from.
// Deprecated: Use VersionMetaData.ABI instead.
var VersionABI = VersionMetaData.ABI

// Version is an auto generated Go binding around an Ethereum contract.
type Version struct {
	VersionCaller     // Read-only binding to the contract
	VersionTransactor // Write-only binding to the contract
	VersionFilterer   // Log filterer for contract events
}

// VersionCaller is an auto generated read-only Go binding around an Ethereum contract.
type VersionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VersionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VersionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VersionSession struct {
	Contract     *Version          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VersionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VersionCallerSession struct {
	Contract *VersionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VersionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VersionTransactorSession struct {
	Contract     *VersionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VersionRaw is an auto generated low-level Go binding around an Ethereum contract.
type VersionRaw struct {
	Contract *Version // Generic contract binding to access the raw methods on
}

// VersionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VersionCallerRaw struct {
	Contract *VersionCaller // Generic read-only contract binding to access the raw methods on
}

// VersionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VersionTransactorRaw struct {
	Contract *VersionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVersion creates a new instance of Version, bound to a specific deployed contract.
func NewVersion(address common.Address, backend bind.ContractBackend) (*Version, error) {
	contract, err := bindVersion(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Version{VersionCaller: VersionCaller{contract: contract}, VersionTransactor: VersionTransactor{contract: contract}, VersionFilterer: VersionFilterer{contract: contract}}, nil
}

// NewVersionCaller creates a new read-only instance of Version, bound to a specific deployed contract.
func NewVersionCaller(address common.Address, caller bind.ContractCaller) (*VersionCaller, error) {
	contract, err := bindVersion(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VersionCaller{contract: contract}, nil
}

// NewVersionTransactor creates a new write-only instance of Version, bound to a specific deployed contract.
func NewVersionTransactor(address common.Address, transactor bind.ContractTransactor) (*VersionTransactor, error) {
	contract, err := bindVersion(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VersionTransactor{contract: contract}, nil
}

// NewVersionFilterer creates a new log filterer instance of Version, bound to a specific deployed contract.
func NewVersionFilterer(address common.Address, filterer bind.ContractFilterer) (*VersionFilterer, error) {
	contract, err := bindVersion(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VersionFilterer{contract: contract}, nil
}

// bindVersion binds a generic wrapper to an already deployed contract.
func bindVersion(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VersionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Version *VersionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Version.Contract.VersionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Version *VersionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Version.Contract.VersionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Version *VersionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Version.Contract.VersionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Version *VersionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Version.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Version *VersionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Version.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Version *VersionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Version.Contract.contract.Transact(opts, method, params...)
}

// TeeNodeVersionStruct is a paid mutator transaction binding the contract method 0x609e259d.
//
// Solidity: function teeNodeVersionStruct((bytes32,bytes32) ) returns()
func (_Version *VersionTransactor) TeeNodeVersionStruct(opts *bind.TransactOpts, arg0 ITeeVersionManagerTeeNodeVersion) (*types.Transaction, error) {
	return _Version.contract.Transact(opts, "teeNodeVersionStruct", arg0)
}

// TeeNodeVersionStruct is a paid mutator transaction binding the contract method 0x609e259d.
//
// Solidity: function teeNodeVersionStruct((bytes32,bytes32) ) returns()
func (_Version *VersionSession) TeeNodeVersionStruct(arg0 ITeeVersionManagerTeeNodeVersion) (*types.Transaction, error) {
	return _Version.Contract.TeeNodeVersionStruct(&_Version.TransactOpts, arg0)
}

// TeeNodeVersionStruct is a paid mutator transaction binding the contract method 0x609e259d.
//
// Solidity: function teeNodeVersionStruct((bytes32,bytes32) ) returns()
func (_Version *VersionTransactorSession) TeeNodeVersionStruct(arg0 ITeeVersionManagerTeeNodeVersion) (*types.Transaction, error) {
	return _Version.Contract.TeeNodeVersionStruct(&_Version.TransactOpts, arg0)
}

// TeeUpgradePathStruct is a paid mutator transaction binding the contract method 0x470b8626.
//
// Solidity: function teeUpgradePathStruct(((bytes32,bytes32)[],(bytes32,bytes32)[]) ) returns()
func (_Version *VersionTransactor) TeeUpgradePathStruct(opts *bind.TransactOpts, arg0 ITeeVersionManagerTeeUpgradePath) (*types.Transaction, error) {
	return _Version.contract.Transact(opts, "teeUpgradePathStruct", arg0)
}

// TeeUpgradePathStruct is a paid mutator transaction binding the contract method 0x470b8626.
//
// Solidity: function teeUpgradePathStruct(((bytes32,bytes32)[],(bytes32,bytes32)[]) ) returns()
func (_Version *VersionSession) TeeUpgradePathStruct(arg0 ITeeVersionManagerTeeUpgradePath) (*types.Transaction, error) {
	return _Version.Contract.TeeUpgradePathStruct(&_Version.TransactOpts, arg0)
}

// TeeUpgradePathStruct is a paid mutator transaction binding the contract method 0x470b8626.
//
// Solidity: function teeUpgradePathStruct(((bytes32,bytes32)[],(bytes32,bytes32)[]) ) returns()
func (_Version *VersionTransactorSession) TeeUpgradePathStruct(arg0 ITeeVersionManagerTeeUpgradePath) (*types.Transaction, error) {
	return _Version.Contract.TeeUpgradePathStruct(&_Version.TransactOpts, arg0)
}
