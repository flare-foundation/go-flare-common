// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletresume

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

// IWalletResumeResumeKeyData is an auto generated low-level Go binding around an user-defined struct.
type IWalletResumeResumeKeyData struct {
	KeyId uint64
	TeeId common.Address
	Nonce *big.Int
}

// WalletResumeMetaData contains all meta data concerning the WalletResume contract.
var WalletResumeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"EmergencyPauseActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongKeyId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"keyIds\",\"type\":\"uint64[]\"}],\"name\":\"WalletKeysNotAvailable\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIWalletResume.ResumeKeyData[]\",\"name\":\"_keysData\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_pausingAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"setPausingAddresses\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// WalletResumeABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletResumeMetaData.ABI instead.
var WalletResumeABI = WalletResumeMetaData.ABI

// WalletResume is an auto generated Go binding around an Ethereum contract.
type WalletResume struct {
	WalletResumeCaller     // Read-only binding to the contract
	WalletResumeTransactor // Write-only binding to the contract
	WalletResumeFilterer   // Log filterer for contract events
}

// WalletResumeCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletResumeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletResumeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletResumeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletResumeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletResumeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletResumeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletResumeSession struct {
	Contract     *WalletResume     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletResumeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletResumeCallerSession struct {
	Contract *WalletResumeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WalletResumeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletResumeTransactorSession struct {
	Contract     *WalletResumeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WalletResumeRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletResumeRaw struct {
	Contract *WalletResume // Generic contract binding to access the raw methods on
}

// WalletResumeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletResumeCallerRaw struct {
	Contract *WalletResumeCaller // Generic read-only contract binding to access the raw methods on
}

// WalletResumeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletResumeTransactorRaw struct {
	Contract *WalletResumeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletResume creates a new instance of WalletResume, bound to a specific deployed contract.
func NewWalletResume(address common.Address, backend bind.ContractBackend) (*WalletResume, error) {
	contract, err := bindWalletResume(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletResume{WalletResumeCaller: WalletResumeCaller{contract: contract}, WalletResumeTransactor: WalletResumeTransactor{contract: contract}, WalletResumeFilterer: WalletResumeFilterer{contract: contract}}, nil
}

// NewWalletResumeCaller creates a new read-only instance of WalletResume, bound to a specific deployed contract.
func NewWalletResumeCaller(address common.Address, caller bind.ContractCaller) (*WalletResumeCaller, error) {
	contract, err := bindWalletResume(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletResumeCaller{contract: contract}, nil
}

// NewWalletResumeTransactor creates a new write-only instance of WalletResume, bound to a specific deployed contract.
func NewWalletResumeTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletResumeTransactor, error) {
	contract, err := bindWalletResume(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletResumeTransactor{contract: contract}, nil
}

// NewWalletResumeFilterer creates a new log filterer instance of WalletResume, bound to a specific deployed contract.
func NewWalletResumeFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletResumeFilterer, error) {
	contract, err := bindWalletResume(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletResumeFilterer{contract: contract}, nil
}

// bindWalletResume binds a generic wrapper to an already deployed contract.
func bindWalletResume(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletResumeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletResume *WalletResumeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletResume.Contract.WalletResumeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletResume *WalletResumeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletResume.Contract.WalletResumeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletResume *WalletResumeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletResume.Contract.WalletResumeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletResume *WalletResumeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletResume.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletResume *WalletResumeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletResume.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletResume *WalletResumeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletResume.Contract.contract.Transact(opts, method, params...)
}

// Resume is a paid mutator transaction binding the contract method 0xa353fe2a.
//
// Solidity: function resume(bytes32 _walletId, (uint64,address,uint256)[] _keysData, address _claimBackAddress) payable returns()
func (_WalletResume *WalletResumeTransactor) Resume(opts *bind.TransactOpts, _walletId [32]byte, _keysData []IWalletResumeResumeKeyData, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletResume.contract.Transact(opts, "resume", _walletId, _keysData, _claimBackAddress)
}

// Resume is a paid mutator transaction binding the contract method 0xa353fe2a.
//
// Solidity: function resume(bytes32 _walletId, (uint64,address,uint256)[] _keysData, address _claimBackAddress) payable returns()
func (_WalletResume *WalletResumeSession) Resume(_walletId [32]byte, _keysData []IWalletResumeResumeKeyData, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletResume.Contract.Resume(&_WalletResume.TransactOpts, _walletId, _keysData, _claimBackAddress)
}

// Resume is a paid mutator transaction binding the contract method 0xa353fe2a.
//
// Solidity: function resume(bytes32 _walletId, (uint64,address,uint256)[] _keysData, address _claimBackAddress) payable returns()
func (_WalletResume *WalletResumeTransactorSession) Resume(_walletId [32]byte, _keysData []IWalletResumeResumeKeyData, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletResume.Contract.Resume(&_WalletResume.TransactOpts, _walletId, _keysData, _claimBackAddress)
}

// SetPausingAddresses is a paid mutator transaction binding the contract method 0xac7355e0.
//
// Solidity: function setPausingAddresses(bytes32 _walletId, address[] _pausingAddresses, address _claimBackAddress) payable returns()
func (_WalletResume *WalletResumeTransactor) SetPausingAddresses(opts *bind.TransactOpts, _walletId [32]byte, _pausingAddresses []common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletResume.contract.Transact(opts, "setPausingAddresses", _walletId, _pausingAddresses, _claimBackAddress)
}

// SetPausingAddresses is a paid mutator transaction binding the contract method 0xac7355e0.
//
// Solidity: function setPausingAddresses(bytes32 _walletId, address[] _pausingAddresses, address _claimBackAddress) payable returns()
func (_WalletResume *WalletResumeSession) SetPausingAddresses(_walletId [32]byte, _pausingAddresses []common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletResume.Contract.SetPausingAddresses(&_WalletResume.TransactOpts, _walletId, _pausingAddresses, _claimBackAddress)
}

// SetPausingAddresses is a paid mutator transaction binding the contract method 0xac7355e0.
//
// Solidity: function setPausingAddresses(bytes32 _walletId, address[] _pausingAddresses, address _claimBackAddress) payable returns()
func (_WalletResume *WalletResumeTransactorSession) SetPausingAddresses(_walletId [32]byte, _pausingAddresses []common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletResume.Contract.SetPausingAddresses(&_WalletResume.TransactOpts, _walletId, _pausingAddresses, _claimBackAddress)
}

// WalletResumeTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the WalletResume contract.
type WalletResumeTeeInstructionsSentIterator struct {
	Event *WalletResumeTeeInstructionsSent // Event containing the contract specifics and raw log

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
func (it *WalletResumeTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletResumeTeeInstructionsSent)
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
		it.Event = new(WalletResumeTeeInstructionsSent)
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
func (it *WalletResumeTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletResumeTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletResumeTeeInstructionsSent represents a TeeInstructionsSent event raised by the WalletResume contract.
type WalletResumeTeeInstructionsSent struct {
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
func (_WalletResume *WalletResumeFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (*WalletResumeTeeInstructionsSentIterator, error) {

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

	logs, sub, err := _WalletResume.contract.FilterLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletResumeTeeInstructionsSentIterator{contract: _WalletResume.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_WalletResume *WalletResumeFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *WalletResumeTeeInstructionsSent, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

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

	logs, sub, err := _WalletResume.contract.WatchLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletResumeTeeInstructionsSent)
				if err := _WalletResume.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
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
func (_WalletResume *WalletResumeFilterer) ParseTeeInstructionsSent(log types.Log) (*WalletResumeTeeInstructionsSent, error) {
	event := new(WalletResumeTeeInstructionsSent)
	if err := _WalletResume.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletResumeWalletKeysNotAvailableIterator is returned from FilterWalletKeysNotAvailable and is used to iterate over the raw logs and unpacked data for WalletKeysNotAvailable events raised by the WalletResume contract.
type WalletResumeWalletKeysNotAvailableIterator struct {
	Event *WalletResumeWalletKeysNotAvailable // Event containing the contract specifics and raw log

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
func (it *WalletResumeWalletKeysNotAvailableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletResumeWalletKeysNotAvailable)
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
		it.Event = new(WalletResumeWalletKeysNotAvailable)
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
func (it *WalletResumeWalletKeysNotAvailableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletResumeWalletKeysNotAvailableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletResumeWalletKeysNotAvailable represents a WalletKeysNotAvailable event raised by the WalletResume contract.
type WalletResumeWalletKeysNotAvailable struct {
	WalletId [32]byte
	KeyIds   []uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletKeysNotAvailable is a free log retrieval operation binding the contract event 0xa82e9445a32f966f6330fc8dc7dbb1a5abf41ce17c2695f248a967d11dc8199c.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint64[] keyIds)
func (_WalletResume *WalletResumeFilterer) FilterWalletKeysNotAvailable(opts *bind.FilterOpts, walletId [][32]byte) (*WalletResumeWalletKeysNotAvailableIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletResume.contract.FilterLogs(opts, "WalletKeysNotAvailable", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletResumeWalletKeysNotAvailableIterator{contract: _WalletResume.contract, event: "WalletKeysNotAvailable", logs: logs, sub: sub}, nil
}

// WatchWalletKeysNotAvailable is a free log subscription operation binding the contract event 0xa82e9445a32f966f6330fc8dc7dbb1a5abf41ce17c2695f248a967d11dc8199c.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint64[] keyIds)
func (_WalletResume *WalletResumeFilterer) WatchWalletKeysNotAvailable(opts *bind.WatchOpts, sink chan<- *WalletResumeWalletKeysNotAvailable, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletResume.contract.WatchLogs(opts, "WalletKeysNotAvailable", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletResumeWalletKeysNotAvailable)
				if err := _WalletResume.contract.UnpackLog(event, "WalletKeysNotAvailable", log); err != nil {
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

// ParseWalletKeysNotAvailable is a log parse operation binding the contract event 0xa82e9445a32f966f6330fc8dc7dbb1a5abf41ce17c2695f248a967d11dc8199c.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint64[] keyIds)
func (_WalletResume *WalletResumeFilterer) ParseWalletKeysNotAvailable(log types.Log) (*WalletResumeWalletKeysNotAvailable, error) {
	event := new(WalletResumeWalletKeysNotAvailable)
	if err := _WalletResume.contract.UnpackLog(event, "WalletKeysNotAvailable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
