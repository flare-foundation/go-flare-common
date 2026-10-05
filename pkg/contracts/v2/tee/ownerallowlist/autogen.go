// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ownerallowlist

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

// OwnerAllowlistMetaData contains all meta data concerning the OwnerAllowlist contract.
var OwnerAllowlistMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AllExtensionOwnersAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AllExtensionOwnersDisallowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeMachineOwnersAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeMachineOwnersDisallowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeWalletProjectOwnersAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeWalletProjectOwnersDisallowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedExtensionOwnersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedExtensionOwnersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeMachineOwnersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeMachineOwnersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeWalletProjectOwnersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeWalletProjectOwnersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"addAllowedExtensionOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"addAllowedTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"addAllowedTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allExtensionOwnersAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_allAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allTeeMachineOwnersAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_allAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allTeeWalletProjectOwnersAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_allAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowAllExtensionOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allowAllTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allowAllTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disallowAllExtensionOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"disallowAllTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"disallowAllTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowedExtensionOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_allowedOwners\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getAllowedTeeMachineOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_allowedOwners\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getAllowedTeeWalletProjectOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_allowedOwners\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"isAllowedExtensionOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"isAllowedTeeMachineOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"isAllowedTeeWalletProjectOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"removeAllowedExtensionOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"removeAllowedTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"removeAllowedTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "OwnerAllowlist",
}

// OwnerAllowlist is an auto generated Go binding around an Ethereum contract.
type OwnerAllowlist struct {
	abi abi.ABI
}

