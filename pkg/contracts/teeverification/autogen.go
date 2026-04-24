// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teeverification

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

// IMachineManagerTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
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

// TeeVerificationMetaData contains all meta data concerning the TeeVerification contract.
var TeeVerificationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccountAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengeTs\",\"type\":\"uint256\"}],\"name\":\"ChallengeExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestBody\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningPolicy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTs\",\"type\":\"uint256\"}],\"name\":\"AvailabilityCheckValidityExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"CosignersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"availabilityCheckValidityDurationSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"signingPolicyValidityDurationInRewardEpochs\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"challengeValidityDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"SettingsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"name\":\"TeeAttestationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"confirmAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getAvailabilityCheckValidity\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_endTs\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_lastSigningPolicyId\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCosigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSettings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_availabilityCheckValidityDurationSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengeValidityDurationSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_testOnTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proofOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestAvailabilityCheckAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_accountAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_testOnTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proofOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestPMWMultisigAccountConfiguredAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestTeeAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"setCosigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_availabilityCheckValidityDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicyValidityDurationInRewardEpochs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_challengeValidityDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"updateSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAvailabilityCheckProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigAccountConfigured.PMWMultisigAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyPMWMultisigAccountConfiguredProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeVerificationABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeVerificationMetaData.ABI instead.
var TeeVerificationABI = TeeVerificationMetaData.ABI

// TeeVerification is an auto generated Go binding around an Ethereum contract.
type TeeVerification struct {
	TeeVerificationCaller     // Read-only binding to the contract
	TeeVerificationTransactor // Write-only binding to the contract
	TeeVerificationFilterer   // Log filterer for contract events
}

// TeeVerificationCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeVerificationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVerificationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeVerificationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVerificationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeVerificationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeVerificationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeVerificationSession struct {
	Contract     *TeeVerification  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeVerificationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeVerificationCallerSession struct {
	Contract *TeeVerificationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TeeVerificationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeVerificationTransactorSession struct {
	Contract     *TeeVerificationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TeeVerificationRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeVerificationRaw struct {
	Contract *TeeVerification // Generic contract binding to access the raw methods on
}

// TeeVerificationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeVerificationCallerRaw struct {
	Contract *TeeVerificationCaller // Generic read-only contract binding to access the raw methods on
}

// TeeVerificationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeVerificationTransactorRaw struct {
	Contract *TeeVerificationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeVerification creates a new instance of TeeVerification, bound to a specific deployed contract.
func NewTeeVerification(address common.Address, backend bind.ContractBackend) (*TeeVerification, error) {
	contract, err := bindTeeVerification(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeVerification{TeeVerificationCaller: TeeVerificationCaller{contract: contract}, TeeVerificationTransactor: TeeVerificationTransactor{contract: contract}, TeeVerificationFilterer: TeeVerificationFilterer{contract: contract}}, nil
}

// NewTeeVerificationCaller creates a new read-only instance of TeeVerification, bound to a specific deployed contract.
func NewTeeVerificationCaller(address common.Address, caller bind.ContractCaller) (*TeeVerificationCaller, error) {
	contract, err := bindTeeVerification(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeVerificationCaller{contract: contract}, nil
}

// NewTeeVerificationTransactor creates a new write-only instance of TeeVerification, bound to a specific deployed contract.
func NewTeeVerificationTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeVerificationTransactor, error) {
	contract, err := bindTeeVerification(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeVerificationTransactor{contract: contract}, nil
}

// NewTeeVerificationFilterer creates a new log filterer instance of TeeVerification, bound to a specific deployed contract.
func NewTeeVerificationFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeVerificationFilterer, error) {
	contract, err := bindTeeVerification(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeVerificationFilterer{contract: contract}, nil
}

// bindTeeVerification binds a generic wrapper to an already deployed contract.
func bindTeeVerification(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeVerificationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeVerification *TeeVerificationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeVerification.Contract.TeeVerificationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeVerification *TeeVerificationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVerification.Contract.TeeVerificationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeVerification *TeeVerificationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeVerification.Contract.TeeVerificationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeVerification *TeeVerificationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeVerification.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeVerification *TeeVerificationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVerification.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeVerification *TeeVerificationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeVerification.Contract.contract.Transact(opts, method, params...)
}

// GetAvailabilityCheckValidity is a free data retrieval call binding the contract method 0x3e1f22f5.
//
// Solidity: function getAvailabilityCheckValidity(address _teeId) view returns(uint64 _endTs, uint32 _lastSigningPolicyId)
func (_TeeVerification *TeeVerificationCaller) GetAvailabilityCheckValidity(opts *bind.CallOpts, _teeId common.Address) (struct {
	EndTs               uint64
	LastSigningPolicyId uint32
}, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "getAvailabilityCheckValidity", _teeId)

	outstruct := new(struct {
		EndTs               uint64
		LastSigningPolicyId uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EndTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.LastSigningPolicyId = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetAvailabilityCheckValidity is a free data retrieval call binding the contract method 0x3e1f22f5.
//
// Solidity: function getAvailabilityCheckValidity(address _teeId) view returns(uint64 _endTs, uint32 _lastSigningPolicyId)
func (_TeeVerification *TeeVerificationSession) GetAvailabilityCheckValidity(_teeId common.Address) (struct {
	EndTs               uint64
	LastSigningPolicyId uint32
}, error) {
	return _TeeVerification.Contract.GetAvailabilityCheckValidity(&_TeeVerification.CallOpts, _teeId)
}

// GetAvailabilityCheckValidity is a free data retrieval call binding the contract method 0x3e1f22f5.
//
// Solidity: function getAvailabilityCheckValidity(address _teeId) view returns(uint64 _endTs, uint32 _lastSigningPolicyId)
func (_TeeVerification *TeeVerificationCallerSession) GetAvailabilityCheckValidity(_teeId common.Address) (struct {
	EndTs               uint64
	LastSigningPolicyId uint32
}, error) {
	return _TeeVerification.Contract.GetAvailabilityCheckValidity(&_TeeVerification.CallOpts, _teeId)
}

// GetCosigners is a free data retrieval call binding the contract method 0x86b8d27e.
//
// Solidity: function getCosigners() view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_TeeVerification *TeeVerificationCaller) GetCosigners(opts *bind.CallOpts) (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "getCosigners")

	outstruct := new(struct {
		Cosigners          []common.Address
		CosignersThreshold uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cosigners = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.CosignersThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetCosigners is a free data retrieval call binding the contract method 0x86b8d27e.
//
// Solidity: function getCosigners() view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_TeeVerification *TeeVerificationSession) GetCosigners() (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	return _TeeVerification.Contract.GetCosigners(&_TeeVerification.CallOpts)
}

// GetCosigners is a free data retrieval call binding the contract method 0x86b8d27e.
//
// Solidity: function getCosigners() view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_TeeVerification *TeeVerificationCallerSession) GetCosigners() (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	return _TeeVerification.Contract.GetCosigners(&_TeeVerification.CallOpts)
}

// GetSettings is a free data retrieval call binding the contract method 0x85b4bb53.
//
// Solidity: function getSettings() view returns(uint256 _availabilityCheckValidityDurationSeconds, uint256 _challengeValidityDurationSeconds)
func (_TeeVerification *TeeVerificationCaller) GetSettings(opts *bind.CallOpts) (struct {
	AvailabilityCheckValidityDurationSeconds *big.Int
	ChallengeValidityDurationSeconds         *big.Int
}, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "getSettings")

	outstruct := new(struct {
		AvailabilityCheckValidityDurationSeconds *big.Int
		ChallengeValidityDurationSeconds         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AvailabilityCheckValidityDurationSeconds = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChallengeValidityDurationSeconds = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSettings is a free data retrieval call binding the contract method 0x85b4bb53.
//
// Solidity: function getSettings() view returns(uint256 _availabilityCheckValidityDurationSeconds, uint256 _challengeValidityDurationSeconds)
func (_TeeVerification *TeeVerificationSession) GetSettings() (struct {
	AvailabilityCheckValidityDurationSeconds *big.Int
	ChallengeValidityDurationSeconds         *big.Int
}, error) {
	return _TeeVerification.Contract.GetSettings(&_TeeVerification.CallOpts)
}

// GetSettings is a free data retrieval call binding the contract method 0x85b4bb53.
//
// Solidity: function getSettings() view returns(uint256 _availabilityCheckValidityDurationSeconds, uint256 _challengeValidityDurationSeconds)
func (_TeeVerification *TeeVerificationCallerSession) GetSettings() (struct {
	AvailabilityCheckValidityDurationSeconds *big.Int
	ChallengeValidityDurationSeconds         *big.Int
}, error) {
	return _TeeVerification.Contract.GetSettings(&_TeeVerification.CallOpts)
}

// ConfirmAvailability is a paid mutator transaction binding the contract method 0x1f78297f.
//
// Solidity: function confirmAvailability(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeVerification *TeeVerificationTransactor) ConfirmAvailability(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "confirmAvailability", _proof)
}

// ConfirmAvailability is a paid mutator transaction binding the contract method 0x1f78297f.
//
// Solidity: function confirmAvailability(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeVerification *TeeVerificationSession) ConfirmAvailability(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.ConfirmAvailability(&_TeeVerification.TransactOpts, _proof)
}

// ConfirmAvailability is a paid mutator transaction binding the contract method 0x1f78297f.
//
// Solidity: function confirmAvailability(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeVerification *TeeVerificationTransactorSession) ConfirmAvailability(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.ConfirmAvailability(&_TeeVerification.TransactOpts, _proof)
}

// RequestAvailabilityCheckAttestation is a paid mutator transaction binding the contract method 0xa3d60348.
//
// Solidity: function requestAvailabilityCheckAttestation(address _teeId, bytes32 _instructionId, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_TeeVerification *TeeVerificationTransactor) RequestAvailabilityCheckAttestation(opts *bind.TransactOpts, _teeId common.Address, _instructionId [32]byte, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "requestAvailabilityCheckAttestation", _teeId, _instructionId, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestAvailabilityCheckAttestation is a paid mutator transaction binding the contract method 0xa3d60348.
//
// Solidity: function requestAvailabilityCheckAttestation(address _teeId, bytes32 _instructionId, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_TeeVerification *TeeVerificationSession) RequestAvailabilityCheckAttestation(_teeId common.Address, _instructionId [32]byte, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestAvailabilityCheckAttestation(&_TeeVerification.TransactOpts, _teeId, _instructionId, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestAvailabilityCheckAttestation is a paid mutator transaction binding the contract method 0xa3d60348.
//
// Solidity: function requestAvailabilityCheckAttestation(address _teeId, bytes32 _instructionId, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_TeeVerification *TeeVerificationTransactorSession) RequestAvailabilityCheckAttestation(_teeId common.Address, _instructionId [32]byte, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestAvailabilityCheckAttestation(&_TeeVerification.TransactOpts, _teeId, _instructionId, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestPMWMultisigAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0x874780b9.
//
// Solidity: function requestPMWMultisigAccountConfiguredAttestation(bytes32 _walletId, bytes32 _sourceId, string _accountAddress, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_TeeVerification *TeeVerificationTransactor) RequestPMWMultisigAccountConfiguredAttestation(opts *bind.TransactOpts, _walletId [32]byte, _sourceId [32]byte, _accountAddress string, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "requestPMWMultisigAccountConfiguredAttestation", _walletId, _sourceId, _accountAddress, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestPMWMultisigAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0x874780b9.
//
// Solidity: function requestPMWMultisigAccountConfiguredAttestation(bytes32 _walletId, bytes32 _sourceId, string _accountAddress, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_TeeVerification *TeeVerificationSession) RequestPMWMultisigAccountConfiguredAttestation(_walletId [32]byte, _sourceId [32]byte, _accountAddress string, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestPMWMultisigAccountConfiguredAttestation(&_TeeVerification.TransactOpts, _walletId, _sourceId, _accountAddress, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestPMWMultisigAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0x874780b9.
//
// Solidity: function requestPMWMultisigAccountConfiguredAttestation(bytes32 _walletId, bytes32 _sourceId, string _accountAddress, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_TeeVerification *TeeVerificationTransactorSession) RequestPMWMultisigAccountConfiguredAttestation(_walletId [32]byte, _sourceId [32]byte, _accountAddress string, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestPMWMultisigAccountConfiguredAttestation(&_TeeVerification.TransactOpts, _walletId, _sourceId, _accountAddress, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestTeeAttestation is a paid mutator transaction binding the contract method 0x7d8dcfb5.
//
// Solidity: function requestTeeAttestation(address _teeId, address _claimBackAddress) payable returns()
func (_TeeVerification *TeeVerificationTransactor) RequestTeeAttestation(opts *bind.TransactOpts, _teeId common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "requestTeeAttestation", _teeId, _claimBackAddress)
}

// RequestTeeAttestation is a paid mutator transaction binding the contract method 0x7d8dcfb5.
//
// Solidity: function requestTeeAttestation(address _teeId, address _claimBackAddress) payable returns()
func (_TeeVerification *TeeVerificationSession) RequestTeeAttestation(_teeId common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestTeeAttestation(&_TeeVerification.TransactOpts, _teeId, _claimBackAddress)
}

// RequestTeeAttestation is a paid mutator transaction binding the contract method 0x7d8dcfb5.
//
// Solidity: function requestTeeAttestation(address _teeId, address _claimBackAddress) payable returns()
func (_TeeVerification *TeeVerificationTransactorSession) RequestTeeAttestation(_teeId common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestTeeAttestation(&_TeeVerification.TransactOpts, _teeId, _claimBackAddress)
}

// SetCosigners is a paid mutator transaction binding the contract method 0xf0781728.
//
// Solidity: function setCosigners(address[] _cosigners, uint64 _cosignersThreshold) returns()
func (_TeeVerification *TeeVerificationTransactor) SetCosigners(opts *bind.TransactOpts, _cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "setCosigners", _cosigners, _cosignersThreshold)
}

// SetCosigners is a paid mutator transaction binding the contract method 0xf0781728.
//
// Solidity: function setCosigners(address[] _cosigners, uint64 _cosignersThreshold) returns()
func (_TeeVerification *TeeVerificationSession) SetCosigners(_cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _TeeVerification.Contract.SetCosigners(&_TeeVerification.TransactOpts, _cosigners, _cosignersThreshold)
}

// SetCosigners is a paid mutator transaction binding the contract method 0xf0781728.
//
// Solidity: function setCosigners(address[] _cosigners, uint64 _cosignersThreshold) returns()
func (_TeeVerification *TeeVerificationTransactorSession) SetCosigners(_cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _TeeVerification.Contract.SetCosigners(&_TeeVerification.TransactOpts, _cosigners, _cosignersThreshold)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x34ef6f09.
//
// Solidity: function updateSettings(uint64 _availabilityCheckValidityDurationSeconds, uint64 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds) returns()
func (_TeeVerification *TeeVerificationTransactor) UpdateSettings(opts *bind.TransactOpts, _availabilityCheckValidityDurationSeconds uint64, _signingPolicyValidityDurationInRewardEpochs uint64, _challengeValidityDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "updateSettings", _availabilityCheckValidityDurationSeconds, _signingPolicyValidityDurationInRewardEpochs, _challengeValidityDurationSeconds)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x34ef6f09.
//
// Solidity: function updateSettings(uint64 _availabilityCheckValidityDurationSeconds, uint64 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds) returns()
func (_TeeVerification *TeeVerificationSession) UpdateSettings(_availabilityCheckValidityDurationSeconds uint64, _signingPolicyValidityDurationInRewardEpochs uint64, _challengeValidityDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeVerification.Contract.UpdateSettings(&_TeeVerification.TransactOpts, _availabilityCheckValidityDurationSeconds, _signingPolicyValidityDurationInRewardEpochs, _challengeValidityDurationSeconds)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x34ef6f09.
//
// Solidity: function updateSettings(uint64 _availabilityCheckValidityDurationSeconds, uint64 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds) returns()
func (_TeeVerification *TeeVerificationTransactorSession) UpdateSettings(_availabilityCheckValidityDurationSeconds uint64, _signingPolicyValidityDurationInRewardEpochs uint64, _challengeValidityDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeVerification.Contract.UpdateSettings(&_TeeVerification.TransactOpts, _availabilityCheckValidityDurationSeconds, _signingPolicyValidityDurationInRewardEpochs, _challengeValidityDurationSeconds)
}

// VerifyAvailabilityCheckProof is a paid mutator transaction binding the contract method 0xd35056ba.
//
// Solidity: function verifyAvailabilityCheckProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns(bool)
func (_TeeVerification *TeeVerificationTransactor) VerifyAvailabilityCheckProof(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "verifyAvailabilityCheckProof", _proof)
}

// VerifyAvailabilityCheckProof is a paid mutator transaction binding the contract method 0xd35056ba.
//
// Solidity: function verifyAvailabilityCheckProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns(bool)
func (_TeeVerification *TeeVerificationSession) VerifyAvailabilityCheckProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.VerifyAvailabilityCheckProof(&_TeeVerification.TransactOpts, _proof)
}

// VerifyAvailabilityCheckProof is a paid mutator transaction binding the contract method 0xd35056ba.
//
// Solidity: function verifyAvailabilityCheckProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns(bool)
func (_TeeVerification *TeeVerificationTransactorSession) VerifyAvailabilityCheckProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.VerifyAvailabilityCheckProof(&_TeeVerification.TransactOpts, _proof)
}

// VerifyPMWMultisigAccountConfiguredProof is a paid mutator transaction binding the contract method 0x6984143d.
//
// Solidity: function verifyPMWMultisigAccountConfiguredProof(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns(bool)
func (_TeeVerification *TeeVerificationTransactor) VerifyPMWMultisigAccountConfiguredProof(opts *bind.TransactOpts, _walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "verifyPMWMultisigAccountConfiguredProof", _walletId, _proof)
}

