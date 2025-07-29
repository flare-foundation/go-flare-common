// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teeinstructions

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

// ITeePaymentsPaymentInstruction is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsPaymentInstruction struct {
	RecipientAddress string
	Amount           *big.Int
	Fee              *big.Int
	PaymentReference [32]byte
}

// TeeInstructionsMetaData contains all meta data concerning the TeeInstructions contract.
var TeeInstructionsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"BatchSettingsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"minFee\",\"type\":\"uint128\"}],\"name\":\"MinFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"initialNonce\",\"type\":\"uint64\"}],\"name\":\"SenderAddressSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getBatchSettings\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getMinFee\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"_minFee\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOpType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getSenderAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_senderAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePayments.PaymentInstruction\",\"name\":\"_paymentInstruction\",\"type\":\"tuple\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_subNonce\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_firstSubNonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePayments.PaymentInstruction[]\",\"name\":\"_paymentInstructions\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_nullify\",\"type\":\"bool[]\"}],\"name\":\"reissue\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"setBatchSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"_minFee\",\"type\":\"uint128\"}],\"name\":\"setMinFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_transactionLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"name\":\"setPaymentLimits\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_senderAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"_initialNonce\",\"type\":\"uint64\"}],\"name\":\"setSenderAddressAndInitialNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeInstructionsABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeInstructionsMetaData.ABI instead.
var TeeInstructionsABI = TeeInstructionsMetaData.ABI

// TeeInstructions is an auto generated Go binding around an Ethereum contract.
type TeeInstructions struct {
	TeeInstructionsCaller     // Read-only binding to the contract
	TeeInstructionsTransactor // Write-only binding to the contract
	TeeInstructionsFilterer   // Log filterer for contract events
}

// TeeInstructionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeInstructionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeInstructionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeInstructionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeInstructionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeInstructionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeInstructionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeInstructionsSession struct {
	Contract     *TeeInstructions  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeInstructionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeInstructionsCallerSession struct {
	Contract *TeeInstructionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TeeInstructionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeInstructionsTransactorSession struct {
	Contract     *TeeInstructionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TeeInstructionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeInstructionsRaw struct {
	Contract *TeeInstructions // Generic contract binding to access the raw methods on
}

// TeeInstructionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeInstructionsCallerRaw struct {
	Contract *TeeInstructionsCaller // Generic read-only contract binding to access the raw methods on
}

// TeeInstructionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeInstructionsTransactorRaw struct {
	Contract *TeeInstructionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeInstructions creates a new instance of TeeInstructions, bound to a specific deployed contract.
func NewTeeInstructions(address common.Address, backend bind.ContractBackend) (*TeeInstructions, error) {
	contract, err := bindTeeInstructions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeInstructions{TeeInstructionsCaller: TeeInstructionsCaller{contract: contract}, TeeInstructionsTransactor: TeeInstructionsTransactor{contract: contract}, TeeInstructionsFilterer: TeeInstructionsFilterer{contract: contract}}, nil
}

// NewTeeInstructionsCaller creates a new read-only instance of TeeInstructions, bound to a specific deployed contract.
func NewTeeInstructionsCaller(address common.Address, caller bind.ContractCaller) (*TeeInstructionsCaller, error) {
	contract, err := bindTeeInstructions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsCaller{contract: contract}, nil
}

// NewTeeInstructionsTransactor creates a new write-only instance of TeeInstructions, bound to a specific deployed contract.
func NewTeeInstructionsTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeInstructionsTransactor, error) {
	contract, err := bindTeeInstructions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsTransactor{contract: contract}, nil
}

// NewTeeInstructionsFilterer creates a new log filterer instance of TeeInstructions, bound to a specific deployed contract.
func NewTeeInstructionsFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeInstructionsFilterer, error) {
	contract, err := bindTeeInstructions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsFilterer{contract: contract}, nil
}

// bindTeeInstructions binds a generic wrapper to an already deployed contract.
func bindTeeInstructions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeInstructionsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeInstructions *TeeInstructionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeInstructions.Contract.TeeInstructionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeInstructions *TeeInstructionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeInstructions.Contract.TeeInstructionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeInstructions *TeeInstructionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeInstructions.Contract.TeeInstructionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeInstructions *TeeInstructionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeInstructions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeInstructions *TeeInstructionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeInstructions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeInstructions *TeeInstructionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeInstructions.Contract.contract.Transact(opts, method, params...)
}

