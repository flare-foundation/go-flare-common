// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teerewardoffersmanager

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

// TeeRewardOffersManagerMetaData contains all meta data concerning the TeeRewardOffersManager contract.
var TeeRewardOffersManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeOwnersPPMValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockInvalidSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"authorizedAmountWei\",\"type\":\"uint256\"}],\"name\":\"DailyAuthorizedInflationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedWei\",\"type\":\"uint256\"}],\"name\":\"InflationReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teeOwnersPPM\",\"type\":\"uint256\"}],\"name\":\"InflationRewardsOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyAuthorizedInflation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInflationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenPoolSupplyData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockedFundsWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalInflationAuthorizedWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalClaimedWei\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_teeOwnersPPM\",\"type\":\"uint24\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInflationAuthorizationReceivedTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInflationReceivedTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveInflation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_toAuthorizeWei\",\"type\":\"uint256\"}],\"name\":\"setDailyAuthorizedInflation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_teeOwnersPPM\",\"type\":\"uint24\"}],\"name\":\"setTeeOwnersPPM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeOwnersPPM\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationAuthorizedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationReceivedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationRewardsOfferedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_currentRewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_currentRewardEpochExpectedEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardEpochDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"triggerRewardEpochSwitchover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeeRewardOffersManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeRewardOffersManagerMetaData.ABI instead.
var TeeRewardOffersManagerABI = TeeRewardOffersManagerMetaData.ABI

// TeeRewardOffersManager is an auto generated Go binding around an Ethereum contract.
type TeeRewardOffersManager struct {
	TeeRewardOffersManagerCaller     // Read-only binding to the contract
	TeeRewardOffersManagerTransactor // Write-only binding to the contract
	TeeRewardOffersManagerFilterer   // Log filterer for contract events
}

// TeeRewardOffersManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeRewardOffersManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeRewardOffersManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeRewardOffersManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeRewardOffersManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeRewardOffersManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeRewardOffersManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeRewardOffersManagerSession struct {
	Contract     *TeeRewardOffersManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TeeRewardOffersManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeRewardOffersManagerCallerSession struct {
	Contract *TeeRewardOffersManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// TeeRewardOffersManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeRewardOffersManagerTransactorSession struct {
	Contract     *TeeRewardOffersManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// TeeRewardOffersManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeRewardOffersManagerRaw struct {
	Contract *TeeRewardOffersManager // Generic contract binding to access the raw methods on
}

// TeeRewardOffersManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeRewardOffersManagerCallerRaw struct {
	Contract *TeeRewardOffersManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TeeRewardOffersManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeRewardOffersManagerTransactorRaw struct {
	Contract *TeeRewardOffersManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeRewardOffersManager creates a new instance of TeeRewardOffersManager, bound to a specific deployed contract.
func NewTeeRewardOffersManager(address common.Address, backend bind.ContractBackend) (*TeeRewardOffersManager, error) {
	contract, err := bindTeeRewardOffersManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManager{TeeRewardOffersManagerCaller: TeeRewardOffersManagerCaller{contract: contract}, TeeRewardOffersManagerTransactor: TeeRewardOffersManagerTransactor{contract: contract}, TeeRewardOffersManagerFilterer: TeeRewardOffersManagerFilterer{contract: contract}}, nil
}

// NewTeeRewardOffersManagerCaller creates a new read-only instance of TeeRewardOffersManager, bound to a specific deployed contract.
func NewTeeRewardOffersManagerCaller(address common.Address, caller bind.ContractCaller) (*TeeRewardOffersManagerCaller, error) {
	contract, err := bindTeeRewardOffersManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManagerCaller{contract: contract}, nil
}

// NewTeeRewardOffersManagerTransactor creates a new write-only instance of TeeRewardOffersManager, bound to a specific deployed contract.
func NewTeeRewardOffersManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeRewardOffersManagerTransactor, error) {
	contract, err := bindTeeRewardOffersManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManagerTransactor{contract: contract}, nil
}

// NewTeeRewardOffersManagerFilterer creates a new log filterer instance of TeeRewardOffersManager, bound to a specific deployed contract.
func NewTeeRewardOffersManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeRewardOffersManagerFilterer, error) {
	contract, err := bindTeeRewardOffersManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManagerFilterer{contract: contract}, nil
}

// bindTeeRewardOffersManager binds a generic wrapper to an already deployed contract.
func bindTeeRewardOffersManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeRewardOffersManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeRewardOffersManager *TeeRewardOffersManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeRewardOffersManager.Contract.TeeRewardOffersManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeRewardOffersManager *TeeRewardOffersManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.TeeRewardOffersManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeRewardOffersManager *TeeRewardOffersManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.TeeRewardOffersManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeRewardOffersManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeRewardOffersManager.Contract.UPGRADEINTERFACEVERSION(&_TeeRewardOffersManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeRewardOffersManager.Contract.UPGRADEINTERFACEVERSION(&_TeeRewardOffersManager.CallOpts)
}

