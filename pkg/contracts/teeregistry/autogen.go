// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teeregistry

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
	State        []byte
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
	StateHash              [32]byte
}

// ITeeRegistryTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
}

// ITeeRegistryTeeMachineWithAttestationData is an auto generated low-level Go binding around an user-defined struct.
type ITeeRegistryTeeMachineWithAttestationData struct {
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

// TeeRegistryMetaData contains all meta data concerning the TeeRegistry contract.
var TeeRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pauseBeforeUpgradeMinDurationSeconds\",\"type\":\"uint256\"}],\"name\":\"PauseBeforeUpgradeMinDurationSecondsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"name\":\"TeeMachinePaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"name\":\"TeeMachinePausedForUpgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"name\":\"TeeMachinePutIntoProduction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"TeeMachineRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldTeeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTeeId\",\"type\":\"address\"}],\"name\":\"TeeMachineReplicationConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldTeeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTeeId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"teeUpgradeId\",\"type\":\"uint256\"}],\"name\":\"TeeMachineReplicationTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"}],\"name\":\"TeeProxyIdSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeId2\",\"type\":\"address\"}],\"name\":\"areTeeMachinesCompatible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTeeId\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"confirmReplicate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveTees\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_urls\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"getRandomTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachine\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structITeeRegistry.TeeMachine\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineStatus\",\"outputs\":[{\"internalType\":\"enumITeeRegistry.TeeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"getTeeMachineWithAttestationData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialTeeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeRegistry.TeeMachineWithAttestationData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"pauseWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oldTeeId\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_teeUpgradeId\",\"type\":\"uint256\"}],\"name\":\"replicateFrom\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeProxyId\",\"type\":\"address\"}],\"name\":\"setTeeProxyId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"}],\"name\":\"toPauseForUpgrade\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFtdcVerification.FtdcSignatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFtdcHub.FtdcResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"challenge\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumITeeAvailabilityCheck.AvailabilityCheckStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"teeTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"initialSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastSigningPolicyId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"stateHash\",\"type\":\"bytes32\"}],\"internalType\":\"structITeeAvailabilityCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"internalType\":\"structITeeAvailabilityCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"toProduction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeRegistryMetaData.ABI instead.
var TeeRegistryABI = TeeRegistryMetaData.ABI

// TeeRegistry is an auto generated Go binding around an Ethereum contract.
type TeeRegistry struct {
	TeeRegistryCaller     // Read-only binding to the contract
	TeeRegistryTransactor // Write-only binding to the contract
	TeeRegistryFilterer   // Log filterer for contract events
}

// TeeRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeRegistrySession struct {
	Contract     *TeeRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeRegistryCallerSession struct {
	Contract *TeeRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TeeRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeRegistryTransactorSession struct {
	Contract     *TeeRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TeeRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeRegistryRaw struct {
	Contract *TeeRegistry // Generic contract binding to access the raw methods on
}

// TeeRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeRegistryCallerRaw struct {
	Contract *TeeRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TeeRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeRegistryTransactorRaw struct {
	Contract *TeeRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeRegistry creates a new instance of TeeRegistry, bound to a specific deployed contract.
func NewTeeRegistry(address common.Address, backend bind.ContractBackend) (*TeeRegistry, error) {
	contract, err := bindTeeRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeRegistry{TeeRegistryCaller: TeeRegistryCaller{contract: contract}, TeeRegistryTransactor: TeeRegistryTransactor{contract: contract}, TeeRegistryFilterer: TeeRegistryFilterer{contract: contract}}, nil
}

// NewTeeRegistryCaller creates a new read-only instance of TeeRegistry, bound to a specific deployed contract.
func NewTeeRegistryCaller(address common.Address, caller bind.ContractCaller) (*TeeRegistryCaller, error) {
	contract, err := bindTeeRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryCaller{contract: contract}, nil
}

// NewTeeRegistryTransactor creates a new write-only instance of TeeRegistry, bound to a specific deployed contract.
func NewTeeRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeRegistryTransactor, error) {
	contract, err := bindTeeRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryTransactor{contract: contract}, nil
}

// NewTeeRegistryFilterer creates a new log filterer instance of TeeRegistry, bound to a specific deployed contract.
func NewTeeRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeRegistryFilterer, error) {
	contract, err := bindTeeRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryFilterer{contract: contract}, nil
}

// bindTeeRegistry binds a generic wrapper to an already deployed contract.
func bindTeeRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeRegistry *TeeRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeRegistry.Contract.TeeRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeRegistry *TeeRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeRegistry.Contract.TeeRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeRegistry *TeeRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeRegistry.Contract.TeeRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeRegistry *TeeRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeRegistry *TeeRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeRegistry *TeeRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeRegistry.Contract.contract.Transact(opts, method, params...)
}

