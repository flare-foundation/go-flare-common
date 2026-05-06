// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ownerallowlist

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

// OwnerAllowlistMetaData contains all meta data concerning the OwnerAllowlist contract.
var OwnerAllowlistMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAlreadyAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerNotInAllowlist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeMachineOwnersAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeMachineOwnersDisallowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeWalletProjectOwnersAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeWalletProjectOwnersDisallowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeMachineOwnersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeMachineOwnersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeWalletProjectOwnersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeWalletProjectOwnersRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"addAllowedTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"addAllowedTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allTeeMachineOwnersAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_allAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allTeeWalletProjectOwnersAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_allAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allowAllTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allowAllTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"disallowAllTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"disallowAllTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getAllowedTeeMachineOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_allowedOwners\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getAllowedTeeWalletProjectOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_allowedOwners\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"isAllowedTeeMachineOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"isAllowedTeeWalletProjectOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"removeAllowedTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"removeAllowedTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnerAllowlistABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnerAllowlistMetaData.ABI instead.
var OwnerAllowlistABI = OwnerAllowlistMetaData.ABI

// OwnerAllowlist is an auto generated Go binding around an Ethereum contract.
type OwnerAllowlist struct {
	OwnerAllowlistCaller     // Read-only binding to the contract
	OwnerAllowlistTransactor // Write-only binding to the contract
	OwnerAllowlistFilterer   // Log filterer for contract events
}

// OwnerAllowlistCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnerAllowlistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerAllowlistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnerAllowlistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerAllowlistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnerAllowlistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerAllowlistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnerAllowlistSession struct {
	Contract     *OwnerAllowlist   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnerAllowlistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnerAllowlistCallerSession struct {
	Contract *OwnerAllowlistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OwnerAllowlistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnerAllowlistTransactorSession struct {
	Contract     *OwnerAllowlistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OwnerAllowlistRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnerAllowlistRaw struct {
	Contract *OwnerAllowlist // Generic contract binding to access the raw methods on
}

// OwnerAllowlistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnerAllowlistCallerRaw struct {
	Contract *OwnerAllowlistCaller // Generic read-only contract binding to access the raw methods on
}

// OwnerAllowlistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnerAllowlistTransactorRaw struct {
	Contract *OwnerAllowlistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnerAllowlist creates a new instance of OwnerAllowlist, bound to a specific deployed contract.
func NewOwnerAllowlist(address common.Address, backend bind.ContractBackend) (*OwnerAllowlist, error) {
	contract, err := bindOwnerAllowlist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnerAllowlist{OwnerAllowlistCaller: OwnerAllowlistCaller{contract: contract}, OwnerAllowlistTransactor: OwnerAllowlistTransactor{contract: contract}, OwnerAllowlistFilterer: OwnerAllowlistFilterer{contract: contract}}, nil
}

// NewOwnerAllowlistCaller creates a new read-only instance of OwnerAllowlist, bound to a specific deployed contract.
func NewOwnerAllowlistCaller(address common.Address, caller bind.ContractCaller) (*OwnerAllowlistCaller, error) {
	contract, err := bindOwnerAllowlist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerAllowlistCaller{contract: contract}, nil
}

// NewOwnerAllowlistTransactor creates a new write-only instance of OwnerAllowlist, bound to a specific deployed contract.
func NewOwnerAllowlistTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnerAllowlistTransactor, error) {
	contract, err := bindOwnerAllowlist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerAllowlistTransactor{contract: contract}, nil
}

// NewOwnerAllowlistFilterer creates a new log filterer instance of OwnerAllowlist, bound to a specific deployed contract.
func NewOwnerAllowlistFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnerAllowlistFilterer, error) {
	contract, err := bindOwnerAllowlist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnerAllowlistFilterer{contract: contract}, nil
}

// bindOwnerAllowlist binds a generic wrapper to an already deployed contract.
func bindOwnerAllowlist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnerAllowlistMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnerAllowlist *OwnerAllowlistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnerAllowlist.Contract.OwnerAllowlistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnerAllowlist *OwnerAllowlistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.OwnerAllowlistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnerAllowlist *OwnerAllowlistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.OwnerAllowlistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnerAllowlist *OwnerAllowlistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnerAllowlist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnerAllowlist *OwnerAllowlistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnerAllowlist *OwnerAllowlistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.contract.Transact(opts, method, params...)
}

// AllTeeMachineOwnersAllowed is a free data retrieval call binding the contract method 0x2a30b9f4.
//
// Solidity: function allTeeMachineOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (_OwnerAllowlist *OwnerAllowlistCaller) AllTeeMachineOwnersAllowed(opts *bind.CallOpts, _extensionId *big.Int) (bool, error) {
	var out []interface{}
	err := _OwnerAllowlist.contract.Call(opts, &out, "allTeeMachineOwnersAllowed", _extensionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllTeeMachineOwnersAllowed is a free data retrieval call binding the contract method 0x2a30b9f4.
//
// Solidity: function allTeeMachineOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (_OwnerAllowlist *OwnerAllowlistSession) AllTeeMachineOwnersAllowed(_extensionId *big.Int) (bool, error) {
	return _OwnerAllowlist.Contract.AllTeeMachineOwnersAllowed(&_OwnerAllowlist.CallOpts, _extensionId)
}

// AllTeeMachineOwnersAllowed is a free data retrieval call binding the contract method 0x2a30b9f4.
//
// Solidity: function allTeeMachineOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (_OwnerAllowlist *OwnerAllowlistCallerSession) AllTeeMachineOwnersAllowed(_extensionId *big.Int) (bool, error) {
	return _OwnerAllowlist.Contract.AllTeeMachineOwnersAllowed(&_OwnerAllowlist.CallOpts, _extensionId)
}

