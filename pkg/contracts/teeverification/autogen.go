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

// IFtdcHubFtdcResponseHeader is an auto generated low-level Go binding around an user-defined struct.
type IFtdcHubFtdcResponseHeader struct {
	AttestationType    [32]byte
	SourceId           [32]byte
	ThresholdBIPS      uint16
	Cosigners          []common.Address
	CosignersThreshold uint64
	Timestamp          uint64
}

// IFtdcVerificationFtdcSignatures is an auto generated low-level Go binding around an user-defined struct.
type IFtdcVerificationFtdcSignatures struct {
	SigningPolicySignatures []byte
	TeeSignatures           []Signature
	CosignerSignatures      []Signature
}

// IPMWMultisigAccountConfiguredProof is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredProof struct {
	Signatures   IFtdcVerificationFtdcSignatures
	Header       IFtdcHubFtdcResponseHeader
	RequestBody  IPMWMultisigAccountConfiguredRequestBody
	ResponseBody IPMWMultisigAccountConfiguredResponseBody
}

// IPMWMultisigAccountConfiguredRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredRequestBody struct {
	WalletAddress string
	PublicKeys    [][]byte
	Threshold     uint64
	OpType        [32]byte
}

// IPMWMultisigAccountConfiguredResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredResponseBody struct {
	Status   uint8
	Sequence uint64
}

// ITeeAvailabilityCheckProof is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckProof struct {
	Signatures   IFtdcVerificationFtdcSignatures
	Header       IFtdcHubFtdcResponseHeader
	RequestBody  ITeeAvailabilityCheckRequestBody
	ResponseBody ITeeAvailabilityCheckResponseBody
}

// ITeeAvailabilityCheckRequestBody is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckRequestBody struct {
	TeeId     common.Address
	Url       string
	Challenge [32]byte
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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengeTs\",\"type\":\"uint256\"}],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"}],\"name\":\"AvailabilityCheckValidityExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"challengeTs\",\"type\":\"uint256\"}],\"name\":\"ChallengeExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtension\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialSigningPolicy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestBody\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningPolicy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletAddressZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTs\",\"type\":\"uint256\"}],\"name\":\"AvailabilityCheckValidityExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"CosignersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"availabilityCheckValidityDurationSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"signingPolicyValidityDurationInRewardEpochs\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"challengeValidityDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"SettingsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"name\":\"TeeAttestationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REG_OP_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TEE_ATTESTATION\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TEE_SOURCE_ID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"confirmAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftdcHub\",\"outputs\":[{\"internalType\":\"contractIFtdcHub\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftdcVerification\",\"outputs\":[{\"internalType\":\"contractIFtdcVerification\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCosigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSettings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_availabilityCheckValidityDurationSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_challengeValidityDurationSeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_availabilityCheckValidityDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"_signingPolicyValidityDurationInRewardEpochs\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_challengeValidityDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relay\",\"outputs\":[{\"internalType\":\"contractIRelay\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_testOnTeeId\",\"type\":\"address\"}],\"name\":\"requestAvailabilityCheckAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_walletAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_testOnTeeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"requestPMWMultisigAccountConfiguredAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"requestTeeAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"setCosigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeExtensionRegistry\",\"outputs\":[{\"internalType\":\"contractITeeExtensionRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeMachineRegistry\",\"outputs\":[{\"internalType\":\"contractITeeMachineRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeReplication\",\"outputs\":[{\"internalType\":\"contractITeeReplication\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeSystemStateVerifier\",\"outputs\":[{\"internalType\":\"contractIITeeSystemStateVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletKeyManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletKeyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletProjectManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletProjectManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_availabilityCheckValidityDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint24\",\"name\":\"_signingPolicyValidityDurationInRewardEpochs\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"_challengeValidityDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"updateSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAvailabilityCheckProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"walletAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigAccountConfigured.PMWMultisigAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyPMWMultisigAccountConfiguredProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// REGOPTYPE is a free data retrieval call binding the contract method 0xf468b2eb.
//
// Solidity: function REG_OP_TYPE() view returns(bytes32)
func (_TeeVerification *TeeVerificationCaller) REGOPTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "REG_OP_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGOPTYPE is a free data retrieval call binding the contract method 0xf468b2eb.
//
// Solidity: function REG_OP_TYPE() view returns(bytes32)
func (_TeeVerification *TeeVerificationSession) REGOPTYPE() ([32]byte, error) {
	return _TeeVerification.Contract.REGOPTYPE(&_TeeVerification.CallOpts)
}

