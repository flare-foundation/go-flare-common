// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teeregistry

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

// ITeeAvailabilityCheckProof is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckProof struct {
	RelayMessage  []byte
	TeeSignatures []Signature
	Data          ITeeAvailabilityCheckResponse
}

// ITeeAvailabilityCheckRequestBody is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckRequestBody struct {
	TeeMachine    ITeeRegistryTeeMachineWithAttestationData
	RewardEpochId *big.Int
}

// ITeeAvailabilityCheckResponse is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckResponse struct {
	AttestationType [32]byte
	SourceId        [32]byte
	ThresholdBIPS   uint16
	Timestamp       uint64
	RequestBody     ITeeAvailabilityCheckRequestBody
	ResponseBody    ITeeAvailabilityCheckResponseBody
}

// ITeeAvailabilityCheckResponseBody is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckResponseBody struct {
	Status uint8
}

// ITeeRegistryTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryTeeMachine struct {
	TeeId common.Address
	Owner common.Address
	Url   string
}

// ITeeRegistryTeeMachineWithAttestationData is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryTeeMachineWithAttestationData struct {
	TeeId    common.Address
	Owner    common.Address
	Url      string
	CodeHash [32]byte
	Platform [32]byte
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// TeeRegistryMetaData contains all meta data concerning the TeeRegistry contract.
var TeeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTs\",\"type\":\"uint256\"}],\"name\":\"AvailabilityCheckValidityExtended\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_backupTeeIds\",\"type\":\"address[]\"}],\"name\":\"arePlatformsCompatible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"relayMessage\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"teeMachine\",\"type\":\"tuple\"},{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Response\",\"name\":\"data\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"confirmAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTeeId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"relayMessage\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"teeMachine\",\"type\":\"tuple\"},{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Response\",\"name\":\"data\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"confirmReplicate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"}],\"name\":\"getCodeHashVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getRandomTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachine\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structITeeRegistry.TeeMachine\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineStatus\",\"outputs\":[{\"internalType\":\"enumITeeRegistry.TeeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineWithAttestationData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_version\",\"type\":\"uint256\"}],\"name\":\"getVersionInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_opTypes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"}],\"name\":\"isOpTypeSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"relayMessage\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"teeMachine\",\"type\":\"tuple\"},{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Response\",\"name\":\"data\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"pauseWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldTeeId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"relayMessage\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"teeMachine\",\"type\":\"tuple\"},{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Response\",\"name\":\"data\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"replicateFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"toPauseForUpgrade\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"relayMessage\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"teeMachine\",\"type\":\"tuple\"},{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Response\",\"name\":\"data\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"toProduction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_testOnTeeId\",\"type\":\"address\"}],\"name\":\"triggerAvailabilityCheck\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeRegistryMetaData.ABI instead.
var TeeRegistryABI = TeeRegistryMetaData.ABI

// TeeRegistry is an auto generated Go binding around an Ethereum contract.
type TeeRegistry struct {
	TeeRegistryCaller     // Read-only binding to the contract
	TeeRegistryTransactor // Write-only binding to the contract
	TeeRegistryFilterer   // Log filterer for contract events
}

// TeeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeRegistrySession struct {
	Contract     *TeeRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeRegistryCallerSession struct {
	Contract *TeeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TeeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeRegistryTransactorSession struct {
	Contract     *TeeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TeeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeRegistryRaw struct {
	Contract *TeeRegistry // Generic contract binding to access the raw methods on
}

// TeeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeRegistryCallerRaw struct {
	Contract *TeeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TeeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeRegistryTransactorRaw struct {
	Contract *TeeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeRegistry creates a new instance of TeeRegistry, bound to a specific deployed contract.
func NewTeeRegistry(address common.Address, backend bind.ContractBackend) (*TeeRegistry, error) {
	contract, err := bindTeeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeRegistry{TeeRegistryCaller: TeeRegistryCaller{contract: contract}, TeeRegistryTransactor: TeeRegistryTransactor{contract: contract}, TeeRegistryFilterer: TeeRegistryFilterer{contract: contract}}, nil
}

// NewTeeRegistryCaller creates a new read-only instance of TeeRegistry, bound to a specific deployed contract.
func NewTeeRegistryCaller(address common.Address, caller bind.ContractCaller) (*TeeRegistryCaller, error) {
	contract, err := bindTeeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryCaller{contract: contract}, nil
}

// NewTeeRegistryTransactor creates a new write-only instance of TeeRegistry, bound to a specific deployed contract.
func NewTeeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeRegistryTransactor, error) {
	contract, err := bindTeeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryTransactor{contract: contract}, nil
}

// NewTeeRegistryFilterer creates a new log filterer instance of TeeRegistry, bound to a specific deployed contract.
func NewTeeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeRegistryFilterer, error) {
	contract, err := bindTeeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryFilterer{contract: contract}, nil
}

// bindTeeRegistry binds a generic wrapper to an already deployed contract.
func bindTeeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeRegistry *TeeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeRegistry.Contract.TeeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeRegistry *TeeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeRegistry.Contract.TeeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeRegistry *TeeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeRegistry.Contract.TeeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeRegistry *TeeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeRegistry *TeeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeRegistry *TeeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeRegistry.Contract.contract.Transact(opts, method, params...)
}

// ArePlatformsCompatible is a free data retrieval call binding the contract method 0x2bcabed9.
//
// Solidity: function arePlatformsCompatible(address _teeId, address[] _backupTeeIds) view returns(bool)
func (_TeeRegistry *TeeRegistryCaller) ArePlatformsCompatible(opts *bind.CallOpts, _teeId common.Address, _backupTeeIds []common.Address) (bool, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "arePlatformsCompatible", _teeId, _backupTeeIds)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ArePlatformsCompatible is a free data retrieval call binding the contract method 0x2bcabed9.
//
// Solidity: function arePlatformsCompatible(address _teeId, address[] _backupTeeIds) view returns(bool)
func (_TeeRegistry *TeeRegistrySession) ArePlatformsCompatible(_teeId common.Address, _backupTeeIds []common.Address) (bool, error) {
	return _TeeRegistry.Contract.ArePlatformsCompatible(&_TeeRegistry.CallOpts, _teeId, _backupTeeIds)
}

// ArePlatformsCompatible is a free data retrieval call binding the contract method 0x2bcabed9.
//
// Solidity: function arePlatformsCompatible(address _teeId, address[] _backupTeeIds) view returns(bool)
func (_TeeRegistry *TeeRegistryCallerSession) ArePlatformsCompatible(_teeId common.Address, _backupTeeIds []common.Address) (bool, error) {
	return _TeeRegistry.Contract.ArePlatformsCompatible(&_TeeRegistry.CallOpts, _teeId, _backupTeeIds)
}

