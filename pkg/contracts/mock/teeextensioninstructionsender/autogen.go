// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teeextensioninstructionsender

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

// TeeExtensionInstructionsSenderMockTransaction is an auto generated low-level Go binding around an user-defined struct.
type TeeExtensionInstructionsSenderMockTransaction struct {
	Nonce                *big.Int
	ChainId              *big.Int
	To                   common.Address
	Data                 []byte
	Value                *big.Int
	GasPrice             *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	Gas                  *big.Int
}

// TeeExtensionInstructionsSenderMockMetaData contains all meta data concerning the TeeExtensionInstructionsSenderMock contract.
var TeeExtensionInstructionsSenderMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractITeeExtensionRegistry\",\"name\":\"_teeExtensionRegistry\",\"type\":\"address\"},{\"internalType\":\"contractITeeWalletProjectManager\",\"name\":\"_teeWalletProjectManager\",\"type\":\"address\"},{\"internalType\":\"contractITeeWalletManager\",\"name\":\"_teeWalletManager\",\"type\":\"address\"},{\"internalType\":\"contractITeeWalletKeyManager\",\"name\":\"_teeWalletKeyManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"OnlyAuthorizationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KEY_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_COMMAND\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OP_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"internalType\":\"structTeeExtensionInstructionsSenderMock.Transaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"signTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeExtensionRegistry\",\"outputs\":[{\"internalType\":\"contractITeeExtensionRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletKeyManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletKeyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletProjectManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletProjectManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TeeExtensionInstructionsSenderMockABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeExtensionInstructionsSenderMockMetaData.ABI instead.
var TeeExtensionInstructionsSenderMockABI = TeeExtensionInstructionsSenderMockMetaData.ABI

// TeeExtensionInstructionsSenderMock is an auto generated Go binding around an Ethereum contract.
type TeeExtensionInstructionsSenderMock struct {
	TeeExtensionInstructionsSenderMockCaller     // Read-only binding to the contract
	TeeExtensionInstructionsSenderMockTransactor // Write-only binding to the contract
	TeeExtensionInstructionsSenderMockFilterer   // Log filterer for contract events
}

// TeeExtensionInstructionsSenderMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeExtensionInstructionsSenderMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeExtensionInstructionsSenderMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeExtensionInstructionsSenderMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeExtensionInstructionsSenderMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeExtensionInstructionsSenderMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeExtensionInstructionsSenderMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeExtensionInstructionsSenderMockSession struct {
	Contract     *TeeExtensionInstructionsSenderMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// TeeExtensionInstructionsSenderMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeExtensionInstructionsSenderMockCallerSession struct {
	Contract *TeeExtensionInstructionsSenderMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// TeeExtensionInstructionsSenderMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeExtensionInstructionsSenderMockTransactorSession struct {
	Contract     *TeeExtensionInstructionsSenderMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// TeeExtensionInstructionsSenderMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeExtensionInstructionsSenderMockRaw struct {
	Contract *TeeExtensionInstructionsSenderMock // Generic contract binding to access the raw methods on
}

// TeeExtensionInstructionsSenderMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeExtensionInstructionsSenderMockCallerRaw struct {
	Contract *TeeExtensionInstructionsSenderMockCaller // Generic read-only contract binding to access the raw methods on
}

// TeeExtensionInstructionsSenderMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeExtensionInstructionsSenderMockTransactorRaw struct {
	Contract *TeeExtensionInstructionsSenderMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeExtensionInstructionsSenderMock creates a new instance of TeeExtensionInstructionsSenderMock, bound to a specific deployed contract.
func NewTeeExtensionInstructionsSenderMock(address common.Address, backend bind.ContractBackend) (*TeeExtensionInstructionsSenderMock, error) {
	contract, err := bindTeeExtensionInstructionsSenderMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionInstructionsSenderMock{TeeExtensionInstructionsSenderMockCaller: TeeExtensionInstructionsSenderMockCaller{contract: contract}, TeeExtensionInstructionsSenderMockTransactor: TeeExtensionInstructionsSenderMockTransactor{contract: contract}, TeeExtensionInstructionsSenderMockFilterer: TeeExtensionInstructionsSenderMockFilterer{contract: contract}}, nil
}

// NewTeeExtensionInstructionsSenderMockCaller creates a new read-only instance of TeeExtensionInstructionsSenderMock, bound to a specific deployed contract.
func NewTeeExtensionInstructionsSenderMockCaller(address common.Address, caller bind.ContractCaller) (*TeeExtensionInstructionsSenderMockCaller, error) {
	contract, err := bindTeeExtensionInstructionsSenderMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionInstructionsSenderMockCaller{contract: contract}, nil
}

// NewTeeExtensionInstructionsSenderMockTransactor creates a new write-only instance of TeeExtensionInstructionsSenderMock, bound to a specific deployed contract.
func NewTeeExtensionInstructionsSenderMockTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeExtensionInstructionsSenderMockTransactor, error) {
	contract, err := bindTeeExtensionInstructionsSenderMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionInstructionsSenderMockTransactor{contract: contract}, nil
}

// NewTeeExtensionInstructionsSenderMockFilterer creates a new log filterer instance of TeeExtensionInstructionsSenderMock, bound to a specific deployed contract.
func NewTeeExtensionInstructionsSenderMockFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeExtensionInstructionsSenderMockFilterer, error) {
	contract, err := bindTeeExtensionInstructionsSenderMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeExtensionInstructionsSenderMockFilterer{contract: contract}, nil
}