// AreTeeMachinesCompatible is a free data retrieval call binding the contract method 0xef545f84.
//
// Solidity: function areTeeMachinesCompatible(address _teeId1, address _teeId2) view returns(bool)
func (_TeeRegistry *TeeRegistryCaller) AreTeeMachinesCompatible(opts *bind.CallOpts, _teeId1 common.Address, _teeId2 common.Address) (bool, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "areTeeMachinesCompatible", _teeId1, _teeId2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AreTeeMachinesCompatible is a free data retrieval call binding the contract method 0xef545f84.
//
// Solidity: function areTeeMachinesCompatible(address _teeId1, address _teeId2) view returns(bool)
func (_TeeRegistry *TeeRegistrySession) AreTeeMachinesCompatible(_teeId1 common.Address, _teeId2 common.Address) (bool, error) {
	return _TeeRegistry.Contract.AreTeeMachinesCompatible(&_TeeRegistry.CallOpts, _teeId1, _teeId2)
}

// AreTeeMachinesCompatible is a free data retrieval call binding the contract method 0xef545f84.
//
// Solidity: function areTeeMachinesCompatible(address _teeId1, address _teeId2) view returns(bool)
func (_TeeRegistry *TeeRegistryCallerSession) AreTeeMachinesCompatible(_teeId1 common.Address, _teeId2 common.Address) (bool, error) {
	return _TeeRegistry.Contract.AreTeeMachinesCompatible(&_TeeRegistry.CallOpts, _teeId1, _teeId2)
}

// GetActiveTees is a free data retrieval call binding the contract method 0x858e53d7.
//
// Solidity: function getActiveTees() view returns(address[] _teeIds, string[] _urls)
func (_TeeRegistry *TeeRegistryCaller) GetActiveTees(opts *bind.CallOpts) (struct {
	TeeIds []common.Address
	Urls   []string
}, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getActiveTees")

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

// GetActiveTees is a free data retrieval call binding the contract method 0x858e53d7.
//
// Solidity: function getActiveTees() view returns(address[] _teeIds, string[] _urls)
func (_TeeRegistry *TeeRegistrySession) GetActiveTees() (struct {
	TeeIds []common.Address
	Urls   []string
}, error) {
	return _TeeRegistry.Contract.GetActiveTees(&_TeeRegistry.CallOpts)
}

// GetActiveTees is a free data retrieval call binding the contract method 0x858e53d7.
//
// Solidity: function getActiveTees() view returns(address[] _teeIds, string[] _urls)
func (_TeeRegistry *TeeRegistryCallerSession) GetActiveTees() (struct {
	TeeIds []common.Address
	Urls   []string
}, error) {
	return _TeeRegistry.Contract.GetActiveTees(&_TeeRegistry.CallOpts)
}

// GetRandomTeeIds is a free data retrieval call binding the contract method 0xb4db023e.
//
// Solidity: function getRandomTeeIds(uint256 _count) view returns(address[])
func (_TeeRegistry *TeeRegistryCaller) GetRandomTeeIds(opts *bind.CallOpts, _count *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getRandomTeeIds", _count)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRandomTeeIds is a free data retrieval call binding the contract method 0xb4db023e.
//
// Solidity: function getRandomTeeIds(uint256 _count) view returns(address[])
func (_TeeRegistry *TeeRegistrySession) GetRandomTeeIds(_count *big.Int) ([]common.Address, error) {
	return _TeeRegistry.Contract.GetRandomTeeIds(&_TeeRegistry.CallOpts, _count)
}

// GetRandomTeeIds is a free data retrieval call binding the contract method 0xb4db023e.
//
// Solidity: function getRandomTeeIds(uint256 _count) view returns(address[])
func (_TeeRegistry *TeeRegistryCallerSession) GetRandomTeeIds(_count *big.Int) ([]common.Address, error) {
	return _TeeRegistry.Contract.GetRandomTeeIds(&_TeeRegistry.CallOpts, _count)
}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_TeeRegistry *TeeRegistryCaller) GetTeeMachine(opts *bind.CallOpts, _teeId common.Address) (ITeeRegistryTeeMachine, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getTeeMachine", _teeId)

	if err != nil {
		return *new(ITeeRegistryTeeMachine), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeeRegistryTeeMachine)).(*ITeeRegistryTeeMachine)

	return out0, err

}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_TeeRegistry *TeeRegistrySession) GetTeeMachine(_teeId common.Address) (ITeeRegistryTeeMachine, error) {
	return _TeeRegistry.Contract.GetTeeMachine(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachine is a free data retrieval call binding the contract method 0x5ede508f.
//
// Solidity: function getTeeMachine(address _teeId) view returns((address,address,string))
func (_TeeRegistry *TeeRegistryCallerSession) GetTeeMachine(_teeId common.Address) (ITeeRegistryTeeMachine, error) {
	return _TeeRegistry.Contract.GetTeeMachine(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineOwner is a free data retrieval call binding the contract method 0xb7e773e9.
//
// Solidity: function getTeeMachineOwner(address _teeId) view returns(address)
func (_TeeRegistry *TeeRegistryCaller) GetTeeMachineOwner(opts *bind.CallOpts, _teeId common.Address) (common.Address, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getTeeMachineOwner", _teeId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeeMachineOwner is a free data retrieval call binding the contract method 0xb7e773e9.
//
// Solidity: function getTeeMachineOwner(address _teeId) view returns(address)
func (_TeeRegistry *TeeRegistrySession) GetTeeMachineOwner(_teeId common.Address) (common.Address, error) {
	return _TeeRegistry.Contract.GetTeeMachineOwner(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineOwner is a free data retrieval call binding the contract method 0xb7e773e9.
//
// Solidity: function getTeeMachineOwner(address _teeId) view returns(address)
func (_TeeRegistry *TeeRegistryCallerSession) GetTeeMachineOwner(_teeId common.Address) (common.Address, error) {
	return _TeeRegistry.Contract.GetTeeMachineOwner(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_TeeRegistry *TeeRegistryCaller) GetTeeMachineStatus(opts *bind.CallOpts, _teeId common.Address) (uint8, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getTeeMachineStatus", _teeId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_TeeRegistry *TeeRegistrySession) GetTeeMachineStatus(_teeId common.Address) (uint8, error) {
	return _TeeRegistry.Contract.GetTeeMachineStatus(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineStatus is a free data retrieval call binding the contract method 0x25e30221.
//
// Solidity: function getTeeMachineStatus(address _teeId) view returns(uint8)
func (_TeeRegistry *TeeRegistryCallerSession) GetTeeMachineStatus(_teeId common.Address) (uint8, error) {
	return _TeeRegistry.Contract.GetTeeMachineStatus(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_TeeRegistry *TeeRegistryCaller) GetTeeMachineWithAttestationData(opts *bind.CallOpts, _teeId common.Address) (ITeeRegistryTeeMachineWithAttestationData, error) {
	var out []interface{}
	err := _TeeRegistry.contract.Call(opts, &out, "getTeeMachineWithAttestationData", _teeId)

	if err != nil {
		return *new(ITeeRegistryTeeMachineWithAttestationData), err
	}

	out0 := *abi.ConvertType(out[0], new(ITeeRegistryTeeMachineWithAttestationData)).(*ITeeRegistryTeeMachineWithAttestationData)

	return out0, err

}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_TeeRegistry *TeeRegistrySession) GetTeeMachineWithAttestationData(_teeId common.Address) (ITeeRegistryTeeMachineWithAttestationData, error) {
	return _TeeRegistry.Contract.GetTeeMachineWithAttestationData(&_TeeRegistry.CallOpts, _teeId)
}

// GetTeeMachineWithAttestationData is a free data retrieval call binding the contract method 0x883e6c9a.
//
// Solidity: function getTeeMachineWithAttestationData(address _teeId) view returns((address,address,string,bytes32,bytes32))
func (_TeeRegistry *TeeRegistryCallerSession) GetTeeMachineWithAttestationData(_teeId common.Address) (ITeeRegistryTeeMachineWithAttestationData, error) {
	return _TeeRegistry.Contract.GetTeeMachineWithAttestationData(&_TeeRegistry.CallOpts, _teeId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x80f2abe3.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (_TeeRegistry *TeeRegistryTransactor) ConfirmOwnership(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "confirmOwnership", _teeId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x80f2abe3.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (_TeeRegistry *TeeRegistrySession) ConfirmOwnership(_teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ConfirmOwnership(&_TeeRegistry.TransactOpts, _teeId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x80f2abe3.
//
// Solidity: function confirmOwnership(address _teeId) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ConfirmOwnership(_teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ConfirmOwnership(&_TeeRegistry.TransactOpts, _teeId)
}

// ConfirmReplicate is a paid mutator transaction binding the contract method 0x78775d5b.
//
// Solidity: function confirmReplicate(address _newTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,bytes32),bytes) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactor) ConfirmReplicate(opts *bind.TransactOpts, _newTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "confirmReplicate", _newTeeId, _proof)
}

// ConfirmReplicate is a paid mutator transaction binding the contract method 0x78775d5b.
//
// Solidity: function confirmReplicate(address _newTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,bytes32),bytes) _proof) returns()
func (_TeeRegistry *TeeRegistrySession) ConfirmReplicate(_newTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ConfirmReplicate(&_TeeRegistry.TransactOpts, _newTeeId, _proof)
}

// ConfirmReplicate is a paid mutator transaction binding the contract method 0x78775d5b.
//
// Solidity: function confirmReplicate(address _newTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,bytes32),bytes) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ConfirmReplicate(_newTeeId common.Address, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ConfirmReplicate(&_TeeRegistry.TransactOpts, _newTeeId, _proof)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _teeId) returns()
func (_TeeRegistry *TeeRegistryTransactor) Pause(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "pause", _teeId)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _teeId) returns()
func (_TeeRegistry *TeeRegistrySession) Pause(_teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.Pause(&_TeeRegistry.TransactOpts, _teeId)
}

// Pause is a paid mutator transaction binding the contract method 0x76a67a51.
//
// Solidity: function pause(address _teeId) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) Pause(_teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.Pause(&_TeeRegistry.TransactOpts, _teeId)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0xa90147fa.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,bytes32),bytes) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactor) PauseWithProof(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "pauseWithProof", _proof)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0xa90147fa.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,bytes32),bytes) _proof) returns()
func (_TeeRegistry *TeeRegistrySession) PauseWithProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.PauseWithProof(&_TeeRegistry.TransactOpts, _proof)
}

// PauseWithProof is a paid mutator transaction binding the contract method 0xa90147fa.
//
// Solidity: function pauseWithProof(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,bytes32),bytes) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) PauseWithProof(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.PauseWithProof(&_TeeRegistry.TransactOpts, _proof)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x38195139.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (_TeeRegistry *TeeRegistryTransactor) ProposeNewOwner(opts *bind.TransactOpts, _teeId common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "proposeNewOwner", _teeId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x38195139.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (_TeeRegistry *TeeRegistrySession) ProposeNewOwner(_teeId common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ProposeNewOwner(&_TeeRegistry.TransactOpts, _teeId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0x38195139.
//
// Solidity: function proposeNewOwner(address _teeId, address _newOwner) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ProposeNewOwner(_teeId common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ProposeNewOwner(&_TeeRegistry.TransactOpts, _teeId, _newOwner)
}

// Register is a paid mutator transaction binding the contract method 0xf56a8a97.
//
// Solidity: function register(address _teeId, address _teeProxyId, string _url, bytes32 _codeHash, bytes32 _platform) payable returns()
func (_TeeRegistry *TeeRegistryTransactor) Register(opts *bind.TransactOpts, _teeId common.Address, _teeProxyId common.Address, _url string, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "register", _teeId, _teeProxyId, _url, _codeHash, _platform)
}

// Register is a paid mutator transaction binding the contract method 0xf56a8a97.
//
// Solidity: function register(address _teeId, address _teeProxyId, string _url, bytes32 _codeHash, bytes32 _platform) payable returns()
func (_TeeRegistry *TeeRegistrySession) Register(_teeId common.Address, _teeProxyId common.Address, _url string, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.Register(&_TeeRegistry.TransactOpts, _teeId, _teeProxyId, _url, _codeHash, _platform)
}

// Register is a paid mutator transaction binding the contract method 0xf56a8a97.
//
// Solidity: function register(address _teeId, address _teeProxyId, string _url, bytes32 _codeHash, bytes32 _platform) payable returns()
func (_TeeRegistry *TeeRegistryTransactorSession) Register(_teeId common.Address, _teeProxyId common.Address, _url string, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _TeeRegistry.Contract.Register(&_TeeRegistry.TransactOpts, _teeId, _teeProxyId, _url, _codeHash, _platform)
}

// ReplicateFrom is a paid mutator transaction binding the contract method 0x5e5711ef.
//
// Solidity: function replicateFrom(address _oldTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,bytes32),bytes) _proof, uint256 _teeUpgradeId) payable returns()
func (_TeeRegistry *TeeRegistryTransactor) ReplicateFrom(opts *bind.TransactOpts, _oldTeeId common.Address, _proof ITeeAvailabilityCheckProof, _teeUpgradeId *big.Int) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "replicateFrom", _oldTeeId, _proof, _teeUpgradeId)
}

// ReplicateFrom is a paid mutator transaction binding the contract method 0x5e5711ef.
//
// Solidity: function replicateFrom(address _oldTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,bytes32),bytes) _proof, uint256 _teeUpgradeId) payable returns()
func (_TeeRegistry *TeeRegistrySession) ReplicateFrom(_oldTeeId common.Address, _proof ITeeAvailabilityCheckProof, _teeUpgradeId *big.Int) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ReplicateFrom(&_TeeRegistry.TransactOpts, _oldTeeId, _proof, _teeUpgradeId)
}

// ReplicateFrom is a paid mutator transaction binding the contract method 0x5e5711ef.
//
// Solidity: function replicateFrom(address _oldTeeId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,bytes32),bytes) _proof, uint256 _teeUpgradeId) payable returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ReplicateFrom(_oldTeeId common.Address, _proof ITeeAvailabilityCheckProof, _teeUpgradeId *big.Int) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ReplicateFrom(&_TeeRegistry.TransactOpts, _oldTeeId, _proof, _teeUpgradeId)
}

// SetTeeProxyId is a paid mutator transaction binding the contract method 0x97538b38.
//
// Solidity: function setTeeProxyId(address _teeId, address _teeProxyId) returns()
func (_TeeRegistry *TeeRegistryTransactor) SetTeeProxyId(opts *bind.TransactOpts, _teeId common.Address, _teeProxyId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "setTeeProxyId", _teeId, _teeProxyId)
}

// SetTeeProxyId is a paid mutator transaction binding the contract method 0x97538b38.
//
// Solidity: function setTeeProxyId(address _teeId, address _teeProxyId) returns()
func (_TeeRegistry *TeeRegistrySession) SetTeeProxyId(_teeId common.Address, _teeProxyId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.SetTeeProxyId(&_TeeRegistry.TransactOpts, _teeId, _teeProxyId)
}

// SetTeeProxyId is a paid mutator transaction binding the contract method 0x97538b38.
//
// Solidity: function setTeeProxyId(address _teeId, address _teeProxyId) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) SetTeeProxyId(_teeId common.Address, _teeProxyId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.SetTeeProxyId(&_TeeRegistry.TransactOpts, _teeId, _teeProxyId)
}

// ToPauseForUpgrade is a paid mutator transaction binding the contract method 0xe143ca78.
//
// Solidity: function toPauseForUpgrade(address _teeId) payable returns()
func (_TeeRegistry *TeeRegistryTransactor) ToPauseForUpgrade(opts *bind.TransactOpts, _teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "toPauseForUpgrade", _teeId)
}

// ToPauseForUpgrade is a paid mutator transaction binding the contract method 0xe143ca78.
//
// Solidity: function toPauseForUpgrade(address _teeId) payable returns()
func (_TeeRegistry *TeeRegistrySession) ToPauseForUpgrade(_teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ToPauseForUpgrade(&_TeeRegistry.TransactOpts, _teeId)
}

// ToPauseForUpgrade is a paid mutator transaction binding the contract method 0xe143ca78.
//
// Solidity: function toPauseForUpgrade(address _teeId) payable returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ToPauseForUpgrade(_teeId common.Address) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ToPauseForUpgrade(&_TeeRegistry.TransactOpts, _teeId)
}

// ToProduction is a paid mutator transaction binding the contract method 0xf54e3975.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,bytes32),bytes) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactor) ToProduction(opts *bind.TransactOpts, _proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.contract.Transact(opts, "toProduction", _proof)
}

