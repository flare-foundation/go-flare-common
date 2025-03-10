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

// ITeeRegistryTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryTeeMachine struct {
	TeeId common.Address
	Owner common.Address
	Url   string
}

// TeeWalletManagerMetaData contains all meta data concerning the TeeWalletManager contract.
var TeeWalletManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"name\":\"WalletCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"addKey\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"cleanUpTeeIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"deleteKey\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"enableWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getFeeFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedOpTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_supportedOpTypes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletBackupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_backupManager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_submitAddress\",\"type\":\"address\"},{\"internalType\":\"enumITeeWalletManager.WalletStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"getWalletKeyPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"getWalletKeyTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletKeysInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_multisigThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint256[]\",\"name\":\"_keyIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"_counter\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletOpType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_walletOwner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletStatus\",\"outputs\":[{\"internalType\":\"enumITeeWalletManager.WalletStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_multisigThreshold\",\"type\":\"uint64\"}],\"name\":\"initializeWallet\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"pauseWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"receivingTees\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structITeeRegistry.TeeMachine[]\",\"name\":\"_receivingTees\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_submitAddress\",\"type\":\"address\"}],\"name\":\"setSubmitAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_backupManager\",\"type\":\"address\"}],\"name\":\"setWalletBackupManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetFeeFactor is a free data retrieval call binding the contract method 0x5220697c.
//
// Solidity: function getFeeFactor(bytes32 _walletId) view returns(uint256 _feeFactor)
func (_TeeWalletManager *TeeWalletManagerCaller) GetFeeFactor(opts *bind.CallOpts, _walletId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getFeeFactor", _walletId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeFactor is a free data retrieval call binding the contract method 0x5220697c.
//
// Solidity: function getFeeFactor(bytes32 _walletId) view returns(uint256 _feeFactor)
func (_TeeWalletManager *TeeWalletManagerSession) GetFeeFactor(_walletId [32]byte) (*big.Int, error) {
	return _TeeWalletManager.Contract.GetFeeFactor(&_TeeWalletManager.CallOpts, _walletId)
}

// GetFeeFactor is a free data retrieval call binding the contract method 0x5220697c.
//
// Solidity: function getFeeFactor(bytes32 _walletId) view returns(uint256 _feeFactor)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetFeeFactor(_walletId [32]byte) (*big.Int, error) {
	return _TeeWalletManager.Contract.GetFeeFactor(&_TeeWalletManager.CallOpts, _walletId)
}

// GetSupportedOpTypes is a free data retrieval call binding the contract method 0x92b7ef79.
//
// Solidity: function getSupportedOpTypes() view returns(bytes32[] _supportedOpTypes)
func (_TeeWalletManager *TeeWalletManagerCaller) GetSupportedOpTypes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getSupportedOpTypes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSupportedOpTypes is a free data retrieval call binding the contract method 0x92b7ef79.
//
// Solidity: function getSupportedOpTypes() view returns(bytes32[] _supportedOpTypes)
func (_TeeWalletManager *TeeWalletManagerSession) GetSupportedOpTypes() ([][32]byte, error) {
	return _TeeWalletManager.Contract.GetSupportedOpTypes(&_TeeWalletManager.CallOpts)
}

// GetSupportedOpTypes is a free data retrieval call binding the contract method 0x92b7ef79.
//
// Solidity: function getSupportedOpTypes() view returns(bytes32[] _supportedOpTypes)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetSupportedOpTypes() ([][32]byte, error) {
	return _TeeWalletManager.Contract.GetSupportedOpTypes(&_TeeWalletManager.CallOpts)
}

