// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package extensiongovernance

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

// ExtensionGovernanceMetaData contains all meta data concerning the ExtensionGovernance contract.
var ExtensionGovernanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSigners\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"governanceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"signersThreshold\",\"type\":\"uint64\"}],\"name\":\"NewTeeGovernanceSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getLatestTeeGovernance\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_signersThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getLatestTeeGovernanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"getTeeGovernance\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_signersThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"getTeeGovernanceThreshold\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"isGovernanceHashValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"isTeeGovernanceSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_signersThreshold\",\"type\":\"uint64\"}],\"name\":\"setNewTeeGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ExtensionGovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use ExtensionGovernanceMetaData.ABI instead.
var ExtensionGovernanceABI = ExtensionGovernanceMetaData.ABI

// ExtensionGovernance is an auto generated Go binding around an Ethereum contract.
type ExtensionGovernance struct {
	ExtensionGovernanceCaller     // Read-only binding to the contract
	ExtensionGovernanceTransactor // Write-only binding to the contract
	ExtensionGovernanceFilterer   // Log filterer for contract events
}

// ExtensionGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtensionGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtensionGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtensionGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtensionGovernanceSession struct {
	Contract     *ExtensionGovernance // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ExtensionGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtensionGovernanceCallerSession struct {
	Contract *ExtensionGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ExtensionGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtensionGovernanceTransactorSession struct {
	Contract     *ExtensionGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ExtensionGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtensionGovernanceRaw struct {
	Contract *ExtensionGovernance // Generic contract binding to access the raw methods on
}

// ExtensionGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtensionGovernanceCallerRaw struct {
	Contract *ExtensionGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// ExtensionGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtensionGovernanceTransactorRaw struct {
	Contract *ExtensionGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtensionGovernance creates a new instance of ExtensionGovernance, bound to a specific deployed contract.
func NewExtensionGovernance(address common.Address, backend bind.ContractBackend) (*ExtensionGovernance, error) {
	contract, err := bindExtensionGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExtensionGovernance{ExtensionGovernanceCaller: ExtensionGovernanceCaller{contract: contract}, ExtensionGovernanceTransactor: ExtensionGovernanceTransactor{contract: contract}, ExtensionGovernanceFilterer: ExtensionGovernanceFilterer{contract: contract}}, nil
}

// NewExtensionGovernanceCaller creates a new read-only instance of ExtensionGovernance, bound to a specific deployed contract.
func NewExtensionGovernanceCaller(address common.Address, caller bind.ContractCaller) (*ExtensionGovernanceCaller, error) {
	contract, err := bindExtensionGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtensionGovernanceCaller{contract: contract}, nil
}

// NewExtensionGovernanceTransactor creates a new write-only instance of ExtensionGovernance, bound to a specific deployed contract.
func NewExtensionGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtensionGovernanceTransactor, error) {
	contract, err := bindExtensionGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtensionGovernanceTransactor{contract: contract}, nil
}

// NewExtensionGovernanceFilterer creates a new log filterer instance of ExtensionGovernance, bound to a specific deployed contract.
func NewExtensionGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtensionGovernanceFilterer, error) {
	contract, err := bindExtensionGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtensionGovernanceFilterer{contract: contract}, nil
}

// bindExtensionGovernance binds a generic wrapper to an already deployed contract.
func bindExtensionGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExtensionGovernanceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtensionGovernance *ExtensionGovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtensionGovernance.Contract.ExtensionGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtensionGovernance *ExtensionGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtensionGovernance.Contract.ExtensionGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtensionGovernance *ExtensionGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtensionGovernance.Contract.ExtensionGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtensionGovernance *ExtensionGovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtensionGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtensionGovernance *ExtensionGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtensionGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtensionGovernance *ExtensionGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtensionGovernance.Contract.contract.Transact(opts, method, params...)
}

