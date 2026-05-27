// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package extensionpausing

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

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// ExtensionPausingMetaData contains all meta data concerning the ExtensionPausing contract.
var ExtensionPausingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"AlreadySigned\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateGovernanceHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoGovernanceHashes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"NotASigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pausingAddress\",\"type\":\"address\"}],\"name\":\"PausingAddressAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausingAddressesNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"governanceHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"pausingAddresses\",\"type\":\"address[]\"}],\"name\":\"NewPausingAddressesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"governanceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"NewPausingAddressesSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"governanceHash\",\"type\":\"bytes32\"}],\"name\":\"TeePausingAddressesThresholdMet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getLatestTeePausingAddresses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_pausingAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_governanceHashes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[][]\",\"name\":\"_signaturesPerHash\",\"type\":\"tuple[][]\"},{\"internalType\":\"bool[]\",\"name\":\"_thresholdMetPerHash\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getTeePausingAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_pausingAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_governanceHashes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[][]\",\"name\":\"_signaturesPerHash\",\"type\":\"tuple[][]\"},{\"internalType\":\"bool[]\",\"name\":\"_thresholdMetPerHash\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"hasSignedTeePausingAddresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_governanceHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_pausingAddresses\",\"type\":\"address[]\"}],\"name\":\"setTeePausingAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signTeePausingAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ExtensionPausingABI is the input ABI used to generate the binding from.
// Deprecated: Use ExtensionPausingMetaData.ABI instead.
var ExtensionPausingABI = ExtensionPausingMetaData.ABI

// ExtensionPausing is an auto generated Go binding around an Ethereum contract.
type ExtensionPausing struct {
	ExtensionPausingCaller     // Read-only binding to the contract
	ExtensionPausingTransactor // Write-only binding to the contract
	ExtensionPausingFilterer   // Log filterer for contract events
}

// ExtensionPausingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtensionPausingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionPausingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtensionPausingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionPausingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtensionPausingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionPausingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtensionPausingSession struct {
	Contract     *ExtensionPausing // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExtensionPausingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtensionPausingCallerSession struct {
	Contract *ExtensionPausingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ExtensionPausingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtensionPausingTransactorSession struct {
	Contract     *ExtensionPausingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExtensionPausingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtensionPausingRaw struct {
	Contract *ExtensionPausing // Generic contract binding to access the raw methods on
}

// ExtensionPausingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtensionPausingCallerRaw struct {
	Contract *ExtensionPausingCaller // Generic read-only contract binding to access the raw methods on
}

// ExtensionPausingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtensionPausingTransactorRaw struct {
	Contract *ExtensionPausingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtensionPausing creates a new instance of ExtensionPausing, bound to a specific deployed contract.
func NewExtensionPausing(address common.Address, backend bind.ContractBackend) (*ExtensionPausing, error) {
	contract, err := bindExtensionPausing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExtensionPausing{ExtensionPausingCaller: ExtensionPausingCaller{contract: contract}, ExtensionPausingTransactor: ExtensionPausingTransactor{contract: contract}, ExtensionPausingFilterer: ExtensionPausingFilterer{contract: contract}}, nil
}

// NewExtensionPausingCaller creates a new read-only instance of ExtensionPausing, bound to a specific deployed contract.
func NewExtensionPausingCaller(address common.Address, caller bind.ContractCaller) (*ExtensionPausingCaller, error) {
	contract, err := bindExtensionPausing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtensionPausingCaller{contract: contract}, nil
}

// NewExtensionPausingTransactor creates a new write-only instance of ExtensionPausing, bound to a specific deployed contract.
func NewExtensionPausingTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtensionPausingTransactor, error) {
	contract, err := bindExtensionPausing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtensionPausingTransactor{contract: contract}, nil
}

// NewExtensionPausingFilterer creates a new log filterer instance of ExtensionPausing, bound to a specific deployed contract.
func NewExtensionPausingFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtensionPausingFilterer, error) {
	contract, err := bindExtensionPausing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtensionPausingFilterer{contract: contract}, nil
}

