// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fumanager

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
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
var FUManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"SampleSize\",\"name\":\"_ss\",\"type\":\"uint256\"},{\"internalType\":\"Range\",\"name\":\"_r\",\"type\":\"uint256\"},{\"internalType\":\"SampleSize\",\"name\":\"_sil\",\"type\":\"uint256\"},{\"internalType\":\"Range\",\"name\":\"_ril\",\"type\":\"uint256\"},{\"internalType\":\"Fee\",\"name\":\"_x\",\"type\":\"uint256\"},{\"internalType\":\"Fee\",\"name\":\"_rip\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dur\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"authorizedAmountWei\",\"type\":\"uint256\"}],\"name\":\"DailyAuthorizedInflationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"Range\",\"name\":\"rangeIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"SampleSize\",\"name\":\"sampleSizeIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"Fee\",\"name\":\"offerAmount\",\"type\":\"uint256\"}],\"name\":\"IncentiveOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedWei\",\"type\":\"uint256\"}],\"name\":\"InflationReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"bytes21\",\"name\":\"feedId\",\"type\":\"bytes21\"},{\"internalType\":\"uint32\",\"name\":\"rewardBandValue\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"inflationShare\",\"type\":\"uint24\"}],\"indexed\":false,\"internalType\":\"structIFastUpdatesConfiguration.FeedConfiguration[]\",\"name\":\"feedConfigurations\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InflationRewardsOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"advance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyAuthorizedInflation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fastUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fastUpdatesConfiguration\",\"outputs\":[{\"internalType\":\"contractIFastUpdatesConfiguration\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseScale\",\"outputs\":[{\"internalType\":\"Scale\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentSampleSizeIncreasePrice\",\"outputs\":[{\"internalType\":\"Fee\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedSampleSize\",\"outputs\":[{\"internalType\":\"SampleSize\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIncentiveDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInflationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPrecision\",\"outputs\":[{\"internalType\":\"Precision\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRange\",\"outputs\":[{\"internalType\":\"Range\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getScale\",\"outputs\":[{\"internalType\":\"Scale\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenPoolSupplyData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockedFundsWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalInflationAuthorizedWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalClaimedWei\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInflationAuthorizationReceivedTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInflationReceivedTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Range\",\"name\":\"rangeIncrease\",\"type\":\"uint256\"},{\"internalType\":\"Range\",\"name\":\"rangeLimit\",\"type\":\"uint256\"}],\"internalType\":\"structIFastUpdateIncentiveManager.IncentiveOffer\",\"name\":\"_offer\",\"type\":\"tuple\"}],\"name\":\"offerIncentive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rangeIncreaseLimit\",\"outputs\":[{\"internalType\":\"Range\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rangeIncreasePrice\",\"outputs\":[{\"internalType\":\"Fee\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveInflation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sampleIncreaseLimit\",\"outputs\":[{\"internalType\":\"SampleSize\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_toAuthorizeWei\",\"type\":\"uint256\"}],\"name\":\"setDailyAuthorizedInflation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"SampleSize\",\"name\":\"_ss\",\"type\":\"uint256\"},{\"internalType\":\"Range\",\"name\":\"_r\",\"type\":\"uint256\"},{\"internalType\":\"Fee\",\"name\":\"_x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dur\",\"type\":\"uint256\"}],\"name\":\"setIncentiveParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Range\",\"name\":\"_lim\",\"type\":\"uint256\"}],\"name\":\"setRangeIncreaseLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Fee\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setRangeIncreasePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"SampleSize\",\"name\":\"_lim\",\"type\":\"uint256\"}],\"name\":\"setSampleIncreaseLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationAuthorizedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationReceivedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationRewardsOfferedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_currentRewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_currentRewardEpochExpectedEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardEpochDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"triggerRewardEpochSwitchover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "FUManager",
}

// FUManager is an auto generated Go binding around an Ethereum contract.
type FUManager struct {
	abi abi.ABI
}

// NewFUManager creates a new instance of FUManager.
func NewFUManager() *FUManager {
	parsed, err := FUManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &FUManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *FUManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint256 _ss, uint256 _r, uint256 _sil, uint256 _ril, uint256 _x, uint256 _rip, uint256 _dur) returns()
