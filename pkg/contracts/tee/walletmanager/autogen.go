// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletmanager

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

// PublicKey is an auto generated low-level Go binding around an user-defined struct.
type PublicKey struct {
	X [32]byte
	Y [32]byte
}

// WalletManagerMetaData contains all meta data concerning the WalletManager contract.
var WalletManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AdminsNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"name\":\"DuplicatedPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdmin\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"name\":\"InvalidAdminPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdminsThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCosignersThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MultisigThresholdNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NotAllAdminsConfirmed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"NotAllCosignersConfirmed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughAdmins\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughKeys\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"WalletAdminConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structPublicKey[]\",\"name\":\"adminsPublicKeys\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"adminsThreshold\",\"type\":\"uint64\"}],\"name\":\"WalletAdminsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"WalletCosignerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"WalletCosignersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletPaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"closeWalletInitialization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"confirmAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"confirmCosigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"createWallet\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"enableWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getProjectWalletIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_walletIds\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAdminsAndThreshold\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_admins\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_adminsThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAdminsPublicKeysAndThreshold\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"_adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"_adminsThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletCosignersAndThreshold\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletProjectId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletStatus\",\"outputs\":[{\"internalType\":\"enumIWalletManager.WalletStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"pauseWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"_adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"_adminsThreshold\",\"type\":\"uint64\"}],\"name\":\"setAdmins\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"setCosigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WalletManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletManagerMetaData.ABI instead.
var WalletManagerABI = WalletManagerMetaData.ABI

// WalletManager is an auto generated Go binding around an Ethereum contract.
type WalletManager struct {
	WalletManagerCaller     // Read-only binding to the contract
	WalletManagerTransactor // Write-only binding to the contract
	WalletManagerFilterer   // Log filterer for contract events
}

// WalletManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletManagerSession struct {
	Contract     *WalletManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletManagerCallerSession struct {
	Contract *WalletManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WalletManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletManagerTransactorSession struct {
	Contract     *WalletManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WalletManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletManagerRaw struct {
	Contract *WalletManager // Generic contract binding to access the raw methods on
}

// WalletManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletManagerCallerRaw struct {
	Contract *WalletManagerCaller // Generic read-only contract binding to access the raw methods on
}

// WalletManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletManagerTransactorRaw struct {
	Contract *WalletManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletManager creates a new instance of WalletManager, bound to a specific deployed contract.
func NewWalletManager(address common.Address, backend bind.ContractBackend) (*WalletManager, error) {
	contract, err := bindWalletManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletManager{WalletManagerCaller: WalletManagerCaller{contract: contract}, WalletManagerTransactor: WalletManagerTransactor{contract: contract}, WalletManagerFilterer: WalletManagerFilterer{contract: contract}}, nil
}

// NewWalletManagerCaller creates a new read-only instance of WalletManager, bound to a specific deployed contract.
func NewWalletManagerCaller(address common.Address, caller bind.ContractCaller) (*WalletManagerCaller, error) {
	contract, err := bindWalletManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletManagerCaller{contract: contract}, nil
}

// NewWalletManagerTransactor creates a new write-only instance of WalletManager, bound to a specific deployed contract.
func NewWalletManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletManagerTransactor, error) {
	contract, err := bindWalletManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletManagerTransactor{contract: contract}, nil
}

// NewWalletManagerFilterer creates a new log filterer instance of WalletManager, bound to a specific deployed contract.
func NewWalletManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletManagerFilterer, error) {
	contract, err := bindWalletManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletManagerFilterer{contract: contract}, nil
}

// bindWalletManager binds a generic wrapper to an already deployed contract.
func bindWalletManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletManager *WalletManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletManager.Contract.WalletManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletManager *WalletManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletManager.Contract.WalletManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletManager *WalletManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletManager.Contract.WalletManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletManager *WalletManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletManager *WalletManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletManager *WalletManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletManager.Contract.contract.Transact(opts, method, params...)
}

