// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// PublicKey is an auto generated low-level Go binding around an user-defined struct.
type PublicKey struct {
	X [32]byte
	Y [32]byte
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes20\",\"name\":\"beneficiary\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"untilRewardEpochId\",\"type\":\"uint32\"}],\"name\":\"BeneficiaryChilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"submitAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"submitSignaturesAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"registrationWeight\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structSignature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"VoterRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"}],\"name\":\"VoterRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes20\",\"name\":\"_beneficiary\",\"type\":\"bytes20\"}],\"name\":\"chilledUntilRewardEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getNumberOfRegisteredVoters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getRegisteredVoters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"isVoterRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxVoters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"newSigningPolicyInitializationStartBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicKeyRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"registerVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "Registry",
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	abi abi.ABI
}

// NewRegistry creates a new instance of Registry.
func NewRegistry() *Registry {
	parsed, err := RegistryMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Registry{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Registry) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackChilledUntilRewardEpochId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3c5cb76f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function chilledUntilRewardEpochId(bytes20 _beneficiary) view returns(uint256 _rewardEpochId)
func (registry *Registry) PackChilledUntilRewardEpochId(beneficiary [20]byte) []byte {
	enc, err := registry.abi.Pack("chilledUntilRewardEpochId", beneficiary)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackChilledUntilRewardEpochId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3c5cb76f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function chilledUntilRewardEpochId(bytes20 _beneficiary) view returns(uint256 _rewardEpochId)
func (registry *Registry) TryPackChilledUntilRewardEpochId(beneficiary [20]byte) ([]byte, error) {
	return registry.abi.Pack("chilledUntilRewardEpochId", beneficiary)
}

// UnpackChilledUntilRewardEpochId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3c5cb76f.
//
// Solidity: function chilledUntilRewardEpochId(bytes20 _beneficiary) view returns(uint256 _rewardEpochId)
func (registry *Registry) UnpackChilledUntilRewardEpochId(data []byte) (*big.Int, error) {
	out, err := registry.abi.Unpack("chilledUntilRewardEpochId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetNumberOfRegisteredVoters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x369e9434.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNumberOfRegisteredVoters(uint256 _rewardEpochId) view returns(uint256)
func (registry *Registry) PackGetNumberOfRegisteredVoters(rewardEpochId *big.Int) []byte {
	enc, err := registry.abi.Pack("getNumberOfRegisteredVoters", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNumberOfRegisteredVoters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x369e9434.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNumberOfRegisteredVoters(uint256 _rewardEpochId) view returns(uint256)
func (registry *Registry) TryPackGetNumberOfRegisteredVoters(rewardEpochId *big.Int) ([]byte, error) {
	return registry.abi.Pack("getNumberOfRegisteredVoters", rewardEpochId)
}

// UnpackGetNumberOfRegisteredVoters is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x369e9434.
//
// Solidity: function getNumberOfRegisteredVoters(uint256 _rewardEpochId) view returns(uint256)
func (registry *Registry) UnpackGetNumberOfRegisteredVoters(data []byte) (*big.Int, error) {
	out, err := registry.abi.Unpack("getNumberOfRegisteredVoters", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetRegisteredVoters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x457c2e47.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRegisteredVoters(uint256 _rewardEpochId) view returns(address[])
func (registry *Registry) PackGetRegisteredVoters(rewardEpochId *big.Int) []byte {
	enc, err := registry.abi.Pack("getRegisteredVoters", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRegisteredVoters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x457c2e47.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRegisteredVoters(uint256 _rewardEpochId) view returns(address[])
func (registry *Registry) TryPackGetRegisteredVoters(rewardEpochId *big.Int) ([]byte, error) {
	return registry.abi.Pack("getRegisteredVoters", rewardEpochId)
}

// UnpackGetRegisteredVoters is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x457c2e47.
//
// Solidity: function getRegisteredVoters(uint256 _rewardEpochId) view returns(address[])
func (registry *Registry) UnpackGetRegisteredVoters(data []byte) ([]common.Address, error) {
	out, err := registry.abi.Unpack("getRegisteredVoters", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackIsVoterRegistered is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f5a9968.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isVoterRegistered(address _voter, uint256 _rewardEpochId) view returns(bool)
func (registry *Registry) PackIsVoterRegistered(voter common.Address, rewardEpochId *big.Int) []byte {
	enc, err := registry.abi.Pack("isVoterRegistered", voter, rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsVoterRegistered is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f5a9968.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isVoterRegistered(address _voter, uint256 _rewardEpochId) view returns(bool)
func (registry *Registry) TryPackIsVoterRegistered(voter common.Address, rewardEpochId *big.Int) ([]byte, error) {
	return registry.abi.Pack("isVoterRegistered", voter, rewardEpochId)
}

// UnpackIsVoterRegistered is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4f5a9968.
//
// Solidity: function isVoterRegistered(address _voter, uint256 _rewardEpochId) view returns(bool)
func (registry *Registry) UnpackIsVoterRegistered(data []byte) (bool, error) {
	out, err := registry.abi.Unpack("isVoterRegistered", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackMaxVoters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5e50a63.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function maxVoters() view returns(uint256)
func (registry *Registry) PackMaxVoters() []byte {
	enc, err := registry.abi.Pack("maxVoters")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMaxVoters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5e50a63.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function maxVoters() view returns(uint256)
func (registry *Registry) TryPackMaxVoters() ([]byte, error) {
	return registry.abi.Pack("maxVoters")
}

// UnpackMaxVoters is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd5e50a63.
//
// Solidity: function maxVoters() view returns(uint256)
func (registry *Registry) UnpackMaxVoters(data []byte) (*big.Int, error) {
	out, err := registry.abi.Unpack("maxVoters", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackNewSigningPolicyInitializationStartBlockNumber is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfff50753.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function newSigningPolicyInitializationStartBlockNumber(uint256 _rewardEpochId) view returns(uint256)
func (registry *Registry) PackNewSigningPolicyInitializationStartBlockNumber(rewardEpochId *big.Int) []byte {
	enc, err := registry.abi.Pack("newSigningPolicyInitializationStartBlockNumber", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNewSigningPolicyInitializationStartBlockNumber is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfff50753.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function newSigningPolicyInitializationStartBlockNumber(uint256 _rewardEpochId) view returns(uint256)
func (registry *Registry) TryPackNewSigningPolicyInitializationStartBlockNumber(rewardEpochId *big.Int) ([]byte, error) {
	return registry.abi.Pack("newSigningPolicyInitializationStartBlockNumber", rewardEpochId)
}

// UnpackNewSigningPolicyInitializationStartBlockNumber is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfff50753.
//
// Solidity: function newSigningPolicyInitializationStartBlockNumber(uint256 _rewardEpochId) view returns(uint256)
func (registry *Registry) UnpackNewSigningPolicyInitializationStartBlockNumber(data []byte) (*big.Int, error) {
	out, err := registry.abi.Unpack("newSigningPolicyInitializationStartBlockNumber", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackPublicKeyRequired is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x92e3e45f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function publicKeyRequired() view returns(bool)
func (registry *Registry) PackPublicKeyRequired() []byte {
	enc, err := registry.abi.Pack("publicKeyRequired")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPublicKeyRequired is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x92e3e45f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function publicKeyRequired() view returns(bool)
func (registry *Registry) TryPackPublicKeyRequired() ([]byte, error) {
	return registry.abi.Pack("publicKeyRequired")
}

// UnpackPublicKeyRequired is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x92e3e45f.
//
// Solidity: function publicKeyRequired() view returns(bool)
func (registry *Registry) UnpackPublicKeyRequired(data []byte) (bool, error) {
	out, err := registry.abi.Unpack("publicKeyRequired", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRegisterVoter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f7d0957.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function registerVoter(address _voter, (uint8,bytes32,bytes32) _signature) returns()
func (registry *Registry) PackRegisterVoter(voter common.Address, signature Signature) []byte {
	enc, err := registry.abi.Pack("registerVoter", voter, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegisterVoter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f7d0957.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function registerVoter(address _voter, (uint8,bytes32,bytes32) _signature) returns()
func (registry *Registry) TryPackRegisterVoter(voter common.Address, signature Signature) ([]byte, error) {
	return registry.abi.Pack("registerVoter", voter, signature)
}

// RegistryBeneficiaryChilled represents a BeneficiaryChilled event raised by the Registry contract.
type RegistryBeneficiaryChilled struct {
	Beneficiary        [20]byte
	UntilRewardEpochId uint32
	Raw                *types.Log // Blockchain specific contextual infos
}

const RegistryBeneficiaryChilledEventName = "BeneficiaryChilled"

// ContractEventName returns the user-defined event name.
func (RegistryBeneficiaryChilled) ContractEventName() string {
	return RegistryBeneficiaryChilledEventName
}

// UnpackBeneficiaryChilledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BeneficiaryChilled(bytes20 indexed beneficiary, uint32 untilRewardEpochId)
func (registry *Registry) UnpackBeneficiaryChilledEvent(log *types.Log) (*RegistryBeneficiaryChilled, error) {
	event := "BeneficiaryChilled"
	if len(log.Topics) == 0 || log.Topics[0] != registry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RegistryBeneficiaryChilled)
	if len(log.Data) > 0 {
		if err := registry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range registry.abi.Events[event].Inputs {
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

// RegistryVoterRegistered represents a VoterRegistered event raised by the Registry contract.
type RegistryVoterRegistered struct {
	Voter                   common.Address
	RewardEpochId           uint32
	SigningPolicyAddress    common.Address
	SubmitAddress           common.Address
	SubmitSignaturesAddress common.Address
	PublicKey               PublicKey
	RegistrationWeight      *big.Int
	Signature               Signature
	Raw                     *types.Log // Blockchain specific contextual infos
}

const RegistryVoterRegisteredEventName = "VoterRegistered"

// ContractEventName returns the user-defined event name.
func (RegistryVoterRegistered) ContractEventName() string {
	return RegistryVoterRegisteredEventName
}

// UnpackVoterRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoterRegistered(address indexed voter, uint32 indexed rewardEpochId, address indexed signingPolicyAddress, address submitAddress, address submitSignaturesAddress, (bytes32,bytes32) publicKey, uint256 registrationWeight, (uint8,bytes32,bytes32) signature)
func (registry *Registry) UnpackVoterRegisteredEvent(log *types.Log) (*RegistryVoterRegistered, error) {
	event := "VoterRegistered"
	if len(log.Topics) == 0 || log.Topics[0] != registry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RegistryVoterRegistered)
	if len(log.Data) > 0 {
		if err := registry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range registry.abi.Events[event].Inputs {
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

// RegistryVoterRemoved represents a VoterRemoved event raised by the Registry contract.
type RegistryVoterRemoved struct {
	Voter         common.Address
	RewardEpochId uint32
	Raw           *types.Log // Blockchain specific contextual infos
}

const RegistryVoterRemovedEventName = "VoterRemoved"

// ContractEventName returns the user-defined event name.
func (RegistryVoterRemoved) ContractEventName() string {
	return RegistryVoterRemovedEventName
}

// UnpackVoterRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VoterRemoved(address indexed voter, uint32 indexed rewardEpochId)
func (registry *Registry) UnpackVoterRemovedEvent(log *types.Log) (*RegistryVoterRemoved, error) {
	event := "VoterRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != registry.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RegistryVoterRemoved)
	if len(log.Data) > 0 {
		if err := registry.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range registry.abi.Events[event].Inputs {
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
