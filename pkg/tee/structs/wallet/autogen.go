// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wallet

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

// ITeeWalletBackupManagerKeyCustodianBackup is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletBackupManagerKeyCustodianBackup struct {
	TeeId               common.Address
	WalletId            [32]byte
	KeyId               *big.Int
	BackupId            *big.Int
	ShamirThreshold     *big.Int
	CustodianPublicKeys [][]byte
}

// ITeeWalletBackupManagerKeyCustodianRestore is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletBackupManagerKeyCustodianRestore struct {
	TeeId               common.Address
	WalletId            [32]byte
	KeyId               *big.Int
	BackupId            *big.Int
	OpType              [32]byte
	PublicKey           []byte
	CustodianPublicKeys [][]byte
}

// ITeeWalletBackupManagerKeyMachineBackup is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletBackupManagerKeyMachineBackup struct {
	TeeMachine        ITeeRegistryTeeMachineWithAttestationData
	WalletId          [32]byte
	KeyId             *big.Int
	BackupId          *big.Int
	ShamirThreshold   *big.Int
	BackupTeeMachines []ITeeRegistryTeeMachineWithAttestationData
}

// ITeeWalletBackupManagerKeyMachineBackupRemove is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletBackupManagerKeyMachineBackupRemove struct {
	TeeIds   []common.Address
	WalletId [32]byte
	KeyId    *big.Int
	BackupId *big.Int
}

// ITeeWalletBackupManagerKeyMachineRestore is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletBackupManagerKeyMachineRestore struct {
	TeeMachine        ITeeRegistryTeeMachineWithAttestationData
	WalletId          [32]byte
	KeyId             *big.Int
	BackupId          *big.Int
	OpType            [32]byte
	PublicKey         []byte
	BackupTeeMachines []ITeeRegistryTeeMachineWithAttestationData
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

// WalletMetaData contains all meta data concerning the Wallet contract.
var WalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shamirThreshold\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"custodianPublicKeys\",\"type\":\"bytes[]\"}],\"internalType\":\"structITeeWalletBackupManager.KeyCustodianBackup\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyCustodianBackupStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backupId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"custodianPublicKeys\",\"type\":\"bytes[]\"}],\"internalType\":\"structITeeWalletBackupManager.KeyCustodianRestore\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyCustodianRestoreStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"}],\"internalType\":\"structITeeWalletManager.KeyDelete\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyDeleteStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeWalletManager.KeyGenerate\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyGenerateStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"teeIds\",\"type\":\"address[]\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backupId\",\"type\":\"uint256\"}],\"internalType\":\"structITeeWalletBackupManager.KeyMachineBackupRemove\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyMachineBackupRemoveStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"teeMachine\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backupId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shamirThreshold\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData[]\",\"name\":\"backupTeeMachines\",\"type\":\"tuple[]\"}],\"internalType\":\"structITeeWalletBackupManager.KeyMachineBackup\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyMachineBackupStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"teeMachine\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backupId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData[]\",\"name\":\"backupTeeMachines\",\"type\":\"tuple[]\"}],\"internalType\":\"structITeeWalletBackupManager.KeyMachineRestore\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyMachineRestoreStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WalletABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletMetaData.ABI instead.
var WalletABI = WalletMetaData.ABI

// Wallet is an auto generated Go binding around an Ethereum contract.
type Wallet struct {
	WalletCaller     // Read-only binding to the contract
	WalletTransactor // Write-only binding to the contract
	WalletFilterer   // Log filterer for contract events
}

// WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSession struct {
	Contract     *Wallet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletCallerSession struct {
	Contract *WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletTransactorSession struct {
	Contract     *WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletRaw struct {
	Contract *Wallet // Generic contract binding to access the raw methods on
}

// WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletCallerRaw struct {
	Contract *WalletCaller // Generic read-only contract binding to access the raw methods on
}

// WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletTransactorRaw struct {
	Contract *WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWallet creates a new instance of Wallet, bound to a specific deployed contract.
func NewWallet(address common.Address, backend bind.ContractBackend) (*Wallet, error) {
	contract, err := bindWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// NewWalletCaller creates a new read-only instance of Wallet, bound to a specific deployed contract.
func NewWalletCaller(address common.Address, caller bind.ContractCaller) (*WalletCaller, error) {
	contract, err := bindWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCaller{contract: contract}, nil
}

// NewWalletTransactor creates a new write-only instance of Wallet, bound to a specific deployed contract.
func NewWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletTransactor, error) {
	contract, err := bindWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletTransactor{contract: contract}, nil
}

// NewWalletFilterer creates a new log filterer instance of Wallet, bound to a specific deployed contract.
func NewWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletFilterer, error) {
	contract, err := bindWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletFilterer{contract: contract}, nil
}

// bindWallet binds a generic wrapper to an already deployed contract.
func bindWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transact(opts, method, params...)
}

// KeyCustodianBackupStruct is a paid mutator transaction binding the contract method 0xd4221c0b.
//
// Solidity: function keyCustodianBackupStruct((address,bytes32,uint256,uint256,uint256,bytes[]) ) returns()
func (_Wallet *WalletTransactor) KeyCustodianBackupStruct(opts *bind.TransactOpts, arg0 ITeeWalletBackupManagerKeyCustodianBackup) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "keyCustodianBackupStruct", arg0)
}

// KeyCustodianBackupStruct is a paid mutator transaction binding the contract method 0xd4221c0b.
//
// Solidity: function keyCustodianBackupStruct((address,bytes32,uint256,uint256,uint256,bytes[]) ) returns()
func (_Wallet *WalletSession) KeyCustodianBackupStruct(arg0 ITeeWalletBackupManagerKeyCustodianBackup) (*types.Transaction, error) {
	return _Wallet.Contract.KeyCustodianBackupStruct(&_Wallet.TransactOpts, arg0)
}

// KeyCustodianBackupStruct is a paid mutator transaction binding the contract method 0xd4221c0b.
//
// Solidity: function keyCustodianBackupStruct((address,bytes32,uint256,uint256,uint256,bytes[]) ) returns()
func (_Wallet *WalletTransactorSession) KeyCustodianBackupStruct(arg0 ITeeWalletBackupManagerKeyCustodianBackup) (*types.Transaction, error) {
	return _Wallet.Contract.KeyCustodianBackupStruct(&_Wallet.TransactOpts, arg0)
}

// KeyCustodianRestoreStruct is a paid mutator transaction binding the contract method 0x471b37d8.
//
// Solidity: function keyCustodianRestoreStruct((address,bytes32,uint256,uint256,bytes32,bytes,bytes[]) ) returns()
func (_Wallet *WalletTransactor) KeyCustodianRestoreStruct(opts *bind.TransactOpts, arg0 ITeeWalletBackupManagerKeyCustodianRestore) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "keyCustodianRestoreStruct", arg0)
}

// KeyCustodianRestoreStruct is a paid mutator transaction binding the contract method 0x471b37d8.
//
// Solidity: function keyCustodianRestoreStruct((address,bytes32,uint256,uint256,bytes32,bytes,bytes[]) ) returns()
func (_Wallet *WalletSession) KeyCustodianRestoreStruct(arg0 ITeeWalletBackupManagerKeyCustodianRestore) (*types.Transaction, error) {
	return _Wallet.Contract.KeyCustodianRestoreStruct(&_Wallet.TransactOpts, arg0)
}

// KeyCustodianRestoreStruct is a paid mutator transaction binding the contract method 0x471b37d8.
//
// Solidity: function keyCustodianRestoreStruct((address,bytes32,uint256,uint256,bytes32,bytes,bytes[]) ) returns()
func (_Wallet *WalletTransactorSession) KeyCustodianRestoreStruct(arg0 ITeeWalletBackupManagerKeyCustodianRestore) (*types.Transaction, error) {
	return _Wallet.Contract.KeyCustodianRestoreStruct(&_Wallet.TransactOpts, arg0)
}

