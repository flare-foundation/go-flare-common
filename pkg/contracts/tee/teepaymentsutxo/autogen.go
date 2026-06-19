// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepaymentsutxo

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

// IPMWMultisigUtxoConfiguredAnchor is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigUtxoConfiguredAnchor struct {
	GenesisAnchorTxid [32]byte
	GenesisAnchorVout uint32
}

// IPMWMultisigUtxoConfiguredProof is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigUtxoConfiguredProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  IPMWMultisigUtxoConfiguredRequestBody
	ResponseBody IPMWMultisigUtxoConfiguredResponseBody
}

// IPMWMultisigUtxoConfiguredRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigUtxoConfiguredRequestBody struct {
	AccountIndex uint32
	PublicKeys   [][]byte
	Threshold    uint64
	Anchors      []IPMWMultisigUtxoConfiguredAnchor
}

// IPMWMultisigUtxoConfiguredResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigUtxoConfiguredResponseBody struct {
	Status          uint8
	AccountAddress  string
	AnchorAddresses []string
}

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

// ITeePaymentsUtxoUtxoAnchorState is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsUtxoUtxoAnchorState struct {
	AnchorAddress     string
	GenesisAnchorTxid [32]byte
	GenesisAnchorVout uint32
	NextNonce         uint64
	AvailableAt       uint64
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// TeePaymentsUtxoMetaData contains all meta data concerning the TeePaymentsUtxo contract.
var TeePaymentsUtxoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountIndexMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorIndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"availableAt\",\"type\":\"uint64\"}],\"name\":\"AnchorNotReady\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuthorizationAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotYetEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchSizeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeFactor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxBatchSizeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoActiveReplacement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNewAnchors\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPaymentInstructions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuthorizationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMWMultisigAccountAddressAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMWMultisigAccountNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentNotInBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReissueRewardEpochChanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReplacementAlreadyFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ScheduledSignaturesUnsupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockInvalidSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongKeyType\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"anchorReuseDelaySeconds\",\"type\":\"uint64\"}],\"name\":\"DefaultAnchorReuseDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxBatchSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxBatchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"MaxBatchSettingsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"anchorCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizationAddress\",\"type\":\"address\"}],\"name\":\"PMWMultisigUtxoAccountAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"anchorCount\",\"type\":\"uint32\"}],\"name\":\"UtxoAnchorsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"UtxoBatchSettingsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchPaymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"replacementId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"firstPaymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"paymentCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"blocks\",\"type\":\"uint256[]\"}],\"name\":\"UtxoReplacementReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchPaymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"replacementId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"firstPaymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"name\":\"UtxoReplacementStarted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Anchor[]\",\"name\":\"anchors\",\"type\":\"tuple[]\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigUtxoConfigured.PMWMultisigUtxoStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"anchorAddresses\",\"type\":\"string[]\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"addAnchors\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Anchor[]\",\"name\":\"anchors\",\"type\":\"tuple[]\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigUtxoConfigured.PMWMultisigUtxoStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"anchorAddresses\",\"type\":\"string[]\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"name\":\"addPMWMultisigAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"contractIIFlareTeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_anchorIndex\",\"type\":\"uint256\"}],\"name\":\"getAnchor\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"anchorAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"nextNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"availableAt\",\"type\":\"uint64\"}],\"internalType\":\"structITeePaymentsUtxo.UtxoAnchorState\",\"name\":\"_anchor\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAnchorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAuthorizationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getBatchSettings\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getDefaultAnchorReuseDelay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getMaxBatchSettings\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_maxBatchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_maxBatchDurationSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"}],\"name\":\"getPaymentFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAccounts\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getWalletId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePaymentsBase.PaymentInstruction\",\"name\":\"_paymentInstruction\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentModel\",\"outputs\":[{\"internalType\":\"enumPaymentModel\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_batchPaymentId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePaymentsBase.PaymentInstruction[]\",\"name\":\"_paymentInstructions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"maxFeePerPayment\",\"type\":\"uint256[]\"},{\"internalType\":\"int16[][]\",\"name\":\"factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"delaysSeconds\",\"type\":\"uint16[]\"}],\"internalType\":\"structITeePaymentsBase.ReissueFeeParams\",\"name\":\"_reissueFeeParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"reissue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_finalized\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"setBatchSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_anchorReuseDelaySeconds\",\"type\":\"uint64\"}],\"name\":\"setDefaultAnchorReuseDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_maxBatchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_maxBatchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"setMaxBatchSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsConfigVerifier\",\"outputs\":[{\"internalType\":\"contractITeePaymentsConfigVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsFeeScheduleManager\",\"outputs\":[{\"internalType\":\"contractITeePaymentsFeeScheduleManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsRegistry\",\"outputs\":[{\"internalType\":\"contractITeePaymentsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeePaymentsUtxoABI is the input ABI used to generate the binding from.
// Deprecated: Use TeePaymentsUtxoMetaData.ABI instead.
var TeePaymentsUtxoABI = TeePaymentsUtxoMetaData.ABI

// TeePaymentsUtxo is an auto generated Go binding around an Ethereum contract.
type TeePaymentsUtxo struct {
	TeePaymentsUtxoCaller     // Read-only binding to the contract
	TeePaymentsUtxoTransactor // Write-only binding to the contract
	TeePaymentsUtxoFilterer   // Log filterer for contract events
}

// TeePaymentsUtxoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeePaymentsUtxoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsUtxoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeePaymentsUtxoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsUtxoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeePaymentsUtxoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsUtxoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeePaymentsUtxoSession struct {
	Contract     *TeePaymentsUtxo  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeePaymentsUtxoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeePaymentsUtxoCallerSession struct {
	Contract *TeePaymentsUtxoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TeePaymentsUtxoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeePaymentsUtxoTransactorSession struct {
	Contract     *TeePaymentsUtxoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TeePaymentsUtxoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeePaymentsUtxoRaw struct {
	Contract *TeePaymentsUtxo // Generic contract binding to access the raw methods on
}

// TeePaymentsUtxoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeePaymentsUtxoCallerRaw struct {
	Contract *TeePaymentsUtxoCaller // Generic read-only contract binding to access the raw methods on
}

// TeePaymentsUtxoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeePaymentsUtxoTransactorRaw struct {
	Contract *TeePaymentsUtxoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeePaymentsUtxo creates a new instance of TeePaymentsUtxo, bound to a specific deployed contract.
func NewTeePaymentsUtxo(address common.Address, backend bind.ContractBackend) (*TeePaymentsUtxo, error) {
	contract, err := bindTeePaymentsUtxo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxo{TeePaymentsUtxoCaller: TeePaymentsUtxoCaller{contract: contract}, TeePaymentsUtxoTransactor: TeePaymentsUtxoTransactor{contract: contract}, TeePaymentsUtxoFilterer: TeePaymentsUtxoFilterer{contract: contract}}, nil
}

// NewTeePaymentsUtxoCaller creates a new read-only instance of TeePaymentsUtxo, bound to a specific deployed contract.
func NewTeePaymentsUtxoCaller(address common.Address, caller bind.ContractCaller) (*TeePaymentsUtxoCaller, error) {
	contract, err := bindTeePaymentsUtxo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoCaller{contract: contract}, nil
}

// NewTeePaymentsUtxoTransactor creates a new write-only instance of TeePaymentsUtxo, bound to a specific deployed contract.
func NewTeePaymentsUtxoTransactor(address common.Address, transactor bind.ContractTransactor) (*TeePaymentsUtxoTransactor, error) {
	contract, err := bindTeePaymentsUtxo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoTransactor{contract: contract}, nil
}

// NewTeePaymentsUtxoFilterer creates a new log filterer instance of TeePaymentsUtxo, bound to a specific deployed contract.
func NewTeePaymentsUtxoFilterer(address common.Address, filterer bind.ContractFilterer) (*TeePaymentsUtxoFilterer, error) {
	contract, err := bindTeePaymentsUtxo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoFilterer{contract: contract}, nil
}

// bindTeePaymentsUtxo binds a generic wrapper to an already deployed contract.
func bindTeePaymentsUtxo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeePaymentsUtxoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePaymentsUtxo *TeePaymentsUtxoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePaymentsUtxo.Contract.TeePaymentsUtxoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePaymentsUtxo *TeePaymentsUtxoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.TeePaymentsUtxoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePaymentsUtxo *TeePaymentsUtxoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.TeePaymentsUtxoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePaymentsUtxo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePaymentsUtxo.Contract.UPGRADEINTERFACEVERSION(&_TeePaymentsUtxo.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePaymentsUtxo.Contract.UPGRADEINTERFACEVERSION(&_TeePaymentsUtxo.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) FlareSystemsManager() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.FlareSystemsManager(&_TeePaymentsUtxo.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) FlareSystemsManager() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.FlareSystemsManager(&_TeePaymentsUtxo.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) FlareTeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "flareTeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) FlareTeeManager() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.FlareTeeManager(&_TeePaymentsUtxo.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) FlareTeeManager() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.FlareTeeManager(&_TeePaymentsUtxo.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) GetAddressUpdater() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.GetAddressUpdater(&_TeePaymentsUtxo.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.GetAddressUpdater(&_TeePaymentsUtxo.CallOpts)
}

// GetAnchor is a free data retrieval call binding the contract method 0xaf278814.
//
// Solidity: function getAnchor((bytes32,string) _account, uint256 _anchorIndex) view returns((string,bytes32,uint32,uint64,uint64) _anchor)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) GetAnchor(opts *bind.CallOpts, _account ITeePaymentsBasePMWMultisigAccount, _anchorIndex *big.Int) (ITeePaymentsUtxoUtxoAnchorState, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "getAnchor", _account, _anchorIndex)

	if err != nil {
		return *new(ITeePaymentsUtxoUtxoAnchorState), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeePaymentsUtxoUtxoAnchorState)).(*ITeePaymentsUtxoUtxoAnchorState)

	return out0, err

}

// GetAnchor is a free data retrieval call binding the contract method 0xaf278814.
//
// Solidity: function getAnchor((bytes32,string) _account, uint256 _anchorIndex) view returns((string,bytes32,uint32,uint64,uint64) _anchor)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) GetAnchor(_account ITeePaymentsBasePMWMultisigAccount, _anchorIndex *big.Int) (ITeePaymentsUtxoUtxoAnchorState, error) {
	return _TeePaymentsUtxo.Contract.GetAnchor(&_TeePaymentsUtxo.CallOpts, _account, _anchorIndex)
}

// GetAnchor is a free data retrieval call binding the contract method 0xaf278814.
//
// Solidity: function getAnchor((bytes32,string) _account, uint256 _anchorIndex) view returns((string,bytes32,uint32,uint64,uint64) _anchor)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) GetAnchor(_account ITeePaymentsBasePMWMultisigAccount, _anchorIndex *big.Int) (ITeePaymentsUtxoUtxoAnchorState, error) {
	return _TeePaymentsUtxo.Contract.GetAnchor(&_TeePaymentsUtxo.CallOpts, _account, _anchorIndex)
}

// GetAnchorCount is a free data retrieval call binding the contract method 0x00dc90fd.
//
// Solidity: function getAnchorCount((bytes32,string) _account) view returns(uint256)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) GetAnchorCount(opts *bind.CallOpts, _account ITeePaymentsBasePMWMultisigAccount) (*big.Int, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "getAnchorCount", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAnchorCount is a free data retrieval call binding the contract method 0x00dc90fd.
//
// Solidity: function getAnchorCount((bytes32,string) _account) view returns(uint256)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) GetAnchorCount(_account ITeePaymentsBasePMWMultisigAccount) (*big.Int, error) {
	return _TeePaymentsUtxo.Contract.GetAnchorCount(&_TeePaymentsUtxo.CallOpts, _account)
}

// GetAnchorCount is a free data retrieval call binding the contract method 0x00dc90fd.
//
// Solidity: function getAnchorCount((bytes32,string) _account) view returns(uint256)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) GetAnchorCount(_account ITeePaymentsBasePMWMultisigAccount) (*big.Int, error) {
	return _TeePaymentsUtxo.Contract.GetAnchorCount(&_TeePaymentsUtxo.CallOpts, _account)
}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) GetAuthorizationAddress(opts *bind.CallOpts, _account ITeePaymentsBasePMWMultisigAccount) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "getAuthorizationAddress", _account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) GetAuthorizationAddress(_account ITeePaymentsBasePMWMultisigAccount) (common.Address, error) {
	return _TeePaymentsUtxo.Contract.GetAuthorizationAddress(&_TeePaymentsUtxo.CallOpts, _account)
}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) GetAuthorizationAddress(_account ITeePaymentsBasePMWMultisigAccount) (common.Address, error) {
	return _TeePaymentsUtxo.Contract.GetAuthorizationAddress(&_TeePaymentsUtxo.CallOpts, _account)
}

// GetBatchSettings is a free data retrieval call binding the contract method 0xb96c3c8f.
//
// Solidity: function getBatchSettings((bytes32,string) _account) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) GetBatchSettings(opts *bind.CallOpts, _account ITeePaymentsBasePMWMultisigAccount) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "getBatchSettings", _account)

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
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) GetBatchSettings(_account ITeePaymentsBasePMWMultisigAccount) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	return _TeePaymentsUtxo.Contract.GetBatchSettings(&_TeePaymentsUtxo.CallOpts, _account)
}

// GetBatchSettings is a free data retrieval call binding the contract method 0xb96c3c8f.
//
// Solidity: function getBatchSettings((bytes32,string) _account) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) GetBatchSettings(_account ITeePaymentsBasePMWMultisigAccount) (struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}, error) {
	return _TeePaymentsUtxo.Contract.GetBatchSettings(&_TeePaymentsUtxo.CallOpts, _account)
}

// GetDefaultAnchorReuseDelay is a free data retrieval call binding the contract method 0x69738527.
//
// Solidity: function getDefaultAnchorReuseDelay(bytes32 _sourceId) view returns(uint64)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) GetDefaultAnchorReuseDelay(opts *bind.CallOpts, _sourceId [32]byte) (uint64, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "getDefaultAnchorReuseDelay", _sourceId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetDefaultAnchorReuseDelay is a free data retrieval call binding the contract method 0x69738527.
//
// Solidity: function getDefaultAnchorReuseDelay(bytes32 _sourceId) view returns(uint64)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) GetDefaultAnchorReuseDelay(_sourceId [32]byte) (uint64, error) {
	return _TeePaymentsUtxo.Contract.GetDefaultAnchorReuseDelay(&_TeePaymentsUtxo.CallOpts, _sourceId)
}

