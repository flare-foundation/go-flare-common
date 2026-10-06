// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package machinepathmanager

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

// IMachinePathManagerApproval is an auto generated low-level Go binding around an user-defined struct.
type IMachinePathManagerApproval struct {
	Signer      common.Address
	BlockNumber uint64
	SafeNonce   uint32
}

// IMachinePathManagerMachinePath is an auto generated low-level Go binding around an user-defined struct.
type IMachinePathManagerMachinePath struct {
	SourceTeeIds      []common.Address
	DestinationTeeIds []common.Address
}

// IMachinePathManagerSafeApprovalArtifact is an auto generated low-level Go binding around an user-defined struct.
type IMachinePathManagerSafeApprovalArtifact struct {
	SafeNonce  *big.Int
	Signatures []byte
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// MachinePathManagerMetaData contains all meta data concerning the MachinePathManager contract.
var MachinePathManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DestinationTeeIdAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"}],\"name\":\"GovernanceHashZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMachinePath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatureType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignaturesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListAlreadyFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ListNotFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoActiveMachinePathList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoDestinationTeeIds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPaths\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSourceTeeIds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeApprovalAlreadyConfirmed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeApprovalNotRecorded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeGovernanceStale\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignerAlreadySigned\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SourceTeeIdAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"counted\",\"type\":\"uint256\"}],\"name\":\"ThresholdNotReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnorderedSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnrecognizedSigner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"safeNonce\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"satisfiedGovernanceHashes\",\"type\":\"bytes32[]\"}],\"name\":\"MachinePathListApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"involvedGovernanceHashes\",\"type\":\"bytes32[]\"}],\"name\":\"MachinePathListFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"governanceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"safeNonce\",\"type\":\"uint256\"}],\"name\":\"MachinePathListSafeApprovalConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"countedGovernanceHashes\",\"type\":\"bytes32[]\"}],\"name\":\"MachinePathListSignatureAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"MachinePathListSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"MachinePathListStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"sourceTeeIds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"destinationTeeIds\",\"type\":\"address[]\"}],\"indexed\":false,\"internalType\":\"structIMachinePathManager.MachinePath[]\",\"name\":\"paths\",\"type\":\"tuple[]\"}],\"name\":\"MachinePathsAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"sourceTeeIds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"destinationTeeIds\",\"type\":\"address[]\"}],\"internalType\":\"structIMachinePathManager.MachinePath[]\",\"name\":\"_paths\",\"type\":\"tuple[]\"}],\"name\":\"addMachinePaths\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"approveMachinePathList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_safeNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signatures\",\"type\":\"bytes\"}],\"name\":\"confirmMachinePathListSafeApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"createNewMachinePathList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"finalizeMachinePathList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getActiveMachinePathListNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getMachinePathList\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"sourceTeeIds\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"destinationTeeIds\",\"type\":\"address[]\"}],\"internalType\":\"structIMachinePathManager.MachinePath[]\",\"name\":\"_paths\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_involvedGovernanceHashes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"_signatures\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"_signed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getMachinePathListApprovals\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"safeNonce\",\"type\":\"uint32\"}],\"internalType\":\"structIMachinePathManager.Approval[]\",\"name\":\"_approvals\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getMachinePathListMessageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"getMachinePathListSafeApprovalArtifact\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"safeNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"internalType\":\"structIMachinePathManager.SafeApprovalArtifact\",\"name\":\"_artifact\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"getMachinePathListSignatureCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getMachinePathListsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"isMachinePathListFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"isMachinePathListSafeApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"isMachinePathListSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sourceTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationTeeId\",\"type\":\"address\"}],\"name\":\"isMachinePathValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signMachinePathList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "MachinePathManager",
}

// MachinePathManager is an auto generated Go binding around an Ethereum contract.
type MachinePathManager struct {
	abi abi.ABI
}

