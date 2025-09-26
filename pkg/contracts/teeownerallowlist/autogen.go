// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teeownerallowlist

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

// TeeOwnerAllowlistMetaData contains all meta data concerning the TeeOwnerAllowlist contract.
var TeeOwnerAllowlistMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAlreadyAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeMachineOwnersAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeWalletProjectOwnersAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeMachineOwnersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeWalletProjectOwnersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"addAllowedTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"addAllowedTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"allTeeMachineOwnersAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"allTeeWalletProjectOwnersAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allowAllTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allowAllTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getAllowedTeeMachineOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getAllowedTeeWalletProjectOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"isAllowedTeeMachineOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"isAllowedTeeWalletProjectOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeExtensionRegistry\",\"outputs\":[{\"internalType\":\"contractITeeExtensionRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeeOwnerAllowlistABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeOwnerAllowlistMetaData.ABI instead.
var TeeOwnerAllowlistABI = TeeOwnerAllowlistMetaData.ABI

// TeeOwnerAllowlist is an auto generated Go binding around an Ethereum contract.
type TeeOwnerAllowlist struct {
	TeeOwnerAllowlistCaller     // Read-only binding to the contract
	TeeOwnerAllowlistTransactor // Write-only binding to the contract
	TeeOwnerAllowlistFilterer   // Log filterer for contract events
}

// TeeOwnerAllowlistCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeOwnerAllowlistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeOwnerAllowlistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeOwnerAllowlistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeOwnerAllowlistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeOwnerAllowlistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeOwnerAllowlistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeOwnerAllowlistSession struct {
	Contract     *TeeOwnerAllowlist // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TeeOwnerAllowlistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeOwnerAllowlistCallerSession struct {
	Contract *TeeOwnerAllowlistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TeeOwnerAllowlistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeOwnerAllowlistTransactorSession struct {
	Contract     *TeeOwnerAllowlistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TeeOwnerAllowlistRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeOwnerAllowlistRaw struct {
	Contract *TeeOwnerAllowlist // Generic contract binding to access the raw methods on
}

// TeeOwnerAllowlistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeOwnerAllowlistCallerRaw struct {
	Contract *TeeOwnerAllowlistCaller // Generic read-only contract binding to access the raw methods on
}

// TeeOwnerAllowlistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeOwnerAllowlistTransactorRaw struct {
	Contract *TeeOwnerAllowlistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeOwnerAllowlist creates a new instance of TeeOwnerAllowlist, bound to a specific deployed contract.
func NewTeeOwnerAllowlist(address common.Address, backend bind.ContractBackend) (*TeeOwnerAllowlist, error) {
	contract, err := bindTeeOwnerAllowlist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlist{TeeOwnerAllowlistCaller: TeeOwnerAllowlistCaller{contract: contract}, TeeOwnerAllowlistTransactor: TeeOwnerAllowlistTransactor{contract: contract}, TeeOwnerAllowlistFilterer: TeeOwnerAllowlistFilterer{contract: contract}}, nil
}

// NewTeeOwnerAllowlistCaller creates a new read-only instance of TeeOwnerAllowlist, bound to a specific deployed contract.
func NewTeeOwnerAllowlistCaller(address common.Address, caller bind.ContractCaller) (*TeeOwnerAllowlistCaller, error) {
	contract, err := bindTeeOwnerAllowlist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistCaller{contract: contract}, nil
}

// NewTeeOwnerAllowlistTransactor creates a new write-only instance of TeeOwnerAllowlist, bound to a specific deployed contract.
func NewTeeOwnerAllowlistTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeOwnerAllowlistTransactor, error) {
	contract, err := bindTeeOwnerAllowlist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistTransactor{contract: contract}, nil
}

// NewTeeOwnerAllowlistFilterer creates a new log filterer instance of TeeOwnerAllowlist, bound to a specific deployed contract.
func NewTeeOwnerAllowlistFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeOwnerAllowlistFilterer, error) {
	contract, err := bindTeeOwnerAllowlist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistFilterer{contract: contract}, nil
}

