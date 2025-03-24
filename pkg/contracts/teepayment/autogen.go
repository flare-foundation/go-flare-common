// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepayment

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
	PaymentReference [32]byte
}

// TeePaymentMetaData contains all meta data concerning the TeePayment contract.
var TeePaymentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"BatchSettingsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"controlAddress\",\"type\":\"address\"}],\"name\":\"ControlAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"maxFee\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxFeeTolerancePPM\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"maxControlFee\",\"type\":\"uint96\"}],\"name\":\"FeesSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getSenderAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_senderAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletSettings\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint96\",\"name\":\"_maxFee\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"_maxFeeTolerancePPM\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_controlAddress\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_maxControlFee\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePayments.PaymentInstruction\",\"name\":\"_paymentInstruction\",\"type\":\"tuple\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_subNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_firstSubNonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePayments.PaymentInstruction[]\",\"name\":\"_paymentInstruction\",\"type\":\"tuple[]\"},{\"internalType\":\"uint96\",\"name\":\"_fee\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"_nullify\",\"type\":\"bool\"}],\"name\":\"reissue\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"setBatchSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_controlAddress\",\"type\":\"address\"}],\"name\":\"setControlAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint96\",\"name\":\"_maxFee\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"_maxFeeTolerancePPM\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"_maxControlFee\",\"type\":\"uint96\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_senderAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"_initialNonce\",\"type\":\"uint64\"}],\"name\":\"setSenderAddressAndInitialNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeePaymentABI is the input ABI used to generate the binding from.
// Deprecated: Use TeePaymentMetaData.ABI instead.
var TeePaymentABI = TeePaymentMetaData.ABI

// TeePayment is an auto generated Go binding around an Ethereum contract.
type TeePayment struct {
	TeePaymentCaller     // Read-only binding to the contract
	TeePaymentTransactor // Write-only binding to the contract
	TeePaymentFilterer   // Log filterer for contract events
}

// TeePaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeePaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeePaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeePaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeePaymentSession struct {
	Contract     *TeePayment       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeePaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeePaymentCallerSession struct {
	Contract *TeePaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TeePaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeePaymentTransactorSession struct {
	Contract     *TeePaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TeePaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeePaymentRaw struct {
	Contract *TeePayment // Generic contract binding to access the raw methods on
}

// TeePaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeePaymentCallerRaw struct {
	Contract *TeePaymentCaller // Generic read-only contract binding to access the raw methods on
}

// TeePaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeePaymentTransactorRaw struct {
	Contract *TeePaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeePayment creates a new instance of TeePayment, bound to a specific deployed contract.
func NewTeePayment(address common.Address, backend bind.ContractBackend) (*TeePayment, error) {
	contract, err := bindTeePayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeePayment{TeePaymentCaller: TeePaymentCaller{contract: contract}, TeePaymentTransactor: TeePaymentTransactor{contract: contract}, TeePaymentFilterer: TeePaymentFilterer{contract: contract}}, nil
}

// NewTeePaymentCaller creates a new read-only instance of TeePayment, bound to a specific deployed contract.
func NewTeePaymentCaller(address common.Address, caller bind.ContractCaller) (*TeePaymentCaller, error) {
	contract, err := bindTeePayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentCaller{contract: contract}, nil
}

// NewTeePaymentTransactor creates a new write-only instance of TeePayment, bound to a specific deployed contract.
func NewTeePaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*TeePaymentTransactor, error) {
	contract, err := bindTeePayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentTransactor{contract: contract}, nil
}

// NewTeePaymentFilterer creates a new log filterer instance of TeePayment, bound to a specific deployed contract.
func NewTeePaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*TeePaymentFilterer, error) {
	contract, err := bindTeePayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeePaymentFilterer{contract: contract}, nil
}

// bindTeePayment binds a generic wrapper to an already deployed contract.
func bindTeePayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeePaymentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePayment *TeePaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePayment.Contract.TeePaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePayment *TeePaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePayment.Contract.TeePaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePayment *TeePaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePayment.Contract.TeePaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePayment *TeePaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePayment *TeePaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePayment *TeePaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePayment.Contract.contract.Transact(opts, method, params...)
}

