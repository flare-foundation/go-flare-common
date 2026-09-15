// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fdc2

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

// IBtcAccountConfiguredAnchor is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountConfiguredAnchor struct {
	GenesisAnchorTxid [32]byte
	GenesisAnchorVout uint32
}

// IBtcAccountConfiguredProof is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountConfiguredProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  IBtcAccountConfiguredRequestBody
	ResponseBody IBtcAccountConfiguredResponseBody
}

// IBtcAccountConfiguredRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountConfiguredRequestBody struct {
	AccountIndex uint32
	PublicKeys   [][]byte
	Threshold    uint64
	Anchors      []IBtcAccountConfiguredAnchor
}

// IBtcAccountConfiguredResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountConfiguredResponseBody struct {
	Status         uint8
	AccountAddress string
}

// IBtcDepositProof is an auto generated low-level Go binding around an user-defined struct.
type IBtcDepositProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  IBtcDepositRequestBody
	ResponseBody IBtcDepositResponseBody
}

// IBtcDepositRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IBtcDepositRequestBody struct {
	TransactionId        [32]byte
	OutputIndex          uint16
	ReceivingAddressHash [32]byte
	MinConfirmations     uint16
	InputIndex           uint16
}

// IBtcDepositResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IBtcDepositResponseBody struct {
	Amount            *big.Int
	Memo              []byte
	BlockNumber       uint64
	BlockTimestamp    uint64
	Confirmations     uint64
	InputAddress      string
	InputAddressHash  [32]byte
	InputSighashTypes []byte
}

// IBtcWalletAddressProof is an auto generated low-level Go binding around an user-defined struct.
type IBtcWalletAddressProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  IBtcWalletAddressRequestBody
	ResponseBody IBtcWalletAddressResponseBody
}

// IBtcWalletAddressRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IBtcWalletAddressRequestBody struct {
	WalletRegistry  common.Address
	WalletId        [32]byte
	AccountIndex    uint32
	DerivationIndex uint32
}

// IBtcWalletAddressResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IBtcWalletAddressResponseBody struct {
	ReceivingAddress string
}

// ICspProposalCheckProof is an auto generated low-level Go binding around an user-defined struct.
type ICspProposalCheckProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  ICspProposalCheckRequestBody
	ResponseBody ICspProposalCheckResponseBody
}

// ICspProposalCheckRequestBody is an auto generated low-level Go binding around an user-defined struct.
type ICspProposalCheckRequestBody struct {
	WalletRegistry     common.Address
	WalletId           [32]byte
	AccountIndex       uint32
	SequencePosition   uint64
	Attempt            uint32
	EligibleGeneration uint64
	PackageHash        [32]byte
}

// ICspProposalCheckResponseBody is an auto generated low-level Go binding around an user-defined struct.
type ICspProposalCheckResponseBody struct {
	ProposerAddress     common.Address
	Score               uint64
	PaymentCount        uint32
	ChainCommitmentHash [32]byte
}

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

// INativeNonceAccountConfiguredProof is an auto generated low-level Go binding around an user-defined struct.
type INativeNonceAccountConfiguredProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  INativeNonceAccountConfiguredRequestBody
	ResponseBody INativeNonceAccountConfiguredResponseBody
}

// INativeNonceAccountConfiguredRequestBody is an auto generated low-level Go binding around an user-defined struct.
type INativeNonceAccountConfiguredRequestBody struct {
	AccountAddress string
	PublicKeys     [][]byte
	Threshold      uint64
}

// INativeNonceAccountConfiguredResponseBody is an auto generated low-level Go binding around an user-defined struct.
type INativeNonceAccountConfiguredResponseBody struct {
	Status   uint8
	Sequence uint64
}

// ITeeAvailabilityCheckProof is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  ITeeAvailabilityCheckRequestBody
	ResponseBody ITeeAvailabilityCheckResponseBody
}

// ITeeAvailabilityCheckRequestBody is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckRequestBody struct {
	TeeId         common.Address
	TeeProxyId    common.Address
	Url           string
	Challenge     [32]byte
	InstructionId [32]byte
}

// ITeeAvailabilityCheckResponseBody is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckResponseBody struct {
	Status                 uint8
	TeeTimestamp           uint64
	CodeHash               [32]byte
	Platform               [32]byte
	InitialSigningPolicyId uint32
	LastSigningPolicyId    uint32
	State                  ITeeAvailabilityCheckTeeState
}

// ITeeAvailabilityCheckTeeState is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckTeeState struct {
	SystemState        []byte
	SystemStateVersion [32]byte
	State              []byte
	StateVersion       [32]byte
}

// IWalletFeeProofProof is an auto generated low-level Go binding around an user-defined struct.
type IWalletFeeProofProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  IWalletFeeProofRequestBody
	ResponseBody IWalletFeeProofResponseBody
}

// IWalletFeeProofRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IWalletFeeProofRequestBody struct {
	OpType         [32]byte
	SenderAddress  string
	FirstPaymentId uint64
	BatchCount     uint64
	UntilTimestamp uint64
}

// IWalletFeeProofResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IWalletFeeProofResponseBody struct {
	LastPaymentId uint64
	ActualFee     *big.Int
	EstimatedFee  *big.Int
}

// IWalletPaymentStatusProof is an auto generated low-level Go binding around an user-defined struct.
type IWalletPaymentStatusProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  IWalletPaymentStatusRequestBody
	ResponseBody IWalletPaymentStatusResponseBody
}

// IWalletPaymentStatusRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IWalletPaymentStatusRequestBody struct {
	OpType        [32]byte
	SenderAddress string
	PaymentId     uint64
	TransactionId [32]byte
}

// IWalletPaymentStatusResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IWalletPaymentStatusResponseBody struct {
	RecipientAddress  string
	TokenId           []byte
	Amount            *big.Int
	MaxFee            *big.Int
	PaymentReference  [32]byte
	TransactionStatus uint8
	RevertReason      string
	ReceivedAmount    *big.Int
	TransactionFee    *big.Int
	TransactionId     [32]byte
	BlockNumber       uint64
	BlockTimestamp    uint64
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// Fdc2MetaData contains all meta data concerning the Fdc2 contract.
var Fdc2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckTeeStateStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcAccountConfigured.Anchor\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"btcAccountConfiguredAnchorStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcAccountConfigured.Anchor[]\",\"name\":\"anchors\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBtcAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIBtcAccountConfigured.BtcAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structIBtcAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIBtcAccountConfigured.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"btcAccountConfiguredProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcAccountConfigured.Anchor[]\",\"name\":\"anchors\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBtcAccountConfigured.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"btcAccountConfiguredRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIBtcAccountConfigured.BtcAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structIBtcAccountConfigured.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"btcAccountConfiguredResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"}],\"internalType\":\"structIFdc2Hub.Fdc2RequestHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"requestBody\",\"type\":\"bytes\"}],\"internalType\":\"structIFdc2Hub.Fdc2AttestationRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"fdc2AttestationRequestStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"}],\"internalType\":\"structIFdc2Hub.Fdc2RequestHeader\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"fdc2RequestHeaderStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"fdc2ResponseHeaderStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"internalType\":\"structINativeNonceAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumINativeNonceAccountConfigured.NativeNonceAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structINativeNonceAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structINativeNonceAccountConfigured.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"nativeNonceAccountConfiguredProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"internalType\":\"structINativeNonceAccountConfigured.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"nativeNonceAccountConfiguredRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumINativeNonceAccountConfigured.NativeNonceAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structINativeNonceAccountConfigured.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"nativeNonceAccountConfiguredResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"firstPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"untilTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIWalletFeeProof.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"lastPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimatedFee\",\"type\":\"uint256\"}],\"internalType\":\"structIWalletFeeProof.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIWalletFeeProof.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"walletFeeProofProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"firstPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"untilTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIWalletFeeProof.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"walletFeeProofRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"lastPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimatedFee\",\"type\":\"uint256\"}],\"internalType\":\"structIWalletFeeProof.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"walletFeeProofResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"internalType\":\"structIWalletPaymentStatus.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"transactionStatus\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transactionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIWalletPaymentStatus.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIWalletPaymentStatus.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"walletPaymentStatusProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"internalType\":\"structIWalletPaymentStatus.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"walletPaymentStatusRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"transactionStatus\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transactionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIWalletPaymentStatus.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"walletPaymentStatusResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"name\":\"walletRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"walletId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accountIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sequencePosition\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"attempt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"eligibleGeneration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"packageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"internalType\":\"structICspProposalCheck.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"cspProposalCheckRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"name\":\"proposerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"score\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"paymentCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"internalType\":\"structICspProposalCheck.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"cspProposalCheckResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"walletRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"walletId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accountIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sequencePosition\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"attempt\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"eligibleGeneration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"packageHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"internalType\":\"structICspProposalCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"proposerAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"score\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"paymentCount\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"internalType\":\"structICspProposalCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structICspProposalCheck.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"cspProposalCheckProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"name\":\"transactionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"outputIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"receivingAddressHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"minConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"inputIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"internalType\":\"structIBtcDeposit.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"btcDepositRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"memo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"inputAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"inputAddressHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"inputSighashTypes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"internalType\":\"structIBtcDeposit.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"btcDepositResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"transactionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"outputIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"receivingAddressHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"minConfirmations\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"inputIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"internalType\":\"structIBtcDeposit.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"memo\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"confirmations\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"inputAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"inputAddressHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"inputSighashTypes\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"internalType\":\"structIBtcDeposit.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIBtcDeposit.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"btcDepositProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"name\":\"walletRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"walletId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accountIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"derivationIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"internalType\":\"structIBtcWalletAddress.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"btcWalletAddressRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"name\":\"receivingAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"internalType\":\"structIBtcWalletAddress.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"btcWalletAddressResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"walletRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"walletId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"accountIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"derivationIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"internalType\":\"structIBtcWalletAddress.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"receivingAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"internalType\":\"structIBtcWalletAddress.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIBtcWalletAddress.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"btcWalletAddressProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Fdc2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Fdc2MetaData.ABI instead.
var Fdc2ABI = Fdc2MetaData.ABI

// Fdc2 is an auto generated Go binding around an Ethereum contract.
type Fdc2 struct {
	Fdc2Caller     // Read-only binding to the contract
	Fdc2Transactor // Write-only binding to the contract
	Fdc2Filterer   // Log filterer for contract events
}

// Fdc2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Fdc2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fdc2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Fdc2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fdc2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Fdc2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fdc2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Fdc2Session struct {
	Contract     *Fdc2             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Fdc2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Fdc2CallerSession struct {
	Contract *Fdc2Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Fdc2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Fdc2TransactorSession struct {
	Contract     *Fdc2Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Fdc2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Fdc2Raw struct {
	Contract *Fdc2 // Generic contract binding to access the raw methods on
}

// Fdc2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Fdc2CallerRaw struct {
	Contract *Fdc2Caller // Generic read-only contract binding to access the raw methods on
}

// Fdc2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Fdc2TransactorRaw struct {
	Contract *Fdc2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFdc2 creates a new instance of Fdc2, bound to a specific deployed contract.
func NewFdc2(address common.Address, backend bind.ContractBackend) (*Fdc2, error) {
	contract, err := bindFdc2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fdc2{Fdc2Caller: Fdc2Caller{contract: contract}, Fdc2Transactor: Fdc2Transactor{contract: contract}, Fdc2Filterer: Fdc2Filterer{contract: contract}}, nil
}

// NewFdc2Caller creates a new read-only instance of Fdc2, bound to a specific deployed contract.
func NewFdc2Caller(address common.Address, caller bind.ContractCaller) (*Fdc2Caller, error) {
	contract, err := bindFdc2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Fdc2Caller{contract: contract}, nil
}

// NewFdc2Transactor creates a new write-only instance of Fdc2, bound to a specific deployed contract.
func NewFdc2Transactor(address common.Address, transactor bind.ContractTransactor) (*Fdc2Transactor, error) {
	contract, err := bindFdc2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Fdc2Transactor{contract: contract}, nil
}

// NewFdc2Filterer creates a new log filterer instance of Fdc2, bound to a specific deployed contract.
func NewFdc2Filterer(address common.Address, filterer bind.ContractFilterer) (*Fdc2Filterer, error) {
	contract, err := bindFdc2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Fdc2Filterer{contract: contract}, nil
}

// bindFdc2 binds a generic wrapper to an already deployed contract.
func bindFdc2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Fdc2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fdc2 *Fdc2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fdc2.Contract.Fdc2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fdc2 *Fdc2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fdc2.Contract.Fdc2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fdc2 *Fdc2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fdc2.Contract.Fdc2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fdc2 *Fdc2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fdc2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fdc2 *Fdc2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fdc2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fdc2 *Fdc2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fdc2.Contract.contract.Transact(opts, method, params...)
}