// REGOPTYPE is a free data retrieval call binding the contract method 0xf468b2eb.
//
// Solidity: function REG_OP_TYPE() view returns(bytes32)
func (_TeeVerification *TeeVerificationCallerSession) REGOPTYPE() ([32]byte, error) {
	return _TeeVerification.Contract.REGOPTYPE(&_TeeVerification.CallOpts)
}

// TEEATTESTATION is a free data retrieval call binding the contract method 0x9a72a205.
//
// Solidity: function TEE_ATTESTATION() view returns(bytes32)
func (_TeeVerification *TeeVerificationCaller) TEEATTESTATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "TEE_ATTESTATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TEEATTESTATION is a free data retrieval call binding the contract method 0x9a72a205.
//
// Solidity: function TEE_ATTESTATION() view returns(bytes32)
func (_TeeVerification *TeeVerificationSession) TEEATTESTATION() ([32]byte, error) {
	return _TeeVerification.Contract.TEEATTESTATION(&_TeeVerification.CallOpts)
}

// TEEATTESTATION is a free data retrieval call binding the contract method 0x9a72a205.
//
// Solidity: function TEE_ATTESTATION() view returns(bytes32)
func (_TeeVerification *TeeVerificationCallerSession) TEEATTESTATION() ([32]byte, error) {
	return _TeeVerification.Contract.TEEATTESTATION(&_TeeVerification.CallOpts)
}

// TEESOURCEID is a free data retrieval call binding the contract method 0x319e1139.
//
// Solidity: function TEE_SOURCE_ID() view returns(bytes32)
func (_TeeVerification *TeeVerificationCaller) TEESOURCEID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "TEE_SOURCE_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TEESOURCEID is a free data retrieval call binding the contract method 0x319e1139.
//
// Solidity: function TEE_SOURCE_ID() view returns(bytes32)
func (_TeeVerification *TeeVerificationSession) TEESOURCEID() ([32]byte, error) {
	return _TeeVerification.Contract.TEESOURCEID(&_TeeVerification.CallOpts)
}

