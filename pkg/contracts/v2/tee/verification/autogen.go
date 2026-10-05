// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package verification

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

// VerificationMetaData contains all meta data concerning the Verification contract.
var VerificationMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengeTs\",\"type\":\"uint256\"}],\"name\":\"ChallengeExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"EmergencyPauseActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestBody\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningPolicy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTs\",\"type\":\"uint256\"}],\"name\":\"AvailabilityCheckValidityExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"CosignersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"availabilityCheckValidityDurationSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"signingPolicyValidityDurationInRewardEpochs\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"challengeValidityDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"SettingsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"name\":\"TeeAttestationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"confirmAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getAvailabilityCheckValidity\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_endTs\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_lastSigningPolicyId\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCosigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSettings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_availabilityCheckValidityDurationSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengeValidityDurationSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_testOnTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proofOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestAvailabilityCheckAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestTeeAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"setCosigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_availabilityCheckValidityDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicyValidityDurationInRewardEpochs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_challengeValidityDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"updateSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Verification",
}

// Verification is an auto generated Go binding around an Ethereum contract.
type Verification struct {
	abi abi.ABI
}

// NewVerification creates a new instance of Verification.
func NewVerification() *Verification {
	parsed, err := VerificationMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Verification{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Verification) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConfirmAvailability is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f78297f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmAvailability(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (verification *Verification) PackConfirmAvailability(proof ITeeAvailabilityCheckProof) []byte {
	enc, err := verification.abi.Pack("confirmAvailability", proof)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmAvailability is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f78297f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmAvailability(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,address,string,bytes32,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (verification *Verification) TryPackConfirmAvailability(proof ITeeAvailabilityCheckProof) ([]byte, error) {
	return verification.abi.Pack("confirmAvailability", proof)
}

// PackGetAvailabilityCheckValidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e1f22f5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAvailabilityCheckValidity(address _teeId) view returns(uint64 _endTs, uint32 _lastSigningPolicyId)
func (verification *Verification) PackGetAvailabilityCheckValidity(teeId common.Address) []byte {
	enc, err := verification.abi.Pack("getAvailabilityCheckValidity", teeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAvailabilityCheckValidity is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e1f22f5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAvailabilityCheckValidity(address _teeId) view returns(uint64 _endTs, uint32 _lastSigningPolicyId)
func (verification *Verification) TryPackGetAvailabilityCheckValidity(teeId common.Address) ([]byte, error) {
	return verification.abi.Pack("getAvailabilityCheckValidity", teeId)
}

// GetAvailabilityCheckValidityOutput serves as a container for the return parameters of contract
// method GetAvailabilityCheckValidity.
type GetAvailabilityCheckValidityOutput struct {
	EndTs               uint64
	LastSigningPolicyId uint32
}

// UnpackGetAvailabilityCheckValidity is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e1f22f5.
//
// Solidity: function getAvailabilityCheckValidity(address _teeId) view returns(uint64 _endTs, uint32 _lastSigningPolicyId)
func (verification *Verification) UnpackGetAvailabilityCheckValidity(data []byte) (GetAvailabilityCheckValidityOutput, error) {
	out, err := verification.abi.Unpack("getAvailabilityCheckValidity", data)
	outstruct := new(GetAvailabilityCheckValidityOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.EndTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.LastSigningPolicyId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	return *outstruct, nil
}

// PackGetCosigners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x86b8d27e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCosigners() view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (verification *Verification) PackGetCosigners() []byte {
	enc, err := verification.abi.Pack("getCosigners")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCosigners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x86b8d27e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCosigners() view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (verification *Verification) TryPackGetCosigners() ([]byte, error) {
	return verification.abi.Pack("getCosigners")
}

// GetCosignersOutput serves as a container for the return parameters of contract
// method GetCosigners.
type GetCosignersOutput struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}

// UnpackGetCosigners is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x86b8d27e.
//
// Solidity: function getCosigners() view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (verification *Verification) UnpackGetCosigners(data []byte) (GetCosignersOutput, error) {
	out, err := verification.abi.Unpack("getCosigners", data)
	outstruct := new(GetCosignersOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Cosigners = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.CosignersThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85b4bb53.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSettings() view returns(uint256 _availabilityCheckValidityDurationSeconds, uint256 _challengeValidityDurationSeconds)
func (verification *Verification) PackGetSettings() []byte {
	enc, err := verification.abi.Pack("getSettings")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85b4bb53.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSettings() view returns(uint256 _availabilityCheckValidityDurationSeconds, uint256 _challengeValidityDurationSeconds)
func (verification *Verification) TryPackGetSettings() ([]byte, error) {
	return verification.abi.Pack("getSettings")
}

// GetSettingsOutput serves as a container for the return parameters of contract
// method GetSettings.
type GetSettingsOutput struct {
	AvailabilityCheckValidityDurationSeconds *big.Int
	ChallengeValidityDurationSeconds         *big.Int
}

// UnpackGetSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x85b4bb53.
//
// Solidity: function getSettings() view returns(uint256 _availabilityCheckValidityDurationSeconds, uint256 _challengeValidityDurationSeconds)
func (verification *Verification) UnpackGetSettings(data []byte) (GetSettingsOutput, error) {
	out, err := verification.abi.Unpack("getSettings", data)
	outstruct := new(GetSettingsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AvailabilityCheckValidityDurationSeconds = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.ChallengeValidityDurationSeconds = abi.ConvertType(out[1], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackRequestAvailabilityCheckAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3d60348.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function requestAvailabilityCheckAttestation(address _teeId, bytes32 _instructionId, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (verification *Verification) PackRequestAvailabilityCheckAttestation(teeId common.Address, instructionId [32]byte, testOnTeeId common.Address, proofOwner common.Address, claimBackAddress common.Address) []byte {
	enc, err := verification.abi.Pack("requestAvailabilityCheckAttestation", teeId, instructionId, testOnTeeId, proofOwner, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRequestAvailabilityCheckAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3d60348.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function requestAvailabilityCheckAttestation(address _teeId, bytes32 _instructionId, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (verification *Verification) TryPackRequestAvailabilityCheckAttestation(teeId common.Address, instructionId [32]byte, testOnTeeId common.Address, proofOwner common.Address, claimBackAddress common.Address) ([]byte, error) {
	return verification.abi.Pack("requestAvailabilityCheckAttestation", teeId, instructionId, testOnTeeId, proofOwner, claimBackAddress)
}

// PackRequestTeeAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d8dcfb5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function requestTeeAttestation(address _teeId, address _claimBackAddress) payable returns()
func (verification *Verification) PackRequestTeeAttestation(teeId common.Address, claimBackAddress common.Address) []byte {
	enc, err := verification.abi.Pack("requestTeeAttestation", teeId, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRequestTeeAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d8dcfb5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function requestTeeAttestation(address _teeId, address _claimBackAddress) payable returns()
func (verification *Verification) TryPackRequestTeeAttestation(teeId common.Address, claimBackAddress common.Address) ([]byte, error) {
	return verification.abi.Pack("requestTeeAttestation", teeId, claimBackAddress)
}

// PackSetCosigners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf0781728.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setCosigners(address[] _cosigners, uint64 _cosignersThreshold) returns()
func (verification *Verification) PackSetCosigners(cosigners []common.Address, cosignersThreshold uint64) []byte {
	enc, err := verification.abi.Pack("setCosigners", cosigners, cosignersThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetCosigners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf0781728.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setCosigners(address[] _cosigners, uint64 _cosignersThreshold) returns()
func (verification *Verification) TryPackSetCosigners(cosigners []common.Address, cosignersThreshold uint64) ([]byte, error) {
	return verification.abi.Pack("setCosigners", cosigners, cosignersThreshold)
}

// PackUpdateSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x34ef6f09.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateSettings(uint64 _availabilityCheckValidityDurationSeconds, uint64 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds) returns()
func (verification *Verification) PackUpdateSettings(availabilityCheckValidityDurationSeconds uint64, signingPolicyValidityDurationInRewardEpochs uint64, challengeValidityDurationSeconds uint64) []byte {
	enc, err := verification.abi.Pack("updateSettings", availabilityCheckValidityDurationSeconds, signingPolicyValidityDurationInRewardEpochs, challengeValidityDurationSeconds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x34ef6f09.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateSettings(uint64 _availabilityCheckValidityDurationSeconds, uint64 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds) returns()
func (verification *Verification) TryPackUpdateSettings(availabilityCheckValidityDurationSeconds uint64, signingPolicyValidityDurationInRewardEpochs uint64, challengeValidityDurationSeconds uint64) ([]byte, error) {
	return verification.abi.Pack("updateSettings", availabilityCheckValidityDurationSeconds, signingPolicyValidityDurationInRewardEpochs, challengeValidityDurationSeconds)
}

// VerificationAvailabilityCheckValidityExtended represents a AvailabilityCheckValidityExtended event raised by the Verification contract.
type VerificationAvailabilityCheckValidityExtended struct {
	TeeId common.Address
	Owner common.Address
	EndTs *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const VerificationAvailabilityCheckValidityExtendedEventName = "AvailabilityCheckValidityExtended"

// ContractEventName returns the user-defined event name.
func (VerificationAvailabilityCheckValidityExtended) ContractEventName() string {
	return VerificationAvailabilityCheckValidityExtendedEventName
}

// UnpackAvailabilityCheckValidityExtendedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AvailabilityCheckValidityExtended(address indexed teeId, address indexed owner, uint256 endTs)
func (verification *Verification) UnpackAvailabilityCheckValidityExtendedEvent(log *types.Log) (*VerificationAvailabilityCheckValidityExtended, error) {
	event := "AvailabilityCheckValidityExtended"
	if len(log.Topics) == 0 || log.Topics[0] != verification.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VerificationAvailabilityCheckValidityExtended)
	if len(log.Data) > 0 {
		if err := verification.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verification.abi.Events[event].Inputs {
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

// VerificationCosignersSet represents a CosignersSet event raised by the Verification contract.
type VerificationCosignersSet struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
	Raw                *types.Log // Blockchain specific contextual infos
}

const VerificationCosignersSetEventName = "CosignersSet"

// ContractEventName returns the user-defined event name.
func (VerificationCosignersSet) ContractEventName() string {
	return VerificationCosignersSetEventName
}

// UnpackCosignersSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CosignersSet(address[] cosigners, uint64 cosignersThreshold)
func (verification *Verification) UnpackCosignersSetEvent(log *types.Log) (*VerificationCosignersSet, error) {
	event := "CosignersSet"
	if len(log.Topics) == 0 || log.Topics[0] != verification.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VerificationCosignersSet)
	if len(log.Data) > 0 {
		if err := verification.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verification.abi.Events[event].Inputs {
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

// VerificationGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Verification contract.
type VerificationGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const VerificationGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (VerificationGovernanceCallTimelocked) ContractEventName() string {
	return VerificationGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (verification *Verification) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*VerificationGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != verification.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VerificationGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := verification.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verification.abi.Events[event].Inputs {
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

// VerificationInitialized represents a Initialized event raised by the Verification contract.
type VerificationInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const VerificationInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (VerificationInitialized) ContractEventName() string {
	return VerificationInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (verification *Verification) UnpackInitializedEvent(log *types.Log) (*VerificationInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != verification.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VerificationInitialized)
	if len(log.Data) > 0 {
		if err := verification.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verification.abi.Events[event].Inputs {
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

// VerificationSettingsUpdated represents a SettingsUpdated event raised by the Verification contract.
type VerificationSettingsUpdated struct {
	AvailabilityCheckValidityDurationSeconds    uint64
	SigningPolicyValidityDurationInRewardEpochs uint64
	ChallengeValidityDurationSeconds            uint64
	Raw                                         *types.Log // Blockchain specific contextual infos
}

const VerificationSettingsUpdatedEventName = "SettingsUpdated"

// ContractEventName returns the user-defined event name.
func (VerificationSettingsUpdated) ContractEventName() string {
	return VerificationSettingsUpdatedEventName
}

// UnpackSettingsUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SettingsUpdated(uint64 availabilityCheckValidityDurationSeconds, uint64 signingPolicyValidityDurationInRewardEpochs, uint64 challengeValidityDurationSeconds)
func (verification *Verification) UnpackSettingsUpdatedEvent(log *types.Log) (*VerificationSettingsUpdated, error) {
	event := "SettingsUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != verification.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VerificationSettingsUpdated)
	if len(log.Data) > 0 {
		if err := verification.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verification.abi.Events[event].Inputs {
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

// VerificationTeeAttestationRequested represents a TeeAttestationRequested event raised by the Verification contract.
type VerificationTeeAttestationRequested struct {
	TeeId     common.Address
	Challenge [32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const VerificationTeeAttestationRequestedEventName = "TeeAttestationRequested"

// ContractEventName returns the user-defined event name.
func (VerificationTeeAttestationRequested) ContractEventName() string {
	return VerificationTeeAttestationRequestedEventName
}

// UnpackTeeAttestationRequestedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeAttestationRequested(address indexed teeId, bytes32 challenge)
func (verification *Verification) UnpackTeeAttestationRequestedEvent(log *types.Log) (*VerificationTeeAttestationRequested, error) {
	event := "TeeAttestationRequested"
	if len(log.Topics) == 0 || log.Topics[0] != verification.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VerificationTeeAttestationRequested)
	if len(log.Data) > 0 {
		if err := verification.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verification.abi.Events[event].Inputs {
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

// VerificationTeeInstructionsSent represents a TeeInstructionsSent event raised by the Verification contract.
type VerificationTeeInstructionsSent struct {
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

const VerificationTeeInstructionsSentEventName = "TeeInstructionsSent"

// ContractEventName returns the user-defined event name.
func (VerificationTeeInstructionsSent) ContractEventName() string {
	return VerificationTeeInstructionsSentEventName
}

// UnpackTeeInstructionsSentEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (verification *Verification) UnpackTeeInstructionsSentEvent(log *types.Log) (*VerificationTeeInstructionsSent, error) {
	event := "TeeInstructionsSent"
	if len(log.Topics) == 0 || log.Topics[0] != verification.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VerificationTeeInstructionsSent)
	if len(log.Data) > 0 {
		if err := verification.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range verification.abi.Events[event].Inputs {
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
func (verification *Verification) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], verification.abi.Errors["AddressAlreadyInSet"].ID.Bytes()[:4]) {
		return verification.UnpackAddressAlreadyInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["AddressNotInSet"].ID.Bytes()[:4]) {
		return verification.UnpackAddressNotInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["AvailabilityCheckTimestampInvalid"].ID.Bytes()[:4]) {
		return verification.UnpackAvailabilityCheckTimestampInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["ChallengeExpired"].ID.Bytes()[:4]) {
		return verification.UnpackChallengeExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["CosignersThresholdNotMet"].ID.Bytes()[:4]) {
		return verification.UnpackCosignersThresholdNotMetError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["CosignersThresholdTooHigh"].ID.Bytes()[:4]) {
		return verification.UnpackCosignersThresholdTooHighError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["DuplicatedCosigner"].ID.Bytes()[:4]) {
		return verification.UnpackDuplicatedCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["EmergencyPauseActive"].ID.Bytes()[:4]) {
		return verification.UnpackEmergencyPauseActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return verification.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["FeeTooLow"].ID.Bytes()[:4]) {
		return verification.UnpackFeeTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidAttestation"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidAttestationError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidKeyType"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidPublicKey"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidRequestBody"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidRequestBodyError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidResponseData"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidResponseDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidSigningAlgo"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidSigningAlgoError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidSigningPolicy"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidSigningPolicyError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["InvalidWalletStatus"].ID.Bytes()[:4]) {
		return verification.UnpackInvalidWalletStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["KeyTypeNotSupported"].ID.Bytes()[:4]) {
		return verification.UnpackKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return verification.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["MessageEmpty"].ID.Bytes()[:4]) {
		return verification.UnpackMessageEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["NoAddresses"].ID.Bytes()[:4]) {
		return verification.UnpackNoAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["NoTeeMachinesSpecified"].ID.Bytes()[:4]) {
		return verification.UnpackNoTeeMachinesSpecifiedError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return verification.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["NotOwnerOrPauser"].ID.Bytes()[:4]) {
		return verification.UnpackNotOwnerOrPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["NotOwnerOrUnpauser"].ID.Bytes()[:4]) {
		return verification.UnpackNotOwnerOrUnpauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["OnlyExtensionOwner"].ID.Bytes()[:4]) {
		return verification.UnpackOnlyExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return verification.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return verification.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return verification.UnpackOnlyOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["OnlyOwnerOrBackupManager"].ID.Bytes()[:4]) {
		return verification.UnpackOnlyOwnerOrBackupManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return verification.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["OnlyProposedOwner"].ID.Bytes()[:4]) {
		return verification.UnpackOnlyProposedOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["OperationCommandEmpty"].ID.Bytes()[:4]) {
		return verification.UnpackOperationCommandEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["OperationTypeEmpty"].ID.Bytes()[:4]) {
		return verification.UnpackOperationTypeEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["OwnerNotAllowed"].ID.Bytes()[:4]) {
		return verification.UnpackOwnerNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return verification.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["TeeNotFound"].ID.Bytes()[:4]) {
		return verification.UnpackTeeNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], verification.abi.Errors["VersionNotSupported"].ID.Bytes()[:4]) {
		return verification.UnpackVersionNotSupportedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// VerificationAddressAlreadyInSet represents a AddressAlreadyInSet error raised by the Verification contract.
type VerificationAddressAlreadyInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressAlreadyInSet(address addr)
func VerificationAddressAlreadyInSetErrorID() common.Hash {
	return common.HexToHash("0xb725c7fb1c3f0f065973d6b9eb55140f970259955af3045742cd5b37902eb920")
}

// UnpackAddressAlreadyInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressAlreadyInSet(address addr)
func (verification *Verification) UnpackAddressAlreadyInSetError(raw []byte) (*VerificationAddressAlreadyInSet, error) {
	out := new(VerificationAddressAlreadyInSet)
	if err := verification.abi.UnpackIntoInterface(out, "AddressAlreadyInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationAddressNotInSet represents a AddressNotInSet error raised by the Verification contract.
type VerificationAddressNotInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressNotInSet(address addr)
func VerificationAddressNotInSetErrorID() common.Hash {
	return common.HexToHash("0x1ff5ecf3e6d385ea01ff0323fe894ed6d25a83ed5f6ff2a2ede4c897408f6be9")
}

// UnpackAddressNotInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressNotInSet(address addr)
func (verification *Verification) UnpackAddressNotInSetError(raw []byte) (*VerificationAddressNotInSet, error) {
	out := new(VerificationAddressNotInSet)
	if err := verification.abi.UnpackIntoInterface(out, "AddressNotInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationAvailabilityCheckTimestampInvalid represents a AvailabilityCheckTimestampInvalid error raised by the Verification contract.
type VerificationAvailabilityCheckTimestampInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func VerificationAvailabilityCheckTimestampInvalidErrorID() common.Hash {
	return common.HexToHash("0x6344a5e74cff78612fa271fa907ca608af3bad10e1e8c8df302e29ebc6298e2d")
}

// UnpackAvailabilityCheckTimestampInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func (verification *Verification) UnpackAvailabilityCheckTimestampInvalidError(raw []byte) (*VerificationAvailabilityCheckTimestampInvalid, error) {
	out := new(VerificationAvailabilityCheckTimestampInvalid)
	if err := verification.abi.UnpackIntoInterface(out, "AvailabilityCheckTimestampInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationChallengeExpired represents a ChallengeExpired error raised by the Verification contract.
type VerificationChallengeExpired struct {
	ChallengeTs *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ChallengeExpired(uint256 challengeTs)
func VerificationChallengeExpiredErrorID() common.Hash {
	return common.HexToHash("0x5e804f6ab65f629e96873f89acb6e5d1456139dc56369bcfc528171e348baa42")
}

// UnpackChallengeExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ChallengeExpired(uint256 challengeTs)
func (verification *Verification) UnpackChallengeExpiredError(raw []byte) (*VerificationChallengeExpired, error) {
	out := new(VerificationChallengeExpired)
	if err := verification.abi.UnpackIntoInterface(out, "ChallengeExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationCosignersThresholdNotMet represents a CosignersThresholdNotMet error raised by the Verification contract.
type VerificationCosignersThresholdNotMet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CosignersThresholdNotMet()
func VerificationCosignersThresholdNotMetErrorID() common.Hash {
	return common.HexToHash("0xff41656670fff33a7dbbd13446e34a04caed47f2d1c8607b1828a6aa9d9051b3")
}

// UnpackCosignersThresholdNotMetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CosignersThresholdNotMet()
func (verification *Verification) UnpackCosignersThresholdNotMetError(raw []byte) (*VerificationCosignersThresholdNotMet, error) {
	out := new(VerificationCosignersThresholdNotMet)
	if err := verification.abi.UnpackIntoInterface(out, "CosignersThresholdNotMet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationCosignersThresholdTooHigh represents a CosignersThresholdTooHigh error raised by the Verification contract.
type VerificationCosignersThresholdTooHigh struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CosignersThresholdTooHigh()
func VerificationCosignersThresholdTooHighErrorID() common.Hash {
	return common.HexToHash("0x3d7053512cbb0b649dba01ac4e3591f104fd7d0957128a39a3c552c77cd5014d")
}

// UnpackCosignersThresholdTooHighError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CosignersThresholdTooHigh()
func (verification *Verification) UnpackCosignersThresholdTooHighError(raw []byte) (*VerificationCosignersThresholdTooHigh, error) {
	out := new(VerificationCosignersThresholdTooHigh)
	if err := verification.abi.UnpackIntoInterface(out, "CosignersThresholdTooHigh", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationDuplicatedCosigner represents a DuplicatedCosigner error raised by the Verification contract.
type VerificationDuplicatedCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func VerificationDuplicatedCosignerErrorID() common.Hash {
	return common.HexToHash("0x53a5890417576fb3043dc3ff1669b0932da2f657d8adb4c11313b12b711f5910")
}

// UnpackDuplicatedCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func (verification *Verification) UnpackDuplicatedCosignerError(raw []byte) (*VerificationDuplicatedCosigner, error) {
	out := new(VerificationDuplicatedCosigner)
	if err := verification.abi.UnpackIntoInterface(out, "DuplicatedCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationEmergencyPauseActive represents a EmergencyPauseActive error raised by the Verification contract.
type VerificationEmergencyPauseActive struct {
	ExtensionId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func VerificationEmergencyPauseActiveErrorID() common.Hash {
	return common.HexToHash("0x480cf2163927c55f46b05a5cac80ac3781076b0e50b1ed6b07995a800937d4b6")
}

// UnpackEmergencyPauseActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func (verification *Verification) UnpackEmergencyPauseActiveError(raw []byte) (*VerificationEmergencyPauseActive, error) {
	out := new(VerificationEmergencyPauseActive)
	if err := verification.abi.UnpackIntoInterface(out, "EmergencyPauseActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationExtensionIdMismatch represents a ExtensionIdMismatch error raised by the Verification contract.
type VerificationExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func VerificationExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (verification *Verification) UnpackExtensionIdMismatchError(raw []byte) (*VerificationExtensionIdMismatch, error) {
	out := new(VerificationExtensionIdMismatch)
	if err := verification.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationFeeTooLow represents a FeeTooLow error raised by the Verification contract.
type VerificationFeeTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeTooLow()
func VerificationFeeTooLowErrorID() common.Hash {
	return common.HexToHash("0x732f94137528ce257864c792a6de53ff98f1203051e0a27c454063f559602458")
}

// UnpackFeeTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeTooLow()
func (verification *Verification) UnpackFeeTooLowError(raw []byte) (*VerificationFeeTooLow, error) {
	out := new(VerificationFeeTooLow)
	if err := verification.abi.UnpackIntoInterface(out, "FeeTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidAddress represents a InvalidAddress error raised by the Verification contract.
type VerificationInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func VerificationInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (verification *Verification) UnpackInvalidAddressError(raw []byte) (*VerificationInvalidAddress, error) {
	out := new(VerificationInvalidAddress)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidAttestation represents a InvalidAttestation error raised by the Verification contract.
type VerificationInvalidAttestation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAttestation()
func VerificationInvalidAttestationErrorID() common.Hash {
	return common.HexToHash("0xbd8ba84d4af301e1b110d1f93f8971934a3ba3afdf4200e7f884f76f2dd27077")
}

// UnpackInvalidAttestationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAttestation()
func (verification *Verification) UnpackInvalidAttestationError(raw []byte) (*VerificationInvalidAttestation, error) {
	out := new(VerificationInvalidAttestation)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidAttestation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidAvailabilityCheckStatus represents a InvalidAvailabilityCheckStatus error raised by the Verification contract.
type VerificationInvalidAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func VerificationInvalidAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0xedee4ff00b842d3ef84031c60001065c540c838c5479bbf0ae56f15868427e44")
}

// UnpackInvalidAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func (verification *Verification) UnpackInvalidAvailabilityCheckStatusError(raw []byte) (*VerificationInvalidAvailabilityCheckStatus, error) {
	out := new(VerificationInvalidAvailabilityCheckStatus)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidCosigner represents a InvalidCosigner error raised by the Verification contract.
type VerificationInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func VerificationInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (verification *Verification) UnpackInvalidCosignerError(raw []byte) (*VerificationInvalidCosigner, error) {
	out := new(VerificationInvalidCosigner)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidDuration represents a InvalidDuration error raised by the Verification contract.
type VerificationInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func VerificationInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (verification *Verification) UnpackInvalidDurationError(raw []byte) (*VerificationInvalidDuration, error) {
	out := new(VerificationInvalidDuration)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the Verification contract.
type VerificationInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func VerificationInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (verification *Verification) UnpackInvalidGovernanceHashError(raw []byte) (*VerificationInvalidGovernanceHash, error) {
	out := new(VerificationInvalidGovernanceHash)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidInitialization represents a InvalidInitialization error raised by the Verification contract.
type VerificationInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func VerificationInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (verification *Verification) UnpackInvalidInitializationError(raw []byte) (*VerificationInvalidInitialization, error) {
	out := new(VerificationInvalidInitialization)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidKeyType represents a InvalidKeyType error raised by the Verification contract.
type VerificationInvalidKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyType()
func VerificationInvalidKeyTypeErrorID() common.Hash {
	return common.HexToHash("0x36dccb440beb343622d8ee0927bce3f69281feea3a5b0610f6e00e2811b557e1")
}

// UnpackInvalidKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyType()
func (verification *Verification) UnpackInvalidKeyTypeError(raw []byte) (*VerificationInvalidKeyType, error) {
	out := new(VerificationInvalidKeyType)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidNonce represents a InvalidNonce error raised by the Verification contract.
type VerificationInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func VerificationInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (verification *Verification) UnpackInvalidNonceError(raw []byte) (*VerificationInvalidNonce, error) {
	out := new(VerificationInvalidNonce)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidPublicKey represents a InvalidPublicKey error raised by the Verification contract.
type VerificationInvalidPublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPublicKey()
func VerificationInvalidPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa2d0fee854985cb3d3a394146b5df1c41d5b4d6fcede7fc57a3c05a49b2f3e73")
}

// UnpackInvalidPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPublicKey()
func (verification *Verification) UnpackInvalidPublicKeyError(raw []byte) (*VerificationInvalidPublicKey, error) {
	out := new(VerificationInvalidPublicKey)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidRequestBody represents a InvalidRequestBody error raised by the Verification contract.
type VerificationInvalidRequestBody struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRequestBody()
func VerificationInvalidRequestBodyErrorID() common.Hash {
	return common.HexToHash("0xfe524bcc2b70aad90ebe31fd595732afd3bed1387a6e1141569409a40554173c")
}

// UnpackInvalidRequestBodyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRequestBody()
func (verification *Verification) UnpackInvalidRequestBodyError(raw []byte) (*VerificationInvalidRequestBody, error) {
	out := new(VerificationInvalidRequestBody)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidRequestBody", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidResponseData represents a InvalidResponseData error raised by the Verification contract.
type VerificationInvalidResponseData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseData()
func VerificationInvalidResponseDataErrorID() common.Hash {
	return common.HexToHash("0x48566617128355a033880805d6462da9bf9c6fe4c61b922997451a71262990af")
}

// UnpackInvalidResponseDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseData()
func (verification *Verification) UnpackInvalidResponseDataError(raw []byte) (*VerificationInvalidResponseData, error) {
	out := new(VerificationInvalidResponseData)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidResponseData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidSigningAlgo represents a InvalidSigningAlgo error raised by the Verification contract.
type VerificationInvalidSigningAlgo struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningAlgo()
func VerificationInvalidSigningAlgoErrorID() common.Hash {
	return common.HexToHash("0x9b7198a697b61e546ef9ee593a6f77d2bf9f1df9dcbba655ce96d7e2c75346c7")
}

// UnpackInvalidSigningAlgoError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningAlgo()
func (verification *Verification) UnpackInvalidSigningAlgoError(raw []byte) (*VerificationInvalidSigningAlgo, error) {
	out := new(VerificationInvalidSigningAlgo)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidSigningAlgo", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidSigningPolicy represents a InvalidSigningPolicy error raised by the Verification contract.
type VerificationInvalidSigningPolicy struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningPolicy()
func VerificationInvalidSigningPolicyErrorID() common.Hash {
	return common.HexToHash("0x3d7cc620812b32543e78b18734710415246f44f2e8d99c9198a7cb7573499dda")
}

// UnpackInvalidSigningPolicyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningPolicy()
func (verification *Verification) UnpackInvalidSigningPolicyError(raw []byte) (*VerificationInvalidSigningPolicy, error) {
	out := new(VerificationInvalidSigningPolicy)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidSigningPolicy", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidThreshold represents a InvalidThreshold error raised by the Verification contract.
type VerificationInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func VerificationInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (verification *Verification) UnpackInvalidThresholdError(raw []byte) (*VerificationInvalidThreshold, error) {
	out := new(VerificationInvalidThreshold)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationInvalidWalletStatus represents a InvalidWalletStatus error raised by the Verification contract.
type VerificationInvalidWalletStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidWalletStatus()
func VerificationInvalidWalletStatusErrorID() common.Hash {
	return common.HexToHash("0xd700120e4896ea7835a5e9e4feb690644ced5cad1c52209afdcbe06b96d91345")
}

// UnpackInvalidWalletStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidWalletStatus()
func (verification *Verification) UnpackInvalidWalletStatusError(raw []byte) (*VerificationInvalidWalletStatus, error) {
	out := new(VerificationInvalidWalletStatus)
	if err := verification.abi.UnpackIntoInterface(out, "InvalidWalletStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationKeyTypeNotSupported represents a KeyTypeNotSupported error raised by the Verification contract.
type VerificationKeyTypeNotSupported struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func VerificationKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x85004df1f47814c2de6fbf03d3f2b1ca8c67c182d0d54278766543b6be3bc603")
}

// UnpackKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func (verification *Verification) UnpackKeyTypeNotSupportedError(raw []byte) (*VerificationKeyTypeNotSupported, error) {
	out := new(VerificationKeyTypeNotSupported)
	if err := verification.abi.UnpackIntoInterface(out, "KeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationLengthsMismatch represents a LengthsMismatch error raised by the Verification contract.
type VerificationLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func VerificationLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (verification *Verification) UnpackLengthsMismatchError(raw []byte) (*VerificationLengthsMismatch, error) {
	out := new(VerificationLengthsMismatch)
	if err := verification.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationMessageEmpty represents a MessageEmpty error raised by the Verification contract.
type VerificationMessageEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MessageEmpty()
func VerificationMessageEmptyErrorID() common.Hash {
	return common.HexToHash("0x3b06a1beedd2da56654bd6f78f50f74e9bc7469bae35c52fd3680e4d95103c79")
}

// UnpackMessageEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MessageEmpty()
func (verification *Verification) UnpackMessageEmptyError(raw []byte) (*VerificationMessageEmpty, error) {
	out := new(VerificationMessageEmpty)
	if err := verification.abi.UnpackIntoInterface(out, "MessageEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationNoAddresses represents a NoAddresses error raised by the Verification contract.
type VerificationNoAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoAddresses()
func VerificationNoAddressesErrorID() common.Hash {
	return common.HexToHash("0xb2e79f05746fa6c9f3d61f9efe85959371cde5030a6c2dd70c92b63578c5cde7")
}

// UnpackNoAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoAddresses()
func (verification *Verification) UnpackNoAddressesError(raw []byte) (*VerificationNoAddresses, error) {
	out := new(VerificationNoAddresses)
	if err := verification.abi.UnpackIntoInterface(out, "NoAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationNoTeeMachinesSpecified represents a NoTeeMachinesSpecified error raised by the Verification contract.
type VerificationNoTeeMachinesSpecified struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoTeeMachinesSpecified()
func VerificationNoTeeMachinesSpecifiedErrorID() common.Hash {
	return common.HexToHash("0xe10387d0118f08d51d80fe03b0c14ed458c46aac2ae9964e0b628978e6938db8")
}

// UnpackNoTeeMachinesSpecifiedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoTeeMachinesSpecified()
func (verification *Verification) UnpackNoTeeMachinesSpecifiedError(raw []byte) (*VerificationNoTeeMachinesSpecified, error) {
	out := new(VerificationNoTeeMachinesSpecified)
	if err := verification.abi.UnpackIntoInterface(out, "NoTeeMachinesSpecified", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationNotInitializing represents a NotInitializing error raised by the Verification contract.
type VerificationNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func VerificationNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (verification *Verification) UnpackNotInitializingError(raw []byte) (*VerificationNotInitializing, error) {
	out := new(VerificationNotInitializing)
	if err := verification.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationNotOwnerOrPauser represents a NotOwnerOrPauser error raised by the Verification contract.
type VerificationNotOwnerOrPauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrPauser(address caller)
func VerificationNotOwnerOrPauserErrorID() common.Hash {
	return common.HexToHash("0x46c0d9af6b895093e022ad2984c54e60192e6e26c5063daf232e825c6e07ff9a")
}

// UnpackNotOwnerOrPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrPauser(address caller)
func (verification *Verification) UnpackNotOwnerOrPauserError(raw []byte) (*VerificationNotOwnerOrPauser, error) {
	out := new(VerificationNotOwnerOrPauser)
	if err := verification.abi.UnpackIntoInterface(out, "NotOwnerOrPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationNotOwnerOrUnpauser represents a NotOwnerOrUnpauser error raised by the Verification contract.
type VerificationNotOwnerOrUnpauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func VerificationNotOwnerOrUnpauserErrorID() common.Hash {
	return common.HexToHash("0x81e700e8f0bd5ba4a6d6a2ce9a1ad0ac9bb1b1b1e62d921a55cbc4b53f1da649")
}

// UnpackNotOwnerOrUnpauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func (verification *Verification) UnpackNotOwnerOrUnpauserError(raw []byte) (*VerificationNotOwnerOrUnpauser, error) {
	out := new(VerificationNotOwnerOrUnpauser)
	if err := verification.abi.UnpackIntoInterface(out, "NotOwnerOrUnpauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationOnlyExtensionOwner represents a OnlyExtensionOwner error raised by the Verification contract.
type VerificationOnlyExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwner()
func VerificationOnlyExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xcd49fd1de5a297cd1ce055d5abffcf2cb25ad9492bd39a703a238af57b2a9488")
}

// UnpackOnlyExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwner()
func (verification *Verification) UnpackOnlyExtensionOwnerError(raw []byte) (*VerificationOnlyExtensionOwner, error) {
	out := new(VerificationOnlyExtensionOwner)
	if err := verification.abi.UnpackIntoInterface(out, "OnlyExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the Verification contract.
type VerificationOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func VerificationOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (verification *Verification) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*VerificationOnlyExtensionOwnerOrOperator, error) {
	out := new(VerificationOnlyExtensionOwnerOrOperator)
	if err := verification.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationOnlyGovernance represents a OnlyGovernance error raised by the Verification contract.
type VerificationOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func VerificationOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (verification *Verification) UnpackOnlyGovernanceError(raw []byte) (*VerificationOnlyGovernance, error) {
	out := new(VerificationOnlyGovernance)
	if err := verification.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationOnlyOwner represents a OnlyOwner error raised by the Verification contract.
type VerificationOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func VerificationOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (verification *Verification) UnpackOnlyOwnerError(raw []byte) (*VerificationOnlyOwner, error) {
	out := new(VerificationOnlyOwner)
	if err := verification.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationOnlyOwnerOrBackupManager represents a OnlyOwnerOrBackupManager error raised by the Verification contract.
type VerificationOnlyOwnerOrBackupManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrBackupManager()
func VerificationOnlyOwnerOrBackupManagerErrorID() common.Hash {
	return common.HexToHash("0x249a6cf6e22092f9c80798e2a0c583c0f36aea0779ef38a5b5cdb6a112a0e6df")
}

// UnpackOnlyOwnerOrBackupManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrBackupManager()
func (verification *Verification) UnpackOnlyOwnerOrBackupManagerError(raw []byte) (*VerificationOnlyOwnerOrBackupManager, error) {
	out := new(VerificationOnlyOwnerOrBackupManager)
	if err := verification.abi.UnpackIntoInterface(out, "OnlyOwnerOrBackupManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the Verification contract.
type VerificationOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func VerificationOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (verification *Verification) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*VerificationOnlyProductionOrPausedStatus, error) {
	out := new(VerificationOnlyProductionOrPausedStatus)
	if err := verification.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationOnlyProposedOwner represents a OnlyProposedOwner error raised by the Verification contract.
type VerificationOnlyProposedOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProposedOwner()
func VerificationOnlyProposedOwnerErrorID() common.Hash {
	return common.HexToHash("0x49cb15061ee558e74f66f937f740fd2a3791a89d2d139862222e5c97cabc02cd")
}

// UnpackOnlyProposedOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProposedOwner()
func (verification *Verification) UnpackOnlyProposedOwnerError(raw []byte) (*VerificationOnlyProposedOwner, error) {
	out := new(VerificationOnlyProposedOwner)
	if err := verification.abi.UnpackIntoInterface(out, "OnlyProposedOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationOperationCommandEmpty represents a OperationCommandEmpty error raised by the Verification contract.
type VerificationOperationCommandEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationCommandEmpty()
func VerificationOperationCommandEmptyErrorID() common.Hash {
	return common.HexToHash("0x038dc5a8067401ab291233391b3d11ff3ce1b21a8aebcf4297e6a9f759f65846")
}

// UnpackOperationCommandEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationCommandEmpty()
func (verification *Verification) UnpackOperationCommandEmptyError(raw []byte) (*VerificationOperationCommandEmpty, error) {
	out := new(VerificationOperationCommandEmpty)
	if err := verification.abi.UnpackIntoInterface(out, "OperationCommandEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationOperationTypeEmpty represents a OperationTypeEmpty error raised by the Verification contract.
type VerificationOperationTypeEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationTypeEmpty()
func VerificationOperationTypeEmptyErrorID() common.Hash {
	return common.HexToHash("0x06e3960d8a83ea8972d826357113f286ce60947bac56680f119fca98a2d6125b")
}

// UnpackOperationTypeEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationTypeEmpty()
func (verification *Verification) UnpackOperationTypeEmptyError(raw []byte) (*VerificationOperationTypeEmpty, error) {
	out := new(VerificationOperationTypeEmpty)
	if err := verification.abi.UnpackIntoInterface(out, "OperationTypeEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationOwnerNotAllowed represents a OwnerNotAllowed error raised by the Verification contract.
type VerificationOwnerNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerNotAllowed()
func VerificationOwnerNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x3d5c9996545c6bf4287c26c1c2d8e6b4c39b92f19ba706c3a756720bf9c4fd24")
}

// UnpackOwnerNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerNotAllowed()
func (verification *Verification) UnpackOwnerNotAllowedError(raw []byte) (*VerificationOwnerNotAllowed, error) {
	out := new(VerificationOwnerNotAllowed)
	if err := verification.abi.UnpackIntoInterface(out, "OwnerNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the Verification contract.
type VerificationTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func VerificationTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (verification *Verification) UnpackTeeMachineNotAvailableError(raw []byte) (*VerificationTeeMachineNotAvailable, error) {
	out := new(VerificationTeeMachineNotAvailable)
	if err := verification.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationTeeNotFound represents a TeeNotFound error raised by the Verification contract.
type VerificationTeeNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeNotFound()
func VerificationTeeNotFoundErrorID() common.Hash {
	return common.HexToHash("0xceb05b6852ed4c7dc3a718eb0d26555398a76877a44dd41dc92c5fe7e6ef8c2d")
}

// UnpackTeeNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeNotFound()
func (verification *Verification) UnpackTeeNotFoundError(raw []byte) (*VerificationTeeNotFound, error) {
	out := new(VerificationTeeNotFound)
	if err := verification.abi.UnpackIntoInterface(out, "TeeNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationVersionNotSupported represents a VersionNotSupported error raised by the Verification contract.
type VerificationVersionNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionNotSupported()
func VerificationVersionNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x6db0b17b06266b37e721024a8ba1221439799ebb779dfd397865e37250abc9a8")
}

// UnpackVersionNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionNotSupported()
func (verification *Verification) UnpackVersionNotSupportedError(raw []byte) (*VerificationVersionNotSupported, error) {
	out := new(VerificationVersionNotSupported)
	if err := verification.abi.UnpackIntoInterface(out, "VersionNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}
