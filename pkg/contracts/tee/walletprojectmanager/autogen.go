// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletprojectmanager

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

// WalletProjectManagerMetaData contains all meta data concerning the WalletProjectManager contract.
var WalletProjectManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SigningAlgoNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotPartOfProject\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotProductionReady\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"backupManager\",\"type\":\"address\"}],\"name\":\"BackupManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"}],\"name\":\"ProjectCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_signingAlgo\",\"type\":\"bytes32\"}],\"name\":\"createProject\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getBackupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_backupManager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getExtensionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getKeyType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getSigningAlgo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_signingAlgo\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_backupManager\",\"type\":\"address\"}],\"name\":\"setBackupManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WalletProjectManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletProjectManagerMetaData.ABI instead.
var WalletProjectManagerABI = WalletProjectManagerMetaData.ABI

// WalletProjectManager is an auto generated Go binding around an Ethereum contract.
type WalletProjectManager struct {
	WalletProjectManagerCaller     // Read-only binding to the contract
	WalletProjectManagerTransactor // Write-only binding to the contract
	WalletProjectManagerFilterer   // Log filterer for contract events
}

// WalletProjectManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletProjectManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletProjectManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletProjectManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletProjectManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletProjectManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletProjectManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletProjectManagerSession struct {
	Contract     *WalletProjectManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// WalletProjectManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletProjectManagerCallerSession struct {
	Contract *WalletProjectManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// WalletProjectManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletProjectManagerTransactorSession struct {
	Contract     *WalletProjectManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// WalletProjectManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletProjectManagerRaw struct {
	Contract *WalletProjectManager // Generic contract binding to access the raw methods on
}

// WalletProjectManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletProjectManagerCallerRaw struct {
	Contract *WalletProjectManagerCaller // Generic read-only contract binding to access the raw methods on
}

// WalletProjectManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletProjectManagerTransactorRaw struct {
	Contract *WalletProjectManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletProjectManager creates a new instance of WalletProjectManager, bound to a specific deployed contract.
func NewWalletProjectManager(address common.Address, backend bind.ContractBackend) (*WalletProjectManager, error) {
	contract, err := bindWalletProjectManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletProjectManager{WalletProjectManagerCaller: WalletProjectManagerCaller{contract: contract}, WalletProjectManagerTransactor: WalletProjectManagerTransactor{contract: contract}, WalletProjectManagerFilterer: WalletProjectManagerFilterer{contract: contract}}, nil
}

// NewWalletProjectManagerCaller creates a new read-only instance of WalletProjectManager, bound to a specific deployed contract.
func NewWalletProjectManagerCaller(address common.Address, caller bind.ContractCaller) (*WalletProjectManagerCaller, error) {
	contract, err := bindWalletProjectManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletProjectManagerCaller{contract: contract}, nil
}

// NewWalletProjectManagerTransactor creates a new write-only instance of WalletProjectManager, bound to a specific deployed contract.
func NewWalletProjectManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletProjectManagerTransactor, error) {
	contract, err := bindWalletProjectManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletProjectManagerTransactor{contract: contract}, nil
}

// NewWalletProjectManagerFilterer creates a new log filterer instance of WalletProjectManager, bound to a specific deployed contract.
func NewWalletProjectManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletProjectManagerFilterer, error) {
	contract, err := bindWalletProjectManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletProjectManagerFilterer{contract: contract}, nil
}

// bindWalletProjectManager binds a generic wrapper to an already deployed contract.
func bindWalletProjectManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletProjectManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletProjectManager *WalletProjectManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletProjectManager.Contract.WalletProjectManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletProjectManager *WalletProjectManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletProjectManager.Contract.WalletProjectManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletProjectManager *WalletProjectManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletProjectManager.Contract.WalletProjectManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletProjectManager *WalletProjectManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletProjectManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletProjectManager *WalletProjectManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletProjectManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletProjectManager *WalletProjectManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletProjectManager.Contract.contract.Transact(opts, method, params...)
}

