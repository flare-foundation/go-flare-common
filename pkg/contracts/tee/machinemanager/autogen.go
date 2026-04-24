// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package machinemanager

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

// IMachineManagerTeeMachineData is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachineData struct {
	ExtensionId  *big.Int
	InitialOwner common.Address
	CodeHash     [32]byte
	Platform     [32]byte
	PublicKey    PublicKey
}

// IMachineManagerTeeMachineWithAttestationData is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachineWithAttestationData struct {
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

// MachineManagerMetaData contains all meta data concerning the MachineManager contract.
var MachineManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengeTs\",\"type\":\"uint256\"}],\"name\":\"ChallengeExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestBody\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseDataOrAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningPolicy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeProxyId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeePublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeePublicKeyOrSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUrl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrExpiredAvailabilityCheckOrDisabledVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTeeReplicationContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooMany\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTs\",\"type\":\"uint256\"}],\"name\":\"AvailabilityCheckValidityExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"name\":\"TeeAttestationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"TeeMachineRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"TeeMachineSettingsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIMachineManager.TeeStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"TeeMachineStatusChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getActiveTeeMachines\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_urls\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"getAllActiveTeeMachines\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_urls\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_totalLength\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getExtensionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getInitialSigningPolicyId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getLastStatusChangeTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getRandomTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachine\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIMachineManager.TeeMachine\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineStatus\",\"outputs\":[{\"internalType\":\"enumIMachineManager.TeeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineWithAttestationData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIMachineManager.TeeMachineWithAttestationData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"pauseWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"internalType\":\"structIMachineManager.TeeMachineData\",\"name\":\"_teeMachineData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"_teeMachineDataSignature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"toProduction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"unban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"updateTeeMachineSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MachineManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MachineManagerMetaData.ABI instead.
var MachineManagerABI = MachineManagerMetaData.ABI

// MachineManager is an auto generated Go binding around an Ethereum contract.
type MachineManager struct {
	MachineManagerCaller     // Read-only binding to the contract
	MachineManagerTransactor // Write-only binding to the contract
	MachineManagerFilterer   // Log filterer for contract events
}

// MachineManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MachineManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MachineManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MachineManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MachineManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MachineManagerSession struct {
	Contract     *MachineManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MachineManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MachineManagerCallerSession struct {
	Contract *MachineManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MachineManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MachineManagerTransactorSession struct {
	Contract     *MachineManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MachineManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MachineManagerRaw struct {
	Contract *MachineManager // Generic contract binding to access the raw methods on
}

// MachineManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MachineManagerCallerRaw struct {
	Contract *MachineManagerCaller // Generic read-only contract binding to access the raw methods on
}

// MachineManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MachineManagerTransactorRaw struct {
	Contract *MachineManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMachineManager creates a new instance of MachineManager, bound to a specific deployed contract.
func NewMachineManager(address common.Address, backend bind.ContractBackend) (*MachineManager, error) {
	contract, err := bindMachineManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MachineManager{MachineManagerCaller: MachineManagerCaller{contract: contract}, MachineManagerTransactor: MachineManagerTransactor{contract: contract}, MachineManagerFilterer: MachineManagerFilterer{contract: contract}}, nil
}

// NewMachineManagerCaller creates a new read-only instance of MachineManager, bound to a specific deployed contract.
func NewMachineManagerCaller(address common.Address, caller bind.ContractCaller) (*MachineManagerCaller, error) {
	contract, err := bindMachineManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MachineManagerCaller{contract: contract}, nil
}

// NewMachineManagerTransactor creates a new write-only instance of MachineManager, bound to a specific deployed contract.
func NewMachineManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*MachineManagerTransactor, error) {
	contract, err := bindMachineManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MachineManagerTransactor{contract: contract}, nil
}

// NewMachineManagerFilterer creates a new log filterer instance of MachineManager, bound to a specific deployed contract.
func NewMachineManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*MachineManagerFilterer, error) {
	contract, err := bindMachineManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MachineManagerFilterer{contract: contract}, nil
}

// bindMachineManager binds a generic wrapper to an already deployed contract.
func bindMachineManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MachineManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MachineManager *MachineManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MachineManager.Contract.MachineManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MachineManager *MachineManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MachineManager.Contract.MachineManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MachineManager *MachineManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MachineManager.Contract.MachineManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MachineManager *MachineManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MachineManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MachineManager *MachineManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MachineManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MachineManager *MachineManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MachineManager.Contract.contract.Transact(opts, method, params...)
}

