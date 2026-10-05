// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package extensiongovernance

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

// ExtensionGovernanceMetaData contains all meta data concerning the ExtensionGovernance contract.
var ExtensionGovernanceMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSigners\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeDomainSeparatorMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"governanceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"signersThreshold\",\"type\":\"uint64\"}],\"name\":\"NewTeeGovernanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"governanceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"safe\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"signersThreshold\",\"type\":\"uint64\"}],\"name\":\"NewTeeSafeGovernanceSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getLatestTeeGovernance\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_signersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_safe\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getLatestTeeGovernanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"getTeeGovernance\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_signersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_safe\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"getTeeGovernanceThreshold\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"isGovernanceHashValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"isTeeGovernanceSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_signersThreshold\",\"type\":\"uint64\"}],\"name\":\"setNewTeeGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_safe\",\"type\":\"address\"}],\"name\":\"setNewTeeGovernanceSafe\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "ExtensionGovernance",
}

// ExtensionGovernance is an auto generated Go binding around an Ethereum contract.
type ExtensionGovernance struct {
	abi abi.ABI
}

// NewExtensionGovernance creates a new instance of ExtensionGovernance.
func NewExtensionGovernance() *ExtensionGovernance {
	parsed, err := ExtensionGovernanceMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ExtensionGovernance{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ExtensionGovernance) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackGetLatestTeeGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e7bd5d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLatestTeeGovernance(uint256 _extensionId) view returns(address[] _signers, uint64 _signersThreshold, address _safe)
func (extensionGovernance *ExtensionGovernance) PackGetLatestTeeGovernance(extensionId *big.Int) []byte {
	enc, err := extensionGovernance.abi.Pack("getLatestTeeGovernance", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLatestTeeGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e7bd5d4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLatestTeeGovernance(uint256 _extensionId) view returns(address[] _signers, uint64 _signersThreshold, address _safe)
func (extensionGovernance *ExtensionGovernance) TryPackGetLatestTeeGovernance(extensionId *big.Int) ([]byte, error) {
	return extensionGovernance.abi.Pack("getLatestTeeGovernance", extensionId)
}

// GetLatestTeeGovernanceOutput serves as a container for the return parameters of contract
// method GetLatestTeeGovernance.
type GetLatestTeeGovernanceOutput struct {
	Signers          []common.Address
	SignersThreshold uint64
	Safe             common.Address
}

// UnpackGetLatestTeeGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e7bd5d4.
//
// Solidity: function getLatestTeeGovernance(uint256 _extensionId) view returns(address[] _signers, uint64 _signersThreshold, address _safe)
func (extensionGovernance *ExtensionGovernance) UnpackGetLatestTeeGovernance(data []byte) (GetLatestTeeGovernanceOutput, error) {
	out, err := extensionGovernance.abi.Unpack("getLatestTeeGovernance", data)
	outstruct := new(GetLatestTeeGovernanceOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Signers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.SignersThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Safe = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackGetLatestTeeGovernanceHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x555f9961.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLatestTeeGovernanceHash(uint256 _extensionId) view returns(bytes32)
func (extensionGovernance *ExtensionGovernance) PackGetLatestTeeGovernanceHash(extensionId *big.Int) []byte {
	enc, err := extensionGovernance.abi.Pack("getLatestTeeGovernanceHash", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLatestTeeGovernanceHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x555f9961.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLatestTeeGovernanceHash(uint256 _extensionId) view returns(bytes32)
func (extensionGovernance *ExtensionGovernance) TryPackGetLatestTeeGovernanceHash(extensionId *big.Int) ([]byte, error) {
	return extensionGovernance.abi.Pack("getLatestTeeGovernanceHash", extensionId)
}

// UnpackGetLatestTeeGovernanceHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x555f9961.
//
// Solidity: function getLatestTeeGovernanceHash(uint256 _extensionId) view returns(bytes32)
func (extensionGovernance *ExtensionGovernance) UnpackGetLatestTeeGovernanceHash(data []byte) ([32]byte, error) {
	out, err := extensionGovernance.abi.Unpack("getLatestTeeGovernanceHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetTeeGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x090f2bfb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTeeGovernance(uint256 _extensionId, bytes32 _governanceHash) view returns(address[] _signers, uint64 _signersThreshold, address _safe)
func (extensionGovernance *ExtensionGovernance) PackGetTeeGovernance(extensionId *big.Int, governanceHash [32]byte) []byte {
	enc, err := extensionGovernance.abi.Pack("getTeeGovernance", extensionId, governanceHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTeeGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x090f2bfb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTeeGovernance(uint256 _extensionId, bytes32 _governanceHash) view returns(address[] _signers, uint64 _signersThreshold, address _safe)
func (extensionGovernance *ExtensionGovernance) TryPackGetTeeGovernance(extensionId *big.Int, governanceHash [32]byte) ([]byte, error) {
	return extensionGovernance.abi.Pack("getTeeGovernance", extensionId, governanceHash)
}

// GetTeeGovernanceOutput serves as a container for the return parameters of contract
// method GetTeeGovernance.
type GetTeeGovernanceOutput struct {
	Signers          []common.Address
	SignersThreshold uint64
	Safe             common.Address
}

// UnpackGetTeeGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x090f2bfb.
//
// Solidity: function getTeeGovernance(uint256 _extensionId, bytes32 _governanceHash) view returns(address[] _signers, uint64 _signersThreshold, address _safe)
func (extensionGovernance *ExtensionGovernance) UnpackGetTeeGovernance(data []byte) (GetTeeGovernanceOutput, error) {
	out, err := extensionGovernance.abi.Unpack("getTeeGovernance", data)
	outstruct := new(GetTeeGovernanceOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Signers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.SignersThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Safe = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	return *outstruct, nil
}

// PackGetTeeGovernanceThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3b13db2b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTeeGovernanceThreshold(uint256 _extensionId, bytes32 _governanceHash) view returns(uint64)
func (extensionGovernance *ExtensionGovernance) PackGetTeeGovernanceThreshold(extensionId *big.Int, governanceHash [32]byte) []byte {
	enc, err := extensionGovernance.abi.Pack("getTeeGovernanceThreshold", extensionId, governanceHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTeeGovernanceThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3b13db2b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTeeGovernanceThreshold(uint256 _extensionId, bytes32 _governanceHash) view returns(uint64)
func (extensionGovernance *ExtensionGovernance) TryPackGetTeeGovernanceThreshold(extensionId *big.Int, governanceHash [32]byte) ([]byte, error) {
	return extensionGovernance.abi.Pack("getTeeGovernanceThreshold", extensionId, governanceHash)
}

// UnpackGetTeeGovernanceThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3b13db2b.
//
// Solidity: function getTeeGovernanceThreshold(uint256 _extensionId, bytes32 _governanceHash) view returns(uint64)
func (extensionGovernance *ExtensionGovernance) UnpackGetTeeGovernanceThreshold(data []byte) (uint64, error) {
	out, err := extensionGovernance.abi.Unpack("getTeeGovernanceThreshold", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackIsGovernanceHashValid is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66d98ef8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isGovernanceHashValid(uint256 _extensionId, bytes32 _governanceHash) view returns(bool)
func (extensionGovernance *ExtensionGovernance) PackIsGovernanceHashValid(extensionId *big.Int, governanceHash [32]byte) []byte {
	enc, err := extensionGovernance.abi.Pack("isGovernanceHashValid", extensionId, governanceHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsGovernanceHashValid is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x66d98ef8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isGovernanceHashValid(uint256 _extensionId, bytes32 _governanceHash) view returns(bool)
func (extensionGovernance *ExtensionGovernance) TryPackIsGovernanceHashValid(extensionId *big.Int, governanceHash [32]byte) ([]byte, error) {
	return extensionGovernance.abi.Pack("isGovernanceHashValid", extensionId, governanceHash)
}

// UnpackIsGovernanceHashValid is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x66d98ef8.
//
// Solidity: function isGovernanceHashValid(uint256 _extensionId, bytes32 _governanceHash) view returns(bool)
func (extensionGovernance *ExtensionGovernance) UnpackIsGovernanceHashValid(data []byte) (bool, error) {
	out, err := extensionGovernance.abi.Unpack("isGovernanceHashValid", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsTeeGovernanceSigner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd8784372.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isTeeGovernanceSigner(uint256 _extensionId, bytes32 _governanceHash, address _signer) view returns(bool)
func (extensionGovernance *ExtensionGovernance) PackIsTeeGovernanceSigner(extensionId *big.Int, governanceHash [32]byte, signer common.Address) []byte {
	enc, err := extensionGovernance.abi.Pack("isTeeGovernanceSigner", extensionId, governanceHash, signer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsTeeGovernanceSigner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd8784372.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isTeeGovernanceSigner(uint256 _extensionId, bytes32 _governanceHash, address _signer) view returns(bool)
func (extensionGovernance *ExtensionGovernance) TryPackIsTeeGovernanceSigner(extensionId *big.Int, governanceHash [32]byte, signer common.Address) ([]byte, error) {
	return extensionGovernance.abi.Pack("isTeeGovernanceSigner", extensionId, governanceHash, signer)
}

// UnpackIsTeeGovernanceSigner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd8784372.
//
// Solidity: function isTeeGovernanceSigner(uint256 _extensionId, bytes32 _governanceHash, address _signer) view returns(bool)
func (extensionGovernance *ExtensionGovernance) UnpackIsTeeGovernanceSigner(data []byte) (bool, error) {
	out, err := extensionGovernance.abi.Unpack("isTeeGovernanceSigner", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSetNewTeeGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49e43748.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setNewTeeGovernance(uint256 _extensionId, address[] _signers, uint64 _signersThreshold) returns()
func (extensionGovernance *ExtensionGovernance) PackSetNewTeeGovernance(extensionId *big.Int, signers []common.Address, signersThreshold uint64) []byte {
	enc, err := extensionGovernance.abi.Pack("setNewTeeGovernance", extensionId, signers, signersThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetNewTeeGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x49e43748.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setNewTeeGovernance(uint256 _extensionId, address[] _signers, uint64 _signersThreshold) returns()
func (extensionGovernance *ExtensionGovernance) TryPackSetNewTeeGovernance(extensionId *big.Int, signers []common.Address, signersThreshold uint64) ([]byte, error) {
	return extensionGovernance.abi.Pack("setNewTeeGovernance", extensionId, signers, signersThreshold)
}

// PackSetNewTeeGovernanceSafe is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e8d6fd5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setNewTeeGovernanceSafe(uint256 _extensionId, address _safe) returns()
func (extensionGovernance *ExtensionGovernance) PackSetNewTeeGovernanceSafe(extensionId *big.Int, safe common.Address) []byte {
	enc, err := extensionGovernance.abi.Pack("setNewTeeGovernanceSafe", extensionId, safe)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetNewTeeGovernanceSafe is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0e8d6fd5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setNewTeeGovernanceSafe(uint256 _extensionId, address _safe) returns()
func (extensionGovernance *ExtensionGovernance) TryPackSetNewTeeGovernanceSafe(extensionId *big.Int, safe common.Address) ([]byte, error) {
	return extensionGovernance.abi.Pack("setNewTeeGovernanceSafe", extensionId, safe)
}

// ExtensionGovernanceNewTeeGovernanceSet represents a NewTeeGovernanceSet event raised by the ExtensionGovernance contract.
type ExtensionGovernanceNewTeeGovernanceSet struct {
	ExtensionId      *big.Int
	GovernanceHash   [32]byte
	Signers          []common.Address
	SignersThreshold uint64
	Raw              *types.Log // Blockchain specific contextual infos
}

const ExtensionGovernanceNewTeeGovernanceSetEventName = "NewTeeGovernanceSet"

// ContractEventName returns the user-defined event name.
func (ExtensionGovernanceNewTeeGovernanceSet) ContractEventName() string {
	return ExtensionGovernanceNewTeeGovernanceSetEventName
}

// UnpackNewTeeGovernanceSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewTeeGovernanceSet(uint256 indexed extensionId, bytes32 indexed governanceHash, address[] signers, uint64 signersThreshold)
func (extensionGovernance *ExtensionGovernance) UnpackNewTeeGovernanceSetEvent(log *types.Log) (*ExtensionGovernanceNewTeeGovernanceSet, error) {
	event := "NewTeeGovernanceSet"
	if len(log.Topics) == 0 || log.Topics[0] != extensionGovernance.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionGovernanceNewTeeGovernanceSet)
	if len(log.Data) > 0 {
		if err := extensionGovernance.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionGovernance.abi.Events[event].Inputs {
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

// ExtensionGovernanceNewTeeSafeGovernanceSet represents a NewTeeSafeGovernanceSet event raised by the ExtensionGovernance contract.
type ExtensionGovernanceNewTeeSafeGovernanceSet struct {
	ExtensionId      *big.Int
	GovernanceHash   [32]byte
	Safe             common.Address
	Signers          []common.Address
	SignersThreshold uint64
	Raw              *types.Log // Blockchain specific contextual infos
}

const ExtensionGovernanceNewTeeSafeGovernanceSetEventName = "NewTeeSafeGovernanceSet"

// ContractEventName returns the user-defined event name.
func (ExtensionGovernanceNewTeeSafeGovernanceSet) ContractEventName() string {
	return ExtensionGovernanceNewTeeSafeGovernanceSetEventName
}

// UnpackNewTeeSafeGovernanceSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewTeeSafeGovernanceSet(uint256 indexed extensionId, bytes32 indexed governanceHash, address safe, address[] signers, uint64 signersThreshold)
func (extensionGovernance *ExtensionGovernance) UnpackNewTeeSafeGovernanceSetEvent(log *types.Log) (*ExtensionGovernanceNewTeeSafeGovernanceSet, error) {
	event := "NewTeeSafeGovernanceSet"
	if len(log.Topics) == 0 || log.Topics[0] != extensionGovernance.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionGovernanceNewTeeSafeGovernanceSet)
	if len(log.Data) > 0 {
		if err := extensionGovernance.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionGovernance.abi.Events[event].Inputs {
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
func (extensionGovernance *ExtensionGovernance) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["AddressAlreadyInSet"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackAddressAlreadyInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["AddressNotInSet"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackAddressNotInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["AvailabilityCheckTimestampInvalid"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackAvailabilityCheckTimestampInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["DuplicatedCosigner"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackDuplicatedCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["InvalidAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackInvalidAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["InvalidKeyType"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackInvalidKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["InvalidPublicKey"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackInvalidPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["InvalidResponseData"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackInvalidResponseDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["InvalidSigner"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackInvalidSignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["InvalidSigningAlgo"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackInvalidSigningAlgoError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["InvalidWalletStatus"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackInvalidWalletStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["KeyTypeNotSupported"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["NoAddresses"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackNoAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["NoSigners"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackNoSignersError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["NotOwnerOrPauser"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackNotOwnerOrPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["NotOwnerOrUnpauser"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackNotOwnerOrUnpauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["OnlyExtensionOwner"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackOnlyExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackOnlyOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["OnlyOwnerOrBackupManager"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackOnlyOwnerOrBackupManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["OnlyProposedOwner"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackOnlyProposedOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["OwnerNotAllowed"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackOwnerNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["SafeDomainSeparatorMismatch"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackSafeDomainSeparatorMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["SignerAlreadyExists"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackSignerAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionGovernance.abi.Errors["VersionNotSupported"].ID.Bytes()[:4]) {
		return extensionGovernance.UnpackVersionNotSupportedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ExtensionGovernanceAddressAlreadyInSet represents a AddressAlreadyInSet error raised by the ExtensionGovernance contract.
type ExtensionGovernanceAddressAlreadyInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressAlreadyInSet(address addr)
func ExtensionGovernanceAddressAlreadyInSetErrorID() common.Hash {
	return common.HexToHash("0xb725c7fb1c3f0f065973d6b9eb55140f970259955af3045742cd5b37902eb920")
}

// UnpackAddressAlreadyInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressAlreadyInSet(address addr)
func (extensionGovernance *ExtensionGovernance) UnpackAddressAlreadyInSetError(raw []byte) (*ExtensionGovernanceAddressAlreadyInSet, error) {
	out := new(ExtensionGovernanceAddressAlreadyInSet)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "AddressAlreadyInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceAddressNotInSet represents a AddressNotInSet error raised by the ExtensionGovernance contract.
type ExtensionGovernanceAddressNotInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressNotInSet(address addr)
func ExtensionGovernanceAddressNotInSetErrorID() common.Hash {
	return common.HexToHash("0x1ff5ecf3e6d385ea01ff0323fe894ed6d25a83ed5f6ff2a2ede4c897408f6be9")
}

// UnpackAddressNotInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressNotInSet(address addr)
func (extensionGovernance *ExtensionGovernance) UnpackAddressNotInSetError(raw []byte) (*ExtensionGovernanceAddressNotInSet, error) {
	out := new(ExtensionGovernanceAddressNotInSet)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "AddressNotInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceAvailabilityCheckTimestampInvalid represents a AvailabilityCheckTimestampInvalid error raised by the ExtensionGovernance contract.
type ExtensionGovernanceAvailabilityCheckTimestampInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func ExtensionGovernanceAvailabilityCheckTimestampInvalidErrorID() common.Hash {
	return common.HexToHash("0x6344a5e74cff78612fa271fa907ca608af3bad10e1e8c8df302e29ebc6298e2d")
}

// UnpackAvailabilityCheckTimestampInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func (extensionGovernance *ExtensionGovernance) UnpackAvailabilityCheckTimestampInvalidError(raw []byte) (*ExtensionGovernanceAvailabilityCheckTimestampInvalid, error) {
	out := new(ExtensionGovernanceAvailabilityCheckTimestampInvalid)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "AvailabilityCheckTimestampInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceDuplicatedCosigner represents a DuplicatedCosigner error raised by the ExtensionGovernance contract.
type ExtensionGovernanceDuplicatedCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func ExtensionGovernanceDuplicatedCosignerErrorID() common.Hash {
	return common.HexToHash("0x53a5890417576fb3043dc3ff1669b0932da2f657d8adb4c11313b12b711f5910")
}

// UnpackDuplicatedCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func (extensionGovernance *ExtensionGovernance) UnpackDuplicatedCosignerError(raw []byte) (*ExtensionGovernanceDuplicatedCosigner, error) {
	out := new(ExtensionGovernanceDuplicatedCosigner)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "DuplicatedCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceExtensionIdMismatch represents a ExtensionIdMismatch error raised by the ExtensionGovernance contract.
type ExtensionGovernanceExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func ExtensionGovernanceExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (extensionGovernance *ExtensionGovernance) UnpackExtensionIdMismatchError(raw []byte) (*ExtensionGovernanceExtensionIdMismatch, error) {
	out := new(ExtensionGovernanceExtensionIdMismatch)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceInvalidAddress represents a InvalidAddress error raised by the ExtensionGovernance contract.
type ExtensionGovernanceInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func ExtensionGovernanceInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (extensionGovernance *ExtensionGovernance) UnpackInvalidAddressError(raw []byte) (*ExtensionGovernanceInvalidAddress, error) {
	out := new(ExtensionGovernanceInvalidAddress)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceInvalidAvailabilityCheckStatus represents a InvalidAvailabilityCheckStatus error raised by the ExtensionGovernance contract.
type ExtensionGovernanceInvalidAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func ExtensionGovernanceInvalidAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0xedee4ff00b842d3ef84031c60001065c540c838c5479bbf0ae56f15868427e44")
}

// UnpackInvalidAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func (extensionGovernance *ExtensionGovernance) UnpackInvalidAvailabilityCheckStatusError(raw []byte) (*ExtensionGovernanceInvalidAvailabilityCheckStatus, error) {
	out := new(ExtensionGovernanceInvalidAvailabilityCheckStatus)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "InvalidAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceInvalidCosigner represents a InvalidCosigner error raised by the ExtensionGovernance contract.
type ExtensionGovernanceInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func ExtensionGovernanceInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (extensionGovernance *ExtensionGovernance) UnpackInvalidCosignerError(raw []byte) (*ExtensionGovernanceInvalidCosigner, error) {
	out := new(ExtensionGovernanceInvalidCosigner)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceInvalidDuration represents a InvalidDuration error raised by the ExtensionGovernance contract.
type ExtensionGovernanceInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func ExtensionGovernanceInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (extensionGovernance *ExtensionGovernance) UnpackInvalidDurationError(raw []byte) (*ExtensionGovernanceInvalidDuration, error) {
	out := new(ExtensionGovernanceInvalidDuration)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the ExtensionGovernance contract.
type ExtensionGovernanceInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func ExtensionGovernanceInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (extensionGovernance *ExtensionGovernance) UnpackInvalidGovernanceHashError(raw []byte) (*ExtensionGovernanceInvalidGovernanceHash, error) {
	out := new(ExtensionGovernanceInvalidGovernanceHash)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceInvalidKeyType represents a InvalidKeyType error raised by the ExtensionGovernance contract.
type ExtensionGovernanceInvalidKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyType()
func ExtensionGovernanceInvalidKeyTypeErrorID() common.Hash {
	return common.HexToHash("0x36dccb440beb343622d8ee0927bce3f69281feea3a5b0610f6e00e2811b557e1")
}

// UnpackInvalidKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyType()
func (extensionGovernance *ExtensionGovernance) UnpackInvalidKeyTypeError(raw []byte) (*ExtensionGovernanceInvalidKeyType, error) {
	out := new(ExtensionGovernanceInvalidKeyType)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "InvalidKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceInvalidNonce represents a InvalidNonce error raised by the ExtensionGovernance contract.
type ExtensionGovernanceInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func ExtensionGovernanceInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (extensionGovernance *ExtensionGovernance) UnpackInvalidNonceError(raw []byte) (*ExtensionGovernanceInvalidNonce, error) {
	out := new(ExtensionGovernanceInvalidNonce)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceInvalidPublicKey represents a InvalidPublicKey error raised by the ExtensionGovernance contract.
type ExtensionGovernanceInvalidPublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPublicKey()
func ExtensionGovernanceInvalidPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa2d0fee854985cb3d3a394146b5df1c41d5b4d6fcede7fc57a3c05a49b2f3e73")
}

// UnpackInvalidPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPublicKey()
func (extensionGovernance *ExtensionGovernance) UnpackInvalidPublicKeyError(raw []byte) (*ExtensionGovernanceInvalidPublicKey, error) {
	out := new(ExtensionGovernanceInvalidPublicKey)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "InvalidPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceInvalidResponseData represents a InvalidResponseData error raised by the ExtensionGovernance contract.
type ExtensionGovernanceInvalidResponseData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseData()
func ExtensionGovernanceInvalidResponseDataErrorID() common.Hash {
	return common.HexToHash("0x48566617128355a033880805d6462da9bf9c6fe4c61b922997451a71262990af")
}

// UnpackInvalidResponseDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseData()
func (extensionGovernance *ExtensionGovernance) UnpackInvalidResponseDataError(raw []byte) (*ExtensionGovernanceInvalidResponseData, error) {
	out := new(ExtensionGovernanceInvalidResponseData)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "InvalidResponseData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceInvalidSigner represents a InvalidSigner error raised by the ExtensionGovernance contract.
type ExtensionGovernanceInvalidSigner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigner()
func ExtensionGovernanceInvalidSignerErrorID() common.Hash {
	return common.HexToHash("0x815e1d64efb74fbe314c20a2b8a2335d18bce12a19165e447fa36bcb35959528")
}

// UnpackInvalidSignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigner()
func (extensionGovernance *ExtensionGovernance) UnpackInvalidSignerError(raw []byte) (*ExtensionGovernanceInvalidSigner, error) {
	out := new(ExtensionGovernanceInvalidSigner)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "InvalidSigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceInvalidSigningAlgo represents a InvalidSigningAlgo error raised by the ExtensionGovernance contract.
type ExtensionGovernanceInvalidSigningAlgo struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningAlgo()
func ExtensionGovernanceInvalidSigningAlgoErrorID() common.Hash {
	return common.HexToHash("0x9b7198a697b61e546ef9ee593a6f77d2bf9f1df9dcbba655ce96d7e2c75346c7")
}

// UnpackInvalidSigningAlgoError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningAlgo()
func (extensionGovernance *ExtensionGovernance) UnpackInvalidSigningAlgoError(raw []byte) (*ExtensionGovernanceInvalidSigningAlgo, error) {
	out := new(ExtensionGovernanceInvalidSigningAlgo)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "InvalidSigningAlgo", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceInvalidThreshold represents a InvalidThreshold error raised by the ExtensionGovernance contract.
type ExtensionGovernanceInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func ExtensionGovernanceInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (extensionGovernance *ExtensionGovernance) UnpackInvalidThresholdError(raw []byte) (*ExtensionGovernanceInvalidThreshold, error) {
	out := new(ExtensionGovernanceInvalidThreshold)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceInvalidWalletStatus represents a InvalidWalletStatus error raised by the ExtensionGovernance contract.
type ExtensionGovernanceInvalidWalletStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidWalletStatus()
func ExtensionGovernanceInvalidWalletStatusErrorID() common.Hash {
	return common.HexToHash("0xd700120e4896ea7835a5e9e4feb690644ced5cad1c52209afdcbe06b96d91345")
}

// UnpackInvalidWalletStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidWalletStatus()
func (extensionGovernance *ExtensionGovernance) UnpackInvalidWalletStatusError(raw []byte) (*ExtensionGovernanceInvalidWalletStatus, error) {
	out := new(ExtensionGovernanceInvalidWalletStatus)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "InvalidWalletStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceKeyTypeNotSupported represents a KeyTypeNotSupported error raised by the ExtensionGovernance contract.
type ExtensionGovernanceKeyTypeNotSupported struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func ExtensionGovernanceKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x85004df1f47814c2de6fbf03d3f2b1ca8c67c182d0d54278766543b6be3bc603")
}

// UnpackKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func (extensionGovernance *ExtensionGovernance) UnpackKeyTypeNotSupportedError(raw []byte) (*ExtensionGovernanceKeyTypeNotSupported, error) {
	out := new(ExtensionGovernanceKeyTypeNotSupported)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "KeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceLengthsMismatch represents a LengthsMismatch error raised by the ExtensionGovernance contract.
type ExtensionGovernanceLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func ExtensionGovernanceLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (extensionGovernance *ExtensionGovernance) UnpackLengthsMismatchError(raw []byte) (*ExtensionGovernanceLengthsMismatch, error) {
	out := new(ExtensionGovernanceLengthsMismatch)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceNoAddresses represents a NoAddresses error raised by the ExtensionGovernance contract.
type ExtensionGovernanceNoAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoAddresses()
func ExtensionGovernanceNoAddressesErrorID() common.Hash {
	return common.HexToHash("0xb2e79f05746fa6c9f3d61f9efe85959371cde5030a6c2dd70c92b63578c5cde7")
}

// UnpackNoAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoAddresses()
func (extensionGovernance *ExtensionGovernance) UnpackNoAddressesError(raw []byte) (*ExtensionGovernanceNoAddresses, error) {
	out := new(ExtensionGovernanceNoAddresses)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "NoAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceNoSigners represents a NoSigners error raised by the ExtensionGovernance contract.
type ExtensionGovernanceNoSigners struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoSigners()
func ExtensionGovernanceNoSignersErrorID() common.Hash {
	return common.HexToHash("0xc7af40f04dbbe7f3c132fc897339fea9aa96418b5a8074edabf89ed0e4594ae0")
}

// UnpackNoSignersError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoSigners()
func (extensionGovernance *ExtensionGovernance) UnpackNoSignersError(raw []byte) (*ExtensionGovernanceNoSigners, error) {
	out := new(ExtensionGovernanceNoSigners)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "NoSigners", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceNotOwnerOrPauser represents a NotOwnerOrPauser error raised by the ExtensionGovernance contract.
type ExtensionGovernanceNotOwnerOrPauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrPauser(address caller)
func ExtensionGovernanceNotOwnerOrPauserErrorID() common.Hash {
	return common.HexToHash("0x46c0d9af6b895093e022ad2984c54e60192e6e26c5063daf232e825c6e07ff9a")
}

// UnpackNotOwnerOrPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrPauser(address caller)
func (extensionGovernance *ExtensionGovernance) UnpackNotOwnerOrPauserError(raw []byte) (*ExtensionGovernanceNotOwnerOrPauser, error) {
	out := new(ExtensionGovernanceNotOwnerOrPauser)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "NotOwnerOrPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceNotOwnerOrUnpauser represents a NotOwnerOrUnpauser error raised by the ExtensionGovernance contract.
type ExtensionGovernanceNotOwnerOrUnpauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func ExtensionGovernanceNotOwnerOrUnpauserErrorID() common.Hash {
	return common.HexToHash("0x81e700e8f0bd5ba4a6d6a2ce9a1ad0ac9bb1b1b1e62d921a55cbc4b53f1da649")
}

// UnpackNotOwnerOrUnpauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func (extensionGovernance *ExtensionGovernance) UnpackNotOwnerOrUnpauserError(raw []byte) (*ExtensionGovernanceNotOwnerOrUnpauser, error) {
	out := new(ExtensionGovernanceNotOwnerOrUnpauser)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "NotOwnerOrUnpauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceOnlyExtensionOwner represents a OnlyExtensionOwner error raised by the ExtensionGovernance contract.
type ExtensionGovernanceOnlyExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwner()
func ExtensionGovernanceOnlyExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xcd49fd1de5a297cd1ce055d5abffcf2cb25ad9492bd39a703a238af57b2a9488")
}

// UnpackOnlyExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwner()
func (extensionGovernance *ExtensionGovernance) UnpackOnlyExtensionOwnerError(raw []byte) (*ExtensionGovernanceOnlyExtensionOwner, error) {
	out := new(ExtensionGovernanceOnlyExtensionOwner)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "OnlyExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the ExtensionGovernance contract.
type ExtensionGovernanceOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func ExtensionGovernanceOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (extensionGovernance *ExtensionGovernance) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*ExtensionGovernanceOnlyExtensionOwnerOrOperator, error) {
	out := new(ExtensionGovernanceOnlyExtensionOwnerOrOperator)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceOnlyOwner represents a OnlyOwner error raised by the ExtensionGovernance contract.
type ExtensionGovernanceOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func ExtensionGovernanceOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (extensionGovernance *ExtensionGovernance) UnpackOnlyOwnerError(raw []byte) (*ExtensionGovernanceOnlyOwner, error) {
	out := new(ExtensionGovernanceOnlyOwner)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceOnlyOwnerOrBackupManager represents a OnlyOwnerOrBackupManager error raised by the ExtensionGovernance contract.
type ExtensionGovernanceOnlyOwnerOrBackupManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrBackupManager()
func ExtensionGovernanceOnlyOwnerOrBackupManagerErrorID() common.Hash {
	return common.HexToHash("0x249a6cf6e22092f9c80798e2a0c583c0f36aea0779ef38a5b5cdb6a112a0e6df")
}

// UnpackOnlyOwnerOrBackupManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrBackupManager()
func (extensionGovernance *ExtensionGovernance) UnpackOnlyOwnerOrBackupManagerError(raw []byte) (*ExtensionGovernanceOnlyOwnerOrBackupManager, error) {
	out := new(ExtensionGovernanceOnlyOwnerOrBackupManager)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "OnlyOwnerOrBackupManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the ExtensionGovernance contract.
type ExtensionGovernanceOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func ExtensionGovernanceOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (extensionGovernance *ExtensionGovernance) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*ExtensionGovernanceOnlyProductionOrPausedStatus, error) {
	out := new(ExtensionGovernanceOnlyProductionOrPausedStatus)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceOnlyProposedOwner represents a OnlyProposedOwner error raised by the ExtensionGovernance contract.
type ExtensionGovernanceOnlyProposedOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProposedOwner()
func ExtensionGovernanceOnlyProposedOwnerErrorID() common.Hash {
	return common.HexToHash("0x49cb15061ee558e74f66f937f740fd2a3791a89d2d139862222e5c97cabc02cd")
}

// UnpackOnlyProposedOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProposedOwner()
func (extensionGovernance *ExtensionGovernance) UnpackOnlyProposedOwnerError(raw []byte) (*ExtensionGovernanceOnlyProposedOwner, error) {
	out := new(ExtensionGovernanceOnlyProposedOwner)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "OnlyProposedOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceOwnerNotAllowed represents a OwnerNotAllowed error raised by the ExtensionGovernance contract.
type ExtensionGovernanceOwnerNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerNotAllowed()
func ExtensionGovernanceOwnerNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x3d5c9996545c6bf4287c26c1c2d8e6b4c39b92f19ba706c3a756720bf9c4fd24")
}

// UnpackOwnerNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerNotAllowed()
func (extensionGovernance *ExtensionGovernance) UnpackOwnerNotAllowedError(raw []byte) (*ExtensionGovernanceOwnerNotAllowed, error) {
	out := new(ExtensionGovernanceOwnerNotAllowed)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "OwnerNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceSafeDomainSeparatorMismatch represents a SafeDomainSeparatorMismatch error raised by the ExtensionGovernance contract.
type ExtensionGovernanceSafeDomainSeparatorMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeDomainSeparatorMismatch()
func ExtensionGovernanceSafeDomainSeparatorMismatchErrorID() common.Hash {
	return common.HexToHash("0x0e631e9beed1622cc2624cfc2691a4006ec1f6734e6803afedd3bfbf5a7f02fb")
}

// UnpackSafeDomainSeparatorMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeDomainSeparatorMismatch()
func (extensionGovernance *ExtensionGovernance) UnpackSafeDomainSeparatorMismatchError(raw []byte) (*ExtensionGovernanceSafeDomainSeparatorMismatch, error) {
	out := new(ExtensionGovernanceSafeDomainSeparatorMismatch)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "SafeDomainSeparatorMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceSignerAlreadyExists represents a SignerAlreadyExists error raised by the ExtensionGovernance contract.
type ExtensionGovernanceSignerAlreadyExists struct {
	Signer common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SignerAlreadyExists(address signer)
func ExtensionGovernanceSignerAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x899ec3691b11a1bae5010d8f4fd4602162278fe24c4c6de775ae361d6bab8333")
}

// UnpackSignerAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SignerAlreadyExists(address signer)
func (extensionGovernance *ExtensionGovernance) UnpackSignerAlreadyExistsError(raw []byte) (*ExtensionGovernanceSignerAlreadyExists, error) {
	out := new(ExtensionGovernanceSignerAlreadyExists)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "SignerAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the ExtensionGovernance contract.
type ExtensionGovernanceTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func ExtensionGovernanceTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (extensionGovernance *ExtensionGovernance) UnpackTeeMachineNotAvailableError(raw []byte) (*ExtensionGovernanceTeeMachineNotAvailable, error) {
	out := new(ExtensionGovernanceTeeMachineNotAvailable)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionGovernanceVersionNotSupported represents a VersionNotSupported error raised by the ExtensionGovernance contract.
type ExtensionGovernanceVersionNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionNotSupported()
func ExtensionGovernanceVersionNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x6db0b17b06266b37e721024a8ba1221439799ebb779dfd397865e37250abc9a8")
}

// UnpackVersionNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionNotSupported()
func (extensionGovernance *ExtensionGovernance) UnpackVersionNotSupportedError(raw []byte) (*ExtensionGovernanceVersionNotSupported, error) {
	out := new(ExtensionGovernanceVersionNotSupported)
	if err := extensionGovernance.abi.UnpackIntoInterface(out, "VersionNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}
