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

// IWalletBackupManagerBackupId is an auto generated low-level Go binding around an user-defined struct.
type IWalletBackupManagerBackupId struct {
	TeeId         common.Address
	WalletId      [32]byte
	KeyId         uint64
	KeyType       [32]byte
	SigningAlgo   [32]byte
	PublicKey     []byte
	RewardEpochId uint32
	RandomNonce   [32]byte
}

// IWalletBackupManagerKeyDataProviderRestore is an auto generated low-level Go binding around an user-defined struct.
type IWalletBackupManagerKeyDataProviderRestore struct {
	TeePublicKey PublicKey
	BackupId     IWalletBackupManagerBackupId
	BackupUrl    string
	Nonce        *big.Int
}

// IWalletKeyManagerKeyConfigConstants is an auto generated low-level Go binding around an user-defined struct.
type IWalletKeyManagerKeyConfigConstants struct {
	AdminsPublicKeys   []PublicKey
	AdminsThreshold    uint64
	Cosigners          []common.Address
	CosignersThreshold uint64
}

// IWalletKeyManagerKeyDelete is an auto generated low-level Go binding around an user-defined struct.
type IWalletKeyManagerKeyDelete struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    uint64
	Nonce    *big.Int
}

// IWalletKeyManagerKeyExistence is an auto generated low-level Go binding around an user-defined struct.
type IWalletKeyManagerKeyExistence struct {
	TeeId           common.Address
	WalletId        [32]byte
	KeyId           uint64
	KeyType         [32]byte
	SigningAlgo     [32]byte
	PublicKey       []byte
	Nonce           *big.Int
	Restored        bool
	ConfigConstants IWalletKeyManagerKeyConfigConstants
	SettingsVersion [32]byte
	Settings        []byte
}

// IWalletKeyManagerKeyGenerate is an auto generated low-level Go binding around an user-defined struct.
type IWalletKeyManagerKeyGenerate struct {
	TeeId           common.Address
	WalletId        [32]byte
	KeyId           uint64
	KeyType         [32]byte
	SigningAlgo     [32]byte
	ConfigConstants IWalletKeyManagerKeyConfigConstants
}

// IWalletResumeResume is an auto generated low-level Go binding around an user-defined struct.
type IWalletResumeResume struct {
	WalletId [32]byte
	KeysData []IWalletResumeResumeKeyData
}

// IWalletResumeResumeKeyData is an auto generated low-level Go binding around an user-defined struct.
type IWalletResumeResumeKeyData struct {
	KeyId uint64
	TeeId common.Address
	Nonce *big.Int
}

// IWalletResumeSetPausingAddresses is an auto generated low-level Go binding around an user-defined struct.
type IWalletResumeSetPausingAddresses struct {
	WalletId         [32]byte
	Nonce            *big.Int
	TeeIdKeyIdPairs  []TeeIdKeyIdPair
	PausingAddresses []common.Address
}

// PublicKey is an auto generated low-level Go binding around an user-defined struct.
type PublicKey struct {
	X [32]byte
	Y [32]byte
}

// TeeIdKeyIdPair is an auto generated low-level Go binding around an user-defined struct.
type TeeIdKeyIdPair struct {
	TeeId common.Address
	KeyId uint64
}