// AvailabilityCheckProofStruct is a paid mutator transaction binding the contract method 0x6d603180.
//
// Solidity: function availabilityCheckProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) ) returns()
func (_Fdc2 *Fdc2Transactor) AvailabilityCheckProofStruct(opts *bind.TransactOpts, arg0 ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "availabilityCheckProofStruct", arg0)
}

// AvailabilityCheckProofStruct is a paid mutator transaction binding the contract method 0x6d603180.
//
// Solidity: function availabilityCheckProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) ) returns()
func (_Fdc2 *Fdc2Session) AvailabilityCheckProofStruct(arg0 ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _Fdc2.Contract.AvailabilityCheckProofStruct(&_Fdc2.TransactOpts, arg0)
}

// AvailabilityCheckProofStruct is a paid mutator transaction binding the contract method 0x6d603180.
//
// Solidity: function availabilityCheckProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) ) returns()
func (_Fdc2 *Fdc2TransactorSession) AvailabilityCheckProofStruct(arg0 ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _Fdc2.Contract.AvailabilityCheckProofStruct(&_Fdc2.TransactOpts, arg0)
}

// AvailabilityCheckRequestBodyStruct is a paid mutator transaction binding the contract method 0x7594627a.
//
// Solidity: function availabilityCheckRequestBodyStruct((address,address,string,bytes32,bytes32) ) returns()
func (_Fdc2 *Fdc2Transactor) AvailabilityCheckRequestBodyStruct(opts *bind.TransactOpts, arg0 ITeeAvailabilityCheckRequestBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "availabilityCheckRequestBodyStruct", arg0)
}

// AvailabilityCheckRequestBodyStruct is a paid mutator transaction binding the contract method 0x7594627a.
//
// Solidity: function availabilityCheckRequestBodyStruct((address,address,string,bytes32,bytes32) ) returns()
func (_Fdc2 *Fdc2Session) AvailabilityCheckRequestBodyStruct(arg0 ITeeAvailabilityCheckRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.AvailabilityCheckRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// AvailabilityCheckRequestBodyStruct is a paid mutator transaction binding the contract method 0x7594627a.
//
// Solidity: function availabilityCheckRequestBodyStruct((address,address,string,bytes32,bytes32) ) returns()
func (_Fdc2 *Fdc2TransactorSession) AvailabilityCheckRequestBodyStruct(arg0 ITeeAvailabilityCheckRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.AvailabilityCheckRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// AvailabilityCheckResponseBodyStruct is a paid mutator transaction binding the contract method 0x972ce998.
//
// Solidity: function availabilityCheckResponseBodyStruct((uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32)) ) returns()
func (_Fdc2 *Fdc2Transactor) AvailabilityCheckResponseBodyStruct(opts *bind.TransactOpts, arg0 ITeeAvailabilityCheckResponseBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "availabilityCheckResponseBodyStruct", arg0)
}

// AvailabilityCheckResponseBodyStruct is a paid mutator transaction binding the contract method 0x972ce998.
//
// Solidity: function availabilityCheckResponseBodyStruct((uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32)) ) returns()
func (_Fdc2 *Fdc2Session) AvailabilityCheckResponseBodyStruct(arg0 ITeeAvailabilityCheckResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.AvailabilityCheckResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// AvailabilityCheckResponseBodyStruct is a paid mutator transaction binding the contract method 0x972ce998.
//
// Solidity: function availabilityCheckResponseBodyStruct((uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32)) ) returns()
func (_Fdc2 *Fdc2TransactorSession) AvailabilityCheckResponseBodyStruct(arg0 ITeeAvailabilityCheckResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.AvailabilityCheckResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// AvailabilityCheckTeeStateStruct is a paid mutator transaction binding the contract method 0xf46db3cf.
//
// Solidity: function availabilityCheckTeeStateStruct((bytes,bytes32,bytes,bytes32) ) returns()
func (_Fdc2 *Fdc2Transactor) AvailabilityCheckTeeStateStruct(opts *bind.TransactOpts, arg0 ITeeAvailabilityCheckTeeState) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "availabilityCheckTeeStateStruct", arg0)
}

// AvailabilityCheckTeeStateStruct is a paid mutator transaction binding the contract method 0xf46db3cf.
//
// Solidity: function availabilityCheckTeeStateStruct((bytes,bytes32,bytes,bytes32) ) returns()
func (_Fdc2 *Fdc2Session) AvailabilityCheckTeeStateStruct(arg0 ITeeAvailabilityCheckTeeState) (*types.Transaction, error) {
	return _Fdc2.Contract.AvailabilityCheckTeeStateStruct(&_Fdc2.TransactOpts, arg0)
}

// AvailabilityCheckTeeStateStruct is a paid mutator transaction binding the contract method 0xf46db3cf.
//
// Solidity: function availabilityCheckTeeStateStruct((bytes,bytes32,bytes,bytes32) ) returns()
func (_Fdc2 *Fdc2TransactorSession) AvailabilityCheckTeeStateStruct(arg0 ITeeAvailabilityCheckTeeState) (*types.Transaction, error) {
	return _Fdc2.Contract.AvailabilityCheckTeeStateStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcAccountConfiguredAnchorStruct is a paid mutator transaction binding the contract method 0xda332262.
//
// Solidity: function btcAccountConfiguredAnchorStruct((bytes32,uint32) ) returns()
func (_Fdc2 *Fdc2Transactor) BtcAccountConfiguredAnchorStruct(opts *bind.TransactOpts, arg0 IBtcAccountConfiguredAnchor) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "btcAccountConfiguredAnchorStruct", arg0)
}

// BtcAccountConfiguredAnchorStruct is a paid mutator transaction binding the contract method 0xda332262.
//
// Solidity: function btcAccountConfiguredAnchorStruct((bytes32,uint32) ) returns()
func (_Fdc2 *Fdc2Session) BtcAccountConfiguredAnchorStruct(arg0 IBtcAccountConfiguredAnchor) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcAccountConfiguredAnchorStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcAccountConfiguredAnchorStruct is a paid mutator transaction binding the contract method 0xda332262.
//
// Solidity: function btcAccountConfiguredAnchorStruct((bytes32,uint32) ) returns()
func (_Fdc2 *Fdc2TransactorSession) BtcAccountConfiguredAnchorStruct(arg0 IBtcAccountConfiguredAnchor) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcAccountConfiguredAnchorStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcAccountConfiguredProofStruct is a paid mutator transaction binding the contract method 0xea3e8eb5.
//
// Solidity: function btcAccountConfiguredProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) ) returns()
func (_Fdc2 *Fdc2Transactor) BtcAccountConfiguredProofStruct(opts *bind.TransactOpts, arg0 IBtcAccountConfiguredProof) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "btcAccountConfiguredProofStruct", arg0)
}

