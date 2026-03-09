// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teevrf

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

// TeeVRFMetaData contains all meta data concerning the TeeVRF contract.
var TeeVRFMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeesForKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonceEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuthorizationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizationAddress\",\"type\":\"address\"}],\"name\":\"VrfAuthorizationAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"name\":\"VrfRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VRF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WALLET_OP_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getVrfAuthorizationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_nonce\",\"type\":\"bytes\"}],\"name\":\"requestVrf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"name\":\"setVrfAuthorizationAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeExtensionRegistry\",\"outputs\":[{\"internalType\":\"contractIITeeExtensionRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeMachineRegistry\",\"outputs\":[{\"internalType\":\"contractITeeMachineRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletKeyManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletKeyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletProjectManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletProjectManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeeVRFABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeVRFMetaData.ABI instead.
var TeeVRFABI = TeeVRFMetaData.ABI

// TeeVRF is an auto generated Go binding around an Ethereum contract.
type TeeVRF struct {
	TeeVRFCaller     // Read-only binding to the contract
	TeeVRFTransactor // Write-only binding to the contract
	TeeVRFFilterer   // Log filterer for contract events
}

// TeeVRFCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeVRFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVRFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeVRFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVRFFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeVRFFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVRFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeVRFSession struct {
	Contract     *TeeVRF           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeVRFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeVRFCallerSession struct {
	Contract *TeeVRFCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TeeVRFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeVRFTransactorSession struct {
	Contract     *TeeVRFTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeVRFRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeVRFRaw struct {
	Contract *TeeVRF // Generic contract binding to access the raw methods on
}

// TeeVRFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeVRFCallerRaw struct {
	Contract *TeeVRFCaller // Generic read-only contract binding to access the raw methods on
}

// TeeVRFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeVRFTransactorRaw struct {
	Contract *TeeVRFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeVRF creates a new instance of TeeVRF, bound to a specific deployed contract.
func NewTeeVRF(address common.Address, backend bind.ContractBackend) (*TeeVRF, error) {
	contract, err := bindTeeVRF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeVRF{TeeVRFCaller: TeeVRFCaller{contract: contract}, TeeVRFTransactor: TeeVRFTransactor{contract: contract}, TeeVRFFilterer: TeeVRFFilterer{contract: contract}}, nil
}

// NewTeeVRFCaller creates a new read-only instance of TeeVRF, bound to a specific deployed contract.
func NewTeeVRFCaller(address common.Address, caller bind.ContractCaller) (*TeeVRFCaller, error) {
	contract, err := bindTeeVRF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeVRFCaller{contract: contract}, nil
}

// NewTeeVRFTransactor creates a new write-only instance of TeeVRF, bound to a specific deployed contract.
func NewTeeVRFTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeVRFTransactor, error) {
	contract, err := bindTeeVRF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeVRFTransactor{contract: contract}, nil
}

// NewTeeVRFFilterer creates a new log filterer instance of TeeVRF, bound to a specific deployed contract.
func NewTeeVRFFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeVRFFilterer, error) {
	contract, err := bindTeeVRF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeVRFFilterer{contract: contract}, nil
}

// bindTeeVRF binds a generic wrapper to an already deployed contract.
func bindTeeVRF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeVRFMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeVRF *TeeVRFRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeVRF.Contract.TeeVRFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeVRF *TeeVRFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVRF.Contract.TeeVRFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeVRF *TeeVRFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeVRF.Contract.TeeVRFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeVRF *TeeVRFCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeVRF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeVRF *TeeVRFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVRF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeVRF *TeeVRFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeVRF.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeVRF *TeeVRFCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeVRF *TeeVRFSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeVRF.Contract.UPGRADEINTERFACEVERSION(&_TeeVRF.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeVRF *TeeVRFCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeVRF.Contract.UPGRADEINTERFACEVERSION(&_TeeVRF.CallOpts)
}

