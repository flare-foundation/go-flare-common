// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teerewardoffersmanager

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

// TeeRewardOffersManagerMetaData contains all meta data concerning the TeeRewardOffersManager contract.
var TeeRewardOffersManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeOwnersPPMValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockCallNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"authorizedAmountWei\",\"type\":\"uint256\"}],\"name\":\"DailyAuthorizedInflationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedWei\",\"type\":\"uint256\"}],\"name\":\"InflationReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teeOwnersPPM\",\"type\":\"uint256\"}],\"name\":\"InflationRewardsOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"teeOwnersPPM\",\"type\":\"uint24\"}],\"name\":\"TeeOwnersPPMSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyAuthorizedInflation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInflationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenPoolSupplyData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockedFundsWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalInflationAuthorizedWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalClaimedWei\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"_teeOwnersPPM\",\"type\":\"uint24\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInflationAuthorizationReceivedTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInflationReceivedTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveInflation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_toAuthorizeWei\",\"type\":\"uint256\"}],\"name\":\"setDailyAuthorizedInflation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_teeOwnersPPM\",\"type\":\"uint24\"}],\"name\":\"setTeeOwnersPPM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeOwnersPPM\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationAuthorizedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationReceivedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationRewardsOfferedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_currentRewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_currentRewardEpochExpectedEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardEpochDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"triggerRewardEpochSwitchover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "TeeRewardOffersManager",
}

// TeeRewardOffersManager is an auto generated Go binding around an Ethereum contract.
type TeeRewardOffersManager struct {
	abi abi.ABI
}

// NewTeeRewardOffersManager creates a new instance of TeeRewardOffersManager.
func NewTeeRewardOffersManager() *TeeRewardOffersManager {
	parsed, err := TeeRewardOffersManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TeeRewardOffersManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TeeRewardOffersManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teeRewardOffersManager *TeeRewardOffersManager) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := teeRewardOffersManager.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16fc2f6d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (teeRewardOffersManager *TeeRewardOffersManager) PackCancelGovernanceCall(encodedCall []byte) []byte {
	enc, err := teeRewardOffersManager.abi.Pack("cancelGovernanceCall", encodedCall)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16fc2f6d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackCancelGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("cancelGovernanceCall", encodedCall)
}

// PackDailyAuthorizedInflation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x708e34ce.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (teeRewardOffersManager *TeeRewardOffersManager) PackDailyAuthorizedInflation() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("dailyAuthorizedInflation")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackDailyAuthorizedInflation() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("dailyAuthorizedInflation")
}

// UnpackDailyAuthorizedInflation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x708e34ce.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackDailyAuthorizedInflation(data []byte) (*big.Int, error) {
	out, err := teeRewardOffersManager.abi.Unpack("dailyAuthorizedInflation", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (teeRewardOffersManager *TeeRewardOffersManager) PackExecuteGovernanceCall(encodedCall []byte) []byte {
	enc, err := teeRewardOffersManager.abi.Pack("executeGovernanceCall", encodedCall)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackExecuteGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("executeGovernanceCall", encodedCall)
}

// PackFlareSystemsManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaae7fc9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareSystemsManager() view returns(address)
func (teeRewardOffersManager *TeeRewardOffersManager) PackFlareSystemsManager() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("flareSystemsManager")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackFlareSystemsManager() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("flareSystemsManager")
}

// UnpackFlareSystemsManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackFlareSystemsManager(data []byte) (common.Address, error) {
	out, err := teeRewardOffersManager.abi.Unpack("flareSystemsManager", data)
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
func (teeRewardOffersManager *TeeRewardOffersManager) PackGetAddressUpdater() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("getAddressUpdater")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackGetAddressUpdater() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := teeRewardOffersManager.abi.Unpack("getAddressUpdater", data)
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
func (teeRewardOffersManager *TeeRewardOffersManager) PackGetContractName() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("getContractName")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackGetContractName() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("getContractName")
}

// UnpackGetContractName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackGetContractName(data []byte) (string, error) {
	out, err := teeRewardOffersManager.abi.Unpack("getContractName", data)
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
func (teeRewardOffersManager *TeeRewardOffersManager) PackGetExpectedBalance() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("getExpectedBalance")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackGetExpectedBalance() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("getExpectedBalance")
}

