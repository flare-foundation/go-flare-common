// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepayment

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

// IFtdcHubFtdcResponseHeader is an auto generated low-level Go binding around an user-defined struct.
type IFtdcHubFtdcResponseHeader struct {
	AttestationType    [32]byte
	SourceId           [32]byte
	ThresholdBIPS      uint16
	Cosigners          []common.Address
	CosignersThreshold uint64
	Timestamp          uint64
}

// IFtdcVerificationFtdcSignatures is an auto generated low-level Go binding around an user-defined struct.
type IFtdcVerificationFtdcSignatures struct {
	SigningPolicySignatures []byte
	TeeSignatures           []Signature
	CosignerSignatures      []Signature
}

// IPMWMultisigAccountConfiguredProof is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredProof struct {
	Signatures   IFtdcVerificationFtdcSignatures
	Header       IFtdcHubFtdcResponseHeader
	RequestBody  IPMWMultisigAccountConfiguredRequestBody
	ResponseBody IPMWMultisigAccountConfiguredResponseBody
}

// IPMWMultisigAccountConfiguredRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredRequestBody struct {
	AccountAddress string
	PublicKeys     [][]byte
	Threshold      uint64
}

// IPMWMultisigAccountConfiguredResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredResponseBody struct {
	Status   uint8
	Sequence uint64
}

// ITeePaymentsPMWMultisigAccount is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsPMWMultisigAccount struct {
	SourceId       [32]byte
	AccountAddress string
}

// ITeePaymentsPaymentInstruction is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsPaymentInstruction struct {
	RecipientAddress string
	TokenId          [32]byte
	Amount           *big.Int
	Fee              *big.Int
	PaymentReference [32]byte
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// TeePaymentMetaData contains all meta data concerning the TeePayment contract.
var TeePaymentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountAddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchDurationTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotYetEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchSizeTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchSizeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DailyLimitBelowTransactionLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxBatchSizeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinFeeNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPaymentInstructions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySubmitAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OpTypeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMWMultisigAccountAddressAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SourceIdAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceIdZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SupportedSourceIdsLengthZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongOpType\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"BatchSettingsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"initialNonce\",\"type\":\"uint64\"}],\"name\":\"PMWMultisigAccountAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SupportedSourceIdAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PAY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REISSUE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SET_PAYMENT_LIMITS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigAccountConfigured.PMWMultisigAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"addPMWMultisigAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_sourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"addSupportedSourceIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getBatchSettings\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOpType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getOpTypeConstants\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedSourceIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAccounts\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getWalletId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_maxBatchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_maxBatchDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_supportedSourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"isSourceIdSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBatchDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBatchSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"tokenId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePayments.PaymentInstruction\",\"name\":\"_paymentInstruction\",\"type\":\"tuple\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_subNonce\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_firstSubNonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"tokenId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePayments.PaymentInstruction[]\",\"name\":\"_paymentInstructions\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_nullify\",\"type\":\"bool[]\"}],\"name\":\"reissue\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"setBatchSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_transactionLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"name\":\"setPaymentLimits\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeExtensionRegistry\",\"outputs\":[{\"internalType\":\"contractITeeExtensionRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeVerification\",\"outputs\":[{\"internalType\":\"contractITeeVerification\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletKeyManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletKeyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletProjectManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletProjectManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeePaymentABI is the input ABI used to generate the binding from.
// Deprecated: Use TeePaymentMetaData.ABI instead.
var TeePaymentABI = TeePaymentMetaData.ABI

// TeePayment is an auto generated Go binding around an Ethereum contract.
type TeePayment struct {
	TeePaymentCaller     // Read-only binding to the contract
	TeePaymentTransactor // Write-only binding to the contract
	TeePaymentFilterer   // Log filterer for contract events
}

// TeePaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeePaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeePaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeePaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeePaymentSession struct {
	Contract     *TeePayment       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeePaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeePaymentCallerSession struct {
	Contract *TeePaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TeePaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeePaymentTransactorSession struct {
	Contract     *TeePaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TeePaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeePaymentRaw struct {
	Contract *TeePayment // Generic contract binding to access the raw methods on
}

// TeePaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeePaymentCallerRaw struct {
	Contract *TeePaymentCaller // Generic read-only contract binding to access the raw methods on
}

// TeePaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeePaymentTransactorRaw struct {
	Contract *TeePaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeePayment creates a new instance of TeePayment, bound to a specific deployed contract.
func NewTeePayment(address common.Address, backend bind.ContractBackend) (*TeePayment, error) {
	contract, err := bindTeePayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeePayment{TeePaymentCaller: TeePaymentCaller{contract: contract}, TeePaymentTransactor: TeePaymentTransactor{contract: contract}, TeePaymentFilterer: TeePaymentFilterer{contract: contract}}, nil
}

// NewTeePaymentCaller creates a new read-only instance of TeePayment, bound to a specific deployed contract.
func NewTeePaymentCaller(address common.Address, caller bind.ContractCaller) (*TeePaymentCaller, error) {
	contract, err := bindTeePayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentCaller{contract: contract}, nil
}

// NewTeePaymentTransactor creates a new write-only instance of TeePayment, bound to a specific deployed contract.
func NewTeePaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*TeePaymentTransactor, error) {
	contract, err := bindTeePayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentTransactor{contract: contract}, nil
}

// NewTeePaymentFilterer creates a new log filterer instance of TeePayment, bound to a specific deployed contract.
func NewTeePaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*TeePaymentFilterer, error) {
	contract, err := bindTeePayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeePaymentFilterer{contract: contract}, nil
}

// bindTeePayment binds a generic wrapper to an already deployed contract.
func bindTeePayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeePaymentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePayment *TeePaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePayment.Contract.TeePaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePayment *TeePaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePayment.Contract.TeePaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePayment *TeePaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePayment.Contract.TeePaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePayment *TeePaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePayment *TeePaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePayment *TeePaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePayment.Contract.contract.Transact(opts, method, params...)
}

