// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepaymentsfeeschedulemanager

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

// ITeePaymentsFeeScheduleManagerFeeSchedule is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsFeeScheduleManagerFeeSchedule struct {
	FactorBIPS   int16
	DelaySeconds uint16
}

// ITeePaymentsFeeScheduleManagerFeeScheduleConfig is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsFeeScheduleManagerFeeScheduleConfig struct {
	MaxSchedules    uint8
	MaxDelaySeconds uint16
}

// ITeePaymentsFeeScheduleManagerFeeScheduleConfigInput is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsFeeScheduleManagerFeeScheduleConfigInput struct {
	MaxDelaySeconds uint16
	MaxSchedules    uint8
	SourceId        [32]byte
}

// ITeePaymentsPMWMultisigAccount is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsPMWMultisigAccount struct {
	SourceId       [32]byte
	AccountAddress string
}

// TeePaymentsFeeScheduleManagerMetaData contains all meta data concerning the TeePaymentsFeeScheduleManager contract.
var TeePaymentsFeeScheduleManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelay\",\"type\":\"uint256\"}],\"name\":\"DelayTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyScheduleNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"FeeScheduleConfigNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeScheduleNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeFactor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"}],\"name\":\"InvalidFeeScheduleConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProjectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SourceLimitsNotConfigured\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySchedules\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"}],\"name\":\"AccountFeeScheduleCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeSchedule[]\",\"name\":\"schedule\",\"type\":\"tuple[]\"}],\"name\":\"AccountFeeScheduleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"sourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"FeeScheduleConfigsCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeScheduleConfigInput[]\",\"name\":\"configs\",\"type\":\"tuple[]\"}],\"name\":\"FeeScheduleConfigsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"ProjectFeeScheduleCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeSchedule[]\",\"name\":\"schedule\",\"type\":\"tuple[]\"}],\"name\":\"ProjectFeeScheduleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"clearAccountFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_sourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"clearFeeScheduleConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"clearProjectFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"contractIIFlareTeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAccountFeeSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_accountHash\",\"type\":\"bytes32\"}],\"name\":\"getEffectiveSchedule\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_feeSchedule\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getFeeScheduleConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeScheduleConfig\",\"name\":\"_config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getProjectFeeSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"name\":\"setAccountFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeScheduleConfigInput[]\",\"name\":\"_configs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeScheduleConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"name\":\"setProjectFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsRegistry\",\"outputs\":[{\"internalType\":\"contractITeePaymentsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"int16[][]\",\"name\":\"_factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"_delaysSeconds\",\"type\":\"uint16[]\"}],\"name\":\"validateAndEncodeSchedules\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_encodedPerPayment\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TeePaymentsFeeScheduleManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeePaymentsFeeScheduleManagerMetaData.ABI instead.
var TeePaymentsFeeScheduleManagerABI = TeePaymentsFeeScheduleManagerMetaData.ABI

// TeePaymentsFeeScheduleManager is an auto generated Go binding around an Ethereum contract.
type TeePaymentsFeeScheduleManager struct {
	TeePaymentsFeeScheduleManagerCaller     // Read-only binding to the contract
	TeePaymentsFeeScheduleManagerTransactor // Write-only binding to the contract
	TeePaymentsFeeScheduleManagerFilterer   // Log filterer for contract events
}

// TeePaymentsFeeScheduleManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeePaymentsFeeScheduleManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsFeeScheduleManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeePaymentsFeeScheduleManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsFeeScheduleManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeePaymentsFeeScheduleManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsFeeScheduleManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeePaymentsFeeScheduleManagerSession struct {
	Contract     *TeePaymentsFeeScheduleManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TeePaymentsFeeScheduleManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeePaymentsFeeScheduleManagerCallerSession struct {
	Contract *TeePaymentsFeeScheduleManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// TeePaymentsFeeScheduleManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeePaymentsFeeScheduleManagerTransactorSession struct {
	Contract     *TeePaymentsFeeScheduleManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// TeePaymentsFeeScheduleManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeePaymentsFeeScheduleManagerRaw struct {
	Contract *TeePaymentsFeeScheduleManager // Generic contract binding to access the raw methods on
}

// TeePaymentsFeeScheduleManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeePaymentsFeeScheduleManagerCallerRaw struct {
	Contract *TeePaymentsFeeScheduleManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TeePaymentsFeeScheduleManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeePaymentsFeeScheduleManagerTransactorRaw struct {
	Contract *TeePaymentsFeeScheduleManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeePaymentsFeeScheduleManager creates a new instance of TeePaymentsFeeScheduleManager, bound to a specific deployed contract.
func NewTeePaymentsFeeScheduleManager(address common.Address, backend bind.ContractBackend) (*TeePaymentsFeeScheduleManager, error) {
	contract, err := bindTeePaymentsFeeScheduleManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManager{TeePaymentsFeeScheduleManagerCaller: TeePaymentsFeeScheduleManagerCaller{contract: contract}, TeePaymentsFeeScheduleManagerTransactor: TeePaymentsFeeScheduleManagerTransactor{contract: contract}, TeePaymentsFeeScheduleManagerFilterer: TeePaymentsFeeScheduleManagerFilterer{contract: contract}}, nil
}

// NewTeePaymentsFeeScheduleManagerCaller creates a new read-only instance of TeePaymentsFeeScheduleManager, bound to a specific deployed contract.
func NewTeePaymentsFeeScheduleManagerCaller(address common.Address, caller bind.ContractCaller) (*TeePaymentsFeeScheduleManagerCaller, error) {
	contract, err := bindTeePaymentsFeeScheduleManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerCaller{contract: contract}, nil
}

// NewTeePaymentsFeeScheduleManagerTransactor creates a new write-only instance of TeePaymentsFeeScheduleManager, bound to a specific deployed contract.
func NewTeePaymentsFeeScheduleManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeePaymentsFeeScheduleManagerTransactor, error) {
	contract, err := bindTeePaymentsFeeScheduleManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerTransactor{contract: contract}, nil
}

