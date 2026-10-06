// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package instructions

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

// IInstructionsTeeInstructionParams is an auto generated low-level Go binding around an user-defined struct.
type IInstructionsTeeInstructionParams struct {
	OpType             [32]byte
	OpCommand          [32]byte
	Message            []byte
	Cosigners          []common.Address
	CosignersThreshold uint64
	ClaimBackAddress   common.Address
}

// IMachineManagerTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
}

// InstructionsMetaData contains all meta data concerning the Instructions contract.
var InstructionsMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"EmergencyPauseActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInstructionsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyInstructionsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemInstructionsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"instructionsSender\",\"type\":\"address\"}],\"name\":\"SystemInstructionsSenderAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"instructionsSender\",\"type\":\"address\"}],\"name\":\"SystemInstructionsSenderNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"name\":\"SystemOpTypeNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"instructionsSenders\",\"type\":\"address[]\"}],\"name\":\"SystemInstructionsSendersRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"instructionsSenders\",\"type\":\"address[]\"}],\"name\":\"SystemInstructionsSendersUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getSystemInstructionsSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_instructionsSenders\",\"type\":\"address[]\"}],\"name\":\"registerSystemInstructionsSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"}],\"internalType\":\"structIInstructions.TeeInstructionParams\",\"name\":\"_instructionParams\",\"type\":\"tuple\"}],\"name\":\"sendInstructions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"}],\"internalType\":\"structIInstructions.TeeInstructionParams\",\"name\":\"_instructionParams\",\"type\":\"tuple\"}],\"name\":\"sendSystemInstructions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_instructionsSenders\",\"type\":\"address[]\"}],\"name\":\"unregisterSystemInstructionsSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Instructions",
}

// Instructions is an auto generated Go binding around an Ethereum contract.
type Instructions struct {
	abi abi.ABI
}

