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
	WalletAddress string
	PublicKeys    [][]byte
	Threshold     uint64
}

// IPMWMultisigAccountConfiguredResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredResponseBody struct {
	Status   uint8
	Sequence uint64
}

// ITeePaymentsPaymentInstruction is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsPaymentInstruction struct {
	RecipientAddress string
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchDurationTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotYetEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchSizeTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchSizeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DailyLimitBelowTransactionLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DefaultWalletNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeBelowMinFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxBatchSizeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinFeeNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinFeeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPaymentInstructions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySubmitAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OpTypeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SourceIdZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletAddressAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletAddressNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongOpType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongProjectId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"BatchSettingsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"minFee\",\"type\":\"uint128\"}],\"name\":\"MinFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"walletAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"initialNonce\",\"type\":\"uint64\"}],\"name\":\"WalletAddressSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PAY\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REISSUE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SET_PAYMENT_LIMITS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getBatchSettings\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getMinFee\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"_minFee\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOpType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getOpTypeConstants\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_maxBatchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_maxBatchDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBatchDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBatchSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePayments.PaymentInstruction\",\"name\":\"_paymentInstruction\",\"type\":\"tuple\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_subNonce\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_firstSubNonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePayments.PaymentInstruction[]\",\"name\":\"_paymentInstructions\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_nullify\",\"type\":\"bool[]\"}],\"name\":\"reissue\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"setBatchSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"_minFee\",\"type\":\"uint128\"}],\"name\":\"setMinFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_transactionLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"name\":\"setPaymentLimits\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"walletAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigAccountConfigured.PMWMultisigAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"setWalletAddressAndInitialNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeInstructions\",\"outputs\":[{\"internalType\":\"contractITeeInstructions\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeVerification\",\"outputs\":[{\"internalType\":\"contractITeeVerification\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletKeyManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletKeyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletProjectManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletProjectManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// GetBatchSettings is a free data retrieval call binding the contract method 0x322dd91b.
//
// Solidity: function getBatchSettings(bytes32 _walletId) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (_TeePayment *TeePaymentCaller) GetBatchSettings(opts *bind.CallOpts, _walletId [32]byte) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "getBatchSettings", _walletId)

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

// GetBatchSettings is a free data retrieval call binding the contract method 0x322dd91b.
//
// Solidity: function getBatchSettings(bytes32 _walletId) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (_TeePayment *TeePaymentSession) GetBatchSettings(_walletId [32]byte) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	return _TeePayment.Contract.GetBatchSettings(&_TeePayment.CallOpts, _walletId)
}

// GetBatchSettings is a free data retrieval call binding the contract method 0x322dd91b.
//
// Solidity: function getBatchSettings(bytes32 _walletId) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (_TeePayment *TeePaymentCallerSession) GetBatchSettings(_walletId [32]byte) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	return _TeePayment.Contract.GetBatchSettings(&_TeePayment.CallOpts, _walletId)
}