// bindTeeOwnerAllowlist binds a generic wrapper to an already deployed contract.
func bindTeeOwnerAllowlist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeOwnerAllowlistMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeOwnerAllowlist *TeeOwnerAllowlistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeOwnerAllowlist.Contract.TeeOwnerAllowlistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeOwnerAllowlist *TeeOwnerAllowlistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.TeeOwnerAllowlistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeOwnerAllowlist *TeeOwnerAllowlistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.TeeOwnerAllowlistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeOwnerAllowlist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeOwnerAllowlist.Contract.UPGRADEINTERFACEVERSION(&_TeeOwnerAllowlist.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeOwnerAllowlist.Contract.UPGRADEINTERFACEVERSION(&_TeeOwnerAllowlist.CallOpts)
}

// AllTeeMachineOwnersAllowed is a free data retrieval call binding the contract method 0x2a30b9f4.
//
// Solidity: function allTeeMachineOwnersAllowed(uint256 extensionId) view returns(bool)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) AllTeeMachineOwnersAllowed(opts *bind.CallOpts, extensionId *big.Int) (bool, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "allTeeMachineOwnersAllowed", extensionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllTeeMachineOwnersAllowed is a free data retrieval call binding the contract method 0x2a30b9f4.
//
// Solidity: function allTeeMachineOwnersAllowed(uint256 extensionId) view returns(bool)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) AllTeeMachineOwnersAllowed(extensionId *big.Int) (bool, error) {
	return _TeeOwnerAllowlist.Contract.AllTeeMachineOwnersAllowed(&_TeeOwnerAllowlist.CallOpts, extensionId)
}

// AllTeeMachineOwnersAllowed is a free data retrieval call binding the contract method 0x2a30b9f4.
//
// Solidity: function allTeeMachineOwnersAllowed(uint256 extensionId) view returns(bool)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) AllTeeMachineOwnersAllowed(extensionId *big.Int) (bool, error) {
	return _TeeOwnerAllowlist.Contract.AllTeeMachineOwnersAllowed(&_TeeOwnerAllowlist.CallOpts, extensionId)
}

// AllTeeWalletProjectOwnersAllowed is a free data retrieval call binding the contract method 0x8398fc2e.
//
// Solidity: function allTeeWalletProjectOwnersAllowed(uint256 extensionId) view returns(bool)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) AllTeeWalletProjectOwnersAllowed(opts *bind.CallOpts, extensionId *big.Int) (bool, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "allTeeWalletProjectOwnersAllowed", extensionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllTeeWalletProjectOwnersAllowed is a free data retrieval call binding the contract method 0x8398fc2e.
//
// Solidity: function allTeeWalletProjectOwnersAllowed(uint256 extensionId) view returns(bool)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) AllTeeWalletProjectOwnersAllowed(extensionId *big.Int) (bool, error) {
	return _TeeOwnerAllowlist.Contract.AllTeeWalletProjectOwnersAllowed(&_TeeOwnerAllowlist.CallOpts, extensionId)
}

// AllTeeWalletProjectOwnersAllowed is a free data retrieval call binding the contract method 0x8398fc2e.
//
// Solidity: function allTeeWalletProjectOwnersAllowed(uint256 extensionId) view returns(bool)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) AllTeeWalletProjectOwnersAllowed(extensionId *big.Int) (bool, error) {
	return _TeeOwnerAllowlist.Contract.AllTeeWalletProjectOwnersAllowed(&_TeeOwnerAllowlist.CallOpts, extensionId)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) GetAddressUpdater() (common.Address, error) {
	return _TeeOwnerAllowlist.Contract.GetAddressUpdater(&_TeeOwnerAllowlist.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeOwnerAllowlist.Contract.GetAddressUpdater(&_TeeOwnerAllowlist.CallOpts)
}

// GetAllowedTeeMachineOwners is a free data retrieval call binding the contract method 0x5f37ea30.
//
// Solidity: function getAllowedTeeMachineOwners(uint256 _extensionId) view returns(address[])
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) GetAllowedTeeMachineOwners(opts *bind.CallOpts, _extensionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "getAllowedTeeMachineOwners", _extensionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowedTeeMachineOwners is a free data retrieval call binding the contract method 0x5f37ea30.
//
// Solidity: function getAllowedTeeMachineOwners(uint256 _extensionId) view returns(address[])
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) GetAllowedTeeMachineOwners(_extensionId *big.Int) ([]common.Address, error) {
	return _TeeOwnerAllowlist.Contract.GetAllowedTeeMachineOwners(&_TeeOwnerAllowlist.CallOpts, _extensionId)
}