// PAY is a free data retrieval call binding the contract method 0x72e7ef29.
//
// Solidity: function PAY() view returns(bytes32)
func (_TeePayment *TeePaymentCaller) PAY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "PAY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAY is a free data retrieval call binding the contract method 0x72e7ef29.
//
// Solidity: function PAY() view returns(bytes32)
func (_TeePayment *TeePaymentSession) PAY() ([32]byte, error) {
	return _TeePayment.Contract.PAY(&_TeePayment.CallOpts)
}

// PAY is a free data retrieval call binding the contract method 0x72e7ef29.
//
// Solidity: function PAY() view returns(bytes32)
func (_TeePayment *TeePaymentCallerSession) PAY() ([32]byte, error) {
	return _TeePayment.Contract.PAY(&_TeePayment.CallOpts)
}

// REISSUE is a free data retrieval call binding the contract method 0x1df89bbc.
//
// Solidity: function REISSUE() view returns(bytes32)
func (_TeePayment *TeePaymentCaller) REISSUE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "REISSUE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REISSUE is a free data retrieval call binding the contract method 0x1df89bbc.
//
// Solidity: function REISSUE() view returns(bytes32)
func (_TeePayment *TeePaymentSession) REISSUE() ([32]byte, error) {
	return _TeePayment.Contract.REISSUE(&_TeePayment.CallOpts)
}

// REISSUE is a free data retrieval call binding the contract method 0x1df89bbc.
//
// Solidity: function REISSUE() view returns(bytes32)
func (_TeePayment *TeePaymentCallerSession) REISSUE() ([32]byte, error) {
	return _TeePayment.Contract.REISSUE(&_TeePayment.CallOpts)
}

// SETPAYMENTLIMITS is a free data retrieval call binding the contract method 0x9c42d648.
//
// Solidity: function SET_PAYMENT_LIMITS() view returns(bytes32)
func (_TeePayment *TeePaymentCaller) SETPAYMENTLIMITS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "SET_PAYMENT_LIMITS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETPAYMENTLIMITS is a free data retrieval call binding the contract method 0x9c42d648.
//
// Solidity: function SET_PAYMENT_LIMITS() view returns(bytes32)
func (_TeePayment *TeePaymentSession) SETPAYMENTLIMITS() ([32]byte, error) {
	return _TeePayment.Contract.SETPAYMENTLIMITS(&_TeePayment.CallOpts)
}