// GetProjectWalletIds is a free data retrieval call binding the contract method 0x92911c30.
//
// Solidity: function getProjectWalletIds(bytes32 _projectId) view returns(bytes32[] _walletIds)
func (_WalletManager *WalletManagerCaller) GetProjectWalletIds(opts *bind.CallOpts, _projectId [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "getProjectWalletIds", _projectId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProjectWalletIds is a free data retrieval call binding the contract method 0x92911c30.
//
// Solidity: function getProjectWalletIds(bytes32 _projectId) view returns(bytes32[] _walletIds)
func (_WalletManager *WalletManagerSession) GetProjectWalletIds(_projectId [32]byte) ([][32]byte, error) {
	return _WalletManager.Contract.GetProjectWalletIds(&_WalletManager.CallOpts, _projectId)
}

// GetProjectWalletIds is a free data retrieval call binding the contract method 0x92911c30.
//
// Solidity: function getProjectWalletIds(bytes32 _projectId) view returns(bytes32[] _walletIds)
func (_WalletManager *WalletManagerCallerSession) GetProjectWalletIds(_projectId [32]byte) ([][32]byte, error) {
	return _WalletManager.Contract.GetProjectWalletIds(&_WalletManager.CallOpts, _projectId)
}

// GetWalletAdminsAndThreshold is a free data retrieval call binding the contract method 0x3309abd0.
//
// Solidity: function getWalletAdminsAndThreshold(bytes32 _walletId) view returns(address[] _admins, uint64 _adminsThreshold)
func (_WalletManager *WalletManagerCaller) GetWalletAdminsAndThreshold(opts *bind.CallOpts, _walletId [32]byte) (struct {
	Admins          []common.Address
	AdminsThreshold uint64
}, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "getWalletAdminsAndThreshold", _walletId)

	outstruct := new(struct {
		Admins          []common.Address
		AdminsThreshold uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Admins = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.AdminsThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetWalletAdminsAndThreshold is a free data retrieval call binding the contract method 0x3309abd0.
//
// Solidity: function getWalletAdminsAndThreshold(bytes32 _walletId) view returns(address[] _admins, uint64 _adminsThreshold)
func (_WalletManager *WalletManagerSession) GetWalletAdminsAndThreshold(_walletId [32]byte) (struct {
	Admins          []common.Address
	AdminsThreshold uint64
}, error) {
	return _WalletManager.Contract.GetWalletAdminsAndThreshold(&_WalletManager.CallOpts, _walletId)
}

// GetWalletAdminsAndThreshold is a free data retrieval call binding the contract method 0x3309abd0.
//
// Solidity: function getWalletAdminsAndThreshold(bytes32 _walletId) view returns(address[] _admins, uint64 _adminsThreshold)
func (_WalletManager *WalletManagerCallerSession) GetWalletAdminsAndThreshold(_walletId [32]byte) (struct {
	Admins          []common.Address
	AdminsThreshold uint64
}, error) {
	return _WalletManager.Contract.GetWalletAdminsAndThreshold(&_WalletManager.CallOpts, _walletId)
}

// GetWalletAdminsPublicKeysAndThreshold is a free data retrieval call binding the contract method 0xf3f0efb4.
//
// Solidity: function getWalletAdminsPublicKeysAndThreshold(bytes32 _walletId) view returns((bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold)
func (_WalletManager *WalletManagerCaller) GetWalletAdminsPublicKeysAndThreshold(opts *bind.CallOpts, _walletId [32]byte) (struct {
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
}, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "getWalletAdminsPublicKeysAndThreshold", _walletId)

	outstruct := new(struct {
		AdminsPublicKeys []PublicKey
		AdminsThreshold  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AdminsPublicKeys = *abi.ConvertType(out[0], new([]PublicKey)).(*[]PublicKey)
	outstruct.AdminsThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetWalletAdminsPublicKeysAndThreshold is a free data retrieval call binding the contract method 0xf3f0efb4.
//
// Solidity: function getWalletAdminsPublicKeysAndThreshold(bytes32 _walletId) view returns((bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold)
func (_WalletManager *WalletManagerSession) GetWalletAdminsPublicKeysAndThreshold(_walletId [32]byte) (struct {
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
}, error) {
	return _WalletManager.Contract.GetWalletAdminsPublicKeysAndThreshold(&_WalletManager.CallOpts, _walletId)
}

// GetWalletAdminsPublicKeysAndThreshold is a free data retrieval call binding the contract method 0xf3f0efb4.
//
// Solidity: function getWalletAdminsPublicKeysAndThreshold(bytes32 _walletId) view returns((bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold)
func (_WalletManager *WalletManagerCallerSession) GetWalletAdminsPublicKeysAndThreshold(_walletId [32]byte) (struct {
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
}, error) {
	return _WalletManager.Contract.GetWalletAdminsPublicKeysAndThreshold(&_WalletManager.CallOpts, _walletId)
}

// GetWalletCosignersAndThreshold is a free data retrieval call binding the contract method 0x881567e5.
//
// Solidity: function getWalletCosignersAndThreshold(bytes32 _walletId) view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_WalletManager *WalletManagerCaller) GetWalletCosignersAndThreshold(opts *bind.CallOpts, _walletId [32]byte) (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "getWalletCosignersAndThreshold", _walletId)

	outstruct := new(struct {
		Cosigners          []common.Address
		CosignersThreshold uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cosigners = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.CosignersThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetWalletCosignersAndThreshold is a free data retrieval call binding the contract method 0x881567e5.
//
// Solidity: function getWalletCosignersAndThreshold(bytes32 _walletId) view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_WalletManager *WalletManagerSession) GetWalletCosignersAndThreshold(_walletId [32]byte) (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	return _WalletManager.Contract.GetWalletCosignersAndThreshold(&_WalletManager.CallOpts, _walletId)
}

// GetWalletCosignersAndThreshold is a free data retrieval call binding the contract method 0x881567e5.
//
// Solidity: function getWalletCosignersAndThreshold(bytes32 _walletId) view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_WalletManager *WalletManagerCallerSession) GetWalletCosignersAndThreshold(_walletId [32]byte) (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	return _WalletManager.Contract.GetWalletCosignersAndThreshold(&_WalletManager.CallOpts, _walletId)
}

// GetWalletProjectId is a free data retrieval call binding the contract method 0xf8d5d64f.
//
// Solidity: function getWalletProjectId(bytes32 _walletId) view returns(bytes32 _projectId)
func (_WalletManager *WalletManagerCaller) GetWalletProjectId(opts *bind.CallOpts, _walletId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "getWalletProjectId", _walletId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWalletProjectId is a free data retrieval call binding the contract method 0xf8d5d64f.
//
// Solidity: function getWalletProjectId(bytes32 _walletId) view returns(bytes32 _projectId)
func (_WalletManager *WalletManagerSession) GetWalletProjectId(_walletId [32]byte) ([32]byte, error) {
	return _WalletManager.Contract.GetWalletProjectId(&_WalletManager.CallOpts, _walletId)
}

// GetWalletProjectId is a free data retrieval call binding the contract method 0xf8d5d64f.
//
// Solidity: function getWalletProjectId(bytes32 _walletId) view returns(bytes32 _projectId)
func (_WalletManager *WalletManagerCallerSession) GetWalletProjectId(_walletId [32]byte) ([32]byte, error) {
	return _WalletManager.Contract.GetWalletProjectId(&_WalletManager.CallOpts, _walletId)
}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 _walletId) view returns(uint8 _status)
func (_WalletManager *WalletManagerCaller) GetWalletStatus(opts *bind.CallOpts, _walletId [32]byte) (uint8, error) {
	var out []interface{}
	err := _WalletManager.contract.Call(opts, &out, "getWalletStatus", _walletId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 _walletId) view returns(uint8 _status)
func (_WalletManager *WalletManagerSession) GetWalletStatus(_walletId [32]byte) (uint8, error) {
	return _WalletManager.Contract.GetWalletStatus(&_WalletManager.CallOpts, _walletId)
}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 _walletId) view returns(uint8 _status)
func (_WalletManager *WalletManagerCallerSession) GetWalletStatus(_walletId [32]byte) (uint8, error) {
	return _WalletManager.Contract.GetWalletStatus(&_WalletManager.CallOpts, _walletId)
}

// CloseWalletInitialization is a paid mutator transaction binding the contract method 0xa3c6d865.
//
// Solidity: function closeWalletInitialization(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerTransactor) CloseWalletInitialization(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "closeWalletInitialization", _walletId)
}

// CloseWalletInitialization is a paid mutator transaction binding the contract method 0xa3c6d865.
//
// Solidity: function closeWalletInitialization(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerSession) CloseWalletInitialization(_walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.Contract.CloseWalletInitialization(&_WalletManager.TransactOpts, _walletId)
}

// CloseWalletInitialization is a paid mutator transaction binding the contract method 0xa3c6d865.
//
// Solidity: function closeWalletInitialization(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerTransactorSession) CloseWalletInitialization(_walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.Contract.CloseWalletInitialization(&_WalletManager.TransactOpts, _walletId)
}

// ConfirmAdmin is a paid mutator transaction binding the contract method 0x9ab3b981.
//
// Solidity: function confirmAdmin(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerTransactor) ConfirmAdmin(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "confirmAdmin", _walletId)
}

// ConfirmAdmin is a paid mutator transaction binding the contract method 0x9ab3b981.
//
// Solidity: function confirmAdmin(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerSession) ConfirmAdmin(_walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.Contract.ConfirmAdmin(&_WalletManager.TransactOpts, _walletId)
}

// ConfirmAdmin is a paid mutator transaction binding the contract method 0x9ab3b981.
//
// Solidity: function confirmAdmin(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerTransactorSession) ConfirmAdmin(_walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.Contract.ConfirmAdmin(&_WalletManager.TransactOpts, _walletId)
}