// ToProduction is a paid mutator transaction binding the contract method 0xf54e3975.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,bytes32),bytes) _proof) returns()
func (_TeeRegistry *TeeRegistrySession) ToProduction(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ToProduction(&_TeeRegistry.TransactOpts, _proof)
}

// ToProduction is a paid mutator transaction binding the contract method 0xf54e3975.
//
// Solidity: function toProduction(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address[],uint64,uint64),(address,string,bytes32),(uint8,uint64,bytes32,bytes32,uint32,uint32,bytes32),bytes) _proof) returns()
func (_TeeRegistry *TeeRegistryTransactorSession) ToProduction(_proof ITeeAvailabilityCheckProof) (*types.Transaction, error) {
	return _TeeRegistry.Contract.ToProduction(&_TeeRegistry.TransactOpts, _proof)
}

// TeeRegistryNewOwnerConfirmedIterator is returned from FilterNewOwnerConfirmed and is used to iterate over the raw logs and unpacked data for NewOwnerConfirmed events raised by the TeeRegistry contract.
type TeeRegistryNewOwnerConfirmedIterator struct {
	Event *TeeRegistryNewOwnerConfirmed // Event containing the contract specifics and raw log

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
func (it *TeeRegistryNewOwnerConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryNewOwnerConfirmed)
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
		it.Event = new(TeeRegistryNewOwnerConfirmed)
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
func (it *TeeRegistryNewOwnerConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryNewOwnerConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryNewOwnerConfirmed represents a NewOwnerConfirmed event raised by the TeeRegistry contract.
type TeeRegistryNewOwnerConfirmed struct {
	TeeId    common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerConfirmed is a free log retrieval operation binding the contract event 0xc813c793f67554890ac6f92415119ac4a7535162dc38fb1120036cbd62069fbf.
//
// Solidity: event NewOwnerConfirmed(address indexed teeId, address indexed newOwner)
func (_TeeRegistry *TeeRegistryFilterer) FilterNewOwnerConfirmed(opts *bind.FilterOpts, teeId []common.Address, newOwner []common.Address) (*TeeRegistryNewOwnerConfirmedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "NewOwnerConfirmed", teeIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryNewOwnerConfirmedIterator{contract: _TeeRegistry.contract, event: "NewOwnerConfirmed", logs: logs, sub: sub}, nil
}

// WatchNewOwnerConfirmed is a free log subscription operation binding the contract event 0xc813c793f67554890ac6f92415119ac4a7535162dc38fb1120036cbd62069fbf.
//
// Solidity: event NewOwnerConfirmed(address indexed teeId, address indexed newOwner)
func (_TeeRegistry *TeeRegistryFilterer) WatchNewOwnerConfirmed(opts *bind.WatchOpts, sink chan<- *TeeRegistryNewOwnerConfirmed, teeId []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "NewOwnerConfirmed", teeIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryNewOwnerConfirmed)
				if err := _TeeRegistry.contract.UnpackLog(event, "NewOwnerConfirmed", log); err != nil {
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
func (_TeeRegistry *TeeRegistryFilterer) ParseNewOwnerConfirmed(log types.Log) (*TeeRegistryNewOwnerConfirmed, error) {
	event := new(TeeRegistryNewOwnerConfirmed)
	if err := _TeeRegistry.contract.UnpackLog(event, "NewOwnerConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRegistryNewOwnerProposedIterator is returned from FilterNewOwnerProposed and is used to iterate over the raw logs and unpacked data for NewOwnerProposed events raised by the TeeRegistry contract.
type TeeRegistryNewOwnerProposedIterator struct {
	Event *TeeRegistryNewOwnerProposed // Event containing the contract specifics and raw log

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
func (it *TeeRegistryNewOwnerProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryNewOwnerProposed)
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
		it.Event = new(TeeRegistryNewOwnerProposed)
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
func (it *TeeRegistryNewOwnerProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryNewOwnerProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryNewOwnerProposed represents a NewOwnerProposed event raised by the TeeRegistry contract.
type TeeRegistryNewOwnerProposed struct {
	TeeId    common.Address
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerProposed is a free log retrieval operation binding the contract event 0x64420d4a41c6ed4de2bccbf33192eea18e576c5b23c79c3a722d4e9534c2e8d8.
//
// Solidity: event NewOwnerProposed(address indexed teeId, address indexed oldOwner, address indexed newOwner)
func (_TeeRegistry *TeeRegistryFilterer) FilterNewOwnerProposed(opts *bind.FilterOpts, teeId []common.Address, oldOwner []common.Address, newOwner []common.Address) (*TeeRegistryNewOwnerProposedIterator, error) {

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

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "NewOwnerProposed", teeIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryNewOwnerProposedIterator{contract: _TeeRegistry.contract, event: "NewOwnerProposed", logs: logs, sub: sub}, nil
}

// WatchNewOwnerProposed is a free log subscription operation binding the contract event 0x64420d4a41c6ed4de2bccbf33192eea18e576c5b23c79c3a722d4e9534c2e8d8.
//
// Solidity: event NewOwnerProposed(address indexed teeId, address indexed oldOwner, address indexed newOwner)
func (_TeeRegistry *TeeRegistryFilterer) WatchNewOwnerProposed(opts *bind.WatchOpts, sink chan<- *TeeRegistryNewOwnerProposed, teeId []common.Address, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "NewOwnerProposed", teeIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryNewOwnerProposed)
				if err := _TeeRegistry.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
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
func (_TeeRegistry *TeeRegistryFilterer) ParseNewOwnerProposed(log types.Log) (*TeeRegistryNewOwnerProposed, error) {
	event := new(TeeRegistryNewOwnerProposed)
	if err := _TeeRegistry.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRegistryPauseBeforeUpgradeMinDurationSecondsSetIterator is returned from FilterPauseBeforeUpgradeMinDurationSecondsSet and is used to iterate over the raw logs and unpacked data for PauseBeforeUpgradeMinDurationSecondsSet events raised by the TeeRegistry contract.
type TeeRegistryPauseBeforeUpgradeMinDurationSecondsSetIterator struct {
	Event *TeeRegistryPauseBeforeUpgradeMinDurationSecondsSet // Event containing the contract specifics and raw log

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
func (it *TeeRegistryPauseBeforeUpgradeMinDurationSecondsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryPauseBeforeUpgradeMinDurationSecondsSet)
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
		it.Event = new(TeeRegistryPauseBeforeUpgradeMinDurationSecondsSet)
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
func (it *TeeRegistryPauseBeforeUpgradeMinDurationSecondsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryPauseBeforeUpgradeMinDurationSecondsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryPauseBeforeUpgradeMinDurationSecondsSet represents a PauseBeforeUpgradeMinDurationSecondsSet event raised by the TeeRegistry contract.
type TeeRegistryPauseBeforeUpgradeMinDurationSecondsSet struct {
	PauseBeforeUpgradeMinDurationSeconds *big.Int
	Raw                                  types.Log // Blockchain specific contextual infos
}

// FilterPauseBeforeUpgradeMinDurationSecondsSet is a free log retrieval operation binding the contract event 0x90d47fa68d02b1c17231c83feaa930838b7cad77a1a9b948fe0a835ba52b16e0.
//
// Solidity: event PauseBeforeUpgradeMinDurationSecondsSet(uint256 pauseBeforeUpgradeMinDurationSeconds)
func (_TeeRegistry *TeeRegistryFilterer) FilterPauseBeforeUpgradeMinDurationSecondsSet(opts *bind.FilterOpts) (*TeeRegistryPauseBeforeUpgradeMinDurationSecondsSetIterator, error) {

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "PauseBeforeUpgradeMinDurationSecondsSet")
	if err != nil {
		return nil, err
	}
	return &TeeRegistryPauseBeforeUpgradeMinDurationSecondsSetIterator{contract: _TeeRegistry.contract, event: "PauseBeforeUpgradeMinDurationSecondsSet", logs: logs, sub: sub}, nil
}

// WatchPauseBeforeUpgradeMinDurationSecondsSet is a free log subscription operation binding the contract event 0x90d47fa68d02b1c17231c83feaa930838b7cad77a1a9b948fe0a835ba52b16e0.
//
// Solidity: event PauseBeforeUpgradeMinDurationSecondsSet(uint256 pauseBeforeUpgradeMinDurationSeconds)
func (_TeeRegistry *TeeRegistryFilterer) WatchPauseBeforeUpgradeMinDurationSecondsSet(opts *bind.WatchOpts, sink chan<- *TeeRegistryPauseBeforeUpgradeMinDurationSecondsSet) (event.Subscription, error) {

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "PauseBeforeUpgradeMinDurationSecondsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryPauseBeforeUpgradeMinDurationSecondsSet)
				if err := _TeeRegistry.contract.UnpackLog(event, "PauseBeforeUpgradeMinDurationSecondsSet", log); err != nil {
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

// ParsePauseBeforeUpgradeMinDurationSecondsSet is a log parse operation binding the contract event 0x90d47fa68d02b1c17231c83feaa930838b7cad77a1a9b948fe0a835ba52b16e0.
//
// Solidity: event PauseBeforeUpgradeMinDurationSecondsSet(uint256 pauseBeforeUpgradeMinDurationSeconds)
func (_TeeRegistry *TeeRegistryFilterer) ParsePauseBeforeUpgradeMinDurationSecondsSet(log types.Log) (*TeeRegistryPauseBeforeUpgradeMinDurationSecondsSet, error) {
	event := new(TeeRegistryPauseBeforeUpgradeMinDurationSecondsSet)
	if err := _TeeRegistry.contract.UnpackLog(event, "PauseBeforeUpgradeMinDurationSecondsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRegistryTeeMachinePausedIterator is returned from FilterTeeMachinePaused and is used to iterate over the raw logs and unpacked data for TeeMachinePaused events raised by the TeeRegistry contract.
type TeeRegistryTeeMachinePausedIterator struct {
	Event *TeeRegistryTeeMachinePaused // Event containing the contract specifics and raw log

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
func (it *TeeRegistryTeeMachinePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryTeeMachinePaused)
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
		it.Event = new(TeeRegistryTeeMachinePaused)
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
func (it *TeeRegistryTeeMachinePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryTeeMachinePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryTeeMachinePaused represents a TeeMachinePaused event raised by the TeeRegistry contract.
type TeeRegistryTeeMachinePaused struct {
	TeeId common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTeeMachinePaused is a free log retrieval operation binding the contract event 0xde5ae0e0f53bc200068db0c7d41dd63811b465bfd5396cf26a087ec901ad81f4.
//
// Solidity: event TeeMachinePaused(address indexed teeId)
func (_TeeRegistry *TeeRegistryFilterer) FilterTeeMachinePaused(opts *bind.FilterOpts, teeId []common.Address) (*TeeRegistryTeeMachinePausedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "TeeMachinePaused", teeIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryTeeMachinePausedIterator{contract: _TeeRegistry.contract, event: "TeeMachinePaused", logs: logs, sub: sub}, nil
}

// WatchTeeMachinePaused is a free log subscription operation binding the contract event 0xde5ae0e0f53bc200068db0c7d41dd63811b465bfd5396cf26a087ec901ad81f4.
//
// Solidity: event TeeMachinePaused(address indexed teeId)
func (_TeeRegistry *TeeRegistryFilterer) WatchTeeMachinePaused(opts *bind.WatchOpts, sink chan<- *TeeRegistryTeeMachinePaused, teeId []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "TeeMachinePaused", teeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryTeeMachinePaused)
				if err := _TeeRegistry.contract.UnpackLog(event, "TeeMachinePaused", log); err != nil {
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

// ParseTeeMachinePaused is a log parse operation binding the contract event 0xde5ae0e0f53bc200068db0c7d41dd63811b465bfd5396cf26a087ec901ad81f4.
//
// Solidity: event TeeMachinePaused(address indexed teeId)
func (_TeeRegistry *TeeRegistryFilterer) ParseTeeMachinePaused(log types.Log) (*TeeRegistryTeeMachinePaused, error) {
	event := new(TeeRegistryTeeMachinePaused)
	if err := _TeeRegistry.contract.UnpackLog(event, "TeeMachinePaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRegistryTeeMachinePausedForUpgradeIterator is returned from FilterTeeMachinePausedForUpgrade and is used to iterate over the raw logs and unpacked data for TeeMachinePausedForUpgrade events raised by the TeeRegistry contract.
type TeeRegistryTeeMachinePausedForUpgradeIterator struct {
	Event *TeeRegistryTeeMachinePausedForUpgrade // Event containing the contract specifics and raw log

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
func (it *TeeRegistryTeeMachinePausedForUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryTeeMachinePausedForUpgrade)
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
		it.Event = new(TeeRegistryTeeMachinePausedForUpgrade)
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
func (it *TeeRegistryTeeMachinePausedForUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryTeeMachinePausedForUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryTeeMachinePausedForUpgrade represents a TeeMachinePausedForUpgrade event raised by the TeeRegistry contract.
type TeeRegistryTeeMachinePausedForUpgrade struct {
	TeeId common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTeeMachinePausedForUpgrade is a free log retrieval operation binding the contract event 0x9956cba12d00f65e289543db454a9d2065e977680aede4b78298feef6b196e26.
//
// Solidity: event TeeMachinePausedForUpgrade(address indexed teeId)
func (_TeeRegistry *TeeRegistryFilterer) FilterTeeMachinePausedForUpgrade(opts *bind.FilterOpts, teeId []common.Address) (*TeeRegistryTeeMachinePausedForUpgradeIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "TeeMachinePausedForUpgrade", teeIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryTeeMachinePausedForUpgradeIterator{contract: _TeeRegistry.contract, event: "TeeMachinePausedForUpgrade", logs: logs, sub: sub}, nil
}

// WatchTeeMachinePausedForUpgrade is a free log subscription operation binding the contract event 0x9956cba12d00f65e289543db454a9d2065e977680aede4b78298feef6b196e26.
//
// Solidity: event TeeMachinePausedForUpgrade(address indexed teeId)
func (_TeeRegistry *TeeRegistryFilterer) WatchTeeMachinePausedForUpgrade(opts *bind.WatchOpts, sink chan<- *TeeRegistryTeeMachinePausedForUpgrade, teeId []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "TeeMachinePausedForUpgrade", teeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryTeeMachinePausedForUpgrade)
				if err := _TeeRegistry.contract.UnpackLog(event, "TeeMachinePausedForUpgrade", log); err != nil {
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

// ParseTeeMachinePausedForUpgrade is a log parse operation binding the contract event 0x9956cba12d00f65e289543db454a9d2065e977680aede4b78298feef6b196e26.
//
// Solidity: event TeeMachinePausedForUpgrade(address indexed teeId)
func (_TeeRegistry *TeeRegistryFilterer) ParseTeeMachinePausedForUpgrade(log types.Log) (*TeeRegistryTeeMachinePausedForUpgrade, error) {
	event := new(TeeRegistryTeeMachinePausedForUpgrade)
	if err := _TeeRegistry.contract.UnpackLog(event, "TeeMachinePausedForUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRegistryTeeMachinePutIntoProductionIterator is returned from FilterTeeMachinePutIntoProduction and is used to iterate over the raw logs and unpacked data for TeeMachinePutIntoProduction events raised by the TeeRegistry contract.
type TeeRegistryTeeMachinePutIntoProductionIterator struct {
	Event *TeeRegistryTeeMachinePutIntoProduction // Event containing the contract specifics and raw log

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
func (it *TeeRegistryTeeMachinePutIntoProductionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryTeeMachinePutIntoProduction)
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
		it.Event = new(TeeRegistryTeeMachinePutIntoProduction)
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
func (it *TeeRegistryTeeMachinePutIntoProductionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryTeeMachinePutIntoProductionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryTeeMachinePutIntoProduction represents a TeeMachinePutIntoProduction event raised by the TeeRegistry contract.
type TeeRegistryTeeMachinePutIntoProduction struct {
	TeeId common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTeeMachinePutIntoProduction is a free log retrieval operation binding the contract event 0x15c68e0e26236bae5dcac76be19be673430ec7f513cda79db7739b24270f3e69.
//
// Solidity: event TeeMachinePutIntoProduction(address indexed teeId)
func (_TeeRegistry *TeeRegistryFilterer) FilterTeeMachinePutIntoProduction(opts *bind.FilterOpts, teeId []common.Address) (*TeeRegistryTeeMachinePutIntoProductionIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "TeeMachinePutIntoProduction", teeIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryTeeMachinePutIntoProductionIterator{contract: _TeeRegistry.contract, event: "TeeMachinePutIntoProduction", logs: logs, sub: sub}, nil
}

// WatchTeeMachinePutIntoProduction is a free log subscription operation binding the contract event 0x15c68e0e26236bae5dcac76be19be673430ec7f513cda79db7739b24270f3e69.
//
// Solidity: event TeeMachinePutIntoProduction(address indexed teeId)
func (_TeeRegistry *TeeRegistryFilterer) WatchTeeMachinePutIntoProduction(opts *bind.WatchOpts, sink chan<- *TeeRegistryTeeMachinePutIntoProduction, teeId []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "TeeMachinePutIntoProduction", teeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryTeeMachinePutIntoProduction)
				if err := _TeeRegistry.contract.UnpackLog(event, "TeeMachinePutIntoProduction", log); err != nil {
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

// ParseTeeMachinePutIntoProduction is a log parse operation binding the contract event 0x15c68e0e26236bae5dcac76be19be673430ec7f513cda79db7739b24270f3e69.
//
// Solidity: event TeeMachinePutIntoProduction(address indexed teeId)
func (_TeeRegistry *TeeRegistryFilterer) ParseTeeMachinePutIntoProduction(log types.Log) (*TeeRegistryTeeMachinePutIntoProduction, error) {
	event := new(TeeRegistryTeeMachinePutIntoProduction)
	if err := _TeeRegistry.contract.UnpackLog(event, "TeeMachinePutIntoProduction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRegistryTeeMachineRegisteredIterator is returned from FilterTeeMachineRegistered and is used to iterate over the raw logs and unpacked data for TeeMachineRegistered events raised by the TeeRegistry contract.
type TeeRegistryTeeMachineRegisteredIterator struct {
	Event *TeeRegistryTeeMachineRegistered // Event containing the contract specifics and raw log

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
func (it *TeeRegistryTeeMachineRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryTeeMachineRegistered)
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
		it.Event = new(TeeRegistryTeeMachineRegistered)
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
func (it *TeeRegistryTeeMachineRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryTeeMachineRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryTeeMachineRegistered represents a TeeMachineRegistered event raised by the TeeRegistry contract.
type TeeRegistryTeeMachineRegistered struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Owner      common.Address
	Url        string
	CodeHash   [32]byte
	Platform   [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeeMachineRegistered is a free log retrieval operation binding the contract event 0x1fc3a1800afade92ba310b5a97338ddf9b625a37e36ac7c1cfd9d2721d1c01cc.
//
// Solidity: event TeeMachineRegistered(address indexed teeId, address indexed teeProxyId, address indexed owner, string url, bytes32 codeHash, bytes32 platform)
func (_TeeRegistry *TeeRegistryFilterer) FilterTeeMachineRegistered(opts *bind.FilterOpts, teeId []common.Address, teeProxyId []common.Address, owner []common.Address) (*TeeRegistryTeeMachineRegisteredIterator, error) {

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

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "TeeMachineRegistered", teeIdRule, teeProxyIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryTeeMachineRegisteredIterator{contract: _TeeRegistry.contract, event: "TeeMachineRegistered", logs: logs, sub: sub}, nil
}

// WatchTeeMachineRegistered is a free log subscription operation binding the contract event 0x1fc3a1800afade92ba310b5a97338ddf9b625a37e36ac7c1cfd9d2721d1c01cc.
//
// Solidity: event TeeMachineRegistered(address indexed teeId, address indexed teeProxyId, address indexed owner, string url, bytes32 codeHash, bytes32 platform)
func (_TeeRegistry *TeeRegistryFilterer) WatchTeeMachineRegistered(opts *bind.WatchOpts, sink chan<- *TeeRegistryTeeMachineRegistered, teeId []common.Address, teeProxyId []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "TeeMachineRegistered", teeIdRule, teeProxyIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryTeeMachineRegistered)
				if err := _TeeRegistry.contract.UnpackLog(event, "TeeMachineRegistered", log); err != nil {
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

// ParseTeeMachineRegistered is a log parse operation binding the contract event 0x1fc3a1800afade92ba310b5a97338ddf9b625a37e36ac7c1cfd9d2721d1c01cc.
//
// Solidity: event TeeMachineRegistered(address indexed teeId, address indexed teeProxyId, address indexed owner, string url, bytes32 codeHash, bytes32 platform)
func (_TeeRegistry *TeeRegistryFilterer) ParseTeeMachineRegistered(log types.Log) (*TeeRegistryTeeMachineRegistered, error) {
	event := new(TeeRegistryTeeMachineRegistered)
	if err := _TeeRegistry.contract.UnpackLog(event, "TeeMachineRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRegistryTeeMachineReplicationConfirmedIterator is returned from FilterTeeMachineReplicationConfirmed and is used to iterate over the raw logs and unpacked data for TeeMachineReplicationConfirmed events raised by the TeeRegistry contract.
type TeeRegistryTeeMachineReplicationConfirmedIterator struct {
	Event *TeeRegistryTeeMachineReplicationConfirmed // Event containing the contract specifics and raw log

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
func (it *TeeRegistryTeeMachineReplicationConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryTeeMachineReplicationConfirmed)
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
		it.Event = new(TeeRegistryTeeMachineReplicationConfirmed)
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
func (it *TeeRegistryTeeMachineReplicationConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryTeeMachineReplicationConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryTeeMachineReplicationConfirmed represents a TeeMachineReplicationConfirmed event raised by the TeeRegistry contract.
type TeeRegistryTeeMachineReplicationConfirmed struct {
	OldTeeId common.Address
	NewTeeId common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTeeMachineReplicationConfirmed is a free log retrieval operation binding the contract event 0xadc0656d4db250edd81c851803025dab7b77ed8409cb0dffa5c9a417edd1e5c6.
//
// Solidity: event TeeMachineReplicationConfirmed(address indexed oldTeeId, address indexed newTeeId)
func (_TeeRegistry *TeeRegistryFilterer) FilterTeeMachineReplicationConfirmed(opts *bind.FilterOpts, oldTeeId []common.Address, newTeeId []common.Address) (*TeeRegistryTeeMachineReplicationConfirmedIterator, error) {

	var oldTeeIdRule []interface{}
	for _, oldTeeIdItem := range oldTeeId {
		oldTeeIdRule = append(oldTeeIdRule, oldTeeIdItem)
	}
	var newTeeIdRule []interface{}
	for _, newTeeIdItem := range newTeeId {
		newTeeIdRule = append(newTeeIdRule, newTeeIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "TeeMachineReplicationConfirmed", oldTeeIdRule, newTeeIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryTeeMachineReplicationConfirmedIterator{contract: _TeeRegistry.contract, event: "TeeMachineReplicationConfirmed", logs: logs, sub: sub}, nil
}

// WatchTeeMachineReplicationConfirmed is a free log subscription operation binding the contract event 0xadc0656d4db250edd81c851803025dab7b77ed8409cb0dffa5c9a417edd1e5c6.
//
// Solidity: event TeeMachineReplicationConfirmed(address indexed oldTeeId, address indexed newTeeId)
func (_TeeRegistry *TeeRegistryFilterer) WatchTeeMachineReplicationConfirmed(opts *bind.WatchOpts, sink chan<- *TeeRegistryTeeMachineReplicationConfirmed, oldTeeId []common.Address, newTeeId []common.Address) (event.Subscription, error) {

	var oldTeeIdRule []interface{}
	for _, oldTeeIdItem := range oldTeeId {
		oldTeeIdRule = append(oldTeeIdRule, oldTeeIdItem)
	}
	var newTeeIdRule []interface{}
	for _, newTeeIdItem := range newTeeId {
		newTeeIdRule = append(newTeeIdRule, newTeeIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "TeeMachineReplicationConfirmed", oldTeeIdRule, newTeeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryTeeMachineReplicationConfirmed)
				if err := _TeeRegistry.contract.UnpackLog(event, "TeeMachineReplicationConfirmed", log); err != nil {
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

// ParseTeeMachineReplicationConfirmed is a log parse operation binding the contract event 0xadc0656d4db250edd81c851803025dab7b77ed8409cb0dffa5c9a417edd1e5c6.
//
// Solidity: event TeeMachineReplicationConfirmed(address indexed oldTeeId, address indexed newTeeId)
func (_TeeRegistry *TeeRegistryFilterer) ParseTeeMachineReplicationConfirmed(log types.Log) (*TeeRegistryTeeMachineReplicationConfirmed, error) {
	event := new(TeeRegistryTeeMachineReplicationConfirmed)
	if err := _TeeRegistry.contract.UnpackLog(event, "TeeMachineReplicationConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRegistryTeeMachineReplicationTriggeredIterator is returned from FilterTeeMachineReplicationTriggered and is used to iterate over the raw logs and unpacked data for TeeMachineReplicationTriggered events raised by the TeeRegistry contract.
type TeeRegistryTeeMachineReplicationTriggeredIterator struct {
	Event *TeeRegistryTeeMachineReplicationTriggered // Event containing the contract specifics and raw log

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
func (it *TeeRegistryTeeMachineReplicationTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryTeeMachineReplicationTriggered)
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
		it.Event = new(TeeRegistryTeeMachineReplicationTriggered)
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
func (it *TeeRegistryTeeMachineReplicationTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryTeeMachineReplicationTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryTeeMachineReplicationTriggered represents a TeeMachineReplicationTriggered event raised by the TeeRegistry contract.
type TeeRegistryTeeMachineReplicationTriggered struct {
	OldTeeId     common.Address
	NewTeeId     common.Address
	TeeUpgradeId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTeeMachineReplicationTriggered is a free log retrieval operation binding the contract event 0xec7149dab14e163102ba012e316ce7c830a771920fe7f043e34f45d10dd928ad.
//
// Solidity: event TeeMachineReplicationTriggered(address indexed oldTeeId, address indexed newTeeId, uint256 teeUpgradeId)
func (_TeeRegistry *TeeRegistryFilterer) FilterTeeMachineReplicationTriggered(opts *bind.FilterOpts, oldTeeId []common.Address, newTeeId []common.Address) (*TeeRegistryTeeMachineReplicationTriggeredIterator, error) {

	var oldTeeIdRule []interface{}
	for _, oldTeeIdItem := range oldTeeId {
		oldTeeIdRule = append(oldTeeIdRule, oldTeeIdItem)
	}
	var newTeeIdRule []interface{}
	for _, newTeeIdItem := range newTeeId {
		newTeeIdRule = append(newTeeIdRule, newTeeIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "TeeMachineReplicationTriggered", oldTeeIdRule, newTeeIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryTeeMachineReplicationTriggeredIterator{contract: _TeeRegistry.contract, event: "TeeMachineReplicationTriggered", logs: logs, sub: sub}, nil
}

// WatchTeeMachineReplicationTriggered is a free log subscription operation binding the contract event 0xec7149dab14e163102ba012e316ce7c830a771920fe7f043e34f45d10dd928ad.
//
// Solidity: event TeeMachineReplicationTriggered(address indexed oldTeeId, address indexed newTeeId, uint256 teeUpgradeId)
func (_TeeRegistry *TeeRegistryFilterer) WatchTeeMachineReplicationTriggered(opts *bind.WatchOpts, sink chan<- *TeeRegistryTeeMachineReplicationTriggered, oldTeeId []common.Address, newTeeId []common.Address) (event.Subscription, error) {

	var oldTeeIdRule []interface{}
	for _, oldTeeIdItem := range oldTeeId {
		oldTeeIdRule = append(oldTeeIdRule, oldTeeIdItem)
	}
	var newTeeIdRule []interface{}
	for _, newTeeIdItem := range newTeeId {
		newTeeIdRule = append(newTeeIdRule, newTeeIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "TeeMachineReplicationTriggered", oldTeeIdRule, newTeeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryTeeMachineReplicationTriggered)
				if err := _TeeRegistry.contract.UnpackLog(event, "TeeMachineReplicationTriggered", log); err != nil {
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

// ParseTeeMachineReplicationTriggered is a log parse operation binding the contract event 0xec7149dab14e163102ba012e316ce7c830a771920fe7f043e34f45d10dd928ad.
//
// Solidity: event TeeMachineReplicationTriggered(address indexed oldTeeId, address indexed newTeeId, uint256 teeUpgradeId)
func (_TeeRegistry *TeeRegistryFilterer) ParseTeeMachineReplicationTriggered(log types.Log) (*TeeRegistryTeeMachineReplicationTriggered, error) {
	event := new(TeeRegistryTeeMachineReplicationTriggered)
	if err := _TeeRegistry.contract.UnpackLog(event, "TeeMachineReplicationTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeRegistryTeeProxyIdSetIterator is returned from FilterTeeProxyIdSet and is used to iterate over the raw logs and unpacked data for TeeProxyIdSet events raised by the TeeRegistry contract.
type TeeRegistryTeeProxyIdSetIterator struct {
	Event *TeeRegistryTeeProxyIdSet // Event containing the contract specifics and raw log

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
func (it *TeeRegistryTeeProxyIdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeRegistryTeeProxyIdSet)
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
		it.Event = new(TeeRegistryTeeProxyIdSet)
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
func (it *TeeRegistryTeeProxyIdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeRegistryTeeProxyIdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeRegistryTeeProxyIdSet represents a TeeProxyIdSet event raised by the TeeRegistry contract.
type TeeRegistryTeeProxyIdSet struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTeeProxyIdSet is a free log retrieval operation binding the contract event 0x42f543e0aaa9ef7d6d77204eff6210ec3d344b013187285f4ba005d9ea568d81.
//
// Solidity: event TeeProxyIdSet(address indexed teeId, address indexed teeProxyId)
func (_TeeRegistry *TeeRegistryFilterer) FilterTeeProxyIdSet(opts *bind.FilterOpts, teeId []common.Address, teeProxyId []common.Address) (*TeeRegistryTeeProxyIdSetIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var teeProxyIdRule []interface{}
	for _, teeProxyIdItem := range teeProxyId {
		teeProxyIdRule = append(teeProxyIdRule, teeProxyIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.FilterLogs(opts, "TeeProxyIdSet", teeIdRule, teeProxyIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeRegistryTeeProxyIdSetIterator{contract: _TeeRegistry.contract, event: "TeeProxyIdSet", logs: logs, sub: sub}, nil
}

// WatchTeeProxyIdSet is a free log subscription operation binding the contract event 0x42f543e0aaa9ef7d6d77204eff6210ec3d344b013187285f4ba005d9ea568d81.
//
// Solidity: event TeeProxyIdSet(address indexed teeId, address indexed teeProxyId)
func (_TeeRegistry *TeeRegistryFilterer) WatchTeeProxyIdSet(opts *bind.WatchOpts, sink chan<- *TeeRegistryTeeProxyIdSet, teeId []common.Address, teeProxyId []common.Address) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var teeProxyIdRule []interface{}
	for _, teeProxyIdItem := range teeProxyId {
		teeProxyIdRule = append(teeProxyIdRule, teeProxyIdItem)
	}

	logs, sub, err := _TeeRegistry.contract.WatchLogs(opts, "TeeProxyIdSet", teeIdRule, teeProxyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeRegistryTeeProxyIdSet)
				if err := _TeeRegistry.contract.UnpackLog(event, "TeeProxyIdSet", log); err != nil {
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

// ParseTeeProxyIdSet is a log parse operation binding the contract event 0x42f543e0aaa9ef7d6d77204eff6210ec3d344b013187285f4ba005d9ea568d81.
//
// Solidity: event TeeProxyIdSet(address indexed teeId, address indexed teeProxyId)
func (_TeeRegistry *TeeRegistryFilterer) ParseTeeProxyIdSet(log types.Log) (*TeeRegistryTeeProxyIdSet, error) {
	event := new(TeeRegistryTeeProxyIdSet)
	if err := _TeeRegistry.contract.UnpackLog(event, "TeeProxyIdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
