// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flareteemanagerinit

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

// FlareTeeManagerInitMetaData contains all meta data concerning the FlareTeeManagerInit contract.
var FlareTeeManagerInitMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DefaultFeeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"graceSeconds\",\"type\":\"uint256\"}],\"name\":\"GracePeriodTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"graceSeconds\",\"type\":\"uint256\"}],\"name\":\"GracePeriodTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AllExtensionOwnersAllowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AllExtensionOwnersDisallowed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultFee\",\"type\":\"uint256\"}],\"name\":\"DefaultFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"graceSeconds\",\"type\":\"uint256\"}],\"name\":\"EmergencyUnpauseGracePeriodSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"availabilityCheckValidityDurationSeconds\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"signingPolicyValidityDurationInRewardEpochs\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"challengeValidityDurationSeconds\",\"type\":\"uint64\"}],\"name\":\"SettingsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_availabilityCheckValidityDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicyValidityDurationInRewardEpochs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_challengeValidityDurationSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_defaultFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_publicExtensionCreationEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_emergencyUnpauseGracePeriodSeconds\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "FlareTeeManagerInit",
}

// FlareTeeManagerInit is an auto generated Go binding around an Ethereum contract.
type FlareTeeManagerInit struct {
	abi abi.ABI
}

// NewFlareTeeManagerInit creates a new instance of FlareTeeManagerInit.
func NewFlareTeeManagerInit() *FlareTeeManagerInit {
	parsed, err := FlareTeeManagerInitMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &FlareTeeManagerInit{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *FlareTeeManagerInit) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (flareTeeManagerInit *FlareTeeManagerInit) PackGetAddressUpdater() []byte {
	enc, err := flareTeeManagerInit.abi.Pack("getAddressUpdater")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (flareTeeManagerInit *FlareTeeManagerInit) TryPackGetAddressUpdater() ([]byte, error) {
	return flareTeeManagerInit.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := flareTeeManagerInit.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84e65dfc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function init(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _availabilityCheckValidityDurationSeconds, uint64 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds, uint256 _defaultFee, bool _publicExtensionCreationEnabled, uint256 _emergencyUnpauseGracePeriodSeconds) returns()
func (flareTeeManagerInit *FlareTeeManagerInit) PackInit(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address, availabilityCheckValidityDurationSeconds uint64, signingPolicyValidityDurationInRewardEpochs uint64, challengeValidityDurationSeconds uint64, defaultFee *big.Int, publicExtensionCreationEnabled bool, emergencyUnpauseGracePeriodSeconds *big.Int) []byte {
	enc, err := flareTeeManagerInit.abi.Pack("init", governanceSettings, initialGovernance, addressUpdater, availabilityCheckValidityDurationSeconds, signingPolicyValidityDurationInRewardEpochs, challengeValidityDurationSeconds, defaultFee, publicExtensionCreationEnabled, emergencyUnpauseGracePeriodSeconds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x84e65dfc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function init(address _governanceSettings, address _initialGovernance, address _addressUpdater, uint64 _availabilityCheckValidityDurationSeconds, uint64 _signingPolicyValidityDurationInRewardEpochs, uint64 _challengeValidityDurationSeconds, uint256 _defaultFee, bool _publicExtensionCreationEnabled, uint256 _emergencyUnpauseGracePeriodSeconds) returns()
func (flareTeeManagerInit *FlareTeeManagerInit) TryPackInit(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address, availabilityCheckValidityDurationSeconds uint64, signingPolicyValidityDurationInRewardEpochs uint64, challengeValidityDurationSeconds uint64, defaultFee *big.Int, publicExtensionCreationEnabled bool, emergencyUnpauseGracePeriodSeconds *big.Int) ([]byte, error) {
	return flareTeeManagerInit.abi.Pack("init", governanceSettings, initialGovernance, addressUpdater, availabilityCheckValidityDurationSeconds, signingPolicyValidityDurationInRewardEpochs, challengeValidityDurationSeconds, defaultFee, publicExtensionCreationEnabled, emergencyUnpauseGracePeriodSeconds)
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (flareTeeManagerInit *FlareTeeManagerInit) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := flareTeeManagerInit.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (flareTeeManagerInit *FlareTeeManagerInit) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return flareTeeManagerInit.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// FlareTeeManagerInitAllExtensionOwnersAllowed represents a AllExtensionOwnersAllowed event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitAllExtensionOwnersAllowed struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const FlareTeeManagerInitAllExtensionOwnersAllowedEventName = "AllExtensionOwnersAllowed"

// ContractEventName returns the user-defined event name.
func (FlareTeeManagerInitAllExtensionOwnersAllowed) ContractEventName() string {
	return FlareTeeManagerInitAllExtensionOwnersAllowedEventName
}

// UnpackAllExtensionOwnersAllowedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllExtensionOwnersAllowed()
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackAllExtensionOwnersAllowedEvent(log *types.Log) (*FlareTeeManagerInitAllExtensionOwnersAllowed, error) {
	event := "AllExtensionOwnersAllowed"
	if len(log.Topics) == 0 || log.Topics[0] != flareTeeManagerInit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareTeeManagerInitAllExtensionOwnersAllowed)
	if len(log.Data) > 0 {
		if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareTeeManagerInit.abi.Events[event].Inputs {
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

// FlareTeeManagerInitAllExtensionOwnersDisallowed represents a AllExtensionOwnersDisallowed event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitAllExtensionOwnersDisallowed struct {
	Raw *types.Log // Blockchain specific contextual infos
}

const FlareTeeManagerInitAllExtensionOwnersDisallowedEventName = "AllExtensionOwnersDisallowed"

// ContractEventName returns the user-defined event name.
func (FlareTeeManagerInitAllExtensionOwnersDisallowed) ContractEventName() string {
	return FlareTeeManagerInitAllExtensionOwnersDisallowedEventName
}

// UnpackAllExtensionOwnersDisallowedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AllExtensionOwnersDisallowed()
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackAllExtensionOwnersDisallowedEvent(log *types.Log) (*FlareTeeManagerInitAllExtensionOwnersDisallowed, error) {
	event := "AllExtensionOwnersDisallowed"
	if len(log.Topics) == 0 || log.Topics[0] != flareTeeManagerInit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareTeeManagerInitAllExtensionOwnersDisallowed)
	if len(log.Data) > 0 {
		if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareTeeManagerInit.abi.Events[event].Inputs {
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

// FlareTeeManagerInitDefaultFeeSet represents a DefaultFeeSet event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitDefaultFeeSet struct {
	DefaultFee *big.Int
	Raw        *types.Log // Blockchain specific contextual infos
}

const FlareTeeManagerInitDefaultFeeSetEventName = "DefaultFeeSet"

// ContractEventName returns the user-defined event name.
func (FlareTeeManagerInitDefaultFeeSet) ContractEventName() string {
	return FlareTeeManagerInitDefaultFeeSetEventName
}

// UnpackDefaultFeeSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DefaultFeeSet(uint256 defaultFee)
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackDefaultFeeSetEvent(log *types.Log) (*FlareTeeManagerInitDefaultFeeSet, error) {
	event := "DefaultFeeSet"
	if len(log.Topics) == 0 || log.Topics[0] != flareTeeManagerInit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareTeeManagerInitDefaultFeeSet)
	if len(log.Data) > 0 {
		if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareTeeManagerInit.abi.Events[event].Inputs {
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

// FlareTeeManagerInitEmergencyUnpauseGracePeriodSet represents a EmergencyUnpauseGracePeriodSet event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitEmergencyUnpauseGracePeriodSet struct {
	GraceSeconds *big.Int
	Raw          *types.Log // Blockchain specific contextual infos
}

const FlareTeeManagerInitEmergencyUnpauseGracePeriodSetEventName = "EmergencyUnpauseGracePeriodSet"

// ContractEventName returns the user-defined event name.
func (FlareTeeManagerInitEmergencyUnpauseGracePeriodSet) ContractEventName() string {
	return FlareTeeManagerInitEmergencyUnpauseGracePeriodSetEventName
}

// UnpackEmergencyUnpauseGracePeriodSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EmergencyUnpauseGracePeriodSet(uint256 graceSeconds)
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackEmergencyUnpauseGracePeriodSetEvent(log *types.Log) (*FlareTeeManagerInitEmergencyUnpauseGracePeriodSet, error) {
	event := "EmergencyUnpauseGracePeriodSet"
	if len(log.Topics) == 0 || log.Topics[0] != flareTeeManagerInit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareTeeManagerInitEmergencyUnpauseGracePeriodSet)
	if len(log.Data) > 0 {
		if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareTeeManagerInit.abi.Events[event].Inputs {
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

// FlareTeeManagerInitGovernanceInitialised represents a GovernanceInitialised event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const FlareTeeManagerInitGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (FlareTeeManagerInitGovernanceInitialised) ContractEventName() string {
	return FlareTeeManagerInitGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackGovernanceInitialisedEvent(log *types.Log) (*FlareTeeManagerInitGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != flareTeeManagerInit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareTeeManagerInitGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareTeeManagerInit.abi.Events[event].Inputs {
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

// FlareTeeManagerInitInitialized represents a Initialized event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const FlareTeeManagerInitInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (FlareTeeManagerInitInitialized) ContractEventName() string {
	return FlareTeeManagerInitInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackInitializedEvent(log *types.Log) (*FlareTeeManagerInitInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != flareTeeManagerInit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareTeeManagerInitInitialized)
	if len(log.Data) > 0 {
		if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareTeeManagerInit.abi.Events[event].Inputs {
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

// FlareTeeManagerInitSettingsUpdated represents a SettingsUpdated event raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitSettingsUpdated struct {
	AvailabilityCheckValidityDurationSeconds    uint64
	SigningPolicyValidityDurationInRewardEpochs uint64
	ChallengeValidityDurationSeconds            uint64
	Raw                                         *types.Log // Blockchain specific contextual infos
}

const FlareTeeManagerInitSettingsUpdatedEventName = "SettingsUpdated"

// ContractEventName returns the user-defined event name.
func (FlareTeeManagerInitSettingsUpdated) ContractEventName() string {
	return FlareTeeManagerInitSettingsUpdatedEventName
}

// UnpackSettingsUpdatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SettingsUpdated(uint64 availabilityCheckValidityDurationSeconds, uint64 signingPolicyValidityDurationInRewardEpochs, uint64 challengeValidityDurationSeconds)
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackSettingsUpdatedEvent(log *types.Log) (*FlareTeeManagerInitSettingsUpdated, error) {
	event := "SettingsUpdated"
	if len(log.Topics) == 0 || log.Topics[0] != flareTeeManagerInit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareTeeManagerInitSettingsUpdated)
	if len(log.Data) > 0 {
		if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareTeeManagerInit.abi.Events[event].Inputs {
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
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], flareTeeManagerInit.abi.Errors["DefaultFeeZero"].ID.Bytes()[:4]) {
		return flareTeeManagerInit.UnpackDefaultFeeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], flareTeeManagerInit.abi.Errors["GovernedAddressZero"].ID.Bytes()[:4]) {
		return flareTeeManagerInit.UnpackGovernedAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], flareTeeManagerInit.abi.Errors["GovernedAlreadyInitialized"].ID.Bytes()[:4]) {
		return flareTeeManagerInit.UnpackGovernedAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], flareTeeManagerInit.abi.Errors["GracePeriodTooLong"].ID.Bytes()[:4]) {
		return flareTeeManagerInit.UnpackGracePeriodTooLongError(raw[4:])
	}
	if bytes.Equal(raw[:4], flareTeeManagerInit.abi.Errors["GracePeriodTooShort"].ID.Bytes()[:4]) {
		return flareTeeManagerInit.UnpackGracePeriodTooShortError(raw[4:])
	}
	if bytes.Equal(raw[:4], flareTeeManagerInit.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return flareTeeManagerInit.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], flareTeeManagerInit.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return flareTeeManagerInit.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], flareTeeManagerInit.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return flareTeeManagerInit.UnpackNotInitializingError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// FlareTeeManagerInitDefaultFeeZero represents a DefaultFeeZero error raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitDefaultFeeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DefaultFeeZero()
func FlareTeeManagerInitDefaultFeeZeroErrorID() common.Hash {
	return common.HexToHash("0x85ed2180a1896c3c873e33b45c04d2bf59b01866e984c0ebe6c53b66730eff4f")
}

// UnpackDefaultFeeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DefaultFeeZero()
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackDefaultFeeZeroError(raw []byte) (*FlareTeeManagerInitDefaultFeeZero, error) {
	out := new(FlareTeeManagerInitDefaultFeeZero)
	if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, "DefaultFeeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// FlareTeeManagerInitGovernedAddressZero represents a GovernedAddressZero error raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitGovernedAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAddressZero()
func FlareTeeManagerInitGovernedAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x695725f6b429bac0cc5a2fa88214e31520e0c69f2fb0e2fb881a7130a564f49f")
}

// UnpackGovernedAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAddressZero()
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackGovernedAddressZeroError(raw []byte) (*FlareTeeManagerInitGovernedAddressZero, error) {
	out := new(FlareTeeManagerInitGovernedAddressZero)
	if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, "GovernedAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// FlareTeeManagerInitGovernedAlreadyInitialized represents a GovernedAlreadyInitialized error raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitGovernedAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAlreadyInitialized()
func FlareTeeManagerInitGovernedAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0xb2f0162156d83828197776d56bbfba0a7a9d5936fad70f0a3014ae18cc7aa33c")
}

// UnpackGovernedAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAlreadyInitialized()
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackGovernedAlreadyInitializedError(raw []byte) (*FlareTeeManagerInitGovernedAlreadyInitialized, error) {
	out := new(FlareTeeManagerInitGovernedAlreadyInitialized)
	if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, "GovernedAlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// FlareTeeManagerInitGracePeriodTooLong represents a GracePeriodTooLong error raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitGracePeriodTooLong struct {
	GraceSeconds *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GracePeriodTooLong(uint256 graceSeconds)
func FlareTeeManagerInitGracePeriodTooLongErrorID() common.Hash {
	return common.HexToHash("0x8ff948d6119974fa3ecbb383b14b7f8685cce2762011e98acf71c1a9158cbad5")
}

// UnpackGracePeriodTooLongError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GracePeriodTooLong(uint256 graceSeconds)
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackGracePeriodTooLongError(raw []byte) (*FlareTeeManagerInitGracePeriodTooLong, error) {
	out := new(FlareTeeManagerInitGracePeriodTooLong)
	if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, "GracePeriodTooLong", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// FlareTeeManagerInitGracePeriodTooShort represents a GracePeriodTooShort error raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitGracePeriodTooShort struct {
	GraceSeconds *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GracePeriodTooShort(uint256 graceSeconds)
func FlareTeeManagerInitGracePeriodTooShortErrorID() common.Hash {
	return common.HexToHash("0xa77e74e78e7adccece83e0a693bd733e0f6da702830477df6f7d58f8b3927e2e")
}

// UnpackGracePeriodTooShortError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GracePeriodTooShort(uint256 graceSeconds)
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackGracePeriodTooShortError(raw []byte) (*FlareTeeManagerInitGracePeriodTooShort, error) {
	out := new(FlareTeeManagerInitGracePeriodTooShort)
	if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, "GracePeriodTooShort", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// FlareTeeManagerInitInvalidDuration represents a InvalidDuration error raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func FlareTeeManagerInitInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackInvalidDurationError(raw []byte) (*FlareTeeManagerInitInvalidDuration, error) {
	out := new(FlareTeeManagerInitInvalidDuration)
	if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// FlareTeeManagerInitInvalidInitialization represents a InvalidInitialization error raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func FlareTeeManagerInitInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackInvalidInitializationError(raw []byte) (*FlareTeeManagerInitInvalidInitialization, error) {
	out := new(FlareTeeManagerInitInvalidInitialization)
	if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// FlareTeeManagerInitNotInitializing represents a NotInitializing error raised by the FlareTeeManagerInit contract.
type FlareTeeManagerInitNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func FlareTeeManagerInitNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (flareTeeManagerInit *FlareTeeManagerInit) UnpackNotInitializingError(raw []byte) (*FlareTeeManagerInitNotInitializing, error) {
	out := new(FlareTeeManagerInitNotInitializing)
	if err := flareTeeManagerInit.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}
