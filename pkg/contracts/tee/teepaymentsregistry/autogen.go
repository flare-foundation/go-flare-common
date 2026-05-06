// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepaymentsregistry

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

// ITeePaymentsRegistrySourceRegistration is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsRegistrySourceRegistration struct {
	SourceId    [32]byte
	TeePayments common.Address
}

// TeePaymentsRegistryMetaData contains all meta data concerning the TeePaymentsRegistry contract.
var TeePaymentsRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"teePayments\",\"type\":\"address\"}],\"name\":\"SourceAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceIdZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SourceNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teePayments\",\"type\":\"address\"}],\"name\":\"TeePaymentsNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"teePayments\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structITeePaymentsRegistry.SourceRegistration[]\",\"name\":\"registrations\",\"type\":\"tuple[]\"}],\"name\":\"SourcesRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"sourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"SourcesUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredSourceIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredTeePaymentsContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teePayments\",\"type\":\"address\"}],\"name\":\"getSourceIdsForTeePayments\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getTeePaymentsForSource\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"isSourceRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"teePayments\",\"type\":\"address\"}],\"internalType\":\"structITeePaymentsRegistry.SourceRegistration[]\",\"name\":\"_registrations\",\"type\":\"tuple[]\"}],\"name\":\"registerSources\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_sourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"unregisterSources\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeePaymentsRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeePaymentsRegistryMetaData.ABI instead.
var TeePaymentsRegistryABI = TeePaymentsRegistryMetaData.ABI

// TeePaymentsRegistry is an auto generated Go binding around an Ethereum contract.
type TeePaymentsRegistry struct {
	TeePaymentsRegistryCaller     // Read-only binding to the contract
	TeePaymentsRegistryTransactor // Write-only binding to the contract
	TeePaymentsRegistryFilterer   // Log filterer for contract events
}

// TeePaymentsRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeePaymentsRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeePaymentsRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeePaymentsRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeePaymentsRegistrySession struct {
	Contract     *TeePaymentsRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TeePaymentsRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeePaymentsRegistryCallerSession struct {
	Contract *TeePaymentsRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TeePaymentsRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeePaymentsRegistryTransactorSession struct {
	Contract     *TeePaymentsRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TeePaymentsRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeePaymentsRegistryRaw struct {
	Contract *TeePaymentsRegistry // Generic contract binding to access the raw methods on
}

// TeePaymentsRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeePaymentsRegistryCallerRaw struct {
	Contract *TeePaymentsRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TeePaymentsRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeePaymentsRegistryTransactorRaw struct {
	Contract *TeePaymentsRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeePaymentsRegistry creates a new instance of TeePaymentsRegistry, bound to a specific deployed contract.
func NewTeePaymentsRegistry(address common.Address, backend bind.ContractBackend) (*TeePaymentsRegistry, error) {
	contract, err := bindTeePaymentsRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsRegistry{TeePaymentsRegistryCaller: TeePaymentsRegistryCaller{contract: contract}, TeePaymentsRegistryTransactor: TeePaymentsRegistryTransactor{contract: contract}, TeePaymentsRegistryFilterer: TeePaymentsRegistryFilterer{contract: contract}}, nil
}

// NewTeePaymentsRegistryCaller creates a new read-only instance of TeePaymentsRegistry, bound to a specific deployed contract.
func NewTeePaymentsRegistryCaller(address common.Address, caller bind.ContractCaller) (*TeePaymentsRegistryCaller, error) {
	contract, err := bindTeePaymentsRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsRegistryCaller{contract: contract}, nil
}

// NewTeePaymentsRegistryTransactor creates a new write-only instance of TeePaymentsRegistry, bound to a specific deployed contract.
func NewTeePaymentsRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TeePaymentsRegistryTransactor, error) {
	contract, err := bindTeePaymentsRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsRegistryTransactor{contract: contract}, nil
}

// NewTeePaymentsRegistryFilterer creates a new log filterer instance of TeePaymentsRegistry, bound to a specific deployed contract.
func NewTeePaymentsRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TeePaymentsRegistryFilterer, error) {
	contract, err := bindTeePaymentsRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsRegistryFilterer{contract: contract}, nil
}

// bindTeePaymentsRegistry binds a generic wrapper to an already deployed contract.
func bindTeePaymentsRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeePaymentsRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePaymentsRegistry *TeePaymentsRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePaymentsRegistry.Contract.TeePaymentsRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePaymentsRegistry *TeePaymentsRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.TeePaymentsRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePaymentsRegistry *TeePaymentsRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.TeePaymentsRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePaymentsRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePaymentsRegistry.Contract.UPGRADEINTERFACEVERSION(&_TeePaymentsRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePaymentsRegistry.Contract.UPGRADEINTERFACEVERSION(&_TeePaymentsRegistry.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) GetAddressUpdater() (common.Address, error) {
	return _TeePaymentsRegistry.Contract.GetAddressUpdater(&_TeePaymentsRegistry.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeePaymentsRegistry.Contract.GetAddressUpdater(&_TeePaymentsRegistry.CallOpts)
}

// GetRegisteredSourceIds is a free data retrieval call binding the contract method 0xec02c8e0.
//
// Solidity: function getRegisteredSourceIds() view returns(bytes32[])
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) GetRegisteredSourceIds(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "getRegisteredSourceIds")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetRegisteredSourceIds is a free data retrieval call binding the contract method 0xec02c8e0.
//
// Solidity: function getRegisteredSourceIds() view returns(bytes32[])
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) GetRegisteredSourceIds() ([][32]byte, error) {
	return _TeePaymentsRegistry.Contract.GetRegisteredSourceIds(&_TeePaymentsRegistry.CallOpts)
}

// GetRegisteredSourceIds is a free data retrieval call binding the contract method 0xec02c8e0.
//
// Solidity: function getRegisteredSourceIds() view returns(bytes32[])
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) GetRegisteredSourceIds() ([][32]byte, error) {
	return _TeePaymentsRegistry.Contract.GetRegisteredSourceIds(&_TeePaymentsRegistry.CallOpts)
}

