// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package externaladdresses

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

// ExternalAddressesMetaData contains all meta data concerning the ExternalAddresses contract.
var ExternalAddressesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"fdc2Hub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2Verification\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relay\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ExternalAddressesABI is the input ABI used to generate the binding from.
// Deprecated: Use ExternalAddressesMetaData.ABI instead.
var ExternalAddressesABI = ExternalAddressesMetaData.ABI

// ExternalAddresses is an auto generated Go binding around an Ethereum contract.
type ExternalAddresses struct {
	ExternalAddressesCaller     // Read-only binding to the contract
	ExternalAddressesTransactor // Write-only binding to the contract
	ExternalAddressesFilterer   // Log filterer for contract events
}

// ExternalAddressesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExternalAddressesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExternalAddressesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExternalAddressesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExternalAddressesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExternalAddressesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExternalAddressesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExternalAddressesSession struct {
	Contract     *ExternalAddresses // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ExternalAddressesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExternalAddressesCallerSession struct {
	Contract *ExternalAddressesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ExternalAddressesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExternalAddressesTransactorSession struct {
	Contract     *ExternalAddressesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ExternalAddressesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExternalAddressesRaw struct {
	Contract *ExternalAddresses // Generic contract binding to access the raw methods on
}

// ExternalAddressesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExternalAddressesCallerRaw struct {
	Contract *ExternalAddressesCaller // Generic read-only contract binding to access the raw methods on
}

// ExternalAddressesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExternalAddressesTransactorRaw struct {
	Contract *ExternalAddressesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExternalAddresses creates a new instance of ExternalAddresses, bound to a specific deployed contract.
