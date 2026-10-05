// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package externaladdresses

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

// ExternalAddressesMetaData contains all meta data concerning the ExternalAddresses contract.
var ExternalAddressesMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"fdc2Hub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2Verification\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relay\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "ExternalAddresses",
}

// ExternalAddresses is an auto generated Go binding around an Ethereum contract.
type ExternalAddresses struct {
	abi abi.ABI
}

// NewExternalAddresses creates a new instance of ExternalAddresses.
func NewExternalAddresses() *ExternalAddresses {
	parsed, err := ExternalAddressesMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ExternalAddresses{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ExternalAddresses) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackFdc2Hub is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7566ff3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fdc2Hub() view returns(address)
func (externalAddresses *ExternalAddresses) PackFdc2Hub() []byte {
	enc, err := externalAddresses.abi.Pack("fdc2Hub")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFdc2Hub is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7566ff3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fdc2Hub() view returns(address)
func (externalAddresses *ExternalAddresses) TryPackFdc2Hub() ([]byte, error) {
	return externalAddresses.abi.Pack("fdc2Hub")
}

// UnpackFdc2Hub is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa7566ff3.
//
// Solidity: function fdc2Hub() view returns(address)
func (externalAddresses *ExternalAddresses) UnpackFdc2Hub(data []byte) (common.Address, error) {
	out, err := externalAddresses.abi.Unpack("fdc2Hub", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFdc2Verification is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf2e9839.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fdc2Verification() view returns(address)
func (externalAddresses *ExternalAddresses) PackFdc2Verification() []byte {
	enc, err := externalAddresses.abi.Pack("fdc2Verification")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFdc2Verification is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf2e9839.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fdc2Verification() view returns(address)
func (externalAddresses *ExternalAddresses) TryPackFdc2Verification() ([]byte, error) {
	return externalAddresses.abi.Pack("fdc2Verification")
}

// UnpackFdc2Verification is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf2e9839.
//
// Solidity: function fdc2Verification() view returns(address)
func (externalAddresses *ExternalAddresses) UnpackFdc2Verification(data []byte) (common.Address, error) {
	out, err := externalAddresses.abi.Unpack("fdc2Verification", data)
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
func (externalAddresses *ExternalAddresses) PackFlareSystemsManager() []byte {
	enc, err := externalAddresses.abi.Pack("flareSystemsManager")
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
func (externalAddresses *ExternalAddresses) TryPackFlareSystemsManager() ([]byte, error) {
	return externalAddresses.abi.Pack("flareSystemsManager")
}

// UnpackFlareSystemsManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (externalAddresses *ExternalAddresses) UnpackFlareSystemsManager(data []byte) (common.Address, error) {
	out, err := externalAddresses.abi.Unpack("flareSystemsManager", data)
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
func (externalAddresses *ExternalAddresses) PackGetAddressUpdater() []byte {
	enc, err := externalAddresses.abi.Pack("getAddressUpdater")
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
func (externalAddresses *ExternalAddresses) TryPackGetAddressUpdater() ([]byte, error) {
	return externalAddresses.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (externalAddresses *ExternalAddresses) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := externalAddresses.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb59589d1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function relay() view returns(address)
func (externalAddresses *ExternalAddresses) PackRelay() []byte {
	enc, err := externalAddresses.abi.Pack("relay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb59589d1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function relay() view returns(address)
func (externalAddresses *ExternalAddresses) TryPackRelay() ([]byte, error) {
	return externalAddresses.abi.Pack("relay")
}

// UnpackRelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (externalAddresses *ExternalAddresses) UnpackRelay(data []byte) (common.Address, error) {
	out, err := externalAddresses.abi.Unpack("relay", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRewardManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f4ef8a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardManager() view returns(address)
func (externalAddresses *ExternalAddresses) PackRewardManager() []byte {
	enc, err := externalAddresses.abi.Pack("rewardManager")
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
func (externalAddresses *ExternalAddresses) TryPackRewardManager() ([]byte, error) {
	return externalAddresses.abi.Pack("rewardManager")
}

// UnpackRewardManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (externalAddresses *ExternalAddresses) UnpackRewardManager(data []byte) (common.Address, error) {
	out, err := externalAddresses.abi.Unpack("rewardManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (externalAddresses *ExternalAddresses) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := externalAddresses.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (externalAddresses *ExternalAddresses) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return externalAddresses.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}
