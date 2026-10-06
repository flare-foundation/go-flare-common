// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diamondgovernance

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

// IDiamondFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// DiamondGovernanceMetaData contains all meta data concerning the DiamondGovernance contract.
var DiamondGovernanceMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotAddFunctionToDiamondThatAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"_selectors\",\"type\":\"bytes4[]\"}],\"name\":\"CannotAddSelectorsToZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotRemoveFunctionThatDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotRemoveImmutableFunction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotReplaceFunctionThatDoesNotExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotReplaceFunctionWithTheSameFunctionFromTheSameFacet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"_selectors\",\"type\":\"bytes4[]\"}],\"name\":\"CannotReplaceFunctionsFromFacetWithZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"CannotReplaceImmutableFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_action\",\"type\":\"uint8\"}],\"name\":\"IncorrectFacetCutAction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_initializationContractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"InitializationFunctionReverted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_message\",\"type\":\"string\"}],\"name\":\"NoBytecodeAtAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facetAddress\",\"type\":\"address\"}],\"name\":\"NoSelectorsProvidedForFacetForCut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facetAddress\",\"type\":\"address\"}],\"name\":\"RemoveFacetAddressMustBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockCallNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamond.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamond.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamond.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamond.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamond.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamond.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "DiamondGovernance",
}

// DiamondGovernance is an auto generated Go binding around an Ethereum contract.
type DiamondGovernance struct {
	abi abi.ABI
}

