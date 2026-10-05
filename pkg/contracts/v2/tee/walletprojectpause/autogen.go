// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletprojectpause

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

// WalletProjectPauseMetaData contains all meta data concerning the WalletProjectPause contract.
var WalletProjectPauseMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWalletIds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"WalletProjectPausersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"WalletProjectPausersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"WalletProjectUnpausersAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"WalletProjectUnpausersRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"walletIds\",\"type\":\"bytes32[]\"}],\"name\":\"WalletsPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"walletIds\",\"type\":\"bytes32[]\"}],\"name\":\"WalletsUnpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addWalletProjectPausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addWalletProjectUnpausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getWalletProjectPausers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getWalletProjectUnpausers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isWalletProjectPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isPauser\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isWalletProjectUnpauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isUnpauser\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_walletIds\",\"type\":\"bytes32[]\"}],\"name\":\"pauseWallets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"removeWalletProjectPausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"removeWalletProjectUnpausers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_walletIds\",\"type\":\"bytes32[]\"}],\"name\":\"unpauseWallets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "WalletProjectPause",
}

// WalletProjectPause is an auto generated Go binding around an Ethereum contract.
type WalletProjectPause struct {
	abi abi.ABI
}

// NewWalletProjectPause creates a new instance of WalletProjectPause.
func NewWalletProjectPause() *WalletProjectPause {
	parsed, err := WalletProjectPauseMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &WalletProjectPause{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *WalletProjectPause) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddWalletProjectPausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2992a55f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addWalletProjectPausers(bytes32 _projectId, address[] _addresses) returns()
func (walletProjectPause *WalletProjectPause) PackAddWalletProjectPausers(projectId [32]byte, addresses []common.Address) []byte {
	enc, err := walletProjectPause.abi.Pack("addWalletProjectPausers", projectId, addresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddWalletProjectPausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2992a55f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addWalletProjectPausers(bytes32 _projectId, address[] _addresses) returns()
func (walletProjectPause *WalletProjectPause) TryPackAddWalletProjectPausers(projectId [32]byte, addresses []common.Address) ([]byte, error) {
	return walletProjectPause.abi.Pack("addWalletProjectPausers", projectId, addresses)
}

// PackAddWalletProjectUnpausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x90a84ccf.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addWalletProjectUnpausers(bytes32 _projectId, address[] _addresses) returns()
func (walletProjectPause *WalletProjectPause) PackAddWalletProjectUnpausers(projectId [32]byte, addresses []common.Address) []byte {
	enc, err := walletProjectPause.abi.Pack("addWalletProjectUnpausers", projectId, addresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddWalletProjectUnpausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x90a84ccf.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addWalletProjectUnpausers(bytes32 _projectId, address[] _addresses) returns()
func (walletProjectPause *WalletProjectPause) TryPackAddWalletProjectUnpausers(projectId [32]byte, addresses []common.Address) ([]byte, error) {
	return walletProjectPause.abi.Pack("addWalletProjectUnpausers", projectId, addresses)
}

// PackGetWalletProjectPausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xefaddfab.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletProjectPausers(bytes32 _projectId) view returns(address[] _addresses)
func (walletProjectPause *WalletProjectPause) PackGetWalletProjectPausers(projectId [32]byte) []byte {
	enc, err := walletProjectPause.abi.Pack("getWalletProjectPausers", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletProjectPausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xefaddfab.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletProjectPausers(bytes32 _projectId) view returns(address[] _addresses)
func (walletProjectPause *WalletProjectPause) TryPackGetWalletProjectPausers(projectId [32]byte) ([]byte, error) {
	return walletProjectPause.abi.Pack("getWalletProjectPausers", projectId)
}

// UnpackGetWalletProjectPausers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xefaddfab.
//
// Solidity: function getWalletProjectPausers(bytes32 _projectId) view returns(address[] _addresses)
func (walletProjectPause *WalletProjectPause) UnpackGetWalletProjectPausers(data []byte) ([]common.Address, error) {
	out, err := walletProjectPause.abi.Unpack("getWalletProjectPausers", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetWalletProjectUnpausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3bb392db.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletProjectUnpausers(bytes32 _projectId) view returns(address[] _addresses)
func (walletProjectPause *WalletProjectPause) PackGetWalletProjectUnpausers(projectId [32]byte) []byte {
	enc, err := walletProjectPause.abi.Pack("getWalletProjectUnpausers", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletProjectUnpausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3bb392db.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletProjectUnpausers(bytes32 _projectId) view returns(address[] _addresses)
func (walletProjectPause *WalletProjectPause) TryPackGetWalletProjectUnpausers(projectId [32]byte) ([]byte, error) {
	return walletProjectPause.abi.Pack("getWalletProjectUnpausers", projectId)
}

// UnpackGetWalletProjectUnpausers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3bb392db.
//
// Solidity: function getWalletProjectUnpausers(bytes32 _projectId) view returns(address[] _addresses)
func (walletProjectPause *WalletProjectPause) UnpackGetWalletProjectUnpausers(data []byte) ([]common.Address, error) {
	out, err := walletProjectPause.abi.Unpack("getWalletProjectUnpausers", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackIsWalletProjectPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe27b3ee.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isWalletProjectPauser(bytes32 _projectId, address _addr) view returns(bool _isPauser)
func (walletProjectPause *WalletProjectPause) PackIsWalletProjectPauser(projectId [32]byte, addr common.Address) []byte {
	enc, err := walletProjectPause.abi.Pack("isWalletProjectPauser", projectId, addr)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsWalletProjectPauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe27b3ee.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isWalletProjectPauser(bytes32 _projectId, address _addr) view returns(bool _isPauser)
func (walletProjectPause *WalletProjectPause) TryPackIsWalletProjectPauser(projectId [32]byte, addr common.Address) ([]byte, error) {
	return walletProjectPause.abi.Pack("isWalletProjectPauser", projectId, addr)
}

// UnpackIsWalletProjectPauser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfe27b3ee.
//
// Solidity: function isWalletProjectPauser(bytes32 _projectId, address _addr) view returns(bool _isPauser)
func (walletProjectPause *WalletProjectPause) UnpackIsWalletProjectPauser(data []byte) (bool, error) {
	out, err := walletProjectPause.abi.Unpack("isWalletProjectPauser", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsWalletProjectUnpauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe5fdda95.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isWalletProjectUnpauser(bytes32 _projectId, address _addr) view returns(bool _isUnpauser)
func (walletProjectPause *WalletProjectPause) PackIsWalletProjectUnpauser(projectId [32]byte, addr common.Address) []byte {
	enc, err := walletProjectPause.abi.Pack("isWalletProjectUnpauser", projectId, addr)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsWalletProjectUnpauser is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe5fdda95.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isWalletProjectUnpauser(bytes32 _projectId, address _addr) view returns(bool _isUnpauser)
func (walletProjectPause *WalletProjectPause) TryPackIsWalletProjectUnpauser(projectId [32]byte, addr common.Address) ([]byte, error) {
	return walletProjectPause.abi.Pack("isWalletProjectUnpauser", projectId, addr)
}

// UnpackIsWalletProjectUnpauser is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe5fdda95.
//
// Solidity: function isWalletProjectUnpauser(bytes32 _projectId, address _addr) view returns(bool _isUnpauser)
func (walletProjectPause *WalletProjectPause) UnpackIsWalletProjectUnpauser(data []byte) (bool, error) {
	out, err := walletProjectPause.abi.Unpack("isWalletProjectUnpauser", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackPauseWallets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7089cca6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pauseWallets(bytes32[] _walletIds) returns()
func (walletProjectPause *WalletProjectPause) PackPauseWallets(walletIds [][32]byte) []byte {
	enc, err := walletProjectPause.abi.Pack("pauseWallets", walletIds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPauseWallets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7089cca6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pauseWallets(bytes32[] _walletIds) returns()
func (walletProjectPause *WalletProjectPause) TryPackPauseWallets(walletIds [][32]byte) ([]byte, error) {
	return walletProjectPause.abi.Pack("pauseWallets", walletIds)
}

// PackRemoveWalletProjectPausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7fb4b937.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeWalletProjectPausers(bytes32 _projectId, address[] _addresses) returns()
func (walletProjectPause *WalletProjectPause) PackRemoveWalletProjectPausers(projectId [32]byte, addresses []common.Address) []byte {
	enc, err := walletProjectPause.abi.Pack("removeWalletProjectPausers", projectId, addresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveWalletProjectPausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7fb4b937.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeWalletProjectPausers(bytes32 _projectId, address[] _addresses) returns()
func (walletProjectPause *WalletProjectPause) TryPackRemoveWalletProjectPausers(projectId [32]byte, addresses []common.Address) ([]byte, error) {
	return walletProjectPause.abi.Pack("removeWalletProjectPausers", projectId, addresses)
}

// PackRemoveWalletProjectUnpausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7250ea34.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeWalletProjectUnpausers(bytes32 _projectId, address[] _addresses) returns()
func (walletProjectPause *WalletProjectPause) PackRemoveWalletProjectUnpausers(projectId [32]byte, addresses []common.Address) []byte {
	enc, err := walletProjectPause.abi.Pack("removeWalletProjectUnpausers", projectId, addresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveWalletProjectUnpausers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7250ea34.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeWalletProjectUnpausers(bytes32 _projectId, address[] _addresses) returns()
func (walletProjectPause *WalletProjectPause) TryPackRemoveWalletProjectUnpausers(projectId [32]byte, addresses []common.Address) ([]byte, error) {
	return walletProjectPause.abi.Pack("removeWalletProjectUnpausers", projectId, addresses)
}

// PackUnpauseWallets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f85f96f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function unpauseWallets(bytes32[] _walletIds) returns()
func (walletProjectPause *WalletProjectPause) PackUnpauseWallets(walletIds [][32]byte) []byte {
	enc, err := walletProjectPause.abi.Pack("unpauseWallets", walletIds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUnpauseWallets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5f85f96f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function unpauseWallets(bytes32[] _walletIds) returns()
func (walletProjectPause *WalletProjectPause) TryPackUnpauseWallets(walletIds [][32]byte) ([]byte, error) {
	return walletProjectPause.abi.Pack("unpauseWallets", walletIds)
}

// WalletProjectPauseWalletProjectPausersAdded represents a WalletProjectPausersAdded event raised by the WalletProjectPause contract.
type WalletProjectPauseWalletProjectPausersAdded struct {
	ProjectId [32]byte
	Addresses []common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletProjectPauseWalletProjectPausersAddedEventName = "WalletProjectPausersAdded"

// ContractEventName returns the user-defined event name.
func (WalletProjectPauseWalletProjectPausersAdded) ContractEventName() string {
	return WalletProjectPauseWalletProjectPausersAddedEventName
}

// UnpackWalletProjectPausersAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletProjectPausersAdded(bytes32 indexed projectId, address[] addresses)
func (walletProjectPause *WalletProjectPause) UnpackWalletProjectPausersAddedEvent(log *types.Log) (*WalletProjectPauseWalletProjectPausersAdded, error) {
	event := "WalletProjectPausersAdded"
	if len(log.Topics) == 0 || log.Topics[0] != walletProjectPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletProjectPauseWalletProjectPausersAdded)
	if len(log.Data) > 0 {
		if err := walletProjectPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletProjectPause.abi.Events[event].Inputs {
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

// WalletProjectPauseWalletProjectPausersRemoved represents a WalletProjectPausersRemoved event raised by the WalletProjectPause contract.
type WalletProjectPauseWalletProjectPausersRemoved struct {
	ProjectId [32]byte
	Addresses []common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletProjectPauseWalletProjectPausersRemovedEventName = "WalletProjectPausersRemoved"

// ContractEventName returns the user-defined event name.
func (WalletProjectPauseWalletProjectPausersRemoved) ContractEventName() string {
	return WalletProjectPauseWalletProjectPausersRemovedEventName
}

// UnpackWalletProjectPausersRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletProjectPausersRemoved(bytes32 indexed projectId, address[] addresses)
func (walletProjectPause *WalletProjectPause) UnpackWalletProjectPausersRemovedEvent(log *types.Log) (*WalletProjectPauseWalletProjectPausersRemoved, error) {
	event := "WalletProjectPausersRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != walletProjectPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletProjectPauseWalletProjectPausersRemoved)
	if len(log.Data) > 0 {
		if err := walletProjectPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletProjectPause.abi.Events[event].Inputs {
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

// WalletProjectPauseWalletProjectUnpausersAdded represents a WalletProjectUnpausersAdded event raised by the WalletProjectPause contract.
type WalletProjectPauseWalletProjectUnpausersAdded struct {
	ProjectId [32]byte
	Addresses []common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletProjectPauseWalletProjectUnpausersAddedEventName = "WalletProjectUnpausersAdded"

// ContractEventName returns the user-defined event name.
func (WalletProjectPauseWalletProjectUnpausersAdded) ContractEventName() string {
	return WalletProjectPauseWalletProjectUnpausersAddedEventName
}

// UnpackWalletProjectUnpausersAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletProjectUnpausersAdded(bytes32 indexed projectId, address[] addresses)
func (walletProjectPause *WalletProjectPause) UnpackWalletProjectUnpausersAddedEvent(log *types.Log) (*WalletProjectPauseWalletProjectUnpausersAdded, error) {
	event := "WalletProjectUnpausersAdded"
	if len(log.Topics) == 0 || log.Topics[0] != walletProjectPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletProjectPauseWalletProjectUnpausersAdded)
	if len(log.Data) > 0 {
		if err := walletProjectPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletProjectPause.abi.Events[event].Inputs {
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

// WalletProjectPauseWalletProjectUnpausersRemoved represents a WalletProjectUnpausersRemoved event raised by the WalletProjectPause contract.
type WalletProjectPauseWalletProjectUnpausersRemoved struct {
	ProjectId [32]byte
	Addresses []common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletProjectPauseWalletProjectUnpausersRemovedEventName = "WalletProjectUnpausersRemoved"

// ContractEventName returns the user-defined event name.
func (WalletProjectPauseWalletProjectUnpausersRemoved) ContractEventName() string {
	return WalletProjectPauseWalletProjectUnpausersRemovedEventName
}

// UnpackWalletProjectUnpausersRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletProjectUnpausersRemoved(bytes32 indexed projectId, address[] addresses)
func (walletProjectPause *WalletProjectPause) UnpackWalletProjectUnpausersRemovedEvent(log *types.Log) (*WalletProjectPauseWalletProjectUnpausersRemoved, error) {
	event := "WalletProjectUnpausersRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != walletProjectPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletProjectPauseWalletProjectUnpausersRemoved)
	if len(log.Data) > 0 {
		if err := walletProjectPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletProjectPause.abi.Events[event].Inputs {
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

// WalletProjectPauseWalletsPaused represents a WalletsPaused event raised by the WalletProjectPause contract.
type WalletProjectPauseWalletsPaused struct {
	WalletIds [][32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletProjectPauseWalletsPausedEventName = "WalletsPaused"

// ContractEventName returns the user-defined event name.
func (WalletProjectPauseWalletsPaused) ContractEventName() string {
	return WalletProjectPauseWalletsPausedEventName
}

// UnpackWalletsPausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletsPaused(bytes32[] walletIds)
func (walletProjectPause *WalletProjectPause) UnpackWalletsPausedEvent(log *types.Log) (*WalletProjectPauseWalletsPaused, error) {
	event := "WalletsPaused"
	if len(log.Topics) == 0 || log.Topics[0] != walletProjectPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletProjectPauseWalletsPaused)
	if len(log.Data) > 0 {
		if err := walletProjectPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletProjectPause.abi.Events[event].Inputs {
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

// WalletProjectPauseWalletsUnpaused represents a WalletsUnpaused event raised by the WalletProjectPause contract.
type WalletProjectPauseWalletsUnpaused struct {
	WalletIds [][32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletProjectPauseWalletsUnpausedEventName = "WalletsUnpaused"

// ContractEventName returns the user-defined event name.
func (WalletProjectPauseWalletsUnpaused) ContractEventName() string {
	return WalletProjectPauseWalletsUnpausedEventName
}

// UnpackWalletsUnpausedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletsUnpaused(bytes32[] walletIds)
func (walletProjectPause *WalletProjectPause) UnpackWalletsUnpausedEvent(log *types.Log) (*WalletProjectPauseWalletsUnpaused, error) {
	event := "WalletsUnpaused"
	if len(log.Topics) == 0 || log.Topics[0] != walletProjectPause.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletProjectPauseWalletsUnpaused)
	if len(log.Data) > 0 {
		if err := walletProjectPause.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletProjectPause.abi.Events[event].Inputs {
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
func (walletProjectPause *WalletProjectPause) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["AddressAlreadyInSet"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackAddressAlreadyInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["AddressNotInSet"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackAddressNotInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["AvailabilityCheckTimestampInvalid"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackAvailabilityCheckTimestampInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["DuplicatedCosigner"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackDuplicatedCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["InvalidAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackInvalidAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["InvalidKeyType"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackInvalidKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["InvalidPublicKey"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackInvalidPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["InvalidResponseData"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackInvalidResponseDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["InvalidSigningAlgo"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackInvalidSigningAlgoError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["InvalidWalletStatus"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackInvalidWalletStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["KeyTypeNotSupported"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["NoAddresses"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackNoAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["NoWalletIds"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackNoWalletIdsError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["NotOwnerOrPauser"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackNotOwnerOrPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["NotOwnerOrUnpauser"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackNotOwnerOrUnpauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["OnlyExtensionOwner"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackOnlyExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackOnlyOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["OnlyOwnerOrBackupManager"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackOnlyOwnerOrBackupManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["OnlyProposedOwner"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackOnlyProposedOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["OwnerNotAllowed"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackOwnerNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletProjectPause.abi.Errors["VersionNotSupported"].ID.Bytes()[:4]) {
		return walletProjectPause.UnpackVersionNotSupportedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// WalletProjectPauseAddressAlreadyInSet represents a AddressAlreadyInSet error raised by the WalletProjectPause contract.
type WalletProjectPauseAddressAlreadyInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressAlreadyInSet(address addr)
func WalletProjectPauseAddressAlreadyInSetErrorID() common.Hash {
	return common.HexToHash("0xb725c7fb1c3f0f065973d6b9eb55140f970259955af3045742cd5b37902eb920")
}

// UnpackAddressAlreadyInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressAlreadyInSet(address addr)
func (walletProjectPause *WalletProjectPause) UnpackAddressAlreadyInSetError(raw []byte) (*WalletProjectPauseAddressAlreadyInSet, error) {
	out := new(WalletProjectPauseAddressAlreadyInSet)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "AddressAlreadyInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseAddressNotInSet represents a AddressNotInSet error raised by the WalletProjectPause contract.
type WalletProjectPauseAddressNotInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressNotInSet(address addr)
func WalletProjectPauseAddressNotInSetErrorID() common.Hash {
	return common.HexToHash("0x1ff5ecf3e6d385ea01ff0323fe894ed6d25a83ed5f6ff2a2ede4c897408f6be9")
}

// UnpackAddressNotInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressNotInSet(address addr)
func (walletProjectPause *WalletProjectPause) UnpackAddressNotInSetError(raw []byte) (*WalletProjectPauseAddressNotInSet, error) {
	out := new(WalletProjectPauseAddressNotInSet)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "AddressNotInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseAvailabilityCheckTimestampInvalid represents a AvailabilityCheckTimestampInvalid error raised by the WalletProjectPause contract.
type WalletProjectPauseAvailabilityCheckTimestampInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func WalletProjectPauseAvailabilityCheckTimestampInvalidErrorID() common.Hash {
	return common.HexToHash("0x6344a5e74cff78612fa271fa907ca608af3bad10e1e8c8df302e29ebc6298e2d")
}

// UnpackAvailabilityCheckTimestampInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func (walletProjectPause *WalletProjectPause) UnpackAvailabilityCheckTimestampInvalidError(raw []byte) (*WalletProjectPauseAvailabilityCheckTimestampInvalid, error) {
	out := new(WalletProjectPauseAvailabilityCheckTimestampInvalid)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "AvailabilityCheckTimestampInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseDuplicatedCosigner represents a DuplicatedCosigner error raised by the WalletProjectPause contract.
type WalletProjectPauseDuplicatedCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func WalletProjectPauseDuplicatedCosignerErrorID() common.Hash {
	return common.HexToHash("0x53a5890417576fb3043dc3ff1669b0932da2f657d8adb4c11313b12b711f5910")
}

// UnpackDuplicatedCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func (walletProjectPause *WalletProjectPause) UnpackDuplicatedCosignerError(raw []byte) (*WalletProjectPauseDuplicatedCosigner, error) {
	out := new(WalletProjectPauseDuplicatedCosigner)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "DuplicatedCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseExtensionIdMismatch represents a ExtensionIdMismatch error raised by the WalletProjectPause contract.
type WalletProjectPauseExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func WalletProjectPauseExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (walletProjectPause *WalletProjectPause) UnpackExtensionIdMismatchError(raw []byte) (*WalletProjectPauseExtensionIdMismatch, error) {
	out := new(WalletProjectPauseExtensionIdMismatch)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseInvalidAddress represents a InvalidAddress error raised by the WalletProjectPause contract.
type WalletProjectPauseInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func WalletProjectPauseInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (walletProjectPause *WalletProjectPause) UnpackInvalidAddressError(raw []byte) (*WalletProjectPauseInvalidAddress, error) {
	out := new(WalletProjectPauseInvalidAddress)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseInvalidAvailabilityCheckStatus represents a InvalidAvailabilityCheckStatus error raised by the WalletProjectPause contract.
type WalletProjectPauseInvalidAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func WalletProjectPauseInvalidAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0xedee4ff00b842d3ef84031c60001065c540c838c5479bbf0ae56f15868427e44")
}

// UnpackInvalidAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func (walletProjectPause *WalletProjectPause) UnpackInvalidAvailabilityCheckStatusError(raw []byte) (*WalletProjectPauseInvalidAvailabilityCheckStatus, error) {
	out := new(WalletProjectPauseInvalidAvailabilityCheckStatus)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "InvalidAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseInvalidCosigner represents a InvalidCosigner error raised by the WalletProjectPause contract.
type WalletProjectPauseInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func WalletProjectPauseInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (walletProjectPause *WalletProjectPause) UnpackInvalidCosignerError(raw []byte) (*WalletProjectPauseInvalidCosigner, error) {
	out := new(WalletProjectPauseInvalidCosigner)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseInvalidDuration represents a InvalidDuration error raised by the WalletProjectPause contract.
type WalletProjectPauseInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func WalletProjectPauseInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (walletProjectPause *WalletProjectPause) UnpackInvalidDurationError(raw []byte) (*WalletProjectPauseInvalidDuration, error) {
	out := new(WalletProjectPauseInvalidDuration)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the WalletProjectPause contract.
type WalletProjectPauseInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func WalletProjectPauseInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (walletProjectPause *WalletProjectPause) UnpackInvalidGovernanceHashError(raw []byte) (*WalletProjectPauseInvalidGovernanceHash, error) {
	out := new(WalletProjectPauseInvalidGovernanceHash)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseInvalidKeyType represents a InvalidKeyType error raised by the WalletProjectPause contract.
type WalletProjectPauseInvalidKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyType()
func WalletProjectPauseInvalidKeyTypeErrorID() common.Hash {
	return common.HexToHash("0x36dccb440beb343622d8ee0927bce3f69281feea3a5b0610f6e00e2811b557e1")
}

// UnpackInvalidKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyType()
func (walletProjectPause *WalletProjectPause) UnpackInvalidKeyTypeError(raw []byte) (*WalletProjectPauseInvalidKeyType, error) {
	out := new(WalletProjectPauseInvalidKeyType)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "InvalidKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseInvalidNonce represents a InvalidNonce error raised by the WalletProjectPause contract.
type WalletProjectPauseInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func WalletProjectPauseInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (walletProjectPause *WalletProjectPause) UnpackInvalidNonceError(raw []byte) (*WalletProjectPauseInvalidNonce, error) {
	out := new(WalletProjectPauseInvalidNonce)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseInvalidPublicKey represents a InvalidPublicKey error raised by the WalletProjectPause contract.
type WalletProjectPauseInvalidPublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPublicKey()
func WalletProjectPauseInvalidPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa2d0fee854985cb3d3a394146b5df1c41d5b4d6fcede7fc57a3c05a49b2f3e73")
}

// UnpackInvalidPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPublicKey()
func (walletProjectPause *WalletProjectPause) UnpackInvalidPublicKeyError(raw []byte) (*WalletProjectPauseInvalidPublicKey, error) {
	out := new(WalletProjectPauseInvalidPublicKey)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "InvalidPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseInvalidResponseData represents a InvalidResponseData error raised by the WalletProjectPause contract.
type WalletProjectPauseInvalidResponseData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseData()
func WalletProjectPauseInvalidResponseDataErrorID() common.Hash {
	return common.HexToHash("0x48566617128355a033880805d6462da9bf9c6fe4c61b922997451a71262990af")
}

// UnpackInvalidResponseDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseData()
func (walletProjectPause *WalletProjectPause) UnpackInvalidResponseDataError(raw []byte) (*WalletProjectPauseInvalidResponseData, error) {
	out := new(WalletProjectPauseInvalidResponseData)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "InvalidResponseData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseInvalidSigningAlgo represents a InvalidSigningAlgo error raised by the WalletProjectPause contract.
type WalletProjectPauseInvalidSigningAlgo struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningAlgo()
func WalletProjectPauseInvalidSigningAlgoErrorID() common.Hash {
	return common.HexToHash("0x9b7198a697b61e546ef9ee593a6f77d2bf9f1df9dcbba655ce96d7e2c75346c7")
}

// UnpackInvalidSigningAlgoError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningAlgo()
func (walletProjectPause *WalletProjectPause) UnpackInvalidSigningAlgoError(raw []byte) (*WalletProjectPauseInvalidSigningAlgo, error) {
	out := new(WalletProjectPauseInvalidSigningAlgo)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "InvalidSigningAlgo", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseInvalidThreshold represents a InvalidThreshold error raised by the WalletProjectPause contract.
type WalletProjectPauseInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func WalletProjectPauseInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (walletProjectPause *WalletProjectPause) UnpackInvalidThresholdError(raw []byte) (*WalletProjectPauseInvalidThreshold, error) {
	out := new(WalletProjectPauseInvalidThreshold)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseInvalidWalletStatus represents a InvalidWalletStatus error raised by the WalletProjectPause contract.
type WalletProjectPauseInvalidWalletStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidWalletStatus()
func WalletProjectPauseInvalidWalletStatusErrorID() common.Hash {
	return common.HexToHash("0xd700120e4896ea7835a5e9e4feb690644ced5cad1c52209afdcbe06b96d91345")
}

// UnpackInvalidWalletStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidWalletStatus()
func (walletProjectPause *WalletProjectPause) UnpackInvalidWalletStatusError(raw []byte) (*WalletProjectPauseInvalidWalletStatus, error) {
	out := new(WalletProjectPauseInvalidWalletStatus)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "InvalidWalletStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseKeyTypeNotSupported represents a KeyTypeNotSupported error raised by the WalletProjectPause contract.
type WalletProjectPauseKeyTypeNotSupported struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func WalletProjectPauseKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x85004df1f47814c2de6fbf03d3f2b1ca8c67c182d0d54278766543b6be3bc603")
}

// UnpackKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func (walletProjectPause *WalletProjectPause) UnpackKeyTypeNotSupportedError(raw []byte) (*WalletProjectPauseKeyTypeNotSupported, error) {
	out := new(WalletProjectPauseKeyTypeNotSupported)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "KeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseLengthsMismatch represents a LengthsMismatch error raised by the WalletProjectPause contract.
type WalletProjectPauseLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func WalletProjectPauseLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (walletProjectPause *WalletProjectPause) UnpackLengthsMismatchError(raw []byte) (*WalletProjectPauseLengthsMismatch, error) {
	out := new(WalletProjectPauseLengthsMismatch)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseNoAddresses represents a NoAddresses error raised by the WalletProjectPause contract.
type WalletProjectPauseNoAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoAddresses()
func WalletProjectPauseNoAddressesErrorID() common.Hash {
	return common.HexToHash("0xb2e79f05746fa6c9f3d61f9efe85959371cde5030a6c2dd70c92b63578c5cde7")
}

// UnpackNoAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoAddresses()
func (walletProjectPause *WalletProjectPause) UnpackNoAddressesError(raw []byte) (*WalletProjectPauseNoAddresses, error) {
	out := new(WalletProjectPauseNoAddresses)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "NoAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseNoWalletIds represents a NoWalletIds error raised by the WalletProjectPause contract.
type WalletProjectPauseNoWalletIds struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoWalletIds()
func WalletProjectPauseNoWalletIdsErrorID() common.Hash {
	return common.HexToHash("0x224f5edc35e78ce6cea3f1e7fbdf67a7b46805bc997e5fd552ef3184117e86bf")
}

// UnpackNoWalletIdsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoWalletIds()
func (walletProjectPause *WalletProjectPause) UnpackNoWalletIdsError(raw []byte) (*WalletProjectPauseNoWalletIds, error) {
	out := new(WalletProjectPauseNoWalletIds)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "NoWalletIds", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseNotOwnerOrPauser represents a NotOwnerOrPauser error raised by the WalletProjectPause contract.
type WalletProjectPauseNotOwnerOrPauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrPauser(address caller)
func WalletProjectPauseNotOwnerOrPauserErrorID() common.Hash {
	return common.HexToHash("0x46c0d9af6b895093e022ad2984c54e60192e6e26c5063daf232e825c6e07ff9a")
}

// UnpackNotOwnerOrPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrPauser(address caller)
func (walletProjectPause *WalletProjectPause) UnpackNotOwnerOrPauserError(raw []byte) (*WalletProjectPauseNotOwnerOrPauser, error) {
	out := new(WalletProjectPauseNotOwnerOrPauser)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "NotOwnerOrPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseNotOwnerOrUnpauser represents a NotOwnerOrUnpauser error raised by the WalletProjectPause contract.
type WalletProjectPauseNotOwnerOrUnpauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func WalletProjectPauseNotOwnerOrUnpauserErrorID() common.Hash {
	return common.HexToHash("0x81e700e8f0bd5ba4a6d6a2ce9a1ad0ac9bb1b1b1e62d921a55cbc4b53f1da649")
}

// UnpackNotOwnerOrUnpauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func (walletProjectPause *WalletProjectPause) UnpackNotOwnerOrUnpauserError(raw []byte) (*WalletProjectPauseNotOwnerOrUnpauser, error) {
	out := new(WalletProjectPauseNotOwnerOrUnpauser)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "NotOwnerOrUnpauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseOnlyExtensionOwner represents a OnlyExtensionOwner error raised by the WalletProjectPause contract.
type WalletProjectPauseOnlyExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwner()
func WalletProjectPauseOnlyExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xcd49fd1de5a297cd1ce055d5abffcf2cb25ad9492bd39a703a238af57b2a9488")
}

// UnpackOnlyExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwner()
func (walletProjectPause *WalletProjectPause) UnpackOnlyExtensionOwnerError(raw []byte) (*WalletProjectPauseOnlyExtensionOwner, error) {
	out := new(WalletProjectPauseOnlyExtensionOwner)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "OnlyExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the WalletProjectPause contract.
type WalletProjectPauseOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func WalletProjectPauseOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (walletProjectPause *WalletProjectPause) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*WalletProjectPauseOnlyExtensionOwnerOrOperator, error) {
	out := new(WalletProjectPauseOnlyExtensionOwnerOrOperator)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseOnlyOwner represents a OnlyOwner error raised by the WalletProjectPause contract.
type WalletProjectPauseOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func WalletProjectPauseOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (walletProjectPause *WalletProjectPause) UnpackOnlyOwnerError(raw []byte) (*WalletProjectPauseOnlyOwner, error) {
	out := new(WalletProjectPauseOnlyOwner)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseOnlyOwnerOrBackupManager represents a OnlyOwnerOrBackupManager error raised by the WalletProjectPause contract.
type WalletProjectPauseOnlyOwnerOrBackupManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrBackupManager()
func WalletProjectPauseOnlyOwnerOrBackupManagerErrorID() common.Hash {
	return common.HexToHash("0x249a6cf6e22092f9c80798e2a0c583c0f36aea0779ef38a5b5cdb6a112a0e6df")
}

// UnpackOnlyOwnerOrBackupManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrBackupManager()
func (walletProjectPause *WalletProjectPause) UnpackOnlyOwnerOrBackupManagerError(raw []byte) (*WalletProjectPauseOnlyOwnerOrBackupManager, error) {
	out := new(WalletProjectPauseOnlyOwnerOrBackupManager)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "OnlyOwnerOrBackupManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the WalletProjectPause contract.
type WalletProjectPauseOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func WalletProjectPauseOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (walletProjectPause *WalletProjectPause) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*WalletProjectPauseOnlyProductionOrPausedStatus, error) {
	out := new(WalletProjectPauseOnlyProductionOrPausedStatus)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseOnlyProposedOwner represents a OnlyProposedOwner error raised by the WalletProjectPause contract.
type WalletProjectPauseOnlyProposedOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProposedOwner()
func WalletProjectPauseOnlyProposedOwnerErrorID() common.Hash {
	return common.HexToHash("0x49cb15061ee558e74f66f937f740fd2a3791a89d2d139862222e5c97cabc02cd")
}

// UnpackOnlyProposedOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProposedOwner()
func (walletProjectPause *WalletProjectPause) UnpackOnlyProposedOwnerError(raw []byte) (*WalletProjectPauseOnlyProposedOwner, error) {
	out := new(WalletProjectPauseOnlyProposedOwner)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "OnlyProposedOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseOwnerNotAllowed represents a OwnerNotAllowed error raised by the WalletProjectPause contract.
type WalletProjectPauseOwnerNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerNotAllowed()
func WalletProjectPauseOwnerNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x3d5c9996545c6bf4287c26c1c2d8e6b4c39b92f19ba706c3a756720bf9c4fd24")
}

// UnpackOwnerNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerNotAllowed()
func (walletProjectPause *WalletProjectPause) UnpackOwnerNotAllowedError(raw []byte) (*WalletProjectPauseOwnerNotAllowed, error) {
	out := new(WalletProjectPauseOwnerNotAllowed)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "OwnerNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the WalletProjectPause contract.
type WalletProjectPauseTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func WalletProjectPauseTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (walletProjectPause *WalletProjectPause) UnpackTeeMachineNotAvailableError(raw []byte) (*WalletProjectPauseTeeMachineNotAvailable, error) {
	out := new(WalletProjectPauseTeeMachineNotAvailable)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletProjectPauseVersionNotSupported represents a VersionNotSupported error raised by the WalletProjectPause contract.
type WalletProjectPauseVersionNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionNotSupported()
func WalletProjectPauseVersionNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x6db0b17b06266b37e721024a8ba1221439799ebb779dfd397865e37250abc9a8")
}

// UnpackVersionNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionNotSupported()
func (walletProjectPause *WalletProjectPause) UnpackVersionNotSupportedError(raw []byte) (*WalletProjectPauseVersionNotSupported, error) {
	out := new(WalletProjectPauseVersionNotSupported)
	if err := walletProjectPause.abi.UnpackIntoInterface(out, "VersionNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}