// GetDefaultAnchorReuseDelay is a free data retrieval call binding the contract method 0x69738527.
//
// Solidity: function getDefaultAnchorReuseDelay(bytes32 _sourceId) view returns(uint64)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) GetDefaultAnchorReuseDelay(_sourceId [32]byte) (uint64, error) {
	return _TeePaymentsUtxo.Contract.GetDefaultAnchorReuseDelay(&_TeePaymentsUtxo.CallOpts, _sourceId)
}

// GetMaxBatchSettings is a free data retrieval call binding the contract method 0xf6cc38e1.
//
// Solidity: function getMaxBatchSettings(bytes32 _sourceId) view returns(uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) GetMaxBatchSettings(opts *bind.CallOpts, _sourceId [32]byte) (struct {
	MaxBatchSize            uint64
	MaxBatchDurationSeconds uint64
}, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "getMaxBatchSettings", _sourceId)

	outstruct := new(struct {
		MaxBatchSize            uint64
		MaxBatchDurationSeconds uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxBatchSize = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.MaxBatchDurationSeconds = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetMaxBatchSettings is a free data retrieval call binding the contract method 0xf6cc38e1.
//
// Solidity: function getMaxBatchSettings(bytes32 _sourceId) view returns(uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) GetMaxBatchSettings(_sourceId [32]byte) (struct {
	MaxBatchSize            uint64
	MaxBatchDurationSeconds uint64
}, error) {
	return _TeePaymentsUtxo.Contract.GetMaxBatchSettings(&_TeePaymentsUtxo.CallOpts, _sourceId)
}

// GetMaxBatchSettings is a free data retrieval call binding the contract method 0xf6cc38e1.
//
// Solidity: function getMaxBatchSettings(bytes32 _sourceId) view returns(uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) GetMaxBatchSettings(_sourceId [32]byte) (struct {
	MaxBatchSize            uint64
	MaxBatchDurationSeconds uint64
}, error) {
	return _TeePaymentsUtxo.Contract.GetMaxBatchSettings(&_TeePaymentsUtxo.CallOpts, _sourceId)
}

// GetPaymentFee is a free data retrieval call binding the contract method 0x57abf78b.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) GetPaymentFee(opts *bind.CallOpts, _account ITeePaymentsBasePMWMultisigAccount, _opCommand [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "getPaymentFee", _account, _opCommand)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPaymentFee is a free data retrieval call binding the contract method 0x57abf78b.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) GetPaymentFee(_account ITeePaymentsBasePMWMultisigAccount, _opCommand [32]byte) (*big.Int, error) {
	return _TeePaymentsUtxo.Contract.GetPaymentFee(&_TeePaymentsUtxo.CallOpts, _account, _opCommand)
}

// GetPaymentFee is a free data retrieval call binding the contract method 0x57abf78b.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) GetPaymentFee(_account ITeePaymentsBasePMWMultisigAccount, _opCommand [32]byte) (*big.Int, error) {
	return _TeePaymentsUtxo.Contract.GetPaymentFee(&_TeePaymentsUtxo.CallOpts, _account, _opCommand)
}

// GetWalletAccounts is a free data retrieval call binding the contract method 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) GetWalletAccounts(opts *bind.CallOpts, _walletId [32]byte) ([]ITeePaymentsBasePMWMultisigAccount, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "getWalletAccounts", _walletId)

	if err != nil {
		return *new([]ITeePaymentsBasePMWMultisigAccount), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITeePaymentsBasePMWMultisigAccount)).(*[]ITeePaymentsBasePMWMultisigAccount)

	return out0, err

}

// GetWalletAccounts is a free data retrieval call binding the contract method 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) GetWalletAccounts(_walletId [32]byte) ([]ITeePaymentsBasePMWMultisigAccount, error) {
	return _TeePaymentsUtxo.Contract.GetWalletAccounts(&_TeePaymentsUtxo.CallOpts, _walletId)
}

// GetWalletAccounts is a free data retrieval call binding the contract method 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) GetWalletAccounts(_walletId [32]byte) ([]ITeePaymentsBasePMWMultisigAccount, error) {
	return _TeePaymentsUtxo.Contract.GetWalletAccounts(&_TeePaymentsUtxo.CallOpts, _walletId)
}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) GetWalletId(opts *bind.CallOpts, _account ITeePaymentsBasePMWMultisigAccount) ([32]byte, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "getWalletId", _account)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) GetWalletId(_account ITeePaymentsBasePMWMultisigAccount) ([32]byte, error) {
	return _TeePaymentsUtxo.Contract.GetWalletId(&_TeePaymentsUtxo.CallOpts, _account)
}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) GetWalletId(_account ITeePaymentsBasePMWMultisigAccount) ([32]byte, error) {
	return _TeePaymentsUtxo.Contract.GetWalletId(&_TeePaymentsUtxo.CallOpts, _account)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) Governance() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.Governance(&_TeePaymentsUtxo.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) Governance() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.Governance(&_TeePaymentsUtxo.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) GovernanceSettings() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.GovernanceSettings(&_TeePaymentsUtxo.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.GovernanceSettings(&_TeePaymentsUtxo.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) Implementation() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.Implementation(&_TeePaymentsUtxo.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) Implementation() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.Implementation(&_TeePaymentsUtxo.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePaymentsUtxo.Contract.IsExecutor(&_TeePaymentsUtxo.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePaymentsUtxo.Contract.IsExecutor(&_TeePaymentsUtxo.CallOpts, _address)
}

// PaymentModel is a free data retrieval call binding the contract method 0xbb9d3dbb.
//
// Solidity: function paymentModel() pure returns(uint8)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) PaymentModel(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "paymentModel")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PaymentModel is a free data retrieval call binding the contract method 0xbb9d3dbb.
//
// Solidity: function paymentModel() pure returns(uint8)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) PaymentModel() (uint8, error) {
	return _TeePaymentsUtxo.Contract.PaymentModel(&_TeePaymentsUtxo.CallOpts)
}

// PaymentModel is a free data retrieval call binding the contract method 0xbb9d3dbb.
//
// Solidity: function paymentModel() pure returns(uint8)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) PaymentModel() (uint8, error) {
	return _TeePaymentsUtxo.Contract.PaymentModel(&_TeePaymentsUtxo.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) ProductionMode() (bool, error) {
	return _TeePaymentsUtxo.Contract.ProductionMode(&_TeePaymentsUtxo.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) ProductionMode() (bool, error) {
	return _TeePaymentsUtxo.Contract.ProductionMode(&_TeePaymentsUtxo.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) ProxiableUUID() ([32]byte, error) {
	return _TeePaymentsUtxo.Contract.ProxiableUUID(&_TeePaymentsUtxo.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeePaymentsUtxo.Contract.ProxiableUUID(&_TeePaymentsUtxo.CallOpts)
}

// TeePaymentsConfigVerifier is a free data retrieval call binding the contract method 0xf71c6c75.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) TeePaymentsConfigVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "teePaymentsConfigVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeePaymentsConfigVerifier is a free data retrieval call binding the contract method 0xf71c6c75.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) TeePaymentsConfigVerifier() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.TeePaymentsConfigVerifier(&_TeePaymentsUtxo.CallOpts)
}

// TeePaymentsConfigVerifier is a free data retrieval call binding the contract method 0xf71c6c75.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) TeePaymentsConfigVerifier() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.TeePaymentsConfigVerifier(&_TeePaymentsUtxo.CallOpts)
}

// TeePaymentsFeeScheduleManager is a free data retrieval call binding the contract method 0xfe9e33e8.
//
// Solidity: function teePaymentsFeeScheduleManager() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) TeePaymentsFeeScheduleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "teePaymentsFeeScheduleManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeePaymentsFeeScheduleManager is a free data retrieval call binding the contract method 0xfe9e33e8.
//
// Solidity: function teePaymentsFeeScheduleManager() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) TeePaymentsFeeScheduleManager() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.TeePaymentsFeeScheduleManager(&_TeePaymentsUtxo.CallOpts)
}

// TeePaymentsFeeScheduleManager is a free data retrieval call binding the contract method 0xfe9e33e8.
//
// Solidity: function teePaymentsFeeScheduleManager() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) TeePaymentsFeeScheduleManager() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.TeePaymentsFeeScheduleManager(&_TeePaymentsUtxo.CallOpts)
}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCaller) TeePaymentsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsUtxo.contract.Call(opts, &out, "teePaymentsRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) TeePaymentsRegistry() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.TeePaymentsRegistry(&_TeePaymentsUtxo.CallOpts)
}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsUtxo *TeePaymentsUtxoCallerSession) TeePaymentsRegistry() (common.Address, error) {
	return _TeePaymentsUtxo.Contract.TeePaymentsRegistry(&_TeePaymentsUtxo.CallOpts)
}

