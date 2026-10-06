// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fdc2hub

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

// IFdc2HubFdc2AttestationRequest is an auto generated low-level Go binding around an user-defined struct.
type IFdc2HubFdc2AttestationRequest struct {
	Header      IFdc2HubFdc2RequestHeader
	RequestBody []byte
}

// IFdc2HubFdc2RequestHeader is an auto generated low-level Go binding around an user-defined struct.
type IFdc2HubFdc2RequestHeader struct {
	AttestationType [32]byte
	SourceId        [32]byte
	ThresholdBIPS   uint16
	ProofOwner      common.Address
}

// Fdc2HubMetaData contains all meta data concerning the Fdc2Hub contract.
var Fdc2HubMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DefaultNumberOfTeesZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"name\":\"DuplicatedTeeId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinThresholdInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MultipleResponsesPossible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NumberOfTeesAndTeeIdsInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockCallNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"AttestationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"defaultNumberOfTees\",\"type\":\"uint8\"}],\"name\":\"DefaultNumberOfTeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minThresholdBIPS\",\"type\":\"uint16\"}],\"name\":\"MinThresholdBIPSSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultNumberOfTees\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2RequestFeeConfigurations\",\"outputs\":[{\"internalType\":\"contractIFdc2RequestFeeConfigurations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"contractIIFlareTeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_minThresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"_defaultNumberOfTees\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minThresholdBIPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"}],\"internalType\":\"structIFdc2Hub.Fdc2RequestHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"requestBody\",\"type\":\"bytes\"}],\"internalType\":\"structIFdc2Hub.Fdc2AttestationRequest\",\"name\":\"_attestationRequest\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfTees\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_defaultNumberOfTees\",\"type\":\"uint8\"}],\"name\":\"setDefaultNumberOfTees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_minThresholdBIPS\",\"type\":\"uint16\"}],\"name\":\"setMinThresholdBIPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "Fdc2Hub",
}

// Fdc2Hub is an auto generated Go binding around an Ethereum contract.
type Fdc2Hub struct {
	abi abi.ABI
}

