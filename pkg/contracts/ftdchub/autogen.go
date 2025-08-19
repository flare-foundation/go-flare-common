// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ftdchub

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

// FtdcHubMetaData contains all meta data concerning the FtdcHub contract.
var FtdcHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CosignersThresholdInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DefaultNumberOfTeesZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinThresholdInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MultipleResponsesPossible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NumberOfTeesAndTeeIdsInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"defaultNumberOfTees\",\"type\":\"uint8\"}],\"name\":\"DefaultNumberOfTeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minThresholdBIPS\",\"type\":\"uint16\"}],\"name\":\"MinThresholdBIPSSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfTees\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_requestBody\",\"type\":\"bytes\"}],\"name\":\"requestAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// FtdcHubABI is the input ABI used to generate the binding from.
// Deprecated: Use FtdcHubMetaData.ABI instead.
var FtdcHubABI = FtdcHubMetaData.ABI

// FtdcHub is an auto generated Go binding around an Ethereum contract.
type FtdcHub struct {
	FtdcHubCaller     // Read-only binding to the contract
	FtdcHubTransactor // Write-only binding to the contract
	FtdcHubFilterer   // Log filterer for contract events
}

// FtdcHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type FtdcHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtdcHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FtdcHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtdcHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FtdcHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtdcHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FtdcHubSession struct {
	Contract     *FtdcHub          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FtdcHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FtdcHubCallerSession struct {
	Contract *FtdcHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FtdcHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FtdcHubTransactorSession struct {
	Contract     *FtdcHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FtdcHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type FtdcHubRaw struct {
	Contract *FtdcHub // Generic contract binding to access the raw methods on
}

// FtdcHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FtdcHubCallerRaw struct {
	Contract *FtdcHubCaller // Generic read-only contract binding to access the raw methods on
}

// FtdcHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FtdcHubTransactorRaw struct {
	Contract *FtdcHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFtdcHub creates a new instance of FtdcHub, bound to a specific deployed contract.
func NewFtdcHub(address common.Address, backend bind.ContractBackend) (*FtdcHub, error) {
	contract, err := bindFtdcHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FtdcHub{FtdcHubCaller: FtdcHubCaller{contract: contract}, FtdcHubTransactor: FtdcHubTransactor{contract: contract}, FtdcHubFilterer: FtdcHubFilterer{contract: contract}}, nil
}

// NewFtdcHubCaller creates a new read-only instance of FtdcHub, bound to a specific deployed contract.
func NewFtdcHubCaller(address common.Address, caller bind.ContractCaller) (*FtdcHubCaller, error) {
	contract, err := bindFtdcHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FtdcHubCaller{contract: contract}, nil
}

// NewFtdcHubTransactor creates a new write-only instance of FtdcHub, bound to a specific deployed contract.
func NewFtdcHubTransactor(address common.Address, transactor bind.ContractTransactor) (*FtdcHubTransactor, error) {
	contract, err := bindFtdcHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FtdcHubTransactor{contract: contract}, nil
}

// NewFtdcHubFilterer creates a new log filterer instance of FtdcHub, bound to a specific deployed contract.
func NewFtdcHubFilterer(address common.Address, filterer bind.ContractFilterer) (*FtdcHubFilterer, error) {
	contract, err := bindFtdcHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FtdcHubFilterer{contract: contract}, nil
}

// bindFtdcHub binds a generic wrapper to an already deployed contract.
func bindFtdcHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FtdcHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FtdcHub *FtdcHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FtdcHub.Contract.FtdcHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FtdcHub *FtdcHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtdcHub.Contract.FtdcHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FtdcHub *FtdcHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FtdcHub.Contract.FtdcHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FtdcHub *FtdcHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FtdcHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FtdcHub *FtdcHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtdcHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FtdcHub *FtdcHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FtdcHub.Contract.contract.Transact(opts, method, params...)
}

// RequestAttestation is a paid mutator transaction binding the contract method 0x8dec79a0.
//
// Solidity: function requestAttestation(uint16 _thresholdBIPS, uint256 _numberOfTees, address[] _teeIds, address[] _cosigners, uint64 _cosignersThreshold, bytes32 _attestationType, bytes32 _sourceId, bytes _requestBody) payable returns()
func (_FtdcHub *FtdcHubTransactor) RequestAttestation(opts *bind.TransactOpts, _thresholdBIPS uint16, _numberOfTees *big.Int, _teeIds []common.Address, _cosigners []common.Address, _cosignersThreshold uint64, _attestationType [32]byte, _sourceId [32]byte, _requestBody []byte) (*types.Transaction, error) {
	return _FtdcHub.contract.Transact(opts, "requestAttestation", _thresholdBIPS, _numberOfTees, _teeIds, _cosigners, _cosignersThreshold, _attestationType, _sourceId, _requestBody)
}

