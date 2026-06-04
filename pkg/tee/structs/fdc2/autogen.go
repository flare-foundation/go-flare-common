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

// IPMWFeeProofProof is an auto generated low-level Go binding around an user-defined struct.
type IPMWFeeProofProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  IPMWFeeProofRequestBody
	ResponseBody IPMWFeeProofResponseBody
}

// IPMWFeeProofRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWFeeProofRequestBody struct {
	OpType         [32]byte
	SenderAddress  string
	FromNonce      uint64
	ToNonce        uint64
	UntilTimestamp uint64
}

// IPMWFeeProofResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWFeeProofResponseBody struct {
	ActualFee    *big.Int
	EstimatedFee *big.Int
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

// IPMWPaymentStatusProof is an auto generated low-level Go binding around an user-defined struct.
type IPMWPaymentStatusProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  IPMWPaymentStatusRequestBody
	ResponseBody IPMWPaymentStatusResponseBody
}

// IPMWPaymentStatusRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWPaymentStatusRequestBody struct {
	OpType        [32]byte
	SenderAddress string
	Nonce         uint64
	SubNonce      uint64
}

// IPMWPaymentStatusResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWPaymentStatusResponseBody struct {
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

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// Fdc2MetaData contains all meta data concerning the Fdc2 contract.
var Fdc2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"availabilityCheckTeeStateStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"}],\"internalType\":\"structIFdc2Hub.Fdc2RequestHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"requestBody\",\"type\":\"bytes\"}],\"internalType\":\"structIFdc2Hub.Fdc2AttestationRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"fdc2AttestationRequestStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"}],\"internalType\":\"structIFdc2Hub.Fdc2RequestHeader\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"fdc2RequestHeaderStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"fdc2ResponseHeaderStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"fromNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"untilTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWFeeProof.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimatedFee\",\"type\":\"uint256\"}],\"internalType\":\"structIPMWFeeProof.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWFeeProof.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwFeeProofProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"fromNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"untilTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWFeeProof.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwFeeProofRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"estimatedFee\",\"type\":\"uint256\"}],\"internalType\":\"structIPMWFeeProof.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwFeeProofResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigAccountConfigured.PMWMultisigAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwMultisigAccountConfiguredProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwMultisigAccountConfiguredRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumIPMWMultisigAccountConfigured.PMWMultisigAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwMultisigAccountConfiguredResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subNonce\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWPaymentStatus.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"transactionStatus\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transactionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWPaymentStatus.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWPaymentStatus.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwPaymentStatusProofStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subNonce\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWPaymentStatus.RequestBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwPaymentStatusRequestBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"transactionStatus\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"revertReason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transactionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWPaymentStatus.ResponseBody\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"pmwPaymentStatusResponseBodyStruct\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// PmwFeeProofProofStruct is a paid mutator transaction binding the contract method 0xc692ef5b.
//
// Solidity: function pmwFeeProofProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,uint64,uint64),(uint256,uint256)) ) returns()
func (_Fdc2 *Fdc2Transactor) PmwFeeProofProofStruct(opts *bind.TransactOpts, arg0 IPMWFeeProofProof) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "pmwFeeProofProofStruct", arg0)
}

// PmwFeeProofProofStruct is a paid mutator transaction binding the contract method 0xc692ef5b.
//
// Solidity: function pmwFeeProofProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,uint64,uint64),(uint256,uint256)) ) returns()
func (_Fdc2 *Fdc2Session) PmwFeeProofProofStruct(arg0 IPMWFeeProofProof) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwFeeProofProofStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwFeeProofProofStruct is a paid mutator transaction binding the contract method 0xc692ef5b.
//
// Solidity: function pmwFeeProofProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,uint64,uint64),(uint256,uint256)) ) returns()
func (_Fdc2 *Fdc2TransactorSession) PmwFeeProofProofStruct(arg0 IPMWFeeProofProof) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwFeeProofProofStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwFeeProofRequestBodyStruct is a paid mutator transaction binding the contract method 0x362fe99f.
//
// Solidity: function pmwFeeProofRequestBodyStruct((bytes32,string,uint64,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2Transactor) PmwFeeProofRequestBodyStruct(opts *bind.TransactOpts, arg0 IPMWFeeProofRequestBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "pmwFeeProofRequestBodyStruct", arg0)
}