// BtcAccountConfiguredProofStruct is a paid mutator transaction binding the contract method 0xea3e8eb5.
//
// Solidity: function btcAccountConfiguredProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) ) returns()
func (_Fdc2 *Fdc2Session) BtcAccountConfiguredProofStruct(arg0 IBtcAccountConfiguredProof) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcAccountConfiguredProofStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcAccountConfiguredProofStruct is a paid mutator transaction binding the contract method 0xea3e8eb5.
//
// Solidity: function btcAccountConfiguredProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) ) returns()
func (_Fdc2 *Fdc2TransactorSession) BtcAccountConfiguredProofStruct(arg0 IBtcAccountConfiguredProof) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcAccountConfiguredProofStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcAccountConfiguredRequestBodyStruct is a paid mutator transaction binding the contract method 0x391598c2.
//
// Solidity: function btcAccountConfiguredRequestBodyStruct((uint32,bytes[],uint64,(bytes32,uint32)[]) ) returns()
func (_Fdc2 *Fdc2Transactor) BtcAccountConfiguredRequestBodyStruct(opts *bind.TransactOpts, arg0 IBtcAccountConfiguredRequestBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "btcAccountConfiguredRequestBodyStruct", arg0)
}

// BtcAccountConfiguredRequestBodyStruct is a paid mutator transaction binding the contract method 0x391598c2.
//
// Solidity: function btcAccountConfiguredRequestBodyStruct((uint32,bytes[],uint64,(bytes32,uint32)[]) ) returns()
func (_Fdc2 *Fdc2Session) BtcAccountConfiguredRequestBodyStruct(arg0 IBtcAccountConfiguredRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcAccountConfiguredRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcAccountConfiguredRequestBodyStruct is a paid mutator transaction binding the contract method 0x391598c2.
//
// Solidity: function btcAccountConfiguredRequestBodyStruct((uint32,bytes[],uint64,(bytes32,uint32)[]) ) returns()
func (_Fdc2 *Fdc2TransactorSession) BtcAccountConfiguredRequestBodyStruct(arg0 IBtcAccountConfiguredRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcAccountConfiguredRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcAccountConfiguredResponseBodyStruct is a paid mutator transaction binding the contract method 0x035333d4.
//
// Solidity: function btcAccountConfiguredResponseBodyStruct((uint8,string) ) returns()
func (_Fdc2 *Fdc2Transactor) BtcAccountConfiguredResponseBodyStruct(opts *bind.TransactOpts, arg0 IBtcAccountConfiguredResponseBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "btcAccountConfiguredResponseBodyStruct", arg0)
}

// BtcAccountConfiguredResponseBodyStruct is a paid mutator transaction binding the contract method 0x035333d4.
//
// Solidity: function btcAccountConfiguredResponseBodyStruct((uint8,string) ) returns()
func (_Fdc2 *Fdc2Session) BtcAccountConfiguredResponseBodyStruct(arg0 IBtcAccountConfiguredResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcAccountConfiguredResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcAccountConfiguredResponseBodyStruct is a paid mutator transaction binding the contract method 0x035333d4.
//
// Solidity: function btcAccountConfiguredResponseBodyStruct((uint8,string) ) returns()
func (_Fdc2 *Fdc2TransactorSession) BtcAccountConfiguredResponseBodyStruct(arg0 IBtcAccountConfiguredResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcAccountConfiguredResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcDepositProofStruct is a paid mutator transaction binding the contract method 0x10006130.
//
// Solidity: function btcDepositProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,uint16,bytes32,uint16,uint16),(uint256,bytes,uint64,uint64,uint64,string,bytes32,bytes)) ) returns()
func (_Fdc2 *Fdc2Transactor) BtcDepositProofStruct(opts *bind.TransactOpts, arg0 IBtcDepositProof) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "btcDepositProofStruct", arg0)
}

// BtcDepositProofStruct is a paid mutator transaction binding the contract method 0x10006130.
//
// Solidity: function btcDepositProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,uint16,bytes32,uint16,uint16),(uint256,bytes,uint64,uint64,uint64,string,bytes32,bytes)) ) returns()
func (_Fdc2 *Fdc2Session) BtcDepositProofStruct(arg0 IBtcDepositProof) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcDepositProofStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcDepositProofStruct is a paid mutator transaction binding the contract method 0x10006130.
//
// Solidity: function btcDepositProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,uint16,bytes32,uint16,uint16),(uint256,bytes,uint64,uint64,uint64,string,bytes32,bytes)) ) returns()
func (_Fdc2 *Fdc2TransactorSession) BtcDepositProofStruct(arg0 IBtcDepositProof) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcDepositProofStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcDepositRequestBodyStruct is a paid mutator transaction binding the contract method 0xef4860fb.
//
// Solidity: function btcDepositRequestBodyStruct((bytes32,uint16,bytes32,uint16,uint16) ) returns()
func (_Fdc2 *Fdc2Transactor) BtcDepositRequestBodyStruct(opts *bind.TransactOpts, arg0 IBtcDepositRequestBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "btcDepositRequestBodyStruct", arg0)
}

