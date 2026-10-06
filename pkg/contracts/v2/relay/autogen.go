// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package relay

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

// IIRelaySigningPolicy is an auto generated low-level Go binding around an user-defined struct.
type IIRelaySigningPolicy struct {
	RewardEpochId      *big.Int
	StartVotingRoundId uint32
	Threshold          uint16
	Seed               *big.Int
	Voters             []common.Address
	Weights            []uint16
}

// RelayMetaData contains all meta data concerning the Relay contract.
var RelayMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signingPolicySetter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_initialRewardEpochId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_startingVotingRoundIdForInitialRewardEpochId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_initialSigningPolicyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"_randomNumberProtocolId\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_firstVotingRoundStartTs\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_votingEpochDurationSeconds\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_firstRewardEpochStartVotingRoundId\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_rewardEpochDurationInVotingEpochs\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"_thresholdIncreaseBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"_messageFinalizationWindowInRewardEpochs\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"protocolId\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"votingRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isSecureRandom\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"ProtocolMessageRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startVotingRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"threshold\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint16[]\",\"name\":\"weights\",\"type\":\"uint16[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signingPolicyBytes\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"SigningPolicyInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"SigningPolicyRelayed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_protocolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_votingRoundId\",\"type\":\"uint256\"}],\"name\":\"getConfirmedMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRandomNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_randomNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isSecureRandom\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_randomTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getVotingRoundId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInitializedRewardEpochData\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"_lastInitializedRewardEpoch\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_startingVotingRoundIdForLastInitializedRewardEpoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"protocolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingRoundId\",\"type\":\"uint256\"}],\"name\":\"merkleRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint32\",\"name\":\"startVotingRoundId\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"threshold\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"voters\",\"type\":\"address[]\"},{\"internalType\":\"uint16[]\",\"name\":\"weights\",\"type\":\"uint16[]\"}],\"internalType\":\"structIIRelay.SigningPolicy\",\"name\":\"_signingPolicy\",\"type\":\"tuple\"}],\"name\":\"setSigningPolicy\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingPolicySetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"startingVotingRoundIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stateData\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"randomNumberProtocolId\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"firstVotingRoundStartTs\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"votingEpochDurationSeconds\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"firstRewardEpochStartVotingRoundId\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"rewardEpochDurationInVotingEpochs\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"thresholdIncreaseBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"randomVotingRoundId\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isSecureRandom\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"lastInitializedRewardEpoch\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"noSigningPolicyRelay\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"messageFinalizationWindowInRewardEpochs\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"toSigningPolicyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "Relay",
}

// Relay is an auto generated Go binding around an Ethereum contract.
type Relay struct {
	abi abi.ABI
}

// NewRelay creates a new instance of Relay.
func NewRelay() *Relay {
	parsed, err := RelayMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &Relay{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *Relay) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _signingPolicySetter, uint32 _initialRewardEpochId, uint32 _startingVotingRoundIdForInitialRewardEpochId, bytes32 _initialSigningPolicyHash, uint8 _randomNumberProtocolId, uint32 _firstVotingRoundStartTs, uint8 _votingEpochDurationSeconds, uint32 _firstRewardEpochStartVotingRoundId, uint16 _rewardEpochDurationInVotingEpochs, uint16 _thresholdIncreaseBIPS, uint32 _messageFinalizationWindowInRewardEpochs) returns()