// DailyAuthorizedInflation is a free data retrieval call binding the contract method 0x708e34ce.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) DailyAuthorizedInflation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "dailyAuthorizedInflation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyAuthorizedInflation is a free data retrieval call binding the contract method 0x708e34ce.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) DailyAuthorizedInflation() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.DailyAuthorizedInflation(&_TeeRewardOffersManager.CallOpts)
}

// DailyAuthorizedInflation is a free data retrieval call binding the contract method 0x708e34ce.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) DailyAuthorizedInflation() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.DailyAuthorizedInflation(&_TeeRewardOffersManager.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) FlareSystemsManager() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.FlareSystemsManager(&_TeeRewardOffersManager.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) FlareSystemsManager() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.FlareSystemsManager(&_TeeRewardOffersManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.GetAddressUpdater(&_TeeRewardOffersManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.GetAddressUpdater(&_TeeRewardOffersManager.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) GetContractName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "getContractName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) GetContractName() (string, error) {
	return _TeeRewardOffersManager.Contract.GetContractName(&_TeeRewardOffersManager.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) GetContractName() (string, error) {
	return _TeeRewardOffersManager.Contract.GetContractName(&_TeeRewardOffersManager.CallOpts)
}

// GetExpectedBalance is a free data retrieval call binding the contract method 0xaf04cd3b.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) GetExpectedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "getExpectedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpectedBalance is a free data retrieval call binding the contract method 0xaf04cd3b.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) GetExpectedBalance() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.GetExpectedBalance(&_TeeRewardOffersManager.CallOpts)
}

// GetExpectedBalance is a free data retrieval call binding the contract method 0xaf04cd3b.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) GetExpectedBalance() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.GetExpectedBalance(&_TeeRewardOffersManager.CallOpts)
}

// GetInflationAddress is a free data retrieval call binding the contract method 0xed39d3f8.
//
// Solidity: function getInflationAddress() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) GetInflationAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "getInflationAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInflationAddress is a free data retrieval call binding the contract method 0xed39d3f8.
//
// Solidity: function getInflationAddress() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) GetInflationAddress() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.GetInflationAddress(&_TeeRewardOffersManager.CallOpts)
}

// GetInflationAddress is a free data retrieval call binding the contract method 0xed39d3f8.
//
// Solidity: function getInflationAddress() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) GetInflationAddress() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.GetInflationAddress(&_TeeRewardOffersManager.CallOpts)
}

// GetTokenPoolSupplyData is a free data retrieval call binding the contract method 0x2dafdbbf.
//
// Solidity: function getTokenPoolSupplyData() view returns(uint256 _lockedFundsWei, uint256 _totalInflationAuthorizedWei, uint256 _totalClaimedWei)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) GetTokenPoolSupplyData(opts *bind.CallOpts) (struct {
	LockedFundsWei              *big.Int
	TotalInflationAuthorizedWei *big.Int
	TotalClaimedWei             *big.Int
}, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "getTokenPoolSupplyData")

	outstruct := new(struct {
		LockedFundsWei              *big.Int
		TotalInflationAuthorizedWei *big.Int
		TotalClaimedWei             *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockedFundsWei = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalInflationAuthorizedWei = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalClaimedWei = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetTokenPoolSupplyData is a free data retrieval call binding the contract method 0x2dafdbbf.
//
// Solidity: function getTokenPoolSupplyData() view returns(uint256 _lockedFundsWei, uint256 _totalInflationAuthorizedWei, uint256 _totalClaimedWei)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) GetTokenPoolSupplyData() (struct {
	LockedFundsWei              *big.Int
	TotalInflationAuthorizedWei *big.Int
	TotalClaimedWei             *big.Int
}, error) {
	return _TeeRewardOffersManager.Contract.GetTokenPoolSupplyData(&_TeeRewardOffersManager.CallOpts)
}