// WalletMetaData contains all meta data concerning the Wallet contract.
var WalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"randomNonce\",\"type\":\"bytes32\"}],\"internalType\":\"structIWalletBackupManager.BackupId\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"backupIdStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"adminsThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"internalType\":\"structIWalletKeyManager.KeyConfigConstants\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyConfigConstantsStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"teePublicKey\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"randomNonce\",\"type\":\"bytes32\"}],\"internalType\":\"structIWalletBackupManager.BackupId\",\"name\":\"backupId\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"backupUrl\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIWalletBackupManager.KeyDataProviderRestore\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyDataProviderRestoreStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIWalletKeyManager.KeyDelete\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyDeleteStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restored\",\"type\":\"bool\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"adminsThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"internalType\":\"structIWalletKeyManager.KeyConfigConstants\",\"name\":\"configConstants\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"settingsVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"settings\",\"type\":\"bytes\"}],\"internalType\":\"structIWalletKeyManager.KeyExistence\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyExistenceStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"adminsThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"internalType\":\"structIWalletKeyManager.KeyConfigConstants\",\"name\":\"configConstants\",\"type\":\"tuple\"}],\"internalType\":\"structIWalletKeyManager.KeyGenerate\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"keyGenerateStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIWalletResume.ResumeKeyData[]\",\"name\":\"keysData\",\"type\":\"tuple[]\"}],\"internalType\":\"structIWalletResume.Resume\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"resumeStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"internalType\":\"structTeeIdKeyIdPair[]\",\"name\":\"teeIdKeyIdPairs\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"pausingAddresses\",\"type\":\"address[]\"}],\"internalType\":\"structIWalletResume.SetPausingAddresses\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"setPausingAddressesStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// BackupIdStruct is a paid mutator transaction binding the contract method 0x61331536.
//
// Solidity: function backupIdStruct((address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32) ) returns()
func (_Wallet *WalletTransactor) BackupIdStruct(opts *bind.TransactOpts, arg0 IWalletBackupManagerBackupId) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "backupIdStruct", arg0)
}

// BackupIdStruct is a paid mutator transaction binding the contract method 0x61331536.
//
// Solidity: function backupIdStruct((address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32) ) returns()
func (_Wallet *WalletSession) BackupIdStruct(arg0 IWalletBackupManagerBackupId) (*types.Transaction, error) {
	return _Wallet.Contract.BackupIdStruct(&_Wallet.TransactOpts, arg0)
}

// BackupIdStruct is a paid mutator transaction binding the contract method 0x61331536.
//
// Solidity: function backupIdStruct((address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32) ) returns()
func (_Wallet *WalletTransactorSession) BackupIdStruct(arg0 IWalletBackupManagerBackupId) (*types.Transaction, error) {
	return _Wallet.Contract.BackupIdStruct(&_Wallet.TransactOpts, arg0)
}

// KeyConfigConstantsStruct is a paid mutator transaction binding the contract method 0x979cc96b.
//
// Solidity: function keyConfigConstantsStruct(((bytes32,bytes32)[],uint64,address[],uint64) ) returns()
func (_Wallet *WalletTransactor) KeyConfigConstantsStruct(opts *bind.TransactOpts, arg0 IWalletKeyManagerKeyConfigConstants) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "keyConfigConstantsStruct", arg0)
}

// KeyConfigConstantsStruct is a paid mutator transaction binding the contract method 0x979cc96b.
//
// Solidity: function keyConfigConstantsStruct(((bytes32,bytes32)[],uint64,address[],uint64) ) returns()
func (_Wallet *WalletSession) KeyConfigConstantsStruct(arg0 IWalletKeyManagerKeyConfigConstants) (*types.Transaction, error) {
	return _Wallet.Contract.KeyConfigConstantsStruct(&_Wallet.TransactOpts, arg0)
}

// KeyConfigConstantsStruct is a paid mutator transaction binding the contract method 0x979cc96b.
//
// Solidity: function keyConfigConstantsStruct(((bytes32,bytes32)[],uint64,address[],uint64) ) returns()
func (_Wallet *WalletTransactorSession) KeyConfigConstantsStruct(arg0 IWalletKeyManagerKeyConfigConstants) (*types.Transaction, error) {
	return _Wallet.Contract.KeyConfigConstantsStruct(&_Wallet.TransactOpts, arg0)
}

// KeyDataProviderRestoreStruct is a paid mutator transaction binding the contract method 0xf3a00746.
//
// Solidity: function keyDataProviderRestoreStruct(((bytes32,bytes32),(address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32),string,uint256) ) returns()
func (_Wallet *WalletTransactor) KeyDataProviderRestoreStruct(opts *bind.TransactOpts, arg0 IWalletBackupManagerKeyDataProviderRestore) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "keyDataProviderRestoreStruct", arg0)
}

