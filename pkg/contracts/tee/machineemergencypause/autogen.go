// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package machineemergencypause

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MachineEmergencyPauseMetaData contains all meta data concerning the MachineEmergencyPause contract.
var MachineEmergencyPauseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"EmergencyPauseActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"EmergencyProtectionActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"ExtensionAlreadyEmergencyPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"ExtensionNotEmergencyPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"graceSeconds\",\"type\":\"uint256\"}],\"name\":\"GracePeriodTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"graceSeconds\",\"type\":\"uint256\"}],\"name\":\"GracePeriodTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"graceSeconds\",\"type\":\"uint256\"}],\"name\":\"EmergencyUnpauseGracePeriodSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"ExtensionEmergencyPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"ExtensionEmergencyPausersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"ExtensionEmergencyPausersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"unpauseTs\",\"type\":\"uint64\"}],\"name\":\"ExtensionEmergencyUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"ExtensionEmergencyUnpausersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"ExtensionEmergencyUnpausersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addExtensionEmergencyPausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addExtensionEmergencyUnpausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"emergencyPauseExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"emergencyUnpauseExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEmergencyUnpauseGracePeriodSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_seconds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getExtensionEmergencyPausers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getExtensionEmergencyUnpausers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getLastUnpauseTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_ts\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"isExtensionEmergencyPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isExtensionEmergencyPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isPauser\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isExtensionEmergencyUnpauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isUnpauser\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"removeExtensionEmergencyPausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"removeExtensionEmergencyUnpausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seconds\",\"type\":\"uint256\"}],\"name\":\"setEmergencyUnpauseGracePeriodSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MachineEmergencyPauseABI is the input ABI used to generate the binding from.
// Deprecated: Use MachineEmergencyPauseMetaData.ABI instead.
var MachineEmergencyPauseABI = MachineEmergencyPauseMetaData.ABI

// MachineEmergencyPause is an auto generated Go binding around an Ethereum contract.
type MachineEmergencyPause struct {
	MachineEmergencyPauseCaller     // Read-only binding to the contract
	MachineEmergencyPauseTransactor // Write-only binding to the contract
	MachineEmergencyPauseFilterer   // Log filterer for contract events
}

// MachineEmergencyPauseCaller is an auto generated read-only Go binding around an Ethereum contract.
type MachineEmergencyPauseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineEmergencyPauseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MachineEmergencyPauseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineEmergencyPauseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MachineEmergencyPauseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineEmergencyPauseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MachineEmergencyPauseSession struct {
	Contract     *MachineEmergencyPause // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MachineEmergencyPauseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MachineEmergencyPauseCallerSession struct {
	Contract *MachineEmergencyPauseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// MachineEmergencyPauseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MachineEmergencyPauseTransactorSession struct {
	Contract     *MachineEmergencyPauseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// MachineEmergencyPauseRaw is an auto generated low-level Go binding around an Ethereum contract.
type MachineEmergencyPauseRaw struct {
	Contract *MachineEmergencyPause // Generic contract binding to access the raw methods on
}

// MachineEmergencyPauseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MachineEmergencyPauseCallerRaw struct {
	Contract *MachineEmergencyPauseCaller // Generic read-only contract binding to access the raw methods on
}

// MachineEmergencyPauseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MachineEmergencyPauseTransactorRaw struct {
	Contract *MachineEmergencyPauseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMachineEmergencyPause creates a new instance of MachineEmergencyPause, bound to a specific deployed contract.
