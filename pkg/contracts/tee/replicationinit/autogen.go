// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package replicationinit

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

// ReplicationInitMetaData contains all meta data concerning the ReplicationInit contract.
var ReplicationInitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pauseBeforeUpgradeMinDurationSeconds\",\"type\":\"uint256\"}],\"name\":\"PauseBeforeUpgradeMinDurationSecondsSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pauseBeforeUpgradeMinDurationSeconds\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ReplicationInitABI is the input ABI used to generate the binding from.
// Deprecated: Use ReplicationInitMetaData.ABI instead.
var ReplicationInitABI = ReplicationInitMetaData.ABI

// ReplicationInit is an auto generated Go binding around an Ethereum contract.
type ReplicationInit struct {
	ReplicationInitCaller     // Read-only binding to the contract
	ReplicationInitTransactor // Write-only binding to the contract
	ReplicationInitFilterer   // Log filterer for contract events
}

// ReplicationInitCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReplicationInitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicationInitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReplicationInitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicationInitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReplicationInitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReplicationInitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReplicationInitSession struct {
	Contract     *ReplicationInit  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReplicationInitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReplicationInitCallerSession struct {
	Contract *ReplicationInitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReplicationInitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReplicationInitTransactorSession struct {
	Contract     *ReplicationInitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReplicationInitRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReplicationInitRaw struct {
	Contract *ReplicationInit // Generic contract binding to access the raw methods on
}

// ReplicationInitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReplicationInitCallerRaw struct {
	Contract *ReplicationInitCaller // Generic read-only contract binding to access the raw methods on
}

// ReplicationInitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReplicationInitTransactorRaw struct {
	Contract *ReplicationInitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReplicationInit creates a new instance of ReplicationInit, bound to a specific deployed contract.
func NewReplicationInit(address common.Address, backend bind.ContractBackend) (*ReplicationInit, error) {
	contract, err := bindReplicationInit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReplicationInit{ReplicationInitCaller: ReplicationInitCaller{contract: contract}, ReplicationInitTransactor: ReplicationInitTransactor{contract: contract}, ReplicationInitFilterer: ReplicationInitFilterer{contract: contract}}, nil
}

// NewReplicationInitCaller creates a new read-only instance of ReplicationInit, bound to a specific deployed contract.
func NewReplicationInitCaller(address common.Address, caller bind.ContractCaller) (*ReplicationInitCaller, error) {
	contract, err := bindReplicationInit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReplicationInitCaller{contract: contract}, nil
}

// NewReplicationInitTransactor creates a new write-only instance of ReplicationInit, bound to a specific deployed contract.
func NewReplicationInitTransactor(address common.Address, transactor bind.ContractTransactor) (*ReplicationInitTransactor, error) {
	contract, err := bindReplicationInit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReplicationInitTransactor{contract: contract}, nil
}

// NewReplicationInitFilterer creates a new log filterer instance of ReplicationInit, bound to a specific deployed contract.
func NewReplicationInitFilterer(address common.Address, filterer bind.ContractFilterer) (*ReplicationInitFilterer, error) {
	contract, err := bindReplicationInit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReplicationInitFilterer{contract: contract}, nil
}

// bindReplicationInit binds a generic wrapper to an already deployed contract.
func bindReplicationInit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReplicationInitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReplicationInit *ReplicationInitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReplicationInit.Contract.ReplicationInitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReplicationInit *ReplicationInitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReplicationInit.Contract.ReplicationInitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReplicationInit *ReplicationInitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReplicationInit.Contract.ReplicationInitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReplicationInit *ReplicationInitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReplicationInit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReplicationInit *ReplicationInitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReplicationInit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReplicationInit *ReplicationInitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReplicationInit.Contract.contract.Transact(opts, method, params...)
}

// Init is a paid mutator transaction binding the contract method 0xb7b0422d.
//
// Solidity: function init(uint256 _pauseBeforeUpgradeMinDurationSeconds) returns()
func (_ReplicationInit *ReplicationInitTransactor) Init(opts *bind.TransactOpts, _pauseBeforeUpgradeMinDurationSeconds *big.Int) (*types.Transaction, error) {
	return _ReplicationInit.contract.Transact(opts, "init", _pauseBeforeUpgradeMinDurationSeconds)
}

// Init is a paid mutator transaction binding the contract method 0xb7b0422d.
//
// Solidity: function init(uint256 _pauseBeforeUpgradeMinDurationSeconds) returns()
func (_ReplicationInit *ReplicationInitSession) Init(_pauseBeforeUpgradeMinDurationSeconds *big.Int) (*types.Transaction, error) {
	return _ReplicationInit.Contract.Init(&_ReplicationInit.TransactOpts, _pauseBeforeUpgradeMinDurationSeconds)
}

