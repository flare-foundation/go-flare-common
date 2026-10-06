// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package offers

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

// IFtsoRewardOffersManagerOffer is an auto generated low-level Go binding around an user-defined struct.
type IFtsoRewardOffersManagerOffer struct {
	Amount                    *big.Int
	FeedId                    [21]byte
	MinRewardedTurnoutBIPS    uint16
	PrimaryBandRewardSharePPM *big.Int
	SecondaryBandWidthPPM     *big.Int
	ClaimBackAddress          common.Address
}

// OffersMetaData contains all meta data concerning the Offers contract.
var OffersMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_minimalRewardsOfferValueWei\",\"type\":\"uint128\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"authorizedAmountWei\",\"type\":\"uint256\"}],\"name\":\"DailyAuthorizedInflationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedWei\",\"type\":\"uint256\"}],\"name\":\"InflationReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"feedIds\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"decimals\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minRewardedTurnoutBIPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"primaryBandRewardSharePPM\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"secondaryBandWidthPPMs\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"mode\",\"type\":\"uint16\"}],\"name\":\"InflationRewardsOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueWei\",\"type\":\"uint256\"}],\"name\":\"MinimalRewardsOfferValueSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bytes21\",\"name\":\"feedId\",\"type\":\"bytes21\"},{\"indexed\":false,\"internalType\":\"int8\",\"name\":\"decimals\",\"type\":\"int8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minRewardedTurnoutBIPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"primaryBandRewardSharePPM\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"secondaryBandWidthPPM\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"}],\"name\":\"RewardsOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyAuthorizedInflation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftsoFeedDecimals\",\"outputs\":[{\"internalType\":\"contractIFtsoFeedDecimals\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftsoInflationConfigurations\",\"outputs\":[{\"internalType\":\"contractIFtsoInflationConfigurations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInflationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenPoolSupplyData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockedFundsWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalInflationAuthorizedWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalClaimedWei\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInflationAuthorizationReceivedTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInflationReceivedTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalRewardsOfferValueWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_nextRewardEpochId\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint120\",\"name\":\"amount\",\"type\":\"uint120\"},{\"internalType\":\"bytes21\",\"name\":\"feedId\",\"type\":\"bytes21\"},{\"internalType\":\"uint16\",\"name\":\"minRewardedTurnoutBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"primaryBandRewardSharePPM\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"secondaryBandWidthPPM\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"}],\"internalType\":\"structIFtsoRewardOffersManager.Offer[]\",\"name\":\"_rewardOffers\",\"type\":\"tuple[]\"}],\"name\":\"offerRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveInflation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_toAuthorizeWei\",\"type\":\"uint256\"}],\"name\":\"setDailyAuthorizedInflation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_minimalRewardsOfferValueWei\",\"type\":\"uint128\"}],\"name\":\"setMinimalRewardsOfferValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationAuthorizedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationReceivedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationRewardsOfferedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_currentRewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_currentRewardEpochExpectedEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardEpochDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"triggerRewardEpochSwitchover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Offers",
}

// Offers is an auto generated Go binding around an Ethereum contract.
type Offers struct {
	abi abi.ABI
}

// NewOffers creates a new instance of Offers.
func NewOffers() *Offers {
	parsed, err := OffersMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Offers{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Offers) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint128 _minimalRewardsOfferValueWei) returns()
