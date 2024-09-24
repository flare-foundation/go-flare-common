// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fumanager

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

// IFastUpdateIncentiveManagerIncentiveOffer is an auto generated low-level Go binding around an user-defined struct.
type IFastUpdateIncentiveManagerIncentiveOffer struct {
	RangeIncrease *big.Int
	RangeLimit    *big.Int
}

// IFastUpdatesConfigurationFeedConfiguration is an auto generated low-level Go binding around an user-defined struct.
type IFastUpdatesConfigurationFeedConfiguration struct {
	FeedId          [21]byte
	RewardBandValue uint32
	InflationShare  *big.Int
}

// FUManagerMetaData contains all meta data concerning the FUManager contract.
var FUManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"SampleSize\",\"name\":\"_ss\",\"type\":\"uint256\"},{\"internalType\":\"Range\",\"name\":\"_r\",\"type\":\"uint256\"},{\"internalType\":\"SampleSize\",\"name\":\"_sil\",\"type\":\"uint256\"},{\"internalType\":\"Range\",\"name\":\"_ril\",\"type\":\"uint256\"},{\"internalType\":\"Fee\",\"name\":\"_x\",\"type\":\"uint256\"},{\"internalType\":\"Fee\",\"name\":\"_rip\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dur\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"authorizedAmountWei\",\"type\":\"uint256\"}],\"name\":\"DailyAuthorizedInflationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"Range\",\"name\":\"rangeIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"SampleSize\",\"name\":\"sampleSizeIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"Fee\",\"name\":\"offerAmount\",\"type\":\"uint256\"}],\"name\":\"IncentiveOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedWei\",\"type\":\"uint256\"}],\"name\":\"InflationReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"bytes21\",\"name\":\"feedId\",\"type\":\"bytes21\"},{\"internalType\":\"uint32\",\"name\":\"rewardBandValue\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"inflationShare\",\"type\":\"uint24\"}],\"indexed\":false,\"internalType\":\"structIFastUpdatesConfiguration.FeedConfiguration[]\",\"name\":\"feedConfigurations\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InflationRewardsOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"advance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyAuthorizedInflation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fastUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fastUpdatesConfiguration\",\"outputs\":[{\"internalType\":\"contractIFastUpdatesConfiguration\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseScale\",\"outputs\":[{\"internalType\":\"Scale\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSampleSizeIncreasePrice\",\"outputs\":[{\"internalType\":\"Fee\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedSampleSize\",\"outputs\":[{\"internalType\":\"SampleSize\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIncentiveDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInflationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrecision\",\"outputs\":[{\"internalType\":\"Precision\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRange\",\"outputs\":[{\"internalType\":\"Range\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getScale\",\"outputs\":[{\"internalType\":\"Scale\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenPoolSupplyData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockedFundsWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalInflationAuthorizedWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalClaimedWei\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInflationAuthorizationReceivedTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInflationReceivedTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Range\",\"name\":\"rangeIncrease\",\"type\":\"uint256\"},{\"internalType\":\"Range\",\"name\":\"rangeLimit\",\"type\":\"uint256\"}],\"internalType\":\"structIFastUpdateIncentiveManager.IncentiveOffer\",\"name\":\"_offer\",\"type\":\"tuple\"}],\"name\":\"offerIncentive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rangeIncreaseLimit\",\"outputs\":[{\"internalType\":\"Range\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rangeIncreasePrice\",\"outputs\":[{\"internalType\":\"Fee\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveInflation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sampleIncreaseLimit\",\"outputs\":[{\"internalType\":\"SampleSize\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_toAuthorizeWei\",\"type\":\"uint256\"}],\"name\":\"setDailyAuthorizedInflation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"SampleSize\",\"name\":\"_ss\",\"type\":\"uint256\"},{\"internalType\":\"Range\",\"name\":\"_r\",\"type\":\"uint256\"},{\"internalType\":\"Fee\",\"name\":\"_x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dur\",\"type\":\"uint256\"}],\"name\":\"setIncentiveParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Range\",\"name\":\"_lim\",\"type\":\"uint256\"}],\"name\":\"setRangeIncreaseLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Fee\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setRangeIncreasePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"SampleSize\",\"name\":\"_lim\",\"type\":\"uint256\"}],\"name\":\"setSampleIncreaseLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationAuthorizedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationReceivedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationRewardsOfferedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_currentRewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_currentRewardEpochExpectedEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardEpochDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"triggerRewardEpochSwitchover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FUManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use FUManagerMetaData.ABI instead.
var FUManagerABI = FUManagerMetaData.ABI

// FUManager is an auto generated Go binding around an Ethereum contract.
type FUManager struct {
	FUManagerCaller     // Read-only binding to the contract
	FUManagerTransactor // Write-only binding to the contract
	FUManagerFilterer   // Log filterer for contract events
}

// FUManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FUManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FUManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FUManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FUManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FUManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FUManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FUManagerSession struct {
	Contract     *FUManager        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FUManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FUManagerCallerSession struct {
	Contract *FUManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FUManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FUManagerTransactorSession struct {
	Contract     *FUManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FUManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FUManagerRaw struct {
	Contract *FUManager // Generic contract binding to access the raw methods on
}

// FUManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FUManagerCallerRaw struct {
	Contract *FUManagerCaller // Generic read-only contract binding to access the raw methods on
}

// FUManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FUManagerTransactorRaw struct {
	Contract *FUManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFUManager creates a new instance of FUManager, bound to a specific deployed contract.
func NewFUManager(address common.Address, backend bind.ContractBackend) (*FUManager, error) {
	contract, err := bindFUManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FUManager{FUManagerCaller: FUManagerCaller{contract: contract}, FUManagerTransactor: FUManagerTransactor{contract: contract}, FUManagerFilterer: FUManagerFilterer{contract: contract}}, nil
}

// NewFUManagerCaller creates a new read-only instance of FUManager, bound to a specific deployed contract.
func NewFUManagerCaller(address common.Address, caller bind.ContractCaller) (*FUManagerCaller, error) {
	contract, err := bindFUManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FUManagerCaller{contract: contract}, nil
}

// NewFUManagerTransactor creates a new write-only instance of FUManager, bound to a specific deployed contract.
func NewFUManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FUManagerTransactor, error) {
	contract, err := bindFUManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FUManagerTransactor{contract: contract}, nil
}

// NewFUManagerFilterer creates a new log filterer instance of FUManager, bound to a specific deployed contract.
func NewFUManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FUManagerFilterer, error) {
	contract, err := bindFUManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FUManagerFilterer{contract: contract}, nil
}

// bindFUManager binds a generic wrapper to an already deployed contract.
func bindFUManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FUManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FUManager *FUManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FUManager.Contract.FUManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FUManager *FUManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FUManager.Contract.FUManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FUManager *FUManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FUManager.Contract.FUManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FUManager *FUManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FUManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FUManager *FUManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FUManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FUManager *FUManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FUManager.Contract.contract.Transact(opts, method, params...)
}

// DailyAuthorizedInflation is a free data retrieval call binding the contract method 0x708e34ce.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (_FUManager *FUManagerCaller) DailyAuthorizedInflation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "dailyAuthorizedInflation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyAuthorizedInflation is a free data retrieval call binding the contract method 0x708e34ce.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (_FUManager *FUManagerSession) DailyAuthorizedInflation() (*big.Int, error) {
	return _FUManager.Contract.DailyAuthorizedInflation(&_FUManager.CallOpts)
}