// GetSenderAddress is a free data retrieval call binding the contract method 0x63453bba.
//
// Solidity: function getSenderAddress(bytes32 _walletId) view returns(string _senderAddress)
func (_TeePayment *TeePaymentCaller) GetSenderAddress(opts *bind.CallOpts, _walletId [32]byte) (string, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "getSenderAddress", _walletId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSenderAddress is a free data retrieval call binding the contract method 0x63453bba.
//
// Solidity: function getSenderAddress(bytes32 _walletId) view returns(string _senderAddress)
func (_TeePayment *TeePaymentSession) GetSenderAddress(_walletId [32]byte) (string, error) {
	return _TeePayment.Contract.GetSenderAddress(&_TeePayment.CallOpts, _walletId)
}

// GetSenderAddress is a free data retrieval call binding the contract method 0x63453bba.
//
// Solidity: function getSenderAddress(bytes32 _walletId) view returns(string _senderAddress)
func (_TeePayment *TeePaymentCallerSession) GetSenderAddress(_walletId [32]byte) (string, error) {
	return _TeePayment.Contract.GetSenderAddress(&_TeePayment.CallOpts, _walletId)
}

// GetWalletSettings is a free data retrieval call binding the contract method 0x73c164b6.
//
// Solidity: function getWalletSettings(bytes32 _walletId) view returns(uint64 _batchSize, uint64 _batchDurationSeconds, uint96 _maxFee, uint32 _maxFeeTolerancePPM, address _controlAddress, uint96 _maxControlFee)
func (_TeePayment *TeePaymentCaller) GetWalletSettings(opts *bind.CallOpts, _walletId [32]byte) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
	MaxFee               *big.Int
	MaxFeeTolerancePPM   uint32
	ControlAddress       common.Address
	MaxControlFee        *big.Int
}, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "getWalletSettings", _walletId)

	outstruct := new(struct {
		BatchSize            uint64
		BatchDurationSeconds uint64
		MaxFee               *big.Int
		MaxFeeTolerancePPM   uint32
		ControlAddress       common.Address
		MaxControlFee        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchSize = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.BatchDurationSeconds = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.MaxFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxFeeTolerancePPM = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.ControlAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.MaxControlFee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetWalletSettings is a free data retrieval call binding the contract method 0x73c164b6.
//
// Solidity: function getWalletSettings(bytes32 _walletId) view returns(uint64 _batchSize, uint64 _batchDurationSeconds, uint96 _maxFee, uint32 _maxFeeTolerancePPM, address _controlAddress, uint96 _maxControlFee)
func (_TeePayment *TeePaymentSession) GetWalletSettings(_walletId [32]byte) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
	MaxFee               *big.Int
	MaxFeeTolerancePPM   uint32
	ControlAddress       common.Address
	MaxControlFee        *big.Int
}, error) {
	return _TeePayment.Contract.GetWalletSettings(&_TeePayment.CallOpts, _walletId)
}

// GetWalletSettings is a free data retrieval call binding the contract method 0x73c164b6.
//
// Solidity: function getWalletSettings(bytes32 _walletId) view returns(uint64 _batchSize, uint64 _batchDurationSeconds, uint96 _maxFee, uint32 _maxFeeTolerancePPM, address _controlAddress, uint96 _maxControlFee)
func (_TeePayment *TeePaymentCallerSession) GetWalletSettings(_walletId [32]byte) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
	MaxFee               *big.Int
	MaxFeeTolerancePPM   uint32
	ControlAddress       common.Address
	MaxControlFee        *big.Int
}, error) {
	return _TeePayment.Contract.GetWalletSettings(&_TeePayment.CallOpts, _walletId)
}

// Pay is a paid mutator transaction binding the contract method 0x56099959.
//
// Solidity: function pay(bytes32 _walletId, (string,uint256,bytes32) _paymentInstruction) payable returns(uint256 _subNonce)
func (_TeePayment *TeePaymentTransactor) Pay(opts *bind.TransactOpts, _walletId [32]byte, _paymentInstruction ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "pay", _walletId, _paymentInstruction)
}

