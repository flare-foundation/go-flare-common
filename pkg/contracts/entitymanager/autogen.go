// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package entitymanager

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

// IEntityManagerVoterAddresses is an auto generated low-level Go binding around an user-defined struct.
type IEntityManagerVoterAddresses struct {
	SubmitAddress           common.Address
	SubmitSignaturesAddress common.Address
	SigningPolicyAddress    common.Address
}

// EntityManagerMetaData contains all meta data concerning the EntityManager contract.
var EntityManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegationAddress\",\"type\":\"address\"}],\"name\":\"DelegationAddressProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegationAddress\",\"type\":\"address\"}],\"name\":\"DelegationAddressRegistrationConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxNodeIdsPerEntity\",\"type\":\"uint256\"}],\"name\":\"MaxNodeIdsPerEntitySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"nodeId\",\"type\":\"bytes20\"}],\"name\":\"NodeIdRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"nodeId\",\"type\":\"bytes20\"}],\"name\":\"NodeIdUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"part1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"part2\",\"type\":\"bytes32\"}],\"name\":\"PublicKeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"part1\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"part2\",\"type\":\"bytes32\"}],\"name\":\"PublicKeyUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"}],\"name\":\"SigningPolicyAddressProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"}],\"name\":\"SigningPolicyAddressRegistrationConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitAddress\",\"type\":\"address\"}],\"name\":\"SubmitAddressProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitAddress\",\"type\":\"address\"}],\"name\":\"SubmitAddressRegistrationConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitSignaturesAddress\",\"type\":\"address\"}],\"name\":\"SubmitSignaturesAddressProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitSignaturesAddress\",\"type\":\"address\"}],\"name\":\"SubmitSignaturesAddressRegistrationConfirmed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"confirmDelegationAddressRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"confirmSigningPolicyAddressRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"confirmSubmitAddressRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"confirmSubmitSignaturesAddressRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getDelegationAddressOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getDelegationAddressOfAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getNodeIdsOf\",\"outputs\":[{\"internalType\":\"bytes20[]\",\"name\":\"\",\"type\":\"bytes20[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getNodeIdsOfAt\",\"outputs\":[{\"internalType\":\"bytes20[]\",\"name\":\"\",\"type\":\"bytes20[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getPublicKeyOf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPublicKeyOfAt\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterAddresses\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"submitSignaturesAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"}],\"internalType\":\"structIEntityManager.VoterAddresses\",\"name\":\"_addresses\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterAddressesAt\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"submitAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"submitSignaturesAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"}],\"internalType\":\"structIEntityManager.VoterAddresses\",\"name\":\"_addresses\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterForDelegationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_nodeId\",\"type\":\"bytes20\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterForNodeId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_part1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_part2\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterForPublicKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signingPolicyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterForSigningPolicyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_submitAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterForSubmitAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_submitSignaturesAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getVoterForSubmitSignaturesAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegationAddress\",\"type\":\"address\"}],\"name\":\"proposeDelegationAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signingPolicyAddress\",\"type\":\"address\"}],\"name\":\"proposeSigningPolicyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_submitAddress\",\"type\":\"address\"}],\"name\":\"proposeSubmitAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_submitSignaturesAddress\",\"type\":\"address\"}],\"name\":\"proposeSubmitSignaturesAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_nodeId\",\"type\":\"bytes20\"},{\"internalType\":\"bytes\",\"name\":\"_certificateRaw\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"registerNodeId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_part1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_part2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_verificationData\",\"type\":\"bytes\"}],\"name\":\"registerPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_nodeId\",\"type\":\"bytes20\"}],\"name\":\"unregisterNodeId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregisterPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EntityManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use EntityManagerMetaData.ABI instead.
var EntityManagerABI = EntityManagerMetaData.ABI

// EntityManager is an auto generated Go binding around an Ethereum contract.
type EntityManager struct {
	EntityManagerCaller     // Read-only binding to the contract
	EntityManagerTransactor // Write-only binding to the contract
	EntityManagerFilterer   // Log filterer for contract events
}

// EntityManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type EntityManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntityManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EntityManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntityManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EntityManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntityManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EntityManagerSession struct {
	Contract     *EntityManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EntityManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EntityManagerCallerSession struct {
	Contract *EntityManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EntityManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EntityManagerTransactorSession struct {
	Contract     *EntityManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EntityManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type EntityManagerRaw struct {
	Contract *EntityManager // Generic contract binding to access the raw methods on
}

// EntityManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EntityManagerCallerRaw struct {
	Contract *EntityManagerCaller // Generic read-only contract binding to access the raw methods on
}

// EntityManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EntityManagerTransactorRaw struct {
	Contract *EntityManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEntityManager creates a new instance of EntityManager, bound to a specific deployed contract.
func NewEntityManager(address common.Address, backend bind.ContractBackend) (*EntityManager, error) {
	contract, err := bindEntityManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EntityManager{EntityManagerCaller: EntityManagerCaller{contract: contract}, EntityManagerTransactor: EntityManagerTransactor{contract: contract}, EntityManagerFilterer: EntityManagerFilterer{contract: contract}}, nil
}

// NewEntityManagerCaller creates a new read-only instance of EntityManager, bound to a specific deployed contract.
func NewEntityManagerCaller(address common.Address, caller bind.ContractCaller) (*EntityManagerCaller, error) {
	contract, err := bindEntityManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EntityManagerCaller{contract: contract}, nil
}

// NewEntityManagerTransactor creates a new write-only instance of EntityManager, bound to a specific deployed contract.
func NewEntityManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*EntityManagerTransactor, error) {
	contract, err := bindEntityManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EntityManagerTransactor{contract: contract}, nil
}

// NewEntityManagerFilterer creates a new log filterer instance of EntityManager, bound to a specific deployed contract.
func NewEntityManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*EntityManagerFilterer, error) {
	contract, err := bindEntityManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EntityManagerFilterer{contract: contract}, nil
}

// bindEntityManager binds a generic wrapper to an already deployed contract.
func bindEntityManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EntityManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntityManager *EntityManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EntityManager.Contract.EntityManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntityManager *EntityManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntityManager.Contract.EntityManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntityManager *EntityManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntityManager.Contract.EntityManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EntityManager *EntityManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EntityManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EntityManager *EntityManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntityManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EntityManager *EntityManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EntityManager.Contract.contract.Transact(opts, method, params...)
}