// GetActiveTeeIds is a free data retrieval call binding the contract method 0x5c434d4f.
//
// Solidity: function getActiveTeeIds() view returns(address[])
func (_TeeRegistry *TeeRegistryCaller) GetActiveTeeIds(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getActiveTeeIds")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveTeeIds is a free data retrieval call binding the contract method 0x5c434d4f.
//
// Solidity: function getActiveTeeIds() view returns(address[])
func (_TeeRegistry *TeeRegistrySession) GetActiveTeeIds() ([]common.Address, error) {
	return _TeeRegistry.Contract.GetActiveTeeIds(&_TeeRegistry.CallOpts)
}

// GetActiveTeeIds is a free data retrieval call binding the contract method 0x5c434d4f.
//
// Solidity: function getActiveTeeIds() view returns(address[])
func (_TeeRegistry *TeeRegistryCallerSession) GetActiveTeeIds() ([]common.Address, error) {
	return _TeeRegistry.Contract.GetActiveTeeIds(&_TeeRegistry.CallOpts)
}

// GetCodeHashVersion is a free data retrieval call binding the contract method 0x0cca1d21.
//
// Solidity: function getCodeHashVersion(bytes32 _codeHash) view returns(uint256)
func (_TeeRegistry *TeeRegistryCaller) GetCodeHashVersion(opts *bind.CallOpts, _codeHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getCodeHashVersion", _codeHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCodeHashVersion is a free data retrieval call binding the contract method 0x0cca1d21.
//
// Solidity: function getCodeHashVersion(bytes32 _codeHash) view returns(uint256)
func (_TeeRegistry *TeeRegistrySession) GetCodeHashVersion(_codeHash [32]byte) (*big.Int, error) {
	return _TeeRegistry.Contract.GetCodeHashVersion(&_TeeRegistry.CallOpts, _codeHash)
}

// GetCodeHashVersion is a free data retrieval call binding the contract method 0x0cca1d21.
//
// Solidity: function getCodeHashVersion(bytes32 _codeHash) view returns(uint256)
func (_TeeRegistry *TeeRegistryCallerSession) GetCodeHashVersion(_codeHash [32]byte) (*big.Int, error) {
	return _TeeRegistry.Contract.GetCodeHashVersion(&_TeeRegistry.CallOpts, _codeHash)
}

// GetRandomTeeIds is a free data retrieval call binding the contract method 0xb4db023e.
//
// Solidity: function getRandomTeeIds(uint256 _count) view returns(address[])
func (_TeeRegistry *TeeRegistryCaller) GetRandomTeeIds(opts *bind.CallOpts, _count *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getRandomTeeIds", _count)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRandomTeeIds is a free data retrieval call binding the contract method 0xb4db023e.
//
// Solidity: function getRandomTeeIds(uint256 _count) view returns(address[])
func (_TeeRegistry *TeeRegistrySession) GetRandomTeeIds(_count *big.Int) ([]common.Address, error) {
	return _TeeRegistry.Contract.GetRandomTeeIds(&_TeeRegistry.CallOpts, _count)
}

// GetRandomTeeIds is a free data retrieval call binding the contract method 0xb4db023e.
//
// Solidity: function getRandomTeeIds(uint256 _count) view returns(address[])
func (_TeeRegistry *TeeRegistryCallerSession) GetRandomTeeIds(_count *big.Int) ([]common.Address, error) {
	return _TeeRegistry.Contract.GetRandomTeeIds(&_TeeRegistry.CallOpts, _count)
}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_TeeRegistry *TeeRegistryCaller) GetTeeMachine(opts *bind.CallOpts, _teeId common.Address) (ITeeRegistryTeeMachine, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getTeeMachine", _teeId)

	if err != nil {
		return *new(ITeeRegistryTeeMachine), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeeRegistryTeeMachine)).(*ITeeRegistryTeeMachine)

	return out0, err

}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_TeeRegistry *TeeRegistrySession) GetTeeMachine(_teeId common.Address) (ITeeRegistryTeeMachine, error) {
	return _TeeRegistry.Contract.GetTeeMachine(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_TeeRegistry *TeeRegistryCallerSession) GetTeeMachine(_teeId common.Address) (ITeeRegistryTeeMachine, error) {
	return _TeeRegistry.Contract.GetTeeMachine(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_TeeRegistry *TeeRegistryCaller) GetTeeMachineStatus(opts *bind.CallOpts, _teeId common.Address) (uint8, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getTeeMachineStatus", _teeId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_TeeRegistry *TeeRegistrySession) GetTeeMachineStatus(_teeId common.Address) (uint8, error) {
	return _TeeRegistry.Contract.GetTeeMachineStatus(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_TeeRegistry *TeeRegistryCallerSession) GetTeeMachineStatus(_teeId common.Address) (uint8, error) {
	return _TeeRegistry.Contract.GetTeeMachineStatus(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineVersion is a free data retrieval call binding the contract method 0x6a1fbb80.
//
// Solidity: function getTeeMachineVersion(address _teeId) view returns(uint256)
func (_TeeRegistry *TeeRegistryCaller) GetTeeMachineVersion(opts *bind.CallOpts, _teeId common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getTeeMachineVersion", _teeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTeeMachineVersion is a free data retrieval call binding the contract method 0x6a1fbb80.
//
// Solidity: function getTeeMachineVersion(address _teeId) view returns(uint256)
func (_TeeRegistry *TeeRegistrySession) GetTeeMachineVersion(_teeId common.Address) (*big.Int, error) {
	return _TeeRegistry.Contract.GetTeeMachineVersion(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineVersion is a free data retrieval call binding the contract method 0x6a1fbb80.
//
// Solidity: function getTeeMachineVersion(address _teeId) view returns(uint256)
func (_TeeRegistry *TeeRegistryCallerSession) GetTeeMachineVersion(_teeId common.Address) (*big.Int, error) {
	return _TeeRegistry.Contract.GetTeeMachineVersion(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_TeeRegistry *TeeRegistryCaller) GetTeeMachineWithAttestationData(opts *bind.CallOpts, _teeId common.Address) (ITeeRegistryTeeMachineWithAttestationData, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getTeeMachineWithAttestationData", _teeId)

	if err != nil {
		return *new(ITeeRegistryTeeMachineWithAttestationData), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeeRegistryTeeMachineWithAttestationData)).(*ITeeRegistryTeeMachineWithAttestationData)

	return out0, err

}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_TeeRegistry *TeeRegistrySession) GetTeeMachineWithAttestationData(_teeId common.Address) (ITeeRegistryTeeMachineWithAttestationData, error) {
	return _TeeRegistry.Contract.GetTeeMachineWithAttestationData(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_TeeRegistry *TeeRegistryCallerSession) GetTeeMachineWithAttestationData(_teeId common.Address) (ITeeRegistryTeeMachineWithAttestationData, error) {
	return _TeeRegistry.Contract.GetTeeMachineWithAttestationData(&_TeeRegistry.CallOpts, _teeId)
}

// GetVersionInfo is a free data retrieval call binding the contract method 0xcb4b9e2c.
//
// Solidity: function getVersionInfo(uint256 _version) view returns(bytes32 _codeHash, bytes32[] _platforms, bytes32[] _opTypes)
func (_TeeRegistry *TeeRegistryCaller) GetVersionInfo(opts *bind.CallOpts, _version *big.Int) (struct {
	CodeHash  [32]byte
	Platforms [][32]byte
	OpTypes   [][32]byte
}, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getVersionInfo", _version)

	outstruct := new(struct {
		CodeHash  [32]byte
		Platforms [][32]byte
		OpTypes   [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CodeHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Platforms = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)
	outstruct.OpTypes = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

// GetVersionInfo is a free data retrieval call binding the contract method 0xcb4b9e2c.
//
// Solidity: function getVersionInfo(uint256 _version) view returns(bytes32 _codeHash, bytes32[] _platforms, bytes32[] _opTypes)
func (_TeeRegistry *TeeRegistrySession) GetVersionInfo(_version *big.Int) (struct {
	CodeHash  [32]byte
	Platforms [][32]byte
	OpTypes   [][32]byte
}, error) {
	return _TeeRegistry.Contract.GetVersionInfo(&_TeeRegistry.CallOpts, _version)
}

// GetVersionInfo is a free data retrieval call binding the contract method 0xcb4b9e2c.
//
// Solidity: function getVersionInfo(uint256 _version) view returns(bytes32 _codeHash, bytes32[] _platforms, bytes32[] _opTypes)
func (_TeeRegistry *TeeRegistryCallerSession) GetVersionInfo(_version *big.Int) (struct {
	CodeHash  [32]byte
	Platforms [][32]byte
	OpTypes   [][32]byte
}, error) {
	return _TeeRegistry.Contract.GetVersionInfo(&_TeeRegistry.CallOpts, _version)
}

// IsOpTypeSupported is a free data retrieval call binding the contract method 0xf132d9c8.
//
// Solidity: function isOpTypeSupported(address _teeId, bytes32 _opType) view returns(bool)
func (_TeeRegistry *TeeRegistryCaller) IsOpTypeSupported(opts *bind.CallOpts, _teeId common.Address, _opType [32]byte) (bool, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "isOpTypeSupported", _teeId, _opType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpTypeSupported is a free data retrieval call binding the contract method 0xf132d9c8.
//
// Solidity: function isOpTypeSupported(address _teeId, bytes32 _opType) view returns(bool)
func (_TeeRegistry *TeeRegistrySession) IsOpTypeSupported(_teeId common.Address, _opType [32]byte) (bool, error) {
	return _TeeRegistry.Contract.IsOpTypeSupported(&_TeeRegistry.CallOpts, _teeId, _opType)
}

// IsOpTypeSupported is a free data retrieval call binding the contract method 0xf132d9c8.
//
// Solidity: function isOpTypeSupported(address _teeId, bytes32 _opType) view returns(bool)
func (_TeeRegistry *TeeRegistryCallerSession) IsOpTypeSupported(_teeId common.Address, _opType [32]byte) (bool, error) {
	return _TeeRegistry.Contract.IsOpTypeSupported(&_TeeRegistry.CallOpts, _teeId, _opType)
}

// ConfirmAvailability is a paid mutator transaction binding the contract method 0xe5456585.
//
// Solidity: function confirmAvailability((bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactor) ConfirmAvailability(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "confirmAvailability", _proof)
}

// ConfirmAvailability is a paid mutator transaction binding the contract method 0xe5456585.
//
// Solidity: function confirmAvailability((bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) returns()
func (_TeeRegistry *TeeRegistrySession) ConfirmAvailability(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ConfirmAvailability(&_TeeRegistry.TransactOpts, _proof)
}

// ConfirmAvailability is a paid mutator transaction binding the contract method 0xe5456585.
//
// Solidity: function confirmAvailability((bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ConfirmAvailability(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ConfirmAvailability(&_TeeRegistry.TransactOpts, _proof)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x80f2abe3.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (_TeeRegistry *TeeRegistryTransactor) ConfirmOwnership(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "confirmOwnership", _teeId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x80f2abe3.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (_TeeRegistry *TeeRegistrySession) ConfirmOwnership(_teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ConfirmOwnership(&_TeeRegistry.TransactOpts, _teeId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x80f2abe3.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ConfirmOwnership(_teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ConfirmOwnership(&_TeeRegistry.TransactOpts, _teeId)
}

// ConfirmReplicate is a paid mutator transaction binding the contract method 0xc5af5b83.
//
// Solidity: function confirmReplicate(address _newTeeId, (bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactor) ConfirmReplicate(opts *bind.TransactOpts, _newTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "confirmReplicate", _newTeeId, _proof)
}

// ConfirmReplicate is a paid mutator transaction binding the contract method 0xc5af5b83.
//
// Solidity: function confirmReplicate(address _newTeeId, (bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) returns()
func (_TeeRegistry *TeeRegistrySession) ConfirmReplicate(_newTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ConfirmReplicate(&_TeeRegistry.TransactOpts, _newTeeId, _proof)
}

// ConfirmReplicate is a paid mutator transaction binding the contract method 0xc5af5b83.
//
// Solidity: function confirmReplicate(address _newTeeId, (bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ConfirmReplicate(_newTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ConfirmReplicate(&_TeeRegistry.TransactOpts, _newTeeId, _proof)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _teeId) returns()
func (_TeeRegistry *TeeRegistryTransactor) Pause(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "pause", _teeId)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _teeId) returns()
func (_TeeRegistry *TeeRegistrySession) Pause(_teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.Pause(&_TeeRegistry.TransactOpts, _teeId)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _teeId) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) Pause(_teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.Pause(&_TeeRegistry.TransactOpts, _teeId)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0x2d5f700d.
//
// Solidity: function pauseWithProof((bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactor) PauseWithProof(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "pauseWithProof", _proof)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0x2d5f700d.
//
// Solidity: function pauseWithProof((bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) returns()
func (_TeeRegistry *TeeRegistrySession) PauseWithProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.PauseWithProof(&_TeeRegistry.TransactOpts, _proof)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0x2d5f700d.
//
// Solidity: function pauseWithProof((bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) PauseWithProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.PauseWithProof(&_TeeRegistry.TransactOpts, _proof)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x38195139.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (_TeeRegistry *TeeRegistryTransactor) ProposeNewOwner(opts *bind.TransactOpts, _teeId common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "proposeNewOwner", _teeId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x38195139.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (_TeeRegistry *TeeRegistrySession) ProposeNewOwner(_teeId common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ProposeNewOwner(&_TeeRegistry.TransactOpts, _teeId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x38195139.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ProposeNewOwner(_teeId common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ProposeNewOwner(&_TeeRegistry.TransactOpts, _teeId, _newOwner)
}

// Register is a paid mutator transaction binding the contract method 0xb6442ccf.
//
// Solidity: function register(address _teeId, string _url, bytes32 _codeHash, bytes32 _platform) payable returns()
func (_TeeRegistry *TeeRegistryTransactor) Register(opts *bind.TransactOpts, _teeId common.Address, _url string, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "register", _teeId, _url, _codeHash, _platform)
}

// Register is a paid mutator transaction binding the contract method 0xb6442ccf.
//
// Solidity: function register(address _teeId, string _url, bytes32 _codeHash, bytes32 _platform) payable returns()
func (_TeeRegistry *TeeRegistrySession) Register(_teeId common.Address, _url string, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.Register(&_TeeRegistry.TransactOpts, _teeId, _url, _codeHash, _platform)
}

// Register is a paid mutator transaction binding the contract method 0xb6442ccf.
//
// Solidity: function register(address _teeId, string _url, bytes32 _codeHash, bytes32 _platform) payable returns()
func (_TeeRegistry *TeeRegistryTransactorSession) Register(_teeId common.Address, _url string, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.Register(&_TeeRegistry.TransactOpts, _teeId, _url, _codeHash, _platform)
}

// ReplicateFrom is a paid mutator transaction binding the contract method 0xfab5e606.
//
// Solidity: function replicateFrom(address _oldTeeId, (bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) payable returns()
func (_TeeRegistry *TeeRegistryTransactor) ReplicateFrom(opts *bind.TransactOpts, _oldTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "replicateFrom", _oldTeeId, _proof)
}

// ReplicateFrom is a paid mutator transaction binding the contract method 0xfab5e606.
//
// Solidity: function replicateFrom(address _oldTeeId, (bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) payable returns()
func (_TeeRegistry *TeeRegistrySession) ReplicateFrom(_oldTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ReplicateFrom(&_TeeRegistry.TransactOpts, _oldTeeId, _proof)
}

// ReplicateFrom is a paid mutator transaction binding the contract method 0xfab5e606.
//
// Solidity: function replicateFrom(address _oldTeeId, (bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) payable returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ReplicateFrom(_oldTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ReplicateFrom(&_TeeRegistry.TransactOpts, _oldTeeId, _proof)
}

// ToPauseForUpgrade is a paid mutator transaction binding the contract method 0xe143ca78.
//
// Solidity: function toPauseForUpgrade(address _teeId) payable returns()
func (_TeeRegistry *TeeRegistryTransactor) ToPauseForUpgrade(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "toPauseForUpgrade", _teeId)
}

// ToPauseForUpgrade is a paid mutator transaction binding the contract method 0xe143ca78.
//
// Solidity: function toPauseForUpgrade(address _teeId) payable returns()
func (_TeeRegistry *TeeRegistrySession) ToPauseForUpgrade(_teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ToPauseForUpgrade(&_TeeRegistry.TransactOpts, _teeId)
}

// ToPauseForUpgrade is a paid mutator transaction binding the contract method 0xe143ca78.
//
// Solidity: function toPauseForUpgrade(address _teeId) payable returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ToPauseForUpgrade(_teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ToPauseForUpgrade(&_TeeRegistry.TransactOpts, _teeId)
}

// ToProduction is a paid mutator transaction binding the contract method 0x7f98637c.
//
// Solidity: function toProduction((bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactor) ToProduction(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "toProduction", _proof)
}

// ToProduction is a paid mutator transaction binding the contract method 0x7f98637c.
//
// Solidity: function toProduction((bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) returns()
func (_TeeRegistry *TeeRegistrySession) ToProduction(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ToProduction(&_TeeRegistry.TransactOpts, _proof)
}

// ToProduction is a paid mutator transaction binding the contract method 0x7f98637c.
//
// Solidity: function toProduction((bytes,(uint8,bytes32,bytes32)[],(bytes32,bytes32,uint16,uint64,((address,address,string,bytes32,bytes32),uint24),(uint8))) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ToProduction(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ToProduction(&_TeeRegistry.TransactOpts, _proof)
}

