// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teewalletmanager

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

// ITeeWalletManagerResumeKeyData is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletManagerResumeKeyData struct {
	KeyId uint64
	TeeId common.Address
	Nonce *big.Int
}

// PublicKey is an auto generated low-level Go binding around an user-defined struct.
type PublicKey struct {
	X [32]byte
	Y [32]byte
}

// TeeWalletManagerMetaData contains all meta data concerning the TeeWalletManager contract.
var TeeWalletManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"WalletAdminConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structPublicKey[]\",\"name\":\"adminsPublicKeys\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"adminsThreshold\",\"type\":\"uint64\"}],\"name\":\"WalletAdminsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"WalletCosignerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"WalletCosignersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletPaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"closeWalletInitialization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"confirmAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"confirmCosigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"createWallet\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"enableWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getProjectWalletIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_walletIds\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAdminsAndThreshold\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"_adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"_adminsThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletCosignersAndThreshold\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletProjectId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletStatus\",\"outputs\":[{\"internalType\":\"enumITeeWalletManager.WalletStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"pauseWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structITeeWalletManager.ResumeKeyData[]\",\"name\":\"_keysData\",\"type\":\"tuple[]\"}],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"_adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"_adminsThreshold\",\"type\":\"uint64\"}],\"name\":\"setAdmins\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"setCosigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_pausingAddresses\",\"type\":\"address[]\"}],\"name\":\"setPausingAddresses\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeeWalletManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeWalletManagerMetaData.ABI instead.
var TeeWalletManagerABI = TeeWalletManagerMetaData.ABI

// TeeWalletManager is an auto generated Go binding around an Ethereum contract.
type TeeWalletManager struct {
	TeeWalletManagerCaller     // Read-only binding to the contract
	TeeWalletManagerTransactor // Write-only binding to the contract
	TeeWalletManagerFilterer   // Log filterer for contract events
}

// TeeWalletManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeWalletManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeWalletManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeWalletManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeWalletManagerSession struct {
	Contract     *TeeWalletManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeWalletManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeWalletManagerCallerSession struct {
	Contract *TeeWalletManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TeeWalletManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeWalletManagerTransactorSession struct {
	Contract     *TeeWalletManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TeeWalletManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeWalletManagerRaw struct {
	Contract *TeeWalletManager // Generic contract binding to access the raw methods on
}

// TeeWalletManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeWalletManagerCallerRaw struct {
	Contract *TeeWalletManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TeeWalletManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeWalletManagerTransactorRaw struct {
	Contract *TeeWalletManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeWalletManager creates a new instance of TeeWalletManager, bound to a specific deployed contract.
func NewTeeWalletManager(address common.Address, backend bind.ContractBackend) (*TeeWalletManager, error) {
	contract, err := bindTeeWalletManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManager{TeeWalletManagerCaller: TeeWalletManagerCaller{contract: contract}, TeeWalletManagerTransactor: TeeWalletManagerTransactor{contract: contract}, TeeWalletManagerFilterer: TeeWalletManagerFilterer{contract: contract}}, nil
}

// NewTeeWalletManagerCaller creates a new read-only instance of TeeWalletManager, bound to a specific deployed contract.
func NewTeeWalletManagerCaller(address common.Address, caller bind.ContractCaller) (*TeeWalletManagerCaller, error) {
	contract, err := bindTeeWalletManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerCaller{contract: contract}, nil
}

// NewTeeWalletManagerTransactor creates a new write-only instance of TeeWalletManager, bound to a specific deployed contract.
func NewTeeWalletManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeWalletManagerTransactor, error) {
	contract, err := bindTeeWalletManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerTransactor{contract: contract}, nil
}

// NewTeeWalletManagerFilterer creates a new log filterer instance of TeeWalletManager, bound to a specific deployed contract.
func NewTeeWalletManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeWalletManagerFilterer, error) {
	contract, err := bindTeeWalletManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerFilterer{contract: contract}, nil
}

// bindTeeWalletManager binds a generic wrapper to an already deployed contract.
func bindTeeWalletManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeWalletManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletManager *TeeWalletManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletManager.Contract.TeeWalletManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletManager *TeeWalletManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.TeeWalletManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletManager *TeeWalletManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.TeeWalletManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletManager *TeeWalletManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletManager *TeeWalletManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletManager *TeeWalletManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.contract.Transact(opts, method, params...)
}

