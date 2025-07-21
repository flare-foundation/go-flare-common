// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tee

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

// PublicKey is an auto generated low-level Go binding around an user-defined struct.
type PublicKey struct {
	X [32]byte
	Y [32]byte
}

// TeeStructsAttestation is an auto generated low-level Go binding around an user-defined struct.
type TeeStructsAttestation struct {
	Challenge                [32]byte
	PublicKey                PublicKey
	InitialSigningPolicyId   uint32
	InitialSigningPolicyHash [32]byte
	LastSigningPolicyId      uint32
	LastSigningPolicyHash    [32]byte
	StateHash                [32]byte
	TeeTimestamp             uint64
}

// TeeStructsInstruction is an auto generated low-level Go binding around an user-defined struct.
type TeeStructsInstruction struct {
	InstructionId          [32]byte
	TeeId                  common.Address
	Timestamp              uint64
	RewardEpochId          uint32
	OpType                 [32]byte
	OpCommand              [32]byte
	OriginalMessage        []byte
	AdditionalFixedMessage []byte
}

// TeeStructsPMWState is an auto generated low-level Go binding around an user-defined struct.
type TeeStructsPMWState struct {
	Status *big.Int
}

// TeeStructsVoteReceipt is an auto generated low-level Go binding around an user-defined struct.
type TeeStructsVoteReceipt struct {
	InstructionHash               [32]byte
	Sequence                      uint64
	Signature                     []byte
	AdditionalVariableMessageHash [32]byte
	Timestamp                     uint64
	VoteHash                      [32]byte
}

// TeeStructsVoteSequenceInit is an auto generated low-level Go binding around an user-defined struct.
type TeeStructsVoteSequenceInit struct {
	InstructionId   [32]byte
	InstructionHash [32]byte
	RewardEpochId   uint32
	TeeId           common.Address
}

// TeeStructsVoteSequenceNext is an auto generated low-level Go binding around an user-defined struct.
type TeeStructsVoteSequenceNext struct {
	VoteHash                      [32]byte
	Sequence                      uint64
	Signature                     []byte
	AdditionalVariableMessageHash [32]byte
	Timestamp                     uint64
}

// TeeMetaData contains all meta data concerning the Tee contract.
var TeeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"initialSigningPolicyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"lastSigningPolicyHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structTeeStructs.Attestation\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"attestationStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"originalMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"additionalFixedMessage\",\"type\":\"bytes\"}],\"internalType\":\"structTeeStructs.Instruction\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"instructionStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"internalType\":\"structTeeStructs.PMWState\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwStateStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"instructionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"additionalVariableMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"voteHash\",\"type\":\"bytes32\"}],\"internalType\":\"structTeeStructs.VoteReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"voteReceiptStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"internalType\":\"structTeeStructs.VoteSequenceInit\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"voteSequenceInitStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"voteHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"additionalVariableMessageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structTeeStructs.VoteSequenceNext\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"voteSequenceNextStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeMetaData.ABI instead.
var TeeABI = TeeMetaData.ABI

// Tee is an auto generated Go binding around an Ethereum contract.
type Tee struct {
	TeeCaller     // Read-only binding to the contract
	TeeTransactor // Write-only binding to the contract
	TeeFilterer   // Log filterer for contract events
}

// TeeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeSession struct {
	Contract     *Tee              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeCallerSession struct {
	Contract *TeeCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TeeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeTransactorSession struct {
	Contract     *TeeTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeRaw struct {
	Contract *Tee // Generic contract binding to access the raw methods on
}

// TeeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeCallerRaw struct {
	Contract *TeeCaller // Generic read-only contract binding to access the raw methods on
}

// TeeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeTransactorRaw struct {
	Contract *TeeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTee creates a new instance of Tee, bound to a specific deployed contract.
func NewTee(address common.Address, backend bind.ContractBackend) (*Tee, error) {
	contract, err := bindTee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tee{TeeCaller: TeeCaller{contract: contract}, TeeTransactor: TeeTransactor{contract: contract}, TeeFilterer: TeeFilterer{contract: contract}}, nil
}

// NewTeeCaller creates a new read-only instance of Tee, bound to a specific deployed contract.
func NewTeeCaller(address common.Address, caller bind.ContractCaller) (*TeeCaller, error) {
	contract, err := bindTee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeCaller{contract: contract}, nil
}

// NewTeeTransactor creates a new write-only instance of Tee, bound to a specific deployed contract.
func NewTeeTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeTransactor, error) {
	contract, err := bindTee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeTransactor{contract: contract}, nil
}