// GetRegisteredTeePaymentsContracts is a free data retrieval call binding the contract method 0xf2e5f7c1.
//
// Solidity: function getRegisteredTeePaymentsContracts() view returns(address[])
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) GetRegisteredTeePaymentsContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "getRegisteredTeePaymentsContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredTeePaymentsContracts is a free data retrieval call binding the contract method 0xf2e5f7c1.
//
// Solidity: function getRegisteredTeePaymentsContracts() view returns(address[])
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) GetRegisteredTeePaymentsContracts() ([]common.Address, error) {
	return _TeePaymentsRegistry.Contract.GetRegisteredTeePaymentsContracts(&_TeePaymentsRegistry.CallOpts)
}

// GetRegisteredTeePaymentsContracts is a free data retrieval call binding the contract method 0xf2e5f7c1.
//
// Solidity: function getRegisteredTeePaymentsContracts() view returns(address[])
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) GetRegisteredTeePaymentsContracts() ([]common.Address, error) {
	return _TeePaymentsRegistry.Contract.GetRegisteredTeePaymentsContracts(&_TeePaymentsRegistry.CallOpts)
}

// GetSourceIdsForTeePayments is a free data retrieval call binding the contract method 0xe4665bd0.
//
// Solidity: function getSourceIdsForTeePayments(address _teePayments) view returns(bytes32[])
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) GetSourceIdsForTeePayments(opts *bind.CallOpts, _teePayments common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "getSourceIdsForTeePayments", _teePayments)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSourceIdsForTeePayments is a free data retrieval call binding the contract method 0xe4665bd0.
//
// Solidity: function getSourceIdsForTeePayments(address _teePayments) view returns(bytes32[])
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) GetSourceIdsForTeePayments(_teePayments common.Address) ([][32]byte, error) {
	return _TeePaymentsRegistry.Contract.GetSourceIdsForTeePayments(&_TeePaymentsRegistry.CallOpts, _teePayments)
}

// GetSourceIdsForTeePayments is a free data retrieval call binding the contract method 0xe4665bd0.
//
// Solidity: function getSourceIdsForTeePayments(address _teePayments) view returns(bytes32[])
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) GetSourceIdsForTeePayments(_teePayments common.Address) ([][32]byte, error) {
	return _TeePaymentsRegistry.Contract.GetSourceIdsForTeePayments(&_TeePaymentsRegistry.CallOpts, _teePayments)
}

