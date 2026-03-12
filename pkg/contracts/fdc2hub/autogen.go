// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fdc2hub

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

// IFdc2HubFdc2AttestationRequest is an auto generated low-level Go binding around an user-defined struct.
type IFdc2HubFdc2AttestationRequest struct {
	Header      IFdc2HubFdc2RequestHeader
	RequestBody []byte
}

// IFdc2HubFdc2RequestHeader is an auto generated low-level Go binding around an user-defined struct.
type IFdc2HubFdc2RequestHeader struct {
	AttestationType [32]byte
	SourceId        [32]byte
	ThresholdBIPS   uint16
	ProofOwner      common.Address
}

// Fdc2HubMetaData contains all meta data concerning the Fdc2Hub contract.
var Fdc2HubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DefaultNumberOfTeesZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"name\":\"DuplicatedTeeId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinThresholdInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MultipleResponsesPossible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NumberOfTeesAndTeeIdsInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"AttestationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"defaultNumberOfTees\",\"type\":\"uint8\"}],\"name\":\"DefaultNumberOfTeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minThresholdBIPS\",\"type\":\"uint16\"}],\"name\":\"MinThresholdBIPSSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FDC2_OP_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultNumberOfTees\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2RequestFeeConfigurations\",\"outputs\":[{\"internalType\":\"contractIFdc2RequestFeeConfigurations\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_minThresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"_defaultNumberOfTees\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minThresholdBIPS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"}],\"internalType\":\"structIFdc2Hub.Fdc2RequestHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"requestBody\",\"type\":\"bytes\"}],\"internalType\":\"structIFdc2Hub.Fdc2AttestationRequest\",\"name\":\"_attestationRequest\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_numberOfTees\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_defaultNumberOfTees\",\"type\":\"uint8\"}],\"name\":\"setDefaultNumberOfTees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_minThresholdBIPS\",\"type\":\"uint16\"}],\"name\":\"setMinThresholdBIPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeExtensionRegistry\",\"outputs\":[{\"internalType\":\"contractIITeeExtensionRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeMachineRegistry\",\"outputs\":[{\"internalType\":\"contractITeeMachineRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeReplication\",\"outputs\":[{\"internalType\":\"contractITeeReplication\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// Fdc2HubABI is the input ABI used to generate the binding from.
// Deprecated: Use Fdc2HubMetaData.ABI instead.
var Fdc2HubABI = Fdc2HubMetaData.ABI

// Fdc2Hub is an auto generated Go binding around an Ethereum contract.
type Fdc2Hub struct {
	Fdc2HubCaller     // Read-only binding to the contract
	Fdc2HubTransactor // Write-only binding to the contract
	Fdc2HubFilterer   // Log filterer for contract events
}

// Fdc2HubCaller is an auto generated read-only Go binding around an Ethereum contract.
type Fdc2HubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fdc2HubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Fdc2HubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fdc2HubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Fdc2HubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fdc2HubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Fdc2HubSession struct {
	Contract     *Fdc2Hub          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Fdc2HubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Fdc2HubCallerSession struct {
	Contract *Fdc2HubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// Fdc2HubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Fdc2HubTransactorSession struct {
	Contract     *Fdc2HubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Fdc2HubRaw is an auto generated low-level Go binding around an Ethereum contract.
type Fdc2HubRaw struct {
	Contract *Fdc2Hub // Generic contract binding to access the raw methods on
}

// Fdc2HubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Fdc2HubCallerRaw struct {
	Contract *Fdc2HubCaller // Generic read-only contract binding to access the raw methods on
}

// Fdc2HubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Fdc2HubTransactorRaw struct {
	Contract *Fdc2HubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFdc2Hub creates a new instance of Fdc2Hub, bound to a specific deployed contract.
func NewFdc2Hub(address common.Address, backend bind.ContractBackend) (*Fdc2Hub, error) {
	contract, err := bindFdc2Hub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fdc2Hub{Fdc2HubCaller: Fdc2HubCaller{contract: contract}, Fdc2HubTransactor: Fdc2HubTransactor{contract: contract}, Fdc2HubFilterer: Fdc2HubFilterer{contract: contract}}, nil
}

// NewFdc2HubCaller creates a new read-only instance of Fdc2Hub, bound to a specific deployed contract.
func NewFdc2HubCaller(address common.Address, caller bind.ContractCaller) (*Fdc2HubCaller, error) {
	contract, err := bindFdc2Hub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Fdc2HubCaller{contract: contract}, nil
}

// NewFdc2HubTransactor creates a new write-only instance of Fdc2Hub, bound to a specific deployed contract.
func NewFdc2HubTransactor(address common.Address, transactor bind.ContractTransactor) (*Fdc2HubTransactor, error) {
	contract, err := bindFdc2Hub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Fdc2HubTransactor{contract: contract}, nil
}

// NewFdc2HubFilterer creates a new log filterer instance of Fdc2Hub, bound to a specific deployed contract.
func NewFdc2HubFilterer(address common.Address, filterer bind.ContractFilterer) (*Fdc2HubFilterer, error) {
	contract, err := bindFdc2Hub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Fdc2HubFilterer{contract: contract}, nil
}

// bindFdc2Hub binds a generic wrapper to an already deployed contract.
func bindFdc2Hub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Fdc2HubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fdc2Hub *Fdc2HubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fdc2Hub.Contract.Fdc2HubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fdc2Hub *Fdc2HubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.Fdc2HubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fdc2Hub *Fdc2HubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.Fdc2HubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fdc2Hub *Fdc2HubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fdc2Hub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fdc2Hub *Fdc2HubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fdc2Hub *Fdc2HubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.contract.Transact(opts, method, params...)
}