// NewTeePaymentsFeeScheduleManagerFilterer creates a new log filterer instance of TeePaymentsFeeScheduleManager, bound to a specific deployed contract.
func NewTeePaymentsFeeScheduleManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeePaymentsFeeScheduleManagerFilterer, error) {
	contract, err := bindTeePaymentsFeeScheduleManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerFilterer{contract: contract}, nil
}

// bindTeePaymentsFeeScheduleManager binds a generic wrapper to an already deployed contract.
func bindTeePaymentsFeeScheduleManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeePaymentsFeeScheduleManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePaymentsFeeScheduleManager.Contract.TeePaymentsFeeScheduleManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.TeePaymentsFeeScheduleManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.TeePaymentsFeeScheduleManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePaymentsFeeScheduleManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePaymentsFeeScheduleManager.Contract.UPGRADEINTERFACEVERSION(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePaymentsFeeScheduleManager.Contract.UPGRADEINTERFACEVERSION(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) FlareTeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "flareTeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) FlareTeeManager() (common.Address, error) {
	return _TeePaymentsFeeScheduleManager.Contract.FlareTeeManager(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) FlareTeeManager() (common.Address, error) {
	return _TeePaymentsFeeScheduleManager.Contract.FlareTeeManager(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// GetAccountFeeSchedule is a free data retrieval call binding the contract method 0x3c9e0c3f.
//
// Solidity: function getAccountFeeSchedule((bytes32,string) _account) view returns((int16,uint16)[] _schedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) GetAccountFeeSchedule(opts *bind.CallOpts, _account ITeePaymentsPMWMultisigAccount) ([]ITeePaymentsFeeScheduleManagerFeeSchedule, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "getAccountFeeSchedule", _account)

	if err != nil {
		return *new([]ITeePaymentsFeeScheduleManagerFeeSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITeePaymentsFeeScheduleManagerFeeSchedule)).(*[]ITeePaymentsFeeScheduleManagerFeeSchedule)

	return out0, err

}

// GetAccountFeeSchedule is a free data retrieval call binding the contract method 0x3c9e0c3f.
//
// Solidity: function getAccountFeeSchedule((bytes32,string) _account) view returns((int16,uint16)[] _schedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) GetAccountFeeSchedule(_account ITeePaymentsPMWMultisigAccount) ([]ITeePaymentsFeeScheduleManagerFeeSchedule, error) {
	return _TeePaymentsFeeScheduleManager.Contract.GetAccountFeeSchedule(&_TeePaymentsFeeScheduleManager.CallOpts, _account)
}

// GetAccountFeeSchedule is a free data retrieval call binding the contract method 0x3c9e0c3f.
//
// Solidity: function getAccountFeeSchedule((bytes32,string) _account) view returns((int16,uint16)[] _schedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) GetAccountFeeSchedule(_account ITeePaymentsPMWMultisigAccount) ([]ITeePaymentsFeeScheduleManagerFeeSchedule, error) {
	return _TeePaymentsFeeScheduleManager.Contract.GetAccountFeeSchedule(&_TeePaymentsFeeScheduleManager.CallOpts, _account)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) GetAddressUpdater() (common.Address, error) {
	return _TeePaymentsFeeScheduleManager.Contract.GetAddressUpdater(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeePaymentsFeeScheduleManager.Contract.GetAddressUpdater(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// GetEffectiveSchedule is a free data retrieval call binding the contract method 0xed8b6f51.
//
// Solidity: function getEffectiveSchedule(bytes32 _projectId, bytes32 _sourceId, bytes32 _accountHash) view returns(bytes _feeSchedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) GetEffectiveSchedule(opts *bind.CallOpts, _projectId [32]byte, _sourceId [32]byte, _accountHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "getEffectiveSchedule", _projectId, _sourceId, _accountHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEffectiveSchedule is a free data retrieval call binding the contract method 0xed8b6f51.
//
// Solidity: function getEffectiveSchedule(bytes32 _projectId, bytes32 _sourceId, bytes32 _accountHash) view returns(bytes _feeSchedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) GetEffectiveSchedule(_projectId [32]byte, _sourceId [32]byte, _accountHash [32]byte) ([]byte, error) {
	return _TeePaymentsFeeScheduleManager.Contract.GetEffectiveSchedule(&_TeePaymentsFeeScheduleManager.CallOpts, _projectId, _sourceId, _accountHash)
}

// GetEffectiveSchedule is a free data retrieval call binding the contract method 0xed8b6f51.
//
// Solidity: function getEffectiveSchedule(bytes32 _projectId, bytes32 _sourceId, bytes32 _accountHash) view returns(bytes _feeSchedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) GetEffectiveSchedule(_projectId [32]byte, _sourceId [32]byte, _accountHash [32]byte) ([]byte, error) {
	return _TeePaymentsFeeScheduleManager.Contract.GetEffectiveSchedule(&_TeePaymentsFeeScheduleManager.CallOpts, _projectId, _sourceId, _accountHash)
}

// GetFeeScheduleConfig is a free data retrieval call binding the contract method 0x16743f21.
//
// Solidity: function getFeeScheduleConfig(bytes32 _sourceId) view returns((uint8,uint16) _config)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) GetFeeScheduleConfig(opts *bind.CallOpts, _sourceId [32]byte) (ITeePaymentsFeeScheduleManagerFeeScheduleConfig, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "getFeeScheduleConfig", _sourceId)

	if err != nil {
		return *new(ITeePaymentsFeeScheduleManagerFeeScheduleConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeePaymentsFeeScheduleManagerFeeScheduleConfig)).(*ITeePaymentsFeeScheduleManagerFeeScheduleConfig)

	return out0, err

}

// GetFeeScheduleConfig is a free data retrieval call binding the contract method 0x16743f21.
//
// Solidity: function getFeeScheduleConfig(bytes32 _sourceId) view returns((uint8,uint16) _config)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) GetFeeScheduleConfig(_sourceId [32]byte) (ITeePaymentsFeeScheduleManagerFeeScheduleConfig, error) {
	return _TeePaymentsFeeScheduleManager.Contract.GetFeeScheduleConfig(&_TeePaymentsFeeScheduleManager.CallOpts, _sourceId)
}