// GetProjectWalletIds is a free data retrieval call binding the contract method 0x92911c30.
//
// Solidity: function getProjectWalletIds(bytes32 _projectId) view returns(bytes32[] _walletIds)
func (_TeeWalletManager *TeeWalletManagerCaller) GetProjectWalletIds(opts *bind.CallOpts, _projectId [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getProjectWalletIds", _projectId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProjectWalletIds is a free data retrieval call binding the contract method 0x92911c30.
//
// Solidity: function getProjectWalletIds(bytes32 _projectId) view returns(bytes32[] _walletIds)
func (_TeeWalletManager *TeeWalletManagerSession) GetProjectWalletIds(_projectId [32]byte) ([][32]byte, error) {
	return _TeeWalletManager.Contract.GetProjectWalletIds(&_TeeWalletManager.CallOpts, _projectId)
}

// GetProjectWalletIds is a free data retrieval call binding the contract method 0x92911c30.
//
// Solidity: function getProjectWalletIds(bytes32 _projectId) view returns(bytes32[] _walletIds)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetProjectWalletIds(_projectId [32]byte) ([][32]byte, error) {
	return _TeeWalletManager.Contract.GetProjectWalletIds(&_TeeWalletManager.CallOpts, _projectId)
}

// GetWalletAdminsAndThreshold is a free data retrieval call binding the contract method 0x3309abd0.
//
// Solidity: function getWalletAdminsAndThreshold(bytes32 _walletId) view returns((bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletAdminsAndThreshold(opts *bind.CallOpts, _walletId [32]byte) (struct {
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
}, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletAdminsAndThreshold", _walletId)

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

