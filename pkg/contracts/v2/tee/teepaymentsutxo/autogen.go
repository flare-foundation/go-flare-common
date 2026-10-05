// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepaymentsutxo

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
	Status         uint8
	AccountAddress string
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

// ITeePaymentsUtxoBatchRecord is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsUtxoBatchRecord struct {
	Nonce         uint64
	BatchEndTs    uint64
	PaymentCount  uint64
	AnchorIndex   uint32
	RewardEpochId *big.Int
}

// ITeePaymentsUtxoUtxoAnchorState is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsUtxoUtxoAnchorState struct {
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
var TeePaymentsUtxoMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountIndexMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorIndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"availableAt\",\"type\":\"uint64\"}],\"name\":\"AnchorNotReady\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuthorizationAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotYetEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchSizeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeFactor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxBatchSizeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoActiveReplacement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNewAnchors\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPaymentInstructions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuthorizationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMWMultisigAccountAddressAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMWMultisigAccountNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentNotInBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReissueRewardEpochChanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReplacementAlreadyFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ScheduledSignaturesUnsupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockCallNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongKeyType\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"anchorReuseDelaySeconds\",\"type\":\"uint64\"}],\"name\":\"AnchorReuseDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxBatchSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxBatchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"MaxBatchSettingsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"anchorCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizationAddress\",\"type\":\"address\"}],\"name\":\"PMWMultisigUtxoAccountAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"anchorCount\",\"type\":\"uint32\"}],\"name\":\"UtxoAnchorsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"UtxoBatchSettingsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchPaymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"replacementId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"firstPaymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"paymentCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"blocks\",\"type\":\"uint256[]\"}],\"name\":\"UtxoReplacementReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"batchPaymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"replacementId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"firstPaymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"}],\"name\":\"UtxoReplacementStarted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Anchor[]\",\"name\":\"anchors\",\"type\":\"tuple[]\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigUtxoConfigured.PMWMultisigUtxoStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"addAnchors\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Anchor[]\",\"name\":\"anchors\",\"type\":\"tuple[]\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigUtxoConfigured.PMWMultisigUtxoStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"name\":\"addPMWMultisigAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressValidator\",\"outputs\":[{\"internalType\":\"contractIAddressValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"contractIIFlareTeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_anchorIndex\",\"type\":\"uint256\"}],\"name\":\"getAnchor\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"nextNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"availableAt\",\"type\":\"uint64\"}],\"internalType\":\"structITeePaymentsUtxo.UtxoAnchorState\",\"name\":\"_anchor\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAnchorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getAnchorReuseDelay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAuthorizationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getBatchPaymentId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchPaymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_batchPaymentId\",\"type\":\"uint64\"}],\"name\":\"getBatchRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"paymentCount\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"anchorIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"internalType\":\"structITeePaymentsUtxo.BatchRecord\",\"name\":\"_batch\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_open\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getBatchSettings\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getMaxBatchSettings\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_maxBatchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_maxBatchDurationSeconds\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getNextPaymentId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nextPaymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"}],\"name\":\"getPaymentFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getPaymentHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_paymentHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAccounts\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getWalletId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePaymentsBase.PaymentInstruction\",\"name\":\"_paymentInstruction\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentModel\",\"outputs\":[{\"internalType\":\"enumPaymentModel\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_batchPaymentId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePaymentsBase.PaymentInstruction[]\",\"name\":\"_paymentInstructions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"maxFeePerPayment\",\"type\":\"uint256[]\"},{\"internalType\":\"int16[][]\",\"name\":\"factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"delaysSeconds\",\"type\":\"uint16[]\"}],\"internalType\":\"structITeePaymentsBase.ReissueFeeParams\",\"name\":\"_reissueFeeParams\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_startNew\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"reissue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_finalized\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_anchorReuseDelaySeconds\",\"type\":\"uint64\"}],\"name\":\"setAnchorReuseDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_batchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_batchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"setBatchSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_maxBatchSize\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_maxBatchDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"setMaxBatchSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsConfigVerifier\",\"outputs\":[{\"internalType\":\"contractITeePaymentsConfigVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsRegistry\",\"outputs\":[{\"internalType\":\"contractITeePaymentsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "TeePaymentsUtxo",
}

// TeePaymentsUtxo is an auto generated Go binding around an Ethereum contract.
type TeePaymentsUtxo struct {
	abi abi.ABI
}

// NewTeePaymentsUtxo creates a new instance of TeePaymentsUtxo.
func NewTeePaymentsUtxo() *TeePaymentsUtxo {
	parsed, err := TeePaymentsUtxoMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TeePaymentsUtxo{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TeePaymentsUtxo) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePaymentsUtxo *TeePaymentsUtxo) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("UPGRADE_INTERFACE_VERSION")
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := teePaymentsUtxo.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackAddAnchors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22344637.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addAnchors(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof) returns()
func (teePaymentsUtxo *TeePaymentsUtxo) PackAddAnchors(proof IPMWMultisigUtxoConfiguredProof) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("addAnchors", proof)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddAnchors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22344637.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addAnchors(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof) returns()
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackAddAnchors(proof IPMWMultisigUtxoConfiguredProof) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("addAnchors", proof)
}

// PackAddPMWMultisigAccount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87233bdc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addPMWMultisigAccount(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof, address _authorizationAddress) returns()
func (teePaymentsUtxo *TeePaymentsUtxo) PackAddPMWMultisigAccount(walletId [32]byte, proof IPMWMultisigUtxoConfiguredProof, authorizationAddress common.Address) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("addPMWMultisigAccount", walletId, proof, authorizationAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddPMWMultisigAccount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87233bdc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addPMWMultisigAccount(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof, address _authorizationAddress) returns()
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackAddPMWMultisigAccount(walletId [32]byte, proof IPMWMultisigUtxoConfiguredProof, authorizationAddress common.Address) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("addPMWMultisigAccount", walletId, proof, authorizationAddress)
}

// PackAddressValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe37facf4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addressValidator() view returns(address)
func (teePaymentsUtxo *TeePaymentsUtxo) PackAddressValidator() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("addressValidator")
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackAddressValidator() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("addressValidator")
}

// UnpackAddressValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe37facf4.
//
// Solidity: function addressValidator() view returns(address)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackAddressValidator(data []byte) (common.Address, error) {
	out, err := teePaymentsUtxo.abi.Unpack("addressValidator", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackCancelGovernanceCall(encodedCall []byte) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("cancelGovernanceCall", encodedCall)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackCancelGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("cancelGovernanceCall", encodedCall)
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (teePaymentsUtxo *TeePaymentsUtxo) PackExecuteGovernanceCall(encodedCall []byte) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("executeGovernanceCall", encodedCall)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackExecuteGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("executeGovernanceCall", encodedCall)
}

// PackFlareSystemsManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaae7fc9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareSystemsManager() view returns(address)
func (teePaymentsUtxo *TeePaymentsUtxo) PackFlareSystemsManager() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("flareSystemsManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFlareSystemsManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaae7fc9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function flareSystemsManager() view returns(address)
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackFlareSystemsManager() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("flareSystemsManager")
}

// UnpackFlareSystemsManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackFlareSystemsManager(data []byte) (common.Address, error) {
	out, err := teePaymentsUtxo.abi.Unpack("flareSystemsManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFlareTeeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x453f7ab4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareTeeManager() view returns(address)
func (teePaymentsUtxo *TeePaymentsUtxo) PackFlareTeeManager() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("flareTeeManager")
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackFlareTeeManager() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("flareTeeManager")
}

// UnpackFlareTeeManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackFlareTeeManager(data []byte) (common.Address, error) {
	out, err := teePaymentsUtxo.abi.Unpack("flareTeeManager", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetAddressUpdater() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getAddressUpdater")
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetAddressUpdater() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAnchor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf278814.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAnchor((bytes32,string) _account, uint256 _anchorIndex) view returns((bytes32,uint32,uint64,uint64) _anchor)
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetAnchor(account ITeePaymentsBasePMWMultisigAccount, anchorIndex *big.Int) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getAnchor", account, anchorIndex)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAnchor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaf278814.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAnchor((bytes32,string) _account, uint256 _anchorIndex) view returns((bytes32,uint32,uint64,uint64) _anchor)
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetAnchor(account ITeePaymentsBasePMWMultisigAccount, anchorIndex *big.Int) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getAnchor", account, anchorIndex)
}

// UnpackGetAnchor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaf278814.
//
// Solidity: function getAnchor((bytes32,string) _account, uint256 _anchorIndex) view returns((bytes32,uint32,uint64,uint64) _anchor)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetAnchor(data []byte) (ITeePaymentsUtxoUtxoAnchorState, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getAnchor", data)
	if err != nil {
		return *new(ITeePaymentsUtxoUtxoAnchorState), err
	}
	out0 := *abi.ConvertType(out[0], new(ITeePaymentsUtxoUtxoAnchorState)).(*ITeePaymentsUtxoUtxoAnchorState)
	return out0, nil
}

// PackGetAnchorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x00dc90fd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAnchorCount((bytes32,string) _account) view returns(uint256)
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetAnchorCount(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getAnchorCount", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAnchorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x00dc90fd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAnchorCount((bytes32,string) _account) view returns(uint256)
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetAnchorCount(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getAnchorCount", account)
}

// UnpackGetAnchorCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x00dc90fd.
//
// Solidity: function getAnchorCount((bytes32,string) _account) view returns(uint256)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetAnchorCount(data []byte) (*big.Int, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getAnchorCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetAnchorReuseDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2b57cf73.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAnchorReuseDelay(bytes32 _sourceId) view returns(uint64)
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetAnchorReuseDelay(sourceId [32]byte) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getAnchorReuseDelay", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAnchorReuseDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2b57cf73.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAnchorReuseDelay(bytes32 _sourceId) view returns(uint64)
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetAnchorReuseDelay(sourceId [32]byte) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getAnchorReuseDelay", sourceId)
}

// UnpackGetAnchorReuseDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2b57cf73.
//
// Solidity: function getAnchorReuseDelay(bytes32 _sourceId) view returns(uint64)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetAnchorReuseDelay(data []byte) (uint64, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getAnchorReuseDelay", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetAuthorizationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x410642e0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetAuthorizationAddress(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getAuthorizationAddress", account)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetAuthorizationAddress(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getAuthorizationAddress", account)
}

// UnpackGetAuthorizationAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetAuthorizationAddress(data []byte) (common.Address, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getAuthorizationAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetBatchPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41952376.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBatchPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(uint64 _batchPaymentId)
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetBatchPaymentId(account ITeePaymentsBasePMWMultisigAccount, paymentId uint64) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getBatchPaymentId", account, paymentId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBatchPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41952376.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBatchPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(uint64 _batchPaymentId)
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetBatchPaymentId(account ITeePaymentsBasePMWMultisigAccount, paymentId uint64) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getBatchPaymentId", account, paymentId)
}

// UnpackGetBatchPaymentId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x41952376.
//
// Solidity: function getBatchPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(uint64 _batchPaymentId)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetBatchPaymentId(data []byte) (uint64, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getBatchPaymentId", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetBatchRecord is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7729adb4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBatchRecord((bytes32,string) _account, uint64 _batchPaymentId) view returns((uint64,uint64,uint64,uint32,uint24) _batch, bool _open)
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetBatchRecord(account ITeePaymentsBasePMWMultisigAccount, batchPaymentId uint64) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getBatchRecord", account, batchPaymentId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBatchRecord is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7729adb4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBatchRecord((bytes32,string) _account, uint64 _batchPaymentId) view returns((uint64,uint64,uint64,uint32,uint24) _batch, bool _open)
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetBatchRecord(account ITeePaymentsBasePMWMultisigAccount, batchPaymentId uint64) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getBatchRecord", account, batchPaymentId)
}

// GetBatchRecordOutput serves as a container for the return parameters of contract
// method GetBatchRecord.
type GetBatchRecordOutput struct {
	Batch ITeePaymentsUtxoBatchRecord
	Open  bool
}

// UnpackGetBatchRecord is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7729adb4.
//
// Solidity: function getBatchRecord((bytes32,string) _account, uint64 _batchPaymentId) view returns((uint64,uint64,uint64,uint32,uint24) _batch, bool _open)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetBatchRecord(data []byte) (GetBatchRecordOutput, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getBatchRecord", data)
	outstruct := new(GetBatchRecordOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Batch = *abi.ConvertType(out[0], new(ITeePaymentsUtxoBatchRecord)).(*ITeePaymentsUtxoBatchRecord)
	outstruct.Open = *abi.ConvertType(out[1], new(bool)).(*bool)
	return *outstruct, nil
}

// PackGetBatchSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb96c3c8f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBatchSettings((bytes32,string) _account) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetBatchSettings(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getBatchSettings", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBatchSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb96c3c8f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBatchSettings((bytes32,string) _account) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetBatchSettings(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getBatchSettings", account)
}

// GetBatchSettingsOutput serves as a container for the return parameters of contract
// method GetBatchSettings.
type GetBatchSettingsOutput struct {
	BatchSize            uint64
	BatchDurationSeconds uint64
}

// UnpackGetBatchSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb96c3c8f.
//
// Solidity: function getBatchSettings((bytes32,string) _account) view returns(uint64 _batchSize, uint64 _batchDurationSeconds)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetBatchSettings(data []byte) (GetBatchSettingsOutput, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getBatchSettings", data)
	outstruct := new(GetBatchSettingsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.BatchSize = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.BatchDurationSeconds = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetMaxBatchSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf6cc38e1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMaxBatchSettings(bytes32 _sourceId) view returns(uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds)
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetMaxBatchSettings(sourceId [32]byte) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getMaxBatchSettings", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMaxBatchSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf6cc38e1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMaxBatchSettings(bytes32 _sourceId) view returns(uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds)
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetMaxBatchSettings(sourceId [32]byte) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getMaxBatchSettings", sourceId)
}

// GetMaxBatchSettingsOutput serves as a container for the return parameters of contract
// method GetMaxBatchSettings.
type GetMaxBatchSettingsOutput struct {
	MaxBatchSize            uint64
	MaxBatchDurationSeconds uint64
}

// UnpackGetMaxBatchSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf6cc38e1.
//
// Solidity: function getMaxBatchSettings(bytes32 _sourceId) view returns(uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetMaxBatchSettings(data []byte) (GetMaxBatchSettingsOutput, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getMaxBatchSettings", data)
	outstruct := new(GetMaxBatchSettingsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.MaxBatchSize = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.MaxBatchDurationSeconds = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetNextPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfb49ac30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNextPaymentId((bytes32,string) _account) view returns(uint64 _nextPaymentId)
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetNextPaymentId(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getNextPaymentId", account)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetNextPaymentId(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getNextPaymentId", account)
}