// GetBatchSettings is a free data retrieval call binding the contract method 0x322dd91b.
//
// Solidity: function getBatchSettings(bytes32 _walletId) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (_TeeInstructions *TeeInstructionsCaller) GetBatchSettings(opts *bind.CallOpts, _walletId [32]byte) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	var out []interface{}
	err := _TeeInstructions.contract.Call(opts, &out, "getBatchSettings", _walletId)

	outstruct := new(struct {
		BatchSize            uint64
		BatchDurationSeconds uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchSize = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.BatchDurationSeconds = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetBatchSettings is a free data retrieval call binding the contract method 0x322dd91b.
//
// Solidity: function getBatchSettings(bytes32 _walletId) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (_TeeInstructions *TeeInstructionsSession) GetBatchSettings(_walletId [32]byte) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	return _TeeInstructions.Contract.GetBatchSettings(&_TeeInstructions.CallOpts, _walletId)
}

// GetBatchSettings is a free data retrieval call binding the contract method 0x322dd91b.
//
// Solidity: function getBatchSettings(bytes32 _walletId) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (_TeeInstructions *TeeInstructionsCallerSession) GetBatchSettings(_walletId [32]byte) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	return _TeeInstructions.Contract.GetBatchSettings(&_TeeInstructions.CallOpts, _walletId)
}

