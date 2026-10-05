// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepaymentsfeeschedulemanager

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// ITeePaymentsBasePMWMultisigAccount is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsBasePMWMultisigAccount struct {
	SourceId       [32]byte
	AccountAddress string
}

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

// TeePaymentsFeeScheduleManagerMetaData contains all meta data concerning the TeePaymentsFeeScheduleManager contract.
var TeePaymentsFeeScheduleManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelay\",\"type\":\"uint256\"}],\"name\":\"DelayTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyScheduleNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"FeeScheduleConfigNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeScheduleNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeFactor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"}],\"name\":\"InvalidFeeScheduleConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProjectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SourceLimitsNotConfigured\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockCallNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySchedules\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"}],\"name\":\"AccountFeeScheduleCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeSchedule[]\",\"name\":\"schedule\",\"type\":\"tuple[]\"}],\"name\":\"AccountFeeScheduleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"sourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"FeeScheduleConfigsCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeScheduleConfigInput[]\",\"name\":\"configs\",\"type\":\"tuple[]\"}],\"name\":\"FeeScheduleConfigsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"ProjectFeeScheduleCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeSchedule[]\",\"name\":\"schedule\",\"type\":\"tuple[]\"}],\"name\":\"ProjectFeeScheduleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"clearAccountFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_sourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"clearFeeScheduleConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"clearProjectFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"contractIIFlareTeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAccountFeeSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_accountHash\",\"type\":\"bytes32\"}],\"name\":\"getEffectiveSchedule\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_feeSchedule\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getFeeScheduleConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeScheduleConfig\",\"name\":\"_config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getProjectFeeSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"name\":\"setAccountFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeScheduleConfigInput[]\",\"name\":\"_configs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeScheduleConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structITeePaymentsFeeScheduleManager.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"name\":\"setProjectFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsRegistry\",\"outputs\":[{\"internalType\":\"contractITeePaymentsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"int16[][]\",\"name\":\"_factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"_delaysSeconds\",\"type\":\"uint16[]\"}],\"name\":\"validateAndEncodeSchedules\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_encodedPerPayment\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "TeePaymentsFeeScheduleManager",
}

// TeePaymentsFeeScheduleManager is an auto generated Go binding around an Ethereum contract.
type TeePaymentsFeeScheduleManager struct {
	abi abi.ABI
}

// NewTeePaymentsFeeScheduleManager creates a new instance of TeePaymentsFeeScheduleManager.
func NewTeePaymentsFeeScheduleManager() *TeePaymentsFeeScheduleManager {
	parsed, err := TeePaymentsFeeScheduleManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TeePaymentsFeeScheduleManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TeePaymentsFeeScheduleManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16fc2f6d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackCancelGovernanceCall(encodedCall []byte) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("cancelGovernanceCall", encodedCall)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16fc2f6d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackCancelGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("cancelGovernanceCall", encodedCall)
}

// PackClearAccountFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf7443f84.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clearAccountFeeSchedule((bytes32,string) _account) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackClearAccountFeeSchedule(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("clearAccountFeeSchedule", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClearAccountFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf7443f84.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clearAccountFeeSchedule((bytes32,string) _account) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackClearAccountFeeSchedule(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("clearAccountFeeSchedule", account)
}

// PackClearFeeScheduleConfigs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11a5c047.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clearFeeScheduleConfigs(bytes32[] _sourceIds) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackClearFeeScheduleConfigs(sourceIds [][32]byte) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("clearFeeScheduleConfigs", sourceIds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClearFeeScheduleConfigs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11a5c047.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clearFeeScheduleConfigs(bytes32[] _sourceIds) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackClearFeeScheduleConfigs(sourceIds [][32]byte) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("clearFeeScheduleConfigs", sourceIds)
}

// PackClearProjectFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7aecd6f7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clearProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackClearProjectFeeSchedule(projectId [32]byte, sourceId [32]byte) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("clearProjectFeeSchedule", projectId, sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClearProjectFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7aecd6f7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clearProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackClearProjectFeeSchedule(projectId [32]byte, sourceId [32]byte) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("clearProjectFeeSchedule", projectId, sourceId)
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackExecuteGovernanceCall(encodedCall []byte) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("executeGovernanceCall", encodedCall)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackExecuteGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("executeGovernanceCall", encodedCall)
}

// PackFlareTeeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x453f7ab4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareTeeManager() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackFlareTeeManager() []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("flareTeeManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFlareTeeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x453f7ab4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function flareTeeManager() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackFlareTeeManager() ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("flareTeeManager")
}

// UnpackFlareTeeManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackFlareTeeManager(data []byte) (common.Address, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("flareTeeManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAccountFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3c9e0c3f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAccountFeeSchedule((bytes32,string) _account) view returns((int16,uint16)[] _schedule)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackGetAccountFeeSchedule(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("getAccountFeeSchedule", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAccountFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3c9e0c3f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAccountFeeSchedule((bytes32,string) _account) view returns((int16,uint16)[] _schedule)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackGetAccountFeeSchedule(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("getAccountFeeSchedule", account)
}

