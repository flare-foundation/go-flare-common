// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teefeecalculator

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

// TeeFeeCalculatorMetaData contains all meta data concerning the TeeFeeCalculator contract.
var TeeFeeCalculatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultFee\",\"type\":\"uint256\"}],\"name\":\"DefaultFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"opTypes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"opCommands\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"}],\"name\":\"OperationFeesSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"name\":\"calculateFeeByTeeIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"}],\"name\":\"getOperationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_defaultFee\",\"type\":\"uint256\"}],\"name\":\"setDefaultFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_opTypes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_opCommands\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"}],\"name\":\"setOperationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeFeeCalculatorABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeFeeCalculatorMetaData.ABI instead.
var TeeFeeCalculatorABI = TeeFeeCalculatorMetaData.ABI

// TeeFeeCalculator is an auto generated Go binding around an Ethereum contract.
type TeeFeeCalculator struct {
	TeeFeeCalculatorCaller     // Read-only binding to the contract
	TeeFeeCalculatorTransactor // Write-only binding to the contract
	TeeFeeCalculatorFilterer   // Log filterer for contract events
}

// TeeFeeCalculatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeFeeCalculatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeFeeCalculatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeFeeCalculatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeFeeCalculatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeFeeCalculatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeFeeCalculatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeFeeCalculatorSession struct {
	Contract     *TeeFeeCalculator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeFeeCalculatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeFeeCalculatorCallerSession struct {
	Contract *TeeFeeCalculatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TeeFeeCalculatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeFeeCalculatorTransactorSession struct {
	Contract     *TeeFeeCalculatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TeeFeeCalculatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeFeeCalculatorRaw struct {
	Contract *TeeFeeCalculator // Generic contract binding to access the raw methods on
}

// TeeFeeCalculatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeFeeCalculatorCallerRaw struct {
	Contract *TeeFeeCalculatorCaller // Generic read-only contract binding to access the raw methods on
}

// TeeFeeCalculatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeFeeCalculatorTransactorRaw struct {
	Contract *TeeFeeCalculatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeFeeCalculator creates a new instance of TeeFeeCalculator, bound to a specific deployed contract.
func NewTeeFeeCalculator(address common.Address, backend bind.ContractBackend) (*TeeFeeCalculator, error) {
	contract, err := bindTeeFeeCalculator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculator{TeeFeeCalculatorCaller: TeeFeeCalculatorCaller{contract: contract}, TeeFeeCalculatorTransactor: TeeFeeCalculatorTransactor{contract: contract}, TeeFeeCalculatorFilterer: TeeFeeCalculatorFilterer{contract: contract}}, nil
}

// NewTeeFeeCalculatorCaller creates a new read-only instance of TeeFeeCalculator, bound to a specific deployed contract.
func NewTeeFeeCalculatorCaller(address common.Address, caller bind.ContractCaller) (*TeeFeeCalculatorCaller, error) {
	contract, err := bindTeeFeeCalculator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorCaller{contract: contract}, nil
}

// NewTeeFeeCalculatorTransactor creates a new write-only instance of TeeFeeCalculator, bound to a specific deployed contract.
func NewTeeFeeCalculatorTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeFeeCalculatorTransactor, error) {
	contract, err := bindTeeFeeCalculator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorTransactor{contract: contract}, nil
}

// NewTeeFeeCalculatorFilterer creates a new log filterer instance of TeeFeeCalculator, bound to a specific deployed contract.
func NewTeeFeeCalculatorFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeFeeCalculatorFilterer, error) {
	contract, err := bindTeeFeeCalculator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorFilterer{contract: contract}, nil
}

// bindTeeFeeCalculator binds a generic wrapper to an already deployed contract.
func bindTeeFeeCalculator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeFeeCalculatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeFeeCalculator *TeeFeeCalculatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeFeeCalculator.Contract.TeeFeeCalculatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeFeeCalculator *TeeFeeCalculatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.TeeFeeCalculatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeFeeCalculator *TeeFeeCalculatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.TeeFeeCalculatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeFeeCalculator *TeeFeeCalculatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeFeeCalculator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.contract.Transact(opts, method, params...)
}

