// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package upgrademanager

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

// IUpgradeManagerTeeNodeVersion is an auto generated low-level Go binding around an user-defined struct.
type IUpgradeManagerTeeNodeVersion struct {
	CodeHash [32]byte
	Platform [32]byte
}

// IUpgradeManagerTeeUpgradePath is an auto generated low-level Go binding around an user-defined struct.
type IUpgradeManagerTeeUpgradePath struct {
	SourceVersions []IUpgradeManagerTeeNodeVersion
	TargetVersions []IUpgradeManagerTeeNodeVersion
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// UpgradeManagerMetaData contains all meta data concerning the UpgradeManager contract.
var UpgradeManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFromGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidToGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUpgradeId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSourceVersions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTargetVersions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoUpgradePaths\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SourceCodeHashAndPlatformNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SourceGovernanceHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SourceVersionAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetCodeHashAndPlatformNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetGovernanceHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetVersionAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeAlreadyFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeAlreadySigned\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpgradeNotFinalized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teeUpgradeId\",\"type\":\"uint256\"}],\"name\":\"TeeUpgradeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teeUpgradeId\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIUpgradeManager.TeeNodeVersion[]\",\"name\":\"sourceVersions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIUpgradeManager.TeeNodeVersion[]\",\"name\":\"targetVersions\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structIUpgradeManager.TeeUpgradePath[]\",\"name\":\"upgradePaths\",\"type\":\"tuple[]\"}],\"name\":\"TeeUpgradePathsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teeUpgradeId\",\"type\":\"uint256\"}],\"name\":\"TeeUpgradeSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teeUpgradeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"TeeUpgradeSourceSignatureAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teeUpgradeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceTeeGovernanceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"targetTeeGovernanceHash\",\"type\":\"bytes32\"}],\"name\":\"TeeUpgradeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"teeUpgradeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"TeeUpgradeTargetSignatureAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_teeUpgradeId\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIUpgradeManager.TeeNodeVersion[]\",\"name\":\"sourceVersions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIUpgradeManager.TeeNodeVersion[]\",\"name\":\"targetVersions\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUpgradeManager.TeeUpgradePath[]\",\"name\":\"_upgradePaths\",\"type\":\"tuple[]\"}],\"name\":\"addTeeUpgradePaths\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceTeeGovernanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_targetTeeGovernanceHash\",\"type\":\"bytes32\"}],\"name\":\"createNewTeeUpgrade\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_teeUpgradeId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_teeUpgradeId\",\"type\":\"uint256\"}],\"name\":\"finalizeTeeUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_teeUpgradeId\",\"type\":\"uint256\"}],\"name\":\"getTeeUpgradeMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_teeUpgradeId\",\"type\":\"uint256\"}],\"name\":\"getTeeUpgradePaths\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIUpgradeManager.TeeNodeVersion[]\",\"name\":\"sourceVersions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIUpgradeManager.TeeNodeVersion[]\",\"name\":\"targetVersions\",\"type\":\"tuple[]\"}],\"internalType\":\"structIUpgradeManager.TeeUpgradePath[]\",\"name\":\"upgradePaths\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_teeUpgradeId\",\"type\":\"uint256\"}],\"name\":\"getTeeUpgradeSignatures\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"sourceTeeGovernanceSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"targetTeeGovernanceSignatures\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTeeUpgradesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_teeUpgradeId\",\"type\":\"uint256\"}],\"name\":\"isTeeUpgradeFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_teeUpgradeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourcePlatform\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_targetCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_targetPlatform\",\"type\":\"bytes32\"}],\"name\":\"isTeeUpgradePathValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_teeUpgradeId\",\"type\":\"uint256\"}],\"name\":\"isTeeUpgradeSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_teeUpgradeId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signTeeUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UpgradeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeManagerMetaData.ABI instead.
var UpgradeManagerABI = UpgradeManagerMetaData.ABI

// UpgradeManager is an auto generated Go binding around an Ethereum contract.
type UpgradeManager struct {
	UpgradeManagerCaller     // Read-only binding to the contract
	UpgradeManagerTransactor // Write-only binding to the contract
	UpgradeManagerFilterer   // Log filterer for contract events
}

// UpgradeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeManagerSession struct {
	Contract     *UpgradeManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeManagerCallerSession struct {
	Contract *UpgradeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// UpgradeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeManagerTransactorSession struct {
	Contract     *UpgradeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// UpgradeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeManagerRaw struct {
	Contract *UpgradeManager // Generic contract binding to access the raw methods on
}

// UpgradeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeManagerCallerRaw struct {
	Contract *UpgradeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeManagerTransactorRaw struct {
	Contract *UpgradeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeManager creates a new instance of UpgradeManager, bound to a specific deployed contract.
func NewUpgradeManager(address common.Address, backend bind.ContractBackend) (*UpgradeManager, error) {
	contract, err := bindUpgradeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeManager{UpgradeManagerCaller: UpgradeManagerCaller{contract: contract}, UpgradeManagerTransactor: UpgradeManagerTransactor{contract: contract}, UpgradeManagerFilterer: UpgradeManagerFilterer{contract: contract}}, nil
}

// NewUpgradeManagerCaller creates a new read-only instance of UpgradeManager, bound to a specific deployed contract.
func NewUpgradeManagerCaller(address common.Address, caller bind.ContractCaller) (*UpgradeManagerCaller, error) {
	contract, err := bindUpgradeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeManagerCaller{contract: contract}, nil
}

// NewUpgradeManagerTransactor creates a new write-only instance of UpgradeManager, bound to a specific deployed contract.
func NewUpgradeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeManagerTransactor, error) {
	contract, err := bindUpgradeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeManagerTransactor{contract: contract}, nil
}

// NewUpgradeManagerFilterer creates a new log filterer instance of UpgradeManager, bound to a specific deployed contract.
func NewUpgradeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeManagerFilterer, error) {
	contract, err := bindUpgradeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeManagerFilterer{contract: contract}, nil
}

// bindUpgradeManager binds a generic wrapper to an already deployed contract.
func bindUpgradeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpgradeManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeManager *UpgradeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeManager.Contract.UpgradeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeManager *UpgradeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeManager.Contract.UpgradeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeManager *UpgradeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeManager.Contract.UpgradeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeManager *UpgradeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeManager *UpgradeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeManager *UpgradeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeManager.Contract.contract.Transact(opts, method, params...)
}

