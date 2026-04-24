// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teevrf

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

// IMachineManagerTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
}

// TeeVRFMetaData contains all meta data concerning the TeeVRF contract.
var TeeVRFMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeesForKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonceEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuthorizationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizationAddress\",\"type\":\"address\"}],\"name\":\"VrfAuthorizationAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"name\":\"VrfRequested\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getVrfAuthorizationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_nonce\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestVrf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"name\":\"setVrfAuthorizationAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeVRFABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeVRFMetaData.ABI instead.
var TeeVRFABI = TeeVRFMetaData.ABI

// TeeVRF is an auto generated Go binding around an Ethereum contract.
type TeeVRF struct {
	TeeVRFCaller     // Read-only binding to the contract
	TeeVRFTransactor // Write-only binding to the contract
	TeeVRFFilterer   // Log filterer for contract events
}

// TeeVRFCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeVRFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVRFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeVRFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVRFFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeVRFFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVRFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeVRFSession struct {
	Contract     *TeeVRF           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeVRFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeVRFCallerSession struct {
	Contract *TeeVRFCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TeeVRFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeVRFTransactorSession struct {
	Contract     *TeeVRFTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeVRFRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeVRFRaw struct {
	Contract *TeeVRF // Generic contract binding to access the raw methods on
}

// TeeVRFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeVRFCallerRaw struct {
	Contract *TeeVRFCaller // Generic read-only contract binding to access the raw methods on
}

// TeeVRFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeVRFTransactorRaw struct {
	Contract *TeeVRFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeVRF creates a new instance of TeeVRF, bound to a specific deployed contract.
func NewTeeVRF(address common.Address, backend bind.ContractBackend) (*TeeVRF, error) {
	contract, err := bindTeeVRF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeVRF{TeeVRFCaller: TeeVRFCaller{contract: contract}, TeeVRFTransactor: TeeVRFTransactor{contract: contract}, TeeVRFFilterer: TeeVRFFilterer{contract: contract}}, nil
}

// NewTeeVRFCaller creates a new read-only instance of TeeVRF, bound to a specific deployed contract.
func NewTeeVRFCaller(address common.Address, caller bind.ContractCaller) (*TeeVRFCaller, error) {
	contract, err := bindTeeVRF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeVRFCaller{contract: contract}, nil
}

// NewTeeVRFTransactor creates a new write-only instance of TeeVRF, bound to a specific deployed contract.
func NewTeeVRFTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeVRFTransactor, error) {
	contract, err := bindTeeVRF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeVRFTransactor{contract: contract}, nil
}

// NewTeeVRFFilterer creates a new log filterer instance of TeeVRF, bound to a specific deployed contract.
func NewTeeVRFFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeVRFFilterer, error) {
	contract, err := bindTeeVRF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeVRFFilterer{contract: contract}, nil
}

// bindTeeVRF binds a generic wrapper to an already deployed contract.
func bindTeeVRF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeVRFMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeVRF *TeeVRFRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeVRF.Contract.TeeVRFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeVRF *TeeVRFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVRF.Contract.TeeVRFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeVRF *TeeVRFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeVRF.Contract.TeeVRFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeVRF *TeeVRFCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeVRF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeVRF *TeeVRFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVRF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeVRF *TeeVRFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeVRF.Contract.contract.Transact(opts, method, params...)
}

// GetVrfAuthorizationAddress is a free data retrieval call binding the contract method 0x918945a0.
//
// Solidity: function getVrfAuthorizationAddress(bytes32 _walletId) view returns(address)
func (_TeeVRF *TeeVRFCaller) GetVrfAuthorizationAddress(opts *bind.CallOpts, _walletId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TeeVRF.contract.Call(opts, &out, "getVrfAuthorizationAddress", _walletId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVrfAuthorizationAddress is a free data retrieval call binding the contract method 0x918945a0.
//
// Solidity: function getVrfAuthorizationAddress(bytes32 _walletId) view returns(address)
func (_TeeVRF *TeeVRFSession) GetVrfAuthorizationAddress(_walletId [32]byte) (common.Address, error) {
	return _TeeVRF.Contract.GetVrfAuthorizationAddress(&_TeeVRF.CallOpts, _walletId)
}

// GetVrfAuthorizationAddress is a free data retrieval call binding the contract method 0x918945a0.
//
// Solidity: function getVrfAuthorizationAddress(bytes32 _walletId) view returns(address)
func (_TeeVRF *TeeVRFCallerSession) GetVrfAuthorizationAddress(_walletId [32]byte) (common.Address, error) {
	return _TeeVRF.Contract.GetVrfAuthorizationAddress(&_TeeVRF.CallOpts, _walletId)
}

// RequestVrf is a paid mutator transaction binding the contract method 0xd7896b11.
//
// Solidity: function requestVrf(bytes32 _walletId, uint64 _keyId, bytes _nonce, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (_TeeVRF *TeeVRFTransactor) RequestVrf(opts *bind.TransactOpts, _walletId [32]byte, _keyId uint64, _nonce []byte, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeVRF.contract.Transact(opts, "requestVrf", _walletId, _keyId, _nonce, _claimBackAddress)
}

// RequestVrf is a paid mutator transaction binding the contract method 0xd7896b11.
//
// Solidity: function requestVrf(bytes32 _walletId, uint64 _keyId, bytes _nonce, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (_TeeVRF *TeeVRFSession) RequestVrf(_walletId [32]byte, _keyId uint64, _nonce []byte, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeVRF.Contract.RequestVrf(&_TeeVRF.TransactOpts, _walletId, _keyId, _nonce, _claimBackAddress)
}

