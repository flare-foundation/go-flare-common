// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package machineemergencypause

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

// MachineEmergencyPauseMetaData contains all meta data concerning the MachineEmergencyPause contract.
var MachineEmergencyPauseMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"EmergencyPauseActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"EmergencyProtectionActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"ExtensionAlreadyEmergencyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"ExtensionNotEmergencyPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"graceSeconds\",\"type\":\"uint256\"}],\"name\":\"GracePeriodTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"graceSeconds\",\"type\":\"uint256\"}],\"name\":\"GracePeriodTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"graceSeconds\",\"type\":\"uint256\"}],\"name\":\"EmergencyUnpauseGracePeriodSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"ExtensionEmergencyPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"ExtensionEmergencyPausersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"ExtensionEmergencyPausersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"unpauseTs\",\"type\":\"uint64\"}],\"name\":\"ExtensionEmergencyUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"ExtensionEmergencyUnpausersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"ExtensionEmergencyUnpausersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addExtensionEmergencyPausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addExtensionEmergencyUnpausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"emergencyPauseExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"emergencyUnpauseExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEmergencyUnpauseGracePeriodSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_seconds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getExtensionEmergencyPausers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getExtensionEmergencyUnpausers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getLastUnpauseTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_ts\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"isExtensionEmergencyPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isExtensionEmergencyPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isPauser\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isExtensionEmergencyUnpauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isUnpauser\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"removeExtensionEmergencyPausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"removeExtensionEmergencyUnpausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seconds\",\"type\":\"uint256\"}],\"name\":\"setEmergencyUnpauseGracePeriodSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "MachineEmergencyPause",
}

// MachineEmergencyPause is an auto generated Go binding around an Ethereum contract.
type MachineEmergencyPause struct {
	abi abi.ABI
}