// UnpackGetExpectedBalance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaf04cd3b.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackGetExpectedBalance(data []byte) (*big.Int, error) {
	out, err := teeRewardOffersManager.abi.Unpack("getExpectedBalance", data)
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
func (teeRewardOffersManager *TeeRewardOffersManager) PackGetInflationAddress() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("getInflationAddress")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackGetInflationAddress() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("getInflationAddress")
}

// UnpackGetInflationAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xed39d3f8.
//
// Solidity: function getInflationAddress() view returns(address)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackGetInflationAddress(data []byte) (common.Address, error) {
	out, err := teeRewardOffersManager.abi.Unpack("getInflationAddress", data)
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
func (teeRewardOffersManager *TeeRewardOffersManager) PackGetTokenPoolSupplyData() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("getTokenPoolSupplyData")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackGetTokenPoolSupplyData() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("getTokenPoolSupplyData")
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
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackGetTokenPoolSupplyData(data []byte) (GetTokenPoolSupplyDataOutput, error) {
	out, err := teeRewardOffersManager.abi.Unpack("getTokenPoolSupplyData", data)
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
func (teeRewardOffersManager *TeeRewardOffersManager) PackGovernance() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("governance")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackGovernance() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := teeRewardOffersManager.abi.Unpack("governance", data)
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
func (teeRewardOffersManager *TeeRewardOffersManager) PackGovernanceSettings() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("governanceSettings")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackGovernanceSettings() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := teeRewardOffersManager.abi.Unpack("governanceSettings", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c60da1b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function implementation() view returns(address)
func (teeRewardOffersManager *TeeRewardOffersManager) PackImplementation() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("implementation")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c60da1b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function implementation() view returns(address)
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackImplementation() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("implementation")
}

// UnpackImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackImplementation(data []byte) (common.Address, error) {
	out, err := teeRewardOffersManager.abi.Unpack("implementation", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4627023.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint24 _teeOwnersPPM) returns()
func (teeRewardOffersManager *TeeRewardOffersManager) PackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address, teeOwnersPPM *big.Int) []byte {
	enc, err := teeRewardOffersManager.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater, teeOwnersPPM)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4627023.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint24 _teeOwnersPPM) returns()
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address, teeOwnersPPM *big.Int) ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater, teeOwnersPPM)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teeRewardOffersManager *TeeRewardOffersManager) PackIsExecutor(address common.Address) []byte {
	enc, err := teeRewardOffersManager.abi.Pack("isExecutor", address)
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := teeRewardOffersManager.abi.Unpack("isExecutor", data)
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
func (teeRewardOffersManager *TeeRewardOffersManager) PackLastInflationAuthorizationReceivedTs() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("lastInflationAuthorizationReceivedTs")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackLastInflationAuthorizationReceivedTs() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("lastInflationAuthorizationReceivedTs")
}

// UnpackLastInflationAuthorizationReceivedTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x473252c4.
//
// Solidity: function lastInflationAuthorizationReceivedTs() view returns(uint256)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackLastInflationAuthorizationReceivedTs(data []byte) (*big.Int, error) {
	out, err := teeRewardOffersManager.abi.Unpack("lastInflationAuthorizationReceivedTs", data)
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
func (teeRewardOffersManager *TeeRewardOffersManager) PackLastInflationReceivedTs() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("lastInflationReceivedTs")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackLastInflationReceivedTs() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("lastInflationReceivedTs")
}

// UnpackLastInflationReceivedTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x12afcf0b.
//
// Solidity: function lastInflationReceivedTs() view returns(uint256)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackLastInflationReceivedTs(data []byte) (*big.Int, error) {
	out, err := teeRewardOffersManager.abi.Unpack("lastInflationReceivedTs", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (teeRewardOffersManager *TeeRewardOffersManager) PackProductionMode() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("productionMode")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackProductionMode() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackProductionMode(data []byte) (bool, error) {
	out, err := teeRewardOffersManager.abi.Unpack("productionMode", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (teeRewardOffersManager *TeeRewardOffersManager) PackProxiableUUID() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackProxiableUUID() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("proxiableUUID")
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := teeRewardOffersManager.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackReceiveInflation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06201f1d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function receiveInflation() payable returns()
func (teeRewardOffersManager *TeeRewardOffersManager) PackReceiveInflation() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("receiveInflation")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackReceiveInflation() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("receiveInflation")
}

// PackRewardManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f4ef8a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardManager() view returns(address)
func (teeRewardOffersManager *TeeRewardOffersManager) PackRewardManager() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("rewardManager")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackRewardManager() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("rewardManager")
}

// UnpackRewardManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackRewardManager(data []byte) (common.Address, error) {
	out, err := teeRewardOffersManager.abi.Unpack("rewardManager", data)
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
func (teeRewardOffersManager *TeeRewardOffersManager) PackSetDailyAuthorizedInflation(toAuthorizeWei *big.Int) []byte {
	enc, err := teeRewardOffersManager.abi.Pack("setDailyAuthorizedInflation", toAuthorizeWei)
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackSetDailyAuthorizedInflation(toAuthorizeWei *big.Int) ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("setDailyAuthorizedInflation", toAuthorizeWei)
}

// PackSetTeeOwnersPPM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9d81e53d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setTeeOwnersPPM(uint24 _teeOwnersPPM) returns()
func (teeRewardOffersManager *TeeRewardOffersManager) PackSetTeeOwnersPPM(teeOwnersPPM *big.Int) []byte {
	enc, err := teeRewardOffersManager.abi.Pack("setTeeOwnersPPM", teeOwnersPPM)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetTeeOwnersPPM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9d81e53d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setTeeOwnersPPM(uint24 _teeOwnersPPM) returns()
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackSetTeeOwnersPPM(teeOwnersPPM *big.Int) ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("setTeeOwnersPPM", teeOwnersPPM)
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (teeRewardOffersManager *TeeRewardOffersManager) PackSwitchToProductionMode() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("switchToProductionMode")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackSwitchToProductionMode() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("switchToProductionMode")
}

// PackTeeOwnersPPM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeeddae62.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function teeOwnersPPM() view returns(uint24)
func (teeRewardOffersManager *TeeRewardOffersManager) PackTeeOwnersPPM() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("teeOwnersPPM")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTeeOwnersPPM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeeddae62.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function teeOwnersPPM() view returns(uint24)
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackTeeOwnersPPM() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("teeOwnersPPM")
}

// UnpackTeeOwnersPPM is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xeeddae62.
//
// Solidity: function teeOwnersPPM() view returns(uint24)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackTeeOwnersPPM(data []byte) (*big.Int, error) {
	out, err := teeRewardOffersManager.abi.Unpack("teeOwnersPPM", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackTotalInflationAuthorizedWei is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd0c1c393.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function totalInflationAuthorizedWei() view returns(uint256)
func (teeRewardOffersManager *TeeRewardOffersManager) PackTotalInflationAuthorizedWei() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("totalInflationAuthorizedWei")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackTotalInflationAuthorizedWei() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("totalInflationAuthorizedWei")
}

// UnpackTotalInflationAuthorizedWei is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd0c1c393.
//
// Solidity: function totalInflationAuthorizedWei() view returns(uint256)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackTotalInflationAuthorizedWei(data []byte) (*big.Int, error) {
	out, err := teeRewardOffersManager.abi.Unpack("totalInflationAuthorizedWei", data)
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
func (teeRewardOffersManager *TeeRewardOffersManager) PackTotalInflationReceivedWei() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("totalInflationReceivedWei")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackTotalInflationReceivedWei() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("totalInflationReceivedWei")
}

// UnpackTotalInflationReceivedWei is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa5555aea.
//
// Solidity: function totalInflationReceivedWei() view returns(uint256)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackTotalInflationReceivedWei(data []byte) (*big.Int, error) {
	out, err := teeRewardOffersManager.abi.Unpack("totalInflationReceivedWei", data)
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
func (teeRewardOffersManager *TeeRewardOffersManager) PackTotalInflationRewardsOfferedWei() []byte {
	enc, err := teeRewardOffersManager.abi.Pack("totalInflationRewardsOfferedWei")
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackTotalInflationRewardsOfferedWei() ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("totalInflationRewardsOfferedWei")
}

