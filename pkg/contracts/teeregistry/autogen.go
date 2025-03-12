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

// TeeRegistryMetaData contains all meta data concerning the TeeRegistry contract.
var TeeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pauseBeforeUpgradeMinDurationSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_availabilityCheckValidityDurationSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSupportedVersion\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AVAILABILITY_CHECK\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REG_OP_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPLICATE_FROM\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TO_PAUSE_FOR_UPGRADE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_version\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_opTypes\",\"type\":\"bytes32[]\"}],\"name\":\"addNewTeeVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_backupTeeIds\",\"type\":\"address[]\"}],\"name\":\"arePlatformsCompatible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availabilityCheckValidityDurationSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newTeeId\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"confirmReplicate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"}],\"name\":\"getCodeHashVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_version\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachine\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structITeeRegistry.TeeMachine\",\"name\":\"_teeMachine\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineStatus\",\"outputs\":[{\"internalType\":\"enumITeeRegistry.TeeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_version\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineWithAttestationData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"_teeMachine\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_version\",\"type\":\"uint256\"}],\"name\":\"getVersionInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_opTypes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"}],\"name\":\"isOpTypeSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSupportedVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseBeforeUpgradeMinDurationSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"pauseWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"name\":\"proposedTeeOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relay\",\"outputs\":[{\"internalType\":\"contractIRelay\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newTeeId\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"replicateFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldTeeId\",\"type\":\"address\"}],\"name\":\"replications\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newTeeId\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_availabilityCheckValidityDurationSeconds\",\"type\":\"uint256\"}],\"name\":\"setAvailabilityCheckValidityDurationSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSupportedVersion\",\"type\":\"uint256\"}],\"name\":\"setMinSupportedVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pauseBeforeUpgradeMinDurationSeconds\",\"type\":\"uint256\"}],\"name\":\"setPauseBeforeUpgradeMinDurationSeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeFeeCalculator\",\"outputs\":[{\"internalType\":\"contractITeeFeeCalculator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeInstructions\",\"outputs\":[{\"internalType\":\"contractITeeInstructions\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"toPauseForUpgrade\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"toProduction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// AVAILABILITYCHECK is a free data retrieval call binding the contract method 0xad30fbcd.
//
// Solidity: function AVAILABILITY_CHECK() view returns(bytes32)
func (_TeeRegistry *TeeRegistryCaller) AVAILABILITYCHECK(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "AVAILABILITY_CHECK")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AVAILABILITYCHECK is a free data retrieval call binding the contract method 0xad30fbcd.
//
// Solidity: function AVAILABILITY_CHECK() view returns(bytes32)
func (_TeeRegistry *TeeRegistrySession) AVAILABILITYCHECK() ([32]byte, error) {
	return _TeeRegistry.Contract.AVAILABILITYCHECK(&_TeeRegistry.CallOpts)
}

// AVAILABILITYCHECK is a free data retrieval call binding the contract method 0xad30fbcd.
//
// Solidity: function AVAILABILITY_CHECK() view returns(bytes32)
func (_TeeRegistry *TeeRegistryCallerSession) AVAILABILITYCHECK() ([32]byte, error) {
	return _TeeRegistry.Contract.AVAILABILITYCHECK(&_TeeRegistry.CallOpts)
}

// REGOPTYPE is a free data retrieval call binding the contract method 0xf468b2eb.
//
// Solidity: function REG_OP_TYPE() view returns(bytes32)
func (_TeeRegistry *TeeRegistryCaller) REGOPTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "REG_OP_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGOPTYPE is a free data retrieval call binding the contract method 0xf468b2eb.
//
// Solidity: function REG_OP_TYPE() view returns(bytes32)
func (_TeeRegistry *TeeRegistrySession) REGOPTYPE() ([32]byte, error) {
	return _TeeRegistry.Contract.REGOPTYPE(&_TeeRegistry.CallOpts)
}

// REGOPTYPE is a free data retrieval call binding the contract method 0xf468b2eb.
//
// Solidity: function REG_OP_TYPE() view returns(bytes32)
func (_TeeRegistry *TeeRegistryCallerSession) REGOPTYPE() ([32]byte, error) {
	return _TeeRegistry.Contract.REGOPTYPE(&_TeeRegistry.CallOpts)
}

// REPLICATEFROM is a free data retrieval call binding the contract method 0x5b24181e.
//
// Solidity: function REPLICATE_FROM() view returns(bytes32)
func (_TeeRegistry *TeeRegistryCaller) REPLICATEFROM(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "REPLICATE_FROM")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPLICATEFROM is a free data retrieval call binding the contract method 0x5b24181e.
//
// Solidity: function REPLICATE_FROM() view returns(bytes32)
func (_TeeRegistry *TeeRegistrySession) REPLICATEFROM() ([32]byte, error) {
	return _TeeRegistry.Contract.REPLICATEFROM(&_TeeRegistry.CallOpts)
}

