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

// ITeeRegistryTeeMachineWithAttestationData is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryTeeMachineWithAttestationData struct {
	TeeId    common.Address
	Owner    common.Address
	Url      string
	CodeHash [32]byte
	Platform [32]byte
}

// ITeeWalletManagerKeyCustodianBackup is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletManagerKeyCustodianBackup struct {
	TeeId               common.Address
	WalletId            [32]byte
	KeyId               *big.Int
	BackupId            *big.Int
	ShamirThreshold     *big.Int
	CustodianPublicKeys [][]byte
}

// ITeeWalletManagerKeyCustodianRestore is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletManagerKeyCustodianRestore struct {
	TeeId               common.Address
	WalletId            [32]byte
	KeyId               *big.Int
	BackupId            *big.Int
	OpType              [32]byte
	PublicKey           []byte
	CustodianPublicKeys [][]byte
}

// ITeeWalletManagerKeyDelete is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletManagerKeyDelete struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    *big.Int
}

// ITeeWalletManagerKeyGenerate is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletManagerKeyGenerate struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    *big.Int
	OpType   [32]byte
}

// ITeeWalletManagerKeyMachineBackup is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletManagerKeyMachineBackup struct {
	TeeMachine        ITeeRegistryTeeMachineWithAttestationData
	WalletId          [32]byte
	KeyId             *big.Int
	BackupId          *big.Int
	ShamirThreshold   *big.Int
	BackupTeeMachines []ITeeRegistryTeeMachineWithAttestationData
}

// ITeeWalletManagerKeyMachineBackupRemove is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletManagerKeyMachineBackupRemove struct {
	TeeIds   []common.Address
	WalletId [32]byte
	KeyId    *big.Int
	BackupId *big.Int
}

// ITeeWalletManagerKeyMachineRestore is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletManagerKeyMachineRestore struct {
	TeeMachine        ITeeRegistryTeeMachineWithAttestationData
	WalletId          [32]byte
	KeyId             *big.Int
	BackupId          *big.Int
	OpType            [32]byte
	PublicKey         []byte
	BackupTeeMachines []ITeeRegistryTeeMachineWithAttestationData
}

// WalletmanagerMetaData contains all meta data concerning the Walletmanager contract.
var WalletmanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shamirThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"custodianPublicKeys\",\"type\":\"bytes[]\"}],\"internalType\":\"structITeeWalletManager.KeyCustodianBackup\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyCustodianBackupStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backupId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"custodianPublicKeys\",\"type\":\"bytes[]\"}],\"internalType\":\"structITeeWalletManager.KeyCustodianRestore\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyCustodianRestoreStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"}],\"internalType\":\"structITeeWalletManager.KeyDelete\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyDeleteStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeWalletManager.KeyGenerate\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyGenerateStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"teeIds\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backupId\",\"type\":\"uint256\"}],\"internalType\":\"structITeeWalletManager.KeyMachineBackupRemove\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyMachineBackupRemoveStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"teeMachine\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shamirThreshold\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData[]\",\"name\":\"backupTeeMachines\",\"type\":\"tuple[]\"}],\"internalType\":\"structITeeWalletManager.KeyMachineBackup\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyMachineBackupStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"teeMachine\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backupId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData[]\",\"name\":\"backupTeeMachines\",\"type\":\"tuple[]\"}],\"internalType\":\"structITeeWalletManager.KeyMachineRestore\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyMachineRestoreStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WalletmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletmanagerMetaData.ABI instead.
var WalletmanagerABI = WalletmanagerMetaData.ABI

// Walletmanager is an auto generated Go binding around an Ethereum contract.
type Walletmanager struct {
	WalletmanagerCaller     // Read-only binding to the contract
	WalletmanagerTransactor // Write-only binding to the contract
	WalletmanagerFilterer   // Log filterer for contract events
}

// WalletmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletmanagerSession struct {
	Contract     *Walletmanager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletmanagerCallerSession struct {
	Contract *WalletmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WalletmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletmanagerTransactorSession struct {
	Contract     *WalletmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WalletmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletmanagerRaw struct {
	Contract *Walletmanager // Generic contract binding to access the raw methods on
}

// WalletmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletmanagerCallerRaw struct {
	Contract *WalletmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// WalletmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletmanagerTransactorRaw struct {
	Contract *WalletmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletmanager creates a new instance of Walletmanager, bound to a specific deployed contract.
func NewWalletmanager(address common.Address, backend bind.ContractBackend) (*Walletmanager, error) {
	contract, err := bindWalletmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Walletmanager{WalletmanagerCaller: WalletmanagerCaller{contract: contract}, WalletmanagerTransactor: WalletmanagerTransactor{contract: contract}, WalletmanagerFilterer: WalletmanagerFilterer{contract: contract}}, nil
}

// NewWalletmanagerCaller creates a new read-only instance of Walletmanager, bound to a specific deployed contract.
func NewWalletmanagerCaller(address common.Address, caller bind.ContractCaller) (*WalletmanagerCaller, error) {
	contract, err := bindWalletmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletmanagerCaller{contract: contract}, nil
}

// NewWalletmanagerTransactor creates a new write-only instance of Walletmanager, bound to a specific deployed contract.
func NewWalletmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletmanagerTransactor, error) {
	contract, err := bindWalletmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletmanagerTransactor{contract: contract}, nil
}

// NewWalletmanagerFilterer creates a new log filterer instance of Walletmanager, bound to a specific deployed contract.
func NewWalletmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletmanagerFilterer, error) {
	contract, err := bindWalletmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletmanagerFilterer{contract: contract}, nil
}

// bindWalletmanager binds a generic wrapper to an already deployed contract.
func bindWalletmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletmanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Walletmanager *WalletmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Walletmanager.Contract.WalletmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Walletmanager *WalletmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Walletmanager.Contract.WalletmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Walletmanager *WalletmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Walletmanager.Contract.WalletmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Walletmanager *WalletmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Walletmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Walletmanager *WalletmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Walletmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Walletmanager *WalletmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Walletmanager.Contract.contract.Transact(opts, method, params...)
}

// KeyCustodianBackupStruct is a paid mutator transaction binding the contract method 0xd4221c0b.
//
// Solidity: function keyCustodianBackupStruct((address,bytes32,uint256,uint256,uint256,bytes[]) ) returns()
func (_Walletmanager *WalletmanagerTransactor) KeyCustodianBackupStruct(opts *bind.TransactOpts, arg0 ITeeWalletManagerKeyCustodianBackup) (*types.Transaction, error) {
	return _Walletmanager.contract.Transact(opts, "keyCustodianBackupStruct", arg0)
}

// KeyCustodianBackupStruct is a paid mutator transaction binding the contract method 0xd4221c0b.
//
// Solidity: function keyCustodianBackupStruct((address,bytes32,uint256,uint256,uint256,bytes[]) ) returns()
func (_Walletmanager *WalletmanagerSession) KeyCustodianBackupStruct(arg0 ITeeWalletManagerKeyCustodianBackup) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyCustodianBackupStruct(&_Walletmanager.TransactOpts, arg0)
}

// KeyCustodianBackupStruct is a paid mutator transaction binding the contract method 0xd4221c0b.
//
// Solidity: function keyCustodianBackupStruct((address,bytes32,uint256,uint256,uint256,bytes[]) ) returns()
func (_Walletmanager *WalletmanagerTransactorSession) KeyCustodianBackupStruct(arg0 ITeeWalletManagerKeyCustodianBackup) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyCustodianBackupStruct(&_Walletmanager.TransactOpts, arg0)
}

// KeyCustodianRestoreStruct is a paid mutator transaction binding the contract method 0x471b37d8.
//
// Solidity: function keyCustodianRestoreStruct((address,bytes32,uint256,uint256,bytes32,bytes,bytes[]) ) returns()
func (_Walletmanager *WalletmanagerTransactor) KeyCustodianRestoreStruct(opts *bind.TransactOpts, arg0 ITeeWalletManagerKeyCustodianRestore) (*types.Transaction, error) {
	return _Walletmanager.contract.Transact(opts, "keyCustodianRestoreStruct", arg0)
}

