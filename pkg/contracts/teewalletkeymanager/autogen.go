// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teewalletkeymanager

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

// IMachineManagerTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
}

// IWalletKeyManagerKeyConfigConstants is an auto generated low-level Go binding around an user-defined struct.
type IWalletKeyManagerKeyConfigConstants struct {
	AdminsPublicKeys   []PublicKey
	AdminsThreshold    uint64
	Cosigners          []common.Address
	CosignersThreshold uint64
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

// PublicKey is an auto generated low-level Go binding around an user-defined struct.
type PublicKey struct {
	X [32]byte
	Y [32]byte
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// TeeIdKeyIdPair is an auto generated low-level Go binding around an user-defined struct.
type TeeIdKeyIdPair struct {
	TeeId common.Address
	KeyId uint64
}

// TeeWalletKeyManagerMetaData contains all meta data concerning the TeeWalletKeyManager contract.
var TeeWalletKeyManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSettings\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyNotGeneratedOnTeeMachine\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyNotRestoredOnTeeMachine\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeIdAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"name\":\"WalletKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"WalletKeyConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"name\":\"WalletKeyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"keyIds\",\"type\":\"uint64[]\"}],\"name\":\"WalletKeysNotAvailable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"multisigThreshold\",\"type\":\"uint64\"}],\"name\":\"WalletMultisigThresholdSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"addKey\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"cleanUpTeeIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restored\",\"type\":\"bool\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"adminsThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"internalType\":\"structIWalletKeyManager.KeyConfigConstants\",\"name\":\"configConstants\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"settingsVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"settings\",\"type\":\"bytes\"}],\"internalType\":\"structIWalletKeyManager.KeyExistence\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"_teeSignature\",\"type\":\"tuple\"}],\"name\":\"confirmKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"deleteKey\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"getWalletKeyPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"getWalletKeyTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletKeysInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_multisigThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64[]\",\"name\":\"_keyIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"_counter\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"receivingTeesAndKeys\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"internalType\":\"structTeeIdKeyIdPair[]\",\"name\":\"_teeIdKeyIdPairs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_multisigThreshold\",\"type\":\"uint64\"}],\"name\":\"setMultisigThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeWalletKeyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeWalletKeyManagerMetaData.ABI instead.
var TeeWalletKeyManagerABI = TeeWalletKeyManagerMetaData.ABI

// TeeWalletKeyManager is an auto generated Go binding around an Ethereum contract.
type TeeWalletKeyManager struct {
	TeeWalletKeyManagerCaller     // Read-only binding to the contract
	TeeWalletKeyManagerTransactor // Write-only binding to the contract
	TeeWalletKeyManagerFilterer   // Log filterer for contract events
}

// TeeWalletKeyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeWalletKeyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletKeyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeWalletKeyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletKeyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeWalletKeyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletKeyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeWalletKeyManagerSession struct {
	Contract     *TeeWalletKeyManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TeeWalletKeyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeWalletKeyManagerCallerSession struct {
	Contract *TeeWalletKeyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TeeWalletKeyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeWalletKeyManagerTransactorSession struct {
	Contract     *TeeWalletKeyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TeeWalletKeyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeWalletKeyManagerRaw struct {
	Contract *TeeWalletKeyManager // Generic contract binding to access the raw methods on
}

// TeeWalletKeyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeWalletKeyManagerCallerRaw struct {
	Contract *TeeWalletKeyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TeeWalletKeyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeWalletKeyManagerTransactorRaw struct {
	Contract *TeeWalletKeyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeWalletKeyManager creates a new instance of TeeWalletKeyManager, bound to a specific deployed contract.
func NewTeeWalletKeyManager(address common.Address, backend bind.ContractBackend) (*TeeWalletKeyManager, error) {
	contract, err := bindTeeWalletKeyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManager{TeeWalletKeyManagerCaller: TeeWalletKeyManagerCaller{contract: contract}, TeeWalletKeyManagerTransactor: TeeWalletKeyManagerTransactor{contract: contract}, TeeWalletKeyManagerFilterer: TeeWalletKeyManagerFilterer{contract: contract}}, nil
}

// NewTeeWalletKeyManagerCaller creates a new read-only instance of TeeWalletKeyManager, bound to a specific deployed contract.
func NewTeeWalletKeyManagerCaller(address common.Address, caller bind.ContractCaller) (*TeeWalletKeyManagerCaller, error) {
	contract, err := bindTeeWalletKeyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerCaller{contract: contract}, nil
}

// NewTeeWalletKeyManagerTransactor creates a new write-only instance of TeeWalletKeyManager, bound to a specific deployed contract.
func NewTeeWalletKeyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeWalletKeyManagerTransactor, error) {
	contract, err := bindTeeWalletKeyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerTransactor{contract: contract}, nil
}

// NewTeeWalletKeyManagerFilterer creates a new log filterer instance of TeeWalletKeyManager, bound to a specific deployed contract.
func NewTeeWalletKeyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeWalletKeyManagerFilterer, error) {
	contract, err := bindTeeWalletKeyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerFilterer{contract: contract}, nil
}

// bindTeeWalletKeyManager binds a generic wrapper to an already deployed contract.
func bindTeeWalletKeyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeWalletKeyManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletKeyManager *TeeWalletKeyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletKeyManager.Contract.TeeWalletKeyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletKeyManager *TeeWalletKeyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.TeeWalletKeyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletKeyManager *TeeWalletKeyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.TeeWalletKeyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletKeyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.contract.Transact(opts, method, params...)
}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0xbb825ac3.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) GetWalletKeyPublicKey(opts *bind.CallOpts, _walletId [32]byte, _keyId uint64) ([]byte, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "getWalletKeyPublicKey", _walletId, _keyId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0xbb825ac3.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) GetWalletKeyPublicKey(_walletId [32]byte, _keyId uint64) ([]byte, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyPublicKey(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0xbb825ac3.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) GetWalletKeyPublicKey(_walletId [32]byte, _keyId uint64) ([]byte, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyPublicKey(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x08480213.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) GetWalletKeyTeeIds(opts *bind.CallOpts, _walletId [32]byte, _keyId uint64) ([]common.Address, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "getWalletKeyTeeIds", _walletId, _keyId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x08480213.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) GetWalletKeyTeeIds(_walletId [32]byte, _keyId uint64) ([]common.Address, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyTeeIds(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x08480213.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) GetWalletKeyTeeIds(_walletId [32]byte, _keyId uint64) ([]common.Address, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyTeeIds(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint64[] _keyIds, uint64 _counter)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) GetWalletKeysInfo(opts *bind.CallOpts, _walletId [32]byte) (struct {
	MultisigThreshold uint64
	KeyIds            []uint64
	Counter           uint64
}, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "getWalletKeysInfo", _walletId)

	outstruct := new(struct {
		MultisigThreshold uint64
		KeyIds            []uint64
		Counter           uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MultisigThreshold = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.KeyIds = *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)
	outstruct.Counter = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint64[] _keyIds, uint64 _counter)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) GetWalletKeysInfo(_walletId [32]byte) (struct {
	MultisigThreshold uint64
	KeyIds            []uint64
	Counter           uint64
}, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeysInfo(&_TeeWalletKeyManager.CallOpts, _walletId)
}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint64[] _keyIds, uint64 _counter)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) GetWalletKeysInfo(_walletId [32]byte) (struct {
	MultisigThreshold uint64
	KeyIds            []uint64
	Counter           uint64
}, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeysInfo(&_TeeWalletKeyManager.CallOpts, _walletId)
}

// AddKey is a paid mutator transaction binding the contract method 0xba9e7a82.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId, address _claimBackAddress) payable returns(uint64 _keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) AddKey(opts *bind.TransactOpts, _teeId common.Address, _walletId [32]byte, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "addKey", _teeId, _walletId, _claimBackAddress)
}