// REPLICATEFROM is a free data retrieval call binding the contract method 0x5b24181e.
//
// Solidity: function REPLICATE_FROM() view returns(bytes32)
func (_TeeRegistry *TeeRegistryCallerSession) REPLICATEFROM() ([32]byte, error) {
	return _TeeRegistry.Contract.REPLICATEFROM(&_TeeRegistry.CallOpts)
}

// TOPAUSEFORUPGRADE is a free data retrieval call binding the contract method 0x143db3e0.
//
// Solidity: function TO_PAUSE_FOR_UPGRADE() view returns(bytes32)
func (_TeeRegistry *TeeRegistryCaller) TOPAUSEFORUPGRADE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "TO_PAUSE_FOR_UPGRADE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOPAUSEFORUPGRADE is a free data retrieval call binding the contract method 0x143db3e0.
//
// Solidity: function TO_PAUSE_FOR_UPGRADE() view returns(bytes32)
func (_TeeRegistry *TeeRegistrySession) TOPAUSEFORUPGRADE() ([32]byte, error) {
	return _TeeRegistry.Contract.TOPAUSEFORUPGRADE(&_TeeRegistry.CallOpts)
}

// TOPAUSEFORUPGRADE is a free data retrieval call binding the contract method 0x143db3e0.
//
// Solidity: function TO_PAUSE_FOR_UPGRADE() view returns(bytes32)
func (_TeeRegistry *TeeRegistryCallerSession) TOPAUSEFORUPGRADE() ([32]byte, error) {
	return _TeeRegistry.Contract.TOPAUSEFORUPGRADE(&_TeeRegistry.CallOpts)
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

// AvailabilityCheckValidityDurationSeconds is a free data retrieval call binding the contract method 0x018666a1.
//
// Solidity: function availabilityCheckValidityDurationSeconds() view returns(uint256)
func (_TeeRegistry *TeeRegistryCaller) AvailabilityCheckValidityDurationSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "availabilityCheckValidityDurationSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailabilityCheckValidityDurationSeconds is a free data retrieval call binding the contract method 0x018666a1.
//
// Solidity: function availabilityCheckValidityDurationSeconds() view returns(uint256)
func (_TeeRegistry *TeeRegistrySession) AvailabilityCheckValidityDurationSeconds() (*big.Int, error) {
	return _TeeRegistry.Contract.AvailabilityCheckValidityDurationSeconds(&_TeeRegistry.CallOpts)
}

// AvailabilityCheckValidityDurationSeconds is a free data retrieval call binding the contract method 0x018666a1.
//
// Solidity: function availabilityCheckValidityDurationSeconds() view returns(uint256)
func (_TeeRegistry *TeeRegistryCallerSession) AvailabilityCheckValidityDurationSeconds() (*big.Int, error) {
	return _TeeRegistry.Contract.AvailabilityCheckValidityDurationSeconds(&_TeeRegistry.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeRegistry *TeeRegistryCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeRegistry *TeeRegistrySession) FlareSystemsManager() (common.Address, error) {
	return _TeeRegistry.Contract.FlareSystemsManager(&_TeeRegistry.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeRegistry *TeeRegistryCallerSession) FlareSystemsManager() (common.Address, error) {
	return _TeeRegistry.Contract.FlareSystemsManager(&_TeeRegistry.CallOpts)
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

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeRegistry *TeeRegistryCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeRegistry *TeeRegistrySession) GetAddressUpdater() (common.Address, error) {
	return _TeeRegistry.Contract.GetAddressUpdater(&_TeeRegistry.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeRegistry *TeeRegistryCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeRegistry.Contract.GetAddressUpdater(&_TeeRegistry.CallOpts)
}

// GetCodeHashVersion is a free data retrieval call binding the contract method 0x0cca1d21.
//
// Solidity: function getCodeHashVersion(bytes32 _codeHash) view returns(uint256 _version)
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
// Solidity: function getCodeHashVersion(bytes32 _codeHash) view returns(uint256 _version)
func (_TeeRegistry *TeeRegistrySession) GetCodeHashVersion(_codeHash [32]byte) (*big.Int, error) {
	return _TeeRegistry.Contract.GetCodeHashVersion(&_TeeRegistry.CallOpts, _codeHash)
}

// GetCodeHashVersion is a free data retrieval call binding the contract method 0x0cca1d21.
//
// Solidity: function getCodeHashVersion(bytes32 _codeHash) view returns(uint256 _version)
func (_TeeRegistry *TeeRegistryCallerSession) GetCodeHashVersion(_codeHash [32]byte) (*big.Int, error) {
	return _TeeRegistry.Contract.GetCodeHashVersion(&_TeeRegistry.CallOpts, _codeHash)
}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string) _teeMachine)
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
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string) _teeMachine)
func (_TeeRegistry *TeeRegistrySession) GetTeeMachine(_teeId common.Address) (ITeeRegistryTeeMachine, error) {
	return _TeeRegistry.Contract.GetTeeMachine(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string) _teeMachine)
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
// Solidity: function getTeeMachineVersion(address _teeId) view returns(uint256 _version)
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
// Solidity: function getTeeMachineVersion(address _teeId) view returns(uint256 _version)
func (_TeeRegistry *TeeRegistrySession) GetTeeMachineVersion(_teeId common.Address) (*big.Int, error) {
	return _TeeRegistry.Contract.GetTeeMachineVersion(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineVersion is a free data retrieval call binding the contract method 0x6a1fbb80.
//
// Solidity: function getTeeMachineVersion(address _teeId) view returns(uint256 _version)
func (_TeeRegistry *TeeRegistryCallerSession) GetTeeMachineVersion(_teeId common.Address) (*big.Int, error) {
	return _TeeRegistry.Contract.GetTeeMachineVersion(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32) _teeMachine)
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
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32) _teeMachine)
func (_TeeRegistry *TeeRegistrySession) GetTeeMachineWithAttestationData(_teeId common.Address) (ITeeRegistryTeeMachineWithAttestationData, error) {
	return _TeeRegistry.Contract.GetTeeMachineWithAttestationData(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32) _teeMachine)
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

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeRegistry *TeeRegistryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeRegistry *TeeRegistrySession) Governance() (common.Address, error) {
	return _TeeRegistry.Contract.Governance(&_TeeRegistry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeRegistry *TeeRegistryCallerSession) Governance() (common.Address, error) {
	return _TeeRegistry.Contract.Governance(&_TeeRegistry.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeRegistry *TeeRegistryCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeRegistry *TeeRegistrySession) GovernanceSettings() (common.Address, error) {
	return _TeeRegistry.Contract.GovernanceSettings(&_TeeRegistry.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeRegistry *TeeRegistryCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeeRegistry.Contract.GovernanceSettings(&_TeeRegistry.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeRegistry *TeeRegistryCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeRegistry *TeeRegistrySession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeRegistry.Contract.IsExecutor(&_TeeRegistry.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeRegistry *TeeRegistryCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeRegistry.Contract.IsExecutor(&_TeeRegistry.CallOpts, _address)
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

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_TeeRegistry *TeeRegistryCaller) LatestVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "latestVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_TeeRegistry *TeeRegistrySession) LatestVersion() (*big.Int, error) {
	return _TeeRegistry.Contract.LatestVersion(&_TeeRegistry.CallOpts)
}

// LatestVersion is a free data retrieval call binding the contract method 0xc07f47d4.
//
// Solidity: function latestVersion() view returns(uint256)
func (_TeeRegistry *TeeRegistryCallerSession) LatestVersion() (*big.Int, error) {
	return _TeeRegistry.Contract.LatestVersion(&_TeeRegistry.CallOpts)
}

// MinSupportedVersion is a free data retrieval call binding the contract method 0x57b28588.
//
// Solidity: function minSupportedVersion() view returns(uint256)
func (_TeeRegistry *TeeRegistryCaller) MinSupportedVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "minSupportedVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSupportedVersion is a free data retrieval call binding the contract method 0x57b28588.
//
// Solidity: function minSupportedVersion() view returns(uint256)
func (_TeeRegistry *TeeRegistrySession) MinSupportedVersion() (*big.Int, error) {
	return _TeeRegistry.Contract.MinSupportedVersion(&_TeeRegistry.CallOpts)
}

// MinSupportedVersion is a free data retrieval call binding the contract method 0x57b28588.
//
// Solidity: function minSupportedVersion() view returns(uint256)
func (_TeeRegistry *TeeRegistryCallerSession) MinSupportedVersion() (*big.Int, error) {
	return _TeeRegistry.Contract.MinSupportedVersion(&_TeeRegistry.CallOpts)
}

// PauseBeforeUpgradeMinDurationSeconds is a free data retrieval call binding the contract method 0x7585fed0.
//
// Solidity: function pauseBeforeUpgradeMinDurationSeconds() view returns(uint256)
func (_TeeRegistry *TeeRegistryCaller) PauseBeforeUpgradeMinDurationSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "pauseBeforeUpgradeMinDurationSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PauseBeforeUpgradeMinDurationSeconds is a free data retrieval call binding the contract method 0x7585fed0.
//
// Solidity: function pauseBeforeUpgradeMinDurationSeconds() view returns(uint256)
func (_TeeRegistry *TeeRegistrySession) PauseBeforeUpgradeMinDurationSeconds() (*big.Int, error) {
	return _TeeRegistry.Contract.PauseBeforeUpgradeMinDurationSeconds(&_TeeRegistry.CallOpts)
}

// PauseBeforeUpgradeMinDurationSeconds is a free data retrieval call binding the contract method 0x7585fed0.
//
// Solidity: function pauseBeforeUpgradeMinDurationSeconds() view returns(uint256)
func (_TeeRegistry *TeeRegistryCallerSession) PauseBeforeUpgradeMinDurationSeconds() (*big.Int, error) {
	return _TeeRegistry.Contract.PauseBeforeUpgradeMinDurationSeconds(&_TeeRegistry.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeRegistry *TeeRegistryCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeRegistry *TeeRegistrySession) ProductionMode() (bool, error) {
	return _TeeRegistry.Contract.ProductionMode(&_TeeRegistry.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeRegistry *TeeRegistryCallerSession) ProductionMode() (bool, error) {
	return _TeeRegistry.Contract.ProductionMode(&_TeeRegistry.CallOpts)
}

// ProposedTeeOwner is a free data retrieval call binding the contract method 0x574d48e7.
//
// Solidity: function proposedTeeOwner(address teeId) view returns(address)
func (_TeeRegistry *TeeRegistryCaller) ProposedTeeOwner(opts *bind.CallOpts, teeId common.Address) (common.Address, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "proposedTeeOwner", teeId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposedTeeOwner is a free data retrieval call binding the contract method 0x574d48e7.
//
// Solidity: function proposedTeeOwner(address teeId) view returns(address)
func (_TeeRegistry *TeeRegistrySession) ProposedTeeOwner(teeId common.Address) (common.Address, error) {
	return _TeeRegistry.Contract.ProposedTeeOwner(&_TeeRegistry.CallOpts, teeId)
}

// ProposedTeeOwner is a free data retrieval call binding the contract method 0x574d48e7.
//
// Solidity: function proposedTeeOwner(address teeId) view returns(address)
func (_TeeRegistry *TeeRegistryCallerSession) ProposedTeeOwner(teeId common.Address) (common.Address, error) {
	return _TeeRegistry.Contract.ProposedTeeOwner(&_TeeRegistry.CallOpts, teeId)
}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_TeeRegistry *TeeRegistryCaller) Relay(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "relay")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_TeeRegistry *TeeRegistrySession) Relay() (common.Address, error) {
	return _TeeRegistry.Contract.Relay(&_TeeRegistry.CallOpts)
}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_TeeRegistry *TeeRegistryCallerSession) Relay() (common.Address, error) {
	return _TeeRegistry.Contract.Relay(&_TeeRegistry.CallOpts)
}

// Replications is a free data retrieval call binding the contract method 0xd436b5e9.
//
// Solidity: function replications(address oldTeeId) view returns(address newTeeId)
func (_TeeRegistry *TeeRegistryCaller) Replications(opts *bind.CallOpts, oldTeeId common.Address) (common.Address, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "replications", oldTeeId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Replications is a free data retrieval call binding the contract method 0xd436b5e9.
//
// Solidity: function replications(address oldTeeId) view returns(address newTeeId)
func (_TeeRegistry *TeeRegistrySession) Replications(oldTeeId common.Address) (common.Address, error) {
	return _TeeRegistry.Contract.Replications(&_TeeRegistry.CallOpts, oldTeeId)
}

// Replications is a free data retrieval call binding the contract method 0xd436b5e9.
//
// Solidity: function replications(address oldTeeId) view returns(address newTeeId)
func (_TeeRegistry *TeeRegistryCallerSession) Replications(oldTeeId common.Address) (common.Address, error) {
	return _TeeRegistry.Contract.Replications(&_TeeRegistry.CallOpts, oldTeeId)
}

// TeeFeeCalculator is a free data retrieval call binding the contract method 0xde4ae7a9.
//
// Solidity: function teeFeeCalculator() view returns(address)
func (_TeeRegistry *TeeRegistryCaller) TeeFeeCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "teeFeeCalculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeFeeCalculator is a free data retrieval call binding the contract method 0xde4ae7a9.
//
// Solidity: function teeFeeCalculator() view returns(address)
func (_TeeRegistry *TeeRegistrySession) TeeFeeCalculator() (common.Address, error) {
	return _TeeRegistry.Contract.TeeFeeCalculator(&_TeeRegistry.CallOpts)
}

// TeeFeeCalculator is a free data retrieval call binding the contract method 0xde4ae7a9.
//
// Solidity: function teeFeeCalculator() view returns(address)
func (_TeeRegistry *TeeRegistryCallerSession) TeeFeeCalculator() (common.Address, error) {
	return _TeeRegistry.Contract.TeeFeeCalculator(&_TeeRegistry.CallOpts)
}

// TeeInstructions is a free data retrieval call binding the contract method 0xa3f000ce.
//
// Solidity: function teeInstructions() view returns(address)
func (_TeeRegistry *TeeRegistryCaller) TeeInstructions(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "teeInstructions")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeInstructions is a free data retrieval call binding the contract method 0xa3f000ce.
//
// Solidity: function teeInstructions() view returns(address)
func (_TeeRegistry *TeeRegistrySession) TeeInstructions() (common.Address, error) {
	return _TeeRegistry.Contract.TeeInstructions(&_TeeRegistry.CallOpts)
}

// TeeInstructions is a free data retrieval call binding the contract method 0xa3f000ce.
//
// Solidity: function teeInstructions() view returns(address)
func (_TeeRegistry *TeeRegistryCallerSession) TeeInstructions() (common.Address, error) {
	return _TeeRegistry.Contract.TeeInstructions(&_TeeRegistry.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeRegistry *TeeRegistryCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "timelockedCalls", selector)

	outstruct := new(struct {
		AllowedAfterTimestamp *big.Int
		EncodedCall           []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowedAfterTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EncodedCall = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeRegistry *TeeRegistrySession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeRegistry.Contract.TimelockedCalls(&_TeeRegistry.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeRegistry *TeeRegistryCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeRegistry.Contract.TimelockedCalls(&_TeeRegistry.CallOpts, selector)
}

// AddNewTeeVersion is a paid mutator transaction binding the contract method 0x3ca61eb7.
//
// Solidity: function addNewTeeVersion(uint256 _version, bytes32 _codeHash, bytes32[] _platforms, bytes32[] _opTypes) returns()
func (_TeeRegistry *TeeRegistryTransactor) AddNewTeeVersion(opts *bind.TransactOpts, _version *big.Int, _codeHash [32]byte, _platforms [][32]byte, _opTypes [][32]byte) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "addNewTeeVersion", _version, _codeHash, _platforms, _opTypes)
}

// AddNewTeeVersion is a paid mutator transaction binding the contract method 0x3ca61eb7.
//
// Solidity: function addNewTeeVersion(uint256 _version, bytes32 _codeHash, bytes32[] _platforms, bytes32[] _opTypes) returns()
func (_TeeRegistry *TeeRegistrySession) AddNewTeeVersion(_version *big.Int, _codeHash [32]byte, _platforms [][32]byte, _opTypes [][32]byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.AddNewTeeVersion(&_TeeRegistry.TransactOpts, _version, _codeHash, _platforms, _opTypes)
}

// AddNewTeeVersion is a paid mutator transaction binding the contract method 0x3ca61eb7.
//
// Solidity: function addNewTeeVersion(uint256 _version, bytes32 _codeHash, bytes32[] _platforms, bytes32[] _opTypes) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) AddNewTeeVersion(_version *big.Int, _codeHash [32]byte, _platforms [][32]byte, _opTypes [][32]byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.AddNewTeeVersion(&_TeeRegistry.TransactOpts, _version, _codeHash, _platforms, _opTypes)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeRegistry *TeeRegistryTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeRegistry *TeeRegistrySession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.CancelGovernanceCall(&_TeeRegistry.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.CancelGovernanceCall(&_TeeRegistry.TransactOpts, _selector)
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

// ConfirmReplicate is a paid mutator transaction binding the contract method 0x864cc798.
//
// Solidity: function confirmReplicate(address _oldTeeId, address _newTeeId, bytes ) payable returns()
func (_TeeRegistry *TeeRegistryTransactor) ConfirmReplicate(opts *bind.TransactOpts, _oldTeeId common.Address, _newTeeId common.Address, arg2 []byte) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "confirmReplicate", _oldTeeId, _newTeeId, arg2)
}

// ConfirmReplicate is a paid mutator transaction binding the contract method 0x864cc798.
//
// Solidity: function confirmReplicate(address _oldTeeId, address _newTeeId, bytes ) payable returns()
func (_TeeRegistry *TeeRegistrySession) ConfirmReplicate(_oldTeeId common.Address, _newTeeId common.Address, arg2 []byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ConfirmReplicate(&_TeeRegistry.TransactOpts, _oldTeeId, _newTeeId, arg2)
}

// ConfirmReplicate is a paid mutator transaction binding the contract method 0x864cc798.
//
// Solidity: function confirmReplicate(address _oldTeeId, address _newTeeId, bytes ) payable returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ConfirmReplicate(_oldTeeId common.Address, _newTeeId common.Address, arg2 []byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ConfirmReplicate(&_TeeRegistry.TransactOpts, _oldTeeId, _newTeeId, arg2)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeRegistry *TeeRegistryTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeRegistry *TeeRegistrySession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ExecuteGovernanceCall(&_TeeRegistry.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ExecuteGovernanceCall(&_TeeRegistry.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeRegistry *TeeRegistryTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeRegistry *TeeRegistrySession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.Initialise(&_TeeRegistry.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.Initialise(&_TeeRegistry.TransactOpts, _governanceSettings, _initialGovernance)
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

// PauseWithProof is a paid mutator transaction binding the contract method 0x396de396.
//
// Solidity: function pauseWithProof(address _teeId, bytes ) returns()
func (_TeeRegistry *TeeRegistryTransactor) PauseWithProof(opts *bind.TransactOpts, _teeId common.Address, arg1 []byte) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "pauseWithProof", _teeId, arg1)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0x396de396.
//
// Solidity: function pauseWithProof(address _teeId, bytes ) returns()
func (_TeeRegistry *TeeRegistrySession) PauseWithProof(_teeId common.Address, arg1 []byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.PauseWithProof(&_TeeRegistry.TransactOpts, _teeId, arg1)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0x396de396.
//
// Solidity: function pauseWithProof(address _teeId, bytes ) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) PauseWithProof(_teeId common.Address, arg1 []byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.PauseWithProof(&_TeeRegistry.TransactOpts, _teeId, arg1)
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

// ReplicateFrom is a paid mutator transaction binding the contract method 0x5fbf767b.
//
// Solidity: function replicateFrom(address _oldTeeId, address _newTeeId, bytes ) payable returns()
func (_TeeRegistry *TeeRegistryTransactor) ReplicateFrom(opts *bind.TransactOpts, _oldTeeId common.Address, _newTeeId common.Address, arg2 []byte) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "replicateFrom", _oldTeeId, _newTeeId, arg2)
}

// ReplicateFrom is a paid mutator transaction binding the contract method 0x5fbf767b.
//
// Solidity: function replicateFrom(address _oldTeeId, address _newTeeId, bytes ) payable returns()
func (_TeeRegistry *TeeRegistrySession) ReplicateFrom(_oldTeeId common.Address, _newTeeId common.Address, arg2 []byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ReplicateFrom(&_TeeRegistry.TransactOpts, _oldTeeId, _newTeeId, arg2)
}

// ReplicateFrom is a paid mutator transaction binding the contract method 0x5fbf767b.
//
// Solidity: function replicateFrom(address _oldTeeId, address _newTeeId, bytes ) payable returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ReplicateFrom(_oldTeeId common.Address, _newTeeId common.Address, arg2 []byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ReplicateFrom(&_TeeRegistry.TransactOpts, _oldTeeId, _newTeeId, arg2)
}

// SetAvailabilityCheckValidityDurationSeconds is a paid mutator transaction binding the contract method 0x22fb83a3.
//
// Solidity: function setAvailabilityCheckValidityDurationSeconds(uint256 _availabilityCheckValidityDurationSeconds) returns()
func (_TeeRegistry *TeeRegistryTransactor) SetAvailabilityCheckValidityDurationSeconds(opts *bind.TransactOpts, _availabilityCheckValidityDurationSeconds *big.Int) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "setAvailabilityCheckValidityDurationSeconds", _availabilityCheckValidityDurationSeconds)
}