// GetTeePaymentsForSource is a free data retrieval call binding the contract method 0x4c1db547.
//
// Solidity: function getTeePaymentsForSource(bytes32 _sourceId) view returns(address)
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) GetTeePaymentsForSource(opts *bind.CallOpts, _sourceId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "getTeePaymentsForSource", _sourceId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeePaymentsForSource is a free data retrieval call binding the contract method 0x4c1db547.
//
// Solidity: function getTeePaymentsForSource(bytes32 _sourceId) view returns(address)
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) GetTeePaymentsForSource(_sourceId [32]byte) (common.Address, error) {
	return _TeePaymentsRegistry.Contract.GetTeePaymentsForSource(&_TeePaymentsRegistry.CallOpts, _sourceId)
}

// GetTeePaymentsForSource is a free data retrieval call binding the contract method 0x4c1db547.
//
// Solidity: function getTeePaymentsForSource(bytes32 _sourceId) view returns(address)
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) GetTeePaymentsForSource(_sourceId [32]byte) (common.Address, error) {
	return _TeePaymentsRegistry.Contract.GetTeePaymentsForSource(&_TeePaymentsRegistry.CallOpts, _sourceId)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) Governance() (common.Address, error) {
	return _TeePaymentsRegistry.Contract.Governance(&_TeePaymentsRegistry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) Governance() (common.Address, error) {
	return _TeePaymentsRegistry.Contract.Governance(&_TeePaymentsRegistry.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) GovernanceSettings() (common.Address, error) {
	return _TeePaymentsRegistry.Contract.GovernanceSettings(&_TeePaymentsRegistry.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeePaymentsRegistry.Contract.GovernanceSettings(&_TeePaymentsRegistry.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) Implementation() (common.Address, error) {
	return _TeePaymentsRegistry.Contract.Implementation(&_TeePaymentsRegistry.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) Implementation() (common.Address, error) {
	return _TeePaymentsRegistry.Contract.Implementation(&_TeePaymentsRegistry.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePaymentsRegistry.Contract.IsExecutor(&_TeePaymentsRegistry.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePaymentsRegistry.Contract.IsExecutor(&_TeePaymentsRegistry.CallOpts, _address)
}

// IsSourceRegistered is a free data retrieval call binding the contract method 0x6848bfa4.
//
// Solidity: function isSourceRegistered(bytes32 _sourceId) view returns(bool)
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) IsSourceRegistered(opts *bind.CallOpts, _sourceId [32]byte) (bool, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "isSourceRegistered", _sourceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSourceRegistered is a free data retrieval call binding the contract method 0x6848bfa4.
//
// Solidity: function isSourceRegistered(bytes32 _sourceId) view returns(bool)
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) IsSourceRegistered(_sourceId [32]byte) (bool, error) {
	return _TeePaymentsRegistry.Contract.IsSourceRegistered(&_TeePaymentsRegistry.CallOpts, _sourceId)
}

// IsSourceRegistered is a free data retrieval call binding the contract method 0x6848bfa4.
//
// Solidity: function isSourceRegistered(bytes32 _sourceId) view returns(bool)
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) IsSourceRegistered(_sourceId [32]byte) (bool, error) {
	return _TeePaymentsRegistry.Contract.IsSourceRegistered(&_TeePaymentsRegistry.CallOpts, _sourceId)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) ProductionMode() (bool, error) {
	return _TeePaymentsRegistry.Contract.ProductionMode(&_TeePaymentsRegistry.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) ProductionMode() (bool, error) {
	return _TeePaymentsRegistry.Contract.ProductionMode(&_TeePaymentsRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _TeePaymentsRegistry.Contract.ProxiableUUID(&_TeePaymentsRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeePaymentsRegistry.Contract.ProxiableUUID(&_TeePaymentsRegistry.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePaymentsRegistry *TeePaymentsRegistryCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeePaymentsRegistry.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeePaymentsRegistry.Contract.TimelockedCalls(&_TeePaymentsRegistry.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePaymentsRegistry *TeePaymentsRegistryCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeePaymentsRegistry.Contract.TimelockedCalls(&_TeePaymentsRegistry.CallOpts, selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsRegistry.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.CancelGovernanceCall(&_TeePaymentsRegistry.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.CancelGovernanceCall(&_TeePaymentsRegistry.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsRegistry.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.ExecuteGovernanceCall(&_TeePaymentsRegistry.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.ExecuteGovernanceCall(&_TeePaymentsRegistry.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePaymentsRegistry.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.Initialise(&_TeePaymentsRegistry.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.Initialise(&_TeePaymentsRegistry.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsRegistry.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.Initialize(&_TeePaymentsRegistry.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.Initialize(&_TeePaymentsRegistry.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// RegisterSources is a paid mutator transaction binding the contract method 0xcb8fecb2.
//
// Solidity: function registerSources((bytes32,address)[] _registrations) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactor) RegisterSources(opts *bind.TransactOpts, _registrations []ITeePaymentsRegistrySourceRegistration) (*types.Transaction, error) {
	return _TeePaymentsRegistry.contract.Transact(opts, "registerSources", _registrations)
}

// RegisterSources is a paid mutator transaction binding the contract method 0xcb8fecb2.
//
// Solidity: function registerSources((bytes32,address)[] _registrations) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) RegisterSources(_registrations []ITeePaymentsRegistrySourceRegistration) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.RegisterSources(&_TeePaymentsRegistry.TransactOpts, _registrations)
}

// RegisterSources is a paid mutator transaction binding the contract method 0xcb8fecb2.
//
// Solidity: function registerSources((bytes32,address)[] _registrations) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactorSession) RegisterSources(_registrations []ITeePaymentsRegistrySourceRegistration) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.RegisterSources(&_TeePaymentsRegistry.TransactOpts, _registrations)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsRegistry.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.SwitchToProductionMode(&_TeePaymentsRegistry.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.SwitchToProductionMode(&_TeePaymentsRegistry.TransactOpts)
}

// UnregisterSources is a paid mutator transaction binding the contract method 0x840c3088.
//
// Solidity: function unregisterSources(bytes32[] _sourceIds) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactor) UnregisterSources(opts *bind.TransactOpts, _sourceIds [][32]byte) (*types.Transaction, error) {
	return _TeePaymentsRegistry.contract.Transact(opts, "unregisterSources", _sourceIds)
}

// UnregisterSources is a paid mutator transaction binding the contract method 0x840c3088.
//
// Solidity: function unregisterSources(bytes32[] _sourceIds) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) UnregisterSources(_sourceIds [][32]byte) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.UnregisterSources(&_TeePaymentsRegistry.TransactOpts, _sourceIds)
}

// UnregisterSources is a paid mutator transaction binding the contract method 0x840c3088.
//
// Solidity: function unregisterSources(bytes32[] _sourceIds) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactorSession) UnregisterSources(_sourceIds [][32]byte) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.UnregisterSources(&_TeePaymentsRegistry.TransactOpts, _sourceIds)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsRegistry.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.UpdateContractAddresses(&_TeePaymentsRegistry.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.UpdateContractAddresses(&_TeePaymentsRegistry.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsRegistry.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsRegistry *TeePaymentsRegistrySession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.UpgradeToAndCall(&_TeePaymentsRegistry.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsRegistry *TeePaymentsRegistryTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsRegistry.Contract.UpgradeToAndCall(&_TeePaymentsRegistry.TransactOpts, _newImplementation, _data)
}

// TeePaymentsRegistryGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryGovernanceCallTimelockedIterator struct {
	Event *TeePaymentsRegistryGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeePaymentsRegistryGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsRegistryGovernanceCallTimelocked)
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
		it.Event = new(TeePaymentsRegistryGovernanceCallTimelocked)
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
func (it *TeePaymentsRegistryGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsRegistryGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsRegistryGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeePaymentsRegistryGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsRegistryGovernanceCallTimelockedIterator{contract: _TeePaymentsRegistry.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeePaymentsRegistryGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsRegistryGovernanceCallTimelocked)
				if err := _TeePaymentsRegistry.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeePaymentsRegistryGovernanceCallTimelocked, error) {
	event := new(TeePaymentsRegistryGovernanceCallTimelocked)
	if err := _TeePaymentsRegistry.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsRegistryGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryGovernanceInitialisedIterator struct {
	Event *TeePaymentsRegistryGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeePaymentsRegistryGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsRegistryGovernanceInitialised)
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
		it.Event = new(TeePaymentsRegistryGovernanceInitialised)
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
func (it *TeePaymentsRegistryGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsRegistryGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsRegistryGovernanceInitialised represents a GovernanceInitialised event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeePaymentsRegistryGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsRegistryGovernanceInitialisedIterator{contract: _TeePaymentsRegistry.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeePaymentsRegistryGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsRegistryGovernanceInitialised)
				if err := _TeePaymentsRegistry.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) ParseGovernanceInitialised(log types.Log) (*TeePaymentsRegistryGovernanceInitialised, error) {
	event := new(TeePaymentsRegistryGovernanceInitialised)
	if err := _TeePaymentsRegistry.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsRegistryGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryGovernedProductionModeEnteredIterator struct {
	Event *TeePaymentsRegistryGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeePaymentsRegistryGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsRegistryGovernedProductionModeEntered)
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
		it.Event = new(TeePaymentsRegistryGovernedProductionModeEntered)
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
func (it *TeePaymentsRegistryGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsRegistryGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsRegistryGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeePaymentsRegistryGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsRegistryGovernedProductionModeEnteredIterator{contract: _TeePaymentsRegistry.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeePaymentsRegistryGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsRegistryGovernedProductionModeEntered)
				if err := _TeePaymentsRegistry.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeePaymentsRegistryGovernedProductionModeEntered, error) {
	event := new(TeePaymentsRegistryGovernedProductionModeEntered)
	if err := _TeePaymentsRegistry.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsRegistrySourcesRegisteredIterator is returned from FilterSourcesRegistered and is used to iterate over the raw logs and unpacked data for SourcesRegistered events raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistrySourcesRegisteredIterator struct {
	Event *TeePaymentsRegistrySourcesRegistered // Event containing the contract specifics and raw log

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
func (it *TeePaymentsRegistrySourcesRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsRegistrySourcesRegistered)
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
		it.Event = new(TeePaymentsRegistrySourcesRegistered)
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
func (it *TeePaymentsRegistrySourcesRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsRegistrySourcesRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsRegistrySourcesRegistered represents a SourcesRegistered event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistrySourcesRegistered struct {
	Registrations []ITeePaymentsRegistrySourceRegistration
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSourcesRegistered is a free log retrieval operation binding the contract event 0x1b652ae86c05b7c28a48232c6bc8cf307c19d0c8b1f4116fdea4ca63ce22cafe.
//
// Solidity: event SourcesRegistered((bytes32,address)[] registrations)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) FilterSourcesRegistered(opts *bind.FilterOpts) (*TeePaymentsRegistrySourcesRegisteredIterator, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.FilterLogs(opts, "SourcesRegistered")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsRegistrySourcesRegisteredIterator{contract: _TeePaymentsRegistry.contract, event: "SourcesRegistered", logs: logs, sub: sub}, nil
}

// WatchSourcesRegistered is a free log subscription operation binding the contract event 0x1b652ae86c05b7c28a48232c6bc8cf307c19d0c8b1f4116fdea4ca63ce22cafe.
//
// Solidity: event SourcesRegistered((bytes32,address)[] registrations)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) WatchSourcesRegistered(opts *bind.WatchOpts, sink chan<- *TeePaymentsRegistrySourcesRegistered) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.WatchLogs(opts, "SourcesRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsRegistrySourcesRegistered)
				if err := _TeePaymentsRegistry.contract.UnpackLog(event, "SourcesRegistered", log); err != nil {
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

// ParseSourcesRegistered is a log parse operation binding the contract event 0x1b652ae86c05b7c28a48232c6bc8cf307c19d0c8b1f4116fdea4ca63ce22cafe.
//
// Solidity: event SourcesRegistered((bytes32,address)[] registrations)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) ParseSourcesRegistered(log types.Log) (*TeePaymentsRegistrySourcesRegistered, error) {
	event := new(TeePaymentsRegistrySourcesRegistered)
	if err := _TeePaymentsRegistry.contract.UnpackLog(event, "SourcesRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsRegistrySourcesUnregisteredIterator is returned from FilterSourcesUnregistered and is used to iterate over the raw logs and unpacked data for SourcesUnregistered events raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistrySourcesUnregisteredIterator struct {
	Event *TeePaymentsRegistrySourcesUnregistered // Event containing the contract specifics and raw log

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
func (it *TeePaymentsRegistrySourcesUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsRegistrySourcesUnregistered)
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
		it.Event = new(TeePaymentsRegistrySourcesUnregistered)
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
func (it *TeePaymentsRegistrySourcesUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsRegistrySourcesUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsRegistrySourcesUnregistered represents a SourcesUnregistered event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistrySourcesUnregistered struct {
	SourceIds [][32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSourcesUnregistered is a free log retrieval operation binding the contract event 0xc7965f9e7313aa28af37168d7877889929a52f22396e7c2b0d8a425d859f6565.
//
// Solidity: event SourcesUnregistered(bytes32[] sourceIds)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) FilterSourcesUnregistered(opts *bind.FilterOpts) (*TeePaymentsRegistrySourcesUnregisteredIterator, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.FilterLogs(opts, "SourcesUnregistered")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsRegistrySourcesUnregisteredIterator{contract: _TeePaymentsRegistry.contract, event: "SourcesUnregistered", logs: logs, sub: sub}, nil
}

// WatchSourcesUnregistered is a free log subscription operation binding the contract event 0xc7965f9e7313aa28af37168d7877889929a52f22396e7c2b0d8a425d859f6565.
//
// Solidity: event SourcesUnregistered(bytes32[] sourceIds)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) WatchSourcesUnregistered(opts *bind.WatchOpts, sink chan<- *TeePaymentsRegistrySourcesUnregistered) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.WatchLogs(opts, "SourcesUnregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsRegistrySourcesUnregistered)
				if err := _TeePaymentsRegistry.contract.UnpackLog(event, "SourcesUnregistered", log); err != nil {
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

// ParseSourcesUnregistered is a log parse operation binding the contract event 0xc7965f9e7313aa28af37168d7877889929a52f22396e7c2b0d8a425d859f6565.
//
// Solidity: event SourcesUnregistered(bytes32[] sourceIds)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) ParseSourcesUnregistered(log types.Log) (*TeePaymentsRegistrySourcesUnregistered, error) {
	event := new(TeePaymentsRegistrySourcesUnregistered)
	if err := _TeePaymentsRegistry.contract.UnpackLog(event, "SourcesUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsRegistryTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryTimelockedGovernanceCallCanceledIterator struct {
	Event *TeePaymentsRegistryTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeePaymentsRegistryTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsRegistryTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeePaymentsRegistryTimelockedGovernanceCallCanceled)
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
func (it *TeePaymentsRegistryTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsRegistryTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsRegistryTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeePaymentsRegistryTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsRegistryTimelockedGovernanceCallCanceledIterator{contract: _TeePaymentsRegistry.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeePaymentsRegistryTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsRegistryTimelockedGovernanceCallCanceled)
				if err := _TeePaymentsRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeePaymentsRegistryTimelockedGovernanceCallCanceled, error) {
	event := new(TeePaymentsRegistryTimelockedGovernanceCallCanceled)
	if err := _TeePaymentsRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsRegistryTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryTimelockedGovernanceCallExecutedIterator struct {
	Event *TeePaymentsRegistryTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeePaymentsRegistryTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsRegistryTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeePaymentsRegistryTimelockedGovernanceCallExecuted)
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
func (it *TeePaymentsRegistryTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsRegistryTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsRegistryTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeePaymentsRegistryTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsRegistryTimelockedGovernanceCallExecutedIterator{contract: _TeePaymentsRegistry.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeePaymentsRegistryTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsRegistry.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsRegistryTimelockedGovernanceCallExecuted)
				if err := _TeePaymentsRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeePaymentsRegistryTimelockedGovernanceCallExecuted, error) {
	event := new(TeePaymentsRegistryTimelockedGovernanceCallExecuted)
	if err := _TeePaymentsRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryUpgradedIterator struct {
	Event *TeePaymentsRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *TeePaymentsRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsRegistryUpgraded)
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
		it.Event = new(TeePaymentsRegistryUpgraded)
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
func (it *TeePaymentsRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsRegistryUpgraded represents a Upgraded event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeePaymentsRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePaymentsRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsRegistryUpgradedIterator{contract: _TeePaymentsRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeePaymentsRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePaymentsRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsRegistryUpgraded)
				if err := _TeePaymentsRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeePaymentsRegistry *TeePaymentsRegistryFilterer) ParseUpgraded(log types.Log) (*TeePaymentsRegistryUpgraded, error) {
	event := new(TeePaymentsRegistryUpgraded)
	if err := _TeePaymentsRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
