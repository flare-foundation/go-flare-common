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

// ITeeKeyExistenceProof is an auto generated low-level Go binding around an user-defined struct.
type ITeeKeyExistenceProof struct {
	RelayMessage  []byte
	TeeSignatures []Signature
	Data          ITeeKeyExistenceResponse
}

// ITeeKeyExistenceRequestBody is an auto generated low-level Go binding around an user-defined struct.
type ITeeKeyExistenceRequestBody struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    *big.Int
}

// ITeeKeyExistenceResponse is an auto generated low-level Go binding around an user-defined struct.
type ITeeKeyExistenceResponse struct {
	AttestationType [32]byte
	SourceId        [32]byte
	ThresholdBIPS   uint16
	Timestamp       uint64
	RequestBody     ITeeKeyExistenceRequestBody
	ResponseBody    ITeeKeyExistenceResponseBody
}

// ITeeKeyExistenceResponseBody is an auto generated low-level Go binding around an user-defined struct.
type ITeeKeyExistenceResponseBody struct {
	OpType     [32]byte
	PublicKey  []byte
	Restored   bool
	AddressStr string
}

// ITeeRegistryTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryTeeMachine struct {
	TeeId common.Address
	Owner common.Address
	Url   string
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
	KeyId *big.Int
}

// TeeWalletKeyManagerMetaData contains all meta data concerning the TeeWalletKeyManager contract.
var TeeWalletKeyManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"}],\"name\":\"WalletKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"addressStr\",\"type\":\"string\"}],\"name\":\"WalletKeyConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"}],\"name\":\"WalletKeyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"keyIds\",\"type\":\"uint256[]\"}],\"name\":\"WalletKeysNotAvailable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multisigThreshold\",\"type\":\"uint256\"}],\"name\":\"WalletMultisigThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletPaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"addKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_keyId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_keyId\",\"type\":\"uint256\"}],\"name\":\"cleanUpTeeIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"relayMessage\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"}],\"internalType\":\"structITeeKeyExistence.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"restored\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"addressStr\",\"type\":\"string\"}],\"internalType\":\"structITeeKeyExistence.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeKeyExistence.Response\",\"name\":\"data\",\"type\":\"tuple\"}],\"internalType\":\"structITeeKeyExistence.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"confirmKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_keyId\",\"type\":\"uint256\"}],\"name\":\"deleteKey\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getFeeFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_keyId\",\"type\":\"uint256\"}],\"name\":\"getWalletKeyAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_addressStr\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_keyId\",\"type\":\"uint256\"}],\"name\":\"getWalletKeyPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_keyId\",\"type\":\"uint256\"}],\"name\":\"getWalletKeyTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletKeysInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_multisigThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_keyIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_counter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"receivingTeesAndKeys\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structITeeRegistry.TeeMachine[]\",\"name\":\"_teeMachines\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"keyId\",\"type\":\"uint256\"}],\"internalType\":\"structTeeIdKeyIdPair[]\",\"name\":\"_teeIdKeyIdPairs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_keyId\",\"type\":\"uint256\"}],\"name\":\"requestKeyExistenceAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_multisigThreshold\",\"type\":\"uint256\"}],\"name\":\"setMultisigThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// GetFeeFactor is a free data retrieval call binding the contract method 0x5220697c.
//
// Solidity: function getFeeFactor(bytes32 _walletId) view returns(uint256 _feeFactor)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) GetFeeFactor(opts *bind.CallOpts, _walletId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "getFeeFactor", _walletId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeFactor is a free data retrieval call binding the contract method 0x5220697c.
//
// Solidity: function getFeeFactor(bytes32 _walletId) view returns(uint256 _feeFactor)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) GetFeeFactor(_walletId [32]byte) (*big.Int, error) {
	return _TeeWalletKeyManager.Contract.GetFeeFactor(&_TeeWalletKeyManager.CallOpts, _walletId)
}