// KeyCustodianRestoreStruct is a paid mutator transaction binding the contract method 0x471b37d8.
//
// Solidity: function keyCustodianRestoreStruct((address,bytes32,uint256,uint256,bytes32,bytes,bytes[]) ) returns()
func (_Walletmanager *WalletmanagerSession) KeyCustodianRestoreStruct(arg0 ITeeWalletManagerKeyCustodianRestore) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyCustodianRestoreStruct(&_Walletmanager.TransactOpts, arg0)
}

// KeyCustodianRestoreStruct is a paid mutator transaction binding the contract method 0x471b37d8.
//
// Solidity: function keyCustodianRestoreStruct((address,bytes32,uint256,uint256,bytes32,bytes,bytes[]) ) returns()
func (_Walletmanager *WalletmanagerTransactorSession) KeyCustodianRestoreStruct(arg0 ITeeWalletManagerKeyCustodianRestore) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyCustodianRestoreStruct(&_Walletmanager.TransactOpts, arg0)
}

// KeyDeleteStruct is a paid mutator transaction binding the contract method 0xde402cf0.
//
// Solidity: function keyDeleteStruct((address,bytes32,uint256) ) returns()
func (_Walletmanager *WalletmanagerTransactor) KeyDeleteStruct(opts *bind.TransactOpts, arg0 ITeeWalletManagerKeyDelete) (*types.Transaction, error) {
	return _Walletmanager.contract.Transact(opts, "keyDeleteStruct", arg0)
}

// KeyDeleteStruct is a paid mutator transaction binding the contract method 0xde402cf0.
//
// Solidity: function keyDeleteStruct((address,bytes32,uint256) ) returns()
func (_Walletmanager *WalletmanagerSession) KeyDeleteStruct(arg0 ITeeWalletManagerKeyDelete) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyDeleteStruct(&_Walletmanager.TransactOpts, arg0)
}

// KeyDeleteStruct is a paid mutator transaction binding the contract method 0xde402cf0.
//
// Solidity: function keyDeleteStruct((address,bytes32,uint256) ) returns()
func (_Walletmanager *WalletmanagerTransactorSession) KeyDeleteStruct(arg0 ITeeWalletManagerKeyDelete) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyDeleteStruct(&_Walletmanager.TransactOpts, arg0)
}

// KeyGenerateStruct is a paid mutator transaction binding the contract method 0xbdca988b.
//
// Solidity: function keyGenerateStruct((address,bytes32,uint256,bytes32) ) returns()
func (_Walletmanager *WalletmanagerTransactor) KeyGenerateStruct(opts *bind.TransactOpts, arg0 ITeeWalletManagerKeyGenerate) (*types.Transaction, error) {
	return _Walletmanager.contract.Transact(opts, "keyGenerateStruct", arg0)
}

// KeyGenerateStruct is a paid mutator transaction binding the contract method 0xbdca988b.
//
// Solidity: function keyGenerateStruct((address,bytes32,uint256,bytes32) ) returns()
func (_Walletmanager *WalletmanagerSession) KeyGenerateStruct(arg0 ITeeWalletManagerKeyGenerate) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyGenerateStruct(&_Walletmanager.TransactOpts, arg0)
}

// KeyGenerateStruct is a paid mutator transaction binding the contract method 0xbdca988b.
//
// Solidity: function keyGenerateStruct((address,bytes32,uint256,bytes32) ) returns()
func (_Walletmanager *WalletmanagerTransactorSession) KeyGenerateStruct(arg0 ITeeWalletManagerKeyGenerate) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyGenerateStruct(&_Walletmanager.TransactOpts, arg0)
}

// KeyMachineBackupRemoveStruct is a paid mutator transaction binding the contract method 0x0f9491fd.
//
// Solidity: function keyMachineBackupRemoveStruct((address[],bytes32,uint256,uint256) ) returns()
func (_Walletmanager *WalletmanagerTransactor) KeyMachineBackupRemoveStruct(opts *bind.TransactOpts, arg0 ITeeWalletManagerKeyMachineBackupRemove) (*types.Transaction, error) {
	return _Walletmanager.contract.Transact(opts, "keyMachineBackupRemoveStruct", arg0)
}

// KeyMachineBackupRemoveStruct is a paid mutator transaction binding the contract method 0x0f9491fd.
//
// Solidity: function keyMachineBackupRemoveStruct((address[],bytes32,uint256,uint256) ) returns()
func (_Walletmanager *WalletmanagerSession) KeyMachineBackupRemoveStruct(arg0 ITeeWalletManagerKeyMachineBackupRemove) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyMachineBackupRemoveStruct(&_Walletmanager.TransactOpts, arg0)
}