// SetAvailabilityCheckValidityDurationSeconds is a paid mutator transaction binding the contract method 0x22fb83a3.
//
// Solidity: function setAvailabilityCheckValidityDurationSeconds(uint256 _availabilityCheckValidityDurationSeconds) returns()
func (_TeeRegistry *TeeRegistrySession) SetAvailabilityCheckValidityDurationSeconds(_availabilityCheckValidityDurationSeconds *big.Int) (*types.Transaction, error) {
	return _TeeRegistry.Contract.SetAvailabilityCheckValidityDurationSeconds(&_TeeRegistry.TransactOpts, _availabilityCheckValidityDurationSeconds)
}

// SetAvailabilityCheckValidityDurationSeconds is a paid mutator transaction binding the contract method 0x22fb83a3.
//
// Solidity: function setAvailabilityCheckValidityDurationSeconds(uint256 _availabilityCheckValidityDurationSeconds) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) SetAvailabilityCheckValidityDurationSeconds(_availabilityCheckValidityDurationSeconds *big.Int) (*types.Transaction, error) {
	return _TeeRegistry.Contract.SetAvailabilityCheckValidityDurationSeconds(&_TeeRegistry.TransactOpts, _availabilityCheckValidityDurationSeconds)
}

// SetMinSupportedVersion is a paid mutator transaction binding the contract method 0x59947aa5.
//
// Solidity: function setMinSupportedVersion(uint256 _minSupportedVersion) returns()
func (_TeeRegistry *TeeRegistryTransactor) SetMinSupportedVersion(opts *bind.TransactOpts, _minSupportedVersion *big.Int) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "setMinSupportedVersion", _minSupportedVersion)
}