// GetWalletBackupManager is a free data retrieval call binding the contract method 0x6e3484a7.
//
// Solidity: function getWalletBackupManager(bytes32 _walletId) view returns(address _backupManager)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletBackupManager(opts *bind.CallOpts, _walletId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletBackupManager", _walletId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWalletBackupManager is a free data retrieval call binding the contract method 0x6e3484a7.
//
// Solidity: function getWalletBackupManager(bytes32 _walletId) view returns(address _backupManager)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletBackupManager(_walletId [32]byte) (common.Address, error) {
	return _TeeWalletManager.Contract.GetWalletBackupManager(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletBackupManager is a free data retrieval call binding the contract method 0x6e3484a7.
//
// Solidity: function getWalletBackupManager(bytes32 _walletId) view returns(address _backupManager)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletBackupManager(_walletId [32]byte) (common.Address, error) {
	return _TeeWalletManager.Contract.GetWalletBackupManager(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletInfo is a free data retrieval call binding the contract method 0x3bb50d7c.
//
// Solidity: function getWalletInfo(bytes32 _walletId) view returns(address _submitAddress, uint8 _status, bytes32 _opType)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletInfo(opts *bind.CallOpts, _walletId [32]byte) (struct {
	SubmitAddress common.Address
	Status        uint8
	OpType        [32]byte
}, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletInfo", _walletId)

	outstruct := new(struct {
		SubmitAddress common.Address
		Status        uint8
		OpType        [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SubmitAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.OpType = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// GetWalletInfo is a free data retrieval call binding the contract method 0x3bb50d7c.
//
// Solidity: function getWalletInfo(bytes32 _walletId) view returns(address _submitAddress, uint8 _status, bytes32 _opType)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletInfo(_walletId [32]byte) (struct {
	SubmitAddress common.Address
	Status        uint8
	OpType        [32]byte
}, error) {
	return _TeeWalletManager.Contract.GetWalletInfo(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletInfo is a free data retrieval call binding the contract method 0x3bb50d7c.
//
// Solidity: function getWalletInfo(bytes32 _walletId) view returns(address _submitAddress, uint8 _status, bytes32 _opType)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletInfo(_walletId [32]byte) (struct {
	SubmitAddress common.Address
	Status        uint8
	OpType        [32]byte
}, error) {
	return _TeeWalletManager.Contract.GetWalletInfo(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0xbb825ac3.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletKeyPublicKey(opts *bind.CallOpts, _walletId [32]byte, _keyId uint64) ([]byte, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletKeyPublicKey", _walletId, _keyId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0xbb825ac3.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletKeyPublicKey(_walletId [32]byte, _keyId uint64) ([]byte, error) {
	return _TeeWalletManager.Contract.GetWalletKeyPublicKey(&_TeeWalletManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0xbb825ac3.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletKeyPublicKey(_walletId [32]byte, _keyId uint64) ([]byte, error) {
	return _TeeWalletManager.Contract.GetWalletKeyPublicKey(&_TeeWalletManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x08480213.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletKeyTeeIds(opts *bind.CallOpts, _walletId [32]byte, _keyId uint64) ([]common.Address, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletKeyTeeIds", _walletId, _keyId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x08480213.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletKeyTeeIds(_walletId [32]byte, _keyId uint64) ([]common.Address, error) {
	return _TeeWalletManager.Contract.GetWalletKeyTeeIds(&_TeeWalletManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x08480213.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletKeyTeeIds(_walletId [32]byte, _keyId uint64) ([]common.Address, error) {
	return _TeeWalletManager.Contract.GetWalletKeyTeeIds(&_TeeWalletManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint256[] _keyIds, uint64 _counter)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletKeysInfo(opts *bind.CallOpts, _walletId [32]byte) (struct {
	MultisigThreshold uint64
	KeyIds            []*big.Int
	Counter           uint64
}, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletKeysInfo", _walletId)

	outstruct := new(struct {
		MultisigThreshold uint64
		KeyIds            []*big.Int
		Counter           uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MultisigThreshold = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.KeyIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Counter = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint256[] _keyIds, uint64 _counter)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletKeysInfo(_walletId [32]byte) (struct {
	MultisigThreshold uint64
	KeyIds            []*big.Int
	Counter           uint64
}, error) {
	return _TeeWalletManager.Contract.GetWalletKeysInfo(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint256[] _keyIds, uint64 _counter)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletKeysInfo(_walletId [32]byte) (struct {
	MultisigThreshold uint64
	KeyIds            []*big.Int
	Counter           uint64
}, error) {
	return _TeeWalletManager.Contract.GetWalletKeysInfo(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletOpType is a free data retrieval call binding the contract method 0x0b09121f.
//
// Solidity: function getWalletOpType(bytes32 _walletId) view returns(bytes32 _opType)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletOpType(opts *bind.CallOpts, _walletId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletOpType", _walletId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWalletOpType is a free data retrieval call binding the contract method 0x0b09121f.
//
// Solidity: function getWalletOpType(bytes32 _walletId) view returns(bytes32 _opType)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletOpType(_walletId [32]byte) ([32]byte, error) {
	return _TeeWalletManager.Contract.GetWalletOpType(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletOpType is a free data retrieval call binding the contract method 0x0b09121f.
//
// Solidity: function getWalletOpType(bytes32 _walletId) view returns(bytes32 _opType)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletOpType(_walletId [32]byte) ([32]byte, error) {
	return _TeeWalletManager.Contract.GetWalletOpType(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletOwner is a free data retrieval call binding the contract method 0xc8ebd47a.
//
// Solidity: function getWalletOwner(bytes32 _walletId) view returns(address _walletOwner)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletOwner(opts *bind.CallOpts, _walletId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletOwner", _walletId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWalletOwner is a free data retrieval call binding the contract method 0xc8ebd47a.
//
// Solidity: function getWalletOwner(bytes32 _walletId) view returns(address _walletOwner)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletOwner(_walletId [32]byte) (common.Address, error) {
	return _TeeWalletManager.Contract.GetWalletOwner(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletOwner is a free data retrieval call binding the contract method 0xc8ebd47a.
//
// Solidity: function getWalletOwner(bytes32 _walletId) view returns(address _walletOwner)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletOwner(_walletId [32]byte) (common.Address, error) {
	return _TeeWalletManager.Contract.GetWalletOwner(&_TeeWalletManager.CallOpts, _walletId)
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

// ReceivingTees is a free data retrieval call binding the contract method 0x0fb83e84.
//
// Solidity: function receivingTees(bytes32 _walletId) view returns((address,address,string)[] _receivingTees)
func (_TeeWalletManager *TeeWalletManagerCaller) ReceivingTees(opts *bind.CallOpts, _walletId [32]byte) ([]ITeeRegistryTeeMachine, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "receivingTees", _walletId)

	if err != nil {
		return *new([]ITeeRegistryTeeMachine), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITeeRegistryTeeMachine)).(*[]ITeeRegistryTeeMachine)

	return out0, err

}

// ReceivingTees is a free data retrieval call binding the contract method 0x0fb83e84.
//
// Solidity: function receivingTees(bytes32 _walletId) view returns((address,address,string)[] _receivingTees)
func (_TeeWalletManager *TeeWalletManagerSession) ReceivingTees(_walletId [32]byte) ([]ITeeRegistryTeeMachine, error) {
	return _TeeWalletManager.Contract.ReceivingTees(&_TeeWalletManager.CallOpts, _walletId)
}

// ReceivingTees is a free data retrieval call binding the contract method 0x0fb83e84.
//
// Solidity: function receivingTees(bytes32 _walletId) view returns((address,address,string)[] _receivingTees)
func (_TeeWalletManager *TeeWalletManagerCallerSession) ReceivingTees(_walletId [32]byte) ([]ITeeRegistryTeeMachine, error) {
	return _TeeWalletManager.Contract.ReceivingTees(&_TeeWalletManager.CallOpts, _walletId)
}

// AddKey is a paid mutator transaction binding the contract method 0x7875fa5c.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId) payable returns(uint64 _keyId)
func (_TeeWalletManager *TeeWalletManagerTransactor) AddKey(opts *bind.TransactOpts, _teeId common.Address, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "addKey", _teeId, _walletId)
}

// AddKey is a paid mutator transaction binding the contract method 0x7875fa5c.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId) payable returns(uint64 _keyId)
func (_TeeWalletManager *TeeWalletManagerSession) AddKey(_teeId common.Address, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.AddKey(&_TeeWalletManager.TransactOpts, _teeId, _walletId)
}

// AddKey is a paid mutator transaction binding the contract method 0x7875fa5c.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId) payable returns(uint64 _keyId)
func (_TeeWalletManager *TeeWalletManagerTransactorSession) AddKey(_teeId common.Address, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.AddKey(&_TeeWalletManager.TransactOpts, _teeId, _walletId)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0xaec4c979.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) CleanUpTeeIds(opts *bind.TransactOpts, _walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "cleanUpTeeIds", _walletId, _keyId)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0xaec4c979.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (_TeeWalletManager *TeeWalletManagerSession) CleanUpTeeIds(_walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.CleanUpTeeIds(&_TeeWalletManager.TransactOpts, _walletId, _keyId)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0xaec4c979.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) CleanUpTeeIds(_walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.CleanUpTeeIds(&_TeeWalletManager.TransactOpts, _walletId, _keyId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x7ff0e872.
//
// Solidity: function confirmOwnership(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) ConfirmOwnership(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "confirmOwnership", _walletId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x7ff0e872.
//
// Solidity: function confirmOwnership(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerSession) ConfirmOwnership(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ConfirmOwnership(&_TeeWalletManager.TransactOpts, _walletId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x7ff0e872.
//
// Solidity: function confirmOwnership(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) ConfirmOwnership(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ConfirmOwnership(&_TeeWalletManager.TransactOpts, _walletId)
}

// DeleteKey is a paid mutator transaction binding the contract method 0x90e31e7e.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId) payable returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) DeleteKey(opts *bind.TransactOpts, _teeId common.Address, _walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "deleteKey", _teeId, _walletId, _keyId)
}

// DeleteKey is a paid mutator transaction binding the contract method 0x90e31e7e.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId) payable returns()
func (_TeeWalletManager *TeeWalletManagerSession) DeleteKey(_teeId common.Address, _walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.DeleteKey(&_TeeWalletManager.TransactOpts, _teeId, _walletId, _keyId)
}

// DeleteKey is a paid mutator transaction binding the contract method 0x90e31e7e.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId) payable returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) DeleteKey(_teeId common.Address, _walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.DeleteKey(&_TeeWalletManager.TransactOpts, _teeId, _walletId, _keyId)
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

// InitializeWallet is a paid mutator transaction binding the contract method 0xdeb39610.
//
// Solidity: function initializeWallet(bytes32 _opType, uint64 _multisigThreshold) payable returns(bytes32 _walletId)
func (_TeeWalletManager *TeeWalletManagerTransactor) InitializeWallet(opts *bind.TransactOpts, _opType [32]byte, _multisigThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "initializeWallet", _opType, _multisigThreshold)
}

// InitializeWallet is a paid mutator transaction binding the contract method 0xdeb39610.
//
// Solidity: function initializeWallet(bytes32 _opType, uint64 _multisigThreshold) payable returns(bytes32 _walletId)
func (_TeeWalletManager *TeeWalletManagerSession) InitializeWallet(_opType [32]byte, _multisigThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.InitializeWallet(&_TeeWalletManager.TransactOpts, _opType, _multisigThreshold)
}

// InitializeWallet is a paid mutator transaction binding the contract method 0xdeb39610.
//
// Solidity: function initializeWallet(bytes32 _opType, uint64 _multisigThreshold) payable returns(bytes32 _walletId)
func (_TeeWalletManager *TeeWalletManagerTransactorSession) InitializeWallet(_opType [32]byte, _multisigThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.InitializeWallet(&_TeeWalletManager.TransactOpts, _opType, _multisigThreshold)
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

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x3ad60dd6.
//
// Solidity: function proposeNewOwner(bytes32 _walletId, address _newOwner) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) ProposeNewOwner(opts *bind.TransactOpts, _walletId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "proposeNewOwner", _walletId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x3ad60dd6.
//
// Solidity: function proposeNewOwner(bytes32 _walletId, address _newOwner) returns()
func (_TeeWalletManager *TeeWalletManagerSession) ProposeNewOwner(_walletId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ProposeNewOwner(&_TeeWalletManager.TransactOpts, _walletId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x3ad60dd6.
//
// Solidity: function proposeNewOwner(bytes32 _walletId, address _newOwner) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) ProposeNewOwner(_walletId [32]byte, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ProposeNewOwner(&_TeeWalletManager.TransactOpts, _walletId, _newOwner)
}

// SetSubmitAddress is a paid mutator transaction binding the contract method 0x9f73ece8.
//
// Solidity: function setSubmitAddress(bytes32 _walletId, address _submitAddress) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) SetSubmitAddress(opts *bind.TransactOpts, _walletId [32]byte, _submitAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "setSubmitAddress", _walletId, _submitAddress)
}

// SetSubmitAddress is a paid mutator transaction binding the contract method 0x9f73ece8.
//
// Solidity: function setSubmitAddress(bytes32 _walletId, address _submitAddress) returns()
func (_TeeWalletManager *TeeWalletManagerSession) SetSubmitAddress(_walletId [32]byte, _submitAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetSubmitAddress(&_TeeWalletManager.TransactOpts, _walletId, _submitAddress)
}

// SetSubmitAddress is a paid mutator transaction binding the contract method 0x9f73ece8.
//
// Solidity: function setSubmitAddress(bytes32 _walletId, address _submitAddress) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) SetSubmitAddress(_walletId [32]byte, _submitAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetSubmitAddress(&_TeeWalletManager.TransactOpts, _walletId, _submitAddress)
}

// SetWalletBackupManager is a paid mutator transaction binding the contract method 0xb6bb79fc.
//
// Solidity: function setWalletBackupManager(bytes32 _walletId, address _backupManager) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) SetWalletBackupManager(opts *bind.TransactOpts, _walletId [32]byte, _backupManager common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "setWalletBackupManager", _walletId, _backupManager)
}

// SetWalletBackupManager is a paid mutator transaction binding the contract method 0xb6bb79fc.
//
// Solidity: function setWalletBackupManager(bytes32 _walletId, address _backupManager) returns()
func (_TeeWalletManager *TeeWalletManagerSession) SetWalletBackupManager(_walletId [32]byte, _backupManager common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetWalletBackupManager(&_TeeWalletManager.TransactOpts, _walletId, _backupManager)
}

// SetWalletBackupManager is a paid mutator transaction binding the contract method 0xb6bb79fc.
//
// Solidity: function setWalletBackupManager(bytes32 _walletId, address _backupManager) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) SetWalletBackupManager(_walletId [32]byte, _backupManager common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetWalletBackupManager(&_TeeWalletManager.TransactOpts, _walletId, _backupManager)
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
	WalletId [32]byte
	Owner    common.Address
	OpType   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletCreated is a free log retrieval operation binding the contract event 0xac5b1cdce61c84ea0a95f1a97d5107f0b5fd19743e6400bce9e8a822fb2bc4ba.
//
// Solidity: event WalletCreated(bytes32 indexed walletId, address indexed owner, bytes32 opType)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletCreated(opts *bind.FilterOpts, walletId [][32]byte, owner []common.Address) (*TeeWalletManagerWalletCreatedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletCreated", walletIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletCreatedIterator{contract: _TeeWalletManager.contract, event: "WalletCreated", logs: logs, sub: sub}, nil
}

// WatchWalletCreated is a free log subscription operation binding the contract event 0xac5b1cdce61c84ea0a95f1a97d5107f0b5fd19743e6400bce9e8a822fb2bc4ba.
//
// Solidity: event WalletCreated(bytes32 indexed walletId, address indexed owner, bytes32 opType)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletCreated(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletCreated, walletId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletCreated", walletIdRule, ownerRule)
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

// ParseWalletCreated is a log parse operation binding the contract event 0xac5b1cdce61c84ea0a95f1a97d5107f0b5fd19743e6400bce9e8a822fb2bc4ba.
//
// Solidity: event WalletCreated(bytes32 indexed walletId, address indexed owner, bytes32 opType)
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletCreated(log types.Log) (*TeeWalletManagerWalletCreated, error) {
	event := new(TeeWalletManagerWalletCreated)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