// VRF is a free data retrieval call binding the contract method 0xa126285b.
//
// Solidity: function VRF() view returns(bytes32)
func (_TeeVRF *TeeVRFCaller) VRF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "VRF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VRF is a free data retrieval call binding the contract method 0xa126285b.
//
// Solidity: function VRF() view returns(bytes32)
func (_TeeVRF *TeeVRFSession) VRF() ([32]byte, error) {
	return _TeeVRF.Contract.VRF(&_TeeVRF.CallOpts)
}

// VRF is a free data retrieval call binding the contract method 0xa126285b.
//
// Solidity: function VRF() view returns(bytes32)
func (_TeeVRF *TeeVRFCallerSession) VRF() ([32]byte, error) {
	return _TeeVRF.Contract.VRF(&_TeeVRF.CallOpts)
}

// WALLETOPTYPE is a free data retrieval call binding the contract method 0xc3500300.
//
// Solidity: function WALLET_OP_TYPE() view returns(bytes32)
func (_TeeVRF *TeeVRFCaller) WALLETOPTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "WALLET_OP_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WALLETOPTYPE is a free data retrieval call binding the contract method 0xc3500300.
//
// Solidity: function WALLET_OP_TYPE() view returns(bytes32)
func (_TeeVRF *TeeVRFSession) WALLETOPTYPE() ([32]byte, error) {
	return _TeeVRF.Contract.WALLETOPTYPE(&_TeeVRF.CallOpts)
}