// AddAnchors is a paid mutator transaction binding the contract method 0xb3fa02b3.
//
// Solidity: function addAnchors(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string,string[])) _proof) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactor) AddAnchors(opts *bind.TransactOpts, _proof IPMWMultisigUtxoConfiguredProof) (*types.Transaction, error) {
	return _TeePaymentsUtxo.contract.Transact(opts, "addAnchors", _proof)
}

// AddAnchors is a paid mutator transaction binding the contract method 0xb3fa02b3.
//
// Solidity: function addAnchors(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string,string[])) _proof) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) AddAnchors(_proof IPMWMultisigUtxoConfiguredProof) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.AddAnchors(&_TeePaymentsUtxo.TransactOpts, _proof)
}

// AddAnchors is a paid mutator transaction binding the contract method 0xb3fa02b3.
//
// Solidity: function addAnchors(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string,string[])) _proof) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorSession) AddAnchors(_proof IPMWMultisigUtxoConfiguredProof) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.AddAnchors(&_TeePaymentsUtxo.TransactOpts, _proof)
}

// AddPMWMultisigAccount is a paid mutator transaction binding the contract method 0x2a0524b6.
//
// Solidity: function addPMWMultisigAccount(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string,string[])) _proof, address _authorizationAddress) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactor) AddPMWMultisigAccount(opts *bind.TransactOpts, _walletId [32]byte, _proof IPMWMultisigUtxoConfiguredProof, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.contract.Transact(opts, "addPMWMultisigAccount", _walletId, _proof, _authorizationAddress)
}

// AddPMWMultisigAccount is a paid mutator transaction binding the contract method 0x2a0524b6.
//
// Solidity: function addPMWMultisigAccount(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string,string[])) _proof, address _authorizationAddress) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) AddPMWMultisigAccount(_walletId [32]byte, _proof IPMWMultisigUtxoConfiguredProof, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.AddPMWMultisigAccount(&_TeePaymentsUtxo.TransactOpts, _walletId, _proof, _authorizationAddress)
}

// AddPMWMultisigAccount is a paid mutator transaction binding the contract method 0x2a0524b6.
//
// Solidity: function addPMWMultisigAccount(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string,string[])) _proof, address _authorizationAddress) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorSession) AddPMWMultisigAccount(_walletId [32]byte, _proof IPMWMultisigUtxoConfiguredProof, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.AddPMWMultisigAccount(&_TeePaymentsUtxo.TransactOpts, _walletId, _proof, _authorizationAddress)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsUtxo.contract.Transact(opts, "cancelGovernanceCall", _encodedCall)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) CancelGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.CancelGovernanceCall(&_TeePaymentsUtxo.TransactOpts, _encodedCall)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorSession) CancelGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.CancelGovernanceCall(&_TeePaymentsUtxo.TransactOpts, _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsUtxo.contract.Transact(opts, "executeGovernanceCall", _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) ExecuteGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.ExecuteGovernanceCall(&_TeePaymentsUtxo.TransactOpts, _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorSession) ExecuteGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.ExecuteGovernanceCall(&_TeePaymentsUtxo.TransactOpts, _encodedCall)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.Initialize(&_TeePaymentsUtxo.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.Initialize(&_TeePaymentsUtxo.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Pay is a paid mutator transaction binding the contract method 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactor) Pay(opts *bind.TransactOpts, _account ITeePaymentsBasePMWMultisigAccount, _paymentInstruction ITeePaymentsBasePaymentInstruction, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.contract.Transact(opts, "pay", _account, _paymentInstruction, _claimBackAddress)
}

// Pay is a paid mutator transaction binding the contract method 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) Pay(_account ITeePaymentsBasePMWMultisigAccount, _paymentInstruction ITeePaymentsBasePaymentInstruction, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.Pay(&_TeePaymentsUtxo.TransactOpts, _account, _paymentInstruction, _claimBackAddress)
}

// Pay is a paid mutator transaction binding the contract method 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorSession) Pay(_account ITeePaymentsBasePMWMultisigAccount, _paymentInstruction ITeePaymentsBasePaymentInstruction, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.Pay(&_TeePaymentsUtxo.TransactOpts, _account, _paymentInstruction, _claimBackAddress)
}

// Reissue is a paid mutator transaction binding the contract method 0x1a82a605.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _batchPaymentId, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns(bool _finalized)
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactor) Reissue(opts *bind.TransactOpts, _account ITeePaymentsBasePMWMultisigAccount, _batchPaymentId uint64, _paymentInstructions []ITeePaymentsBasePaymentInstruction, _reissueFeeParams ITeePaymentsBaseReissueFeeParams, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.contract.Transact(opts, "reissue", _account, _batchPaymentId, _paymentInstructions, _reissueFeeParams, _claimBackAddress)
}

// Reissue is a paid mutator transaction binding the contract method 0x1a82a605.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _batchPaymentId, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns(bool _finalized)
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) Reissue(_account ITeePaymentsBasePMWMultisigAccount, _batchPaymentId uint64, _paymentInstructions []ITeePaymentsBasePaymentInstruction, _reissueFeeParams ITeePaymentsBaseReissueFeeParams, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.Reissue(&_TeePaymentsUtxo.TransactOpts, _account, _batchPaymentId, _paymentInstructions, _reissueFeeParams, _claimBackAddress)
}