// TriggerAvailabilityCheck is a paid mutator transaction binding the contract method 0xa1751302.
//
// Solidity: function triggerAvailabilityCheck(address _teeId, address _testOnTeeId) payable returns()
func (_TeeRegistry *TeeRegistryTransactor) TriggerAvailabilityCheck(opts *bind.TransactOpts, _teeId common.Address, _testOnTeeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "triggerAvailabilityCheck", _teeId, _testOnTeeId)
}

// TriggerAvailabilityCheck is a paid mutator transaction binding the contract method 0xa1751302.
//
// Solidity: function triggerAvailabilityCheck(address _teeId, address _testOnTeeId) payable returns()
func (_TeeRegistry *TeeRegistrySession) TriggerAvailabilityCheck(_teeId common.Address, _testOnTeeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.TriggerAvailabilityCheck(&_TeeRegistry.TransactOpts, _teeId, _testOnTeeId)
}

// TriggerAvailabilityCheck is a paid mutator transaction binding the contract method 0xa1751302.
//
// Solidity: function triggerAvailabilityCheck(address _teeId, address _testOnTeeId) payable returns()
func (_TeeRegistry *TeeRegistryTransactorSession) TriggerAvailabilityCheck(_teeId common.Address, _testOnTeeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.TriggerAvailabilityCheck(&_TeeRegistry.TransactOpts, _teeId, _testOnTeeId)
}

