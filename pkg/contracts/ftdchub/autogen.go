// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ftdchub

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

// FtdcHubMetaData contains all meta data concerning the FtdcHub contract.
var FtdcHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_minThresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"_defaultNumberOfTees\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CosignersThresholdInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DefaultNumberOfTeesZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinThresholdInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MultipleResponsesPossible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NumberOfTeesAndTeeIdsInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"requestBody\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"AttestationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"defaultNumberOfTees\",\"type\":\"uint8\"}],\"name\":\"DefaultNumberOfTeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minThresholdBIPS\",\"type\":\"uint16\"}],\"name\":\"MinThresholdBIPSSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FTDC_OP_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultNumberOfTees\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftdcRequestFeeConfigurations\",\"outputs\":[{\"internalType\":\"contractIFtdcRequestFeeConfigurations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minThresholdBIPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfTees\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_requestBody\",\"type\":\"bytes\"}],\"name\":\"requestAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_defaultNumberOfTees\",\"type\":\"uint8\"}],\"name\":\"setDefaultNumberOfTees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_minThresholdBIPS\",\"type\":\"uint16\"}],\"name\":\"setMinThresholdBIPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeExtensionRegistry\",\"outputs\":[{\"internalType\":\"contractIITeeExtensionRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeMachineRegistry\",\"outputs\":[{\"internalType\":\"contractITeeMachineRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeReplication\",\"outputs\":[{\"internalType\":\"contractITeeReplication\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FtdcHubABI is the input ABI used to generate the binding from.
// Deprecated: Use FtdcHubMetaData.ABI instead.
var FtdcHubABI = FtdcHubMetaData.ABI

// FtdcHub is an auto generated Go binding around an Ethereum contract.
type FtdcHub struct {
	FtdcHubCaller     // Read-only binding to the contract
	FtdcHubTransactor // Write-only binding to the contract
	FtdcHubFilterer   // Log filterer for contract events
}

// FtdcHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type FtdcHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtdcHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FtdcHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtdcHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FtdcHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FtdcHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FtdcHubSession struct {
	Contract     *FtdcHub          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FtdcHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FtdcHubCallerSession struct {
	Contract *FtdcHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FtdcHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FtdcHubTransactorSession struct {
	Contract     *FtdcHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FtdcHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type FtdcHubRaw struct {
	Contract *FtdcHub // Generic contract binding to access the raw methods on
}

// FtdcHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FtdcHubCallerRaw struct {
	Contract *FtdcHubCaller // Generic read-only contract binding to access the raw methods on
}

// FtdcHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FtdcHubTransactorRaw struct {
	Contract *FtdcHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFtdcHub creates a new instance of FtdcHub, bound to a specific deployed contract.
func NewFtdcHub(address common.Address, backend bind.ContractBackend) (*FtdcHub, error) {
	contract, err := bindFtdcHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FtdcHub{FtdcHubCaller: FtdcHubCaller{contract: contract}, FtdcHubTransactor: FtdcHubTransactor{contract: contract}, FtdcHubFilterer: FtdcHubFilterer{contract: contract}}, nil
}

// NewFtdcHubCaller creates a new read-only instance of FtdcHub, bound to a specific deployed contract.
func NewFtdcHubCaller(address common.Address, caller bind.ContractCaller) (*FtdcHubCaller, error) {
	contract, err := bindFtdcHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FtdcHubCaller{contract: contract}, nil
}

// NewFtdcHubTransactor creates a new write-only instance of FtdcHub, bound to a specific deployed contract.
func NewFtdcHubTransactor(address common.Address, transactor bind.ContractTransactor) (*FtdcHubTransactor, error) {
	contract, err := bindFtdcHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FtdcHubTransactor{contract: contract}, nil
}

// NewFtdcHubFilterer creates a new log filterer instance of FtdcHub, bound to a specific deployed contract.
func NewFtdcHubFilterer(address common.Address, filterer bind.ContractFilterer) (*FtdcHubFilterer, error) {
	contract, err := bindFtdcHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FtdcHubFilterer{contract: contract}, nil
}

// bindFtdcHub binds a generic wrapper to an already deployed contract.
func bindFtdcHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FtdcHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FtdcHub *FtdcHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FtdcHub.Contract.FtdcHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FtdcHub *FtdcHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtdcHub.Contract.FtdcHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FtdcHub *FtdcHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FtdcHub.Contract.FtdcHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FtdcHub *FtdcHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FtdcHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FtdcHub *FtdcHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtdcHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FtdcHub *FtdcHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FtdcHub.Contract.contract.Transact(opts, method, params...)
}

