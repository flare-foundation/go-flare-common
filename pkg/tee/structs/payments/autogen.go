// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package payments

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

// ITeePaymentsBasePMWMultisigAccount is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsBasePMWMultisigAccount struct {
	SourceId       [32]byte
	AccountAddress string
}

// ITeePaymentsBasePaymentInstruction is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsBasePaymentInstruction struct {
	RecipientAddress string
	TokenId          []byte
	Amount           *big.Int
	MaxFee           *big.Int
	PaymentReference [32]byte
}

// ITeePaymentsBaseReissueFeeParams is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsBaseReissueFeeParams struct {
	MaxFeePerPayment      []*big.Int
	FactorsBIPSPerPayment [][]int16
	DelaysSeconds         []uint16
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
	PaymentId        uint64
}

// ITeePaymentsUtxoUtxoAnchorState is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsUtxoUtxoAnchorState struct {
	GenesisAnchorTxid [32]byte
	GenesisAnchorVout uint32
	NextNonce         uint64
	AvailableAt       uint64
}

// ITeePaymentsUtxoUtxoPaymentInstructionMessage is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsUtxoUtxoPaymentInstructionMessage struct {
	WalletId         [32]byte
	TeeIdKeyIdPairs  []TeeIdKeyIdPair
	SourceId         [32]byte
	AccountAddress   string
	AccountIndex     uint32
	AnchorIndex      uint32
	RecipientAddress string
	TokenId          []byte
	Amount           *big.Int
	MaxFee           *big.Int
	FeeSchedule      []byte
	PaymentReference [32]byte
	Nonce            uint64
	PaymentId        uint64
	BatchPaymentId   uint64
	BatchEndTs       uint64
}

// TeeIdKeyIdPair is an auto generated low-level Go binding around an user-defined struct.
type TeeIdKeyIdPair struct {
	TeeId common.Address
	KeyId uint64
}

// TeePaymentsMetaData contains all meta data concerning the TeePayments contract.
var TeePaymentsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"internalType\":\"structTeeIdKeyIdPair[]\",\"name\":\"teeIdKeyIdPairs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeSchedule\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"}],\"internalType\":\"structITeePayments.PaymentInstructionMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"paymentInstructionMessageStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePaymentsBase.PaymentInstruction\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"paymentInstructionStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwMultisigAccountStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"maxFeePerPayment\",\"type\":\"uint256[]\"},{\"internalType\":\"int16[][]\",\"name\":\"factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"delaysSeconds\",\"type\":\"uint16[]\"}],\"internalType\":\"structITeePaymentsBase.ReissueFeeParams\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"reissueFeeParamsStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"nextNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"availableAt\",\"type\":\"uint64\"}],\"internalType\":\"structITeePaymentsUtxo.UtxoAnchorState\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"utxoAnchorStateStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"internalType\":\"structTeeIdKeyIdPair[]\",\"name\":\"teeIdKeyIdPairs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"anchorIndex\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"feeSchedule\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchEndTs\",\"type\":\"uint64\"}],\"internalType\":\"structITeePaymentsUtxo.UtxoPaymentInstructionMessage\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"utxoPaymentInstructionMessageStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeePaymentsABI is the input ABI used to generate the binding from.
// Deprecated: Use TeePaymentsMetaData.ABI instead.
var TeePaymentsABI = TeePaymentsMetaData.ABI

// TeePayments is an auto generated Go binding around an Ethereum contract.
type TeePayments struct {
	TeePaymentsCaller     // Read-only binding to the contract
	TeePaymentsTransactor // Write-only binding to the contract
	TeePaymentsFilterer   // Log filterer for contract events
}

// TeePaymentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeePaymentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeePaymentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeePaymentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeePaymentsSession struct {
	Contract     *TeePayments      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeePaymentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeePaymentsCallerSession struct {
	Contract *TeePaymentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TeePaymentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeePaymentsTransactorSession struct {
	Contract     *TeePaymentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TeePaymentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeePaymentsRaw struct {
	Contract *TeePayments // Generic contract binding to access the raw methods on
}

// TeePaymentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeePaymentsCallerRaw struct {
	Contract *TeePaymentsCaller // Generic read-only contract binding to access the raw methods on
}

// TeePaymentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeePaymentsTransactorRaw struct {
	Contract *TeePaymentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeePayments creates a new instance of TeePayments, bound to a specific deployed contract.
func NewTeePayments(address common.Address, backend bind.ContractBackend) (*TeePayments, error) {
	contract, err := bindTeePayments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeePayments{TeePaymentsCaller: TeePaymentsCaller{contract: contract}, TeePaymentsTransactor: TeePaymentsTransactor{contract: contract}, TeePaymentsFilterer: TeePaymentsFilterer{contract: contract}}, nil
}

// NewTeePaymentsCaller creates a new read-only instance of TeePayments, bound to a specific deployed contract.
func NewTeePaymentsCaller(address common.Address, caller bind.ContractCaller) (*TeePaymentsCaller, error) {
	contract, err := bindTeePayments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsCaller{contract: contract}, nil
}

// NewTeePaymentsTransactor creates a new write-only instance of TeePayments, bound to a specific deployed contract.
func NewTeePaymentsTransactor(address common.Address, transactor bind.ContractTransactor) (*TeePaymentsTransactor, error) {
	contract, err := bindTeePayments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsTransactor{contract: contract}, nil
}

// NewTeePaymentsFilterer creates a new log filterer instance of TeePayments, bound to a specific deployed contract.
func NewTeePaymentsFilterer(address common.Address, filterer bind.ContractFilterer) (*TeePaymentsFilterer, error) {
	contract, err := bindTeePayments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFilterer{contract: contract}, nil
}

// bindTeePayments binds a generic wrapper to an already deployed contract.
func bindTeePayments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeePaymentsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePayments *TeePaymentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePayments.Contract.TeePaymentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePayments *TeePaymentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePayments.Contract.TeePaymentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePayments *TeePaymentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePayments.Contract.TeePaymentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePayments *TeePaymentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePayments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePayments *TeePaymentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePayments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePayments *TeePaymentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePayments.Contract.contract.Transact(opts, method, params...)
}

// PaymentInstructionMessageStruct is a paid mutator transaction binding the contract method 0x4b311981.
//
// Solidity: function paymentInstructionMessageStruct((bytes32,(address,uint64)[],bytes32,string,string,bytes,uint256,uint256,bytes,bytes32,uint64,uint64) ) returns()
func (_TeePayments *TeePaymentsTransactor) PaymentInstructionMessageStruct(opts *bind.TransactOpts, arg0 ITeePaymentsPaymentInstructionMessage) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "paymentInstructionMessageStruct", arg0)
}

// PaymentInstructionMessageStruct is a paid mutator transaction binding the contract method 0x4b311981.
//
// Solidity: function paymentInstructionMessageStruct((bytes32,(address,uint64)[],bytes32,string,string,bytes,uint256,uint256,bytes,bytes32,uint64,uint64) ) returns()
func (_TeePayments *TeePaymentsSession) PaymentInstructionMessageStruct(arg0 ITeePaymentsPaymentInstructionMessage) (*types.Transaction, error) {
	return _TeePayments.Contract.PaymentInstructionMessageStruct(&_TeePayments.TransactOpts, arg0)
}

// PaymentInstructionMessageStruct is a paid mutator transaction binding the contract method 0x4b311981.
//
// Solidity: function paymentInstructionMessageStruct((bytes32,(address,uint64)[],bytes32,string,string,bytes,uint256,uint256,bytes,bytes32,uint64,uint64) ) returns()
func (_TeePayments *TeePaymentsTransactorSession) PaymentInstructionMessageStruct(arg0 ITeePaymentsPaymentInstructionMessage) (*types.Transaction, error) {
	return _TeePayments.Contract.PaymentInstructionMessageStruct(&_TeePayments.TransactOpts, arg0)
}

