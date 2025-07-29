// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teewalletprojectmanager

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

// TeeWalletProjectManagerMetaData contains all meta data concerning the TeeWalletProjectManager contract.
var TeeWalletProjectManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"backupManager\",\"type\":\"address\"}],\"name\":\"BackupManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"DefaultWalletSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"submitAddress\",\"type\":\"address\"}],\"name\":\"ProjectCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_submitAddress\",\"type\":\"address\"}],\"name\":\"createProject\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getBackupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_backupManager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getDefaultWalletInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_submitAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getOpType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getOpTypeConstants\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getSubmitAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_submitAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_backupManager\",\"type\":\"address\"}],\"name\":\"setBackupManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"setDefaultWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeWalletProjectManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeWalletProjectManagerMetaData.ABI instead.
var TeeWalletProjectManagerABI = TeeWalletProjectManagerMetaData.ABI

// TeeWalletProjectManager is an auto generated Go binding around an Ethereum contract.
type TeeWalletProjectManager struct {
	TeeWalletProjectManagerCaller     // Read-only binding to the contract
	TeeWalletProjectManagerTransactor // Write-only binding to the contract
	TeeWalletProjectManagerFilterer   // Log filterer for contract events
}

// TeeWalletProjectManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeWalletProjectManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletProjectManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeWalletProjectManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletProjectManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeWalletProjectManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletProjectManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeWalletProjectManagerSession struct {
	Contract     *TeeWalletProjectManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TeeWalletProjectManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeWalletProjectManagerCallerSession struct {
	Contract *TeeWalletProjectManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// TeeWalletProjectManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeWalletProjectManagerTransactorSession struct {
	Contract     *TeeWalletProjectManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// TeeWalletProjectManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeWalletProjectManagerRaw struct {
	Contract *TeeWalletProjectManager // Generic contract binding to access the raw methods on
}

// TeeWalletProjectManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeWalletProjectManagerCallerRaw struct {
	Contract *TeeWalletProjectManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TeeWalletProjectManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeWalletProjectManagerTransactorRaw struct {
	Contract *TeeWalletProjectManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeWalletProjectManager creates a new instance of TeeWalletProjectManager, bound to a specific deployed contract.
func NewTeeWalletProjectManager(address common.Address, backend bind.ContractBackend) (*TeeWalletProjectManager, error) {
	contract, err := bindTeeWalletProjectManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManager{TeeWalletProjectManagerCaller: TeeWalletProjectManagerCaller{contract: contract}, TeeWalletProjectManagerTransactor: TeeWalletProjectManagerTransactor{contract: contract}, TeeWalletProjectManagerFilterer: TeeWalletProjectManagerFilterer{contract: contract}}, nil
}

// NewTeeWalletProjectManagerCaller creates a new read-only instance of TeeWalletProjectManager, bound to a specific deployed contract.
func NewTeeWalletProjectManagerCaller(address common.Address, caller bind.ContractCaller) (*TeeWalletProjectManagerCaller, error) {
	contract, err := bindTeeWalletProjectManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerCaller{contract: contract}, nil
}

// NewTeeWalletProjectManagerTransactor creates a new write-only instance of TeeWalletProjectManager, bound to a specific deployed contract.
func NewTeeWalletProjectManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeWalletProjectManagerTransactor, error) {
	contract, err := bindTeeWalletProjectManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerTransactor{contract: contract}, nil
}

// NewTeeWalletProjectManagerFilterer creates a new log filterer instance of TeeWalletProjectManager, bound to a specific deployed contract.
func NewTeeWalletProjectManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeWalletProjectManagerFilterer, error) {
	contract, err := bindTeeWalletProjectManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerFilterer{contract: contract}, nil
}

