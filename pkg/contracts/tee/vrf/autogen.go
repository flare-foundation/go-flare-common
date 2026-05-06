// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf

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

// VRFMetaData contains all meta data concerning the VRF contract.
var VRFMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeesForKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonceEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuthorizationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizationAddress\",\"type\":\"address\"}],\"name\":\"VrfAuthorizationAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"name\":\"VrfRequested\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getVrfAuthorizationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_nonce\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestVrf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"name\":\"setVrfAuthorizationAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VRFABI is the input ABI used to generate the binding from.
// Deprecated: Use VRFMetaData.ABI instead.
var VRFABI = VRFMetaData.ABI

// VRF is an auto generated Go binding around an Ethereum contract.
type VRF struct {
	VRFCaller     // Read-only binding to the contract
	VRFTransactor // Write-only binding to the contract
	VRFFilterer   // Log filterer for contract events
}

// VRFCaller is an auto generated read-only Go binding around an Ethereum contract.
type VRFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFSession struct {
	Contract     *VRF              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFCallerSession struct {
	Contract *VRFCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VRFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFTransactorSession struct {
	Contract     *VRFTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFRaw is an auto generated low-level Go binding around an Ethereum contract.
type VRFRaw struct {
	Contract *VRF // Generic contract binding to access the raw methods on
}

// VRFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFCallerRaw struct {
	Contract *VRFCaller // Generic read-only contract binding to access the raw methods on
}

// VRFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFTransactorRaw struct {
	Contract *VRFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVRF creates a new instance of VRF, bound to a specific deployed contract.
func NewVRF(address common.Address, backend bind.ContractBackend) (*VRF, error) {
	contract, err := bindVRF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRF{VRFCaller: VRFCaller{contract: contract}, VRFTransactor: VRFTransactor{contract: contract}, VRFFilterer: VRFFilterer{contract: contract}}, nil
}

// NewVRFCaller creates a new read-only instance of VRF, bound to a specific deployed contract.
func NewVRFCaller(address common.Address, caller bind.ContractCaller) (*VRFCaller, error) {
	contract, err := bindVRF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCaller{contract: contract}, nil
}

// NewVRFTransactor creates a new write-only instance of VRF, bound to a specific deployed contract.
func NewVRFTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFTransactor, error) {
	contract, err := bindVRF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFTransactor{contract: contract}, nil
}

// NewVRFFilterer creates a new log filterer instance of VRF, bound to a specific deployed contract.
func NewVRFFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFFilterer, error) {
	contract, err := bindVRF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFFilterer{contract: contract}, nil
}

// bindVRF binds a generic wrapper to an already deployed contract.
func bindVRF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRF *VRFRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRF.Contract.VRFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRF *VRFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRF.Contract.VRFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRF *VRFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRF.Contract.VRFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRF *VRFCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRF *VRFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRF *VRFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRF.Contract.contract.Transact(opts, method, params...)
}

// GetVrfAuthorizationAddress is a free data retrieval call binding the contract method 0x918945a0.
//
// Solidity: function getVrfAuthorizationAddress(bytes32 _walletId) view returns(address)
func (_VRF *VRFCaller) GetVrfAuthorizationAddress(opts *bind.CallOpts, _walletId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _VRF.contract.Call(opts, &out, "getVrfAuthorizationAddress", _walletId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVrfAuthorizationAddress is a free data retrieval call binding the contract method 0x918945a0.
//
// Solidity: function getVrfAuthorizationAddress(bytes32 _walletId) view returns(address)
func (_VRF *VRFSession) GetVrfAuthorizationAddress(_walletId [32]byte) (common.Address, error) {
	return _VRF.Contract.GetVrfAuthorizationAddress(&_VRF.CallOpts, _walletId)
}

// GetVrfAuthorizationAddress is a free data retrieval call binding the contract method 0x918945a0.
//
// Solidity: function getVrfAuthorizationAddress(bytes32 _walletId) view returns(address)
func (_VRF *VRFCallerSession) GetVrfAuthorizationAddress(_walletId [32]byte) (common.Address, error) {
	return _VRF.Contract.GetVrfAuthorizationAddress(&_VRF.CallOpts, _walletId)
}

// RequestVrf is a paid mutator transaction binding the contract method 0xd7896b11.
//
// Solidity: function requestVrf(bytes32 _walletId, uint64 _keyId, bytes _nonce, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (_VRF *VRFTransactor) RequestVrf(opts *bind.TransactOpts, _walletId [32]byte, _keyId uint64, _nonce []byte, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "requestVrf", _walletId, _keyId, _nonce, _claimBackAddress)
}

// RequestVrf is a paid mutator transaction binding the contract method 0xd7896b11.
//
// Solidity: function requestVrf(bytes32 _walletId, uint64 _keyId, bytes _nonce, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (_VRF *VRFSession) RequestVrf(_walletId [32]byte, _keyId uint64, _nonce []byte, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _VRF.Contract.RequestVrf(&_VRF.TransactOpts, _walletId, _keyId, _nonce, _claimBackAddress)
}