// GetFeeFactor is a free data retrieval call binding the contract method 0x5220697c.
//
// Solidity: function getFeeFactor(bytes32 _walletId) view returns(uint256 _feeFactor)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) GetFeeFactor(_walletId [32]byte) (*big.Int, error) {
	return _TeeWalletKeyManager.Contract.GetFeeFactor(&_TeeWalletKeyManager.CallOpts, _walletId)
}

// GetWalletKeyAddress is a free data retrieval call binding the contract method 0x777c8861.
//
// Solidity: function getWalletKeyAddress(bytes32 _walletId, uint256 _keyId) view returns(string _addressStr)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) GetWalletKeyAddress(opts *bind.CallOpts, _walletId [32]byte, _keyId *big.Int) (string, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "getWalletKeyAddress", _walletId, _keyId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetWalletKeyAddress is a free data retrieval call binding the contract method 0x777c8861.
//
// Solidity: function getWalletKeyAddress(bytes32 _walletId, uint256 _keyId) view returns(string _addressStr)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) GetWalletKeyAddress(_walletId [32]byte, _keyId *big.Int) (string, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyAddress(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyAddress is a free data retrieval call binding the contract method 0x777c8861.
//
// Solidity: function getWalletKeyAddress(bytes32 _walletId, uint256 _keyId) view returns(string _addressStr)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) GetWalletKeyAddress(_walletId [32]byte, _keyId *big.Int) (string, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyAddress(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0x3170536a.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint256 _keyId) view returns(bytes _publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) GetWalletKeyPublicKey(opts *bind.CallOpts, _walletId [32]byte, _keyId *big.Int) ([]byte, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "getWalletKeyPublicKey", _walletId, _keyId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0x3170536a.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint256 _keyId) view returns(bytes _publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) GetWalletKeyPublicKey(_walletId [32]byte, _keyId *big.Int) ([]byte, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyPublicKey(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0x3170536a.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint256 _keyId) view returns(bytes _publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) GetWalletKeyPublicKey(_walletId [32]byte, _keyId *big.Int) ([]byte, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyPublicKey(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x006a82b9.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint256 _keyId) view returns(address[] _teeIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) GetWalletKeyTeeIds(opts *bind.CallOpts, _walletId [32]byte, _keyId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "getWalletKeyTeeIds", _walletId, _keyId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x006a82b9.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint256 _keyId) view returns(address[] _teeIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) GetWalletKeyTeeIds(_walletId [32]byte, _keyId *big.Int) ([]common.Address, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyTeeIds(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x006a82b9.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint256 _keyId) view returns(address[] _teeIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) GetWalletKeyTeeIds(_walletId [32]byte, _keyId *big.Int) ([]common.Address, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyTeeIds(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint256 _multisigThreshold, uint256[] _keyIds, uint256 _counter)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) GetWalletKeysInfo(opts *bind.CallOpts, _walletId [32]byte) (struct {
	MultisigThreshold *big.Int
	KeyIds            []*big.Int
	Counter           *big.Int
}, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "getWalletKeysInfo", _walletId)

	outstruct := new(struct {
		MultisigThreshold *big.Int
		KeyIds            []*big.Int
		Counter           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MultisigThreshold = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.KeyIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Counter = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint256 _multisigThreshold, uint256[] _keyIds, uint256 _counter)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) GetWalletKeysInfo(_walletId [32]byte) (struct {
	MultisigThreshold *big.Int
	KeyIds            []*big.Int
	Counter           *big.Int
}, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeysInfo(&_TeeWalletKeyManager.CallOpts, _walletId)
}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint256 _multisigThreshold, uint256[] _keyIds, uint256 _counter)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) GetWalletKeysInfo(_walletId [32]byte) (struct {
	MultisigThreshold *big.Int
	KeyIds            []*big.Int
	Counter           *big.Int
}, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeysInfo(&_TeeWalletKeyManager.CallOpts, _walletId)
}