// GetMinFee is a free data retrieval call binding the contract method 0x8f937b01.
//
// Solidity: function getMinFee(bytes32 _walletId) view returns(uint128 _minFee)
func (_TeePayment *TeePaymentCaller) GetMinFee(opts *bind.CallOpts, _walletId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "getMinFee", _walletId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinFee is a free data retrieval call binding the contract method 0x8f937b01.
//
// Solidity: function getMinFee(bytes32 _walletId) view returns(uint128 _minFee)
func (_TeePayment *TeePaymentSession) GetMinFee(_walletId [32]byte) (*big.Int, error) {
	return _TeePayment.Contract.GetMinFee(&_TeePayment.CallOpts, _walletId)
}

// GetMinFee is a free data retrieval call binding the contract method 0x8f937b01.
//
// Solidity: function getMinFee(bytes32 _walletId) view returns(uint128 _minFee)
func (_TeePayment *TeePaymentCallerSession) GetMinFee(_walletId [32]byte) (*big.Int, error) {
	return _TeePayment.Contract.GetMinFee(&_TeePayment.CallOpts, _walletId)
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

// GetWalletAddress is a free data retrieval call binding the contract method 0x206d54db.
//
// Solidity: function getWalletAddress(bytes32 _walletId) view returns(string)
func (_TeePayment *TeePaymentCaller) GetWalletAddress(opts *bind.CallOpts, _walletId [32]byte) (string, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "getWalletAddress", _walletId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetWalletAddress is a free data retrieval call binding the contract method 0x206d54db.
//
// Solidity: function getWalletAddress(bytes32 _walletId) view returns(string)
func (_TeePayment *TeePaymentSession) GetWalletAddress(_walletId [32]byte) (string, error) {
	return _TeePayment.Contract.GetWalletAddress(&_TeePayment.CallOpts, _walletId)
}

// GetWalletAddress is a free data retrieval call binding the contract method 0x206d54db.
//
// Solidity: function getWalletAddress(bytes32 _walletId) view returns(string)
func (_TeePayment *TeePaymentCallerSession) GetWalletAddress(_walletId [32]byte) (string, error) {
	return _TeePayment.Contract.GetWalletAddress(&_TeePayment.CallOpts, _walletId)
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

// TeeInstructions is a free data retrieval call binding the contract method 0xa3f000ce.
//
// Solidity: function teeInstructions() view returns(address)
func (_TeePayment *TeePaymentCaller) TeeInstructions(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayment.contract.Call(opts, &out, "teeInstructions")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeInstructions is a free data retrieval call binding the contract method 0xa3f000ce.
//
// Solidity: function teeInstructions() view returns(address)
func (_TeePayment *TeePaymentSession) TeeInstructions() (common.Address, error) {
	return _TeePayment.Contract.TeeInstructions(&_TeePayment.CallOpts)
}

// TeeInstructions is a free data retrieval call binding the contract method 0xa3f000ce.
//
// Solidity: function teeInstructions() view returns(address)
func (_TeePayment *TeePaymentCallerSession) TeeInstructions() (common.Address, error) {
	return _TeePayment.Contract.TeeInstructions(&_TeePayment.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x74a5d32e.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds, bytes32 _opType, bytes32 _sourceId) returns()
func (_TeePayment *TeePaymentTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _maxBatchSize uint64, _maxBatchDurationSeconds uint64, _opType [32]byte, _sourceId [32]byte) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater, _maxBatchSize, _maxBatchDurationSeconds, _opType, _sourceId)
}

// Initialize is a paid mutator transaction binding the contract method 0x74a5d32e.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds, bytes32 _opType, bytes32 _sourceId) returns()
func (_TeePayment *TeePaymentSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _maxBatchSize uint64, _maxBatchDurationSeconds uint64, _opType [32]byte, _sourceId [32]byte) (*types.Transaction, error) {
	return _TeePayment.Contract.Initialize(&_TeePayment.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _maxBatchSize, _maxBatchDurationSeconds, _opType, _sourceId)
}

// Initialize is a paid mutator transaction binding the contract method 0x74a5d32e.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds, bytes32 _opType, bytes32 _sourceId) returns()
func (_TeePayment *TeePaymentTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _maxBatchSize uint64, _maxBatchDurationSeconds uint64, _opType [32]byte, _sourceId [32]byte) (*types.Transaction, error) {
	return _TeePayment.Contract.Initialize(&_TeePayment.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _maxBatchSize, _maxBatchDurationSeconds, _opType, _sourceId)
}

// Pay is a paid mutator transaction binding the contract method 0xe253d4c7.
//
// Solidity: function pay(bytes32 _projectId, bytes32 _walletId, (string,uint256,uint256,bytes32) _paymentInstruction) payable returns(uint64 _nonce, uint64 _subNonce)
func (_TeePayment *TeePaymentTransactor) Pay(opts *bind.TransactOpts, _projectId [32]byte, _walletId [32]byte, _paymentInstruction ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "pay", _projectId, _walletId, _paymentInstruction)
}

// Pay is a paid mutator transaction binding the contract method 0xe253d4c7.
//
// Solidity: function pay(bytes32 _projectId, bytes32 _walletId, (string,uint256,uint256,bytes32) _paymentInstruction) payable returns(uint64 _nonce, uint64 _subNonce)
func (_TeePayment *TeePaymentSession) Pay(_projectId [32]byte, _walletId [32]byte, _paymentInstruction ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _TeePayment.Contract.Pay(&_TeePayment.TransactOpts, _projectId, _walletId, _paymentInstruction)
}

// Pay is a paid mutator transaction binding the contract method 0xe253d4c7.
//
// Solidity: function pay(bytes32 _projectId, bytes32 _walletId, (string,uint256,uint256,bytes32) _paymentInstruction) payable returns(uint64 _nonce, uint64 _subNonce)
func (_TeePayment *TeePaymentTransactorSession) Pay(_projectId [32]byte, _walletId [32]byte, _paymentInstruction ITeePaymentsPaymentInstruction) (*types.Transaction, error) {
	return _TeePayment.Contract.Pay(&_TeePayment.TransactOpts, _projectId, _walletId, _paymentInstruction)
}

// Reissue is a paid mutator transaction binding the contract method 0xcac836b8.
//
// Solidity: function reissue(bytes32 _walletId, uint64 _nonce, uint64 _firstSubNonce, (string,uint256,uint256,bytes32)[] _paymentInstructions, uint256[] _fees, bool[] _nullify) payable returns()
func (_TeePayment *TeePaymentTransactor) Reissue(opts *bind.TransactOpts, _walletId [32]byte, _nonce uint64, _firstSubNonce uint64, _paymentInstructions []ITeePaymentsPaymentInstruction, _fees []*big.Int, _nullify []bool) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "reissue", _walletId, _nonce, _firstSubNonce, _paymentInstructions, _fees, _nullify)
}

// Reissue is a paid mutator transaction binding the contract method 0xcac836b8.
//
// Solidity: function reissue(bytes32 _walletId, uint64 _nonce, uint64 _firstSubNonce, (string,uint256,uint256,bytes32)[] _paymentInstructions, uint256[] _fees, bool[] _nullify) payable returns()
func (_TeePayment *TeePaymentSession) Reissue(_walletId [32]byte, _nonce uint64, _firstSubNonce uint64, _paymentInstructions []ITeePaymentsPaymentInstruction, _fees []*big.Int, _nullify []bool) (*types.Transaction, error) {
	return _TeePayment.Contract.Reissue(&_TeePayment.TransactOpts, _walletId, _nonce, _firstSubNonce, _paymentInstructions, _fees, _nullify)
}

// Reissue is a paid mutator transaction binding the contract method 0xcac836b8.
//
// Solidity: function reissue(bytes32 _walletId, uint64 _nonce, uint64 _firstSubNonce, (string,uint256,uint256,bytes32)[] _paymentInstructions, uint256[] _fees, bool[] _nullify) payable returns()
func (_TeePayment *TeePaymentTransactorSession) Reissue(_walletId [32]byte, _nonce uint64, _firstSubNonce uint64, _paymentInstructions []ITeePaymentsPaymentInstruction, _fees []*big.Int, _nullify []bool) (*types.Transaction, error) {
	return _TeePayment.Contract.Reissue(&_TeePayment.TransactOpts, _walletId, _nonce, _firstSubNonce, _paymentInstructions, _fees, _nullify)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xf680cbad.
//
// Solidity: function setBatchSettings(bytes32 _walletId, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePayment *TeePaymentTransactor) SetBatchSettings(opts *bind.TransactOpts, _walletId [32]byte, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "setBatchSettings", _walletId, _batchSize, _batchDurationSeconds)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xf680cbad.
//
// Solidity: function setBatchSettings(bytes32 _walletId, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePayment *TeePaymentSession) SetBatchSettings(_walletId [32]byte, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePayment.Contract.SetBatchSettings(&_TeePayment.TransactOpts, _walletId, _batchSize, _batchDurationSeconds)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xf680cbad.
//
// Solidity: function setBatchSettings(bytes32 _walletId, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePayment *TeePaymentTransactorSession) SetBatchSettings(_walletId [32]byte, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePayment.Contract.SetBatchSettings(&_TeePayment.TransactOpts, _walletId, _batchSize, _batchDurationSeconds)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x88da8aa6.
//
// Solidity: function setMinFee(bytes32 _walletId, uint128 _minFee) returns()
func (_TeePayment *TeePaymentTransactor) SetMinFee(opts *bind.TransactOpts, _walletId [32]byte, _minFee *big.Int) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "setMinFee", _walletId, _minFee)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x88da8aa6.
//
// Solidity: function setMinFee(bytes32 _walletId, uint128 _minFee) returns()
func (_TeePayment *TeePaymentSession) SetMinFee(_walletId [32]byte, _minFee *big.Int) (*types.Transaction, error) {
	return _TeePayment.Contract.SetMinFee(&_TeePayment.TransactOpts, _walletId, _minFee)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x88da8aa6.
//
// Solidity: function setMinFee(bytes32 _walletId, uint128 _minFee) returns()
func (_TeePayment *TeePaymentTransactorSession) SetMinFee(_walletId [32]byte, _minFee *big.Int) (*types.Transaction, error) {
	return _TeePayment.Contract.SetMinFee(&_TeePayment.TransactOpts, _walletId, _minFee)
}

// SetPaymentLimits is a paid mutator transaction binding the contract method 0xb01f700c.
//
// Solidity: function setPaymentLimits(bytes32 _walletId, uint256 _transactionLimit, uint256 _dailyLimit) payable returns()
func (_TeePayment *TeePaymentTransactor) SetPaymentLimits(opts *bind.TransactOpts, _walletId [32]byte, _transactionLimit *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "setPaymentLimits", _walletId, _transactionLimit, _dailyLimit)
}

// SetPaymentLimits is a paid mutator transaction binding the contract method 0xb01f700c.
//
// Solidity: function setPaymentLimits(bytes32 _walletId, uint256 _transactionLimit, uint256 _dailyLimit) payable returns()
func (_TeePayment *TeePaymentSession) SetPaymentLimits(_walletId [32]byte, _transactionLimit *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _TeePayment.Contract.SetPaymentLimits(&_TeePayment.TransactOpts, _walletId, _transactionLimit, _dailyLimit)
}

// SetPaymentLimits is a paid mutator transaction binding the contract method 0xb01f700c.
//
// Solidity: function setPaymentLimits(bytes32 _walletId, uint256 _transactionLimit, uint256 _dailyLimit) payable returns()
func (_TeePayment *TeePaymentTransactorSession) SetPaymentLimits(_walletId [32]byte, _transactionLimit *big.Int, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _TeePayment.Contract.SetPaymentLimits(&_TeePayment.TransactOpts, _walletId, _transactionLimit, _dailyLimit)
}

// SetWalletAddressAndInitialNonce is a paid mutator transaction binding the contract method 0x4874b439.
//
// Solidity: function setWalletAddressAndInitialNonce(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns()
func (_TeePayment *TeePaymentTransactor) SetWalletAddressAndInitialNonce(opts *bind.TransactOpts, _walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeePayment.contract.Transact(opts, "setWalletAddressAndInitialNonce", _walletId, _proof)
}

// SetWalletAddressAndInitialNonce is a paid mutator transaction binding the contract method 0x4874b439.
//
// Solidity: function setWalletAddressAndInitialNonce(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns()
func (_TeePayment *TeePaymentSession) SetWalletAddressAndInitialNonce(_walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeePayment.Contract.SetWalletAddressAndInitialNonce(&_TeePayment.TransactOpts, _walletId, _proof)
}

// SetWalletAddressAndInitialNonce is a paid mutator transaction binding the contract method 0x4874b439.
//
// Solidity: function setWalletAddressAndInitialNonce(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns()
func (_TeePayment *TeePaymentTransactorSession) SetWalletAddressAndInitialNonce(_walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeePayment.Contract.SetWalletAddressAndInitialNonce(&_TeePayment.TransactOpts, _walletId, _proof)
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
	BatchSize            uint64
	BatchDurationSeconds uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBatchSettingsSet is a free log retrieval operation binding the contract event 0xfa2da2b02c35a83b9a69ec27c4d6fbf1c4cf31d416e5783d45a39bdfbc36deff.
//
// Solidity: event BatchSettingsSet(bytes32 indexed walletId, uint64 batchSize, uint64 batchDurationSeconds)
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

// WatchBatchSettingsSet is a free log subscription operation binding the contract event 0xfa2da2b02c35a83b9a69ec27c4d6fbf1c4cf31d416e5783d45a39bdfbc36deff.
//
// Solidity: event BatchSettingsSet(bytes32 indexed walletId, uint64 batchSize, uint64 batchDurationSeconds)
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

// ParseBatchSettingsSet is a log parse operation binding the contract event 0xfa2da2b02c35a83b9a69ec27c4d6fbf1c4cf31d416e5783d45a39bdfbc36deff.
//
// Solidity: event BatchSettingsSet(bytes32 indexed walletId, uint64 batchSize, uint64 batchDurationSeconds)
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

// TeePaymentMinFeeSetIterator is returned from FilterMinFeeSet and is used to iterate over the raw logs and unpacked data for MinFeeSet events raised by the TeePayment contract.
type TeePaymentMinFeeSetIterator struct {
	Event *TeePaymentMinFeeSet // Event containing the contract specifics and raw log

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
func (it *TeePaymentMinFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentMinFeeSet)
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
		it.Event = new(TeePaymentMinFeeSet)
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
func (it *TeePaymentMinFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentMinFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentMinFeeSet represents a MinFeeSet event raised by the TeePayment contract.
type TeePaymentMinFeeSet struct {
	WalletId [32]byte
	MinFee   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinFeeSet is a free log retrieval operation binding the contract event 0xc37f37cead8822d50df03488bfae0fcd52d892c3d18eb730441eccbedcddd66f.
//
// Solidity: event MinFeeSet(bytes32 indexed walletId, uint128 minFee)
func (_TeePayment *TeePaymentFilterer) FilterMinFeeSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeePaymentMinFeeSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "MinFeeSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentMinFeeSetIterator{contract: _TeePayment.contract, event: "MinFeeSet", logs: logs, sub: sub}, nil
}

// WatchMinFeeSet is a free log subscription operation binding the contract event 0xc37f37cead8822d50df03488bfae0fcd52d892c3d18eb730441eccbedcddd66f.
//
// Solidity: event MinFeeSet(bytes32 indexed walletId, uint128 minFee)
func (_TeePayment *TeePaymentFilterer) WatchMinFeeSet(opts *bind.WatchOpts, sink chan<- *TeePaymentMinFeeSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "MinFeeSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentMinFeeSet)
				if err := _TeePayment.contract.UnpackLog(event, "MinFeeSet", log); err != nil {
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

// ParseMinFeeSet is a log parse operation binding the contract event 0xc37f37cead8822d50df03488bfae0fcd52d892c3d18eb730441eccbedcddd66f.
//
// Solidity: event MinFeeSet(bytes32 indexed walletId, uint128 minFee)
func (_TeePayment *TeePaymentFilterer) ParseMinFeeSet(log types.Log) (*TeePaymentMinFeeSet, error) {
	event := new(TeePaymentMinFeeSet)
	if err := _TeePayment.contract.UnpackLog(event, "MinFeeSet", log); err != nil {
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

// TeePaymentWalletAddressSetIterator is returned from FilterWalletAddressSet and is used to iterate over the raw logs and unpacked data for WalletAddressSet events raised by the TeePayment contract.
type TeePaymentWalletAddressSetIterator struct {
	Event *TeePaymentWalletAddressSet // Event containing the contract specifics and raw log

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
func (it *TeePaymentWalletAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentWalletAddressSet)
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
		it.Event = new(TeePaymentWalletAddressSet)
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
func (it *TeePaymentWalletAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentWalletAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentWalletAddressSet represents a WalletAddressSet event raised by the TeePayment contract.
type TeePaymentWalletAddressSet struct {
	WalletId      [32]byte
	WalletAddress string
	InitialNonce  uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWalletAddressSet is a free log retrieval operation binding the contract event 0xdd4418741bb390c3c70db3b6a6af62f3befd21c0cbc950b75d2799132691594c.
//
// Solidity: event WalletAddressSet(bytes32 indexed walletId, string walletAddress, uint64 initialNonce)
func (_TeePayment *TeePaymentFilterer) FilterWalletAddressSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeePaymentWalletAddressSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.FilterLogs(opts, "WalletAddressSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentWalletAddressSetIterator{contract: _TeePayment.contract, event: "WalletAddressSet", logs: logs, sub: sub}, nil
}

// WatchWalletAddressSet is a free log subscription operation binding the contract event 0xdd4418741bb390c3c70db3b6a6af62f3befd21c0cbc950b75d2799132691594c.
//
// Solidity: event WalletAddressSet(bytes32 indexed walletId, string walletAddress, uint64 initialNonce)
func (_TeePayment *TeePaymentFilterer) WatchWalletAddressSet(opts *bind.WatchOpts, sink chan<- *TeePaymentWalletAddressSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayment.contract.WatchLogs(opts, "WalletAddressSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentWalletAddressSet)
				if err := _TeePayment.contract.UnpackLog(event, "WalletAddressSet", log); err != nil {
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

// ParseWalletAddressSet is a log parse operation binding the contract event 0xdd4418741bb390c3c70db3b6a6af62f3befd21c0cbc950b75d2799132691594c.
//
// Solidity: event WalletAddressSet(bytes32 indexed walletId, string walletAddress, uint64 initialNonce)
func (_TeePayment *TeePaymentFilterer) ParseWalletAddressSet(log types.Log) (*TeePaymentWalletAddressSet, error) {
	event := new(TeePaymentWalletAddressSet)
	if err := _TeePayment.contract.UnpackLog(event, "WalletAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
