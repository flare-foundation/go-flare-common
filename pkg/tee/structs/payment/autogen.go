// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package payment

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

// ITeePaymentsLimitsManagerSetPaymentLimitsMessage is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsLimitsManagerSetPaymentLimitsMessage struct {
	WalletId         [32]byte
	SourceId         [32]byte
	AccountAddress   string
	Nonce            *big.Int
	TeeIdKeyIdPairs  []TeeIdKeyIdPair
	TransactionLimit *big.Int
	DailyLimit       *big.Int
}

// ITeePaymentsPMWMultisigAccount is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsPMWMultisigAccount struct {
	SourceId       [32]byte
	AccountAddress string
}

// ITeePaymentsPaymentInstruction is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsPaymentInstruction struct {
	RecipientAddress string
	TokenId          []byte
	Amount           *big.Int
	MaxFee           *big.Int
	PaymentReference [32]byte
}

// ITeePaymentsPaymentInstructionMessage is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsPaymentInstructionMessage struct {
	WalletId         [32]byte
	TeeIdKeyIdPairs  []TeeIdKeyIdPair
	SourceId         [32]byte
	SenderAddress    string
	RecipientAddress string
	TokenId          []byte
	Amount           *big.Int
	MaxFee           *big.Int
	FeeSchedule      []byte
	PaymentReference [32]byte
	Nonce            uint64
	SubNonce         uint64
	BatchEndTs       uint64
}

// ITeePaymentsReissueFeeParams is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsReissueFeeParams struct {
	MaxFeePerPayment      []*big.Int
	FactorsBIPSPerPayment [][]int16
	DelaysSeconds         []uint16
}

// TeeIdKeyIdPair is an auto generated low-level Go binding around an user-defined struct.
type TeeIdKeyIdPair struct {
	TeeId common.Address
	KeyId uint64
}

// PaymentMetaData contains all meta data concerning the Payment contract.
var PaymentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"internalType\":\"structTeeIdKeyIdPair[]\",\"name\":\"teeIdKeyIdPairs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeSchedule\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchEndTs\",\"type\":\"uint64\"}],\"internalType\":\"structITeePayments.PaymentInstructionMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"paymentInstructionMessageStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePayments.PaymentInstruction\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"paymentInstructionStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwMultisigAccountStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"maxFeePerPayment\",\"type\":\"uint256[]\"},{\"internalType\":\"int16[][]\",\"name\":\"factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"delaysSeconds\",\"type\":\"uint16[]\"}],\"internalType\":\"structITeePayments.ReissueFeeParams\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"reissueFeeParamsStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"internalType\":\"structTeeIdKeyIdPair[]\",\"name\":\"teeIdKeyIdPairs\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"transactionLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dailyLimit\",\"type\":\"uint256\"}],\"internalType\":\"structITeePaymentsLimitsManager.SetPaymentLimitsMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"setPaymentLimitsStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PaymentABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymentMetaData.ABI instead.
var PaymentABI = PaymentMetaData.ABI

// Payment is an auto generated Go binding around an Ethereum contract.
type Payment struct {
	PaymentCaller     // Read-only binding to the contract
	PaymentTransactor // Write-only binding to the contract
	PaymentFilterer   // Log filterer for contract events
}

// PaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentSession struct {
	Contract     *Payment          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentCallerSession struct {
	Contract *PaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentTransactorSession struct {
	Contract     *PaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentRaw struct {
	Contract *Payment // Generic contract binding to access the raw methods on
}

// PaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentCallerRaw struct {
	Contract *PaymentCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentTransactorRaw struct {
	Contract *PaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayment creates a new instance of Payment, bound to a specific deployed contract.
func NewPayment(address common.Address, backend bind.ContractBackend) (*Payment, error) {
	contract, err := bindPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Payment{PaymentCaller: PaymentCaller{contract: contract}, PaymentTransactor: PaymentTransactor{contract: contract}, PaymentFilterer: PaymentFilterer{contract: contract}}, nil
}

// NewPaymentCaller creates a new read-only instance of Payment, bound to a specific deployed contract.
func NewPaymentCaller(address common.Address, caller bind.ContractCaller) (*PaymentCaller, error) {
	contract, err := bindPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentCaller{contract: contract}, nil
}

// NewPaymentTransactor creates a new write-only instance of Payment, bound to a specific deployed contract.
func NewPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentTransactor, error) {
	contract, err := bindPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentTransactor{contract: contract}, nil
}

// NewPaymentFilterer creates a new log filterer instance of Payment, bound to a specific deployed contract.
func NewPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentFilterer, error) {
	contract, err := bindPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentFilterer{contract: contract}, nil
}

// bindPayment binds a generic wrapper to an already deployed contract.
func bindPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payment *PaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Payment.Contract.PaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payment *PaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.Contract.PaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payment *PaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payment.Contract.PaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payment *PaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Payment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payment *PaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payment *PaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payment.Contract.contract.Transact(opts, method, params...)
}

// PaymentInstructionMessageStruct is a paid mutator transaction binding the contract method 0x97a864e9.
//
// Solidity: function paymentInstructionMessageStruct((bytes32,(address,uint64)[],bytes32,string,string,bytes,uint256,uint256,bytes,bytes32,uint64,uint64,uint64) ) returns()
func (_Payment *PaymentTransactor) PaymentInstructionMessageStruct(opts *bind.TransactOpts, arg0 ITeePaymentsPaymentInstructionMessage) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "paymentInstructionMessageStruct", arg0)
}

// PaymentInstructionMessageStruct is a paid mutator transaction binding the contract method 0x97a864e9.
//
// Solidity: function paymentInstructionMessageStruct((bytes32,(address,uint64)[],bytes32,string,string,bytes,uint256,uint256,bytes,bytes32,uint64,uint64,uint64) ) returns()
func (_Payment *PaymentSession) PaymentInstructionMessageStruct(arg0 ITeePaymentsPaymentInstructionMessage) (*types.Transaction, error) {
	return _Payment.Contract.PaymentInstructionMessageStruct(&_Payment.TransactOpts, arg0)
}

// PaymentInstructionMessageStruct is a paid mutator transaction binding the contract method 0x97a864e9.
//
// Solidity: function paymentInstructionMessageStruct((bytes32,(address,uint64)[],bytes32,string,string,bytes,uint256,uint256,bytes,bytes32,uint64,uint64,uint64) ) returns()
func (_Payment *PaymentTransactorSession) PaymentInstructionMessageStruct(arg0 ITeePaymentsPaymentInstructionMessage) (*types.Transaction, error) {
	return _Payment.Contract.PaymentInstructionMessageStruct(&_Payment.TransactOpts, arg0)
}

// PaymentInstructionStruct is a paid mutator transaction binding the contract method 0x3ef0b061.
//
// Solidity: function paymentInstructionStruct((string,bytes,uint256,uint256,bytes32) ) returns()
func (_Payment *PaymentTransactor) PaymentInstructionStruct(opts *bind.TransactOpts, arg0 ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "paymentInstructionStruct", arg0)
}

// PaymentInstructionStruct is a paid mutator transaction binding the contract method 0x3ef0b061.
//
// Solidity: function paymentInstructionStruct((string,bytes,uint256,uint256,bytes32) ) returns()
func (_Payment *PaymentSession) PaymentInstructionStruct(arg0 ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _Payment.Contract.PaymentInstructionStruct(&_Payment.TransactOpts, arg0)
}

// PaymentInstructionStruct is a paid mutator transaction binding the contract method 0x3ef0b061.
//
// Solidity: function paymentInstructionStruct((string,bytes,uint256,uint256,bytes32) ) returns()
func (_Payment *PaymentTransactorSession) PaymentInstructionStruct(arg0 ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _Payment.Contract.PaymentInstructionStruct(&_Payment.TransactOpts, arg0)
}

// PmwMultisigAccountStruct is a paid mutator transaction binding the contract method 0x8d06b168.
//
// Solidity: function pmwMultisigAccountStruct((bytes32,string) ) returns()
func (_Payment *PaymentTransactor) PmwMultisigAccountStruct(opts *bind.TransactOpts, arg0 ITeePaymentsPMWMultisigAccount) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "pmwMultisigAccountStruct", arg0)
}

