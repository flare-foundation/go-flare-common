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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuthorizationAddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SigningAlgoNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotPartOfProject\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotProductionReady\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"backupManager\",\"type\":\"address\"}],\"name\":\"BackupManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizationAddress\",\"type\":\"address\"}],\"name\":\"ProjectCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_signingAlgo\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"name\":\"createProject\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getAuthorizationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getBackupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_backupManager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getExtensionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getKeyType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_projectOwner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getSigningAlgo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_signingAlgo\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"projectCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"}],\"name\":\"proposedProjectOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_backupManager\",\"type\":\"address\"}],\"name\":\"setBackupManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeExtensionRegistry\",\"outputs\":[{\"internalType\":\"contractITeeExtensionRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeOwnerAllowlist\",\"outputs\":[{\"internalType\":\"contractITeeOwnerAllowlist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeWalletProjectManager.Contract.UPGRADEINTERFACEVERSION(&_TeeWalletProjectManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeWalletProjectManager.Contract.UPGRADEINTERFACEVERSION(&_TeeWalletProjectManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GetAddressUpdater(&_TeeWalletProjectManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GetAddressUpdater(&_TeeWalletProjectManager.CallOpts)
}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x33fa0007.
//
// Solidity: function getAuthorizationAddress(bytes32 _projectId) view returns(address _authorizationAddress)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) GetAuthorizationAddress(opts *bind.CallOpts, _projectId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "getAuthorizationAddress", _projectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x33fa0007.
//
// Solidity: function getAuthorizationAddress(bytes32 _projectId) view returns(address _authorizationAddress)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) GetAuthorizationAddress(_projectId [32]byte) (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GetAuthorizationAddress(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x33fa0007.
//
// Solidity: function getAuthorizationAddress(bytes32 _projectId) view returns(address _authorizationAddress)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) GetAuthorizationAddress(_projectId [32]byte) (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GetAuthorizationAddress(&_TeeWalletProjectManager.CallOpts, _projectId)
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

