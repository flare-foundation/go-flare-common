// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ftdcrequestfeeconfigurations

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

// FtdcRequestFeeConfigurationsMetaData contains all meta data concerning the FtdcRequestFeeConfigurations contract.
var FtdcRequestFeeConfigurationsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeMustBeGreaterThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TypeAndSourceCombinationNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"source\",\"type\":\"bytes32\"}],\"name\":\"TypeAndSourceFeeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"source\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TypeAndSourceFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_type\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_source\",\"type\":\"bytes32\"}],\"name\":\"getTypeAndSourceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_type\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_source\",\"type\":\"bytes32\"}],\"name\":\"removeTypeAndSourceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_types\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sources\",\"type\":\"bytes32[]\"}],\"name\":\"removeTypeAndSourceFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_type\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_source\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setTypeAndSourceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_types\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sources\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"}],\"name\":\"setTypeAndSourceFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// FtdcRequestFeeConfigurationsABI is the input ABI used to generate the binding from.
// Deprecated: Use FtdcRequestFeeConfigurationsMetaData.ABI instead.
var FtdcRequestFeeConfigurationsABI = FtdcRequestFeeConfigurationsMetaData.ABI

// FtdcRequestFeeConfigurations is an auto generated Go binding around an Ethereum contract.
type FtdcRequestFeeConfigurations struct {
	FtdcRequestFeeConfigurationsCaller     // Read-only binding to the contract
	FtdcRequestFeeConfigurationsTransactor // Write-only binding to the contract
	FtdcRequestFeeConfigurationsFilterer   // Log filterer for contract events
}

// FtdcRequestFeeConfigurationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type FtdcRequestFeeConfigurationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtdcRequestFeeConfigurationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FtdcRequestFeeConfigurationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtdcRequestFeeConfigurationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FtdcRequestFeeConfigurationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtdcRequestFeeConfigurationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FtdcRequestFeeConfigurationsSession struct {
	Contract     *FtdcRequestFeeConfigurations // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// FtdcRequestFeeConfigurationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FtdcRequestFeeConfigurationsCallerSession struct {
	Contract *FtdcRequestFeeConfigurationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// FtdcRequestFeeConfigurationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FtdcRequestFeeConfigurationsTransactorSession struct {
	Contract     *FtdcRequestFeeConfigurationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// FtdcRequestFeeConfigurationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type FtdcRequestFeeConfigurationsRaw struct {
	Contract *FtdcRequestFeeConfigurations // Generic contract binding to access the raw methods on
}

// FtdcRequestFeeConfigurationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FtdcRequestFeeConfigurationsCallerRaw struct {
	Contract *FtdcRequestFeeConfigurationsCaller // Generic read-only contract binding to access the raw methods on
}

// FtdcRequestFeeConfigurationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FtdcRequestFeeConfigurationsTransactorRaw struct {
	Contract *FtdcRequestFeeConfigurationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFtdcRequestFeeConfigurations creates a new instance of FtdcRequestFeeConfigurations, bound to a specific deployed contract.
func NewFtdcRequestFeeConfigurations(address common.Address, backend bind.ContractBackend) (*FtdcRequestFeeConfigurations, error) {
	contract, err := bindFtdcRequestFeeConfigurations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FtdcRequestFeeConfigurations{FtdcRequestFeeConfigurationsCaller: FtdcRequestFeeConfigurationsCaller{contract: contract}, FtdcRequestFeeConfigurationsTransactor: FtdcRequestFeeConfigurationsTransactor{contract: contract}, FtdcRequestFeeConfigurationsFilterer: FtdcRequestFeeConfigurationsFilterer{contract: contract}}, nil
}

// NewFtdcRequestFeeConfigurationsCaller creates a new read-only instance of FtdcRequestFeeConfigurations, bound to a specific deployed contract.
func NewFtdcRequestFeeConfigurationsCaller(address common.Address, caller bind.ContractCaller) (*FtdcRequestFeeConfigurationsCaller, error) {
	contract, err := bindFtdcRequestFeeConfigurations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FtdcRequestFeeConfigurationsCaller{contract: contract}, nil
}

