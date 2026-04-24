// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletbackupmanager

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

// IWalletBackupManagerBackupId is an auto generated low-level Go binding around an user-defined struct.
type IWalletBackupManagerBackupId struct {
	TeeId         common.Address
	WalletId      [32]byte
	KeyId         uint64
	KeyType       [32]byte
	SigningAlgo   [32]byte
	PublicKey     []byte
	RewardEpochId uint32
	RandomNonce   [32]byte
}

// WalletBackupManagerMetaData contains all meta data concerning the WalletBackupManager contract.
var WalletBackupManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardEpochId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeMachine\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyAlreadyAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyNotConfirmed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedRewardEpochId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"BackupRestoreTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"randomNonce\",\"type\":\"bytes32\"}],\"internalType\":\"structIWalletBackupManager.BackupId\",\"name\":\"_backupId\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_backupUrl\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"backupRestore\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// WalletBackupManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletBackupManagerMetaData.ABI instead.
var WalletBackupManagerABI = WalletBackupManagerMetaData.ABI

// WalletBackupManager is an auto generated Go binding around an Ethereum contract.
type WalletBackupManager struct {
	WalletBackupManagerCaller     // Read-only binding to the contract
	WalletBackupManagerTransactor // Write-only binding to the contract
	WalletBackupManagerFilterer   // Log filterer for contract events
}

// WalletBackupManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletBackupManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletBackupManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletBackupManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletBackupManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletBackupManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletBackupManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletBackupManagerSession struct {
	Contract     *WalletBackupManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WalletBackupManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletBackupManagerCallerSession struct {
	Contract *WalletBackupManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// WalletBackupManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletBackupManagerTransactorSession struct {
	Contract     *WalletBackupManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// WalletBackupManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletBackupManagerRaw struct {
	Contract *WalletBackupManager // Generic contract binding to access the raw methods on
}

// WalletBackupManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletBackupManagerCallerRaw struct {
	Contract *WalletBackupManagerCaller // Generic read-only contract binding to access the raw methods on
}

// WalletBackupManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletBackupManagerTransactorRaw struct {
	Contract *WalletBackupManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletBackupManager creates a new instance of WalletBackupManager, bound to a specific deployed contract.
func NewWalletBackupManager(address common.Address, backend bind.ContractBackend) (*WalletBackupManager, error) {
	contract, err := bindWalletBackupManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletBackupManager{WalletBackupManagerCaller: WalletBackupManagerCaller{contract: contract}, WalletBackupManagerTransactor: WalletBackupManagerTransactor{contract: contract}, WalletBackupManagerFilterer: WalletBackupManagerFilterer{contract: contract}}, nil
}

// NewWalletBackupManagerCaller creates a new read-only instance of WalletBackupManager, bound to a specific deployed contract.
func NewWalletBackupManagerCaller(address common.Address, caller bind.ContractCaller) (*WalletBackupManagerCaller, error) {
	contract, err := bindWalletBackupManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletBackupManagerCaller{contract: contract}, nil
}

// NewWalletBackupManagerTransactor creates a new write-only instance of WalletBackupManager, bound to a specific deployed contract.
func NewWalletBackupManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletBackupManagerTransactor, error) {
	contract, err := bindWalletBackupManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletBackupManagerTransactor{contract: contract}, nil
}

// NewWalletBackupManagerFilterer creates a new log filterer instance of WalletBackupManager, bound to a specific deployed contract.
func NewWalletBackupManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletBackupManagerFilterer, error) {
	contract, err := bindWalletBackupManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletBackupManagerFilterer{contract: contract}, nil
}

// bindWalletBackupManager binds a generic wrapper to an already deployed contract.
func bindWalletBackupManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletBackupManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletBackupManager *WalletBackupManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletBackupManager.Contract.WalletBackupManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletBackupManager *WalletBackupManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletBackupManager.Contract.WalletBackupManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletBackupManager *WalletBackupManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletBackupManager.Contract.WalletBackupManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletBackupManager *WalletBackupManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletBackupManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletBackupManager *WalletBackupManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletBackupManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletBackupManager *WalletBackupManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletBackupManager.Contract.contract.Transact(opts, method, params...)
}

// BackupRestore is a paid mutator transaction binding the contract method 0x20737dd7.
//
// Solidity: function backupRestore(address _teeId, (address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32) _backupId, string _backupUrl, address _claimBackAddress) payable returns()
func (_WalletBackupManager *WalletBackupManagerTransactor) BackupRestore(opts *bind.TransactOpts, _teeId common.Address, _backupId IWalletBackupManagerBackupId, _backupUrl string, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletBackupManager.contract.Transact(opts, "backupRestore", _teeId, _backupId, _backupUrl, _claimBackAddress)
}

// BackupRestore is a paid mutator transaction binding the contract method 0x20737dd7.
//
// Solidity: function backupRestore(address _teeId, (address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32) _backupId, string _backupUrl, address _claimBackAddress) payable returns()
func (_WalletBackupManager *WalletBackupManagerSession) BackupRestore(_teeId common.Address, _backupId IWalletBackupManagerBackupId, _backupUrl string, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletBackupManager.Contract.BackupRestore(&_WalletBackupManager.TransactOpts, _teeId, _backupId, _backupUrl, _claimBackAddress)
}

