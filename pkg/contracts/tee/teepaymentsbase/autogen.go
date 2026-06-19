// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepaymentsbase

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

// ITeePaymentsBasePMWMultisigAccount is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsBasePMWMultisigAccount struct {
	SourceId       [32]byte
	AccountAddress string
}

// ITeePaymentsBasePaymentInstruction is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsBasePaymentInstruction struct {
	RecipientAddress string
	TokenId          []byte
	Amount           *big.Int
	MaxFee           *big.Int
	PaymentReference [32]byte
}

// ITeePaymentsBaseReissueFeeParams is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsBaseReissueFeeParams struct {
	MaxFeePerPayment      []*big.Int
	FactorsBIPSPerPayment [][]int16
	DelaysSeconds         []uint16
}

// TeePaymentsBaseMetaData contains all meta data concerning the TeePaymentsBase contract.
var TeePaymentsBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuthorizationAddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuthorizationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMWMultisigAccountAddressAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMWMultisigAccountNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockInvalidSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongKeyType\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"contractIIFlareTeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAuthorizationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"}],\"name\":\"getPaymentFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAccounts\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getWalletId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePaymentsBase.PaymentInstruction\",\"name\":\"_paymentInstruction\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentModel\",\"outputs\":[{\"internalType\":\"enumPaymentModel\",\"name\":\"_paymentModel\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_batchPaymentId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePaymentsBase.PaymentInstruction[]\",\"name\":\"_paymentInstructions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"maxFeePerPayment\",\"type\":\"uint256[]\"},{\"internalType\":\"int16[][]\",\"name\":\"factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"delaysSeconds\",\"type\":\"uint16[]\"}],\"internalType\":\"structITeePaymentsBase.ReissueFeeParams\",\"name\":\"_reissueFeeParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"reissue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_finalized\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsConfigVerifier\",\"outputs\":[{\"internalType\":\"contractITeePaymentsConfigVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsFeeScheduleManager\",\"outputs\":[{\"internalType\":\"contractITeePaymentsFeeScheduleManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsRegistry\",\"outputs\":[{\"internalType\":\"contractITeePaymentsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeePaymentsBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use TeePaymentsBaseMetaData.ABI instead.
var TeePaymentsBaseABI = TeePaymentsBaseMetaData.ABI

// TeePaymentsBase is an auto generated Go binding around an Ethereum contract.
type TeePaymentsBase struct {
	TeePaymentsBaseCaller     // Read-only binding to the contract
	TeePaymentsBaseTransactor // Write-only binding to the contract
	TeePaymentsBaseFilterer   // Log filterer for contract events
}

// TeePaymentsBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeePaymentsBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeePaymentsBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeePaymentsBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeePaymentsBaseSession struct {
	Contract     *TeePaymentsBase  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeePaymentsBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeePaymentsBaseCallerSession struct {
	Contract *TeePaymentsBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TeePaymentsBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeePaymentsBaseTransactorSession struct {
	Contract     *TeePaymentsBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TeePaymentsBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeePaymentsBaseRaw struct {
	Contract *TeePaymentsBase // Generic contract binding to access the raw methods on
}

// TeePaymentsBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeePaymentsBaseCallerRaw struct {
	Contract *TeePaymentsBaseCaller // Generic read-only contract binding to access the raw methods on
}

// TeePaymentsBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeePaymentsBaseTransactorRaw struct {
	Contract *TeePaymentsBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeePaymentsBase creates a new instance of TeePaymentsBase, bound to a specific deployed contract.
func NewTeePaymentsBase(address common.Address, backend bind.ContractBackend) (*TeePaymentsBase, error) {
	contract, err := bindTeePaymentsBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsBase{TeePaymentsBaseCaller: TeePaymentsBaseCaller{contract: contract}, TeePaymentsBaseTransactor: TeePaymentsBaseTransactor{contract: contract}, TeePaymentsBaseFilterer: TeePaymentsBaseFilterer{contract: contract}}, nil
}

// NewTeePaymentsBaseCaller creates a new read-only instance of TeePaymentsBase, bound to a specific deployed contract.
func NewTeePaymentsBaseCaller(address common.Address, caller bind.ContractCaller) (*TeePaymentsBaseCaller, error) {
	contract, err := bindTeePaymentsBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsBaseCaller{contract: contract}, nil
}

// NewTeePaymentsBaseTransactor creates a new write-only instance of TeePaymentsBase, bound to a specific deployed contract.
func NewTeePaymentsBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*TeePaymentsBaseTransactor, error) {
	contract, err := bindTeePaymentsBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsBaseTransactor{contract: contract}, nil
}

