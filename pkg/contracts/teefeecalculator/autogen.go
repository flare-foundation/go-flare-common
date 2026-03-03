// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teefeecalculator

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

// TeeFeeCalculatorMetaData contains all meta data concerning the TeeFeeCalculator contract.
var TeeFeeCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultFee\",\"type\":\"uint256\"}],\"name\":\"DefaultFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperationFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"name\":\"calculateFeeByTeeIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"}],\"name\":\"getOperationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_defaultFee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_defaultFee\",\"type\":\"uint256\"}],\"name\":\"setDefaultFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_opTypes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_opCommands\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"}],\"name\":\"setOperationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeeFeeCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeFeeCalculatorMetaData.ABI instead.
var TeeFeeCalculatorABI = TeeFeeCalculatorMetaData.ABI

// TeeFeeCalculator is an auto generated Go binding around an Ethereum contract.
type TeeFeeCalculator struct {
	TeeFeeCalculatorCaller     // Read-only binding to the contract
	TeeFeeCalculatorTransactor // Write-only binding to the contract
	TeeFeeCalculatorFilterer   // Log filterer for contract events
}

// TeeFeeCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeFeeCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeFeeCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeFeeCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeFeeCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeFeeCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeFeeCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeFeeCalculatorSession struct {
	Contract     *TeeFeeCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeFeeCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeFeeCalculatorCallerSession struct {
	Contract *TeeFeeCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TeeFeeCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeFeeCalculatorTransactorSession struct {
	Contract     *TeeFeeCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TeeFeeCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeFeeCalculatorRaw struct {
	Contract *TeeFeeCalculator // Generic contract binding to access the raw methods on
}

// TeeFeeCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeFeeCalculatorCallerRaw struct {
	Contract *TeeFeeCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// TeeFeeCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeFeeCalculatorTransactorRaw struct {
	Contract *TeeFeeCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeFeeCalculator creates a new instance of TeeFeeCalculator, bound to a specific deployed contract.
func NewTeeFeeCalculator(address common.Address, backend bind.ContractBackend) (*TeeFeeCalculator, error) {
	contract, err := bindTeeFeeCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculator{TeeFeeCalculatorCaller: TeeFeeCalculatorCaller{contract: contract}, TeeFeeCalculatorTransactor: TeeFeeCalculatorTransactor{contract: contract}, TeeFeeCalculatorFilterer: TeeFeeCalculatorFilterer{contract: contract}}, nil
}

// NewTeeFeeCalculatorCaller creates a new read-only instance of TeeFeeCalculator, bound to a specific deployed contract.
func NewTeeFeeCalculatorCaller(address common.Address, caller bind.ContractCaller) (*TeeFeeCalculatorCaller, error) {
	contract, err := bindTeeFeeCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorCaller{contract: contract}, nil
}

// NewTeeFeeCalculatorTransactor creates a new write-only instance of TeeFeeCalculator, bound to a specific deployed contract.
func NewTeeFeeCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeFeeCalculatorTransactor, error) {
	contract, err := bindTeeFeeCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorTransactor{contract: contract}, nil
}

// NewTeeFeeCalculatorFilterer creates a new log filterer instance of TeeFeeCalculator, bound to a specific deployed contract.
func NewTeeFeeCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeFeeCalculatorFilterer, error) {
	contract, err := bindTeeFeeCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorFilterer{contract: contract}, nil
}

// bindTeeFeeCalculator binds a generic wrapper to an already deployed contract.
func bindTeeFeeCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeFeeCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeFeeCalculator *TeeFeeCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeFeeCalculator.Contract.TeeFeeCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeFeeCalculator *TeeFeeCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.TeeFeeCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeFeeCalculator *TeeFeeCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.TeeFeeCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeFeeCalculator *TeeFeeCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeFeeCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeFeeCalculator *TeeFeeCalculatorSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeFeeCalculator.Contract.UPGRADEINTERFACEVERSION(&_TeeFeeCalculator.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeFeeCalculator.Contract.UPGRADEINTERFACEVERSION(&_TeeFeeCalculator.CallOpts)
}