// GetWalletAdminsAndThreshold is a free data retrieval call binding the contract method 0x3309abd0.
//
// Solidity: function getWalletAdminsAndThreshold(bytes32 _walletId) view returns((bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletAdminsAndThreshold(_walletId [32]byte) (struct {
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
}, error) {
	return _TeeWalletManager.Contract.GetWalletAdminsAndThreshold(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletAdminsAndThreshold is a free data retrieval call binding the contract method 0x3309abd0.
//
// Solidity: function getWalletAdminsAndThreshold(bytes32 _walletId) view returns((bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletAdminsAndThreshold(_walletId [32]byte) (struct {
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
}, error) {
	return _TeeWalletManager.Contract.GetWalletAdminsAndThreshold(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletCosignersAndThreshold is a free data retrieval call binding the contract method 0x881567e5.
//
// Solidity: function getWalletCosignersAndThreshold(bytes32 _walletId) view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletCosignersAndThreshold(opts *bind.CallOpts, _walletId [32]byte) (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletCosignersAndThreshold", _walletId)

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
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletCosignersAndThreshold(_walletId [32]byte) (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	return _TeeWalletManager.Contract.GetWalletCosignersAndThreshold(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletCosignersAndThreshold is a free data retrieval call binding the contract method 0x881567e5.
//
// Solidity: function getWalletCosignersAndThreshold(bytes32 _walletId) view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletCosignersAndThreshold(_walletId [32]byte) (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	return _TeeWalletManager.Contract.GetWalletCosignersAndThreshold(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletProjectId is a free data retrieval call binding the contract method 0xf8d5d64f.
//
// Solidity: function getWalletProjectId(bytes32 _walletId) view returns(bytes32 _projectId)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletProjectId(opts *bind.CallOpts, _walletId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletProjectId", _walletId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWalletProjectId is a free data retrieval call binding the contract method 0xf8d5d64f.
//
// Solidity: function getWalletProjectId(bytes32 _walletId) view returns(bytes32 _projectId)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletProjectId(_walletId [32]byte) ([32]byte, error) {
	return _TeeWalletManager.Contract.GetWalletProjectId(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletProjectId is a free data retrieval call binding the contract method 0xf8d5d64f.
//
// Solidity: function getWalletProjectId(bytes32 _walletId) view returns(bytes32 _projectId)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletProjectId(_walletId [32]byte) ([32]byte, error) {
	return _TeeWalletManager.Contract.GetWalletProjectId(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 _walletId) view returns(uint8 _status)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletStatus(opts *bind.CallOpts, _walletId [32]byte) (uint8, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletStatus", _walletId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 _walletId) view returns(uint8 _status)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletStatus(_walletId [32]byte) (uint8, error) {
	return _TeeWalletManager.Contract.GetWalletStatus(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 _walletId) view returns(uint8 _status)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletStatus(_walletId [32]byte) (uint8, error) {
	return _TeeWalletManager.Contract.GetWalletStatus(&_TeeWalletManager.CallOpts, _walletId)
}

// CloseWalletInitialization is a paid mutator transaction binding the contract method 0xa3c6d865.
//
// Solidity: function closeWalletInitialization(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) CloseWalletInitialization(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "closeWalletInitialization", _walletId)
}

// CloseWalletInitialization is a paid mutator transaction binding the contract method 0xa3c6d865.
//
// Solidity: function closeWalletInitialization(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerSession) CloseWalletInitialization(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.CloseWalletInitialization(&_TeeWalletManager.TransactOpts, _walletId)
}

// CloseWalletInitialization is a paid mutator transaction binding the contract method 0xa3c6d865.
//
// Solidity: function closeWalletInitialization(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) CloseWalletInitialization(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.CloseWalletInitialization(&_TeeWalletManager.TransactOpts, _walletId)
}

// ConfirmAdmin is a paid mutator transaction binding the contract method 0x9ab3b981.
//
// Solidity: function confirmAdmin(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) ConfirmAdmin(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "confirmAdmin", _walletId)
}

// ConfirmAdmin is a paid mutator transaction binding the contract method 0x9ab3b981.
//
// Solidity: function confirmAdmin(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerSession) ConfirmAdmin(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ConfirmAdmin(&_TeeWalletManager.TransactOpts, _walletId)
}

// ConfirmAdmin is a paid mutator transaction binding the contract method 0x9ab3b981.
//
// Solidity: function confirmAdmin(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) ConfirmAdmin(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ConfirmAdmin(&_TeeWalletManager.TransactOpts, _walletId)
}

// ConfirmCosigner is a paid mutator transaction binding the contract method 0x77f74c1a.
//
// Solidity: function confirmCosigner(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) ConfirmCosigner(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "confirmCosigner", _walletId)
}

// ConfirmCosigner is a paid mutator transaction binding the contract method 0x77f74c1a.
//
// Solidity: function confirmCosigner(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerSession) ConfirmCosigner(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ConfirmCosigner(&_TeeWalletManager.TransactOpts, _walletId)
}

// ConfirmCosigner is a paid mutator transaction binding the contract method 0x77f74c1a.
//
// Solidity: function confirmCosigner(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) ConfirmCosigner(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ConfirmCosigner(&_TeeWalletManager.TransactOpts, _walletId)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x1d647605.
//
// Solidity: function createWallet(bytes32 _projectId) returns(bytes32 _walletId)
func (_TeeWalletManager *TeeWalletManagerTransactor) CreateWallet(opts *bind.TransactOpts, _projectId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "createWallet", _projectId)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x1d647605.
//
// Solidity: function createWallet(bytes32 _projectId) returns(bytes32 _walletId)
func (_TeeWalletManager *TeeWalletManagerSession) CreateWallet(_projectId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.CreateWallet(&_TeeWalletManager.TransactOpts, _projectId)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x1d647605.
//
// Solidity: function createWallet(bytes32 _projectId) returns(bytes32 _walletId)
func (_TeeWalletManager *TeeWalletManagerTransactorSession) CreateWallet(_projectId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.CreateWallet(&_TeeWalletManager.TransactOpts, _projectId)
}

// EnableWallet is a paid mutator transaction binding the contract method 0x676f358c.
//
// Solidity: function enableWallet(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) EnableWallet(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "enableWallet", _walletId)
}