// KeyDataProviderRestoreStruct is a paid mutator transaction binding the contract method 0xf3a00746.
//
// Solidity: function keyDataProviderRestoreStruct(((bytes32,bytes32),(address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32),string,uint256) ) returns()
func (_Wallet *WalletSession) KeyDataProviderRestoreStruct(arg0 IWalletBackupManagerKeyDataProviderRestore) (*types.Transaction, error) {
	return _Wallet.Contract.KeyDataProviderRestoreStruct(&_Wallet.TransactOpts, arg0)
}

// KeyDataProviderRestoreStruct is a paid mutator transaction binding the contract method 0xf3a00746.
//
// Solidity: function keyDataProviderRestoreStruct(((bytes32,bytes32),(address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32),string,uint256) ) returns()
func (_Wallet *WalletTransactorSession) KeyDataProviderRestoreStruct(arg0 IWalletBackupManagerKeyDataProviderRestore) (*types.Transaction, error) {
	return _Wallet.Contract.KeyDataProviderRestoreStruct(&_Wallet.TransactOpts, arg0)
}

// KeyDeleteStruct is a paid mutator transaction binding the contract method 0x1dd54316.
//
// Solidity: function keyDeleteStruct((address,bytes32,uint64,uint256) ) returns()
func (_Wallet *WalletTransactor) KeyDeleteStruct(opts *bind.TransactOpts, arg0 IWalletKeyManagerKeyDelete) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "keyDeleteStruct", arg0)
}

// KeyDeleteStruct is a paid mutator transaction binding the contract method 0x1dd54316.
//
// Solidity: function keyDeleteStruct((address,bytes32,uint64,uint256) ) returns()
func (_Wallet *WalletSession) KeyDeleteStruct(arg0 IWalletKeyManagerKeyDelete) (*types.Transaction, error) {
	return _Wallet.Contract.KeyDeleteStruct(&_Wallet.TransactOpts, arg0)
}

// KeyDeleteStruct is a paid mutator transaction binding the contract method 0x1dd54316.
//
// Solidity: function keyDeleteStruct((address,bytes32,uint64,uint256) ) returns()
func (_Wallet *WalletTransactorSession) KeyDeleteStruct(arg0 IWalletKeyManagerKeyDelete) (*types.Transaction, error) {
	return _Wallet.Contract.KeyDeleteStruct(&_Wallet.TransactOpts, arg0)
}

// KeyExistenceStruct is a paid mutator transaction binding the contract method 0x641ca408.
//
// Solidity: function keyExistenceStruct((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) ) returns()
func (_Wallet *WalletTransactor) KeyExistenceStruct(opts *bind.TransactOpts, arg0 IWalletKeyManagerKeyExistence) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "keyExistenceStruct", arg0)
}

// KeyExistenceStruct is a paid mutator transaction binding the contract method 0x641ca408.
//
// Solidity: function keyExistenceStruct((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) ) returns()
func (_Wallet *WalletSession) KeyExistenceStruct(arg0 IWalletKeyManagerKeyExistence) (*types.Transaction, error) {
	return _Wallet.Contract.KeyExistenceStruct(&_Wallet.TransactOpts, arg0)
}

// KeyExistenceStruct is a paid mutator transaction binding the contract method 0x641ca408.
//
// Solidity: function keyExistenceStruct((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) ) returns()
func (_Wallet *WalletTransactorSession) KeyExistenceStruct(arg0 IWalletKeyManagerKeyExistence) (*types.Transaction, error) {
	return _Wallet.Contract.KeyExistenceStruct(&_Wallet.TransactOpts, arg0)
}