// BtcDepositRequestBodyStruct is a paid mutator transaction binding the contract method 0xef4860fb.
//
// Solidity: function btcDepositRequestBodyStruct((bytes32,uint16,bytes32,uint16,uint16) ) returns()
func (_Fdc2 *Fdc2Session) BtcDepositRequestBodyStruct(arg0 IBtcDepositRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcDepositRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcDepositRequestBodyStruct is a paid mutator transaction binding the contract method 0xef4860fb.
//
// Solidity: function btcDepositRequestBodyStruct((bytes32,uint16,bytes32,uint16,uint16) ) returns()
func (_Fdc2 *Fdc2TransactorSession) BtcDepositRequestBodyStruct(arg0 IBtcDepositRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcDepositRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcDepositResponseBodyStruct is a paid mutator transaction binding the contract method 0x89246ed6.
//
// Solidity: function btcDepositResponseBodyStruct((uint256,bytes,uint64,uint64,uint64,string,bytes32,bytes) ) returns()
func (_Fdc2 *Fdc2Transactor) BtcDepositResponseBodyStruct(opts *bind.TransactOpts, arg0 IBtcDepositResponseBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "btcDepositResponseBodyStruct", arg0)
}

// BtcDepositResponseBodyStruct is a paid mutator transaction binding the contract method 0x89246ed6.
//
// Solidity: function btcDepositResponseBodyStruct((uint256,bytes,uint64,uint64,uint64,string,bytes32,bytes) ) returns()
func (_Fdc2 *Fdc2Session) BtcDepositResponseBodyStruct(arg0 IBtcDepositResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcDepositResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcDepositResponseBodyStruct is a paid mutator transaction binding the contract method 0x89246ed6.
//
// Solidity: function btcDepositResponseBodyStruct((uint256,bytes,uint64,uint64,uint64,string,bytes32,bytes) ) returns()
func (_Fdc2 *Fdc2TransactorSession) BtcDepositResponseBodyStruct(arg0 IBtcDepositResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcDepositResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcWalletAddressProofStruct is a paid mutator transaction binding the contract method 0xd124a2e7.
//
// Solidity: function btcWalletAddressProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,bytes32,uint32,uint32),(string)) ) returns()
func (_Fdc2 *Fdc2Transactor) BtcWalletAddressProofStruct(opts *bind.TransactOpts, arg0 IBtcWalletAddressProof) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "btcWalletAddressProofStruct", arg0)
}

// BtcWalletAddressProofStruct is a paid mutator transaction binding the contract method 0xd124a2e7.
//
// Solidity: function btcWalletAddressProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,bytes32,uint32,uint32),(string)) ) returns()
func (_Fdc2 *Fdc2Session) BtcWalletAddressProofStruct(arg0 IBtcWalletAddressProof) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcWalletAddressProofStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcWalletAddressProofStruct is a paid mutator transaction binding the contract method 0xd124a2e7.
//
// Solidity: function btcWalletAddressProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,bytes32,uint32,uint32),(string)) ) returns()
func (_Fdc2 *Fdc2TransactorSession) BtcWalletAddressProofStruct(arg0 IBtcWalletAddressProof) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcWalletAddressProofStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcWalletAddressRequestBodyStruct is a paid mutator transaction binding the contract method 0xefc943ea.
//
// Solidity: function btcWalletAddressRequestBodyStruct((address,bytes32,uint32,uint32) ) returns()
func (_Fdc2 *Fdc2Transactor) BtcWalletAddressRequestBodyStruct(opts *bind.TransactOpts, arg0 IBtcWalletAddressRequestBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "btcWalletAddressRequestBodyStruct", arg0)
}

// BtcWalletAddressRequestBodyStruct is a paid mutator transaction binding the contract method 0xefc943ea.
//
// Solidity: function btcWalletAddressRequestBodyStruct((address,bytes32,uint32,uint32) ) returns()
func (_Fdc2 *Fdc2Session) BtcWalletAddressRequestBodyStruct(arg0 IBtcWalletAddressRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcWalletAddressRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcWalletAddressRequestBodyStruct is a paid mutator transaction binding the contract method 0xefc943ea.
//
// Solidity: function btcWalletAddressRequestBodyStruct((address,bytes32,uint32,uint32) ) returns()
func (_Fdc2 *Fdc2TransactorSession) BtcWalletAddressRequestBodyStruct(arg0 IBtcWalletAddressRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcWalletAddressRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcWalletAddressResponseBodyStruct is a paid mutator transaction binding the contract method 0x11565336.
//
// Solidity: function btcWalletAddressResponseBodyStruct((string) ) returns()
func (_Fdc2 *Fdc2Transactor) BtcWalletAddressResponseBodyStruct(opts *bind.TransactOpts, arg0 IBtcWalletAddressResponseBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "btcWalletAddressResponseBodyStruct", arg0)
}

// BtcWalletAddressResponseBodyStruct is a paid mutator transaction binding the contract method 0x11565336.
//
// Solidity: function btcWalletAddressResponseBodyStruct((string) ) returns()
func (_Fdc2 *Fdc2Session) BtcWalletAddressResponseBodyStruct(arg0 IBtcWalletAddressResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcWalletAddressResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// BtcWalletAddressResponseBodyStruct is a paid mutator transaction binding the contract method 0x11565336.
//
// Solidity: function btcWalletAddressResponseBodyStruct((string) ) returns()
func (_Fdc2 *Fdc2TransactorSession) BtcWalletAddressResponseBodyStruct(arg0 IBtcWalletAddressResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.BtcWalletAddressResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// CspProposalCheckProofStruct is a paid mutator transaction binding the contract method 0x2e272ca1.
//
// Solidity: function cspProposalCheckProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,bytes32,uint32,uint64,uint32,uint64,bytes32),(address,uint64,uint32,bytes32)) ) returns()
func (_Fdc2 *Fdc2Transactor) CspProposalCheckProofStruct(opts *bind.TransactOpts, arg0 ICspProposalCheckProof) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "cspProposalCheckProofStruct", arg0)
}

// CspProposalCheckProofStruct is a paid mutator transaction binding the contract method 0x2e272ca1.
//
// Solidity: function cspProposalCheckProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,bytes32,uint32,uint64,uint32,uint64,bytes32),(address,uint64,uint32,bytes32)) ) returns()
func (_Fdc2 *Fdc2Session) CspProposalCheckProofStruct(arg0 ICspProposalCheckProof) (*types.Transaction, error) {
	return _Fdc2.Contract.CspProposalCheckProofStruct(&_Fdc2.TransactOpts, arg0)
}

// CspProposalCheckProofStruct is a paid mutator transaction binding the contract method 0x2e272ca1.
//
// Solidity: function cspProposalCheckProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,bytes32,uint32,uint64,uint32,uint64,bytes32),(address,uint64,uint32,bytes32)) ) returns()
func (_Fdc2 *Fdc2TransactorSession) CspProposalCheckProofStruct(arg0 ICspProposalCheckProof) (*types.Transaction, error) {
	return _Fdc2.Contract.CspProposalCheckProofStruct(&_Fdc2.TransactOpts, arg0)
}

// CspProposalCheckRequestBodyStruct is a paid mutator transaction binding the contract method 0xe8e3dc48.
//
// Solidity: function cspProposalCheckRequestBodyStruct((address,bytes32,uint32,uint64,uint32,uint64,bytes32) ) returns()
func (_Fdc2 *Fdc2Transactor) CspProposalCheckRequestBodyStruct(opts *bind.TransactOpts, arg0 ICspProposalCheckRequestBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "cspProposalCheckRequestBodyStruct", arg0)
}