// Init is a paid mutator transaction binding the contract method 0xb7b0422d.
//
// Solidity: function init(uint256 _pauseBeforeUpgradeMinDurationSeconds) returns()
func (_ReplicationInit *ReplicationInitTransactorSession) Init(_pauseBeforeUpgradeMinDurationSeconds *big.Int) (*types.Transaction, error) {
	return _ReplicationInit.Contract.Init(&_ReplicationInit.TransactOpts, _pauseBeforeUpgradeMinDurationSeconds)
}

// ReplicationInitPauseBeforeUpgradeMinDurationSecondsSetIterator is returned from FilterPauseBeforeUpgradeMinDurationSecondsSet and is used to iterate over the raw logs and unpacked data for PauseBeforeUpgradeMinDurationSecondsSet events raised by the ReplicationInit contract.
type ReplicationInitPauseBeforeUpgradeMinDurationSecondsSetIterator struct {
	Event *ReplicationInitPauseBeforeUpgradeMinDurationSecondsSet // Event containing the contract specifics and raw log

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
func (it *ReplicationInitPauseBeforeUpgradeMinDurationSecondsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReplicationInitPauseBeforeUpgradeMinDurationSecondsSet)
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
		it.Event = new(ReplicationInitPauseBeforeUpgradeMinDurationSecondsSet)
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
func (it *ReplicationInitPauseBeforeUpgradeMinDurationSecondsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReplicationInitPauseBeforeUpgradeMinDurationSecondsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReplicationInitPauseBeforeUpgradeMinDurationSecondsSet represents a PauseBeforeUpgradeMinDurationSecondsSet event raised by the ReplicationInit contract.
type ReplicationInitPauseBeforeUpgradeMinDurationSecondsSet struct {
	PauseBeforeUpgradeMinDurationSeconds *big.Int
	Raw                                  types.Log // Blockchain specific contextual infos
}

// FilterPauseBeforeUpgradeMinDurationSecondsSet is a free log retrieval operation binding the contract event 0x90d47fa68d02b1c17231c83feaa930838b7cad77a1a9b948fe0a835ba52b16e0.
//
// Solidity: event PauseBeforeUpgradeMinDurationSecondsSet(uint256 pauseBeforeUpgradeMinDurationSeconds)
func (_ReplicationInit *ReplicationInitFilterer) FilterPauseBeforeUpgradeMinDurationSecondsSet(opts *bind.FilterOpts) (*ReplicationInitPauseBeforeUpgradeMinDurationSecondsSetIterator, error) {

	logs, sub, err := _ReplicationInit.contract.FilterLogs(opts, "PauseBeforeUpgradeMinDurationSecondsSet")
	if err != nil {
		return nil, err
	}
	return &ReplicationInitPauseBeforeUpgradeMinDurationSecondsSetIterator{contract: _ReplicationInit.contract, event: "PauseBeforeUpgradeMinDurationSecondsSet", logs: logs, sub: sub}, nil
}

// WatchPauseBeforeUpgradeMinDurationSecondsSet is a free log subscription operation binding the contract event 0x90d47fa68d02b1c17231c83feaa930838b7cad77a1a9b948fe0a835ba52b16e0.
//
// Solidity: event PauseBeforeUpgradeMinDurationSecondsSet(uint256 pauseBeforeUpgradeMinDurationSeconds)
func (_ReplicationInit *ReplicationInitFilterer) WatchPauseBeforeUpgradeMinDurationSecondsSet(opts *bind.WatchOpts, sink chan<- *ReplicationInitPauseBeforeUpgradeMinDurationSecondsSet) (event.Subscription, error) {

	logs, sub, err := _ReplicationInit.contract.WatchLogs(opts, "PauseBeforeUpgradeMinDurationSecondsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReplicationInitPauseBeforeUpgradeMinDurationSecondsSet)
				if err := _ReplicationInit.contract.UnpackLog(event, "PauseBeforeUpgradeMinDurationSecondsSet", log); err != nil {
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

// ParsePauseBeforeUpgradeMinDurationSecondsSet is a log parse operation binding the contract event 0x90d47fa68d02b1c17231c83feaa930838b7cad77a1a9b948fe0a835ba52b16e0.
//
// Solidity: event PauseBeforeUpgradeMinDurationSecondsSet(uint256 pauseBeforeUpgradeMinDurationSeconds)
func (_ReplicationInit *ReplicationInitFilterer) ParsePauseBeforeUpgradeMinDurationSecondsSet(log types.Log) (*ReplicationInitPauseBeforeUpgradeMinDurationSecondsSet, error) {
	event := new(ReplicationInitPauseBeforeUpgradeMinDurationSecondsSet)
	if err := _ReplicationInit.contract.UnpackLog(event, "PauseBeforeUpgradeMinDurationSecondsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