// FDC2OPTYPE is a free data retrieval call binding the contract method 0x6b845d3c.
//
// Solidity: function FDC2_OP_TYPE() view returns(bytes32)
func (_Fdc2Hub *Fdc2HubCaller) FDC2OPTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "FDC2_OP_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FDC2OPTYPE is a free data retrieval call binding the contract method 0x6b845d3c.
//
// Solidity: function FDC2_OP_TYPE() view returns(bytes32)
func (_Fdc2Hub *Fdc2HubSession) FDC2OPTYPE() ([32]byte, error) {
	return _Fdc2Hub.Contract.FDC2OPTYPE(&_Fdc2Hub.CallOpts)
}

// FDC2OPTYPE is a free data retrieval call binding the contract method 0x6b845d3c.
//
// Solidity: function FDC2_OP_TYPE() view returns(bytes32)
func (_Fdc2Hub *Fdc2HubCallerSession) FDC2OPTYPE() ([32]byte, error) {
	return _Fdc2Hub.Contract.FDC2OPTYPE(&_Fdc2Hub.CallOpts)
}

// PROVE is a free data retrieval call binding the contract method 0x14f20ca6.
//
// Solidity: function PROVE() view returns(bytes32)
func (_Fdc2Hub *Fdc2HubCaller) PROVE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "PROVE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROVE is a free data retrieval call binding the contract method 0x14f20ca6.
//
// Solidity: function PROVE() view returns(bytes32)
func (_Fdc2Hub *Fdc2HubSession) PROVE() ([32]byte, error) {
	return _Fdc2Hub.Contract.PROVE(&_Fdc2Hub.CallOpts)
}

