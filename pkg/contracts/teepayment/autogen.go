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

// ITeeRegistryTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
}

// TeePaymentMetaData contains all meta data concerning the TeePayment contract.
var TeePaymentMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structITeeRegistry.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getInstructionInitiators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structITeeRegistry.TeeMachine[]\",\"name\":\"_teeMachines\",\"type\":\"tuple[]\"},{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendInstructions\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// GetInstructionInitiators is a free data retrieval call binding the contract method 0xf1233d2d.
//
// Solidity: function getInstructionInitiators() view returns(address[])
func (_TeePayment *TeePaymentCaller) GetInstructionInitiators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "getInstructionInitiators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetInstructionInitiators is a free data retrieval call binding the contract method 0xf1233d2d.
//
// Solidity: function getInstructionInitiators() view returns(address[])
func (_TeePayment *TeePaymentSession) GetInstructionInitiators() ([]common.Address, error) {
	return _TeePayment.Contract.GetInstructionInitiators(&_TeePayment.CallOpts)
}

// GetInstructionInitiators is a free data retrieval call binding the contract method 0xf1233d2d.
//
// Solidity: function getInstructionInitiators() view returns(address[])
func (_TeePayment *TeePaymentCallerSession) GetInstructionInitiators() ([]common.Address, error) {
	return _TeePayment.Contract.GetInstructionInitiators(&_TeePayment.CallOpts)
}

// SendInstructions is a paid mutator transaction binding the contract method 0xe9320ea1.
//
// Solidity: function sendInstructions(bytes32 _instructionId, (address,address,string)[] _teeMachines, uint24 _rewardEpochId, bytes32 _opType, bytes32 _opCommand, bytes _message) payable returns()
func (_TeePayment *TeePaymentTransactor) SendInstructions(opts *bind.TransactOpts, _instructionId [32]byte, _teeMachines []ITeeRegistryTeeMachine, _rewardEpochId *big.Int, _opType [32]byte, _opCommand [32]byte, _message []byte) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "sendInstructions", _instructionId, _teeMachines, _rewardEpochId, _opType, _opCommand, _message)
}

// SendInstructions is a paid mutator transaction binding the contract method 0xe9320ea1.
//
// Solidity: function sendInstructions(bytes32 _instructionId, (address,address,string)[] _teeMachines, uint24 _rewardEpochId, bytes32 _opType, bytes32 _opCommand, bytes _message) payable returns()
func (_TeePayment *TeePaymentSession) SendInstructions(_instructionId [32]byte, _teeMachines []ITeeRegistryTeeMachine, _rewardEpochId *big.Int, _opType [32]byte, _opCommand [32]byte, _message []byte) (*types.Transaction, error) {
	return _TeePayment.Contract.SendInstructions(&_TeePayment.TransactOpts, _instructionId, _teeMachines, _rewardEpochId, _opType, _opCommand, _message)
}

// SendInstructions is a paid mutator transaction binding the contract method 0xe9320ea1.
//
// Solidity: function sendInstructions(bytes32 _instructionId, (address,address,string)[] _teeMachines, uint24 _rewardEpochId, bytes32 _opType, bytes32 _opCommand, bytes _message) payable returns()
func (_TeePayment *TeePaymentTransactorSession) SendInstructions(_instructionId [32]byte, _teeMachines []ITeeRegistryTeeMachine, _rewardEpochId *big.Int, _opType [32]byte, _opCommand [32]byte, _message []byte) (*types.Transaction, error) {
	return _TeePayment.Contract.SendInstructions(&_TeePayment.TransactOpts, _instructionId, _teeMachines, _rewardEpochId, _opType, _opCommand, _message)
}

// TeePaymentTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the TeePayment contract.
type TeePaymentTeeInstructionsSentIterator struct {
	Event *TeePaymentTeeInstructionsSent // Event containing the contract specifics and raw log

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
func (it *TeePaymentTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentTeeInstructionsSent)
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
		it.Event = new(TeePaymentTeeInstructionsSent)
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
func (it *TeePaymentTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentTeeInstructionsSent represents a TeeInstructionsSent event raised by the TeePayment contract.
type TeePaymentTeeInstructionsSent struct {
	InstructionId [32]byte
	RewardEpochId uint32
	TeeMachines   []ITeeRegistryTeeMachine
	OpType        [32]byte
	OpCommand     [32]byte
	Message       []byte
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTeeInstructionsSent is a free log retrieval operation binding the contract event 0x0b155d08a134ef49412293aa17be27ead7cc04b35ada16085a36c4355fd6f40d.
//
// Solidity: event TeeInstructionsSent(bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, uint256 fee)
func (_TeePayment *TeePaymentFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, instructionId [][32]byte, rewardEpochId []uint32) (*TeePaymentTeeInstructionsSentIterator, error) {

	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "TeeInstructionsSent", instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentTeeInstructionsSentIterator{contract: _TeePayment.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0x0b155d08a134ef49412293aa17be27ead7cc04b35ada16085a36c4355fd6f40d.
//
// Solidity: event TeeInstructionsSent(bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, uint256 fee)
func (_TeePayment *TeePaymentFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *TeePaymentTeeInstructionsSent, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "TeeInstructionsSent", instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentTeeInstructionsSent)
				if err := _TeePayment.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
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

// ParseTeeInstructionsSent is a log parse operation binding the contract event 0x0b155d08a134ef49412293aa17be27ead7cc04b35ada16085a36c4355fd6f40d.
//
// Solidity: event TeeInstructionsSent(bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, uint256 fee)
func (_TeePayment *TeePaymentFilterer) ParseTeeInstructionsSent(log types.Log) (*TeePaymentTeeInstructionsSent, error) {
	event := new(TeePaymentTeeInstructionsSent)
	if err := _TeePayment.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
