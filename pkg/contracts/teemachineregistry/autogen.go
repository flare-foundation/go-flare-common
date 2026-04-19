// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teemachineregistry

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

// IMachineManagerFacetTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerFacetTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
}

// IMachineManagerFacetTeeMachineData is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerFacetTeeMachineData struct {
	ExtensionId  *big.Int
	InitialOwner common.Address
	CodeHash     [32]byte
	Platform     [32]byte
	PublicKey    PublicKey
}

// IMachineManagerFacetTeeMachineWithAttestationData is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerFacetTeeMachineWithAttestationData struct {
	TeeId        common.Address
	InitialTeeId common.Address
	Url          string
	CodeHash     [32]byte
	Platform     [32]byte
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

// PublicKey is an auto generated low-level Go binding around an user-defined struct.
type PublicKey struct {
	X [32]byte
	Y [32]byte
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// TeeMachineRegistryMetaData contains all meta data concerning the TeeMachineRegistry contract.
var TeeMachineRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengeTs\",\"type\":\"uint256\"}],\"name\":\"ChallengeExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestBody\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseDataOrAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningPolicy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeProxyId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeePublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeePublicKeyOrSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUrl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrExpiredAvailabilityCheckOrDisabledVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTeeReplicationContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooMany\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTs\",\"type\":\"uint256\"}],\"name\":\"AvailabilityCheckValidityExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"name\":\"TeeAttestationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManagerFacet.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"TeeMachineRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"TeeMachineSettingsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIMachineManagerFacet.TeeStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"TeeMachineStatusChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getActiveTeeMachines\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_urls\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"getAllActiveTeeMachines\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_urls\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_totalLength\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getExtensionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getInitialSigningPolicyId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getLastStatusChangeTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getRandomTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachine\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIMachineManagerFacet.TeeMachine\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineStatus\",\"outputs\":[{\"internalType\":\"enumIMachineManagerFacet.TeeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineWithAttestationData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIMachineManagerFacet.TeeMachineWithAttestationData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"pauseWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"internalType\":\"structIMachineManagerFacet.TeeMachineData\",\"name\":\"_teeMachineData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"_teeMachineDataSignature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"toProduction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"unban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"updateTeeMachineSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeMachineRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeMachineRegistryMetaData.ABI instead.
var TeeMachineRegistryABI = TeeMachineRegistryMetaData.ABI

// TeeMachineRegistry is an auto generated Go binding around an Ethereum contract.
type TeeMachineRegistry struct {
	TeeMachineRegistryCaller     // Read-only binding to the contract
	TeeMachineRegistryTransactor // Write-only binding to the contract
	TeeMachineRegistryFilterer   // Log filterer for contract events
}

// TeeMachineRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeMachineRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeMachineRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeMachineRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeMachineRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeMachineRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeMachineRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeMachineRegistrySession struct {
	Contract     *TeeMachineRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TeeMachineRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeMachineRegistryCallerSession struct {
	Contract *TeeMachineRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TeeMachineRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeMachineRegistryTransactorSession struct {
	Contract     *TeeMachineRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TeeMachineRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeMachineRegistryRaw struct {
	Contract *TeeMachineRegistry // Generic contract binding to access the raw methods on
}

// TeeMachineRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeMachineRegistryCallerRaw struct {
	Contract *TeeMachineRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TeeMachineRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeMachineRegistryTransactorRaw struct {
	Contract *TeeMachineRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeMachineRegistry creates a new instance of TeeMachineRegistry, bound to a specific deployed contract.
func NewTeeMachineRegistry(address common.Address, backend bind.ContractBackend) (*TeeMachineRegistry, error) {
	contract, err := bindTeeMachineRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistry{TeeMachineRegistryCaller: TeeMachineRegistryCaller{contract: contract}, TeeMachineRegistryTransactor: TeeMachineRegistryTransactor{contract: contract}, TeeMachineRegistryFilterer: TeeMachineRegistryFilterer{contract: contract}}, nil
}

// NewTeeMachineRegistryCaller creates a new read-only instance of TeeMachineRegistry, bound to a specific deployed contract.
func NewTeeMachineRegistryCaller(address common.Address, caller bind.ContractCaller) (*TeeMachineRegistryCaller, error) {
	contract, err := bindTeeMachineRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryCaller{contract: contract}, nil
}

// NewTeeMachineRegistryTransactor creates a new write-only instance of TeeMachineRegistry, bound to a specific deployed contract.
func NewTeeMachineRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeMachineRegistryTransactor, error) {
	contract, err := bindTeeMachineRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryTransactor{contract: contract}, nil
}

// NewTeeMachineRegistryFilterer creates a new log filterer instance of TeeMachineRegistry, bound to a specific deployed contract.
func NewTeeMachineRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeMachineRegistryFilterer, error) {
	contract, err := bindTeeMachineRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryFilterer{contract: contract}, nil
}

// bindTeeMachineRegistry binds a generic wrapper to an already deployed contract.
func bindTeeMachineRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeMachineRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeMachineRegistry *TeeMachineRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeMachineRegistry.Contract.TeeMachineRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeMachineRegistry *TeeMachineRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.TeeMachineRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeMachineRegistry *TeeMachineRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.TeeMachineRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeMachineRegistry *TeeMachineRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeMachineRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeMachineRegistry *TeeMachineRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeMachineRegistry *TeeMachineRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetActiveTeeMachines is a free data retrieval call binding the contract method 0x45f1f0f0.
//
// Solidity: function getActiveTeeMachines(uint256 _extensionId) view returns(address[] _teeIds, string[] _urls)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetActiveTeeMachines(opts *bind.CallOpts, _extensionId *big.Int) (struct {
	TeeIds []common.Address
	Urls   []string
}, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getActiveTeeMachines", _extensionId)

	outstruct := new(struct {
		TeeIds []common.Address
		Urls   []string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TeeIds = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Urls = *abi.ConvertType(out[1], new([]string)).(*[]string)

	return *outstruct, err

}

// GetActiveTeeMachines is a free data retrieval call binding the contract method 0x45f1f0f0.
//
// Solidity: function getActiveTeeMachines(uint256 _extensionId) view returns(address[] _teeIds, string[] _urls)
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetActiveTeeMachines(_extensionId *big.Int) (struct {
	TeeIds []common.Address
	Urls   []string
}, error) {
	return _TeeMachineRegistry.Contract.GetActiveTeeMachines(&_TeeMachineRegistry.CallOpts, _extensionId)
}

// GetActiveTeeMachines is a free data retrieval call binding the contract method 0x45f1f0f0.
//
// Solidity: function getActiveTeeMachines(uint256 _extensionId) view returns(address[] _teeIds, string[] _urls)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetActiveTeeMachines(_extensionId *big.Int) (struct {
	TeeIds []common.Address
	Urls   []string
}, error) {
	return _TeeMachineRegistry.Contract.GetActiveTeeMachines(&_TeeMachineRegistry.CallOpts, _extensionId)
}

