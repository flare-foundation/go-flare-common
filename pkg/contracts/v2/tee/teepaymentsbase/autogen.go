// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepaymentsbase

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
var TeePaymentsBaseMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuthorizationAddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuthorizationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMWMultisigAccountAddressAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMWMultisigAccountNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockCallNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongKeyType\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressValidator\",\"outputs\":[{\"internalType\":\"contractIAddressValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"contractIIFlareTeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAuthorizationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getNextPaymentId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nextPaymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"}],\"name\":\"getPaymentFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getPaymentHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_paymentHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAccounts\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getWalletId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePaymentsBase.PaymentInstruction\",\"name\":\"_paymentInstruction\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentModel\",\"outputs\":[{\"internalType\":\"enumPaymentModel\",\"name\":\"_paymentModel\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_batchPaymentId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePaymentsBase.PaymentInstruction[]\",\"name\":\"_paymentInstructions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"maxFeePerPayment\",\"type\":\"uint256[]\"},{\"internalType\":\"int16[][]\",\"name\":\"factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"delaysSeconds\",\"type\":\"uint16[]\"}],\"internalType\":\"structITeePaymentsBase.ReissueFeeParams\",\"name\":\"_reissueFeeParams\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_startNew\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"reissue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_finalized\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsConfigVerifier\",\"outputs\":[{\"internalType\":\"contractITeePaymentsConfigVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsRegistry\",\"outputs\":[{\"internalType\":\"contractITeePaymentsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "TeePaymentsBase",
}

// TeePaymentsBase is an auto generated Go binding around an Ethereum contract.
type TeePaymentsBase struct {
	abi abi.ABI
}