// bindTeeExtensionInstructionsSenderMock binds a generic wrapper to an already deployed contract.
func bindTeeExtensionInstructionsSenderMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeExtensionInstructionsSenderMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeExtensionInstructionsSenderMock.Contract.TeeExtensionInstructionsSenderMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.TeeExtensionInstructionsSenderMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.TeeExtensionInstructionsSenderMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeExtensionInstructionsSenderMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.contract.Transact(opts, method, params...)
}

// KEYTYPE is a free data retrieval call binding the contract method 0xa2adba83.
//
// Solidity: function KEY_TYPE() view returns(bytes32)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCaller) KEYTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeExtensionInstructionsSenderMock.contract.Call(opts, &out, "KEY_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KEYTYPE is a free data retrieval call binding the contract method 0xa2adba83.
//
// Solidity: function KEY_TYPE() view returns(bytes32)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockSession) KEYTYPE() ([32]byte, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.KEYTYPE(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// KEYTYPE is a free data retrieval call binding the contract method 0xa2adba83.
//
// Solidity: function KEY_TYPE() view returns(bytes32)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCallerSession) KEYTYPE() ([32]byte, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.KEYTYPE(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// OPCOMMAND is a free data retrieval call binding the contract method 0xc0816aae.
//
// Solidity: function OP_COMMAND() view returns(bytes32)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCaller) OPCOMMAND(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeExtensionInstructionsSenderMock.contract.Call(opts, &out, "OP_COMMAND")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPCOMMAND is a free data retrieval call binding the contract method 0xc0816aae.
//
// Solidity: function OP_COMMAND() view returns(bytes32)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockSession) OPCOMMAND() ([32]byte, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.OPCOMMAND(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// OPCOMMAND is a free data retrieval call binding the contract method 0xc0816aae.
//
// Solidity: function OP_COMMAND() view returns(bytes32)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCallerSession) OPCOMMAND() ([32]byte, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.OPCOMMAND(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// OPTYPE is a free data retrieval call binding the contract method 0xb2713759.
//
// Solidity: function OP_TYPE() view returns(bytes32)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCaller) OPTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeExtensionInstructionsSenderMock.contract.Call(opts, &out, "OP_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPTYPE is a free data retrieval call binding the contract method 0xb2713759.
//
// Solidity: function OP_TYPE() view returns(bytes32)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockSession) OPTYPE() ([32]byte, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.OPTYPE(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// OPTYPE is a free data retrieval call binding the contract method 0xb2713759.
//
// Solidity: function OP_TYPE() view returns(bytes32)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCallerSession) OPTYPE() ([32]byte, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.OPTYPE(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCaller) TeeExtensionRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionInstructionsSenderMock.contract.Call(opts, &out, "teeExtensionRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.TeeExtensionRegistry(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCallerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.TeeExtensionRegistry(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCaller) TeeWalletKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionInstructionsSenderMock.contract.Call(opts, &out, "teeWalletKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockSession) TeeWalletKeyManager() (common.Address, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.TeeWalletKeyManager(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCallerSession) TeeWalletKeyManager() (common.Address, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.TeeWalletKeyManager(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCaller) TeeWalletManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionInstructionsSenderMock.contract.Call(opts, &out, "teeWalletManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockSession) TeeWalletManager() (common.Address, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.TeeWalletManager(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCallerSession) TeeWalletManager() (common.Address, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.TeeWalletManager(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCaller) TeeWalletProjectManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeExtensionInstructionsSenderMock.contract.Call(opts, &out, "teeWalletProjectManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.TeeWalletProjectManager(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockCallerSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.TeeWalletProjectManager(&_TeeExtensionInstructionsSenderMock.CallOpts)
}

// SignTransaction is a paid mutator transaction binding the contract method 0x862231da.
//
// Solidity: function signTransaction(bytes32 _walletId, (uint256,uint256,address,bytes,uint256,uint256,uint256,uint256,uint256) _transaction) payable returns()
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockTransactor) SignTransaction(opts *bind.TransactOpts, _walletId [32]byte, _transaction TeeExtensionInstructionsSenderMockTransaction) (*types.Transaction, error) {
	return _TeeExtensionInstructionsSenderMock.contract.Transact(opts, "signTransaction", _walletId, _transaction)
}

// SignTransaction is a paid mutator transaction binding the contract method 0x862231da.
//
// Solidity: function signTransaction(bytes32 _walletId, (uint256,uint256,address,bytes,uint256,uint256,uint256,uint256,uint256) _transaction) payable returns()
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockSession) SignTransaction(_walletId [32]byte, _transaction TeeExtensionInstructionsSenderMockTransaction) (*types.Transaction, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.SignTransaction(&_TeeExtensionInstructionsSenderMock.TransactOpts, _walletId, _transaction)
}

// SignTransaction is a paid mutator transaction binding the contract method 0x862231da.
//
// Solidity: function signTransaction(bytes32 _walletId, (uint256,uint256,address,bytes,uint256,uint256,uint256,uint256,uint256) _transaction) payable returns()
func (_TeeExtensionInstructionsSenderMock *TeeExtensionInstructionsSenderMockTransactorSession) SignTransaction(_walletId [32]byte, _transaction TeeExtensionInstructionsSenderMockTransaction) (*types.Transaction, error) {
	return _TeeExtensionInstructionsSenderMock.Contract.SignTransaction(&_TeeExtensionInstructionsSenderMock.TransactOpts, _walletId, _transaction)
}