// GetDelegationAddressOf is a free data retrieval call binding the contract method 0xa5d70de0.
//
// Solidity: function getDelegationAddressOf(address _voter) view returns(address)
func (_EntityManager *EntityManagerCaller) GetDelegationAddressOf(opts *bind.CallOpts, _voter common.Address) (common.Address, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getDelegationAddressOf", _voter)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDelegationAddressOf is a free data retrieval call binding the contract method 0xa5d70de0.
//
// Solidity: function getDelegationAddressOf(address _voter) view returns(address)
func (_EntityManager *EntityManagerSession) GetDelegationAddressOf(_voter common.Address) (common.Address, error) {
	return _EntityManager.Contract.GetDelegationAddressOf(&_EntityManager.CallOpts, _voter)
}

// GetDelegationAddressOf is a free data retrieval call binding the contract method 0xa5d70de0.
//
// Solidity: function getDelegationAddressOf(address _voter) view returns(address)
func (_EntityManager *EntityManagerCallerSession) GetDelegationAddressOf(_voter common.Address) (common.Address, error) {
	return _EntityManager.Contract.GetDelegationAddressOf(&_EntityManager.CallOpts, _voter)
}

// GetDelegationAddressOfAt is a free data retrieval call binding the contract method 0x2bedbf7a.
//
// Solidity: function getDelegationAddressOfAt(address _voter, uint256 _blockNumber) view returns(address)
func (_EntityManager *EntityManagerCaller) GetDelegationAddressOfAt(opts *bind.CallOpts, _voter common.Address, _blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getDelegationAddressOfAt", _voter, _blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDelegationAddressOfAt is a free data retrieval call binding the contract method 0x2bedbf7a.
//
// Solidity: function getDelegationAddressOfAt(address _voter, uint256 _blockNumber) view returns(address)
func (_EntityManager *EntityManagerSession) GetDelegationAddressOfAt(_voter common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetDelegationAddressOfAt(&_EntityManager.CallOpts, _voter, _blockNumber)
}

// GetDelegationAddressOfAt is a free data retrieval call binding the contract method 0x2bedbf7a.
//
// Solidity: function getDelegationAddressOfAt(address _voter, uint256 _blockNumber) view returns(address)
func (_EntityManager *EntityManagerCallerSession) GetDelegationAddressOfAt(_voter common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetDelegationAddressOfAt(&_EntityManager.CallOpts, _voter, _blockNumber)
}

// GetNodeIdsOf is a free data retrieval call binding the contract method 0xe6450b97.
//
// Solidity: function getNodeIdsOf(address _voter) view returns(bytes20[])
func (_EntityManager *EntityManagerCaller) GetNodeIdsOf(opts *bind.CallOpts, _voter common.Address) ([][20]byte, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getNodeIdsOf", _voter)

	if err != nil {
		return *new([][20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][20]byte)).(*[][20]byte)

	return out0, err

}

// GetNodeIdsOf is a free data retrieval call binding the contract method 0xe6450b97.
//
// Solidity: function getNodeIdsOf(address _voter) view returns(bytes20[])
func (_EntityManager *EntityManagerSession) GetNodeIdsOf(_voter common.Address) ([][20]byte, error) {
	return _EntityManager.Contract.GetNodeIdsOf(&_EntityManager.CallOpts, _voter)
}

// GetNodeIdsOf is a free data retrieval call binding the contract method 0xe6450b97.
//
// Solidity: function getNodeIdsOf(address _voter) view returns(bytes20[])
func (_EntityManager *EntityManagerCallerSession) GetNodeIdsOf(_voter common.Address) ([][20]byte, error) {
	return _EntityManager.Contract.GetNodeIdsOf(&_EntityManager.CallOpts, _voter)
}

// GetNodeIdsOfAt is a free data retrieval call binding the contract method 0x5b4be5f2.
//
// Solidity: function getNodeIdsOfAt(address _voter, uint256 _blockNumber) view returns(bytes20[])
func (_EntityManager *EntityManagerCaller) GetNodeIdsOfAt(opts *bind.CallOpts, _voter common.Address, _blockNumber *big.Int) ([][20]byte, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getNodeIdsOfAt", _voter, _blockNumber)

	if err != nil {
		return *new([][20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][20]byte)).(*[][20]byte)

	return out0, err

}

// GetNodeIdsOfAt is a free data retrieval call binding the contract method 0x5b4be5f2.
//
// Solidity: function getNodeIdsOfAt(address _voter, uint256 _blockNumber) view returns(bytes20[])
func (_EntityManager *EntityManagerSession) GetNodeIdsOfAt(_voter common.Address, _blockNumber *big.Int) ([][20]byte, error) {
	return _EntityManager.Contract.GetNodeIdsOfAt(&_EntityManager.CallOpts, _voter, _blockNumber)
}

// GetNodeIdsOfAt is a free data retrieval call binding the contract method 0x5b4be5f2.
//
// Solidity: function getNodeIdsOfAt(address _voter, uint256 _blockNumber) view returns(bytes20[])
func (_EntityManager *EntityManagerCallerSession) GetNodeIdsOfAt(_voter common.Address, _blockNumber *big.Int) ([][20]byte, error) {
	return _EntityManager.Contract.GetNodeIdsOfAt(&_EntityManager.CallOpts, _voter, _blockNumber)
}

// GetPublicKeyOf is a free data retrieval call binding the contract method 0x75e68605.
//
// Solidity: function getPublicKeyOf(address _voter) view returns(bytes32, bytes32)
func (_EntityManager *EntityManagerCaller) GetPublicKeyOf(opts *bind.CallOpts, _voter common.Address) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getPublicKeyOf", _voter)

	if err != nil {
		return *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetPublicKeyOf is a free data retrieval call binding the contract method 0x75e68605.
//
// Solidity: function getPublicKeyOf(address _voter) view returns(bytes32, bytes32)
func (_EntityManager *EntityManagerSession) GetPublicKeyOf(_voter common.Address) ([32]byte, [32]byte, error) {
	return _EntityManager.Contract.GetPublicKeyOf(&_EntityManager.CallOpts, _voter)
}

// GetPublicKeyOf is a free data retrieval call binding the contract method 0x75e68605.
//
// Solidity: function getPublicKeyOf(address _voter) view returns(bytes32, bytes32)
func (_EntityManager *EntityManagerCallerSession) GetPublicKeyOf(_voter common.Address) ([32]byte, [32]byte, error) {
	return _EntityManager.Contract.GetPublicKeyOf(&_EntityManager.CallOpts, _voter)
}

// GetPublicKeyOfAt is a free data retrieval call binding the contract method 0x11104561.
//
// Solidity: function getPublicKeyOfAt(address _voter, uint256 _blockNumber) view returns(bytes32, bytes32)
func (_EntityManager *EntityManagerCaller) GetPublicKeyOfAt(opts *bind.CallOpts, _voter common.Address, _blockNumber *big.Int) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getPublicKeyOfAt", _voter, _blockNumber)

	if err != nil {
		return *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// GetPublicKeyOfAt is a free data retrieval call binding the contract method 0x11104561.
//
// Solidity: function getPublicKeyOfAt(address _voter, uint256 _blockNumber) view returns(bytes32, bytes32)
func (_EntityManager *EntityManagerSession) GetPublicKeyOfAt(_voter common.Address, _blockNumber *big.Int) ([32]byte, [32]byte, error) {
	return _EntityManager.Contract.GetPublicKeyOfAt(&_EntityManager.CallOpts, _voter, _blockNumber)
}

// GetPublicKeyOfAt is a free data retrieval call binding the contract method 0x11104561.
//
// Solidity: function getPublicKeyOfAt(address _voter, uint256 _blockNumber) view returns(bytes32, bytes32)
func (_EntityManager *EntityManagerCallerSession) GetPublicKeyOfAt(_voter common.Address, _blockNumber *big.Int) ([32]byte, [32]byte, error) {
	return _EntityManager.Contract.GetPublicKeyOfAt(&_EntityManager.CallOpts, _voter, _blockNumber)
}

// GetVoterAddresses is a free data retrieval call binding the contract method 0xe5771dbc.
//
// Solidity: function getVoterAddresses(address _voter) view returns((address,address,address) _addresses)
func (_EntityManager *EntityManagerCaller) GetVoterAddresses(opts *bind.CallOpts, _voter common.Address) (IEntityManagerVoterAddresses, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getVoterAddresses", _voter)

	if err != nil {
		return *new(IEntityManagerVoterAddresses), err
	}

	out0 := *abi.ConvertType(out[0], new(IEntityManagerVoterAddresses)).(*IEntityManagerVoterAddresses)

	return out0, err

}

// GetVoterAddresses is a free data retrieval call binding the contract method 0xe5771dbc.
//
// Solidity: function getVoterAddresses(address _voter) view returns((address,address,address) _addresses)
func (_EntityManager *EntityManagerSession) GetVoterAddresses(_voter common.Address) (IEntityManagerVoterAddresses, error) {
	return _EntityManager.Contract.GetVoterAddresses(&_EntityManager.CallOpts, _voter)
}

// GetVoterAddresses is a free data retrieval call binding the contract method 0xe5771dbc.
//
// Solidity: function getVoterAddresses(address _voter) view returns((address,address,address) _addresses)
func (_EntityManager *EntityManagerCallerSession) GetVoterAddresses(_voter common.Address) (IEntityManagerVoterAddresses, error) {
	return _EntityManager.Contract.GetVoterAddresses(&_EntityManager.CallOpts, _voter)
}

// GetVoterAddressesAt is a free data retrieval call binding the contract method 0xb2519ac8.
//
// Solidity: function getVoterAddressesAt(address _voter, uint256 _blockNumber) view returns((address,address,address) _addresses)
func (_EntityManager *EntityManagerCaller) GetVoterAddressesAt(opts *bind.CallOpts, _voter common.Address, _blockNumber *big.Int) (IEntityManagerVoterAddresses, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getVoterAddressesAt", _voter, _blockNumber)

	if err != nil {
		return *new(IEntityManagerVoterAddresses), err
	}

	out0 := *abi.ConvertType(out[0], new(IEntityManagerVoterAddresses)).(*IEntityManagerVoterAddresses)

	return out0, err

}

// GetVoterAddressesAt is a free data retrieval call binding the contract method 0xb2519ac8.
//
// Solidity: function getVoterAddressesAt(address _voter, uint256 _blockNumber) view returns((address,address,address) _addresses)
func (_EntityManager *EntityManagerSession) GetVoterAddressesAt(_voter common.Address, _blockNumber *big.Int) (IEntityManagerVoterAddresses, error) {
	return _EntityManager.Contract.GetVoterAddressesAt(&_EntityManager.CallOpts, _voter, _blockNumber)
}

// GetVoterAddressesAt is a free data retrieval call binding the contract method 0xb2519ac8.
//
// Solidity: function getVoterAddressesAt(address _voter, uint256 _blockNumber) view returns((address,address,address) _addresses)
func (_EntityManager *EntityManagerCallerSession) GetVoterAddressesAt(_voter common.Address, _blockNumber *big.Int) (IEntityManagerVoterAddresses, error) {
	return _EntityManager.Contract.GetVoterAddressesAt(&_EntityManager.CallOpts, _voter, _blockNumber)
}

// GetVoterForDelegationAddress is a free data retrieval call binding the contract method 0xbc8c4a5f.
//
// Solidity: function getVoterForDelegationAddress(address _delegationAddress, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerCaller) GetVoterForDelegationAddress(opts *bind.CallOpts, _delegationAddress common.Address, _blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getVoterForDelegationAddress", _delegationAddress, _blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoterForDelegationAddress is a free data retrieval call binding the contract method 0xbc8c4a5f.
//
// Solidity: function getVoterForDelegationAddress(address _delegationAddress, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerSession) GetVoterForDelegationAddress(_delegationAddress common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetVoterForDelegationAddress(&_EntityManager.CallOpts, _delegationAddress, _blockNumber)
}

// GetVoterForDelegationAddress is a free data retrieval call binding the contract method 0xbc8c4a5f.
//
// Solidity: function getVoterForDelegationAddress(address _delegationAddress, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerCallerSession) GetVoterForDelegationAddress(_delegationAddress common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetVoterForDelegationAddress(&_EntityManager.CallOpts, _delegationAddress, _blockNumber)
}

// GetVoterForNodeId is a free data retrieval call binding the contract method 0x0a812fff.
//
// Solidity: function getVoterForNodeId(bytes20 _nodeId, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerCaller) GetVoterForNodeId(opts *bind.CallOpts, _nodeId [20]byte, _blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getVoterForNodeId", _nodeId, _blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoterForNodeId is a free data retrieval call binding the contract method 0x0a812fff.
//
// Solidity: function getVoterForNodeId(bytes20 _nodeId, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerSession) GetVoterForNodeId(_nodeId [20]byte, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetVoterForNodeId(&_EntityManager.CallOpts, _nodeId, _blockNumber)
}

// GetVoterForNodeId is a free data retrieval call binding the contract method 0x0a812fff.
//
// Solidity: function getVoterForNodeId(bytes20 _nodeId, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerCallerSession) GetVoterForNodeId(_nodeId [20]byte, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetVoterForNodeId(&_EntityManager.CallOpts, _nodeId, _blockNumber)
}

// GetVoterForPublicKey is a free data retrieval call binding the contract method 0x09c80fa6.
//
// Solidity: function getVoterForPublicKey(bytes32 _part1, bytes32 _part2, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerCaller) GetVoterForPublicKey(opts *bind.CallOpts, _part1 [32]byte, _part2 [32]byte, _blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getVoterForPublicKey", _part1, _part2, _blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoterForPublicKey is a free data retrieval call binding the contract method 0x09c80fa6.
//
// Solidity: function getVoterForPublicKey(bytes32 _part1, bytes32 _part2, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerSession) GetVoterForPublicKey(_part1 [32]byte, _part2 [32]byte, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetVoterForPublicKey(&_EntityManager.CallOpts, _part1, _part2, _blockNumber)
}

// GetVoterForPublicKey is a free data retrieval call binding the contract method 0x09c80fa6.
//
// Solidity: function getVoterForPublicKey(bytes32 _part1, bytes32 _part2, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerCallerSession) GetVoterForPublicKey(_part1 [32]byte, _part2 [32]byte, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetVoterForPublicKey(&_EntityManager.CallOpts, _part1, _part2, _blockNumber)
}

// GetVoterForSigningPolicyAddress is a free data retrieval call binding the contract method 0xe178ec4c.
//
// Solidity: function getVoterForSigningPolicyAddress(address _signingPolicyAddress, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerCaller) GetVoterForSigningPolicyAddress(opts *bind.CallOpts, _signingPolicyAddress common.Address, _blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getVoterForSigningPolicyAddress", _signingPolicyAddress, _blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoterForSigningPolicyAddress is a free data retrieval call binding the contract method 0xe178ec4c.
//
// Solidity: function getVoterForSigningPolicyAddress(address _signingPolicyAddress, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerSession) GetVoterForSigningPolicyAddress(_signingPolicyAddress common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetVoterForSigningPolicyAddress(&_EntityManager.CallOpts, _signingPolicyAddress, _blockNumber)
}