// CalculateFeeByTeeIds is a free data retrieval call binding the contract method 0x81174ec7.
//
// Solidity: function calculateFeeByTeeIds(bytes32 _opType, bytes32 _opCommand, address[] _teeIds) view returns(uint256 _fee)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) CalculateFeeByTeeIds(opts *bind.CallOpts, _opType [32]byte, _opCommand [32]byte, _teeIds []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "calculateFeeByTeeIds", _opType, _opCommand, _teeIds)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFeeByTeeIds is a free data retrieval call binding the contract method 0x81174ec7.
//
// Solidity: function calculateFeeByTeeIds(bytes32 _opType, bytes32 _opCommand, address[] _teeIds) view returns(uint256 _fee)
func (_TeeFeeCalculator *TeeFeeCalculatorSession) CalculateFeeByTeeIds(_opType [32]byte, _opCommand [32]byte, _teeIds []common.Address) (*big.Int, error) {
	return _TeeFeeCalculator.Contract.CalculateFeeByTeeIds(&_TeeFeeCalculator.CallOpts, _opType, _opCommand, _teeIds)
}

// CalculateFeeByTeeIds is a free data retrieval call binding the contract method 0x81174ec7.
//
// Solidity: function calculateFeeByTeeIds(bytes32 _opType, bytes32 _opCommand, address[] _teeIds) view returns(uint256 _fee)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) CalculateFeeByTeeIds(_opType [32]byte, _opCommand [32]byte, _teeIds []common.Address) (*big.Int, error) {
	return _TeeFeeCalculator.Contract.CalculateFeeByTeeIds(&_TeeFeeCalculator.CallOpts, _opType, _opCommand, _teeIds)
}

// GetDefaultFee is a free data retrieval call binding the contract method 0x21bacf28.
//
// Solidity: function getDefaultFee() view returns(uint256)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) GetDefaultFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "getDefaultFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDefaultFee is a free data retrieval call binding the contract method 0x21bacf28.
//
// Solidity: function getDefaultFee() view returns(uint256)
func (_TeeFeeCalculator *TeeFeeCalculatorSession) GetDefaultFee() (*big.Int, error) {
	return _TeeFeeCalculator.Contract.GetDefaultFee(&_TeeFeeCalculator.CallOpts)
}

// GetDefaultFee is a free data retrieval call binding the contract method 0x21bacf28.
//
// Solidity: function getDefaultFee() view returns(uint256)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) GetDefaultFee() (*big.Int, error) {
	return _TeeFeeCalculator.Contract.GetDefaultFee(&_TeeFeeCalculator.CallOpts)
}

// GetOperationFee is a free data retrieval call binding the contract method 0x3d5e0390.
//
// Solidity: function getOperationFee(bytes32 _opType, bytes32 _opCommand) view returns(uint256)
func (_TeeFeeCalculator *TeeFeeCalculatorCaller) GetOperationFee(opts *bind.CallOpts, _opType [32]byte, _opCommand [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeeFeeCalculator.contract.Call(opts, &out, "getOperationFee", _opType, _opCommand)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperationFee is a free data retrieval call binding the contract method 0x3d5e0390.
//
// Solidity: function getOperationFee(bytes32 _opType, bytes32 _opCommand) view returns(uint256)
func (_TeeFeeCalculator *TeeFeeCalculatorSession) GetOperationFee(_opType [32]byte, _opCommand [32]byte) (*big.Int, error) {
	return _TeeFeeCalculator.Contract.GetOperationFee(&_TeeFeeCalculator.CallOpts, _opType, _opCommand)
}

// GetOperationFee is a free data retrieval call binding the contract method 0x3d5e0390.
//
// Solidity: function getOperationFee(bytes32 _opType, bytes32 _opCommand) view returns(uint256)
func (_TeeFeeCalculator *TeeFeeCalculatorCallerSession) GetOperationFee(_opType [32]byte, _opCommand [32]byte) (*big.Int, error) {
	return _TeeFeeCalculator.Contract.GetOperationFee(&_TeeFeeCalculator.CallOpts, _opType, _opCommand)
}

// SetDefaultFee is a paid mutator transaction binding the contract method 0xc93a6c84.
//
// Solidity: function setDefaultFee(uint256 _defaultFee) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactor) SetDefaultFee(opts *bind.TransactOpts, _defaultFee *big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.contract.Transact(opts, "setDefaultFee", _defaultFee)
}

// SetDefaultFee is a paid mutator transaction binding the contract method 0xc93a6c84.
//
// Solidity: function setDefaultFee(uint256 _defaultFee) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorSession) SetDefaultFee(_defaultFee *big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.SetDefaultFee(&_TeeFeeCalculator.TransactOpts, _defaultFee)
}

// SetDefaultFee is a paid mutator transaction binding the contract method 0xc93a6c84.
//
// Solidity: function setDefaultFee(uint256 _defaultFee) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorSession) SetDefaultFee(_defaultFee *big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.SetDefaultFee(&_TeeFeeCalculator.TransactOpts, _defaultFee)
}