func NewMachineEmergencyPause(address common.Address, backend bind.ContractBackend) (*MachineEmergencyPause, error) {
	contract, err := bindMachineEmergencyPause(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MachineEmergencyPause{MachineEmergencyPauseCaller: MachineEmergencyPauseCaller{contract: contract}, MachineEmergencyPauseTransactor: MachineEmergencyPauseTransactor{contract: contract}, MachineEmergencyPauseFilterer: MachineEmergencyPauseFilterer{contract: contract}}, nil
}

// NewMachineEmergencyPauseCaller creates a new read-only instance of MachineEmergencyPause, bound to a specific deployed contract.
func NewMachineEmergencyPauseCaller(address common.Address, caller bind.ContractCaller) (*MachineEmergencyPauseCaller, error) {
	contract, err := bindMachineEmergencyPause(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MachineEmergencyPauseCaller{contract: contract}, nil
}

// NewMachineEmergencyPauseTransactor creates a new write-only instance of MachineEmergencyPause, bound to a specific deployed contract.
func NewMachineEmergencyPauseTransactor(address common.Address, transactor bind.ContractTransactor) (*MachineEmergencyPauseTransactor, error) {
	contract, err := bindMachineEmergencyPause(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MachineEmergencyPauseTransactor{contract: contract}, nil
}

// NewMachineEmergencyPauseFilterer creates a new log filterer instance of MachineEmergencyPause, bound to a specific deployed contract.
func NewMachineEmergencyPauseFilterer(address common.Address, filterer bind.ContractFilterer) (*MachineEmergencyPauseFilterer, error) {
	contract, err := bindMachineEmergencyPause(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MachineEmergencyPauseFilterer{contract: contract}, nil
}

// bindMachineEmergencyPause binds a generic wrapper to an already deployed contract.
func bindMachineEmergencyPause(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MachineEmergencyPauseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MachineEmergencyPause *MachineEmergencyPauseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MachineEmergencyPause.Contract.MachineEmergencyPauseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MachineEmergencyPause *MachineEmergencyPauseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.MachineEmergencyPauseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MachineEmergencyPause *MachineEmergencyPauseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.MachineEmergencyPauseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MachineEmergencyPause *MachineEmergencyPauseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MachineEmergencyPause.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MachineEmergencyPause *MachineEmergencyPauseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MachineEmergencyPause *MachineEmergencyPauseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.contract.Transact(opts, method, params...)
}

// GetEmergencyUnpauseGracePeriodSeconds is a free data retrieval call binding the contract method 0x92f8d8a3.
//
// Solidity: function getEmergencyUnpauseGracePeriodSeconds() view returns(uint256 _seconds)
func (_MachineEmergencyPause *MachineEmergencyPauseCaller) GetEmergencyUnpauseGracePeriodSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MachineEmergencyPause.contract.Call(opts, &out, "getEmergencyUnpauseGracePeriodSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEmergencyUnpauseGracePeriodSeconds is a free data retrieval call binding the contract method 0x92f8d8a3.
//
// Solidity: function getEmergencyUnpauseGracePeriodSeconds() view returns(uint256 _seconds)
func (_MachineEmergencyPause *MachineEmergencyPauseSession) GetEmergencyUnpauseGracePeriodSeconds() (*big.Int, error) {
	return _MachineEmergencyPause.Contract.GetEmergencyUnpauseGracePeriodSeconds(&_MachineEmergencyPause.CallOpts)
}

// GetEmergencyUnpauseGracePeriodSeconds is a free data retrieval call binding the contract method 0x92f8d8a3.
//
// Solidity: function getEmergencyUnpauseGracePeriodSeconds() view returns(uint256 _seconds)
func (_MachineEmergencyPause *MachineEmergencyPauseCallerSession) GetEmergencyUnpauseGracePeriodSeconds() (*big.Int, error) {
	return _MachineEmergencyPause.Contract.GetEmergencyUnpauseGracePeriodSeconds(&_MachineEmergencyPause.CallOpts)
}

// GetExtensionEmergencyPausers is a free data retrieval call binding the contract method 0x930a0f03.
//
// Solidity: function getExtensionEmergencyPausers(uint256 _extensionId) view returns(address[] _addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseCaller) GetExtensionEmergencyPausers(opts *bind.CallOpts, _extensionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _MachineEmergencyPause.contract.Call(opts, &out, "getExtensionEmergencyPausers", _extensionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetExtensionEmergencyPausers is a free data retrieval call binding the contract method 0x930a0f03.
//
// Solidity: function getExtensionEmergencyPausers(uint256 _extensionId) view returns(address[] _addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseSession) GetExtensionEmergencyPausers(_extensionId *big.Int) ([]common.Address, error) {
	return _MachineEmergencyPause.Contract.GetExtensionEmergencyPausers(&_MachineEmergencyPause.CallOpts, _extensionId)
}

// GetExtensionEmergencyPausers is a free data retrieval call binding the contract method 0x930a0f03.
//
// Solidity: function getExtensionEmergencyPausers(uint256 _extensionId) view returns(address[] _addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseCallerSession) GetExtensionEmergencyPausers(_extensionId *big.Int) ([]common.Address, error) {
	return _MachineEmergencyPause.Contract.GetExtensionEmergencyPausers(&_MachineEmergencyPause.CallOpts, _extensionId)
}

// GetExtensionEmergencyUnpausers is a free data retrieval call binding the contract method 0xbeefb796.
//
// Solidity: function getExtensionEmergencyUnpausers(uint256 _extensionId) view returns(address[] _addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseCaller) GetExtensionEmergencyUnpausers(opts *bind.CallOpts, _extensionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _MachineEmergencyPause.contract.Call(opts, &out, "getExtensionEmergencyUnpausers", _extensionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetExtensionEmergencyUnpausers is a free data retrieval call binding the contract method 0xbeefb796.
//
// Solidity: function getExtensionEmergencyUnpausers(uint256 _extensionId) view returns(address[] _addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseSession) GetExtensionEmergencyUnpausers(_extensionId *big.Int) ([]common.Address, error) {
	return _MachineEmergencyPause.Contract.GetExtensionEmergencyUnpausers(&_MachineEmergencyPause.CallOpts, _extensionId)
}

// GetExtensionEmergencyUnpausers is a free data retrieval call binding the contract method 0xbeefb796.
//
// Solidity: function getExtensionEmergencyUnpausers(uint256 _extensionId) view returns(address[] _addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseCallerSession) GetExtensionEmergencyUnpausers(_extensionId *big.Int) ([]common.Address, error) {
	return _MachineEmergencyPause.Contract.GetExtensionEmergencyUnpausers(&_MachineEmergencyPause.CallOpts, _extensionId)
}

// GetLastUnpauseTs is a free data retrieval call binding the contract method 0x87ba2360.
//
// Solidity: function getLastUnpauseTs(uint256 _extensionId) view returns(uint64 _ts)
func (_MachineEmergencyPause *MachineEmergencyPauseCaller) GetLastUnpauseTs(opts *bind.CallOpts, _extensionId *big.Int) (uint64, error) {
	var out []interface{}
	err := _MachineEmergencyPause.contract.Call(opts, &out, "getLastUnpauseTs", _extensionId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLastUnpauseTs is a free data retrieval call binding the contract method 0x87ba2360.
//
// Solidity: function getLastUnpauseTs(uint256 _extensionId) view returns(uint64 _ts)
func (_MachineEmergencyPause *MachineEmergencyPauseSession) GetLastUnpauseTs(_extensionId *big.Int) (uint64, error) {
	return _MachineEmergencyPause.Contract.GetLastUnpauseTs(&_MachineEmergencyPause.CallOpts, _extensionId)
}

// GetLastUnpauseTs is a free data retrieval call binding the contract method 0x87ba2360.
//
// Solidity: function getLastUnpauseTs(uint256 _extensionId) view returns(uint64 _ts)
func (_MachineEmergencyPause *MachineEmergencyPauseCallerSession) GetLastUnpauseTs(_extensionId *big.Int) (uint64, error) {
	return _MachineEmergencyPause.Contract.GetLastUnpauseTs(&_MachineEmergencyPause.CallOpts, _extensionId)
}

// IsExtensionEmergencyPaused is a free data retrieval call binding the contract method 0xd8121be4.
//
// Solidity: function isExtensionEmergencyPaused(uint256 _extensionId) view returns(bool _paused)
func (_MachineEmergencyPause *MachineEmergencyPauseCaller) IsExtensionEmergencyPaused(opts *bind.CallOpts, _extensionId *big.Int) (bool, error) {
	var out []interface{}
	err := _MachineEmergencyPause.contract.Call(opts, &out, "isExtensionEmergencyPaused", _extensionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExtensionEmergencyPaused is a free data retrieval call binding the contract method 0xd8121be4.
//
// Solidity: function isExtensionEmergencyPaused(uint256 _extensionId) view returns(bool _paused)
func (_MachineEmergencyPause *MachineEmergencyPauseSession) IsExtensionEmergencyPaused(_extensionId *big.Int) (bool, error) {
	return _MachineEmergencyPause.Contract.IsExtensionEmergencyPaused(&_MachineEmergencyPause.CallOpts, _extensionId)
}

// IsExtensionEmergencyPaused is a free data retrieval call binding the contract method 0xd8121be4.
//
// Solidity: function isExtensionEmergencyPaused(uint256 _extensionId) view returns(bool _paused)
func (_MachineEmergencyPause *MachineEmergencyPauseCallerSession) IsExtensionEmergencyPaused(_extensionId *big.Int) (bool, error) {
	return _MachineEmergencyPause.Contract.IsExtensionEmergencyPaused(&_MachineEmergencyPause.CallOpts, _extensionId)
}

// IsExtensionEmergencyPauser is a free data retrieval call binding the contract method 0x192364a5.
//
// Solidity: function isExtensionEmergencyPauser(uint256 _extensionId, address _addr) view returns(bool _isPauser)
func (_MachineEmergencyPause *MachineEmergencyPauseCaller) IsExtensionEmergencyPauser(opts *bind.CallOpts, _extensionId *big.Int, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _MachineEmergencyPause.contract.Call(opts, &out, "isExtensionEmergencyPauser", _extensionId, _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExtensionEmergencyPauser is a free data retrieval call binding the contract method 0x192364a5.
//
// Solidity: function isExtensionEmergencyPauser(uint256 _extensionId, address _addr) view returns(bool _isPauser)
func (_MachineEmergencyPause *MachineEmergencyPauseSession) IsExtensionEmergencyPauser(_extensionId *big.Int, _addr common.Address) (bool, error) {
	return _MachineEmergencyPause.Contract.IsExtensionEmergencyPauser(&_MachineEmergencyPause.CallOpts, _extensionId, _addr)
}

// IsExtensionEmergencyPauser is a free data retrieval call binding the contract method 0x192364a5.
//
// Solidity: function isExtensionEmergencyPauser(uint256 _extensionId, address _addr) view returns(bool _isPauser)
func (_MachineEmergencyPause *MachineEmergencyPauseCallerSession) IsExtensionEmergencyPauser(_extensionId *big.Int, _addr common.Address) (bool, error) {
	return _MachineEmergencyPause.Contract.IsExtensionEmergencyPauser(&_MachineEmergencyPause.CallOpts, _extensionId, _addr)
}

// IsExtensionEmergencyUnpauser is a free data retrieval call binding the contract method 0xc61b79ff.
//
// Solidity: function isExtensionEmergencyUnpauser(uint256 _extensionId, address _addr) view returns(bool _isUnpauser)
func (_MachineEmergencyPause *MachineEmergencyPauseCaller) IsExtensionEmergencyUnpauser(opts *bind.CallOpts, _extensionId *big.Int, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _MachineEmergencyPause.contract.Call(opts, &out, "isExtensionEmergencyUnpauser", _extensionId, _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExtensionEmergencyUnpauser is a free data retrieval call binding the contract method 0xc61b79ff.
//
// Solidity: function isExtensionEmergencyUnpauser(uint256 _extensionId, address _addr) view returns(bool _isUnpauser)
func (_MachineEmergencyPause *MachineEmergencyPauseSession) IsExtensionEmergencyUnpauser(_extensionId *big.Int, _addr common.Address) (bool, error) {
	return _MachineEmergencyPause.Contract.IsExtensionEmergencyUnpauser(&_MachineEmergencyPause.CallOpts, _extensionId, _addr)
}

// IsExtensionEmergencyUnpauser is a free data retrieval call binding the contract method 0xc61b79ff.
//
// Solidity: function isExtensionEmergencyUnpauser(uint256 _extensionId, address _addr) view returns(bool _isUnpauser)
func (_MachineEmergencyPause *MachineEmergencyPauseCallerSession) IsExtensionEmergencyUnpauser(_extensionId *big.Int, _addr common.Address) (bool, error) {
	return _MachineEmergencyPause.Contract.IsExtensionEmergencyUnpauser(&_MachineEmergencyPause.CallOpts, _extensionId, _addr)
}

// AddExtensionEmergencyPausers is a paid mutator transaction binding the contract method 0x6d02fabc.
//
// Solidity: function addExtensionEmergencyPausers(uint256 _extensionId, address[] _addresses) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactor) AddExtensionEmergencyPausers(opts *bind.TransactOpts, _extensionId *big.Int, _addresses []common.Address) (*types.Transaction, error) {
	return _MachineEmergencyPause.contract.Transact(opts, "addExtensionEmergencyPausers", _extensionId, _addresses)
}

// AddExtensionEmergencyPausers is a paid mutator transaction binding the contract method 0x6d02fabc.
//
// Solidity: function addExtensionEmergencyPausers(uint256 _extensionId, address[] _addresses) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseSession) AddExtensionEmergencyPausers(_extensionId *big.Int, _addresses []common.Address) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.AddExtensionEmergencyPausers(&_MachineEmergencyPause.TransactOpts, _extensionId, _addresses)
}

// AddExtensionEmergencyPausers is a paid mutator transaction binding the contract method 0x6d02fabc.
//
// Solidity: function addExtensionEmergencyPausers(uint256 _extensionId, address[] _addresses) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactorSession) AddExtensionEmergencyPausers(_extensionId *big.Int, _addresses []common.Address) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.AddExtensionEmergencyPausers(&_MachineEmergencyPause.TransactOpts, _extensionId, _addresses)
}

// AddExtensionEmergencyUnpausers is a paid mutator transaction binding the contract method 0xd2c21c0e.
//
// Solidity: function addExtensionEmergencyUnpausers(uint256 _extensionId, address[] _addresses) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactor) AddExtensionEmergencyUnpausers(opts *bind.TransactOpts, _extensionId *big.Int, _addresses []common.Address) (*types.Transaction, error) {
	return _MachineEmergencyPause.contract.Transact(opts, "addExtensionEmergencyUnpausers", _extensionId, _addresses)
}

// AddExtensionEmergencyUnpausers is a paid mutator transaction binding the contract method 0xd2c21c0e.
//
// Solidity: function addExtensionEmergencyUnpausers(uint256 _extensionId, address[] _addresses) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseSession) AddExtensionEmergencyUnpausers(_extensionId *big.Int, _addresses []common.Address) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.AddExtensionEmergencyUnpausers(&_MachineEmergencyPause.TransactOpts, _extensionId, _addresses)
}

// AddExtensionEmergencyUnpausers is a paid mutator transaction binding the contract method 0xd2c21c0e.
//
// Solidity: function addExtensionEmergencyUnpausers(uint256 _extensionId, address[] _addresses) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactorSession) AddExtensionEmergencyUnpausers(_extensionId *big.Int, _addresses []common.Address) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.AddExtensionEmergencyUnpausers(&_MachineEmergencyPause.TransactOpts, _extensionId, _addresses)
}

// EmergencyPauseExtension is a paid mutator transaction binding the contract method 0x54a77684.
//
// Solidity: function emergencyPauseExtension(uint256 _extensionId) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactor) EmergencyPauseExtension(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _MachineEmergencyPause.contract.Transact(opts, "emergencyPauseExtension", _extensionId)
}

// EmergencyPauseExtension is a paid mutator transaction binding the contract method 0x54a77684.
//
// Solidity: function emergencyPauseExtension(uint256 _extensionId) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseSession) EmergencyPauseExtension(_extensionId *big.Int) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.EmergencyPauseExtension(&_MachineEmergencyPause.TransactOpts, _extensionId)
}

// EmergencyPauseExtension is a paid mutator transaction binding the contract method 0x54a77684.
//
// Solidity: function emergencyPauseExtension(uint256 _extensionId) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactorSession) EmergencyPauseExtension(_extensionId *big.Int) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.EmergencyPauseExtension(&_MachineEmergencyPause.TransactOpts, _extensionId)
}

