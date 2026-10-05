// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractregistry

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

// ContractRegistryMetaData contains all meta data concerning the ContractRegistry contract.
var ContractRegistryMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllContracts\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_nameHash\",\"type\":\"bytes32\"}],\"name\":\"getContractAddressByHash\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"getContractAddressByName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_nameHashes\",\"type\":\"bytes32[]\"}],\"name\":\"getContractAddressesByHash\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_names\",\"type\":\"string[]\"}],\"name\":\"getContractAddressesByName\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "ContractRegistry",
}

// ContractRegistry is an auto generated Go binding around an Ethereum contract.
type ContractRegistry struct {
	abi abi.ABI
}

// NewContractRegistry creates a new instance of ContractRegistry.
func NewContractRegistry() *ContractRegistry {
	parsed, err := ContractRegistryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ContractRegistry{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ContractRegistry) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _addressUpdater) returns()
func (contractRegistry *ContractRegistry) PackConstructor(_addressUpdater common.Address) []byte {
	enc, err := contractRegistry.abi.Pack("", _addressUpdater)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (contractRegistry *ContractRegistry) PackGetAddressUpdater() []byte {
	enc, err := contractRegistry.abi.Pack("getAddressUpdater")
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
func (contractRegistry *ContractRegistry) TryPackGetAddressUpdater() ([]byte, error) {
	return contractRegistry.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (contractRegistry *ContractRegistry) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := contractRegistry.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAllContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18d3ce96.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAllContracts() view returns(string[], address[])
func (contractRegistry *ContractRegistry) PackGetAllContracts() []byte {
	enc, err := contractRegistry.abi.Pack("getAllContracts")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAllContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x18d3ce96.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAllContracts() view returns(string[], address[])
func (contractRegistry *ContractRegistry) TryPackGetAllContracts() ([]byte, error) {
	return contractRegistry.abi.Pack("getAllContracts")
}

// GetAllContractsOutput serves as a container for the return parameters of contract
// method GetAllContracts.
type GetAllContractsOutput struct {
	Arg0 []string
	Arg1 []common.Address
}

// UnpackGetAllContracts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x18d3ce96.
//
// Solidity: function getAllContracts() view returns(string[], address[])
func (contractRegistry *ContractRegistry) UnpackGetAllContracts(data []byte) (GetAllContractsOutput, error) {
	out, err := contractRegistry.abi.Unpack("getAllContracts", data)
	outstruct := new(GetAllContractsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Arg0 = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Arg1 = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	return *outstruct, nil
}

// PackGetContractAddressByHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x159354a2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getContractAddressByHash(bytes32 _nameHash) view returns(address)
func (contractRegistry *ContractRegistry) PackGetContractAddressByHash(nameHash [32]byte) []byte {
	enc, err := contractRegistry.abi.Pack("getContractAddressByHash", nameHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetContractAddressByHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x159354a2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getContractAddressByHash(bytes32 _nameHash) view returns(address)
func (contractRegistry *ContractRegistry) TryPackGetContractAddressByHash(nameHash [32]byte) ([]byte, error) {
	return contractRegistry.abi.Pack("getContractAddressByHash", nameHash)
}

// UnpackGetContractAddressByHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x159354a2.
//
// Solidity: function getContractAddressByHash(bytes32 _nameHash) view returns(address)
func (contractRegistry *ContractRegistry) UnpackGetContractAddressByHash(data []byte) (common.Address, error) {
	out, err := contractRegistry.abi.Unpack("getContractAddressByHash", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetContractAddressByName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82760fca.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getContractAddressByName(string _name) view returns(address)
func (contractRegistry *ContractRegistry) PackGetContractAddressByName(name string) []byte {
	enc, err := contractRegistry.abi.Pack("getContractAddressByName", name)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetContractAddressByName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x82760fca.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getContractAddressByName(string _name) view returns(address)
func (contractRegistry *ContractRegistry) TryPackGetContractAddressByName(name string) ([]byte, error) {
	return contractRegistry.abi.Pack("getContractAddressByName", name)
}

// UnpackGetContractAddressByName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x82760fca.
//
// Solidity: function getContractAddressByName(string _name) view returns(address)
func (contractRegistry *ContractRegistry) UnpackGetContractAddressByName(data []byte) (common.Address, error) {
	out, err := contractRegistry.abi.Unpack("getContractAddressByName", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetContractAddressesByHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e11e2d1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getContractAddressesByHash(bytes32[] _nameHashes) view returns(address[])
func (contractRegistry *ContractRegistry) PackGetContractAddressesByHash(nameHashes [][32]byte) []byte {
	enc, err := contractRegistry.abi.Pack("getContractAddressesByHash", nameHashes)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetContractAddressesByHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e11e2d1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getContractAddressesByHash(bytes32[] _nameHashes) view returns(address[])
func (contractRegistry *ContractRegistry) TryPackGetContractAddressesByHash(nameHashes [][32]byte) ([]byte, error) {
	return contractRegistry.abi.Pack("getContractAddressesByHash", nameHashes)
}

// UnpackGetContractAddressesByHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5e11e2d1.
//
// Solidity: function getContractAddressesByHash(bytes32[] _nameHashes) view returns(address[])
func (contractRegistry *ContractRegistry) UnpackGetContractAddressesByHash(data []byte) ([]common.Address, error) {
	out, err := contractRegistry.abi.Unpack("getContractAddressesByHash", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetContractAddressesByName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x76d2b1af.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getContractAddressesByName(string[] _names) view returns(address[])
func (contractRegistry *ContractRegistry) PackGetContractAddressesByName(names []string) []byte {
	enc, err := contractRegistry.abi.Pack("getContractAddressesByName", names)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetContractAddressesByName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x76d2b1af.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getContractAddressesByName(string[] _names) view returns(address[])
func (contractRegistry *ContractRegistry) TryPackGetContractAddressesByName(names []string) ([]byte, error) {
	return contractRegistry.abi.Pack("getContractAddressesByName", names)
}

// UnpackGetContractAddressesByName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x76d2b1af.
//
// Solidity: function getContractAddressesByName(string[] _names) view returns(address[])
func (contractRegistry *ContractRegistry) UnpackGetContractAddressesByName(data []byte) ([]common.Address, error) {
	out, err := contractRegistry.abi.Unpack("getContractAddressesByName", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (contractRegistry *ContractRegistry) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := contractRegistry.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (contractRegistry *ContractRegistry) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return contractRegistry.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}
