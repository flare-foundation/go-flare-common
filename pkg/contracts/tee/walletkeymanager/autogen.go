// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletkeymanager

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

// WalletKeyManagerMetaData contains all meta data concerning the WalletKeyManager contract.
var WalletKeyManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSettings\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyNotGeneratedOnTeeMachine\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyNotRestoredOnTeeMachine\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeIdAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"name\":\"WalletKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"WalletKeyConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"name\":\"WalletKeyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"keyIds\",\"type\":\"uint64[]\"}],\"name\":\"WalletKeysNotAvailable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"multisigThreshold\",\"type\":\"uint64\"}],\"name\":\"WalletMultisigThresholdSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"addKey\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"cleanUpTeeIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restored\",\"type\":\"bool\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"adminsThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"internalType\":\"structIWalletKeyManager.KeyConfigConstants\",\"name\":\"configConstants\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"settingsVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"settings\",\"type\":\"bytes\"}],\"internalType\":\"structIWalletKeyManager.KeyExistence\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"_teeSignature\",\"type\":\"tuple\"}],\"name\":\"confirmKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"deleteKey\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"getWalletKeyPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"getWalletKeyTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletKeysInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_multisigThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64[]\",\"name\":\"_keyIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"_counter\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"receivingTeesAndKeys\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"internalType\":\"structTeeIdKeyIdPair[]\",\"name\":\"_teeIdKeyIdPairs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_multisigThreshold\",\"type\":\"uint64\"}],\"name\":\"setMultisigThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WalletKeyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletKeyManagerMetaData.ABI instead.
var WalletKeyManagerABI = WalletKeyManagerMetaData.ABI

// WalletKeyManager is an auto generated Go binding around an Ethereum contract.
type WalletKeyManager struct {
	WalletKeyManagerCaller     // Read-only binding to the contract
	WalletKeyManagerTransactor // Write-only binding to the contract
	WalletKeyManagerFilterer   // Log filterer for contract events
}

// WalletKeyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletKeyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletKeyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletKeyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletKeyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletKeyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletKeyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletKeyManagerSession struct {
	Contract     *WalletKeyManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletKeyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletKeyManagerCallerSession struct {
	Contract *WalletKeyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// WalletKeyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletKeyManagerTransactorSession struct {
	Contract     *WalletKeyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// WalletKeyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletKeyManagerRaw struct {
	Contract *WalletKeyManager // Generic contract binding to access the raw methods on
}

// WalletKeyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletKeyManagerCallerRaw struct {
	Contract *WalletKeyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// WalletKeyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletKeyManagerTransactorRaw struct {
	Contract *WalletKeyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletKeyManager creates a new instance of WalletKeyManager, bound to a specific deployed contract.
func NewWalletKeyManager(address common.Address, backend bind.ContractBackend) (*WalletKeyManager, error) {
	contract, err := bindWalletKeyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletKeyManager{WalletKeyManagerCaller: WalletKeyManagerCaller{contract: contract}, WalletKeyManagerTransactor: WalletKeyManagerTransactor{contract: contract}, WalletKeyManagerFilterer: WalletKeyManagerFilterer{contract: contract}}, nil
}

// NewWalletKeyManagerCaller creates a new read-only instance of WalletKeyManager, bound to a specific deployed contract.
func NewWalletKeyManagerCaller(address common.Address, caller bind.ContractCaller) (*WalletKeyManagerCaller, error) {
	contract, err := bindWalletKeyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletKeyManagerCaller{contract: contract}, nil
}

// NewWalletKeyManagerTransactor creates a new write-only instance of WalletKeyManager, bound to a specific deployed contract.
func NewWalletKeyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletKeyManagerTransactor, error) {
	contract, err := bindWalletKeyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletKeyManagerTransactor{contract: contract}, nil
}

// NewWalletKeyManagerFilterer creates a new log filterer instance of WalletKeyManager, bound to a specific deployed contract.
func NewWalletKeyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletKeyManagerFilterer, error) {
	contract, err := bindWalletKeyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletKeyManagerFilterer{contract: contract}, nil
}