// GetAllActiveTeeMachines is a free data retrieval call binding the contract method 0xa5578af4.
//
// Solidity: function getAllActiveTeeMachines(uint256 _start, uint256 _end) view returns(address[] _teeIds, string[] _urls, uint256 _totalLength)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetAllActiveTeeMachines(opts *bind.CallOpts, _start *big.Int, _end *big.Int) (struct {
	TeeIds      []common.Address
	Urls        []string
	TotalLength *big.Int
}, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getAllActiveTeeMachines", _start, _end)

	outstruct := new(struct {
		TeeIds      []common.Address
		Urls        []string
		TotalLength *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TeeIds = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Urls = *abi.ConvertType(out[1], new([]string)).(*[]string)
	outstruct.TotalLength = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAllActiveTeeMachines is a free data retrieval call binding the contract method 0xa5578af4.
//
// Solidity: function getAllActiveTeeMachines(uint256 _start, uint256 _end) view returns(address[] _teeIds, string[] _urls, uint256 _totalLength)
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetAllActiveTeeMachines(_start *big.Int, _end *big.Int) (struct {
	TeeIds      []common.Address
	Urls        []string
	TotalLength *big.Int
}, error) {
	return _TeeMachineRegistry.Contract.GetAllActiveTeeMachines(&_TeeMachineRegistry.CallOpts, _start, _end)
}

// GetAllActiveTeeMachines is a free data retrieval call binding the contract method 0xa5578af4.
//
// Solidity: function getAllActiveTeeMachines(uint256 _start, uint256 _end) view returns(address[] _teeIds, string[] _urls, uint256 _totalLength)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetAllActiveTeeMachines(_start *big.Int, _end *big.Int) (struct {
	TeeIds      []common.Address
	Urls        []string
	TotalLength *big.Int
}, error) {
	return _TeeMachineRegistry.Contract.GetAllActiveTeeMachines(&_TeeMachineRegistry.CallOpts, _start, _end)
}