// GetActiveTeeMachines is a free data retrieval call binding the contract method 0x45f1f0f0.
//
// Solidity: function getActiveTeeMachines(uint256 _extensionId) view returns(address[] _teeIds, string[] _urls)
func (_MachineManager *MachineManagerCaller) GetActiveTeeMachines(opts *bind.CallOpts, _extensionId *big.Int) (struct {
	TeeIds []common.Address
	Urls   []string
}, error) {
	var out []interface{}
	err := _MachineManager.contract.Call(opts, &out, "getActiveTeeMachines", _extensionId)

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
func (_MachineManager *MachineManagerSession) GetActiveTeeMachines(_extensionId *big.Int) (struct {
	TeeIds []common.Address
	Urls   []string
}, error) {
	return _MachineManager.Contract.GetActiveTeeMachines(&_MachineManager.CallOpts, _extensionId)
}

// GetActiveTeeMachines is a free data retrieval call binding the contract method 0x45f1f0f0.
//
// Solidity: function getActiveTeeMachines(uint256 _extensionId) view returns(address[] _teeIds, string[] _urls)
func (_MachineManager *MachineManagerCallerSession) GetActiveTeeMachines(_extensionId *big.Int) (struct {
	TeeIds []common.Address
	Urls   []string
}, error) {
	return _MachineManager.Contract.GetActiveTeeMachines(&_MachineManager.CallOpts, _extensionId)
}

// GetAllActiveTeeMachines is a free data retrieval call binding the contract method 0xa5578af4.
//
// Solidity: function getAllActiveTeeMachines(uint256 _start, uint256 _end) view returns(address[] _teeIds, string[] _urls, uint256 _totalLength)
func (_MachineManager *MachineManagerCaller) GetAllActiveTeeMachines(opts *bind.CallOpts, _start *big.Int, _end *big.Int) (struct {
	TeeIds      []common.Address
	Urls        []string
	TotalLength *big.Int
}, error) {
	var out []interface{}
	err := _MachineManager.contract.Call(opts, &out, "getAllActiveTeeMachines", _start, _end)

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
func (_MachineManager *MachineManagerSession) GetAllActiveTeeMachines(_start *big.Int, _end *big.Int) (struct {
	TeeIds      []common.Address
	Urls        []string
	TotalLength *big.Int
}, error) {
	return _MachineManager.Contract.GetAllActiveTeeMachines(&_MachineManager.CallOpts, _start, _end)
}

// GetAllActiveTeeMachines is a free data retrieval call binding the contract method 0xa5578af4.
//
// Solidity: function getAllActiveTeeMachines(uint256 _start, uint256 _end) view returns(address[] _teeIds, string[] _urls, uint256 _totalLength)
func (_MachineManager *MachineManagerCallerSession) GetAllActiveTeeMachines(_start *big.Int, _end *big.Int) (struct {
	TeeIds      []common.Address
	Urls        []string
	TotalLength *big.Int
}, error) {
	return _MachineManager.Contract.GetAllActiveTeeMachines(&_MachineManager.CallOpts, _start, _end)
}

// GetExtensionId is a free data retrieval call binding the contract method 0xaa5bb892.
//
// Solidity: function getExtensionId(address _teeId) view returns(uint256)
func (_MachineManager *MachineManagerCaller) GetExtensionId(opts *bind.CallOpts, _teeId common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MachineManager.contract.Call(opts, &out, "getExtensionId", _teeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExtensionId is a free data retrieval call binding the contract method 0xaa5bb892.
//
// Solidity: function getExtensionId(address _teeId) view returns(uint256)
func (_MachineManager *MachineManagerSession) GetExtensionId(_teeId common.Address) (*big.Int, error) {
	return _MachineManager.Contract.GetExtensionId(&_MachineManager.CallOpts, _teeId)
}

// GetExtensionId is a free data retrieval call binding the contract method 0xaa5bb892.
//
// Solidity: function getExtensionId(address _teeId) view returns(uint256)
func (_MachineManager *MachineManagerCallerSession) GetExtensionId(_teeId common.Address) (*big.Int, error) {
	return _MachineManager.Contract.GetExtensionId(&_MachineManager.CallOpts, _teeId)
}

// GetInitialSigningPolicyId is a free data retrieval call binding the contract method 0xb2e51edd.
//
// Solidity: function getInitialSigningPolicyId(address _teeId) view returns(uint32)
func (_MachineManager *MachineManagerCaller) GetInitialSigningPolicyId(opts *bind.CallOpts, _teeId common.Address) (uint32, error) {
	var out []interface{}
	err := _MachineManager.contract.Call(opts, &out, "getInitialSigningPolicyId", _teeId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetInitialSigningPolicyId is a free data retrieval call binding the contract method 0xb2e51edd.
//
// Solidity: function getInitialSigningPolicyId(address _teeId) view returns(uint32)
func (_MachineManager *MachineManagerSession) GetInitialSigningPolicyId(_teeId common.Address) (uint32, error) {
	return _MachineManager.Contract.GetInitialSigningPolicyId(&_MachineManager.CallOpts, _teeId)
}

// GetInitialSigningPolicyId is a free data retrieval call binding the contract method 0xb2e51edd.
//
// Solidity: function getInitialSigningPolicyId(address _teeId) view returns(uint32)
func (_MachineManager *MachineManagerCallerSession) GetInitialSigningPolicyId(_teeId common.Address) (uint32, error) {
	return _MachineManager.Contract.GetInitialSigningPolicyId(&_MachineManager.CallOpts, _teeId)
}

// GetLastStatusChangeTs is a free data retrieval call binding the contract method 0x1e17eb5a.
//
// Solidity: function getLastStatusChangeTs(address _teeId) view returns(uint256)
func (_MachineManager *MachineManagerCaller) GetLastStatusChangeTs(opts *bind.CallOpts, _teeId common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MachineManager.contract.Call(opts, &out, "getLastStatusChangeTs", _teeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastStatusChangeTs is a free data retrieval call binding the contract method 0x1e17eb5a.
//
// Solidity: function getLastStatusChangeTs(address _teeId) view returns(uint256)
func (_MachineManager *MachineManagerSession) GetLastStatusChangeTs(_teeId common.Address) (*big.Int, error) {
	return _MachineManager.Contract.GetLastStatusChangeTs(&_MachineManager.CallOpts, _teeId)
}

// GetLastStatusChangeTs is a free data retrieval call binding the contract method 0x1e17eb5a.
//
// Solidity: function getLastStatusChangeTs(address _teeId) view returns(uint256)
func (_MachineManager *MachineManagerCallerSession) GetLastStatusChangeTs(_teeId common.Address) (*big.Int, error) {
	return _MachineManager.Contract.GetLastStatusChangeTs(&_MachineManager.CallOpts, _teeId)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address _teeId) view returns((bytes32,bytes32))
func (_MachineManager *MachineManagerCaller) GetPublicKey(opts *bind.CallOpts, _teeId common.Address) (PublicKey, error) {
	var out []interface{}
	err := _MachineManager.contract.Call(opts, &out, "getPublicKey", _teeId)

	if err != nil {
		return *new(PublicKey), err
	}

	out0 := *abi.ConvertType(out[0], new(PublicKey)).(*PublicKey)

	return out0, err

}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address _teeId) view returns((bytes32,bytes32))
func (_MachineManager *MachineManagerSession) GetPublicKey(_teeId common.Address) (PublicKey, error) {
	return _MachineManager.Contract.GetPublicKey(&_MachineManager.CallOpts, _teeId)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address _teeId) view returns((bytes32,bytes32))
func (_MachineManager *MachineManagerCallerSession) GetPublicKey(_teeId common.Address) (PublicKey, error) {
	return _MachineManager.Contract.GetPublicKey(&_MachineManager.CallOpts, _teeId)
}

// GetRandomTeeIds is a free data retrieval call binding the contract method 0xfeeabcbf.
//
// Solidity: function getRandomTeeIds(uint256 _extensionId, uint256 _count) view returns(address[] _teeIds)
func (_MachineManager *MachineManagerCaller) GetRandomTeeIds(opts *bind.CallOpts, _extensionId *big.Int, _count *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _MachineManager.contract.Call(opts, &out, "getRandomTeeIds", _extensionId, _count)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRandomTeeIds is a free data retrieval call binding the contract method 0xfeeabcbf.
//
// Solidity: function getRandomTeeIds(uint256 _extensionId, uint256 _count) view returns(address[] _teeIds)
func (_MachineManager *MachineManagerSession) GetRandomTeeIds(_extensionId *big.Int, _count *big.Int) ([]common.Address, error) {
	return _MachineManager.Contract.GetRandomTeeIds(&_MachineManager.CallOpts, _extensionId, _count)
}

// GetRandomTeeIds is a free data retrieval call binding the contract method 0xfeeabcbf.
//
// Solidity: function getRandomTeeIds(uint256 _extensionId, uint256 _count) view returns(address[] _teeIds)
func (_MachineManager *MachineManagerCallerSession) GetRandomTeeIds(_extensionId *big.Int, _count *big.Int) ([]common.Address, error) {
	return _MachineManager.Contract.GetRandomTeeIds(&_MachineManager.CallOpts, _extensionId, _count)
}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_MachineManager *MachineManagerCaller) GetTeeMachine(opts *bind.CallOpts, _teeId common.Address) (IMachineManagerTeeMachine, error) {
	var out []interface{}
	err := _MachineManager.contract.Call(opts, &out, "getTeeMachine", _teeId)

	if err != nil {
		return *new(IMachineManagerTeeMachine), err
	}

	out0 := *abi.ConvertType(out[0], new(IMachineManagerTeeMachine)).(*IMachineManagerTeeMachine)

	return out0, err

}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_MachineManager *MachineManagerSession) GetTeeMachine(_teeId common.Address) (IMachineManagerTeeMachine, error) {
	return _MachineManager.Contract.GetTeeMachine(&_MachineManager.CallOpts, _teeId)
}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_MachineManager *MachineManagerCallerSession) GetTeeMachine(_teeId common.Address) (IMachineManagerTeeMachine, error) {
	return _MachineManager.Contract.GetTeeMachine(&_MachineManager.CallOpts, _teeId)
}

// GetTeeMachineOwner is a free data retrieval call binding the contract method 0xb7e773e9.
//
// Solidity: function getTeeMachineOwner(address _teeId) view returns(address)
func (_MachineManager *MachineManagerCaller) GetTeeMachineOwner(opts *bind.CallOpts, _teeId common.Address) (common.Address, error) {
	var out []interface{}
	err := _MachineManager.contract.Call(opts, &out, "getTeeMachineOwner", _teeId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeeMachineOwner is a free data retrieval call binding the contract method 0xb7e773e9.
//
// Solidity: function getTeeMachineOwner(address _teeId) view returns(address)
func (_MachineManager *MachineManagerSession) GetTeeMachineOwner(_teeId common.Address) (common.Address, error) {
	return _MachineManager.Contract.GetTeeMachineOwner(&_MachineManager.CallOpts, _teeId)
}

// GetTeeMachineOwner is a free data retrieval call binding the contract method 0xb7e773e9.
//
// Solidity: function getTeeMachineOwner(address _teeId) view returns(address)
func (_MachineManager *MachineManagerCallerSession) GetTeeMachineOwner(_teeId common.Address) (common.Address, error) {
	return _MachineManager.Contract.GetTeeMachineOwner(&_MachineManager.CallOpts, _teeId)
}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_MachineManager *MachineManagerCaller) GetTeeMachineStatus(opts *bind.CallOpts, _teeId common.Address) (uint8, error) {
	var out []interface{}
	err := _MachineManager.contract.Call(opts, &out, "getTeeMachineStatus", _teeId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_MachineManager *MachineManagerSession) GetTeeMachineStatus(_teeId common.Address) (uint8, error) {
	return _MachineManager.Contract.GetTeeMachineStatus(&_MachineManager.CallOpts, _teeId)
}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_MachineManager *MachineManagerCallerSession) GetTeeMachineStatus(_teeId common.Address) (uint8, error) {
	return _MachineManager.Contract.GetTeeMachineStatus(&_MachineManager.CallOpts, _teeId)
}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_MachineManager *MachineManagerCaller) GetTeeMachineWithAttestationData(opts *bind.CallOpts, _teeId common.Address) (IMachineManagerTeeMachineWithAttestationData, error) {
	var out []interface{}
	err := _MachineManager.contract.Call(opts, &out, "getTeeMachineWithAttestationData", _teeId)

	if err != nil {
		return *new(IMachineManagerTeeMachineWithAttestationData), err
	}

	out0 := *abi.ConvertType(out[0], new(IMachineManagerTeeMachineWithAttestationData)).(*IMachineManagerTeeMachineWithAttestationData)

	return out0, err

}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_MachineManager *MachineManagerSession) GetTeeMachineWithAttestationData(_teeId common.Address) (IMachineManagerTeeMachineWithAttestationData, error) {
	return _MachineManager.Contract.GetTeeMachineWithAttestationData(&_MachineManager.CallOpts, _teeId)
}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_MachineManager *MachineManagerCallerSession) GetTeeMachineWithAttestationData(_teeId common.Address) (IMachineManagerTeeMachineWithAttestationData, error) {
	return _MachineManager.Contract.GetTeeMachineWithAttestationData(&_MachineManager.CallOpts, _teeId)
}

// Ban is a paid mutator transaction binding the contract method 0x97c3ccd8.
//
// Solidity: function ban(address _teeId) returns()
func (_MachineManager *MachineManagerTransactor) Ban(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _MachineManager.contract.Transact(opts, "ban", _teeId)
}

// Ban is a paid mutator transaction binding the contract method 0x97c3ccd8.
//
// Solidity: function ban(address _teeId) returns()
func (_MachineManager *MachineManagerSession) Ban(_teeId common.Address) (*types.Transaction, error) {
	return _MachineManager.Contract.Ban(&_MachineManager.TransactOpts, _teeId)
}

// Ban is a paid mutator transaction binding the contract method 0x97c3ccd8.
//
// Solidity: function ban(address _teeId) returns()
func (_MachineManager *MachineManagerTransactorSession) Ban(_teeId common.Address) (*types.Transaction, error) {
	return _MachineManager.Contract.Ban(&_MachineManager.TransactOpts, _teeId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x80f2abe3.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (_MachineManager *MachineManagerTransactor) ConfirmOwnership(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _MachineManager.contract.Transact(opts, "confirmOwnership", _teeId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x80f2abe3.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (_MachineManager *MachineManagerSession) ConfirmOwnership(_teeId common.Address) (*types.Transaction, error) {
	return _MachineManager.Contract.ConfirmOwnership(&_MachineManager.TransactOpts, _teeId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x80f2abe3.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (_MachineManager *MachineManagerTransactorSession) ConfirmOwnership(_teeId common.Address) (*types.Transaction, error) {
	return _MachineManager.Contract.ConfirmOwnership(&_MachineManager.TransactOpts, _teeId)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _teeId) returns()
func (_MachineManager *MachineManagerTransactor) Pause(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _MachineManager.contract.Transact(opts, "pause", _teeId)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _teeId) returns()
func (_MachineManager *MachineManagerSession) Pause(_teeId common.Address) (*types.Transaction, error) {
	return _MachineManager.Contract.Pause(&_MachineManager.TransactOpts, _teeId)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _teeId) returns()
func (_MachineManager *MachineManagerTransactorSession) Pause(_teeId common.Address) (*types.Transaction, error) {
	return _MachineManager.Contract.Pause(&_MachineManager.TransactOpts, _teeId)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0xd0824eac.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_MachineManager *MachineManagerTransactor) PauseWithProof(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _MachineManager.contract.Transact(opts, "pauseWithProof", _proof)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0xd0824eac.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_MachineManager *MachineManagerSession) PauseWithProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _MachineManager.Contract.PauseWithProof(&_MachineManager.TransactOpts, _proof)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0xd0824eac.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_MachineManager *MachineManagerTransactorSession) PauseWithProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _MachineManager.Contract.PauseWithProof(&_MachineManager.TransactOpts, _proof)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x38195139.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (_MachineManager *MachineManagerTransactor) ProposeNewOwner(opts *bind.TransactOpts, _teeId common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _MachineManager.contract.Transact(opts, "proposeNewOwner", _teeId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x38195139.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (_MachineManager *MachineManagerSession) ProposeNewOwner(_teeId common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _MachineManager.Contract.ProposeNewOwner(&_MachineManager.TransactOpts, _teeId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x38195139.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (_MachineManager *MachineManagerTransactorSession) ProposeNewOwner(_teeId common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _MachineManager.Contract.ProposeNewOwner(&_MachineManager.TransactOpts, _teeId, _newOwner)
}

// Register is a paid mutator transaction binding the contract method 0x95ef52e0.
//
// Solidity: function register((uint256,address,bytes32,bytes32,(bytes32,bytes32)) _teeMachineData, (uint8,bytes32,bytes32) _teeMachineDataSignature, address _teeProxyId, string _url, address _claimBackAddress) payable returns()
func (_MachineManager *MachineManagerTransactor) Register(opts *bind.TransactOpts, _teeMachineData IMachineManagerTeeMachineData, _teeMachineDataSignature Signature, _teeProxyId common.Address, _url string, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _MachineManager.contract.Transact(opts, "register", _teeMachineData, _teeMachineDataSignature, _teeProxyId, _url, _claimBackAddress)
}

// Register is a paid mutator transaction binding the contract method 0x95ef52e0.
//
// Solidity: function register((uint256,address,bytes32,bytes32,(bytes32,bytes32)) _teeMachineData, (uint8,bytes32,bytes32) _teeMachineDataSignature, address _teeProxyId, string _url, address _claimBackAddress) payable returns()
func (_MachineManager *MachineManagerSession) Register(_teeMachineData IMachineManagerTeeMachineData, _teeMachineDataSignature Signature, _teeProxyId common.Address, _url string, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _MachineManager.Contract.Register(&_MachineManager.TransactOpts, _teeMachineData, _teeMachineDataSignature, _teeProxyId, _url, _claimBackAddress)
}

// Register is a paid mutator transaction binding the contract method 0x95ef52e0.
//
// Solidity: function register((uint256,address,bytes32,bytes32,(bytes32,bytes32)) _teeMachineData, (uint8,bytes32,bytes32) _teeMachineDataSignature, address _teeProxyId, string _url, address _claimBackAddress) payable returns()
func (_MachineManager *MachineManagerTransactorSession) Register(_teeMachineData IMachineManagerTeeMachineData, _teeMachineDataSignature Signature, _teeProxyId common.Address, _url string, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _MachineManager.Contract.Register(&_MachineManager.TransactOpts, _teeMachineData, _teeMachineDataSignature, _teeProxyId, _url, _claimBackAddress)
}

// ToProduction is a paid mutator transaction binding the contract method 0x4d40e8fb.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_MachineManager *MachineManagerTransactor) ToProduction(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _MachineManager.contract.Transact(opts, "toProduction", _proof)
}

// ToProduction is a paid mutator transaction binding the contract method 0x4d40e8fb.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_MachineManager *MachineManagerSession) ToProduction(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _MachineManager.Contract.ToProduction(&_MachineManager.TransactOpts, _proof)
}

// ToProduction is a paid mutator transaction binding the contract method 0x4d40e8fb.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_MachineManager *MachineManagerTransactorSession) ToProduction(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _MachineManager.Contract.ToProduction(&_MachineManager.TransactOpts, _proof)
}

// Unban is a paid mutator transaction binding the contract method 0xb9f14557.
//
// Solidity: function unban(address _teeId) returns()
func (_MachineManager *MachineManagerTransactor) Unban(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _MachineManager.contract.Transact(opts, "unban", _teeId)
}

// Unban is a paid mutator transaction binding the contract method 0xb9f14557.
//
// Solidity: function unban(address _teeId) returns()
func (_MachineManager *MachineManagerSession) Unban(_teeId common.Address) (*types.Transaction, error) {
	return _MachineManager.Contract.Unban(&_MachineManager.TransactOpts, _teeId)
}

// Unban is a paid mutator transaction binding the contract method 0xb9f14557.
//
// Solidity: function unban(address _teeId) returns()
func (_MachineManager *MachineManagerTransactorSession) Unban(_teeId common.Address) (*types.Transaction, error) {
	return _MachineManager.Contract.Unban(&_MachineManager.TransactOpts, _teeId)
}

// UpdateTeeMachineSettings is a paid mutator transaction binding the contract method 0x06ed5da4.
//
// Solidity: function updateTeeMachineSettings(address _teeId, address _teeProxyId, string _url) returns()
func (_MachineManager *MachineManagerTransactor) UpdateTeeMachineSettings(opts *bind.TransactOpts, _teeId common.Address, _teeProxyId common.Address, _url string) (*types.Transaction, error) {
	return _MachineManager.contract.Transact(opts, "updateTeeMachineSettings", _teeId, _teeProxyId, _url)
}

// UpdateTeeMachineSettings is a paid mutator transaction binding the contract method 0x06ed5da4.
//
// Solidity: function updateTeeMachineSettings(address _teeId, address _teeProxyId, string _url) returns()
func (_MachineManager *MachineManagerSession) UpdateTeeMachineSettings(_teeId common.Address, _teeProxyId common.Address, _url string) (*types.Transaction, error) {
	return _MachineManager.Contract.UpdateTeeMachineSettings(&_MachineManager.TransactOpts, _teeId, _teeProxyId, _url)
}

// UpdateTeeMachineSettings is a paid mutator transaction binding the contract method 0x06ed5da4.
//
// Solidity: function updateTeeMachineSettings(address _teeId, address _teeProxyId, string _url) returns()
func (_MachineManager *MachineManagerTransactorSession) UpdateTeeMachineSettings(_teeId common.Address, _teeProxyId common.Address, _url string) (*types.Transaction, error) {
	return _MachineManager.Contract.UpdateTeeMachineSettings(&_MachineManager.TransactOpts, _teeId, _teeProxyId, _url)
}

// MachineManagerAvailabilityCheckValidityExtendedIterator is returned from FilterAvailabilityCheckValidityExtended and is used to iterate over the raw logs and unpacked data for AvailabilityCheckValidityExtended events raised by the MachineManager contract.
type MachineManagerAvailabilityCheckValidityExtendedIterator struct {
	Event *MachineManagerAvailabilityCheckValidityExtended // Event containing the contract specifics and raw log

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
func (it *MachineManagerAvailabilityCheckValidityExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineManagerAvailabilityCheckValidityExtended)
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
		it.Event = new(MachineManagerAvailabilityCheckValidityExtended)
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
func (it *MachineManagerAvailabilityCheckValidityExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineManagerAvailabilityCheckValidityExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineManagerAvailabilityCheckValidityExtended represents a AvailabilityCheckValidityExtended event raised by the MachineManager contract.
type MachineManagerAvailabilityCheckValidityExtended struct {
	TeeId common.Address
	Owner common.Address
	EndTs *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAvailabilityCheckValidityExtended is a free log retrieval operation binding the contract event 0xbcea1901d8026bcd3401ce1c443ba0a3e5e16680271d1dcc84cfdd36a0ca44b3.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (_MachineManager *MachineManagerFilterer) FilterAvailabilityCheckValidityExtended(opts *bind.FilterOpts, teeId []common.Address, owner []common.Address) (*MachineManagerAvailabilityCheckValidityExtendedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MachineManager.contract.FilterLogs(opts, "AvailabilityCheckValidityExtended", teeIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MachineManagerAvailabilityCheckValidityExtendedIterator{contract: _MachineManager.contract, event: "AvailabilityCheckValidityExtended", logs: logs, sub: sub}, nil
}

// WatchAvailabilityCheckValidityExtended is a free log subscription operation binding the contract event 0xbcea1901d8026bcd3401ce1c443ba0a3e5e16680271d1dcc84cfdd36a0ca44b3.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (_MachineManager *MachineManagerFilterer) WatchAvailabilityCheckValidityExtended(opts *bind.WatchOpts, sink chan<- *MachineManagerAvailabilityCheckValidityExtended, teeId []common.Address, owner []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MachineManager.contract.WatchLogs(opts, "AvailabilityCheckValidityExtended", teeIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineManagerAvailabilityCheckValidityExtended)
				if err := _MachineManager.contract.UnpackLog(event, "AvailabilityCheckValidityExtended", log); err != nil {
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
func (_MachineManager *MachineManagerFilterer) ParseAvailabilityCheckValidityExtended(log types.Log) (*MachineManagerAvailabilityCheckValidityExtended, error) {
	event := new(MachineManagerAvailabilityCheckValidityExtended)
	if err := _MachineManager.contract.UnpackLog(event, "AvailabilityCheckValidityExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineManagerNewOwnerConfirmedIterator is returned from FilterNewOwnerConfirmed and is used to iterate over the raw logs and unpacked data for NewOwnerConfirmed events raised by the MachineManager contract.
type MachineManagerNewOwnerConfirmedIterator struct {
	Event *MachineManagerNewOwnerConfirmed // Event containing the contract specifics and raw log

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
func (it *MachineManagerNewOwnerConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineManagerNewOwnerConfirmed)
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
		it.Event = new(MachineManagerNewOwnerConfirmed)
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
func (it *MachineManagerNewOwnerConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineManagerNewOwnerConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineManagerNewOwnerConfirmed represents a NewOwnerConfirmed event raised by the MachineManager contract.
type MachineManagerNewOwnerConfirmed struct {
	TeeId    common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerConfirmed is a free log retrieval operation binding the contract event 0xc813c793f67554890ac6f92415119ac4a7535162dc38fb1120036cbd62069fbf.
//
// Solidity: event NewOwnerConfirmed(address indexed teeId, address indexed newOwner)
func (_MachineManager *MachineManagerFilterer) FilterNewOwnerConfirmed(opts *bind.FilterOpts, teeId []common.Address, newOwner []common.Address) (*MachineManagerNewOwnerConfirmedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MachineManager.contract.FilterLogs(opts, "NewOwnerConfirmed", teeIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MachineManagerNewOwnerConfirmedIterator{contract: _MachineManager.contract, event: "NewOwnerConfirmed", logs: logs, sub: sub}, nil
}

// WatchNewOwnerConfirmed is a free log subscription operation binding the contract event 0xc813c793f67554890ac6f92415119ac4a7535162dc38fb1120036cbd62069fbf.
//
// Solidity: event NewOwnerConfirmed(address indexed teeId, address indexed newOwner)
func (_MachineManager *MachineManagerFilterer) WatchNewOwnerConfirmed(opts *bind.WatchOpts, sink chan<- *MachineManagerNewOwnerConfirmed, teeId []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MachineManager.contract.WatchLogs(opts, "NewOwnerConfirmed", teeIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineManagerNewOwnerConfirmed)
				if err := _MachineManager.contract.UnpackLog(event, "NewOwnerConfirmed", log); err != nil {
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
func (_MachineManager *MachineManagerFilterer) ParseNewOwnerConfirmed(log types.Log) (*MachineManagerNewOwnerConfirmed, error) {
	event := new(MachineManagerNewOwnerConfirmed)
	if err := _MachineManager.contract.UnpackLog(event, "NewOwnerConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineManagerNewOwnerProposedIterator is returned from FilterNewOwnerProposed and is used to iterate over the raw logs and unpacked data for NewOwnerProposed events raised by the MachineManager contract.
type MachineManagerNewOwnerProposedIterator struct {
	Event *MachineManagerNewOwnerProposed // Event containing the contract specifics and raw log

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
func (it *MachineManagerNewOwnerProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineManagerNewOwnerProposed)
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
		it.Event = new(MachineManagerNewOwnerProposed)
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
func (it *MachineManagerNewOwnerProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineManagerNewOwnerProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineManagerNewOwnerProposed represents a NewOwnerProposed event raised by the MachineManager contract.
type MachineManagerNewOwnerProposed struct {
	TeeId    common.Address
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerProposed is a free log retrieval operation binding the contract event 0x64420d4a41c6ed4de2bccbf33192eea18e576c5b23c79c3a722d4e9534c2e8d8.
//
// Solidity: event NewOwnerProposed(address indexed teeId, address indexed oldOwner, address indexed newOwner)
func (_MachineManager *MachineManagerFilterer) FilterNewOwnerProposed(opts *bind.FilterOpts, teeId []common.Address, oldOwner []common.Address, newOwner []common.Address) (*MachineManagerNewOwnerProposedIterator, error) {

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

	logs, sub, err := _MachineManager.contract.FilterLogs(opts, "NewOwnerProposed", teeIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MachineManagerNewOwnerProposedIterator{contract: _MachineManager.contract, event: "NewOwnerProposed", logs: logs, sub: sub}, nil
}

// WatchNewOwnerProposed is a free log subscription operation binding the contract event 0x64420d4a41c6ed4de2bccbf33192eea18e576c5b23c79c3a722d4e9534c2e8d8.
//
// Solidity: event NewOwnerProposed(address indexed teeId, address indexed oldOwner, address indexed newOwner)
func (_MachineManager *MachineManagerFilterer) WatchNewOwnerProposed(opts *bind.WatchOpts, sink chan<- *MachineManagerNewOwnerProposed, teeId []common.Address, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MachineManager.contract.WatchLogs(opts, "NewOwnerProposed", teeIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineManagerNewOwnerProposed)
				if err := _MachineManager.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
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
func (_MachineManager *MachineManagerFilterer) ParseNewOwnerProposed(log types.Log) (*MachineManagerNewOwnerProposed, error) {
	event := new(MachineManagerNewOwnerProposed)
	if err := _MachineManager.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineManagerTeeAttestationRequestedIterator is returned from FilterTeeAttestationRequested and is used to iterate over the raw logs and unpacked data for TeeAttestationRequested events raised by the MachineManager contract.
type MachineManagerTeeAttestationRequestedIterator struct {
	Event *MachineManagerTeeAttestationRequested // Event containing the contract specifics and raw log

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
func (it *MachineManagerTeeAttestationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineManagerTeeAttestationRequested)
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
		it.Event = new(MachineManagerTeeAttestationRequested)
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
func (it *MachineManagerTeeAttestationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineManagerTeeAttestationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineManagerTeeAttestationRequested represents a TeeAttestationRequested event raised by the MachineManager contract.
type MachineManagerTeeAttestationRequested struct {
	TeeId     common.Address
	Challenge [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTeeAttestationRequested is a free log retrieval operation binding the contract event 0x88097792b658dddf5ef685fd49ad9d66c9838d549030d9b4bf44c80dd2076db1.
//
// Solidity: event TeeAttestationRequested(address indexed teeId, bytes32 challenge)
func (_MachineManager *MachineManagerFilterer) FilterTeeAttestationRequested(opts *bind.FilterOpts, teeId []common.Address) (*MachineManagerTeeAttestationRequestedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _MachineManager.contract.FilterLogs(opts, "TeeAttestationRequested", teeIdRule)
	if err != nil {
		return nil, err
	}
	return &MachineManagerTeeAttestationRequestedIterator{contract: _MachineManager.contract, event: "TeeAttestationRequested", logs: logs, sub: sub}, nil
}

// WatchTeeAttestationRequested is a free log subscription operation binding the contract event 0x88097792b658dddf5ef685fd49ad9d66c9838d549030d9b4bf44c80dd2076db1.
//
// Solidity: event TeeAttestationRequested(address indexed teeId, bytes32 challenge)
func (_MachineManager *MachineManagerFilterer) WatchTeeAttestationRequested(opts *bind.WatchOpts, sink chan<- *MachineManagerTeeAttestationRequested, teeId []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _MachineManager.contract.WatchLogs(opts, "TeeAttestationRequested", teeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineManagerTeeAttestationRequested)
				if err := _MachineManager.contract.UnpackLog(event, "TeeAttestationRequested", log); err != nil {
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
func (_MachineManager *MachineManagerFilterer) ParseTeeAttestationRequested(log types.Log) (*MachineManagerTeeAttestationRequested, error) {
	event := new(MachineManagerTeeAttestationRequested)
	if err := _MachineManager.contract.UnpackLog(event, "TeeAttestationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineManagerTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the MachineManager contract.
type MachineManagerTeeInstructionsSentIterator struct {
	Event *MachineManagerTeeInstructionsSent // Event containing the contract specifics and raw log

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
func (it *MachineManagerTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineManagerTeeInstructionsSent)
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
		it.Event = new(MachineManagerTeeInstructionsSent)
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
func (it *MachineManagerTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineManagerTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineManagerTeeInstructionsSent represents a TeeInstructionsSent event raised by the MachineManager contract.
type MachineManagerTeeInstructionsSent struct {
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
func (_MachineManager *MachineManagerFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (*MachineManagerTeeInstructionsSentIterator, error) {

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

	logs, sub, err := _MachineManager.contract.FilterLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &MachineManagerTeeInstructionsSentIterator{contract: _MachineManager.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_MachineManager *MachineManagerFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *MachineManagerTeeInstructionsSent, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

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

	logs, sub, err := _MachineManager.contract.WatchLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineManagerTeeInstructionsSent)
				if err := _MachineManager.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
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
func (_MachineManager *MachineManagerFilterer) ParseTeeInstructionsSent(log types.Log) (*MachineManagerTeeInstructionsSent, error) {
	event := new(MachineManagerTeeInstructionsSent)
	if err := _MachineManager.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineManagerTeeMachineRegisteredIterator is returned from FilterTeeMachineRegistered and is used to iterate over the raw logs and unpacked data for TeeMachineRegistered events raised by the MachineManager contract.
type MachineManagerTeeMachineRegisteredIterator struct {
	Event *MachineManagerTeeMachineRegistered // Event containing the contract specifics and raw log

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
func (it *MachineManagerTeeMachineRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineManagerTeeMachineRegistered)
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
		it.Event = new(MachineManagerTeeMachineRegistered)
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
func (it *MachineManagerTeeMachineRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineManagerTeeMachineRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineManagerTeeMachineRegistered represents a TeeMachineRegistered event raised by the MachineManager contract.
type MachineManagerTeeMachineRegistered struct {
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
func (_MachineManager *MachineManagerFilterer) FilterTeeMachineRegistered(opts *bind.FilterOpts, teeId []common.Address, teeProxyId []common.Address, owner []common.Address) (*MachineManagerTeeMachineRegisteredIterator, error) {

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

	logs, sub, err := _MachineManager.contract.FilterLogs(opts, "TeeMachineRegistered", teeIdRule, teeProxyIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MachineManagerTeeMachineRegisteredIterator{contract: _MachineManager.contract, event: "TeeMachineRegistered", logs: logs, sub: sub}, nil
}

// WatchTeeMachineRegistered is a free log subscription operation binding the contract event 0x4a4f6c10f9306f09cfbee5108b4cc42e71d5b6cd887eb42a98c483cb5680164a.
//
// Solidity: event TeeMachineRegistered(address indexed teeId, address indexed teeProxyId, address indexed owner, uint256 extensionId, string url, bytes32 codeHash, bytes32 platform)
func (_MachineManager *MachineManagerFilterer) WatchTeeMachineRegistered(opts *bind.WatchOpts, sink chan<- *MachineManagerTeeMachineRegistered, teeId []common.Address, teeProxyId []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MachineManager.contract.WatchLogs(opts, "TeeMachineRegistered", teeIdRule, teeProxyIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineManagerTeeMachineRegistered)
				if err := _MachineManager.contract.UnpackLog(event, "TeeMachineRegistered", log); err != nil {
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
func (_MachineManager *MachineManagerFilterer) ParseTeeMachineRegistered(log types.Log) (*MachineManagerTeeMachineRegistered, error) {
	event := new(MachineManagerTeeMachineRegistered)
	if err := _MachineManager.contract.UnpackLog(event, "TeeMachineRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineManagerTeeMachineSettingsUpdatedIterator is returned from FilterTeeMachineSettingsUpdated and is used to iterate over the raw logs and unpacked data for TeeMachineSettingsUpdated events raised by the MachineManager contract.
type MachineManagerTeeMachineSettingsUpdatedIterator struct {
	Event *MachineManagerTeeMachineSettingsUpdated // Event containing the contract specifics and raw log

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
func (it *MachineManagerTeeMachineSettingsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineManagerTeeMachineSettingsUpdated)
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
		it.Event = new(MachineManagerTeeMachineSettingsUpdated)
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
func (it *MachineManagerTeeMachineSettingsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineManagerTeeMachineSettingsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineManagerTeeMachineSettingsUpdated represents a TeeMachineSettingsUpdated event raised by the MachineManager contract.
type MachineManagerTeeMachineSettingsUpdated struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeeMachineSettingsUpdated is a free log retrieval operation binding the contract event 0x2ee0145d87c2e0db71dfa482a5d5fbec4a34c66dba75742b647608507c11bb46.
//
// Solidity: event TeeMachineSettingsUpdated(address indexed teeId, address indexed teeProxyId, string url)
func (_MachineManager *MachineManagerFilterer) FilterTeeMachineSettingsUpdated(opts *bind.FilterOpts, teeId []common.Address, teeProxyId []common.Address) (*MachineManagerTeeMachineSettingsUpdatedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var teeProxyIdRule []interface{}
	for _, teeProxyIdItem := range teeProxyId {
		teeProxyIdRule = append(teeProxyIdRule, teeProxyIdItem)
	}

	logs, sub, err := _MachineManager.contract.FilterLogs(opts, "TeeMachineSettingsUpdated", teeIdRule, teeProxyIdRule)
	if err != nil {
		return nil, err
	}
	return &MachineManagerTeeMachineSettingsUpdatedIterator{contract: _MachineManager.contract, event: "TeeMachineSettingsUpdated", logs: logs, sub: sub}, nil
}

// WatchTeeMachineSettingsUpdated is a free log subscription operation binding the contract event 0x2ee0145d87c2e0db71dfa482a5d5fbec4a34c66dba75742b647608507c11bb46.
//
// Solidity: event TeeMachineSettingsUpdated(address indexed teeId, address indexed teeProxyId, string url)
func (_MachineManager *MachineManagerFilterer) WatchTeeMachineSettingsUpdated(opts *bind.WatchOpts, sink chan<- *MachineManagerTeeMachineSettingsUpdated, teeId []common.Address, teeProxyId []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var teeProxyIdRule []interface{}
	for _, teeProxyIdItem := range teeProxyId {
		teeProxyIdRule = append(teeProxyIdRule, teeProxyIdItem)
	}

	logs, sub, err := _MachineManager.contract.WatchLogs(opts, "TeeMachineSettingsUpdated", teeIdRule, teeProxyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineManagerTeeMachineSettingsUpdated)
				if err := _MachineManager.contract.UnpackLog(event, "TeeMachineSettingsUpdated", log); err != nil {
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
func (_MachineManager *MachineManagerFilterer) ParseTeeMachineSettingsUpdated(log types.Log) (*MachineManagerTeeMachineSettingsUpdated, error) {
	event := new(MachineManagerTeeMachineSettingsUpdated)
	if err := _MachineManager.contract.UnpackLog(event, "TeeMachineSettingsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MachineManagerTeeMachineStatusChangedIterator is returned from FilterTeeMachineStatusChanged and is used to iterate over the raw logs and unpacked data for TeeMachineStatusChanged events raised by the MachineManager contract.
type MachineManagerTeeMachineStatusChangedIterator struct {
	Event *MachineManagerTeeMachineStatusChanged // Event containing the contract specifics and raw log

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
func (it *MachineManagerTeeMachineStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MachineManagerTeeMachineStatusChanged)
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
		it.Event = new(MachineManagerTeeMachineStatusChanged)
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
func (it *MachineManagerTeeMachineStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MachineManagerTeeMachineStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MachineManagerTeeMachineStatusChanged represents a TeeMachineStatusChanged event raised by the MachineManager contract.
type MachineManagerTeeMachineStatusChanged struct {
	TeeId     common.Address
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTeeMachineStatusChanged is a free log retrieval operation binding the contract event 0x15cd9c7dcca8fe809f6f492069627d36efb27df330c4e6e1c0840751dba73ab1.
//
// Solidity: event TeeMachineStatusChanged(address indexed teeId, uint8 indexed newStatus)
func (_MachineManager *MachineManagerFilterer) FilterTeeMachineStatusChanged(opts *bind.FilterOpts, teeId []common.Address, newStatus []uint8) (*MachineManagerTeeMachineStatusChangedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _MachineManager.contract.FilterLogs(opts, "TeeMachineStatusChanged", teeIdRule, newStatusRule)
	if err != nil {
		return nil, err
	}
	return &MachineManagerTeeMachineStatusChangedIterator{contract: _MachineManager.contract, event: "TeeMachineStatusChanged", logs: logs, sub: sub}, nil
}

// WatchTeeMachineStatusChanged is a free log subscription operation binding the contract event 0x15cd9c7dcca8fe809f6f492069627d36efb27df330c4e6e1c0840751dba73ab1.
//
// Solidity: event TeeMachineStatusChanged(address indexed teeId, uint8 indexed newStatus)
func (_MachineManager *MachineManagerFilterer) WatchTeeMachineStatusChanged(opts *bind.WatchOpts, sink chan<- *MachineManagerTeeMachineStatusChanged, teeId []common.Address, newStatus []uint8) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _MachineManager.contract.WatchLogs(opts, "TeeMachineStatusChanged", teeIdRule, newStatusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MachineManagerTeeMachineStatusChanged)
				if err := _MachineManager.contract.UnpackLog(event, "TeeMachineStatusChanged", log); err != nil {
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
func (_MachineManager *MachineManagerFilterer) ParseTeeMachineStatusChanged(log types.Log) (*MachineManagerTeeMachineStatusChanged, error) {
	event := new(MachineManagerTeeMachineStatusChanged)
	if err := _MachineManager.contract.UnpackLog(event, "TeeMachineStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
