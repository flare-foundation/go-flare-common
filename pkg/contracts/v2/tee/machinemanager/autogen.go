// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package machinemanager

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

// IMachineManagerTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
}

// IMachineManagerTeeMachineData is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachineData struct {
	ExtensionId    *big.Int
	InitialOwner   common.Address
	CodeHash       [32]byte
	Platform       [32]byte
	PublicKey      PublicKey
	GovernanceHash [32]byte
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
var MachineManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengeTs\",\"type\":\"uint256\"}],\"name\":\"ChallengeExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"EmergencyPauseActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"EmergencyProtectionActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestBody\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseDataOrAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningPolicy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeProxyId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeePublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeePublicKeyOrSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUrl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrExpiredAvailabilityCheck\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooMany\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTs\",\"type\":\"uint256\"}],\"name\":\"AvailabilityCheckValidityExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"name\":\"TeeAttestationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"governanceHash\",\"type\":\"bytes32\"}],\"name\":\"TeeMachineRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"TeeMachineSettingsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIMachineManager.TeeStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"TeeMachineStatusChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"ban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getActiveTeeMachines\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_urls\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"getAllActiveTeeMachines\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_urls\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_totalLength\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getExtensionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getInitialSigningPolicyId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getLastStatusChangeTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getRandomTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachine\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIMachineManager.TeeMachine\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineStatus\",\"outputs\":[{\"internalType\":\"enumIMachineManager.TeeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineWithAttestationData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structIMachineManager.TeeMachineWithAttestationData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"pauseWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"governanceHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIMachineManager.TeeMachineData\",\"name\":\"_teeMachineData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"_teeMachineDataSignature\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"toProduction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"unban\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"updateTeeMachineSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "MachineManager",
}

// MachineManager is an auto generated Go binding around an Ethereum contract.
type MachineManager struct {
	abi abi.ABI
}