// AllTeeWalletProjectOwnersAllowed is a free data retrieval call binding the contract method 0x8398fc2e.
//
// Solidity: function allTeeWalletProjectOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (_OwnerAllowlist *OwnerAllowlistCaller) AllTeeWalletProjectOwnersAllowed(opts *bind.CallOpts, _extensionId *big.Int) (bool, error) {
	var out []interface{}
	err := _OwnerAllowlist.contract.Call(opts, &out, "allTeeWalletProjectOwnersAllowed", _extensionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllTeeWalletProjectOwnersAllowed is a free data retrieval call binding the contract method 0x8398fc2e.
//
// Solidity: function allTeeWalletProjectOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (_OwnerAllowlist *OwnerAllowlistSession) AllTeeWalletProjectOwnersAllowed(_extensionId *big.Int) (bool, error) {
	return _OwnerAllowlist.Contract.AllTeeWalletProjectOwnersAllowed(&_OwnerAllowlist.CallOpts, _extensionId)
}

// AllTeeWalletProjectOwnersAllowed is a free data retrieval call binding the contract method 0x8398fc2e.
//
// Solidity: function allTeeWalletProjectOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (_OwnerAllowlist *OwnerAllowlistCallerSession) AllTeeWalletProjectOwnersAllowed(_extensionId *big.Int) (bool, error) {
	return _OwnerAllowlist.Contract.AllTeeWalletProjectOwnersAllowed(&_OwnerAllowlist.CallOpts, _extensionId)
}

// GetAllowedTeeMachineOwners is a free data retrieval call binding the contract method 0x5f37ea30.
//
// Solidity: function getAllowedTeeMachineOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (_OwnerAllowlist *OwnerAllowlistCaller) GetAllowedTeeMachineOwners(opts *bind.CallOpts, _extensionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _OwnerAllowlist.contract.Call(opts, &out, "getAllowedTeeMachineOwners", _extensionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowedTeeMachineOwners is a free data retrieval call binding the contract method 0x5f37ea30.
//
// Solidity: function getAllowedTeeMachineOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (_OwnerAllowlist *OwnerAllowlistSession) GetAllowedTeeMachineOwners(_extensionId *big.Int) ([]common.Address, error) {
	return _OwnerAllowlist.Contract.GetAllowedTeeMachineOwners(&_OwnerAllowlist.CallOpts, _extensionId)
}

// GetAllowedTeeMachineOwners is a free data retrieval call binding the contract method 0x5f37ea30.
//
// Solidity: function getAllowedTeeMachineOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (_OwnerAllowlist *OwnerAllowlistCallerSession) GetAllowedTeeMachineOwners(_extensionId *big.Int) ([]common.Address, error) {
	return _OwnerAllowlist.Contract.GetAllowedTeeMachineOwners(&_OwnerAllowlist.CallOpts, _extensionId)
}

// GetAllowedTeeWalletProjectOwners is a free data retrieval call binding the contract method 0xd11340ed.
//
// Solidity: function getAllowedTeeWalletProjectOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (_OwnerAllowlist *OwnerAllowlistCaller) GetAllowedTeeWalletProjectOwners(opts *bind.CallOpts, _extensionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _OwnerAllowlist.contract.Call(opts, &out, "getAllowedTeeWalletProjectOwners", _extensionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowedTeeWalletProjectOwners is a free data retrieval call binding the contract method 0xd11340ed.
//
// Solidity: function getAllowedTeeWalletProjectOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (_OwnerAllowlist *OwnerAllowlistSession) GetAllowedTeeWalletProjectOwners(_extensionId *big.Int) ([]common.Address, error) {
	return _OwnerAllowlist.Contract.GetAllowedTeeWalletProjectOwners(&_OwnerAllowlist.CallOpts, _extensionId)
}

// GetAllowedTeeWalletProjectOwners is a free data retrieval call binding the contract method 0xd11340ed.
//
// Solidity: function getAllowedTeeWalletProjectOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (_OwnerAllowlist *OwnerAllowlistCallerSession) GetAllowedTeeWalletProjectOwners(_extensionId *big.Int) ([]common.Address, error) {
	return _OwnerAllowlist.Contract.GetAllowedTeeWalletProjectOwners(&_OwnerAllowlist.CallOpts, _extensionId)
}

// IsAllowedTeeMachineOwner is a free data retrieval call binding the contract method 0x84f5da36.
//
// Solidity: function isAllowedTeeMachineOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_OwnerAllowlist *OwnerAllowlistCaller) IsAllowedTeeMachineOwner(opts *bind.CallOpts, _extensionId *big.Int, _owner common.Address) (bool, error) {
	var out []interface{}
	err := _OwnerAllowlist.contract.Call(opts, &out, "isAllowedTeeMachineOwner", _extensionId, _owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedTeeMachineOwner is a free data retrieval call binding the contract method 0x84f5da36.
//
// Solidity: function isAllowedTeeMachineOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_OwnerAllowlist *OwnerAllowlistSession) IsAllowedTeeMachineOwner(_extensionId *big.Int, _owner common.Address) (bool, error) {
	return _OwnerAllowlist.Contract.IsAllowedTeeMachineOwner(&_OwnerAllowlist.CallOpts, _extensionId, _owner)
}

// IsAllowedTeeMachineOwner is a free data retrieval call binding the contract method 0x84f5da36.
//
// Solidity: function isAllowedTeeMachineOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_OwnerAllowlist *OwnerAllowlistCallerSession) IsAllowedTeeMachineOwner(_extensionId *big.Int, _owner common.Address) (bool, error) {
	return _OwnerAllowlist.Contract.IsAllowedTeeMachineOwner(&_OwnerAllowlist.CallOpts, _extensionId, _owner)
}