// NewMachineEmergencyPause creates a new instance of MachineEmergencyPause.
func NewMachineEmergencyPause() *MachineEmergencyPause {
	parsed, err := MachineEmergencyPauseMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MachineEmergencyPause{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MachineEmergencyPause) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddExtensionEmergencyPausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6d02fabc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addExtensionEmergencyPausers(uint256 _extensionId, address[] _addresses) returns()
func (machineEmergencyPause *MachineEmergencyPause) PackAddExtensionEmergencyPausers(extensionId *big.Int, addresses []common.Address) []byte {
	enc, err := machineEmergencyPause.abi.Pack("addExtensionEmergencyPausers", extensionId, addresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddExtensionEmergencyPausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6d02fabc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addExtensionEmergencyPausers(uint256 _extensionId, address[] _addresses) returns()
func (machineEmergencyPause *MachineEmergencyPause) TryPackAddExtensionEmergencyPausers(extensionId *big.Int, addresses []common.Address) ([]byte, error) {
	return machineEmergencyPause.abi.Pack("addExtensionEmergencyPausers", extensionId, addresses)
}

// PackAddExtensionEmergencyUnpausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2c21c0e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addExtensionEmergencyUnpausers(uint256 _extensionId, address[] _addresses) returns()
func (machineEmergencyPause *MachineEmergencyPause) PackAddExtensionEmergencyUnpausers(extensionId *big.Int, addresses []common.Address) []byte {
	enc, err := machineEmergencyPause.abi.Pack("addExtensionEmergencyUnpausers", extensionId, addresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddExtensionEmergencyUnpausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2c21c0e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addExtensionEmergencyUnpausers(uint256 _extensionId, address[] _addresses) returns()
func (machineEmergencyPause *MachineEmergencyPause) TryPackAddExtensionEmergencyUnpausers(extensionId *big.Int, addresses []common.Address) ([]byte, error) {
	return machineEmergencyPause.abi.Pack("addExtensionEmergencyUnpausers", extensionId, addresses)
}

// PackEmergencyPauseExtension is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54a77684.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function emergencyPauseExtension(uint256 _extensionId) returns()
func (machineEmergencyPause *MachineEmergencyPause) PackEmergencyPauseExtension(extensionId *big.Int) []byte {
	enc, err := machineEmergencyPause.abi.Pack("emergencyPauseExtension", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEmergencyPauseExtension is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54a77684.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function emergencyPauseExtension(uint256 _extensionId) returns()
func (machineEmergencyPause *MachineEmergencyPause) TryPackEmergencyPauseExtension(extensionId *big.Int) ([]byte, error) {
	return machineEmergencyPause.abi.Pack("emergencyPauseExtension", extensionId)
}

// PackEmergencyUnpauseExtension is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x255fb084.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function emergencyUnpauseExtension(uint256 _extensionId) returns()
func (machineEmergencyPause *MachineEmergencyPause) PackEmergencyUnpauseExtension(extensionId *big.Int) []byte {
	enc, err := machineEmergencyPause.abi.Pack("emergencyUnpauseExtension", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEmergencyUnpauseExtension is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x255fb084.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function emergencyUnpauseExtension(uint256 _extensionId) returns()
func (machineEmergencyPause *MachineEmergencyPause) TryPackEmergencyUnpauseExtension(extensionId *big.Int) ([]byte, error) {
	return machineEmergencyPause.abi.Pack("emergencyUnpauseExtension", extensionId)
}

// PackGetEmergencyUnpauseGracePeriodSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x92f8d8a3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getEmergencyUnpauseGracePeriodSeconds() view returns(uint256 _seconds)
func (machineEmergencyPause *MachineEmergencyPause) PackGetEmergencyUnpauseGracePeriodSeconds() []byte {
	enc, err := machineEmergencyPause.abi.Pack("getEmergencyUnpauseGracePeriodSeconds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetEmergencyUnpauseGracePeriodSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x92f8d8a3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getEmergencyUnpauseGracePeriodSeconds() view returns(uint256 _seconds)
func (machineEmergencyPause *MachineEmergencyPause) TryPackGetEmergencyUnpauseGracePeriodSeconds() ([]byte, error) {
	return machineEmergencyPause.abi.Pack("getEmergencyUnpauseGracePeriodSeconds")
}

// UnpackGetEmergencyUnpauseGracePeriodSeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x92f8d8a3.
//
// Solidity: function getEmergencyUnpauseGracePeriodSeconds() view returns(uint256 _seconds)
func (machineEmergencyPause *MachineEmergencyPause) UnpackGetEmergencyUnpauseGracePeriodSeconds(data []byte) (*big.Int, error) {
	out, err := machineEmergencyPause.abi.Unpack("getEmergencyUnpauseGracePeriodSeconds", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetExtensionEmergencyPausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x930a0f03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getExtensionEmergencyPausers(uint256 _extensionId) view returns(address[] _addresses)
func (machineEmergencyPause *MachineEmergencyPause) PackGetExtensionEmergencyPausers(extensionId *big.Int) []byte {
	enc, err := machineEmergencyPause.abi.Pack("getExtensionEmergencyPausers", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetExtensionEmergencyPausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x930a0f03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getExtensionEmergencyPausers(uint256 _extensionId) view returns(address[] _addresses)
func (machineEmergencyPause *MachineEmergencyPause) TryPackGetExtensionEmergencyPausers(extensionId *big.Int) ([]byte, error) {
	return machineEmergencyPause.abi.Pack("getExtensionEmergencyPausers", extensionId)
}

// UnpackGetExtensionEmergencyPausers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x930a0f03.
//
// Solidity: function getExtensionEmergencyPausers(uint256 _extensionId) view returns(address[] _addresses)
func (machineEmergencyPause *MachineEmergencyPause) UnpackGetExtensionEmergencyPausers(data []byte) ([]common.Address, error) {
	out, err := machineEmergencyPause.abi.Unpack("getExtensionEmergencyPausers", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetExtensionEmergencyUnpausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbeefb796.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getExtensionEmergencyUnpausers(uint256 _extensionId) view returns(address[] _addresses)
func (machineEmergencyPause *MachineEmergencyPause) PackGetExtensionEmergencyUnpausers(extensionId *big.Int) []byte {
	enc, err := machineEmergencyPause.abi.Pack("getExtensionEmergencyUnpausers", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetExtensionEmergencyUnpausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbeefb796.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getExtensionEmergencyUnpausers(uint256 _extensionId) view returns(address[] _addresses)
func (machineEmergencyPause *MachineEmergencyPause) TryPackGetExtensionEmergencyUnpausers(extensionId *big.Int) ([]byte, error) {
	return machineEmergencyPause.abi.Pack("getExtensionEmergencyUnpausers", extensionId)
}

// UnpackGetExtensionEmergencyUnpausers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbeefb796.
//
// Solidity: function getExtensionEmergencyUnpausers(uint256 _extensionId) view returns(address[] _addresses)
func (machineEmergencyPause *MachineEmergencyPause) UnpackGetExtensionEmergencyUnpausers(data []byte) ([]common.Address, error) {
	out, err := machineEmergencyPause.abi.Unpack("getExtensionEmergencyUnpausers", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetLastUnpauseTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87ba2360.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLastUnpauseTs(uint256 _extensionId) view returns(uint64 _ts)
func (machineEmergencyPause *MachineEmergencyPause) PackGetLastUnpauseTs(extensionId *big.Int) []byte {
	enc, err := machineEmergencyPause.abi.Pack("getLastUnpauseTs", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLastUnpauseTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87ba2360.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLastUnpauseTs(uint256 _extensionId) view returns(uint64 _ts)
func (machineEmergencyPause *MachineEmergencyPause) TryPackGetLastUnpauseTs(extensionId *big.Int) ([]byte, error) {
	return machineEmergencyPause.abi.Pack("getLastUnpauseTs", extensionId)
}

// UnpackGetLastUnpauseTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87ba2360.
//
// Solidity: function getLastUnpauseTs(uint256 _extensionId) view returns(uint64 _ts)
func (machineEmergencyPause *MachineEmergencyPause) UnpackGetLastUnpauseTs(data []byte) (uint64, error) {
	out, err := machineEmergencyPause.abi.Unpack("getLastUnpauseTs", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackIsExtensionEmergencyPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd8121be4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExtensionEmergencyPaused(uint256 _extensionId) view returns(bool _paused)
func (machineEmergencyPause *MachineEmergencyPause) PackIsExtensionEmergencyPaused(extensionId *big.Int) []byte {
	enc, err := machineEmergencyPause.abi.Pack("isExtensionEmergencyPaused", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsExtensionEmergencyPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd8121be4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isExtensionEmergencyPaused(uint256 _extensionId) view returns(bool _paused)
func (machineEmergencyPause *MachineEmergencyPause) TryPackIsExtensionEmergencyPaused(extensionId *big.Int) ([]byte, error) {
	return machineEmergencyPause.abi.Pack("isExtensionEmergencyPaused", extensionId)
}

// UnpackIsExtensionEmergencyPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd8121be4.
//
// Solidity: function isExtensionEmergencyPaused(uint256 _extensionId) view returns(bool _paused)
func (machineEmergencyPause *MachineEmergencyPause) UnpackIsExtensionEmergencyPaused(data []byte) (bool, error) {
	out, err := machineEmergencyPause.abi.Unpack("isExtensionEmergencyPaused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsExtensionEmergencyPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x192364a5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExtensionEmergencyPauser(uint256 _extensionId, address _addr) view returns(bool _isPauser)
func (machineEmergencyPause *MachineEmergencyPause) PackIsExtensionEmergencyPauser(extensionId *big.Int, addr common.Address) []byte {
	enc, err := machineEmergencyPause.abi.Pack("isExtensionEmergencyPauser", extensionId, addr)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsExtensionEmergencyPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x192364a5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isExtensionEmergencyPauser(uint256 _extensionId, address _addr) view returns(bool _isPauser)
func (machineEmergencyPause *MachineEmergencyPause) TryPackIsExtensionEmergencyPauser(extensionId *big.Int, addr common.Address) ([]byte, error) {
	return machineEmergencyPause.abi.Pack("isExtensionEmergencyPauser", extensionId, addr)
}

// UnpackIsExtensionEmergencyPauser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x192364a5.
//
// Solidity: function isExtensionEmergencyPauser(uint256 _extensionId, address _addr) view returns(bool _isPauser)
func (machineEmergencyPause *MachineEmergencyPause) UnpackIsExtensionEmergencyPauser(data []byte) (bool, error) {
	out, err := machineEmergencyPause.abi.Unpack("isExtensionEmergencyPauser", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsExtensionEmergencyUnpauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc61b79ff.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExtensionEmergencyUnpauser(uint256 _extensionId, address _addr) view returns(bool _isUnpauser)
func (machineEmergencyPause *MachineEmergencyPause) PackIsExtensionEmergencyUnpauser(extensionId *big.Int, addr common.Address) []byte {
	enc, err := machineEmergencyPause.abi.Pack("isExtensionEmergencyUnpauser", extensionId, addr)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsExtensionEmergencyUnpauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc61b79ff.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isExtensionEmergencyUnpauser(uint256 _extensionId, address _addr) view returns(bool _isUnpauser)
func (machineEmergencyPause *MachineEmergencyPause) TryPackIsExtensionEmergencyUnpauser(extensionId *big.Int, addr common.Address) ([]byte, error) {
	return machineEmergencyPause.abi.Pack("isExtensionEmergencyUnpauser", extensionId, addr)
}

// UnpackIsExtensionEmergencyUnpauser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc61b79ff.
//
// Solidity: function isExtensionEmergencyUnpauser(uint256 _extensionId, address _addr) view returns(bool _isUnpauser)
func (machineEmergencyPause *MachineEmergencyPause) UnpackIsExtensionEmergencyUnpauser(data []byte) (bool, error) {
	out, err := machineEmergencyPause.abi.Unpack("isExtensionEmergencyUnpauser", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRemoveExtensionEmergencyPausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa53840ab.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeExtensionEmergencyPausers(uint256 _extensionId, address[] _addresses) returns()
func (machineEmergencyPause *MachineEmergencyPause) PackRemoveExtensionEmergencyPausers(extensionId *big.Int, addresses []common.Address) []byte {
	enc, err := machineEmergencyPause.abi.Pack("removeExtensionEmergencyPausers", extensionId, addresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveExtensionEmergencyPausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa53840ab.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeExtensionEmergencyPausers(uint256 _extensionId, address[] _addresses) returns()
func (machineEmergencyPause *MachineEmergencyPause) TryPackRemoveExtensionEmergencyPausers(extensionId *big.Int, addresses []common.Address) ([]byte, error) {
	return machineEmergencyPause.abi.Pack("removeExtensionEmergencyPausers", extensionId, addresses)
}

// PackRemoveExtensionEmergencyUnpausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x98b5cfe0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeExtensionEmergencyUnpausers(uint256 _extensionId, address[] _addresses) returns()
func (machineEmergencyPause *MachineEmergencyPause) PackRemoveExtensionEmergencyUnpausers(extensionId *big.Int, addresses []common.Address) []byte {
	enc, err := machineEmergencyPause.abi.Pack("removeExtensionEmergencyUnpausers", extensionId, addresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveExtensionEmergencyUnpausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x98b5cfe0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeExtensionEmergencyUnpausers(uint256 _extensionId, address[] _addresses) returns()
func (machineEmergencyPause *MachineEmergencyPause) TryPackRemoveExtensionEmergencyUnpausers(extensionId *big.Int, addresses []common.Address) ([]byte, error) {
	return machineEmergencyPause.abi.Pack("removeExtensionEmergencyUnpausers", extensionId, addresses)
}

// PackSetEmergencyUnpauseGracePeriodSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8fb045d5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setEmergencyUnpauseGracePeriodSeconds(uint256 _seconds) returns()
func (machineEmergencyPause *MachineEmergencyPause) PackSetEmergencyUnpauseGracePeriodSeconds(seconds *big.Int) []byte {
	enc, err := machineEmergencyPause.abi.Pack("setEmergencyUnpauseGracePeriodSeconds", seconds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetEmergencyUnpauseGracePeriodSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8fb045d5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setEmergencyUnpauseGracePeriodSeconds(uint256 _seconds) returns()
func (machineEmergencyPause *MachineEmergencyPause) TryPackSetEmergencyUnpauseGracePeriodSeconds(seconds *big.Int) ([]byte, error) {
	return machineEmergencyPause.abi.Pack("setEmergencyUnpauseGracePeriodSeconds", seconds)
}

// MachineEmergencyPauseEmergencyUnpauseGracePeriodSet represents a EmergencyUnpauseGracePeriodSet event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseEmergencyUnpauseGracePeriodSet struct {
	GraceSeconds *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const MachineEmergencyPauseEmergencyUnpauseGracePeriodSetEventName = "EmergencyUnpauseGracePeriodSet"

// ContractEventName returns the user-defined event name.
func (MachineEmergencyPauseEmergencyUnpauseGracePeriodSet) ContractEventName() string {
	return MachineEmergencyPauseEmergencyUnpauseGracePeriodSetEventName
}

// UnpackEmergencyUnpauseGracePeriodSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EmergencyUnpauseGracePeriodSet(uint256 graceSeconds)
func (machineEmergencyPause *MachineEmergencyPause) UnpackEmergencyUnpauseGracePeriodSetEvent(log *types.Log) (*MachineEmergencyPauseEmergencyUnpauseGracePeriodSet, error) {
	event := "EmergencyUnpauseGracePeriodSet"
	if len(log.Topics) == 0 || log.Topics[0] != machineEmergencyPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineEmergencyPauseEmergencyUnpauseGracePeriodSet)
	if len(log.Data) > 0 {
		if err := machineEmergencyPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineEmergencyPause.abi.Events[event].Inputs {
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

// MachineEmergencyPauseExtensionEmergencyPaused represents a ExtensionEmergencyPaused event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyPaused struct {
	ExtensionId *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const MachineEmergencyPauseExtensionEmergencyPausedEventName = "ExtensionEmergencyPaused"

// ContractEventName returns the user-defined event name.
func (MachineEmergencyPauseExtensionEmergencyPaused) ContractEventName() string {
	return MachineEmergencyPauseExtensionEmergencyPausedEventName
}

// UnpackExtensionEmergencyPausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ExtensionEmergencyPaused(uint256 indexed extensionId)
func (machineEmergencyPause *MachineEmergencyPause) UnpackExtensionEmergencyPausedEvent(log *types.Log) (*MachineEmergencyPauseExtensionEmergencyPaused, error) {
	event := "ExtensionEmergencyPaused"
	if len(log.Topics) == 0 || log.Topics[0] != machineEmergencyPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineEmergencyPauseExtensionEmergencyPaused)
	if len(log.Data) > 0 {
		if err := machineEmergencyPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineEmergencyPause.abi.Events[event].Inputs {
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

// MachineEmergencyPauseExtensionEmergencyPausersAdded represents a ExtensionEmergencyPausersAdded event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyPausersAdded struct {
	ExtensionId *big.Int
	Addresses   []common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const MachineEmergencyPauseExtensionEmergencyPausersAddedEventName = "ExtensionEmergencyPausersAdded"

// ContractEventName returns the user-defined event name.
func (MachineEmergencyPauseExtensionEmergencyPausersAdded) ContractEventName() string {
	return MachineEmergencyPauseExtensionEmergencyPausersAddedEventName
}

// UnpackExtensionEmergencyPausersAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ExtensionEmergencyPausersAdded(uint256 indexed extensionId, address[] addresses)
func (machineEmergencyPause *MachineEmergencyPause) UnpackExtensionEmergencyPausersAddedEvent(log *types.Log) (*MachineEmergencyPauseExtensionEmergencyPausersAdded, error) {
	event := "ExtensionEmergencyPausersAdded"
	if len(log.Topics) == 0 || log.Topics[0] != machineEmergencyPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineEmergencyPauseExtensionEmergencyPausersAdded)
	if len(log.Data) > 0 {
		if err := machineEmergencyPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineEmergencyPause.abi.Events[event].Inputs {
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

// MachineEmergencyPauseExtensionEmergencyPausersRemoved represents a ExtensionEmergencyPausersRemoved event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyPausersRemoved struct {
	ExtensionId *big.Int
	Addresses   []common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const MachineEmergencyPauseExtensionEmergencyPausersRemovedEventName = "ExtensionEmergencyPausersRemoved"

// ContractEventName returns the user-defined event name.
func (MachineEmergencyPauseExtensionEmergencyPausersRemoved) ContractEventName() string {
	return MachineEmergencyPauseExtensionEmergencyPausersRemovedEventName
}

// UnpackExtensionEmergencyPausersRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ExtensionEmergencyPausersRemoved(uint256 indexed extensionId, address[] addresses)
func (machineEmergencyPause *MachineEmergencyPause) UnpackExtensionEmergencyPausersRemovedEvent(log *types.Log) (*MachineEmergencyPauseExtensionEmergencyPausersRemoved, error) {
	event := "ExtensionEmergencyPausersRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != machineEmergencyPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineEmergencyPauseExtensionEmergencyPausersRemoved)
	if len(log.Data) > 0 {
		if err := machineEmergencyPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineEmergencyPause.abi.Events[event].Inputs {
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

// MachineEmergencyPauseExtensionEmergencyUnpaused represents a ExtensionEmergencyUnpaused event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyUnpaused struct {
	ExtensionId *big.Int
	UnpauseTs   uint64
	Raw         *types.Log // Blockchain specific contextual infos
}

const MachineEmergencyPauseExtensionEmergencyUnpausedEventName = "ExtensionEmergencyUnpaused"

// ContractEventName returns the user-defined event name.
func (MachineEmergencyPauseExtensionEmergencyUnpaused) ContractEventName() string {
	return MachineEmergencyPauseExtensionEmergencyUnpausedEventName
}

// UnpackExtensionEmergencyUnpausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ExtensionEmergencyUnpaused(uint256 indexed extensionId, uint64 unpauseTs)
func (machineEmergencyPause *MachineEmergencyPause) UnpackExtensionEmergencyUnpausedEvent(log *types.Log) (*MachineEmergencyPauseExtensionEmergencyUnpaused, error) {
	event := "ExtensionEmergencyUnpaused"
	if len(log.Topics) == 0 || log.Topics[0] != machineEmergencyPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineEmergencyPauseExtensionEmergencyUnpaused)
	if len(log.Data) > 0 {
		if err := machineEmergencyPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineEmergencyPause.abi.Events[event].Inputs {
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

// MachineEmergencyPauseExtensionEmergencyUnpausersAdded represents a ExtensionEmergencyUnpausersAdded event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyUnpausersAdded struct {
	ExtensionId *big.Int
	Addresses   []common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const MachineEmergencyPauseExtensionEmergencyUnpausersAddedEventName = "ExtensionEmergencyUnpausersAdded"

// ContractEventName returns the user-defined event name.
func (MachineEmergencyPauseExtensionEmergencyUnpausersAdded) ContractEventName() string {
	return MachineEmergencyPauseExtensionEmergencyUnpausersAddedEventName
}

// UnpackExtensionEmergencyUnpausersAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ExtensionEmergencyUnpausersAdded(uint256 indexed extensionId, address[] addresses)
func (machineEmergencyPause *MachineEmergencyPause) UnpackExtensionEmergencyUnpausersAddedEvent(log *types.Log) (*MachineEmergencyPauseExtensionEmergencyUnpausersAdded, error) {
	event := "ExtensionEmergencyUnpausersAdded"
	if len(log.Topics) == 0 || log.Topics[0] != machineEmergencyPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineEmergencyPauseExtensionEmergencyUnpausersAdded)
	if len(log.Data) > 0 {
		if err := machineEmergencyPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineEmergencyPause.abi.Events[event].Inputs {
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

// MachineEmergencyPauseExtensionEmergencyUnpausersRemoved represents a ExtensionEmergencyUnpausersRemoved event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyUnpausersRemoved struct {
	ExtensionId *big.Int
	Addresses   []common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const MachineEmergencyPauseExtensionEmergencyUnpausersRemovedEventName = "ExtensionEmergencyUnpausersRemoved"

// ContractEventName returns the user-defined event name.
func (MachineEmergencyPauseExtensionEmergencyUnpausersRemoved) ContractEventName() string {
	return MachineEmergencyPauseExtensionEmergencyUnpausersRemovedEventName
}

// UnpackExtensionEmergencyUnpausersRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ExtensionEmergencyUnpausersRemoved(uint256 indexed extensionId, address[] addresses)
func (machineEmergencyPause *MachineEmergencyPause) UnpackExtensionEmergencyUnpausersRemovedEvent(log *types.Log) (*MachineEmergencyPauseExtensionEmergencyUnpausersRemoved, error) {
	event := "ExtensionEmergencyUnpausersRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != machineEmergencyPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineEmergencyPauseExtensionEmergencyUnpausersRemoved)
	if len(log.Data) > 0 {
		if err := machineEmergencyPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineEmergencyPause.abi.Events[event].Inputs {
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

// MachineEmergencyPauseGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const MachineEmergencyPauseGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (MachineEmergencyPauseGovernanceCallTimelocked) ContractEventName() string {
	return MachineEmergencyPauseGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (machineEmergencyPause *MachineEmergencyPause) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*MachineEmergencyPauseGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != machineEmergencyPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineEmergencyPauseGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := machineEmergencyPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineEmergencyPause.abi.Events[event].Inputs {
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

// MachineEmergencyPauseInitialized represents a Initialized event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const MachineEmergencyPauseInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (MachineEmergencyPauseInitialized) ContractEventName() string {
	return MachineEmergencyPauseInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (machineEmergencyPause *MachineEmergencyPause) UnpackInitializedEvent(log *types.Log) (*MachineEmergencyPauseInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != machineEmergencyPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineEmergencyPauseInitialized)
	if len(log.Data) > 0 {
		if err := machineEmergencyPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineEmergencyPause.abi.Events[event].Inputs {
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
func (machineEmergencyPause *MachineEmergencyPause) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["AddressAlreadyInSet"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackAddressAlreadyInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["AddressNotInSet"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackAddressNotInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["AvailabilityCheckTimestampInvalid"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackAvailabilityCheckTimestampInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["DuplicatedCosigner"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackDuplicatedCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["EmergencyPauseActive"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackEmergencyPauseActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["EmergencyProtectionActive"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackEmergencyProtectionActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["ExtensionAlreadyEmergencyPaused"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackExtensionAlreadyEmergencyPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["ExtensionNotEmergencyPaused"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackExtensionNotEmergencyPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["GracePeriodTooLong"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackGracePeriodTooLongError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["GracePeriodTooShort"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackGracePeriodTooShortError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["InvalidAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackInvalidAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["InvalidKeyType"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackInvalidKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["InvalidPublicKey"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackInvalidPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["InvalidResponseData"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackInvalidResponseDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["InvalidSigningAlgo"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackInvalidSigningAlgoError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["InvalidWalletStatus"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackInvalidWalletStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["KeyTypeNotSupported"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["NoAddresses"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackNoAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["NotOwnerOrPauser"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackNotOwnerOrPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["NotOwnerOrUnpauser"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackNotOwnerOrUnpauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["OnlyExtensionOwner"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackOnlyExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackOnlyOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["OnlyOwnerOrBackupManager"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackOnlyOwnerOrBackupManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["OnlyProposedOwner"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackOnlyProposedOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["OwnerNotAllowed"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackOwnerNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineEmergencyPause.abi.Errors["VersionNotSupported"].ID.Bytes()[:4]) {
		return machineEmergencyPause.UnpackVersionNotSupportedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// MachineEmergencyPauseAddressAlreadyInSet represents a AddressAlreadyInSet error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseAddressAlreadyInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressAlreadyInSet(address addr)
func MachineEmergencyPauseAddressAlreadyInSetErrorID() common.Hash {
	return common.HexToHash("0xb725c7fb1c3f0f065973d6b9eb55140f970259955af3045742cd5b37902eb920")
}

// UnpackAddressAlreadyInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressAlreadyInSet(address addr)
func (machineEmergencyPause *MachineEmergencyPause) UnpackAddressAlreadyInSetError(raw []byte) (*MachineEmergencyPauseAddressAlreadyInSet, error) {
	out := new(MachineEmergencyPauseAddressAlreadyInSet)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "AddressAlreadyInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseAddressNotInSet represents a AddressNotInSet error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseAddressNotInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressNotInSet(address addr)
func MachineEmergencyPauseAddressNotInSetErrorID() common.Hash {
	return common.HexToHash("0x1ff5ecf3e6d385ea01ff0323fe894ed6d25a83ed5f6ff2a2ede4c897408f6be9")
}

// UnpackAddressNotInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressNotInSet(address addr)
func (machineEmergencyPause *MachineEmergencyPause) UnpackAddressNotInSetError(raw []byte) (*MachineEmergencyPauseAddressNotInSet, error) {
	out := new(MachineEmergencyPauseAddressNotInSet)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "AddressNotInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseAvailabilityCheckTimestampInvalid represents a AvailabilityCheckTimestampInvalid error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseAvailabilityCheckTimestampInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func MachineEmergencyPauseAvailabilityCheckTimestampInvalidErrorID() common.Hash {
	return common.HexToHash("0x6344a5e74cff78612fa271fa907ca608af3bad10e1e8c8df302e29ebc6298e2d")
}

// UnpackAvailabilityCheckTimestampInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func (machineEmergencyPause *MachineEmergencyPause) UnpackAvailabilityCheckTimestampInvalidError(raw []byte) (*MachineEmergencyPauseAvailabilityCheckTimestampInvalid, error) {
	out := new(MachineEmergencyPauseAvailabilityCheckTimestampInvalid)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "AvailabilityCheckTimestampInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseDuplicatedCosigner represents a DuplicatedCosigner error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseDuplicatedCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func MachineEmergencyPauseDuplicatedCosignerErrorID() common.Hash {
	return common.HexToHash("0x53a5890417576fb3043dc3ff1669b0932da2f657d8adb4c11313b12b711f5910")
}

// UnpackDuplicatedCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func (machineEmergencyPause *MachineEmergencyPause) UnpackDuplicatedCosignerError(raw []byte) (*MachineEmergencyPauseDuplicatedCosigner, error) {
	out := new(MachineEmergencyPauseDuplicatedCosigner)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "DuplicatedCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseEmergencyPauseActive represents a EmergencyPauseActive error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseEmergencyPauseActive struct {
	ExtensionId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func MachineEmergencyPauseEmergencyPauseActiveErrorID() common.Hash {
	return common.HexToHash("0x480cf2163927c55f46b05a5cac80ac3781076b0e50b1ed6b07995a800937d4b6")
}

// UnpackEmergencyPauseActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func (machineEmergencyPause *MachineEmergencyPause) UnpackEmergencyPauseActiveError(raw []byte) (*MachineEmergencyPauseEmergencyPauseActive, error) {
	out := new(MachineEmergencyPauseEmergencyPauseActive)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "EmergencyPauseActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseEmergencyProtectionActive represents a EmergencyProtectionActive error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseEmergencyProtectionActive struct {
	ExtensionId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmergencyProtectionActive(uint256 extensionId)
func MachineEmergencyPauseEmergencyProtectionActiveErrorID() common.Hash {
	return common.HexToHash("0x93136b5bc178e37d8a37403ffe4d8ca41116b64440923a070fa269ad41a394bf")
}

// UnpackEmergencyProtectionActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmergencyProtectionActive(uint256 extensionId)
func (machineEmergencyPause *MachineEmergencyPause) UnpackEmergencyProtectionActiveError(raw []byte) (*MachineEmergencyPauseEmergencyProtectionActive, error) {
	out := new(MachineEmergencyPauseEmergencyProtectionActive)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "EmergencyProtectionActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseExtensionAlreadyEmergencyPaused represents a ExtensionAlreadyEmergencyPaused error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionAlreadyEmergencyPaused struct {
	ExtensionId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionAlreadyEmergencyPaused(uint256 extensionId)
func MachineEmergencyPauseExtensionAlreadyEmergencyPausedErrorID() common.Hash {
	return common.HexToHash("0xfed3231a21a06ac8b977e13a75d9bcae39195a0e0dec258dd9e2e050753301b6")
}

// UnpackExtensionAlreadyEmergencyPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionAlreadyEmergencyPaused(uint256 extensionId)
func (machineEmergencyPause *MachineEmergencyPause) UnpackExtensionAlreadyEmergencyPausedError(raw []byte) (*MachineEmergencyPauseExtensionAlreadyEmergencyPaused, error) {
	out := new(MachineEmergencyPauseExtensionAlreadyEmergencyPaused)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "ExtensionAlreadyEmergencyPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseExtensionIdMismatch represents a ExtensionIdMismatch error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func MachineEmergencyPauseExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (machineEmergencyPause *MachineEmergencyPause) UnpackExtensionIdMismatchError(raw []byte) (*MachineEmergencyPauseExtensionIdMismatch, error) {
	out := new(MachineEmergencyPauseExtensionIdMismatch)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseExtensionNotEmergencyPaused represents a ExtensionNotEmergencyPaused error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionNotEmergencyPaused struct {
	ExtensionId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionNotEmergencyPaused(uint256 extensionId)
func MachineEmergencyPauseExtensionNotEmergencyPausedErrorID() common.Hash {
	return common.HexToHash("0x8e730727309ffb777fdfc02f8794c29f01c687ccc809862e6bf128b311870d8c")
}

// UnpackExtensionNotEmergencyPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionNotEmergencyPaused(uint256 extensionId)
func (machineEmergencyPause *MachineEmergencyPause) UnpackExtensionNotEmergencyPausedError(raw []byte) (*MachineEmergencyPauseExtensionNotEmergencyPaused, error) {
	out := new(MachineEmergencyPauseExtensionNotEmergencyPaused)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "ExtensionNotEmergencyPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseGracePeriodTooLong represents a GracePeriodTooLong error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseGracePeriodTooLong struct {
	GraceSeconds *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GracePeriodTooLong(uint256 graceSeconds)
func MachineEmergencyPauseGracePeriodTooLongErrorID() common.Hash {
	return common.HexToHash("0x8ff948d6119974fa3ecbb383b14b7f8685cce2762011e98acf71c1a9158cbad5")
}

// UnpackGracePeriodTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GracePeriodTooLong(uint256 graceSeconds)
func (machineEmergencyPause *MachineEmergencyPause) UnpackGracePeriodTooLongError(raw []byte) (*MachineEmergencyPauseGracePeriodTooLong, error) {
	out := new(MachineEmergencyPauseGracePeriodTooLong)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "GracePeriodTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseGracePeriodTooShort represents a GracePeriodTooShort error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseGracePeriodTooShort struct {
	GraceSeconds *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GracePeriodTooShort(uint256 graceSeconds)
func MachineEmergencyPauseGracePeriodTooShortErrorID() common.Hash {
	return common.HexToHash("0xa77e74e78e7adccece83e0a693bd733e0f6da702830477df6f7d58f8b3927e2e")
}

// UnpackGracePeriodTooShortError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GracePeriodTooShort(uint256 graceSeconds)
func (machineEmergencyPause *MachineEmergencyPause) UnpackGracePeriodTooShortError(raw []byte) (*MachineEmergencyPauseGracePeriodTooShort, error) {
	out := new(MachineEmergencyPauseGracePeriodTooShort)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "GracePeriodTooShort", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseInvalidAddress represents a InvalidAddress error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func MachineEmergencyPauseInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (machineEmergencyPause *MachineEmergencyPause) UnpackInvalidAddressError(raw []byte) (*MachineEmergencyPauseInvalidAddress, error) {
	out := new(MachineEmergencyPauseInvalidAddress)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseInvalidAvailabilityCheckStatus represents a InvalidAvailabilityCheckStatus error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInvalidAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func MachineEmergencyPauseInvalidAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0xedee4ff00b842d3ef84031c60001065c540c838c5479bbf0ae56f15868427e44")
}

// UnpackInvalidAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func (machineEmergencyPause *MachineEmergencyPause) UnpackInvalidAvailabilityCheckStatusError(raw []byte) (*MachineEmergencyPauseInvalidAvailabilityCheckStatus, error) {
	out := new(MachineEmergencyPauseInvalidAvailabilityCheckStatus)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "InvalidAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseInvalidCosigner represents a InvalidCosigner error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func MachineEmergencyPauseInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (machineEmergencyPause *MachineEmergencyPause) UnpackInvalidCosignerError(raw []byte) (*MachineEmergencyPauseInvalidCosigner, error) {
	out := new(MachineEmergencyPauseInvalidCosigner)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseInvalidDuration represents a InvalidDuration error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func MachineEmergencyPauseInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (machineEmergencyPause *MachineEmergencyPause) UnpackInvalidDurationError(raw []byte) (*MachineEmergencyPauseInvalidDuration, error) {
	out := new(MachineEmergencyPauseInvalidDuration)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func MachineEmergencyPauseInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (machineEmergencyPause *MachineEmergencyPause) UnpackInvalidGovernanceHashError(raw []byte) (*MachineEmergencyPauseInvalidGovernanceHash, error) {
	out := new(MachineEmergencyPauseInvalidGovernanceHash)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseInvalidInitialization represents a InvalidInitialization error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func MachineEmergencyPauseInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (machineEmergencyPause *MachineEmergencyPause) UnpackInvalidInitializationError(raw []byte) (*MachineEmergencyPauseInvalidInitialization, error) {
	out := new(MachineEmergencyPauseInvalidInitialization)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseInvalidKeyType represents a InvalidKeyType error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInvalidKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyType()
func MachineEmergencyPauseInvalidKeyTypeErrorID() common.Hash {
	return common.HexToHash("0x36dccb440beb343622d8ee0927bce3f69281feea3a5b0610f6e00e2811b557e1")
}

// UnpackInvalidKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyType()
func (machineEmergencyPause *MachineEmergencyPause) UnpackInvalidKeyTypeError(raw []byte) (*MachineEmergencyPauseInvalidKeyType, error) {
	out := new(MachineEmergencyPauseInvalidKeyType)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "InvalidKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseInvalidNonce represents a InvalidNonce error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func MachineEmergencyPauseInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (machineEmergencyPause *MachineEmergencyPause) UnpackInvalidNonceError(raw []byte) (*MachineEmergencyPauseInvalidNonce, error) {
	out := new(MachineEmergencyPauseInvalidNonce)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseInvalidPublicKey represents a InvalidPublicKey error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInvalidPublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPublicKey()
func MachineEmergencyPauseInvalidPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa2d0fee854985cb3d3a394146b5df1c41d5b4d6fcede7fc57a3c05a49b2f3e73")
}

// UnpackInvalidPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPublicKey()
func (machineEmergencyPause *MachineEmergencyPause) UnpackInvalidPublicKeyError(raw []byte) (*MachineEmergencyPauseInvalidPublicKey, error) {
	out := new(MachineEmergencyPauseInvalidPublicKey)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "InvalidPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseInvalidResponseData represents a InvalidResponseData error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInvalidResponseData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseData()
func MachineEmergencyPauseInvalidResponseDataErrorID() common.Hash {
	return common.HexToHash("0x48566617128355a033880805d6462da9bf9c6fe4c61b922997451a71262990af")
}

// UnpackInvalidResponseDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseData()
func (machineEmergencyPause *MachineEmergencyPause) UnpackInvalidResponseDataError(raw []byte) (*MachineEmergencyPauseInvalidResponseData, error) {
	out := new(MachineEmergencyPauseInvalidResponseData)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "InvalidResponseData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseInvalidSigningAlgo represents a InvalidSigningAlgo error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInvalidSigningAlgo struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningAlgo()
func MachineEmergencyPauseInvalidSigningAlgoErrorID() common.Hash {
	return common.HexToHash("0x9b7198a697b61e546ef9ee593a6f77d2bf9f1df9dcbba655ce96d7e2c75346c7")
}

// UnpackInvalidSigningAlgoError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningAlgo()
func (machineEmergencyPause *MachineEmergencyPause) UnpackInvalidSigningAlgoError(raw []byte) (*MachineEmergencyPauseInvalidSigningAlgo, error) {
	out := new(MachineEmergencyPauseInvalidSigningAlgo)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "InvalidSigningAlgo", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseInvalidThreshold represents a InvalidThreshold error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func MachineEmergencyPauseInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (machineEmergencyPause *MachineEmergencyPause) UnpackInvalidThresholdError(raw []byte) (*MachineEmergencyPauseInvalidThreshold, error) {
	out := new(MachineEmergencyPauseInvalidThreshold)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseInvalidWalletStatus represents a InvalidWalletStatus error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInvalidWalletStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidWalletStatus()
func MachineEmergencyPauseInvalidWalletStatusErrorID() common.Hash {
	return common.HexToHash("0xd700120e4896ea7835a5e9e4feb690644ced5cad1c52209afdcbe06b96d91345")
}

// UnpackInvalidWalletStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidWalletStatus()
func (machineEmergencyPause *MachineEmergencyPause) UnpackInvalidWalletStatusError(raw []byte) (*MachineEmergencyPauseInvalidWalletStatus, error) {
	out := new(MachineEmergencyPauseInvalidWalletStatus)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "InvalidWalletStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseKeyTypeNotSupported represents a KeyTypeNotSupported error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseKeyTypeNotSupported struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func MachineEmergencyPauseKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x85004df1f47814c2de6fbf03d3f2b1ca8c67c182d0d54278766543b6be3bc603")
}

// UnpackKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func (machineEmergencyPause *MachineEmergencyPause) UnpackKeyTypeNotSupportedError(raw []byte) (*MachineEmergencyPauseKeyTypeNotSupported, error) {
	out := new(MachineEmergencyPauseKeyTypeNotSupported)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "KeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseLengthsMismatch represents a LengthsMismatch error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func MachineEmergencyPauseLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (machineEmergencyPause *MachineEmergencyPause) UnpackLengthsMismatchError(raw []byte) (*MachineEmergencyPauseLengthsMismatch, error) {
	out := new(MachineEmergencyPauseLengthsMismatch)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseNoAddresses represents a NoAddresses error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseNoAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoAddresses()
func MachineEmergencyPauseNoAddressesErrorID() common.Hash {
	return common.HexToHash("0xb2e79f05746fa6c9f3d61f9efe85959371cde5030a6c2dd70c92b63578c5cde7")
}

// UnpackNoAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoAddresses()
func (machineEmergencyPause *MachineEmergencyPause) UnpackNoAddressesError(raw []byte) (*MachineEmergencyPauseNoAddresses, error) {
	out := new(MachineEmergencyPauseNoAddresses)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "NoAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseNotInitializing represents a NotInitializing error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func MachineEmergencyPauseNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (machineEmergencyPause *MachineEmergencyPause) UnpackNotInitializingError(raw []byte) (*MachineEmergencyPauseNotInitializing, error) {
	out := new(MachineEmergencyPauseNotInitializing)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseNotOwnerOrPauser represents a NotOwnerOrPauser error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseNotOwnerOrPauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrPauser(address caller)
func MachineEmergencyPauseNotOwnerOrPauserErrorID() common.Hash {
	return common.HexToHash("0x46c0d9af6b895093e022ad2984c54e60192e6e26c5063daf232e825c6e07ff9a")
}

// UnpackNotOwnerOrPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrPauser(address caller)
func (machineEmergencyPause *MachineEmergencyPause) UnpackNotOwnerOrPauserError(raw []byte) (*MachineEmergencyPauseNotOwnerOrPauser, error) {
	out := new(MachineEmergencyPauseNotOwnerOrPauser)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "NotOwnerOrPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseNotOwnerOrUnpauser represents a NotOwnerOrUnpauser error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseNotOwnerOrUnpauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func MachineEmergencyPauseNotOwnerOrUnpauserErrorID() common.Hash {
	return common.HexToHash("0x81e700e8f0bd5ba4a6d6a2ce9a1ad0ac9bb1b1b1e62d921a55cbc4b53f1da649")
}

// UnpackNotOwnerOrUnpauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func (machineEmergencyPause *MachineEmergencyPause) UnpackNotOwnerOrUnpauserError(raw []byte) (*MachineEmergencyPauseNotOwnerOrUnpauser, error) {
	out := new(MachineEmergencyPauseNotOwnerOrUnpauser)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "NotOwnerOrUnpauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseOnlyExtensionOwner represents a OnlyExtensionOwner error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseOnlyExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwner()
func MachineEmergencyPauseOnlyExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xcd49fd1de5a297cd1ce055d5abffcf2cb25ad9492bd39a703a238af57b2a9488")
}

// UnpackOnlyExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwner()
func (machineEmergencyPause *MachineEmergencyPause) UnpackOnlyExtensionOwnerError(raw []byte) (*MachineEmergencyPauseOnlyExtensionOwner, error) {
	out := new(MachineEmergencyPauseOnlyExtensionOwner)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "OnlyExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func MachineEmergencyPauseOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (machineEmergencyPause *MachineEmergencyPause) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*MachineEmergencyPauseOnlyExtensionOwnerOrOperator, error) {
	out := new(MachineEmergencyPauseOnlyExtensionOwnerOrOperator)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseOnlyGovernance represents a OnlyGovernance error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func MachineEmergencyPauseOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (machineEmergencyPause *MachineEmergencyPause) UnpackOnlyGovernanceError(raw []byte) (*MachineEmergencyPauseOnlyGovernance, error) {
	out := new(MachineEmergencyPauseOnlyGovernance)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseOnlyOwner represents a OnlyOwner error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func MachineEmergencyPauseOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (machineEmergencyPause *MachineEmergencyPause) UnpackOnlyOwnerError(raw []byte) (*MachineEmergencyPauseOnlyOwner, error) {
	out := new(MachineEmergencyPauseOnlyOwner)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseOnlyOwnerOrBackupManager represents a OnlyOwnerOrBackupManager error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseOnlyOwnerOrBackupManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrBackupManager()
func MachineEmergencyPauseOnlyOwnerOrBackupManagerErrorID() common.Hash {
	return common.HexToHash("0x249a6cf6e22092f9c80798e2a0c583c0f36aea0779ef38a5b5cdb6a112a0e6df")
}

// UnpackOnlyOwnerOrBackupManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrBackupManager()
func (machineEmergencyPause *MachineEmergencyPause) UnpackOnlyOwnerOrBackupManagerError(raw []byte) (*MachineEmergencyPauseOnlyOwnerOrBackupManager, error) {
	out := new(MachineEmergencyPauseOnlyOwnerOrBackupManager)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "OnlyOwnerOrBackupManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func MachineEmergencyPauseOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (machineEmergencyPause *MachineEmergencyPause) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*MachineEmergencyPauseOnlyProductionOrPausedStatus, error) {
	out := new(MachineEmergencyPauseOnlyProductionOrPausedStatus)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseOnlyProposedOwner represents a OnlyProposedOwner error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseOnlyProposedOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProposedOwner()
func MachineEmergencyPauseOnlyProposedOwnerErrorID() common.Hash {
	return common.HexToHash("0x49cb15061ee558e74f66f937f740fd2a3791a89d2d139862222e5c97cabc02cd")
}

// UnpackOnlyProposedOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProposedOwner()
func (machineEmergencyPause *MachineEmergencyPause) UnpackOnlyProposedOwnerError(raw []byte) (*MachineEmergencyPauseOnlyProposedOwner, error) {
	out := new(MachineEmergencyPauseOnlyProposedOwner)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "OnlyProposedOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseOwnerNotAllowed represents a OwnerNotAllowed error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseOwnerNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerNotAllowed()
func MachineEmergencyPauseOwnerNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x3d5c9996545c6bf4287c26c1c2d8e6b4c39b92f19ba706c3a756720bf9c4fd24")
}

// UnpackOwnerNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerNotAllowed()
func (machineEmergencyPause *MachineEmergencyPause) UnpackOwnerNotAllowedError(raw []byte) (*MachineEmergencyPauseOwnerNotAllowed, error) {
	out := new(MachineEmergencyPauseOwnerNotAllowed)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "OwnerNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func MachineEmergencyPauseTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (machineEmergencyPause *MachineEmergencyPause) UnpackTeeMachineNotAvailableError(raw []byte) (*MachineEmergencyPauseTeeMachineNotAvailable, error) {
	out := new(MachineEmergencyPauseTeeMachineNotAvailable)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineEmergencyPauseVersionNotSupported represents a VersionNotSupported error raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseVersionNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionNotSupported()
func MachineEmergencyPauseVersionNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x6db0b17b06266b37e721024a8ba1221439799ebb779dfd397865e37250abc9a8")
}

// UnpackVersionNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionNotSupported()
func (machineEmergencyPause *MachineEmergencyPause) UnpackVersionNotSupportedError(raw []byte) (*MachineEmergencyPauseVersionNotSupported, error) {
	out := new(MachineEmergencyPauseVersionNotSupported)
	if err := machineEmergencyPause.abi.UnpackIntoInterface(out, "VersionNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}
