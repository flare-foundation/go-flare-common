// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package extensionmanager

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

// ExtensionManagerMetaData contains all meta data concerning the ExtensionManager contract.
var ExtensionManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CodeHashPlatformAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CodeHashZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCodeHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInstructionsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPlatform\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReservedExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyTypeEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoKeyTypes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPlatforms\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"NoSigningAlgos\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"PlatformAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PlatformEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"PlatformNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReservedExtensionIdAlreadyAssigned\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"}],\"name\":\"SigningAlgoAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SigningAlgoEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"}],\"name\":\"SigningAlgoNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SystemOwnedExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"UnsupportedPlatform\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"platforms\",\"type\":\"bytes32[]\"}],\"name\":\"CodeHashPlatformsDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"ExtensionOperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTypes\",\"type\":\"bytes32[]\"}],\"name\":\"SupportedKeyTypesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTypes\",\"type\":\"bytes32[]\"}],\"name\":\"SupportedKeyTypesRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTypes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[][]\",\"name\":\"signingAlgosByKeyType\",\"type\":\"bytes32[][]\"}],\"name\":\"SystemSupportedKeyTypesAndSigningAlgosAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTypes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[][]\",\"name\":\"signingAlgosByKeyType\",\"type\":\"bytes32[][]\"}],\"name\":\"SystemSupportedKeyTypesAndSigningAlgosRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"platforms\",\"type\":\"bytes32[]\"}],\"name\":\"SystemSupportedPlatformsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"platforms\",\"type\":\"bytes32[]\"}],\"name\":\"SystemSupportedPlatformsRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"teeExtensionStateVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"TeeExtensionContractsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"TeeExtensionRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIMachineManager.TeeStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"TeeMachineStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"platforms\",\"type\":\"bytes32[]\"}],\"name\":\"TeeVersionAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_keyTypes\",\"type\":\"bytes32[]\"}],\"name\":\"addSupportedKeyTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_keyTypes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_signingAlgosByKeyType\",\"type\":\"bytes32[][]\"}],\"name\":\"addSystemSupportedKeyTypesAndSigningAlgos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"}],\"name\":\"addSystemSupportedPlatforms\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"}],\"name\":\"addTeeVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"}],\"name\":\"disableCodeHashPlatforms\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"}],\"name\":\"getCodeHashInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getExtensionOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getExtensionOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getSupportedCodeHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_supportedCodeHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getSupportedKeyTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_supportedKeyTypes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSystemSupportedKeyTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSystemSupportedPlatforms\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"}],\"name\":\"getSystemSupportedSigningAlgos\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getTeeExtensionInstructionsSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getTeeExtensionStateVerifier\",\"outputs\":[{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"isCodeHashPlatformDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"isCodeHashPlatformSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"}],\"name\":\"isKeyTypeSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_signingAlgo\",\"type\":\"bytes32\"}],\"name\":\"isSigningAlgoSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextPublicExtensionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"_teeExtensionStateVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"registerReserved\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_keyTypes\",\"type\":\"bytes32[]\"}],\"name\":\"removeSupportedKeyTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_keyTypes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_signingAlgosByKeyType\",\"type\":\"bytes32[][]\"}],\"name\":\"removeSystemSupportedKeyTypesAndSigningAlgos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"}],\"name\":\"removeSystemSupportedPlatforms\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"_teeExtensionStateVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"setExtensionContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setExtensionOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "ExtensionManager",
}

// ExtensionManager is an auto generated Go binding around an Ethereum contract.
type ExtensionManager struct {
	abi abi.ABI
}

