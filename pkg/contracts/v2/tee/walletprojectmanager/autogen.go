// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletprojectmanager

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

// WalletProjectManagerMetaData contains all meta data concerning the WalletProjectManager contract.
var WalletProjectManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SigningAlgoNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotPartOfProject\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotProductionReady\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"backupManager\",\"type\":\"address\"}],\"name\":\"BackupManagerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"}],\"name\":\"ProjectCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_signingAlgo\",\"type\":\"bytes32\"}],\"name\":\"createProject\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getBackupManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_backupManager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getExtensionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getKeyType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getSigningAlgo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_signingAlgo\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_backupManager\",\"type\":\"address\"}],\"name\":\"setBackupManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "WalletProjectManager",
}

// WalletProjectManager is an auto generated Go binding around an Ethereum contract.
type WalletProjectManager struct {
	abi abi.ABI
}

// NewWalletProjectManager creates a new instance of WalletProjectManager.
func NewWalletProjectManager() *WalletProjectManager {
	parsed, err := WalletProjectManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &WalletProjectManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *WalletProjectManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConfirmOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ff0e872.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmOwnership(bytes32 _projectId) returns()
func (walletProjectManager *WalletProjectManager) PackConfirmOwnership(projectId [32]byte) []byte {
	enc, err := walletProjectManager.abi.Pack("confirmOwnership", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7ff0e872.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmOwnership(bytes32 _projectId) returns()
func (walletProjectManager *WalletProjectManager) TryPackConfirmOwnership(projectId [32]byte) ([]byte, error) {
	return walletProjectManager.abi.Pack("confirmOwnership", projectId)
}

// PackCreateProject is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa650c968.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createProject(uint256 _extensionId, bytes32 _keyType, bytes32 _signingAlgo) returns(bytes32 _projectId)
func (walletProjectManager *WalletProjectManager) PackCreateProject(extensionId *big.Int, keyType [32]byte, signingAlgo [32]byte) []byte {
	enc, err := walletProjectManager.abi.Pack("createProject", extensionId, keyType, signingAlgo)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreateProject is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa650c968.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createProject(uint256 _extensionId, bytes32 _keyType, bytes32 _signingAlgo) returns(bytes32 _projectId)
func (walletProjectManager *WalletProjectManager) TryPackCreateProject(extensionId *big.Int, keyType [32]byte, signingAlgo [32]byte) ([]byte, error) {
	return walletProjectManager.abi.Pack("createProject", extensionId, keyType, signingAlgo)
}

// UnpackCreateProject is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa650c968.
//
// Solidity: function createProject(uint256 _extensionId, bytes32 _keyType, bytes32 _signingAlgo) returns(bytes32 _projectId)
func (walletProjectManager *WalletProjectManager) UnpackCreateProject(data []byte) ([32]byte, error) {
	out, err := walletProjectManager.abi.Unpack("createProject", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetBackupManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7bb77989.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBackupManager(bytes32 _projectId) view returns(address _backupManager)
func (walletProjectManager *WalletProjectManager) PackGetBackupManager(projectId [32]byte) []byte {
	enc, err := walletProjectManager.abi.Pack("getBackupManager", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBackupManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7bb77989.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBackupManager(bytes32 _projectId) view returns(address _backupManager)
func (walletProjectManager *WalletProjectManager) TryPackGetBackupManager(projectId [32]byte) ([]byte, error) {
	return walletProjectManager.abi.Pack("getBackupManager", projectId)
}

// UnpackGetBackupManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7bb77989.
//
// Solidity: function getBackupManager(bytes32 _projectId) view returns(address _backupManager)
func (walletProjectManager *WalletProjectManager) UnpackGetBackupManager(data []byte) (common.Address, error) {
	out, err := walletProjectManager.abi.Unpack("getBackupManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetExtensionId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13def0e3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getExtensionId(bytes32 _projectId) view returns(uint256 _extensionId)
func (walletProjectManager *WalletProjectManager) PackGetExtensionId(projectId [32]byte) []byte {
	enc, err := walletProjectManager.abi.Pack("getExtensionId", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetExtensionId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x13def0e3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getExtensionId(bytes32 _projectId) view returns(uint256 _extensionId)
func (walletProjectManager *WalletProjectManager) TryPackGetExtensionId(projectId [32]byte) ([]byte, error) {
	return walletProjectManager.abi.Pack("getExtensionId", projectId)
}

// UnpackGetExtensionId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x13def0e3.
//
// Solidity: function getExtensionId(bytes32 _projectId) view returns(uint256 _extensionId)
func (walletProjectManager *WalletProjectManager) UnpackGetExtensionId(data []byte) (*big.Int, error) {
	out, err := walletProjectManager.abi.Unpack("getExtensionId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetKeyType is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x779d385a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getKeyType(bytes32 _projectId) view returns(bytes32 _keyType)
func (walletProjectManager *WalletProjectManager) PackGetKeyType(projectId [32]byte) []byte {
	enc, err := walletProjectManager.abi.Pack("getKeyType", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetKeyType is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x779d385a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getKeyType(bytes32 _projectId) view returns(bytes32 _keyType)
func (walletProjectManager *WalletProjectManager) TryPackGetKeyType(projectId [32]byte) ([]byte, error) {
	return walletProjectManager.abi.Pack("getKeyType", projectId)
}

// UnpackGetKeyType is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x779d385a.
//
// Solidity: function getKeyType(bytes32 _projectId) view returns(bytes32 _keyType)
func (walletProjectManager *WalletProjectManager) UnpackGetKeyType(data []byte) ([32]byte, error) {
	out, err := walletProjectManager.abi.Unpack("getKeyType", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdeb931a2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getOwner(bytes32 _projectId) view returns(address _owner)
func (walletProjectManager *WalletProjectManager) PackGetOwner(projectId [32]byte) []byte {
	enc, err := walletProjectManager.abi.Pack("getOwner", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdeb931a2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getOwner(bytes32 _projectId) view returns(address _owner)
func (walletProjectManager *WalletProjectManager) TryPackGetOwner(projectId [32]byte) ([]byte, error) {
	return walletProjectManager.abi.Pack("getOwner", projectId)
}

// UnpackGetOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdeb931a2.
//
// Solidity: function getOwner(bytes32 _projectId) view returns(address _owner)
func (walletProjectManager *WalletProjectManager) UnpackGetOwner(data []byte) (common.Address, error) {
	out, err := walletProjectManager.abi.Unpack("getOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetSigningAlgo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x17566896.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSigningAlgo(bytes32 _projectId) view returns(bytes32 _signingAlgo)
func (walletProjectManager *WalletProjectManager) PackGetSigningAlgo(projectId [32]byte) []byte {
	enc, err := walletProjectManager.abi.Pack("getSigningAlgo", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSigningAlgo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x17566896.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSigningAlgo(bytes32 _projectId) view returns(bytes32 _signingAlgo)
func (walletProjectManager *WalletProjectManager) TryPackGetSigningAlgo(projectId [32]byte) ([]byte, error) {
	return walletProjectManager.abi.Pack("getSigningAlgo", projectId)
}

// UnpackGetSigningAlgo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x17566896.
//
// Solidity: function getSigningAlgo(bytes32 _projectId) view returns(bytes32 _signingAlgo)
func (walletProjectManager *WalletProjectManager) UnpackGetSigningAlgo(data []byte) ([32]byte, error) {
	out, err := walletProjectManager.abi.Unpack("getSigningAlgo", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackProposeNewOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3ad60dd6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposeNewOwner(bytes32 _projectId, address _newOwner) returns()
func (walletProjectManager *WalletProjectManager) PackProposeNewOwner(projectId [32]byte, newOwner common.Address) []byte {
	enc, err := walletProjectManager.abi.Pack("proposeNewOwner", projectId, newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposeNewOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3ad60dd6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposeNewOwner(bytes32 _projectId, address _newOwner) returns()
func (walletProjectManager *WalletProjectManager) TryPackProposeNewOwner(projectId [32]byte, newOwner common.Address) ([]byte, error) {
	return walletProjectManager.abi.Pack("proposeNewOwner", projectId, newOwner)
}

// PackSetBackupManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x755fd7c2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setBackupManager(bytes32 _projectId, address _backupManager) returns()
func (walletProjectManager *WalletProjectManager) PackSetBackupManager(projectId [32]byte, backupManager common.Address) []byte {
	enc, err := walletProjectManager.abi.Pack("setBackupManager", projectId, backupManager)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetBackupManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x755fd7c2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setBackupManager(bytes32 _projectId, address _backupManager) returns()
func (walletProjectManager *WalletProjectManager) TryPackSetBackupManager(projectId [32]byte, backupManager common.Address) ([]byte, error) {
	return walletProjectManager.abi.Pack("setBackupManager", projectId, backupManager)
}

// WalletProjectManagerBackupManagerSet represents a BackupManagerSet event raised by the WalletProjectManager contract.
type WalletProjectManagerBackupManagerSet struct {
	ProjectId     [32]byte
	BackupManager common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const WalletProjectManagerBackupManagerSetEventName = "BackupManagerSet"

// ContractEventName returns the user-defined event name.
func (WalletProjectManagerBackupManagerSet) ContractEventName() string {
	return WalletProjectManagerBackupManagerSetEventName
}

// UnpackBackupManagerSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BackupManagerSet(bytes32 indexed projectId, address indexed backupManager)
func (walletProjectManager *WalletProjectManager) UnpackBackupManagerSetEvent(log *types.Log) (*WalletProjectManagerBackupManagerSet, error) {
	event := "BackupManagerSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletProjectManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletProjectManagerBackupManagerSet)
	if len(log.Data) > 0 {
		if err := walletProjectManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletProjectManager.abi.Events[event].Inputs {
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

// WalletProjectManagerNewOwnerProposed represents a NewOwnerProposed event raised by the WalletProjectManager contract.
type WalletProjectManagerNewOwnerProposed struct {
	ProjectId [32]byte
	NewOwner  common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletProjectManagerNewOwnerProposedEventName = "NewOwnerProposed"

// ContractEventName returns the user-defined event name.
func (WalletProjectManagerNewOwnerProposed) ContractEventName() string {
	return WalletProjectManagerNewOwnerProposedEventName
}

// UnpackNewOwnerProposedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewOwnerProposed(bytes32 indexed projectId, address indexed newOwner)
func (walletProjectManager *WalletProjectManager) UnpackNewOwnerProposedEvent(log *types.Log) (*WalletProjectManagerNewOwnerProposed, error) {
	event := "NewOwnerProposed"
	if len(log.Topics) == 0 || log.Topics[0] != walletProjectManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletProjectManagerNewOwnerProposed)
	if len(log.Data) > 0 {
		if err := walletProjectManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletProjectManager.abi.Events[event].Inputs {
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

// WalletProjectManagerOwnershipConfirmed represents a OwnershipConfirmed event raised by the WalletProjectManager contract.
type WalletProjectManagerOwnershipConfirmed struct {
	ProjectId [32]byte
	NewOwner  common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletProjectManagerOwnershipConfirmedEventName = "OwnershipConfirmed"

// ContractEventName returns the user-defined event name.
func (WalletProjectManagerOwnershipConfirmed) ContractEventName() string {
	return WalletProjectManagerOwnershipConfirmedEventName
}

// UnpackOwnershipConfirmedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipConfirmed(bytes32 indexed projectId, address indexed newOwner)
func (walletProjectManager *WalletProjectManager) UnpackOwnershipConfirmedEvent(log *types.Log) (*WalletProjectManagerOwnershipConfirmed, error) {
	event := "OwnershipConfirmed"
	if len(log.Topics) == 0 || log.Topics[0] != walletProjectManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletProjectManagerOwnershipConfirmed)
	if len(log.Data) > 0 {
		if err := walletProjectManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletProjectManager.abi.Events[event].Inputs {
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

// WalletProjectManagerProjectCreated represents a ProjectCreated event raised by the WalletProjectManager contract.
type WalletProjectManagerProjectCreated struct {
	ProjectId   [32]byte
	Owner       common.Address
	ExtensionId *big.Int
	KeyType     [32]byte
	SigningAlgo [32]byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const WalletProjectManagerProjectCreatedEventName = "ProjectCreated"

// ContractEventName returns the user-defined event name.
func (WalletProjectManagerProjectCreated) ContractEventName() string {
	return WalletProjectManagerProjectCreatedEventName
}

// UnpackProjectCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProjectCreated(bytes32 indexed projectId, address indexed owner, uint256 extensionId, bytes32 keyType, bytes32 signingAlgo)
func (walletProjectManager *WalletProjectManager) UnpackProjectCreatedEvent(log *types.Log) (*WalletProjectManagerProjectCreated, error) {
	event := "ProjectCreated"
	if len(log.Topics) == 0 || log.Topics[0] != walletProjectManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletProjectManagerProjectCreated)
	if len(log.Data) > 0 {
		if err := walletProjectManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletProjectManager.abi.Events[event].Inputs {
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
func (walletProjectManager *WalletProjectManager) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["AddressAlreadyInSet"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackAddressAlreadyInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["AddressNotInSet"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackAddressNotInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["AvailabilityCheckTimestampInvalid"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackAvailabilityCheckTimestampInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["DuplicatedCosigner"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackDuplicatedCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["InvalidAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackInvalidAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["InvalidKeyType"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackInvalidKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["InvalidPublicKey"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackInvalidPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["InvalidResponseData"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackInvalidResponseDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["InvalidSigningAlgo"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackInvalidSigningAlgoError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["InvalidWalletStatus"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackInvalidWalletStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["KeyTypeNotSupported"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["NoAddresses"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackNoAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["NotOwnerOrPauser"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackNotOwnerOrPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["NotOwnerOrUnpauser"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackNotOwnerOrUnpauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["OnlyExtensionOwner"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackOnlyExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackOnlyOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["OnlyOwnerOrBackupManager"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackOnlyOwnerOrBackupManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["OnlyProposedOwner"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackOnlyProposedOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["OwnerNotAllowed"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackOwnerNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["SigningAlgoNotSupported"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackSigningAlgoNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["VersionNotSupported"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackVersionNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["WalletNotPartOfProject"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackWalletNotPartOfProjectError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectManager.abi.Errors["WalletNotProductionReady"].ID.Bytes()[:4]) {
		return walletProjectManager.UnpackWalletNotProductionReadyError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// WalletProjectManagerAddressAlreadyInSet represents a AddressAlreadyInSet error raised by the WalletProjectManager contract.
type WalletProjectManagerAddressAlreadyInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressAlreadyInSet(address addr)
func WalletProjectManagerAddressAlreadyInSetErrorID() common.Hash {
	return common.HexToHash("0xb725c7fb1c3f0f065973d6b9eb55140f970259955af3045742cd5b37902eb920")
}

// UnpackAddressAlreadyInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressAlreadyInSet(address addr)
func (walletProjectManager *WalletProjectManager) UnpackAddressAlreadyInSetError(raw []byte) (*WalletProjectManagerAddressAlreadyInSet, error) {
	out := new(WalletProjectManagerAddressAlreadyInSet)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "AddressAlreadyInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerAddressNotInSet represents a AddressNotInSet error raised by the WalletProjectManager contract.
type WalletProjectManagerAddressNotInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressNotInSet(address addr)
func WalletProjectManagerAddressNotInSetErrorID() common.Hash {
	return common.HexToHash("0x1ff5ecf3e6d385ea01ff0323fe894ed6d25a83ed5f6ff2a2ede4c897408f6be9")
}

// UnpackAddressNotInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressNotInSet(address addr)
func (walletProjectManager *WalletProjectManager) UnpackAddressNotInSetError(raw []byte) (*WalletProjectManagerAddressNotInSet, error) {
	out := new(WalletProjectManagerAddressNotInSet)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "AddressNotInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerAvailabilityCheckTimestampInvalid represents a AvailabilityCheckTimestampInvalid error raised by the WalletProjectManager contract.
type WalletProjectManagerAvailabilityCheckTimestampInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func WalletProjectManagerAvailabilityCheckTimestampInvalidErrorID() common.Hash {
	return common.HexToHash("0x6344a5e74cff78612fa271fa907ca608af3bad10e1e8c8df302e29ebc6298e2d")
}

// UnpackAvailabilityCheckTimestampInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func (walletProjectManager *WalletProjectManager) UnpackAvailabilityCheckTimestampInvalidError(raw []byte) (*WalletProjectManagerAvailabilityCheckTimestampInvalid, error) {
	out := new(WalletProjectManagerAvailabilityCheckTimestampInvalid)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "AvailabilityCheckTimestampInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerDuplicatedCosigner represents a DuplicatedCosigner error raised by the WalletProjectManager contract.
type WalletProjectManagerDuplicatedCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func WalletProjectManagerDuplicatedCosignerErrorID() common.Hash {
	return common.HexToHash("0x53a5890417576fb3043dc3ff1669b0932da2f657d8adb4c11313b12b711f5910")
}

// UnpackDuplicatedCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func (walletProjectManager *WalletProjectManager) UnpackDuplicatedCosignerError(raw []byte) (*WalletProjectManagerDuplicatedCosigner, error) {
	out := new(WalletProjectManagerDuplicatedCosigner)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "DuplicatedCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerExtensionIdMismatch represents a ExtensionIdMismatch error raised by the WalletProjectManager contract.
type WalletProjectManagerExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func WalletProjectManagerExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (walletProjectManager *WalletProjectManager) UnpackExtensionIdMismatchError(raw []byte) (*WalletProjectManagerExtensionIdMismatch, error) {
	out := new(WalletProjectManagerExtensionIdMismatch)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerInvalidAddress represents a InvalidAddress error raised by the WalletProjectManager contract.
type WalletProjectManagerInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func WalletProjectManagerInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (walletProjectManager *WalletProjectManager) UnpackInvalidAddressError(raw []byte) (*WalletProjectManagerInvalidAddress, error) {
	out := new(WalletProjectManagerInvalidAddress)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerInvalidAvailabilityCheckStatus represents a InvalidAvailabilityCheckStatus error raised by the WalletProjectManager contract.
type WalletProjectManagerInvalidAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func WalletProjectManagerInvalidAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0xedee4ff00b842d3ef84031c60001065c540c838c5479bbf0ae56f15868427e44")
}

// UnpackInvalidAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func (walletProjectManager *WalletProjectManager) UnpackInvalidAvailabilityCheckStatusError(raw []byte) (*WalletProjectManagerInvalidAvailabilityCheckStatus, error) {
	out := new(WalletProjectManagerInvalidAvailabilityCheckStatus)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "InvalidAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerInvalidCosigner represents a InvalidCosigner error raised by the WalletProjectManager contract.
type WalletProjectManagerInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func WalletProjectManagerInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (walletProjectManager *WalletProjectManager) UnpackInvalidCosignerError(raw []byte) (*WalletProjectManagerInvalidCosigner, error) {
	out := new(WalletProjectManagerInvalidCosigner)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerInvalidDuration represents a InvalidDuration error raised by the WalletProjectManager contract.
type WalletProjectManagerInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func WalletProjectManagerInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (walletProjectManager *WalletProjectManager) UnpackInvalidDurationError(raw []byte) (*WalletProjectManagerInvalidDuration, error) {
	out := new(WalletProjectManagerInvalidDuration)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the WalletProjectManager contract.
type WalletProjectManagerInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func WalletProjectManagerInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (walletProjectManager *WalletProjectManager) UnpackInvalidGovernanceHashError(raw []byte) (*WalletProjectManagerInvalidGovernanceHash, error) {
	out := new(WalletProjectManagerInvalidGovernanceHash)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerInvalidKeyType represents a InvalidKeyType error raised by the WalletProjectManager contract.
type WalletProjectManagerInvalidKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyType()
func WalletProjectManagerInvalidKeyTypeErrorID() common.Hash {
	return common.HexToHash("0x36dccb440beb343622d8ee0927bce3f69281feea3a5b0610f6e00e2811b557e1")
}

// UnpackInvalidKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyType()
func (walletProjectManager *WalletProjectManager) UnpackInvalidKeyTypeError(raw []byte) (*WalletProjectManagerInvalidKeyType, error) {
	out := new(WalletProjectManagerInvalidKeyType)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "InvalidKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerInvalidNonce represents a InvalidNonce error raised by the WalletProjectManager contract.
type WalletProjectManagerInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func WalletProjectManagerInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (walletProjectManager *WalletProjectManager) UnpackInvalidNonceError(raw []byte) (*WalletProjectManagerInvalidNonce, error) {
	out := new(WalletProjectManagerInvalidNonce)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerInvalidPublicKey represents a InvalidPublicKey error raised by the WalletProjectManager contract.
type WalletProjectManagerInvalidPublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPublicKey()
func WalletProjectManagerInvalidPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa2d0fee854985cb3d3a394146b5df1c41d5b4d6fcede7fc57a3c05a49b2f3e73")
}

// UnpackInvalidPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPublicKey()
func (walletProjectManager *WalletProjectManager) UnpackInvalidPublicKeyError(raw []byte) (*WalletProjectManagerInvalidPublicKey, error) {
	out := new(WalletProjectManagerInvalidPublicKey)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "InvalidPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerInvalidResponseData represents a InvalidResponseData error raised by the WalletProjectManager contract.
type WalletProjectManagerInvalidResponseData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseData()
func WalletProjectManagerInvalidResponseDataErrorID() common.Hash {
	return common.HexToHash("0x48566617128355a033880805d6462da9bf9c6fe4c61b922997451a71262990af")
}

// UnpackInvalidResponseDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseData()
func (walletProjectManager *WalletProjectManager) UnpackInvalidResponseDataError(raw []byte) (*WalletProjectManagerInvalidResponseData, error) {
	out := new(WalletProjectManagerInvalidResponseData)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "InvalidResponseData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerInvalidSigningAlgo represents a InvalidSigningAlgo error raised by the WalletProjectManager contract.
type WalletProjectManagerInvalidSigningAlgo struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningAlgo()
func WalletProjectManagerInvalidSigningAlgoErrorID() common.Hash {
	return common.HexToHash("0x9b7198a697b61e546ef9ee593a6f77d2bf9f1df9dcbba655ce96d7e2c75346c7")
}

// UnpackInvalidSigningAlgoError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningAlgo()
func (walletProjectManager *WalletProjectManager) UnpackInvalidSigningAlgoError(raw []byte) (*WalletProjectManagerInvalidSigningAlgo, error) {
	out := new(WalletProjectManagerInvalidSigningAlgo)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "InvalidSigningAlgo", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerInvalidThreshold represents a InvalidThreshold error raised by the WalletProjectManager contract.
type WalletProjectManagerInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func WalletProjectManagerInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (walletProjectManager *WalletProjectManager) UnpackInvalidThresholdError(raw []byte) (*WalletProjectManagerInvalidThreshold, error) {
	out := new(WalletProjectManagerInvalidThreshold)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerInvalidWalletStatus represents a InvalidWalletStatus error raised by the WalletProjectManager contract.
type WalletProjectManagerInvalidWalletStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidWalletStatus()
func WalletProjectManagerInvalidWalletStatusErrorID() common.Hash {
	return common.HexToHash("0xd700120e4896ea7835a5e9e4feb690644ced5cad1c52209afdcbe06b96d91345")
}

// UnpackInvalidWalletStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidWalletStatus()
func (walletProjectManager *WalletProjectManager) UnpackInvalidWalletStatusError(raw []byte) (*WalletProjectManagerInvalidWalletStatus, error) {
	out := new(WalletProjectManagerInvalidWalletStatus)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "InvalidWalletStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerKeyTypeNotSupported represents a KeyTypeNotSupported error raised by the WalletProjectManager contract.
type WalletProjectManagerKeyTypeNotSupported struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func WalletProjectManagerKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x85004df1f47814c2de6fbf03d3f2b1ca8c67c182d0d54278766543b6be3bc603")
}

// UnpackKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func (walletProjectManager *WalletProjectManager) UnpackKeyTypeNotSupportedError(raw []byte) (*WalletProjectManagerKeyTypeNotSupported, error) {
	out := new(WalletProjectManagerKeyTypeNotSupported)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "KeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerLengthsMismatch represents a LengthsMismatch error raised by the WalletProjectManager contract.
type WalletProjectManagerLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func WalletProjectManagerLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (walletProjectManager *WalletProjectManager) UnpackLengthsMismatchError(raw []byte) (*WalletProjectManagerLengthsMismatch, error) {
	out := new(WalletProjectManagerLengthsMismatch)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerNoAddresses represents a NoAddresses error raised by the WalletProjectManager contract.
type WalletProjectManagerNoAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoAddresses()
func WalletProjectManagerNoAddressesErrorID() common.Hash {
	return common.HexToHash("0xb2e79f05746fa6c9f3d61f9efe85959371cde5030a6c2dd70c92b63578c5cde7")
}

// UnpackNoAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoAddresses()
func (walletProjectManager *WalletProjectManager) UnpackNoAddressesError(raw []byte) (*WalletProjectManagerNoAddresses, error) {
	out := new(WalletProjectManagerNoAddresses)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "NoAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerNotOwnerOrPauser represents a NotOwnerOrPauser error raised by the WalletProjectManager contract.
type WalletProjectManagerNotOwnerOrPauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrPauser(address caller)
func WalletProjectManagerNotOwnerOrPauserErrorID() common.Hash {
	return common.HexToHash("0x46c0d9af6b895093e022ad2984c54e60192e6e26c5063daf232e825c6e07ff9a")
}

// UnpackNotOwnerOrPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrPauser(address caller)
func (walletProjectManager *WalletProjectManager) UnpackNotOwnerOrPauserError(raw []byte) (*WalletProjectManagerNotOwnerOrPauser, error) {
	out := new(WalletProjectManagerNotOwnerOrPauser)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "NotOwnerOrPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerNotOwnerOrUnpauser represents a NotOwnerOrUnpauser error raised by the WalletProjectManager contract.
type WalletProjectManagerNotOwnerOrUnpauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func WalletProjectManagerNotOwnerOrUnpauserErrorID() common.Hash {
	return common.HexToHash("0x81e700e8f0bd5ba4a6d6a2ce9a1ad0ac9bb1b1b1e62d921a55cbc4b53f1da649")
}

// UnpackNotOwnerOrUnpauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func (walletProjectManager *WalletProjectManager) UnpackNotOwnerOrUnpauserError(raw []byte) (*WalletProjectManagerNotOwnerOrUnpauser, error) {
	out := new(WalletProjectManagerNotOwnerOrUnpauser)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "NotOwnerOrUnpauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerOnlyExtensionOwner represents a OnlyExtensionOwner error raised by the WalletProjectManager contract.
type WalletProjectManagerOnlyExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwner()
func WalletProjectManagerOnlyExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xcd49fd1de5a297cd1ce055d5abffcf2cb25ad9492bd39a703a238af57b2a9488")
}

// UnpackOnlyExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwner()
func (walletProjectManager *WalletProjectManager) UnpackOnlyExtensionOwnerError(raw []byte) (*WalletProjectManagerOnlyExtensionOwner, error) {
	out := new(WalletProjectManagerOnlyExtensionOwner)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "OnlyExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the WalletProjectManager contract.
type WalletProjectManagerOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func WalletProjectManagerOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (walletProjectManager *WalletProjectManager) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*WalletProjectManagerOnlyExtensionOwnerOrOperator, error) {
	out := new(WalletProjectManagerOnlyExtensionOwnerOrOperator)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerOnlyOwner represents a OnlyOwner error raised by the WalletProjectManager contract.
type WalletProjectManagerOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func WalletProjectManagerOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (walletProjectManager *WalletProjectManager) UnpackOnlyOwnerError(raw []byte) (*WalletProjectManagerOnlyOwner, error) {
	out := new(WalletProjectManagerOnlyOwner)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerOnlyOwnerOrBackupManager represents a OnlyOwnerOrBackupManager error raised by the WalletProjectManager contract.
type WalletProjectManagerOnlyOwnerOrBackupManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrBackupManager()
func WalletProjectManagerOnlyOwnerOrBackupManagerErrorID() common.Hash {
	return common.HexToHash("0x249a6cf6e22092f9c80798e2a0c583c0f36aea0779ef38a5b5cdb6a112a0e6df")
}

// UnpackOnlyOwnerOrBackupManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrBackupManager()
func (walletProjectManager *WalletProjectManager) UnpackOnlyOwnerOrBackupManagerError(raw []byte) (*WalletProjectManagerOnlyOwnerOrBackupManager, error) {
	out := new(WalletProjectManagerOnlyOwnerOrBackupManager)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "OnlyOwnerOrBackupManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the WalletProjectManager contract.
type WalletProjectManagerOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func WalletProjectManagerOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (walletProjectManager *WalletProjectManager) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*WalletProjectManagerOnlyProductionOrPausedStatus, error) {
	out := new(WalletProjectManagerOnlyProductionOrPausedStatus)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerOnlyProposedOwner represents a OnlyProposedOwner error raised by the WalletProjectManager contract.
type WalletProjectManagerOnlyProposedOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProposedOwner()
func WalletProjectManagerOnlyProposedOwnerErrorID() common.Hash {
	return common.HexToHash("0x49cb15061ee558e74f66f937f740fd2a3791a89d2d139862222e5c97cabc02cd")
}

// UnpackOnlyProposedOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProposedOwner()
func (walletProjectManager *WalletProjectManager) UnpackOnlyProposedOwnerError(raw []byte) (*WalletProjectManagerOnlyProposedOwner, error) {
	out := new(WalletProjectManagerOnlyProposedOwner)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "OnlyProposedOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerOwnerNotAllowed represents a OwnerNotAllowed error raised by the WalletProjectManager contract.
type WalletProjectManagerOwnerNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerNotAllowed()
func WalletProjectManagerOwnerNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x3d5c9996545c6bf4287c26c1c2d8e6b4c39b92f19ba706c3a756720bf9c4fd24")
}

// UnpackOwnerNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerNotAllowed()
func (walletProjectManager *WalletProjectManager) UnpackOwnerNotAllowedError(raw []byte) (*WalletProjectManagerOwnerNotAllowed, error) {
	out := new(WalletProjectManagerOwnerNotAllowed)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "OwnerNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerSigningAlgoNotSupported represents a SigningAlgoNotSupported error raised by the WalletProjectManager contract.
type WalletProjectManagerSigningAlgoNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SigningAlgoNotSupported()
func WalletProjectManagerSigningAlgoNotSupportedErrorID() common.Hash {
	return common.HexToHash("0xc6d79a2f35820d5b8724cdb275dfb4af83bea4fd323b401243b85be069d73848")
}

// UnpackSigningAlgoNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SigningAlgoNotSupported()
func (walletProjectManager *WalletProjectManager) UnpackSigningAlgoNotSupportedError(raw []byte) (*WalletProjectManagerSigningAlgoNotSupported, error) {
	out := new(WalletProjectManagerSigningAlgoNotSupported)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "SigningAlgoNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the WalletProjectManager contract.
type WalletProjectManagerTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func WalletProjectManagerTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (walletProjectManager *WalletProjectManager) UnpackTeeMachineNotAvailableError(raw []byte) (*WalletProjectManagerTeeMachineNotAvailable, error) {
	out := new(WalletProjectManagerTeeMachineNotAvailable)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerVersionNotSupported represents a VersionNotSupported error raised by the WalletProjectManager contract.
type WalletProjectManagerVersionNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionNotSupported()
func WalletProjectManagerVersionNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x6db0b17b06266b37e721024a8ba1221439799ebb779dfd397865e37250abc9a8")
}

// UnpackVersionNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionNotSupported()
func (walletProjectManager *WalletProjectManager) UnpackVersionNotSupportedError(raw []byte) (*WalletProjectManagerVersionNotSupported, error) {
	out := new(WalletProjectManagerVersionNotSupported)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "VersionNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerWalletNotPartOfProject represents a WalletNotPartOfProject error raised by the WalletProjectManager contract.
type WalletProjectManagerWalletNotPartOfProject struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WalletNotPartOfProject()
func WalletProjectManagerWalletNotPartOfProjectErrorID() common.Hash {
	return common.HexToHash("0xb17858ce8143328aa746a80ee13b0cb7109167ab4a9cd812e533859a299e9728")
}

// UnpackWalletNotPartOfProjectError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WalletNotPartOfProject()
func (walletProjectManager *WalletProjectManager) UnpackWalletNotPartOfProjectError(raw []byte) (*WalletProjectManagerWalletNotPartOfProject, error) {
	out := new(WalletProjectManagerWalletNotPartOfProject)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "WalletNotPartOfProject", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectManagerWalletNotProductionReady represents a WalletNotProductionReady error raised by the WalletProjectManager contract.
type WalletProjectManagerWalletNotProductionReady struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WalletNotProductionReady()
func WalletProjectManagerWalletNotProductionReadyErrorID() common.Hash {
	return common.HexToHash("0xb65a0c136617b5427f37f9704b49dea1a234453e7bb3d5d7aebc523e2ede9f2d")
}

// UnpackWalletNotProductionReadyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WalletNotProductionReady()
func (walletProjectManager *WalletProjectManager) UnpackWalletNotProductionReadyError(raw []byte) (*WalletProjectManagerWalletNotProductionReady, error) {
	out := new(WalletProjectManagerWalletNotProductionReady)
	if err := walletProjectManager.abi.UnpackIntoInterface(out, "WalletNotProductionReady", raw); err != nil {
		return nil, err
	}
	return out, nil
}
