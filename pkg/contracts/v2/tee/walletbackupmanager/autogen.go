// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletbackupmanager

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

// IMachineManagerTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
}

// IWalletBackupManagerBackupId is an auto generated low-level Go binding around an user-defined struct.
type IWalletBackupManagerBackupId struct {
	TeeId         common.Address
	WalletId      [32]byte
	KeyId         uint64
	KeyType       [32]byte
	SigningAlgo   [32]byte
	PublicKey     []byte
	RewardEpochId uint32
	RandomNonce   [32]byte
}

// WalletBackupManagerMetaData contains all meta data concerning the WalletBackupManager contract.
var WalletBackupManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"EmergencyPauseActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMachinePath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardEpochId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeMachine\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyAlreadyAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyNotConfirmed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoActiveMachinePathList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SourceTeeDoesNotHoldKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"BackupRestoreTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceTeeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationTeeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"backupInstructionId\",\"type\":\"bytes32\"}],\"name\":\"DirectBackupTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destinationTeeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"backupInstructionId\",\"type\":\"bytes32\"}],\"name\":\"DirectRestoreTriggered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"randomNonce\",\"type\":\"bytes32\"}],\"internalType\":\"structIWalletBackupManager.BackupId\",\"name\":\"_backupId\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_backupUrl\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"backupRestore\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sourceTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destinationTeeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"directBackup\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destinationTeeId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"randomNonce\",\"type\":\"bytes32\"}],\"internalType\":\"structIWalletBackupManager.BackupId\",\"name\":\"_backupId\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_backupInstructionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"directRestore\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "WalletBackupManager",
}

// WalletBackupManager is an auto generated Go binding around an Ethereum contract.
type WalletBackupManager struct {
	abi abi.ABI
}