// UnpackTotalInflationRewardsOfferedWei is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbd76b69c.
//
// Solidity: function totalInflationRewardsOfferedWei() view returns(uint256)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackTotalInflationRewardsOfferedWei(data []byte) (*big.Int, error) {
	out, err := teeRewardOffersManager.abi.Unpack("totalInflationRewardsOfferedWei", data)
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
func (teeRewardOffersManager *TeeRewardOffersManager) PackTriggerRewardEpochSwitchover(currentRewardEpochId *big.Int, currentRewardEpochExpectedEndTs uint64, rewardEpochDurationSeconds uint64) []byte {
	enc, err := teeRewardOffersManager.abi.Pack("triggerRewardEpochSwitchover", currentRewardEpochId, currentRewardEpochExpectedEndTs, rewardEpochDurationSeconds)
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackTriggerRewardEpochSwitchover(currentRewardEpochId *big.Int, currentRewardEpochExpectedEndTs uint64, rewardEpochDurationSeconds uint64) ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("triggerRewardEpochSwitchover", currentRewardEpochId, currentRewardEpochExpectedEndTs, rewardEpochDurationSeconds)
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (teeRewardOffersManager *TeeRewardOffersManager) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := teeRewardOffersManager.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (teeRewardOffersManager *TeeRewardOffersManager) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := teeRewardOffersManager.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (teeRewardOffersManager *TeeRewardOffersManager) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return teeRewardOffersManager.abi.Pack("upgradeToAndCall", newImplementation, data)
}

// TeeRewardOffersManagerDailyAuthorizedInflationSet represents a DailyAuthorizedInflationSet event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerDailyAuthorizedInflationSet struct {
	AuthorizedAmountWei *big.Int
	Raw                 *types.Log // Blockchain specific contextual infos
}

const TeeRewardOffersManagerDailyAuthorizedInflationSetEventName = "DailyAuthorizedInflationSet"

// ContractEventName returns the user-defined event name.
func (TeeRewardOffersManagerDailyAuthorizedInflationSet) ContractEventName() string {
	return TeeRewardOffersManagerDailyAuthorizedInflationSetEventName
}

// UnpackDailyAuthorizedInflationSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DailyAuthorizedInflationSet(uint256 authorizedAmountWei)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackDailyAuthorizedInflationSetEvent(log *types.Log) (*TeeRewardOffersManagerDailyAuthorizedInflationSet, error) {
	event := "DailyAuthorizedInflationSet"
	if len(log.Topics) == 0 || log.Topics[0] != teeRewardOffersManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeeRewardOffersManagerDailyAuthorizedInflationSet)
	if len(log.Data) > 0 {
		if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teeRewardOffersManager.abi.Events[event].Inputs {
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

// TeeRewardOffersManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const TeeRewardOffersManagerGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (TeeRewardOffersManagerGovernanceCallTimelocked) ContractEventName() string {
	return TeeRewardOffersManagerGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*TeeRewardOffersManagerGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != teeRewardOffersManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeeRewardOffersManagerGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teeRewardOffersManager.abi.Events[event].Inputs {
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

// TeeRewardOffersManagerGovernanceInitialised represents a GovernanceInitialised event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const TeeRewardOffersManagerGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (TeeRewardOffersManagerGovernanceInitialised) ContractEventName() string {
	return TeeRewardOffersManagerGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackGovernanceInitialisedEvent(log *types.Log) (*TeeRewardOffersManagerGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != teeRewardOffersManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeeRewardOffersManagerGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teeRewardOffersManager.abi.Events[event].Inputs {
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

// TeeRewardOffersManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const TeeRewardOffersManagerGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (TeeRewardOffersManagerGovernedProductionModeEntered) ContractEventName() string {
	return TeeRewardOffersManagerGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*TeeRewardOffersManagerGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != teeRewardOffersManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeeRewardOffersManagerGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teeRewardOffersManager.abi.Events[event].Inputs {
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

// TeeRewardOffersManagerInflationReceived represents a InflationReceived event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerInflationReceived struct {
	AmountReceivedWei *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const TeeRewardOffersManagerInflationReceivedEventName = "InflationReceived"

// ContractEventName returns the user-defined event name.
func (TeeRewardOffersManagerInflationReceived) ContractEventName() string {
	return TeeRewardOffersManagerInflationReceivedEventName
}

// UnpackInflationReceivedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InflationReceived(uint256 amountReceivedWei)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackInflationReceivedEvent(log *types.Log) (*TeeRewardOffersManagerInflationReceived, error) {
	event := "InflationReceived"
	if len(log.Topics) == 0 || log.Topics[0] != teeRewardOffersManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeeRewardOffersManagerInflationReceived)
	if len(log.Data) > 0 {
		if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teeRewardOffersManager.abi.Events[event].Inputs {
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

// TeeRewardOffersManagerInflationRewardsOffered represents a InflationRewardsOffered event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerInflationRewardsOffered struct {
	RewardEpochId *big.Int
	Amount        *big.Int
	TeeOwnersPPM  *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const TeeRewardOffersManagerInflationRewardsOfferedEventName = "InflationRewardsOffered"

// ContractEventName returns the user-defined event name.
func (TeeRewardOffersManagerInflationRewardsOffered) ContractEventName() string {
	return TeeRewardOffersManagerInflationRewardsOfferedEventName
}

// UnpackInflationRewardsOfferedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InflationRewardsOffered(uint24 indexed rewardEpochId, uint256 amount, uint256 teeOwnersPPM)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackInflationRewardsOfferedEvent(log *types.Log) (*TeeRewardOffersManagerInflationRewardsOffered, error) {
	event := "InflationRewardsOffered"
	if len(log.Topics) == 0 || log.Topics[0] != teeRewardOffersManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeeRewardOffersManagerInflationRewardsOffered)
	if len(log.Data) > 0 {
		if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teeRewardOffersManager.abi.Events[event].Inputs {
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

// TeeRewardOffersManagerInitialized represents a Initialized event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const TeeRewardOffersManagerInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (TeeRewardOffersManagerInitialized) ContractEventName() string {
	return TeeRewardOffersManagerInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackInitializedEvent(log *types.Log) (*TeeRewardOffersManagerInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != teeRewardOffersManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeeRewardOffersManagerInitialized)
	if len(log.Data) > 0 {
		if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teeRewardOffersManager.abi.Events[event].Inputs {
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

// TeeRewardOffersManagerTeeOwnersPPMSet represents a TeeOwnersPPMSet event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerTeeOwnersPPMSet struct {
	TeeOwnersPPM *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const TeeRewardOffersManagerTeeOwnersPPMSetEventName = "TeeOwnersPPMSet"

// ContractEventName returns the user-defined event name.
func (TeeRewardOffersManagerTeeOwnersPPMSet) ContractEventName() string {
	return TeeRewardOffersManagerTeeOwnersPPMSetEventName
}

// UnpackTeeOwnersPPMSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeOwnersPPMSet(uint24 teeOwnersPPM)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackTeeOwnersPPMSetEvent(log *types.Log) (*TeeRewardOffersManagerTeeOwnersPPMSet, error) {
	event := "TeeOwnersPPMSet"
	if len(log.Topics) == 0 || log.Topics[0] != teeRewardOffersManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeeRewardOffersManagerTeeOwnersPPMSet)
	if len(log.Data) > 0 {
		if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teeRewardOffersManager.abi.Events[event].Inputs {
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

// TeeRewardOffersManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeeRewardOffersManagerTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (TeeRewardOffersManagerTimelockedGovernanceCallCanceled) ContractEventName() string {
	return TeeRewardOffersManagerTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*TeeRewardOffersManagerTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != teeRewardOffersManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeeRewardOffersManagerTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teeRewardOffersManager.abi.Events[event].Inputs {
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

// TeeRewardOffersManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeeRewardOffersManagerTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (TeeRewardOffersManagerTimelockedGovernanceCallExecuted) ContractEventName() string {
	return TeeRewardOffersManagerTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*TeeRewardOffersManagerTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != teeRewardOffersManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeeRewardOffersManagerTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teeRewardOffersManager.abi.Events[event].Inputs {
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

// TeeRewardOffersManagerUpgraded represents a Upgraded event raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const TeeRewardOffersManagerUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (TeeRewardOffersManagerUpgraded) ContractEventName() string {
	return TeeRewardOffersManagerUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackUpgradedEvent(log *types.Log) (*TeeRewardOffersManagerUpgraded, error) {
	event := "Upgraded"
	if len(log.Topics) == 0 || log.Topics[0] != teeRewardOffersManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeeRewardOffersManagerUpgraded)
	if len(log.Data) > 0 {
		if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teeRewardOffersManager.abi.Events[event].Inputs {
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

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["AlreadyInProductionMode"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackAlreadyInProductionModeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["GovernedAddressZero"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackGovernedAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["GovernedAlreadyInitialized"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackGovernedAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["InvalidTeeOwnersPPMValue"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackInvalidTeeOwnersPPMValueError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["OnlyExecutor"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["TimelockCallNotFound"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackTimelockCallNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["TimelockNotAllowedYet"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackTimelockNotAllowedYetError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], teeRewardOffersManager.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return teeRewardOffersManager.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// TeeRewardOffersManagerAddressEmptyCode represents a AddressEmptyCode error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func TeeRewardOffersManagerAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackAddressEmptyCodeError(raw []byte) (*TeeRewardOffersManagerAddressEmptyCode, error) {
	out := new(TeeRewardOffersManagerAddressEmptyCode)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerAlreadyInProductionMode represents a AlreadyInProductionMode error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerAlreadyInProductionMode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInProductionMode()
func TeeRewardOffersManagerAlreadyInProductionModeErrorID() common.Hash {
	return common.HexToHash("0x72a23ec2b1fdbc3de7d5f411b572fcf8ad8409803196930f63ddf40a2a2b2594")
}

// UnpackAlreadyInProductionModeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInProductionMode()
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackAlreadyInProductionModeError(raw []byte) (*TeeRewardOffersManagerAlreadyInProductionMode, error) {
	out := new(TeeRewardOffersManagerAlreadyInProductionMode)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "AlreadyInProductionMode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func TeeRewardOffersManagerERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackERC1967InvalidImplementationError(raw []byte) (*TeeRewardOffersManagerERC1967InvalidImplementation, error) {
	out := new(TeeRewardOffersManagerERC1967InvalidImplementation)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerERC1967NonPayable represents a ERC1967NonPayable error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func TeeRewardOffersManagerERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackERC1967NonPayableError(raw []byte) (*TeeRewardOffersManagerERC1967NonPayable, error) {
	out := new(TeeRewardOffersManagerERC1967NonPayable)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerFailedCall represents a FailedCall error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func TeeRewardOffersManagerFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackFailedCallError(raw []byte) (*TeeRewardOffersManagerFailedCall, error) {
	out := new(TeeRewardOffersManagerFailedCall)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerGovernedAddressZero represents a GovernedAddressZero error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerGovernedAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAddressZero()
func TeeRewardOffersManagerGovernedAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x695725f6b429bac0cc5a2fa88214e31520e0c69f2fb0e2fb881a7130a564f49f")
}

// UnpackGovernedAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAddressZero()
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackGovernedAddressZeroError(raw []byte) (*TeeRewardOffersManagerGovernedAddressZero, error) {
	out := new(TeeRewardOffersManagerGovernedAddressZero)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "GovernedAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerGovernedAlreadyInitialized represents a GovernedAlreadyInitialized error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerGovernedAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAlreadyInitialized()
func TeeRewardOffersManagerGovernedAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0xb2f0162156d83828197776d56bbfba0a7a9d5936fad70f0a3014ae18cc7aa33c")
}

// UnpackGovernedAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAlreadyInitialized()
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackGovernedAlreadyInitializedError(raw []byte) (*TeeRewardOffersManagerGovernedAlreadyInitialized, error) {
	out := new(TeeRewardOffersManagerGovernedAlreadyInitialized)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "GovernedAlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerInvalidInitialization represents a InvalidInitialization error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func TeeRewardOffersManagerInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackInvalidInitializationError(raw []byte) (*TeeRewardOffersManagerInvalidInitialization, error) {
	out := new(TeeRewardOffersManagerInvalidInitialization)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerInvalidTeeOwnersPPMValue represents a InvalidTeeOwnersPPMValue error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerInvalidTeeOwnersPPMValue struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidTeeOwnersPPMValue()
func TeeRewardOffersManagerInvalidTeeOwnersPPMValueErrorID() common.Hash {
	return common.HexToHash("0x226c1ea6d6dfa92af9da3ae1b55e740dea6476cda74e9d00f9145ed6e267a1bf")
}

// UnpackInvalidTeeOwnersPPMValueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidTeeOwnersPPMValue()
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackInvalidTeeOwnersPPMValueError(raw []byte) (*TeeRewardOffersManagerInvalidTeeOwnersPPMValue, error) {
	out := new(TeeRewardOffersManagerInvalidTeeOwnersPPMValue)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "InvalidTeeOwnersPPMValue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerNotInitializing represents a NotInitializing error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func TeeRewardOffersManagerNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackNotInitializingError(raw []byte) (*TeeRewardOffersManagerNotInitializing, error) {
	out := new(TeeRewardOffersManagerNotInitializing)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerOnlyExecutor represents a OnlyExecutor error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerOnlyExecutor struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExecutor()
func TeeRewardOffersManagerOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x7fb6be02aaf40e420b0e7eb67e43a3feee949e8932a3c64e33b591a436670841")
}

// UnpackOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExecutor()
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackOnlyExecutorError(raw []byte) (*TeeRewardOffersManagerOnlyExecutor, error) {
	out := new(TeeRewardOffersManagerOnlyExecutor)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "OnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerOnlyGovernance represents a OnlyGovernance error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func TeeRewardOffersManagerOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackOnlyGovernanceError(raw []byte) (*TeeRewardOffersManagerOnlyGovernance, error) {
	out := new(TeeRewardOffersManagerOnlyGovernance)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerTimelockCallNotFound represents a TimelockCallNotFound error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerTimelockCallNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockCallNotFound()
func TeeRewardOffersManagerTimelockCallNotFoundErrorID() common.Hash {
	return common.HexToHash("0x262ac987dd09120380d393a0c52cdef1ca58949d1107dc26d8da9f7b527eb33d")
}

// UnpackTimelockCallNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockCallNotFound()
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackTimelockCallNotFoundError(raw []byte) (*TeeRewardOffersManagerTimelockCallNotFound, error) {
	out := new(TeeRewardOffersManagerTimelockCallNotFound)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "TimelockCallNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerTimelockNotAllowedYet represents a TimelockNotAllowedYet error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerTimelockNotAllowedYet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockNotAllowedYet()
func TeeRewardOffersManagerTimelockNotAllowedYetErrorID() common.Hash {
	return common.HexToHash("0x6124e5c2c545bec1d02048c68500eebdabeae0cc69995a629ebccf8a62f69083")
}

// UnpackTimelockNotAllowedYetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockNotAllowedYet()
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackTimelockNotAllowedYetError(raw []byte) (*TeeRewardOffersManagerTimelockNotAllowedYet, error) {
	out := new(TeeRewardOffersManagerTimelockNotAllowedYet)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "TimelockNotAllowedYet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func TeeRewardOffersManagerUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*TeeRewardOffersManagerUUPSUnauthorizedCallContext, error) {
	out := new(TeeRewardOffersManagerUUPSUnauthorizedCallContext)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeeRewardOffersManagerUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the TeeRewardOffersManager contract.
type TeeRewardOffersManagerUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func TeeRewardOffersManagerUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (teeRewardOffersManager *TeeRewardOffersManager) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*TeeRewardOffersManagerUUPSUnsupportedProxiableUUID, error) {
	out := new(TeeRewardOffersManagerUUPSUnsupportedProxiableUUID)
	if err := teeRewardOffersManager.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