// NewDiamondGovernance creates a new instance of DiamondGovernance.
func NewDiamondGovernance() *DiamondGovernance {
	parsed, err := DiamondGovernanceMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &DiamondGovernance{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *DiamondGovernance) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16fc2f6d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (diamondGovernance *DiamondGovernance) PackCancelGovernanceCall(encodedCall []byte) []byte {
	enc, err := diamondGovernance.abi.Pack("cancelGovernanceCall", encodedCall)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16fc2f6d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (diamondGovernance *DiamondGovernance) TryPackCancelGovernanceCall(encodedCall []byte) ([]byte, error) {
	return diamondGovernance.abi.Pack("cancelGovernanceCall", encodedCall)
}

// PackDiamondCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f931c1c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (diamondGovernance *DiamondGovernance) PackDiamondCut(diamondCut []IDiamondFacetCut, init common.Address, calldata []byte) []byte {
	enc, err := diamondGovernance.abi.Pack("diamondCut", diamondCut, init, calldata)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDiamondCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f931c1c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (diamondGovernance *DiamondGovernance) TryPackDiamondCut(diamondCut []IDiamondFacetCut, init common.Address, calldata []byte) ([]byte, error) {
	return diamondGovernance.abi.Pack("diamondCut", diamondCut, init, calldata)
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (diamondGovernance *DiamondGovernance) PackExecuteGovernanceCall(encodedCall []byte) []byte {
	enc, err := diamondGovernance.abi.Pack("executeGovernanceCall", encodedCall)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (diamondGovernance *DiamondGovernance) TryPackExecuteGovernanceCall(encodedCall []byte) ([]byte, error) {
	return diamondGovernance.abi.Pack("executeGovernanceCall", encodedCall)
}

// PackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governance() view returns(address)
func (diamondGovernance *DiamondGovernance) PackGovernance() []byte {
	enc, err := diamondGovernance.abi.Pack("governance")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function governance() view returns(address)
func (diamondGovernance *DiamondGovernance) TryPackGovernance() ([]byte, error) {
	return diamondGovernance.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (diamondGovernance *DiamondGovernance) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := diamondGovernance.abi.Unpack("governance", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGovernanceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62354e03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governanceSettings() view returns(address)
func (diamondGovernance *DiamondGovernance) PackGovernanceSettings() []byte {
	enc, err := diamondGovernance.abi.Pack("governanceSettings")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGovernanceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62354e03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function governanceSettings() view returns(address)
func (diamondGovernance *DiamondGovernance) TryPackGovernanceSettings() ([]byte, error) {
	return diamondGovernance.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (diamondGovernance *DiamondGovernance) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := diamondGovernance.abi.Unpack("governanceSettings", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (diamondGovernance *DiamondGovernance) PackIsExecutor(address common.Address) []byte {
	enc, err := diamondGovernance.abi.Pack("isExecutor", address)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (diamondGovernance *DiamondGovernance) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return diamondGovernance.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (diamondGovernance *DiamondGovernance) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := diamondGovernance.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (diamondGovernance *DiamondGovernance) PackProductionMode() []byte {
	enc, err := diamondGovernance.abi.Pack("productionMode")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function productionMode() view returns(bool)
func (diamondGovernance *DiamondGovernance) TryPackProductionMode() ([]byte, error) {
	return diamondGovernance.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (diamondGovernance *DiamondGovernance) UnpackProductionMode(data []byte) (bool, error) {
	out, err := diamondGovernance.abi.Unpack("productionMode", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (diamondGovernance *DiamondGovernance) PackSwitchToProductionMode() []byte {
	enc, err := diamondGovernance.abi.Pack("switchToProductionMode")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function switchToProductionMode() returns()
func (diamondGovernance *DiamondGovernance) TryPackSwitchToProductionMode() ([]byte, error) {
	return diamondGovernance.abi.Pack("switchToProductionMode")
}

// DiamondGovernanceDiamondCut represents a DiamondCut event raised by the DiamondGovernance contract.
type DiamondGovernanceDiamondCut struct {
	DiamondCut []IDiamondFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const DiamondGovernanceDiamondCutEventName = "DiamondCut"

// ContractEventName returns the user-defined event name.
func (DiamondGovernanceDiamondCut) ContractEventName() string {
	return DiamondGovernanceDiamondCutEventName
}

// UnpackDiamondCutEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (diamondGovernance *DiamondGovernance) UnpackDiamondCutEvent(log *types.Log) (*DiamondGovernanceDiamondCut, error) {
	event := "DiamondCut"
	if len(log.Topics) == 0 || log.Topics[0] != diamondGovernance.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DiamondGovernanceDiamondCut)
	if len(log.Data) > 0 {
		if err := diamondGovernance.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diamondGovernance.abi.Events[event].Inputs {
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

// DiamondGovernanceDiamondCut0 represents a DiamondCut0 event raised by the DiamondGovernance contract.
type DiamondGovernanceDiamondCut0 struct {
	DiamondCut []IDiamondFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const DiamondGovernanceDiamondCut0EventName = "DiamondCut0"

// ContractEventName returns the user-defined event name.
func (DiamondGovernanceDiamondCut0) ContractEventName() string {
	return DiamondGovernanceDiamondCut0EventName
}

// UnpackDiamondCut0Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (diamondGovernance *DiamondGovernance) UnpackDiamondCut0Event(log *types.Log) (*DiamondGovernanceDiamondCut0, error) {
	event := "DiamondCut0"
	if len(log.Topics) == 0 || log.Topics[0] != diamondGovernance.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DiamondGovernanceDiamondCut0)
	if len(log.Data) > 0 {
		if err := diamondGovernance.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diamondGovernance.abi.Events[event].Inputs {
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

// DiamondGovernanceGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the DiamondGovernance contract.
type DiamondGovernanceGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const DiamondGovernanceGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (DiamondGovernanceGovernanceCallTimelocked) ContractEventName() string {
	return DiamondGovernanceGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (diamondGovernance *DiamondGovernance) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*DiamondGovernanceGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != diamondGovernance.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DiamondGovernanceGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := diamondGovernance.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diamondGovernance.abi.Events[event].Inputs {
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

// DiamondGovernanceGovernanceInitialised represents a GovernanceInitialised event raised by the DiamondGovernance contract.
type DiamondGovernanceGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const DiamondGovernanceGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (DiamondGovernanceGovernanceInitialised) ContractEventName() string {
	return DiamondGovernanceGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (diamondGovernance *DiamondGovernance) UnpackGovernanceInitialisedEvent(log *types.Log) (*DiamondGovernanceGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != diamondGovernance.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DiamondGovernanceGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := diamondGovernance.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diamondGovernance.abi.Events[event].Inputs {
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

// DiamondGovernanceGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the DiamondGovernance contract.
type DiamondGovernanceGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const DiamondGovernanceGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (DiamondGovernanceGovernedProductionModeEntered) ContractEventName() string {
	return DiamondGovernanceGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (diamondGovernance *DiamondGovernance) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*DiamondGovernanceGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != diamondGovernance.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DiamondGovernanceGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := diamondGovernance.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diamondGovernance.abi.Events[event].Inputs {
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

// DiamondGovernanceInitialized represents a Initialized event raised by the DiamondGovernance contract.
type DiamondGovernanceInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const DiamondGovernanceInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (DiamondGovernanceInitialized) ContractEventName() string {
	return DiamondGovernanceInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (diamondGovernance *DiamondGovernance) UnpackInitializedEvent(log *types.Log) (*DiamondGovernanceInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != diamondGovernance.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DiamondGovernanceInitialized)
	if len(log.Data) > 0 {
		if err := diamondGovernance.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diamondGovernance.abi.Events[event].Inputs {
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

// DiamondGovernanceTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the DiamondGovernance contract.
type DiamondGovernanceTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const DiamondGovernanceTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (DiamondGovernanceTimelockedGovernanceCallCanceled) ContractEventName() string {
	return DiamondGovernanceTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (diamondGovernance *DiamondGovernance) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*DiamondGovernanceTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != diamondGovernance.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DiamondGovernanceTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := diamondGovernance.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diamondGovernance.abi.Events[event].Inputs {
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

// DiamondGovernanceTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the DiamondGovernance contract.
type DiamondGovernanceTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const DiamondGovernanceTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (DiamondGovernanceTimelockedGovernanceCallExecuted) ContractEventName() string {
	return DiamondGovernanceTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (diamondGovernance *DiamondGovernance) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*DiamondGovernanceTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != diamondGovernance.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(DiamondGovernanceTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := diamondGovernance.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range diamondGovernance.abi.Events[event].Inputs {
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
func (diamondGovernance *DiamondGovernance) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["AlreadyInProductionMode"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackAlreadyInProductionModeError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["CannotAddFunctionToDiamondThatAlreadyExists"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackCannotAddFunctionToDiamondThatAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["CannotAddSelectorsToZeroAddress"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackCannotAddSelectorsToZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["CannotRemoveFunctionThatDoesNotExist"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackCannotRemoveFunctionThatDoesNotExistError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["CannotRemoveImmutableFunction"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackCannotRemoveImmutableFunctionError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["CannotReplaceFunctionThatDoesNotExists"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackCannotReplaceFunctionThatDoesNotExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["CannotReplaceFunctionWithTheSameFunctionFromTheSameFacet"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackCannotReplaceFunctionWithTheSameFunctionFromTheSameFacetError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["CannotReplaceFunctionsFromFacetWithZeroAddress"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackCannotReplaceFunctionsFromFacetWithZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["CannotReplaceImmutableFunction"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackCannotReplaceImmutableFunctionError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["GovernedAddressZero"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackGovernedAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["GovernedAlreadyInitialized"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackGovernedAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["IncorrectFacetCutAction"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackIncorrectFacetCutActionError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["InitializationFunctionReverted"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackInitializationFunctionRevertedError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["NoBytecodeAtAddress"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackNoBytecodeAtAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["NoSelectorsProvidedForFacetForCut"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackNoSelectorsProvidedForFacetForCutError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["OnlyExecutor"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["RemoveFacetAddressMustBeZeroAddress"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackRemoveFacetAddressMustBeZeroAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["TimelockCallNotFound"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackTimelockCallNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], diamondGovernance.abi.Errors["TimelockNotAllowedYet"].ID.Bytes()[:4]) {
		return diamondGovernance.UnpackTimelockNotAllowedYetError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// DiamondGovernanceAlreadyInProductionMode represents a AlreadyInProductionMode error raised by the DiamondGovernance contract.
type DiamondGovernanceAlreadyInProductionMode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInProductionMode()
func DiamondGovernanceAlreadyInProductionModeErrorID() common.Hash {
	return common.HexToHash("0x72a23ec2b1fdbc3de7d5f411b572fcf8ad8409803196930f63ddf40a2a2b2594")
}

// UnpackAlreadyInProductionModeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInProductionMode()
func (diamondGovernance *DiamondGovernance) UnpackAlreadyInProductionModeError(raw []byte) (*DiamondGovernanceAlreadyInProductionMode, error) {
	out := new(DiamondGovernanceAlreadyInProductionMode)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "AlreadyInProductionMode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceCannotAddFunctionToDiamondThatAlreadyExists represents a CannotAddFunctionToDiamondThatAlreadyExists error raised by the DiamondGovernance contract.
type DiamondGovernanceCannotAddFunctionToDiamondThatAlreadyExists struct {
	Selector [4]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotAddFunctionToDiamondThatAlreadyExists(bytes4 _selector)
func DiamondGovernanceCannotAddFunctionToDiamondThatAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0xebbf5d073128f7e2a4f2b1c8052d80d9606e9a1bfae6b0a2090b6c7dfa3b9b1f")
}

// UnpackCannotAddFunctionToDiamondThatAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotAddFunctionToDiamondThatAlreadyExists(bytes4 _selector)
func (diamondGovernance *DiamondGovernance) UnpackCannotAddFunctionToDiamondThatAlreadyExistsError(raw []byte) (*DiamondGovernanceCannotAddFunctionToDiamondThatAlreadyExists, error) {
	out := new(DiamondGovernanceCannotAddFunctionToDiamondThatAlreadyExists)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "CannotAddFunctionToDiamondThatAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceCannotAddSelectorsToZeroAddress represents a CannotAddSelectorsToZeroAddress error raised by the DiamondGovernance contract.
type DiamondGovernanceCannotAddSelectorsToZeroAddress struct {
	Selectors [][4]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotAddSelectorsToZeroAddress(bytes4[] _selectors)
func DiamondGovernanceCannotAddSelectorsToZeroAddressErrorID() common.Hash {
	return common.HexToHash("0x0ae3681ca6bb92cef652575526bb043f0f6f307616fad5b63e1c77c1afab36c6")
}

// UnpackCannotAddSelectorsToZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotAddSelectorsToZeroAddress(bytes4[] _selectors)
func (diamondGovernance *DiamondGovernance) UnpackCannotAddSelectorsToZeroAddressError(raw []byte) (*DiamondGovernanceCannotAddSelectorsToZeroAddress, error) {
	out := new(DiamondGovernanceCannotAddSelectorsToZeroAddress)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "CannotAddSelectorsToZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceCannotRemoveFunctionThatDoesNotExist represents a CannotRemoveFunctionThatDoesNotExist error raised by the DiamondGovernance contract.
type DiamondGovernanceCannotRemoveFunctionThatDoesNotExist struct {
	Selector [4]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotRemoveFunctionThatDoesNotExist(bytes4 _selector)
func DiamondGovernanceCannotRemoveFunctionThatDoesNotExistErrorID() common.Hash {
	return common.HexToHash("0x7a08a22d0efea4f9bf13381790fd48da1f400978a1400fce86a57d2eb5bd6195")
}

// UnpackCannotRemoveFunctionThatDoesNotExistError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotRemoveFunctionThatDoesNotExist(bytes4 _selector)
func (diamondGovernance *DiamondGovernance) UnpackCannotRemoveFunctionThatDoesNotExistError(raw []byte) (*DiamondGovernanceCannotRemoveFunctionThatDoesNotExist, error) {
	out := new(DiamondGovernanceCannotRemoveFunctionThatDoesNotExist)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "CannotRemoveFunctionThatDoesNotExist", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceCannotRemoveImmutableFunction represents a CannotRemoveImmutableFunction error raised by the DiamondGovernance contract.
type DiamondGovernanceCannotRemoveImmutableFunction struct {
	Selector [4]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotRemoveImmutableFunction(bytes4 _selector)
func DiamondGovernanceCannotRemoveImmutableFunctionErrorID() common.Hash {
	return common.HexToHash("0x6fafeb08728b9d836e4aff337a72ac92446bd88cb42d0f146bf7bbc80dd7f271")
}

// UnpackCannotRemoveImmutableFunctionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotRemoveImmutableFunction(bytes4 _selector)
func (diamondGovernance *DiamondGovernance) UnpackCannotRemoveImmutableFunctionError(raw []byte) (*DiamondGovernanceCannotRemoveImmutableFunction, error) {
	out := new(DiamondGovernanceCannotRemoveImmutableFunction)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "CannotRemoveImmutableFunction", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceCannotReplaceFunctionThatDoesNotExists represents a CannotReplaceFunctionThatDoesNotExists error raised by the DiamondGovernance contract.
type DiamondGovernanceCannotReplaceFunctionThatDoesNotExists struct {
	Selector [4]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotReplaceFunctionThatDoesNotExists(bytes4 _selector)
func DiamondGovernanceCannotReplaceFunctionThatDoesNotExistsErrorID() common.Hash {
	return common.HexToHash("0x7479f93954bf02fdc4f90ad7558f6ff9f2d52027eb6c4dad443c083f6c70ca63")
}

// UnpackCannotReplaceFunctionThatDoesNotExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotReplaceFunctionThatDoesNotExists(bytes4 _selector)
func (diamondGovernance *DiamondGovernance) UnpackCannotReplaceFunctionThatDoesNotExistsError(raw []byte) (*DiamondGovernanceCannotReplaceFunctionThatDoesNotExists, error) {
	out := new(DiamondGovernanceCannotReplaceFunctionThatDoesNotExists)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "CannotReplaceFunctionThatDoesNotExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceCannotReplaceFunctionWithTheSameFunctionFromTheSameFacet represents a CannotReplaceFunctionWithTheSameFunctionFromTheSameFacet error raised by the DiamondGovernance contract.
type DiamondGovernanceCannotReplaceFunctionWithTheSameFunctionFromTheSameFacet struct {
	Selector [4]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotReplaceFunctionWithTheSameFunctionFromTheSameFacet(bytes4 _selector)
func DiamondGovernanceCannotReplaceFunctionWithTheSameFunctionFromTheSameFacetErrorID() common.Hash {
	return common.HexToHash("0x358d9d1a0a28c86b5b327b120c2e53a616ea3c5d331d3dcb436fd20e28aa9de8")
}

// UnpackCannotReplaceFunctionWithTheSameFunctionFromTheSameFacetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotReplaceFunctionWithTheSameFunctionFromTheSameFacet(bytes4 _selector)
func (diamondGovernance *DiamondGovernance) UnpackCannotReplaceFunctionWithTheSameFunctionFromTheSameFacetError(raw []byte) (*DiamondGovernanceCannotReplaceFunctionWithTheSameFunctionFromTheSameFacet, error) {
	out := new(DiamondGovernanceCannotReplaceFunctionWithTheSameFunctionFromTheSameFacet)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "CannotReplaceFunctionWithTheSameFunctionFromTheSameFacet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceCannotReplaceFunctionsFromFacetWithZeroAddress represents a CannotReplaceFunctionsFromFacetWithZeroAddress error raised by the DiamondGovernance contract.
type DiamondGovernanceCannotReplaceFunctionsFromFacetWithZeroAddress struct {
	Selectors [][4]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotReplaceFunctionsFromFacetWithZeroAddress(bytes4[] _selectors)
func DiamondGovernanceCannotReplaceFunctionsFromFacetWithZeroAddressErrorID() common.Hash {
	return common.HexToHash("0xcd98a96f95158128481f7e36badb3a3c6259c14bd6f66d74a5d72651c76847e6")
}

// UnpackCannotReplaceFunctionsFromFacetWithZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotReplaceFunctionsFromFacetWithZeroAddress(bytes4[] _selectors)
func (diamondGovernance *DiamondGovernance) UnpackCannotReplaceFunctionsFromFacetWithZeroAddressError(raw []byte) (*DiamondGovernanceCannotReplaceFunctionsFromFacetWithZeroAddress, error) {
	out := new(DiamondGovernanceCannotReplaceFunctionsFromFacetWithZeroAddress)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "CannotReplaceFunctionsFromFacetWithZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceCannotReplaceImmutableFunction represents a CannotReplaceImmutableFunction error raised by the DiamondGovernance contract.
type DiamondGovernanceCannotReplaceImmutableFunction struct {
	Selector [4]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotReplaceImmutableFunction(bytes4 _selector)
func DiamondGovernanceCannotReplaceImmutableFunctionErrorID() common.Hash {
	return common.HexToHash("0x520300da3f8bf9794a770829a87024e7adc7e2793266f77e7d2c7f904868893f")
}

// UnpackCannotReplaceImmutableFunctionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotReplaceImmutableFunction(bytes4 _selector)
func (diamondGovernance *DiamondGovernance) UnpackCannotReplaceImmutableFunctionError(raw []byte) (*DiamondGovernanceCannotReplaceImmutableFunction, error) {
	out := new(DiamondGovernanceCannotReplaceImmutableFunction)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "CannotReplaceImmutableFunction", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceGovernedAddressZero represents a GovernedAddressZero error raised by the DiamondGovernance contract.
type DiamondGovernanceGovernedAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAddressZero()
func DiamondGovernanceGovernedAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x695725f6b429bac0cc5a2fa88214e31520e0c69f2fb0e2fb881a7130a564f49f")
}

// UnpackGovernedAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAddressZero()
func (diamondGovernance *DiamondGovernance) UnpackGovernedAddressZeroError(raw []byte) (*DiamondGovernanceGovernedAddressZero, error) {
	out := new(DiamondGovernanceGovernedAddressZero)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "GovernedAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceGovernedAlreadyInitialized represents a GovernedAlreadyInitialized error raised by the DiamondGovernance contract.
type DiamondGovernanceGovernedAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAlreadyInitialized()
func DiamondGovernanceGovernedAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0xb2f0162156d83828197776d56bbfba0a7a9d5936fad70f0a3014ae18cc7aa33c")
}

// UnpackGovernedAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAlreadyInitialized()
func (diamondGovernance *DiamondGovernance) UnpackGovernedAlreadyInitializedError(raw []byte) (*DiamondGovernanceGovernedAlreadyInitialized, error) {
	out := new(DiamondGovernanceGovernedAlreadyInitialized)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "GovernedAlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceIncorrectFacetCutAction represents a IncorrectFacetCutAction error raised by the DiamondGovernance contract.
type DiamondGovernanceIncorrectFacetCutAction struct {
	Action uint8
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error IncorrectFacetCutAction(uint8 _action)
func DiamondGovernanceIncorrectFacetCutActionErrorID() common.Hash {
	return common.HexToHash("0x7fe9a41e4ca196c41352f22728aee97d4bf601f71dfb5e87e4eb04641d3f2c86")
}

// UnpackIncorrectFacetCutActionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error IncorrectFacetCutAction(uint8 _action)
func (diamondGovernance *DiamondGovernance) UnpackIncorrectFacetCutActionError(raw []byte) (*DiamondGovernanceIncorrectFacetCutAction, error) {
	out := new(DiamondGovernanceIncorrectFacetCutAction)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "IncorrectFacetCutAction", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceInitializationFunctionReverted represents a InitializationFunctionReverted error raised by the DiamondGovernance contract.
type DiamondGovernanceInitializationFunctionReverted struct {
	InitializationContractAddress common.Address
	Calldata                      []byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InitializationFunctionReverted(address _initializationContractAddress, bytes _calldata)
func DiamondGovernanceInitializationFunctionRevertedErrorID() common.Hash {
	return common.HexToHash("0x192105d7c9705d7758a80e15689daed7a9a10e21cf07f39c489e10b763120253")
}

// UnpackInitializationFunctionRevertedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InitializationFunctionReverted(address _initializationContractAddress, bytes _calldata)
func (diamondGovernance *DiamondGovernance) UnpackInitializationFunctionRevertedError(raw []byte) (*DiamondGovernanceInitializationFunctionReverted, error) {
	out := new(DiamondGovernanceInitializationFunctionReverted)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "InitializationFunctionReverted", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceInvalidInitialization represents a InvalidInitialization error raised by the DiamondGovernance contract.
type DiamondGovernanceInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func DiamondGovernanceInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (diamondGovernance *DiamondGovernance) UnpackInvalidInitializationError(raw []byte) (*DiamondGovernanceInvalidInitialization, error) {
	out := new(DiamondGovernanceInvalidInitialization)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceNoBytecodeAtAddress represents a NoBytecodeAtAddress error raised by the DiamondGovernance contract.
type DiamondGovernanceNoBytecodeAtAddress struct {
	ContractAddress common.Address
	Message         string
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoBytecodeAtAddress(address _contractAddress, string _message)
func DiamondGovernanceNoBytecodeAtAddressErrorID() common.Hash {
	return common.HexToHash("0x919834b9b707737dab30ad35affa328cdfde8817b8f50a2e80d16818d30fb53c")
}

// UnpackNoBytecodeAtAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoBytecodeAtAddress(address _contractAddress, string _message)
func (diamondGovernance *DiamondGovernance) UnpackNoBytecodeAtAddressError(raw []byte) (*DiamondGovernanceNoBytecodeAtAddress, error) {
	out := new(DiamondGovernanceNoBytecodeAtAddress)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "NoBytecodeAtAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceNoSelectorsProvidedForFacetForCut represents a NoSelectorsProvidedForFacetForCut error raised by the DiamondGovernance contract.
type DiamondGovernanceNoSelectorsProvidedForFacetForCut struct {
	FacetAddress common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoSelectorsProvidedForFacetForCut(address _facetAddress)
func DiamondGovernanceNoSelectorsProvidedForFacetForCutErrorID() common.Hash {
	return common.HexToHash("0xe767f91fc6034a704f6c864b7195a6442825339e5bbeef5e9c575e1424990d6d")
}

// UnpackNoSelectorsProvidedForFacetForCutError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoSelectorsProvidedForFacetForCut(address _facetAddress)
func (diamondGovernance *DiamondGovernance) UnpackNoSelectorsProvidedForFacetForCutError(raw []byte) (*DiamondGovernanceNoSelectorsProvidedForFacetForCut, error) {
	out := new(DiamondGovernanceNoSelectorsProvidedForFacetForCut)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "NoSelectorsProvidedForFacetForCut", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceNotInitializing represents a NotInitializing error raised by the DiamondGovernance contract.
type DiamondGovernanceNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func DiamondGovernanceNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (diamondGovernance *DiamondGovernance) UnpackNotInitializingError(raw []byte) (*DiamondGovernanceNotInitializing, error) {
	out := new(DiamondGovernanceNotInitializing)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceOnlyExecutor represents a OnlyExecutor error raised by the DiamondGovernance contract.
type DiamondGovernanceOnlyExecutor struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExecutor()
func DiamondGovernanceOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x7fb6be02aaf40e420b0e7eb67e43a3feee949e8932a3c64e33b591a436670841")
}

// UnpackOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExecutor()
func (diamondGovernance *DiamondGovernance) UnpackOnlyExecutorError(raw []byte) (*DiamondGovernanceOnlyExecutor, error) {
	out := new(DiamondGovernanceOnlyExecutor)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "OnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceOnlyGovernance represents a OnlyGovernance error raised by the DiamondGovernance contract.
type DiamondGovernanceOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func DiamondGovernanceOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (diamondGovernance *DiamondGovernance) UnpackOnlyGovernanceError(raw []byte) (*DiamondGovernanceOnlyGovernance, error) {
	out := new(DiamondGovernanceOnlyGovernance)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceRemoveFacetAddressMustBeZeroAddress represents a RemoveFacetAddressMustBeZeroAddress error raised by the DiamondGovernance contract.
type DiamondGovernanceRemoveFacetAddressMustBeZeroAddress struct {
	FacetAddress common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error RemoveFacetAddressMustBeZeroAddress(address _facetAddress)
func DiamondGovernanceRemoveFacetAddressMustBeZeroAddressErrorID() common.Hash {
	return common.HexToHash("0xd091bc81dd807ff038f064cb6673ce29f80ea293efc70736022051150cfcbd31")
}

// UnpackRemoveFacetAddressMustBeZeroAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error RemoveFacetAddressMustBeZeroAddress(address _facetAddress)
func (diamondGovernance *DiamondGovernance) UnpackRemoveFacetAddressMustBeZeroAddressError(raw []byte) (*DiamondGovernanceRemoveFacetAddressMustBeZeroAddress, error) {
	out := new(DiamondGovernanceRemoveFacetAddressMustBeZeroAddress)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "RemoveFacetAddressMustBeZeroAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceTimelockCallNotFound represents a TimelockCallNotFound error raised by the DiamondGovernance contract.
type DiamondGovernanceTimelockCallNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockCallNotFound()
func DiamondGovernanceTimelockCallNotFoundErrorID() common.Hash {
	return common.HexToHash("0x262ac987dd09120380d393a0c52cdef1ca58949d1107dc26d8da9f7b527eb33d")
}

// UnpackTimelockCallNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockCallNotFound()
func (diamondGovernance *DiamondGovernance) UnpackTimelockCallNotFoundError(raw []byte) (*DiamondGovernanceTimelockCallNotFound, error) {
	out := new(DiamondGovernanceTimelockCallNotFound)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "TimelockCallNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// DiamondGovernanceTimelockNotAllowedYet represents a TimelockNotAllowedYet error raised by the DiamondGovernance contract.
type DiamondGovernanceTimelockNotAllowedYet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockNotAllowedYet()
func DiamondGovernanceTimelockNotAllowedYetErrorID() common.Hash {
	return common.HexToHash("0x6124e5c2c545bec1d02048c68500eebdabeae0cc69995a629ebccf8a62f69083")
}

// UnpackTimelockNotAllowedYetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockNotAllowedYet()
func (diamondGovernance *DiamondGovernance) UnpackTimelockNotAllowedYetError(raw []byte) (*DiamondGovernanceTimelockNotAllowedYet, error) {
	out := new(DiamondGovernanceTimelockNotAllowedYet)
	if err := diamondGovernance.abi.UnpackIntoInterface(out, "TimelockNotAllowedYet", raw); err != nil {
		return nil, err
	}
	return out, nil
}