// ConfirmCosigner is a paid mutator transaction binding the contract method 0x77f74c1a.
//
// Solidity: function confirmCosigner(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerTransactor) ConfirmCosigner(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "confirmCosigner", _walletId)
}

// ConfirmCosigner is a paid mutator transaction binding the contract method 0x77f74c1a.
//
// Solidity: function confirmCosigner(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerSession) ConfirmCosigner(_walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.Contract.ConfirmCosigner(&_WalletManager.TransactOpts, _walletId)
}

// ConfirmCosigner is a paid mutator transaction binding the contract method 0x77f74c1a.
//
// Solidity: function confirmCosigner(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerTransactorSession) ConfirmCosigner(_walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.Contract.ConfirmCosigner(&_WalletManager.TransactOpts, _walletId)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x1d647605.
//
// Solidity: function createWallet(bytes32 _projectId) returns(bytes32 _walletId)
func (_WalletManager *WalletManagerTransactor) CreateWallet(opts *bind.TransactOpts, _projectId [32]byte) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "createWallet", _projectId)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x1d647605.
//
// Solidity: function createWallet(bytes32 _projectId) returns(bytes32 _walletId)
func (_WalletManager *WalletManagerSession) CreateWallet(_projectId [32]byte) (*types.Transaction, error) {
	return _WalletManager.Contract.CreateWallet(&_WalletManager.TransactOpts, _projectId)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x1d647605.
//
// Solidity: function createWallet(bytes32 _projectId) returns(bytes32 _walletId)
func (_WalletManager *WalletManagerTransactorSession) CreateWallet(_projectId [32]byte) (*types.Transaction, error) {
	return _WalletManager.Contract.CreateWallet(&_WalletManager.TransactOpts, _projectId)
}

// EnableWallet is a paid mutator transaction binding the contract method 0x676f358c.
//
// Solidity: function enableWallet(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerTransactor) EnableWallet(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "enableWallet", _walletId)
}