// bindExtensionPausing binds a generic wrapper to an already deployed contract.
func bindExtensionPausing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExtensionPausingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtensionPausing *ExtensionPausingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtensionPausing.Contract.ExtensionPausingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtensionPausing *ExtensionPausingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtensionPausing.Contract.ExtensionPausingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtensionPausing *ExtensionPausingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtensionPausing.Contract.ExtensionPausingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtensionPausing *ExtensionPausingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtensionPausing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtensionPausing *ExtensionPausingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtensionPausing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtensionPausing *ExtensionPausingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtensionPausing.Contract.contract.Transact(opts, method, params...)
}

// GetLatestTeePausingAddresses is a free data retrieval call binding the contract method 0xdbc0797d.
//
// Solidity: function getLatestTeePausingAddresses(uint256 _extensionId) view returns(uint256 _nonce, address[] _pausingAddresses, bytes32[] _governanceHashes, (uint8,bytes32,bytes32)[][] _signaturesPerHash, bool[] _thresholdMetPerHash)
func (_ExtensionPausing *ExtensionPausingCaller) GetLatestTeePausingAddresses(opts *bind.CallOpts, _extensionId *big.Int) (struct {
	Nonce               *big.Int
	PausingAddresses    []common.Address
	GovernanceHashes    [][32]byte
	SignaturesPerHash   [][]Signature
	ThresholdMetPerHash []bool
}, error) {
	var out []interface{}
	err := _ExtensionPausing.contract.Call(opts, &out, "getLatestTeePausingAddresses", _extensionId)

	outstruct := new(struct {
		Nonce               *big.Int
		PausingAddresses    []common.Address
		GovernanceHashes    [][32]byte
		SignaturesPerHash   [][]Signature
		ThresholdMetPerHash []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Nonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PausingAddresses = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.GovernanceHashes = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)
	outstruct.SignaturesPerHash = *abi.ConvertType(out[3], new([][]Signature)).(*[][]Signature)
	outstruct.ThresholdMetPerHash = *abi.ConvertType(out[4], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetLatestTeePausingAddresses is a free data retrieval call binding the contract method 0xdbc0797d.
//
// Solidity: function getLatestTeePausingAddresses(uint256 _extensionId) view returns(uint256 _nonce, address[] _pausingAddresses, bytes32[] _governanceHashes, (uint8,bytes32,bytes32)[][] _signaturesPerHash, bool[] _thresholdMetPerHash)
func (_ExtensionPausing *ExtensionPausingSession) GetLatestTeePausingAddresses(_extensionId *big.Int) (struct {
	Nonce               *big.Int
	PausingAddresses    []common.Address
	GovernanceHashes    [][32]byte
	SignaturesPerHash   [][]Signature
	ThresholdMetPerHash []bool
}, error) {
	return _ExtensionPausing.Contract.GetLatestTeePausingAddresses(&_ExtensionPausing.CallOpts, _extensionId)
}

// GetLatestTeePausingAddresses is a free data retrieval call binding the contract method 0xdbc0797d.
//
// Solidity: function getLatestTeePausingAddresses(uint256 _extensionId) view returns(uint256 _nonce, address[] _pausingAddresses, bytes32[] _governanceHashes, (uint8,bytes32,bytes32)[][] _signaturesPerHash, bool[] _thresholdMetPerHash)
func (_ExtensionPausing *ExtensionPausingCallerSession) GetLatestTeePausingAddresses(_extensionId *big.Int) (struct {
	Nonce               *big.Int
	PausingAddresses    []common.Address
	GovernanceHashes    [][32]byte
	SignaturesPerHash   [][]Signature
	ThresholdMetPerHash []bool
}, error) {
	return _ExtensionPausing.Contract.GetLatestTeePausingAddresses(&_ExtensionPausing.CallOpts, _extensionId)
}

// GetTeePausingAddresses is a free data retrieval call binding the contract method 0xaaee816a.
//
// Solidity: function getTeePausingAddresses(uint256 _extensionId, uint256 _nonce) view returns(address[] _pausingAddresses, bytes32[] _governanceHashes, (uint8,bytes32,bytes32)[][] _signaturesPerHash, bool[] _thresholdMetPerHash)
func (_ExtensionPausing *ExtensionPausingCaller) GetTeePausingAddresses(opts *bind.CallOpts, _extensionId *big.Int, _nonce *big.Int) (struct {
	PausingAddresses    []common.Address
	GovernanceHashes    [][32]byte
	SignaturesPerHash   [][]Signature
	ThresholdMetPerHash []bool
}, error) {
	var out []interface{}
	err := _ExtensionPausing.contract.Call(opts, &out, "getTeePausingAddresses", _extensionId, _nonce)

	outstruct := new(struct {
		PausingAddresses    []common.Address
		GovernanceHashes    [][32]byte
		SignaturesPerHash   [][]Signature
		ThresholdMetPerHash []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PausingAddresses = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.GovernanceHashes = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)
	outstruct.SignaturesPerHash = *abi.ConvertType(out[2], new([][]Signature)).(*[][]Signature)
	outstruct.ThresholdMetPerHash = *abi.ConvertType(out[3], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetTeePausingAddresses is a free data retrieval call binding the contract method 0xaaee816a.
//
// Solidity: function getTeePausingAddresses(uint256 _extensionId, uint256 _nonce) view returns(address[] _pausingAddresses, bytes32[] _governanceHashes, (uint8,bytes32,bytes32)[][] _signaturesPerHash, bool[] _thresholdMetPerHash)
func (_ExtensionPausing *ExtensionPausingSession) GetTeePausingAddresses(_extensionId *big.Int, _nonce *big.Int) (struct {
	PausingAddresses    []common.Address
	GovernanceHashes    [][32]byte
	SignaturesPerHash   [][]Signature
	ThresholdMetPerHash []bool
}, error) {
	return _ExtensionPausing.Contract.GetTeePausingAddresses(&_ExtensionPausing.CallOpts, _extensionId, _nonce)
}

// GetTeePausingAddresses is a free data retrieval call binding the contract method 0xaaee816a.
//
// Solidity: function getTeePausingAddresses(uint256 _extensionId, uint256 _nonce) view returns(address[] _pausingAddresses, bytes32[] _governanceHashes, (uint8,bytes32,bytes32)[][] _signaturesPerHash, bool[] _thresholdMetPerHash)
func (_ExtensionPausing *ExtensionPausingCallerSession) GetTeePausingAddresses(_extensionId *big.Int, _nonce *big.Int) (struct {
	PausingAddresses    []common.Address
	GovernanceHashes    [][32]byte
	SignaturesPerHash   [][]Signature
	ThresholdMetPerHash []bool
}, error) {
	return _ExtensionPausing.Contract.GetTeePausingAddresses(&_ExtensionPausing.CallOpts, _extensionId, _nonce)
}

// HasSignedTeePausingAddresses is a free data retrieval call binding the contract method 0x2030ec94.
//
// Solidity: function hasSignedTeePausingAddresses(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash, address _signer) view returns(bool)
func (_ExtensionPausing *ExtensionPausingCaller) HasSignedTeePausingAddresses(opts *bind.CallOpts, _extensionId *big.Int, _nonce *big.Int, _governanceHash [32]byte, _signer common.Address) (bool, error) {
	var out []interface{}
	err := _ExtensionPausing.contract.Call(opts, &out, "hasSignedTeePausingAddresses", _extensionId, _nonce, _governanceHash, _signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSignedTeePausingAddresses is a free data retrieval call binding the contract method 0x2030ec94.
//
// Solidity: function hasSignedTeePausingAddresses(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash, address _signer) view returns(bool)
func (_ExtensionPausing *ExtensionPausingSession) HasSignedTeePausingAddresses(_extensionId *big.Int, _nonce *big.Int, _governanceHash [32]byte, _signer common.Address) (bool, error) {
	return _ExtensionPausing.Contract.HasSignedTeePausingAddresses(&_ExtensionPausing.CallOpts, _extensionId, _nonce, _governanceHash, _signer)
}

// HasSignedTeePausingAddresses is a free data retrieval call binding the contract method 0x2030ec94.
//
// Solidity: function hasSignedTeePausingAddresses(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash, address _signer) view returns(bool)
func (_ExtensionPausing *ExtensionPausingCallerSession) HasSignedTeePausingAddresses(_extensionId *big.Int, _nonce *big.Int, _governanceHash [32]byte, _signer common.Address) (bool, error) {
	return _ExtensionPausing.Contract.HasSignedTeePausingAddresses(&_ExtensionPausing.CallOpts, _extensionId, _nonce, _governanceHash, _signer)
}

// SetTeePausingAddresses is a paid mutator transaction binding the contract method 0x9bad9f14.
//
// Solidity: function setTeePausingAddresses(uint256 _extensionId, bytes32[] _governanceHashes, address[] _pausingAddresses) returns()
func (_ExtensionPausing *ExtensionPausingTransactor) SetTeePausingAddresses(opts *bind.TransactOpts, _extensionId *big.Int, _governanceHashes [][32]byte, _pausingAddresses []common.Address) (*types.Transaction, error) {
	return _ExtensionPausing.contract.Transact(opts, "setTeePausingAddresses", _extensionId, _governanceHashes, _pausingAddresses)
}

// SetTeePausingAddresses is a paid mutator transaction binding the contract method 0x9bad9f14.
//
// Solidity: function setTeePausingAddresses(uint256 _extensionId, bytes32[] _governanceHashes, address[] _pausingAddresses) returns()
func (_ExtensionPausing *ExtensionPausingSession) SetTeePausingAddresses(_extensionId *big.Int, _governanceHashes [][32]byte, _pausingAddresses []common.Address) (*types.Transaction, error) {
	return _ExtensionPausing.Contract.SetTeePausingAddresses(&_ExtensionPausing.TransactOpts, _extensionId, _governanceHashes, _pausingAddresses)
}

// SetTeePausingAddresses is a paid mutator transaction binding the contract method 0x9bad9f14.
//
// Solidity: function setTeePausingAddresses(uint256 _extensionId, bytes32[] _governanceHashes, address[] _pausingAddresses) returns()
func (_ExtensionPausing *ExtensionPausingTransactorSession) SetTeePausingAddresses(_extensionId *big.Int, _governanceHashes [][32]byte, _pausingAddresses []common.Address) (*types.Transaction, error) {
	return _ExtensionPausing.Contract.SetTeePausingAddresses(&_ExtensionPausing.TransactOpts, _extensionId, _governanceHashes, _pausingAddresses)
}

// SignTeePausingAddresses is a paid mutator transaction binding the contract method 0x511cb613.
//
// Solidity: function signTeePausingAddresses(uint256 _extensionId, uint256 _nonce, (uint8,bytes32,bytes32) _signature) returns()
func (_ExtensionPausing *ExtensionPausingTransactor) SignTeePausingAddresses(opts *bind.TransactOpts, _extensionId *big.Int, _nonce *big.Int, _signature Signature) (*types.Transaction, error) {
	return _ExtensionPausing.contract.Transact(opts, "signTeePausingAddresses", _extensionId, _nonce, _signature)
}

// SignTeePausingAddresses is a paid mutator transaction binding the contract method 0x511cb613.
//
// Solidity: function signTeePausingAddresses(uint256 _extensionId, uint256 _nonce, (uint8,bytes32,bytes32) _signature) returns()
func (_ExtensionPausing *ExtensionPausingSession) SignTeePausingAddresses(_extensionId *big.Int, _nonce *big.Int, _signature Signature) (*types.Transaction, error) {
	return _ExtensionPausing.Contract.SignTeePausingAddresses(&_ExtensionPausing.TransactOpts, _extensionId, _nonce, _signature)
}

// SignTeePausingAddresses is a paid mutator transaction binding the contract method 0x511cb613.
//
// Solidity: function signTeePausingAddresses(uint256 _extensionId, uint256 _nonce, (uint8,bytes32,bytes32) _signature) returns()
func (_ExtensionPausing *ExtensionPausingTransactorSession) SignTeePausingAddresses(_extensionId *big.Int, _nonce *big.Int, _signature Signature) (*types.Transaction, error) {
	return _ExtensionPausing.Contract.SignTeePausingAddresses(&_ExtensionPausing.TransactOpts, _extensionId, _nonce, _signature)
}

// ExtensionPausingNewPausingAddressesSetIterator is returned from FilterNewPausingAddressesSet and is used to iterate over the raw logs and unpacked data for NewPausingAddressesSet events raised by the ExtensionPausing contract.
type ExtensionPausingNewPausingAddressesSetIterator struct {
	Event *ExtensionPausingNewPausingAddressesSet // Event containing the contract specifics and raw log

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
func (it *ExtensionPausingNewPausingAddressesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionPausingNewPausingAddressesSet)
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
		it.Event = new(ExtensionPausingNewPausingAddressesSet)
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
func (it *ExtensionPausingNewPausingAddressesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionPausingNewPausingAddressesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionPausingNewPausingAddressesSet represents a NewPausingAddressesSet event raised by the ExtensionPausing contract.
type ExtensionPausingNewPausingAddressesSet struct {
	ExtensionId      *big.Int
	Nonce            *big.Int
	GovernanceHashes [][32]byte
	PausingAddresses []common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPausingAddressesSet is a free log retrieval operation binding the contract event 0xb0050bcebd4aac2732f354f8b2d2c08b7de3f52963885e66b1dec95ec87df13d.
//
// Solidity: event NewPausingAddressesSet(uint256 indexed extensionId, uint256 indexed nonce, bytes32[] governanceHashes, address[] pausingAddresses)
func (_ExtensionPausing *ExtensionPausingFilterer) FilterNewPausingAddressesSet(opts *bind.FilterOpts, extensionId []*big.Int, nonce []*big.Int) (*ExtensionPausingNewPausingAddressesSetIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ExtensionPausing.contract.FilterLogs(opts, "NewPausingAddressesSet", extensionIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionPausingNewPausingAddressesSetIterator{contract: _ExtensionPausing.contract, event: "NewPausingAddressesSet", logs: logs, sub: sub}, nil
}

// WatchNewPausingAddressesSet is a free log subscription operation binding the contract event 0xb0050bcebd4aac2732f354f8b2d2c08b7de3f52963885e66b1dec95ec87df13d.
//
// Solidity: event NewPausingAddressesSet(uint256 indexed extensionId, uint256 indexed nonce, bytes32[] governanceHashes, address[] pausingAddresses)
func (_ExtensionPausing *ExtensionPausingFilterer) WatchNewPausingAddressesSet(opts *bind.WatchOpts, sink chan<- *ExtensionPausingNewPausingAddressesSet, extensionId []*big.Int, nonce []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _ExtensionPausing.contract.WatchLogs(opts, "NewPausingAddressesSet", extensionIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionPausingNewPausingAddressesSet)
				if err := _ExtensionPausing.contract.UnpackLog(event, "NewPausingAddressesSet", log); err != nil {
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

// ParseNewPausingAddressesSet is a log parse operation binding the contract event 0xb0050bcebd4aac2732f354f8b2d2c08b7de3f52963885e66b1dec95ec87df13d.
//
// Solidity: event NewPausingAddressesSet(uint256 indexed extensionId, uint256 indexed nonce, bytes32[] governanceHashes, address[] pausingAddresses)
func (_ExtensionPausing *ExtensionPausingFilterer) ParseNewPausingAddressesSet(log types.Log) (*ExtensionPausingNewPausingAddressesSet, error) {
	event := new(ExtensionPausingNewPausingAddressesSet)
	if err := _ExtensionPausing.contract.UnpackLog(event, "NewPausingAddressesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionPausingNewPausingAddressesSignedIterator is returned from FilterNewPausingAddressesSigned and is used to iterate over the raw logs and unpacked data for NewPausingAddressesSigned events raised by the ExtensionPausing contract.
type ExtensionPausingNewPausingAddressesSignedIterator struct {
	Event *ExtensionPausingNewPausingAddressesSigned // Event containing the contract specifics and raw log

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
func (it *ExtensionPausingNewPausingAddressesSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionPausingNewPausingAddressesSigned)
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
		it.Event = new(ExtensionPausingNewPausingAddressesSigned)
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
func (it *ExtensionPausingNewPausingAddressesSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionPausingNewPausingAddressesSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionPausingNewPausingAddressesSigned represents a NewPausingAddressesSigned event raised by the ExtensionPausing contract.
type ExtensionPausingNewPausingAddressesSigned struct {
	ExtensionId    *big.Int
	Nonce          *big.Int
	GovernanceHash [32]byte
	Signer         common.Address
	Signature      Signature
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPausingAddressesSigned is a free log retrieval operation binding the contract event 0x88c185ddecef8e42cabf4b20ca6f5f4b5f43f31aa04ecc8d49c1fcefcd2d0de1.
//
// Solidity: event NewPausingAddressesSigned(uint256 indexed extensionId, uint256 indexed nonce, bytes32 indexed governanceHash, address signer, (uint8,bytes32,bytes32) signature)
func (_ExtensionPausing *ExtensionPausingFilterer) FilterNewPausingAddressesSigned(opts *bind.FilterOpts, extensionId []*big.Int, nonce []*big.Int, governanceHash [][32]byte) (*ExtensionPausingNewPausingAddressesSignedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var governanceHashRule []interface{}
	for _, governanceHashItem := range governanceHash {
		governanceHashRule = append(governanceHashRule, governanceHashItem)
	}

	logs, sub, err := _ExtensionPausing.contract.FilterLogs(opts, "NewPausingAddressesSigned", extensionIdRule, nonceRule, governanceHashRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionPausingNewPausingAddressesSignedIterator{contract: _ExtensionPausing.contract, event: "NewPausingAddressesSigned", logs: logs, sub: sub}, nil
}

// WatchNewPausingAddressesSigned is a free log subscription operation binding the contract event 0x88c185ddecef8e42cabf4b20ca6f5f4b5f43f31aa04ecc8d49c1fcefcd2d0de1.
//
// Solidity: event NewPausingAddressesSigned(uint256 indexed extensionId, uint256 indexed nonce, bytes32 indexed governanceHash, address signer, (uint8,bytes32,bytes32) signature)
func (_ExtensionPausing *ExtensionPausingFilterer) WatchNewPausingAddressesSigned(opts *bind.WatchOpts, sink chan<- *ExtensionPausingNewPausingAddressesSigned, extensionId []*big.Int, nonce []*big.Int, governanceHash [][32]byte) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var governanceHashRule []interface{}
	for _, governanceHashItem := range governanceHash {
		governanceHashRule = append(governanceHashRule, governanceHashItem)
	}

	logs, sub, err := _ExtensionPausing.contract.WatchLogs(opts, "NewPausingAddressesSigned", extensionIdRule, nonceRule, governanceHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionPausingNewPausingAddressesSigned)
				if err := _ExtensionPausing.contract.UnpackLog(event, "NewPausingAddressesSigned", log); err != nil {
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

// ParseNewPausingAddressesSigned is a log parse operation binding the contract event 0x88c185ddecef8e42cabf4b20ca6f5f4b5f43f31aa04ecc8d49c1fcefcd2d0de1.
//
// Solidity: event NewPausingAddressesSigned(uint256 indexed extensionId, uint256 indexed nonce, bytes32 indexed governanceHash, address signer, (uint8,bytes32,bytes32) signature)
func (_ExtensionPausing *ExtensionPausingFilterer) ParseNewPausingAddressesSigned(log types.Log) (*ExtensionPausingNewPausingAddressesSigned, error) {
	event := new(ExtensionPausingNewPausingAddressesSigned)
	if err := _ExtensionPausing.contract.UnpackLog(event, "NewPausingAddressesSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionPausingTeePausingAddressesThresholdMetIterator is returned from FilterTeePausingAddressesThresholdMet and is used to iterate over the raw logs and unpacked data for TeePausingAddressesThresholdMet events raised by the ExtensionPausing contract.
type ExtensionPausingTeePausingAddressesThresholdMetIterator struct {
	Event *ExtensionPausingTeePausingAddressesThresholdMet // Event containing the contract specifics and raw log

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
func (it *ExtensionPausingTeePausingAddressesThresholdMetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionPausingTeePausingAddressesThresholdMet)
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
		it.Event = new(ExtensionPausingTeePausingAddressesThresholdMet)
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
func (it *ExtensionPausingTeePausingAddressesThresholdMetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionPausingTeePausingAddressesThresholdMetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionPausingTeePausingAddressesThresholdMet represents a TeePausingAddressesThresholdMet event raised by the ExtensionPausing contract.
type ExtensionPausingTeePausingAddressesThresholdMet struct {
	ExtensionId    *big.Int
	Nonce          *big.Int
	GovernanceHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTeePausingAddressesThresholdMet is a free log retrieval operation binding the contract event 0x51693d2775948f897e94c6ab0a014f9691805bdc9df6915e2b7ccb90ba80fe0c.
//
// Solidity: event TeePausingAddressesThresholdMet(uint256 indexed extensionId, uint256 indexed nonce, bytes32 indexed governanceHash)
func (_ExtensionPausing *ExtensionPausingFilterer) FilterTeePausingAddressesThresholdMet(opts *bind.FilterOpts, extensionId []*big.Int, nonce []*big.Int, governanceHash [][32]byte) (*ExtensionPausingTeePausingAddressesThresholdMetIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var governanceHashRule []interface{}
	for _, governanceHashItem := range governanceHash {
		governanceHashRule = append(governanceHashRule, governanceHashItem)
	}

	logs, sub, err := _ExtensionPausing.contract.FilterLogs(opts, "TeePausingAddressesThresholdMet", extensionIdRule, nonceRule, governanceHashRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionPausingTeePausingAddressesThresholdMetIterator{contract: _ExtensionPausing.contract, event: "TeePausingAddressesThresholdMet", logs: logs, sub: sub}, nil
}

// WatchTeePausingAddressesThresholdMet is a free log subscription operation binding the contract event 0x51693d2775948f897e94c6ab0a014f9691805bdc9df6915e2b7ccb90ba80fe0c.
//
// Solidity: event TeePausingAddressesThresholdMet(uint256 indexed extensionId, uint256 indexed nonce, bytes32 indexed governanceHash)
func (_ExtensionPausing *ExtensionPausingFilterer) WatchTeePausingAddressesThresholdMet(opts *bind.WatchOpts, sink chan<- *ExtensionPausingTeePausingAddressesThresholdMet, extensionId []*big.Int, nonce []*big.Int, governanceHash [][32]byte) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var governanceHashRule []interface{}
	for _, governanceHashItem := range governanceHash {
		governanceHashRule = append(governanceHashRule, governanceHashItem)
	}

	logs, sub, err := _ExtensionPausing.contract.WatchLogs(opts, "TeePausingAddressesThresholdMet", extensionIdRule, nonceRule, governanceHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionPausingTeePausingAddressesThresholdMet)
				if err := _ExtensionPausing.contract.UnpackLog(event, "TeePausingAddressesThresholdMet", log); err != nil {
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

// ParseTeePausingAddressesThresholdMet is a log parse operation binding the contract event 0x51693d2775948f897e94c6ab0a014f9691805bdc9df6915e2b7ccb90ba80fe0c.
//
// Solidity: event TeePausingAddressesThresholdMet(uint256 indexed extensionId, uint256 indexed nonce, bytes32 indexed governanceHash)
func (_ExtensionPausing *ExtensionPausingFilterer) ParseTeePausingAddressesThresholdMet(log types.Log) (*ExtensionPausingTeePausingAddressesThresholdMet, error) {
	event := new(ExtensionPausingTeePausingAddressesThresholdMet)
	if err := _ExtensionPausing.contract.UnpackLog(event, "TeePausingAddressesThresholdMet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