// EmergencyUnpauseExtension is a paid mutator transaction binding the contract method 0x255fb084.
//
// Solidity: function emergencyUnpauseExtension(uint256 _extensionId) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactor) EmergencyUnpauseExtension(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _MachineEmergencyPause.contract.Transact(opts, "emergencyUnpauseExtension", _extensionId)
}

// EmergencyUnpauseExtension is a paid mutator transaction binding the contract method 0x255fb084.
//
// Solidity: function emergencyUnpauseExtension(uint256 _extensionId) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseSession) EmergencyUnpauseExtension(_extensionId *big.Int) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.EmergencyUnpauseExtension(&_MachineEmergencyPause.TransactOpts, _extensionId)
}

// EmergencyUnpauseExtension is a paid mutator transaction binding the contract method 0x255fb084.
//
// Solidity: function emergencyUnpauseExtension(uint256 _extensionId) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactorSession) EmergencyUnpauseExtension(_extensionId *big.Int) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.EmergencyUnpauseExtension(&_MachineEmergencyPause.TransactOpts, _extensionId)
}

// RemoveExtensionEmergencyPausers is a paid mutator transaction binding the contract method 0xa53840ab.
//
// Solidity: function removeExtensionEmergencyPausers(uint256 _extensionId, address[] _addresses) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactor) RemoveExtensionEmergencyPausers(opts *bind.TransactOpts, _extensionId *big.Int, _addresses []common.Address) (*types.Transaction, error) {
	return _MachineEmergencyPause.contract.Transact(opts, "removeExtensionEmergencyPausers", _extensionId, _addresses)
}