// bindTeeWalletProjectManager binds a generic wrapper to an already deployed contract.
func bindTeeWalletProjectManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeWalletProjectManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletProjectManager *TeeWalletProjectManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletProjectManager.Contract.TeeWalletProjectManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletProjectManager *TeeWalletProjectManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.TeeWalletProjectManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletProjectManager *TeeWalletProjectManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.TeeWalletProjectManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletProjectManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.contract.Transact(opts, method, params...)
}

// GetBackupManager is a free data retrieval call binding the contract method 0x7bb77989.
//
// Solidity: function getBackupManager(bytes32 _projectId) view returns(address _backupManager)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) GetBackupManager(opts *bind.CallOpts, _projectId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "getBackupManager", _projectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBackupManager is a free data retrieval call binding the contract method 0x7bb77989.
//
// Solidity: function getBackupManager(bytes32 _projectId) view returns(address _backupManager)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) GetBackupManager(_projectId [32]byte) (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GetBackupManager(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetBackupManager is a free data retrieval call binding the contract method 0x7bb77989.
//
// Solidity: function getBackupManager(bytes32 _projectId) view returns(address _backupManager)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) GetBackupManager(_projectId [32]byte) (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GetBackupManager(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetDefaultWalletInfo is a free data retrieval call binding the contract method 0x1a9f0cb0.
//
// Solidity: function getDefaultWalletInfo(bytes32 _projectId) view returns(bytes32 _walletId, bytes32 _opType, address _submitAddress)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) GetDefaultWalletInfo(opts *bind.CallOpts, _projectId [32]byte) (struct {
	WalletId      [32]byte
	OpType        [32]byte
	SubmitAddress common.Address
}, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "getDefaultWalletInfo", _projectId)

	outstruct := new(struct {
		WalletId      [32]byte
		OpType        [32]byte
		SubmitAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WalletId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.OpType = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SubmitAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetDefaultWalletInfo is a free data retrieval call binding the contract method 0x1a9f0cb0.
//
// Solidity: function getDefaultWalletInfo(bytes32 _projectId) view returns(bytes32 _walletId, bytes32 _opType, address _submitAddress)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) GetDefaultWalletInfo(_projectId [32]byte) (struct {
	WalletId      [32]byte
	OpType        [32]byte
	SubmitAddress common.Address
}, error) {
	return _TeeWalletProjectManager.Contract.GetDefaultWalletInfo(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetDefaultWalletInfo is a free data retrieval call binding the contract method 0x1a9f0cb0.
//
// Solidity: function getDefaultWalletInfo(bytes32 _projectId) view returns(bytes32 _walletId, bytes32 _opType, address _submitAddress)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) GetDefaultWalletInfo(_projectId [32]byte) (struct {
	WalletId      [32]byte
	OpType        [32]byte
	SubmitAddress common.Address
}, error) {
	return _TeeWalletProjectManager.Contract.GetDefaultWalletInfo(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetOpType is a free data retrieval call binding the contract method 0x1409b249.
//
// Solidity: function getOpType(bytes32 _projectId) view returns(bytes32 _opType)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) GetOpType(opts *bind.CallOpts, _projectId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "getOpType", _projectId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOpType is a free data retrieval call binding the contract method 0x1409b249.
//
// Solidity: function getOpType(bytes32 _projectId) view returns(bytes32 _opType)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) GetOpType(_projectId [32]byte) ([32]byte, error) {
	return _TeeWalletProjectManager.Contract.GetOpType(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetOpType is a free data retrieval call binding the contract method 0x1409b249.
//
// Solidity: function getOpType(bytes32 _projectId) view returns(bytes32 _opType)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) GetOpType(_projectId [32]byte) ([32]byte, error) {
	return _TeeWalletProjectManager.Contract.GetOpType(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetOpTypeConstants is a free data retrieval call binding the contract method 0x727b7d21.
//
// Solidity: function getOpTypeConstants(bytes32 _projectId) view returns(bytes)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) GetOpTypeConstants(opts *bind.CallOpts, _projectId [32]byte) ([]byte, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "getOpTypeConstants", _projectId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetOpTypeConstants is a free data retrieval call binding the contract method 0x727b7d21.
//
// Solidity: function getOpTypeConstants(bytes32 _projectId) view returns(bytes)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) GetOpTypeConstants(_projectId [32]byte) ([]byte, error) {
	return _TeeWalletProjectManager.Contract.GetOpTypeConstants(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetOpTypeConstants is a free data retrieval call binding the contract method 0x727b7d21.
//
// Solidity: function getOpTypeConstants(bytes32 _projectId) view returns(bytes)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) GetOpTypeConstants(_projectId [32]byte) ([]byte, error) {
	return _TeeWalletProjectManager.Contract.GetOpTypeConstants(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 _projectId) view returns(address _owner)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) GetOwner(opts *bind.CallOpts, _projectId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "getOwner", _projectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 _projectId) view returns(address _owner)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) GetOwner(_projectId [32]byte) (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GetOwner(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 _projectId) view returns(address _owner)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) GetOwner(_projectId [32]byte) (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GetOwner(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetSubmitAddress is a free data retrieval call binding the contract method 0x14edb8d7.
//
// Solidity: function getSubmitAddress(bytes32 _projectId) view returns(address _submitAddress)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) GetSubmitAddress(opts *bind.CallOpts, _projectId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "getSubmitAddress", _projectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubmitAddress is a free data retrieval call binding the contract method 0x14edb8d7.
//
// Solidity: function getSubmitAddress(bytes32 _projectId) view returns(address _submitAddress)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) GetSubmitAddress(_projectId [32]byte) (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GetSubmitAddress(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetSubmitAddress is a free data retrieval call binding the contract method 0x14edb8d7.
//
// Solidity: function getSubmitAddress(bytes32 _projectId) view returns(address _submitAddress)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) GetSubmitAddress(_projectId [32]byte) (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GetSubmitAddress(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x7ff0e872.
//
// Solidity: function confirmOwnership(bytes32 _projectId) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactor) ConfirmOwnership(opts *bind.TransactOpts, _projectId [32]byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.contract.Transact(opts, "confirmOwnership", _projectId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x7ff0e872.
//
// Solidity: function confirmOwnership(bytes32 _projectId) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) ConfirmOwnership(_projectId [32]byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.ConfirmOwnership(&_TeeWalletProjectManager.TransactOpts, _projectId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x7ff0e872.
//
// Solidity: function confirmOwnership(bytes32 _projectId) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorSession) ConfirmOwnership(_projectId [32]byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.ConfirmOwnership(&_TeeWalletProjectManager.TransactOpts, _projectId)
}

// CreateProject is a paid mutator transaction binding the contract method 0xef565a1c.
//
// Solidity: function createProject(uint256 _extensionId, bytes32 _opType, address _submitAddress) returns(bytes32 _projectId)
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactor) CreateProject(opts *bind.TransactOpts, _extensionId *big.Int, _opType [32]byte, _submitAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.contract.Transact(opts, "createProject", _extensionId, _opType, _submitAddress)
}

// CreateProject is a paid mutator transaction binding the contract method 0xef565a1c.
//
// Solidity: function createProject(uint256 _extensionId, bytes32 _opType, address _submitAddress) returns(bytes32 _projectId)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) CreateProject(_extensionId *big.Int, _opType [32]byte, _submitAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.CreateProject(&_TeeWalletProjectManager.TransactOpts, _extensionId, _opType, _submitAddress)
}

// CreateProject is a paid mutator transaction binding the contract method 0xef565a1c.
//
// Solidity: function createProject(uint256 _extensionId, bytes32 _opType, address _submitAddress) returns(bytes32 _projectId)
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorSession) CreateProject(_extensionId *big.Int, _opType [32]byte, _submitAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.CreateProject(&_TeeWalletProjectManager.TransactOpts, _extensionId, _opType, _submitAddress)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x3ad60dd6.
//
// Solidity: function proposeNewOwner(bytes32 _projectId, address _newOwner) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactor) ProposeNewOwner(opts *bind.TransactOpts, _projectId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.contract.Transact(opts, "proposeNewOwner", _projectId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x3ad60dd6.
//
// Solidity: function proposeNewOwner(bytes32 _projectId, address _newOwner) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) ProposeNewOwner(_projectId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.ProposeNewOwner(&_TeeWalletProjectManager.TransactOpts, _projectId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x3ad60dd6.
//
// Solidity: function proposeNewOwner(bytes32 _projectId, address _newOwner) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorSession) ProposeNewOwner(_projectId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.ProposeNewOwner(&_TeeWalletProjectManager.TransactOpts, _projectId, _newOwner)
}

// SetBackupManager is a paid mutator transaction binding the contract method 0x755fd7c2.
//
// Solidity: function setBackupManager(bytes32 _projectId, address _backupManager) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactor) SetBackupManager(opts *bind.TransactOpts, _projectId [32]byte, _backupManager common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.contract.Transact(opts, "setBackupManager", _projectId, _backupManager)
}

