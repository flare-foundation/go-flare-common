// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package operationfees

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

// OperationFeesMetaData contains all meta data concerning the OperationFees contract.
var OperationFeesMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DefaultFeeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultFee\",\"type\":\"uint256\"}],\"name\":\"DefaultFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"opTypes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"opCommands\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"}],\"name\":\"OperationFeesSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"name\":\"calculateFeeByTeeIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"}],\"name\":\"calculateFeeByWalletId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"}],\"name\":\"getOperationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_defaultFee\",\"type\":\"uint256\"}],\"name\":\"setDefaultFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_opTypes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_opCommands\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"}],\"name\":\"setOperationFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "OperationFees",
}

// OperationFees is an auto generated Go binding around an Ethereum contract.
type OperationFees struct {
	abi abi.ABI
}

// NewOperationFees creates a new instance of OperationFees.
func NewOperationFees() *OperationFees {
	parsed, err := OperationFeesMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &OperationFees{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *OperationFees) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCalculateFeeByTeeIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x81174ec7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function calculateFeeByTeeIds(bytes32 _opType, bytes32 _opCommand, address[] _teeIds) view returns(uint256 _fee)
func (operationFees *OperationFees) PackCalculateFeeByTeeIds(opType [32]byte, opCommand [32]byte, teeIds []common.Address) []byte {
	enc, err := operationFees.abi.Pack("calculateFeeByTeeIds", opType, opCommand, teeIds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCalculateFeeByTeeIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x81174ec7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function calculateFeeByTeeIds(bytes32 _opType, bytes32 _opCommand, address[] _teeIds) view returns(uint256 _fee)
func (operationFees *OperationFees) TryPackCalculateFeeByTeeIds(opType [32]byte, opCommand [32]byte, teeIds []common.Address) ([]byte, error) {
	return operationFees.abi.Pack("calculateFeeByTeeIds", opType, opCommand, teeIds)
}

// UnpackCalculateFeeByTeeIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x81174ec7.
//
// Solidity: function calculateFeeByTeeIds(bytes32 _opType, bytes32 _opCommand, address[] _teeIds) view returns(uint256 _fee)
func (operationFees *OperationFees) UnpackCalculateFeeByTeeIds(data []byte) (*big.Int, error) {
	out, err := operationFees.abi.Unpack("calculateFeeByTeeIds", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCalculateFeeByWalletId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb4e28b6e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function calculateFeeByWalletId(bytes32 _walletId, bytes32 _opType, bytes32 _opCommand) view returns(uint256 _fee)
func (operationFees *OperationFees) PackCalculateFeeByWalletId(walletId [32]byte, opType [32]byte, opCommand [32]byte) []byte {
	enc, err := operationFees.abi.Pack("calculateFeeByWalletId", walletId, opType, opCommand)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCalculateFeeByWalletId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb4e28b6e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function calculateFeeByWalletId(bytes32 _walletId, bytes32 _opType, bytes32 _opCommand) view returns(uint256 _fee)
func (operationFees *OperationFees) TryPackCalculateFeeByWalletId(walletId [32]byte, opType [32]byte, opCommand [32]byte) ([]byte, error) {
	return operationFees.abi.Pack("calculateFeeByWalletId", walletId, opType, opCommand)
}

// UnpackCalculateFeeByWalletId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb4e28b6e.
//
// Solidity: function calculateFeeByWalletId(bytes32 _walletId, bytes32 _opType, bytes32 _opCommand) view returns(uint256 _fee)
func (operationFees *OperationFees) UnpackCalculateFeeByWalletId(data []byte) (*big.Int, error) {
	out, err := operationFees.abi.Unpack("calculateFeeByWalletId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetDefaultFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x21bacf28.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getDefaultFee() view returns(uint256)
func (operationFees *OperationFees) PackGetDefaultFee() []byte {
	enc, err := operationFees.abi.Pack("getDefaultFee")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetDefaultFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x21bacf28.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getDefaultFee() view returns(uint256)
func (operationFees *OperationFees) TryPackGetDefaultFee() ([]byte, error) {
	return operationFees.abi.Pack("getDefaultFee")
}

// UnpackGetDefaultFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x21bacf28.
//
// Solidity: function getDefaultFee() view returns(uint256)
func (operationFees *OperationFees) UnpackGetDefaultFee(data []byte) (*big.Int, error) {
	out, err := operationFees.abi.Unpack("getDefaultFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetOperationFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3d5e0390.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getOperationFee(bytes32 _opType, bytes32 _opCommand) view returns(uint256)
func (operationFees *OperationFees) PackGetOperationFee(opType [32]byte, opCommand [32]byte) []byte {
	enc, err := operationFees.abi.Pack("getOperationFee", opType, opCommand)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetOperationFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3d5e0390.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getOperationFee(bytes32 _opType, bytes32 _opCommand) view returns(uint256)
func (operationFees *OperationFees) TryPackGetOperationFee(opType [32]byte, opCommand [32]byte) ([]byte, error) {
	return operationFees.abi.Pack("getOperationFee", opType, opCommand)
}

// UnpackGetOperationFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3d5e0390.
//
// Solidity: function getOperationFee(bytes32 _opType, bytes32 _opCommand) view returns(uint256)
func (operationFees *OperationFees) UnpackGetOperationFee(data []byte) (*big.Int, error) {
	out, err := operationFees.abi.Unpack("getOperationFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSetDefaultFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc93a6c84.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setDefaultFee(uint256 _defaultFee) returns()
func (operationFees *OperationFees) PackSetDefaultFee(defaultFee *big.Int) []byte {
	enc, err := operationFees.abi.Pack("setDefaultFee", defaultFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetDefaultFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc93a6c84.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setDefaultFee(uint256 _defaultFee) returns()
func (operationFees *OperationFees) TryPackSetDefaultFee(defaultFee *big.Int) ([]byte, error) {
	return operationFees.abi.Pack("setDefaultFee", defaultFee)
}

// PackSetOperationFees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaa259c3e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setOperationFees(bytes32[] _opTypes, bytes32[] _opCommands, uint256[] _fees) returns()
func (operationFees *OperationFees) PackSetOperationFees(opTypes [][32]byte, opCommands [][32]byte, fees []*big.Int) []byte {
	enc, err := operationFees.abi.Pack("setOperationFees", opTypes, opCommands, fees)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetOperationFees is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaa259c3e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setOperationFees(bytes32[] _opTypes, bytes32[] _opCommands, uint256[] _fees) returns()
func (operationFees *OperationFees) TryPackSetOperationFees(opTypes [][32]byte, opCommands [][32]byte, fees []*big.Int) ([]byte, error) {
	return operationFees.abi.Pack("setOperationFees", opTypes, opCommands, fees)
}

// OperationFeesDefaultFeeSet represents a DefaultFeeSet event raised by the OperationFees contract.
type OperationFeesDefaultFeeSet struct {
	DefaultFee *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const OperationFeesDefaultFeeSetEventName = "DefaultFeeSet"

// ContractEventName returns the user-defined event name.
func (OperationFeesDefaultFeeSet) ContractEventName() string {
	return OperationFeesDefaultFeeSetEventName
}

// UnpackDefaultFeeSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultFeeSet(uint256 defaultFee)
func (operationFees *OperationFees) UnpackDefaultFeeSetEvent(log *types.Log) (*OperationFeesDefaultFeeSet, error) {
	event := "DefaultFeeSet"
	if len(log.Topics) == 0 || log.Topics[0] != operationFees.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OperationFeesDefaultFeeSet)
	if len(log.Data) > 0 {
		if err := operationFees.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range operationFees.abi.Events[event].Inputs {
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

// OperationFeesGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the OperationFees contract.
type OperationFeesGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const OperationFeesGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (OperationFeesGovernanceCallTimelocked) ContractEventName() string {
	return OperationFeesGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (operationFees *OperationFees) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*OperationFeesGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != operationFees.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OperationFeesGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := operationFees.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range operationFees.abi.Events[event].Inputs {
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

// OperationFeesInitialized represents a Initialized event raised by the OperationFees contract.
type OperationFeesInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const OperationFeesInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (OperationFeesInitialized) ContractEventName() string {
	return OperationFeesInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (operationFees *OperationFees) UnpackInitializedEvent(log *types.Log) (*OperationFeesInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != operationFees.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OperationFeesInitialized)
	if len(log.Data) > 0 {
		if err := operationFees.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range operationFees.abi.Events[event].Inputs {
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

// OperationFeesOperationFeesSet represents a OperationFeesSet event raised by the OperationFees contract.
type OperationFeesOperationFeesSet struct {
	OpTypes    [][32]byte
	OpCommands [][32]byte
	Fees       []*big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const OperationFeesOperationFeesSetEventName = "OperationFeesSet"

// ContractEventName returns the user-defined event name.
func (OperationFeesOperationFeesSet) ContractEventName() string {
	return OperationFeesOperationFeesSetEventName
}

// UnpackOperationFeesSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OperationFeesSet(bytes32[] opTypes, bytes32[] opCommands, uint256[] fees)
func (operationFees *OperationFees) UnpackOperationFeesSetEvent(log *types.Log) (*OperationFeesOperationFeesSet, error) {
	event := "OperationFeesSet"
	if len(log.Topics) == 0 || log.Topics[0] != operationFees.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(OperationFeesOperationFeesSet)
	if len(log.Data) > 0 {
		if err := operationFees.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range operationFees.abi.Events[event].Inputs {
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
func (operationFees *OperationFees) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], operationFees.abi.Errors["AddressAlreadyInSet"].ID.Bytes()[:4]) {
		return operationFees.UnpackAddressAlreadyInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["AddressNotInSet"].ID.Bytes()[:4]) {
		return operationFees.UnpackAddressNotInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["AvailabilityCheckTimestampInvalid"].ID.Bytes()[:4]) {
		return operationFees.UnpackAvailabilityCheckTimestampInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["DefaultFeeZero"].ID.Bytes()[:4]) {
		return operationFees.UnpackDefaultFeeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["DuplicatedCosigner"].ID.Bytes()[:4]) {
		return operationFees.UnpackDuplicatedCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return operationFees.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return operationFees.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["InvalidAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return operationFees.UnpackInvalidAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return operationFees.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return operationFees.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return operationFees.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return operationFees.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["InvalidKeyType"].ID.Bytes()[:4]) {
		return operationFees.UnpackInvalidKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return operationFees.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["InvalidPublicKey"].ID.Bytes()[:4]) {
		return operationFees.UnpackInvalidPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["InvalidResponseData"].ID.Bytes()[:4]) {
		return operationFees.UnpackInvalidResponseDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["InvalidSigningAlgo"].ID.Bytes()[:4]) {
		return operationFees.UnpackInvalidSigningAlgoError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return operationFees.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["InvalidWalletStatus"].ID.Bytes()[:4]) {
		return operationFees.UnpackInvalidWalletStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["KeyTypeNotSupported"].ID.Bytes()[:4]) {
		return operationFees.UnpackKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return operationFees.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["NoAddresses"].ID.Bytes()[:4]) {
		return operationFees.UnpackNoAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return operationFees.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["NotOwnerOrPauser"].ID.Bytes()[:4]) {
		return operationFees.UnpackNotOwnerOrPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["NotOwnerOrUnpauser"].ID.Bytes()[:4]) {
		return operationFees.UnpackNotOwnerOrUnpauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["OnlyExtensionOwner"].ID.Bytes()[:4]) {
		return operationFees.UnpackOnlyExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return operationFees.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return operationFees.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return operationFees.UnpackOnlyOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["OnlyOwnerOrBackupManager"].ID.Bytes()[:4]) {
		return operationFees.UnpackOnlyOwnerOrBackupManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return operationFees.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["OnlyProposedOwner"].ID.Bytes()[:4]) {
		return operationFees.UnpackOnlyProposedOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["OwnerNotAllowed"].ID.Bytes()[:4]) {
		return operationFees.UnpackOwnerNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return operationFees.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["TeeNotFound"].ID.Bytes()[:4]) {
		return operationFees.UnpackTeeNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["ThresholdNotMet"].ID.Bytes()[:4]) {
		return operationFees.UnpackThresholdNotMetError(raw[4:])
	}
	if bytes.Equal(raw[:4], operationFees.abi.Errors["VersionNotSupported"].ID.Bytes()[:4]) {
		return operationFees.UnpackVersionNotSupportedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// OperationFeesAddressAlreadyInSet represents a AddressAlreadyInSet error raised by the OperationFees contract.
type OperationFeesAddressAlreadyInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressAlreadyInSet(address addr)
func OperationFeesAddressAlreadyInSetErrorID() common.Hash {
	return common.HexToHash("0xb725c7fb1c3f0f065973d6b9eb55140f970259955af3045742cd5b37902eb920")
}

// UnpackAddressAlreadyInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressAlreadyInSet(address addr)
func (operationFees *OperationFees) UnpackAddressAlreadyInSetError(raw []byte) (*OperationFeesAddressAlreadyInSet, error) {
	out := new(OperationFeesAddressAlreadyInSet)
	if err := operationFees.abi.UnpackIntoInterface(out, "AddressAlreadyInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesAddressNotInSet represents a AddressNotInSet error raised by the OperationFees contract.
type OperationFeesAddressNotInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressNotInSet(address addr)
func OperationFeesAddressNotInSetErrorID() common.Hash {
	return common.HexToHash("0x1ff5ecf3e6d385ea01ff0323fe894ed6d25a83ed5f6ff2a2ede4c897408f6be9")
}

// UnpackAddressNotInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressNotInSet(address addr)
func (operationFees *OperationFees) UnpackAddressNotInSetError(raw []byte) (*OperationFeesAddressNotInSet, error) {
	out := new(OperationFeesAddressNotInSet)
	if err := operationFees.abi.UnpackIntoInterface(out, "AddressNotInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesAvailabilityCheckTimestampInvalid represents a AvailabilityCheckTimestampInvalid error raised by the OperationFees contract.
type OperationFeesAvailabilityCheckTimestampInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func OperationFeesAvailabilityCheckTimestampInvalidErrorID() common.Hash {
	return common.HexToHash("0x6344a5e74cff78612fa271fa907ca608af3bad10e1e8c8df302e29ebc6298e2d")
}

// UnpackAvailabilityCheckTimestampInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func (operationFees *OperationFees) UnpackAvailabilityCheckTimestampInvalidError(raw []byte) (*OperationFeesAvailabilityCheckTimestampInvalid, error) {
	out := new(OperationFeesAvailabilityCheckTimestampInvalid)
	if err := operationFees.abi.UnpackIntoInterface(out, "AvailabilityCheckTimestampInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesDefaultFeeZero represents a DefaultFeeZero error raised by the OperationFees contract.
type OperationFeesDefaultFeeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DefaultFeeZero()
func OperationFeesDefaultFeeZeroErrorID() common.Hash {
	return common.HexToHash("0x85ed2180a1896c3c873e33b45c04d2bf59b01866e984c0ebe6c53b66730eff4f")
}

// UnpackDefaultFeeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DefaultFeeZero()
func (operationFees *OperationFees) UnpackDefaultFeeZeroError(raw []byte) (*OperationFeesDefaultFeeZero, error) {
	out := new(OperationFeesDefaultFeeZero)
	if err := operationFees.abi.UnpackIntoInterface(out, "DefaultFeeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesDuplicatedCosigner represents a DuplicatedCosigner error raised by the OperationFees contract.
type OperationFeesDuplicatedCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func OperationFeesDuplicatedCosignerErrorID() common.Hash {
	return common.HexToHash("0x53a5890417576fb3043dc3ff1669b0932da2f657d8adb4c11313b12b711f5910")
}

// UnpackDuplicatedCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func (operationFees *OperationFees) UnpackDuplicatedCosignerError(raw []byte) (*OperationFeesDuplicatedCosigner, error) {
	out := new(OperationFeesDuplicatedCosigner)
	if err := operationFees.abi.UnpackIntoInterface(out, "DuplicatedCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesExtensionIdMismatch represents a ExtensionIdMismatch error raised by the OperationFees contract.
type OperationFeesExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func OperationFeesExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (operationFees *OperationFees) UnpackExtensionIdMismatchError(raw []byte) (*OperationFeesExtensionIdMismatch, error) {
	out := new(OperationFeesExtensionIdMismatch)
	if err := operationFees.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesInvalidAddress represents a InvalidAddress error raised by the OperationFees contract.
type OperationFeesInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func OperationFeesInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (operationFees *OperationFees) UnpackInvalidAddressError(raw []byte) (*OperationFeesInvalidAddress, error) {
	out := new(OperationFeesInvalidAddress)
	if err := operationFees.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesInvalidAvailabilityCheckStatus represents a InvalidAvailabilityCheckStatus error raised by the OperationFees contract.
type OperationFeesInvalidAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func OperationFeesInvalidAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0xedee4ff00b842d3ef84031c60001065c540c838c5479bbf0ae56f15868427e44")
}

// UnpackInvalidAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func (operationFees *OperationFees) UnpackInvalidAvailabilityCheckStatusError(raw []byte) (*OperationFeesInvalidAvailabilityCheckStatus, error) {
	out := new(OperationFeesInvalidAvailabilityCheckStatus)
	if err := operationFees.abi.UnpackIntoInterface(out, "InvalidAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesInvalidCosigner represents a InvalidCosigner error raised by the OperationFees contract.
type OperationFeesInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func OperationFeesInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (operationFees *OperationFees) UnpackInvalidCosignerError(raw []byte) (*OperationFeesInvalidCosigner, error) {
	out := new(OperationFeesInvalidCosigner)
	if err := operationFees.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesInvalidDuration represents a InvalidDuration error raised by the OperationFees contract.
type OperationFeesInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func OperationFeesInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (operationFees *OperationFees) UnpackInvalidDurationError(raw []byte) (*OperationFeesInvalidDuration, error) {
	out := new(OperationFeesInvalidDuration)
	if err := operationFees.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the OperationFees contract.
type OperationFeesInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func OperationFeesInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (operationFees *OperationFees) UnpackInvalidGovernanceHashError(raw []byte) (*OperationFeesInvalidGovernanceHash, error) {
	out := new(OperationFeesInvalidGovernanceHash)
	if err := operationFees.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesInvalidInitialization represents a InvalidInitialization error raised by the OperationFees contract.
type OperationFeesInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func OperationFeesInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (operationFees *OperationFees) UnpackInvalidInitializationError(raw []byte) (*OperationFeesInvalidInitialization, error) {
	out := new(OperationFeesInvalidInitialization)
	if err := operationFees.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesInvalidKeyType represents a InvalidKeyType error raised by the OperationFees contract.
type OperationFeesInvalidKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyType()
func OperationFeesInvalidKeyTypeErrorID() common.Hash {
	return common.HexToHash("0x36dccb440beb343622d8ee0927bce3f69281feea3a5b0610f6e00e2811b557e1")
}

// UnpackInvalidKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyType()
func (operationFees *OperationFees) UnpackInvalidKeyTypeError(raw []byte) (*OperationFeesInvalidKeyType, error) {
	out := new(OperationFeesInvalidKeyType)
	if err := operationFees.abi.UnpackIntoInterface(out, "InvalidKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesInvalidNonce represents a InvalidNonce error raised by the OperationFees contract.
type OperationFeesInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func OperationFeesInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (operationFees *OperationFees) UnpackInvalidNonceError(raw []byte) (*OperationFeesInvalidNonce, error) {
	out := new(OperationFeesInvalidNonce)
	if err := operationFees.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesInvalidPublicKey represents a InvalidPublicKey error raised by the OperationFees contract.
type OperationFeesInvalidPublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPublicKey()
func OperationFeesInvalidPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa2d0fee854985cb3d3a394146b5df1c41d5b4d6fcede7fc57a3c05a49b2f3e73")
}

// UnpackInvalidPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPublicKey()
func (operationFees *OperationFees) UnpackInvalidPublicKeyError(raw []byte) (*OperationFeesInvalidPublicKey, error) {
	out := new(OperationFeesInvalidPublicKey)
	if err := operationFees.abi.UnpackIntoInterface(out, "InvalidPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesInvalidResponseData represents a InvalidResponseData error raised by the OperationFees contract.
type OperationFeesInvalidResponseData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseData()
func OperationFeesInvalidResponseDataErrorID() common.Hash {
	return common.HexToHash("0x48566617128355a033880805d6462da9bf9c6fe4c61b922997451a71262990af")
}

// UnpackInvalidResponseDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseData()
func (operationFees *OperationFees) UnpackInvalidResponseDataError(raw []byte) (*OperationFeesInvalidResponseData, error) {
	out := new(OperationFeesInvalidResponseData)
	if err := operationFees.abi.UnpackIntoInterface(out, "InvalidResponseData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesInvalidSigningAlgo represents a InvalidSigningAlgo error raised by the OperationFees contract.
type OperationFeesInvalidSigningAlgo struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningAlgo()
func OperationFeesInvalidSigningAlgoErrorID() common.Hash {
	return common.HexToHash("0x9b7198a697b61e546ef9ee593a6f77d2bf9f1df9dcbba655ce96d7e2c75346c7")
}

// UnpackInvalidSigningAlgoError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningAlgo()
func (operationFees *OperationFees) UnpackInvalidSigningAlgoError(raw []byte) (*OperationFeesInvalidSigningAlgo, error) {
	out := new(OperationFeesInvalidSigningAlgo)
	if err := operationFees.abi.UnpackIntoInterface(out, "InvalidSigningAlgo", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesInvalidThreshold represents a InvalidThreshold error raised by the OperationFees contract.
type OperationFeesInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func OperationFeesInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (operationFees *OperationFees) UnpackInvalidThresholdError(raw []byte) (*OperationFeesInvalidThreshold, error) {
	out := new(OperationFeesInvalidThreshold)
	if err := operationFees.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesInvalidWalletStatus represents a InvalidWalletStatus error raised by the OperationFees contract.
type OperationFeesInvalidWalletStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidWalletStatus()
func OperationFeesInvalidWalletStatusErrorID() common.Hash {
	return common.HexToHash("0xd700120e4896ea7835a5e9e4feb690644ced5cad1c52209afdcbe06b96d91345")
}

// UnpackInvalidWalletStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidWalletStatus()
func (operationFees *OperationFees) UnpackInvalidWalletStatusError(raw []byte) (*OperationFeesInvalidWalletStatus, error) {
	out := new(OperationFeesInvalidWalletStatus)
	if err := operationFees.abi.UnpackIntoInterface(out, "InvalidWalletStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesKeyTypeNotSupported represents a KeyTypeNotSupported error raised by the OperationFees contract.
type OperationFeesKeyTypeNotSupported struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func OperationFeesKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x85004df1f47814c2de6fbf03d3f2b1ca8c67c182d0d54278766543b6be3bc603")
}

// UnpackKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func (operationFees *OperationFees) UnpackKeyTypeNotSupportedError(raw []byte) (*OperationFeesKeyTypeNotSupported, error) {
	out := new(OperationFeesKeyTypeNotSupported)
	if err := operationFees.abi.UnpackIntoInterface(out, "KeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesLengthsMismatch represents a LengthsMismatch error raised by the OperationFees contract.
type OperationFeesLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func OperationFeesLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (operationFees *OperationFees) UnpackLengthsMismatchError(raw []byte) (*OperationFeesLengthsMismatch, error) {
	out := new(OperationFeesLengthsMismatch)
	if err := operationFees.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesNoAddresses represents a NoAddresses error raised by the OperationFees contract.
type OperationFeesNoAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoAddresses()
func OperationFeesNoAddressesErrorID() common.Hash {
	return common.HexToHash("0xb2e79f05746fa6c9f3d61f9efe85959371cde5030a6c2dd70c92b63578c5cde7")
}

// UnpackNoAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoAddresses()
func (operationFees *OperationFees) UnpackNoAddressesError(raw []byte) (*OperationFeesNoAddresses, error) {
	out := new(OperationFeesNoAddresses)
	if err := operationFees.abi.UnpackIntoInterface(out, "NoAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesNotInitializing represents a NotInitializing error raised by the OperationFees contract.
type OperationFeesNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func OperationFeesNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (operationFees *OperationFees) UnpackNotInitializingError(raw []byte) (*OperationFeesNotInitializing, error) {
	out := new(OperationFeesNotInitializing)
	if err := operationFees.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesNotOwnerOrPauser represents a NotOwnerOrPauser error raised by the OperationFees contract.
type OperationFeesNotOwnerOrPauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrPauser(address caller)
func OperationFeesNotOwnerOrPauserErrorID() common.Hash {
	return common.HexToHash("0x46c0d9af6b895093e022ad2984c54e60192e6e26c5063daf232e825c6e07ff9a")
}

// UnpackNotOwnerOrPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrPauser(address caller)
func (operationFees *OperationFees) UnpackNotOwnerOrPauserError(raw []byte) (*OperationFeesNotOwnerOrPauser, error) {
	out := new(OperationFeesNotOwnerOrPauser)
	if err := operationFees.abi.UnpackIntoInterface(out, "NotOwnerOrPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesNotOwnerOrUnpauser represents a NotOwnerOrUnpauser error raised by the OperationFees contract.
type OperationFeesNotOwnerOrUnpauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func OperationFeesNotOwnerOrUnpauserErrorID() common.Hash {
	return common.HexToHash("0x81e700e8f0bd5ba4a6d6a2ce9a1ad0ac9bb1b1b1e62d921a55cbc4b53f1da649")
}

// UnpackNotOwnerOrUnpauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func (operationFees *OperationFees) UnpackNotOwnerOrUnpauserError(raw []byte) (*OperationFeesNotOwnerOrUnpauser, error) {
	out := new(OperationFeesNotOwnerOrUnpauser)
	if err := operationFees.abi.UnpackIntoInterface(out, "NotOwnerOrUnpauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesOnlyExtensionOwner represents a OnlyExtensionOwner error raised by the OperationFees contract.
type OperationFeesOnlyExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwner()
func OperationFeesOnlyExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xcd49fd1de5a297cd1ce055d5abffcf2cb25ad9492bd39a703a238af57b2a9488")
}

// UnpackOnlyExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwner()
func (operationFees *OperationFees) UnpackOnlyExtensionOwnerError(raw []byte) (*OperationFeesOnlyExtensionOwner, error) {
	out := new(OperationFeesOnlyExtensionOwner)
	if err := operationFees.abi.UnpackIntoInterface(out, "OnlyExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the OperationFees contract.
type OperationFeesOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func OperationFeesOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (operationFees *OperationFees) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*OperationFeesOnlyExtensionOwnerOrOperator, error) {
	out := new(OperationFeesOnlyExtensionOwnerOrOperator)
	if err := operationFees.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesOnlyGovernance represents a OnlyGovernance error raised by the OperationFees contract.
type OperationFeesOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func OperationFeesOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (operationFees *OperationFees) UnpackOnlyGovernanceError(raw []byte) (*OperationFeesOnlyGovernance, error) {
	out := new(OperationFeesOnlyGovernance)
	if err := operationFees.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesOnlyOwner represents a OnlyOwner error raised by the OperationFees contract.
type OperationFeesOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func OperationFeesOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (operationFees *OperationFees) UnpackOnlyOwnerError(raw []byte) (*OperationFeesOnlyOwner, error) {
	out := new(OperationFeesOnlyOwner)
	if err := operationFees.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesOnlyOwnerOrBackupManager represents a OnlyOwnerOrBackupManager error raised by the OperationFees contract.
type OperationFeesOnlyOwnerOrBackupManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrBackupManager()
func OperationFeesOnlyOwnerOrBackupManagerErrorID() common.Hash {
	return common.HexToHash("0x249a6cf6e22092f9c80798e2a0c583c0f36aea0779ef38a5b5cdb6a112a0e6df")
}

// UnpackOnlyOwnerOrBackupManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrBackupManager()
func (operationFees *OperationFees) UnpackOnlyOwnerOrBackupManagerError(raw []byte) (*OperationFeesOnlyOwnerOrBackupManager, error) {
	out := new(OperationFeesOnlyOwnerOrBackupManager)
	if err := operationFees.abi.UnpackIntoInterface(out, "OnlyOwnerOrBackupManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the OperationFees contract.
type OperationFeesOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func OperationFeesOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (operationFees *OperationFees) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*OperationFeesOnlyProductionOrPausedStatus, error) {
	out := new(OperationFeesOnlyProductionOrPausedStatus)
	if err := operationFees.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesOnlyProposedOwner represents a OnlyProposedOwner error raised by the OperationFees contract.
type OperationFeesOnlyProposedOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProposedOwner()
func OperationFeesOnlyProposedOwnerErrorID() common.Hash {
	return common.HexToHash("0x49cb15061ee558e74f66f937f740fd2a3791a89d2d139862222e5c97cabc02cd")
}

// UnpackOnlyProposedOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProposedOwner()
func (operationFees *OperationFees) UnpackOnlyProposedOwnerError(raw []byte) (*OperationFeesOnlyProposedOwner, error) {
	out := new(OperationFeesOnlyProposedOwner)
	if err := operationFees.abi.UnpackIntoInterface(out, "OnlyProposedOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesOwnerNotAllowed represents a OwnerNotAllowed error raised by the OperationFees contract.
type OperationFeesOwnerNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerNotAllowed()
func OperationFeesOwnerNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x3d5c9996545c6bf4287c26c1c2d8e6b4c39b92f19ba706c3a756720bf9c4fd24")
}

// UnpackOwnerNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerNotAllowed()
func (operationFees *OperationFees) UnpackOwnerNotAllowedError(raw []byte) (*OperationFeesOwnerNotAllowed, error) {
	out := new(OperationFeesOwnerNotAllowed)
	if err := operationFees.abi.UnpackIntoInterface(out, "OwnerNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the OperationFees contract.
type OperationFeesTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func OperationFeesTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (operationFees *OperationFees) UnpackTeeMachineNotAvailableError(raw []byte) (*OperationFeesTeeMachineNotAvailable, error) {
	out := new(OperationFeesTeeMachineNotAvailable)
	if err := operationFees.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesTeeNotFound represents a TeeNotFound error raised by the OperationFees contract.
type OperationFeesTeeNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeNotFound()
func OperationFeesTeeNotFoundErrorID() common.Hash {
	return common.HexToHash("0xceb05b6852ed4c7dc3a718eb0d26555398a76877a44dd41dc92c5fe7e6ef8c2d")
}

// UnpackTeeNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeNotFound()
func (operationFees *OperationFees) UnpackTeeNotFoundError(raw []byte) (*OperationFeesTeeNotFound, error) {
	out := new(OperationFeesTeeNotFound)
	if err := operationFees.abi.UnpackIntoInterface(out, "TeeNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesThresholdNotMet represents a ThresholdNotMet error raised by the OperationFees contract.
type OperationFeesThresholdNotMet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ThresholdNotMet()
func OperationFeesThresholdNotMetErrorID() common.Hash {
	return common.HexToHash("0x59fa4a935269513c1e751ea6c6fe2e0a1920bf90b5e062c629e36320baf7378d")
}

// UnpackThresholdNotMetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ThresholdNotMet()
func (operationFees *OperationFees) UnpackThresholdNotMetError(raw []byte) (*OperationFeesThresholdNotMet, error) {
	out := new(OperationFeesThresholdNotMet)
	if err := operationFees.abi.UnpackIntoInterface(out, "ThresholdNotMet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// OperationFeesVersionNotSupported represents a VersionNotSupported error raised by the OperationFees contract.
type OperationFeesVersionNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionNotSupported()
func OperationFeesVersionNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x6db0b17b06266b37e721024a8ba1221439799ebb779dfd397865e37250abc9a8")
}

// UnpackVersionNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionNotSupported()
func (operationFees *OperationFees) UnpackVersionNotSupportedError(raw []byte) (*OperationFeesVersionNotSupported, error) {
	out := new(OperationFeesVersionNotSupported)
	if err := operationFees.abi.UnpackIntoInterface(out, "VersionNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}