// CalculateFeeByTeeIds is a free data retrieval call binding the contract method 0x81174ec7.
//
// Solidity: function calculateFeeByTeeIds(bytes32 _opType, bytes32 _opCommand, address[] _teeIds) view returns(uint256 _fee)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) CalculateFeeByTeeIds(opts *bind.CallOpts, _opType [32]byte, _opCommand [32]byte, _teeIds []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "calculateFeeByTeeIds", _opType, _opCommand, _teeIds)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFeeByTeeIds is a free data retrieval call binding the contract method 0x81174ec7.
//
// Solidity: function calculateFeeByTeeIds(bytes32 _opType, bytes32 _opCommand, address[] _teeIds) view returns(uint256 _fee)
func (_TeeFeeCalculator *TeeFeeCalculatorSession) CalculateFeeByTeeIds(_opType [32]byte, _opCommand [32]byte, _teeIds []common.Address) (*big.Int, error) {
	return _TeeFeeCalculator.Contract.CalculateFeeByTeeIds(&_TeeFeeCalculator.CallOpts, _opType, _opCommand, _teeIds)
}

// CalculateFeeByTeeIds is a free data retrieval call binding the contract method 0x81174ec7.
//
// Solidity: function calculateFeeByTeeIds(bytes32 _opType, bytes32 _opCommand, address[] _teeIds) view returns(uint256 _fee)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) CalculateFeeByTeeIds(_opType [32]byte, _opCommand [32]byte, _teeIds []common.Address) (*big.Int, error) {
	return _TeeFeeCalculator.Contract.CalculateFeeByTeeIds(&_TeeFeeCalculator.CallOpts, _opType, _opCommand, _teeIds)
}

// GetDefaultFee is a free data retrieval call binding the contract method 0x21bacf28.
//
// Solidity: function getDefaultFee() view returns(uint256)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) GetDefaultFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "getDefaultFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDefaultFee is a free data retrieval call binding the contract method 0x21bacf28.
//
// Solidity: function getDefaultFee() view returns(uint256)
func (_TeeFeeCalculator *TeeFeeCalculatorSession) GetDefaultFee() (*big.Int, error) {
	return _TeeFeeCalculator.Contract.GetDefaultFee(&_TeeFeeCalculator.CallOpts)
}

// GetDefaultFee is a free data retrieval call binding the contract method 0x21bacf28.
//
// Solidity: function getDefaultFee() view returns(uint256)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) GetDefaultFee() (*big.Int, error) {
	return _TeeFeeCalculator.Contract.GetDefaultFee(&_TeeFeeCalculator.CallOpts)
}

// GetOperationFee is a free data retrieval call binding the contract method 0x3d5e0390.
//
// Solidity: function getOperationFee(bytes32 _opType, bytes32 _opCommand) view returns(uint256)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) GetOperationFee(opts *bind.CallOpts, _opType [32]byte, _opCommand [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "getOperationFee", _opType, _opCommand)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperationFee is a free data retrieval call binding the contract method 0x3d5e0390.
//
// Solidity: function getOperationFee(bytes32 _opType, bytes32 _opCommand) view returns(uint256)
func (_TeeFeeCalculator *TeeFeeCalculatorSession) GetOperationFee(_opType [32]byte, _opCommand [32]byte) (*big.Int, error) {
	return _TeeFeeCalculator.Contract.GetOperationFee(&_TeeFeeCalculator.CallOpts, _opType, _opCommand)
}