// SetOperationFees is a paid mutator transaction binding the contract method 0xaa259c3e.
//
// Solidity: function setOperationFees(bytes32[] _opTypes, bytes32[] _opCommands, uint256[] _fees) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactor) SetOperationFees(opts *bind.TransactOpts, _opTypes [][32]byte, _opCommands [][32]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.contract.Transact(opts, "setOperationFees", _opTypes, _opCommands, _fees)
}

// SetOperationFees is a paid mutator transaction binding the contract method 0xaa259c3e.
//
// Solidity: function setOperationFees(bytes32[] _opTypes, bytes32[] _opCommands, uint256[] _fees) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorSession) SetOperationFees(_opTypes [][32]byte, _opCommands [][32]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.SetOperationFees(&_TeeFeeCalculator.TransactOpts, _opTypes, _opCommands, _fees)
}

// SetOperationFees is a paid mutator transaction binding the contract method 0xaa259c3e.
//
// Solidity: function setOperationFees(bytes32[] _opTypes, bytes32[] _opCommands, uint256[] _fees) returns()
func (_TeeFeeCalculator *TeeFeeCalculatorTransactorSession) SetOperationFees(_opTypes [][32]byte, _opCommands [][32]byte, _fees []*big.Int) (*types.Transaction, error) {
	return _TeeFeeCalculator.Contract.SetOperationFees(&_TeeFeeCalculator.TransactOpts, _opTypes, _opCommands, _fees)
}