// Reissue is a paid mutator transaction binding the contract method 0x1a82a605.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _batchPaymentId, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns(bool _finalized)
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorSession) Reissue(_account ITeePaymentsBasePMWMultisigAccount, _batchPaymentId uint64, _paymentInstructions []ITeePaymentsBasePaymentInstruction, _reissueFeeParams ITeePaymentsBaseReissueFeeParams, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.Reissue(&_TeePaymentsUtxo.TransactOpts, _account, _batchPaymentId, _paymentInstructions, _reissueFeeParams, _claimBackAddress)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xd302be77.
//
// Solidity: function setBatchSettings((bytes32,string) _account, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactor) SetBatchSettings(opts *bind.TransactOpts, _account ITeePaymentsBasePMWMultisigAccount, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePaymentsUtxo.contract.Transact(opts, "setBatchSettings", _account, _batchSize, _batchDurationSeconds)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xd302be77.
//
// Solidity: function setBatchSettings((bytes32,string) _account, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) SetBatchSettings(_account ITeePaymentsBasePMWMultisigAccount, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.SetBatchSettings(&_TeePaymentsUtxo.TransactOpts, _account, _batchSize, _batchDurationSeconds)
}

// SetBatchSettings is a paid mutator transaction binding the contract method 0xd302be77.
//
// Solidity: function setBatchSettings((bytes32,string) _account, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorSession) SetBatchSettings(_account ITeePaymentsBasePMWMultisigAccount, _batchSize uint64, _batchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.SetBatchSettings(&_TeePaymentsUtxo.TransactOpts, _account, _batchSize, _batchDurationSeconds)
}

// SetDefaultAnchorReuseDelay is a paid mutator transaction binding the contract method 0xabe2259e.
//
// Solidity: function setDefaultAnchorReuseDelay(bytes32 _sourceId, uint64 _anchorReuseDelaySeconds) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactor) SetDefaultAnchorReuseDelay(opts *bind.TransactOpts, _sourceId [32]byte, _anchorReuseDelaySeconds uint64) (*types.Transaction, error) {
	return _TeePaymentsUtxo.contract.Transact(opts, "setDefaultAnchorReuseDelay", _sourceId, _anchorReuseDelaySeconds)
}

// SetDefaultAnchorReuseDelay is a paid mutator transaction binding the contract method 0xabe2259e.
//
// Solidity: function setDefaultAnchorReuseDelay(bytes32 _sourceId, uint64 _anchorReuseDelaySeconds) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) SetDefaultAnchorReuseDelay(_sourceId [32]byte, _anchorReuseDelaySeconds uint64) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.SetDefaultAnchorReuseDelay(&_TeePaymentsUtxo.TransactOpts, _sourceId, _anchorReuseDelaySeconds)
}

// SetDefaultAnchorReuseDelay is a paid mutator transaction binding the contract method 0xabe2259e.
//
// Solidity: function setDefaultAnchorReuseDelay(bytes32 _sourceId, uint64 _anchorReuseDelaySeconds) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorSession) SetDefaultAnchorReuseDelay(_sourceId [32]byte, _anchorReuseDelaySeconds uint64) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.SetDefaultAnchorReuseDelay(&_TeePaymentsUtxo.TransactOpts, _sourceId, _anchorReuseDelaySeconds)
}

// SetMaxBatchSettings is a paid mutator transaction binding the contract method 0xcb42368a.
//
// Solidity: function setMaxBatchSettings(bytes32 _sourceId, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactor) SetMaxBatchSettings(opts *bind.TransactOpts, _sourceId [32]byte, _maxBatchSize uint64, _maxBatchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePaymentsUtxo.contract.Transact(opts, "setMaxBatchSettings", _sourceId, _maxBatchSize, _maxBatchDurationSeconds)
}

// SetMaxBatchSettings is a paid mutator transaction binding the contract method 0xcb42368a.
//
// Solidity: function setMaxBatchSettings(bytes32 _sourceId, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) SetMaxBatchSettings(_sourceId [32]byte, _maxBatchSize uint64, _maxBatchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.SetMaxBatchSettings(&_TeePaymentsUtxo.TransactOpts, _sourceId, _maxBatchSize, _maxBatchDurationSeconds)
}

// SetMaxBatchSettings is a paid mutator transaction binding the contract method 0xcb42368a.
//
// Solidity: function setMaxBatchSettings(bytes32 _sourceId, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorSession) SetMaxBatchSettings(_sourceId [32]byte, _maxBatchSize uint64, _maxBatchDurationSeconds uint64) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.SetMaxBatchSettings(&_TeePaymentsUtxo.TransactOpts, _sourceId, _maxBatchSize, _maxBatchDurationSeconds)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsUtxo.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.SwitchToProductionMode(&_TeePaymentsUtxo.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.SwitchToProductionMode(&_TeePaymentsUtxo.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.UpdateContractAddresses(&_TeePaymentsUtxo.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.UpdateContractAddresses(&_TeePaymentsUtxo.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsUtxo.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.UpgradeToAndCall(&_TeePaymentsUtxo.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsUtxo *TeePaymentsUtxoTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsUtxo.Contract.UpgradeToAndCall(&_TeePaymentsUtxo.TransactOpts, _newImplementation, _data)
}