// FTDCOPTYPE is a free data retrieval call binding the contract method 0x9be93414.
//
// Solidity: function FTDC_OP_TYPE() view returns(bytes32)
func (_FtdcHub *FtdcHubCaller) FTDCOPTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "FTDC_OP_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FTDCOPTYPE is a free data retrieval call binding the contract method 0x9be93414.
//
// Solidity: function FTDC_OP_TYPE() view returns(bytes32)
func (_FtdcHub *FtdcHubSession) FTDCOPTYPE() ([32]byte, error) {
	return _FtdcHub.Contract.FTDCOPTYPE(&_FtdcHub.CallOpts)
}

// FTDCOPTYPE is a free data retrieval call binding the contract method 0x9be93414.
//
// Solidity: function FTDC_OP_TYPE() view returns(bytes32)
func (_FtdcHub *FtdcHubCallerSession) FTDCOPTYPE() ([32]byte, error) {
	return _FtdcHub.Contract.FTDCOPTYPE(&_FtdcHub.CallOpts)
}

// PROVE is a free data retrieval call binding the contract method 0x14f20ca6.
//
// Solidity: function PROVE() view returns(bytes32)
func (_FtdcHub *FtdcHubCaller) PROVE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "PROVE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROVE is a free data retrieval call binding the contract method 0x14f20ca6.
//
// Solidity: function PROVE() view returns(bytes32)
func (_FtdcHub *FtdcHubSession) PROVE() ([32]byte, error) {
	return _FtdcHub.Contract.PROVE(&_FtdcHub.CallOpts)
}

// PROVE is a free data retrieval call binding the contract method 0x14f20ca6.
//
// Solidity: function PROVE() view returns(bytes32)
func (_FtdcHub *FtdcHubCallerSession) PROVE() ([32]byte, error) {
	return _FtdcHub.Contract.PROVE(&_FtdcHub.CallOpts)
}

// DefaultNumberOfTees is a free data retrieval call binding the contract method 0x1b92dbb4.
//
// Solidity: function defaultNumberOfTees() view returns(uint8)
func (_FtdcHub *FtdcHubCaller) DefaultNumberOfTees(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "defaultNumberOfTees")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultNumberOfTees is a free data retrieval call binding the contract method 0x1b92dbb4.
//
// Solidity: function defaultNumberOfTees() view returns(uint8)
func (_FtdcHub *FtdcHubSession) DefaultNumberOfTees() (uint8, error) {
	return _FtdcHub.Contract.DefaultNumberOfTees(&_FtdcHub.CallOpts)
}