// AddKey is a paid mutator transaction binding the contract method 0x7875fa5c.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId) payable returns(uint256 _keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) AddKey(opts *bind.TransactOpts, _teeId common.Address, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "addKey", _teeId, _walletId)
}

// AddKey is a paid mutator transaction binding the contract method 0x7875fa5c.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId) payable returns(uint256 _keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) AddKey(_teeId common.Address, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.AddKey(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId)
}

// AddKey is a paid mutator transaction binding the contract method 0x7875fa5c.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId) payable returns(uint256 _keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) AddKey(_teeId common.Address, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.AddKey(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0x17b57f74.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint256 _keyId) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) CleanUpTeeIds(opts *bind.TransactOpts, _walletId [32]byte, _keyId *big.Int) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "cleanUpTeeIds", _walletId, _keyId)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0x17b57f74.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint256 _keyId) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) CleanUpTeeIds(_walletId [32]byte, _keyId *big.Int) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.CleanUpTeeIds(&_TeeWalletKeyManager.TransactOpts, _walletId, _keyId)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0x17b57f74.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint256 _keyId) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) CleanUpTeeIds(_walletId [32]byte, _keyId *big.Int) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.CleanUpTeeIds(&_TeeWalletKeyManager.TransactOpts, _walletId, _keyId)
}

// ConfirmKey is a paid mutator transaction binding the contract method 0x30bf7350.
//
// Solidity: function confirmKey((bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,(address,bytes32,uint256),(bytes32,bytes,bool,string))) _proof) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) ConfirmKey(opts *bind.TransactOpts, _proof ITeeKeyExistenceProof) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "confirmKey", _proof)
}

// ConfirmKey is a paid mutator transaction binding the contract method 0x30bf7350.
//
// Solidity: function confirmKey((bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,(address,bytes32,uint256),(bytes32,bytes,bool,string))) _proof) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) ConfirmKey(_proof ITeeKeyExistenceProof) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ConfirmKey(&_TeeWalletKeyManager.TransactOpts, _proof)
}

// ConfirmKey is a paid mutator transaction binding the contract method 0x30bf7350.
//
// Solidity: function confirmKey((bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,(address,bytes32,uint256),(bytes32,bytes,bool,string))) _proof) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) ConfirmKey(_proof ITeeKeyExistenceProof) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ConfirmKey(&_TeeWalletKeyManager.TransactOpts, _proof)
}

// DeleteKey is a paid mutator transaction binding the contract method 0xfefcd702.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint256 _keyId) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) DeleteKey(opts *bind.TransactOpts, _teeId common.Address, _walletId [32]byte, _keyId *big.Int) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "deleteKey", _teeId, _walletId, _keyId)
}

// DeleteKey is a paid mutator transaction binding the contract method 0xfefcd702.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint256 _keyId) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) DeleteKey(_teeId common.Address, _walletId [32]byte, _keyId *big.Int) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.DeleteKey(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId, _keyId)
}

// DeleteKey is a paid mutator transaction binding the contract method 0xfefcd702.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint256 _keyId) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) DeleteKey(_teeId common.Address, _walletId [32]byte, _keyId *big.Int) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.DeleteKey(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId, _keyId)
}

// ReceivingTeesAndKeys is a paid mutator transaction binding the contract method 0xfaddcc06.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,address,string)[] _teeMachines, (address,uint256)[] _teeIdKeyIdPairs)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) ReceivingTeesAndKeys(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "receivingTeesAndKeys", _walletId)
}

// ReceivingTeesAndKeys is a paid mutator transaction binding the contract method 0xfaddcc06.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,address,string)[] _teeMachines, (address,uint256)[] _teeIdKeyIdPairs)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) ReceivingTeesAndKeys(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ReceivingTeesAndKeys(&_TeeWalletKeyManager.TransactOpts, _walletId)
}