// NewExtensionManager creates a new instance of ExtensionManager.
func NewExtensionManager() *ExtensionManager {
	parsed, err := ExtensionManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &ExtensionManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *ExtensionManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddSupportedKeyTypes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x174c5e26.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (extensionManager *ExtensionManager) PackAddSupportedKeyTypes(extensionId *big.Int, keyTypes [][32]byte) []byte {
	enc, err := extensionManager.abi.Pack("addSupportedKeyTypes", extensionId, keyTypes)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddSupportedKeyTypes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x174c5e26.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (extensionManager *ExtensionManager) TryPackAddSupportedKeyTypes(extensionId *big.Int, keyTypes [][32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("addSupportedKeyTypes", extensionId, keyTypes)
}

// PackAddSystemSupportedKeyTypesAndSigningAlgos is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38c77e46.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addSystemSupportedKeyTypesAndSigningAlgos(bytes32[] _keyTypes, bytes32[][] _signingAlgosByKeyType) returns()
func (extensionManager *ExtensionManager) PackAddSystemSupportedKeyTypesAndSigningAlgos(keyTypes [][32]byte, signingAlgosByKeyType [][][32]byte) []byte {
	enc, err := extensionManager.abi.Pack("addSystemSupportedKeyTypesAndSigningAlgos", keyTypes, signingAlgosByKeyType)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddSystemSupportedKeyTypesAndSigningAlgos is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x38c77e46.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addSystemSupportedKeyTypesAndSigningAlgos(bytes32[] _keyTypes, bytes32[][] _signingAlgosByKeyType) returns()
func (extensionManager *ExtensionManager) TryPackAddSystemSupportedKeyTypesAndSigningAlgos(keyTypes [][32]byte, signingAlgosByKeyType [][][32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("addSystemSupportedKeyTypesAndSigningAlgos", keyTypes, signingAlgosByKeyType)
}

// PackAddSystemSupportedPlatforms is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0c14a66b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addSystemSupportedPlatforms(bytes32[] _platforms) returns()
func (extensionManager *ExtensionManager) PackAddSystemSupportedPlatforms(platforms [][32]byte) []byte {
	enc, err := extensionManager.abi.Pack("addSystemSupportedPlatforms", platforms)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddSystemSupportedPlatforms is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0c14a66b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addSystemSupportedPlatforms(bytes32[] _platforms) returns()
func (extensionManager *ExtensionManager) TryPackAddSystemSupportedPlatforms(platforms [][32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("addSystemSupportedPlatforms", platforms)
}

// PackAddTeeVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2cf2ca0e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addTeeVersion(uint256 _extensionId, bytes32 _version, bytes32 _codeHash, bytes32[] _platforms) returns()
func (extensionManager *ExtensionManager) PackAddTeeVersion(extensionId *big.Int, version [32]byte, codeHash [32]byte, platforms [][32]byte) []byte {
	enc, err := extensionManager.abi.Pack("addTeeVersion", extensionId, version, codeHash, platforms)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddTeeVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2cf2ca0e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addTeeVersion(uint256 _extensionId, bytes32 _version, bytes32 _codeHash, bytes32[] _platforms) returns()
func (extensionManager *ExtensionManager) TryPackAddTeeVersion(extensionId *big.Int, version [32]byte, codeHash [32]byte, platforms [][32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("addTeeVersion", extensionId, version, codeHash, platforms)
}

// PackConfirmOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x290fbb03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmOwnership(uint256 _extensionId) returns()
func (extensionManager *ExtensionManager) PackConfirmOwnership(extensionId *big.Int) []byte {
	enc, err := extensionManager.abi.Pack("confirmOwnership", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x290fbb03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmOwnership(uint256 _extensionId) returns()
func (extensionManager *ExtensionManager) TryPackConfirmOwnership(extensionId *big.Int) ([]byte, error) {
	return extensionManager.abi.Pack("confirmOwnership", extensionId)
}

// PackDisableCodeHashPlatforms is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc7ef55f8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function disableCodeHashPlatforms(uint256 _extensionId, bytes32 _codeHash, bytes32[] _platforms) returns()
func (extensionManager *ExtensionManager) PackDisableCodeHashPlatforms(extensionId *big.Int, codeHash [32]byte, platforms [][32]byte) []byte {
	enc, err := extensionManager.abi.Pack("disableCodeHashPlatforms", extensionId, codeHash, platforms)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDisableCodeHashPlatforms is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc7ef55f8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function disableCodeHashPlatforms(uint256 _extensionId, bytes32 _codeHash, bytes32[] _platforms) returns()
func (extensionManager *ExtensionManager) TryPackDisableCodeHashPlatforms(extensionId *big.Int, codeHash [32]byte, platforms [][32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("disableCodeHashPlatforms", extensionId, codeHash, platforms)
}

// PackGetCodeHashInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62672305.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCodeHashInfo(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32 _version, bytes32[] _platforms)
func (extensionManager *ExtensionManager) PackGetCodeHashInfo(extensionId *big.Int, codeHash [32]byte) []byte {
	enc, err := extensionManager.abi.Pack("getCodeHashInfo", extensionId, codeHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCodeHashInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62672305.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCodeHashInfo(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32 _version, bytes32[] _platforms)
func (extensionManager *ExtensionManager) TryPackGetCodeHashInfo(extensionId *big.Int, codeHash [32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("getCodeHashInfo", extensionId, codeHash)
}

// GetCodeHashInfoOutput serves as a container for the return parameters of contract
// method GetCodeHashInfo.
type GetCodeHashInfoOutput struct {
	Version   [32]byte
	Platforms [][32]byte
}

// UnpackGetCodeHashInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62672305.
//
// Solidity: function getCodeHashInfo(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32 _version, bytes32[] _platforms)
func (extensionManager *ExtensionManager) UnpackGetCodeHashInfo(data []byte) (GetCodeHashInfoOutput, error) {
	out, err := extensionManager.abi.Unpack("getCodeHashInfo", data)
	outstruct := new(GetCodeHashInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Version = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Platforms = *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)
	return *outstruct, nil
}

// PackGetExtensionOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb69caaca.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getExtensionOperator(uint256 _extensionId) view returns(address)
func (extensionManager *ExtensionManager) PackGetExtensionOperator(extensionId *big.Int) []byte {
	enc, err := extensionManager.abi.Pack("getExtensionOperator", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetExtensionOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb69caaca.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getExtensionOperator(uint256 _extensionId) view returns(address)
func (extensionManager *ExtensionManager) TryPackGetExtensionOperator(extensionId *big.Int) ([]byte, error) {
	return extensionManager.abi.Pack("getExtensionOperator", extensionId)
}

// UnpackGetExtensionOperator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb69caaca.
//
// Solidity: function getExtensionOperator(uint256 _extensionId) view returns(address)
func (extensionManager *ExtensionManager) UnpackGetExtensionOperator(data []byte) (common.Address, error) {
	out, err := extensionManager.abi.Unpack("getExtensionOperator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetExtensionOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e46e380.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getExtensionOwner(uint256 _extensionId) view returns(address)
func (extensionManager *ExtensionManager) PackGetExtensionOwner(extensionId *big.Int) []byte {
	enc, err := extensionManager.abi.Pack("getExtensionOwner", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetExtensionOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5e46e380.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getExtensionOwner(uint256 _extensionId) view returns(address)
func (extensionManager *ExtensionManager) TryPackGetExtensionOwner(extensionId *big.Int) ([]byte, error) {
	return extensionManager.abi.Pack("getExtensionOwner", extensionId)
}

// UnpackGetExtensionOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5e46e380.
//
// Solidity: function getExtensionOwner(uint256 _extensionId) view returns(address)
func (extensionManager *ExtensionManager) UnpackGetExtensionOwner(data []byte) (common.Address, error) {
	out, err := extensionManager.abi.Unpack("getExtensionOwner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetSupportedCodeHashes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ebb72b2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSupportedCodeHashes(uint256 _extensionId) view returns(bytes32[] _supportedCodeHashes)
func (extensionManager *ExtensionManager) PackGetSupportedCodeHashes(extensionId *big.Int) []byte {
	enc, err := extensionManager.abi.Pack("getSupportedCodeHashes", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSupportedCodeHashes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ebb72b2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSupportedCodeHashes(uint256 _extensionId) view returns(bytes32[] _supportedCodeHashes)
func (extensionManager *ExtensionManager) TryPackGetSupportedCodeHashes(extensionId *big.Int) ([]byte, error) {
	return extensionManager.abi.Pack("getSupportedCodeHashes", extensionId)
}

// UnpackGetSupportedCodeHashes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ebb72b2.
//
// Solidity: function getSupportedCodeHashes(uint256 _extensionId) view returns(bytes32[] _supportedCodeHashes)
func (extensionManager *ExtensionManager) UnpackGetSupportedCodeHashes(data []byte) ([][32]byte, error) {
	out, err := extensionManager.abi.Unpack("getSupportedCodeHashes", data)
	if err != nil {
		return *new([][32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	return out0, nil
}

// PackGetSupportedKeyTypes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8bab3177.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSupportedKeyTypes(uint256 _extensionId) view returns(bytes32[] _supportedKeyTypes)
func (extensionManager *ExtensionManager) PackGetSupportedKeyTypes(extensionId *big.Int) []byte {
	enc, err := extensionManager.abi.Pack("getSupportedKeyTypes", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSupportedKeyTypes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8bab3177.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSupportedKeyTypes(uint256 _extensionId) view returns(bytes32[] _supportedKeyTypes)
func (extensionManager *ExtensionManager) TryPackGetSupportedKeyTypes(extensionId *big.Int) ([]byte, error) {
	return extensionManager.abi.Pack("getSupportedKeyTypes", extensionId)
}

// UnpackGetSupportedKeyTypes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8bab3177.
//
// Solidity: function getSupportedKeyTypes(uint256 _extensionId) view returns(bytes32[] _supportedKeyTypes)
func (extensionManager *ExtensionManager) UnpackGetSupportedKeyTypes(data []byte) ([][32]byte, error) {
	out, err := extensionManager.abi.Unpack("getSupportedKeyTypes", data)
	if err != nil {
		return *new([][32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	return out0, nil
}

// PackGetSystemSupportedKeyTypes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x50a15fcc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSystemSupportedKeyTypes() view returns(bytes32[])
func (extensionManager *ExtensionManager) PackGetSystemSupportedKeyTypes() []byte {
	enc, err := extensionManager.abi.Pack("getSystemSupportedKeyTypes")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSystemSupportedKeyTypes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x50a15fcc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSystemSupportedKeyTypes() view returns(bytes32[])
func (extensionManager *ExtensionManager) TryPackGetSystemSupportedKeyTypes() ([]byte, error) {
	return extensionManager.abi.Pack("getSystemSupportedKeyTypes")
}

// UnpackGetSystemSupportedKeyTypes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x50a15fcc.
//
// Solidity: function getSystemSupportedKeyTypes() view returns(bytes32[])
func (extensionManager *ExtensionManager) UnpackGetSystemSupportedKeyTypes(data []byte) ([][32]byte, error) {
	out, err := extensionManager.abi.Unpack("getSystemSupportedKeyTypes", data)
	if err != nil {
		return *new([][32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	return out0, nil
}

// PackGetSystemSupportedPlatforms is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe38c8964.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSystemSupportedPlatforms() view returns(bytes32[])
func (extensionManager *ExtensionManager) PackGetSystemSupportedPlatforms() []byte {
	enc, err := extensionManager.abi.Pack("getSystemSupportedPlatforms")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSystemSupportedPlatforms is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe38c8964.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSystemSupportedPlatforms() view returns(bytes32[])
func (extensionManager *ExtensionManager) TryPackGetSystemSupportedPlatforms() ([]byte, error) {
	return extensionManager.abi.Pack("getSystemSupportedPlatforms")
}

// UnpackGetSystemSupportedPlatforms is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe38c8964.
//
// Solidity: function getSystemSupportedPlatforms() view returns(bytes32[])
func (extensionManager *ExtensionManager) UnpackGetSystemSupportedPlatforms(data []byte) ([][32]byte, error) {
	out, err := extensionManager.abi.Unpack("getSystemSupportedPlatforms", data)
	if err != nil {
		return *new([][32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	return out0, nil
}

// PackGetSystemSupportedSigningAlgos is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc6cba800.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSystemSupportedSigningAlgos(bytes32 _keyType) view returns(bytes32[])
func (extensionManager *ExtensionManager) PackGetSystemSupportedSigningAlgos(keyType [32]byte) []byte {
	enc, err := extensionManager.abi.Pack("getSystemSupportedSigningAlgos", keyType)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSystemSupportedSigningAlgos is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc6cba800.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSystemSupportedSigningAlgos(bytes32 _keyType) view returns(bytes32[])
func (extensionManager *ExtensionManager) TryPackGetSystemSupportedSigningAlgos(keyType [32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("getSystemSupportedSigningAlgos", keyType)
}

// UnpackGetSystemSupportedSigningAlgos is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc6cba800.
//
// Solidity: function getSystemSupportedSigningAlgos(bytes32 _keyType) view returns(bytes32[])
func (extensionManager *ExtensionManager) UnpackGetSystemSupportedSigningAlgos(data []byte) ([][32]byte, error) {
	out, err := extensionManager.abi.Unpack("getSystemSupportedSigningAlgos", data)
	if err != nil {
		return *new([][32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	return out0, nil
}

// PackGetTeeExtensionInstructionsSender is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c177358.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTeeExtensionInstructionsSender(uint256 _extensionId) view returns(address)
func (extensionManager *ExtensionManager) PackGetTeeExtensionInstructionsSender(extensionId *big.Int) []byte {
	enc, err := extensionManager.abi.Pack("getTeeExtensionInstructionsSender", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTeeExtensionInstructionsSender is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c177358.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTeeExtensionInstructionsSender(uint256 _extensionId) view returns(address)
func (extensionManager *ExtensionManager) TryPackGetTeeExtensionInstructionsSender(extensionId *big.Int) ([]byte, error) {
	return extensionManager.abi.Pack("getTeeExtensionInstructionsSender", extensionId)
}

// UnpackGetTeeExtensionInstructionsSender is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2c177358.
//
// Solidity: function getTeeExtensionInstructionsSender(uint256 _extensionId) view returns(address)
func (extensionManager *ExtensionManager) UnpackGetTeeExtensionInstructionsSender(data []byte) (common.Address, error) {
	out, err := extensionManager.abi.Unpack("getTeeExtensionInstructionsSender", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetTeeExtensionStateVerifier is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x709ebdcd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getTeeExtensionStateVerifier(uint256 _extensionId) view returns(address)
func (extensionManager *ExtensionManager) PackGetTeeExtensionStateVerifier(extensionId *big.Int) []byte {
	enc, err := extensionManager.abi.Pack("getTeeExtensionStateVerifier", extensionId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetTeeExtensionStateVerifier is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x709ebdcd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getTeeExtensionStateVerifier(uint256 _extensionId) view returns(address)
func (extensionManager *ExtensionManager) TryPackGetTeeExtensionStateVerifier(extensionId *big.Int) ([]byte, error) {
	return extensionManager.abi.Pack("getTeeExtensionStateVerifier", extensionId)
}

// UnpackGetTeeExtensionStateVerifier is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x709ebdcd.
//
// Solidity: function getTeeExtensionStateVerifier(uint256 _extensionId) view returns(address)
func (extensionManager *ExtensionManager) UnpackGetTeeExtensionStateVerifier(data []byte) (common.Address, error) {
	out, err := extensionManager.abi.Unpack("getTeeExtensionStateVerifier", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackIsCodeHashPlatformDisabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2a19c98e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isCodeHashPlatformDisabled(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (extensionManager *ExtensionManager) PackIsCodeHashPlatformDisabled(extensionId *big.Int, codeHash [32]byte, platform [32]byte) []byte {
	enc, err := extensionManager.abi.Pack("isCodeHashPlatformDisabled", extensionId, codeHash, platform)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsCodeHashPlatformDisabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2a19c98e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isCodeHashPlatformDisabled(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (extensionManager *ExtensionManager) TryPackIsCodeHashPlatformDisabled(extensionId *big.Int, codeHash [32]byte, platform [32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("isCodeHashPlatformDisabled", extensionId, codeHash, platform)
}

// UnpackIsCodeHashPlatformDisabled is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2a19c98e.
//
// Solidity: function isCodeHashPlatformDisabled(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (extensionManager *ExtensionManager) UnpackIsCodeHashPlatformDisabled(data []byte) (bool, error) {
	out, err := extensionManager.abi.Unpack("isCodeHashPlatformDisabled", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsCodeHashPlatformSupported is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1d925c6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isCodeHashPlatformSupported(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (extensionManager *ExtensionManager) PackIsCodeHashPlatformSupported(extensionId *big.Int, codeHash [32]byte, platform [32]byte) []byte {
	enc, err := extensionManager.abi.Pack("isCodeHashPlatformSupported", extensionId, codeHash, platform)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsCodeHashPlatformSupported is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1d925c6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isCodeHashPlatformSupported(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (extensionManager *ExtensionManager) TryPackIsCodeHashPlatformSupported(extensionId *big.Int, codeHash [32]byte, platform [32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("isCodeHashPlatformSupported", extensionId, codeHash, platform)
}

// UnpackIsCodeHashPlatformSupported is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1d925c6.
//
// Solidity: function isCodeHashPlatformSupported(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (extensionManager *ExtensionManager) UnpackIsCodeHashPlatformSupported(data []byte) (bool, error) {
	out, err := extensionManager.abi.Unpack("isCodeHashPlatformSupported", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsKeyTypeSupported is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf4ac4279.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isKeyTypeSupported(uint256 _extensionId, bytes32 _keyType) view returns(bool)
func (extensionManager *ExtensionManager) PackIsKeyTypeSupported(extensionId *big.Int, keyType [32]byte) []byte {
	enc, err := extensionManager.abi.Pack("isKeyTypeSupported", extensionId, keyType)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsKeyTypeSupported is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf4ac4279.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isKeyTypeSupported(uint256 _extensionId, bytes32 _keyType) view returns(bool)
func (extensionManager *ExtensionManager) TryPackIsKeyTypeSupported(extensionId *big.Int, keyType [32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("isKeyTypeSupported", extensionId, keyType)
}

// UnpackIsKeyTypeSupported is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf4ac4279.
//
// Solidity: function isKeyTypeSupported(uint256 _extensionId, bytes32 _keyType) view returns(bool)
func (extensionManager *ExtensionManager) UnpackIsKeyTypeSupported(data []byte) (bool, error) {
	out, err := extensionManager.abi.Unpack("isKeyTypeSupported", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsSigningAlgoSupported is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62487d98.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isSigningAlgoSupported(bytes32 _keyType, bytes32 _signingAlgo) view returns(bool)
func (extensionManager *ExtensionManager) PackIsSigningAlgoSupported(keyType [32]byte, signingAlgo [32]byte) []byte {
	enc, err := extensionManager.abi.Pack("isSigningAlgoSupported", keyType, signingAlgo)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsSigningAlgoSupported is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62487d98.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isSigningAlgoSupported(bytes32 _keyType, bytes32 _signingAlgo) view returns(bool)
func (extensionManager *ExtensionManager) TryPackIsSigningAlgoSupported(keyType [32]byte, signingAlgo [32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("isSigningAlgoSupported", keyType, signingAlgo)
}

// UnpackIsSigningAlgoSupported is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62487d98.
//
// Solidity: function isSigningAlgoSupported(bytes32 _keyType, bytes32 _signingAlgo) view returns(bool)
func (extensionManager *ExtensionManager) UnpackIsSigningAlgoSupported(data []byte) (bool, error) {
	out, err := extensionManager.abi.Unpack("isSigningAlgoSupported", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackNextPublicExtensionId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x27582ad5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nextPublicExtensionId() view returns(uint256)
func (extensionManager *ExtensionManager) PackNextPublicExtensionId() []byte {
	enc, err := extensionManager.abi.Pack("nextPublicExtensionId")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNextPublicExtensionId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x27582ad5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function nextPublicExtensionId() view returns(uint256)
func (extensionManager *ExtensionManager) TryPackNextPublicExtensionId() ([]byte, error) {
	return extensionManager.abi.Pack("nextPublicExtensionId")
}

// UnpackNextPublicExtensionId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x27582ad5.
//
// Solidity: function nextPublicExtensionId() view returns(uint256)
func (extensionManager *ExtensionManager) UnpackNextPublicExtensionId(data []byte) (*big.Int, error) {
	out, err := extensionManager.abi.Unpack("nextPublicExtensionId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProposeNewOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe947a09f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proposeNewOwner(uint256 _extensionId, address _newOwner) returns()
func (extensionManager *ExtensionManager) PackProposeNewOwner(extensionId *big.Int, newOwner common.Address) []byte {
	enc, err := extensionManager.abi.Pack("proposeNewOwner", extensionId, newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProposeNewOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe947a09f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proposeNewOwner(uint256 _extensionId, address _newOwner) returns()
func (extensionManager *ExtensionManager) TryPackProposeNewOwner(extensionId *big.Int, newOwner common.Address) ([]byte, error) {
	return extensionManager.abi.Pack("proposeNewOwner", extensionId, newOwner)
}

// PackRegister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaa677354.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function register(address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns(uint256 _extensionId)
func (extensionManager *ExtensionManager) PackRegister(teeExtensionStateVerifier common.Address, teeExtensionInstructionsSender common.Address) []byte {
	enc, err := extensionManager.abi.Pack("register", teeExtensionStateVerifier, teeExtensionInstructionsSender)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegister is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaa677354.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function register(address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns(uint256 _extensionId)
func (extensionManager *ExtensionManager) TryPackRegister(teeExtensionStateVerifier common.Address, teeExtensionInstructionsSender common.Address) ([]byte, error) {
	return extensionManager.abi.Pack("register", teeExtensionStateVerifier, teeExtensionInstructionsSender)
}

// UnpackRegister is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaa677354.
//
// Solidity: function register(address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns(uint256 _extensionId)
func (extensionManager *ExtensionManager) UnpackRegister(data []byte) (*big.Int, error) {
	out, err := extensionManager.abi.Unpack("register", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRegisterReserved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf59ac18.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function registerReserved(uint256 _extensionId, address _owner) returns()
func (extensionManager *ExtensionManager) PackRegisterReserved(extensionId *big.Int, owner common.Address) []byte {
	enc, err := extensionManager.abi.Pack("registerReserved", extensionId, owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegisterReserved is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf59ac18.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function registerReserved(uint256 _extensionId, address _owner) returns()
func (extensionManager *ExtensionManager) TryPackRegisterReserved(extensionId *big.Int, owner common.Address) ([]byte, error) {
	return extensionManager.abi.Pack("registerReserved", extensionId, owner)
}

// PackRemoveSupportedKeyTypes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xecb0233f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (extensionManager *ExtensionManager) PackRemoveSupportedKeyTypes(extensionId *big.Int, keyTypes [][32]byte) []byte {
	enc, err := extensionManager.abi.Pack("removeSupportedKeyTypes", extensionId, keyTypes)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveSupportedKeyTypes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xecb0233f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (extensionManager *ExtensionManager) TryPackRemoveSupportedKeyTypes(extensionId *big.Int, keyTypes [][32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("removeSupportedKeyTypes", extensionId, keyTypes)
}

// PackRemoveSystemSupportedKeyTypesAndSigningAlgos is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36d34083.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeSystemSupportedKeyTypesAndSigningAlgos(bytes32[] _keyTypes, bytes32[][] _signingAlgosByKeyType) returns()
func (extensionManager *ExtensionManager) PackRemoveSystemSupportedKeyTypesAndSigningAlgos(keyTypes [][32]byte, signingAlgosByKeyType [][][32]byte) []byte {
	enc, err := extensionManager.abi.Pack("removeSystemSupportedKeyTypesAndSigningAlgos", keyTypes, signingAlgosByKeyType)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveSystemSupportedKeyTypesAndSigningAlgos is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x36d34083.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeSystemSupportedKeyTypesAndSigningAlgos(bytes32[] _keyTypes, bytes32[][] _signingAlgosByKeyType) returns()
func (extensionManager *ExtensionManager) TryPackRemoveSystemSupportedKeyTypesAndSigningAlgos(keyTypes [][32]byte, signingAlgosByKeyType [][][32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("removeSystemSupportedKeyTypesAndSigningAlgos", keyTypes, signingAlgosByKeyType)
}

// PackRemoveSystemSupportedPlatforms is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x97874fed.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeSystemSupportedPlatforms(bytes32[] _platforms) returns()
func (extensionManager *ExtensionManager) PackRemoveSystemSupportedPlatforms(platforms [][32]byte) []byte {
	enc, err := extensionManager.abi.Pack("removeSystemSupportedPlatforms", platforms)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveSystemSupportedPlatforms is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x97874fed.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeSystemSupportedPlatforms(bytes32[] _platforms) returns()
func (extensionManager *ExtensionManager) TryPackRemoveSystemSupportedPlatforms(platforms [][32]byte) ([]byte, error) {
	return extensionManager.abi.Pack("removeSystemSupportedPlatforms", platforms)
}

// PackSetExtensionContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6df0108f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setExtensionContracts(uint256 _extensionId, address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns()
func (extensionManager *ExtensionManager) PackSetExtensionContracts(extensionId *big.Int, teeExtensionStateVerifier common.Address, teeExtensionInstructionsSender common.Address) []byte {
	enc, err := extensionManager.abi.Pack("setExtensionContracts", extensionId, teeExtensionStateVerifier, teeExtensionInstructionsSender)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetExtensionContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6df0108f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setExtensionContracts(uint256 _extensionId, address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns()
func (extensionManager *ExtensionManager) TryPackSetExtensionContracts(extensionId *big.Int, teeExtensionStateVerifier common.Address, teeExtensionInstructionsSender common.Address) ([]byte, error) {
	return extensionManager.abi.Pack("setExtensionContracts", extensionId, teeExtensionStateVerifier, teeExtensionInstructionsSender)
}

// PackSetExtensionOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x868843b9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setExtensionOperator(uint256 _extensionId, address _operator) returns()
func (extensionManager *ExtensionManager) PackSetExtensionOperator(extensionId *big.Int, operator common.Address) []byte {
	enc, err := extensionManager.abi.Pack("setExtensionOperator", extensionId, operator)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetExtensionOperator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x868843b9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setExtensionOperator(uint256 _extensionId, address _operator) returns()
func (extensionManager *ExtensionManager) TryPackSetExtensionOperator(extensionId *big.Int, operator common.Address) ([]byte, error) {
	return extensionManager.abi.Pack("setExtensionOperator", extensionId, operator)
}

// ExtensionManagerCodeHashPlatformsDisabled represents a CodeHashPlatformsDisabled event raised by the ExtensionManager contract.
type ExtensionManagerCodeHashPlatformsDisabled struct {
	ExtensionId *big.Int
	CodeHash    [32]byte
	Platforms   [][32]byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerCodeHashPlatformsDisabledEventName = "CodeHashPlatformsDisabled"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerCodeHashPlatformsDisabled) ContractEventName() string {
	return ExtensionManagerCodeHashPlatformsDisabledEventName
}

// UnpackCodeHashPlatformsDisabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CodeHashPlatformsDisabled(uint256 indexed extensionId, bytes32 indexed codeHash, bytes32[] platforms)
func (extensionManager *ExtensionManager) UnpackCodeHashPlatformsDisabledEvent(log *types.Log) (*ExtensionManagerCodeHashPlatformsDisabled, error) {
	event := "CodeHashPlatformsDisabled"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerCodeHashPlatformsDisabled)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerExtensionOperatorSet represents a ExtensionOperatorSet event raised by the ExtensionManager contract.
type ExtensionManagerExtensionOperatorSet struct {
	ExtensionId *big.Int
	OldOperator common.Address
	NewOperator common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerExtensionOperatorSetEventName = "ExtensionOperatorSet"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerExtensionOperatorSet) ContractEventName() string {
	return ExtensionManagerExtensionOperatorSetEventName
}

// UnpackExtensionOperatorSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ExtensionOperatorSet(uint256 indexed extensionId, address indexed oldOperator, address indexed newOperator)
func (extensionManager *ExtensionManager) UnpackExtensionOperatorSetEvent(log *types.Log) (*ExtensionManagerExtensionOperatorSet, error) {
	event := "ExtensionOperatorSet"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerExtensionOperatorSet)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the ExtensionManager contract.
type ExtensionManagerGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerGovernanceCallTimelocked) ContractEventName() string {
	return ExtensionManagerGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (extensionManager *ExtensionManager) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*ExtensionManagerGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerInitialized represents a Initialized event raised by the ExtensionManager contract.
type ExtensionManagerInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerInitialized) ContractEventName() string {
	return ExtensionManagerInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (extensionManager *ExtensionManager) UnpackInitializedEvent(log *types.Log) (*ExtensionManagerInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerInitialized)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerNewOwnerConfirmed represents a NewOwnerConfirmed event raised by the ExtensionManager contract.
type ExtensionManagerNewOwnerConfirmed struct {
	ExtensionId *big.Int
	NewOwner    common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerNewOwnerConfirmedEventName = "NewOwnerConfirmed"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerNewOwnerConfirmed) ContractEventName() string {
	return ExtensionManagerNewOwnerConfirmedEventName
}

// UnpackNewOwnerConfirmedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewOwnerConfirmed(uint256 indexed extensionId, address indexed newOwner)
func (extensionManager *ExtensionManager) UnpackNewOwnerConfirmedEvent(log *types.Log) (*ExtensionManagerNewOwnerConfirmed, error) {
	event := "NewOwnerConfirmed"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerNewOwnerConfirmed)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerNewOwnerProposed represents a NewOwnerProposed event raised by the ExtensionManager contract.
type ExtensionManagerNewOwnerProposed struct {
	ExtensionId *big.Int
	OldOwner    common.Address
	NewOwner    common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerNewOwnerProposedEventName = "NewOwnerProposed"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerNewOwnerProposed) ContractEventName() string {
	return ExtensionManagerNewOwnerProposedEventName
}

// UnpackNewOwnerProposedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewOwnerProposed(uint256 indexed extensionId, address indexed oldOwner, address indexed newOwner)
func (extensionManager *ExtensionManager) UnpackNewOwnerProposedEvent(log *types.Log) (*ExtensionManagerNewOwnerProposed, error) {
	event := "NewOwnerProposed"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerNewOwnerProposed)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerSupportedKeyTypesAdded represents a SupportedKeyTypesAdded event raised by the ExtensionManager contract.
type ExtensionManagerSupportedKeyTypesAdded struct {
	ExtensionId *big.Int
	KeyTypes    [][32]byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerSupportedKeyTypesAddedEventName = "SupportedKeyTypesAdded"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerSupportedKeyTypesAdded) ContractEventName() string {
	return ExtensionManagerSupportedKeyTypesAddedEventName
}

// UnpackSupportedKeyTypesAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SupportedKeyTypesAdded(uint256 indexed extensionId, bytes32[] keyTypes)
func (extensionManager *ExtensionManager) UnpackSupportedKeyTypesAddedEvent(log *types.Log) (*ExtensionManagerSupportedKeyTypesAdded, error) {
	event := "SupportedKeyTypesAdded"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerSupportedKeyTypesAdded)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerSupportedKeyTypesRemoved represents a SupportedKeyTypesRemoved event raised by the ExtensionManager contract.
type ExtensionManagerSupportedKeyTypesRemoved struct {
	ExtensionId *big.Int
	KeyTypes    [][32]byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerSupportedKeyTypesRemovedEventName = "SupportedKeyTypesRemoved"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerSupportedKeyTypesRemoved) ContractEventName() string {
	return ExtensionManagerSupportedKeyTypesRemovedEventName
}

// UnpackSupportedKeyTypesRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SupportedKeyTypesRemoved(uint256 indexed extensionId, bytes32[] keyTypes)
func (extensionManager *ExtensionManager) UnpackSupportedKeyTypesRemovedEvent(log *types.Log) (*ExtensionManagerSupportedKeyTypesRemoved, error) {
	event := "SupportedKeyTypesRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerSupportedKeyTypesRemoved)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded represents a SystemSupportedKeyTypesAndSigningAlgosAdded event raised by the ExtensionManager contract.
type ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded struct {
	KeyTypes              [][32]byte
	SigningAlgosByKeyType [][][32]byte
	Raw                   *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAddedEventName = "SystemSupportedKeyTypesAndSigningAlgosAdded"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded) ContractEventName() string {
	return ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAddedEventName
}

// UnpackSystemSupportedKeyTypesAndSigningAlgosAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SystemSupportedKeyTypesAndSigningAlgosAdded(bytes32[] keyTypes, bytes32[][] signingAlgosByKeyType)
func (extensionManager *ExtensionManager) UnpackSystemSupportedKeyTypesAndSigningAlgosAddedEvent(log *types.Log) (*ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded, error) {
	event := "SystemSupportedKeyTypesAndSigningAlgosAdded"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosRemoved represents a SystemSupportedKeyTypesAndSigningAlgosRemoved event raised by the ExtensionManager contract.
type ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosRemoved struct {
	KeyTypes              [][32]byte
	SigningAlgosByKeyType [][][32]byte
	Raw                   *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosRemovedEventName = "SystemSupportedKeyTypesAndSigningAlgosRemoved"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosRemoved) ContractEventName() string {
	return ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosRemovedEventName
}

// UnpackSystemSupportedKeyTypesAndSigningAlgosRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SystemSupportedKeyTypesAndSigningAlgosRemoved(bytes32[] keyTypes, bytes32[][] signingAlgosByKeyType)
func (extensionManager *ExtensionManager) UnpackSystemSupportedKeyTypesAndSigningAlgosRemovedEvent(log *types.Log) (*ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosRemoved, error) {
	event := "SystemSupportedKeyTypesAndSigningAlgosRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosRemoved)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerSystemSupportedPlatformsAdded represents a SystemSupportedPlatformsAdded event raised by the ExtensionManager contract.
type ExtensionManagerSystemSupportedPlatformsAdded struct {
	Platforms [][32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerSystemSupportedPlatformsAddedEventName = "SystemSupportedPlatformsAdded"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerSystemSupportedPlatformsAdded) ContractEventName() string {
	return ExtensionManagerSystemSupportedPlatformsAddedEventName
}

// UnpackSystemSupportedPlatformsAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SystemSupportedPlatformsAdded(bytes32[] platforms)
func (extensionManager *ExtensionManager) UnpackSystemSupportedPlatformsAddedEvent(log *types.Log) (*ExtensionManagerSystemSupportedPlatformsAdded, error) {
	event := "SystemSupportedPlatformsAdded"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerSystemSupportedPlatformsAdded)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerSystemSupportedPlatformsRemoved represents a SystemSupportedPlatformsRemoved event raised by the ExtensionManager contract.
type ExtensionManagerSystemSupportedPlatformsRemoved struct {
	Platforms [][32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerSystemSupportedPlatformsRemovedEventName = "SystemSupportedPlatformsRemoved"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerSystemSupportedPlatformsRemoved) ContractEventName() string {
	return ExtensionManagerSystemSupportedPlatformsRemovedEventName
}

// UnpackSystemSupportedPlatformsRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SystemSupportedPlatformsRemoved(bytes32[] platforms)
func (extensionManager *ExtensionManager) UnpackSystemSupportedPlatformsRemovedEvent(log *types.Log) (*ExtensionManagerSystemSupportedPlatformsRemoved, error) {
	event := "SystemSupportedPlatformsRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerSystemSupportedPlatformsRemoved)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerTeeExtensionContractsSet represents a TeeExtensionContractsSet event raised by the ExtensionManager contract.
type ExtensionManagerTeeExtensionContractsSet struct {
	ExtensionId                    *big.Int
	TeeExtensionStateVerifier      common.Address
	TeeExtensionInstructionsSender common.Address
	Raw                            *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerTeeExtensionContractsSetEventName = "TeeExtensionContractsSet"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerTeeExtensionContractsSet) ContractEventName() string {
	return ExtensionManagerTeeExtensionContractsSetEventName
}

// UnpackTeeExtensionContractsSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeExtensionContractsSet(uint256 indexed extensionId, address indexed teeExtensionStateVerifier, address indexed teeExtensionInstructionsSender)
func (extensionManager *ExtensionManager) UnpackTeeExtensionContractsSetEvent(log *types.Log) (*ExtensionManagerTeeExtensionContractsSet, error) {
	event := "TeeExtensionContractsSet"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerTeeExtensionContractsSet)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerTeeExtensionRegistered represents a TeeExtensionRegistered event raised by the ExtensionManager contract.
type ExtensionManagerTeeExtensionRegistered struct {
	ExtensionId *big.Int
	Owner       common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerTeeExtensionRegisteredEventName = "TeeExtensionRegistered"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerTeeExtensionRegistered) ContractEventName() string {
	return ExtensionManagerTeeExtensionRegisteredEventName
}

// UnpackTeeExtensionRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeExtensionRegistered(uint256 indexed extensionId, address indexed owner)
func (extensionManager *ExtensionManager) UnpackTeeExtensionRegisteredEvent(log *types.Log) (*ExtensionManagerTeeExtensionRegistered, error) {
	event := "TeeExtensionRegistered"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerTeeExtensionRegistered)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerTeeMachineStatusChanged represents a TeeMachineStatusChanged event raised by the ExtensionManager contract.
type ExtensionManagerTeeMachineStatusChanged struct {
	TeeId     common.Address
	NewStatus uint8
	Raw       *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerTeeMachineStatusChangedEventName = "TeeMachineStatusChanged"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerTeeMachineStatusChanged) ContractEventName() string {
	return ExtensionManagerTeeMachineStatusChangedEventName
}

// UnpackTeeMachineStatusChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeMachineStatusChanged(address indexed teeId, uint8 indexed newStatus)
func (extensionManager *ExtensionManager) UnpackTeeMachineStatusChangedEvent(log *types.Log) (*ExtensionManagerTeeMachineStatusChanged, error) {
	event := "TeeMachineStatusChanged"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerTeeMachineStatusChanged)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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

// ExtensionManagerTeeVersionAdded represents a TeeVersionAdded event raised by the ExtensionManager contract.
type ExtensionManagerTeeVersionAdded struct {
	ExtensionId *big.Int
	Version     [32]byte
	CodeHash    [32]byte
	Platforms   [][32]byte
	Raw         *types.Log // Blockchain specific contextual infos
}

const ExtensionManagerTeeVersionAddedEventName = "TeeVersionAdded"

// ContractEventName returns the user-defined event name.
func (ExtensionManagerTeeVersionAdded) ContractEventName() string {
	return ExtensionManagerTeeVersionAddedEventName
}

// UnpackTeeVersionAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeVersionAdded(uint256 indexed extensionId, bytes32 version, bytes32 indexed codeHash, bytes32[] platforms)
func (extensionManager *ExtensionManager) UnpackTeeVersionAddedEvent(log *types.Log) (*ExtensionManagerTeeVersionAdded, error) {
	event := "TeeVersionAdded"
	if len(log.Topics) == 0 || log.Topics[0] != extensionManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(ExtensionManagerTeeVersionAdded)
	if len(log.Data) > 0 {
		if err := extensionManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range extensionManager.abi.Events[event].Inputs {
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
func (extensionManager *ExtensionManager) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["AddressAlreadyInSet"].ID.Bytes()[:4]) {
		return extensionManager.UnpackAddressAlreadyInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["AddressNotInSet"].ID.Bytes()[:4]) {
		return extensionManager.UnpackAddressNotInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["AvailabilityCheckTimestampInvalid"].ID.Bytes()[:4]) {
		return extensionManager.UnpackAvailabilityCheckTimestampInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["CodeHashPlatformAlreadyDisabled"].ID.Bytes()[:4]) {
		return extensionManager.UnpackCodeHashPlatformAlreadyDisabledError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["CodeHashZero"].ID.Bytes()[:4]) {
		return extensionManager.UnpackCodeHashZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["DuplicatedCosigner"].ID.Bytes()[:4]) {
		return extensionManager.UnpackDuplicatedCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return extensionManager.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidCodeHash"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidCodeHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidExtensionOwner"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidInstructionsSender"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidInstructionsSenderError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidKeyType"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidPlatform"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidPlatformError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidPublicKey"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidReservedExtensionId"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidReservedExtensionIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidResponseData"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidResponseDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidSigningAlgo"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidSigningAlgoError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["InvalidWalletStatus"].ID.Bytes()[:4]) {
		return extensionManager.UnpackInvalidWalletStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["KeyTypeAlreadyExists"].ID.Bytes()[:4]) {
		return extensionManager.UnpackKeyTypeAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["KeyTypeEmpty"].ID.Bytes()[:4]) {
		return extensionManager.UnpackKeyTypeEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["KeyTypeNotSupported"].ID.Bytes()[:4]) {
		return extensionManager.UnpackKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return extensionManager.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["NoAddresses"].ID.Bytes()[:4]) {
		return extensionManager.UnpackNoAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["NoKeyTypes"].ID.Bytes()[:4]) {
		return extensionManager.UnpackNoKeyTypesError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["NoPlatforms"].ID.Bytes()[:4]) {
		return extensionManager.UnpackNoPlatformsError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["NoSigningAlgos"].ID.Bytes()[:4]) {
		return extensionManager.UnpackNoSigningAlgosError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["NotAllowedExtensionOwner"].ID.Bytes()[:4]) {
		return extensionManager.UnpackNotAllowedExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return extensionManager.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["NotOwnerOrPauser"].ID.Bytes()[:4]) {
		return extensionManager.UnpackNotOwnerOrPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["NotOwnerOrUnpauser"].ID.Bytes()[:4]) {
		return extensionManager.UnpackNotOwnerOrUnpauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["OnlyExtensionOwner"].ID.Bytes()[:4]) {
		return extensionManager.UnpackOnlyExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return extensionManager.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return extensionManager.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return extensionManager.UnpackOnlyOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["OnlyOwnerOrBackupManager"].ID.Bytes()[:4]) {
		return extensionManager.UnpackOnlyOwnerOrBackupManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return extensionManager.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["OnlyProposedOwner"].ID.Bytes()[:4]) {
		return extensionManager.UnpackOnlyProposedOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["OwnerNotAllowed"].ID.Bytes()[:4]) {
		return extensionManager.UnpackOwnerNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["PlatformAlreadyExists"].ID.Bytes()[:4]) {
		return extensionManager.UnpackPlatformAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["PlatformEmpty"].ID.Bytes()[:4]) {
		return extensionManager.UnpackPlatformEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["PlatformNotFound"].ID.Bytes()[:4]) {
		return extensionManager.UnpackPlatformNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["ReservedExtensionIdAlreadyAssigned"].ID.Bytes()[:4]) {
		return extensionManager.UnpackReservedExtensionIdAlreadyAssignedError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["SigningAlgoAlreadyExists"].ID.Bytes()[:4]) {
		return extensionManager.UnpackSigningAlgoAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["SigningAlgoEmpty"].ID.Bytes()[:4]) {
		return extensionManager.UnpackSigningAlgoEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["SigningAlgoNotFound"].ID.Bytes()[:4]) {
		return extensionManager.UnpackSigningAlgoNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["SystemOwnedExtensionId"].ID.Bytes()[:4]) {
		return extensionManager.UnpackSystemOwnedExtensionIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return extensionManager.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["UnsupportedPlatform"].ID.Bytes()[:4]) {
		return extensionManager.UnpackUnsupportedPlatformError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["VersionAlreadyExists"].ID.Bytes()[:4]) {
		return extensionManager.UnpackVersionAlreadyExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["VersionEmpty"].ID.Bytes()[:4]) {
		return extensionManager.UnpackVersionEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], extensionManager.abi.Errors["VersionNotSupported"].ID.Bytes()[:4]) {
		return extensionManager.UnpackVersionNotSupportedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// ExtensionManagerAddressAlreadyInSet represents a AddressAlreadyInSet error raised by the ExtensionManager contract.
type ExtensionManagerAddressAlreadyInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressAlreadyInSet(address addr)
func ExtensionManagerAddressAlreadyInSetErrorID() common.Hash {
	return common.HexToHash("0xb725c7fb1c3f0f065973d6b9eb55140f970259955af3045742cd5b37902eb920")
}

// UnpackAddressAlreadyInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressAlreadyInSet(address addr)
func (extensionManager *ExtensionManager) UnpackAddressAlreadyInSetError(raw []byte) (*ExtensionManagerAddressAlreadyInSet, error) {
	out := new(ExtensionManagerAddressAlreadyInSet)
	if err := extensionManager.abi.UnpackIntoInterface(out, "AddressAlreadyInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerAddressNotInSet represents a AddressNotInSet error raised by the ExtensionManager contract.
type ExtensionManagerAddressNotInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressNotInSet(address addr)
func ExtensionManagerAddressNotInSetErrorID() common.Hash {
	return common.HexToHash("0x1ff5ecf3e6d385ea01ff0323fe894ed6d25a83ed5f6ff2a2ede4c897408f6be9")
}

// UnpackAddressNotInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressNotInSet(address addr)
func (extensionManager *ExtensionManager) UnpackAddressNotInSetError(raw []byte) (*ExtensionManagerAddressNotInSet, error) {
	out := new(ExtensionManagerAddressNotInSet)
	if err := extensionManager.abi.UnpackIntoInterface(out, "AddressNotInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerAvailabilityCheckTimestampInvalid represents a AvailabilityCheckTimestampInvalid error raised by the ExtensionManager contract.
type ExtensionManagerAvailabilityCheckTimestampInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func ExtensionManagerAvailabilityCheckTimestampInvalidErrorID() common.Hash {
	return common.HexToHash("0x6344a5e74cff78612fa271fa907ca608af3bad10e1e8c8df302e29ebc6298e2d")
}

// UnpackAvailabilityCheckTimestampInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func (extensionManager *ExtensionManager) UnpackAvailabilityCheckTimestampInvalidError(raw []byte) (*ExtensionManagerAvailabilityCheckTimestampInvalid, error) {
	out := new(ExtensionManagerAvailabilityCheckTimestampInvalid)
	if err := extensionManager.abi.UnpackIntoInterface(out, "AvailabilityCheckTimestampInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerCodeHashPlatformAlreadyDisabled represents a CodeHashPlatformAlreadyDisabled error raised by the ExtensionManager contract.
type ExtensionManagerCodeHashPlatformAlreadyDisabled struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CodeHashPlatformAlreadyDisabled()
func ExtensionManagerCodeHashPlatformAlreadyDisabledErrorID() common.Hash {
	return common.HexToHash("0xf4595ad5dedf107f0316503ba8bcc6ce761b2b4f2684a220aa655de3695f59ea")
}

// UnpackCodeHashPlatformAlreadyDisabledError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CodeHashPlatformAlreadyDisabled()
func (extensionManager *ExtensionManager) UnpackCodeHashPlatformAlreadyDisabledError(raw []byte) (*ExtensionManagerCodeHashPlatformAlreadyDisabled, error) {
	out := new(ExtensionManagerCodeHashPlatformAlreadyDisabled)
	if err := extensionManager.abi.UnpackIntoInterface(out, "CodeHashPlatformAlreadyDisabled", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerCodeHashZero represents a CodeHashZero error raised by the ExtensionManager contract.
type ExtensionManagerCodeHashZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CodeHashZero()
func ExtensionManagerCodeHashZeroErrorID() common.Hash {
	return common.HexToHash("0x70dd96566d5b147aed2efbe406553b3c35353d50b94a4200336b8cd955a9bb68")
}

// UnpackCodeHashZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CodeHashZero()
func (extensionManager *ExtensionManager) UnpackCodeHashZeroError(raw []byte) (*ExtensionManagerCodeHashZero, error) {
	out := new(ExtensionManagerCodeHashZero)
	if err := extensionManager.abi.UnpackIntoInterface(out, "CodeHashZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerDuplicatedCosigner represents a DuplicatedCosigner error raised by the ExtensionManager contract.
type ExtensionManagerDuplicatedCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func ExtensionManagerDuplicatedCosignerErrorID() common.Hash {
	return common.HexToHash("0x53a5890417576fb3043dc3ff1669b0932da2f657d8adb4c11313b12b711f5910")
}

// UnpackDuplicatedCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func (extensionManager *ExtensionManager) UnpackDuplicatedCosignerError(raw []byte) (*ExtensionManagerDuplicatedCosigner, error) {
	out := new(ExtensionManagerDuplicatedCosigner)
	if err := extensionManager.abi.UnpackIntoInterface(out, "DuplicatedCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerExtensionIdMismatch represents a ExtensionIdMismatch error raised by the ExtensionManager contract.
type ExtensionManagerExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func ExtensionManagerExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (extensionManager *ExtensionManager) UnpackExtensionIdMismatchError(raw []byte) (*ExtensionManagerExtensionIdMismatch, error) {
	out := new(ExtensionManagerExtensionIdMismatch)
	if err := extensionManager.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidAddress represents a InvalidAddress error raised by the ExtensionManager contract.
type ExtensionManagerInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func ExtensionManagerInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (extensionManager *ExtensionManager) UnpackInvalidAddressError(raw []byte) (*ExtensionManagerInvalidAddress, error) {
	out := new(ExtensionManagerInvalidAddress)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidAvailabilityCheckStatus represents a InvalidAvailabilityCheckStatus error raised by the ExtensionManager contract.
type ExtensionManagerInvalidAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func ExtensionManagerInvalidAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0xedee4ff00b842d3ef84031c60001065c540c838c5479bbf0ae56f15868427e44")
}

// UnpackInvalidAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func (extensionManager *ExtensionManager) UnpackInvalidAvailabilityCheckStatusError(raw []byte) (*ExtensionManagerInvalidAvailabilityCheckStatus, error) {
	out := new(ExtensionManagerInvalidAvailabilityCheckStatus)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidCodeHash represents a InvalidCodeHash error raised by the ExtensionManager contract.
type ExtensionManagerInvalidCodeHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCodeHash()
func ExtensionManagerInvalidCodeHashErrorID() common.Hash {
	return common.HexToHash("0x8f84fb2496c57bdbdd3c1f135efc28a30a089cd0a7109349a9ca81714a44b7bd")
}

// UnpackInvalidCodeHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCodeHash()
func (extensionManager *ExtensionManager) UnpackInvalidCodeHashError(raw []byte) (*ExtensionManagerInvalidCodeHash, error) {
	out := new(ExtensionManagerInvalidCodeHash)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidCodeHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidCosigner represents a InvalidCosigner error raised by the ExtensionManager contract.
type ExtensionManagerInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func ExtensionManagerInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (extensionManager *ExtensionManager) UnpackInvalidCosignerError(raw []byte) (*ExtensionManagerInvalidCosigner, error) {
	out := new(ExtensionManagerInvalidCosigner)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidDuration represents a InvalidDuration error raised by the ExtensionManager contract.
type ExtensionManagerInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func ExtensionManagerInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (extensionManager *ExtensionManager) UnpackInvalidDurationError(raw []byte) (*ExtensionManagerInvalidDuration, error) {
	out := new(ExtensionManagerInvalidDuration)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidExtensionOwner represents a InvalidExtensionOwner error raised by the ExtensionManager contract.
type ExtensionManagerInvalidExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidExtensionOwner()
func ExtensionManagerInvalidExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xe4f7bb4686f3a681127cca1fcd379b73e07f76f34a2a7db21dff316047c70d7f")
}

// UnpackInvalidExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidExtensionOwner()
func (extensionManager *ExtensionManager) UnpackInvalidExtensionOwnerError(raw []byte) (*ExtensionManagerInvalidExtensionOwner, error) {
	out := new(ExtensionManagerInvalidExtensionOwner)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the ExtensionManager contract.
type ExtensionManagerInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func ExtensionManagerInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (extensionManager *ExtensionManager) UnpackInvalidGovernanceHashError(raw []byte) (*ExtensionManagerInvalidGovernanceHash, error) {
	out := new(ExtensionManagerInvalidGovernanceHash)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidInitialization represents a InvalidInitialization error raised by the ExtensionManager contract.
type ExtensionManagerInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func ExtensionManagerInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (extensionManager *ExtensionManager) UnpackInvalidInitializationError(raw []byte) (*ExtensionManagerInvalidInitialization, error) {
	out := new(ExtensionManagerInvalidInitialization)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidInstructionsSender represents a InvalidInstructionsSender error raised by the ExtensionManager contract.
type ExtensionManagerInvalidInstructionsSender struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInstructionsSender()
func ExtensionManagerInvalidInstructionsSenderErrorID() common.Hash {
	return common.HexToHash("0x5bed0dd65877c8b84b4527862f2be859e8495bb33a1d40cf0df0156d3973c72e")
}

// UnpackInvalidInstructionsSenderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInstructionsSender()
func (extensionManager *ExtensionManager) UnpackInvalidInstructionsSenderError(raw []byte) (*ExtensionManagerInvalidInstructionsSender, error) {
	out := new(ExtensionManagerInvalidInstructionsSender)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidInstructionsSender", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidKeyType represents a InvalidKeyType error raised by the ExtensionManager contract.
type ExtensionManagerInvalidKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyType()
func ExtensionManagerInvalidKeyTypeErrorID() common.Hash {
	return common.HexToHash("0x36dccb440beb343622d8ee0927bce3f69281feea3a5b0610f6e00e2811b557e1")
}

// UnpackInvalidKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyType()
func (extensionManager *ExtensionManager) UnpackInvalidKeyTypeError(raw []byte) (*ExtensionManagerInvalidKeyType, error) {
	out := new(ExtensionManagerInvalidKeyType)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidNonce represents a InvalidNonce error raised by the ExtensionManager contract.
type ExtensionManagerInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func ExtensionManagerInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (extensionManager *ExtensionManager) UnpackInvalidNonceError(raw []byte) (*ExtensionManagerInvalidNonce, error) {
	out := new(ExtensionManagerInvalidNonce)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidPlatform represents a InvalidPlatform error raised by the ExtensionManager contract.
type ExtensionManagerInvalidPlatform struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPlatform()
func ExtensionManagerInvalidPlatformErrorID() common.Hash {
	return common.HexToHash("0x62f0746f1a5c097931c0a251929011e54b21baef7b0120f5aad2bb81b99a30e4")
}

// UnpackInvalidPlatformError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPlatform()
func (extensionManager *ExtensionManager) UnpackInvalidPlatformError(raw []byte) (*ExtensionManagerInvalidPlatform, error) {
	out := new(ExtensionManagerInvalidPlatform)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidPlatform", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidPublicKey represents a InvalidPublicKey error raised by the ExtensionManager contract.
type ExtensionManagerInvalidPublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPublicKey()
func ExtensionManagerInvalidPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa2d0fee854985cb3d3a394146b5df1c41d5b4d6fcede7fc57a3c05a49b2f3e73")
}

// UnpackInvalidPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPublicKey()
func (extensionManager *ExtensionManager) UnpackInvalidPublicKeyError(raw []byte) (*ExtensionManagerInvalidPublicKey, error) {
	out := new(ExtensionManagerInvalidPublicKey)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidReservedExtensionId represents a InvalidReservedExtensionId error raised by the ExtensionManager contract.
type ExtensionManagerInvalidReservedExtensionId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidReservedExtensionId()
func ExtensionManagerInvalidReservedExtensionIdErrorID() common.Hash {
	return common.HexToHash("0xa635e195e1d9e53aceefd9a4db03241d570f0faaa781e2f595c75928a1e44a80")
}

// UnpackInvalidReservedExtensionIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidReservedExtensionId()
func (extensionManager *ExtensionManager) UnpackInvalidReservedExtensionIdError(raw []byte) (*ExtensionManagerInvalidReservedExtensionId, error) {
	out := new(ExtensionManagerInvalidReservedExtensionId)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidReservedExtensionId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidResponseData represents a InvalidResponseData error raised by the ExtensionManager contract.
type ExtensionManagerInvalidResponseData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseData()
func ExtensionManagerInvalidResponseDataErrorID() common.Hash {
	return common.HexToHash("0x48566617128355a033880805d6462da9bf9c6fe4c61b922997451a71262990af")
}

// UnpackInvalidResponseDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseData()
func (extensionManager *ExtensionManager) UnpackInvalidResponseDataError(raw []byte) (*ExtensionManagerInvalidResponseData, error) {
	out := new(ExtensionManagerInvalidResponseData)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidResponseData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidSigningAlgo represents a InvalidSigningAlgo error raised by the ExtensionManager contract.
type ExtensionManagerInvalidSigningAlgo struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningAlgo()
func ExtensionManagerInvalidSigningAlgoErrorID() common.Hash {
	return common.HexToHash("0x9b7198a697b61e546ef9ee593a6f77d2bf9f1df9dcbba655ce96d7e2c75346c7")
}

// UnpackInvalidSigningAlgoError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningAlgo()
func (extensionManager *ExtensionManager) UnpackInvalidSigningAlgoError(raw []byte) (*ExtensionManagerInvalidSigningAlgo, error) {
	out := new(ExtensionManagerInvalidSigningAlgo)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidSigningAlgo", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidThreshold represents a InvalidThreshold error raised by the ExtensionManager contract.
type ExtensionManagerInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func ExtensionManagerInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (extensionManager *ExtensionManager) UnpackInvalidThresholdError(raw []byte) (*ExtensionManagerInvalidThreshold, error) {
	out := new(ExtensionManagerInvalidThreshold)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerInvalidWalletStatus represents a InvalidWalletStatus error raised by the ExtensionManager contract.
type ExtensionManagerInvalidWalletStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidWalletStatus()
func ExtensionManagerInvalidWalletStatusErrorID() common.Hash {
	return common.HexToHash("0xd700120e4896ea7835a5e9e4feb690644ced5cad1c52209afdcbe06b96d91345")
}

// UnpackInvalidWalletStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidWalletStatus()
func (extensionManager *ExtensionManager) UnpackInvalidWalletStatusError(raw []byte) (*ExtensionManagerInvalidWalletStatus, error) {
	out := new(ExtensionManagerInvalidWalletStatus)
	if err := extensionManager.abi.UnpackIntoInterface(out, "InvalidWalletStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerKeyTypeAlreadyExists represents a KeyTypeAlreadyExists error raised by the ExtensionManager contract.
type ExtensionManagerKeyTypeAlreadyExists struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeAlreadyExists(bytes32 keyType)
func ExtensionManagerKeyTypeAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x25b06a192170c4ad643d794e7d7db577956500627d59f81aae91e79f0ecc2649")
}

// UnpackKeyTypeAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeAlreadyExists(bytes32 keyType)
func (extensionManager *ExtensionManager) UnpackKeyTypeAlreadyExistsError(raw []byte) (*ExtensionManagerKeyTypeAlreadyExists, error) {
	out := new(ExtensionManagerKeyTypeAlreadyExists)
	if err := extensionManager.abi.UnpackIntoInterface(out, "KeyTypeAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerKeyTypeEmpty represents a KeyTypeEmpty error raised by the ExtensionManager contract.
type ExtensionManagerKeyTypeEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeEmpty()
func ExtensionManagerKeyTypeEmptyErrorID() common.Hash {
	return common.HexToHash("0xe2a144a5f53c25c20b4c58d15a771d0cd9f4bf209da0c848df15657c51d02dca")
}

// UnpackKeyTypeEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeEmpty()
func (extensionManager *ExtensionManager) UnpackKeyTypeEmptyError(raw []byte) (*ExtensionManagerKeyTypeEmpty, error) {
	out := new(ExtensionManagerKeyTypeEmpty)
	if err := extensionManager.abi.UnpackIntoInterface(out, "KeyTypeEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerKeyTypeNotSupported represents a KeyTypeNotSupported error raised by the ExtensionManager contract.
type ExtensionManagerKeyTypeNotSupported struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func ExtensionManagerKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x85004df1f47814c2de6fbf03d3f2b1ca8c67c182d0d54278766543b6be3bc603")
}

// UnpackKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func (extensionManager *ExtensionManager) UnpackKeyTypeNotSupportedError(raw []byte) (*ExtensionManagerKeyTypeNotSupported, error) {
	out := new(ExtensionManagerKeyTypeNotSupported)
	if err := extensionManager.abi.UnpackIntoInterface(out, "KeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerLengthsMismatch represents a LengthsMismatch error raised by the ExtensionManager contract.
type ExtensionManagerLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func ExtensionManagerLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (extensionManager *ExtensionManager) UnpackLengthsMismatchError(raw []byte) (*ExtensionManagerLengthsMismatch, error) {
	out := new(ExtensionManagerLengthsMismatch)
	if err := extensionManager.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerNoAddresses represents a NoAddresses error raised by the ExtensionManager contract.
type ExtensionManagerNoAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoAddresses()
func ExtensionManagerNoAddressesErrorID() common.Hash {
	return common.HexToHash("0xb2e79f05746fa6c9f3d61f9efe85959371cde5030a6c2dd70c92b63578c5cde7")
}

// UnpackNoAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoAddresses()
func (extensionManager *ExtensionManager) UnpackNoAddressesError(raw []byte) (*ExtensionManagerNoAddresses, error) {
	out := new(ExtensionManagerNoAddresses)
	if err := extensionManager.abi.UnpackIntoInterface(out, "NoAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerNoKeyTypes represents a NoKeyTypes error raised by the ExtensionManager contract.
type ExtensionManagerNoKeyTypes struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoKeyTypes()
func ExtensionManagerNoKeyTypesErrorID() common.Hash {
	return common.HexToHash("0xc448b33ddb2fd84ff67cc0066f161b1d204ed0b1481db7915df3a0b7e20da2b7")
}

// UnpackNoKeyTypesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoKeyTypes()
func (extensionManager *ExtensionManager) UnpackNoKeyTypesError(raw []byte) (*ExtensionManagerNoKeyTypes, error) {
	out := new(ExtensionManagerNoKeyTypes)
	if err := extensionManager.abi.UnpackIntoInterface(out, "NoKeyTypes", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerNoPlatforms represents a NoPlatforms error raised by the ExtensionManager contract.
type ExtensionManagerNoPlatforms struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoPlatforms()
func ExtensionManagerNoPlatformsErrorID() common.Hash {
	return common.HexToHash("0x27c348052935c75ea61c74906c098d14cab296073930875fe474c5af1d35a8c1")
}

// UnpackNoPlatformsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoPlatforms()
func (extensionManager *ExtensionManager) UnpackNoPlatformsError(raw []byte) (*ExtensionManagerNoPlatforms, error) {
	out := new(ExtensionManagerNoPlatforms)
	if err := extensionManager.abi.UnpackIntoInterface(out, "NoPlatforms", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerNoSigningAlgos represents a NoSigningAlgos error raised by the ExtensionManager contract.
type ExtensionManagerNoSigningAlgos struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoSigningAlgos(bytes32 keyType)
func ExtensionManagerNoSigningAlgosErrorID() common.Hash {
	return common.HexToHash("0xbe1a21733eff41aa17caf006c6e92ce348ebedac6f173851e3d6fc4159b94cc7")
}

// UnpackNoSigningAlgosError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoSigningAlgos(bytes32 keyType)
func (extensionManager *ExtensionManager) UnpackNoSigningAlgosError(raw []byte) (*ExtensionManagerNoSigningAlgos, error) {
	out := new(ExtensionManagerNoSigningAlgos)
	if err := extensionManager.abi.UnpackIntoInterface(out, "NoSigningAlgos", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerNotAllowedExtensionOwner represents a NotAllowedExtensionOwner error raised by the ExtensionManager contract.
type ExtensionManagerNotAllowedExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotAllowedExtensionOwner()
func ExtensionManagerNotAllowedExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xee3f9812d033ce9a92b488b3eaf18fb279498d558a69e003bd9266c74a3bf6ab")
}

// UnpackNotAllowedExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotAllowedExtensionOwner()
func (extensionManager *ExtensionManager) UnpackNotAllowedExtensionOwnerError(raw []byte) (*ExtensionManagerNotAllowedExtensionOwner, error) {
	out := new(ExtensionManagerNotAllowedExtensionOwner)
	if err := extensionManager.abi.UnpackIntoInterface(out, "NotAllowedExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerNotInitializing represents a NotInitializing error raised by the ExtensionManager contract.
type ExtensionManagerNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func ExtensionManagerNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (extensionManager *ExtensionManager) UnpackNotInitializingError(raw []byte) (*ExtensionManagerNotInitializing, error) {
	out := new(ExtensionManagerNotInitializing)
	if err := extensionManager.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerNotOwnerOrPauser represents a NotOwnerOrPauser error raised by the ExtensionManager contract.
type ExtensionManagerNotOwnerOrPauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrPauser(address caller)
func ExtensionManagerNotOwnerOrPauserErrorID() common.Hash {
	return common.HexToHash("0x46c0d9af6b895093e022ad2984c54e60192e6e26c5063daf232e825c6e07ff9a")
}

// UnpackNotOwnerOrPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrPauser(address caller)
func (extensionManager *ExtensionManager) UnpackNotOwnerOrPauserError(raw []byte) (*ExtensionManagerNotOwnerOrPauser, error) {
	out := new(ExtensionManagerNotOwnerOrPauser)
	if err := extensionManager.abi.UnpackIntoInterface(out, "NotOwnerOrPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerNotOwnerOrUnpauser represents a NotOwnerOrUnpauser error raised by the ExtensionManager contract.
type ExtensionManagerNotOwnerOrUnpauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func ExtensionManagerNotOwnerOrUnpauserErrorID() common.Hash {
	return common.HexToHash("0x81e700e8f0bd5ba4a6d6a2ce9a1ad0ac9bb1b1b1e62d921a55cbc4b53f1da649")
}

// UnpackNotOwnerOrUnpauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func (extensionManager *ExtensionManager) UnpackNotOwnerOrUnpauserError(raw []byte) (*ExtensionManagerNotOwnerOrUnpauser, error) {
	out := new(ExtensionManagerNotOwnerOrUnpauser)
	if err := extensionManager.abi.UnpackIntoInterface(out, "NotOwnerOrUnpauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerOnlyExtensionOwner represents a OnlyExtensionOwner error raised by the ExtensionManager contract.
type ExtensionManagerOnlyExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwner()
func ExtensionManagerOnlyExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xcd49fd1de5a297cd1ce055d5abffcf2cb25ad9492bd39a703a238af57b2a9488")
}

// UnpackOnlyExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwner()
func (extensionManager *ExtensionManager) UnpackOnlyExtensionOwnerError(raw []byte) (*ExtensionManagerOnlyExtensionOwner, error) {
	out := new(ExtensionManagerOnlyExtensionOwner)
	if err := extensionManager.abi.UnpackIntoInterface(out, "OnlyExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the ExtensionManager contract.
type ExtensionManagerOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func ExtensionManagerOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (extensionManager *ExtensionManager) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*ExtensionManagerOnlyExtensionOwnerOrOperator, error) {
	out := new(ExtensionManagerOnlyExtensionOwnerOrOperator)
	if err := extensionManager.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerOnlyGovernance represents a OnlyGovernance error raised by the ExtensionManager contract.
type ExtensionManagerOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func ExtensionManagerOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (extensionManager *ExtensionManager) UnpackOnlyGovernanceError(raw []byte) (*ExtensionManagerOnlyGovernance, error) {
	out := new(ExtensionManagerOnlyGovernance)
	if err := extensionManager.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerOnlyOwner represents a OnlyOwner error raised by the ExtensionManager contract.
type ExtensionManagerOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func ExtensionManagerOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (extensionManager *ExtensionManager) UnpackOnlyOwnerError(raw []byte) (*ExtensionManagerOnlyOwner, error) {
	out := new(ExtensionManagerOnlyOwner)
	if err := extensionManager.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerOnlyOwnerOrBackupManager represents a OnlyOwnerOrBackupManager error raised by the ExtensionManager contract.
type ExtensionManagerOnlyOwnerOrBackupManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrBackupManager()
func ExtensionManagerOnlyOwnerOrBackupManagerErrorID() common.Hash {
	return common.HexToHash("0x249a6cf6e22092f9c80798e2a0c583c0f36aea0779ef38a5b5cdb6a112a0e6df")
}

// UnpackOnlyOwnerOrBackupManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrBackupManager()
func (extensionManager *ExtensionManager) UnpackOnlyOwnerOrBackupManagerError(raw []byte) (*ExtensionManagerOnlyOwnerOrBackupManager, error) {
	out := new(ExtensionManagerOnlyOwnerOrBackupManager)
	if err := extensionManager.abi.UnpackIntoInterface(out, "OnlyOwnerOrBackupManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the ExtensionManager contract.
type ExtensionManagerOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func ExtensionManagerOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (extensionManager *ExtensionManager) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*ExtensionManagerOnlyProductionOrPausedStatus, error) {
	out := new(ExtensionManagerOnlyProductionOrPausedStatus)
	if err := extensionManager.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerOnlyProposedOwner represents a OnlyProposedOwner error raised by the ExtensionManager contract.
type ExtensionManagerOnlyProposedOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProposedOwner()
func ExtensionManagerOnlyProposedOwnerErrorID() common.Hash {
	return common.HexToHash("0x49cb15061ee558e74f66f937f740fd2a3791a89d2d139862222e5c97cabc02cd")
}

// UnpackOnlyProposedOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProposedOwner()
func (extensionManager *ExtensionManager) UnpackOnlyProposedOwnerError(raw []byte) (*ExtensionManagerOnlyProposedOwner, error) {
	out := new(ExtensionManagerOnlyProposedOwner)
	if err := extensionManager.abi.UnpackIntoInterface(out, "OnlyProposedOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerOwnerNotAllowed represents a OwnerNotAllowed error raised by the ExtensionManager contract.
type ExtensionManagerOwnerNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerNotAllowed()
func ExtensionManagerOwnerNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x3d5c9996545c6bf4287c26c1c2d8e6b4c39b92f19ba706c3a756720bf9c4fd24")
}

// UnpackOwnerNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerNotAllowed()
func (extensionManager *ExtensionManager) UnpackOwnerNotAllowedError(raw []byte) (*ExtensionManagerOwnerNotAllowed, error) {
	out := new(ExtensionManagerOwnerNotAllowed)
	if err := extensionManager.abi.UnpackIntoInterface(out, "OwnerNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerPlatformAlreadyExists represents a PlatformAlreadyExists error raised by the ExtensionManager contract.
type ExtensionManagerPlatformAlreadyExists struct {
	Platform [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PlatformAlreadyExists(bytes32 platform)
func ExtensionManagerPlatformAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0xb57a12b46c1d78c126c394c0042eabdd8d9c16d2356e0a7a57fad1a10adfdd8d")
}

// UnpackPlatformAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PlatformAlreadyExists(bytes32 platform)
func (extensionManager *ExtensionManager) UnpackPlatformAlreadyExistsError(raw []byte) (*ExtensionManagerPlatformAlreadyExists, error) {
	out := new(ExtensionManagerPlatformAlreadyExists)
	if err := extensionManager.abi.UnpackIntoInterface(out, "PlatformAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerPlatformEmpty represents a PlatformEmpty error raised by the ExtensionManager contract.
type ExtensionManagerPlatformEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PlatformEmpty()
func ExtensionManagerPlatformEmptyErrorID() common.Hash {
	return common.HexToHash("0x3d98d7260ca455a0a395aac9826dc9e850bce1cc9ed271e940d566c7e73e16ce")
}

// UnpackPlatformEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PlatformEmpty()
func (extensionManager *ExtensionManager) UnpackPlatformEmptyError(raw []byte) (*ExtensionManagerPlatformEmpty, error) {
	out := new(ExtensionManagerPlatformEmpty)
	if err := extensionManager.abi.UnpackIntoInterface(out, "PlatformEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerPlatformNotFound represents a PlatformNotFound error raised by the ExtensionManager contract.
type ExtensionManagerPlatformNotFound struct {
	Platform [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PlatformNotFound(bytes32 platform)
func ExtensionManagerPlatformNotFoundErrorID() common.Hash {
	return common.HexToHash("0x1df04ab2ad3da4c4747c4f8082795f36bb9263e098812c2643bd1507cda9269c")
}

// UnpackPlatformNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PlatformNotFound(bytes32 platform)
func (extensionManager *ExtensionManager) UnpackPlatformNotFoundError(raw []byte) (*ExtensionManagerPlatformNotFound, error) {
	out := new(ExtensionManagerPlatformNotFound)
	if err := extensionManager.abi.UnpackIntoInterface(out, "PlatformNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerReservedExtensionIdAlreadyAssigned represents a ReservedExtensionIdAlreadyAssigned error raised by the ExtensionManager contract.
type ExtensionManagerReservedExtensionIdAlreadyAssigned struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReservedExtensionIdAlreadyAssigned()
func ExtensionManagerReservedExtensionIdAlreadyAssignedErrorID() common.Hash {
	return common.HexToHash("0x69330ce437ea323be06a1b43e644523c21215dbbc758061356d745ce14c71349")
}

// UnpackReservedExtensionIdAlreadyAssignedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReservedExtensionIdAlreadyAssigned()
func (extensionManager *ExtensionManager) UnpackReservedExtensionIdAlreadyAssignedError(raw []byte) (*ExtensionManagerReservedExtensionIdAlreadyAssigned, error) {
	out := new(ExtensionManagerReservedExtensionIdAlreadyAssigned)
	if err := extensionManager.abi.UnpackIntoInterface(out, "ReservedExtensionIdAlreadyAssigned", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerSigningAlgoAlreadyExists represents a SigningAlgoAlreadyExists error raised by the ExtensionManager contract.
type ExtensionManagerSigningAlgoAlreadyExists struct {
	KeyType     [32]byte
	SigningAlgo [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SigningAlgoAlreadyExists(bytes32 keyType, bytes32 signingAlgo)
func ExtensionManagerSigningAlgoAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x646308c9301cdd87e37fa1b5970a872fc2aad39fd4157776b199ca19c0d54395")
}

// UnpackSigningAlgoAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SigningAlgoAlreadyExists(bytes32 keyType, bytes32 signingAlgo)
func (extensionManager *ExtensionManager) UnpackSigningAlgoAlreadyExistsError(raw []byte) (*ExtensionManagerSigningAlgoAlreadyExists, error) {
	out := new(ExtensionManagerSigningAlgoAlreadyExists)
	if err := extensionManager.abi.UnpackIntoInterface(out, "SigningAlgoAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerSigningAlgoEmpty represents a SigningAlgoEmpty error raised by the ExtensionManager contract.
type ExtensionManagerSigningAlgoEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SigningAlgoEmpty()
func ExtensionManagerSigningAlgoEmptyErrorID() common.Hash {
	return common.HexToHash("0x1c96eff2925c4790d507bc64ebf172386223c44a0c904a1d72ecc4c93a9168ed")
}

// UnpackSigningAlgoEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SigningAlgoEmpty()
func (extensionManager *ExtensionManager) UnpackSigningAlgoEmptyError(raw []byte) (*ExtensionManagerSigningAlgoEmpty, error) {
	out := new(ExtensionManagerSigningAlgoEmpty)
	if err := extensionManager.abi.UnpackIntoInterface(out, "SigningAlgoEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerSigningAlgoNotFound represents a SigningAlgoNotFound error raised by the ExtensionManager contract.
type ExtensionManagerSigningAlgoNotFound struct {
	KeyType     [32]byte
	SigningAlgo [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SigningAlgoNotFound(bytes32 keyType, bytes32 signingAlgo)
func ExtensionManagerSigningAlgoNotFoundErrorID() common.Hash {
	return common.HexToHash("0xd267ebb79f522e825ffc575f27e870e8a01497627e91af1c04dbf6a4eaf6f66f")
}

// UnpackSigningAlgoNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SigningAlgoNotFound(bytes32 keyType, bytes32 signingAlgo)
func (extensionManager *ExtensionManager) UnpackSigningAlgoNotFoundError(raw []byte) (*ExtensionManagerSigningAlgoNotFound, error) {
	out := new(ExtensionManagerSigningAlgoNotFound)
	if err := extensionManager.abi.UnpackIntoInterface(out, "SigningAlgoNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerSystemOwnedExtensionId represents a SystemOwnedExtensionId error raised by the ExtensionManager contract.
type ExtensionManagerSystemOwnedExtensionId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SystemOwnedExtensionId()
func ExtensionManagerSystemOwnedExtensionIdErrorID() common.Hash {
	return common.HexToHash("0x22fb069bd8b96ad5f902cd7f78d7f660f9a00d7b6c25b9f452e20482de214a2f")
}

// UnpackSystemOwnedExtensionIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SystemOwnedExtensionId()
func (extensionManager *ExtensionManager) UnpackSystemOwnedExtensionIdError(raw []byte) (*ExtensionManagerSystemOwnedExtensionId, error) {
	out := new(ExtensionManagerSystemOwnedExtensionId)
	if err := extensionManager.abi.UnpackIntoInterface(out, "SystemOwnedExtensionId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the ExtensionManager contract.
type ExtensionManagerTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func ExtensionManagerTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (extensionManager *ExtensionManager) UnpackTeeMachineNotAvailableError(raw []byte) (*ExtensionManagerTeeMachineNotAvailable, error) {
	out := new(ExtensionManagerTeeMachineNotAvailable)
	if err := extensionManager.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerUnsupportedPlatform represents a UnsupportedPlatform error raised by the ExtensionManager contract.
type ExtensionManagerUnsupportedPlatform struct {
	Platform [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedPlatform(bytes32 platform)
func ExtensionManagerUnsupportedPlatformErrorID() common.Hash {
	return common.HexToHash("0xaff97f9376cfc9b9a52133ac76e01c2e8f7767cb478110b2987ba842793421fa")
}

// UnpackUnsupportedPlatformError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedPlatform(bytes32 platform)
func (extensionManager *ExtensionManager) UnpackUnsupportedPlatformError(raw []byte) (*ExtensionManagerUnsupportedPlatform, error) {
	out := new(ExtensionManagerUnsupportedPlatform)
	if err := extensionManager.abi.UnpackIntoInterface(out, "UnsupportedPlatform", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerVersionAlreadyExists represents a VersionAlreadyExists error raised by the ExtensionManager contract.
type ExtensionManagerVersionAlreadyExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionAlreadyExists()
func ExtensionManagerVersionAlreadyExistsErrorID() common.Hash {
	return common.HexToHash("0x13db702b902fc830978b5d9709400c3a7d363828f944a655aec9698bbd9abe29")
}

// UnpackVersionAlreadyExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionAlreadyExists()
func (extensionManager *ExtensionManager) UnpackVersionAlreadyExistsError(raw []byte) (*ExtensionManagerVersionAlreadyExists, error) {
	out := new(ExtensionManagerVersionAlreadyExists)
	if err := extensionManager.abi.UnpackIntoInterface(out, "VersionAlreadyExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerVersionEmpty represents a VersionEmpty error raised by the ExtensionManager contract.
type ExtensionManagerVersionEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionEmpty()
func ExtensionManagerVersionEmptyErrorID() common.Hash {
	return common.HexToHash("0xd4dde0ae5d6d46bbe7c328853080e44c95bb07c6472bb940569702cc4f153a9d")
}

// UnpackVersionEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionEmpty()
func (extensionManager *ExtensionManager) UnpackVersionEmptyError(raw []byte) (*ExtensionManagerVersionEmpty, error) {
	out := new(ExtensionManagerVersionEmpty)
	if err := extensionManager.abi.UnpackIntoInterface(out, "VersionEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// ExtensionManagerVersionNotSupported represents a VersionNotSupported error raised by the ExtensionManager contract.
type ExtensionManagerVersionNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionNotSupported()
func ExtensionManagerVersionNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x6db0b17b06266b37e721024a8ba1221439799ebb779dfd397865e37250abc9a8")
}

// UnpackVersionNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionNotSupported()
func (extensionManager *ExtensionManager) UnpackVersionNotSupportedError(raw []byte) (*ExtensionManagerVersionNotSupported, error) {
	out := new(ExtensionManagerVersionNotSupported)
	if err := extensionManager.abi.UnpackIntoInterface(out, "VersionNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}
