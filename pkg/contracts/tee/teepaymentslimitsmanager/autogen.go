// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepaymentslimitsmanager

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

// ITeePaymentsPMWMultisigAccount is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsPMWMultisigAccount struct {
	SourceId       [32]byte
	AccountAddress string
}

// TeePaymentsLimitsManagerMetaData contains all meta data concerning the TeePaymentsLimitsManager contract.
var TeePaymentsLimitsManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DailyLimitBelowTransactionLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dailyLimit\",\"type\":\"uint256\"}],\"name\":\"PaymentLimitsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"contractIIFlareTeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getPaymentLimitsNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_transactionLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dailyLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"setPaymentLimits\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsRegistry\",\"outputs\":[{\"internalType\":\"contractITeePaymentsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeePaymentsLimitsManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeePaymentsLimitsManagerMetaData.ABI instead.
var TeePaymentsLimitsManagerABI = TeePaymentsLimitsManagerMetaData.ABI

// TeePaymentsLimitsManager is an auto generated Go binding around an Ethereum contract.
type TeePaymentsLimitsManager struct {
	TeePaymentsLimitsManagerCaller     // Read-only binding to the contract
	TeePaymentsLimitsManagerTransactor // Write-only binding to the contract
	TeePaymentsLimitsManagerFilterer   // Log filterer for contract events
}

// TeePaymentsLimitsManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeePaymentsLimitsManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsLimitsManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeePaymentsLimitsManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsLimitsManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeePaymentsLimitsManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsLimitsManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeePaymentsLimitsManagerSession struct {
	Contract     *TeePaymentsLimitsManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TeePaymentsLimitsManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeePaymentsLimitsManagerCallerSession struct {
	Contract *TeePaymentsLimitsManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// TeePaymentsLimitsManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeePaymentsLimitsManagerTransactorSession struct {
	Contract     *TeePaymentsLimitsManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// TeePaymentsLimitsManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeePaymentsLimitsManagerRaw struct {
	Contract *TeePaymentsLimitsManager // Generic contract binding to access the raw methods on
}

// TeePaymentsLimitsManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeePaymentsLimitsManagerCallerRaw struct {
	Contract *TeePaymentsLimitsManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TeePaymentsLimitsManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeePaymentsLimitsManagerTransactorRaw struct {
	Contract *TeePaymentsLimitsManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeePaymentsLimitsManager creates a new instance of TeePaymentsLimitsManager, bound to a specific deployed contract.
func NewTeePaymentsLimitsManager(address common.Address, backend bind.ContractBackend) (*TeePaymentsLimitsManager, error) {
	contract, err := bindTeePaymentsLimitsManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsLimitsManager{TeePaymentsLimitsManagerCaller: TeePaymentsLimitsManagerCaller{contract: contract}, TeePaymentsLimitsManagerTransactor: TeePaymentsLimitsManagerTransactor{contract: contract}, TeePaymentsLimitsManagerFilterer: TeePaymentsLimitsManagerFilterer{contract: contract}}, nil
}

// NewTeePaymentsLimitsManagerCaller creates a new read-only instance of TeePaymentsLimitsManager, bound to a specific deployed contract.
func NewTeePaymentsLimitsManagerCaller(address common.Address, caller bind.ContractCaller) (*TeePaymentsLimitsManagerCaller, error) {
	contract, err := bindTeePaymentsLimitsManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsLimitsManagerCaller{contract: contract}, nil
}

// NewTeePaymentsLimitsManagerTransactor creates a new write-only instance of TeePaymentsLimitsManager, bound to a specific deployed contract.
func NewTeePaymentsLimitsManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeePaymentsLimitsManagerTransactor, error) {
	contract, err := bindTeePaymentsLimitsManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsLimitsManagerTransactor{contract: contract}, nil
}

// NewTeePaymentsLimitsManagerFilterer creates a new log filterer instance of TeePaymentsLimitsManager, bound to a specific deployed contract.
func NewTeePaymentsLimitsManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeePaymentsLimitsManagerFilterer, error) {
	contract, err := bindTeePaymentsLimitsManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsLimitsManagerFilterer{contract: contract}, nil
}