// IsAllowedTeeWalletProjectOwner is a free data retrieval call binding the contract method 0xab1c7ea1.
//
// Solidity: function isAllowedTeeWalletProjectOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_OwnerAllowlist *OwnerAllowlistCaller) IsAllowedTeeWalletProjectOwner(opts *bind.CallOpts, _extensionId *big.Int, _owner common.Address) (bool, error) {
	var out []interface{}
	err := _OwnerAllowlist.contract.Call(opts, &out, "isAllowedTeeWalletProjectOwner", _extensionId, _owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedTeeWalletProjectOwner is a free data retrieval call binding the contract method 0xab1c7ea1.
//
// Solidity: function isAllowedTeeWalletProjectOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_OwnerAllowlist *OwnerAllowlistSession) IsAllowedTeeWalletProjectOwner(_extensionId *big.Int, _owner common.Address) (bool, error) {
	return _OwnerAllowlist.Contract.IsAllowedTeeWalletProjectOwner(&_OwnerAllowlist.CallOpts, _extensionId, _owner)
}

// IsAllowedTeeWalletProjectOwner is a free data retrieval call binding the contract method 0xab1c7ea1.
//
// Solidity: function isAllowedTeeWalletProjectOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_OwnerAllowlist *OwnerAllowlistCallerSession) IsAllowedTeeWalletProjectOwner(_extensionId *big.Int, _owner common.Address) (bool, error) {
	return _OwnerAllowlist.Contract.IsAllowedTeeWalletProjectOwner(&_OwnerAllowlist.CallOpts, _extensionId, _owner)
}

// AddAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0xa8f3cab7.
//
// Solidity: function addAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactor) AddAllowedTeeMachineOwners(opts *bind.TransactOpts, _extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _OwnerAllowlist.contract.Transact(opts, "addAllowedTeeMachineOwners", _extensionId, _owners)
}

// AddAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0xa8f3cab7.
//
// Solidity: function addAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_OwnerAllowlist *OwnerAllowlistSession) AddAllowedTeeMachineOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.AddAllowedTeeMachineOwners(&_OwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// AddAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0xa8f3cab7.
//
// Solidity: function addAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactorSession) AddAllowedTeeMachineOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.AddAllowedTeeMachineOwners(&_OwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// AddAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0xd1c34746.
//
// Solidity: function addAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactor) AddAllowedTeeWalletProjectOwners(opts *bind.TransactOpts, _extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _OwnerAllowlist.contract.Transact(opts, "addAllowedTeeWalletProjectOwners", _extensionId, _owners)
}

// AddAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0xd1c34746.
//
// Solidity: function addAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_OwnerAllowlist *OwnerAllowlistSession) AddAllowedTeeWalletProjectOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.AddAllowedTeeWalletProjectOwners(&_OwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// AddAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0xd1c34746.
//
// Solidity: function addAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactorSession) AddAllowedTeeWalletProjectOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.AddAllowedTeeWalletProjectOwners(&_OwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// AllowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0x1b489f2d.
//
// Solidity: function allowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactor) AllowAllTeeMachineOwners(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _OwnerAllowlist.contract.Transact(opts, "allowAllTeeMachineOwners", _extensionId)
}

// AllowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0x1b489f2d.
//
// Solidity: function allowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_OwnerAllowlist *OwnerAllowlistSession) AllowAllTeeMachineOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.AllowAllTeeMachineOwners(&_OwnerAllowlist.TransactOpts, _extensionId)
}

// AllowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0x1b489f2d.
//
// Solidity: function allowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactorSession) AllowAllTeeMachineOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.AllowAllTeeMachineOwners(&_OwnerAllowlist.TransactOpts, _extensionId)
}

// AllowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x61f957e7.
//
// Solidity: function allowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactor) AllowAllTeeWalletProjectOwners(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _OwnerAllowlist.contract.Transact(opts, "allowAllTeeWalletProjectOwners", _extensionId)
}

// AllowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x61f957e7.
//
// Solidity: function allowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_OwnerAllowlist *OwnerAllowlistSession) AllowAllTeeWalletProjectOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.AllowAllTeeWalletProjectOwners(&_OwnerAllowlist.TransactOpts, _extensionId)
}

// AllowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x61f957e7.
//
// Solidity: function allowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactorSession) AllowAllTeeWalletProjectOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.AllowAllTeeWalletProjectOwners(&_OwnerAllowlist.TransactOpts, _extensionId)
}

// DisallowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0xb86936ba.
//
// Solidity: function disallowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactor) DisallowAllTeeMachineOwners(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _OwnerAllowlist.contract.Transact(opts, "disallowAllTeeMachineOwners", _extensionId)
}

// DisallowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0xb86936ba.
//
// Solidity: function disallowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_OwnerAllowlist *OwnerAllowlistSession) DisallowAllTeeMachineOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.DisallowAllTeeMachineOwners(&_OwnerAllowlist.TransactOpts, _extensionId)
}

// DisallowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0xb86936ba.
//
// Solidity: function disallowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactorSession) DisallowAllTeeMachineOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.DisallowAllTeeMachineOwners(&_OwnerAllowlist.TransactOpts, _extensionId)
}

// DisallowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x919fd33a.
//
// Solidity: function disallowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactor) DisallowAllTeeWalletProjectOwners(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _OwnerAllowlist.contract.Transact(opts, "disallowAllTeeWalletProjectOwners", _extensionId)
}

// DisallowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x919fd33a.
//
// Solidity: function disallowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_OwnerAllowlist *OwnerAllowlistSession) DisallowAllTeeWalletProjectOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.DisallowAllTeeWalletProjectOwners(&_OwnerAllowlist.TransactOpts, _extensionId)
}

// DisallowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x919fd33a.
//
// Solidity: function disallowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactorSession) DisallowAllTeeWalletProjectOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.DisallowAllTeeWalletProjectOwners(&_OwnerAllowlist.TransactOpts, _extensionId)
}

// RemoveAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0x5d371c3c.
//
// Solidity: function removeAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactor) RemoveAllowedTeeMachineOwners(opts *bind.TransactOpts, _extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _OwnerAllowlist.contract.Transact(opts, "removeAllowedTeeMachineOwners", _extensionId, _owners)
}

// RemoveAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0x5d371c3c.
//
// Solidity: function removeAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_OwnerAllowlist *OwnerAllowlistSession) RemoveAllowedTeeMachineOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.RemoveAllowedTeeMachineOwners(&_OwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// RemoveAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0x5d371c3c.
//
// Solidity: function removeAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactorSession) RemoveAllowedTeeMachineOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.RemoveAllowedTeeMachineOwners(&_OwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// RemoveAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x62a2513d.
//
// Solidity: function removeAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactor) RemoveAllowedTeeWalletProjectOwners(opts *bind.TransactOpts, _extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _OwnerAllowlist.contract.Transact(opts, "removeAllowedTeeWalletProjectOwners", _extensionId, _owners)
}

// RemoveAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x62a2513d.
//
// Solidity: function removeAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_OwnerAllowlist *OwnerAllowlistSession) RemoveAllowedTeeWalletProjectOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.RemoveAllowedTeeWalletProjectOwners(&_OwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// RemoveAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x62a2513d.
//
// Solidity: function removeAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_OwnerAllowlist *OwnerAllowlistTransactorSession) RemoveAllowedTeeWalletProjectOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _OwnerAllowlist.Contract.RemoveAllowedTeeWalletProjectOwners(&_OwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// OwnerAllowlistAllTeeMachineOwnersAllowedIterator is returned from FilterAllTeeMachineOwnersAllowed and is used to iterate over the raw logs and unpacked data for AllTeeMachineOwnersAllowed events raised by the OwnerAllowlist contract.
type OwnerAllowlistAllTeeMachineOwnersAllowedIterator struct {
	Event *OwnerAllowlistAllTeeMachineOwnersAllowed // Event containing the contract specifics and raw log

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
func (it *OwnerAllowlistAllTeeMachineOwnersAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerAllowlistAllTeeMachineOwnersAllowed)
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
		it.Event = new(OwnerAllowlistAllTeeMachineOwnersAllowed)
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
func (it *OwnerAllowlistAllTeeMachineOwnersAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerAllowlistAllTeeMachineOwnersAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerAllowlistAllTeeMachineOwnersAllowed represents a AllTeeMachineOwnersAllowed event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllTeeMachineOwnersAllowed struct {
	ExtensionId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllTeeMachineOwnersAllowed is a free log retrieval operation binding the contract event 0x1bc41c6dd61d3a656898250dfc58a14ddd4735f4480ae4c1581072ad19fcd8a5.
//
// Solidity: event AllTeeMachineOwnersAllowed(uint256 extensionId)
func (_OwnerAllowlist *OwnerAllowlistFilterer) FilterAllTeeMachineOwnersAllowed(opts *bind.FilterOpts) (*OwnerAllowlistAllTeeMachineOwnersAllowedIterator, error) {

	logs, sub, err := _OwnerAllowlist.contract.FilterLogs(opts, "AllTeeMachineOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return &OwnerAllowlistAllTeeMachineOwnersAllowedIterator{contract: _OwnerAllowlist.contract, event: "AllTeeMachineOwnersAllowed", logs: logs, sub: sub}, nil
}

// WatchAllTeeMachineOwnersAllowed is a free log subscription operation binding the contract event 0x1bc41c6dd61d3a656898250dfc58a14ddd4735f4480ae4c1581072ad19fcd8a5.
//
// Solidity: event AllTeeMachineOwnersAllowed(uint256 extensionId)
func (_OwnerAllowlist *OwnerAllowlistFilterer) WatchAllTeeMachineOwnersAllowed(opts *bind.WatchOpts, sink chan<- *OwnerAllowlistAllTeeMachineOwnersAllowed) (event.Subscription, error) {

	logs, sub, err := _OwnerAllowlist.contract.WatchLogs(opts, "AllTeeMachineOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerAllowlistAllTeeMachineOwnersAllowed)
				if err := _OwnerAllowlist.contract.UnpackLog(event, "AllTeeMachineOwnersAllowed", log); err != nil {
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

// ParseAllTeeMachineOwnersAllowed is a log parse operation binding the contract event 0x1bc41c6dd61d3a656898250dfc58a14ddd4735f4480ae4c1581072ad19fcd8a5.
//
// Solidity: event AllTeeMachineOwnersAllowed(uint256 extensionId)
func (_OwnerAllowlist *OwnerAllowlistFilterer) ParseAllTeeMachineOwnersAllowed(log types.Log) (*OwnerAllowlistAllTeeMachineOwnersAllowed, error) {
	event := new(OwnerAllowlistAllTeeMachineOwnersAllowed)
	if err := _OwnerAllowlist.contract.UnpackLog(event, "AllTeeMachineOwnersAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerAllowlistAllTeeMachineOwnersDisallowedIterator is returned from FilterAllTeeMachineOwnersDisallowed and is used to iterate over the raw logs and unpacked data for AllTeeMachineOwnersDisallowed events raised by the OwnerAllowlist contract.
type OwnerAllowlistAllTeeMachineOwnersDisallowedIterator struct {
	Event *OwnerAllowlistAllTeeMachineOwnersDisallowed // Event containing the contract specifics and raw log

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
func (it *OwnerAllowlistAllTeeMachineOwnersDisallowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerAllowlistAllTeeMachineOwnersDisallowed)
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
		it.Event = new(OwnerAllowlistAllTeeMachineOwnersDisallowed)
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
func (it *OwnerAllowlistAllTeeMachineOwnersDisallowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerAllowlistAllTeeMachineOwnersDisallowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerAllowlistAllTeeMachineOwnersDisallowed represents a AllTeeMachineOwnersDisallowed event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllTeeMachineOwnersDisallowed struct {
	ExtensionId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllTeeMachineOwnersDisallowed is a free log retrieval operation binding the contract event 0x76fc91f6aa03fa961257e10347e87c3844fd0547a18ab3967b5a52db80c353d3.
//
// Solidity: event AllTeeMachineOwnersDisallowed(uint256 extensionId)
func (_OwnerAllowlist *OwnerAllowlistFilterer) FilterAllTeeMachineOwnersDisallowed(opts *bind.FilterOpts) (*OwnerAllowlistAllTeeMachineOwnersDisallowedIterator, error) {

	logs, sub, err := _OwnerAllowlist.contract.FilterLogs(opts, "AllTeeMachineOwnersDisallowed")
	if err != nil {
		return nil, err
	}
	return &OwnerAllowlistAllTeeMachineOwnersDisallowedIterator{contract: _OwnerAllowlist.contract, event: "AllTeeMachineOwnersDisallowed", logs: logs, sub: sub}, nil
}

// WatchAllTeeMachineOwnersDisallowed is a free log subscription operation binding the contract event 0x76fc91f6aa03fa961257e10347e87c3844fd0547a18ab3967b5a52db80c353d3.
//
// Solidity: event AllTeeMachineOwnersDisallowed(uint256 extensionId)
func (_OwnerAllowlist *OwnerAllowlistFilterer) WatchAllTeeMachineOwnersDisallowed(opts *bind.WatchOpts, sink chan<- *OwnerAllowlistAllTeeMachineOwnersDisallowed) (event.Subscription, error) {

	logs, sub, err := _OwnerAllowlist.contract.WatchLogs(opts, "AllTeeMachineOwnersDisallowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerAllowlistAllTeeMachineOwnersDisallowed)
				if err := _OwnerAllowlist.contract.UnpackLog(event, "AllTeeMachineOwnersDisallowed", log); err != nil {
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

// ParseAllTeeMachineOwnersDisallowed is a log parse operation binding the contract event 0x76fc91f6aa03fa961257e10347e87c3844fd0547a18ab3967b5a52db80c353d3.
//
// Solidity: event AllTeeMachineOwnersDisallowed(uint256 extensionId)
func (_OwnerAllowlist *OwnerAllowlistFilterer) ParseAllTeeMachineOwnersDisallowed(log types.Log) (*OwnerAllowlistAllTeeMachineOwnersDisallowed, error) {
	event := new(OwnerAllowlistAllTeeMachineOwnersDisallowed)
	if err := _OwnerAllowlist.contract.UnpackLog(event, "AllTeeMachineOwnersDisallowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator is returned from FilterAllTeeWalletProjectOwnersAllowed and is used to iterate over the raw logs and unpacked data for AllTeeWalletProjectOwnersAllowed events raised by the OwnerAllowlist contract.
type OwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator struct {
	Event *OwnerAllowlistAllTeeWalletProjectOwnersAllowed // Event containing the contract specifics and raw log

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
func (it *OwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerAllowlistAllTeeWalletProjectOwnersAllowed)
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
		it.Event = new(OwnerAllowlistAllTeeWalletProjectOwnersAllowed)
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
func (it *OwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerAllowlistAllTeeWalletProjectOwnersAllowed represents a AllTeeWalletProjectOwnersAllowed event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllTeeWalletProjectOwnersAllowed struct {
	ExtensionId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllTeeWalletProjectOwnersAllowed is a free log retrieval operation binding the contract event 0x1936686ee6942840bf2283355eb720261b36284a52b74604253556ee61c08854.
//
// Solidity: event AllTeeWalletProjectOwnersAllowed(uint256 extensionId)
func (_OwnerAllowlist *OwnerAllowlistFilterer) FilterAllTeeWalletProjectOwnersAllowed(opts *bind.FilterOpts) (*OwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator, error) {

	logs, sub, err := _OwnerAllowlist.contract.FilterLogs(opts, "AllTeeWalletProjectOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return &OwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator{contract: _OwnerAllowlist.contract, event: "AllTeeWalletProjectOwnersAllowed", logs: logs, sub: sub}, nil
}

// WatchAllTeeWalletProjectOwnersAllowed is a free log subscription operation binding the contract event 0x1936686ee6942840bf2283355eb720261b36284a52b74604253556ee61c08854.
//
// Solidity: event AllTeeWalletProjectOwnersAllowed(uint256 extensionId)
func (_OwnerAllowlist *OwnerAllowlistFilterer) WatchAllTeeWalletProjectOwnersAllowed(opts *bind.WatchOpts, sink chan<- *OwnerAllowlistAllTeeWalletProjectOwnersAllowed) (event.Subscription, error) {

	logs, sub, err := _OwnerAllowlist.contract.WatchLogs(opts, "AllTeeWalletProjectOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerAllowlistAllTeeWalletProjectOwnersAllowed)
				if err := _OwnerAllowlist.contract.UnpackLog(event, "AllTeeWalletProjectOwnersAllowed", log); err != nil {
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

// ParseAllTeeWalletProjectOwnersAllowed is a log parse operation binding the contract event 0x1936686ee6942840bf2283355eb720261b36284a52b74604253556ee61c08854.
//
// Solidity: event AllTeeWalletProjectOwnersAllowed(uint256 extensionId)
func (_OwnerAllowlist *OwnerAllowlistFilterer) ParseAllTeeWalletProjectOwnersAllowed(log types.Log) (*OwnerAllowlistAllTeeWalletProjectOwnersAllowed, error) {
	event := new(OwnerAllowlistAllTeeWalletProjectOwnersAllowed)
	if err := _OwnerAllowlist.contract.UnpackLog(event, "AllTeeWalletProjectOwnersAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator is returned from FilterAllTeeWalletProjectOwnersDisallowed and is used to iterate over the raw logs and unpacked data for AllTeeWalletProjectOwnersDisallowed events raised by the OwnerAllowlist contract.
type OwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator struct {
	Event *OwnerAllowlistAllTeeWalletProjectOwnersDisallowed // Event containing the contract specifics and raw log

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
func (it *OwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerAllowlistAllTeeWalletProjectOwnersDisallowed)
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
		it.Event = new(OwnerAllowlistAllTeeWalletProjectOwnersDisallowed)
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
func (it *OwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerAllowlistAllTeeWalletProjectOwnersDisallowed represents a AllTeeWalletProjectOwnersDisallowed event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllTeeWalletProjectOwnersDisallowed struct {
	ExtensionId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllTeeWalletProjectOwnersDisallowed is a free log retrieval operation binding the contract event 0x1c6abd0ba2324be35a8b8dac7bef3bf25c5d2d2cc45c216ee3b38a82da6948e8.
//
// Solidity: event AllTeeWalletProjectOwnersDisallowed(uint256 extensionId)
func (_OwnerAllowlist *OwnerAllowlistFilterer) FilterAllTeeWalletProjectOwnersDisallowed(opts *bind.FilterOpts) (*OwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator, error) {

	logs, sub, err := _OwnerAllowlist.contract.FilterLogs(opts, "AllTeeWalletProjectOwnersDisallowed")
	if err != nil {
		return nil, err
	}
	return &OwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator{contract: _OwnerAllowlist.contract, event: "AllTeeWalletProjectOwnersDisallowed", logs: logs, sub: sub}, nil
}

// WatchAllTeeWalletProjectOwnersDisallowed is a free log subscription operation binding the contract event 0x1c6abd0ba2324be35a8b8dac7bef3bf25c5d2d2cc45c216ee3b38a82da6948e8.
//
// Solidity: event AllTeeWalletProjectOwnersDisallowed(uint256 extensionId)
func (_OwnerAllowlist *OwnerAllowlistFilterer) WatchAllTeeWalletProjectOwnersDisallowed(opts *bind.WatchOpts, sink chan<- *OwnerAllowlistAllTeeWalletProjectOwnersDisallowed) (event.Subscription, error) {

	logs, sub, err := _OwnerAllowlist.contract.WatchLogs(opts, "AllTeeWalletProjectOwnersDisallowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerAllowlistAllTeeWalletProjectOwnersDisallowed)
				if err := _OwnerAllowlist.contract.UnpackLog(event, "AllTeeWalletProjectOwnersDisallowed", log); err != nil {
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

// ParseAllTeeWalletProjectOwnersDisallowed is a log parse operation binding the contract event 0x1c6abd0ba2324be35a8b8dac7bef3bf25c5d2d2cc45c216ee3b38a82da6948e8.
//
// Solidity: event AllTeeWalletProjectOwnersDisallowed(uint256 extensionId)
func (_OwnerAllowlist *OwnerAllowlistFilterer) ParseAllTeeWalletProjectOwnersDisallowed(log types.Log) (*OwnerAllowlistAllTeeWalletProjectOwnersDisallowed, error) {
	event := new(OwnerAllowlistAllTeeWalletProjectOwnersDisallowed)
	if err := _OwnerAllowlist.contract.UnpackLog(event, "AllTeeWalletProjectOwnersDisallowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerAllowlistAllowedTeeMachineOwnersAddedIterator is returned from FilterAllowedTeeMachineOwnersAdded and is used to iterate over the raw logs and unpacked data for AllowedTeeMachineOwnersAdded events raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedTeeMachineOwnersAddedIterator struct {
	Event *OwnerAllowlistAllowedTeeMachineOwnersAdded // Event containing the contract specifics and raw log

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
func (it *OwnerAllowlistAllowedTeeMachineOwnersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerAllowlistAllowedTeeMachineOwnersAdded)
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
		it.Event = new(OwnerAllowlistAllowedTeeMachineOwnersAdded)
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
func (it *OwnerAllowlistAllowedTeeMachineOwnersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerAllowlistAllowedTeeMachineOwnersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerAllowlistAllowedTeeMachineOwnersAdded represents a AllowedTeeMachineOwnersAdded event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedTeeMachineOwnersAdded struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllowedTeeMachineOwnersAdded is a free log retrieval operation binding the contract event 0xd392883c2acc8ffc90b60bb8783db156d990502913d5b159e209cd604475337c.
//
// Solidity: event AllowedTeeMachineOwnersAdded(uint256 extensionId, address[] owners)
func (_OwnerAllowlist *OwnerAllowlistFilterer) FilterAllowedTeeMachineOwnersAdded(opts *bind.FilterOpts) (*OwnerAllowlistAllowedTeeMachineOwnersAddedIterator, error) {

	logs, sub, err := _OwnerAllowlist.contract.FilterLogs(opts, "AllowedTeeMachineOwnersAdded")
	if err != nil {
		return nil, err
	}
	return &OwnerAllowlistAllowedTeeMachineOwnersAddedIterator{contract: _OwnerAllowlist.contract, event: "AllowedTeeMachineOwnersAdded", logs: logs, sub: sub}, nil
}

// WatchAllowedTeeMachineOwnersAdded is a free log subscription operation binding the contract event 0xd392883c2acc8ffc90b60bb8783db156d990502913d5b159e209cd604475337c.
//
// Solidity: event AllowedTeeMachineOwnersAdded(uint256 extensionId, address[] owners)
func (_OwnerAllowlist *OwnerAllowlistFilterer) WatchAllowedTeeMachineOwnersAdded(opts *bind.WatchOpts, sink chan<- *OwnerAllowlistAllowedTeeMachineOwnersAdded) (event.Subscription, error) {

	logs, sub, err := _OwnerAllowlist.contract.WatchLogs(opts, "AllowedTeeMachineOwnersAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerAllowlistAllowedTeeMachineOwnersAdded)
				if err := _OwnerAllowlist.contract.UnpackLog(event, "AllowedTeeMachineOwnersAdded", log); err != nil {
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

// ParseAllowedTeeMachineOwnersAdded is a log parse operation binding the contract event 0xd392883c2acc8ffc90b60bb8783db156d990502913d5b159e209cd604475337c.
//
// Solidity: event AllowedTeeMachineOwnersAdded(uint256 extensionId, address[] owners)
func (_OwnerAllowlist *OwnerAllowlistFilterer) ParseAllowedTeeMachineOwnersAdded(log types.Log) (*OwnerAllowlistAllowedTeeMachineOwnersAdded, error) {
	event := new(OwnerAllowlistAllowedTeeMachineOwnersAdded)
	if err := _OwnerAllowlist.contract.UnpackLog(event, "AllowedTeeMachineOwnersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerAllowlistAllowedTeeMachineOwnersRemovedIterator is returned from FilterAllowedTeeMachineOwnersRemoved and is used to iterate over the raw logs and unpacked data for AllowedTeeMachineOwnersRemoved events raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedTeeMachineOwnersRemovedIterator struct {
	Event *OwnerAllowlistAllowedTeeMachineOwnersRemoved // Event containing the contract specifics and raw log

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
func (it *OwnerAllowlistAllowedTeeMachineOwnersRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerAllowlistAllowedTeeMachineOwnersRemoved)
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
		it.Event = new(OwnerAllowlistAllowedTeeMachineOwnersRemoved)
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
func (it *OwnerAllowlistAllowedTeeMachineOwnersRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerAllowlistAllowedTeeMachineOwnersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerAllowlistAllowedTeeMachineOwnersRemoved represents a AllowedTeeMachineOwnersRemoved event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedTeeMachineOwnersRemoved struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllowedTeeMachineOwnersRemoved is a free log retrieval operation binding the contract event 0x8444f01824cfebbd6772bffe7855e88270c338f6879216e756f86cd54c53ddf6.
//
// Solidity: event AllowedTeeMachineOwnersRemoved(uint256 extensionId, address[] owners)
func (_OwnerAllowlist *OwnerAllowlistFilterer) FilterAllowedTeeMachineOwnersRemoved(opts *bind.FilterOpts) (*OwnerAllowlistAllowedTeeMachineOwnersRemovedIterator, error) {

	logs, sub, err := _OwnerAllowlist.contract.FilterLogs(opts, "AllowedTeeMachineOwnersRemoved")
	if err != nil {
		return nil, err
	}
	return &OwnerAllowlistAllowedTeeMachineOwnersRemovedIterator{contract: _OwnerAllowlist.contract, event: "AllowedTeeMachineOwnersRemoved", logs: logs, sub: sub}, nil
}

// WatchAllowedTeeMachineOwnersRemoved is a free log subscription operation binding the contract event 0x8444f01824cfebbd6772bffe7855e88270c338f6879216e756f86cd54c53ddf6.
//
// Solidity: event AllowedTeeMachineOwnersRemoved(uint256 extensionId, address[] owners)
func (_OwnerAllowlist *OwnerAllowlistFilterer) WatchAllowedTeeMachineOwnersRemoved(opts *bind.WatchOpts, sink chan<- *OwnerAllowlistAllowedTeeMachineOwnersRemoved) (event.Subscription, error) {

	logs, sub, err := _OwnerAllowlist.contract.WatchLogs(opts, "AllowedTeeMachineOwnersRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerAllowlistAllowedTeeMachineOwnersRemoved)
				if err := _OwnerAllowlist.contract.UnpackLog(event, "AllowedTeeMachineOwnersRemoved", log); err != nil {
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

// ParseAllowedTeeMachineOwnersRemoved is a log parse operation binding the contract event 0x8444f01824cfebbd6772bffe7855e88270c338f6879216e756f86cd54c53ddf6.
//
// Solidity: event AllowedTeeMachineOwnersRemoved(uint256 extensionId, address[] owners)
func (_OwnerAllowlist *OwnerAllowlistFilterer) ParseAllowedTeeMachineOwnersRemoved(log types.Log) (*OwnerAllowlistAllowedTeeMachineOwnersRemoved, error) {
	event := new(OwnerAllowlistAllowedTeeMachineOwnersRemoved)
	if err := _OwnerAllowlist.contract.UnpackLog(event, "AllowedTeeMachineOwnersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator is returned from FilterAllowedTeeWalletProjectOwnersAdded and is used to iterate over the raw logs and unpacked data for AllowedTeeWalletProjectOwnersAdded events raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator struct {
	Event *OwnerAllowlistAllowedTeeWalletProjectOwnersAdded // Event containing the contract specifics and raw log

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
func (it *OwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerAllowlistAllowedTeeWalletProjectOwnersAdded)
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
		it.Event = new(OwnerAllowlistAllowedTeeWalletProjectOwnersAdded)
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
func (it *OwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerAllowlistAllowedTeeWalletProjectOwnersAdded represents a AllowedTeeWalletProjectOwnersAdded event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedTeeWalletProjectOwnersAdded struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllowedTeeWalletProjectOwnersAdded is a free log retrieval operation binding the contract event 0x60340c199493a9a48f50ae5af13ebefed3adeeb1c4f0a5476d59654680ffdb0d.
//
// Solidity: event AllowedTeeWalletProjectOwnersAdded(uint256 extensionId, address[] owners)
func (_OwnerAllowlist *OwnerAllowlistFilterer) FilterAllowedTeeWalletProjectOwnersAdded(opts *bind.FilterOpts) (*OwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator, error) {

	logs, sub, err := _OwnerAllowlist.contract.FilterLogs(opts, "AllowedTeeWalletProjectOwnersAdded")
	if err != nil {
		return nil, err
	}
	return &OwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator{contract: _OwnerAllowlist.contract, event: "AllowedTeeWalletProjectOwnersAdded", logs: logs, sub: sub}, nil
}

// WatchAllowedTeeWalletProjectOwnersAdded is a free log subscription operation binding the contract event 0x60340c199493a9a48f50ae5af13ebefed3adeeb1c4f0a5476d59654680ffdb0d.
//
// Solidity: event AllowedTeeWalletProjectOwnersAdded(uint256 extensionId, address[] owners)
func (_OwnerAllowlist *OwnerAllowlistFilterer) WatchAllowedTeeWalletProjectOwnersAdded(opts *bind.WatchOpts, sink chan<- *OwnerAllowlistAllowedTeeWalletProjectOwnersAdded) (event.Subscription, error) {

	logs, sub, err := _OwnerAllowlist.contract.WatchLogs(opts, "AllowedTeeWalletProjectOwnersAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerAllowlistAllowedTeeWalletProjectOwnersAdded)
				if err := _OwnerAllowlist.contract.UnpackLog(event, "AllowedTeeWalletProjectOwnersAdded", log); err != nil {
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

// ParseAllowedTeeWalletProjectOwnersAdded is a log parse operation binding the contract event 0x60340c199493a9a48f50ae5af13ebefed3adeeb1c4f0a5476d59654680ffdb0d.
//
// Solidity: event AllowedTeeWalletProjectOwnersAdded(uint256 extensionId, address[] owners)
func (_OwnerAllowlist *OwnerAllowlistFilterer) ParseAllowedTeeWalletProjectOwnersAdded(log types.Log) (*OwnerAllowlistAllowedTeeWalletProjectOwnersAdded, error) {
	event := new(OwnerAllowlistAllowedTeeWalletProjectOwnersAdded)
	if err := _OwnerAllowlist.contract.UnpackLog(event, "AllowedTeeWalletProjectOwnersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator is returned from FilterAllowedTeeWalletProjectOwnersRemoved and is used to iterate over the raw logs and unpacked data for AllowedTeeWalletProjectOwnersRemoved events raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator struct {
	Event *OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved // Event containing the contract specifics and raw log

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
func (it *OwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved)
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
		it.Event = new(OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved)
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
func (it *OwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved represents a AllowedTeeWalletProjectOwnersRemoved event raised by the OwnerAllowlist contract.
type OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllowedTeeWalletProjectOwnersRemoved is a free log retrieval operation binding the contract event 0x676aa060014299ae54ae19d4e1fd6fe0b4681aad07cc8ddf7979673262e3fcac.
//
// Solidity: event AllowedTeeWalletProjectOwnersRemoved(uint256 extensionId, address[] owners)
func (_OwnerAllowlist *OwnerAllowlistFilterer) FilterAllowedTeeWalletProjectOwnersRemoved(opts *bind.FilterOpts) (*OwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator, error) {

	logs, sub, err := _OwnerAllowlist.contract.FilterLogs(opts, "AllowedTeeWalletProjectOwnersRemoved")
	if err != nil {
		return nil, err
	}
	return &OwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator{contract: _OwnerAllowlist.contract, event: "AllowedTeeWalletProjectOwnersRemoved", logs: logs, sub: sub}, nil
}

// WatchAllowedTeeWalletProjectOwnersRemoved is a free log subscription operation binding the contract event 0x676aa060014299ae54ae19d4e1fd6fe0b4681aad07cc8ddf7979673262e3fcac.
//
// Solidity: event AllowedTeeWalletProjectOwnersRemoved(uint256 extensionId, address[] owners)
func (_OwnerAllowlist *OwnerAllowlistFilterer) WatchAllowedTeeWalletProjectOwnersRemoved(opts *bind.WatchOpts, sink chan<- *OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved) (event.Subscription, error) {

	logs, sub, err := _OwnerAllowlist.contract.WatchLogs(opts, "AllowedTeeWalletProjectOwnersRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved)
				if err := _OwnerAllowlist.contract.UnpackLog(event, "AllowedTeeWalletProjectOwnersRemoved", log); err != nil {
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

// ParseAllowedTeeWalletProjectOwnersRemoved is a log parse operation binding the contract event 0x676aa060014299ae54ae19d4e1fd6fe0b4681aad07cc8ddf7979673262e3fcac.
//
// Solidity: event AllowedTeeWalletProjectOwnersRemoved(uint256 extensionId, address[] owners)
func (_OwnerAllowlist *OwnerAllowlistFilterer) ParseAllowedTeeWalletProjectOwnersRemoved(log types.Log) (*OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved, error) {
	event := new(OwnerAllowlistAllowedTeeWalletProjectOwnersRemoved)
	if err := _OwnerAllowlist.contract.UnpackLog(event, "AllowedTeeWalletProjectOwnersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