// SetMinSupportedVersion is a paid mutator transaction binding the contract method 0x59947aa5.
//
// Solidity: function setMinSupportedVersion(uint256 _minSupportedVersion) returns()
func (_TeeRegistry *TeeRegistrySession) SetMinSupportedVersion(_minSupportedVersion *big.Int) (*types.Transaction, error) {
	return _TeeRegistry.Contract.SetMinSupportedVersion(&_TeeRegistry.TransactOpts, _minSupportedVersion)
}

// SetMinSupportedVersion is a paid mutator transaction binding the contract method 0x59947aa5.
//
// Solidity: function setMinSupportedVersion(uint256 _minSupportedVersion) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) SetMinSupportedVersion(_minSupportedVersion *big.Int) (*types.Transaction, error) {
	return _TeeRegistry.Contract.SetMinSupportedVersion(&_TeeRegistry.TransactOpts, _minSupportedVersion)
}

// SetPauseBeforeUpgradeMinDurationSeconds is a paid mutator transaction binding the contract method 0x5319f918.
//
// Solidity: function setPauseBeforeUpgradeMinDurationSeconds(uint256 _pauseBeforeUpgradeMinDurationSeconds) returns()
func (_TeeRegistry *TeeRegistryTransactor) SetPauseBeforeUpgradeMinDurationSeconds(opts *bind.TransactOpts, _pauseBeforeUpgradeMinDurationSeconds *big.Int) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "setPauseBeforeUpgradeMinDurationSeconds", _pauseBeforeUpgradeMinDurationSeconds)
}

