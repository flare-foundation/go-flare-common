// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diamondgovernance

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

// IDiamondFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// DiamondGovernanceMetaData contains all meta data concerning the DiamondGovernance contract.
var DiamondGovernanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotAddFunctionToDiamondThatAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"_selectors\",\"type\":\"bytes4[]\"}],\"name\":\"CannotAddSelectorsToZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotRemoveFunctionThatDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotRemoveImmutableFunction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotReplaceFunctionThatDoesNotExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotReplaceFunctionWithTheSameFunctionFromTheSameFacet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"_selectors\",\"type\":\"bytes4[]\"}],\"name\":\"CannotReplaceFunctionsFromFacetWithZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotReplaceImmutableFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_action\",\"type\":\"uint8\"}],\"name\":\"IncorrectFacetCutAction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_initializationContractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"InitializationFunctionReverted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"NoBytecodeAtAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facetAddress\",\"type\":\"address\"}],\"name\":\"NoSelectorsProvidedForFacetForCut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facetAddress\",\"type\":\"address\"}],\"name\":\"RemoveFacetAddressMustBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockInvalidSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamond.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamond.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamond.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamond.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamond.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamond.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DiamondGovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondGovernanceMetaData.ABI instead.
var DiamondGovernanceABI = DiamondGovernanceMetaData.ABI

// DiamondGovernance is an auto generated Go binding around an Ethereum contract.
type DiamondGovernance struct {
	DiamondGovernanceCaller     // Read-only binding to the contract
	DiamondGovernanceTransactor // Write-only binding to the contract
	DiamondGovernanceFilterer   // Log filterer for contract events
}

// DiamondGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondGovernanceSession struct {
	Contract     *DiamondGovernance // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DiamondGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondGovernanceCallerSession struct {
	Contract *DiamondGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DiamondGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondGovernanceTransactorSession struct {
	Contract     *DiamondGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DiamondGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondGovernanceRaw struct {
	Contract *DiamondGovernance // Generic contract binding to access the raw methods on
}

// DiamondGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondGovernanceCallerRaw struct {
	Contract *DiamondGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondGovernanceTransactorRaw struct {
	Contract *DiamondGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamondGovernance creates a new instance of DiamondGovernance, bound to a specific deployed contract.
func NewDiamondGovernance(address common.Address, backend bind.ContractBackend) (*DiamondGovernance, error) {
	contract, err := bindDiamondGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiamondGovernance{DiamondGovernanceCaller: DiamondGovernanceCaller{contract: contract}, DiamondGovernanceTransactor: DiamondGovernanceTransactor{contract: contract}, DiamondGovernanceFilterer: DiamondGovernanceFilterer{contract: contract}}, nil
}

// NewDiamondGovernanceCaller creates a new read-only instance of DiamondGovernance, bound to a specific deployed contract.
func NewDiamondGovernanceCaller(address common.Address, caller bind.ContractCaller) (*DiamondGovernanceCaller, error) {
	contract, err := bindDiamondGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondGovernanceCaller{contract: contract}, nil
}

// NewDiamondGovernanceTransactor creates a new write-only instance of DiamondGovernance, bound to a specific deployed contract.
func NewDiamondGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondGovernanceTransactor, error) {
	contract, err := bindDiamondGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondGovernanceTransactor{contract: contract}, nil
}

// NewDiamondGovernanceFilterer creates a new log filterer instance of DiamondGovernance, bound to a specific deployed contract.
func NewDiamondGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondGovernanceFilterer, error) {
	contract, err := bindDiamondGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondGovernanceFilterer{contract: contract}, nil
}