// KeyDeleteStruct is a paid mutator transaction binding the contract method 0xde402cf0.
//
// Solidity: function keyDeleteStruct((address,bytes32,uint256) ) returns()
func (_Wallet *WalletTransactor) KeyDeleteStruct(opts *bind.TransactOpts, arg0 ITeeWalletManagerKeyDelete) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "keyDeleteStruct", arg0)
}

// KeyDeleteStruct is a paid mutator transaction binding the contract method 0xde402cf0.
//
// Solidity: function keyDeleteStruct((address,bytes32,uint256) ) returns()
func (_Wallet *WalletSession) KeyDeleteStruct(arg0 ITeeWalletManagerKeyDelete) (*types.Transaction, error) {
	return _Wallet.Contract.KeyDeleteStruct(&_Wallet.TransactOpts, arg0)
}

// KeyDeleteStruct is a paid mutator transaction binding the contract method 0xde402cf0.
//
// Solidity: function keyDeleteStruct((address,bytes32,uint256) ) returns()
func (_Wallet *WalletTransactorSession) KeyDeleteStruct(arg0 ITeeWalletManagerKeyDelete) (*types.Transaction, error) {
	return _Wallet.Contract.KeyDeleteStruct(&_Wallet.TransactOpts, arg0)
}

// KeyGenerateStruct is a paid mutator transaction binding the contract method 0xbdca988b.
//
// Solidity: function keyGenerateStruct((address,bytes32,uint256,bytes32) ) returns()
func (_Wallet *WalletTransactor) KeyGenerateStruct(opts *bind.TransactOpts, arg0 ITeeWalletManagerKeyGenerate) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "keyGenerateStruct", arg0)
}

// KeyGenerateStruct is a paid mutator transaction binding the contract method 0xbdca988b.
//
// Solidity: function keyGenerateStruct((address,bytes32,uint256,bytes32) ) returns()
func (_Wallet *WalletSession) KeyGenerateStruct(arg0 ITeeWalletManagerKeyGenerate) (*types.Transaction, error) {
	return _Wallet.Contract.KeyGenerateStruct(&_Wallet.TransactOpts, arg0)
}

// KeyGenerateStruct is a paid mutator transaction binding the contract method 0xbdca988b.
//
// Solidity: function keyGenerateStruct((address,bytes32,uint256,bytes32) ) returns()
func (_Wallet *WalletTransactorSession) KeyGenerateStruct(arg0 ITeeWalletManagerKeyGenerate) (*types.Transaction, error) {
	return _Wallet.Contract.KeyGenerateStruct(&_Wallet.TransactOpts, arg0)
}

// KeyMachineBackupRemoveStruct is a paid mutator transaction binding the contract method 0x0f9491fd.
//
// Solidity: function keyMachineBackupRemoveStruct((address[],bytes32,uint256,uint256) ) returns()
func (_Wallet *WalletTransactor) KeyMachineBackupRemoveStruct(opts *bind.TransactOpts, arg0 ITeeWalletBackupManagerKeyMachineBackupRemove) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "keyMachineBackupRemoveStruct", arg0)
}

// KeyMachineBackupRemoveStruct is a paid mutator transaction binding the contract method 0x0f9491fd.
//
// Solidity: function keyMachineBackupRemoveStruct((address[],bytes32,uint256,uint256) ) returns()
func (_Wallet *WalletSession) KeyMachineBackupRemoveStruct(arg0 ITeeWalletBackupManagerKeyMachineBackupRemove) (*types.Transaction, error) {
	return _Wallet.Contract.KeyMachineBackupRemoveStruct(&_Wallet.TransactOpts, arg0)
}