// Pay is a paid mutator transaction binding the contract method 0x56099959.
//
// Solidity: function pay(bytes32 _walletId, (string,uint256,bytes32) _paymentInstruction) payable returns(uint256 _subNonce)
func (_TeePayment *TeePaymentSession) Pay(_walletId [32]byte, _paymentInstruction ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _TeePayment.Contract.Pay(&_TeePayment.TransactOpts, _walletId, _paymentInstruction)
}

// Pay is a paid mutator transaction binding the contract method 0x56099959.
//
// Solidity: function pay(bytes32 _walletId, (string,uint256,bytes32) _paymentInstruction) payable returns(uint256 _subNonce)
func (_TeePayment *TeePaymentTransactorSession) Pay(_walletId [32]byte, _paymentInstruction ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _TeePayment.Contract.Pay(&_TeePayment.TransactOpts, _walletId, _paymentInstruction)
}

// Reissue is a paid mutator transaction binding the contract method 0xdaf1eafd.
//
// Solidity: function reissue(bytes32 _walletId, uint64 _nonce, uint64 _firstSubNonce, (string,uint256,bytes32)[] _paymentInstruction, uint96 _fee, bool _nullify) payable returns()
func (_TeePayment *TeePaymentTransactor) Reissue(opts *bind.TransactOpts, _walletId [32]byte, _nonce uint64, _firstSubNonce uint64, _paymentInstruction []ITeePaymentsPaymentInstruction, _fee *big.Int, _nullify bool) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "reissue", _walletId, _nonce, _firstSubNonce, _paymentInstruction, _fee, _nullify)
}

// Reissue is a paid mutator transaction binding the contract method 0xdaf1eafd.
//
// Solidity: function reissue(bytes32 _walletId, uint64 _nonce, uint64 _firstSubNonce, (string,uint256,bytes32)[] _paymentInstruction, uint96 _fee, bool _nullify) payable returns()
func (_TeePayment *TeePaymentSession) Reissue(_walletId [32]byte, _nonce uint64, _firstSubNonce uint64, _paymentInstruction []ITeePaymentsPaymentInstruction, _fee *big.Int, _nullify bool) (*types.Transaction, error) {
	return _TeePayment.Contract.Reissue(&_TeePayment.TransactOpts, _walletId, _nonce, _firstSubNonce, _paymentInstruction, _fee, _nullify)
}

// Reissue is a paid mutator transaction binding the contract method 0xdaf1eafd.
//
// Solidity: function reissue(bytes32 _walletId, uint64 _nonce, uint64 _firstSubNonce, (string,uint256,bytes32)[] _paymentInstruction, uint96 _fee, bool _nullify) payable returns()
func (_TeePayment *TeePaymentTransactorSession) Reissue(_walletId [32]byte, _nonce uint64, _firstSubNonce uint64, _paymentInstruction []ITeePaymentsPaymentInstruction, _fee *big.Int, _nullify bool) (*types.Transaction, error) {
	return _TeePayment.Contract.Reissue(&_TeePayment.TransactOpts, _walletId, _nonce, _firstSubNonce, _paymentInstruction, _fee, _nullify)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xf680cbad.
//
// Solidity: function setBatchSettings(bytes32 _walletId, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePayment *TeePaymentTransactor) SetBatchSettings(opts *bind.TransactOpts, _walletId [32]byte, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "setBatchSettings", _walletId, _batchSize, _batchDurationSeconds)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xf680cbad.
//
// Solidity: function setBatchSettings(bytes32 _walletId, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePayment *TeePaymentSession) SetBatchSettings(_walletId [32]byte, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePayment.Contract.SetBatchSettings(&_TeePayment.TransactOpts, _walletId, _batchSize, _batchDurationSeconds)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xf680cbad.
//
// Solidity: function setBatchSettings(bytes32 _walletId, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePayment *TeePaymentTransactorSession) SetBatchSettings(_walletId [32]byte, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePayment.Contract.SetBatchSettings(&_TeePayment.TransactOpts, _walletId, _batchSize, _batchDurationSeconds)
}