// UnpackGetAccountFeeSchedule is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3c9e0c3f.
//
// Solidity: function getAccountFeeSchedule((bytes32,string) _account) view returns((int16,uint16)[] _schedule)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackGetAccountFeeSchedule(data []byte) ([]ITeePaymentsFeeScheduleManagerFeeSchedule, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("getAccountFeeSchedule", data)
	if err != nil {
		return *new([]ITeePaymentsFeeScheduleManagerFeeSchedule), err
	}
	out0 := *abi.ConvertType(out[0], new([]ITeePaymentsFeeScheduleManagerFeeSchedule)).(*[]ITeePaymentsFeeScheduleManagerFeeSchedule)
	return out0, nil
}

// PackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackGetAddressUpdater() []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("getAddressUpdater")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackGetAddressUpdater() ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetEffectiveSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed8b6f51.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getEffectiveSchedule(bytes32 _projectId, bytes32 _sourceId, bytes32 _accountHash) view returns(bytes _feeSchedule)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackGetEffectiveSchedule(projectId [32]byte, sourceId [32]byte, accountHash [32]byte) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("getEffectiveSchedule", projectId, sourceId, accountHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetEffectiveSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed8b6f51.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getEffectiveSchedule(bytes32 _projectId, bytes32 _sourceId, bytes32 _accountHash) view returns(bytes _feeSchedule)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackGetEffectiveSchedule(projectId [32]byte, sourceId [32]byte, accountHash [32]byte) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("getEffectiveSchedule", projectId, sourceId, accountHash)
}

// UnpackGetEffectiveSchedule is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xed8b6f51.
//
// Solidity: function getEffectiveSchedule(bytes32 _projectId, bytes32 _sourceId, bytes32 _accountHash) view returns(bytes _feeSchedule)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackGetEffectiveSchedule(data []byte) ([]byte, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("getEffectiveSchedule", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackGetFeeScheduleConfig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16743f21.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getFeeScheduleConfig(bytes32 _sourceId) view returns((uint8,uint16) _config)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackGetFeeScheduleConfig(sourceId [32]byte) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("getFeeScheduleConfig", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetFeeScheduleConfig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16743f21.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getFeeScheduleConfig(bytes32 _sourceId) view returns((uint8,uint16) _config)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackGetFeeScheduleConfig(sourceId [32]byte) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("getFeeScheduleConfig", sourceId)
}

// UnpackGetFeeScheduleConfig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x16743f21.
//
// Solidity: function getFeeScheduleConfig(bytes32 _sourceId) view returns((uint8,uint16) _config)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackGetFeeScheduleConfig(data []byte) (ITeePaymentsFeeScheduleManagerFeeScheduleConfig, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("getFeeScheduleConfig", data)
	if err != nil {
		return *new(ITeePaymentsFeeScheduleManagerFeeScheduleConfig), err
	}
	out0 := *abi.ConvertType(out[0], new(ITeePaymentsFeeScheduleManagerFeeScheduleConfig)).(*ITeePaymentsFeeScheduleManagerFeeScheduleConfig)
	return out0, nil
}

// PackGetProjectFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x30d1bb93.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) view returns((int16,uint16)[] _schedule)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackGetProjectFeeSchedule(projectId [32]byte, sourceId [32]byte) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("getProjectFeeSchedule", projectId, sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetProjectFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x30d1bb93.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) view returns((int16,uint16)[] _schedule)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackGetProjectFeeSchedule(projectId [32]byte, sourceId [32]byte) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("getProjectFeeSchedule", projectId, sourceId)
}

// UnpackGetProjectFeeSchedule is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x30d1bb93.
//
// Solidity: function getProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) view returns((int16,uint16)[] _schedule)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackGetProjectFeeSchedule(data []byte) ([]ITeePaymentsFeeScheduleManagerFeeSchedule, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("getProjectFeeSchedule", data)
	if err != nil {
		return *new([]ITeePaymentsFeeScheduleManagerFeeSchedule), err
	}
	out0 := *abi.ConvertType(out[0], new([]ITeePaymentsFeeScheduleManagerFeeSchedule)).(*[]ITeePaymentsFeeScheduleManagerFeeSchedule)
	return out0, nil
}

// PackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governance() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackGovernance() []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("governance")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function governance() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackGovernance() ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("governance", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGovernanceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62354e03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governanceSettings() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackGovernanceSettings() []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("governanceSettings")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGovernanceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62354e03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function governanceSettings() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackGovernanceSettings() ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("governanceSettings", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c60da1b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function implementation() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackImplementation() []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("implementation")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c60da1b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function implementation() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackImplementation() ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("implementation")
}

// UnpackImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackImplementation(data []byte) (common.Address, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("implementation", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc0c53b8b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc0c53b8b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackIsExecutor(address common.Address) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("isExecutor", address)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackProductionMode() []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("productionMode")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function productionMode() view returns(bool)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackProductionMode() ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackProductionMode(data []byte) (bool, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("productionMode", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackProxiableUUID() []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackProxiableUUID() ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("proxiableUUID")
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackSetAccountFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32ab4577.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setAccountFeeSchedule((bytes32,string) _account, (int16,uint16)[] _schedule) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackSetAccountFeeSchedule(account ITeePaymentsBasePMWMultisigAccount, schedule []ITeePaymentsFeeScheduleManagerFeeSchedule) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("setAccountFeeSchedule", account, schedule)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetAccountFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32ab4577.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setAccountFeeSchedule((bytes32,string) _account, (int16,uint16)[] _schedule) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackSetAccountFeeSchedule(account ITeePaymentsBasePMWMultisigAccount, schedule []ITeePaymentsFeeScheduleManagerFeeSchedule) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("setAccountFeeSchedule", account, schedule)
}

// PackSetFeeScheduleConfigs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b7f8f29.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFeeScheduleConfigs((uint16,uint8,bytes32)[] _configs) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackSetFeeScheduleConfigs(configs []ITeePaymentsFeeScheduleManagerFeeScheduleConfigInput) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("setFeeScheduleConfigs", configs)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFeeScheduleConfigs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b7f8f29.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFeeScheduleConfigs((uint16,uint8,bytes32)[] _configs) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackSetFeeScheduleConfigs(configs []ITeePaymentsFeeScheduleManagerFeeScheduleConfigInput) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("setFeeScheduleConfigs", configs)
}

// PackSetProjectFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe27d19bd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId, (int16,uint16)[] _schedule) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackSetProjectFeeSchedule(projectId [32]byte, sourceId [32]byte, schedule []ITeePaymentsFeeScheduleManagerFeeSchedule) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("setProjectFeeSchedule", projectId, sourceId, schedule)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetProjectFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe27d19bd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId, (int16,uint16)[] _schedule) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackSetProjectFeeSchedule(projectId [32]byte, sourceId [32]byte, schedule []ITeePaymentsFeeScheduleManagerFeeSchedule) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("setProjectFeeSchedule", projectId, sourceId, schedule)
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackSwitchToProductionMode() []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("switchToProductionMode")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function switchToProductionMode() returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackSwitchToProductionMode() ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("switchToProductionMode")
}

// PackTeePaymentsRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaef828de.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackTeePaymentsRegistry() []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("teePaymentsRegistry")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTeePaymentsRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaef828de.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackTeePaymentsRegistry() ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("teePaymentsRegistry")
}

// UnpackTeePaymentsRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackTeePaymentsRegistry(data []byte) (common.Address, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("teePaymentsRegistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("upgradeToAndCall", newImplementation, data)
}

// PackValidateAndEncodeSchedules is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3d699135.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function validateAndEncodeSchedules(bytes32 _sourceId, int16[][] _factorsBIPSPerPayment, uint16[] _delaysSeconds) view returns(bytes[] _encodedPerPayment)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) PackValidateAndEncodeSchedules(sourceId [32]byte, factorsBIPSPerPayment [][]int16, delaysSeconds []uint16) []byte {
	enc, err := teePaymentsFeeScheduleManager.abi.Pack("validateAndEncodeSchedules", sourceId, factorsBIPSPerPayment, delaysSeconds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackValidateAndEncodeSchedules is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3d699135.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function validateAndEncodeSchedules(bytes32 _sourceId, int16[][] _factorsBIPSPerPayment, uint16[] _delaysSeconds) view returns(bytes[] _encodedPerPayment)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) TryPackValidateAndEncodeSchedules(sourceId [32]byte, factorsBIPSPerPayment [][]int16, delaysSeconds []uint16) ([]byte, error) {
	return teePaymentsFeeScheduleManager.abi.Pack("validateAndEncodeSchedules", sourceId, factorsBIPSPerPayment, delaysSeconds)
}

// UnpackValidateAndEncodeSchedules is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3d699135.
//
// Solidity: function validateAndEncodeSchedules(bytes32 _sourceId, int16[][] _factorsBIPSPerPayment, uint16[] _delaysSeconds) view returns(bytes[] _encodedPerPayment)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackValidateAndEncodeSchedules(data []byte) ([][]byte, error) {
	out, err := teePaymentsFeeScheduleManager.abi.Unpack("validateAndEncodeSchedules", data)
	if err != nil {
		return *new([][]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	return out0, nil
}

// TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared represents a AccountFeeScheduleCleared event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared struct {
	ProjectId      [32]byte
	SourceId       [32]byte
	AccountAddress string
	AccountHash    [32]byte
	Raw            *types.Log // Blockchain specific contextual infos
}

const TeePaymentsFeeScheduleManagerAccountFeeScheduleClearedEventName = "AccountFeeScheduleCleared"

// ContractEventName returns the user-defined event name.
func (TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared) ContractEventName() string {
	return TeePaymentsFeeScheduleManagerAccountFeeScheduleClearedEventName
}

// UnpackAccountFeeScheduleClearedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AccountFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackAccountFeeScheduleClearedEvent(log *types.Log) (*TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared, error) {
	event := "AccountFeeScheduleCleared"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsFeeScheduleManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsFeeScheduleManagerAccountFeeScheduleCleared)
	if len(log.Data) > 0 {
		if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsFeeScheduleManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TeePaymentsFeeScheduleManagerAccountFeeScheduleSet represents a AccountFeeScheduleSet event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerAccountFeeScheduleSet struct {
	ProjectId      [32]byte
	SourceId       [32]byte
	AccountAddress string
	AccountHash    [32]byte
	Schedule       []ITeePaymentsFeeScheduleManagerFeeSchedule
	Raw            *types.Log // Blockchain specific contextual infos
}

const TeePaymentsFeeScheduleManagerAccountFeeScheduleSetEventName = "AccountFeeScheduleSet"

// ContractEventName returns the user-defined event name.
func (TeePaymentsFeeScheduleManagerAccountFeeScheduleSet) ContractEventName() string {
	return TeePaymentsFeeScheduleManagerAccountFeeScheduleSetEventName
}

// UnpackAccountFeeScheduleSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AccountFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash, (int16,uint16)[] schedule)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackAccountFeeScheduleSetEvent(log *types.Log) (*TeePaymentsFeeScheduleManagerAccountFeeScheduleSet, error) {
	event := "AccountFeeScheduleSet"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsFeeScheduleManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsFeeScheduleManagerAccountFeeScheduleSet)
	if len(log.Data) > 0 {
		if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsFeeScheduleManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared represents a FeeScheduleConfigsCleared event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared struct {
	SourceIds [][32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const TeePaymentsFeeScheduleManagerFeeScheduleConfigsClearedEventName = "FeeScheduleConfigsCleared"

// ContractEventName returns the user-defined event name.
func (TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared) ContractEventName() string {
	return TeePaymentsFeeScheduleManagerFeeScheduleConfigsClearedEventName
}

// UnpackFeeScheduleConfigsClearedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeScheduleConfigsCleared(bytes32[] sourceIds)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackFeeScheduleConfigsClearedEvent(log *types.Log) (*TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared, error) {
	event := "FeeScheduleConfigsCleared"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsFeeScheduleManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsFeeScheduleManagerFeeScheduleConfigsCleared)
	if len(log.Data) > 0 {
		if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsFeeScheduleManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet represents a FeeScheduleConfigsSet event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet struct {
	Configs []ITeePaymentsFeeScheduleManagerFeeScheduleConfigInput
	Raw     *types.Log // Blockchain specific contextual infos
}

const TeePaymentsFeeScheduleManagerFeeScheduleConfigsSetEventName = "FeeScheduleConfigsSet"

// ContractEventName returns the user-defined event name.
func (TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet) ContractEventName() string {
	return TeePaymentsFeeScheduleManagerFeeScheduleConfigsSetEventName
}

// UnpackFeeScheduleConfigsSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeScheduleConfigsSet((uint16,uint8,bytes32)[] configs)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackFeeScheduleConfigsSetEvent(log *types.Log) (*TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet, error) {
	event := "FeeScheduleConfigsSet"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsFeeScheduleManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsFeeScheduleManagerFeeScheduleConfigsSet)
	if len(log.Data) > 0 {
		if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsFeeScheduleManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TeePaymentsFeeScheduleManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const TeePaymentsFeeScheduleManagerGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (TeePaymentsFeeScheduleManagerGovernanceCallTimelocked) ContractEventName() string {
	return TeePaymentsFeeScheduleManagerGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*TeePaymentsFeeScheduleManagerGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsFeeScheduleManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsFeeScheduleManagerGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsFeeScheduleManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TeePaymentsFeeScheduleManagerGovernanceInitialised represents a GovernanceInitialised event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const TeePaymentsFeeScheduleManagerGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (TeePaymentsFeeScheduleManagerGovernanceInitialised) ContractEventName() string {
	return TeePaymentsFeeScheduleManagerGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackGovernanceInitialisedEvent(log *types.Log) (*TeePaymentsFeeScheduleManagerGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsFeeScheduleManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsFeeScheduleManagerGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsFeeScheduleManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TeePaymentsFeeScheduleManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const TeePaymentsFeeScheduleManagerGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (TeePaymentsFeeScheduleManagerGovernedProductionModeEntered) ContractEventName() string {
	return TeePaymentsFeeScheduleManagerGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*TeePaymentsFeeScheduleManagerGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsFeeScheduleManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsFeeScheduleManagerGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsFeeScheduleManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TeePaymentsFeeScheduleManagerInitialized represents a Initialized event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const TeePaymentsFeeScheduleManagerInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (TeePaymentsFeeScheduleManagerInitialized) ContractEventName() string {
	return TeePaymentsFeeScheduleManagerInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackInitializedEvent(log *types.Log) (*TeePaymentsFeeScheduleManagerInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsFeeScheduleManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsFeeScheduleManagerInitialized)
	if len(log.Data) > 0 {
		if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsFeeScheduleManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared represents a ProjectFeeScheduleCleared event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared struct {
	ProjectId [32]byte
	SourceId  [32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const TeePaymentsFeeScheduleManagerProjectFeeScheduleClearedEventName = "ProjectFeeScheduleCleared"

// ContractEventName returns the user-defined event name.
func (TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared) ContractEventName() string {
	return TeePaymentsFeeScheduleManagerProjectFeeScheduleClearedEventName
}

// UnpackProjectFeeScheduleClearedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProjectFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackProjectFeeScheduleClearedEvent(log *types.Log) (*TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared, error) {
	event := "ProjectFeeScheduleCleared"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsFeeScheduleManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsFeeScheduleManagerProjectFeeScheduleCleared)
	if len(log.Data) > 0 {
		if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsFeeScheduleManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TeePaymentsFeeScheduleManagerProjectFeeScheduleSet represents a ProjectFeeScheduleSet event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerProjectFeeScheduleSet struct {
	ProjectId [32]byte
	SourceId  [32]byte
	Schedule  []ITeePaymentsFeeScheduleManagerFeeSchedule
	Raw       *types.Log // Blockchain specific contextual infos
}

const TeePaymentsFeeScheduleManagerProjectFeeScheduleSetEventName = "ProjectFeeScheduleSet"

// ContractEventName returns the user-defined event name.
func (TeePaymentsFeeScheduleManagerProjectFeeScheduleSet) ContractEventName() string {
	return TeePaymentsFeeScheduleManagerProjectFeeScheduleSetEventName
}

// UnpackProjectFeeScheduleSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProjectFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, (int16,uint16)[] schedule)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackProjectFeeScheduleSetEvent(log *types.Log) (*TeePaymentsFeeScheduleManagerProjectFeeScheduleSet, error) {
	event := "ProjectFeeScheduleSet"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsFeeScheduleManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsFeeScheduleManagerProjectFeeScheduleSet)
	if len(log.Data) > 0 {
		if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsFeeScheduleManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled) ContractEventName() string {
	return TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsFeeScheduleManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsFeeScheduleManagerTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsFeeScheduleManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted) ContractEventName() string {
	return TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsFeeScheduleManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsFeeScheduleManagerTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsFeeScheduleManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// TeePaymentsFeeScheduleManagerUpgraded represents a Upgraded event raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const TeePaymentsFeeScheduleManagerUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (TeePaymentsFeeScheduleManagerUpgraded) ContractEventName() string {
	return TeePaymentsFeeScheduleManagerUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackUpgradedEvent(log *types.Log) (*TeePaymentsFeeScheduleManagerUpgraded, error) {
	event := "Upgraded"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsFeeScheduleManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsFeeScheduleManagerUpgraded)
	if len(log.Data) > 0 {
		if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsFeeScheduleManager.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["AccountNotRegistered"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackAccountNotRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["AlreadyInProductionMode"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackAlreadyInProductionModeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["DelayTooLarge"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackDelayTooLargeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["EmptyScheduleNotAllowed"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackEmptyScheduleNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["FeeScheduleConfigNotSet"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackFeeScheduleConfigNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["FeeScheduleNotSet"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackFeeScheduleNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["GovernedAddressZero"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackGovernedAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["GovernedAlreadyInitialized"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackGovernedAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["InvalidFeeDelay"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackInvalidFeeDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["InvalidFeeFactor"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackInvalidFeeFactorError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["InvalidFeeScheduleConfig"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackInvalidFeeScheduleConfigError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["OnlyExecutor"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["OnlyProjectOwner"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackOnlyProjectOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["OnlySystemExtensionId"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackOnlySystemExtensionIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["SourceLimitsNotConfigured"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackSourceLimitsNotConfiguredError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["TimelockCallNotFound"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackTimelockCallNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["TimelockNotAllowedYet"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackTimelockNotAllowedYetError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["TooManySchedules"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackTooManySchedulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsFeeScheduleManager.abi.Errors["UnsupportedSourceId"].ID.Bytes()[:4]) {
		return teePaymentsFeeScheduleManager.UnpackUnsupportedSourceIdError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// TeePaymentsFeeScheduleManagerAccountNotRegistered represents a AccountNotRegistered error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerAccountNotRegistered struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccountNotRegistered()
func TeePaymentsFeeScheduleManagerAccountNotRegisteredErrorID() common.Hash {
	return common.HexToHash("0x9704dd5103977e212a4ccc5d11b13fb9a4e452fb5b457b2b39c0fcad36ab2cf0")
}

// UnpackAccountNotRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccountNotRegistered()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackAccountNotRegisteredError(raw []byte) (*TeePaymentsFeeScheduleManagerAccountNotRegistered, error) {
	out := new(TeePaymentsFeeScheduleManagerAccountNotRegistered)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "AccountNotRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerAddressEmptyCode represents a AddressEmptyCode error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func TeePaymentsFeeScheduleManagerAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackAddressEmptyCodeError(raw []byte) (*TeePaymentsFeeScheduleManagerAddressEmptyCode, error) {
	out := new(TeePaymentsFeeScheduleManagerAddressEmptyCode)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerAlreadyInProductionMode represents a AlreadyInProductionMode error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerAlreadyInProductionMode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInProductionMode()
func TeePaymentsFeeScheduleManagerAlreadyInProductionModeErrorID() common.Hash {
	return common.HexToHash("0x72a23ec2b1fdbc3de7d5f411b572fcf8ad8409803196930f63ddf40a2a2b2594")
}

// UnpackAlreadyInProductionModeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInProductionMode()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackAlreadyInProductionModeError(raw []byte) (*TeePaymentsFeeScheduleManagerAlreadyInProductionMode, error) {
	out := new(TeePaymentsFeeScheduleManagerAlreadyInProductionMode)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "AlreadyInProductionMode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerDelayTooLarge represents a DelayTooLarge error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerDelayTooLarge struct {
	Delay    *big.Int
	MaxDelay *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DelayTooLarge(uint256 delay, uint256 maxDelay)
func TeePaymentsFeeScheduleManagerDelayTooLargeErrorID() common.Hash {
	return common.HexToHash("0xf77cfffc23bacd92259f81786482f236072e2ca49bedf646dcdb1d309c5f0ff4")
}

// UnpackDelayTooLargeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DelayTooLarge(uint256 delay, uint256 maxDelay)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackDelayTooLargeError(raw []byte) (*TeePaymentsFeeScheduleManagerDelayTooLarge, error) {
	out := new(TeePaymentsFeeScheduleManagerDelayTooLarge)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "DelayTooLarge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func TeePaymentsFeeScheduleManagerERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackERC1967InvalidImplementationError(raw []byte) (*TeePaymentsFeeScheduleManagerERC1967InvalidImplementation, error) {
	out := new(TeePaymentsFeeScheduleManagerERC1967InvalidImplementation)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerERC1967NonPayable represents a ERC1967NonPayable error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func TeePaymentsFeeScheduleManagerERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackERC1967NonPayableError(raw []byte) (*TeePaymentsFeeScheduleManagerERC1967NonPayable, error) {
	out := new(TeePaymentsFeeScheduleManagerERC1967NonPayable)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerEmptyScheduleNotAllowed represents a EmptyScheduleNotAllowed error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerEmptyScheduleNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmptyScheduleNotAllowed()
func TeePaymentsFeeScheduleManagerEmptyScheduleNotAllowedErrorID() common.Hash {
	return common.HexToHash("0xcdc0f5e297e9ba099b219e61725f34d2cf720b212c9e7d68571ab793321ba39d")
}

// UnpackEmptyScheduleNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmptyScheduleNotAllowed()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackEmptyScheduleNotAllowedError(raw []byte) (*TeePaymentsFeeScheduleManagerEmptyScheduleNotAllowed, error) {
	out := new(TeePaymentsFeeScheduleManagerEmptyScheduleNotAllowed)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "EmptyScheduleNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerFailedCall represents a FailedCall error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func TeePaymentsFeeScheduleManagerFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackFailedCallError(raw []byte) (*TeePaymentsFeeScheduleManagerFailedCall, error) {
	out := new(TeePaymentsFeeScheduleManagerFailedCall)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerFeeScheduleConfigNotSet represents a FeeScheduleConfigNotSet error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerFeeScheduleConfigNotSet struct {
	SourceId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeScheduleConfigNotSet(bytes32 sourceId)
func TeePaymentsFeeScheduleManagerFeeScheduleConfigNotSetErrorID() common.Hash {
	return common.HexToHash("0x42b58a254883973ca2465c0acbab0b789fd25da4ab5f1bea80bbfd12320073d5")
}

// UnpackFeeScheduleConfigNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeScheduleConfigNotSet(bytes32 sourceId)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackFeeScheduleConfigNotSetError(raw []byte) (*TeePaymentsFeeScheduleManagerFeeScheduleConfigNotSet, error) {
	out := new(TeePaymentsFeeScheduleManagerFeeScheduleConfigNotSet)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "FeeScheduleConfigNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerFeeScheduleNotSet represents a FeeScheduleNotSet error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerFeeScheduleNotSet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeScheduleNotSet()
func TeePaymentsFeeScheduleManagerFeeScheduleNotSetErrorID() common.Hash {
	return common.HexToHash("0x4f47ce423907d3f49d5eb099e6293bf6be71471950c89ba028907cabdaae2e73")
}

// UnpackFeeScheduleNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeScheduleNotSet()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackFeeScheduleNotSetError(raw []byte) (*TeePaymentsFeeScheduleManagerFeeScheduleNotSet, error) {
	out := new(TeePaymentsFeeScheduleManagerFeeScheduleNotSet)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "FeeScheduleNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerGovernedAddressZero represents a GovernedAddressZero error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerGovernedAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAddressZero()
func TeePaymentsFeeScheduleManagerGovernedAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x695725f6b429bac0cc5a2fa88214e31520e0c69f2fb0e2fb881a7130a564f49f")
}

// UnpackGovernedAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAddressZero()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackGovernedAddressZeroError(raw []byte) (*TeePaymentsFeeScheduleManagerGovernedAddressZero, error) {
	out := new(TeePaymentsFeeScheduleManagerGovernedAddressZero)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "GovernedAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerGovernedAlreadyInitialized represents a GovernedAlreadyInitialized error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerGovernedAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAlreadyInitialized()
func TeePaymentsFeeScheduleManagerGovernedAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0xb2f0162156d83828197776d56bbfba0a7a9d5936fad70f0a3014ae18cc7aa33c")
}

// UnpackGovernedAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAlreadyInitialized()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackGovernedAlreadyInitializedError(raw []byte) (*TeePaymentsFeeScheduleManagerGovernedAlreadyInitialized, error) {
	out := new(TeePaymentsFeeScheduleManagerGovernedAlreadyInitialized)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "GovernedAlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerInvalidFeeDelay represents a InvalidFeeDelay error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerInvalidFeeDelay struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeDelay(uint256 index)
func TeePaymentsFeeScheduleManagerInvalidFeeDelayErrorID() common.Hash {
	return common.HexToHash("0x2a9c9c04432dee2444e275db8772f3255e5685176b3891e4a344c3c7c4854620")
}

// UnpackInvalidFeeDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeDelay(uint256 index)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackInvalidFeeDelayError(raw []byte) (*TeePaymentsFeeScheduleManagerInvalidFeeDelay, error) {
	out := new(TeePaymentsFeeScheduleManagerInvalidFeeDelay)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "InvalidFeeDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerInvalidFeeFactor represents a InvalidFeeFactor error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerInvalidFeeFactor struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeFactor(uint256 index)
func TeePaymentsFeeScheduleManagerInvalidFeeFactorErrorID() common.Hash {
	return common.HexToHash("0x89619400c26b8355799d991b2b4ac022351c0c4f21a4c0f3b4ee2ef70d0c8737")
}

// UnpackInvalidFeeFactorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeFactor(uint256 index)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackInvalidFeeFactorError(raw []byte) (*TeePaymentsFeeScheduleManagerInvalidFeeFactor, error) {
	out := new(TeePaymentsFeeScheduleManagerInvalidFeeFactor)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "InvalidFeeFactor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerInvalidFeeScheduleConfig represents a InvalidFeeScheduleConfig error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerInvalidFeeScheduleConfig struct {
	SourceId        [32]byte
	MaxSchedules    uint8
	MaxDelaySeconds uint16
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeScheduleConfig(bytes32 sourceId, uint8 maxSchedules, uint16 maxDelaySeconds)
func TeePaymentsFeeScheduleManagerInvalidFeeScheduleConfigErrorID() common.Hash {
	return common.HexToHash("0x3749abf55999564327d44208b64402fca0d55fb8e7c8c042042af79acd07d579")
}

// UnpackInvalidFeeScheduleConfigError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeScheduleConfig(bytes32 sourceId, uint8 maxSchedules, uint16 maxDelaySeconds)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackInvalidFeeScheduleConfigError(raw []byte) (*TeePaymentsFeeScheduleManagerInvalidFeeScheduleConfig, error) {
	out := new(TeePaymentsFeeScheduleManagerInvalidFeeScheduleConfig)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "InvalidFeeScheduleConfig", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerInvalidInitialization represents a InvalidInitialization error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func TeePaymentsFeeScheduleManagerInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackInvalidInitializationError(raw []byte) (*TeePaymentsFeeScheduleManagerInvalidInitialization, error) {
	out := new(TeePaymentsFeeScheduleManagerInvalidInitialization)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerLengthsMismatch represents a LengthsMismatch error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func TeePaymentsFeeScheduleManagerLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackLengthsMismatchError(raw []byte) (*TeePaymentsFeeScheduleManagerLengthsMismatch, error) {
	out := new(TeePaymentsFeeScheduleManagerLengthsMismatch)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerNotInitializing represents a NotInitializing error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func TeePaymentsFeeScheduleManagerNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackNotInitializingError(raw []byte) (*TeePaymentsFeeScheduleManagerNotInitializing, error) {
	out := new(TeePaymentsFeeScheduleManagerNotInitializing)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerOnlyExecutor represents a OnlyExecutor error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerOnlyExecutor struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExecutor()
func TeePaymentsFeeScheduleManagerOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x7fb6be02aaf40e420b0e7eb67e43a3feee949e8932a3c64e33b591a436670841")
}

// UnpackOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExecutor()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackOnlyExecutorError(raw []byte) (*TeePaymentsFeeScheduleManagerOnlyExecutor, error) {
	out := new(TeePaymentsFeeScheduleManagerOnlyExecutor)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "OnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerOnlyGovernance represents a OnlyGovernance error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func TeePaymentsFeeScheduleManagerOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackOnlyGovernanceError(raw []byte) (*TeePaymentsFeeScheduleManagerOnlyGovernance, error) {
	out := new(TeePaymentsFeeScheduleManagerOnlyGovernance)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerOnlyProjectOwner represents a OnlyProjectOwner error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerOnlyProjectOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProjectOwner()
func TeePaymentsFeeScheduleManagerOnlyProjectOwnerErrorID() common.Hash {
	return common.HexToHash("0x9e140988363c55adef362d3d92091fcd2495ebb514cc5070cdc0816e6a555dd6")
}

// UnpackOnlyProjectOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProjectOwner()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackOnlyProjectOwnerError(raw []byte) (*TeePaymentsFeeScheduleManagerOnlyProjectOwner, error) {
	out := new(TeePaymentsFeeScheduleManagerOnlyProjectOwner)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "OnlyProjectOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerOnlySystemExtensionId represents a OnlySystemExtensionId error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerOnlySystemExtensionId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemExtensionId()
func TeePaymentsFeeScheduleManagerOnlySystemExtensionIdErrorID() common.Hash {
	return common.HexToHash("0xbd49094d5a86ee3f3c4b39274f5f98391dde2cd16618a917e121333ac3d43ee4")
}

// UnpackOnlySystemExtensionIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemExtensionId()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackOnlySystemExtensionIdError(raw []byte) (*TeePaymentsFeeScheduleManagerOnlySystemExtensionId, error) {
	out := new(TeePaymentsFeeScheduleManagerOnlySystemExtensionId)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "OnlySystemExtensionId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerSourceLimitsNotConfigured represents a SourceLimitsNotConfigured error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerSourceLimitsNotConfigured struct {
	SourceId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceLimitsNotConfigured(bytes32 sourceId)
func TeePaymentsFeeScheduleManagerSourceLimitsNotConfiguredErrorID() common.Hash {
	return common.HexToHash("0x09198876d604df9387eeb94a3b2a10ecb3233765b8e0c28b96dd6be17e7bbe75")
}

// UnpackSourceLimitsNotConfiguredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceLimitsNotConfigured(bytes32 sourceId)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackSourceLimitsNotConfiguredError(raw []byte) (*TeePaymentsFeeScheduleManagerSourceLimitsNotConfigured, error) {
	out := new(TeePaymentsFeeScheduleManagerSourceLimitsNotConfigured)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "SourceLimitsNotConfigured", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerTimelockCallNotFound represents a TimelockCallNotFound error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerTimelockCallNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockCallNotFound()
func TeePaymentsFeeScheduleManagerTimelockCallNotFoundErrorID() common.Hash {
	return common.HexToHash("0x262ac987dd09120380d393a0c52cdef1ca58949d1107dc26d8da9f7b527eb33d")
}

// UnpackTimelockCallNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockCallNotFound()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackTimelockCallNotFoundError(raw []byte) (*TeePaymentsFeeScheduleManagerTimelockCallNotFound, error) {
	out := new(TeePaymentsFeeScheduleManagerTimelockCallNotFound)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "TimelockCallNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerTimelockNotAllowedYet represents a TimelockNotAllowedYet error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerTimelockNotAllowedYet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockNotAllowedYet()
func TeePaymentsFeeScheduleManagerTimelockNotAllowedYetErrorID() common.Hash {
	return common.HexToHash("0x6124e5c2c545bec1d02048c68500eebdabeae0cc69995a629ebccf8a62f69083")
}

// UnpackTimelockNotAllowedYetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockNotAllowedYet()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackTimelockNotAllowedYetError(raw []byte) (*TeePaymentsFeeScheduleManagerTimelockNotAllowedYet, error) {
	out := new(TeePaymentsFeeScheduleManagerTimelockNotAllowedYet)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "TimelockNotAllowedYet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerTooManySchedules represents a TooManySchedules error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerTooManySchedules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TooManySchedules()
func TeePaymentsFeeScheduleManagerTooManySchedulesErrorID() common.Hash {
	return common.HexToHash("0x6dd8671e1b420f7157257ab2728fbfd77b6ef516af0a461b4affaaee232305eb")
}

// UnpackTooManySchedulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TooManySchedules()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackTooManySchedulesError(raw []byte) (*TeePaymentsFeeScheduleManagerTooManySchedules, error) {
	out := new(TeePaymentsFeeScheduleManagerTooManySchedules)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "TooManySchedules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func TeePaymentsFeeScheduleManagerUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*TeePaymentsFeeScheduleManagerUUPSUnauthorizedCallContext, error) {
	out := new(TeePaymentsFeeScheduleManagerUUPSUnauthorizedCallContext)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func TeePaymentsFeeScheduleManagerUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*TeePaymentsFeeScheduleManagerUUPSUnsupportedProxiableUUID, error) {
	out := new(TeePaymentsFeeScheduleManagerUUPSUnsupportedProxiableUUID)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFeeScheduleManagerUnsupportedSourceId represents a UnsupportedSourceId error raised by the TeePaymentsFeeScheduleManager contract.
type TeePaymentsFeeScheduleManagerUnsupportedSourceId struct {
	SourceId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedSourceId(bytes32 sourceId)
func TeePaymentsFeeScheduleManagerUnsupportedSourceIdErrorID() common.Hash {
	return common.HexToHash("0xe9a2a60c442e03fe02a63f743afaa6a578a16c5c1cfe959ea4878b57953601ff")
}

// UnpackUnsupportedSourceIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedSourceId(bytes32 sourceId)
func (teePaymentsFeeScheduleManager *TeePaymentsFeeScheduleManager) UnpackUnsupportedSourceIdError(raw []byte) (*TeePaymentsFeeScheduleManagerUnsupportedSourceId, error) {
	out := new(TeePaymentsFeeScheduleManagerUnsupportedSourceId)
	if err := teePaymentsFeeScheduleManager.abi.UnpackIntoInterface(out, "UnsupportedSourceId", raw); err != nil {
		return nil, err
	}
	return out, nil
}
