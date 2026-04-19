// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teeownerallowlist

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

// TeeOwnerAllowlistMetaData contains all meta data concerning the TeeOwnerAllowlist contract.
var TeeOwnerAllowlistMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAlreadyAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerNotInAllowlist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeMachineOwnersAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeMachineOwnersDisallowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeWalletProjectOwnersAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"AllTeeWalletProjectOwnersDisallowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeMachineOwnersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeMachineOwnersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeWalletProjectOwnersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"}],\"name\":\"AllowedTeeWalletProjectOwnersRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"addAllowedTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"addAllowedTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allTeeMachineOwnersAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_allAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allTeeWalletProjectOwnersAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_allAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allowAllTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"allowAllTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"disallowAllTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"disallowAllTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getAllowedTeeMachineOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_allowedOwners\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getAllowedTeeWalletProjectOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_allowedOwners\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"isAllowedTeeMachineOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"isAllowedTeeWalletProjectOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isAllowed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"removeAllowedTeeMachineOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"}],\"name\":\"removeAllowedTeeWalletProjectOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeOwnerAllowlistABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeOwnerAllowlistMetaData.ABI instead.
var TeeOwnerAllowlistABI = TeeOwnerAllowlistMetaData.ABI

// TeeOwnerAllowlist is an auto generated Go binding around an Ethereum contract.
type TeeOwnerAllowlist struct {
	TeeOwnerAllowlistCaller     // Read-only binding to the contract
	TeeOwnerAllowlistTransactor // Write-only binding to the contract
	TeeOwnerAllowlistFilterer   // Log filterer for contract events
}

// TeeOwnerAllowlistCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeOwnerAllowlistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeOwnerAllowlistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeOwnerAllowlistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeOwnerAllowlistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeOwnerAllowlistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeOwnerAllowlistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeOwnerAllowlistSession struct {
	Contract     *TeeOwnerAllowlist // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TeeOwnerAllowlistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeOwnerAllowlistCallerSession struct {
	Contract *TeeOwnerAllowlistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// TeeOwnerAllowlistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeOwnerAllowlistTransactorSession struct {
	Contract     *TeeOwnerAllowlistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// TeeOwnerAllowlistRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeOwnerAllowlistRaw struct {
	Contract *TeeOwnerAllowlist // Generic contract binding to access the raw methods on
}

// TeeOwnerAllowlistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeOwnerAllowlistCallerRaw struct {
	Contract *TeeOwnerAllowlistCaller // Generic read-only contract binding to access the raw methods on
}

// TeeOwnerAllowlistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeOwnerAllowlistTransactorRaw struct {
	Contract *TeeOwnerAllowlistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeOwnerAllowlist creates a new instance of TeeOwnerAllowlist, bound to a specific deployed contract.
func NewTeeOwnerAllowlist(address common.Address, backend bind.ContractBackend) (*TeeOwnerAllowlist, error) {
	contract, err := bindTeeOwnerAllowlist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlist{TeeOwnerAllowlistCaller: TeeOwnerAllowlistCaller{contract: contract}, TeeOwnerAllowlistTransactor: TeeOwnerAllowlistTransactor{contract: contract}, TeeOwnerAllowlistFilterer: TeeOwnerAllowlistFilterer{contract: contract}}, nil
}

// NewTeeOwnerAllowlistCaller creates a new read-only instance of TeeOwnerAllowlist, bound to a specific deployed contract.
func NewTeeOwnerAllowlistCaller(address common.Address, caller bind.ContractCaller) (*TeeOwnerAllowlistCaller, error) {
	contract, err := bindTeeOwnerAllowlist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistCaller{contract: contract}, nil
}

// NewTeeOwnerAllowlistTransactor creates a new write-only instance of TeeOwnerAllowlist, bound to a specific deployed contract.
func NewTeeOwnerAllowlistTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeOwnerAllowlistTransactor, error) {
	contract, err := bindTeeOwnerAllowlist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistTransactor{contract: contract}, nil
}

// NewTeeOwnerAllowlistFilterer creates a new log filterer instance of TeeOwnerAllowlist, bound to a specific deployed contract.
func NewTeeOwnerAllowlistFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeOwnerAllowlistFilterer, error) {
	contract, err := bindTeeOwnerAllowlist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistFilterer{contract: contract}, nil
}