// ReceivingTeesAndKeys is a paid mutator transaction binding the contract method 0xfaddcc06.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,address,string)[] _teeMachines, (address,uint256)[] _teeIdKeyIdPairs)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) ReceivingTeesAndKeys(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ReceivingTeesAndKeys(&_TeeWalletKeyManager.TransactOpts, _walletId)
}

// RequestKeyExistenceAttestation is a paid mutator transaction binding the contract method 0xafcbe870.
//
// Solidity: function requestKeyExistenceAttestation(address _teeId, bytes32 _walletId, uint256 _keyId) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) RequestKeyExistenceAttestation(opts *bind.TransactOpts, _teeId common.Address, _walletId [32]byte, _keyId *big.Int) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "requestKeyExistenceAttestation", _teeId, _walletId, _keyId)
}

// RequestKeyExistenceAttestation is a paid mutator transaction binding the contract method 0xafcbe870.
//
// Solidity: function requestKeyExistenceAttestation(address _teeId, bytes32 _walletId, uint256 _keyId) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) RequestKeyExistenceAttestation(_teeId common.Address, _walletId [32]byte, _keyId *big.Int) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.RequestKeyExistenceAttestation(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId, _keyId)
}

// RequestKeyExistenceAttestation is a paid mutator transaction binding the contract method 0xafcbe870.
//
// Solidity: function requestKeyExistenceAttestation(address _teeId, bytes32 _walletId, uint256 _keyId) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) RequestKeyExistenceAttestation(_teeId common.Address, _walletId [32]byte, _keyId *big.Int) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.RequestKeyExistenceAttestation(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId, _keyId)
}

// SetMultisigThreshold is a paid mutator transaction binding the contract method 0xf6696007.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint256 _multisigThreshold) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) SetMultisigThreshold(opts *bind.TransactOpts, _walletId [32]byte, _multisigThreshold *big.Int) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "setMultisigThreshold", _walletId, _multisigThreshold)
}

// SetMultisigThreshold is a paid mutator transaction binding the contract method 0xf6696007.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint256 _multisigThreshold) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) SetMultisigThreshold(_walletId [32]byte, _multisigThreshold *big.Int) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.SetMultisigThreshold(&_TeeWalletKeyManager.TransactOpts, _walletId, _multisigThreshold)
}

// SetMultisigThreshold is a paid mutator transaction binding the contract method 0xf6696007.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint256 _multisigThreshold) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) SetMultisigThreshold(_walletId [32]byte, _multisigThreshold *big.Int) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.SetMultisigThreshold(&_TeeWalletKeyManager.TransactOpts, _walletId, _multisigThreshold)
}

