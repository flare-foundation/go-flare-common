// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletprojectpause

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

// WalletProjectPauseMetaData contains all meta data concerning the WalletProjectPause contract.
var WalletProjectPauseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWalletIds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"WalletProjectPausersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"WalletProjectPausersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"WalletProjectUnpausersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"WalletProjectUnpausersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"walletIds\",\"type\":\"bytes32[]\"}],\"name\":\"WalletsPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"walletIds\",\"type\":\"bytes32[]\"}],\"name\":\"WalletsUnpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addWalletProjectPausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addWalletProjectUnpausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getWalletProjectPausers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getWalletProjectUnpausers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isWalletProjectPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isPauser\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isWalletProjectUnpauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isUnpauser\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_walletIds\",\"type\":\"bytes32[]\"}],\"name\":\"pauseWallets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"removeWalletProjectPausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"removeWalletProjectUnpausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_walletIds\",\"type\":\"bytes32[]\"}],\"name\":\"unpauseWallets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WalletProjectPauseABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletProjectPauseMetaData.ABI instead.
var WalletProjectPauseABI = WalletProjectPauseMetaData.ABI

// WalletProjectPause is an auto generated Go binding around an Ethereum contract.
type WalletProjectPause struct {
	WalletProjectPauseCaller     // Read-only binding to the contract
	WalletProjectPauseTransactor // Write-only binding to the contract
	WalletProjectPauseFilterer   // Log filterer for contract events
}

// WalletProjectPauseCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletProjectPauseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletProjectPauseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletProjectPauseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletProjectPauseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletProjectPauseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletProjectPauseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletProjectPauseSession struct {
	Contract     *WalletProjectPause // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WalletProjectPauseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletProjectPauseCallerSession struct {
	Contract *WalletProjectPauseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// WalletProjectPauseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletProjectPauseTransactorSession struct {
	Contract     *WalletProjectPauseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// WalletProjectPauseRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletProjectPauseRaw struct {
	Contract *WalletProjectPause // Generic contract binding to access the raw methods on
}

// WalletProjectPauseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletProjectPauseCallerRaw struct {
	Contract *WalletProjectPauseCaller // Generic read-only contract binding to access the raw methods on
}

// WalletProjectPauseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletProjectPauseTransactorRaw struct {
	Contract *WalletProjectPauseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletProjectPause creates a new instance of WalletProjectPause, bound to a specific deployed contract.