// TeeRegistryAvailabilityCheckValidityExtendedIterator is returned from FilterAvailabilityCheckValidityExtended and is used to iterate over the raw logs and unpacked data for AvailabilityCheckValidityExtended events raised by the TeeRegistry contract.
type TeeRegistryAvailabilityCheckValidityExtendedIterator struct {
	Event *TeeRegistryAvailabilityCheckValidityExtended // Event containing the contract specifics and raw log

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
func (it *TeeRegistryAvailabilityCheckValidityExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryAvailabilityCheckValidityExtended)
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
		it.Event = new(TeeRegistryAvailabilityCheckValidityExtended)
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
func (it *TeeRegistryAvailabilityCheckValidityExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryAvailabilityCheckValidityExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryAvailabilityCheckValidityExtended represents a AvailabilityCheckValidityExtended event raised by the TeeRegistry contract.
type TeeRegistryAvailabilityCheckValidityExtended struct {
	TeeId common.Address
	EndTs *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAvailabilityCheckValidityExtended is a free log retrieval operation binding the contract event 0x7e7a3ec0073c8e220c8a849aaa8a363fb1171a59e8d42f0ca56378df1bf4388b.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, uint256 endTs)
func (_TeeRegistry *TeeRegistryFilterer) FilterAvailabilityCheckValidityExtended(opts *bind.FilterOpts, teeId []common.Address) (*TeeRegistryAvailabilityCheckValidityExtendedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "AvailabilityCheckValidityExtended", teeIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryAvailabilityCheckValidityExtendedIterator{contract: _TeeRegistry.contract, event: "AvailabilityCheckValidityExtended", logs: logs, sub: sub}, nil
}

// WatchAvailabilityCheckValidityExtended is a free log subscription operation binding the contract event 0x7e7a3ec0073c8e220c8a849aaa8a363fb1171a59e8d42f0ca56378df1bf4388b.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, uint256 endTs)
func (_TeeRegistry *TeeRegistryFilterer) WatchAvailabilityCheckValidityExtended(opts *bind.WatchOpts, sink chan<- *TeeRegistryAvailabilityCheckValidityExtended, teeId []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "AvailabilityCheckValidityExtended", teeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryAvailabilityCheckValidityExtended)
				if err := _TeeRegistry.contract.UnpackLog(event, "AvailabilityCheckValidityExtended", log); err != nil {
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

// ParseAvailabilityCheckValidityExtended is a log parse operation binding the contract event 0x7e7a3ec0073c8e220c8a849aaa8a363fb1171a59e8d42f0ca56378df1bf4388b.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, uint256 endTs)
func (_TeeRegistry *TeeRegistryFilterer) ParseAvailabilityCheckValidityExtended(log types.Log) (*TeeRegistryAvailabilityCheckValidityExtended, error) {
	event := new(TeeRegistryAvailabilityCheckValidityExtended)
	if err := _TeeRegistry.contract.UnpackLog(event, "AvailabilityCheckValidityExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