// KeyMachineBackupRemoveStruct is a paid mutator transaction binding the contract method 0x0f9491fd.
//
// Solidity: function keyMachineBackupRemoveStruct((address[],bytes32,uint256,uint256) ) returns()
func (_Walletmanager *WalletmanagerTransactorSession) KeyMachineBackupRemoveStruct(arg0 ITeeWalletManagerKeyMachineBackupRemove) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyMachineBackupRemoveStruct(&_Walletmanager.TransactOpts, arg0)
}

// KeyMachineBackupStruct is a paid mutator transaction binding the contract method 0x75dc0c12.
//
// Solidity: function keyMachineBackupStruct(((address,address,string,bytes32,bytes32),bytes32,uint256,uint256,uint256,(address,address,string,bytes32,bytes32)[]) ) returns()
func (_Walletmanager *WalletmanagerTransactor) KeyMachineBackupStruct(opts *bind.TransactOpts, arg0 ITeeWalletManagerKeyMachineBackup) (*types.Transaction, error) {
	return _Walletmanager.contract.Transact(opts, "keyMachineBackupStruct", arg0)
}

// KeyMachineBackupStruct is a paid mutator transaction binding the contract method 0x75dc0c12.
//
// Solidity: function keyMachineBackupStruct(((address,address,string,bytes32,bytes32),bytes32,uint256,uint256,uint256,(address,address,string,bytes32,bytes32)[]) ) returns()
func (_Walletmanager *WalletmanagerSession) KeyMachineBackupStruct(arg0 ITeeWalletManagerKeyMachineBackup) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyMachineBackupStruct(&_Walletmanager.TransactOpts, arg0)
}

// KeyMachineBackupStruct is a paid mutator transaction binding the contract method 0x75dc0c12.
//
// Solidity: function keyMachineBackupStruct(((address,address,string,bytes32,bytes32),bytes32,uint256,uint256,uint256,(address,address,string,bytes32,bytes32)[]) ) returns()
func (_Walletmanager *WalletmanagerTransactorSession) KeyMachineBackupStruct(arg0 ITeeWalletManagerKeyMachineBackup) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyMachineBackupStruct(&_Walletmanager.TransactOpts, arg0)
}

// KeyMachineRestoreStruct is a paid mutator transaction binding the contract method 0x1c89aa89.
//
// Solidity: function keyMachineRestoreStruct(((address,address,string,bytes32,bytes32),bytes32,uint256,uint256,bytes32,bytes,(address,address,string,bytes32,bytes32)[]) ) returns()
func (_Walletmanager *WalletmanagerTransactor) KeyMachineRestoreStruct(opts *bind.TransactOpts, arg0 ITeeWalletManagerKeyMachineRestore) (*types.Transaction, error) {
	return _Walletmanager.contract.Transact(opts, "keyMachineRestoreStruct", arg0)
}

// KeyMachineRestoreStruct is a paid mutator transaction binding the contract method 0x1c89aa89.
//
// Solidity: function keyMachineRestoreStruct(((address,address,string,bytes32,bytes32),bytes32,uint256,uint256,bytes32,bytes,(address,address,string,bytes32,bytes32)[]) ) returns()
func (_Walletmanager *WalletmanagerSession) KeyMachineRestoreStruct(arg0 ITeeWalletManagerKeyMachineRestore) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyMachineRestoreStruct(&_Walletmanager.TransactOpts, arg0)
}

// KeyMachineRestoreStruct is a paid mutator transaction binding the contract method 0x1c89aa89.
//
// Solidity: function keyMachineRestoreStruct(((address,address,string,bytes32,bytes32),bytes32,uint256,uint256,bytes32,bytes,(address,address,string,bytes32,bytes32)[]) ) returns()
func (_Walletmanager *WalletmanagerTransactorSession) KeyMachineRestoreStruct(arg0 ITeeWalletManagerKeyMachineRestore) (*types.Transaction, error) {
	return _Walletmanager.Contract.KeyMachineRestoreStruct(&_Walletmanager.TransactOpts, arg0)
}