// GetOperationFee is a free data retrieval call binding the contract method 0x3d5e0390.
//
// Solidity: function getOperationFee(bytes32 _opType, bytes32 _opCommand) view returns(uint256)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) GetOperationFee(_opType [32]byte, _opCommand [32]byte) (*big.Int, error) {
	return _TeeFeeCalculator.Contract.GetOperationFee(&_TeeFeeCalculator.CallOpts, _opType, _opCommand)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeFeeCalculator *TeeFeeCalculatorSession) Governance() (common.Address, error) {
	return _TeeFeeCalculator.Contract.Governance(&_TeeFeeCalculator.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) Governance() (common.Address, error) {
	return _TeeFeeCalculator.Contract.Governance(&_TeeFeeCalculator.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeFeeCalculator *TeeFeeCalculatorSession) GovernanceSettings() (common.Address, error) {
	return _TeeFeeCalculator.Contract.GovernanceSettings(&_TeeFeeCalculator.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeeFeeCalculator.Contract.GovernanceSettings(&_TeeFeeCalculator.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeFeeCalculator *TeeFeeCalculatorSession) Implementation() (common.Address, error) {
	return _TeeFeeCalculator.Contract.Implementation(&_TeeFeeCalculator.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) Implementation() (common.Address, error) {
	return _TeeFeeCalculator.Contract.Implementation(&_TeeFeeCalculator.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeFeeCalculator *TeeFeeCalculatorSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeFeeCalculator.Contract.IsExecutor(&_TeeFeeCalculator.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeFeeCalculator.Contract.IsExecutor(&_TeeFeeCalculator.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeFeeCalculator *TeeFeeCalculatorSession) ProductionMode() (bool, error) {
	return _TeeFeeCalculator.Contract.ProductionMode(&_TeeFeeCalculator.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) ProductionMode() (bool, error) {
	return _TeeFeeCalculator.Contract.ProductionMode(&_TeeFeeCalculator.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeFeeCalculator *TeeFeeCalculatorSession) ProxiableUUID() ([32]byte, error) {
	return _TeeFeeCalculator.Contract.ProxiableUUID(&_TeeFeeCalculator.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeFeeCalculator.Contract.ProxiableUUID(&_TeeFeeCalculator.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_TeeFeeCalculator *TeeFeeCalculatorSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeFeeCalculator.Contract.TimelockedCalls(&_TeeFeeCalculator.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeFeeCalculator.Contract.TimelockedCalls(&_TeeFeeCalculator.CallOpts, selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeFeeCalculator.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.CancelGovernanceCall(&_TeeFeeCalculator.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.CancelGovernanceCall(&_TeeFeeCalculator.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeFeeCalculator.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.ExecuteGovernanceCall(&_TeeFeeCalculator.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.ExecuteGovernanceCall(&_TeeFeeCalculator.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeFeeCalculator.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.Initialise(&_TeeFeeCalculator.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.Initialise(&_TeeFeeCalculator.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, uint256 _defaultFee) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _defaultFee *big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _defaultFee)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, uint256 _defaultFee) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _defaultFee *big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.Initialize(&_TeeFeeCalculator.TransactOpts, _governanceSettings, _initialGovernance, _defaultFee)
}

// Initialize is a paid mutator transaction binding the contract method 0x1794bb3c.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, uint256 _defaultFee) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _defaultFee *big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.Initialize(&_TeeFeeCalculator.TransactOpts, _governanceSettings, _initialGovernance, _defaultFee)
}

// SetDefaultFee is a paid mutator transaction binding the contract method 0xc93a6c84.
//
// Solidity: function setDefaultFee(uint256 _defaultFee) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactor) SetDefaultFee(opts *bind.TransactOpts, _defaultFee *big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.contract.Transact(opts, "setDefaultFee", _defaultFee)
}

// SetDefaultFee is a paid mutator transaction binding the contract method 0xc93a6c84.
//
// Solidity: function setDefaultFee(uint256 _defaultFee) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorSession) SetDefaultFee(_defaultFee *big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.SetDefaultFee(&_TeeFeeCalculator.TransactOpts, _defaultFee)
}

// SetDefaultFee is a paid mutator transaction binding the contract method 0xc93a6c84.
//
// Solidity: function setDefaultFee(uint256 _defaultFee) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorSession) SetDefaultFee(_defaultFee *big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.SetDefaultFee(&_TeeFeeCalculator.TransactOpts, _defaultFee)
}

// SetOperationFees is a paid mutator transaction binding the contract method 0xaa259c3e.
//
// Solidity: function setOperationFees(bytes32[] _opTypes, bytes32[] _opCommands, uint256[] _fees) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactor) SetOperationFees(opts *bind.TransactOpts, _opTypes [][32]byte, _opCommands [][32]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.contract.Transact(opts, "setOperationFees", _opTypes, _opCommands, _fees)
}

// SetOperationFees is a paid mutator transaction binding the contract method 0xaa259c3e.
//
// Solidity: function setOperationFees(bytes32[] _opTypes, bytes32[] _opCommands, uint256[] _fees) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorSession) SetOperationFees(_opTypes [][32]byte, _opCommands [][32]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.SetOperationFees(&_TeeFeeCalculator.TransactOpts, _opTypes, _opCommands, _fees)
}

// SetOperationFees is a paid mutator transaction binding the contract method 0xaa259c3e.
//
// Solidity: function setOperationFees(bytes32[] _opTypes, bytes32[] _opCommands, uint256[] _fees) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorSession) SetOperationFees(_opTypes [][32]byte, _opCommands [][32]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.SetOperationFees(&_TeeFeeCalculator.TransactOpts, _opTypes, _opCommands, _fees)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeFeeCalculator.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeFeeCalculator *TeeFeeCalculatorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.SwitchToProductionMode(&_TeeFeeCalculator.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.SwitchToProductionMode(&_TeeFeeCalculator.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeFeeCalculator.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeFeeCalculator *TeeFeeCalculatorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.UpgradeToAndCall(&_TeeFeeCalculator.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.UpgradeToAndCall(&_TeeFeeCalculator.TransactOpts, _newImplementation, _data)
}

// TeeFeeCalculatorDefaultFeeSetIterator is returned from FilterDefaultFeeSet and is used to iterate over the raw logs and unpacked data for DefaultFeeSet events raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorDefaultFeeSetIterator struct {
	Event *TeeFeeCalculatorDefaultFeeSet // Event containing the contract specifics and raw log

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
func (it *TeeFeeCalculatorDefaultFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeFeeCalculatorDefaultFeeSet)
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
		it.Event = new(TeeFeeCalculatorDefaultFeeSet)
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
func (it *TeeFeeCalculatorDefaultFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeFeeCalculatorDefaultFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeFeeCalculatorDefaultFeeSet represents a DefaultFeeSet event raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorDefaultFeeSet struct {
	DefaultFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDefaultFeeSet is a free log retrieval operation binding the contract event 0xe3c879f1bacd84281e6f3b2c940aee391b4ea5d58d41f2f9ae7808469ac38127.
//
// Solidity: event DefaultFeeSet(uint256 defaultFee)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) FilterDefaultFeeSet(opts *bind.FilterOpts) (*TeeFeeCalculatorDefaultFeeSetIterator, error) {

	logs, sub, err := _TeeFeeCalculator.contract.FilterLogs(opts, "DefaultFeeSet")
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorDefaultFeeSetIterator{contract: _TeeFeeCalculator.contract, event: "DefaultFeeSet", logs: logs, sub: sub}, nil
}

// WatchDefaultFeeSet is a free log subscription operation binding the contract event 0xe3c879f1bacd84281e6f3b2c940aee391b4ea5d58d41f2f9ae7808469ac38127.
//
// Solidity: event DefaultFeeSet(uint256 defaultFee)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) WatchDefaultFeeSet(opts *bind.WatchOpts, sink chan<- *TeeFeeCalculatorDefaultFeeSet) (event.Subscription, error) {

	logs, sub, err := _TeeFeeCalculator.contract.WatchLogs(opts, "DefaultFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeFeeCalculatorDefaultFeeSet)
				if err := _TeeFeeCalculator.contract.UnpackLog(event, "DefaultFeeSet", log); err != nil {
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

// ParseDefaultFeeSet is a log parse operation binding the contract event 0xe3c879f1bacd84281e6f3b2c940aee391b4ea5d58d41f2f9ae7808469ac38127.
//
// Solidity: event DefaultFeeSet(uint256 defaultFee)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) ParseDefaultFeeSet(log types.Log) (*TeeFeeCalculatorDefaultFeeSet, error) {
	event := new(TeeFeeCalculatorDefaultFeeSet)
	if err := _TeeFeeCalculator.contract.UnpackLog(event, "DefaultFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeFeeCalculatorGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorGovernanceCallTimelockedIterator struct {
	Event *TeeFeeCalculatorGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeeFeeCalculatorGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeFeeCalculatorGovernanceCallTimelocked)
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
		it.Event = new(TeeFeeCalculatorGovernanceCallTimelocked)
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
func (it *TeeFeeCalculatorGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeFeeCalculatorGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeFeeCalculatorGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeFeeCalculatorGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeFeeCalculator.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorGovernanceCallTimelockedIterator{contract: _TeeFeeCalculator.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeFeeCalculatorGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeFeeCalculator.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeFeeCalculatorGovernanceCallTimelocked)
				if err := _TeeFeeCalculator.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeFeeCalculatorGovernanceCallTimelocked, error) {
	event := new(TeeFeeCalculatorGovernanceCallTimelocked)
	if err := _TeeFeeCalculator.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeFeeCalculatorGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorGovernanceInitialisedIterator struct {
	Event *TeeFeeCalculatorGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeeFeeCalculatorGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeFeeCalculatorGovernanceInitialised)
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
		it.Event = new(TeeFeeCalculatorGovernanceInitialised)
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
func (it *TeeFeeCalculatorGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeFeeCalculatorGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeFeeCalculatorGovernanceInitialised represents a GovernanceInitialised event raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeFeeCalculatorGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeFeeCalculator.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorGovernanceInitialisedIterator{contract: _TeeFeeCalculator.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeFeeCalculatorGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeFeeCalculator.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeFeeCalculatorGovernanceInitialised)
				if err := _TeeFeeCalculator.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) ParseGovernanceInitialised(log types.Log) (*TeeFeeCalculatorGovernanceInitialised, error) {
	event := new(TeeFeeCalculatorGovernanceInitialised)
	if err := _TeeFeeCalculator.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeFeeCalculatorGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorGovernedProductionModeEnteredIterator struct {
	Event *TeeFeeCalculatorGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeeFeeCalculatorGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeFeeCalculatorGovernedProductionModeEntered)
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
		it.Event = new(TeeFeeCalculatorGovernedProductionModeEntered)
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
func (it *TeeFeeCalculatorGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeFeeCalculatorGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeFeeCalculatorGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeeFeeCalculatorGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeeFeeCalculator.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorGovernedProductionModeEnteredIterator{contract: _TeeFeeCalculator.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeeFeeCalculatorGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeeFeeCalculator.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeFeeCalculatorGovernedProductionModeEntered)
				if err := _TeeFeeCalculator.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeeFeeCalculatorGovernedProductionModeEntered, error) {
	event := new(TeeFeeCalculatorGovernedProductionModeEntered)
	if err := _TeeFeeCalculator.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeFeeCalculatorOperationFeeSetIterator is returned from FilterOperationFeeSet and is used to iterate over the raw logs and unpacked data for OperationFeeSet events raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorOperationFeeSetIterator struct {
	Event *TeeFeeCalculatorOperationFeeSet // Event containing the contract specifics and raw log

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
func (it *TeeFeeCalculatorOperationFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeFeeCalculatorOperationFeeSet)
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
		it.Event = new(TeeFeeCalculatorOperationFeeSet)
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
func (it *TeeFeeCalculatorOperationFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeFeeCalculatorOperationFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeFeeCalculatorOperationFeeSet represents a OperationFeeSet event raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorOperationFeeSet struct {
	OpType    [32]byte
	OpCommand [32]byte
	Fee       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOperationFeeSet is a free log retrieval operation binding the contract event 0xba6398ad47dfb412ec486fc07a04e352309f53b2073e0d72df66009d2aa1ee65.
//
// Solidity: event OperationFeeSet(bytes32 opType, bytes32 opCommand, uint256 fee)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) FilterOperationFeeSet(opts *bind.FilterOpts) (*TeeFeeCalculatorOperationFeeSetIterator, error) {

	logs, sub, err := _TeeFeeCalculator.contract.FilterLogs(opts, "OperationFeeSet")
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorOperationFeeSetIterator{contract: _TeeFeeCalculator.contract, event: "OperationFeeSet", logs: logs, sub: sub}, nil
}

// WatchOperationFeeSet is a free log subscription operation binding the contract event 0xba6398ad47dfb412ec486fc07a04e352309f53b2073e0d72df66009d2aa1ee65.
//
// Solidity: event OperationFeeSet(bytes32 opType, bytes32 opCommand, uint256 fee)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) WatchOperationFeeSet(opts *bind.WatchOpts, sink chan<- *TeeFeeCalculatorOperationFeeSet) (event.Subscription, error) {

	logs, sub, err := _TeeFeeCalculator.contract.WatchLogs(opts, "OperationFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeFeeCalculatorOperationFeeSet)
				if err := _TeeFeeCalculator.contract.UnpackLog(event, "OperationFeeSet", log); err != nil {
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

// ParseOperationFeeSet is a log parse operation binding the contract event 0xba6398ad47dfb412ec486fc07a04e352309f53b2073e0d72df66009d2aa1ee65.
//
// Solidity: event OperationFeeSet(bytes32 opType, bytes32 opCommand, uint256 fee)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) ParseOperationFeeSet(log types.Log) (*TeeFeeCalculatorOperationFeeSet, error) {
	event := new(TeeFeeCalculatorOperationFeeSet)
	if err := _TeeFeeCalculator.contract.UnpackLog(event, "OperationFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeFeeCalculatorTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorTimelockedGovernanceCallCanceledIterator struct {
	Event *TeeFeeCalculatorTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeeFeeCalculatorTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeFeeCalculatorTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeeFeeCalculatorTimelockedGovernanceCallCanceled)
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
func (it *TeeFeeCalculatorTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeFeeCalculatorTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeFeeCalculatorTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeeFeeCalculatorTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeeFeeCalculator.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorTimelockedGovernanceCallCanceledIterator{contract: _TeeFeeCalculator.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeeFeeCalculatorTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeeFeeCalculator.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeFeeCalculatorTimelockedGovernanceCallCanceled)
				if err := _TeeFeeCalculator.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeeFeeCalculatorTimelockedGovernanceCallCanceled, error) {
	event := new(TeeFeeCalculatorTimelockedGovernanceCallCanceled)
	if err := _TeeFeeCalculator.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeFeeCalculatorTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorTimelockedGovernanceCallExecutedIterator struct {
	Event *TeeFeeCalculatorTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeeFeeCalculatorTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeFeeCalculatorTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeeFeeCalculatorTimelockedGovernanceCallExecuted)
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
func (it *TeeFeeCalculatorTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeFeeCalculatorTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeFeeCalculatorTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeeFeeCalculatorTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeeFeeCalculator.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorTimelockedGovernanceCallExecutedIterator{contract: _TeeFeeCalculator.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeeFeeCalculatorTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeeFeeCalculator.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeFeeCalculatorTimelockedGovernanceCallExecuted)
				if err := _TeeFeeCalculator.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeeFeeCalculatorTimelockedGovernanceCallExecuted, error) {
	event := new(TeeFeeCalculatorTimelockedGovernanceCallExecuted)
	if err := _TeeFeeCalculator.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeFeeCalculatorUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorUpgradedIterator struct {
	Event *TeeFeeCalculatorUpgraded // Event containing the contract specifics and raw log

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
func (it *TeeFeeCalculatorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeFeeCalculatorUpgraded)
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
		it.Event = new(TeeFeeCalculatorUpgraded)
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
func (it *TeeFeeCalculatorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeFeeCalculatorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeFeeCalculatorUpgraded represents a Upgraded event raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeeFeeCalculatorUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeFeeCalculator.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorUpgradedIterator{contract: _TeeFeeCalculator.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeeFeeCalculatorUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeFeeCalculator.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeFeeCalculatorUpgraded)
				if err := _TeeFeeCalculator.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) ParseUpgraded(log types.Log) (*TeeFeeCalculatorUpgraded, error) {
	event := new(TeeFeeCalculatorUpgraded)
	if err := _TeeFeeCalculator.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