func NewWalletProjectPause(address common.Address, backend bind.ContractBackend) (*WalletProjectPause, error) {
	contract, err := bindWalletProjectPause(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletProjectPause{WalletProjectPauseCaller: WalletProjectPauseCaller{contract: contract}, WalletProjectPauseTransactor: WalletProjectPauseTransactor{contract: contract}, WalletProjectPauseFilterer: WalletProjectPauseFilterer{contract: contract}}, nil
}

// NewWalletProjectPauseCaller creates a new read-only instance of WalletProjectPause, bound to a specific deployed contract.
func NewWalletProjectPauseCaller(address common.Address, caller bind.ContractCaller) (*WalletProjectPauseCaller, error) {
	contract, err := bindWalletProjectPause(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletProjectPauseCaller{contract: contract}, nil
}

// NewWalletProjectPauseTransactor creates a new write-only instance of WalletProjectPause, bound to a specific deployed contract.
func NewWalletProjectPauseTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletProjectPauseTransactor, error) {
	contract, err := bindWalletProjectPause(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletProjectPauseTransactor{contract: contract}, nil
}

// NewWalletProjectPauseFilterer creates a new log filterer instance of WalletProjectPause, bound to a specific deployed contract.
func NewWalletProjectPauseFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletProjectPauseFilterer, error) {
	contract, err := bindWalletProjectPause(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletProjectPauseFilterer{contract: contract}, nil
}

// bindWalletProjectPause binds a generic wrapper to an already deployed contract.
func bindWalletProjectPause(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletProjectPauseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletProjectPause *WalletProjectPauseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletProjectPause.Contract.WalletProjectPauseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletProjectPause *WalletProjectPauseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.WalletProjectPauseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletProjectPause *WalletProjectPauseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.WalletProjectPauseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletProjectPause *WalletProjectPauseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletProjectPause.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletProjectPause *WalletProjectPauseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletProjectPause *WalletProjectPauseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.contract.Transact(opts, method, params...)
}

// GetWalletProjectPausers is a free data retrieval call binding the contract method 0xefaddfab.
//
// Solidity: function getWalletProjectPausers(bytes32 _projectId) view returns(address[] _addresses)
func (_WalletProjectPause *WalletProjectPauseCaller) GetWalletProjectPausers(opts *bind.CallOpts, _projectId [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _WalletProjectPause.contract.Call(opts, &out, "getWalletProjectPausers", _projectId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWalletProjectPausers is a free data retrieval call binding the contract method 0xefaddfab.
//
// Solidity: function getWalletProjectPausers(bytes32 _projectId) view returns(address[] _addresses)
func (_WalletProjectPause *WalletProjectPauseSession) GetWalletProjectPausers(_projectId [32]byte) ([]common.Address, error) {
	return _WalletProjectPause.Contract.GetWalletProjectPausers(&_WalletProjectPause.CallOpts, _projectId)
}

// GetWalletProjectPausers is a free data retrieval call binding the contract method 0xefaddfab.
//
// Solidity: function getWalletProjectPausers(bytes32 _projectId) view returns(address[] _addresses)
func (_WalletProjectPause *WalletProjectPauseCallerSession) GetWalletProjectPausers(_projectId [32]byte) ([]common.Address, error) {
	return _WalletProjectPause.Contract.GetWalletProjectPausers(&_WalletProjectPause.CallOpts, _projectId)
}

// GetWalletProjectUnpausers is a free data retrieval call binding the contract method 0x3bb392db.
//
// Solidity: function getWalletProjectUnpausers(bytes32 _projectId) view returns(address[] _addresses)
func (_WalletProjectPause *WalletProjectPauseCaller) GetWalletProjectUnpausers(opts *bind.CallOpts, _projectId [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _WalletProjectPause.contract.Call(opts, &out, "getWalletProjectUnpausers", _projectId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWalletProjectUnpausers is a free data retrieval call binding the contract method 0x3bb392db.
//
// Solidity: function getWalletProjectUnpausers(bytes32 _projectId) view returns(address[] _addresses)
func (_WalletProjectPause *WalletProjectPauseSession) GetWalletProjectUnpausers(_projectId [32]byte) ([]common.Address, error) {
	return _WalletProjectPause.Contract.GetWalletProjectUnpausers(&_WalletProjectPause.CallOpts, _projectId)
}

// GetWalletProjectUnpausers is a free data retrieval call binding the contract method 0x3bb392db.
//
// Solidity: function getWalletProjectUnpausers(bytes32 _projectId) view returns(address[] _addresses)
func (_WalletProjectPause *WalletProjectPauseCallerSession) GetWalletProjectUnpausers(_projectId [32]byte) ([]common.Address, error) {
	return _WalletProjectPause.Contract.GetWalletProjectUnpausers(&_WalletProjectPause.CallOpts, _projectId)
}

// IsWalletProjectPauser is a free data retrieval call binding the contract method 0xfe27b3ee.
//
// Solidity: function isWalletProjectPauser(bytes32 _projectId, address _addr) view returns(bool _isPauser)
func (_WalletProjectPause *WalletProjectPauseCaller) IsWalletProjectPauser(opts *bind.CallOpts, _projectId [32]byte, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _WalletProjectPause.contract.Call(opts, &out, "isWalletProjectPauser", _projectId, _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWalletProjectPauser is a free data retrieval call binding the contract method 0xfe27b3ee.
//
// Solidity: function isWalletProjectPauser(bytes32 _projectId, address _addr) view returns(bool _isPauser)
func (_WalletProjectPause *WalletProjectPauseSession) IsWalletProjectPauser(_projectId [32]byte, _addr common.Address) (bool, error) {
	return _WalletProjectPause.Contract.IsWalletProjectPauser(&_WalletProjectPause.CallOpts, _projectId, _addr)
}

// IsWalletProjectPauser is a free data retrieval call binding the contract method 0xfe27b3ee.
//
// Solidity: function isWalletProjectPauser(bytes32 _projectId, address _addr) view returns(bool _isPauser)
func (_WalletProjectPause *WalletProjectPauseCallerSession) IsWalletProjectPauser(_projectId [32]byte, _addr common.Address) (bool, error) {
	return _WalletProjectPause.Contract.IsWalletProjectPauser(&_WalletProjectPause.CallOpts, _projectId, _addr)
}

// IsWalletProjectUnpauser is a free data retrieval call binding the contract method 0xe5fdda95.
//
// Solidity: function isWalletProjectUnpauser(bytes32 _projectId, address _addr) view returns(bool _isUnpauser)
func (_WalletProjectPause *WalletProjectPauseCaller) IsWalletProjectUnpauser(opts *bind.CallOpts, _projectId [32]byte, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _WalletProjectPause.contract.Call(opts, &out, "isWalletProjectUnpauser", _projectId, _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWalletProjectUnpauser is a free data retrieval call binding the contract method 0xe5fdda95.
//
// Solidity: function isWalletProjectUnpauser(bytes32 _projectId, address _addr) view returns(bool _isUnpauser)
func (_WalletProjectPause *WalletProjectPauseSession) IsWalletProjectUnpauser(_projectId [32]byte, _addr common.Address) (bool, error) {
	return _WalletProjectPause.Contract.IsWalletProjectUnpauser(&_WalletProjectPause.CallOpts, _projectId, _addr)
}

// IsWalletProjectUnpauser is a free data retrieval call binding the contract method 0xe5fdda95.
//
// Solidity: function isWalletProjectUnpauser(bytes32 _projectId, address _addr) view returns(bool _isUnpauser)
func (_WalletProjectPause *WalletProjectPauseCallerSession) IsWalletProjectUnpauser(_projectId [32]byte, _addr common.Address) (bool, error) {
	return _WalletProjectPause.Contract.IsWalletProjectUnpauser(&_WalletProjectPause.CallOpts, _projectId, _addr)
}

// AddWalletProjectPausers is a paid mutator transaction binding the contract method 0x2992a55f.
//
// Solidity: function addWalletProjectPausers(bytes32 _projectId, address[] _addresses) returns()
func (_WalletProjectPause *WalletProjectPauseTransactor) AddWalletProjectPausers(opts *bind.TransactOpts, _projectId [32]byte, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletProjectPause.contract.Transact(opts, "addWalletProjectPausers", _projectId, _addresses)
}

// AddWalletProjectPausers is a paid mutator transaction binding the contract method 0x2992a55f.
//
// Solidity: function addWalletProjectPausers(bytes32 _projectId, address[] _addresses) returns()
func (_WalletProjectPause *WalletProjectPauseSession) AddWalletProjectPausers(_projectId [32]byte, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.AddWalletProjectPausers(&_WalletProjectPause.TransactOpts, _projectId, _addresses)
}

// AddWalletProjectPausers is a paid mutator transaction binding the contract method 0x2992a55f.
//
// Solidity: function addWalletProjectPausers(bytes32 _projectId, address[] _addresses) returns()
func (_WalletProjectPause *WalletProjectPauseTransactorSession) AddWalletProjectPausers(_projectId [32]byte, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.AddWalletProjectPausers(&_WalletProjectPause.TransactOpts, _projectId, _addresses)
}

// AddWalletProjectUnpausers is a paid mutator transaction binding the contract method 0x90a84ccf.
//
// Solidity: function addWalletProjectUnpausers(bytes32 _projectId, address[] _addresses) returns()
func (_WalletProjectPause *WalletProjectPauseTransactor) AddWalletProjectUnpausers(opts *bind.TransactOpts, _projectId [32]byte, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletProjectPause.contract.Transact(opts, "addWalletProjectUnpausers", _projectId, _addresses)
}

// AddWalletProjectUnpausers is a paid mutator transaction binding the contract method 0x90a84ccf.
//
// Solidity: function addWalletProjectUnpausers(bytes32 _projectId, address[] _addresses) returns()
func (_WalletProjectPause *WalletProjectPauseSession) AddWalletProjectUnpausers(_projectId [32]byte, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.AddWalletProjectUnpausers(&_WalletProjectPause.TransactOpts, _projectId, _addresses)
}

// AddWalletProjectUnpausers is a paid mutator transaction binding the contract method 0x90a84ccf.
//
// Solidity: function addWalletProjectUnpausers(bytes32 _projectId, address[] _addresses) returns()
func (_WalletProjectPause *WalletProjectPauseTransactorSession) AddWalletProjectUnpausers(_projectId [32]byte, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.AddWalletProjectUnpausers(&_WalletProjectPause.TransactOpts, _projectId, _addresses)
}

// PauseWallets is a paid mutator transaction binding the contract method 0x7089cca6.
//
// Solidity: function pauseWallets(bytes32[] _walletIds) returns()
func (_WalletProjectPause *WalletProjectPauseTransactor) PauseWallets(opts *bind.TransactOpts, _walletIds [][32]byte) (*types.Transaction, error) {
	return _WalletProjectPause.contract.Transact(opts, "pauseWallets", _walletIds)
}

// PauseWallets is a paid mutator transaction binding the contract method 0x7089cca6.
//
// Solidity: function pauseWallets(bytes32[] _walletIds) returns()
func (_WalletProjectPause *WalletProjectPauseSession) PauseWallets(_walletIds [][32]byte) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.PauseWallets(&_WalletProjectPause.TransactOpts, _walletIds)
}

// PauseWallets is a paid mutator transaction binding the contract method 0x7089cca6.
//
// Solidity: function pauseWallets(bytes32[] _walletIds) returns()
func (_WalletProjectPause *WalletProjectPauseTransactorSession) PauseWallets(_walletIds [][32]byte) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.PauseWallets(&_WalletProjectPause.TransactOpts, _walletIds)
}

// RemoveWalletProjectPausers is a paid mutator transaction binding the contract method 0x7fb4b937.
//
// Solidity: function removeWalletProjectPausers(bytes32 _projectId, address[] _addresses) returns()
func (_WalletProjectPause *WalletProjectPauseTransactor) RemoveWalletProjectPausers(opts *bind.TransactOpts, _projectId [32]byte, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletProjectPause.contract.Transact(opts, "removeWalletProjectPausers", _projectId, _addresses)
}

// RemoveWalletProjectPausers is a paid mutator transaction binding the contract method 0x7fb4b937.
//
// Solidity: function removeWalletProjectPausers(bytes32 _projectId, address[] _addresses) returns()
func (_WalletProjectPause *WalletProjectPauseSession) RemoveWalletProjectPausers(_projectId [32]byte, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.RemoveWalletProjectPausers(&_WalletProjectPause.TransactOpts, _projectId, _addresses)
}

// RemoveWalletProjectPausers is a paid mutator transaction binding the contract method 0x7fb4b937.
//
// Solidity: function removeWalletProjectPausers(bytes32 _projectId, address[] _addresses) returns()
func (_WalletProjectPause *WalletProjectPauseTransactorSession) RemoveWalletProjectPausers(_projectId [32]byte, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.RemoveWalletProjectPausers(&_WalletProjectPause.TransactOpts, _projectId, _addresses)
}

// RemoveWalletProjectUnpausers is a paid mutator transaction binding the contract method 0x7250ea34.
//
// Solidity: function removeWalletProjectUnpausers(bytes32 _projectId, address[] _addresses) returns()
func (_WalletProjectPause *WalletProjectPauseTransactor) RemoveWalletProjectUnpausers(opts *bind.TransactOpts, _projectId [32]byte, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletProjectPause.contract.Transact(opts, "removeWalletProjectUnpausers", _projectId, _addresses)
}

// RemoveWalletProjectUnpausers is a paid mutator transaction binding the contract method 0x7250ea34.
//
// Solidity: function removeWalletProjectUnpausers(bytes32 _projectId, address[] _addresses) returns()
func (_WalletProjectPause *WalletProjectPauseSession) RemoveWalletProjectUnpausers(_projectId [32]byte, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.RemoveWalletProjectUnpausers(&_WalletProjectPause.TransactOpts, _projectId, _addresses)
}

// RemoveWalletProjectUnpausers is a paid mutator transaction binding the contract method 0x7250ea34.
//
// Solidity: function removeWalletProjectUnpausers(bytes32 _projectId, address[] _addresses) returns()
func (_WalletProjectPause *WalletProjectPauseTransactorSession) RemoveWalletProjectUnpausers(_projectId [32]byte, _addresses []common.Address) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.RemoveWalletProjectUnpausers(&_WalletProjectPause.TransactOpts, _projectId, _addresses)
}

// UnpauseWallets is a paid mutator transaction binding the contract method 0x5f85f96f.
//
// Solidity: function unpauseWallets(bytes32[] _walletIds) returns()
func (_WalletProjectPause *WalletProjectPauseTransactor) UnpauseWallets(opts *bind.TransactOpts, _walletIds [][32]byte) (*types.Transaction, error) {
	return _WalletProjectPause.contract.Transact(opts, "unpauseWallets", _walletIds)
}

// UnpauseWallets is a paid mutator transaction binding the contract method 0x5f85f96f.
//
// Solidity: function unpauseWallets(bytes32[] _walletIds) returns()
func (_WalletProjectPause *WalletProjectPauseSession) UnpauseWallets(_walletIds [][32]byte) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.UnpauseWallets(&_WalletProjectPause.TransactOpts, _walletIds)
}

// UnpauseWallets is a paid mutator transaction binding the contract method 0x5f85f96f.
//
// Solidity: function unpauseWallets(bytes32[] _walletIds) returns()
func (_WalletProjectPause *WalletProjectPauseTransactorSession) UnpauseWallets(_walletIds [][32]byte) (*types.Transaction, error) {
	return _WalletProjectPause.Contract.UnpauseWallets(&_WalletProjectPause.TransactOpts, _walletIds)
}

// WalletProjectPauseWalletProjectPausersAddedIterator is returned from FilterWalletProjectPausersAdded and is used to iterate over the raw logs and unpacked data for WalletProjectPausersAdded events raised by the WalletProjectPause contract.
type WalletProjectPauseWalletProjectPausersAddedIterator struct {
	Event *WalletProjectPauseWalletProjectPausersAdded // Event containing the contract specifics and raw log

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
func (it *WalletProjectPauseWalletProjectPausersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProjectPauseWalletProjectPausersAdded)
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
		it.Event = new(WalletProjectPauseWalletProjectPausersAdded)
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
func (it *WalletProjectPauseWalletProjectPausersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProjectPauseWalletProjectPausersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProjectPauseWalletProjectPausersAdded represents a WalletProjectPausersAdded event raised by the WalletProjectPause contract.
type WalletProjectPauseWalletProjectPausersAdded struct {
	ProjectId [32]byte
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletProjectPausersAdded is a free log retrieval operation binding the contract event 0x47dcc7cedaf5e2438829fc5add80740bc32d8f6303077f46fc87fc00503d6ca5.
//
// Solidity: event WalletProjectPausersAdded(bytes32 indexed projectId, address[] addresses)
func (_WalletProjectPause *WalletProjectPauseFilterer) FilterWalletProjectPausersAdded(opts *bind.FilterOpts, projectId [][32]byte) (*WalletProjectPauseWalletProjectPausersAddedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _WalletProjectPause.contract.FilterLogs(opts, "WalletProjectPausersAdded", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletProjectPauseWalletProjectPausersAddedIterator{contract: _WalletProjectPause.contract, event: "WalletProjectPausersAdded", logs: logs, sub: sub}, nil
}

// WatchWalletProjectPausersAdded is a free log subscription operation binding the contract event 0x47dcc7cedaf5e2438829fc5add80740bc32d8f6303077f46fc87fc00503d6ca5.
//
// Solidity: event WalletProjectPausersAdded(bytes32 indexed projectId, address[] addresses)
func (_WalletProjectPause *WalletProjectPauseFilterer) WatchWalletProjectPausersAdded(opts *bind.WatchOpts, sink chan<- *WalletProjectPauseWalletProjectPausersAdded, projectId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _WalletProjectPause.contract.WatchLogs(opts, "WalletProjectPausersAdded", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProjectPauseWalletProjectPausersAdded)
				if err := _WalletProjectPause.contract.UnpackLog(event, "WalletProjectPausersAdded", log); err != nil {
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

// ParseWalletProjectPausersAdded is a log parse operation binding the contract event 0x47dcc7cedaf5e2438829fc5add80740bc32d8f6303077f46fc87fc00503d6ca5.
//
// Solidity: event WalletProjectPausersAdded(bytes32 indexed projectId, address[] addresses)
func (_WalletProjectPause *WalletProjectPauseFilterer) ParseWalletProjectPausersAdded(log types.Log) (*WalletProjectPauseWalletProjectPausersAdded, error) {
	event := new(WalletProjectPauseWalletProjectPausersAdded)
	if err := _WalletProjectPause.contract.UnpackLog(event, "WalletProjectPausersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProjectPauseWalletProjectPausersRemovedIterator is returned from FilterWalletProjectPausersRemoved and is used to iterate over the raw logs and unpacked data for WalletProjectPausersRemoved events raised by the WalletProjectPause contract.
type WalletProjectPauseWalletProjectPausersRemovedIterator struct {
	Event *WalletProjectPauseWalletProjectPausersRemoved // Event containing the contract specifics and raw log

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
func (it *WalletProjectPauseWalletProjectPausersRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProjectPauseWalletProjectPausersRemoved)
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
		it.Event = new(WalletProjectPauseWalletProjectPausersRemoved)
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
func (it *WalletProjectPauseWalletProjectPausersRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProjectPauseWalletProjectPausersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProjectPauseWalletProjectPausersRemoved represents a WalletProjectPausersRemoved event raised by the WalletProjectPause contract.
type WalletProjectPauseWalletProjectPausersRemoved struct {
	ProjectId [32]byte
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletProjectPausersRemoved is a free log retrieval operation binding the contract event 0x15dae0511985f5dbaee75c976bd0682bad4322fd90cdfd5ba0416a960a747dd2.
//
// Solidity: event WalletProjectPausersRemoved(bytes32 indexed projectId, address[] addresses)
func (_WalletProjectPause *WalletProjectPauseFilterer) FilterWalletProjectPausersRemoved(opts *bind.FilterOpts, projectId [][32]byte) (*WalletProjectPauseWalletProjectPausersRemovedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _WalletProjectPause.contract.FilterLogs(opts, "WalletProjectPausersRemoved", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletProjectPauseWalletProjectPausersRemovedIterator{contract: _WalletProjectPause.contract, event: "WalletProjectPausersRemoved", logs: logs, sub: sub}, nil
}

// WatchWalletProjectPausersRemoved is a free log subscription operation binding the contract event 0x15dae0511985f5dbaee75c976bd0682bad4322fd90cdfd5ba0416a960a747dd2.
//
// Solidity: event WalletProjectPausersRemoved(bytes32 indexed projectId, address[] addresses)
func (_WalletProjectPause *WalletProjectPauseFilterer) WatchWalletProjectPausersRemoved(opts *bind.WatchOpts, sink chan<- *WalletProjectPauseWalletProjectPausersRemoved, projectId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _WalletProjectPause.contract.WatchLogs(opts, "WalletProjectPausersRemoved", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProjectPauseWalletProjectPausersRemoved)
				if err := _WalletProjectPause.contract.UnpackLog(event, "WalletProjectPausersRemoved", log); err != nil {
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

// ParseWalletProjectPausersRemoved is a log parse operation binding the contract event 0x15dae0511985f5dbaee75c976bd0682bad4322fd90cdfd5ba0416a960a747dd2.
//
// Solidity: event WalletProjectPausersRemoved(bytes32 indexed projectId, address[] addresses)
func (_WalletProjectPause *WalletProjectPauseFilterer) ParseWalletProjectPausersRemoved(log types.Log) (*WalletProjectPauseWalletProjectPausersRemoved, error) {
	event := new(WalletProjectPauseWalletProjectPausersRemoved)
	if err := _WalletProjectPause.contract.UnpackLog(event, "WalletProjectPausersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProjectPauseWalletProjectUnpausersAddedIterator is returned from FilterWalletProjectUnpausersAdded and is used to iterate over the raw logs and unpacked data for WalletProjectUnpausersAdded events raised by the WalletProjectPause contract.
type WalletProjectPauseWalletProjectUnpausersAddedIterator struct {
	Event *WalletProjectPauseWalletProjectUnpausersAdded // Event containing the contract specifics and raw log

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
func (it *WalletProjectPauseWalletProjectUnpausersAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProjectPauseWalletProjectUnpausersAdded)
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
		it.Event = new(WalletProjectPauseWalletProjectUnpausersAdded)
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
func (it *WalletProjectPauseWalletProjectUnpausersAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProjectPauseWalletProjectUnpausersAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProjectPauseWalletProjectUnpausersAdded represents a WalletProjectUnpausersAdded event raised by the WalletProjectPause contract.
type WalletProjectPauseWalletProjectUnpausersAdded struct {
	ProjectId [32]byte
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletProjectUnpausersAdded is a free log retrieval operation binding the contract event 0x91f773d11bba59c35174341bf35b18319b51cafa3c8eb66f2f5f4ad4926959d8.
//
// Solidity: event WalletProjectUnpausersAdded(bytes32 indexed projectId, address[] addresses)
func (_WalletProjectPause *WalletProjectPauseFilterer) FilterWalletProjectUnpausersAdded(opts *bind.FilterOpts, projectId [][32]byte) (*WalletProjectPauseWalletProjectUnpausersAddedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _WalletProjectPause.contract.FilterLogs(opts, "WalletProjectUnpausersAdded", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletProjectPauseWalletProjectUnpausersAddedIterator{contract: _WalletProjectPause.contract, event: "WalletProjectUnpausersAdded", logs: logs, sub: sub}, nil
}

// WatchWalletProjectUnpausersAdded is a free log subscription operation binding the contract event 0x91f773d11bba59c35174341bf35b18319b51cafa3c8eb66f2f5f4ad4926959d8.
//
// Solidity: event WalletProjectUnpausersAdded(bytes32 indexed projectId, address[] addresses)
func (_WalletProjectPause *WalletProjectPauseFilterer) WatchWalletProjectUnpausersAdded(opts *bind.WatchOpts, sink chan<- *WalletProjectPauseWalletProjectUnpausersAdded, projectId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _WalletProjectPause.contract.WatchLogs(opts, "WalletProjectUnpausersAdded", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProjectPauseWalletProjectUnpausersAdded)
				if err := _WalletProjectPause.contract.UnpackLog(event, "WalletProjectUnpausersAdded", log); err != nil {
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

// ParseWalletProjectUnpausersAdded is a log parse operation binding the contract event 0x91f773d11bba59c35174341bf35b18319b51cafa3c8eb66f2f5f4ad4926959d8.
//
// Solidity: event WalletProjectUnpausersAdded(bytes32 indexed projectId, address[] addresses)
func (_WalletProjectPause *WalletProjectPauseFilterer) ParseWalletProjectUnpausersAdded(log types.Log) (*WalletProjectPauseWalletProjectUnpausersAdded, error) {
	event := new(WalletProjectPauseWalletProjectUnpausersAdded)
	if err := _WalletProjectPause.contract.UnpackLog(event, "WalletProjectUnpausersAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProjectPauseWalletProjectUnpausersRemovedIterator is returned from FilterWalletProjectUnpausersRemoved and is used to iterate over the raw logs and unpacked data for WalletProjectUnpausersRemoved events raised by the WalletProjectPause contract.
type WalletProjectPauseWalletProjectUnpausersRemovedIterator struct {
	Event *WalletProjectPauseWalletProjectUnpausersRemoved // Event containing the contract specifics and raw log

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
func (it *WalletProjectPauseWalletProjectUnpausersRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProjectPauseWalletProjectUnpausersRemoved)
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
		it.Event = new(WalletProjectPauseWalletProjectUnpausersRemoved)
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
func (it *WalletProjectPauseWalletProjectUnpausersRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProjectPauseWalletProjectUnpausersRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProjectPauseWalletProjectUnpausersRemoved represents a WalletProjectUnpausersRemoved event raised by the WalletProjectPause contract.
type WalletProjectPauseWalletProjectUnpausersRemoved struct {
	ProjectId [32]byte
	Addresses []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletProjectUnpausersRemoved is a free log retrieval operation binding the contract event 0x5d0dab859c19dcdb279454666b1738fd267bdc5ebb33773dd80e3cbc53c85c9d.
//
// Solidity: event WalletProjectUnpausersRemoved(bytes32 indexed projectId, address[] addresses)
func (_WalletProjectPause *WalletProjectPauseFilterer) FilterWalletProjectUnpausersRemoved(opts *bind.FilterOpts, projectId [][32]byte) (*WalletProjectPauseWalletProjectUnpausersRemovedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _WalletProjectPause.contract.FilterLogs(opts, "WalletProjectUnpausersRemoved", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletProjectPauseWalletProjectUnpausersRemovedIterator{contract: _WalletProjectPause.contract, event: "WalletProjectUnpausersRemoved", logs: logs, sub: sub}, nil
}

// WatchWalletProjectUnpausersRemoved is a free log subscription operation binding the contract event 0x5d0dab859c19dcdb279454666b1738fd267bdc5ebb33773dd80e3cbc53c85c9d.
//
// Solidity: event WalletProjectUnpausersRemoved(bytes32 indexed projectId, address[] addresses)
func (_WalletProjectPause *WalletProjectPauseFilterer) WatchWalletProjectUnpausersRemoved(opts *bind.WatchOpts, sink chan<- *WalletProjectPauseWalletProjectUnpausersRemoved, projectId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _WalletProjectPause.contract.WatchLogs(opts, "WalletProjectUnpausersRemoved", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProjectPauseWalletProjectUnpausersRemoved)
				if err := _WalletProjectPause.contract.UnpackLog(event, "WalletProjectUnpausersRemoved", log); err != nil {
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

// ParseWalletProjectUnpausersRemoved is a log parse operation binding the contract event 0x5d0dab859c19dcdb279454666b1738fd267bdc5ebb33773dd80e3cbc53c85c9d.
//
// Solidity: event WalletProjectUnpausersRemoved(bytes32 indexed projectId, address[] addresses)
func (_WalletProjectPause *WalletProjectPauseFilterer) ParseWalletProjectUnpausersRemoved(log types.Log) (*WalletProjectPauseWalletProjectUnpausersRemoved, error) {
	event := new(WalletProjectPauseWalletProjectUnpausersRemoved)
	if err := _WalletProjectPause.contract.UnpackLog(event, "WalletProjectUnpausersRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProjectPauseWalletsPausedIterator is returned from FilterWalletsPaused and is used to iterate over the raw logs and unpacked data for WalletsPaused events raised by the WalletProjectPause contract.
type WalletProjectPauseWalletsPausedIterator struct {
	Event *WalletProjectPauseWalletsPaused // Event containing the contract specifics and raw log

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
func (it *WalletProjectPauseWalletsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProjectPauseWalletsPaused)
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
		it.Event = new(WalletProjectPauseWalletsPaused)
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
func (it *WalletProjectPauseWalletsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProjectPauseWalletsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProjectPauseWalletsPaused represents a WalletsPaused event raised by the WalletProjectPause contract.
type WalletProjectPauseWalletsPaused struct {
	WalletIds [][32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletsPaused is a free log retrieval operation binding the contract event 0xaf4f5cfb5a824adbe87460f7723d07293dca679879f5c889759c64bd521cfc23.
//
// Solidity: event WalletsPaused(bytes32[] walletIds)
func (_WalletProjectPause *WalletProjectPauseFilterer) FilterWalletsPaused(opts *bind.FilterOpts) (*WalletProjectPauseWalletsPausedIterator, error) {

	logs, sub, err := _WalletProjectPause.contract.FilterLogs(opts, "WalletsPaused")
	if err != nil {
		return nil, err
	}
	return &WalletProjectPauseWalletsPausedIterator{contract: _WalletProjectPause.contract, event: "WalletsPaused", logs: logs, sub: sub}, nil
}

// WatchWalletsPaused is a free log subscription operation binding the contract event 0xaf4f5cfb5a824adbe87460f7723d07293dca679879f5c889759c64bd521cfc23.
//
// Solidity: event WalletsPaused(bytes32[] walletIds)
func (_WalletProjectPause *WalletProjectPauseFilterer) WatchWalletsPaused(opts *bind.WatchOpts, sink chan<- *WalletProjectPauseWalletsPaused) (event.Subscription, error) {

	logs, sub, err := _WalletProjectPause.contract.WatchLogs(opts, "WalletsPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProjectPauseWalletsPaused)
				if err := _WalletProjectPause.contract.UnpackLog(event, "WalletsPaused", log); err != nil {
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

// ParseWalletsPaused is a log parse operation binding the contract event 0xaf4f5cfb5a824adbe87460f7723d07293dca679879f5c889759c64bd521cfc23.
//
// Solidity: event WalletsPaused(bytes32[] walletIds)
func (_WalletProjectPause *WalletProjectPauseFilterer) ParseWalletsPaused(log types.Log) (*WalletProjectPauseWalletsPaused, error) {
	event := new(WalletProjectPauseWalletsPaused)
	if err := _WalletProjectPause.contract.UnpackLog(event, "WalletsPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletProjectPauseWalletsUnpausedIterator is returned from FilterWalletsUnpaused and is used to iterate over the raw logs and unpacked data for WalletsUnpaused events raised by the WalletProjectPause contract.
type WalletProjectPauseWalletsUnpausedIterator struct {
	Event *WalletProjectPauseWalletsUnpaused // Event containing the contract specifics and raw log

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
func (it *WalletProjectPauseWalletsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletProjectPauseWalletsUnpaused)
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
		it.Event = new(WalletProjectPauseWalletsUnpaused)
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
func (it *WalletProjectPauseWalletsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletProjectPauseWalletsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletProjectPauseWalletsUnpaused represents a WalletsUnpaused event raised by the WalletProjectPause contract.
type WalletProjectPauseWalletsUnpaused struct {
	WalletIds [][32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletsUnpaused is a free log retrieval operation binding the contract event 0x276b89df3d722ab36196a1fda86f4e297a0be12e60543b9f4aaabfdc0f32fc89.
//
// Solidity: event WalletsUnpaused(bytes32[] walletIds)
func (_WalletProjectPause *WalletProjectPauseFilterer) FilterWalletsUnpaused(opts *bind.FilterOpts) (*WalletProjectPauseWalletsUnpausedIterator, error) {

	logs, sub, err := _WalletProjectPause.contract.FilterLogs(opts, "WalletsUnpaused")
	if err != nil {
		return nil, err
	}
	return &WalletProjectPauseWalletsUnpausedIterator{contract: _WalletProjectPause.contract, event: "WalletsUnpaused", logs: logs, sub: sub}, nil
}

// WatchWalletsUnpaused is a free log subscription operation binding the contract event 0x276b89df3d722ab36196a1fda86f4e297a0be12e60543b9f4aaabfdc0f32fc89.
//
// Solidity: event WalletsUnpaused(bytes32[] walletIds)
func (_WalletProjectPause *WalletProjectPauseFilterer) WatchWalletsUnpaused(opts *bind.WatchOpts, sink chan<- *WalletProjectPauseWalletsUnpaused) (event.Subscription, error) {

	logs, sub, err := _WalletProjectPause.contract.WatchLogs(opts, "WalletsUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletProjectPauseWalletsUnpaused)
				if err := _WalletProjectPause.contract.UnpackLog(event, "WalletsUnpaused", log); err != nil {
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

// ParseWalletsUnpaused is a log parse operation binding the contract event 0x276b89df3d722ab36196a1fda86f4e297a0be12e60543b9f4aaabfdc0f32fc89.
//
// Solidity: event WalletsUnpaused(bytes32[] walletIds)
func (_WalletProjectPause *WalletProjectPauseFilterer) ParseWalletsUnpaused(log types.Log) (*WalletProjectPauseWalletsUnpaused, error) {
	event := new(WalletProjectPauseWalletsUnpaused)
	if err := _WalletProjectPause.contract.UnpackLog(event, "WalletsUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