// TEESOURCEID is a free data retrieval call binding the contract method 0x319e1139.
//
// Solidity: function TEE_SOURCE_ID() view returns(bytes32)
func (_TeeVerification *TeeVerificationCallerSession) TEESOURCEID() ([32]byte, error) {
	return _TeeVerification.Contract.TEESOURCEID(&_TeeVerification.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeVerification *TeeVerificationCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeVerification *TeeVerificationSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeVerification.Contract.UPGRADEINTERFACEVERSION(&_TeeVerification.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeVerification *TeeVerificationCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeVerification.Contract.UPGRADEINTERFACEVERSION(&_TeeVerification.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeVerification *TeeVerificationCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeVerification *TeeVerificationSession) FlareSystemsManager() (common.Address, error) {
	return _TeeVerification.Contract.FlareSystemsManager(&_TeeVerification.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) FlareSystemsManager() (common.Address, error) {
	return _TeeVerification.Contract.FlareSystemsManager(&_TeeVerification.CallOpts)
}

// FtdcHub is a free data retrieval call binding the contract method 0xfdf8d81f.
//
// Solidity: function ftdcHub() view returns(address)
func (_TeeVerification *TeeVerificationCaller) FtdcHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "ftdcHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FtdcHub is a free data retrieval call binding the contract method 0xfdf8d81f.
//
// Solidity: function ftdcHub() view returns(address)
func (_TeeVerification *TeeVerificationSession) FtdcHub() (common.Address, error) {
	return _TeeVerification.Contract.FtdcHub(&_TeeVerification.CallOpts)
}

// FtdcHub is a free data retrieval call binding the contract method 0xfdf8d81f.
//
// Solidity: function ftdcHub() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) FtdcHub() (common.Address, error) {
	return _TeeVerification.Contract.FtdcHub(&_TeeVerification.CallOpts)
}

// FtdcVerification is a free data retrieval call binding the contract method 0x2d5a3dcf.
//
// Solidity: function ftdcVerification() view returns(address)
func (_TeeVerification *TeeVerificationCaller) FtdcVerification(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "ftdcVerification")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FtdcVerification is a free data retrieval call binding the contract method 0x2d5a3dcf.
//
// Solidity: function ftdcVerification() view returns(address)
func (_TeeVerification *TeeVerificationSession) FtdcVerification() (common.Address, error) {
	return _TeeVerification.Contract.FtdcVerification(&_TeeVerification.CallOpts)
}

// FtdcVerification is a free data retrieval call binding the contract method 0x2d5a3dcf.
//
// Solidity: function ftdcVerification() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) FtdcVerification() (common.Address, error) {
	return _TeeVerification.Contract.FtdcVerification(&_TeeVerification.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeVerification *TeeVerificationCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeVerification *TeeVerificationSession) GetAddressUpdater() (common.Address, error) {
	return _TeeVerification.Contract.GetAddressUpdater(&_TeeVerification.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeVerification *TeeVerificationCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeVerification.Contract.GetAddressUpdater(&_TeeVerification.CallOpts)
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

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeVerification *TeeVerificationCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeVerification *TeeVerificationSession) Governance() (common.Address, error) {
	return _TeeVerification.Contract.Governance(&_TeeVerification.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) Governance() (common.Address, error) {
	return _TeeVerification.Contract.Governance(&_TeeVerification.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeVerification *TeeVerificationCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeVerification *TeeVerificationSession) GovernanceSettings() (common.Address, error) {
	return _TeeVerification.Contract.GovernanceSettings(&_TeeVerification.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeeVerification.Contract.GovernanceSettings(&_TeeVerification.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeVerification *TeeVerificationCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeVerification *TeeVerificationSession) Implementation() (common.Address, error) {
	return _TeeVerification.Contract.Implementation(&_TeeVerification.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) Implementation() (common.Address, error) {
	return _TeeVerification.Contract.Implementation(&_TeeVerification.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeVerification *TeeVerificationCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeVerification *TeeVerificationSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeVerification.Contract.IsExecutor(&_TeeVerification.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeVerification *TeeVerificationCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeVerification.Contract.IsExecutor(&_TeeVerification.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeVerification *TeeVerificationCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeVerification *TeeVerificationSession) ProductionMode() (bool, error) {
	return _TeeVerification.Contract.ProductionMode(&_TeeVerification.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeVerification *TeeVerificationCallerSession) ProductionMode() (bool, error) {
	return _TeeVerification.Contract.ProductionMode(&_TeeVerification.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeVerification *TeeVerificationCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeVerification *TeeVerificationSession) ProxiableUUID() ([32]byte, error) {
	return _TeeVerification.Contract.ProxiableUUID(&_TeeVerification.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeVerification *TeeVerificationCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeVerification.Contract.ProxiableUUID(&_TeeVerification.CallOpts)
}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_TeeVerification *TeeVerificationCaller) Relay(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "relay")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_TeeVerification *TeeVerificationSession) Relay() (common.Address, error) {
	return _TeeVerification.Contract.Relay(&_TeeVerification.CallOpts)
}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) Relay() (common.Address, error) {
	return _TeeVerification.Contract.Relay(&_TeeVerification.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeVerification *TeeVerificationCaller) TeeExtensionRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "teeExtensionRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeVerification *TeeVerificationSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeVerification.Contract.TeeExtensionRegistry(&_TeeVerification.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeVerification.Contract.TeeExtensionRegistry(&_TeeVerification.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeVerification *TeeVerificationCaller) TeeMachineRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "teeMachineRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeVerification *TeeVerificationSession) TeeMachineRegistry() (common.Address, error) {
	return _TeeVerification.Contract.TeeMachineRegistry(&_TeeVerification.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) TeeMachineRegistry() (common.Address, error) {
	return _TeeVerification.Contract.TeeMachineRegistry(&_TeeVerification.CallOpts)
}

// TeeReplication is a free data retrieval call binding the contract method 0x0f10ad79.
//
// Solidity: function teeReplication() view returns(address)
func (_TeeVerification *TeeVerificationCaller) TeeReplication(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "teeReplication")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeReplication is a free data retrieval call binding the contract method 0x0f10ad79.
//
// Solidity: function teeReplication() view returns(address)
func (_TeeVerification *TeeVerificationSession) TeeReplication() (common.Address, error) {
	return _TeeVerification.Contract.TeeReplication(&_TeeVerification.CallOpts)
}

// TeeReplication is a free data retrieval call binding the contract method 0x0f10ad79.
//
// Solidity: function teeReplication() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) TeeReplication() (common.Address, error) {
	return _TeeVerification.Contract.TeeReplication(&_TeeVerification.CallOpts)
}

// TeeSystemStateVerifier is a free data retrieval call binding the contract method 0xb3e8e44d.
//
// Solidity: function teeSystemStateVerifier() view returns(address)
func (_TeeVerification *TeeVerificationCaller) TeeSystemStateVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "teeSystemStateVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeSystemStateVerifier is a free data retrieval call binding the contract method 0xb3e8e44d.
//
// Solidity: function teeSystemStateVerifier() view returns(address)
func (_TeeVerification *TeeVerificationSession) TeeSystemStateVerifier() (common.Address, error) {
	return _TeeVerification.Contract.TeeSystemStateVerifier(&_TeeVerification.CallOpts)
}

// TeeSystemStateVerifier is a free data retrieval call binding the contract method 0xb3e8e44d.
//
// Solidity: function teeSystemStateVerifier() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) TeeSystemStateVerifier() (common.Address, error) {
	return _TeeVerification.Contract.TeeSystemStateVerifier(&_TeeVerification.CallOpts)
}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeVerification *TeeVerificationCaller) TeeWalletKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "teeWalletKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeVerification *TeeVerificationSession) TeeWalletKeyManager() (common.Address, error) {
	return _TeeVerification.Contract.TeeWalletKeyManager(&_TeeVerification.CallOpts)
}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) TeeWalletKeyManager() (common.Address, error) {
	return _TeeVerification.Contract.TeeWalletKeyManager(&_TeeVerification.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeVerification *TeeVerificationCaller) TeeWalletManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "teeWalletManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeVerification *TeeVerificationSession) TeeWalletManager() (common.Address, error) {
	return _TeeVerification.Contract.TeeWalletManager(&_TeeVerification.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) TeeWalletManager() (common.Address, error) {
	return _TeeVerification.Contract.TeeWalletManager(&_TeeVerification.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeVerification *TeeVerificationCaller) TeeWalletProjectManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "teeWalletProjectManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeVerification *TeeVerificationSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeeVerification.Contract.TeeWalletProjectManager(&_TeeVerification.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeVerification *TeeVerificationCallerSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeeVerification.Contract.TeeWalletProjectManager(&_TeeVerification.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeVerification *TeeVerificationCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeeVerification.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_TeeVerification *TeeVerificationSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeVerification.Contract.TimelockedCalls(&_TeeVerification.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeVerification *TeeVerificationCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeVerification.Contract.TimelockedCalls(&_TeeVerification.CallOpts, selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeVerification *TeeVerificationTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeVerification *TeeVerificationSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeVerification.Contract.CancelGovernanceCall(&_TeeVerification.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeVerification *TeeVerificationTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeVerification.Contract.CancelGovernanceCall(&_TeeVerification.TransactOpts, _selector)
}

// ConfirmAvailability is a paid mutator transaction binding the contract method 0x21dd39a1.
//
// Solidity: function confirmAvailability(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeVerification *TeeVerificationTransactor) ConfirmAvailability(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "confirmAvailability", _proof)
}

// ConfirmAvailability is a paid mutator transaction binding the contract method 0x21dd39a1.
//
// Solidity: function confirmAvailability(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeVerification *TeeVerificationSession) ConfirmAvailability(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.ConfirmAvailability(&_TeeVerification.TransactOpts, _proof)
}

// ConfirmAvailability is a paid mutator transaction binding the contract method 0x21dd39a1.
//
// Solidity: function confirmAvailability(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeVerification *TeeVerificationTransactorSession) ConfirmAvailability(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.ConfirmAvailability(&_TeeVerification.TransactOpts, _proof)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeVerification *TeeVerificationTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeVerification *TeeVerificationSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeVerification.Contract.ExecuteGovernanceCall(&_TeeVerification.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeVerification *TeeVerificationTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeVerification.Contract.ExecuteGovernanceCall(&_TeeVerification.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeVerification *TeeVerificationTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeVerification *TeeVerificationSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.Initialise(&_TeeVerification.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeVerification *TeeVerificationTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.Initialise(&_TeeVerification.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc7e1903d.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _availabilityCheckValidityDurationSeconds, uint24 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds) returns()
func (_TeeVerification *TeeVerificationTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _availabilityCheckValidityDurationSeconds uint64, _signingPolicyValidityDurationInRewardEpochs *big.Int, _challengeValidityDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater, _availabilityCheckValidityDurationSeconds, _signingPolicyValidityDurationInRewardEpochs, _challengeValidityDurationSeconds)
}

// Initialize is a paid mutator transaction binding the contract method 0xc7e1903d.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _availabilityCheckValidityDurationSeconds, uint24 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds) returns()
func (_TeeVerification *TeeVerificationSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _availabilityCheckValidityDurationSeconds uint64, _signingPolicyValidityDurationInRewardEpochs *big.Int, _challengeValidityDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeVerification.Contract.Initialize(&_TeeVerification.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _availabilityCheckValidityDurationSeconds, _signingPolicyValidityDurationInRewardEpochs, _challengeValidityDurationSeconds)
}

// Initialize is a paid mutator transaction binding the contract method 0xc7e1903d.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _availabilityCheckValidityDurationSeconds, uint24 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds) returns()
func (_TeeVerification *TeeVerificationTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _availabilityCheckValidityDurationSeconds uint64, _signingPolicyValidityDurationInRewardEpochs *big.Int, _challengeValidityDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeVerification.Contract.Initialize(&_TeeVerification.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater, _availabilityCheckValidityDurationSeconds, _signingPolicyValidityDurationInRewardEpochs, _challengeValidityDurationSeconds)
}

// RequestAvailabilityCheckAttestation is a paid mutator transaction binding the contract method 0x0c336e35.
//
// Solidity: function requestAvailabilityCheckAttestation(address _teeId, address _testOnTeeId) payable returns()
func (_TeeVerification *TeeVerificationTransactor) RequestAvailabilityCheckAttestation(opts *bind.TransactOpts, _teeId common.Address, _testOnTeeId common.Address) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "requestAvailabilityCheckAttestation", _teeId, _testOnTeeId)
}

// RequestAvailabilityCheckAttestation is a paid mutator transaction binding the contract method 0x0c336e35.
//
// Solidity: function requestAvailabilityCheckAttestation(address _teeId, address _testOnTeeId) payable returns()
func (_TeeVerification *TeeVerificationSession) RequestAvailabilityCheckAttestation(_teeId common.Address, _testOnTeeId common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestAvailabilityCheckAttestation(&_TeeVerification.TransactOpts, _teeId, _testOnTeeId)
}

// RequestAvailabilityCheckAttestation is a paid mutator transaction binding the contract method 0x0c336e35.
//
// Solidity: function requestAvailabilityCheckAttestation(address _teeId, address _testOnTeeId) payable returns()
func (_TeeVerification *TeeVerificationTransactorSession) RequestAvailabilityCheckAttestation(_teeId common.Address, _testOnTeeId common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestAvailabilityCheckAttestation(&_TeeVerification.TransactOpts, _teeId, _testOnTeeId)
}

// RequestPMWMultisigAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0x17417614.
//
// Solidity: function requestPMWMultisigAccountConfiguredAttestation(bytes32 _walletId, string _walletAddress, address _testOnTeeId, bytes32 _sourceId) payable returns()
func (_TeeVerification *TeeVerificationTransactor) RequestPMWMultisigAccountConfiguredAttestation(opts *bind.TransactOpts, _walletId [32]byte, _walletAddress string, _testOnTeeId common.Address, _sourceId [32]byte) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "requestPMWMultisigAccountConfiguredAttestation", _walletId, _walletAddress, _testOnTeeId, _sourceId)
}

// RequestPMWMultisigAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0x17417614.
//
// Solidity: function requestPMWMultisigAccountConfiguredAttestation(bytes32 _walletId, string _walletAddress, address _testOnTeeId, bytes32 _sourceId) payable returns()
func (_TeeVerification *TeeVerificationSession) RequestPMWMultisigAccountConfiguredAttestation(_walletId [32]byte, _walletAddress string, _testOnTeeId common.Address, _sourceId [32]byte) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestPMWMultisigAccountConfiguredAttestation(&_TeeVerification.TransactOpts, _walletId, _walletAddress, _testOnTeeId, _sourceId)
}

// RequestPMWMultisigAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0x17417614.
//
// Solidity: function requestPMWMultisigAccountConfiguredAttestation(bytes32 _walletId, string _walletAddress, address _testOnTeeId, bytes32 _sourceId) payable returns()
func (_TeeVerification *TeeVerificationTransactorSession) RequestPMWMultisigAccountConfiguredAttestation(_walletId [32]byte, _walletAddress string, _testOnTeeId common.Address, _sourceId [32]byte) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestPMWMultisigAccountConfiguredAttestation(&_TeeVerification.TransactOpts, _walletId, _walletAddress, _testOnTeeId, _sourceId)
}

