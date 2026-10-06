// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fdc2requestfeeconfigurations

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

// Fdc2RequestFeeConfigurationsMetaData contains all meta data concerning the Fdc2RequestFeeConfigurations contract.
var Fdc2RequestFeeConfigurationsMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeMustBeGreaterThanZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockCallNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TypeAndSourceCombinationNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"source\",\"type\":\"bytes32\"}],\"name\":\"TypeAndSourceFeeRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"source\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TypeAndSourceFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_typ\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_source\",\"type\":\"bytes32\"}],\"name\":\"getTypeAndSourceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_typ\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_source\",\"type\":\"bytes32\"}],\"name\":\"removeTypeAndSourceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_types\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sources\",\"type\":\"bytes32[]\"}],\"name\":\"removeTypeAndSourceFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_typ\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_source\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setTypeAndSourceFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_types\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_sources\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"}],\"name\":\"setTypeAndSourceFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "Fdc2RequestFeeConfigurations",
}

// Fdc2RequestFeeConfigurations is an auto generated Go binding around an Ethereum contract.
type Fdc2RequestFeeConfigurations struct {
	abi abi.ABI
}

// NewFdc2RequestFeeConfigurations creates a new instance of Fdc2RequestFeeConfigurations.
func NewFdc2RequestFeeConfigurations() *Fdc2RequestFeeConfigurations {
	parsed, err := Fdc2RequestFeeConfigurationsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Fdc2RequestFeeConfigurations{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Fdc2RequestFeeConfigurations) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("UPGRADE_INTERFACE_VERSION")
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := fdc2RequestFeeConfigurations.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackCancelGovernanceCall(encodedCall []byte) []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("cancelGovernanceCall", encodedCall)
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackCancelGovernanceCall(encodedCall []byte) ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("cancelGovernanceCall", encodedCall)
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackExecuteGovernanceCall(encodedCall []byte) []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("executeGovernanceCall", encodedCall)
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackExecuteGovernanceCall(encodedCall []byte) ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("executeGovernanceCall", encodedCall)
}

// PackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackGetAddressUpdater() []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("getAddressUpdater")
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackGetAddressUpdater() ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := fdc2RequestFeeConfigurations.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetTypeAndSourceFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ebe55b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTypeAndSourceFee(bytes32 _typ, bytes32 _source) view returns(uint256 _fee)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackGetTypeAndSourceFee(typ [32]byte, source [32]byte) []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("getTypeAndSourceFee", typ, source)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTypeAndSourceFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ebe55b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTypeAndSourceFee(bytes32 _typ, bytes32 _source) view returns(uint256 _fee)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackGetTypeAndSourceFee(typ [32]byte, source [32]byte) ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("getTypeAndSourceFee", typ, source)
}

// UnpackGetTypeAndSourceFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5ebe55b3.
//
// Solidity: function getTypeAndSourceFee(bytes32 _typ, bytes32 _source) view returns(uint256 _fee)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackGetTypeAndSourceFee(data []byte) (*big.Int, error) {
	out, err := fdc2RequestFeeConfigurations.abi.Unpack("getTypeAndSourceFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governance() view returns(address)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackGovernance() []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("governance")
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackGovernance() ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := fdc2RequestFeeConfigurations.abi.Unpack("governance", data)
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackGovernanceSettings() []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("governanceSettings")
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackGovernanceSettings() ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := fdc2RequestFeeConfigurations.abi.Unpack("governanceSettings", data)
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackImplementation() []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("implementation")
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackImplementation() ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("implementation")
}

// UnpackImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackImplementation(data []byte) (common.Address, error) {
	out, err := fdc2RequestFeeConfigurations.abi.Unpack("implementation", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc0c53b8b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc0c53b8b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackIsExecutor(address common.Address) []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("isExecutor", address)
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := fdc2RequestFeeConfigurations.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackProductionMode() []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("productionMode")
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackProductionMode() ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackProductionMode(data []byte) (bool, error) {
	out, err := fdc2RequestFeeConfigurations.abi.Unpack("productionMode", data)
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackProxiableUUID() []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("proxiableUUID")
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackProxiableUUID() ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("proxiableUUID")
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := fdc2RequestFeeConfigurations.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackRemoveTypeAndSourceFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda42a778.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeTypeAndSourceFee(bytes32 _typ, bytes32 _source) returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackRemoveTypeAndSourceFee(typ [32]byte, source [32]byte) []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("removeTypeAndSourceFee", typ, source)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveTypeAndSourceFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xda42a778.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeTypeAndSourceFee(bytes32 _typ, bytes32 _source) returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackRemoveTypeAndSourceFee(typ [32]byte, source [32]byte) ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("removeTypeAndSourceFee", typ, source)
}

// PackRemoveTypeAndSourceFees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2f6bca8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeTypeAndSourceFees(bytes32[] _types, bytes32[] _sources) returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackRemoveTypeAndSourceFees(types [][32]byte, sources [][32]byte) []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("removeTypeAndSourceFees", types, sources)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveTypeAndSourceFees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe2f6bca8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeTypeAndSourceFees(bytes32[] _types, bytes32[] _sources) returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackRemoveTypeAndSourceFees(types [][32]byte, sources [][32]byte) ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("removeTypeAndSourceFees", types, sources)
}

// PackSetTypeAndSourceFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b1bacb7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setTypeAndSourceFee(bytes32 _typ, bytes32 _source, uint256 _fee) returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackSetTypeAndSourceFee(typ [32]byte, source [32]byte, fee *big.Int) []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("setTypeAndSourceFee", typ, source, fee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetTypeAndSourceFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5b1bacb7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setTypeAndSourceFee(bytes32 _typ, bytes32 _source, uint256 _fee) returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackSetTypeAndSourceFee(typ [32]byte, source [32]byte, fee *big.Int) ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("setTypeAndSourceFee", typ, source, fee)
}

// PackSetTypeAndSourceFees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3da84157.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setTypeAndSourceFees(bytes32[] _types, bytes32[] _sources, uint256[] _fees) returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackSetTypeAndSourceFees(types [][32]byte, sources [][32]byte, fees []*big.Int) []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("setTypeAndSourceFees", types, sources, fees)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetTypeAndSourceFees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3da84157.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setTypeAndSourceFees(bytes32[] _types, bytes32[] _sources, uint256[] _fees) returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackSetTypeAndSourceFees(types [][32]byte, sources [][32]byte, fees []*big.Int) ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("setTypeAndSourceFees", types, sources, fees)
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackSwitchToProductionMode() []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("switchToProductionMode")
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackSwitchToProductionMode() ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("switchToProductionMode")
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := fdc2RequestFeeConfigurations.abi.Pack("upgradeToAndCall", newImplementation, data)
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return fdc2RequestFeeConfigurations.abi.Pack("upgradeToAndCall", newImplementation, data)
}

// Fdc2RequestFeeConfigurationsGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const Fdc2RequestFeeConfigurationsGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (Fdc2RequestFeeConfigurationsGovernanceCallTimelocked) ContractEventName() string {
	return Fdc2RequestFeeConfigurationsGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*Fdc2RequestFeeConfigurationsGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2RequestFeeConfigurations.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2RequestFeeConfigurationsGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2RequestFeeConfigurations.abi.Events[event].Inputs {
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

// Fdc2RequestFeeConfigurationsGovernanceInitialised represents a GovernanceInitialised event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const Fdc2RequestFeeConfigurationsGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (Fdc2RequestFeeConfigurationsGovernanceInitialised) ContractEventName() string {
	return Fdc2RequestFeeConfigurationsGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackGovernanceInitialisedEvent(log *types.Log) (*Fdc2RequestFeeConfigurationsGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2RequestFeeConfigurations.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2RequestFeeConfigurationsGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2RequestFeeConfigurations.abi.Events[event].Inputs {
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

// Fdc2RequestFeeConfigurationsGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const Fdc2RequestFeeConfigurationsGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (Fdc2RequestFeeConfigurationsGovernedProductionModeEntered) ContractEventName() string {
	return Fdc2RequestFeeConfigurationsGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*Fdc2RequestFeeConfigurationsGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2RequestFeeConfigurations.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2RequestFeeConfigurationsGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2RequestFeeConfigurations.abi.Events[event].Inputs {
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

// Fdc2RequestFeeConfigurationsInitialized represents a Initialized event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const Fdc2RequestFeeConfigurationsInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (Fdc2RequestFeeConfigurationsInitialized) ContractEventName() string {
	return Fdc2RequestFeeConfigurationsInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackInitializedEvent(log *types.Log) (*Fdc2RequestFeeConfigurationsInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2RequestFeeConfigurations.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2RequestFeeConfigurationsInitialized)
	if len(log.Data) > 0 {
		if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2RequestFeeConfigurations.abi.Events[event].Inputs {
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

// Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled) ContractEventName() string {
	return Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2RequestFeeConfigurations.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2RequestFeeConfigurationsTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2RequestFeeConfigurations.abi.Events[event].Inputs {
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

// Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted) ContractEventName() string {
	return Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2RequestFeeConfigurations.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2RequestFeeConfigurationsTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2RequestFeeConfigurations.abi.Events[event].Inputs {
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

// Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved represents a TypeAndSourceFeeRemoved event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved struct {
	AttestationType [32]byte
	Source          [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemovedEventName = "TypeAndSourceFeeRemoved"

// ContractEventName returns the user-defined event name.
func (Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved) ContractEventName() string {
	return Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemovedEventName
}

// UnpackTypeAndSourceFeeRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TypeAndSourceFeeRemoved(bytes32 indexed attestationType, bytes32 indexed source)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackTypeAndSourceFeeRemovedEvent(log *types.Log) (*Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved, error) {
	event := "TypeAndSourceFeeRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2RequestFeeConfigurations.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2RequestFeeConfigurationsTypeAndSourceFeeRemoved)
	if len(log.Data) > 0 {
		if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2RequestFeeConfigurations.abi.Events[event].Inputs {
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

// Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet represents a TypeAndSourceFeeSet event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet struct {
	AttestationType [32]byte
	Source          [32]byte
	Fee             *big.Int
	Raw             *types.Log // Blockchain specific contextual infos
}

const Fdc2RequestFeeConfigurationsTypeAndSourceFeeSetEventName = "TypeAndSourceFeeSet"

// ContractEventName returns the user-defined event name.
func (Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet) ContractEventName() string {
	return Fdc2RequestFeeConfigurationsTypeAndSourceFeeSetEventName
}

// UnpackTypeAndSourceFeeSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TypeAndSourceFeeSet(bytes32 indexed attestationType, bytes32 indexed source, uint256 fee)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackTypeAndSourceFeeSetEvent(log *types.Log) (*Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet, error) {
	event := "TypeAndSourceFeeSet"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2RequestFeeConfigurations.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2RequestFeeConfigurationsTypeAndSourceFeeSet)
	if len(log.Data) > 0 {
		if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2RequestFeeConfigurations.abi.Events[event].Inputs {
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

// Fdc2RequestFeeConfigurationsUpgraded represents a Upgraded event raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const Fdc2RequestFeeConfigurationsUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (Fdc2RequestFeeConfigurationsUpgraded) ContractEventName() string {
	return Fdc2RequestFeeConfigurationsUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackUpgradedEvent(log *types.Log) (*Fdc2RequestFeeConfigurationsUpgraded, error) {
	event := "Upgraded"
	if len(log.Topics) == 0 || log.Topics[0] != fdc2RequestFeeConfigurations.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(Fdc2RequestFeeConfigurationsUpgraded)
	if len(log.Data) > 0 {
		if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fdc2RequestFeeConfigurations.abi.Events[event].Inputs {
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
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["AlreadyInProductionMode"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackAlreadyInProductionModeError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["FeeMustBeGreaterThanZero"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackFeeMustBeGreaterThanZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["FeeNotSet"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackFeeNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["GovernedAddressZero"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackGovernedAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["GovernedAlreadyInitialized"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackGovernedAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["OnlyExecutor"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["TimelockCallNotFound"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackTimelockCallNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["TimelockNotAllowedYet"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackTimelockNotAllowedYetError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["TypeAndSourceCombinationNotSupported"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackTypeAndSourceCombinationNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], fdc2RequestFeeConfigurations.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return fdc2RequestFeeConfigurations.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// Fdc2RequestFeeConfigurationsAddressEmptyCode represents a AddressEmptyCode error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func Fdc2RequestFeeConfigurationsAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackAddressEmptyCodeError(raw []byte) (*Fdc2RequestFeeConfigurationsAddressEmptyCode, error) {
	out := new(Fdc2RequestFeeConfigurationsAddressEmptyCode)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsAlreadyInProductionMode represents a AlreadyInProductionMode error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsAlreadyInProductionMode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInProductionMode()
func Fdc2RequestFeeConfigurationsAlreadyInProductionModeErrorID() common.Hash {
	return common.HexToHash("0x72a23ec2b1fdbc3de7d5f411b572fcf8ad8409803196930f63ddf40a2a2b2594")
}

// UnpackAlreadyInProductionModeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInProductionMode()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackAlreadyInProductionModeError(raw []byte) (*Fdc2RequestFeeConfigurationsAlreadyInProductionMode, error) {
	out := new(Fdc2RequestFeeConfigurationsAlreadyInProductionMode)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "AlreadyInProductionMode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func Fdc2RequestFeeConfigurationsERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackERC1967InvalidImplementationError(raw []byte) (*Fdc2RequestFeeConfigurationsERC1967InvalidImplementation, error) {
	out := new(Fdc2RequestFeeConfigurationsERC1967InvalidImplementation)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsERC1967NonPayable represents a ERC1967NonPayable error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func Fdc2RequestFeeConfigurationsERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackERC1967NonPayableError(raw []byte) (*Fdc2RequestFeeConfigurationsERC1967NonPayable, error) {
	out := new(Fdc2RequestFeeConfigurationsERC1967NonPayable)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsFailedCall represents a FailedCall error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func Fdc2RequestFeeConfigurationsFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackFailedCallError(raw []byte) (*Fdc2RequestFeeConfigurationsFailedCall, error) {
	out := new(Fdc2RequestFeeConfigurationsFailedCall)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsFeeMustBeGreaterThanZero represents a FeeMustBeGreaterThanZero error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsFeeMustBeGreaterThanZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeMustBeGreaterThanZero()
func Fdc2RequestFeeConfigurationsFeeMustBeGreaterThanZeroErrorID() common.Hash {
	return common.HexToHash("0x3d85076c5d8fa2564d997d20c5015ec91bc8a10891ddccf10d93847192eabd23")
}

// UnpackFeeMustBeGreaterThanZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeMustBeGreaterThanZero()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackFeeMustBeGreaterThanZeroError(raw []byte) (*Fdc2RequestFeeConfigurationsFeeMustBeGreaterThanZero, error) {
	out := new(Fdc2RequestFeeConfigurationsFeeMustBeGreaterThanZero)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "FeeMustBeGreaterThanZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsFeeNotSet represents a FeeNotSet error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsFeeNotSet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeNotSet()
func Fdc2RequestFeeConfigurationsFeeNotSetErrorID() common.Hash {
	return common.HexToHash("0xda3424084cce5799e74f636863c1ae18dc226ce936b17767335a7fada011c395")
}

// UnpackFeeNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeNotSet()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackFeeNotSetError(raw []byte) (*Fdc2RequestFeeConfigurationsFeeNotSet, error) {
	out := new(Fdc2RequestFeeConfigurationsFeeNotSet)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "FeeNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsGovernedAddressZero represents a GovernedAddressZero error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsGovernedAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAddressZero()
func Fdc2RequestFeeConfigurationsGovernedAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x695725f6b429bac0cc5a2fa88214e31520e0c69f2fb0e2fb881a7130a564f49f")
}

// UnpackGovernedAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAddressZero()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackGovernedAddressZeroError(raw []byte) (*Fdc2RequestFeeConfigurationsGovernedAddressZero, error) {
	out := new(Fdc2RequestFeeConfigurationsGovernedAddressZero)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "GovernedAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsGovernedAlreadyInitialized represents a GovernedAlreadyInitialized error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsGovernedAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAlreadyInitialized()
func Fdc2RequestFeeConfigurationsGovernedAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0xb2f0162156d83828197776d56bbfba0a7a9d5936fad70f0a3014ae18cc7aa33c")
}

// UnpackGovernedAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAlreadyInitialized()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackGovernedAlreadyInitializedError(raw []byte) (*Fdc2RequestFeeConfigurationsGovernedAlreadyInitialized, error) {
	out := new(Fdc2RequestFeeConfigurationsGovernedAlreadyInitialized)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "GovernedAlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsInvalidInitialization represents a InvalidInitialization error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func Fdc2RequestFeeConfigurationsInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackInvalidInitializationError(raw []byte) (*Fdc2RequestFeeConfigurationsInvalidInitialization, error) {
	out := new(Fdc2RequestFeeConfigurationsInvalidInitialization)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsLengthsMismatch represents a LengthsMismatch error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func Fdc2RequestFeeConfigurationsLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackLengthsMismatchError(raw []byte) (*Fdc2RequestFeeConfigurationsLengthsMismatch, error) {
	out := new(Fdc2RequestFeeConfigurationsLengthsMismatch)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsNotInitializing represents a NotInitializing error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func Fdc2RequestFeeConfigurationsNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackNotInitializingError(raw []byte) (*Fdc2RequestFeeConfigurationsNotInitializing, error) {
	out := new(Fdc2RequestFeeConfigurationsNotInitializing)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsOnlyExecutor represents a OnlyExecutor error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsOnlyExecutor struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExecutor()
func Fdc2RequestFeeConfigurationsOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x7fb6be02aaf40e420b0e7eb67e43a3feee949e8932a3c64e33b591a436670841")
}

// UnpackOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExecutor()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackOnlyExecutorError(raw []byte) (*Fdc2RequestFeeConfigurationsOnlyExecutor, error) {
	out := new(Fdc2RequestFeeConfigurationsOnlyExecutor)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "OnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsOnlyGovernance represents a OnlyGovernance error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func Fdc2RequestFeeConfigurationsOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackOnlyGovernanceError(raw []byte) (*Fdc2RequestFeeConfigurationsOnlyGovernance, error) {
	out := new(Fdc2RequestFeeConfigurationsOnlyGovernance)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsTimelockCallNotFound represents a TimelockCallNotFound error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTimelockCallNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockCallNotFound()
func Fdc2RequestFeeConfigurationsTimelockCallNotFoundErrorID() common.Hash {
	return common.HexToHash("0x262ac987dd09120380d393a0c52cdef1ca58949d1107dc26d8da9f7b527eb33d")
}

// UnpackTimelockCallNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockCallNotFound()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackTimelockCallNotFoundError(raw []byte) (*Fdc2RequestFeeConfigurationsTimelockCallNotFound, error) {
	out := new(Fdc2RequestFeeConfigurationsTimelockCallNotFound)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "TimelockCallNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsTimelockNotAllowedYet represents a TimelockNotAllowedYet error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTimelockNotAllowedYet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockNotAllowedYet()
func Fdc2RequestFeeConfigurationsTimelockNotAllowedYetErrorID() common.Hash {
	return common.HexToHash("0x6124e5c2c545bec1d02048c68500eebdabeae0cc69995a629ebccf8a62f69083")
}

// UnpackTimelockNotAllowedYetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockNotAllowedYet()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackTimelockNotAllowedYetError(raw []byte) (*Fdc2RequestFeeConfigurationsTimelockNotAllowedYet, error) {
	out := new(Fdc2RequestFeeConfigurationsTimelockNotAllowedYet)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "TimelockNotAllowedYet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsTypeAndSourceCombinationNotSupported represents a TypeAndSourceCombinationNotSupported error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsTypeAndSourceCombinationNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TypeAndSourceCombinationNotSupported()
func Fdc2RequestFeeConfigurationsTypeAndSourceCombinationNotSupportedErrorID() common.Hash {
	return common.HexToHash("0xbc5759d88213ca3e83a20bbcd2b610665f65a7e853a4fb97905fae26e8f93cb7")
}

// UnpackTypeAndSourceCombinationNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TypeAndSourceCombinationNotSupported()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackTypeAndSourceCombinationNotSupportedError(raw []byte) (*Fdc2RequestFeeConfigurationsTypeAndSourceCombinationNotSupported, error) {
	out := new(Fdc2RequestFeeConfigurationsTypeAndSourceCombinationNotSupported)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "TypeAndSourceCombinationNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func Fdc2RequestFeeConfigurationsUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*Fdc2RequestFeeConfigurationsUUPSUnauthorizedCallContext, error) {
	out := new(Fdc2RequestFeeConfigurationsUUPSUnauthorizedCallContext)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// Fdc2RequestFeeConfigurationsUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the Fdc2RequestFeeConfigurations contract.
type Fdc2RequestFeeConfigurationsUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func Fdc2RequestFeeConfigurationsUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (fdc2RequestFeeConfigurations *Fdc2RequestFeeConfigurations) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*Fdc2RequestFeeConfigurationsUUPSUnsupportedProxiableUUID, error) {
	out := new(Fdc2RequestFeeConfigurationsUUPSUnsupportedProxiableUUID)
	if err := fdc2RequestFeeConfigurations.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}