// RemoveExtensionEmergencyPausers is a paid mutator transaction binding the contract method 0xa53840ab.
//
// Solidity: function removeExtensionEmergencyPausers(uint256 _extensionId, address[] _addresses) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseSession) RemoveExtensionEmergencyPausers(_extensionId *big.Int, _addresses []common.Address) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.RemoveExtensionEmergencyPausers(&_MachineEmergencyPause.TransactOpts, _extensionId, _addresses)
}

// RemoveExtensionEmergencyPausers is a paid mutator transaction binding the contract method 0xa53840ab.
//
// Solidity: function removeExtensionEmergencyPausers(uint256 _extensionId, address[] _addresses) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactorSession) RemoveExtensionEmergencyPausers(_extensionId *big.Int, _addresses []common.Address) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.RemoveExtensionEmergencyPausers(&_MachineEmergencyPause.TransactOpts, _extensionId, _addresses)
}

// RemoveExtensionEmergencyUnpausers is a paid mutator transaction binding the contract method 0x98b5cfe0.
//
// Solidity: function removeExtensionEmergencyUnpausers(uint256 _extensionId, address[] _addresses) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactor) RemoveExtensionEmergencyUnpausers(opts *bind.TransactOpts, _extensionId *big.Int, _addresses []common.Address) (*types.Transaction, error) {
	return _MachineEmergencyPause.contract.Transact(opts, "removeExtensionEmergencyUnpausers", _extensionId, _addresses)
}

// RemoveExtensionEmergencyUnpausers is a paid mutator transaction binding the contract method 0x98b5cfe0.
//
// Solidity: function removeExtensionEmergencyUnpausers(uint256 _extensionId, address[] _addresses) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseSession) RemoveExtensionEmergencyUnpausers(_extensionId *big.Int, _addresses []common.Address) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.RemoveExtensionEmergencyUnpausers(&_MachineEmergencyPause.TransactOpts, _extensionId, _addresses)
}

// RemoveExtensionEmergencyUnpausers is a paid mutator transaction binding the contract method 0x98b5cfe0.
//
// Solidity: function removeExtensionEmergencyUnpausers(uint256 _extensionId, address[] _addresses) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactorSession) RemoveExtensionEmergencyUnpausers(_extensionId *big.Int, _addresses []common.Address) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.RemoveExtensionEmergencyUnpausers(&_MachineEmergencyPause.TransactOpts, _extensionId, _addresses)
}

// SetEmergencyUnpauseGracePeriodSeconds is a paid mutator transaction binding the contract method 0x8fb045d5.
//
// Solidity: function setEmergencyUnpauseGracePeriodSeconds(uint256 _seconds) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactor) SetEmergencyUnpauseGracePeriodSeconds(opts *bind.TransactOpts, _seconds *big.Int) (*types.Transaction, error) {
	return _MachineEmergencyPause.contract.Transact(opts, "setEmergencyUnpauseGracePeriodSeconds", _seconds)
}