// SETPAYMENTLIMITS is a free data retrieval call binding the contract method 0x9c42d648.
//
// Solidity: function SET_PAYMENT_LIMITS() view returns(bytes32)
func (_TeePayment *TeePaymentCallerSession) SETPAYMENTLIMITS() ([32]byte, error) {
	return _TeePayment.Contract.SETPAYMENTLIMITS(&_TeePayment.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePayment *TeePaymentCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePayment *TeePaymentSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePayment.Contract.UPGRADEINTERFACEVERSION(&_TeePayment.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePayment *TeePaymentCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePayment.Contract.UPGRADEINTERFACEVERSION(&_TeePayment.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePayment *TeePaymentCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePayment *TeePaymentSession) FlareSystemsManager() (common.Address, error) {
	return _TeePayment.Contract.FlareSystemsManager(&_TeePayment.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePayment *TeePaymentCallerSession) FlareSystemsManager() (common.Address, error) {
	return _TeePayment.Contract.FlareSystemsManager(&_TeePayment.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePayment *TeePaymentCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePayment *TeePaymentSession) GetAddressUpdater() (common.Address, error) {
	return _TeePayment.Contract.GetAddressUpdater(&_TeePayment.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePayment *TeePaymentCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeePayment.Contract.GetAddressUpdater(&_TeePayment.CallOpts)
}

// GetBatchSettings is a free data retrieval call binding the contract method 0xb96c3c8f.
//
// Solidity: function getBatchSettings((bytes32,string) _account) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (_TeePayment *TeePaymentCaller) GetBatchSettings(opts *bind.CallOpts, _account ITeePaymentsPMWMultisigAccount) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "getBatchSettings", _account)

	outstruct := new(struct {
		BatchSize            uint64
		BatchDurationSeconds uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BatchSize = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.BatchDurationSeconds = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetBatchSettings is a free data retrieval call binding the contract method 0xb96c3c8f.
//
// Solidity: function getBatchSettings((bytes32,string) _account) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (_TeePayment *TeePaymentSession) GetBatchSettings(_account ITeePaymentsPMWMultisigAccount) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	return _TeePayment.Contract.GetBatchSettings(&_TeePayment.CallOpts, _account)
}

// GetBatchSettings is a free data retrieval call binding the contract method 0xb96c3c8f.
//
// Solidity: function getBatchSettings((bytes32,string) _account) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (_TeePayment *TeePaymentCallerSession) GetBatchSettings(_account ITeePaymentsPMWMultisigAccount) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	return _TeePayment.Contract.GetBatchSettings(&_TeePayment.CallOpts, _account)
}

// GetOpType is a free data retrieval call binding the contract method 0xa3e2d39a.
//
// Solidity: function getOpType() view returns(bytes32)
func (_TeePayment *TeePaymentCaller) GetOpType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "getOpType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOpType is a free data retrieval call binding the contract method 0xa3e2d39a.
//
// Solidity: function getOpType() view returns(bytes32)
func (_TeePayment *TeePaymentSession) GetOpType() ([32]byte, error) {
	return _TeePayment.Contract.GetOpType(&_TeePayment.CallOpts)
}

// GetOpType is a free data retrieval call binding the contract method 0xa3e2d39a.
//
// Solidity: function getOpType() view returns(bytes32)
func (_TeePayment *TeePaymentCallerSession) GetOpType() ([32]byte, error) {
	return _TeePayment.Contract.GetOpType(&_TeePayment.CallOpts)
}

// GetOpTypeConstants is a free data retrieval call binding the contract method 0x727b7d21.
//
// Solidity: function getOpTypeConstants(bytes32 _projectId) view returns(bytes)
func (_TeePayment *TeePaymentCaller) GetOpTypeConstants(opts *bind.CallOpts, _projectId [32]byte) ([]byte, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "getOpTypeConstants", _projectId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetOpTypeConstants is a free data retrieval call binding the contract method 0x727b7d21.
//
// Solidity: function getOpTypeConstants(bytes32 _projectId) view returns(bytes)
func (_TeePayment *TeePaymentSession) GetOpTypeConstants(_projectId [32]byte) ([]byte, error) {
	return _TeePayment.Contract.GetOpTypeConstants(&_TeePayment.CallOpts, _projectId)
}

// GetOpTypeConstants is a free data retrieval call binding the contract method 0x727b7d21.
//
// Solidity: function getOpTypeConstants(bytes32 _projectId) view returns(bytes)
func (_TeePayment *TeePaymentCallerSession) GetOpTypeConstants(_projectId [32]byte) ([]byte, error) {
	return _TeePayment.Contract.GetOpTypeConstants(&_TeePayment.CallOpts, _projectId)
}

// GetSupportedSourceIds is a free data retrieval call binding the contract method 0x3141aae6.
//
// Solidity: function getSupportedSourceIds() view returns(bytes32[])
func (_TeePayment *TeePaymentCaller) GetSupportedSourceIds(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "getSupportedSourceIds")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSupportedSourceIds is a free data retrieval call binding the contract method 0x3141aae6.
//
// Solidity: function getSupportedSourceIds() view returns(bytes32[])
func (_TeePayment *TeePaymentSession) GetSupportedSourceIds() ([][32]byte, error) {
	return _TeePayment.Contract.GetSupportedSourceIds(&_TeePayment.CallOpts)
}

// GetSupportedSourceIds is a free data retrieval call binding the contract method 0x3141aae6.
//
// Solidity: function getSupportedSourceIds() view returns(bytes32[])
func (_TeePayment *TeePaymentCallerSession) GetSupportedSourceIds() ([][32]byte, error) {
	return _TeePayment.Contract.GetSupportedSourceIds(&_TeePayment.CallOpts)
}

// GetWalletAccounts is a free data retrieval call binding the contract method 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (_TeePayment *TeePaymentCaller) GetWalletAccounts(opts *bind.CallOpts, _walletId [32]byte) ([]ITeePaymentsPMWMultisigAccount, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "getWalletAccounts", _walletId)

	if err != nil {
		return *new([]ITeePaymentsPMWMultisigAccount), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITeePaymentsPMWMultisigAccount)).(*[]ITeePaymentsPMWMultisigAccount)

	return out0, err

}

// GetWalletAccounts is a free data retrieval call binding the contract method 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (_TeePayment *TeePaymentSession) GetWalletAccounts(_walletId [32]byte) ([]ITeePaymentsPMWMultisigAccount, error) {
	return _TeePayment.Contract.GetWalletAccounts(&_TeePayment.CallOpts, _walletId)
}

// GetWalletAccounts is a free data retrieval call binding the contract method 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (_TeePayment *TeePaymentCallerSession) GetWalletAccounts(_walletId [32]byte) ([]ITeePaymentsPMWMultisigAccount, error) {
	return _TeePayment.Contract.GetWalletAccounts(&_TeePayment.CallOpts, _walletId)
}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (_TeePayment *TeePaymentCaller) GetWalletId(opts *bind.CallOpts, _account ITeePaymentsPMWMultisigAccount) ([32]byte, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "getWalletId", _account)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (_TeePayment *TeePaymentSession) GetWalletId(_account ITeePaymentsPMWMultisigAccount) ([32]byte, error) {
	return _TeePayment.Contract.GetWalletId(&_TeePayment.CallOpts, _account)
}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (_TeePayment *TeePaymentCallerSession) GetWalletId(_account ITeePaymentsPMWMultisigAccount) ([32]byte, error) {
	return _TeePayment.Contract.GetWalletId(&_TeePayment.CallOpts, _account)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePayment *TeePaymentCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePayment *TeePaymentSession) Governance() (common.Address, error) {
	return _TeePayment.Contract.Governance(&_TeePayment.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePayment *TeePaymentCallerSession) Governance() (common.Address, error) {
	return _TeePayment.Contract.Governance(&_TeePayment.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePayment *TeePaymentCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePayment *TeePaymentSession) GovernanceSettings() (common.Address, error) {
	return _TeePayment.Contract.GovernanceSettings(&_TeePayment.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePayment *TeePaymentCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeePayment.Contract.GovernanceSettings(&_TeePayment.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePayment *TeePaymentCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePayment *TeePaymentSession) Implementation() (common.Address, error) {
	return _TeePayment.Contract.Implementation(&_TeePayment.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePayment *TeePaymentCallerSession) Implementation() (common.Address, error) {
	return _TeePayment.Contract.Implementation(&_TeePayment.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePayment *TeePaymentCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePayment *TeePaymentSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePayment.Contract.IsExecutor(&_TeePayment.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePayment *TeePaymentCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePayment.Contract.IsExecutor(&_TeePayment.CallOpts, _address)
}

// IsSourceIdSupported is a free data retrieval call binding the contract method 0x7e0bbbf3.
//
// Solidity: function isSourceIdSupported(bytes32 _sourceId) view returns(bool)
func (_TeePayment *TeePaymentCaller) IsSourceIdSupported(opts *bind.CallOpts, _sourceId [32]byte) (bool, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "isSourceIdSupported", _sourceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSourceIdSupported is a free data retrieval call binding the contract method 0x7e0bbbf3.
//
// Solidity: function isSourceIdSupported(bytes32 _sourceId) view returns(bool)
func (_TeePayment *TeePaymentSession) IsSourceIdSupported(_sourceId [32]byte) (bool, error) {
	return _TeePayment.Contract.IsSourceIdSupported(&_TeePayment.CallOpts, _sourceId)
}

// IsSourceIdSupported is a free data retrieval call binding the contract method 0x7e0bbbf3.
//
// Solidity: function isSourceIdSupported(bytes32 _sourceId) view returns(bool)
func (_TeePayment *TeePaymentCallerSession) IsSourceIdSupported(_sourceId [32]byte) (bool, error) {
	return _TeePayment.Contract.IsSourceIdSupported(&_TeePayment.CallOpts, _sourceId)
}

// MaxBatchDurationSeconds is a free data retrieval call binding the contract method 0xfdef3a1e.
//
// Solidity: function maxBatchDurationSeconds() view returns(uint64)
func (_TeePayment *TeePaymentCaller) MaxBatchDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "maxBatchDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxBatchDurationSeconds is a free data retrieval call binding the contract method 0xfdef3a1e.
//
// Solidity: function maxBatchDurationSeconds() view returns(uint64)
func (_TeePayment *TeePaymentSession) MaxBatchDurationSeconds() (uint64, error) {
	return _TeePayment.Contract.MaxBatchDurationSeconds(&_TeePayment.CallOpts)
}

// MaxBatchDurationSeconds is a free data retrieval call binding the contract method 0xfdef3a1e.
//
// Solidity: function maxBatchDurationSeconds() view returns(uint64)
func (_TeePayment *TeePaymentCallerSession) MaxBatchDurationSeconds() (uint64, error) {
	return _TeePayment.Contract.MaxBatchDurationSeconds(&_TeePayment.CallOpts)
}

// MaxBatchSize is a free data retrieval call binding the contract method 0x2913daa0.
//
// Solidity: function maxBatchSize() view returns(uint64)
func (_TeePayment *TeePaymentCaller) MaxBatchSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "maxBatchSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxBatchSize is a free data retrieval call binding the contract method 0x2913daa0.
//
// Solidity: function maxBatchSize() view returns(uint64)
func (_TeePayment *TeePaymentSession) MaxBatchSize() (uint64, error) {
	return _TeePayment.Contract.MaxBatchSize(&_TeePayment.CallOpts)
}

// MaxBatchSize is a free data retrieval call binding the contract method 0x2913daa0.
//
// Solidity: function maxBatchSize() view returns(uint64)
func (_TeePayment *TeePaymentCallerSession) MaxBatchSize() (uint64, error) {
	return _TeePayment.Contract.MaxBatchSize(&_TeePayment.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePayment *TeePaymentCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePayment *TeePaymentSession) ProductionMode() (bool, error) {
	return _TeePayment.Contract.ProductionMode(&_TeePayment.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePayment *TeePaymentCallerSession) ProductionMode() (bool, error) {
	return _TeePayment.Contract.ProductionMode(&_TeePayment.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePayment *TeePaymentCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePayment *TeePaymentSession) ProxiableUUID() ([32]byte, error) {
	return _TeePayment.Contract.ProxiableUUID(&_TeePayment.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePayment *TeePaymentCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeePayment.Contract.ProxiableUUID(&_TeePayment.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeePayment *TeePaymentCaller) TeeExtensionRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "teeExtensionRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeePayment *TeePaymentSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeePayment.Contract.TeeExtensionRegistry(&_TeePayment.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeePayment *TeePaymentCallerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeePayment.Contract.TeeExtensionRegistry(&_TeePayment.CallOpts)
}

// TeeVerification is a free data retrieval call binding the contract method 0x70dcd962.
//
// Solidity: function teeVerification() view returns(address)
func (_TeePayment *TeePaymentCaller) TeeVerification(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "teeVerification")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeVerification is a free data retrieval call binding the contract method 0x70dcd962.
//
// Solidity: function teeVerification() view returns(address)
func (_TeePayment *TeePaymentSession) TeeVerification() (common.Address, error) {
	return _TeePayment.Contract.TeeVerification(&_TeePayment.CallOpts)
}

// TeeVerification is a free data retrieval call binding the contract method 0x70dcd962.
//
// Solidity: function teeVerification() view returns(address)
func (_TeePayment *TeePaymentCallerSession) TeeVerification() (common.Address, error) {
	return _TeePayment.Contract.TeeVerification(&_TeePayment.CallOpts)
}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeePayment *TeePaymentCaller) TeeWalletKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "teeWalletKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeePayment *TeePaymentSession) TeeWalletKeyManager() (common.Address, error) {
	return _TeePayment.Contract.TeeWalletKeyManager(&_TeePayment.CallOpts)
}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeePayment *TeePaymentCallerSession) TeeWalletKeyManager() (common.Address, error) {
	return _TeePayment.Contract.TeeWalletKeyManager(&_TeePayment.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeePayment *TeePaymentCaller) TeeWalletManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "teeWalletManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeePayment *TeePaymentSession) TeeWalletManager() (common.Address, error) {
	return _TeePayment.Contract.TeeWalletManager(&_TeePayment.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeePayment *TeePaymentCallerSession) TeeWalletManager() (common.Address, error) {
	return _TeePayment.Contract.TeeWalletManager(&_TeePayment.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeePayment *TeePaymentCaller) TeeWalletProjectManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "teeWalletProjectManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeePayment *TeePaymentSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeePayment.Contract.TeeWalletProjectManager(&_TeePayment.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeePayment *TeePaymentCallerSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeePayment.Contract.TeeWalletProjectManager(&_TeePayment.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePayment *TeePaymentCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_TeePayment *TeePaymentSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeePayment.Contract.TimelockedCalls(&_TeePayment.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePayment *TeePaymentCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeePayment.Contract.TimelockedCalls(&_TeePayment.CallOpts, selector)
}

// AddPMWMultisigAccount is a paid mutator transaction binding the contract method 0xdfda8e99.
//
// Solidity: function addPMWMultisigAccount(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns()
func (_TeePayment *TeePaymentTransactor) AddPMWMultisigAccount(opts *bind.TransactOpts, _walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "addPMWMultisigAccount", _walletId, _proof)
}

// AddPMWMultisigAccount is a paid mutator transaction binding the contract method 0xdfda8e99.
//
// Solidity: function addPMWMultisigAccount(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns()
func (_TeePayment *TeePaymentSession) AddPMWMultisigAccount(_walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeePayment.Contract.AddPMWMultisigAccount(&_TeePayment.TransactOpts, _walletId, _proof)
}

// AddPMWMultisigAccount is a paid mutator transaction binding the contract method 0xdfda8e99.
//
// Solidity: function addPMWMultisigAccount(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns()
func (_TeePayment *TeePaymentTransactorSession) AddPMWMultisigAccount(_walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeePayment.Contract.AddPMWMultisigAccount(&_TeePayment.TransactOpts, _walletId, _proof)
}

// AddSupportedSourceIds is a paid mutator transaction binding the contract method 0x7e57ebed.
//
// Solidity: function addSupportedSourceIds(bytes32[] _sourceIds) returns()
func (_TeePayment *TeePaymentTransactor) AddSupportedSourceIds(opts *bind.TransactOpts, _sourceIds [][32]byte) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "addSupportedSourceIds", _sourceIds)
}

// AddSupportedSourceIds is a paid mutator transaction binding the contract method 0x7e57ebed.
//
// Solidity: function addSupportedSourceIds(bytes32[] _sourceIds) returns()
func (_TeePayment *TeePaymentSession) AddSupportedSourceIds(_sourceIds [][32]byte) (*types.Transaction, error) {
	return _TeePayment.Contract.AddSupportedSourceIds(&_TeePayment.TransactOpts, _sourceIds)
}

// AddSupportedSourceIds is a paid mutator transaction binding the contract method 0x7e57ebed.
//
// Solidity: function addSupportedSourceIds(bytes32[] _sourceIds) returns()
func (_TeePayment *TeePaymentTransactorSession) AddSupportedSourceIds(_sourceIds [][32]byte) (*types.Transaction, error) {
	return _TeePayment.Contract.AddSupportedSourceIds(&_TeePayment.TransactOpts, _sourceIds)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePayment *TeePaymentTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePayment *TeePaymentSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePayment.Contract.CancelGovernanceCall(&_TeePayment.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePayment *TeePaymentTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePayment.Contract.CancelGovernanceCall(&_TeePayment.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePayment *TeePaymentTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePayment *TeePaymentSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePayment.Contract.ExecuteGovernanceCall(&_TeePayment.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePayment *TeePaymentTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePayment.Contract.ExecuteGovernanceCall(&_TeePayment.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePayment *TeePaymentTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePayment *TeePaymentSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePayment.Contract.Initialise(&_TeePayment.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePayment *TeePaymentTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePayment.Contract.Initialise(&_TeePayment.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xbdaf95f3.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds, bytes32 _opType, bytes32[] _supportedSourceIds) returns()
func (_TeePayment *TeePaymentTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _maxBatchSize uint64, _maxBatchDurationSeconds uint64, _opType [32]byte, _supportedSourceIds [][32]byte) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater, _maxBatchSize, _maxBatchDurationSeconds, _opType, _supportedSourceIds)
}

// Initialize is a paid mutator transaction binding the contract method 0xbdaf95f3.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds, bytes32 _opType, bytes32[] _supportedSourceIds) returns()
func (_TeePayment *TeePaymentSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _maxBatchSize uint64, _maxBatchDurationSeconds uint64, _opType [32]byte, _supportedSourceIds [][32]byte) (*types.Transaction, error) {
	return _TeePayment.Contract.Initialize(&_TeePayment.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _maxBatchSize, _maxBatchDurationSeconds, _opType, _supportedSourceIds)
}

// Initialize is a paid mutator transaction binding the contract method 0xbdaf95f3.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds, bytes32 _opType, bytes32[] _supportedSourceIds) returns()
func (_TeePayment *TeePaymentTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _maxBatchSize uint64, _maxBatchDurationSeconds uint64, _opType [32]byte, _supportedSourceIds [][32]byte) (*types.Transaction, error) {
	return _TeePayment.Contract.Initialize(&_TeePayment.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _maxBatchSize, _maxBatchDurationSeconds, _opType, _supportedSourceIds)
}

// Pay is a paid mutator transaction binding the contract method 0x69666b83.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes32,uint256,uint256,bytes32) _paymentInstruction) payable returns(uint64 _nonce, uint64 _subNonce)
func (_TeePayment *TeePaymentTransactor) Pay(opts *bind.TransactOpts, _account ITeePaymentsPMWMultisigAccount, _paymentInstruction ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "pay", _account, _paymentInstruction)
}

// Pay is a paid mutator transaction binding the contract method 0x69666b83.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes32,uint256,uint256,bytes32) _paymentInstruction) payable returns(uint64 _nonce, uint64 _subNonce)
func (_TeePayment *TeePaymentSession) Pay(_account ITeePaymentsPMWMultisigAccount, _paymentInstruction ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _TeePayment.Contract.Pay(&_TeePayment.TransactOpts, _account, _paymentInstruction)
}

// Pay is a paid mutator transaction binding the contract method 0x69666b83.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes32,uint256,uint256,bytes32) _paymentInstruction) payable returns(uint64 _nonce, uint64 _subNonce)
func (_TeePayment *TeePaymentTransactorSession) Pay(_account ITeePaymentsPMWMultisigAccount, _paymentInstruction ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _TeePayment.Contract.Pay(&_TeePayment.TransactOpts, _account, _paymentInstruction)
}

// Reissue is a paid mutator transaction binding the contract method 0x404157e8.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _nonce, uint64 _firstSubNonce, (string,bytes32,uint256,uint256,bytes32)[] _paymentInstructions, uint256[] _fees, bool[] _nullify) payable returns()
func (_TeePayment *TeePaymentTransactor) Reissue(opts *bind.TransactOpts, _account ITeePaymentsPMWMultisigAccount, _nonce uint64, _firstSubNonce uint64, _paymentInstructions []ITeePaymentsPaymentInstruction, _fees []*big.Int, _nullify []bool) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "reissue", _account, _nonce, _firstSubNonce, _paymentInstructions, _fees, _nullify)
}

// Reissue is a paid mutator transaction binding the contract method 0x404157e8.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _nonce, uint64 _firstSubNonce, (string,bytes32,uint256,uint256,bytes32)[] _paymentInstructions, uint256[] _fees, bool[] _nullify) payable returns()
func (_TeePayment *TeePaymentSession) Reissue(_account ITeePaymentsPMWMultisigAccount, _nonce uint64, _firstSubNonce uint64, _paymentInstructions []ITeePaymentsPaymentInstruction, _fees []*big.Int, _nullify []bool) (*types.Transaction, error) {
	return _TeePayment.Contract.Reissue(&_TeePayment.TransactOpts, _account, _nonce, _firstSubNonce, _paymentInstructions, _fees, _nullify)
}

// Reissue is a paid mutator transaction binding the contract method 0x404157e8.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _nonce, uint64 _firstSubNonce, (string,bytes32,uint256,uint256,bytes32)[] _paymentInstructions, uint256[] _fees, bool[] _nullify) payable returns()
func (_TeePayment *TeePaymentTransactorSession) Reissue(_account ITeePaymentsPMWMultisigAccount, _nonce uint64, _firstSubNonce uint64, _paymentInstructions []ITeePaymentsPaymentInstruction, _fees []*big.Int, _nullify []bool) (*types.Transaction, error) {
	return _TeePayment.Contract.Reissue(&_TeePayment.TransactOpts, _account, _nonce, _firstSubNonce, _paymentInstructions, _fees, _nullify)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xd302be77.
//
// Solidity: function setBatchSettings((bytes32,string) _account, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePayment *TeePaymentTransactor) SetBatchSettings(opts *bind.TransactOpts, _account ITeePaymentsPMWMultisigAccount, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "setBatchSettings", _account, _batchSize, _batchDurationSeconds)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xd302be77.
//
// Solidity: function setBatchSettings((bytes32,string) _account, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePayment *TeePaymentSession) SetBatchSettings(_account ITeePaymentsPMWMultisigAccount, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePayment.Contract.SetBatchSettings(&_TeePayment.TransactOpts, _account, _batchSize, _batchDurationSeconds)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xd302be77.
//
// Solidity: function setBatchSettings((bytes32,string) _account, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePayment *TeePaymentTransactorSession) SetBatchSettings(_account ITeePaymentsPMWMultisigAccount, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePayment.Contract.SetBatchSettings(&_TeePayment.TransactOpts, _account, _batchSize, _batchDurationSeconds)
}

// SetPaymentLimits is a paid mutator transaction binding the contract method 0x547ec094.
//
// Solidity: function setPaymentLimits((bytes32,string) _account, uint256 _transactionLimit, uint256 _dailyLimit) payable returns()
func (_TeePayment *TeePaymentTransactor) SetPaymentLimits(opts *bind.TransactOpts, _account ITeePaymentsPMWMultisigAccount, _transactionLimit *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "setPaymentLimits", _account, _transactionLimit, _dailyLimit)
}

// SetPaymentLimits is a paid mutator transaction binding the contract method 0x547ec094.
//
// Solidity: function setPaymentLimits((bytes32,string) _account, uint256 _transactionLimit, uint256 _dailyLimit) payable returns()
func (_TeePayment *TeePaymentSession) SetPaymentLimits(_account ITeePaymentsPMWMultisigAccount, _transactionLimit *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _TeePayment.Contract.SetPaymentLimits(&_TeePayment.TransactOpts, _account, _transactionLimit, _dailyLimit)
}

// SetPaymentLimits is a paid mutator transaction binding the contract method 0x547ec094.
//
// Solidity: function setPaymentLimits((bytes32,string) _account, uint256 _transactionLimit, uint256 _dailyLimit) payable returns()
func (_TeePayment *TeePaymentTransactorSession) SetPaymentLimits(_account ITeePaymentsPMWMultisigAccount, _transactionLimit *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _TeePayment.Contract.SetPaymentLimits(&_TeePayment.TransactOpts, _account, _transactionLimit, _dailyLimit)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePayment *TeePaymentTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePayment *TeePaymentSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePayment.Contract.SwitchToProductionMode(&_TeePayment.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePayment *TeePaymentTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePayment.Contract.SwitchToProductionMode(&_TeePayment.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePayment *TeePaymentTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePayment *TeePaymentSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePayment.Contract.UpdateContractAddresses(&_TeePayment.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePayment *TeePaymentTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePayment.Contract.UpdateContractAddresses(&_TeePayment.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePayment *TeePaymentTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePayment *TeePaymentSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePayment.Contract.UpgradeToAndCall(&_TeePayment.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePayment *TeePaymentTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePayment.Contract.UpgradeToAndCall(&_TeePayment.TransactOpts, _newImplementation, _data)
}

// TeePaymentBatchSettingsSetIterator is returned from FilterBatchSettingsSet and is used to iterate over the raw logs and unpacked data for BatchSettingsSet events raised by the TeePayment contract.
type TeePaymentBatchSettingsSetIterator struct {
	Event *TeePaymentBatchSettingsSet // Event containing the contract specifics and raw log

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
func (it *TeePaymentBatchSettingsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentBatchSettingsSet)
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
		it.Event = new(TeePaymentBatchSettingsSet)
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
func (it *TeePaymentBatchSettingsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentBatchSettingsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentBatchSettingsSet represents a BatchSettingsSet event raised by the TeePayment contract.
type TeePaymentBatchSettingsSet struct {
	WalletId             [32]byte
	SourceId             [32]byte
	AccountAddress       string
	BatchSize            uint64
	BatchDurationSeconds uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBatchSettingsSet is a free log retrieval operation binding the contract event 0x3d8272157336eb8d7637921596918b24705cbf5dd80e65f14bc3f36c6a1d2466.
//
// Solidity: event BatchSettingsSet(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeePayment *TeePaymentFilterer) FilterBatchSettingsSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeePaymentBatchSettingsSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "BatchSettingsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentBatchSettingsSetIterator{contract: _TeePayment.contract, event: "BatchSettingsSet", logs: logs, sub: sub}, nil
}

// WatchBatchSettingsSet is a free log subscription operation binding the contract event 0x3d8272157336eb8d7637921596918b24705cbf5dd80e65f14bc3f36c6a1d2466.
//
// Solidity: event BatchSettingsSet(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeePayment *TeePaymentFilterer) WatchBatchSettingsSet(opts *bind.WatchOpts, sink chan<- *TeePaymentBatchSettingsSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "BatchSettingsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentBatchSettingsSet)
				if err := _TeePayment.contract.UnpackLog(event, "BatchSettingsSet", log); err != nil {
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

// ParseBatchSettingsSet is a log parse operation binding the contract event 0x3d8272157336eb8d7637921596918b24705cbf5dd80e65f14bc3f36c6a1d2466.
//
// Solidity: event BatchSettingsSet(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeePayment *TeePaymentFilterer) ParseBatchSettingsSet(log types.Log) (*TeePaymentBatchSettingsSet, error) {
	event := new(TeePaymentBatchSettingsSet)
	if err := _TeePayment.contract.UnpackLog(event, "BatchSettingsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeePayment contract.
type TeePaymentGovernanceCallTimelockedIterator struct {
	Event *TeePaymentGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeePaymentGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentGovernanceCallTimelocked)
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
		it.Event = new(TeePaymentGovernanceCallTimelocked)
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
func (it *TeePaymentGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePayment contract.
type TeePaymentGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePayment *TeePaymentFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeePaymentGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeePaymentGovernanceCallTimelockedIterator{contract: _TeePayment.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePayment *TeePaymentFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeePaymentGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentGovernanceCallTimelocked)
				if err := _TeePayment.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeePayment *TeePaymentFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeePaymentGovernanceCallTimelocked, error) {
	event := new(TeePaymentGovernanceCallTimelocked)
	if err := _TeePayment.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeePayment contract.
type TeePaymentGovernanceInitialisedIterator struct {
	Event *TeePaymentGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeePaymentGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentGovernanceInitialised)
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
		it.Event = new(TeePaymentGovernanceInitialised)
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
func (it *TeePaymentGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentGovernanceInitialised represents a GovernanceInitialised event raised by the TeePayment contract.
type TeePaymentGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePayment *TeePaymentFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeePaymentGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeePaymentGovernanceInitialisedIterator{contract: _TeePayment.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePayment *TeePaymentFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeePaymentGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentGovernanceInitialised)
				if err := _TeePayment.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeePayment *TeePaymentFilterer) ParseGovernanceInitialised(log types.Log) (*TeePaymentGovernanceInitialised, error) {
	event := new(TeePaymentGovernanceInitialised)
	if err := _TeePayment.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeePayment contract.
type TeePaymentGovernedProductionModeEnteredIterator struct {
	Event *TeePaymentGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeePaymentGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentGovernedProductionModeEntered)
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
		it.Event = new(TeePaymentGovernedProductionModeEntered)
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
func (it *TeePaymentGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePayment contract.
type TeePaymentGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePayment *TeePaymentFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeePaymentGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeePaymentGovernedProductionModeEnteredIterator{contract: _TeePayment.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePayment *TeePaymentFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeePaymentGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentGovernedProductionModeEntered)
				if err := _TeePayment.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeePayment *TeePaymentFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeePaymentGovernedProductionModeEntered, error) {
	event := new(TeePaymentGovernedProductionModeEntered)
	if err := _TeePayment.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentPMWMultisigAccountAddedIterator is returned from FilterPMWMultisigAccountAdded and is used to iterate over the raw logs and unpacked data for PMWMultisigAccountAdded events raised by the TeePayment contract.
type TeePaymentPMWMultisigAccountAddedIterator struct {
	Event *TeePaymentPMWMultisigAccountAdded // Event containing the contract specifics and raw log

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
func (it *TeePaymentPMWMultisigAccountAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentPMWMultisigAccountAdded)
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
		it.Event = new(TeePaymentPMWMultisigAccountAdded)
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
func (it *TeePaymentPMWMultisigAccountAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentPMWMultisigAccountAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentPMWMultisigAccountAdded represents a PMWMultisigAccountAdded event raised by the TeePayment contract.
type TeePaymentPMWMultisigAccountAdded struct {
	WalletId       [32]byte
	SourceId       [32]byte
	AccountAddress string
	InitialNonce   uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPMWMultisigAccountAdded is a free log retrieval operation binding the contract event 0x1a154754ef3a3e481040abeec2f919b1ed69d93cb469ea99d50261eae497cadd.
//
// Solidity: event PMWMultisigAccountAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 initialNonce)
func (_TeePayment *TeePaymentFilterer) FilterPMWMultisigAccountAdded(opts *bind.FilterOpts, walletId [][32]byte) (*TeePaymentPMWMultisigAccountAddedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "PMWMultisigAccountAdded", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentPMWMultisigAccountAddedIterator{contract: _TeePayment.contract, event: "PMWMultisigAccountAdded", logs: logs, sub: sub}, nil
}

// WatchPMWMultisigAccountAdded is a free log subscription operation binding the contract event 0x1a154754ef3a3e481040abeec2f919b1ed69d93cb469ea99d50261eae497cadd.
//
// Solidity: event PMWMultisigAccountAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 initialNonce)
func (_TeePayment *TeePaymentFilterer) WatchPMWMultisigAccountAdded(opts *bind.WatchOpts, sink chan<- *TeePaymentPMWMultisigAccountAdded, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "PMWMultisigAccountAdded", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentPMWMultisigAccountAdded)
				if err := _TeePayment.contract.UnpackLog(event, "PMWMultisigAccountAdded", log); err != nil {
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

// ParsePMWMultisigAccountAdded is a log parse operation binding the contract event 0x1a154754ef3a3e481040abeec2f919b1ed69d93cb469ea99d50261eae497cadd.
//
// Solidity: event PMWMultisigAccountAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 initialNonce)
func (_TeePayment *TeePaymentFilterer) ParsePMWMultisigAccountAdded(log types.Log) (*TeePaymentPMWMultisigAccountAdded, error) {
	event := new(TeePaymentPMWMultisigAccountAdded)
	if err := _TeePayment.contract.UnpackLog(event, "PMWMultisigAccountAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentSupportedSourceIdAddedIterator is returned from FilterSupportedSourceIdAdded and is used to iterate over the raw logs and unpacked data for SupportedSourceIdAdded events raised by the TeePayment contract.
type TeePaymentSupportedSourceIdAddedIterator struct {
	Event *TeePaymentSupportedSourceIdAdded // Event containing the contract specifics and raw log

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
func (it *TeePaymentSupportedSourceIdAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentSupportedSourceIdAdded)
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
		it.Event = new(TeePaymentSupportedSourceIdAdded)
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
func (it *TeePaymentSupportedSourceIdAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentSupportedSourceIdAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentSupportedSourceIdAdded represents a SupportedSourceIdAdded event raised by the TeePayment contract.
type TeePaymentSupportedSourceIdAdded struct {
	OpType   [32]byte
	SourceId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSupportedSourceIdAdded is a free log retrieval operation binding the contract event 0x11562ae33c7b3979b6a1851b0f2d8a4d41e35a41dd9a2e021b56700aa01894ff.
//
// Solidity: event SupportedSourceIdAdded(bytes32 indexed opType, bytes32 indexed sourceId)
func (_TeePayment *TeePaymentFilterer) FilterSupportedSourceIdAdded(opts *bind.FilterOpts, opType [][32]byte, sourceId [][32]byte) (*TeePaymentSupportedSourceIdAddedIterator, error) {

	var opTypeRule []interface{}
	for _, opTypeItem := range opType {
		opTypeRule = append(opTypeRule, opTypeItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "SupportedSourceIdAdded", opTypeRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentSupportedSourceIdAddedIterator{contract: _TeePayment.contract, event: "SupportedSourceIdAdded", logs: logs, sub: sub}, nil
}

// WatchSupportedSourceIdAdded is a free log subscription operation binding the contract event 0x11562ae33c7b3979b6a1851b0f2d8a4d41e35a41dd9a2e021b56700aa01894ff.
//
// Solidity: event SupportedSourceIdAdded(bytes32 indexed opType, bytes32 indexed sourceId)
func (_TeePayment *TeePaymentFilterer) WatchSupportedSourceIdAdded(opts *bind.WatchOpts, sink chan<- *TeePaymentSupportedSourceIdAdded, opType [][32]byte, sourceId [][32]byte) (event.Subscription, error) {

	var opTypeRule []interface{}
	for _, opTypeItem := range opType {
		opTypeRule = append(opTypeRule, opTypeItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "SupportedSourceIdAdded", opTypeRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentSupportedSourceIdAdded)
				if err := _TeePayment.contract.UnpackLog(event, "SupportedSourceIdAdded", log); err != nil {
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

// ParseSupportedSourceIdAdded is a log parse operation binding the contract event 0x11562ae33c7b3979b6a1851b0f2d8a4d41e35a41dd9a2e021b56700aa01894ff.
//
// Solidity: event SupportedSourceIdAdded(bytes32 indexed opType, bytes32 indexed sourceId)
func (_TeePayment *TeePaymentFilterer) ParseSupportedSourceIdAdded(log types.Log) (*TeePaymentSupportedSourceIdAdded, error) {
	event := new(TeePaymentSupportedSourceIdAdded)
	if err := _TeePayment.contract.UnpackLog(event, "SupportedSourceIdAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeePayment contract.
type TeePaymentTimelockedGovernanceCallCanceledIterator struct {
	Event *TeePaymentTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeePaymentTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeePaymentTimelockedGovernanceCallCanceled)
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
func (it *TeePaymentTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePayment contract.
type TeePaymentTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeePayment *TeePaymentFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeePaymentTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeePaymentTimelockedGovernanceCallCanceledIterator{contract: _TeePayment.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeePayment *TeePaymentFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeePaymentTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentTimelockedGovernanceCallCanceled)
				if err := _TeePayment.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeePayment *TeePaymentFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeePaymentTimelockedGovernanceCallCanceled, error) {
	event := new(TeePaymentTimelockedGovernanceCallCanceled)
	if err := _TeePayment.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeePayment contract.
type TeePaymentTimelockedGovernanceCallExecutedIterator struct {
	Event *TeePaymentTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeePaymentTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeePaymentTimelockedGovernanceCallExecuted)
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
func (it *TeePaymentTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePayment contract.
type TeePaymentTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeePayment *TeePaymentFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeePaymentTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeePaymentTimelockedGovernanceCallExecutedIterator{contract: _TeePayment.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeePayment *TeePaymentFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeePaymentTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentTimelockedGovernanceCallExecuted)
				if err := _TeePayment.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeePayment *TeePaymentFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeePaymentTimelockedGovernanceCallExecuted, error) {
	event := new(TeePaymentTimelockedGovernanceCallExecuted)
	if err := _TeePayment.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeePayment contract.
type TeePaymentUpgradedIterator struct {
	Event *TeePaymentUpgraded // Event containing the contract specifics and raw log

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
func (it *TeePaymentUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentUpgraded)
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
		it.Event = new(TeePaymentUpgraded)
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
func (it *TeePaymentUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentUpgraded represents a Upgraded event raised by the TeePayment contract.
type TeePaymentUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePayment *TeePaymentFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeePaymentUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentUpgradedIterator{contract: _TeePayment.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePayment *TeePaymentFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeePaymentUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentUpgraded)
				if err := _TeePayment.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeePayment *TeePaymentFilterer) ParseUpgraded(log types.Log) (*TeePaymentUpgraded, error) {
	event := new(TeePaymentUpgraded)
	if err := _TeePayment.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