// TeePaymentsUtxoDefaultAnchorReuseDelaySetIterator is returned from FilterDefaultAnchorReuseDelaySet and is used to iterate over the raw logs and unpacked data for DefaultAnchorReuseDelaySet events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoDefaultAnchorReuseDelaySetIterator struct {
	Event *TeePaymentsUtxoDefaultAnchorReuseDelaySet // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoDefaultAnchorReuseDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoDefaultAnchorReuseDelaySet)
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
		it.Event = new(TeePaymentsUtxoDefaultAnchorReuseDelaySet)
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
func (it *TeePaymentsUtxoDefaultAnchorReuseDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoDefaultAnchorReuseDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoDefaultAnchorReuseDelaySet represents a DefaultAnchorReuseDelaySet event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoDefaultAnchorReuseDelaySet struct {
	SourceId                [32]byte
	AnchorReuseDelaySeconds uint64
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterDefaultAnchorReuseDelaySet is a free log retrieval operation binding the contract event 0x92a7abeb13b7d1b4fd7ae54df522f9e39d8ba9864bbe8b28f1daf260797ba9b9.
//
// Solidity: event DefaultAnchorReuseDelaySet(bytes32 indexed sourceId, uint64 anchorReuseDelaySeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterDefaultAnchorReuseDelaySet(opts *bind.FilterOpts, sourceId [][32]byte) (*TeePaymentsUtxoDefaultAnchorReuseDelaySetIterator, error) {

	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "DefaultAnchorReuseDelaySet", sourceIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoDefaultAnchorReuseDelaySetIterator{contract: _TeePaymentsUtxo.contract, event: "DefaultAnchorReuseDelaySet", logs: logs, sub: sub}, nil
}

// WatchDefaultAnchorReuseDelaySet is a free log subscription operation binding the contract event 0x92a7abeb13b7d1b4fd7ae54df522f9e39d8ba9864bbe8b28f1daf260797ba9b9.
//
// Solidity: event DefaultAnchorReuseDelaySet(bytes32 indexed sourceId, uint64 anchorReuseDelaySeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchDefaultAnchorReuseDelaySet(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoDefaultAnchorReuseDelaySet, sourceId [][32]byte) (event.Subscription, error) {

	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "DefaultAnchorReuseDelaySet", sourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoDefaultAnchorReuseDelaySet)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "DefaultAnchorReuseDelaySet", log); err != nil {
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

// ParseDefaultAnchorReuseDelaySet is a log parse operation binding the contract event 0x92a7abeb13b7d1b4fd7ae54df522f9e39d8ba9864bbe8b28f1daf260797ba9b9.
//
// Solidity: event DefaultAnchorReuseDelaySet(bytes32 indexed sourceId, uint64 anchorReuseDelaySeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParseDefaultAnchorReuseDelaySet(log types.Log) (*TeePaymentsUtxoDefaultAnchorReuseDelaySet, error) {
	event := new(TeePaymentsUtxoDefaultAnchorReuseDelaySet)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "DefaultAnchorReuseDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUtxoGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoGovernanceCallTimelockedIterator struct {
	Event *TeePaymentsUtxoGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoGovernanceCallTimelocked)
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
		it.Event = new(TeePaymentsUtxoGovernanceCallTimelocked)
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
func (it *TeePaymentsUtxoGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeePaymentsUtxoGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoGovernanceCallTimelockedIterator{contract: _TeePaymentsUtxo.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoGovernanceCallTimelocked)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeePaymentsUtxoGovernanceCallTimelocked, error) {
	event := new(TeePaymentsUtxoGovernanceCallTimelocked)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUtxoGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoGovernanceInitialisedIterator struct {
	Event *TeePaymentsUtxoGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoGovernanceInitialised)
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
		it.Event = new(TeePaymentsUtxoGovernanceInitialised)
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
func (it *TeePaymentsUtxoGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoGovernanceInitialised represents a GovernanceInitialised event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeePaymentsUtxoGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoGovernanceInitialisedIterator{contract: _TeePaymentsUtxo.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoGovernanceInitialised)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParseGovernanceInitialised(log types.Log) (*TeePaymentsUtxoGovernanceInitialised, error) {
	event := new(TeePaymentsUtxoGovernanceInitialised)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUtxoGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoGovernedProductionModeEnteredIterator struct {
	Event *TeePaymentsUtxoGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoGovernedProductionModeEntered)
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
		it.Event = new(TeePaymentsUtxoGovernedProductionModeEntered)
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
func (it *TeePaymentsUtxoGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeePaymentsUtxoGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoGovernedProductionModeEnteredIterator{contract: _TeePaymentsUtxo.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoGovernedProductionModeEntered)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeePaymentsUtxoGovernedProductionModeEntered, error) {
	event := new(TeePaymentsUtxoGovernedProductionModeEntered)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUtxoInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoInitializedIterator struct {
	Event *TeePaymentsUtxoInitialized // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoInitialized)
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
		it.Event = new(TeePaymentsUtxoInitialized)
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
func (it *TeePaymentsUtxoInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoInitialized represents a Initialized event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterInitialized(opts *bind.FilterOpts) (*TeePaymentsUtxoInitializedIterator, error) {

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoInitializedIterator{contract: _TeePaymentsUtxo.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoInitialized) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoInitialized)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParseInitialized(log types.Log) (*TeePaymentsUtxoInitialized, error) {
	event := new(TeePaymentsUtxoInitialized)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUtxoMaxBatchSettingsSetIterator is returned from FilterMaxBatchSettingsSet and is used to iterate over the raw logs and unpacked data for MaxBatchSettingsSet events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoMaxBatchSettingsSetIterator struct {
	Event *TeePaymentsUtxoMaxBatchSettingsSet // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoMaxBatchSettingsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoMaxBatchSettingsSet)
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
		it.Event = new(TeePaymentsUtxoMaxBatchSettingsSet)
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
func (it *TeePaymentsUtxoMaxBatchSettingsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoMaxBatchSettingsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoMaxBatchSettingsSet represents a MaxBatchSettingsSet event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoMaxBatchSettingsSet struct {
	SourceId                [32]byte
	MaxBatchSize            uint64
	MaxBatchDurationSeconds uint64
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterMaxBatchSettingsSet is a free log retrieval operation binding the contract event 0x42c5686e1b1d2704efd24a239b1204e823efad44594cda54f15b21c1d958494f.
//
// Solidity: event MaxBatchSettingsSet(bytes32 indexed sourceId, uint64 maxBatchSize, uint64 maxBatchDurationSeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterMaxBatchSettingsSet(opts *bind.FilterOpts, sourceId [][32]byte) (*TeePaymentsUtxoMaxBatchSettingsSetIterator, error) {

	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "MaxBatchSettingsSet", sourceIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoMaxBatchSettingsSetIterator{contract: _TeePaymentsUtxo.contract, event: "MaxBatchSettingsSet", logs: logs, sub: sub}, nil
}

// WatchMaxBatchSettingsSet is a free log subscription operation binding the contract event 0x42c5686e1b1d2704efd24a239b1204e823efad44594cda54f15b21c1d958494f.
//
// Solidity: event MaxBatchSettingsSet(bytes32 indexed sourceId, uint64 maxBatchSize, uint64 maxBatchDurationSeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchMaxBatchSettingsSet(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoMaxBatchSettingsSet, sourceId [][32]byte) (event.Subscription, error) {

	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "MaxBatchSettingsSet", sourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoMaxBatchSettingsSet)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "MaxBatchSettingsSet", log); err != nil {
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

// ParseMaxBatchSettingsSet is a log parse operation binding the contract event 0x42c5686e1b1d2704efd24a239b1204e823efad44594cda54f15b21c1d958494f.
//
// Solidity: event MaxBatchSettingsSet(bytes32 indexed sourceId, uint64 maxBatchSize, uint64 maxBatchDurationSeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParseMaxBatchSettingsSet(log types.Log) (*TeePaymentsUtxoMaxBatchSettingsSet, error) {
	event := new(TeePaymentsUtxoMaxBatchSettingsSet)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "MaxBatchSettingsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUtxoPMWMultisigUtxoAccountAddedIterator is returned from FilterPMWMultisigUtxoAccountAdded and is used to iterate over the raw logs and unpacked data for PMWMultisigUtxoAccountAdded events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoPMWMultisigUtxoAccountAddedIterator struct {
	Event *TeePaymentsUtxoPMWMultisigUtxoAccountAdded // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoPMWMultisigUtxoAccountAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoPMWMultisigUtxoAccountAdded)
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
		it.Event = new(TeePaymentsUtxoPMWMultisigUtxoAccountAdded)
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
func (it *TeePaymentsUtxoPMWMultisigUtxoAccountAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoPMWMultisigUtxoAccountAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoPMWMultisigUtxoAccountAdded represents a PMWMultisigUtxoAccountAdded event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoPMWMultisigUtxoAccountAdded struct {
	WalletId             [32]byte
	SourceId             [32]byte
	AccountAddress       string
	AccountIndex         uint32
	AnchorCount          uint32
	AuthorizationAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPMWMultisigUtxoAccountAdded is a free log retrieval operation binding the contract event 0x71383f39dd17f41e1d080d4f01d248b2d7c436c68aa9f7ed9ec375c947badcbc.
//
// Solidity: event PMWMultisigUtxoAccountAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount, address authorizationAddress)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterPMWMultisigUtxoAccountAdded(opts *bind.FilterOpts, walletId [][32]byte) (*TeePaymentsUtxoPMWMultisigUtxoAccountAddedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "PMWMultisigUtxoAccountAdded", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoPMWMultisigUtxoAccountAddedIterator{contract: _TeePaymentsUtxo.contract, event: "PMWMultisigUtxoAccountAdded", logs: logs, sub: sub}, nil
}

// WatchPMWMultisigUtxoAccountAdded is a free log subscription operation binding the contract event 0x71383f39dd17f41e1d080d4f01d248b2d7c436c68aa9f7ed9ec375c947badcbc.
//
// Solidity: event PMWMultisigUtxoAccountAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount, address authorizationAddress)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchPMWMultisigUtxoAccountAdded(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoPMWMultisigUtxoAccountAdded, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "PMWMultisigUtxoAccountAdded", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoPMWMultisigUtxoAccountAdded)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "PMWMultisigUtxoAccountAdded", log); err != nil {
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

// ParsePMWMultisigUtxoAccountAdded is a log parse operation binding the contract event 0x71383f39dd17f41e1d080d4f01d248b2d7c436c68aa9f7ed9ec375c947badcbc.
//
// Solidity: event PMWMultisigUtxoAccountAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount, address authorizationAddress)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParsePMWMultisigUtxoAccountAdded(log types.Log) (*TeePaymentsUtxoPMWMultisigUtxoAccountAdded, error) {
	event := new(TeePaymentsUtxoPMWMultisigUtxoAccountAdded)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "PMWMultisigUtxoAccountAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUtxoTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoTimelockedGovernanceCallCanceledIterator struct {
	Event *TeePaymentsUtxoTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeePaymentsUtxoTimelockedGovernanceCallCanceled)
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
func (it *TeePaymentsUtxoTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeePaymentsUtxoTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoTimelockedGovernanceCallCanceledIterator{contract: _TeePaymentsUtxo.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoTimelockedGovernanceCallCanceled)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeePaymentsUtxoTimelockedGovernanceCallCanceled, error) {
	event := new(TeePaymentsUtxoTimelockedGovernanceCallCanceled)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUtxoTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoTimelockedGovernanceCallExecutedIterator struct {
	Event *TeePaymentsUtxoTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeePaymentsUtxoTimelockedGovernanceCallExecuted)
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
func (it *TeePaymentsUtxoTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeePaymentsUtxoTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoTimelockedGovernanceCallExecutedIterator{contract: _TeePaymentsUtxo.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoTimelockedGovernanceCallExecuted)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeePaymentsUtxoTimelockedGovernanceCallExecuted, error) {
	event := new(TeePaymentsUtxoTimelockedGovernanceCallExecuted)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUtxoUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUpgradedIterator struct {
	Event *TeePaymentsUtxoUpgraded // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoUpgraded)
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
		it.Event = new(TeePaymentsUtxoUpgraded)
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
func (it *TeePaymentsUtxoUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoUpgraded represents a Upgraded event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeePaymentsUtxoUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoUpgradedIterator{contract: _TeePaymentsUtxo.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoUpgraded)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParseUpgraded(log types.Log) (*TeePaymentsUtxoUpgraded, error) {
	event := new(TeePaymentsUtxoUpgraded)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUtxoUtxoAnchorsAddedIterator is returned from FilterUtxoAnchorsAdded and is used to iterate over the raw logs and unpacked data for UtxoAnchorsAdded events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUtxoAnchorsAddedIterator struct {
	Event *TeePaymentsUtxoUtxoAnchorsAdded // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoUtxoAnchorsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoUtxoAnchorsAdded)
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
		it.Event = new(TeePaymentsUtxoUtxoAnchorsAdded)
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
func (it *TeePaymentsUtxoUtxoAnchorsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoUtxoAnchorsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoUtxoAnchorsAdded represents a UtxoAnchorsAdded event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUtxoAnchorsAdded struct {
	WalletId       [32]byte
	SourceId       [32]byte
	AccountAddress string
	AccountIndex   uint32
	AnchorCount    uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUtxoAnchorsAdded is a free log retrieval operation binding the contract event 0x181d960fe1b444001068e1f7b220327715170cba45596b12cc66686bf55d1147.
//
// Solidity: event UtxoAnchorsAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterUtxoAnchorsAdded(opts *bind.FilterOpts, walletId [][32]byte) (*TeePaymentsUtxoUtxoAnchorsAddedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "UtxoAnchorsAdded", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoUtxoAnchorsAddedIterator{contract: _TeePaymentsUtxo.contract, event: "UtxoAnchorsAdded", logs: logs, sub: sub}, nil
}

// WatchUtxoAnchorsAdded is a free log subscription operation binding the contract event 0x181d960fe1b444001068e1f7b220327715170cba45596b12cc66686bf55d1147.
//
// Solidity: event UtxoAnchorsAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchUtxoAnchorsAdded(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoUtxoAnchorsAdded, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "UtxoAnchorsAdded", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoUtxoAnchorsAdded)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "UtxoAnchorsAdded", log); err != nil {
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

// ParseUtxoAnchorsAdded is a log parse operation binding the contract event 0x181d960fe1b444001068e1f7b220327715170cba45596b12cc66686bf55d1147.
//
// Solidity: event UtxoAnchorsAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParseUtxoAnchorsAdded(log types.Log) (*TeePaymentsUtxoUtxoAnchorsAdded, error) {
	event := new(TeePaymentsUtxoUtxoAnchorsAdded)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "UtxoAnchorsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUtxoUtxoBatchSettingsSetIterator is returned from FilterUtxoBatchSettingsSet and is used to iterate over the raw logs and unpacked data for UtxoBatchSettingsSet events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUtxoBatchSettingsSetIterator struct {
	Event *TeePaymentsUtxoUtxoBatchSettingsSet // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoUtxoBatchSettingsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoUtxoBatchSettingsSet)
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
		it.Event = new(TeePaymentsUtxoUtxoBatchSettingsSet)
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
func (it *TeePaymentsUtxoUtxoBatchSettingsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoUtxoBatchSettingsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoUtxoBatchSettingsSet represents a UtxoBatchSettingsSet event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUtxoBatchSettingsSet struct {
	WalletId             [32]byte
	SourceId             [32]byte
	AccountAddress       string
	BatchSize            uint64
	BatchDurationSeconds uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUtxoBatchSettingsSet is a free log retrieval operation binding the contract event 0xcbe8cc700fee4dbb5181f3a0e1e11ae203f60dc344c20566a2c02663a685d2f3.
//
// Solidity: event UtxoBatchSettingsSet(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterUtxoBatchSettingsSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeePaymentsUtxoUtxoBatchSettingsSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "UtxoBatchSettingsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoUtxoBatchSettingsSetIterator{contract: _TeePaymentsUtxo.contract, event: "UtxoBatchSettingsSet", logs: logs, sub: sub}, nil
}

// WatchUtxoBatchSettingsSet is a free log subscription operation binding the contract event 0xcbe8cc700fee4dbb5181f3a0e1e11ae203f60dc344c20566a2c02663a685d2f3.
//
// Solidity: event UtxoBatchSettingsSet(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchUtxoBatchSettingsSet(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoUtxoBatchSettingsSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "UtxoBatchSettingsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoUtxoBatchSettingsSet)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "UtxoBatchSettingsSet", log); err != nil {
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

// ParseUtxoBatchSettingsSet is a log parse operation binding the contract event 0xcbe8cc700fee4dbb5181f3a0e1e11ae203f60dc344c20566a2c02663a685d2f3.
//
// Solidity: event UtxoBatchSettingsSet(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 batchSize, uint64 batchDurationSeconds)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParseUtxoBatchSettingsSet(log types.Log) (*TeePaymentsUtxoUtxoBatchSettingsSet, error) {
	event := new(TeePaymentsUtxoUtxoBatchSettingsSet)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "UtxoBatchSettingsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUtxoUtxoReplacementReadyIterator is returned from FilterUtxoReplacementReady and is used to iterate over the raw logs and unpacked data for UtxoReplacementReady events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUtxoReplacementReadyIterator struct {
	Event *TeePaymentsUtxoUtxoReplacementReady // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoUtxoReplacementReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoUtxoReplacementReady)
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
		it.Event = new(TeePaymentsUtxoUtxoReplacementReady)
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
func (it *TeePaymentsUtxoUtxoReplacementReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoUtxoReplacementReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoUtxoReplacementReady represents a UtxoReplacementReady event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUtxoReplacementReady struct {
	WalletId       [32]byte
	AccountHash    [32]byte
	BatchPaymentId uint64
	ReplacementId  uint64
	FirstPaymentId uint64
	PaymentCount   uint64
	Blocks         []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUtxoReplacementReady is a free log retrieval operation binding the contract event 0xead2b565baa9113ac15b2bafe426e5776645beea0e572837cb8b60d9a4ade04c.
//
// Solidity: event UtxoReplacementReady(bytes32 indexed walletId, bytes32 indexed accountHash, uint64 batchPaymentId, uint64 replacementId, uint64 firstPaymentId, uint64 paymentCount, uint256[] blocks)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterUtxoReplacementReady(opts *bind.FilterOpts, walletId [][32]byte, accountHash [][32]byte) (*TeePaymentsUtxoUtxoReplacementReadyIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountHashRule []interface{}
	for _, accountHashItem := range accountHash {
		accountHashRule = append(accountHashRule, accountHashItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "UtxoReplacementReady", walletIdRule, accountHashRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoUtxoReplacementReadyIterator{contract: _TeePaymentsUtxo.contract, event: "UtxoReplacementReady", logs: logs, sub: sub}, nil
}

// WatchUtxoReplacementReady is a free log subscription operation binding the contract event 0xead2b565baa9113ac15b2bafe426e5776645beea0e572837cb8b60d9a4ade04c.
//
// Solidity: event UtxoReplacementReady(bytes32 indexed walletId, bytes32 indexed accountHash, uint64 batchPaymentId, uint64 replacementId, uint64 firstPaymentId, uint64 paymentCount, uint256[] blocks)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchUtxoReplacementReady(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoUtxoReplacementReady, walletId [][32]byte, accountHash [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountHashRule []interface{}
	for _, accountHashItem := range accountHash {
		accountHashRule = append(accountHashRule, accountHashItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "UtxoReplacementReady", walletIdRule, accountHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoUtxoReplacementReady)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "UtxoReplacementReady", log); err != nil {
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

// ParseUtxoReplacementReady is a log parse operation binding the contract event 0xead2b565baa9113ac15b2bafe426e5776645beea0e572837cb8b60d9a4ade04c.
//
// Solidity: event UtxoReplacementReady(bytes32 indexed walletId, bytes32 indexed accountHash, uint64 batchPaymentId, uint64 replacementId, uint64 firstPaymentId, uint64 paymentCount, uint256[] blocks)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParseUtxoReplacementReady(log types.Log) (*TeePaymentsUtxoUtxoReplacementReady, error) {
	event := new(TeePaymentsUtxoUtxoReplacementReady)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "UtxoReplacementReady", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsUtxoUtxoReplacementStartedIterator is returned from FilterUtxoReplacementStarted and is used to iterate over the raw logs and unpacked data for UtxoReplacementStarted events raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUtxoReplacementStartedIterator struct {
	Event *TeePaymentsUtxoUtxoReplacementStarted // Event containing the contract specifics and raw log

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
func (it *TeePaymentsUtxoUtxoReplacementStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsUtxoUtxoReplacementStarted)
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
		it.Event = new(TeePaymentsUtxoUtxoReplacementStarted)
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
func (it *TeePaymentsUtxoUtxoReplacementStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsUtxoUtxoReplacementStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsUtxoUtxoReplacementStarted represents a UtxoReplacementStarted event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUtxoReplacementStarted struct {
	WalletId       [32]byte
	AccountHash    [32]byte
	BatchPaymentId uint64
	ReplacementId  uint64
	FirstPaymentId uint64
	StartBlock     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUtxoReplacementStarted is a free log retrieval operation binding the contract event 0x3b5a986b5bdb4613927f5a7b3a92d444dc1e3766bed93dc980a676d992640328.
//
// Solidity: event UtxoReplacementStarted(bytes32 indexed walletId, bytes32 indexed accountHash, uint64 batchPaymentId, uint64 replacementId, uint64 firstPaymentId, uint256 startBlock)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) FilterUtxoReplacementStarted(opts *bind.FilterOpts, walletId [][32]byte, accountHash [][32]byte) (*TeePaymentsUtxoUtxoReplacementStartedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountHashRule []interface{}
	for _, accountHashItem := range accountHash {
		accountHashRule = append(accountHashRule, accountHashItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.FilterLogs(opts, "UtxoReplacementStarted", walletIdRule, accountHashRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsUtxoUtxoReplacementStartedIterator{contract: _TeePaymentsUtxo.contract, event: "UtxoReplacementStarted", logs: logs, sub: sub}, nil
}

// WatchUtxoReplacementStarted is a free log subscription operation binding the contract event 0x3b5a986b5bdb4613927f5a7b3a92d444dc1e3766bed93dc980a676d992640328.
//
// Solidity: event UtxoReplacementStarted(bytes32 indexed walletId, bytes32 indexed accountHash, uint64 batchPaymentId, uint64 replacementId, uint64 firstPaymentId, uint256 startBlock)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) WatchUtxoReplacementStarted(opts *bind.WatchOpts, sink chan<- *TeePaymentsUtxoUtxoReplacementStarted, walletId [][32]byte, accountHash [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountHashRule []interface{}
	for _, accountHashItem := range accountHash {
		accountHashRule = append(accountHashRule, accountHashItem)
	}

	logs, sub, err := _TeePaymentsUtxo.contract.WatchLogs(opts, "UtxoReplacementStarted", walletIdRule, accountHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsUtxoUtxoReplacementStarted)
				if err := _TeePaymentsUtxo.contract.UnpackLog(event, "UtxoReplacementStarted", log); err != nil {
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

// ParseUtxoReplacementStarted is a log parse operation binding the contract event 0x3b5a986b5bdb4613927f5a7b3a92d444dc1e3766bed93dc980a676d992640328.
//
// Solidity: event UtxoReplacementStarted(bytes32 indexed walletId, bytes32 indexed accountHash, uint64 batchPaymentId, uint64 replacementId, uint64 firstPaymentId, uint256 startBlock)
func (_TeePaymentsUtxo *TeePaymentsUtxoFilterer) ParseUtxoReplacementStarted(log types.Log) (*TeePaymentsUtxoUtxoReplacementStarted, error) {
	event := new(TeePaymentsUtxoUtxoReplacementStarted)
	if err := _TeePaymentsUtxo.contract.UnpackLog(event, "UtxoReplacementStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