// NewFtdcRequestFeeConfigurationsTransactor creates a new write-only instance of FtdcRequestFeeConfigurations, bound to a specific deployed contract.
func NewFtdcRequestFeeConfigurationsTransactor(address common.Address, transactor bind.ContractTransactor) (*FtdcRequestFeeConfigurationsTransactor, error) {
	contract, err := bindFtdcRequestFeeConfigurations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FtdcRequestFeeConfigurationsTransactor{contract: contract}, nil
}

// NewFtdcRequestFeeConfigurationsFilterer creates a new log filterer instance of FtdcRequestFeeConfigurations, bound to a specific deployed contract.
func NewFtdcRequestFeeConfigurationsFilterer(address common.Address, filterer bind.ContractFilterer) (*FtdcRequestFeeConfigurationsFilterer, error) {
	contract, err := bindFtdcRequestFeeConfigurations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FtdcRequestFeeConfigurationsFilterer{contract: contract}, nil
}

// bindFtdcRequestFeeConfigurations binds a generic wrapper to an already deployed contract.
func bindFtdcRequestFeeConfigurations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FtdcRequestFeeConfigurationsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FtdcRequestFeeConfigurations.Contract.FtdcRequestFeeConfigurationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.FtdcRequestFeeConfigurationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.FtdcRequestFeeConfigurationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FtdcRequestFeeConfigurations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FtdcRequestFeeConfigurations.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _FtdcRequestFeeConfigurations.Contract.UPGRADEINTERFACEVERSION(&_FtdcRequestFeeConfigurations.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _FtdcRequestFeeConfigurations.Contract.UPGRADEINTERFACEVERSION(&_FtdcRequestFeeConfigurations.CallOpts)
}

// GetTypeAndSourceFee is a free data retrieval call binding the contract method 0x5ebe55b3.
//
// Solidity: function getTypeAndSourceFee(bytes32 _type, bytes32 _source) view returns(uint256 _fee)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCaller) GetTypeAndSourceFee(opts *bind.CallOpts, _type [32]byte, _source [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _FtdcRequestFeeConfigurations.contract.Call(opts, &out, "getTypeAndSourceFee", _type, _source)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTypeAndSourceFee is a free data retrieval call binding the contract method 0x5ebe55b3.
//
// Solidity: function getTypeAndSourceFee(bytes32 _type, bytes32 _source) view returns(uint256 _fee)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) GetTypeAndSourceFee(_type [32]byte, _source [32]byte) (*big.Int, error) {
	return _FtdcRequestFeeConfigurations.Contract.GetTypeAndSourceFee(&_FtdcRequestFeeConfigurations.CallOpts, _type, _source)
}