// RequestAttestation is a paid mutator transaction binding the contract method 0x8dec79a0.
//
// Solidity: function requestAttestation(uint16 _thresholdBIPS, uint256 _numberOfTees, address[] _teeIds, address[] _cosigners, uint64 _cosignersThreshold, bytes32 _attestationType, bytes32 _sourceId, bytes _requestBody) payable returns()
func (_FtdcHub *FtdcHubSession) RequestAttestation(_thresholdBIPS uint16, _numberOfTees *big.Int, _teeIds []common.Address, _cosigners []common.Address, _cosignersThreshold uint64, _attestationType [32]byte, _sourceId [32]byte, _requestBody []byte) (*types.Transaction, error) {
	return _FtdcHub.Contract.RequestAttestation(&_FtdcHub.TransactOpts, _thresholdBIPS, _numberOfTees, _teeIds, _cosigners, _cosignersThreshold, _attestationType, _sourceId, _requestBody)
}

// RequestAttestation is a paid mutator transaction binding the contract method 0x8dec79a0.
//
// Solidity: function requestAttestation(uint16 _thresholdBIPS, uint256 _numberOfTees, address[] _teeIds, address[] _cosigners, uint64 _cosignersThreshold, bytes32 _attestationType, bytes32 _sourceId, bytes _requestBody) payable returns()
func (_FtdcHub *FtdcHubTransactorSession) RequestAttestation(_thresholdBIPS uint16, _numberOfTees *big.Int, _teeIds []common.Address, _cosigners []common.Address, _cosignersThreshold uint64, _attestationType [32]byte, _sourceId [32]byte, _requestBody []byte) (*types.Transaction, error) {
	return _FtdcHub.Contract.RequestAttestation(&_FtdcHub.TransactOpts, _thresholdBIPS, _numberOfTees, _teeIds, _cosigners, _cosignersThreshold, _attestationType, _sourceId, _requestBody)
}