// SetEmergencyUnpauseGracePeriodSeconds is a paid mutator transaction binding the contract method 0x8fb045d5.
//
// Solidity: function setEmergencyUnpauseGracePeriodSeconds(uint256 _seconds) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseSession) SetEmergencyUnpauseGracePeriodSeconds(_seconds *big.Int) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.SetEmergencyUnpauseGracePeriodSeconds(&_MachineEmergencyPause.TransactOpts, _seconds)
}

// SetEmergencyUnpauseGracePeriodSeconds is a paid mutator transaction binding the contract method 0x8fb045d5.
//
// Solidity: function setEmergencyUnpauseGracePeriodSeconds(uint256 _seconds) returns()
func (_MachineEmergencyPause *MachineEmergencyPauseTransactorSession) SetEmergencyUnpauseGracePeriodSeconds(_seconds *big.Int) (*types.Transaction, error) {
	return _MachineEmergencyPause.Contract.SetEmergencyUnpauseGracePeriodSeconds(&_MachineEmergencyPause.TransactOpts, _seconds)
}

// MachineEmergencyPauseEmergencyUnpauseGracePeriodSetIterator is returned from FilterEmergencyUnpauseGracePeriodSet and is used to iterate over the raw logs and unpacked data for EmergencyUnpauseGracePeriodSet events raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseEmergencyUnpauseGracePeriodSetIterator struct {
	Event *MachineEmergencyPauseEmergencyUnpauseGracePeriodSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MachineEmergencyPauseEmergencyUnpauseGracePeriodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineEmergencyPauseEmergencyUnpauseGracePeriodSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MachineEmergencyPauseEmergencyUnpauseGracePeriodSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MachineEmergencyPauseEmergencyUnpauseGracePeriodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineEmergencyPauseEmergencyUnpauseGracePeriodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineEmergencyPauseEmergencyUnpauseGracePeriodSet represents a EmergencyUnpauseGracePeriodSet event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseEmergencyUnpauseGracePeriodSet struct {
	GraceSeconds *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmergencyUnpauseGracePeriodSet is a free log retrieval operation binding the contract event 0xd3e133eed54192c593c1994bcda078ca197c9b9af58465e5e0d84296e23f0b16.
//
// Solidity: event EmergencyUnpauseGracePeriodSet(uint256 graceSeconds)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) FilterEmergencyUnpauseGracePeriodSet(opts *bind.FilterOpts) (*MachineEmergencyPauseEmergencyUnpauseGracePeriodSetIterator, error) {

	logs, sub, err := _MachineEmergencyPause.contract.FilterLogs(opts, "EmergencyUnpauseGracePeriodSet")
	if err != nil {
		return nil, err
	}
	return &MachineEmergencyPauseEmergencyUnpauseGracePeriodSetIterator{contract: _MachineEmergencyPause.contract, event: "EmergencyUnpauseGracePeriodSet", logs: logs, sub: sub}, nil
}