// SetControlAddress is a paid mutator transaction binding the contract method 0x1ae659c8.
//
// Solidity: function setControlAddress(bytes32 _walletId, address _controlAddress) returns()
func (_TeePayment *TeePaymentTransactor) SetControlAddress(opts *bind.TransactOpts, _walletId [32]byte, _controlAddress common.Address) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "setControlAddress", _walletId, _controlAddress)
}

// SetControlAddress is a paid mutator transaction binding the contract method 0x1ae659c8.
//
// Solidity: function setControlAddress(bytes32 _walletId, address _controlAddress) returns()
func (_TeePayment *TeePaymentSession) SetControlAddress(_walletId [32]byte, _controlAddress common.Address) (*types.Transaction, error) {
	return _TeePayment.Contract.SetControlAddress(&_TeePayment.TransactOpts, _walletId, _controlAddress)
}

// SetControlAddress is a paid mutator transaction binding the contract method 0x1ae659c8.
//
// Solidity: function setControlAddress(bytes32 _walletId, address _controlAddress) returns()
func (_TeePayment *TeePaymentTransactorSession) SetControlAddress(_walletId [32]byte, _controlAddress common.Address) (*types.Transaction, error) {
	return _TeePayment.Contract.SetControlAddress(&_TeePayment.TransactOpts, _walletId, _controlAddress)
}

// SetFees is a paid mutator transaction binding the contract method 0xbe2ab5f4.
//
// Solidity: function setFees(bytes32 _walletId, uint96 _maxFee, uint32 _maxFeeTolerancePPM, uint96 _maxControlFee) returns()
func (_TeePayment *TeePaymentTransactor) SetFees(opts *bind.TransactOpts, _walletId [32]byte, _maxFee *big.Int, _maxFeeTolerancePPM uint32, _maxControlFee *big.Int) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "setFees", _walletId, _maxFee, _maxFeeTolerancePPM, _maxControlFee)
}

// SetFees is a paid mutator transaction binding the contract method 0xbe2ab5f4.
//
// Solidity: function setFees(bytes32 _walletId, uint96 _maxFee, uint32 _maxFeeTolerancePPM, uint96 _maxControlFee) returns()
func (_TeePayment *TeePaymentSession) SetFees(_walletId [32]byte, _maxFee *big.Int, _maxFeeTolerancePPM uint32, _maxControlFee *big.Int) (*types.Transaction, error) {
	return _TeePayment.Contract.SetFees(&_TeePayment.TransactOpts, _walletId, _maxFee, _maxFeeTolerancePPM, _maxControlFee)
}

// SetFees is a paid mutator transaction binding the contract method 0xbe2ab5f4.
//
// Solidity: function setFees(bytes32 _walletId, uint96 _maxFee, uint32 _maxFeeTolerancePPM, uint96 _maxControlFee) returns()
func (_TeePayment *TeePaymentTransactorSession) SetFees(_walletId [32]byte, _maxFee *big.Int, _maxFeeTolerancePPM uint32, _maxControlFee *big.Int) (*types.Transaction, error) {
	return _TeePayment.Contract.SetFees(&_TeePayment.TransactOpts, _walletId, _maxFee, _maxFeeTolerancePPM, _maxControlFee)
}

// SetSenderAddressAndInitialNonce is a paid mutator transaction binding the contract method 0xb9a853c3.
//
// Solidity: function setSenderAddressAndInitialNonce(bytes32 _walletId, string _senderAddress, uint64 _initialNonce) returns()
func (_TeePayment *TeePaymentTransactor) SetSenderAddressAndInitialNonce(opts *bind.TransactOpts, _walletId [32]byte, _senderAddress string, _initialNonce uint64) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "setSenderAddressAndInitialNonce", _walletId, _senderAddress, _initialNonce)
}