// PROVE is a free data retrieval call binding the contract method 0x14f20ca6.
//
// Solidity: function PROVE() view returns(bytes32)
func (_Fdc2Hub *Fdc2HubCallerSession) PROVE() ([32]byte, error) {
	return _Fdc2Hub.Contract.PROVE(&_Fdc2Hub.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Fdc2Hub *Fdc2HubCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Fdc2Hub *Fdc2HubSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Fdc2Hub.Contract.UPGRADEINTERFACEVERSION(&_Fdc2Hub.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Fdc2Hub *Fdc2HubCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Fdc2Hub.Contract.UPGRADEINTERFACEVERSION(&_Fdc2Hub.CallOpts)
}

// DefaultNumberOfTees is a free data retrieval call binding the contract method 0x1b92dbb4.
//
// Solidity: function defaultNumberOfTees() view returns(uint8)
func (_Fdc2Hub *Fdc2HubCaller) DefaultNumberOfTees(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "defaultNumberOfTees")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultNumberOfTees is a free data retrieval call binding the contract method 0x1b92dbb4.
//
// Solidity: function defaultNumberOfTees() view returns(uint8)
func (_Fdc2Hub *Fdc2HubSession) DefaultNumberOfTees() (uint8, error) {
	return _Fdc2Hub.Contract.DefaultNumberOfTees(&_Fdc2Hub.CallOpts)
}

// DefaultNumberOfTees is a free data retrieval call binding the contract method 0x1b92dbb4.
//
// Solidity: function defaultNumberOfTees() view returns(uint8)
func (_Fdc2Hub *Fdc2HubCallerSession) DefaultNumberOfTees() (uint8, error) {
	return _Fdc2Hub.Contract.DefaultNumberOfTees(&_Fdc2Hub.CallOpts)
}

// Fdc2RequestFeeConfigurations is a free data retrieval call binding the contract method 0x285434a2.
//
// Solidity: function fdc2RequestFeeConfigurations() view returns(address)
func (_Fdc2Hub *Fdc2HubCaller) Fdc2RequestFeeConfigurations(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "fdc2RequestFeeConfigurations")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fdc2RequestFeeConfigurations is a free data retrieval call binding the contract method 0x285434a2.
//
// Solidity: function fdc2RequestFeeConfigurations() view returns(address)
func (_Fdc2Hub *Fdc2HubSession) Fdc2RequestFeeConfigurations() (common.Address, error) {
	return _Fdc2Hub.Contract.Fdc2RequestFeeConfigurations(&_Fdc2Hub.CallOpts)
}

// Fdc2RequestFeeConfigurations is a free data retrieval call binding the contract method 0x285434a2.
//
// Solidity: function fdc2RequestFeeConfigurations() view returns(address)
func (_Fdc2Hub *Fdc2HubCallerSession) Fdc2RequestFeeConfigurations() (common.Address, error) {
	return _Fdc2Hub.Contract.Fdc2RequestFeeConfigurations(&_Fdc2Hub.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_Fdc2Hub *Fdc2HubCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_Fdc2Hub *Fdc2HubSession) FlareSystemsManager() (common.Address, error) {
	return _Fdc2Hub.Contract.FlareSystemsManager(&_Fdc2Hub.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_Fdc2Hub *Fdc2HubCallerSession) FlareSystemsManager() (common.Address, error) {
	return _Fdc2Hub.Contract.FlareSystemsManager(&_Fdc2Hub.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Fdc2Hub *Fdc2HubCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Fdc2Hub *Fdc2HubSession) GetAddressUpdater() (common.Address, error) {
	return _Fdc2Hub.Contract.GetAddressUpdater(&_Fdc2Hub.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_Fdc2Hub *Fdc2HubCallerSession) GetAddressUpdater() (common.Address, error) {
	return _Fdc2Hub.Contract.GetAddressUpdater(&_Fdc2Hub.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Fdc2Hub *Fdc2HubCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Fdc2Hub *Fdc2HubSession) Governance() (common.Address, error) {
	return _Fdc2Hub.Contract.Governance(&_Fdc2Hub.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Fdc2Hub *Fdc2HubCallerSession) Governance() (common.Address, error) {
	return _Fdc2Hub.Contract.Governance(&_Fdc2Hub.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Fdc2Hub *Fdc2HubCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Fdc2Hub *Fdc2HubSession) GovernanceSettings() (common.Address, error) {
	return _Fdc2Hub.Contract.GovernanceSettings(&_Fdc2Hub.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_Fdc2Hub *Fdc2HubCallerSession) GovernanceSettings() (common.Address, error) {
	return _Fdc2Hub.Contract.GovernanceSettings(&_Fdc2Hub.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Fdc2Hub *Fdc2HubCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Fdc2Hub *Fdc2HubSession) Implementation() (common.Address, error) {
	return _Fdc2Hub.Contract.Implementation(&_Fdc2Hub.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Fdc2Hub *Fdc2HubCallerSession) Implementation() (common.Address, error) {
	return _Fdc2Hub.Contract.Implementation(&_Fdc2Hub.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_Fdc2Hub *Fdc2HubCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_Fdc2Hub *Fdc2HubSession) IsExecutor(_address common.Address) (bool, error) {
	return _Fdc2Hub.Contract.IsExecutor(&_Fdc2Hub.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_Fdc2Hub *Fdc2HubCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _Fdc2Hub.Contract.IsExecutor(&_Fdc2Hub.CallOpts, _address)
}

// MinThresholdBIPS is a free data retrieval call binding the contract method 0xeabf5bac.
//
// Solidity: function minThresholdBIPS() view returns(uint16)
func (_Fdc2Hub *Fdc2HubCaller) MinThresholdBIPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "minThresholdBIPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MinThresholdBIPS is a free data retrieval call binding the contract method 0xeabf5bac.
//
// Solidity: function minThresholdBIPS() view returns(uint16)
func (_Fdc2Hub *Fdc2HubSession) MinThresholdBIPS() (uint16, error) {
	return _Fdc2Hub.Contract.MinThresholdBIPS(&_Fdc2Hub.CallOpts)
}

// MinThresholdBIPS is a free data retrieval call binding the contract method 0xeabf5bac.
//
// Solidity: function minThresholdBIPS() view returns(uint16)
func (_Fdc2Hub *Fdc2HubCallerSession) MinThresholdBIPS() (uint16, error) {
	return _Fdc2Hub.Contract.MinThresholdBIPS(&_Fdc2Hub.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Fdc2Hub *Fdc2HubCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Fdc2Hub *Fdc2HubSession) ProductionMode() (bool, error) {
	return _Fdc2Hub.Contract.ProductionMode(&_Fdc2Hub.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_Fdc2Hub *Fdc2HubCallerSession) ProductionMode() (bool, error) {
	return _Fdc2Hub.Contract.ProductionMode(&_Fdc2Hub.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Fdc2Hub *Fdc2HubCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Fdc2Hub *Fdc2HubSession) ProxiableUUID() ([32]byte, error) {
	return _Fdc2Hub.Contract.ProxiableUUID(&_Fdc2Hub.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Fdc2Hub *Fdc2HubCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Fdc2Hub.Contract.ProxiableUUID(&_Fdc2Hub.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_Fdc2Hub *Fdc2HubCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_Fdc2Hub *Fdc2HubSession) RewardManager() (common.Address, error) {
	return _Fdc2Hub.Contract.RewardManager(&_Fdc2Hub.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_Fdc2Hub *Fdc2HubCallerSession) RewardManager() (common.Address, error) {
	return _Fdc2Hub.Contract.RewardManager(&_Fdc2Hub.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_Fdc2Hub *Fdc2HubCaller) TeeExtensionRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "teeExtensionRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_Fdc2Hub *Fdc2HubSession) TeeExtensionRegistry() (common.Address, error) {
	return _Fdc2Hub.Contract.TeeExtensionRegistry(&_Fdc2Hub.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_Fdc2Hub *Fdc2HubCallerSession) TeeExtensionRegistry() (common.Address, error) {
	return _Fdc2Hub.Contract.TeeExtensionRegistry(&_Fdc2Hub.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_Fdc2Hub *Fdc2HubCaller) TeeMachineRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "teeMachineRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_Fdc2Hub *Fdc2HubSession) TeeMachineRegistry() (common.Address, error) {
	return _Fdc2Hub.Contract.TeeMachineRegistry(&_Fdc2Hub.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_Fdc2Hub *Fdc2HubCallerSession) TeeMachineRegistry() (common.Address, error) {
	return _Fdc2Hub.Contract.TeeMachineRegistry(&_Fdc2Hub.CallOpts)
}

// TeeReplication is a free data retrieval call binding the contract method 0x0f10ad79.
//
// Solidity: function teeReplication() view returns(address)
func (_Fdc2Hub *Fdc2HubCaller) TeeReplication(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "teeReplication")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeReplication is a free data retrieval call binding the contract method 0x0f10ad79.
//
// Solidity: function teeReplication() view returns(address)
func (_Fdc2Hub *Fdc2HubSession) TeeReplication() (common.Address, error) {
	return _Fdc2Hub.Contract.TeeReplication(&_Fdc2Hub.CallOpts)
}

// TeeReplication is a free data retrieval call binding the contract method 0x0f10ad79.
//
// Solidity: function teeReplication() view returns(address)
func (_Fdc2Hub *Fdc2HubCallerSession) TeeReplication() (common.Address, error) {
	return _Fdc2Hub.Contract.TeeReplication(&_Fdc2Hub.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Fdc2Hub *Fdc2HubCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _Fdc2Hub.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_Fdc2Hub *Fdc2HubSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _Fdc2Hub.Contract.TimelockedCalls(&_Fdc2Hub.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Fdc2Hub *Fdc2HubCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _Fdc2Hub.Contract.TimelockedCalls(&_Fdc2Hub.CallOpts, selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Fdc2Hub *Fdc2HubTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _Fdc2Hub.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Fdc2Hub *Fdc2HubSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.CancelGovernanceCall(&_Fdc2Hub.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_Fdc2Hub *Fdc2HubTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.CancelGovernanceCall(&_Fdc2Hub.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Fdc2Hub *Fdc2HubTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _Fdc2Hub.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Fdc2Hub *Fdc2HubSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.ExecuteGovernanceCall(&_Fdc2Hub.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_Fdc2Hub *Fdc2HubTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.ExecuteGovernanceCall(&_Fdc2Hub.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_Fdc2Hub *Fdc2HubTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Fdc2Hub.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_Fdc2Hub *Fdc2HubSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.Initialise(&_Fdc2Hub.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_Fdc2Hub *Fdc2HubTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.Initialise(&_Fdc2Hub.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xd6b35a06.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint16 _minThresholdBIPS, uint8 _defaultNumberOfTees) returns()
func (_Fdc2Hub *Fdc2HubTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _minThresholdBIPS uint16, _defaultNumberOfTees uint8) (*types.Transaction, error) {
	return _Fdc2Hub.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater, _minThresholdBIPS, _defaultNumberOfTees)
}

// Initialize is a paid mutator transaction binding the contract method 0xd6b35a06.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint16 _minThresholdBIPS, uint8 _defaultNumberOfTees) returns()
func (_Fdc2Hub *Fdc2HubSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _minThresholdBIPS uint16, _defaultNumberOfTees uint8) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.Initialize(&_Fdc2Hub.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _minThresholdBIPS, _defaultNumberOfTees)
}

// Initialize is a paid mutator transaction binding the contract method 0xd6b35a06.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint16 _minThresholdBIPS, uint8 _defaultNumberOfTees) returns()
func (_Fdc2Hub *Fdc2HubTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _minThresholdBIPS uint16, _defaultNumberOfTees uint8) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.Initialize(&_Fdc2Hub.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _minThresholdBIPS, _defaultNumberOfTees)
}

// RequestAttestation is a paid mutator transaction binding the contract method 0x1398dddb.
//
// Solidity: function requestAttestation(((bytes32,bytes32,uint16,address),bytes) _attestationRequest, uint256 _numberOfTees, address[] _teeIds, address[] _cosigners, uint64 _cosignersThreshold, address _claimBackAddress) payable returns()
func (_Fdc2Hub *Fdc2HubTransactor) RequestAttestation(opts *bind.TransactOpts, _attestationRequest IFdc2HubFdc2AttestationRequest, _numberOfTees *big.Int, _teeIds []common.Address, _cosigners []common.Address, _cosignersThreshold uint64, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _Fdc2Hub.contract.Transact(opts, "requestAttestation", _attestationRequest, _numberOfTees, _teeIds, _cosigners, _cosignersThreshold, _claimBackAddress)
}

// RequestAttestation is a paid mutator transaction binding the contract method 0x1398dddb.
//
// Solidity: function requestAttestation(((bytes32,bytes32,uint16,address),bytes) _attestationRequest, uint256 _numberOfTees, address[] _teeIds, address[] _cosigners, uint64 _cosignersThreshold, address _claimBackAddress) payable returns()
func (_Fdc2Hub *Fdc2HubSession) RequestAttestation(_attestationRequest IFdc2HubFdc2AttestationRequest, _numberOfTees *big.Int, _teeIds []common.Address, _cosigners []common.Address, _cosignersThreshold uint64, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.RequestAttestation(&_Fdc2Hub.TransactOpts, _attestationRequest, _numberOfTees, _teeIds, _cosigners, _cosignersThreshold, _claimBackAddress)
}

// RequestAttestation is a paid mutator transaction binding the contract method 0x1398dddb.
//
// Solidity: function requestAttestation(((bytes32,bytes32,uint16,address),bytes) _attestationRequest, uint256 _numberOfTees, address[] _teeIds, address[] _cosigners, uint64 _cosignersThreshold, address _claimBackAddress) payable returns()
func (_Fdc2Hub *Fdc2HubTransactorSession) RequestAttestation(_attestationRequest IFdc2HubFdc2AttestationRequest, _numberOfTees *big.Int, _teeIds []common.Address, _cosigners []common.Address, _cosignersThreshold uint64, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.RequestAttestation(&_Fdc2Hub.TransactOpts, _attestationRequest, _numberOfTees, _teeIds, _cosigners, _cosignersThreshold, _claimBackAddress)
}

// SetDefaultNumberOfTees is a paid mutator transaction binding the contract method 0x8f46aa36.
//
// Solidity: function setDefaultNumberOfTees(uint8 _defaultNumberOfTees) returns()
func (_Fdc2Hub *Fdc2HubTransactor) SetDefaultNumberOfTees(opts *bind.TransactOpts, _defaultNumberOfTees uint8) (*types.Transaction, error) {
	return _Fdc2Hub.contract.Transact(opts, "setDefaultNumberOfTees", _defaultNumberOfTees)
}

// SetDefaultNumberOfTees is a paid mutator transaction binding the contract method 0x8f46aa36.
//
// Solidity: function setDefaultNumberOfTees(uint8 _defaultNumberOfTees) returns()
func (_Fdc2Hub *Fdc2HubSession) SetDefaultNumberOfTees(_defaultNumberOfTees uint8) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.SetDefaultNumberOfTees(&_Fdc2Hub.TransactOpts, _defaultNumberOfTees)
}

// SetDefaultNumberOfTees is a paid mutator transaction binding the contract method 0x8f46aa36.
//
// Solidity: function setDefaultNumberOfTees(uint8 _defaultNumberOfTees) returns()
func (_Fdc2Hub *Fdc2HubTransactorSession) SetDefaultNumberOfTees(_defaultNumberOfTees uint8) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.SetDefaultNumberOfTees(&_Fdc2Hub.TransactOpts, _defaultNumberOfTees)
}

// SetMinThresholdBIPS is a paid mutator transaction binding the contract method 0x8b8e70f5.
//
// Solidity: function setMinThresholdBIPS(uint16 _minThresholdBIPS) returns()
func (_Fdc2Hub *Fdc2HubTransactor) SetMinThresholdBIPS(opts *bind.TransactOpts, _minThresholdBIPS uint16) (*types.Transaction, error) {
	return _Fdc2Hub.contract.Transact(opts, "setMinThresholdBIPS", _minThresholdBIPS)
}

// SetMinThresholdBIPS is a paid mutator transaction binding the contract method 0x8b8e70f5.
//
// Solidity: function setMinThresholdBIPS(uint16 _minThresholdBIPS) returns()
func (_Fdc2Hub *Fdc2HubSession) SetMinThresholdBIPS(_minThresholdBIPS uint16) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.SetMinThresholdBIPS(&_Fdc2Hub.TransactOpts, _minThresholdBIPS)
}

// SetMinThresholdBIPS is a paid mutator transaction binding the contract method 0x8b8e70f5.
//
// Solidity: function setMinThresholdBIPS(uint16 _minThresholdBIPS) returns()
func (_Fdc2Hub *Fdc2HubTransactorSession) SetMinThresholdBIPS(_minThresholdBIPS uint16) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.SetMinThresholdBIPS(&_Fdc2Hub.TransactOpts, _minThresholdBIPS)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Fdc2Hub *Fdc2HubTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fdc2Hub.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Fdc2Hub *Fdc2HubSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _Fdc2Hub.Contract.SwitchToProductionMode(&_Fdc2Hub.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_Fdc2Hub *Fdc2HubTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _Fdc2Hub.Contract.SwitchToProductionMode(&_Fdc2Hub.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Fdc2Hub *Fdc2HubTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Fdc2Hub.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Fdc2Hub *Fdc2HubSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.UpdateContractAddresses(&_Fdc2Hub.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_Fdc2Hub *Fdc2HubTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.UpdateContractAddresses(&_Fdc2Hub.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_Fdc2Hub *Fdc2HubTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _Fdc2Hub.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_Fdc2Hub *Fdc2HubSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.UpgradeToAndCall(&_Fdc2Hub.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_Fdc2Hub *Fdc2HubTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _Fdc2Hub.Contract.UpgradeToAndCall(&_Fdc2Hub.TransactOpts, _newImplementation, _data)
}

// Fdc2HubAttestationRequestedIterator is returned from FilterAttestationRequested and is used to iterate over the raw logs and unpacked data for AttestationRequested events raised by the Fdc2Hub contract.
type Fdc2HubAttestationRequestedIterator struct {
	Event *Fdc2HubAttestationRequested // Event containing the contract specifics and raw log

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
func (it *Fdc2HubAttestationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2HubAttestationRequested)
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
		it.Event = new(Fdc2HubAttestationRequested)
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
func (it *Fdc2HubAttestationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2HubAttestationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2HubAttestationRequested represents a AttestationRequested event raised by the Fdc2Hub contract.
type Fdc2HubAttestationRequested struct {
	InstructionId    [32]byte
	AttestationType  [32]byte
	SourceId         [32]byte
	ProofOwner       common.Address
	ClaimBackAddress common.Address
	Fee              *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAttestationRequested is a free log retrieval operation binding the contract event 0x57c4413905bb1b444f93a5eab5a942fae34c0fcaa1c25cc595ce0b990310f5de.
//
// Solidity: event AttestationRequested(bytes32 indexed instructionId, bytes32 indexed attestationType, bytes32 indexed sourceId, address proofOwner, address claimBackAddress, uint256 fee)
func (_Fdc2Hub *Fdc2HubFilterer) FilterAttestationRequested(opts *bind.FilterOpts, instructionId [][32]byte, attestationType [][32]byte, sourceId [][32]byte) (*Fdc2HubAttestationRequestedIterator, error) {

	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var attestationTypeRule []interface{}
	for _, attestationTypeItem := range attestationType {
		attestationTypeRule = append(attestationTypeRule, attestationTypeItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _Fdc2Hub.contract.FilterLogs(opts, "AttestationRequested", instructionIdRule, attestationTypeRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return &Fdc2HubAttestationRequestedIterator{contract: _Fdc2Hub.contract, event: "AttestationRequested", logs: logs, sub: sub}, nil
}

// WatchAttestationRequested is a free log subscription operation binding the contract event 0x57c4413905bb1b444f93a5eab5a942fae34c0fcaa1c25cc595ce0b990310f5de.
//
// Solidity: event AttestationRequested(bytes32 indexed instructionId, bytes32 indexed attestationType, bytes32 indexed sourceId, address proofOwner, address claimBackAddress, uint256 fee)
func (_Fdc2Hub *Fdc2HubFilterer) WatchAttestationRequested(opts *bind.WatchOpts, sink chan<- *Fdc2HubAttestationRequested, instructionId [][32]byte, attestationType [][32]byte, sourceId [][32]byte) (event.Subscription, error) {

	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var attestationTypeRule []interface{}
	for _, attestationTypeItem := range attestationType {
		attestationTypeRule = append(attestationTypeRule, attestationTypeItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _Fdc2Hub.contract.WatchLogs(opts, "AttestationRequested", instructionIdRule, attestationTypeRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2HubAttestationRequested)
				if err := _Fdc2Hub.contract.UnpackLog(event, "AttestationRequested", log); err != nil {
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

// ParseAttestationRequested is a log parse operation binding the contract event 0x57c4413905bb1b444f93a5eab5a942fae34c0fcaa1c25cc595ce0b990310f5de.
//
// Solidity: event AttestationRequested(bytes32 indexed instructionId, bytes32 indexed attestationType, bytes32 indexed sourceId, address proofOwner, address claimBackAddress, uint256 fee)
func (_Fdc2Hub *Fdc2HubFilterer) ParseAttestationRequested(log types.Log) (*Fdc2HubAttestationRequested, error) {
	event := new(Fdc2HubAttestationRequested)
	if err := _Fdc2Hub.contract.UnpackLog(event, "AttestationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2HubDefaultNumberOfTeesSetIterator is returned from FilterDefaultNumberOfTeesSet and is used to iterate over the raw logs and unpacked data for DefaultNumberOfTeesSet events raised by the Fdc2Hub contract.
type Fdc2HubDefaultNumberOfTeesSetIterator struct {
	Event *Fdc2HubDefaultNumberOfTeesSet // Event containing the contract specifics and raw log

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
func (it *Fdc2HubDefaultNumberOfTeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2HubDefaultNumberOfTeesSet)
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
		it.Event = new(Fdc2HubDefaultNumberOfTeesSet)
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
func (it *Fdc2HubDefaultNumberOfTeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2HubDefaultNumberOfTeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2HubDefaultNumberOfTeesSet represents a DefaultNumberOfTeesSet event raised by the Fdc2Hub contract.
type Fdc2HubDefaultNumberOfTeesSet struct {
	DefaultNumberOfTees uint8
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDefaultNumberOfTeesSet is a free log retrieval operation binding the contract event 0xfd28c45cd39f6ed5f3bcbd90a07ba9adc16d49ba1440a07595ee4d6bd8f8dafd.
//
// Solidity: event DefaultNumberOfTeesSet(uint8 defaultNumberOfTees)
func (_Fdc2Hub *Fdc2HubFilterer) FilterDefaultNumberOfTeesSet(opts *bind.FilterOpts) (*Fdc2HubDefaultNumberOfTeesSetIterator, error) {

	logs, sub, err := _Fdc2Hub.contract.FilterLogs(opts, "DefaultNumberOfTeesSet")
	if err != nil {
		return nil, err
	}
	return &Fdc2HubDefaultNumberOfTeesSetIterator{contract: _Fdc2Hub.contract, event: "DefaultNumberOfTeesSet", logs: logs, sub: sub}, nil
}

// WatchDefaultNumberOfTeesSet is a free log subscription operation binding the contract event 0xfd28c45cd39f6ed5f3bcbd90a07ba9adc16d49ba1440a07595ee4d6bd8f8dafd.
//
// Solidity: event DefaultNumberOfTeesSet(uint8 defaultNumberOfTees)
func (_Fdc2Hub *Fdc2HubFilterer) WatchDefaultNumberOfTeesSet(opts *bind.WatchOpts, sink chan<- *Fdc2HubDefaultNumberOfTeesSet) (event.Subscription, error) {

	logs, sub, err := _Fdc2Hub.contract.WatchLogs(opts, "DefaultNumberOfTeesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2HubDefaultNumberOfTeesSet)
				if err := _Fdc2Hub.contract.UnpackLog(event, "DefaultNumberOfTeesSet", log); err != nil {
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
func (_Fdc2Hub *Fdc2HubFilterer) ParseDefaultNumberOfTeesSet(log types.Log) (*Fdc2HubDefaultNumberOfTeesSet, error) {
	event := new(Fdc2HubDefaultNumberOfTeesSet)
	if err := _Fdc2Hub.contract.UnpackLog(event, "DefaultNumberOfTeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2HubGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the Fdc2Hub contract.
type Fdc2HubGovernanceCallTimelockedIterator struct {
	Event *Fdc2HubGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *Fdc2HubGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2HubGovernanceCallTimelocked)
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
		it.Event = new(Fdc2HubGovernanceCallTimelocked)
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
func (it *Fdc2HubGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2HubGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2HubGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Fdc2Hub contract.
type Fdc2HubGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Fdc2Hub *Fdc2HubFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*Fdc2HubGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _Fdc2Hub.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &Fdc2HubGovernanceCallTimelockedIterator{contract: _Fdc2Hub.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_Fdc2Hub *Fdc2HubFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *Fdc2HubGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _Fdc2Hub.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2HubGovernanceCallTimelocked)
				if err := _Fdc2Hub.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_Fdc2Hub *Fdc2HubFilterer) ParseGovernanceCallTimelocked(log types.Log) (*Fdc2HubGovernanceCallTimelocked, error) {
	event := new(Fdc2HubGovernanceCallTimelocked)
	if err := _Fdc2Hub.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2HubGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the Fdc2Hub contract.
type Fdc2HubGovernanceInitialisedIterator struct {
	Event *Fdc2HubGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *Fdc2HubGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2HubGovernanceInitialised)
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
		it.Event = new(Fdc2HubGovernanceInitialised)
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
func (it *Fdc2HubGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2HubGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2HubGovernanceInitialised represents a GovernanceInitialised event raised by the Fdc2Hub contract.
type Fdc2HubGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_Fdc2Hub *Fdc2HubFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*Fdc2HubGovernanceInitialisedIterator, error) {

	logs, sub, err := _Fdc2Hub.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &Fdc2HubGovernanceInitialisedIterator{contract: _Fdc2Hub.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_Fdc2Hub *Fdc2HubFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *Fdc2HubGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _Fdc2Hub.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2HubGovernanceInitialised)
				if err := _Fdc2Hub.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_Fdc2Hub *Fdc2HubFilterer) ParseGovernanceInitialised(log types.Log) (*Fdc2HubGovernanceInitialised, error) {
	event := new(Fdc2HubGovernanceInitialised)
	if err := _Fdc2Hub.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2HubGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the Fdc2Hub contract.
type Fdc2HubGovernedProductionModeEnteredIterator struct {
	Event *Fdc2HubGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *Fdc2HubGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2HubGovernedProductionModeEntered)
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
		it.Event = new(Fdc2HubGovernedProductionModeEntered)
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
func (it *Fdc2HubGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2HubGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2HubGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the Fdc2Hub contract.
type Fdc2HubGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_Fdc2Hub *Fdc2HubFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*Fdc2HubGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _Fdc2Hub.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &Fdc2HubGovernedProductionModeEnteredIterator{contract: _Fdc2Hub.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_Fdc2Hub *Fdc2HubFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *Fdc2HubGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _Fdc2Hub.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2HubGovernedProductionModeEntered)
				if err := _Fdc2Hub.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_Fdc2Hub *Fdc2HubFilterer) ParseGovernedProductionModeEntered(log types.Log) (*Fdc2HubGovernedProductionModeEntered, error) {
	event := new(Fdc2HubGovernedProductionModeEntered)
	if err := _Fdc2Hub.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2HubMinThresholdBIPSSetIterator is returned from FilterMinThresholdBIPSSet and is used to iterate over the raw logs and unpacked data for MinThresholdBIPSSet events raised by the Fdc2Hub contract.
type Fdc2HubMinThresholdBIPSSetIterator struct {
	Event *Fdc2HubMinThresholdBIPSSet // Event containing the contract specifics and raw log

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
func (it *Fdc2HubMinThresholdBIPSSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2HubMinThresholdBIPSSet)
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
		it.Event = new(Fdc2HubMinThresholdBIPSSet)
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
func (it *Fdc2HubMinThresholdBIPSSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2HubMinThresholdBIPSSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2HubMinThresholdBIPSSet represents a MinThresholdBIPSSet event raised by the Fdc2Hub contract.
type Fdc2HubMinThresholdBIPSSet struct {
	MinThresholdBIPS uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMinThresholdBIPSSet is a free log retrieval operation binding the contract event 0xc9abb127b1b92ec942674d22e590b5a81e559d7a65f23c63965166a3ffacb92d.
//
// Solidity: event MinThresholdBIPSSet(uint16 minThresholdBIPS)
func (_Fdc2Hub *Fdc2HubFilterer) FilterMinThresholdBIPSSet(opts *bind.FilterOpts) (*Fdc2HubMinThresholdBIPSSetIterator, error) {

	logs, sub, err := _Fdc2Hub.contract.FilterLogs(opts, "MinThresholdBIPSSet")
	if err != nil {
		return nil, err
	}
	return &Fdc2HubMinThresholdBIPSSetIterator{contract: _Fdc2Hub.contract, event: "MinThresholdBIPSSet", logs: logs, sub: sub}, nil
}

// WatchMinThresholdBIPSSet is a free log subscription operation binding the contract event 0xc9abb127b1b92ec942674d22e590b5a81e559d7a65f23c63965166a3ffacb92d.
//
// Solidity: event MinThresholdBIPSSet(uint16 minThresholdBIPS)
func (_Fdc2Hub *Fdc2HubFilterer) WatchMinThresholdBIPSSet(opts *bind.WatchOpts, sink chan<- *Fdc2HubMinThresholdBIPSSet) (event.Subscription, error) {

	logs, sub, err := _Fdc2Hub.contract.WatchLogs(opts, "MinThresholdBIPSSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2HubMinThresholdBIPSSet)
				if err := _Fdc2Hub.contract.UnpackLog(event, "MinThresholdBIPSSet", log); err != nil {
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
func (_Fdc2Hub *Fdc2HubFilterer) ParseMinThresholdBIPSSet(log types.Log) (*Fdc2HubMinThresholdBIPSSet, error) {
	event := new(Fdc2HubMinThresholdBIPSSet)
	if err := _Fdc2Hub.contract.UnpackLog(event, "MinThresholdBIPSSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2HubTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the Fdc2Hub contract.
type Fdc2HubTimelockedGovernanceCallCanceledIterator struct {
	Event *Fdc2HubTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *Fdc2HubTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2HubTimelockedGovernanceCallCanceled)
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
		it.Event = new(Fdc2HubTimelockedGovernanceCallCanceled)
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
func (it *Fdc2HubTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2HubTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2HubTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the Fdc2Hub contract.
type Fdc2HubTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_Fdc2Hub *Fdc2HubFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*Fdc2HubTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _Fdc2Hub.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &Fdc2HubTimelockedGovernanceCallCanceledIterator{contract: _Fdc2Hub.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_Fdc2Hub *Fdc2HubFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *Fdc2HubTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _Fdc2Hub.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2HubTimelockedGovernanceCallCanceled)
				if err := _Fdc2Hub.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_Fdc2Hub *Fdc2HubFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*Fdc2HubTimelockedGovernanceCallCanceled, error) {
	event := new(Fdc2HubTimelockedGovernanceCallCanceled)
	if err := _Fdc2Hub.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2HubTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the Fdc2Hub contract.
type Fdc2HubTimelockedGovernanceCallExecutedIterator struct {
	Event *Fdc2HubTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *Fdc2HubTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2HubTimelockedGovernanceCallExecuted)
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
		it.Event = new(Fdc2HubTimelockedGovernanceCallExecuted)
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
func (it *Fdc2HubTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2HubTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2HubTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the Fdc2Hub contract.
type Fdc2HubTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_Fdc2Hub *Fdc2HubFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*Fdc2HubTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _Fdc2Hub.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &Fdc2HubTimelockedGovernanceCallExecutedIterator{contract: _Fdc2Hub.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_Fdc2Hub *Fdc2HubFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *Fdc2HubTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _Fdc2Hub.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2HubTimelockedGovernanceCallExecuted)
				if err := _Fdc2Hub.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_Fdc2Hub *Fdc2HubFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*Fdc2HubTimelockedGovernanceCallExecuted, error) {
	event := new(Fdc2HubTimelockedGovernanceCallExecuted)
	if err := _Fdc2Hub.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Fdc2HubUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Fdc2Hub contract.
type Fdc2HubUpgradedIterator struct {
	Event *Fdc2HubUpgraded // Event containing the contract specifics and raw log

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
func (it *Fdc2HubUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fdc2HubUpgraded)
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
		it.Event = new(Fdc2HubUpgraded)
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
func (it *Fdc2HubUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fdc2HubUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fdc2HubUpgraded represents a Upgraded event raised by the Fdc2Hub contract.
type Fdc2HubUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Fdc2Hub *Fdc2HubFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*Fdc2HubUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Fdc2Hub.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &Fdc2HubUpgradedIterator{contract: _Fdc2Hub.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Fdc2Hub *Fdc2HubFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *Fdc2HubUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Fdc2Hub.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fdc2HubUpgraded)
				if err := _Fdc2Hub.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Fdc2Hub *Fdc2HubFilterer) ParseUpgraded(log types.Log) (*Fdc2HubUpgraded, error) {
	event := new(Fdc2HubUpgraded)
	if err := _Fdc2Hub.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