// GetTokenPoolSupplyData is a free data retrieval call binding the contract method 0x2dafdbbf.
//
// Solidity: function getTokenPoolSupplyData() view returns(uint256 _lockedFundsWei, uint256 _totalInflationAuthorizedWei, uint256 _totalClaimedWei)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) GetTokenPoolSupplyData() (struct {
	LockedFundsWei              *big.Int
	TotalInflationAuthorizedWei *big.Int
	TotalClaimedWei             *big.Int
}, error) {
	return _TeeRewardOffersManager.Contract.GetTokenPoolSupplyData(&_TeeRewardOffersManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) Governance() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.Governance(&_TeeRewardOffersManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) Governance() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.Governance(&_TeeRewardOffersManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) GovernanceSettings() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.GovernanceSettings(&_TeeRewardOffersManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.GovernanceSettings(&_TeeRewardOffersManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) Implementation() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.Implementation(&_TeeRewardOffersManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) Implementation() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.Implementation(&_TeeRewardOffersManager.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeRewardOffersManager.Contract.IsExecutor(&_TeeRewardOffersManager.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeRewardOffersManager.Contract.IsExecutor(&_TeeRewardOffersManager.CallOpts, _address)
}

// LastInflationAuthorizationReceivedTs is a free data retrieval call binding the contract method 0x473252c4.
//
// Solidity: function lastInflationAuthorizationReceivedTs() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) LastInflationAuthorizationReceivedTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "lastInflationAuthorizationReceivedTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastInflationAuthorizationReceivedTs is a free data retrieval call binding the contract method 0x473252c4.
//
// Solidity: function lastInflationAuthorizationReceivedTs() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) LastInflationAuthorizationReceivedTs() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.LastInflationAuthorizationReceivedTs(&_TeeRewardOffersManager.CallOpts)
}

// LastInflationAuthorizationReceivedTs is a free data retrieval call binding the contract method 0x473252c4.
//
// Solidity: function lastInflationAuthorizationReceivedTs() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) LastInflationAuthorizationReceivedTs() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.LastInflationAuthorizationReceivedTs(&_TeeRewardOffersManager.CallOpts)
}

// LastInflationReceivedTs is a free data retrieval call binding the contract method 0x12afcf0b.
//
// Solidity: function lastInflationReceivedTs() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) LastInflationReceivedTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "lastInflationReceivedTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastInflationReceivedTs is a free data retrieval call binding the contract method 0x12afcf0b.
//
// Solidity: function lastInflationReceivedTs() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) LastInflationReceivedTs() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.LastInflationReceivedTs(&_TeeRewardOffersManager.CallOpts)
}

// LastInflationReceivedTs is a free data retrieval call binding the contract method 0x12afcf0b.
//
// Solidity: function lastInflationReceivedTs() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) LastInflationReceivedTs() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.LastInflationReceivedTs(&_TeeRewardOffersManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) ProductionMode() (bool, error) {
	return _TeeRewardOffersManager.Contract.ProductionMode(&_TeeRewardOffersManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) ProductionMode() (bool, error) {
	return _TeeRewardOffersManager.Contract.ProductionMode(&_TeeRewardOffersManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeRewardOffersManager.Contract.ProxiableUUID(&_TeeRewardOffersManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeRewardOffersManager.Contract.ProxiableUUID(&_TeeRewardOffersManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) RewardManager() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.RewardManager(&_TeeRewardOffersManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) RewardManager() (common.Address, error) {
	return _TeeRewardOffersManager.Contract.RewardManager(&_TeeRewardOffersManager.CallOpts)
}

// TeeOwnersPPM is a free data retrieval call binding the contract method 0xeeddae62.
//
// Solidity: function teeOwnersPPM() view returns(uint24)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) TeeOwnersPPM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "teeOwnersPPM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeeOwnersPPM is a free data retrieval call binding the contract method 0xeeddae62.
//
// Solidity: function teeOwnersPPM() view returns(uint24)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) TeeOwnersPPM() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.TeeOwnersPPM(&_TeeRewardOffersManager.CallOpts)
}

// TeeOwnersPPM is a free data retrieval call binding the contract method 0xeeddae62.
//
// Solidity: function teeOwnersPPM() view returns(uint24)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) TeeOwnersPPM() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.TeeOwnersPPM(&_TeeRewardOffersManager.CallOpts)
}

// TotalInflationAuthorizedWei is a free data retrieval call binding the contract method 0xd0c1c393.
//
// Solidity: function totalInflationAuthorizedWei() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) TotalInflationAuthorizedWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "totalInflationAuthorizedWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInflationAuthorizedWei is a free data retrieval call binding the contract method 0xd0c1c393.
//
// Solidity: function totalInflationAuthorizedWei() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) TotalInflationAuthorizedWei() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.TotalInflationAuthorizedWei(&_TeeRewardOffersManager.CallOpts)
}

// TotalInflationAuthorizedWei is a free data retrieval call binding the contract method 0xd0c1c393.
//
// Solidity: function totalInflationAuthorizedWei() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) TotalInflationAuthorizedWei() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.TotalInflationAuthorizedWei(&_TeeRewardOffersManager.CallOpts)
}