// GetFeeScheduleConfig is a free data retrieval call binding the contract method 0x16743f21.
//
// Solidity: function getFeeScheduleConfig(bytes32 _sourceId) view returns((uint8,uint16) _config)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) GetFeeScheduleConfig(_sourceId [32]byte) (ITeePaymentsFeeScheduleManagerFeeScheduleConfig, error) {
	return _TeePaymentsFeeScheduleManager.Contract.GetFeeScheduleConfig(&_TeePaymentsFeeScheduleManager.CallOpts, _sourceId)
}

// GetProjectFeeSchedule is a free data retrieval call binding the contract method 0x30d1bb93.
//
// Solidity: function getProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) view returns((int16,uint16)[] _schedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) GetProjectFeeSchedule(opts *bind.CallOpts, _projectId [32]byte, _sourceId [32]byte) ([]ITeePaymentsFeeScheduleManagerFeeSchedule, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "getProjectFeeSchedule", _projectId, _sourceId)

	if err != nil {
		return *new([]ITeePaymentsFeeScheduleManagerFeeSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITeePaymentsFeeScheduleManagerFeeSchedule)).(*[]ITeePaymentsFeeScheduleManagerFeeSchedule)

	return out0, err

}

// GetProjectFeeSchedule is a free data retrieval call binding the contract method 0x30d1bb93.
//
// Solidity: function getProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) view returns((int16,uint16)[] _schedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) GetProjectFeeSchedule(_projectId [32]byte, _sourceId [32]byte) ([]ITeePaymentsFeeScheduleManagerFeeSchedule, error) {
	return _TeePaymentsFeeScheduleManager.Contract.GetProjectFeeSchedule(&_TeePaymentsFeeScheduleManager.CallOpts, _projectId, _sourceId)
}