func NewExternalAddresses(address common.Address, backend bind.ContractBackend) (*ExternalAddresses, error) {
	contract, err := bindExternalAddresses(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExternalAddresses{ExternalAddressesCaller: ExternalAddressesCaller{contract: contract}, ExternalAddressesTransactor: ExternalAddressesTransactor{contract: contract}, ExternalAddressesFilterer: ExternalAddressesFilterer{contract: contract}}, nil
}

// NewExternalAddressesCaller creates a new read-only instance of ExternalAddresses, bound to a specific deployed contract.
func NewExternalAddressesCaller(address common.Address, caller bind.ContractCaller) (*ExternalAddressesCaller, error) {
	contract, err := bindExternalAddresses(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExternalAddressesCaller{contract: contract}, nil
}

// NewExternalAddressesTransactor creates a new write-only instance of ExternalAddresses, bound to a specific deployed contract.
func NewExternalAddressesTransactor(address common.Address, transactor bind.ContractTransactor) (*ExternalAddressesTransactor, error) {
	contract, err := bindExternalAddresses(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExternalAddressesTransactor{contract: contract}, nil
}

// NewExternalAddressesFilterer creates a new log filterer instance of ExternalAddresses, bound to a specific deployed contract.
func NewExternalAddressesFilterer(address common.Address, filterer bind.ContractFilterer) (*ExternalAddressesFilterer, error) {
	contract, err := bindExternalAddresses(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExternalAddressesFilterer{contract: contract}, nil
}

// bindExternalAddresses binds a generic wrapper to an already deployed contract.
func bindExternalAddresses(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExternalAddressesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExternalAddresses *ExternalAddressesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExternalAddresses.Contract.ExternalAddressesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExternalAddresses *ExternalAddressesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternalAddresses.Contract.ExternalAddressesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExternalAddresses *ExternalAddressesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExternalAddresses.Contract.ExternalAddressesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExternalAddresses *ExternalAddressesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExternalAddresses.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExternalAddresses *ExternalAddressesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExternalAddresses.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExternalAddresses *ExternalAddressesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExternalAddresses.Contract.contract.Transact(opts, method, params...)
}

// Fdc2Hub is a free data retrieval call binding the contract method 0xa7566ff3.
//
// Solidity: function fdc2Hub() view returns(address)
func (_ExternalAddresses *ExternalAddressesCaller) Fdc2Hub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternalAddresses.contract.Call(opts, &out, "fdc2Hub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fdc2Hub is a free data retrieval call binding the contract method 0xa7566ff3.
//
// Solidity: function fdc2Hub() view returns(address)
func (_ExternalAddresses *ExternalAddressesSession) Fdc2Hub() (common.Address, error) {
	return _ExternalAddresses.Contract.Fdc2Hub(&_ExternalAddresses.CallOpts)
}

// Fdc2Hub is a free data retrieval call binding the contract method 0xa7566ff3.
//
// Solidity: function fdc2Hub() view returns(address)
func (_ExternalAddresses *ExternalAddressesCallerSession) Fdc2Hub() (common.Address, error) {
	return _ExternalAddresses.Contract.Fdc2Hub(&_ExternalAddresses.CallOpts)
}

// Fdc2Verification is a free data retrieval call binding the contract method 0xbf2e9839.
//
// Solidity: function fdc2Verification() view returns(address)
func (_ExternalAddresses *ExternalAddressesCaller) Fdc2Verification(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternalAddresses.contract.Call(opts, &out, "fdc2Verification")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fdc2Verification is a free data retrieval call binding the contract method 0xbf2e9839.
//
// Solidity: function fdc2Verification() view returns(address)
func (_ExternalAddresses *ExternalAddressesSession) Fdc2Verification() (common.Address, error) {
	return _ExternalAddresses.Contract.Fdc2Verification(&_ExternalAddresses.CallOpts)
}

// Fdc2Verification is a free data retrieval call binding the contract method 0xbf2e9839.
//
// Solidity: function fdc2Verification() view returns(address)
func (_ExternalAddresses *ExternalAddressesCallerSession) Fdc2Verification() (common.Address, error) {
	return _ExternalAddresses.Contract.Fdc2Verification(&_ExternalAddresses.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_ExternalAddresses *ExternalAddressesCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternalAddresses.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_ExternalAddresses *ExternalAddressesSession) FlareSystemsManager() (common.Address, error) {
	return _ExternalAddresses.Contract.FlareSystemsManager(&_ExternalAddresses.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_ExternalAddresses *ExternalAddressesCallerSession) FlareSystemsManager() (common.Address, error) {
	return _ExternalAddresses.Contract.FlareSystemsManager(&_ExternalAddresses.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_ExternalAddresses *ExternalAddressesCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternalAddresses.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_ExternalAddresses *ExternalAddressesSession) GetAddressUpdater() (common.Address, error) {
	return _ExternalAddresses.Contract.GetAddressUpdater(&_ExternalAddresses.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_ExternalAddresses *ExternalAddressesCallerSession) GetAddressUpdater() (common.Address, error) {
	return _ExternalAddresses.Contract.GetAddressUpdater(&_ExternalAddresses.CallOpts)
}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_ExternalAddresses *ExternalAddressesCaller) Relay(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternalAddresses.contract.Call(opts, &out, "relay")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_ExternalAddresses *ExternalAddressesSession) Relay() (common.Address, error) {
	return _ExternalAddresses.Contract.Relay(&_ExternalAddresses.CallOpts)
}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_ExternalAddresses *ExternalAddressesCallerSession) Relay() (common.Address, error) {
	return _ExternalAddresses.Contract.Relay(&_ExternalAddresses.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_ExternalAddresses *ExternalAddressesCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExternalAddresses.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_ExternalAddresses *ExternalAddressesSession) RewardManager() (common.Address, error) {
	return _ExternalAddresses.Contract.RewardManager(&_ExternalAddresses.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_ExternalAddresses *ExternalAddressesCallerSession) RewardManager() (common.Address, error) {
	return _ExternalAddresses.Contract.RewardManager(&_ExternalAddresses.CallOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_ExternalAddresses *ExternalAddressesTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _ExternalAddresses.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_ExternalAddresses *ExternalAddressesSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _ExternalAddresses.Contract.UpdateContractAddresses(&_ExternalAddresses.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_ExternalAddresses *ExternalAddressesTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _ExternalAddresses.Contract.UpdateContractAddresses(&_ExternalAddresses.TransactOpts, _contractNameHashes, _contractAddresses)
}