// RequestVrf is a paid mutator transaction binding the contract method 0xd7896b11.
//
// Solidity: function requestVrf(bytes32 _walletId, uint64 _keyId, bytes _nonce, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (_TeeVRF *TeeVRFTransactorSession) RequestVrf(_walletId [32]byte, _keyId uint64, _nonce []byte, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeVRF.Contract.RequestVrf(&_TeeVRF.TransactOpts, _walletId, _keyId, _nonce, _claimBackAddress)
}

// SetVrfAuthorizationAddress is a paid mutator transaction binding the contract method 0x76b87bcf.
//
// Solidity: function setVrfAuthorizationAddress(bytes32 _walletId, address _authorizationAddress) returns()
func (_TeeVRF *TeeVRFTransactor) SetVrfAuthorizationAddress(opts *bind.TransactOpts, _walletId [32]byte, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeeVRF.contract.Transact(opts, "setVrfAuthorizationAddress", _walletId, _authorizationAddress)
}

// SetVrfAuthorizationAddress is a paid mutator transaction binding the contract method 0x76b87bcf.
//
// Solidity: function setVrfAuthorizationAddress(bytes32 _walletId, address _authorizationAddress) returns()
func (_TeeVRF *TeeVRFSession) SetVrfAuthorizationAddress(_walletId [32]byte, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeeVRF.Contract.SetVrfAuthorizationAddress(&_TeeVRF.TransactOpts, _walletId, _authorizationAddress)
}

// SetVrfAuthorizationAddress is a paid mutator transaction binding the contract method 0x76b87bcf.
//
// Solidity: function setVrfAuthorizationAddress(bytes32 _walletId, address _authorizationAddress) returns()
func (_TeeVRF *TeeVRFTransactorSession) SetVrfAuthorizationAddress(_walletId [32]byte, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeeVRF.Contract.SetVrfAuthorizationAddress(&_TeeVRF.TransactOpts, _walletId, _authorizationAddress)
}

// TeeVRFTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the TeeVRF contract.
type TeeVRFTeeInstructionsSentIterator struct {
	Event *TeeVRFTeeInstructionsSent // Event containing the contract specifics and raw log

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
func (it *TeeVRFTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVRFTeeInstructionsSent)
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
		it.Event = new(TeeVRFTeeInstructionsSent)
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
func (it *TeeVRFTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVRFTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVRFTeeInstructionsSent represents a TeeInstructionsSent event raised by the TeeVRF contract.
type TeeVRFTeeInstructionsSent struct {
	ExtensionId        *big.Int
	InstructionId      [32]byte
	RewardEpochId      uint32
	TeeMachines        []IMachineManagerTeeMachine
	OpType             [32]byte
	OpCommand          [32]byte
	Message            []byte
	Cosigners          []common.Address
	CosignersThreshold uint64
	ClaimBackAddress   common.Address
	Fee                *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTeeInstructionsSent is a free log retrieval operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_TeeVRF *TeeVRFFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (*TeeVRFTeeInstructionsSentIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeVRF.contract.FilterLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeVRFTeeInstructionsSentIterator{contract: _TeeVRF.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_TeeVRF *TeeVRFFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *TeeVRFTeeInstructionsSent, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeVRF.contract.WatchLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVRFTeeInstructionsSent)
				if err := _TeeVRF.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
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

// ParseTeeInstructionsSent is a log parse operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_TeeVRF *TeeVRFFilterer) ParseTeeInstructionsSent(log types.Log) (*TeeVRFTeeInstructionsSent, error) {
	event := new(TeeVRFTeeInstructionsSent)
	if err := _TeeVRF.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVRFVrfAuthorizationAddressSetIterator is returned from FilterVrfAuthorizationAddressSet and is used to iterate over the raw logs and unpacked data for VrfAuthorizationAddressSet events raised by the TeeVRF contract.
type TeeVRFVrfAuthorizationAddressSetIterator struct {
	Event *TeeVRFVrfAuthorizationAddressSet // Event containing the contract specifics and raw log

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
func (it *TeeVRFVrfAuthorizationAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVRFVrfAuthorizationAddressSet)
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
		it.Event = new(TeeVRFVrfAuthorizationAddressSet)
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
func (it *TeeVRFVrfAuthorizationAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVRFVrfAuthorizationAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVRFVrfAuthorizationAddressSet represents a VrfAuthorizationAddressSet event raised by the TeeVRF contract.
type TeeVRFVrfAuthorizationAddressSet struct {
	WalletId             [32]byte
	AuthorizationAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterVrfAuthorizationAddressSet is a free log retrieval operation binding the contract event 0xdf8b8dac41b0ec7f0b1415d39e5f0885052aa11dbb18bb13043d64a3f1cb9c3b.
//
// Solidity: event VrfAuthorizationAddressSet(bytes32 indexed walletId, address authorizationAddress)
func (_TeeVRF *TeeVRFFilterer) FilterVrfAuthorizationAddressSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeeVRFVrfAuthorizationAddressSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeVRF.contract.FilterLogs(opts, "VrfAuthorizationAddressSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeVRFVrfAuthorizationAddressSetIterator{contract: _TeeVRF.contract, event: "VrfAuthorizationAddressSet", logs: logs, sub: sub}, nil
}

// WatchVrfAuthorizationAddressSet is a free log subscription operation binding the contract event 0xdf8b8dac41b0ec7f0b1415d39e5f0885052aa11dbb18bb13043d64a3f1cb9c3b.
//
// Solidity: event VrfAuthorizationAddressSet(bytes32 indexed walletId, address authorizationAddress)
func (_TeeVRF *TeeVRFFilterer) WatchVrfAuthorizationAddressSet(opts *bind.WatchOpts, sink chan<- *TeeVRFVrfAuthorizationAddressSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeVRF.contract.WatchLogs(opts, "VrfAuthorizationAddressSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVRFVrfAuthorizationAddressSet)
				if err := _TeeVRF.contract.UnpackLog(event, "VrfAuthorizationAddressSet", log); err != nil {
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

// ParseVrfAuthorizationAddressSet is a log parse operation binding the contract event 0xdf8b8dac41b0ec7f0b1415d39e5f0885052aa11dbb18bb13043d64a3f1cb9c3b.
//
// Solidity: event VrfAuthorizationAddressSet(bytes32 indexed walletId, address authorizationAddress)
func (_TeeVRF *TeeVRFFilterer) ParseVrfAuthorizationAddressSet(log types.Log) (*TeeVRFVrfAuthorizationAddressSet, error) {
	event := new(TeeVRFVrfAuthorizationAddressSet)
	if err := _TeeVRF.contract.UnpackLog(event, "VrfAuthorizationAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVRFVrfRequestedIterator is returned from FilterVrfRequested and is used to iterate over the raw logs and unpacked data for VrfRequested events raised by the TeeVRF contract.
type TeeVRFVrfRequestedIterator struct {
	Event *TeeVRFVrfRequested // Event containing the contract specifics and raw log

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
func (it *TeeVRFVrfRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVRFVrfRequested)
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
		it.Event = new(TeeVRFVrfRequested)
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
func (it *TeeVRFVrfRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVRFVrfRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVRFVrfRequested represents a VrfRequested event raised by the TeeVRF contract.
type TeeVRFVrfRequested struct {
	WalletId      [32]byte
	KeyId         uint64
	InstructionId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVrfRequested is a free log retrieval operation binding the contract event 0x52749cdb7f850186051c4aff58b749e22b3e30c23dbd43ac8c8f7a9c5eec5ddd.
//
// Solidity: event VrfRequested(bytes32 indexed walletId, uint64 keyId, bytes32 instructionId)
func (_TeeVRF *TeeVRFFilterer) FilterVrfRequested(opts *bind.FilterOpts, walletId [][32]byte) (*TeeVRFVrfRequestedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeVRF.contract.FilterLogs(opts, "VrfRequested", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeVRFVrfRequestedIterator{contract: _TeeVRF.contract, event: "VrfRequested", logs: logs, sub: sub}, nil
}

// WatchVrfRequested is a free log subscription operation binding the contract event 0x52749cdb7f850186051c4aff58b749e22b3e30c23dbd43ac8c8f7a9c5eec5ddd.
//
// Solidity: event VrfRequested(bytes32 indexed walletId, uint64 keyId, bytes32 instructionId)
func (_TeeVRF *TeeVRFFilterer) WatchVrfRequested(opts *bind.WatchOpts, sink chan<- *TeeVRFVrfRequested, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeVRF.contract.WatchLogs(opts, "VrfRequested", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVRFVrfRequested)
				if err := _TeeVRF.contract.UnpackLog(event, "VrfRequested", log); err != nil {
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

// ParseVrfRequested is a log parse operation binding the contract event 0x52749cdb7f850186051c4aff58b749e22b3e30c23dbd43ac8c8f7a9c5eec5ddd.
//
// Solidity: event VrfRequested(bytes32 indexed walletId, uint64 keyId, bytes32 instructionId)
func (_TeeVRF *TeeVRFFilterer) ParseVrfRequested(log types.Log) (*TeeVRFVrfRequested, error) {
	event := new(TeeVRFVrfRequested)
	if err := _TeeVRF.contract.UnpackLog(event, "VrfRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