// EnableWallet is a paid mutator transaction binding the contract method 0x676f358c.
//
// Solidity: function enableWallet(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerSession) EnableWallet(_walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.Contract.EnableWallet(&_WalletManager.TransactOpts, _walletId)
}

// EnableWallet is a paid mutator transaction binding the contract method 0x676f358c.
//
// Solidity: function enableWallet(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerTransactorSession) EnableWallet(_walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.Contract.EnableWallet(&_WalletManager.TransactOpts, _walletId)
}

// PauseWallet is a paid mutator transaction binding the contract method 0x963a310d.
//
// Solidity: function pauseWallet(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerTransactor) PauseWallet(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "pauseWallet", _walletId)
}

// PauseWallet is a paid mutator transaction binding the contract method 0x963a310d.
//
// Solidity: function pauseWallet(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerSession) PauseWallet(_walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.Contract.PauseWallet(&_WalletManager.TransactOpts, _walletId)
}

// PauseWallet is a paid mutator transaction binding the contract method 0x963a310d.
//
// Solidity: function pauseWallet(bytes32 _walletId) returns()
func (_WalletManager *WalletManagerTransactorSession) PauseWallet(_walletId [32]byte) (*types.Transaction, error) {
	return _WalletManager.Contract.PauseWallet(&_WalletManager.TransactOpts, _walletId)
}

// SetAdmins is a paid mutator transaction binding the contract method 0x42e1aa9a.
//
// Solidity: function setAdmins(bytes32 _walletId, (bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold) returns()
func (_WalletManager *WalletManagerTransactor) SetAdmins(opts *bind.TransactOpts, _walletId [32]byte, _adminsPublicKeys []PublicKey, _adminsThreshold uint64) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "setAdmins", _walletId, _adminsPublicKeys, _adminsThreshold)
}

// SetAdmins is a paid mutator transaction binding the contract method 0x42e1aa9a.
//
// Solidity: function setAdmins(bytes32 _walletId, (bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold) returns()
func (_WalletManager *WalletManagerSession) SetAdmins(_walletId [32]byte, _adminsPublicKeys []PublicKey, _adminsThreshold uint64) (*types.Transaction, error) {
	return _WalletManager.Contract.SetAdmins(&_WalletManager.TransactOpts, _walletId, _adminsPublicKeys, _adminsThreshold)
}

// SetAdmins is a paid mutator transaction binding the contract method 0x42e1aa9a.
//
// Solidity: function setAdmins(bytes32 _walletId, (bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold) returns()
func (_WalletManager *WalletManagerTransactorSession) SetAdmins(_walletId [32]byte, _adminsPublicKeys []PublicKey, _adminsThreshold uint64) (*types.Transaction, error) {
	return _WalletManager.Contract.SetAdmins(&_WalletManager.TransactOpts, _walletId, _adminsPublicKeys, _adminsThreshold)
}

// SetCosigners is a paid mutator transaction binding the contract method 0x3e00a2d4.
//
// Solidity: function setCosigners(bytes32 _walletId, address[] _cosigners, uint64 _cosignersThreshold) returns()
func (_WalletManager *WalletManagerTransactor) SetCosigners(opts *bind.TransactOpts, _walletId [32]byte, _cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _WalletManager.contract.Transact(opts, "setCosigners", _walletId, _cosigners, _cosignersThreshold)
}

// SetCosigners is a paid mutator transaction binding the contract method 0x3e00a2d4.
//
// Solidity: function setCosigners(bytes32 _walletId, address[] _cosigners, uint64 _cosignersThreshold) returns()
func (_WalletManager *WalletManagerSession) SetCosigners(_walletId [32]byte, _cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _WalletManager.Contract.SetCosigners(&_WalletManager.TransactOpts, _walletId, _cosigners, _cosignersThreshold)
}

// SetCosigners is a paid mutator transaction binding the contract method 0x3e00a2d4.
//
// Solidity: function setCosigners(bytes32 _walletId, address[] _cosigners, uint64 _cosignersThreshold) returns()
func (_WalletManager *WalletManagerTransactorSession) SetCosigners(_walletId [32]byte, _cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _WalletManager.Contract.SetCosigners(&_WalletManager.TransactOpts, _walletId, _cosigners, _cosignersThreshold)
}