func (offers *Offers) PackConstructor(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _minimalRewardsOfferValueWei *big.Int) []byte {
	enc, err := offers.abi.Pack("", _governanceSettings, _initialGovernance, _addressUpdater, _minimalRewardsOfferValueWei)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67fc4029.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (offers *Offers) PackCancelGovernanceCall(selector [4]byte) []byte {
	enc, err := offers.abi.Pack("cancelGovernanceCall", selector)
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
func (offers *Offers) TryPackCancelGovernanceCall(selector [4]byte) ([]byte, error) {
	return offers.abi.Pack("cancelGovernanceCall", selector)
}

// PackDailyAuthorizedInflation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x708e34ce.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (offers *Offers) PackDailyAuthorizedInflation() []byte {
	enc, err := offers.abi.Pack("dailyAuthorizedInflation")
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
func (offers *Offers) TryPackDailyAuthorizedInflation() ([]byte, error) {
	return offers.abi.Pack("dailyAuthorizedInflation")
}

// UnpackDailyAuthorizedInflation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x708e34ce.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (offers *Offers) UnpackDailyAuthorizedInflation(data []byte) (*big.Int, error) {
	out, err := offers.abi.Unpack("dailyAuthorizedInflation", data)
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
func (offers *Offers) PackExecuteGovernanceCall(selector [4]byte) []byte {
	enc, err := offers.abi.Pack("executeGovernanceCall", selector)
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
func (offers *Offers) TryPackExecuteGovernanceCall(selector [4]byte) ([]byte, error) {
	return offers.abi.Pack("executeGovernanceCall", selector)
}

// PackFlareSystemsManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaae7fc9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareSystemsManager() view returns(address)
func (offers *Offers) PackFlareSystemsManager() []byte {
	enc, err := offers.abi.Pack("flareSystemsManager")
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
func (offers *Offers) TryPackFlareSystemsManager() ([]byte, error) {
	return offers.abi.Pack("flareSystemsManager")
}

// UnpackFlareSystemsManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (offers *Offers) UnpackFlareSystemsManager(data []byte) (common.Address, error) {
	out, err := offers.abi.Unpack("flareSystemsManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFtsoFeedDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9065c974.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ftsoFeedDecimals() view returns(address)
func (offers *Offers) PackFtsoFeedDecimals() []byte {
	enc, err := offers.abi.Pack("ftsoFeedDecimals")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFtsoFeedDecimals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9065c974.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ftsoFeedDecimals() view returns(address)
func (offers *Offers) TryPackFtsoFeedDecimals() ([]byte, error) {
	return offers.abi.Pack("ftsoFeedDecimals")
}

// UnpackFtsoFeedDecimals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9065c974.
//
// Solidity: function ftsoFeedDecimals() view returns(address)
func (offers *Offers) UnpackFtsoFeedDecimals(data []byte) (common.Address, error) {
	out, err := offers.abi.Unpack("ftsoFeedDecimals", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFtsoInflationConfigurations is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc27fb624.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ftsoInflationConfigurations() view returns(address)
func (offers *Offers) PackFtsoInflationConfigurations() []byte {
	enc, err := offers.abi.Pack("ftsoInflationConfigurations")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFtsoInflationConfigurations is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc27fb624.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ftsoInflationConfigurations() view returns(address)
func (offers *Offers) TryPackFtsoInflationConfigurations() ([]byte, error) {
	return offers.abi.Pack("ftsoInflationConfigurations")
}

// UnpackFtsoInflationConfigurations is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc27fb624.
//
// Solidity: function ftsoInflationConfigurations() view returns(address)
func (offers *Offers) UnpackFtsoInflationConfigurations(data []byte) (common.Address, error) {
	out, err := offers.abi.Unpack("ftsoInflationConfigurations", data)
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
func (offers *Offers) PackGetAddressUpdater() []byte {
	enc, err := offers.abi.Pack("getAddressUpdater")
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
func (offers *Offers) TryPackGetAddressUpdater() ([]byte, error) {
	return offers.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (offers *Offers) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := offers.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetContractName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5f5ba72.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getContractName() pure returns(string)
func (offers *Offers) PackGetContractName() []byte {
	enc, err := offers.abi.Pack("getContractName")
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
func (offers *Offers) TryPackGetContractName() ([]byte, error) {
	return offers.abi.Pack("getContractName")
}

// UnpackGetContractName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (offers *Offers) UnpackGetContractName(data []byte) (string, error) {
	out, err := offers.abi.Unpack("getContractName", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackGetExpectedBalance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf04cd3b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (offers *Offers) PackGetExpectedBalance() []byte {
	enc, err := offers.abi.Pack("getExpectedBalance")
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
func (offers *Offers) TryPackGetExpectedBalance() ([]byte, error) {
	return offers.abi.Pack("getExpectedBalance")
}

// UnpackGetExpectedBalance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaf04cd3b.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (offers *Offers) UnpackGetExpectedBalance(data []byte) (*big.Int, error) {
	out, err := offers.abi.Unpack("getExpectedBalance", data)
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
func (offers *Offers) PackGetInflationAddress() []byte {
	enc, err := offers.abi.Pack("getInflationAddress")
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
func (offers *Offers) TryPackGetInflationAddress() ([]byte, error) {
	return offers.abi.Pack("getInflationAddress")
}

// UnpackGetInflationAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xed39d3f8.
//
// Solidity: function getInflationAddress() view returns(address)
func (offers *Offers) UnpackGetInflationAddress(data []byte) (common.Address, error) {
	out, err := offers.abi.Unpack("getInflationAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetTokenPoolSupplyData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2dafdbbf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTokenPoolSupplyData() view returns(uint256 _lockedFundsWei, uint256 _totalInflationAuthorizedWei, uint256 _totalClaimedWei)
func (offers *Offers) PackGetTokenPoolSupplyData() []byte {
	enc, err := offers.abi.Pack("getTokenPoolSupplyData")
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
func (offers *Offers) TryPackGetTokenPoolSupplyData() ([]byte, error) {
	return offers.abi.Pack("getTokenPoolSupplyData")
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
func (offers *Offers) UnpackGetTokenPoolSupplyData(data []byte) (GetTokenPoolSupplyDataOutput, error) {
	out, err := offers.abi.Unpack("getTokenPoolSupplyData", data)
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
func (offers *Offers) PackGovernance() []byte {
	enc, err := offers.abi.Pack("governance")
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
func (offers *Offers) TryPackGovernance() ([]byte, error) {
	return offers.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (offers *Offers) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := offers.abi.Unpack("governance", data)
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
func (offers *Offers) PackGovernanceSettings() []byte {
	enc, err := offers.abi.Pack("governanceSettings")
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
func (offers *Offers) TryPackGovernanceSettings() ([]byte, error) {
	return offers.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (offers *Offers) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := offers.abi.Unpack("governanceSettings", data)
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
func (offers *Offers) PackInitialise(governanceSettings common.Address, initialGovernance common.Address) []byte {
	enc, err := offers.abi.Pack("initialise", governanceSettings, initialGovernance)
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
func (offers *Offers) TryPackInitialise(governanceSettings common.Address, initialGovernance common.Address) ([]byte, error) {
	return offers.abi.Pack("initialise", governanceSettings, initialGovernance)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (offers *Offers) PackIsExecutor(address common.Address) []byte {
	enc, err := offers.abi.Pack("isExecutor", address)
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
func (offers *Offers) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return offers.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (offers *Offers) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := offers.abi.Unpack("isExecutor", data)
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
func (offers *Offers) PackLastInflationAuthorizationReceivedTs() []byte {
	enc, err := offers.abi.Pack("lastInflationAuthorizationReceivedTs")
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
func (offers *Offers) TryPackLastInflationAuthorizationReceivedTs() ([]byte, error) {
	return offers.abi.Pack("lastInflationAuthorizationReceivedTs")
}

// UnpackLastInflationAuthorizationReceivedTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x473252c4.
//
// Solidity: function lastInflationAuthorizationReceivedTs() view returns(uint256)
func (offers *Offers) UnpackLastInflationAuthorizationReceivedTs(data []byte) (*big.Int, error) {
	out, err := offers.abi.Unpack("lastInflationAuthorizationReceivedTs", data)
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
func (offers *Offers) PackLastInflationReceivedTs() []byte {
	enc, err := offers.abi.Pack("lastInflationReceivedTs")
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
func (offers *Offers) TryPackLastInflationReceivedTs() ([]byte, error) {
	return offers.abi.Pack("lastInflationReceivedTs")
}

// UnpackLastInflationReceivedTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x12afcf0b.
//
// Solidity: function lastInflationReceivedTs() view returns(uint256)
func (offers *Offers) UnpackLastInflationReceivedTs(data []byte) (*big.Int, error) {
	out, err := offers.abi.Unpack("lastInflationReceivedTs", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackMinimalRewardsOfferValueWei is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc85f3a46.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function minimalRewardsOfferValueWei() view returns(uint256)
func (offers *Offers) PackMinimalRewardsOfferValueWei() []byte {
	enc, err := offers.abi.Pack("minimalRewardsOfferValueWei")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMinimalRewardsOfferValueWei is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc85f3a46.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function minimalRewardsOfferValueWei() view returns(uint256)
func (offers *Offers) TryPackMinimalRewardsOfferValueWei() ([]byte, error) {
	return offers.abi.Pack("minimalRewardsOfferValueWei")
}

// UnpackMinimalRewardsOfferValueWei is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc85f3a46.
//
// Solidity: function minimalRewardsOfferValueWei() view returns(uint256)
func (offers *Offers) UnpackMinimalRewardsOfferValueWei(data []byte) (*big.Int, error) {
	out, err := offers.abi.Unpack("minimalRewardsOfferValueWei", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackOfferRewards is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb75a424f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function offerRewards(uint24 _nextRewardEpochId, (uint120,bytes21,uint16,uint24,uint24,address)[] _rewardOffers) payable returns()
func (offers *Offers) PackOfferRewards(nextRewardEpochId *big.Int, rewardOffers []IFtsoRewardOffersManagerOffer) []byte {
	enc, err := offers.abi.Pack("offerRewards", nextRewardEpochId, rewardOffers)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackOfferRewards is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb75a424f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function offerRewards(uint24 _nextRewardEpochId, (uint120,bytes21,uint16,uint24,uint24,address)[] _rewardOffers) payable returns()
func (offers *Offers) TryPackOfferRewards(nextRewardEpochId *big.Int, rewardOffers []IFtsoRewardOffersManagerOffer) ([]byte, error) {
	return offers.abi.Pack("offerRewards", nextRewardEpochId, rewardOffers)
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (offers *Offers) PackProductionMode() []byte {
	enc, err := offers.abi.Pack("productionMode")
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
func (offers *Offers) TryPackProductionMode() ([]byte, error) {
	return offers.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (offers *Offers) UnpackProductionMode(data []byte) (bool, error) {
	out, err := offers.abi.Unpack("productionMode", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackReceiveInflation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06201f1d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function receiveInflation() payable returns()
func (offers *Offers) PackReceiveInflation() []byte {
	enc, err := offers.abi.Pack("receiveInflation")
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
func (offers *Offers) TryPackReceiveInflation() ([]byte, error) {
	return offers.abi.Pack("receiveInflation")
}

// PackRewardManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f4ef8a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardManager() view returns(address)
func (offers *Offers) PackRewardManager() []byte {
	enc, err := offers.abi.Pack("rewardManager")
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
func (offers *Offers) TryPackRewardManager() ([]byte, error) {
	return offers.abi.Pack("rewardManager")
}

// UnpackRewardManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (offers *Offers) UnpackRewardManager(data []byte) (common.Address, error) {
	out, err := offers.abi.Unpack("rewardManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetDailyAuthorizedInflation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2739563.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setDailyAuthorizedInflation(uint256 _toAuthorizeWei) returns()
func (offers *Offers) PackSetDailyAuthorizedInflation(toAuthorizeWei *big.Int) []byte {
	enc, err := offers.abi.Pack("setDailyAuthorizedInflation", toAuthorizeWei)
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
func (offers *Offers) TryPackSetDailyAuthorizedInflation(toAuthorizeWei *big.Int) ([]byte, error) {
	return offers.abi.Pack("setDailyAuthorizedInflation", toAuthorizeWei)
}

// PackSetMinimalRewardsOfferValue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc9f19dd2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMinimalRewardsOfferValue(uint128 _minimalRewardsOfferValueWei) returns()
func (offers *Offers) PackSetMinimalRewardsOfferValue(minimalRewardsOfferValueWei *big.Int) []byte {
	enc, err := offers.abi.Pack("setMinimalRewardsOfferValue", minimalRewardsOfferValueWei)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMinimalRewardsOfferValue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc9f19dd2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMinimalRewardsOfferValue(uint128 _minimalRewardsOfferValueWei) returns()
func (offers *Offers) TryPackSetMinimalRewardsOfferValue(minimalRewardsOfferValueWei *big.Int) ([]byte, error) {
	return offers.abi.Pack("setMinimalRewardsOfferValue", minimalRewardsOfferValueWei)
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (offers *Offers) PackSwitchToProductionMode() []byte {
	enc, err := offers.abi.Pack("switchToProductionMode")
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
func (offers *Offers) TryPackSwitchToProductionMode() ([]byte, error) {
	return offers.abi.Pack("switchToProductionMode")
}

// PackTimelockedCalls is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74e6310e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (offers *Offers) PackTimelockedCalls(selector [4]byte) []byte {
	enc, err := offers.abi.Pack("timelockedCalls", selector)
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
func (offers *Offers) TryPackTimelockedCalls(selector [4]byte) ([]byte, error) {
	return offers.abi.Pack("timelockedCalls", selector)
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
func (offers *Offers) UnpackTimelockedCalls(data []byte) (TimelockedCallsOutput, error) {
	out, err := offers.abi.Unpack("timelockedCalls", data)
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
func (offers *Offers) PackTotalInflationAuthorizedWei() []byte {
	enc, err := offers.abi.Pack("totalInflationAuthorizedWei")
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
func (offers *Offers) TryPackTotalInflationAuthorizedWei() ([]byte, error) {
	return offers.abi.Pack("totalInflationAuthorizedWei")
}

// UnpackTotalInflationAuthorizedWei is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd0c1c393.
//
// Solidity: function totalInflationAuthorizedWei() view returns(uint256)
func (offers *Offers) UnpackTotalInflationAuthorizedWei(data []byte) (*big.Int, error) {
	out, err := offers.abi.Unpack("totalInflationAuthorizedWei", data)
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
func (offers *Offers) PackTotalInflationReceivedWei() []byte {
	enc, err := offers.abi.Pack("totalInflationReceivedWei")
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
func (offers *Offers) TryPackTotalInflationReceivedWei() ([]byte, error) {
	return offers.abi.Pack("totalInflationReceivedWei")
}

// UnpackTotalInflationReceivedWei is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa5555aea.
//
// Solidity: function totalInflationReceivedWei() view returns(uint256)
func (offers *Offers) UnpackTotalInflationReceivedWei(data []byte) (*big.Int, error) {
	out, err := offers.abi.Unpack("totalInflationReceivedWei", data)
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
func (offers *Offers) PackTotalInflationRewardsOfferedWei() []byte {
	enc, err := offers.abi.Pack("totalInflationRewardsOfferedWei")
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
func (offers *Offers) TryPackTotalInflationRewardsOfferedWei() ([]byte, error) {
	return offers.abi.Pack("totalInflationRewardsOfferedWei")
}

// UnpackTotalInflationRewardsOfferedWei is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbd76b69c.
//
// Solidity: function totalInflationRewardsOfferedWei() view returns(uint256)
func (offers *Offers) UnpackTotalInflationRewardsOfferedWei(data []byte) (*big.Int, error) {
	out, err := offers.abi.Unpack("totalInflationRewardsOfferedWei", data)
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
func (offers *Offers) PackTriggerRewardEpochSwitchover(currentRewardEpochId *big.Int, currentRewardEpochExpectedEndTs uint64, rewardEpochDurationSeconds uint64) []byte {
	enc, err := offers.abi.Pack("triggerRewardEpochSwitchover", currentRewardEpochId, currentRewardEpochExpectedEndTs, rewardEpochDurationSeconds)
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
func (offers *Offers) TryPackTriggerRewardEpochSwitchover(currentRewardEpochId *big.Int, currentRewardEpochExpectedEndTs uint64, rewardEpochDurationSeconds uint64) ([]byte, error) {
	return offers.abi.Pack("triggerRewardEpochSwitchover", currentRewardEpochId, currentRewardEpochExpectedEndTs, rewardEpochDurationSeconds)
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (offers *Offers) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := offers.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (offers *Offers) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return offers.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// OffersDailyAuthorizedInflationSet represents a DailyAuthorizedInflationSet event raised by the Offers contract.
type OffersDailyAuthorizedInflationSet struct {
	AuthorizedAmountWei *big.Int
	Raw                 *types.Log // Blockchain specific contextual infos
}

const OffersDailyAuthorizedInflationSetEventName = "DailyAuthorizedInflationSet"

// ContractEventName returns the user-defined event name.
func (OffersDailyAuthorizedInflationSet) ContractEventName() string {
	return OffersDailyAuthorizedInflationSetEventName
}

// UnpackDailyAuthorizedInflationSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DailyAuthorizedInflationSet(uint256 authorizedAmountWei)
func (offers *Offers) UnpackDailyAuthorizedInflationSetEvent(log *types.Log) (*OffersDailyAuthorizedInflationSet, error) {
	event := "DailyAuthorizedInflationSet"
	if len(log.Topics) == 0 || log.Topics[0] != offers.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OffersDailyAuthorizedInflationSet)
	if len(log.Data) > 0 {
		if err := offers.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range offers.abi.Events[event].Inputs {
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

// OffersGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Offers contract.
type OffersGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   *types.Log // Blockchain specific contextual infos
}

const OffersGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (OffersGovernanceCallTimelocked) ContractEventName() string {
	return OffersGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (offers *Offers) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*OffersGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != offers.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OffersGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := offers.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range offers.abi.Events[event].Inputs {
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

// OffersGovernanceInitialised represents a GovernanceInitialised event raised by the Offers contract.
type OffersGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const OffersGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (OffersGovernanceInitialised) ContractEventName() string {
	return OffersGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (offers *Offers) UnpackGovernanceInitialisedEvent(log *types.Log) (*OffersGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != offers.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OffersGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := offers.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range offers.abi.Events[event].Inputs {
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

// OffersGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the Offers contract.
type OffersGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const OffersGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (OffersGovernedProductionModeEntered) ContractEventName() string {
	return OffersGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (offers *Offers) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*OffersGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != offers.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OffersGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := offers.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range offers.abi.Events[event].Inputs {
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

// OffersInflationReceived represents a InflationReceived event raised by the Offers contract.
type OffersInflationReceived struct {
	AmountReceivedWei *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const OffersInflationReceivedEventName = "InflationReceived"

// ContractEventName returns the user-defined event name.
func (OffersInflationReceived) ContractEventName() string {
	return OffersInflationReceivedEventName
}

// UnpackInflationReceivedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InflationReceived(uint256 amountReceivedWei)
func (offers *Offers) UnpackInflationReceivedEvent(log *types.Log) (*OffersInflationReceived, error) {
	event := "InflationReceived"
	if len(log.Topics) == 0 || log.Topics[0] != offers.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OffersInflationReceived)
	if len(log.Data) > 0 {
		if err := offers.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range offers.abi.Events[event].Inputs {
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

// OffersInflationRewardsOffered represents a InflationRewardsOffered event raised by the Offers contract.
type OffersInflationRewardsOffered struct {
	RewardEpochId             *big.Int
	FeedIds                   []byte
	Decimals                  []byte
	Amount                    *big.Int
	MinRewardedTurnoutBIPS    uint16
	PrimaryBandRewardSharePPM *big.Int
	SecondaryBandWidthPPMs    []byte
	Mode                      uint16
	Raw                       *types.Log // Blockchain specific contextual infos
}

const OffersInflationRewardsOfferedEventName = "InflationRewardsOffered"

// ContractEventName returns the user-defined event name.
func (OffersInflationRewardsOffered) ContractEventName() string {
	return OffersInflationRewardsOfferedEventName
}

// UnpackInflationRewardsOfferedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InflationRewardsOffered(uint24 indexed rewardEpochId, bytes feedIds, bytes decimals, uint256 amount, uint16 minRewardedTurnoutBIPS, uint24 primaryBandRewardSharePPM, bytes secondaryBandWidthPPMs, uint16 mode)
func (offers *Offers) UnpackInflationRewardsOfferedEvent(log *types.Log) (*OffersInflationRewardsOffered, error) {
	event := "InflationRewardsOffered"
	if len(log.Topics) == 0 || log.Topics[0] != offers.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OffersInflationRewardsOffered)
	if len(log.Data) > 0 {
		if err := offers.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range offers.abi.Events[event].Inputs {
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

// OffersMinimalRewardsOfferValueSet represents a MinimalRewardsOfferValueSet event raised by the Offers contract.
type OffersMinimalRewardsOfferValueSet struct {
	ValueWei *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const OffersMinimalRewardsOfferValueSetEventName = "MinimalRewardsOfferValueSet"

// ContractEventName returns the user-defined event name.
func (OffersMinimalRewardsOfferValueSet) ContractEventName() string {
	return OffersMinimalRewardsOfferValueSetEventName
}

// UnpackMinimalRewardsOfferValueSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinimalRewardsOfferValueSet(uint256 valueWei)
func (offers *Offers) UnpackMinimalRewardsOfferValueSetEvent(log *types.Log) (*OffersMinimalRewardsOfferValueSet, error) {
	event := "MinimalRewardsOfferValueSet"
	if len(log.Topics) == 0 || log.Topics[0] != offers.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OffersMinimalRewardsOfferValueSet)
	if len(log.Data) > 0 {
		if err := offers.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range offers.abi.Events[event].Inputs {
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

// OffersRewardsOffered represents a RewardsOffered event raised by the Offers contract.
type OffersRewardsOffered struct {
	RewardEpochId             *big.Int
	FeedId                    [21]byte
	Decimals                  int8
	Amount                    *big.Int
	MinRewardedTurnoutBIPS    uint16
	PrimaryBandRewardSharePPM *big.Int
	SecondaryBandWidthPPM     *big.Int
	ClaimBackAddress          common.Address
	Raw                       *types.Log // Blockchain specific contextual infos
}

const OffersRewardsOfferedEventName = "RewardsOffered"

// ContractEventName returns the user-defined event name.
func (OffersRewardsOffered) ContractEventName() string {
	return OffersRewardsOfferedEventName
}

// UnpackRewardsOfferedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RewardsOffered(uint24 indexed rewardEpochId, bytes21 feedId, int8 decimals, uint256 amount, uint16 minRewardedTurnoutBIPS, uint24 primaryBandRewardSharePPM, uint24 secondaryBandWidthPPM, address claimBackAddress)
func (offers *Offers) UnpackRewardsOfferedEvent(log *types.Log) (*OffersRewardsOffered, error) {
	event := "RewardsOffered"
	if len(log.Topics) == 0 || log.Topics[0] != offers.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OffersRewardsOffered)
	if len(log.Data) > 0 {
		if err := offers.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range offers.abi.Events[event].Inputs {
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

// OffersTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the Offers contract.
type OffersTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const OffersTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (OffersTimelockedGovernanceCallCanceled) ContractEventName() string {
	return OffersTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (offers *Offers) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*OffersTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != offers.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OffersTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := offers.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range offers.abi.Events[event].Inputs {
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

// OffersTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the Offers contract.
type OffersTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const OffersTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (OffersTimelockedGovernanceCallExecuted) ContractEventName() string {
	return OffersTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (offers *Offers) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*OffersTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != offers.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OffersTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := offers.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range offers.abi.Events[event].Inputs {
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
