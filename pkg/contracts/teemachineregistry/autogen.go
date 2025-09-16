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

// ITeeAvailabilityCheckProof is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckProof struct {
	Signatures   IFtdcVerificationFtdcSignatures
	Header       IFtdcHubFtdcResponseHeader
	RequestBody  ITeeAvailabilityCheckRequestBody
	ResponseBody ITeeAvailabilityCheckResponseBody
}

// ITeeAvailabilityCheckRequestBody is an auto generated low-level Go binding around an user-defined struct.
type ITeeAvailabilityCheckRequestBody struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
	Challenge  [32]byte
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

// ITeeMachineRegistryTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type ITeeMachineRegistryTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
}

// ITeeMachineRegistryTeeMachineWithAttestationData is an auto generated low-level Go binding around an user-defined struct.
type ITeeMachineRegistryTeeMachineWithAttestationData struct {
	TeeId        common.Address
	InitialTeeId common.Address
	Url          string
	CodeHash     [32]byte
	Platform     [32]byte
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// TeeMachineRegistryMetaData contains all meta data concerning the TeeMachineRegistry contract.
var TeeMachineRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AcTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNewStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseDataOrAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeProxyId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUrl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrDisabledVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTeeReplicationContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooMany\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"TeeMachineRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"TeeMachineSettingsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumITeeMachineRegistry.TeeStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"TeeMachineStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REG_OP_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"enumITeeMachineRegistry.TeeStatus\",\"name\":\"_newStatus\",\"type\":\"uint8\"}],\"name\":\"changeStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getActiveTeeMachines\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_urls\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"getAllActiveTeeMachines\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_urls\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"_totalLength\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getExtensionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getInitialSigningPolicyId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getLastStatusChangeTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getRandomTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachine\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structITeeMachineRegistry.TeeMachine\",\"name\":\"_teeMachine\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineStatus\",\"outputs\":[{\"internalType\":\"enumITeeMachineRegistry.TeeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineWithAttestationData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeMachineRegistry.TeeMachineWithAttestationData\",\"name\":\"_teeMachine\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"pauseWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"name\":\"proposedTeeOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relay\",\"outputs\":[{\"internalType\":\"contractIRelay\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTeeId\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"replicate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeExtensionRegistry\",\"outputs\":[{\"internalType\":\"contractITeeExtensionRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeOwnerAllowlist\",\"outputs\":[{\"internalType\":\"contractITeeOwnerAllowlist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeReplication\",\"outputs\":[{\"internalType\":\"contractITeeReplication\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeVerification\",\"outputs\":[{\"internalType\":\"contractITeeVerification\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"systemState\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"systemStateVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"stateVersion\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.TeeState\",\"name\":\"state\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"toProduction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"updateTeeMachineSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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

// REGOPTYPE is a free data retrieval call binding the contract method 0xf468b2eb.
//
// Solidity: function REG_OP_TYPE() view returns(bytes32)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) REGOPTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "REG_OP_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REGOPTYPE is a free data retrieval call binding the contract method 0xf468b2eb.
//
// Solidity: function REG_OP_TYPE() view returns(bytes32)
func (_TeeMachineRegistry *TeeMachineRegistrySession) REGOPTYPE() ([32]byte, error) {
	return _TeeMachineRegistry.Contract.REGOPTYPE(&_TeeMachineRegistry.CallOpts)
}

// REGOPTYPE is a free data retrieval call binding the contract method 0xf468b2eb.
//
// Solidity: function REG_OP_TYPE() view returns(bytes32)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) REGOPTYPE() ([32]byte, error) {
	return _TeeMachineRegistry.Contract.REGOPTYPE(&_TeeMachineRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeMachineRegistry *TeeMachineRegistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeMachineRegistry.Contract.UPGRADEINTERFACEVERSION(&_TeeMachineRegistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeMachineRegistry.Contract.UPGRADEINTERFACEVERSION(&_TeeMachineRegistry.CallOpts)
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

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetAddressUpdater() (common.Address, error) {
	return _TeeMachineRegistry.Contract.GetAddressUpdater(&_TeeMachineRegistry.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeMachineRegistry.Contract.GetAddressUpdater(&_TeeMachineRegistry.CallOpts)
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
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string) _teeMachine)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetTeeMachine(opts *bind.CallOpts, _teeId common.Address) (ITeeMachineRegistryTeeMachine, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getTeeMachine", _teeId)

	if err != nil {
		return *new(ITeeMachineRegistryTeeMachine), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeeMachineRegistryTeeMachine)).(*ITeeMachineRegistryTeeMachine)

	return out0, err

}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string) _teeMachine)
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetTeeMachine(_teeId common.Address) (ITeeMachineRegistryTeeMachine, error) {
	return _TeeMachineRegistry.Contract.GetTeeMachine(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string) _teeMachine)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetTeeMachine(_teeId common.Address) (ITeeMachineRegistryTeeMachine, error) {
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
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32) _teeMachine)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GetTeeMachineWithAttestationData(opts *bind.CallOpts, _teeId common.Address) (ITeeMachineRegistryTeeMachineWithAttestationData, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "getTeeMachineWithAttestationData", _teeId)

	if err != nil {
		return *new(ITeeMachineRegistryTeeMachineWithAttestationData), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeeMachineRegistryTeeMachineWithAttestationData)).(*ITeeMachineRegistryTeeMachineWithAttestationData)

	return out0, err

}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32) _teeMachine)
func (_TeeMachineRegistry *TeeMachineRegistrySession) GetTeeMachineWithAttestationData(_teeId common.Address) (ITeeMachineRegistryTeeMachineWithAttestationData, error) {
	return _TeeMachineRegistry.Contract.GetTeeMachineWithAttestationData(&_TeeMachineRegistry.CallOpts, _teeId)
}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32) _teeMachine)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GetTeeMachineWithAttestationData(_teeId common.Address) (ITeeMachineRegistryTeeMachineWithAttestationData, error) {
	return _TeeMachineRegistry.Contract.GetTeeMachineWithAttestationData(&_TeeMachineRegistry.CallOpts, _teeId)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistrySession) Governance() (common.Address, error) {
	return _TeeMachineRegistry.Contract.Governance(&_TeeMachineRegistry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) Governance() (common.Address, error) {
	return _TeeMachineRegistry.Contract.Governance(&_TeeMachineRegistry.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistrySession) GovernanceSettings() (common.Address, error) {
	return _TeeMachineRegistry.Contract.GovernanceSettings(&_TeeMachineRegistry.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeeMachineRegistry.Contract.GovernanceSettings(&_TeeMachineRegistry.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistrySession) Implementation() (common.Address, error) {
	return _TeeMachineRegistry.Contract.Implementation(&_TeeMachineRegistry.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) Implementation() (common.Address, error) {
	return _TeeMachineRegistry.Contract.Implementation(&_TeeMachineRegistry.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeMachineRegistry *TeeMachineRegistrySession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeMachineRegistry.Contract.IsExecutor(&_TeeMachineRegistry.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeMachineRegistry.Contract.IsExecutor(&_TeeMachineRegistry.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeMachineRegistry *TeeMachineRegistrySession) ProductionMode() (bool, error) {
	return _TeeMachineRegistry.Contract.ProductionMode(&_TeeMachineRegistry.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) ProductionMode() (bool, error) {
	return _TeeMachineRegistry.Contract.ProductionMode(&_TeeMachineRegistry.CallOpts)
}

// ProposedTeeOwner is a free data retrieval call binding the contract method 0x574d48e7.
//
// Solidity: function proposedTeeOwner(address teeId) view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) ProposedTeeOwner(opts *bind.CallOpts, teeId common.Address) (common.Address, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "proposedTeeOwner", teeId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposedTeeOwner is a free data retrieval call binding the contract method 0x574d48e7.
//
// Solidity: function proposedTeeOwner(address teeId) view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistrySession) ProposedTeeOwner(teeId common.Address) (common.Address, error) {
	return _TeeMachineRegistry.Contract.ProposedTeeOwner(&_TeeMachineRegistry.CallOpts, teeId)
}

// ProposedTeeOwner is a free data retrieval call binding the contract method 0x574d48e7.
//
// Solidity: function proposedTeeOwner(address teeId) view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) ProposedTeeOwner(teeId common.Address) (common.Address, error) {
	return _TeeMachineRegistry.Contract.ProposedTeeOwner(&_TeeMachineRegistry.CallOpts, teeId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeMachineRegistry *TeeMachineRegistrySession) ProxiableUUID() ([32]byte, error) {
	return _TeeMachineRegistry.Contract.ProxiableUUID(&_TeeMachineRegistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeMachineRegistry.Contract.ProxiableUUID(&_TeeMachineRegistry.CallOpts)
}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) Relay(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "relay")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistrySession) Relay() (common.Address, error) {
	return _TeeMachineRegistry.Contract.Relay(&_TeeMachineRegistry.CallOpts)
}

// Relay is a free data retrieval call binding the contract method 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) Relay() (common.Address, error) {
	return _TeeMachineRegistry.Contract.Relay(&_TeeMachineRegistry.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) TeeExtensionRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "teeExtensionRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistrySession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeMachineRegistry.Contract.TeeExtensionRegistry(&_TeeMachineRegistry.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeMachineRegistry.Contract.TeeExtensionRegistry(&_TeeMachineRegistry.CallOpts)
}

// TeeOwnerAllowlist is a free data retrieval call binding the contract method 0x327282d4.
//
// Solidity: function teeOwnerAllowlist() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) TeeOwnerAllowlist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "teeOwnerAllowlist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeOwnerAllowlist is a free data retrieval call binding the contract method 0x327282d4.
//
// Solidity: function teeOwnerAllowlist() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistrySession) TeeOwnerAllowlist() (common.Address, error) {
	return _TeeMachineRegistry.Contract.TeeOwnerAllowlist(&_TeeMachineRegistry.CallOpts)
}

// TeeOwnerAllowlist is a free data retrieval call binding the contract method 0x327282d4.
//
// Solidity: function teeOwnerAllowlist() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) TeeOwnerAllowlist() (common.Address, error) {
	return _TeeMachineRegistry.Contract.TeeOwnerAllowlist(&_TeeMachineRegistry.CallOpts)
}

// TeeReplication is a free data retrieval call binding the contract method 0x0f10ad79.
//
// Solidity: function teeReplication() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) TeeReplication(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "teeReplication")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeReplication is a free data retrieval call binding the contract method 0x0f10ad79.
//
// Solidity: function teeReplication() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistrySession) TeeReplication() (common.Address, error) {
	return _TeeMachineRegistry.Contract.TeeReplication(&_TeeMachineRegistry.CallOpts)
}

// TeeReplication is a free data retrieval call binding the contract method 0x0f10ad79.
//
// Solidity: function teeReplication() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) TeeReplication() (common.Address, error) {
	return _TeeMachineRegistry.Contract.TeeReplication(&_TeeMachineRegistry.CallOpts)
}

// TeeVerification is a free data retrieval call binding the contract method 0x70dcd962.
//
// Solidity: function teeVerification() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) TeeVerification(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "teeVerification")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeVerification is a free data retrieval call binding the contract method 0x70dcd962.
//
// Solidity: function teeVerification() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistrySession) TeeVerification() (common.Address, error) {
	return _TeeMachineRegistry.Contract.TeeVerification(&_TeeMachineRegistry.CallOpts)
}

// TeeVerification is a free data retrieval call binding the contract method 0x70dcd962.
//
// Solidity: function teeVerification() view returns(address)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) TeeVerification() (common.Address, error) {
	return _TeeMachineRegistry.Contract.TeeVerification(&_TeeMachineRegistry.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeMachineRegistry *TeeMachineRegistryCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeeMachineRegistry.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_TeeMachineRegistry *TeeMachineRegistrySession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeMachineRegistry.Contract.TimelockedCalls(&_TeeMachineRegistry.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeMachineRegistry *TeeMachineRegistryCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeMachineRegistry.Contract.TimelockedCalls(&_TeeMachineRegistry.CallOpts, selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.CancelGovernanceCall(&_TeeMachineRegistry.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.CancelGovernanceCall(&_TeeMachineRegistry.TransactOpts, _selector)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xff63e7a8.
//
// Solidity: function changeStatus(address _teeId, uint8 _newStatus) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) ChangeStatus(opts *bind.TransactOpts, _teeId common.Address, _newStatus uint8) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "changeStatus", _teeId, _newStatus)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xff63e7a8.
//
// Solidity: function changeStatus(address _teeId, uint8 _newStatus) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) ChangeStatus(_teeId common.Address, _newStatus uint8) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.ChangeStatus(&_TeeMachineRegistry.TransactOpts, _teeId, _newStatus)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xff63e7a8.
//
// Solidity: function changeStatus(address _teeId, uint8 _newStatus) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) ChangeStatus(_teeId common.Address, _newStatus uint8) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.ChangeStatus(&_TeeMachineRegistry.TransactOpts, _teeId, _newStatus)
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

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.ExecuteGovernanceCall(&_TeeMachineRegistry.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.ExecuteGovernanceCall(&_TeeMachineRegistry.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Initialise(&_TeeMachineRegistry.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Initialise(&_TeeMachineRegistry.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Initialize(&_TeeMachineRegistry.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Initialize(&_TeeMachineRegistry.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
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

// PauseWithProof is a paid mutator transaction binding the contract method 0x77ab7d69.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) PauseWithProof(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "pauseWithProof", _proof)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0x77ab7d69.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) PauseWithProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.PauseWithProof(&_TeeMachineRegistry.TransactOpts, _proof)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0x77ab7d69.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
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

// Register is a paid mutator transaction binding the contract method 0x87e6164c.
//
// Solidity: function register(uint256 _extensionId, address _teeId, address _teeProxyId, string _url, bytes32 _codeHash, bytes32 _platform) payable returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) Register(opts *bind.TransactOpts, _extensionId *big.Int, _teeId common.Address, _teeProxyId common.Address, _url string, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "register", _extensionId, _teeId, _teeProxyId, _url, _codeHash, _platform)
}

// Register is a paid mutator transaction binding the contract method 0x87e6164c.
//
// Solidity: function register(uint256 _extensionId, address _teeId, address _teeProxyId, string _url, bytes32 _codeHash, bytes32 _platform) payable returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) Register(_extensionId *big.Int, _teeId common.Address, _teeProxyId common.Address, _url string, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Register(&_TeeMachineRegistry.TransactOpts, _extensionId, _teeId, _teeProxyId, _url, _codeHash, _platform)
}

// Register is a paid mutator transaction binding the contract method 0x87e6164c.
//
// Solidity: function register(uint256 _extensionId, address _teeId, address _teeProxyId, string _url, bytes32 _codeHash, bytes32 _platform) payable returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) Register(_extensionId *big.Int, _teeId common.Address, _teeProxyId common.Address, _url string, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Register(&_TeeMachineRegistry.TransactOpts, _extensionId, _teeId, _teeProxyId, _url, _codeHash, _platform)
}

// Replicate is a paid mutator transaction binding the contract method 0x907bde05.
//
// Solidity: function replicate(address _newTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) Replicate(opts *bind.TransactOpts, _newTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "replicate", _newTeeId, _proof)
}

// Replicate is a paid mutator transaction binding the contract method 0x907bde05.
//
// Solidity: function replicate(address _newTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) Replicate(_newTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Replicate(&_TeeMachineRegistry.TransactOpts, _newTeeId, _proof)
}

// Replicate is a paid mutator transaction binding the contract method 0x907bde05.
//
// Solidity: function replicate(address _newTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) Replicate(_newTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.Replicate(&_TeeMachineRegistry.TransactOpts, _newTeeId, _proof)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.SwitchToProductionMode(&_TeeMachineRegistry.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.SwitchToProductionMode(&_TeeMachineRegistry.TransactOpts)
}

// ToProduction is a paid mutator transaction binding the contract method 0x2972de67.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) ToProduction(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "toProduction", _proof)
}

// ToProduction is a paid mutator transaction binding the contract method 0x2972de67.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) ToProduction(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.ToProduction(&_TeeMachineRegistry.TransactOpts, _proof)
}

// ToProduction is a paid mutator transaction binding the contract method 0x2972de67.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,(bytes,bytes32,bytes,bytes32))) _proof) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) ToProduction(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.ToProduction(&_TeeMachineRegistry.TransactOpts, _proof)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.UpdateContractAddresses(&_TeeMachineRegistry.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.UpdateContractAddresses(&_TeeMachineRegistry.TransactOpts, _contractNameHashes, _contractAddresses)
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

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeMachineRegistry.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeMachineRegistry *TeeMachineRegistrySession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.UpgradeToAndCall(&_TeeMachineRegistry.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeMachineRegistry *TeeMachineRegistryTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeMachineRegistry.Contract.UpgradeToAndCall(&_TeeMachineRegistry.TransactOpts, _newImplementation, _data)
}

// TeeMachineRegistryGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryGovernanceCallTimelockedIterator struct {
	Event *TeeMachineRegistryGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryGovernanceCallTimelocked)
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
		it.Event = new(TeeMachineRegistryGovernanceCallTimelocked)
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
func (it *TeeMachineRegistryGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeMachineRegistryGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryGovernanceCallTimelockedIterator{contract: _TeeMachineRegistry.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryGovernanceCallTimelocked)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeMachineRegistryGovernanceCallTimelocked, error) {
	event := new(TeeMachineRegistryGovernanceCallTimelocked)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeMachineRegistryGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryGovernanceInitialisedIterator struct {
	Event *TeeMachineRegistryGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryGovernanceInitialised)
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
		it.Event = new(TeeMachineRegistryGovernanceInitialised)
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
func (it *TeeMachineRegistryGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryGovernanceInitialised represents a GovernanceInitialised event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeMachineRegistryGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryGovernanceInitialisedIterator{contract: _TeeMachineRegistry.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryGovernanceInitialised)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseGovernanceInitialised(log types.Log) (*TeeMachineRegistryGovernanceInitialised, error) {
	event := new(TeeMachineRegistryGovernanceInitialised)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeMachineRegistryGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryGovernedProductionModeEnteredIterator struct {
	Event *TeeMachineRegistryGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryGovernedProductionModeEntered)
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
		it.Event = new(TeeMachineRegistryGovernedProductionModeEntered)
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
func (it *TeeMachineRegistryGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeeMachineRegistryGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryGovernedProductionModeEnteredIterator{contract: _TeeMachineRegistry.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryGovernedProductionModeEntered)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeeMachineRegistryGovernedProductionModeEntered, error) {
	event := new(TeeMachineRegistryGovernedProductionModeEntered)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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

// TeeMachineRegistryTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTimelockedGovernanceCallCanceledIterator struct {
	Event *TeeMachineRegistryTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeeMachineRegistryTimelockedGovernanceCallCanceled)
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
func (it *TeeMachineRegistryTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeeMachineRegistryTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryTimelockedGovernanceCallCanceledIterator{contract: _TeeMachineRegistry.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryTimelockedGovernanceCallCanceled)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeeMachineRegistryTimelockedGovernanceCallCanceled, error) {
	event := new(TeeMachineRegistryTimelockedGovernanceCallCanceled)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeMachineRegistryTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTimelockedGovernanceCallExecutedIterator struct {
	Event *TeeMachineRegistryTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeeMachineRegistryTimelockedGovernanceCallExecuted)
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
func (it *TeeMachineRegistryTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeeMachineRegistryTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryTimelockedGovernanceCallExecutedIterator{contract: _TeeMachineRegistry.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryTimelockedGovernanceCallExecuted)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeeMachineRegistryTimelockedGovernanceCallExecuted, error) {
	event := new(TeeMachineRegistryTimelockedGovernanceCallExecuted)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeMachineRegistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeeMachineRegistry contract.
type TeeMachineRegistryUpgradedIterator struct {
	Event *TeeMachineRegistryUpgraded // Event containing the contract specifics and raw log

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
func (it *TeeMachineRegistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeMachineRegistryUpgraded)
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
		it.Event = new(TeeMachineRegistryUpgraded)
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
func (it *TeeMachineRegistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeMachineRegistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeMachineRegistryUpgraded represents a Upgraded event raised by the TeeMachineRegistry contract.
type TeeMachineRegistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeeMachineRegistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeeMachineRegistryUpgradedIterator{contract: _TeeMachineRegistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeeMachineRegistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeMachineRegistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeMachineRegistryUpgraded)
				if err := _TeeMachineRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeeMachineRegistry *TeeMachineRegistryFilterer) ParseUpgraded(log types.Log) (*TeeMachineRegistryUpgraded, error) {
	event := new(TeeMachineRegistryUpgraded)
	if err := _TeeMachineRegistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