// GetVoterForSigningPolicyAddress is a free data retrieval call binding the contract method 0xe178ec4c.
//
// Solidity: function getVoterForSigningPolicyAddress(address _signingPolicyAddress, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerCallerSession) GetVoterForSigningPolicyAddress(_signingPolicyAddress common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetVoterForSigningPolicyAddress(&_EntityManager.CallOpts, _signingPolicyAddress, _blockNumber)
}

// GetVoterForSubmitAddress is a free data retrieval call binding the contract method 0x89292b0f.
//
// Solidity: function getVoterForSubmitAddress(address _submitAddress, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerCaller) GetVoterForSubmitAddress(opts *bind.CallOpts, _submitAddress common.Address, _blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getVoterForSubmitAddress", _submitAddress, _blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoterForSubmitAddress is a free data retrieval call binding the contract method 0x89292b0f.
//
// Solidity: function getVoterForSubmitAddress(address _submitAddress, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerSession) GetVoterForSubmitAddress(_submitAddress common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetVoterForSubmitAddress(&_EntityManager.CallOpts, _submitAddress, _blockNumber)
}

// GetVoterForSubmitAddress is a free data retrieval call binding the contract method 0x89292b0f.
//
// Solidity: function getVoterForSubmitAddress(address _submitAddress, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerCallerSession) GetVoterForSubmitAddress(_submitAddress common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetVoterForSubmitAddress(&_EntityManager.CallOpts, _submitAddress, _blockNumber)
}

// GetVoterForSubmitSignaturesAddress is a free data retrieval call binding the contract method 0x61bc320f.
//
// Solidity: function getVoterForSubmitSignaturesAddress(address _submitSignaturesAddress, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerCaller) GetVoterForSubmitSignaturesAddress(opts *bind.CallOpts, _submitSignaturesAddress common.Address, _blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EntityManager.contract.Call(opts, &out, "getVoterForSubmitSignaturesAddress", _submitSignaturesAddress, _blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVoterForSubmitSignaturesAddress is a free data retrieval call binding the contract method 0x61bc320f.
//
// Solidity: function getVoterForSubmitSignaturesAddress(address _submitSignaturesAddress, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerSession) GetVoterForSubmitSignaturesAddress(_submitSignaturesAddress common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetVoterForSubmitSignaturesAddress(&_EntityManager.CallOpts, _submitSignaturesAddress, _blockNumber)
}

// GetVoterForSubmitSignaturesAddress is a free data retrieval call binding the contract method 0x61bc320f.
//
// Solidity: function getVoterForSubmitSignaturesAddress(address _submitSignaturesAddress, uint256 _blockNumber) view returns(address _voter)
func (_EntityManager *EntityManagerCallerSession) GetVoterForSubmitSignaturesAddress(_submitSignaturesAddress common.Address, _blockNumber *big.Int) (common.Address, error) {
	return _EntityManager.Contract.GetVoterForSubmitSignaturesAddress(&_EntityManager.CallOpts, _submitSignaturesAddress, _blockNumber)
}

// ConfirmDelegationAddressRegistration is a paid mutator transaction binding the contract method 0xe5360006.
//
// Solidity: function confirmDelegationAddressRegistration(address _voter) returns()
func (_EntityManager *EntityManagerTransactor) ConfirmDelegationAddressRegistration(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _EntityManager.contract.Transact(opts, "confirmDelegationAddressRegistration", _voter)
}

// ConfirmDelegationAddressRegistration is a paid mutator transaction binding the contract method 0xe5360006.
//
// Solidity: function confirmDelegationAddressRegistration(address _voter) returns()
func (_EntityManager *EntityManagerSession) ConfirmDelegationAddressRegistration(_voter common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ConfirmDelegationAddressRegistration(&_EntityManager.TransactOpts, _voter)
}

// ConfirmDelegationAddressRegistration is a paid mutator transaction binding the contract method 0xe5360006.
//
// Solidity: function confirmDelegationAddressRegistration(address _voter) returns()
func (_EntityManager *EntityManagerTransactorSession) ConfirmDelegationAddressRegistration(_voter common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ConfirmDelegationAddressRegistration(&_EntityManager.TransactOpts, _voter)
}

// ConfirmSigningPolicyAddressRegistration is a paid mutator transaction binding the contract method 0xdb8d5125.
//
// Solidity: function confirmSigningPolicyAddressRegistration(address _voter) returns()
func (_EntityManager *EntityManagerTransactor) ConfirmSigningPolicyAddressRegistration(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _EntityManager.contract.Transact(opts, "confirmSigningPolicyAddressRegistration", _voter)
}

// ConfirmSigningPolicyAddressRegistration is a paid mutator transaction binding the contract method 0xdb8d5125.
//
// Solidity: function confirmSigningPolicyAddressRegistration(address _voter) returns()
func (_EntityManager *EntityManagerSession) ConfirmSigningPolicyAddressRegistration(_voter common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ConfirmSigningPolicyAddressRegistration(&_EntityManager.TransactOpts, _voter)
}

// ConfirmSigningPolicyAddressRegistration is a paid mutator transaction binding the contract method 0xdb8d5125.
//
// Solidity: function confirmSigningPolicyAddressRegistration(address _voter) returns()
func (_EntityManager *EntityManagerTransactorSession) ConfirmSigningPolicyAddressRegistration(_voter common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ConfirmSigningPolicyAddressRegistration(&_EntityManager.TransactOpts, _voter)
}

// ConfirmSubmitAddressRegistration is a paid mutator transaction binding the contract method 0x5f1fb56a.
//
// Solidity: function confirmSubmitAddressRegistration(address _voter) returns()
func (_EntityManager *EntityManagerTransactor) ConfirmSubmitAddressRegistration(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _EntityManager.contract.Transact(opts, "confirmSubmitAddressRegistration", _voter)
}

// ConfirmSubmitAddressRegistration is a paid mutator transaction binding the contract method 0x5f1fb56a.
//
// Solidity: function confirmSubmitAddressRegistration(address _voter) returns()
func (_EntityManager *EntityManagerSession) ConfirmSubmitAddressRegistration(_voter common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ConfirmSubmitAddressRegistration(&_EntityManager.TransactOpts, _voter)
}

// ConfirmSubmitAddressRegistration is a paid mutator transaction binding the contract method 0x5f1fb56a.
//
// Solidity: function confirmSubmitAddressRegistration(address _voter) returns()
func (_EntityManager *EntityManagerTransactorSession) ConfirmSubmitAddressRegistration(_voter common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ConfirmSubmitAddressRegistration(&_EntityManager.TransactOpts, _voter)
}

// ConfirmSubmitSignaturesAddressRegistration is a paid mutator transaction binding the contract method 0xbda1a84e.
//
// Solidity: function confirmSubmitSignaturesAddressRegistration(address _voter) returns()
func (_EntityManager *EntityManagerTransactor) ConfirmSubmitSignaturesAddressRegistration(opts *bind.TransactOpts, _voter common.Address) (*types.Transaction, error) {
	return _EntityManager.contract.Transact(opts, "confirmSubmitSignaturesAddressRegistration", _voter)
}

// ConfirmSubmitSignaturesAddressRegistration is a paid mutator transaction binding the contract method 0xbda1a84e.
//
// Solidity: function confirmSubmitSignaturesAddressRegistration(address _voter) returns()
func (_EntityManager *EntityManagerSession) ConfirmSubmitSignaturesAddressRegistration(_voter common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ConfirmSubmitSignaturesAddressRegistration(&_EntityManager.TransactOpts, _voter)
}

// ConfirmSubmitSignaturesAddressRegistration is a paid mutator transaction binding the contract method 0xbda1a84e.
//
// Solidity: function confirmSubmitSignaturesAddressRegistration(address _voter) returns()
func (_EntityManager *EntityManagerTransactorSession) ConfirmSubmitSignaturesAddressRegistration(_voter common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ConfirmSubmitSignaturesAddressRegistration(&_EntityManager.TransactOpts, _voter)
}

// ProposeDelegationAddress is a paid mutator transaction binding the contract method 0x5712b9f3.
//
// Solidity: function proposeDelegationAddress(address _delegationAddress) returns()
func (_EntityManager *EntityManagerTransactor) ProposeDelegationAddress(opts *bind.TransactOpts, _delegationAddress common.Address) (*types.Transaction, error) {
	return _EntityManager.contract.Transact(opts, "proposeDelegationAddress", _delegationAddress)
}

// ProposeDelegationAddress is a paid mutator transaction binding the contract method 0x5712b9f3.
//
// Solidity: function proposeDelegationAddress(address _delegationAddress) returns()
func (_EntityManager *EntityManagerSession) ProposeDelegationAddress(_delegationAddress common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ProposeDelegationAddress(&_EntityManager.TransactOpts, _delegationAddress)
}

// ProposeDelegationAddress is a paid mutator transaction binding the contract method 0x5712b9f3.
//
// Solidity: function proposeDelegationAddress(address _delegationAddress) returns()
func (_EntityManager *EntityManagerTransactorSession) ProposeDelegationAddress(_delegationAddress common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ProposeDelegationAddress(&_EntityManager.TransactOpts, _delegationAddress)
}

// ProposeSigningPolicyAddress is a paid mutator transaction binding the contract method 0x69b5fed9.
//
// Solidity: function proposeSigningPolicyAddress(address _signingPolicyAddress) returns()
func (_EntityManager *EntityManagerTransactor) ProposeSigningPolicyAddress(opts *bind.TransactOpts, _signingPolicyAddress common.Address) (*types.Transaction, error) {
	return _EntityManager.contract.Transact(opts, "proposeSigningPolicyAddress", _signingPolicyAddress)
}

// ProposeSigningPolicyAddress is a paid mutator transaction binding the contract method 0x69b5fed9.
//
// Solidity: function proposeSigningPolicyAddress(address _signingPolicyAddress) returns()
func (_EntityManager *EntityManagerSession) ProposeSigningPolicyAddress(_signingPolicyAddress common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ProposeSigningPolicyAddress(&_EntityManager.TransactOpts, _signingPolicyAddress)
}

// ProposeSigningPolicyAddress is a paid mutator transaction binding the contract method 0x69b5fed9.
//
// Solidity: function proposeSigningPolicyAddress(address _signingPolicyAddress) returns()
func (_EntityManager *EntityManagerTransactorSession) ProposeSigningPolicyAddress(_signingPolicyAddress common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ProposeSigningPolicyAddress(&_EntityManager.TransactOpts, _signingPolicyAddress)
}

// ProposeSubmitAddress is a paid mutator transaction binding the contract method 0x87607e41.
//
// Solidity: function proposeSubmitAddress(address _submitAddress) returns()
func (_EntityManager *EntityManagerTransactor) ProposeSubmitAddress(opts *bind.TransactOpts, _submitAddress common.Address) (*types.Transaction, error) {
	return _EntityManager.contract.Transact(opts, "proposeSubmitAddress", _submitAddress)
}