// bindTeePaymentsLimitsManager binds a generic wrapper to an already deployed contract.
func bindTeePaymentsLimitsManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeePaymentsLimitsManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePaymentsLimitsManager.Contract.TeePaymentsLimitsManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.TeePaymentsLimitsManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.TeePaymentsLimitsManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePaymentsLimitsManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeePaymentsLimitsManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePaymentsLimitsManager.Contract.UPGRADEINTERFACEVERSION(&_TeePaymentsLimitsManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePaymentsLimitsManager.Contract.UPGRADEINTERFACEVERSION(&_TeePaymentsLimitsManager.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCaller) FlareTeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsLimitsManager.contract.Call(opts, &out, "flareTeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) FlareTeeManager() (common.Address, error) {
	return _TeePaymentsLimitsManager.Contract.FlareTeeManager(&_TeePaymentsLimitsManager.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCallerSession) FlareTeeManager() (common.Address, error) {
	return _TeePaymentsLimitsManager.Contract.FlareTeeManager(&_TeePaymentsLimitsManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsLimitsManager.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) GetAddressUpdater() (common.Address, error) {
	return _TeePaymentsLimitsManager.Contract.GetAddressUpdater(&_TeePaymentsLimitsManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeePaymentsLimitsManager.Contract.GetAddressUpdater(&_TeePaymentsLimitsManager.CallOpts)
}

// GetPaymentLimitsNonce is a free data retrieval call binding the contract method 0x3ee8b8dc.
//
// Solidity: function getPaymentLimitsNonce((bytes32,string) _account) view returns(uint256 _nonce)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCaller) GetPaymentLimitsNonce(opts *bind.CallOpts, _account ITeePaymentsPMWMultisigAccount) (*big.Int, error) {
	var out []interface{}
	err := _TeePaymentsLimitsManager.contract.Call(opts, &out, "getPaymentLimitsNonce", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPaymentLimitsNonce is a free data retrieval call binding the contract method 0x3ee8b8dc.
//
// Solidity: function getPaymentLimitsNonce((bytes32,string) _account) view returns(uint256 _nonce)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) GetPaymentLimitsNonce(_account ITeePaymentsPMWMultisigAccount) (*big.Int, error) {
	return _TeePaymentsLimitsManager.Contract.GetPaymentLimitsNonce(&_TeePaymentsLimitsManager.CallOpts, _account)
}

// GetPaymentLimitsNonce is a free data retrieval call binding the contract method 0x3ee8b8dc.
//
// Solidity: function getPaymentLimitsNonce((bytes32,string) _account) view returns(uint256 _nonce)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCallerSession) GetPaymentLimitsNonce(_account ITeePaymentsPMWMultisigAccount) (*big.Int, error) {
	return _TeePaymentsLimitsManager.Contract.GetPaymentLimitsNonce(&_TeePaymentsLimitsManager.CallOpts, _account)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsLimitsManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) Governance() (common.Address, error) {
	return _TeePaymentsLimitsManager.Contract.Governance(&_TeePaymentsLimitsManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCallerSession) Governance() (common.Address, error) {
	return _TeePaymentsLimitsManager.Contract.Governance(&_TeePaymentsLimitsManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsLimitsManager.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) GovernanceSettings() (common.Address, error) {
	return _TeePaymentsLimitsManager.Contract.GovernanceSettings(&_TeePaymentsLimitsManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeePaymentsLimitsManager.Contract.GovernanceSettings(&_TeePaymentsLimitsManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsLimitsManager.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) Implementation() (common.Address, error) {
	return _TeePaymentsLimitsManager.Contract.Implementation(&_TeePaymentsLimitsManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCallerSession) Implementation() (common.Address, error) {
	return _TeePaymentsLimitsManager.Contract.Implementation(&_TeePaymentsLimitsManager.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeePaymentsLimitsManager.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePaymentsLimitsManager.Contract.IsExecutor(&_TeePaymentsLimitsManager.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePaymentsLimitsManager.Contract.IsExecutor(&_TeePaymentsLimitsManager.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeePaymentsLimitsManager.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) ProductionMode() (bool, error) {
	return _TeePaymentsLimitsManager.Contract.ProductionMode(&_TeePaymentsLimitsManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCallerSession) ProductionMode() (bool, error) {
	return _TeePaymentsLimitsManager.Contract.ProductionMode(&_TeePaymentsLimitsManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePaymentsLimitsManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) ProxiableUUID() ([32]byte, error) {
	return _TeePaymentsLimitsManager.Contract.ProxiableUUID(&_TeePaymentsLimitsManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeePaymentsLimitsManager.Contract.ProxiableUUID(&_TeePaymentsLimitsManager.CallOpts)
}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCaller) TeePaymentsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsLimitsManager.contract.Call(opts, &out, "teePaymentsRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) TeePaymentsRegistry() (common.Address, error) {
	return _TeePaymentsLimitsManager.Contract.TeePaymentsRegistry(&_TeePaymentsLimitsManager.CallOpts)
}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCallerSession) TeePaymentsRegistry() (common.Address, error) {
	return _TeePaymentsLimitsManager.Contract.TeePaymentsRegistry(&_TeePaymentsLimitsManager.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeePaymentsLimitsManager.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeePaymentsLimitsManager.Contract.TimelockedCalls(&_TeePaymentsLimitsManager.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeePaymentsLimitsManager.Contract.TimelockedCalls(&_TeePaymentsLimitsManager.CallOpts, selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.CancelGovernanceCall(&_TeePaymentsLimitsManager.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.CancelGovernanceCall(&_TeePaymentsLimitsManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.ExecuteGovernanceCall(&_TeePaymentsLimitsManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.ExecuteGovernanceCall(&_TeePaymentsLimitsManager.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.Initialise(&_TeePaymentsLimitsManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.Initialise(&_TeePaymentsLimitsManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.Initialize(&_TeePaymentsLimitsManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.Initialize(&_TeePaymentsLimitsManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// SetPaymentLimits is a paid mutator transaction binding the contract method 0x635a6323.
//
// Solidity: function setPaymentLimits((bytes32,string) _account, uint256 _transactionLimit, uint256 _dailyLimit, address _claimBackAddress) payable returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactor) SetPaymentLimits(opts *bind.TransactOpts, _account ITeePaymentsPMWMultisigAccount, _transactionLimit *big.Int, _dailyLimit *big.Int, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.contract.Transact(opts, "setPaymentLimits", _account, _transactionLimit, _dailyLimit, _claimBackAddress)
}

// SetPaymentLimits is a paid mutator transaction binding the contract method 0x635a6323.
//
// Solidity: function setPaymentLimits((bytes32,string) _account, uint256 _transactionLimit, uint256 _dailyLimit, address _claimBackAddress) payable returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) SetPaymentLimits(_account ITeePaymentsPMWMultisigAccount, _transactionLimit *big.Int, _dailyLimit *big.Int, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.SetPaymentLimits(&_TeePaymentsLimitsManager.TransactOpts, _account, _transactionLimit, _dailyLimit, _claimBackAddress)
}

// SetPaymentLimits is a paid mutator transaction binding the contract method 0x635a6323.
//
// Solidity: function setPaymentLimits((bytes32,string) _account, uint256 _transactionLimit, uint256 _dailyLimit, address _claimBackAddress) payable returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactorSession) SetPaymentLimits(_account ITeePaymentsPMWMultisigAccount, _transactionLimit *big.Int, _dailyLimit *big.Int, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.SetPaymentLimits(&_TeePaymentsLimitsManager.TransactOpts, _account, _transactionLimit, _dailyLimit, _claimBackAddress)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.SwitchToProductionMode(&_TeePaymentsLimitsManager.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.SwitchToProductionMode(&_TeePaymentsLimitsManager.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.UpdateContractAddresses(&_TeePaymentsLimitsManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.UpdateContractAddresses(&_TeePaymentsLimitsManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.UpgradeToAndCall(&_TeePaymentsLimitsManager.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsLimitsManager.Contract.UpgradeToAndCall(&_TeePaymentsLimitsManager.TransactOpts, _newImplementation, _data)
}

// TeePaymentsLimitsManagerGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerGovernanceCallTimelockedIterator struct {
	Event *TeePaymentsLimitsManagerGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeePaymentsLimitsManagerGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsLimitsManagerGovernanceCallTimelocked)
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
		it.Event = new(TeePaymentsLimitsManagerGovernanceCallTimelocked)
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
func (it *TeePaymentsLimitsManagerGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsLimitsManagerGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsLimitsManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeePaymentsLimitsManagerGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeePaymentsLimitsManager.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsLimitsManagerGovernanceCallTimelockedIterator{contract: _TeePaymentsLimitsManager.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeePaymentsLimitsManagerGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsLimitsManager.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsLimitsManagerGovernanceCallTimelocked)
				if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeePaymentsLimitsManagerGovernanceCallTimelocked, error) {
	event := new(TeePaymentsLimitsManagerGovernanceCallTimelocked)
	if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsLimitsManagerGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerGovernanceInitialisedIterator struct {
	Event *TeePaymentsLimitsManagerGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeePaymentsLimitsManagerGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsLimitsManagerGovernanceInitialised)
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
		it.Event = new(TeePaymentsLimitsManagerGovernanceInitialised)
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
func (it *TeePaymentsLimitsManagerGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsLimitsManagerGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsLimitsManagerGovernanceInitialised represents a GovernanceInitialised event raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeePaymentsLimitsManagerGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeePaymentsLimitsManager.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsLimitsManagerGovernanceInitialisedIterator{contract: _TeePaymentsLimitsManager.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeePaymentsLimitsManagerGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsLimitsManager.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsLimitsManagerGovernanceInitialised)
				if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) ParseGovernanceInitialised(log types.Log) (*TeePaymentsLimitsManagerGovernanceInitialised, error) {
	event := new(TeePaymentsLimitsManagerGovernanceInitialised)
	if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsLimitsManagerGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerGovernedProductionModeEnteredIterator struct {
	Event *TeePaymentsLimitsManagerGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeePaymentsLimitsManagerGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsLimitsManagerGovernedProductionModeEntered)
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
		it.Event = new(TeePaymentsLimitsManagerGovernedProductionModeEntered)
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
func (it *TeePaymentsLimitsManagerGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsLimitsManagerGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsLimitsManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeePaymentsLimitsManagerGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeePaymentsLimitsManager.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsLimitsManagerGovernedProductionModeEnteredIterator{contract: _TeePaymentsLimitsManager.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeePaymentsLimitsManagerGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsLimitsManager.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsLimitsManagerGovernedProductionModeEntered)
				if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeePaymentsLimitsManagerGovernedProductionModeEntered, error) {
	event := new(TeePaymentsLimitsManagerGovernedProductionModeEntered)
	if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsLimitsManagerPaymentLimitsSetIterator is returned from FilterPaymentLimitsSet and is used to iterate over the raw logs and unpacked data for PaymentLimitsSet events raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerPaymentLimitsSetIterator struct {
	Event *TeePaymentsLimitsManagerPaymentLimitsSet // Event containing the contract specifics and raw log

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
func (it *TeePaymentsLimitsManagerPaymentLimitsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsLimitsManagerPaymentLimitsSet)
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
		it.Event = new(TeePaymentsLimitsManagerPaymentLimitsSet)
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
func (it *TeePaymentsLimitsManagerPaymentLimitsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsLimitsManagerPaymentLimitsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsLimitsManagerPaymentLimitsSet represents a PaymentLimitsSet event raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerPaymentLimitsSet struct {
	WalletId         [32]byte
	SourceId         [32]byte
	AccountAddress   string
	TransactionLimit *big.Int
	DailyLimit       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPaymentLimitsSet is a free log retrieval operation binding the contract event 0x535dac6e2ff89b763f52fb267334f91ca4e04fd947b660cbe7200a6de51388df.
//
// Solidity: event PaymentLimitsSet(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint256 transactionLimit, uint256 dailyLimit)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) FilterPaymentLimitsSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeePaymentsLimitsManagerPaymentLimitsSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePaymentsLimitsManager.contract.FilterLogs(opts, "PaymentLimitsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsLimitsManagerPaymentLimitsSetIterator{contract: _TeePaymentsLimitsManager.contract, event: "PaymentLimitsSet", logs: logs, sub: sub}, nil
}

// WatchPaymentLimitsSet is a free log subscription operation binding the contract event 0x535dac6e2ff89b763f52fb267334f91ca4e04fd947b660cbe7200a6de51388df.
//
// Solidity: event PaymentLimitsSet(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint256 transactionLimit, uint256 dailyLimit)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) WatchPaymentLimitsSet(opts *bind.WatchOpts, sink chan<- *TeePaymentsLimitsManagerPaymentLimitsSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePaymentsLimitsManager.contract.WatchLogs(opts, "PaymentLimitsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsLimitsManagerPaymentLimitsSet)
				if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "PaymentLimitsSet", log); err != nil {
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

// ParsePaymentLimitsSet is a log parse operation binding the contract event 0x535dac6e2ff89b763f52fb267334f91ca4e04fd947b660cbe7200a6de51388df.
//
// Solidity: event PaymentLimitsSet(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint256 transactionLimit, uint256 dailyLimit)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) ParsePaymentLimitsSet(log types.Log) (*TeePaymentsLimitsManagerPaymentLimitsSet, error) {
	event := new(TeePaymentsLimitsManagerPaymentLimitsSet)
	if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "PaymentLimitsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsLimitsManagerTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerTimelockedGovernanceCallCanceledIterator struct {
	Event *TeePaymentsLimitsManagerTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeePaymentsLimitsManagerTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsLimitsManagerTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeePaymentsLimitsManagerTimelockedGovernanceCallCanceled)
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
func (it *TeePaymentsLimitsManagerTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsLimitsManagerTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsLimitsManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeePaymentsLimitsManagerTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeePaymentsLimitsManager.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsLimitsManagerTimelockedGovernanceCallCanceledIterator{contract: _TeePaymentsLimitsManager.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeePaymentsLimitsManagerTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsLimitsManager.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsLimitsManagerTimelockedGovernanceCallCanceled)
				if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeePaymentsLimitsManagerTimelockedGovernanceCallCanceled, error) {
	event := new(TeePaymentsLimitsManagerTimelockedGovernanceCallCanceled)
	if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsLimitsManagerTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerTimelockedGovernanceCallExecutedIterator struct {
	Event *TeePaymentsLimitsManagerTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeePaymentsLimitsManagerTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsLimitsManagerTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeePaymentsLimitsManagerTimelockedGovernanceCallExecuted)
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
func (it *TeePaymentsLimitsManagerTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsLimitsManagerTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsLimitsManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeePaymentsLimitsManagerTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeePaymentsLimitsManager.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsLimitsManagerTimelockedGovernanceCallExecutedIterator{contract: _TeePaymentsLimitsManager.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeePaymentsLimitsManagerTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsLimitsManager.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsLimitsManagerTimelockedGovernanceCallExecuted)
				if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeePaymentsLimitsManagerTimelockedGovernanceCallExecuted, error) {
	event := new(TeePaymentsLimitsManagerTimelockedGovernanceCallExecuted)
	if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsLimitsManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerUpgradedIterator struct {
	Event *TeePaymentsLimitsManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *TeePaymentsLimitsManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsLimitsManagerUpgraded)
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
		it.Event = new(TeePaymentsLimitsManagerUpgraded)
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
func (it *TeePaymentsLimitsManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsLimitsManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsLimitsManagerUpgraded represents a Upgraded event raised by the TeePaymentsLimitsManager contract.
type TeePaymentsLimitsManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeePaymentsLimitsManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePaymentsLimitsManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsLimitsManagerUpgradedIterator{contract: _TeePaymentsLimitsManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeePaymentsLimitsManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePaymentsLimitsManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsLimitsManagerUpgraded)
				if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeePaymentsLimitsManager *TeePaymentsLimitsManagerFilterer) ParseUpgraded(log types.Log) (*TeePaymentsLimitsManagerUpgraded, error) {
	event := new(TeePaymentsLimitsManagerUpgraded)
	if err := _TeePaymentsLimitsManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