// GetAllowedTeeMachineOwners is a free data retrieval call binding the contract method 0x5f37ea30.
//
// Solidity: function getAllowedTeeMachineOwners(uint256 _extensionId) view returns(address[])
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) GetAllowedTeeMachineOwners(_extensionId *big.Int) ([]common.Address, error) {
	return _TeeOwnerAllowlist.Contract.GetAllowedTeeMachineOwners(&_TeeOwnerAllowlist.CallOpts, _extensionId)
}

// GetAllowedTeeWalletProjectOwners is a free data retrieval call binding the contract method 0xd11340ed.
//
// Solidity: function getAllowedTeeWalletProjectOwners(uint256 _extensionId) view returns(address[])
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) GetAllowedTeeWalletProjectOwners(opts *bind.CallOpts, _extensionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "getAllowedTeeWalletProjectOwners", _extensionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowedTeeWalletProjectOwners is a free data retrieval call binding the contract method 0xd11340ed.
//
// Solidity: function getAllowedTeeWalletProjectOwners(uint256 _extensionId) view returns(address[])
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) GetAllowedTeeWalletProjectOwners(_extensionId *big.Int) ([]common.Address, error) {
	return _TeeOwnerAllowlist.Contract.GetAllowedTeeWalletProjectOwners(&_TeeOwnerAllowlist.CallOpts, _extensionId)
}