// NewMachineManager creates a new instance of MachineManager.
func NewMachineManager() *MachineManager {
	parsed, err := MachineManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MachineManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MachineManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackBan is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x97c3ccd8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ban(address _teeId) returns()
func (machineManager *MachineManager) PackBan(teeId common.Address) []byte {
	enc, err := machineManager.abi.Pack("ban", teeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBan is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x97c3ccd8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ban(address _teeId) returns()
func (machineManager *MachineManager) TryPackBan(teeId common.Address) ([]byte, error) {
	return machineManager.abi.Pack("ban", teeId)
}

// PackConfirmOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x80f2abe3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (machineManager *MachineManager) PackConfirmOwnership(teeId common.Address) []byte {
	enc, err := machineManager.abi.Pack("confirmOwnership", teeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x80f2abe3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (machineManager *MachineManager) TryPackConfirmOwnership(teeId common.Address) ([]byte, error) {
	return machineManager.abi.Pack("confirmOwnership", teeId)
}

// PackGetActiveTeeMachines is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x45f1f0f0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getActiveTeeMachines(uint256 _extensionId) view returns(address[] _teeIds, string[] _urls)
func (machineManager *MachineManager) PackGetActiveTeeMachines(extensionId *big.Int) []byte {
	enc, err := machineManager.abi.Pack("getActiveTeeMachines", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetActiveTeeMachines is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x45f1f0f0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getActiveTeeMachines(uint256 _extensionId) view returns(address[] _teeIds, string[] _urls)
func (machineManager *MachineManager) TryPackGetActiveTeeMachines(extensionId *big.Int) ([]byte, error) {
	return machineManager.abi.Pack("getActiveTeeMachines", extensionId)
}

// GetActiveTeeMachinesOutput serves as a container for the return parameters of contract
// method GetActiveTeeMachines.
type GetActiveTeeMachinesOutput struct {
	TeeIds []common.Address
	Urls   []string
}

// UnpackGetActiveTeeMachines is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x45f1f0f0.
//
// Solidity: function getActiveTeeMachines(uint256 _extensionId) view returns(address[] _teeIds, string[] _urls)
func (machineManager *MachineManager) UnpackGetActiveTeeMachines(data []byte) (GetActiveTeeMachinesOutput, error) {
	out, err := machineManager.abi.Unpack("getActiveTeeMachines", data)
	outstruct := new(GetActiveTeeMachinesOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.TeeIds = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Urls = *abi.ConvertType(out[1], new([]string)).(*[]string)
	return *outstruct, nil
}

// PackGetAllActiveTeeMachines is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa5578af4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAllActiveTeeMachines(uint256 _start, uint256 _end) view returns(address[] _teeIds, string[] _urls, uint256 _totalLength)
func (machineManager *MachineManager) PackGetAllActiveTeeMachines(start *big.Int, end *big.Int) []byte {
	enc, err := machineManager.abi.Pack("getAllActiveTeeMachines", start, end)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAllActiveTeeMachines is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa5578af4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAllActiveTeeMachines(uint256 _start, uint256 _end) view returns(address[] _teeIds, string[] _urls, uint256 _totalLength)
func (machineManager *MachineManager) TryPackGetAllActiveTeeMachines(start *big.Int, end *big.Int) ([]byte, error) {
	return machineManager.abi.Pack("getAllActiveTeeMachines", start, end)
}

// GetAllActiveTeeMachinesOutput serves as a container for the return parameters of contract
// method GetAllActiveTeeMachines.
type GetAllActiveTeeMachinesOutput struct {
	TeeIds      []common.Address
	Urls        []string
	TotalLength *big.Int
}

// UnpackGetAllActiveTeeMachines is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa5578af4.
//
// Solidity: function getAllActiveTeeMachines(uint256 _start, uint256 _end) view returns(address[] _teeIds, string[] _urls, uint256 _totalLength)
func (machineManager *MachineManager) UnpackGetAllActiveTeeMachines(data []byte) (GetAllActiveTeeMachinesOutput, error) {
	out, err := machineManager.abi.Unpack("getAllActiveTeeMachines", data)
	outstruct := new(GetAllActiveTeeMachinesOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.TeeIds = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Urls = *abi.ConvertType(out[1], new([]string)).(*[]string)
	outstruct.TotalLength = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGetExtensionId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaa5bb892.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getExtensionId(address _teeId) view returns(uint256)
func (machineManager *MachineManager) PackGetExtensionId(teeId common.Address) []byte {
	enc, err := machineManager.abi.Pack("getExtensionId", teeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetExtensionId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaa5bb892.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getExtensionId(address _teeId) view returns(uint256)
func (machineManager *MachineManager) TryPackGetExtensionId(teeId common.Address) ([]byte, error) {
	return machineManager.abi.Pack("getExtensionId", teeId)
}

// UnpackGetExtensionId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaa5bb892.
//
// Solidity: function getExtensionId(address _teeId) view returns(uint256)
func (machineManager *MachineManager) UnpackGetExtensionId(data []byte) (*big.Int, error) {
	out, err := machineManager.abi.Unpack("getExtensionId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetInitialSigningPolicyId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb2e51edd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getInitialSigningPolicyId(address _teeId) view returns(uint32)
func (machineManager *MachineManager) PackGetInitialSigningPolicyId(teeId common.Address) []byte {
	enc, err := machineManager.abi.Pack("getInitialSigningPolicyId", teeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetInitialSigningPolicyId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb2e51edd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getInitialSigningPolicyId(address _teeId) view returns(uint32)
func (machineManager *MachineManager) TryPackGetInitialSigningPolicyId(teeId common.Address) ([]byte, error) {
	return machineManager.abi.Pack("getInitialSigningPolicyId", teeId)
}

// UnpackGetInitialSigningPolicyId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb2e51edd.
//
// Solidity: function getInitialSigningPolicyId(address _teeId) view returns(uint32)
func (machineManager *MachineManager) UnpackGetInitialSigningPolicyId(data []byte) (uint32, error) {
	out, err := machineManager.abi.Unpack("getInitialSigningPolicyId", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackGetLastStatusChangeTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e17eb5a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLastStatusChangeTs(address _teeId) view returns(uint256)
func (machineManager *MachineManager) PackGetLastStatusChangeTs(teeId common.Address) []byte {
	enc, err := machineManager.abi.Pack("getLastStatusChangeTs", teeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLastStatusChangeTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e17eb5a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLastStatusChangeTs(address _teeId) view returns(uint256)
func (machineManager *MachineManager) TryPackGetLastStatusChangeTs(teeId common.Address) ([]byte, error) {
	return machineManager.abi.Pack("getLastStatusChangeTs", teeId)
}

// UnpackGetLastStatusChangeTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1e17eb5a.
//
// Solidity: function getLastStatusChangeTs(address _teeId) view returns(uint256)
func (machineManager *MachineManager) UnpackGetLastStatusChangeTs(data []byte) (*big.Int, error) {
	out, err := machineManager.abi.Unpack("getLastStatusChangeTs", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetPublicKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x857cdbb8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPublicKey(address _teeId) view returns((bytes32,bytes32))
func (machineManager *MachineManager) PackGetPublicKey(teeId common.Address) []byte {
	enc, err := machineManager.abi.Pack("getPublicKey", teeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPublicKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x857cdbb8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPublicKey(address _teeId) view returns((bytes32,bytes32))
func (machineManager *MachineManager) TryPackGetPublicKey(teeId common.Address) ([]byte, error) {
	return machineManager.abi.Pack("getPublicKey", teeId)
}

// UnpackGetPublicKey is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x857cdbb8.
//
// Solidity: function getPublicKey(address _teeId) view returns((bytes32,bytes32))
func (machineManager *MachineManager) UnpackGetPublicKey(data []byte) (PublicKey, error) {
	out, err := machineManager.abi.Unpack("getPublicKey", data)
	if err != nil {
		return *new(PublicKey), err
	}
	out0 := *abi.ConvertType(out[0], new(PublicKey)).(*PublicKey)
	return out0, nil
}

// PackGetRandomTeeIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfeeabcbf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRandomTeeIds(uint256 _extensionId, uint256 _count) view returns(address[] _teeIds)
func (machineManager *MachineManager) PackGetRandomTeeIds(extensionId *big.Int, count *big.Int) []byte {
	enc, err := machineManager.abi.Pack("getRandomTeeIds", extensionId, count)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRandomTeeIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfeeabcbf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRandomTeeIds(uint256 _extensionId, uint256 _count) view returns(address[] _teeIds)
func (machineManager *MachineManager) TryPackGetRandomTeeIds(extensionId *big.Int, count *big.Int) ([]byte, error) {
	return machineManager.abi.Pack("getRandomTeeIds", extensionId, count)
}

// UnpackGetRandomTeeIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfeeabcbf.
//
// Solidity: function getRandomTeeIds(uint256 _extensionId, uint256 _count) view returns(address[] _teeIds)
func (machineManager *MachineManager) UnpackGetRandomTeeIds(data []byte) ([]common.Address, error) {
	out, err := machineManager.abi.Unpack("getRandomTeeIds", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetTeeMachine is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ede508f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (machineManager *MachineManager) PackGetTeeMachine(teeId common.Address) []byte {
	enc, err := machineManager.abi.Pack("getTeeMachine", teeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTeeMachine is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ede508f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (machineManager *MachineManager) TryPackGetTeeMachine(teeId common.Address) ([]byte, error) {
	return machineManager.abi.Pack("getTeeMachine", teeId)
}

// UnpackGetTeeMachine is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (machineManager *MachineManager) UnpackGetTeeMachine(data []byte) (IMachineManagerTeeMachine, error) {
	out, err := machineManager.abi.Unpack("getTeeMachine", data)
	if err != nil {
		return *new(IMachineManagerTeeMachine), err
	}
	out0 := *abi.ConvertType(out[0], new(IMachineManagerTeeMachine)).(*IMachineManagerTeeMachine)
	return out0, nil
}

// PackGetTeeMachineOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb7e773e9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTeeMachineOwner(address _teeId) view returns(address)
func (machineManager *MachineManager) PackGetTeeMachineOwner(teeId common.Address) []byte {
	enc, err := machineManager.abi.Pack("getTeeMachineOwner", teeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTeeMachineOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb7e773e9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTeeMachineOwner(address _teeId) view returns(address)
func (machineManager *MachineManager) TryPackGetTeeMachineOwner(teeId common.Address) ([]byte, error) {
	return machineManager.abi.Pack("getTeeMachineOwner", teeId)
}

// UnpackGetTeeMachineOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb7e773e9.
//
// Solidity: function getTeeMachineOwner(address _teeId) view returns(address)
func (machineManager *MachineManager) UnpackGetTeeMachineOwner(data []byte) (common.Address, error) {
	out, err := machineManager.abi.Unpack("getTeeMachineOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetTeeMachineStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x25e30221.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (machineManager *MachineManager) PackGetTeeMachineStatus(teeId common.Address) []byte {
	enc, err := machineManager.abi.Pack("getTeeMachineStatus", teeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTeeMachineStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x25e30221.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (machineManager *MachineManager) TryPackGetTeeMachineStatus(teeId common.Address) ([]byte, error) {
	return machineManager.abi.Pack("getTeeMachineStatus", teeId)
}

// UnpackGetTeeMachineStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (machineManager *MachineManager) UnpackGetTeeMachineStatus(data []byte) (uint8, error) {
	out, err := machineManager.abi.Unpack("getTeeMachineStatus", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackGetTeeMachineWithAttestationData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x883e6c9a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (machineManager *MachineManager) PackGetTeeMachineWithAttestationData(teeId common.Address) []byte {
	enc, err := machineManager.abi.Pack("getTeeMachineWithAttestationData", teeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTeeMachineWithAttestationData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x883e6c9a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (machineManager *MachineManager) TryPackGetTeeMachineWithAttestationData(teeId common.Address) ([]byte, error) {
	return machineManager.abi.Pack("getTeeMachineWithAttestationData", teeId)
}

// UnpackGetTeeMachineWithAttestationData is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (machineManager *MachineManager) UnpackGetTeeMachineWithAttestationData(data []byte) (IMachineManagerTeeMachineWithAttestationData, error) {
	out, err := machineManager.abi.Unpack("getTeeMachineWithAttestationData", data)
	if err != nil {
		return *new(IMachineManagerTeeMachineWithAttestationData), err
	}
	out0 := *abi.ConvertType(out[0], new(IMachineManagerTeeMachineWithAttestationData)).(*IMachineManagerTeeMachineWithAttestationData)
	return out0, nil
}

// PackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x76a67a51.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pause(address _teeId) returns()
func (machineManager *MachineManager) PackPause(teeId common.Address) []byte {
	enc, err := machineManager.abi.Pack("pause", teeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPause is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x76a67a51.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pause(address _teeId) returns()
func (machineManager *MachineManager) TryPackPause(teeId common.Address) ([]byte, error) {
	return machineManager.abi.Pack("pause", teeId)
}

// PackPauseWithProof is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd0824eac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (machineManager *MachineManager) PackPauseWithProof(proof ITeeAvailabilityCheckProof) []byte {
	enc, err := machineManager.abi.Pack("pauseWithProof", proof)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPauseWithProof is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd0824eac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (machineManager *MachineManager) TryPackPauseWithProof(proof ITeeAvailabilityCheckProof) ([]byte, error) {
	return machineManager.abi.Pack("pauseWithProof", proof)
}

// PackProposeNewOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38195139.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (machineManager *MachineManager) PackProposeNewOwner(teeId common.Address, newOwner common.Address) []byte {
	enc, err := machineManager.abi.Pack("proposeNewOwner", teeId, newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposeNewOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38195139.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (machineManager *MachineManager) TryPackProposeNewOwner(teeId common.Address, newOwner common.Address) ([]byte, error) {
	return machineManager.abi.Pack("proposeNewOwner", teeId, newOwner)
}

// PackRegister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2b4d762.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function register((uint256,address,bytes32,bytes32,(bytes32,bytes32),bytes32) _teeMachineData, (uint8,bytes32,bytes32) _teeMachineDataSignature, address _teeProxyId, string _url, address _claimBackAddress) payable returns()
func (machineManager *MachineManager) PackRegister(teeMachineData IMachineManagerTeeMachineData, teeMachineDataSignature Signature, teeProxyId common.Address, url string, claimBackAddress common.Address) []byte {
	enc, err := machineManager.abi.Pack("register", teeMachineData, teeMachineDataSignature, teeProxyId, url, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2b4d762.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function register((uint256,address,bytes32,bytes32,(bytes32,bytes32),bytes32) _teeMachineData, (uint8,bytes32,bytes32) _teeMachineDataSignature, address _teeProxyId, string _url, address _claimBackAddress) payable returns()
func (machineManager *MachineManager) TryPackRegister(teeMachineData IMachineManagerTeeMachineData, teeMachineDataSignature Signature, teeProxyId common.Address, url string, claimBackAddress common.Address) ([]byte, error) {
	return machineManager.abi.Pack("register", teeMachineData, teeMachineDataSignature, teeProxyId, url, claimBackAddress)
}

// PackToProduction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4d40e8fb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (machineManager *MachineManager) PackToProduction(proof ITeeAvailabilityCheckProof) []byte {
	enc, err := machineManager.abi.Pack("toProduction", proof)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackToProduction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4d40e8fb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (machineManager *MachineManager) TryPackToProduction(proof ITeeAvailabilityCheckProof) ([]byte, error) {
	return machineManager.abi.Pack("toProduction", proof)
}

// PackUnban is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb9f14557.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unban(address _teeId) returns()
func (machineManager *MachineManager) PackUnban(teeId common.Address) []byte {
	enc, err := machineManager.abi.Pack("unban", teeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnban is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb9f14557.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unban(address _teeId) returns()
func (machineManager *MachineManager) TryPackUnban(teeId common.Address) ([]byte, error) {
	return machineManager.abi.Pack("unban", teeId)
}

// PackUpdateTeeMachineSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06ed5da4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateTeeMachineSettings(address _teeId, address _teeProxyId, string _url) returns()
func (machineManager *MachineManager) PackUpdateTeeMachineSettings(teeId common.Address, teeProxyId common.Address, url string) []byte {
	enc, err := machineManager.abi.Pack("updateTeeMachineSettings", teeId, teeProxyId, url)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateTeeMachineSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06ed5da4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateTeeMachineSettings(address _teeId, address _teeProxyId, string _url) returns()
func (machineManager *MachineManager) TryPackUpdateTeeMachineSettings(teeId common.Address, teeProxyId common.Address, url string) ([]byte, error) {
	return machineManager.abi.Pack("updateTeeMachineSettings", teeId, teeProxyId, url)
}

// MachineManagerAvailabilityCheckValidityExtended represents a AvailabilityCheckValidityExtended event raised by the MachineManager contract.
type MachineManagerAvailabilityCheckValidityExtended struct {
	TeeId common.Address
	Owner common.Address
	EndTs *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const MachineManagerAvailabilityCheckValidityExtendedEventName = "AvailabilityCheckValidityExtended"

// ContractEventName returns the user-defined event name.
func (MachineManagerAvailabilityCheckValidityExtended) ContractEventName() string {
	return MachineManagerAvailabilityCheckValidityExtendedEventName
}

// UnpackAvailabilityCheckValidityExtendedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (machineManager *MachineManager) UnpackAvailabilityCheckValidityExtendedEvent(log *types.Log) (*MachineManagerAvailabilityCheckValidityExtended, error) {
	event := "AvailabilityCheckValidityExtended"
	if len(log.Topics) == 0 || log.Topics[0] != machineManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineManagerAvailabilityCheckValidityExtended)
	if len(log.Data) > 0 {
		if err := machineManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineManager.abi.Events[event].Inputs {
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

// MachineManagerNewOwnerConfirmed represents a NewOwnerConfirmed event raised by the MachineManager contract.
type MachineManagerNewOwnerConfirmed struct {
	TeeId    common.Address
	NewOwner common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const MachineManagerNewOwnerConfirmedEventName = "NewOwnerConfirmed"

// ContractEventName returns the user-defined event name.
func (MachineManagerNewOwnerConfirmed) ContractEventName() string {
	return MachineManagerNewOwnerConfirmedEventName
}

// UnpackNewOwnerConfirmedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewOwnerConfirmed(address indexed teeId, address indexed newOwner)
func (machineManager *MachineManager) UnpackNewOwnerConfirmedEvent(log *types.Log) (*MachineManagerNewOwnerConfirmed, error) {
	event := "NewOwnerConfirmed"
	if len(log.Topics) == 0 || log.Topics[0] != machineManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineManagerNewOwnerConfirmed)
	if len(log.Data) > 0 {
		if err := machineManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineManager.abi.Events[event].Inputs {
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

// MachineManagerNewOwnerProposed represents a NewOwnerProposed event raised by the MachineManager contract.
type MachineManagerNewOwnerProposed struct {
	TeeId    common.Address
	OldOwner common.Address
	NewOwner common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const MachineManagerNewOwnerProposedEventName = "NewOwnerProposed"

// ContractEventName returns the user-defined event name.
func (MachineManagerNewOwnerProposed) ContractEventName() string {
	return MachineManagerNewOwnerProposedEventName
}

// UnpackNewOwnerProposedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewOwnerProposed(address indexed teeId, address indexed oldOwner, address indexed newOwner)
func (machineManager *MachineManager) UnpackNewOwnerProposedEvent(log *types.Log) (*MachineManagerNewOwnerProposed, error) {
	event := "NewOwnerProposed"
	if len(log.Topics) == 0 || log.Topics[0] != machineManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineManagerNewOwnerProposed)
	if len(log.Data) > 0 {
		if err := machineManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineManager.abi.Events[event].Inputs {
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

// MachineManagerTeeAttestationRequested represents a TeeAttestationRequested event raised by the MachineManager contract.
type MachineManagerTeeAttestationRequested struct {
	TeeId     common.Address
	Challenge [32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const MachineManagerTeeAttestationRequestedEventName = "TeeAttestationRequested"

// ContractEventName returns the user-defined event name.
func (MachineManagerTeeAttestationRequested) ContractEventName() string {
	return MachineManagerTeeAttestationRequestedEventName
}

// UnpackTeeAttestationRequestedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeAttestationRequested(address indexed teeId, bytes32 challenge)
func (machineManager *MachineManager) UnpackTeeAttestationRequestedEvent(log *types.Log) (*MachineManagerTeeAttestationRequested, error) {
	event := "TeeAttestationRequested"
	if len(log.Topics) == 0 || log.Topics[0] != machineManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineManagerTeeAttestationRequested)
	if len(log.Data) > 0 {
		if err := machineManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineManager.abi.Events[event].Inputs {
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
	Raw                *types.Log // Blockchain specific contextual infos
}

const MachineManagerTeeInstructionsSentEventName = "TeeInstructionsSent"

// ContractEventName returns the user-defined event name.
func (MachineManagerTeeInstructionsSent) ContractEventName() string {
	return MachineManagerTeeInstructionsSentEventName
}

// UnpackTeeInstructionsSentEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (machineManager *MachineManager) UnpackTeeInstructionsSentEvent(log *types.Log) (*MachineManagerTeeInstructionsSent, error) {
	event := "TeeInstructionsSent"
	if len(log.Topics) == 0 || log.Topics[0] != machineManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineManagerTeeInstructionsSent)
	if len(log.Data) > 0 {
		if err := machineManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineManager.abi.Events[event].Inputs {
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

// MachineManagerTeeMachineRegistered represents a TeeMachineRegistered event raised by the MachineManager contract.
type MachineManagerTeeMachineRegistered struct {
	TeeId          common.Address
	TeeProxyId     common.Address
	Owner          common.Address
	ExtensionId    *big.Int
	Url            string
	CodeHash       [32]byte
	Platform       [32]byte
	GovernanceHash [32]byte
	Raw            *types.Log // Blockchain specific contextual infos
}

const MachineManagerTeeMachineRegisteredEventName = "TeeMachineRegistered"

// ContractEventName returns the user-defined event name.
func (MachineManagerTeeMachineRegistered) ContractEventName() string {
	return MachineManagerTeeMachineRegisteredEventName
}

// UnpackTeeMachineRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeMachineRegistered(address indexed teeId, address indexed teeProxyId, address indexed owner, uint256 extensionId, string url, bytes32 codeHash, bytes32 platform, bytes32 governanceHash)
func (machineManager *MachineManager) UnpackTeeMachineRegisteredEvent(log *types.Log) (*MachineManagerTeeMachineRegistered, error) {
	event := "TeeMachineRegistered"
	if len(log.Topics) == 0 || log.Topics[0] != machineManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineManagerTeeMachineRegistered)
	if len(log.Data) > 0 {
		if err := machineManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineManager.abi.Events[event].Inputs {
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

// MachineManagerTeeMachineSettingsUpdated represents a TeeMachineSettingsUpdated event raised by the MachineManager contract.
type MachineManagerTeeMachineSettingsUpdated struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
	Raw        *types.Log // Blockchain specific contextual infos
}

const MachineManagerTeeMachineSettingsUpdatedEventName = "TeeMachineSettingsUpdated"

// ContractEventName returns the user-defined event name.
func (MachineManagerTeeMachineSettingsUpdated) ContractEventName() string {
	return MachineManagerTeeMachineSettingsUpdatedEventName
}

// UnpackTeeMachineSettingsUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeMachineSettingsUpdated(address indexed teeId, address indexed teeProxyId, string url)
func (machineManager *MachineManager) UnpackTeeMachineSettingsUpdatedEvent(log *types.Log) (*MachineManagerTeeMachineSettingsUpdated, error) {
	event := "TeeMachineSettingsUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != machineManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineManagerTeeMachineSettingsUpdated)
	if len(log.Data) > 0 {
		if err := machineManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineManager.abi.Events[event].Inputs {
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

// MachineManagerTeeMachineStatusChanged represents a TeeMachineStatusChanged event raised by the MachineManager contract.
type MachineManagerTeeMachineStatusChanged struct {
	TeeId     common.Address
	NewStatus uint8
	Raw       *types.Log // Blockchain specific contextual infos
}

const MachineManagerTeeMachineStatusChangedEventName = "TeeMachineStatusChanged"

// ContractEventName returns the user-defined event name.
func (MachineManagerTeeMachineStatusChanged) ContractEventName() string {
	return MachineManagerTeeMachineStatusChangedEventName
}

// UnpackTeeMachineStatusChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeMachineStatusChanged(address indexed teeId, uint8 indexed newStatus)
func (machineManager *MachineManager) UnpackTeeMachineStatusChangedEvent(log *types.Log) (*MachineManagerTeeMachineStatusChanged, error) {
	event := "TeeMachineStatusChanged"
	if len(log.Topics) == 0 || log.Topics[0] != machineManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachineManagerTeeMachineStatusChanged)
	if len(log.Data) > 0 {
		if err := machineManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machineManager.abi.Events[event].Inputs {
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
func (machineManager *MachineManager) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], machineManager.abi.Errors["AddressAlreadyInSet"].ID.Bytes()[:4]) {
		return machineManager.UnpackAddressAlreadyInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["AddressNotInSet"].ID.Bytes()[:4]) {
		return machineManager.UnpackAddressNotInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["AlreadyRegistered"].ID.Bytes()[:4]) {
		return machineManager.UnpackAlreadyRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["AvailabilityCheckTimestampInvalid"].ID.Bytes()[:4]) {
		return machineManager.UnpackAvailabilityCheckTimestampInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["ChallengeExpired"].ID.Bytes()[:4]) {
		return machineManager.UnpackChallengeExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["CosignersThresholdNotMet"].ID.Bytes()[:4]) {
		return machineManager.UnpackCosignersThresholdNotMetError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["CosignersThresholdTooHigh"].ID.Bytes()[:4]) {
		return machineManager.UnpackCosignersThresholdTooHighError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["DuplicatedCosigner"].ID.Bytes()[:4]) {
		return machineManager.UnpackDuplicatedCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return machineManager.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return machineManager.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return machineManager.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["EmergencyPauseActive"].ID.Bytes()[:4]) {
		return machineManager.UnpackEmergencyPauseActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["EmergencyProtectionActive"].ID.Bytes()[:4]) {
		return machineManager.UnpackEmergencyProtectionActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return machineManager.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["FeeTooLow"].ID.Bytes()[:4]) {
		return machineManager.UnpackFeeTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidAttestation"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidAttestationError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidKeyType"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidNewStatus"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidNewStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidPublicKey"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidRequestBody"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidRequestBodyError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidResponseData"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidResponseDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidResponseDataOrAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidResponseDataOrAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidSigningAlgo"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidSigningAlgoError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidSigningPolicy"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidSigningPolicyError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidTeeProxyId"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidTeeProxyIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidTeePublicKey"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidTeePublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidTeePublicKeyOrSignature"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidTeePublicKeyOrSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidTeeStatus"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidTeeStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidUrl"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidUrlError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["InvalidWalletStatus"].ID.Bytes()[:4]) {
		return machineManager.UnpackInvalidWalletStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["KeyTypeNotSupported"].ID.Bytes()[:4]) {
		return machineManager.UnpackKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return machineManager.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["MessageEmpty"].ID.Bytes()[:4]) {
		return machineManager.UnpackMessageEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["NoAddresses"].ID.Bytes()[:4]) {
		return machineManager.UnpackNoAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["NoTeeMachinesSpecified"].ID.Bytes()[:4]) {
		return machineManager.UnpackNoTeeMachinesSpecifiedError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["NotOwnerOrPauser"].ID.Bytes()[:4]) {
		return machineManager.UnpackNotOwnerOrPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["NotOwnerOrUnpauser"].ID.Bytes()[:4]) {
		return machineManager.UnpackNotOwnerOrUnpauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["OnlyExtensionOwner"].ID.Bytes()[:4]) {
		return machineManager.UnpackOnlyExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return machineManager.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return machineManager.UnpackOnlyOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["OnlyOwnerOrBackupManager"].ID.Bytes()[:4]) {
		return machineManager.UnpackOnlyOwnerOrBackupManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["OnlyOwnerOrExpiredAvailabilityCheck"].ID.Bytes()[:4]) {
		return machineManager.UnpackOnlyOwnerOrExpiredAvailabilityCheckError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return machineManager.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["OnlyProposedOwner"].ID.Bytes()[:4]) {
		return machineManager.UnpackOnlyProposedOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["OperationCommandEmpty"].ID.Bytes()[:4]) {
		return machineManager.UnpackOperationCommandEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["OperationTypeEmpty"].ID.Bytes()[:4]) {
		return machineManager.UnpackOperationTypeEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["OwnerMismatch"].ID.Bytes()[:4]) {
		return machineManager.UnpackOwnerMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["OwnerNotAllowed"].ID.Bytes()[:4]) {
		return machineManager.UnpackOwnerNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return machineManager.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["TeeNotFound"].ID.Bytes()[:4]) {
		return machineManager.UnpackTeeNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["TooMany"].ID.Bytes()[:4]) {
		return machineManager.UnpackTooManyError(raw[4:])
	}
	if bytes.Equal(raw[:4], machineManager.abi.Errors["VersionNotSupported"].ID.Bytes()[:4]) {
		return machineManager.UnpackVersionNotSupportedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// MachineManagerAddressAlreadyInSet represents a AddressAlreadyInSet error raised by the MachineManager contract.
type MachineManagerAddressAlreadyInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressAlreadyInSet(address addr)
func MachineManagerAddressAlreadyInSetErrorID() common.Hash {
	return common.HexToHash("0xb725c7fb1c3f0f065973d6b9eb55140f970259955af3045742cd5b37902eb920")
}

// UnpackAddressAlreadyInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressAlreadyInSet(address addr)
func (machineManager *MachineManager) UnpackAddressAlreadyInSetError(raw []byte) (*MachineManagerAddressAlreadyInSet, error) {
	out := new(MachineManagerAddressAlreadyInSet)
	if err := machineManager.abi.UnpackIntoInterface(out, "AddressAlreadyInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerAddressNotInSet represents a AddressNotInSet error raised by the MachineManager contract.
type MachineManagerAddressNotInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressNotInSet(address addr)
func MachineManagerAddressNotInSetErrorID() common.Hash {
	return common.HexToHash("0x1ff5ecf3e6d385ea01ff0323fe894ed6d25a83ed5f6ff2a2ede4c897408f6be9")
}

// UnpackAddressNotInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressNotInSet(address addr)
func (machineManager *MachineManager) UnpackAddressNotInSetError(raw []byte) (*MachineManagerAddressNotInSet, error) {
	out := new(MachineManagerAddressNotInSet)
	if err := machineManager.abi.UnpackIntoInterface(out, "AddressNotInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerAlreadyRegistered represents a AlreadyRegistered error raised by the MachineManager contract.
type MachineManagerAlreadyRegistered struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyRegistered()
func MachineManagerAlreadyRegisteredErrorID() common.Hash {
	return common.HexToHash("0x3a81d6fc79a5635b1f06f5006b0b65f8bc72a09a41893acedbabc02c660943ca")
}

// UnpackAlreadyRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyRegistered()
func (machineManager *MachineManager) UnpackAlreadyRegisteredError(raw []byte) (*MachineManagerAlreadyRegistered, error) {
	out := new(MachineManagerAlreadyRegistered)
	if err := machineManager.abi.UnpackIntoInterface(out, "AlreadyRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerAvailabilityCheckTimestampInvalid represents a AvailabilityCheckTimestampInvalid error raised by the MachineManager contract.
type MachineManagerAvailabilityCheckTimestampInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func MachineManagerAvailabilityCheckTimestampInvalidErrorID() common.Hash {
	return common.HexToHash("0x6344a5e74cff78612fa271fa907ca608af3bad10e1e8c8df302e29ebc6298e2d")
}

// UnpackAvailabilityCheckTimestampInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func (machineManager *MachineManager) UnpackAvailabilityCheckTimestampInvalidError(raw []byte) (*MachineManagerAvailabilityCheckTimestampInvalid, error) {
	out := new(MachineManagerAvailabilityCheckTimestampInvalid)
	if err := machineManager.abi.UnpackIntoInterface(out, "AvailabilityCheckTimestampInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerChallengeExpired represents a ChallengeExpired error raised by the MachineManager contract.
type MachineManagerChallengeExpired struct {
	ChallengeTs *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ChallengeExpired(uint256 challengeTs)
func MachineManagerChallengeExpiredErrorID() common.Hash {
	return common.HexToHash("0x5e804f6ab65f629e96873f89acb6e5d1456139dc56369bcfc528171e348baa42")
}

// UnpackChallengeExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ChallengeExpired(uint256 challengeTs)
func (machineManager *MachineManager) UnpackChallengeExpiredError(raw []byte) (*MachineManagerChallengeExpired, error) {
	out := new(MachineManagerChallengeExpired)
	if err := machineManager.abi.UnpackIntoInterface(out, "ChallengeExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerCosignersThresholdNotMet represents a CosignersThresholdNotMet error raised by the MachineManager contract.
type MachineManagerCosignersThresholdNotMet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CosignersThresholdNotMet()
func MachineManagerCosignersThresholdNotMetErrorID() common.Hash {
	return common.HexToHash("0xff41656670fff33a7dbbd13446e34a04caed47f2d1c8607b1828a6aa9d9051b3")
}

// UnpackCosignersThresholdNotMetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CosignersThresholdNotMet()
func (machineManager *MachineManager) UnpackCosignersThresholdNotMetError(raw []byte) (*MachineManagerCosignersThresholdNotMet, error) {
	out := new(MachineManagerCosignersThresholdNotMet)
	if err := machineManager.abi.UnpackIntoInterface(out, "CosignersThresholdNotMet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerCosignersThresholdTooHigh represents a CosignersThresholdTooHigh error raised by the MachineManager contract.
type MachineManagerCosignersThresholdTooHigh struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CosignersThresholdTooHigh()
func MachineManagerCosignersThresholdTooHighErrorID() common.Hash {
	return common.HexToHash("0x3d7053512cbb0b649dba01ac4e3591f104fd7d0957128a39a3c552c77cd5014d")
}

// UnpackCosignersThresholdTooHighError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CosignersThresholdTooHigh()
func (machineManager *MachineManager) UnpackCosignersThresholdTooHighError(raw []byte) (*MachineManagerCosignersThresholdTooHigh, error) {
	out := new(MachineManagerCosignersThresholdTooHigh)
	if err := machineManager.abi.UnpackIntoInterface(out, "CosignersThresholdTooHigh", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerDuplicatedCosigner represents a DuplicatedCosigner error raised by the MachineManager contract.
type MachineManagerDuplicatedCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func MachineManagerDuplicatedCosignerErrorID() common.Hash {
	return common.HexToHash("0x53a5890417576fb3043dc3ff1669b0932da2f657d8adb4c11313b12b711f5910")
}

// UnpackDuplicatedCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func (machineManager *MachineManager) UnpackDuplicatedCosignerError(raw []byte) (*MachineManagerDuplicatedCosigner, error) {
	out := new(MachineManagerDuplicatedCosigner)
	if err := machineManager.abi.UnpackIntoInterface(out, "DuplicatedCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the MachineManager contract.
type MachineManagerECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func MachineManagerECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (machineManager *MachineManager) UnpackECDSAInvalidSignatureError(raw []byte) (*MachineManagerECDSAInvalidSignature, error) {
	out := new(MachineManagerECDSAInvalidSignature)
	if err := machineManager.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the MachineManager contract.
type MachineManagerECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func MachineManagerECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (machineManager *MachineManager) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*MachineManagerECDSAInvalidSignatureLength, error) {
	out := new(MachineManagerECDSAInvalidSignatureLength)
	if err := machineManager.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the MachineManager contract.
type MachineManagerECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func MachineManagerECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (machineManager *MachineManager) UnpackECDSAInvalidSignatureSError(raw []byte) (*MachineManagerECDSAInvalidSignatureS, error) {
	out := new(MachineManagerECDSAInvalidSignatureS)
	if err := machineManager.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerEmergencyPauseActive represents a EmergencyPauseActive error raised by the MachineManager contract.
type MachineManagerEmergencyPauseActive struct {
	ExtensionId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func MachineManagerEmergencyPauseActiveErrorID() common.Hash {
	return common.HexToHash("0x480cf2163927c55f46b05a5cac80ac3781076b0e50b1ed6b07995a800937d4b6")
}

// UnpackEmergencyPauseActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func (machineManager *MachineManager) UnpackEmergencyPauseActiveError(raw []byte) (*MachineManagerEmergencyPauseActive, error) {
	out := new(MachineManagerEmergencyPauseActive)
	if err := machineManager.abi.UnpackIntoInterface(out, "EmergencyPauseActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerEmergencyProtectionActive represents a EmergencyProtectionActive error raised by the MachineManager contract.
type MachineManagerEmergencyProtectionActive struct {
	ExtensionId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmergencyProtectionActive(uint256 extensionId)
func MachineManagerEmergencyProtectionActiveErrorID() common.Hash {
	return common.HexToHash("0x93136b5bc178e37d8a37403ffe4d8ca41116b64440923a070fa269ad41a394bf")
}

// UnpackEmergencyProtectionActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmergencyProtectionActive(uint256 extensionId)
func (machineManager *MachineManager) UnpackEmergencyProtectionActiveError(raw []byte) (*MachineManagerEmergencyProtectionActive, error) {
	out := new(MachineManagerEmergencyProtectionActive)
	if err := machineManager.abi.UnpackIntoInterface(out, "EmergencyProtectionActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerExtensionIdMismatch represents a ExtensionIdMismatch error raised by the MachineManager contract.
type MachineManagerExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func MachineManagerExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (machineManager *MachineManager) UnpackExtensionIdMismatchError(raw []byte) (*MachineManagerExtensionIdMismatch, error) {
	out := new(MachineManagerExtensionIdMismatch)
	if err := machineManager.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerFeeTooLow represents a FeeTooLow error raised by the MachineManager contract.
type MachineManagerFeeTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeTooLow()
func MachineManagerFeeTooLowErrorID() common.Hash {
	return common.HexToHash("0x732f94137528ce257864c792a6de53ff98f1203051e0a27c454063f559602458")
}

// UnpackFeeTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeTooLow()
func (machineManager *MachineManager) UnpackFeeTooLowError(raw []byte) (*MachineManagerFeeTooLow, error) {
	out := new(MachineManagerFeeTooLow)
	if err := machineManager.abi.UnpackIntoInterface(out, "FeeTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidAddress represents a InvalidAddress error raised by the MachineManager contract.
type MachineManagerInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func MachineManagerInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (machineManager *MachineManager) UnpackInvalidAddressError(raw []byte) (*MachineManagerInvalidAddress, error) {
	out := new(MachineManagerInvalidAddress)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidAttestation represents a InvalidAttestation error raised by the MachineManager contract.
type MachineManagerInvalidAttestation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAttestation()
func MachineManagerInvalidAttestationErrorID() common.Hash {
	return common.HexToHash("0xbd8ba84d4af301e1b110d1f93f8971934a3ba3afdf4200e7f884f76f2dd27077")
}

// UnpackInvalidAttestationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAttestation()
func (machineManager *MachineManager) UnpackInvalidAttestationError(raw []byte) (*MachineManagerInvalidAttestation, error) {
	out := new(MachineManagerInvalidAttestation)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidAttestation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidAvailabilityCheckStatus represents a InvalidAvailabilityCheckStatus error raised by the MachineManager contract.
type MachineManagerInvalidAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func MachineManagerInvalidAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0xedee4ff00b842d3ef84031c60001065c540c838c5479bbf0ae56f15868427e44")
}

// UnpackInvalidAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func (machineManager *MachineManager) UnpackInvalidAvailabilityCheckStatusError(raw []byte) (*MachineManagerInvalidAvailabilityCheckStatus, error) {
	out := new(MachineManagerInvalidAvailabilityCheckStatus)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidCosigner represents a InvalidCosigner error raised by the MachineManager contract.
type MachineManagerInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func MachineManagerInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (machineManager *MachineManager) UnpackInvalidCosignerError(raw []byte) (*MachineManagerInvalidCosigner, error) {
	out := new(MachineManagerInvalidCosigner)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidDuration represents a InvalidDuration error raised by the MachineManager contract.
type MachineManagerInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func MachineManagerInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (machineManager *MachineManager) UnpackInvalidDurationError(raw []byte) (*MachineManagerInvalidDuration, error) {
	out := new(MachineManagerInvalidDuration)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the MachineManager contract.
type MachineManagerInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func MachineManagerInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (machineManager *MachineManager) UnpackInvalidGovernanceHashError(raw []byte) (*MachineManagerInvalidGovernanceHash, error) {
	out := new(MachineManagerInvalidGovernanceHash)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidKeyType represents a InvalidKeyType error raised by the MachineManager contract.
type MachineManagerInvalidKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyType()
func MachineManagerInvalidKeyTypeErrorID() common.Hash {
	return common.HexToHash("0x36dccb440beb343622d8ee0927bce3f69281feea3a5b0610f6e00e2811b557e1")
}

// UnpackInvalidKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyType()
func (machineManager *MachineManager) UnpackInvalidKeyTypeError(raw []byte) (*MachineManagerInvalidKeyType, error) {
	out := new(MachineManagerInvalidKeyType)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidNewStatus represents a InvalidNewStatus error raised by the MachineManager contract.
type MachineManagerInvalidNewStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNewStatus()
func MachineManagerInvalidNewStatusErrorID() common.Hash {
	return common.HexToHash("0x42ab00592290b3c07f555e63f1f1fffd50cfc2ab8256154885afe07541f9b78b")
}

// UnpackInvalidNewStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNewStatus()
func (machineManager *MachineManager) UnpackInvalidNewStatusError(raw []byte) (*MachineManagerInvalidNewStatus, error) {
	out := new(MachineManagerInvalidNewStatus)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidNewStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidNonce represents a InvalidNonce error raised by the MachineManager contract.
type MachineManagerInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func MachineManagerInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (machineManager *MachineManager) UnpackInvalidNonceError(raw []byte) (*MachineManagerInvalidNonce, error) {
	out := new(MachineManagerInvalidNonce)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidPublicKey represents a InvalidPublicKey error raised by the MachineManager contract.
type MachineManagerInvalidPublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPublicKey()
func MachineManagerInvalidPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa2d0fee854985cb3d3a394146b5df1c41d5b4d6fcede7fc57a3c05a49b2f3e73")
}

// UnpackInvalidPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPublicKey()
func (machineManager *MachineManager) UnpackInvalidPublicKeyError(raw []byte) (*MachineManagerInvalidPublicKey, error) {
	out := new(MachineManagerInvalidPublicKey)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidRequestBody represents a InvalidRequestBody error raised by the MachineManager contract.
type MachineManagerInvalidRequestBody struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRequestBody()
func MachineManagerInvalidRequestBodyErrorID() common.Hash {
	return common.HexToHash("0xfe524bcc2b70aad90ebe31fd595732afd3bed1387a6e1141569409a40554173c")
}

// UnpackInvalidRequestBodyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRequestBody()
func (machineManager *MachineManager) UnpackInvalidRequestBodyError(raw []byte) (*MachineManagerInvalidRequestBody, error) {
	out := new(MachineManagerInvalidRequestBody)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidRequestBody", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidResponseData represents a InvalidResponseData error raised by the MachineManager contract.
type MachineManagerInvalidResponseData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseData()
func MachineManagerInvalidResponseDataErrorID() common.Hash {
	return common.HexToHash("0x48566617128355a033880805d6462da9bf9c6fe4c61b922997451a71262990af")
}

// UnpackInvalidResponseDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseData()
func (machineManager *MachineManager) UnpackInvalidResponseDataError(raw []byte) (*MachineManagerInvalidResponseData, error) {
	out := new(MachineManagerInvalidResponseData)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidResponseData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidResponseDataOrAvailabilityCheckStatus represents a InvalidResponseDataOrAvailabilityCheckStatus error raised by the MachineManager contract.
type MachineManagerInvalidResponseDataOrAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseDataOrAvailabilityCheckStatus()
func MachineManagerInvalidResponseDataOrAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0x1cc301de281a1519ee1feb7ab01ed8bb0b921df792e8526e446c68a9b08a0288")
}

// UnpackInvalidResponseDataOrAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseDataOrAvailabilityCheckStatus()
func (machineManager *MachineManager) UnpackInvalidResponseDataOrAvailabilityCheckStatusError(raw []byte) (*MachineManagerInvalidResponseDataOrAvailabilityCheckStatus, error) {
	out := new(MachineManagerInvalidResponseDataOrAvailabilityCheckStatus)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidResponseDataOrAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidSigningAlgo represents a InvalidSigningAlgo error raised by the MachineManager contract.
type MachineManagerInvalidSigningAlgo struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningAlgo()
func MachineManagerInvalidSigningAlgoErrorID() common.Hash {
	return common.HexToHash("0x9b7198a697b61e546ef9ee593a6f77d2bf9f1df9dcbba655ce96d7e2c75346c7")
}

// UnpackInvalidSigningAlgoError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningAlgo()
func (machineManager *MachineManager) UnpackInvalidSigningAlgoError(raw []byte) (*MachineManagerInvalidSigningAlgo, error) {
	out := new(MachineManagerInvalidSigningAlgo)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidSigningAlgo", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidSigningPolicy represents a InvalidSigningPolicy error raised by the MachineManager contract.
type MachineManagerInvalidSigningPolicy struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningPolicy()
func MachineManagerInvalidSigningPolicyErrorID() common.Hash {
	return common.HexToHash("0x3d7cc620812b32543e78b18734710415246f44f2e8d99c9198a7cb7573499dda")
}

// UnpackInvalidSigningPolicyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningPolicy()
func (machineManager *MachineManager) UnpackInvalidSigningPolicyError(raw []byte) (*MachineManagerInvalidSigningPolicy, error) {
	out := new(MachineManagerInvalidSigningPolicy)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidSigningPolicy", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidTeeProxyId represents a InvalidTeeProxyId error raised by the MachineManager contract.
type MachineManagerInvalidTeeProxyId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidTeeProxyId()
func MachineManagerInvalidTeeProxyIdErrorID() common.Hash {
	return common.HexToHash("0xd8aca4a05deb49e8a61120b4a51af7a64cee367d06a7b519e3cd7f8362b3181a")
}

// UnpackInvalidTeeProxyIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidTeeProxyId()
func (machineManager *MachineManager) UnpackInvalidTeeProxyIdError(raw []byte) (*MachineManagerInvalidTeeProxyId, error) {
	out := new(MachineManagerInvalidTeeProxyId)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidTeeProxyId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidTeePublicKey represents a InvalidTeePublicKey error raised by the MachineManager contract.
type MachineManagerInvalidTeePublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidTeePublicKey()
func MachineManagerInvalidTeePublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa602bd8acedfcf0cbb7e9d954c2e86d3123f1adff5f7e3e920336a33d04a392e")
}

// UnpackInvalidTeePublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidTeePublicKey()
func (machineManager *MachineManager) UnpackInvalidTeePublicKeyError(raw []byte) (*MachineManagerInvalidTeePublicKey, error) {
	out := new(MachineManagerInvalidTeePublicKey)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidTeePublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidTeePublicKeyOrSignature represents a InvalidTeePublicKeyOrSignature error raised by the MachineManager contract.
type MachineManagerInvalidTeePublicKeyOrSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidTeePublicKeyOrSignature()
func MachineManagerInvalidTeePublicKeyOrSignatureErrorID() common.Hash {
	return common.HexToHash("0xa6f4fbdeae9232aabe4ca35be781de76c4d69a863d85d8a08cbbd472bcb938e0")
}

// UnpackInvalidTeePublicKeyOrSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidTeePublicKeyOrSignature()
func (machineManager *MachineManager) UnpackInvalidTeePublicKeyOrSignatureError(raw []byte) (*MachineManagerInvalidTeePublicKeyOrSignature, error) {
	out := new(MachineManagerInvalidTeePublicKeyOrSignature)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidTeePublicKeyOrSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidTeeStatus represents a InvalidTeeStatus error raised by the MachineManager contract.
type MachineManagerInvalidTeeStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidTeeStatus()
func MachineManagerInvalidTeeStatusErrorID() common.Hash {
	return common.HexToHash("0xdbc8b0c17d99454121c4999e401072aa57d8436c6da35989829d6f29414c232d")
}

// UnpackInvalidTeeStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidTeeStatus()
func (machineManager *MachineManager) UnpackInvalidTeeStatusError(raw []byte) (*MachineManagerInvalidTeeStatus, error) {
	out := new(MachineManagerInvalidTeeStatus)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidTeeStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidThreshold represents a InvalidThreshold error raised by the MachineManager contract.
type MachineManagerInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func MachineManagerInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (machineManager *MachineManager) UnpackInvalidThresholdError(raw []byte) (*MachineManagerInvalidThreshold, error) {
	out := new(MachineManagerInvalidThreshold)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidUrl represents a InvalidUrl error raised by the MachineManager contract.
type MachineManagerInvalidUrl struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidUrl()
func MachineManagerInvalidUrlErrorID() common.Hash {
	return common.HexToHash("0xe7bc872f355f40107fd25870b13981db57bc67caa988d2e8e204eb6db09b5477")
}

// UnpackInvalidUrlError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidUrl()
func (machineManager *MachineManager) UnpackInvalidUrlError(raw []byte) (*MachineManagerInvalidUrl, error) {
	out := new(MachineManagerInvalidUrl)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidUrl", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerInvalidWalletStatus represents a InvalidWalletStatus error raised by the MachineManager contract.
type MachineManagerInvalidWalletStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidWalletStatus()
func MachineManagerInvalidWalletStatusErrorID() common.Hash {
	return common.HexToHash("0xd700120e4896ea7835a5e9e4feb690644ced5cad1c52209afdcbe06b96d91345")
}

// UnpackInvalidWalletStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidWalletStatus()
func (machineManager *MachineManager) UnpackInvalidWalletStatusError(raw []byte) (*MachineManagerInvalidWalletStatus, error) {
	out := new(MachineManagerInvalidWalletStatus)
	if err := machineManager.abi.UnpackIntoInterface(out, "InvalidWalletStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerKeyTypeNotSupported represents a KeyTypeNotSupported error raised by the MachineManager contract.
type MachineManagerKeyTypeNotSupported struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func MachineManagerKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x85004df1f47814c2de6fbf03d3f2b1ca8c67c182d0d54278766543b6be3bc603")
}

// UnpackKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func (machineManager *MachineManager) UnpackKeyTypeNotSupportedError(raw []byte) (*MachineManagerKeyTypeNotSupported, error) {
	out := new(MachineManagerKeyTypeNotSupported)
	if err := machineManager.abi.UnpackIntoInterface(out, "KeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerLengthsMismatch represents a LengthsMismatch error raised by the MachineManager contract.
type MachineManagerLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func MachineManagerLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (machineManager *MachineManager) UnpackLengthsMismatchError(raw []byte) (*MachineManagerLengthsMismatch, error) {
	out := new(MachineManagerLengthsMismatch)
	if err := machineManager.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerMessageEmpty represents a MessageEmpty error raised by the MachineManager contract.
type MachineManagerMessageEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MessageEmpty()
func MachineManagerMessageEmptyErrorID() common.Hash {
	return common.HexToHash("0x3b06a1beedd2da56654bd6f78f50f74e9bc7469bae35c52fd3680e4d95103c79")
}

// UnpackMessageEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MessageEmpty()
func (machineManager *MachineManager) UnpackMessageEmptyError(raw []byte) (*MachineManagerMessageEmpty, error) {
	out := new(MachineManagerMessageEmpty)
	if err := machineManager.abi.UnpackIntoInterface(out, "MessageEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerNoAddresses represents a NoAddresses error raised by the MachineManager contract.
type MachineManagerNoAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoAddresses()
func MachineManagerNoAddressesErrorID() common.Hash {
	return common.HexToHash("0xb2e79f05746fa6c9f3d61f9efe85959371cde5030a6c2dd70c92b63578c5cde7")
}

// UnpackNoAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoAddresses()
func (machineManager *MachineManager) UnpackNoAddressesError(raw []byte) (*MachineManagerNoAddresses, error) {
	out := new(MachineManagerNoAddresses)
	if err := machineManager.abi.UnpackIntoInterface(out, "NoAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerNoTeeMachinesSpecified represents a NoTeeMachinesSpecified error raised by the MachineManager contract.
type MachineManagerNoTeeMachinesSpecified struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoTeeMachinesSpecified()
func MachineManagerNoTeeMachinesSpecifiedErrorID() common.Hash {
	return common.HexToHash("0xe10387d0118f08d51d80fe03b0c14ed458c46aac2ae9964e0b628978e6938db8")
}

// UnpackNoTeeMachinesSpecifiedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoTeeMachinesSpecified()
func (machineManager *MachineManager) UnpackNoTeeMachinesSpecifiedError(raw []byte) (*MachineManagerNoTeeMachinesSpecified, error) {
	out := new(MachineManagerNoTeeMachinesSpecified)
	if err := machineManager.abi.UnpackIntoInterface(out, "NoTeeMachinesSpecified", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerNotOwnerOrPauser represents a NotOwnerOrPauser error raised by the MachineManager contract.
type MachineManagerNotOwnerOrPauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrPauser(address caller)
func MachineManagerNotOwnerOrPauserErrorID() common.Hash {
	return common.HexToHash("0x46c0d9af6b895093e022ad2984c54e60192e6e26c5063daf232e825c6e07ff9a")
}

// UnpackNotOwnerOrPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrPauser(address caller)
func (machineManager *MachineManager) UnpackNotOwnerOrPauserError(raw []byte) (*MachineManagerNotOwnerOrPauser, error) {
	out := new(MachineManagerNotOwnerOrPauser)
	if err := machineManager.abi.UnpackIntoInterface(out, "NotOwnerOrPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerNotOwnerOrUnpauser represents a NotOwnerOrUnpauser error raised by the MachineManager contract.
type MachineManagerNotOwnerOrUnpauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func MachineManagerNotOwnerOrUnpauserErrorID() common.Hash {
	return common.HexToHash("0x81e700e8f0bd5ba4a6d6a2ce9a1ad0ac9bb1b1b1e62d921a55cbc4b53f1da649")
}

// UnpackNotOwnerOrUnpauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func (machineManager *MachineManager) UnpackNotOwnerOrUnpauserError(raw []byte) (*MachineManagerNotOwnerOrUnpauser, error) {
	out := new(MachineManagerNotOwnerOrUnpauser)
	if err := machineManager.abi.UnpackIntoInterface(out, "NotOwnerOrUnpauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerOnlyExtensionOwner represents a OnlyExtensionOwner error raised by the MachineManager contract.
type MachineManagerOnlyExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwner()
func MachineManagerOnlyExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xcd49fd1de5a297cd1ce055d5abffcf2cb25ad9492bd39a703a238af57b2a9488")
}

// UnpackOnlyExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwner()
func (machineManager *MachineManager) UnpackOnlyExtensionOwnerError(raw []byte) (*MachineManagerOnlyExtensionOwner, error) {
	out := new(MachineManagerOnlyExtensionOwner)
	if err := machineManager.abi.UnpackIntoInterface(out, "OnlyExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the MachineManager contract.
type MachineManagerOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func MachineManagerOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (machineManager *MachineManager) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*MachineManagerOnlyExtensionOwnerOrOperator, error) {
	out := new(MachineManagerOnlyExtensionOwnerOrOperator)
	if err := machineManager.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerOnlyOwner represents a OnlyOwner error raised by the MachineManager contract.
type MachineManagerOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func MachineManagerOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (machineManager *MachineManager) UnpackOnlyOwnerError(raw []byte) (*MachineManagerOnlyOwner, error) {
	out := new(MachineManagerOnlyOwner)
	if err := machineManager.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerOnlyOwnerOrBackupManager represents a OnlyOwnerOrBackupManager error raised by the MachineManager contract.
type MachineManagerOnlyOwnerOrBackupManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrBackupManager()
func MachineManagerOnlyOwnerOrBackupManagerErrorID() common.Hash {
	return common.HexToHash("0x249a6cf6e22092f9c80798e2a0c583c0f36aea0779ef38a5b5cdb6a112a0e6df")
}

// UnpackOnlyOwnerOrBackupManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrBackupManager()
func (machineManager *MachineManager) UnpackOnlyOwnerOrBackupManagerError(raw []byte) (*MachineManagerOnlyOwnerOrBackupManager, error) {
	out := new(MachineManagerOnlyOwnerOrBackupManager)
	if err := machineManager.abi.UnpackIntoInterface(out, "OnlyOwnerOrBackupManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerOnlyOwnerOrExpiredAvailabilityCheck represents a OnlyOwnerOrExpiredAvailabilityCheck error raised by the MachineManager contract.
type MachineManagerOnlyOwnerOrExpiredAvailabilityCheck struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrExpiredAvailabilityCheck()
func MachineManagerOnlyOwnerOrExpiredAvailabilityCheckErrorID() common.Hash {
	return common.HexToHash("0x43cc8a7918d4d7d7208562a0c5939b1374ce32a9b7402d4ca659adfe6e72d7e0")
}

// UnpackOnlyOwnerOrExpiredAvailabilityCheckError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrExpiredAvailabilityCheck()
func (machineManager *MachineManager) UnpackOnlyOwnerOrExpiredAvailabilityCheckError(raw []byte) (*MachineManagerOnlyOwnerOrExpiredAvailabilityCheck, error) {
	out := new(MachineManagerOnlyOwnerOrExpiredAvailabilityCheck)
	if err := machineManager.abi.UnpackIntoInterface(out, "OnlyOwnerOrExpiredAvailabilityCheck", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the MachineManager contract.
type MachineManagerOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func MachineManagerOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (machineManager *MachineManager) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*MachineManagerOnlyProductionOrPausedStatus, error) {
	out := new(MachineManagerOnlyProductionOrPausedStatus)
	if err := machineManager.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerOnlyProposedOwner represents a OnlyProposedOwner error raised by the MachineManager contract.
type MachineManagerOnlyProposedOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProposedOwner()
func MachineManagerOnlyProposedOwnerErrorID() common.Hash {
	return common.HexToHash("0x49cb15061ee558e74f66f937f740fd2a3791a89d2d139862222e5c97cabc02cd")
}

// UnpackOnlyProposedOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProposedOwner()
func (machineManager *MachineManager) UnpackOnlyProposedOwnerError(raw []byte) (*MachineManagerOnlyProposedOwner, error) {
	out := new(MachineManagerOnlyProposedOwner)
	if err := machineManager.abi.UnpackIntoInterface(out, "OnlyProposedOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerOperationCommandEmpty represents a OperationCommandEmpty error raised by the MachineManager contract.
type MachineManagerOperationCommandEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationCommandEmpty()
func MachineManagerOperationCommandEmptyErrorID() common.Hash {
	return common.HexToHash("0x038dc5a8067401ab291233391b3d11ff3ce1b21a8aebcf4297e6a9f759f65846")
}

// UnpackOperationCommandEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationCommandEmpty()
func (machineManager *MachineManager) UnpackOperationCommandEmptyError(raw []byte) (*MachineManagerOperationCommandEmpty, error) {
	out := new(MachineManagerOperationCommandEmpty)
	if err := machineManager.abi.UnpackIntoInterface(out, "OperationCommandEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerOperationTypeEmpty represents a OperationTypeEmpty error raised by the MachineManager contract.
type MachineManagerOperationTypeEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationTypeEmpty()
func MachineManagerOperationTypeEmptyErrorID() common.Hash {
	return common.HexToHash("0x06e3960d8a83ea8972d826357113f286ce60947bac56680f119fca98a2d6125b")
}

// UnpackOperationTypeEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationTypeEmpty()
func (machineManager *MachineManager) UnpackOperationTypeEmptyError(raw []byte) (*MachineManagerOperationTypeEmpty, error) {
	out := new(MachineManagerOperationTypeEmpty)
	if err := machineManager.abi.UnpackIntoInterface(out, "OperationTypeEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerOwnerMismatch represents a OwnerMismatch error raised by the MachineManager contract.
type MachineManagerOwnerMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerMismatch()
func MachineManagerOwnerMismatchErrorID() common.Hash {
	return common.HexToHash("0xa8c81623c6a24998fb51aeb63135226db1bb39f681e9395555b26f3627f0e744")
}

// UnpackOwnerMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerMismatch()
func (machineManager *MachineManager) UnpackOwnerMismatchError(raw []byte) (*MachineManagerOwnerMismatch, error) {
	out := new(MachineManagerOwnerMismatch)
	if err := machineManager.abi.UnpackIntoInterface(out, "OwnerMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerOwnerNotAllowed represents a OwnerNotAllowed error raised by the MachineManager contract.
type MachineManagerOwnerNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerNotAllowed()
func MachineManagerOwnerNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x3d5c9996545c6bf4287c26c1c2d8e6b4c39b92f19ba706c3a756720bf9c4fd24")
}

// UnpackOwnerNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerNotAllowed()
func (machineManager *MachineManager) UnpackOwnerNotAllowedError(raw []byte) (*MachineManagerOwnerNotAllowed, error) {
	out := new(MachineManagerOwnerNotAllowed)
	if err := machineManager.abi.UnpackIntoInterface(out, "OwnerNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the MachineManager contract.
type MachineManagerTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func MachineManagerTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (machineManager *MachineManager) UnpackTeeMachineNotAvailableError(raw []byte) (*MachineManagerTeeMachineNotAvailable, error) {
	out := new(MachineManagerTeeMachineNotAvailable)
	if err := machineManager.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerTeeNotFound represents a TeeNotFound error raised by the MachineManager contract.
type MachineManagerTeeNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeNotFound()
func MachineManagerTeeNotFoundErrorID() common.Hash {
	return common.HexToHash("0xceb05b6852ed4c7dc3a718eb0d26555398a76877a44dd41dc92c5fe7e6ef8c2d")
}

// UnpackTeeNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeNotFound()
func (machineManager *MachineManager) UnpackTeeNotFoundError(raw []byte) (*MachineManagerTeeNotFound, error) {
	out := new(MachineManagerTeeNotFound)
	if err := machineManager.abi.UnpackIntoInterface(out, "TeeNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerTooMany represents a TooMany error raised by the MachineManager contract.
type MachineManagerTooMany struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TooMany()
func MachineManagerTooManyErrorID() common.Hash {
	return common.HexToHash("0xd65ac61eaed3332b3b71dfc563e4a76b8b5917c5e456ad30c9e7669c992d1de3")
}

// UnpackTooManyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TooMany()
func (machineManager *MachineManager) UnpackTooManyError(raw []byte) (*MachineManagerTooMany, error) {
	out := new(MachineManagerTooMany)
	if err := machineManager.abi.UnpackIntoInterface(out, "TooMany", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachineManagerVersionNotSupported represents a VersionNotSupported error raised by the MachineManager contract.
type MachineManagerVersionNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionNotSupported()
func MachineManagerVersionNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x6db0b17b06266b37e721024a8ba1221439799ebb779dfd397865e37250abc9a8")
}

// UnpackVersionNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionNotSupported()
func (machineManager *MachineManager) UnpackVersionNotSupportedError(raw []byte) (*MachineManagerVersionNotSupported, error) {
	out := new(MachineManagerVersionNotSupported)
	if err := machineManager.abi.UnpackIntoInterface(out, "VersionNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}