// NewTeePaymentsBaseFilterer creates a new log filterer instance of TeePaymentsBase, bound to a specific deployed contract.
func NewTeePaymentsBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*TeePaymentsBaseFilterer, error) {
	contract, err := bindTeePaymentsBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsBaseFilterer{contract: contract}, nil
}

// bindTeePaymentsBase binds a generic wrapper to an already deployed contract.
func bindTeePaymentsBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeePaymentsBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePaymentsBase *TeePaymentsBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePaymentsBase.Contract.TeePaymentsBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePaymentsBase *TeePaymentsBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.TeePaymentsBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePaymentsBase *TeePaymentsBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.TeePaymentsBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePaymentsBase *TeePaymentsBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePaymentsBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePaymentsBase *TeePaymentsBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePaymentsBase *TeePaymentsBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsBase *TeePaymentsBaseCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsBase *TeePaymentsBaseSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePaymentsBase.Contract.UPGRADEINTERFACEVERSION(&_TeePaymentsBase.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePaymentsBase.Contract.UPGRADEINTERFACEVERSION(&_TeePaymentsBase.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseSession) FlareSystemsManager() (common.Address, error) {
	return _TeePaymentsBase.Contract.FlareSystemsManager(&_TeePaymentsBase.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) FlareSystemsManager() (common.Address, error) {
	return _TeePaymentsBase.Contract.FlareSystemsManager(&_TeePaymentsBase.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCaller) FlareTeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "flareTeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseSession) FlareTeeManager() (common.Address, error) {
	return _TeePaymentsBase.Contract.FlareTeeManager(&_TeePaymentsBase.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) FlareTeeManager() (common.Address, error) {
	return _TeePaymentsBase.Contract.FlareTeeManager(&_TeePaymentsBase.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsBase *TeePaymentsBaseCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsBase *TeePaymentsBaseSession) GetAddressUpdater() (common.Address, error) {
	return _TeePaymentsBase.Contract.GetAddressUpdater(&_TeePaymentsBase.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeePaymentsBase.Contract.GetAddressUpdater(&_TeePaymentsBase.CallOpts)
}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (_TeePaymentsBase *TeePaymentsBaseCaller) GetAuthorizationAddress(opts *bind.CallOpts, _account ITeePaymentsBasePMWMultisigAccount) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "getAuthorizationAddress", _account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (_TeePaymentsBase *TeePaymentsBaseSession) GetAuthorizationAddress(_account ITeePaymentsBasePMWMultisigAccount) (common.Address, error) {
	return _TeePaymentsBase.Contract.GetAuthorizationAddress(&_TeePaymentsBase.CallOpts, _account)
}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) GetAuthorizationAddress(_account ITeePaymentsBasePMWMultisigAccount) (common.Address, error) {
	return _TeePaymentsBase.Contract.GetAuthorizationAddress(&_TeePaymentsBase.CallOpts, _account)
}

// GetPaymentFee is a free data retrieval call binding the contract method 0x57abf78b.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (_TeePaymentsBase *TeePaymentsBaseCaller) GetPaymentFee(opts *bind.CallOpts, _account ITeePaymentsBasePMWMultisigAccount, _opCommand [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "getPaymentFee", _account, _opCommand)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPaymentFee is a free data retrieval call binding the contract method 0x57abf78b.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (_TeePaymentsBase *TeePaymentsBaseSession) GetPaymentFee(_account ITeePaymentsBasePMWMultisigAccount, _opCommand [32]byte) (*big.Int, error) {
	return _TeePaymentsBase.Contract.GetPaymentFee(&_TeePaymentsBase.CallOpts, _account, _opCommand)
}

// GetPaymentFee is a free data retrieval call binding the contract method 0x57abf78b.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) GetPaymentFee(_account ITeePaymentsBasePMWMultisigAccount, _opCommand [32]byte) (*big.Int, error) {
	return _TeePaymentsBase.Contract.GetPaymentFee(&_TeePaymentsBase.CallOpts, _account, _opCommand)
}

// GetWalletAccounts is a free data retrieval call binding the contract method 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (_TeePaymentsBase *TeePaymentsBaseCaller) GetWalletAccounts(opts *bind.CallOpts, _walletId [32]byte) ([]ITeePaymentsBasePMWMultisigAccount, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "getWalletAccounts", _walletId)

	if err != nil {
		return *new([]ITeePaymentsBasePMWMultisigAccount), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITeePaymentsBasePMWMultisigAccount)).(*[]ITeePaymentsBasePMWMultisigAccount)

	return out0, err

}