// GetProjectFeeSchedule is a free data retrieval call binding the contract method 0x30d1bb93.
//
// Solidity: function getProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) view returns((int16,uint16)[] _schedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) GetProjectFeeSchedule(_projectId [32]byte, _sourceId [32]byte) ([]ITeePaymentsFeeScheduleManagerFeeSchedule, error) {
	return _TeePaymentsFeeScheduleManager.Contract.GetProjectFeeSchedule(&_TeePaymentsFeeScheduleManager.CallOpts, _projectId, _sourceId)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) Governance() (common.Address, error) {
	return _TeePaymentsFeeScheduleManager.Contract.Governance(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) Governance() (common.Address, error) {
	return _TeePaymentsFeeScheduleManager.Contract.Governance(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) GovernanceSettings() (common.Address, error) {
	return _TeePaymentsFeeScheduleManager.Contract.GovernanceSettings(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeePaymentsFeeScheduleManager.Contract.GovernanceSettings(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) Implementation() (common.Address, error) {
	return _TeePaymentsFeeScheduleManager.Contract.Implementation(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) Implementation() (common.Address, error) {
	return _TeePaymentsFeeScheduleManager.Contract.Implementation(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePaymentsFeeScheduleManager.Contract.IsExecutor(&_TeePaymentsFeeScheduleManager.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePaymentsFeeScheduleManager.Contract.IsExecutor(&_TeePaymentsFeeScheduleManager.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) ProductionMode() (bool, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ProductionMode(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) ProductionMode() (bool, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ProductionMode(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) ProxiableUUID() ([32]byte, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ProxiableUUID(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ProxiableUUID(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) TeePaymentsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "teePaymentsRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) TeePaymentsRegistry() (common.Address, error) {
	return _TeePaymentsFeeScheduleManager.Contract.TeePaymentsRegistry(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) TeePaymentsRegistry() (common.Address, error) {
	return _TeePaymentsFeeScheduleManager.Contract.TeePaymentsRegistry(&_TeePaymentsFeeScheduleManager.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeePaymentsFeeScheduleManager.Contract.TimelockedCalls(&_TeePaymentsFeeScheduleManager.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeePaymentsFeeScheduleManager.Contract.TimelockedCalls(&_TeePaymentsFeeScheduleManager.CallOpts, selector)
}

// ValidateAndEncodeSchedules is a free data retrieval call binding the contract method 0x3d699135.
//
// Solidity: function validateAndEncodeSchedules(bytes32 _sourceId, int16[][] _factorsBIPSPerPayment, uint16[] _delaysSeconds) view returns(bytes[] _encodedPerPayment)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCaller) ValidateAndEncodeSchedules(opts *bind.CallOpts, _sourceId [32]byte, _factorsBIPSPerPayment [][]int16, _delaysSeconds []uint16) ([][]byte, error) {
	var out []interface{}
	err := _TeePaymentsFeeScheduleManager.contract.Call(opts, &out, "validateAndEncodeSchedules", _sourceId, _factorsBIPSPerPayment, _delaysSeconds)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// ValidateAndEncodeSchedules is a free data retrieval call binding the contract method 0x3d699135.
//
// Solidity: function validateAndEncodeSchedules(bytes32 _sourceId, int16[][] _factorsBIPSPerPayment, uint16[] _delaysSeconds) view returns(bytes[] _encodedPerPayment)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) ValidateAndEncodeSchedules(_sourceId [32]byte, _factorsBIPSPerPayment [][]int16, _delaysSeconds []uint16) ([][]byte, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ValidateAndEncodeSchedules(&_TeePaymentsFeeScheduleManager.CallOpts, _sourceId, _factorsBIPSPerPayment, _delaysSeconds)
}

// ValidateAndEncodeSchedules is a free data retrieval call binding the contract method 0x3d699135.
//
// Solidity: function validateAndEncodeSchedules(bytes32 _sourceId, int16[][] _factorsBIPSPerPayment, uint16[] _delaysSeconds) view returns(bytes[] _encodedPerPayment)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerCallerSession) ValidateAndEncodeSchedules(_sourceId [32]byte, _factorsBIPSPerPayment [][]int16, _delaysSeconds []uint16) ([][]byte, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ValidateAndEncodeSchedules(&_TeePaymentsFeeScheduleManager.CallOpts, _sourceId, _factorsBIPSPerPayment, _delaysSeconds)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.CancelGovernanceCall(&_TeePaymentsFeeScheduleManager.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.CancelGovernanceCall(&_TeePaymentsFeeScheduleManager.TransactOpts, _selector)
}

// ClearAccountFeeSchedule is a paid mutator transaction binding the contract method 0xf7443f84.
//
// Solidity: function clearAccountFeeSchedule((bytes32,string) _account) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactor) ClearAccountFeeSchedule(opts *bind.TransactOpts, _account ITeePaymentsPMWMultisigAccount) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.contract.Transact(opts, "clearAccountFeeSchedule", _account)
}

// ClearAccountFeeSchedule is a paid mutator transaction binding the contract method 0xf7443f84.
//
// Solidity: function clearAccountFeeSchedule((bytes32,string) _account) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) ClearAccountFeeSchedule(_account ITeePaymentsPMWMultisigAccount) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ClearAccountFeeSchedule(&_TeePaymentsFeeScheduleManager.TransactOpts, _account)
}

// ClearAccountFeeSchedule is a paid mutator transaction binding the contract method 0xf7443f84.
//
// Solidity: function clearAccountFeeSchedule((bytes32,string) _account) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorSession) ClearAccountFeeSchedule(_account ITeePaymentsPMWMultisigAccount) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ClearAccountFeeSchedule(&_TeePaymentsFeeScheduleManager.TransactOpts, _account)
}

// ClearFeeScheduleConfigs is a paid mutator transaction binding the contract method 0x11a5c047.
//
// Solidity: function clearFeeScheduleConfigs(bytes32[] _sourceIds) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactor) ClearFeeScheduleConfigs(opts *bind.TransactOpts, _sourceIds [][32]byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.contract.Transact(opts, "clearFeeScheduleConfigs", _sourceIds)
}

// ClearFeeScheduleConfigs is a paid mutator transaction binding the contract method 0x11a5c047.
//
// Solidity: function clearFeeScheduleConfigs(bytes32[] _sourceIds) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) ClearFeeScheduleConfigs(_sourceIds [][32]byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ClearFeeScheduleConfigs(&_TeePaymentsFeeScheduleManager.TransactOpts, _sourceIds)
}

// ClearFeeScheduleConfigs is a paid mutator transaction binding the contract method 0x11a5c047.
//
// Solidity: function clearFeeScheduleConfigs(bytes32[] _sourceIds) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorSession) ClearFeeScheduleConfigs(_sourceIds [][32]byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ClearFeeScheduleConfigs(&_TeePaymentsFeeScheduleManager.TransactOpts, _sourceIds)
}

// ClearProjectFeeSchedule is a paid mutator transaction binding the contract method 0x7aecd6f7.
//
// Solidity: function clearProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactor) ClearProjectFeeSchedule(opts *bind.TransactOpts, _projectId [32]byte, _sourceId [32]byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.contract.Transact(opts, "clearProjectFeeSchedule", _projectId, _sourceId)
}

// ClearProjectFeeSchedule is a paid mutator transaction binding the contract method 0x7aecd6f7.
//
// Solidity: function clearProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) ClearProjectFeeSchedule(_projectId [32]byte, _sourceId [32]byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ClearProjectFeeSchedule(&_TeePaymentsFeeScheduleManager.TransactOpts, _projectId, _sourceId)
}

// ClearProjectFeeSchedule is a paid mutator transaction binding the contract method 0x7aecd6f7.
//
// Solidity: function clearProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorSession) ClearProjectFeeSchedule(_projectId [32]byte, _sourceId [32]byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ClearProjectFeeSchedule(&_TeePaymentsFeeScheduleManager.TransactOpts, _projectId, _sourceId)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ExecuteGovernanceCall(&_TeePaymentsFeeScheduleManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.ExecuteGovernanceCall(&_TeePaymentsFeeScheduleManager.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.Initialise(&_TeePaymentsFeeScheduleManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.Initialise(&_TeePaymentsFeeScheduleManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.Initialize(&_TeePaymentsFeeScheduleManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.Initialize(&_TeePaymentsFeeScheduleManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// SetAccountFeeSchedule is a paid mutator transaction binding the contract method 0x32ab4577.
//
// Solidity: function setAccountFeeSchedule((bytes32,string) _account, (int16,uint16)[] _schedule) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactor) SetAccountFeeSchedule(opts *bind.TransactOpts, _account ITeePaymentsPMWMultisigAccount, _schedule []ITeePaymentsFeeScheduleManagerFeeSchedule) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.contract.Transact(opts, "setAccountFeeSchedule", _account, _schedule)
}

// SetAccountFeeSchedule is a paid mutator transaction binding the contract method 0x32ab4577.
//
// Solidity: function setAccountFeeSchedule((bytes32,string) _account, (int16,uint16)[] _schedule) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) SetAccountFeeSchedule(_account ITeePaymentsPMWMultisigAccount, _schedule []ITeePaymentsFeeScheduleManagerFeeSchedule) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.SetAccountFeeSchedule(&_TeePaymentsFeeScheduleManager.TransactOpts, _account, _schedule)
}

// SetAccountFeeSchedule is a paid mutator transaction binding the contract method 0x32ab4577.
//
// Solidity: function setAccountFeeSchedule((bytes32,string) _account, (int16,uint16)[] _schedule) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorSession) SetAccountFeeSchedule(_account ITeePaymentsPMWMultisigAccount, _schedule []ITeePaymentsFeeScheduleManagerFeeSchedule) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.SetAccountFeeSchedule(&_TeePaymentsFeeScheduleManager.TransactOpts, _account, _schedule)
}

// SetFeeScheduleConfigs is a paid mutator transaction binding the contract method 0x9b7f8f29.
//
// Solidity: function setFeeScheduleConfigs((uint16,uint8,bytes32)[] _configs) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactor) SetFeeScheduleConfigs(opts *bind.TransactOpts, _configs []ITeePaymentsFeeScheduleManagerFeeScheduleConfigInput) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.contract.Transact(opts, "setFeeScheduleConfigs", _configs)
}

