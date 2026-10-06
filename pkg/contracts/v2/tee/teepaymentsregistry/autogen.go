// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepaymentsregistry

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

// ITeePaymentsRegistrySourceConfig is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsRegistrySourceConfig struct {
	KeyType      [32]byte
	OpType       [32]byte
	PaymentModel uint8
	TeePayments  common.Address
}

// ITeePaymentsRegistrySourceRegistration is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsRegistrySourceRegistration struct {
	KeyType      [32]byte
	OpType       [32]byte
	PaymentModel uint8
	SourceId     [32]byte
	TeePayments  common.Address
}

// TeePaymentsRegistryMetaData contains all meta data concerning the TeePaymentsRegistry contract.
var TeePaymentsRegistryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"teePayments\",\"type\":\"address\"}],\"name\":\"SourceAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceIdZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceKeyTypeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SourceNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceOpTypeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourcePaymentModelUnknown\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teePayments\",\"type\":\"address\"}],\"name\":\"TeePaymentsNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockCallNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teePayments\",\"type\":\"address\"},{\"internalType\":\"enumPaymentModel\",\"name\":\"expected\",\"type\":\"uint8\"},{\"internalType\":\"enumPaymentModel\",\"name\":\"actual\",\"type\":\"uint8\"}],\"name\":\"WrongPaymentModel\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"enumPaymentModel\",\"name\":\"paymentModel\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"teePayments\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structITeePaymentsRegistry.SourceRegistration[]\",\"name\":\"registrations\",\"type\":\"tuple[]\"}],\"name\":\"SourcesRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"sourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"SourcesUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredSourceIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredTeePaymentsContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getSourceConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"enumPaymentModel\",\"name\":\"paymentModel\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"teePayments\",\"type\":\"address\"}],\"internalType\":\"structITeePaymentsRegistry.SourceConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teePayments\",\"type\":\"address\"}],\"name\":\"getSourceIdsForTeePayments\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getSourceKeyTypeAndTeePayments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_teePayments\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getSourceOpTypeAndTeePayments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_teePayments\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getSourcePaymentModel\",\"outputs\":[{\"internalType\":\"enumPaymentModel\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getTeePaymentsForSource\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"isSourceRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"enumPaymentModel\",\"name\":\"paymentModel\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"teePayments\",\"type\":\"address\"}],\"internalType\":\"structITeePaymentsRegistry.SourceRegistration[]\",\"name\":\"_registrations\",\"type\":\"tuple[]\"}],\"name\":\"registerSources\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_sourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"unregisterSources\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "TeePaymentsRegistry",
}

// TeePaymentsRegistry is an auto generated Go binding around an Ethereum contract.
type TeePaymentsRegistry struct {
	abi abi.ABI
}

// NewTeePaymentsRegistry creates a new instance of TeePaymentsRegistry.
func NewTeePaymentsRegistry() *TeePaymentsRegistry {
	parsed, err := TeePaymentsRegistryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TeePaymentsRegistry{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TeePaymentsRegistry) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePaymentsRegistry *TeePaymentsRegistry) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := teePaymentsRegistry.abi.Pack("UPGRADE_INTERFACE_VERSION")
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := teePaymentsRegistry.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
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
func (teePaymentsRegistry *TeePaymentsRegistry) PackCancelGovernanceCall(encodedCall []byte) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("cancelGovernanceCall", encodedCall)
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackCancelGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("cancelGovernanceCall", encodedCall)
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (teePaymentsRegistry *TeePaymentsRegistry) PackExecuteGovernanceCall(encodedCall []byte) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("executeGovernanceCall", encodedCall)
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackExecuteGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("executeGovernanceCall", encodedCall)
}

// PackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (teePaymentsRegistry *TeePaymentsRegistry) PackGetAddressUpdater() []byte {
	enc, err := teePaymentsRegistry.abi.Pack("getAddressUpdater")
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackGetAddressUpdater() ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := teePaymentsRegistry.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetRegisteredSourceIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xec02c8e0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRegisteredSourceIds() view returns(bytes32[])
func (teePaymentsRegistry *TeePaymentsRegistry) PackGetRegisteredSourceIds() []byte {
	enc, err := teePaymentsRegistry.abi.Pack("getRegisteredSourceIds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRegisteredSourceIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xec02c8e0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRegisteredSourceIds() view returns(bytes32[])
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackGetRegisteredSourceIds() ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("getRegisteredSourceIds")
}

// UnpackGetRegisteredSourceIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec02c8e0.
//
// Solidity: function getRegisteredSourceIds() view returns(bytes32[])
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGetRegisteredSourceIds(data []byte) ([][32]byte, error) {
	out, err := teePaymentsRegistry.abi.Unpack("getRegisteredSourceIds", data)
	if err != nil {
		return *new([][32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	return out0, nil
}

// PackGetRegisteredTeePaymentsContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2e5f7c1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRegisteredTeePaymentsContracts() view returns(address[])
func (teePaymentsRegistry *TeePaymentsRegistry) PackGetRegisteredTeePaymentsContracts() []byte {
	enc, err := teePaymentsRegistry.abi.Pack("getRegisteredTeePaymentsContracts")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRegisteredTeePaymentsContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2e5f7c1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRegisteredTeePaymentsContracts() view returns(address[])
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackGetRegisteredTeePaymentsContracts() ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("getRegisteredTeePaymentsContracts")
}

// UnpackGetRegisteredTeePaymentsContracts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf2e5f7c1.
//
// Solidity: function getRegisteredTeePaymentsContracts() view returns(address[])
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGetRegisteredTeePaymentsContracts(data []byte) ([]common.Address, error) {
	out, err := teePaymentsRegistry.abi.Unpack("getRegisteredTeePaymentsContracts", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetSourceConfig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a5b5cfd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSourceConfig(bytes32 _sourceId) view returns((bytes32,bytes32,uint8,address))
func (teePaymentsRegistry *TeePaymentsRegistry) PackGetSourceConfig(sourceId [32]byte) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("getSourceConfig", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSourceConfig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a5b5cfd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSourceConfig(bytes32 _sourceId) view returns((bytes32,bytes32,uint8,address))
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackGetSourceConfig(sourceId [32]byte) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("getSourceConfig", sourceId)
}

// UnpackGetSourceConfig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9a5b5cfd.
//
// Solidity: function getSourceConfig(bytes32 _sourceId) view returns((bytes32,bytes32,uint8,address))
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGetSourceConfig(data []byte) (ITeePaymentsRegistrySourceConfig, error) {
	out, err := teePaymentsRegistry.abi.Unpack("getSourceConfig", data)
	if err != nil {
		return *new(ITeePaymentsRegistrySourceConfig), err
	}
	out0 := *abi.ConvertType(out[0], new(ITeePaymentsRegistrySourceConfig)).(*ITeePaymentsRegistrySourceConfig)
	return out0, nil
}

// PackGetSourceIdsForTeePayments is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe4665bd0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSourceIdsForTeePayments(address _teePayments) view returns(bytes32[])
func (teePaymentsRegistry *TeePaymentsRegistry) PackGetSourceIdsForTeePayments(teePayments common.Address) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("getSourceIdsForTeePayments", teePayments)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSourceIdsForTeePayments is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe4665bd0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSourceIdsForTeePayments(address _teePayments) view returns(bytes32[])
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackGetSourceIdsForTeePayments(teePayments common.Address) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("getSourceIdsForTeePayments", teePayments)
}

// UnpackGetSourceIdsForTeePayments is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe4665bd0.
//
// Solidity: function getSourceIdsForTeePayments(address _teePayments) view returns(bytes32[])
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGetSourceIdsForTeePayments(data []byte) ([][32]byte, error) {
	out, err := teePaymentsRegistry.abi.Unpack("getSourceIdsForTeePayments", data)
	if err != nil {
		return *new([][32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	return out0, nil
}

// PackGetSourceKeyTypeAndTeePayments is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x828def57.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSourceKeyTypeAndTeePayments(bytes32 _sourceId) view returns(bytes32 _keyType, address _teePayments)
func (teePaymentsRegistry *TeePaymentsRegistry) PackGetSourceKeyTypeAndTeePayments(sourceId [32]byte) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("getSourceKeyTypeAndTeePayments", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSourceKeyTypeAndTeePayments is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x828def57.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSourceKeyTypeAndTeePayments(bytes32 _sourceId) view returns(bytes32 _keyType, address _teePayments)
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackGetSourceKeyTypeAndTeePayments(sourceId [32]byte) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("getSourceKeyTypeAndTeePayments", sourceId)
}

// GetSourceKeyTypeAndTeePaymentsOutput serves as a container for the return parameters of contract
// method GetSourceKeyTypeAndTeePayments.
type GetSourceKeyTypeAndTeePaymentsOutput struct {
	KeyType     [32]byte
	TeePayments common.Address
}

// UnpackGetSourceKeyTypeAndTeePayments is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x828def57.
//
// Solidity: function getSourceKeyTypeAndTeePayments(bytes32 _sourceId) view returns(bytes32 _keyType, address _teePayments)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGetSourceKeyTypeAndTeePayments(data []byte) (GetSourceKeyTypeAndTeePaymentsOutput, error) {
	out, err := teePaymentsRegistry.abi.Unpack("getSourceKeyTypeAndTeePayments", data)
	outstruct := new(GetSourceKeyTypeAndTeePaymentsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.KeyType = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.TeePayments = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackGetSourceOpTypeAndTeePayments is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x274b42ba.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSourceOpTypeAndTeePayments(bytes32 _sourceId) view returns(bytes32 _opType, address _teePayments)
func (teePaymentsRegistry *TeePaymentsRegistry) PackGetSourceOpTypeAndTeePayments(sourceId [32]byte) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("getSourceOpTypeAndTeePayments", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSourceOpTypeAndTeePayments is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x274b42ba.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSourceOpTypeAndTeePayments(bytes32 _sourceId) view returns(bytes32 _opType, address _teePayments)
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackGetSourceOpTypeAndTeePayments(sourceId [32]byte) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("getSourceOpTypeAndTeePayments", sourceId)
}

// GetSourceOpTypeAndTeePaymentsOutput serves as a container for the return parameters of contract
// method GetSourceOpTypeAndTeePayments.
type GetSourceOpTypeAndTeePaymentsOutput struct {
	OpType      [32]byte
	TeePayments common.Address
}

// UnpackGetSourceOpTypeAndTeePayments is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x274b42ba.
//
// Solidity: function getSourceOpTypeAndTeePayments(bytes32 _sourceId) view returns(bytes32 _opType, address _teePayments)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGetSourceOpTypeAndTeePayments(data []byte) (GetSourceOpTypeAndTeePaymentsOutput, error) {
	out, err := teePaymentsRegistry.abi.Unpack("getSourceOpTypeAndTeePayments", data)
	outstruct := new(GetSourceOpTypeAndTeePaymentsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.OpType = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.TeePayments = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackGetSourcePaymentModel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4ffa436.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSourcePaymentModel(bytes32 _sourceId) view returns(uint8)
func (teePaymentsRegistry *TeePaymentsRegistry) PackGetSourcePaymentModel(sourceId [32]byte) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("getSourcePaymentModel", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSourcePaymentModel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4ffa436.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSourcePaymentModel(bytes32 _sourceId) view returns(uint8)
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackGetSourcePaymentModel(sourceId [32]byte) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("getSourcePaymentModel", sourceId)
}

// UnpackGetSourcePaymentModel is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc4ffa436.
//
// Solidity: function getSourcePaymentModel(bytes32 _sourceId) view returns(uint8)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGetSourcePaymentModel(data []byte) (uint8, error) {
	out, err := teePaymentsRegistry.abi.Unpack("getSourcePaymentModel", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackGetTeePaymentsForSource is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4c1db547.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTeePaymentsForSource(bytes32 _sourceId) view returns(address)
func (teePaymentsRegistry *TeePaymentsRegistry) PackGetTeePaymentsForSource(sourceId [32]byte) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("getTeePaymentsForSource", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTeePaymentsForSource is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4c1db547.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTeePaymentsForSource(bytes32 _sourceId) view returns(address)
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackGetTeePaymentsForSource(sourceId [32]byte) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("getTeePaymentsForSource", sourceId)
}

// UnpackGetTeePaymentsForSource is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4c1db547.
//
// Solidity: function getTeePaymentsForSource(bytes32 _sourceId) view returns(address)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGetTeePaymentsForSource(data []byte) (common.Address, error) {
	out, err := teePaymentsRegistry.abi.Unpack("getTeePaymentsForSource", data)
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
func (teePaymentsRegistry *TeePaymentsRegistry) PackGovernance() []byte {
	enc, err := teePaymentsRegistry.abi.Pack("governance")
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackGovernance() ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := teePaymentsRegistry.abi.Unpack("governance", data)
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
func (teePaymentsRegistry *TeePaymentsRegistry) PackGovernanceSettings() []byte {
	enc, err := teePaymentsRegistry.abi.Pack("governanceSettings")
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackGovernanceSettings() ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := teePaymentsRegistry.abi.Unpack("governanceSettings", data)
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
func (teePaymentsRegistry *TeePaymentsRegistry) PackImplementation() []byte {
	enc, err := teePaymentsRegistry.abi.Pack("implementation")
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackImplementation() ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("implementation")
}

// UnpackImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackImplementation(data []byte) (common.Address, error) {
	out, err := teePaymentsRegistry.abi.Unpack("implementation", data)
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
func (teePaymentsRegistry *TeePaymentsRegistry) PackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teePaymentsRegistry *TeePaymentsRegistry) PackIsExecutor(address common.Address) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("isExecutor", address)
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := teePaymentsRegistry.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsSourceRegistered is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6848bfa4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isSourceRegistered(bytes32 _sourceId) view returns(bool)
func (teePaymentsRegistry *TeePaymentsRegistry) PackIsSourceRegistered(sourceId [32]byte) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("isSourceRegistered", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsSourceRegistered is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6848bfa4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isSourceRegistered(bytes32 _sourceId) view returns(bool)
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackIsSourceRegistered(sourceId [32]byte) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("isSourceRegistered", sourceId)
}

// UnpackIsSourceRegistered is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6848bfa4.
//
// Solidity: function isSourceRegistered(bytes32 _sourceId) view returns(bool)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackIsSourceRegistered(data []byte) (bool, error) {
	out, err := teePaymentsRegistry.abi.Unpack("isSourceRegistered", data)
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
func (teePaymentsRegistry *TeePaymentsRegistry) PackProductionMode() []byte {
	enc, err := teePaymentsRegistry.abi.Pack("productionMode")
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackProductionMode() ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackProductionMode(data []byte) (bool, error) {
	out, err := teePaymentsRegistry.abi.Unpack("productionMode", data)
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
func (teePaymentsRegistry *TeePaymentsRegistry) PackProxiableUUID() []byte {
	enc, err := teePaymentsRegistry.abi.Pack("proxiableUUID")
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackProxiableUUID() ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("proxiableUUID")
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := teePaymentsRegistry.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackRegisterSources is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42d8f760.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function registerSources((bytes32,bytes32,uint8,bytes32,address)[] _registrations) returns()
func (teePaymentsRegistry *TeePaymentsRegistry) PackRegisterSources(registrations []ITeePaymentsRegistrySourceRegistration) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("registerSources", registrations)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegisterSources is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42d8f760.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function registerSources((bytes32,bytes32,uint8,bytes32,address)[] _registrations) returns()
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackRegisterSources(registrations []ITeePaymentsRegistrySourceRegistration) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("registerSources", registrations)
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (teePaymentsRegistry *TeePaymentsRegistry) PackSwitchToProductionMode() []byte {
	enc, err := teePaymentsRegistry.abi.Pack("switchToProductionMode")
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackSwitchToProductionMode() ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("switchToProductionMode")
}

// PackUnregisterSources is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x840c3088.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unregisterSources(bytes32[] _sourceIds) returns()
func (teePaymentsRegistry *TeePaymentsRegistry) PackUnregisterSources(sourceIds [][32]byte) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("unregisterSources", sourceIds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnregisterSources is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x840c3088.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unregisterSources(bytes32[] _sourceIds) returns()
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackUnregisterSources(sourceIds [][32]byte) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("unregisterSources", sourceIds)
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (teePaymentsRegistry *TeePaymentsRegistry) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (teePaymentsRegistry *TeePaymentsRegistry) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := teePaymentsRegistry.abi.Pack("upgradeToAndCall", newImplementation, data)
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
func (teePaymentsRegistry *TeePaymentsRegistry) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return teePaymentsRegistry.abi.Pack("upgradeToAndCall", newImplementation, data)
}

// TeePaymentsRegistryGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const TeePaymentsRegistryGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (TeePaymentsRegistryGovernanceCallTimelocked) ContractEventName() string {
	return TeePaymentsRegistryGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*TeePaymentsRegistryGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsRegistryGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsRegistry.abi.Events[event].Inputs {
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

// TeePaymentsRegistryGovernanceInitialised represents a GovernanceInitialised event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const TeePaymentsRegistryGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (TeePaymentsRegistryGovernanceInitialised) ContractEventName() string {
	return TeePaymentsRegistryGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGovernanceInitialisedEvent(log *types.Log) (*TeePaymentsRegistryGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsRegistryGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsRegistry.abi.Events[event].Inputs {
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

// TeePaymentsRegistryGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const TeePaymentsRegistryGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (TeePaymentsRegistryGovernedProductionModeEntered) ContractEventName() string {
	return TeePaymentsRegistryGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*TeePaymentsRegistryGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsRegistryGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsRegistry.abi.Events[event].Inputs {
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

// TeePaymentsRegistryInitialized represents a Initialized event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const TeePaymentsRegistryInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (TeePaymentsRegistryInitialized) ContractEventName() string {
	return TeePaymentsRegistryInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackInitializedEvent(log *types.Log) (*TeePaymentsRegistryInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsRegistryInitialized)
	if len(log.Data) > 0 {
		if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsRegistry.abi.Events[event].Inputs {
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

// TeePaymentsRegistrySourcesRegistered represents a SourcesRegistered event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistrySourcesRegistered struct {
	Registrations []ITeePaymentsRegistrySourceRegistration
	Raw           *types.Log // Blockchain specific contextual infos
}

const TeePaymentsRegistrySourcesRegisteredEventName = "SourcesRegistered"

// ContractEventName returns the user-defined event name.
func (TeePaymentsRegistrySourcesRegistered) ContractEventName() string {
	return TeePaymentsRegistrySourcesRegisteredEventName
}

// UnpackSourcesRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SourcesRegistered((bytes32,bytes32,uint8,bytes32,address)[] registrations)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackSourcesRegisteredEvent(log *types.Log) (*TeePaymentsRegistrySourcesRegistered, error) {
	event := "SourcesRegistered"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsRegistrySourcesRegistered)
	if len(log.Data) > 0 {
		if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsRegistry.abi.Events[event].Inputs {
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

// TeePaymentsRegistrySourcesUnregistered represents a SourcesUnregistered event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistrySourcesUnregistered struct {
	SourceIds [][32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const TeePaymentsRegistrySourcesUnregisteredEventName = "SourcesUnregistered"

// ContractEventName returns the user-defined event name.
func (TeePaymentsRegistrySourcesUnregistered) ContractEventName() string {
	return TeePaymentsRegistrySourcesUnregisteredEventName
}

// UnpackSourcesUnregisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SourcesUnregistered(bytes32[] sourceIds)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackSourcesUnregisteredEvent(log *types.Log) (*TeePaymentsRegistrySourcesUnregistered, error) {
	event := "SourcesUnregistered"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsRegistrySourcesUnregistered)
	if len(log.Data) > 0 {
		if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsRegistry.abi.Events[event].Inputs {
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

// TeePaymentsRegistryTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeePaymentsRegistryTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (TeePaymentsRegistryTimelockedGovernanceCallCanceled) ContractEventName() string {
	return TeePaymentsRegistryTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*TeePaymentsRegistryTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsRegistryTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsRegistry.abi.Events[event].Inputs {
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

// TeePaymentsRegistryTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeePaymentsRegistryTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (TeePaymentsRegistryTimelockedGovernanceCallExecuted) ContractEventName() string {
	return TeePaymentsRegistryTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*TeePaymentsRegistryTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsRegistryTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsRegistry.abi.Events[event].Inputs {
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

// TeePaymentsRegistryUpgraded represents a Upgraded event raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const TeePaymentsRegistryUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (TeePaymentsRegistryUpgraded) ContractEventName() string {
	return TeePaymentsRegistryUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackUpgradedEvent(log *types.Log) (*TeePaymentsRegistryUpgraded, error) {
	event := "Upgraded"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsRegistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsRegistryUpgraded)
	if len(log.Data) > 0 {
		if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsRegistry.abi.Events[event].Inputs {
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
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["AlreadyInProductionMode"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackAlreadyInProductionModeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["GovernedAddressZero"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackGovernedAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["GovernedAlreadyInitialized"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackGovernedAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["OnlyExecutor"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["SourceAlreadyRegistered"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackSourceAlreadyRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["SourceIdZero"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackSourceIdZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["SourceKeyTypeZero"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackSourceKeyTypeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["SourceNotRegistered"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackSourceNotRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["SourceOpTypeZero"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackSourceOpTypeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["SourcePaymentModelUnknown"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackSourcePaymentModelUnknownError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["TeePaymentsNotContract"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackTeePaymentsNotContractError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["TimelockCallNotFound"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackTimelockCallNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["TimelockNotAllowedYet"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackTimelockNotAllowedYetError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsRegistry.abi.Errors["WrongPaymentModel"].ID.Bytes()[:4]) {
		return teePaymentsRegistry.UnpackWrongPaymentModelError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// TeePaymentsRegistryAddressEmptyCode represents a AddressEmptyCode error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func TeePaymentsRegistryAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackAddressEmptyCodeError(raw []byte) (*TeePaymentsRegistryAddressEmptyCode, error) {
	out := new(TeePaymentsRegistryAddressEmptyCode)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryAlreadyInProductionMode represents a AlreadyInProductionMode error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryAlreadyInProductionMode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInProductionMode()
func TeePaymentsRegistryAlreadyInProductionModeErrorID() common.Hash {
	return common.HexToHash("0x72a23ec2b1fdbc3de7d5f411b572fcf8ad8409803196930f63ddf40a2a2b2594")
}

// UnpackAlreadyInProductionModeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInProductionMode()
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackAlreadyInProductionModeError(raw []byte) (*TeePaymentsRegistryAlreadyInProductionMode, error) {
	out := new(TeePaymentsRegistryAlreadyInProductionMode)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "AlreadyInProductionMode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func TeePaymentsRegistryERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackERC1967InvalidImplementationError(raw []byte) (*TeePaymentsRegistryERC1967InvalidImplementation, error) {
	out := new(TeePaymentsRegistryERC1967InvalidImplementation)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryERC1967NonPayable represents a ERC1967NonPayable error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func TeePaymentsRegistryERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackERC1967NonPayableError(raw []byte) (*TeePaymentsRegistryERC1967NonPayable, error) {
	out := new(TeePaymentsRegistryERC1967NonPayable)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryFailedCall represents a FailedCall error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func TeePaymentsRegistryFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackFailedCallError(raw []byte) (*TeePaymentsRegistryFailedCall, error) {
	out := new(TeePaymentsRegistryFailedCall)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryGovernedAddressZero represents a GovernedAddressZero error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryGovernedAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAddressZero()
func TeePaymentsRegistryGovernedAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x695725f6b429bac0cc5a2fa88214e31520e0c69f2fb0e2fb881a7130a564f49f")
}

// UnpackGovernedAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAddressZero()
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGovernedAddressZeroError(raw []byte) (*TeePaymentsRegistryGovernedAddressZero, error) {
	out := new(TeePaymentsRegistryGovernedAddressZero)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "GovernedAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryGovernedAlreadyInitialized represents a GovernedAlreadyInitialized error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryGovernedAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAlreadyInitialized()
func TeePaymentsRegistryGovernedAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0xb2f0162156d83828197776d56bbfba0a7a9d5936fad70f0a3014ae18cc7aa33c")
}

// UnpackGovernedAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAlreadyInitialized()
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackGovernedAlreadyInitializedError(raw []byte) (*TeePaymentsRegistryGovernedAlreadyInitialized, error) {
	out := new(TeePaymentsRegistryGovernedAlreadyInitialized)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "GovernedAlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryInvalidInitialization represents a InvalidInitialization error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func TeePaymentsRegistryInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackInvalidInitializationError(raw []byte) (*TeePaymentsRegistryInvalidInitialization, error) {
	out := new(TeePaymentsRegistryInvalidInitialization)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryNotInitializing represents a NotInitializing error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func TeePaymentsRegistryNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackNotInitializingError(raw []byte) (*TeePaymentsRegistryNotInitializing, error) {
	out := new(TeePaymentsRegistryNotInitializing)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryOnlyExecutor represents a OnlyExecutor error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryOnlyExecutor struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExecutor()
func TeePaymentsRegistryOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x7fb6be02aaf40e420b0e7eb67e43a3feee949e8932a3c64e33b591a436670841")
}

// UnpackOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExecutor()
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackOnlyExecutorError(raw []byte) (*TeePaymentsRegistryOnlyExecutor, error) {
	out := new(TeePaymentsRegistryOnlyExecutor)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "OnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryOnlyGovernance represents a OnlyGovernance error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func TeePaymentsRegistryOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackOnlyGovernanceError(raw []byte) (*TeePaymentsRegistryOnlyGovernance, error) {
	out := new(TeePaymentsRegistryOnlyGovernance)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistrySourceAlreadyRegistered represents a SourceAlreadyRegistered error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistrySourceAlreadyRegistered struct {
	SourceId    [32]byte
	TeePayments common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceAlreadyRegistered(bytes32 sourceId, address teePayments)
func TeePaymentsRegistrySourceAlreadyRegisteredErrorID() common.Hash {
	return common.HexToHash("0xf1a092009bfc94d25889b5bbb284a7b764e44206acec482d4f39b5e18040c971")
}

// UnpackSourceAlreadyRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceAlreadyRegistered(bytes32 sourceId, address teePayments)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackSourceAlreadyRegisteredError(raw []byte) (*TeePaymentsRegistrySourceAlreadyRegistered, error) {
	out := new(TeePaymentsRegistrySourceAlreadyRegistered)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "SourceAlreadyRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistrySourceIdZero represents a SourceIdZero error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistrySourceIdZero struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceIdZero(uint256 index)
func TeePaymentsRegistrySourceIdZeroErrorID() common.Hash {
	return common.HexToHash("0xd9950937ebabdb658eaf457d54197e233dd4841a64bb53f18118a853631876e7")
}

// UnpackSourceIdZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceIdZero(uint256 index)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackSourceIdZeroError(raw []byte) (*TeePaymentsRegistrySourceIdZero, error) {
	out := new(TeePaymentsRegistrySourceIdZero)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "SourceIdZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistrySourceKeyTypeZero represents a SourceKeyTypeZero error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistrySourceKeyTypeZero struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceKeyTypeZero(uint256 index)
func TeePaymentsRegistrySourceKeyTypeZeroErrorID() common.Hash {
	return common.HexToHash("0xeeb24635fe3f0a4f44f49b89472e645d49ec073ab20e68b83785acdbc3484f63")
}

// UnpackSourceKeyTypeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceKeyTypeZero(uint256 index)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackSourceKeyTypeZeroError(raw []byte) (*TeePaymentsRegistrySourceKeyTypeZero, error) {
	out := new(TeePaymentsRegistrySourceKeyTypeZero)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "SourceKeyTypeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistrySourceNotRegistered represents a SourceNotRegistered error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistrySourceNotRegistered struct {
	SourceId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceNotRegistered(bytes32 sourceId)
func TeePaymentsRegistrySourceNotRegisteredErrorID() common.Hash {
	return common.HexToHash("0x321b06cf0196e728deff60d765e36d0ef5386b39e55e899738403838fb35f40a")
}

// UnpackSourceNotRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceNotRegistered(bytes32 sourceId)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackSourceNotRegisteredError(raw []byte) (*TeePaymentsRegistrySourceNotRegistered, error) {
	out := new(TeePaymentsRegistrySourceNotRegistered)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "SourceNotRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistrySourceOpTypeZero represents a SourceOpTypeZero error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistrySourceOpTypeZero struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceOpTypeZero(uint256 index)
func TeePaymentsRegistrySourceOpTypeZeroErrorID() common.Hash {
	return common.HexToHash("0xef1189f6c1f88aa0a6fd81ffc2c6c160ee2269443b0c873bcf7ce43b2cfef8a7")
}

// UnpackSourceOpTypeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceOpTypeZero(uint256 index)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackSourceOpTypeZeroError(raw []byte) (*TeePaymentsRegistrySourceOpTypeZero, error) {
	out := new(TeePaymentsRegistrySourceOpTypeZero)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "SourceOpTypeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistrySourcePaymentModelUnknown represents a SourcePaymentModelUnknown error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistrySourcePaymentModelUnknown struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourcePaymentModelUnknown(uint256 index)
func TeePaymentsRegistrySourcePaymentModelUnknownErrorID() common.Hash {
	return common.HexToHash("0x287d290df524bb6943a708cad1441bc00da738b9b88d19b86e8c763e72bb6706")
}

// UnpackSourcePaymentModelUnknownError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourcePaymentModelUnknown(uint256 index)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackSourcePaymentModelUnknownError(raw []byte) (*TeePaymentsRegistrySourcePaymentModelUnknown, error) {
	out := new(TeePaymentsRegistrySourcePaymentModelUnknown)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "SourcePaymentModelUnknown", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryTeePaymentsNotContract represents a TeePaymentsNotContract error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryTeePaymentsNotContract struct {
	TeePayments common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeePaymentsNotContract(address teePayments)
func TeePaymentsRegistryTeePaymentsNotContractErrorID() common.Hash {
	return common.HexToHash("0xb7164ffc084e827e00e15f92609dc160222d1fd4d4a655df9f73d1084110f588")
}

// UnpackTeePaymentsNotContractError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeePaymentsNotContract(address teePayments)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackTeePaymentsNotContractError(raw []byte) (*TeePaymentsRegistryTeePaymentsNotContract, error) {
	out := new(TeePaymentsRegistryTeePaymentsNotContract)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "TeePaymentsNotContract", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryTimelockCallNotFound represents a TimelockCallNotFound error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryTimelockCallNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockCallNotFound()
func TeePaymentsRegistryTimelockCallNotFoundErrorID() common.Hash {
	return common.HexToHash("0x262ac987dd09120380d393a0c52cdef1ca58949d1107dc26d8da9f7b527eb33d")
}

// UnpackTimelockCallNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockCallNotFound()
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackTimelockCallNotFoundError(raw []byte) (*TeePaymentsRegistryTimelockCallNotFound, error) {
	out := new(TeePaymentsRegistryTimelockCallNotFound)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "TimelockCallNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryTimelockNotAllowedYet represents a TimelockNotAllowedYet error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryTimelockNotAllowedYet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockNotAllowedYet()
func TeePaymentsRegistryTimelockNotAllowedYetErrorID() common.Hash {
	return common.HexToHash("0x6124e5c2c545bec1d02048c68500eebdabeae0cc69995a629ebccf8a62f69083")
}

// UnpackTimelockNotAllowedYetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockNotAllowedYet()
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackTimelockNotAllowedYetError(raw []byte) (*TeePaymentsRegistryTimelockNotAllowedYet, error) {
	out := new(TeePaymentsRegistryTimelockNotAllowedYet)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "TimelockNotAllowedYet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func TeePaymentsRegistryUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*TeePaymentsRegistryUUPSUnauthorizedCallContext, error) {
	out := new(TeePaymentsRegistryUUPSUnauthorizedCallContext)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func TeePaymentsRegistryUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*TeePaymentsRegistryUUPSUnsupportedProxiableUUID, error) {
	out := new(TeePaymentsRegistryUUPSUnsupportedProxiableUUID)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsRegistryWrongPaymentModel represents a WrongPaymentModel error raised by the TeePaymentsRegistry contract.
type TeePaymentsRegistryWrongPaymentModel struct {
	TeePayments common.Address
	Expected    uint8
	Actual      uint8
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WrongPaymentModel(address teePayments, uint8 expected, uint8 actual)
func TeePaymentsRegistryWrongPaymentModelErrorID() common.Hash {
	return common.HexToHash("0x25396cade563dd0c1040be8b5eb842c8be4f38b9a1b71d2a90697b8411624fc4")
}

// UnpackWrongPaymentModelError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WrongPaymentModel(address teePayments, uint8 expected, uint8 actual)
func (teePaymentsRegistry *TeePaymentsRegistry) UnpackWrongPaymentModelError(raw []byte) (*TeePaymentsRegistryWrongPaymentModel, error) {
	out := new(TeePaymentsRegistryWrongPaymentModel)
	if err := teePaymentsRegistry.abi.UnpackIntoInterface(out, "WrongPaymentModel", raw); err != nil {
		return nil, err
	}
	return out, nil
}
