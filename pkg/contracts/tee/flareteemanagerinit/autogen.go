// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flareteemanagerinit

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

// FlareTeeManagerInitMetaData contains all meta data concerning the FlareTeeManagerInit contract.
var FlareTeeManagerInitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"graceSeconds\",\"type\":\"uint256\"}],\"name\":\"GracePeriodTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"graceSeconds\",\"type\":\"uint256\"}],\"name\":\"GracePeriodTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AllExtensionOwnersAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AllExtensionOwnersDisallowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultFee\",\"type\":\"uint256\"}],\"name\":\"DefaultFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"graceSeconds\",\"type\":\"uint256\"}],\"name\":\"EmergencyUnpauseGracePeriodSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"availabilityCheckValidityDurationSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"signingPolicyValidityDurationInRewardEpochs\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"challengeValidityDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"SettingsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_availabilityCheckValidityDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicyValidityDurationInRewardEpochs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_challengeValidityDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_defaultFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_publicExtensionCreationEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_emergencyUnpauseGracePeriodSeconds\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FlareTeeManagerInitABI is the input ABI used to generate the binding from.
// Deprecated: Use FlareTeeManagerInitMetaData.ABI instead.
var FlareTeeManagerInitABI = FlareTeeManagerInitMetaData.ABI

// FlareTeeManagerInit is an auto generated Go binding around an Ethereum contract.
type FlareTeeManagerInit struct {
	FlareTeeManagerInitCaller     // Read-only binding to the contract
	FlareTeeManagerInitTransactor // Write-only binding to the contract
	FlareTeeManagerInitFilterer   // Log filterer for contract events
}

// FlareTeeManagerInitCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlareTeeManagerInitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlareTeeManagerInitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlareTeeManagerInitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlareTeeManagerInitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlareTeeManagerInitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlareTeeManagerInitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlareTeeManagerInitSession struct {
	Contract     *FlareTeeManagerInit // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FlareTeeManagerInitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlareTeeManagerInitCallerSession struct {
	Contract *FlareTeeManagerInitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// FlareTeeManagerInitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlareTeeManagerInitTransactorSession struct {
	Contract     *FlareTeeManagerInitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// FlareTeeManagerInitRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlareTeeManagerInitRaw struct {
	Contract *FlareTeeManagerInit // Generic contract binding to access the raw methods on
}

// FlareTeeManagerInitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlareTeeManagerInitCallerRaw struct {
	Contract *FlareTeeManagerInitCaller // Generic read-only contract binding to access the raw methods on
}

// FlareTeeManagerInitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlareTeeManagerInitTransactorRaw struct {
	Contract *FlareTeeManagerInitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlareTeeManagerInit creates a new instance of FlareTeeManagerInit, bound to a specific deployed contract.
func NewFlareTeeManagerInit(address common.Address, backend bind.ContractBackend) (*FlareTeeManagerInit, error) {
	contract, err := bindFlareTeeManagerInit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlareTeeManagerInit{FlareTeeManagerInitCaller: FlareTeeManagerInitCaller{contract: contract}, FlareTeeManagerInitTransactor: FlareTeeManagerInitTransactor{contract: contract}, FlareTeeManagerInitFilterer: FlareTeeManagerInitFilterer{contract: contract}}, nil
}

// NewFlareTeeManagerInitCaller creates a new read-only instance of FlareTeeManagerInit, bound to a specific deployed contract.
func NewFlareTeeManagerInitCaller(address common.Address, caller bind.ContractCaller) (*FlareTeeManagerInitCaller, error) {
	contract, err := bindFlareTeeManagerInit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlareTeeManagerInitCaller{contract: contract}, nil
}

// NewFlareTeeManagerInitTransactor creates a new write-only instance of FlareTeeManagerInit, bound to a specific deployed contract.
func NewFlareTeeManagerInitTransactor(address common.Address, transactor bind.ContractTransactor) (*FlareTeeManagerInitTransactor, error) {
	contract, err := bindFlareTeeManagerInit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlareTeeManagerInitTransactor{contract: contract}, nil
}

// NewFlareTeeManagerInitFilterer creates a new log filterer instance of FlareTeeManagerInit, bound to a specific deployed contract.
func NewFlareTeeManagerInitFilterer(address common.Address, filterer bind.ContractFilterer) (*FlareTeeManagerInitFilterer, error) {
	contract, err := bindFlareTeeManagerInit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlareTeeManagerInitFilterer{contract: contract}, nil
}