// NewInstructions creates a new instance of Instructions.
func NewInstructions() *Instructions {
	parsed, err := InstructionsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Instructions{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Instructions) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetSystemInstructionsSenders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa37768e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSystemInstructionsSenders() view returns(address[])
func (instructions *Instructions) PackGetSystemInstructionsSenders() []byte {
	enc, err := instructions.abi.Pack("getSystemInstructionsSenders")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSystemInstructionsSenders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa37768e5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSystemInstructionsSenders() view returns(address[])
func (instructions *Instructions) TryPackGetSystemInstructionsSenders() ([]byte, error) {
	return instructions.abi.Pack("getSystemInstructionsSenders")
}

// UnpackGetSystemInstructionsSenders is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa37768e5.
//
// Solidity: function getSystemInstructionsSenders() view returns(address[])
func (instructions *Instructions) UnpackGetSystemInstructionsSenders(data []byte) ([]common.Address, error) {
	out, err := instructions.abi.Unpack("getSystemInstructionsSenders", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackRegisterSystemInstructionsSenders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4d9003cf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function registerSystemInstructionsSenders(address[] _instructionsSenders) returns()
func (instructions *Instructions) PackRegisterSystemInstructionsSenders(instructionsSenders []common.Address) []byte {
	enc, err := instructions.abi.Pack("registerSystemInstructionsSenders", instructionsSenders)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegisterSystemInstructionsSenders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4d9003cf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function registerSystemInstructionsSenders(address[] _instructionsSenders) returns()
func (instructions *Instructions) TryPackRegisterSystemInstructionsSenders(instructionsSenders []common.Address) ([]byte, error) {
	return instructions.abi.Pack("registerSystemInstructionsSenders", instructionsSenders)
}

// PackSendInstructions is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf731df53.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sendInstructions(address[] _teeIds, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (instructions *Instructions) PackSendInstructions(teeIds []common.Address, instructionParams IInstructionsTeeInstructionParams) []byte {
	enc, err := instructions.abi.Pack("sendInstructions", teeIds, instructionParams)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSendInstructions is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf731df53.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sendInstructions(address[] _teeIds, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (instructions *Instructions) TryPackSendInstructions(teeIds []common.Address, instructionParams IInstructionsTeeInstructionParams) ([]byte, error) {
	return instructions.abi.Pack("sendInstructions", teeIds, instructionParams)
}

// UnpackSendInstructions is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf731df53.
//
// Solidity: function sendInstructions(address[] _teeIds, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (instructions *Instructions) UnpackSendInstructions(data []byte) ([32]byte, error) {
	out, err := instructions.abi.Unpack("sendInstructions", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackSendSystemInstructions is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1456290e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function sendSystemInstructions(bytes32 _instructionId, address[] _teeIds, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (instructions *Instructions) PackSendSystemInstructions(instructionId [32]byte, teeIds []common.Address, instructionParams IInstructionsTeeInstructionParams) []byte {
	enc, err := instructions.abi.Pack("sendSystemInstructions", instructionId, teeIds, instructionParams)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSendSystemInstructions is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1456290e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function sendSystemInstructions(bytes32 _instructionId, address[] _teeIds, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (instructions *Instructions) TryPackSendSystemInstructions(instructionId [32]byte, teeIds []common.Address, instructionParams IInstructionsTeeInstructionParams) ([]byte, error) {
	return instructions.abi.Pack("sendSystemInstructions", instructionId, teeIds, instructionParams)
}

// UnpackSendSystemInstructions is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1456290e.
//
// Solidity: function sendSystemInstructions(bytes32 _instructionId, address[] _teeIds, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (instructions *Instructions) UnpackSendSystemInstructions(data []byte) ([32]byte, error) {
	out, err := instructions.abi.Unpack("sendSystemInstructions", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackUnregisterSystemInstructionsSenders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf123abf3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unregisterSystemInstructionsSenders(address[] _instructionsSenders) returns()
func (instructions *Instructions) PackUnregisterSystemInstructionsSenders(instructionsSenders []common.Address) []byte {
	enc, err := instructions.abi.Pack("unregisterSystemInstructionsSenders", instructionsSenders)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnregisterSystemInstructionsSenders is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf123abf3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unregisterSystemInstructionsSenders(address[] _instructionsSenders) returns()
func (instructions *Instructions) TryPackUnregisterSystemInstructionsSenders(instructionsSenders []common.Address) ([]byte, error) {
	return instructions.abi.Pack("unregisterSystemInstructionsSenders", instructionsSenders)
}

// InstructionsGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the Instructions contract.
type InstructionsGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const InstructionsGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (InstructionsGovernanceCallTimelocked) ContractEventName() string {
	return InstructionsGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (instructions *Instructions) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*InstructionsGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != instructions.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(InstructionsGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := instructions.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range instructions.abi.Events[event].Inputs {
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

// InstructionsInitialized represents a Initialized event raised by the Instructions contract.
type InstructionsInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const InstructionsInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (InstructionsInitialized) ContractEventName() string {
	return InstructionsInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (instructions *Instructions) UnpackInitializedEvent(log *types.Log) (*InstructionsInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != instructions.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(InstructionsInitialized)
	if len(log.Data) > 0 {
		if err := instructions.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range instructions.abi.Events[event].Inputs {
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

// InstructionsSystemInstructionsSendersRegistered represents a SystemInstructionsSendersRegistered event raised by the Instructions contract.
type InstructionsSystemInstructionsSendersRegistered struct {
	InstructionsSenders []common.Address
	Raw                 *types.Log // Blockchain specific contextual infos
}

const InstructionsSystemInstructionsSendersRegisteredEventName = "SystemInstructionsSendersRegistered"

// ContractEventName returns the user-defined event name.
func (InstructionsSystemInstructionsSendersRegistered) ContractEventName() string {
	return InstructionsSystemInstructionsSendersRegisteredEventName
}

// UnpackSystemInstructionsSendersRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SystemInstructionsSendersRegistered(address[] instructionsSenders)
func (instructions *Instructions) UnpackSystemInstructionsSendersRegisteredEvent(log *types.Log) (*InstructionsSystemInstructionsSendersRegistered, error) {
	event := "SystemInstructionsSendersRegistered"
	if len(log.Topics) == 0 || log.Topics[0] != instructions.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(InstructionsSystemInstructionsSendersRegistered)
	if len(log.Data) > 0 {
		if err := instructions.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range instructions.abi.Events[event].Inputs {
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

// InstructionsSystemInstructionsSendersUnregistered represents a SystemInstructionsSendersUnregistered event raised by the Instructions contract.
type InstructionsSystemInstructionsSendersUnregistered struct {
	InstructionsSenders []common.Address
	Raw                 *types.Log // Blockchain specific contextual infos
}

const InstructionsSystemInstructionsSendersUnregisteredEventName = "SystemInstructionsSendersUnregistered"

// ContractEventName returns the user-defined event name.
func (InstructionsSystemInstructionsSendersUnregistered) ContractEventName() string {
	return InstructionsSystemInstructionsSendersUnregisteredEventName
}

// UnpackSystemInstructionsSendersUnregisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SystemInstructionsSendersUnregistered(address[] instructionsSenders)
func (instructions *Instructions) UnpackSystemInstructionsSendersUnregisteredEvent(log *types.Log) (*InstructionsSystemInstructionsSendersUnregistered, error) {
	event := "SystemInstructionsSendersUnregistered"
	if len(log.Topics) == 0 || log.Topics[0] != instructions.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(InstructionsSystemInstructionsSendersUnregistered)
	if len(log.Data) > 0 {
		if err := instructions.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range instructions.abi.Events[event].Inputs {
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

// InstructionsTeeInstructionsSent represents a TeeInstructionsSent event raised by the Instructions contract.
type InstructionsTeeInstructionsSent struct {
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

const InstructionsTeeInstructionsSentEventName = "TeeInstructionsSent"

// ContractEventName returns the user-defined event name.
func (InstructionsTeeInstructionsSent) ContractEventName() string {
	return InstructionsTeeInstructionsSentEventName
}

// UnpackTeeInstructionsSentEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (instructions *Instructions) UnpackTeeInstructionsSentEvent(log *types.Log) (*InstructionsTeeInstructionsSent, error) {
	event := "TeeInstructionsSent"
	if len(log.Topics) == 0 || log.Topics[0] != instructions.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(InstructionsTeeInstructionsSent)
	if len(log.Data) > 0 {
		if err := instructions.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range instructions.abi.Events[event].Inputs {
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
func (instructions *Instructions) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], instructions.abi.Errors["AddressAlreadyInSet"].ID.Bytes()[:4]) {
		return instructions.UnpackAddressAlreadyInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["AddressNotInSet"].ID.Bytes()[:4]) {
		return instructions.UnpackAddressNotInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["AvailabilityCheckTimestampInvalid"].ID.Bytes()[:4]) {
		return instructions.UnpackAvailabilityCheckTimestampInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["CosignersThresholdTooHigh"].ID.Bytes()[:4]) {
		return instructions.UnpackCosignersThresholdTooHighError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["DuplicatedCosigner"].ID.Bytes()[:4]) {
		return instructions.UnpackDuplicatedCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["EmergencyPauseActive"].ID.Bytes()[:4]) {
		return instructions.UnpackEmergencyPauseActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return instructions.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["FeeTooLow"].ID.Bytes()[:4]) {
		return instructions.UnpackFeeTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidInstructionsSender"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidInstructionsSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidKeyType"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidPublicKey"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidResponseData"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidResponseDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidSigningAlgo"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidSigningAlgoError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["InvalidWalletStatus"].ID.Bytes()[:4]) {
		return instructions.UnpackInvalidWalletStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["KeyTypeNotSupported"].ID.Bytes()[:4]) {
		return instructions.UnpackKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return instructions.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["MessageEmpty"].ID.Bytes()[:4]) {
		return instructions.UnpackMessageEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["NoAddresses"].ID.Bytes()[:4]) {
		return instructions.UnpackNoAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["NoTeeMachinesSpecified"].ID.Bytes()[:4]) {
		return instructions.UnpackNoTeeMachinesSpecifiedError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return instructions.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["NotOwnerOrPauser"].ID.Bytes()[:4]) {
		return instructions.UnpackNotOwnerOrPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["NotOwnerOrUnpauser"].ID.Bytes()[:4]) {
		return instructions.UnpackNotOwnerOrUnpauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["OnlyExtensionOwner"].ID.Bytes()[:4]) {
		return instructions.UnpackOnlyExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return instructions.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return instructions.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["OnlyInstructionsSender"].ID.Bytes()[:4]) {
		return instructions.UnpackOnlyInstructionsSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return instructions.UnpackOnlyOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["OnlyOwnerOrBackupManager"].ID.Bytes()[:4]) {
		return instructions.UnpackOnlyOwnerOrBackupManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return instructions.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["OnlyProposedOwner"].ID.Bytes()[:4]) {
		return instructions.UnpackOnlyProposedOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["OnlySystemInstructionsSender"].ID.Bytes()[:4]) {
		return instructions.UnpackOnlySystemInstructionsSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["OperationCommandEmpty"].ID.Bytes()[:4]) {
		return instructions.UnpackOperationCommandEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["OperationTypeEmpty"].ID.Bytes()[:4]) {
		return instructions.UnpackOperationTypeEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["OwnerNotAllowed"].ID.Bytes()[:4]) {
		return instructions.UnpackOwnerNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["SystemInstructionsSenderAlreadyExists"].ID.Bytes()[:4]) {
		return instructions.UnpackSystemInstructionsSenderAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["SystemInstructionsSenderNotFound"].ID.Bytes()[:4]) {
		return instructions.UnpackSystemInstructionsSenderNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["SystemOpTypeNotAllowed"].ID.Bytes()[:4]) {
		return instructions.UnpackSystemOpTypeNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return instructions.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["TeeNotFound"].ID.Bytes()[:4]) {
		return instructions.UnpackTeeNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], instructions.abi.Errors["VersionNotSupported"].ID.Bytes()[:4]) {
		return instructions.UnpackVersionNotSupportedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// InstructionsAddressAlreadyInSet represents a AddressAlreadyInSet error raised by the Instructions contract.
type InstructionsAddressAlreadyInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressAlreadyInSet(address addr)
func InstructionsAddressAlreadyInSetErrorID() common.Hash {
	return common.HexToHash("0xb725c7fb1c3f0f065973d6b9eb55140f970259955af3045742cd5b37902eb920")
}

// UnpackAddressAlreadyInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressAlreadyInSet(address addr)
func (instructions *Instructions) UnpackAddressAlreadyInSetError(raw []byte) (*InstructionsAddressAlreadyInSet, error) {
	out := new(InstructionsAddressAlreadyInSet)
	if err := instructions.abi.UnpackIntoInterface(out, "AddressAlreadyInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsAddressNotInSet represents a AddressNotInSet error raised by the Instructions contract.
type InstructionsAddressNotInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressNotInSet(address addr)
func InstructionsAddressNotInSetErrorID() common.Hash {
	return common.HexToHash("0x1ff5ecf3e6d385ea01ff0323fe894ed6d25a83ed5f6ff2a2ede4c897408f6be9")
}

// UnpackAddressNotInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressNotInSet(address addr)
func (instructions *Instructions) UnpackAddressNotInSetError(raw []byte) (*InstructionsAddressNotInSet, error) {
	out := new(InstructionsAddressNotInSet)
	if err := instructions.abi.UnpackIntoInterface(out, "AddressNotInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsAvailabilityCheckTimestampInvalid represents a AvailabilityCheckTimestampInvalid error raised by the Instructions contract.
type InstructionsAvailabilityCheckTimestampInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func InstructionsAvailabilityCheckTimestampInvalidErrorID() common.Hash {
	return common.HexToHash("0x6344a5e74cff78612fa271fa907ca608af3bad10e1e8c8df302e29ebc6298e2d")
}

// UnpackAvailabilityCheckTimestampInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func (instructions *Instructions) UnpackAvailabilityCheckTimestampInvalidError(raw []byte) (*InstructionsAvailabilityCheckTimestampInvalid, error) {
	out := new(InstructionsAvailabilityCheckTimestampInvalid)
	if err := instructions.abi.UnpackIntoInterface(out, "AvailabilityCheckTimestampInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsCosignersThresholdTooHigh represents a CosignersThresholdTooHigh error raised by the Instructions contract.
type InstructionsCosignersThresholdTooHigh struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CosignersThresholdTooHigh()
func InstructionsCosignersThresholdTooHighErrorID() common.Hash {
	return common.HexToHash("0x3d7053512cbb0b649dba01ac4e3591f104fd7d0957128a39a3c552c77cd5014d")
}

// UnpackCosignersThresholdTooHighError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CosignersThresholdTooHigh()
func (instructions *Instructions) UnpackCosignersThresholdTooHighError(raw []byte) (*InstructionsCosignersThresholdTooHigh, error) {
	out := new(InstructionsCosignersThresholdTooHigh)
	if err := instructions.abi.UnpackIntoInterface(out, "CosignersThresholdTooHigh", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsDuplicatedCosigner represents a DuplicatedCosigner error raised by the Instructions contract.
type InstructionsDuplicatedCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func InstructionsDuplicatedCosignerErrorID() common.Hash {
	return common.HexToHash("0x53a5890417576fb3043dc3ff1669b0932da2f657d8adb4c11313b12b711f5910")
}

// UnpackDuplicatedCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func (instructions *Instructions) UnpackDuplicatedCosignerError(raw []byte) (*InstructionsDuplicatedCosigner, error) {
	out := new(InstructionsDuplicatedCosigner)
	if err := instructions.abi.UnpackIntoInterface(out, "DuplicatedCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsEmergencyPauseActive represents a EmergencyPauseActive error raised by the Instructions contract.
type InstructionsEmergencyPauseActive struct {
	ExtensionId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func InstructionsEmergencyPauseActiveErrorID() common.Hash {
	return common.HexToHash("0x480cf2163927c55f46b05a5cac80ac3781076b0e50b1ed6b07995a800937d4b6")
}

// UnpackEmergencyPauseActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func (instructions *Instructions) UnpackEmergencyPauseActiveError(raw []byte) (*InstructionsEmergencyPauseActive, error) {
	out := new(InstructionsEmergencyPauseActive)
	if err := instructions.abi.UnpackIntoInterface(out, "EmergencyPauseActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsExtensionIdMismatch represents a ExtensionIdMismatch error raised by the Instructions contract.
type InstructionsExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func InstructionsExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (instructions *Instructions) UnpackExtensionIdMismatchError(raw []byte) (*InstructionsExtensionIdMismatch, error) {
	out := new(InstructionsExtensionIdMismatch)
	if err := instructions.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsFeeTooLow represents a FeeTooLow error raised by the Instructions contract.
type InstructionsFeeTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeTooLow()
func InstructionsFeeTooLowErrorID() common.Hash {
	return common.HexToHash("0x732f94137528ce257864c792a6de53ff98f1203051e0a27c454063f559602458")
}

// UnpackFeeTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeTooLow()
func (instructions *Instructions) UnpackFeeTooLowError(raw []byte) (*InstructionsFeeTooLow, error) {
	out := new(InstructionsFeeTooLow)
	if err := instructions.abi.UnpackIntoInterface(out, "FeeTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidAddress represents a InvalidAddress error raised by the Instructions contract.
type InstructionsInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func InstructionsInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (instructions *Instructions) UnpackInvalidAddressError(raw []byte) (*InstructionsInvalidAddress, error) {
	out := new(InstructionsInvalidAddress)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidAvailabilityCheckStatus represents a InvalidAvailabilityCheckStatus error raised by the Instructions contract.
type InstructionsInvalidAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func InstructionsInvalidAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0xedee4ff00b842d3ef84031c60001065c540c838c5479bbf0ae56f15868427e44")
}

// UnpackInvalidAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func (instructions *Instructions) UnpackInvalidAvailabilityCheckStatusError(raw []byte) (*InstructionsInvalidAvailabilityCheckStatus, error) {
	out := new(InstructionsInvalidAvailabilityCheckStatus)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidCosigner represents a InvalidCosigner error raised by the Instructions contract.
type InstructionsInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func InstructionsInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (instructions *Instructions) UnpackInvalidCosignerError(raw []byte) (*InstructionsInvalidCosigner, error) {
	out := new(InstructionsInvalidCosigner)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidDuration represents a InvalidDuration error raised by the Instructions contract.
type InstructionsInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func InstructionsInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (instructions *Instructions) UnpackInvalidDurationError(raw []byte) (*InstructionsInvalidDuration, error) {
	out := new(InstructionsInvalidDuration)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the Instructions contract.
type InstructionsInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func InstructionsInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (instructions *Instructions) UnpackInvalidGovernanceHashError(raw []byte) (*InstructionsInvalidGovernanceHash, error) {
	out := new(InstructionsInvalidGovernanceHash)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidInitialization represents a InvalidInitialization error raised by the Instructions contract.
type InstructionsInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func InstructionsInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (instructions *Instructions) UnpackInvalidInitializationError(raw []byte) (*InstructionsInvalidInitialization, error) {
	out := new(InstructionsInvalidInitialization)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidInstructionsSender represents a InvalidInstructionsSender error raised by the Instructions contract.
type InstructionsInvalidInstructionsSender struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInstructionsSender()
func InstructionsInvalidInstructionsSenderErrorID() common.Hash {
	return common.HexToHash("0x5bed0dd65877c8b84b4527862f2be859e8495bb33a1d40cf0df0156d3973c72e")
}

// UnpackInvalidInstructionsSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInstructionsSender()
func (instructions *Instructions) UnpackInvalidInstructionsSenderError(raw []byte) (*InstructionsInvalidInstructionsSender, error) {
	out := new(InstructionsInvalidInstructionsSender)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidInstructionsSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidKeyType represents a InvalidKeyType error raised by the Instructions contract.
type InstructionsInvalidKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyType()
func InstructionsInvalidKeyTypeErrorID() common.Hash {
	return common.HexToHash("0x36dccb440beb343622d8ee0927bce3f69281feea3a5b0610f6e00e2811b557e1")
}

// UnpackInvalidKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyType()
func (instructions *Instructions) UnpackInvalidKeyTypeError(raw []byte) (*InstructionsInvalidKeyType, error) {
	out := new(InstructionsInvalidKeyType)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidNonce represents a InvalidNonce error raised by the Instructions contract.
type InstructionsInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func InstructionsInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (instructions *Instructions) UnpackInvalidNonceError(raw []byte) (*InstructionsInvalidNonce, error) {
	out := new(InstructionsInvalidNonce)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidPublicKey represents a InvalidPublicKey error raised by the Instructions contract.
type InstructionsInvalidPublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPublicKey()
func InstructionsInvalidPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa2d0fee854985cb3d3a394146b5df1c41d5b4d6fcede7fc57a3c05a49b2f3e73")
}

// UnpackInvalidPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPublicKey()
func (instructions *Instructions) UnpackInvalidPublicKeyError(raw []byte) (*InstructionsInvalidPublicKey, error) {
	out := new(InstructionsInvalidPublicKey)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidResponseData represents a InvalidResponseData error raised by the Instructions contract.
type InstructionsInvalidResponseData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseData()
func InstructionsInvalidResponseDataErrorID() common.Hash {
	return common.HexToHash("0x48566617128355a033880805d6462da9bf9c6fe4c61b922997451a71262990af")
}

// UnpackInvalidResponseDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseData()
func (instructions *Instructions) UnpackInvalidResponseDataError(raw []byte) (*InstructionsInvalidResponseData, error) {
	out := new(InstructionsInvalidResponseData)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidResponseData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidSigningAlgo represents a InvalidSigningAlgo error raised by the Instructions contract.
type InstructionsInvalidSigningAlgo struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningAlgo()
func InstructionsInvalidSigningAlgoErrorID() common.Hash {
	return common.HexToHash("0x9b7198a697b61e546ef9ee593a6f77d2bf9f1df9dcbba655ce96d7e2c75346c7")
}

// UnpackInvalidSigningAlgoError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningAlgo()
func (instructions *Instructions) UnpackInvalidSigningAlgoError(raw []byte) (*InstructionsInvalidSigningAlgo, error) {
	out := new(InstructionsInvalidSigningAlgo)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidSigningAlgo", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidThreshold represents a InvalidThreshold error raised by the Instructions contract.
type InstructionsInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func InstructionsInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (instructions *Instructions) UnpackInvalidThresholdError(raw []byte) (*InstructionsInvalidThreshold, error) {
	out := new(InstructionsInvalidThreshold)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsInvalidWalletStatus represents a InvalidWalletStatus error raised by the Instructions contract.
type InstructionsInvalidWalletStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidWalletStatus()
func InstructionsInvalidWalletStatusErrorID() common.Hash {
	return common.HexToHash("0xd700120e4896ea7835a5e9e4feb690644ced5cad1c52209afdcbe06b96d91345")
}

// UnpackInvalidWalletStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidWalletStatus()
func (instructions *Instructions) UnpackInvalidWalletStatusError(raw []byte) (*InstructionsInvalidWalletStatus, error) {
	out := new(InstructionsInvalidWalletStatus)
	if err := instructions.abi.UnpackIntoInterface(out, "InvalidWalletStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsKeyTypeNotSupported represents a KeyTypeNotSupported error raised by the Instructions contract.
type InstructionsKeyTypeNotSupported struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func InstructionsKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x85004df1f47814c2de6fbf03d3f2b1ca8c67c182d0d54278766543b6be3bc603")
}

// UnpackKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func (instructions *Instructions) UnpackKeyTypeNotSupportedError(raw []byte) (*InstructionsKeyTypeNotSupported, error) {
	out := new(InstructionsKeyTypeNotSupported)
	if err := instructions.abi.UnpackIntoInterface(out, "KeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsLengthsMismatch represents a LengthsMismatch error raised by the Instructions contract.
type InstructionsLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func InstructionsLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (instructions *Instructions) UnpackLengthsMismatchError(raw []byte) (*InstructionsLengthsMismatch, error) {
	out := new(InstructionsLengthsMismatch)
	if err := instructions.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsMessageEmpty represents a MessageEmpty error raised by the Instructions contract.
type InstructionsMessageEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MessageEmpty()
func InstructionsMessageEmptyErrorID() common.Hash {
	return common.HexToHash("0x3b06a1beedd2da56654bd6f78f50f74e9bc7469bae35c52fd3680e4d95103c79")
}

// UnpackMessageEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MessageEmpty()
func (instructions *Instructions) UnpackMessageEmptyError(raw []byte) (*InstructionsMessageEmpty, error) {
	out := new(InstructionsMessageEmpty)
	if err := instructions.abi.UnpackIntoInterface(out, "MessageEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsNoAddresses represents a NoAddresses error raised by the Instructions contract.
type InstructionsNoAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoAddresses()
func InstructionsNoAddressesErrorID() common.Hash {
	return common.HexToHash("0xb2e79f05746fa6c9f3d61f9efe85959371cde5030a6c2dd70c92b63578c5cde7")
}

// UnpackNoAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoAddresses()
func (instructions *Instructions) UnpackNoAddressesError(raw []byte) (*InstructionsNoAddresses, error) {
	out := new(InstructionsNoAddresses)
	if err := instructions.abi.UnpackIntoInterface(out, "NoAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsNoTeeMachinesSpecified represents a NoTeeMachinesSpecified error raised by the Instructions contract.
type InstructionsNoTeeMachinesSpecified struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoTeeMachinesSpecified()
func InstructionsNoTeeMachinesSpecifiedErrorID() common.Hash {
	return common.HexToHash("0xe10387d0118f08d51d80fe03b0c14ed458c46aac2ae9964e0b628978e6938db8")
}

// UnpackNoTeeMachinesSpecifiedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoTeeMachinesSpecified()
func (instructions *Instructions) UnpackNoTeeMachinesSpecifiedError(raw []byte) (*InstructionsNoTeeMachinesSpecified, error) {
	out := new(InstructionsNoTeeMachinesSpecified)
	if err := instructions.abi.UnpackIntoInterface(out, "NoTeeMachinesSpecified", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsNotInitializing represents a NotInitializing error raised by the Instructions contract.
type InstructionsNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func InstructionsNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (instructions *Instructions) UnpackNotInitializingError(raw []byte) (*InstructionsNotInitializing, error) {
	out := new(InstructionsNotInitializing)
	if err := instructions.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsNotOwnerOrPauser represents a NotOwnerOrPauser error raised by the Instructions contract.
type InstructionsNotOwnerOrPauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrPauser(address caller)
func InstructionsNotOwnerOrPauserErrorID() common.Hash {
	return common.HexToHash("0x46c0d9af6b895093e022ad2984c54e60192e6e26c5063daf232e825c6e07ff9a")
}

// UnpackNotOwnerOrPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrPauser(address caller)
func (instructions *Instructions) UnpackNotOwnerOrPauserError(raw []byte) (*InstructionsNotOwnerOrPauser, error) {
	out := new(InstructionsNotOwnerOrPauser)
	if err := instructions.abi.UnpackIntoInterface(out, "NotOwnerOrPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsNotOwnerOrUnpauser represents a NotOwnerOrUnpauser error raised by the Instructions contract.
type InstructionsNotOwnerOrUnpauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func InstructionsNotOwnerOrUnpauserErrorID() common.Hash {
	return common.HexToHash("0x81e700e8f0bd5ba4a6d6a2ce9a1ad0ac9bb1b1b1e62d921a55cbc4b53f1da649")
}

// UnpackNotOwnerOrUnpauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func (instructions *Instructions) UnpackNotOwnerOrUnpauserError(raw []byte) (*InstructionsNotOwnerOrUnpauser, error) {
	out := new(InstructionsNotOwnerOrUnpauser)
	if err := instructions.abi.UnpackIntoInterface(out, "NotOwnerOrUnpauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsOnlyExtensionOwner represents a OnlyExtensionOwner error raised by the Instructions contract.
type InstructionsOnlyExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwner()
func InstructionsOnlyExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xcd49fd1de5a297cd1ce055d5abffcf2cb25ad9492bd39a703a238af57b2a9488")
}

// UnpackOnlyExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwner()
func (instructions *Instructions) UnpackOnlyExtensionOwnerError(raw []byte) (*InstructionsOnlyExtensionOwner, error) {
	out := new(InstructionsOnlyExtensionOwner)
	if err := instructions.abi.UnpackIntoInterface(out, "OnlyExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the Instructions contract.
type InstructionsOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func InstructionsOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (instructions *Instructions) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*InstructionsOnlyExtensionOwnerOrOperator, error) {
	out := new(InstructionsOnlyExtensionOwnerOrOperator)
	if err := instructions.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsOnlyGovernance represents a OnlyGovernance error raised by the Instructions contract.
type InstructionsOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func InstructionsOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (instructions *Instructions) UnpackOnlyGovernanceError(raw []byte) (*InstructionsOnlyGovernance, error) {
	out := new(InstructionsOnlyGovernance)
	if err := instructions.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsOnlyInstructionsSender represents a OnlyInstructionsSender error raised by the Instructions contract.
type InstructionsOnlyInstructionsSender struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyInstructionsSender()
func InstructionsOnlyInstructionsSenderErrorID() common.Hash {
	return common.HexToHash("0x552fdce2b6beab0198f6649b0feafb81e257123e380e5fe1f4b63850c359905f")
}

// UnpackOnlyInstructionsSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyInstructionsSender()
func (instructions *Instructions) UnpackOnlyInstructionsSenderError(raw []byte) (*InstructionsOnlyInstructionsSender, error) {
	out := new(InstructionsOnlyInstructionsSender)
	if err := instructions.abi.UnpackIntoInterface(out, "OnlyInstructionsSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsOnlyOwner represents a OnlyOwner error raised by the Instructions contract.
type InstructionsOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func InstructionsOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (instructions *Instructions) UnpackOnlyOwnerError(raw []byte) (*InstructionsOnlyOwner, error) {
	out := new(InstructionsOnlyOwner)
	if err := instructions.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsOnlyOwnerOrBackupManager represents a OnlyOwnerOrBackupManager error raised by the Instructions contract.
type InstructionsOnlyOwnerOrBackupManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrBackupManager()
func InstructionsOnlyOwnerOrBackupManagerErrorID() common.Hash {
	return common.HexToHash("0x249a6cf6e22092f9c80798e2a0c583c0f36aea0779ef38a5b5cdb6a112a0e6df")
}

// UnpackOnlyOwnerOrBackupManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrBackupManager()
func (instructions *Instructions) UnpackOnlyOwnerOrBackupManagerError(raw []byte) (*InstructionsOnlyOwnerOrBackupManager, error) {
	out := new(InstructionsOnlyOwnerOrBackupManager)
	if err := instructions.abi.UnpackIntoInterface(out, "OnlyOwnerOrBackupManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the Instructions contract.
type InstructionsOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func InstructionsOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (instructions *Instructions) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*InstructionsOnlyProductionOrPausedStatus, error) {
	out := new(InstructionsOnlyProductionOrPausedStatus)
	if err := instructions.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsOnlyProposedOwner represents a OnlyProposedOwner error raised by the Instructions contract.
type InstructionsOnlyProposedOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProposedOwner()
func InstructionsOnlyProposedOwnerErrorID() common.Hash {
	return common.HexToHash("0x49cb15061ee558e74f66f937f740fd2a3791a89d2d139862222e5c97cabc02cd")
}

// UnpackOnlyProposedOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProposedOwner()
func (instructions *Instructions) UnpackOnlyProposedOwnerError(raw []byte) (*InstructionsOnlyProposedOwner, error) {
	out := new(InstructionsOnlyProposedOwner)
	if err := instructions.abi.UnpackIntoInterface(out, "OnlyProposedOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsOnlySystemInstructionsSender represents a OnlySystemInstructionsSender error raised by the Instructions contract.
type InstructionsOnlySystemInstructionsSender struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemInstructionsSender()
func InstructionsOnlySystemInstructionsSenderErrorID() common.Hash {
	return common.HexToHash("0x6d91a66a3b94c696672b5467f1a4a0020afd670f91bab82a178ea9b1a53daede")
}

// UnpackOnlySystemInstructionsSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemInstructionsSender()
func (instructions *Instructions) UnpackOnlySystemInstructionsSenderError(raw []byte) (*InstructionsOnlySystemInstructionsSender, error) {
	out := new(InstructionsOnlySystemInstructionsSender)
	if err := instructions.abi.UnpackIntoInterface(out, "OnlySystemInstructionsSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsOperationCommandEmpty represents a OperationCommandEmpty error raised by the Instructions contract.
type InstructionsOperationCommandEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationCommandEmpty()
func InstructionsOperationCommandEmptyErrorID() common.Hash {
	return common.HexToHash("0x038dc5a8067401ab291233391b3d11ff3ce1b21a8aebcf4297e6a9f759f65846")
}

// UnpackOperationCommandEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationCommandEmpty()
func (instructions *Instructions) UnpackOperationCommandEmptyError(raw []byte) (*InstructionsOperationCommandEmpty, error) {
	out := new(InstructionsOperationCommandEmpty)
	if err := instructions.abi.UnpackIntoInterface(out, "OperationCommandEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsOperationTypeEmpty represents a OperationTypeEmpty error raised by the Instructions contract.
type InstructionsOperationTypeEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationTypeEmpty()
func InstructionsOperationTypeEmptyErrorID() common.Hash {
	return common.HexToHash("0x06e3960d8a83ea8972d826357113f286ce60947bac56680f119fca98a2d6125b")
}

// UnpackOperationTypeEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationTypeEmpty()
func (instructions *Instructions) UnpackOperationTypeEmptyError(raw []byte) (*InstructionsOperationTypeEmpty, error) {
	out := new(InstructionsOperationTypeEmpty)
	if err := instructions.abi.UnpackIntoInterface(out, "OperationTypeEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsOwnerNotAllowed represents a OwnerNotAllowed error raised by the Instructions contract.
type InstructionsOwnerNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerNotAllowed()
func InstructionsOwnerNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x3d5c9996545c6bf4287c26c1c2d8e6b4c39b92f19ba706c3a756720bf9c4fd24")
}

// UnpackOwnerNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerNotAllowed()
func (instructions *Instructions) UnpackOwnerNotAllowedError(raw []byte) (*InstructionsOwnerNotAllowed, error) {
	out := new(InstructionsOwnerNotAllowed)
	if err := instructions.abi.UnpackIntoInterface(out, "OwnerNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsSystemInstructionsSenderAlreadyExists represents a SystemInstructionsSenderAlreadyExists error raised by the Instructions contract.
type InstructionsSystemInstructionsSenderAlreadyExists struct {
	InstructionsSender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SystemInstructionsSenderAlreadyExists(address instructionsSender)
func InstructionsSystemInstructionsSenderAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x4619b3f77d436f95c709556d25aae2911c5a2bbb4825bc9e1ed4b6ab881d9894")
}

// UnpackSystemInstructionsSenderAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SystemInstructionsSenderAlreadyExists(address instructionsSender)
func (instructions *Instructions) UnpackSystemInstructionsSenderAlreadyExistsError(raw []byte) (*InstructionsSystemInstructionsSenderAlreadyExists, error) {
	out := new(InstructionsSystemInstructionsSenderAlreadyExists)
	if err := instructions.abi.UnpackIntoInterface(out, "SystemInstructionsSenderAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsSystemInstructionsSenderNotFound represents a SystemInstructionsSenderNotFound error raised by the Instructions contract.
type InstructionsSystemInstructionsSenderNotFound struct {
	InstructionsSender common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SystemInstructionsSenderNotFound(address instructionsSender)
func InstructionsSystemInstructionsSenderNotFoundErrorID() common.Hash {
	return common.HexToHash("0xa1168ac04b16d431fe3730377abfaf1a006295294f618aa119c90476d3c68317")
}

// UnpackSystemInstructionsSenderNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SystemInstructionsSenderNotFound(address instructionsSender)
func (instructions *Instructions) UnpackSystemInstructionsSenderNotFoundError(raw []byte) (*InstructionsSystemInstructionsSenderNotFound, error) {
	out := new(InstructionsSystemInstructionsSenderNotFound)
	if err := instructions.abi.UnpackIntoInterface(out, "SystemInstructionsSenderNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsSystemOpTypeNotAllowed represents a SystemOpTypeNotAllowed error raised by the Instructions contract.
type InstructionsSystemOpTypeNotAllowed struct {
	OpType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SystemOpTypeNotAllowed(bytes32 opType)
func InstructionsSystemOpTypeNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x76751df70f5b49136a3ccee5bb02e7269b6c80583770ddca2581a20f96de6aa8")
}

// UnpackSystemOpTypeNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SystemOpTypeNotAllowed(bytes32 opType)
func (instructions *Instructions) UnpackSystemOpTypeNotAllowedError(raw []byte) (*InstructionsSystemOpTypeNotAllowed, error) {
	out := new(InstructionsSystemOpTypeNotAllowed)
	if err := instructions.abi.UnpackIntoInterface(out, "SystemOpTypeNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the Instructions contract.
type InstructionsTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func InstructionsTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (instructions *Instructions) UnpackTeeMachineNotAvailableError(raw []byte) (*InstructionsTeeMachineNotAvailable, error) {
	out := new(InstructionsTeeMachineNotAvailable)
	if err := instructions.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsTeeNotFound represents a TeeNotFound error raised by the Instructions contract.
type InstructionsTeeNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeNotFound()
func InstructionsTeeNotFoundErrorID() common.Hash {
	return common.HexToHash("0xceb05b6852ed4c7dc3a718eb0d26555398a76877a44dd41dc92c5fe7e6ef8c2d")
}

// UnpackTeeNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeNotFound()
func (instructions *Instructions) UnpackTeeNotFoundError(raw []byte) (*InstructionsTeeNotFound, error) {
	out := new(InstructionsTeeNotFound)
	if err := instructions.abi.UnpackIntoInterface(out, "TeeNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// InstructionsVersionNotSupported represents a VersionNotSupported error raised by the Instructions contract.
type InstructionsVersionNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionNotSupported()
func InstructionsVersionNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x6db0b17b06266b37e721024a8ba1221439799ebb779dfd397865e37250abc9a8")
}

// UnpackVersionNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionNotSupported()
func (instructions *Instructions) UnpackVersionNotSupportedError(raw []byte) (*InstructionsVersionNotSupported, error) {
	out := new(InstructionsVersionNotSupported)
	if err := instructions.abi.UnpackIntoInterface(out, "VersionNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}