// DefaultNumberOfTees is a free data retrieval call binding the contract method 0x1b92dbb4.
//
// Solidity: function defaultNumberOfTees() view returns(uint8)
func (_FtdcHub *FtdcHubCallerSession) DefaultNumberOfTees() (uint8, error) {
	return _FtdcHub.Contract.DefaultNumberOfTees(&_FtdcHub.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_FtdcHub *FtdcHubCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_FtdcHub *FtdcHubSession) FlareSystemsManager() (common.Address, error) {
	return _FtdcHub.Contract.FlareSystemsManager(&_FtdcHub.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_FtdcHub *FtdcHubCallerSession) FlareSystemsManager() (common.Address, error) {
	return _FtdcHub.Contract.FlareSystemsManager(&_FtdcHub.CallOpts)
}

// FtdcRequestFeeConfigurations is a free data retrieval call binding the contract method 0x2d581eb1.
//
// Solidity: function ftdcRequestFeeConfigurations() view returns(address)
func (_FtdcHub *FtdcHubCaller) FtdcRequestFeeConfigurations(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "ftdcRequestFeeConfigurations")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FtdcRequestFeeConfigurations is a free data retrieval call binding the contract method 0x2d581eb1.
//
// Solidity: function ftdcRequestFeeConfigurations() view returns(address)
func (_FtdcHub *FtdcHubSession) FtdcRequestFeeConfigurations() (common.Address, error) {
	return _FtdcHub.Contract.FtdcRequestFeeConfigurations(&_FtdcHub.CallOpts)
}

// FtdcRequestFeeConfigurations is a free data retrieval call binding the contract method 0x2d581eb1.
//
// Solidity: function ftdcRequestFeeConfigurations() view returns(address)
func (_FtdcHub *FtdcHubCallerSession) FtdcRequestFeeConfigurations() (common.Address, error) {
	return _FtdcHub.Contract.FtdcRequestFeeConfigurations(&_FtdcHub.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FtdcHub *FtdcHubCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FtdcHub *FtdcHubSession) GetAddressUpdater() (common.Address, error) {
	return _FtdcHub.Contract.GetAddressUpdater(&_FtdcHub.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FtdcHub *FtdcHubCallerSession) GetAddressUpdater() (common.Address, error) {
	return _FtdcHub.Contract.GetAddressUpdater(&_FtdcHub.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FtdcHub *FtdcHubCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FtdcHub *FtdcHubSession) Governance() (common.Address, error) {
	return _FtdcHub.Contract.Governance(&_FtdcHub.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FtdcHub *FtdcHubCallerSession) Governance() (common.Address, error) {
	return _FtdcHub.Contract.Governance(&_FtdcHub.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FtdcHub *FtdcHubCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FtdcHub *FtdcHubSession) GovernanceSettings() (common.Address, error) {
	return _FtdcHub.Contract.GovernanceSettings(&_FtdcHub.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FtdcHub *FtdcHubCallerSession) GovernanceSettings() (common.Address, error) {
	return _FtdcHub.Contract.GovernanceSettings(&_FtdcHub.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FtdcHub *FtdcHubCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FtdcHub *FtdcHubSession) IsExecutor(_address common.Address) (bool, error) {
	return _FtdcHub.Contract.IsExecutor(&_FtdcHub.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FtdcHub *FtdcHubCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _FtdcHub.Contract.IsExecutor(&_FtdcHub.CallOpts, _address)
}

// MinThresholdBIPS is a free data retrieval call binding the contract method 0xeabf5bac.
//
// Solidity: function minThresholdBIPS() view returns(uint16)
func (_FtdcHub *FtdcHubCaller) MinThresholdBIPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "minThresholdBIPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MinThresholdBIPS is a free data retrieval call binding the contract method 0xeabf5bac.
//
// Solidity: function minThresholdBIPS() view returns(uint16)
func (_FtdcHub *FtdcHubSession) MinThresholdBIPS() (uint16, error) {
	return _FtdcHub.Contract.MinThresholdBIPS(&_FtdcHub.CallOpts)
}

// MinThresholdBIPS is a free data retrieval call binding the contract method 0xeabf5bac.
//
// Solidity: function minThresholdBIPS() view returns(uint16)
func (_FtdcHub *FtdcHubCallerSession) MinThresholdBIPS() (uint16, error) {
	return _FtdcHub.Contract.MinThresholdBIPS(&_FtdcHub.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FtdcHub *FtdcHubCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FtdcHub *FtdcHubSession) ProductionMode() (bool, error) {
	return _FtdcHub.Contract.ProductionMode(&_FtdcHub.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FtdcHub *FtdcHubCallerSession) ProductionMode() (bool, error) {
	return _FtdcHub.Contract.ProductionMode(&_FtdcHub.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FtdcHub *FtdcHubCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FtdcHub *FtdcHubSession) RewardManager() (common.Address, error) {
	return _FtdcHub.Contract.RewardManager(&_FtdcHub.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_FtdcHub *FtdcHubCallerSession) RewardManager() (common.Address, error) {
	return _FtdcHub.Contract.RewardManager(&_FtdcHub.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_FtdcHub *FtdcHubCaller) TeeExtensionRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "teeExtensionRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_FtdcHub *FtdcHubSession) TeeExtensionRegistry() (common.Address, error) {
	return _FtdcHub.Contract.TeeExtensionRegistry(&_FtdcHub.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_FtdcHub *FtdcHubCallerSession) TeeExtensionRegistry() (common.Address, error) {
	return _FtdcHub.Contract.TeeExtensionRegistry(&_FtdcHub.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_FtdcHub *FtdcHubCaller) TeeMachineRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "teeMachineRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_FtdcHub *FtdcHubSession) TeeMachineRegistry() (common.Address, error) {
	return _FtdcHub.Contract.TeeMachineRegistry(&_FtdcHub.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_FtdcHub *FtdcHubCallerSession) TeeMachineRegistry() (common.Address, error) {
	return _FtdcHub.Contract.TeeMachineRegistry(&_FtdcHub.CallOpts)
}

// TeeReplication is a free data retrieval call binding the contract method 0x0f10ad79.
//
// Solidity: function teeReplication() view returns(address)
func (_FtdcHub *FtdcHubCaller) TeeReplication(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "teeReplication")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeReplication is a free data retrieval call binding the contract method 0x0f10ad79.
//
// Solidity: function teeReplication() view returns(address)
func (_FtdcHub *FtdcHubSession) TeeReplication() (common.Address, error) {
	return _FtdcHub.Contract.TeeReplication(&_FtdcHub.CallOpts)
}

// TeeReplication is a free data retrieval call binding the contract method 0x0f10ad79.
//
// Solidity: function teeReplication() view returns(address)
func (_FtdcHub *FtdcHubCallerSession) TeeReplication() (common.Address, error) {
	return _FtdcHub.Contract.TeeReplication(&_FtdcHub.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtdcHub *FtdcHubCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _FtdcHub.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_FtdcHub *FtdcHubSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FtdcHub.Contract.TimelockedCalls(&_FtdcHub.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtdcHub *FtdcHubCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FtdcHub.Contract.TimelockedCalls(&_FtdcHub.CallOpts, selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FtdcHub *FtdcHubTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FtdcHub.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FtdcHub *FtdcHubSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtdcHub.Contract.CancelGovernanceCall(&_FtdcHub.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FtdcHub *FtdcHubTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtdcHub.Contract.CancelGovernanceCall(&_FtdcHub.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FtdcHub *FtdcHubTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FtdcHub.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FtdcHub *FtdcHubSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtdcHub.Contract.ExecuteGovernanceCall(&_FtdcHub.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FtdcHub *FtdcHubTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FtdcHub.Contract.ExecuteGovernanceCall(&_FtdcHub.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FtdcHub *FtdcHubTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FtdcHub.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FtdcHub *FtdcHubSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FtdcHub.Contract.Initialise(&_FtdcHub.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FtdcHub *FtdcHubTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FtdcHub.Contract.Initialise(&_FtdcHub.TransactOpts, _governanceSettings, _initialGovernance)
}

// RequestAttestation is a paid mutator transaction binding the contract method 0x8dec79a0.
//
// Solidity: function requestAttestation(uint16 _thresholdBIPS, uint256 _numberOfTees, address[] _teeIds, address[] _cosigners, uint64 _cosignersThreshold, bytes32 _attestationType, bytes32 _sourceId, bytes _requestBody) payable returns()
func (_FtdcHub *FtdcHubTransactor) RequestAttestation(opts *bind.TransactOpts, _thresholdBIPS uint16, _numberOfTees *big.Int, _teeIds []common.Address, _cosigners []common.Address, _cosignersThreshold uint64, _attestationType [32]byte, _sourceId [32]byte, _requestBody []byte) (*types.Transaction, error) {
	return _FtdcHub.contract.Transact(opts, "requestAttestation", _thresholdBIPS, _numberOfTees, _teeIds, _cosigners, _cosignersThreshold, _attestationType, _sourceId, _requestBody)
}

// RequestAttestation is a paid mutator transaction binding the contract method 0x8dec79a0.
//
// Solidity: function requestAttestation(uint16 _thresholdBIPS, uint256 _numberOfTees, address[] _teeIds, address[] _cosigners, uint64 _cosignersThreshold, bytes32 _attestationType, bytes32 _sourceId, bytes _requestBody) payable returns()
func (_FtdcHub *FtdcHubSession) RequestAttestation(_thresholdBIPS uint16, _numberOfTees *big.Int, _teeIds []common.Address, _cosigners []common.Address, _cosignersThreshold uint64, _attestationType [32]byte, _sourceId [32]byte, _requestBody []byte) (*types.Transaction, error) {
	return _FtdcHub.Contract.RequestAttestation(&_FtdcHub.TransactOpts, _thresholdBIPS, _numberOfTees, _teeIds, _cosigners, _cosignersThreshold, _attestationType, _sourceId, _requestBody)
}

// RequestAttestation is a paid mutator transaction binding the contract method 0x8dec79a0.
//
// Solidity: function requestAttestation(uint16 _thresholdBIPS, uint256 _numberOfTees, address[] _teeIds, address[] _cosigners, uint64 _cosignersThreshold, bytes32 _attestationType, bytes32 _sourceId, bytes _requestBody) payable returns()
func (_FtdcHub *FtdcHubTransactorSession) RequestAttestation(_thresholdBIPS uint16, _numberOfTees *big.Int, _teeIds []common.Address, _cosigners []common.Address, _cosignersThreshold uint64, _attestationType [32]byte, _sourceId [32]byte, _requestBody []byte) (*types.Transaction, error) {
	return _FtdcHub.Contract.RequestAttestation(&_FtdcHub.TransactOpts, _thresholdBIPS, _numberOfTees, _teeIds, _cosigners, _cosignersThreshold, _attestationType, _sourceId, _requestBody)
}

// SetDefaultNumberOfTees is a paid mutator transaction binding the contract method 0x8f46aa36.
//
// Solidity: function setDefaultNumberOfTees(uint8 _defaultNumberOfTees) returns()
func (_FtdcHub *FtdcHubTransactor) SetDefaultNumberOfTees(opts *bind.TransactOpts, _defaultNumberOfTees uint8) (*types.Transaction, error) {
	return _FtdcHub.contract.Transact(opts, "setDefaultNumberOfTees", _defaultNumberOfTees)
}

// SetDefaultNumberOfTees is a paid mutator transaction binding the contract method 0x8f46aa36.
//
// Solidity: function setDefaultNumberOfTees(uint8 _defaultNumberOfTees) returns()
func (_FtdcHub *FtdcHubSession) SetDefaultNumberOfTees(_defaultNumberOfTees uint8) (*types.Transaction, error) {
	return _FtdcHub.Contract.SetDefaultNumberOfTees(&_FtdcHub.TransactOpts, _defaultNumberOfTees)
}

// SetDefaultNumberOfTees is a paid mutator transaction binding the contract method 0x8f46aa36.
//
// Solidity: function setDefaultNumberOfTees(uint8 _defaultNumberOfTees) returns()
func (_FtdcHub *FtdcHubTransactorSession) SetDefaultNumberOfTees(_defaultNumberOfTees uint8) (*types.Transaction, error) {
	return _FtdcHub.Contract.SetDefaultNumberOfTees(&_FtdcHub.TransactOpts, _defaultNumberOfTees)
}

// SetMinThresholdBIPS is a paid mutator transaction binding the contract method 0x8b8e70f5.
//
// Solidity: function setMinThresholdBIPS(uint16 _minThresholdBIPS) returns()
func (_FtdcHub *FtdcHubTransactor) SetMinThresholdBIPS(opts *bind.TransactOpts, _minThresholdBIPS uint16) (*types.Transaction, error) {
	return _FtdcHub.contract.Transact(opts, "setMinThresholdBIPS", _minThresholdBIPS)
}

// SetMinThresholdBIPS is a paid mutator transaction binding the contract method 0x8b8e70f5.
//
// Solidity: function setMinThresholdBIPS(uint16 _minThresholdBIPS) returns()
func (_FtdcHub *FtdcHubSession) SetMinThresholdBIPS(_minThresholdBIPS uint16) (*types.Transaction, error) {
	return _FtdcHub.Contract.SetMinThresholdBIPS(&_FtdcHub.TransactOpts, _minThresholdBIPS)
}

// SetMinThresholdBIPS is a paid mutator transaction binding the contract method 0x8b8e70f5.
//
// Solidity: function setMinThresholdBIPS(uint16 _minThresholdBIPS) returns()
func (_FtdcHub *FtdcHubTransactorSession) SetMinThresholdBIPS(_minThresholdBIPS uint16) (*types.Transaction, error) {
	return _FtdcHub.Contract.SetMinThresholdBIPS(&_FtdcHub.TransactOpts, _minThresholdBIPS)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FtdcHub *FtdcHubTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FtdcHub.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FtdcHub *FtdcHubSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FtdcHub.Contract.SwitchToProductionMode(&_FtdcHub.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FtdcHub *FtdcHubTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FtdcHub.Contract.SwitchToProductionMode(&_FtdcHub.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FtdcHub *FtdcHubTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FtdcHub.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FtdcHub *FtdcHubSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FtdcHub.Contract.UpdateContractAddresses(&_FtdcHub.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FtdcHub *FtdcHubTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FtdcHub.Contract.UpdateContractAddresses(&_FtdcHub.TransactOpts, _contractNameHashes, _contractAddresses)
}

// FtdcHubAttestationRequestedIterator is returned from FilterAttestationRequested and is used to iterate over the raw logs and unpacked data for AttestationRequested events raised by the FtdcHub contract.
type FtdcHubAttestationRequestedIterator struct {
	Event *FtdcHubAttestationRequested // Event containing the contract specifics and raw log

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
func (it *FtdcHubAttestationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcHubAttestationRequested)
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
		it.Event = new(FtdcHubAttestationRequested)
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
func (it *FtdcHubAttestationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcHubAttestationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcHubAttestationRequested represents a AttestationRequested event raised by the FtdcHub contract.
type FtdcHubAttestationRequested struct {
	AttestationType [32]byte
	SourceId        [32]byte
	RequestBody     []byte
	Fee             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAttestationRequested is a free log retrieval operation binding the contract event 0xf1602f3bc1606d2384ad0bed8b431d963c3f9b14fbe6f0ebf31ed8d5d6b2d9c5.
//
// Solidity: event AttestationRequested(bytes32 attestationType, bytes32 sourceId, bytes requestBody, uint256 fee)
func (_FtdcHub *FtdcHubFilterer) FilterAttestationRequested(opts *bind.FilterOpts) (*FtdcHubAttestationRequestedIterator, error) {

	logs, sub, err := _FtdcHub.contract.FilterLogs(opts, "AttestationRequested")
	if err != nil {
		return nil, err
	}
	return &FtdcHubAttestationRequestedIterator{contract: _FtdcHub.contract, event: "AttestationRequested", logs: logs, sub: sub}, nil
}

// WatchAttestationRequested is a free log subscription operation binding the contract event 0xf1602f3bc1606d2384ad0bed8b431d963c3f9b14fbe6f0ebf31ed8d5d6b2d9c5.
//
// Solidity: event AttestationRequested(bytes32 attestationType, bytes32 sourceId, bytes requestBody, uint256 fee)
func (_FtdcHub *FtdcHubFilterer) WatchAttestationRequested(opts *bind.WatchOpts, sink chan<- *FtdcHubAttestationRequested) (event.Subscription, error) {

	logs, sub, err := _FtdcHub.contract.WatchLogs(opts, "AttestationRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcHubAttestationRequested)
				if err := _FtdcHub.contract.UnpackLog(event, "AttestationRequested", log); err != nil {
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

// ParseAttestationRequested is a log parse operation binding the contract event 0xf1602f3bc1606d2384ad0bed8b431d963c3f9b14fbe6f0ebf31ed8d5d6b2d9c5.
//
// Solidity: event AttestationRequested(bytes32 attestationType, bytes32 sourceId, bytes requestBody, uint256 fee)
func (_FtdcHub *FtdcHubFilterer) ParseAttestationRequested(log types.Log) (*FtdcHubAttestationRequested, error) {
	event := new(FtdcHubAttestationRequested)
	if err := _FtdcHub.contract.UnpackLog(event, "AttestationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcHubDefaultNumberOfTeesSetIterator is returned from FilterDefaultNumberOfTeesSet and is used to iterate over the raw logs and unpacked data for DefaultNumberOfTeesSet events raised by the FtdcHub contract.
type FtdcHubDefaultNumberOfTeesSetIterator struct {
	Event *FtdcHubDefaultNumberOfTeesSet // Event containing the contract specifics and raw log

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
func (it *FtdcHubDefaultNumberOfTeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcHubDefaultNumberOfTeesSet)
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
		it.Event = new(FtdcHubDefaultNumberOfTeesSet)
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
func (it *FtdcHubDefaultNumberOfTeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcHubDefaultNumberOfTeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcHubDefaultNumberOfTeesSet represents a DefaultNumberOfTeesSet event raised by the FtdcHub contract.
type FtdcHubDefaultNumberOfTeesSet struct {
	DefaultNumberOfTees uint8
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDefaultNumberOfTeesSet is a free log retrieval operation binding the contract event 0xfd28c45cd39f6ed5f3bcbd90a07ba9adc16d49ba1440a07595ee4d6bd8f8dafd.
//
// Solidity: event DefaultNumberOfTeesSet(uint8 defaultNumberOfTees)
func (_FtdcHub *FtdcHubFilterer) FilterDefaultNumberOfTeesSet(opts *bind.FilterOpts) (*FtdcHubDefaultNumberOfTeesSetIterator, error) {

	logs, sub, err := _FtdcHub.contract.FilterLogs(opts, "DefaultNumberOfTeesSet")
	if err != nil {
		return nil, err
	}
	return &FtdcHubDefaultNumberOfTeesSetIterator{contract: _FtdcHub.contract, event: "DefaultNumberOfTeesSet", logs: logs, sub: sub}, nil
}

// WatchDefaultNumberOfTeesSet is a free log subscription operation binding the contract event 0xfd28c45cd39f6ed5f3bcbd90a07ba9adc16d49ba1440a07595ee4d6bd8f8dafd.
//
// Solidity: event DefaultNumberOfTeesSet(uint8 defaultNumberOfTees)
func (_FtdcHub *FtdcHubFilterer) WatchDefaultNumberOfTeesSet(opts *bind.WatchOpts, sink chan<- *FtdcHubDefaultNumberOfTeesSet) (event.Subscription, error) {

	logs, sub, err := _FtdcHub.contract.WatchLogs(opts, "DefaultNumberOfTeesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcHubDefaultNumberOfTeesSet)
				if err := _FtdcHub.contract.UnpackLog(event, "DefaultNumberOfTeesSet", log); err != nil {
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

// ParseDefaultNumberOfTeesSet is a log parse operation binding the contract event 0xfd28c45cd39f6ed5f3bcbd90a07ba9adc16d49ba1440a07595ee4d6bd8f8dafd.
//
// Solidity: event DefaultNumberOfTeesSet(uint8 defaultNumberOfTees)
func (_FtdcHub *FtdcHubFilterer) ParseDefaultNumberOfTeesSet(log types.Log) (*FtdcHubDefaultNumberOfTeesSet, error) {
	event := new(FtdcHubDefaultNumberOfTeesSet)
	if err := _FtdcHub.contract.UnpackLog(event, "DefaultNumberOfTeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcHubGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the FtdcHub contract.
type FtdcHubGovernanceCallTimelockedIterator struct {
	Event *FtdcHubGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *FtdcHubGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcHubGovernanceCallTimelocked)
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
		it.Event = new(FtdcHubGovernanceCallTimelocked)
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
func (it *FtdcHubGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcHubGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcHubGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FtdcHub contract.
type FtdcHubGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtdcHub *FtdcHubFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*FtdcHubGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _FtdcHub.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &FtdcHubGovernanceCallTimelockedIterator{contract: _FtdcHub.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FtdcHub *FtdcHubFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *FtdcHubGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _FtdcHub.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcHubGovernanceCallTimelocked)
				if err := _FtdcHub.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_FtdcHub *FtdcHubFilterer) ParseGovernanceCallTimelocked(log types.Log) (*FtdcHubGovernanceCallTimelocked, error) {
	event := new(FtdcHubGovernanceCallTimelocked)
	if err := _FtdcHub.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcHubGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the FtdcHub contract.
type FtdcHubGovernanceInitialisedIterator struct {
	Event *FtdcHubGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *FtdcHubGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcHubGovernanceInitialised)
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
		it.Event = new(FtdcHubGovernanceInitialised)
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
func (it *FtdcHubGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcHubGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcHubGovernanceInitialised represents a GovernanceInitialised event raised by the FtdcHub contract.
type FtdcHubGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FtdcHub *FtdcHubFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*FtdcHubGovernanceInitialisedIterator, error) {

	logs, sub, err := _FtdcHub.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &FtdcHubGovernanceInitialisedIterator{contract: _FtdcHub.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FtdcHub *FtdcHubFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *FtdcHubGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _FtdcHub.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcHubGovernanceInitialised)
				if err := _FtdcHub.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_FtdcHub *FtdcHubFilterer) ParseGovernanceInitialised(log types.Log) (*FtdcHubGovernanceInitialised, error) {
	event := new(FtdcHubGovernanceInitialised)
	if err := _FtdcHub.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcHubGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the FtdcHub contract.
type FtdcHubGovernedProductionModeEnteredIterator struct {
	Event *FtdcHubGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *FtdcHubGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcHubGovernedProductionModeEntered)
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
		it.Event = new(FtdcHubGovernedProductionModeEntered)
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
func (it *FtdcHubGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcHubGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcHubGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FtdcHub contract.
type FtdcHubGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FtdcHub *FtdcHubFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*FtdcHubGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _FtdcHub.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &FtdcHubGovernedProductionModeEnteredIterator{contract: _FtdcHub.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FtdcHub *FtdcHubFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *FtdcHubGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _FtdcHub.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcHubGovernedProductionModeEntered)
				if err := _FtdcHub.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_FtdcHub *FtdcHubFilterer) ParseGovernedProductionModeEntered(log types.Log) (*FtdcHubGovernedProductionModeEntered, error) {
	event := new(FtdcHubGovernedProductionModeEntered)
	if err := _FtdcHub.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcHubMinThresholdBIPSSetIterator is returned from FilterMinThresholdBIPSSet and is used to iterate over the raw logs and unpacked data for MinThresholdBIPSSet events raised by the FtdcHub contract.
type FtdcHubMinThresholdBIPSSetIterator struct {
	Event *FtdcHubMinThresholdBIPSSet // Event containing the contract specifics and raw log

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
func (it *FtdcHubMinThresholdBIPSSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcHubMinThresholdBIPSSet)
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
		it.Event = new(FtdcHubMinThresholdBIPSSet)
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
func (it *FtdcHubMinThresholdBIPSSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcHubMinThresholdBIPSSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcHubMinThresholdBIPSSet represents a MinThresholdBIPSSet event raised by the FtdcHub contract.
type FtdcHubMinThresholdBIPSSet struct {
	MinThresholdBIPS uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMinThresholdBIPSSet is a free log retrieval operation binding the contract event 0xc9abb127b1b92ec942674d22e590b5a81e559d7a65f23c63965166a3ffacb92d.
//
// Solidity: event MinThresholdBIPSSet(uint16 minThresholdBIPS)
func (_FtdcHub *FtdcHubFilterer) FilterMinThresholdBIPSSet(opts *bind.FilterOpts) (*FtdcHubMinThresholdBIPSSetIterator, error) {

	logs, sub, err := _FtdcHub.contract.FilterLogs(opts, "MinThresholdBIPSSet")
	if err != nil {
		return nil, err
	}
	return &FtdcHubMinThresholdBIPSSetIterator{contract: _FtdcHub.contract, event: "MinThresholdBIPSSet", logs: logs, sub: sub}, nil
}

// WatchMinThresholdBIPSSet is a free log subscription operation binding the contract event 0xc9abb127b1b92ec942674d22e590b5a81e559d7a65f23c63965166a3ffacb92d.
//
// Solidity: event MinThresholdBIPSSet(uint16 minThresholdBIPS)
func (_FtdcHub *FtdcHubFilterer) WatchMinThresholdBIPSSet(opts *bind.WatchOpts, sink chan<- *FtdcHubMinThresholdBIPSSet) (event.Subscription, error) {

	logs, sub, err := _FtdcHub.contract.WatchLogs(opts, "MinThresholdBIPSSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcHubMinThresholdBIPSSet)
				if err := _FtdcHub.contract.UnpackLog(event, "MinThresholdBIPSSet", log); err != nil {
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

// ParseMinThresholdBIPSSet is a log parse operation binding the contract event 0xc9abb127b1b92ec942674d22e590b5a81e559d7a65f23c63965166a3ffacb92d.
//
// Solidity: event MinThresholdBIPSSet(uint16 minThresholdBIPS)
func (_FtdcHub *FtdcHubFilterer) ParseMinThresholdBIPSSet(log types.Log) (*FtdcHubMinThresholdBIPSSet, error) {
	event := new(FtdcHubMinThresholdBIPSSet)
	if err := _FtdcHub.contract.UnpackLog(event, "MinThresholdBIPSSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcHubTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the FtdcHub contract.
type FtdcHubTimelockedGovernanceCallCanceledIterator struct {
	Event *FtdcHubTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *FtdcHubTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcHubTimelockedGovernanceCallCanceled)
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
		it.Event = new(FtdcHubTimelockedGovernanceCallCanceled)
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
func (it *FtdcHubTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcHubTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcHubTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FtdcHub contract.
type FtdcHubTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FtdcHub *FtdcHubFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*FtdcHubTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _FtdcHub.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &FtdcHubTimelockedGovernanceCallCanceledIterator{contract: _FtdcHub.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FtdcHub *FtdcHubFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *FtdcHubTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _FtdcHub.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcHubTimelockedGovernanceCallCanceled)
				if err := _FtdcHub.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_FtdcHub *FtdcHubFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*FtdcHubTimelockedGovernanceCallCanceled, error) {
	event := new(FtdcHubTimelockedGovernanceCallCanceled)
	if err := _FtdcHub.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FtdcHubTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the FtdcHub contract.
type FtdcHubTimelockedGovernanceCallExecutedIterator struct {
	Event *FtdcHubTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *FtdcHubTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FtdcHubTimelockedGovernanceCallExecuted)
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
		it.Event = new(FtdcHubTimelockedGovernanceCallExecuted)
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
func (it *FtdcHubTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FtdcHubTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FtdcHubTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FtdcHub contract.
type FtdcHubTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FtdcHub *FtdcHubFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*FtdcHubTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _FtdcHub.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &FtdcHubTimelockedGovernanceCallExecutedIterator{contract: _FtdcHub.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FtdcHub *FtdcHubFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *FtdcHubTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _FtdcHub.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FtdcHubTimelockedGovernanceCallExecuted)
				if err := _FtdcHub.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_FtdcHub *FtdcHubFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*FtdcHubTimelockedGovernanceCallExecuted, error) {
	event := new(FtdcHubTimelockedGovernanceCallExecuted)
	if err := _FtdcHub.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
