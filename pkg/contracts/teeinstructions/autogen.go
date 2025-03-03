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

// ITeeRegistryTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryTeeMachine struct {
	TeeId common.Address
	Owner common.Address
	Url   string
}

// TeeInstructionsMetaData contains all meta data concerning the TeeInstructions contract.
var TeeInstructionsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structITeeRegistry.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getInstructionInitiators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structITeeRegistry.TeeMachine[]\",\"name\":\"_teeMachines\",\"type\":\"tuple[]\"},{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendInstructions\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// GetInstructionInitiators is a free data retrieval call binding the contract method 0xf1233d2d.
//
// Solidity: function getInstructionInitiators() view returns(address[])
func (_TeeInstructions *TeeInstructionsCaller) GetInstructionInitiators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TeeInstructions.contract.Call(opts, &out, "getInstructionInitiators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetInstructionInitiators is a free data retrieval call binding the contract method 0xf1233d2d.
//
// Solidity: function getInstructionInitiators() view returns(address[])
func (_TeeInstructions *TeeInstructionsSession) GetInstructionInitiators() ([]common.Address, error) {
	return _TeeInstructions.Contract.GetInstructionInitiators(&_TeeInstructions.CallOpts)
}

// GetInstructionInitiators is a free data retrieval call binding the contract method 0xf1233d2d.
//
// Solidity: function getInstructionInitiators() view returns(address[])
func (_TeeInstructions *TeeInstructionsCallerSession) GetInstructionInitiators() ([]common.Address, error) {
	return _TeeInstructions.Contract.GetInstructionInitiators(&_TeeInstructions.CallOpts)
}

// SendInstructions is a paid mutator transaction binding the contract method 0xe9320ea1.
//
// Solidity: function sendInstructions(bytes32 _instructionId, (address,address,string)[] _teeMachines, uint24 _rewardEpochId, bytes32 _opType, bytes32 _opCommand, bytes _message) payable returns()
func (_TeeInstructions *TeeInstructionsTransactor) SendInstructions(opts *bind.TransactOpts, _instructionId [32]byte, _teeMachines []ITeeRegistryTeeMachine, _rewardEpochId *big.Int, _opType [32]byte, _opCommand [32]byte, _message []byte) (*types.Transaction, error) {
	return _TeeInstructions.contract.Transact(opts, "sendInstructions", _instructionId, _teeMachines, _rewardEpochId, _opType, _opCommand, _message)
}

// SendInstructions is a paid mutator transaction binding the contract method 0xe9320ea1.
//
// Solidity: function sendInstructions(bytes32 _instructionId, (address,address,string)[] _teeMachines, uint24 _rewardEpochId, bytes32 _opType, bytes32 _opCommand, bytes _message) payable returns()
func (_TeeInstructions *TeeInstructionsSession) SendInstructions(_instructionId [32]byte, _teeMachines []ITeeRegistryTeeMachine, _rewardEpochId *big.Int, _opType [32]byte, _opCommand [32]byte, _message []byte) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SendInstructions(&_TeeInstructions.TransactOpts, _instructionId, _teeMachines, _rewardEpochId, _opType, _opCommand, _message)
}

// SendInstructions is a paid mutator transaction binding the contract method 0xe9320ea1.
//
// Solidity: function sendInstructions(bytes32 _instructionId, (address,address,string)[] _teeMachines, uint24 _rewardEpochId, bytes32 _opType, bytes32 _opCommand, bytes _message) payable returns()
func (_TeeInstructions *TeeInstructionsTransactorSession) SendInstructions(_instructionId [32]byte, _teeMachines []ITeeRegistryTeeMachine, _rewardEpochId *big.Int, _opType [32]byte, _opCommand [32]byte, _message []byte) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SendInstructions(&_TeeInstructions.TransactOpts, _instructionId, _teeMachines, _rewardEpochId, _opType, _opCommand, _message)
}

// TeeInstructionsTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the TeeInstructions contract.
type TeeInstructionsTeeInstructionsSentIterator struct {
	Event *TeeInstructionsTeeInstructionsSent // Event containing the contract specifics and raw log

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
func (it *TeeInstructionsTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeInstructionsTeeInstructionsSent)
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
		it.Event = new(TeeInstructionsTeeInstructionsSent)
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
func (it *TeeInstructionsTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeInstructionsTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeInstructionsTeeInstructionsSent represents a TeeInstructionsSent event raised by the TeeInstructions contract.
type TeeInstructionsTeeInstructionsSent struct {
	InstructionId [32]byte
	RewardEpochId *big.Int
	TeeMachines   []ITeeRegistryTeeMachine
	OpType        [32]byte
	OpCommand     [32]byte
	Message       []byte
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTeeInstructionsSent is a free log retrieval operation binding the contract event 0xde9f7143f1f7c75c63309d31548e8decbeaefc90b4f7f72b15c4a83d45f6ea77.
//
// Solidity: event TeeInstructionsSent(bytes32 indexed instructionId, uint24 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, uint256 fee)
func (_TeeInstructions *TeeInstructionsFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, instructionId [][32]byte, rewardEpochId []*big.Int) (*TeeInstructionsTeeInstructionsSentIterator, error) {

	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeInstructions.contract.FilterLogs(opts, "TeeInstructionsSent", instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsTeeInstructionsSentIterator{contract: _TeeInstructions.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0xde9f7143f1f7c75c63309d31548e8decbeaefc90b4f7f72b15c4a83d45f6ea77.
//
// Solidity: event TeeInstructionsSent(bytes32 indexed instructionId, uint24 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, uint256 fee)
func (_TeeInstructions *TeeInstructionsFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *TeeInstructionsTeeInstructionsSent, instructionId [][32]byte, rewardEpochId []*big.Int) (event.Subscription, error) {

	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeInstructions.contract.WatchLogs(opts, "TeeInstructionsSent", instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeInstructionsTeeInstructionsSent)
				if err := _TeeInstructions.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
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

// ParseTeeInstructionsSent is a log parse operation binding the contract event 0xde9f7143f1f7c75c63309d31548e8decbeaefc90b4f7f72b15c4a83d45f6ea77.
//
// Solidity: event TeeInstructionsSent(bytes32 indexed instructionId, uint24 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, uint256 fee)
func (_TeeInstructions *TeeInstructionsFilterer) ParseTeeInstructionsSent(log types.Log) (*TeeInstructionsTeeInstructionsSent, error) {
	event := new(TeeInstructionsTeeInstructionsSent)
	if err := _TeeInstructions.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