// SetBackupManager is a paid mutator transaction binding the contract method 0x755fd7c2.
//
// Solidity: function setBackupManager(bytes32 _projectId, address _backupManager) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) SetBackupManager(_projectId [32]byte, _backupManager common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.SetBackupManager(&_TeeWalletProjectManager.TransactOpts, _projectId, _backupManager)
}

// SetBackupManager is a paid mutator transaction binding the contract method 0x755fd7c2.
//
// Solidity: function setBackupManager(bytes32 _projectId, address _backupManager) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorSession) SetBackupManager(_projectId [32]byte, _backupManager common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.SetBackupManager(&_TeeWalletProjectManager.TransactOpts, _projectId, _backupManager)
}

// SetDefaultWallet is a paid mutator transaction binding the contract method 0x14ce9290.
//
// Solidity: function setDefaultWallet(bytes32 _projectId, bytes32 _walletId) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactor) SetDefaultWallet(opts *bind.TransactOpts, _projectId [32]byte, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.contract.Transact(opts, "setDefaultWallet", _projectId, _walletId)
}

// SetDefaultWallet is a paid mutator transaction binding the contract method 0x14ce9290.
//
// Solidity: function setDefaultWallet(bytes32 _projectId, bytes32 _walletId) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) SetDefaultWallet(_projectId [32]byte, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.SetDefaultWallet(&_TeeWalletProjectManager.TransactOpts, _projectId, _walletId)
}

