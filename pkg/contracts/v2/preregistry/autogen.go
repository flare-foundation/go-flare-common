// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package preregistry

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

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// PreregistryMetaData contains all meta data concerning the Preregistry contract.
var PreregistryMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"}],\"name\":\"VoterPreRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"}],\"name\":\"VoterRegistrationFailed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getPreRegisteredVoters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterSignature\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"isVoterPreRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"preRegisterVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Preregistry",
}

// Preregistry is an auto generated Go binding around an Ethereum contract.
type Preregistry struct {
	abi abi.ABI
}

// NewPreregistry creates a new instance of Preregistry.
func NewPreregistry() *Preregistry {
	parsed, err := PreregistryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Preregistry{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Preregistry) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetPreRegisteredVoters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x010f079d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPreRegisteredVoters(uint256 _rewardEpochId) view returns(address[])
func (preregistry *Preregistry) PackGetPreRegisteredVoters(rewardEpochId *big.Int) []byte {
	enc, err := preregistry.abi.Pack("getPreRegisteredVoters", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPreRegisteredVoters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x010f079d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPreRegisteredVoters(uint256 _rewardEpochId) view returns(address[])
func (preregistry *Preregistry) TryPackGetPreRegisteredVoters(rewardEpochId *big.Int) ([]byte, error) {
	return preregistry.abi.Pack("getPreRegisteredVoters", rewardEpochId)
}

// UnpackGetPreRegisteredVoters is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x010f079d.
//
// Solidity: function getPreRegisteredVoters(uint256 _rewardEpochId) view returns(address[])
func (preregistry *Preregistry) UnpackGetPreRegisteredVoters(data []byte) ([]common.Address, error) {
	out, err := preregistry.abi.Unpack("getPreRegisteredVoters", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetVoterSignature is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3beabf52.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterSignature(uint256 _rewardEpochId, address _voter) view returns((uint8,bytes32,bytes32))
func (preregistry *Preregistry) PackGetVoterSignature(rewardEpochId *big.Int, voter common.Address) []byte {
	enc, err := preregistry.abi.Pack("getVoterSignature", rewardEpochId, voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterSignature is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3beabf52.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterSignature(uint256 _rewardEpochId, address _voter) view returns((uint8,bytes32,bytes32))
func (preregistry *Preregistry) TryPackGetVoterSignature(rewardEpochId *big.Int, voter common.Address) ([]byte, error) {
	return preregistry.abi.Pack("getVoterSignature", rewardEpochId, voter)
}

// UnpackGetVoterSignature is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3beabf52.
//
// Solidity: function getVoterSignature(uint256 _rewardEpochId, address _voter) view returns((uint8,bytes32,bytes32))
func (preregistry *Preregistry) UnpackGetVoterSignature(data []byte) (Signature, error) {
	out, err := preregistry.abi.Unpack("getVoterSignature", data)
	if err != nil {
		return *new(Signature), err
	}
	out0 := *abi.ConvertType(out[0], new(Signature)).(*Signature)
	return out0, nil
}

// PackIsVoterPreRegistered is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x98041046.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isVoterPreRegistered(uint256 _rewardEpochId, address _voter) view returns(bool)
func (preregistry *Preregistry) PackIsVoterPreRegistered(rewardEpochId *big.Int, voter common.Address) []byte {
	enc, err := preregistry.abi.Pack("isVoterPreRegistered", rewardEpochId, voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsVoterPreRegistered is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x98041046.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isVoterPreRegistered(uint256 _rewardEpochId, address _voter) view returns(bool)
func (preregistry *Preregistry) TryPackIsVoterPreRegistered(rewardEpochId *big.Int, voter common.Address) ([]byte, error) {
	return preregistry.abi.Pack("isVoterPreRegistered", rewardEpochId, voter)
}

// UnpackIsVoterPreRegistered is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x98041046.
//
// Solidity: function isVoterPreRegistered(uint256 _rewardEpochId, address _voter) view returns(bool)
func (preregistry *Preregistry) UnpackIsVoterPreRegistered(data []byte) (bool, error) {
	out, err := preregistry.abi.Unpack("isVoterPreRegistered", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackPreRegisterVoter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xea8eb77d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function preRegisterVoter(address _voter, (uint8,bytes32,bytes32) _signature) returns()
func (preregistry *Preregistry) PackPreRegisterVoter(voter common.Address, signature Signature) []byte {
	enc, err := preregistry.abi.Pack("preRegisterVoter", voter, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPreRegisterVoter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xea8eb77d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function preRegisterVoter(address _voter, (uint8,bytes32,bytes32) _signature) returns()
func (preregistry *Preregistry) TryPackPreRegisterVoter(voter common.Address, signature Signature) ([]byte, error) {
	return preregistry.abi.Pack("preRegisterVoter", voter, signature)
}

// PreregistryVoterPreRegistered represents a VoterPreRegistered event raised by the Preregistry contract.
type PreregistryVoterPreRegistered struct {
	Voter         common.Address
	RewardEpochId uint32
	Raw           *types.Log // Blockchain specific contextual infos
}

const PreregistryVoterPreRegisteredEventName = "VoterPreRegistered"

// ContractEventName returns the user-defined event name.
func (PreregistryVoterPreRegistered) ContractEventName() string {
	return PreregistryVoterPreRegisteredEventName
}

// UnpackVoterPreRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoterPreRegistered(address indexed voter, uint32 indexed rewardEpochId)
func (preregistry *Preregistry) UnpackVoterPreRegisteredEvent(log *types.Log) (*PreregistryVoterPreRegistered, error) {
	event := "VoterPreRegistered"
	if len(log.Topics) == 0 || log.Topics[0] != preregistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PreregistryVoterPreRegistered)
	if len(log.Data) > 0 {
		if err := preregistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range preregistry.abi.Events[event].Inputs {
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

// PreregistryVoterRegistrationFailed represents a VoterRegistrationFailed event raised by the Preregistry contract.
type PreregistryVoterRegistrationFailed struct {
	Voter         common.Address
	RewardEpochId uint32
	Raw           *types.Log // Blockchain specific contextual infos
}

const PreregistryVoterRegistrationFailedEventName = "VoterRegistrationFailed"

// ContractEventName returns the user-defined event name.
func (PreregistryVoterRegistrationFailed) ContractEventName() string {
	return PreregistryVoterRegistrationFailedEventName
}

// UnpackVoterRegistrationFailedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoterRegistrationFailed(address indexed voter, uint32 indexed rewardEpochId)
func (preregistry *Preregistry) UnpackVoterRegistrationFailedEvent(log *types.Log) (*PreregistryVoterRegistrationFailed, error) {
	event := "VoterRegistrationFailed"
	if len(log.Topics) == 0 || log.Topics[0] != preregistry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(PreregistryVoterRegistrationFailed)
	if len(log.Data) > 0 {
		if err := preregistry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range preregistry.abi.Events[event].Inputs {
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