// GetLatestTeeGovernance is a free data retrieval call binding the contract method 0x3e7bd5d4.
//
// Solidity: function getLatestTeeGovernance(uint256 _extensionId) view returns(address[] _signers, uint64 _signersThreshold)
func (_ExtensionGovernance *ExtensionGovernanceCaller) GetLatestTeeGovernance(opts *bind.CallOpts, _extensionId *big.Int) (struct {
	Signers          []common.Address
	SignersThreshold uint64
}, error) {
	var out []interface{}
	err := _ExtensionGovernance.contract.Call(opts, &out, "getLatestTeeGovernance", _extensionId)

	outstruct := new(struct {
		Signers          []common.Address
		SignersThreshold uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Signers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.SignersThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetLatestTeeGovernance is a free data retrieval call binding the contract method 0x3e7bd5d4.
//
// Solidity: function getLatestTeeGovernance(uint256 _extensionId) view returns(address[] _signers, uint64 _signersThreshold)
func (_ExtensionGovernance *ExtensionGovernanceSession) GetLatestTeeGovernance(_extensionId *big.Int) (struct {
	Signers          []common.Address
	SignersThreshold uint64
}, error) {
	return _ExtensionGovernance.Contract.GetLatestTeeGovernance(&_ExtensionGovernance.CallOpts, _extensionId)
}

// GetLatestTeeGovernance is a free data retrieval call binding the contract method 0x3e7bd5d4.
//
// Solidity: function getLatestTeeGovernance(uint256 _extensionId) view returns(address[] _signers, uint64 _signersThreshold)
func (_ExtensionGovernance *ExtensionGovernanceCallerSession) GetLatestTeeGovernance(_extensionId *big.Int) (struct {
	Signers          []common.Address
	SignersThreshold uint64
}, error) {
	return _ExtensionGovernance.Contract.GetLatestTeeGovernance(&_ExtensionGovernance.CallOpts, _extensionId)
}

// GetLatestTeeGovernanceHash is a free data retrieval call binding the contract method 0x555f9961.
//
// Solidity: function getLatestTeeGovernanceHash(uint256 _extensionId) view returns(bytes32)
func (_ExtensionGovernance *ExtensionGovernanceCaller) GetLatestTeeGovernanceHash(opts *bind.CallOpts, _extensionId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ExtensionGovernance.contract.Call(opts, &out, "getLatestTeeGovernanceHash", _extensionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLatestTeeGovernanceHash is a free data retrieval call binding the contract method 0x555f9961.
//
// Solidity: function getLatestTeeGovernanceHash(uint256 _extensionId) view returns(bytes32)
func (_ExtensionGovernance *ExtensionGovernanceSession) GetLatestTeeGovernanceHash(_extensionId *big.Int) ([32]byte, error) {
	return _ExtensionGovernance.Contract.GetLatestTeeGovernanceHash(&_ExtensionGovernance.CallOpts, _extensionId)
}

// GetLatestTeeGovernanceHash is a free data retrieval call binding the contract method 0x555f9961.
//
// Solidity: function getLatestTeeGovernanceHash(uint256 _extensionId) view returns(bytes32)
func (_ExtensionGovernance *ExtensionGovernanceCallerSession) GetLatestTeeGovernanceHash(_extensionId *big.Int) ([32]byte, error) {
	return _ExtensionGovernance.Contract.GetLatestTeeGovernanceHash(&_ExtensionGovernance.CallOpts, _extensionId)
}

// GetTeeGovernance is a free data retrieval call binding the contract method 0x090f2bfb.
//
// Solidity: function getTeeGovernance(uint256 _extensionId, bytes32 _governanceHash) view returns(address[] _signers, uint64 _signersThreshold)
func (_ExtensionGovernance *ExtensionGovernanceCaller) GetTeeGovernance(opts *bind.CallOpts, _extensionId *big.Int, _governanceHash [32]byte) (struct {
	Signers          []common.Address
	SignersThreshold uint64
}, error) {
	var out []interface{}
	err := _ExtensionGovernance.contract.Call(opts, &out, "getTeeGovernance", _extensionId, _governanceHash)

	outstruct := new(struct {
		Signers          []common.Address
		SignersThreshold uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Signers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.SignersThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetTeeGovernance is a free data retrieval call binding the contract method 0x090f2bfb.
//
// Solidity: function getTeeGovernance(uint256 _extensionId, bytes32 _governanceHash) view returns(address[] _signers, uint64 _signersThreshold)
func (_ExtensionGovernance *ExtensionGovernanceSession) GetTeeGovernance(_extensionId *big.Int, _governanceHash [32]byte) (struct {
	Signers          []common.Address
	SignersThreshold uint64
}, error) {
	return _ExtensionGovernance.Contract.GetTeeGovernance(&_ExtensionGovernance.CallOpts, _extensionId, _governanceHash)
}

// GetTeeGovernance is a free data retrieval call binding the contract method 0x090f2bfb.
//
// Solidity: function getTeeGovernance(uint256 _extensionId, bytes32 _governanceHash) view returns(address[] _signers, uint64 _signersThreshold)
func (_ExtensionGovernance *ExtensionGovernanceCallerSession) GetTeeGovernance(_extensionId *big.Int, _governanceHash [32]byte) (struct {
	Signers          []common.Address
	SignersThreshold uint64
}, error) {
	return _ExtensionGovernance.Contract.GetTeeGovernance(&_ExtensionGovernance.CallOpts, _extensionId, _governanceHash)
}

// GetTeeGovernanceThreshold is a free data retrieval call binding the contract method 0x3b13db2b.
//
// Solidity: function getTeeGovernanceThreshold(uint256 _extensionId, bytes32 _governanceHash) view returns(uint64)
func (_ExtensionGovernance *ExtensionGovernanceCaller) GetTeeGovernanceThreshold(opts *bind.CallOpts, _extensionId *big.Int, _governanceHash [32]byte) (uint64, error) {
	var out []interface{}
	err := _ExtensionGovernance.contract.Call(opts, &out, "getTeeGovernanceThreshold", _extensionId, _governanceHash)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTeeGovernanceThreshold is a free data retrieval call binding the contract method 0x3b13db2b.
//
// Solidity: function getTeeGovernanceThreshold(uint256 _extensionId, bytes32 _governanceHash) view returns(uint64)
func (_ExtensionGovernance *ExtensionGovernanceSession) GetTeeGovernanceThreshold(_extensionId *big.Int, _governanceHash [32]byte) (uint64, error) {
	return _ExtensionGovernance.Contract.GetTeeGovernanceThreshold(&_ExtensionGovernance.CallOpts, _extensionId, _governanceHash)
}

// GetTeeGovernanceThreshold is a free data retrieval call binding the contract method 0x3b13db2b.
//
// Solidity: function getTeeGovernanceThreshold(uint256 _extensionId, bytes32 _governanceHash) view returns(uint64)
func (_ExtensionGovernance *ExtensionGovernanceCallerSession) GetTeeGovernanceThreshold(_extensionId *big.Int, _governanceHash [32]byte) (uint64, error) {
	return _ExtensionGovernance.Contract.GetTeeGovernanceThreshold(&_ExtensionGovernance.CallOpts, _extensionId, _governanceHash)
}

// IsGovernanceHashValid is a free data retrieval call binding the contract method 0x66d98ef8.
//
// Solidity: function isGovernanceHashValid(uint256 _extensionId, bytes32 _governanceHash) view returns(bool)
func (_ExtensionGovernance *ExtensionGovernanceCaller) IsGovernanceHashValid(opts *bind.CallOpts, _extensionId *big.Int, _governanceHash [32]byte) (bool, error) {
	var out []interface{}
	err := _ExtensionGovernance.contract.Call(opts, &out, "isGovernanceHashValid", _extensionId, _governanceHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernanceHashValid is a free data retrieval call binding the contract method 0x66d98ef8.
//
// Solidity: function isGovernanceHashValid(uint256 _extensionId, bytes32 _governanceHash) view returns(bool)
func (_ExtensionGovernance *ExtensionGovernanceSession) IsGovernanceHashValid(_extensionId *big.Int, _governanceHash [32]byte) (bool, error) {
	return _ExtensionGovernance.Contract.IsGovernanceHashValid(&_ExtensionGovernance.CallOpts, _extensionId, _governanceHash)
}

// IsGovernanceHashValid is a free data retrieval call binding the contract method 0x66d98ef8.
//
// Solidity: function isGovernanceHashValid(uint256 _extensionId, bytes32 _governanceHash) view returns(bool)
func (_ExtensionGovernance *ExtensionGovernanceCallerSession) IsGovernanceHashValid(_extensionId *big.Int, _governanceHash [32]byte) (bool, error) {
	return _ExtensionGovernance.Contract.IsGovernanceHashValid(&_ExtensionGovernance.CallOpts, _extensionId, _governanceHash)
}

// IsTeeGovernanceSigner is a free data retrieval call binding the contract method 0xd8784372.
//
// Solidity: function isTeeGovernanceSigner(uint256 _extensionId, bytes32 _governanceHash, address _signer) view returns(bool)
func (_ExtensionGovernance *ExtensionGovernanceCaller) IsTeeGovernanceSigner(opts *bind.CallOpts, _extensionId *big.Int, _governanceHash [32]byte, _signer common.Address) (bool, error) {
	var out []interface{}
	err := _ExtensionGovernance.contract.Call(opts, &out, "isTeeGovernanceSigner", _extensionId, _governanceHash, _signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeeGovernanceSigner is a free data retrieval call binding the contract method 0xd8784372.
//
// Solidity: function isTeeGovernanceSigner(uint256 _extensionId, bytes32 _governanceHash, address _signer) view returns(bool)
func (_ExtensionGovernance *ExtensionGovernanceSession) IsTeeGovernanceSigner(_extensionId *big.Int, _governanceHash [32]byte, _signer common.Address) (bool, error) {
	return _ExtensionGovernance.Contract.IsTeeGovernanceSigner(&_ExtensionGovernance.CallOpts, _extensionId, _governanceHash, _signer)
}

// IsTeeGovernanceSigner is a free data retrieval call binding the contract method 0xd8784372.
//
// Solidity: function isTeeGovernanceSigner(uint256 _extensionId, bytes32 _governanceHash, address _signer) view returns(bool)
func (_ExtensionGovernance *ExtensionGovernanceCallerSession) IsTeeGovernanceSigner(_extensionId *big.Int, _governanceHash [32]byte, _signer common.Address) (bool, error) {
	return _ExtensionGovernance.Contract.IsTeeGovernanceSigner(&_ExtensionGovernance.CallOpts, _extensionId, _governanceHash, _signer)
}

// SetNewTeeGovernance is a paid mutator transaction binding the contract method 0x49e43748.
//
// Solidity: function setNewTeeGovernance(uint256 _extensionId, address[] _signers, uint64 _signersThreshold) returns()
func (_ExtensionGovernance *ExtensionGovernanceTransactor) SetNewTeeGovernance(opts *bind.TransactOpts, _extensionId *big.Int, _signers []common.Address, _signersThreshold uint64) (*types.Transaction, error) {
	return _ExtensionGovernance.contract.Transact(opts, "setNewTeeGovernance", _extensionId, _signers, _signersThreshold)
}

// SetNewTeeGovernance is a paid mutator transaction binding the contract method 0x49e43748.
//
// Solidity: function setNewTeeGovernance(uint256 _extensionId, address[] _signers, uint64 _signersThreshold) returns()
func (_ExtensionGovernance *ExtensionGovernanceSession) SetNewTeeGovernance(_extensionId *big.Int, _signers []common.Address, _signersThreshold uint64) (*types.Transaction, error) {
	return _ExtensionGovernance.Contract.SetNewTeeGovernance(&_ExtensionGovernance.TransactOpts, _extensionId, _signers, _signersThreshold)
}

// SetNewTeeGovernance is a paid mutator transaction binding the contract method 0x49e43748.
//
// Solidity: function setNewTeeGovernance(uint256 _extensionId, address[] _signers, uint64 _signersThreshold) returns()
func (_ExtensionGovernance *ExtensionGovernanceTransactorSession) SetNewTeeGovernance(_extensionId *big.Int, _signers []common.Address, _signersThreshold uint64) (*types.Transaction, error) {
	return _ExtensionGovernance.Contract.SetNewTeeGovernance(&_ExtensionGovernance.TransactOpts, _extensionId, _signers, _signersThreshold)
}

// ExtensionGovernanceNewTeeGovernanceSetIterator is returned from FilterNewTeeGovernanceSet and is used to iterate over the raw logs and unpacked data for NewTeeGovernanceSet events raised by the ExtensionGovernance contract.
type ExtensionGovernanceNewTeeGovernanceSetIterator struct {
	Event *ExtensionGovernanceNewTeeGovernanceSet // Event containing the contract specifics and raw log

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
func (it *ExtensionGovernanceNewTeeGovernanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionGovernanceNewTeeGovernanceSet)
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
		it.Event = new(ExtensionGovernanceNewTeeGovernanceSet)
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
func (it *ExtensionGovernanceNewTeeGovernanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionGovernanceNewTeeGovernanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionGovernanceNewTeeGovernanceSet represents a NewTeeGovernanceSet event raised by the ExtensionGovernance contract.
type ExtensionGovernanceNewTeeGovernanceSet struct {
	ExtensionId      *big.Int
	GovernanceHash   [32]byte
	Signers          []common.Address
	SignersThreshold uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewTeeGovernanceSet is a free log retrieval operation binding the contract event 0x4f69db036ae5a154d908999072bb0414178748284f1b9c3a8449b9ee9e3932bd.
//
// Solidity: event NewTeeGovernanceSet(uint256 indexed extensionId, bytes32 indexed governanceHash, address[] signers, uint64 signersThreshold)
func (_ExtensionGovernance *ExtensionGovernanceFilterer) FilterNewTeeGovernanceSet(opts *bind.FilterOpts, extensionId []*big.Int, governanceHash [][32]byte) (*ExtensionGovernanceNewTeeGovernanceSetIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var governanceHashRule []interface{}
	for _, governanceHashItem := range governanceHash {
		governanceHashRule = append(governanceHashRule, governanceHashItem)
	}

	logs, sub, err := _ExtensionGovernance.contract.FilterLogs(opts, "NewTeeGovernanceSet", extensionIdRule, governanceHashRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionGovernanceNewTeeGovernanceSetIterator{contract: _ExtensionGovernance.contract, event: "NewTeeGovernanceSet", logs: logs, sub: sub}, nil
}

// WatchNewTeeGovernanceSet is a free log subscription operation binding the contract event 0x4f69db036ae5a154d908999072bb0414178748284f1b9c3a8449b9ee9e3932bd.
//
// Solidity: event NewTeeGovernanceSet(uint256 indexed extensionId, bytes32 indexed governanceHash, address[] signers, uint64 signersThreshold)
func (_ExtensionGovernance *ExtensionGovernanceFilterer) WatchNewTeeGovernanceSet(opts *bind.WatchOpts, sink chan<- *ExtensionGovernanceNewTeeGovernanceSet, extensionId []*big.Int, governanceHash [][32]byte) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var governanceHashRule []interface{}
	for _, governanceHashItem := range governanceHash {
		governanceHashRule = append(governanceHashRule, governanceHashItem)
	}

	logs, sub, err := _ExtensionGovernance.contract.WatchLogs(opts, "NewTeeGovernanceSet", extensionIdRule, governanceHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionGovernanceNewTeeGovernanceSet)
				if err := _ExtensionGovernance.contract.UnpackLog(event, "NewTeeGovernanceSet", log); err != nil {
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

// ParseNewTeeGovernanceSet is a log parse operation binding the contract event 0x4f69db036ae5a154d908999072bb0414178748284f1b9c3a8449b9ee9e3932bd.
//
// Solidity: event NewTeeGovernanceSet(uint256 indexed extensionId, bytes32 indexed governanceHash, address[] signers, uint64 signersThreshold)
func (_ExtensionGovernance *ExtensionGovernanceFilterer) ParseNewTeeGovernanceSet(log types.Log) (*ExtensionGovernanceNewTeeGovernanceSet, error) {
	event := new(ExtensionGovernanceNewTeeGovernanceSet)
	if err := _ExtensionGovernance.contract.UnpackLog(event, "NewTeeGovernanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