// BackupRestore is a paid mutator transaction binding the contract method 0x20737dd7.
//
// Solidity: function backupRestore(address _teeId, (address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32) _backupId, string _backupUrl, address _claimBackAddress) payable returns()
func (_WalletBackupManager *WalletBackupManagerTransactorSession) BackupRestore(_teeId common.Address, _backupId IWalletBackupManagerBackupId, _backupUrl string, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletBackupManager.Contract.BackupRestore(&_WalletBackupManager.TransactOpts, _teeId, _backupId, _backupUrl, _claimBackAddress)
}

// WalletBackupManagerBackupRestoreTriggeredIterator is returned from FilterBackupRestoreTriggered and is used to iterate over the raw logs and unpacked data for BackupRestoreTriggered events raised by the WalletBackupManager contract.
type WalletBackupManagerBackupRestoreTriggeredIterator struct {
	Event *WalletBackupManagerBackupRestoreTriggered // Event containing the contract specifics and raw log

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
func (it *WalletBackupManagerBackupRestoreTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletBackupManagerBackupRestoreTriggered)
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
		it.Event = new(WalletBackupManagerBackupRestoreTriggered)
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
func (it *WalletBackupManagerBackupRestoreTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletBackupManagerBackupRestoreTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletBackupManagerBackupRestoreTriggered represents a BackupRestoreTriggered event raised by the WalletBackupManager contract.
type WalletBackupManagerBackupRestoreTriggered struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    uint64
	Nonce    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBackupRestoreTriggered is a free log retrieval operation binding the contract event 0x0696822220fe833afbe2d89fd4f60ff50190fbda02ff94256d2301ca4097f812.
//
// Solidity: event BackupRestoreTriggered(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId, uint256 nonce)
func (_WalletBackupManager *WalletBackupManagerFilterer) FilterBackupRestoreTriggered(opts *bind.FilterOpts, teeId []common.Address, walletId [][32]byte, keyId []uint64) (*WalletBackupManagerBackupRestoreTriggeredIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _WalletBackupManager.contract.FilterLogs(opts, "BackupRestoreTriggered", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletBackupManagerBackupRestoreTriggeredIterator{contract: _WalletBackupManager.contract, event: "BackupRestoreTriggered", logs: logs, sub: sub}, nil
}

// WatchBackupRestoreTriggered is a free log subscription operation binding the contract event 0x0696822220fe833afbe2d89fd4f60ff50190fbda02ff94256d2301ca4097f812.
//
// Solidity: event BackupRestoreTriggered(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId, uint256 nonce)
func (_WalletBackupManager *WalletBackupManagerFilterer) WatchBackupRestoreTriggered(opts *bind.WatchOpts, sink chan<- *WalletBackupManagerBackupRestoreTriggered, teeId []common.Address, walletId [][32]byte, keyId []uint64) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _WalletBackupManager.contract.WatchLogs(opts, "BackupRestoreTriggered", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletBackupManagerBackupRestoreTriggered)
				if err := _WalletBackupManager.contract.UnpackLog(event, "BackupRestoreTriggered", log); err != nil {
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

// ParseBackupRestoreTriggered is a log parse operation binding the contract event 0x0696822220fe833afbe2d89fd4f60ff50190fbda02ff94256d2301ca4097f812.
//
// Solidity: event BackupRestoreTriggered(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId, uint256 nonce)
func (_WalletBackupManager *WalletBackupManagerFilterer) ParseBackupRestoreTriggered(log types.Log) (*WalletBackupManagerBackupRestoreTriggered, error) {
	event := new(WalletBackupManagerBackupRestoreTriggered)
	if err := _WalletBackupManager.contract.UnpackLog(event, "BackupRestoreTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletBackupManagerTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the WalletBackupManager contract.
type WalletBackupManagerTeeInstructionsSentIterator struct {
	Event *WalletBackupManagerTeeInstructionsSent // Event containing the contract specifics and raw log

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
func (it *WalletBackupManagerTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletBackupManagerTeeInstructionsSent)
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
		it.Event = new(WalletBackupManagerTeeInstructionsSent)
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
func (it *WalletBackupManagerTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletBackupManagerTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletBackupManagerTeeInstructionsSent represents a TeeInstructionsSent event raised by the WalletBackupManager contract.
type WalletBackupManagerTeeInstructionsSent struct {
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
func (_WalletBackupManager *WalletBackupManagerFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (*WalletBackupManagerTeeInstructionsSentIterator, error) {

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

	logs, sub, err := _WalletBackupManager.contract.FilterLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletBackupManagerTeeInstructionsSentIterator{contract: _WalletBackupManager.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_WalletBackupManager *WalletBackupManagerFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *WalletBackupManagerTeeInstructionsSent, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

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

	logs, sub, err := _WalletBackupManager.contract.WatchLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletBackupManagerTeeInstructionsSent)
				if err := _WalletBackupManager.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
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
func (_WalletBackupManager *WalletBackupManagerFilterer) ParseTeeInstructionsSent(log types.Log) (*WalletBackupManagerTeeInstructionsSent, error) {
	event := new(WalletBackupManagerTeeInstructionsSent)
	if err := _WalletBackupManager.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
