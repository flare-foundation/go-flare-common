// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepayments

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

// IFdc2HubFdc2ResponseHeader is an auto generated low-level Go binding around an user-defined struct.
type IFdc2HubFdc2ResponseHeader struct {
	AttestationType    [32]byte
	SourceId           [32]byte
	ThresholdBIPS      uint16
	ProofOwner         common.Address
	Cosigners          []common.Address
	CosignersThreshold uint64
	Timestamp          uint64
}

// IFdc2VerificationFdc2Signatures is an auto generated low-level Go binding around an user-defined struct.
type IFdc2VerificationFdc2Signatures struct {
	SigningPolicySignatures []byte
	TeeSignatures           []Signature
	CosignerSignatures      []Signature
}

// IPMWMultisigAccountConfiguredProof is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
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
	TokenId          []byte
	Amount           *big.Int
	MaxFee           *big.Int
	PaymentReference [32]byte
}

// ITeePaymentsReissueFeeParams is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsReissueFeeParams struct {
	MaxFeePerPayment      []*big.Int
	FactorsBIPSPerPayment [][]int16
	DelaysSeconds         []uint16
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// TeePaymentsMetaData contains all meta data concerning the TeePayments contract.
var TeePaymentsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountAddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuthorizationAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchDurationTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotYetEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchSizeTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchSizeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyTypeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxBatchSizeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinFeeNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPaymentInstructions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuthorizationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OpTypeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMWMultisigAccountAddressAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RecipientIsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongKeyType\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"BatchSettingsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"initialNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"PMWMultisigAccountAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigAccountConfigured.PMWMultisigAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"name\":\"addPMWMultisigAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"contractIIFlareTeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAuthorizationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getBatchSettings\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKeyType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOpType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAccounts\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getWalletId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_maxBatchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_maxBatchDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBatchDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBatchSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePayments.PaymentInstruction\",\"name\":\"_paymentInstruction\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_subNonce\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_firstSubNonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePayments.PaymentInstruction[]\",\"name\":\"_paymentInstructions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"maxFeePerPayment\",\"type\":\"uint256[]\"},{\"internalType\":\"int16[][]\",\"name\":\"factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"delaysSeconds\",\"type\":\"uint16[]\"}],\"internalType\":\"structITeePayments.ReissueFeeParams\",\"name\":\"_reissueFeeParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"reissue\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePayments.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"setBatchSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsFeeScheduleManager\",\"outputs\":[{\"internalType\":\"contractITeePaymentsFeeScheduleManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsRegistry\",\"outputs\":[{\"internalType\":\"contractITeePaymentsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeePaymentsABI is the input ABI used to generate the binding from.
// Deprecated: Use TeePaymentsMetaData.ABI instead.
var TeePaymentsABI = TeePaymentsMetaData.ABI

// TeePayments is an auto generated Go binding around an Ethereum contract.
type TeePayments struct {
	TeePaymentsCaller     // Read-only binding to the contract
	TeePaymentsTransactor // Write-only binding to the contract
	TeePaymentsFilterer   // Log filterer for contract events
}

// TeePaymentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeePaymentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeePaymentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeePaymentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeePaymentsSession struct {
	Contract     *TeePayments      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeePaymentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeePaymentsCallerSession struct {
	Contract *TeePaymentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TeePaymentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeePaymentsTransactorSession struct {
	Contract     *TeePaymentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TeePaymentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeePaymentsRaw struct {
	Contract *TeePayments // Generic contract binding to access the raw methods on
}

// TeePaymentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeePaymentsCallerRaw struct {
	Contract *TeePaymentsCaller // Generic read-only contract binding to access the raw methods on
}

// TeePaymentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeePaymentsTransactorRaw struct {
	Contract *TeePaymentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeePayments creates a new instance of TeePayments, bound to a specific deployed contract.
func NewTeePayments(address common.Address, backend bind.ContractBackend) (*TeePayments, error) {
	contract, err := bindTeePayments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeePayments{TeePaymentsCaller: TeePaymentsCaller{contract: contract}, TeePaymentsTransactor: TeePaymentsTransactor{contract: contract}, TeePaymentsFilterer: TeePaymentsFilterer{contract: contract}}, nil
}

// NewTeePaymentsCaller creates a new read-only instance of TeePayments, bound to a specific deployed contract.
func NewTeePaymentsCaller(address common.Address, caller bind.ContractCaller) (*TeePaymentsCaller, error) {
	contract, err := bindTeePayments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsCaller{contract: contract}, nil
}

// NewTeePaymentsTransactor creates a new write-only instance of TeePayments, bound to a specific deployed contract.
func NewTeePaymentsTransactor(address common.Address, transactor bind.ContractTransactor) (*TeePaymentsTransactor, error) {
	contract, err := bindTeePayments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsTransactor{contract: contract}, nil
}

// NewTeePaymentsFilterer creates a new log filterer instance of TeePayments, bound to a specific deployed contract.
func NewTeePaymentsFilterer(address common.Address, filterer bind.ContractFilterer) (*TeePaymentsFilterer, error) {
	contract, err := bindTeePayments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsFilterer{contract: contract}, nil
}

// bindTeePayments binds a generic wrapper to an already deployed contract.
func bindTeePayments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeePaymentsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePayments *TeePaymentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePayments.Contract.TeePaymentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePayments *TeePaymentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePayments.Contract.TeePaymentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePayments *TeePaymentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePayments.Contract.TeePaymentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePayments *TeePaymentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePayments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePayments *TeePaymentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePayments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePayments *TeePaymentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePayments.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePayments *TeePaymentsCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePayments *TeePaymentsSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePayments.Contract.UPGRADEINTERFACEVERSION(&_TeePayments.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePayments *TeePaymentsCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePayments.Contract.UPGRADEINTERFACEVERSION(&_TeePayments.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePayments *TeePaymentsCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePayments *TeePaymentsSession) FlareSystemsManager() (common.Address, error) {
	return _TeePayments.Contract.FlareSystemsManager(&_TeePayments.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePayments *TeePaymentsCallerSession) FlareSystemsManager() (common.Address, error) {
	return _TeePayments.Contract.FlareSystemsManager(&_TeePayments.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePayments *TeePaymentsCaller) FlareTeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "flareTeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePayments *TeePaymentsSession) FlareTeeManager() (common.Address, error) {
	return _TeePayments.Contract.FlareTeeManager(&_TeePayments.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePayments *TeePaymentsCallerSession) FlareTeeManager() (common.Address, error) {
	return _TeePayments.Contract.FlareTeeManager(&_TeePayments.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePayments *TeePaymentsCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePayments *TeePaymentsSession) GetAddressUpdater() (common.Address, error) {
	return _TeePayments.Contract.GetAddressUpdater(&_TeePayments.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePayments *TeePaymentsCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeePayments.Contract.GetAddressUpdater(&_TeePayments.CallOpts)
}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (_TeePayments *TeePaymentsCaller) GetAuthorizationAddress(opts *bind.CallOpts, _account ITeePaymentsPMWMultisigAccount) (common.Address, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "getAuthorizationAddress", _account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (_TeePayments *TeePaymentsSession) GetAuthorizationAddress(_account ITeePaymentsPMWMultisigAccount) (common.Address, error) {
	return _TeePayments.Contract.GetAuthorizationAddress(&_TeePayments.CallOpts, _account)
}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (_TeePayments *TeePaymentsCallerSession) GetAuthorizationAddress(_account ITeePaymentsPMWMultisigAccount) (common.Address, error) {
	return _TeePayments.Contract.GetAuthorizationAddress(&_TeePayments.CallOpts, _account)
}

// GetBatchSettings is a free data retrieval call binding the contract method 0xb96c3c8f.
//
// Solidity: function getBatchSettings((bytes32,string) _account) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (_TeePayments *TeePaymentsCaller) GetBatchSettings(opts *bind.CallOpts, _account ITeePaymentsPMWMultisigAccount) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "getBatchSettings", _account)

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
func (_TeePayments *TeePaymentsSession) GetBatchSettings(_account ITeePaymentsPMWMultisigAccount) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	return _TeePayments.Contract.GetBatchSettings(&_TeePayments.CallOpts, _account)
}

// GetBatchSettings is a free data retrieval call binding the contract method 0xb96c3c8f.
//
// Solidity: function getBatchSettings((bytes32,string) _account) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (_TeePayments *TeePaymentsCallerSession) GetBatchSettings(_account ITeePaymentsPMWMultisigAccount) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	return _TeePayments.Contract.GetBatchSettings(&_TeePayments.CallOpts, _account)
}

// GetKeyType is a free data retrieval call binding the contract method 0xd4f0b37c.
//
// Solidity: function getKeyType() view returns(bytes32)
func (_TeePayments *TeePaymentsCaller) GetKeyType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "getKeyType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKeyType is a free data retrieval call binding the contract method 0xd4f0b37c.
//
// Solidity: function getKeyType() view returns(bytes32)
func (_TeePayments *TeePaymentsSession) GetKeyType() ([32]byte, error) {
	return _TeePayments.Contract.GetKeyType(&_TeePayments.CallOpts)
}

// GetKeyType is a free data retrieval call binding the contract method 0xd4f0b37c.
//
// Solidity: function getKeyType() view returns(bytes32)
func (_TeePayments *TeePaymentsCallerSession) GetKeyType() ([32]byte, error) {
	return _TeePayments.Contract.GetKeyType(&_TeePayments.CallOpts)
}

// GetOpType is a free data retrieval call binding the contract method 0xa3e2d39a.
//
// Solidity: function getOpType() view returns(bytes32)
func (_TeePayments *TeePaymentsCaller) GetOpType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "getOpType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOpType is a free data retrieval call binding the contract method 0xa3e2d39a.
//
// Solidity: function getOpType() view returns(bytes32)
func (_TeePayments *TeePaymentsSession) GetOpType() ([32]byte, error) {
	return _TeePayments.Contract.GetOpType(&_TeePayments.CallOpts)
}

// GetOpType is a free data retrieval call binding the contract method 0xa3e2d39a.
//
// Solidity: function getOpType() view returns(bytes32)
func (_TeePayments *TeePaymentsCallerSession) GetOpType() ([32]byte, error) {
	return _TeePayments.Contract.GetOpType(&_TeePayments.CallOpts)
}

// GetWalletAccounts is a free data retrieval call binding the contract method 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (_TeePayments *TeePaymentsCaller) GetWalletAccounts(opts *bind.CallOpts, _walletId [32]byte) ([]ITeePaymentsPMWMultisigAccount, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "getWalletAccounts", _walletId)

	if err != nil {
		return *new([]ITeePaymentsPMWMultisigAccount), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITeePaymentsPMWMultisigAccount)).(*[]ITeePaymentsPMWMultisigAccount)

	return out0, err

}

// GetWalletAccounts is a free data retrieval call binding the contract method 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (_TeePayments *TeePaymentsSession) GetWalletAccounts(_walletId [32]byte) ([]ITeePaymentsPMWMultisigAccount, error) {
	return _TeePayments.Contract.GetWalletAccounts(&_TeePayments.CallOpts, _walletId)
}

// GetWalletAccounts is a free data retrieval call binding the contract method 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (_TeePayments *TeePaymentsCallerSession) GetWalletAccounts(_walletId [32]byte) ([]ITeePaymentsPMWMultisigAccount, error) {
	return _TeePayments.Contract.GetWalletAccounts(&_TeePayments.CallOpts, _walletId)
}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (_TeePayments *TeePaymentsCaller) GetWalletId(opts *bind.CallOpts, _account ITeePaymentsPMWMultisigAccount) ([32]byte, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "getWalletId", _account)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (_TeePayments *TeePaymentsSession) GetWalletId(_account ITeePaymentsPMWMultisigAccount) ([32]byte, error) {
	return _TeePayments.Contract.GetWalletId(&_TeePayments.CallOpts, _account)
}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (_TeePayments *TeePaymentsCallerSession) GetWalletId(_account ITeePaymentsPMWMultisigAccount) ([32]byte, error) {
	return _TeePayments.Contract.GetWalletId(&_TeePayments.CallOpts, _account)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePayments *TeePaymentsCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePayments *TeePaymentsSession) Governance() (common.Address, error) {
	return _TeePayments.Contract.Governance(&_TeePayments.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePayments *TeePaymentsCallerSession) Governance() (common.Address, error) {
	return _TeePayments.Contract.Governance(&_TeePayments.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePayments *TeePaymentsCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePayments *TeePaymentsSession) GovernanceSettings() (common.Address, error) {
	return _TeePayments.Contract.GovernanceSettings(&_TeePayments.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePayments *TeePaymentsCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeePayments.Contract.GovernanceSettings(&_TeePayments.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePayments *TeePaymentsCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePayments *TeePaymentsSession) Implementation() (common.Address, error) {
	return _TeePayments.Contract.Implementation(&_TeePayments.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePayments *TeePaymentsCallerSession) Implementation() (common.Address, error) {
	return _TeePayments.Contract.Implementation(&_TeePayments.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePayments *TeePaymentsCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePayments *TeePaymentsSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePayments.Contract.IsExecutor(&_TeePayments.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePayments *TeePaymentsCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePayments.Contract.IsExecutor(&_TeePayments.CallOpts, _address)
}

// MaxBatchDurationSeconds is a free data retrieval call binding the contract method 0xfdef3a1e.
//
// Solidity: function maxBatchDurationSeconds() view returns(uint64)
func (_TeePayments *TeePaymentsCaller) MaxBatchDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "maxBatchDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxBatchDurationSeconds is a free data retrieval call binding the contract method 0xfdef3a1e.
//
// Solidity: function maxBatchDurationSeconds() view returns(uint64)
func (_TeePayments *TeePaymentsSession) MaxBatchDurationSeconds() (uint64, error) {
	return _TeePayments.Contract.MaxBatchDurationSeconds(&_TeePayments.CallOpts)
}

// MaxBatchDurationSeconds is a free data retrieval call binding the contract method 0xfdef3a1e.
//
// Solidity: function maxBatchDurationSeconds() view returns(uint64)
func (_TeePayments *TeePaymentsCallerSession) MaxBatchDurationSeconds() (uint64, error) {
	return _TeePayments.Contract.MaxBatchDurationSeconds(&_TeePayments.CallOpts)
}

// MaxBatchSize is a free data retrieval call binding the contract method 0x2913daa0.
//
// Solidity: function maxBatchSize() view returns(uint64)
func (_TeePayments *TeePaymentsCaller) MaxBatchSize(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "maxBatchSize")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaxBatchSize is a free data retrieval call binding the contract method 0x2913daa0.
//
// Solidity: function maxBatchSize() view returns(uint64)
func (_TeePayments *TeePaymentsSession) MaxBatchSize() (uint64, error) {
	return _TeePayments.Contract.MaxBatchSize(&_TeePayments.CallOpts)
}

// MaxBatchSize is a free data retrieval call binding the contract method 0x2913daa0.
//
// Solidity: function maxBatchSize() view returns(uint64)
func (_TeePayments *TeePaymentsCallerSession) MaxBatchSize() (uint64, error) {
	return _TeePayments.Contract.MaxBatchSize(&_TeePayments.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePayments *TeePaymentsCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePayments *TeePaymentsSession) ProductionMode() (bool, error) {
	return _TeePayments.Contract.ProductionMode(&_TeePayments.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePayments *TeePaymentsCallerSession) ProductionMode() (bool, error) {
	return _TeePayments.Contract.ProductionMode(&_TeePayments.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePayments *TeePaymentsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePayments *TeePaymentsSession) ProxiableUUID() ([32]byte, error) {
	return _TeePayments.Contract.ProxiableUUID(&_TeePayments.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePayments *TeePaymentsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeePayments.Contract.ProxiableUUID(&_TeePayments.CallOpts)
}

// TeePaymentsFeeScheduleManager is a free data retrieval call binding the contract method 0xfe9e33e8.
//
// Solidity: function teePaymentsFeeScheduleManager() view returns(address)
func (_TeePayments *TeePaymentsCaller) TeePaymentsFeeScheduleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "teePaymentsFeeScheduleManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeePaymentsFeeScheduleManager is a free data retrieval call binding the contract method 0xfe9e33e8.
//
// Solidity: function teePaymentsFeeScheduleManager() view returns(address)
func (_TeePayments *TeePaymentsSession) TeePaymentsFeeScheduleManager() (common.Address, error) {
	return _TeePayments.Contract.TeePaymentsFeeScheduleManager(&_TeePayments.CallOpts)
}

// TeePaymentsFeeScheduleManager is a free data retrieval call binding the contract method 0xfe9e33e8.
//
// Solidity: function teePaymentsFeeScheduleManager() view returns(address)
func (_TeePayments *TeePaymentsCallerSession) TeePaymentsFeeScheduleManager() (common.Address, error) {
	return _TeePayments.Contract.TeePaymentsFeeScheduleManager(&_TeePayments.CallOpts)
}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePayments *TeePaymentsCaller) TeePaymentsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "teePaymentsRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePayments *TeePaymentsSession) TeePaymentsRegistry() (common.Address, error) {
	return _TeePayments.Contract.TeePaymentsRegistry(&_TeePayments.CallOpts)
}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePayments *TeePaymentsCallerSession) TeePaymentsRegistry() (common.Address, error) {
	return _TeePayments.Contract.TeePaymentsRegistry(&_TeePayments.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePayments *TeePaymentsCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeePayments.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_TeePayments *TeePaymentsSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeePayments.Contract.TimelockedCalls(&_TeePayments.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePayments *TeePaymentsCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeePayments.Contract.TimelockedCalls(&_TeePayments.CallOpts, selector)
}

// AddPMWMultisigAccount is a paid mutator transaction binding the contract method 0x895d1bb4.
//
// Solidity: function addPMWMultisigAccount(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof, address _authorizationAddress) returns()
func (_TeePayments *TeePaymentsTransactor) AddPMWMultisigAccount(opts *bind.TransactOpts, _walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "addPMWMultisigAccount", _walletId, _proof, _authorizationAddress)
}

// AddPMWMultisigAccount is a paid mutator transaction binding the contract method 0x895d1bb4.
//
// Solidity: function addPMWMultisigAccount(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof, address _authorizationAddress) returns()
func (_TeePayments *TeePaymentsSession) AddPMWMultisigAccount(_walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeePayments.Contract.AddPMWMultisigAccount(&_TeePayments.TransactOpts, _walletId, _proof, _authorizationAddress)
}

// AddPMWMultisigAccount is a paid mutator transaction binding the contract method 0x895d1bb4.
//
// Solidity: function addPMWMultisigAccount(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof, address _authorizationAddress) returns()
func (_TeePayments *TeePaymentsTransactorSession) AddPMWMultisigAccount(_walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeePayments.Contract.AddPMWMultisigAccount(&_TeePayments.TransactOpts, _walletId, _proof, _authorizationAddress)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePayments *TeePaymentsTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePayments *TeePaymentsSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePayments.Contract.CancelGovernanceCall(&_TeePayments.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeePayments *TeePaymentsTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePayments.Contract.CancelGovernanceCall(&_TeePayments.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePayments *TeePaymentsTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePayments *TeePaymentsSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePayments.Contract.ExecuteGovernanceCall(&_TeePayments.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeePayments *TeePaymentsTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeePayments.Contract.ExecuteGovernanceCall(&_TeePayments.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePayments *TeePaymentsTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePayments *TeePaymentsSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePayments.Contract.Initialise(&_TeePayments.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeePayments *TeePaymentsTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeePayments.Contract.Initialise(&_TeePayments.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x74a5d32e.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds, bytes32 _opType, bytes32 _keyType) returns()
func (_TeePayments *TeePaymentsTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _maxBatchSize uint64, _maxBatchDurationSeconds uint64, _opType [32]byte, _keyType [32]byte) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater, _maxBatchSize, _maxBatchDurationSeconds, _opType, _keyType)
}

// Initialize is a paid mutator transaction binding the contract method 0x74a5d32e.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds, bytes32 _opType, bytes32 _keyType) returns()
func (_TeePayments *TeePaymentsSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _maxBatchSize uint64, _maxBatchDurationSeconds uint64, _opType [32]byte, _keyType [32]byte) (*types.Transaction, error) {
	return _TeePayments.Contract.Initialize(&_TeePayments.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _maxBatchSize, _maxBatchDurationSeconds, _opType, _keyType)
}

// Initialize is a paid mutator transaction binding the contract method 0x74a5d32e.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds, bytes32 _opType, bytes32 _keyType) returns()
func (_TeePayments *TeePaymentsTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _maxBatchSize uint64, _maxBatchDurationSeconds uint64, _opType [32]byte, _keyType [32]byte) (*types.Transaction, error) {
	return _TeePayments.Contract.Initialize(&_TeePayments.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _maxBatchSize, _maxBatchDurationSeconds, _opType, _keyType)
}

// Pay is a paid mutator transaction binding the contract method 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _nonce, uint64 _subNonce)
func (_TeePayments *TeePaymentsTransactor) Pay(opts *bind.TransactOpts, _account ITeePaymentsPMWMultisigAccount, _paymentInstruction ITeePaymentsPaymentInstruction, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "pay", _account, _paymentInstruction, _claimBackAddress)
}

// Pay is a paid mutator transaction binding the contract method 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _nonce, uint64 _subNonce)
func (_TeePayments *TeePaymentsSession) Pay(_account ITeePaymentsPMWMultisigAccount, _paymentInstruction ITeePaymentsPaymentInstruction, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePayments.Contract.Pay(&_TeePayments.TransactOpts, _account, _paymentInstruction, _claimBackAddress)
}

// Pay is a paid mutator transaction binding the contract method 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _nonce, uint64 _subNonce)
func (_TeePayments *TeePaymentsTransactorSession) Pay(_account ITeePaymentsPMWMultisigAccount, _paymentInstruction ITeePaymentsPaymentInstruction, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePayments.Contract.Pay(&_TeePayments.TransactOpts, _account, _paymentInstruction, _claimBackAddress)
}

// Reissue is a paid mutator transaction binding the contract method 0xb250c225.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _nonce, uint64 _firstSubNonce, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns()
func (_TeePayments *TeePaymentsTransactor) Reissue(opts *bind.TransactOpts, _account ITeePaymentsPMWMultisigAccount, _nonce uint64, _firstSubNonce uint64, _paymentInstructions []ITeePaymentsPaymentInstruction, _reissueFeeParams ITeePaymentsReissueFeeParams, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "reissue", _account, _nonce, _firstSubNonce, _paymentInstructions, _reissueFeeParams, _claimBackAddress)
}

// Reissue is a paid mutator transaction binding the contract method 0xb250c225.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _nonce, uint64 _firstSubNonce, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns()
func (_TeePayments *TeePaymentsSession) Reissue(_account ITeePaymentsPMWMultisigAccount, _nonce uint64, _firstSubNonce uint64, _paymentInstructions []ITeePaymentsPaymentInstruction, _reissueFeeParams ITeePaymentsReissueFeeParams, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePayments.Contract.Reissue(&_TeePayments.TransactOpts, _account, _nonce, _firstSubNonce, _paymentInstructions, _reissueFeeParams, _claimBackAddress)
}

// Reissue is a paid mutator transaction binding the contract method 0xb250c225.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _nonce, uint64 _firstSubNonce, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns()
func (_TeePayments *TeePaymentsTransactorSession) Reissue(_account ITeePaymentsPMWMultisigAccount, _nonce uint64, _firstSubNonce uint64, _paymentInstructions []ITeePaymentsPaymentInstruction, _reissueFeeParams ITeePaymentsReissueFeeParams, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePayments.Contract.Reissue(&_TeePayments.TransactOpts, _account, _nonce, _firstSubNonce, _paymentInstructions, _reissueFeeParams, _claimBackAddress)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xd302be77.
//
// Solidity: function setBatchSettings((bytes32,string) _account, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePayments *TeePaymentsTransactor) SetBatchSettings(opts *bind.TransactOpts, _account ITeePaymentsPMWMultisigAccount, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "setBatchSettings", _account, _batchSize, _batchDurationSeconds)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xd302be77.
//
// Solidity: function setBatchSettings((bytes32,string) _account, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePayments *TeePaymentsSession) SetBatchSettings(_account ITeePaymentsPMWMultisigAccount, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePayments.Contract.SetBatchSettings(&_TeePayments.TransactOpts, _account, _batchSize, _batchDurationSeconds)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xd302be77.
//
// Solidity: function setBatchSettings((bytes32,string) _account, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePayments *TeePaymentsTransactorSession) SetBatchSettings(_account ITeePaymentsPMWMultisigAccount, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePayments.Contract.SetBatchSettings(&_TeePayments.TransactOpts, _account, _batchSize, _batchDurationSeconds)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePayments *TeePaymentsTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePayments *TeePaymentsSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePayments.Contract.SwitchToProductionMode(&_TeePayments.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePayments *TeePaymentsTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePayments.Contract.SwitchToProductionMode(&_TeePayments.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePayments *TeePaymentsTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePayments *TeePaymentsSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePayments.Contract.UpdateContractAddresses(&_TeePayments.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePayments *TeePaymentsTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePayments.Contract.UpdateContractAddresses(&_TeePayments.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePayments *TeePaymentsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePayments.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePayments *TeePaymentsSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePayments.Contract.UpgradeToAndCall(&_TeePayments.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePayments *TeePaymentsTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePayments.Contract.UpgradeToAndCall(&_TeePayments.TransactOpts, _newImplementation, _data)
}

// TeePaymentsBatchSettingsSetIterator is returned from FilterBatchSettingsSet and is used to iterate over the raw logs and unpacked data for BatchSettingsSet events raised by the TeePayments contract.
type TeePaymentsBatchSettingsSetIterator struct {
	Event *TeePaymentsBatchSettingsSet // Event containing the contract specifics and raw log

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
func (it *TeePaymentsBatchSettingsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsBatchSettingsSet)
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
		it.Event = new(TeePaymentsBatchSettingsSet)
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
func (it *TeePaymentsBatchSettingsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsBatchSettingsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsBatchSettingsSet represents a BatchSettingsSet event raised by the TeePayments contract.
type TeePaymentsBatchSettingsSet struct {
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
func (_TeePayments *TeePaymentsFilterer) FilterBatchSettingsSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeePaymentsBatchSettingsSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayments.contract.FilterLogs(opts, "BatchSettingsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsBatchSettingsSetIterator{contract: _TeePayments.contract, event: "BatchSettingsSet", logs: logs, sub: sub}, nil
}

// WatchBatchSettingsSet is a free log subscription operation binding the contract event 0x3d8272157336eb8d7637921596918b24705cbf5dd80e65f14bc3f36c6a1d2466.
//
// Solidity: event BatchSettingsSet(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeePayments *TeePaymentsFilterer) WatchBatchSettingsSet(opts *bind.WatchOpts, sink chan<- *TeePaymentsBatchSettingsSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayments.contract.WatchLogs(opts, "BatchSettingsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsBatchSettingsSet)
				if err := _TeePayments.contract.UnpackLog(event, "BatchSettingsSet", log); err != nil {
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
func (_TeePayments *TeePaymentsFilterer) ParseBatchSettingsSet(log types.Log) (*TeePaymentsBatchSettingsSet, error) {
	event := new(TeePaymentsBatchSettingsSet)
	if err := _TeePayments.contract.UnpackLog(event, "BatchSettingsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeePayments contract.
type TeePaymentsGovernanceCallTimelockedIterator struct {
	Event *TeePaymentsGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeePaymentsGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsGovernanceCallTimelocked)
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
		it.Event = new(TeePaymentsGovernanceCallTimelocked)
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
func (it *TeePaymentsGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePayments contract.
type TeePaymentsGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePayments *TeePaymentsFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeePaymentsGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeePayments.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsGovernanceCallTimelockedIterator{contract: _TeePayments.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeePayments *TeePaymentsFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeePaymentsGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeePayments.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsGovernanceCallTimelocked)
				if err := _TeePayments.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeePayments *TeePaymentsFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeePaymentsGovernanceCallTimelocked, error) {
	event := new(TeePaymentsGovernanceCallTimelocked)
	if err := _TeePayments.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeePayments contract.
type TeePaymentsGovernanceInitialisedIterator struct {
	Event *TeePaymentsGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeePaymentsGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsGovernanceInitialised)
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
		it.Event = new(TeePaymentsGovernanceInitialised)
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
func (it *TeePaymentsGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsGovernanceInitialised represents a GovernanceInitialised event raised by the TeePayments contract.
type TeePaymentsGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePayments *TeePaymentsFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeePaymentsGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeePayments.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsGovernanceInitialisedIterator{contract: _TeePayments.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePayments *TeePaymentsFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeePaymentsGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeePayments.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsGovernanceInitialised)
				if err := _TeePayments.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeePayments *TeePaymentsFilterer) ParseGovernanceInitialised(log types.Log) (*TeePaymentsGovernanceInitialised, error) {
	event := new(TeePaymentsGovernanceInitialised)
	if err := _TeePayments.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeePayments contract.
type TeePaymentsGovernedProductionModeEnteredIterator struct {
	Event *TeePaymentsGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeePaymentsGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsGovernedProductionModeEntered)
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
		it.Event = new(TeePaymentsGovernedProductionModeEntered)
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
func (it *TeePaymentsGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePayments contract.
type TeePaymentsGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePayments *TeePaymentsFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeePaymentsGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeePayments.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsGovernedProductionModeEnteredIterator{contract: _TeePayments.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePayments *TeePaymentsFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeePaymentsGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeePayments.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsGovernedProductionModeEntered)
				if err := _TeePayments.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeePayments *TeePaymentsFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeePaymentsGovernedProductionModeEntered, error) {
	event := new(TeePaymentsGovernedProductionModeEntered)
	if err := _TeePayments.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsPMWMultisigAccountAddedIterator is returned from FilterPMWMultisigAccountAdded and is used to iterate over the raw logs and unpacked data for PMWMultisigAccountAdded events raised by the TeePayments contract.
type TeePaymentsPMWMultisigAccountAddedIterator struct {
	Event *TeePaymentsPMWMultisigAccountAdded // Event containing the contract specifics and raw log

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
func (it *TeePaymentsPMWMultisigAccountAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsPMWMultisigAccountAdded)
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
		it.Event = new(TeePaymentsPMWMultisigAccountAdded)
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
func (it *TeePaymentsPMWMultisigAccountAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsPMWMultisigAccountAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsPMWMultisigAccountAdded represents a PMWMultisigAccountAdded event raised by the TeePayments contract.
type TeePaymentsPMWMultisigAccountAdded struct {
	WalletId             [32]byte
	SourceId             [32]byte
	AccountAddress       string
	InitialNonce         uint64
	AuthorizationAddress common.Address
	BatchSize            uint64
	BatchDurationSeconds uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPMWMultisigAccountAdded is a free log retrieval operation binding the contract event 0xa43e4d705923e9d5e070e31f77dc6e19098c055be0bc847f73cc082f175511a5.
//
// Solidity: event PMWMultisigAccountAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 initialNonce, address authorizationAddress, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeePayments *TeePaymentsFilterer) FilterPMWMultisigAccountAdded(opts *bind.FilterOpts, walletId [][32]byte) (*TeePaymentsPMWMultisigAccountAddedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayments.contract.FilterLogs(opts, "PMWMultisigAccountAdded", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsPMWMultisigAccountAddedIterator{contract: _TeePayments.contract, event: "PMWMultisigAccountAdded", logs: logs, sub: sub}, nil
}

// WatchPMWMultisigAccountAdded is a free log subscription operation binding the contract event 0xa43e4d705923e9d5e070e31f77dc6e19098c055be0bc847f73cc082f175511a5.
//
// Solidity: event PMWMultisigAccountAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 initialNonce, address authorizationAddress, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeePayments *TeePaymentsFilterer) WatchPMWMultisigAccountAdded(opts *bind.WatchOpts, sink chan<- *TeePaymentsPMWMultisigAccountAdded, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePayments.contract.WatchLogs(opts, "PMWMultisigAccountAdded", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsPMWMultisigAccountAdded)
				if err := _TeePayments.contract.UnpackLog(event, "PMWMultisigAccountAdded", log); err != nil {
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

// ParsePMWMultisigAccountAdded is a log parse operation binding the contract event 0xa43e4d705923e9d5e070e31f77dc6e19098c055be0bc847f73cc082f175511a5.
//
// Solidity: event PMWMultisigAccountAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 initialNonce, address authorizationAddress, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeePayments *TeePaymentsFilterer) ParsePMWMultisigAccountAdded(log types.Log) (*TeePaymentsPMWMultisigAccountAdded, error) {
	event := new(TeePaymentsPMWMultisigAccountAdded)
	if err := _TeePayments.contract.UnpackLog(event, "PMWMultisigAccountAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeePayments contract.
type TeePaymentsTimelockedGovernanceCallCanceledIterator struct {
	Event *TeePaymentsTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeePaymentsTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeePaymentsTimelockedGovernanceCallCanceled)
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
func (it *TeePaymentsTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePayments contract.
type TeePaymentsTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeePayments *TeePaymentsFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeePaymentsTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeePayments.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsTimelockedGovernanceCallCanceledIterator{contract: _TeePayments.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeePayments *TeePaymentsFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeePaymentsTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeePayments.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsTimelockedGovernanceCallCanceled)
				if err := _TeePayments.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeePayments *TeePaymentsFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeePaymentsTimelockedGovernanceCallCanceled, error) {
	event := new(TeePaymentsTimelockedGovernanceCallCanceled)
	if err := _TeePayments.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeePayments contract.
type TeePaymentsTimelockedGovernanceCallExecutedIterator struct {
	Event *TeePaymentsTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeePaymentsTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeePaymentsTimelockedGovernanceCallExecuted)
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
func (it *TeePaymentsTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePayments contract.
type TeePaymentsTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeePayments *TeePaymentsFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeePaymentsTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeePayments.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsTimelockedGovernanceCallExecutedIterator{contract: _TeePayments.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeePayments *TeePaymentsFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeePaymentsTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeePayments.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsTimelockedGovernanceCallExecuted)
				if err := _TeePayments.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeePayments *TeePaymentsFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeePaymentsTimelockedGovernanceCallExecuted, error) {
	event := new(TeePaymentsTimelockedGovernanceCallExecuted)
	if err := _TeePayments.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeePayments contract.
type TeePaymentsUpgradedIterator struct {
	Event *TeePaymentsUpgraded // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUpgraded)
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
		it.Event = new(TeePaymentsUpgraded)
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
func (it *TeePaymentsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUpgraded represents a Upgraded event raised by the TeePayments contract.
type TeePaymentsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePayments *TeePaymentsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeePaymentsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePayments.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUpgradedIterator{contract: _TeePayments.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePayments *TeePaymentsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeePaymentsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePayments.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUpgraded)
				if err := _TeePayments.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeePayments *TeePaymentsFilterer) ParseUpgraded(log types.Log) (*TeePaymentsUpgraded, error) {
	event := new(TeePaymentsUpgraded)
	if err := _TeePayments.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