// EnableWallet is a paid mutator transaction binding the contract method 0x676f358c.
//
// Solidity: function enableWallet(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerSession) EnableWallet(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.EnableWallet(&_TeeWalletManager.TransactOpts, _walletId)
}

// EnableWallet is a paid mutator transaction binding the contract method 0x676f358c.
//
// Solidity: function enableWallet(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) EnableWallet(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.EnableWallet(&_TeeWalletManager.TransactOpts, _walletId)
}

// PauseWallet is a paid mutator transaction binding the contract method 0x963a310d.
//
// Solidity: function pauseWallet(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) PauseWallet(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "pauseWallet", _walletId)
}

// PauseWallet is a paid mutator transaction binding the contract method 0x963a310d.
//
// Solidity: function pauseWallet(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerSession) PauseWallet(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.PauseWallet(&_TeeWalletManager.TransactOpts, _walletId)
}

// PauseWallet is a paid mutator transaction binding the contract method 0x963a310d.
//
// Solidity: function pauseWallet(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) PauseWallet(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.PauseWallet(&_TeeWalletManager.TransactOpts, _walletId)
}

// Resume is a paid mutator transaction binding the contract method 0x33921204.
//
// Solidity: function resume(bytes32 _walletId, (uint64,address,uint256)[] _keysData) payable returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) Resume(opts *bind.TransactOpts, _walletId [32]byte, _keysData []ITeeWalletManagerResumeKeyData) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "resume", _walletId, _keysData)
}

// Resume is a paid mutator transaction binding the contract method 0x33921204.
//
// Solidity: function resume(bytes32 _walletId, (uint64,address,uint256)[] _keysData) payable returns()
func (_TeeWalletManager *TeeWalletManagerSession) Resume(_walletId [32]byte, _keysData []ITeeWalletManagerResumeKeyData) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.Resume(&_TeeWalletManager.TransactOpts, _walletId, _keysData)
}

// Resume is a paid mutator transaction binding the contract method 0x33921204.
//
// Solidity: function resume(bytes32 _walletId, (uint64,address,uint256)[] _keysData) payable returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) Resume(_walletId [32]byte, _keysData []ITeeWalletManagerResumeKeyData) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.Resume(&_TeeWalletManager.TransactOpts, _walletId, _keysData)
}

// SetAdmins is a paid mutator transaction binding the contract method 0x42e1aa9a.
//
// Solidity: function setAdmins(bytes32 _walletId, (bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) SetAdmins(opts *bind.TransactOpts, _walletId [32]byte, _adminsPublicKeys []PublicKey, _adminsThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "setAdmins", _walletId, _adminsPublicKeys, _adminsThreshold)
}

// SetAdmins is a paid mutator transaction binding the contract method 0x42e1aa9a.
//
// Solidity: function setAdmins(bytes32 _walletId, (bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold) returns()
func (_TeeWalletManager *TeeWalletManagerSession) SetAdmins(_walletId [32]byte, _adminsPublicKeys []PublicKey, _adminsThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetAdmins(&_TeeWalletManager.TransactOpts, _walletId, _adminsPublicKeys, _adminsThreshold)
}