// GetAllowedTeeWalletProjectOwners is a free data retrieval call binding the contract method 0xd11340ed.
//
// Solidity: function getAllowedTeeWalletProjectOwners(uint256 _extensionId) view returns(address[])
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) GetAllowedTeeWalletProjectOwners(_extensionId *big.Int) ([]common.Address, error) {
	return _TeeOwnerAllowlist.Contract.GetAllowedTeeWalletProjectOwners(&_TeeOwnerAllowlist.CallOpts, _extensionId)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) Governance() (common.Address, error) {
	return _TeeOwnerAllowlist.Contract.Governance(&_TeeOwnerAllowlist.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) Governance() (common.Address, error) {
	return _TeeOwnerAllowlist.Contract.Governance(&_TeeOwnerAllowlist.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) GovernanceSettings() (common.Address, error) {
	return _TeeOwnerAllowlist.Contract.GovernanceSettings(&_TeeOwnerAllowlist.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeeOwnerAllowlist.Contract.GovernanceSettings(&_TeeOwnerAllowlist.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) Implementation() (common.Address, error) {
	return _TeeOwnerAllowlist.Contract.Implementation(&_TeeOwnerAllowlist.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) Implementation() (common.Address, error) {
	return _TeeOwnerAllowlist.Contract.Implementation(&_TeeOwnerAllowlist.CallOpts)
}

// IsAllowedTeeMachineOwner is a free data retrieval call binding the contract method 0x84f5da36.
//
// Solidity: function isAllowedTeeMachineOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) IsAllowedTeeMachineOwner(opts *bind.CallOpts, _extensionId *big.Int, _owner common.Address) (bool, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "isAllowedTeeMachineOwner", _extensionId, _owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedTeeMachineOwner is a free data retrieval call binding the contract method 0x84f5da36.
//
// Solidity: function isAllowedTeeMachineOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) IsAllowedTeeMachineOwner(_extensionId *big.Int, _owner common.Address) (bool, error) {
	return _TeeOwnerAllowlist.Contract.IsAllowedTeeMachineOwner(&_TeeOwnerAllowlist.CallOpts, _extensionId, _owner)
}

// IsAllowedTeeMachineOwner is a free data retrieval call binding the contract method 0x84f5da36.
//
// Solidity: function isAllowedTeeMachineOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) IsAllowedTeeMachineOwner(_extensionId *big.Int, _owner common.Address) (bool, error) {
	return _TeeOwnerAllowlist.Contract.IsAllowedTeeMachineOwner(&_TeeOwnerAllowlist.CallOpts, _extensionId, _owner)
}

// IsAllowedTeeWalletProjectOwner is a free data retrieval call binding the contract method 0xab1c7ea1.
//
// Solidity: function isAllowedTeeWalletProjectOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) IsAllowedTeeWalletProjectOwner(opts *bind.CallOpts, _extensionId *big.Int, _owner common.Address) (bool, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "isAllowedTeeWalletProjectOwner", _extensionId, _owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedTeeWalletProjectOwner is a free data retrieval call binding the contract method 0xab1c7ea1.
//
// Solidity: function isAllowedTeeWalletProjectOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) IsAllowedTeeWalletProjectOwner(_extensionId *big.Int, _owner common.Address) (bool, error) {
	return _TeeOwnerAllowlist.Contract.IsAllowedTeeWalletProjectOwner(&_TeeOwnerAllowlist.CallOpts, _extensionId, _owner)
}

// IsAllowedTeeWalletProjectOwner is a free data retrieval call binding the contract method 0xab1c7ea1.
//
// Solidity: function isAllowedTeeWalletProjectOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) IsAllowedTeeWalletProjectOwner(_extensionId *big.Int, _owner common.Address) (bool, error) {
	return _TeeOwnerAllowlist.Contract.IsAllowedTeeWalletProjectOwner(&_TeeOwnerAllowlist.CallOpts, _extensionId, _owner)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeOwnerAllowlist.Contract.IsExecutor(&_TeeOwnerAllowlist.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeOwnerAllowlist.Contract.IsExecutor(&_TeeOwnerAllowlist.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) ProductionMode() (bool, error) {
	return _TeeOwnerAllowlist.Contract.ProductionMode(&_TeeOwnerAllowlist.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) ProductionMode() (bool, error) {
	return _TeeOwnerAllowlist.Contract.ProductionMode(&_TeeOwnerAllowlist.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) ProxiableUUID() ([32]byte, error) {
	return _TeeOwnerAllowlist.Contract.ProxiableUUID(&_TeeOwnerAllowlist.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeOwnerAllowlist.Contract.ProxiableUUID(&_TeeOwnerAllowlist.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) TeeExtensionRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "teeExtensionRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeOwnerAllowlist.Contract.TeeExtensionRegistry(&_TeeOwnerAllowlist.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeOwnerAllowlist.Contract.TeeExtensionRegistry(&_TeeOwnerAllowlist.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeOwnerAllowlist.Contract.TimelockedCalls(&_TeeOwnerAllowlist.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeOwnerAllowlist.Contract.TimelockedCalls(&_TeeOwnerAllowlist.CallOpts, selector)
}

// AddAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0xa8f3cab7.
//
// Solidity: function addAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) AddAllowedTeeMachineOwners(opts *bind.TransactOpts, _extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "addAllowedTeeMachineOwners", _extensionId, _owners)
}

// AddAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0xa8f3cab7.
//
// Solidity: function addAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) AddAllowedTeeMachineOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AddAllowedTeeMachineOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// AddAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0xa8f3cab7.
//
// Solidity: function addAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) AddAllowedTeeMachineOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AddAllowedTeeMachineOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// AddAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0xd1c34746.
//
// Solidity: function addAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) AddAllowedTeeWalletProjectOwners(opts *bind.TransactOpts, _extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "addAllowedTeeWalletProjectOwners", _extensionId, _owners)
}

// AddAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0xd1c34746.
//
// Solidity: function addAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) AddAllowedTeeWalletProjectOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AddAllowedTeeWalletProjectOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// AddAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0xd1c34746.
//
// Solidity: function addAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) AddAllowedTeeWalletProjectOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AddAllowedTeeWalletProjectOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// AllowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0x1b489f2d.
//
// Solidity: function allowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) AllowAllTeeMachineOwners(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "allowAllTeeMachineOwners", _extensionId)
}

// AllowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0x1b489f2d.
//
// Solidity: function allowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) AllowAllTeeMachineOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AllowAllTeeMachineOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId)
}

// AllowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0x1b489f2d.
//
// Solidity: function allowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) AllowAllTeeMachineOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AllowAllTeeMachineOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId)
}

// AllowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x61f957e7.
//
// Solidity: function allowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) AllowAllTeeWalletProjectOwners(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "allowAllTeeWalletProjectOwners", _extensionId)
}

// AllowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x61f957e7.
//
// Solidity: function allowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) AllowAllTeeWalletProjectOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AllowAllTeeWalletProjectOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId)
}

// AllowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x61f957e7.
//
// Solidity: function allowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) AllowAllTeeWalletProjectOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AllowAllTeeWalletProjectOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.CancelGovernanceCall(&_TeeOwnerAllowlist.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.CancelGovernanceCall(&_TeeOwnerAllowlist.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.ExecuteGovernanceCall(&_TeeOwnerAllowlist.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.ExecuteGovernanceCall(&_TeeOwnerAllowlist.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.Initialise(&_TeeOwnerAllowlist.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.Initialise(&_TeeOwnerAllowlist.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.Initialize(&_TeeOwnerAllowlist.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.Initialize(&_TeeOwnerAllowlist.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.SwitchToProductionMode(&_TeeOwnerAllowlist.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.SwitchToProductionMode(&_TeeOwnerAllowlist.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.UpdateContractAddresses(&_TeeOwnerAllowlist.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.UpdateContractAddresses(&_TeeOwnerAllowlist.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.UpgradeToAndCall(&_TeeOwnerAllowlist.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.UpgradeToAndCall(&_TeeOwnerAllowlist.TransactOpts, _newImplementation, _data)
}

// TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator is returned from FilterAllTeeMachineOwnersAllowed and is used to iterate over the raw logs and unpacked data for AllTeeMachineOwnersAllowed events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator struct {
	Event *TeeOwnerAllowlistAllTeeMachineOwnersAllowed // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistAllTeeMachineOwnersAllowed)
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
		it.Event = new(TeeOwnerAllowlistAllTeeMachineOwnersAllowed)
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
func (it *TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistAllTeeMachineOwnersAllowed represents a AllTeeMachineOwnersAllowed event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllTeeMachineOwnersAllowed struct {
	ExtensionId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllTeeMachineOwnersAllowed is a free log retrieval operation binding the contract event 0x1bc41c6dd61d3a656898250dfc58a14ddd4735f4480ae4c1581072ad19fcd8a5.
//
// Solidity: event AllTeeMachineOwnersAllowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterAllTeeMachineOwnersAllowed(opts *bind.FilterOpts) (*TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "AllTeeMachineOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator{contract: _TeeOwnerAllowlist.contract, event: "AllTeeMachineOwnersAllowed", logs: logs, sub: sub}, nil
}

// WatchAllTeeMachineOwnersAllowed is a free log subscription operation binding the contract event 0x1bc41c6dd61d3a656898250dfc58a14ddd4735f4480ae4c1581072ad19fcd8a5.
//
// Solidity: event AllTeeMachineOwnersAllowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchAllTeeMachineOwnersAllowed(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistAllTeeMachineOwnersAllowed) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "AllTeeMachineOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistAllTeeMachineOwnersAllowed)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllTeeMachineOwnersAllowed", log); err != nil {
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

// ParseAllTeeMachineOwnersAllowed is a log parse operation binding the contract event 0x1bc41c6dd61d3a656898250dfc58a14ddd4735f4480ae4c1581072ad19fcd8a5.
//
// Solidity: event AllTeeMachineOwnersAllowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseAllTeeMachineOwnersAllowed(log types.Log) (*TeeOwnerAllowlistAllTeeMachineOwnersAllowed, error) {
	event := new(TeeOwnerAllowlistAllTeeMachineOwnersAllowed)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllTeeMachineOwnersAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator is returned from FilterAllTeeWalletProjectOwnersAllowed and is used to iterate over the raw logs and unpacked data for AllTeeWalletProjectOwnersAllowed events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator struct {
	Event *TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed)
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
		it.Event = new(TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed)
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
func (it *TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed represents a AllTeeWalletProjectOwnersAllowed event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed struct {
	ExtensionId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllTeeWalletProjectOwnersAllowed is a free log retrieval operation binding the contract event 0x1936686ee6942840bf2283355eb720261b36284a52b74604253556ee61c08854.
//
// Solidity: event AllTeeWalletProjectOwnersAllowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterAllTeeWalletProjectOwnersAllowed(opts *bind.FilterOpts) (*TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "AllTeeWalletProjectOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator{contract: _TeeOwnerAllowlist.contract, event: "AllTeeWalletProjectOwnersAllowed", logs: logs, sub: sub}, nil
}

// WatchAllTeeWalletProjectOwnersAllowed is a free log subscription operation binding the contract event 0x1936686ee6942840bf2283355eb720261b36284a52b74604253556ee61c08854.
//
// Solidity: event AllTeeWalletProjectOwnersAllowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchAllTeeWalletProjectOwnersAllowed(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "AllTeeWalletProjectOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllTeeWalletProjectOwnersAllowed", log); err != nil {
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

// ParseAllTeeWalletProjectOwnersAllowed is a log parse operation binding the contract event 0x1936686ee6942840bf2283355eb720261b36284a52b74604253556ee61c08854.
//
// Solidity: event AllTeeWalletProjectOwnersAllowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseAllTeeWalletProjectOwnersAllowed(log types.Log) (*TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed, error) {
	event := new(TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllTeeWalletProjectOwnersAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator is returned from FilterAllowedTeeMachineOwnersAdded and is used to iterate over the raw logs and unpacked data for AllowedTeeMachineOwnersAdded events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator struct {
	Event *TeeOwnerAllowlistAllowedTeeMachineOwnersAdded // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistAllowedTeeMachineOwnersAdded)
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
		it.Event = new(TeeOwnerAllowlistAllowedTeeMachineOwnersAdded)
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
func (it *TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistAllowedTeeMachineOwnersAdded represents a AllowedTeeMachineOwnersAdded event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllowedTeeMachineOwnersAdded struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllowedTeeMachineOwnersAdded is a free log retrieval operation binding the contract event 0xd392883c2acc8ffc90b60bb8783db156d990502913d5b159e209cd604475337c.
//
// Solidity: event AllowedTeeMachineOwnersAdded(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterAllowedTeeMachineOwnersAdded(opts *bind.FilterOpts) (*TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "AllowedTeeMachineOwnersAdded")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator{contract: _TeeOwnerAllowlist.contract, event: "AllowedTeeMachineOwnersAdded", logs: logs, sub: sub}, nil
}

// WatchAllowedTeeMachineOwnersAdded is a free log subscription operation binding the contract event 0xd392883c2acc8ffc90b60bb8783db156d990502913d5b159e209cd604475337c.
//
// Solidity: event AllowedTeeMachineOwnersAdded(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchAllowedTeeMachineOwnersAdded(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistAllowedTeeMachineOwnersAdded) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "AllowedTeeMachineOwnersAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistAllowedTeeMachineOwnersAdded)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllowedTeeMachineOwnersAdded", log); err != nil {
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

// ParseAllowedTeeMachineOwnersAdded is a log parse operation binding the contract event 0xd392883c2acc8ffc90b60bb8783db156d990502913d5b159e209cd604475337c.
//
// Solidity: event AllowedTeeMachineOwnersAdded(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseAllowedTeeMachineOwnersAdded(log types.Log) (*TeeOwnerAllowlistAllowedTeeMachineOwnersAdded, error) {
	event := new(TeeOwnerAllowlistAllowedTeeMachineOwnersAdded)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllowedTeeMachineOwnersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator is returned from FilterAllowedTeeWalletProjectOwnersAdded and is used to iterate over the raw logs and unpacked data for AllowedTeeWalletProjectOwnersAdded events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator struct {
	Event *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded)
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
		it.Event = new(TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded)
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
func (it *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded represents a AllowedTeeWalletProjectOwnersAdded event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllowedTeeWalletProjectOwnersAdded is a free log retrieval operation binding the contract event 0x60340c199493a9a48f50ae5af13ebefed3adeeb1c4f0a5476d59654680ffdb0d.
//
// Solidity: event AllowedTeeWalletProjectOwnersAdded(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterAllowedTeeWalletProjectOwnersAdded(opts *bind.FilterOpts) (*TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "AllowedTeeWalletProjectOwnersAdded")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator{contract: _TeeOwnerAllowlist.contract, event: "AllowedTeeWalletProjectOwnersAdded", logs: logs, sub: sub}, nil
}

// WatchAllowedTeeWalletProjectOwnersAdded is a free log subscription operation binding the contract event 0x60340c199493a9a48f50ae5af13ebefed3adeeb1c4f0a5476d59654680ffdb0d.
//
// Solidity: event AllowedTeeWalletProjectOwnersAdded(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchAllowedTeeWalletProjectOwnersAdded(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "AllowedTeeWalletProjectOwnersAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllowedTeeWalletProjectOwnersAdded", log); err != nil {
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

// ParseAllowedTeeWalletProjectOwnersAdded is a log parse operation binding the contract event 0x60340c199493a9a48f50ae5af13ebefed3adeeb1c4f0a5476d59654680ffdb0d.
//
// Solidity: event AllowedTeeWalletProjectOwnersAdded(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseAllowedTeeWalletProjectOwnersAdded(log types.Log) (*TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded, error) {
	event := new(TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllowedTeeWalletProjectOwnersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistGovernanceCallTimelockedIterator struct {
	Event *TeeOwnerAllowlistGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistGovernanceCallTimelocked)
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
		it.Event = new(TeeOwnerAllowlistGovernanceCallTimelocked)
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
func (it *TeeOwnerAllowlistGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeOwnerAllowlistGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistGovernanceCallTimelockedIterator{contract: _TeeOwnerAllowlist.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistGovernanceCallTimelocked)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeOwnerAllowlistGovernanceCallTimelocked, error) {
	event := new(TeeOwnerAllowlistGovernanceCallTimelocked)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistGovernanceInitialisedIterator struct {
	Event *TeeOwnerAllowlistGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistGovernanceInitialised)
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
		it.Event = new(TeeOwnerAllowlistGovernanceInitialised)
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
func (it *TeeOwnerAllowlistGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistGovernanceInitialised represents a GovernanceInitialised event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeOwnerAllowlistGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistGovernanceInitialisedIterator{contract: _TeeOwnerAllowlist.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistGovernanceInitialised)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseGovernanceInitialised(log types.Log) (*TeeOwnerAllowlistGovernanceInitialised, error) {
	event := new(TeeOwnerAllowlistGovernanceInitialised)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistGovernedProductionModeEnteredIterator struct {
	Event *TeeOwnerAllowlistGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistGovernedProductionModeEntered)
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
		it.Event = new(TeeOwnerAllowlistGovernedProductionModeEntered)
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
func (it *TeeOwnerAllowlistGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeeOwnerAllowlistGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistGovernedProductionModeEnteredIterator{contract: _TeeOwnerAllowlist.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistGovernedProductionModeEntered)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeeOwnerAllowlistGovernedProductionModeEntered, error) {
	event := new(TeeOwnerAllowlistGovernedProductionModeEntered)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistTimelockedGovernanceCallCanceledIterator struct {
	Event *TeeOwnerAllowlistTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeeOwnerAllowlistTimelockedGovernanceCallCanceled)
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
func (it *TeeOwnerAllowlistTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeeOwnerAllowlistTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistTimelockedGovernanceCallCanceledIterator{contract: _TeeOwnerAllowlist.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistTimelockedGovernanceCallCanceled)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeeOwnerAllowlistTimelockedGovernanceCallCanceled, error) {
	event := new(TeeOwnerAllowlistTimelockedGovernanceCallCanceled)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistTimelockedGovernanceCallExecutedIterator struct {
	Event *TeeOwnerAllowlistTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeeOwnerAllowlistTimelockedGovernanceCallExecuted)
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
func (it *TeeOwnerAllowlistTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeeOwnerAllowlistTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistTimelockedGovernanceCallExecutedIterator{contract: _TeeOwnerAllowlist.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistTimelockedGovernanceCallExecuted)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeeOwnerAllowlistTimelockedGovernanceCallExecuted, error) {
	event := new(TeeOwnerAllowlistTimelockedGovernanceCallExecuted)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistUpgradedIterator struct {
	Event *TeeOwnerAllowlistUpgraded // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistUpgraded)
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
		it.Event = new(TeeOwnerAllowlistUpgraded)
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
func (it *TeeOwnerAllowlistUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistUpgraded represents a Upgraded event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeeOwnerAllowlistUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistUpgradedIterator{contract: _TeeOwnerAllowlist.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistUpgraded)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseUpgraded(log types.Log) (*TeeOwnerAllowlistUpgraded, error) {
	event := new(TeeOwnerAllowlistUpgraded)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