// SetPauseBeforeUpgradeMinDurationSeconds is a paid mutator transaction binding the contract method 0x5319f918.
//
// Solidity: function setPauseBeforeUpgradeMinDurationSeconds(uint256 _pauseBeforeUpgradeMinDurationSeconds) returns()
func (_TeeRegistry *TeeRegistrySession) SetPauseBeforeUpgradeMinDurationSeconds(_pauseBeforeUpgradeMinDurationSeconds *big.Int) (*types.Transaction, error) {
	return _TeeRegistry.Contract.SetPauseBeforeUpgradeMinDurationSeconds(&_TeeRegistry.TransactOpts, _pauseBeforeUpgradeMinDurationSeconds)
}

// SetPauseBeforeUpgradeMinDurationSeconds is a paid mutator transaction binding the contract method 0x5319f918.
//
// Solidity: function setPauseBeforeUpgradeMinDurationSeconds(uint256 _pauseBeforeUpgradeMinDurationSeconds) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) SetPauseBeforeUpgradeMinDurationSeconds(_pauseBeforeUpgradeMinDurationSeconds *big.Int) (*types.Transaction, error) {
	return _TeeRegistry.Contract.SetPauseBeforeUpgradeMinDurationSeconds(&_TeeRegistry.TransactOpts, _pauseBeforeUpgradeMinDurationSeconds)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeRegistry *TeeRegistryTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeRegistry *TeeRegistrySession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeRegistry.Contract.SwitchToProductionMode(&_TeeRegistry.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeRegistry *TeeRegistryTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeRegistry.Contract.SwitchToProductionMode(&_TeeRegistry.TransactOpts)
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

// ToProduction is a paid mutator transaction binding the contract method 0x430e77d1.
//
// Solidity: function toProduction(address _teeId, bytes ) returns()
func (_TeeRegistry *TeeRegistryTransactor) ToProduction(opts *bind.TransactOpts, _teeId common.Address, arg1 []byte) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "toProduction", _teeId, arg1)
}