// ProposeSubmitAddress is a paid mutator transaction binding the contract method 0x87607e41.
//
// Solidity: function proposeSubmitAddress(address _submitAddress) returns()
func (_EntityManager *EntityManagerSession) ProposeSubmitAddress(_submitAddress common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ProposeSubmitAddress(&_EntityManager.TransactOpts, _submitAddress)
}

// ProposeSubmitAddress is a paid mutator transaction binding the contract method 0x87607e41.
//
// Solidity: function proposeSubmitAddress(address _submitAddress) returns()
func (_EntityManager *EntityManagerTransactorSession) ProposeSubmitAddress(_submitAddress common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ProposeSubmitAddress(&_EntityManager.TransactOpts, _submitAddress)
}

// ProposeSubmitSignaturesAddress is a paid mutator transaction binding the contract method 0x36edd357.
//
// Solidity: function proposeSubmitSignaturesAddress(address _submitSignaturesAddress) returns()
func (_EntityManager *EntityManagerTransactor) ProposeSubmitSignaturesAddress(opts *bind.TransactOpts, _submitSignaturesAddress common.Address) (*types.Transaction, error) {
	return _EntityManager.contract.Transact(opts, "proposeSubmitSignaturesAddress", _submitSignaturesAddress)
}

// ProposeSubmitSignaturesAddress is a paid mutator transaction binding the contract method 0x36edd357.
//
// Solidity: function proposeSubmitSignaturesAddress(address _submitSignaturesAddress) returns()
func (_EntityManager *EntityManagerSession) ProposeSubmitSignaturesAddress(_submitSignaturesAddress common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ProposeSubmitSignaturesAddress(&_EntityManager.TransactOpts, _submitSignaturesAddress)
}

// ProposeSubmitSignaturesAddress is a paid mutator transaction binding the contract method 0x36edd357.
//
// Solidity: function proposeSubmitSignaturesAddress(address _submitSignaturesAddress) returns()
func (_EntityManager *EntityManagerTransactorSession) ProposeSubmitSignaturesAddress(_submitSignaturesAddress common.Address) (*types.Transaction, error) {
	return _EntityManager.Contract.ProposeSubmitSignaturesAddress(&_EntityManager.TransactOpts, _submitSignaturesAddress)
}

// RegisterNodeId is a paid mutator transaction binding the contract method 0x019476a7.
//
// Solidity: function registerNodeId(bytes20 _nodeId, bytes _certificateRaw, bytes _signature) returns()
func (_EntityManager *EntityManagerTransactor) RegisterNodeId(opts *bind.TransactOpts, _nodeId [20]byte, _certificateRaw []byte, _signature []byte) (*types.Transaction, error) {
	return _EntityManager.contract.Transact(opts, "registerNodeId", _nodeId, _certificateRaw, _signature)
}

// RegisterNodeId is a paid mutator transaction binding the contract method 0x019476a7.
//
// Solidity: function registerNodeId(bytes20 _nodeId, bytes _certificateRaw, bytes _signature) returns()
func (_EntityManager *EntityManagerSession) RegisterNodeId(_nodeId [20]byte, _certificateRaw []byte, _signature []byte) (*types.Transaction, error) {
	return _EntityManager.Contract.RegisterNodeId(&_EntityManager.TransactOpts, _nodeId, _certificateRaw, _signature)
}

// RegisterNodeId is a paid mutator transaction binding the contract method 0x019476a7.
//
// Solidity: function registerNodeId(bytes20 _nodeId, bytes _certificateRaw, bytes _signature) returns()
func (_EntityManager *EntityManagerTransactorSession) RegisterNodeId(_nodeId [20]byte, _certificateRaw []byte, _signature []byte) (*types.Transaction, error) {
	return _EntityManager.Contract.RegisterNodeId(&_EntityManager.TransactOpts, _nodeId, _certificateRaw, _signature)
}

// RegisterPublicKey is a paid mutator transaction binding the contract method 0xb1863cff.
//
// Solidity: function registerPublicKey(bytes32 _part1, bytes32 _part2, bytes _verificationData) returns()
func (_EntityManager *EntityManagerTransactor) RegisterPublicKey(opts *bind.TransactOpts, _part1 [32]byte, _part2 [32]byte, _verificationData []byte) (*types.Transaction, error) {
	return _EntityManager.contract.Transact(opts, "registerPublicKey", _part1, _part2, _verificationData)
}

// RegisterPublicKey is a paid mutator transaction binding the contract method 0xb1863cff.
//
// Solidity: function registerPublicKey(bytes32 _part1, bytes32 _part2, bytes _verificationData) returns()
func (_EntityManager *EntityManagerSession) RegisterPublicKey(_part1 [32]byte, _part2 [32]byte, _verificationData []byte) (*types.Transaction, error) {
	return _EntityManager.Contract.RegisterPublicKey(&_EntityManager.TransactOpts, _part1, _part2, _verificationData)
}

// RegisterPublicKey is a paid mutator transaction binding the contract method 0xb1863cff.
//
// Solidity: function registerPublicKey(bytes32 _part1, bytes32 _part2, bytes _verificationData) returns()
func (_EntityManager *EntityManagerTransactorSession) RegisterPublicKey(_part1 [32]byte, _part2 [32]byte, _verificationData []byte) (*types.Transaction, error) {
	return _EntityManager.Contract.RegisterPublicKey(&_EntityManager.TransactOpts, _part1, _part2, _verificationData)
}

// UnregisterNodeId is a paid mutator transaction binding the contract method 0x18c03f94.
//
// Solidity: function unregisterNodeId(bytes20 _nodeId) returns()
func (_EntityManager *EntityManagerTransactor) UnregisterNodeId(opts *bind.TransactOpts, _nodeId [20]byte) (*types.Transaction, error) {
	return _EntityManager.contract.Transact(opts, "unregisterNodeId", _nodeId)
}

// UnregisterNodeId is a paid mutator transaction binding the contract method 0x18c03f94.
//
// Solidity: function unregisterNodeId(bytes20 _nodeId) returns()
func (_EntityManager *EntityManagerSession) UnregisterNodeId(_nodeId [20]byte) (*types.Transaction, error) {
	return _EntityManager.Contract.UnregisterNodeId(&_EntityManager.TransactOpts, _nodeId)
}

// UnregisterNodeId is a paid mutator transaction binding the contract method 0x18c03f94.
//
// Solidity: function unregisterNodeId(bytes20 _nodeId) returns()
func (_EntityManager *EntityManagerTransactorSession) UnregisterNodeId(_nodeId [20]byte) (*types.Transaction, error) {
	return _EntityManager.Contract.UnregisterNodeId(&_EntityManager.TransactOpts, _nodeId)
}

// UnregisterPublicKey is a paid mutator transaction binding the contract method 0xe1e3cdac.
//
// Solidity: function unregisterPublicKey() returns()
func (_EntityManager *EntityManagerTransactor) UnregisterPublicKey(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EntityManager.contract.Transact(opts, "unregisterPublicKey")
}

// UnregisterPublicKey is a paid mutator transaction binding the contract method 0xe1e3cdac.
//
// Solidity: function unregisterPublicKey() returns()
func (_EntityManager *EntityManagerSession) UnregisterPublicKey() (*types.Transaction, error) {
	return _EntityManager.Contract.UnregisterPublicKey(&_EntityManager.TransactOpts)
}

// UnregisterPublicKey is a paid mutator transaction binding the contract method 0xe1e3cdac.
//
// Solidity: function unregisterPublicKey() returns()
func (_EntityManager *EntityManagerTransactorSession) UnregisterPublicKey() (*types.Transaction, error) {
	return _EntityManager.Contract.UnregisterPublicKey(&_EntityManager.TransactOpts)
}