// KeyMachineBackupRemoveStruct is a paid mutator transaction binding the contract method 0x0f9491fd.
//
// Solidity: function keyMachineBackupRemoveStruct((address[],bytes32,uint256,uint256) ) returns()
func (_Wallet *WalletTransactorSession) KeyMachineBackupRemoveStruct(arg0 ITeeWalletBackupManagerKeyMachineBackupRemove) (*types.Transaction, error) {
	return _Wallet.Contract.KeyMachineBackupRemoveStruct(&_Wallet.TransactOpts, arg0)
}

// KeyMachineBackupStruct is a paid mutator transaction binding the contract method 0x75dc0c12.
//
// Solidity: function keyMachineBackupStruct(((address,address,string,bytes32,bytes32),bytes32,uint256,uint256,uint256,(address,address,string,bytes32,bytes32)[]) ) returns()
func (_Wallet *WalletTransactor) KeyMachineBackupStruct(opts *bind.TransactOpts, arg0 ITeeWalletBackupManagerKeyMachineBackup) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "keyMachineBackupStruct", arg0)
}

// KeyMachineBackupStruct is a paid mutator transaction binding the contract method 0x75dc0c12.
//
// Solidity: function keyMachineBackupStruct(((address,address,string,bytes32,bytes32),bytes32,uint256,uint256,uint256,(address,address,string,bytes32,bytes32)[]) ) returns()
func (_Wallet *WalletSession) KeyMachineBackupStruct(arg0 ITeeWalletBackupManagerKeyMachineBackup) (*types.Transaction, error) {
	return _Wallet.Contract.KeyMachineBackupStruct(&_Wallet.TransactOpts, arg0)
}

// KeyMachineBackupStruct is a paid mutator transaction binding the contract method 0x75dc0c12.
//
// Solidity: function keyMachineBackupStruct(((address,address,string,bytes32,bytes32),bytes32,uint256,uint256,uint256,(address,address,string,bytes32,bytes32)[]) ) returns()
func (_Wallet *WalletTransactorSession) KeyMachineBackupStruct(arg0 ITeeWalletBackupManagerKeyMachineBackup) (*types.Transaction, error) {
	return _Wallet.Contract.KeyMachineBackupStruct(&_Wallet.TransactOpts, arg0)
}

// KeyMachineRestoreStruct is a paid mutator transaction binding the contract method 0x1c89aa89.
//
// Solidity: function keyMachineRestoreStruct(((address,address,string,bytes32,bytes32),bytes32,uint256,uint256,bytes32,bytes,(address,address,string,bytes32,bytes32)[]) ) returns()
func (_Wallet *WalletTransactor) KeyMachineRestoreStruct(opts *bind.TransactOpts, arg0 ITeeWalletBackupManagerKeyMachineRestore) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "keyMachineRestoreStruct", arg0)
}

// KeyMachineRestoreStruct is a paid mutator transaction binding the contract method 0x1c89aa89.
//
// Solidity: function keyMachineRestoreStruct(((address,address,string,bytes32,bytes32),bytes32,uint256,uint256,bytes32,bytes,(address,address,string,bytes32,bytes32)[]) ) returns()
func (_Wallet *WalletSession) KeyMachineRestoreStruct(arg0 ITeeWalletBackupManagerKeyMachineRestore) (*types.Transaction, error) {
	return _Wallet.Contract.KeyMachineRestoreStruct(&_Wallet.TransactOpts, arg0)
}

// KeyMachineRestoreStruct is a paid mutator transaction binding the contract method 0x1c89aa89.
//
// Solidity: function keyMachineRestoreStruct(((address,address,string,bytes32,bytes32),bytes32,uint256,uint256,bytes32,bytes,(address,address,string,bytes32,bytes32)[]) ) returns()
func (_Wallet *WalletTransactorSession) KeyMachineRestoreStruct(arg0 ITeeWalletBackupManagerKeyMachineRestore) (*types.Transaction, error) {
	return _Wallet.Contract.KeyMachineRestoreStruct(&_Wallet.TransactOpts, arg0)
}