// SetDefaultWallet is a paid mutator transaction binding the contract method 0x14ce9290.
//
// Solidity: function setDefaultWallet(bytes32 _projectId, bytes32 _walletId) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorSession) SetDefaultWallet(_projectId [32]byte, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.SetDefaultWallet(&_TeeWalletProjectManager.TransactOpts, _projectId, _walletId)
}

// TeeWalletProjectManagerBackupManagerSetIterator is returned from FilterBackupManagerSet and is used to iterate over the raw logs and unpacked data for BackupManagerSet events raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerBackupManagerSetIterator struct {
	Event *TeeWalletProjectManagerBackupManagerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeWalletProjectManagerBackupManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletProjectManagerBackupManagerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeWalletProjectManagerBackupManagerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeWalletProjectManagerBackupManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletProjectManagerBackupManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletProjectManagerBackupManagerSet represents a BackupManagerSet event raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerBackupManagerSet struct {
	ProjectId     [32]byte
	BackupManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBackupManagerSet is a free log retrieval operation binding the contract event 0x15ba31a03fdfef8efb2a9ebd2e02cb87cc38cb96de404cc4603067bb2360076e.
//
// Solidity: event BackupManagerSet(bytes32 indexed projectId, address indexed backupManager)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) FilterBackupManagerSet(opts *bind.FilterOpts, projectId [][32]byte, backupManager []common.Address) (*TeeWalletProjectManagerBackupManagerSetIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var backupManagerRule []interface{}
	for _, backupManagerItem := range backupManager {
		backupManagerRule = append(backupManagerRule, backupManagerItem)
	}

	logs, sub, err := _TeeWalletProjectManager.contract.FilterLogs(opts, "BackupManagerSet", projectIdRule, backupManagerRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerBackupManagerSetIterator{contract: _TeeWalletProjectManager.contract, event: "BackupManagerSet", logs: logs, sub: sub}, nil
}

// WatchBackupManagerSet is a free log subscription operation binding the contract event 0x15ba31a03fdfef8efb2a9ebd2e02cb87cc38cb96de404cc4603067bb2360076e.
//
// Solidity: event BackupManagerSet(bytes32 indexed projectId, address indexed backupManager)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) WatchBackupManagerSet(opts *bind.WatchOpts, sink chan<- *TeeWalletProjectManagerBackupManagerSet, projectId [][32]byte, backupManager []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var backupManagerRule []interface{}
	for _, backupManagerItem := range backupManager {
		backupManagerRule = append(backupManagerRule, backupManagerItem)
	}

	logs, sub, err := _TeeWalletProjectManager.contract.WatchLogs(opts, "BackupManagerSet", projectIdRule, backupManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletProjectManagerBackupManagerSet)
				if err := _TeeWalletProjectManager.contract.UnpackLog(event, "BackupManagerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBackupManagerSet is a log parse operation binding the contract event 0x15ba31a03fdfef8efb2a9ebd2e02cb87cc38cb96de404cc4603067bb2360076e.
//
// Solidity: event BackupManagerSet(bytes32 indexed projectId, address indexed backupManager)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) ParseBackupManagerSet(log types.Log) (*TeeWalletProjectManagerBackupManagerSet, error) {
	event := new(TeeWalletProjectManagerBackupManagerSet)
	if err := _TeeWalletProjectManager.contract.UnpackLog(event, "BackupManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletProjectManagerDefaultWalletSetIterator is returned from FilterDefaultWalletSet and is used to iterate over the raw logs and unpacked data for DefaultWalletSet events raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerDefaultWalletSetIterator struct {
	Event *TeeWalletProjectManagerDefaultWalletSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeWalletProjectManagerDefaultWalletSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletProjectManagerDefaultWalletSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeWalletProjectManagerDefaultWalletSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeWalletProjectManagerDefaultWalletSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletProjectManagerDefaultWalletSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletProjectManagerDefaultWalletSet represents a DefaultWalletSet event raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerDefaultWalletSet struct {
	ProjectId [32]byte
	WalletId  [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDefaultWalletSet is a free log retrieval operation binding the contract event 0xec54818f21763e2755ddc9d073c02b5a06e0291b600219046210461045bfae3e.
//
// Solidity: event DefaultWalletSet(bytes32 indexed projectId, bytes32 indexed walletId)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) FilterDefaultWalletSet(opts *bind.FilterOpts, projectId [][32]byte, walletId [][32]byte) (*TeeWalletProjectManagerDefaultWalletSetIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletProjectManager.contract.FilterLogs(opts, "DefaultWalletSet", projectIdRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerDefaultWalletSetIterator{contract: _TeeWalletProjectManager.contract, event: "DefaultWalletSet", logs: logs, sub: sub}, nil
}

// WatchDefaultWalletSet is a free log subscription operation binding the contract event 0xec54818f21763e2755ddc9d073c02b5a06e0291b600219046210461045bfae3e.
//
// Solidity: event DefaultWalletSet(bytes32 indexed projectId, bytes32 indexed walletId)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) WatchDefaultWalletSet(opts *bind.WatchOpts, sink chan<- *TeeWalletProjectManagerDefaultWalletSet, projectId [][32]byte, walletId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletProjectManager.contract.WatchLogs(opts, "DefaultWalletSet", projectIdRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletProjectManagerDefaultWalletSet)
				if err := _TeeWalletProjectManager.contract.UnpackLog(event, "DefaultWalletSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultWalletSet is a log parse operation binding the contract event 0xec54818f21763e2755ddc9d073c02b5a06e0291b600219046210461045bfae3e.
//
// Solidity: event DefaultWalletSet(bytes32 indexed projectId, bytes32 indexed walletId)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) ParseDefaultWalletSet(log types.Log) (*TeeWalletProjectManagerDefaultWalletSet, error) {
	event := new(TeeWalletProjectManagerDefaultWalletSet)
	if err := _TeeWalletProjectManager.contract.UnpackLog(event, "DefaultWalletSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletProjectManagerNewOwnerProposedIterator is returned from FilterNewOwnerProposed and is used to iterate over the raw logs and unpacked data for NewOwnerProposed events raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerNewOwnerProposedIterator struct {
	Event *TeeWalletProjectManagerNewOwnerProposed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeWalletProjectManagerNewOwnerProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletProjectManagerNewOwnerProposed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeWalletProjectManagerNewOwnerProposed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeWalletProjectManagerNewOwnerProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletProjectManagerNewOwnerProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletProjectManagerNewOwnerProposed represents a NewOwnerProposed event raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerNewOwnerProposed struct {
	ProjectId [32]byte
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerProposed is a free log retrieval operation binding the contract event 0x90bc1e7c886602fdb10896f74fe5f65b1728518f0154de79478f10890351e725.
//
// Solidity: event NewOwnerProposed(bytes32 indexed projectId, address indexed newOwner)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) FilterNewOwnerProposed(opts *bind.FilterOpts, projectId [][32]byte, newOwner []common.Address) (*TeeWalletProjectManagerNewOwnerProposedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeWalletProjectManager.contract.FilterLogs(opts, "NewOwnerProposed", projectIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerNewOwnerProposedIterator{contract: _TeeWalletProjectManager.contract, event: "NewOwnerProposed", logs: logs, sub: sub}, nil
}

// WatchNewOwnerProposed is a free log subscription operation binding the contract event 0x90bc1e7c886602fdb10896f74fe5f65b1728518f0154de79478f10890351e725.
//
// Solidity: event NewOwnerProposed(bytes32 indexed projectId, address indexed newOwner)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) WatchNewOwnerProposed(opts *bind.WatchOpts, sink chan<- *TeeWalletProjectManagerNewOwnerProposed, projectId [][32]byte, newOwner []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeWalletProjectManager.contract.WatchLogs(opts, "NewOwnerProposed", projectIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletProjectManagerNewOwnerProposed)
				if err := _TeeWalletProjectManager.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewOwnerProposed is a log parse operation binding the contract event 0x90bc1e7c886602fdb10896f74fe5f65b1728518f0154de79478f10890351e725.
//
// Solidity: event NewOwnerProposed(bytes32 indexed projectId, address indexed newOwner)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) ParseNewOwnerProposed(log types.Log) (*TeeWalletProjectManagerNewOwnerProposed, error) {
	event := new(TeeWalletProjectManagerNewOwnerProposed)
	if err := _TeeWalletProjectManager.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletProjectManagerOwnershipConfirmedIterator is returned from FilterOwnershipConfirmed and is used to iterate over the raw logs and unpacked data for OwnershipConfirmed events raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerOwnershipConfirmedIterator struct {
	Event *TeeWalletProjectManagerOwnershipConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeWalletProjectManagerOwnershipConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletProjectManagerOwnershipConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeWalletProjectManagerOwnershipConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeWalletProjectManagerOwnershipConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletProjectManagerOwnershipConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletProjectManagerOwnershipConfirmed represents a OwnershipConfirmed event raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerOwnershipConfirmed struct {
	ProjectId [32]byte
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnershipConfirmed is a free log retrieval operation binding the contract event 0xf80ea9c0a069a69c7b1582cd4391b2bf86e63347d60b636a1200036b502305ba.
//
// Solidity: event OwnershipConfirmed(bytes32 indexed projectId, address indexed newOwner)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) FilterOwnershipConfirmed(opts *bind.FilterOpts, projectId [][32]byte, newOwner []common.Address) (*TeeWalletProjectManagerOwnershipConfirmedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeWalletProjectManager.contract.FilterLogs(opts, "OwnershipConfirmed", projectIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerOwnershipConfirmedIterator{contract: _TeeWalletProjectManager.contract, event: "OwnershipConfirmed", logs: logs, sub: sub}, nil
}

// WatchOwnershipConfirmed is a free log subscription operation binding the contract event 0xf80ea9c0a069a69c7b1582cd4391b2bf86e63347d60b636a1200036b502305ba.
//
// Solidity: event OwnershipConfirmed(bytes32 indexed projectId, address indexed newOwner)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) WatchOwnershipConfirmed(opts *bind.WatchOpts, sink chan<- *TeeWalletProjectManagerOwnershipConfirmed, projectId [][32]byte, newOwner []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeWalletProjectManager.contract.WatchLogs(opts, "OwnershipConfirmed", projectIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletProjectManagerOwnershipConfirmed)
				if err := _TeeWalletProjectManager.contract.UnpackLog(event, "OwnershipConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipConfirmed is a log parse operation binding the contract event 0xf80ea9c0a069a69c7b1582cd4391b2bf86e63347d60b636a1200036b502305ba.
//
// Solidity: event OwnershipConfirmed(bytes32 indexed projectId, address indexed newOwner)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) ParseOwnershipConfirmed(log types.Log) (*TeeWalletProjectManagerOwnershipConfirmed, error) {
	event := new(TeeWalletProjectManagerOwnershipConfirmed)
	if err := _TeeWalletProjectManager.contract.UnpackLog(event, "OwnershipConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletProjectManagerProjectCreatedIterator is returned from FilterProjectCreated and is used to iterate over the raw logs and unpacked data for ProjectCreated events raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerProjectCreatedIterator struct {
	Event *TeeWalletProjectManagerProjectCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TeeWalletProjectManagerProjectCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletProjectManagerProjectCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TeeWalletProjectManagerProjectCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TeeWalletProjectManagerProjectCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletProjectManagerProjectCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletProjectManagerProjectCreated represents a ProjectCreated event raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerProjectCreated struct {
	ProjectId     [32]byte
	Owner         common.Address
	ExtensionId   *big.Int
	OpType        [32]byte
	SubmitAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProjectCreated is a free log retrieval operation binding the contract event 0x298a5f56e1c0f474f194b7963de86b48f28c9b42b450babfa242ff1072c3766d.
//
// Solidity: event ProjectCreated(bytes32 indexed projectId, address indexed owner, uint256 extensionId, bytes32 opType, address submitAddress)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) FilterProjectCreated(opts *bind.FilterOpts, projectId [][32]byte, owner []common.Address) (*TeeWalletProjectManagerProjectCreatedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeWalletProjectManager.contract.FilterLogs(opts, "ProjectCreated", projectIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerProjectCreatedIterator{contract: _TeeWalletProjectManager.contract, event: "ProjectCreated", logs: logs, sub: sub}, nil
}

// WatchProjectCreated is a free log subscription operation binding the contract event 0x298a5f56e1c0f474f194b7963de86b48f28c9b42b450babfa242ff1072c3766d.
//
// Solidity: event ProjectCreated(bytes32 indexed projectId, address indexed owner, uint256 extensionId, bytes32 opType, address submitAddress)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) WatchProjectCreated(opts *bind.WatchOpts, sink chan<- *TeeWalletProjectManagerProjectCreated, projectId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeWalletProjectManager.contract.WatchLogs(opts, "ProjectCreated", projectIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletProjectManagerProjectCreated)
				if err := _TeeWalletProjectManager.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProjectCreated is a log parse operation binding the contract event 0x298a5f56e1c0f474f194b7963de86b48f28c9b42b450babfa242ff1072c3766d.
//
// Solidity: event ProjectCreated(bytes32 indexed projectId, address indexed owner, uint256 extensionId, bytes32 opType, address submitAddress)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) ParseProjectCreated(log types.Log) (*TeeWalletProjectManagerProjectCreated, error) {
	event := new(TeeWalletProjectManagerProjectCreated)
	if err := _TeeWalletProjectManager.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