// bindTeeOwnerAllowlist binds a generic wrapper to an already deployed contract.
func bindTeeOwnerAllowlist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeOwnerAllowlistMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeOwnerAllowlist *TeeOwnerAllowlistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeOwnerAllowlist.Contract.TeeOwnerAllowlistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeOwnerAllowlist *TeeOwnerAllowlistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.TeeOwnerAllowlistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeOwnerAllowlist *TeeOwnerAllowlistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.TeeOwnerAllowlistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeOwnerAllowlist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.contract.Transact(opts, method, params...)
}

// AllTeeMachineOwnersAllowed is a free data retrieval call binding the contract method 0x2a30b9f4.
//
// Solidity: function allTeeMachineOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) AllTeeMachineOwnersAllowed(opts *bind.CallOpts, _extensionId *big.Int) (bool, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "allTeeMachineOwnersAllowed", _extensionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllTeeMachineOwnersAllowed is a free data retrieval call binding the contract method 0x2a30b9f4.
//
// Solidity: function allTeeMachineOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) AllTeeMachineOwnersAllowed(_extensionId *big.Int) (bool, error) {
	return _TeeOwnerAllowlist.Contract.AllTeeMachineOwnersAllowed(&_TeeOwnerAllowlist.CallOpts, _extensionId)
}

// AllTeeMachineOwnersAllowed is a free data retrieval call binding the contract method 0x2a30b9f4.
//
// Solidity: function allTeeMachineOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) AllTeeMachineOwnersAllowed(_extensionId *big.Int) (bool, error) {
	return _TeeOwnerAllowlist.Contract.AllTeeMachineOwnersAllowed(&_TeeOwnerAllowlist.CallOpts, _extensionId)
}

// AllTeeWalletProjectOwnersAllowed is a free data retrieval call binding the contract method 0x8398fc2e.
//
// Solidity: function allTeeWalletProjectOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) AllTeeWalletProjectOwnersAllowed(opts *bind.CallOpts, _extensionId *big.Int) (bool, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "allTeeWalletProjectOwnersAllowed", _extensionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllTeeWalletProjectOwnersAllowed is a free data retrieval call binding the contract method 0x8398fc2e.
//
// Solidity: function allTeeWalletProjectOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) AllTeeWalletProjectOwnersAllowed(_extensionId *big.Int) (bool, error) {
	return _TeeOwnerAllowlist.Contract.AllTeeWalletProjectOwnersAllowed(&_TeeOwnerAllowlist.CallOpts, _extensionId)
}

// AllTeeWalletProjectOwnersAllowed is a free data retrieval call binding the contract method 0x8398fc2e.
//
// Solidity: function allTeeWalletProjectOwnersAllowed(uint256 _extensionId) view returns(bool _allAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) AllTeeWalletProjectOwnersAllowed(_extensionId *big.Int) (bool, error) {
	return _TeeOwnerAllowlist.Contract.AllTeeWalletProjectOwnersAllowed(&_TeeOwnerAllowlist.CallOpts, _extensionId)
}

// GetAllowedTeeMachineOwners is a free data retrieval call binding the contract method 0x5f37ea30.
//
// Solidity: function getAllowedTeeMachineOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) GetAllowedTeeMachineOwners(opts *bind.CallOpts, _extensionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "getAllowedTeeMachineOwners", _extensionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowedTeeMachineOwners is a free data retrieval call binding the contract method 0x5f37ea30.
//
// Solidity: function getAllowedTeeMachineOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) GetAllowedTeeMachineOwners(_extensionId *big.Int) ([]common.Address, error) {
	return _TeeOwnerAllowlist.Contract.GetAllowedTeeMachineOwners(&_TeeOwnerAllowlist.CallOpts, _extensionId)
}

// GetAllowedTeeMachineOwners is a free data retrieval call binding the contract method 0x5f37ea30.
//
// Solidity: function getAllowedTeeMachineOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) GetAllowedTeeMachineOwners(_extensionId *big.Int) ([]common.Address, error) {
	return _TeeOwnerAllowlist.Contract.GetAllowedTeeMachineOwners(&_TeeOwnerAllowlist.CallOpts, _extensionId)
}