// WALLETOPTYPE is a free data retrieval call binding the contract method 0xc3500300.
//
// Solidity: function WALLET_OP_TYPE() view returns(bytes32)
func (_TeeVRF *TeeVRFCallerSession) WALLETOPTYPE() ([32]byte, error) {
	return _TeeVRF.Contract.WALLETOPTYPE(&_TeeVRF.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeVRF *TeeVRFCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeVRF *TeeVRFSession) GetAddressUpdater() (common.Address, error) {
	return _TeeVRF.Contract.GetAddressUpdater(&_TeeVRF.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeVRF *TeeVRFCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeVRF.Contract.GetAddressUpdater(&_TeeVRF.CallOpts)
}

// GetVrfAuthorizationAddress is a free data retrieval call binding the contract method 0x918945a0.
//
// Solidity: function getVrfAuthorizationAddress(bytes32 _walletId) view returns(address)
func (_TeeVRF *TeeVRFCaller) GetVrfAuthorizationAddress(opts *bind.CallOpts, _walletId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "getVrfAuthorizationAddress", _walletId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVrfAuthorizationAddress is a free data retrieval call binding the contract method 0x918945a0.
//
// Solidity: function getVrfAuthorizationAddress(bytes32 _walletId) view returns(address)
func (_TeeVRF *TeeVRFSession) GetVrfAuthorizationAddress(_walletId [32]byte) (common.Address, error) {
	return _TeeVRF.Contract.GetVrfAuthorizationAddress(&_TeeVRF.CallOpts, _walletId)
}

// GetVrfAuthorizationAddress is a free data retrieval call binding the contract method 0x918945a0.
//
// Solidity: function getVrfAuthorizationAddress(bytes32 _walletId) view returns(address)
func (_TeeVRF *TeeVRFCallerSession) GetVrfAuthorizationAddress(_walletId [32]byte) (common.Address, error) {
	return _TeeVRF.Contract.GetVrfAuthorizationAddress(&_TeeVRF.CallOpts, _walletId)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeVRF *TeeVRFCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeVRF *TeeVRFSession) Governance() (common.Address, error) {
	return _TeeVRF.Contract.Governance(&_TeeVRF.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeVRF *TeeVRFCallerSession) Governance() (common.Address, error) {
	return _TeeVRF.Contract.Governance(&_TeeVRF.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeVRF *TeeVRFCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeVRF *TeeVRFSession) GovernanceSettings() (common.Address, error) {
	return _TeeVRF.Contract.GovernanceSettings(&_TeeVRF.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeVRF *TeeVRFCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeeVRF.Contract.GovernanceSettings(&_TeeVRF.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeVRF *TeeVRFCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeVRF *TeeVRFSession) Implementation() (common.Address, error) {
	return _TeeVRF.Contract.Implementation(&_TeeVRF.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeVRF *TeeVRFCallerSession) Implementation() (common.Address, error) {
	return _TeeVRF.Contract.Implementation(&_TeeVRF.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeVRF *TeeVRFCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeVRF *TeeVRFSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeVRF.Contract.IsExecutor(&_TeeVRF.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeVRF *TeeVRFCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeVRF.Contract.IsExecutor(&_TeeVRF.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeVRF *TeeVRFCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeVRF *TeeVRFSession) ProductionMode() (bool, error) {
	return _TeeVRF.Contract.ProductionMode(&_TeeVRF.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeVRF *TeeVRFCallerSession) ProductionMode() (bool, error) {
	return _TeeVRF.Contract.ProductionMode(&_TeeVRF.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeVRF *TeeVRFCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeVRF *TeeVRFSession) ProxiableUUID() ([32]byte, error) {
	return _TeeVRF.Contract.ProxiableUUID(&_TeeVRF.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeVRF *TeeVRFCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeVRF.Contract.ProxiableUUID(&_TeeVRF.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeVRF *TeeVRFCaller) TeeExtensionRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "teeExtensionRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeVRF *TeeVRFSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeVRF.Contract.TeeExtensionRegistry(&_TeeVRF.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeVRF *TeeVRFCallerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeVRF.Contract.TeeExtensionRegistry(&_TeeVRF.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeVRF *TeeVRFCaller) TeeMachineRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "teeMachineRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeVRF *TeeVRFSession) TeeMachineRegistry() (common.Address, error) {
	return _TeeVRF.Contract.TeeMachineRegistry(&_TeeVRF.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeVRF *TeeVRFCallerSession) TeeMachineRegistry() (common.Address, error) {
	return _TeeVRF.Contract.TeeMachineRegistry(&_TeeVRF.CallOpts)
}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeVRF *TeeVRFCaller) TeeWalletKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "teeWalletKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeVRF *TeeVRFSession) TeeWalletKeyManager() (common.Address, error) {
	return _TeeVRF.Contract.TeeWalletKeyManager(&_TeeVRF.CallOpts)
}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeVRF *TeeVRFCallerSession) TeeWalletKeyManager() (common.Address, error) {
	return _TeeVRF.Contract.TeeWalletKeyManager(&_TeeVRF.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeVRF *TeeVRFCaller) TeeWalletManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "teeWalletManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeVRF *TeeVRFSession) TeeWalletManager() (common.Address, error) {
	return _TeeVRF.Contract.TeeWalletManager(&_TeeVRF.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeVRF *TeeVRFCallerSession) TeeWalletManager() (common.Address, error) {
	return _TeeVRF.Contract.TeeWalletManager(&_TeeVRF.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeVRF *TeeVRFCaller) TeeWalletProjectManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "teeWalletProjectManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeVRF *TeeVRFSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeeVRF.Contract.TeeWalletProjectManager(&_TeeVRF.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeVRF *TeeVRFCallerSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeeVRF.Contract.TeeWalletProjectManager(&_TeeVRF.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeVRF *TeeVRFCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_TeeVRF *TeeVRFSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeVRF.Contract.TimelockedCalls(&_TeeVRF.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeVRF *TeeVRFCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeVRF.Contract.TimelockedCalls(&_TeeVRF.CallOpts, selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeVRF *TeeVRFTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeVRF.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeVRF *TeeVRFSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeVRF.Contract.CancelGovernanceCall(&_TeeVRF.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeVRF *TeeVRFTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeVRF.Contract.CancelGovernanceCall(&_TeeVRF.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeVRF *TeeVRFTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeVRF.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeVRF *TeeVRFSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeVRF.Contract.ExecuteGovernanceCall(&_TeeVRF.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeVRF *TeeVRFTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeVRF.Contract.ExecuteGovernanceCall(&_TeeVRF.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeVRF *TeeVRFTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeVRF.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeVRF *TeeVRFSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeVRF.Contract.Initialise(&_TeeVRF.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeVRF *TeeVRFTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeVRF.Contract.Initialise(&_TeeVRF.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeVRF *TeeVRFTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeVRF.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeVRF *TeeVRFSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeVRF.Contract.Initialize(&_TeeVRF.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeVRF *TeeVRFTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeVRF.Contract.Initialize(&_TeeVRF.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// RequestVrf is a paid mutator transaction binding the contract method 0x0198d2e2.
//
// Solidity: function requestVrf(bytes32 _walletId, uint64 _keyId, bytes _nonce) payable returns(bytes32 _instructionId)
func (_TeeVRF *TeeVRFTransactor) RequestVrf(opts *bind.TransactOpts, _walletId [32]byte, _keyId uint64, _nonce []byte) (*types.Transaction, error) {
	return _TeeVRF.contract.Transact(opts, "requestVrf", _walletId, _keyId, _nonce)
}

// RequestVrf is a paid mutator transaction binding the contract method 0x0198d2e2.
//
// Solidity: function requestVrf(bytes32 _walletId, uint64 _keyId, bytes _nonce) payable returns(bytes32 _instructionId)
func (_TeeVRF *TeeVRFSession) RequestVrf(_walletId [32]byte, _keyId uint64, _nonce []byte) (*types.Transaction, error) {
	return _TeeVRF.Contract.RequestVrf(&_TeeVRF.TransactOpts, _walletId, _keyId, _nonce)
}

// RequestVrf is a paid mutator transaction binding the contract method 0x0198d2e2.
//
// Solidity: function requestVrf(bytes32 _walletId, uint64 _keyId, bytes _nonce) payable returns(bytes32 _instructionId)
func (_TeeVRF *TeeVRFTransactorSession) RequestVrf(_walletId [32]byte, _keyId uint64, _nonce []byte) (*types.Transaction, error) {
	return _TeeVRF.Contract.RequestVrf(&_TeeVRF.TransactOpts, _walletId, _keyId, _nonce)
}

// SetVrfAuthorizationAddress is a paid mutator transaction binding the contract method 0x76b87bcf.
//
// Solidity: function setVrfAuthorizationAddress(bytes32 _walletId, address _authorizationAddress) returns()
func (_TeeVRF *TeeVRFTransactor) SetVrfAuthorizationAddress(opts *bind.TransactOpts, _walletId [32]byte, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeeVRF.contract.Transact(opts, "setVrfAuthorizationAddress", _walletId, _authorizationAddress)
}

// SetVrfAuthorizationAddress is a paid mutator transaction binding the contract method 0x76b87bcf.
//
// Solidity: function setVrfAuthorizationAddress(bytes32 _walletId, address _authorizationAddress) returns()
func (_TeeVRF *TeeVRFSession) SetVrfAuthorizationAddress(_walletId [32]byte, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeeVRF.Contract.SetVrfAuthorizationAddress(&_TeeVRF.TransactOpts, _walletId, _authorizationAddress)
}

// SetVrfAuthorizationAddress is a paid mutator transaction binding the contract method 0x76b87bcf.
//
// Solidity: function setVrfAuthorizationAddress(bytes32 _walletId, address _authorizationAddress) returns()
func (_TeeVRF *TeeVRFTransactorSession) SetVrfAuthorizationAddress(_walletId [32]byte, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeeVRF.Contract.SetVrfAuthorizationAddress(&_TeeVRF.TransactOpts, _walletId, _authorizationAddress)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeVRF *TeeVRFTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVRF.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeVRF *TeeVRFSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeVRF.Contract.SwitchToProductionMode(&_TeeVRF.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeVRF *TeeVRFTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeVRF.Contract.SwitchToProductionMode(&_TeeVRF.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeVRF *TeeVRFTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeVRF.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeVRF *TeeVRFSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeVRF.Contract.UpdateContractAddresses(&_TeeVRF.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeVRF *TeeVRFTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeVRF.Contract.UpdateContractAddresses(&_TeeVRF.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeVRF *TeeVRFTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeVRF.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeVRF *TeeVRFSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeVRF.Contract.UpgradeToAndCall(&_TeeVRF.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeVRF *TeeVRFTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeVRF.Contract.UpgradeToAndCall(&_TeeVRF.TransactOpts, _newImplementation, _data)
}

// TeeVRFGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeVRF contract.
type TeeVRFGovernanceCallTimelockedIterator struct {
	Event *TeeVRFGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeeVRFGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVRFGovernanceCallTimelocked)
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
		it.Event = new(TeeVRFGovernanceCallTimelocked)
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
func (it *TeeVRFGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVRFGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVRFGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeVRF contract.
type TeeVRFGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeVRF *TeeVRFFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeVRFGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeVRF.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeVRFGovernanceCallTimelockedIterator{contract: _TeeVRF.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeVRF *TeeVRFFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeVRFGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeVRF.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVRFGovernanceCallTimelocked)
				if err := _TeeVRF.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeeVRF *TeeVRFFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeVRFGovernanceCallTimelocked, error) {
	event := new(TeeVRFGovernanceCallTimelocked)
	if err := _TeeVRF.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVRFGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeVRF contract.
type TeeVRFGovernanceInitialisedIterator struct {
	Event *TeeVRFGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeeVRFGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVRFGovernanceInitialised)
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
		it.Event = new(TeeVRFGovernanceInitialised)
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
func (it *TeeVRFGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVRFGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVRFGovernanceInitialised represents a GovernanceInitialised event raised by the TeeVRF contract.
type TeeVRFGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeVRF *TeeVRFFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeVRFGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeVRF.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeVRFGovernanceInitialisedIterator{contract: _TeeVRF.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeVRF *TeeVRFFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeVRFGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeVRF.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVRFGovernanceInitialised)
				if err := _TeeVRF.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeeVRF *TeeVRFFilterer) ParseGovernanceInitialised(log types.Log) (*TeeVRFGovernanceInitialised, error) {
	event := new(TeeVRFGovernanceInitialised)
	if err := _TeeVRF.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVRFGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeeVRF contract.
type TeeVRFGovernedProductionModeEnteredIterator struct {
	Event *TeeVRFGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeeVRFGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVRFGovernedProductionModeEntered)
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
		it.Event = new(TeeVRFGovernedProductionModeEntered)
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
func (it *TeeVRFGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVRFGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVRFGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeeVRF contract.
type TeeVRFGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeVRF *TeeVRFFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeeVRFGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeeVRF.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeeVRFGovernedProductionModeEnteredIterator{contract: _TeeVRF.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeVRF *TeeVRFFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeeVRFGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeeVRF.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVRFGovernedProductionModeEntered)
				if err := _TeeVRF.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeeVRF *TeeVRFFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeeVRFGovernedProductionModeEntered, error) {
	event := new(TeeVRFGovernedProductionModeEntered)
	if err := _TeeVRF.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVRFTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeeVRF contract.
type TeeVRFTimelockedGovernanceCallCanceledIterator struct {
	Event *TeeVRFTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeeVRFTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVRFTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeeVRFTimelockedGovernanceCallCanceled)
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
func (it *TeeVRFTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVRFTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVRFTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeeVRF contract.
type TeeVRFTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeVRF *TeeVRFFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeeVRFTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeeVRF.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeeVRFTimelockedGovernanceCallCanceledIterator{contract: _TeeVRF.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeVRF *TeeVRFFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeeVRFTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeeVRF.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVRFTimelockedGovernanceCallCanceled)
				if err := _TeeVRF.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeeVRF *TeeVRFFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeeVRFTimelockedGovernanceCallCanceled, error) {
	event := new(TeeVRFTimelockedGovernanceCallCanceled)
	if err := _TeeVRF.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVRFTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeeVRF contract.
type TeeVRFTimelockedGovernanceCallExecutedIterator struct {
	Event *TeeVRFTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeeVRFTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVRFTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeeVRFTimelockedGovernanceCallExecuted)
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
func (it *TeeVRFTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVRFTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVRFTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeeVRF contract.
type TeeVRFTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeVRF *TeeVRFFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeeVRFTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeeVRF.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeeVRFTimelockedGovernanceCallExecutedIterator{contract: _TeeVRF.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeVRF *TeeVRFFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeeVRFTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeeVRF.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVRFTimelockedGovernanceCallExecuted)
				if err := _TeeVRF.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeeVRF *TeeVRFFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeeVRFTimelockedGovernanceCallExecuted, error) {
	event := new(TeeVRFTimelockedGovernanceCallExecuted)
	if err := _TeeVRF.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVRFUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeeVRF contract.
type TeeVRFUpgradedIterator struct {
	Event *TeeVRFUpgraded // Event containing the contract specifics and raw log

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
func (it *TeeVRFUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVRFUpgraded)
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
		it.Event = new(TeeVRFUpgraded)
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
func (it *TeeVRFUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVRFUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVRFUpgraded represents a Upgraded event raised by the TeeVRF contract.
type TeeVRFUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeVRF *TeeVRFFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeeVRFUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeVRF.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeeVRFUpgradedIterator{contract: _TeeVRF.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeVRF *TeeVRFFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeeVRFUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeVRF.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVRFUpgraded)
				if err := _TeeVRF.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeeVRF *TeeVRFFilterer) ParseUpgraded(log types.Log) (*TeeVRFUpgraded, error) {
	event := new(TeeVRFUpgraded)
	if err := _TeeVRF.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVRFVrfAuthorizationAddressSetIterator is returned from FilterVrfAuthorizationAddressSet and is used to iterate over the raw logs and unpacked data for VrfAuthorizationAddressSet events raised by the TeeVRF contract.
type TeeVRFVrfAuthorizationAddressSetIterator struct {
	Event *TeeVRFVrfAuthorizationAddressSet // Event containing the contract specifics and raw log

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
func (it *TeeVRFVrfAuthorizationAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVRFVrfAuthorizationAddressSet)
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
		it.Event = new(TeeVRFVrfAuthorizationAddressSet)
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
func (it *TeeVRFVrfAuthorizationAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVRFVrfAuthorizationAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVRFVrfAuthorizationAddressSet represents a VrfAuthorizationAddressSet event raised by the TeeVRF contract.
type TeeVRFVrfAuthorizationAddressSet struct {
	WalletId             [32]byte
	AuthorizationAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterVrfAuthorizationAddressSet is a free log retrieval operation binding the contract event 0xdf8b8dac41b0ec7f0b1415d39e5f0885052aa11dbb18bb13043d64a3f1cb9c3b.
//
// Solidity: event VrfAuthorizationAddressSet(bytes32 indexed walletId, address authorizationAddress)
func (_TeeVRF *TeeVRFFilterer) FilterVrfAuthorizationAddressSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeeVRFVrfAuthorizationAddressSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeVRF.contract.FilterLogs(opts, "VrfAuthorizationAddressSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeVRFVrfAuthorizationAddressSetIterator{contract: _TeeVRF.contract, event: "VrfAuthorizationAddressSet", logs: logs, sub: sub}, nil
}

// WatchVrfAuthorizationAddressSet is a free log subscription operation binding the contract event 0xdf8b8dac41b0ec7f0b1415d39e5f0885052aa11dbb18bb13043d64a3f1cb9c3b.
//
// Solidity: event VrfAuthorizationAddressSet(bytes32 indexed walletId, address authorizationAddress)
func (_TeeVRF *TeeVRFFilterer) WatchVrfAuthorizationAddressSet(opts *bind.WatchOpts, sink chan<- *TeeVRFVrfAuthorizationAddressSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeVRF.contract.WatchLogs(opts, "VrfAuthorizationAddressSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVRFVrfAuthorizationAddressSet)
				if err := _TeeVRF.contract.UnpackLog(event, "VrfAuthorizationAddressSet", log); err != nil {
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

// ParseVrfAuthorizationAddressSet is a log parse operation binding the contract event 0xdf8b8dac41b0ec7f0b1415d39e5f0885052aa11dbb18bb13043d64a3f1cb9c3b.
//
// Solidity: event VrfAuthorizationAddressSet(bytes32 indexed walletId, address authorizationAddress)
func (_TeeVRF *TeeVRFFilterer) ParseVrfAuthorizationAddressSet(log types.Log) (*TeeVRFVrfAuthorizationAddressSet, error) {
	event := new(TeeVRFVrfAuthorizationAddressSet)
	if err := _TeeVRF.contract.UnpackLog(event, "VrfAuthorizationAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVRFVrfRequestedIterator is returned from FilterVrfRequested and is used to iterate over the raw logs and unpacked data for VrfRequested events raised by the TeeVRF contract.
type TeeVRFVrfRequestedIterator struct {
	Event *TeeVRFVrfRequested // Event containing the contract specifics and raw log

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
func (it *TeeVRFVrfRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVRFVrfRequested)
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
		it.Event = new(TeeVRFVrfRequested)
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
func (it *TeeVRFVrfRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVRFVrfRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVRFVrfRequested represents a VrfRequested event raised by the TeeVRF contract.
type TeeVRFVrfRequested struct {
	WalletId      [32]byte
	KeyId         uint64
	InstructionId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVrfRequested is a free log retrieval operation binding the contract event 0x52749cdb7f850186051c4aff58b749e22b3e30c23dbd43ac8c8f7a9c5eec5ddd.
//
// Solidity: event VrfRequested(bytes32 indexed walletId, uint64 keyId, bytes32 instructionId)
func (_TeeVRF *TeeVRFFilterer) FilterVrfRequested(opts *bind.FilterOpts, walletId [][32]byte) (*TeeVRFVrfRequestedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeVRF.contract.FilterLogs(opts, "VrfRequested", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeVRFVrfRequestedIterator{contract: _TeeVRF.contract, event: "VrfRequested", logs: logs, sub: sub}, nil
}

// WatchVrfRequested is a free log subscription operation binding the contract event 0x52749cdb7f850186051c4aff58b749e22b3e30c23dbd43ac8c8f7a9c5eec5ddd.
//
// Solidity: event VrfRequested(bytes32 indexed walletId, uint64 keyId, bytes32 instructionId)
func (_TeeVRF *TeeVRFFilterer) WatchVrfRequested(opts *bind.WatchOpts, sink chan<- *TeeVRFVrfRequested, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeVRF.contract.WatchLogs(opts, "VrfRequested", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVRFVrfRequested)
				if err := _TeeVRF.contract.UnpackLog(event, "VrfRequested", log); err != nil {
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

// ParseVrfRequested is a log parse operation binding the contract event 0x52749cdb7f850186051c4aff58b749e22b3e30c23dbd43ac8c8f7a9c5eec5ddd.
//
// Solidity: event VrfRequested(bytes32 indexed walletId, uint64 keyId, bytes32 instructionId)
func (_TeeVRF *TeeVRFFilterer) ParseVrfRequested(log types.Log) (*TeeVRFVrfRequested, error) {
	event := new(TeeVRFVrfRequested)
	if err := _TeeVRF.contract.UnpackLog(event, "VrfRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
