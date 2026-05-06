// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fdc2requestfeeconfigurations

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

// Fdc2RequestFeeConfigurationsMetaData contains all meta data concerning the Fdc2RequestFeeConfigurations contract.
var Fdc2RequestFeeConfigurationsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeMustBeGreaterThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TypeAndSourceCombinationNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"source\",\"type\":\"bytes32\"}],\"name\":\"TypeAndSourceFeeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"source\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TypeAndSourceFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_type\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_source\",\"type\":\"bytes32\"}],\"name\":\"getTypeAndSourceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_type\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_source\",\"type\":\"bytes32\"}],\"name\":\"removeTypeAndSourceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_types\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sources\",\"type\":\"bytes32[]\"}],\"name\":\"removeTypeAndSourceFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_type\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_source\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setTypeAndSourceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_types\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sources\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"}],\"name\":\"setTypeAndSourceFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// Fdc2RequestFeeConfigurationsABI is the input ABI used to generate the binding from.
// Deprecated: Use Fdc2RequestFeeConfigurationsMetaData.ABI instead.
var Fdc2RequestFeeConfigurationsABI = Fdc2RequestFeeConfigurationsMetaData.ABI

// Fdc2RequestFeeConfigurations is an auto generated Go binding around an Ethereum contract.
type Fdc2RequestFeeConfigurations struct {
	Fdc2RequestFeeConfigurationsCaller     // Read-only binding to the contract
	Fdc2RequestFeeConfigurationsTransactor // Write-only binding to the contract
	Fdc2RequestFeeConfigurationsFilterer   // Log filterer for contract events
}

// Fdc2RequestFeeConfigurationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type Fdc2RequestFeeConfigurationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fdc2RequestFeeConfigurationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Fdc2RequestFeeConfigurationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fdc2RequestFeeConfigurationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Fdc2RequestFeeConfigurationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fdc2RequestFeeConfigurationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Fdc2RequestFeeConfigurationsSession struct {
	Contract     *Fdc2RequestFeeConfigurations // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// Fdc2RequestFeeConfigurationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Fdc2RequestFeeConfigurationsCallerSession struct {
	Contract *Fdc2RequestFeeConfigurationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// Fdc2RequestFeeConfigurationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Fdc2RequestFeeConfigurationsTransactorSession struct {
	Contract     *Fdc2RequestFeeConfigurationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// Fdc2RequestFeeConfigurationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type Fdc2RequestFeeConfigurationsRaw struct {
	Contract *Fdc2RequestFeeConfigurations // Generic contract binding to access the raw methods on
}

// Fdc2RequestFeeConfigurationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Fdc2RequestFeeConfigurationsCallerRaw struct {
	Contract *Fdc2RequestFeeConfigurationsCaller // Generic read-only contract binding to access the raw methods on
}

// Fdc2RequestFeeConfigurationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Fdc2RequestFeeConfigurationsTransactorRaw struct {
	Contract *Fdc2RequestFeeConfigurationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFdc2RequestFeeConfigurations creates a new instance of Fdc2RequestFeeConfigurations, bound to a specific deployed contract.
func NewFdc2RequestFeeConfigurations(address common.Address, backend bind.ContractBackend) (*Fdc2RequestFeeConfigurations, error) {
	contract, err := bindFdc2RequestFeeConfigurations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fdc2RequestFeeConfigurations{Fdc2RequestFeeConfigurationsCaller: Fdc2RequestFeeConfigurationsCaller{contract: contract}, Fdc2RequestFeeConfigurationsTransactor: Fdc2RequestFeeConfigurationsTransactor{contract: contract}, Fdc2RequestFeeConfigurationsFilterer: Fdc2RequestFeeConfigurationsFilterer{contract: contract}}, nil
}

// NewFdc2RequestFeeConfigurationsCaller creates a new read-only instance of Fdc2RequestFeeConfigurations, bound to a specific deployed contract.
func NewFdc2RequestFeeConfigurationsCaller(address common.Address, caller bind.ContractCaller) (*Fdc2RequestFeeConfigurationsCaller, error) {
	contract, err := bindFdc2RequestFeeConfigurations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Fdc2RequestFeeConfigurationsCaller{contract: contract}, nil
}

// NewFdc2RequestFeeConfigurationsTransactor creates a new write-only instance of Fdc2RequestFeeConfigurations, bound to a specific deployed contract.
func NewFdc2RequestFeeConfigurationsTransactor(address common.Address, transactor bind.ContractTransactor) (*Fdc2RequestFeeConfigurationsTransactor, error) {
	contract, err := bindFdc2RequestFeeConfigurations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Fdc2RequestFeeConfigurationsTransactor{contract: contract}, nil
}