// TeeFeeCalculatorDefaultFeeSetIterator is returned from FilterDefaultFeeSet and is used to iterate over the raw logs and unpacked data for DefaultFeeSet events raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorDefaultFeeSetIterator struct {
	Event *TeeFeeCalculatorDefaultFeeSet // Event containing the contract specifics and raw log

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
func (it *TeeFeeCalculatorDefaultFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeFeeCalculatorDefaultFeeSet)
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
		it.Event = new(TeeFeeCalculatorDefaultFeeSet)
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
func (it *TeeFeeCalculatorDefaultFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeFeeCalculatorDefaultFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeFeeCalculatorDefaultFeeSet represents a DefaultFeeSet event raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorDefaultFeeSet struct {
	DefaultFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDefaultFeeSet is a free log retrieval operation binding the contract event 0xe3c879f1bacd84281e6f3b2c940aee391b4ea5d58d41f2f9ae7808469ac38127.
//
// Solidity: event DefaultFeeSet(uint256 defaultFee)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) FilterDefaultFeeSet(opts *bind.FilterOpts) (*TeeFeeCalculatorDefaultFeeSetIterator, error) {

	logs, sub, err := _TeeFeeCalculator.contract.FilterLogs(opts, "DefaultFeeSet")
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorDefaultFeeSetIterator{contract: _TeeFeeCalculator.contract, event: "DefaultFeeSet", logs: logs, sub: sub}, nil
}

// WatchDefaultFeeSet is a free log subscription operation binding the contract event 0xe3c879f1bacd84281e6f3b2c940aee391b4ea5d58d41f2f9ae7808469ac38127.
//
// Solidity: event DefaultFeeSet(uint256 defaultFee)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) WatchDefaultFeeSet(opts *bind.WatchOpts, sink chan<- *TeeFeeCalculatorDefaultFeeSet) (event.Subscription, error) {

	logs, sub, err := _TeeFeeCalculator.contract.WatchLogs(opts, "DefaultFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeFeeCalculatorDefaultFeeSet)
				if err := _TeeFeeCalculator.contract.UnpackLog(event, "DefaultFeeSet", log); err != nil {
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

// ParseDefaultFeeSet is a log parse operation binding the contract event 0xe3c879f1bacd84281e6f3b2c940aee391b4ea5d58d41f2f9ae7808469ac38127.
//
// Solidity: event DefaultFeeSet(uint256 defaultFee)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) ParseDefaultFeeSet(log types.Log) (*TeeFeeCalculatorDefaultFeeSet, error) {
	event := new(TeeFeeCalculatorDefaultFeeSet)
	if err := _TeeFeeCalculator.contract.UnpackLog(event, "DefaultFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeFeeCalculatorGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorGovernanceCallTimelockedIterator struct {
	Event *TeeFeeCalculatorGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeeFeeCalculatorGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeFeeCalculatorGovernanceCallTimelocked)
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
		it.Event = new(TeeFeeCalculatorGovernanceCallTimelocked)
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
func (it *TeeFeeCalculatorGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeFeeCalculatorGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeFeeCalculatorGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeFeeCalculatorGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeFeeCalculator.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorGovernanceCallTimelockedIterator{contract: _TeeFeeCalculator.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeFeeCalculatorGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeFeeCalculator.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeFeeCalculatorGovernanceCallTimelocked)
				if err := _TeeFeeCalculator.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeFeeCalculatorGovernanceCallTimelocked, error) {
	event := new(TeeFeeCalculatorGovernanceCallTimelocked)
	if err := _TeeFeeCalculator.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeFeeCalculatorGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorGovernanceInitialisedIterator struct {
	Event *TeeFeeCalculatorGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeeFeeCalculatorGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeFeeCalculatorGovernanceInitialised)
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
		it.Event = new(TeeFeeCalculatorGovernanceInitialised)
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
func (it *TeeFeeCalculatorGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeFeeCalculatorGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeFeeCalculatorGovernanceInitialised represents a GovernanceInitialised event raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeFeeCalculatorGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeFeeCalculator.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorGovernanceInitialisedIterator{contract: _TeeFeeCalculator.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeFeeCalculatorGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeFeeCalculator.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeFeeCalculatorGovernanceInitialised)
				if err := _TeeFeeCalculator.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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

// ParseGovernanceInitialised is a log parse operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) ParseGovernanceInitialised(log types.Log) (*TeeFeeCalculatorGovernanceInitialised, error) {
	event := new(TeeFeeCalculatorGovernanceInitialised)
	if err := _TeeFeeCalculator.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeFeeCalculatorOperationFeesSetIterator is returned from FilterOperationFeesSet and is used to iterate over the raw logs and unpacked data for OperationFeesSet events raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorOperationFeesSetIterator struct {
	Event *TeeFeeCalculatorOperationFeesSet // Event containing the contract specifics and raw log

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
func (it *TeeFeeCalculatorOperationFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeFeeCalculatorOperationFeesSet)
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
		it.Event = new(TeeFeeCalculatorOperationFeesSet)
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
func (it *TeeFeeCalculatorOperationFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeFeeCalculatorOperationFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeFeeCalculatorOperationFeesSet represents a OperationFeesSet event raised by the TeeFeeCalculator contract.
type TeeFeeCalculatorOperationFeesSet struct {
	OpTypes    [][32]byte
	OpCommands [][32]byte
	Fees       []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperationFeesSet is a free log retrieval operation binding the contract event 0x1c36e02b57b3a0177cc8393cf08249650669360707214b808180ae8f18c98ca5.
//
// Solidity: event OperationFeesSet(bytes32[] opTypes, bytes32[] opCommands, uint256[] fees)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) FilterOperationFeesSet(opts *bind.FilterOpts) (*TeeFeeCalculatorOperationFeesSetIterator, error) {

	logs, sub, err := _TeeFeeCalculator.contract.FilterLogs(opts, "OperationFeesSet")
	if err != nil {
		return nil, err
	}
	return &TeeFeeCalculatorOperationFeesSetIterator{contract: _TeeFeeCalculator.contract, event: "OperationFeesSet", logs: logs, sub: sub}, nil
}

// WatchOperationFeesSet is a free log subscription operation binding the contract event 0x1c36e02b57b3a0177cc8393cf08249650669360707214b808180ae8f18c98ca5.
//
// Solidity: event OperationFeesSet(bytes32[] opTypes, bytes32[] opCommands, uint256[] fees)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) WatchOperationFeesSet(opts *bind.WatchOpts, sink chan<- *TeeFeeCalculatorOperationFeesSet) (event.Subscription, error) {

	logs, sub, err := _TeeFeeCalculator.contract.WatchLogs(opts, "OperationFeesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeFeeCalculatorOperationFeesSet)
				if err := _TeeFeeCalculator.contract.UnpackLog(event, "OperationFeesSet", log); err != nil {
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

// ParseOperationFeesSet is a log parse operation binding the contract event 0x1c36e02b57b3a0177cc8393cf08249650669360707214b808180ae8f18c98ca5.
//
// Solidity: event OperationFeesSet(bytes32[] opTypes, bytes32[] opCommands, uint256[] fees)
func (_TeeFeeCalculator *TeeFeeCalculatorFilterer) ParseOperationFeesSet(log types.Log) (*TeeFeeCalculatorOperationFeesSet, error) {
	event := new(TeeFeeCalculatorOperationFeesSet)
	if err := _TeeFeeCalculator.contract.UnpackLog(event, "OperationFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