// AddKey is a paid mutator transaction binding the contract method 0xba9e7a82.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId, address _claimBackAddress) payable returns(uint64 _keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) AddKey(_teeId common.Address, _walletId [32]byte, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.AddKey(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId, _claimBackAddress)
}

// AddKey is a paid mutator transaction binding the contract method 0xba9e7a82.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId, address _claimBackAddress) payable returns(uint64 _keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) AddKey(_teeId common.Address, _walletId [32]byte, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.AddKey(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId, _claimBackAddress)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0xaec4c979.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) CleanUpTeeIds(opts *bind.TransactOpts, _walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "cleanUpTeeIds", _walletId, _keyId)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0xaec4c979.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) CleanUpTeeIds(_walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.CleanUpTeeIds(&_TeeWalletKeyManager.TransactOpts, _walletId, _keyId)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0xaec4c979.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) CleanUpTeeIds(_walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.CleanUpTeeIds(&_TeeWalletKeyManager.TransactOpts, _walletId, _keyId)
}

// ConfirmKey is a paid mutator transaction binding the contract method 0xa02d6fe1.
//
// Solidity: function confirmKey((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) _proof, (uint8,bytes32,bytes32) _teeSignature) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) ConfirmKey(opts *bind.TransactOpts, _proof IWalletKeyManagerKeyExistence, _teeSignature Signature) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "confirmKey", _proof, _teeSignature)
}

