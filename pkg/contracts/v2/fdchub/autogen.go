// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fdchub

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

// IFdcInflationConfigurationsFdcConfiguration is an auto generated low-level Go binding around an user-defined struct.
type IFdcInflationConfigurationsFdcConfiguration struct {
	AttestationType      [32]byte
	Source               [32]byte
	InflationShare       *big.Int
	MinRequestsThreshold uint8
	Mode                 *big.Int
}

// FdcHubMetaData contains all meta data concerning the FdcHub contract.
var FdcHubMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_requestsOffsetSeconds\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"AttestationRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"authorizedAmountWei\",\"type\":\"uint256\"}],\"name\":\"DailyAuthorizedInflationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountReceivedWei\",\"type\":\"uint256\"}],\"name\":\"InflationReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"source\",\"type\":\"bytes32\"},{\"internalType\":\"uint24\",\"name\":\"inflationShare\",\"type\":\"uint24\"},{\"internalType\":\"uint8\",\"name\":\"minRequestsThreshold\",\"type\":\"uint8\"},{\"internalType\":\"uint224\",\"name\":\"mode\",\"type\":\"uint224\"}],\"indexed\":false,\"internalType\":\"structIFdcInflationConfigurations.FdcConfiguration[]\",\"name\":\"fdcConfigurations\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InflationRewardsOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"requestsOffsetSeconds\",\"type\":\"uint8\"}],\"name\":\"RequestsOffsetSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyAuthorizedInflation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdcInflationConfigurations\",\"outputs\":[{\"internalType\":\"contractIFdcInflationConfigurations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdcRequestFeeConfigurations\",\"outputs\":[{\"internalType\":\"contractIFdcRequestFeeConfigurations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInflationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenPoolSupplyData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockedFundsWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalInflationAuthorizedWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalClaimedWei\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInflationAuthorizationReceivedTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInflationReceivedTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveInflation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"requestAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestsOffsetSeconds\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_toAuthorizeWei\",\"type\":\"uint256\"}],\"name\":\"setDailyAuthorizedInflation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_requestsOffsetSeconds\",\"type\":\"uint8\"}],\"name\":\"setRequestsOffset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationAuthorizedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationReceivedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalInflationRewardsOfferedWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_currentRewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_currentRewardEpochExpectedEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardEpochDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"triggerRewardEpochSwitchover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "FdcHub",
}

// FdcHub is an auto generated Go binding around an Ethereum contract.
type FdcHub struct {
	abi abi.ABI
}

// NewFdcHub creates a new instance of FdcHub.
func NewFdcHub() *FdcHub {
	parsed, err := FdcHubMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &FdcHub{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *FdcHub) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint8 _requestsOffsetSeconds) returns()