// DailyAuthorizedInflation is a free data retrieval call binding the contract method 0x708e34ce.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (_FUManager *FUManagerCallerSession) DailyAuthorizedInflation() (*big.Int, error) {
	return _FUManager.Contract.DailyAuthorizedInflation(&_FUManager.CallOpts)
}

// FastUpdater is a free data retrieval call binding the contract method 0xd29a4fa9.
//
// Solidity: function fastUpdater() view returns(address)
func (_FUManager *FUManagerCaller) FastUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "fastUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FastUpdater is a free data retrieval call binding the contract method 0xd29a4fa9.
//
// Solidity: function fastUpdater() view returns(address)
func (_FUManager *FUManagerSession) FastUpdater() (common.Address, error) {
	return _FUManager.Contract.FastUpdater(&_FUManager.CallOpts)
}

// FastUpdater is a free data retrieval call binding the contract method 0xd29a4fa9.
//
// Solidity: function fastUpdater() view returns(address)
func (_FUManager *FUManagerCallerSession) FastUpdater() (common.Address, error) {
	return _FUManager.Contract.FastUpdater(&_FUManager.CallOpts)
}

// FastUpdatesConfiguration is a free data retrieval call binding the contract method 0xc10f489a.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (_FUManager *FUManagerCaller) FastUpdatesConfiguration(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "fastUpdatesConfiguration")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FastUpdatesConfiguration is a free data retrieval call binding the contract method 0xc10f489a.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (_FUManager *FUManagerSession) FastUpdatesConfiguration() (common.Address, error) {
	return _FUManager.Contract.FastUpdatesConfiguration(&_FUManager.CallOpts)
}

