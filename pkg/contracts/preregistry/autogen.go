// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package preregistry

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

// IVoterRegistrySignature is an auto generated low-level Go binding around an user-defined struct.
type IVoterRegistrySignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// PreregistryMetaData contains all meta data concerning the Preregistry contract.
var PreregistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"VoterPreRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"VoterRegistrationFailed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"entityManager\",\"outputs\":[{\"internalType\":\"contractIIEntityManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getPreRegisteredVoters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"isVoterPreRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIVoterRegistry.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"preRegisterVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"triggerVoterRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistry\",\"outputs\":[{\"internalType\":\"contractIIVoterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PreregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PreregistryMetaData.ABI instead.
var PreregistryABI = PreregistryMetaData.ABI

// Preregistry is an auto generated Go binding around an Ethereum contract.
type Preregistry struct {
	PreregistryCaller     // Read-only binding to the contract
	PreregistryTransactor // Write-only binding to the contract
	PreregistryFilterer   // Log filterer for contract events
}

// PreregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PreregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PreregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PreregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PreregistrySession struct {
	Contract     *Preregistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PreregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PreregistryCallerSession struct {
	Contract *PreregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PreregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PreregistryTransactorSession struct {
	Contract     *PreregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PreregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PreregistryRaw struct {
	Contract *Preregistry // Generic contract binding to access the raw methods on
}

// PreregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PreregistryCallerRaw struct {
	Contract *PreregistryCaller // Generic read-only contract binding to access the raw methods on
}

// PreregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PreregistryTransactorRaw struct {
	Contract *PreregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPreregistry creates a new instance of Preregistry, bound to a specific deployed contract.
func NewPreregistry(address common.Address, backend bind.ContractBackend) (*Preregistry, error) {
	contract, err := bindPreregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Preregistry{PreregistryCaller: PreregistryCaller{contract: contract}, PreregistryTransactor: PreregistryTransactor{contract: contract}, PreregistryFilterer: PreregistryFilterer{contract: contract}}, nil
}

// NewPreregistryCaller creates a new read-only instance of Preregistry, bound to a specific deployed contract.
func NewPreregistryCaller(address common.Address, caller bind.ContractCaller) (*PreregistryCaller, error) {
	contract, err := bindPreregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PreregistryCaller{contract: contract}, nil
}

// NewPreregistryTransactor creates a new write-only instance of Preregistry, bound to a specific deployed contract.
func NewPreregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PreregistryTransactor, error) {
	contract, err := bindPreregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PreregistryTransactor{contract: contract}, nil
}

// NewPreregistryFilterer creates a new log filterer instance of Preregistry, bound to a specific deployed contract.
func NewPreregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PreregistryFilterer, error) {
	contract, err := bindPreregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PreregistryFilterer{contract: contract}, nil
}

// bindPreregistry binds a generic wrapper to an already deployed contract.
func bindPreregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PreregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Preregistry *PreregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Preregistry.Contract.PreregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Preregistry *PreregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Preregistry.Contract.PreregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Preregistry *PreregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Preregistry.Contract.PreregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Preregistry *PreregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Preregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Preregistry *PreregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Preregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Preregistry *PreregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Preregistry.Contract.contract.Transact(opts, method, params...)
}