// NewTeeFilterer creates a new log filterer instance of Tee, bound to a specific deployed contract.
func NewTeeFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeFilterer, error) {
	contract, err := bindTee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeFilterer{contract: contract}, nil
}

// bindTee binds a generic wrapper to an already deployed contract.
func bindTee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tee *TeeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tee.Contract.TeeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tee *TeeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tee.Contract.TeeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tee *TeeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tee.Contract.TeeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tee *TeeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tee *TeeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tee *TeeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tee.Contract.contract.Transact(opts, method, params...)
}

// AttestationStruct is a paid mutator transaction binding the contract method 0xb567ed85.
//
// Solidity: function attestationStruct((bytes32,(bytes32,bytes32),uint32,bytes32,uint32,bytes32,bytes32,uint64) ) returns()
func (_Tee *TeeTransactor) AttestationStruct(opts *bind.TransactOpts, arg0 TeeStructsAttestation) (*types.Transaction, error) {
	return _Tee.contract.Transact(opts, "attestationStruct", arg0)
}

// AttestationStruct is a paid mutator transaction binding the contract method 0xb567ed85.
//
// Solidity: function attestationStruct((bytes32,(bytes32,bytes32),uint32,bytes32,uint32,bytes32,bytes32,uint64) ) returns()
func (_Tee *TeeSession) AttestationStruct(arg0 TeeStructsAttestation) (*types.Transaction, error) {
	return _Tee.Contract.AttestationStruct(&_Tee.TransactOpts, arg0)
}

// AttestationStruct is a paid mutator transaction binding the contract method 0xb567ed85.
//
// Solidity: function attestationStruct((bytes32,(bytes32,bytes32),uint32,bytes32,uint32,bytes32,bytes32,uint64) ) returns()
func (_Tee *TeeTransactorSession) AttestationStruct(arg0 TeeStructsAttestation) (*types.Transaction, error) {
	return _Tee.Contract.AttestationStruct(&_Tee.TransactOpts, arg0)
}

// InstructionStruct is a paid mutator transaction binding the contract method 0x02c349a1.
//
// Solidity: function instructionStruct((bytes32,address,uint64,uint32,bytes32,bytes32,bytes,bytes) ) returns()
func (_Tee *TeeTransactor) InstructionStruct(opts *bind.TransactOpts, arg0 TeeStructsInstruction) (*types.Transaction, error) {
	return _Tee.contract.Transact(opts, "instructionStruct", arg0)
}

// InstructionStruct is a paid mutator transaction binding the contract method 0x02c349a1.
//
// Solidity: function instructionStruct((bytes32,address,uint64,uint32,bytes32,bytes32,bytes,bytes) ) returns()
func (_Tee *TeeSession) InstructionStruct(arg0 TeeStructsInstruction) (*types.Transaction, error) {
	return _Tee.Contract.InstructionStruct(&_Tee.TransactOpts, arg0)
}

// InstructionStruct is a paid mutator transaction binding the contract method 0x02c349a1.
//
// Solidity: function instructionStruct((bytes32,address,uint64,uint32,bytes32,bytes32,bytes,bytes) ) returns()
func (_Tee *TeeTransactorSession) InstructionStruct(arg0 TeeStructsInstruction) (*types.Transaction, error) {
	return _Tee.Contract.InstructionStruct(&_Tee.TransactOpts, arg0)
}

// PmwStateStruct is a paid mutator transaction binding the contract method 0x2bbf19a4.
//
// Solidity: function pmwStateStruct((uint256) ) returns()
func (_Tee *TeeTransactor) PmwStateStruct(opts *bind.TransactOpts, arg0 TeeStructsPMWState) (*types.Transaction, error) {
	return _Tee.contract.Transact(opts, "pmwStateStruct", arg0)
}

// PmwStateStruct is a paid mutator transaction binding the contract method 0x2bbf19a4.
//
// Solidity: function pmwStateStruct((uint256) ) returns()
func (_Tee *TeeSession) PmwStateStruct(arg0 TeeStructsPMWState) (*types.Transaction, error) {
	return _Tee.Contract.PmwStateStruct(&_Tee.TransactOpts, arg0)
}

// PmwStateStruct is a paid mutator transaction binding the contract method 0x2bbf19a4.
//
// Solidity: function pmwStateStruct((uint256) ) returns()
func (_Tee *TeeTransactorSession) PmwStateStruct(arg0 TeeStructsPMWState) (*types.Transaction, error) {
	return _Tee.Contract.PmwStateStruct(&_Tee.TransactOpts, arg0)
}