// SetAdmins is a paid mutator transaction binding the contract method 0x42e1aa9a.
//
// Solidity: function setAdmins(bytes32 _walletId, (bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) SetAdmins(_walletId [32]byte, _adminsPublicKeys []PublicKey, _adminsThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetAdmins(&_TeeWalletManager.TransactOpts, _walletId, _adminsPublicKeys, _adminsThreshold)
}

// SetCosigners is a paid mutator transaction binding the contract method 0x3e00a2d4.
//
// Solidity: function setCosigners(bytes32 _walletId, address[] _cosigners, uint64 _cosignersThreshold) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) SetCosigners(opts *bind.TransactOpts, _walletId [32]byte, _cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "setCosigners", _walletId, _cosigners, _cosignersThreshold)
}

// SetCosigners is a paid mutator transaction binding the contract method 0x3e00a2d4.
//
// Solidity: function setCosigners(bytes32 _walletId, address[] _cosigners, uint64 _cosignersThreshold) returns()
func (_TeeWalletManager *TeeWalletManagerSession) SetCosigners(_walletId [32]byte, _cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetCosigners(&_TeeWalletManager.TransactOpts, _walletId, _cosigners, _cosignersThreshold)
}

// SetCosigners is a paid mutator transaction binding the contract method 0x3e00a2d4.
//
// Solidity: function setCosigners(bytes32 _walletId, address[] _cosigners, uint64 _cosignersThreshold) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) SetCosigners(_walletId [32]byte, _cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetCosigners(&_TeeWalletManager.TransactOpts, _walletId, _cosigners, _cosignersThreshold)
}

// SetPausingAddresses is a paid mutator transaction binding the contract method 0x4fb81385.
//
// Solidity: function setPausingAddresses(bytes32 _walletId, address[] _pausingAddresses) payable returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) SetPausingAddresses(opts *bind.TransactOpts, _walletId [32]byte, _pausingAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "setPausingAddresses", _walletId, _pausingAddresses)
}

// SetPausingAddresses is a paid mutator transaction binding the contract method 0x4fb81385.
//
// Solidity: function setPausingAddresses(bytes32 _walletId, address[] _pausingAddresses) payable returns()
func (_TeeWalletManager *TeeWalletManagerSession) SetPausingAddresses(_walletId [32]byte, _pausingAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetPausingAddresses(&_TeeWalletManager.TransactOpts, _walletId, _pausingAddresses)
}

// SetPausingAddresses is a paid mutator transaction binding the contract method 0x4fb81385.
//
// Solidity: function setPausingAddresses(bytes32 _walletId, address[] _pausingAddresses) payable returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) SetPausingAddresses(_walletId [32]byte, _pausingAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetPausingAddresses(&_TeeWalletManager.TransactOpts, _walletId, _pausingAddresses)
}