// KeyGenerateStruct is a paid mutator transaction binding the contract method 0x88cfe3a4.
//
// Solidity: function keyGenerateStruct((address,bytes32,uint64,bytes32,bytes32,((bytes32,bytes32)[],uint64,address[],uint64)) ) returns()
func (_Wallet *WalletTransactor) KeyGenerateStruct(opts *bind.TransactOpts, arg0 IWalletKeyManagerKeyGenerate) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "keyGenerateStruct", arg0)
}

// KeyGenerateStruct is a paid mutator transaction binding the contract method 0x88cfe3a4.
//
// Solidity: function keyGenerateStruct((address,bytes32,uint64,bytes32,bytes32,((bytes32,bytes32)[],uint64,address[],uint64)) ) returns()
func (_Wallet *WalletSession) KeyGenerateStruct(arg0 IWalletKeyManagerKeyGenerate) (*types.Transaction, error) {
	return _Wallet.Contract.KeyGenerateStruct(&_Wallet.TransactOpts, arg0)
}

// KeyGenerateStruct is a paid mutator transaction binding the contract method 0x88cfe3a4.
//
// Solidity: function keyGenerateStruct((address,bytes32,uint64,bytes32,bytes32,((bytes32,bytes32)[],uint64,address[],uint64)) ) returns()
func (_Wallet *WalletTransactorSession) KeyGenerateStruct(arg0 IWalletKeyManagerKeyGenerate) (*types.Transaction, error) {
	return _Wallet.Contract.KeyGenerateStruct(&_Wallet.TransactOpts, arg0)
}

// ResumeStruct is a paid mutator transaction binding the contract method 0x996da6d5.
//
// Solidity: function resumeStruct((bytes32,(uint64,address,uint256)[]) ) returns()
func (_Wallet *WalletTransactor) ResumeStruct(opts *bind.TransactOpts, arg0 IWalletResumeResume) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "resumeStruct", arg0)
}

// ResumeStruct is a paid mutator transaction binding the contract method 0x996da6d5.
//
// Solidity: function resumeStruct((bytes32,(uint64,address,uint256)[]) ) returns()
func (_Wallet *WalletSession) ResumeStruct(arg0 IWalletResumeResume) (*types.Transaction, error) {
	return _Wallet.Contract.ResumeStruct(&_Wallet.TransactOpts, arg0)
}

// ResumeStruct is a paid mutator transaction binding the contract method 0x996da6d5.
//
// Solidity: function resumeStruct((bytes32,(uint64,address,uint256)[]) ) returns()
func (_Wallet *WalletTransactorSession) ResumeStruct(arg0 IWalletResumeResume) (*types.Transaction, error) {
	return _Wallet.Contract.ResumeStruct(&_Wallet.TransactOpts, arg0)
}

// SetPausingAddressesStruct is a paid mutator transaction binding the contract method 0xb1a9cdac.
//
// Solidity: function setPausingAddressesStruct((bytes32,uint256,(address,uint64)[],address[]) ) returns()
func (_Wallet *WalletTransactor) SetPausingAddressesStruct(opts *bind.TransactOpts, arg0 IWalletResumeSetPausingAddresses) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setPausingAddressesStruct", arg0)
}

// SetPausingAddressesStruct is a paid mutator transaction binding the contract method 0xb1a9cdac.
//
// Solidity: function setPausingAddressesStruct((bytes32,uint256,(address,uint64)[],address[]) ) returns()
func (_Wallet *WalletSession) SetPausingAddressesStruct(arg0 IWalletResumeSetPausingAddresses) (*types.Transaction, error) {
	return _Wallet.Contract.SetPausingAddressesStruct(&_Wallet.TransactOpts, arg0)
}

// SetPausingAddressesStruct is a paid mutator transaction binding the contract method 0xb1a9cdac.
//
// Solidity: function setPausingAddressesStruct((bytes32,uint256,(address,uint64)[],address[]) ) returns()
func (_Wallet *WalletTransactorSession) SetPausingAddressesStruct(arg0 IWalletResumeSetPausingAddresses) (*types.Transaction, error) {
	return _Wallet.Contract.SetPausingAddressesStruct(&_Wallet.TransactOpts, arg0)
}