// GetExtensionId is a free data retrieval call binding the contract method 0x13def0e3.
//
// Solidity: function getExtensionId(bytes32 _projectId) view returns(uint256 _extensionId)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) GetExtensionId(opts *bind.CallOpts, _projectId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "getExtensionId", _projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExtensionId is a free data retrieval call binding the contract method 0x13def0e3.
//
// Solidity: function getExtensionId(bytes32 _projectId) view returns(uint256 _extensionId)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) GetExtensionId(_projectId [32]byte) (*big.Int, error) {
	return _TeeWalletProjectManager.Contract.GetExtensionId(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetExtensionId is a free data retrieval call binding the contract method 0x13def0e3.
//
// Solidity: function getExtensionId(bytes32 _projectId) view returns(uint256 _extensionId)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) GetExtensionId(_projectId [32]byte) (*big.Int, error) {
	return _TeeWalletProjectManager.Contract.GetExtensionId(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetKeyType is a free data retrieval call binding the contract method 0x779d385a.
//
// Solidity: function getKeyType(bytes32 _projectId) view returns(bytes32 _keyType)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) GetKeyType(opts *bind.CallOpts, _projectId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "getKeyType", _projectId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKeyType is a free data retrieval call binding the contract method 0x779d385a.
//
// Solidity: function getKeyType(bytes32 _projectId) view returns(bytes32 _keyType)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) GetKeyType(_projectId [32]byte) ([32]byte, error) {
	return _TeeWalletProjectManager.Contract.GetKeyType(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetKeyType is a free data retrieval call binding the contract method 0x779d385a.
//
// Solidity: function getKeyType(bytes32 _projectId) view returns(bytes32 _keyType)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) GetKeyType(_projectId [32]byte) ([32]byte, error) {
	return _TeeWalletProjectManager.Contract.GetKeyType(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 _projectId) view returns(address _projectOwner)
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
// Solidity: function getOwner(bytes32 _projectId) view returns(address _projectOwner)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) GetOwner(_projectId [32]byte) (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GetOwner(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetOwner is a free data retrieval call binding the contract method 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 _projectId) view returns(address _projectOwner)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) GetOwner(_projectId [32]byte) (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GetOwner(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetSigningAlgo is a free data retrieval call binding the contract method 0x17566896.
//
// Solidity: function getSigningAlgo(bytes32 _projectId) view returns(bytes32 _signingAlgo)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) GetSigningAlgo(opts *bind.CallOpts, _projectId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "getSigningAlgo", _projectId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSigningAlgo is a free data retrieval call binding the contract method 0x17566896.
//
// Solidity: function getSigningAlgo(bytes32 _projectId) view returns(bytes32 _signingAlgo)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) GetSigningAlgo(_projectId [32]byte) ([32]byte, error) {
	return _TeeWalletProjectManager.Contract.GetSigningAlgo(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// GetSigningAlgo is a free data retrieval call binding the contract method 0x17566896.
//
// Solidity: function getSigningAlgo(bytes32 _projectId) view returns(bytes32 _signingAlgo)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) GetSigningAlgo(_projectId [32]byte) ([32]byte, error) {
	return _TeeWalletProjectManager.Contract.GetSigningAlgo(&_TeeWalletProjectManager.CallOpts, _projectId)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) Governance() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.Governance(&_TeeWalletProjectManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) Governance() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.Governance(&_TeeWalletProjectManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) GovernanceSettings() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GovernanceSettings(&_TeeWalletProjectManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.GovernanceSettings(&_TeeWalletProjectManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) Implementation() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.Implementation(&_TeeWalletProjectManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) Implementation() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.Implementation(&_TeeWalletProjectManager.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeWalletProjectManager.Contract.IsExecutor(&_TeeWalletProjectManager.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeWalletProjectManager.Contract.IsExecutor(&_TeeWalletProjectManager.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) ProductionMode() (bool, error) {
	return _TeeWalletProjectManager.Contract.ProductionMode(&_TeeWalletProjectManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) ProductionMode() (bool, error) {
	return _TeeWalletProjectManager.Contract.ProductionMode(&_TeeWalletProjectManager.CallOpts)
}

// ProjectCounter is a free data retrieval call binding the contract method 0xf8a518ed.
//
// Solidity: function projectCounter() view returns(uint256)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) ProjectCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "projectCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectCounter is a free data retrieval call binding the contract method 0xf8a518ed.
//
// Solidity: function projectCounter() view returns(uint256)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) ProjectCounter() (*big.Int, error) {
	return _TeeWalletProjectManager.Contract.ProjectCounter(&_TeeWalletProjectManager.CallOpts)
}

// ProjectCounter is a free data retrieval call binding the contract method 0xf8a518ed.
//
// Solidity: function projectCounter() view returns(uint256)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) ProjectCounter() (*big.Int, error) {
	return _TeeWalletProjectManager.Contract.ProjectCounter(&_TeeWalletProjectManager.CallOpts)
}

// ProposedProjectOwner is a free data retrieval call binding the contract method 0xc8eb84ba.
//
// Solidity: function proposedProjectOwner(bytes32 projectId) view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) ProposedProjectOwner(opts *bind.CallOpts, projectId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "proposedProjectOwner", projectId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposedProjectOwner is a free data retrieval call binding the contract method 0xc8eb84ba.
//
// Solidity: function proposedProjectOwner(bytes32 projectId) view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) ProposedProjectOwner(projectId [32]byte) (common.Address, error) {
	return _TeeWalletProjectManager.Contract.ProposedProjectOwner(&_TeeWalletProjectManager.CallOpts, projectId)
}

// ProposedProjectOwner is a free data retrieval call binding the contract method 0xc8eb84ba.
//
// Solidity: function proposedProjectOwner(bytes32 projectId) view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) ProposedProjectOwner(projectId [32]byte) (common.Address, error) {
	return _TeeWalletProjectManager.Contract.ProposedProjectOwner(&_TeeWalletProjectManager.CallOpts, projectId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeWalletProjectManager.Contract.ProxiableUUID(&_TeeWalletProjectManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeWalletProjectManager.Contract.ProxiableUUID(&_TeeWalletProjectManager.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) TeeExtensionRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "teeExtensionRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.TeeExtensionRegistry(&_TeeWalletProjectManager.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.TeeExtensionRegistry(&_TeeWalletProjectManager.CallOpts)
}

// TeeOwnerAllowlist is a free data retrieval call binding the contract method 0x327282d4.
//
// Solidity: function teeOwnerAllowlist() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) TeeOwnerAllowlist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "teeOwnerAllowlist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeOwnerAllowlist is a free data retrieval call binding the contract method 0x327282d4.
//
// Solidity: function teeOwnerAllowlist() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) TeeOwnerAllowlist() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.TeeOwnerAllowlist(&_TeeWalletProjectManager.CallOpts)
}

// TeeOwnerAllowlist is a free data retrieval call binding the contract method 0x327282d4.
//
// Solidity: function teeOwnerAllowlist() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) TeeOwnerAllowlist() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.TeeOwnerAllowlist(&_TeeWalletProjectManager.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) TeeWalletManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "teeWalletManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) TeeWalletManager() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.TeeWalletManager(&_TeeWalletProjectManager.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) TeeWalletManager() (common.Address, error) {
	return _TeeWalletProjectManager.Contract.TeeWalletManager(&_TeeWalletProjectManager.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeeWalletProjectManager.contract.Call(opts, &out, "timelockedCalls", selector)

	outstruct := new(struct {
		AllowedAfterTimestamp *big.Int
		EncodedCall           []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowedAfterTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EncodedCall = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeWalletProjectManager.Contract.TimelockedCalls(&_TeeWalletProjectManager.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletProjectManager *TeeWalletProjectManagerCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeWalletProjectManager.Contract.TimelockedCalls(&_TeeWalletProjectManager.CallOpts, selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.CancelGovernanceCall(&_TeeWalletProjectManager.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.CancelGovernanceCall(&_TeeWalletProjectManager.TransactOpts, _selector)
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

// CreateProject is a paid mutator transaction binding the contract method 0xa632bcb6.
//
// Solidity: function createProject(uint256 _extensionId, bytes32 _keyType, bytes32 _signingAlgo, address _authorizationAddress) returns(bytes32 _projectId)
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactor) CreateProject(opts *bind.TransactOpts, _extensionId *big.Int, _keyType [32]byte, _signingAlgo [32]byte, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.contract.Transact(opts, "createProject", _extensionId, _keyType, _signingAlgo, _authorizationAddress)
}

// CreateProject is a paid mutator transaction binding the contract method 0xa632bcb6.
//
// Solidity: function createProject(uint256 _extensionId, bytes32 _keyType, bytes32 _signingAlgo, address _authorizationAddress) returns(bytes32 _projectId)
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) CreateProject(_extensionId *big.Int, _keyType [32]byte, _signingAlgo [32]byte, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.CreateProject(&_TeeWalletProjectManager.TransactOpts, _extensionId, _keyType, _signingAlgo, _authorizationAddress)
}

// CreateProject is a paid mutator transaction binding the contract method 0xa632bcb6.
//
// Solidity: function createProject(uint256 _extensionId, bytes32 _keyType, bytes32 _signingAlgo, address _authorizationAddress) returns(bytes32 _projectId)
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorSession) CreateProject(_extensionId *big.Int, _keyType [32]byte, _signingAlgo [32]byte, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.CreateProject(&_TeeWalletProjectManager.TransactOpts, _extensionId, _keyType, _signingAlgo, _authorizationAddress)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.ExecuteGovernanceCall(&_TeeWalletProjectManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.ExecuteGovernanceCall(&_TeeWalletProjectManager.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.Initialise(&_TeeWalletProjectManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.Initialise(&_TeeWalletProjectManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.Initialize(&_TeeWalletProjectManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.Initialize(&_TeeWalletProjectManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
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

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletProjectManager.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.SwitchToProductionMode(&_TeeWalletProjectManager.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.SwitchToProductionMode(&_TeeWalletProjectManager.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.UpdateContractAddresses(&_TeeWalletProjectManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.UpdateContractAddresses(&_TeeWalletProjectManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.UpgradeToAndCall(&_TeeWalletProjectManager.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeWalletProjectManager *TeeWalletProjectManagerTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeWalletProjectManager.Contract.UpgradeToAndCall(&_TeeWalletProjectManager.TransactOpts, _newImplementation, _data)
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

// TeeWalletProjectManagerGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerGovernanceCallTimelockedIterator struct {
	Event *TeeWalletProjectManagerGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeeWalletProjectManagerGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletProjectManagerGovernanceCallTimelocked)
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
		it.Event = new(TeeWalletProjectManagerGovernanceCallTimelocked)
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
func (it *TeeWalletProjectManagerGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletProjectManagerGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletProjectManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeWalletProjectManagerGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeWalletProjectManager.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerGovernanceCallTimelockedIterator{contract: _TeeWalletProjectManager.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeWalletProjectManagerGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeWalletProjectManager.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletProjectManagerGovernanceCallTimelocked)
				if err := _TeeWalletProjectManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeWalletProjectManagerGovernanceCallTimelocked, error) {
	event := new(TeeWalletProjectManagerGovernanceCallTimelocked)
	if err := _TeeWalletProjectManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletProjectManagerGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerGovernanceInitialisedIterator struct {
	Event *TeeWalletProjectManagerGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeeWalletProjectManagerGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletProjectManagerGovernanceInitialised)
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
		it.Event = new(TeeWalletProjectManagerGovernanceInitialised)
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
func (it *TeeWalletProjectManagerGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletProjectManagerGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletProjectManagerGovernanceInitialised represents a GovernanceInitialised event raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeWalletProjectManagerGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeWalletProjectManager.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerGovernanceInitialisedIterator{contract: _TeeWalletProjectManager.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeWalletProjectManagerGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeWalletProjectManager.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletProjectManagerGovernanceInitialised)
				if err := _TeeWalletProjectManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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

// ParseGovernanceInitialised is a log parse operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) ParseGovernanceInitialised(log types.Log) (*TeeWalletProjectManagerGovernanceInitialised, error) {
	event := new(TeeWalletProjectManagerGovernanceInitialised)
	if err := _TeeWalletProjectManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletProjectManagerGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerGovernedProductionModeEnteredIterator struct {
	Event *TeeWalletProjectManagerGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeeWalletProjectManagerGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletProjectManagerGovernedProductionModeEntered)
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
		it.Event = new(TeeWalletProjectManagerGovernedProductionModeEntered)
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
func (it *TeeWalletProjectManagerGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletProjectManagerGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletProjectManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeeWalletProjectManagerGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeeWalletProjectManager.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerGovernedProductionModeEnteredIterator{contract: _TeeWalletProjectManager.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeeWalletProjectManagerGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeeWalletProjectManager.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletProjectManagerGovernedProductionModeEntered)
				if err := _TeeWalletProjectManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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

// ParseGovernedProductionModeEntered is a log parse operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeeWalletProjectManagerGovernedProductionModeEntered, error) {
	event := new(TeeWalletProjectManagerGovernedProductionModeEntered)
	if err := _TeeWalletProjectManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
	ProjectId            [32]byte
	Owner                common.Address
	ExtensionId          *big.Int
	KeyType              [32]byte
	SigningAlgo          [32]byte
	AuthorizationAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterProjectCreated is a free log retrieval operation binding the contract event 0xcc296665527b84fb4e1da4690e3104d656f4cb08a150dc6a856b1af5aa671514.
//
// Solidity: event ProjectCreated(bytes32 indexed projectId, address indexed owner, uint256 extensionId, bytes32 keyType, bytes32 signingAlgo, address authorizationAddress)
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

// WatchProjectCreated is a free log subscription operation binding the contract event 0xcc296665527b84fb4e1da4690e3104d656f4cb08a150dc6a856b1af5aa671514.
//
// Solidity: event ProjectCreated(bytes32 indexed projectId, address indexed owner, uint256 extensionId, bytes32 keyType, bytes32 signingAlgo, address authorizationAddress)
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

// ParseProjectCreated is a log parse operation binding the contract event 0xcc296665527b84fb4e1da4690e3104d656f4cb08a150dc6a856b1af5aa671514.
//
// Solidity: event ProjectCreated(bytes32 indexed projectId, address indexed owner, uint256 extensionId, bytes32 keyType, bytes32 signingAlgo, address authorizationAddress)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) ParseProjectCreated(log types.Log) (*TeeWalletProjectManagerProjectCreated, error) {
	event := new(TeeWalletProjectManagerProjectCreated)
	if err := _TeeWalletProjectManager.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletProjectManagerTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerTimelockedGovernanceCallCanceledIterator struct {
	Event *TeeWalletProjectManagerTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeeWalletProjectManagerTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletProjectManagerTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeeWalletProjectManagerTimelockedGovernanceCallCanceled)
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
func (it *TeeWalletProjectManagerTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletProjectManagerTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletProjectManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeeWalletProjectManagerTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeeWalletProjectManager.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerTimelockedGovernanceCallCanceledIterator{contract: _TeeWalletProjectManager.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeeWalletProjectManagerTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeeWalletProjectManager.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletProjectManagerTimelockedGovernanceCallCanceled)
				if err := _TeeWalletProjectManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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

// ParseTimelockedGovernanceCallCanceled is a log parse operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeeWalletProjectManagerTimelockedGovernanceCallCanceled, error) {
	event := new(TeeWalletProjectManagerTimelockedGovernanceCallCanceled)
	if err := _TeeWalletProjectManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletProjectManagerTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerTimelockedGovernanceCallExecutedIterator struct {
	Event *TeeWalletProjectManagerTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeeWalletProjectManagerTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletProjectManagerTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeeWalletProjectManagerTimelockedGovernanceCallExecuted)
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
func (it *TeeWalletProjectManagerTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletProjectManagerTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletProjectManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeeWalletProjectManagerTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeeWalletProjectManager.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerTimelockedGovernanceCallExecutedIterator{contract: _TeeWalletProjectManager.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeeWalletProjectManagerTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeeWalletProjectManager.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletProjectManagerTimelockedGovernanceCallExecuted)
				if err := _TeeWalletProjectManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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

// ParseTimelockedGovernanceCallExecuted is a log parse operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeeWalletProjectManagerTimelockedGovernanceCallExecuted, error) {
	event := new(TeeWalletProjectManagerTimelockedGovernanceCallExecuted)
	if err := _TeeWalletProjectManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletProjectManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerUpgradedIterator struct {
	Event *TeeWalletProjectManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *TeeWalletProjectManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletProjectManagerUpgraded)
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
		it.Event = new(TeeWalletProjectManagerUpgraded)
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
func (it *TeeWalletProjectManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletProjectManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletProjectManagerUpgraded represents a Upgraded event raised by the TeeWalletProjectManager contract.
type TeeWalletProjectManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeeWalletProjectManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeWalletProjectManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletProjectManagerUpgradedIterator{contract: _TeeWalletProjectManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeeWalletProjectManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeWalletProjectManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletProjectManagerUpgraded)
				if err := _TeeWalletProjectManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeWalletProjectManager *TeeWalletProjectManagerFilterer) ParseUpgraded(log types.Log) (*TeeWalletProjectManagerUpgraded, error) {
	event := new(TeeWalletProjectManagerUpgraded)
	if err := _TeeWalletProjectManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