// PmwFeeProofRequestBodyStruct is a paid mutator transaction binding the contract method 0x362fe99f.
//
// Solidity: function pmwFeeProofRequestBodyStruct((bytes32,string,uint64,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2Session) PmwFeeProofRequestBodyStruct(arg0 IPMWFeeProofRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwFeeProofRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwFeeProofRequestBodyStruct is a paid mutator transaction binding the contract method 0x362fe99f.
//
// Solidity: function pmwFeeProofRequestBodyStruct((bytes32,string,uint64,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2TransactorSession) PmwFeeProofRequestBodyStruct(arg0 IPMWFeeProofRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwFeeProofRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwFeeProofResponseBodyStruct is a paid mutator transaction binding the contract method 0x8bc5ed5e.
//
// Solidity: function pmwFeeProofResponseBodyStruct((uint256,uint256) ) returns()
func (_Fdc2 *Fdc2Transactor) PmwFeeProofResponseBodyStruct(opts *bind.TransactOpts, arg0 IPMWFeeProofResponseBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "pmwFeeProofResponseBodyStruct", arg0)
}

// PmwFeeProofResponseBodyStruct is a paid mutator transaction binding the contract method 0x8bc5ed5e.
//
// Solidity: function pmwFeeProofResponseBodyStruct((uint256,uint256) ) returns()
func (_Fdc2 *Fdc2Session) PmwFeeProofResponseBodyStruct(arg0 IPMWFeeProofResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwFeeProofResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwFeeProofResponseBodyStruct is a paid mutator transaction binding the contract method 0x8bc5ed5e.
//
// Solidity: function pmwFeeProofResponseBodyStruct((uint256,uint256) ) returns()
func (_Fdc2 *Fdc2TransactorSession) PmwFeeProofResponseBodyStruct(arg0 IPMWFeeProofResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwFeeProofResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwMultisigAccountConfiguredProofStruct is a paid mutator transaction binding the contract method 0x1cfce1b7.
//
// Solidity: function pmwMultisigAccountConfiguredProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) ) returns()
func (_Fdc2 *Fdc2Transactor) PmwMultisigAccountConfiguredProofStruct(opts *bind.TransactOpts, arg0 IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "pmwMultisigAccountConfiguredProofStruct", arg0)
}

// PmwMultisigAccountConfiguredProofStruct is a paid mutator transaction binding the contract method 0x1cfce1b7.
//
// Solidity: function pmwMultisigAccountConfiguredProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) ) returns()
func (_Fdc2 *Fdc2Session) PmwMultisigAccountConfiguredProofStruct(arg0 IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwMultisigAccountConfiguredProofStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwMultisigAccountConfiguredProofStruct is a paid mutator transaction binding the contract method 0x1cfce1b7.
//
// Solidity: function pmwMultisigAccountConfiguredProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) ) returns()
func (_Fdc2 *Fdc2TransactorSession) PmwMultisigAccountConfiguredProofStruct(arg0 IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwMultisigAccountConfiguredProofStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwMultisigAccountConfiguredRequestBodyStruct is a paid mutator transaction binding the contract method 0xc801f3d6.
//
// Solidity: function pmwMultisigAccountConfiguredRequestBodyStruct((string,bytes[],uint64) ) returns()
func (_Fdc2 *Fdc2Transactor) PmwMultisigAccountConfiguredRequestBodyStruct(opts *bind.TransactOpts, arg0 IPMWMultisigAccountConfiguredRequestBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "pmwMultisigAccountConfiguredRequestBodyStruct", arg0)
}

// PmwMultisigAccountConfiguredRequestBodyStruct is a paid mutator transaction binding the contract method 0xc801f3d6.
//
// Solidity: function pmwMultisigAccountConfiguredRequestBodyStruct((string,bytes[],uint64) ) returns()
func (_Fdc2 *Fdc2Session) PmwMultisigAccountConfiguredRequestBodyStruct(arg0 IPMWMultisigAccountConfiguredRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwMultisigAccountConfiguredRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwMultisigAccountConfiguredRequestBodyStruct is a paid mutator transaction binding the contract method 0xc801f3d6.
//
// Solidity: function pmwMultisigAccountConfiguredRequestBodyStruct((string,bytes[],uint64) ) returns()
func (_Fdc2 *Fdc2TransactorSession) PmwMultisigAccountConfiguredRequestBodyStruct(arg0 IPMWMultisigAccountConfiguredRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwMultisigAccountConfiguredRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwMultisigAccountConfiguredResponseBodyStruct is a paid mutator transaction binding the contract method 0xe9681a8f.
//
// Solidity: function pmwMultisigAccountConfiguredResponseBodyStruct((uint8,uint64) ) returns()
func (_Fdc2 *Fdc2Transactor) PmwMultisigAccountConfiguredResponseBodyStruct(opts *bind.TransactOpts, arg0 IPMWMultisigAccountConfiguredResponseBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "pmwMultisigAccountConfiguredResponseBodyStruct", arg0)
}

// PmwMultisigAccountConfiguredResponseBodyStruct is a paid mutator transaction binding the contract method 0xe9681a8f.
//
// Solidity: function pmwMultisigAccountConfiguredResponseBodyStruct((uint8,uint64) ) returns()
func (_Fdc2 *Fdc2Session) PmwMultisigAccountConfiguredResponseBodyStruct(arg0 IPMWMultisigAccountConfiguredResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwMultisigAccountConfiguredResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwMultisigAccountConfiguredResponseBodyStruct is a paid mutator transaction binding the contract method 0xe9681a8f.
//
// Solidity: function pmwMultisigAccountConfiguredResponseBodyStruct((uint8,uint64) ) returns()
func (_Fdc2 *Fdc2TransactorSession) PmwMultisigAccountConfiguredResponseBodyStruct(arg0 IPMWMultisigAccountConfiguredResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwMultisigAccountConfiguredResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwPaymentStatusProofStruct is a paid mutator transaction binding the contract method 0x1115acdc.
//
// Solidity: function pmwPaymentStatusProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,uint64),(string,bytes,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64)) ) returns()
func (_Fdc2 *Fdc2Transactor) PmwPaymentStatusProofStruct(opts *bind.TransactOpts, arg0 IPMWPaymentStatusProof) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "pmwPaymentStatusProofStruct", arg0)
}

// PmwPaymentStatusProofStruct is a paid mutator transaction binding the contract method 0x1115acdc.
//
// Solidity: function pmwPaymentStatusProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,uint64),(string,bytes,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64)) ) returns()
func (_Fdc2 *Fdc2Session) PmwPaymentStatusProofStruct(arg0 IPMWPaymentStatusProof) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwPaymentStatusProofStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwPaymentStatusProofStruct is a paid mutator transaction binding the contract method 0x1115acdc.
//
// Solidity: function pmwPaymentStatusProofStruct(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,uint64),(string,bytes,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64)) ) returns()
func (_Fdc2 *Fdc2TransactorSession) PmwPaymentStatusProofStruct(arg0 IPMWPaymentStatusProof) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwPaymentStatusProofStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwPaymentStatusRequestBodyStruct is a paid mutator transaction binding the contract method 0x7bd394cd.
//
// Solidity: function pmwPaymentStatusRequestBodyStruct((bytes32,string,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2Transactor) PmwPaymentStatusRequestBodyStruct(opts *bind.TransactOpts, arg0 IPMWPaymentStatusRequestBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "pmwPaymentStatusRequestBodyStruct", arg0)
}

// PmwPaymentStatusRequestBodyStruct is a paid mutator transaction binding the contract method 0x7bd394cd.
//
// Solidity: function pmwPaymentStatusRequestBodyStruct((bytes32,string,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2Session) PmwPaymentStatusRequestBodyStruct(arg0 IPMWPaymentStatusRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwPaymentStatusRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwPaymentStatusRequestBodyStruct is a paid mutator transaction binding the contract method 0x7bd394cd.
//
// Solidity: function pmwPaymentStatusRequestBodyStruct((bytes32,string,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2TransactorSession) PmwPaymentStatusRequestBodyStruct(arg0 IPMWPaymentStatusRequestBody) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwPaymentStatusRequestBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwPaymentStatusResponseBodyStruct is a paid mutator transaction binding the contract method 0xc18329f2.
//
// Solidity: function pmwPaymentStatusResponseBodyStruct((string,bytes,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2Transactor) PmwPaymentStatusResponseBodyStruct(opts *bind.TransactOpts, arg0 IPMWPaymentStatusResponseBody) (*types.Transaction, error) {
	return _Fdc2.contract.Transact(opts, "pmwPaymentStatusResponseBodyStruct", arg0)
}

// PmwPaymentStatusResponseBodyStruct is a paid mutator transaction binding the contract method 0xc18329f2.
//
// Solidity: function pmwPaymentStatusResponseBodyStruct((string,bytes,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2Session) PmwPaymentStatusResponseBodyStruct(arg0 IPMWPaymentStatusResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwPaymentStatusResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}

// PmwPaymentStatusResponseBodyStruct is a paid mutator transaction binding the contract method 0xc18329f2.
//
// Solidity: function pmwPaymentStatusResponseBodyStruct((string,bytes,uint256,uint256,bytes32,uint8,string,uint256,uint256,bytes32,uint64,uint64) ) returns()
func (_Fdc2 *Fdc2TransactorSession) PmwPaymentStatusResponseBodyStruct(arg0 IPMWPaymentStatusResponseBody) (*types.Transaction, error) {
	return _Fdc2.Contract.PmwPaymentStatusResponseBodyStruct(&_Fdc2.TransactOpts, arg0)
}