// FastUpdatesConfiguration is a free data retrieval call binding the contract method 0xc10f489a.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (_FUManager *FUManagerCallerSession) FastUpdatesConfiguration() (common.Address, error) {
	return _FUManager.Contract.FastUpdatesConfiguration(&_FUManager.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_FUManager *FUManagerCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_FUManager *FUManagerSession) FlareSystemsManager() (common.Address, error) {
	return _FUManager.Contract.FlareSystemsManager(&_FUManager.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_FUManager *FUManagerCallerSession) FlareSystemsManager() (common.Address, error) {
	return _FUManager.Contract.FlareSystemsManager(&_FUManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FUManager *FUManagerCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FUManager *FUManagerSession) GetAddressUpdater() (common.Address, error) {
	return _FUManager.Contract.GetAddressUpdater(&_FUManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FUManager *FUManagerCallerSession) GetAddressUpdater() (common.Address, error) {
	return _FUManager.Contract.GetAddressUpdater(&_FUManager.CallOpts)
}

// GetBaseScale is a free data retrieval call binding the contract method 0x7a68533f.
//
// Solidity: function getBaseScale() view returns(uint256)
func (_FUManager *FUManagerCaller) GetBaseScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "getBaseScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseScale is a free data retrieval call binding the contract method 0x7a68533f.
//
// Solidity: function getBaseScale() view returns(uint256)
func (_FUManager *FUManagerSession) GetBaseScale() (*big.Int, error) {
	return _FUManager.Contract.GetBaseScale(&_FUManager.CallOpts)
}

// GetBaseScale is a free data retrieval call binding the contract method 0x7a68533f.
//
// Solidity: function getBaseScale() view returns(uint256)
func (_FUManager *FUManagerCallerSession) GetBaseScale() (*big.Int, error) {
	return _FUManager.Contract.GetBaseScale(&_FUManager.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FUManager *FUManagerCaller) GetContractName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "getContractName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FUManager *FUManagerSession) GetContractName() (string, error) {
	return _FUManager.Contract.GetContractName(&_FUManager.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FUManager *FUManagerCallerSession) GetContractName() (string, error) {
	return _FUManager.Contract.GetContractName(&_FUManager.CallOpts)
}

// GetCurrentSampleSizeIncreasePrice is a free data retrieval call binding the contract method 0x2de490c3.
//
// Solidity: function getCurrentSampleSizeIncreasePrice() view returns(uint256)
func (_FUManager *FUManagerCaller) GetCurrentSampleSizeIncreasePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "getCurrentSampleSizeIncreasePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentSampleSizeIncreasePrice is a free data retrieval call binding the contract method 0x2de490c3.
//
// Solidity: function getCurrentSampleSizeIncreasePrice() view returns(uint256)
func (_FUManager *FUManagerSession) GetCurrentSampleSizeIncreasePrice() (*big.Int, error) {
	return _FUManager.Contract.GetCurrentSampleSizeIncreasePrice(&_FUManager.CallOpts)
}

// GetCurrentSampleSizeIncreasePrice is a free data retrieval call binding the contract method 0x2de490c3.
//
// Solidity: function getCurrentSampleSizeIncreasePrice() view returns(uint256)
func (_FUManager *FUManagerCallerSession) GetCurrentSampleSizeIncreasePrice() (*big.Int, error) {
	return _FUManager.Contract.GetCurrentSampleSizeIncreasePrice(&_FUManager.CallOpts)
}

// GetExpectedBalance is a free data retrieval call binding the contract method 0xaf04cd3b.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (_FUManager *FUManagerCaller) GetExpectedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "getExpectedBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpectedBalance is a free data retrieval call binding the contract method 0xaf04cd3b.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (_FUManager *FUManagerSession) GetExpectedBalance() (*big.Int, error) {
	return _FUManager.Contract.GetExpectedBalance(&_FUManager.CallOpts)
}

// GetExpectedBalance is a free data retrieval call binding the contract method 0xaf04cd3b.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (_FUManager *FUManagerCallerSession) GetExpectedBalance() (*big.Int, error) {
	return _FUManager.Contract.GetExpectedBalance(&_FUManager.CallOpts)
}

// GetExpectedSampleSize is a free data retrieval call binding the contract method 0x6d62b413.
//
// Solidity: function getExpectedSampleSize() view returns(uint256)
func (_FUManager *FUManagerCaller) GetExpectedSampleSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "getExpectedSampleSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExpectedSampleSize is a free data retrieval call binding the contract method 0x6d62b413.
//
// Solidity: function getExpectedSampleSize() view returns(uint256)
func (_FUManager *FUManagerSession) GetExpectedSampleSize() (*big.Int, error) {
	return _FUManager.Contract.GetExpectedSampleSize(&_FUManager.CallOpts)
}

// GetExpectedSampleSize is a free data retrieval call binding the contract method 0x6d62b413.
//
// Solidity: function getExpectedSampleSize() view returns(uint256)
func (_FUManager *FUManagerCallerSession) GetExpectedSampleSize() (*big.Int, error) {
	return _FUManager.Contract.GetExpectedSampleSize(&_FUManager.CallOpts)
}

// GetIncentiveDuration is a free data retrieval call binding the contract method 0xdd8dca9f.
//
// Solidity: function getIncentiveDuration() view returns(uint256)
func (_FUManager *FUManagerCaller) GetIncentiveDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "getIncentiveDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIncentiveDuration is a free data retrieval call binding the contract method 0xdd8dca9f.
//
// Solidity: function getIncentiveDuration() view returns(uint256)
func (_FUManager *FUManagerSession) GetIncentiveDuration() (*big.Int, error) {
	return _FUManager.Contract.GetIncentiveDuration(&_FUManager.CallOpts)
}

// GetIncentiveDuration is a free data retrieval call binding the contract method 0xdd8dca9f.
//
// Solidity: function getIncentiveDuration() view returns(uint256)
func (_FUManager *FUManagerCallerSession) GetIncentiveDuration() (*big.Int, error) {
	return _FUManager.Contract.GetIncentiveDuration(&_FUManager.CallOpts)
}

// GetInflationAddress is a free data retrieval call binding the contract method 0xed39d3f8.
//
// Solidity: function getInflationAddress() view returns(address)
func (_FUManager *FUManagerCaller) GetInflationAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "getInflationAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInflationAddress is a free data retrieval call binding the contract method 0xed39d3f8.
//
// Solidity: function getInflationAddress() view returns(address)
func (_FUManager *FUManagerSession) GetInflationAddress() (common.Address, error) {
	return _FUManager.Contract.GetInflationAddress(&_FUManager.CallOpts)
}

// GetInflationAddress is a free data retrieval call binding the contract method 0xed39d3f8.
//
// Solidity: function getInflationAddress() view returns(address)
func (_FUManager *FUManagerCallerSession) GetInflationAddress() (common.Address, error) {
	return _FUManager.Contract.GetInflationAddress(&_FUManager.CallOpts)
}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() view returns(uint256)
func (_FUManager *FUManagerCaller) GetPrecision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "getPrecision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() view returns(uint256)
func (_FUManager *FUManagerSession) GetPrecision() (*big.Int, error) {
	return _FUManager.Contract.GetPrecision(&_FUManager.CallOpts)
}

// GetPrecision is a free data retrieval call binding the contract method 0x9670c0bc.
//
// Solidity: function getPrecision() view returns(uint256)
func (_FUManager *FUManagerCallerSession) GetPrecision() (*big.Int, error) {
	return _FUManager.Contract.GetPrecision(&_FUManager.CallOpts)
}

// GetRange is a free data retrieval call binding the contract method 0x9b85961f.
//
// Solidity: function getRange() view returns(uint256)
func (_FUManager *FUManagerCaller) GetRange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "getRange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRange is a free data retrieval call binding the contract method 0x9b85961f.
//
// Solidity: function getRange() view returns(uint256)
func (_FUManager *FUManagerSession) GetRange() (*big.Int, error) {
	return _FUManager.Contract.GetRange(&_FUManager.CallOpts)
}

// GetRange is a free data retrieval call binding the contract method 0x9b85961f.
//
// Solidity: function getRange() view returns(uint256)
func (_FUManager *FUManagerCallerSession) GetRange() (*big.Int, error) {
	return _FUManager.Contract.GetRange(&_FUManager.CallOpts)
}

// GetScale is a free data retrieval call binding the contract method 0xb5cddab8.
//
// Solidity: function getScale() view returns(uint256)
func (_FUManager *FUManagerCaller) GetScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "getScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetScale is a free data retrieval call binding the contract method 0xb5cddab8.
//
// Solidity: function getScale() view returns(uint256)
func (_FUManager *FUManagerSession) GetScale() (*big.Int, error) {
	return _FUManager.Contract.GetScale(&_FUManager.CallOpts)
}

// GetScale is a free data retrieval call binding the contract method 0xb5cddab8.
//
// Solidity: function getScale() view returns(uint256)
func (_FUManager *FUManagerCallerSession) GetScale() (*big.Int, error) {
	return _FUManager.Contract.GetScale(&_FUManager.CallOpts)
}

// GetTokenPoolSupplyData is a free data retrieval call binding the contract method 0x2dafdbbf.
//
// Solidity: function getTokenPoolSupplyData() view returns(uint256 _lockedFundsWei, uint256 _totalInflationAuthorizedWei, uint256 _totalClaimedWei)
func (_FUManager *FUManagerCaller) GetTokenPoolSupplyData(opts *bind.CallOpts) (struct {
	LockedFundsWei              *big.Int
	TotalInflationAuthorizedWei *big.Int
	TotalClaimedWei             *big.Int
}, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "getTokenPoolSupplyData")

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
func (_FUManager *FUManagerSession) GetTokenPoolSupplyData() (struct {
	LockedFundsWei              *big.Int
	TotalInflationAuthorizedWei *big.Int
	TotalClaimedWei             *big.Int
}, error) {
	return _FUManager.Contract.GetTokenPoolSupplyData(&_FUManager.CallOpts)
}

// GetTokenPoolSupplyData is a free data retrieval call binding the contract method 0x2dafdbbf.
//
// Solidity: function getTokenPoolSupplyData() view returns(uint256 _lockedFundsWei, uint256 _totalInflationAuthorizedWei, uint256 _totalClaimedWei)
func (_FUManager *FUManagerCallerSession) GetTokenPoolSupplyData() (struct {
	LockedFundsWei              *big.Int
	TotalInflationAuthorizedWei *big.Int
	TotalClaimedWei             *big.Int
}, error) {
	return _FUManager.Contract.GetTokenPoolSupplyData(&_FUManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FUManager *FUManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FUManager *FUManagerSession) Governance() (common.Address, error) {
	return _FUManager.Contract.Governance(&_FUManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FUManager *FUManagerCallerSession) Governance() (common.Address, error) {
	return _FUManager.Contract.Governance(&_FUManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FUManager *FUManagerCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FUManager *FUManagerSession) GovernanceSettings() (common.Address, error) {
	return _FUManager.Contract.GovernanceSettings(&_FUManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FUManager *FUManagerCallerSession) GovernanceSettings() (common.Address, error) {
	return _FUManager.Contract.GovernanceSettings(&_FUManager.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FUManager *FUManagerCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FUManager *FUManagerSession) IsExecutor(_address common.Address) (bool, error) {
	return _FUManager.Contract.IsExecutor(&_FUManager.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FUManager *FUManagerCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _FUManager.Contract.IsExecutor(&_FUManager.CallOpts, _address)
}

// LastInflationAuthorizationReceivedTs is a free data retrieval call binding the contract method 0x473252c4.
//
// Solidity: function lastInflationAuthorizationReceivedTs() view returns(uint256)
func (_FUManager *FUManagerCaller) LastInflationAuthorizationReceivedTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "lastInflationAuthorizationReceivedTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastInflationAuthorizationReceivedTs is a free data retrieval call binding the contract method 0x473252c4.
//
// Solidity: function lastInflationAuthorizationReceivedTs() view returns(uint256)
func (_FUManager *FUManagerSession) LastInflationAuthorizationReceivedTs() (*big.Int, error) {
	return _FUManager.Contract.LastInflationAuthorizationReceivedTs(&_FUManager.CallOpts)
}

// LastInflationAuthorizationReceivedTs is a free data retrieval call binding the contract method 0x473252c4.
//
// Solidity: function lastInflationAuthorizationReceivedTs() view returns(uint256)
func (_FUManager *FUManagerCallerSession) LastInflationAuthorizationReceivedTs() (*big.Int, error) {
	return _FUManager.Contract.LastInflationAuthorizationReceivedTs(&_FUManager.CallOpts)
}

// LastInflationReceivedTs is a free data retrieval call binding the contract method 0x12afcf0b.
//
// Solidity: function lastInflationReceivedTs() view returns(uint256)
func (_FUManager *FUManagerCaller) LastInflationReceivedTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "lastInflationReceivedTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastInflationReceivedTs is a free data retrieval call binding the contract method 0x12afcf0b.
//
// Solidity: function lastInflationReceivedTs() view returns(uint256)
func (_FUManager *FUManagerSession) LastInflationReceivedTs() (*big.Int, error) {
	return _FUManager.Contract.LastInflationReceivedTs(&_FUManager.CallOpts)
}

// LastInflationReceivedTs is a free data retrieval call binding the contract method 0x12afcf0b.
//
// Solidity: function lastInflationReceivedTs() view returns(uint256)
func (_FUManager *FUManagerCallerSession) LastInflationReceivedTs() (*big.Int, error) {
	return _FUManager.Contract.LastInflationReceivedTs(&_FUManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FUManager *FUManagerCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FUManager *FUManagerSession) ProductionMode() (bool, error) {
	return _FUManager.Contract.ProductionMode(&_FUManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FUManager *FUManagerCallerSession) ProductionMode() (bool, error) {
	return _FUManager.Contract.ProductionMode(&_FUManager.CallOpts)
}

// RangeIncreaseLimit is a free data retrieval call binding the contract method 0x74f3eff9.
//
// Solidity: function rangeIncreaseLimit() view returns(uint256)
func (_FUManager *FUManagerCaller) RangeIncreaseLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "rangeIncreaseLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RangeIncreaseLimit is a free data retrieval call binding the contract method 0x74f3eff9.
//
// Solidity: function rangeIncreaseLimit() view returns(uint256)
func (_FUManager *FUManagerSession) RangeIncreaseLimit() (*big.Int, error) {
	return _FUManager.Contract.RangeIncreaseLimit(&_FUManager.CallOpts)
}

// RangeIncreaseLimit is a free data retrieval call binding the contract method 0x74f3eff9.
//
// Solidity: function rangeIncreaseLimit() view returns(uint256)
func (_FUManager *FUManagerCallerSession) RangeIncreaseLimit() (*big.Int, error) {
	return _FUManager.Contract.RangeIncreaseLimit(&_FUManager.CallOpts)
}

// RangeIncreasePrice is a free data retrieval call binding the contract method 0x52545a7c.
//
// Solidity: function rangeIncreasePrice() view returns(uint256)
func (_FUManager *FUManagerCaller) RangeIncreasePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "rangeIncreasePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RangeIncreasePrice is a free data retrieval call binding the contract method 0x52545a7c.
//
// Solidity: function rangeIncreasePrice() view returns(uint256)
func (_FUManager *FUManagerSession) RangeIncreasePrice() (*big.Int, error) {
	return _FUManager.Contract.RangeIncreasePrice(&_FUManager.CallOpts)
}

// RangeIncreasePrice is a free data retrieval call binding the contract method 0x52545a7c.
//
// Solidity: function rangeIncreasePrice() view returns(uint256)
func (_FUManager *FUManagerCallerSession) RangeIncreasePrice() (*big.Int, error) {
	return _FUManager.Contract.RangeIncreasePrice(&_FUManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FUManager *FUManagerCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FUManager *FUManagerSession) RewardManager() (common.Address, error) {
	return _FUManager.Contract.RewardManager(&_FUManager.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FUManager *FUManagerCallerSession) RewardManager() (common.Address, error) {
	return _FUManager.Contract.RewardManager(&_FUManager.CallOpts)
}

// SampleIncreaseLimit is a free data retrieval call binding the contract method 0xd4ab8f94.
//
// Solidity: function sampleIncreaseLimit() view returns(uint256)
func (_FUManager *FUManagerCaller) SampleIncreaseLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "sampleIncreaseLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SampleIncreaseLimit is a free data retrieval call binding the contract method 0xd4ab8f94.
//
// Solidity: function sampleIncreaseLimit() view returns(uint256)
func (_FUManager *FUManagerSession) SampleIncreaseLimit() (*big.Int, error) {
	return _FUManager.Contract.SampleIncreaseLimit(&_FUManager.CallOpts)
}

// SampleIncreaseLimit is a free data retrieval call binding the contract method 0xd4ab8f94.
//
// Solidity: function sampleIncreaseLimit() view returns(uint256)
func (_FUManager *FUManagerCallerSession) SampleIncreaseLimit() (*big.Int, error) {
	return _FUManager.Contract.SampleIncreaseLimit(&_FUManager.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FUManager *FUManagerCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_FUManager *FUManagerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FUManager.Contract.TimelockedCalls(&_FUManager.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FUManager *FUManagerCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FUManager.Contract.TimelockedCalls(&_FUManager.CallOpts, selector)
}

// TotalInflationAuthorizedWei is a free data retrieval call binding the contract method 0xd0c1c393.
//
// Solidity: function totalInflationAuthorizedWei() view returns(uint256)
func (_FUManager *FUManagerCaller) TotalInflationAuthorizedWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "totalInflationAuthorizedWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInflationAuthorizedWei is a free data retrieval call binding the contract method 0xd0c1c393.
//
// Solidity: function totalInflationAuthorizedWei() view returns(uint256)
func (_FUManager *FUManagerSession) TotalInflationAuthorizedWei() (*big.Int, error) {
	return _FUManager.Contract.TotalInflationAuthorizedWei(&_FUManager.CallOpts)
}

// TotalInflationAuthorizedWei is a free data retrieval call binding the contract method 0xd0c1c393.
//
// Solidity: function totalInflationAuthorizedWei() view returns(uint256)
func (_FUManager *FUManagerCallerSession) TotalInflationAuthorizedWei() (*big.Int, error) {
	return _FUManager.Contract.TotalInflationAuthorizedWei(&_FUManager.CallOpts)
}

// TotalInflationReceivedWei is a free data retrieval call binding the contract method 0xa5555aea.
//
// Solidity: function totalInflationReceivedWei() view returns(uint256)
func (_FUManager *FUManagerCaller) TotalInflationReceivedWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "totalInflationReceivedWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInflationReceivedWei is a free data retrieval call binding the contract method 0xa5555aea.
//
// Solidity: function totalInflationReceivedWei() view returns(uint256)
func (_FUManager *FUManagerSession) TotalInflationReceivedWei() (*big.Int, error) {
	return _FUManager.Contract.TotalInflationReceivedWei(&_FUManager.CallOpts)
}

// TotalInflationReceivedWei is a free data retrieval call binding the contract method 0xa5555aea.
//
// Solidity: function totalInflationReceivedWei() view returns(uint256)
func (_FUManager *FUManagerCallerSession) TotalInflationReceivedWei() (*big.Int, error) {
	return _FUManager.Contract.TotalInflationReceivedWei(&_FUManager.CallOpts)
}

// TotalInflationRewardsOfferedWei is a free data retrieval call binding the contract method 0xbd76b69c.
//
// Solidity: function totalInflationRewardsOfferedWei() view returns(uint256)
func (_FUManager *FUManagerCaller) TotalInflationRewardsOfferedWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUManager.contract.Call(opts, &out, "totalInflationRewardsOfferedWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalInflationRewardsOfferedWei is a free data retrieval call binding the contract method 0xbd76b69c.
//
// Solidity: function totalInflationRewardsOfferedWei() view returns(uint256)
func (_FUManager *FUManagerSession) TotalInflationRewardsOfferedWei() (*big.Int, error) {
	return _FUManager.Contract.TotalInflationRewardsOfferedWei(&_FUManager.CallOpts)
}

// TotalInflationRewardsOfferedWei is a free data retrieval call binding the contract method 0xbd76b69c.
//
// Solidity: function totalInflationRewardsOfferedWei() view returns(uint256)
func (_FUManager *FUManagerCallerSession) TotalInflationRewardsOfferedWei() (*big.Int, error) {
	return _FUManager.Contract.TotalInflationRewardsOfferedWei(&_FUManager.CallOpts)
}

// Advance is a paid mutator transaction binding the contract method 0xea105ac7.
//
// Solidity: function advance() returns()
func (_FUManager *FUManagerTransactor) Advance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "advance")
}

// Advance is a paid mutator transaction binding the contract method 0xea105ac7.
//
// Solidity: function advance() returns()
func (_FUManager *FUManagerSession) Advance() (*types.Transaction, error) {
	return _FUManager.Contract.Advance(&_FUManager.TransactOpts)
}

// Advance is a paid mutator transaction binding the contract method 0xea105ac7.
//
// Solidity: function advance() returns()
func (_FUManager *FUManagerTransactorSession) Advance() (*types.Transaction, error) {
	return _FUManager.Contract.Advance(&_FUManager.TransactOpts)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FUManager *FUManagerTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FUManager *FUManagerSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FUManager.Contract.CancelGovernanceCall(&_FUManager.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FUManager *FUManagerTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FUManager.Contract.CancelGovernanceCall(&_FUManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FUManager *FUManagerTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FUManager *FUManagerSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FUManager.Contract.ExecuteGovernanceCall(&_FUManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FUManager *FUManagerTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FUManager.Contract.ExecuteGovernanceCall(&_FUManager.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FUManager *FUManagerTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FUManager *FUManagerSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FUManager.Contract.Initialise(&_FUManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FUManager *FUManagerTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FUManager.Contract.Initialise(&_FUManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// OfferIncentive is a paid mutator transaction binding the contract method 0x36247180.
//
// Solidity: function offerIncentive((uint256,uint256) _offer) payable returns()
func (_FUManager *FUManagerTransactor) OfferIncentive(opts *bind.TransactOpts, _offer IFastUpdateIncentiveManagerIncentiveOffer) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "offerIncentive", _offer)
}

// OfferIncentive is a paid mutator transaction binding the contract method 0x36247180.
//
// Solidity: function offerIncentive((uint256,uint256) _offer) payable returns()
func (_FUManager *FUManagerSession) OfferIncentive(_offer IFastUpdateIncentiveManagerIncentiveOffer) (*types.Transaction, error) {
	return _FUManager.Contract.OfferIncentive(&_FUManager.TransactOpts, _offer)
}

// OfferIncentive is a paid mutator transaction binding the contract method 0x36247180.
//
// Solidity: function offerIncentive((uint256,uint256) _offer) payable returns()
func (_FUManager *FUManagerTransactorSession) OfferIncentive(_offer IFastUpdateIncentiveManagerIncentiveOffer) (*types.Transaction, error) {
	return _FUManager.Contract.OfferIncentive(&_FUManager.TransactOpts, _offer)
}

// ReceiveInflation is a paid mutator transaction binding the contract method 0x06201f1d.
//
// Solidity: function receiveInflation() payable returns()
func (_FUManager *FUManagerTransactor) ReceiveInflation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "receiveInflation")
}

// ReceiveInflation is a paid mutator transaction binding the contract method 0x06201f1d.
//
// Solidity: function receiveInflation() payable returns()
func (_FUManager *FUManagerSession) ReceiveInflation() (*types.Transaction, error) {
	return _FUManager.Contract.ReceiveInflation(&_FUManager.TransactOpts)
}

// ReceiveInflation is a paid mutator transaction binding the contract method 0x06201f1d.
//
// Solidity: function receiveInflation() payable returns()
func (_FUManager *FUManagerTransactorSession) ReceiveInflation() (*types.Transaction, error) {
	return _FUManager.Contract.ReceiveInflation(&_FUManager.TransactOpts)
}

// SetDailyAuthorizedInflation is a paid mutator transaction binding the contract method 0xe2739563.
//
// Solidity: function setDailyAuthorizedInflation(uint256 _toAuthorizeWei) returns()
func (_FUManager *FUManagerTransactor) SetDailyAuthorizedInflation(opts *bind.TransactOpts, _toAuthorizeWei *big.Int) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "setDailyAuthorizedInflation", _toAuthorizeWei)
}

// SetDailyAuthorizedInflation is a paid mutator transaction binding the contract method 0xe2739563.
//
// Solidity: function setDailyAuthorizedInflation(uint256 _toAuthorizeWei) returns()
func (_FUManager *FUManagerSession) SetDailyAuthorizedInflation(_toAuthorizeWei *big.Int) (*types.Transaction, error) {
	return _FUManager.Contract.SetDailyAuthorizedInflation(&_FUManager.TransactOpts, _toAuthorizeWei)
}

// SetDailyAuthorizedInflation is a paid mutator transaction binding the contract method 0xe2739563.
//
// Solidity: function setDailyAuthorizedInflation(uint256 _toAuthorizeWei) returns()
func (_FUManager *FUManagerTransactorSession) SetDailyAuthorizedInflation(_toAuthorizeWei *big.Int) (*types.Transaction, error) {
	return _FUManager.Contract.SetDailyAuthorizedInflation(&_FUManager.TransactOpts, _toAuthorizeWei)
}

// SetIncentiveParameters is a paid mutator transaction binding the contract method 0x75d71307.
//
// Solidity: function setIncentiveParameters(uint256 _ss, uint256 _r, uint256 _x, uint256 _dur) returns()
func (_FUManager *FUManagerTransactor) SetIncentiveParameters(opts *bind.TransactOpts, _ss *big.Int, _r *big.Int, _x *big.Int, _dur *big.Int) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "setIncentiveParameters", _ss, _r, _x, _dur)
}

// SetIncentiveParameters is a paid mutator transaction binding the contract method 0x75d71307.
//
// Solidity: function setIncentiveParameters(uint256 _ss, uint256 _r, uint256 _x, uint256 _dur) returns()
func (_FUManager *FUManagerSession) SetIncentiveParameters(_ss *big.Int, _r *big.Int, _x *big.Int, _dur *big.Int) (*types.Transaction, error) {
	return _FUManager.Contract.SetIncentiveParameters(&_FUManager.TransactOpts, _ss, _r, _x, _dur)
}

// SetIncentiveParameters is a paid mutator transaction binding the contract method 0x75d71307.
//
// Solidity: function setIncentiveParameters(uint256 _ss, uint256 _r, uint256 _x, uint256 _dur) returns()
func (_FUManager *FUManagerTransactorSession) SetIncentiveParameters(_ss *big.Int, _r *big.Int, _x *big.Int, _dur *big.Int) (*types.Transaction, error) {
	return _FUManager.Contract.SetIncentiveParameters(&_FUManager.TransactOpts, _ss, _r, _x, _dur)
}

// SetRangeIncreaseLimit is a paid mutator transaction binding the contract method 0x864578e8.
//
// Solidity: function setRangeIncreaseLimit(uint256 _lim) returns()
func (_FUManager *FUManagerTransactor) SetRangeIncreaseLimit(opts *bind.TransactOpts, _lim *big.Int) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "setRangeIncreaseLimit", _lim)
}

// SetRangeIncreaseLimit is a paid mutator transaction binding the contract method 0x864578e8.
//
// Solidity: function setRangeIncreaseLimit(uint256 _lim) returns()
func (_FUManager *FUManagerSession) SetRangeIncreaseLimit(_lim *big.Int) (*types.Transaction, error) {
	return _FUManager.Contract.SetRangeIncreaseLimit(&_FUManager.TransactOpts, _lim)
}

// SetRangeIncreaseLimit is a paid mutator transaction binding the contract method 0x864578e8.
//
// Solidity: function setRangeIncreaseLimit(uint256 _lim) returns()
func (_FUManager *FUManagerTransactorSession) SetRangeIncreaseLimit(_lim *big.Int) (*types.Transaction, error) {
	return _FUManager.Contract.SetRangeIncreaseLimit(&_FUManager.TransactOpts, _lim)
}

// SetRangeIncreasePrice is a paid mutator transaction binding the contract method 0x0d6e9537.
//
// Solidity: function setRangeIncreasePrice(uint256 _price) returns()
func (_FUManager *FUManagerTransactor) SetRangeIncreasePrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "setRangeIncreasePrice", _price)
}

// SetRangeIncreasePrice is a paid mutator transaction binding the contract method 0x0d6e9537.
//
// Solidity: function setRangeIncreasePrice(uint256 _price) returns()
func (_FUManager *FUManagerSession) SetRangeIncreasePrice(_price *big.Int) (*types.Transaction, error) {
	return _FUManager.Contract.SetRangeIncreasePrice(&_FUManager.TransactOpts, _price)
}

// SetRangeIncreasePrice is a paid mutator transaction binding the contract method 0x0d6e9537.
//
// Solidity: function setRangeIncreasePrice(uint256 _price) returns()
func (_FUManager *FUManagerTransactorSession) SetRangeIncreasePrice(_price *big.Int) (*types.Transaction, error) {
	return _FUManager.Contract.SetRangeIncreasePrice(&_FUManager.TransactOpts, _price)
}

// SetSampleIncreaseLimit is a paid mutator transaction binding the contract method 0xf7690bfe.
//
// Solidity: function setSampleIncreaseLimit(uint256 _lim) returns()
func (_FUManager *FUManagerTransactor) SetSampleIncreaseLimit(opts *bind.TransactOpts, _lim *big.Int) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "setSampleIncreaseLimit", _lim)
}

// SetSampleIncreaseLimit is a paid mutator transaction binding the contract method 0xf7690bfe.
//
// Solidity: function setSampleIncreaseLimit(uint256 _lim) returns()
func (_FUManager *FUManagerSession) SetSampleIncreaseLimit(_lim *big.Int) (*types.Transaction, error) {
	return _FUManager.Contract.SetSampleIncreaseLimit(&_FUManager.TransactOpts, _lim)
}

// SetSampleIncreaseLimit is a paid mutator transaction binding the contract method 0xf7690bfe.
//
// Solidity: function setSampleIncreaseLimit(uint256 _lim) returns()
func (_FUManager *FUManagerTransactorSession) SetSampleIncreaseLimit(_lim *big.Int) (*types.Transaction, error) {
	return _FUManager.Contract.SetSampleIncreaseLimit(&_FUManager.TransactOpts, _lim)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FUManager *FUManagerTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FUManager *FUManagerSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FUManager.Contract.SwitchToProductionMode(&_FUManager.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FUManager *FUManagerTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FUManager.Contract.SwitchToProductionMode(&_FUManager.TransactOpts)
}

// TriggerRewardEpochSwitchover is a paid mutator transaction binding the contract method 0x91f25679.
//
// Solidity: function triggerRewardEpochSwitchover(uint24 _currentRewardEpochId, uint64 _currentRewardEpochExpectedEndTs, uint64 _rewardEpochDurationSeconds) returns()
func (_FUManager *FUManagerTransactor) TriggerRewardEpochSwitchover(opts *bind.TransactOpts, _currentRewardEpochId *big.Int, _currentRewardEpochExpectedEndTs uint64, _rewardEpochDurationSeconds uint64) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "triggerRewardEpochSwitchover", _currentRewardEpochId, _currentRewardEpochExpectedEndTs, _rewardEpochDurationSeconds)
}

// TriggerRewardEpochSwitchover is a paid mutator transaction binding the contract method 0x91f25679.
//
// Solidity: function triggerRewardEpochSwitchover(uint24 _currentRewardEpochId, uint64 _currentRewardEpochExpectedEndTs, uint64 _rewardEpochDurationSeconds) returns()
func (_FUManager *FUManagerSession) TriggerRewardEpochSwitchover(_currentRewardEpochId *big.Int, _currentRewardEpochExpectedEndTs uint64, _rewardEpochDurationSeconds uint64) (*types.Transaction, error) {
	return _FUManager.Contract.TriggerRewardEpochSwitchover(&_FUManager.TransactOpts, _currentRewardEpochId, _currentRewardEpochExpectedEndTs, _rewardEpochDurationSeconds)
}

// TriggerRewardEpochSwitchover is a paid mutator transaction binding the contract method 0x91f25679.
//
// Solidity: function triggerRewardEpochSwitchover(uint24 _currentRewardEpochId, uint64 _currentRewardEpochExpectedEndTs, uint64 _rewardEpochDurationSeconds) returns()
func (_FUManager *FUManagerTransactorSession) TriggerRewardEpochSwitchover(_currentRewardEpochId *big.Int, _currentRewardEpochExpectedEndTs uint64, _rewardEpochDurationSeconds uint64) (*types.Transaction, error) {
	return _FUManager.Contract.TriggerRewardEpochSwitchover(&_FUManager.TransactOpts, _currentRewardEpochId, _currentRewardEpochExpectedEndTs, _rewardEpochDurationSeconds)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FUManager *FUManagerTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FUManager.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FUManager *FUManagerSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FUManager.Contract.UpdateContractAddresses(&_FUManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FUManager *FUManagerTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FUManager.Contract.UpdateContractAddresses(&_FUManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// FUManagerDailyAuthorizedInflationSetIterator is returned from FilterDailyAuthorizedInflationSet and is used to iterate over the raw logs and unpacked data for DailyAuthorizedInflationSet events raised by the FUManager contract.
type FUManagerDailyAuthorizedInflationSetIterator struct {
	Event *FUManagerDailyAuthorizedInflationSet // Event containing the contract specifics and raw log

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
func (it *FUManagerDailyAuthorizedInflationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUManagerDailyAuthorizedInflationSet)
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
		it.Event = new(FUManagerDailyAuthorizedInflationSet)
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
func (it *FUManagerDailyAuthorizedInflationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUManagerDailyAuthorizedInflationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUManagerDailyAuthorizedInflationSet represents a DailyAuthorizedInflationSet event raised by the FUManager contract.
type FUManagerDailyAuthorizedInflationSet struct {
	AuthorizedAmountWei *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDailyAuthorizedInflationSet is a free log retrieval operation binding the contract event 0x187f32a0f765499f15b3bb52ed0aebf6015059f230f2ace7e701e60a47669595.
//
// Solidity: event DailyAuthorizedInflationSet(uint256 authorizedAmountWei)
func (_FUManager *FUManagerFilterer) FilterDailyAuthorizedInflationSet(opts *bind.FilterOpts) (*FUManagerDailyAuthorizedInflationSetIterator, error) {

	logs, sub, err := _FUManager.contract.FilterLogs(opts, "DailyAuthorizedInflationSet")
	if err != nil {
		return nil, err
	}
	return &FUManagerDailyAuthorizedInflationSetIterator{contract: _FUManager.contract, event: "DailyAuthorizedInflationSet", logs: logs, sub: sub}, nil
}

// WatchDailyAuthorizedInflationSet is a free log subscription operation binding the contract event 0x187f32a0f765499f15b3bb52ed0aebf6015059f230f2ace7e701e60a47669595.
//
// Solidity: event DailyAuthorizedInflationSet(uint256 authorizedAmountWei)
func (_FUManager *FUManagerFilterer) WatchDailyAuthorizedInflationSet(opts *bind.WatchOpts, sink chan<- *FUManagerDailyAuthorizedInflationSet) (event.Subscription, error) {

	logs, sub, err := _FUManager.contract.WatchLogs(opts, "DailyAuthorizedInflationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUManagerDailyAuthorizedInflationSet)
				if err := _FUManager.contract.UnpackLog(event, "DailyAuthorizedInflationSet", log); err != nil {
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
func (_FUManager *FUManagerFilterer) ParseDailyAuthorizedInflationSet(log types.Log) (*FUManagerDailyAuthorizedInflationSet, error) {
	event := new(FUManagerDailyAuthorizedInflationSet)
	if err := _FUManager.contract.UnpackLog(event, "DailyAuthorizedInflationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUManagerGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the FUManager contract.
type FUManagerGovernanceCallTimelockedIterator struct {
	Event *FUManagerGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *FUManagerGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUManagerGovernanceCallTimelocked)
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
		it.Event = new(FUManagerGovernanceCallTimelocked)
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
func (it *FUManagerGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUManagerGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FUManager contract.
type FUManagerGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FUManager *FUManagerFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*FUManagerGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _FUManager.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &FUManagerGovernanceCallTimelockedIterator{contract: _FUManager.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FUManager *FUManagerFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *FUManagerGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _FUManager.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUManagerGovernanceCallTimelocked)
				if err := _FUManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_FUManager *FUManagerFilterer) ParseGovernanceCallTimelocked(log types.Log) (*FUManagerGovernanceCallTimelocked, error) {
	event := new(FUManagerGovernanceCallTimelocked)
	if err := _FUManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUManagerGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the FUManager contract.
type FUManagerGovernanceInitialisedIterator struct {
	Event *FUManagerGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *FUManagerGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUManagerGovernanceInitialised)
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
		it.Event = new(FUManagerGovernanceInitialised)
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
func (it *FUManagerGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUManagerGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUManagerGovernanceInitialised represents a GovernanceInitialised event raised by the FUManager contract.
type FUManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FUManager *FUManagerFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*FUManagerGovernanceInitialisedIterator, error) {

	logs, sub, err := _FUManager.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &FUManagerGovernanceInitialisedIterator{contract: _FUManager.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FUManager *FUManagerFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *FUManagerGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _FUManager.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUManagerGovernanceInitialised)
				if err := _FUManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_FUManager *FUManagerFilterer) ParseGovernanceInitialised(log types.Log) (*FUManagerGovernanceInitialised, error) {
	event := new(FUManagerGovernanceInitialised)
	if err := _FUManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUManagerGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the FUManager contract.
type FUManagerGovernedProductionModeEnteredIterator struct {
	Event *FUManagerGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *FUManagerGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUManagerGovernedProductionModeEntered)
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
		it.Event = new(FUManagerGovernedProductionModeEntered)
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
func (it *FUManagerGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUManagerGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FUManager contract.
type FUManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FUManager *FUManagerFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*FUManagerGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _FUManager.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &FUManagerGovernedProductionModeEnteredIterator{contract: _FUManager.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FUManager *FUManagerFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *FUManagerGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _FUManager.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUManagerGovernedProductionModeEntered)
				if err := _FUManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_FUManager *FUManagerFilterer) ParseGovernedProductionModeEntered(log types.Log) (*FUManagerGovernedProductionModeEntered, error) {
	event := new(FUManagerGovernedProductionModeEntered)
	if err := _FUManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUManagerIncentiveOfferedIterator is returned from FilterIncentiveOffered and is used to iterate over the raw logs and unpacked data for IncentiveOffered events raised by the FUManager contract.
type FUManagerIncentiveOfferedIterator struct {
	Event *FUManagerIncentiveOffered // Event containing the contract specifics and raw log

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
func (it *FUManagerIncentiveOfferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUManagerIncentiveOffered)
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
		it.Event = new(FUManagerIncentiveOffered)
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
func (it *FUManagerIncentiveOfferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUManagerIncentiveOfferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUManagerIncentiveOffered represents a IncentiveOffered event raised by the FUManager contract.
type FUManagerIncentiveOffered struct {
	RewardEpochId      *big.Int
	RangeIncrease      *big.Int
	SampleSizeIncrease *big.Int
	OfferAmount        *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterIncentiveOffered is a free log retrieval operation binding the contract event 0x1c5543607841f3a87aa841c3bfa973bf64f4d545b1d9c12af3cd5831ecf82603.
//
// Solidity: event IncentiveOffered(uint24 indexed rewardEpochId, uint256 rangeIncrease, uint256 sampleSizeIncrease, uint256 offerAmount)
func (_FUManager *FUManagerFilterer) FilterIncentiveOffered(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*FUManagerIncentiveOfferedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FUManager.contract.FilterLogs(opts, "IncentiveOffered", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &FUManagerIncentiveOfferedIterator{contract: _FUManager.contract, event: "IncentiveOffered", logs: logs, sub: sub}, nil
}

// WatchIncentiveOffered is a free log subscription operation binding the contract event 0x1c5543607841f3a87aa841c3bfa973bf64f4d545b1d9c12af3cd5831ecf82603.
//
// Solidity: event IncentiveOffered(uint24 indexed rewardEpochId, uint256 rangeIncrease, uint256 sampleSizeIncrease, uint256 offerAmount)
func (_FUManager *FUManagerFilterer) WatchIncentiveOffered(opts *bind.WatchOpts, sink chan<- *FUManagerIncentiveOffered, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FUManager.contract.WatchLogs(opts, "IncentiveOffered", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUManagerIncentiveOffered)
				if err := _FUManager.contract.UnpackLog(event, "IncentiveOffered", log); err != nil {
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

// ParseIncentiveOffered is a log parse operation binding the contract event 0x1c5543607841f3a87aa841c3bfa973bf64f4d545b1d9c12af3cd5831ecf82603.
//
// Solidity: event IncentiveOffered(uint24 indexed rewardEpochId, uint256 rangeIncrease, uint256 sampleSizeIncrease, uint256 offerAmount)
func (_FUManager *FUManagerFilterer) ParseIncentiveOffered(log types.Log) (*FUManagerIncentiveOffered, error) {
	event := new(FUManagerIncentiveOffered)
	if err := _FUManager.contract.UnpackLog(event, "IncentiveOffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUManagerInflationReceivedIterator is returned from FilterInflationReceived and is used to iterate over the raw logs and unpacked data for InflationReceived events raised by the FUManager contract.
type FUManagerInflationReceivedIterator struct {
	Event *FUManagerInflationReceived // Event containing the contract specifics and raw log

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
func (it *FUManagerInflationReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUManagerInflationReceived)
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
		it.Event = new(FUManagerInflationReceived)
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
func (it *FUManagerInflationReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUManagerInflationReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUManagerInflationReceived represents a InflationReceived event raised by the FUManager contract.
type FUManagerInflationReceived struct {
	AmountReceivedWei *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInflationReceived is a free log retrieval operation binding the contract event 0x95c4e29cc99bc027cfc3cd719d6fd973d5f0317061885fbb322b9b17d8d35d37.
//
// Solidity: event InflationReceived(uint256 amountReceivedWei)
func (_FUManager *FUManagerFilterer) FilterInflationReceived(opts *bind.FilterOpts) (*FUManagerInflationReceivedIterator, error) {

	logs, sub, err := _FUManager.contract.FilterLogs(opts, "InflationReceived")
	if err != nil {
		return nil, err
	}
	return &FUManagerInflationReceivedIterator{contract: _FUManager.contract, event: "InflationReceived", logs: logs, sub: sub}, nil
}

// WatchInflationReceived is a free log subscription operation binding the contract event 0x95c4e29cc99bc027cfc3cd719d6fd973d5f0317061885fbb322b9b17d8d35d37.
//
// Solidity: event InflationReceived(uint256 amountReceivedWei)
func (_FUManager *FUManagerFilterer) WatchInflationReceived(opts *bind.WatchOpts, sink chan<- *FUManagerInflationReceived) (event.Subscription, error) {

	logs, sub, err := _FUManager.contract.WatchLogs(opts, "InflationReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUManagerInflationReceived)
				if err := _FUManager.contract.UnpackLog(event, "InflationReceived", log); err != nil {
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
func (_FUManager *FUManagerFilterer) ParseInflationReceived(log types.Log) (*FUManagerInflationReceived, error) {
	event := new(FUManagerInflationReceived)
	if err := _FUManager.contract.UnpackLog(event, "InflationReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUManagerInflationRewardsOfferedIterator is returned from FilterInflationRewardsOffered and is used to iterate over the raw logs and unpacked data for InflationRewardsOffered events raised by the FUManager contract.
type FUManagerInflationRewardsOfferedIterator struct {
	Event *FUManagerInflationRewardsOffered // Event containing the contract specifics and raw log

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
func (it *FUManagerInflationRewardsOfferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUManagerInflationRewardsOffered)
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
		it.Event = new(FUManagerInflationRewardsOffered)
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
func (it *FUManagerInflationRewardsOfferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUManagerInflationRewardsOfferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUManagerInflationRewardsOffered represents a InflationRewardsOffered event raised by the FUManager contract.
type FUManagerInflationRewardsOffered struct {
	RewardEpochId      *big.Int
	FeedConfigurations []IFastUpdatesConfigurationFeedConfiguration
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInflationRewardsOffered is a free log retrieval operation binding the contract event 0x58575ff9908663af0451165c3cefcb802da242d63261f6d9df3be0e05366e4da.
//
// Solidity: event InflationRewardsOffered(uint24 indexed rewardEpochId, (bytes21,uint32,uint24)[] feedConfigurations, uint256 amount)
func (_FUManager *FUManagerFilterer) FilterInflationRewardsOffered(opts *bind.FilterOpts, rewardEpochId []*big.Int) (*FUManagerInflationRewardsOfferedIterator, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FUManager.contract.FilterLogs(opts, "InflationRewardsOffered", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &FUManagerInflationRewardsOfferedIterator{contract: _FUManager.contract, event: "InflationRewardsOffered", logs: logs, sub: sub}, nil
}

// WatchInflationRewardsOffered is a free log subscription operation binding the contract event 0x58575ff9908663af0451165c3cefcb802da242d63261f6d9df3be0e05366e4da.
//
// Solidity: event InflationRewardsOffered(uint24 indexed rewardEpochId, (bytes21,uint32,uint24)[] feedConfigurations, uint256 amount)
func (_FUManager *FUManagerFilterer) WatchInflationRewardsOffered(opts *bind.WatchOpts, sink chan<- *FUManagerInflationRewardsOffered, rewardEpochId []*big.Int) (event.Subscription, error) {

	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _FUManager.contract.WatchLogs(opts, "InflationRewardsOffered", rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUManagerInflationRewardsOffered)
				if err := _FUManager.contract.UnpackLog(event, "InflationRewardsOffered", log); err != nil {
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

// ParseInflationRewardsOffered is a log parse operation binding the contract event 0x58575ff9908663af0451165c3cefcb802da242d63261f6d9df3be0e05366e4da.
//
// Solidity: event InflationRewardsOffered(uint24 indexed rewardEpochId, (bytes21,uint32,uint24)[] feedConfigurations, uint256 amount)
func (_FUManager *FUManagerFilterer) ParseInflationRewardsOffered(log types.Log) (*FUManagerInflationRewardsOffered, error) {
	event := new(FUManagerInflationRewardsOffered)
	if err := _FUManager.contract.UnpackLog(event, "InflationRewardsOffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUManagerTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the FUManager contract.
type FUManagerTimelockedGovernanceCallCanceledIterator struct {
	Event *FUManagerTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *FUManagerTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUManagerTimelockedGovernanceCallCanceled)
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
		it.Event = new(FUManagerTimelockedGovernanceCallCanceled)
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
func (it *FUManagerTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUManagerTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FUManager contract.
type FUManagerTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FUManager *FUManagerFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*FUManagerTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _FUManager.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &FUManagerTimelockedGovernanceCallCanceledIterator{contract: _FUManager.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FUManager *FUManagerFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *FUManagerTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _FUManager.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUManagerTimelockedGovernanceCallCanceled)
				if err := _FUManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_FUManager *FUManagerFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*FUManagerTimelockedGovernanceCallCanceled, error) {
	event := new(FUManagerTimelockedGovernanceCallCanceled)
	if err := _FUManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUManagerTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the FUManager contract.
type FUManagerTimelockedGovernanceCallExecutedIterator struct {
	Event *FUManagerTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *FUManagerTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUManagerTimelockedGovernanceCallExecuted)
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
		it.Event = new(FUManagerTimelockedGovernanceCallExecuted)
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
func (it *FUManagerTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUManagerTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FUManager contract.
type FUManagerTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FUManager *FUManagerFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*FUManagerTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _FUManager.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &FUManagerTimelockedGovernanceCallExecutedIterator{contract: _FUManager.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FUManager *FUManagerFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *FUManagerTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _FUManager.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUManagerTimelockedGovernanceCallExecuted)
				if err := _FUManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_FUManager *FUManagerFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*FUManagerTimelockedGovernanceCallExecuted, error) {
	event := new(FUManagerTimelockedGovernanceCallExecuted)
	if err := _FUManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