// SetFeeScheduleConfigs is a paid mutator transaction binding the contract method 0x9b7f8f29.
//
// Solidity: function setFeeScheduleConfigs((uint16,uint8,bytes32)[] _configs) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) SetFeeScheduleConfigs(_configs []ITeePaymentsFeeScheduleManagerFeeScheduleConfigInput) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.SetFeeScheduleConfigs(&_TeePaymentsFeeScheduleManager.TransactOpts, _configs)
}

// SetFeeScheduleConfigs is a paid mutator transaction binding the contract method 0x9b7f8f29.
//
// Solidity: function setFeeScheduleConfigs((uint16,uint8,bytes32)[] _configs) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorSession) SetFeeScheduleConfigs(_configs []ITeePaymentsFeeScheduleManagerFeeScheduleConfigInput) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.SetFeeScheduleConfigs(&_TeePaymentsFeeScheduleManager.TransactOpts, _configs)
}

// SetProjectFeeSchedule is a paid mutator transaction binding the contract method 0xe27d19bd.
//
// Solidity: function setProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId, (int16,uint16)[] _schedule) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactor) SetProjectFeeSchedule(opts *bind.TransactOpts, _projectId [32]byte, _sourceId [32]byte, _schedule []ITeePaymentsFeeScheduleManagerFeeSchedule) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.contract.Transact(opts, "setProjectFeeSchedule", _projectId, _sourceId, _schedule)
}

// SetProjectFeeSchedule is a paid mutator transaction binding the contract method 0xe27d19bd.
//
// Solidity: function setProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId, (int16,uint16)[] _schedule) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) SetProjectFeeSchedule(_projectId [32]byte, _sourceId [32]byte, _schedule []ITeePaymentsFeeScheduleManagerFeeSchedule) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.SetProjectFeeSchedule(&_TeePaymentsFeeScheduleManager.TransactOpts, _projectId, _sourceId, _schedule)
}

// SetProjectFeeSchedule is a paid mutator transaction binding the contract method 0xe27d19bd.
//
// Solidity: function setProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId, (int16,uint16)[] _schedule) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorSession) SetProjectFeeSchedule(_projectId [32]byte, _sourceId [32]byte, _schedule []ITeePaymentsFeeScheduleManagerFeeSchedule) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.SetProjectFeeSchedule(&_TeePaymentsFeeScheduleManager.TransactOpts, _projectId, _sourceId, _schedule)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.SwitchToProductionMode(&_TeePaymentsFeeScheduleManager.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.SwitchToProductionMode(&_TeePaymentsFeeScheduleManager.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.UpdateContractAddresses(&_TeePaymentsFeeScheduleManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.UpdateContractAddresses(&_TeePaymentsFeeScheduleManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.UpgradeToAndCall(&_TeePaymentsFeeScheduleManager.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsFeeScheduleManager.Contract.UpgradeToAndCall(&_TeePaymentsFeeScheduleManager.TransactOpts, _newImplementation, _data)
}