// bindFlareTeeManagerInit binds a generic wrapper to an already deployed contract.
func bindFlareTeeManagerInit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FlareTeeManagerInitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlareTeeManagerInit *FlareTeeManagerInitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlareTeeManagerInit.Contract.FlareTeeManagerInitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlareTeeManagerInit *FlareTeeManagerInitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlareTeeManagerInit.Contract.FlareTeeManagerInitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlareTeeManagerInit *FlareTeeManagerInitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlareTeeManagerInit.Contract.FlareTeeManagerInitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlareTeeManagerInit *FlareTeeManagerInitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlareTeeManagerInit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlareTeeManagerInit *FlareTeeManagerInitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlareTeeManagerInit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlareTeeManagerInit *FlareTeeManagerInitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlareTeeManagerInit.Contract.contract.Transact(opts, method, params...)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FlareTeeManagerInit *FlareTeeManagerInitCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FlareTeeManagerInit.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FlareTeeManagerInit *FlareTeeManagerInitSession) GetAddressUpdater() (common.Address, error) {
	return _FlareTeeManagerInit.Contract.GetAddressUpdater(&_FlareTeeManagerInit.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FlareTeeManagerInit *FlareTeeManagerInitCallerSession) GetAddressUpdater() (common.Address, error) {
	return _FlareTeeManagerInit.Contract.GetAddressUpdater(&_FlareTeeManagerInit.CallOpts)
}

// Init is a paid mutator transaction binding the contract method 0x84e65dfc.
//
// Solidity: function init(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _availabilityCheckValidityDurationSeconds, uint64 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds, uint256 _defaultFee, bool _publicExtensionCreationEnabled, uint256 _emergencyUnpauseGracePeriodSeconds) returns()
func (_FlareTeeManagerInit *FlareTeeManagerInitTransactor) Init(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _availabilityCheckValidityDurationSeconds uint64, _signingPolicyValidityDurationInRewardEpochs uint64, _challengeValidityDurationSeconds uint64, _defaultFee *big.Int, _publicExtensionCreationEnabled bool, _emergencyUnpauseGracePeriodSeconds *big.Int) (*types.Transaction, error) {
	return _FlareTeeManagerInit.contract.Transact(opts, "init", _governanceSettings, _initialGovernance, _addressUpdater, _availabilityCheckValidityDurationSeconds, _signingPolicyValidityDurationInRewardEpochs, _challengeValidityDurationSeconds, _defaultFee, _publicExtensionCreationEnabled, _emergencyUnpauseGracePeriodSeconds)
}

// Init is a paid mutator transaction binding the contract method 0x84e65dfc.
//
// Solidity: function init(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _availabilityCheckValidityDurationSeconds, uint64 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds, uint256 _defaultFee, bool _publicExtensionCreationEnabled, uint256 _emergencyUnpauseGracePeriodSeconds) returns()
func (_FlareTeeManagerInit *FlareTeeManagerInitSession) Init(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _availabilityCheckValidityDurationSeconds uint64, _signingPolicyValidityDurationInRewardEpochs uint64, _challengeValidityDurationSeconds uint64, _defaultFee *big.Int, _publicExtensionCreationEnabled bool, _emergencyUnpauseGracePeriodSeconds *big.Int) (*types.Transaction, error) {
	return _FlareTeeManagerInit.Contract.Init(&_FlareTeeManagerInit.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _availabilityCheckValidityDurationSeconds, _signingPolicyValidityDurationInRewardEpochs, _challengeValidityDurationSeconds, _defaultFee, _publicExtensionCreationEnabled, _emergencyUnpauseGracePeriodSeconds)
}

// Init is a paid mutator transaction binding the contract method 0x84e65dfc.
//
// Solidity: function init(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _availabilityCheckValidityDurationSeconds, uint64 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds, uint256 _defaultFee, bool _publicExtensionCreationEnabled, uint256 _emergencyUnpauseGracePeriodSeconds) returns()
func (_FlareTeeManagerInit *FlareTeeManagerInitTransactorSession) Init(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _availabilityCheckValidityDurationSeconds uint64, _signingPolicyValidityDurationInRewardEpochs uint64, _challengeValidityDurationSeconds uint64, _defaultFee *big.Int, _publicExtensionCreationEnabled bool, _emergencyUnpauseGracePeriodSeconds *big.Int) (*types.Transaction, error) {
	return _FlareTeeManagerInit.Contract.Init(&_FlareTeeManagerInit.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _availabilityCheckValidityDurationSeconds, _signingPolicyValidityDurationInRewardEpochs, _challengeValidityDurationSeconds, _defaultFee, _publicExtensionCreationEnabled, _emergencyUnpauseGracePeriodSeconds)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FlareTeeManagerInit *FlareTeeManagerInitTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FlareTeeManagerInit.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FlareTeeManagerInit *FlareTeeManagerInitSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FlareTeeManagerInit.Contract.UpdateContractAddresses(&_FlareTeeManagerInit.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FlareTeeManagerInit *FlareTeeManagerInitTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FlareTeeManagerInit.Contract.UpdateContractAddresses(&_FlareTeeManagerInit.TransactOpts, _contractNameHashes, _contractAddresses)
}

// FlareTeeManagerInitAllExtensionOwnersAllowedIterator is returned from FilterAllExtensionOwnersAllowed and is used to iterate over the raw logs and unpacked data for AllExtensionOwnersAllowed events raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitAllExtensionOwnersAllowedIterator struct {
	Event *FlareTeeManagerInitAllExtensionOwnersAllowed // Event containing the contract specifics and raw log

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
func (it *FlareTeeManagerInitAllExtensionOwnersAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareTeeManagerInitAllExtensionOwnersAllowed)
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
		it.Event = new(FlareTeeManagerInitAllExtensionOwnersAllowed)
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
func (it *FlareTeeManagerInitAllExtensionOwnersAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareTeeManagerInitAllExtensionOwnersAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareTeeManagerInitAllExtensionOwnersAllowed represents a AllExtensionOwnersAllowed event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitAllExtensionOwnersAllowed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllExtensionOwnersAllowed is a free log retrieval operation binding the contract event 0x187952b217dbb616bfaa4b45d52b804e332e3af0cc43291f4c21105e6f99a1a9.
//
// Solidity: event AllExtensionOwnersAllowed()
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) FilterAllExtensionOwnersAllowed(opts *bind.FilterOpts) (*FlareTeeManagerInitAllExtensionOwnersAllowedIterator, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.FilterLogs(opts, "AllExtensionOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return &FlareTeeManagerInitAllExtensionOwnersAllowedIterator{contract: _FlareTeeManagerInit.contract, event: "AllExtensionOwnersAllowed", logs: logs, sub: sub}, nil
}

// WatchAllExtensionOwnersAllowed is a free log subscription operation binding the contract event 0x187952b217dbb616bfaa4b45d52b804e332e3af0cc43291f4c21105e6f99a1a9.
//
// Solidity: event AllExtensionOwnersAllowed()
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) WatchAllExtensionOwnersAllowed(opts *bind.WatchOpts, sink chan<- *FlareTeeManagerInitAllExtensionOwnersAllowed) (event.Subscription, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.WatchLogs(opts, "AllExtensionOwnersAllowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareTeeManagerInitAllExtensionOwnersAllowed)
				if err := _FlareTeeManagerInit.contract.UnpackLog(event, "AllExtensionOwnersAllowed", log); err != nil {
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

// ParseAllExtensionOwnersAllowed is a log parse operation binding the contract event 0x187952b217dbb616bfaa4b45d52b804e332e3af0cc43291f4c21105e6f99a1a9.
//
// Solidity: event AllExtensionOwnersAllowed()
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) ParseAllExtensionOwnersAllowed(log types.Log) (*FlareTeeManagerInitAllExtensionOwnersAllowed, error) {
	event := new(FlareTeeManagerInitAllExtensionOwnersAllowed)
	if err := _FlareTeeManagerInit.contract.UnpackLog(event, "AllExtensionOwnersAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareTeeManagerInitAllExtensionOwnersDisallowedIterator is returned from FilterAllExtensionOwnersDisallowed and is used to iterate over the raw logs and unpacked data for AllExtensionOwnersDisallowed events raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitAllExtensionOwnersDisallowedIterator struct {
	Event *FlareTeeManagerInitAllExtensionOwnersDisallowed // Event containing the contract specifics and raw log

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
func (it *FlareTeeManagerInitAllExtensionOwnersDisallowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareTeeManagerInitAllExtensionOwnersDisallowed)
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
		it.Event = new(FlareTeeManagerInitAllExtensionOwnersDisallowed)
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
func (it *FlareTeeManagerInitAllExtensionOwnersDisallowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareTeeManagerInitAllExtensionOwnersDisallowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareTeeManagerInitAllExtensionOwnersDisallowed represents a AllExtensionOwnersDisallowed event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitAllExtensionOwnersDisallowed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllExtensionOwnersDisallowed is a free log retrieval operation binding the contract event 0x26b7975912f19fefe2f76735b4a0d6d37eee9a185598010b9ba0a8cd2bbbe243.
//
// Solidity: event AllExtensionOwnersDisallowed()
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) FilterAllExtensionOwnersDisallowed(opts *bind.FilterOpts) (*FlareTeeManagerInitAllExtensionOwnersDisallowedIterator, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.FilterLogs(opts, "AllExtensionOwnersDisallowed")
	if err != nil {
		return nil, err
	}
	return &FlareTeeManagerInitAllExtensionOwnersDisallowedIterator{contract: _FlareTeeManagerInit.contract, event: "AllExtensionOwnersDisallowed", logs: logs, sub: sub}, nil
}

// WatchAllExtensionOwnersDisallowed is a free log subscription operation binding the contract event 0x26b7975912f19fefe2f76735b4a0d6d37eee9a185598010b9ba0a8cd2bbbe243.
//
// Solidity: event AllExtensionOwnersDisallowed()
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) WatchAllExtensionOwnersDisallowed(opts *bind.WatchOpts, sink chan<- *FlareTeeManagerInitAllExtensionOwnersDisallowed) (event.Subscription, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.WatchLogs(opts, "AllExtensionOwnersDisallowed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareTeeManagerInitAllExtensionOwnersDisallowed)
				if err := _FlareTeeManagerInit.contract.UnpackLog(event, "AllExtensionOwnersDisallowed", log); err != nil {
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

// ParseAllExtensionOwnersDisallowed is a log parse operation binding the contract event 0x26b7975912f19fefe2f76735b4a0d6d37eee9a185598010b9ba0a8cd2bbbe243.
//
// Solidity: event AllExtensionOwnersDisallowed()
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) ParseAllExtensionOwnersDisallowed(log types.Log) (*FlareTeeManagerInitAllExtensionOwnersDisallowed, error) {
	event := new(FlareTeeManagerInitAllExtensionOwnersDisallowed)
	if err := _FlareTeeManagerInit.contract.UnpackLog(event, "AllExtensionOwnersDisallowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareTeeManagerInitDefaultFeeSetIterator is returned from FilterDefaultFeeSet and is used to iterate over the raw logs and unpacked data for DefaultFeeSet events raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitDefaultFeeSetIterator struct {
	Event *FlareTeeManagerInitDefaultFeeSet // Event containing the contract specifics and raw log

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
func (it *FlareTeeManagerInitDefaultFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareTeeManagerInitDefaultFeeSet)
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
		it.Event = new(FlareTeeManagerInitDefaultFeeSet)
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
func (it *FlareTeeManagerInitDefaultFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareTeeManagerInitDefaultFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareTeeManagerInitDefaultFeeSet represents a DefaultFeeSet event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitDefaultFeeSet struct {
	DefaultFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDefaultFeeSet is a free log retrieval operation binding the contract event 0xe3c879f1bacd84281e6f3b2c940aee391b4ea5d58d41f2f9ae7808469ac38127.
//
// Solidity: event DefaultFeeSet(uint256 defaultFee)
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) FilterDefaultFeeSet(opts *bind.FilterOpts) (*FlareTeeManagerInitDefaultFeeSetIterator, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.FilterLogs(opts, "DefaultFeeSet")
	if err != nil {
		return nil, err
	}
	return &FlareTeeManagerInitDefaultFeeSetIterator{contract: _FlareTeeManagerInit.contract, event: "DefaultFeeSet", logs: logs, sub: sub}, nil
}

// WatchDefaultFeeSet is a free log subscription operation binding the contract event 0xe3c879f1bacd84281e6f3b2c940aee391b4ea5d58d41f2f9ae7808469ac38127.
//
// Solidity: event DefaultFeeSet(uint256 defaultFee)
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) WatchDefaultFeeSet(opts *bind.WatchOpts, sink chan<- *FlareTeeManagerInitDefaultFeeSet) (event.Subscription, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.WatchLogs(opts, "DefaultFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareTeeManagerInitDefaultFeeSet)
				if err := _FlareTeeManagerInit.contract.UnpackLog(event, "DefaultFeeSet", log); err != nil {
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
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) ParseDefaultFeeSet(log types.Log) (*FlareTeeManagerInitDefaultFeeSet, error) {
	event := new(FlareTeeManagerInitDefaultFeeSet)
	if err := _FlareTeeManagerInit.contract.UnpackLog(event, "DefaultFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareTeeManagerInitEmergencyUnpauseGracePeriodSetIterator is returned from FilterEmergencyUnpauseGracePeriodSet and is used to iterate over the raw logs and unpacked data for EmergencyUnpauseGracePeriodSet events raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitEmergencyUnpauseGracePeriodSetIterator struct {
	Event *FlareTeeManagerInitEmergencyUnpauseGracePeriodSet // Event containing the contract specifics and raw log

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
func (it *FlareTeeManagerInitEmergencyUnpauseGracePeriodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareTeeManagerInitEmergencyUnpauseGracePeriodSet)
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
		it.Event = new(FlareTeeManagerInitEmergencyUnpauseGracePeriodSet)
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
func (it *FlareTeeManagerInitEmergencyUnpauseGracePeriodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareTeeManagerInitEmergencyUnpauseGracePeriodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareTeeManagerInitEmergencyUnpauseGracePeriodSet represents a EmergencyUnpauseGracePeriodSet event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitEmergencyUnpauseGracePeriodSet struct {
	GraceSeconds *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEmergencyUnpauseGracePeriodSet is a free log retrieval operation binding the contract event 0xd3e133eed54192c593c1994bcda078ca197c9b9af58465e5e0d84296e23f0b16.
//
// Solidity: event EmergencyUnpauseGracePeriodSet(uint256 graceSeconds)
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) FilterEmergencyUnpauseGracePeriodSet(opts *bind.FilterOpts) (*FlareTeeManagerInitEmergencyUnpauseGracePeriodSetIterator, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.FilterLogs(opts, "EmergencyUnpauseGracePeriodSet")
	if err != nil {
		return nil, err
	}
	return &FlareTeeManagerInitEmergencyUnpauseGracePeriodSetIterator{contract: _FlareTeeManagerInit.contract, event: "EmergencyUnpauseGracePeriodSet", logs: logs, sub: sub}, nil
}

// WatchEmergencyUnpauseGracePeriodSet is a free log subscription operation binding the contract event 0xd3e133eed54192c593c1994bcda078ca197c9b9af58465e5e0d84296e23f0b16.
//
// Solidity: event EmergencyUnpauseGracePeriodSet(uint256 graceSeconds)
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) WatchEmergencyUnpauseGracePeriodSet(opts *bind.WatchOpts, sink chan<- *FlareTeeManagerInitEmergencyUnpauseGracePeriodSet) (event.Subscription, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.WatchLogs(opts, "EmergencyUnpauseGracePeriodSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareTeeManagerInitEmergencyUnpauseGracePeriodSet)
				if err := _FlareTeeManagerInit.contract.UnpackLog(event, "EmergencyUnpauseGracePeriodSet", log); err != nil {
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

// ParseEmergencyUnpauseGracePeriodSet is a log parse operation binding the contract event 0xd3e133eed54192c593c1994bcda078ca197c9b9af58465e5e0d84296e23f0b16.
//
// Solidity: event EmergencyUnpauseGracePeriodSet(uint256 graceSeconds)
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) ParseEmergencyUnpauseGracePeriodSet(log types.Log) (*FlareTeeManagerInitEmergencyUnpauseGracePeriodSet, error) {
	event := new(FlareTeeManagerInitEmergencyUnpauseGracePeriodSet)
	if err := _FlareTeeManagerInit.contract.UnpackLog(event, "EmergencyUnpauseGracePeriodSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareTeeManagerInitGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitGovernanceInitialisedIterator struct {
	Event *FlareTeeManagerInitGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *FlareTeeManagerInitGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareTeeManagerInitGovernanceInitialised)
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
		it.Event = new(FlareTeeManagerInitGovernanceInitialised)
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
func (it *FlareTeeManagerInitGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareTeeManagerInitGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareTeeManagerInitGovernanceInitialised represents a GovernanceInitialised event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*FlareTeeManagerInitGovernanceInitialisedIterator, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &FlareTeeManagerInitGovernanceInitialisedIterator{contract: _FlareTeeManagerInit.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *FlareTeeManagerInitGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareTeeManagerInitGovernanceInitialised)
				if err := _FlareTeeManagerInit.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) ParseGovernanceInitialised(log types.Log) (*FlareTeeManagerInitGovernanceInitialised, error) {
	event := new(FlareTeeManagerInitGovernanceInitialised)
	if err := _FlareTeeManagerInit.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareTeeManagerInitInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitInitializedIterator struct {
	Event *FlareTeeManagerInitInitialized // Event containing the contract specifics and raw log

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
func (it *FlareTeeManagerInitInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareTeeManagerInitInitialized)
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
		it.Event = new(FlareTeeManagerInitInitialized)
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
func (it *FlareTeeManagerInitInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareTeeManagerInitInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareTeeManagerInitInitialized represents a Initialized event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) FilterInitialized(opts *bind.FilterOpts) (*FlareTeeManagerInitInitializedIterator, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FlareTeeManagerInitInitializedIterator{contract: _FlareTeeManagerInit.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FlareTeeManagerInitInitialized) (event.Subscription, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareTeeManagerInitInitialized)
				if err := _FlareTeeManagerInit.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) ParseInitialized(log types.Log) (*FlareTeeManagerInitInitialized, error) {
	event := new(FlareTeeManagerInitInitialized)
	if err := _FlareTeeManagerInit.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlareTeeManagerInitSettingsUpdatedIterator is returned from FilterSettingsUpdated and is used to iterate over the raw logs and unpacked data for SettingsUpdated events raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitSettingsUpdatedIterator struct {
	Event *FlareTeeManagerInitSettingsUpdated // Event containing the contract specifics and raw log

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
func (it *FlareTeeManagerInitSettingsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlareTeeManagerInitSettingsUpdated)
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
		it.Event = new(FlareTeeManagerInitSettingsUpdated)
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
func (it *FlareTeeManagerInitSettingsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlareTeeManagerInitSettingsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlareTeeManagerInitSettingsUpdated represents a SettingsUpdated event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitSettingsUpdated struct {
	AvailabilityCheckValidityDurationSeconds    uint64
	SigningPolicyValidityDurationInRewardEpochs uint64
	ChallengeValidityDurationSeconds            uint64
	Raw                                         types.Log // Blockchain specific contextual infos
}

// FilterSettingsUpdated is a free log retrieval operation binding the contract event 0x61d24ae6ba3e0c2182628db244b31540bc1b1ffdc9f16fd32ffa52caeb9868b9.
//
// Solidity: event SettingsUpdated(uint64 availabilityCheckValidityDurationSeconds, uint64 signingPolicyValidityDurationInRewardEpochs, uint64 challengeValidityDurationSeconds)
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) FilterSettingsUpdated(opts *bind.FilterOpts) (*FlareTeeManagerInitSettingsUpdatedIterator, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.FilterLogs(opts, "SettingsUpdated")
	if err != nil {
		return nil, err
	}
	return &FlareTeeManagerInitSettingsUpdatedIterator{contract: _FlareTeeManagerInit.contract, event: "SettingsUpdated", logs: logs, sub: sub}, nil
}

// WatchSettingsUpdated is a free log subscription operation binding the contract event 0x61d24ae6ba3e0c2182628db244b31540bc1b1ffdc9f16fd32ffa52caeb9868b9.
//
// Solidity: event SettingsUpdated(uint64 availabilityCheckValidityDurationSeconds, uint64 signingPolicyValidityDurationInRewardEpochs, uint64 challengeValidityDurationSeconds)
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) WatchSettingsUpdated(opts *bind.WatchOpts, sink chan<- *FlareTeeManagerInitSettingsUpdated) (event.Subscription, error) {

	logs, sub, err := _FlareTeeManagerInit.contract.WatchLogs(opts, "SettingsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlareTeeManagerInitSettingsUpdated)
				if err := _FlareTeeManagerInit.contract.UnpackLog(event, "SettingsUpdated", log); err != nil {
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

// ParseSettingsUpdated is a log parse operation binding the contract event 0x61d24ae6ba3e0c2182628db244b31540bc1b1ffdc9f16fd32ffa52caeb9868b9.
//
// Solidity: event SettingsUpdated(uint64 availabilityCheckValidityDurationSeconds, uint64 signingPolicyValidityDurationInRewardEpochs, uint64 challengeValidityDurationSeconds)
func (_FlareTeeManagerInit *FlareTeeManagerInitFilterer) ParseSettingsUpdated(log types.Log) (*FlareTeeManagerInitSettingsUpdated, error) {
	event := new(FlareTeeManagerInitSettingsUpdated)
	if err := _FlareTeeManagerInit.contract.UnpackLog(event, "SettingsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