// GetWalletAccounts is a free data retrieval call binding the contract method 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (_TeePaymentsBase *TeePaymentsBaseSession) GetWalletAccounts(_walletId [32]byte) ([]ITeePaymentsBasePMWMultisigAccount, error) {
	return _TeePaymentsBase.Contract.GetWalletAccounts(&_TeePaymentsBase.CallOpts, _walletId)
}

// GetWalletAccounts is a free data retrieval call binding the contract method 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) GetWalletAccounts(_walletId [32]byte) ([]ITeePaymentsBasePMWMultisigAccount, error) {
	return _TeePaymentsBase.Contract.GetWalletAccounts(&_TeePaymentsBase.CallOpts, _walletId)
}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (_TeePaymentsBase *TeePaymentsBaseCaller) GetWalletId(opts *bind.CallOpts, _account ITeePaymentsBasePMWMultisigAccount) ([32]byte, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "getWalletId", _account)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (_TeePaymentsBase *TeePaymentsBaseSession) GetWalletId(_account ITeePaymentsBasePMWMultisigAccount) ([32]byte, error) {
	return _TeePaymentsBase.Contract.GetWalletId(&_TeePaymentsBase.CallOpts, _account)
}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) GetWalletId(_account ITeePaymentsBasePMWMultisigAccount) ([32]byte, error) {
	return _TeePaymentsBase.Contract.GetWalletId(&_TeePaymentsBase.CallOpts, _account)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseSession) Governance() (common.Address, error) {
	return _TeePaymentsBase.Contract.Governance(&_TeePaymentsBase.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) Governance() (common.Address, error) {
	return _TeePaymentsBase.Contract.Governance(&_TeePaymentsBase.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseSession) GovernanceSettings() (common.Address, error) {
	return _TeePaymentsBase.Contract.GovernanceSettings(&_TeePaymentsBase.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeePaymentsBase.Contract.GovernanceSettings(&_TeePaymentsBase.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseSession) Implementation() (common.Address, error) {
	return _TeePaymentsBase.Contract.Implementation(&_TeePaymentsBase.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) Implementation() (common.Address, error) {
	return _TeePaymentsBase.Contract.Implementation(&_TeePaymentsBase.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsBase *TeePaymentsBaseCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsBase *TeePaymentsBaseSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePaymentsBase.Contract.IsExecutor(&_TeePaymentsBase.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePaymentsBase.Contract.IsExecutor(&_TeePaymentsBase.CallOpts, _address)
}

// PaymentModel is a free data retrieval call binding the contract method 0xbb9d3dbb.
//
// Solidity: function paymentModel() pure returns(uint8 _paymentModel)
func (_TeePaymentsBase *TeePaymentsBaseCaller) PaymentModel(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "paymentModel")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PaymentModel is a free data retrieval call binding the contract method 0xbb9d3dbb.
//
// Solidity: function paymentModel() pure returns(uint8 _paymentModel)
func (_TeePaymentsBase *TeePaymentsBaseSession) PaymentModel() (uint8, error) {
	return _TeePaymentsBase.Contract.PaymentModel(&_TeePaymentsBase.CallOpts)
}

// PaymentModel is a free data retrieval call binding the contract method 0xbb9d3dbb.
//
// Solidity: function paymentModel() pure returns(uint8 _paymentModel)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) PaymentModel() (uint8, error) {
	return _TeePaymentsBase.Contract.PaymentModel(&_TeePaymentsBase.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsBase *TeePaymentsBaseCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsBase *TeePaymentsBaseSession) ProductionMode() (bool, error) {
	return _TeePaymentsBase.Contract.ProductionMode(&_TeePaymentsBase.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) ProductionMode() (bool, error) {
	return _TeePaymentsBase.Contract.ProductionMode(&_TeePaymentsBase.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsBase *TeePaymentsBaseCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsBase *TeePaymentsBaseSession) ProxiableUUID() ([32]byte, error) {
	return _TeePaymentsBase.Contract.ProxiableUUID(&_TeePaymentsBase.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeePaymentsBase.Contract.ProxiableUUID(&_TeePaymentsBase.CallOpts)
}

// TeePaymentsConfigVerifier is a free data retrieval call binding the contract method 0xf71c6c75.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCaller) TeePaymentsConfigVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "teePaymentsConfigVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeePaymentsConfigVerifier is a free data retrieval call binding the contract method 0xf71c6c75.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseSession) TeePaymentsConfigVerifier() (common.Address, error) {
	return _TeePaymentsBase.Contract.TeePaymentsConfigVerifier(&_TeePaymentsBase.CallOpts)
}

// TeePaymentsConfigVerifier is a free data retrieval call binding the contract method 0xf71c6c75.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) TeePaymentsConfigVerifier() (common.Address, error) {
	return _TeePaymentsBase.Contract.TeePaymentsConfigVerifier(&_TeePaymentsBase.CallOpts)
}

// TeePaymentsFeeScheduleManager is a free data retrieval call binding the contract method 0xfe9e33e8.
//
// Solidity: function teePaymentsFeeScheduleManager() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCaller) TeePaymentsFeeScheduleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "teePaymentsFeeScheduleManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeePaymentsFeeScheduleManager is a free data retrieval call binding the contract method 0xfe9e33e8.
//
// Solidity: function teePaymentsFeeScheduleManager() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseSession) TeePaymentsFeeScheduleManager() (common.Address, error) {
	return _TeePaymentsBase.Contract.TeePaymentsFeeScheduleManager(&_TeePaymentsBase.CallOpts)
}

// TeePaymentsFeeScheduleManager is a free data retrieval call binding the contract method 0xfe9e33e8.
//
// Solidity: function teePaymentsFeeScheduleManager() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) TeePaymentsFeeScheduleManager() (common.Address, error) {
	return _TeePaymentsBase.Contract.TeePaymentsFeeScheduleManager(&_TeePaymentsBase.CallOpts)
}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCaller) TeePaymentsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsBase.contract.Call(opts, &out, "teePaymentsRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseSession) TeePaymentsRegistry() (common.Address, error) {
	return _TeePaymentsBase.Contract.TeePaymentsRegistry(&_TeePaymentsBase.CallOpts)
}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsBase *TeePaymentsBaseCallerSession) TeePaymentsRegistry() (common.Address, error) {
	return _TeePaymentsBase.Contract.TeePaymentsRegistry(&_TeePaymentsBase.CallOpts)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsBase *TeePaymentsBaseTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsBase.contract.Transact(opts, "cancelGovernanceCall", _encodedCall)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsBase *TeePaymentsBaseSession) CancelGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.CancelGovernanceCall(&_TeePaymentsBase.TransactOpts, _encodedCall)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsBase *TeePaymentsBaseTransactorSession) CancelGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.CancelGovernanceCall(&_TeePaymentsBase.TransactOpts, _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsBase *TeePaymentsBaseTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsBase.contract.Transact(opts, "executeGovernanceCall", _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsBase *TeePaymentsBaseSession) ExecuteGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.ExecuteGovernanceCall(&_TeePaymentsBase.TransactOpts, _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsBase *TeePaymentsBaseTransactorSession) ExecuteGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.ExecuteGovernanceCall(&_TeePaymentsBase.TransactOpts, _encodedCall)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsBase *TeePaymentsBaseTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsBase.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsBase *TeePaymentsBaseSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.Initialize(&_TeePaymentsBase.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsBase *TeePaymentsBaseTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.Initialize(&_TeePaymentsBase.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Pay is a paid mutator transaction binding the contract method 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (_TeePaymentsBase *TeePaymentsBaseTransactor) Pay(opts *bind.TransactOpts, _account ITeePaymentsBasePMWMultisigAccount, _paymentInstruction ITeePaymentsBasePaymentInstruction, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsBase.contract.Transact(opts, "pay", _account, _paymentInstruction, _claimBackAddress)
}

// Pay is a paid mutator transaction binding the contract method 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (_TeePaymentsBase *TeePaymentsBaseSession) Pay(_account ITeePaymentsBasePMWMultisigAccount, _paymentInstruction ITeePaymentsBasePaymentInstruction, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.Pay(&_TeePaymentsBase.TransactOpts, _account, _paymentInstruction, _claimBackAddress)
}

// Pay is a paid mutator transaction binding the contract method 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (_TeePaymentsBase *TeePaymentsBaseTransactorSession) Pay(_account ITeePaymentsBasePMWMultisigAccount, _paymentInstruction ITeePaymentsBasePaymentInstruction, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.Pay(&_TeePaymentsBase.TransactOpts, _account, _paymentInstruction, _claimBackAddress)
}

// Reissue is a paid mutator transaction binding the contract method 0x1a82a605.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _batchPaymentId, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns(bool _finalized)
func (_TeePaymentsBase *TeePaymentsBaseTransactor) Reissue(opts *bind.TransactOpts, _account ITeePaymentsBasePMWMultisigAccount, _batchPaymentId uint64, _paymentInstructions []ITeePaymentsBasePaymentInstruction, _reissueFeeParams ITeePaymentsBaseReissueFeeParams, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsBase.contract.Transact(opts, "reissue", _account, _batchPaymentId, _paymentInstructions, _reissueFeeParams, _claimBackAddress)
}

// Reissue is a paid mutator transaction binding the contract method 0x1a82a605.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _batchPaymentId, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns(bool _finalized)
func (_TeePaymentsBase *TeePaymentsBaseSession) Reissue(_account ITeePaymentsBasePMWMultisigAccount, _batchPaymentId uint64, _paymentInstructions []ITeePaymentsBasePaymentInstruction, _reissueFeeParams ITeePaymentsBaseReissueFeeParams, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.Reissue(&_TeePaymentsBase.TransactOpts, _account, _batchPaymentId, _paymentInstructions, _reissueFeeParams, _claimBackAddress)
}

// Reissue is a paid mutator transaction binding the contract method 0x1a82a605.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _batchPaymentId, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns(bool _finalized)
func (_TeePaymentsBase *TeePaymentsBaseTransactorSession) Reissue(_account ITeePaymentsBasePMWMultisigAccount, _batchPaymentId uint64, _paymentInstructions []ITeePaymentsBasePaymentInstruction, _reissueFeeParams ITeePaymentsBaseReissueFeeParams, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.Reissue(&_TeePaymentsBase.TransactOpts, _account, _batchPaymentId, _paymentInstructions, _reissueFeeParams, _claimBackAddress)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsBase *TeePaymentsBaseTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsBase.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsBase *TeePaymentsBaseSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.SwitchToProductionMode(&_TeePaymentsBase.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsBase *TeePaymentsBaseTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.SwitchToProductionMode(&_TeePaymentsBase.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsBase *TeePaymentsBaseTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsBase.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsBase *TeePaymentsBaseSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.UpdateContractAddresses(&_TeePaymentsBase.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsBase *TeePaymentsBaseTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.UpdateContractAddresses(&_TeePaymentsBase.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsBase *TeePaymentsBaseTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsBase.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsBase *TeePaymentsBaseSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.UpgradeToAndCall(&_TeePaymentsBase.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsBase *TeePaymentsBaseTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsBase.Contract.UpgradeToAndCall(&_TeePaymentsBase.TransactOpts, _newImplementation, _data)
}

// TeePaymentsBaseGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeePaymentsBase contract.
type TeePaymentsBaseGovernanceCallTimelockedIterator struct {
	Event *TeePaymentsBaseGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeePaymentsBaseGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsBaseGovernanceCallTimelocked)
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
		it.Event = new(TeePaymentsBaseGovernanceCallTimelocked)
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
func (it *TeePaymentsBaseGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsBaseGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsBaseGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePaymentsBase contract.
type TeePaymentsBaseGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeePaymentsBaseGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeePaymentsBase.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsBaseGovernanceCallTimelockedIterator{contract: _TeePaymentsBase.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeePaymentsBaseGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsBase.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsBaseGovernanceCallTimelocked)
				if err := _TeePaymentsBase.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeePaymentsBaseGovernanceCallTimelocked, error) {
	event := new(TeePaymentsBaseGovernanceCallTimelocked)
	if err := _TeePaymentsBase.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsBaseGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeePaymentsBase contract.
type TeePaymentsBaseGovernanceInitialisedIterator struct {
	Event *TeePaymentsBaseGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeePaymentsBaseGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsBaseGovernanceInitialised)
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
		it.Event = new(TeePaymentsBaseGovernanceInitialised)
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
func (it *TeePaymentsBaseGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsBaseGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsBaseGovernanceInitialised represents a GovernanceInitialised event raised by the TeePaymentsBase contract.
type TeePaymentsBaseGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeePaymentsBaseGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeePaymentsBase.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsBaseGovernanceInitialisedIterator{contract: _TeePaymentsBase.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeePaymentsBaseGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsBase.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsBaseGovernanceInitialised)
				if err := _TeePaymentsBase.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeePaymentsBase *TeePaymentsBaseFilterer) ParseGovernanceInitialised(log types.Log) (*TeePaymentsBaseGovernanceInitialised, error) {
	event := new(TeePaymentsBaseGovernanceInitialised)
	if err := _TeePaymentsBase.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsBaseGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeePaymentsBase contract.
type TeePaymentsBaseGovernedProductionModeEnteredIterator struct {
	Event *TeePaymentsBaseGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeePaymentsBaseGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsBaseGovernedProductionModeEntered)
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
		it.Event = new(TeePaymentsBaseGovernedProductionModeEntered)
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
func (it *TeePaymentsBaseGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsBaseGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsBaseGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePaymentsBase contract.
type TeePaymentsBaseGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeePaymentsBaseGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeePaymentsBase.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsBaseGovernedProductionModeEnteredIterator{contract: _TeePaymentsBase.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeePaymentsBaseGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsBase.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsBaseGovernedProductionModeEntered)
				if err := _TeePaymentsBase.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeePaymentsBase *TeePaymentsBaseFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeePaymentsBaseGovernedProductionModeEntered, error) {
	event := new(TeePaymentsBaseGovernedProductionModeEntered)
	if err := _TeePaymentsBase.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TeePaymentsBase contract.
type TeePaymentsBaseInitializedIterator struct {
	Event *TeePaymentsBaseInitialized // Event containing the contract specifics and raw log

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
func (it *TeePaymentsBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsBaseInitialized)
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
		it.Event = new(TeePaymentsBaseInitialized)
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
func (it *TeePaymentsBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsBaseInitialized represents a Initialized event raised by the TeePaymentsBase contract.
type TeePaymentsBaseInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*TeePaymentsBaseInitializedIterator, error) {

	logs, sub, err := _TeePaymentsBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsBaseInitializedIterator{contract: _TeePaymentsBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TeePaymentsBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsBaseInitialized)
				if err := _TeePaymentsBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) ParseInitialized(log types.Log) (*TeePaymentsBaseInitialized, error) {
	event := new(TeePaymentsBaseInitialized)
	if err := _TeePaymentsBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsBaseTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeePaymentsBase contract.
type TeePaymentsBaseTimelockedGovernanceCallCanceledIterator struct {
	Event *TeePaymentsBaseTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeePaymentsBaseTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsBaseTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeePaymentsBaseTimelockedGovernanceCallCanceled)
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
func (it *TeePaymentsBaseTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsBaseTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsBaseTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePaymentsBase contract.
type TeePaymentsBaseTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeePaymentsBaseTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeePaymentsBase.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsBaseTimelockedGovernanceCallCanceledIterator{contract: _TeePaymentsBase.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeePaymentsBaseTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsBase.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsBaseTimelockedGovernanceCallCanceled)
				if err := _TeePaymentsBase.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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

// ParseTimelockedGovernanceCallCanceled is a log parse operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeePaymentsBaseTimelockedGovernanceCallCanceled, error) {
	event := new(TeePaymentsBaseTimelockedGovernanceCallCanceled)
	if err := _TeePaymentsBase.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsBaseTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeePaymentsBase contract.
type TeePaymentsBaseTimelockedGovernanceCallExecutedIterator struct {
	Event *TeePaymentsBaseTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeePaymentsBaseTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsBaseTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeePaymentsBaseTimelockedGovernanceCallExecuted)
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
func (it *TeePaymentsBaseTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsBaseTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsBaseTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePaymentsBase contract.
type TeePaymentsBaseTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeePaymentsBaseTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeePaymentsBase.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsBaseTimelockedGovernanceCallExecutedIterator{contract: _TeePaymentsBase.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeePaymentsBaseTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsBase.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsBaseTimelockedGovernanceCallExecuted)
				if err := _TeePaymentsBase.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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

// ParseTimelockedGovernanceCallExecuted is a log parse operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeePaymentsBaseTimelockedGovernanceCallExecuted, error) {
	event := new(TeePaymentsBaseTimelockedGovernanceCallExecuted)
	if err := _TeePaymentsBase.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsBaseUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeePaymentsBase contract.
type TeePaymentsBaseUpgradedIterator struct {
	Event *TeePaymentsBaseUpgraded // Event containing the contract specifics and raw log

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
func (it *TeePaymentsBaseUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsBaseUpgraded)
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
		it.Event = new(TeePaymentsBaseUpgraded)
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
func (it *TeePaymentsBaseUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsBaseUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsBaseUpgraded represents a Upgraded event raised by the TeePaymentsBase contract.
type TeePaymentsBaseUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeePaymentsBaseUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePaymentsBase.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsBaseUpgradedIterator{contract: _TeePaymentsBase.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePaymentsBase *TeePaymentsBaseFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeePaymentsBaseUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePaymentsBase.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsBaseUpgraded)
				if err := _TeePaymentsBase.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeePaymentsBase *TeePaymentsBaseFilterer) ParseUpgraded(log types.Log) (*TeePaymentsBaseUpgraded, error) {
	event := new(TeePaymentsBaseUpgraded)
	if err := _TeePaymentsBase.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