// TeeWalletManagerWalletAdminConfirmedIterator is returned from FilterWalletAdminConfirmed and is used to iterate over the raw logs and unpacked data for WalletAdminConfirmed events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletAdminConfirmedIterator struct {
	Event *TeeWalletManagerWalletAdminConfirmed // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletAdminConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletAdminConfirmed)
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
		it.Event = new(TeeWalletManagerWalletAdminConfirmed)
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
func (it *TeeWalletManagerWalletAdminConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletAdminConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletAdminConfirmed represents a WalletAdminConfirmed event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletAdminConfirmed struct {
	WalletId [32]byte
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletAdminConfirmed is a free log retrieval operation binding the contract event 0xefdffb32952334eb7adb154fb07bfa263335e73f8ad33a54aba665efe346cb08.
//
// Solidity: event WalletAdminConfirmed(bytes32 indexed walletId, address indexed admin)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletAdminConfirmed(opts *bind.FilterOpts, walletId [][32]byte, admin []common.Address) (*TeeWalletManagerWalletAdminConfirmedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletAdminConfirmed", walletIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletAdminConfirmedIterator{contract: _TeeWalletManager.contract, event: "WalletAdminConfirmed", logs: logs, sub: sub}, nil
}

// WatchWalletAdminConfirmed is a free log subscription operation binding the contract event 0xefdffb32952334eb7adb154fb07bfa263335e73f8ad33a54aba665efe346cb08.
//
// Solidity: event WalletAdminConfirmed(bytes32 indexed walletId, address indexed admin)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletAdminConfirmed(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletAdminConfirmed, walletId [][32]byte, admin []common.Address) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletAdminConfirmed", walletIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletAdminConfirmed)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletAdminConfirmed", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletAdminConfirmed(log types.Log) (*TeeWalletManagerWalletAdminConfirmed, error) {
	event := new(TeeWalletManagerWalletAdminConfirmed)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletAdminConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletAdminsSetIterator is returned from FilterWalletAdminsSet and is used to iterate over the raw logs and unpacked data for WalletAdminsSet events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletAdminsSetIterator struct {
	Event *TeeWalletManagerWalletAdminsSet // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletAdminsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletAdminsSet)
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
		it.Event = new(TeeWalletManagerWalletAdminsSet)
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
func (it *TeeWalletManagerWalletAdminsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletAdminsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletAdminsSet represents a WalletAdminsSet event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletAdminsSet struct {
	WalletId         [32]byte
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWalletAdminsSet is a free log retrieval operation binding the contract event 0xfac536eb8a80dfc8ca75bf6543d4b1f435e119551bafd3733fcb28850d29d8aa.
//
// Solidity: event WalletAdminsSet(bytes32 indexed walletId, (bytes32,bytes32)[] adminsPublicKeys, uint64 adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletAdminsSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletManagerWalletAdminsSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletAdminsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletAdminsSetIterator{contract: _TeeWalletManager.contract, event: "WalletAdminsSet", logs: logs, sub: sub}, nil
}

// WatchWalletAdminsSet is a free log subscription operation binding the contract event 0xfac536eb8a80dfc8ca75bf6543d4b1f435e119551bafd3733fcb28850d29d8aa.
//
// Solidity: event WalletAdminsSet(bytes32 indexed walletId, (bytes32,bytes32)[] adminsPublicKeys, uint64 adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletAdminsSet(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletAdminsSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletAdminsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletAdminsSet)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletAdminsSet", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletAdminsSet(log types.Log) (*TeeWalletManagerWalletAdminsSet, error) {
	event := new(TeeWalletManagerWalletAdminsSet)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletAdminsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletCosignerConfirmedIterator is returned from FilterWalletCosignerConfirmed and is used to iterate over the raw logs and unpacked data for WalletCosignerConfirmed events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletCosignerConfirmedIterator struct {
	Event *TeeWalletManagerWalletCosignerConfirmed // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletCosignerConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletCosignerConfirmed)
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
		it.Event = new(TeeWalletManagerWalletCosignerConfirmed)
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
func (it *TeeWalletManagerWalletCosignerConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletCosignerConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletCosignerConfirmed represents a WalletCosignerConfirmed event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletCosignerConfirmed struct {
	WalletId [32]byte
	Cosigner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletCosignerConfirmed is a free log retrieval operation binding the contract event 0x10ab8f96b70aef2c884291a57e690b6d9e4d97ca15e4c8b2514626b8d0cd023b.
//
// Solidity: event WalletCosignerConfirmed(bytes32 indexed walletId, address indexed cosigner)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletCosignerConfirmed(opts *bind.FilterOpts, walletId [][32]byte, cosigner []common.Address) (*TeeWalletManagerWalletCosignerConfirmedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var cosignerRule []interface{}
	for _, cosignerItem := range cosigner {
		cosignerRule = append(cosignerRule, cosignerItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletCosignerConfirmed", walletIdRule, cosignerRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletCosignerConfirmedIterator{contract: _TeeWalletManager.contract, event: "WalletCosignerConfirmed", logs: logs, sub: sub}, nil
}

// WatchWalletCosignerConfirmed is a free log subscription operation binding the contract event 0x10ab8f96b70aef2c884291a57e690b6d9e4d97ca15e4c8b2514626b8d0cd023b.
//
// Solidity: event WalletCosignerConfirmed(bytes32 indexed walletId, address indexed cosigner)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletCosignerConfirmed(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletCosignerConfirmed, walletId [][32]byte, cosigner []common.Address) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var cosignerRule []interface{}
	for _, cosignerItem := range cosigner {
		cosignerRule = append(cosignerRule, cosignerItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletCosignerConfirmed", walletIdRule, cosignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletCosignerConfirmed)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletCosignerConfirmed", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletCosignerConfirmed(log types.Log) (*TeeWalletManagerWalletCosignerConfirmed, error) {
	event := new(TeeWalletManagerWalletCosignerConfirmed)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletCosignerConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletCosignersSetIterator is returned from FilterWalletCosignersSet and is used to iterate over the raw logs and unpacked data for WalletCosignersSet events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletCosignersSetIterator struct {
	Event *TeeWalletManagerWalletCosignersSet // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletCosignersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletCosignersSet)
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
		it.Event = new(TeeWalletManagerWalletCosignersSet)
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
func (it *TeeWalletManagerWalletCosignersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletCosignersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletCosignersSet represents a WalletCosignersSet event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletCosignersSet struct {
	WalletId           [32]byte
	Cosigners          []common.Address
	CosignersThreshold uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWalletCosignersSet is a free log retrieval operation binding the contract event 0x48419616e657a6824f128cbf23c09b8b4750e1c50999dc05dd869a97ac4a9c31.
//
// Solidity: event WalletCosignersSet(bytes32 indexed walletId, address[] cosigners, uint64 cosignersThreshold)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletCosignersSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletManagerWalletCosignersSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletCosignersSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletCosignersSetIterator{contract: _TeeWalletManager.contract, event: "WalletCosignersSet", logs: logs, sub: sub}, nil
}

// WatchWalletCosignersSet is a free log subscription operation binding the contract event 0x48419616e657a6824f128cbf23c09b8b4750e1c50999dc05dd869a97ac4a9c31.
//
// Solidity: event WalletCosignersSet(bytes32 indexed walletId, address[] cosigners, uint64 cosignersThreshold)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletCosignersSet(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletCosignersSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletCosignersSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletCosignersSet)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletCosignersSet", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletCosignersSet(log types.Log) (*TeeWalletManagerWalletCosignersSet, error) {
	event := new(TeeWalletManagerWalletCosignersSet)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletCosignersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletCreatedIterator is returned from FilterWalletCreated and is used to iterate over the raw logs and unpacked data for WalletCreated events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletCreatedIterator struct {
	Event *TeeWalletManagerWalletCreated // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletCreated)
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
		it.Event = new(TeeWalletManagerWalletCreated)
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
func (it *TeeWalletManagerWalletCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletCreated represents a WalletCreated event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletCreated struct {
	ProjectId [32]byte
	WalletId  [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletCreated is a free log retrieval operation binding the contract event 0xbe8f27cef1f3d94120c9c547c3614f5b992fdb0c0a497cc920fde06546291ab4.
//
// Solidity: event WalletCreated(bytes32 indexed projectId, bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletCreated(opts *bind.FilterOpts, projectId [][32]byte, walletId [][32]byte) (*TeeWalletManagerWalletCreatedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletCreated", projectIdRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletCreatedIterator{contract: _TeeWalletManager.contract, event: "WalletCreated", logs: logs, sub: sub}, nil
}

// WatchWalletCreated is a free log subscription operation binding the contract event 0xbe8f27cef1f3d94120c9c547c3614f5b992fdb0c0a497cc920fde06546291ab4.
//
// Solidity: event WalletCreated(bytes32 indexed projectId, bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletCreated(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletCreated, projectId [][32]byte, walletId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletCreated", projectIdRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletCreated)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletCreated", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletCreated(log types.Log) (*TeeWalletManagerWalletCreated, error) {
	event := new(TeeWalletManagerWalletCreated)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletEnabledIterator is returned from FilterWalletEnabled and is used to iterate over the raw logs and unpacked data for WalletEnabled events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletEnabledIterator struct {
	Event *TeeWalletManagerWalletEnabled // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletEnabled)
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
		it.Event = new(TeeWalletManagerWalletEnabled)
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
func (it *TeeWalletManagerWalletEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletEnabled represents a WalletEnabled event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletEnabled struct {
	WalletId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletEnabled is a free log retrieval operation binding the contract event 0x9dd94a650e5c2e19300b404a2e83ff44d26999ca24996bbff99dcd70f81661bf.
//
// Solidity: event WalletEnabled(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletEnabled(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletManagerWalletEnabledIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletEnabled", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletEnabledIterator{contract: _TeeWalletManager.contract, event: "WalletEnabled", logs: logs, sub: sub}, nil
}

// WatchWalletEnabled is a free log subscription operation binding the contract event 0x9dd94a650e5c2e19300b404a2e83ff44d26999ca24996bbff99dcd70f81661bf.
//
// Solidity: event WalletEnabled(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletEnabled(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletEnabled, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletEnabled", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletEnabled)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletEnabled", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletEnabled(log types.Log) (*TeeWalletManagerWalletEnabled, error) {
	event := new(TeeWalletManagerWalletEnabled)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletInitializedIterator is returned from FilterWalletInitialized and is used to iterate over the raw logs and unpacked data for WalletInitialized events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletInitializedIterator struct {
	Event *TeeWalletManagerWalletInitialized // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletInitialized)
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
		it.Event = new(TeeWalletManagerWalletInitialized)
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
func (it *TeeWalletManagerWalletInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletInitialized represents a WalletInitialized event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletInitialized struct {
	WalletId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletInitialized is a free log retrieval operation binding the contract event 0x09032cc96c5d9d0239f5528329abee7a1c72f41813ebdd51e4f51f8f962c7223.
//
// Solidity: event WalletInitialized(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletInitialized(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletManagerWalletInitializedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletInitialized", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletInitializedIterator{contract: _TeeWalletManager.contract, event: "WalletInitialized", logs: logs, sub: sub}, nil
}

// WatchWalletInitialized is a free log subscription operation binding the contract event 0x09032cc96c5d9d0239f5528329abee7a1c72f41813ebdd51e4f51f8f962c7223.
//
// Solidity: event WalletInitialized(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletInitialized(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletInitialized, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletInitialized", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletInitialized)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletInitialized", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletInitialized(log types.Log) (*TeeWalletManagerWalletInitialized, error) {
	event := new(TeeWalletManagerWalletInitialized)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletPausedIterator is returned from FilterWalletPaused and is used to iterate over the raw logs and unpacked data for WalletPaused events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletPausedIterator struct {
	Event *TeeWalletManagerWalletPaused // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletPaused)
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
		it.Event = new(TeeWalletManagerWalletPaused)
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
func (it *TeeWalletManagerWalletPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletPaused represents a WalletPaused event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletPaused struct {
	WalletId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletPaused is a free log retrieval operation binding the contract event 0x5bc3bc472409bb4331cff333acd11b5d5b6d97052d783936292f19955b25eef3.
//
// Solidity: event WalletPaused(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletPaused(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletManagerWalletPausedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletPaused", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletPausedIterator{contract: _TeeWalletManager.contract, event: "WalletPaused", logs: logs, sub: sub}, nil
}

// WatchWalletPaused is a free log subscription operation binding the contract event 0x5bc3bc472409bb4331cff333acd11b5d5b6d97052d783936292f19955b25eef3.
//
// Solidity: event WalletPaused(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletPaused(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletPaused, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletPaused", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletPaused)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletPaused", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletPaused(log types.Log) (*TeeWalletManagerWalletPaused, error) {
	event := new(TeeWalletManagerWalletPaused)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