// GetTeeUpgradeMessageHash is a free data retrieval call binding the contract method 0x497b2b16.
//
// Solidity: function getTeeUpgradeMessageHash(uint256 _teeUpgradeId) view returns(bytes32)
func (_UpgradeManager *UpgradeManagerCaller) GetTeeUpgradeMessageHash(opts *bind.CallOpts, _teeUpgradeId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _UpgradeManager.contract.Call(opts, &out, "getTeeUpgradeMessageHash", _teeUpgradeId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTeeUpgradeMessageHash is a free data retrieval call binding the contract method 0x497b2b16.
//
// Solidity: function getTeeUpgradeMessageHash(uint256 _teeUpgradeId) view returns(bytes32)
func (_UpgradeManager *UpgradeManagerSession) GetTeeUpgradeMessageHash(_teeUpgradeId *big.Int) ([32]byte, error) {
	return _UpgradeManager.Contract.GetTeeUpgradeMessageHash(&_UpgradeManager.CallOpts, _teeUpgradeId)
}

// GetTeeUpgradeMessageHash is a free data retrieval call binding the contract method 0x497b2b16.
//
// Solidity: function getTeeUpgradeMessageHash(uint256 _teeUpgradeId) view returns(bytes32)
func (_UpgradeManager *UpgradeManagerCallerSession) GetTeeUpgradeMessageHash(_teeUpgradeId *big.Int) ([32]byte, error) {
	return _UpgradeManager.Contract.GetTeeUpgradeMessageHash(&_UpgradeManager.CallOpts, _teeUpgradeId)
}

// GetTeeUpgradePaths is a free data retrieval call binding the contract method 0x832b1f07.
//
// Solidity: function getTeeUpgradePaths(uint256 _teeUpgradeId) view returns(((bytes32,bytes32)[],(bytes32,bytes32)[])[] upgradePaths)
func (_UpgradeManager *UpgradeManagerCaller) GetTeeUpgradePaths(opts *bind.CallOpts, _teeUpgradeId *big.Int) ([]IUpgradeManagerTeeUpgradePath, error) {
	var out []interface{}
	err := _UpgradeManager.contract.Call(opts, &out, "getTeeUpgradePaths", _teeUpgradeId)

	if err != nil {
		return *new([]IUpgradeManagerTeeUpgradePath), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUpgradeManagerTeeUpgradePath)).(*[]IUpgradeManagerTeeUpgradePath)

	return out0, err

}

// GetTeeUpgradePaths is a free data retrieval call binding the contract method 0x832b1f07.
//
// Solidity: function getTeeUpgradePaths(uint256 _teeUpgradeId) view returns(((bytes32,bytes32)[],(bytes32,bytes32)[])[] upgradePaths)
func (_UpgradeManager *UpgradeManagerSession) GetTeeUpgradePaths(_teeUpgradeId *big.Int) ([]IUpgradeManagerTeeUpgradePath, error) {
	return _UpgradeManager.Contract.GetTeeUpgradePaths(&_UpgradeManager.CallOpts, _teeUpgradeId)
}

// GetTeeUpgradePaths is a free data retrieval call binding the contract method 0x832b1f07.
//
// Solidity: function getTeeUpgradePaths(uint256 _teeUpgradeId) view returns(((bytes32,bytes32)[],(bytes32,bytes32)[])[] upgradePaths)
func (_UpgradeManager *UpgradeManagerCallerSession) GetTeeUpgradePaths(_teeUpgradeId *big.Int) ([]IUpgradeManagerTeeUpgradePath, error) {
	return _UpgradeManager.Contract.GetTeeUpgradePaths(&_UpgradeManager.CallOpts, _teeUpgradeId)
}

// GetTeeUpgradeSignatures is a free data retrieval call binding the contract method 0x31df785c.
//
// Solidity: function getTeeUpgradeSignatures(uint256 _teeUpgradeId) view returns((uint8,bytes32,bytes32)[] sourceTeeGovernanceSignatures, (uint8,bytes32,bytes32)[] targetTeeGovernanceSignatures)
func (_UpgradeManager *UpgradeManagerCaller) GetTeeUpgradeSignatures(opts *bind.CallOpts, _teeUpgradeId *big.Int) (struct {
	SourceTeeGovernanceSignatures []Signature
	TargetTeeGovernanceSignatures []Signature
}, error) {
	var out []interface{}
	err := _UpgradeManager.contract.Call(opts, &out, "getTeeUpgradeSignatures", _teeUpgradeId)

	outstruct := new(struct {
		SourceTeeGovernanceSignatures []Signature
		TargetTeeGovernanceSignatures []Signature
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SourceTeeGovernanceSignatures = *abi.ConvertType(out[0], new([]Signature)).(*[]Signature)
	outstruct.TargetTeeGovernanceSignatures = *abi.ConvertType(out[1], new([]Signature)).(*[]Signature)

	return *outstruct, err

}

// GetTeeUpgradeSignatures is a free data retrieval call binding the contract method 0x31df785c.
//
// Solidity: function getTeeUpgradeSignatures(uint256 _teeUpgradeId) view returns((uint8,bytes32,bytes32)[] sourceTeeGovernanceSignatures, (uint8,bytes32,bytes32)[] targetTeeGovernanceSignatures)
func (_UpgradeManager *UpgradeManagerSession) GetTeeUpgradeSignatures(_teeUpgradeId *big.Int) (struct {
	SourceTeeGovernanceSignatures []Signature
	TargetTeeGovernanceSignatures []Signature
}, error) {
	return _UpgradeManager.Contract.GetTeeUpgradeSignatures(&_UpgradeManager.CallOpts, _teeUpgradeId)
}

// GetTeeUpgradeSignatures is a free data retrieval call binding the contract method 0x31df785c.
//
// Solidity: function getTeeUpgradeSignatures(uint256 _teeUpgradeId) view returns((uint8,bytes32,bytes32)[] sourceTeeGovernanceSignatures, (uint8,bytes32,bytes32)[] targetTeeGovernanceSignatures)
func (_UpgradeManager *UpgradeManagerCallerSession) GetTeeUpgradeSignatures(_teeUpgradeId *big.Int) (struct {
	SourceTeeGovernanceSignatures []Signature
	TargetTeeGovernanceSignatures []Signature
}, error) {
	return _UpgradeManager.Contract.GetTeeUpgradeSignatures(&_UpgradeManager.CallOpts, _teeUpgradeId)
}

// GetTeeUpgradesCount is a free data retrieval call binding the contract method 0x1839c376.
//
// Solidity: function getTeeUpgradesCount() view returns(uint256)
func (_UpgradeManager *UpgradeManagerCaller) GetTeeUpgradesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UpgradeManager.contract.Call(opts, &out, "getTeeUpgradesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTeeUpgradesCount is a free data retrieval call binding the contract method 0x1839c376.
//
// Solidity: function getTeeUpgradesCount() view returns(uint256)
func (_UpgradeManager *UpgradeManagerSession) GetTeeUpgradesCount() (*big.Int, error) {
	return _UpgradeManager.Contract.GetTeeUpgradesCount(&_UpgradeManager.CallOpts)
}

// GetTeeUpgradesCount is a free data retrieval call binding the contract method 0x1839c376.
//
// Solidity: function getTeeUpgradesCount() view returns(uint256)
func (_UpgradeManager *UpgradeManagerCallerSession) GetTeeUpgradesCount() (*big.Int, error) {
	return _UpgradeManager.Contract.GetTeeUpgradesCount(&_UpgradeManager.CallOpts)
}

// IsTeeUpgradeFinalized is a free data retrieval call binding the contract method 0xc4a953d5.
//
// Solidity: function isTeeUpgradeFinalized(uint256 _teeUpgradeId) view returns(bool)
func (_UpgradeManager *UpgradeManagerCaller) IsTeeUpgradeFinalized(opts *bind.CallOpts, _teeUpgradeId *big.Int) (bool, error) {
	var out []interface{}
	err := _UpgradeManager.contract.Call(opts, &out, "isTeeUpgradeFinalized", _teeUpgradeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeeUpgradeFinalized is a free data retrieval call binding the contract method 0xc4a953d5.
//
// Solidity: function isTeeUpgradeFinalized(uint256 _teeUpgradeId) view returns(bool)
func (_UpgradeManager *UpgradeManagerSession) IsTeeUpgradeFinalized(_teeUpgradeId *big.Int) (bool, error) {
	return _UpgradeManager.Contract.IsTeeUpgradeFinalized(&_UpgradeManager.CallOpts, _teeUpgradeId)
}

// IsTeeUpgradeFinalized is a free data retrieval call binding the contract method 0xc4a953d5.
//
// Solidity: function isTeeUpgradeFinalized(uint256 _teeUpgradeId) view returns(bool)
func (_UpgradeManager *UpgradeManagerCallerSession) IsTeeUpgradeFinalized(_teeUpgradeId *big.Int) (bool, error) {
	return _UpgradeManager.Contract.IsTeeUpgradeFinalized(&_UpgradeManager.CallOpts, _teeUpgradeId)
}

// IsTeeUpgradePathValid is a free data retrieval call binding the contract method 0xa9913bb8.
//
// Solidity: function isTeeUpgradePathValid(uint256 _teeUpgradeId, uint256 _extensionId, bytes32 _sourceCodeHash, bytes32 _sourcePlatform, bytes32 _targetCodeHash, bytes32 _targetPlatform) view returns(bool)
func (_UpgradeManager *UpgradeManagerCaller) IsTeeUpgradePathValid(opts *bind.CallOpts, _teeUpgradeId *big.Int, _extensionId *big.Int, _sourceCodeHash [32]byte, _sourcePlatform [32]byte, _targetCodeHash [32]byte, _targetPlatform [32]byte) (bool, error) {
	var out []interface{}
	err := _UpgradeManager.contract.Call(opts, &out, "isTeeUpgradePathValid", _teeUpgradeId, _extensionId, _sourceCodeHash, _sourcePlatform, _targetCodeHash, _targetPlatform)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeeUpgradePathValid is a free data retrieval call binding the contract method 0xa9913bb8.
//
// Solidity: function isTeeUpgradePathValid(uint256 _teeUpgradeId, uint256 _extensionId, bytes32 _sourceCodeHash, bytes32 _sourcePlatform, bytes32 _targetCodeHash, bytes32 _targetPlatform) view returns(bool)
func (_UpgradeManager *UpgradeManagerSession) IsTeeUpgradePathValid(_teeUpgradeId *big.Int, _extensionId *big.Int, _sourceCodeHash [32]byte, _sourcePlatform [32]byte, _targetCodeHash [32]byte, _targetPlatform [32]byte) (bool, error) {
	return _UpgradeManager.Contract.IsTeeUpgradePathValid(&_UpgradeManager.CallOpts, _teeUpgradeId, _extensionId, _sourceCodeHash, _sourcePlatform, _targetCodeHash, _targetPlatform)
}

// IsTeeUpgradePathValid is a free data retrieval call binding the contract method 0xa9913bb8.
//
// Solidity: function isTeeUpgradePathValid(uint256 _teeUpgradeId, uint256 _extensionId, bytes32 _sourceCodeHash, bytes32 _sourcePlatform, bytes32 _targetCodeHash, bytes32 _targetPlatform) view returns(bool)
func (_UpgradeManager *UpgradeManagerCallerSession) IsTeeUpgradePathValid(_teeUpgradeId *big.Int, _extensionId *big.Int, _sourceCodeHash [32]byte, _sourcePlatform [32]byte, _targetCodeHash [32]byte, _targetPlatform [32]byte) (bool, error) {
	return _UpgradeManager.Contract.IsTeeUpgradePathValid(&_UpgradeManager.CallOpts, _teeUpgradeId, _extensionId, _sourceCodeHash, _sourcePlatform, _targetCodeHash, _targetPlatform)
}

// IsTeeUpgradeSigned is a free data retrieval call binding the contract method 0x6d90d753.
//
// Solidity: function isTeeUpgradeSigned(uint256 _teeUpgradeId) view returns(bool)
func (_UpgradeManager *UpgradeManagerCaller) IsTeeUpgradeSigned(opts *bind.CallOpts, _teeUpgradeId *big.Int) (bool, error) {
	var out []interface{}
	err := _UpgradeManager.contract.Call(opts, &out, "isTeeUpgradeSigned", _teeUpgradeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTeeUpgradeSigned is a free data retrieval call binding the contract method 0x6d90d753.
//
// Solidity: function isTeeUpgradeSigned(uint256 _teeUpgradeId) view returns(bool)
func (_UpgradeManager *UpgradeManagerSession) IsTeeUpgradeSigned(_teeUpgradeId *big.Int) (bool, error) {
	return _UpgradeManager.Contract.IsTeeUpgradeSigned(&_UpgradeManager.CallOpts, _teeUpgradeId)
}

// IsTeeUpgradeSigned is a free data retrieval call binding the contract method 0x6d90d753.
//
// Solidity: function isTeeUpgradeSigned(uint256 _teeUpgradeId) view returns(bool)
func (_UpgradeManager *UpgradeManagerCallerSession) IsTeeUpgradeSigned(_teeUpgradeId *big.Int) (bool, error) {
	return _UpgradeManager.Contract.IsTeeUpgradeSigned(&_UpgradeManager.CallOpts, _teeUpgradeId)
}

// AddTeeUpgradePaths is a paid mutator transaction binding the contract method 0x0042a695.
//
// Solidity: function addTeeUpgradePaths(uint256 _teeUpgradeId, ((bytes32,bytes32)[],(bytes32,bytes32)[])[] _upgradePaths) returns()
func (_UpgradeManager *UpgradeManagerTransactor) AddTeeUpgradePaths(opts *bind.TransactOpts, _teeUpgradeId *big.Int, _upgradePaths []IUpgradeManagerTeeUpgradePath) (*types.Transaction, error) {
	return _UpgradeManager.contract.Transact(opts, "addTeeUpgradePaths", _teeUpgradeId, _upgradePaths)
}

// AddTeeUpgradePaths is a paid mutator transaction binding the contract method 0x0042a695.
//
// Solidity: function addTeeUpgradePaths(uint256 _teeUpgradeId, ((bytes32,bytes32)[],(bytes32,bytes32)[])[] _upgradePaths) returns()
func (_UpgradeManager *UpgradeManagerSession) AddTeeUpgradePaths(_teeUpgradeId *big.Int, _upgradePaths []IUpgradeManagerTeeUpgradePath) (*types.Transaction, error) {
	return _UpgradeManager.Contract.AddTeeUpgradePaths(&_UpgradeManager.TransactOpts, _teeUpgradeId, _upgradePaths)
}

// AddTeeUpgradePaths is a paid mutator transaction binding the contract method 0x0042a695.
//
// Solidity: function addTeeUpgradePaths(uint256 _teeUpgradeId, ((bytes32,bytes32)[],(bytes32,bytes32)[])[] _upgradePaths) returns()
func (_UpgradeManager *UpgradeManagerTransactorSession) AddTeeUpgradePaths(_teeUpgradeId *big.Int, _upgradePaths []IUpgradeManagerTeeUpgradePath) (*types.Transaction, error) {
	return _UpgradeManager.Contract.AddTeeUpgradePaths(&_UpgradeManager.TransactOpts, _teeUpgradeId, _upgradePaths)
}

// CreateNewTeeUpgrade is a paid mutator transaction binding the contract method 0x9bec3cc8.
//
// Solidity: function createNewTeeUpgrade(uint256 _extensionId, bytes32 _sourceTeeGovernanceHash, bytes32 _targetTeeGovernanceHash) returns(uint256 _teeUpgradeId)
func (_UpgradeManager *UpgradeManagerTransactor) CreateNewTeeUpgrade(opts *bind.TransactOpts, _extensionId *big.Int, _sourceTeeGovernanceHash [32]byte, _targetTeeGovernanceHash [32]byte) (*types.Transaction, error) {
	return _UpgradeManager.contract.Transact(opts, "createNewTeeUpgrade", _extensionId, _sourceTeeGovernanceHash, _targetTeeGovernanceHash)
}

// CreateNewTeeUpgrade is a paid mutator transaction binding the contract method 0x9bec3cc8.
//
// Solidity: function createNewTeeUpgrade(uint256 _extensionId, bytes32 _sourceTeeGovernanceHash, bytes32 _targetTeeGovernanceHash) returns(uint256 _teeUpgradeId)
func (_UpgradeManager *UpgradeManagerSession) CreateNewTeeUpgrade(_extensionId *big.Int, _sourceTeeGovernanceHash [32]byte, _targetTeeGovernanceHash [32]byte) (*types.Transaction, error) {
	return _UpgradeManager.Contract.CreateNewTeeUpgrade(&_UpgradeManager.TransactOpts, _extensionId, _sourceTeeGovernanceHash, _targetTeeGovernanceHash)
}

// CreateNewTeeUpgrade is a paid mutator transaction binding the contract method 0x9bec3cc8.
//
// Solidity: function createNewTeeUpgrade(uint256 _extensionId, bytes32 _sourceTeeGovernanceHash, bytes32 _targetTeeGovernanceHash) returns(uint256 _teeUpgradeId)
func (_UpgradeManager *UpgradeManagerTransactorSession) CreateNewTeeUpgrade(_extensionId *big.Int, _sourceTeeGovernanceHash [32]byte, _targetTeeGovernanceHash [32]byte) (*types.Transaction, error) {
	return _UpgradeManager.Contract.CreateNewTeeUpgrade(&_UpgradeManager.TransactOpts, _extensionId, _sourceTeeGovernanceHash, _targetTeeGovernanceHash)
}

// FinalizeTeeUpgrade is a paid mutator transaction binding the contract method 0x8dcdeafc.
//
// Solidity: function finalizeTeeUpgrade(uint256 _teeUpgradeId) returns()
func (_UpgradeManager *UpgradeManagerTransactor) FinalizeTeeUpgrade(opts *bind.TransactOpts, _teeUpgradeId *big.Int) (*types.Transaction, error) {
	return _UpgradeManager.contract.Transact(opts, "finalizeTeeUpgrade", _teeUpgradeId)
}

// FinalizeTeeUpgrade is a paid mutator transaction binding the contract method 0x8dcdeafc.
//
// Solidity: function finalizeTeeUpgrade(uint256 _teeUpgradeId) returns()
func (_UpgradeManager *UpgradeManagerSession) FinalizeTeeUpgrade(_teeUpgradeId *big.Int) (*types.Transaction, error) {
	return _UpgradeManager.Contract.FinalizeTeeUpgrade(&_UpgradeManager.TransactOpts, _teeUpgradeId)
}

// FinalizeTeeUpgrade is a paid mutator transaction binding the contract method 0x8dcdeafc.
//
// Solidity: function finalizeTeeUpgrade(uint256 _teeUpgradeId) returns()
func (_UpgradeManager *UpgradeManagerTransactorSession) FinalizeTeeUpgrade(_teeUpgradeId *big.Int) (*types.Transaction, error) {
	return _UpgradeManager.Contract.FinalizeTeeUpgrade(&_UpgradeManager.TransactOpts, _teeUpgradeId)
}

// SignTeeUpgrade is a paid mutator transaction binding the contract method 0x481ff52f.
//
// Solidity: function signTeeUpgrade(uint256 _teeUpgradeId, (uint8,bytes32,bytes32) _signature) returns()
func (_UpgradeManager *UpgradeManagerTransactor) SignTeeUpgrade(opts *bind.TransactOpts, _teeUpgradeId *big.Int, _signature Signature) (*types.Transaction, error) {
	return _UpgradeManager.contract.Transact(opts, "signTeeUpgrade", _teeUpgradeId, _signature)
}

// SignTeeUpgrade is a paid mutator transaction binding the contract method 0x481ff52f.
//
// Solidity: function signTeeUpgrade(uint256 _teeUpgradeId, (uint8,bytes32,bytes32) _signature) returns()
func (_UpgradeManager *UpgradeManagerSession) SignTeeUpgrade(_teeUpgradeId *big.Int, _signature Signature) (*types.Transaction, error) {
	return _UpgradeManager.Contract.SignTeeUpgrade(&_UpgradeManager.TransactOpts, _teeUpgradeId, _signature)
}

// SignTeeUpgrade is a paid mutator transaction binding the contract method 0x481ff52f.
//
// Solidity: function signTeeUpgrade(uint256 _teeUpgradeId, (uint8,bytes32,bytes32) _signature) returns()
func (_UpgradeManager *UpgradeManagerTransactorSession) SignTeeUpgrade(_teeUpgradeId *big.Int, _signature Signature) (*types.Transaction, error) {
	return _UpgradeManager.Contract.SignTeeUpgrade(&_UpgradeManager.TransactOpts, _teeUpgradeId, _signature)
}

// UpgradeManagerTeeUpgradeFinalizedIterator is returned from FilterTeeUpgradeFinalized and is used to iterate over the raw logs and unpacked data for TeeUpgradeFinalized events raised by the UpgradeManager contract.
type UpgradeManagerTeeUpgradeFinalizedIterator struct {
	Event *UpgradeManagerTeeUpgradeFinalized // Event containing the contract specifics and raw log

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
func (it *UpgradeManagerTeeUpgradeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeManagerTeeUpgradeFinalized)
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
		it.Event = new(UpgradeManagerTeeUpgradeFinalized)
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
func (it *UpgradeManagerTeeUpgradeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeManagerTeeUpgradeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeManagerTeeUpgradeFinalized represents a TeeUpgradeFinalized event raised by the UpgradeManager contract.
type UpgradeManagerTeeUpgradeFinalized struct {
	TeeUpgradeId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTeeUpgradeFinalized is a free log retrieval operation binding the contract event 0x8d233b65016936f5c465d2cb9e04111b097656ae0ac31e3412a8d5eb196a50f1.
//
// Solidity: event TeeUpgradeFinalized(uint256 indexed teeUpgradeId)
func (_UpgradeManager *UpgradeManagerFilterer) FilterTeeUpgradeFinalized(opts *bind.FilterOpts, teeUpgradeId []*big.Int) (*UpgradeManagerTeeUpgradeFinalizedIterator, error) {

	var teeUpgradeIdRule []interface{}
	for _, teeUpgradeIdItem := range teeUpgradeId {
		teeUpgradeIdRule = append(teeUpgradeIdRule, teeUpgradeIdItem)
	}

	logs, sub, err := _UpgradeManager.contract.FilterLogs(opts, "TeeUpgradeFinalized", teeUpgradeIdRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeManagerTeeUpgradeFinalizedIterator{contract: _UpgradeManager.contract, event: "TeeUpgradeFinalized", logs: logs, sub: sub}, nil
}

// WatchTeeUpgradeFinalized is a free log subscription operation binding the contract event 0x8d233b65016936f5c465d2cb9e04111b097656ae0ac31e3412a8d5eb196a50f1.
//
// Solidity: event TeeUpgradeFinalized(uint256 indexed teeUpgradeId)
func (_UpgradeManager *UpgradeManagerFilterer) WatchTeeUpgradeFinalized(opts *bind.WatchOpts, sink chan<- *UpgradeManagerTeeUpgradeFinalized, teeUpgradeId []*big.Int) (event.Subscription, error) {

	var teeUpgradeIdRule []interface{}
	for _, teeUpgradeIdItem := range teeUpgradeId {
		teeUpgradeIdRule = append(teeUpgradeIdRule, teeUpgradeIdItem)
	}

	logs, sub, err := _UpgradeManager.contract.WatchLogs(opts, "TeeUpgradeFinalized", teeUpgradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeManagerTeeUpgradeFinalized)
				if err := _UpgradeManager.contract.UnpackLog(event, "TeeUpgradeFinalized", log); err != nil {
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

// ParseTeeUpgradeFinalized is a log parse operation binding the contract event 0x8d233b65016936f5c465d2cb9e04111b097656ae0ac31e3412a8d5eb196a50f1.
//
// Solidity: event TeeUpgradeFinalized(uint256 indexed teeUpgradeId)
func (_UpgradeManager *UpgradeManagerFilterer) ParseTeeUpgradeFinalized(log types.Log) (*UpgradeManagerTeeUpgradeFinalized, error) {
	event := new(UpgradeManagerTeeUpgradeFinalized)
	if err := _UpgradeManager.contract.UnpackLog(event, "TeeUpgradeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeManagerTeeUpgradePathsAddedIterator is returned from FilterTeeUpgradePathsAdded and is used to iterate over the raw logs and unpacked data for TeeUpgradePathsAdded events raised by the UpgradeManager contract.
type UpgradeManagerTeeUpgradePathsAddedIterator struct {
	Event *UpgradeManagerTeeUpgradePathsAdded // Event containing the contract specifics and raw log

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
func (it *UpgradeManagerTeeUpgradePathsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeManagerTeeUpgradePathsAdded)
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
		it.Event = new(UpgradeManagerTeeUpgradePathsAdded)
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
func (it *UpgradeManagerTeeUpgradePathsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeManagerTeeUpgradePathsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeManagerTeeUpgradePathsAdded represents a TeeUpgradePathsAdded event raised by the UpgradeManager contract.
type UpgradeManagerTeeUpgradePathsAdded struct {
	TeeUpgradeId *big.Int
	UpgradePaths []IUpgradeManagerTeeUpgradePath
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTeeUpgradePathsAdded is a free log retrieval operation binding the contract event 0x69ec533f8683746130fbb60afeeebce498eb1289700489bdb347878bb176c0ef.
//
// Solidity: event TeeUpgradePathsAdded(uint256 indexed teeUpgradeId, ((bytes32,bytes32)[],(bytes32,bytes32)[])[] upgradePaths)
func (_UpgradeManager *UpgradeManagerFilterer) FilterTeeUpgradePathsAdded(opts *bind.FilterOpts, teeUpgradeId []*big.Int) (*UpgradeManagerTeeUpgradePathsAddedIterator, error) {

	var teeUpgradeIdRule []interface{}
	for _, teeUpgradeIdItem := range teeUpgradeId {
		teeUpgradeIdRule = append(teeUpgradeIdRule, teeUpgradeIdItem)
	}

	logs, sub, err := _UpgradeManager.contract.FilterLogs(opts, "TeeUpgradePathsAdded", teeUpgradeIdRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeManagerTeeUpgradePathsAddedIterator{contract: _UpgradeManager.contract, event: "TeeUpgradePathsAdded", logs: logs, sub: sub}, nil
}

// WatchTeeUpgradePathsAdded is a free log subscription operation binding the contract event 0x69ec533f8683746130fbb60afeeebce498eb1289700489bdb347878bb176c0ef.
//
// Solidity: event TeeUpgradePathsAdded(uint256 indexed teeUpgradeId, ((bytes32,bytes32)[],(bytes32,bytes32)[])[] upgradePaths)
func (_UpgradeManager *UpgradeManagerFilterer) WatchTeeUpgradePathsAdded(opts *bind.WatchOpts, sink chan<- *UpgradeManagerTeeUpgradePathsAdded, teeUpgradeId []*big.Int) (event.Subscription, error) {

	var teeUpgradeIdRule []interface{}
	for _, teeUpgradeIdItem := range teeUpgradeId {
		teeUpgradeIdRule = append(teeUpgradeIdRule, teeUpgradeIdItem)
	}

	logs, sub, err := _UpgradeManager.contract.WatchLogs(opts, "TeeUpgradePathsAdded", teeUpgradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeManagerTeeUpgradePathsAdded)
				if err := _UpgradeManager.contract.UnpackLog(event, "TeeUpgradePathsAdded", log); err != nil {
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

// ParseTeeUpgradePathsAdded is a log parse operation binding the contract event 0x69ec533f8683746130fbb60afeeebce498eb1289700489bdb347878bb176c0ef.
//
// Solidity: event TeeUpgradePathsAdded(uint256 indexed teeUpgradeId, ((bytes32,bytes32)[],(bytes32,bytes32)[])[] upgradePaths)
func (_UpgradeManager *UpgradeManagerFilterer) ParseTeeUpgradePathsAdded(log types.Log) (*UpgradeManagerTeeUpgradePathsAdded, error) {
	event := new(UpgradeManagerTeeUpgradePathsAdded)
	if err := _UpgradeManager.contract.UnpackLog(event, "TeeUpgradePathsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeManagerTeeUpgradeSignedIterator is returned from FilterTeeUpgradeSigned and is used to iterate over the raw logs and unpacked data for TeeUpgradeSigned events raised by the UpgradeManager contract.
type UpgradeManagerTeeUpgradeSignedIterator struct {
	Event *UpgradeManagerTeeUpgradeSigned // Event containing the contract specifics and raw log

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
func (it *UpgradeManagerTeeUpgradeSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeManagerTeeUpgradeSigned)
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
		it.Event = new(UpgradeManagerTeeUpgradeSigned)
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
func (it *UpgradeManagerTeeUpgradeSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeManagerTeeUpgradeSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeManagerTeeUpgradeSigned represents a TeeUpgradeSigned event raised by the UpgradeManager contract.
type UpgradeManagerTeeUpgradeSigned struct {
	TeeUpgradeId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTeeUpgradeSigned is a free log retrieval operation binding the contract event 0x67ac64a3de7c920a35088d3e845a44f4e14eb98e2cd87dddbe14f1ea9cef7cde.
//
// Solidity: event TeeUpgradeSigned(uint256 indexed teeUpgradeId)
func (_UpgradeManager *UpgradeManagerFilterer) FilterTeeUpgradeSigned(opts *bind.FilterOpts, teeUpgradeId []*big.Int) (*UpgradeManagerTeeUpgradeSignedIterator, error) {

	var teeUpgradeIdRule []interface{}
	for _, teeUpgradeIdItem := range teeUpgradeId {
		teeUpgradeIdRule = append(teeUpgradeIdRule, teeUpgradeIdItem)
	}

	logs, sub, err := _UpgradeManager.contract.FilterLogs(opts, "TeeUpgradeSigned", teeUpgradeIdRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeManagerTeeUpgradeSignedIterator{contract: _UpgradeManager.contract, event: "TeeUpgradeSigned", logs: logs, sub: sub}, nil
}

// WatchTeeUpgradeSigned is a free log subscription operation binding the contract event 0x67ac64a3de7c920a35088d3e845a44f4e14eb98e2cd87dddbe14f1ea9cef7cde.
//
// Solidity: event TeeUpgradeSigned(uint256 indexed teeUpgradeId)
func (_UpgradeManager *UpgradeManagerFilterer) WatchTeeUpgradeSigned(opts *bind.WatchOpts, sink chan<- *UpgradeManagerTeeUpgradeSigned, teeUpgradeId []*big.Int) (event.Subscription, error) {

	var teeUpgradeIdRule []interface{}
	for _, teeUpgradeIdItem := range teeUpgradeId {
		teeUpgradeIdRule = append(teeUpgradeIdRule, teeUpgradeIdItem)
	}

	logs, sub, err := _UpgradeManager.contract.WatchLogs(opts, "TeeUpgradeSigned", teeUpgradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeManagerTeeUpgradeSigned)
				if err := _UpgradeManager.contract.UnpackLog(event, "TeeUpgradeSigned", log); err != nil {
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

// ParseTeeUpgradeSigned is a log parse operation binding the contract event 0x67ac64a3de7c920a35088d3e845a44f4e14eb98e2cd87dddbe14f1ea9cef7cde.
//
// Solidity: event TeeUpgradeSigned(uint256 indexed teeUpgradeId)
func (_UpgradeManager *UpgradeManagerFilterer) ParseTeeUpgradeSigned(log types.Log) (*UpgradeManagerTeeUpgradeSigned, error) {
	event := new(UpgradeManagerTeeUpgradeSigned)
	if err := _UpgradeManager.contract.UnpackLog(event, "TeeUpgradeSigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeManagerTeeUpgradeSourceSignatureAddedIterator is returned from FilterTeeUpgradeSourceSignatureAdded and is used to iterate over the raw logs and unpacked data for TeeUpgradeSourceSignatureAdded events raised by the UpgradeManager contract.
type UpgradeManagerTeeUpgradeSourceSignatureAddedIterator struct {
	Event *UpgradeManagerTeeUpgradeSourceSignatureAdded // Event containing the contract specifics and raw log

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
func (it *UpgradeManagerTeeUpgradeSourceSignatureAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeManagerTeeUpgradeSourceSignatureAdded)
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
		it.Event = new(UpgradeManagerTeeUpgradeSourceSignatureAdded)
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
func (it *UpgradeManagerTeeUpgradeSourceSignatureAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeManagerTeeUpgradeSourceSignatureAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeManagerTeeUpgradeSourceSignatureAdded represents a TeeUpgradeSourceSignatureAdded event raised by the UpgradeManager contract.
type UpgradeManagerTeeUpgradeSourceSignatureAdded struct {
	TeeUpgradeId *big.Int
	Signer       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTeeUpgradeSourceSignatureAdded is a free log retrieval operation binding the contract event 0x65b858f34e826c0fafe2e62dbed49a48096b7da6e03873406c7ac48778fe16ad.
//
// Solidity: event TeeUpgradeSourceSignatureAdded(uint256 indexed teeUpgradeId, address indexed signer)
func (_UpgradeManager *UpgradeManagerFilterer) FilterTeeUpgradeSourceSignatureAdded(opts *bind.FilterOpts, teeUpgradeId []*big.Int, signer []common.Address) (*UpgradeManagerTeeUpgradeSourceSignatureAddedIterator, error) {

	var teeUpgradeIdRule []interface{}
	for _, teeUpgradeIdItem := range teeUpgradeId {
		teeUpgradeIdRule = append(teeUpgradeIdRule, teeUpgradeIdItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _UpgradeManager.contract.FilterLogs(opts, "TeeUpgradeSourceSignatureAdded", teeUpgradeIdRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeManagerTeeUpgradeSourceSignatureAddedIterator{contract: _UpgradeManager.contract, event: "TeeUpgradeSourceSignatureAdded", logs: logs, sub: sub}, nil
}

// WatchTeeUpgradeSourceSignatureAdded is a free log subscription operation binding the contract event 0x65b858f34e826c0fafe2e62dbed49a48096b7da6e03873406c7ac48778fe16ad.
//
// Solidity: event TeeUpgradeSourceSignatureAdded(uint256 indexed teeUpgradeId, address indexed signer)
func (_UpgradeManager *UpgradeManagerFilterer) WatchTeeUpgradeSourceSignatureAdded(opts *bind.WatchOpts, sink chan<- *UpgradeManagerTeeUpgradeSourceSignatureAdded, teeUpgradeId []*big.Int, signer []common.Address) (event.Subscription, error) {

	var teeUpgradeIdRule []interface{}
	for _, teeUpgradeIdItem := range teeUpgradeId {
		teeUpgradeIdRule = append(teeUpgradeIdRule, teeUpgradeIdItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _UpgradeManager.contract.WatchLogs(opts, "TeeUpgradeSourceSignatureAdded", teeUpgradeIdRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeManagerTeeUpgradeSourceSignatureAdded)
				if err := _UpgradeManager.contract.UnpackLog(event, "TeeUpgradeSourceSignatureAdded", log); err != nil {
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

// ParseTeeUpgradeSourceSignatureAdded is a log parse operation binding the contract event 0x65b858f34e826c0fafe2e62dbed49a48096b7da6e03873406c7ac48778fe16ad.
//
// Solidity: event TeeUpgradeSourceSignatureAdded(uint256 indexed teeUpgradeId, address indexed signer)
func (_UpgradeManager *UpgradeManagerFilterer) ParseTeeUpgradeSourceSignatureAdded(log types.Log) (*UpgradeManagerTeeUpgradeSourceSignatureAdded, error) {
	event := new(UpgradeManagerTeeUpgradeSourceSignatureAdded)
	if err := _UpgradeManager.contract.UnpackLog(event, "TeeUpgradeSourceSignatureAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeManagerTeeUpgradeStartedIterator is returned from FilterTeeUpgradeStarted and is used to iterate over the raw logs and unpacked data for TeeUpgradeStarted events raised by the UpgradeManager contract.
type UpgradeManagerTeeUpgradeStartedIterator struct {
	Event *UpgradeManagerTeeUpgradeStarted // Event containing the contract specifics and raw log

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
func (it *UpgradeManagerTeeUpgradeStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeManagerTeeUpgradeStarted)
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
		it.Event = new(UpgradeManagerTeeUpgradeStarted)
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
func (it *UpgradeManagerTeeUpgradeStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeManagerTeeUpgradeStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeManagerTeeUpgradeStarted represents a TeeUpgradeStarted event raised by the UpgradeManager contract.
type UpgradeManagerTeeUpgradeStarted struct {
	ExtensionId             *big.Int
	TeeUpgradeId            *big.Int
	SourceTeeGovernanceHash [32]byte
	TargetTeeGovernanceHash [32]byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterTeeUpgradeStarted is a free log retrieval operation binding the contract event 0x9d617452e1975b0e6b83c660eb65f8be917080769ec01118aab23880fbf44aed.
//
// Solidity: event TeeUpgradeStarted(uint256 indexed extensionId, uint256 indexed teeUpgradeId, bytes32 sourceTeeGovernanceHash, bytes32 targetTeeGovernanceHash)
func (_UpgradeManager *UpgradeManagerFilterer) FilterTeeUpgradeStarted(opts *bind.FilterOpts, extensionId []*big.Int, teeUpgradeId []*big.Int) (*UpgradeManagerTeeUpgradeStartedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var teeUpgradeIdRule []interface{}
	for _, teeUpgradeIdItem := range teeUpgradeId {
		teeUpgradeIdRule = append(teeUpgradeIdRule, teeUpgradeIdItem)
	}

	logs, sub, err := _UpgradeManager.contract.FilterLogs(opts, "TeeUpgradeStarted", extensionIdRule, teeUpgradeIdRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeManagerTeeUpgradeStartedIterator{contract: _UpgradeManager.contract, event: "TeeUpgradeStarted", logs: logs, sub: sub}, nil
}

// WatchTeeUpgradeStarted is a free log subscription operation binding the contract event 0x9d617452e1975b0e6b83c660eb65f8be917080769ec01118aab23880fbf44aed.
//
// Solidity: event TeeUpgradeStarted(uint256 indexed extensionId, uint256 indexed teeUpgradeId, bytes32 sourceTeeGovernanceHash, bytes32 targetTeeGovernanceHash)
func (_UpgradeManager *UpgradeManagerFilterer) WatchTeeUpgradeStarted(opts *bind.WatchOpts, sink chan<- *UpgradeManagerTeeUpgradeStarted, extensionId []*big.Int, teeUpgradeId []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var teeUpgradeIdRule []interface{}
	for _, teeUpgradeIdItem := range teeUpgradeId {
		teeUpgradeIdRule = append(teeUpgradeIdRule, teeUpgradeIdItem)
	}

	logs, sub, err := _UpgradeManager.contract.WatchLogs(opts, "TeeUpgradeStarted", extensionIdRule, teeUpgradeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeManagerTeeUpgradeStarted)
				if err := _UpgradeManager.contract.UnpackLog(event, "TeeUpgradeStarted", log); err != nil {
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

// ParseTeeUpgradeStarted is a log parse operation binding the contract event 0x9d617452e1975b0e6b83c660eb65f8be917080769ec01118aab23880fbf44aed.
//
// Solidity: event TeeUpgradeStarted(uint256 indexed extensionId, uint256 indexed teeUpgradeId, bytes32 sourceTeeGovernanceHash, bytes32 targetTeeGovernanceHash)
func (_UpgradeManager *UpgradeManagerFilterer) ParseTeeUpgradeStarted(log types.Log) (*UpgradeManagerTeeUpgradeStarted, error) {
	event := new(UpgradeManagerTeeUpgradeStarted)
	if err := _UpgradeManager.contract.UnpackLog(event, "TeeUpgradeStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UpgradeManagerTeeUpgradeTargetSignatureAddedIterator is returned from FilterTeeUpgradeTargetSignatureAdded and is used to iterate over the raw logs and unpacked data for TeeUpgradeTargetSignatureAdded events raised by the UpgradeManager contract.
type UpgradeManagerTeeUpgradeTargetSignatureAddedIterator struct {
	Event *UpgradeManagerTeeUpgradeTargetSignatureAdded // Event containing the contract specifics and raw log

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
func (it *UpgradeManagerTeeUpgradeTargetSignatureAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeManagerTeeUpgradeTargetSignatureAdded)
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
		it.Event = new(UpgradeManagerTeeUpgradeTargetSignatureAdded)
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
func (it *UpgradeManagerTeeUpgradeTargetSignatureAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeManagerTeeUpgradeTargetSignatureAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeManagerTeeUpgradeTargetSignatureAdded represents a TeeUpgradeTargetSignatureAdded event raised by the UpgradeManager contract.
type UpgradeManagerTeeUpgradeTargetSignatureAdded struct {
	TeeUpgradeId *big.Int
	Signer       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTeeUpgradeTargetSignatureAdded is a free log retrieval operation binding the contract event 0x89f157dbe621729de7f13085be27cea2358e1eecda931c9926ad841581aeaee7.
//
// Solidity: event TeeUpgradeTargetSignatureAdded(uint256 indexed teeUpgradeId, address indexed signer)
func (_UpgradeManager *UpgradeManagerFilterer) FilterTeeUpgradeTargetSignatureAdded(opts *bind.FilterOpts, teeUpgradeId []*big.Int, signer []common.Address) (*UpgradeManagerTeeUpgradeTargetSignatureAddedIterator, error) {

	var teeUpgradeIdRule []interface{}
	for _, teeUpgradeIdItem := range teeUpgradeId {
		teeUpgradeIdRule = append(teeUpgradeIdRule, teeUpgradeIdItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _UpgradeManager.contract.FilterLogs(opts, "TeeUpgradeTargetSignatureAdded", teeUpgradeIdRule, signerRule)
	if err != nil {
		return nil, err
	}
	return &UpgradeManagerTeeUpgradeTargetSignatureAddedIterator{contract: _UpgradeManager.contract, event: "TeeUpgradeTargetSignatureAdded", logs: logs, sub: sub}, nil
}

// WatchTeeUpgradeTargetSignatureAdded is a free log subscription operation binding the contract event 0x89f157dbe621729de7f13085be27cea2358e1eecda931c9926ad841581aeaee7.
//
// Solidity: event TeeUpgradeTargetSignatureAdded(uint256 indexed teeUpgradeId, address indexed signer)
func (_UpgradeManager *UpgradeManagerFilterer) WatchTeeUpgradeTargetSignatureAdded(opts *bind.WatchOpts, sink chan<- *UpgradeManagerTeeUpgradeTargetSignatureAdded, teeUpgradeId []*big.Int, signer []common.Address) (event.Subscription, error) {

	var teeUpgradeIdRule []interface{}
	for _, teeUpgradeIdItem := range teeUpgradeId {
		teeUpgradeIdRule = append(teeUpgradeIdRule, teeUpgradeIdItem)
	}
	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _UpgradeManager.contract.WatchLogs(opts, "TeeUpgradeTargetSignatureAdded", teeUpgradeIdRule, signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeManagerTeeUpgradeTargetSignatureAdded)
				if err := _UpgradeManager.contract.UnpackLog(event, "TeeUpgradeTargetSignatureAdded", log); err != nil {
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

// ParseTeeUpgradeTargetSignatureAdded is a log parse operation binding the contract event 0x89f157dbe621729de7f13085be27cea2358e1eecda931c9926ad841581aeaee7.
//
// Solidity: event TeeUpgradeTargetSignatureAdded(uint256 indexed teeUpgradeId, address indexed signer)
func (_UpgradeManager *UpgradeManagerFilterer) ParseTeeUpgradeTargetSignatureAdded(log types.Log) (*UpgradeManagerTeeUpgradeTargetSignatureAdded, error) {
	event := new(UpgradeManagerTeeUpgradeTargetSignatureAdded)
	if err := _UpgradeManager.contract.UnpackLog(event, "TeeUpgradeTargetSignatureAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