// TotalInflationReceivedWei is a free data retrieval call binding the contract method 0xa5555aea.
//
// Solidity: function totalInflationReceivedWei() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) TotalInflationReceivedWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "totalInflationReceivedWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInflationReceivedWei is a free data retrieval call binding the contract method 0xa5555aea.
//
// Solidity: function totalInflationReceivedWei() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) TotalInflationReceivedWei() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.TotalInflationReceivedWei(&_TeeRewardOffersManager.CallOpts)
}

// TotalInflationReceivedWei is a free data retrieval call binding the contract method 0xa5555aea.
//
// Solidity: function totalInflationReceivedWei() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) TotalInflationReceivedWei() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.TotalInflationReceivedWei(&_TeeRewardOffersManager.CallOpts)
}

// TotalInflationRewardsOfferedWei is a free data retrieval call binding the contract method 0xbd76b69c.
//
// Solidity: function totalInflationRewardsOfferedWei() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCaller) TotalInflationRewardsOfferedWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeRewardOffersManager.contract.Call(opts, &out, "totalInflationRewardsOfferedWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInflationRewardsOfferedWei is a free data retrieval call binding the contract method 0xbd76b69c.
//
// Solidity: function totalInflationRewardsOfferedWei() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) TotalInflationRewardsOfferedWei() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.TotalInflationRewardsOfferedWei(&_TeeRewardOffersManager.CallOpts)
}

// TotalInflationRewardsOfferedWei is a free data retrieval call binding the contract method 0xbd76b69c.
//
// Solidity: function totalInflationRewardsOfferedWei() view returns(uint256)
func (_TeeRewardOffersManager *TeeRewardOffersManagerCallerSession) TotalInflationRewardsOfferedWei() (*big.Int, error) {
	return _TeeRewardOffersManager.Contract.TotalInflationRewardsOfferedWei(&_TeeRewardOffersManager.CallOpts)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _encodedCall []byte) (*types.Transaction, error) {
	return _TeeRewardOffersManager.contract.Transact(opts, "cancelGovernanceCall", _encodedCall)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) CancelGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.CancelGovernanceCall(&_TeeRewardOffersManager.TransactOpts, _encodedCall)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactorSession) CancelGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.CancelGovernanceCall(&_TeeRewardOffersManager.TransactOpts, _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _encodedCall []byte) (*types.Transaction, error) {
	return _TeeRewardOffersManager.contract.Transact(opts, "executeGovernanceCall", _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) ExecuteGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.ExecuteGovernanceCall(&_TeeRewardOffersManager.TransactOpts, _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactorSession) ExecuteGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.ExecuteGovernanceCall(&_TeeRewardOffersManager.TransactOpts, _encodedCall)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4627023.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint24 _teeOwnersPPM) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _teeOwnersPPM *big.Int) (*types.Transaction, error) {
	return _TeeRewardOffersManager.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater, _teeOwnersPPM)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4627023.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint24 _teeOwnersPPM) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _teeOwnersPPM *big.Int) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.Initialize(&_TeeRewardOffersManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _teeOwnersPPM)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4627023.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint24 _teeOwnersPPM) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _teeOwnersPPM *big.Int) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.Initialize(&_TeeRewardOffersManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _teeOwnersPPM)
}

// ReceiveInflation is a paid mutator transaction binding the contract method 0x06201f1d.
//
// Solidity: function receiveInflation() payable returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactor) ReceiveInflation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeRewardOffersManager.contract.Transact(opts, "receiveInflation")
}

// ReceiveInflation is a paid mutator transaction binding the contract method 0x06201f1d.
//
// Solidity: function receiveInflation() payable returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) ReceiveInflation() (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.ReceiveInflation(&_TeeRewardOffersManager.TransactOpts)
}

// ReceiveInflation is a paid mutator transaction binding the contract method 0x06201f1d.
//
// Solidity: function receiveInflation() payable returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactorSession) ReceiveInflation() (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.ReceiveInflation(&_TeeRewardOffersManager.TransactOpts)
}

// SetDailyAuthorizedInflation is a paid mutator transaction binding the contract method 0xe2739563.
//
// Solidity: function setDailyAuthorizedInflation(uint256 _toAuthorizeWei) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactor) SetDailyAuthorizedInflation(opts *bind.TransactOpts, _toAuthorizeWei *big.Int) (*types.Transaction, error) {
	return _TeeRewardOffersManager.contract.Transact(opts, "setDailyAuthorizedInflation", _toAuthorizeWei)
}

// SetDailyAuthorizedInflation is a paid mutator transaction binding the contract method 0xe2739563.
//
// Solidity: function setDailyAuthorizedInflation(uint256 _toAuthorizeWei) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) SetDailyAuthorizedInflation(_toAuthorizeWei *big.Int) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.SetDailyAuthorizedInflation(&_TeeRewardOffersManager.TransactOpts, _toAuthorizeWei)
}