// WatchEmergencyUnpauseGracePeriodSet is a free log subscription operation binding the contract event 0xd3e133eed54192c593c1994bcda078ca197c9b9af58465e5e0d84296e23f0b16.
//
// Solidity: event EmergencyUnpauseGracePeriodSet(uint256 graceSeconds)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) WatchEmergencyUnpauseGracePeriodSet(opts *bind.WatchOpts, sink chan<- *MachineEmergencyPauseEmergencyUnpauseGracePeriodSet) (event.Subscription, error) {

	logs, sub, err := _MachineEmergencyPause.contract.WatchLogs(opts, "EmergencyUnpauseGracePeriodSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineEmergencyPauseEmergencyUnpauseGracePeriodSet)
				if err := _MachineEmergencyPause.contract.UnpackLog(event, "EmergencyUnpauseGracePeriodSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmergencyUnpauseGracePeriodSet is a log parse operation binding the contract event 0xd3e133eed54192c593c1994bcda078ca197c9b9af58465e5e0d84296e23f0b16.
//
// Solidity: event EmergencyUnpauseGracePeriodSet(uint256 graceSeconds)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) ParseEmergencyUnpauseGracePeriodSet(log types.Log) (*MachineEmergencyPauseEmergencyUnpauseGracePeriodSet, error) {
	event := new(MachineEmergencyPauseEmergencyUnpauseGracePeriodSet)
	if err := _MachineEmergencyPause.contract.UnpackLog(event, "EmergencyUnpauseGracePeriodSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineEmergencyPauseExtensionEmergencyPausedIterator is returned from FilterExtensionEmergencyPaused and is used to iterate over the raw logs and unpacked data for ExtensionEmergencyPaused events raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyPausedIterator struct {
	Event *MachineEmergencyPauseExtensionEmergencyPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MachineEmergencyPauseExtensionEmergencyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineEmergencyPauseExtensionEmergencyPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MachineEmergencyPauseExtensionEmergencyPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MachineEmergencyPauseExtensionEmergencyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineEmergencyPauseExtensionEmergencyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineEmergencyPauseExtensionEmergencyPaused represents a ExtensionEmergencyPaused event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyPaused struct {
	ExtensionId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExtensionEmergencyPaused is a free log retrieval operation binding the contract event 0x1b344be4973044de8ecc45085b9f5c2716b7f9d81ccc7ee616ed7889e6fd0b89.
//
// Solidity: event ExtensionEmergencyPaused(uint256 indexed extensionId)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) FilterExtensionEmergencyPaused(opts *bind.FilterOpts, extensionId []*big.Int) (*MachineEmergencyPauseExtensionEmergencyPausedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _MachineEmergencyPause.contract.FilterLogs(opts, "ExtensionEmergencyPaused", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return &MachineEmergencyPauseExtensionEmergencyPausedIterator{contract: _MachineEmergencyPause.contract, event: "ExtensionEmergencyPaused", logs: logs, sub: sub}, nil
}

// WatchExtensionEmergencyPaused is a free log subscription operation binding the contract event 0x1b344be4973044de8ecc45085b9f5c2716b7f9d81ccc7ee616ed7889e6fd0b89.
//
// Solidity: event ExtensionEmergencyPaused(uint256 indexed extensionId)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) WatchExtensionEmergencyPaused(opts *bind.WatchOpts, sink chan<- *MachineEmergencyPauseExtensionEmergencyPaused, extensionId []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _MachineEmergencyPause.contract.WatchLogs(opts, "ExtensionEmergencyPaused", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineEmergencyPauseExtensionEmergencyPaused)
				if err := _MachineEmergencyPause.contract.UnpackLog(event, "ExtensionEmergencyPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtensionEmergencyPaused is a log parse operation binding the contract event 0x1b344be4973044de8ecc45085b9f5c2716b7f9d81ccc7ee616ed7889e6fd0b89.
//
// Solidity: event ExtensionEmergencyPaused(uint256 indexed extensionId)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) ParseExtensionEmergencyPaused(log types.Log) (*MachineEmergencyPauseExtensionEmergencyPaused, error) {
	event := new(MachineEmergencyPauseExtensionEmergencyPaused)
	if err := _MachineEmergencyPause.contract.UnpackLog(event, "ExtensionEmergencyPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineEmergencyPauseExtensionEmergencyPausersAddedIterator is returned from FilterExtensionEmergencyPausersAdded and is used to iterate over the raw logs and unpacked data for ExtensionEmergencyPausersAdded events raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyPausersAddedIterator struct {
	Event *MachineEmergencyPauseExtensionEmergencyPausersAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MachineEmergencyPauseExtensionEmergencyPausersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineEmergencyPauseExtensionEmergencyPausersAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MachineEmergencyPauseExtensionEmergencyPausersAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MachineEmergencyPauseExtensionEmergencyPausersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineEmergencyPauseExtensionEmergencyPausersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineEmergencyPauseExtensionEmergencyPausersAdded represents a ExtensionEmergencyPausersAdded event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyPausersAdded struct {
	ExtensionId *big.Int
	Addresses   []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExtensionEmergencyPausersAdded is a free log retrieval operation binding the contract event 0x8564c6eaa4c74ee321aeb69d83f505548d2383a35fb9507ddde73dada0907243.
//
// Solidity: event ExtensionEmergencyPausersAdded(uint256 indexed extensionId, address[] addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) FilterExtensionEmergencyPausersAdded(opts *bind.FilterOpts, extensionId []*big.Int) (*MachineEmergencyPauseExtensionEmergencyPausersAddedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _MachineEmergencyPause.contract.FilterLogs(opts, "ExtensionEmergencyPausersAdded", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return &MachineEmergencyPauseExtensionEmergencyPausersAddedIterator{contract: _MachineEmergencyPause.contract, event: "ExtensionEmergencyPausersAdded", logs: logs, sub: sub}, nil
}

// WatchExtensionEmergencyPausersAdded is a free log subscription operation binding the contract event 0x8564c6eaa4c74ee321aeb69d83f505548d2383a35fb9507ddde73dada0907243.
//
// Solidity: event ExtensionEmergencyPausersAdded(uint256 indexed extensionId, address[] addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) WatchExtensionEmergencyPausersAdded(opts *bind.WatchOpts, sink chan<- *MachineEmergencyPauseExtensionEmergencyPausersAdded, extensionId []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _MachineEmergencyPause.contract.WatchLogs(opts, "ExtensionEmergencyPausersAdded", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineEmergencyPauseExtensionEmergencyPausersAdded)
				if err := _MachineEmergencyPause.contract.UnpackLog(event, "ExtensionEmergencyPausersAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtensionEmergencyPausersAdded is a log parse operation binding the contract event 0x8564c6eaa4c74ee321aeb69d83f505548d2383a35fb9507ddde73dada0907243.
//
// Solidity: event ExtensionEmergencyPausersAdded(uint256 indexed extensionId, address[] addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) ParseExtensionEmergencyPausersAdded(log types.Log) (*MachineEmergencyPauseExtensionEmergencyPausersAdded, error) {
	event := new(MachineEmergencyPauseExtensionEmergencyPausersAdded)
	if err := _MachineEmergencyPause.contract.UnpackLog(event, "ExtensionEmergencyPausersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineEmergencyPauseExtensionEmergencyPausersRemovedIterator is returned from FilterExtensionEmergencyPausersRemoved and is used to iterate over the raw logs and unpacked data for ExtensionEmergencyPausersRemoved events raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyPausersRemovedIterator struct {
	Event *MachineEmergencyPauseExtensionEmergencyPausersRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MachineEmergencyPauseExtensionEmergencyPausersRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineEmergencyPauseExtensionEmergencyPausersRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MachineEmergencyPauseExtensionEmergencyPausersRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MachineEmergencyPauseExtensionEmergencyPausersRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineEmergencyPauseExtensionEmergencyPausersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineEmergencyPauseExtensionEmergencyPausersRemoved represents a ExtensionEmergencyPausersRemoved event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyPausersRemoved struct {
	ExtensionId *big.Int
	Addresses   []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExtensionEmergencyPausersRemoved is a free log retrieval operation binding the contract event 0x6d16af833cb648d9ffdd60ba9299fa7c2c6a9daf00e87589326c00e56dffed08.
//
// Solidity: event ExtensionEmergencyPausersRemoved(uint256 indexed extensionId, address[] addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) FilterExtensionEmergencyPausersRemoved(opts *bind.FilterOpts, extensionId []*big.Int) (*MachineEmergencyPauseExtensionEmergencyPausersRemovedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _MachineEmergencyPause.contract.FilterLogs(opts, "ExtensionEmergencyPausersRemoved", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return &MachineEmergencyPauseExtensionEmergencyPausersRemovedIterator{contract: _MachineEmergencyPause.contract, event: "ExtensionEmergencyPausersRemoved", logs: logs, sub: sub}, nil
}

// WatchExtensionEmergencyPausersRemoved is a free log subscription operation binding the contract event 0x6d16af833cb648d9ffdd60ba9299fa7c2c6a9daf00e87589326c00e56dffed08.
//
// Solidity: event ExtensionEmergencyPausersRemoved(uint256 indexed extensionId, address[] addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) WatchExtensionEmergencyPausersRemoved(opts *bind.WatchOpts, sink chan<- *MachineEmergencyPauseExtensionEmergencyPausersRemoved, extensionId []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _MachineEmergencyPause.contract.WatchLogs(opts, "ExtensionEmergencyPausersRemoved", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineEmergencyPauseExtensionEmergencyPausersRemoved)
				if err := _MachineEmergencyPause.contract.UnpackLog(event, "ExtensionEmergencyPausersRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtensionEmergencyPausersRemoved is a log parse operation binding the contract event 0x6d16af833cb648d9ffdd60ba9299fa7c2c6a9daf00e87589326c00e56dffed08.
//
// Solidity: event ExtensionEmergencyPausersRemoved(uint256 indexed extensionId, address[] addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) ParseExtensionEmergencyPausersRemoved(log types.Log) (*MachineEmergencyPauseExtensionEmergencyPausersRemoved, error) {
	event := new(MachineEmergencyPauseExtensionEmergencyPausersRemoved)
	if err := _MachineEmergencyPause.contract.UnpackLog(event, "ExtensionEmergencyPausersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineEmergencyPauseExtensionEmergencyUnpausedIterator is returned from FilterExtensionEmergencyUnpaused and is used to iterate over the raw logs and unpacked data for ExtensionEmergencyUnpaused events raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyUnpausedIterator struct {
	Event *MachineEmergencyPauseExtensionEmergencyUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MachineEmergencyPauseExtensionEmergencyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineEmergencyPauseExtensionEmergencyUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MachineEmergencyPauseExtensionEmergencyUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MachineEmergencyPauseExtensionEmergencyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineEmergencyPauseExtensionEmergencyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineEmergencyPauseExtensionEmergencyUnpaused represents a ExtensionEmergencyUnpaused event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyUnpaused struct {
	ExtensionId *big.Int
	UnpauseTs   uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExtensionEmergencyUnpaused is a free log retrieval operation binding the contract event 0x0c1c08d84f944ffe762960a9bde1be060b21a106f2f6fb70ae7d57a0d8653653.
//
// Solidity: event ExtensionEmergencyUnpaused(uint256 indexed extensionId, uint64 unpauseTs)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) FilterExtensionEmergencyUnpaused(opts *bind.FilterOpts, extensionId []*big.Int) (*MachineEmergencyPauseExtensionEmergencyUnpausedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _MachineEmergencyPause.contract.FilterLogs(opts, "ExtensionEmergencyUnpaused", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return &MachineEmergencyPauseExtensionEmergencyUnpausedIterator{contract: _MachineEmergencyPause.contract, event: "ExtensionEmergencyUnpaused", logs: logs, sub: sub}, nil
}

// WatchExtensionEmergencyUnpaused is a free log subscription operation binding the contract event 0x0c1c08d84f944ffe762960a9bde1be060b21a106f2f6fb70ae7d57a0d8653653.
//
// Solidity: event ExtensionEmergencyUnpaused(uint256 indexed extensionId, uint64 unpauseTs)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) WatchExtensionEmergencyUnpaused(opts *bind.WatchOpts, sink chan<- *MachineEmergencyPauseExtensionEmergencyUnpaused, extensionId []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _MachineEmergencyPause.contract.WatchLogs(opts, "ExtensionEmergencyUnpaused", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineEmergencyPauseExtensionEmergencyUnpaused)
				if err := _MachineEmergencyPause.contract.UnpackLog(event, "ExtensionEmergencyUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtensionEmergencyUnpaused is a log parse operation binding the contract event 0x0c1c08d84f944ffe762960a9bde1be060b21a106f2f6fb70ae7d57a0d8653653.
//
// Solidity: event ExtensionEmergencyUnpaused(uint256 indexed extensionId, uint64 unpauseTs)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) ParseExtensionEmergencyUnpaused(log types.Log) (*MachineEmergencyPauseExtensionEmergencyUnpaused, error) {
	event := new(MachineEmergencyPauseExtensionEmergencyUnpaused)
	if err := _MachineEmergencyPause.contract.UnpackLog(event, "ExtensionEmergencyUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineEmergencyPauseExtensionEmergencyUnpausersAddedIterator is returned from FilterExtensionEmergencyUnpausersAdded and is used to iterate over the raw logs and unpacked data for ExtensionEmergencyUnpausersAdded events raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyUnpausersAddedIterator struct {
	Event *MachineEmergencyPauseExtensionEmergencyUnpausersAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MachineEmergencyPauseExtensionEmergencyUnpausersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineEmergencyPauseExtensionEmergencyUnpausersAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MachineEmergencyPauseExtensionEmergencyUnpausersAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MachineEmergencyPauseExtensionEmergencyUnpausersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineEmergencyPauseExtensionEmergencyUnpausersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineEmergencyPauseExtensionEmergencyUnpausersAdded represents a ExtensionEmergencyUnpausersAdded event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyUnpausersAdded struct {
	ExtensionId *big.Int
	Addresses   []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExtensionEmergencyUnpausersAdded is a free log retrieval operation binding the contract event 0x5565f4ced151a1301327e18a390de31e626be39fd9616437af2e9213c58a4154.
//
// Solidity: event ExtensionEmergencyUnpausersAdded(uint256 indexed extensionId, address[] addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) FilterExtensionEmergencyUnpausersAdded(opts *bind.FilterOpts, extensionId []*big.Int) (*MachineEmergencyPauseExtensionEmergencyUnpausersAddedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _MachineEmergencyPause.contract.FilterLogs(opts, "ExtensionEmergencyUnpausersAdded", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return &MachineEmergencyPauseExtensionEmergencyUnpausersAddedIterator{contract: _MachineEmergencyPause.contract, event: "ExtensionEmergencyUnpausersAdded", logs: logs, sub: sub}, nil
}

// WatchExtensionEmergencyUnpausersAdded is a free log subscription operation binding the contract event 0x5565f4ced151a1301327e18a390de31e626be39fd9616437af2e9213c58a4154.
//
// Solidity: event ExtensionEmergencyUnpausersAdded(uint256 indexed extensionId, address[] addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) WatchExtensionEmergencyUnpausersAdded(opts *bind.WatchOpts, sink chan<- *MachineEmergencyPauseExtensionEmergencyUnpausersAdded, extensionId []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _MachineEmergencyPause.contract.WatchLogs(opts, "ExtensionEmergencyUnpausersAdded", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineEmergencyPauseExtensionEmergencyUnpausersAdded)
				if err := _MachineEmergencyPause.contract.UnpackLog(event, "ExtensionEmergencyUnpausersAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtensionEmergencyUnpausersAdded is a log parse operation binding the contract event 0x5565f4ced151a1301327e18a390de31e626be39fd9616437af2e9213c58a4154.
//
// Solidity: event ExtensionEmergencyUnpausersAdded(uint256 indexed extensionId, address[] addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) ParseExtensionEmergencyUnpausersAdded(log types.Log) (*MachineEmergencyPauseExtensionEmergencyUnpausersAdded, error) {
	event := new(MachineEmergencyPauseExtensionEmergencyUnpausersAdded)
	if err := _MachineEmergencyPause.contract.UnpackLog(event, "ExtensionEmergencyUnpausersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineEmergencyPauseExtensionEmergencyUnpausersRemovedIterator is returned from FilterExtensionEmergencyUnpausersRemoved and is used to iterate over the raw logs and unpacked data for ExtensionEmergencyUnpausersRemoved events raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyUnpausersRemovedIterator struct {
	Event *MachineEmergencyPauseExtensionEmergencyUnpausersRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MachineEmergencyPauseExtensionEmergencyUnpausersRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineEmergencyPauseExtensionEmergencyUnpausersRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MachineEmergencyPauseExtensionEmergencyUnpausersRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MachineEmergencyPauseExtensionEmergencyUnpausersRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineEmergencyPauseExtensionEmergencyUnpausersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineEmergencyPauseExtensionEmergencyUnpausersRemoved represents a ExtensionEmergencyUnpausersRemoved event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseExtensionEmergencyUnpausersRemoved struct {
	ExtensionId *big.Int
	Addresses   []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExtensionEmergencyUnpausersRemoved is a free log retrieval operation binding the contract event 0xd4f869e52dfb90a1ace337f7f9b4677fe210c1e2679c344b1f78a9b0bbd5cf11.
//
// Solidity: event ExtensionEmergencyUnpausersRemoved(uint256 indexed extensionId, address[] addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) FilterExtensionEmergencyUnpausersRemoved(opts *bind.FilterOpts, extensionId []*big.Int) (*MachineEmergencyPauseExtensionEmergencyUnpausersRemovedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _MachineEmergencyPause.contract.FilterLogs(opts, "ExtensionEmergencyUnpausersRemoved", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return &MachineEmergencyPauseExtensionEmergencyUnpausersRemovedIterator{contract: _MachineEmergencyPause.contract, event: "ExtensionEmergencyUnpausersRemoved", logs: logs, sub: sub}, nil
}

// WatchExtensionEmergencyUnpausersRemoved is a free log subscription operation binding the contract event 0xd4f869e52dfb90a1ace337f7f9b4677fe210c1e2679c344b1f78a9b0bbd5cf11.
//
// Solidity: event ExtensionEmergencyUnpausersRemoved(uint256 indexed extensionId, address[] addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) WatchExtensionEmergencyUnpausersRemoved(opts *bind.WatchOpts, sink chan<- *MachineEmergencyPauseExtensionEmergencyUnpausersRemoved, extensionId []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _MachineEmergencyPause.contract.WatchLogs(opts, "ExtensionEmergencyUnpausersRemoved", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineEmergencyPauseExtensionEmergencyUnpausersRemoved)
				if err := _MachineEmergencyPause.contract.UnpackLog(event, "ExtensionEmergencyUnpausersRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExtensionEmergencyUnpausersRemoved is a log parse operation binding the contract event 0xd4f869e52dfb90a1ace337f7f9b4677fe210c1e2679c344b1f78a9b0bbd5cf11.
//
// Solidity: event ExtensionEmergencyUnpausersRemoved(uint256 indexed extensionId, address[] addresses)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) ParseExtensionEmergencyUnpausersRemoved(log types.Log) (*MachineEmergencyPauseExtensionEmergencyUnpausersRemoved, error) {
	event := new(MachineEmergencyPauseExtensionEmergencyUnpausersRemoved)
	if err := _MachineEmergencyPause.contract.UnpackLog(event, "ExtensionEmergencyUnpausersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineEmergencyPauseGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseGovernanceCallTimelockedIterator struct {
	Event *MachineEmergencyPauseGovernanceCallTimelocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MachineEmergencyPauseGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineEmergencyPauseGovernanceCallTimelocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MachineEmergencyPauseGovernanceCallTimelocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MachineEmergencyPauseGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineEmergencyPauseGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineEmergencyPauseGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*MachineEmergencyPauseGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _MachineEmergencyPause.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &MachineEmergencyPauseGovernanceCallTimelockedIterator{contract: _MachineEmergencyPause.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *MachineEmergencyPauseGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _MachineEmergencyPause.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineEmergencyPauseGovernanceCallTimelocked)
				if err := _MachineEmergencyPause.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) ParseGovernanceCallTimelocked(log types.Log) (*MachineEmergencyPauseGovernanceCallTimelocked, error) {
	event := new(MachineEmergencyPauseGovernanceCallTimelocked)
	if err := _MachineEmergencyPause.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineEmergencyPauseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInitializedIterator struct {
	Event *MachineEmergencyPauseInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MachineEmergencyPauseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineEmergencyPauseInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MachineEmergencyPauseInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MachineEmergencyPauseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineEmergencyPauseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineEmergencyPauseInitialized represents a Initialized event raised by the MachineEmergencyPause contract.
type MachineEmergencyPauseInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) FilterInitialized(opts *bind.FilterOpts) (*MachineEmergencyPauseInitializedIterator, error) {

	logs, sub, err := _MachineEmergencyPause.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MachineEmergencyPauseInitializedIterator{contract: _MachineEmergencyPause.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MachineEmergencyPauseInitialized) (event.Subscription, error) {

	logs, sub, err := _MachineEmergencyPause.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineEmergencyPauseInitialized)
				if err := _MachineEmergencyPause.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MachineEmergencyPause *MachineEmergencyPauseFilterer) ParseInitialized(log types.Log) (*MachineEmergencyPauseInitialized, error) {
	event := new(MachineEmergencyPauseInitialized)
	if err := _MachineEmergencyPause.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