// NewWalletBackupManager creates a new instance of WalletBackupManager.
func NewWalletBackupManager() *WalletBackupManager {
	parsed, err := WalletBackupManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &WalletBackupManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *WalletBackupManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackBackupRestore is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20737dd7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function backupRestore(address _teeId, (address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32) _backupId, string _backupUrl, address _claimBackAddress) payable returns()
func (walletBackupManager *WalletBackupManager) PackBackupRestore(teeId common.Address, backupId IWalletBackupManagerBackupId, backupUrl string, claimBackAddress common.Address) []byte {
	enc, err := walletBackupManager.abi.Pack("backupRestore", teeId, backupId, backupUrl, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBackupRestore is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20737dd7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function backupRestore(address _teeId, (address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32) _backupId, string _backupUrl, address _claimBackAddress) payable returns()
func (walletBackupManager *WalletBackupManager) TryPackBackupRestore(teeId common.Address, backupId IWalletBackupManagerBackupId, backupUrl string, claimBackAddress common.Address) ([]byte, error) {
	return walletBackupManager.abi.Pack("backupRestore", teeId, backupId, backupUrl, claimBackAddress)
}

// PackDirectBackup is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x390e69f7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function directBackup(address _sourceTeeId, address _destinationTeeId, bytes32 _walletId, uint64 _keyId, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (walletBackupManager *WalletBackupManager) PackDirectBackup(sourceTeeId common.Address, destinationTeeId common.Address, walletId [32]byte, keyId uint64, claimBackAddress common.Address) []byte {
	enc, err := walletBackupManager.abi.Pack("directBackup", sourceTeeId, destinationTeeId, walletId, keyId, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDirectBackup is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x390e69f7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function directBackup(address _sourceTeeId, address _destinationTeeId, bytes32 _walletId, uint64 _keyId, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (walletBackupManager *WalletBackupManager) TryPackDirectBackup(sourceTeeId common.Address, destinationTeeId common.Address, walletId [32]byte, keyId uint64, claimBackAddress common.Address) ([]byte, error) {
	return walletBackupManager.abi.Pack("directBackup", sourceTeeId, destinationTeeId, walletId, keyId, claimBackAddress)
}

// UnpackDirectBackup is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x390e69f7.
//
// Solidity: function directBackup(address _sourceTeeId, address _destinationTeeId, bytes32 _walletId, uint64 _keyId, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (walletBackupManager *WalletBackupManager) UnpackDirectBackup(data []byte) ([32]byte, error) {
	out, err := walletBackupManager.abi.Unpack("directBackup", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackDirectRestore is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1de9013.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function directRestore(address _destinationTeeId, (address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32) _backupId, bytes32 _backupInstructionId, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (walletBackupManager *WalletBackupManager) PackDirectRestore(destinationTeeId common.Address, backupId IWalletBackupManagerBackupId, backupInstructionId [32]byte, claimBackAddress common.Address) []byte {
	enc, err := walletBackupManager.abi.Pack("directRestore", destinationTeeId, backupId, backupInstructionId, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDirectRestore is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb1de9013.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function directRestore(address _destinationTeeId, (address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32) _backupId, bytes32 _backupInstructionId, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (walletBackupManager *WalletBackupManager) TryPackDirectRestore(destinationTeeId common.Address, backupId IWalletBackupManagerBackupId, backupInstructionId [32]byte, claimBackAddress common.Address) ([]byte, error) {
	return walletBackupManager.abi.Pack("directRestore", destinationTeeId, backupId, backupInstructionId, claimBackAddress)
}

// UnpackDirectRestore is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb1de9013.
//
// Solidity: function directRestore(address _destinationTeeId, (address,bytes32,uint64,bytes32,bytes32,bytes,uint32,bytes32) _backupId, bytes32 _backupInstructionId, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (walletBackupManager *WalletBackupManager) UnpackDirectRestore(data []byte) ([32]byte, error) {
	out, err := walletBackupManager.abi.Unpack("directRestore", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// WalletBackupManagerBackupRestoreTriggered represents a BackupRestoreTriggered event raised by the WalletBackupManager contract.
type WalletBackupManagerBackupRestoreTriggered struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    uint64
	Nonce    *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const WalletBackupManagerBackupRestoreTriggeredEventName = "BackupRestoreTriggered"

// ContractEventName returns the user-defined event name.
func (WalletBackupManagerBackupRestoreTriggered) ContractEventName() string {
	return WalletBackupManagerBackupRestoreTriggeredEventName
}

// UnpackBackupRestoreTriggeredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BackupRestoreTriggered(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId, uint256 nonce)
func (walletBackupManager *WalletBackupManager) UnpackBackupRestoreTriggeredEvent(log *types.Log) (*WalletBackupManagerBackupRestoreTriggered, error) {
	event := "BackupRestoreTriggered"
	if len(log.Topics) == 0 || log.Topics[0] != walletBackupManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletBackupManagerBackupRestoreTriggered)
	if len(log.Data) > 0 {
		if err := walletBackupManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletBackupManager.abi.Events[event].Inputs {
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

// WalletBackupManagerDirectBackupTriggered represents a DirectBackupTriggered event raised by the WalletBackupManager contract.
type WalletBackupManagerDirectBackupTriggered struct {
	SourceTeeId         common.Address
	DestinationTeeId    common.Address
	WalletId            [32]byte
	KeyId               uint64
	BackupInstructionId [32]byte
	Raw                 *types.Log // Blockchain specific contextual infos
}

const WalletBackupManagerDirectBackupTriggeredEventName = "DirectBackupTriggered"

// ContractEventName returns the user-defined event name.
func (WalletBackupManagerDirectBackupTriggered) ContractEventName() string {
	return WalletBackupManagerDirectBackupTriggeredEventName
}

// UnpackDirectBackupTriggeredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DirectBackupTriggered(address indexed sourceTeeId, address indexed destinationTeeId, bytes32 indexed walletId, uint64 keyId, bytes32 backupInstructionId)
func (walletBackupManager *WalletBackupManager) UnpackDirectBackupTriggeredEvent(log *types.Log) (*WalletBackupManagerDirectBackupTriggered, error) {
	event := "DirectBackupTriggered"
	if len(log.Topics) == 0 || log.Topics[0] != walletBackupManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletBackupManagerDirectBackupTriggered)
	if len(log.Data) > 0 {
		if err := walletBackupManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletBackupManager.abi.Events[event].Inputs {
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

// WalletBackupManagerDirectRestoreTriggered represents a DirectRestoreTriggered event raised by the WalletBackupManager contract.
type WalletBackupManagerDirectRestoreTriggered struct {
	DestinationTeeId    common.Address
	WalletId            [32]byte
	KeyId               uint64
	DestinationNonce    *big.Int
	BackupInstructionId [32]byte
	Raw                 *types.Log // Blockchain specific contextual infos
}

const WalletBackupManagerDirectRestoreTriggeredEventName = "DirectRestoreTriggered"

// ContractEventName returns the user-defined event name.
func (WalletBackupManagerDirectRestoreTriggered) ContractEventName() string {
	return WalletBackupManagerDirectRestoreTriggeredEventName
}

// UnpackDirectRestoreTriggeredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DirectRestoreTriggered(address indexed destinationTeeId, bytes32 indexed walletId, uint64 indexed keyId, uint256 destinationNonce, bytes32 backupInstructionId)
func (walletBackupManager *WalletBackupManager) UnpackDirectRestoreTriggeredEvent(log *types.Log) (*WalletBackupManagerDirectRestoreTriggered, error) {
	event := "DirectRestoreTriggered"
	if len(log.Topics) == 0 || log.Topics[0] != walletBackupManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletBackupManagerDirectRestoreTriggered)
	if len(log.Data) > 0 {
		if err := walletBackupManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletBackupManager.abi.Events[event].Inputs {
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

// WalletBackupManagerTeeInstructionsSent represents a TeeInstructionsSent event raised by the WalletBackupManager contract.
type WalletBackupManagerTeeInstructionsSent struct {
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

const WalletBackupManagerTeeInstructionsSentEventName = "TeeInstructionsSent"

// ContractEventName returns the user-defined event name.
func (WalletBackupManagerTeeInstructionsSent) ContractEventName() string {
	return WalletBackupManagerTeeInstructionsSentEventName
}

// UnpackTeeInstructionsSentEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (walletBackupManager *WalletBackupManager) UnpackTeeInstructionsSentEvent(log *types.Log) (*WalletBackupManagerTeeInstructionsSent, error) {
	event := "TeeInstructionsSent"
	if len(log.Topics) == 0 || log.Topics[0] != walletBackupManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletBackupManagerTeeInstructionsSent)
	if len(log.Data) > 0 {
		if err := walletBackupManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletBackupManager.abi.Events[event].Inputs {
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
func (walletBackupManager *WalletBackupManager) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["AddressAlreadyInSet"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackAddressAlreadyInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["AddressNotInSet"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackAddressNotInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["AvailabilityCheckTimestampInvalid"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackAvailabilityCheckTimestampInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["CosignersThresholdTooHigh"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackCosignersThresholdTooHighError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["DuplicatedCosigner"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackDuplicatedCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["EmergencyPauseActive"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackEmergencyPauseActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["FeeTooLow"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackFeeTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidKeyId"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidKeyIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidKeyType"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidMachinePath"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidMachinePathError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidPublicKey"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidResponseData"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidResponseDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidRewardEpochId"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidRewardEpochIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidSigningAlgo"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidSigningAlgoError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidTeeMachine"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidTeeMachineError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["InvalidWalletStatus"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackInvalidWalletStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["KeyAlreadyAvailable"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackKeyAlreadyAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["KeyNotConfirmed"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackKeyNotConfirmedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["KeyTypeNotSupported"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["MessageEmpty"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackMessageEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["NoActiveMachinePathList"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackNoActiveMachinePathListError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["NoAddresses"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackNoAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["NoTeeMachinesSpecified"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackNoTeeMachinesSpecifiedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["NotOwnerOrPauser"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackNotOwnerOrPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["NotOwnerOrUnpauser"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackNotOwnerOrUnpauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["OnlyExtensionOwner"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackOnlyExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackOnlyOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["OnlyOwnerOrBackupManager"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackOnlyOwnerOrBackupManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["OnlyProposedOwner"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackOnlyProposedOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["OperationCommandEmpty"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackOperationCommandEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["OperationTypeEmpty"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackOperationTypeEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["OwnerNotAllowed"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackOwnerNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["SourceTeeDoesNotHoldKey"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackSourceTeeDoesNotHoldKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["TeeNotFound"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackTeeNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletBackupManager.abi.Errors["VersionNotSupported"].ID.Bytes()[:4]) {
		return walletBackupManager.UnpackVersionNotSupportedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// WalletBackupManagerAddressAlreadyInSet represents a AddressAlreadyInSet error raised by the WalletBackupManager contract.
type WalletBackupManagerAddressAlreadyInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressAlreadyInSet(address addr)
func WalletBackupManagerAddressAlreadyInSetErrorID() common.Hash {
	return common.HexToHash("0xb725c7fb1c3f0f065973d6b9eb55140f970259955af3045742cd5b37902eb920")
}

// UnpackAddressAlreadyInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressAlreadyInSet(address addr)
func (walletBackupManager *WalletBackupManager) UnpackAddressAlreadyInSetError(raw []byte) (*WalletBackupManagerAddressAlreadyInSet, error) {
	out := new(WalletBackupManagerAddressAlreadyInSet)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "AddressAlreadyInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerAddressNotInSet represents a AddressNotInSet error raised by the WalletBackupManager contract.
type WalletBackupManagerAddressNotInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressNotInSet(address addr)
func WalletBackupManagerAddressNotInSetErrorID() common.Hash {
	return common.HexToHash("0x1ff5ecf3e6d385ea01ff0323fe894ed6d25a83ed5f6ff2a2ede4c897408f6be9")
}

// UnpackAddressNotInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressNotInSet(address addr)
func (walletBackupManager *WalletBackupManager) UnpackAddressNotInSetError(raw []byte) (*WalletBackupManagerAddressNotInSet, error) {
	out := new(WalletBackupManagerAddressNotInSet)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "AddressNotInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerAvailabilityCheckTimestampInvalid represents a AvailabilityCheckTimestampInvalid error raised by the WalletBackupManager contract.
type WalletBackupManagerAvailabilityCheckTimestampInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func WalletBackupManagerAvailabilityCheckTimestampInvalidErrorID() common.Hash {
	return common.HexToHash("0x6344a5e74cff78612fa271fa907ca608af3bad10e1e8c8df302e29ebc6298e2d")
}

// UnpackAvailabilityCheckTimestampInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func (walletBackupManager *WalletBackupManager) UnpackAvailabilityCheckTimestampInvalidError(raw []byte) (*WalletBackupManagerAvailabilityCheckTimestampInvalid, error) {
	out := new(WalletBackupManagerAvailabilityCheckTimestampInvalid)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "AvailabilityCheckTimestampInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerCosignersThresholdTooHigh represents a CosignersThresholdTooHigh error raised by the WalletBackupManager contract.
type WalletBackupManagerCosignersThresholdTooHigh struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CosignersThresholdTooHigh()
func WalletBackupManagerCosignersThresholdTooHighErrorID() common.Hash {
	return common.HexToHash("0x3d7053512cbb0b649dba01ac4e3591f104fd7d0957128a39a3c552c77cd5014d")
}

// UnpackCosignersThresholdTooHighError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CosignersThresholdTooHigh()
func (walletBackupManager *WalletBackupManager) UnpackCosignersThresholdTooHighError(raw []byte) (*WalletBackupManagerCosignersThresholdTooHigh, error) {
	out := new(WalletBackupManagerCosignersThresholdTooHigh)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "CosignersThresholdTooHigh", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerDuplicatedCosigner represents a DuplicatedCosigner error raised by the WalletBackupManager contract.
type WalletBackupManagerDuplicatedCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func WalletBackupManagerDuplicatedCosignerErrorID() common.Hash {
	return common.HexToHash("0x53a5890417576fb3043dc3ff1669b0932da2f657d8adb4c11313b12b711f5910")
}

// UnpackDuplicatedCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func (walletBackupManager *WalletBackupManager) UnpackDuplicatedCosignerError(raw []byte) (*WalletBackupManagerDuplicatedCosigner, error) {
	out := new(WalletBackupManagerDuplicatedCosigner)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "DuplicatedCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerEmergencyPauseActive represents a EmergencyPauseActive error raised by the WalletBackupManager contract.
type WalletBackupManagerEmergencyPauseActive struct {
	ExtensionId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func WalletBackupManagerEmergencyPauseActiveErrorID() common.Hash {
	return common.HexToHash("0x480cf2163927c55f46b05a5cac80ac3781076b0e50b1ed6b07995a800937d4b6")
}

// UnpackEmergencyPauseActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func (walletBackupManager *WalletBackupManager) UnpackEmergencyPauseActiveError(raw []byte) (*WalletBackupManagerEmergencyPauseActive, error) {
	out := new(WalletBackupManagerEmergencyPauseActive)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "EmergencyPauseActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerExtensionIdMismatch represents a ExtensionIdMismatch error raised by the WalletBackupManager contract.
type WalletBackupManagerExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func WalletBackupManagerExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (walletBackupManager *WalletBackupManager) UnpackExtensionIdMismatchError(raw []byte) (*WalletBackupManagerExtensionIdMismatch, error) {
	out := new(WalletBackupManagerExtensionIdMismatch)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerFeeTooLow represents a FeeTooLow error raised by the WalletBackupManager contract.
type WalletBackupManagerFeeTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeTooLow()
func WalletBackupManagerFeeTooLowErrorID() common.Hash {
	return common.HexToHash("0x732f94137528ce257864c792a6de53ff98f1203051e0a27c454063f559602458")
}

// UnpackFeeTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeTooLow()
func (walletBackupManager *WalletBackupManager) UnpackFeeTooLowError(raw []byte) (*WalletBackupManagerFeeTooLow, error) {
	out := new(WalletBackupManagerFeeTooLow)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "FeeTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidAddress represents a InvalidAddress error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func WalletBackupManagerInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (walletBackupManager *WalletBackupManager) UnpackInvalidAddressError(raw []byte) (*WalletBackupManagerInvalidAddress, error) {
	out := new(WalletBackupManagerInvalidAddress)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidAvailabilityCheckStatus represents a InvalidAvailabilityCheckStatus error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func WalletBackupManagerInvalidAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0xedee4ff00b842d3ef84031c60001065c540c838c5479bbf0ae56f15868427e44")
}

// UnpackInvalidAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func (walletBackupManager *WalletBackupManager) UnpackInvalidAvailabilityCheckStatusError(raw []byte) (*WalletBackupManagerInvalidAvailabilityCheckStatus, error) {
	out := new(WalletBackupManagerInvalidAvailabilityCheckStatus)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidCosigner represents a InvalidCosigner error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func WalletBackupManagerInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (walletBackupManager *WalletBackupManager) UnpackInvalidCosignerError(raw []byte) (*WalletBackupManagerInvalidCosigner, error) {
	out := new(WalletBackupManagerInvalidCosigner)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidDuration represents a InvalidDuration error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func WalletBackupManagerInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (walletBackupManager *WalletBackupManager) UnpackInvalidDurationError(raw []byte) (*WalletBackupManagerInvalidDuration, error) {
	out := new(WalletBackupManagerInvalidDuration)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func WalletBackupManagerInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (walletBackupManager *WalletBackupManager) UnpackInvalidGovernanceHashError(raw []byte) (*WalletBackupManagerInvalidGovernanceHash, error) {
	out := new(WalletBackupManagerInvalidGovernanceHash)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidKeyId represents a InvalidKeyId error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidKeyId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyId()
func WalletBackupManagerInvalidKeyIdErrorID() common.Hash {
	return common.HexToHash("0xb0aeb53e4a83bd3cc8ab2c55982004930d12942a439bdf74bdc4a5d0805f9e0c")
}

// UnpackInvalidKeyIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyId()
func (walletBackupManager *WalletBackupManager) UnpackInvalidKeyIdError(raw []byte) (*WalletBackupManagerInvalidKeyId, error) {
	out := new(WalletBackupManagerInvalidKeyId)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidKeyId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidKeyType represents a InvalidKeyType error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyType()
func WalletBackupManagerInvalidKeyTypeErrorID() common.Hash {
	return common.HexToHash("0x36dccb440beb343622d8ee0927bce3f69281feea3a5b0610f6e00e2811b557e1")
}

// UnpackInvalidKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyType()
func (walletBackupManager *WalletBackupManager) UnpackInvalidKeyTypeError(raw []byte) (*WalletBackupManagerInvalidKeyType, error) {
	out := new(WalletBackupManagerInvalidKeyType)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidMachinePath represents a InvalidMachinePath error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidMachinePath struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidMachinePath()
func WalletBackupManagerInvalidMachinePathErrorID() common.Hash {
	return common.HexToHash("0xf8a0ba22852d6de75ca0c6f463d7f4f70e97b45e489a445014ec85848bf9713b")
}

// UnpackInvalidMachinePathError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidMachinePath()
func (walletBackupManager *WalletBackupManager) UnpackInvalidMachinePathError(raw []byte) (*WalletBackupManagerInvalidMachinePath, error) {
	out := new(WalletBackupManagerInvalidMachinePath)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidMachinePath", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidNonce represents a InvalidNonce error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func WalletBackupManagerInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (walletBackupManager *WalletBackupManager) UnpackInvalidNonceError(raw []byte) (*WalletBackupManagerInvalidNonce, error) {
	out := new(WalletBackupManagerInvalidNonce)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidPublicKey represents a InvalidPublicKey error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidPublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPublicKey()
func WalletBackupManagerInvalidPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa2d0fee854985cb3d3a394146b5df1c41d5b4d6fcede7fc57a3c05a49b2f3e73")
}

// UnpackInvalidPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPublicKey()
func (walletBackupManager *WalletBackupManager) UnpackInvalidPublicKeyError(raw []byte) (*WalletBackupManagerInvalidPublicKey, error) {
	out := new(WalletBackupManagerInvalidPublicKey)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidResponseData represents a InvalidResponseData error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidResponseData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseData()
func WalletBackupManagerInvalidResponseDataErrorID() common.Hash {
	return common.HexToHash("0x48566617128355a033880805d6462da9bf9c6fe4c61b922997451a71262990af")
}

// UnpackInvalidResponseDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseData()
func (walletBackupManager *WalletBackupManager) UnpackInvalidResponseDataError(raw []byte) (*WalletBackupManagerInvalidResponseData, error) {
	out := new(WalletBackupManagerInvalidResponseData)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidResponseData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidRewardEpochId represents a InvalidRewardEpochId error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidRewardEpochId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRewardEpochId()
func WalletBackupManagerInvalidRewardEpochIdErrorID() common.Hash {
	return common.HexToHash("0x16cb8656d2604b0aec9b192e56b5cced5dfaf6d807e2da3f56b8a429e07530f8")
}

// UnpackInvalidRewardEpochIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRewardEpochId()
func (walletBackupManager *WalletBackupManager) UnpackInvalidRewardEpochIdError(raw []byte) (*WalletBackupManagerInvalidRewardEpochId, error) {
	out := new(WalletBackupManagerInvalidRewardEpochId)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidRewardEpochId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidSigningAlgo represents a InvalidSigningAlgo error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidSigningAlgo struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningAlgo()
func WalletBackupManagerInvalidSigningAlgoErrorID() common.Hash {
	return common.HexToHash("0x9b7198a697b61e546ef9ee593a6f77d2bf9f1df9dcbba655ce96d7e2c75346c7")
}

// UnpackInvalidSigningAlgoError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningAlgo()
func (walletBackupManager *WalletBackupManager) UnpackInvalidSigningAlgoError(raw []byte) (*WalletBackupManagerInvalidSigningAlgo, error) {
	out := new(WalletBackupManagerInvalidSigningAlgo)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidSigningAlgo", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidTeeMachine represents a InvalidTeeMachine error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidTeeMachine struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidTeeMachine()
func WalletBackupManagerInvalidTeeMachineErrorID() common.Hash {
	return common.HexToHash("0x5d2417b3b272c7a898abca5e310dd9ca23da671ef358eeb52dcad542b2987108")
}

// UnpackInvalidTeeMachineError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidTeeMachine()
func (walletBackupManager *WalletBackupManager) UnpackInvalidTeeMachineError(raw []byte) (*WalletBackupManagerInvalidTeeMachine, error) {
	out := new(WalletBackupManagerInvalidTeeMachine)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidTeeMachine", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidThreshold represents a InvalidThreshold error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func WalletBackupManagerInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (walletBackupManager *WalletBackupManager) UnpackInvalidThresholdError(raw []byte) (*WalletBackupManagerInvalidThreshold, error) {
	out := new(WalletBackupManagerInvalidThreshold)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerInvalidWalletStatus represents a InvalidWalletStatus error raised by the WalletBackupManager contract.
type WalletBackupManagerInvalidWalletStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidWalletStatus()
func WalletBackupManagerInvalidWalletStatusErrorID() common.Hash {
	return common.HexToHash("0xd700120e4896ea7835a5e9e4feb690644ced5cad1c52209afdcbe06b96d91345")
}

// UnpackInvalidWalletStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidWalletStatus()
func (walletBackupManager *WalletBackupManager) UnpackInvalidWalletStatusError(raw []byte) (*WalletBackupManagerInvalidWalletStatus, error) {
	out := new(WalletBackupManagerInvalidWalletStatus)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "InvalidWalletStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerKeyAlreadyAvailable represents a KeyAlreadyAvailable error raised by the WalletBackupManager contract.
type WalletBackupManagerKeyAlreadyAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyAlreadyAvailable()
func WalletBackupManagerKeyAlreadyAvailableErrorID() common.Hash {
	return common.HexToHash("0xf075e441bd19e4850eb60abe55ec35697892ad75b87e50818f22d0f063ef4675")
}

// UnpackKeyAlreadyAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyAlreadyAvailable()
func (walletBackupManager *WalletBackupManager) UnpackKeyAlreadyAvailableError(raw []byte) (*WalletBackupManagerKeyAlreadyAvailable, error) {
	out := new(WalletBackupManagerKeyAlreadyAvailable)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "KeyAlreadyAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerKeyNotConfirmed represents a KeyNotConfirmed error raised by the WalletBackupManager contract.
type WalletBackupManagerKeyNotConfirmed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyNotConfirmed()
func WalletBackupManagerKeyNotConfirmedErrorID() common.Hash {
	return common.HexToHash("0x87c78cc790b5be0545b0c2a6598626fb2d852bb89ac5ebbdba786b484ec50f36")
}

// UnpackKeyNotConfirmedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyNotConfirmed()
func (walletBackupManager *WalletBackupManager) UnpackKeyNotConfirmedError(raw []byte) (*WalletBackupManagerKeyNotConfirmed, error) {
	out := new(WalletBackupManagerKeyNotConfirmed)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "KeyNotConfirmed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerKeyTypeNotSupported represents a KeyTypeNotSupported error raised by the WalletBackupManager contract.
type WalletBackupManagerKeyTypeNotSupported struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func WalletBackupManagerKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x85004df1f47814c2de6fbf03d3f2b1ca8c67c182d0d54278766543b6be3bc603")
}

// UnpackKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func (walletBackupManager *WalletBackupManager) UnpackKeyTypeNotSupportedError(raw []byte) (*WalletBackupManagerKeyTypeNotSupported, error) {
	out := new(WalletBackupManagerKeyTypeNotSupported)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "KeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerLengthsMismatch represents a LengthsMismatch error raised by the WalletBackupManager contract.
type WalletBackupManagerLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func WalletBackupManagerLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (walletBackupManager *WalletBackupManager) UnpackLengthsMismatchError(raw []byte) (*WalletBackupManagerLengthsMismatch, error) {
	out := new(WalletBackupManagerLengthsMismatch)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerMessageEmpty represents a MessageEmpty error raised by the WalletBackupManager contract.
type WalletBackupManagerMessageEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MessageEmpty()
func WalletBackupManagerMessageEmptyErrorID() common.Hash {
	return common.HexToHash("0x3b06a1beedd2da56654bd6f78f50f74e9bc7469bae35c52fd3680e4d95103c79")
}

// UnpackMessageEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MessageEmpty()
func (walletBackupManager *WalletBackupManager) UnpackMessageEmptyError(raw []byte) (*WalletBackupManagerMessageEmpty, error) {
	out := new(WalletBackupManagerMessageEmpty)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "MessageEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerNoActiveMachinePathList represents a NoActiveMachinePathList error raised by the WalletBackupManager contract.
type WalletBackupManagerNoActiveMachinePathList struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoActiveMachinePathList()
func WalletBackupManagerNoActiveMachinePathListErrorID() common.Hash {
	return common.HexToHash("0xb0dae5768e38e4e9d9ec3149b08b494bbe09867f8733737d3ea98e00826652f2")
}

// UnpackNoActiveMachinePathListError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoActiveMachinePathList()
func (walletBackupManager *WalletBackupManager) UnpackNoActiveMachinePathListError(raw []byte) (*WalletBackupManagerNoActiveMachinePathList, error) {
	out := new(WalletBackupManagerNoActiveMachinePathList)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "NoActiveMachinePathList", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerNoAddresses represents a NoAddresses error raised by the WalletBackupManager contract.
type WalletBackupManagerNoAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoAddresses()
func WalletBackupManagerNoAddressesErrorID() common.Hash {
	return common.HexToHash("0xb2e79f05746fa6c9f3d61f9efe85959371cde5030a6c2dd70c92b63578c5cde7")
}

// UnpackNoAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoAddresses()
func (walletBackupManager *WalletBackupManager) UnpackNoAddressesError(raw []byte) (*WalletBackupManagerNoAddresses, error) {
	out := new(WalletBackupManagerNoAddresses)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "NoAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerNoTeeMachinesSpecified represents a NoTeeMachinesSpecified error raised by the WalletBackupManager contract.
type WalletBackupManagerNoTeeMachinesSpecified struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoTeeMachinesSpecified()
func WalletBackupManagerNoTeeMachinesSpecifiedErrorID() common.Hash {
	return common.HexToHash("0xe10387d0118f08d51d80fe03b0c14ed458c46aac2ae9964e0b628978e6938db8")
}

// UnpackNoTeeMachinesSpecifiedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoTeeMachinesSpecified()
func (walletBackupManager *WalletBackupManager) UnpackNoTeeMachinesSpecifiedError(raw []byte) (*WalletBackupManagerNoTeeMachinesSpecified, error) {
	out := new(WalletBackupManagerNoTeeMachinesSpecified)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "NoTeeMachinesSpecified", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerNotOwnerOrPauser represents a NotOwnerOrPauser error raised by the WalletBackupManager contract.
type WalletBackupManagerNotOwnerOrPauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrPauser(address caller)
func WalletBackupManagerNotOwnerOrPauserErrorID() common.Hash {
	return common.HexToHash("0x46c0d9af6b895093e022ad2984c54e60192e6e26c5063daf232e825c6e07ff9a")
}

// UnpackNotOwnerOrPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrPauser(address caller)
func (walletBackupManager *WalletBackupManager) UnpackNotOwnerOrPauserError(raw []byte) (*WalletBackupManagerNotOwnerOrPauser, error) {
	out := new(WalletBackupManagerNotOwnerOrPauser)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "NotOwnerOrPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerNotOwnerOrUnpauser represents a NotOwnerOrUnpauser error raised by the WalletBackupManager contract.
type WalletBackupManagerNotOwnerOrUnpauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func WalletBackupManagerNotOwnerOrUnpauserErrorID() common.Hash {
	return common.HexToHash("0x81e700e8f0bd5ba4a6d6a2ce9a1ad0ac9bb1b1b1e62d921a55cbc4b53f1da649")
}

// UnpackNotOwnerOrUnpauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func (walletBackupManager *WalletBackupManager) UnpackNotOwnerOrUnpauserError(raw []byte) (*WalletBackupManagerNotOwnerOrUnpauser, error) {
	out := new(WalletBackupManagerNotOwnerOrUnpauser)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "NotOwnerOrUnpauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerOnlyExtensionOwner represents a OnlyExtensionOwner error raised by the WalletBackupManager contract.
type WalletBackupManagerOnlyExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwner()
func WalletBackupManagerOnlyExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xcd49fd1de5a297cd1ce055d5abffcf2cb25ad9492bd39a703a238af57b2a9488")
}

// UnpackOnlyExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwner()
func (walletBackupManager *WalletBackupManager) UnpackOnlyExtensionOwnerError(raw []byte) (*WalletBackupManagerOnlyExtensionOwner, error) {
	out := new(WalletBackupManagerOnlyExtensionOwner)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "OnlyExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the WalletBackupManager contract.
type WalletBackupManagerOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func WalletBackupManagerOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (walletBackupManager *WalletBackupManager) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*WalletBackupManagerOnlyExtensionOwnerOrOperator, error) {
	out := new(WalletBackupManagerOnlyExtensionOwnerOrOperator)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerOnlyOwner represents a OnlyOwner error raised by the WalletBackupManager contract.
type WalletBackupManagerOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func WalletBackupManagerOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (walletBackupManager *WalletBackupManager) UnpackOnlyOwnerError(raw []byte) (*WalletBackupManagerOnlyOwner, error) {
	out := new(WalletBackupManagerOnlyOwner)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerOnlyOwnerOrBackupManager represents a OnlyOwnerOrBackupManager error raised by the WalletBackupManager contract.
type WalletBackupManagerOnlyOwnerOrBackupManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrBackupManager()
func WalletBackupManagerOnlyOwnerOrBackupManagerErrorID() common.Hash {
	return common.HexToHash("0x249a6cf6e22092f9c80798e2a0c583c0f36aea0779ef38a5b5cdb6a112a0e6df")
}

// UnpackOnlyOwnerOrBackupManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrBackupManager()
func (walletBackupManager *WalletBackupManager) UnpackOnlyOwnerOrBackupManagerError(raw []byte) (*WalletBackupManagerOnlyOwnerOrBackupManager, error) {
	out := new(WalletBackupManagerOnlyOwnerOrBackupManager)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "OnlyOwnerOrBackupManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the WalletBackupManager contract.
type WalletBackupManagerOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func WalletBackupManagerOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (walletBackupManager *WalletBackupManager) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*WalletBackupManagerOnlyProductionOrPausedStatus, error) {
	out := new(WalletBackupManagerOnlyProductionOrPausedStatus)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerOnlyProposedOwner represents a OnlyProposedOwner error raised by the WalletBackupManager contract.
type WalletBackupManagerOnlyProposedOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProposedOwner()
func WalletBackupManagerOnlyProposedOwnerErrorID() common.Hash {
	return common.HexToHash("0x49cb15061ee558e74f66f937f740fd2a3791a89d2d139862222e5c97cabc02cd")
}

// UnpackOnlyProposedOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProposedOwner()
func (walletBackupManager *WalletBackupManager) UnpackOnlyProposedOwnerError(raw []byte) (*WalletBackupManagerOnlyProposedOwner, error) {
	out := new(WalletBackupManagerOnlyProposedOwner)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "OnlyProposedOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerOperationCommandEmpty represents a OperationCommandEmpty error raised by the WalletBackupManager contract.
type WalletBackupManagerOperationCommandEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationCommandEmpty()
func WalletBackupManagerOperationCommandEmptyErrorID() common.Hash {
	return common.HexToHash("0x038dc5a8067401ab291233391b3d11ff3ce1b21a8aebcf4297e6a9f759f65846")
}

// UnpackOperationCommandEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationCommandEmpty()
func (walletBackupManager *WalletBackupManager) UnpackOperationCommandEmptyError(raw []byte) (*WalletBackupManagerOperationCommandEmpty, error) {
	out := new(WalletBackupManagerOperationCommandEmpty)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "OperationCommandEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerOperationTypeEmpty represents a OperationTypeEmpty error raised by the WalletBackupManager contract.
type WalletBackupManagerOperationTypeEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationTypeEmpty()
func WalletBackupManagerOperationTypeEmptyErrorID() common.Hash {
	return common.HexToHash("0x06e3960d8a83ea8972d826357113f286ce60947bac56680f119fca98a2d6125b")
}

// UnpackOperationTypeEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationTypeEmpty()
func (walletBackupManager *WalletBackupManager) UnpackOperationTypeEmptyError(raw []byte) (*WalletBackupManagerOperationTypeEmpty, error) {
	out := new(WalletBackupManagerOperationTypeEmpty)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "OperationTypeEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerOwnerNotAllowed represents a OwnerNotAllowed error raised by the WalletBackupManager contract.
type WalletBackupManagerOwnerNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerNotAllowed()
func WalletBackupManagerOwnerNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x3d5c9996545c6bf4287c26c1c2d8e6b4c39b92f19ba706c3a756720bf9c4fd24")
}

// UnpackOwnerNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerNotAllowed()
func (walletBackupManager *WalletBackupManager) UnpackOwnerNotAllowedError(raw []byte) (*WalletBackupManagerOwnerNotAllowed, error) {
	out := new(WalletBackupManagerOwnerNotAllowed)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "OwnerNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerSourceTeeDoesNotHoldKey represents a SourceTeeDoesNotHoldKey error raised by the WalletBackupManager contract.
type WalletBackupManagerSourceTeeDoesNotHoldKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceTeeDoesNotHoldKey()
func WalletBackupManagerSourceTeeDoesNotHoldKeyErrorID() common.Hash {
	return common.HexToHash("0xad7f3b2423c8af25e7d0db49bdcd5fadf7332fca434abb1e71c803397a846039")
}

// UnpackSourceTeeDoesNotHoldKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceTeeDoesNotHoldKey()
func (walletBackupManager *WalletBackupManager) UnpackSourceTeeDoesNotHoldKeyError(raw []byte) (*WalletBackupManagerSourceTeeDoesNotHoldKey, error) {
	out := new(WalletBackupManagerSourceTeeDoesNotHoldKey)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "SourceTeeDoesNotHoldKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the WalletBackupManager contract.
type WalletBackupManagerTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func WalletBackupManagerTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (walletBackupManager *WalletBackupManager) UnpackTeeMachineNotAvailableError(raw []byte) (*WalletBackupManagerTeeMachineNotAvailable, error) {
	out := new(WalletBackupManagerTeeMachineNotAvailable)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerTeeNotFound represents a TeeNotFound error raised by the WalletBackupManager contract.
type WalletBackupManagerTeeNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeNotFound()
func WalletBackupManagerTeeNotFoundErrorID() common.Hash {
	return common.HexToHash("0xceb05b6852ed4c7dc3a718eb0d26555398a76877a44dd41dc92c5fe7e6ef8c2d")
}

// UnpackTeeNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeNotFound()
func (walletBackupManager *WalletBackupManager) UnpackTeeNotFoundError(raw []byte) (*WalletBackupManagerTeeNotFound, error) {
	out := new(WalletBackupManagerTeeNotFound)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "TeeNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletBackupManagerVersionNotSupported represents a VersionNotSupported error raised by the WalletBackupManager contract.
type WalletBackupManagerVersionNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionNotSupported()
func WalletBackupManagerVersionNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x6db0b17b06266b37e721024a8ba1221439799ebb779dfd397865e37250abc9a8")
}

// UnpackVersionNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionNotSupported()
func (walletBackupManager *WalletBackupManager) UnpackVersionNotSupportedError(raw []byte) (*WalletBackupManagerVersionNotSupported, error) {
	out := new(WalletBackupManagerVersionNotSupported)
	if err := walletBackupManager.abi.UnpackIntoInterface(out, "VersionNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}
