// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletmanager

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

// WalletManagerMetaData contains all meta data concerning the WalletManager contract.
var WalletManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AdminsNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"name\":\"DuplicatedPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdmin\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"name\":\"InvalidAdminPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdminsThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCosignersThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MultisigThresholdNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NotAllAdminsConfirmed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"NotAllCosignersConfirmed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughAdmins\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughKeys\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyAdmins\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyCosigners\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"WalletAdminConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structPublicKey[]\",\"name\":\"adminsPublicKeys\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"adminsThreshold\",\"type\":\"uint64\"}],\"name\":\"WalletAdminsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"WalletCosignerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"WalletCosignersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletInitialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"closeWalletInitialization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"confirmAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"confirmCosigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"createWallet\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"enableWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getProjectWalletIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_walletIds\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAdminsAndThreshold\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_admins\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_adminsThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAdminsPublicKeysAndThreshold\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"_adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"_adminsThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletCosignersAndThreshold\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletProjectId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletStatus\",\"outputs\":[{\"internalType\":\"enumIWalletManager.WalletStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"_adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"_adminsThreshold\",\"type\":\"uint64\"}],\"name\":\"setAdmins\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"setCosigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "WalletManager",
}

// WalletManager is an auto generated Go binding around an Ethereum contract.
type WalletManager struct {
	abi abi.ABI
}