// CspProposalCheckRequestBodyStruct is a paid mutator transaction binding the contract method 0xe8e3dc48.
//
// Solidity: function cspProposalCheckRequestBodyStruct((address,bytes32,uint32,uint64,uint32,uint64,bytes32) ) returns()
func (_Fdc2 *Fdc2Session) CspProposalCheckRequestBodyStruct(arg0 ICspProposalCheckRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.CspProposalCheckRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// CspProposalCheckRequestBodyStruct is a paid mutator transaction binding the contract method 0xe8e3dc48.
//
// Solidity: function cspProposalCheckRequestBodyStruct((address,bytes32,uint32,uint64,uint32,uint64,bytes32) ) returns()
func (_Fdc2 *Fdc2TransactorSession) CspProposalCheckRequestBodyStruct(arg0 ICspProposalCheckRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.CspProposalCheckRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// CspProposalCheckResponseBodyStruct is a paid mutator transaction binding the contract method 0x4e0df24b.
//
// Solidity: function cspProposalCheckResponseBodyStruct((address,uint64,uint32,bytes32) ) returns()
func (_Fdc2 *Fdc2Transactor) CspProposalCheckResponseBodyStruct(opts *bind.TransactOpts, arg0 ICspProposalCheckResponseBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "cspProposalCheckResponseBodyStruct", arg0)
}

// CspProposalCheckResponseBodyStruct is a paid mutator transaction binding the contract method 0x4e0df24b.
//
// Solidity: function cspProposalCheckResponseBodyStruct((address,uint64,uint32,bytes32) ) returns()
func (_Fdc2 *Fdc2Session) CspProposalCheckResponseBodyStruct(arg0 ICspProposalCheckResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.CspProposalCheckResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// CspProposalCheckResponseBodyStruct is a paid mutator transaction binding the contract method 0x4e0df24b.
//
// Solidity: function cspProposalCheckResponseBodyStruct((address,uint64,uint32,bytes32) ) returns()
func (_Fdc2 *Fdc2TransactorSession) CspProposalCheckResponseBodyStruct(arg0 ICspProposalCheckResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.CspProposalCheckResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// Fdc2AttestationRequestStruct is a paid mutator transaction binding the contract method 0xdfce7e8c.
//
// Solidity: function fdc2AttestationRequestStruct(((bytes32,bytes32,uint16,address),bytes) ) returns()
func (_Fdc2 *Fdc2Transactor) Fdc2AttestationRequestStruct(opts *bind.TransactOpts, arg0 IFdc2HubFdc2AttestationRequest) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "fdc2AttestationRequestStruct", arg0)
}

// Fdc2AttestationRequestStruct is a paid mutator transaction binding the contract method 0xdfce7e8c.
//
// Solidity: function fdc2AttestationRequestStruct(((bytes32,bytes32,uint16,address),bytes) ) returns()
func (_Fdc2 *Fdc2Session) Fdc2AttestationRequestStruct(arg0 IFdc2HubFdc2AttestationRequest) (*types.Transaction, error) {
	return _Fdc2.Contract.Fdc2AttestationRequestStruct(&_Fdc2.TransactOpts, arg0)
}

// Fdc2AttestationRequestStruct is a paid mutator transaction binding the contract method 0xdfce7e8c.
//
// Solidity: function fdc2AttestationRequestStruct(((bytes32,bytes32,uint16,address),bytes) ) returns()
func (_Fdc2 *Fdc2TransactorSession) Fdc2AttestationRequestStruct(arg0 IFdc2HubFdc2AttestationRequest) (*types.Transaction, error) {
	return _Fdc2.Contract.Fdc2AttestationRequestStruct(&_Fdc2.TransactOpts, arg0)
}

// Fdc2RequestHeaderStruct is a paid mutator transaction binding the contract method 0x88bb2414.
//
// Solidity: function fdc2RequestHeaderStruct((bytes32,bytes32,uint16,address) ) returns()
func (_Fdc2 *Fdc2Transactor) Fdc2RequestHeaderStruct(opts *bind.TransactOpts, arg0 IFdc2HubFdc2RequestHeader) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "fdc2RequestHeaderStruct", arg0)
}

// Fdc2RequestHeaderStruct is a paid mutator transaction binding the contract method 0x88bb2414.
//
// Solidity: function fdc2RequestHeaderStruct((bytes32,bytes32,uint16,address) ) returns()
func (_Fdc2 *Fdc2Session) Fdc2RequestHeaderStruct(arg0 IFdc2HubFdc2RequestHeader) (*types.Transaction, error) {
	return _Fdc2.Contract.Fdc2RequestHeaderStruct(&_Fdc2.TransactOpts, arg0)
}

// Fdc2RequestHeaderStruct is a paid mutator transaction binding the contract method 0x88bb2414.
//
// Solidity: function fdc2RequestHeaderStruct((bytes32,bytes32,uint16,address) ) returns()
func (_Fdc2 *Fdc2TransactorSession) Fdc2RequestHeaderStruct(arg0 IFdc2HubFdc2RequestHeader) (*types.Transaction, error) {
	return _Fdc2.Contract.Fdc2RequestHeaderStruct(&_Fdc2.TransactOpts, arg0)
}

// Fdc2ResponseHeaderStruct is a paid mutator transaction binding the contract method 0x4c2c8ceb.
//
// Solidity: function fdc2ResponseHeaderStruct((bytes32,bytes32,uint16,address,address[],uint64,uint64) ) returns()
func (_Fdc2 *Fdc2Transactor) Fdc2ResponseHeaderStruct(opts *bind.TransactOpts, arg0 IFdc2HubFdc2ResponseHeader) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "fdc2ResponseHeaderStruct", arg0)
}

// Fdc2ResponseHeaderStruct is a paid mutator transaction binding the contract method 0x4c2c8ceb.
//
// Solidity: function fdc2ResponseHeaderStruct((bytes32,bytes32,uint16,address,address[],uint64,uint64) ) returns()
func (_Fdc2 *Fdc2Session) Fdc2ResponseHeaderStruct(arg0 IFdc2HubFdc2ResponseHeader) (*types.Transaction, error) {
	return _Fdc2.Contract.Fdc2ResponseHeaderStruct(&_Fdc2.TransactOpts, arg0)
}