// NewOwnerAllowlist creates a new instance of OwnerAllowlist.
func NewOwnerAllowlist() *OwnerAllowlist {
	parsed, err := OwnerAllowlistMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &OwnerAllowlist{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *OwnerAllowlist) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddAllowedExtensionOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x55414fd1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addAllowedExtensionOwners(address[] _owners) returns()
func (ownerAllowlist *OwnerAllowlist) PackAddAllowedExtensionOwners(owners []common.Address) []byte {
	enc, err := ownerAllowlist.abi.Pack("addAllowedExtensionOwners", owners)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddAllowedExtensionOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x55414fd1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addAllowedExtensionOwners(address[] _owners) returns()
func (ownerAllowlist *OwnerAllowlist) TryPackAddAllowedExtensionOwners(owners []common.Address) ([]byte, error) {
	return ownerAllowlist.abi.Pack("addAllowedExtensionOwners", owners)
}

// PackAddAllowedTeeMachineOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8f3cab7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (ownerAllowlist *OwnerAllowlist) PackAddAllowedTeeMachineOwners(extensionId *big.Int, owners []common.Address) []byte {
	enc, err := ownerAllowlist.abi.Pack("addAllowedTeeMachineOwners", extensionId, owners)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddAllowedTeeMachineOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa8f3cab7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (ownerAllowlist *OwnerAllowlist) TryPackAddAllowedTeeMachineOwners(extensionId *big.Int, owners []common.Address) ([]byte, error) {
	return ownerAllowlist.abi.Pack("addAllowedTeeMachineOwners", extensionId, owners)
}

// PackAddAllowedTeeWalletProjectOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1c34746.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (ownerAllowlist *OwnerAllowlist) PackAddAllowedTeeWalletProjectOwners(extensionId *big.Int, owners []common.Address) []byte {
	enc, err := ownerAllowlist.abi.Pack("addAllowedTeeWalletProjectOwners", extensionId, owners)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddAllowedTeeWalletProjectOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1c34746.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (ownerAllowlist *OwnerAllowlist) TryPackAddAllowedTeeWalletProjectOwners(extensionId *big.Int, owners []common.Address) ([]byte, error) {
	return ownerAllowlist.abi.Pack("addAllowedTeeWalletProjectOwners", extensionId, owners)
}

// PackAllExtensionOwnersAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x899aa572.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allExtensionOwnersAllowed() view returns(bool _allAllowed)
func (ownerAllowlist *OwnerAllowlist) PackAllExtensionOwnersAllowed() []byte {
	enc, err := ownerAllowlist.abi.Pack("allExtensionOwnersAllowed")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllExtensionOwnersAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x899aa572.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allExtensionOwnersAllowed() view returns(bool _allAllowed)
func (ownerAllowlist *OwnerAllowlist) TryPackAllExtensionOwnersAllowed() ([]byte, error) {
	return ownerAllowlist.abi.Pack("allExtensionOwnersAllowed")
}

// UnpackAllExtensionOwnersAllowed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x899aa572.
//
// Solidity: function allExtensionOwnersAllowed() view returns(bool _allAllowed)
func (ownerAllowlist *OwnerAllowlist) UnpackAllExtensionOwnersAllowed(data []byte) (bool, error) {
	out, err := ownerAllowlist.abi.Unpack("allExtensionOwnersAllowed", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackAllTeeMachineOwnersAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2a30b9f4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allTeeMachineOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (ownerAllowlist *OwnerAllowlist) PackAllTeeMachineOwnersAllowed(extensionId *big.Int) []byte {
	enc, err := ownerAllowlist.abi.Pack("allTeeMachineOwnersAllowed", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllTeeMachineOwnersAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2a30b9f4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allTeeMachineOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (ownerAllowlist *OwnerAllowlist) TryPackAllTeeMachineOwnersAllowed(extensionId *big.Int) ([]byte, error) {
	return ownerAllowlist.abi.Pack("allTeeMachineOwnersAllowed", extensionId)
}

// UnpackAllTeeMachineOwnersAllowed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2a30b9f4.
//
// Solidity: function allTeeMachineOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (ownerAllowlist *OwnerAllowlist) UnpackAllTeeMachineOwnersAllowed(data []byte) (bool, error) {
	out, err := ownerAllowlist.abi.Unpack("allTeeMachineOwnersAllowed", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackAllTeeWalletProjectOwnersAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8398fc2e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allTeeWalletProjectOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (ownerAllowlist *OwnerAllowlist) PackAllTeeWalletProjectOwnersAllowed(extensionId *big.Int) []byte {
	enc, err := ownerAllowlist.abi.Pack("allTeeWalletProjectOwnersAllowed", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllTeeWalletProjectOwnersAllowed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8398fc2e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allTeeWalletProjectOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (ownerAllowlist *OwnerAllowlist) TryPackAllTeeWalletProjectOwnersAllowed(extensionId *big.Int) ([]byte, error) {
	return ownerAllowlist.abi.Pack("allTeeWalletProjectOwnersAllowed", extensionId)
}

// UnpackAllTeeWalletProjectOwnersAllowed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8398fc2e.
//
// Solidity: function allTeeWalletProjectOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (ownerAllowlist *OwnerAllowlist) UnpackAllTeeWalletProjectOwnersAllowed(data []byte) (bool, error) {
	out, err := ownerAllowlist.abi.Unpack("allTeeWalletProjectOwnersAllowed", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackAllowAllExtensionOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef10f5ad.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allowAllExtensionOwners() returns()
func (ownerAllowlist *OwnerAllowlist) PackAllowAllExtensionOwners() []byte {
	enc, err := ownerAllowlist.abi.Pack("allowAllExtensionOwners")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllowAllExtensionOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef10f5ad.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allowAllExtensionOwners() returns()
func (ownerAllowlist *OwnerAllowlist) TryPackAllowAllExtensionOwners() ([]byte, error) {
	return ownerAllowlist.abi.Pack("allowAllExtensionOwners")
}

// PackAllowAllTeeMachineOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1b489f2d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allowAllTeeMachineOwners(uint256 _extensionId) returns()
func (ownerAllowlist *OwnerAllowlist) PackAllowAllTeeMachineOwners(extensionId *big.Int) []byte {
	enc, err := ownerAllowlist.abi.Pack("allowAllTeeMachineOwners", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllowAllTeeMachineOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1b489f2d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allowAllTeeMachineOwners(uint256 _extensionId) returns()
func (ownerAllowlist *OwnerAllowlist) TryPackAllowAllTeeMachineOwners(extensionId *big.Int) ([]byte, error) {
	return ownerAllowlist.abi.Pack("allowAllTeeMachineOwners", extensionId)
}

// PackAllowAllTeeWalletProjectOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x61f957e7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function allowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (ownerAllowlist *OwnerAllowlist) PackAllowAllTeeWalletProjectOwners(extensionId *big.Int) []byte {
	enc, err := ownerAllowlist.abi.Pack("allowAllTeeWalletProjectOwners", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAllowAllTeeWalletProjectOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x61f957e7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function allowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (ownerAllowlist *OwnerAllowlist) TryPackAllowAllTeeWalletProjectOwners(extensionId *big.Int) ([]byte, error) {
	return ownerAllowlist.abi.Pack("allowAllTeeWalletProjectOwners", extensionId)
}

// PackDisallowAllExtensionOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x12282f33.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function disallowAllExtensionOwners() returns()
func (ownerAllowlist *OwnerAllowlist) PackDisallowAllExtensionOwners() []byte {
	enc, err := ownerAllowlist.abi.Pack("disallowAllExtensionOwners")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDisallowAllExtensionOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x12282f33.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function disallowAllExtensionOwners() returns()
func (ownerAllowlist *OwnerAllowlist) TryPackDisallowAllExtensionOwners() ([]byte, error) {
	return ownerAllowlist.abi.Pack("disallowAllExtensionOwners")
}

// PackDisallowAllTeeMachineOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb86936ba.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function disallowAllTeeMachineOwners(uint256 _extensionId) returns()
func (ownerAllowlist *OwnerAllowlist) PackDisallowAllTeeMachineOwners(extensionId *big.Int) []byte {
	enc, err := ownerAllowlist.abi.Pack("disallowAllTeeMachineOwners", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDisallowAllTeeMachineOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb86936ba.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function disallowAllTeeMachineOwners(uint256 _extensionId) returns()
func (ownerAllowlist *OwnerAllowlist) TryPackDisallowAllTeeMachineOwners(extensionId *big.Int) ([]byte, error) {
	return ownerAllowlist.abi.Pack("disallowAllTeeMachineOwners", extensionId)
}

// PackDisallowAllTeeWalletProjectOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x919fd33a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function disallowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (ownerAllowlist *OwnerAllowlist) PackDisallowAllTeeWalletProjectOwners(extensionId *big.Int) []byte {
	enc, err := ownerAllowlist.abi.Pack("disallowAllTeeWalletProjectOwners", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDisallowAllTeeWalletProjectOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x919fd33a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function disallowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (ownerAllowlist *OwnerAllowlist) TryPackDisallowAllTeeWalletProjectOwners(extensionId *big.Int) ([]byte, error) {
	return ownerAllowlist.abi.Pack("disallowAllTeeWalletProjectOwners", extensionId)
}

// PackGetAllowedExtensionOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6affbc67.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAllowedExtensionOwners() view returns(address[] _allowedOwners)
func (ownerAllowlist *OwnerAllowlist) PackGetAllowedExtensionOwners() []byte {
	enc, err := ownerAllowlist.abi.Pack("getAllowedExtensionOwners")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAllowedExtensionOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6affbc67.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAllowedExtensionOwners() view returns(address[] _allowedOwners)
func (ownerAllowlist *OwnerAllowlist) TryPackGetAllowedExtensionOwners() ([]byte, error) {
	return ownerAllowlist.abi.Pack("getAllowedExtensionOwners")
}

// UnpackGetAllowedExtensionOwners is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6affbc67.
//
// Solidity: function getAllowedExtensionOwners() view returns(address[] _allowedOwners)
func (ownerAllowlist *OwnerAllowlist) UnpackGetAllowedExtensionOwners(data []byte) ([]common.Address, error) {
	out, err := ownerAllowlist.abi.Unpack("getAllowedExtensionOwners", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetAllowedTeeMachineOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f37ea30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAllowedTeeMachineOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (ownerAllowlist *OwnerAllowlist) PackGetAllowedTeeMachineOwners(extensionId *big.Int) []byte {
	enc, err := ownerAllowlist.abi.Pack("getAllowedTeeMachineOwners", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAllowedTeeMachineOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f37ea30.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAllowedTeeMachineOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (ownerAllowlist *OwnerAllowlist) TryPackGetAllowedTeeMachineOwners(extensionId *big.Int) ([]byte, error) {
	return ownerAllowlist.abi.Pack("getAllowedTeeMachineOwners", extensionId)
}

// UnpackGetAllowedTeeMachineOwners is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5f37ea30.
//
// Solidity: function getAllowedTeeMachineOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (ownerAllowlist *OwnerAllowlist) UnpackGetAllowedTeeMachineOwners(data []byte) ([]common.Address, error) {
	out, err := ownerAllowlist.abi.Unpack("getAllowedTeeMachineOwners", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetAllowedTeeWalletProjectOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd11340ed.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAllowedTeeWalletProjectOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (ownerAllowlist *OwnerAllowlist) PackGetAllowedTeeWalletProjectOwners(extensionId *big.Int) []byte {
	enc, err := ownerAllowlist.abi.Pack("getAllowedTeeWalletProjectOwners", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAllowedTeeWalletProjectOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd11340ed.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAllowedTeeWalletProjectOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (ownerAllowlist *OwnerAllowlist) TryPackGetAllowedTeeWalletProjectOwners(extensionId *big.Int) ([]byte, error) {
	return ownerAllowlist.abi.Pack("getAllowedTeeWalletProjectOwners", extensionId)
}

// UnpackGetAllowedTeeWalletProjectOwners is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd11340ed.
//
// Solidity: function getAllowedTeeWalletProjectOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (ownerAllowlist *OwnerAllowlist) UnpackGetAllowedTeeWalletProjectOwners(data []byte) ([]common.Address, error) {
	out, err := ownerAllowlist.abi.Unpack("getAllowedTeeWalletProjectOwners", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackIsAllowedExtensionOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3886a10.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isAllowedExtensionOwner(address _owner) view returns(bool _isAllowed)
func (ownerAllowlist *OwnerAllowlist) PackIsAllowedExtensionOwner(owner common.Address) []byte {
	enc, err := ownerAllowlist.abi.Pack("isAllowedExtensionOwner", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsAllowedExtensionOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3886a10.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isAllowedExtensionOwner(address _owner) view returns(bool _isAllowed)
func (ownerAllowlist *OwnerAllowlist) TryPackIsAllowedExtensionOwner(owner common.Address) ([]byte, error) {
	return ownerAllowlist.abi.Pack("isAllowedExtensionOwner", owner)
}

// UnpackIsAllowedExtensionOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd3886a10.
//
// Solidity: function isAllowedExtensionOwner(address _owner) view returns(bool _isAllowed)
func (ownerAllowlist *OwnerAllowlist) UnpackIsAllowedExtensionOwner(data []byte) (bool, error) {
	out, err := ownerAllowlist.abi.Unpack("isAllowedExtensionOwner", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsAllowedTeeMachineOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84f5da36.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isAllowedTeeMachineOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (ownerAllowlist *OwnerAllowlist) PackIsAllowedTeeMachineOwner(extensionId *big.Int, owner common.Address) []byte {
	enc, err := ownerAllowlist.abi.Pack("isAllowedTeeMachineOwner", extensionId, owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsAllowedTeeMachineOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84f5da36.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isAllowedTeeMachineOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (ownerAllowlist *OwnerAllowlist) TryPackIsAllowedTeeMachineOwner(extensionId *big.Int, owner common.Address) ([]byte, error) {
	return ownerAllowlist.abi.Pack("isAllowedTeeMachineOwner", extensionId, owner)
}

// UnpackIsAllowedTeeMachineOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x84f5da36.
//
// Solidity: function isAllowedTeeMachineOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (ownerAllowlist *OwnerAllowlist) UnpackIsAllowedTeeMachineOwner(data []byte) (bool, error) {
	out, err := ownerAllowlist.abi.Unpack("isAllowedTeeMachineOwner", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsAllowedTeeWalletProjectOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab1c7ea1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isAllowedTeeWalletProjectOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (ownerAllowlist *OwnerAllowlist) PackIsAllowedTeeWalletProjectOwner(extensionId *big.Int, owner common.Address) []byte {
	enc, err := ownerAllowlist.abi.Pack("isAllowedTeeWalletProjectOwner", extensionId, owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsAllowedTeeWalletProjectOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab1c7ea1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isAllowedTeeWalletProjectOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (ownerAllowlist *OwnerAllowlist) TryPackIsAllowedTeeWalletProjectOwner(extensionId *big.Int, owner common.Address) ([]byte, error) {
	return ownerAllowlist.abi.Pack("isAllowedTeeWalletProjectOwner", extensionId, owner)
}

// UnpackIsAllowedTeeWalletProjectOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xab1c7ea1.
//
// Solidity: function isAllowedTeeWalletProjectOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (ownerAllowlist *OwnerAllowlist) UnpackIsAllowedTeeWalletProjectOwner(data []byte) (bool, error) {
	out, err := ownerAllowlist.abi.Unpack("isAllowedTeeWalletProjectOwner", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRemoveAllowedExtensionOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d9ef68a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeAllowedExtensionOwners(address[] _owners) returns()
func (ownerAllowlist *OwnerAllowlist) PackRemoveAllowedExtensionOwners(owners []common.Address) []byte {
	enc, err := ownerAllowlist.abi.Pack("removeAllowedExtensionOwners", owners)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveAllowedExtensionOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d9ef68a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeAllowedExtensionOwners(address[] _owners) returns()
func (ownerAllowlist *OwnerAllowlist) TryPackRemoveAllowedExtensionOwners(owners []common.Address) ([]byte, error) {
	return ownerAllowlist.abi.Pack("removeAllowedExtensionOwners", owners)
}

// PackRemoveAllowedTeeMachineOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d371c3c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (ownerAllowlist *OwnerAllowlist) PackRemoveAllowedTeeMachineOwners(extensionId *big.Int, owners []common.Address) []byte {
	enc, err := ownerAllowlist.abi.Pack("removeAllowedTeeMachineOwners", extensionId, owners)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveAllowedTeeMachineOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d371c3c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (ownerAllowlist *OwnerAllowlist) TryPackRemoveAllowedTeeMachineOwners(extensionId *big.Int, owners []common.Address) ([]byte, error) {
	return ownerAllowlist.abi.Pack("removeAllowedTeeMachineOwners", extensionId, owners)
}

// PackRemoveAllowedTeeWalletProjectOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62a2513d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (ownerAllowlist *OwnerAllowlist) PackRemoveAllowedTeeWalletProjectOwners(extensionId *big.Int, owners []common.Address) []byte {
	enc, err := ownerAllowlist.abi.Pack("removeAllowedTeeWalletProjectOwners", extensionId, owners)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveAllowedTeeWalletProjectOwners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62a2513d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (ownerAllowlist *OwnerAllowlist) TryPackRemoveAllowedTeeWalletProjectOwners(extensionId *big.Int, owners []common.Address) ([]byte, error) {
	return ownerAllowlist.abi.Pack("removeAllowedTeeWalletProjectOwners", extensionId, owners)
}

// OwnerAllowlistAllExtensionOwnersAllowed represents a AllExtensionOwnersAllowed event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllExtensionOwnersAllowed struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistAllExtensionOwnersAllowedEventName = "AllExtensionOwnersAllowed"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistAllExtensionOwnersAllowed) ContractEventName() string {
	return OwnerAllowlistAllExtensionOwnersAllowedEventName
}

// UnpackAllExtensionOwnersAllowedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllExtensionOwnersAllowed()
func (ownerAllowlist *OwnerAllowlist) UnpackAllExtensionOwnersAllowedEvent(log *types.Log) (*OwnerAllowlistAllExtensionOwnersAllowed, error) {
	event := "AllExtensionOwnersAllowed"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistAllExtensionOwnersAllowed)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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

// OwnerAllowlistAllExtensionOwnersDisallowed represents a AllExtensionOwnersDisallowed event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllExtensionOwnersDisallowed struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistAllExtensionOwnersDisallowedEventName = "AllExtensionOwnersDisallowed"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistAllExtensionOwnersDisallowed) ContractEventName() string {
	return OwnerAllowlistAllExtensionOwnersDisallowedEventName
}

// UnpackAllExtensionOwnersDisallowedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllExtensionOwnersDisallowed()
func (ownerAllowlist *OwnerAllowlist) UnpackAllExtensionOwnersDisallowedEvent(log *types.Log) (*OwnerAllowlistAllExtensionOwnersDisallowed, error) {
	event := "AllExtensionOwnersDisallowed"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistAllExtensionOwnersDisallowed)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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

// OwnerAllowlistAllTeeMachineOwnersAllowed represents a AllTeeMachineOwnersAllowed event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllTeeMachineOwnersAllowed struct {
	ExtensionId *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistAllTeeMachineOwnersAllowedEventName = "AllTeeMachineOwnersAllowed"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistAllTeeMachineOwnersAllowed) ContractEventName() string {
	return OwnerAllowlistAllTeeMachineOwnersAllowedEventName
}

// UnpackAllTeeMachineOwnersAllowedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllTeeMachineOwnersAllowed(uint256 indexed extensionId)
func (ownerAllowlist *OwnerAllowlist) UnpackAllTeeMachineOwnersAllowedEvent(log *types.Log) (*OwnerAllowlistAllTeeMachineOwnersAllowed, error) {
	event := "AllTeeMachineOwnersAllowed"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistAllTeeMachineOwnersAllowed)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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

// OwnerAllowlistAllTeeMachineOwnersDisallowed represents a AllTeeMachineOwnersDisallowed event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllTeeMachineOwnersDisallowed struct {
	ExtensionId *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistAllTeeMachineOwnersDisallowedEventName = "AllTeeMachineOwnersDisallowed"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistAllTeeMachineOwnersDisallowed) ContractEventName() string {
	return OwnerAllowlistAllTeeMachineOwnersDisallowedEventName
}

// UnpackAllTeeMachineOwnersDisallowedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllTeeMachineOwnersDisallowed(uint256 indexed extensionId)
func (ownerAllowlist *OwnerAllowlist) UnpackAllTeeMachineOwnersDisallowedEvent(log *types.Log) (*OwnerAllowlistAllTeeMachineOwnersDisallowed, error) {
	event := "AllTeeMachineOwnersDisallowed"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistAllTeeMachineOwnersDisallowed)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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

// OwnerAllowlistAllTeeWalletProjectOwnersAllowed represents a AllTeeWalletProjectOwnersAllowed event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllTeeWalletProjectOwnersAllowed struct {
	ExtensionId *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistAllTeeWalletProjectOwnersAllowedEventName = "AllTeeWalletProjectOwnersAllowed"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistAllTeeWalletProjectOwnersAllowed) ContractEventName() string {
	return OwnerAllowlistAllTeeWalletProjectOwnersAllowedEventName
}

// UnpackAllTeeWalletProjectOwnersAllowedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllTeeWalletProjectOwnersAllowed(uint256 indexed extensionId)
func (ownerAllowlist *OwnerAllowlist) UnpackAllTeeWalletProjectOwnersAllowedEvent(log *types.Log) (*OwnerAllowlistAllTeeWalletProjectOwnersAllowed, error) {
	event := "AllTeeWalletProjectOwnersAllowed"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistAllTeeWalletProjectOwnersAllowed)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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

// OwnerAllowlistAllTeeWalletProjectOwnersDisallowed represents a AllTeeWalletProjectOwnersDisallowed event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllTeeWalletProjectOwnersDisallowed struct {
	ExtensionId *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistAllTeeWalletProjectOwnersDisallowedEventName = "AllTeeWalletProjectOwnersDisallowed"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistAllTeeWalletProjectOwnersDisallowed) ContractEventName() string {
	return OwnerAllowlistAllTeeWalletProjectOwnersDisallowedEventName
}

// UnpackAllTeeWalletProjectOwnersDisallowedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllTeeWalletProjectOwnersDisallowed(uint256 indexed extensionId)
func (ownerAllowlist *OwnerAllowlist) UnpackAllTeeWalletProjectOwnersDisallowedEvent(log *types.Log) (*OwnerAllowlistAllTeeWalletProjectOwnersDisallowed, error) {
	event := "AllTeeWalletProjectOwnersDisallowed"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistAllTeeWalletProjectOwnersDisallowed)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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

// OwnerAllowlistAllowedExtensionOwnersAdded represents a AllowedExtensionOwnersAdded event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedExtensionOwnersAdded struct {
	Owners []common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistAllowedExtensionOwnersAddedEventName = "AllowedExtensionOwnersAdded"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistAllowedExtensionOwnersAdded) ContractEventName() string {
	return OwnerAllowlistAllowedExtensionOwnersAddedEventName
}

// UnpackAllowedExtensionOwnersAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllowedExtensionOwnersAdded(address[] owners)
func (ownerAllowlist *OwnerAllowlist) UnpackAllowedExtensionOwnersAddedEvent(log *types.Log) (*OwnerAllowlistAllowedExtensionOwnersAdded, error) {
	event := "AllowedExtensionOwnersAdded"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistAllowedExtensionOwnersAdded)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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

// OwnerAllowlistAllowedExtensionOwnersRemoved represents a AllowedExtensionOwnersRemoved event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedExtensionOwnersRemoved struct {
	Owners []common.Address
	Raw    *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistAllowedExtensionOwnersRemovedEventName = "AllowedExtensionOwnersRemoved"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistAllowedExtensionOwnersRemoved) ContractEventName() string {
	return OwnerAllowlistAllowedExtensionOwnersRemovedEventName
}

// UnpackAllowedExtensionOwnersRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllowedExtensionOwnersRemoved(address[] owners)
func (ownerAllowlist *OwnerAllowlist) UnpackAllowedExtensionOwnersRemovedEvent(log *types.Log) (*OwnerAllowlistAllowedExtensionOwnersRemoved, error) {
	event := "AllowedExtensionOwnersRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistAllowedExtensionOwnersRemoved)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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

// OwnerAllowlistAllowedTeeMachineOwnersAdded represents a AllowedTeeMachineOwnersAdded event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedTeeMachineOwnersAdded struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistAllowedTeeMachineOwnersAddedEventName = "AllowedTeeMachineOwnersAdded"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistAllowedTeeMachineOwnersAdded) ContractEventName() string {
	return OwnerAllowlistAllowedTeeMachineOwnersAddedEventName
}

// UnpackAllowedTeeMachineOwnersAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllowedTeeMachineOwnersAdded(uint256 indexed extensionId, address[] owners)
func (ownerAllowlist *OwnerAllowlist) UnpackAllowedTeeMachineOwnersAddedEvent(log *types.Log) (*OwnerAllowlistAllowedTeeMachineOwnersAdded, error) {
	event := "AllowedTeeMachineOwnersAdded"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistAllowedTeeMachineOwnersAdded)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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

// OwnerAllowlistAllowedTeeMachineOwnersRemoved represents a AllowedTeeMachineOwnersRemoved event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedTeeMachineOwnersRemoved struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistAllowedTeeMachineOwnersRemovedEventName = "AllowedTeeMachineOwnersRemoved"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistAllowedTeeMachineOwnersRemoved) ContractEventName() string {
	return OwnerAllowlistAllowedTeeMachineOwnersRemovedEventName
}

// UnpackAllowedTeeMachineOwnersRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllowedTeeMachineOwnersRemoved(uint256 indexed extensionId, address[] owners)
func (ownerAllowlist *OwnerAllowlist) UnpackAllowedTeeMachineOwnersRemovedEvent(log *types.Log) (*OwnerAllowlistAllowedTeeMachineOwnersRemoved, error) {
	event := "AllowedTeeMachineOwnersRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistAllowedTeeMachineOwnersRemoved)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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

// OwnerAllowlistAllowedTeeWalletProjectOwnersAdded represents a AllowedTeeWalletProjectOwnersAdded event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedTeeWalletProjectOwnersAdded struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistAllowedTeeWalletProjectOwnersAddedEventName = "AllowedTeeWalletProjectOwnersAdded"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistAllowedTeeWalletProjectOwnersAdded) ContractEventName() string {
	return OwnerAllowlistAllowedTeeWalletProjectOwnersAddedEventName
}

// UnpackAllowedTeeWalletProjectOwnersAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllowedTeeWalletProjectOwnersAdded(uint256 indexed extensionId, address[] owners)
func (ownerAllowlist *OwnerAllowlist) UnpackAllowedTeeWalletProjectOwnersAddedEvent(log *types.Log) (*OwnerAllowlistAllowedTeeWalletProjectOwnersAdded, error) {
	event := "AllowedTeeWalletProjectOwnersAdded"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistAllowedTeeWalletProjectOwnersAdded)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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

// OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved represents a AllowedTeeWalletProjectOwnersRemoved event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistAllowedTeeWalletProjectOwnersRemovedEventName = "AllowedTeeWalletProjectOwnersRemoved"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved) ContractEventName() string {
	return OwnerAllowlistAllowedTeeWalletProjectOwnersRemovedEventName
}

// UnpackAllowedTeeWalletProjectOwnersRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllowedTeeWalletProjectOwnersRemoved(uint256 indexed extensionId, address[] owners)
func (ownerAllowlist *OwnerAllowlist) UnpackAllowedTeeWalletProjectOwnersRemovedEvent(log *types.Log) (*OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved, error) {
	event := "AllowedTeeWalletProjectOwnersRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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

// OwnerAllowlistGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the OwnerAllowlist contract.
type OwnerAllowlistGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistGovernanceCallTimelocked) ContractEventName() string {
	return OwnerAllowlistGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (ownerAllowlist *OwnerAllowlist) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*OwnerAllowlistGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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

// OwnerAllowlistInitialized represents a Initialized event raised by the OwnerAllowlist contract.
type OwnerAllowlistInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const OwnerAllowlistInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (OwnerAllowlistInitialized) ContractEventName() string {
	return OwnerAllowlistInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (ownerAllowlist *OwnerAllowlist) UnpackInitializedEvent(log *types.Log) (*OwnerAllowlistInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != ownerAllowlist.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OwnerAllowlistInitialized)
	if len(log.Data) > 0 {
		if err := ownerAllowlist.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range ownerAllowlist.abi.Events[event].Inputs {
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
func (ownerAllowlist *OwnerAllowlist) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["AddressAlreadyInSet"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackAddressAlreadyInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["AddressNotInSet"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackAddressNotInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["AvailabilityCheckTimestampInvalid"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackAvailabilityCheckTimestampInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["DuplicatedCosigner"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackDuplicatedCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["InvalidAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackInvalidAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["InvalidKeyType"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackInvalidKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["InvalidPublicKey"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackInvalidPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["InvalidResponseData"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackInvalidResponseDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["InvalidSigningAlgo"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackInvalidSigningAlgoError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["InvalidWalletStatus"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackInvalidWalletStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["KeyTypeNotSupported"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["NoAddresses"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackNoAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["NotOwnerOrPauser"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackNotOwnerOrPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["NotOwnerOrUnpauser"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackNotOwnerOrUnpauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["OnlyExtensionOwner"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackOnlyExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackOnlyOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["OnlyOwnerOrBackupManager"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackOnlyOwnerOrBackupManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["OnlyProposedOwner"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackOnlyProposedOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["OwnerNotAllowed"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackOwnerNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], ownerAllowlist.abi.Errors["VersionNotSupported"].ID.Bytes()[:4]) {
		return ownerAllowlist.UnpackVersionNotSupportedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// OwnerAllowlistAddressAlreadyInSet represents a AddressAlreadyInSet error raised by the OwnerAllowlist contract.
type OwnerAllowlistAddressAlreadyInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressAlreadyInSet(address addr)
func OwnerAllowlistAddressAlreadyInSetErrorID() common.Hash {
	return common.HexToHash("0xb725c7fb1c3f0f065973d6b9eb55140f970259955af3045742cd5b37902eb920")
}

// UnpackAddressAlreadyInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressAlreadyInSet(address addr)
func (ownerAllowlist *OwnerAllowlist) UnpackAddressAlreadyInSetError(raw []byte) (*OwnerAllowlistAddressAlreadyInSet, error) {
	out := new(OwnerAllowlistAddressAlreadyInSet)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "AddressAlreadyInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistAddressNotInSet represents a AddressNotInSet error raised by the OwnerAllowlist contract.
type OwnerAllowlistAddressNotInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressNotInSet(address addr)
func OwnerAllowlistAddressNotInSetErrorID() common.Hash {
	return common.HexToHash("0x1ff5ecf3e6d385ea01ff0323fe894ed6d25a83ed5f6ff2a2ede4c897408f6be9")
}

// UnpackAddressNotInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressNotInSet(address addr)
func (ownerAllowlist *OwnerAllowlist) UnpackAddressNotInSetError(raw []byte) (*OwnerAllowlistAddressNotInSet, error) {
	out := new(OwnerAllowlistAddressNotInSet)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "AddressNotInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistAvailabilityCheckTimestampInvalid represents a AvailabilityCheckTimestampInvalid error raised by the OwnerAllowlist contract.
type OwnerAllowlistAvailabilityCheckTimestampInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func OwnerAllowlistAvailabilityCheckTimestampInvalidErrorID() common.Hash {
	return common.HexToHash("0x6344a5e74cff78612fa271fa907ca608af3bad10e1e8c8df302e29ebc6298e2d")
}

// UnpackAvailabilityCheckTimestampInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func (ownerAllowlist *OwnerAllowlist) UnpackAvailabilityCheckTimestampInvalidError(raw []byte) (*OwnerAllowlistAvailabilityCheckTimestampInvalid, error) {
	out := new(OwnerAllowlistAvailabilityCheckTimestampInvalid)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "AvailabilityCheckTimestampInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistDuplicatedCosigner represents a DuplicatedCosigner error raised by the OwnerAllowlist contract.
type OwnerAllowlistDuplicatedCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func OwnerAllowlistDuplicatedCosignerErrorID() common.Hash {
	return common.HexToHash("0x53a5890417576fb3043dc3ff1669b0932da2f657d8adb4c11313b12b711f5910")
}

// UnpackDuplicatedCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func (ownerAllowlist *OwnerAllowlist) UnpackDuplicatedCosignerError(raw []byte) (*OwnerAllowlistDuplicatedCosigner, error) {
	out := new(OwnerAllowlistDuplicatedCosigner)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "DuplicatedCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistExtensionIdMismatch represents a ExtensionIdMismatch error raised by the OwnerAllowlist contract.
type OwnerAllowlistExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func OwnerAllowlistExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (ownerAllowlist *OwnerAllowlist) UnpackExtensionIdMismatchError(raw []byte) (*OwnerAllowlistExtensionIdMismatch, error) {
	out := new(OwnerAllowlistExtensionIdMismatch)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistInvalidAddress represents a InvalidAddress error raised by the OwnerAllowlist contract.
type OwnerAllowlistInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func OwnerAllowlistInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (ownerAllowlist *OwnerAllowlist) UnpackInvalidAddressError(raw []byte) (*OwnerAllowlistInvalidAddress, error) {
	out := new(OwnerAllowlistInvalidAddress)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistInvalidAvailabilityCheckStatus represents a InvalidAvailabilityCheckStatus error raised by the OwnerAllowlist contract.
type OwnerAllowlistInvalidAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func OwnerAllowlistInvalidAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0xedee4ff00b842d3ef84031c60001065c540c838c5479bbf0ae56f15868427e44")
}

// UnpackInvalidAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func (ownerAllowlist *OwnerAllowlist) UnpackInvalidAvailabilityCheckStatusError(raw []byte) (*OwnerAllowlistInvalidAvailabilityCheckStatus, error) {
	out := new(OwnerAllowlistInvalidAvailabilityCheckStatus)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "InvalidAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistInvalidCosigner represents a InvalidCosigner error raised by the OwnerAllowlist contract.
type OwnerAllowlistInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func OwnerAllowlistInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (ownerAllowlist *OwnerAllowlist) UnpackInvalidCosignerError(raw []byte) (*OwnerAllowlistInvalidCosigner, error) {
	out := new(OwnerAllowlistInvalidCosigner)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistInvalidDuration represents a InvalidDuration error raised by the OwnerAllowlist contract.
type OwnerAllowlistInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func OwnerAllowlistInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (ownerAllowlist *OwnerAllowlist) UnpackInvalidDurationError(raw []byte) (*OwnerAllowlistInvalidDuration, error) {
	out := new(OwnerAllowlistInvalidDuration)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the OwnerAllowlist contract.
type OwnerAllowlistInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func OwnerAllowlistInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (ownerAllowlist *OwnerAllowlist) UnpackInvalidGovernanceHashError(raw []byte) (*OwnerAllowlistInvalidGovernanceHash, error) {
	out := new(OwnerAllowlistInvalidGovernanceHash)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistInvalidInitialization represents a InvalidInitialization error raised by the OwnerAllowlist contract.
type OwnerAllowlistInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func OwnerAllowlistInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (ownerAllowlist *OwnerAllowlist) UnpackInvalidInitializationError(raw []byte) (*OwnerAllowlistInvalidInitialization, error) {
	out := new(OwnerAllowlistInvalidInitialization)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistInvalidKeyType represents a InvalidKeyType error raised by the OwnerAllowlist contract.
type OwnerAllowlistInvalidKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyType()
func OwnerAllowlistInvalidKeyTypeErrorID() common.Hash {
	return common.HexToHash("0x36dccb440beb343622d8ee0927bce3f69281feea3a5b0610f6e00e2811b557e1")
}

// UnpackInvalidKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyType()
func (ownerAllowlist *OwnerAllowlist) UnpackInvalidKeyTypeError(raw []byte) (*OwnerAllowlistInvalidKeyType, error) {
	out := new(OwnerAllowlistInvalidKeyType)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "InvalidKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistInvalidNonce represents a InvalidNonce error raised by the OwnerAllowlist contract.
type OwnerAllowlistInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func OwnerAllowlistInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (ownerAllowlist *OwnerAllowlist) UnpackInvalidNonceError(raw []byte) (*OwnerAllowlistInvalidNonce, error) {
	out := new(OwnerAllowlistInvalidNonce)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistInvalidPublicKey represents a InvalidPublicKey error raised by the OwnerAllowlist contract.
type OwnerAllowlistInvalidPublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPublicKey()
func OwnerAllowlistInvalidPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa2d0fee854985cb3d3a394146b5df1c41d5b4d6fcede7fc57a3c05a49b2f3e73")
}

// UnpackInvalidPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPublicKey()
func (ownerAllowlist *OwnerAllowlist) UnpackInvalidPublicKeyError(raw []byte) (*OwnerAllowlistInvalidPublicKey, error) {
	out := new(OwnerAllowlistInvalidPublicKey)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "InvalidPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistInvalidResponseData represents a InvalidResponseData error raised by the OwnerAllowlist contract.
type OwnerAllowlistInvalidResponseData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseData()
func OwnerAllowlistInvalidResponseDataErrorID() common.Hash {
	return common.HexToHash("0x48566617128355a033880805d6462da9bf9c6fe4c61b922997451a71262990af")
}

// UnpackInvalidResponseDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseData()
func (ownerAllowlist *OwnerAllowlist) UnpackInvalidResponseDataError(raw []byte) (*OwnerAllowlistInvalidResponseData, error) {
	out := new(OwnerAllowlistInvalidResponseData)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "InvalidResponseData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistInvalidSigningAlgo represents a InvalidSigningAlgo error raised by the OwnerAllowlist contract.
type OwnerAllowlistInvalidSigningAlgo struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningAlgo()
func OwnerAllowlistInvalidSigningAlgoErrorID() common.Hash {
	return common.HexToHash("0x9b7198a697b61e546ef9ee593a6f77d2bf9f1df9dcbba655ce96d7e2c75346c7")
}

// UnpackInvalidSigningAlgoError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningAlgo()
func (ownerAllowlist *OwnerAllowlist) UnpackInvalidSigningAlgoError(raw []byte) (*OwnerAllowlistInvalidSigningAlgo, error) {
	out := new(OwnerAllowlistInvalidSigningAlgo)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "InvalidSigningAlgo", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistInvalidThreshold represents a InvalidThreshold error raised by the OwnerAllowlist contract.
type OwnerAllowlistInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func OwnerAllowlistInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (ownerAllowlist *OwnerAllowlist) UnpackInvalidThresholdError(raw []byte) (*OwnerAllowlistInvalidThreshold, error) {
	out := new(OwnerAllowlistInvalidThreshold)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistInvalidWalletStatus represents a InvalidWalletStatus error raised by the OwnerAllowlist contract.
type OwnerAllowlistInvalidWalletStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidWalletStatus()
func OwnerAllowlistInvalidWalletStatusErrorID() common.Hash {
	return common.HexToHash("0xd700120e4896ea7835a5e9e4feb690644ced5cad1c52209afdcbe06b96d91345")
}

// UnpackInvalidWalletStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidWalletStatus()
func (ownerAllowlist *OwnerAllowlist) UnpackInvalidWalletStatusError(raw []byte) (*OwnerAllowlistInvalidWalletStatus, error) {
	out := new(OwnerAllowlistInvalidWalletStatus)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "InvalidWalletStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistKeyTypeNotSupported represents a KeyTypeNotSupported error raised by the OwnerAllowlist contract.
type OwnerAllowlistKeyTypeNotSupported struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func OwnerAllowlistKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x85004df1f47814c2de6fbf03d3f2b1ca8c67c182d0d54278766543b6be3bc603")
}

// UnpackKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func (ownerAllowlist *OwnerAllowlist) UnpackKeyTypeNotSupportedError(raw []byte) (*OwnerAllowlistKeyTypeNotSupported, error) {
	out := new(OwnerAllowlistKeyTypeNotSupported)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "KeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistLengthsMismatch represents a LengthsMismatch error raised by the OwnerAllowlist contract.
type OwnerAllowlistLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func OwnerAllowlistLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (ownerAllowlist *OwnerAllowlist) UnpackLengthsMismatchError(raw []byte) (*OwnerAllowlistLengthsMismatch, error) {
	out := new(OwnerAllowlistLengthsMismatch)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistNoAddresses represents a NoAddresses error raised by the OwnerAllowlist contract.
type OwnerAllowlistNoAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoAddresses()
func OwnerAllowlistNoAddressesErrorID() common.Hash {
	return common.HexToHash("0xb2e79f05746fa6c9f3d61f9efe85959371cde5030a6c2dd70c92b63578c5cde7")
}

// UnpackNoAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoAddresses()
func (ownerAllowlist *OwnerAllowlist) UnpackNoAddressesError(raw []byte) (*OwnerAllowlistNoAddresses, error) {
	out := new(OwnerAllowlistNoAddresses)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "NoAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistNotInitializing represents a NotInitializing error raised by the OwnerAllowlist contract.
type OwnerAllowlistNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func OwnerAllowlistNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (ownerAllowlist *OwnerAllowlist) UnpackNotInitializingError(raw []byte) (*OwnerAllowlistNotInitializing, error) {
	out := new(OwnerAllowlistNotInitializing)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistNotOwnerOrPauser represents a NotOwnerOrPauser error raised by the OwnerAllowlist contract.
type OwnerAllowlistNotOwnerOrPauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrPauser(address caller)
func OwnerAllowlistNotOwnerOrPauserErrorID() common.Hash {
	return common.HexToHash("0x46c0d9af6b895093e022ad2984c54e60192e6e26c5063daf232e825c6e07ff9a")
}

// UnpackNotOwnerOrPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrPauser(address caller)
func (ownerAllowlist *OwnerAllowlist) UnpackNotOwnerOrPauserError(raw []byte) (*OwnerAllowlistNotOwnerOrPauser, error) {
	out := new(OwnerAllowlistNotOwnerOrPauser)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "NotOwnerOrPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistNotOwnerOrUnpauser represents a NotOwnerOrUnpauser error raised by the OwnerAllowlist contract.
type OwnerAllowlistNotOwnerOrUnpauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func OwnerAllowlistNotOwnerOrUnpauserErrorID() common.Hash {
	return common.HexToHash("0x81e700e8f0bd5ba4a6d6a2ce9a1ad0ac9bb1b1b1e62d921a55cbc4b53f1da649")
}

// UnpackNotOwnerOrUnpauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func (ownerAllowlist *OwnerAllowlist) UnpackNotOwnerOrUnpauserError(raw []byte) (*OwnerAllowlistNotOwnerOrUnpauser, error) {
	out := new(OwnerAllowlistNotOwnerOrUnpauser)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "NotOwnerOrUnpauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistOnlyExtensionOwner represents a OnlyExtensionOwner error raised by the OwnerAllowlist contract.
type OwnerAllowlistOnlyExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwner()
func OwnerAllowlistOnlyExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xcd49fd1de5a297cd1ce055d5abffcf2cb25ad9492bd39a703a238af57b2a9488")
}

// UnpackOnlyExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwner()
func (ownerAllowlist *OwnerAllowlist) UnpackOnlyExtensionOwnerError(raw []byte) (*OwnerAllowlistOnlyExtensionOwner, error) {
	out := new(OwnerAllowlistOnlyExtensionOwner)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "OnlyExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the OwnerAllowlist contract.
type OwnerAllowlistOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func OwnerAllowlistOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (ownerAllowlist *OwnerAllowlist) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*OwnerAllowlistOnlyExtensionOwnerOrOperator, error) {
	out := new(OwnerAllowlistOnlyExtensionOwnerOrOperator)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistOnlyGovernance represents a OnlyGovernance error raised by the OwnerAllowlist contract.
type OwnerAllowlistOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func OwnerAllowlistOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (ownerAllowlist *OwnerAllowlist) UnpackOnlyGovernanceError(raw []byte) (*OwnerAllowlistOnlyGovernance, error) {
	out := new(OwnerAllowlistOnlyGovernance)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistOnlyOwner represents a OnlyOwner error raised by the OwnerAllowlist contract.
type OwnerAllowlistOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func OwnerAllowlistOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (ownerAllowlist *OwnerAllowlist) UnpackOnlyOwnerError(raw []byte) (*OwnerAllowlistOnlyOwner, error) {
	out := new(OwnerAllowlistOnlyOwner)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistOnlyOwnerOrBackupManager represents a OnlyOwnerOrBackupManager error raised by the OwnerAllowlist contract.
type OwnerAllowlistOnlyOwnerOrBackupManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrBackupManager()
func OwnerAllowlistOnlyOwnerOrBackupManagerErrorID() common.Hash {
	return common.HexToHash("0x249a6cf6e22092f9c80798e2a0c583c0f36aea0779ef38a5b5cdb6a112a0e6df")
}

// UnpackOnlyOwnerOrBackupManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrBackupManager()
func (ownerAllowlist *OwnerAllowlist) UnpackOnlyOwnerOrBackupManagerError(raw []byte) (*OwnerAllowlistOnlyOwnerOrBackupManager, error) {
	out := new(OwnerAllowlistOnlyOwnerOrBackupManager)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "OnlyOwnerOrBackupManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the OwnerAllowlist contract.
type OwnerAllowlistOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func OwnerAllowlistOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (ownerAllowlist *OwnerAllowlist) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*OwnerAllowlistOnlyProductionOrPausedStatus, error) {
	out := new(OwnerAllowlistOnlyProductionOrPausedStatus)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistOnlyProposedOwner represents a OnlyProposedOwner error raised by the OwnerAllowlist contract.
type OwnerAllowlistOnlyProposedOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProposedOwner()
func OwnerAllowlistOnlyProposedOwnerErrorID() common.Hash {
	return common.HexToHash("0x49cb15061ee558e74f66f937f740fd2a3791a89d2d139862222e5c97cabc02cd")
}

// UnpackOnlyProposedOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProposedOwner()
func (ownerAllowlist *OwnerAllowlist) UnpackOnlyProposedOwnerError(raw []byte) (*OwnerAllowlistOnlyProposedOwner, error) {
	out := new(OwnerAllowlistOnlyProposedOwner)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "OnlyProposedOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistOwnerNotAllowed represents a OwnerNotAllowed error raised by the OwnerAllowlist contract.
type OwnerAllowlistOwnerNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerNotAllowed()
func OwnerAllowlistOwnerNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x3d5c9996545c6bf4287c26c1c2d8e6b4c39b92f19ba706c3a756720bf9c4fd24")
}

// UnpackOwnerNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerNotAllowed()
func (ownerAllowlist *OwnerAllowlist) UnpackOwnerNotAllowedError(raw []byte) (*OwnerAllowlistOwnerNotAllowed, error) {
	out := new(OwnerAllowlistOwnerNotAllowed)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "OwnerNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the OwnerAllowlist contract.
type OwnerAllowlistTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func OwnerAllowlistTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (ownerAllowlist *OwnerAllowlist) UnpackTeeMachineNotAvailableError(raw []byte) (*OwnerAllowlistTeeMachineNotAvailable, error) {
	out := new(OwnerAllowlistTeeMachineNotAvailable)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerAllowlistVersionNotSupported represents a VersionNotSupported error raised by the OwnerAllowlist contract.
type OwnerAllowlistVersionNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionNotSupported()
func OwnerAllowlistVersionNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x6db0b17b06266b37e721024a8ba1221439799ebb779dfd397865e37250abc9a8")
}

// UnpackVersionNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionNotSupported()
func (ownerAllowlist *OwnerAllowlist) UnpackVersionNotSupportedError(raw []byte) (*OwnerAllowlistVersionNotSupported, error) {
	out := new(OwnerAllowlistVersionNotSupported)
	if err := ownerAllowlist.abi.UnpackIntoInterface(out, "VersionNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}