// GetBackupManager is a free data retrieval call binding the contract method 0x7bb77989.
//
// Solidity: function getBackupManager(bytes32 _projectId) view returns(address _backupManager)
func (_WalletProjectManager *WalletProjectManagerCaller) GetBackupManager(opts *bind.CallOpts, _projectId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletProjectManager.contract.Call(opts, &out, "getBackupManager", _projectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBackupManager is a free data retrieval call binding the contract method 0x7bb77989.
//
// Solidity: function getBackupManager(bytes32 _projectId) view returns(address _backupManager)
func (_WalletProjectManager *WalletProjectManagerSession) GetBackupManager(_projectId [32]byte) (common.Address, error) {
	return _WalletProjectManager.Contract.GetBackupManager(&_WalletProjectManager.CallOpts, _projectId)
}

// GetBackupManager is a free data retrieval call binding the contract method 0x7bb77989.
//
// Solidity: function getBackupManager(bytes32 _projectId) view returns(address _backupManager)
func (_WalletProjectManager *WalletProjectManagerCallerSession) GetBackupManager(_projectId [32]byte) (common.Address, error) {
	return _WalletProjectManager.Contract.GetBackupManager(&_WalletProjectManager.CallOpts, _projectId)
}

// GetExtensionId is a free data retrieval call binding the contract method 0x13def0e3.
//
// Solidity: function getExtensionId(bytes32 _projectId) view returns(uint256 _extensionId)
func (_WalletProjectManager *WalletProjectManagerCaller) GetExtensionId(opts *bind.CallOpts, _projectId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _WalletProjectManager.contract.Call(opts, &out, "getExtensionId", _projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExtensionId is a free data retrieval call binding the contract method 0x13def0e3.
//
// Solidity: function getExtensionId(bytes32 _projectId) view returns(uint256 _extensionId)
func (_WalletProjectManager *WalletProjectManagerSession) GetExtensionId(_projectId [32]byte) (*big.Int, error) {
	return _WalletProjectManager.Contract.GetExtensionId(&_WalletProjectManager.CallOpts, _projectId)
}

// GetExtensionId is a free data retrieval call binding the contract method 0x13def0e3.
//
// Solidity: function getExtensionId(bytes32 _projectId) view returns(uint256 _extensionId)
func (_WalletProjectManager *WalletProjectManagerCallerSession) GetExtensionId(_projectId [32]byte) (*big.Int, error) {
	return _WalletProjectManager.Contract.GetExtensionId(&_WalletProjectManager.CallOpts, _projectId)
}

// GetKeyType is a free data retrieval call binding the contract method 0x779d385a.
//
// Solidity: function getKeyType(bytes32 _projectId) view returns(bytes32 _keyType)
func (_WalletProjectManager *WalletProjectManagerCaller) GetKeyType(opts *bind.CallOpts, _projectId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletProjectManager.contract.Call(opts, &out, "getKeyType", _projectId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKeyType is a free data retrieval call binding the contract method 0x779d385a.
//
// Solidity: function getKeyType(bytes32 _projectId) view returns(bytes32 _keyType)
func (_WalletProjectManager *WalletProjectManagerSession) GetKeyType(_projectId [32]byte) ([32]byte, error) {
	return _WalletProjectManager.Contract.GetKeyType(&_WalletProjectManager.CallOpts, _projectId)
}

// GetKeyType is a free data retrieval call binding the contract method 0x779d385a.
//
// Solidity: function getKeyType(bytes32 _projectId) view returns(bytes32 _keyType)
func (_WalletProjectManager *WalletProjectManagerCallerSession) GetKeyType(_projectId [32]byte) ([32]byte, error) {
	return _WalletProjectManager.Contract.GetKeyType(&_WalletProjectManager.CallOpts, _projectId)
}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 _projectId) view returns(address _owner)
func (_WalletProjectManager *WalletProjectManagerCaller) GetOwner(opts *bind.CallOpts, _projectId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletProjectManager.contract.Call(opts, &out, "getOwner", _projectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 _projectId) view returns(address _owner)
func (_WalletProjectManager *WalletProjectManagerSession) GetOwner(_projectId [32]byte) (common.Address, error) {
	return _WalletProjectManager.Contract.GetOwner(&_WalletProjectManager.CallOpts, _projectId)
}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 _projectId) view returns(address _owner)
func (_WalletProjectManager *WalletProjectManagerCallerSession) GetOwner(_projectId [32]byte) (common.Address, error) {
	return _WalletProjectManager.Contract.GetOwner(&_WalletProjectManager.CallOpts, _projectId)
}

// GetSigningAlgo is a free data retrieval call binding the contract method 0x17566896.
//
// Solidity: function getSigningAlgo(bytes32 _projectId) view returns(bytes32 _signingAlgo)
func (_WalletProjectManager *WalletProjectManagerCaller) GetSigningAlgo(opts *bind.CallOpts, _projectId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletProjectManager.contract.Call(opts, &out, "getSigningAlgo", _projectId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSigningAlgo is a free data retrieval call binding the contract method 0x17566896.
//
// Solidity: function getSigningAlgo(bytes32 _projectId) view returns(bytes32 _signingAlgo)
func (_WalletProjectManager *WalletProjectManagerSession) GetSigningAlgo(_projectId [32]byte) ([32]byte, error) {
	return _WalletProjectManager.Contract.GetSigningAlgo(&_WalletProjectManager.CallOpts, _projectId)
}

// GetSigningAlgo is a free data retrieval call binding the contract method 0x17566896.
//
// Solidity: function getSigningAlgo(bytes32 _projectId) view returns(bytes32 _signingAlgo)
func (_WalletProjectManager *WalletProjectManagerCallerSession) GetSigningAlgo(_projectId [32]byte) ([32]byte, error) {
	return _WalletProjectManager.Contract.GetSigningAlgo(&_WalletProjectManager.CallOpts, _projectId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x7ff0e872.
//
// Solidity: function confirmOwnership(bytes32 _projectId) returns()
func (_WalletProjectManager *WalletProjectManagerTransactor) ConfirmOwnership(opts *bind.TransactOpts, _projectId [32]byte) (*types.Transaction, error) {
	return _WalletProjectManager.contract.Transact(opts, "confirmOwnership", _projectId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x7ff0e872.
//
// Solidity: function confirmOwnership(bytes32 _projectId) returns()
func (_WalletProjectManager *WalletProjectManagerSession) ConfirmOwnership(_projectId [32]byte) (*types.Transaction, error) {
	return _WalletProjectManager.Contract.ConfirmOwnership(&_WalletProjectManager.TransactOpts, _projectId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x7ff0e872.
//
// Solidity: function confirmOwnership(bytes32 _projectId) returns()
func (_WalletProjectManager *WalletProjectManagerTransactorSession) ConfirmOwnership(_projectId [32]byte) (*types.Transaction, error) {
	return _WalletProjectManager.Contract.ConfirmOwnership(&_WalletProjectManager.TransactOpts, _projectId)
}

// CreateProject is a paid mutator transaction binding the contract method 0xa650c968.
//
// Solidity: function createProject(uint256 _extensionId, bytes32 _keyType, bytes32 _signingAlgo) returns(bytes32 _projectId)
func (_WalletProjectManager *WalletProjectManagerTransactor) CreateProject(opts *bind.TransactOpts, _extensionId *big.Int, _keyType [32]byte, _signingAlgo [32]byte) (*types.Transaction, error) {
	return _WalletProjectManager.contract.Transact(opts, "createProject", _extensionId, _keyType, _signingAlgo)
}

// CreateProject is a paid mutator transaction binding the contract method 0xa650c968.
//
// Solidity: function createProject(uint256 _extensionId, bytes32 _keyType, bytes32 _signingAlgo) returns(bytes32 _projectId)
func (_WalletProjectManager *WalletProjectManagerSession) CreateProject(_extensionId *big.Int, _keyType [32]byte, _signingAlgo [32]byte) (*types.Transaction, error) {
	return _WalletProjectManager.Contract.CreateProject(&_WalletProjectManager.TransactOpts, _extensionId, _keyType, _signingAlgo)
}

// CreateProject is a paid mutator transaction binding the contract method 0xa650c968.
//
// Solidity: function createProject(uint256 _extensionId, bytes32 _keyType, bytes32 _signingAlgo) returns(bytes32 _projectId)
func (_WalletProjectManager *WalletProjectManagerTransactorSession) CreateProject(_extensionId *big.Int, _keyType [32]byte, _signingAlgo [32]byte) (*types.Transaction, error) {
	return _WalletProjectManager.Contract.CreateProject(&_WalletProjectManager.TransactOpts, _extensionId, _keyType, _signingAlgo)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x3ad60dd6.
//
// Solidity: function proposeNewOwner(bytes32 _projectId, address _newOwner) returns()
func (_WalletProjectManager *WalletProjectManagerTransactor) ProposeNewOwner(opts *bind.TransactOpts, _projectId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _WalletProjectManager.contract.Transact(opts, "proposeNewOwner", _projectId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x3ad60dd6.
//
// Solidity: function proposeNewOwner(bytes32 _projectId, address _newOwner) returns()
func (_WalletProjectManager *WalletProjectManagerSession) ProposeNewOwner(_projectId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _WalletProjectManager.Contract.ProposeNewOwner(&_WalletProjectManager.TransactOpts, _projectId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x3ad60dd6.
//
// Solidity: function proposeNewOwner(bytes32 _projectId, address _newOwner) returns()
func (_WalletProjectManager *WalletProjectManagerTransactorSession) ProposeNewOwner(_projectId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _WalletProjectManager.Contract.ProposeNewOwner(&_WalletProjectManager.TransactOpts, _projectId, _newOwner)
}

// SetBackupManager is a paid mutator transaction binding the contract method 0x755fd7c2.
//
// Solidity: function setBackupManager(bytes32 _projectId, address _backupManager) returns()
func (_WalletProjectManager *WalletProjectManagerTransactor) SetBackupManager(opts *bind.TransactOpts, _projectId [32]byte, _backupManager common.Address) (*types.Transaction, error) {
	return _WalletProjectManager.contract.Transact(opts, "setBackupManager", _projectId, _backupManager)
}

// SetBackupManager is a paid mutator transaction binding the contract method 0x755fd7c2.
//
// Solidity: function setBackupManager(bytes32 _projectId, address _backupManager) returns()
func (_WalletProjectManager *WalletProjectManagerSession) SetBackupManager(_projectId [32]byte, _backupManager common.Address) (*types.Transaction, error) {
	return _WalletProjectManager.Contract.SetBackupManager(&_WalletProjectManager.TransactOpts, _projectId, _backupManager)
}

// SetBackupManager is a paid mutator transaction binding the contract method 0x755fd7c2.
//
// Solidity: function setBackupManager(bytes32 _projectId, address _backupManager) returns()
func (_WalletProjectManager *WalletProjectManagerTransactorSession) SetBackupManager(_projectId [32]byte, _backupManager common.Address) (*types.Transaction, error) {
	return _WalletProjectManager.Contract.SetBackupManager(&_WalletProjectManager.TransactOpts, _projectId, _backupManager)
}

// WalletProjectManagerBackupManagerSetIterator is returned from FilterBackupManagerSet and is used to iterate over the raw logs and unpacked data for BackupManagerSet events raised by the WalletProjectManager contract.
type WalletProjectManagerBackupManagerSetIterator struct {
	Event *WalletProjectManagerBackupManagerSet // Event containing the contract specifics and raw log

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
func (it *WalletProjectManagerBackupManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProjectManagerBackupManagerSet)
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
		it.Event = new(WalletProjectManagerBackupManagerSet)
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
func (it *WalletProjectManagerBackupManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProjectManagerBackupManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProjectManagerBackupManagerSet represents a BackupManagerSet event raised by the WalletProjectManager contract.
type WalletProjectManagerBackupManagerSet struct {
	ProjectId     [32]byte
	BackupManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBackupManagerSet is a free log retrieval operation binding the contract event 0x15ba31a03fdfef8efb2a9ebd2e02cb87cc38cb96de404cc4603067bb2360076e.
//
// Solidity: event BackupManagerSet(bytes32 indexed projectId, address indexed backupManager)
func (_WalletProjectManager *WalletProjectManagerFilterer) FilterBackupManagerSet(opts *bind.FilterOpts, projectId [][32]byte, backupManager []common.Address) (*WalletProjectManagerBackupManagerSetIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var backupManagerRule []interface{}
	for _, backupManagerItem := range backupManager {
		backupManagerRule = append(backupManagerRule, backupManagerItem)
	}

	logs, sub, err := _WalletProjectManager.contract.FilterLogs(opts, "BackupManagerSet", projectIdRule, backupManagerRule)
	if err != nil {
		return nil, err
	}
	return &WalletProjectManagerBackupManagerSetIterator{contract: _WalletProjectManager.contract, event: "BackupManagerSet", logs: logs, sub: sub}, nil
}

// WatchBackupManagerSet is a free log subscription operation binding the contract event 0x15ba31a03fdfef8efb2a9ebd2e02cb87cc38cb96de404cc4603067bb2360076e.
//
// Solidity: event BackupManagerSet(bytes32 indexed projectId, address indexed backupManager)
func (_WalletProjectManager *WalletProjectManagerFilterer) WatchBackupManagerSet(opts *bind.WatchOpts, sink chan<- *WalletProjectManagerBackupManagerSet, projectId [][32]byte, backupManager []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var backupManagerRule []interface{}
	for _, backupManagerItem := range backupManager {
		backupManagerRule = append(backupManagerRule, backupManagerItem)
	}

	logs, sub, err := _WalletProjectManager.contract.WatchLogs(opts, "BackupManagerSet", projectIdRule, backupManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProjectManagerBackupManagerSet)
				if err := _WalletProjectManager.contract.UnpackLog(event, "BackupManagerSet", log); err != nil {
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
func (_WalletProjectManager *WalletProjectManagerFilterer) ParseBackupManagerSet(log types.Log) (*WalletProjectManagerBackupManagerSet, error) {
	event := new(WalletProjectManagerBackupManagerSet)
	if err := _WalletProjectManager.contract.UnpackLog(event, "BackupManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProjectManagerNewOwnerProposedIterator is returned from FilterNewOwnerProposed and is used to iterate over the raw logs and unpacked data for NewOwnerProposed events raised by the WalletProjectManager contract.
type WalletProjectManagerNewOwnerProposedIterator struct {
	Event *WalletProjectManagerNewOwnerProposed // Event containing the contract specifics and raw log

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
func (it *WalletProjectManagerNewOwnerProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProjectManagerNewOwnerProposed)
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
		it.Event = new(WalletProjectManagerNewOwnerProposed)
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
func (it *WalletProjectManagerNewOwnerProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProjectManagerNewOwnerProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProjectManagerNewOwnerProposed represents a NewOwnerProposed event raised by the WalletProjectManager contract.
type WalletProjectManagerNewOwnerProposed struct {
	ProjectId [32]byte
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerProposed is a free log retrieval operation binding the contract event 0x90bc1e7c886602fdb10896f74fe5f65b1728518f0154de79478f10890351e725.
//
// Solidity: event NewOwnerProposed(bytes32 indexed projectId, address indexed newOwner)
func (_WalletProjectManager *WalletProjectManagerFilterer) FilterNewOwnerProposed(opts *bind.FilterOpts, projectId [][32]byte, newOwner []common.Address) (*WalletProjectManagerNewOwnerProposedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WalletProjectManager.contract.FilterLogs(opts, "NewOwnerProposed", projectIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WalletProjectManagerNewOwnerProposedIterator{contract: _WalletProjectManager.contract, event: "NewOwnerProposed", logs: logs, sub: sub}, nil
}

// WatchNewOwnerProposed is a free log subscription operation binding the contract event 0x90bc1e7c886602fdb10896f74fe5f65b1728518f0154de79478f10890351e725.
//
// Solidity: event NewOwnerProposed(bytes32 indexed projectId, address indexed newOwner)
func (_WalletProjectManager *WalletProjectManagerFilterer) WatchNewOwnerProposed(opts *bind.WatchOpts, sink chan<- *WalletProjectManagerNewOwnerProposed, projectId [][32]byte, newOwner []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WalletProjectManager.contract.WatchLogs(opts, "NewOwnerProposed", projectIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProjectManagerNewOwnerProposed)
				if err := _WalletProjectManager.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
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
func (_WalletProjectManager *WalletProjectManagerFilterer) ParseNewOwnerProposed(log types.Log) (*WalletProjectManagerNewOwnerProposed, error) {
	event := new(WalletProjectManagerNewOwnerProposed)
	if err := _WalletProjectManager.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProjectManagerOwnershipConfirmedIterator is returned from FilterOwnershipConfirmed and is used to iterate over the raw logs and unpacked data for OwnershipConfirmed events raised by the WalletProjectManager contract.
type WalletProjectManagerOwnershipConfirmedIterator struct {
	Event *WalletProjectManagerOwnershipConfirmed // Event containing the contract specifics and raw log

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
func (it *WalletProjectManagerOwnershipConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProjectManagerOwnershipConfirmed)
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
		it.Event = new(WalletProjectManagerOwnershipConfirmed)
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
func (it *WalletProjectManagerOwnershipConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProjectManagerOwnershipConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProjectManagerOwnershipConfirmed represents a OwnershipConfirmed event raised by the WalletProjectManager contract.
type WalletProjectManagerOwnershipConfirmed struct {
	ProjectId [32]byte
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnershipConfirmed is a free log retrieval operation binding the contract event 0xf80ea9c0a069a69c7b1582cd4391b2bf86e63347d60b636a1200036b502305ba.
//
// Solidity: event OwnershipConfirmed(bytes32 indexed projectId, address indexed newOwner)
func (_WalletProjectManager *WalletProjectManagerFilterer) FilterOwnershipConfirmed(opts *bind.FilterOpts, projectId [][32]byte, newOwner []common.Address) (*WalletProjectManagerOwnershipConfirmedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WalletProjectManager.contract.FilterLogs(opts, "OwnershipConfirmed", projectIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WalletProjectManagerOwnershipConfirmedIterator{contract: _WalletProjectManager.contract, event: "OwnershipConfirmed", logs: logs, sub: sub}, nil
}

// WatchOwnershipConfirmed is a free log subscription operation binding the contract event 0xf80ea9c0a069a69c7b1582cd4391b2bf86e63347d60b636a1200036b502305ba.
//
// Solidity: event OwnershipConfirmed(bytes32 indexed projectId, address indexed newOwner)
func (_WalletProjectManager *WalletProjectManagerFilterer) WatchOwnershipConfirmed(opts *bind.WatchOpts, sink chan<- *WalletProjectManagerOwnershipConfirmed, projectId [][32]byte, newOwner []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WalletProjectManager.contract.WatchLogs(opts, "OwnershipConfirmed", projectIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProjectManagerOwnershipConfirmed)
				if err := _WalletProjectManager.contract.UnpackLog(event, "OwnershipConfirmed", log); err != nil {
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
func (_WalletProjectManager *WalletProjectManagerFilterer) ParseOwnershipConfirmed(log types.Log) (*WalletProjectManagerOwnershipConfirmed, error) {
	event := new(WalletProjectManagerOwnershipConfirmed)
	if err := _WalletProjectManager.contract.UnpackLog(event, "OwnershipConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProjectManagerProjectCreatedIterator is returned from FilterProjectCreated and is used to iterate over the raw logs and unpacked data for ProjectCreated events raised by the WalletProjectManager contract.
type WalletProjectManagerProjectCreatedIterator struct {
	Event *WalletProjectManagerProjectCreated // Event containing the contract specifics and raw log

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
func (it *WalletProjectManagerProjectCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProjectManagerProjectCreated)
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
		it.Event = new(WalletProjectManagerProjectCreated)
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
func (it *WalletProjectManagerProjectCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProjectManagerProjectCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProjectManagerProjectCreated represents a ProjectCreated event raised by the WalletProjectManager contract.
type WalletProjectManagerProjectCreated struct {
	ProjectId   [32]byte
	Owner       common.Address
	ExtensionId *big.Int
	KeyType     [32]byte
	SigningAlgo [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProjectCreated is a free log retrieval operation binding the contract event 0x6a98f0d33cd18de1a10aa099b478292aa9c6b801ce4826fabc4a674eaad6a968.
//
// Solidity: event ProjectCreated(bytes32 indexed projectId, address indexed owner, uint256 extensionId, bytes32 keyType, bytes32 signingAlgo)
func (_WalletProjectManager *WalletProjectManagerFilterer) FilterProjectCreated(opts *bind.FilterOpts, projectId [][32]byte, owner []common.Address) (*WalletProjectManagerProjectCreatedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WalletProjectManager.contract.FilterLogs(opts, "ProjectCreated", projectIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &WalletProjectManagerProjectCreatedIterator{contract: _WalletProjectManager.contract, event: "ProjectCreated", logs: logs, sub: sub}, nil
}

// WatchProjectCreated is a free log subscription operation binding the contract event 0x6a98f0d33cd18de1a10aa099b478292aa9c6b801ce4826fabc4a674eaad6a968.
//
// Solidity: event ProjectCreated(bytes32 indexed projectId, address indexed owner, uint256 extensionId, bytes32 keyType, bytes32 signingAlgo)
func (_WalletProjectManager *WalletProjectManagerFilterer) WatchProjectCreated(opts *bind.WatchOpts, sink chan<- *WalletProjectManagerProjectCreated, projectId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _WalletProjectManager.contract.WatchLogs(opts, "ProjectCreated", projectIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProjectManagerProjectCreated)
				if err := _WalletProjectManager.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
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

// ParseProjectCreated is a log parse operation binding the contract event 0x6a98f0d33cd18de1a10aa099b478292aa9c6b801ce4826fabc4a674eaad6a968.
//
// Solidity: event ProjectCreated(bytes32 indexed projectId, address indexed owner, uint256 extensionId, bytes32 keyType, bytes32 signingAlgo)
func (_WalletProjectManager *WalletProjectManagerFilterer) ParseProjectCreated(log types.Log) (*WalletProjectManagerProjectCreated, error) {
	event := new(WalletProjectManagerProjectCreated)
	if err := _WalletProjectManager.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