// RequestTeeAttestation is a paid mutator transaction binding the contract method 0x84d6fe27.
//
// Solidity: function requestTeeAttestation(address _teeId) payable returns()
func (_TeeVerification *TeeVerificationTransactor) RequestTeeAttestation(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "requestTeeAttestation", _teeId)
}

// RequestTeeAttestation is a paid mutator transaction binding the contract method 0x84d6fe27.
//
// Solidity: function requestTeeAttestation(address _teeId) payable returns()
func (_TeeVerification *TeeVerificationSession) RequestTeeAttestation(_teeId common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestTeeAttestation(&_TeeVerification.TransactOpts, _teeId)
}

// RequestTeeAttestation is a paid mutator transaction binding the contract method 0x84d6fe27.
//
// Solidity: function requestTeeAttestation(address _teeId) payable returns()
func (_TeeVerification *TeeVerificationTransactorSession) RequestTeeAttestation(_teeId common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.RequestTeeAttestation(&_TeeVerification.TransactOpts, _teeId)
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

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeVerification *TeeVerificationTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeVerification *TeeVerificationSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeVerification.Contract.SwitchToProductionMode(&_TeeVerification.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeVerification *TeeVerificationTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeVerification.Contract.SwitchToProductionMode(&_TeeVerification.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeVerification *TeeVerificationTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeVerification *TeeVerificationSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.UpdateContractAddresses(&_TeeVerification.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeVerification *TeeVerificationTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeVerification.Contract.UpdateContractAddresses(&_TeeVerification.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0xc9f9d906.
//
// Solidity: function updateSettings(uint64 _availabilityCheckValidityDurationSeconds, uint24 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds) returns()
func (_TeeVerification *TeeVerificationTransactor) UpdateSettings(opts *bind.TransactOpts, _availabilityCheckValidityDurationSeconds uint64, _signingPolicyValidityDurationInRewardEpochs *big.Int, _challengeValidityDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "updateSettings", _availabilityCheckValidityDurationSeconds, _signingPolicyValidityDurationInRewardEpochs, _challengeValidityDurationSeconds)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0xc9f9d906.
//
// Solidity: function updateSettings(uint64 _availabilityCheckValidityDurationSeconds, uint24 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds) returns()
func (_TeeVerification *TeeVerificationSession) UpdateSettings(_availabilityCheckValidityDurationSeconds uint64, _signingPolicyValidityDurationInRewardEpochs *big.Int, _challengeValidityDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeVerification.Contract.UpdateSettings(&_TeeVerification.TransactOpts, _availabilityCheckValidityDurationSeconds, _signingPolicyValidityDurationInRewardEpochs, _challengeValidityDurationSeconds)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0xc9f9d906.
//
// Solidity: function updateSettings(uint64 _availabilityCheckValidityDurationSeconds, uint24 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds) returns()
func (_TeeVerification *TeeVerificationTransactorSession) UpdateSettings(_availabilityCheckValidityDurationSeconds uint64, _signingPolicyValidityDurationInRewardEpochs *big.Int, _challengeValidityDurationSeconds uint64) (*types.Transaction, error) {
	return _TeeVerification.Contract.UpdateSettings(&_TeeVerification.TransactOpts, _availabilityCheckValidityDurationSeconds, _signingPolicyValidityDurationInRewardEpochs, _challengeValidityDurationSeconds)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeVerification *TeeVerificationTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeVerification *TeeVerificationSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeVerification.Contract.UpgradeToAndCall(&_TeeVerification.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeVerification *TeeVerificationTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeVerification.Contract.UpgradeToAndCall(&_TeeVerification.TransactOpts, _newImplementation, _data)
}

// VerifyAvailabilityCheckProof is a paid mutator transaction binding the contract method 0xa8b5b815.
//
// Solidity: function verifyAvailabilityCheckProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns(bool)
func (_TeeVerification *TeeVerificationTransactor) VerifyAvailabilityCheckProof(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "verifyAvailabilityCheckProof", _proof)
}

// VerifyAvailabilityCheckProof is a paid mutator transaction binding the contract method 0xa8b5b815.
//
// Solidity: function verifyAvailabilityCheckProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns(bool)
func (_TeeVerification *TeeVerificationSession) VerifyAvailabilityCheckProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.VerifyAvailabilityCheckProof(&_TeeVerification.TransactOpts, _proof)
}

// VerifyAvailabilityCheckProof is a paid mutator transaction binding the contract method 0xa8b5b815.
//
// Solidity: function verifyAvailabilityCheckProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns(bool)
func (_TeeVerification *TeeVerificationTransactorSession) VerifyAvailabilityCheckProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.VerifyAvailabilityCheckProof(&_TeeVerification.TransactOpts, _proof)
}

// VerifyPMWMultisigAccountConfiguredProof is a paid mutator transaction binding the contract method 0x64b6417c.
//
// Solidity: function verifyPMWMultisigAccountConfiguredProof(bytes32 _walletId, bytes32 _sourceId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(string,bytes[],uint64,bytes32),(uint8,uint64)) _proof) returns(bool)
func (_TeeVerification *TeeVerificationTransactor) VerifyPMWMultisigAccountConfiguredProof(opts *bind.TransactOpts, _walletId [32]byte, _sourceId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeeVerification.contract.Transact(opts, "verifyPMWMultisigAccountConfiguredProof", _walletId, _sourceId, _proof)
}

// VerifyPMWMultisigAccountConfiguredProof is a paid mutator transaction binding the contract method 0x64b6417c.
//
// Solidity: function verifyPMWMultisigAccountConfiguredProof(bytes32 _walletId, bytes32 _sourceId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(string,bytes[],uint64,bytes32),(uint8,uint64)) _proof) returns(bool)
func (_TeeVerification *TeeVerificationSession) VerifyPMWMultisigAccountConfiguredProof(_walletId [32]byte, _sourceId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.VerifyPMWMultisigAccountConfiguredProof(&_TeeVerification.TransactOpts, _walletId, _sourceId, _proof)
}

// VerifyPMWMultisigAccountConfiguredProof is a paid mutator transaction binding the contract method 0x64b6417c.
//
// Solidity: function verifyPMWMultisigAccountConfiguredProof(bytes32 _walletId, bytes32 _sourceId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(string,bytes[],uint64,bytes32),(uint8,uint64)) _proof) returns(bool)
func (_TeeVerification *TeeVerificationTransactorSession) VerifyPMWMultisigAccountConfiguredProof(_walletId [32]byte, _sourceId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeeVerification.Contract.VerifyPMWMultisigAccountConfiguredProof(&_TeeVerification.TransactOpts, _walletId, _sourceId, _proof)
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
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeVerification *TeeVerificationFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeVerificationGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeVerificationGovernanceCallTimelockedIterator{contract: _TeeVerification.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
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

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
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

// TeeVerificationGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeeVerification contract.
type TeeVerificationGovernedProductionModeEnteredIterator struct {
	Event *TeeVerificationGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeeVerificationGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationGovernedProductionModeEntered)
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
		it.Event = new(TeeVerificationGovernedProductionModeEntered)
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
func (it *TeeVerificationGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeeVerification contract.
type TeeVerificationGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeVerification *TeeVerificationFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeeVerificationGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeeVerificationGovernedProductionModeEnteredIterator{contract: _TeeVerification.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeVerification *TeeVerificationFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeeVerificationGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationGovernedProductionModeEntered)
				if err := _TeeVerification.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeeVerification *TeeVerificationFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeeVerificationGovernedProductionModeEntered, error) {
	event := new(TeeVerificationGovernedProductionModeEntered)
	if err := _TeeVerification.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
	SigningPolicyValidityDurationInRewardEpochs *big.Int
	ChallengeValidityDurationSeconds            uint64
	Raw                                         types.Log // Blockchain specific contextual infos
}

// FilterSettingsUpdated is a free log retrieval operation binding the contract event 0x27f607220ff59c6102bb2da2cfa5debe408ec0a6046f7039b77defb3f137a346.
//
// Solidity: event SettingsUpdated(uint64 availabilityCheckValidityDurationSeconds, uint24 signingPolicyValidityDurationInRewardEpochs, uint64 challengeValidityDurationSeconds)
func (_TeeVerification *TeeVerificationFilterer) FilterSettingsUpdated(opts *bind.FilterOpts) (*TeeVerificationSettingsUpdatedIterator, error) {

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "SettingsUpdated")
	if err != nil {
		return nil, err
	}
	return &TeeVerificationSettingsUpdatedIterator{contract: _TeeVerification.contract, event: "SettingsUpdated", logs: logs, sub: sub}, nil
}

// WatchSettingsUpdated is a free log subscription operation binding the contract event 0x27f607220ff59c6102bb2da2cfa5debe408ec0a6046f7039b77defb3f137a346.
//
// Solidity: event SettingsUpdated(uint64 availabilityCheckValidityDurationSeconds, uint24 signingPolicyValidityDurationInRewardEpochs, uint64 challengeValidityDurationSeconds)
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

// ParseSettingsUpdated is a log parse operation binding the contract event 0x27f607220ff59c6102bb2da2cfa5debe408ec0a6046f7039b77defb3f137a346.
//
// Solidity: event SettingsUpdated(uint64 availabilityCheckValidityDurationSeconds, uint24 signingPolicyValidityDurationInRewardEpochs, uint64 challengeValidityDurationSeconds)
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

// TeeVerificationTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeeVerification contract.
type TeeVerificationTimelockedGovernanceCallCanceledIterator struct {
	Event *TeeVerificationTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeeVerificationTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeeVerificationTimelockedGovernanceCallCanceled)
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
func (it *TeeVerificationTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeeVerification contract.
type TeeVerificationTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeVerification *TeeVerificationFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeeVerificationTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeeVerificationTimelockedGovernanceCallCanceledIterator{contract: _TeeVerification.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeVerification *TeeVerificationFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeeVerificationTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationTimelockedGovernanceCallCanceled)
				if err := _TeeVerification.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeeVerification *TeeVerificationFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeeVerificationTimelockedGovernanceCallCanceled, error) {
	event := new(TeeVerificationTimelockedGovernanceCallCanceled)
	if err := _TeeVerification.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVerificationTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeeVerification contract.
type TeeVerificationTimelockedGovernanceCallExecutedIterator struct {
	Event *TeeVerificationTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeeVerificationTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeeVerificationTimelockedGovernanceCallExecuted)
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
func (it *TeeVerificationTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeeVerification contract.
type TeeVerificationTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeVerification *TeeVerificationFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeeVerificationTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeeVerificationTimelockedGovernanceCallExecutedIterator{contract: _TeeVerification.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeVerification *TeeVerificationFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeeVerificationTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationTimelockedGovernanceCallExecuted)
				if err := _TeeVerification.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeeVerification *TeeVerificationFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeeVerificationTimelockedGovernanceCallExecuted, error) {
	event := new(TeeVerificationTimelockedGovernanceCallExecuted)
	if err := _TeeVerification.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeVerificationUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeeVerification contract.
type TeeVerificationUpgradedIterator struct {
	Event *TeeVerificationUpgraded // Event containing the contract specifics and raw log

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
func (it *TeeVerificationUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeVerificationUpgraded)
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
		it.Event = new(TeeVerificationUpgraded)
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
func (it *TeeVerificationUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeVerificationUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeVerificationUpgraded represents a Upgraded event raised by the TeeVerification contract.
type TeeVerificationUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeVerification *TeeVerificationFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeeVerificationUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeVerification.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeeVerificationUpgradedIterator{contract: _TeeVerification.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeVerification *TeeVerificationFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeeVerificationUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeVerification.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeVerificationUpgraded)
				if err := _TeeVerification.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeeVerification *TeeVerificationFilterer) ParseUpgraded(log types.Log) (*TeeVerificationUpgraded, error) {
	event := new(TeeVerificationUpgraded)
	if err := _TeeVerification.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