// SetDailyAuthorizedInflation is a paid mutator transaction binding the contract method 0xe2739563.
//
// Solidity: function setDailyAuthorizedInflation(uint256 _toAuthorizeWei) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactorSession) SetDailyAuthorizedInflation(_toAuthorizeWei *big.Int) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.SetDailyAuthorizedInflation(&_TeeRewardOffersManager.TransactOpts, _toAuthorizeWei)
}

// SetTeeOwnersPPM is a paid mutator transaction binding the contract method 0x9d81e53d.
//
// Solidity: function setTeeOwnersPPM(uint24 _teeOwnersPPM) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactor) SetTeeOwnersPPM(opts *bind.TransactOpts, _teeOwnersPPM *big.Int) (*types.Transaction, error) {
	return _TeeRewardOffersManager.contract.Transact(opts, "setTeeOwnersPPM", _teeOwnersPPM)
}

// SetTeeOwnersPPM is a paid mutator transaction binding the contract method 0x9d81e53d.
//
// Solidity: function setTeeOwnersPPM(uint24 _teeOwnersPPM) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) SetTeeOwnersPPM(_teeOwnersPPM *big.Int) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.SetTeeOwnersPPM(&_TeeRewardOffersManager.TransactOpts, _teeOwnersPPM)
}

// SetTeeOwnersPPM is a paid mutator transaction binding the contract method 0x9d81e53d.
//
// Solidity: function setTeeOwnersPPM(uint24 _teeOwnersPPM) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactorSession) SetTeeOwnersPPM(_teeOwnersPPM *big.Int) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.SetTeeOwnersPPM(&_TeeRewardOffersManager.TransactOpts, _teeOwnersPPM)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeRewardOffersManager.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.SwitchToProductionMode(&_TeeRewardOffersManager.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.SwitchToProductionMode(&_TeeRewardOffersManager.TransactOpts)
}

// TriggerRewardEpochSwitchover is a paid mutator transaction binding the contract method 0x91f25679.
//
// Solidity: function triggerRewardEpochSwitchover(uint24 _currentRewardEpochId, uint64 _currentRewardEpochExpectedEndTs, uint64 _rewardEpochDurationSeconds) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactor) TriggerRewardEpochSwitchover(opts *bind.TransactOpts, _currentRewardEpochId *big.Int, _currentRewardEpochExpectedEndTs uint64, _rewardEpochDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeRewardOffersManager.contract.Transact(opts, "triggerRewardEpochSwitchover", _currentRewardEpochId, _currentRewardEpochExpectedEndTs, _rewardEpochDurationSeconds)
}

// TriggerRewardEpochSwitchover is a paid mutator transaction binding the contract method 0x91f25679.
//
// Solidity: function triggerRewardEpochSwitchover(uint24 _currentRewardEpochId, uint64 _currentRewardEpochExpectedEndTs, uint64 _rewardEpochDurationSeconds) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) TriggerRewardEpochSwitchover(_currentRewardEpochId *big.Int, _currentRewardEpochExpectedEndTs uint64, _rewardEpochDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.TriggerRewardEpochSwitchover(&_TeeRewardOffersManager.TransactOpts, _currentRewardEpochId, _currentRewardEpochExpectedEndTs, _rewardEpochDurationSeconds)
}

// TriggerRewardEpochSwitchover is a paid mutator transaction binding the contract method 0x91f25679.
//
// Solidity: function triggerRewardEpochSwitchover(uint24 _currentRewardEpochId, uint64 _currentRewardEpochExpectedEndTs, uint64 _rewardEpochDurationSeconds) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactorSession) TriggerRewardEpochSwitchover(_currentRewardEpochId *big.Int, _currentRewardEpochExpectedEndTs uint64, _rewardEpochDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.TriggerRewardEpochSwitchover(&_TeeRewardOffersManager.TransactOpts, _currentRewardEpochId, _currentRewardEpochExpectedEndTs, _rewardEpochDurationSeconds)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeRewardOffersManager.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.UpdateContractAddresses(&_TeeRewardOffersManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.UpdateContractAddresses(&_TeeRewardOffersManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeRewardOffersManager.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.UpgradeToAndCall(&_TeeRewardOffersManager.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeRewardOffersManager *TeeRewardOffersManagerTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeRewardOffersManager.Contract.UpgradeToAndCall(&_TeeRewardOffersManager.TransactOpts, _newImplementation, _data)
}