// GetTypeAndSourceFee is a free data retrieval call binding the contract method 0x5ebe55b3.
//
// Solidity: function getTypeAndSourceFee(bytes32 _type, bytes32 _source) view returns(uint256 _fee)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCallerSession) GetTypeAndSourceFee(_type [32]byte, _source [32]byte) (*big.Int, error) {
	return _FtdcRequestFeeConfigurations.Contract.GetTypeAndSourceFee(&_FtdcRequestFeeConfigurations.CallOpts, _type, _source)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtdcRequestFeeConfigurations.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) Governance() (common.Address, error) {
	return _FtdcRequestFeeConfigurations.Contract.Governance(&_FtdcRequestFeeConfigurations.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCallerSession) Governance() (common.Address, error) {
	return _FtdcRequestFeeConfigurations.Contract.Governance(&_FtdcRequestFeeConfigurations.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtdcRequestFeeConfigurations.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) GovernanceSettings() (common.Address, error) {
	return _FtdcRequestFeeConfigurations.Contract.GovernanceSettings(&_FtdcRequestFeeConfigurations.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCallerSession) GovernanceSettings() (common.Address, error) {
	return _FtdcRequestFeeConfigurations.Contract.GovernanceSettings(&_FtdcRequestFeeConfigurations.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtdcRequestFeeConfigurations.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) Implementation() (common.Address, error) {
	return _FtdcRequestFeeConfigurations.Contract.Implementation(&_FtdcRequestFeeConfigurations.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCallerSession) Implementation() (common.Address, error) {
	return _FtdcRequestFeeConfigurations.Contract.Implementation(&_FtdcRequestFeeConfigurations.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _FtdcRequestFeeConfigurations.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) IsExecutor(_address common.Address) (bool, error) {
	return _FtdcRequestFeeConfigurations.Contract.IsExecutor(&_FtdcRequestFeeConfigurations.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _FtdcRequestFeeConfigurations.Contract.IsExecutor(&_FtdcRequestFeeConfigurations.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FtdcRequestFeeConfigurations.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) ProductionMode() (bool, error) {
	return _FtdcRequestFeeConfigurations.Contract.ProductionMode(&_FtdcRequestFeeConfigurations.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCallerSession) ProductionMode() (bool, error) {
	return _FtdcRequestFeeConfigurations.Contract.ProductionMode(&_FtdcRequestFeeConfigurations.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FtdcRequestFeeConfigurations.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) ProxiableUUID() ([32]byte, error) {
	return _FtdcRequestFeeConfigurations.Contract.ProxiableUUID(&_FtdcRequestFeeConfigurations.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _FtdcRequestFeeConfigurations.Contract.ProxiableUUID(&_FtdcRequestFeeConfigurations.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _FtdcRequestFeeConfigurations.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FtdcRequestFeeConfigurations.Contract.TimelockedCalls(&_FtdcRequestFeeConfigurations.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FtdcRequestFeeConfigurations.Contract.TimelockedCalls(&_FtdcRequestFeeConfigurations.CallOpts, selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.CancelGovernanceCall(&_FtdcRequestFeeConfigurations.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.CancelGovernanceCall(&_FtdcRequestFeeConfigurations.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.ExecuteGovernanceCall(&_FtdcRequestFeeConfigurations.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.ExecuteGovernanceCall(&_FtdcRequestFeeConfigurations.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.Initialise(&_FtdcRequestFeeConfigurations.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.Initialise(&_FtdcRequestFeeConfigurations.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.Initialize(&_FtdcRequestFeeConfigurations.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.Initialize(&_FtdcRequestFeeConfigurations.TransactOpts, _governanceSettings, _initialGovernance)
}

// RemoveTypeAndSourceFee is a paid mutator transaction binding the contract method 0xda42a778.
//
// Solidity: function removeTypeAndSourceFee(bytes32 _type, bytes32 _source) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactor) RemoveTypeAndSourceFee(opts *bind.TransactOpts, _type [32]byte, _source [32]byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.contract.Transact(opts, "removeTypeAndSourceFee", _type, _source)
}

// RemoveTypeAndSourceFee is a paid mutator transaction binding the contract method 0xda42a778.
//
// Solidity: function removeTypeAndSourceFee(bytes32 _type, bytes32 _source) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) RemoveTypeAndSourceFee(_type [32]byte, _source [32]byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.RemoveTypeAndSourceFee(&_FtdcRequestFeeConfigurations.TransactOpts, _type, _source)
}

// RemoveTypeAndSourceFee is a paid mutator transaction binding the contract method 0xda42a778.
//
// Solidity: function removeTypeAndSourceFee(bytes32 _type, bytes32 _source) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactorSession) RemoveTypeAndSourceFee(_type [32]byte, _source [32]byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.RemoveTypeAndSourceFee(&_FtdcRequestFeeConfigurations.TransactOpts, _type, _source)
}

// RemoveTypeAndSourceFees is a paid mutator transaction binding the contract method 0xe2f6bca8.
//
// Solidity: function removeTypeAndSourceFees(bytes32[] _types, bytes32[] _sources) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactor) RemoveTypeAndSourceFees(opts *bind.TransactOpts, _types [][32]byte, _sources [][32]byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.contract.Transact(opts, "removeTypeAndSourceFees", _types, _sources)
}

// RemoveTypeAndSourceFees is a paid mutator transaction binding the contract method 0xe2f6bca8.
//
// Solidity: function removeTypeAndSourceFees(bytes32[] _types, bytes32[] _sources) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) RemoveTypeAndSourceFees(_types [][32]byte, _sources [][32]byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.RemoveTypeAndSourceFees(&_FtdcRequestFeeConfigurations.TransactOpts, _types, _sources)
}

// RemoveTypeAndSourceFees is a paid mutator transaction binding the contract method 0xe2f6bca8.
//
// Solidity: function removeTypeAndSourceFees(bytes32[] _types, bytes32[] _sources) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactorSession) RemoveTypeAndSourceFees(_types [][32]byte, _sources [][32]byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.RemoveTypeAndSourceFees(&_FtdcRequestFeeConfigurations.TransactOpts, _types, _sources)
}

// SetTypeAndSourceFee is a paid mutator transaction binding the contract method 0x5b1bacb7.
//
// Solidity: function setTypeAndSourceFee(bytes32 _type, bytes32 _source, uint256 _fee) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactor) SetTypeAndSourceFee(opts *bind.TransactOpts, _type [32]byte, _source [32]byte, _fee *big.Int) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.contract.Transact(opts, "setTypeAndSourceFee", _type, _source, _fee)
}

// SetTypeAndSourceFee is a paid mutator transaction binding the contract method 0x5b1bacb7.
//
// Solidity: function setTypeAndSourceFee(bytes32 _type, bytes32 _source, uint256 _fee) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) SetTypeAndSourceFee(_type [32]byte, _source [32]byte, _fee *big.Int) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.SetTypeAndSourceFee(&_FtdcRequestFeeConfigurations.TransactOpts, _type, _source, _fee)
}

// SetTypeAndSourceFee is a paid mutator transaction binding the contract method 0x5b1bacb7.
//
// Solidity: function setTypeAndSourceFee(bytes32 _type, bytes32 _source, uint256 _fee) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactorSession) SetTypeAndSourceFee(_type [32]byte, _source [32]byte, _fee *big.Int) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.SetTypeAndSourceFee(&_FtdcRequestFeeConfigurations.TransactOpts, _type, _source, _fee)
}

// SetTypeAndSourceFees is a paid mutator transaction binding the contract method 0x3da84157.
//
// Solidity: function setTypeAndSourceFees(bytes32[] _types, bytes32[] _sources, uint256[] _fees) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactor) SetTypeAndSourceFees(opts *bind.TransactOpts, _types [][32]byte, _sources [][32]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.contract.Transact(opts, "setTypeAndSourceFees", _types, _sources, _fees)
}

// SetTypeAndSourceFees is a paid mutator transaction binding the contract method 0x3da84157.
//
// Solidity: function setTypeAndSourceFees(bytes32[] _types, bytes32[] _sources, uint256[] _fees) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) SetTypeAndSourceFees(_types [][32]byte, _sources [][32]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.SetTypeAndSourceFees(&_FtdcRequestFeeConfigurations.TransactOpts, _types, _sources, _fees)
}

// SetTypeAndSourceFees is a paid mutator transaction binding the contract method 0x3da84157.
//
// Solidity: function setTypeAndSourceFees(bytes32[] _types, bytes32[] _sources, uint256[] _fees) returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactorSession) SetTypeAndSourceFees(_types [][32]byte, _sources [][32]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.SetTypeAndSourceFees(&_FtdcRequestFeeConfigurations.TransactOpts, _types, _sources, _fees)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.SwitchToProductionMode(&_FtdcRequestFeeConfigurations.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.SwitchToProductionMode(&_FtdcRequestFeeConfigurations.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.UpgradeToAndCall(&_FtdcRequestFeeConfigurations.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _FtdcRequestFeeConfigurations.Contract.UpgradeToAndCall(&_FtdcRequestFeeConfigurations.TransactOpts, _newImplementation, _data)
}

// FtdcRequestFeeConfigurationsGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsGovernanceCallTimelockedIterator struct {
	Event *FtdcRequestFeeConfigurationsGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *FtdcRequestFeeConfigurationsGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcRequestFeeConfigurationsGovernanceCallTimelocked)
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
		it.Event = new(FtdcRequestFeeConfigurationsGovernanceCallTimelocked)
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
func (it *FtdcRequestFeeConfigurationsGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcRequestFeeConfigurationsGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcRequestFeeConfigurationsGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*FtdcRequestFeeConfigurationsGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &FtdcRequestFeeConfigurationsGovernanceCallTimelockedIterator{contract: _FtdcRequestFeeConfigurations.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *FtdcRequestFeeConfigurationsGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcRequestFeeConfigurationsGovernanceCallTimelocked)
				if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) ParseGovernanceCallTimelocked(log types.Log) (*FtdcRequestFeeConfigurationsGovernanceCallTimelocked, error) {
	event := new(FtdcRequestFeeConfigurationsGovernanceCallTimelocked)
	if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcRequestFeeConfigurationsGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsGovernanceInitialisedIterator struct {
	Event *FtdcRequestFeeConfigurationsGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *FtdcRequestFeeConfigurationsGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcRequestFeeConfigurationsGovernanceInitialised)
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
		it.Event = new(FtdcRequestFeeConfigurationsGovernanceInitialised)
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
func (it *FtdcRequestFeeConfigurationsGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcRequestFeeConfigurationsGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcRequestFeeConfigurationsGovernanceInitialised represents a GovernanceInitialised event raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*FtdcRequestFeeConfigurationsGovernanceInitialisedIterator, error) {

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &FtdcRequestFeeConfigurationsGovernanceInitialisedIterator{contract: _FtdcRequestFeeConfigurations.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *FtdcRequestFeeConfigurationsGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcRequestFeeConfigurationsGovernanceInitialised)
				if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) ParseGovernanceInitialised(log types.Log) (*FtdcRequestFeeConfigurationsGovernanceInitialised, error) {
	event := new(FtdcRequestFeeConfigurationsGovernanceInitialised)
	if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcRequestFeeConfigurationsGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsGovernedProductionModeEnteredIterator struct {
	Event *FtdcRequestFeeConfigurationsGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *FtdcRequestFeeConfigurationsGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcRequestFeeConfigurationsGovernedProductionModeEntered)
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
		it.Event = new(FtdcRequestFeeConfigurationsGovernedProductionModeEntered)
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
func (it *FtdcRequestFeeConfigurationsGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcRequestFeeConfigurationsGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcRequestFeeConfigurationsGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*FtdcRequestFeeConfigurationsGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &FtdcRequestFeeConfigurationsGovernedProductionModeEnteredIterator{contract: _FtdcRequestFeeConfigurations.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *FtdcRequestFeeConfigurationsGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcRequestFeeConfigurationsGovernedProductionModeEntered)
				if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) ParseGovernedProductionModeEntered(log types.Log) (*FtdcRequestFeeConfigurationsGovernedProductionModeEntered, error) {
	event := new(FtdcRequestFeeConfigurationsGovernedProductionModeEntered)
	if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator struct {
	Event *FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceled)
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
		it.Event = new(FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceled)
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
func (it *FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceledIterator{contract: _FtdcRequestFeeConfigurations.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceled)
				if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceled, error) {
	event := new(FtdcRequestFeeConfigurationsTimelockedGovernanceCallCanceled)
	if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator struct {
	Event *FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecuted)
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
		it.Event = new(FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecuted)
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
func (it *FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecutedIterator{contract: _FtdcRequestFeeConfigurations.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecuted)
				if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecuted, error) {
	event := new(FtdcRequestFeeConfigurationsTimelockedGovernanceCallExecuted)
	if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcRequestFeeConfigurationsTypeAndSourceFeeRemovedIterator is returned from FilterTypeAndSourceFeeRemoved and is used to iterate over the raw logs and unpacked data for TypeAndSourceFeeRemoved events raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsTypeAndSourceFeeRemovedIterator struct {
	Event *FtdcRequestFeeConfigurationsTypeAndSourceFeeRemoved // Event containing the contract specifics and raw log

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
func (it *FtdcRequestFeeConfigurationsTypeAndSourceFeeRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcRequestFeeConfigurationsTypeAndSourceFeeRemoved)
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
		it.Event = new(FtdcRequestFeeConfigurationsTypeAndSourceFeeRemoved)
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
func (it *FtdcRequestFeeConfigurationsTypeAndSourceFeeRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcRequestFeeConfigurationsTypeAndSourceFeeRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcRequestFeeConfigurationsTypeAndSourceFeeRemoved represents a TypeAndSourceFeeRemoved event raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsTypeAndSourceFeeRemoved struct {
	AttestationType [32]byte
	Source          [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTypeAndSourceFeeRemoved is a free log retrieval operation binding the contract event 0xa186aa06a988f2cd41162e05c078a9f69e58eb1a7233361c2604a1a49721d521.
//
// Solidity: event TypeAndSourceFeeRemoved(bytes32 indexed attestationType, bytes32 indexed source)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) FilterTypeAndSourceFeeRemoved(opts *bind.FilterOpts, attestationType [][32]byte, source [][32]byte) (*FtdcRequestFeeConfigurationsTypeAndSourceFeeRemovedIterator, error) {

	var attestationTypeRule []interface{}
	for _, attestationTypeItem := range attestationType {
		attestationTypeRule = append(attestationTypeRule, attestationTypeItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.FilterLogs(opts, "TypeAndSourceFeeRemoved", attestationTypeRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return &FtdcRequestFeeConfigurationsTypeAndSourceFeeRemovedIterator{contract: _FtdcRequestFeeConfigurations.contract, event: "TypeAndSourceFeeRemoved", logs: logs, sub: sub}, nil
}

// WatchTypeAndSourceFeeRemoved is a free log subscription operation binding the contract event 0xa186aa06a988f2cd41162e05c078a9f69e58eb1a7233361c2604a1a49721d521.
//
// Solidity: event TypeAndSourceFeeRemoved(bytes32 indexed attestationType, bytes32 indexed source)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) WatchTypeAndSourceFeeRemoved(opts *bind.WatchOpts, sink chan<- *FtdcRequestFeeConfigurationsTypeAndSourceFeeRemoved, attestationType [][32]byte, source [][32]byte) (event.Subscription, error) {

	var attestationTypeRule []interface{}
	for _, attestationTypeItem := range attestationType {
		attestationTypeRule = append(attestationTypeRule, attestationTypeItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.WatchLogs(opts, "TypeAndSourceFeeRemoved", attestationTypeRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcRequestFeeConfigurationsTypeAndSourceFeeRemoved)
				if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "TypeAndSourceFeeRemoved", log); err != nil {
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
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) ParseTypeAndSourceFeeRemoved(log types.Log) (*FtdcRequestFeeConfigurationsTypeAndSourceFeeRemoved, error) {
	event := new(FtdcRequestFeeConfigurationsTypeAndSourceFeeRemoved)
	if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "TypeAndSourceFeeRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcRequestFeeConfigurationsTypeAndSourceFeeSetIterator is returned from FilterTypeAndSourceFeeSet and is used to iterate over the raw logs and unpacked data for TypeAndSourceFeeSet events raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsTypeAndSourceFeeSetIterator struct {
	Event *FtdcRequestFeeConfigurationsTypeAndSourceFeeSet // Event containing the contract specifics and raw log

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
func (it *FtdcRequestFeeConfigurationsTypeAndSourceFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcRequestFeeConfigurationsTypeAndSourceFeeSet)
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
		it.Event = new(FtdcRequestFeeConfigurationsTypeAndSourceFeeSet)
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
func (it *FtdcRequestFeeConfigurationsTypeAndSourceFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcRequestFeeConfigurationsTypeAndSourceFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcRequestFeeConfigurationsTypeAndSourceFeeSet represents a TypeAndSourceFeeSet event raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsTypeAndSourceFeeSet struct {
	AttestationType [32]byte
	Source          [32]byte
	Fee             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTypeAndSourceFeeSet is a free log retrieval operation binding the contract event 0x53fbcb0688fd3904cd33601e03a5eba0c6bba217d8c9312333ea0c116ea91382.
//
// Solidity: event TypeAndSourceFeeSet(bytes32 indexed attestationType, bytes32 indexed source, uint256 fee)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) FilterTypeAndSourceFeeSet(opts *bind.FilterOpts, attestationType [][32]byte, source [][32]byte) (*FtdcRequestFeeConfigurationsTypeAndSourceFeeSetIterator, error) {

	var attestationTypeRule []interface{}
	for _, attestationTypeItem := range attestationType {
		attestationTypeRule = append(attestationTypeRule, attestationTypeItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.FilterLogs(opts, "TypeAndSourceFeeSet", attestationTypeRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return &FtdcRequestFeeConfigurationsTypeAndSourceFeeSetIterator{contract: _FtdcRequestFeeConfigurations.contract, event: "TypeAndSourceFeeSet", logs: logs, sub: sub}, nil
}

// WatchTypeAndSourceFeeSet is a free log subscription operation binding the contract event 0x53fbcb0688fd3904cd33601e03a5eba0c6bba217d8c9312333ea0c116ea91382.
//
// Solidity: event TypeAndSourceFeeSet(bytes32 indexed attestationType, bytes32 indexed source, uint256 fee)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) WatchTypeAndSourceFeeSet(opts *bind.WatchOpts, sink chan<- *FtdcRequestFeeConfigurationsTypeAndSourceFeeSet, attestationType [][32]byte, source [][32]byte) (event.Subscription, error) {

	var attestationTypeRule []interface{}
	for _, attestationTypeItem := range attestationType {
		attestationTypeRule = append(attestationTypeRule, attestationTypeItem)
	}
	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.WatchLogs(opts, "TypeAndSourceFeeSet", attestationTypeRule, sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcRequestFeeConfigurationsTypeAndSourceFeeSet)
				if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "TypeAndSourceFeeSet", log); err != nil {
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
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) ParseTypeAndSourceFeeSet(log types.Log) (*FtdcRequestFeeConfigurationsTypeAndSourceFeeSet, error) {
	event := new(FtdcRequestFeeConfigurationsTypeAndSourceFeeSet)
	if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "TypeAndSourceFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcRequestFeeConfigurationsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsUpgradedIterator struct {
	Event *FtdcRequestFeeConfigurationsUpgraded // Event containing the contract specifics and raw log

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
func (it *FtdcRequestFeeConfigurationsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcRequestFeeConfigurationsUpgraded)
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
		it.Event = new(FtdcRequestFeeConfigurationsUpgraded)
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
func (it *FtdcRequestFeeConfigurationsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcRequestFeeConfigurationsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcRequestFeeConfigurationsUpgraded represents a Upgraded event raised by the FtdcRequestFeeConfigurations contract.
type FtdcRequestFeeConfigurationsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FtdcRequestFeeConfigurationsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FtdcRequestFeeConfigurationsUpgradedIterator{contract: _FtdcRequestFeeConfigurations.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FtdcRequestFeeConfigurationsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FtdcRequestFeeConfigurations.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcRequestFeeConfigurationsUpgraded)
				if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_FtdcRequestFeeConfigurations *FtdcRequestFeeConfigurationsFilterer) ParseUpgraded(log types.Log) (*FtdcRequestFeeConfigurationsUpgraded, error) {
	event := new(FtdcRequestFeeConfigurationsUpgraded)
	if err := _FtdcRequestFeeConfigurations.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