// Fdc2ResponseHeaderStruct is a paid mutator transaction binding the contract method 0x4c2c8ceb.
//
// Solidity: function fdc2ResponseHeaderStruct((bytes32,bytes32,uint16,address,address[],uint64,uint64) ) returns()
func (_Fdc2 *Fdc2TransactorSession) Fdc2ResponseHeaderStruct(arg0 IFdc2HubFdc2ResponseHeader) (*types.Transaction, error) {
	return _Fdc2.Contract.Fdc2ResponseHeaderStruct(&_Fdc2.TransactOpts, arg0)
}

// NativeNonceAccountConfiguredProofStruct is a paid mutator transaction binding the contract method 0x1b4bf46c.
//
// Solidity: function nativeNonceAccountConfiguredProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) ) returns()
func (_Fdc2 *Fdc2Transactor) NativeNonceAccountConfiguredProofStruct(opts *bind.TransactOpts, arg0 INativeNonceAccountConfiguredProof) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "nativeNonceAccountConfiguredProofStruct", arg0)
}

// NativeNonceAccountConfiguredProofStruct is a paid mutator transaction binding the contract method 0x1b4bf46c.
//
// Solidity: function nativeNonceAccountConfiguredProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) ) returns()
func (_Fdc2 *Fdc2Session) NativeNonceAccountConfiguredProofStruct(arg0 INativeNonceAccountConfiguredProof) (*types.Transaction, error) {
	return _Fdc2.Contract.NativeNonceAccountConfiguredProofStruct(&_Fdc2.TransactOpts, arg0)
}

// NativeNonceAccountConfiguredProofStruct is a paid mutator transaction binding the contract method 0x1b4bf46c.
//
// Solidity: function nativeNonceAccountConfiguredProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) ) returns()
func (_Fdc2 *Fdc2TransactorSession) NativeNonceAccountConfiguredProofStruct(arg0 INativeNonceAccountConfiguredProof) (*types.Transaction, error) {
	return _Fdc2.Contract.NativeNonceAccountConfiguredProofStruct(&_Fdc2.TransactOpts, arg0)
}

// NativeNonceAccountConfiguredRequestBodyStruct is a paid mutator transaction binding the contract method 0x1d7c6400.
//
// Solidity: function nativeNonceAccountConfiguredRequestBodyStruct((string,bytes[],uint64) ) returns()
func (_Fdc2 *Fdc2Transactor) NativeNonceAccountConfiguredRequestBodyStruct(opts *bind.TransactOpts, arg0 INativeNonceAccountConfiguredRequestBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "nativeNonceAccountConfiguredRequestBodyStruct", arg0)
}

// NativeNonceAccountConfiguredRequestBodyStruct is a paid mutator transaction binding the contract method 0x1d7c6400.
//
// Solidity: function nativeNonceAccountConfiguredRequestBodyStruct((string,bytes[],uint64) ) returns()
func (_Fdc2 *Fdc2Session) NativeNonceAccountConfiguredRequestBodyStruct(arg0 INativeNonceAccountConfiguredRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.NativeNonceAccountConfiguredRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// NativeNonceAccountConfiguredRequestBodyStruct is a paid mutator transaction binding the contract method 0x1d7c6400.
//
// Solidity: function nativeNonceAccountConfiguredRequestBodyStruct((string,bytes[],uint64) ) returns()
func (_Fdc2 *Fdc2TransactorSession) NativeNonceAccountConfiguredRequestBodyStruct(arg0 INativeNonceAccountConfiguredRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.NativeNonceAccountConfiguredRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// NativeNonceAccountConfiguredResponseBodyStruct is a paid mutator transaction binding the contract method 0x3adb48ea.
//
// Solidity: function nativeNonceAccountConfiguredResponseBodyStruct((uint8,uint64) ) returns()
func (_Fdc2 *Fdc2Transactor) NativeNonceAccountConfiguredResponseBodyStruct(opts *bind.TransactOpts, arg0 INativeNonceAccountConfiguredResponseBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "nativeNonceAccountConfiguredResponseBodyStruct", arg0)
}

// NativeNonceAccountConfiguredResponseBodyStruct is a paid mutator transaction binding the contract method 0x3adb48ea.
//
// Solidity: function nativeNonceAccountConfiguredResponseBodyStruct((uint8,uint64) ) returns()
func (_Fdc2 *Fdc2Session) NativeNonceAccountConfiguredResponseBodyStruct(arg0 INativeNonceAccountConfiguredResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.NativeNonceAccountConfiguredResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// NativeNonceAccountConfiguredResponseBodyStruct is a paid mutator transaction binding the contract method 0x3adb48ea.
//
// Solidity: function nativeNonceAccountConfiguredResponseBodyStruct((uint8,uint64) ) returns()
func (_Fdc2 *Fdc2TransactorSession) NativeNonceAccountConfiguredResponseBodyStruct(arg0 INativeNonceAccountConfiguredResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.NativeNonceAccountConfiguredResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// WalletFeeProofProofStruct is a paid mutator transaction binding the contract method 0xc61418cb.
//
// Solidity: function walletFeeProofProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,uint64,uint64),(uint64,uint256,uint256)) ) returns()
func (_Fdc2 *Fdc2Transactor) WalletFeeProofProofStruct(opts *bind.TransactOpts, arg0 IWalletFeeProofProof) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "walletFeeProofProofStruct", arg0)
}

// WalletFeeProofProofStruct is a paid mutator transaction binding the contract method 0xc61418cb.
//
// Solidity: function walletFeeProofProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,uint64,uint64),(uint64,uint256,uint256)) ) returns()
func (_Fdc2 *Fdc2Session) WalletFeeProofProofStruct(arg0 IWalletFeeProofProof) (*types.Transaction, error) {
	return _Fdc2.Contract.WalletFeeProofProofStruct(&_Fdc2.TransactOpts, arg0)
}

// WalletFeeProofProofStruct is a paid mutator transaction binding the contract method 0xc61418cb.
//
// Solidity: function walletFeeProofProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,uint64,uint64),(uint64,uint256,uint256)) ) returns()
func (_Fdc2 *Fdc2TransactorSession) WalletFeeProofProofStruct(arg0 IWalletFeeProofProof) (*types.Transaction, error) {
	return _Fdc2.Contract.WalletFeeProofProofStruct(&_Fdc2.TransactOpts, arg0)
}

// WalletFeeProofRequestBodyStruct is a paid mutator transaction binding the contract method 0x760297f8.
//
// Solidity: function walletFeeProofRequestBodyStruct((bytes32,string,uint64,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2Transactor) WalletFeeProofRequestBodyStruct(opts *bind.TransactOpts, arg0 IWalletFeeProofRequestBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "walletFeeProofRequestBodyStruct", arg0)
}