// NewTeePaymentsBase creates a new instance of TeePaymentsBase.
func NewTeePaymentsBase() *TeePaymentsBase {
	parsed, err := TeePaymentsBaseMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TeePaymentsBase{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TeePaymentsBase) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePaymentsBase *TeePaymentsBase) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := teePaymentsBase.abi.Pack("UPGRADE_INTERFACE_VERSION")
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
func (teePaymentsBase *TeePaymentsBase) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return teePaymentsBase.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePaymentsBase *TeePaymentsBase) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := teePaymentsBase.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackAddressValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe37facf4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addressValidator() view returns(address)
func (teePaymentsBase *TeePaymentsBase) PackAddressValidator() []byte {
	enc, err := teePaymentsBase.abi.Pack("addressValidator")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddressValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe37facf4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addressValidator() view returns(address)
func (teePaymentsBase *TeePaymentsBase) TryPackAddressValidator() ([]byte, error) {
	return teePaymentsBase.abi.Pack("addressValidator")
}

// UnpackAddressValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe37facf4.
//
// Solidity: function addressValidator() view returns(address)
func (teePaymentsBase *TeePaymentsBase) UnpackAddressValidator(data []byte) (common.Address, error) {
	out, err := teePaymentsBase.abi.Unpack("addressValidator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16fc2f6d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (teePaymentsBase *TeePaymentsBase) PackCancelGovernanceCall(encodedCall []byte) []byte {
	enc, err := teePaymentsBase.abi.Pack("cancelGovernanceCall", encodedCall)
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
func (teePaymentsBase *TeePaymentsBase) TryPackCancelGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teePaymentsBase.abi.Pack("cancelGovernanceCall", encodedCall)
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (teePaymentsBase *TeePaymentsBase) PackExecuteGovernanceCall(encodedCall []byte) []byte {
	enc, err := teePaymentsBase.abi.Pack("executeGovernanceCall", encodedCall)
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
func (teePaymentsBase *TeePaymentsBase) TryPackExecuteGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teePaymentsBase.abi.Pack("executeGovernanceCall", encodedCall)
}

// PackFlareTeeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x453f7ab4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareTeeManager() view returns(address)
func (teePaymentsBase *TeePaymentsBase) PackFlareTeeManager() []byte {
	enc, err := teePaymentsBase.abi.Pack("flareTeeManager")
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
func (teePaymentsBase *TeePaymentsBase) TryPackFlareTeeManager() ([]byte, error) {
	return teePaymentsBase.abi.Pack("flareTeeManager")
}

// UnpackFlareTeeManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (teePaymentsBase *TeePaymentsBase) UnpackFlareTeeManager(data []byte) (common.Address, error) {
	out, err := teePaymentsBase.abi.Unpack("flareTeeManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (teePaymentsBase *TeePaymentsBase) PackGetAddressUpdater() []byte {
	enc, err := teePaymentsBase.abi.Pack("getAddressUpdater")
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
func (teePaymentsBase *TeePaymentsBase) TryPackGetAddressUpdater() ([]byte, error) {
	return teePaymentsBase.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (teePaymentsBase *TeePaymentsBase) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := teePaymentsBase.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAuthorizationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x410642e0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (teePaymentsBase *TeePaymentsBase) PackGetAuthorizationAddress(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePaymentsBase.abi.Pack("getAuthorizationAddress", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAuthorizationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x410642e0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (teePaymentsBase *TeePaymentsBase) TryPackGetAuthorizationAddress(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePaymentsBase.abi.Pack("getAuthorizationAddress", account)
}

// UnpackGetAuthorizationAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (teePaymentsBase *TeePaymentsBase) UnpackGetAuthorizationAddress(data []byte) (common.Address, error) {
	out, err := teePaymentsBase.abi.Unpack("getAuthorizationAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetNextPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfb49ac30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNextPaymentId((bytes32,string) _account) view returns(uint64 _nextPaymentId)
func (teePaymentsBase *TeePaymentsBase) PackGetNextPaymentId(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePaymentsBase.abi.Pack("getNextPaymentId", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNextPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfb49ac30.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNextPaymentId((bytes32,string) _account) view returns(uint64 _nextPaymentId)
func (teePaymentsBase *TeePaymentsBase) TryPackGetNextPaymentId(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePaymentsBase.abi.Pack("getNextPaymentId", account)
}

// UnpackGetNextPaymentId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfb49ac30.
//
// Solidity: function getNextPaymentId((bytes32,string) _account) view returns(uint64 _nextPaymentId)
func (teePaymentsBase *TeePaymentsBase) UnpackGetNextPaymentId(data []byte) (uint64, error) {
	out, err := teePaymentsBase.abi.Unpack("getNextPaymentId", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetPaymentFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57abf78b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (teePaymentsBase *TeePaymentsBase) PackGetPaymentFee(account ITeePaymentsBasePMWMultisigAccount, opCommand [32]byte) []byte {
	enc, err := teePaymentsBase.abi.Pack("getPaymentFee", account, opCommand)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPaymentFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57abf78b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (teePaymentsBase *TeePaymentsBase) TryPackGetPaymentFee(account ITeePaymentsBasePMWMultisigAccount, opCommand [32]byte) ([]byte, error) {
	return teePaymentsBase.abi.Pack("getPaymentFee", account, opCommand)
}

// UnpackGetPaymentFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x57abf78b.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (teePaymentsBase *TeePaymentsBase) UnpackGetPaymentFee(data []byte) (*big.Int, error) {
	out, err := teePaymentsBase.abi.Unpack("getPaymentFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetPaymentHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0079dfe8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPaymentHash((bytes32,string) _account, uint64 _paymentId) view returns(bytes32 _paymentHash)
func (teePaymentsBase *TeePaymentsBase) PackGetPaymentHash(account ITeePaymentsBasePMWMultisigAccount, paymentId uint64) []byte {
	enc, err := teePaymentsBase.abi.Pack("getPaymentHash", account, paymentId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPaymentHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0079dfe8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPaymentHash((bytes32,string) _account, uint64 _paymentId) view returns(bytes32 _paymentHash)
func (teePaymentsBase *TeePaymentsBase) TryPackGetPaymentHash(account ITeePaymentsBasePMWMultisigAccount, paymentId uint64) ([]byte, error) {
	return teePaymentsBase.abi.Pack("getPaymentHash", account, paymentId)
}

// UnpackGetPaymentHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0079dfe8.
//
// Solidity: function getPaymentHash((bytes32,string) _account, uint64 _paymentId) view returns(bytes32 _paymentHash)
func (teePaymentsBase *TeePaymentsBase) UnpackGetPaymentHash(data []byte) ([32]byte, error) {
	out, err := teePaymentsBase.abi.Unpack("getPaymentHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetWalletAccounts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a54c1b0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (teePaymentsBase *TeePaymentsBase) PackGetWalletAccounts(walletId [32]byte) []byte {
	enc, err := teePaymentsBase.abi.Pack("getWalletAccounts", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletAccounts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a54c1b0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (teePaymentsBase *TeePaymentsBase) TryPackGetWalletAccounts(walletId [32]byte) ([]byte, error) {
	return teePaymentsBase.abi.Pack("getWalletAccounts", walletId)
}

// UnpackGetWalletAccounts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (teePaymentsBase *TeePaymentsBase) UnpackGetWalletAccounts(data []byte) ([]ITeePaymentsBasePMWMultisigAccount, error) {
	out, err := teePaymentsBase.abi.Unpack("getWalletAccounts", data)
	if err != nil {
		return *new([]ITeePaymentsBasePMWMultisigAccount), err
	}
	out0 := *abi.ConvertType(out[0], new([]ITeePaymentsBasePMWMultisigAccount)).(*[]ITeePaymentsBasePMWMultisigAccount)
	return out0, nil
}

// PackGetWalletId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5623b3f5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (teePaymentsBase *TeePaymentsBase) PackGetWalletId(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePaymentsBase.abi.Pack("getWalletId", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5623b3f5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (teePaymentsBase *TeePaymentsBase) TryPackGetWalletId(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePaymentsBase.abi.Pack("getWalletId", account)
}

// UnpackGetWalletId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (teePaymentsBase *TeePaymentsBase) UnpackGetWalletId(data []byte) ([32]byte, error) {
	out, err := teePaymentsBase.abi.Unpack("getWalletId", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governance() view returns(address)
func (teePaymentsBase *TeePaymentsBase) PackGovernance() []byte {
	enc, err := teePaymentsBase.abi.Pack("governance")
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
func (teePaymentsBase *TeePaymentsBase) TryPackGovernance() ([]byte, error) {
	return teePaymentsBase.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (teePaymentsBase *TeePaymentsBase) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := teePaymentsBase.abi.Unpack("governance", data)
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
func (teePaymentsBase *TeePaymentsBase) PackGovernanceSettings() []byte {
	enc, err := teePaymentsBase.abi.Pack("governanceSettings")
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
func (teePaymentsBase *TeePaymentsBase) TryPackGovernanceSettings() ([]byte, error) {
	return teePaymentsBase.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (teePaymentsBase *TeePaymentsBase) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := teePaymentsBase.abi.Unpack("governanceSettings", data)
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
func (teePaymentsBase *TeePaymentsBase) PackImplementation() []byte {
	enc, err := teePaymentsBase.abi.Pack("implementation")
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
func (teePaymentsBase *TeePaymentsBase) TryPackImplementation() ([]byte, error) {
	return teePaymentsBase.abi.Pack("implementation")
}

// UnpackImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (teePaymentsBase *TeePaymentsBase) UnpackImplementation(data []byte) (common.Address, error) {
	out, err := teePaymentsBase.abi.Unpack("implementation", data)
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
func (teePaymentsBase *TeePaymentsBase) PackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) []byte {
	enc, err := teePaymentsBase.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
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
func (teePaymentsBase *TeePaymentsBase) TryPackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) ([]byte, error) {
	return teePaymentsBase.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teePaymentsBase *TeePaymentsBase) PackIsExecutor(address common.Address) []byte {
	enc, err := teePaymentsBase.abi.Pack("isExecutor", address)
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
func (teePaymentsBase *TeePaymentsBase) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return teePaymentsBase.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teePaymentsBase *TeePaymentsBase) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := teePaymentsBase.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackPay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x009ce938.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (teePaymentsBase *TeePaymentsBase) PackPay(account ITeePaymentsBasePMWMultisigAccount, paymentInstruction ITeePaymentsBasePaymentInstruction, claimBackAddress common.Address) []byte {
	enc, err := teePaymentsBase.abi.Pack("pay", account, paymentInstruction, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x009ce938.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (teePaymentsBase *TeePaymentsBase) TryPackPay(account ITeePaymentsBasePMWMultisigAccount, paymentInstruction ITeePaymentsBasePaymentInstruction, claimBackAddress common.Address) ([]byte, error) {
	return teePaymentsBase.abi.Pack("pay", account, paymentInstruction, claimBackAddress)
}

// UnpackPay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (teePaymentsBase *TeePaymentsBase) UnpackPay(data []byte) (uint64, error) {
	out, err := teePaymentsBase.abi.Unpack("pay", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackPaymentModel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb9d3dbb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function paymentModel() pure returns(uint8 _paymentModel)
func (teePaymentsBase *TeePaymentsBase) PackPaymentModel() []byte {
	enc, err := teePaymentsBase.abi.Pack("paymentModel")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPaymentModel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb9d3dbb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function paymentModel() pure returns(uint8 _paymentModel)
func (teePaymentsBase *TeePaymentsBase) TryPackPaymentModel() ([]byte, error) {
	return teePaymentsBase.abi.Pack("paymentModel")
}

// UnpackPaymentModel is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbb9d3dbb.
//
// Solidity: function paymentModel() pure returns(uint8 _paymentModel)
func (teePaymentsBase *TeePaymentsBase) UnpackPaymentModel(data []byte) (uint8, error) {
	out, err := teePaymentsBase.abi.Unpack("paymentModel", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (teePaymentsBase *TeePaymentsBase) PackProductionMode() []byte {
	enc, err := teePaymentsBase.abi.Pack("productionMode")
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
func (teePaymentsBase *TeePaymentsBase) TryPackProductionMode() ([]byte, error) {
	return teePaymentsBase.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (teePaymentsBase *TeePaymentsBase) UnpackProductionMode(data []byte) (bool, error) {
	out, err := teePaymentsBase.abi.Unpack("productionMode", data)
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
func (teePaymentsBase *TeePaymentsBase) PackProxiableUUID() []byte {
	enc, err := teePaymentsBase.abi.Pack("proxiableUUID")
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
func (teePaymentsBase *TeePaymentsBase) TryPackProxiableUUID() ([]byte, error) {
	return teePaymentsBase.abi.Pack("proxiableUUID")
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (teePaymentsBase *TeePaymentsBase) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := teePaymentsBase.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackReissue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1abcad2c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _batchPaymentId, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, bool _startNew, address _claimBackAddress) payable returns(bool _finalized)
func (teePaymentsBase *TeePaymentsBase) PackReissue(account ITeePaymentsBasePMWMultisigAccount, batchPaymentId uint64, paymentInstructions []ITeePaymentsBasePaymentInstruction, reissueFeeParams ITeePaymentsBaseReissueFeeParams, startNew bool, claimBackAddress common.Address) []byte {
	enc, err := teePaymentsBase.abi.Pack("reissue", account, batchPaymentId, paymentInstructions, reissueFeeParams, startNew, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReissue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1abcad2c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _batchPaymentId, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, bool _startNew, address _claimBackAddress) payable returns(bool _finalized)
func (teePaymentsBase *TeePaymentsBase) TryPackReissue(account ITeePaymentsBasePMWMultisigAccount, batchPaymentId uint64, paymentInstructions []ITeePaymentsBasePaymentInstruction, reissueFeeParams ITeePaymentsBaseReissueFeeParams, startNew bool, claimBackAddress common.Address) ([]byte, error) {
	return teePaymentsBase.abi.Pack("reissue", account, batchPaymentId, paymentInstructions, reissueFeeParams, startNew, claimBackAddress)
}

// UnpackReissue is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1abcad2c.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _batchPaymentId, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, bool _startNew, address _claimBackAddress) payable returns(bool _finalized)
func (teePaymentsBase *TeePaymentsBase) UnpackReissue(data []byte) (bool, error) {
	out, err := teePaymentsBase.abi.Unpack("reissue", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (teePaymentsBase *TeePaymentsBase) PackSwitchToProductionMode() []byte {
	enc, err := teePaymentsBase.abi.Pack("switchToProductionMode")
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
func (teePaymentsBase *TeePaymentsBase) TryPackSwitchToProductionMode() ([]byte, error) {
	return teePaymentsBase.abi.Pack("switchToProductionMode")
}

// PackTeePaymentsConfigVerifier is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf71c6c75.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (teePaymentsBase *TeePaymentsBase) PackTeePaymentsConfigVerifier() []byte {
	enc, err := teePaymentsBase.abi.Pack("teePaymentsConfigVerifier")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTeePaymentsConfigVerifier is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf71c6c75.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (teePaymentsBase *TeePaymentsBase) TryPackTeePaymentsConfigVerifier() ([]byte, error) {
	return teePaymentsBase.abi.Pack("teePaymentsConfigVerifier")
}

// UnpackTeePaymentsConfigVerifier is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf71c6c75.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (teePaymentsBase *TeePaymentsBase) UnpackTeePaymentsConfigVerifier(data []byte) (common.Address, error) {
	out, err := teePaymentsBase.abi.Unpack("teePaymentsConfigVerifier", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackTeePaymentsRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaef828de.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (teePaymentsBase *TeePaymentsBase) PackTeePaymentsRegistry() []byte {
	enc, err := teePaymentsBase.abi.Pack("teePaymentsRegistry")
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
func (teePaymentsBase *TeePaymentsBase) TryPackTeePaymentsRegistry() ([]byte, error) {
	return teePaymentsBase.abi.Pack("teePaymentsRegistry")
}

// UnpackTeePaymentsRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (teePaymentsBase *TeePaymentsBase) UnpackTeePaymentsRegistry(data []byte) (common.Address, error) {
	out, err := teePaymentsBase.abi.Unpack("teePaymentsRegistry", data)
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
func (teePaymentsBase *TeePaymentsBase) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := teePaymentsBase.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (teePaymentsBase *TeePaymentsBase) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return teePaymentsBase.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (teePaymentsBase *TeePaymentsBase) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := teePaymentsBase.abi.Pack("upgradeToAndCall", newImplementation, data)
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
func (teePaymentsBase *TeePaymentsBase) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return teePaymentsBase.abi.Pack("upgradeToAndCall", newImplementation, data)
}

// TeePaymentsBaseGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePaymentsBase contract.
type TeePaymentsBaseGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const TeePaymentsBaseGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (TeePaymentsBaseGovernanceCallTimelocked) ContractEventName() string {
	return TeePaymentsBaseGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (teePaymentsBase *TeePaymentsBase) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*TeePaymentsBaseGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsBase.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsBaseGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := teePaymentsBase.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsBase.abi.Events[event].Inputs {
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

// TeePaymentsBaseGovernanceInitialised represents a GovernanceInitialised event raised by the TeePaymentsBase contract.
type TeePaymentsBaseGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const TeePaymentsBaseGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (TeePaymentsBaseGovernanceInitialised) ContractEventName() string {
	return TeePaymentsBaseGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (teePaymentsBase *TeePaymentsBase) UnpackGovernanceInitialisedEvent(log *types.Log) (*TeePaymentsBaseGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsBase.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsBaseGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := teePaymentsBase.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsBase.abi.Events[event].Inputs {
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

// TeePaymentsBaseGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePaymentsBase contract.
type TeePaymentsBaseGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const TeePaymentsBaseGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (TeePaymentsBaseGovernedProductionModeEntered) ContractEventName() string {
	return TeePaymentsBaseGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (teePaymentsBase *TeePaymentsBase) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*TeePaymentsBaseGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsBase.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsBaseGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := teePaymentsBase.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsBase.abi.Events[event].Inputs {
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

// TeePaymentsBaseInitialized represents a Initialized event raised by the TeePaymentsBase contract.
type TeePaymentsBaseInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const TeePaymentsBaseInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (TeePaymentsBaseInitialized) ContractEventName() string {
	return TeePaymentsBaseInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (teePaymentsBase *TeePaymentsBase) UnpackInitializedEvent(log *types.Log) (*TeePaymentsBaseInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsBase.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsBaseInitialized)
	if len(log.Data) > 0 {
		if err := teePaymentsBase.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsBase.abi.Events[event].Inputs {
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

// TeePaymentsBaseTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePaymentsBase contract.
type TeePaymentsBaseTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeePaymentsBaseTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (TeePaymentsBaseTimelockedGovernanceCallCanceled) ContractEventName() string {
	return TeePaymentsBaseTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (teePaymentsBase *TeePaymentsBase) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*TeePaymentsBaseTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsBase.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsBaseTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := teePaymentsBase.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsBase.abi.Events[event].Inputs {
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

// TeePaymentsBaseTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePaymentsBase contract.
type TeePaymentsBaseTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeePaymentsBaseTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (TeePaymentsBaseTimelockedGovernanceCallExecuted) ContractEventName() string {
	return TeePaymentsBaseTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (teePaymentsBase *TeePaymentsBase) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*TeePaymentsBaseTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsBase.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsBaseTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := teePaymentsBase.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsBase.abi.Events[event].Inputs {
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

// TeePaymentsBaseUpgraded represents a Upgraded event raised by the TeePaymentsBase contract.
type TeePaymentsBaseUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const TeePaymentsBaseUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (TeePaymentsBaseUpgraded) ContractEventName() string {
	return TeePaymentsBaseUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (teePaymentsBase *TeePaymentsBase) UnpackUpgradedEvent(log *types.Log) (*TeePaymentsBaseUpgraded, error) {
	event := "Upgraded"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsBase.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsBaseUpgraded)
	if len(log.Data) > 0 {
		if err := teePaymentsBase.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsBase.abi.Events[event].Inputs {
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
func (teePaymentsBase *TeePaymentsBase) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["AlreadyInProductionMode"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackAlreadyInProductionModeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["AuthorizationAddressZero"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackAuthorizationAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["GovernedAddressZero"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackGovernedAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["GovernedAlreadyInitialized"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackGovernedAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["InvalidPaymentId"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackInvalidPaymentIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["InvalidRecipientAddress"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackInvalidRecipientAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["OnlyAuthorizationAddress"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackOnlyAuthorizationAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["OnlyExecutor"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["OnlySystemExtensionId"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackOnlySystemExtensionIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["OnlyWalletOwner"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackOnlyWalletOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["PMWMultisigAccountAddressAlreadySet"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackPMWMultisigAccountAddressAlreadySetError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["PMWMultisigAccountNotRegistered"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackPMWMultisigAccountNotRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["PaymentAmountZero"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackPaymentAmountZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["PaymentHashMismatch"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackPaymentHashMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["TimelockCallNotFound"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackTimelockCallNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["TimelockNotAllowedYet"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackTimelockNotAllowedYetError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["UnsupportedSourceId"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackUnsupportedSourceIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["WalletNotInProduction"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackWalletNotInProductionError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsBase.abi.Errors["WrongKeyType"].ID.Bytes()[:4]) {
		return teePaymentsBase.UnpackWrongKeyTypeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// TeePaymentsBaseAddressEmptyCode represents a AddressEmptyCode error raised by the TeePaymentsBase contract.
type TeePaymentsBaseAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func TeePaymentsBaseAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (teePaymentsBase *TeePaymentsBase) UnpackAddressEmptyCodeError(raw []byte) (*TeePaymentsBaseAddressEmptyCode, error) {
	out := new(TeePaymentsBaseAddressEmptyCode)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseAlreadyInProductionMode represents a AlreadyInProductionMode error raised by the TeePaymentsBase contract.
type TeePaymentsBaseAlreadyInProductionMode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInProductionMode()
func TeePaymentsBaseAlreadyInProductionModeErrorID() common.Hash {
	return common.HexToHash("0x72a23ec2b1fdbc3de7d5f411b572fcf8ad8409803196930f63ddf40a2a2b2594")
}

// UnpackAlreadyInProductionModeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInProductionMode()
func (teePaymentsBase *TeePaymentsBase) UnpackAlreadyInProductionModeError(raw []byte) (*TeePaymentsBaseAlreadyInProductionMode, error) {
	out := new(TeePaymentsBaseAlreadyInProductionMode)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "AlreadyInProductionMode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseAuthorizationAddressZero represents a AuthorizationAddressZero error raised by the TeePaymentsBase contract.
type TeePaymentsBaseAuthorizationAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AuthorizationAddressZero()
func TeePaymentsBaseAuthorizationAddressZeroErrorID() common.Hash {
	return common.HexToHash("0xe4472ca7994bab3d6f38c9bdc243c438c59c017951973a129c8674da85c299a2")
}

// UnpackAuthorizationAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AuthorizationAddressZero()
func (teePaymentsBase *TeePaymentsBase) UnpackAuthorizationAddressZeroError(raw []byte) (*TeePaymentsBaseAuthorizationAddressZero, error) {
	out := new(TeePaymentsBaseAuthorizationAddressZero)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "AuthorizationAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the TeePaymentsBase contract.
type TeePaymentsBaseERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func TeePaymentsBaseERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (teePaymentsBase *TeePaymentsBase) UnpackERC1967InvalidImplementationError(raw []byte) (*TeePaymentsBaseERC1967InvalidImplementation, error) {
	out := new(TeePaymentsBaseERC1967InvalidImplementation)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseERC1967NonPayable represents a ERC1967NonPayable error raised by the TeePaymentsBase contract.
type TeePaymentsBaseERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func TeePaymentsBaseERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (teePaymentsBase *TeePaymentsBase) UnpackERC1967NonPayableError(raw []byte) (*TeePaymentsBaseERC1967NonPayable, error) {
	out := new(TeePaymentsBaseERC1967NonPayable)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseFailedCall represents a FailedCall error raised by the TeePaymentsBase contract.
type TeePaymentsBaseFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func TeePaymentsBaseFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (teePaymentsBase *TeePaymentsBase) UnpackFailedCallError(raw []byte) (*TeePaymentsBaseFailedCall, error) {
	out := new(TeePaymentsBaseFailedCall)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseGovernedAddressZero represents a GovernedAddressZero error raised by the TeePaymentsBase contract.
type TeePaymentsBaseGovernedAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAddressZero()
func TeePaymentsBaseGovernedAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x695725f6b429bac0cc5a2fa88214e31520e0c69f2fb0e2fb881a7130a564f49f")
}

// UnpackGovernedAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAddressZero()
func (teePaymentsBase *TeePaymentsBase) UnpackGovernedAddressZeroError(raw []byte) (*TeePaymentsBaseGovernedAddressZero, error) {
	out := new(TeePaymentsBaseGovernedAddressZero)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "GovernedAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseGovernedAlreadyInitialized represents a GovernedAlreadyInitialized error raised by the TeePaymentsBase contract.
type TeePaymentsBaseGovernedAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAlreadyInitialized()
func TeePaymentsBaseGovernedAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0xb2f0162156d83828197776d56bbfba0a7a9d5936fad70f0a3014ae18cc7aa33c")
}

// UnpackGovernedAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAlreadyInitialized()
func (teePaymentsBase *TeePaymentsBase) UnpackGovernedAlreadyInitializedError(raw []byte) (*TeePaymentsBaseGovernedAlreadyInitialized, error) {
	out := new(TeePaymentsBaseGovernedAlreadyInitialized)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "GovernedAlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseInvalidInitialization represents a InvalidInitialization error raised by the TeePaymentsBase contract.
type TeePaymentsBaseInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func TeePaymentsBaseInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (teePaymentsBase *TeePaymentsBase) UnpackInvalidInitializationError(raw []byte) (*TeePaymentsBaseInvalidInitialization, error) {
	out := new(TeePaymentsBaseInvalidInitialization)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseInvalidPaymentId represents a InvalidPaymentId error raised by the TeePaymentsBase contract.
type TeePaymentsBaseInvalidPaymentId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPaymentId()
func TeePaymentsBaseInvalidPaymentIdErrorID() common.Hash {
	return common.HexToHash("0xb11be8c4d37d5cf7116085b4f06afe80345db997f69552f4e6a826a12eea0c0f")
}

// UnpackInvalidPaymentIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPaymentId()
func (teePaymentsBase *TeePaymentsBase) UnpackInvalidPaymentIdError(raw []byte) (*TeePaymentsBaseInvalidPaymentId, error) {
	out := new(TeePaymentsBaseInvalidPaymentId)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "InvalidPaymentId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseInvalidRecipientAddress represents a InvalidRecipientAddress error raised by the TeePaymentsBase contract.
type TeePaymentsBaseInvalidRecipientAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRecipientAddress()
func TeePaymentsBaseInvalidRecipientAddressErrorID() common.Hash {
	return common.HexToHash("0x44d99fea8d272f0b82a2cdf27d16958dace6a85ad6302ff443201a39bfc4820a")
}

// UnpackInvalidRecipientAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRecipientAddress()
func (teePaymentsBase *TeePaymentsBase) UnpackInvalidRecipientAddressError(raw []byte) (*TeePaymentsBaseInvalidRecipientAddress, error) {
	out := new(TeePaymentsBaseInvalidRecipientAddress)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "InvalidRecipientAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseLengthsMismatch represents a LengthsMismatch error raised by the TeePaymentsBase contract.
type TeePaymentsBaseLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func TeePaymentsBaseLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (teePaymentsBase *TeePaymentsBase) UnpackLengthsMismatchError(raw []byte) (*TeePaymentsBaseLengthsMismatch, error) {
	out := new(TeePaymentsBaseLengthsMismatch)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseNotInitializing represents a NotInitializing error raised by the TeePaymentsBase contract.
type TeePaymentsBaseNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func TeePaymentsBaseNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (teePaymentsBase *TeePaymentsBase) UnpackNotInitializingError(raw []byte) (*TeePaymentsBaseNotInitializing, error) {
	out := new(TeePaymentsBaseNotInitializing)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseOnlyAuthorizationAddress represents a OnlyAuthorizationAddress error raised by the TeePaymentsBase contract.
type TeePaymentsBaseOnlyAuthorizationAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyAuthorizationAddress()
func TeePaymentsBaseOnlyAuthorizationAddressErrorID() common.Hash {
	return common.HexToHash("0xdc65e71c09e77be7539dadee1dc86849e01c3e995f0c8d49e625476c2a1b1ed5")
}

// UnpackOnlyAuthorizationAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyAuthorizationAddress()
func (teePaymentsBase *TeePaymentsBase) UnpackOnlyAuthorizationAddressError(raw []byte) (*TeePaymentsBaseOnlyAuthorizationAddress, error) {
	out := new(TeePaymentsBaseOnlyAuthorizationAddress)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "OnlyAuthorizationAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseOnlyExecutor represents a OnlyExecutor error raised by the TeePaymentsBase contract.
type TeePaymentsBaseOnlyExecutor struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExecutor()
func TeePaymentsBaseOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x7fb6be02aaf40e420b0e7eb67e43a3feee949e8932a3c64e33b591a436670841")
}

// UnpackOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExecutor()
func (teePaymentsBase *TeePaymentsBase) UnpackOnlyExecutorError(raw []byte) (*TeePaymentsBaseOnlyExecutor, error) {
	out := new(TeePaymentsBaseOnlyExecutor)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "OnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseOnlyGovernance represents a OnlyGovernance error raised by the TeePaymentsBase contract.
type TeePaymentsBaseOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func TeePaymentsBaseOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (teePaymentsBase *TeePaymentsBase) UnpackOnlyGovernanceError(raw []byte) (*TeePaymentsBaseOnlyGovernance, error) {
	out := new(TeePaymentsBaseOnlyGovernance)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the TeePaymentsBase contract.
type TeePaymentsBaseOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func TeePaymentsBaseOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (teePaymentsBase *TeePaymentsBase) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*TeePaymentsBaseOnlyProductionOrPausedStatus, error) {
	out := new(TeePaymentsBaseOnlyProductionOrPausedStatus)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseOnlySystemExtensionId represents a OnlySystemExtensionId error raised by the TeePaymentsBase contract.
type TeePaymentsBaseOnlySystemExtensionId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemExtensionId()
func TeePaymentsBaseOnlySystemExtensionIdErrorID() common.Hash {
	return common.HexToHash("0xbd49094d5a86ee3f3c4b39274f5f98391dde2cd16618a917e121333ac3d43ee4")
}

// UnpackOnlySystemExtensionIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemExtensionId()
func (teePaymentsBase *TeePaymentsBase) UnpackOnlySystemExtensionIdError(raw []byte) (*TeePaymentsBaseOnlySystemExtensionId, error) {
	out := new(TeePaymentsBaseOnlySystemExtensionId)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "OnlySystemExtensionId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseOnlyWalletOwner represents a OnlyWalletOwner error raised by the TeePaymentsBase contract.
type TeePaymentsBaseOnlyWalletOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyWalletOwner()
func TeePaymentsBaseOnlyWalletOwnerErrorID() common.Hash {
	return common.HexToHash("0x3132d62dfddd1d8b6a9dd805001453bbb6a84b746ebe51fd650df1a6704a8143")
}

// UnpackOnlyWalletOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyWalletOwner()
func (teePaymentsBase *TeePaymentsBase) UnpackOnlyWalletOwnerError(raw []byte) (*TeePaymentsBaseOnlyWalletOwner, error) {
	out := new(TeePaymentsBaseOnlyWalletOwner)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "OnlyWalletOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBasePMWMultisigAccountAddressAlreadySet represents a PMWMultisigAccountAddressAlreadySet error raised by the TeePaymentsBase contract.
type TeePaymentsBasePMWMultisigAccountAddressAlreadySet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PMWMultisigAccountAddressAlreadySet()
func TeePaymentsBasePMWMultisigAccountAddressAlreadySetErrorID() common.Hash {
	return common.HexToHash("0xe3e804ecd31f69f9fde81b0b1994f9d3057e672dfe1cb39f32b68f667aa752e6")
}

// UnpackPMWMultisigAccountAddressAlreadySetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PMWMultisigAccountAddressAlreadySet()
func (teePaymentsBase *TeePaymentsBase) UnpackPMWMultisigAccountAddressAlreadySetError(raw []byte) (*TeePaymentsBasePMWMultisigAccountAddressAlreadySet, error) {
	out := new(TeePaymentsBasePMWMultisigAccountAddressAlreadySet)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "PMWMultisigAccountAddressAlreadySet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBasePMWMultisigAccountNotRegistered represents a PMWMultisigAccountNotRegistered error raised by the TeePaymentsBase contract.
type TeePaymentsBasePMWMultisigAccountNotRegistered struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PMWMultisigAccountNotRegistered()
func TeePaymentsBasePMWMultisigAccountNotRegisteredErrorID() common.Hash {
	return common.HexToHash("0x9d78814a2f7e263f4d4cc3000f076b4c90004105130f1880c0dc94acd2464be9")
}

// UnpackPMWMultisigAccountNotRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PMWMultisigAccountNotRegistered()
func (teePaymentsBase *TeePaymentsBase) UnpackPMWMultisigAccountNotRegisteredError(raw []byte) (*TeePaymentsBasePMWMultisigAccountNotRegistered, error) {
	out := new(TeePaymentsBasePMWMultisigAccountNotRegistered)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "PMWMultisigAccountNotRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBasePaymentAmountZero represents a PaymentAmountZero error raised by the TeePaymentsBase contract.
type TeePaymentsBasePaymentAmountZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentAmountZero()
func TeePaymentsBasePaymentAmountZeroErrorID() common.Hash {
	return common.HexToHash("0xa74a1c28a64fc70e7d718ae1b3af31f436e1c55b3754d607ef7ec829ec74034c")
}

// UnpackPaymentAmountZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentAmountZero()
func (teePaymentsBase *TeePaymentsBase) UnpackPaymentAmountZeroError(raw []byte) (*TeePaymentsBasePaymentAmountZero, error) {
	out := new(TeePaymentsBasePaymentAmountZero)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "PaymentAmountZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBasePaymentHashMismatch represents a PaymentHashMismatch error raised by the TeePaymentsBase contract.
type TeePaymentsBasePaymentHashMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentHashMismatch()
func TeePaymentsBasePaymentHashMismatchErrorID() common.Hash {
	return common.HexToHash("0x85628ea5f0ea30f18d9205be9d6ae9bd4b189b3dc00aeba346ab8f4c0e25674e")
}

// UnpackPaymentHashMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentHashMismatch()
func (teePaymentsBase *TeePaymentsBase) UnpackPaymentHashMismatchError(raw []byte) (*TeePaymentsBasePaymentHashMismatch, error) {
	out := new(TeePaymentsBasePaymentHashMismatch)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "PaymentHashMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseTimelockCallNotFound represents a TimelockCallNotFound error raised by the TeePaymentsBase contract.
type TeePaymentsBaseTimelockCallNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockCallNotFound()
func TeePaymentsBaseTimelockCallNotFoundErrorID() common.Hash {
	return common.HexToHash("0x262ac987dd09120380d393a0c52cdef1ca58949d1107dc26d8da9f7b527eb33d")
}

// UnpackTimelockCallNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockCallNotFound()
func (teePaymentsBase *TeePaymentsBase) UnpackTimelockCallNotFoundError(raw []byte) (*TeePaymentsBaseTimelockCallNotFound, error) {
	out := new(TeePaymentsBaseTimelockCallNotFound)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "TimelockCallNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseTimelockNotAllowedYet represents a TimelockNotAllowedYet error raised by the TeePaymentsBase contract.
type TeePaymentsBaseTimelockNotAllowedYet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockNotAllowedYet()
func TeePaymentsBaseTimelockNotAllowedYetErrorID() common.Hash {
	return common.HexToHash("0x6124e5c2c545bec1d02048c68500eebdabeae0cc69995a629ebccf8a62f69083")
}

// UnpackTimelockNotAllowedYetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockNotAllowedYet()
func (teePaymentsBase *TeePaymentsBase) UnpackTimelockNotAllowedYetError(raw []byte) (*TeePaymentsBaseTimelockNotAllowedYet, error) {
	out := new(TeePaymentsBaseTimelockNotAllowedYet)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "TimelockNotAllowedYet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the TeePaymentsBase contract.
type TeePaymentsBaseUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func TeePaymentsBaseUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (teePaymentsBase *TeePaymentsBase) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*TeePaymentsBaseUUPSUnauthorizedCallContext, error) {
	out := new(TeePaymentsBaseUUPSUnauthorizedCallContext)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the TeePaymentsBase contract.
type TeePaymentsBaseUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func TeePaymentsBaseUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (teePaymentsBase *TeePaymentsBase) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*TeePaymentsBaseUUPSUnsupportedProxiableUUID, error) {
	out := new(TeePaymentsBaseUUPSUnsupportedProxiableUUID)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseUnsupportedSourceId represents a UnsupportedSourceId error raised by the TeePaymentsBase contract.
type TeePaymentsBaseUnsupportedSourceId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedSourceId()
func TeePaymentsBaseUnsupportedSourceIdErrorID() common.Hash {
	return common.HexToHash("0xb9c2749e785a86b27011ac62f2eeace91e285734582a3d21db2e093d45e3124a")
}

// UnpackUnsupportedSourceIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedSourceId()
func (teePaymentsBase *TeePaymentsBase) UnpackUnsupportedSourceIdError(raw []byte) (*TeePaymentsBaseUnsupportedSourceId, error) {
	out := new(TeePaymentsBaseUnsupportedSourceId)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "UnsupportedSourceId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseWalletNotInProduction represents a WalletNotInProduction error raised by the TeePaymentsBase contract.
type TeePaymentsBaseWalletNotInProduction struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WalletNotInProduction()
func TeePaymentsBaseWalletNotInProductionErrorID() common.Hash {
	return common.HexToHash("0x2b7f97d6ba9d96a13708bf15f7c7ed50584fa90c4e3f979e82b455f65581f045")
}

// UnpackWalletNotInProductionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WalletNotInProduction()
func (teePaymentsBase *TeePaymentsBase) UnpackWalletNotInProductionError(raw []byte) (*TeePaymentsBaseWalletNotInProduction, error) {
	out := new(TeePaymentsBaseWalletNotInProduction)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "WalletNotInProduction", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsBaseWrongKeyType represents a WrongKeyType error raised by the TeePaymentsBase contract.
type TeePaymentsBaseWrongKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WrongKeyType()
func TeePaymentsBaseWrongKeyTypeErrorID() common.Hash {
	return common.HexToHash("0xfb224bce70b4062685f138c7eed15a72b903495475f7cd7843835692a7c58bbd")
}

// UnpackWrongKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WrongKeyType()
func (teePaymentsBase *TeePaymentsBase) UnpackWrongKeyTypeError(raw []byte) (*TeePaymentsBaseWrongKeyType, error) {
	out := new(TeePaymentsBaseWrongKeyType)
	if err := teePaymentsBase.abi.UnpackIntoInterface(out, "WrongKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}