func (relay *Relay) PackConstructor(_signingPolicySetter common.Address, _initialRewardEpochId uint32, _startingVotingRoundIdForInitialRewardEpochId uint32, _initialSigningPolicyHash [32]byte, _randomNumberProtocolId uint8, _firstVotingRoundStartTs uint32, _votingEpochDurationSeconds uint8, _firstRewardEpochStartVotingRoundId uint32, _rewardEpochDurationInVotingEpochs uint16, _thresholdIncreaseBIPS uint16, _messageFinalizationWindowInRewardEpochs uint32) []byte {
	enc, err := relay.abi.Pack("", _signingPolicySetter, _initialRewardEpochId, _startingVotingRoundIdForInitialRewardEpochId, _initialSigningPolicyHash, _randomNumberProtocolId, _firstVotingRoundStartTs, _votingEpochDurationSeconds, _firstRewardEpochStartVotingRoundId, _rewardEpochDurationInVotingEpochs, _thresholdIncreaseBIPS, _messageFinalizationWindowInRewardEpochs)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetConfirmedMerkleRoot is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22c3f6fa.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getConfirmedMerkleRoot(uint256 _protocolId, uint256 _votingRoundId) view returns(bytes32)
func (relay *Relay) PackGetConfirmedMerkleRoot(protocolId *big.Int, votingRoundId *big.Int) []byte {
	enc, err := relay.abi.Pack("getConfirmedMerkleRoot", protocolId, votingRoundId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetConfirmedMerkleRoot is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22c3f6fa.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getConfirmedMerkleRoot(uint256 _protocolId, uint256 _votingRoundId) view returns(bytes32)
func (relay *Relay) TryPackGetConfirmedMerkleRoot(protocolId *big.Int, votingRoundId *big.Int) ([]byte, error) {
	return relay.abi.Pack("getConfirmedMerkleRoot", protocolId, votingRoundId)
}

// UnpackGetConfirmedMerkleRoot is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x22c3f6fa.
//
// Solidity: function getConfirmedMerkleRoot(uint256 _protocolId, uint256 _votingRoundId) view returns(bytes32)
func (relay *Relay) UnpackGetConfirmedMerkleRoot(data []byte) ([32]byte, error) {
	out, err := relay.abi.Unpack("getConfirmedMerkleRoot", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetRandomNumber is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdbdff2c1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRandomNumber() view returns(uint256 _randomNumber, bool _isSecureRandom, uint256 _randomTimestamp)
func (relay *Relay) PackGetRandomNumber() []byte {
	enc, err := relay.abi.Pack("getRandomNumber")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRandomNumber is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdbdff2c1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRandomNumber() view returns(uint256 _randomNumber, bool _isSecureRandom, uint256 _randomTimestamp)
func (relay *Relay) TryPackGetRandomNumber() ([]byte, error) {
	return relay.abi.Pack("getRandomNumber")
}

// GetRandomNumberOutput serves as a container for the return parameters of contract
// method GetRandomNumber.
type GetRandomNumberOutput struct {
	RandomNumber    *big.Int
	IsSecureRandom  bool
	RandomTimestamp *big.Int
}

// UnpackGetRandomNumber is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdbdff2c1.
//
// Solidity: function getRandomNumber() view returns(uint256 _randomNumber, bool _isSecureRandom, uint256 _randomTimestamp)
func (relay *Relay) UnpackGetRandomNumber(data []byte) (GetRandomNumberOutput, error) {
	out, err := relay.abi.Unpack("getRandomNumber", data)
	outstruct := new(GetRandomNumberOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.RandomNumber = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.IsSecureRandom = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RandomTimestamp = abi.ConvertType(out[2], new(big.Int)).(*big.Int)
	return *outstruct, nil
}

// PackGetVotingRoundId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab97db37.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVotingRoundId(uint256 _timestamp) view returns(uint256)
func (relay *Relay) PackGetVotingRoundId(timestamp *big.Int) []byte {
	enc, err := relay.abi.Pack("getVotingRoundId", timestamp)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVotingRoundId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab97db37.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVotingRoundId(uint256 _timestamp) view returns(uint256)
func (relay *Relay) TryPackGetVotingRoundId(timestamp *big.Int) ([]byte, error) {
	return relay.abi.Pack("getVotingRoundId", timestamp)
}

// UnpackGetVotingRoundId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xab97db37.
//
// Solidity: function getVotingRoundId(uint256 _timestamp) view returns(uint256)
func (relay *Relay) UnpackGetVotingRoundId(data []byte) (*big.Int, error) {
	out, err := relay.abi.Unpack("getVotingRoundId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackLastInitializedRewardEpochData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8af0c307.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function lastInitializedRewardEpochData() view returns(uint32 _lastInitializedRewardEpoch, uint32 _startingVotingRoundIdForLastInitializedRewardEpoch)
func (relay *Relay) PackLastInitializedRewardEpochData() []byte {
	enc, err := relay.abi.Pack("lastInitializedRewardEpochData")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLastInitializedRewardEpochData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8af0c307.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function lastInitializedRewardEpochData() view returns(uint32 _lastInitializedRewardEpoch, uint32 _startingVotingRoundIdForLastInitializedRewardEpoch)
func (relay *Relay) TryPackLastInitializedRewardEpochData() ([]byte, error) {
	return relay.abi.Pack("lastInitializedRewardEpochData")
}

// LastInitializedRewardEpochDataOutput serves as a container for the return parameters of contract
// method LastInitializedRewardEpochData.
type LastInitializedRewardEpochDataOutput struct {
	LastInitializedRewardEpoch                         uint32
	StartingVotingRoundIdForLastInitializedRewardEpoch uint32
}

// UnpackLastInitializedRewardEpochData is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8af0c307.
//
// Solidity: function lastInitializedRewardEpochData() view returns(uint32 _lastInitializedRewardEpoch, uint32 _startingVotingRoundIdForLastInitializedRewardEpoch)
func (relay *Relay) UnpackLastInitializedRewardEpochData(data []byte) (LastInitializedRewardEpochDataOutput, error) {
	out, err := relay.abi.Unpack("lastInitializedRewardEpochData", data)
	outstruct := new(LastInitializedRewardEpochDataOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.LastInitializedRewardEpoch = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.StartingVotingRoundIdForLastInitializedRewardEpoch = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	return *outstruct, nil
}

// PackMerkleRoots is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x39436b00.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function merkleRoots(uint256 protocolId, uint256 votingRoundId) view returns(bytes32)
func (relay *Relay) PackMerkleRoots(protocolId *big.Int, votingRoundId *big.Int) []byte {
	enc, err := relay.abi.Pack("merkleRoots", protocolId, votingRoundId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMerkleRoots is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x39436b00.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function merkleRoots(uint256 protocolId, uint256 votingRoundId) view returns(bytes32)
func (relay *Relay) TryPackMerkleRoots(protocolId *big.Int, votingRoundId *big.Int) ([]byte, error) {
	return relay.abi.Pack("merkleRoots", protocolId, votingRoundId)
}

// UnpackMerkleRoots is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x39436b00.
//
// Solidity: function merkleRoots(uint256 protocolId, uint256 votingRoundId) view returns(bytes32)
func (relay *Relay) UnpackMerkleRoots(data []byte) ([32]byte, error) {
	out, err := relay.abi.Unpack("merkleRoots", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackRelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb59589d1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function relay() returns()
func (relay *Relay) PackRelay() []byte {
	enc, err := relay.abi.Pack("relay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb59589d1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function relay() returns()
func (relay *Relay) TryPackRelay() ([]byte, error) {
	return relay.abi.Pack("relay")
}

// PackSetSigningPolicy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83534125.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSigningPolicy((uint24,uint32,uint16,uint256,address[],uint16[]) _signingPolicy) returns(bytes32)
func (relay *Relay) PackSetSigningPolicy(signingPolicy IIRelaySigningPolicy) []byte {
	enc, err := relay.abi.Pack("setSigningPolicy", signingPolicy)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSigningPolicy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x83534125.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSigningPolicy((uint24,uint32,uint16,uint256,address[],uint16[]) _signingPolicy) returns(bytes32)
func (relay *Relay) TryPackSetSigningPolicy(signingPolicy IIRelaySigningPolicy) ([]byte, error) {
	return relay.abi.Pack("setSigningPolicy", signingPolicy)
}

// UnpackSetSigningPolicy is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x83534125.
//
// Solidity: function setSigningPolicy((uint24,uint32,uint16,uint256,address[],uint16[]) _signingPolicy) returns(bytes32)
func (relay *Relay) UnpackSetSigningPolicy(data []byte) ([32]byte, error) {
	out, err := relay.abi.Unpack("setSigningPolicy", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackSigningPolicySetter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9dbe8ed.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signingPolicySetter() view returns(address)
func (relay *Relay) PackSigningPolicySetter() []byte {
	enc, err := relay.abi.Pack("signingPolicySetter")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSigningPolicySetter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa9dbe8ed.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signingPolicySetter() view returns(address)
func (relay *Relay) TryPackSigningPolicySetter() ([]byte, error) {
	return relay.abi.Pack("signingPolicySetter")
}

// UnpackSigningPolicySetter is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa9dbe8ed.
//
// Solidity: function signingPolicySetter() view returns(address)
func (relay *Relay) UnpackSigningPolicySetter(data []byte) (common.Address, error) {
	out, err := relay.abi.Unpack("signingPolicySetter", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackStartingVotingRoundIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7297c0a2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function startingVotingRoundIds(uint256 rewardEpochId) view returns(uint256)
func (relay *Relay) PackStartingVotingRoundIds(rewardEpochId *big.Int) []byte {
	enc, err := relay.abi.Pack("startingVotingRoundIds", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackStartingVotingRoundIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7297c0a2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function startingVotingRoundIds(uint256 rewardEpochId) view returns(uint256)
func (relay *Relay) TryPackStartingVotingRoundIds(rewardEpochId *big.Int) ([]byte, error) {
	return relay.abi.Pack("startingVotingRoundIds", rewardEpochId)
}

// UnpackStartingVotingRoundIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7297c0a2.
//
// Solidity: function startingVotingRoundIds(uint256 rewardEpochId) view returns(uint256)
func (relay *Relay) UnpackStartingVotingRoundIds(data []byte) (*big.Int, error) {
	out, err := relay.abi.Unpack("startingVotingRoundIds", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackStateData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e8fb36a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function stateData() view returns(uint8 randomNumberProtocolId, uint32 firstVotingRoundStartTs, uint8 votingEpochDurationSeconds, uint32 firstRewardEpochStartVotingRoundId, uint16 rewardEpochDurationInVotingEpochs, uint16 thresholdIncreaseBIPS, uint32 randomVotingRoundId, bool isSecureRandom, uint32 lastInitializedRewardEpoch, bool noSigningPolicyRelay, uint32 messageFinalizationWindowInRewardEpochs)
func (relay *Relay) PackStateData() []byte {
	enc, err := relay.abi.Pack("stateData")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackStateData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e8fb36a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function stateData() view returns(uint8 randomNumberProtocolId, uint32 firstVotingRoundStartTs, uint8 votingEpochDurationSeconds, uint32 firstRewardEpochStartVotingRoundId, uint16 rewardEpochDurationInVotingEpochs, uint16 thresholdIncreaseBIPS, uint32 randomVotingRoundId, bool isSecureRandom, uint32 lastInitializedRewardEpoch, bool noSigningPolicyRelay, uint32 messageFinalizationWindowInRewardEpochs)
func (relay *Relay) TryPackStateData() ([]byte, error) {
	return relay.abi.Pack("stateData")
}

// StateDataOutput serves as a container for the return parameters of contract
// method StateData.
type StateDataOutput struct {
	RandomNumberProtocolId                  uint8
	FirstVotingRoundStartTs                 uint32
	VotingEpochDurationSeconds              uint8
	FirstRewardEpochStartVotingRoundId      uint32
	RewardEpochDurationInVotingEpochs       uint16
	ThresholdIncreaseBIPS                   uint16
	RandomVotingRoundId                     uint32
	IsSecureRandom                          bool
	LastInitializedRewardEpoch              uint32
	NoSigningPolicyRelay                    bool
	MessageFinalizationWindowInRewardEpochs uint32
}

// UnpackStateData is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1e8fb36a.
//
// Solidity: function stateData() view returns(uint8 randomNumberProtocolId, uint32 firstVotingRoundStartTs, uint8 votingEpochDurationSeconds, uint32 firstRewardEpochStartVotingRoundId, uint16 rewardEpochDurationInVotingEpochs, uint16 thresholdIncreaseBIPS, uint32 randomVotingRoundId, bool isSecureRandom, uint32 lastInitializedRewardEpoch, bool noSigningPolicyRelay, uint32 messageFinalizationWindowInRewardEpochs)
func (relay *Relay) UnpackStateData(data []byte) (StateDataOutput, error) {
	out, err := relay.abi.Unpack("stateData", data)
	outstruct := new(StateDataOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.RandomNumberProtocolId = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.FirstVotingRoundStartTs = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.VotingEpochDurationSeconds = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.FirstRewardEpochStartVotingRoundId = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.RewardEpochDurationInVotingEpochs = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.ThresholdIncreaseBIPS = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.RandomVotingRoundId = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.IsSecureRandom = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.LastInitializedRewardEpoch = *abi.ConvertType(out[8], new(uint32)).(*uint32)
	outstruct.NoSigningPolicyRelay = *abi.ConvertType(out[9], new(bool)).(*bool)
	outstruct.MessageFinalizationWindowInRewardEpochs = *abi.ConvertType(out[10], new(uint32)).(*uint32)
	return *outstruct, nil
}

// PackToSigningPolicyHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0c85bf07.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function toSigningPolicyHash(uint256 rewardEpochId) view returns(bytes32)
func (relay *Relay) PackToSigningPolicyHash(rewardEpochId *big.Int) []byte {
	enc, err := relay.abi.Pack("toSigningPolicyHash", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackToSigningPolicyHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0c85bf07.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function toSigningPolicyHash(uint256 rewardEpochId) view returns(bytes32)
func (relay *Relay) TryPackToSigningPolicyHash(rewardEpochId *big.Int) ([]byte, error) {
	return relay.abi.Pack("toSigningPolicyHash", rewardEpochId)
}

// UnpackToSigningPolicyHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0c85bf07.
//
// Solidity: function toSigningPolicyHash(uint256 rewardEpochId) view returns(bytes32)
func (relay *Relay) UnpackToSigningPolicyHash(data []byte) ([32]byte, error) {
	out, err := relay.abi.Unpack("toSigningPolicyHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// RelayProtocolMessageRelayed represents a ProtocolMessageRelayed event raised by the Relay contract.
type RelayProtocolMessageRelayed struct {
	ProtocolId     uint8
	VotingRoundId  uint32
	IsSecureRandom bool
	MerkleRoot     [32]byte
	Raw            *types.Log // Blockchain specific contextual infos
}

const RelayProtocolMessageRelayedEventName = "ProtocolMessageRelayed"

// ContractEventName returns the user-defined event name.
func (RelayProtocolMessageRelayed) ContractEventName() string {
	return RelayProtocolMessageRelayedEventName
}

// UnpackProtocolMessageRelayedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProtocolMessageRelayed(uint8 indexed protocolId, uint32 indexed votingRoundId, bool isSecureRandom, bytes32 merkleRoot)
func (relay *Relay) UnpackProtocolMessageRelayedEvent(log *types.Log) (*RelayProtocolMessageRelayed, error) {
	event := "ProtocolMessageRelayed"
	if len(log.Topics) == 0 || log.Topics[0] != relay.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RelayProtocolMessageRelayed)
	if len(log.Data) > 0 {
		if err := relay.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range relay.abi.Events[event].Inputs {
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

// RelaySigningPolicyInitialized represents a SigningPolicyInitialized event raised by the Relay contract.
type RelaySigningPolicyInitialized struct {
	RewardEpochId      *big.Int
	StartVotingRoundId uint32
	Threshold          uint16
	Seed               *big.Int
	Voters             []common.Address
	Weights            []uint16
	SigningPolicyBytes []byte
	Timestamp          uint64
	Raw                *types.Log // Blockchain specific contextual infos
}

const RelaySigningPolicyInitializedEventName = "SigningPolicyInitialized"

// ContractEventName returns the user-defined event name.
func (RelaySigningPolicyInitialized) ContractEventName() string {
	return RelaySigningPolicyInitializedEventName
}

// UnpackSigningPolicyInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SigningPolicyInitialized(uint24 indexed rewardEpochId, uint32 startVotingRoundId, uint16 threshold, uint256 seed, address[] voters, uint16[] weights, bytes signingPolicyBytes, uint64 timestamp)
func (relay *Relay) UnpackSigningPolicyInitializedEvent(log *types.Log) (*RelaySigningPolicyInitialized, error) {
	event := "SigningPolicyInitialized"
	if len(log.Topics) == 0 || log.Topics[0] != relay.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RelaySigningPolicyInitialized)
	if len(log.Data) > 0 {
		if err := relay.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range relay.abi.Events[event].Inputs {
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

// RelaySigningPolicyRelayed represents a SigningPolicyRelayed event raised by the Relay contract.
type RelaySigningPolicyRelayed struct {
	RewardEpochId *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const RelaySigningPolicyRelayedEventName = "SigningPolicyRelayed"

// ContractEventName returns the user-defined event name.
func (RelaySigningPolicyRelayed) ContractEventName() string {
	return RelaySigningPolicyRelayedEventName
}

// UnpackSigningPolicyRelayedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SigningPolicyRelayed(uint256 indexed rewardEpochId)
func (relay *Relay) UnpackSigningPolicyRelayedEvent(log *types.Log) (*RelaySigningPolicyRelayed, error) {
	event := "SigningPolicyRelayed"
	if len(log.Topics) == 0 || log.Topics[0] != relay.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(RelaySigningPolicyRelayed)
	if len(log.Data) > 0 {
		if err := relay.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range relay.abi.Events[event].Inputs {
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