// GetMinFee is a free data retrieval call binding the contract method 0x8f937b01.
//
// Solidity: function getMinFee(bytes32 _walletId) view returns(uint128 _minFee)
func (_TeeInstructions *TeeInstructionsCaller) GetMinFee(opts *bind.CallOpts, _walletId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeeInstructions.contract.Call(opts, &out, "getMinFee", _walletId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinFee is a free data retrieval call binding the contract method 0x8f937b01.
//
// Solidity: function getMinFee(bytes32 _walletId) view returns(uint128 _minFee)
func (_TeeInstructions *TeeInstructionsSession) GetMinFee(_walletId [32]byte) (*big.Int, error) {
	return _TeeInstructions.Contract.GetMinFee(&_TeeInstructions.CallOpts, _walletId)
}

// GetMinFee is a free data retrieval call binding the contract method 0x8f937b01.
//
// Solidity: function getMinFee(bytes32 _walletId) view returns(uint128 _minFee)
func (_TeeInstructions *TeeInstructionsCallerSession) GetMinFee(_walletId [32]byte) (*big.Int, error) {
	return _TeeInstructions.Contract.GetMinFee(&_TeeInstructions.CallOpts, _walletId)
}

// GetOpType is a free data retrieval call binding the contract method 0xa3e2d39a.
//
// Solidity: function getOpType() view returns(bytes32)
func (_TeeInstructions *TeeInstructionsCaller) GetOpType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeInstructions.contract.Call(opts, &out, "getOpType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOpType is a free data retrieval call binding the contract method 0xa3e2d39a.
//
// Solidity: function getOpType() view returns(bytes32)
func (_TeeInstructions *TeeInstructionsSession) GetOpType() ([32]byte, error) {
	return _TeeInstructions.Contract.GetOpType(&_TeeInstructions.CallOpts)
}

// GetOpType is a free data retrieval call binding the contract method 0xa3e2d39a.
//
// Solidity: function getOpType() view returns(bytes32)
func (_TeeInstructions *TeeInstructionsCallerSession) GetOpType() ([32]byte, error) {
	return _TeeInstructions.Contract.GetOpType(&_TeeInstructions.CallOpts)
}

// GetSenderAddress is a free data retrieval call binding the contract method 0x63453bba.
//
// Solidity: function getSenderAddress(bytes32 _walletId) view returns(string _senderAddress)
func (_TeeInstructions *TeeInstructionsCaller) GetSenderAddress(opts *bind.CallOpts, _walletId [32]byte) (string, error) {
	var out []interface{}
	err := _TeeInstructions.contract.Call(opts, &out, "getSenderAddress", _walletId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSenderAddress is a free data retrieval call binding the contract method 0x63453bba.
//
// Solidity: function getSenderAddress(bytes32 _walletId) view returns(string _senderAddress)
func (_TeeInstructions *TeeInstructionsSession) GetSenderAddress(_walletId [32]byte) (string, error) {
	return _TeeInstructions.Contract.GetSenderAddress(&_TeeInstructions.CallOpts, _walletId)
}

// GetSenderAddress is a free data retrieval call binding the contract method 0x63453bba.
//
// Solidity: function getSenderAddress(bytes32 _walletId) view returns(string _senderAddress)
func (_TeeInstructions *TeeInstructionsCallerSession) GetSenderAddress(_walletId [32]byte) (string, error) {
	return _TeeInstructions.Contract.GetSenderAddress(&_TeeInstructions.CallOpts, _walletId)
}

// Pay is a paid mutator transaction binding the contract method 0xe253d4c7.
//
// Solidity: function pay(bytes32 _projectId, bytes32 _walletId, (string,uint256,uint256,bytes32) _paymentInstruction) payable returns(uint64 _nonce, uint64 _subNonce)
func (_TeeInstructions *TeeInstructionsTransactor) Pay(opts *bind.TransactOpts, _projectId [32]byte, _walletId [32]byte, _paymentInstruction ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _TeeInstructions.contract.Transact(opts, "pay", _projectId, _walletId, _paymentInstruction)
}

// Pay is a paid mutator transaction binding the contract method 0xe253d4c7.
//
// Solidity: function pay(bytes32 _projectId, bytes32 _walletId, (string,uint256,uint256,bytes32) _paymentInstruction) payable returns(uint64 _nonce, uint64 _subNonce)
func (_TeeInstructions *TeeInstructionsSession) Pay(_projectId [32]byte, _walletId [32]byte, _paymentInstruction ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _TeeInstructions.Contract.Pay(&_TeeInstructions.TransactOpts, _projectId, _walletId, _paymentInstruction)
}

// Pay is a paid mutator transaction binding the contract method 0xe253d4c7.
//
// Solidity: function pay(bytes32 _projectId, bytes32 _walletId, (string,uint256,uint256,bytes32) _paymentInstruction) payable returns(uint64 _nonce, uint64 _subNonce)
func (_TeeInstructions *TeeInstructionsTransactorSession) Pay(_projectId [32]byte, _walletId [32]byte, _paymentInstruction ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _TeeInstructions.Contract.Pay(&_TeeInstructions.TransactOpts, _projectId, _walletId, _paymentInstruction)
}

// Reissue is a paid mutator transaction binding the contract method 0xcac836b8.
//
// Solidity: function reissue(bytes32 _walletId, uint64 _nonce, uint64 _firstSubNonce, (string,uint256,uint256,bytes32)[] _paymentInstructions, uint256[] _fees, bool[] _nullify) payable returns()
func (_TeeInstructions *TeeInstructionsTransactor) Reissue(opts *bind.TransactOpts, _walletId [32]byte, _nonce uint64, _firstSubNonce uint64, _paymentInstructions []ITeePaymentsPaymentInstruction, _fees []*big.Int, _nullify []bool) (*types.Transaction, error) {
	return _TeeInstructions.contract.Transact(opts, "reissue", _walletId, _nonce, _firstSubNonce, _paymentInstructions, _fees, _nullify)
}

// Reissue is a paid mutator transaction binding the contract method 0xcac836b8.
//
// Solidity: function reissue(bytes32 _walletId, uint64 _nonce, uint64 _firstSubNonce, (string,uint256,uint256,bytes32)[] _paymentInstructions, uint256[] _fees, bool[] _nullify) payable returns()
func (_TeeInstructions *TeeInstructionsSession) Reissue(_walletId [32]byte, _nonce uint64, _firstSubNonce uint64, _paymentInstructions []ITeePaymentsPaymentInstruction, _fees []*big.Int, _nullify []bool) (*types.Transaction, error) {
	return _TeeInstructions.Contract.Reissue(&_TeeInstructions.TransactOpts, _walletId, _nonce, _firstSubNonce, _paymentInstructions, _fees, _nullify)
}

// Reissue is a paid mutator transaction binding the contract method 0xcac836b8.
//
// Solidity: function reissue(bytes32 _walletId, uint64 _nonce, uint64 _firstSubNonce, (string,uint256,uint256,bytes32)[] _paymentInstructions, uint256[] _fees, bool[] _nullify) payable returns()
func (_TeeInstructions *TeeInstructionsTransactorSession) Reissue(_walletId [32]byte, _nonce uint64, _firstSubNonce uint64, _paymentInstructions []ITeePaymentsPaymentInstruction, _fees []*big.Int, _nullify []bool) (*types.Transaction, error) {
	return _TeeInstructions.Contract.Reissue(&_TeeInstructions.TransactOpts, _walletId, _nonce, _firstSubNonce, _paymentInstructions, _fees, _nullify)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xf680cbad.
//
// Solidity: function setBatchSettings(bytes32 _walletId, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeeInstructions *TeeInstructionsTransactor) SetBatchSettings(opts *bind.TransactOpts, _walletId [32]byte, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeInstructions.contract.Transact(opts, "setBatchSettings", _walletId, _batchSize, _batchDurationSeconds)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xf680cbad.
//
// Solidity: function setBatchSettings(bytes32 _walletId, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeeInstructions *TeeInstructionsSession) SetBatchSettings(_walletId [32]byte, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SetBatchSettings(&_TeeInstructions.TransactOpts, _walletId, _batchSize, _batchDurationSeconds)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xf680cbad.
//
// Solidity: function setBatchSettings(bytes32 _walletId, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeeInstructions *TeeInstructionsTransactorSession) SetBatchSettings(_walletId [32]byte, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SetBatchSettings(&_TeeInstructions.TransactOpts, _walletId, _batchSize, _batchDurationSeconds)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x88da8aa6.
//
// Solidity: function setMinFee(bytes32 _walletId, uint128 _minFee) returns()
func (_TeeInstructions *TeeInstructionsTransactor) SetMinFee(opts *bind.TransactOpts, _walletId [32]byte, _minFee *big.Int) (*types.Transaction, error) {
	return _TeeInstructions.contract.Transact(opts, "setMinFee", _walletId, _minFee)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x88da8aa6.
//
// Solidity: function setMinFee(bytes32 _walletId, uint128 _minFee) returns()
func (_TeeInstructions *TeeInstructionsSession) SetMinFee(_walletId [32]byte, _minFee *big.Int) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SetMinFee(&_TeeInstructions.TransactOpts, _walletId, _minFee)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x88da8aa6.
//
// Solidity: function setMinFee(bytes32 _walletId, uint128 _minFee) returns()
func (_TeeInstructions *TeeInstructionsTransactorSession) SetMinFee(_walletId [32]byte, _minFee *big.Int) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SetMinFee(&_TeeInstructions.TransactOpts, _walletId, _minFee)
}

// SetPaymentLimits is a paid mutator transaction binding the contract method 0xb01f700c.
//
// Solidity: function setPaymentLimits(bytes32 _walletId, uint256 _transactionLimit, uint256 _dailyLimit) payable returns()
func (_TeeInstructions *TeeInstructionsTransactor) SetPaymentLimits(opts *bind.TransactOpts, _walletId [32]byte, _transactionLimit *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _TeeInstructions.contract.Transact(opts, "setPaymentLimits", _walletId, _transactionLimit, _dailyLimit)
}

// SetPaymentLimits is a paid mutator transaction binding the contract method 0xb01f700c.
//
// Solidity: function setPaymentLimits(bytes32 _walletId, uint256 _transactionLimit, uint256 _dailyLimit) payable returns()
func (_TeeInstructions *TeeInstructionsSession) SetPaymentLimits(_walletId [32]byte, _transactionLimit *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SetPaymentLimits(&_TeeInstructions.TransactOpts, _walletId, _transactionLimit, _dailyLimit)
}

// SetPaymentLimits is a paid mutator transaction binding the contract method 0xb01f700c.
//
// Solidity: function setPaymentLimits(bytes32 _walletId, uint256 _transactionLimit, uint256 _dailyLimit) payable returns()
func (_TeeInstructions *TeeInstructionsTransactorSession) SetPaymentLimits(_walletId [32]byte, _transactionLimit *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SetPaymentLimits(&_TeeInstructions.TransactOpts, _walletId, _transactionLimit, _dailyLimit)
}

// SetSenderAddressAndInitialNonce is a paid mutator transaction binding the contract method 0xb9a853c3.
//
// Solidity: function setSenderAddressAndInitialNonce(bytes32 _walletId, string _senderAddress, uint64 _initialNonce) returns()
func (_TeeInstructions *TeeInstructionsTransactor) SetSenderAddressAndInitialNonce(opts *bind.TransactOpts, _walletId [32]byte, _senderAddress string, _initialNonce uint64) (*types.Transaction, error) {
	return _TeeInstructions.contract.Transact(opts, "setSenderAddressAndInitialNonce", _walletId, _senderAddress, _initialNonce)
}

// SetSenderAddressAndInitialNonce is a paid mutator transaction binding the contract method 0xb9a853c3.
//
// Solidity: function setSenderAddressAndInitialNonce(bytes32 _walletId, string _senderAddress, uint64 _initialNonce) returns()
func (_TeeInstructions *TeeInstructionsSession) SetSenderAddressAndInitialNonce(_walletId [32]byte, _senderAddress string, _initialNonce uint64) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SetSenderAddressAndInitialNonce(&_TeeInstructions.TransactOpts, _walletId, _senderAddress, _initialNonce)
}

// SetSenderAddressAndInitialNonce is a paid mutator transaction binding the contract method 0xb9a853c3.
//
// Solidity: function setSenderAddressAndInitialNonce(bytes32 _walletId, string _senderAddress, uint64 _initialNonce) returns()
func (_TeeInstructions *TeeInstructionsTransactorSession) SetSenderAddressAndInitialNonce(_walletId [32]byte, _senderAddress string, _initialNonce uint64) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SetSenderAddressAndInitialNonce(&_TeeInstructions.TransactOpts, _walletId, _senderAddress, _initialNonce)
}

// TeeInstructionsBatchSettingsSetIterator is returned from FilterBatchSettingsSet and is used to iterate over the raw logs and unpacked data for BatchSettingsSet events raised by the TeeInstructions contract.
type TeeInstructionsBatchSettingsSetIterator struct {
	Event *TeeInstructionsBatchSettingsSet // Event containing the contract specifics and raw log

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
func (it *TeeInstructionsBatchSettingsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeInstructionsBatchSettingsSet)
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
		it.Event = new(TeeInstructionsBatchSettingsSet)
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
func (it *TeeInstructionsBatchSettingsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeInstructionsBatchSettingsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeInstructionsBatchSettingsSet represents a BatchSettingsSet event raised by the TeeInstructions contract.
type TeeInstructionsBatchSettingsSet struct {
	WalletId             [32]byte
	BatchSize            uint64
	BatchDurationSeconds uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBatchSettingsSet is a free log retrieval operation binding the contract event 0xfa2da2b02c35a83b9a69ec27c4d6fbf1c4cf31d416e5783d45a39bdfbc36deff.
//
// Solidity: event BatchSettingsSet(bytes32 indexed walletId, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeeInstructions *TeeInstructionsFilterer) FilterBatchSettingsSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeeInstructionsBatchSettingsSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeInstructions.contract.FilterLogs(opts, "BatchSettingsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsBatchSettingsSetIterator{contract: _TeeInstructions.contract, event: "BatchSettingsSet", logs: logs, sub: sub}, nil
}

// WatchBatchSettingsSet is a free log subscription operation binding the contract event 0xfa2da2b02c35a83b9a69ec27c4d6fbf1c4cf31d416e5783d45a39bdfbc36deff.
//
// Solidity: event BatchSettingsSet(bytes32 indexed walletId, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeeInstructions *TeeInstructionsFilterer) WatchBatchSettingsSet(opts *bind.WatchOpts, sink chan<- *TeeInstructionsBatchSettingsSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeInstructions.contract.WatchLogs(opts, "BatchSettingsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeInstructionsBatchSettingsSet)
				if err := _TeeInstructions.contract.UnpackLog(event, "BatchSettingsSet", log); err != nil {
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

// ParseBatchSettingsSet is a log parse operation binding the contract event 0xfa2da2b02c35a83b9a69ec27c4d6fbf1c4cf31d416e5783d45a39bdfbc36deff.
//
// Solidity: event BatchSettingsSet(bytes32 indexed walletId, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeeInstructions *TeeInstructionsFilterer) ParseBatchSettingsSet(log types.Log) (*TeeInstructionsBatchSettingsSet, error) {
	event := new(TeeInstructionsBatchSettingsSet)
	if err := _TeeInstructions.contract.UnpackLog(event, "BatchSettingsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeInstructionsMinFeeSetIterator is returned from FilterMinFeeSet and is used to iterate over the raw logs and unpacked data for MinFeeSet events raised by the TeeInstructions contract.
type TeeInstructionsMinFeeSetIterator struct {
	Event *TeeInstructionsMinFeeSet // Event containing the contract specifics and raw log

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
func (it *TeeInstructionsMinFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeInstructionsMinFeeSet)
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
		it.Event = new(TeeInstructionsMinFeeSet)
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
func (it *TeeInstructionsMinFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeInstructionsMinFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeInstructionsMinFeeSet represents a MinFeeSet event raised by the TeeInstructions contract.
type TeeInstructionsMinFeeSet struct {
	WalletId [32]byte
	MinFee   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinFeeSet is a free log retrieval operation binding the contract event 0xc37f37cead8822d50df03488bfae0fcd52d892c3d18eb730441eccbedcddd66f.
//
// Solidity: event MinFeeSet(bytes32 indexed walletId, uint128 minFee)
func (_TeeInstructions *TeeInstructionsFilterer) FilterMinFeeSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeeInstructionsMinFeeSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeInstructions.contract.FilterLogs(opts, "MinFeeSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsMinFeeSetIterator{contract: _TeeInstructions.contract, event: "MinFeeSet", logs: logs, sub: sub}, nil
}

// WatchMinFeeSet is a free log subscription operation binding the contract event 0xc37f37cead8822d50df03488bfae0fcd52d892c3d18eb730441eccbedcddd66f.
//
// Solidity: event MinFeeSet(bytes32 indexed walletId, uint128 minFee)
func (_TeeInstructions *TeeInstructionsFilterer) WatchMinFeeSet(opts *bind.WatchOpts, sink chan<- *TeeInstructionsMinFeeSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeInstructions.contract.WatchLogs(opts, "MinFeeSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeInstructionsMinFeeSet)
				if err := _TeeInstructions.contract.UnpackLog(event, "MinFeeSet", log); err != nil {
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

// ParseMinFeeSet is a log parse operation binding the contract event 0xc37f37cead8822d50df03488bfae0fcd52d892c3d18eb730441eccbedcddd66f.
//
// Solidity: event MinFeeSet(bytes32 indexed walletId, uint128 minFee)
func (_TeeInstructions *TeeInstructionsFilterer) ParseMinFeeSet(log types.Log) (*TeeInstructionsMinFeeSet, error) {
	event := new(TeeInstructionsMinFeeSet)
	if err := _TeeInstructions.contract.UnpackLog(event, "MinFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeInstructionsSenderAddressSetIterator is returned from FilterSenderAddressSet and is used to iterate over the raw logs and unpacked data for SenderAddressSet events raised by the TeeInstructions contract.
type TeeInstructionsSenderAddressSetIterator struct {
	Event *TeeInstructionsSenderAddressSet // Event containing the contract specifics and raw log

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
func (it *TeeInstructionsSenderAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeInstructionsSenderAddressSet)
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
		it.Event = new(TeeInstructionsSenderAddressSet)
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
func (it *TeeInstructionsSenderAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeInstructionsSenderAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeInstructionsSenderAddressSet represents a SenderAddressSet event raised by the TeeInstructions contract.
type TeeInstructionsSenderAddressSet struct {
	WalletId      [32]byte
	SenderAddress string
	InitialNonce  uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSenderAddressSet is a free log retrieval operation binding the contract event 0xe0990a7f0cc91ee9137337748e6a3fcb6cdea7499992320c28efd5609e88d23f.
//
// Solidity: event SenderAddressSet(bytes32 indexed walletId, string senderAddress, uint64 initialNonce)
func (_TeeInstructions *TeeInstructionsFilterer) FilterSenderAddressSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeeInstructionsSenderAddressSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeInstructions.contract.FilterLogs(opts, "SenderAddressSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsSenderAddressSetIterator{contract: _TeeInstructions.contract, event: "SenderAddressSet", logs: logs, sub: sub}, nil
}

// WatchSenderAddressSet is a free log subscription operation binding the contract event 0xe0990a7f0cc91ee9137337748e6a3fcb6cdea7499992320c28efd5609e88d23f.
//
// Solidity: event SenderAddressSet(bytes32 indexed walletId, string senderAddress, uint64 initialNonce)
func (_TeeInstructions *TeeInstructionsFilterer) WatchSenderAddressSet(opts *bind.WatchOpts, sink chan<- *TeeInstructionsSenderAddressSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeInstructions.contract.WatchLogs(opts, "SenderAddressSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeInstructionsSenderAddressSet)
				if err := _TeeInstructions.contract.UnpackLog(event, "SenderAddressSet", log); err != nil {
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

// ParseSenderAddressSet is a log parse operation binding the contract event 0xe0990a7f0cc91ee9137337748e6a3fcb6cdea7499992320c28efd5609e88d23f.
//
// Solidity: event SenderAddressSet(bytes32 indexed walletId, string senderAddress, uint64 initialNonce)
func (_TeeInstructions *TeeInstructionsFilterer) ParseSenderAddressSet(log types.Log) (*TeeInstructionsSenderAddressSet, error) {
	event := new(TeeInstructionsSenderAddressSet)
	if err := _TeeInstructions.contract.UnpackLog(event, "SenderAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