// ConfirmKey is a paid mutator transaction binding the contract method 0xa02d6fe1.
//
// Solidity: function confirmKey((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) _proof, (uint8,bytes32,bytes32) _teeSignature) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) ConfirmKey(_proof IWalletKeyManagerKeyExistence, _teeSignature Signature) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ConfirmKey(&_TeeWalletKeyManager.TransactOpts, _proof, _teeSignature)
}

// ConfirmKey is a paid mutator transaction binding the contract method 0xa02d6fe1.
//
// Solidity: function confirmKey((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) _proof, (uint8,bytes32,bytes32) _teeSignature) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) ConfirmKey(_proof IWalletKeyManagerKeyExistence, _teeSignature Signature) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ConfirmKey(&_TeeWalletKeyManager.TransactOpts, _proof, _teeSignature)
}

// DeleteKey is a paid mutator transaction binding the contract method 0x270de236.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId, address _claimBackAddress) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) DeleteKey(opts *bind.TransactOpts, _teeId common.Address, _walletId [32]byte, _keyId uint64, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "deleteKey", _teeId, _walletId, _keyId, _claimBackAddress)
}

// DeleteKey is a paid mutator transaction binding the contract method 0x270de236.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId, address _claimBackAddress) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) DeleteKey(_teeId common.Address, _walletId [32]byte, _keyId uint64, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.DeleteKey(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId, _keyId, _claimBackAddress)
}

// DeleteKey is a paid mutator transaction binding the contract method 0x270de236.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId, address _claimBackAddress) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) DeleteKey(_teeId common.Address, _walletId [32]byte, _keyId uint64, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.DeleteKey(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId, _keyId, _claimBackAddress)
}

// ReceivingTeesAndKeys is a paid mutator transaction binding the contract method 0xfaddcc06.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,uint64)[] _teeIdKeyIdPairs)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) ReceivingTeesAndKeys(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "receivingTeesAndKeys", _walletId)
}

// ReceivingTeesAndKeys is a paid mutator transaction binding the contract method 0xfaddcc06.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,uint64)[] _teeIdKeyIdPairs)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) ReceivingTeesAndKeys(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ReceivingTeesAndKeys(&_TeeWalletKeyManager.TransactOpts, _walletId)
}

// ReceivingTeesAndKeys is a paid mutator transaction binding the contract method 0xfaddcc06.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,uint64)[] _teeIdKeyIdPairs)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) ReceivingTeesAndKeys(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ReceivingTeesAndKeys(&_TeeWalletKeyManager.TransactOpts, _walletId)
}

// SetMultisigThreshold is a paid mutator transaction binding the contract method 0x62101d0a.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint64 _multisigThreshold) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) SetMultisigThreshold(opts *bind.TransactOpts, _walletId [32]byte, _multisigThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "setMultisigThreshold", _walletId, _multisigThreshold)
}

// SetMultisigThreshold is a paid mutator transaction binding the contract method 0x62101d0a.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint64 _multisigThreshold) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) SetMultisigThreshold(_walletId [32]byte, _multisigThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.SetMultisigThreshold(&_TeeWalletKeyManager.TransactOpts, _walletId, _multisigThreshold)
}

