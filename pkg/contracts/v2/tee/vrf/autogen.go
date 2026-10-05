// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrf

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

// VRFMetaData contains all meta data concerning the VRF contract.
var VRFMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"EmergencyPauseActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeesForKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonceEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuthorizationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizationAddress\",\"type\":\"address\"}],\"name\":\"VrfAuthorizationAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"name\":\"VrfRequested\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getVrfAuthorizationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_nonce\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestVrf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"name\":\"setVrfAuthorizationAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "VRF",
}

// VRF is an auto generated Go binding around an Ethereum contract.
type VRF struct {
	abi abi.ABI
}

// NewVRF creates a new instance of VRF.
func NewVRF() *VRF {
	parsed, err := VRFMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &VRF{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *VRF) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetVrfAuthorizationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x918945a0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVrfAuthorizationAddress(bytes32 _walletId) view returns(address)
func (vRF *VRF) PackGetVrfAuthorizationAddress(walletId [32]byte) []byte {
	enc, err := vRF.abi.Pack("getVrfAuthorizationAddress", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVrfAuthorizationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x918945a0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVrfAuthorizationAddress(bytes32 _walletId) view returns(address)
func (vRF *VRF) TryPackGetVrfAuthorizationAddress(walletId [32]byte) ([]byte, error) {
	return vRF.abi.Pack("getVrfAuthorizationAddress", walletId)
}

// UnpackGetVrfAuthorizationAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x918945a0.
//
// Solidity: function getVrfAuthorizationAddress(bytes32 _walletId) view returns(address)
func (vRF *VRF) UnpackGetVrfAuthorizationAddress(data []byte) (common.Address, error) {
	out, err := vRF.abi.Unpack("getVrfAuthorizationAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRequestVrf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd7896b11.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function requestVrf(bytes32 _walletId, uint64 _keyId, bytes _nonce, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (vRF *VRF) PackRequestVrf(walletId [32]byte, keyId uint64, nonce []byte, claimBackAddress common.Address) []byte {
	enc, err := vRF.abi.Pack("requestVrf", walletId, keyId, nonce, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRequestVrf is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd7896b11.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function requestVrf(bytes32 _walletId, uint64 _keyId, bytes _nonce, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (vRF *VRF) TryPackRequestVrf(walletId [32]byte, keyId uint64, nonce []byte, claimBackAddress common.Address) ([]byte, error) {
	return vRF.abi.Pack("requestVrf", walletId, keyId, nonce, claimBackAddress)
}

// UnpackRequestVrf is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd7896b11.
//
// Solidity: function requestVrf(bytes32 _walletId, uint64 _keyId, bytes _nonce, address _claimBackAddress) payable returns(bytes32 _instructionId)
func (vRF *VRF) UnpackRequestVrf(data []byte) ([32]byte, error) {
	out, err := vRF.abi.Unpack("requestVrf", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackSetVrfAuthorizationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x76b87bcf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setVrfAuthorizationAddress(bytes32 _walletId, address _authorizationAddress) returns()
func (vRF *VRF) PackSetVrfAuthorizationAddress(walletId [32]byte, authorizationAddress common.Address) []byte {
	enc, err := vRF.abi.Pack("setVrfAuthorizationAddress", walletId, authorizationAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetVrfAuthorizationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x76b87bcf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setVrfAuthorizationAddress(bytes32 _walletId, address _authorizationAddress) returns()
func (vRF *VRF) TryPackSetVrfAuthorizationAddress(walletId [32]byte, authorizationAddress common.Address) ([]byte, error) {
	return vRF.abi.Pack("setVrfAuthorizationAddress", walletId, authorizationAddress)
}

// VRFTeeInstructionsSent represents a TeeInstructionsSent event raised by the VRF contract.
type VRFTeeInstructionsSent struct {
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

const VRFTeeInstructionsSentEventName = "TeeInstructionsSent"

// ContractEventName returns the user-defined event name.
func (VRFTeeInstructionsSent) ContractEventName() string {
	return VRFTeeInstructionsSentEventName
}

// UnpackTeeInstructionsSentEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (vRF *VRF) UnpackTeeInstructionsSentEvent(log *types.Log) (*VRFTeeInstructionsSent, error) {
	event := "TeeInstructionsSent"
	if len(log.Topics) == 0 || log.Topics[0] != vRF.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VRFTeeInstructionsSent)
	if len(log.Data) > 0 {
		if err := vRF.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vRF.abi.Events[event].Inputs {
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

// VRFVrfAuthorizationAddressSet represents a VrfAuthorizationAddressSet event raised by the VRF contract.
type VRFVrfAuthorizationAddressSet struct {
	WalletId             [32]byte
	AuthorizationAddress common.Address
	Raw                  *types.Log // Blockchain specific contextual infos
}

const VRFVrfAuthorizationAddressSetEventName = "VrfAuthorizationAddressSet"

// ContractEventName returns the user-defined event name.
func (VRFVrfAuthorizationAddressSet) ContractEventName() string {
	return VRFVrfAuthorizationAddressSetEventName
}

// UnpackVrfAuthorizationAddressSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VrfAuthorizationAddressSet(bytes32 indexed walletId, address authorizationAddress)
func (vRF *VRF) UnpackVrfAuthorizationAddressSetEvent(log *types.Log) (*VRFVrfAuthorizationAddressSet, error) {
	event := "VrfAuthorizationAddressSet"
	if len(log.Topics) == 0 || log.Topics[0] != vRF.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VRFVrfAuthorizationAddressSet)
	if len(log.Data) > 0 {
		if err := vRF.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vRF.abi.Events[event].Inputs {
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

// VRFVrfRequested represents a VrfRequested event raised by the VRF contract.
type VRFVrfRequested struct {
	WalletId      [32]byte
	KeyId         uint64
	InstructionId [32]byte
	Raw           *types.Log // Blockchain specific contextual infos
}

const VRFVrfRequestedEventName = "VrfRequested"

// ContractEventName returns the user-defined event name.
func (VRFVrfRequested) ContractEventName() string {
	return VRFVrfRequestedEventName
}

// UnpackVrfRequestedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VrfRequested(bytes32 indexed walletId, uint64 indexed keyId, bytes32 indexed instructionId)
func (vRF *VRF) UnpackVrfRequestedEvent(log *types.Log) (*VRFVrfRequested, error) {
	event := "VrfRequested"
	if len(log.Topics) == 0 || log.Topics[0] != vRF.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(VRFVrfRequested)
	if len(log.Data) > 0 {
		if err := vRF.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range vRF.abi.Events[event].Inputs {
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
func (vRF *VRF) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], vRF.abi.Errors["CosignersThresholdTooHigh"].ID.Bytes()[:4]) {
		return vRF.UnpackCosignersThresholdTooHighError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["EmergencyPauseActive"].ID.Bytes()[:4]) {
		return vRF.UnpackEmergencyPauseActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return vRF.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["FeeTooLow"].ID.Bytes()[:4]) {
		return vRF.UnpackFeeTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["MessageEmpty"].ID.Bytes()[:4]) {
		return vRF.UnpackMessageEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["NoTeeMachinesSpecified"].ID.Bytes()[:4]) {
		return vRF.UnpackNoTeeMachinesSpecifiedError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["NoTeesForKey"].ID.Bytes()[:4]) {
		return vRF.UnpackNoTeesForKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["NonceEmpty"].ID.Bytes()[:4]) {
		return vRF.UnpackNonceEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["OnlyAuthorizationAddress"].ID.Bytes()[:4]) {
		return vRF.UnpackOnlyAuthorizationAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["OnlyWalletOwner"].ID.Bytes()[:4]) {
		return vRF.UnpackOnlyWalletOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["OperationCommandEmpty"].ID.Bytes()[:4]) {
		return vRF.UnpackOperationCommandEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["OperationTypeEmpty"].ID.Bytes()[:4]) {
		return vRF.UnpackOperationTypeEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return vRF.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["TeeNotFound"].ID.Bytes()[:4]) {
		return vRF.UnpackTeeNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], vRF.abi.Errors["WalletNotInProduction"].ID.Bytes()[:4]) {
		return vRF.UnpackWalletNotInProductionError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// VRFCosignersThresholdTooHigh represents a CosignersThresholdTooHigh error raised by the VRF contract.
type VRFCosignersThresholdTooHigh struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CosignersThresholdTooHigh()
func VRFCosignersThresholdTooHighErrorID() common.Hash {
	return common.HexToHash("0x3d7053512cbb0b649dba01ac4e3591f104fd7d0957128a39a3c552c77cd5014d")
}

// UnpackCosignersThresholdTooHighError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CosignersThresholdTooHigh()
func (vRF *VRF) UnpackCosignersThresholdTooHighError(raw []byte) (*VRFCosignersThresholdTooHigh, error) {
	out := new(VRFCosignersThresholdTooHigh)
	if err := vRF.abi.UnpackIntoInterface(out, "CosignersThresholdTooHigh", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFEmergencyPauseActive represents a EmergencyPauseActive error raised by the VRF contract.
type VRFEmergencyPauseActive struct {
	ExtensionId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func VRFEmergencyPauseActiveErrorID() common.Hash {
	return common.HexToHash("0x480cf2163927c55f46b05a5cac80ac3781076b0e50b1ed6b07995a800937d4b6")
}

// UnpackEmergencyPauseActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func (vRF *VRF) UnpackEmergencyPauseActiveError(raw []byte) (*VRFEmergencyPauseActive, error) {
	out := new(VRFEmergencyPauseActive)
	if err := vRF.abi.UnpackIntoInterface(out, "EmergencyPauseActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFExtensionIdMismatch represents a ExtensionIdMismatch error raised by the VRF contract.
type VRFExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func VRFExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (vRF *VRF) UnpackExtensionIdMismatchError(raw []byte) (*VRFExtensionIdMismatch, error) {
	out := new(VRFExtensionIdMismatch)
	if err := vRF.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFFeeTooLow represents a FeeTooLow error raised by the VRF contract.
type VRFFeeTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeTooLow()
func VRFFeeTooLowErrorID() common.Hash {
	return common.HexToHash("0x732f94137528ce257864c792a6de53ff98f1203051e0a27c454063f559602458")
}

// UnpackFeeTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeTooLow()
func (vRF *VRF) UnpackFeeTooLowError(raw []byte) (*VRFFeeTooLow, error) {
	out := new(VRFFeeTooLow)
	if err := vRF.abi.UnpackIntoInterface(out, "FeeTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFMessageEmpty represents a MessageEmpty error raised by the VRF contract.
type VRFMessageEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MessageEmpty()
func VRFMessageEmptyErrorID() common.Hash {
	return common.HexToHash("0x3b06a1beedd2da56654bd6f78f50f74e9bc7469bae35c52fd3680e4d95103c79")
}

// UnpackMessageEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MessageEmpty()
func (vRF *VRF) UnpackMessageEmptyError(raw []byte) (*VRFMessageEmpty, error) {
	out := new(VRFMessageEmpty)
	if err := vRF.abi.UnpackIntoInterface(out, "MessageEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFNoTeeMachinesSpecified represents a NoTeeMachinesSpecified error raised by the VRF contract.
type VRFNoTeeMachinesSpecified struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoTeeMachinesSpecified()
func VRFNoTeeMachinesSpecifiedErrorID() common.Hash {
	return common.HexToHash("0xe10387d0118f08d51d80fe03b0c14ed458c46aac2ae9964e0b628978e6938db8")
}

// UnpackNoTeeMachinesSpecifiedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoTeeMachinesSpecified()
func (vRF *VRF) UnpackNoTeeMachinesSpecifiedError(raw []byte) (*VRFNoTeeMachinesSpecified, error) {
	out := new(VRFNoTeeMachinesSpecified)
	if err := vRF.abi.UnpackIntoInterface(out, "NoTeeMachinesSpecified", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFNoTeesForKey represents a NoTeesForKey error raised by the VRF contract.
type VRFNoTeesForKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoTeesForKey()
func VRFNoTeesForKeyErrorID() common.Hash {
	return common.HexToHash("0x3b6e5cf253077822a16fcbafda9348dbce3ae4a87e4902694c9153d203d81d34")
}

// UnpackNoTeesForKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoTeesForKey()
func (vRF *VRF) UnpackNoTeesForKeyError(raw []byte) (*VRFNoTeesForKey, error) {
	out := new(VRFNoTeesForKey)
	if err := vRF.abi.UnpackIntoInterface(out, "NoTeesForKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFNonceEmpty represents a NonceEmpty error raised by the VRF contract.
type VRFNonceEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NonceEmpty()
func VRFNonceEmptyErrorID() common.Hash {
	return common.HexToHash("0x03a8c4f7fdae0fe041c63a7cefe960a84663ed9f64e0e56ddf9cd15f02fb6317")
}

// UnpackNonceEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NonceEmpty()
func (vRF *VRF) UnpackNonceEmptyError(raw []byte) (*VRFNonceEmpty, error) {
	out := new(VRFNonceEmpty)
	if err := vRF.abi.UnpackIntoInterface(out, "NonceEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFOnlyAuthorizationAddress represents a OnlyAuthorizationAddress error raised by the VRF contract.
type VRFOnlyAuthorizationAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyAuthorizationAddress()
func VRFOnlyAuthorizationAddressErrorID() common.Hash {
	return common.HexToHash("0xdc65e71c09e77be7539dadee1dc86849e01c3e995f0c8d49e625476c2a1b1ed5")
}

// UnpackOnlyAuthorizationAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyAuthorizationAddress()
func (vRF *VRF) UnpackOnlyAuthorizationAddressError(raw []byte) (*VRFOnlyAuthorizationAddress, error) {
	out := new(VRFOnlyAuthorizationAddress)
	if err := vRF.abi.UnpackIntoInterface(out, "OnlyAuthorizationAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFOnlyWalletOwner represents a OnlyWalletOwner error raised by the VRF contract.
type VRFOnlyWalletOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyWalletOwner()
func VRFOnlyWalletOwnerErrorID() common.Hash {
	return common.HexToHash("0x3132d62dfddd1d8b6a9dd805001453bbb6a84b746ebe51fd650df1a6704a8143")
}

// UnpackOnlyWalletOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyWalletOwner()
func (vRF *VRF) UnpackOnlyWalletOwnerError(raw []byte) (*VRFOnlyWalletOwner, error) {
	out := new(VRFOnlyWalletOwner)
	if err := vRF.abi.UnpackIntoInterface(out, "OnlyWalletOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFOperationCommandEmpty represents a OperationCommandEmpty error raised by the VRF contract.
type VRFOperationCommandEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationCommandEmpty()
func VRFOperationCommandEmptyErrorID() common.Hash {
	return common.HexToHash("0x038dc5a8067401ab291233391b3d11ff3ce1b21a8aebcf4297e6a9f759f65846")
}

// UnpackOperationCommandEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationCommandEmpty()
func (vRF *VRF) UnpackOperationCommandEmptyError(raw []byte) (*VRFOperationCommandEmpty, error) {
	out := new(VRFOperationCommandEmpty)
	if err := vRF.abi.UnpackIntoInterface(out, "OperationCommandEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFOperationTypeEmpty represents a OperationTypeEmpty error raised by the VRF contract.
type VRFOperationTypeEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationTypeEmpty()
func VRFOperationTypeEmptyErrorID() common.Hash {
	return common.HexToHash("0x06e3960d8a83ea8972d826357113f286ce60947bac56680f119fca98a2d6125b")
}

// UnpackOperationTypeEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationTypeEmpty()
func (vRF *VRF) UnpackOperationTypeEmptyError(raw []byte) (*VRFOperationTypeEmpty, error) {
	out := new(VRFOperationTypeEmpty)
	if err := vRF.abi.UnpackIntoInterface(out, "OperationTypeEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the VRF contract.
type VRFTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func VRFTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (vRF *VRF) UnpackTeeMachineNotAvailableError(raw []byte) (*VRFTeeMachineNotAvailable, error) {
	out := new(VRFTeeMachineNotAvailable)
	if err := vRF.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFTeeNotFound represents a TeeNotFound error raised by the VRF contract.
type VRFTeeNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeNotFound()
func VRFTeeNotFoundErrorID() common.Hash {
	return common.HexToHash("0xceb05b6852ed4c7dc3a718eb0d26555398a76877a44dd41dc92c5fe7e6ef8c2d")
}

// UnpackTeeNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeNotFound()
func (vRF *VRF) UnpackTeeNotFoundError(raw []byte) (*VRFTeeNotFound, error) {
	out := new(VRFTeeNotFound)
	if err := vRF.abi.UnpackIntoInterface(out, "TeeNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// VRFWalletNotInProduction represents a WalletNotInProduction error raised by the VRF contract.
type VRFWalletNotInProduction struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WalletNotInProduction()
func VRFWalletNotInProductionErrorID() common.Hash {
	return common.HexToHash("0x2b7f97d6ba9d96a13708bf15f7c7ed50584fa90c4e3f979e82b455f65581f045")
}

// UnpackWalletNotInProductionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WalletNotInProduction()
func (vRF *VRF) UnpackWalletNotInProductionError(raw []byte) (*VRFWalletNotInProduction, error) {
	out := new(VRFWalletNotInProduction)
	if err := vRF.abi.UnpackIntoInterface(out, "WalletNotInProduction", raw); err != nil {
		return nil, err
	}
	return out, nil
}