// GetExtensionId is a free data retrieval call binding the contract method 0xaa5bb892.
//
// Solidity: function getExtensionId(address _teeId) view returns(uint256)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetExtensionId(opts *bind.CallOpts, _teeId common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getExtensionId", _teeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExtensionId is a free data retrieval call binding the contract method 0xaa5bb892.
//
// Solidity: function getExtensionId(address _teeId) view returns(uint256)
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetExtensionId(_teeId common.Address) (*big.Int, error) {
	return _TeeMachineRegistry.Contract.GetExtensionId(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetExtensionId is a free data retrieval call binding the contract method 0xaa5bb892.
//
// Solidity: function getExtensionId(address _teeId) view returns(uint256)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetExtensionId(_teeId common.Address) (*big.Int, error) {
	return _TeeMachineRegistry.Contract.GetExtensionId(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetInitialSigningPolicyId is a free data retrieval call binding the contract method 0xb2e51edd.
//
// Solidity: function getInitialSigningPolicyId(address _teeId) view returns(uint32)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetInitialSigningPolicyId(opts *bind.CallOpts, _teeId common.Address) (uint32, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getInitialSigningPolicyId", _teeId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetInitialSigningPolicyId is a free data retrieval call binding the contract method 0xb2e51edd.
//
// Solidity: function getInitialSigningPolicyId(address _teeId) view returns(uint32)
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetInitialSigningPolicyId(_teeId common.Address) (uint32, error) {
	return _TeeMachineRegistry.Contract.GetInitialSigningPolicyId(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetInitialSigningPolicyId is a free data retrieval call binding the contract method 0xb2e51edd.
//
// Solidity: function getInitialSigningPolicyId(address _teeId) view returns(uint32)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetInitialSigningPolicyId(_teeId common.Address) (uint32, error) {
	return _TeeMachineRegistry.Contract.GetInitialSigningPolicyId(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetLastStatusChangeTs is a free data retrieval call binding the contract method 0x1e17eb5a.
//
// Solidity: function getLastStatusChangeTs(address _teeId) view returns(uint256)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetLastStatusChangeTs(opts *bind.CallOpts, _teeId common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getLastStatusChangeTs", _teeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastStatusChangeTs is a free data retrieval call binding the contract method 0x1e17eb5a.
//
// Solidity: function getLastStatusChangeTs(address _teeId) view returns(uint256)
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetLastStatusChangeTs(_teeId common.Address) (*big.Int, error) {
	return _TeeMachineRegistry.Contract.GetLastStatusChangeTs(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetLastStatusChangeTs is a free data retrieval call binding the contract method 0x1e17eb5a.
//
// Solidity: function getLastStatusChangeTs(address _teeId) view returns(uint256)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetLastStatusChangeTs(_teeId common.Address) (*big.Int, error) {
	return _TeeMachineRegistry.Contract.GetLastStatusChangeTs(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address _teeId) view returns((bytes32,bytes32))
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetPublicKey(opts *bind.CallOpts, _teeId common.Address) (PublicKey, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getPublicKey", _teeId)

	if err != nil {
		return *new(PublicKey), err
	}

	out0 := *abi.ConvertType(out[0], new(PublicKey)).(*PublicKey)

	return out0, err

}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address _teeId) view returns((bytes32,bytes32))
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetPublicKey(_teeId common.Address) (PublicKey, error) {
	return _TeeMachineRegistry.Contract.GetPublicKey(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address _teeId) view returns((bytes32,bytes32))
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetPublicKey(_teeId common.Address) (PublicKey, error) {
	return _TeeMachineRegistry.Contract.GetPublicKey(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetRandomTeeIds is a free data retrieval call binding the contract method 0xfeeabcbf.
//
// Solidity: function getRandomTeeIds(uint256 _extensionId, uint256 _count) view returns(address[] _teeIds)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetRandomTeeIds(opts *bind.CallOpts, _extensionId *big.Int, _count *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getRandomTeeIds", _extensionId, _count)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRandomTeeIds is a free data retrieval call binding the contract method 0xfeeabcbf.
//
// Solidity: function getRandomTeeIds(uint256 _extensionId, uint256 _count) view returns(address[] _teeIds)
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetRandomTeeIds(_extensionId *big.Int, _count *big.Int) ([]common.Address, error) {
	return _TeeMachineRegistry.Contract.GetRandomTeeIds(&_TeeMachineRegistry.CallOpts, _extensionId, _count)
}

// GetRandomTeeIds is a free data retrieval call binding the contract method 0xfeeabcbf.
//
// Solidity: function getRandomTeeIds(uint256 _extensionId, uint256 _count) view returns(address[] _teeIds)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetRandomTeeIds(_extensionId *big.Int, _count *big.Int) ([]common.Address, error) {
	return _TeeMachineRegistry.Contract.GetRandomTeeIds(&_TeeMachineRegistry.CallOpts, _extensionId, _count)
}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetTeeMachine(opts *bind.CallOpts, _teeId common.Address) (IMachineManagerFacetTeeMachine, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getTeeMachine", _teeId)

	if err != nil {
		return *new(IMachineManagerFacetTeeMachine), err
	}

	out0 := *abi.ConvertType(out[0], new(IMachineManagerFacetTeeMachine)).(*IMachineManagerFacetTeeMachine)

	return out0, err

}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetTeeMachine(_teeId common.Address) (IMachineManagerFacetTeeMachine, error) {
	return _TeeMachineRegistry.Contract.GetTeeMachine(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetTeeMachine(_teeId common.Address) (IMachineManagerFacetTeeMachine, error) {
	return _TeeMachineRegistry.Contract.GetTeeMachine(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetTeeMachineOwner is a free data retrieval call binding the contract method 0xb7e773e9.
//
// Solidity: function getTeeMachineOwner(address _teeId) view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetTeeMachineOwner(opts *bind.CallOpts, _teeId common.Address) (common.Address, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getTeeMachineOwner", _teeId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeeMachineOwner is a free data retrieval call binding the contract method 0xb7e773e9.
//
// Solidity: function getTeeMachineOwner(address _teeId) view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetTeeMachineOwner(_teeId common.Address) (common.Address, error) {
	return _TeeMachineRegistry.Contract.GetTeeMachineOwner(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetTeeMachineOwner is a free data retrieval call binding the contract method 0xb7e773e9.
//
// Solidity: function getTeeMachineOwner(address _teeId) view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetTeeMachineOwner(_teeId common.Address) (common.Address, error) {
	return _TeeMachineRegistry.Contract.GetTeeMachineOwner(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetTeeMachineStatus(opts *bind.CallOpts, _teeId common.Address) (uint8, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getTeeMachineStatus", _teeId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetTeeMachineStatus(_teeId common.Address) (uint8, error) {
	return _TeeMachineRegistry.Contract.GetTeeMachineStatus(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetTeeMachineStatus(_teeId common.Address) (uint8, error) {
	return _TeeMachineRegistry.Contract.GetTeeMachineStatus(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetTeeMachineWithAttestationData(opts *bind.CallOpts, _teeId common.Address) (IMachineManagerFacetTeeMachineWithAttestationData, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getTeeMachineWithAttestationData", _teeId)

	if err != nil {
		return *new(IMachineManagerFacetTeeMachineWithAttestationData), err
	}

	out0 := *abi.ConvertType(out[0], new(IMachineManagerFacetTeeMachineWithAttestationData)).(*IMachineManagerFacetTeeMachineWithAttestationData)

	return out0, err

}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetTeeMachineWithAttestationData(_teeId common.Address) (IMachineManagerFacetTeeMachineWithAttestationData, error) {
	return _TeeMachineRegistry.Contract.GetTeeMachineWithAttestationData(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetTeeMachineWithAttestationData(_teeId common.Address) (IMachineManagerFacetTeeMachineWithAttestationData, error) {
	return _TeeMachineRegistry.Contract.GetTeeMachineWithAttestationData(&_TeeMachineRegistry.CallOpts, _teeId)
}

// Ban is a paid mutator transaction binding the contract method 0x97c3ccd8.
//
// Solidity: function ban(address _teeId) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) Ban(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "ban", _teeId)
}

// Ban is a paid mutator transaction binding the contract method 0x97c3ccd8.
//
// Solidity: function ban(address _teeId) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) Ban(_teeId common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Ban(&_TeeMachineRegistry.TransactOpts, _teeId)
}

// Ban is a paid mutator transaction binding the contract method 0x97c3ccd8.
//
// Solidity: function ban(address _teeId) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) Ban(_teeId common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Ban(&_TeeMachineRegistry.TransactOpts, _teeId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x80f2abe3.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) ConfirmOwnership(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "confirmOwnership", _teeId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x80f2abe3.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) ConfirmOwnership(_teeId common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.ConfirmOwnership(&_TeeMachineRegistry.TransactOpts, _teeId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x80f2abe3.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) ConfirmOwnership(_teeId common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.ConfirmOwnership(&_TeeMachineRegistry.TransactOpts, _teeId)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _teeId) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) Pause(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "pause", _teeId)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _teeId) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) Pause(_teeId common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Pause(&_TeeMachineRegistry.TransactOpts, _teeId)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _teeId) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) Pause(_teeId common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Pause(&_TeeMachineRegistry.TransactOpts, _teeId)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0xd0824eac.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) PauseWithProof(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "pauseWithProof", _proof)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0xd0824eac.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) PauseWithProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.PauseWithProof(&_TeeMachineRegistry.TransactOpts, _proof)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0xd0824eac.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) PauseWithProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.PauseWithProof(&_TeeMachineRegistry.TransactOpts, _proof)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x38195139.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) ProposeNewOwner(opts *bind.TransactOpts, _teeId common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "proposeNewOwner", _teeId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x38195139.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) ProposeNewOwner(_teeId common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.ProposeNewOwner(&_TeeMachineRegistry.TransactOpts, _teeId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x38195139.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) ProposeNewOwner(_teeId common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.ProposeNewOwner(&_TeeMachineRegistry.TransactOpts, _teeId, _newOwner)
}

// Register is a paid mutator transaction binding the contract method 0x95ef52e0.
//
// Solidity: function register((uint256,address,bytes32,bytes32,(bytes32,bytes32)) _teeMachineData, (uint8,bytes32,bytes32) _teeMachineDataSignature, address _teeProxyId, string _url, address _claimBackAddress) payable returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) Register(opts *bind.TransactOpts, _teeMachineData IMachineManagerFacetTeeMachineData, _teeMachineDataSignature Signature, _teeProxyId common.Address, _url string, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "register", _teeMachineData, _teeMachineDataSignature, _teeProxyId, _url, _claimBackAddress)
}

// Register is a paid mutator transaction binding the contract method 0x95ef52e0.
//
// Solidity: function register((uint256,address,bytes32,bytes32,(bytes32,bytes32)) _teeMachineData, (uint8,bytes32,bytes32) _teeMachineDataSignature, address _teeProxyId, string _url, address _claimBackAddress) payable returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) Register(_teeMachineData IMachineManagerFacetTeeMachineData, _teeMachineDataSignature Signature, _teeProxyId common.Address, _url string, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Register(&_TeeMachineRegistry.TransactOpts, _teeMachineData, _teeMachineDataSignature, _teeProxyId, _url, _claimBackAddress)
}

// Register is a paid mutator transaction binding the contract method 0x95ef52e0.
//
// Solidity: function register((uint256,address,bytes32,bytes32,(bytes32,bytes32)) _teeMachineData, (uint8,bytes32,bytes32) _teeMachineDataSignature, address _teeProxyId, string _url, address _claimBackAddress) payable returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) Register(_teeMachineData IMachineManagerFacetTeeMachineData, _teeMachineDataSignature Signature, _teeProxyId common.Address, _url string, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Register(&_TeeMachineRegistry.TransactOpts, _teeMachineData, _teeMachineDataSignature, _teeProxyId, _url, _claimBackAddress)
}

// ToProduction is a paid mutator transaction binding the contract method 0x4d40e8fb.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) ToProduction(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "toProduction", _proof)
}

// ToProduction is a paid mutator transaction binding the contract method 0x4d40e8fb.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) ToProduction(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.ToProduction(&_TeeMachineRegistry.TransactOpts, _proof)
}

// ToProduction is a paid mutator transaction binding the contract method 0x4d40e8fb.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) ToProduction(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.ToProduction(&_TeeMachineRegistry.TransactOpts, _proof)
}

// Unban is a paid mutator transaction binding the contract method 0xb9f14557.
//
// Solidity: function unban(address _teeId) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) Unban(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "unban", _teeId)
}

// Unban is a paid mutator transaction binding the contract method 0xb9f14557.
//
// Solidity: function unban(address _teeId) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) Unban(_teeId common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Unban(&_TeeMachineRegistry.TransactOpts, _teeId)
}

// Unban is a paid mutator transaction binding the contract method 0xb9f14557.
//
// Solidity: function unban(address _teeId) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) Unban(_teeId common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Unban(&_TeeMachineRegistry.TransactOpts, _teeId)
}

// UpdateTeeMachineSettings is a paid mutator transaction binding the contract method 0x06ed5da4.
//
// Solidity: function updateTeeMachineSettings(address _teeId, address _teeProxyId, string _url) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) UpdateTeeMachineSettings(opts *bind.TransactOpts, _teeId common.Address, _teeProxyId common.Address, _url string) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "updateTeeMachineSettings", _teeId, _teeProxyId, _url)
}

// UpdateTeeMachineSettings is a paid mutator transaction binding the contract method 0x06ed5da4.
//
// Solidity: function updateTeeMachineSettings(address _teeId, address _teeProxyId, string _url) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) UpdateTeeMachineSettings(_teeId common.Address, _teeProxyId common.Address, _url string) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.UpdateTeeMachineSettings(&_TeeMachineRegistry.TransactOpts, _teeId, _teeProxyId, _url)
}

// UpdateTeeMachineSettings is a paid mutator transaction binding the contract method 0x06ed5da4.
//
// Solidity: function updateTeeMachineSettings(address _teeId, address _teeProxyId, string _url) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) UpdateTeeMachineSettings(_teeId common.Address, _teeProxyId common.Address, _url string) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.UpdateTeeMachineSettings(&_TeeMachineRegistry.TransactOpts, _teeId, _teeProxyId, _url)
}

// TeeMachineRegistryAvailabilityCheckValidityExtendedIterator is returned from FilterAvailabilityCheckValidityExtended and is used to iterate over the raw logs and unpacked data for AvailabilityCheckValidityExtended events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryAvailabilityCheckValidityExtendedIterator struct {
	Event *TeeMachineRegistryAvailabilityCheckValidityExtended // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryAvailabilityCheckValidityExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryAvailabilityCheckValidityExtended)
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
		it.Event = new(TeeMachineRegistryAvailabilityCheckValidityExtended)
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
func (it *TeeMachineRegistryAvailabilityCheckValidityExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryAvailabilityCheckValidityExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryAvailabilityCheckValidityExtended represents a AvailabilityCheckValidityExtended event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryAvailabilityCheckValidityExtended struct {
	TeeId common.Address
	Owner common.Address
	EndTs *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAvailabilityCheckValidityExtended is a free log retrieval operation binding the contract event 0xbcea1901d8026bcd3401ce1c443ba0a3e5e16680271d1dcc84cfdd36a0ca44b3.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterAvailabilityCheckValidityExtended(opts *bind.FilterOpts, teeId []common.Address, owner []common.Address) (*TeeMachineRegistryAvailabilityCheckValidityExtendedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "AvailabilityCheckValidityExtended", teeIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryAvailabilityCheckValidityExtendedIterator{contract: _TeeMachineRegistry.contract, event: "AvailabilityCheckValidityExtended", logs: logs, sub: sub}, nil
}

// WatchAvailabilityCheckValidityExtended is a free log subscription operation binding the contract event 0xbcea1901d8026bcd3401ce1c443ba0a3e5e16680271d1dcc84cfdd36a0ca44b3.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchAvailabilityCheckValidityExtended(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryAvailabilityCheckValidityExtended, teeId []common.Address, owner []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "AvailabilityCheckValidityExtended", teeIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryAvailabilityCheckValidityExtended)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "AvailabilityCheckValidityExtended", log); err != nil {
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
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseAvailabilityCheckValidityExtended(log types.Log) (*TeeMachineRegistryAvailabilityCheckValidityExtended, error) {
	event := new(TeeMachineRegistryAvailabilityCheckValidityExtended)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "AvailabilityCheckValidityExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeMachineRegistryNewOwnerConfirmedIterator is returned from FilterNewOwnerConfirmed and is used to iterate over the raw logs and unpacked data for NewOwnerConfirmed events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryNewOwnerConfirmedIterator struct {
	Event *TeeMachineRegistryNewOwnerConfirmed // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryNewOwnerConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryNewOwnerConfirmed)
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
		it.Event = new(TeeMachineRegistryNewOwnerConfirmed)
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
func (it *TeeMachineRegistryNewOwnerConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryNewOwnerConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryNewOwnerConfirmed represents a NewOwnerConfirmed event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryNewOwnerConfirmed struct {
	TeeId    common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerConfirmed is a free log retrieval operation binding the contract event 0xc813c793f67554890ac6f92415119ac4a7535162dc38fb1120036cbd62069fbf.
//
// Solidity: event NewOwnerConfirmed(address indexed teeId, address indexed newOwner)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterNewOwnerConfirmed(opts *bind.FilterOpts, teeId []common.Address, newOwner []common.Address) (*TeeMachineRegistryNewOwnerConfirmedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "NewOwnerConfirmed", teeIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryNewOwnerConfirmedIterator{contract: _TeeMachineRegistry.contract, event: "NewOwnerConfirmed", logs: logs, sub: sub}, nil
}

// WatchNewOwnerConfirmed is a free log subscription operation binding the contract event 0xc813c793f67554890ac6f92415119ac4a7535162dc38fb1120036cbd62069fbf.
//
// Solidity: event NewOwnerConfirmed(address indexed teeId, address indexed newOwner)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchNewOwnerConfirmed(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryNewOwnerConfirmed, teeId []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "NewOwnerConfirmed", teeIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryNewOwnerConfirmed)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "NewOwnerConfirmed", log); err != nil {
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

// ParseNewOwnerConfirmed is a log parse operation binding the contract event 0xc813c793f67554890ac6f92415119ac4a7535162dc38fb1120036cbd62069fbf.
//
// Solidity: event NewOwnerConfirmed(address indexed teeId, address indexed newOwner)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseNewOwnerConfirmed(log types.Log) (*TeeMachineRegistryNewOwnerConfirmed, error) {
	event := new(TeeMachineRegistryNewOwnerConfirmed)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "NewOwnerConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeMachineRegistryNewOwnerProposedIterator is returned from FilterNewOwnerProposed and is used to iterate over the raw logs and unpacked data for NewOwnerProposed events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryNewOwnerProposedIterator struct {
	Event *TeeMachineRegistryNewOwnerProposed // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryNewOwnerProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryNewOwnerProposed)
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
		it.Event = new(TeeMachineRegistryNewOwnerProposed)
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
func (it *TeeMachineRegistryNewOwnerProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryNewOwnerProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryNewOwnerProposed represents a NewOwnerProposed event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryNewOwnerProposed struct {
	TeeId    common.Address
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerProposed is a free log retrieval operation binding the contract event 0x64420d4a41c6ed4de2bccbf33192eea18e576c5b23c79c3a722d4e9534c2e8d8.
//
// Solidity: event NewOwnerProposed(address indexed teeId, address indexed oldOwner, address indexed newOwner)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterNewOwnerProposed(opts *bind.FilterOpts, teeId []common.Address, oldOwner []common.Address, newOwner []common.Address) (*TeeMachineRegistryNewOwnerProposedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "NewOwnerProposed", teeIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryNewOwnerProposedIterator{contract: _TeeMachineRegistry.contract, event: "NewOwnerProposed", logs: logs, sub: sub}, nil
}

// WatchNewOwnerProposed is a free log subscription operation binding the contract event 0x64420d4a41c6ed4de2bccbf33192eea18e576c5b23c79c3a722d4e9534c2e8d8.
//
// Solidity: event NewOwnerProposed(address indexed teeId, address indexed oldOwner, address indexed newOwner)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchNewOwnerProposed(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryNewOwnerProposed, teeId []common.Address, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "NewOwnerProposed", teeIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryNewOwnerProposed)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
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

// ParseNewOwnerProposed is a log parse operation binding the contract event 0x64420d4a41c6ed4de2bccbf33192eea18e576c5b23c79c3a722d4e9534c2e8d8.
//
// Solidity: event NewOwnerProposed(address indexed teeId, address indexed oldOwner, address indexed newOwner)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseNewOwnerProposed(log types.Log) (*TeeMachineRegistryNewOwnerProposed, error) {
	event := new(TeeMachineRegistryNewOwnerProposed)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeMachineRegistryTeeAttestationRequestedIterator is returned from FilterTeeAttestationRequested and is used to iterate over the raw logs and unpacked data for TeeAttestationRequested events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTeeAttestationRequestedIterator struct {
	Event *TeeMachineRegistryTeeAttestationRequested // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryTeeAttestationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryTeeAttestationRequested)
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
		it.Event = new(TeeMachineRegistryTeeAttestationRequested)
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
func (it *TeeMachineRegistryTeeAttestationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryTeeAttestationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryTeeAttestationRequested represents a TeeAttestationRequested event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTeeAttestationRequested struct {
	TeeId     common.Address
	Challenge [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTeeAttestationRequested is a free log retrieval operation binding the contract event 0x88097792b658dddf5ef685fd49ad9d66c9838d549030d9b4bf44c80dd2076db1.
//
// Solidity: event TeeAttestationRequested(address indexed teeId, bytes32 challenge)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterTeeAttestationRequested(opts *bind.FilterOpts, teeId []common.Address) (*TeeMachineRegistryTeeAttestationRequestedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "TeeAttestationRequested", teeIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryTeeAttestationRequestedIterator{contract: _TeeMachineRegistry.contract, event: "TeeAttestationRequested", logs: logs, sub: sub}, nil
}

// WatchTeeAttestationRequested is a free log subscription operation binding the contract event 0x88097792b658dddf5ef685fd49ad9d66c9838d549030d9b4bf44c80dd2076db1.
//
// Solidity: event TeeAttestationRequested(address indexed teeId, bytes32 challenge)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchTeeAttestationRequested(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryTeeAttestationRequested, teeId []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "TeeAttestationRequested", teeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryTeeAttestationRequested)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "TeeAttestationRequested", log); err != nil {
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
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseTeeAttestationRequested(log types.Log) (*TeeMachineRegistryTeeAttestationRequested, error) {
	event := new(TeeMachineRegistryTeeAttestationRequested)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "TeeAttestationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeMachineRegistryTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTeeInstructionsSentIterator struct {
	Event *TeeMachineRegistryTeeInstructionsSent // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryTeeInstructionsSent)
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
		it.Event = new(TeeMachineRegistryTeeInstructionsSent)
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
func (it *TeeMachineRegistryTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryTeeInstructionsSent represents a TeeInstructionsSent event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTeeInstructionsSent struct {
	ExtensionId        *big.Int
	InstructionId      [32]byte
	RewardEpochId      uint32
	TeeMachines        []IMachineManagerFacetTeeMachine
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
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (*TeeMachineRegistryTeeInstructionsSentIterator, error) {

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

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryTeeInstructionsSentIterator{contract: _TeeMachineRegistry.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryTeeInstructionsSent, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

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

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryTeeInstructionsSent)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
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
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseTeeInstructionsSent(log types.Log) (*TeeMachineRegistryTeeInstructionsSent, error) {
	event := new(TeeMachineRegistryTeeInstructionsSent)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeMachineRegistryTeeMachineRegisteredIterator is returned from FilterTeeMachineRegistered and is used to iterate over the raw logs and unpacked data for TeeMachineRegistered events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTeeMachineRegisteredIterator struct {
	Event *TeeMachineRegistryTeeMachineRegistered // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryTeeMachineRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryTeeMachineRegistered)
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
		it.Event = new(TeeMachineRegistryTeeMachineRegistered)
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
func (it *TeeMachineRegistryTeeMachineRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryTeeMachineRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryTeeMachineRegistered represents a TeeMachineRegistered event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTeeMachineRegistered struct {
	TeeId       common.Address
	TeeProxyId  common.Address
	Owner       common.Address
	ExtensionId *big.Int
	Url         string
	CodeHash    [32]byte
	Platform    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTeeMachineRegistered is a free log retrieval operation binding the contract event 0x4a4f6c10f9306f09cfbee5108b4cc42e71d5b6cd887eb42a98c483cb5680164a.
//
// Solidity: event TeeMachineRegistered(address indexed teeId, address indexed teeProxyId, address indexed owner, uint256 extensionId, string url, bytes32 codeHash, bytes32 platform)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterTeeMachineRegistered(opts *bind.FilterOpts, teeId []common.Address, teeProxyId []common.Address, owner []common.Address) (*TeeMachineRegistryTeeMachineRegisteredIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var teeProxyIdRule []interface{}
	for _, teeProxyIdItem := range teeProxyId {
		teeProxyIdRule = append(teeProxyIdRule, teeProxyIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "TeeMachineRegistered", teeIdRule, teeProxyIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryTeeMachineRegisteredIterator{contract: _TeeMachineRegistry.contract, event: "TeeMachineRegistered", logs: logs, sub: sub}, nil
}

// WatchTeeMachineRegistered is a free log subscription operation binding the contract event 0x4a4f6c10f9306f09cfbee5108b4cc42e71d5b6cd887eb42a98c483cb5680164a.
//
// Solidity: event TeeMachineRegistered(address indexed teeId, address indexed teeProxyId, address indexed owner, uint256 extensionId, string url, bytes32 codeHash, bytes32 platform)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchTeeMachineRegistered(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryTeeMachineRegistered, teeId []common.Address, teeProxyId []common.Address, owner []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var teeProxyIdRule []interface{}
	for _, teeProxyIdItem := range teeProxyId {
		teeProxyIdRule = append(teeProxyIdRule, teeProxyIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "TeeMachineRegistered", teeIdRule, teeProxyIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryTeeMachineRegistered)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "TeeMachineRegistered", log); err != nil {
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

// ParseTeeMachineRegistered is a log parse operation binding the contract event 0x4a4f6c10f9306f09cfbee5108b4cc42e71d5b6cd887eb42a98c483cb5680164a.
//
// Solidity: event TeeMachineRegistered(address indexed teeId, address indexed teeProxyId, address indexed owner, uint256 extensionId, string url, bytes32 codeHash, bytes32 platform)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseTeeMachineRegistered(log types.Log) (*TeeMachineRegistryTeeMachineRegistered, error) {
	event := new(TeeMachineRegistryTeeMachineRegistered)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "TeeMachineRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeMachineRegistryTeeMachineSettingsUpdatedIterator is returned from FilterTeeMachineSettingsUpdated and is used to iterate over the raw logs and unpacked data for TeeMachineSettingsUpdated events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTeeMachineSettingsUpdatedIterator struct {
	Event *TeeMachineRegistryTeeMachineSettingsUpdated // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryTeeMachineSettingsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryTeeMachineSettingsUpdated)
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
		it.Event = new(TeeMachineRegistryTeeMachineSettingsUpdated)
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
func (it *TeeMachineRegistryTeeMachineSettingsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryTeeMachineSettingsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryTeeMachineSettingsUpdated represents a TeeMachineSettingsUpdated event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTeeMachineSettingsUpdated struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeeMachineSettingsUpdated is a free log retrieval operation binding the contract event 0x2ee0145d87c2e0db71dfa482a5d5fbec4a34c66dba75742b647608507c11bb46.
//
// Solidity: event TeeMachineSettingsUpdated(address indexed teeId, address indexed teeProxyId, string url)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterTeeMachineSettingsUpdated(opts *bind.FilterOpts, teeId []common.Address, teeProxyId []common.Address) (*TeeMachineRegistryTeeMachineSettingsUpdatedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var teeProxyIdRule []interface{}
	for _, teeProxyIdItem := range teeProxyId {
		teeProxyIdRule = append(teeProxyIdRule, teeProxyIdItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "TeeMachineSettingsUpdated", teeIdRule, teeProxyIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryTeeMachineSettingsUpdatedIterator{contract: _TeeMachineRegistry.contract, event: "TeeMachineSettingsUpdated", logs: logs, sub: sub}, nil
}

// WatchTeeMachineSettingsUpdated is a free log subscription operation binding the contract event 0x2ee0145d87c2e0db71dfa482a5d5fbec4a34c66dba75742b647608507c11bb46.
//
// Solidity: event TeeMachineSettingsUpdated(address indexed teeId, address indexed teeProxyId, string url)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchTeeMachineSettingsUpdated(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryTeeMachineSettingsUpdated, teeId []common.Address, teeProxyId []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var teeProxyIdRule []interface{}
	for _, teeProxyIdItem := range teeProxyId {
		teeProxyIdRule = append(teeProxyIdRule, teeProxyIdItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "TeeMachineSettingsUpdated", teeIdRule, teeProxyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryTeeMachineSettingsUpdated)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "TeeMachineSettingsUpdated", log); err != nil {
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

// ParseTeeMachineSettingsUpdated is a log parse operation binding the contract event 0x2ee0145d87c2e0db71dfa482a5d5fbec4a34c66dba75742b647608507c11bb46.
//
// Solidity: event TeeMachineSettingsUpdated(address indexed teeId, address indexed teeProxyId, string url)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseTeeMachineSettingsUpdated(log types.Log) (*TeeMachineRegistryTeeMachineSettingsUpdated, error) {
	event := new(TeeMachineRegistryTeeMachineSettingsUpdated)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "TeeMachineSettingsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeMachineRegistryTeeMachineStatusChangedIterator is returned from FilterTeeMachineStatusChanged and is used to iterate over the raw logs and unpacked data for TeeMachineStatusChanged events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTeeMachineStatusChangedIterator struct {
	Event *TeeMachineRegistryTeeMachineStatusChanged // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryTeeMachineStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryTeeMachineStatusChanged)
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
		it.Event = new(TeeMachineRegistryTeeMachineStatusChanged)
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
func (it *TeeMachineRegistryTeeMachineStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryTeeMachineStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryTeeMachineStatusChanged represents a TeeMachineStatusChanged event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTeeMachineStatusChanged struct {
	TeeId     common.Address
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTeeMachineStatusChanged is a free log retrieval operation binding the contract event 0x15cd9c7dcca8fe809f6f492069627d36efb27df330c4e6e1c0840751dba73ab1.
//
// Solidity: event TeeMachineStatusChanged(address indexed teeId, uint8 indexed newStatus)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterTeeMachineStatusChanged(opts *bind.FilterOpts, teeId []common.Address, newStatus []uint8) (*TeeMachineRegistryTeeMachineStatusChangedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "TeeMachineStatusChanged", teeIdRule, newStatusRule)
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryTeeMachineStatusChangedIterator{contract: _TeeMachineRegistry.contract, event: "TeeMachineStatusChanged", logs: logs, sub: sub}, nil
}

// WatchTeeMachineStatusChanged is a free log subscription operation binding the contract event 0x15cd9c7dcca8fe809f6f492069627d36efb27df330c4e6e1c0840751dba73ab1.
//
// Solidity: event TeeMachineStatusChanged(address indexed teeId, uint8 indexed newStatus)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchTeeMachineStatusChanged(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryTeeMachineStatusChanged, teeId []common.Address, newStatus []uint8) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "TeeMachineStatusChanged", teeIdRule, newStatusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryTeeMachineStatusChanged)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "TeeMachineStatusChanged", log); err != nil {
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

// ParseTeeMachineStatusChanged is a log parse operation binding the contract event 0x15cd9c7dcca8fe809f6f492069627d36efb27df330c4e6e1c0840751dba73ab1.
//
// Solidity: event TeeMachineStatusChanged(address indexed teeId, uint8 indexed newStatus)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseTeeMachineStatusChanged(log types.Log) (*TeeMachineRegistryTeeMachineStatusChanged, error) {
	event := new(TeeMachineRegistryTeeMachineStatusChanged)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "TeeMachineStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