// VerifyPMWMultisigAccountConfiguredProof is a paid mutator transaction binding the contract method 0x6984143d.
//
// Solidity: function verifyPMWMultisigAccountConfiguredProof(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns(bool)
func (_TeeVerification *TeeVerificationSession) VerifyPMWMultisigAccountConfiguredProof(_walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.VerifyPMWMultisigAccountConfiguredProof(&_TeeVerification.TransactOpts, _walletId, _proof)
}

// VerifyPMWMultisigAccountConfiguredProof is a paid mutator transaction binding the contract method 0x6984143d.
//
// Solidity: function verifyPMWMultisigAccountConfiguredProof(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns(bool)
func (_TeeVerification *TeeVerificationTransactorSession) VerifyPMWMultisigAccountConfiguredProof(_walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.VerifyPMWMultisigAccountConfiguredProof(&_TeeVerification.TransactOpts, _walletId, _proof)
}

// TeeVerificationAvailabilityCheckValidityExtendedIterator is returned from FilterAvailabilityCheckValidityExtended and is used to iterate over the raw logs and unpacked data for AvailabilityCheckValidityExtended events raised by the TeeVerification contract.
type TeeVerificationAvailabilityCheckValidityExtendedIterator struct {
	Event *TeeVerificationAvailabilityCheckValidityExtended // Event containing the contract specifics and raw log

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
func (it *TeeVerificationAvailabilityCheckValidityExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationAvailabilityCheckValidityExtended)
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
		it.Event = new(TeeVerificationAvailabilityCheckValidityExtended)
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
func (it *TeeVerificationAvailabilityCheckValidityExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationAvailabilityCheckValidityExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationAvailabilityCheckValidityExtended represents a AvailabilityCheckValidityExtended event raised by the TeeVerification contract.
type TeeVerificationAvailabilityCheckValidityExtended struct {
	TeeId common.Address
	Owner common.Address
	EndTs *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAvailabilityCheckValidityExtended is a free log retrieval operation binding the contract event 0xbcea1901d8026bcd3401ce1c443ba0a3e5e16680271d1dcc84cfdd36a0ca44b3.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (_TeeVerification *TeeVerificationFilterer) FilterAvailabilityCheckValidityExtended(opts *bind.FilterOpts, teeId []common.Address, owner []common.Address) (*TeeVerificationAvailabilityCheckValidityExtendedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "AvailabilityCheckValidityExtended", teeIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TeeVerificationAvailabilityCheckValidityExtendedIterator{contract: _TeeVerification.contract, event: "AvailabilityCheckValidityExtended", logs: logs, sub: sub}, nil
}

// WatchAvailabilityCheckValidityExtended is a free log subscription operation binding the contract event 0xbcea1901d8026bcd3401ce1c443ba0a3e5e16680271d1dcc84cfdd36a0ca44b3.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (_TeeVerification *TeeVerificationFilterer) WatchAvailabilityCheckValidityExtended(opts *bind.WatchOpts, sink chan<- *TeeVerificationAvailabilityCheckValidityExtended, teeId []common.Address, owner []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "AvailabilityCheckValidityExtended", teeIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationAvailabilityCheckValidityExtended)
				if err := _TeeVerification.contract.UnpackLog(event, "AvailabilityCheckValidityExtended", log); err != nil {
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

// ParseAvailabilityCheckValidityExtended is a log parse operation binding the contract event 0xbcea1901d8026bcd3401ce1c443ba0a3e5e16680271d1dcc84cfdd36a0ca44b3.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (_TeeVerification *TeeVerificationFilterer) ParseAvailabilityCheckValidityExtended(log types.Log) (*TeeVerificationAvailabilityCheckValidityExtended, error) {
	event := new(TeeVerificationAvailabilityCheckValidityExtended)
	if err := _TeeVerification.contract.UnpackLog(event, "AvailabilityCheckValidityExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVerificationCosignersSetIterator is returned from FilterCosignersSet and is used to iterate over the raw logs and unpacked data for CosignersSet events raised by the TeeVerification contract.
type TeeVerificationCosignersSetIterator struct {
	Event *TeeVerificationCosignersSet // Event containing the contract specifics and raw log

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
func (it *TeeVerificationCosignersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationCosignersSet)
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
		it.Event = new(TeeVerificationCosignersSet)
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
func (it *TeeVerificationCosignersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationCosignersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationCosignersSet represents a CosignersSet event raised by the TeeVerification contract.
type TeeVerificationCosignersSet struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCosignersSet is a free log retrieval operation binding the contract event 0x19ba4a7aa64550135f3db2da04c8c13f2331d3cc9b32dd06339ba2b48a71e9b8.
//
// Solidity: event CosignersSet(address[] cosigners, uint64 cosignersThreshold)
func (_TeeVerification *TeeVerificationFilterer) FilterCosignersSet(opts *bind.FilterOpts) (*TeeVerificationCosignersSetIterator, error) {

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "CosignersSet")
	if err != nil {
		return nil, err
	}
	return &TeeVerificationCosignersSetIterator{contract: _TeeVerification.contract, event: "CosignersSet", logs: logs, sub: sub}, nil
}

// WatchCosignersSet is a free log subscription operation binding the contract event 0x19ba4a7aa64550135f3db2da04c8c13f2331d3cc9b32dd06339ba2b48a71e9b8.
//
// Solidity: event CosignersSet(address[] cosigners, uint64 cosignersThreshold)
func (_TeeVerification *TeeVerificationFilterer) WatchCosignersSet(opts *bind.WatchOpts, sink chan<- *TeeVerificationCosignersSet) (event.Subscription, error) {

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "CosignersSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationCosignersSet)
				if err := _TeeVerification.contract.UnpackLog(event, "CosignersSet", log); err != nil {
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

// ParseCosignersSet is a log parse operation binding the contract event 0x19ba4a7aa64550135f3db2da04c8c13f2331d3cc9b32dd06339ba2b48a71e9b8.
//
// Solidity: event CosignersSet(address[] cosigners, uint64 cosignersThreshold)
func (_TeeVerification *TeeVerificationFilterer) ParseCosignersSet(log types.Log) (*TeeVerificationCosignersSet, error) {
	event := new(TeeVerificationCosignersSet)
	if err := _TeeVerification.contract.UnpackLog(event, "CosignersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVerificationGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeVerification contract.
type TeeVerificationGovernanceCallTimelockedIterator struct {
	Event *TeeVerificationGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeeVerificationGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationGovernanceCallTimelocked)
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
		it.Event = new(TeeVerificationGovernanceCallTimelocked)
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
func (it *TeeVerificationGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeVerification contract.
type TeeVerificationGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeeVerification *TeeVerificationFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeVerificationGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeVerificationGovernanceCallTimelockedIterator{contract: _TeeVerification.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeeVerification *TeeVerificationFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeVerificationGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationGovernanceCallTimelocked)
				if err := _TeeVerification.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeeVerification *TeeVerificationFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeVerificationGovernanceCallTimelocked, error) {
	event := new(TeeVerificationGovernanceCallTimelocked)
	if err := _TeeVerification.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVerificationGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeVerification contract.
type TeeVerificationGovernanceInitialisedIterator struct {
	Event *TeeVerificationGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeeVerificationGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationGovernanceInitialised)
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
		it.Event = new(TeeVerificationGovernanceInitialised)
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
func (it *TeeVerificationGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationGovernanceInitialised represents a GovernanceInitialised event raised by the TeeVerification contract.
type TeeVerificationGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeVerification *TeeVerificationFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeVerificationGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeVerificationGovernanceInitialisedIterator{contract: _TeeVerification.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeVerification *TeeVerificationFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeVerificationGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationGovernanceInitialised)
				if err := _TeeVerification.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeeVerification *TeeVerificationFilterer) ParseGovernanceInitialised(log types.Log) (*TeeVerificationGovernanceInitialised, error) {
	event := new(TeeVerificationGovernanceInitialised)
	if err := _TeeVerification.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVerificationSettingsUpdatedIterator is returned from FilterSettingsUpdated and is used to iterate over the raw logs and unpacked data for SettingsUpdated events raised by the TeeVerification contract.
type TeeVerificationSettingsUpdatedIterator struct {
	Event *TeeVerificationSettingsUpdated // Event containing the contract specifics and raw log

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
func (it *TeeVerificationSettingsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationSettingsUpdated)
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
		it.Event = new(TeeVerificationSettingsUpdated)
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
func (it *TeeVerificationSettingsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationSettingsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationSettingsUpdated represents a SettingsUpdated event raised by the TeeVerification contract.
type TeeVerificationSettingsUpdated struct {
	AvailabilityCheckValidityDurationSeconds    uint64
	SigningPolicyValidityDurationInRewardEpochs uint64
	ChallengeValidityDurationSeconds            uint64
	Raw                                         types.Log // Blockchain specific contextual infos
}

// FilterSettingsUpdated is a free log retrieval operation binding the contract event 0x61d24ae6ba3e0c2182628db244b31540bc1b1ffdc9f16fd32ffa52caeb9868b9.
//
// Solidity: event SettingsUpdated(uint64 availabilityCheckValidityDurationSeconds, uint64 signingPolicyValidityDurationInRewardEpochs, uint64 challengeValidityDurationSeconds)
func (_TeeVerification *TeeVerificationFilterer) FilterSettingsUpdated(opts *bind.FilterOpts) (*TeeVerificationSettingsUpdatedIterator, error) {

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "SettingsUpdated")
	if err != nil {
		return nil, err
	}
	return &TeeVerificationSettingsUpdatedIterator{contract: _TeeVerification.contract, event: "SettingsUpdated", logs: logs, sub: sub}, nil
}

// WatchSettingsUpdated is a free log subscription operation binding the contract event 0x61d24ae6ba3e0c2182628db244b31540bc1b1ffdc9f16fd32ffa52caeb9868b9.
//
// Solidity: event SettingsUpdated(uint64 availabilityCheckValidityDurationSeconds, uint64 signingPolicyValidityDurationInRewardEpochs, uint64 challengeValidityDurationSeconds)
func (_TeeVerification *TeeVerificationFilterer) WatchSettingsUpdated(opts *bind.WatchOpts, sink chan<- *TeeVerificationSettingsUpdated) (event.Subscription, error) {

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "SettingsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationSettingsUpdated)
				if err := _TeeVerification.contract.UnpackLog(event, "SettingsUpdated", log); err != nil {
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

// ParseSettingsUpdated is a log parse operation binding the contract event 0x61d24ae6ba3e0c2182628db244b31540bc1b1ffdc9f16fd32ffa52caeb9868b9.
//
// Solidity: event SettingsUpdated(uint64 availabilityCheckValidityDurationSeconds, uint64 signingPolicyValidityDurationInRewardEpochs, uint64 challengeValidityDurationSeconds)
func (_TeeVerification *TeeVerificationFilterer) ParseSettingsUpdated(log types.Log) (*TeeVerificationSettingsUpdated, error) {
	event := new(TeeVerificationSettingsUpdated)
	if err := _TeeVerification.contract.UnpackLog(event, "SettingsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVerificationTeeAttestationRequestedIterator is returned from FilterTeeAttestationRequested and is used to iterate over the raw logs and unpacked data for TeeAttestationRequested events raised by the TeeVerification contract.
type TeeVerificationTeeAttestationRequestedIterator struct {
	Event *TeeVerificationTeeAttestationRequested // Event containing the contract specifics and raw log

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
func (it *TeeVerificationTeeAttestationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationTeeAttestationRequested)
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
		it.Event = new(TeeVerificationTeeAttestationRequested)
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
func (it *TeeVerificationTeeAttestationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationTeeAttestationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationTeeAttestationRequested represents a TeeAttestationRequested event raised by the TeeVerification contract.
type TeeVerificationTeeAttestationRequested struct {
	TeeId     common.Address
	Challenge [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTeeAttestationRequested is a free log retrieval operation binding the contract event 0x88097792b658dddf5ef685fd49ad9d66c9838d549030d9b4bf44c80dd2076db1.
//
// Solidity: event TeeAttestationRequested(address indexed teeId, bytes32 challenge)
func (_TeeVerification *TeeVerificationFilterer) FilterTeeAttestationRequested(opts *bind.FilterOpts, teeId []common.Address) (*TeeVerificationTeeAttestationRequestedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "TeeAttestationRequested", teeIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeVerificationTeeAttestationRequestedIterator{contract: _TeeVerification.contract, event: "TeeAttestationRequested", logs: logs, sub: sub}, nil
}

// WatchTeeAttestationRequested is a free log subscription operation binding the contract event 0x88097792b658dddf5ef685fd49ad9d66c9838d549030d9b4bf44c80dd2076db1.
//
// Solidity: event TeeAttestationRequested(address indexed teeId, bytes32 challenge)
func (_TeeVerification *TeeVerificationFilterer) WatchTeeAttestationRequested(opts *bind.WatchOpts, sink chan<- *TeeVerificationTeeAttestationRequested, teeId []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "TeeAttestationRequested", teeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationTeeAttestationRequested)
				if err := _TeeVerification.contract.UnpackLog(event, "TeeAttestationRequested", log); err != nil {
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

// ParseTeeAttestationRequested is a log parse operation binding the contract event 0x88097792b658dddf5ef685fd49ad9d66c9838d549030d9b4bf44c80dd2076db1.
//
// Solidity: event TeeAttestationRequested(address indexed teeId, bytes32 challenge)
func (_TeeVerification *TeeVerificationFilterer) ParseTeeAttestationRequested(log types.Log) (*TeeVerificationTeeAttestationRequested, error) {
	event := new(TeeVerificationTeeAttestationRequested)
	if err := _TeeVerification.contract.UnpackLog(event, "TeeAttestationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVerificationTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the TeeVerification contract.
type TeeVerificationTeeInstructionsSentIterator struct {
	Event *TeeVerificationTeeInstructionsSent // Event containing the contract specifics and raw log

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
func (it *TeeVerificationTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationTeeInstructionsSent)
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
		it.Event = new(TeeVerificationTeeInstructionsSent)
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
func (it *TeeVerificationTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationTeeInstructionsSent represents a TeeInstructionsSent event raised by the TeeVerification contract.
type TeeVerificationTeeInstructionsSent struct {
	ExtensionId        *big.Int
	InstructionId      [32]byte
	RewardEpochId      uint32
	TeeMachines        []IMachineManagerTeeMachine
	OpType             [32]byte
	OpCommand          [32]byte
	Message            []byte
	Cosigners          []common.Address
	CosignersThreshold uint64
	ClaimBackAddress   common.Address
	Fee                *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTeeInstructionsSent is a free log retrieval operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_TeeVerification *TeeVerificationFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (*TeeVerificationTeeInstructionsSentIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeVerificationTeeInstructionsSentIterator{contract: _TeeVerification.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_TeeVerification *TeeVerificationFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *TeeVerificationTeeInstructionsSent, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationTeeInstructionsSent)
				if err := _TeeVerification.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
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

// ParseTeeInstructionsSent is a log parse operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_TeeVerification *TeeVerificationFilterer) ParseTeeInstructionsSent(log types.Log) (*TeeVerificationTeeInstructionsSent, error) {
	event := new(TeeVerificationTeeInstructionsSent)
	if err := _TeeVerification.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