// TeeWalletKeyManagerWalletEnabledIterator is returned from FilterWalletEnabled and is used to iterate over the raw logs and unpacked data for WalletEnabled events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletEnabledIterator struct {
	Event *TeeWalletKeyManagerWalletEnabled // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerWalletEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerWalletEnabled)
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
		it.Event = new(TeeWalletKeyManagerWalletEnabled)
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
func (it *TeeWalletKeyManagerWalletEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerWalletEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerWalletEnabled represents a WalletEnabled event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletEnabled struct {
	WalletId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletEnabled is a free log retrieval operation binding the contract event 0x9dd94a650e5c2e19300b404a2e83ff44d26999ca24996bbff99dcd70f81661bf.
//
// Solidity: event WalletEnabled(bytes32 indexed walletId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletEnabled(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletKeyManagerWalletEnabledIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "WalletEnabled", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerWalletEnabledIterator{contract: _TeeWalletKeyManager.contract, event: "WalletEnabled", logs: logs, sub: sub}, nil
}

// WatchWalletEnabled is a free log subscription operation binding the contract event 0x9dd94a650e5c2e19300b404a2e83ff44d26999ca24996bbff99dcd70f81661bf.
//
// Solidity: event WalletEnabled(bytes32 indexed walletId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletEnabled(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletEnabled, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "WalletEnabled", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerWalletEnabled)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletEnabled", log); err != nil {
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
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseWalletEnabled(log types.Log) (*TeeWalletKeyManagerWalletEnabled, error) {
	event := new(TeeWalletKeyManagerWalletEnabled)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletEnabled", log); err != nil {
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
	KeyId    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletKeyAdded is a free log retrieval operation binding the contract event 0x8a498086639687da3d3c2c0821e82b29aaff6b2c0de5b502dad677a6eebf31fd.
//
// Solidity: event WalletKeyAdded(address indexed teeId, bytes32 indexed walletId, uint256 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletKeyAdded(opts *bind.FilterOpts, teeId []common.Address, walletId [][32]byte, keyId []*big.Int) (*TeeWalletKeyManagerWalletKeyAddedIterator, error) {

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

// WatchWalletKeyAdded is a free log subscription operation binding the contract event 0x8a498086639687da3d3c2c0821e82b29aaff6b2c0de5b502dad677a6eebf31fd.
//
// Solidity: event WalletKeyAdded(address indexed teeId, bytes32 indexed walletId, uint256 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletKeyAdded(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletKeyAdded, teeId []common.Address, walletId [][32]byte, keyId []*big.Int) (event.Subscription, error) {

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

// ParseWalletKeyAdded is a log parse operation binding the contract event 0x8a498086639687da3d3c2c0821e82b29aaff6b2c0de5b502dad677a6eebf31fd.
//
// Solidity: event WalletKeyAdded(address indexed teeId, bytes32 indexed walletId, uint256 indexed keyId)
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
	TeeId      common.Address
	WalletId   [32]byte
	KeyId      *big.Int
	PublicKey  []byte
	AddressStr string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWalletKeyConfirmed is a free log retrieval operation binding the contract event 0xc1841ba51fd962fa4e570f5ffdf73374776001c2b25c89600e314f241a790a0b.
//
// Solidity: event WalletKeyConfirmed(address indexed teeId, bytes32 indexed walletId, uint256 indexed keyId, bytes publicKey, string addressStr)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletKeyConfirmed(opts *bind.FilterOpts, teeId []common.Address, walletId [][32]byte, keyId []*big.Int) (*TeeWalletKeyManagerWalletKeyConfirmedIterator, error) {

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

// WatchWalletKeyConfirmed is a free log subscription operation binding the contract event 0xc1841ba51fd962fa4e570f5ffdf73374776001c2b25c89600e314f241a790a0b.
//
// Solidity: event WalletKeyConfirmed(address indexed teeId, bytes32 indexed walletId, uint256 indexed keyId, bytes publicKey, string addressStr)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletKeyConfirmed(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletKeyConfirmed, teeId []common.Address, walletId [][32]byte, keyId []*big.Int) (event.Subscription, error) {

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

// ParseWalletKeyConfirmed is a log parse operation binding the contract event 0xc1841ba51fd962fa4e570f5ffdf73374776001c2b25c89600e314f241a790a0b.
//
// Solidity: event WalletKeyConfirmed(address indexed teeId, bytes32 indexed walletId, uint256 indexed keyId, bytes publicKey, string addressStr)
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
	KeyId    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletKeyDeleted is a free log retrieval operation binding the contract event 0x3ff28933bfff03aa51030336c57b8fa074b8ea4c8efb0c97c856b51f35080437.
//
// Solidity: event WalletKeyDeleted(address indexed teeId, bytes32 indexed walletId, uint256 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletKeyDeleted(opts *bind.FilterOpts, teeId []common.Address, walletId [][32]byte, keyId []*big.Int) (*TeeWalletKeyManagerWalletKeyDeletedIterator, error) {

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

// WatchWalletKeyDeleted is a free log subscription operation binding the contract event 0x3ff28933bfff03aa51030336c57b8fa074b8ea4c8efb0c97c856b51f35080437.
//
// Solidity: event WalletKeyDeleted(address indexed teeId, bytes32 indexed walletId, uint256 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletKeyDeleted(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletKeyDeleted, teeId []common.Address, walletId [][32]byte, keyId []*big.Int) (event.Subscription, error) {

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

// ParseWalletKeyDeleted is a log parse operation binding the contract event 0x3ff28933bfff03aa51030336c57b8fa074b8ea4c8efb0c97c856b51f35080437.
//
// Solidity: event WalletKeyDeleted(address indexed teeId, bytes32 indexed walletId, uint256 indexed keyId)
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
	KeyIds   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletKeysNotAvailable is a free log retrieval operation binding the contract event 0x081ee386be51ce531a26f6e71426d77d046d2e6ebb881a91609e4cc1f648768e.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint256[] keyIds)
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

// WatchWalletKeysNotAvailable is a free log subscription operation binding the contract event 0x081ee386be51ce531a26f6e71426d77d046d2e6ebb881a91609e4cc1f648768e.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint256[] keyIds)
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

// ParseWalletKeysNotAvailable is a log parse operation binding the contract event 0x081ee386be51ce531a26f6e71426d77d046d2e6ebb881a91609e4cc1f648768e.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint256[] keyIds)
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
	MultisigThreshold *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWalletMultisigThresholdSet is a free log retrieval operation binding the contract event 0x5339483a85c4620ec4dd2dd18dcebef7991e8fdd3ef9a171da72de8bf4bc9784.
//
// Solidity: event WalletMultisigThresholdSet(bytes32 indexed walletId, uint256 multisigThreshold)
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

// WatchWalletMultisigThresholdSet is a free log subscription operation binding the contract event 0x5339483a85c4620ec4dd2dd18dcebef7991e8fdd3ef9a171da72de8bf4bc9784.
//
// Solidity: event WalletMultisigThresholdSet(bytes32 indexed walletId, uint256 multisigThreshold)
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

// ParseWalletMultisigThresholdSet is a log parse operation binding the contract event 0x5339483a85c4620ec4dd2dd18dcebef7991e8fdd3ef9a171da72de8bf4bc9784.
//
// Solidity: event WalletMultisigThresholdSet(bytes32 indexed walletId, uint256 multisigThreshold)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseWalletMultisigThresholdSet(log types.Log) (*TeeWalletKeyManagerWalletMultisigThresholdSet, error) {
	event := new(TeeWalletKeyManagerWalletMultisigThresholdSet)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletMultisigThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerWalletPausedIterator is returned from FilterWalletPaused and is used to iterate over the raw logs and unpacked data for WalletPaused events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletPausedIterator struct {
	Event *TeeWalletKeyManagerWalletPaused // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerWalletPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerWalletPaused)
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
		it.Event = new(TeeWalletKeyManagerWalletPaused)
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
func (it *TeeWalletKeyManagerWalletPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerWalletPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerWalletPaused represents a WalletPaused event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletPaused struct {
	WalletId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletPaused is a free log retrieval operation binding the contract event 0x5bc3bc472409bb4331cff333acd11b5d5b6d97052d783936292f19955b25eef3.
//
// Solidity: event WalletPaused(bytes32 indexed walletId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletPaused(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletKeyManagerWalletPausedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "WalletPaused", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerWalletPausedIterator{contract: _TeeWalletKeyManager.contract, event: "WalletPaused", logs: logs, sub: sub}, nil
}

// WatchWalletPaused is a free log subscription operation binding the contract event 0x5bc3bc472409bb4331cff333acd11b5d5b6d97052d783936292f19955b25eef3.
//
// Solidity: event WalletPaused(bytes32 indexed walletId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletPaused(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletPaused, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "WalletPaused", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerWalletPaused)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletPaused", log); err != nil {
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
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseWalletPaused(log types.Log) (*TeeWalletKeyManagerWalletPaused, error) {
	event := new(TeeWalletKeyManagerWalletPaused)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