// ToProduction is a paid mutator transaction binding the contract method 0x430e77d1.
//
// Solidity: function toProduction(address _teeId, bytes ) returns()
func (_TeeRegistry *TeeRegistrySession) ToProduction(_teeId common.Address, arg1 []byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ToProduction(&_TeeRegistry.TransactOpts, _teeId, arg1)
}

// ToProduction is a paid mutator transaction binding the contract method 0x430e77d1.
//
// Solidity: function toProduction(address _teeId, bytes ) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ToProduction(_teeId common.Address, arg1 []byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ToProduction(&_TeeRegistry.TransactOpts, _teeId, arg1)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeRegistry *TeeRegistryTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeRegistry *TeeRegistrySession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.UpdateContractAddresses(&_TeeRegistry.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.UpdateContractAddresses(&_TeeRegistry.TransactOpts, _contractNameHashes, _contractAddresses)
}

// TeeRegistryGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeRegistry contract.
type TeeRegistryGovernanceCallTimelockedIterator struct {
	Event *TeeRegistryGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeeRegistryGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryGovernanceCallTimelocked)
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
		it.Event = new(TeeRegistryGovernanceCallTimelocked)
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
func (it *TeeRegistryGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeRegistry contract.
type TeeRegistryGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeRegistry *TeeRegistryFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeRegistryGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeRegistryGovernanceCallTimelockedIterator{contract: _TeeRegistry.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeRegistry *TeeRegistryFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeRegistryGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryGovernanceCallTimelocked)
				if err := _TeeRegistry.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeRegistry *TeeRegistryFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeRegistryGovernanceCallTimelocked, error) {
	event := new(TeeRegistryGovernanceCallTimelocked)
	if err := _TeeRegistry.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRegistryGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeRegistry contract.
type TeeRegistryGovernanceInitialisedIterator struct {
	Event *TeeRegistryGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeeRegistryGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryGovernanceInitialised)
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
		it.Event = new(TeeRegistryGovernanceInitialised)
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
func (it *TeeRegistryGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryGovernanceInitialised represents a GovernanceInitialised event raised by the TeeRegistry contract.
type TeeRegistryGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeRegistry *TeeRegistryFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeRegistryGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeRegistryGovernanceInitialisedIterator{contract: _TeeRegistry.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeRegistry *TeeRegistryFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeRegistryGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryGovernanceInitialised)
				if err := _TeeRegistry.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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

// ParseGovernanceInitialised is a log parse operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeRegistry *TeeRegistryFilterer) ParseGovernanceInitialised(log types.Log) (*TeeRegistryGovernanceInitialised, error) {
	event := new(TeeRegistryGovernanceInitialised)
	if err := _TeeRegistry.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRegistryGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeeRegistry contract.
type TeeRegistryGovernedProductionModeEnteredIterator struct {
	Event *TeeRegistryGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeeRegistryGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryGovernedProductionModeEntered)
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
		it.Event = new(TeeRegistryGovernedProductionModeEntered)
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
func (it *TeeRegistryGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeeRegistry contract.
type TeeRegistryGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeRegistry *TeeRegistryFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeeRegistryGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeeRegistryGovernedProductionModeEnteredIterator{contract: _TeeRegistry.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeRegistry *TeeRegistryFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeeRegistryGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryGovernedProductionModeEntered)
				if err := _TeeRegistry.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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

// ParseGovernedProductionModeEntered is a log parse operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeRegistry *TeeRegistryFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeeRegistryGovernedProductionModeEntered, error) {
	event := new(TeeRegistryGovernedProductionModeEntered)
	if err := _TeeRegistry.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRegistryTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeeRegistry contract.
type TeeRegistryTimelockedGovernanceCallCanceledIterator struct {
	Event *TeeRegistryTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeeRegistryTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeeRegistryTimelockedGovernanceCallCanceled)
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
func (it *TeeRegistryTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeeRegistry contract.
type TeeRegistryTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeRegistry *TeeRegistryFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeeRegistryTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeeRegistryTimelockedGovernanceCallCanceledIterator{contract: _TeeRegistry.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeRegistry *TeeRegistryFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeeRegistryTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryTimelockedGovernanceCallCanceled)
				if err := _TeeRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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

// ParseTimelockedGovernanceCallCanceled is a log parse operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeRegistry *TeeRegistryFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeeRegistryTimelockedGovernanceCallCanceled, error) {
	event := new(TeeRegistryTimelockedGovernanceCallCanceled)
	if err := _TeeRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRegistryTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeeRegistry contract.
type TeeRegistryTimelockedGovernanceCallExecutedIterator struct {
	Event *TeeRegistryTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeeRegistryTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeeRegistryTimelockedGovernanceCallExecuted)
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
func (it *TeeRegistryTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeeRegistry contract.
type TeeRegistryTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeRegistry *TeeRegistryFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeeRegistryTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeeRegistryTimelockedGovernanceCallExecutedIterator{contract: _TeeRegistry.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeRegistry *TeeRegistryFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeeRegistryTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryTimelockedGovernanceCallExecuted)
				if err := _TeeRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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

// ParseTimelockedGovernanceCallExecuted is a log parse operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeRegistry *TeeRegistryFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeeRegistryTimelockedGovernanceCallExecuted, error) {
	event := new(TeeRegistryTimelockedGovernanceCallExecuted)
	if err := _TeeRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