// SetSenderAddressAndInitialNonce is a paid mutator transaction binding the contract method 0xb9a853c3.
//
// Solidity: function setSenderAddressAndInitialNonce(bytes32 _walletId, string _senderAddress, uint64 _initialNonce) returns()
func (_TeePayment *TeePaymentSession) SetSenderAddressAndInitialNonce(_walletId [32]byte, _senderAddress string, _initialNonce uint64) (*types.Transaction, error) {
	return _TeePayment.Contract.SetSenderAddressAndInitialNonce(&_TeePayment.TransactOpts, _walletId, _senderAddress, _initialNonce)
}

// SetSenderAddressAndInitialNonce is a paid mutator transaction binding the contract method 0xb9a853c3.
//
// Solidity: function setSenderAddressAndInitialNonce(bytes32 _walletId, string _senderAddress, uint64 _initialNonce) returns()
func (_TeePayment *TeePaymentTransactorSession) SetSenderAddressAndInitialNonce(_walletId [32]byte, _senderAddress string, _initialNonce uint64) (*types.Transaction, error) {
	return _TeePayment.Contract.SetSenderAddressAndInitialNonce(&_TeePayment.TransactOpts, _walletId, _senderAddress, _initialNonce)
}

// TeePaymentBatchSettingsSetIterator is returned from FilterBatchSettingsSet and is used to iterate over the raw logs and unpacked data for BatchSettingsSet events raised by the TeePayment contract.
type TeePaymentBatchSettingsSetIterator struct {
	Event *TeePaymentBatchSettingsSet // Event containing the contract specifics and raw log

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
func (it *TeePaymentBatchSettingsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentBatchSettingsSet)
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
		it.Event = new(TeePaymentBatchSettingsSet)
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
func (it *TeePaymentBatchSettingsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentBatchSettingsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentBatchSettingsSet represents a BatchSettingsSet event raised by the TeePayment contract.
type TeePaymentBatchSettingsSet struct {
	WalletId             [32]byte
	BatchSize            uint64
	BatchDurationSeconds uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBatchSettingsSet is a free log retrieval operation binding the contract event 0xfa2da2b02c35a83b9a69ec27c4d6fbf1c4cf31d416e5783d45a39bdfbc36deff.
//
// Solidity: event BatchSettingsSet(bytes32 indexed walletId, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeePayment *TeePaymentFilterer) FilterBatchSettingsSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeePaymentBatchSettingsSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "BatchSettingsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentBatchSettingsSetIterator{contract: _TeePayment.contract, event: "BatchSettingsSet", logs: logs, sub: sub}, nil
}

// WatchBatchSettingsSet is a free log subscription operation binding the contract event 0xfa2da2b02c35a83b9a69ec27c4d6fbf1c4cf31d416e5783d45a39bdfbc36deff.
//
// Solidity: event BatchSettingsSet(bytes32 indexed walletId, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeePayment *TeePaymentFilterer) WatchBatchSettingsSet(opts *bind.WatchOpts, sink chan<- *TeePaymentBatchSettingsSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "BatchSettingsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentBatchSettingsSet)
				if err := _TeePayment.contract.UnpackLog(event, "BatchSettingsSet", log); err != nil {
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
func (_TeePayment *TeePaymentFilterer) ParseBatchSettingsSet(log types.Log) (*TeePaymentBatchSettingsSet, error) {
	event := new(TeePaymentBatchSettingsSet)
	if err := _TeePayment.contract.UnpackLog(event, "BatchSettingsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentControlAddressSetIterator is returned from FilterControlAddressSet and is used to iterate over the raw logs and unpacked data for ControlAddressSet events raised by the TeePayment contract.
type TeePaymentControlAddressSetIterator struct {
	Event *TeePaymentControlAddressSet // Event containing the contract specifics and raw log

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
func (it *TeePaymentControlAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentControlAddressSet)
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
		it.Event = new(TeePaymentControlAddressSet)
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
func (it *TeePaymentControlAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentControlAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentControlAddressSet represents a ControlAddressSet event raised by the TeePayment contract.
type TeePaymentControlAddressSet struct {
	WalletId       [32]byte
	ControlAddress common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterControlAddressSet is a free log retrieval operation binding the contract event 0x996301706402921996c1da682c26acf8bfe8381043ca0c62636312f6ced8e44a.
//
// Solidity: event ControlAddressSet(bytes32 indexed walletId, address controlAddress)
func (_TeePayment *TeePaymentFilterer) FilterControlAddressSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeePaymentControlAddressSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "ControlAddressSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentControlAddressSetIterator{contract: _TeePayment.contract, event: "ControlAddressSet", logs: logs, sub: sub}, nil
}

// WatchControlAddressSet is a free log subscription operation binding the contract event 0x996301706402921996c1da682c26acf8bfe8381043ca0c62636312f6ced8e44a.
//
// Solidity: event ControlAddressSet(bytes32 indexed walletId, address controlAddress)
func (_TeePayment *TeePaymentFilterer) WatchControlAddressSet(opts *bind.WatchOpts, sink chan<- *TeePaymentControlAddressSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "ControlAddressSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentControlAddressSet)
				if err := _TeePayment.contract.UnpackLog(event, "ControlAddressSet", log); err != nil {
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

// ParseControlAddressSet is a log parse operation binding the contract event 0x996301706402921996c1da682c26acf8bfe8381043ca0c62636312f6ced8e44a.
//
// Solidity: event ControlAddressSet(bytes32 indexed walletId, address controlAddress)
func (_TeePayment *TeePaymentFilterer) ParseControlAddressSet(log types.Log) (*TeePaymentControlAddressSet, error) {
	event := new(TeePaymentControlAddressSet)
	if err := _TeePayment.contract.UnpackLog(event, "ControlAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentFeesSetIterator is returned from FilterFeesSet and is used to iterate over the raw logs and unpacked data for FeesSet events raised by the TeePayment contract.
type TeePaymentFeesSetIterator struct {
	Event *TeePaymentFeesSet // Event containing the contract specifics and raw log

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
func (it *TeePaymentFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentFeesSet)
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
		it.Event = new(TeePaymentFeesSet)
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
func (it *TeePaymentFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentFeesSet represents a FeesSet event raised by the TeePayment contract.
type TeePaymentFeesSet struct {
	WalletId           [32]byte
	MaxFee             *big.Int
	MaxFeeTolerancePPM uint32
	MaxControlFee      *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFeesSet is a free log retrieval operation binding the contract event 0x6104aa749f5a8c1d1c75a44525f12df2f0127c671fcfb9fab4afafb1f46d4c40.
//
// Solidity: event FeesSet(bytes32 indexed walletId, uint96 maxFee, uint32 maxFeeTolerancePPM, uint96 maxControlFee)
func (_TeePayment *TeePaymentFilterer) FilterFeesSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeePaymentFeesSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "FeesSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentFeesSetIterator{contract: _TeePayment.contract, event: "FeesSet", logs: logs, sub: sub}, nil
}

// WatchFeesSet is a free log subscription operation binding the contract event 0x6104aa749f5a8c1d1c75a44525f12df2f0127c671fcfb9fab4afafb1f46d4c40.
//
// Solidity: event FeesSet(bytes32 indexed walletId, uint96 maxFee, uint32 maxFeeTolerancePPM, uint96 maxControlFee)
func (_TeePayment *TeePaymentFilterer) WatchFeesSet(opts *bind.WatchOpts, sink chan<- *TeePaymentFeesSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "FeesSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentFeesSet)
				if err := _TeePayment.contract.UnpackLog(event, "FeesSet", log); err != nil {
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

// ParseFeesSet is a log parse operation binding the contract event 0x6104aa749f5a8c1d1c75a44525f12df2f0127c671fcfb9fab4afafb1f46d4c40.
//
// Solidity: event FeesSet(bytes32 indexed walletId, uint96 maxFee, uint32 maxFeeTolerancePPM, uint96 maxControlFee)
func (_TeePayment *TeePaymentFilterer) ParseFeesSet(log types.Log) (*TeePaymentFeesSet, error) {
	event := new(TeePaymentFeesSet)
	if err := _TeePayment.contract.UnpackLog(event, "FeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