// UnpackGetNextPaymentId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfb49ac30.
//
// Solidity: function getNextPaymentId((bytes32,string) _account) view returns(uint64 _nextPaymentId)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetNextPaymentId(data []byte) (uint64, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getNextPaymentId", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetPaymentFee(account ITeePaymentsBasePMWMultisigAccount, opCommand [32]byte) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getPaymentFee", account, opCommand)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetPaymentFee(account ITeePaymentsBasePMWMultisigAccount, opCommand [32]byte) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getPaymentFee", account, opCommand)
}

// UnpackGetPaymentFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x57abf78b.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetPaymentFee(data []byte) (*big.Int, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getPaymentFee", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetPaymentHash(account ITeePaymentsBasePMWMultisigAccount, paymentId uint64) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getPaymentHash", account, paymentId)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetPaymentHash(account ITeePaymentsBasePMWMultisigAccount, paymentId uint64) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getPaymentHash", account, paymentId)
}

// UnpackGetPaymentHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0079dfe8.
//
// Solidity: function getPaymentHash((bytes32,string) _account, uint64 _paymentId) view returns(bytes32 _paymentHash)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetPaymentHash(data []byte) ([32]byte, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getPaymentHash", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetWalletAccounts(walletId [32]byte) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getWalletAccounts", walletId)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetWalletAccounts(walletId [32]byte) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getWalletAccounts", walletId)
}

// UnpackGetWalletAccounts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetWalletAccounts(data []byte) ([]ITeePaymentsBasePMWMultisigAccount, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getWalletAccounts", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackGetWalletId(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("getWalletId", account)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGetWalletId(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("getWalletId", account)
}

// UnpackGetWalletId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGetWalletId(data []byte) ([32]byte, error) {
	out, err := teePaymentsUtxo.abi.Unpack("getWalletId", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackGovernance() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("governance")
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGovernance() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := teePaymentsUtxo.abi.Unpack("governance", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackGovernanceSettings() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("governanceSettings")
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackGovernanceSettings() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := teePaymentsUtxo.abi.Unpack("governanceSettings", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackImplementation() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("implementation")
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackImplementation() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("implementation")
}

// UnpackImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackImplementation(data []byte) (common.Address, error) {
	out, err := teePaymentsUtxo.abi.Unpack("implementation", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teePaymentsUtxo *TeePaymentsUtxo) PackIsExecutor(address common.Address) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("isExecutor", address)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := teePaymentsUtxo.abi.Unpack("isExecutor", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackPay(account ITeePaymentsBasePMWMultisigAccount, paymentInstruction ITeePaymentsBasePaymentInstruction, claimBackAddress common.Address) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("pay", account, paymentInstruction, claimBackAddress)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackPay(account ITeePaymentsBasePMWMultisigAccount, paymentInstruction ITeePaymentsBasePaymentInstruction, claimBackAddress common.Address) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("pay", account, paymentInstruction, claimBackAddress)
}

// UnpackPay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackPay(data []byte) (uint64, error) {
	out, err := teePaymentsUtxo.abi.Unpack("pay", data)
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
// Solidity: function paymentModel() pure returns(uint8)
func (teePaymentsUtxo *TeePaymentsUtxo) PackPaymentModel() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("paymentModel")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPaymentModel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb9d3dbb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function paymentModel() pure returns(uint8)
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackPaymentModel() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("paymentModel")
}

// UnpackPaymentModel is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbb9d3dbb.
//
// Solidity: function paymentModel() pure returns(uint8)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackPaymentModel(data []byte) (uint8, error) {
	out, err := teePaymentsUtxo.abi.Unpack("paymentModel", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackProductionMode() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("productionMode")
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackProductionMode() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackProductionMode(data []byte) (bool, error) {
	out, err := teePaymentsUtxo.abi.Unpack("productionMode", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackProxiableUUID() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("proxiableUUID")
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackProxiableUUID() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("proxiableUUID")
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := teePaymentsUtxo.abi.Unpack("proxiableUUID", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackReissue(account ITeePaymentsBasePMWMultisigAccount, batchPaymentId uint64, paymentInstructions []ITeePaymentsBasePaymentInstruction, reissueFeeParams ITeePaymentsBaseReissueFeeParams, startNew bool, claimBackAddress common.Address) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("reissue", account, batchPaymentId, paymentInstructions, reissueFeeParams, startNew, claimBackAddress)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackReissue(account ITeePaymentsBasePMWMultisigAccount, batchPaymentId uint64, paymentInstructions []ITeePaymentsBasePaymentInstruction, reissueFeeParams ITeePaymentsBaseReissueFeeParams, startNew bool, claimBackAddress common.Address) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("reissue", account, batchPaymentId, paymentInstructions, reissueFeeParams, startNew, claimBackAddress)
}

// UnpackReissue is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1abcad2c.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _batchPaymentId, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, bool _startNew, address _claimBackAddress) payable returns(bool _finalized)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackReissue(data []byte) (bool, error) {
	out, err := teePaymentsUtxo.abi.Unpack("reissue", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSetAnchorReuseDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x729a0c5a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setAnchorReuseDelay(bytes32 _sourceId, uint64 _anchorReuseDelaySeconds) returns()
func (teePaymentsUtxo *TeePaymentsUtxo) PackSetAnchorReuseDelay(sourceId [32]byte, anchorReuseDelaySeconds uint64) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("setAnchorReuseDelay", sourceId, anchorReuseDelaySeconds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetAnchorReuseDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x729a0c5a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setAnchorReuseDelay(bytes32 _sourceId, uint64 _anchorReuseDelaySeconds) returns()
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackSetAnchorReuseDelay(sourceId [32]byte, anchorReuseDelaySeconds uint64) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("setAnchorReuseDelay", sourceId, anchorReuseDelaySeconds)
}

// PackSetBatchSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd302be77.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setBatchSettings((bytes32,string) _account, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (teePaymentsUtxo *TeePaymentsUtxo) PackSetBatchSettings(account ITeePaymentsBasePMWMultisigAccount, batchSize uint64, batchDurationSeconds uint64) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("setBatchSettings", account, batchSize, batchDurationSeconds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetBatchSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd302be77.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setBatchSettings((bytes32,string) _account, uint64 _batchSize, uint64 _batchDurationSeconds) returns()
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackSetBatchSettings(account ITeePaymentsBasePMWMultisigAccount, batchSize uint64, batchDurationSeconds uint64) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("setBatchSettings", account, batchSize, batchDurationSeconds)
}

// PackSetMaxBatchSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcb42368a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMaxBatchSettings(bytes32 _sourceId, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds) returns()
func (teePaymentsUtxo *TeePaymentsUtxo) PackSetMaxBatchSettings(sourceId [32]byte, maxBatchSize uint64, maxBatchDurationSeconds uint64) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("setMaxBatchSettings", sourceId, maxBatchSize, maxBatchDurationSeconds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMaxBatchSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcb42368a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMaxBatchSettings(bytes32 _sourceId, uint64 _maxBatchSize, uint64 _maxBatchDurationSeconds) returns()
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackSetMaxBatchSettings(sourceId [32]byte, maxBatchSize uint64, maxBatchDurationSeconds uint64) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("setMaxBatchSettings", sourceId, maxBatchSize, maxBatchDurationSeconds)
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (teePaymentsUtxo *TeePaymentsUtxo) PackSwitchToProductionMode() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("switchToProductionMode")
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackSwitchToProductionMode() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("switchToProductionMode")
}

// PackTeePaymentsConfigVerifier is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf71c6c75.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (teePaymentsUtxo *TeePaymentsUtxo) PackTeePaymentsConfigVerifier() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("teePaymentsConfigVerifier")
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackTeePaymentsConfigVerifier() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("teePaymentsConfigVerifier")
}

// UnpackTeePaymentsConfigVerifier is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf71c6c75.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackTeePaymentsConfigVerifier(data []byte) (common.Address, error) {
	out, err := teePaymentsUtxo.abi.Unpack("teePaymentsConfigVerifier", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackTeePaymentsRegistry() []byte {
	enc, err := teePaymentsUtxo.abi.Pack("teePaymentsRegistry")
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackTeePaymentsRegistry() ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("teePaymentsRegistry")
}

// UnpackTeePaymentsRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackTeePaymentsRegistry(data []byte) (common.Address, error) {
	out, err := teePaymentsUtxo.abi.Unpack("teePaymentsRegistry", data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (teePaymentsUtxo *TeePaymentsUtxo) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := teePaymentsUtxo.abi.Pack("upgradeToAndCall", newImplementation, data)
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
func (teePaymentsUtxo *TeePaymentsUtxo) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return teePaymentsUtxo.abi.Pack("upgradeToAndCall", newImplementation, data)
}

// TeePaymentsUtxoAnchorReuseDelaySet represents a AnchorReuseDelaySet event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoAnchorReuseDelaySet struct {
	SourceId                [32]byte
	AnchorReuseDelaySeconds uint64
	Raw                     *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoAnchorReuseDelaySetEventName = "AnchorReuseDelaySet"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoAnchorReuseDelaySet) ContractEventName() string {
	return TeePaymentsUtxoAnchorReuseDelaySetEventName
}

// UnpackAnchorReuseDelaySetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AnchorReuseDelaySet(bytes32 indexed sourceId, uint64 anchorReuseDelaySeconds)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackAnchorReuseDelaySetEvent(log *types.Log) (*TeePaymentsUtxoAnchorReuseDelaySet, error) {
	event := "AnchorReuseDelaySet"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoAnchorReuseDelaySet)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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

// TeePaymentsUtxoGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoGovernanceCallTimelocked) ContractEventName() string {
	return TeePaymentsUtxoGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*TeePaymentsUtxoGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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

// TeePaymentsUtxoGovernanceInitialised represents a GovernanceInitialised event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoGovernanceInitialised) ContractEventName() string {
	return TeePaymentsUtxoGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGovernanceInitialisedEvent(log *types.Log) (*TeePaymentsUtxoGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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

// TeePaymentsUtxoGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoGovernedProductionModeEntered) ContractEventName() string {
	return TeePaymentsUtxoGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*TeePaymentsUtxoGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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

// TeePaymentsUtxoInitialized represents a Initialized event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoInitialized) ContractEventName() string {
	return TeePaymentsUtxoInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackInitializedEvent(log *types.Log) (*TeePaymentsUtxoInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoInitialized)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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

// TeePaymentsUtxoMaxBatchSettingsSet represents a MaxBatchSettingsSet event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoMaxBatchSettingsSet struct {
	SourceId                [32]byte
	MaxBatchSize            uint64
	MaxBatchDurationSeconds uint64
	Raw                     *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoMaxBatchSettingsSetEventName = "MaxBatchSettingsSet"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoMaxBatchSettingsSet) ContractEventName() string {
	return TeePaymentsUtxoMaxBatchSettingsSetEventName
}

// UnpackMaxBatchSettingsSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MaxBatchSettingsSet(bytes32 indexed sourceId, uint64 maxBatchSize, uint64 maxBatchDurationSeconds)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackMaxBatchSettingsSetEvent(log *types.Log) (*TeePaymentsUtxoMaxBatchSettingsSet, error) {
	event := "MaxBatchSettingsSet"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoMaxBatchSettingsSet)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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

// TeePaymentsUtxoPMWMultisigUtxoAccountAdded represents a PMWMultisigUtxoAccountAdded event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoPMWMultisigUtxoAccountAdded struct {
	WalletId             [32]byte
	SourceId             [32]byte
	AccountAddress       string
	AccountIndex         uint32
	AnchorCount          uint32
	AuthorizationAddress common.Address
	Raw                  *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoPMWMultisigUtxoAccountAddedEventName = "PMWMultisigUtxoAccountAdded"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoPMWMultisigUtxoAccountAdded) ContractEventName() string {
	return TeePaymentsUtxoPMWMultisigUtxoAccountAddedEventName
}

// UnpackPMWMultisigUtxoAccountAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PMWMultisigUtxoAccountAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount, address authorizationAddress)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackPMWMultisigUtxoAccountAddedEvent(log *types.Log) (*TeePaymentsUtxoPMWMultisigUtxoAccountAdded, error) {
	event := "PMWMultisigUtxoAccountAdded"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoPMWMultisigUtxoAccountAdded)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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

// TeePaymentsUtxoTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoTimelockedGovernanceCallCanceled) ContractEventName() string {
	return TeePaymentsUtxoTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*TeePaymentsUtxoTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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

// TeePaymentsUtxoTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoTimelockedGovernanceCallExecuted) ContractEventName() string {
	return TeePaymentsUtxoTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*TeePaymentsUtxoTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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

// TeePaymentsUtxoUpgraded represents a Upgraded event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoUpgraded) ContractEventName() string {
	return TeePaymentsUtxoUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackUpgradedEvent(log *types.Log) (*TeePaymentsUtxoUpgraded, error) {
	event := "Upgraded"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoUpgraded)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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

// TeePaymentsUtxoUtxoAnchorsAdded represents a UtxoAnchorsAdded event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUtxoAnchorsAdded struct {
	WalletId       [32]byte
	SourceId       [32]byte
	AccountAddress string
	AccountIndex   uint32
	AnchorCount    uint32
	Raw            *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoUtxoAnchorsAddedEventName = "UtxoAnchorsAdded"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoUtxoAnchorsAdded) ContractEventName() string {
	return TeePaymentsUtxoUtxoAnchorsAddedEventName
}

// UnpackUtxoAnchorsAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UtxoAnchorsAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackUtxoAnchorsAddedEvent(log *types.Log) (*TeePaymentsUtxoUtxoAnchorsAdded, error) {
	event := "UtxoAnchorsAdded"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoUtxoAnchorsAdded)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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

// TeePaymentsUtxoUtxoBatchSettingsSet represents a UtxoBatchSettingsSet event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUtxoBatchSettingsSet struct {
	WalletId             [32]byte
	SourceId             [32]byte
	AccountAddress       string
	BatchSize            uint64
	BatchDurationSeconds uint64
	Raw                  *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoUtxoBatchSettingsSetEventName = "UtxoBatchSettingsSet"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoUtxoBatchSettingsSet) ContractEventName() string {
	return TeePaymentsUtxoUtxoBatchSettingsSetEventName
}

// UnpackUtxoBatchSettingsSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UtxoBatchSettingsSet(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint64 batchSize, uint64 batchDurationSeconds)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackUtxoBatchSettingsSetEvent(log *types.Log) (*TeePaymentsUtxoUtxoBatchSettingsSet, error) {
	event := "UtxoBatchSettingsSet"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoUtxoBatchSettingsSet)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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

// TeePaymentsUtxoUtxoReplacementReady represents a UtxoReplacementReady event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUtxoReplacementReady struct {
	WalletId       [32]byte
	AccountHash    [32]byte
	BatchPaymentId uint64
	ReplacementId  uint64
	FirstPaymentId uint64
	PaymentCount   uint64
	Blocks         []*big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoUtxoReplacementReadyEventName = "UtxoReplacementReady"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoUtxoReplacementReady) ContractEventName() string {
	return TeePaymentsUtxoUtxoReplacementReadyEventName
}

// UnpackUtxoReplacementReadyEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UtxoReplacementReady(bytes32 indexed walletId, bytes32 indexed accountHash, uint64 batchPaymentId, uint64 replacementId, uint64 firstPaymentId, uint64 paymentCount, uint256[] blocks)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackUtxoReplacementReadyEvent(log *types.Log) (*TeePaymentsUtxoUtxoReplacementReady, error) {
	event := "UtxoReplacementReady"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoUtxoReplacementReady)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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

// TeePaymentsUtxoUtxoReplacementStarted represents a UtxoReplacementStarted event raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUtxoReplacementStarted struct {
	WalletId       [32]byte
	AccountHash    [32]byte
	BatchPaymentId uint64
	ReplacementId  uint64
	FirstPaymentId uint64
	StartBlock     *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUtxoUtxoReplacementStartedEventName = "UtxoReplacementStarted"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUtxoUtxoReplacementStarted) ContractEventName() string {
	return TeePaymentsUtxoUtxoReplacementStartedEventName
}

// UnpackUtxoReplacementStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UtxoReplacementStarted(bytes32 indexed walletId, bytes32 indexed accountHash, uint64 batchPaymentId, uint64 replacementId, uint64 firstPaymentId, uint256 startBlock)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackUtxoReplacementStartedEvent(log *types.Log) (*TeePaymentsUtxoUtxoReplacementStarted, error) {
	event := "UtxoReplacementStarted"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsUtxo.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUtxoUtxoReplacementStarted)
	if len(log.Data) > 0 {
		if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsUtxo.abi.Events[event].Inputs {
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
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["AccountIndexMismatch"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackAccountIndexMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["AlreadyInProductionMode"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackAlreadyInProductionModeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["AnchorIndexOutOfBounds"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackAnchorIndexOutOfBoundsError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["AnchorMismatch"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackAnchorMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["AnchorNotReady"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackAnchorNotReadyError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["AuthorizationAddressZero"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackAuthorizationAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["BatchNotYetEnded"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackBatchNotYetEndedError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["BatchSizeZero"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackBatchSizeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["GovernedAddressZero"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackGovernedAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["GovernedAlreadyInitialized"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackGovernedAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["InvalidFeeFactor"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackInvalidFeeFactorError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["InvalidPaymentId"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackInvalidPaymentIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["InvalidRecipientAddress"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackInvalidRecipientAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["MaxBatchSizeZero"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackMaxBatchSizeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["NoActiveReplacement"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackNoActiveReplacementError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["NoNewAnchors"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackNoNewAnchorsError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["NoPaymentInstructions"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackNoPaymentInstructionsError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["OnlyAuthorizationAddress"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackOnlyAuthorizationAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["OnlyExecutor"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["OnlySystemExtensionId"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackOnlySystemExtensionIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["OnlyWalletOwner"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackOnlyWalletOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["PMWMultisigAccountAddressAlreadySet"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackPMWMultisigAccountAddressAlreadySetError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["PMWMultisigAccountNotRegistered"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackPMWMultisigAccountNotRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["PaymentAmountZero"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackPaymentAmountZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["PaymentHashMismatch"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackPaymentHashMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["PaymentNotInBatch"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackPaymentNotInBatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["ReissueRewardEpochChanged"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackReissueRewardEpochChangedError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["ReplacementAlreadyFinalized"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackReplacementAlreadyFinalizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["ScheduledSignaturesUnsupported"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackScheduledSignaturesUnsupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["TimelockCallNotFound"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackTimelockCallNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["TimelockNotAllowedYet"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackTimelockNotAllowedYetError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["UnsupportedSourceId"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackUnsupportedSourceIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["WalletNotInProduction"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackWalletNotInProductionError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsUtxo.abi.Errors["WrongKeyType"].ID.Bytes()[:4]) {
		return teePaymentsUtxo.UnpackWrongKeyTypeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// TeePaymentsUtxoAccountIndexMismatch represents a AccountIndexMismatch error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoAccountIndexMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccountIndexMismatch()
func TeePaymentsUtxoAccountIndexMismatchErrorID() common.Hash {
	return common.HexToHash("0x27313024826f37a5dc59237f5f36a6005c35c5374c7feb74c5211b91b777fe75")
}

// UnpackAccountIndexMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccountIndexMismatch()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackAccountIndexMismatchError(raw []byte) (*TeePaymentsUtxoAccountIndexMismatch, error) {
	out := new(TeePaymentsUtxoAccountIndexMismatch)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "AccountIndexMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoAddressEmptyCode represents a AddressEmptyCode error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func TeePaymentsUtxoAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackAddressEmptyCodeError(raw []byte) (*TeePaymentsUtxoAddressEmptyCode, error) {
	out := new(TeePaymentsUtxoAddressEmptyCode)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoAlreadyInProductionMode represents a AlreadyInProductionMode error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoAlreadyInProductionMode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInProductionMode()
func TeePaymentsUtxoAlreadyInProductionModeErrorID() common.Hash {
	return common.HexToHash("0x72a23ec2b1fdbc3de7d5f411b572fcf8ad8409803196930f63ddf40a2a2b2594")
}

// UnpackAlreadyInProductionModeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInProductionMode()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackAlreadyInProductionModeError(raw []byte) (*TeePaymentsUtxoAlreadyInProductionMode, error) {
	out := new(TeePaymentsUtxoAlreadyInProductionMode)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "AlreadyInProductionMode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoAnchorIndexOutOfBounds represents a AnchorIndexOutOfBounds error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoAnchorIndexOutOfBounds struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AnchorIndexOutOfBounds()
func TeePaymentsUtxoAnchorIndexOutOfBoundsErrorID() common.Hash {
	return common.HexToHash("0x2de58279687b2e18582d62fa0fa14630123921cc78446039d46922005d9e20a6")
}

// UnpackAnchorIndexOutOfBoundsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AnchorIndexOutOfBounds()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackAnchorIndexOutOfBoundsError(raw []byte) (*TeePaymentsUtxoAnchorIndexOutOfBounds, error) {
	out := new(TeePaymentsUtxoAnchorIndexOutOfBounds)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "AnchorIndexOutOfBounds", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoAnchorMismatch represents a AnchorMismatch error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoAnchorMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AnchorMismatch()
func TeePaymentsUtxoAnchorMismatchErrorID() common.Hash {
	return common.HexToHash("0x53be431d1f094556035616469b0ed7a4a966a3bb752407dd890067a5ff782e0b")
}

// UnpackAnchorMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AnchorMismatch()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackAnchorMismatchError(raw []byte) (*TeePaymentsUtxoAnchorMismatch, error) {
	out := new(TeePaymentsUtxoAnchorMismatch)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "AnchorMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoAnchorNotReady represents a AnchorNotReady error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoAnchorNotReady struct {
	AvailableAt uint64
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AnchorNotReady(uint64 availableAt)
func TeePaymentsUtxoAnchorNotReadyErrorID() common.Hash {
	return common.HexToHash("0xf709bc644b6eb5014a7ada4a6ca63ad922c1239bcdb0c4667f898c1f782d562e")
}

// UnpackAnchorNotReadyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AnchorNotReady(uint64 availableAt)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackAnchorNotReadyError(raw []byte) (*TeePaymentsUtxoAnchorNotReady, error) {
	out := new(TeePaymentsUtxoAnchorNotReady)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "AnchorNotReady", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoAuthorizationAddressZero represents a AuthorizationAddressZero error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoAuthorizationAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AuthorizationAddressZero()
func TeePaymentsUtxoAuthorizationAddressZeroErrorID() common.Hash {
	return common.HexToHash("0xe4472ca7994bab3d6f38c9bdc243c438c59c017951973a129c8674da85c299a2")
}

// UnpackAuthorizationAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AuthorizationAddressZero()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackAuthorizationAddressZeroError(raw []byte) (*TeePaymentsUtxoAuthorizationAddressZero, error) {
	out := new(TeePaymentsUtxoAuthorizationAddressZero)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "AuthorizationAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoBatchNotYetEnded represents a BatchNotYetEnded error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoBatchNotYetEnded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BatchNotYetEnded()
func TeePaymentsUtxoBatchNotYetEndedErrorID() common.Hash {
	return common.HexToHash("0x2ec22b8c0d5d91f85ff6b0731315817f219ee094d28bdc47e320d1c9d6a1fb1b")
}

// UnpackBatchNotYetEndedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BatchNotYetEnded()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackBatchNotYetEndedError(raw []byte) (*TeePaymentsUtxoBatchNotYetEnded, error) {
	out := new(TeePaymentsUtxoBatchNotYetEnded)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "BatchNotYetEnded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoBatchSizeZero represents a BatchSizeZero error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoBatchSizeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BatchSizeZero()
func TeePaymentsUtxoBatchSizeZeroErrorID() common.Hash {
	return common.HexToHash("0x032c1d2a2f459e21459554a4d63ba77fdd3530203b8c6f35d7a243651afa832a")
}

// UnpackBatchSizeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BatchSizeZero()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackBatchSizeZeroError(raw []byte) (*TeePaymentsUtxoBatchSizeZero, error) {
	out := new(TeePaymentsUtxoBatchSizeZero)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "BatchSizeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func TeePaymentsUtxoERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackERC1967InvalidImplementationError(raw []byte) (*TeePaymentsUtxoERC1967InvalidImplementation, error) {
	out := new(TeePaymentsUtxoERC1967InvalidImplementation)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoERC1967NonPayable represents a ERC1967NonPayable error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func TeePaymentsUtxoERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackERC1967NonPayableError(raw []byte) (*TeePaymentsUtxoERC1967NonPayable, error) {
	out := new(TeePaymentsUtxoERC1967NonPayable)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoFailedCall represents a FailedCall error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func TeePaymentsUtxoFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackFailedCallError(raw []byte) (*TeePaymentsUtxoFailedCall, error) {
	out := new(TeePaymentsUtxoFailedCall)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoGovernedAddressZero represents a GovernedAddressZero error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoGovernedAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAddressZero()
func TeePaymentsUtxoGovernedAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x695725f6b429bac0cc5a2fa88214e31520e0c69f2fb0e2fb881a7130a564f49f")
}

// UnpackGovernedAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAddressZero()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGovernedAddressZeroError(raw []byte) (*TeePaymentsUtxoGovernedAddressZero, error) {
	out := new(TeePaymentsUtxoGovernedAddressZero)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "GovernedAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoGovernedAlreadyInitialized represents a GovernedAlreadyInitialized error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoGovernedAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAlreadyInitialized()
func TeePaymentsUtxoGovernedAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0xb2f0162156d83828197776d56bbfba0a7a9d5936fad70f0a3014ae18cc7aa33c")
}

// UnpackGovernedAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAlreadyInitialized()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackGovernedAlreadyInitializedError(raw []byte) (*TeePaymentsUtxoGovernedAlreadyInitialized, error) {
	out := new(TeePaymentsUtxoGovernedAlreadyInitialized)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "GovernedAlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoInvalidFeeFactor represents a InvalidFeeFactor error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoInvalidFeeFactor struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeFactor(uint256 index)
func TeePaymentsUtxoInvalidFeeFactorErrorID() common.Hash {
	return common.HexToHash("0x89619400c26b8355799d991b2b4ac022351c0c4f21a4c0f3b4ee2ef70d0c8737")
}

// UnpackInvalidFeeFactorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeFactor(uint256 index)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackInvalidFeeFactorError(raw []byte) (*TeePaymentsUtxoInvalidFeeFactor, error) {
	out := new(TeePaymentsUtxoInvalidFeeFactor)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "InvalidFeeFactor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoInvalidInitialization represents a InvalidInitialization error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func TeePaymentsUtxoInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackInvalidInitializationError(raw []byte) (*TeePaymentsUtxoInvalidInitialization, error) {
	out := new(TeePaymentsUtxoInvalidInitialization)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoInvalidPaymentId represents a InvalidPaymentId error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoInvalidPaymentId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPaymentId()
func TeePaymentsUtxoInvalidPaymentIdErrorID() common.Hash {
	return common.HexToHash("0xb11be8c4d37d5cf7116085b4f06afe80345db997f69552f4e6a826a12eea0c0f")
}

// UnpackInvalidPaymentIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPaymentId()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackInvalidPaymentIdError(raw []byte) (*TeePaymentsUtxoInvalidPaymentId, error) {
	out := new(TeePaymentsUtxoInvalidPaymentId)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "InvalidPaymentId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoInvalidRecipientAddress represents a InvalidRecipientAddress error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoInvalidRecipientAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRecipientAddress()
func TeePaymentsUtxoInvalidRecipientAddressErrorID() common.Hash {
	return common.HexToHash("0x44d99fea8d272f0b82a2cdf27d16958dace6a85ad6302ff443201a39bfc4820a")
}

// UnpackInvalidRecipientAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRecipientAddress()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackInvalidRecipientAddressError(raw []byte) (*TeePaymentsUtxoInvalidRecipientAddress, error) {
	out := new(TeePaymentsUtxoInvalidRecipientAddress)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "InvalidRecipientAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoLengthsMismatch represents a LengthsMismatch error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func TeePaymentsUtxoLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackLengthsMismatchError(raw []byte) (*TeePaymentsUtxoLengthsMismatch, error) {
	out := new(TeePaymentsUtxoLengthsMismatch)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoMaxBatchSizeZero represents a MaxBatchSizeZero error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoMaxBatchSizeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MaxBatchSizeZero()
func TeePaymentsUtxoMaxBatchSizeZeroErrorID() common.Hash {
	return common.HexToHash("0x001cacce9473248782736bba2fe090bfa6bf7a8a235aa347fbb7caa70ccaa2ea")
}

// UnpackMaxBatchSizeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MaxBatchSizeZero()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackMaxBatchSizeZeroError(raw []byte) (*TeePaymentsUtxoMaxBatchSizeZero, error) {
	out := new(TeePaymentsUtxoMaxBatchSizeZero)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "MaxBatchSizeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoNoActiveReplacement represents a NoActiveReplacement error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoNoActiveReplacement struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoActiveReplacement()
func TeePaymentsUtxoNoActiveReplacementErrorID() common.Hash {
	return common.HexToHash("0xf98592b32bcb27eecbd55257176f90b2be7305ef1268e9eaf7e7b78e45b73ca5")
}

// UnpackNoActiveReplacementError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoActiveReplacement()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackNoActiveReplacementError(raw []byte) (*TeePaymentsUtxoNoActiveReplacement, error) {
	out := new(TeePaymentsUtxoNoActiveReplacement)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "NoActiveReplacement", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoNoNewAnchors represents a NoNewAnchors error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoNoNewAnchors struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoNewAnchors()
func TeePaymentsUtxoNoNewAnchorsErrorID() common.Hash {
	return common.HexToHash("0xb01793c2dd316a7313f6307f05a3be47b867e01e1f48b301608f1f92885c2d6d")
}

// UnpackNoNewAnchorsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoNewAnchors()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackNoNewAnchorsError(raw []byte) (*TeePaymentsUtxoNoNewAnchors, error) {
	out := new(TeePaymentsUtxoNoNewAnchors)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "NoNewAnchors", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoNoPaymentInstructions represents a NoPaymentInstructions error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoNoPaymentInstructions struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoPaymentInstructions()
func TeePaymentsUtxoNoPaymentInstructionsErrorID() common.Hash {
	return common.HexToHash("0xf77ba6c12a501330783ab3f6aed0d4dc523570aee3e8753b61dc0e992206dd6c")
}

// UnpackNoPaymentInstructionsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoPaymentInstructions()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackNoPaymentInstructionsError(raw []byte) (*TeePaymentsUtxoNoPaymentInstructions, error) {
	out := new(TeePaymentsUtxoNoPaymentInstructions)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "NoPaymentInstructions", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoNotInitializing represents a NotInitializing error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func TeePaymentsUtxoNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackNotInitializingError(raw []byte) (*TeePaymentsUtxoNotInitializing, error) {
	out := new(TeePaymentsUtxoNotInitializing)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoOnlyAuthorizationAddress represents a OnlyAuthorizationAddress error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoOnlyAuthorizationAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyAuthorizationAddress()
func TeePaymentsUtxoOnlyAuthorizationAddressErrorID() common.Hash {
	return common.HexToHash("0xdc65e71c09e77be7539dadee1dc86849e01c3e995f0c8d49e625476c2a1b1ed5")
}

// UnpackOnlyAuthorizationAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyAuthorizationAddress()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackOnlyAuthorizationAddressError(raw []byte) (*TeePaymentsUtxoOnlyAuthorizationAddress, error) {
	out := new(TeePaymentsUtxoOnlyAuthorizationAddress)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "OnlyAuthorizationAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoOnlyExecutor represents a OnlyExecutor error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoOnlyExecutor struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExecutor()
func TeePaymentsUtxoOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x7fb6be02aaf40e420b0e7eb67e43a3feee949e8932a3c64e33b591a436670841")
}

// UnpackOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExecutor()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackOnlyExecutorError(raw []byte) (*TeePaymentsUtxoOnlyExecutor, error) {
	out := new(TeePaymentsUtxoOnlyExecutor)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "OnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoOnlyGovernance represents a OnlyGovernance error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func TeePaymentsUtxoOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackOnlyGovernanceError(raw []byte) (*TeePaymentsUtxoOnlyGovernance, error) {
	out := new(TeePaymentsUtxoOnlyGovernance)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func TeePaymentsUtxoOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*TeePaymentsUtxoOnlyProductionOrPausedStatus, error) {
	out := new(TeePaymentsUtxoOnlyProductionOrPausedStatus)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoOnlySystemExtensionId represents a OnlySystemExtensionId error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoOnlySystemExtensionId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemExtensionId()
func TeePaymentsUtxoOnlySystemExtensionIdErrorID() common.Hash {
	return common.HexToHash("0xbd49094d5a86ee3f3c4b39274f5f98391dde2cd16618a917e121333ac3d43ee4")
}

// UnpackOnlySystemExtensionIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemExtensionId()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackOnlySystemExtensionIdError(raw []byte) (*TeePaymentsUtxoOnlySystemExtensionId, error) {
	out := new(TeePaymentsUtxoOnlySystemExtensionId)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "OnlySystemExtensionId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoOnlyWalletOwner represents a OnlyWalletOwner error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoOnlyWalletOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyWalletOwner()
func TeePaymentsUtxoOnlyWalletOwnerErrorID() common.Hash {
	return common.HexToHash("0x3132d62dfddd1d8b6a9dd805001453bbb6a84b746ebe51fd650df1a6704a8143")
}

// UnpackOnlyWalletOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyWalletOwner()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackOnlyWalletOwnerError(raw []byte) (*TeePaymentsUtxoOnlyWalletOwner, error) {
	out := new(TeePaymentsUtxoOnlyWalletOwner)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "OnlyWalletOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoPMWMultisigAccountAddressAlreadySet represents a PMWMultisigAccountAddressAlreadySet error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoPMWMultisigAccountAddressAlreadySet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PMWMultisigAccountAddressAlreadySet()
func TeePaymentsUtxoPMWMultisigAccountAddressAlreadySetErrorID() common.Hash {
	return common.HexToHash("0xe3e804ecd31f69f9fde81b0b1994f9d3057e672dfe1cb39f32b68f667aa752e6")
}

// UnpackPMWMultisigAccountAddressAlreadySetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PMWMultisigAccountAddressAlreadySet()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackPMWMultisigAccountAddressAlreadySetError(raw []byte) (*TeePaymentsUtxoPMWMultisigAccountAddressAlreadySet, error) {
	out := new(TeePaymentsUtxoPMWMultisigAccountAddressAlreadySet)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "PMWMultisigAccountAddressAlreadySet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoPMWMultisigAccountNotRegistered represents a PMWMultisigAccountNotRegistered error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoPMWMultisigAccountNotRegistered struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PMWMultisigAccountNotRegistered()
func TeePaymentsUtxoPMWMultisigAccountNotRegisteredErrorID() common.Hash {
	return common.HexToHash("0x9d78814a2f7e263f4d4cc3000f076b4c90004105130f1880c0dc94acd2464be9")
}

// UnpackPMWMultisigAccountNotRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PMWMultisigAccountNotRegistered()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackPMWMultisigAccountNotRegisteredError(raw []byte) (*TeePaymentsUtxoPMWMultisigAccountNotRegistered, error) {
	out := new(TeePaymentsUtxoPMWMultisigAccountNotRegistered)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "PMWMultisigAccountNotRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoPaymentAmountZero represents a PaymentAmountZero error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoPaymentAmountZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentAmountZero()
func TeePaymentsUtxoPaymentAmountZeroErrorID() common.Hash {
	return common.HexToHash("0xa74a1c28a64fc70e7d718ae1b3af31f436e1c55b3754d607ef7ec829ec74034c")
}

// UnpackPaymentAmountZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentAmountZero()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackPaymentAmountZeroError(raw []byte) (*TeePaymentsUtxoPaymentAmountZero, error) {
	out := new(TeePaymentsUtxoPaymentAmountZero)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "PaymentAmountZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoPaymentHashMismatch represents a PaymentHashMismatch error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoPaymentHashMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentHashMismatch()
func TeePaymentsUtxoPaymentHashMismatchErrorID() common.Hash {
	return common.HexToHash("0x85628ea5f0ea30f18d9205be9d6ae9bd4b189b3dc00aeba346ab8f4c0e25674e")
}

// UnpackPaymentHashMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentHashMismatch()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackPaymentHashMismatchError(raw []byte) (*TeePaymentsUtxoPaymentHashMismatch, error) {
	out := new(TeePaymentsUtxoPaymentHashMismatch)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "PaymentHashMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoPaymentNotInBatch represents a PaymentNotInBatch error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoPaymentNotInBatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentNotInBatch()
func TeePaymentsUtxoPaymentNotInBatchErrorID() common.Hash {
	return common.HexToHash("0x0514c319b299aa4dd34cdc47de28afafdc2bf6184d6a274e846659612969994c")
}

// UnpackPaymentNotInBatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentNotInBatch()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackPaymentNotInBatchError(raw []byte) (*TeePaymentsUtxoPaymentNotInBatch, error) {
	out := new(TeePaymentsUtxoPaymentNotInBatch)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "PaymentNotInBatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoReissueRewardEpochChanged represents a ReissueRewardEpochChanged error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoReissueRewardEpochChanged struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReissueRewardEpochChanged()
func TeePaymentsUtxoReissueRewardEpochChangedErrorID() common.Hash {
	return common.HexToHash("0xf13e54f4f286092adf499378106be96f2b406030e50f5134805d6aee90dac4c7")
}

// UnpackReissueRewardEpochChangedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReissueRewardEpochChanged()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackReissueRewardEpochChangedError(raw []byte) (*TeePaymentsUtxoReissueRewardEpochChanged, error) {
	out := new(TeePaymentsUtxoReissueRewardEpochChanged)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "ReissueRewardEpochChanged", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoReplacementAlreadyFinalized represents a ReplacementAlreadyFinalized error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoReplacementAlreadyFinalized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReplacementAlreadyFinalized()
func TeePaymentsUtxoReplacementAlreadyFinalizedErrorID() common.Hash {
	return common.HexToHash("0x9a9fc08856f923d05d201c6644b1558abb8e68115e5da93d267823c918532338")
}

// UnpackReplacementAlreadyFinalizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReplacementAlreadyFinalized()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackReplacementAlreadyFinalizedError(raw []byte) (*TeePaymentsUtxoReplacementAlreadyFinalized, error) {
	out := new(TeePaymentsUtxoReplacementAlreadyFinalized)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "ReplacementAlreadyFinalized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoScheduledSignaturesUnsupported represents a ScheduledSignaturesUnsupported error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoScheduledSignaturesUnsupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ScheduledSignaturesUnsupported()
func TeePaymentsUtxoScheduledSignaturesUnsupportedErrorID() common.Hash {
	return common.HexToHash("0x214a99dbcb28037b83881b0f6833643a04410fc2678ee379bb16afec2060772e")
}

// UnpackScheduledSignaturesUnsupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ScheduledSignaturesUnsupported()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackScheduledSignaturesUnsupportedError(raw []byte) (*TeePaymentsUtxoScheduledSignaturesUnsupported, error) {
	out := new(TeePaymentsUtxoScheduledSignaturesUnsupported)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "ScheduledSignaturesUnsupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoTimelockCallNotFound represents a TimelockCallNotFound error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoTimelockCallNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockCallNotFound()
func TeePaymentsUtxoTimelockCallNotFoundErrorID() common.Hash {
	return common.HexToHash("0x262ac987dd09120380d393a0c52cdef1ca58949d1107dc26d8da9f7b527eb33d")
}

// UnpackTimelockCallNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockCallNotFound()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackTimelockCallNotFoundError(raw []byte) (*TeePaymentsUtxoTimelockCallNotFound, error) {
	out := new(TeePaymentsUtxoTimelockCallNotFound)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "TimelockCallNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoTimelockNotAllowedYet represents a TimelockNotAllowedYet error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoTimelockNotAllowedYet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockNotAllowedYet()
func TeePaymentsUtxoTimelockNotAllowedYetErrorID() common.Hash {
	return common.HexToHash("0x6124e5c2c545bec1d02048c68500eebdabeae0cc69995a629ebccf8a62f69083")
}

// UnpackTimelockNotAllowedYetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockNotAllowedYet()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackTimelockNotAllowedYetError(raw []byte) (*TeePaymentsUtxoTimelockNotAllowedYet, error) {
	out := new(TeePaymentsUtxoTimelockNotAllowedYet)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "TimelockNotAllowedYet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func TeePaymentsUtxoUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*TeePaymentsUtxoUUPSUnauthorizedCallContext, error) {
	out := new(TeePaymentsUtxoUUPSUnauthorizedCallContext)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func TeePaymentsUtxoUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*TeePaymentsUtxoUUPSUnsupportedProxiableUUID, error) {
	out := new(TeePaymentsUtxoUUPSUnsupportedProxiableUUID)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoUnsupportedSourceId represents a UnsupportedSourceId error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoUnsupportedSourceId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedSourceId()
func TeePaymentsUtxoUnsupportedSourceIdErrorID() common.Hash {
	return common.HexToHash("0xb9c2749e785a86b27011ac62f2eeace91e285734582a3d21db2e093d45e3124a")
}

// UnpackUnsupportedSourceIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedSourceId()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackUnsupportedSourceIdError(raw []byte) (*TeePaymentsUtxoUnsupportedSourceId, error) {
	out := new(TeePaymentsUtxoUnsupportedSourceId)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "UnsupportedSourceId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoWalletNotInProduction represents a WalletNotInProduction error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoWalletNotInProduction struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WalletNotInProduction()
func TeePaymentsUtxoWalletNotInProductionErrorID() common.Hash {
	return common.HexToHash("0x2b7f97d6ba9d96a13708bf15f7c7ed50584fa90c4e3f979e82b455f65581f045")
}

// UnpackWalletNotInProductionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WalletNotInProduction()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackWalletNotInProductionError(raw []byte) (*TeePaymentsUtxoWalletNotInProduction, error) {
	out := new(TeePaymentsUtxoWalletNotInProduction)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "WalletNotInProduction", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUtxoWrongKeyType represents a WrongKeyType error raised by the TeePaymentsUtxo contract.
type TeePaymentsUtxoWrongKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WrongKeyType()
func TeePaymentsUtxoWrongKeyTypeErrorID() common.Hash {
	return common.HexToHash("0xfb224bce70b4062685f138c7eed15a72b903495475f7cd7843835692a7c58bbd")
}

// UnpackWrongKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WrongKeyType()
func (teePaymentsUtxo *TeePaymentsUtxo) UnpackWrongKeyTypeError(raw []byte) (*TeePaymentsUtxoWrongKeyType, error) {
	out := new(TeePaymentsUtxoWrongKeyType)
	if err := teePaymentsUtxo.abi.UnpackIntoInterface(out, "WrongKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}