// NewFdc2RequestFeeConfigurationsFilterer creates a new log filterer instance of Fdc2RequestFeeConfigurations, bound to a specific deployed contract.
func NewFdc2RequestFeeConfigurationsFilterer(address common.Address, filterer bind.ContractFilterer) (*Fdc2RequestFeeConfigurationsFilterer, error) {
	contract, err := bindFdc2RequestFeeConfigurations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Fdc2RequestFeeConfigurationsFilterer{contract: contract}, nil
}

// bindFdc2RequestFeeConfigurations binds a generic wrapper to an already deployed contract.
func bindFdc2RequestFeeConfigurations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Fdc2RequestFeeConfigurationsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fdc2RequestFeeConfigurations.Contract.Fdc2RequestFeeConfigurationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.Fdc2RequestFeeConfigurationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.Fdc2RequestFeeConfigurationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fdc2RequestFeeConfigurations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Fdc2RequestFeeConfigurations.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Fdc2RequestFeeConfigurations.Contract.UPGRADEINTERFACEVERSION(&_Fdc2RequestFeeConfigurations.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Fdc2RequestFeeConfigurations.Contract.UPGRADEINTERFACEVERSION(&_Fdc2RequestFeeConfigurations.CallOpts)
}

// GetTypeAndSourceFee is a free data retrieval call binding the contract method 0x5ebe55b3.
//
// Solidity: function getTypeAndSourceFee(bytes32 _type, bytes32 _source) view returns(uint256 _fee)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCaller) GetTypeAndSourceFee(opts *bind.CallOpts, _type [32]byte, _source [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Fdc2RequestFeeConfigurations.contract.Call(opts, &out, "getTypeAndSourceFee", _type, _source)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTypeAndSourceFee is a free data retrieval call binding the contract method 0x5ebe55b3.
//
// Solidity: function getTypeAndSourceFee(bytes32 _type, bytes32 _source) view returns(uint256 _fee)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) GetTypeAndSourceFee(_type [32]byte, _source [32]byte) (*big.Int, error) {
	return _Fdc2RequestFeeConfigurations.Contract.GetTypeAndSourceFee(&_Fdc2RequestFeeConfigurations.CallOpts, _type, _source)
}