// FtdcHubDefaultNumberOfTeesSetIterator is returned from FilterDefaultNumberOfTeesSet and is used to iterate over the raw logs and unpacked data for DefaultNumberOfTeesSet events raised by the FtdcHub contract.
type FtdcHubDefaultNumberOfTeesSetIterator struct {
	Event *FtdcHubDefaultNumberOfTeesSet // Event containing the contract specifics and raw log

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
func (it *FtdcHubDefaultNumberOfTeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcHubDefaultNumberOfTeesSet)
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
		it.Event = new(FtdcHubDefaultNumberOfTeesSet)
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
func (it *FtdcHubDefaultNumberOfTeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcHubDefaultNumberOfTeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcHubDefaultNumberOfTeesSet represents a DefaultNumberOfTeesSet event raised by the FtdcHub contract.
type FtdcHubDefaultNumberOfTeesSet struct {
	DefaultNumberOfTees uint8
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDefaultNumberOfTeesSet is a free log retrieval operation binding the contract event 0xfd28c45cd39f6ed5f3bcbd90a07ba9adc16d49ba1440a07595ee4d6bd8f8dafd.
//
// Solidity: event DefaultNumberOfTeesSet(uint8 defaultNumberOfTees)
func (_FtdcHub *FtdcHubFilterer) FilterDefaultNumberOfTeesSet(opts *bind.FilterOpts) (*FtdcHubDefaultNumberOfTeesSetIterator, error) {

	logs, sub, err := _FtdcHub.contract.FilterLogs(opts, "DefaultNumberOfTeesSet")
	if err != nil {
		return nil, err
	}
	return &FtdcHubDefaultNumberOfTeesSetIterator{contract: _FtdcHub.contract, event: "DefaultNumberOfTeesSet", logs: logs, sub: sub}, nil
}

// WatchDefaultNumberOfTeesSet is a free log subscription operation binding the contract event 0xfd28c45cd39f6ed5f3bcbd90a07ba9adc16d49ba1440a07595ee4d6bd8f8dafd.
//
// Solidity: event DefaultNumberOfTeesSet(uint8 defaultNumberOfTees)
func (_FtdcHub *FtdcHubFilterer) WatchDefaultNumberOfTeesSet(opts *bind.WatchOpts, sink chan<- *FtdcHubDefaultNumberOfTeesSet) (event.Subscription, error) {

	logs, sub, err := _FtdcHub.contract.WatchLogs(opts, "DefaultNumberOfTeesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcHubDefaultNumberOfTeesSet)
				if err := _FtdcHub.contract.UnpackLog(event, "DefaultNumberOfTeesSet", log); err != nil {
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

// ParseDefaultNumberOfTeesSet is a log parse operation binding the contract event 0xfd28c45cd39f6ed5f3bcbd90a07ba9adc16d49ba1440a07595ee4d6bd8f8dafd.
//
// Solidity: event DefaultNumberOfTeesSet(uint8 defaultNumberOfTees)
func (_FtdcHub *FtdcHubFilterer) ParseDefaultNumberOfTeesSet(log types.Log) (*FtdcHubDefaultNumberOfTeesSet, error) {
	event := new(FtdcHubDefaultNumberOfTeesSet)
	if err := _FtdcHub.contract.UnpackLog(event, "DefaultNumberOfTeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcHubMinThresholdBIPSSetIterator is returned from FilterMinThresholdBIPSSet and is used to iterate over the raw logs and unpacked data for MinThresholdBIPSSet events raised by the FtdcHub contract.
type FtdcHubMinThresholdBIPSSetIterator struct {
	Event *FtdcHubMinThresholdBIPSSet // Event containing the contract specifics and raw log

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
func (it *FtdcHubMinThresholdBIPSSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcHubMinThresholdBIPSSet)
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
		it.Event = new(FtdcHubMinThresholdBIPSSet)
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
func (it *FtdcHubMinThresholdBIPSSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcHubMinThresholdBIPSSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcHubMinThresholdBIPSSet represents a MinThresholdBIPSSet event raised by the FtdcHub contract.
type FtdcHubMinThresholdBIPSSet struct {
	MinThresholdBIPS uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMinThresholdBIPSSet is a free log retrieval operation binding the contract event 0xc9abb127b1b92ec942674d22e590b5a81e559d7a65f23c63965166a3ffacb92d.
//
// Solidity: event MinThresholdBIPSSet(uint16 minThresholdBIPS)
func (_FtdcHub *FtdcHubFilterer) FilterMinThresholdBIPSSet(opts *bind.FilterOpts) (*FtdcHubMinThresholdBIPSSetIterator, error) {

	logs, sub, err := _FtdcHub.contract.FilterLogs(opts, "MinThresholdBIPSSet")
	if err != nil {
		return nil, err
	}
	return &FtdcHubMinThresholdBIPSSetIterator{contract: _FtdcHub.contract, event: "MinThresholdBIPSSet", logs: logs, sub: sub}, nil
}

// WatchMinThresholdBIPSSet is a free log subscription operation binding the contract event 0xc9abb127b1b92ec942674d22e590b5a81e559d7a65f23c63965166a3ffacb92d.
//
// Solidity: event MinThresholdBIPSSet(uint16 minThresholdBIPS)
func (_FtdcHub *FtdcHubFilterer) WatchMinThresholdBIPSSet(opts *bind.WatchOpts, sink chan<- *FtdcHubMinThresholdBIPSSet) (event.Subscription, error) {

	logs, sub, err := _FtdcHub.contract.WatchLogs(opts, "MinThresholdBIPSSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcHubMinThresholdBIPSSet)
				if err := _FtdcHub.contract.UnpackLog(event, "MinThresholdBIPSSet", log); err != nil {
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

// ParseMinThresholdBIPSSet is a log parse operation binding the contract event 0xc9abb127b1b92ec942674d22e590b5a81e559d7a65f23c63965166a3ffacb92d.
//
// Solidity: event MinThresholdBIPSSet(uint16 minThresholdBIPS)
func (_FtdcHub *FtdcHubFilterer) ParseMinThresholdBIPSSet(log types.Log) (*FtdcHubMinThresholdBIPSSet, error) {
	event := new(FtdcHubMinThresholdBIPSSet)
	if err := _FtdcHub.contract.UnpackLog(event, "MinThresholdBIPSSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