// WalletFeeProofRequestBodyStruct is a paid mutator transaction binding the contract method 0x760297f8.
//
// Solidity: function walletFeeProofRequestBodyStruct((bytes32,string,uint64,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2Session) WalletFeeProofRequestBodyStruct(arg0 IWalletFeeProofRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.WalletFeeProofRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// WalletFeeProofRequestBodyStruct is a paid mutator transaction binding the contract method 0x760297f8.
//
// Solidity: function walletFeeProofRequestBodyStruct((bytes32,string,uint64,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2TransactorSession) WalletFeeProofRequestBodyStruct(arg0 IWalletFeeProofRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.WalletFeeProofRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// WalletFeeProofResponseBodyStruct is a paid mutator transaction binding the contract method 0x4a03bdff.
//
// Solidity: function walletFeeProofResponseBodyStruct((uint64,uint256,uint256) ) returns()
func (_Fdc2 *Fdc2Transactor) WalletFeeProofResponseBodyStruct(opts *bind.TransactOpts, arg0 IWalletFeeProofResponseBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "walletFeeProofResponseBodyStruct", arg0)
}

// WalletFeeProofResponseBodyStruct is a paid mutator transaction binding the contract method 0x4a03bdff.
//
// Solidity: function walletFeeProofResponseBodyStruct((uint64,uint256,uint256) ) returns()
func (_Fdc2 *Fdc2Session) WalletFeeProofResponseBodyStruct(arg0 IWalletFeeProofResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.WalletFeeProofResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// WalletFeeProofResponseBodyStruct is a paid mutator transaction binding the contract method 0x4a03bdff.
//
// Solidity: function walletFeeProofResponseBodyStruct((uint64,uint256,uint256) ) returns()
func (_Fdc2 *Fdc2TransactorSession) WalletFeeProofResponseBodyStruct(arg0 IWalletFeeProofResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.WalletFeeProofResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// WalletPaymentStatusProofStruct is a paid mutator transaction binding the contract method 0xd71a2fde.
//
// Solidity: function walletPaymentStatusProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,bytes32),(string,bytes,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64)) ) returns()
func (_Fdc2 *Fdc2Transactor) WalletPaymentStatusProofStruct(opts *bind.TransactOpts, arg0 IWalletPaymentStatusProof) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "walletPaymentStatusProofStruct", arg0)
}

// WalletPaymentStatusProofStruct is a paid mutator transaction binding the contract method 0xd71a2fde.
//
// Solidity: function walletPaymentStatusProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,bytes32),(string,bytes,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64)) ) returns()
func (_Fdc2 *Fdc2Session) WalletPaymentStatusProofStruct(arg0 IWalletPaymentStatusProof) (*types.Transaction, error) {
	return _Fdc2.Contract.WalletPaymentStatusProofStruct(&_Fdc2.TransactOpts, arg0)
}

// WalletPaymentStatusProofStruct is a paid mutator transaction binding the contract method 0xd71a2fde.
//
// Solidity: function walletPaymentStatusProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,bytes32),(string,bytes,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64)) ) returns()
func (_Fdc2 *Fdc2TransactorSession) WalletPaymentStatusProofStruct(arg0 IWalletPaymentStatusProof) (*types.Transaction, error) {
	return _Fdc2.Contract.WalletPaymentStatusProofStruct(&_Fdc2.TransactOpts, arg0)
}

// WalletPaymentStatusRequestBodyStruct is a paid mutator transaction binding the contract method 0x25be73ca.
//
// Solidity: function walletPaymentStatusRequestBodyStruct((bytes32,string,uint64,bytes32) ) returns()
func (_Fdc2 *Fdc2Transactor) WalletPaymentStatusRequestBodyStruct(opts *bind.TransactOpts, arg0 IWalletPaymentStatusRequestBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "walletPaymentStatusRequestBodyStruct", arg0)
}

// WalletPaymentStatusRequestBodyStruct is a paid mutator transaction binding the contract method 0x25be73ca.
//
// Solidity: function walletPaymentStatusRequestBodyStruct((bytes32,string,uint64,bytes32) ) returns()
func (_Fdc2 *Fdc2Session) WalletPaymentStatusRequestBodyStruct(arg0 IWalletPaymentStatusRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.WalletPaymentStatusRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// WalletPaymentStatusRequestBodyStruct is a paid mutator transaction binding the contract method 0x25be73ca.
//
// Solidity: function walletPaymentStatusRequestBodyStruct((bytes32,string,uint64,bytes32) ) returns()
func (_Fdc2 *Fdc2TransactorSession) WalletPaymentStatusRequestBodyStruct(arg0 IWalletPaymentStatusRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.WalletPaymentStatusRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// WalletPaymentStatusResponseBodyStruct is a paid mutator transaction binding the contract method 0x8abc1d0e.
//
// Solidity: function walletPaymentStatusResponseBodyStruct((string,bytes,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2Transactor) WalletPaymentStatusResponseBodyStruct(opts *bind.TransactOpts, arg0 IWalletPaymentStatusResponseBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "walletPaymentStatusResponseBodyStruct", arg0)
}

// WalletPaymentStatusResponseBodyStruct is a paid mutator transaction binding the contract method 0x8abc1d0e.
//
// Solidity: function walletPaymentStatusResponseBodyStruct((string,bytes,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2Session) WalletPaymentStatusResponseBodyStruct(arg0 IWalletPaymentStatusResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.WalletPaymentStatusResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// WalletPaymentStatusResponseBodyStruct is a paid mutator transaction binding the contract method 0x8abc1d0e.
//
// Solidity: function walletPaymentStatusResponseBodyStruct((string,bytes,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2TransactorSession) WalletPaymentStatusResponseBodyStruct(arg0 IWalletPaymentStatusResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.WalletPaymentStatusResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}