// bindDiamondGovernance binds a generic wrapper to an already deployed contract.
func bindDiamondGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DiamondGovernanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondGovernance *DiamondGovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondGovernance.Contract.DiamondGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondGovernance *DiamondGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondGovernance.Contract.DiamondGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondGovernance *DiamondGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondGovernance.Contract.DiamondGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondGovernance *DiamondGovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondGovernance *DiamondGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondGovernance *DiamondGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondGovernance.Contract.contract.Transact(opts, method, params...)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_DiamondGovernance *DiamondGovernanceCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DiamondGovernance.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_DiamondGovernance *DiamondGovernanceSession) Governance() (common.Address, error) {
	return _DiamondGovernance.Contract.Governance(&_DiamondGovernance.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_DiamondGovernance *DiamondGovernanceCallerSession) Governance() (common.Address, error) {
	return _DiamondGovernance.Contract.Governance(&_DiamondGovernance.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_DiamondGovernance *DiamondGovernanceCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DiamondGovernance.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_DiamondGovernance *DiamondGovernanceSession) GovernanceSettings() (common.Address, error) {
	return _DiamondGovernance.Contract.GovernanceSettings(&_DiamondGovernance.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_DiamondGovernance *DiamondGovernanceCallerSession) GovernanceSettings() (common.Address, error) {
	return _DiamondGovernance.Contract.GovernanceSettings(&_DiamondGovernance.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_DiamondGovernance *DiamondGovernanceCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _DiamondGovernance.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_DiamondGovernance *DiamondGovernanceSession) IsExecutor(_address common.Address) (bool, error) {
	return _DiamondGovernance.Contract.IsExecutor(&_DiamondGovernance.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_DiamondGovernance *DiamondGovernanceCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _DiamondGovernance.Contract.IsExecutor(&_DiamondGovernance.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_DiamondGovernance *DiamondGovernanceCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DiamondGovernance.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_DiamondGovernance *DiamondGovernanceSession) ProductionMode() (bool, error) {
	return _DiamondGovernance.Contract.ProductionMode(&_DiamondGovernance.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_DiamondGovernance *DiamondGovernanceCallerSession) ProductionMode() (bool, error) {
	return _DiamondGovernance.Contract.ProductionMode(&_DiamondGovernance.CallOpts)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_DiamondGovernance *DiamondGovernanceTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _encodedCall []byte) (*types.Transaction, error) {
	return _DiamondGovernance.contract.Transact(opts, "cancelGovernanceCall", _encodedCall)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_DiamondGovernance *DiamondGovernanceSession) CancelGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _DiamondGovernance.Contract.CancelGovernanceCall(&_DiamondGovernance.TransactOpts, _encodedCall)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_DiamondGovernance *DiamondGovernanceTransactorSession) CancelGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _DiamondGovernance.Contract.CancelGovernanceCall(&_DiamondGovernance.TransactOpts, _encodedCall)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondGovernance *DiamondGovernanceTransactor) DiamondCut(opts *bind.TransactOpts, _diamondCut []IDiamondFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondGovernance.contract.Transact(opts, "diamondCut", _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondGovernance *DiamondGovernanceSession) DiamondCut(_diamondCut []IDiamondFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondGovernance.Contract.DiamondCut(&_DiamondGovernance.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondGovernance *DiamondGovernanceTransactorSession) DiamondCut(_diamondCut []IDiamondFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondGovernance.Contract.DiamondCut(&_DiamondGovernance.TransactOpts, _diamondCut, _init, _calldata)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_DiamondGovernance *DiamondGovernanceTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _encodedCall []byte) (*types.Transaction, error) {
	return _DiamondGovernance.contract.Transact(opts, "executeGovernanceCall", _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_DiamondGovernance *DiamondGovernanceSession) ExecuteGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _DiamondGovernance.Contract.ExecuteGovernanceCall(&_DiamondGovernance.TransactOpts, _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_DiamondGovernance *DiamondGovernanceTransactorSession) ExecuteGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _DiamondGovernance.Contract.ExecuteGovernanceCall(&_DiamondGovernance.TransactOpts, _encodedCall)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_DiamondGovernance *DiamondGovernanceTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondGovernance.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_DiamondGovernance *DiamondGovernanceSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _DiamondGovernance.Contract.SwitchToProductionMode(&_DiamondGovernance.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_DiamondGovernance *DiamondGovernanceTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _DiamondGovernance.Contract.SwitchToProductionMode(&_DiamondGovernance.TransactOpts)
}

// DiamondGovernanceDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the DiamondGovernance contract.
type DiamondGovernanceDiamondCutIterator struct {
	Event *DiamondGovernanceDiamondCut // Event containing the contract specifics and raw log

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
func (it *DiamondGovernanceDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondGovernanceDiamondCut)
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
		it.Event = new(DiamondGovernanceDiamondCut)
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
func (it *DiamondGovernanceDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondGovernanceDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondGovernanceDiamondCut represents a DiamondCut event raised by the DiamondGovernance contract.
type DiamondGovernanceDiamondCut struct {
	DiamondCut []IDiamondFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondGovernance *DiamondGovernanceFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*DiamondGovernanceDiamondCutIterator, error) {

	logs, sub, err := _DiamondGovernance.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &DiamondGovernanceDiamondCutIterator{contract: _DiamondGovernance.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondGovernance *DiamondGovernanceFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *DiamondGovernanceDiamondCut) (event.Subscription, error) {

	logs, sub, err := _DiamondGovernance.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondGovernanceDiamondCut)
				if err := _DiamondGovernance.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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

// ParseDiamondCut is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondGovernance *DiamondGovernanceFilterer) ParseDiamondCut(log types.Log) (*DiamondGovernanceDiamondCut, error) {
	event := new(DiamondGovernanceDiamondCut)
	if err := _DiamondGovernance.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiamondGovernanceDiamondCut0Iterator is returned from FilterDiamondCut0 and is used to iterate over the raw logs and unpacked data for DiamondCut0 events raised by the DiamondGovernance contract.
type DiamondGovernanceDiamondCut0Iterator struct {
	Event *DiamondGovernanceDiamondCut0 // Event containing the contract specifics and raw log

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
func (it *DiamondGovernanceDiamondCut0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondGovernanceDiamondCut0)
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
		it.Event = new(DiamondGovernanceDiamondCut0)
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
func (it *DiamondGovernanceDiamondCut0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondGovernanceDiamondCut0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondGovernanceDiamondCut0 represents a DiamondCut0 event raised by the DiamondGovernance contract.
type DiamondGovernanceDiamondCut0 struct {
	DiamondCut []IDiamondFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut0 is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondGovernance *DiamondGovernanceFilterer) FilterDiamondCut0(opts *bind.FilterOpts) (*DiamondGovernanceDiamondCut0Iterator, error) {

	logs, sub, err := _DiamondGovernance.contract.FilterLogs(opts, "DiamondCut0")
	if err != nil {
		return nil, err
	}
	return &DiamondGovernanceDiamondCut0Iterator{contract: _DiamondGovernance.contract, event: "DiamondCut0", logs: logs, sub: sub}, nil
}

// WatchDiamondCut0 is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondGovernance *DiamondGovernanceFilterer) WatchDiamondCut0(opts *bind.WatchOpts, sink chan<- *DiamondGovernanceDiamondCut0) (event.Subscription, error) {

	logs, sub, err := _DiamondGovernance.contract.WatchLogs(opts, "DiamondCut0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondGovernanceDiamondCut0)
				if err := _DiamondGovernance.contract.UnpackLog(event, "DiamondCut0", log); err != nil {
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

// ParseDiamondCut0 is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondGovernance *DiamondGovernanceFilterer) ParseDiamondCut0(log types.Log) (*DiamondGovernanceDiamondCut0, error) {
	event := new(DiamondGovernanceDiamondCut0)
	if err := _DiamondGovernance.contract.UnpackLog(event, "DiamondCut0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiamondGovernanceGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the DiamondGovernance contract.
type DiamondGovernanceGovernanceCallTimelockedIterator struct {
	Event *DiamondGovernanceGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *DiamondGovernanceGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondGovernanceGovernanceCallTimelocked)
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
		it.Event = new(DiamondGovernanceGovernanceCallTimelocked)
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
func (it *DiamondGovernanceGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondGovernanceGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondGovernanceGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the DiamondGovernance contract.
type DiamondGovernanceGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_DiamondGovernance *DiamondGovernanceFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*DiamondGovernanceGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _DiamondGovernance.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &DiamondGovernanceGovernanceCallTimelockedIterator{contract: _DiamondGovernance.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_DiamondGovernance *DiamondGovernanceFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *DiamondGovernanceGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _DiamondGovernance.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondGovernanceGovernanceCallTimelocked)
				if err := _DiamondGovernance.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_DiamondGovernance *DiamondGovernanceFilterer) ParseGovernanceCallTimelocked(log types.Log) (*DiamondGovernanceGovernanceCallTimelocked, error) {
	event := new(DiamondGovernanceGovernanceCallTimelocked)
	if err := _DiamondGovernance.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiamondGovernanceGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the DiamondGovernance contract.
type DiamondGovernanceGovernanceInitialisedIterator struct {
	Event *DiamondGovernanceGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *DiamondGovernanceGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondGovernanceGovernanceInitialised)
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
		it.Event = new(DiamondGovernanceGovernanceInitialised)
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
func (it *DiamondGovernanceGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondGovernanceGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondGovernanceGovernanceInitialised represents a GovernanceInitialised event raised by the DiamondGovernance contract.
type DiamondGovernanceGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_DiamondGovernance *DiamondGovernanceFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*DiamondGovernanceGovernanceInitialisedIterator, error) {

	logs, sub, err := _DiamondGovernance.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &DiamondGovernanceGovernanceInitialisedIterator{contract: _DiamondGovernance.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_DiamondGovernance *DiamondGovernanceFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *DiamondGovernanceGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _DiamondGovernance.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondGovernanceGovernanceInitialised)
				if err := _DiamondGovernance.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_DiamondGovernance *DiamondGovernanceFilterer) ParseGovernanceInitialised(log types.Log) (*DiamondGovernanceGovernanceInitialised, error) {
	event := new(DiamondGovernanceGovernanceInitialised)
	if err := _DiamondGovernance.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiamondGovernanceGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the DiamondGovernance contract.
type DiamondGovernanceGovernedProductionModeEnteredIterator struct {
	Event *DiamondGovernanceGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *DiamondGovernanceGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondGovernanceGovernedProductionModeEntered)
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
		it.Event = new(DiamondGovernanceGovernedProductionModeEntered)
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
func (it *DiamondGovernanceGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondGovernanceGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondGovernanceGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the DiamondGovernance contract.
type DiamondGovernanceGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_DiamondGovernance *DiamondGovernanceFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*DiamondGovernanceGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _DiamondGovernance.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &DiamondGovernanceGovernedProductionModeEnteredIterator{contract: _DiamondGovernance.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_DiamondGovernance *DiamondGovernanceFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *DiamondGovernanceGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _DiamondGovernance.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondGovernanceGovernedProductionModeEntered)
				if err := _DiamondGovernance.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_DiamondGovernance *DiamondGovernanceFilterer) ParseGovernedProductionModeEntered(log types.Log) (*DiamondGovernanceGovernedProductionModeEntered, error) {
	event := new(DiamondGovernanceGovernedProductionModeEntered)
	if err := _DiamondGovernance.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiamondGovernanceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DiamondGovernance contract.
type DiamondGovernanceInitializedIterator struct {
	Event *DiamondGovernanceInitialized // Event containing the contract specifics and raw log

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
func (it *DiamondGovernanceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondGovernanceInitialized)
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
		it.Event = new(DiamondGovernanceInitialized)
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
func (it *DiamondGovernanceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondGovernanceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondGovernanceInitialized represents a Initialized event raised by the DiamondGovernance contract.
type DiamondGovernanceInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DiamondGovernance *DiamondGovernanceFilterer) FilterInitialized(opts *bind.FilterOpts) (*DiamondGovernanceInitializedIterator, error) {

	logs, sub, err := _DiamondGovernance.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DiamondGovernanceInitializedIterator{contract: _DiamondGovernance.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DiamondGovernance *DiamondGovernanceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DiamondGovernanceInitialized) (event.Subscription, error) {

	logs, sub, err := _DiamondGovernance.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondGovernanceInitialized)
				if err := _DiamondGovernance.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DiamondGovernance *DiamondGovernanceFilterer) ParseInitialized(log types.Log) (*DiamondGovernanceInitialized, error) {
	event := new(DiamondGovernanceInitialized)
	if err := _DiamondGovernance.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiamondGovernanceTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the DiamondGovernance contract.
type DiamondGovernanceTimelockedGovernanceCallCanceledIterator struct {
	Event *DiamondGovernanceTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *DiamondGovernanceTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondGovernanceTimelockedGovernanceCallCanceled)
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
		it.Event = new(DiamondGovernanceTimelockedGovernanceCallCanceled)
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
func (it *DiamondGovernanceTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondGovernanceTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondGovernanceTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the DiamondGovernance contract.
type DiamondGovernanceTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_DiamondGovernance *DiamondGovernanceFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*DiamondGovernanceTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _DiamondGovernance.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &DiamondGovernanceTimelockedGovernanceCallCanceledIterator{contract: _DiamondGovernance.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_DiamondGovernance *DiamondGovernanceFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *DiamondGovernanceTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _DiamondGovernance.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondGovernanceTimelockedGovernanceCallCanceled)
				if err := _DiamondGovernance.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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

// ParseTimelockedGovernanceCallCanceled is a log parse operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_DiamondGovernance *DiamondGovernanceFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*DiamondGovernanceTimelockedGovernanceCallCanceled, error) {
	event := new(DiamondGovernanceTimelockedGovernanceCallCanceled)
	if err := _DiamondGovernance.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiamondGovernanceTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the DiamondGovernance contract.
type DiamondGovernanceTimelockedGovernanceCallExecutedIterator struct {
	Event *DiamondGovernanceTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *DiamondGovernanceTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondGovernanceTimelockedGovernanceCallExecuted)
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
		it.Event = new(DiamondGovernanceTimelockedGovernanceCallExecuted)
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
func (it *DiamondGovernanceTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondGovernanceTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondGovernanceTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the DiamondGovernance contract.
type DiamondGovernanceTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_DiamondGovernance *DiamondGovernanceFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*DiamondGovernanceTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _DiamondGovernance.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &DiamondGovernanceTimelockedGovernanceCallExecutedIterator{contract: _DiamondGovernance.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_DiamondGovernance *DiamondGovernanceFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *DiamondGovernanceTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _DiamondGovernance.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondGovernanceTimelockedGovernanceCallExecuted)
				if err := _DiamondGovernance.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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

// ParseTimelockedGovernanceCallExecuted is a log parse operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_DiamondGovernance *DiamondGovernanceFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*DiamondGovernanceTimelockedGovernanceCallExecuted, error) {
	event := new(DiamondGovernanceTimelockedGovernanceCallExecuted)
	if err := _DiamondGovernance.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