func (fdcHub *FdcHub) PackConstructor(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _requestsOffsetSeconds uint8) []byte {
	enc, err := fdcHub.abi.Pack("", _governanceSettings, _initialGovernance, _addressUpdater, _requestsOffsetSeconds)
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
func (fdcHub *FdcHub) PackCancelGovernanceCall(selector [4]byte) []byte {
	enc, err := fdcHub.abi.Pack("cancelGovernanceCall", selector)
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
func (fdcHub *FdcHub) TryPackCancelGovernanceCall(selector [4]byte) ([]byte, error) {
	return fdcHub.abi.Pack("cancelGovernanceCall", selector)
}

// PackDailyAuthorizedInflation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x708e34ce.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (fdcHub *FdcHub) PackDailyAuthorizedInflation() []byte {
	enc, err := fdcHub.abi.Pack("dailyAuthorizedInflation")
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
func (fdcHub *FdcHub) TryPackDailyAuthorizedInflation() ([]byte, error) {
	return fdcHub.abi.Pack("dailyAuthorizedInflation")
}

// UnpackDailyAuthorizedInflation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x708e34ce.
//
// Solidity: function dailyAuthorizedInflation() view returns(uint256)
func (fdcHub *FdcHub) UnpackDailyAuthorizedInflation(data []byte) (*big.Int, error) {
	out, err := fdcHub.abi.Unpack("dailyAuthorizedInflation", data)
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
func (fdcHub *FdcHub) PackExecuteGovernanceCall(selector [4]byte) []byte {
	enc, err := fdcHub.abi.Pack("executeGovernanceCall", selector)
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
func (fdcHub *FdcHub) TryPackExecuteGovernanceCall(selector [4]byte) ([]byte, error) {
	return fdcHub.abi.Pack("executeGovernanceCall", selector)
}

// PackFdcInflationConfigurations is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4c5a1d28.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fdcInflationConfigurations() view returns(address)
func (fdcHub *FdcHub) PackFdcInflationConfigurations() []byte {
	enc, err := fdcHub.abi.Pack("fdcInflationConfigurations")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFdcInflationConfigurations is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4c5a1d28.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fdcInflationConfigurations() view returns(address)
func (fdcHub *FdcHub) TryPackFdcInflationConfigurations() ([]byte, error) {
	return fdcHub.abi.Pack("fdcInflationConfigurations")
}

// UnpackFdcInflationConfigurations is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4c5a1d28.
//
// Solidity: function fdcInflationConfigurations() view returns(address)
func (fdcHub *FdcHub) UnpackFdcInflationConfigurations(data []byte) (common.Address, error) {
	out, err := fdcHub.abi.Unpack("fdcInflationConfigurations", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFdcRequestFeeConfigurations is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x116ea702.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fdcRequestFeeConfigurations() view returns(address)
func (fdcHub *FdcHub) PackFdcRequestFeeConfigurations() []byte {
	enc, err := fdcHub.abi.Pack("fdcRequestFeeConfigurations")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFdcRequestFeeConfigurations is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x116ea702.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fdcRequestFeeConfigurations() view returns(address)
func (fdcHub *FdcHub) TryPackFdcRequestFeeConfigurations() ([]byte, error) {
	return fdcHub.abi.Pack("fdcRequestFeeConfigurations")
}

// UnpackFdcRequestFeeConfigurations is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x116ea702.
//
// Solidity: function fdcRequestFeeConfigurations() view returns(address)
func (fdcHub *FdcHub) UnpackFdcRequestFeeConfigurations(data []byte) (common.Address, error) {
	out, err := fdcHub.abi.Unpack("fdcRequestFeeConfigurations", data)
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
func (fdcHub *FdcHub) PackFlareSystemsManager() []byte {
	enc, err := fdcHub.abi.Pack("flareSystemsManager")
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
func (fdcHub *FdcHub) TryPackFlareSystemsManager() ([]byte, error) {
	return fdcHub.abi.Pack("flareSystemsManager")
}

// UnpackFlareSystemsManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (fdcHub *FdcHub) UnpackFlareSystemsManager(data []byte) (common.Address, error) {
	out, err := fdcHub.abi.Unpack("flareSystemsManager", data)
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
func (fdcHub *FdcHub) PackGetAddressUpdater() []byte {
	enc, err := fdcHub.abi.Pack("getAddressUpdater")
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
func (fdcHub *FdcHub) TryPackGetAddressUpdater() ([]byte, error) {
	return fdcHub.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (fdcHub *FdcHub) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := fdcHub.abi.Unpack("getAddressUpdater", data)
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
func (fdcHub *FdcHub) PackGetContractName() []byte {
	enc, err := fdcHub.abi.Pack("getContractName")
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
func (fdcHub *FdcHub) TryPackGetContractName() ([]byte, error) {
	return fdcHub.abi.Pack("getContractName")
}

// UnpackGetContractName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (fdcHub *FdcHub) UnpackGetContractName(data []byte) (string, error) {
	out, err := fdcHub.abi.Unpack("getContractName", data)
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
func (fdcHub *FdcHub) PackGetExpectedBalance() []byte {
	enc, err := fdcHub.abi.Pack("getExpectedBalance")
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
func (fdcHub *FdcHub) TryPackGetExpectedBalance() ([]byte, error) {
	return fdcHub.abi.Pack("getExpectedBalance")
}

// UnpackGetExpectedBalance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaf04cd3b.
//
// Solidity: function getExpectedBalance() view returns(uint256)
func (fdcHub *FdcHub) UnpackGetExpectedBalance(data []byte) (*big.Int, error) {
	out, err := fdcHub.abi.Unpack("getExpectedBalance", data)
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
func (fdcHub *FdcHub) PackGetInflationAddress() []byte {
	enc, err := fdcHub.abi.Pack("getInflationAddress")
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
func (fdcHub *FdcHub) TryPackGetInflationAddress() ([]byte, error) {
	return fdcHub.abi.Pack("getInflationAddress")
}

// UnpackGetInflationAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xed39d3f8.
//
// Solidity: function getInflationAddress() view returns(address)
func (fdcHub *FdcHub) UnpackGetInflationAddress(data []byte) (common.Address, error) {
	out, err := fdcHub.abi.Unpack("getInflationAddress", data)
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
func (fdcHub *FdcHub) PackGetTokenPoolSupplyData() []byte {
	enc, err := fdcHub.abi.Pack("getTokenPoolSupplyData")
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
func (fdcHub *FdcHub) TryPackGetTokenPoolSupplyData() ([]byte, error) {
	return fdcHub.abi.Pack("getTokenPoolSupplyData")
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
func (fdcHub *FdcHub) UnpackGetTokenPoolSupplyData(data []byte) (GetTokenPoolSupplyDataOutput, error) {
	out, err := fdcHub.abi.Unpack("getTokenPoolSupplyData", data)
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
func (fdcHub *FdcHub) PackGovernance() []byte {
	enc, err := fdcHub.abi.Pack("governance")
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
func (fdcHub *FdcHub) TryPackGovernance() ([]byte, error) {
	return fdcHub.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (fdcHub *FdcHub) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := fdcHub.abi.Unpack("governance", data)
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
func (fdcHub *FdcHub) PackGovernanceSettings() []byte {
	enc, err := fdcHub.abi.Pack("governanceSettings")
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
func (fdcHub *FdcHub) TryPackGovernanceSettings() ([]byte, error) {
	return fdcHub.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (fdcHub *FdcHub) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := fdcHub.abi.Unpack("governanceSettings", data)
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
func (fdcHub *FdcHub) PackInitialise(governanceSettings common.Address, initialGovernance common.Address) []byte {
	enc, err := fdcHub.abi.Pack("initialise", governanceSettings, initialGovernance)
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
func (fdcHub *FdcHub) TryPackInitialise(governanceSettings common.Address, initialGovernance common.Address) ([]byte, error) {
	return fdcHub.abi.Pack("initialise", governanceSettings, initialGovernance)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (fdcHub *FdcHub) PackIsExecutor(address common.Address) []byte {
	enc, err := fdcHub.abi.Pack("isExecutor", address)
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
func (fdcHub *FdcHub) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return fdcHub.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (fdcHub *FdcHub) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := fdcHub.abi.Unpack("isExecutor", data)
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
func (fdcHub *FdcHub) PackLastInflationAuthorizationReceivedTs() []byte {
	enc, err := fdcHub.abi.Pack("lastInflationAuthorizationReceivedTs")
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
func (fdcHub *FdcHub) TryPackLastInflationAuthorizationReceivedTs() ([]byte, error) {
	return fdcHub.abi.Pack("lastInflationAuthorizationReceivedTs")
}

// UnpackLastInflationAuthorizationReceivedTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x473252c4.
//
// Solidity: function lastInflationAuthorizationReceivedTs() view returns(uint256)
func (fdcHub *FdcHub) UnpackLastInflationAuthorizationReceivedTs(data []byte) (*big.Int, error) {
	out, err := fdcHub.abi.Unpack("lastInflationAuthorizationReceivedTs", data)
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
func (fdcHub *FdcHub) PackLastInflationReceivedTs() []byte {
	enc, err := fdcHub.abi.Pack("lastInflationReceivedTs")
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
func (fdcHub *FdcHub) TryPackLastInflationReceivedTs() ([]byte, error) {
	return fdcHub.abi.Pack("lastInflationReceivedTs")
}

// UnpackLastInflationReceivedTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x12afcf0b.
//
// Solidity: function lastInflationReceivedTs() view returns(uint256)
func (fdcHub *FdcHub) UnpackLastInflationReceivedTs(data []byte) (*big.Int, error) {
	out, err := fdcHub.abi.Unpack("lastInflationReceivedTs", data)
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
func (fdcHub *FdcHub) PackProductionMode() []byte {
	enc, err := fdcHub.abi.Pack("productionMode")
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
func (fdcHub *FdcHub) TryPackProductionMode() ([]byte, error) {
	return fdcHub.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (fdcHub *FdcHub) UnpackProductionMode(data []byte) (bool, error) {
	out, err := fdcHub.abi.Unpack("productionMode", data)
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
func (fdcHub *FdcHub) PackReceiveInflation() []byte {
	enc, err := fdcHub.abi.Pack("receiveInflation")
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
func (fdcHub *FdcHub) TryPackReceiveInflation() ([]byte, error) {
	return fdcHub.abi.Pack("receiveInflation")
}

// PackRequestAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6238f354.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function requestAttestation(bytes _data) payable returns()
func (fdcHub *FdcHub) PackRequestAttestation(data []byte) []byte {
	enc, err := fdcHub.abi.Pack("requestAttestation", data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRequestAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6238f354.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function requestAttestation(bytes _data) payable returns()
func (fdcHub *FdcHub) TryPackRequestAttestation(data []byte) ([]byte, error) {
	return fdcHub.abi.Pack("requestAttestation", data)
}

// PackRequestsOffsetSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x94d019f1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function requestsOffsetSeconds() view returns(uint8)
func (fdcHub *FdcHub) PackRequestsOffsetSeconds() []byte {
	enc, err := fdcHub.abi.Pack("requestsOffsetSeconds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRequestsOffsetSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x94d019f1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function requestsOffsetSeconds() view returns(uint8)
func (fdcHub *FdcHub) TryPackRequestsOffsetSeconds() ([]byte, error) {
	return fdcHub.abi.Pack("requestsOffsetSeconds")
}

// UnpackRequestsOffsetSeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x94d019f1.
//
// Solidity: function requestsOffsetSeconds() view returns(uint8)
func (fdcHub *FdcHub) UnpackRequestsOffsetSeconds(data []byte) (uint8, error) {
	out, err := fdcHub.abi.Unpack("requestsOffsetSeconds", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackRewardManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f4ef8a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardManager() view returns(address)
func (fdcHub *FdcHub) PackRewardManager() []byte {
	enc, err := fdcHub.abi.Pack("rewardManager")
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
func (fdcHub *FdcHub) TryPackRewardManager() ([]byte, error) {
	return fdcHub.abi.Pack("rewardManager")
}

// UnpackRewardManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (fdcHub *FdcHub) UnpackRewardManager(data []byte) (common.Address, error) {
	out, err := fdcHub.abi.Unpack("rewardManager", data)
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
func (fdcHub *FdcHub) PackSetDailyAuthorizedInflation(toAuthorizeWei *big.Int) []byte {
	enc, err := fdcHub.abi.Pack("setDailyAuthorizedInflation", toAuthorizeWei)
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
func (fdcHub *FdcHub) TryPackSetDailyAuthorizedInflation(toAuthorizeWei *big.Int) ([]byte, error) {
	return fdcHub.abi.Pack("setDailyAuthorizedInflation", toAuthorizeWei)
}

// PackSetRequestsOffset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbfda6086.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setRequestsOffset(uint8 _requestsOffsetSeconds) returns()
func (fdcHub *FdcHub) PackSetRequestsOffset(requestsOffsetSeconds uint8) []byte {
	enc, err := fdcHub.abi.Pack("setRequestsOffset", requestsOffsetSeconds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetRequestsOffset is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbfda6086.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setRequestsOffset(uint8 _requestsOffsetSeconds) returns()
func (fdcHub *FdcHub) TryPackSetRequestsOffset(requestsOffsetSeconds uint8) ([]byte, error) {
	return fdcHub.abi.Pack("setRequestsOffset", requestsOffsetSeconds)
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (fdcHub *FdcHub) PackSwitchToProductionMode() []byte {
	enc, err := fdcHub.abi.Pack("switchToProductionMode")
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
func (fdcHub *FdcHub) TryPackSwitchToProductionMode() ([]byte, error) {
	return fdcHub.abi.Pack("switchToProductionMode")
}

// PackTimelockedCalls is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74e6310e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (fdcHub *FdcHub) PackTimelockedCalls(selector [4]byte) []byte {
	enc, err := fdcHub.abi.Pack("timelockedCalls", selector)
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
func (fdcHub *FdcHub) TryPackTimelockedCalls(selector [4]byte) ([]byte, error) {
	return fdcHub.abi.Pack("timelockedCalls", selector)
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
func (fdcHub *FdcHub) UnpackTimelockedCalls(data []byte) (TimelockedCallsOutput, error) {
	out, err := fdcHub.abi.Unpack("timelockedCalls", data)
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
func (fdcHub *FdcHub) PackTotalInflationAuthorizedWei() []byte {
	enc, err := fdcHub.abi.Pack("totalInflationAuthorizedWei")
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
func (fdcHub *FdcHub) TryPackTotalInflationAuthorizedWei() ([]byte, error) {
	return fdcHub.abi.Pack("totalInflationAuthorizedWei")
}

// UnpackTotalInflationAuthorizedWei is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd0c1c393.
//
// Solidity: function totalInflationAuthorizedWei() view returns(uint256)
func (fdcHub *FdcHub) UnpackTotalInflationAuthorizedWei(data []byte) (*big.Int, error) {
	out, err := fdcHub.abi.Unpack("totalInflationAuthorizedWei", data)
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
func (fdcHub *FdcHub) PackTotalInflationReceivedWei() []byte {
	enc, err := fdcHub.abi.Pack("totalInflationReceivedWei")
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
func (fdcHub *FdcHub) TryPackTotalInflationReceivedWei() ([]byte, error) {
	return fdcHub.abi.Pack("totalInflationReceivedWei")
}

// UnpackTotalInflationReceivedWei is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa5555aea.
//
// Solidity: function totalInflationReceivedWei() view returns(uint256)
func (fdcHub *FdcHub) UnpackTotalInflationReceivedWei(data []byte) (*big.Int, error) {
	out, err := fdcHub.abi.Unpack("totalInflationReceivedWei", data)
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
func (fdcHub *FdcHub) PackTotalInflationRewardsOfferedWei() []byte {
	enc, err := fdcHub.abi.Pack("totalInflationRewardsOfferedWei")
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
func (fdcHub *FdcHub) TryPackTotalInflationRewardsOfferedWei() ([]byte, error) {
	return fdcHub.abi.Pack("totalInflationRewardsOfferedWei")
}

// UnpackTotalInflationRewardsOfferedWei is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbd76b69c.
//
// Solidity: function totalInflationRewardsOfferedWei() view returns(uint256)
func (fdcHub *FdcHub) UnpackTotalInflationRewardsOfferedWei(data []byte) (*big.Int, error) {
	out, err := fdcHub.abi.Unpack("totalInflationRewardsOfferedWei", data)
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
func (fdcHub *FdcHub) PackTriggerRewardEpochSwitchover(currentRewardEpochId *big.Int, currentRewardEpochExpectedEndTs uint64, rewardEpochDurationSeconds uint64) []byte {
	enc, err := fdcHub.abi.Pack("triggerRewardEpochSwitchover", currentRewardEpochId, currentRewardEpochExpectedEndTs, rewardEpochDurationSeconds)
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
func (fdcHub *FdcHub) TryPackTriggerRewardEpochSwitchover(currentRewardEpochId *big.Int, currentRewardEpochExpectedEndTs uint64, rewardEpochDurationSeconds uint64) ([]byte, error) {
	return fdcHub.abi.Pack("triggerRewardEpochSwitchover", currentRewardEpochId, currentRewardEpochExpectedEndTs, rewardEpochDurationSeconds)
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (fdcHub *FdcHub) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := fdcHub.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (fdcHub *FdcHub) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return fdcHub.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// FdcHubAttestationRequest represents a AttestationRequest event raised by the FdcHub contract.
type FdcHubAttestationRequest struct {
	Data []byte
	Fee  *big.Int
	Raw  *types.Log // Blockchain specific contextual infos
}

const FdcHubAttestationRequestEventName = "AttestationRequest"

// ContractEventName returns the user-defined event name.
func (FdcHubAttestationRequest) ContractEventName() string {
	return FdcHubAttestationRequestEventName
}

// UnpackAttestationRequestEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttestationRequest(bytes data, uint256 fee)
func (fdcHub *FdcHub) UnpackAttestationRequestEvent(log *types.Log) (*FdcHubAttestationRequest, error) {
	event := "AttestationRequest"
	if len(log.Topics) == 0 || log.Topics[0] != fdcHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FdcHubAttestationRequest)
	if len(log.Data) > 0 {
		if err := fdcHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdcHub.abi.Events[event].Inputs {
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

// FdcHubDailyAuthorizedInflationSet represents a DailyAuthorizedInflationSet event raised by the FdcHub contract.
type FdcHubDailyAuthorizedInflationSet struct {
	AuthorizedAmountWei *big.Int
	Raw                 *types.Log // Blockchain specific contextual infos
}

const FdcHubDailyAuthorizedInflationSetEventName = "DailyAuthorizedInflationSet"

// ContractEventName returns the user-defined event name.
func (FdcHubDailyAuthorizedInflationSet) ContractEventName() string {
	return FdcHubDailyAuthorizedInflationSetEventName
}

// UnpackDailyAuthorizedInflationSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DailyAuthorizedInflationSet(uint256 authorizedAmountWei)
func (fdcHub *FdcHub) UnpackDailyAuthorizedInflationSetEvent(log *types.Log) (*FdcHubDailyAuthorizedInflationSet, error) {
	event := "DailyAuthorizedInflationSet"
	if len(log.Topics) == 0 || log.Topics[0] != fdcHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FdcHubDailyAuthorizedInflationSet)
	if len(log.Data) > 0 {
		if err := fdcHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdcHub.abi.Events[event].Inputs {
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

// FdcHubGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FdcHub contract.
type FdcHubGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   *types.Log // Blockchain specific contextual infos
}

const FdcHubGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (FdcHubGovernanceCallTimelocked) ContractEventName() string {
	return FdcHubGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (fdcHub *FdcHub) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*FdcHubGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != fdcHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FdcHubGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := fdcHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdcHub.abi.Events[event].Inputs {
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

// FdcHubGovernanceInitialised represents a GovernanceInitialised event raised by the FdcHub contract.
type FdcHubGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const FdcHubGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (FdcHubGovernanceInitialised) ContractEventName() string {
	return FdcHubGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (fdcHub *FdcHub) UnpackGovernanceInitialisedEvent(log *types.Log) (*FdcHubGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != fdcHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FdcHubGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := fdcHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdcHub.abi.Events[event].Inputs {
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

// FdcHubGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FdcHub contract.
type FdcHubGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const FdcHubGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (FdcHubGovernedProductionModeEntered) ContractEventName() string {
	return FdcHubGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (fdcHub *FdcHub) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*FdcHubGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != fdcHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FdcHubGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := fdcHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdcHub.abi.Events[event].Inputs {
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

// FdcHubInflationReceived represents a InflationReceived event raised by the FdcHub contract.
type FdcHubInflationReceived struct {
	AmountReceivedWei *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const FdcHubInflationReceivedEventName = "InflationReceived"

// ContractEventName returns the user-defined event name.
func (FdcHubInflationReceived) ContractEventName() string {
	return FdcHubInflationReceivedEventName
}

// UnpackInflationReceivedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InflationReceived(uint256 amountReceivedWei)
func (fdcHub *FdcHub) UnpackInflationReceivedEvent(log *types.Log) (*FdcHubInflationReceived, error) {
	event := "InflationReceived"
	if len(log.Topics) == 0 || log.Topics[0] != fdcHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FdcHubInflationReceived)
	if len(log.Data) > 0 {
		if err := fdcHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdcHub.abi.Events[event].Inputs {
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

// FdcHubInflationRewardsOffered represents a InflationRewardsOffered event raised by the FdcHub contract.
type FdcHubInflationRewardsOffered struct {
	RewardEpochId     *big.Int
	FdcConfigurations []IFdcInflationConfigurationsFdcConfiguration
	Amount            *big.Int
	Raw               *types.Log // Blockchain specific contextual infos
}

const FdcHubInflationRewardsOfferedEventName = "InflationRewardsOffered"

// ContractEventName returns the user-defined event name.
func (FdcHubInflationRewardsOffered) ContractEventName() string {
	return FdcHubInflationRewardsOfferedEventName
}

// UnpackInflationRewardsOfferedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InflationRewardsOffered(uint24 indexed rewardEpochId, (bytes32,bytes32,uint24,uint8,uint224)[] fdcConfigurations, uint256 amount)
func (fdcHub *FdcHub) UnpackInflationRewardsOfferedEvent(log *types.Log) (*FdcHubInflationRewardsOffered, error) {
	event := "InflationRewardsOffered"
	if len(log.Topics) == 0 || log.Topics[0] != fdcHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FdcHubInflationRewardsOffered)
	if len(log.Data) > 0 {
		if err := fdcHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdcHub.abi.Events[event].Inputs {
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

// FdcHubRequestsOffsetSet represents a RequestsOffsetSet event raised by the FdcHub contract.
type FdcHubRequestsOffsetSet struct {
	RequestsOffsetSeconds uint8
	Raw                   *types.Log // Blockchain specific contextual infos
}

const FdcHubRequestsOffsetSetEventName = "RequestsOffsetSet"

// ContractEventName returns the user-defined event name.
func (FdcHubRequestsOffsetSet) ContractEventName() string {
	return FdcHubRequestsOffsetSetEventName
}

// UnpackRequestsOffsetSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RequestsOffsetSet(uint8 requestsOffsetSeconds)
func (fdcHub *FdcHub) UnpackRequestsOffsetSetEvent(log *types.Log) (*FdcHubRequestsOffsetSet, error) {
	event := "RequestsOffsetSet"
	if len(log.Topics) == 0 || log.Topics[0] != fdcHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FdcHubRequestsOffsetSet)
	if len(log.Data) > 0 {
		if err := fdcHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdcHub.abi.Events[event].Inputs {
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

// FdcHubTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FdcHub contract.
type FdcHubTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const FdcHubTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (FdcHubTimelockedGovernanceCallCanceled) ContractEventName() string {
	return FdcHubTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (fdcHub *FdcHub) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*FdcHubTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != fdcHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FdcHubTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := fdcHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdcHub.abi.Events[event].Inputs {
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

// FdcHubTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FdcHub contract.
type FdcHubTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const FdcHubTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (FdcHubTimelockedGovernanceCallExecuted) ContractEventName() string {
	return FdcHubTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (fdcHub *FdcHub) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*FdcHubTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != fdcHub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FdcHubTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := fdcHub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdcHub.abi.Events[event].Inputs {
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
