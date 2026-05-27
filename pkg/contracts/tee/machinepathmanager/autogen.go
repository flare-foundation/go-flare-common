// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package machinepathmanager

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

// IMachinePathManagerMachinePath is an auto generated low-level Go binding around an user-defined struct.
type IMachinePathManagerMachinePath struct {
	SourceTeeIds      []common.Address
	DestinationTeeIds []common.Address
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// MachinePathManagerMetaData contains all meta data concerning the MachinePathManager contract.
var MachinePathManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DestinationTeeIdAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMachinePath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListAlreadyFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListAlreadySigned\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListNotFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoActiveMachinePathList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoDestinationTeeIds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPaths\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSourceTeeIds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerAlreadySigned\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SourceTeeIdAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeIdNotEligible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnrecognizedSigner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"involvedGovernanceHashes\",\"type\":\"bytes32[]\"}],\"name\":\"MachinePathListFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"countedGovernanceHashes\",\"type\":\"bytes32[]\"}],\"name\":\"MachinePathListSignatureAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"MachinePathListSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"MachinePathListStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"sourceTeeIds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"destinationTeeIds\",\"type\":\"address[]\"}],\"indexed\":false,\"internalType\":\"structIMachinePathManager.MachinePath[]\",\"name\":\"paths\",\"type\":\"tuple[]\"}],\"name\":\"MachinePathsAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"sourceTeeIds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"destinationTeeIds\",\"type\":\"address[]\"}],\"internalType\":\"structIMachinePathManager.MachinePath[]\",\"name\":\"_paths\",\"type\":\"tuple[]\"}],\"name\":\"addMachinePaths\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"createNewMachinePathList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"finalizeMachinePathList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getActiveMachinePathListNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getMachinePathList\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"sourceTeeIds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"destinationTeeIds\",\"type\":\"address[]\"}],\"internalType\":\"structIMachinePathManager.MachinePath[]\",\"name\":\"_paths\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_involvedGovernanceHashes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"_signed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getMachinePathListMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"getMachinePathListSignatureCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getMachinePathListsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"isMachinePathListFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"isMachinePathListSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sourceTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationTeeId\",\"type\":\"address\"}],\"name\":\"isMachinePathValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signMachinePathList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MachinePathManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MachinePathManagerMetaData.ABI instead.
var MachinePathManagerABI = MachinePathManagerMetaData.ABI

// MachinePathManager is an auto generated Go binding around an Ethereum contract.
type MachinePathManager struct {
	MachinePathManagerCaller     // Read-only binding to the contract
	MachinePathManagerTransactor // Write-only binding to the contract
	MachinePathManagerFilterer   // Log filterer for contract events
}

// MachinePathManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MachinePathManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachinePathManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MachinePathManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachinePathManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MachinePathManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachinePathManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MachinePathManagerSession struct {
	Contract     *MachinePathManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MachinePathManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MachinePathManagerCallerSession struct {
	Contract *MachinePathManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MachinePathManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MachinePathManagerTransactorSession struct {
	Contract     *MachinePathManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MachinePathManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MachinePathManagerRaw struct {
	Contract *MachinePathManager // Generic contract binding to access the raw methods on
}

// MachinePathManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MachinePathManagerCallerRaw struct {
	Contract *MachinePathManagerCaller // Generic read-only contract binding to access the raw methods on
}

// MachinePathManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MachinePathManagerTransactorRaw struct {
	Contract *MachinePathManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMachinePathManager creates a new instance of MachinePathManager, bound to a specific deployed contract.
func NewMachinePathManager(address common.Address, backend bind.ContractBackend) (*MachinePathManager, error) {
	contract, err := bindMachinePathManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MachinePathManager{MachinePathManagerCaller: MachinePathManagerCaller{contract: contract}, MachinePathManagerTransactor: MachinePathManagerTransactor{contract: contract}, MachinePathManagerFilterer: MachinePathManagerFilterer{contract: contract}}, nil
}

// NewMachinePathManagerCaller creates a new read-only instance of MachinePathManager, bound to a specific deployed contract.
func NewMachinePathManagerCaller(address common.Address, caller bind.ContractCaller) (*MachinePathManagerCaller, error) {
	contract, err := bindMachinePathManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MachinePathManagerCaller{contract: contract}, nil
}

// NewMachinePathManagerTransactor creates a new write-only instance of MachinePathManager, bound to a specific deployed contract.
func NewMachinePathManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*MachinePathManagerTransactor, error) {
	contract, err := bindMachinePathManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MachinePathManagerTransactor{contract: contract}, nil
}

// NewMachinePathManagerFilterer creates a new log filterer instance of MachinePathManager, bound to a specific deployed contract.
func NewMachinePathManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*MachinePathManagerFilterer, error) {
	contract, err := bindMachinePathManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MachinePathManagerFilterer{contract: contract}, nil
}