// bindWalletKeyManager binds a generic wrapper to an already deployed contract.
func bindWalletKeyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletKeyManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletKeyManager *WalletKeyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletKeyManager.Contract.WalletKeyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletKeyManager *WalletKeyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.WalletKeyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletKeyManager *WalletKeyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.WalletKeyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletKeyManager *WalletKeyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletKeyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletKeyManager *WalletKeyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletKeyManager *WalletKeyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.contract.Transact(opts, method, params...)
}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0xbb825ac3.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (_WalletKeyManager *WalletKeyManagerCaller) GetWalletKeyPublicKey(opts *bind.CallOpts, _walletId [32]byte, _keyId uint64) ([]byte, error) {
	var out []interface{}
	err := _WalletKeyManager.contract.Call(opts, &out, "getWalletKeyPublicKey", _walletId, _keyId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0xbb825ac3.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (_WalletKeyManager *WalletKeyManagerSession) GetWalletKeyPublicKey(_walletId [32]byte, _keyId uint64) ([]byte, error) {
	return _WalletKeyManager.Contract.GetWalletKeyPublicKey(&_WalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0xbb825ac3.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (_WalletKeyManager *WalletKeyManagerCallerSession) GetWalletKeyPublicKey(_walletId [32]byte, _keyId uint64) ([]byte, error) {
	return _WalletKeyManager.Contract.GetWalletKeyPublicKey(&_WalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x08480213.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (_WalletKeyManager *WalletKeyManagerCaller) GetWalletKeyTeeIds(opts *bind.CallOpts, _walletId [32]byte, _keyId uint64) ([]common.Address, error) {
	var out []interface{}
	err := _WalletKeyManager.contract.Call(opts, &out, "getWalletKeyTeeIds", _walletId, _keyId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x08480213.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (_WalletKeyManager *WalletKeyManagerSession) GetWalletKeyTeeIds(_walletId [32]byte, _keyId uint64) ([]common.Address, error) {
	return _WalletKeyManager.Contract.GetWalletKeyTeeIds(&_WalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x08480213.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (_WalletKeyManager *WalletKeyManagerCallerSession) GetWalletKeyTeeIds(_walletId [32]byte, _keyId uint64) ([]common.Address, error) {
	return _WalletKeyManager.Contract.GetWalletKeyTeeIds(&_WalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint64[] _keyIds, uint64 _counter)
func (_WalletKeyManager *WalletKeyManagerCaller) GetWalletKeysInfo(opts *bind.CallOpts, _walletId [32]byte) (struct {
	MultisigThreshold uint64
	KeyIds            []uint64
	Counter           uint64
}, error) {
	var out []interface{}
	err := _WalletKeyManager.contract.Call(opts, &out, "getWalletKeysInfo", _walletId)

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
func (_WalletKeyManager *WalletKeyManagerSession) GetWalletKeysInfo(_walletId [32]byte) (struct {
	MultisigThreshold uint64
	KeyIds            []uint64
	Counter           uint64
}, error) {
	return _WalletKeyManager.Contract.GetWalletKeysInfo(&_WalletKeyManager.CallOpts, _walletId)
}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint64[] _keyIds, uint64 _counter)
func (_WalletKeyManager *WalletKeyManagerCallerSession) GetWalletKeysInfo(_walletId [32]byte) (struct {
	MultisigThreshold uint64
	KeyIds            []uint64
	Counter           uint64
}, error) {
	return _WalletKeyManager.Contract.GetWalletKeysInfo(&_WalletKeyManager.CallOpts, _walletId)
}

// AddKey is a paid mutator transaction binding the contract method 0xba9e7a82.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId, address _claimBackAddress) payable returns(uint64 _keyId)
func (_WalletKeyManager *WalletKeyManagerTransactor) AddKey(opts *bind.TransactOpts, _teeId common.Address, _walletId [32]byte, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletKeyManager.contract.Transact(opts, "addKey", _teeId, _walletId, _claimBackAddress)
}

// AddKey is a paid mutator transaction binding the contract method 0xba9e7a82.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId, address _claimBackAddress) payable returns(uint64 _keyId)
func (_WalletKeyManager *WalletKeyManagerSession) AddKey(_teeId common.Address, _walletId [32]byte, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.AddKey(&_WalletKeyManager.TransactOpts, _teeId, _walletId, _claimBackAddress)
}

// AddKey is a paid mutator transaction binding the contract method 0xba9e7a82.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId, address _claimBackAddress) payable returns(uint64 _keyId)
func (_WalletKeyManager *WalletKeyManagerTransactorSession) AddKey(_teeId common.Address, _walletId [32]byte, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.AddKey(&_WalletKeyManager.TransactOpts, _teeId, _walletId, _claimBackAddress)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0xaec4c979.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (_WalletKeyManager *WalletKeyManagerTransactor) CleanUpTeeIds(opts *bind.TransactOpts, _walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _WalletKeyManager.contract.Transact(opts, "cleanUpTeeIds", _walletId, _keyId)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0xaec4c979.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (_WalletKeyManager *WalletKeyManagerSession) CleanUpTeeIds(_walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.CleanUpTeeIds(&_WalletKeyManager.TransactOpts, _walletId, _keyId)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0xaec4c979.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (_WalletKeyManager *WalletKeyManagerTransactorSession) CleanUpTeeIds(_walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.CleanUpTeeIds(&_WalletKeyManager.TransactOpts, _walletId, _keyId)
}

// ConfirmKey is a paid mutator transaction binding the contract method 0xa02d6fe1.
//
// Solidity: function confirmKey((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) _proof, (uint8,bytes32,bytes32) _teeSignature) returns()
func (_WalletKeyManager *WalletKeyManagerTransactor) ConfirmKey(opts *bind.TransactOpts, _proof IWalletKeyManagerKeyExistence, _teeSignature Signature) (*types.Transaction, error) {
	return _WalletKeyManager.contract.Transact(opts, "confirmKey", _proof, _teeSignature)
}

// ConfirmKey is a paid mutator transaction binding the contract method 0xa02d6fe1.
//
// Solidity: function confirmKey((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) _proof, (uint8,bytes32,bytes32) _teeSignature) returns()
func (_WalletKeyManager *WalletKeyManagerSession) ConfirmKey(_proof IWalletKeyManagerKeyExistence, _teeSignature Signature) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.ConfirmKey(&_WalletKeyManager.TransactOpts, _proof, _teeSignature)
}

// ConfirmKey is a paid mutator transaction binding the contract method 0xa02d6fe1.
//
// Solidity: function confirmKey((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) _proof, (uint8,bytes32,bytes32) _teeSignature) returns()
func (_WalletKeyManager *WalletKeyManagerTransactorSession) ConfirmKey(_proof IWalletKeyManagerKeyExistence, _teeSignature Signature) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.ConfirmKey(&_WalletKeyManager.TransactOpts, _proof, _teeSignature)
}

// DeleteKey is a paid mutator transaction binding the contract method 0x270de236.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId, address _claimBackAddress) payable returns()
func (_WalletKeyManager *WalletKeyManagerTransactor) DeleteKey(opts *bind.TransactOpts, _teeId common.Address, _walletId [32]byte, _keyId uint64, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletKeyManager.contract.Transact(opts, "deleteKey", _teeId, _walletId, _keyId, _claimBackAddress)
}

// DeleteKey is a paid mutator transaction binding the contract method 0x270de236.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId, address _claimBackAddress) payable returns()
func (_WalletKeyManager *WalletKeyManagerSession) DeleteKey(_teeId common.Address, _walletId [32]byte, _keyId uint64, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.DeleteKey(&_WalletKeyManager.TransactOpts, _teeId, _walletId, _keyId, _claimBackAddress)
}

// DeleteKey is a paid mutator transaction binding the contract method 0x270de236.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId, address _claimBackAddress) payable returns()
func (_WalletKeyManager *WalletKeyManagerTransactorSession) DeleteKey(_teeId common.Address, _walletId [32]byte, _keyId uint64, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.DeleteKey(&_WalletKeyManager.TransactOpts, _teeId, _walletId, _keyId, _claimBackAddress)
}

// ReceivingTeesAndKeys is a paid mutator transaction binding the contract method 0xfaddcc06.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,uint64)[] _teeIdKeyIdPairs)
func (_WalletKeyManager *WalletKeyManagerTransactor) ReceivingTeesAndKeys(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _WalletKeyManager.contract.Transact(opts, "receivingTeesAndKeys", _walletId)
}

// ReceivingTeesAndKeys is a paid mutator transaction binding the contract method 0xfaddcc06.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,uint64)[] _teeIdKeyIdPairs)
func (_WalletKeyManager *WalletKeyManagerSession) ReceivingTeesAndKeys(_walletId [32]byte) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.ReceivingTeesAndKeys(&_WalletKeyManager.TransactOpts, _walletId)
}

// ReceivingTeesAndKeys is a paid mutator transaction binding the contract method 0xfaddcc06.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,uint64)[] _teeIdKeyIdPairs)
func (_WalletKeyManager *WalletKeyManagerTransactorSession) ReceivingTeesAndKeys(_walletId [32]byte) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.ReceivingTeesAndKeys(&_WalletKeyManager.TransactOpts, _walletId)
}

// SetMultisigThreshold is a paid mutator transaction binding the contract method 0x62101d0a.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint64 _multisigThreshold) returns()
func (_WalletKeyManager *WalletKeyManagerTransactor) SetMultisigThreshold(opts *bind.TransactOpts, _walletId [32]byte, _multisigThreshold uint64) (*types.Transaction, error) {
	return _WalletKeyManager.contract.Transact(opts, "setMultisigThreshold", _walletId, _multisigThreshold)
}

// SetMultisigThreshold is a paid mutator transaction binding the contract method 0x62101d0a.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint64 _multisigThreshold) returns()
func (_WalletKeyManager *WalletKeyManagerSession) SetMultisigThreshold(_walletId [32]byte, _multisigThreshold uint64) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.SetMultisigThreshold(&_WalletKeyManager.TransactOpts, _walletId, _multisigThreshold)
}

// SetMultisigThreshold is a paid mutator transaction binding the contract method 0x62101d0a.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint64 _multisigThreshold) returns()
func (_WalletKeyManager *WalletKeyManagerTransactorSession) SetMultisigThreshold(_walletId [32]byte, _multisigThreshold uint64) (*types.Transaction, error) {
	return _WalletKeyManager.Contract.SetMultisigThreshold(&_WalletKeyManager.TransactOpts, _walletId, _multisigThreshold)
}

// WalletKeyManagerTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the WalletKeyManager contract.
type WalletKeyManagerTeeInstructionsSentIterator struct {
	Event *WalletKeyManagerTeeInstructionsSent // Event containing the contract specifics and raw log

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
func (it *WalletKeyManagerTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletKeyManagerTeeInstructionsSent)
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
		it.Event = new(WalletKeyManagerTeeInstructionsSent)
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
func (it *WalletKeyManagerTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletKeyManagerTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletKeyManagerTeeInstructionsSent represents a TeeInstructionsSent event raised by the WalletKeyManager contract.
type WalletKeyManagerTeeInstructionsSent struct {
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
func (_WalletKeyManager *WalletKeyManagerFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (*WalletKeyManagerTeeInstructionsSentIterator, error) {

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

	logs, sub, err := _WalletKeyManager.contract.FilterLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletKeyManagerTeeInstructionsSentIterator{contract: _WalletKeyManager.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_WalletKeyManager *WalletKeyManagerFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *WalletKeyManagerTeeInstructionsSent, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

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

	logs, sub, err := _WalletKeyManager.contract.WatchLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletKeyManagerTeeInstructionsSent)
				if err := _WalletKeyManager.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
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
func (_WalletKeyManager *WalletKeyManagerFilterer) ParseTeeInstructionsSent(log types.Log) (*WalletKeyManagerTeeInstructionsSent, error) {
	event := new(WalletKeyManagerTeeInstructionsSent)
	if err := _WalletKeyManager.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletKeyManagerWalletKeyAddedIterator is returned from FilterWalletKeyAdded and is used to iterate over the raw logs and unpacked data for WalletKeyAdded events raised by the WalletKeyManager contract.
type WalletKeyManagerWalletKeyAddedIterator struct {
	Event *WalletKeyManagerWalletKeyAdded // Event containing the contract specifics and raw log

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
func (it *WalletKeyManagerWalletKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletKeyManagerWalletKeyAdded)
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
		it.Event = new(WalletKeyManagerWalletKeyAdded)
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
func (it *WalletKeyManagerWalletKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletKeyManagerWalletKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletKeyManagerWalletKeyAdded represents a WalletKeyAdded event raised by the WalletKeyManager contract.
type WalletKeyManagerWalletKeyAdded struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletKeyAdded is a free log retrieval operation binding the contract event 0x8798bb666f83ef50fc32aae926ecc61732c5f6f056210ade84ff0adeda273955.
//
// Solidity: event WalletKeyAdded(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_WalletKeyManager *WalletKeyManagerFilterer) FilterWalletKeyAdded(opts *bind.FilterOpts, teeId []common.Address, walletId [][32]byte, keyId []uint64) (*WalletKeyManagerWalletKeyAddedIterator, error) {

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

	logs, sub, err := _WalletKeyManager.contract.FilterLogs(opts, "WalletKeyAdded", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletKeyManagerWalletKeyAddedIterator{contract: _WalletKeyManager.contract, event: "WalletKeyAdded", logs: logs, sub: sub}, nil
}

// WatchWalletKeyAdded is a free log subscription operation binding the contract event 0x8798bb666f83ef50fc32aae926ecc61732c5f6f056210ade84ff0adeda273955.
//
// Solidity: event WalletKeyAdded(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_WalletKeyManager *WalletKeyManagerFilterer) WatchWalletKeyAdded(opts *bind.WatchOpts, sink chan<- *WalletKeyManagerWalletKeyAdded, teeId []common.Address, walletId [][32]byte, keyId []uint64) (event.Subscription, error) {

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

	logs, sub, err := _WalletKeyManager.contract.WatchLogs(opts, "WalletKeyAdded", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletKeyManagerWalletKeyAdded)
				if err := _WalletKeyManager.contract.UnpackLog(event, "WalletKeyAdded", log); err != nil {
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
func (_WalletKeyManager *WalletKeyManagerFilterer) ParseWalletKeyAdded(log types.Log) (*WalletKeyManagerWalletKeyAdded, error) {
	event := new(WalletKeyManagerWalletKeyAdded)
	if err := _WalletKeyManager.contract.UnpackLog(event, "WalletKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletKeyManagerWalletKeyConfirmedIterator is returned from FilterWalletKeyConfirmed and is used to iterate over the raw logs and unpacked data for WalletKeyConfirmed events raised by the WalletKeyManager contract.
type WalletKeyManagerWalletKeyConfirmedIterator struct {
	Event *WalletKeyManagerWalletKeyConfirmed // Event containing the contract specifics and raw log

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
func (it *WalletKeyManagerWalletKeyConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletKeyManagerWalletKeyConfirmed)
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
		it.Event = new(WalletKeyManagerWalletKeyConfirmed)
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
func (it *WalletKeyManagerWalletKeyConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletKeyManagerWalletKeyConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletKeyManagerWalletKeyConfirmed represents a WalletKeyConfirmed event raised by the WalletKeyManager contract.
type WalletKeyManagerWalletKeyConfirmed struct {
	TeeId     common.Address
	WalletId  [32]byte
	KeyId     uint64
	PublicKey []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletKeyConfirmed is a free log retrieval operation binding the contract event 0x9eccb667b5a2a1ebae9bbef46a76d93aee881470e069aa53cdbcb7af7a2b3b21.
//
// Solidity: event WalletKeyConfirmed(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId, bytes publicKey)
func (_WalletKeyManager *WalletKeyManagerFilterer) FilterWalletKeyConfirmed(opts *bind.FilterOpts, teeId []common.Address, walletId [][32]byte, keyId []uint64) (*WalletKeyManagerWalletKeyConfirmedIterator, error) {

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

	logs, sub, err := _WalletKeyManager.contract.FilterLogs(opts, "WalletKeyConfirmed", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletKeyManagerWalletKeyConfirmedIterator{contract: _WalletKeyManager.contract, event: "WalletKeyConfirmed", logs: logs, sub: sub}, nil
}

// WatchWalletKeyConfirmed is a free log subscription operation binding the contract event 0x9eccb667b5a2a1ebae9bbef46a76d93aee881470e069aa53cdbcb7af7a2b3b21.
//
// Solidity: event WalletKeyConfirmed(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId, bytes publicKey)
func (_WalletKeyManager *WalletKeyManagerFilterer) WatchWalletKeyConfirmed(opts *bind.WatchOpts, sink chan<- *WalletKeyManagerWalletKeyConfirmed, teeId []common.Address, walletId [][32]byte, keyId []uint64) (event.Subscription, error) {

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

	logs, sub, err := _WalletKeyManager.contract.WatchLogs(opts, "WalletKeyConfirmed", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletKeyManagerWalletKeyConfirmed)
				if err := _WalletKeyManager.contract.UnpackLog(event, "WalletKeyConfirmed", log); err != nil {
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
func (_WalletKeyManager *WalletKeyManagerFilterer) ParseWalletKeyConfirmed(log types.Log) (*WalletKeyManagerWalletKeyConfirmed, error) {
	event := new(WalletKeyManagerWalletKeyConfirmed)
	if err := _WalletKeyManager.contract.UnpackLog(event, "WalletKeyConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletKeyManagerWalletKeyDeletedIterator is returned from FilterWalletKeyDeleted and is used to iterate over the raw logs and unpacked data for WalletKeyDeleted events raised by the WalletKeyManager contract.
type WalletKeyManagerWalletKeyDeletedIterator struct {
	Event *WalletKeyManagerWalletKeyDeleted // Event containing the contract specifics and raw log

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
func (it *WalletKeyManagerWalletKeyDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletKeyManagerWalletKeyDeleted)
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
		it.Event = new(WalletKeyManagerWalletKeyDeleted)
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
func (it *WalletKeyManagerWalletKeyDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletKeyManagerWalletKeyDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletKeyManagerWalletKeyDeleted represents a WalletKeyDeleted event raised by the WalletKeyManager contract.
type WalletKeyManagerWalletKeyDeleted struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletKeyDeleted is a free log retrieval operation binding the contract event 0x067c5c5fab799f3b1a41b54608f7f1fd5ec48e8575327ff3da9e4cc149f8226f.
//
// Solidity: event WalletKeyDeleted(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_WalletKeyManager *WalletKeyManagerFilterer) FilterWalletKeyDeleted(opts *bind.FilterOpts, teeId []common.Address, walletId [][32]byte, keyId []uint64) (*WalletKeyManagerWalletKeyDeletedIterator, error) {

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

	logs, sub, err := _WalletKeyManager.contract.FilterLogs(opts, "WalletKeyDeleted", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletKeyManagerWalletKeyDeletedIterator{contract: _WalletKeyManager.contract, event: "WalletKeyDeleted", logs: logs, sub: sub}, nil
}

// WatchWalletKeyDeleted is a free log subscription operation binding the contract event 0x067c5c5fab799f3b1a41b54608f7f1fd5ec48e8575327ff3da9e4cc149f8226f.
//
// Solidity: event WalletKeyDeleted(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_WalletKeyManager *WalletKeyManagerFilterer) WatchWalletKeyDeleted(opts *bind.WatchOpts, sink chan<- *WalletKeyManagerWalletKeyDeleted, teeId []common.Address, walletId [][32]byte, keyId []uint64) (event.Subscription, error) {

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

	logs, sub, err := _WalletKeyManager.contract.WatchLogs(opts, "WalletKeyDeleted", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletKeyManagerWalletKeyDeleted)
				if err := _WalletKeyManager.contract.UnpackLog(event, "WalletKeyDeleted", log); err != nil {
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
func (_WalletKeyManager *WalletKeyManagerFilterer) ParseWalletKeyDeleted(log types.Log) (*WalletKeyManagerWalletKeyDeleted, error) {
	event := new(WalletKeyManagerWalletKeyDeleted)
	if err := _WalletKeyManager.contract.UnpackLog(event, "WalletKeyDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletKeyManagerWalletKeysNotAvailableIterator is returned from FilterWalletKeysNotAvailable and is used to iterate over the raw logs and unpacked data for WalletKeysNotAvailable events raised by the WalletKeyManager contract.
type WalletKeyManagerWalletKeysNotAvailableIterator struct {
	Event *WalletKeyManagerWalletKeysNotAvailable // Event containing the contract specifics and raw log

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
func (it *WalletKeyManagerWalletKeysNotAvailableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletKeyManagerWalletKeysNotAvailable)
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
		it.Event = new(WalletKeyManagerWalletKeysNotAvailable)
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
func (it *WalletKeyManagerWalletKeysNotAvailableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletKeyManagerWalletKeysNotAvailableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletKeyManagerWalletKeysNotAvailable represents a WalletKeysNotAvailable event raised by the WalletKeyManager contract.
type WalletKeyManagerWalletKeysNotAvailable struct {
	WalletId [32]byte
	KeyIds   []uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletKeysNotAvailable is a free log retrieval operation binding the contract event 0xa82e9445a32f966f6330fc8dc7dbb1a5abf41ce17c2695f248a967d11dc8199c.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint64[] keyIds)
func (_WalletKeyManager *WalletKeyManagerFilterer) FilterWalletKeysNotAvailable(opts *bind.FilterOpts, walletId [][32]byte) (*WalletKeyManagerWalletKeysNotAvailableIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletKeyManager.contract.FilterLogs(opts, "WalletKeysNotAvailable", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletKeyManagerWalletKeysNotAvailableIterator{contract: _WalletKeyManager.contract, event: "WalletKeysNotAvailable", logs: logs, sub: sub}, nil
}

// WatchWalletKeysNotAvailable is a free log subscription operation binding the contract event 0xa82e9445a32f966f6330fc8dc7dbb1a5abf41ce17c2695f248a967d11dc8199c.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint64[] keyIds)
func (_WalletKeyManager *WalletKeyManagerFilterer) WatchWalletKeysNotAvailable(opts *bind.WatchOpts, sink chan<- *WalletKeyManagerWalletKeysNotAvailable, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletKeyManager.contract.WatchLogs(opts, "WalletKeysNotAvailable", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletKeyManagerWalletKeysNotAvailable)
				if err := _WalletKeyManager.contract.UnpackLog(event, "WalletKeysNotAvailable", log); err != nil {
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
func (_WalletKeyManager *WalletKeyManagerFilterer) ParseWalletKeysNotAvailable(log types.Log) (*WalletKeyManagerWalletKeysNotAvailable, error) {
	event := new(WalletKeyManagerWalletKeysNotAvailable)
	if err := _WalletKeyManager.contract.UnpackLog(event, "WalletKeysNotAvailable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletKeyManagerWalletMultisigThresholdSetIterator is returned from FilterWalletMultisigThresholdSet and is used to iterate over the raw logs and unpacked data for WalletMultisigThresholdSet events raised by the WalletKeyManager contract.
type WalletKeyManagerWalletMultisigThresholdSetIterator struct {
	Event *WalletKeyManagerWalletMultisigThresholdSet // Event containing the contract specifics and raw log

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
func (it *WalletKeyManagerWalletMultisigThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletKeyManagerWalletMultisigThresholdSet)
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
		it.Event = new(WalletKeyManagerWalletMultisigThresholdSet)
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
func (it *WalletKeyManagerWalletMultisigThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletKeyManagerWalletMultisigThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletKeyManagerWalletMultisigThresholdSet represents a WalletMultisigThresholdSet event raised by the WalletKeyManager contract.
type WalletKeyManagerWalletMultisigThresholdSet struct {
	WalletId          [32]byte
	MultisigThreshold uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWalletMultisigThresholdSet is a free log retrieval operation binding the contract event 0x3874a09a3d4be5b9ad2530ce7aeac3a302ade2cdc68a57336f7c59bad64cad0c.
//
// Solidity: event WalletMultisigThresholdSet(bytes32 indexed walletId, uint64 multisigThreshold)
func (_WalletKeyManager *WalletKeyManagerFilterer) FilterWalletMultisigThresholdSet(opts *bind.FilterOpts, walletId [][32]byte) (*WalletKeyManagerWalletMultisigThresholdSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletKeyManager.contract.FilterLogs(opts, "WalletMultisigThresholdSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletKeyManagerWalletMultisigThresholdSetIterator{contract: _WalletKeyManager.contract, event: "WalletMultisigThresholdSet", logs: logs, sub: sub}, nil
}

// WatchWalletMultisigThresholdSet is a free log subscription operation binding the contract event 0x3874a09a3d4be5b9ad2530ce7aeac3a302ade2cdc68a57336f7c59bad64cad0c.
//
// Solidity: event WalletMultisigThresholdSet(bytes32 indexed walletId, uint64 multisigThreshold)
func (_WalletKeyManager *WalletKeyManagerFilterer) WatchWalletMultisigThresholdSet(opts *bind.WatchOpts, sink chan<- *WalletKeyManagerWalletMultisigThresholdSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletKeyManager.contract.WatchLogs(opts, "WalletMultisigThresholdSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletKeyManagerWalletMultisigThresholdSet)
				if err := _WalletKeyManager.contract.UnpackLog(event, "WalletMultisigThresholdSet", log); err != nil {
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
func (_WalletKeyManager *WalletKeyManagerFilterer) ParseWalletMultisigThresholdSet(log types.Log) (*WalletKeyManagerWalletMultisigThresholdSet, error) {
	event := new(WalletKeyManagerWalletMultisigThresholdSet)
	if err := _WalletKeyManager.contract.UnpackLog(event, "WalletMultisigThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