// WalletManagerWalletAdminConfirmedIterator is returned from FilterWalletAdminConfirmed and is used to iterate over the raw logs and unpacked data for WalletAdminConfirmed events raised by the WalletManager contract.
type WalletManagerWalletAdminConfirmedIterator struct {
	Event *WalletManagerWalletAdminConfirmed // Event containing the contract specifics and raw log

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
func (it *WalletManagerWalletAdminConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerWalletAdminConfirmed)
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
		it.Event = new(WalletManagerWalletAdminConfirmed)
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
func (it *WalletManagerWalletAdminConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerWalletAdminConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerWalletAdminConfirmed represents a WalletAdminConfirmed event raised by the WalletManager contract.
type WalletManagerWalletAdminConfirmed struct {
	WalletId [32]byte
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletAdminConfirmed is a free log retrieval operation binding the contract event 0xefdffb32952334eb7adb154fb07bfa263335e73f8ad33a54aba665efe346cb08.
//
// Solidity: event WalletAdminConfirmed(bytes32 indexed walletId, address indexed admin)
func (_WalletManager *WalletManagerFilterer) FilterWalletAdminConfirmed(opts *bind.FilterOpts, walletId [][32]byte, admin []common.Address) (*WalletManagerWalletAdminConfirmedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "WalletAdminConfirmed", walletIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &WalletManagerWalletAdminConfirmedIterator{contract: _WalletManager.contract, event: "WalletAdminConfirmed", logs: logs, sub: sub}, nil
}

// WatchWalletAdminConfirmed is a free log subscription operation binding the contract event 0xefdffb32952334eb7adb154fb07bfa263335e73f8ad33a54aba665efe346cb08.
//
// Solidity: event WalletAdminConfirmed(bytes32 indexed walletId, address indexed admin)
func (_WalletManager *WalletManagerFilterer) WatchWalletAdminConfirmed(opts *bind.WatchOpts, sink chan<- *WalletManagerWalletAdminConfirmed, walletId [][32]byte, admin []common.Address) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "WalletAdminConfirmed", walletIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerWalletAdminConfirmed)
				if err := _WalletManager.contract.UnpackLog(event, "WalletAdminConfirmed", log); err != nil {
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

// ParseWalletAdminConfirmed is a log parse operation binding the contract event 0xefdffb32952334eb7adb154fb07bfa263335e73f8ad33a54aba665efe346cb08.
//
// Solidity: event WalletAdminConfirmed(bytes32 indexed walletId, address indexed admin)
func (_WalletManager *WalletManagerFilterer) ParseWalletAdminConfirmed(log types.Log) (*WalletManagerWalletAdminConfirmed, error) {
	event := new(WalletManagerWalletAdminConfirmed)
	if err := _WalletManager.contract.UnpackLog(event, "WalletAdminConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletManagerWalletAdminsSetIterator is returned from FilterWalletAdminsSet and is used to iterate over the raw logs and unpacked data for WalletAdminsSet events raised by the WalletManager contract.
type WalletManagerWalletAdminsSetIterator struct {
	Event *WalletManagerWalletAdminsSet // Event containing the contract specifics and raw log

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
func (it *WalletManagerWalletAdminsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerWalletAdminsSet)
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
		it.Event = new(WalletManagerWalletAdminsSet)
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
func (it *WalletManagerWalletAdminsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerWalletAdminsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerWalletAdminsSet represents a WalletAdminsSet event raised by the WalletManager contract.
type WalletManagerWalletAdminsSet struct {
	WalletId         [32]byte
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWalletAdminsSet is a free log retrieval operation binding the contract event 0xfac536eb8a80dfc8ca75bf6543d4b1f435e119551bafd3733fcb28850d29d8aa.
//
// Solidity: event WalletAdminsSet(bytes32 indexed walletId, (bytes32,bytes32)[] adminsPublicKeys, uint64 adminsThreshold)
func (_WalletManager *WalletManagerFilterer) FilterWalletAdminsSet(opts *bind.FilterOpts, walletId [][32]byte) (*WalletManagerWalletAdminsSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "WalletAdminsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletManagerWalletAdminsSetIterator{contract: _WalletManager.contract, event: "WalletAdminsSet", logs: logs, sub: sub}, nil
}

// WatchWalletAdminsSet is a free log subscription operation binding the contract event 0xfac536eb8a80dfc8ca75bf6543d4b1f435e119551bafd3733fcb28850d29d8aa.
//
// Solidity: event WalletAdminsSet(bytes32 indexed walletId, (bytes32,bytes32)[] adminsPublicKeys, uint64 adminsThreshold)
func (_WalletManager *WalletManagerFilterer) WatchWalletAdminsSet(opts *bind.WatchOpts, sink chan<- *WalletManagerWalletAdminsSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "WalletAdminsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerWalletAdminsSet)
				if err := _WalletManager.contract.UnpackLog(event, "WalletAdminsSet", log); err != nil {
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

// ParseWalletAdminsSet is a log parse operation binding the contract event 0xfac536eb8a80dfc8ca75bf6543d4b1f435e119551bafd3733fcb28850d29d8aa.
//
// Solidity: event WalletAdminsSet(bytes32 indexed walletId, (bytes32,bytes32)[] adminsPublicKeys, uint64 adminsThreshold)
func (_WalletManager *WalletManagerFilterer) ParseWalletAdminsSet(log types.Log) (*WalletManagerWalletAdminsSet, error) {
	event := new(WalletManagerWalletAdminsSet)
	if err := _WalletManager.contract.UnpackLog(event, "WalletAdminsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletManagerWalletCosignerConfirmedIterator is returned from FilterWalletCosignerConfirmed and is used to iterate over the raw logs and unpacked data for WalletCosignerConfirmed events raised by the WalletManager contract.
type WalletManagerWalletCosignerConfirmedIterator struct {
	Event *WalletManagerWalletCosignerConfirmed // Event containing the contract specifics and raw log

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
func (it *WalletManagerWalletCosignerConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerWalletCosignerConfirmed)
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
		it.Event = new(WalletManagerWalletCosignerConfirmed)
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
func (it *WalletManagerWalletCosignerConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerWalletCosignerConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerWalletCosignerConfirmed represents a WalletCosignerConfirmed event raised by the WalletManager contract.
type WalletManagerWalletCosignerConfirmed struct {
	WalletId [32]byte
	Cosigner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletCosignerConfirmed is a free log retrieval operation binding the contract event 0x10ab8f96b70aef2c884291a57e690b6d9e4d97ca15e4c8b2514626b8d0cd023b.
//
// Solidity: event WalletCosignerConfirmed(bytes32 indexed walletId, address indexed cosigner)
func (_WalletManager *WalletManagerFilterer) FilterWalletCosignerConfirmed(opts *bind.FilterOpts, walletId [][32]byte, cosigner []common.Address) (*WalletManagerWalletCosignerConfirmedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var cosignerRule []interface{}
	for _, cosignerItem := range cosigner {
		cosignerRule = append(cosignerRule, cosignerItem)
	}

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "WalletCosignerConfirmed", walletIdRule, cosignerRule)
	if err != nil {
		return nil, err
	}
	return &WalletManagerWalletCosignerConfirmedIterator{contract: _WalletManager.contract, event: "WalletCosignerConfirmed", logs: logs, sub: sub}, nil
}

// WatchWalletCosignerConfirmed is a free log subscription operation binding the contract event 0x10ab8f96b70aef2c884291a57e690b6d9e4d97ca15e4c8b2514626b8d0cd023b.
//
// Solidity: event WalletCosignerConfirmed(bytes32 indexed walletId, address indexed cosigner)
func (_WalletManager *WalletManagerFilterer) WatchWalletCosignerConfirmed(opts *bind.WatchOpts, sink chan<- *WalletManagerWalletCosignerConfirmed, walletId [][32]byte, cosigner []common.Address) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var cosignerRule []interface{}
	for _, cosignerItem := range cosigner {
		cosignerRule = append(cosignerRule, cosignerItem)
	}

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "WalletCosignerConfirmed", walletIdRule, cosignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerWalletCosignerConfirmed)
				if err := _WalletManager.contract.UnpackLog(event, "WalletCosignerConfirmed", log); err != nil {
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

// ParseWalletCosignerConfirmed is a log parse operation binding the contract event 0x10ab8f96b70aef2c884291a57e690b6d9e4d97ca15e4c8b2514626b8d0cd023b.
//
// Solidity: event WalletCosignerConfirmed(bytes32 indexed walletId, address indexed cosigner)
func (_WalletManager *WalletManagerFilterer) ParseWalletCosignerConfirmed(log types.Log) (*WalletManagerWalletCosignerConfirmed, error) {
	event := new(WalletManagerWalletCosignerConfirmed)
	if err := _WalletManager.contract.UnpackLog(event, "WalletCosignerConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletManagerWalletCosignersSetIterator is returned from FilterWalletCosignersSet and is used to iterate over the raw logs and unpacked data for WalletCosignersSet events raised by the WalletManager contract.
type WalletManagerWalletCosignersSetIterator struct {
	Event *WalletManagerWalletCosignersSet // Event containing the contract specifics and raw log

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
func (it *WalletManagerWalletCosignersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerWalletCosignersSet)
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
		it.Event = new(WalletManagerWalletCosignersSet)
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
func (it *WalletManagerWalletCosignersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerWalletCosignersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerWalletCosignersSet represents a WalletCosignersSet event raised by the WalletManager contract.
type WalletManagerWalletCosignersSet struct {
	WalletId           [32]byte
	Cosigners          []common.Address
	CosignersThreshold uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWalletCosignersSet is a free log retrieval operation binding the contract event 0x48419616e657a6824f128cbf23c09b8b4750e1c50999dc05dd869a97ac4a9c31.
//
// Solidity: event WalletCosignersSet(bytes32 indexed walletId, address[] cosigners, uint64 cosignersThreshold)
func (_WalletManager *WalletManagerFilterer) FilterWalletCosignersSet(opts *bind.FilterOpts, walletId [][32]byte) (*WalletManagerWalletCosignersSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "WalletCosignersSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletManagerWalletCosignersSetIterator{contract: _WalletManager.contract, event: "WalletCosignersSet", logs: logs, sub: sub}, nil
}

// WatchWalletCosignersSet is a free log subscription operation binding the contract event 0x48419616e657a6824f128cbf23c09b8b4750e1c50999dc05dd869a97ac4a9c31.
//
// Solidity: event WalletCosignersSet(bytes32 indexed walletId, address[] cosigners, uint64 cosignersThreshold)
func (_WalletManager *WalletManagerFilterer) WatchWalletCosignersSet(opts *bind.WatchOpts, sink chan<- *WalletManagerWalletCosignersSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "WalletCosignersSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerWalletCosignersSet)
				if err := _WalletManager.contract.UnpackLog(event, "WalletCosignersSet", log); err != nil {
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

// ParseWalletCosignersSet is a log parse operation binding the contract event 0x48419616e657a6824f128cbf23c09b8b4750e1c50999dc05dd869a97ac4a9c31.
//
// Solidity: event WalletCosignersSet(bytes32 indexed walletId, address[] cosigners, uint64 cosignersThreshold)
func (_WalletManager *WalletManagerFilterer) ParseWalletCosignersSet(log types.Log) (*WalletManagerWalletCosignersSet, error) {
	event := new(WalletManagerWalletCosignersSet)
	if err := _WalletManager.contract.UnpackLog(event, "WalletCosignersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletManagerWalletCreatedIterator is returned from FilterWalletCreated and is used to iterate over the raw logs and unpacked data for WalletCreated events raised by the WalletManager contract.
type WalletManagerWalletCreatedIterator struct {
	Event *WalletManagerWalletCreated // Event containing the contract specifics and raw log

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
func (it *WalletManagerWalletCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerWalletCreated)
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
		it.Event = new(WalletManagerWalletCreated)
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
func (it *WalletManagerWalletCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerWalletCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerWalletCreated represents a WalletCreated event raised by the WalletManager contract.
type WalletManagerWalletCreated struct {
	ProjectId [32]byte
	WalletId  [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletCreated is a free log retrieval operation binding the contract event 0xbe8f27cef1f3d94120c9c547c3614f5b992fdb0c0a497cc920fde06546291ab4.
//
// Solidity: event WalletCreated(bytes32 indexed projectId, bytes32 indexed walletId)
func (_WalletManager *WalletManagerFilterer) FilterWalletCreated(opts *bind.FilterOpts, projectId [][32]byte, walletId [][32]byte) (*WalletManagerWalletCreatedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "WalletCreated", projectIdRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletManagerWalletCreatedIterator{contract: _WalletManager.contract, event: "WalletCreated", logs: logs, sub: sub}, nil
}

// WatchWalletCreated is a free log subscription operation binding the contract event 0xbe8f27cef1f3d94120c9c547c3614f5b992fdb0c0a497cc920fde06546291ab4.
//
// Solidity: event WalletCreated(bytes32 indexed projectId, bytes32 indexed walletId)
func (_WalletManager *WalletManagerFilterer) WatchWalletCreated(opts *bind.WatchOpts, sink chan<- *WalletManagerWalletCreated, projectId [][32]byte, walletId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "WalletCreated", projectIdRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerWalletCreated)
				if err := _WalletManager.contract.UnpackLog(event, "WalletCreated", log); err != nil {
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

// ParseWalletCreated is a log parse operation binding the contract event 0xbe8f27cef1f3d94120c9c547c3614f5b992fdb0c0a497cc920fde06546291ab4.
//
// Solidity: event WalletCreated(bytes32 indexed projectId, bytes32 indexed walletId)
func (_WalletManager *WalletManagerFilterer) ParseWalletCreated(log types.Log) (*WalletManagerWalletCreated, error) {
	event := new(WalletManagerWalletCreated)
	if err := _WalletManager.contract.UnpackLog(event, "WalletCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletManagerWalletEnabledIterator is returned from FilterWalletEnabled and is used to iterate over the raw logs and unpacked data for WalletEnabled events raised by the WalletManager contract.
type WalletManagerWalletEnabledIterator struct {
	Event *WalletManagerWalletEnabled // Event containing the contract specifics and raw log

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
func (it *WalletManagerWalletEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerWalletEnabled)
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
		it.Event = new(WalletManagerWalletEnabled)
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
func (it *WalletManagerWalletEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerWalletEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerWalletEnabled represents a WalletEnabled event raised by the WalletManager contract.
type WalletManagerWalletEnabled struct {
	WalletId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletEnabled is a free log retrieval operation binding the contract event 0x9dd94a650e5c2e19300b404a2e83ff44d26999ca24996bbff99dcd70f81661bf.
//
// Solidity: event WalletEnabled(bytes32 indexed walletId)
func (_WalletManager *WalletManagerFilterer) FilterWalletEnabled(opts *bind.FilterOpts, walletId [][32]byte) (*WalletManagerWalletEnabledIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "WalletEnabled", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletManagerWalletEnabledIterator{contract: _WalletManager.contract, event: "WalletEnabled", logs: logs, sub: sub}, nil
}

// WatchWalletEnabled is a free log subscription operation binding the contract event 0x9dd94a650e5c2e19300b404a2e83ff44d26999ca24996bbff99dcd70f81661bf.
//
// Solidity: event WalletEnabled(bytes32 indexed walletId)
func (_WalletManager *WalletManagerFilterer) WatchWalletEnabled(opts *bind.WatchOpts, sink chan<- *WalletManagerWalletEnabled, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "WalletEnabled", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerWalletEnabled)
				if err := _WalletManager.contract.UnpackLog(event, "WalletEnabled", log); err != nil {
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

// ParseWalletEnabled is a log parse operation binding the contract event 0x9dd94a650e5c2e19300b404a2e83ff44d26999ca24996bbff99dcd70f81661bf.
//
// Solidity: event WalletEnabled(bytes32 indexed walletId)
func (_WalletManager *WalletManagerFilterer) ParseWalletEnabled(log types.Log) (*WalletManagerWalletEnabled, error) {
	event := new(WalletManagerWalletEnabled)
	if err := _WalletManager.contract.UnpackLog(event, "WalletEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletManagerWalletInitializedIterator is returned from FilterWalletInitialized and is used to iterate over the raw logs and unpacked data for WalletInitialized events raised by the WalletManager contract.
type WalletManagerWalletInitializedIterator struct {
	Event *WalletManagerWalletInitialized // Event containing the contract specifics and raw log

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
func (it *WalletManagerWalletInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerWalletInitialized)
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
		it.Event = new(WalletManagerWalletInitialized)
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
func (it *WalletManagerWalletInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerWalletInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerWalletInitialized represents a WalletInitialized event raised by the WalletManager contract.
type WalletManagerWalletInitialized struct {
	WalletId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletInitialized is a free log retrieval operation binding the contract event 0x09032cc96c5d9d0239f5528329abee7a1c72f41813ebdd51e4f51f8f962c7223.
//
// Solidity: event WalletInitialized(bytes32 indexed walletId)
func (_WalletManager *WalletManagerFilterer) FilterWalletInitialized(opts *bind.FilterOpts, walletId [][32]byte) (*WalletManagerWalletInitializedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "WalletInitialized", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletManagerWalletInitializedIterator{contract: _WalletManager.contract, event: "WalletInitialized", logs: logs, sub: sub}, nil
}

// WatchWalletInitialized is a free log subscription operation binding the contract event 0x09032cc96c5d9d0239f5528329abee7a1c72f41813ebdd51e4f51f8f962c7223.
//
// Solidity: event WalletInitialized(bytes32 indexed walletId)
func (_WalletManager *WalletManagerFilterer) WatchWalletInitialized(opts *bind.WatchOpts, sink chan<- *WalletManagerWalletInitialized, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "WalletInitialized", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerWalletInitialized)
				if err := _WalletManager.contract.UnpackLog(event, "WalletInitialized", log); err != nil {
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

// ParseWalletInitialized is a log parse operation binding the contract event 0x09032cc96c5d9d0239f5528329abee7a1c72f41813ebdd51e4f51f8f962c7223.
//
// Solidity: event WalletInitialized(bytes32 indexed walletId)
func (_WalletManager *WalletManagerFilterer) ParseWalletInitialized(log types.Log) (*WalletManagerWalletInitialized, error) {
	event := new(WalletManagerWalletInitialized)
	if err := _WalletManager.contract.UnpackLog(event, "WalletInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletManagerWalletPausedIterator is returned from FilterWalletPaused and is used to iterate over the raw logs and unpacked data for WalletPaused events raised by the WalletManager contract.
type WalletManagerWalletPausedIterator struct {
	Event *WalletManagerWalletPaused // Event containing the contract specifics and raw log

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
func (it *WalletManagerWalletPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletManagerWalletPaused)
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
		it.Event = new(WalletManagerWalletPaused)
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
func (it *WalletManagerWalletPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletManagerWalletPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletManagerWalletPaused represents a WalletPaused event raised by the WalletManager contract.
type WalletManagerWalletPaused struct {
	WalletId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletPaused is a free log retrieval operation binding the contract event 0x5bc3bc472409bb4331cff333acd11b5d5b6d97052d783936292f19955b25eef3.
//
// Solidity: event WalletPaused(bytes32 indexed walletId)
func (_WalletManager *WalletManagerFilterer) FilterWalletPaused(opts *bind.FilterOpts, walletId [][32]byte) (*WalletManagerWalletPausedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletManager.contract.FilterLogs(opts, "WalletPaused", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletManagerWalletPausedIterator{contract: _WalletManager.contract, event: "WalletPaused", logs: logs, sub: sub}, nil
}

// WatchWalletPaused is a free log subscription operation binding the contract event 0x5bc3bc472409bb4331cff333acd11b5d5b6d97052d783936292f19955b25eef3.
//
// Solidity: event WalletPaused(bytes32 indexed walletId)
func (_WalletManager *WalletManagerFilterer) WatchWalletPaused(opts *bind.WatchOpts, sink chan<- *WalletManagerWalletPaused, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletManager.contract.WatchLogs(opts, "WalletPaused", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletManagerWalletPaused)
				if err := _WalletManager.contract.UnpackLog(event, "WalletPaused", log); err != nil {
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

// ParseWalletPaused is a log parse operation binding the contract event 0x5bc3bc472409bb4331cff333acd11b5d5b6d97052d783936292f19955b25eef3.
//
// Solidity: event WalletPaused(bytes32 indexed walletId)
func (_WalletManager *WalletManagerFilterer) ParseWalletPaused(log types.Log) (*WalletManagerWalletPaused, error) {
	event := new(WalletManagerWalletPaused)
	if err := _WalletManager.contract.UnpackLog(event, "WalletPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