// SetMultisigThreshold is a paid mutator transaction binding the contract method 0x62101d0a.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint64 _multisigThreshold) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) SetMultisigThreshold(_walletId [32]byte, _multisigThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.SetMultisigThreshold(&_TeeWalletKeyManager.TransactOpts, _walletId, _multisigThreshold)
}

// TeeWalletKeyManagerTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerTeeInstructionsSentIterator struct {
	Event *TeeWalletKeyManagerTeeInstructionsSent // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerTeeInstructionsSent)
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
		it.Event = new(TeeWalletKeyManagerTeeInstructionsSent)
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
func (it *TeeWalletKeyManagerTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerTeeInstructionsSent represents a TeeInstructionsSent event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerTeeInstructionsSent struct {
	ExtensionId        *big.Int
	InstructionId      [32]byte
	RewardEpochId      uint32
	TeeMachines        []IMachineManagerTeeMachine
	OpType             [32]byte
	OpCommand          [32]byte
	Message            []byte
	Cosigners          []common.Address
	CosignersThreshold uint64
	ClaimBackAddress   common.Address
	Fee                *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTeeInstructionsSent is a free log retrieval operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (*TeeWalletKeyManagerTeeInstructionsSentIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerTeeInstructionsSentIterator{contract: _TeeWalletKeyManager.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerTeeInstructionsSent, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerTeeInstructionsSent)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
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

// ParseTeeInstructionsSent is a log parse operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseTeeInstructionsSent(log types.Log) (*TeeWalletKeyManagerTeeInstructionsSent, error) {
	event := new(TeeWalletKeyManagerTeeInstructionsSent)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerWalletKeyAddedIterator is returned from FilterWalletKeyAdded and is used to iterate over the raw logs and unpacked data for WalletKeyAdded events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeyAddedIterator struct {
	Event *TeeWalletKeyManagerWalletKeyAdded // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerWalletKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerWalletKeyAdded)
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
		it.Event = new(TeeWalletKeyManagerWalletKeyAdded)
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
func (it *TeeWalletKeyManagerWalletKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerWalletKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerWalletKeyAdded represents a WalletKeyAdded event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeyAdded struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletKeyAdded is a free log retrieval operation binding the contract event 0x8798bb666f83ef50fc32aae926ecc61732c5f6f056210ade84ff0adeda273955.
//
// Solidity: event WalletKeyAdded(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletKeyAdded(opts *bind.FilterOpts, teeId []common.Address, walletId [][32]byte, keyId []uint64) (*TeeWalletKeyManagerWalletKeyAddedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "WalletKeyAdded", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerWalletKeyAddedIterator{contract: _TeeWalletKeyManager.contract, event: "WalletKeyAdded", logs: logs, sub: sub}, nil
}

// WatchWalletKeyAdded is a free log subscription operation binding the contract event 0x8798bb666f83ef50fc32aae926ecc61732c5f6f056210ade84ff0adeda273955.
//
// Solidity: event WalletKeyAdded(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletKeyAdded(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletKeyAdded, teeId []common.Address, walletId [][32]byte, keyId []uint64) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "WalletKeyAdded", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerWalletKeyAdded)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeyAdded", log); err != nil {
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

// ParseWalletKeyAdded is a log parse operation binding the contract event 0x8798bb666f83ef50fc32aae926ecc61732c5f6f056210ade84ff0adeda273955.
//
// Solidity: event WalletKeyAdded(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseWalletKeyAdded(log types.Log) (*TeeWalletKeyManagerWalletKeyAdded, error) {
	event := new(TeeWalletKeyManagerWalletKeyAdded)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerWalletKeyConfirmedIterator is returned from FilterWalletKeyConfirmed and is used to iterate over the raw logs and unpacked data for WalletKeyConfirmed events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeyConfirmedIterator struct {
	Event *TeeWalletKeyManagerWalletKeyConfirmed // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerWalletKeyConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerWalletKeyConfirmed)
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
		it.Event = new(TeeWalletKeyManagerWalletKeyConfirmed)
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
func (it *TeeWalletKeyManagerWalletKeyConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerWalletKeyConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerWalletKeyConfirmed represents a WalletKeyConfirmed event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeyConfirmed struct {
	TeeId     common.Address
	WalletId  [32]byte
	KeyId     uint64
	PublicKey []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletKeyConfirmed is a free log retrieval operation binding the contract event 0x9eccb667b5a2a1ebae9bbef46a76d93aee881470e069aa53cdbcb7af7a2b3b21.
//
// Solidity: event WalletKeyConfirmed(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId, bytes publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletKeyConfirmed(opts *bind.FilterOpts, teeId []common.Address, walletId [][32]byte, keyId []uint64) (*TeeWalletKeyManagerWalletKeyConfirmedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "WalletKeyConfirmed", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerWalletKeyConfirmedIterator{contract: _TeeWalletKeyManager.contract, event: "WalletKeyConfirmed", logs: logs, sub: sub}, nil
}

// WatchWalletKeyConfirmed is a free log subscription operation binding the contract event 0x9eccb667b5a2a1ebae9bbef46a76d93aee881470e069aa53cdbcb7af7a2b3b21.
//
// Solidity: event WalletKeyConfirmed(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId, bytes publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletKeyConfirmed(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletKeyConfirmed, teeId []common.Address, walletId [][32]byte, keyId []uint64) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "WalletKeyConfirmed", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerWalletKeyConfirmed)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeyConfirmed", log); err != nil {
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

// ParseWalletKeyConfirmed is a log parse operation binding the contract event 0x9eccb667b5a2a1ebae9bbef46a76d93aee881470e069aa53cdbcb7af7a2b3b21.
//
// Solidity: event WalletKeyConfirmed(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId, bytes publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseWalletKeyConfirmed(log types.Log) (*TeeWalletKeyManagerWalletKeyConfirmed, error) {
	event := new(TeeWalletKeyManagerWalletKeyConfirmed)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeyConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerWalletKeyDeletedIterator is returned from FilterWalletKeyDeleted and is used to iterate over the raw logs and unpacked data for WalletKeyDeleted events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeyDeletedIterator struct {
	Event *TeeWalletKeyManagerWalletKeyDeleted // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerWalletKeyDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerWalletKeyDeleted)
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
		it.Event = new(TeeWalletKeyManagerWalletKeyDeleted)
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
func (it *TeeWalletKeyManagerWalletKeyDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerWalletKeyDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerWalletKeyDeleted represents a WalletKeyDeleted event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeyDeleted struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletKeyDeleted is a free log retrieval operation binding the contract event 0x067c5c5fab799f3b1a41b54608f7f1fd5ec48e8575327ff3da9e4cc149f8226f.
//
// Solidity: event WalletKeyDeleted(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletKeyDeleted(opts *bind.FilterOpts, teeId []common.Address, walletId [][32]byte, keyId []uint64) (*TeeWalletKeyManagerWalletKeyDeletedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "WalletKeyDeleted", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerWalletKeyDeletedIterator{contract: _TeeWalletKeyManager.contract, event: "WalletKeyDeleted", logs: logs, sub: sub}, nil
}

// WatchWalletKeyDeleted is a free log subscription operation binding the contract event 0x067c5c5fab799f3b1a41b54608f7f1fd5ec48e8575327ff3da9e4cc149f8226f.
//
// Solidity: event WalletKeyDeleted(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletKeyDeleted(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletKeyDeleted, teeId []common.Address, walletId [][32]byte, keyId []uint64) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "WalletKeyDeleted", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerWalletKeyDeleted)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeyDeleted", log); err != nil {
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

// ParseWalletKeyDeleted is a log parse operation binding the contract event 0x067c5c5fab799f3b1a41b54608f7f1fd5ec48e8575327ff3da9e4cc149f8226f.
//
// Solidity: event WalletKeyDeleted(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseWalletKeyDeleted(log types.Log) (*TeeWalletKeyManagerWalletKeyDeleted, error) {
	event := new(TeeWalletKeyManagerWalletKeyDeleted)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeyDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerWalletKeysNotAvailableIterator is returned from FilterWalletKeysNotAvailable and is used to iterate over the raw logs and unpacked data for WalletKeysNotAvailable events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeysNotAvailableIterator struct {
	Event *TeeWalletKeyManagerWalletKeysNotAvailable // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerWalletKeysNotAvailableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerWalletKeysNotAvailable)
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
		it.Event = new(TeeWalletKeyManagerWalletKeysNotAvailable)
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
func (it *TeeWalletKeyManagerWalletKeysNotAvailableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerWalletKeysNotAvailableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerWalletKeysNotAvailable represents a WalletKeysNotAvailable event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeysNotAvailable struct {
	WalletId [32]byte
	KeyIds   []uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletKeysNotAvailable is a free log retrieval operation binding the contract event 0xa82e9445a32f966f6330fc8dc7dbb1a5abf41ce17c2695f248a967d11dc8199c.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint64[] keyIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletKeysNotAvailable(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletKeyManagerWalletKeysNotAvailableIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "WalletKeysNotAvailable", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerWalletKeysNotAvailableIterator{contract: _TeeWalletKeyManager.contract, event: "WalletKeysNotAvailable", logs: logs, sub: sub}, nil
}

// WatchWalletKeysNotAvailable is a free log subscription operation binding the contract event 0xa82e9445a32f966f6330fc8dc7dbb1a5abf41ce17c2695f248a967d11dc8199c.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint64[] keyIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletKeysNotAvailable(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletKeysNotAvailable, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "WalletKeysNotAvailable", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerWalletKeysNotAvailable)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeysNotAvailable", log); err != nil {
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

// ParseWalletKeysNotAvailable is a log parse operation binding the contract event 0xa82e9445a32f966f6330fc8dc7dbb1a5abf41ce17c2695f248a967d11dc8199c.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint64[] keyIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseWalletKeysNotAvailable(log types.Log) (*TeeWalletKeyManagerWalletKeysNotAvailable, error) {
	event := new(TeeWalletKeyManagerWalletKeysNotAvailable)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeysNotAvailable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerWalletMultisigThresholdSetIterator is returned from FilterWalletMultisigThresholdSet and is used to iterate over the raw logs and unpacked data for WalletMultisigThresholdSet events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletMultisigThresholdSetIterator struct {
	Event *TeeWalletKeyManagerWalletMultisigThresholdSet // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerWalletMultisigThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerWalletMultisigThresholdSet)
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
		it.Event = new(TeeWalletKeyManagerWalletMultisigThresholdSet)
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
func (it *TeeWalletKeyManagerWalletMultisigThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerWalletMultisigThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerWalletMultisigThresholdSet represents a WalletMultisigThresholdSet event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletMultisigThresholdSet struct {
	WalletId          [32]byte
	MultisigThreshold uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWalletMultisigThresholdSet is a free log retrieval operation binding the contract event 0x3874a09a3d4be5b9ad2530ce7aeac3a302ade2cdc68a57336f7c59bad64cad0c.
//
// Solidity: event WalletMultisigThresholdSet(bytes32 indexed walletId, uint64 multisigThreshold)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletMultisigThresholdSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletKeyManagerWalletMultisigThresholdSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "WalletMultisigThresholdSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerWalletMultisigThresholdSetIterator{contract: _TeeWalletKeyManager.contract, event: "WalletMultisigThresholdSet", logs: logs, sub: sub}, nil
}

// WatchWalletMultisigThresholdSet is a free log subscription operation binding the contract event 0x3874a09a3d4be5b9ad2530ce7aeac3a302ade2cdc68a57336f7c59bad64cad0c.
//
// Solidity: event WalletMultisigThresholdSet(bytes32 indexed walletId, uint64 multisigThreshold)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletMultisigThresholdSet(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletMultisigThresholdSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "WalletMultisigThresholdSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerWalletMultisigThresholdSet)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletMultisigThresholdSet", log); err != nil {
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

// ParseWalletMultisigThresholdSet is a log parse operation binding the contract event 0x3874a09a3d4be5b9ad2530ce7aeac3a302ade2cdc68a57336f7c59bad64cad0c.
//
// Solidity: event WalletMultisigThresholdSet(bytes32 indexed walletId, uint64 multisigThreshold)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseWalletMultisigThresholdSet(log types.Log) (*TeeWalletKeyManagerWalletMultisigThresholdSet, error) {
	event := new(TeeWalletKeyManagerWalletMultisigThresholdSet)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletMultisigThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