// NewWalletManager creates a new instance of WalletManager.
func NewWalletManager() *WalletManager {
	parsed, err := WalletManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &WalletManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *WalletManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackCloseWalletInitialization is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3c6d865.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function closeWalletInitialization(bytes32 _walletId) returns()
func (walletManager *WalletManager) PackCloseWalletInitialization(walletId [32]byte) []byte {
	enc, err := walletManager.abi.Pack("closeWalletInitialization", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCloseWalletInitialization is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3c6d865.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function closeWalletInitialization(bytes32 _walletId) returns()
func (walletManager *WalletManager) TryPackCloseWalletInitialization(walletId [32]byte) ([]byte, error) {
	return walletManager.abi.Pack("closeWalletInitialization", walletId)
}

// PackConfirmAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ab3b981.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmAdmin(bytes32 _walletId) returns()
func (walletManager *WalletManager) PackConfirmAdmin(walletId [32]byte) []byte {
	enc, err := walletManager.abi.Pack("confirmAdmin", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmAdmin is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ab3b981.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmAdmin(bytes32 _walletId) returns()
func (walletManager *WalletManager) TryPackConfirmAdmin(walletId [32]byte) ([]byte, error) {
	return walletManager.abi.Pack("confirmAdmin", walletId)
}

// PackConfirmCosigner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x77f74c1a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmCosigner(bytes32 _walletId) returns()
func (walletManager *WalletManager) PackConfirmCosigner(walletId [32]byte) []byte {
	enc, err := walletManager.abi.Pack("confirmCosigner", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmCosigner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x77f74c1a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmCosigner(bytes32 _walletId) returns()
func (walletManager *WalletManager) TryPackConfirmCosigner(walletId [32]byte) ([]byte, error) {
	return walletManager.abi.Pack("confirmCosigner", walletId)
}

// PackCreateWallet is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1d647605.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createWallet(bytes32 _projectId) returns(bytes32 _walletId)
func (walletManager *WalletManager) PackCreateWallet(projectId [32]byte) []byte {
	enc, err := walletManager.abi.Pack("createWallet", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreateWallet is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1d647605.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createWallet(bytes32 _projectId) returns(bytes32 _walletId)
func (walletManager *WalletManager) TryPackCreateWallet(projectId [32]byte) ([]byte, error) {
	return walletManager.abi.Pack("createWallet", projectId)
}

// UnpackCreateWallet is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1d647605.
//
// Solidity: function createWallet(bytes32 _projectId) returns(bytes32 _walletId)
func (walletManager *WalletManager) UnpackCreateWallet(data []byte) ([32]byte, error) {
	out, err := walletManager.abi.Unpack("createWallet", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackEnableWallet is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x676f358c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function enableWallet(bytes32 _walletId) returns()
func (walletManager *WalletManager) PackEnableWallet(walletId [32]byte) []byte {
	enc, err := walletManager.abi.Pack("enableWallet", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackEnableWallet is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x676f358c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function enableWallet(bytes32 _walletId) returns()
func (walletManager *WalletManager) TryPackEnableWallet(walletId [32]byte) ([]byte, error) {
	return walletManager.abi.Pack("enableWallet", walletId)
}

// PackGetProjectWalletIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x92911c30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getProjectWalletIds(bytes32 _projectId) view returns(bytes32[] _walletIds)
func (walletManager *WalletManager) PackGetProjectWalletIds(projectId [32]byte) []byte {
	enc, err := walletManager.abi.Pack("getProjectWalletIds", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetProjectWalletIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x92911c30.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getProjectWalletIds(bytes32 _projectId) view returns(bytes32[] _walletIds)
func (walletManager *WalletManager) TryPackGetProjectWalletIds(projectId [32]byte) ([]byte, error) {
	return walletManager.abi.Pack("getProjectWalletIds", projectId)
}

// UnpackGetProjectWalletIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x92911c30.
//
// Solidity: function getProjectWalletIds(bytes32 _projectId) view returns(bytes32[] _walletIds)
func (walletManager *WalletManager) UnpackGetProjectWalletIds(data []byte) ([][32]byte, error) {
	out, err := walletManager.abi.Unpack("getProjectWalletIds", data)
	if err != nil {
		return *new([][32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	return out0, nil
}

// PackGetWalletAdminsAndThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3309abd0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletAdminsAndThreshold(bytes32 _walletId) view returns(address[] _admins, uint64 _adminsThreshold)
func (walletManager *WalletManager) PackGetWalletAdminsAndThreshold(walletId [32]byte) []byte {
	enc, err := walletManager.abi.Pack("getWalletAdminsAndThreshold", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletAdminsAndThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3309abd0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletAdminsAndThreshold(bytes32 _walletId) view returns(address[] _admins, uint64 _adminsThreshold)
func (walletManager *WalletManager) TryPackGetWalletAdminsAndThreshold(walletId [32]byte) ([]byte, error) {
	return walletManager.abi.Pack("getWalletAdminsAndThreshold", walletId)
}

// GetWalletAdminsAndThresholdOutput serves as a container for the return parameters of contract
// method GetWalletAdminsAndThreshold.
type GetWalletAdminsAndThresholdOutput struct {
	Admins          []common.Address
	AdminsThreshold uint64
}

// UnpackGetWalletAdminsAndThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3309abd0.
//
// Solidity: function getWalletAdminsAndThreshold(bytes32 _walletId) view returns(address[] _admins, uint64 _adminsThreshold)
func (walletManager *WalletManager) UnpackGetWalletAdminsAndThreshold(data []byte) (GetWalletAdminsAndThresholdOutput, error) {
	out, err := walletManager.abi.Unpack("getWalletAdminsAndThreshold", data)
	outstruct := new(GetWalletAdminsAndThresholdOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Admins = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.AdminsThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetWalletAdminsPublicKeysAndThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3f0efb4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletAdminsPublicKeysAndThreshold(bytes32 _walletId) view returns((bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold)
func (walletManager *WalletManager) PackGetWalletAdminsPublicKeysAndThreshold(walletId [32]byte) []byte {
	enc, err := walletManager.abi.Pack("getWalletAdminsPublicKeysAndThreshold", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletAdminsPublicKeysAndThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3f0efb4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletAdminsPublicKeysAndThreshold(bytes32 _walletId) view returns((bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold)
func (walletManager *WalletManager) TryPackGetWalletAdminsPublicKeysAndThreshold(walletId [32]byte) ([]byte, error) {
	return walletManager.abi.Pack("getWalletAdminsPublicKeysAndThreshold", walletId)
}

// GetWalletAdminsPublicKeysAndThresholdOutput serves as a container for the return parameters of contract
// method GetWalletAdminsPublicKeysAndThreshold.
type GetWalletAdminsPublicKeysAndThresholdOutput struct {
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
}

// UnpackGetWalletAdminsPublicKeysAndThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf3f0efb4.
//
// Solidity: function getWalletAdminsPublicKeysAndThreshold(bytes32 _walletId) view returns((bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold)
func (walletManager *WalletManager) UnpackGetWalletAdminsPublicKeysAndThreshold(data []byte) (GetWalletAdminsPublicKeysAndThresholdOutput, error) {
	out, err := walletManager.abi.Unpack("getWalletAdminsPublicKeysAndThreshold", data)
	outstruct := new(GetWalletAdminsPublicKeysAndThresholdOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AdminsPublicKeys = *abi.ConvertType(out[0], new([]PublicKey)).(*[]PublicKey)
	outstruct.AdminsThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetWalletCosignersAndThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x881567e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletCosignersAndThreshold(bytes32 _walletId) view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (walletManager *WalletManager) PackGetWalletCosignersAndThreshold(walletId [32]byte) []byte {
	enc, err := walletManager.abi.Pack("getWalletCosignersAndThreshold", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletCosignersAndThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x881567e5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletCosignersAndThreshold(bytes32 _walletId) view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (walletManager *WalletManager) TryPackGetWalletCosignersAndThreshold(walletId [32]byte) ([]byte, error) {
	return walletManager.abi.Pack("getWalletCosignersAndThreshold", walletId)
}

// GetWalletCosignersAndThresholdOutput serves as a container for the return parameters of contract
// method GetWalletCosignersAndThreshold.
type GetWalletCosignersAndThresholdOutput struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}

// UnpackGetWalletCosignersAndThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x881567e5.
//
// Solidity: function getWalletCosignersAndThreshold(bytes32 _walletId) view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (walletManager *WalletManager) UnpackGetWalletCosignersAndThreshold(data []byte) (GetWalletCosignersAndThresholdOutput, error) {
	out, err := walletManager.abi.Unpack("getWalletCosignersAndThreshold", data)
	outstruct := new(GetWalletCosignersAndThresholdOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Cosigners = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.CosignersThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetWalletProjectId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8d5d64f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletProjectId(bytes32 _walletId) view returns(bytes32 _projectId)
func (walletManager *WalletManager) PackGetWalletProjectId(walletId [32]byte) []byte {
	enc, err := walletManager.abi.Pack("getWalletProjectId", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletProjectId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8d5d64f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletProjectId(bytes32 _walletId) view returns(bytes32 _projectId)
func (walletManager *WalletManager) TryPackGetWalletProjectId(walletId [32]byte) ([]byte, error) {
	return walletManager.abi.Pack("getWalletProjectId", walletId)
}

// UnpackGetWalletProjectId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf8d5d64f.
//
// Solidity: function getWalletProjectId(bytes32 _walletId) view returns(bytes32 _projectId)
func (walletManager *WalletManager) UnpackGetWalletProjectId(data []byte) ([32]byte, error) {
	out, err := walletManager.abi.Unpack("getWalletProjectId", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetWalletStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbfb5dcf9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletStatus(bytes32 _walletId) view returns(uint8 _status)
func (walletManager *WalletManager) PackGetWalletStatus(walletId [32]byte) []byte {
	enc, err := walletManager.abi.Pack("getWalletStatus", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbfb5dcf9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletStatus(bytes32 _walletId) view returns(uint8 _status)
func (walletManager *WalletManager) TryPackGetWalletStatus(walletId [32]byte) ([]byte, error) {
	return walletManager.abi.Pack("getWalletStatus", walletId)
}

// UnpackGetWalletStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 _walletId) view returns(uint8 _status)
func (walletManager *WalletManager) UnpackGetWalletStatus(data []byte) (uint8, error) {
	out, err := walletManager.abi.Unpack("getWalletStatus", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackSetAdmins is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42e1aa9a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setAdmins(bytes32 _walletId, (bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold) returns()
func (walletManager *WalletManager) PackSetAdmins(walletId [32]byte, adminsPublicKeys []PublicKey, adminsThreshold uint64) []byte {
	enc, err := walletManager.abi.Pack("setAdmins", walletId, adminsPublicKeys, adminsThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetAdmins is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x42e1aa9a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setAdmins(bytes32 _walletId, (bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold) returns()
func (walletManager *WalletManager) TryPackSetAdmins(walletId [32]byte, adminsPublicKeys []PublicKey, adminsThreshold uint64) ([]byte, error) {
	return walletManager.abi.Pack("setAdmins", walletId, adminsPublicKeys, adminsThreshold)
}

// PackSetCosigners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e00a2d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setCosigners(bytes32 _walletId, address[] _cosigners, uint64 _cosignersThreshold) returns()
func (walletManager *WalletManager) PackSetCosigners(walletId [32]byte, cosigners []common.Address, cosignersThreshold uint64) []byte {
	enc, err := walletManager.abi.Pack("setCosigners", walletId, cosigners, cosignersThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetCosigners is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e00a2d4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setCosigners(bytes32 _walletId, address[] _cosigners, uint64 _cosignersThreshold) returns()
func (walletManager *WalletManager) TryPackSetCosigners(walletId [32]byte, cosigners []common.Address, cosignersThreshold uint64) ([]byte, error) {
	return walletManager.abi.Pack("setCosigners", walletId, cosigners, cosignersThreshold)
}

// WalletManagerWalletAdminConfirmed represents a WalletAdminConfirmed event raised by the WalletManager contract.
type WalletManagerWalletAdminConfirmed struct {
	WalletId [32]byte
	Admin    common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const WalletManagerWalletAdminConfirmedEventName = "WalletAdminConfirmed"

// ContractEventName returns the user-defined event name.
func (WalletManagerWalletAdminConfirmed) ContractEventName() string {
	return WalletManagerWalletAdminConfirmedEventName
}

// UnpackWalletAdminConfirmedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletAdminConfirmed(bytes32 indexed walletId, address indexed admin)
func (walletManager *WalletManager) UnpackWalletAdminConfirmedEvent(log *types.Log) (*WalletManagerWalletAdminConfirmed, error) {
	event := "WalletAdminConfirmed"
	if len(log.Topics) == 0 || log.Topics[0] != walletManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletManagerWalletAdminConfirmed)
	if len(log.Data) > 0 {
		if err := walletManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletManager.abi.Events[event].Inputs {
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

// WalletManagerWalletAdminsSet represents a WalletAdminsSet event raised by the WalletManager contract.
type WalletManagerWalletAdminsSet struct {
	WalletId         [32]byte
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
	Raw              *types.Log // Blockchain specific contextual infos
}

const WalletManagerWalletAdminsSetEventName = "WalletAdminsSet"

// ContractEventName returns the user-defined event name.
func (WalletManagerWalletAdminsSet) ContractEventName() string {
	return WalletManagerWalletAdminsSetEventName
}

// UnpackWalletAdminsSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletAdminsSet(bytes32 indexed walletId, (bytes32,bytes32)[] adminsPublicKeys, uint64 adminsThreshold)
func (walletManager *WalletManager) UnpackWalletAdminsSetEvent(log *types.Log) (*WalletManagerWalletAdminsSet, error) {
	event := "WalletAdminsSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletManagerWalletAdminsSet)
	if len(log.Data) > 0 {
		if err := walletManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletManager.abi.Events[event].Inputs {
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

// WalletManagerWalletCosignerConfirmed represents a WalletCosignerConfirmed event raised by the WalletManager contract.
type WalletManagerWalletCosignerConfirmed struct {
	WalletId [32]byte
	Cosigner common.Address
	Raw      *types.Log // Blockchain specific contextual infos
}

const WalletManagerWalletCosignerConfirmedEventName = "WalletCosignerConfirmed"

// ContractEventName returns the user-defined event name.
func (WalletManagerWalletCosignerConfirmed) ContractEventName() string {
	return WalletManagerWalletCosignerConfirmedEventName
}

// UnpackWalletCosignerConfirmedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletCosignerConfirmed(bytes32 indexed walletId, address indexed cosigner)
func (walletManager *WalletManager) UnpackWalletCosignerConfirmedEvent(log *types.Log) (*WalletManagerWalletCosignerConfirmed, error) {
	event := "WalletCosignerConfirmed"
	if len(log.Topics) == 0 || log.Topics[0] != walletManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletManagerWalletCosignerConfirmed)
	if len(log.Data) > 0 {
		if err := walletManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletManager.abi.Events[event].Inputs {
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

// WalletManagerWalletCosignersSet represents a WalletCosignersSet event raised by the WalletManager contract.
type WalletManagerWalletCosignersSet struct {
	WalletId           [32]byte
	Cosigners          []common.Address
	CosignersThreshold uint64
	Raw                *types.Log // Blockchain specific contextual infos
}

const WalletManagerWalletCosignersSetEventName = "WalletCosignersSet"

// ContractEventName returns the user-defined event name.
func (WalletManagerWalletCosignersSet) ContractEventName() string {
	return WalletManagerWalletCosignersSetEventName
}

// UnpackWalletCosignersSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletCosignersSet(bytes32 indexed walletId, address[] cosigners, uint64 cosignersThreshold)
func (walletManager *WalletManager) UnpackWalletCosignersSetEvent(log *types.Log) (*WalletManagerWalletCosignersSet, error) {
	event := "WalletCosignersSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletManagerWalletCosignersSet)
	if len(log.Data) > 0 {
		if err := walletManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletManager.abi.Events[event].Inputs {
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

// WalletManagerWalletCreated represents a WalletCreated event raised by the WalletManager contract.
type WalletManagerWalletCreated struct {
	ProjectId [32]byte
	WalletId  [32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletManagerWalletCreatedEventName = "WalletCreated"

// ContractEventName returns the user-defined event name.
func (WalletManagerWalletCreated) ContractEventName() string {
	return WalletManagerWalletCreatedEventName
}

// UnpackWalletCreatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletCreated(bytes32 indexed projectId, bytes32 indexed walletId)
func (walletManager *WalletManager) UnpackWalletCreatedEvent(log *types.Log) (*WalletManagerWalletCreated, error) {
	event := "WalletCreated"
	if len(log.Topics) == 0 || log.Topics[0] != walletManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletManagerWalletCreated)
	if len(log.Data) > 0 {
		if err := walletManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletManager.abi.Events[event].Inputs {
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

// WalletManagerWalletEnabled represents a WalletEnabled event raised by the WalletManager contract.
type WalletManagerWalletEnabled struct {
	WalletId [32]byte
	Raw      *types.Log // Blockchain specific contextual infos
}

const WalletManagerWalletEnabledEventName = "WalletEnabled"

// ContractEventName returns the user-defined event name.
func (WalletManagerWalletEnabled) ContractEventName() string {
	return WalletManagerWalletEnabledEventName
}

// UnpackWalletEnabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletEnabled(bytes32 indexed walletId)
func (walletManager *WalletManager) UnpackWalletEnabledEvent(log *types.Log) (*WalletManagerWalletEnabled, error) {
	event := "WalletEnabled"
	if len(log.Topics) == 0 || log.Topics[0] != walletManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletManagerWalletEnabled)
	if len(log.Data) > 0 {
		if err := walletManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletManager.abi.Events[event].Inputs {
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

// WalletManagerWalletInitialized represents a WalletInitialized event raised by the WalletManager contract.
type WalletManagerWalletInitialized struct {
	WalletId [32]byte
	Raw      *types.Log // Blockchain specific contextual infos
}

const WalletManagerWalletInitializedEventName = "WalletInitialized"

// ContractEventName returns the user-defined event name.
func (WalletManagerWalletInitialized) ContractEventName() string {
	return WalletManagerWalletInitializedEventName
}

// UnpackWalletInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletInitialized(bytes32 indexed walletId)
func (walletManager *WalletManager) UnpackWalletInitializedEvent(log *types.Log) (*WalletManagerWalletInitialized, error) {
	event := "WalletInitialized"
	if len(log.Topics) == 0 || log.Topics[0] != walletManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletManagerWalletInitialized)
	if len(log.Data) > 0 {
		if err := walletManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletManager.abi.Events[event].Inputs {
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
func (walletManager *WalletManager) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], walletManager.abi.Errors["AddressAlreadyInSet"].ID.Bytes()[:4]) {
		return walletManager.UnpackAddressAlreadyInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["AddressNotInSet"].ID.Bytes()[:4]) {
		return walletManager.UnpackAddressNotInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["AdminsNotSet"].ID.Bytes()[:4]) {
		return walletManager.UnpackAdminsNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["AvailabilityCheckTimestampInvalid"].ID.Bytes()[:4]) {
		return walletManager.UnpackAvailabilityCheckTimestampInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["DuplicatedCosigner"].ID.Bytes()[:4]) {
		return walletManager.UnpackDuplicatedCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["DuplicatedPublicKey"].ID.Bytes()[:4]) {
		return walletManager.UnpackDuplicatedPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return walletManager.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidAdmin"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidAdminError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidAdminPublicKey"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidAdminPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidAdminsThreshold"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidAdminsThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidCosignersThreshold"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidCosignersThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidKeyType"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidPublicKey"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidResponseData"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidResponseDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidSigningAlgo"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidSigningAlgoError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["InvalidWalletStatus"].ID.Bytes()[:4]) {
		return walletManager.UnpackInvalidWalletStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["KeyTypeNotSupported"].ID.Bytes()[:4]) {
		return walletManager.UnpackKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return walletManager.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["MultisigThresholdNotSet"].ID.Bytes()[:4]) {
		return walletManager.UnpackMultisigThresholdNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["NoAddresses"].ID.Bytes()[:4]) {
		return walletManager.UnpackNoAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["NotAllAdminsConfirmed"].ID.Bytes()[:4]) {
		return walletManager.UnpackNotAllAdminsConfirmedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["NotAllCosignersConfirmed"].ID.Bytes()[:4]) {
		return walletManager.UnpackNotAllCosignersConfirmedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["NotEnoughAdmins"].ID.Bytes()[:4]) {
		return walletManager.UnpackNotEnoughAdminsError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["NotEnoughKeys"].ID.Bytes()[:4]) {
		return walletManager.UnpackNotEnoughKeysError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["NotOwnerOrPauser"].ID.Bytes()[:4]) {
		return walletManager.UnpackNotOwnerOrPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["NotOwnerOrUnpauser"].ID.Bytes()[:4]) {
		return walletManager.UnpackNotOwnerOrUnpauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["OnlyExtensionOwner"].ID.Bytes()[:4]) {
		return walletManager.UnpackOnlyExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return walletManager.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return walletManager.UnpackOnlyOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["OnlyOwnerOrBackupManager"].ID.Bytes()[:4]) {
		return walletManager.UnpackOnlyOwnerOrBackupManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return walletManager.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["OnlyProposedOwner"].ID.Bytes()[:4]) {
		return walletManager.UnpackOnlyProposedOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["OwnerNotAllowed"].ID.Bytes()[:4]) {
		return walletManager.UnpackOwnerNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return walletManager.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["TooManyAdmins"].ID.Bytes()[:4]) {
		return walletManager.UnpackTooManyAdminsError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["TooManyCosigners"].ID.Bytes()[:4]) {
		return walletManager.UnpackTooManyCosignersError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletManager.abi.Errors["VersionNotSupported"].ID.Bytes()[:4]) {
		return walletManager.UnpackVersionNotSupportedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// WalletManagerAddressAlreadyInSet represents a AddressAlreadyInSet error raised by the WalletManager contract.
type WalletManagerAddressAlreadyInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressAlreadyInSet(address addr)
func WalletManagerAddressAlreadyInSetErrorID() common.Hash {
	return common.HexToHash("0xb725c7fb1c3f0f065973d6b9eb55140f970259955af3045742cd5b37902eb920")
}

// UnpackAddressAlreadyInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressAlreadyInSet(address addr)
func (walletManager *WalletManager) UnpackAddressAlreadyInSetError(raw []byte) (*WalletManagerAddressAlreadyInSet, error) {
	out := new(WalletManagerAddressAlreadyInSet)
	if err := walletManager.abi.UnpackIntoInterface(out, "AddressAlreadyInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerAddressNotInSet represents a AddressNotInSet error raised by the WalletManager contract.
type WalletManagerAddressNotInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressNotInSet(address addr)
func WalletManagerAddressNotInSetErrorID() common.Hash {
	return common.HexToHash("0x1ff5ecf3e6d385ea01ff0323fe894ed6d25a83ed5f6ff2a2ede4c897408f6be9")
}

// UnpackAddressNotInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressNotInSet(address addr)
func (walletManager *WalletManager) UnpackAddressNotInSetError(raw []byte) (*WalletManagerAddressNotInSet, error) {
	out := new(WalletManagerAddressNotInSet)
	if err := walletManager.abi.UnpackIntoInterface(out, "AddressNotInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerAdminsNotSet represents a AdminsNotSet error raised by the WalletManager contract.
type WalletManagerAdminsNotSet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AdminsNotSet()
func WalletManagerAdminsNotSetErrorID() common.Hash {
	return common.HexToHash("0x7d64e8b6547bfafa2a1fd13b4bec4139914c0e527bfb48c8ff0e7c9fc1e95571")
}

// UnpackAdminsNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AdminsNotSet()
func (walletManager *WalletManager) UnpackAdminsNotSetError(raw []byte) (*WalletManagerAdminsNotSet, error) {
	out := new(WalletManagerAdminsNotSet)
	if err := walletManager.abi.UnpackIntoInterface(out, "AdminsNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerAvailabilityCheckTimestampInvalid represents a AvailabilityCheckTimestampInvalid error raised by the WalletManager contract.
type WalletManagerAvailabilityCheckTimestampInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func WalletManagerAvailabilityCheckTimestampInvalidErrorID() common.Hash {
	return common.HexToHash("0x6344a5e74cff78612fa271fa907ca608af3bad10e1e8c8df302e29ebc6298e2d")
}

// UnpackAvailabilityCheckTimestampInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func (walletManager *WalletManager) UnpackAvailabilityCheckTimestampInvalidError(raw []byte) (*WalletManagerAvailabilityCheckTimestampInvalid, error) {
	out := new(WalletManagerAvailabilityCheckTimestampInvalid)
	if err := walletManager.abi.UnpackIntoInterface(out, "AvailabilityCheckTimestampInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerDuplicatedCosigner represents a DuplicatedCosigner error raised by the WalletManager contract.
type WalletManagerDuplicatedCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func WalletManagerDuplicatedCosignerErrorID() common.Hash {
	return common.HexToHash("0x53a5890417576fb3043dc3ff1669b0932da2f657d8adb4c11313b12b711f5910")
}

// UnpackDuplicatedCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func (walletManager *WalletManager) UnpackDuplicatedCosignerError(raw []byte) (*WalletManagerDuplicatedCosigner, error) {
	out := new(WalletManagerDuplicatedCosigner)
	if err := walletManager.abi.UnpackIntoInterface(out, "DuplicatedCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerDuplicatedPublicKey represents a DuplicatedPublicKey error raised by the WalletManager contract.
type WalletManagerDuplicatedPublicKey struct {
	PublicKey PublicKey
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedPublicKey((bytes32,bytes32) publicKey)
func WalletManagerDuplicatedPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xe616ec728dba41c6b7b5e093a3d6791c994ae0db55876c565734bc265c2486e9")
}

// UnpackDuplicatedPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedPublicKey((bytes32,bytes32) publicKey)
func (walletManager *WalletManager) UnpackDuplicatedPublicKeyError(raw []byte) (*WalletManagerDuplicatedPublicKey, error) {
	out := new(WalletManagerDuplicatedPublicKey)
	if err := walletManager.abi.UnpackIntoInterface(out, "DuplicatedPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerExtensionIdMismatch represents a ExtensionIdMismatch error raised by the WalletManager contract.
type WalletManagerExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func WalletManagerExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (walletManager *WalletManager) UnpackExtensionIdMismatchError(raw []byte) (*WalletManagerExtensionIdMismatch, error) {
	out := new(WalletManagerExtensionIdMismatch)
	if err := walletManager.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidAddress represents a InvalidAddress error raised by the WalletManager contract.
type WalletManagerInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func WalletManagerInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (walletManager *WalletManager) UnpackInvalidAddressError(raw []byte) (*WalletManagerInvalidAddress, error) {
	out := new(WalletManagerInvalidAddress)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidAdmin represents a InvalidAdmin error raised by the WalletManager contract.
type WalletManagerInvalidAdmin struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAdmin()
func WalletManagerInvalidAdminErrorID() common.Hash {
	return common.HexToHash("0xb5eba9f01c7cbc6d71522850c06349630c7deaed3ac3e944712461ed3625f10e")
}

// UnpackInvalidAdminError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAdmin()
func (walletManager *WalletManager) UnpackInvalidAdminError(raw []byte) (*WalletManagerInvalidAdmin, error) {
	out := new(WalletManagerInvalidAdmin)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidAdmin", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidAdminPublicKey represents a InvalidAdminPublicKey error raised by the WalletManager contract.
type WalletManagerInvalidAdminPublicKey struct {
	PublicKey PublicKey
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAdminPublicKey((bytes32,bytes32) publicKey)
func WalletManagerInvalidAdminPublicKeyErrorID() common.Hash {
	return common.HexToHash("0x944808702f4427fb6cc5e8012427fb5e8f5e479fc0d2e1fb223aa11a23f68455")
}

// UnpackInvalidAdminPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAdminPublicKey((bytes32,bytes32) publicKey)
func (walletManager *WalletManager) UnpackInvalidAdminPublicKeyError(raw []byte) (*WalletManagerInvalidAdminPublicKey, error) {
	out := new(WalletManagerInvalidAdminPublicKey)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidAdminPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidAdminsThreshold represents a InvalidAdminsThreshold error raised by the WalletManager contract.
type WalletManagerInvalidAdminsThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAdminsThreshold()
func WalletManagerInvalidAdminsThresholdErrorID() common.Hash {
	return common.HexToHash("0x35488e9941655d1f8a104bd6ba67445eb59fe2e353c374e725d44d1de9440b48")
}

// UnpackInvalidAdminsThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAdminsThreshold()
func (walletManager *WalletManager) UnpackInvalidAdminsThresholdError(raw []byte) (*WalletManagerInvalidAdminsThreshold, error) {
	out := new(WalletManagerInvalidAdminsThreshold)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidAdminsThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidAvailabilityCheckStatus represents a InvalidAvailabilityCheckStatus error raised by the WalletManager contract.
type WalletManagerInvalidAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func WalletManagerInvalidAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0xedee4ff00b842d3ef84031c60001065c540c838c5479bbf0ae56f15868427e44")
}

// UnpackInvalidAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func (walletManager *WalletManager) UnpackInvalidAvailabilityCheckStatusError(raw []byte) (*WalletManagerInvalidAvailabilityCheckStatus, error) {
	out := new(WalletManagerInvalidAvailabilityCheckStatus)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidCosigner represents a InvalidCosigner error raised by the WalletManager contract.
type WalletManagerInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func WalletManagerInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (walletManager *WalletManager) UnpackInvalidCosignerError(raw []byte) (*WalletManagerInvalidCosigner, error) {
	out := new(WalletManagerInvalidCosigner)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidCosignersThreshold represents a InvalidCosignersThreshold error raised by the WalletManager contract.
type WalletManagerInvalidCosignersThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosignersThreshold()
func WalletManagerInvalidCosignersThresholdErrorID() common.Hash {
	return common.HexToHash("0x00c60d44e7ecea6f0cb356f1b262cbabe7b56b832c62a6f8543adc4440bc8f12")
}

// UnpackInvalidCosignersThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosignersThreshold()
func (walletManager *WalletManager) UnpackInvalidCosignersThresholdError(raw []byte) (*WalletManagerInvalidCosignersThreshold, error) {
	out := new(WalletManagerInvalidCosignersThreshold)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidCosignersThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidDuration represents a InvalidDuration error raised by the WalletManager contract.
type WalletManagerInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func WalletManagerInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (walletManager *WalletManager) UnpackInvalidDurationError(raw []byte) (*WalletManagerInvalidDuration, error) {
	out := new(WalletManagerInvalidDuration)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the WalletManager contract.
type WalletManagerInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func WalletManagerInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (walletManager *WalletManager) UnpackInvalidGovernanceHashError(raw []byte) (*WalletManagerInvalidGovernanceHash, error) {
	out := new(WalletManagerInvalidGovernanceHash)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidKeyType represents a InvalidKeyType error raised by the WalletManager contract.
type WalletManagerInvalidKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyType()
func WalletManagerInvalidKeyTypeErrorID() common.Hash {
	return common.HexToHash("0x36dccb440beb343622d8ee0927bce3f69281feea3a5b0610f6e00e2811b557e1")
}

// UnpackInvalidKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyType()
func (walletManager *WalletManager) UnpackInvalidKeyTypeError(raw []byte) (*WalletManagerInvalidKeyType, error) {
	out := new(WalletManagerInvalidKeyType)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidNonce represents a InvalidNonce error raised by the WalletManager contract.
type WalletManagerInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func WalletManagerInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (walletManager *WalletManager) UnpackInvalidNonceError(raw []byte) (*WalletManagerInvalidNonce, error) {
	out := new(WalletManagerInvalidNonce)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidPublicKey represents a InvalidPublicKey error raised by the WalletManager contract.
type WalletManagerInvalidPublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPublicKey()
func WalletManagerInvalidPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa2d0fee854985cb3d3a394146b5df1c41d5b4d6fcede7fc57a3c05a49b2f3e73")
}

// UnpackInvalidPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPublicKey()
func (walletManager *WalletManager) UnpackInvalidPublicKeyError(raw []byte) (*WalletManagerInvalidPublicKey, error) {
	out := new(WalletManagerInvalidPublicKey)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidResponseData represents a InvalidResponseData error raised by the WalletManager contract.
type WalletManagerInvalidResponseData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseData()
func WalletManagerInvalidResponseDataErrorID() common.Hash {
	return common.HexToHash("0x48566617128355a033880805d6462da9bf9c6fe4c61b922997451a71262990af")
}

// UnpackInvalidResponseDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseData()
func (walletManager *WalletManager) UnpackInvalidResponseDataError(raw []byte) (*WalletManagerInvalidResponseData, error) {
	out := new(WalletManagerInvalidResponseData)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidResponseData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidSigningAlgo represents a InvalidSigningAlgo error raised by the WalletManager contract.
type WalletManagerInvalidSigningAlgo struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningAlgo()
func WalletManagerInvalidSigningAlgoErrorID() common.Hash {
	return common.HexToHash("0x9b7198a697b61e546ef9ee593a6f77d2bf9f1df9dcbba655ce96d7e2c75346c7")
}

// UnpackInvalidSigningAlgoError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningAlgo()
func (walletManager *WalletManager) UnpackInvalidSigningAlgoError(raw []byte) (*WalletManagerInvalidSigningAlgo, error) {
	out := new(WalletManagerInvalidSigningAlgo)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidSigningAlgo", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidThreshold represents a InvalidThreshold error raised by the WalletManager contract.
type WalletManagerInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func WalletManagerInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (walletManager *WalletManager) UnpackInvalidThresholdError(raw []byte) (*WalletManagerInvalidThreshold, error) {
	out := new(WalletManagerInvalidThreshold)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerInvalidWalletStatus represents a InvalidWalletStatus error raised by the WalletManager contract.
type WalletManagerInvalidWalletStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidWalletStatus()
func WalletManagerInvalidWalletStatusErrorID() common.Hash {
	return common.HexToHash("0xd700120e4896ea7835a5e9e4feb690644ced5cad1c52209afdcbe06b96d91345")
}

// UnpackInvalidWalletStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidWalletStatus()
func (walletManager *WalletManager) UnpackInvalidWalletStatusError(raw []byte) (*WalletManagerInvalidWalletStatus, error) {
	out := new(WalletManagerInvalidWalletStatus)
	if err := walletManager.abi.UnpackIntoInterface(out, "InvalidWalletStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerKeyTypeNotSupported represents a KeyTypeNotSupported error raised by the WalletManager contract.
type WalletManagerKeyTypeNotSupported struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func WalletManagerKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x85004df1f47814c2de6fbf03d3f2b1ca8c67c182d0d54278766543b6be3bc603")
}

// UnpackKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func (walletManager *WalletManager) UnpackKeyTypeNotSupportedError(raw []byte) (*WalletManagerKeyTypeNotSupported, error) {
	out := new(WalletManagerKeyTypeNotSupported)
	if err := walletManager.abi.UnpackIntoInterface(out, "KeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerLengthsMismatch represents a LengthsMismatch error raised by the WalletManager contract.
type WalletManagerLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func WalletManagerLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (walletManager *WalletManager) UnpackLengthsMismatchError(raw []byte) (*WalletManagerLengthsMismatch, error) {
	out := new(WalletManagerLengthsMismatch)
	if err := walletManager.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerMultisigThresholdNotSet represents a MultisigThresholdNotSet error raised by the WalletManager contract.
type WalletManagerMultisigThresholdNotSet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MultisigThresholdNotSet()
func WalletManagerMultisigThresholdNotSetErrorID() common.Hash {
	return common.HexToHash("0xa754d3ded2abb7c4de5671addd714e8219d8ecdfa97d7ba66229a03266256634")
}

// UnpackMultisigThresholdNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MultisigThresholdNotSet()
func (walletManager *WalletManager) UnpackMultisigThresholdNotSetError(raw []byte) (*WalletManagerMultisigThresholdNotSet, error) {
	out := new(WalletManagerMultisigThresholdNotSet)
	if err := walletManager.abi.UnpackIntoInterface(out, "MultisigThresholdNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerNoAddresses represents a NoAddresses error raised by the WalletManager contract.
type WalletManagerNoAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoAddresses()
func WalletManagerNoAddressesErrorID() common.Hash {
	return common.HexToHash("0xb2e79f05746fa6c9f3d61f9efe85959371cde5030a6c2dd70c92b63578c5cde7")
}

// UnpackNoAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoAddresses()
func (walletManager *WalletManager) UnpackNoAddressesError(raw []byte) (*WalletManagerNoAddresses, error) {
	out := new(WalletManagerNoAddresses)
	if err := walletManager.abi.UnpackIntoInterface(out, "NoAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerNotAllAdminsConfirmed represents a NotAllAdminsConfirmed error raised by the WalletManager contract.
type WalletManagerNotAllAdminsConfirmed struct {
	Admin common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAllAdminsConfirmed(address admin)
func WalletManagerNotAllAdminsConfirmedErrorID() common.Hash {
	return common.HexToHash("0xd9b8524adf7b8a09fa04589a8fa7e76710132c3707141a64a595fdce0f94326a")
}

// UnpackNotAllAdminsConfirmedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAllAdminsConfirmed(address admin)
func (walletManager *WalletManager) UnpackNotAllAdminsConfirmedError(raw []byte) (*WalletManagerNotAllAdminsConfirmed, error) {
	out := new(WalletManagerNotAllAdminsConfirmed)
	if err := walletManager.abi.UnpackIntoInterface(out, "NotAllAdminsConfirmed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerNotAllCosignersConfirmed represents a NotAllCosignersConfirmed error raised by the WalletManager contract.
type WalletManagerNotAllCosignersConfirmed struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAllCosignersConfirmed(address cosigner)
func WalletManagerNotAllCosignersConfirmedErrorID() common.Hash {
	return common.HexToHash("0x944205148b694a955b67ee577441d714d54bc39606eb44f1eed7cec0cf838bfb")
}

// UnpackNotAllCosignersConfirmedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAllCosignersConfirmed(address cosigner)
func (walletManager *WalletManager) UnpackNotAllCosignersConfirmedError(raw []byte) (*WalletManagerNotAllCosignersConfirmed, error) {
	out := new(WalletManagerNotAllCosignersConfirmed)
	if err := walletManager.abi.UnpackIntoInterface(out, "NotAllCosignersConfirmed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerNotEnoughAdmins represents a NotEnoughAdmins error raised by the WalletManager contract.
type WalletManagerNotEnoughAdmins struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotEnoughAdmins()
func WalletManagerNotEnoughAdminsErrorID() common.Hash {
	return common.HexToHash("0xff7e6cf8335a6d16a1d23eaea73973afbc302691980067139e5635a738bf8e13")
}

// UnpackNotEnoughAdminsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotEnoughAdmins()
func (walletManager *WalletManager) UnpackNotEnoughAdminsError(raw []byte) (*WalletManagerNotEnoughAdmins, error) {
	out := new(WalletManagerNotEnoughAdmins)
	if err := walletManager.abi.UnpackIntoInterface(out, "NotEnoughAdmins", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerNotEnoughKeys represents a NotEnoughKeys error raised by the WalletManager contract.
type WalletManagerNotEnoughKeys struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotEnoughKeys()
func WalletManagerNotEnoughKeysErrorID() common.Hash {
	return common.HexToHash("0x2f27a9543372f42ae13dc175d18b0b075bad70c284b7919a973990e9f1fe36c6")
}

// UnpackNotEnoughKeysError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotEnoughKeys()
func (walletManager *WalletManager) UnpackNotEnoughKeysError(raw []byte) (*WalletManagerNotEnoughKeys, error) {
	out := new(WalletManagerNotEnoughKeys)
	if err := walletManager.abi.UnpackIntoInterface(out, "NotEnoughKeys", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerNotOwnerOrPauser represents a NotOwnerOrPauser error raised by the WalletManager contract.
type WalletManagerNotOwnerOrPauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrPauser(address caller)
func WalletManagerNotOwnerOrPauserErrorID() common.Hash {
	return common.HexToHash("0x46c0d9af6b895093e022ad2984c54e60192e6e26c5063daf232e825c6e07ff9a")
}

// UnpackNotOwnerOrPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrPauser(address caller)
func (walletManager *WalletManager) UnpackNotOwnerOrPauserError(raw []byte) (*WalletManagerNotOwnerOrPauser, error) {
	out := new(WalletManagerNotOwnerOrPauser)
	if err := walletManager.abi.UnpackIntoInterface(out, "NotOwnerOrPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerNotOwnerOrUnpauser represents a NotOwnerOrUnpauser error raised by the WalletManager contract.
type WalletManagerNotOwnerOrUnpauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func WalletManagerNotOwnerOrUnpauserErrorID() common.Hash {
	return common.HexToHash("0x81e700e8f0bd5ba4a6d6a2ce9a1ad0ac9bb1b1b1e62d921a55cbc4b53f1da649")
}

// UnpackNotOwnerOrUnpauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func (walletManager *WalletManager) UnpackNotOwnerOrUnpauserError(raw []byte) (*WalletManagerNotOwnerOrUnpauser, error) {
	out := new(WalletManagerNotOwnerOrUnpauser)
	if err := walletManager.abi.UnpackIntoInterface(out, "NotOwnerOrUnpauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerOnlyExtensionOwner represents a OnlyExtensionOwner error raised by the WalletManager contract.
type WalletManagerOnlyExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwner()
func WalletManagerOnlyExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xcd49fd1de5a297cd1ce055d5abffcf2cb25ad9492bd39a703a238af57b2a9488")
}

// UnpackOnlyExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwner()
func (walletManager *WalletManager) UnpackOnlyExtensionOwnerError(raw []byte) (*WalletManagerOnlyExtensionOwner, error) {
	out := new(WalletManagerOnlyExtensionOwner)
	if err := walletManager.abi.UnpackIntoInterface(out, "OnlyExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the WalletManager contract.
type WalletManagerOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func WalletManagerOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (walletManager *WalletManager) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*WalletManagerOnlyExtensionOwnerOrOperator, error) {
	out := new(WalletManagerOnlyExtensionOwnerOrOperator)
	if err := walletManager.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerOnlyOwner represents a OnlyOwner error raised by the WalletManager contract.
type WalletManagerOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func WalletManagerOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (walletManager *WalletManager) UnpackOnlyOwnerError(raw []byte) (*WalletManagerOnlyOwner, error) {
	out := new(WalletManagerOnlyOwner)
	if err := walletManager.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerOnlyOwnerOrBackupManager represents a OnlyOwnerOrBackupManager error raised by the WalletManager contract.
type WalletManagerOnlyOwnerOrBackupManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrBackupManager()
func WalletManagerOnlyOwnerOrBackupManagerErrorID() common.Hash {
	return common.HexToHash("0x249a6cf6e22092f9c80798e2a0c583c0f36aea0779ef38a5b5cdb6a112a0e6df")
}

// UnpackOnlyOwnerOrBackupManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrBackupManager()
func (walletManager *WalletManager) UnpackOnlyOwnerOrBackupManagerError(raw []byte) (*WalletManagerOnlyOwnerOrBackupManager, error) {
	out := new(WalletManagerOnlyOwnerOrBackupManager)
	if err := walletManager.abi.UnpackIntoInterface(out, "OnlyOwnerOrBackupManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the WalletManager contract.
type WalletManagerOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func WalletManagerOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (walletManager *WalletManager) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*WalletManagerOnlyProductionOrPausedStatus, error) {
	out := new(WalletManagerOnlyProductionOrPausedStatus)
	if err := walletManager.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerOnlyProposedOwner represents a OnlyProposedOwner error raised by the WalletManager contract.
type WalletManagerOnlyProposedOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProposedOwner()
func WalletManagerOnlyProposedOwnerErrorID() common.Hash {
	return common.HexToHash("0x49cb15061ee558e74f66f937f740fd2a3791a89d2d139862222e5c97cabc02cd")
}

// UnpackOnlyProposedOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProposedOwner()
func (walletManager *WalletManager) UnpackOnlyProposedOwnerError(raw []byte) (*WalletManagerOnlyProposedOwner, error) {
	out := new(WalletManagerOnlyProposedOwner)
	if err := walletManager.abi.UnpackIntoInterface(out, "OnlyProposedOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerOwnerNotAllowed represents a OwnerNotAllowed error raised by the WalletManager contract.
type WalletManagerOwnerNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerNotAllowed()
func WalletManagerOwnerNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x3d5c9996545c6bf4287c26c1c2d8e6b4c39b92f19ba706c3a756720bf9c4fd24")
}

// UnpackOwnerNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerNotAllowed()
func (walletManager *WalletManager) UnpackOwnerNotAllowedError(raw []byte) (*WalletManagerOwnerNotAllowed, error) {
	out := new(WalletManagerOwnerNotAllowed)
	if err := walletManager.abi.UnpackIntoInterface(out, "OwnerNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the WalletManager contract.
type WalletManagerTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func WalletManagerTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (walletManager *WalletManager) UnpackTeeMachineNotAvailableError(raw []byte) (*WalletManagerTeeMachineNotAvailable, error) {
	out := new(WalletManagerTeeMachineNotAvailable)
	if err := walletManager.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerTooManyAdmins represents a TooManyAdmins error raised by the WalletManager contract.
type WalletManagerTooManyAdmins struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TooManyAdmins()
func WalletManagerTooManyAdminsErrorID() common.Hash {
	return common.HexToHash("0x775f8d17413938ad665d73a714e3cefb0ea5bbe1f3ac8048560dfb59e710fd45")
}

// UnpackTooManyAdminsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TooManyAdmins()
func (walletManager *WalletManager) UnpackTooManyAdminsError(raw []byte) (*WalletManagerTooManyAdmins, error) {
	out := new(WalletManagerTooManyAdmins)
	if err := walletManager.abi.UnpackIntoInterface(out, "TooManyAdmins", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerTooManyCosigners represents a TooManyCosigners error raised by the WalletManager contract.
type WalletManagerTooManyCosigners struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TooManyCosigners()
func WalletManagerTooManyCosignersErrorID() common.Hash {
	return common.HexToHash("0x81654976590dafc3dfaf9a3e6dfd333f8ef1a1f66061a2340eeedb0a194dffbe")
}

// UnpackTooManyCosignersError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TooManyCosigners()
func (walletManager *WalletManager) UnpackTooManyCosignersError(raw []byte) (*WalletManagerTooManyCosigners, error) {
	out := new(WalletManagerTooManyCosigners)
	if err := walletManager.abi.UnpackIntoInterface(out, "TooManyCosigners", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletManagerVersionNotSupported represents a VersionNotSupported error raised by the WalletManager contract.
type WalletManagerVersionNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionNotSupported()
func WalletManagerVersionNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x6db0b17b06266b37e721024a8ba1221439799ebb779dfd397865e37250abc9a8")
}

// UnpackVersionNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionNotSupported()
func (walletManager *WalletManager) UnpackVersionNotSupportedError(raw []byte) (*WalletManagerVersionNotSupported, error) {
	out := new(WalletManagerVersionNotSupported)
	if err := walletManager.abi.UnpackIntoInterface(out, "VersionNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}