// GetAllowedTeeWalletProjectOwners is a free data retrieval call binding the contract method 0xd11340ed.
//
// Solidity: function getAllowedTeeWalletProjectOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) GetAllowedTeeWalletProjectOwners(opts *bind.CallOpts, _extensionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "getAllowedTeeWalletProjectOwners", _extensionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllowedTeeWalletProjectOwners is a free data retrieval call binding the contract method 0xd11340ed.
//
// Solidity: function getAllowedTeeWalletProjectOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) GetAllowedTeeWalletProjectOwners(_extensionId *big.Int) ([]common.Address, error) {
	return _TeeOwnerAllowlist.Contract.GetAllowedTeeWalletProjectOwners(&_TeeOwnerAllowlist.CallOpts, _extensionId)
}

// GetAllowedTeeWalletProjectOwners is a free data retrieval call binding the contract method 0xd11340ed.
//
// Solidity: function getAllowedTeeWalletProjectOwners(uint256 _extensionId) view returns(address[] _allowedOwners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) GetAllowedTeeWalletProjectOwners(_extensionId *big.Int) ([]common.Address, error) {
	return _TeeOwnerAllowlist.Contract.GetAllowedTeeWalletProjectOwners(&_TeeOwnerAllowlist.CallOpts, _extensionId)
}

// IsAllowedTeeMachineOwner is a free data retrieval call binding the contract method 0x84f5da36.
//
// Solidity: function isAllowedTeeMachineOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) IsAllowedTeeMachineOwner(opts *bind.CallOpts, _extensionId *big.Int, _owner common.Address) (bool, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "isAllowedTeeMachineOwner", _extensionId, _owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedTeeMachineOwner is a free data retrieval call binding the contract method 0x84f5da36.
//
// Solidity: function isAllowedTeeMachineOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) IsAllowedTeeMachineOwner(_extensionId *big.Int, _owner common.Address) (bool, error) {
	return _TeeOwnerAllowlist.Contract.IsAllowedTeeMachineOwner(&_TeeOwnerAllowlist.CallOpts, _extensionId, _owner)
}

// IsAllowedTeeMachineOwner is a free data retrieval call binding the contract method 0x84f5da36.
//
// Solidity: function isAllowedTeeMachineOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) IsAllowedTeeMachineOwner(_extensionId *big.Int, _owner common.Address) (bool, error) {
	return _TeeOwnerAllowlist.Contract.IsAllowedTeeMachineOwner(&_TeeOwnerAllowlist.CallOpts, _extensionId, _owner)
}

// IsAllowedTeeWalletProjectOwner is a free data retrieval call binding the contract method 0xab1c7ea1.
//
// Solidity: function isAllowedTeeWalletProjectOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCaller) IsAllowedTeeWalletProjectOwner(opts *bind.CallOpts, _extensionId *big.Int, _owner common.Address) (bool, error) {
	var out []interface{}
	err := _TeeOwnerAllowlist.contract.Call(opts, &out, "isAllowedTeeWalletProjectOwner", _extensionId, _owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedTeeWalletProjectOwner is a free data retrieval call binding the contract method 0xab1c7ea1.
//
// Solidity: function isAllowedTeeWalletProjectOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) IsAllowedTeeWalletProjectOwner(_extensionId *big.Int, _owner common.Address) (bool, error) {
	return _TeeOwnerAllowlist.Contract.IsAllowedTeeWalletProjectOwner(&_TeeOwnerAllowlist.CallOpts, _extensionId, _owner)
}

// IsAllowedTeeWalletProjectOwner is a free data retrieval call binding the contract method 0xab1c7ea1.
//
// Solidity: function isAllowedTeeWalletProjectOwner(uint256 _extensionId, address _owner) view returns(bool _isAllowed)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistCallerSession) IsAllowedTeeWalletProjectOwner(_extensionId *big.Int, _owner common.Address) (bool, error) {
	return _TeeOwnerAllowlist.Contract.IsAllowedTeeWalletProjectOwner(&_TeeOwnerAllowlist.CallOpts, _extensionId, _owner)
}

// AddAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0xa8f3cab7.
//
// Solidity: function addAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) AddAllowedTeeMachineOwners(opts *bind.TransactOpts, _extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "addAllowedTeeMachineOwners", _extensionId, _owners)
}

// AddAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0xa8f3cab7.
//
// Solidity: function addAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) AddAllowedTeeMachineOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AddAllowedTeeMachineOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// AddAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0xa8f3cab7.
//
// Solidity: function addAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) AddAllowedTeeMachineOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AddAllowedTeeMachineOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// AddAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0xd1c34746.
//
// Solidity: function addAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) AddAllowedTeeWalletProjectOwners(opts *bind.TransactOpts, _extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "addAllowedTeeWalletProjectOwners", _extensionId, _owners)
}

// AddAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0xd1c34746.
//
// Solidity: function addAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) AddAllowedTeeWalletProjectOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AddAllowedTeeWalletProjectOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// AddAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0xd1c34746.
//
// Solidity: function addAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) AddAllowedTeeWalletProjectOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AddAllowedTeeWalletProjectOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// AllowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0x1b489f2d.
//
// Solidity: function allowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) AllowAllTeeMachineOwners(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "allowAllTeeMachineOwners", _extensionId)
}

// AllowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0x1b489f2d.
//
// Solidity: function allowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) AllowAllTeeMachineOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AllowAllTeeMachineOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId)
}

// AllowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0x1b489f2d.
//
// Solidity: function allowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) AllowAllTeeMachineOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AllowAllTeeMachineOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId)
}

// AllowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x61f957e7.
//
// Solidity: function allowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) AllowAllTeeWalletProjectOwners(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "allowAllTeeWalletProjectOwners", _extensionId)
}

// AllowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x61f957e7.
//
// Solidity: function allowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) AllowAllTeeWalletProjectOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AllowAllTeeWalletProjectOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId)
}

// AllowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x61f957e7.
//
// Solidity: function allowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) AllowAllTeeWalletProjectOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.AllowAllTeeWalletProjectOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId)
}

// DisallowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0xb86936ba.
//
// Solidity: function disallowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) DisallowAllTeeMachineOwners(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "disallowAllTeeMachineOwners", _extensionId)
}

// DisallowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0xb86936ba.
//
// Solidity: function disallowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) DisallowAllTeeMachineOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.DisallowAllTeeMachineOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId)
}

// DisallowAllTeeMachineOwners is a paid mutator transaction binding the contract method 0xb86936ba.
//
// Solidity: function disallowAllTeeMachineOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) DisallowAllTeeMachineOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.DisallowAllTeeMachineOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId)
}

// DisallowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x919fd33a.
//
// Solidity: function disallowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) DisallowAllTeeWalletProjectOwners(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "disallowAllTeeWalletProjectOwners", _extensionId)
}

// DisallowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x919fd33a.
//
// Solidity: function disallowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) DisallowAllTeeWalletProjectOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.DisallowAllTeeWalletProjectOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId)
}

// DisallowAllTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x919fd33a.
//
// Solidity: function disallowAllTeeWalletProjectOwners(uint256 _extensionId) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) DisallowAllTeeWalletProjectOwners(_extensionId *big.Int) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.DisallowAllTeeWalletProjectOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId)
}

// RemoveAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0x5d371c3c.
//
// Solidity: function removeAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) RemoveAllowedTeeMachineOwners(opts *bind.TransactOpts, _extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "removeAllowedTeeMachineOwners", _extensionId, _owners)
}

// RemoveAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0x5d371c3c.
//
// Solidity: function removeAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) RemoveAllowedTeeMachineOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.RemoveAllowedTeeMachineOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// RemoveAllowedTeeMachineOwners is a paid mutator transaction binding the contract method 0x5d371c3c.
//
// Solidity: function removeAllowedTeeMachineOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) RemoveAllowedTeeMachineOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.RemoveAllowedTeeMachineOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// RemoveAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x62a2513d.
//
// Solidity: function removeAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactor) RemoveAllowedTeeWalletProjectOwners(opts *bind.TransactOpts, _extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.contract.Transact(opts, "removeAllowedTeeWalletProjectOwners", _extensionId, _owners)
}

// RemoveAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x62a2513d.
//
// Solidity: function removeAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistSession) RemoveAllowedTeeWalletProjectOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.RemoveAllowedTeeWalletProjectOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// RemoveAllowedTeeWalletProjectOwners is a paid mutator transaction binding the contract method 0x62a2513d.
//
// Solidity: function removeAllowedTeeWalletProjectOwners(uint256 _extensionId, address[] _owners) returns()
func (_TeeOwnerAllowlist *TeeOwnerAllowlistTransactorSession) RemoveAllowedTeeWalletProjectOwners(_extensionId *big.Int, _owners []common.Address) (*types.Transaction, error) {
	return _TeeOwnerAllowlist.Contract.RemoveAllowedTeeWalletProjectOwners(&_TeeOwnerAllowlist.TransactOpts, _extensionId, _owners)
}

// TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator is returned from FilterAllTeeMachineOwnersAllowed and is used to iterate over the raw logs and unpacked data for AllTeeMachineOwnersAllowed events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator struct {
	Event *TeeOwnerAllowlistAllTeeMachineOwnersAllowed // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistAllTeeMachineOwnersAllowed)
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
		it.Event = new(TeeOwnerAllowlistAllTeeMachineOwnersAllowed)
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
func (it *TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistAllTeeMachineOwnersAllowed represents a AllTeeMachineOwnersAllowed event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllTeeMachineOwnersAllowed struct {
	ExtensionId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllTeeMachineOwnersAllowed is a free log retrieval operation binding the contract event 0x1bc41c6dd61d3a656898250dfc58a14ddd4735f4480ae4c1581072ad19fcd8a5.
//
// Solidity: event AllTeeMachineOwnersAllowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterAllTeeMachineOwnersAllowed(opts *bind.FilterOpts) (*TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "AllTeeMachineOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistAllTeeMachineOwnersAllowedIterator{contract: _TeeOwnerAllowlist.contract, event: "AllTeeMachineOwnersAllowed", logs: logs, sub: sub}, nil
}

// WatchAllTeeMachineOwnersAllowed is a free log subscription operation binding the contract event 0x1bc41c6dd61d3a656898250dfc58a14ddd4735f4480ae4c1581072ad19fcd8a5.
//
// Solidity: event AllTeeMachineOwnersAllowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchAllTeeMachineOwnersAllowed(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistAllTeeMachineOwnersAllowed) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "AllTeeMachineOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistAllTeeMachineOwnersAllowed)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllTeeMachineOwnersAllowed", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseAllTeeMachineOwnersAllowed(log types.Log) (*TeeOwnerAllowlistAllTeeMachineOwnersAllowed, error) {
	event := new(TeeOwnerAllowlistAllTeeMachineOwnersAllowed)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllTeeMachineOwnersAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistAllTeeMachineOwnersDisallowedIterator is returned from FilterAllTeeMachineOwnersDisallowed and is used to iterate over the raw logs and unpacked data for AllTeeMachineOwnersDisallowed events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllTeeMachineOwnersDisallowedIterator struct {
	Event *TeeOwnerAllowlistAllTeeMachineOwnersDisallowed // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistAllTeeMachineOwnersDisallowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistAllTeeMachineOwnersDisallowed)
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
		it.Event = new(TeeOwnerAllowlistAllTeeMachineOwnersDisallowed)
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
func (it *TeeOwnerAllowlistAllTeeMachineOwnersDisallowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistAllTeeMachineOwnersDisallowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistAllTeeMachineOwnersDisallowed represents a AllTeeMachineOwnersDisallowed event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllTeeMachineOwnersDisallowed struct {
	ExtensionId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllTeeMachineOwnersDisallowed is a free log retrieval operation binding the contract event 0x76fc91f6aa03fa961257e10347e87c3844fd0547a18ab3967b5a52db80c353d3.
//
// Solidity: event AllTeeMachineOwnersDisallowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterAllTeeMachineOwnersDisallowed(opts *bind.FilterOpts) (*TeeOwnerAllowlistAllTeeMachineOwnersDisallowedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "AllTeeMachineOwnersDisallowed")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistAllTeeMachineOwnersDisallowedIterator{contract: _TeeOwnerAllowlist.contract, event: "AllTeeMachineOwnersDisallowed", logs: logs, sub: sub}, nil
}

// WatchAllTeeMachineOwnersDisallowed is a free log subscription operation binding the contract event 0x76fc91f6aa03fa961257e10347e87c3844fd0547a18ab3967b5a52db80c353d3.
//
// Solidity: event AllTeeMachineOwnersDisallowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchAllTeeMachineOwnersDisallowed(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistAllTeeMachineOwnersDisallowed) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "AllTeeMachineOwnersDisallowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistAllTeeMachineOwnersDisallowed)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllTeeMachineOwnersDisallowed", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseAllTeeMachineOwnersDisallowed(log types.Log) (*TeeOwnerAllowlistAllTeeMachineOwnersDisallowed, error) {
	event := new(TeeOwnerAllowlistAllTeeMachineOwnersDisallowed)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllTeeMachineOwnersDisallowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator is returned from FilterAllTeeWalletProjectOwnersAllowed and is used to iterate over the raw logs and unpacked data for AllTeeWalletProjectOwnersAllowed events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator struct {
	Event *TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed)
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
		it.Event = new(TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed)
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
func (it *TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed represents a AllTeeWalletProjectOwnersAllowed event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed struct {
	ExtensionId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllTeeWalletProjectOwnersAllowed is a free log retrieval operation binding the contract event 0x1936686ee6942840bf2283355eb720261b36284a52b74604253556ee61c08854.
//
// Solidity: event AllTeeWalletProjectOwnersAllowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterAllTeeWalletProjectOwnersAllowed(opts *bind.FilterOpts) (*TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "AllTeeWalletProjectOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowedIterator{contract: _TeeOwnerAllowlist.contract, event: "AllTeeWalletProjectOwnersAllowed", logs: logs, sub: sub}, nil
}

// WatchAllTeeWalletProjectOwnersAllowed is a free log subscription operation binding the contract event 0x1936686ee6942840bf2283355eb720261b36284a52b74604253556ee61c08854.
//
// Solidity: event AllTeeWalletProjectOwnersAllowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchAllTeeWalletProjectOwnersAllowed(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "AllTeeWalletProjectOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllTeeWalletProjectOwnersAllowed", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseAllTeeWalletProjectOwnersAllowed(log types.Log) (*TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed, error) {
	event := new(TeeOwnerAllowlistAllTeeWalletProjectOwnersAllowed)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllTeeWalletProjectOwnersAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator is returned from FilterAllTeeWalletProjectOwnersDisallowed and is used to iterate over the raw logs and unpacked data for AllTeeWalletProjectOwnersDisallowed events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator struct {
	Event *TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowed // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowed)
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
		it.Event = new(TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowed)
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
func (it *TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowed represents a AllTeeWalletProjectOwnersDisallowed event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowed struct {
	ExtensionId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllTeeWalletProjectOwnersDisallowed is a free log retrieval operation binding the contract event 0x1c6abd0ba2324be35a8b8dac7bef3bf25c5d2d2cc45c216ee3b38a82da6948e8.
//
// Solidity: event AllTeeWalletProjectOwnersDisallowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterAllTeeWalletProjectOwnersDisallowed(opts *bind.FilterOpts) (*TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "AllTeeWalletProjectOwnersDisallowed")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowedIterator{contract: _TeeOwnerAllowlist.contract, event: "AllTeeWalletProjectOwnersDisallowed", logs: logs, sub: sub}, nil
}

// WatchAllTeeWalletProjectOwnersDisallowed is a free log subscription operation binding the contract event 0x1c6abd0ba2324be35a8b8dac7bef3bf25c5d2d2cc45c216ee3b38a82da6948e8.
//
// Solidity: event AllTeeWalletProjectOwnersDisallowed(uint256 extensionId)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchAllTeeWalletProjectOwnersDisallowed(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowed) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "AllTeeWalletProjectOwnersDisallowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowed)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllTeeWalletProjectOwnersDisallowed", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseAllTeeWalletProjectOwnersDisallowed(log types.Log) (*TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowed, error) {
	event := new(TeeOwnerAllowlistAllTeeWalletProjectOwnersDisallowed)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllTeeWalletProjectOwnersDisallowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator is returned from FilterAllowedTeeMachineOwnersAdded and is used to iterate over the raw logs and unpacked data for AllowedTeeMachineOwnersAdded events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator struct {
	Event *TeeOwnerAllowlistAllowedTeeMachineOwnersAdded // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistAllowedTeeMachineOwnersAdded)
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
		it.Event = new(TeeOwnerAllowlistAllowedTeeMachineOwnersAdded)
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
func (it *TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistAllowedTeeMachineOwnersAdded represents a AllowedTeeMachineOwnersAdded event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllowedTeeMachineOwnersAdded struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllowedTeeMachineOwnersAdded is a free log retrieval operation binding the contract event 0xd392883c2acc8ffc90b60bb8783db156d990502913d5b159e209cd604475337c.
//
// Solidity: event AllowedTeeMachineOwnersAdded(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterAllowedTeeMachineOwnersAdded(opts *bind.FilterOpts) (*TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "AllowedTeeMachineOwnersAdded")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistAllowedTeeMachineOwnersAddedIterator{contract: _TeeOwnerAllowlist.contract, event: "AllowedTeeMachineOwnersAdded", logs: logs, sub: sub}, nil
}

// WatchAllowedTeeMachineOwnersAdded is a free log subscription operation binding the contract event 0xd392883c2acc8ffc90b60bb8783db156d990502913d5b159e209cd604475337c.
//
// Solidity: event AllowedTeeMachineOwnersAdded(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchAllowedTeeMachineOwnersAdded(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistAllowedTeeMachineOwnersAdded) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "AllowedTeeMachineOwnersAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistAllowedTeeMachineOwnersAdded)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllowedTeeMachineOwnersAdded", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseAllowedTeeMachineOwnersAdded(log types.Log) (*TeeOwnerAllowlistAllowedTeeMachineOwnersAdded, error) {
	event := new(TeeOwnerAllowlistAllowedTeeMachineOwnersAdded)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllowedTeeMachineOwnersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistAllowedTeeMachineOwnersRemovedIterator is returned from FilterAllowedTeeMachineOwnersRemoved and is used to iterate over the raw logs and unpacked data for AllowedTeeMachineOwnersRemoved events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllowedTeeMachineOwnersRemovedIterator struct {
	Event *TeeOwnerAllowlistAllowedTeeMachineOwnersRemoved // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistAllowedTeeMachineOwnersRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistAllowedTeeMachineOwnersRemoved)
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
		it.Event = new(TeeOwnerAllowlistAllowedTeeMachineOwnersRemoved)
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
func (it *TeeOwnerAllowlistAllowedTeeMachineOwnersRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistAllowedTeeMachineOwnersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistAllowedTeeMachineOwnersRemoved represents a AllowedTeeMachineOwnersRemoved event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllowedTeeMachineOwnersRemoved struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllowedTeeMachineOwnersRemoved is a free log retrieval operation binding the contract event 0x8444f01824cfebbd6772bffe7855e88270c338f6879216e756f86cd54c53ddf6.
//
// Solidity: event AllowedTeeMachineOwnersRemoved(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterAllowedTeeMachineOwnersRemoved(opts *bind.FilterOpts) (*TeeOwnerAllowlistAllowedTeeMachineOwnersRemovedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "AllowedTeeMachineOwnersRemoved")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistAllowedTeeMachineOwnersRemovedIterator{contract: _TeeOwnerAllowlist.contract, event: "AllowedTeeMachineOwnersRemoved", logs: logs, sub: sub}, nil
}

// WatchAllowedTeeMachineOwnersRemoved is a free log subscription operation binding the contract event 0x8444f01824cfebbd6772bffe7855e88270c338f6879216e756f86cd54c53ddf6.
//
// Solidity: event AllowedTeeMachineOwnersRemoved(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchAllowedTeeMachineOwnersRemoved(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistAllowedTeeMachineOwnersRemoved) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "AllowedTeeMachineOwnersRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistAllowedTeeMachineOwnersRemoved)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllowedTeeMachineOwnersRemoved", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseAllowedTeeMachineOwnersRemoved(log types.Log) (*TeeOwnerAllowlistAllowedTeeMachineOwnersRemoved, error) {
	event := new(TeeOwnerAllowlistAllowedTeeMachineOwnersRemoved)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllowedTeeMachineOwnersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator is returned from FilterAllowedTeeWalletProjectOwnersAdded and is used to iterate over the raw logs and unpacked data for AllowedTeeWalletProjectOwnersAdded events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator struct {
	Event *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded)
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
		it.Event = new(TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded)
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
func (it *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded represents a AllowedTeeWalletProjectOwnersAdded event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllowedTeeWalletProjectOwnersAdded is a free log retrieval operation binding the contract event 0x60340c199493a9a48f50ae5af13ebefed3adeeb1c4f0a5476d59654680ffdb0d.
//
// Solidity: event AllowedTeeWalletProjectOwnersAdded(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterAllowedTeeWalletProjectOwnersAdded(opts *bind.FilterOpts) (*TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "AllowedTeeWalletProjectOwnersAdded")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAddedIterator{contract: _TeeOwnerAllowlist.contract, event: "AllowedTeeWalletProjectOwnersAdded", logs: logs, sub: sub}, nil
}

// WatchAllowedTeeWalletProjectOwnersAdded is a free log subscription operation binding the contract event 0x60340c199493a9a48f50ae5af13ebefed3adeeb1c4f0a5476d59654680ffdb0d.
//
// Solidity: event AllowedTeeWalletProjectOwnersAdded(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchAllowedTeeWalletProjectOwnersAdded(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "AllowedTeeWalletProjectOwnersAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllowedTeeWalletProjectOwnersAdded", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseAllowedTeeWalletProjectOwnersAdded(log types.Log) (*TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded, error) {
	event := new(TeeOwnerAllowlistAllowedTeeWalletProjectOwnersAdded)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllowedTeeWalletProjectOwnersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator is returned from FilterAllowedTeeWalletProjectOwnersRemoved and is used to iterate over the raw logs and unpacked data for AllowedTeeWalletProjectOwnersRemoved events raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator struct {
	Event *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemoved // Event containing the contract specifics and raw log

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
func (it *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemoved)
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
		it.Event = new(TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemoved)
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
func (it *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemoved represents a AllowedTeeWalletProjectOwnersRemoved event raised by the TeeOwnerAllowlist contract.
type TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemoved struct {
	ExtensionId *big.Int
	Owners      []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllowedTeeWalletProjectOwnersRemoved is a free log retrieval operation binding the contract event 0x676aa060014299ae54ae19d4e1fd6fe0b4681aad07cc8ddf7979673262e3fcac.
//
// Solidity: event AllowedTeeWalletProjectOwnersRemoved(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) FilterAllowedTeeWalletProjectOwnersRemoved(opts *bind.FilterOpts) (*TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.FilterLogs(opts, "AllowedTeeWalletProjectOwnersRemoved")
	if err != nil {
		return nil, err
	}
	return &TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemovedIterator{contract: _TeeOwnerAllowlist.contract, event: "AllowedTeeWalletProjectOwnersRemoved", logs: logs, sub: sub}, nil
}

// WatchAllowedTeeWalletProjectOwnersRemoved is a free log subscription operation binding the contract event 0x676aa060014299ae54ae19d4e1fd6fe0b4681aad07cc8ddf7979673262e3fcac.
//
// Solidity: event AllowedTeeWalletProjectOwnersRemoved(uint256 extensionId, address[] owners)
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) WatchAllowedTeeWalletProjectOwnersRemoved(opts *bind.WatchOpts, sink chan<- *TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemoved) (event.Subscription, error) {

	logs, sub, err := _TeeOwnerAllowlist.contract.WatchLogs(opts, "AllowedTeeWalletProjectOwnersRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemoved)
				if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllowedTeeWalletProjectOwnersRemoved", log); err != nil {
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
func (_TeeOwnerAllowlist *TeeOwnerAllowlistFilterer) ParseAllowedTeeWalletProjectOwnersRemoved(log types.Log) (*TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemoved, error) {
	event := new(TeeOwnerAllowlistAllowedTeeWalletProjectOwnersRemoved)
	if err := _TeeOwnerAllowlist.contract.UnpackLog(event, "AllowedTeeWalletProjectOwnersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