// EntityManager is a free data retrieval call binding the contract method 0x50b1d61b.
//
// Solidity: function entityManager() view returns(address)
func (_Preregistry *PreregistryCaller) EntityManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Preregistry.contract.Call(opts, &out, "entityManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntityManager is a free data retrieval call binding the contract method 0x50b1d61b.
//
// Solidity: function entityManager() view returns(address)
func (_Preregistry *PreregistrySession) EntityManager() (common.Address, error) {
	return _Preregistry.Contract.EntityManager(&_Preregistry.CallOpts)
}

// EntityManager is a free data retrieval call binding the contract method 0x50b1d61b.
//
// Solidity: function entityManager() view returns(address)
func (_Preregistry *PreregistryCallerSession) EntityManager() (common.Address, error) {
	return _Preregistry.Contract.EntityManager(&_Preregistry.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_Preregistry *PreregistryCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Preregistry.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_Preregistry *PreregistrySession) FlareSystemsManager() (common.Address, error) {
	return _Preregistry.Contract.FlareSystemsManager(&_Preregistry.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_Preregistry *PreregistryCallerSession) FlareSystemsManager() (common.Address, error) {
	return _Preregistry.Contract.FlareSystemsManager(&_Preregistry.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Preregistry *PreregistryCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Preregistry.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Preregistry *PreregistrySession) GetAddressUpdater() (common.Address, error) {
	return _Preregistry.Contract.GetAddressUpdater(&_Preregistry.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Preregistry *PreregistryCallerSession) GetAddressUpdater() (common.Address, error) {
	return _Preregistry.Contract.GetAddressUpdater(&_Preregistry.CallOpts)
}

// GetPreRegisteredVoters is a free data retrieval call binding the contract method 0xf116d1dd.
//
// Solidity: function getPreRegisteredVoters(uint24 _rewardEpochId) view returns(address[])
func (_Preregistry *PreregistryCaller) GetPreRegisteredVoters(opts *bind.CallOpts, _rewardEpochId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Preregistry.contract.Call(opts, &out, "getPreRegisteredVoters", _rewardEpochId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPreRegisteredVoters is a free data retrieval call binding the contract method 0xf116d1dd.
//
// Solidity: function getPreRegisteredVoters(uint24 _rewardEpochId) view returns(address[])
func (_Preregistry *PreregistrySession) GetPreRegisteredVoters(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _Preregistry.Contract.GetPreRegisteredVoters(&_Preregistry.CallOpts, _rewardEpochId)
}

// GetPreRegisteredVoters is a free data retrieval call binding the contract method 0xf116d1dd.
//
// Solidity: function getPreRegisteredVoters(uint24 _rewardEpochId) view returns(address[])
func (_Preregistry *PreregistryCallerSession) GetPreRegisteredVoters(_rewardEpochId *big.Int) ([]common.Address, error) {
	return _Preregistry.Contract.GetPreRegisteredVoters(&_Preregistry.CallOpts, _rewardEpochId)
}

// IsVoterPreRegistered is a free data retrieval call binding the contract method 0x19549e92.
//
// Solidity: function isVoterPreRegistered(uint24 _rewardEpochId, address _voter) view returns(bool)
func (_Preregistry *PreregistryCaller) IsVoterPreRegistered(opts *bind.CallOpts, _rewardEpochId *big.Int, _voter common.Address) (bool, error) {
	var out []interface{}
	err := _Preregistry.contract.Call(opts, &out, "isVoterPreRegistered", _rewardEpochId, _voter)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoterPreRegistered is a free data retrieval call binding the contract method 0x19549e92.
//
// Solidity: function isVoterPreRegistered(uint24 _rewardEpochId, address _voter) view returns(bool)
func (_Preregistry *PreregistrySession) IsVoterPreRegistered(_rewardEpochId *big.Int, _voter common.Address) (bool, error) {
	return _Preregistry.Contract.IsVoterPreRegistered(&_Preregistry.CallOpts, _rewardEpochId, _voter)
}

// IsVoterPreRegistered is a free data retrieval call binding the contract method 0x19549e92.
//
// Solidity: function isVoterPreRegistered(uint24 _rewardEpochId, address _voter) view returns(bool)
func (_Preregistry *PreregistryCallerSession) IsVoterPreRegistered(_rewardEpochId *big.Int, _voter common.Address) (bool, error) {
	return _Preregistry.Contract.IsVoterPreRegistered(&_Preregistry.CallOpts, _rewardEpochId, _voter)
}

// VoterRegistry is a free data retrieval call binding the contract method 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (_Preregistry *PreregistryCaller) VoterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Preregistry.contract.Call(opts, &out, "voterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoterRegistry is a free data retrieval call binding the contract method 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (_Preregistry *PreregistrySession) VoterRegistry() (common.Address, error) {
	return _Preregistry.Contract.VoterRegistry(&_Preregistry.CallOpts)
}

// VoterRegistry is a free data retrieval call binding the contract method 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (_Preregistry *PreregistryCallerSession) VoterRegistry() (common.Address, error) {
	return _Preregistry.Contract.VoterRegistry(&_Preregistry.CallOpts)
}

// PreRegisterVoter is a paid mutator transaction binding the contract method 0xea8eb77d.
//
// Solidity: function preRegisterVoter(address _voter, (uint8,bytes32,bytes32) _signature) returns()
func (_Preregistry *PreregistryTransactor) PreRegisterVoter(opts *bind.TransactOpts, _voter common.Address, _signature IVoterRegistrySignature) (*types.Transaction, error) {
	return _Preregistry.contract.Transact(opts, "preRegisterVoter", _voter, _signature)
}

// PreRegisterVoter is a paid mutator transaction binding the contract method 0xea8eb77d.
//
// Solidity: function preRegisterVoter(address _voter, (uint8,bytes32,bytes32) _signature) returns()
func (_Preregistry *PreregistrySession) PreRegisterVoter(_voter common.Address, _signature IVoterRegistrySignature) (*types.Transaction, error) {
	return _Preregistry.Contract.PreRegisterVoter(&_Preregistry.TransactOpts, _voter, _signature)
}

// PreRegisterVoter is a paid mutator transaction binding the contract method 0xea8eb77d.
//
// Solidity: function preRegisterVoter(address _voter, (uint8,bytes32,bytes32) _signature) returns()
func (_Preregistry *PreregistryTransactorSession) PreRegisterVoter(_voter common.Address, _signature IVoterRegistrySignature) (*types.Transaction, error) {
	return _Preregistry.Contract.PreRegisterVoter(&_Preregistry.TransactOpts, _voter, _signature)
}

// TriggerVoterRegistration is a paid mutator transaction binding the contract method 0xf707f612.
//
// Solidity: function triggerVoterRegistration(uint24 _rewardEpochId) returns()
func (_Preregistry *PreregistryTransactor) TriggerVoterRegistration(opts *bind.TransactOpts, _rewardEpochId *big.Int) (*types.Transaction, error) {
	return _Preregistry.contract.Transact(opts, "triggerVoterRegistration", _rewardEpochId)
}

// TriggerVoterRegistration is a paid mutator transaction binding the contract method 0xf707f612.
//
// Solidity: function triggerVoterRegistration(uint24 _rewardEpochId) returns()
func (_Preregistry *PreregistrySession) TriggerVoterRegistration(_rewardEpochId *big.Int) (*types.Transaction, error) {
	return _Preregistry.Contract.TriggerVoterRegistration(&_Preregistry.TransactOpts, _rewardEpochId)
}

// TriggerVoterRegistration is a paid mutator transaction binding the contract method 0xf707f612.
//
// Solidity: function triggerVoterRegistration(uint24 _rewardEpochId) returns()
func (_Preregistry *PreregistryTransactorSession) TriggerVoterRegistration(_rewardEpochId *big.Int) (*types.Transaction, error) {
	return _Preregistry.Contract.TriggerVoterRegistration(&_Preregistry.TransactOpts, _rewardEpochId)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Preregistry *PreregistryTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Preregistry.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Preregistry *PreregistrySession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Preregistry.Contract.UpdateContractAddresses(&_Preregistry.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Preregistry *PreregistryTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Preregistry.Contract.UpdateContractAddresses(&_Preregistry.TransactOpts, _contractNameHashes, _contractAddresses)
}

// PreregistryVoterPreRegisteredIterator is returned from FilterVoterPreRegistered and is used to iterate over the raw logs and unpacked data for VoterPreRegistered events raised by the Preregistry contract.
type PreregistryVoterPreRegisteredIterator struct {
	Event *PreregistryVoterPreRegistered // Event containing the contract specifics and raw log

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
func (it *PreregistryVoterPreRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreregistryVoterPreRegistered)
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
		it.Event = new(PreregistryVoterPreRegistered)
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
func (it *PreregistryVoterPreRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreregistryVoterPreRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreregistryVoterPreRegistered represents a VoterPreRegistered event raised by the Preregistry contract.
type PreregistryVoterPreRegistered struct {
	Voter         common.Address
	RewardEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVoterPreRegistered is a free log retrieval operation binding the contract event 0x33dfb7ce87c082c21fdb6b5833ba631717f7fd8d8bf8311426fe740ecd2e71fd.
//
// Solidity: event VoterPreRegistered(address indexed voter, uint256 indexed rewardEpochId)
func (_Preregistry *PreregistryFilterer) FilterVoterPreRegistered(opts *bind.FilterOpts, voter []common.Address, rewardEpochId []*big.Int) (*PreregistryVoterPreRegisteredIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _Preregistry.contract.FilterLogs(opts, "VoterPreRegistered", voterRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &PreregistryVoterPreRegisteredIterator{contract: _Preregistry.contract, event: "VoterPreRegistered", logs: logs, sub: sub}, nil
}

// WatchVoterPreRegistered is a free log subscription operation binding the contract event 0x33dfb7ce87c082c21fdb6b5833ba631717f7fd8d8bf8311426fe740ecd2e71fd.
//
// Solidity: event VoterPreRegistered(address indexed voter, uint256 indexed rewardEpochId)
func (_Preregistry *PreregistryFilterer) WatchVoterPreRegistered(opts *bind.WatchOpts, sink chan<- *PreregistryVoterPreRegistered, voter []common.Address, rewardEpochId []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _Preregistry.contract.WatchLogs(opts, "VoterPreRegistered", voterRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreregistryVoterPreRegistered)
				if err := _Preregistry.contract.UnpackLog(event, "VoterPreRegistered", log); err != nil {
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

// ParseVoterPreRegistered is a log parse operation binding the contract event 0x33dfb7ce87c082c21fdb6b5833ba631717f7fd8d8bf8311426fe740ecd2e71fd.
//
// Solidity: event VoterPreRegistered(address indexed voter, uint256 indexed rewardEpochId)
func (_Preregistry *PreregistryFilterer) ParseVoterPreRegistered(log types.Log) (*PreregistryVoterPreRegistered, error) {
	event := new(PreregistryVoterPreRegistered)
	if err := _Preregistry.contract.UnpackLog(event, "VoterPreRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PreregistryVoterRegistrationFailedIterator is returned from FilterVoterRegistrationFailed and is used to iterate over the raw logs and unpacked data for VoterRegistrationFailed events raised by the Preregistry contract.
type PreregistryVoterRegistrationFailedIterator struct {
	Event *PreregistryVoterRegistrationFailed // Event containing the contract specifics and raw log

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
func (it *PreregistryVoterRegistrationFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreregistryVoterRegistrationFailed)
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
		it.Event = new(PreregistryVoterRegistrationFailed)
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
func (it *PreregistryVoterRegistrationFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreregistryVoterRegistrationFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreregistryVoterRegistrationFailed represents a VoterRegistrationFailed event raised by the Preregistry contract.
type PreregistryVoterRegistrationFailed struct {
	Voter         common.Address
	RewardEpochId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterVoterRegistrationFailed is a free log retrieval operation binding the contract event 0x925794c2e759216f395d05fc6e6c1ff4632b87077831f9a8fe78d448203ccddc.
//
// Solidity: event VoterRegistrationFailed(address indexed voter, uint256 indexed rewardEpochId)
func (_Preregistry *PreregistryFilterer) FilterVoterRegistrationFailed(opts *bind.FilterOpts, voter []common.Address, rewardEpochId []*big.Int) (*PreregistryVoterRegistrationFailedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _Preregistry.contract.FilterLogs(opts, "VoterRegistrationFailed", voterRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &PreregistryVoterRegistrationFailedIterator{contract: _Preregistry.contract, event: "VoterRegistrationFailed", logs: logs, sub: sub}, nil
}

// WatchVoterRegistrationFailed is a free log subscription operation binding the contract event 0x925794c2e759216f395d05fc6e6c1ff4632b87077831f9a8fe78d448203ccddc.
//
// Solidity: event VoterRegistrationFailed(address indexed voter, uint256 indexed rewardEpochId)
func (_Preregistry *PreregistryFilterer) WatchVoterRegistrationFailed(opts *bind.WatchOpts, sink chan<- *PreregistryVoterRegistrationFailed, voter []common.Address, rewardEpochId []*big.Int) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _Preregistry.contract.WatchLogs(opts, "VoterRegistrationFailed", voterRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreregistryVoterRegistrationFailed)
				if err := _Preregistry.contract.UnpackLog(event, "VoterRegistrationFailed", log); err != nil {
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

// ParseVoterRegistrationFailed is a log parse operation binding the contract event 0x925794c2e759216f395d05fc6e6c1ff4632b87077831f9a8fe78d448203ccddc.
//
// Solidity: event VoterRegistrationFailed(address indexed voter, uint256 indexed rewardEpochId)
func (_Preregistry *PreregistryFilterer) ParseVoterRegistrationFailed(log types.Log) (*PreregistryVoterRegistrationFailed, error) {
	event := new(PreregistryVoterRegistrationFailed)
	if err := _Preregistry.contract.UnpackLog(event, "VoterRegistrationFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