// PmwMultisigAccountStruct is a paid mutator transaction binding the contract method 0x8d06b168.
//
// Solidity: function pmwMultisigAccountStruct((bytes32,string) ) returns()
func (_Payment *PaymentSession) PmwMultisigAccountStruct(arg0 ITeePaymentsPMWMultisigAccount) (*types.Transaction, error) {
	return _Payment.Contract.PmwMultisigAccountStruct(&_Payment.TransactOpts, arg0)
}

// PmwMultisigAccountStruct is a paid mutator transaction binding the contract method 0x8d06b168.
//
// Solidity: function pmwMultisigAccountStruct((bytes32,string) ) returns()
func (_Payment *PaymentTransactorSession) PmwMultisigAccountStruct(arg0 ITeePaymentsPMWMultisigAccount) (*types.Transaction, error) {
	return _Payment.Contract.PmwMultisigAccountStruct(&_Payment.TransactOpts, arg0)
}

// ReissueFeeParamsStruct is a paid mutator transaction binding the contract method 0xc9a7c35e.
//
// Solidity: function reissueFeeParamsStruct((uint256[],int16[][],uint16[]) ) returns()
func (_Payment *PaymentTransactor) ReissueFeeParamsStruct(opts *bind.TransactOpts, arg0 ITeePaymentsReissueFeeParams) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "reissueFeeParamsStruct", arg0)
}

// ReissueFeeParamsStruct is a paid mutator transaction binding the contract method 0xc9a7c35e.
//
// Solidity: function reissueFeeParamsStruct((uint256[],int16[][],uint16[]) ) returns()
func (_Payment *PaymentSession) ReissueFeeParamsStruct(arg0 ITeePaymentsReissueFeeParams) (*types.Transaction, error) {
	return _Payment.Contract.ReissueFeeParamsStruct(&_Payment.TransactOpts, arg0)
}

// ReissueFeeParamsStruct is a paid mutator transaction binding the contract method 0xc9a7c35e.
//
// Solidity: function reissueFeeParamsStruct((uint256[],int16[][],uint16[]) ) returns()
func (_Payment *PaymentTransactorSession) ReissueFeeParamsStruct(arg0 ITeePaymentsReissueFeeParams) (*types.Transaction, error) {
	return _Payment.Contract.ReissueFeeParamsStruct(&_Payment.TransactOpts, arg0)
}

// SetPaymentLimitsStruct is a paid mutator transaction binding the contract method 0xfdfc8892.
//
// Solidity: function setPaymentLimitsStruct((bytes32,bytes32,string,uint256,(address,uint64)[],uint256,uint256) ) returns()
func (_Payment *PaymentTransactor) SetPaymentLimitsStruct(opts *bind.TransactOpts, arg0 ITeePaymentsLimitsManagerSetPaymentLimitsMessage) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "setPaymentLimitsStruct", arg0)
}

// SetPaymentLimitsStruct is a paid mutator transaction binding the contract method 0xfdfc8892.
//
// Solidity: function setPaymentLimitsStruct((bytes32,bytes32,string,uint256,(address,uint64)[],uint256,uint256) ) returns()
func (_Payment *PaymentSession) SetPaymentLimitsStruct(arg0 ITeePaymentsLimitsManagerSetPaymentLimitsMessage) (*types.Transaction, error) {
	return _Payment.Contract.SetPaymentLimitsStruct(&_Payment.TransactOpts, arg0)
}

// SetPaymentLimitsStruct is a paid mutator transaction binding the contract method 0xfdfc8892.
//
// Solidity: function setPaymentLimitsStruct((bytes32,bytes32,string,uint256,(address,uint64)[],uint256,uint256) ) returns()
func (_Payment *PaymentTransactorSession) SetPaymentLimitsStruct(arg0 ITeePaymentsLimitsManagerSetPaymentLimitsMessage) (*types.Transaction, error) {
	return _Payment.Contract.SetPaymentLimitsStruct(&_Payment.TransactOpts, arg0)
}