// EntityManagerDelegationAddressProposedIterator is returned from FilterDelegationAddressProposed and is used to iterate over the raw logs and unpacked data for DelegationAddressProposed events raised by the EntityManager contract.
type EntityManagerDelegationAddressProposedIterator struct {
	Event *EntityManagerDelegationAddressProposed // Event containing the contract specifics and raw log

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
func (it *EntityManagerDelegationAddressProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityManagerDelegationAddressProposed)
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
		it.Event = new(EntityManagerDelegationAddressProposed)
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
func (it *EntityManagerDelegationAddressProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityManagerDelegationAddressProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityManagerDelegationAddressProposed represents a DelegationAddressProposed event raised by the EntityManager contract.
type EntityManagerDelegationAddressProposed struct {
	Voter             common.Address
	DelegationAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDelegationAddressProposed is a free log retrieval operation binding the contract event 0xed9e746b8b92f73a860c368d322985a6a0ab72ab1a3068eb69f708b495126e0f.
//
// Solidity: event DelegationAddressProposed(address indexed voter, address indexed delegationAddress)
func (_EntityManager *EntityManagerFilterer) FilterDelegationAddressProposed(opts *bind.FilterOpts, voter []common.Address, delegationAddress []common.Address) (*EntityManagerDelegationAddressProposedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var delegationAddressRule []interface{}
	for _, delegationAddressItem := range delegationAddress {
		delegationAddressRule = append(delegationAddressRule, delegationAddressItem)
	}

	logs, sub, err := _EntityManager.contract.FilterLogs(opts, "DelegationAddressProposed", voterRule, delegationAddressRule)
	if err != nil {
		return nil, err
	}
	return &EntityManagerDelegationAddressProposedIterator{contract: _EntityManager.contract, event: "DelegationAddressProposed", logs: logs, sub: sub}, nil
}

// WatchDelegationAddressProposed is a free log subscription operation binding the contract event 0xed9e746b8b92f73a860c368d322985a6a0ab72ab1a3068eb69f708b495126e0f.
//
// Solidity: event DelegationAddressProposed(address indexed voter, address indexed delegationAddress)
func (_EntityManager *EntityManagerFilterer) WatchDelegationAddressProposed(opts *bind.WatchOpts, sink chan<- *EntityManagerDelegationAddressProposed, voter []common.Address, delegationAddress []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var delegationAddressRule []interface{}
	for _, delegationAddressItem := range delegationAddress {
		delegationAddressRule = append(delegationAddressRule, delegationAddressItem)
	}

	logs, sub, err := _EntityManager.contract.WatchLogs(opts, "DelegationAddressProposed", voterRule, delegationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityManagerDelegationAddressProposed)
				if err := _EntityManager.contract.UnpackLog(event, "DelegationAddressProposed", log); err != nil {
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

// ParseDelegationAddressProposed is a log parse operation binding the contract event 0xed9e746b8b92f73a860c368d322985a6a0ab72ab1a3068eb69f708b495126e0f.
//
// Solidity: event DelegationAddressProposed(address indexed voter, address indexed delegationAddress)
func (_EntityManager *EntityManagerFilterer) ParseDelegationAddressProposed(log types.Log) (*EntityManagerDelegationAddressProposed, error) {
	event := new(EntityManagerDelegationAddressProposed)
	if err := _EntityManager.contract.UnpackLog(event, "DelegationAddressProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityManagerDelegationAddressRegistrationConfirmedIterator is returned from FilterDelegationAddressRegistrationConfirmed and is used to iterate over the raw logs and unpacked data for DelegationAddressRegistrationConfirmed events raised by the EntityManager contract.
type EntityManagerDelegationAddressRegistrationConfirmedIterator struct {
	Event *EntityManagerDelegationAddressRegistrationConfirmed // Event containing the contract specifics and raw log

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
func (it *EntityManagerDelegationAddressRegistrationConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityManagerDelegationAddressRegistrationConfirmed)
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
		it.Event = new(EntityManagerDelegationAddressRegistrationConfirmed)
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
func (it *EntityManagerDelegationAddressRegistrationConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityManagerDelegationAddressRegistrationConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityManagerDelegationAddressRegistrationConfirmed represents a DelegationAddressRegistrationConfirmed event raised by the EntityManager contract.
type EntityManagerDelegationAddressRegistrationConfirmed struct {
	Voter             common.Address
	DelegationAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDelegationAddressRegistrationConfirmed is a free log retrieval operation binding the contract event 0x614066d64432bd11a481cb202e15398e6b92baffd86c2573f54f9e9cd6c29ee4.
//
// Solidity: event DelegationAddressRegistrationConfirmed(address indexed voter, address indexed delegationAddress)
func (_EntityManager *EntityManagerFilterer) FilterDelegationAddressRegistrationConfirmed(opts *bind.FilterOpts, voter []common.Address, delegationAddress []common.Address) (*EntityManagerDelegationAddressRegistrationConfirmedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var delegationAddressRule []interface{}
	for _, delegationAddressItem := range delegationAddress {
		delegationAddressRule = append(delegationAddressRule, delegationAddressItem)
	}

	logs, sub, err := _EntityManager.contract.FilterLogs(opts, "DelegationAddressRegistrationConfirmed", voterRule, delegationAddressRule)
	if err != nil {
		return nil, err
	}
	return &EntityManagerDelegationAddressRegistrationConfirmedIterator{contract: _EntityManager.contract, event: "DelegationAddressRegistrationConfirmed", logs: logs, sub: sub}, nil
}

// WatchDelegationAddressRegistrationConfirmed is a free log subscription operation binding the contract event 0x614066d64432bd11a481cb202e15398e6b92baffd86c2573f54f9e9cd6c29ee4.
//
// Solidity: event DelegationAddressRegistrationConfirmed(address indexed voter, address indexed delegationAddress)
func (_EntityManager *EntityManagerFilterer) WatchDelegationAddressRegistrationConfirmed(opts *bind.WatchOpts, sink chan<- *EntityManagerDelegationAddressRegistrationConfirmed, voter []common.Address, delegationAddress []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var delegationAddressRule []interface{}
	for _, delegationAddressItem := range delegationAddress {
		delegationAddressRule = append(delegationAddressRule, delegationAddressItem)
	}

	logs, sub, err := _EntityManager.contract.WatchLogs(opts, "DelegationAddressRegistrationConfirmed", voterRule, delegationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityManagerDelegationAddressRegistrationConfirmed)
				if err := _EntityManager.contract.UnpackLog(event, "DelegationAddressRegistrationConfirmed", log); err != nil {
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

// ParseDelegationAddressRegistrationConfirmed is a log parse operation binding the contract event 0x614066d64432bd11a481cb202e15398e6b92baffd86c2573f54f9e9cd6c29ee4.
//
// Solidity: event DelegationAddressRegistrationConfirmed(address indexed voter, address indexed delegationAddress)
func (_EntityManager *EntityManagerFilterer) ParseDelegationAddressRegistrationConfirmed(log types.Log) (*EntityManagerDelegationAddressRegistrationConfirmed, error) {
	event := new(EntityManagerDelegationAddressRegistrationConfirmed)
	if err := _EntityManager.contract.UnpackLog(event, "DelegationAddressRegistrationConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityManagerMaxNodeIdsPerEntitySetIterator is returned from FilterMaxNodeIdsPerEntitySet and is used to iterate over the raw logs and unpacked data for MaxNodeIdsPerEntitySet events raised by the EntityManager contract.
type EntityManagerMaxNodeIdsPerEntitySetIterator struct {
	Event *EntityManagerMaxNodeIdsPerEntitySet // Event containing the contract specifics and raw log

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
func (it *EntityManagerMaxNodeIdsPerEntitySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityManagerMaxNodeIdsPerEntitySet)
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
		it.Event = new(EntityManagerMaxNodeIdsPerEntitySet)
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
func (it *EntityManagerMaxNodeIdsPerEntitySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityManagerMaxNodeIdsPerEntitySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityManagerMaxNodeIdsPerEntitySet represents a MaxNodeIdsPerEntitySet event raised by the EntityManager contract.
type EntityManagerMaxNodeIdsPerEntitySet struct {
	MaxNodeIdsPerEntity *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMaxNodeIdsPerEntitySet is a free log retrieval operation binding the contract event 0xce75fee15f61236d01d8324caa31f041cd043f94ae234d463ae5b88bd2ad7e57.
//
// Solidity: event MaxNodeIdsPerEntitySet(uint256 maxNodeIdsPerEntity)
func (_EntityManager *EntityManagerFilterer) FilterMaxNodeIdsPerEntitySet(opts *bind.FilterOpts) (*EntityManagerMaxNodeIdsPerEntitySetIterator, error) {

	logs, sub, err := _EntityManager.contract.FilterLogs(opts, "MaxNodeIdsPerEntitySet")
	if err != nil {
		return nil, err
	}
	return &EntityManagerMaxNodeIdsPerEntitySetIterator{contract: _EntityManager.contract, event: "MaxNodeIdsPerEntitySet", logs: logs, sub: sub}, nil
}

// WatchMaxNodeIdsPerEntitySet is a free log subscription operation binding the contract event 0xce75fee15f61236d01d8324caa31f041cd043f94ae234d463ae5b88bd2ad7e57.
//
// Solidity: event MaxNodeIdsPerEntitySet(uint256 maxNodeIdsPerEntity)
func (_EntityManager *EntityManagerFilterer) WatchMaxNodeIdsPerEntitySet(opts *bind.WatchOpts, sink chan<- *EntityManagerMaxNodeIdsPerEntitySet) (event.Subscription, error) {

	logs, sub, err := _EntityManager.contract.WatchLogs(opts, "MaxNodeIdsPerEntitySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityManagerMaxNodeIdsPerEntitySet)
				if err := _EntityManager.contract.UnpackLog(event, "MaxNodeIdsPerEntitySet", log); err != nil {
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

// ParseMaxNodeIdsPerEntitySet is a log parse operation binding the contract event 0xce75fee15f61236d01d8324caa31f041cd043f94ae234d463ae5b88bd2ad7e57.
//
// Solidity: event MaxNodeIdsPerEntitySet(uint256 maxNodeIdsPerEntity)
func (_EntityManager *EntityManagerFilterer) ParseMaxNodeIdsPerEntitySet(log types.Log) (*EntityManagerMaxNodeIdsPerEntitySet, error) {
	event := new(EntityManagerMaxNodeIdsPerEntitySet)
	if err := _EntityManager.contract.UnpackLog(event, "MaxNodeIdsPerEntitySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityManagerNodeIdRegisteredIterator is returned from FilterNodeIdRegistered and is used to iterate over the raw logs and unpacked data for NodeIdRegistered events raised by the EntityManager contract.
type EntityManagerNodeIdRegisteredIterator struct {
	Event *EntityManagerNodeIdRegistered // Event containing the contract specifics and raw log

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
func (it *EntityManagerNodeIdRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityManagerNodeIdRegistered)
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
		it.Event = new(EntityManagerNodeIdRegistered)
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
func (it *EntityManagerNodeIdRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityManagerNodeIdRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityManagerNodeIdRegistered represents a NodeIdRegistered event raised by the EntityManager contract.
type EntityManagerNodeIdRegistered struct {
	Voter  common.Address
	NodeId [20]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNodeIdRegistered is a free log retrieval operation binding the contract event 0xcc83e5a4ea6e4a90196f405439628da7c929529e42fbb6af2e388f9c6a2fd6f0.
//
// Solidity: event NodeIdRegistered(address indexed voter, bytes20 indexed nodeId)
func (_EntityManager *EntityManagerFilterer) FilterNodeIdRegistered(opts *bind.FilterOpts, voter []common.Address, nodeId [][20]byte) (*EntityManagerNodeIdRegisteredIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _EntityManager.contract.FilterLogs(opts, "NodeIdRegistered", voterRule, nodeIdRule)
	if err != nil {
		return nil, err
	}
	return &EntityManagerNodeIdRegisteredIterator{contract: _EntityManager.contract, event: "NodeIdRegistered", logs: logs, sub: sub}, nil
}

// WatchNodeIdRegistered is a free log subscription operation binding the contract event 0xcc83e5a4ea6e4a90196f405439628da7c929529e42fbb6af2e388f9c6a2fd6f0.
//
// Solidity: event NodeIdRegistered(address indexed voter, bytes20 indexed nodeId)
func (_EntityManager *EntityManagerFilterer) WatchNodeIdRegistered(opts *bind.WatchOpts, sink chan<- *EntityManagerNodeIdRegistered, voter []common.Address, nodeId [][20]byte) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _EntityManager.contract.WatchLogs(opts, "NodeIdRegistered", voterRule, nodeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityManagerNodeIdRegistered)
				if err := _EntityManager.contract.UnpackLog(event, "NodeIdRegistered", log); err != nil {
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

// ParseNodeIdRegistered is a log parse operation binding the contract event 0xcc83e5a4ea6e4a90196f405439628da7c929529e42fbb6af2e388f9c6a2fd6f0.
//
// Solidity: event NodeIdRegistered(address indexed voter, bytes20 indexed nodeId)
func (_EntityManager *EntityManagerFilterer) ParseNodeIdRegistered(log types.Log) (*EntityManagerNodeIdRegistered, error) {
	event := new(EntityManagerNodeIdRegistered)
	if err := _EntityManager.contract.UnpackLog(event, "NodeIdRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityManagerNodeIdUnregisteredIterator is returned from FilterNodeIdUnregistered and is used to iterate over the raw logs and unpacked data for NodeIdUnregistered events raised by the EntityManager contract.
type EntityManagerNodeIdUnregisteredIterator struct {
	Event *EntityManagerNodeIdUnregistered // Event containing the contract specifics and raw log

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
func (it *EntityManagerNodeIdUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityManagerNodeIdUnregistered)
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
		it.Event = new(EntityManagerNodeIdUnregistered)
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
func (it *EntityManagerNodeIdUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityManagerNodeIdUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityManagerNodeIdUnregistered represents a NodeIdUnregistered event raised by the EntityManager contract.
type EntityManagerNodeIdUnregistered struct {
	Voter  common.Address
	NodeId [20]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNodeIdUnregistered is a free log retrieval operation binding the contract event 0x44d1a340b419d3f42da5eb85de7715aecbaf6376855ef4f64a1e033e094231d2.
//
// Solidity: event NodeIdUnregistered(address indexed voter, bytes20 indexed nodeId)
func (_EntityManager *EntityManagerFilterer) FilterNodeIdUnregistered(opts *bind.FilterOpts, voter []common.Address, nodeId [][20]byte) (*EntityManagerNodeIdUnregisteredIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _EntityManager.contract.FilterLogs(opts, "NodeIdUnregistered", voterRule, nodeIdRule)
	if err != nil {
		return nil, err
	}
	return &EntityManagerNodeIdUnregisteredIterator{contract: _EntityManager.contract, event: "NodeIdUnregistered", logs: logs, sub: sub}, nil
}

// WatchNodeIdUnregistered is a free log subscription operation binding the contract event 0x44d1a340b419d3f42da5eb85de7715aecbaf6376855ef4f64a1e033e094231d2.
//
// Solidity: event NodeIdUnregistered(address indexed voter, bytes20 indexed nodeId)
func (_EntityManager *EntityManagerFilterer) WatchNodeIdUnregistered(opts *bind.WatchOpts, sink chan<- *EntityManagerNodeIdUnregistered, voter []common.Address, nodeId [][20]byte) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _EntityManager.contract.WatchLogs(opts, "NodeIdUnregistered", voterRule, nodeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityManagerNodeIdUnregistered)
				if err := _EntityManager.contract.UnpackLog(event, "NodeIdUnregistered", log); err != nil {
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

// ParseNodeIdUnregistered is a log parse operation binding the contract event 0x44d1a340b419d3f42da5eb85de7715aecbaf6376855ef4f64a1e033e094231d2.
//
// Solidity: event NodeIdUnregistered(address indexed voter, bytes20 indexed nodeId)
func (_EntityManager *EntityManagerFilterer) ParseNodeIdUnregistered(log types.Log) (*EntityManagerNodeIdUnregistered, error) {
	event := new(EntityManagerNodeIdUnregistered)
	if err := _EntityManager.contract.UnpackLog(event, "NodeIdUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityManagerPublicKeyRegisteredIterator is returned from FilterPublicKeyRegistered and is used to iterate over the raw logs and unpacked data for PublicKeyRegistered events raised by the EntityManager contract.
type EntityManagerPublicKeyRegisteredIterator struct {
	Event *EntityManagerPublicKeyRegistered // Event containing the contract specifics and raw log

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
func (it *EntityManagerPublicKeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityManagerPublicKeyRegistered)
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
		it.Event = new(EntityManagerPublicKeyRegistered)
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
func (it *EntityManagerPublicKeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityManagerPublicKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityManagerPublicKeyRegistered represents a PublicKeyRegistered event raised by the EntityManager contract.
type EntityManagerPublicKeyRegistered struct {
	Voter common.Address
	Part1 [32]byte
	Part2 [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyRegistered is a free log retrieval operation binding the contract event 0x857427b1f3d6fc86cfeeb2c86d75772cf01de4519a0ff8d413418b55e7680745.
//
// Solidity: event PublicKeyRegistered(address indexed voter, bytes32 indexed part1, bytes32 indexed part2)
func (_EntityManager *EntityManagerFilterer) FilterPublicKeyRegistered(opts *bind.FilterOpts, voter []common.Address, part1 [][32]byte, part2 [][32]byte) (*EntityManagerPublicKeyRegisteredIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var part1Rule []interface{}
	for _, part1Item := range part1 {
		part1Rule = append(part1Rule, part1Item)
	}
	var part2Rule []interface{}
	for _, part2Item := range part2 {
		part2Rule = append(part2Rule, part2Item)
	}

	logs, sub, err := _EntityManager.contract.FilterLogs(opts, "PublicKeyRegistered", voterRule, part1Rule, part2Rule)
	if err != nil {
		return nil, err
	}
	return &EntityManagerPublicKeyRegisteredIterator{contract: _EntityManager.contract, event: "PublicKeyRegistered", logs: logs, sub: sub}, nil
}

// WatchPublicKeyRegistered is a free log subscription operation binding the contract event 0x857427b1f3d6fc86cfeeb2c86d75772cf01de4519a0ff8d413418b55e7680745.
//
// Solidity: event PublicKeyRegistered(address indexed voter, bytes32 indexed part1, bytes32 indexed part2)
func (_EntityManager *EntityManagerFilterer) WatchPublicKeyRegistered(opts *bind.WatchOpts, sink chan<- *EntityManagerPublicKeyRegistered, voter []common.Address, part1 [][32]byte, part2 [][32]byte) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var part1Rule []interface{}
	for _, part1Item := range part1 {
		part1Rule = append(part1Rule, part1Item)
	}
	var part2Rule []interface{}
	for _, part2Item := range part2 {
		part2Rule = append(part2Rule, part2Item)
	}

	logs, sub, err := _EntityManager.contract.WatchLogs(opts, "PublicKeyRegistered", voterRule, part1Rule, part2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityManagerPublicKeyRegistered)
				if err := _EntityManager.contract.UnpackLog(event, "PublicKeyRegistered", log); err != nil {
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

// ParsePublicKeyRegistered is a log parse operation binding the contract event 0x857427b1f3d6fc86cfeeb2c86d75772cf01de4519a0ff8d413418b55e7680745.
//
// Solidity: event PublicKeyRegistered(address indexed voter, bytes32 indexed part1, bytes32 indexed part2)
func (_EntityManager *EntityManagerFilterer) ParsePublicKeyRegistered(log types.Log) (*EntityManagerPublicKeyRegistered, error) {
	event := new(EntityManagerPublicKeyRegistered)
	if err := _EntityManager.contract.UnpackLog(event, "PublicKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityManagerPublicKeyUnregisteredIterator is returned from FilterPublicKeyUnregistered and is used to iterate over the raw logs and unpacked data for PublicKeyUnregistered events raised by the EntityManager contract.
type EntityManagerPublicKeyUnregisteredIterator struct {
	Event *EntityManagerPublicKeyUnregistered // Event containing the contract specifics and raw log

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
func (it *EntityManagerPublicKeyUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityManagerPublicKeyUnregistered)
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
		it.Event = new(EntityManagerPublicKeyUnregistered)
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
func (it *EntityManagerPublicKeyUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityManagerPublicKeyUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityManagerPublicKeyUnregistered represents a PublicKeyUnregistered event raised by the EntityManager contract.
type EntityManagerPublicKeyUnregistered struct {
	Voter common.Address
	Part1 [32]byte
	Part2 [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyUnregistered is a free log retrieval operation binding the contract event 0xc82d1b006dff077f4b0f4d725ec132b77e92dd46cd9712e660de1b13181ed238.
//
// Solidity: event PublicKeyUnregistered(address indexed voter, bytes32 indexed part1, bytes32 indexed part2)
func (_EntityManager *EntityManagerFilterer) FilterPublicKeyUnregistered(opts *bind.FilterOpts, voter []common.Address, part1 [][32]byte, part2 [][32]byte) (*EntityManagerPublicKeyUnregisteredIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var part1Rule []interface{}
	for _, part1Item := range part1 {
		part1Rule = append(part1Rule, part1Item)
	}
	var part2Rule []interface{}
	for _, part2Item := range part2 {
		part2Rule = append(part2Rule, part2Item)
	}

	logs, sub, err := _EntityManager.contract.FilterLogs(opts, "PublicKeyUnregistered", voterRule, part1Rule, part2Rule)
	if err != nil {
		return nil, err
	}
	return &EntityManagerPublicKeyUnregisteredIterator{contract: _EntityManager.contract, event: "PublicKeyUnregistered", logs: logs, sub: sub}, nil
}

// WatchPublicKeyUnregistered is a free log subscription operation binding the contract event 0xc82d1b006dff077f4b0f4d725ec132b77e92dd46cd9712e660de1b13181ed238.
//
// Solidity: event PublicKeyUnregistered(address indexed voter, bytes32 indexed part1, bytes32 indexed part2)
func (_EntityManager *EntityManagerFilterer) WatchPublicKeyUnregistered(opts *bind.WatchOpts, sink chan<- *EntityManagerPublicKeyUnregistered, voter []common.Address, part1 [][32]byte, part2 [][32]byte) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var part1Rule []interface{}
	for _, part1Item := range part1 {
		part1Rule = append(part1Rule, part1Item)
	}
	var part2Rule []interface{}
	for _, part2Item := range part2 {
		part2Rule = append(part2Rule, part2Item)
	}

	logs, sub, err := _EntityManager.contract.WatchLogs(opts, "PublicKeyUnregistered", voterRule, part1Rule, part2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityManagerPublicKeyUnregistered)
				if err := _EntityManager.contract.UnpackLog(event, "PublicKeyUnregistered", log); err != nil {
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

// ParsePublicKeyUnregistered is a log parse operation binding the contract event 0xc82d1b006dff077f4b0f4d725ec132b77e92dd46cd9712e660de1b13181ed238.
//
// Solidity: event PublicKeyUnregistered(address indexed voter, bytes32 indexed part1, bytes32 indexed part2)
func (_EntityManager *EntityManagerFilterer) ParsePublicKeyUnregistered(log types.Log) (*EntityManagerPublicKeyUnregistered, error) {
	event := new(EntityManagerPublicKeyUnregistered)
	if err := _EntityManager.contract.UnpackLog(event, "PublicKeyUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityManagerSigningPolicyAddressProposedIterator is returned from FilterSigningPolicyAddressProposed and is used to iterate over the raw logs and unpacked data for SigningPolicyAddressProposed events raised by the EntityManager contract.
type EntityManagerSigningPolicyAddressProposedIterator struct {
	Event *EntityManagerSigningPolicyAddressProposed // Event containing the contract specifics and raw log

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
func (it *EntityManagerSigningPolicyAddressProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityManagerSigningPolicyAddressProposed)
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
		it.Event = new(EntityManagerSigningPolicyAddressProposed)
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
func (it *EntityManagerSigningPolicyAddressProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityManagerSigningPolicyAddressProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityManagerSigningPolicyAddressProposed represents a SigningPolicyAddressProposed event raised by the EntityManager contract.
type EntityManagerSigningPolicyAddressProposed struct {
	Voter                common.Address
	SigningPolicyAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSigningPolicyAddressProposed is a free log retrieval operation binding the contract event 0xcc50114c186aff0cb8d59f83a59f057e0514f1ea68b9b1c0b993f8a7cde815eb.
//
// Solidity: event SigningPolicyAddressProposed(address indexed voter, address indexed signingPolicyAddress)
func (_EntityManager *EntityManagerFilterer) FilterSigningPolicyAddressProposed(opts *bind.FilterOpts, voter []common.Address, signingPolicyAddress []common.Address) (*EntityManagerSigningPolicyAddressProposedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}

	logs, sub, err := _EntityManager.contract.FilterLogs(opts, "SigningPolicyAddressProposed", voterRule, signingPolicyAddressRule)
	if err != nil {
		return nil, err
	}
	return &EntityManagerSigningPolicyAddressProposedIterator{contract: _EntityManager.contract, event: "SigningPolicyAddressProposed", logs: logs, sub: sub}, nil
}

// WatchSigningPolicyAddressProposed is a free log subscription operation binding the contract event 0xcc50114c186aff0cb8d59f83a59f057e0514f1ea68b9b1c0b993f8a7cde815eb.
//
// Solidity: event SigningPolicyAddressProposed(address indexed voter, address indexed signingPolicyAddress)
func (_EntityManager *EntityManagerFilterer) WatchSigningPolicyAddressProposed(opts *bind.WatchOpts, sink chan<- *EntityManagerSigningPolicyAddressProposed, voter []common.Address, signingPolicyAddress []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}

	logs, sub, err := _EntityManager.contract.WatchLogs(opts, "SigningPolicyAddressProposed", voterRule, signingPolicyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityManagerSigningPolicyAddressProposed)
				if err := _EntityManager.contract.UnpackLog(event, "SigningPolicyAddressProposed", log); err != nil {
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

// ParseSigningPolicyAddressProposed is a log parse operation binding the contract event 0xcc50114c186aff0cb8d59f83a59f057e0514f1ea68b9b1c0b993f8a7cde815eb.
//
// Solidity: event SigningPolicyAddressProposed(address indexed voter, address indexed signingPolicyAddress)
func (_EntityManager *EntityManagerFilterer) ParseSigningPolicyAddressProposed(log types.Log) (*EntityManagerSigningPolicyAddressProposed, error) {
	event := new(EntityManagerSigningPolicyAddressProposed)
	if err := _EntityManager.contract.UnpackLog(event, "SigningPolicyAddressProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityManagerSigningPolicyAddressRegistrationConfirmedIterator is returned from FilterSigningPolicyAddressRegistrationConfirmed and is used to iterate over the raw logs and unpacked data for SigningPolicyAddressRegistrationConfirmed events raised by the EntityManager contract.
type EntityManagerSigningPolicyAddressRegistrationConfirmedIterator struct {
	Event *EntityManagerSigningPolicyAddressRegistrationConfirmed // Event containing the contract specifics and raw log

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
func (it *EntityManagerSigningPolicyAddressRegistrationConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityManagerSigningPolicyAddressRegistrationConfirmed)
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
		it.Event = new(EntityManagerSigningPolicyAddressRegistrationConfirmed)
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
func (it *EntityManagerSigningPolicyAddressRegistrationConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityManagerSigningPolicyAddressRegistrationConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityManagerSigningPolicyAddressRegistrationConfirmed represents a SigningPolicyAddressRegistrationConfirmed event raised by the EntityManager contract.
type EntityManagerSigningPolicyAddressRegistrationConfirmed struct {
	Voter                common.Address
	SigningPolicyAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSigningPolicyAddressRegistrationConfirmed is a free log retrieval operation binding the contract event 0x9048f894e2dc8c92dc12516c12b68f8240a4fa83c046943471f4e2ea8c6f2ba8.
//
// Solidity: event SigningPolicyAddressRegistrationConfirmed(address indexed voter, address indexed signingPolicyAddress)
func (_EntityManager *EntityManagerFilterer) FilterSigningPolicyAddressRegistrationConfirmed(opts *bind.FilterOpts, voter []common.Address, signingPolicyAddress []common.Address) (*EntityManagerSigningPolicyAddressRegistrationConfirmedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}

	logs, sub, err := _EntityManager.contract.FilterLogs(opts, "SigningPolicyAddressRegistrationConfirmed", voterRule, signingPolicyAddressRule)
	if err != nil {
		return nil, err
	}
	return &EntityManagerSigningPolicyAddressRegistrationConfirmedIterator{contract: _EntityManager.contract, event: "SigningPolicyAddressRegistrationConfirmed", logs: logs, sub: sub}, nil
}

// WatchSigningPolicyAddressRegistrationConfirmed is a free log subscription operation binding the contract event 0x9048f894e2dc8c92dc12516c12b68f8240a4fa83c046943471f4e2ea8c6f2ba8.
//
// Solidity: event SigningPolicyAddressRegistrationConfirmed(address indexed voter, address indexed signingPolicyAddress)
func (_EntityManager *EntityManagerFilterer) WatchSigningPolicyAddressRegistrationConfirmed(opts *bind.WatchOpts, sink chan<- *EntityManagerSigningPolicyAddressRegistrationConfirmed, voter []common.Address, signingPolicyAddress []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}

	logs, sub, err := _EntityManager.contract.WatchLogs(opts, "SigningPolicyAddressRegistrationConfirmed", voterRule, signingPolicyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityManagerSigningPolicyAddressRegistrationConfirmed)
				if err := _EntityManager.contract.UnpackLog(event, "SigningPolicyAddressRegistrationConfirmed", log); err != nil {
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

// ParseSigningPolicyAddressRegistrationConfirmed is a log parse operation binding the contract event 0x9048f894e2dc8c92dc12516c12b68f8240a4fa83c046943471f4e2ea8c6f2ba8.
//
// Solidity: event SigningPolicyAddressRegistrationConfirmed(address indexed voter, address indexed signingPolicyAddress)
func (_EntityManager *EntityManagerFilterer) ParseSigningPolicyAddressRegistrationConfirmed(log types.Log) (*EntityManagerSigningPolicyAddressRegistrationConfirmed, error) {
	event := new(EntityManagerSigningPolicyAddressRegistrationConfirmed)
	if err := _EntityManager.contract.UnpackLog(event, "SigningPolicyAddressRegistrationConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityManagerSubmitAddressProposedIterator is returned from FilterSubmitAddressProposed and is used to iterate over the raw logs and unpacked data for SubmitAddressProposed events raised by the EntityManager contract.
type EntityManagerSubmitAddressProposedIterator struct {
	Event *EntityManagerSubmitAddressProposed // Event containing the contract specifics and raw log

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
func (it *EntityManagerSubmitAddressProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityManagerSubmitAddressProposed)
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
		it.Event = new(EntityManagerSubmitAddressProposed)
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
func (it *EntityManagerSubmitAddressProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityManagerSubmitAddressProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityManagerSubmitAddressProposed represents a SubmitAddressProposed event raised by the EntityManager contract.
type EntityManagerSubmitAddressProposed struct {
	Voter         common.Address
	SubmitAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmitAddressProposed is a free log retrieval operation binding the contract event 0x6a2d1988f57bc4ea8524ed144eabaac3bd34483c8e45307a895418c99474e9aa.
//
// Solidity: event SubmitAddressProposed(address indexed voter, address indexed submitAddress)
func (_EntityManager *EntityManagerFilterer) FilterSubmitAddressProposed(opts *bind.FilterOpts, voter []common.Address, submitAddress []common.Address) (*EntityManagerSubmitAddressProposedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var submitAddressRule []interface{}
	for _, submitAddressItem := range submitAddress {
		submitAddressRule = append(submitAddressRule, submitAddressItem)
	}

	logs, sub, err := _EntityManager.contract.FilterLogs(opts, "SubmitAddressProposed", voterRule, submitAddressRule)
	if err != nil {
		return nil, err
	}
	return &EntityManagerSubmitAddressProposedIterator{contract: _EntityManager.contract, event: "SubmitAddressProposed", logs: logs, sub: sub}, nil
}

// WatchSubmitAddressProposed is a free log subscription operation binding the contract event 0x6a2d1988f57bc4ea8524ed144eabaac3bd34483c8e45307a895418c99474e9aa.
//
// Solidity: event SubmitAddressProposed(address indexed voter, address indexed submitAddress)
func (_EntityManager *EntityManagerFilterer) WatchSubmitAddressProposed(opts *bind.WatchOpts, sink chan<- *EntityManagerSubmitAddressProposed, voter []common.Address, submitAddress []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var submitAddressRule []interface{}
	for _, submitAddressItem := range submitAddress {
		submitAddressRule = append(submitAddressRule, submitAddressItem)
	}

	logs, sub, err := _EntityManager.contract.WatchLogs(opts, "SubmitAddressProposed", voterRule, submitAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityManagerSubmitAddressProposed)
				if err := _EntityManager.contract.UnpackLog(event, "SubmitAddressProposed", log); err != nil {
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

// ParseSubmitAddressProposed is a log parse operation binding the contract event 0x6a2d1988f57bc4ea8524ed144eabaac3bd34483c8e45307a895418c99474e9aa.
//
// Solidity: event SubmitAddressProposed(address indexed voter, address indexed submitAddress)
func (_EntityManager *EntityManagerFilterer) ParseSubmitAddressProposed(log types.Log) (*EntityManagerSubmitAddressProposed, error) {
	event := new(EntityManagerSubmitAddressProposed)
	if err := _EntityManager.contract.UnpackLog(event, "SubmitAddressProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityManagerSubmitAddressRegistrationConfirmedIterator is returned from FilterSubmitAddressRegistrationConfirmed and is used to iterate over the raw logs and unpacked data for SubmitAddressRegistrationConfirmed events raised by the EntityManager contract.
type EntityManagerSubmitAddressRegistrationConfirmedIterator struct {
	Event *EntityManagerSubmitAddressRegistrationConfirmed // Event containing the contract specifics and raw log

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
func (it *EntityManagerSubmitAddressRegistrationConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityManagerSubmitAddressRegistrationConfirmed)
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
		it.Event = new(EntityManagerSubmitAddressRegistrationConfirmed)
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
func (it *EntityManagerSubmitAddressRegistrationConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityManagerSubmitAddressRegistrationConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityManagerSubmitAddressRegistrationConfirmed represents a SubmitAddressRegistrationConfirmed event raised by the EntityManager contract.
type EntityManagerSubmitAddressRegistrationConfirmed struct {
	Voter         common.Address
	SubmitAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmitAddressRegistrationConfirmed is a free log retrieval operation binding the contract event 0x82e3417c0fa3f7348ea9fcadb3c243c43ac30b003c0e1749990612a8d10a9674.
//
// Solidity: event SubmitAddressRegistrationConfirmed(address indexed voter, address indexed submitAddress)
func (_EntityManager *EntityManagerFilterer) FilterSubmitAddressRegistrationConfirmed(opts *bind.FilterOpts, voter []common.Address, submitAddress []common.Address) (*EntityManagerSubmitAddressRegistrationConfirmedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var submitAddressRule []interface{}
	for _, submitAddressItem := range submitAddress {
		submitAddressRule = append(submitAddressRule, submitAddressItem)
	}

	logs, sub, err := _EntityManager.contract.FilterLogs(opts, "SubmitAddressRegistrationConfirmed", voterRule, submitAddressRule)
	if err != nil {
		return nil, err
	}
	return &EntityManagerSubmitAddressRegistrationConfirmedIterator{contract: _EntityManager.contract, event: "SubmitAddressRegistrationConfirmed", logs: logs, sub: sub}, nil
}

// WatchSubmitAddressRegistrationConfirmed is a free log subscription operation binding the contract event 0x82e3417c0fa3f7348ea9fcadb3c243c43ac30b003c0e1749990612a8d10a9674.
//
// Solidity: event SubmitAddressRegistrationConfirmed(address indexed voter, address indexed submitAddress)
func (_EntityManager *EntityManagerFilterer) WatchSubmitAddressRegistrationConfirmed(opts *bind.WatchOpts, sink chan<- *EntityManagerSubmitAddressRegistrationConfirmed, voter []common.Address, submitAddress []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var submitAddressRule []interface{}
	for _, submitAddressItem := range submitAddress {
		submitAddressRule = append(submitAddressRule, submitAddressItem)
	}

	logs, sub, err := _EntityManager.contract.WatchLogs(opts, "SubmitAddressRegistrationConfirmed", voterRule, submitAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityManagerSubmitAddressRegistrationConfirmed)
				if err := _EntityManager.contract.UnpackLog(event, "SubmitAddressRegistrationConfirmed", log); err != nil {
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

// ParseSubmitAddressRegistrationConfirmed is a log parse operation binding the contract event 0x82e3417c0fa3f7348ea9fcadb3c243c43ac30b003c0e1749990612a8d10a9674.
//
// Solidity: event SubmitAddressRegistrationConfirmed(address indexed voter, address indexed submitAddress)
func (_EntityManager *EntityManagerFilterer) ParseSubmitAddressRegistrationConfirmed(log types.Log) (*EntityManagerSubmitAddressRegistrationConfirmed, error) {
	event := new(EntityManagerSubmitAddressRegistrationConfirmed)
	if err := _EntityManager.contract.UnpackLog(event, "SubmitAddressRegistrationConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityManagerSubmitSignaturesAddressProposedIterator is returned from FilterSubmitSignaturesAddressProposed and is used to iterate over the raw logs and unpacked data for SubmitSignaturesAddressProposed events raised by the EntityManager contract.
type EntityManagerSubmitSignaturesAddressProposedIterator struct {
	Event *EntityManagerSubmitSignaturesAddressProposed // Event containing the contract specifics and raw log

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
func (it *EntityManagerSubmitSignaturesAddressProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityManagerSubmitSignaturesAddressProposed)
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
		it.Event = new(EntityManagerSubmitSignaturesAddressProposed)
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
func (it *EntityManagerSubmitSignaturesAddressProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityManagerSubmitSignaturesAddressProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityManagerSubmitSignaturesAddressProposed represents a SubmitSignaturesAddressProposed event raised by the EntityManager contract.
type EntityManagerSubmitSignaturesAddressProposed struct {
	Voter                   common.Address
	SubmitSignaturesAddress common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSubmitSignaturesAddressProposed is a free log retrieval operation binding the contract event 0x1362ec372c84a894b58be817f44d2d1b2f9cb444493e30bf344d31c5f6894cc1.
//
// Solidity: event SubmitSignaturesAddressProposed(address indexed voter, address indexed submitSignaturesAddress)
func (_EntityManager *EntityManagerFilterer) FilterSubmitSignaturesAddressProposed(opts *bind.FilterOpts, voter []common.Address, submitSignaturesAddress []common.Address) (*EntityManagerSubmitSignaturesAddressProposedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var submitSignaturesAddressRule []interface{}
	for _, submitSignaturesAddressItem := range submitSignaturesAddress {
		submitSignaturesAddressRule = append(submitSignaturesAddressRule, submitSignaturesAddressItem)
	}

	logs, sub, err := _EntityManager.contract.FilterLogs(opts, "SubmitSignaturesAddressProposed", voterRule, submitSignaturesAddressRule)
	if err != nil {
		return nil, err
	}
	return &EntityManagerSubmitSignaturesAddressProposedIterator{contract: _EntityManager.contract, event: "SubmitSignaturesAddressProposed", logs: logs, sub: sub}, nil
}

// WatchSubmitSignaturesAddressProposed is a free log subscription operation binding the contract event 0x1362ec372c84a894b58be817f44d2d1b2f9cb444493e30bf344d31c5f6894cc1.
//
// Solidity: event SubmitSignaturesAddressProposed(address indexed voter, address indexed submitSignaturesAddress)
func (_EntityManager *EntityManagerFilterer) WatchSubmitSignaturesAddressProposed(opts *bind.WatchOpts, sink chan<- *EntityManagerSubmitSignaturesAddressProposed, voter []common.Address, submitSignaturesAddress []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var submitSignaturesAddressRule []interface{}
	for _, submitSignaturesAddressItem := range submitSignaturesAddress {
		submitSignaturesAddressRule = append(submitSignaturesAddressRule, submitSignaturesAddressItem)
	}

	logs, sub, err := _EntityManager.contract.WatchLogs(opts, "SubmitSignaturesAddressProposed", voterRule, submitSignaturesAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityManagerSubmitSignaturesAddressProposed)
				if err := _EntityManager.contract.UnpackLog(event, "SubmitSignaturesAddressProposed", log); err != nil {
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

// ParseSubmitSignaturesAddressProposed is a log parse operation binding the contract event 0x1362ec372c84a894b58be817f44d2d1b2f9cb444493e30bf344d31c5f6894cc1.
//
// Solidity: event SubmitSignaturesAddressProposed(address indexed voter, address indexed submitSignaturesAddress)
func (_EntityManager *EntityManagerFilterer) ParseSubmitSignaturesAddressProposed(log types.Log) (*EntityManagerSubmitSignaturesAddressProposed, error) {
	event := new(EntityManagerSubmitSignaturesAddressProposed)
	if err := _EntityManager.contract.UnpackLog(event, "SubmitSignaturesAddressProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EntityManagerSubmitSignaturesAddressRegistrationConfirmedIterator is returned from FilterSubmitSignaturesAddressRegistrationConfirmed and is used to iterate over the raw logs and unpacked data for SubmitSignaturesAddressRegistrationConfirmed events raised by the EntityManager contract.
type EntityManagerSubmitSignaturesAddressRegistrationConfirmedIterator struct {
	Event *EntityManagerSubmitSignaturesAddressRegistrationConfirmed // Event containing the contract specifics and raw log

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
func (it *EntityManagerSubmitSignaturesAddressRegistrationConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EntityManagerSubmitSignaturesAddressRegistrationConfirmed)
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
		it.Event = new(EntityManagerSubmitSignaturesAddressRegistrationConfirmed)
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
func (it *EntityManagerSubmitSignaturesAddressRegistrationConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EntityManagerSubmitSignaturesAddressRegistrationConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EntityManagerSubmitSignaturesAddressRegistrationConfirmed represents a SubmitSignaturesAddressRegistrationConfirmed event raised by the EntityManager contract.
type EntityManagerSubmitSignaturesAddressRegistrationConfirmed struct {
	Voter                   common.Address
	SubmitSignaturesAddress common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterSubmitSignaturesAddressRegistrationConfirmed is a free log retrieval operation binding the contract event 0x19542a5b3b9df6306e5277bdfd5865021928d36f752509fda85bc749378f4a12.
//
// Solidity: event SubmitSignaturesAddressRegistrationConfirmed(address indexed voter, address indexed submitSignaturesAddress)
func (_EntityManager *EntityManagerFilterer) FilterSubmitSignaturesAddressRegistrationConfirmed(opts *bind.FilterOpts, voter []common.Address, submitSignaturesAddress []common.Address) (*EntityManagerSubmitSignaturesAddressRegistrationConfirmedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var submitSignaturesAddressRule []interface{}
	for _, submitSignaturesAddressItem := range submitSignaturesAddress {
		submitSignaturesAddressRule = append(submitSignaturesAddressRule, submitSignaturesAddressItem)
	}

	logs, sub, err := _EntityManager.contract.FilterLogs(opts, "SubmitSignaturesAddressRegistrationConfirmed", voterRule, submitSignaturesAddressRule)
	if err != nil {
		return nil, err
	}
	return &EntityManagerSubmitSignaturesAddressRegistrationConfirmedIterator{contract: _EntityManager.contract, event: "SubmitSignaturesAddressRegistrationConfirmed", logs: logs, sub: sub}, nil
}

// WatchSubmitSignaturesAddressRegistrationConfirmed is a free log subscription operation binding the contract event 0x19542a5b3b9df6306e5277bdfd5865021928d36f752509fda85bc749378f4a12.
//
// Solidity: event SubmitSignaturesAddressRegistrationConfirmed(address indexed voter, address indexed submitSignaturesAddress)
func (_EntityManager *EntityManagerFilterer) WatchSubmitSignaturesAddressRegistrationConfirmed(opts *bind.WatchOpts, sink chan<- *EntityManagerSubmitSignaturesAddressRegistrationConfirmed, voter []common.Address, submitSignaturesAddress []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}
	var submitSignaturesAddressRule []interface{}
	for _, submitSignaturesAddressItem := range submitSignaturesAddress {
		submitSignaturesAddressRule = append(submitSignaturesAddressRule, submitSignaturesAddressItem)
	}

	logs, sub, err := _EntityManager.contract.WatchLogs(opts, "SubmitSignaturesAddressRegistrationConfirmed", voterRule, submitSignaturesAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EntityManagerSubmitSignaturesAddressRegistrationConfirmed)
				if err := _EntityManager.contract.UnpackLog(event, "SubmitSignaturesAddressRegistrationConfirmed", log); err != nil {
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

// ParseSubmitSignaturesAddressRegistrationConfirmed is a log parse operation binding the contract event 0x19542a5b3b9df6306e5277bdfd5865021928d36f752509fda85bc749378f4a12.
//
// Solidity: event SubmitSignaturesAddressRegistrationConfirmed(address indexed voter, address indexed submitSignaturesAddress)
func (_EntityManager *EntityManagerFilterer) ParseSubmitSignaturesAddressRegistrationConfirmed(log types.Log) (*EntityManagerSubmitSignaturesAddressRegistrationConfirmed, error) {
	event := new(EntityManagerSubmitSignaturesAddressRegistrationConfirmed)
	if err := _EntityManager.contract.UnpackLog(event, "SubmitSignaturesAddressRegistrationConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