// RequestVrf is a paid mutator transaction binding the contract method 0xd7896b11.
//
// Solidity: function requestVrf(bytes32 _walletId, uint64 _keyId, bytes _nonce, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (_VRF *VRFTransactorSession) RequestVrf(_walletId [32]byte, _keyId uint64, _nonce []byte, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _VRF.Contract.RequestVrf(&_VRF.TransactOpts, _walletId, _keyId, _nonce, _claimBackAddress)
}

// SetVrfAuthorizationAddress is a paid mutator transaction binding the contract method 0x76b87bcf.
//
// Solidity: function setVrfAuthorizationAddress(bytes32 _walletId, address _authorizationAddress) returns()
func (_VRF *VRFTransactor) SetVrfAuthorizationAddress(opts *bind.TransactOpts, _walletId [32]byte, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _VRF.contract.Transact(opts, "setVrfAuthorizationAddress", _walletId, _authorizationAddress)
}

// SetVrfAuthorizationAddress is a paid mutator transaction binding the contract method 0x76b87bcf.
//
// Solidity: function setVrfAuthorizationAddress(bytes32 _walletId, address _authorizationAddress) returns()
func (_VRF *VRFSession) SetVrfAuthorizationAddress(_walletId [32]byte, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _VRF.Contract.SetVrfAuthorizationAddress(&_VRF.TransactOpts, _walletId, _authorizationAddress)
}

// SetVrfAuthorizationAddress is a paid mutator transaction binding the contract method 0x76b87bcf.
//
// Solidity: function setVrfAuthorizationAddress(bytes32 _walletId, address _authorizationAddress) returns()
func (_VRF *VRFTransactorSession) SetVrfAuthorizationAddress(_walletId [32]byte, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _VRF.Contract.SetVrfAuthorizationAddress(&_VRF.TransactOpts, _walletId, _authorizationAddress)
}

// VRFTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the VRF contract.
type VRFTeeInstructionsSentIterator struct {
	Event *VRFTeeInstructionsSent // Event containing the contract specifics and raw log

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
func (it *VRFTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFTeeInstructionsSent)
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
		it.Event = new(VRFTeeInstructionsSent)
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
func (it *VRFTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFTeeInstructionsSent represents a TeeInstructionsSent event raised by the VRF contract.
type VRFTeeInstructionsSent struct {
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
func (_VRF *VRFFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (*VRFTeeInstructionsSentIterator, error) {

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

	logs, sub, err := _VRF.contract.FilterLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFTeeInstructionsSentIterator{contract: _VRF.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_VRF *VRFFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *VRFTeeInstructionsSent, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

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

	logs, sub, err := _VRF.contract.WatchLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFTeeInstructionsSent)
				if err := _VRF.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
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
func (_VRF *VRFFilterer) ParseTeeInstructionsSent(log types.Log) (*VRFTeeInstructionsSent, error) {
	event := new(VRFTeeInstructionsSent)
	if err := _VRF.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFVrfAuthorizationAddressSetIterator is returned from FilterVrfAuthorizationAddressSet and is used to iterate over the raw logs and unpacked data for VrfAuthorizationAddressSet events raised by the VRF contract.
type VRFVrfAuthorizationAddressSetIterator struct {
	Event *VRFVrfAuthorizationAddressSet // Event containing the contract specifics and raw log

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
func (it *VRFVrfAuthorizationAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFVrfAuthorizationAddressSet)
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
		it.Event = new(VRFVrfAuthorizationAddressSet)
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
func (it *VRFVrfAuthorizationAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFVrfAuthorizationAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFVrfAuthorizationAddressSet represents a VrfAuthorizationAddressSet event raised by the VRF contract.
type VRFVrfAuthorizationAddressSet struct {
	WalletId             [32]byte
	AuthorizationAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterVrfAuthorizationAddressSet is a free log retrieval operation binding the contract event 0xdf8b8dac41b0ec7f0b1415d39e5f0885052aa11dbb18bb13043d64a3f1cb9c3b.
//
// Solidity: event VrfAuthorizationAddressSet(bytes32 indexed walletId, address authorizationAddress)
func (_VRF *VRFFilterer) FilterVrfAuthorizationAddressSet(opts *bind.FilterOpts, walletId [][32]byte) (*VRFVrfAuthorizationAddressSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _VRF.contract.FilterLogs(opts, "VrfAuthorizationAddressSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFVrfAuthorizationAddressSetIterator{contract: _VRF.contract, event: "VrfAuthorizationAddressSet", logs: logs, sub: sub}, nil
}

// WatchVrfAuthorizationAddressSet is a free log subscription operation binding the contract event 0xdf8b8dac41b0ec7f0b1415d39e5f0885052aa11dbb18bb13043d64a3f1cb9c3b.
//
// Solidity: event VrfAuthorizationAddressSet(bytes32 indexed walletId, address authorizationAddress)
func (_VRF *VRFFilterer) WatchVrfAuthorizationAddressSet(opts *bind.WatchOpts, sink chan<- *VRFVrfAuthorizationAddressSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _VRF.contract.WatchLogs(opts, "VrfAuthorizationAddressSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFVrfAuthorizationAddressSet)
				if err := _VRF.contract.UnpackLog(event, "VrfAuthorizationAddressSet", log); err != nil {
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
func (_VRF *VRFFilterer) ParseVrfAuthorizationAddressSet(log types.Log) (*VRFVrfAuthorizationAddressSet, error) {
	event := new(VRFVrfAuthorizationAddressSet)
	if err := _VRF.contract.UnpackLog(event, "VrfAuthorizationAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VRFVrfRequestedIterator is returned from FilterVrfRequested and is used to iterate over the raw logs and unpacked data for VrfRequested events raised by the VRF contract.
type VRFVrfRequestedIterator struct {
	Event *VRFVrfRequested // Event containing the contract specifics and raw log

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
func (it *VRFVrfRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFVrfRequested)
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
		it.Event = new(VRFVrfRequested)
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
func (it *VRFVrfRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFVrfRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFVrfRequested represents a VrfRequested event raised by the VRF contract.
type VRFVrfRequested struct {
	WalletId      [32]byte
	KeyId         uint64
	InstructionId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVrfRequested is a free log retrieval operation binding the contract event 0x52749cdb7f850186051c4aff58b749e22b3e30c23dbd43ac8c8f7a9c5eec5ddd.
//
// Solidity: event VrfRequested(bytes32 indexed walletId, uint64 keyId, bytes32 instructionId)
func (_VRF *VRFFilterer) FilterVrfRequested(opts *bind.FilterOpts, walletId [][32]byte) (*VRFVrfRequestedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _VRF.contract.FilterLogs(opts, "VrfRequested", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFVrfRequestedIterator{contract: _VRF.contract, event: "VrfRequested", logs: logs, sub: sub}, nil
}

// WatchVrfRequested is a free log subscription operation binding the contract event 0x52749cdb7f850186051c4aff58b749e22b3e30c23dbd43ac8c8f7a9c5eec5ddd.
//
// Solidity: event VrfRequested(bytes32 indexed walletId, uint64 keyId, bytes32 instructionId)
func (_VRF *VRFFilterer) WatchVrfRequested(opts *bind.WatchOpts, sink chan<- *VRFVrfRequested, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _VRF.contract.WatchLogs(opts, "VrfRequested", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFVrfRequested)
				if err := _VRF.contract.UnpackLog(event, "VrfRequested", log); err != nil {
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
func (_VRF *VRFFilterer) ParseVrfRequested(log types.Log) (*VRFVrfRequested, error) {
	event := new(VRFVrfRequested)
	if err := _VRF.contract.UnpackLog(event, "VrfRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