// PaymentInstructionStruct is a paid mutator transaction binding the contract method 0x3ef0b061.
//
// Solidity: function paymentInstructionStruct((string,bytes,uint256,uint256,bytes32) ) returns()
func (_TeePayments *TeePaymentsTransactor) PaymentInstructionStruct(opts *bind.TransactOpts, arg0 ITeePaymentsBasePaymentInstruction) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "paymentInstructionStruct", arg0)
}

// PaymentInstructionStruct is a paid mutator transaction binding the contract method 0x3ef0b061.
//
// Solidity: function paymentInstructionStruct((string,bytes,uint256,uint256,bytes32) ) returns()
func (_TeePayments *TeePaymentsSession) PaymentInstructionStruct(arg0 ITeePaymentsBasePaymentInstruction) (*types.Transaction, error) {
	return _TeePayments.Contract.PaymentInstructionStruct(&_TeePayments.TransactOpts, arg0)
}

// PaymentInstructionStruct is a paid mutator transaction binding the contract method 0x3ef0b061.
//
// Solidity: function paymentInstructionStruct((string,bytes,uint256,uint256,bytes32) ) returns()
func (_TeePayments *TeePaymentsTransactorSession) PaymentInstructionStruct(arg0 ITeePaymentsBasePaymentInstruction) (*types.Transaction, error) {
	return _TeePayments.Contract.PaymentInstructionStruct(&_TeePayments.TransactOpts, arg0)
}

// PmwMultisigAccountStruct is a paid mutator transaction binding the contract method 0x8d06b168.
//
// Solidity: function pmwMultisigAccountStruct((bytes32,string) ) returns()
func (_TeePayments *TeePaymentsTransactor) PmwMultisigAccountStruct(opts *bind.TransactOpts, arg0 ITeePaymentsBasePMWMultisigAccount) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "pmwMultisigAccountStruct", arg0)
}

// PmwMultisigAccountStruct is a paid mutator transaction binding the contract method 0x8d06b168.
//
// Solidity: function pmwMultisigAccountStruct((bytes32,string) ) returns()
func (_TeePayments *TeePaymentsSession) PmwMultisigAccountStruct(arg0 ITeePaymentsBasePMWMultisigAccount) (*types.Transaction, error) {
	return _TeePayments.Contract.PmwMultisigAccountStruct(&_TeePayments.TransactOpts, arg0)
}

// PmwMultisigAccountStruct is a paid mutator transaction binding the contract method 0x8d06b168.
//
// Solidity: function pmwMultisigAccountStruct((bytes32,string) ) returns()
func (_TeePayments *TeePaymentsTransactorSession) PmwMultisigAccountStruct(arg0 ITeePaymentsBasePMWMultisigAccount) (*types.Transaction, error) {
	return _TeePayments.Contract.PmwMultisigAccountStruct(&_TeePayments.TransactOpts, arg0)
}

// ReissueFeeParamsStruct is a paid mutator transaction binding the contract method 0xc9a7c35e.
//
// Solidity: function reissueFeeParamsStruct((uint256[],int16[][],uint16[]) ) returns()
func (_TeePayments *TeePaymentsTransactor) ReissueFeeParamsStruct(opts *bind.TransactOpts, arg0 ITeePaymentsBaseReissueFeeParams) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "reissueFeeParamsStruct", arg0)
}

// ReissueFeeParamsStruct is a paid mutator transaction binding the contract method 0xc9a7c35e.
//
// Solidity: function reissueFeeParamsStruct((uint256[],int16[][],uint16[]) ) returns()
func (_TeePayments *TeePaymentsSession) ReissueFeeParamsStruct(arg0 ITeePaymentsBaseReissueFeeParams) (*types.Transaction, error) {
	return _TeePayments.Contract.ReissueFeeParamsStruct(&_TeePayments.TransactOpts, arg0)
}