// NewMachinePathManager creates a new instance of MachinePathManager.
func NewMachinePathManager() *MachinePathManager {
	parsed, err := MachinePathManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &MachinePathManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *MachinePathManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddMachinePaths is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac6dad90.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addMachinePaths(uint256 _extensionId, uint256 _nonce, (address[],address[])[] _paths) returns()
func (machinePathManager *MachinePathManager) PackAddMachinePaths(extensionId *big.Int, nonce *big.Int, paths []IMachinePathManagerMachinePath) []byte {
	enc, err := machinePathManager.abi.Pack("addMachinePaths", extensionId, nonce, paths)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddMachinePaths is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xac6dad90.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addMachinePaths(uint256 _extensionId, uint256 _nonce, (address[],address[])[] _paths) returns()
func (machinePathManager *MachinePathManager) TryPackAddMachinePaths(extensionId *big.Int, nonce *big.Int, paths []IMachinePathManagerMachinePath) ([]byte, error) {
	return machinePathManager.abi.Pack("addMachinePaths", extensionId, nonce, paths)
}

// PackApproveMachinePathList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1def5f18.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function approveMachinePathList(uint256 _extensionId, uint256 _nonce, bytes32 _messageHash) returns()
func (machinePathManager *MachinePathManager) PackApproveMachinePathList(extensionId *big.Int, nonce *big.Int, messageHash [32]byte) []byte {
	enc, err := machinePathManager.abi.Pack("approveMachinePathList", extensionId, nonce, messageHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackApproveMachinePathList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1def5f18.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function approveMachinePathList(uint256 _extensionId, uint256 _nonce, bytes32 _messageHash) returns()
func (machinePathManager *MachinePathManager) TryPackApproveMachinePathList(extensionId *big.Int, nonce *big.Int, messageHash [32]byte) ([]byte, error) {
	return machinePathManager.abi.Pack("approveMachinePathList", extensionId, nonce, messageHash)
}

// PackConfirmMachinePathListSafeApproval is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d993477.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmMachinePathListSafeApproval(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash, uint256 _safeNonce, bytes _signatures) returns()
func (machinePathManager *MachinePathManager) PackConfirmMachinePathListSafeApproval(extensionId *big.Int, nonce *big.Int, governanceHash [32]byte, safeNonce *big.Int, signatures []byte) []byte {
	enc, err := machinePathManager.abi.Pack("confirmMachinePathListSafeApproval", extensionId, nonce, governanceHash, safeNonce, signatures)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmMachinePathListSafeApproval is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d993477.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmMachinePathListSafeApproval(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash, uint256 _safeNonce, bytes _signatures) returns()
func (machinePathManager *MachinePathManager) TryPackConfirmMachinePathListSafeApproval(extensionId *big.Int, nonce *big.Int, governanceHash [32]byte, safeNonce *big.Int, signatures []byte) ([]byte, error) {
	return machinePathManager.abi.Pack("confirmMachinePathListSafeApproval", extensionId, nonce, governanceHash, safeNonce, signatures)
}

// PackCreateNewMachinePathList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07ae64b2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createNewMachinePathList(uint256 _extensionId) returns(uint256 _nonce)
func (machinePathManager *MachinePathManager) PackCreateNewMachinePathList(extensionId *big.Int) []byte {
	enc, err := machinePathManager.abi.Pack("createNewMachinePathList", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreateNewMachinePathList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07ae64b2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createNewMachinePathList(uint256 _extensionId) returns(uint256 _nonce)
func (machinePathManager *MachinePathManager) TryPackCreateNewMachinePathList(extensionId *big.Int) ([]byte, error) {
	return machinePathManager.abi.Pack("createNewMachinePathList", extensionId)
}

// UnpackCreateNewMachinePathList is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x07ae64b2.
//
// Solidity: function createNewMachinePathList(uint256 _extensionId) returns(uint256 _nonce)
func (machinePathManager *MachinePathManager) UnpackCreateNewMachinePathList(data []byte) (*big.Int, error) {
	out, err := machinePathManager.abi.Unpack("createNewMachinePathList", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackFinalizeMachinePathList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7fb88568.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function finalizeMachinePathList(uint256 _extensionId, uint256 _nonce) returns()
func (machinePathManager *MachinePathManager) PackFinalizeMachinePathList(extensionId *big.Int, nonce *big.Int) []byte {
	enc, err := machinePathManager.abi.Pack("finalizeMachinePathList", extensionId, nonce)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFinalizeMachinePathList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7fb88568.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function finalizeMachinePathList(uint256 _extensionId, uint256 _nonce) returns()
func (machinePathManager *MachinePathManager) TryPackFinalizeMachinePathList(extensionId *big.Int, nonce *big.Int) ([]byte, error) {
	return machinePathManager.abi.Pack("finalizeMachinePathList", extensionId, nonce)
}

// PackGetActiveMachinePathListNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb9b27a4b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getActiveMachinePathListNonce(uint256 _extensionId) view returns(uint256 _nonce)
func (machinePathManager *MachinePathManager) PackGetActiveMachinePathListNonce(extensionId *big.Int) []byte {
	enc, err := machinePathManager.abi.Pack("getActiveMachinePathListNonce", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetActiveMachinePathListNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb9b27a4b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getActiveMachinePathListNonce(uint256 _extensionId) view returns(uint256 _nonce)
func (machinePathManager *MachinePathManager) TryPackGetActiveMachinePathListNonce(extensionId *big.Int) ([]byte, error) {
	return machinePathManager.abi.Pack("getActiveMachinePathListNonce", extensionId)
}

// UnpackGetActiveMachinePathListNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb9b27a4b.
//
// Solidity: function getActiveMachinePathListNonce(uint256 _extensionId) view returns(uint256 _nonce)
func (machinePathManager *MachinePathManager) UnpackGetActiveMachinePathListNonce(data []byte) (*big.Int, error) {
	out, err := machinePathManager.abi.Unpack("getActiveMachinePathListNonce", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetMachinePathList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02473311.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMachinePathList(uint256 _extensionId, uint256 _nonce) view returns((address[],address[])[] _paths, bytes32[] _involvedGovernanceHashes, (uint8,bytes32,bytes32)[] _signatures, bool _signed)
func (machinePathManager *MachinePathManager) PackGetMachinePathList(extensionId *big.Int, nonce *big.Int) []byte {
	enc, err := machinePathManager.abi.Pack("getMachinePathList", extensionId, nonce)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMachinePathList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x02473311.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMachinePathList(uint256 _extensionId, uint256 _nonce) view returns((address[],address[])[] _paths, bytes32[] _involvedGovernanceHashes, (uint8,bytes32,bytes32)[] _signatures, bool _signed)
func (machinePathManager *MachinePathManager) TryPackGetMachinePathList(extensionId *big.Int, nonce *big.Int) ([]byte, error) {
	return machinePathManager.abi.Pack("getMachinePathList", extensionId, nonce)
}

// GetMachinePathListOutput serves as a container for the return parameters of contract
// method GetMachinePathList.
type GetMachinePathListOutput struct {
	Paths                    []IMachinePathManagerMachinePath
	InvolvedGovernanceHashes [][32]byte
	Signatures               []Signature
	Signed                   bool
}

// UnpackGetMachinePathList is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x02473311.
//
// Solidity: function getMachinePathList(uint256 _extensionId, uint256 _nonce) view returns((address[],address[])[] _paths, bytes32[] _involvedGovernanceHashes, (uint8,bytes32,bytes32)[] _signatures, bool _signed)
func (machinePathManager *MachinePathManager) UnpackGetMachinePathList(data []byte) (GetMachinePathListOutput, error) {
	out, err := machinePathManager.abi.Unpack("getMachinePathList", data)
	outstruct := new(GetMachinePathListOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Paths = *abi.ConvertType(out[0], new([]IMachinePathManagerMachinePath)).(*[]IMachinePathManagerMachinePath)
	outstruct.InvolvedGovernanceHashes = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)
	outstruct.Signatures = *abi.ConvertType(out[2], new([]Signature)).(*[]Signature)
	outstruct.Signed = *abi.ConvertType(out[3], new(bool)).(*bool)
	return *outstruct, nil
}

// PackGetMachinePathListApprovals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfd41976f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMachinePathListApprovals(uint256 _extensionId, uint256 _nonce) view returns((address,uint64,uint32)[] _approvals)
func (machinePathManager *MachinePathManager) PackGetMachinePathListApprovals(extensionId *big.Int, nonce *big.Int) []byte {
	enc, err := machinePathManager.abi.Pack("getMachinePathListApprovals", extensionId, nonce)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMachinePathListApprovals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfd41976f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMachinePathListApprovals(uint256 _extensionId, uint256 _nonce) view returns((address,uint64,uint32)[] _approvals)
func (machinePathManager *MachinePathManager) TryPackGetMachinePathListApprovals(extensionId *big.Int, nonce *big.Int) ([]byte, error) {
	return machinePathManager.abi.Pack("getMachinePathListApprovals", extensionId, nonce)
}

// UnpackGetMachinePathListApprovals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfd41976f.
//
// Solidity: function getMachinePathListApprovals(uint256 _extensionId, uint256 _nonce) view returns((address,uint64,uint32)[] _approvals)
func (machinePathManager *MachinePathManager) UnpackGetMachinePathListApprovals(data []byte) ([]IMachinePathManagerApproval, error) {
	out, err := machinePathManager.abi.Unpack("getMachinePathListApprovals", data)
	if err != nil {
		return *new([]IMachinePathManagerApproval), err
	}
	out0 := *abi.ConvertType(out[0], new([]IMachinePathManagerApproval)).(*[]IMachinePathManagerApproval)
	return out0, nil
}

// PackGetMachinePathListMessageHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb27f4702.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMachinePathListMessageHash(uint256 _extensionId, uint256 _nonce) view returns(bytes32)
func (machinePathManager *MachinePathManager) PackGetMachinePathListMessageHash(extensionId *big.Int, nonce *big.Int) []byte {
	enc, err := machinePathManager.abi.Pack("getMachinePathListMessageHash", extensionId, nonce)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMachinePathListMessageHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb27f4702.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMachinePathListMessageHash(uint256 _extensionId, uint256 _nonce) view returns(bytes32)
func (machinePathManager *MachinePathManager) TryPackGetMachinePathListMessageHash(extensionId *big.Int, nonce *big.Int) ([]byte, error) {
	return machinePathManager.abi.Pack("getMachinePathListMessageHash", extensionId, nonce)
}

// UnpackGetMachinePathListMessageHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb27f4702.
//
// Solidity: function getMachinePathListMessageHash(uint256 _extensionId, uint256 _nonce) view returns(bytes32)
func (machinePathManager *MachinePathManager) UnpackGetMachinePathListMessageHash(data []byte) ([32]byte, error) {
	out, err := machinePathManager.abi.Unpack("getMachinePathListMessageHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetMachinePathListSafeApprovalArtifact is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42b46389.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMachinePathListSafeApprovalArtifact(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash) view returns((uint256,bytes) _artifact)
func (machinePathManager *MachinePathManager) PackGetMachinePathListSafeApprovalArtifact(extensionId *big.Int, nonce *big.Int, governanceHash [32]byte) []byte {
	enc, err := machinePathManager.abi.Pack("getMachinePathListSafeApprovalArtifact", extensionId, nonce, governanceHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMachinePathListSafeApprovalArtifact is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42b46389.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMachinePathListSafeApprovalArtifact(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash) view returns((uint256,bytes) _artifact)
func (machinePathManager *MachinePathManager) TryPackGetMachinePathListSafeApprovalArtifact(extensionId *big.Int, nonce *big.Int, governanceHash [32]byte) ([]byte, error) {
	return machinePathManager.abi.Pack("getMachinePathListSafeApprovalArtifact", extensionId, nonce, governanceHash)
}

// UnpackGetMachinePathListSafeApprovalArtifact is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x42b46389.
//
// Solidity: function getMachinePathListSafeApprovalArtifact(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash) view returns((uint256,bytes) _artifact)
func (machinePathManager *MachinePathManager) UnpackGetMachinePathListSafeApprovalArtifact(data []byte) (IMachinePathManagerSafeApprovalArtifact, error) {
	out, err := machinePathManager.abi.Unpack("getMachinePathListSafeApprovalArtifact", data)
	if err != nil {
		return *new(IMachinePathManagerSafeApprovalArtifact), err
	}
	out0 := *abi.ConvertType(out[0], new(IMachinePathManagerSafeApprovalArtifact)).(*IMachinePathManagerSafeApprovalArtifact)
	return out0, nil
}

// PackGetMachinePathListSignatureCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4acbad8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMachinePathListSignatureCount(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash) view returns(uint64)
func (machinePathManager *MachinePathManager) PackGetMachinePathListSignatureCount(extensionId *big.Int, nonce *big.Int, governanceHash [32]byte) []byte {
	enc, err := machinePathManager.abi.Pack("getMachinePathListSignatureCount", extensionId, nonce, governanceHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMachinePathListSignatureCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4acbad8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMachinePathListSignatureCount(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash) view returns(uint64)
func (machinePathManager *MachinePathManager) TryPackGetMachinePathListSignatureCount(extensionId *big.Int, nonce *big.Int, governanceHash [32]byte) ([]byte, error) {
	return machinePathManager.abi.Pack("getMachinePathListSignatureCount", extensionId, nonce, governanceHash)
}

// UnpackGetMachinePathListSignatureCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa4acbad8.
//
// Solidity: function getMachinePathListSignatureCount(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash) view returns(uint64)
func (machinePathManager *MachinePathManager) UnpackGetMachinePathListSignatureCount(data []byte) (uint64, error) {
	out, err := machinePathManager.abi.Unpack("getMachinePathListSignatureCount", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetMachinePathListsCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa71b1c71.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getMachinePathListsCount(uint256 _extensionId) view returns(uint256)
func (machinePathManager *MachinePathManager) PackGetMachinePathListsCount(extensionId *big.Int) []byte {
	enc, err := machinePathManager.abi.Pack("getMachinePathListsCount", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetMachinePathListsCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa71b1c71.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getMachinePathListsCount(uint256 _extensionId) view returns(uint256)
func (machinePathManager *MachinePathManager) TryPackGetMachinePathListsCount(extensionId *big.Int) ([]byte, error) {
	return machinePathManager.abi.Pack("getMachinePathListsCount", extensionId)
}

// UnpackGetMachinePathListsCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa71b1c71.
//
// Solidity: function getMachinePathListsCount(uint256 _extensionId) view returns(uint256)
func (machinePathManager *MachinePathManager) UnpackGetMachinePathListsCount(data []byte) (*big.Int, error) {
	out, err := machinePathManager.abi.Unpack("getMachinePathListsCount", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackIsMachinePathListFinalized is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb367ffed.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isMachinePathListFinalized(uint256 _extensionId, uint256 _nonce) view returns(bool)
func (machinePathManager *MachinePathManager) PackIsMachinePathListFinalized(extensionId *big.Int, nonce *big.Int) []byte {
	enc, err := machinePathManager.abi.Pack("isMachinePathListFinalized", extensionId, nonce)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsMachinePathListFinalized is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb367ffed.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isMachinePathListFinalized(uint256 _extensionId, uint256 _nonce) view returns(bool)
func (machinePathManager *MachinePathManager) TryPackIsMachinePathListFinalized(extensionId *big.Int, nonce *big.Int) ([]byte, error) {
	return machinePathManager.abi.Pack("isMachinePathListFinalized", extensionId, nonce)
}

// UnpackIsMachinePathListFinalized is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb367ffed.
//
// Solidity: function isMachinePathListFinalized(uint256 _extensionId, uint256 _nonce) view returns(bool)
func (machinePathManager *MachinePathManager) UnpackIsMachinePathListFinalized(data []byte) (bool, error) {
	out, err := machinePathManager.abi.Unpack("isMachinePathListFinalized", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsMachinePathListSafeApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x489fec55.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isMachinePathListSafeApproved(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash) view returns(bool)
func (machinePathManager *MachinePathManager) PackIsMachinePathListSafeApproved(extensionId *big.Int, nonce *big.Int, governanceHash [32]byte) []byte {
	enc, err := machinePathManager.abi.Pack("isMachinePathListSafeApproved", extensionId, nonce, governanceHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsMachinePathListSafeApproved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x489fec55.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isMachinePathListSafeApproved(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash) view returns(bool)
func (machinePathManager *MachinePathManager) TryPackIsMachinePathListSafeApproved(extensionId *big.Int, nonce *big.Int, governanceHash [32]byte) ([]byte, error) {
	return machinePathManager.abi.Pack("isMachinePathListSafeApproved", extensionId, nonce, governanceHash)
}

// UnpackIsMachinePathListSafeApproved is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x489fec55.
//
// Solidity: function isMachinePathListSafeApproved(uint256 _extensionId, uint256 _nonce, bytes32 _governanceHash) view returns(bool)
func (machinePathManager *MachinePathManager) UnpackIsMachinePathListSafeApproved(data []byte) (bool, error) {
	out, err := machinePathManager.abi.Unpack("isMachinePathListSafeApproved", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsMachinePathListSigned is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70838db9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isMachinePathListSigned(uint256 _extensionId, uint256 _nonce) view returns(bool)
func (machinePathManager *MachinePathManager) PackIsMachinePathListSigned(extensionId *big.Int, nonce *big.Int) []byte {
	enc, err := machinePathManager.abi.Pack("isMachinePathListSigned", extensionId, nonce)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsMachinePathListSigned is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70838db9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isMachinePathListSigned(uint256 _extensionId, uint256 _nonce) view returns(bool)
func (machinePathManager *MachinePathManager) TryPackIsMachinePathListSigned(extensionId *big.Int, nonce *big.Int) ([]byte, error) {
	return machinePathManager.abi.Pack("isMachinePathListSigned", extensionId, nonce)
}

// UnpackIsMachinePathListSigned is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70838db9.
//
// Solidity: function isMachinePathListSigned(uint256 _extensionId, uint256 _nonce) view returns(bool)
func (machinePathManager *MachinePathManager) UnpackIsMachinePathListSigned(data []byte) (bool, error) {
	out, err := machinePathManager.abi.Unpack("isMachinePathListSigned", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsMachinePathValid is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x76ffee5c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isMachinePathValid(uint256 _extensionId, address _sourceTeeId, address _destinationTeeId) view returns(bool)
func (machinePathManager *MachinePathManager) PackIsMachinePathValid(extensionId *big.Int, sourceTeeId common.Address, destinationTeeId common.Address) []byte {
	enc, err := machinePathManager.abi.Pack("isMachinePathValid", extensionId, sourceTeeId, destinationTeeId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsMachinePathValid is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x76ffee5c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isMachinePathValid(uint256 _extensionId, address _sourceTeeId, address _destinationTeeId) view returns(bool)
func (machinePathManager *MachinePathManager) TryPackIsMachinePathValid(extensionId *big.Int, sourceTeeId common.Address, destinationTeeId common.Address) ([]byte, error) {
	return machinePathManager.abi.Pack("isMachinePathValid", extensionId, sourceTeeId, destinationTeeId)
}

// UnpackIsMachinePathValid is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x76ffee5c.
//
// Solidity: function isMachinePathValid(uint256 _extensionId, address _sourceTeeId, address _destinationTeeId) view returns(bool)
func (machinePathManager *MachinePathManager) UnpackIsMachinePathValid(data []byte) (bool, error) {
	out, err := machinePathManager.abi.Unpack("isMachinePathValid", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSignMachinePathList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcce3332d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signMachinePathList(uint256 _extensionId, uint256 _nonce, (uint8,bytes32,bytes32) _signature) returns()
func (machinePathManager *MachinePathManager) PackSignMachinePathList(extensionId *big.Int, nonce *big.Int, signature Signature) []byte {
	enc, err := machinePathManager.abi.Pack("signMachinePathList", extensionId, nonce, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSignMachinePathList is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcce3332d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signMachinePathList(uint256 _extensionId, uint256 _nonce, (uint8,bytes32,bytes32) _signature) returns()
func (machinePathManager *MachinePathManager) TryPackSignMachinePathList(extensionId *big.Int, nonce *big.Int, signature Signature) ([]byte, error) {
	return machinePathManager.abi.Pack("signMachinePathList", extensionId, nonce, signature)
}

// MachinePathManagerMachinePathListApproved represents a MachinePathListApproved event raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListApproved struct {
	ExtensionId               *big.Int
	Nonce                     *big.Int
	Safe                      common.Address
	SafeNonce                 uint32
	SatisfiedGovernanceHashes [][32]byte
	Raw                       *types.Log // Blockchain specific contextual infos
}

const MachinePathManagerMachinePathListApprovedEventName = "MachinePathListApproved"

// ContractEventName returns the user-defined event name.
func (MachinePathManagerMachinePathListApproved) ContractEventName() string {
	return MachinePathManagerMachinePathListApprovedEventName
}

// UnpackMachinePathListApprovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MachinePathListApproved(uint256 indexed extensionId, uint256 indexed nonce, address indexed safe, uint32 safeNonce, bytes32[] satisfiedGovernanceHashes)
func (machinePathManager *MachinePathManager) UnpackMachinePathListApprovedEvent(log *types.Log) (*MachinePathManagerMachinePathListApproved, error) {
	event := "MachinePathListApproved"
	if len(log.Topics) == 0 || log.Topics[0] != machinePathManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachinePathManagerMachinePathListApproved)
	if len(log.Data) > 0 {
		if err := machinePathManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machinePathManager.abi.Events[event].Inputs {
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

// MachinePathManagerMachinePathListFinalized represents a MachinePathListFinalized event raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListFinalized struct {
	ExtensionId              *big.Int
	Nonce                    *big.Int
	InvolvedGovernanceHashes [][32]byte
	Raw                      *types.Log // Blockchain specific contextual infos
}

const MachinePathManagerMachinePathListFinalizedEventName = "MachinePathListFinalized"

// ContractEventName returns the user-defined event name.
func (MachinePathManagerMachinePathListFinalized) ContractEventName() string {
	return MachinePathManagerMachinePathListFinalizedEventName
}

// UnpackMachinePathListFinalizedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MachinePathListFinalized(uint256 indexed extensionId, uint256 indexed nonce, bytes32[] involvedGovernanceHashes)
func (machinePathManager *MachinePathManager) UnpackMachinePathListFinalizedEvent(log *types.Log) (*MachinePathManagerMachinePathListFinalized, error) {
	event := "MachinePathListFinalized"
	if len(log.Topics) == 0 || log.Topics[0] != machinePathManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachinePathManagerMachinePathListFinalized)
	if len(log.Data) > 0 {
		if err := machinePathManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machinePathManager.abi.Events[event].Inputs {
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

// MachinePathManagerMachinePathListSafeApprovalConfirmed represents a MachinePathListSafeApprovalConfirmed event raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListSafeApprovalConfirmed struct {
	ExtensionId    *big.Int
	Nonce          *big.Int
	GovernanceHash [32]byte
	Safe           common.Address
	SafeNonce      *big.Int
	Raw            *types.Log // Blockchain specific contextual infos
}

const MachinePathManagerMachinePathListSafeApprovalConfirmedEventName = "MachinePathListSafeApprovalConfirmed"

// ContractEventName returns the user-defined event name.
func (MachinePathManagerMachinePathListSafeApprovalConfirmed) ContractEventName() string {
	return MachinePathManagerMachinePathListSafeApprovalConfirmedEventName
}

// UnpackMachinePathListSafeApprovalConfirmedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MachinePathListSafeApprovalConfirmed(uint256 indexed extensionId, uint256 indexed nonce, bytes32 indexed governanceHash, address safe, uint256 safeNonce)
func (machinePathManager *MachinePathManager) UnpackMachinePathListSafeApprovalConfirmedEvent(log *types.Log) (*MachinePathManagerMachinePathListSafeApprovalConfirmed, error) {
	event := "MachinePathListSafeApprovalConfirmed"
	if len(log.Topics) == 0 || log.Topics[0] != machinePathManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachinePathManagerMachinePathListSafeApprovalConfirmed)
	if len(log.Data) > 0 {
		if err := machinePathManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machinePathManager.abi.Events[event].Inputs {
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

// MachinePathManagerMachinePathListSignatureAdded represents a MachinePathListSignatureAdded event raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListSignatureAdded struct {
	ExtensionId             *big.Int
	Nonce                   *big.Int
	Signer                  common.Address
	CountedGovernanceHashes [][32]byte
	Raw                     *types.Log // Blockchain specific contextual infos
}

const MachinePathManagerMachinePathListSignatureAddedEventName = "MachinePathListSignatureAdded"

// ContractEventName returns the user-defined event name.
func (MachinePathManagerMachinePathListSignatureAdded) ContractEventName() string {
	return MachinePathManagerMachinePathListSignatureAddedEventName
}

// UnpackMachinePathListSignatureAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MachinePathListSignatureAdded(uint256 indexed extensionId, uint256 indexed nonce, address indexed signer, bytes32[] countedGovernanceHashes)
func (machinePathManager *MachinePathManager) UnpackMachinePathListSignatureAddedEvent(log *types.Log) (*MachinePathManagerMachinePathListSignatureAdded, error) {
	event := "MachinePathListSignatureAdded"
	if len(log.Topics) == 0 || log.Topics[0] != machinePathManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachinePathManagerMachinePathListSignatureAdded)
	if len(log.Data) > 0 {
		if err := machinePathManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machinePathManager.abi.Events[event].Inputs {
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

// MachinePathManagerMachinePathListSigned represents a MachinePathListSigned event raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListSigned struct {
	ExtensionId *big.Int
	Nonce       *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const MachinePathManagerMachinePathListSignedEventName = "MachinePathListSigned"

// ContractEventName returns the user-defined event name.
func (MachinePathManagerMachinePathListSigned) ContractEventName() string {
	return MachinePathManagerMachinePathListSignedEventName
}

// UnpackMachinePathListSignedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MachinePathListSigned(uint256 indexed extensionId, uint256 indexed nonce)
func (machinePathManager *MachinePathManager) UnpackMachinePathListSignedEvent(log *types.Log) (*MachinePathManagerMachinePathListSigned, error) {
	event := "MachinePathListSigned"
	if len(log.Topics) == 0 || log.Topics[0] != machinePathManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachinePathManagerMachinePathListSigned)
	if len(log.Data) > 0 {
		if err := machinePathManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machinePathManager.abi.Events[event].Inputs {
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

// MachinePathManagerMachinePathListStarted represents a MachinePathListStarted event raised by the MachinePathManager contract.
type MachinePathManagerMachinePathListStarted struct {
	ExtensionId *big.Int
	Nonce       *big.Int
	Raw         *types.Log // Blockchain specific contextual infos
}

const MachinePathManagerMachinePathListStartedEventName = "MachinePathListStarted"

// ContractEventName returns the user-defined event name.
func (MachinePathManagerMachinePathListStarted) ContractEventName() string {
	return MachinePathManagerMachinePathListStartedEventName
}

// UnpackMachinePathListStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MachinePathListStarted(uint256 indexed extensionId, uint256 indexed nonce)
func (machinePathManager *MachinePathManager) UnpackMachinePathListStartedEvent(log *types.Log) (*MachinePathManagerMachinePathListStarted, error) {
	event := "MachinePathListStarted"
	if len(log.Topics) == 0 || log.Topics[0] != machinePathManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachinePathManagerMachinePathListStarted)
	if len(log.Data) > 0 {
		if err := machinePathManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machinePathManager.abi.Events[event].Inputs {
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

// MachinePathManagerMachinePathsAdded represents a MachinePathsAdded event raised by the MachinePathManager contract.
type MachinePathManagerMachinePathsAdded struct {
	ExtensionId *big.Int
	Nonce       *big.Int
	Paths       []IMachinePathManagerMachinePath
	Raw         *types.Log // Blockchain specific contextual infos
}

const MachinePathManagerMachinePathsAddedEventName = "MachinePathsAdded"

// ContractEventName returns the user-defined event name.
func (MachinePathManagerMachinePathsAdded) ContractEventName() string {
	return MachinePathManagerMachinePathsAddedEventName
}

// UnpackMachinePathsAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event MachinePathsAdded(uint256 indexed extensionId, uint256 indexed nonce, (address[],address[])[] paths)
func (machinePathManager *MachinePathManager) UnpackMachinePathsAddedEvent(log *types.Log) (*MachinePathManagerMachinePathsAdded, error) {
	event := "MachinePathsAdded"
	if len(log.Topics) == 0 || log.Topics[0] != machinePathManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(MachinePathManagerMachinePathsAdded)
	if len(log.Data) > 0 {
		if err := machinePathManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range machinePathManager.abi.Events[event].Inputs {
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
func (machinePathManager *MachinePathManager) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["DestinationTeeIdAlreadyExists"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackDestinationTeeIdAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["GovernanceHashZero"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackGovernanceHashZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["InvalidMachinePath"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackInvalidMachinePathError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["InvalidSignatureType"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackInvalidSignatureTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["InvalidSignaturesLength"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackInvalidSignaturesLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["ListAlreadyFinalized"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackListAlreadyFinalizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["ListNotFinalized"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackListNotFinalizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["MessageHashMismatch"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackMessageHashMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["NoActiveMachinePathList"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackNoActiveMachinePathListError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["NoDestinationTeeIds"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackNoDestinationTeeIdsError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["NoPaths"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackNoPathsError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["NoSourceTeeIds"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackNoSourceTeeIdsError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["SafeApprovalAlreadyConfirmed"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackSafeApprovalAlreadyConfirmedError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["SafeApprovalNotRecorded"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackSafeApprovalNotRecordedError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["SafeGovernanceStale"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackSafeGovernanceStaleError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["SignerAlreadySigned"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackSignerAlreadySignedError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["SourceTeeIdAlreadyExists"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackSourceTeeIdAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["TeeNotFound"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackTeeNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["ThresholdNotReached"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackThresholdNotReachedError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["UnorderedSignatures"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackUnorderedSignaturesError(raw[4:])
	}
	if bytes.Equal(raw[:4], machinePathManager.abi.Errors["UnrecognizedSigner"].ID.Bytes()[:4]) {
		return machinePathManager.UnpackUnrecognizedSignerError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// MachinePathManagerDestinationTeeIdAlreadyExists represents a DestinationTeeIdAlreadyExists error raised by the MachinePathManager contract.
type MachinePathManagerDestinationTeeIdAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DestinationTeeIdAlreadyExists()
func MachinePathManagerDestinationTeeIdAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0xd648095f08c7920e227c3a076812cfc18651fe4461018eea9cdae13eb8186c44")
}

// UnpackDestinationTeeIdAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DestinationTeeIdAlreadyExists()
func (machinePathManager *MachinePathManager) UnpackDestinationTeeIdAlreadyExistsError(raw []byte) (*MachinePathManagerDestinationTeeIdAlreadyExists, error) {
	out := new(MachinePathManagerDestinationTeeIdAlreadyExists)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "DestinationTeeIdAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the MachinePathManager contract.
type MachinePathManagerECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func MachinePathManagerECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (machinePathManager *MachinePathManager) UnpackECDSAInvalidSignatureError(raw []byte) (*MachinePathManagerECDSAInvalidSignature, error) {
	out := new(MachinePathManagerECDSAInvalidSignature)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the MachinePathManager contract.
type MachinePathManagerECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func MachinePathManagerECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (machinePathManager *MachinePathManager) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*MachinePathManagerECDSAInvalidSignatureLength, error) {
	out := new(MachinePathManagerECDSAInvalidSignatureLength)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the MachinePathManager contract.
type MachinePathManagerECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func MachinePathManagerECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (machinePathManager *MachinePathManager) UnpackECDSAInvalidSignatureSError(raw []byte) (*MachinePathManagerECDSAInvalidSignatureS, error) {
	out := new(MachinePathManagerECDSAInvalidSignatureS)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerExtensionIdMismatch represents a ExtensionIdMismatch error raised by the MachinePathManager contract.
type MachinePathManagerExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func MachinePathManagerExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (machinePathManager *MachinePathManager) UnpackExtensionIdMismatchError(raw []byte) (*MachinePathManagerExtensionIdMismatch, error) {
	out := new(MachinePathManagerExtensionIdMismatch)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerGovernanceHashZero represents a GovernanceHashZero error raised by the MachinePathManager contract.
type MachinePathManagerGovernanceHashZero struct {
	TeeId common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernanceHashZero(address teeId)
func MachinePathManagerGovernanceHashZeroErrorID() common.Hash {
	return common.HexToHash("0x9309bca4c98210603f9027adaa91d9a8df4b825acf17b17fb1525a891ac09437")
}

// UnpackGovernanceHashZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernanceHashZero(address teeId)
func (machinePathManager *MachinePathManager) UnpackGovernanceHashZeroError(raw []byte) (*MachinePathManagerGovernanceHashZero, error) {
	out := new(MachinePathManagerGovernanceHashZero)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "GovernanceHashZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the MachinePathManager contract.
type MachinePathManagerInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func MachinePathManagerInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (machinePathManager *MachinePathManager) UnpackInvalidGovernanceHashError(raw []byte) (*MachinePathManagerInvalidGovernanceHash, error) {
	out := new(MachinePathManagerInvalidGovernanceHash)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerInvalidMachinePath represents a InvalidMachinePath error raised by the MachinePathManager contract.
type MachinePathManagerInvalidMachinePath struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidMachinePath()
func MachinePathManagerInvalidMachinePathErrorID() common.Hash {
	return common.HexToHash("0xf8a0ba22852d6de75ca0c6f463d7f4f70e97b45e489a445014ec85848bf9713b")
}

// UnpackInvalidMachinePathError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidMachinePath()
func (machinePathManager *MachinePathManager) UnpackInvalidMachinePathError(raw []byte) (*MachinePathManagerInvalidMachinePath, error) {
	out := new(MachinePathManagerInvalidMachinePath)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "InvalidMachinePath", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerInvalidNonce represents a InvalidNonce error raised by the MachinePathManager contract.
type MachinePathManagerInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func MachinePathManagerInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (machinePathManager *MachinePathManager) UnpackInvalidNonceError(raw []byte) (*MachinePathManagerInvalidNonce, error) {
	out := new(MachinePathManagerInvalidNonce)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerInvalidSignatureType represents a InvalidSignatureType error raised by the MachinePathManager contract.
type MachinePathManagerInvalidSignatureType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSignatureType()
func MachinePathManagerInvalidSignatureTypeErrorID() common.Hash {
	return common.HexToHash("0x60cd402d1c2cdf3cb1f7f894473c09df7f5c23063b38f8628ecf7b6e2f56df88")
}

// UnpackInvalidSignatureTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSignatureType()
func (machinePathManager *MachinePathManager) UnpackInvalidSignatureTypeError(raw []byte) (*MachinePathManagerInvalidSignatureType, error) {
	out := new(MachinePathManagerInvalidSignatureType)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "InvalidSignatureType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerInvalidSignaturesLength represents a InvalidSignaturesLength error raised by the MachinePathManager contract.
type MachinePathManagerInvalidSignaturesLength struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSignaturesLength()
func MachinePathManagerInvalidSignaturesLengthErrorID() common.Hash {
	return common.HexToHash("0xfb6c73045e15dde90f7a6787480292632f59a7ea6c90880eeafcb19635c006cf")
}

// UnpackInvalidSignaturesLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSignaturesLength()
func (machinePathManager *MachinePathManager) UnpackInvalidSignaturesLengthError(raw []byte) (*MachinePathManagerInvalidSignaturesLength, error) {
	out := new(MachinePathManagerInvalidSignaturesLength)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "InvalidSignaturesLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerListAlreadyFinalized represents a ListAlreadyFinalized error raised by the MachinePathManager contract.
type MachinePathManagerListAlreadyFinalized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ListAlreadyFinalized()
func MachinePathManagerListAlreadyFinalizedErrorID() common.Hash {
	return common.HexToHash("0x79197006f154f3c64ab365667c7545bdf3496a18e3b077138c6483ca27298a2f")
}

// UnpackListAlreadyFinalizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ListAlreadyFinalized()
func (machinePathManager *MachinePathManager) UnpackListAlreadyFinalizedError(raw []byte) (*MachinePathManagerListAlreadyFinalized, error) {
	out := new(MachinePathManagerListAlreadyFinalized)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "ListAlreadyFinalized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerListNotFinalized represents a ListNotFinalized error raised by the MachinePathManager contract.
type MachinePathManagerListNotFinalized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ListNotFinalized()
func MachinePathManagerListNotFinalizedErrorID() common.Hash {
	return common.HexToHash("0x06d05554359806b626ef25af9b70d5a3df2096f759ce1aa438cdaa202a251d35")
}

// UnpackListNotFinalizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ListNotFinalized()
func (machinePathManager *MachinePathManager) UnpackListNotFinalizedError(raw []byte) (*MachinePathManagerListNotFinalized, error) {
	out := new(MachinePathManagerListNotFinalized)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "ListNotFinalized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerMessageHashMismatch represents a MessageHashMismatch error raised by the MachinePathManager contract.
type MachinePathManagerMessageHashMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MessageHashMismatch()
func MachinePathManagerMessageHashMismatchErrorID() common.Hash {
	return common.HexToHash("0x93961fd56738def850ad116f0869a20eef493a29e4f12e42d503b18bc44d31a4")
}

// UnpackMessageHashMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MessageHashMismatch()
func (machinePathManager *MachinePathManager) UnpackMessageHashMismatchError(raw []byte) (*MachinePathManagerMessageHashMismatch, error) {
	out := new(MachinePathManagerMessageHashMismatch)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "MessageHashMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerNoActiveMachinePathList represents a NoActiveMachinePathList error raised by the MachinePathManager contract.
type MachinePathManagerNoActiveMachinePathList struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoActiveMachinePathList()
func MachinePathManagerNoActiveMachinePathListErrorID() common.Hash {
	return common.HexToHash("0xb0dae5768e38e4e9d9ec3149b08b494bbe09867f8733737d3ea98e00826652f2")
}

// UnpackNoActiveMachinePathListError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoActiveMachinePathList()
func (machinePathManager *MachinePathManager) UnpackNoActiveMachinePathListError(raw []byte) (*MachinePathManagerNoActiveMachinePathList, error) {
	out := new(MachinePathManagerNoActiveMachinePathList)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "NoActiveMachinePathList", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerNoDestinationTeeIds represents a NoDestinationTeeIds error raised by the MachinePathManager contract.
type MachinePathManagerNoDestinationTeeIds struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoDestinationTeeIds()
func MachinePathManagerNoDestinationTeeIdsErrorID() common.Hash {
	return common.HexToHash("0xcc22648432050d22cc42eafc327c736a29926bb913a662c76fd4f5c83e97061c")
}

// UnpackNoDestinationTeeIdsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoDestinationTeeIds()
func (machinePathManager *MachinePathManager) UnpackNoDestinationTeeIdsError(raw []byte) (*MachinePathManagerNoDestinationTeeIds, error) {
	out := new(MachinePathManagerNoDestinationTeeIds)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "NoDestinationTeeIds", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerNoPaths represents a NoPaths error raised by the MachinePathManager contract.
type MachinePathManagerNoPaths struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoPaths()
func MachinePathManagerNoPathsErrorID() common.Hash {
	return common.HexToHash("0xb209c5871a7461ae96adccf6561662b00064829073cdee63dd7cdc47a4811de5")
}

// UnpackNoPathsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoPaths()
func (machinePathManager *MachinePathManager) UnpackNoPathsError(raw []byte) (*MachinePathManagerNoPaths, error) {
	out := new(MachinePathManagerNoPaths)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "NoPaths", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerNoSourceTeeIds represents a NoSourceTeeIds error raised by the MachinePathManager contract.
type MachinePathManagerNoSourceTeeIds struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoSourceTeeIds()
func MachinePathManagerNoSourceTeeIdsErrorID() common.Hash {
	return common.HexToHash("0xa26eb84f3cfcacb4e47c87606c1ed2f5edfa9d05cc465bd2f8ea0e972b96f0af")
}

// UnpackNoSourceTeeIdsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoSourceTeeIds()
func (machinePathManager *MachinePathManager) UnpackNoSourceTeeIdsError(raw []byte) (*MachinePathManagerNoSourceTeeIds, error) {
	out := new(MachinePathManagerNoSourceTeeIds)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "NoSourceTeeIds", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the MachinePathManager contract.
type MachinePathManagerOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func MachinePathManagerOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (machinePathManager *MachinePathManager) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*MachinePathManagerOnlyExtensionOwnerOrOperator, error) {
	out := new(MachinePathManagerOnlyExtensionOwnerOrOperator)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerSafeApprovalAlreadyConfirmed represents a SafeApprovalAlreadyConfirmed error raised by the MachinePathManager contract.
type MachinePathManagerSafeApprovalAlreadyConfirmed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeApprovalAlreadyConfirmed()
func MachinePathManagerSafeApprovalAlreadyConfirmedErrorID() common.Hash {
	return common.HexToHash("0x1e5d73a6295dd6aad7f8b9101fec00f3cf3993dcbe13880ff1f81a7153ee5cf2")
}

// UnpackSafeApprovalAlreadyConfirmedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeApprovalAlreadyConfirmed()
func (machinePathManager *MachinePathManager) UnpackSafeApprovalAlreadyConfirmedError(raw []byte) (*MachinePathManagerSafeApprovalAlreadyConfirmed, error) {
	out := new(MachinePathManagerSafeApprovalAlreadyConfirmed)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "SafeApprovalAlreadyConfirmed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerSafeApprovalNotRecorded represents a SafeApprovalNotRecorded error raised by the MachinePathManager contract.
type MachinePathManagerSafeApprovalNotRecorded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeApprovalNotRecorded()
func MachinePathManagerSafeApprovalNotRecordedErrorID() common.Hash {
	return common.HexToHash("0x3358274e0faf47a43bb567b5b234120a835638e802c059ad95e390965b297471")
}

// UnpackSafeApprovalNotRecordedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeApprovalNotRecorded()
func (machinePathManager *MachinePathManager) UnpackSafeApprovalNotRecordedError(raw []byte) (*MachinePathManagerSafeApprovalNotRecorded, error) {
	out := new(MachinePathManagerSafeApprovalNotRecorded)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "SafeApprovalNotRecorded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerSafeGovernanceStale represents a SafeGovernanceStale error raised by the MachinePathManager contract.
type MachinePathManagerSafeGovernanceStale struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeGovernanceStale()
func MachinePathManagerSafeGovernanceStaleErrorID() common.Hash {
	return common.HexToHash("0xa0e8878d8ad235e5c5519e1e9d836eb20ceae72d64fc9a2f03683d7b18b1d7bc")
}

// UnpackSafeGovernanceStaleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeGovernanceStale()
func (machinePathManager *MachinePathManager) UnpackSafeGovernanceStaleError(raw []byte) (*MachinePathManagerSafeGovernanceStale, error) {
	out := new(MachinePathManagerSafeGovernanceStale)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "SafeGovernanceStale", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerSignerAlreadySigned represents a SignerAlreadySigned error raised by the MachinePathManager contract.
type MachinePathManagerSignerAlreadySigned struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SignerAlreadySigned()
func MachinePathManagerSignerAlreadySignedErrorID() common.Hash {
	return common.HexToHash("0xe05524d5a4ec4e0da60827b8018799f1bb404e897caf9d618c6a09ab3a7ebb00")
}

// UnpackSignerAlreadySignedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SignerAlreadySigned()
func (machinePathManager *MachinePathManager) UnpackSignerAlreadySignedError(raw []byte) (*MachinePathManagerSignerAlreadySigned, error) {
	out := new(MachinePathManagerSignerAlreadySigned)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "SignerAlreadySigned", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerSourceTeeIdAlreadyExists represents a SourceTeeIdAlreadyExists error raised by the MachinePathManager contract.
type MachinePathManagerSourceTeeIdAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceTeeIdAlreadyExists()
func MachinePathManagerSourceTeeIdAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x0c17d5aa5a16f9da3ba60606dfcd7d02bc655d80265ac88b59c41f5fe1c4f6c0")
}

// UnpackSourceTeeIdAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceTeeIdAlreadyExists()
func (machinePathManager *MachinePathManager) UnpackSourceTeeIdAlreadyExistsError(raw []byte) (*MachinePathManagerSourceTeeIdAlreadyExists, error) {
	out := new(MachinePathManagerSourceTeeIdAlreadyExists)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "SourceTeeIdAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerTeeNotFound represents a TeeNotFound error raised by the MachinePathManager contract.
type MachinePathManagerTeeNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeNotFound()
func MachinePathManagerTeeNotFoundErrorID() common.Hash {
	return common.HexToHash("0xceb05b6852ed4c7dc3a718eb0d26555398a76877a44dd41dc92c5fe7e6ef8c2d")
}

// UnpackTeeNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeNotFound()
func (machinePathManager *MachinePathManager) UnpackTeeNotFoundError(raw []byte) (*MachinePathManagerTeeNotFound, error) {
	out := new(MachinePathManagerTeeNotFound)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "TeeNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerThresholdNotReached represents a ThresholdNotReached error raised by the MachinePathManager contract.
type MachinePathManagerThresholdNotReached struct {
	Required *big.Int
	Counted  *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ThresholdNotReached(uint256 required, uint256 counted)
func MachinePathManagerThresholdNotReachedErrorID() common.Hash {
	return common.HexToHash("0xccf7c5f5fe9a2e53ec10b9e03f1280b37c3d051c388cc7e173e58f9641e9b844")
}

// UnpackThresholdNotReachedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ThresholdNotReached(uint256 required, uint256 counted)
func (machinePathManager *MachinePathManager) UnpackThresholdNotReachedError(raw []byte) (*MachinePathManagerThresholdNotReached, error) {
	out := new(MachinePathManagerThresholdNotReached)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "ThresholdNotReached", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerUnorderedSignatures represents a UnorderedSignatures error raised by the MachinePathManager contract.
type MachinePathManagerUnorderedSignatures struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnorderedSignatures()
func MachinePathManagerUnorderedSignaturesErrorID() common.Hash {
	return common.HexToHash("0x37fd368afc5914653c11331ed835e1e3ea5b6a5b5c3c1d3692d8fea563194fb1")
}

// UnpackUnorderedSignaturesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnorderedSignatures()
func (machinePathManager *MachinePathManager) UnpackUnorderedSignaturesError(raw []byte) (*MachinePathManagerUnorderedSignatures, error) {
	out := new(MachinePathManagerUnorderedSignatures)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "UnorderedSignatures", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// MachinePathManagerUnrecognizedSigner represents a UnrecognizedSigner error raised by the MachinePathManager contract.
type MachinePathManagerUnrecognizedSigner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnrecognizedSigner()
func MachinePathManagerUnrecognizedSignerErrorID() common.Hash {
	return common.HexToHash("0x0376f89c0b4cb304a274936daba3f2f7fb372d688998bd9714ac06eb3ddfe806")
}

// UnpackUnrecognizedSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnrecognizedSigner()
func (machinePathManager *MachinePathManager) UnpackUnrecognizedSignerError(raw []byte) (*MachinePathManagerUnrecognizedSigner, error) {
	out := new(MachinePathManagerUnrecognizedSigner)
	if err := machinePathManager.abi.UnpackIntoInterface(out, "UnrecognizedSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}