// TeeRewardOffersManagerDailyAuthorizedInflationSetIterator is returned from FilterDailyAuthorizedInflationSet and is used to iterate over the raw logs and unpacked data for DailyAuthorizedInflationSet events raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerDailyAuthorizedInflationSetIterator struct {
	Event *TeeRewardOffersManagerDailyAuthorizedInflationSet // Event containing the contract specifics and raw log

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
func (it *TeeRewardOffersManagerDailyAuthorizedInflationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRewardOffersManagerDailyAuthorizedInflationSet)
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
		it.Event = new(TeeRewardOffersManagerDailyAuthorizedInflationSet)
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
func (it *TeeRewardOffersManagerDailyAuthorizedInflationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRewardOffersManagerDailyAuthorizedInflationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRewardOffersManagerDailyAuthorizedInflationSet represents a DailyAuthorizedInflationSet event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerDailyAuthorizedInflationSet struct {
	AuthorizedAmountWei *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDailyAuthorizedInflationSet is a free log retrieval operation binding the contract event 0x187f32a0f765499f15b3bb52ed0aebf6015059f230f2ace7e701e60a47669595.
//
// Solidity: event DailyAuthorizedInflationSet(uint256 authorizedAmountWei)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) FilterDailyAuthorizedInflationSet(opts *bind.FilterOpts) (*TeeRewardOffersManagerDailyAuthorizedInflationSetIterator, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.FilterLogs(opts, "DailyAuthorizedInflationSet")
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManagerDailyAuthorizedInflationSetIterator{contract: _TeeRewardOffersManager.contract, event: "DailyAuthorizedInflationSet", logs: logs, sub: sub}, nil
}

// WatchDailyAuthorizedInflationSet is a free log subscription operation binding the contract event 0x187f32a0f765499f15b3bb52ed0aebf6015059f230f2ace7e701e60a47669595.
//
// Solidity: event DailyAuthorizedInflationSet(uint256 authorizedAmountWei)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) WatchDailyAuthorizedInflationSet(opts *bind.WatchOpts, sink chan<- *TeeRewardOffersManagerDailyAuthorizedInflationSet) (event.Subscription, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.WatchLogs(opts, "DailyAuthorizedInflationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRewardOffersManagerDailyAuthorizedInflationSet)
				if err := _TeeRewardOffersManager.contract.UnpackLog(event, "DailyAuthorizedInflationSet", log); err != nil {
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

// ParseDailyAuthorizedInflationSet is a log parse operation binding the contract event 0x187f32a0f765499f15b3bb52ed0aebf6015059f230f2ace7e701e60a47669595.
//
// Solidity: event DailyAuthorizedInflationSet(uint256 authorizedAmountWei)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) ParseDailyAuthorizedInflationSet(log types.Log) (*TeeRewardOffersManagerDailyAuthorizedInflationSet, error) {
	event := new(TeeRewardOffersManagerDailyAuthorizedInflationSet)
	if err := _TeeRewardOffersManager.contract.UnpackLog(event, "DailyAuthorizedInflationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRewardOffersManagerGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerGovernanceCallTimelockedIterator struct {
	Event *TeeRewardOffersManagerGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeeRewardOffersManagerGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRewardOffersManagerGovernanceCallTimelocked)
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
		it.Event = new(TeeRewardOffersManagerGovernanceCallTimelocked)
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
func (it *TeeRewardOffersManagerGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRewardOffersManagerGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRewardOffersManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeRewardOffersManagerGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManagerGovernanceCallTimelockedIterator{contract: _TeeRewardOffersManager.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeRewardOffersManagerGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRewardOffersManagerGovernanceCallTimelocked)
				if err := _TeeRewardOffersManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeRewardOffersManagerGovernanceCallTimelocked, error) {
	event := new(TeeRewardOffersManagerGovernanceCallTimelocked)
	if err := _TeeRewardOffersManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRewardOffersManagerGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerGovernanceInitialisedIterator struct {
	Event *TeeRewardOffersManagerGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeeRewardOffersManagerGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRewardOffersManagerGovernanceInitialised)
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
		it.Event = new(TeeRewardOffersManagerGovernanceInitialised)
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
func (it *TeeRewardOffersManagerGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRewardOffersManagerGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRewardOffersManagerGovernanceInitialised represents a GovernanceInitialised event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeRewardOffersManagerGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManagerGovernanceInitialisedIterator{contract: _TeeRewardOffersManager.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeRewardOffersManagerGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRewardOffersManagerGovernanceInitialised)
				if err := _TeeRewardOffersManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) ParseGovernanceInitialised(log types.Log) (*TeeRewardOffersManagerGovernanceInitialised, error) {
	event := new(TeeRewardOffersManagerGovernanceInitialised)
	if err := _TeeRewardOffersManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRewardOffersManagerGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerGovernedProductionModeEnteredIterator struct {
	Event *TeeRewardOffersManagerGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeeRewardOffersManagerGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRewardOffersManagerGovernedProductionModeEntered)
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
		it.Event = new(TeeRewardOffersManagerGovernedProductionModeEntered)
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
func (it *TeeRewardOffersManagerGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRewardOffersManagerGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRewardOffersManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeeRewardOffersManagerGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManagerGovernedProductionModeEnteredIterator{contract: _TeeRewardOffersManager.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeeRewardOffersManagerGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRewardOffersManagerGovernedProductionModeEntered)
				if err := _TeeRewardOffersManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeeRewardOffersManagerGovernedProductionModeEntered, error) {
	event := new(TeeRewardOffersManagerGovernedProductionModeEntered)
	if err := _TeeRewardOffersManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRewardOffersManagerInflationReceivedIterator is returned from FilterInflationReceived and is used to iterate over the raw logs and unpacked data for InflationReceived events raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerInflationReceivedIterator struct {
	Event *TeeRewardOffersManagerInflationReceived // Event containing the contract specifics and raw log

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
func (it *TeeRewardOffersManagerInflationReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRewardOffersManagerInflationReceived)
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
		it.Event = new(TeeRewardOffersManagerInflationReceived)
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
func (it *TeeRewardOffersManagerInflationReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRewardOffersManagerInflationReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRewardOffersManagerInflationReceived represents a InflationReceived event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerInflationReceived struct {
	AmountReceivedWei *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInflationReceived is a free log retrieval operation binding the contract event 0x95c4e29cc99bc027cfc3cd719d6fd973d5f0317061885fbb322b9b17d8d35d37.
//
// Solidity: event InflationReceived(uint256 amountReceivedWei)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) FilterInflationReceived(opts *bind.FilterOpts) (*TeeRewardOffersManagerInflationReceivedIterator, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.FilterLogs(opts, "InflationReceived")
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManagerInflationReceivedIterator{contract: _TeeRewardOffersManager.contract, event: "InflationReceived", logs: logs, sub: sub}, nil
}

// WatchInflationReceived is a free log subscription operation binding the contract event 0x95c4e29cc99bc027cfc3cd719d6fd973d5f0317061885fbb322b9b17d8d35d37.
//
// Solidity: event InflationReceived(uint256 amountReceivedWei)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) WatchInflationReceived(opts *bind.WatchOpts, sink chan<- *TeeRewardOffersManagerInflationReceived) (event.Subscription, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.WatchLogs(opts, "InflationReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRewardOffersManagerInflationReceived)
				if err := _TeeRewardOffersManager.contract.UnpackLog(event, "InflationReceived", log); err != nil {
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

// ParseInflationReceived is a log parse operation binding the contract event 0x95c4e29cc99bc027cfc3cd719d6fd973d5f0317061885fbb322b9b17d8d35d37.
//
// Solidity: event InflationReceived(uint256 amountReceivedWei)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) ParseInflationReceived(log types.Log) (*TeeRewardOffersManagerInflationReceived, error) {
	event := new(TeeRewardOffersManagerInflationReceived)
	if err := _TeeRewardOffersManager.contract.UnpackLog(event, "InflationReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRewardOffersManagerInflationRewardsOfferedIterator is returned from FilterInflationRewardsOffered and is used to iterate over the raw logs and unpacked data for InflationRewardsOffered events raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerInflationRewardsOfferedIterator struct {
	Event *TeeRewardOffersManagerInflationRewardsOffered // Event containing the contract specifics and raw log

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
func (it *TeeRewardOffersManagerInflationRewardsOfferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRewardOffersManagerInflationRewardsOffered)
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
		it.Event = new(TeeRewardOffersManagerInflationRewardsOffered)
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
func (it *TeeRewardOffersManagerInflationRewardsOfferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRewardOffersManagerInflationRewardsOfferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRewardOffersManagerInflationRewardsOffered represents a InflationRewardsOffered event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerInflationRewardsOffered struct {
	RewardEpochId *big.Int
	Amount        *big.Int
	TeeOwnersPPM  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInflationRewardsOffered is a free log retrieval operation binding the contract event 0x380801a5750b79ca9b887859642bc58ce7853a8a1d2ecd9778bd34fed0e9bc18.
//
// Solidity: event InflationRewardsOffered(uint24 indexed rewardEpochId, uint256 amount, uint256 teeOwnersPPM)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) FilterInflationRewardsOffered(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*TeeRewardOffersManagerInflationRewardsOfferedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeRewardOffersManager.contract.FilterLogs(opts, "InflationRewardsOffered", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManagerInflationRewardsOfferedIterator{contract: _TeeRewardOffersManager.contract, event: "InflationRewardsOffered", logs: logs, sub: sub}, nil
}

// WatchInflationRewardsOffered is a free log subscription operation binding the contract event 0x380801a5750b79ca9b887859642bc58ce7853a8a1d2ecd9778bd34fed0e9bc18.
//
// Solidity: event InflationRewardsOffered(uint24 indexed rewardEpochId, uint256 amount, uint256 teeOwnersPPM)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) WatchInflationRewardsOffered(opts *bind.WatchOpts, sink chan<- *TeeRewardOffersManagerInflationRewardsOffered, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeRewardOffersManager.contract.WatchLogs(opts, "InflationRewardsOffered", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRewardOffersManagerInflationRewardsOffered)
				if err := _TeeRewardOffersManager.contract.UnpackLog(event, "InflationRewardsOffered", log); err != nil {
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

// ParseInflationRewardsOffered is a log parse operation binding the contract event 0x380801a5750b79ca9b887859642bc58ce7853a8a1d2ecd9778bd34fed0e9bc18.
//
// Solidity: event InflationRewardsOffered(uint24 indexed rewardEpochId, uint256 amount, uint256 teeOwnersPPM)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) ParseInflationRewardsOffered(log types.Log) (*TeeRewardOffersManagerInflationRewardsOffered, error) {
	event := new(TeeRewardOffersManagerInflationRewardsOffered)
	if err := _TeeRewardOffersManager.contract.UnpackLog(event, "InflationRewardsOffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRewardOffersManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerInitializedIterator struct {
	Event *TeeRewardOffersManagerInitialized // Event containing the contract specifics and raw log

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
func (it *TeeRewardOffersManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRewardOffersManagerInitialized)
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
		it.Event = new(TeeRewardOffersManagerInitialized)
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
func (it *TeeRewardOffersManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRewardOffersManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRewardOffersManagerInitialized represents a Initialized event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*TeeRewardOffersManagerInitializedIterator, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManagerInitializedIterator{contract: _TeeRewardOffersManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TeeRewardOffersManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRewardOffersManagerInitialized)
				if err := _TeeRewardOffersManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) ParseInitialized(log types.Log) (*TeeRewardOffersManagerInitialized, error) {
	event := new(TeeRewardOffersManagerInitialized)
	if err := _TeeRewardOffersManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRewardOffersManagerTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerTimelockedGovernanceCallCanceledIterator struct {
	Event *TeeRewardOffersManagerTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeeRewardOffersManagerTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRewardOffersManagerTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeeRewardOffersManagerTimelockedGovernanceCallCanceled)
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
func (it *TeeRewardOffersManagerTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRewardOffersManagerTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRewardOffersManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeeRewardOffersManagerTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManagerTimelockedGovernanceCallCanceledIterator{contract: _TeeRewardOffersManager.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeeRewardOffersManagerTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRewardOffersManagerTimelockedGovernanceCallCanceled)
				if err := _TeeRewardOffersManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeeRewardOffersManagerTimelockedGovernanceCallCanceled, error) {
	event := new(TeeRewardOffersManagerTimelockedGovernanceCallCanceled)
	if err := _TeeRewardOffersManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRewardOffersManagerTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerTimelockedGovernanceCallExecutedIterator struct {
	Event *TeeRewardOffersManagerTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeeRewardOffersManagerTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRewardOffersManagerTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeeRewardOffersManagerTimelockedGovernanceCallExecuted)
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
func (it *TeeRewardOffersManagerTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRewardOffersManagerTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRewardOffersManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeeRewardOffersManagerTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManagerTimelockedGovernanceCallExecutedIterator{contract: _TeeRewardOffersManager.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeeRewardOffersManagerTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeeRewardOffersManager.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRewardOffersManagerTimelockedGovernanceCallExecuted)
				if err := _TeeRewardOffersManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeeRewardOffersManagerTimelockedGovernanceCallExecuted, error) {
	event := new(TeeRewardOffersManagerTimelockedGovernanceCallExecuted)
	if err := _TeeRewardOffersManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRewardOffersManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerUpgradedIterator struct {
	Event *TeeRewardOffersManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *TeeRewardOffersManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRewardOffersManagerUpgraded)
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
		it.Event = new(TeeRewardOffersManagerUpgraded)
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
func (it *TeeRewardOffersManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRewardOffersManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRewardOffersManagerUpgraded represents a Upgraded event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeeRewardOffersManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeRewardOffersManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeeRewardOffersManagerUpgradedIterator{contract: _TeeRewardOffersManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeeRewardOffersManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeRewardOffersManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRewardOffersManagerUpgraded)
				if err := _TeeRewardOffersManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeeRewardOffersManager *TeeRewardOffersManagerFilterer) ParseUpgraded(log types.Log) (*TeeRewardOffersManagerUpgraded, error) {
	event := new(TeeRewardOffersManagerUpgraded)
	if err := _TeeRewardOffersManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
