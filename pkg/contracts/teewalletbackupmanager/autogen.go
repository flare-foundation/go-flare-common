// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teewalletbackupmanager

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

// ITeeWalletBackupManagerBackupId is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletBackupManagerBackupId struct {
	TeeId         common.Address
	WalletId      [32]byte
	KeyId         uint64
	OpType        [32]byte
	PublicKey     []byte
	RewardEpochId *big.Int
}

// TeeWalletBackupManagerMetaData contains all meta data concerning the TeeWalletBackupManager contract.
var TeeWalletBackupManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"internalType\":\"structITeeWalletBackupManager.BackupId\",\"name\":\"_backupId\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_backupUrl\",\"type\":\"string\"}],\"name\":\"backupRestore\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeeWalletBackupManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeWalletBackupManagerMetaData.ABI instead.
var TeeWalletBackupManagerABI = TeeWalletBackupManagerMetaData.ABI

// TeeWalletBackupManager is an auto generated Go binding around an Ethereum contract.
type TeeWalletBackupManager struct {
	TeeWalletBackupManagerCaller     // Read-only binding to the contract
	TeeWalletBackupManagerTransactor // Write-only binding to the contract
	TeeWalletBackupManagerFilterer   // Log filterer for contract events
}

// TeeWalletBackupManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeWalletBackupManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletBackupManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeWalletBackupManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletBackupManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeWalletBackupManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletBackupManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeWalletBackupManagerSession struct {
	Contract     *TeeWalletBackupManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TeeWalletBackupManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeWalletBackupManagerCallerSession struct {
	Contract *TeeWalletBackupManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// TeeWalletBackupManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeWalletBackupManagerTransactorSession struct {
	Contract     *TeeWalletBackupManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// TeeWalletBackupManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeWalletBackupManagerRaw struct {
	Contract *TeeWalletBackupManager // Generic contract binding to access the raw methods on
}

// TeeWalletBackupManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeWalletBackupManagerCallerRaw struct {
	Contract *TeeWalletBackupManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TeeWalletBackupManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeWalletBackupManagerTransactorRaw struct {
	Contract *TeeWalletBackupManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeWalletBackupManager creates a new instance of TeeWalletBackupManager, bound to a specific deployed contract.
func NewTeeWalletBackupManager(address common.Address, backend bind.ContractBackend) (*TeeWalletBackupManager, error) {
	contract, err := bindTeeWalletBackupManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManager{TeeWalletBackupManagerCaller: TeeWalletBackupManagerCaller{contract: contract}, TeeWalletBackupManagerTransactor: TeeWalletBackupManagerTransactor{contract: contract}, TeeWalletBackupManagerFilterer: TeeWalletBackupManagerFilterer{contract: contract}}, nil
}

// NewTeeWalletBackupManagerCaller creates a new read-only instance of TeeWalletBackupManager, bound to a specific deployed contract.
func NewTeeWalletBackupManagerCaller(address common.Address, caller bind.ContractCaller) (*TeeWalletBackupManagerCaller, error) {
	contract, err := bindTeeWalletBackupManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManagerCaller{contract: contract}, nil
}

// NewTeeWalletBackupManagerTransactor creates a new write-only instance of TeeWalletBackupManager, bound to a specific deployed contract.
func NewTeeWalletBackupManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeWalletBackupManagerTransactor, error) {
	contract, err := bindTeeWalletBackupManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManagerTransactor{contract: contract}, nil
}

// NewTeeWalletBackupManagerFilterer creates a new log filterer instance of TeeWalletBackupManager, bound to a specific deployed contract.
func NewTeeWalletBackupManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeWalletBackupManagerFilterer, error) {
	contract, err := bindTeeWalletBackupManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManagerFilterer{contract: contract}, nil
}

// bindTeeWalletBackupManager binds a generic wrapper to an already deployed contract.
func bindTeeWalletBackupManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeWalletBackupManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletBackupManager *TeeWalletBackupManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletBackupManager.Contract.TeeWalletBackupManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletBackupManager *TeeWalletBackupManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.TeeWalletBackupManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletBackupManager *TeeWalletBackupManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.TeeWalletBackupManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletBackupManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.contract.Transact(opts, method, params...)
}

// BackupRestore is a paid mutator transaction binding the contract method 0xac097729.
//
// Solidity: function backupRestore(address _teeId, (address,bytes32,uint64,bytes32,bytes,uint24) _backupId, string _backupUrl) payable returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactor) BackupRestore(opts *bind.TransactOpts, _teeId common.Address, _backupId ITeeWalletBackupManagerBackupId, _backupUrl string) (*types.Transaction, error) {
	return _TeeWalletBackupManager.contract.Transact(opts, "backupRestore", _teeId, _backupId, _backupUrl)
}

// BackupRestore is a paid mutator transaction binding the contract method 0xac097729.
//
// Solidity: function backupRestore(address _teeId, (address,bytes32,uint64,bytes32,bytes,uint24) _backupId, string _backupUrl) payable returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) BackupRestore(_teeId common.Address, _backupId ITeeWalletBackupManagerBackupId, _backupUrl string) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.BackupRestore(&_TeeWalletBackupManager.TransactOpts, _teeId, _backupId, _backupUrl)
}

// BackupRestore is a paid mutator transaction binding the contract method 0xac097729.
//
// Solidity: function backupRestore(address _teeId, (address,bytes32,uint64,bytes32,bytes,uint24) _backupId, string _backupUrl) payable returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactorSession) BackupRestore(_teeId common.Address, _backupId ITeeWalletBackupManagerBackupId, _backupUrl string) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.BackupRestore(&_TeeWalletBackupManager.TransactOpts, _teeId, _backupId, _backupUrl)
}