// TeePaymentsFeeScheduleManagerAccountFeeScheduleClearedIterator is returned from FilterAccountFeeScheduleCleared and is used to iterate over the raw logs and unpacked data for AccountFeeScheduleCleared events raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerAccountFeeScheduleClearedIterator struct {
	Event *TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared // Event containing the contract specifics and raw log

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
func (it *TeePaymentsFeeScheduleManagerAccountFeeScheduleClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared)
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
		it.Event = new(TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared)
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
func (it *TeePaymentsFeeScheduleManagerAccountFeeScheduleClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsFeeScheduleManagerAccountFeeScheduleClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared represents a AccountFeeScheduleCleared event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared struct {
	ProjectId      [32]byte
	SourceId       [32]byte
	AccountAddress string
	AccountHash    [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountFeeScheduleCleared is a free log retrieval operation binding the contract event 0x2426e08a74c1afcafde131339b2e2fa2289d879322564f7e49d16bfe6bb0fd7b.
//
// Solidity: event AccountFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) FilterAccountFeeScheduleCleared(opts *bind.FilterOpts, projectId [][32]byte, sourceId [][32]byte, accountHash [][32]byte) (*TeePaymentsFeeScheduleManagerAccountFeeScheduleClearedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	var accountHashRule []interface{}
	for _, accountHashItem := range accountHash {
		accountHashRule = append(accountHashRule, accountHashItem)
	}

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.FilterLogs(opts, "AccountFeeScheduleCleared", projectIdRule, sourceIdRule, accountHashRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerAccountFeeScheduleClearedIterator{contract: _TeePaymentsFeeScheduleManager.contract, event: "AccountFeeScheduleCleared", logs: logs, sub: sub}, nil
}

// WatchAccountFeeScheduleCleared is a free log subscription operation binding the contract event 0x2426e08a74c1afcafde131339b2e2fa2289d879322564f7e49d16bfe6bb0fd7b.
//
// Solidity: event AccountFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) WatchAccountFeeScheduleCleared(opts *bind.WatchOpts, sink chan<- *TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared, projectId [][32]byte, sourceId [][32]byte, accountHash [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	var accountHashRule []interface{}
	for _, accountHashItem := range accountHash {
		accountHashRule = append(accountHashRule, accountHashItem)
	}

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.WatchLogs(opts, "AccountFeeScheduleCleared", projectIdRule, sourceIdRule, accountHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared)
				if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "AccountFeeScheduleCleared", log); err != nil {
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

// ParseAccountFeeScheduleCleared is a log parse operation binding the contract event 0x2426e08a74c1afcafde131339b2e2fa2289d879322564f7e49d16bfe6bb0fd7b.
//
// Solidity: event AccountFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) ParseAccountFeeScheduleCleared(log types.Log) (*TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared, error) {
	event := new(TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared)
	if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "AccountFeeScheduleCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsFeeScheduleManagerAccountFeeScheduleSetIterator is returned from FilterAccountFeeScheduleSet and is used to iterate over the raw logs and unpacked data for AccountFeeScheduleSet events raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerAccountFeeScheduleSetIterator struct {
	Event *TeePaymentsFeeScheduleManagerAccountFeeScheduleSet // Event containing the contract specifics and raw log

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
func (it *TeePaymentsFeeScheduleManagerAccountFeeScheduleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsFeeScheduleManagerAccountFeeScheduleSet)
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
		it.Event = new(TeePaymentsFeeScheduleManagerAccountFeeScheduleSet)
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
func (it *TeePaymentsFeeScheduleManagerAccountFeeScheduleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsFeeScheduleManagerAccountFeeScheduleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsFeeScheduleManagerAccountFeeScheduleSet represents a AccountFeeScheduleSet event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerAccountFeeScheduleSet struct {
	ProjectId      [32]byte
	SourceId       [32]byte
	AccountAddress string
	AccountHash    [32]byte
	Schedule       []ITeePaymentsFeeScheduleManagerFeeSchedule
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountFeeScheduleSet is a free log retrieval operation binding the contract event 0x84562203b55c4f8e7837356a2b059908a1da339c2081ce5d934d3e0fb119b98f.
//
// Solidity: event AccountFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash, (int16,uint16)[] schedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) FilterAccountFeeScheduleSet(opts *bind.FilterOpts, projectId [][32]byte, sourceId [][32]byte, accountHash [][32]byte) (*TeePaymentsFeeScheduleManagerAccountFeeScheduleSetIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	var accountHashRule []interface{}
	for _, accountHashItem := range accountHash {
		accountHashRule = append(accountHashRule, accountHashItem)
	}

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.FilterLogs(opts, "AccountFeeScheduleSet", projectIdRule, sourceIdRule, accountHashRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerAccountFeeScheduleSetIterator{contract: _TeePaymentsFeeScheduleManager.contract, event: "AccountFeeScheduleSet", logs: logs, sub: sub}, nil
}

// WatchAccountFeeScheduleSet is a free log subscription operation binding the contract event 0x84562203b55c4f8e7837356a2b059908a1da339c2081ce5d934d3e0fb119b98f.
//
// Solidity: event AccountFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash, (int16,uint16)[] schedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) WatchAccountFeeScheduleSet(opts *bind.WatchOpts, sink chan<- *TeePaymentsFeeScheduleManagerAccountFeeScheduleSet, projectId [][32]byte, sourceId [][32]byte, accountHash [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	var accountHashRule []interface{}
	for _, accountHashItem := range accountHash {
		accountHashRule = append(accountHashRule, accountHashItem)
	}

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.WatchLogs(opts, "AccountFeeScheduleSet", projectIdRule, sourceIdRule, accountHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsFeeScheduleManagerAccountFeeScheduleSet)
				if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "AccountFeeScheduleSet", log); err != nil {
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

// ParseAccountFeeScheduleSet is a log parse operation binding the contract event 0x84562203b55c4f8e7837356a2b059908a1da339c2081ce5d934d3e0fb119b98f.
//
// Solidity: event AccountFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash, (int16,uint16)[] schedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) ParseAccountFeeScheduleSet(log types.Log) (*TeePaymentsFeeScheduleManagerAccountFeeScheduleSet, error) {
	event := new(TeePaymentsFeeScheduleManagerAccountFeeScheduleSet)
	if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "AccountFeeScheduleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsFeeScheduleManagerFeeScheduleConfigsClearedIterator is returned from FilterFeeScheduleConfigsCleared and is used to iterate over the raw logs and unpacked data for FeeScheduleConfigsCleared events raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerFeeScheduleConfigsClearedIterator struct {
	Event *TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared // Event containing the contract specifics and raw log

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
func (it *TeePaymentsFeeScheduleManagerFeeScheduleConfigsClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared)
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
		it.Event = new(TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared)
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
func (it *TeePaymentsFeeScheduleManagerFeeScheduleConfigsClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsFeeScheduleManagerFeeScheduleConfigsClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared represents a FeeScheduleConfigsCleared event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared struct {
	SourceIds [][32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeScheduleConfigsCleared is a free log retrieval operation binding the contract event 0xa65a10cbf89a889142f7004143171a2381448115ee506c5a8a1c4c8444f76016.
//
// Solidity: event FeeScheduleConfigsCleared(bytes32[] sourceIds)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) FilterFeeScheduleConfigsCleared(opts *bind.FilterOpts) (*TeePaymentsFeeScheduleManagerFeeScheduleConfigsClearedIterator, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.FilterLogs(opts, "FeeScheduleConfigsCleared")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerFeeScheduleConfigsClearedIterator{contract: _TeePaymentsFeeScheduleManager.contract, event: "FeeScheduleConfigsCleared", logs: logs, sub: sub}, nil
}

// WatchFeeScheduleConfigsCleared is a free log subscription operation binding the contract event 0xa65a10cbf89a889142f7004143171a2381448115ee506c5a8a1c4c8444f76016.
//
// Solidity: event FeeScheduleConfigsCleared(bytes32[] sourceIds)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) WatchFeeScheduleConfigsCleared(opts *bind.WatchOpts, sink chan<- *TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.WatchLogs(opts, "FeeScheduleConfigsCleared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared)
				if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "FeeScheduleConfigsCleared", log); err != nil {
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

// ParseFeeScheduleConfigsCleared is a log parse operation binding the contract event 0xa65a10cbf89a889142f7004143171a2381448115ee506c5a8a1c4c8444f76016.
//
// Solidity: event FeeScheduleConfigsCleared(bytes32[] sourceIds)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) ParseFeeScheduleConfigsCleared(log types.Log) (*TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared, error) {
	event := new(TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared)
	if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "FeeScheduleConfigsCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsFeeScheduleManagerFeeScheduleConfigsSetIterator is returned from FilterFeeScheduleConfigsSet and is used to iterate over the raw logs and unpacked data for FeeScheduleConfigsSet events raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerFeeScheduleConfigsSetIterator struct {
	Event *TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet // Event containing the contract specifics and raw log

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
func (it *TeePaymentsFeeScheduleManagerFeeScheduleConfigsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet)
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
		it.Event = new(TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet)
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
func (it *TeePaymentsFeeScheduleManagerFeeScheduleConfigsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsFeeScheduleManagerFeeScheduleConfigsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet represents a FeeScheduleConfigsSet event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet struct {
	Configs []ITeePaymentsFeeScheduleManagerFeeScheduleConfigInput
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeScheduleConfigsSet is a free log retrieval operation binding the contract event 0x6a0718b26b38a09770c97425ea4ecdb442e8d8ab07960c3e00fd32c6397646c4.
//
// Solidity: event FeeScheduleConfigsSet((uint16,uint8,bytes32)[] configs)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) FilterFeeScheduleConfigsSet(opts *bind.FilterOpts) (*TeePaymentsFeeScheduleManagerFeeScheduleConfigsSetIterator, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.FilterLogs(opts, "FeeScheduleConfigsSet")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerFeeScheduleConfigsSetIterator{contract: _TeePaymentsFeeScheduleManager.contract, event: "FeeScheduleConfigsSet", logs: logs, sub: sub}, nil
}

// WatchFeeScheduleConfigsSet is a free log subscription operation binding the contract event 0x6a0718b26b38a09770c97425ea4ecdb442e8d8ab07960c3e00fd32c6397646c4.
//
// Solidity: event FeeScheduleConfigsSet((uint16,uint8,bytes32)[] configs)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) WatchFeeScheduleConfigsSet(opts *bind.WatchOpts, sink chan<- *TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.WatchLogs(opts, "FeeScheduleConfigsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet)
				if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "FeeScheduleConfigsSet", log); err != nil {
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

// ParseFeeScheduleConfigsSet is a log parse operation binding the contract event 0x6a0718b26b38a09770c97425ea4ecdb442e8d8ab07960c3e00fd32c6397646c4.
//
// Solidity: event FeeScheduleConfigsSet((uint16,uint8,bytes32)[] configs)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) ParseFeeScheduleConfigsSet(log types.Log) (*TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet, error) {
	event := new(TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet)
	if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "FeeScheduleConfigsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsFeeScheduleManagerGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerGovernanceCallTimelockedIterator struct {
	Event *TeePaymentsFeeScheduleManagerGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeePaymentsFeeScheduleManagerGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsFeeScheduleManagerGovernanceCallTimelocked)
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
		it.Event = new(TeePaymentsFeeScheduleManagerGovernanceCallTimelocked)
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
func (it *TeePaymentsFeeScheduleManagerGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsFeeScheduleManagerGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsFeeScheduleManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeePaymentsFeeScheduleManagerGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerGovernanceCallTimelockedIterator{contract: _TeePaymentsFeeScheduleManager.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeePaymentsFeeScheduleManagerGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsFeeScheduleManagerGovernanceCallTimelocked)
				if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeePaymentsFeeScheduleManagerGovernanceCallTimelocked, error) {
	event := new(TeePaymentsFeeScheduleManagerGovernanceCallTimelocked)
	if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsFeeScheduleManagerGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerGovernanceInitialisedIterator struct {
	Event *TeePaymentsFeeScheduleManagerGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeePaymentsFeeScheduleManagerGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsFeeScheduleManagerGovernanceInitialised)
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
		it.Event = new(TeePaymentsFeeScheduleManagerGovernanceInitialised)
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
func (it *TeePaymentsFeeScheduleManagerGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsFeeScheduleManagerGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsFeeScheduleManagerGovernanceInitialised represents a GovernanceInitialised event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeePaymentsFeeScheduleManagerGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerGovernanceInitialisedIterator{contract: _TeePaymentsFeeScheduleManager.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeePaymentsFeeScheduleManagerGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsFeeScheduleManagerGovernanceInitialised)
				if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) ParseGovernanceInitialised(log types.Log) (*TeePaymentsFeeScheduleManagerGovernanceInitialised, error) {
	event := new(TeePaymentsFeeScheduleManagerGovernanceInitialised)
	if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsFeeScheduleManagerGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerGovernedProductionModeEnteredIterator struct {
	Event *TeePaymentsFeeScheduleManagerGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeePaymentsFeeScheduleManagerGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsFeeScheduleManagerGovernedProductionModeEntered)
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
		it.Event = new(TeePaymentsFeeScheduleManagerGovernedProductionModeEntered)
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
func (it *TeePaymentsFeeScheduleManagerGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsFeeScheduleManagerGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsFeeScheduleManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeePaymentsFeeScheduleManagerGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerGovernedProductionModeEnteredIterator{contract: _TeePaymentsFeeScheduleManager.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeePaymentsFeeScheduleManagerGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsFeeScheduleManagerGovernedProductionModeEntered)
				if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeePaymentsFeeScheduleManagerGovernedProductionModeEntered, error) {
	event := new(TeePaymentsFeeScheduleManagerGovernedProductionModeEntered)
	if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsFeeScheduleManagerProjectFeeScheduleClearedIterator is returned from FilterProjectFeeScheduleCleared and is used to iterate over the raw logs and unpacked data for ProjectFeeScheduleCleared events raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerProjectFeeScheduleClearedIterator struct {
	Event *TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared // Event containing the contract specifics and raw log

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
func (it *TeePaymentsFeeScheduleManagerProjectFeeScheduleClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared)
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
		it.Event = new(TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared)
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
func (it *TeePaymentsFeeScheduleManagerProjectFeeScheduleClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsFeeScheduleManagerProjectFeeScheduleClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared represents a ProjectFeeScheduleCleared event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared struct {
	ProjectId [32]byte
	SourceId  [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectFeeScheduleCleared is a free log retrieval operation binding the contract event 0xe3f22e2cb60f792203c33a8b3ce8faae7df971c5bd56aa1beca573c904909aa2.
//
// Solidity: event ProjectFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) FilterProjectFeeScheduleCleared(opts *bind.FilterOpts, projectId [][32]byte, sourceId [][32]byte) (*TeePaymentsFeeScheduleManagerProjectFeeScheduleClearedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.FilterLogs(opts, "ProjectFeeScheduleCleared", projectIdRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerProjectFeeScheduleClearedIterator{contract: _TeePaymentsFeeScheduleManager.contract, event: "ProjectFeeScheduleCleared", logs: logs, sub: sub}, nil
}

// WatchProjectFeeScheduleCleared is a free log subscription operation binding the contract event 0xe3f22e2cb60f792203c33a8b3ce8faae7df971c5bd56aa1beca573c904909aa2.
//
// Solidity: event ProjectFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) WatchProjectFeeScheduleCleared(opts *bind.WatchOpts, sink chan<- *TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared, projectId [][32]byte, sourceId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.WatchLogs(opts, "ProjectFeeScheduleCleared", projectIdRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared)
				if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "ProjectFeeScheduleCleared", log); err != nil {
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

// ParseProjectFeeScheduleCleared is a log parse operation binding the contract event 0xe3f22e2cb60f792203c33a8b3ce8faae7df971c5bd56aa1beca573c904909aa2.
//
// Solidity: event ProjectFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) ParseProjectFeeScheduleCleared(log types.Log) (*TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared, error) {
	event := new(TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared)
	if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "ProjectFeeScheduleCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsFeeScheduleManagerProjectFeeScheduleSetIterator is returned from FilterProjectFeeScheduleSet and is used to iterate over the raw logs and unpacked data for ProjectFeeScheduleSet events raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerProjectFeeScheduleSetIterator struct {
	Event *TeePaymentsFeeScheduleManagerProjectFeeScheduleSet // Event containing the contract specifics and raw log

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
func (it *TeePaymentsFeeScheduleManagerProjectFeeScheduleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsFeeScheduleManagerProjectFeeScheduleSet)
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
		it.Event = new(TeePaymentsFeeScheduleManagerProjectFeeScheduleSet)
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
func (it *TeePaymentsFeeScheduleManagerProjectFeeScheduleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsFeeScheduleManagerProjectFeeScheduleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsFeeScheduleManagerProjectFeeScheduleSet represents a ProjectFeeScheduleSet event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerProjectFeeScheduleSet struct {
	ProjectId [32]byte
	SourceId  [32]byte
	Schedule  []ITeePaymentsFeeScheduleManagerFeeSchedule
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectFeeScheduleSet is a free log retrieval operation binding the contract event 0xbf2c68fcd2f4ce03009cd21ea065e36803d64647a5840a1b97cdf4bdfc1dacfe.
//
// Solidity: event ProjectFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, (int16,uint16)[] schedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) FilterProjectFeeScheduleSet(opts *bind.FilterOpts, projectId [][32]byte, sourceId [][32]byte) (*TeePaymentsFeeScheduleManagerProjectFeeScheduleSetIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.FilterLogs(opts, "ProjectFeeScheduleSet", projectIdRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerProjectFeeScheduleSetIterator{contract: _TeePaymentsFeeScheduleManager.contract, event: "ProjectFeeScheduleSet", logs: logs, sub: sub}, nil
}

// WatchProjectFeeScheduleSet is a free log subscription operation binding the contract event 0xbf2c68fcd2f4ce03009cd21ea065e36803d64647a5840a1b97cdf4bdfc1dacfe.
//
// Solidity: event ProjectFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, (int16,uint16)[] schedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) WatchProjectFeeScheduleSet(opts *bind.WatchOpts, sink chan<- *TeePaymentsFeeScheduleManagerProjectFeeScheduleSet, projectId [][32]byte, sourceId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.WatchLogs(opts, "ProjectFeeScheduleSet", projectIdRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsFeeScheduleManagerProjectFeeScheduleSet)
				if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "ProjectFeeScheduleSet", log); err != nil {
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

// ParseProjectFeeScheduleSet is a log parse operation binding the contract event 0xbf2c68fcd2f4ce03009cd21ea065e36803d64647a5840a1b97cdf4bdfc1dacfe.
//
// Solidity: event ProjectFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, (int16,uint16)[] schedule)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) ParseProjectFeeScheduleSet(log types.Log) (*TeePaymentsFeeScheduleManagerProjectFeeScheduleSet, error) {
	event := new(TeePaymentsFeeScheduleManagerProjectFeeScheduleSet)
	if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "ProjectFeeScheduleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceledIterator struct {
	Event *TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled)
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
func (it *TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceledIterator{contract: _TeePaymentsFeeScheduleManager.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled)
				if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled, error) {
	event := new(TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled)
	if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecutedIterator struct {
	Event *TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted)
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
func (it *TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecutedIterator{contract: _TeePaymentsFeeScheduleManager.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted)
				if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted, error) {
	event := new(TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted)
	if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsFeeScheduleManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerUpgradedIterator struct {
	Event *TeePaymentsFeeScheduleManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *TeePaymentsFeeScheduleManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsFeeScheduleManagerUpgraded)
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
		it.Event = new(TeePaymentsFeeScheduleManagerUpgraded)
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
func (it *TeePaymentsFeeScheduleManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsFeeScheduleManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsFeeScheduleManagerUpgraded represents a Upgraded event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeePaymentsFeeScheduleManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFeeScheduleManagerUpgradedIterator{contract: _TeePaymentsFeeScheduleManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeePaymentsFeeScheduleManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePaymentsFeeScheduleManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsFeeScheduleManagerUpgraded)
				if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManagerFilterer) ParseUpgraded(log types.Log) (*TeePaymentsFeeScheduleManagerUpgraded, error) {
	event := new(TeePaymentsFeeScheduleManagerUpgraded)
	if err := _TeePaymentsFeeScheduleManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