// ReissueFeeParamsStruct is a paid mutator transaction binding the contract method 0xc9a7c35e.
//
// Solidity: function reissueFeeParamsStruct((uint256[],int16[][],uint16[]) ) returns()
func (_TeePayments *TeePaymentsTransactorSession) ReissueFeeParamsStruct(arg0 ITeePaymentsBaseReissueFeeParams) (*types.Transaction, error) {
	return _TeePayments.Contract.ReissueFeeParamsStruct(&_TeePayments.TransactOpts, arg0)
}

// UtxoAnchorStateStruct is a paid mutator transaction binding the contract method 0x8ae285d3.
//
// Solidity: function utxoAnchorStateStruct((bytes32,uint32,uint64,uint64) ) returns()
func (_TeePayments *TeePaymentsTransactor) UtxoAnchorStateStruct(opts *bind.TransactOpts, arg0 ITeePaymentsUtxoUtxoAnchorState) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "utxoAnchorStateStruct", arg0)
}

// UtxoAnchorStateStruct is a paid mutator transaction binding the contract method 0x8ae285d3.
//
// Solidity: function utxoAnchorStateStruct((bytes32,uint32,uint64,uint64) ) returns()
func (_TeePayments *TeePaymentsSession) UtxoAnchorStateStruct(arg0 ITeePaymentsUtxoUtxoAnchorState) (*types.Transaction, error) {
	return _TeePayments.Contract.UtxoAnchorStateStruct(&_TeePayments.TransactOpts, arg0)
}

// UtxoAnchorStateStruct is a paid mutator transaction binding the contract method 0x8ae285d3.
//
// Solidity: function utxoAnchorStateStruct((bytes32,uint32,uint64,uint64) ) returns()
func (_TeePayments *TeePaymentsTransactorSession) UtxoAnchorStateStruct(arg0 ITeePaymentsUtxoUtxoAnchorState) (*types.Transaction, error) {
	return _TeePayments.Contract.UtxoAnchorStateStruct(&_TeePayments.TransactOpts, arg0)
}

// UtxoPaymentInstructionMessageStruct is a paid mutator transaction binding the contract method 0x1995f2c6.
//
// Solidity: function utxoPaymentInstructionMessageStruct((bytes32,(address,uint64)[],bytes32,string,uint32,uint32,string,bytes,uint256,uint256,bytes,bytes32,uint64,uint64,uint64,uint64) ) returns()
func (_TeePayments *TeePaymentsTransactor) UtxoPaymentInstructionMessageStruct(opts *bind.TransactOpts, arg0 ITeePaymentsUtxoUtxoPaymentInstructionMessage) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "utxoPaymentInstructionMessageStruct", arg0)
}

// UtxoPaymentInstructionMessageStruct is a paid mutator transaction binding the contract method 0x1995f2c6.
//
// Solidity: function utxoPaymentInstructionMessageStruct((bytes32,(address,uint64)[],bytes32,string,uint32,uint32,string,bytes,uint256,uint256,bytes,bytes32,uint64,uint64,uint64,uint64) ) returns()
func (_TeePayments *TeePaymentsSession) UtxoPaymentInstructionMessageStruct(arg0 ITeePaymentsUtxoUtxoPaymentInstructionMessage) (*types.Transaction, error) {
	return _TeePayments.Contract.UtxoPaymentInstructionMessageStruct(&_TeePayments.TransactOpts, arg0)
}

// UtxoPaymentInstructionMessageStruct is a paid mutator transaction binding the contract method 0x1995f2c6.
//
// Solidity: function utxoPaymentInstructionMessageStruct((bytes32,(address,uint64)[],bytes32,string,uint32,uint32,string,bytes,uint256,uint256,bytes,bytes32,uint64,uint64,uint64,uint64) ) returns()
func (_TeePayments *TeePaymentsTransactorSession) UtxoPaymentInstructionMessageStruct(arg0 ITeePaymentsUtxoUtxoPaymentInstructionMessage) (*types.Transaction, error) {
	return _TeePayments.Contract.UtxoPaymentInstructionMessageStruct(&_TeePayments.TransactOpts, arg0)
}