// NewFdc2Hub creates a new instance of Fdc2Hub.
func NewFdc2Hub() *Fdc2Hub {
	parsed, err := Fdc2HubMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Fdc2Hub{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Fdc2Hub) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (fdc2Hub *Fdc2Hub) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := fdc2Hub.abi.Pack("UPGRADE_INTERFACE_VERSION")
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
func (fdc2Hub *Fdc2Hub) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return fdc2Hub.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (fdc2Hub *Fdc2Hub) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := fdc2Hub.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
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
func (fdc2Hub *Fdc2Hub) PackCancelGovernanceCall(encodedCall []byte) []byte {
	enc, err := fdc2Hub.abi.Pack("cancelGovernanceCall", encodedCall)
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
func (fdc2Hub *Fdc2Hub) TryPackCancelGovernanceCall(encodedCall []byte) ([]byte, error) {
	return fdc2Hub.abi.Pack("cancelGovernanceCall", encodedCall)
}

// PackDefaultNumberOfTees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1b92dbb4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function defaultNumberOfTees() view returns(uint8)
func (fdc2Hub *Fdc2Hub) PackDefaultNumberOfTees() []byte {
	enc, err := fdc2Hub.abi.Pack("defaultNumberOfTees")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDefaultNumberOfTees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1b92dbb4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function defaultNumberOfTees() view returns(uint8)
func (fdc2Hub *Fdc2Hub) TryPackDefaultNumberOfTees() ([]byte, error) {
	return fdc2Hub.abi.Pack("defaultNumberOfTees")
}

// UnpackDefaultNumberOfTees is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1b92dbb4.
//
// Solidity: function defaultNumberOfTees() view returns(uint8)
func (fdc2Hub *Fdc2Hub) UnpackDefaultNumberOfTees(data []byte) (uint8, error) {
	out, err := fdc2Hub.abi.Unpack("defaultNumberOfTees", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (fdc2Hub *Fdc2Hub) PackExecuteGovernanceCall(encodedCall []byte) []byte {
	enc, err := fdc2Hub.abi.Pack("executeGovernanceCall", encodedCall)
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
func (fdc2Hub *Fdc2Hub) TryPackExecuteGovernanceCall(encodedCall []byte) ([]byte, error) {
	return fdc2Hub.abi.Pack("executeGovernanceCall", encodedCall)
}

// PackFdc2RequestFeeConfigurations is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x285434a2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fdc2RequestFeeConfigurations() view returns(address)
func (fdc2Hub *Fdc2Hub) PackFdc2RequestFeeConfigurations() []byte {
	enc, err := fdc2Hub.abi.Pack("fdc2RequestFeeConfigurations")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFdc2RequestFeeConfigurations is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x285434a2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fdc2RequestFeeConfigurations() view returns(address)
func (fdc2Hub *Fdc2Hub) TryPackFdc2RequestFeeConfigurations() ([]byte, error) {
	return fdc2Hub.abi.Pack("fdc2RequestFeeConfigurations")
}

// UnpackFdc2RequestFeeConfigurations is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x285434a2.
//
// Solidity: function fdc2RequestFeeConfigurations() view returns(address)
func (fdc2Hub *Fdc2Hub) UnpackFdc2RequestFeeConfigurations(data []byte) (common.Address, error) {
	out, err := fdc2Hub.abi.Unpack("fdc2RequestFeeConfigurations", data)
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
func (fdc2Hub *Fdc2Hub) PackFlareSystemsManager() []byte {
	enc, err := fdc2Hub.abi.Pack("flareSystemsManager")
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
func (fdc2Hub *Fdc2Hub) TryPackFlareSystemsManager() ([]byte, error) {
	return fdc2Hub.abi.Pack("flareSystemsManager")
}

// UnpackFlareSystemsManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (fdc2Hub *Fdc2Hub) UnpackFlareSystemsManager(data []byte) (common.Address, error) {
	out, err := fdc2Hub.abi.Unpack("flareSystemsManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFlareTeeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x453f7ab4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareTeeManager() view returns(address)
func (fdc2Hub *Fdc2Hub) PackFlareTeeManager() []byte {
	enc, err := fdc2Hub.abi.Pack("flareTeeManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFlareTeeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x453f7ab4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function flareTeeManager() view returns(address)
func (fdc2Hub *Fdc2Hub) TryPackFlareTeeManager() ([]byte, error) {
	return fdc2Hub.abi.Pack("flareTeeManager")
}

// UnpackFlareTeeManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (fdc2Hub *Fdc2Hub) UnpackFlareTeeManager(data []byte) (common.Address, error) {
	out, err := fdc2Hub.abi.Unpack("flareTeeManager", data)
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
func (fdc2Hub *Fdc2Hub) PackGetAddressUpdater() []byte {
	enc, err := fdc2Hub.abi.Pack("getAddressUpdater")
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
func (fdc2Hub *Fdc2Hub) TryPackGetAddressUpdater() ([]byte, error) {
	return fdc2Hub.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (fdc2Hub *Fdc2Hub) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := fdc2Hub.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governance() view returns(address)
func (fdc2Hub *Fdc2Hub) PackGovernance() []byte {
	enc, err := fdc2Hub.abi.Pack("governance")
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
func (fdc2Hub *Fdc2Hub) TryPackGovernance() ([]byte, error) {
	return fdc2Hub.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (fdc2Hub *Fdc2Hub) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := fdc2Hub.abi.Unpack("governance", data)
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
func (fdc2Hub *Fdc2Hub) PackGovernanceSettings() []byte {
	enc, err := fdc2Hub.abi.Pack("governanceSettings")
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
func (fdc2Hub *Fdc2Hub) TryPackGovernanceSettings() ([]byte, error) {
	return fdc2Hub.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (fdc2Hub *Fdc2Hub) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := fdc2Hub.abi.Unpack("governanceSettings", data)
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
func (fdc2Hub *Fdc2Hub) PackImplementation() []byte {
	enc, err := fdc2Hub.abi.Pack("implementation")
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
func (fdc2Hub *Fdc2Hub) TryPackImplementation() ([]byte, error) {
	return fdc2Hub.abi.Pack("implementation")
}

// UnpackImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (fdc2Hub *Fdc2Hub) UnpackImplementation(data []byte) (common.Address, error) {
	out, err := fdc2Hub.abi.Unpack("implementation", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd6b35a06.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint16 _minThresholdBIPS, uint8 _defaultNumberOfTees) returns()
func (fdc2Hub *Fdc2Hub) PackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address, minThresholdBIPS uint16, defaultNumberOfTees uint8) []byte {
	enc, err := fdc2Hub.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater, minThresholdBIPS, defaultNumberOfTees)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd6b35a06.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint16 _minThresholdBIPS, uint8 _defaultNumberOfTees) returns()
func (fdc2Hub *Fdc2Hub) TryPackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address, minThresholdBIPS uint16, defaultNumberOfTees uint8) ([]byte, error) {
	return fdc2Hub.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater, minThresholdBIPS, defaultNumberOfTees)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (fdc2Hub *Fdc2Hub) PackIsExecutor(address common.Address) []byte {
	enc, err := fdc2Hub.abi.Pack("isExecutor", address)
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
func (fdc2Hub *Fdc2Hub) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return fdc2Hub.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (fdc2Hub *Fdc2Hub) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := fdc2Hub.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackMinThresholdBIPS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeabf5bac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function minThresholdBIPS() view returns(uint16)
func (fdc2Hub *Fdc2Hub) PackMinThresholdBIPS() []byte {
	enc, err := fdc2Hub.abi.Pack("minThresholdBIPS")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMinThresholdBIPS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeabf5bac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function minThresholdBIPS() view returns(uint16)
func (fdc2Hub *Fdc2Hub) TryPackMinThresholdBIPS() ([]byte, error) {
	return fdc2Hub.abi.Pack("minThresholdBIPS")
}

// UnpackMinThresholdBIPS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xeabf5bac.
//
// Solidity: function minThresholdBIPS() view returns(uint16)
func (fdc2Hub *Fdc2Hub) UnpackMinThresholdBIPS(data []byte) (uint16, error) {
	out, err := fdc2Hub.abi.Unpack("minThresholdBIPS", data)
	if err != nil {
		return *new(uint16), err
	}
	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	return out0, nil
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (fdc2Hub *Fdc2Hub) PackProductionMode() []byte {
	enc, err := fdc2Hub.abi.Pack("productionMode")
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
func (fdc2Hub *Fdc2Hub) TryPackProductionMode() ([]byte, error) {
	return fdc2Hub.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (fdc2Hub *Fdc2Hub) UnpackProductionMode(data []byte) (bool, error) {
	out, err := fdc2Hub.abi.Unpack("productionMode", data)
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
func (fdc2Hub *Fdc2Hub) PackProxiableUUID() []byte {
	enc, err := fdc2Hub.abi.Pack("proxiableUUID")
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
func (fdc2Hub *Fdc2Hub) TryPackProxiableUUID() ([]byte, error) {
	return fdc2Hub.abi.Pack("proxiableUUID")
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (fdc2Hub *Fdc2Hub) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := fdc2Hub.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackRequestAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1398dddb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function requestAttestation(((bytes32,bytes32,uint16,address),bytes) _attestationRequest, uint256 _numberOfTees, address[] _teeIds, address[] _cosigners, uint64 _cosignersThreshold, address _claimBackAddress) payable returns()
func (fdc2Hub *Fdc2Hub) PackRequestAttestation(attestationRequest IFdc2HubFdc2AttestationRequest, numberOfTees *big.Int, teeIds []common.Address, cosigners []common.Address, cosignersThreshold uint64, claimBackAddress common.Address) []byte {
	enc, err := fdc2Hub.abi.Pack("requestAttestation", attestationRequest, numberOfTees, teeIds, cosigners, cosignersThreshold, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRequestAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1398dddb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function requestAttestation(((bytes32,bytes32,uint16,address),bytes) _attestationRequest, uint256 _numberOfTees, address[] _teeIds, address[] _cosigners, uint64 _cosignersThreshold, address _claimBackAddress) payable returns()
func (fdc2Hub *Fdc2Hub) TryPackRequestAttestation(attestationRequest IFdc2HubFdc2AttestationRequest, numberOfTees *big.Int, teeIds []common.Address, cosigners []common.Address, cosignersThreshold uint64, claimBackAddress common.Address) ([]byte, error) {
	return fdc2Hub.abi.Pack("requestAttestation", attestationRequest, numberOfTees, teeIds, cosigners, cosignersThreshold, claimBackAddress)
}

// PackRewardManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f4ef8a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardManager() view returns(address)
func (fdc2Hub *Fdc2Hub) PackRewardManager() []byte {
	enc, err := fdc2Hub.abi.Pack("rewardManager")
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
func (fdc2Hub *Fdc2Hub) TryPackRewardManager() ([]byte, error) {
	return fdc2Hub.abi.Pack("rewardManager")
}

// UnpackRewardManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (fdc2Hub *Fdc2Hub) UnpackRewardManager(data []byte) (common.Address, error) {
	out, err := fdc2Hub.abi.Unpack("rewardManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetDefaultNumberOfTees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f46aa36.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setDefaultNumberOfTees(uint8 _defaultNumberOfTees) returns()
func (fdc2Hub *Fdc2Hub) PackSetDefaultNumberOfTees(defaultNumberOfTees uint8) []byte {
	enc, err := fdc2Hub.abi.Pack("setDefaultNumberOfTees", defaultNumberOfTees)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetDefaultNumberOfTees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f46aa36.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setDefaultNumberOfTees(uint8 _defaultNumberOfTees) returns()
func (fdc2Hub *Fdc2Hub) TryPackSetDefaultNumberOfTees(defaultNumberOfTees uint8) ([]byte, error) {
	return fdc2Hub.abi.Pack("setDefaultNumberOfTees", defaultNumberOfTees)
}

// PackSetMinThresholdBIPS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8b8e70f5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMinThresholdBIPS(uint16 _minThresholdBIPS) returns()
func (fdc2Hub *Fdc2Hub) PackSetMinThresholdBIPS(minThresholdBIPS uint16) []byte {
	enc, err := fdc2Hub.abi.Pack("setMinThresholdBIPS", minThresholdBIPS)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMinThresholdBIPS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8b8e70f5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMinThresholdBIPS(uint16 _minThresholdBIPS) returns()
func (fdc2Hub *Fdc2Hub) TryPackSetMinThresholdBIPS(minThresholdBIPS uint16) ([]byte, error) {
	return fdc2Hub.abi.Pack("setMinThresholdBIPS", minThresholdBIPS)
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (fdc2Hub *Fdc2Hub) PackSwitchToProductionMode() []byte {
	enc, err := fdc2Hub.abi.Pack("switchToProductionMode")
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
func (fdc2Hub *Fdc2Hub) TryPackSwitchToProductionMode() ([]byte, error) {
	return fdc2Hub.abi.Pack("switchToProductionMode")
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (fdc2Hub *Fdc2Hub) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := fdc2Hub.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (fdc2Hub *Fdc2Hub) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return fdc2Hub.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (fdc2Hub *Fdc2Hub) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := fdc2Hub.abi.Pack("upgradeToAndCall", newImplementation, data)
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
func (fdc2Hub *Fdc2Hub) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return fdc2Hub.abi.Pack("upgradeToAndCall", newImplementation, data)
}

// Fdc2HubAttestationRequested represents a AttestationRequested event raised by the Fdc2Hub contract.
type Fdc2HubAttestationRequested struct {
	InstructionId    [32]byte
	AttestationType  [32]byte
	SourceId         [32]byte
	ProofOwner       common.Address
	ClaimBackAddress common.Address
	Fee              *big.Int
	Raw              *types.Log // Blockchain specific contextual infos
}

const Fdc2HubAttestationRequestedEventName = "AttestationRequested"

// ContractEventName returns the user-defined event name.
func (Fdc2HubAttestationRequested) ContractEventName() string {
	return Fdc2HubAttestationRequestedEventName
}

// UnpackAttestationRequestedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AttestationRequested(bytes32 indexed instructionId, bytes32 indexed attestationType, bytes32 indexed sourceId, address proofOwner, address claimBackAddress, uint256 fee)
func (fdc2Hub *Fdc2Hub) UnpackAttestationRequestedEvent(log *types.Log) (*Fdc2HubAttestationRequested, error) {
	event := "AttestationRequested"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2Hub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2HubAttestationRequested)
	if len(log.Data) > 0 {
		if err := fdc2Hub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2Hub.abi.Events[event].Inputs {
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

// Fdc2HubDefaultNumberOfTeesSet represents a DefaultNumberOfTeesSet event raised by the Fdc2Hub contract.
type Fdc2HubDefaultNumberOfTeesSet struct {
	DefaultNumberOfTees uint8
	Raw                 *types.Log // Blockchain specific contextual infos
}

const Fdc2HubDefaultNumberOfTeesSetEventName = "DefaultNumberOfTeesSet"

// ContractEventName returns the user-defined event name.
func (Fdc2HubDefaultNumberOfTeesSet) ContractEventName() string {
	return Fdc2HubDefaultNumberOfTeesSetEventName
}

// UnpackDefaultNumberOfTeesSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultNumberOfTeesSet(uint8 defaultNumberOfTees)
func (fdc2Hub *Fdc2Hub) UnpackDefaultNumberOfTeesSetEvent(log *types.Log) (*Fdc2HubDefaultNumberOfTeesSet, error) {
	event := "DefaultNumberOfTeesSet"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2Hub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2HubDefaultNumberOfTeesSet)
	if len(log.Data) > 0 {
		if err := fdc2Hub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2Hub.abi.Events[event].Inputs {
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

// Fdc2HubGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Fdc2Hub contract.
type Fdc2HubGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const Fdc2HubGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (Fdc2HubGovernanceCallTimelocked) ContractEventName() string {
	return Fdc2HubGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (fdc2Hub *Fdc2Hub) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*Fdc2HubGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2Hub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2HubGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := fdc2Hub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2Hub.abi.Events[event].Inputs {
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

// Fdc2HubGovernanceInitialised represents a GovernanceInitialised event raised by the Fdc2Hub contract.
type Fdc2HubGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const Fdc2HubGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (Fdc2HubGovernanceInitialised) ContractEventName() string {
	return Fdc2HubGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (fdc2Hub *Fdc2Hub) UnpackGovernanceInitialisedEvent(log *types.Log) (*Fdc2HubGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2Hub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2HubGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := fdc2Hub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2Hub.abi.Events[event].Inputs {
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

// Fdc2HubGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the Fdc2Hub contract.
type Fdc2HubGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const Fdc2HubGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (Fdc2HubGovernedProductionModeEntered) ContractEventName() string {
	return Fdc2HubGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (fdc2Hub *Fdc2Hub) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*Fdc2HubGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2Hub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2HubGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := fdc2Hub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2Hub.abi.Events[event].Inputs {
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

// Fdc2HubInitialized represents a Initialized event raised by the Fdc2Hub contract.
type Fdc2HubInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const Fdc2HubInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (Fdc2HubInitialized) ContractEventName() string {
	return Fdc2HubInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (fdc2Hub *Fdc2Hub) UnpackInitializedEvent(log *types.Log) (*Fdc2HubInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2Hub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2HubInitialized)
	if len(log.Data) > 0 {
		if err := fdc2Hub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2Hub.abi.Events[event].Inputs {
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

// Fdc2HubMinThresholdBIPSSet represents a MinThresholdBIPSSet event raised by the Fdc2Hub contract.
type Fdc2HubMinThresholdBIPSSet struct {
	MinThresholdBIPS uint16
	Raw              *types.Log // Blockchain specific contextual infos
}

const Fdc2HubMinThresholdBIPSSetEventName = "MinThresholdBIPSSet"

// ContractEventName returns the user-defined event name.
func (Fdc2HubMinThresholdBIPSSet) ContractEventName() string {
	return Fdc2HubMinThresholdBIPSSetEventName
}

// UnpackMinThresholdBIPSSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MinThresholdBIPSSet(uint16 minThresholdBIPS)
func (fdc2Hub *Fdc2Hub) UnpackMinThresholdBIPSSetEvent(log *types.Log) (*Fdc2HubMinThresholdBIPSSet, error) {
	event := "MinThresholdBIPSSet"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2Hub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2HubMinThresholdBIPSSet)
	if len(log.Data) > 0 {
		if err := fdc2Hub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2Hub.abi.Events[event].Inputs {
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

// Fdc2HubTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the Fdc2Hub contract.
type Fdc2HubTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const Fdc2HubTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (Fdc2HubTimelockedGovernanceCallCanceled) ContractEventName() string {
	return Fdc2HubTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (fdc2Hub *Fdc2Hub) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*Fdc2HubTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2Hub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2HubTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := fdc2Hub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2Hub.abi.Events[event].Inputs {
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

// Fdc2HubTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the Fdc2Hub contract.
type Fdc2HubTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const Fdc2HubTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (Fdc2HubTimelockedGovernanceCallExecuted) ContractEventName() string {
	return Fdc2HubTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (fdc2Hub *Fdc2Hub) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*Fdc2HubTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2Hub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2HubTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := fdc2Hub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2Hub.abi.Events[event].Inputs {
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

// Fdc2HubUpgraded represents a Upgraded event raised by the Fdc2Hub contract.
type Fdc2HubUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const Fdc2HubUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (Fdc2HubUpgraded) ContractEventName() string {
	return Fdc2HubUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (fdc2Hub *Fdc2Hub) UnpackUpgradedEvent(log *types.Log) (*Fdc2HubUpgraded, error) {
	event := "Upgraded"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2Hub.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2HubUpgraded)
	if len(log.Data) > 0 {
		if err := fdc2Hub.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2Hub.abi.Events[event].Inputs {
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
func (fdc2Hub *Fdc2Hub) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["AlreadyInProductionMode"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackAlreadyInProductionModeError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["CosignersThresholdInvalid"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackCosignersThresholdInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["DefaultNumberOfTeesZero"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackDefaultNumberOfTeesZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["DuplicatedTeeId"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackDuplicatedTeeIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["FeeTooLow"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackFeeTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["GovernedAddressZero"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackGovernedAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["GovernedAlreadyInitialized"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackGovernedAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["MinThresholdInvalid"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackMinThresholdInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["MultipleResponsesPossible"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackMultipleResponsesPossibleError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["NumberOfTeesAndTeeIdsInvalid"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackNumberOfTeesAndTeeIdsInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["OnlyExecutor"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["OnlySystemExtensionId"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackOnlySystemExtensionIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["ThresholdInvalid"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackThresholdInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["TimelockCallNotFound"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackTimelockCallNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["TimelockNotAllowedYet"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackTimelockNotAllowedYetError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2Hub.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return fdc2Hub.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// Fdc2HubAddressEmptyCode represents a AddressEmptyCode error raised by the Fdc2Hub contract.
type Fdc2HubAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func Fdc2HubAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (fdc2Hub *Fdc2Hub) UnpackAddressEmptyCodeError(raw []byte) (*Fdc2HubAddressEmptyCode, error) {
	out := new(Fdc2HubAddressEmptyCode)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubAlreadyInProductionMode represents a AlreadyInProductionMode error raised by the Fdc2Hub contract.
type Fdc2HubAlreadyInProductionMode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInProductionMode()
func Fdc2HubAlreadyInProductionModeErrorID() common.Hash {
	return common.HexToHash("0x72a23ec2b1fdbc3de7d5f411b572fcf8ad8409803196930f63ddf40a2a2b2594")
}

// UnpackAlreadyInProductionModeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInProductionMode()
func (fdc2Hub *Fdc2Hub) UnpackAlreadyInProductionModeError(raw []byte) (*Fdc2HubAlreadyInProductionMode, error) {
	out := new(Fdc2HubAlreadyInProductionMode)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "AlreadyInProductionMode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubCosignersThresholdInvalid represents a CosignersThresholdInvalid error raised by the Fdc2Hub contract.
type Fdc2HubCosignersThresholdInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CosignersThresholdInvalid()
func Fdc2HubCosignersThresholdInvalidErrorID() common.Hash {
	return common.HexToHash("0x90cf003b7d8e363fd0e715b342c918e95f12cc504e7870ce696831c1680bc0f3")
}

// UnpackCosignersThresholdInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CosignersThresholdInvalid()
func (fdc2Hub *Fdc2Hub) UnpackCosignersThresholdInvalidError(raw []byte) (*Fdc2HubCosignersThresholdInvalid, error) {
	out := new(Fdc2HubCosignersThresholdInvalid)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "CosignersThresholdInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubDefaultNumberOfTeesZero represents a DefaultNumberOfTeesZero error raised by the Fdc2Hub contract.
type Fdc2HubDefaultNumberOfTeesZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DefaultNumberOfTeesZero()
func Fdc2HubDefaultNumberOfTeesZeroErrorID() common.Hash {
	return common.HexToHash("0xb3c6b18cab7179a0bd3cbf8bb13973182737192cf9944d5e9c7a1c4b4891089f")
}

// UnpackDefaultNumberOfTeesZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DefaultNumberOfTeesZero()
func (fdc2Hub *Fdc2Hub) UnpackDefaultNumberOfTeesZeroError(raw []byte) (*Fdc2HubDefaultNumberOfTeesZero, error) {
	out := new(Fdc2HubDefaultNumberOfTeesZero)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "DefaultNumberOfTeesZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubDuplicatedTeeId represents a DuplicatedTeeId error raised by the Fdc2Hub contract.
type Fdc2HubDuplicatedTeeId struct {
	TeeId common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedTeeId(address teeId)
func Fdc2HubDuplicatedTeeIdErrorID() common.Hash {
	return common.HexToHash("0x592160bfd0bde9fc197118ed960686b86f74f857f1f949bb4d91bb15c8cd0c98")
}

// UnpackDuplicatedTeeIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedTeeId(address teeId)
func (fdc2Hub *Fdc2Hub) UnpackDuplicatedTeeIdError(raw []byte) (*Fdc2HubDuplicatedTeeId, error) {
	out := new(Fdc2HubDuplicatedTeeId)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "DuplicatedTeeId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the Fdc2Hub contract.
type Fdc2HubERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func Fdc2HubERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (fdc2Hub *Fdc2Hub) UnpackERC1967InvalidImplementationError(raw []byte) (*Fdc2HubERC1967InvalidImplementation, error) {
	out := new(Fdc2HubERC1967InvalidImplementation)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubERC1967NonPayable represents a ERC1967NonPayable error raised by the Fdc2Hub contract.
type Fdc2HubERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func Fdc2HubERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (fdc2Hub *Fdc2Hub) UnpackERC1967NonPayableError(raw []byte) (*Fdc2HubERC1967NonPayable, error) {
	out := new(Fdc2HubERC1967NonPayable)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubFailedCall represents a FailedCall error raised by the Fdc2Hub contract.
type Fdc2HubFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func Fdc2HubFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (fdc2Hub *Fdc2Hub) UnpackFailedCallError(raw []byte) (*Fdc2HubFailedCall, error) {
	out := new(Fdc2HubFailedCall)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubFeeTooLow represents a FeeTooLow error raised by the Fdc2Hub contract.
type Fdc2HubFeeTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeTooLow()
func Fdc2HubFeeTooLowErrorID() common.Hash {
	return common.HexToHash("0x732f94137528ce257864c792a6de53ff98f1203051e0a27c454063f559602458")
}

// UnpackFeeTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeTooLow()
func (fdc2Hub *Fdc2Hub) UnpackFeeTooLowError(raw []byte) (*Fdc2HubFeeTooLow, error) {
	out := new(Fdc2HubFeeTooLow)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "FeeTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubGovernedAddressZero represents a GovernedAddressZero error raised by the Fdc2Hub contract.
type Fdc2HubGovernedAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAddressZero()
func Fdc2HubGovernedAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x695725f6b429bac0cc5a2fa88214e31520e0c69f2fb0e2fb881a7130a564f49f")
}

// UnpackGovernedAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAddressZero()
func (fdc2Hub *Fdc2Hub) UnpackGovernedAddressZeroError(raw []byte) (*Fdc2HubGovernedAddressZero, error) {
	out := new(Fdc2HubGovernedAddressZero)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "GovernedAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubGovernedAlreadyInitialized represents a GovernedAlreadyInitialized error raised by the Fdc2Hub contract.
type Fdc2HubGovernedAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAlreadyInitialized()
func Fdc2HubGovernedAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0xb2f0162156d83828197776d56bbfba0a7a9d5936fad70f0a3014ae18cc7aa33c")
}

// UnpackGovernedAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAlreadyInitialized()
func (fdc2Hub *Fdc2Hub) UnpackGovernedAlreadyInitializedError(raw []byte) (*Fdc2HubGovernedAlreadyInitialized, error) {
	out := new(Fdc2HubGovernedAlreadyInitialized)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "GovernedAlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubInvalidInitialization represents a InvalidInitialization error raised by the Fdc2Hub contract.
type Fdc2HubInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func Fdc2HubInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (fdc2Hub *Fdc2Hub) UnpackInvalidInitializationError(raw []byte) (*Fdc2HubInvalidInitialization, error) {
	out := new(Fdc2HubInvalidInitialization)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubMinThresholdInvalid represents a MinThresholdInvalid error raised by the Fdc2Hub contract.
type Fdc2HubMinThresholdInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MinThresholdInvalid()
func Fdc2HubMinThresholdInvalidErrorID() common.Hash {
	return common.HexToHash("0x403c5d42f4fbd428276d17dd1239935a43077e9263b008c176333fea69cdc2eb")
}

// UnpackMinThresholdInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MinThresholdInvalid()
func (fdc2Hub *Fdc2Hub) UnpackMinThresholdInvalidError(raw []byte) (*Fdc2HubMinThresholdInvalid, error) {
	out := new(Fdc2HubMinThresholdInvalid)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "MinThresholdInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubMultipleResponsesPossible represents a MultipleResponsesPossible error raised by the Fdc2Hub contract.
type Fdc2HubMultipleResponsesPossible struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MultipleResponsesPossible()
func Fdc2HubMultipleResponsesPossibleErrorID() common.Hash {
	return common.HexToHash("0x47de09652836543fdc64a541dd7d60af49ab938b25e92d597f5f42b33c2bdb90")
}

// UnpackMultipleResponsesPossibleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MultipleResponsesPossible()
func (fdc2Hub *Fdc2Hub) UnpackMultipleResponsesPossibleError(raw []byte) (*Fdc2HubMultipleResponsesPossible, error) {
	out := new(Fdc2HubMultipleResponsesPossible)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "MultipleResponsesPossible", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubNotInitializing represents a NotInitializing error raised by the Fdc2Hub contract.
type Fdc2HubNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func Fdc2HubNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (fdc2Hub *Fdc2Hub) UnpackNotInitializingError(raw []byte) (*Fdc2HubNotInitializing, error) {
	out := new(Fdc2HubNotInitializing)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubNumberOfTeesAndTeeIdsInvalid represents a NumberOfTeesAndTeeIdsInvalid error raised by the Fdc2Hub contract.
type Fdc2HubNumberOfTeesAndTeeIdsInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NumberOfTeesAndTeeIdsInvalid()
func Fdc2HubNumberOfTeesAndTeeIdsInvalidErrorID() common.Hash {
	return common.HexToHash("0xabe06402b264ff4212e22994f791f881d85d62694956f1c463307a7331fd627d")
}

// UnpackNumberOfTeesAndTeeIdsInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NumberOfTeesAndTeeIdsInvalid()
func (fdc2Hub *Fdc2Hub) UnpackNumberOfTeesAndTeeIdsInvalidError(raw []byte) (*Fdc2HubNumberOfTeesAndTeeIdsInvalid, error) {
	out := new(Fdc2HubNumberOfTeesAndTeeIdsInvalid)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "NumberOfTeesAndTeeIdsInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubOnlyExecutor represents a OnlyExecutor error raised by the Fdc2Hub contract.
type Fdc2HubOnlyExecutor struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExecutor()
func Fdc2HubOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x7fb6be02aaf40e420b0e7eb67e43a3feee949e8932a3c64e33b591a436670841")
}

// UnpackOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExecutor()
func (fdc2Hub *Fdc2Hub) UnpackOnlyExecutorError(raw []byte) (*Fdc2HubOnlyExecutor, error) {
	out := new(Fdc2HubOnlyExecutor)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "OnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubOnlyGovernance represents a OnlyGovernance error raised by the Fdc2Hub contract.
type Fdc2HubOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func Fdc2HubOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (fdc2Hub *Fdc2Hub) UnpackOnlyGovernanceError(raw []byte) (*Fdc2HubOnlyGovernance, error) {
	out := new(Fdc2HubOnlyGovernance)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubOnlySystemExtensionId represents a OnlySystemExtensionId error raised by the Fdc2Hub contract.
type Fdc2HubOnlySystemExtensionId struct {
	TeeId common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemExtensionId(address teeId)
func Fdc2HubOnlySystemExtensionIdErrorID() common.Hash {
	return common.HexToHash("0x5ad200f03eec3d14783c311ecea8604906fbedec26a70135b63c30fab23494b0")
}

// UnpackOnlySystemExtensionIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemExtensionId(address teeId)
func (fdc2Hub *Fdc2Hub) UnpackOnlySystemExtensionIdError(raw []byte) (*Fdc2HubOnlySystemExtensionId, error) {
	out := new(Fdc2HubOnlySystemExtensionId)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "OnlySystemExtensionId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the Fdc2Hub contract.
type Fdc2HubTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func Fdc2HubTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (fdc2Hub *Fdc2Hub) UnpackTeeMachineNotAvailableError(raw []byte) (*Fdc2HubTeeMachineNotAvailable, error) {
	out := new(Fdc2HubTeeMachineNotAvailable)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubThresholdInvalid represents a ThresholdInvalid error raised by the Fdc2Hub contract.
type Fdc2HubThresholdInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ThresholdInvalid()
func Fdc2HubThresholdInvalidErrorID() common.Hash {
	return common.HexToHash("0x9309a4f4719058160d4c096778e43007d7d5dee7df84e80429f28273ee7eb282")
}

// UnpackThresholdInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ThresholdInvalid()
func (fdc2Hub *Fdc2Hub) UnpackThresholdInvalidError(raw []byte) (*Fdc2HubThresholdInvalid, error) {
	out := new(Fdc2HubThresholdInvalid)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "ThresholdInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubTimelockCallNotFound represents a TimelockCallNotFound error raised by the Fdc2Hub contract.
type Fdc2HubTimelockCallNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockCallNotFound()
func Fdc2HubTimelockCallNotFoundErrorID() common.Hash {
	return common.HexToHash("0x262ac987dd09120380d393a0c52cdef1ca58949d1107dc26d8da9f7b527eb33d")
}

// UnpackTimelockCallNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockCallNotFound()
func (fdc2Hub *Fdc2Hub) UnpackTimelockCallNotFoundError(raw []byte) (*Fdc2HubTimelockCallNotFound, error) {
	out := new(Fdc2HubTimelockCallNotFound)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "TimelockCallNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubTimelockNotAllowedYet represents a TimelockNotAllowedYet error raised by the Fdc2Hub contract.
type Fdc2HubTimelockNotAllowedYet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockNotAllowedYet()
func Fdc2HubTimelockNotAllowedYetErrorID() common.Hash {
	return common.HexToHash("0x6124e5c2c545bec1d02048c68500eebdabeae0cc69995a629ebccf8a62f69083")
}

// UnpackTimelockNotAllowedYetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockNotAllowedYet()
func (fdc2Hub *Fdc2Hub) UnpackTimelockNotAllowedYetError(raw []byte) (*Fdc2HubTimelockNotAllowedYet, error) {
	out := new(Fdc2HubTimelockNotAllowedYet)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "TimelockNotAllowedYet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the Fdc2Hub contract.
type Fdc2HubUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func Fdc2HubUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (fdc2Hub *Fdc2Hub) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*Fdc2HubUUPSUnauthorizedCallContext, error) {
	out := new(Fdc2HubUUPSUnauthorizedCallContext)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2HubUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the Fdc2Hub contract.
type Fdc2HubUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func Fdc2HubUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (fdc2Hub *Fdc2Hub) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*Fdc2HubUUPSUnsupportedProxiableUUID, error) {
	out := new(Fdc2HubUUPSUnsupportedProxiableUUID)
	if err := fdc2Hub.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