// VoteReceiptStruct is a paid mutator transaction binding the contract method 0x2ac04628.
//
// Solidity: function voteReceiptStruct((bytes32,uint64,bytes,bytes32,uint64,bytes32) ) returns()
func (_Tee *TeeTransactor) VoteReceiptStruct(opts *bind.TransactOpts, arg0 TeeStructsVoteReceipt) (*types.Transaction, error) {
	return _Tee.contract.Transact(opts, "voteReceiptStruct", arg0)
}

// VoteReceiptStruct is a paid mutator transaction binding the contract method 0x2ac04628.
//
// Solidity: function voteReceiptStruct((bytes32,uint64,bytes,bytes32,uint64,bytes32) ) returns()
func (_Tee *TeeSession) VoteReceiptStruct(arg0 TeeStructsVoteReceipt) (*types.Transaction, error) {
	return _Tee.Contract.VoteReceiptStruct(&_Tee.TransactOpts, arg0)
}

// VoteReceiptStruct is a paid mutator transaction binding the contract method 0x2ac04628.
//
// Solidity: function voteReceiptStruct((bytes32,uint64,bytes,bytes32,uint64,bytes32) ) returns()
func (_Tee *TeeTransactorSession) VoteReceiptStruct(arg0 TeeStructsVoteReceipt) (*types.Transaction, error) {
	return _Tee.Contract.VoteReceiptStruct(&_Tee.TransactOpts, arg0)
}

// VoteSequenceInitStruct is a paid mutator transaction binding the contract method 0x70874ce3.
//
// Solidity: function voteSequenceInitStruct((bytes32,bytes32,uint32,address) ) returns()
func (_Tee *TeeTransactor) VoteSequenceInitStruct(opts *bind.TransactOpts, arg0 TeeStructsVoteSequenceInit) (*types.Transaction, error) {
	return _Tee.contract.Transact(opts, "voteSequenceInitStruct", arg0)
}

// VoteSequenceInitStruct is a paid mutator transaction binding the contract method 0x70874ce3.
//
// Solidity: function voteSequenceInitStruct((bytes32,bytes32,uint32,address) ) returns()
func (_Tee *TeeSession) VoteSequenceInitStruct(arg0 TeeStructsVoteSequenceInit) (*types.Transaction, error) {
	return _Tee.Contract.VoteSequenceInitStruct(&_Tee.TransactOpts, arg0)
}

// VoteSequenceInitStruct is a paid mutator transaction binding the contract method 0x70874ce3.
//
// Solidity: function voteSequenceInitStruct((bytes32,bytes32,uint32,address) ) returns()
func (_Tee *TeeTransactorSession) VoteSequenceInitStruct(arg0 TeeStructsVoteSequenceInit) (*types.Transaction, error) {
	return _Tee.Contract.VoteSequenceInitStruct(&_Tee.TransactOpts, arg0)
}

// VoteSequenceNextStruct is a paid mutator transaction binding the contract method 0x7797af66.
//
// Solidity: function voteSequenceNextStruct((bytes32,uint64,bytes,bytes32,uint64) ) returns()
func (_Tee *TeeTransactor) VoteSequenceNextStruct(opts *bind.TransactOpts, arg0 TeeStructsVoteSequenceNext) (*types.Transaction, error) {
	return _Tee.contract.Transact(opts, "voteSequenceNextStruct", arg0)
}

// VoteSequenceNextStruct is a paid mutator transaction binding the contract method 0x7797af66.
//
// Solidity: function voteSequenceNextStruct((bytes32,uint64,bytes,bytes32,uint64) ) returns()
func (_Tee *TeeSession) VoteSequenceNextStruct(arg0 TeeStructsVoteSequenceNext) (*types.Transaction, error) {
	return _Tee.Contract.VoteSequenceNextStruct(&_Tee.TransactOpts, arg0)
}

// VoteSequenceNextStruct is a paid mutator transaction binding the contract method 0x7797af66.
//
// Solidity: function voteSequenceNextStruct((bytes32,uint64,bytes,bytes32,uint64) ) returns()
func (_Tee *TeeTransactorSession) VoteSequenceNextStruct(arg0 TeeStructsVoteSequenceNext) (*types.Transaction, error) {
	return _Tee.Contract.VoteSequenceNextStruct(&_Tee.TransactOpts, arg0)
}