// GetTypeAndSourceFee is a free data retrieval call binding the contract method 0x5ebe55b3.
//
// Solidity: function getTypeAndSourceFee(bytes32 _type, bytes32 _source) view returns(uint256 _fee)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCallerSession) GetTypeAndSourceFee(_type [32]byte, _source [32]byte) (*big.Int, error) {
	return _Fdc2RequestFeeConfigurations.Contract.GetTypeAndSourceFee(&_Fdc2RequestFeeConfigurations.CallOpts, _type, _source)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fdc2RequestFeeConfigurations.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) Governance() (common.Address, error) {
	return _Fdc2RequestFeeConfigurations.Contract.Governance(&_Fdc2RequestFeeConfigurations.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCallerSession) Governance() (common.Address, error) {
	return _Fdc2RequestFeeConfigurations.Contract.Governance(&_Fdc2RequestFeeConfigurations.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fdc2RequestFeeConfigurations.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) GovernanceSettings() (common.Address, error) {
	return _Fdc2RequestFeeConfigurations.Contract.GovernanceSettings(&_Fdc2RequestFeeConfigurations.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCallerSession) GovernanceSettings() (common.Address, error) {
	return _Fdc2RequestFeeConfigurations.Contract.GovernanceSettings(&_Fdc2RequestFeeConfigurations.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fdc2RequestFeeConfigurations.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) Implementation() (common.Address, error) {
	return _Fdc2RequestFeeConfigurations.Contract.Implementation(&_Fdc2RequestFeeConfigurations.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCallerSession) Implementation() (common.Address, error) {
	return _Fdc2RequestFeeConfigurations.Contract.Implementation(&_Fdc2RequestFeeConfigurations.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Fdc2RequestFeeConfigurations.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) IsExecutor(_address common.Address) (bool, error) {
	return _Fdc2RequestFeeConfigurations.Contract.IsExecutor(&_Fdc2RequestFeeConfigurations.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _Fdc2RequestFeeConfigurations.Contract.IsExecutor(&_Fdc2RequestFeeConfigurations.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Fdc2RequestFeeConfigurations.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) ProductionMode() (bool, error) {
	return _Fdc2RequestFeeConfigurations.Contract.ProductionMode(&_Fdc2RequestFeeConfigurations.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCallerSession) ProductionMode() (bool, error) {
	return _Fdc2RequestFeeConfigurations.Contract.ProductionMode(&_Fdc2RequestFeeConfigurations.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fdc2RequestFeeConfigurations.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) ProxiableUUID() ([32]byte, error) {
	return _Fdc2RequestFeeConfigurations.Contract.ProxiableUUID(&_Fdc2RequestFeeConfigurations.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Fdc2RequestFeeConfigurations.Contract.ProxiableUUID(&_Fdc2RequestFeeConfigurations.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _Fdc2RequestFeeConfigurations.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _Fdc2RequestFeeConfigurations.Contract.TimelockedCalls(&_Fdc2RequestFeeConfigurations.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _Fdc2RequestFeeConfigurations.Contract.TimelockedCalls(&_Fdc2RequestFeeConfigurations.CallOpts, selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.CancelGovernanceCall(&_Fdc2RequestFeeConfigurations.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.CancelGovernanceCall(&_Fdc2RequestFeeConfigurations.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.ExecuteGovernanceCall(&_Fdc2RequestFeeConfigurations.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.ExecuteGovernanceCall(&_Fdc2RequestFeeConfigurations.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.Initialise(&_Fdc2RequestFeeConfigurations.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.Initialise(&_Fdc2RequestFeeConfigurations.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.Initialize(&_Fdc2RequestFeeConfigurations.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.Initialize(&_Fdc2RequestFeeConfigurations.TransactOpts, _governanceSettings, _initialGovernance)
}

// RemoveTypeAndSourceFee is a paid mutator transaction binding the contract method 0xda42a778.
//
// Solidity: function removeTypeAndSourceFee(bytes32 _type, bytes32 _source) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactor) RemoveTypeAndSourceFee(opts *bind.TransactOpts, _type [32]byte, _source [32]byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.contract.Transact(opts, "removeTypeAndSourceFee", _type, _source)
}

// RemoveTypeAndSourceFee is a paid mutator transaction binding the contract method 0xda42a778.
//
// Solidity: function removeTypeAndSourceFee(bytes32 _type, bytes32 _source) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) RemoveTypeAndSourceFee(_type [32]byte, _source [32]byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.RemoveTypeAndSourceFee(&_Fdc2RequestFeeConfigurations.TransactOpts, _type, _source)
}

// RemoveTypeAndSourceFee is a paid mutator transaction binding the contract method 0xda42a778.
//
// Solidity: function removeTypeAndSourceFee(bytes32 _type, bytes32 _source) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactorSession) RemoveTypeAndSourceFee(_type [32]byte, _source [32]byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.RemoveTypeAndSourceFee(&_Fdc2RequestFeeConfigurations.TransactOpts, _type, _source)
}

// RemoveTypeAndSourceFees is a paid mutator transaction binding the contract method 0xe2f6bca8.
//
// Solidity: function removeTypeAndSourceFees(bytes32[] _types, bytes32[] _sources) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactor) RemoveTypeAndSourceFees(opts *bind.TransactOpts, _types [][32]byte, _sources [][32]byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.contract.Transact(opts, "removeTypeAndSourceFees", _types, _sources)
}

// RemoveTypeAndSourceFees is a paid mutator transaction binding the contract method 0xe2f6bca8.
//
// Solidity: function removeTypeAndSourceFees(bytes32[] _types, bytes32[] _sources) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) RemoveTypeAndSourceFees(_types [][32]byte, _sources [][32]byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.RemoveTypeAndSourceFees(&_Fdc2RequestFeeConfigurations.TransactOpts, _types, _sources)
}

// RemoveTypeAndSourceFees is a paid mutator transaction binding the contract method 0xe2f6bca8.
//
// Solidity: function removeTypeAndSourceFees(bytes32[] _types, bytes32[] _sources) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactorSession) RemoveTypeAndSourceFees(_types [][32]byte, _sources [][32]byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.RemoveTypeAndSourceFees(&_Fdc2RequestFeeConfigurations.TransactOpts, _types, _sources)
}

// SetTypeAndSourceFee is a paid mutator transaction binding the contract method 0x5b1bacb7.
//
// Solidity: function setTypeAndSourceFee(bytes32 _type, bytes32 _source, uint256 _fee) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactor) SetTypeAndSourceFee(opts *bind.TransactOpts, _type [32]byte, _source [32]byte, _fee *big.Int) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.contract.Transact(opts, "setTypeAndSourceFee", _type, _source, _fee)
}

// SetTypeAndSourceFee is a paid mutator transaction binding the contract method 0x5b1bacb7.
//
// Solidity: function setTypeAndSourceFee(bytes32 _type, bytes32 _source, uint256 _fee) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) SetTypeAndSourceFee(_type [32]byte, _source [32]byte, _fee *big.Int) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.SetTypeAndSourceFee(&_Fdc2RequestFeeConfigurations.TransactOpts, _type, _source, _fee)
}

// SetTypeAndSourceFee is a paid mutator transaction binding the contract method 0x5b1bacb7.
//
// Solidity: function setTypeAndSourceFee(bytes32 _type, bytes32 _source, uint256 _fee) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactorSession) SetTypeAndSourceFee(_type [32]byte, _source [32]byte, _fee *big.Int) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.SetTypeAndSourceFee(&_Fdc2RequestFeeConfigurations.TransactOpts, _type, _source, _fee)
}

// SetTypeAndSourceFees is a paid mutator transaction binding the contract method 0x3da84157.
//
// Solidity: function setTypeAndSourceFees(bytes32[] _types, bytes32[] _sources, uint256[] _fees) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactor) SetTypeAndSourceFees(opts *bind.TransactOpts, _types [][32]byte, _sources [][32]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.contract.Transact(opts, "setTypeAndSourceFees", _types, _sources, _fees)
}

// SetTypeAndSourceFees is a paid mutator transaction binding the contract method 0x3da84157.
//
// Solidity: function setTypeAndSourceFees(bytes32[] _types, bytes32[] _sources, uint256[] _fees) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) SetTypeAndSourceFees(_types [][32]byte, _sources [][32]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.SetTypeAndSourceFees(&_Fdc2RequestFeeConfigurations.TransactOpts, _types, _sources, _fees)
}

// SetTypeAndSourceFees is a paid mutator transaction binding the contract method 0x3da84157.
//
// Solidity: function setTypeAndSourceFees(bytes32[] _types, bytes32[] _sources, uint256[] _fees) returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactorSession) SetTypeAndSourceFees(_types [][32]byte, _sources [][32]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.SetTypeAndSourceFees(&_Fdc2RequestFeeConfigurations.TransactOpts, _types, _sources, _fees)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.SwitchToProductionMode(&_Fdc2RequestFeeConfigurations.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.SwitchToProductionMode(&_Fdc2RequestFeeConfigurations.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.UpgradeToAndCall(&_Fdc2RequestFeeConfigurations.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _Fdc2RequestFeeConfigurations.Contract.UpgradeToAndCall(&_Fdc2RequestFeeConfigurations.TransactOpts, _newImplementation, _data)
}

// Fdc2RequestFeeConfigurationsGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsGovernanceCallTimelockedIterator struct {
	Event *Fdc2RequestFeeConfigurationsGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *Fdc2RequestFeeConfigurationsGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2RequestFeeConfigurationsGovernanceCallTimelocked)
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
		it.Event = new(Fdc2RequestFeeConfigurationsGovernanceCallTimelocked)
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
func (it *Fdc2RequestFeeConfigurationsGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2RequestFeeConfigurationsGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2RequestFeeConfigurationsGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*Fdc2RequestFeeConfigurationsGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &Fdc2RequestFeeConfigurationsGovernanceCallTimelockedIterator{contract: _Fdc2RequestFeeConfigurations.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *Fdc2RequestFeeConfigurationsGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2RequestFeeConfigurationsGovernanceCallTimelocked)
				if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) ParseGovernanceCallTimelocked(log types.Log) (*Fdc2RequestFeeConfigurationsGovernanceCallTimelocked, error) {
	event := new(Fdc2RequestFeeConfigurationsGovernanceCallTimelocked)
	if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2RequestFeeConfigurationsGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsGovernanceInitialisedIterator struct {
	Event *Fdc2RequestFeeConfigurationsGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *Fdc2RequestFeeConfigurationsGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2RequestFeeConfigurationsGovernanceInitialised)
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
		it.Event = new(Fdc2RequestFeeConfigurationsGovernanceInitialised)
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
func (it *Fdc2RequestFeeConfigurationsGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2RequestFeeConfigurationsGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2RequestFeeConfigurationsGovernanceInitialised represents a GovernanceInitialised event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*Fdc2RequestFeeConfigurationsGovernanceInitialisedIterator, error) {

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &Fdc2RequestFeeConfigurationsGovernanceInitialisedIterator{contract: _Fdc2RequestFeeConfigurations.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *Fdc2RequestFeeConfigurationsGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2RequestFeeConfigurationsGovernanceInitialised)
				if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) ParseGovernanceInitialised(log types.Log) (*Fdc2RequestFeeConfigurationsGovernanceInitialised, error) {
	event := new(Fdc2RequestFeeConfigurationsGovernanceInitialised)
	if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2RequestFeeConfigurationsGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsGovernedProductionModeEnteredIterator struct {
	Event *Fdc2RequestFeeConfigurationsGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *Fdc2RequestFeeConfigurationsGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2RequestFeeConfigurationsGovernedProductionModeEntered)
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
		it.Event = new(Fdc2RequestFeeConfigurationsGovernedProductionModeEntered)
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
func (it *Fdc2RequestFeeConfigurationsGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2RequestFeeConfigurationsGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2RequestFeeConfigurationsGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*Fdc2RequestFeeConfigurationsGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &Fdc2RequestFeeConfigurationsGovernedProductionModeEnteredIterator{contract: _Fdc2RequestFeeConfigurations.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *Fdc2RequestFeeConfigurationsGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2RequestFeeConfigurationsGovernedProductionModeEntered)
				if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) ParseGovernedProductionModeEntered(log types.Log) (*Fdc2RequestFeeConfigurationsGovernedProductionModeEntered, error) {
	event := new(Fdc2RequestFeeConfigurationsGovernedProductionModeEntered)
	if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator struct {
	Event *Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled)
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
		it.Event = new(Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled)
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
func (it *Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator{contract: _Fdc2RequestFeeConfigurations.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled)
				if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled, error) {
	event := new(Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled)
	if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator struct {
	Event *Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted)
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
		it.Event = new(Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted)
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
func (it *Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator{contract: _Fdc2RequestFeeConfigurations.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted)
				if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted, error) {
	event := new(Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted)
	if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemovedIterator is returned from FilterTypeAndSourceFeeRemoved and is used to iterate over the raw logs and unpacked data for TypeAndSourceFeeRemoved events raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemovedIterator struct {
	Event *Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved // Event containing the contract specifics and raw log

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
func (it *Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved)
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
		it.Event = new(Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved)
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
func (it *Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved represents a TypeAndSourceFeeRemoved event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved struct {
	AttestationType [32]byte
	Source          [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTypeAndSourceFeeRemoved is a free log retrieval operation binding the contract event 0xa186aa06a988f2cd41162e05c078a9f69e58eb1a7233361c2604a1a49721d521.
//
// Solidity: event TypeAndSourceFeeRemoved(bytes32 indexed attestationType, bytes32 indexed source)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) FilterTypeAndSourceFeeRemoved(opts *bind.FilterOpts, attestationType [][32]byte, source [][32]byte) (*Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemovedIterator, error) {

	var attestationTypeRule []interface{}
	for _, attestationTypeItem := range attestationType {
		attestationTypeRule = append(attestationTypeRule, attestationTypeItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.FilterLogs(opts, "TypeAndSourceFeeRemoved", attestationTypeRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return &Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemovedIterator{contract: _Fdc2RequestFeeConfigurations.contract, event: "TypeAndSourceFeeRemoved", logs: logs, sub: sub}, nil
}

// WatchTypeAndSourceFeeRemoved is a free log subscription operation binding the contract event 0xa186aa06a988f2cd41162e05c078a9f69e58eb1a7233361c2604a1a49721d521.
//
// Solidity: event TypeAndSourceFeeRemoved(bytes32 indexed attestationType, bytes32 indexed source)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) WatchTypeAndSourceFeeRemoved(opts *bind.WatchOpts, sink chan<- *Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved, attestationType [][32]byte, source [][32]byte) (event.Subscription, error) {

	var attestationTypeRule []interface{}
	for _, attestationTypeItem := range attestationType {
		attestationTypeRule = append(attestationTypeRule, attestationTypeItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.WatchLogs(opts, "TypeAndSourceFeeRemoved", attestationTypeRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved)
				if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "TypeAndSourceFeeRemoved", log); err != nil {
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

// ParseTypeAndSourceFeeRemoved is a log parse operation binding the contract event 0xa186aa06a988f2cd41162e05c078a9f69e58eb1a7233361c2604a1a49721d521.
//
// Solidity: event TypeAndSourceFeeRemoved(bytes32 indexed attestationType, bytes32 indexed source)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) ParseTypeAndSourceFeeRemoved(log types.Log) (*Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved, error) {
	event := new(Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved)
	if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "TypeAndSourceFeeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2RequestFeeConfigurationsTypeAndSourceFeeSetIterator is returned from FilterTypeAndSourceFeeSet and is used to iterate over the raw logs and unpacked data for TypeAndSourceFeeSet events raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTypeAndSourceFeeSetIterator struct {
	Event *Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet // Event containing the contract specifics and raw log

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
func (it *Fdc2RequestFeeConfigurationsTypeAndSourceFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet)
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
		it.Event = new(Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet)
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
func (it *Fdc2RequestFeeConfigurationsTypeAndSourceFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2RequestFeeConfigurationsTypeAndSourceFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet represents a TypeAndSourceFeeSet event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet struct {
	AttestationType [32]byte
	Source          [32]byte
	Fee             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTypeAndSourceFeeSet is a free log retrieval operation binding the contract event 0x53fbcb0688fd3904cd33601e03a5eba0c6bba217d8c9312333ea0c116ea91382.
//
// Solidity: event TypeAndSourceFeeSet(bytes32 indexed attestationType, bytes32 indexed source, uint256 fee)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) FilterTypeAndSourceFeeSet(opts *bind.FilterOpts, attestationType [][32]byte, source [][32]byte) (*Fdc2RequestFeeConfigurationsTypeAndSourceFeeSetIterator, error) {

	var attestationTypeRule []interface{}
	for _, attestationTypeItem := range attestationType {
		attestationTypeRule = append(attestationTypeRule, attestationTypeItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.FilterLogs(opts, "TypeAndSourceFeeSet", attestationTypeRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return &Fdc2RequestFeeConfigurationsTypeAndSourceFeeSetIterator{contract: _Fdc2RequestFeeConfigurations.contract, event: "TypeAndSourceFeeSet", logs: logs, sub: sub}, nil
}

// WatchTypeAndSourceFeeSet is a free log subscription operation binding the contract event 0x53fbcb0688fd3904cd33601e03a5eba0c6bba217d8c9312333ea0c116ea91382.
//
// Solidity: event TypeAndSourceFeeSet(bytes32 indexed attestationType, bytes32 indexed source, uint256 fee)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) WatchTypeAndSourceFeeSet(opts *bind.WatchOpts, sink chan<- *Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet, attestationType [][32]byte, source [][32]byte) (event.Subscription, error) {

	var attestationTypeRule []interface{}
	for _, attestationTypeItem := range attestationType {
		attestationTypeRule = append(attestationTypeRule, attestationTypeItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.WatchLogs(opts, "TypeAndSourceFeeSet", attestationTypeRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet)
				if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "TypeAndSourceFeeSet", log); err != nil {
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

// ParseTypeAndSourceFeeSet is a log parse operation binding the contract event 0x53fbcb0688fd3904cd33601e03a5eba0c6bba217d8c9312333ea0c116ea91382.
//
// Solidity: event TypeAndSourceFeeSet(bytes32 indexed attestationType, bytes32 indexed source, uint256 fee)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) ParseTypeAndSourceFeeSet(log types.Log) (*Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet, error) {
	event := new(Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet)
	if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "TypeAndSourceFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2RequestFeeConfigurationsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsUpgradedIterator struct {
	Event *Fdc2RequestFeeConfigurationsUpgraded // Event containing the contract specifics and raw log

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
func (it *Fdc2RequestFeeConfigurationsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2RequestFeeConfigurationsUpgraded)
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
		it.Event = new(Fdc2RequestFeeConfigurationsUpgraded)
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
func (it *Fdc2RequestFeeConfigurationsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2RequestFeeConfigurationsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2RequestFeeConfigurationsUpgraded represents a Upgraded event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*Fdc2RequestFeeConfigurationsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &Fdc2RequestFeeConfigurationsUpgradedIterator{contract: _Fdc2RequestFeeConfigurations.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *Fdc2RequestFeeConfigurationsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Fdc2RequestFeeConfigurations.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2RequestFeeConfigurationsUpgraded)
				if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurationsFilterer) ParseUpgraded(log types.Log) (*Fdc2RequestFeeConfigurationsUpgraded, error) {
	event := new(Fdc2RequestFeeConfigurationsUpgraded)
	if err := _Fdc2RequestFeeConfigurations.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