func (fUManager *FUManager) PackConstructor(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _ss *big.Int, _r *big.Int, _sil *big.Int, _ril *big.Int, _x *big.Int, _rip *big.Int, _dur *big.Int) []byte {
	enc, err := fUManager.abi.Pack("", _governanceSettings, _initialGovernance, _addressUpdater, _ss, _r, _sil, _ril, _x, _rip, _dur)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackAdvance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xea105ac7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function advance() returns()
func (fUManager *FUManager) PackAdvance() []byte {
	enc, err := fUManager.abi.Pack("advance")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAdvance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xea105ac7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function advance() returns()
func (fUManager *FUManager) TryPackAdvance() ([]byte, error) {
	return fUManager.abi.Pack("advance")
}

// PackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67fc4029.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (fUManager *FUManager) PackCancelGovernanceCall(selector [4]byte) []byte {
	enc, err := fUManager.abi.Pack("cancelGovernanceCall", selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67fc4029.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (fUManager *FUManager) TryPackCancelGovernanceCall(selector [4]byte) ([]byte, error) {
	return fUManager.abi.Pack("cancelGovernanceCall", selector)
}

// PackDailyAuthorizedInflation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x708e34ce.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (fUManager *FUManager) PackDailyAuthorizedInflation() []byte {
	enc, err := fUManager.abi.Pack("dailyAuthorizedInflation")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDailyAuthorizedInflation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x708e34ce.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (fUManager *FUManager) TryPackDailyAuthorizedInflation() ([]byte, error) {
	return fUManager.abi.Pack("dailyAuthorizedInflation")
}

// UnpackDailyAuthorizedInflation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x708e34ce.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (fUManager *FUManager) UnpackDailyAuthorizedInflation(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("dailyAuthorizedInflation", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ff27079.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (fUManager *FUManager) PackExecuteGovernanceCall(selector [4]byte) []byte {
	enc, err := fUManager.abi.Pack("executeGovernanceCall", selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ff27079.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (fUManager *FUManager) TryPackExecuteGovernanceCall(selector [4]byte) ([]byte, error) {
	return fUManager.abi.Pack("executeGovernanceCall", selector)
}

// PackFastUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd29a4fa9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fastUpdater() view returns(address)
func (fUManager *FUManager) PackFastUpdater() []byte {
	enc, err := fUManager.abi.Pack("fastUpdater")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFastUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd29a4fa9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fastUpdater() view returns(address)
func (fUManager *FUManager) TryPackFastUpdater() ([]byte, error) {
	return fUManager.abi.Pack("fastUpdater")
}

// UnpackFastUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd29a4fa9.
//
// Solidity: function fastUpdater() view returns(address)
func (fUManager *FUManager) UnpackFastUpdater(data []byte) (common.Address, error) {
	out, err := fUManager.abi.Unpack("fastUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFastUpdatesConfiguration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc10f489a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (fUManager *FUManager) PackFastUpdatesConfiguration() []byte {
	enc, err := fUManager.abi.Pack("fastUpdatesConfiguration")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFastUpdatesConfiguration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc10f489a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (fUManager *FUManager) TryPackFastUpdatesConfiguration() ([]byte, error) {
	return fUManager.abi.Pack("fastUpdatesConfiguration")
}

// UnpackFastUpdatesConfiguration is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc10f489a.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (fUManager *FUManager) UnpackFastUpdatesConfiguration(data []byte) (common.Address, error) {
	out, err := fUManager.abi.Unpack("fastUpdatesConfiguration", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFlareSystemsManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaae7fc9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareSystemsManager() view returns(address)
func (fUManager *FUManager) PackFlareSystemsManager() []byte {
	enc, err := fUManager.abi.Pack("flareSystemsManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFlareSystemsManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaae7fc9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function flareSystemsManager() view returns(address)
func (fUManager *FUManager) TryPackFlareSystemsManager() ([]byte, error) {
	return fUManager.abi.Pack("flareSystemsManager")
}

// UnpackFlareSystemsManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (fUManager *FUManager) UnpackFlareSystemsManager(data []byte) (common.Address, error) {
	out, err := fUManager.abi.Unpack("flareSystemsManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (fUManager *FUManager) PackGetAddressUpdater() []byte {
	enc, err := fUManager.abi.Pack("getAddressUpdater")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (fUManager *FUManager) TryPackGetAddressUpdater() ([]byte, error) {
	return fUManager.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (fUManager *FUManager) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := fUManager.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetBaseScale is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a68533f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBaseScale() view returns(uint256)
func (fUManager *FUManager) PackGetBaseScale() []byte {
	enc, err := fUManager.abi.Pack("getBaseScale")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBaseScale is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a68533f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBaseScale() view returns(uint256)
func (fUManager *FUManager) TryPackGetBaseScale() ([]byte, error) {
	return fUManager.abi.Pack("getBaseScale")
}

// UnpackGetBaseScale is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7a68533f.
//
// Solidity: function getBaseScale() view returns(uint256)
func (fUManager *FUManager) UnpackGetBaseScale(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("getBaseScale", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetContractName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5f5ba72.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getContractName() pure returns(string)
func (fUManager *FUManager) PackGetContractName() []byte {
	enc, err := fUManager.abi.Pack("getContractName")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetContractName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5f5ba72.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getContractName() pure returns(string)
func (fUManager *FUManager) TryPackGetContractName() ([]byte, error) {
	return fUManager.abi.Pack("getContractName")
}

// UnpackGetContractName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (fUManager *FUManager) UnpackGetContractName(data []byte) (string, error) {
	out, err := fUManager.abi.Unpack("getContractName", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackGetCurrentSampleSizeIncreasePrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2de490c3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCurrentSampleSizeIncreasePrice() view returns(uint256)
func (fUManager *FUManager) PackGetCurrentSampleSizeIncreasePrice() []byte {
	enc, err := fUManager.abi.Pack("getCurrentSampleSizeIncreasePrice")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCurrentSampleSizeIncreasePrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2de490c3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCurrentSampleSizeIncreasePrice() view returns(uint256)
func (fUManager *FUManager) TryPackGetCurrentSampleSizeIncreasePrice() ([]byte, error) {
	return fUManager.abi.Pack("getCurrentSampleSizeIncreasePrice")
}

// UnpackGetCurrentSampleSizeIncreasePrice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2de490c3.
//
// Solidity: function getCurrentSampleSizeIncreasePrice() view returns(uint256)
func (fUManager *FUManager) UnpackGetCurrentSampleSizeIncreasePrice(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("getCurrentSampleSizeIncreasePrice", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetExpectedBalance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf04cd3b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (fUManager *FUManager) PackGetExpectedBalance() []byte {
	enc, err := fUManager.abi.Pack("getExpectedBalance")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetExpectedBalance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf04cd3b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (fUManager *FUManager) TryPackGetExpectedBalance() ([]byte, error) {
	return fUManager.abi.Pack("getExpectedBalance")
}

// UnpackGetExpectedBalance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaf04cd3b.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (fUManager *FUManager) UnpackGetExpectedBalance(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("getExpectedBalance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetExpectedSampleSize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6d62b413.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getExpectedSampleSize() view returns(uint256)
func (fUManager *FUManager) PackGetExpectedSampleSize() []byte {
	enc, err := fUManager.abi.Pack("getExpectedSampleSize")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetExpectedSampleSize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6d62b413.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getExpectedSampleSize() view returns(uint256)
func (fUManager *FUManager) TryPackGetExpectedSampleSize() ([]byte, error) {
	return fUManager.abi.Pack("getExpectedSampleSize")
}

// UnpackGetExpectedSampleSize is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6d62b413.
//
// Solidity: function getExpectedSampleSize() view returns(uint256)
func (fUManager *FUManager) UnpackGetExpectedSampleSize(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("getExpectedSampleSize", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetIncentiveDuration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd8dca9f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getIncentiveDuration() view returns(uint256)
func (fUManager *FUManager) PackGetIncentiveDuration() []byte {
	enc, err := fUManager.abi.Pack("getIncentiveDuration")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetIncentiveDuration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdd8dca9f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getIncentiveDuration() view returns(uint256)
func (fUManager *FUManager) TryPackGetIncentiveDuration() ([]byte, error) {
	return fUManager.abi.Pack("getIncentiveDuration")
}

// UnpackGetIncentiveDuration is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdd8dca9f.
//
// Solidity: function getIncentiveDuration() view returns(uint256)
func (fUManager *FUManager) UnpackGetIncentiveDuration(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("getIncentiveDuration", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetInflationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed39d3f8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getInflationAddress() view returns(address)
func (fUManager *FUManager) PackGetInflationAddress() []byte {
	enc, err := fUManager.abi.Pack("getInflationAddress")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetInflationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed39d3f8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getInflationAddress() view returns(address)
func (fUManager *FUManager) TryPackGetInflationAddress() ([]byte, error) {
	return fUManager.abi.Pack("getInflationAddress")
}

// UnpackGetInflationAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xed39d3f8.
//
// Solidity: function getInflationAddress() view returns(address)
func (fUManager *FUManager) UnpackGetInflationAddress(data []byte) (common.Address, error) {
	out, err := fUManager.abi.Unpack("getInflationAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetPrecision is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9670c0bc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPrecision() view returns(uint256)
func (fUManager *FUManager) PackGetPrecision() []byte {
	enc, err := fUManager.abi.Pack("getPrecision")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPrecision is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9670c0bc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPrecision() view returns(uint256)
func (fUManager *FUManager) TryPackGetPrecision() ([]byte, error) {
	return fUManager.abi.Pack("getPrecision")
}

// UnpackGetPrecision is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9670c0bc.
//
// Solidity: function getPrecision() view returns(uint256)
func (fUManager *FUManager) UnpackGetPrecision(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("getPrecision", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetRange is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b85961f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRange() view returns(uint256)
func (fUManager *FUManager) PackGetRange() []byte {
	enc, err := fUManager.abi.Pack("getRange")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRange is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b85961f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRange() view returns(uint256)
func (fUManager *FUManager) TryPackGetRange() ([]byte, error) {
	return fUManager.abi.Pack("getRange")
}

// UnpackGetRange is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9b85961f.
//
// Solidity: function getRange() view returns(uint256)
func (fUManager *FUManager) UnpackGetRange(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("getRange", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetScale is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb5cddab8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getScale() view returns(uint256)
func (fUManager *FUManager) PackGetScale() []byte {
	enc, err := fUManager.abi.Pack("getScale")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetScale is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb5cddab8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getScale() view returns(uint256)
func (fUManager *FUManager) TryPackGetScale() ([]byte, error) {
	return fUManager.abi.Pack("getScale")
}

// UnpackGetScale is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb5cddab8.
//
// Solidity: function getScale() view returns(uint256)
func (fUManager *FUManager) UnpackGetScale(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("getScale", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetTokenPoolSupplyData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2dafdbbf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTokenPoolSupplyData() view returns(uint256 _lockedFundsWei, uint256 _totalInflationAuthorizedWei, uint256 _totalClaimedWei)
func (fUManager *FUManager) PackGetTokenPoolSupplyData() []byte {
	enc, err := fUManager.abi.Pack("getTokenPoolSupplyData")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTokenPoolSupplyData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2dafdbbf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTokenPoolSupplyData() view returns(uint256 _lockedFundsWei, uint256 _totalInflationAuthorizedWei, uint256 _totalClaimedWei)
func (fUManager *FUManager) TryPackGetTokenPoolSupplyData() ([]byte, error) {
	return fUManager.abi.Pack("getTokenPoolSupplyData")
}

// GetTokenPoolSupplyDataOutput serves as a container for the return parameters of contract
// method GetTokenPoolSupplyData.
type GetTokenPoolSupplyDataOutput struct {
	LockedFundsWei              *big.Int
	TotalInflationAuthorizedWei *big.Int
	TotalClaimedWei             *big.Int
}

// UnpackGetTokenPoolSupplyData is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2dafdbbf.
//
// Solidity: function getTokenPoolSupplyData() view returns(uint256 _lockedFundsWei, uint256 _totalInflationAuthorizedWei, uint256 _totalClaimedWei)
func (fUManager *FUManager) UnpackGetTokenPoolSupplyData(data []byte) (GetTokenPoolSupplyDataOutput, error) {
	out, err := fUManager.abi.Unpack("getTokenPoolSupplyData", data)
	outstruct := new(GetTokenPoolSupplyDataOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.LockedFundsWei = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.TotalInflationAuthorizedWei = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	outstruct.TotalClaimedWei = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governance() view returns(address)
func (fUManager *FUManager) PackGovernance() []byte {
	enc, err := fUManager.abi.Pack("governance")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function governance() view returns(address)
func (fUManager *FUManager) TryPackGovernance() ([]byte, error) {
	return fUManager.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (fUManager *FUManager) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := fUManager.abi.Unpack("governance", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGovernanceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62354e03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governanceSettings() view returns(address)
func (fUManager *FUManager) PackGovernanceSettings() []byte {
	enc, err := fUManager.abi.Pack("governanceSettings")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGovernanceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62354e03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function governanceSettings() view returns(address)
func (fUManager *FUManager) TryPackGovernanceSettings() ([]byte, error) {
	return fUManager.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (fUManager *FUManager) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := fUManager.abi.Unpack("governanceSettings", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitialise is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef88bf13.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (fUManager *FUManager) PackInitialise(governanceSettings common.Address, initialGovernance common.Address) []byte {
	enc, err := fUManager.abi.Pack("initialise", governanceSettings, initialGovernance)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialise is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef88bf13.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (fUManager *FUManager) TryPackInitialise(governanceSettings common.Address, initialGovernance common.Address) ([]byte, error) {
	return fUManager.abi.Pack("initialise", governanceSettings, initialGovernance)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (fUManager *FUManager) PackIsExecutor(address common.Address) []byte {
	enc, err := fUManager.abi.Pack("isExecutor", address)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (fUManager *FUManager) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return fUManager.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (fUManager *FUManager) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := fUManager.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackLastInflationAuthorizationReceivedTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x473252c4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function lastInflationAuthorizationReceivedTs() view returns(uint256)
func (fUManager *FUManager) PackLastInflationAuthorizationReceivedTs() []byte {
	enc, err := fUManager.abi.Pack("lastInflationAuthorizationReceivedTs")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLastInflationAuthorizationReceivedTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x473252c4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function lastInflationAuthorizationReceivedTs() view returns(uint256)
func (fUManager *FUManager) TryPackLastInflationAuthorizationReceivedTs() ([]byte, error) {
	return fUManager.abi.Pack("lastInflationAuthorizationReceivedTs")
}

// UnpackLastInflationAuthorizationReceivedTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x473252c4.
//
// Solidity: function lastInflationAuthorizationReceivedTs() view returns(uint256)
func (fUManager *FUManager) UnpackLastInflationAuthorizationReceivedTs(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("lastInflationAuthorizationReceivedTs", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackLastInflationReceivedTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x12afcf0b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function lastInflationReceivedTs() view returns(uint256)
func (fUManager *FUManager) PackLastInflationReceivedTs() []byte {
	enc, err := fUManager.abi.Pack("lastInflationReceivedTs")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLastInflationReceivedTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x12afcf0b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function lastInflationReceivedTs() view returns(uint256)
func (fUManager *FUManager) TryPackLastInflationReceivedTs() ([]byte, error) {
	return fUManager.abi.Pack("lastInflationReceivedTs")
}

// UnpackLastInflationReceivedTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x12afcf0b.
//
// Solidity: function lastInflationReceivedTs() view returns(uint256)
func (fUManager *FUManager) UnpackLastInflationReceivedTs(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("lastInflationReceivedTs", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOfferIncentive is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36247180.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function offerIncentive((uint256,uint256) _offer) payable returns()
func (fUManager *FUManager) PackOfferIncentive(offer IFastUpdateIncentiveManagerIncentiveOffer) []byte {
	enc, err := fUManager.abi.Pack("offerIncentive", offer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOfferIncentive is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36247180.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function offerIncentive((uint256,uint256) _offer) payable returns()
func (fUManager *FUManager) TryPackOfferIncentive(offer IFastUpdateIncentiveManagerIncentiveOffer) ([]byte, error) {
	return fUManager.abi.Pack("offerIncentive", offer)
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (fUManager *FUManager) PackProductionMode() []byte {
	enc, err := fUManager.abi.Pack("productionMode")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function productionMode() view returns(bool)
func (fUManager *FUManager) TryPackProductionMode() ([]byte, error) {
	return fUManager.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (fUManager *FUManager) UnpackProductionMode(data []byte) (bool, error) {
	out, err := fUManager.abi.Unpack("productionMode", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRangeIncreaseLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74f3eff9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rangeIncreaseLimit() view returns(uint256)
func (fUManager *FUManager) PackRangeIncreaseLimit() []byte {
	enc, err := fUManager.abi.Pack("rangeIncreaseLimit")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRangeIncreaseLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74f3eff9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rangeIncreaseLimit() view returns(uint256)
func (fUManager *FUManager) TryPackRangeIncreaseLimit() ([]byte, error) {
	return fUManager.abi.Pack("rangeIncreaseLimit")
}

// UnpackRangeIncreaseLimit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x74f3eff9.
//
// Solidity: function rangeIncreaseLimit() view returns(uint256)
func (fUManager *FUManager) UnpackRangeIncreaseLimit(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("rangeIncreaseLimit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRangeIncreasePrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52545a7c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rangeIncreasePrice() view returns(uint256)
func (fUManager *FUManager) PackRangeIncreasePrice() []byte {
	enc, err := fUManager.abi.Pack("rangeIncreasePrice")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRangeIncreasePrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52545a7c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rangeIncreasePrice() view returns(uint256)
func (fUManager *FUManager) TryPackRangeIncreasePrice() ([]byte, error) {
	return fUManager.abi.Pack("rangeIncreasePrice")
}

// UnpackRangeIncreasePrice is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52545a7c.
//
// Solidity: function rangeIncreasePrice() view returns(uint256)
func (fUManager *FUManager) UnpackRangeIncreasePrice(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("rangeIncreasePrice", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackReceiveInflation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06201f1d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function receiveInflation() payable returns()
func (fUManager *FUManager) PackReceiveInflation() []byte {
	enc, err := fUManager.abi.Pack("receiveInflation")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReceiveInflation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06201f1d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function receiveInflation() payable returns()
func (fUManager *FUManager) TryPackReceiveInflation() ([]byte, error) {
	return fUManager.abi.Pack("receiveInflation")
}

// PackRewardManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f4ef8a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardManager() view returns(address)
func (fUManager *FUManager) PackRewardManager() []byte {
	enc, err := fUManager.abi.Pack("rewardManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRewardManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f4ef8a6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rewardManager() view returns(address)
func (fUManager *FUManager) TryPackRewardManager() ([]byte, error) {
	return fUManager.abi.Pack("rewardManager")
}

// UnpackRewardManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (fUManager *FUManager) UnpackRewardManager(data []byte) (common.Address, error) {
	out, err := fUManager.abi.Unpack("rewardManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSampleIncreaseLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd4ab8f94.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sampleIncreaseLimit() view returns(uint256)
func (fUManager *FUManager) PackSampleIncreaseLimit() []byte {
	enc, err := fUManager.abi.Pack("sampleIncreaseLimit")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSampleIncreaseLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd4ab8f94.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sampleIncreaseLimit() view returns(uint256)
func (fUManager *FUManager) TryPackSampleIncreaseLimit() ([]byte, error) {
	return fUManager.abi.Pack("sampleIncreaseLimit")
}

// UnpackSampleIncreaseLimit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd4ab8f94.
//
// Solidity: function sampleIncreaseLimit() view returns(uint256)
func (fUManager *FUManager) UnpackSampleIncreaseLimit(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("sampleIncreaseLimit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSetDailyAuthorizedInflation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2739563.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setDailyAuthorizedInflation(uint256 _toAuthorizeWei) returns()
func (fUManager *FUManager) PackSetDailyAuthorizedInflation(toAuthorizeWei *big.Int) []byte {
	enc, err := fUManager.abi.Pack("setDailyAuthorizedInflation", toAuthorizeWei)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetDailyAuthorizedInflation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2739563.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setDailyAuthorizedInflation(uint256 _toAuthorizeWei) returns()
func (fUManager *FUManager) TryPackSetDailyAuthorizedInflation(toAuthorizeWei *big.Int) ([]byte, error) {
	return fUManager.abi.Pack("setDailyAuthorizedInflation", toAuthorizeWei)
}

// PackSetIncentiveParameters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x75d71307.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setIncentiveParameters(uint256 _ss, uint256 _r, uint256 _x, uint256 _dur) returns()
func (fUManager *FUManager) PackSetIncentiveParameters(ss *big.Int, r *big.Int, x *big.Int, dur *big.Int) []byte {
	enc, err := fUManager.abi.Pack("setIncentiveParameters", ss, r, x, dur)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetIncentiveParameters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x75d71307.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setIncentiveParameters(uint256 _ss, uint256 _r, uint256 _x, uint256 _dur) returns()
func (fUManager *FUManager) TryPackSetIncentiveParameters(ss *big.Int, r *big.Int, x *big.Int, dur *big.Int) ([]byte, error) {
	return fUManager.abi.Pack("setIncentiveParameters", ss, r, x, dur)
}

// PackSetRangeIncreaseLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x864578e8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setRangeIncreaseLimit(uint256 _lim) returns()
func (fUManager *FUManager) PackSetRangeIncreaseLimit(lim *big.Int) []byte {
	enc, err := fUManager.abi.Pack("setRangeIncreaseLimit", lim)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetRangeIncreaseLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x864578e8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setRangeIncreaseLimit(uint256 _lim) returns()
func (fUManager *FUManager) TryPackSetRangeIncreaseLimit(lim *big.Int) ([]byte, error) {
	return fUManager.abi.Pack("setRangeIncreaseLimit", lim)
}

// PackSetRangeIncreasePrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d6e9537.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setRangeIncreasePrice(uint256 _price) returns()
func (fUManager *FUManager) PackSetRangeIncreasePrice(price *big.Int) []byte {
	enc, err := fUManager.abi.Pack("setRangeIncreasePrice", price)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetRangeIncreasePrice is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d6e9537.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setRangeIncreasePrice(uint256 _price) returns()
func (fUManager *FUManager) TryPackSetRangeIncreasePrice(price *big.Int) ([]byte, error) {
	return fUManager.abi.Pack("setRangeIncreasePrice", price)
}

// PackSetSampleIncreaseLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf7690bfe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSampleIncreaseLimit(uint256 _lim) returns()
func (fUManager *FUManager) PackSetSampleIncreaseLimit(lim *big.Int) []byte {
	enc, err := fUManager.abi.Pack("setSampleIncreaseLimit", lim)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSampleIncreaseLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf7690bfe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSampleIncreaseLimit(uint256 _lim) returns()
func (fUManager *FUManager) TryPackSetSampleIncreaseLimit(lim *big.Int) ([]byte, error) {
	return fUManager.abi.Pack("setSampleIncreaseLimit", lim)
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (fUManager *FUManager) PackSwitchToProductionMode() []byte {
	enc, err := fUManager.abi.Pack("switchToProductionMode")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function switchToProductionMode() returns()
func (fUManager *FUManager) TryPackSwitchToProductionMode() ([]byte, error) {
	return fUManager.abi.Pack("switchToProductionMode")
}

// PackTimelockedCalls is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74e6310e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (fUManager *FUManager) PackTimelockedCalls(selector [4]byte) []byte {
	enc, err := fUManager.abi.Pack("timelockedCalls", selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTimelockedCalls is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74e6310e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (fUManager *FUManager) TryPackTimelockedCalls(selector [4]byte) ([]byte, error) {
	return fUManager.abi.Pack("timelockedCalls", selector)
}

// TimelockedCallsOutput serves as a container for the return parameters of contract
// method TimelockedCalls.
type TimelockedCallsOutput struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}

// UnpackTimelockedCalls is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (fUManager *FUManager) UnpackTimelockedCalls(data []byte) (TimelockedCallsOutput, error) {
	out, err := fUManager.abi.Unpack("timelockedCalls", data)
	outstruct := new(TimelockedCallsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AllowedAfterTimestamp = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.EncodedCall = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	return *outstruct, nil
}

// PackTotalInflationAuthorizedWei is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd0c1c393.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function totalInflationAuthorizedWei() view returns(uint256)
func (fUManager *FUManager) PackTotalInflationAuthorizedWei() []byte {
	enc, err := fUManager.abi.Pack("totalInflationAuthorizedWei")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTotalInflationAuthorizedWei is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd0c1c393.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function totalInflationAuthorizedWei() view returns(uint256)
func (fUManager *FUManager) TryPackTotalInflationAuthorizedWei() ([]byte, error) {
	return fUManager.abi.Pack("totalInflationAuthorizedWei")
}

// UnpackTotalInflationAuthorizedWei is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd0c1c393.
//
// Solidity: function totalInflationAuthorizedWei() view returns(uint256)
func (fUManager *FUManager) UnpackTotalInflationAuthorizedWei(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("totalInflationAuthorizedWei", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTotalInflationReceivedWei is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa5555aea.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function totalInflationReceivedWei() view returns(uint256)
func (fUManager *FUManager) PackTotalInflationReceivedWei() []byte {
	enc, err := fUManager.abi.Pack("totalInflationReceivedWei")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTotalInflationReceivedWei is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa5555aea.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function totalInflationReceivedWei() view returns(uint256)
func (fUManager *FUManager) TryPackTotalInflationReceivedWei() ([]byte, error) {
	return fUManager.abi.Pack("totalInflationReceivedWei")
}

// UnpackTotalInflationReceivedWei is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa5555aea.
//
// Solidity: function totalInflationReceivedWei() view returns(uint256)
func (fUManager *FUManager) UnpackTotalInflationReceivedWei(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("totalInflationReceivedWei", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTotalInflationRewardsOfferedWei is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd76b69c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function totalInflationRewardsOfferedWei() view returns(uint256)
func (fUManager *FUManager) PackTotalInflationRewardsOfferedWei() []byte {
	enc, err := fUManager.abi.Pack("totalInflationRewardsOfferedWei")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTotalInflationRewardsOfferedWei is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd76b69c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function totalInflationRewardsOfferedWei() view returns(uint256)
func (fUManager *FUManager) TryPackTotalInflationRewardsOfferedWei() ([]byte, error) {
	return fUManager.abi.Pack("totalInflationRewardsOfferedWei")
}

// UnpackTotalInflationRewardsOfferedWei is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbd76b69c.
//
// Solidity: function totalInflationRewardsOfferedWei() view returns(uint256)
func (fUManager *FUManager) UnpackTotalInflationRewardsOfferedWei(data []byte) (*big.Int, error) {
	out, err := fUManager.abi.Unpack("totalInflationRewardsOfferedWei", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTriggerRewardEpochSwitchover is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91f25679.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function triggerRewardEpochSwitchover(uint24 _currentRewardEpochId, uint64 _currentRewardEpochExpectedEndTs, uint64 _rewardEpochDurationSeconds) returns()
func (fUManager *FUManager) PackTriggerRewardEpochSwitchover(currentRewardEpochId *big.Int, currentRewardEpochExpectedEndTs uint64, rewardEpochDurationSeconds uint64) []byte {
	enc, err := fUManager.abi.Pack("triggerRewardEpochSwitchover", currentRewardEpochId, currentRewardEpochExpectedEndTs, rewardEpochDurationSeconds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTriggerRewardEpochSwitchover is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x91f25679.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function triggerRewardEpochSwitchover(uint24 _currentRewardEpochId, uint64 _currentRewardEpochExpectedEndTs, uint64 _rewardEpochDurationSeconds) returns()
func (fUManager *FUManager) TryPackTriggerRewardEpochSwitchover(currentRewardEpochId *big.Int, currentRewardEpochExpectedEndTs uint64, rewardEpochDurationSeconds uint64) ([]byte, error) {
	return fUManager.abi.Pack("triggerRewardEpochSwitchover", currentRewardEpochId, currentRewardEpochExpectedEndTs, rewardEpochDurationSeconds)
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (fUManager *FUManager) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := fUManager.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (fUManager *FUManager) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return fUManager.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// FUManagerDailyAuthorizedInflationSet represents a DailyAuthorizedInflationSet event raised by the FUManager contract.
type FUManagerDailyAuthorizedInflationSet struct {
	AuthorizedAmountWei *big.Int
	Raw                 *types.Log // Blockchain specific contextual infos
}

const FUManagerDailyAuthorizedInflationSetEventName = "DailyAuthorizedInflationSet"

// ContractEventName returns the user-defined event name.
func (FUManagerDailyAuthorizedInflationSet) ContractEventName() string {
	return FUManagerDailyAuthorizedInflationSetEventName
}

// UnpackDailyAuthorizedInflationSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DailyAuthorizedInflationSet(uint256 authorizedAmountWei)
func (fUManager *FUManager) UnpackDailyAuthorizedInflationSetEvent(log *types.Log) (*FUManagerDailyAuthorizedInflationSet, error) {
	event := "DailyAuthorizedInflationSet"
	if len(log.Topics) == 0 || log.Topics[0] != fUManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUManagerDailyAuthorizedInflationSet)
	if len(log.Data) > 0 {
		if err := fUManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// FUManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FUManager contract.
type FUManagerGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   *types.Log // Blockchain specific contextual infos
}

const FUManagerGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (FUManagerGovernanceCallTimelocked) ContractEventName() string {
	return FUManagerGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (fUManager *FUManager) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*FUManagerGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != fUManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUManagerGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := fUManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// FUManagerGovernanceInitialised represents a GovernanceInitialised event raised by the FUManager contract.
type FUManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const FUManagerGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (FUManagerGovernanceInitialised) ContractEventName() string {
	return FUManagerGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (fUManager *FUManager) UnpackGovernanceInitialisedEvent(log *types.Log) (*FUManagerGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != fUManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUManagerGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := fUManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// FUManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FUManager contract.
type FUManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const FUManagerGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (FUManagerGovernedProductionModeEntered) ContractEventName() string {
	return FUManagerGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (fUManager *FUManager) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*FUManagerGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != fUManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUManagerGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := fUManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// FUManagerIncentiveOffered represents a IncentiveOffered event raised by the FUManager contract.
type FUManagerIncentiveOffered struct {
	RewardEpochId      *big.Int
	RangeIncrease      *big.Int
	SampleSizeIncrease *big.Int
	OfferAmount        *big.Int
	Raw                *types.Log // Blockchain specific contextual infos
}

const FUManagerIncentiveOfferedEventName = "IncentiveOffered"

// ContractEventName returns the user-defined event name.
func (FUManagerIncentiveOffered) ContractEventName() string {
	return FUManagerIncentiveOfferedEventName
}

// UnpackIncentiveOfferedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event IncentiveOffered(uint24 indexed rewardEpochId, uint256 rangeIncrease, uint256 sampleSizeIncrease, uint256 offerAmount)
func (fUManager *FUManager) UnpackIncentiveOfferedEvent(log *types.Log) (*FUManagerIncentiveOffered, error) {
	event := "IncentiveOffered"
	if len(log.Topics) == 0 || log.Topics[0] != fUManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUManagerIncentiveOffered)
	if len(log.Data) > 0 {
		if err := fUManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// FUManagerInflationReceived represents a InflationReceived event raised by the FUManager contract.
type FUManagerInflationReceived struct {
	AmountReceivedWei *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const FUManagerInflationReceivedEventName = "InflationReceived"

// ContractEventName returns the user-defined event name.
func (FUManagerInflationReceived) ContractEventName() string {
	return FUManagerInflationReceivedEventName
}

// UnpackInflationReceivedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InflationReceived(uint256 amountReceivedWei)
func (fUManager *FUManager) UnpackInflationReceivedEvent(log *types.Log) (*FUManagerInflationReceived, error) {
	event := "InflationReceived"
	if len(log.Topics) == 0 || log.Topics[0] != fUManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUManagerInflationReceived)
	if len(log.Data) > 0 {
		if err := fUManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// FUManagerInflationRewardsOffered represents a InflationRewardsOffered event raised by the FUManager contract.
type FUManagerInflationRewardsOffered struct {
	RewardEpochId      *big.Int
	FeedConfigurations []IFastUpdatesConfigurationFeedConfiguration
	Amount             *big.Int
	Raw                *types.Log // Blockchain specific contextual infos
}

const FUManagerInflationRewardsOfferedEventName = "InflationRewardsOffered"

// ContractEventName returns the user-defined event name.
func (FUManagerInflationRewardsOffered) ContractEventName() string {
	return FUManagerInflationRewardsOfferedEventName
}

// UnpackInflationRewardsOfferedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InflationRewardsOffered(uint24 indexed rewardEpochId, (bytes21,uint32,uint24)[] feedConfigurations, uint256 amount)
func (fUManager *FUManager) UnpackInflationRewardsOfferedEvent(log *types.Log) (*FUManagerInflationRewardsOffered, error) {
	event := "InflationRewardsOffered"
	if len(log.Topics) == 0 || log.Topics[0] != fUManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUManagerInflationRewardsOffered)
	if len(log.Data) > 0 {
		if err := fUManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// FUManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FUManager contract.
type FUManagerTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const FUManagerTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (FUManagerTimelockedGovernanceCallCanceled) ContractEventName() string {
	return FUManagerTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (fUManager *FUManager) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*FUManagerTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != fUManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUManagerTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := fUManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// FUManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FUManager contract.
type FUManagerTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const FUManagerTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (FUManagerTimelockedGovernanceCallExecuted) ContractEventName() string {
	return FUManagerTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (fUManager *FUManager) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*FUManagerTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != fUManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUManagerTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := fUManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}