// bindMachinePathManager binds a generic wrapper to an already deployed contract.
func bindMachinePathManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MachinePathManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MachinePathManager *MachinePathManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MachinePathManager.Contract.MachinePathManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MachinePathManager *MachinePathManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MachinePathManager.Contract.MachinePathManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MachinePathManager *MachinePathManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MachinePathManager.Contract.MachinePathManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MachinePathManager *MachinePathManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MachinePathManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MachinePathManager *MachinePathManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MachinePathManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MachinePathManager *MachinePathManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MachinePathManager.Contract.contract.Transact(opts, method, params...)
}

// GetActiveMachinePathListNonce is a free data retrieval call binding the contract method 0xb9b27a4b.
//
// Solidity: function getActiveMachinePathListNonce(uint256 _extensionId) view returns(uint256 _nonce)
func (_MachinePathManager *MachinePathManagerCaller) GetActiveMachinePathListNonce(opts *bind.CallOpts, _extensionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MachinePathManager.contract.Call(opts, &out, "getActiveMachinePathListNonce", _extensionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveMachinePathListNonce is a free data retrieval call binding the contract method 0xb9b27a4b.
//
// Solidity: function getActiveMachinePathListNonce(uint256 _extensionId) view returns(uint256 _nonce)
func (_MachinePathManager *MachinePathManagerSession) GetActiveMachinePathListNonce(_extensionId *big.Int) (*big.Int, error) {
	return _MachinePathManager.Contract.GetActiveMachinePathListNonce(&_MachinePathManager.CallOpts, _extensionId)
}

// GetActiveMachinePathListNonce is a free data retrieval call binding the contract method 0xb9b27a4b.
//
// Solidity: function getActiveMachinePathListNonce(uint256 _extensionId) view returns(uint256 _nonce)
func (_MachinePathManager *MachinePathManagerCallerSession) GetActiveMachinePathListNonce(_extensionId *big.Int) (*big.Int, error) {
	return _MachinePathManager.Contract.GetActiveMachinePathListNonce(&_MachinePathManager.CallOpts, _extensionId)
}

// GetMachinePathList is a free data retrieval call binding the contract method 0x02473311.
//
// Solidity: function getMachinePathList(uint256 _extensionId, uint256 _nonce) view returns((address[],address[])[] _paths, bytes32[] _involvedGovernanceHashes, (uint8,bytes32,bytes32)[] _signatures, bool _signed)
func (_MachinePathManager *MachinePathManagerCaller) GetMachinePathList(opts *bind.CallOpts, _extensionId *big.Int, _nonce *big.Int) (struct {
	Paths                    []IMachinePathManagerMachinePath
	InvolvedGovernanceHashes [][32]byte
	Signatures               []Signature
	Signed                   bool
}, error) {
	var out []interface{}
	err := _MachinePathManager.contract.Call(opts, &out, "getMachinePathList", _extensionId, _nonce)

	outstruct := new(struct {
		Paths                    []IMachinePathManagerMachinePath
		InvolvedGovernanceHashes [][32]byte
		Signatures               []Signature
		Signed                   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Paths = *abi.ConvertType(out[0], new([]IMachinePathManagerMachinePath)).(*[]IMachinePathManagerMachinePath)
	outstruct.InvolvedGovernanceHashes = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)
	outstruct.Signatures = *abi.ConvertType(out[2], new([]Signature)).(*[]Signature)
	outstruct.Signed = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// GetMachinePathList is a free data retrieval call binding the contract method 0x02473311.
//
// Solidity: function getMachinePathList(uint256 _extensionId, uint256 _nonce) view returns((address[],address[])[] _paths, bytes32[] _involvedGovernanceHashes, (uint8,bytes32,bytes32)[] _signatures, bool _signed)
func (_MachinePathManager *MachinePathManagerSession) GetMachinePathList(_extensionId *big.Int, _nonce *big.Int) (struct {
	Paths                    []IMachinePathManagerMachinePath
	InvolvedGovernanceHashes [][32]byte
	Signatures               []Signature
	Signed                   bool
}, error) {
	return _MachinePathManager.Contract.GetMachinePathList(&_MachinePathManager.CallOpts, _extensionId, _nonce)
}

// GetMachinePathList is a free data retrieval call binding the contract method 0x02473311.
//
// Solidity: function getMachinePathList(uint256 _extensionId, uint256 _nonce) view returns((address[],address[])[] _paths, bytes32[] _involvedGovernanceHashes, (uint8,bytes32,bytes32)[] _signatures, bool _signed)
func (_MachinePathManager *MachinePathManagerCallerSession) GetMachinePathList(_extensionId *big.Int, _nonce *big.Int) (struct {
	Paths                    []IMachinePathManagerMachinePath
	InvolvedGovernanceHashes [][32]byte
	Signatures               []Signature
	Signed                   bool
}, error) {
	return _MachinePathManager.Contract.GetMachinePathList(&_MachinePathManager.CallOpts, _extensionId, _nonce)
}

// GetMachinePathListMessageHash is a free data retrieval call binding the contract method 0xb27f4702.
//
// Solidity: function getMachinePathListMessageHash(uint256 _extensionId, uint256 _nonce) view returns(bytes32)
func (_MachinePathManager *MachinePathManagerCaller) GetMachinePathListMessageHash(opts *bind.CallOpts, _extensionId *big.Int, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MachinePathManager.contract.Call(opts, &out, "getMachinePathListMessageHash", _extensionId, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetMachinePathListMessageHash is a free data retrieval call binding the contract method 0xb27f4702.
//
// Solidity: function getMachinePathListMessageHash(uint256 _extensionId, uint256 _nonce) view returns(bytes32)
func (_MachinePathManager *MachinePathManagerSession) GetMachinePathListMessageHash(_extensionId *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _MachinePathManager.Contract.GetMachinePathListMessageHash(&_MachinePathManager.CallOpts, _extensionId, _nonce)
}

// GetMachinePathListMessageHash is a free data retrieval call binding the contract method 0xb27f4702.
//
// Solidity: function getMachinePathListMessageHash(uint256 _extensionId, uint256 _nonce) view returns(bytes32)
func (_MachinePathManager *MachinePathManagerCallerSession) GetMachinePathListMessageHash(_extensionId *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _MachinePathManager.Contract.GetMachinePathListMessageHash(&_MachinePathManager.CallOpts, _extensionId, _nonce)
}

// GetMachinePathListSignatureCount is a free data retrieval call binding the contract method 0xa4acbad8.
//
// Solidity: function getMachinePathListSignatureCount(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash) view returns(uint64)
func (_MachinePathManager *MachinePathManagerCaller) GetMachinePathListSignatureCount(opts *bind.CallOpts, _extensionId *big.Int, _nonce *big.Int, _governanceHash [32]byte) (uint64, error) {
	var out []interface{}
	err := _MachinePathManager.contract.Call(opts, &out, "getMachinePathListSignatureCount", _extensionId, _nonce, _governanceHash)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetMachinePathListSignatureCount is a free data retrieval call binding the contract method 0xa4acbad8.
//
// Solidity: function getMachinePathListSignatureCount(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash) view returns(uint64)
func (_MachinePathManager *MachinePathManagerSession) GetMachinePathListSignatureCount(_extensionId *big.Int, _nonce *big.Int, _governanceHash [32]byte) (uint64, error) {
	return _MachinePathManager.Contract.GetMachinePathListSignatureCount(&_MachinePathManager.CallOpts, _extensionId, _nonce, _governanceHash)
}

// GetMachinePathListSignatureCount is a free data retrieval call binding the contract method 0xa4acbad8.
//
// Solidity: function getMachinePathListSignatureCount(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash) view returns(uint64)
func (_MachinePathManager *MachinePathManagerCallerSession) GetMachinePathListSignatureCount(_extensionId *big.Int, _nonce *big.Int, _governanceHash [32]byte) (uint64, error) {
	return _MachinePathManager.Contract.GetMachinePathListSignatureCount(&_MachinePathManager.CallOpts, _extensionId, _nonce, _governanceHash)
}

// GetMachinePathListsCount is a free data retrieval call binding the contract method 0xa71b1c71.
//
// Solidity: function getMachinePathListsCount(uint256 _extensionId) view returns(uint256)
func (_MachinePathManager *MachinePathManagerCaller) GetMachinePathListsCount(opts *bind.CallOpts, _extensionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MachinePathManager.contract.Call(opts, &out, "getMachinePathListsCount", _extensionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMachinePathListsCount is a free data retrieval call binding the contract method 0xa71b1c71.
//
// Solidity: function getMachinePathListsCount(uint256 _extensionId) view returns(uint256)
func (_MachinePathManager *MachinePathManagerSession) GetMachinePathListsCount(_extensionId *big.Int) (*big.Int, error) {
	return _MachinePathManager.Contract.GetMachinePathListsCount(&_MachinePathManager.CallOpts, _extensionId)
}

// GetMachinePathListsCount is a free data retrieval call binding the contract method 0xa71b1c71.
//
// Solidity: function getMachinePathListsCount(uint256 _extensionId) view returns(uint256)
func (_MachinePathManager *MachinePathManagerCallerSession) GetMachinePathListsCount(_extensionId *big.Int) (*big.Int, error) {
	return _MachinePathManager.Contract.GetMachinePathListsCount(&_MachinePathManager.CallOpts, _extensionId)
}

// IsMachinePathListFinalized is a free data retrieval call binding the contract method 0xb367ffed.
//
// Solidity: function isMachinePathListFinalized(uint256 _extensionId, uint256 _nonce) view returns(bool)
func (_MachinePathManager *MachinePathManagerCaller) IsMachinePathListFinalized(opts *bind.CallOpts, _extensionId *big.Int, _nonce *big.Int) (bool, error) {
	var out []interface{}
	err := _MachinePathManager.contract.Call(opts, &out, "isMachinePathListFinalized", _extensionId, _nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMachinePathListFinalized is a free data retrieval call binding the contract method 0xb367ffed.
//
// Solidity: function isMachinePathListFinalized(uint256 _extensionId, uint256 _nonce) view returns(bool)
func (_MachinePathManager *MachinePathManagerSession) IsMachinePathListFinalized(_extensionId *big.Int, _nonce *big.Int) (bool, error) {
	return _MachinePathManager.Contract.IsMachinePathListFinalized(&_MachinePathManager.CallOpts, _extensionId, _nonce)
}

// IsMachinePathListFinalized is a free data retrieval call binding the contract method 0xb367ffed.
//
// Solidity: function isMachinePathListFinalized(uint256 _extensionId, uint256 _nonce) view returns(bool)
func (_MachinePathManager *MachinePathManagerCallerSession) IsMachinePathListFinalized(_extensionId *big.Int, _nonce *big.Int) (bool, error) {
	return _MachinePathManager.Contract.IsMachinePathListFinalized(&_MachinePathManager.CallOpts, _extensionId, _nonce)
}

// IsMachinePathListSigned is a free data retrieval call binding the contract method 0x70838db9.
//
// Solidity: function isMachinePathListSigned(uint256 _extensionId, uint256 _nonce) view returns(bool)
func (_MachinePathManager *MachinePathManagerCaller) IsMachinePathListSigned(opts *bind.CallOpts, _extensionId *big.Int, _nonce *big.Int) (bool, error) {
	var out []interface{}
	err := _MachinePathManager.contract.Call(opts, &out, "isMachinePathListSigned", _extensionId, _nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMachinePathListSigned is a free data retrieval call binding the contract method 0x70838db9.
//
// Solidity: function isMachinePathListSigned(uint256 _extensionId, uint256 _nonce) view returns(bool)
func (_MachinePathManager *MachinePathManagerSession) IsMachinePathListSigned(_extensionId *big.Int, _nonce *big.Int) (bool, error) {
	return _MachinePathManager.Contract.IsMachinePathListSigned(&_MachinePathManager.CallOpts, _extensionId, _nonce)
}

// IsMachinePathListSigned is a free data retrieval call binding the contract method 0x70838db9.
//
// Solidity: function isMachinePathListSigned(uint256 _extensionId, uint256 _nonce) view returns(bool)
func (_MachinePathManager *MachinePathManagerCallerSession) IsMachinePathListSigned(_extensionId *big.Int, _nonce *big.Int) (bool, error) {
	return _MachinePathManager.Contract.IsMachinePathListSigned(&_MachinePathManager.CallOpts, _extensionId, _nonce)
}

// IsMachinePathValid is a free data retrieval call binding the contract method 0x76ffee5c.
//
// Solidity: function isMachinePathValid(uint256 _extensionId, address _sourceTeeId, address _destinationTeeId) view returns(bool)
func (_MachinePathManager *MachinePathManagerCaller) IsMachinePathValid(opts *bind.CallOpts, _extensionId *big.Int, _sourceTeeId common.Address, _destinationTeeId common.Address) (bool, error) {
	var out []interface{}
	err := _MachinePathManager.contract.Call(opts, &out, "isMachinePathValid", _extensionId, _sourceTeeId, _destinationTeeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMachinePathValid is a free data retrieval call binding the contract method 0x76ffee5c.
//
// Solidity: function isMachinePathValid(uint256 _extensionId, address _sourceTeeId, address _destinationTeeId) view returns(bool)
func (_MachinePathManager *MachinePathManagerSession) IsMachinePathValid(_extensionId *big.Int, _sourceTeeId common.Address, _destinationTeeId common.Address) (bool, error) {
	return _MachinePathManager.Contract.IsMachinePathValid(&_MachinePathManager.CallOpts, _extensionId, _sourceTeeId, _destinationTeeId)
}

// IsMachinePathValid is a free data retrieval call binding the contract method 0x76ffee5c.
//
// Solidity: function isMachinePathValid(uint256 _extensionId, address _sourceTeeId, address _destinationTeeId) view returns(bool)
func (_MachinePathManager *MachinePathManagerCallerSession) IsMachinePathValid(_extensionId *big.Int, _sourceTeeId common.Address, _destinationTeeId common.Address) (bool, error) {
	return _MachinePathManager.Contract.IsMachinePathValid(&_MachinePathManager.CallOpts, _extensionId, _sourceTeeId, _destinationTeeId)
}

// AddMachinePaths is a paid mutator transaction binding the contract method 0xac6dad90.
//
// Solidity: function addMachinePaths(uint256 _extensionId, uint256 _nonce, (address[],address[])[] _paths) returns()
func (_MachinePathManager *MachinePathManagerTransactor) AddMachinePaths(opts *bind.TransactOpts, _extensionId *big.Int, _nonce *big.Int, _paths []IMachinePathManagerMachinePath) (*types.Transaction, error) {
	return _MachinePathManager.contract.Transact(opts, "addMachinePaths", _extensionId, _nonce, _paths)
}

// AddMachinePaths is a paid mutator transaction binding the contract method 0xac6dad90.
//
// Solidity: function addMachinePaths(uint256 _extensionId, uint256 _nonce, (address[],address[])[] _paths) returns()
func (_MachinePathManager *MachinePathManagerSession) AddMachinePaths(_extensionId *big.Int, _nonce *big.Int, _paths []IMachinePathManagerMachinePath) (*types.Transaction, error) {
	return _MachinePathManager.Contract.AddMachinePaths(&_MachinePathManager.TransactOpts, _extensionId, _nonce, _paths)
}

// AddMachinePaths is a paid mutator transaction binding the contract method 0xac6dad90.
//
// Solidity: function addMachinePaths(uint256 _extensionId, uint256 _nonce, (address[],address[])[] _paths) returns()
func (_MachinePathManager *MachinePathManagerTransactorSession) AddMachinePaths(_extensionId *big.Int, _nonce *big.Int, _paths []IMachinePathManagerMachinePath) (*types.Transaction, error) {
	return _MachinePathManager.Contract.AddMachinePaths(&_MachinePathManager.TransactOpts, _extensionId, _nonce, _paths)
}

// CreateNewMachinePathList is a paid mutator transaction binding the contract method 0x07ae64b2.
//
// Solidity: function createNewMachinePathList(uint256 _extensionId) returns(uint256 _nonce)
func (_MachinePathManager *MachinePathManagerTransactor) CreateNewMachinePathList(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _MachinePathManager.contract.Transact(opts, "createNewMachinePathList", _extensionId)
}

// CreateNewMachinePathList is a paid mutator transaction binding the contract method 0x07ae64b2.
//
// Solidity: function createNewMachinePathList(uint256 _extensionId) returns(uint256 _nonce)
func (_MachinePathManager *MachinePathManagerSession) CreateNewMachinePathList(_extensionId *big.Int) (*types.Transaction, error) {
	return _MachinePathManager.Contract.CreateNewMachinePathList(&_MachinePathManager.TransactOpts, _extensionId)
}

// CreateNewMachinePathList is a paid mutator transaction binding the contract method 0x07ae64b2.
//
// Solidity: function createNewMachinePathList(uint256 _extensionId) returns(uint256 _nonce)
func (_MachinePathManager *MachinePathManagerTransactorSession) CreateNewMachinePathList(_extensionId *big.Int) (*types.Transaction, error) {
	return _MachinePathManager.Contract.CreateNewMachinePathList(&_MachinePathManager.TransactOpts, _extensionId)
}

// FinalizeMachinePathList is a paid mutator transaction binding the contract method 0x7fb88568.
//
// Solidity: function finalizeMachinePathList(uint256 _extensionId, uint256 _nonce) returns()
func (_MachinePathManager *MachinePathManagerTransactor) FinalizeMachinePathList(opts *bind.TransactOpts, _extensionId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _MachinePathManager.contract.Transact(opts, "finalizeMachinePathList", _extensionId, _nonce)
}

// FinalizeMachinePathList is a paid mutator transaction binding the contract method 0x7fb88568.
//
// Solidity: function finalizeMachinePathList(uint256 _extensionId, uint256 _nonce) returns()
func (_MachinePathManager *MachinePathManagerSession) FinalizeMachinePathList(_extensionId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _MachinePathManager.Contract.FinalizeMachinePathList(&_MachinePathManager.TransactOpts, _extensionId, _nonce)
}

// FinalizeMachinePathList is a paid mutator transaction binding the contract method 0x7fb88568.
//
// Solidity: function finalizeMachinePathList(uint256 _extensionId, uint256 _nonce) returns()
func (_MachinePathManager *MachinePathManagerTransactorSession) FinalizeMachinePathList(_extensionId *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _MachinePathManager.Contract.FinalizeMachinePathList(&_MachinePathManager.TransactOpts, _extensionId, _nonce)
}

// SignMachinePathList is a paid mutator transaction binding the contract method 0xcce3332d.
//
// Solidity: function signMachinePathList(uint256 _extensionId, uint256 _nonce, (uint8,bytes32,bytes32) _signature) returns()
func (_MachinePathManager *MachinePathManagerTransactor) SignMachinePathList(opts *bind.TransactOpts, _extensionId *big.Int, _nonce *big.Int, _signature Signature) (*types.Transaction, error) {
	return _MachinePathManager.contract.Transact(opts, "signMachinePathList", _extensionId, _nonce, _signature)
}

// SignMachinePathList is a paid mutator transaction binding the contract method 0xcce3332d.
//
// Solidity: function signMachinePathList(uint256 _extensionId, uint256 _nonce, (uint8,bytes32,bytes32) _signature) returns()
func (_MachinePathManager *MachinePathManagerSession) SignMachinePathList(_extensionId *big.Int, _nonce *big.Int, _signature Signature) (*types.Transaction, error) {
	return _MachinePathManager.Contract.SignMachinePathList(&_MachinePathManager.TransactOpts, _extensionId, _nonce, _signature)
}

// SignMachinePathList is a paid mutator transaction binding the contract method 0xcce3332d.
//
// Solidity: function signMachinePathList(uint256 _extensionId, uint256 _nonce, (uint8,bytes32,bytes32) _signature) returns()
func (_MachinePathManager *MachinePathManagerTransactorSession) SignMachinePathList(_extensionId *big.Int, _nonce *big.Int, _signature Signature) (*types.Transaction, error) {
	return _MachinePathManager.Contract.SignMachinePathList(&_MachinePathManager.TransactOpts, _extensionId, _nonce, _signature)
}

// MachinePathManagerMachinePathListFinalizedIterator is returned from FilterMachinePathListFinalized and is used to iterate over the raw logs and unpacked data for MachinePathListFinalized events raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListFinalizedIterator struct {
	Event *MachinePathManagerMachinePathListFinalized // Event containing the contract specifics and raw log

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
func (it *MachinePathManagerMachinePathListFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachinePathManagerMachinePathListFinalized)
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
		it.Event = new(MachinePathManagerMachinePathListFinalized)
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
func (it *MachinePathManagerMachinePathListFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachinePathManagerMachinePathListFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachinePathManagerMachinePathListFinalized represents a MachinePathListFinalized event raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListFinalized struct {
	ExtensionId              *big.Int
	Nonce                    *big.Int
	InvolvedGovernanceHashes [][32]byte
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterMachinePathListFinalized is a free log retrieval operation binding the contract event 0x9f237b84521640079b688c3779dab9b800ecda4287ed347f6c24210148a6247c.
//
// Solidity: event MachinePathListFinalized(uint256 indexed extensionId, uint256 indexed nonce, bytes32[] involvedGovernanceHashes)
func (_MachinePathManager *MachinePathManagerFilterer) FilterMachinePathListFinalized(opts *bind.FilterOpts, extensionId []*big.Int, nonce []*big.Int) (*MachinePathManagerMachinePathListFinalizedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MachinePathManager.contract.FilterLogs(opts, "MachinePathListFinalized", extensionIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &MachinePathManagerMachinePathListFinalizedIterator{contract: _MachinePathManager.contract, event: "MachinePathListFinalized", logs: logs, sub: sub}, nil
}

// WatchMachinePathListFinalized is a free log subscription operation binding the contract event 0x9f237b84521640079b688c3779dab9b800ecda4287ed347f6c24210148a6247c.
//
// Solidity: event MachinePathListFinalized(uint256 indexed extensionId, uint256 indexed nonce, bytes32[] involvedGovernanceHashes)
func (_MachinePathManager *MachinePathManagerFilterer) WatchMachinePathListFinalized(opts *bind.WatchOpts, sink chan<- *MachinePathManagerMachinePathListFinalized, extensionId []*big.Int, nonce []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MachinePathManager.contract.WatchLogs(opts, "MachinePathListFinalized", extensionIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachinePathManagerMachinePathListFinalized)
				if err := _MachinePathManager.contract.UnpackLog(event, "MachinePathListFinalized", log); err != nil {
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

// ParseMachinePathListFinalized is a log parse operation binding the contract event 0x9f237b84521640079b688c3779dab9b800ecda4287ed347f6c24210148a6247c.
//
// Solidity: event MachinePathListFinalized(uint256 indexed extensionId, uint256 indexed nonce, bytes32[] involvedGovernanceHashes)
func (_MachinePathManager *MachinePathManagerFilterer) ParseMachinePathListFinalized(log types.Log) (*MachinePathManagerMachinePathListFinalized, error) {
	event := new(MachinePathManagerMachinePathListFinalized)
	if err := _MachinePathManager.contract.UnpackLog(event, "MachinePathListFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachinePathManagerMachinePathListSignatureAddedIterator is returned from FilterMachinePathListSignatureAdded and is used to iterate over the raw logs and unpacked data for MachinePathListSignatureAdded events raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListSignatureAddedIterator struct {
	Event *MachinePathManagerMachinePathListSignatureAdded // Event containing the contract specifics and raw log

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
func (it *MachinePathManagerMachinePathListSignatureAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachinePathManagerMachinePathListSignatureAdded)
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
		it.Event = new(MachinePathManagerMachinePathListSignatureAdded)
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
func (it *MachinePathManagerMachinePathListSignatureAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachinePathManagerMachinePathListSignatureAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachinePathManagerMachinePathListSignatureAdded represents a MachinePathListSignatureAdded event raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListSignatureAdded struct {
	ExtensionId             *big.Int
	Nonce                   *big.Int
	Signer                  common.Address
	CountedGovernanceHashes [][32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMachinePathListSignatureAdded is a free log retrieval operation binding the contract event 0xd3005d11653c4a8380af635c860cd8ab5692a8aef3e307cb39a7346e620cd665.
//
// Solidity: event MachinePathListSignatureAdded(uint256 indexed extensionId, uint256 indexed nonce, address indexed signer, bytes32[] countedGovernanceHashes)
func (_MachinePathManager *MachinePathManagerFilterer) FilterMachinePathListSignatureAdded(opts *bind.FilterOpts, extensionId []*big.Int, nonce []*big.Int, signer []common.Address) (*MachinePathManagerMachinePathListSignatureAddedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _MachinePathManager.contract.FilterLogs(opts, "MachinePathListSignatureAdded", extensionIdRule, nonceRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &MachinePathManagerMachinePathListSignatureAddedIterator{contract: _MachinePathManager.contract, event: "MachinePathListSignatureAdded", logs: logs, sub: sub}, nil
}

// WatchMachinePathListSignatureAdded is a free log subscription operation binding the contract event 0xd3005d11653c4a8380af635c860cd8ab5692a8aef3e307cb39a7346e620cd665.
//
// Solidity: event MachinePathListSignatureAdded(uint256 indexed extensionId, uint256 indexed nonce, address indexed signer, bytes32[] countedGovernanceHashes)
func (_MachinePathManager *MachinePathManagerFilterer) WatchMachinePathListSignatureAdded(opts *bind.WatchOpts, sink chan<- *MachinePathManagerMachinePathListSignatureAdded, extensionId []*big.Int, nonce []*big.Int, signer []common.Address) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _MachinePathManager.contract.WatchLogs(opts, "MachinePathListSignatureAdded", extensionIdRule, nonceRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachinePathManagerMachinePathListSignatureAdded)
				if err := _MachinePathManager.contract.UnpackLog(event, "MachinePathListSignatureAdded", log); err != nil {
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

// ParseMachinePathListSignatureAdded is a log parse operation binding the contract event 0xd3005d11653c4a8380af635c860cd8ab5692a8aef3e307cb39a7346e620cd665.
//
// Solidity: event MachinePathListSignatureAdded(uint256 indexed extensionId, uint256 indexed nonce, address indexed signer, bytes32[] countedGovernanceHashes)
func (_MachinePathManager *MachinePathManagerFilterer) ParseMachinePathListSignatureAdded(log types.Log) (*MachinePathManagerMachinePathListSignatureAdded, error) {
	event := new(MachinePathManagerMachinePathListSignatureAdded)
	if err := _MachinePathManager.contract.UnpackLog(event, "MachinePathListSignatureAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachinePathManagerMachinePathListSignedIterator is returned from FilterMachinePathListSigned and is used to iterate over the raw logs and unpacked data for MachinePathListSigned events raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListSignedIterator struct {
	Event *MachinePathManagerMachinePathListSigned // Event containing the contract specifics and raw log

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
func (it *MachinePathManagerMachinePathListSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachinePathManagerMachinePathListSigned)
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
		it.Event = new(MachinePathManagerMachinePathListSigned)
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
func (it *MachinePathManagerMachinePathListSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachinePathManagerMachinePathListSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachinePathManagerMachinePathListSigned represents a MachinePathListSigned event raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListSigned struct {
	ExtensionId *big.Int
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMachinePathListSigned is a free log retrieval operation binding the contract event 0x81f37ab661bc938e2691a96974841d81d7d97d19a5f2af658db75332f2439d2c.
//
// Solidity: event MachinePathListSigned(uint256 indexed extensionId, uint256 indexed nonce)
func (_MachinePathManager *MachinePathManagerFilterer) FilterMachinePathListSigned(opts *bind.FilterOpts, extensionId []*big.Int, nonce []*big.Int) (*MachinePathManagerMachinePathListSignedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MachinePathManager.contract.FilterLogs(opts, "MachinePathListSigned", extensionIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &MachinePathManagerMachinePathListSignedIterator{contract: _MachinePathManager.contract, event: "MachinePathListSigned", logs: logs, sub: sub}, nil
}

// WatchMachinePathListSigned is a free log subscription operation binding the contract event 0x81f37ab661bc938e2691a96974841d81d7d97d19a5f2af658db75332f2439d2c.
//
// Solidity: event MachinePathListSigned(uint256 indexed extensionId, uint256 indexed nonce)
func (_MachinePathManager *MachinePathManagerFilterer) WatchMachinePathListSigned(opts *bind.WatchOpts, sink chan<- *MachinePathManagerMachinePathListSigned, extensionId []*big.Int, nonce []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MachinePathManager.contract.WatchLogs(opts, "MachinePathListSigned", extensionIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachinePathManagerMachinePathListSigned)
				if err := _MachinePathManager.contract.UnpackLog(event, "MachinePathListSigned", log); err != nil {
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

// ParseMachinePathListSigned is a log parse operation binding the contract event 0x81f37ab661bc938e2691a96974841d81d7d97d19a5f2af658db75332f2439d2c.
//
// Solidity: event MachinePathListSigned(uint256 indexed extensionId, uint256 indexed nonce)
func (_MachinePathManager *MachinePathManagerFilterer) ParseMachinePathListSigned(log types.Log) (*MachinePathManagerMachinePathListSigned, error) {
	event := new(MachinePathManagerMachinePathListSigned)
	if err := _MachinePathManager.contract.UnpackLog(event, "MachinePathListSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachinePathManagerMachinePathListStartedIterator is returned from FilterMachinePathListStarted and is used to iterate over the raw logs and unpacked data for MachinePathListStarted events raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListStartedIterator struct {
	Event *MachinePathManagerMachinePathListStarted // Event containing the contract specifics and raw log

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
func (it *MachinePathManagerMachinePathListStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachinePathManagerMachinePathListStarted)
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
		it.Event = new(MachinePathManagerMachinePathListStarted)
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
func (it *MachinePathManagerMachinePathListStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachinePathManagerMachinePathListStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachinePathManagerMachinePathListStarted represents a MachinePathListStarted event raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListStarted struct {
	ExtensionId *big.Int
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMachinePathListStarted is a free log retrieval operation binding the contract event 0xb8c47517b99bea16eb9150177c5a4711f957fb53de48047f3ce604a19631fa64.
//
// Solidity: event MachinePathListStarted(uint256 indexed extensionId, uint256 indexed nonce)
func (_MachinePathManager *MachinePathManagerFilterer) FilterMachinePathListStarted(opts *bind.FilterOpts, extensionId []*big.Int, nonce []*big.Int) (*MachinePathManagerMachinePathListStartedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MachinePathManager.contract.FilterLogs(opts, "MachinePathListStarted", extensionIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &MachinePathManagerMachinePathListStartedIterator{contract: _MachinePathManager.contract, event: "MachinePathListStarted", logs: logs, sub: sub}, nil
}

// WatchMachinePathListStarted is a free log subscription operation binding the contract event 0xb8c47517b99bea16eb9150177c5a4711f957fb53de48047f3ce604a19631fa64.
//
// Solidity: event MachinePathListStarted(uint256 indexed extensionId, uint256 indexed nonce)
func (_MachinePathManager *MachinePathManagerFilterer) WatchMachinePathListStarted(opts *bind.WatchOpts, sink chan<- *MachinePathManagerMachinePathListStarted, extensionId []*big.Int, nonce []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MachinePathManager.contract.WatchLogs(opts, "MachinePathListStarted", extensionIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachinePathManagerMachinePathListStarted)
				if err := _MachinePathManager.contract.UnpackLog(event, "MachinePathListStarted", log); err != nil {
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

// ParseMachinePathListStarted is a log parse operation binding the contract event 0xb8c47517b99bea16eb9150177c5a4711f957fb53de48047f3ce604a19631fa64.
//
// Solidity: event MachinePathListStarted(uint256 indexed extensionId, uint256 indexed nonce)
func (_MachinePathManager *MachinePathManagerFilterer) ParseMachinePathListStarted(log types.Log) (*MachinePathManagerMachinePathListStarted, error) {
	event := new(MachinePathManagerMachinePathListStarted)
	if err := _MachinePathManager.contract.UnpackLog(event, "MachinePathListStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachinePathManagerMachinePathsAddedIterator is returned from FilterMachinePathsAdded and is used to iterate over the raw logs and unpacked data for MachinePathsAdded events raised by the MachinePathManager contract.
type MachinePathManagerMachinePathsAddedIterator struct {
	Event *MachinePathManagerMachinePathsAdded // Event containing the contract specifics and raw log

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
func (it *MachinePathManagerMachinePathsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachinePathManagerMachinePathsAdded)
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
		it.Event = new(MachinePathManagerMachinePathsAdded)
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
func (it *MachinePathManagerMachinePathsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachinePathManagerMachinePathsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachinePathManagerMachinePathsAdded represents a MachinePathsAdded event raised by the MachinePathManager contract.
type MachinePathManagerMachinePathsAdded struct {
	ExtensionId *big.Int
	Nonce       *big.Int
	Paths       []IMachinePathManagerMachinePath
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMachinePathsAdded is a free log retrieval operation binding the contract event 0xb431778feef3dfa1ab02d748f50bfbac8ad13947791c2d21cd52e808f52562d0.
//
// Solidity: event MachinePathsAdded(uint256 indexed extensionId, uint256 indexed nonce, (address[],address[])[] paths)
func (_MachinePathManager *MachinePathManagerFilterer) FilterMachinePathsAdded(opts *bind.FilterOpts, extensionId []*big.Int, nonce []*big.Int) (*MachinePathManagerMachinePathsAddedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MachinePathManager.contract.FilterLogs(opts, "MachinePathsAdded", extensionIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &MachinePathManagerMachinePathsAddedIterator{contract: _MachinePathManager.contract, event: "MachinePathsAdded", logs: logs, sub: sub}, nil
}

// WatchMachinePathsAdded is a free log subscription operation binding the contract event 0xb431778feef3dfa1ab02d748f50bfbac8ad13947791c2d21cd52e808f52562d0.
//
// Solidity: event MachinePathsAdded(uint256 indexed extensionId, uint256 indexed nonce, (address[],address[])[] paths)
func (_MachinePathManager *MachinePathManagerFilterer) WatchMachinePathsAdded(opts *bind.WatchOpts, sink chan<- *MachinePathManagerMachinePathsAdded, extensionId []*big.Int, nonce []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MachinePathManager.contract.WatchLogs(opts, "MachinePathsAdded", extensionIdRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachinePathManagerMachinePathsAdded)
				if err := _MachinePathManager.contract.UnpackLog(event, "MachinePathsAdded", log); err != nil {
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

// ParseMachinePathsAdded is a log parse operation binding the contract event 0xb431778feef3dfa1ab02d748f50bfbac8ad13947791c2d21cd52e808f52562d0.
//
// Solidity: event MachinePathsAdded(uint256 indexed extensionId, uint256 indexed nonce, (address[],address[])[] paths)
func (_MachinePathManager *MachinePathManagerFilterer) ParseMachinePathsAdded(log types.Log) (*MachinePathManagerMachinePathsAdded, error) {
	event := new(MachinePathManagerMachinePathsAdded)
	if err := _MachinePathManager.contract.UnpackLog(event, "MachinePathsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
