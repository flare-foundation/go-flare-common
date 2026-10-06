// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletkeymanager

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

// IWalletKeyManagerKeyConfigConstants is an auto generated low-level Go binding around an user-defined struct.
type IWalletKeyManagerKeyConfigConstants struct {
	AdminsPublicKeys   []PublicKey
	AdminsThreshold    uint64
	Cosigners          []common.Address
	CosignersThreshold uint64
}

// IWalletKeyManagerKeyExistence is an auto generated low-level Go binding around an user-defined struct.
type IWalletKeyManagerKeyExistence struct {
	TeeId           common.Address
	WalletId        [32]byte
	KeyId           uint64
	KeyType         [32]byte
	SigningAlgo     [32]byte
	PublicKey       []byte
	Nonce           *big.Int
	Restored        bool
	ConfigConstants IWalletKeyManagerKeyConfigConstants
	SettingsVersion [32]byte
	Settings        []byte
}

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

// TeeIdKeyIdPair is an auto generated low-level Go binding around an user-defined struct.
type TeeIdKeyIdPair struct {
	TeeId common.Address
	KeyId uint64
}

// WalletKeyManagerMetaData contains all meta data concerning the WalletKeyManager contract.
var WalletKeyManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyInSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressNotInSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"}],\"name\":\"EmergencyPauseActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSettings\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyNotGeneratedOnTeeMachine\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyNotRestoredOnTeeMachine\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrPauser\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"NotOwnerOrUnpauser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwnerOrOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeIdAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManager.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"name\":\"WalletKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"WalletKeyConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"name\":\"WalletKeyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"keyIds\",\"type\":\"uint64[]\"}],\"name\":\"WalletKeysNotAvailable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"multisigThreshold\",\"type\":\"uint64\"}],\"name\":\"WalletMultisigThresholdSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"addKey\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"cleanUpTeeIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restored\",\"type\":\"bool\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"adminsThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"internalType\":\"structIWalletKeyManager.KeyConfigConstants\",\"name\":\"configConstants\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"settingsVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"settings\",\"type\":\"bytes\"}],\"internalType\":\"structIWalletKeyManager.KeyExistence\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"_teeSignature\",\"type\":\"tuple\"}],\"name\":\"confirmKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"deleteKey\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"getKeyNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_teeHoldsKey\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getReceivingTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"getWalletKeyPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"getWalletKeyTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletKeysInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_multisigThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64[]\",\"name\":\"_keyIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"_counter\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletPublicKeys\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_multisigThreshold\",\"type\":\"uint64\"},{\"internalType\":\"bytes[]\",\"name\":\"_publicKeys\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"receivingTeesAndKeys\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"internalType\":\"structTeeIdKeyIdPair[]\",\"name\":\"_teeIdKeyIdPairs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_multisigThreshold\",\"type\":\"uint64\"}],\"name\":\"setMultisigThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "WalletKeyManager",
}

// WalletKeyManager is an auto generated Go binding around an Ethereum contract.
type WalletKeyManager struct {
	abi abi.ABI
}

// NewWalletKeyManager creates a new instance of WalletKeyManager.
func NewWalletKeyManager() *WalletKeyManager {
	parsed, err := WalletKeyManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &WalletKeyManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *WalletKeyManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba9e7a82.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId, address _claimBackAddress) payable returns(uint64 _keyId)
func (walletKeyManager *WalletKeyManager) PackAddKey(teeId common.Address, walletId [32]byte, claimBackAddress common.Address) []byte {
	enc, err := walletKeyManager.abi.Pack("addKey", teeId, walletId, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xba9e7a82.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId, address _claimBackAddress) payable returns(uint64 _keyId)
func (walletKeyManager *WalletKeyManager) TryPackAddKey(teeId common.Address, walletId [32]byte, claimBackAddress common.Address) ([]byte, error) {
	return walletKeyManager.abi.Pack("addKey", teeId, walletId, claimBackAddress)
}

// UnpackAddKey is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xba9e7a82.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId, address _claimBackAddress) payable returns(uint64 _keyId)
func (walletKeyManager *WalletKeyManager) UnpackAddKey(data []byte) (uint64, error) {
	out, err := walletKeyManager.abi.Unpack("addKey", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackCleanUpTeeIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaec4c979.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (walletKeyManager *WalletKeyManager) PackCleanUpTeeIds(walletId [32]byte, keyId uint64) []byte {
	enc, err := walletKeyManager.abi.Pack("cleanUpTeeIds", walletId, keyId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCleanUpTeeIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaec4c979.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (walletKeyManager *WalletKeyManager) TryPackCleanUpTeeIds(walletId [32]byte, keyId uint64) ([]byte, error) {
	return walletKeyManager.abi.Pack("cleanUpTeeIds", walletId, keyId)
}

// PackConfirmKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa02d6fe1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function confirmKey((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) _proof, (uint8,bytes32,bytes32) _teeSignature) returns()
func (walletKeyManager *WalletKeyManager) PackConfirmKey(proof IWalletKeyManagerKeyExistence, teeSignature Signature) []byte {
	enc, err := walletKeyManager.abi.Pack("confirmKey", proof, teeSignature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConfirmKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa02d6fe1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function confirmKey((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) _proof, (uint8,bytes32,bytes32) _teeSignature) returns()
func (walletKeyManager *WalletKeyManager) TryPackConfirmKey(proof IWalletKeyManagerKeyExistence, teeSignature Signature) ([]byte, error) {
	return walletKeyManager.abi.Pack("confirmKey", proof, teeSignature)
}

// PackDeleteKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x270de236.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId, address _claimBackAddress) payable returns()
func (walletKeyManager *WalletKeyManager) PackDeleteKey(teeId common.Address, walletId [32]byte, keyId uint64, claimBackAddress common.Address) []byte {
	enc, err := walletKeyManager.abi.Pack("deleteKey", teeId, walletId, keyId, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDeleteKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x270de236.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId, address _claimBackAddress) payable returns()
func (walletKeyManager *WalletKeyManager) TryPackDeleteKey(teeId common.Address, walletId [32]byte, keyId uint64, claimBackAddress common.Address) ([]byte, error) {
	return walletKeyManager.abi.Pack("deleteKey", teeId, walletId, keyId, claimBackAddress)
}

// PackGetKeyNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x60d8ff13.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getKeyNonce(address _teeId, bytes32 _walletId, uint64 _keyId) view returns(uint256 _nonce, bool _teeHoldsKey)
func (walletKeyManager *WalletKeyManager) PackGetKeyNonce(teeId common.Address, walletId [32]byte, keyId uint64) []byte {
	enc, err := walletKeyManager.abi.Pack("getKeyNonce", teeId, walletId, keyId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetKeyNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x60d8ff13.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getKeyNonce(address _teeId, bytes32 _walletId, uint64 _keyId) view returns(uint256 _nonce, bool _teeHoldsKey)
func (walletKeyManager *WalletKeyManager) TryPackGetKeyNonce(teeId common.Address, walletId [32]byte, keyId uint64) ([]byte, error) {
	return walletKeyManager.abi.Pack("getKeyNonce", teeId, walletId, keyId)
}

// GetKeyNonceOutput serves as a container for the return parameters of contract
// method GetKeyNonce.
type GetKeyNonceOutput struct {
	Nonce       *big.Int
	TeeHoldsKey bool
}

// UnpackGetKeyNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x60d8ff13.
//
// Solidity: function getKeyNonce(address _teeId, bytes32 _walletId, uint64 _keyId) view returns(uint256 _nonce, bool _teeHoldsKey)
func (walletKeyManager *WalletKeyManager) UnpackGetKeyNonce(data []byte) (GetKeyNonceOutput, error) {
	out, err := walletKeyManager.abi.Unpack("getKeyNonce", data)
	outstruct := new(GetKeyNonceOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Nonce = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.TeeHoldsKey = *abi.ConvertType(out[1], new(bool)).(*bool)
	return *outstruct, nil
}

// PackGetReceivingTeeIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8184ce4b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getReceivingTeeIds(bytes32 _walletId) view returns(address[] _teeIds)
func (walletKeyManager *WalletKeyManager) PackGetReceivingTeeIds(walletId [32]byte) []byte {
	enc, err := walletKeyManager.abi.Pack("getReceivingTeeIds", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetReceivingTeeIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8184ce4b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getReceivingTeeIds(bytes32 _walletId) view returns(address[] _teeIds)
func (walletKeyManager *WalletKeyManager) TryPackGetReceivingTeeIds(walletId [32]byte) ([]byte, error) {
	return walletKeyManager.abi.Pack("getReceivingTeeIds", walletId)
}

// UnpackGetReceivingTeeIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8184ce4b.
//
// Solidity: function getReceivingTeeIds(bytes32 _walletId) view returns(address[] _teeIds)
func (walletKeyManager *WalletKeyManager) UnpackGetReceivingTeeIds(data []byte) ([]common.Address, error) {
	out, err := walletKeyManager.abi.Unpack("getReceivingTeeIds", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetWalletKeyPublicKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb825ac3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (walletKeyManager *WalletKeyManager) PackGetWalletKeyPublicKey(walletId [32]byte, keyId uint64) []byte {
	enc, err := walletKeyManager.abi.Pack("getWalletKeyPublicKey", walletId, keyId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletKeyPublicKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb825ac3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (walletKeyManager *WalletKeyManager) TryPackGetWalletKeyPublicKey(walletId [32]byte, keyId uint64) ([]byte, error) {
	return walletKeyManager.abi.Pack("getWalletKeyPublicKey", walletId, keyId)
}

// UnpackGetWalletKeyPublicKey is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbb825ac3.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (walletKeyManager *WalletKeyManager) UnpackGetWalletKeyPublicKey(data []byte) ([]byte, error) {
	out, err := walletKeyManager.abi.Unpack("getWalletKeyPublicKey", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackGetWalletKeyTeeIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x08480213.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (walletKeyManager *WalletKeyManager) PackGetWalletKeyTeeIds(walletId [32]byte, keyId uint64) []byte {
	enc, err := walletKeyManager.abi.Pack("getWalletKeyTeeIds", walletId, keyId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletKeyTeeIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x08480213.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (walletKeyManager *WalletKeyManager) TryPackGetWalletKeyTeeIds(walletId [32]byte, keyId uint64) ([]byte, error) {
	return walletKeyManager.abi.Pack("getWalletKeyTeeIds", walletId, keyId)
}

// UnpackGetWalletKeyTeeIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x08480213.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (walletKeyManager *WalletKeyManager) UnpackGetWalletKeyTeeIds(data []byte) ([]common.Address, error) {
	out, err := walletKeyManager.abi.Unpack("getWalletKeyTeeIds", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetWalletKeysInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8338292b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint64[] _keyIds, uint64 _counter)
func (walletKeyManager *WalletKeyManager) PackGetWalletKeysInfo(walletId [32]byte) []byte {
	enc, err := walletKeyManager.abi.Pack("getWalletKeysInfo", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletKeysInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8338292b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint64[] _keyIds, uint64 _counter)
func (walletKeyManager *WalletKeyManager) TryPackGetWalletKeysInfo(walletId [32]byte) ([]byte, error) {
	return walletKeyManager.abi.Pack("getWalletKeysInfo", walletId)
}

// GetWalletKeysInfoOutput serves as a container for the return parameters of contract
// method GetWalletKeysInfo.
type GetWalletKeysInfoOutput struct {
	MultisigThreshold uint64
	KeyIds            []uint64
	Counter           uint64
}

// UnpackGetWalletKeysInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint64[] _keyIds, uint64 _counter)
func (walletKeyManager *WalletKeyManager) UnpackGetWalletKeysInfo(data []byte) (GetWalletKeysInfoOutput, error) {
	out, err := walletKeyManager.abi.Unpack("getWalletKeysInfo", data)
	outstruct := new(GetWalletKeysInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.MultisigThreshold = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.KeyIds = *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)
	outstruct.Counter = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetWalletPublicKeys is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6205085f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletPublicKeys(bytes32 _walletId) view returns(uint64 _multisigThreshold, bytes[] _publicKeys)
func (walletKeyManager *WalletKeyManager) PackGetWalletPublicKeys(walletId [32]byte) []byte {
	enc, err := walletKeyManager.abi.Pack("getWalletPublicKeys", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletPublicKeys is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6205085f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletPublicKeys(bytes32 _walletId) view returns(uint64 _multisigThreshold, bytes[] _publicKeys)
func (walletKeyManager *WalletKeyManager) TryPackGetWalletPublicKeys(walletId [32]byte) ([]byte, error) {
	return walletKeyManager.abi.Pack("getWalletPublicKeys", walletId)
}

// GetWalletPublicKeysOutput serves as a container for the return parameters of contract
// method GetWalletPublicKeys.
type GetWalletPublicKeysOutput struct {
	MultisigThreshold uint64
	PublicKeys        [][]byte
}

// UnpackGetWalletPublicKeys is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6205085f.
//
// Solidity: function getWalletPublicKeys(bytes32 _walletId) view returns(uint64 _multisigThreshold, bytes[] _publicKeys)
func (walletKeyManager *WalletKeyManager) UnpackGetWalletPublicKeys(data []byte) (GetWalletPublicKeysOutput, error) {
	out, err := walletKeyManager.abi.Unpack("getWalletPublicKeys", data)
	outstruct := new(GetWalletPublicKeysOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.MultisigThreshold = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.PublicKeys = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)
	return *outstruct, nil
}

// PackReceivingTeesAndKeys is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaddcc06.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,uint64)[] _teeIdKeyIdPairs)
func (walletKeyManager *WalletKeyManager) PackReceivingTeesAndKeys(walletId [32]byte) []byte {
	enc, err := walletKeyManager.abi.Pack("receivingTeesAndKeys", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReceivingTeesAndKeys is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaddcc06.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,uint64)[] _teeIdKeyIdPairs)
func (walletKeyManager *WalletKeyManager) TryPackReceivingTeesAndKeys(walletId [32]byte) ([]byte, error) {
	return walletKeyManager.abi.Pack("receivingTeesAndKeys", walletId)
}

// UnpackReceivingTeesAndKeys is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaddcc06.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,uint64)[] _teeIdKeyIdPairs)
func (walletKeyManager *WalletKeyManager) UnpackReceivingTeesAndKeys(data []byte) ([]TeeIdKeyIdPair, error) {
	out, err := walletKeyManager.abi.Unpack("receivingTeesAndKeys", data)
	if err != nil {
		return *new([]TeeIdKeyIdPair), err
	}
	out0 := *abi.ConvertType(out[0], new([]TeeIdKeyIdPair)).(*[]TeeIdKeyIdPair)
	return out0, nil
}

// PackSetMultisigThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62101d0a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint64 _multisigThreshold) returns()
func (walletKeyManager *WalletKeyManager) PackSetMultisigThreshold(walletId [32]byte, multisigThreshold uint64) []byte {
	enc, err := walletKeyManager.abi.Pack("setMultisigThreshold", walletId, multisigThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetMultisigThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62101d0a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint64 _multisigThreshold) returns()
func (walletKeyManager *WalletKeyManager) TryPackSetMultisigThreshold(walletId [32]byte, multisigThreshold uint64) ([]byte, error) {
	return walletKeyManager.abi.Pack("setMultisigThreshold", walletId, multisigThreshold)
}

// WalletKeyManagerTeeInstructionsSent represents a TeeInstructionsSent event raised by the WalletKeyManager contract.
type WalletKeyManagerTeeInstructionsSent struct {
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

const WalletKeyManagerTeeInstructionsSentEventName = "TeeInstructionsSent"

// ContractEventName returns the user-defined event name.
func (WalletKeyManagerTeeInstructionsSent) ContractEventName() string {
	return WalletKeyManagerTeeInstructionsSentEventName
}

// UnpackTeeInstructionsSentEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (walletKeyManager *WalletKeyManager) UnpackTeeInstructionsSentEvent(log *types.Log) (*WalletKeyManagerTeeInstructionsSent, error) {
	event := "TeeInstructionsSent"
	if len(log.Topics) == 0 || log.Topics[0] != walletKeyManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletKeyManagerTeeInstructionsSent)
	if len(log.Data) > 0 {
		if err := walletKeyManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletKeyManager.abi.Events[event].Inputs {
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

// WalletKeyManagerWalletKeyAdded represents a WalletKeyAdded event raised by the WalletKeyManager contract.
type WalletKeyManagerWalletKeyAdded struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    uint64
	Raw      *types.Log // Blockchain specific contextual infos
}

const WalletKeyManagerWalletKeyAddedEventName = "WalletKeyAdded"

// ContractEventName returns the user-defined event name.
func (WalletKeyManagerWalletKeyAdded) ContractEventName() string {
	return WalletKeyManagerWalletKeyAddedEventName
}

// UnpackWalletKeyAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletKeyAdded(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (walletKeyManager *WalletKeyManager) UnpackWalletKeyAddedEvent(log *types.Log) (*WalletKeyManagerWalletKeyAdded, error) {
	event := "WalletKeyAdded"
	if len(log.Topics) == 0 || log.Topics[0] != walletKeyManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletKeyManagerWalletKeyAdded)
	if len(log.Data) > 0 {
		if err := walletKeyManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletKeyManager.abi.Events[event].Inputs {
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

// WalletKeyManagerWalletKeyConfirmed represents a WalletKeyConfirmed event raised by the WalletKeyManager contract.
type WalletKeyManagerWalletKeyConfirmed struct {
	TeeId     common.Address
	WalletId  [32]byte
	KeyId     uint64
	PublicKey []byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletKeyManagerWalletKeyConfirmedEventName = "WalletKeyConfirmed"

// ContractEventName returns the user-defined event name.
func (WalletKeyManagerWalletKeyConfirmed) ContractEventName() string {
	return WalletKeyManagerWalletKeyConfirmedEventName
}

// UnpackWalletKeyConfirmedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletKeyConfirmed(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId, bytes publicKey)
func (walletKeyManager *WalletKeyManager) UnpackWalletKeyConfirmedEvent(log *types.Log) (*WalletKeyManagerWalletKeyConfirmed, error) {
	event := "WalletKeyConfirmed"
	if len(log.Topics) == 0 || log.Topics[0] != walletKeyManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletKeyManagerWalletKeyConfirmed)
	if len(log.Data) > 0 {
		if err := walletKeyManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletKeyManager.abi.Events[event].Inputs {
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

// WalletKeyManagerWalletKeyDeleted represents a WalletKeyDeleted event raised by the WalletKeyManager contract.
type WalletKeyManagerWalletKeyDeleted struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    uint64
	Raw      *types.Log // Blockchain specific contextual infos
}

const WalletKeyManagerWalletKeyDeletedEventName = "WalletKeyDeleted"

// ContractEventName returns the user-defined event name.
func (WalletKeyManagerWalletKeyDeleted) ContractEventName() string {
	return WalletKeyManagerWalletKeyDeletedEventName
}

// UnpackWalletKeyDeletedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletKeyDeleted(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (walletKeyManager *WalletKeyManager) UnpackWalletKeyDeletedEvent(log *types.Log) (*WalletKeyManagerWalletKeyDeleted, error) {
	event := "WalletKeyDeleted"
	if len(log.Topics) == 0 || log.Topics[0] != walletKeyManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletKeyManagerWalletKeyDeleted)
	if len(log.Data) > 0 {
		if err := walletKeyManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletKeyManager.abi.Events[event].Inputs {
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

// WalletKeyManagerWalletKeysNotAvailable represents a WalletKeysNotAvailable event raised by the WalletKeyManager contract.
type WalletKeyManagerWalletKeysNotAvailable struct {
	WalletId [32]byte
	KeyIds   []uint64
	Raw      *types.Log // Blockchain specific contextual infos
}

const WalletKeyManagerWalletKeysNotAvailableEventName = "WalletKeysNotAvailable"

// ContractEventName returns the user-defined event name.
func (WalletKeyManagerWalletKeysNotAvailable) ContractEventName() string {
	return WalletKeyManagerWalletKeysNotAvailableEventName
}

// UnpackWalletKeysNotAvailableEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint64[] keyIds)
func (walletKeyManager *WalletKeyManager) UnpackWalletKeysNotAvailableEvent(log *types.Log) (*WalletKeyManagerWalletKeysNotAvailable, error) {
	event := "WalletKeysNotAvailable"
	if len(log.Topics) == 0 || log.Topics[0] != walletKeyManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletKeyManagerWalletKeysNotAvailable)
	if len(log.Data) > 0 {
		if err := walletKeyManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletKeyManager.abi.Events[event].Inputs {
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

// WalletKeyManagerWalletMultisigThresholdSet represents a WalletMultisigThresholdSet event raised by the WalletKeyManager contract.
type WalletKeyManagerWalletMultisigThresholdSet struct {
	WalletId          [32]byte
	MultisigThreshold uint64
	Raw               *types.Log // Blockchain specific contextual infos
}

const WalletKeyManagerWalletMultisigThresholdSetEventName = "WalletMultisigThresholdSet"

// ContractEventName returns the user-defined event name.
func (WalletKeyManagerWalletMultisigThresholdSet) ContractEventName() string {
	return WalletKeyManagerWalletMultisigThresholdSetEventName
}

// UnpackWalletMultisigThresholdSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WalletMultisigThresholdSet(bytes32 indexed walletId, uint64 multisigThreshold)
func (walletKeyManager *WalletKeyManager) UnpackWalletMultisigThresholdSetEvent(log *types.Log) (*WalletKeyManagerWalletMultisigThresholdSet, error) {
	event := "WalletMultisigThresholdSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletKeyManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletKeyManagerWalletMultisigThresholdSet)
	if len(log.Data) > 0 {
		if err := walletKeyManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletKeyManager.abi.Events[event].Inputs {
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
func (walletKeyManager *WalletKeyManager) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["AddressAlreadyInSet"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackAddressAlreadyInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["AddressNotInSet"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackAddressNotInSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["AvailabilityCheckTimestampInvalid"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackAvailabilityCheckTimestampInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["CosignersThresholdTooHigh"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackCosignersThresholdTooHighError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["DuplicatedCosigner"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackDuplicatedCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["EmergencyPauseActive"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackEmergencyPauseActiveError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["ExtensionIdMismatch"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackExtensionIdMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["FeeTooLow"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackFeeTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidAddress"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidAvailabilityCheckStatus"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidAvailabilityCheckStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidDuration"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidDurationError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidGovernanceHash"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidGovernanceHashError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidKeyId"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidKeyIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidKeyType"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidNonce"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidNonceError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidPublicKey"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidPublicKeyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidResponseData"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidResponseDataError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidSettings"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidSettingsError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidSigningAlgo"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidSigningAlgoError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidTeeSignature"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidTeeSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidThreshold"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidThresholdError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["InvalidWalletStatus"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackInvalidWalletStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["KeyNotGeneratedOnTeeMachine"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackKeyNotGeneratedOnTeeMachineError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["KeyNotRestoredOnTeeMachine"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackKeyNotRestoredOnTeeMachineError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["KeyTypeNotSupported"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["MessageEmpty"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackMessageEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["NoAddresses"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackNoAddressesError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["NoTeeMachinesSpecified"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackNoTeeMachinesSpecifiedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["NotOwnerOrPauser"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackNotOwnerOrPauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["NotOwnerOrUnpauser"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackNotOwnerOrUnpauserError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["OnlyExtensionOwner"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackOnlyExtensionOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["OnlyExtensionOwnerOrOperator"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackOnlyExtensionOwnerOrOperatorError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["OnlyOwner"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackOnlyOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["OnlyOwnerOrBackupManager"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackOnlyOwnerOrBackupManagerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["OnlyProposedOwner"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackOnlyProposedOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["OperationCommandEmpty"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackOperationCommandEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["OperationTypeEmpty"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackOperationTypeEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["OwnerNotAllowed"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackOwnerNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["TeeIdAlreadyAdded"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackTeeIdAlreadyAddedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["TeeMachineNotAvailable"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackTeeMachineNotAvailableError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["TeeNotFound"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackTeeNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["ThresholdNotMet"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackThresholdNotMetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletKeyManager.abi.Errors["VersionNotSupported"].ID.Bytes()[:4]) {
		return walletKeyManager.UnpackVersionNotSupportedError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// WalletKeyManagerAddressAlreadyInSet represents a AddressAlreadyInSet error raised by the WalletKeyManager contract.
type WalletKeyManagerAddressAlreadyInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressAlreadyInSet(address addr)
func WalletKeyManagerAddressAlreadyInSetErrorID() common.Hash {
	return common.HexToHash("0xb725c7fb1c3f0f065973d6b9eb55140f970259955af3045742cd5b37902eb920")
}

// UnpackAddressAlreadyInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressAlreadyInSet(address addr)
func (walletKeyManager *WalletKeyManager) UnpackAddressAlreadyInSetError(raw []byte) (*WalletKeyManagerAddressAlreadyInSet, error) {
	out := new(WalletKeyManagerAddressAlreadyInSet)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "AddressAlreadyInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerAddressNotInSet represents a AddressNotInSet error raised by the WalletKeyManager contract.
type WalletKeyManagerAddressNotInSet struct {
	Addr common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressNotInSet(address addr)
func WalletKeyManagerAddressNotInSetErrorID() common.Hash {
	return common.HexToHash("0x1ff5ecf3e6d385ea01ff0323fe894ed6d25a83ed5f6ff2a2ede4c897408f6be9")
}

// UnpackAddressNotInSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressNotInSet(address addr)
func (walletKeyManager *WalletKeyManager) UnpackAddressNotInSetError(raw []byte) (*WalletKeyManagerAddressNotInSet, error) {
	out := new(WalletKeyManagerAddressNotInSet)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "AddressNotInSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerAvailabilityCheckTimestampInvalid represents a AvailabilityCheckTimestampInvalid error raised by the WalletKeyManager contract.
type WalletKeyManagerAvailabilityCheckTimestampInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func WalletKeyManagerAvailabilityCheckTimestampInvalidErrorID() common.Hash {
	return common.HexToHash("0x6344a5e74cff78612fa271fa907ca608af3bad10e1e8c8df302e29ebc6298e2d")
}

// UnpackAvailabilityCheckTimestampInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AvailabilityCheckTimestampInvalid()
func (walletKeyManager *WalletKeyManager) UnpackAvailabilityCheckTimestampInvalidError(raw []byte) (*WalletKeyManagerAvailabilityCheckTimestampInvalid, error) {
	out := new(WalletKeyManagerAvailabilityCheckTimestampInvalid)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "AvailabilityCheckTimestampInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerCosignersThresholdTooHigh represents a CosignersThresholdTooHigh error raised by the WalletKeyManager contract.
type WalletKeyManagerCosignersThresholdTooHigh struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CosignersThresholdTooHigh()
func WalletKeyManagerCosignersThresholdTooHighErrorID() common.Hash {
	return common.HexToHash("0x3d7053512cbb0b649dba01ac4e3591f104fd7d0957128a39a3c552c77cd5014d")
}

// UnpackCosignersThresholdTooHighError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CosignersThresholdTooHigh()
func (walletKeyManager *WalletKeyManager) UnpackCosignersThresholdTooHighError(raw []byte) (*WalletKeyManagerCosignersThresholdTooHigh, error) {
	out := new(WalletKeyManagerCosignersThresholdTooHigh)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "CosignersThresholdTooHigh", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerDuplicatedCosigner represents a DuplicatedCosigner error raised by the WalletKeyManager contract.
type WalletKeyManagerDuplicatedCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func WalletKeyManagerDuplicatedCosignerErrorID() common.Hash {
	return common.HexToHash("0x53a5890417576fb3043dc3ff1669b0932da2f657d8adb4c11313b12b711f5910")
}

// UnpackDuplicatedCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicatedCosigner(address cosigner)
func (walletKeyManager *WalletKeyManager) UnpackDuplicatedCosignerError(raw []byte) (*WalletKeyManagerDuplicatedCosigner, error) {
	out := new(WalletKeyManagerDuplicatedCosigner)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "DuplicatedCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the WalletKeyManager contract.
type WalletKeyManagerECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func WalletKeyManagerECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (walletKeyManager *WalletKeyManager) UnpackECDSAInvalidSignatureError(raw []byte) (*WalletKeyManagerECDSAInvalidSignature, error) {
	out := new(WalletKeyManagerECDSAInvalidSignature)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the WalletKeyManager contract.
type WalletKeyManagerECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func WalletKeyManagerECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (walletKeyManager *WalletKeyManager) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*WalletKeyManagerECDSAInvalidSignatureLength, error) {
	out := new(WalletKeyManagerECDSAInvalidSignatureLength)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the WalletKeyManager contract.
type WalletKeyManagerECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func WalletKeyManagerECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (walletKeyManager *WalletKeyManager) UnpackECDSAInvalidSignatureSError(raw []byte) (*WalletKeyManagerECDSAInvalidSignatureS, error) {
	out := new(WalletKeyManagerECDSAInvalidSignatureS)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerEmergencyPauseActive represents a EmergencyPauseActive error raised by the WalletKeyManager contract.
type WalletKeyManagerEmergencyPauseActive struct {
	ExtensionId *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func WalletKeyManagerEmergencyPauseActiveErrorID() common.Hash {
	return common.HexToHash("0x480cf2163927c55f46b05a5cac80ac3781076b0e50b1ed6b07995a800937d4b6")
}

// UnpackEmergencyPauseActiveError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmergencyPauseActive(uint256 extensionId)
func (walletKeyManager *WalletKeyManager) UnpackEmergencyPauseActiveError(raw []byte) (*WalletKeyManagerEmergencyPauseActive, error) {
	out := new(WalletKeyManagerEmergencyPauseActive)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "EmergencyPauseActive", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerExtensionIdMismatch represents a ExtensionIdMismatch error raised by the WalletKeyManager contract.
type WalletKeyManagerExtensionIdMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ExtensionIdMismatch()
func WalletKeyManagerExtensionIdMismatchErrorID() common.Hash {
	return common.HexToHash("0x78969c25ba5fb1d5055b45511bfef89eba6076b306b02133216054eb56f33e49")
}

// UnpackExtensionIdMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ExtensionIdMismatch()
func (walletKeyManager *WalletKeyManager) UnpackExtensionIdMismatchError(raw []byte) (*WalletKeyManagerExtensionIdMismatch, error) {
	out := new(WalletKeyManagerExtensionIdMismatch)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "ExtensionIdMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerFeeTooLow represents a FeeTooLow error raised by the WalletKeyManager contract.
type WalletKeyManagerFeeTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeTooLow()
func WalletKeyManagerFeeTooLowErrorID() common.Hash {
	return common.HexToHash("0x732f94137528ce257864c792a6de53ff98f1203051e0a27c454063f559602458")
}

// UnpackFeeTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeTooLow()
func (walletKeyManager *WalletKeyManager) UnpackFeeTooLowError(raw []byte) (*WalletKeyManagerFeeTooLow, error) {
	out := new(WalletKeyManagerFeeTooLow)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "FeeTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidAddress represents a InvalidAddress error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAddress()
func WalletKeyManagerInvalidAddressErrorID() common.Hash {
	return common.HexToHash("0xe6c4247b90bd06996a32d386bb770af9c0018dd1b0ebbb2df2c4499f1eda7b16")
}

// UnpackInvalidAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAddress()
func (walletKeyManager *WalletKeyManager) UnpackInvalidAddressError(raw []byte) (*WalletKeyManagerInvalidAddress, error) {
	out := new(WalletKeyManagerInvalidAddress)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidAvailabilityCheckStatus represents a InvalidAvailabilityCheckStatus error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidAvailabilityCheckStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func WalletKeyManagerInvalidAvailabilityCheckStatusErrorID() common.Hash {
	return common.HexToHash("0xedee4ff00b842d3ef84031c60001065c540c838c5479bbf0ae56f15868427e44")
}

// UnpackInvalidAvailabilityCheckStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAvailabilityCheckStatus()
func (walletKeyManager *WalletKeyManager) UnpackInvalidAvailabilityCheckStatusError(raw []byte) (*WalletKeyManagerInvalidAvailabilityCheckStatus, error) {
	out := new(WalletKeyManagerInvalidAvailabilityCheckStatus)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidAvailabilityCheckStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidCosigner represents a InvalidCosigner error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func WalletKeyManagerInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (walletKeyManager *WalletKeyManager) UnpackInvalidCosignerError(raw []byte) (*WalletKeyManagerInvalidCosigner, error) {
	out := new(WalletKeyManagerInvalidCosigner)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidDuration represents a InvalidDuration error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidDuration struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDuration()
func WalletKeyManagerInvalidDurationErrorID() common.Hash {
	return common.HexToHash("0x76166401494e5a511541446d9197f2558d8ed87f455106fcd343b5ffbb4fc7bb")
}

// UnpackInvalidDurationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDuration()
func (walletKeyManager *WalletKeyManager) UnpackInvalidDurationError(raw []byte) (*WalletKeyManagerInvalidDuration, error) {
	out := new(WalletKeyManagerInvalidDuration)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidDuration", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidGovernanceHash represents a InvalidGovernanceHash error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidGovernanceHash struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGovernanceHash()
func WalletKeyManagerInvalidGovernanceHashErrorID() common.Hash {
	return common.HexToHash("0x343ca6286bce3d98937ea43eca2b4d6e2ec2be1a6d379fd9a54faa59408f0d1f")
}

// UnpackInvalidGovernanceHashError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGovernanceHash()
func (walletKeyManager *WalletKeyManager) UnpackInvalidGovernanceHashError(raw []byte) (*WalletKeyManagerInvalidGovernanceHash, error) {
	out := new(WalletKeyManagerInvalidGovernanceHash)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidGovernanceHash", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidKeyId represents a InvalidKeyId error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidKeyId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyId()
func WalletKeyManagerInvalidKeyIdErrorID() common.Hash {
	return common.HexToHash("0xb0aeb53e4a83bd3cc8ab2c55982004930d12942a439bdf74bdc4a5d0805f9e0c")
}

// UnpackInvalidKeyIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyId()
func (walletKeyManager *WalletKeyManager) UnpackInvalidKeyIdError(raw []byte) (*WalletKeyManagerInvalidKeyId, error) {
	out := new(WalletKeyManagerInvalidKeyId)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidKeyId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidKeyType represents a InvalidKeyType error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidKeyType()
func WalletKeyManagerInvalidKeyTypeErrorID() common.Hash {
	return common.HexToHash("0x36dccb440beb343622d8ee0927bce3f69281feea3a5b0610f6e00e2811b557e1")
}

// UnpackInvalidKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidKeyType()
func (walletKeyManager *WalletKeyManager) UnpackInvalidKeyTypeError(raw []byte) (*WalletKeyManagerInvalidKeyType, error) {
	out := new(WalletKeyManagerInvalidKeyType)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidNonce represents a InvalidNonce error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidNonce struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidNonce()
func WalletKeyManagerInvalidNonceErrorID() common.Hash {
	return common.HexToHash("0x756688fec2871909d72599c334b663ffcc94654c438569966c7fd3ab3a351f34")
}

// UnpackInvalidNonceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidNonce()
func (walletKeyManager *WalletKeyManager) UnpackInvalidNonceError(raw []byte) (*WalletKeyManagerInvalidNonce, error) {
	out := new(WalletKeyManagerInvalidNonce)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidNonce", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidPublicKey represents a InvalidPublicKey error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidPublicKey struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPublicKey()
func WalletKeyManagerInvalidPublicKeyErrorID() common.Hash {
	return common.HexToHash("0xa2d0fee854985cb3d3a394146b5df1c41d5b4d6fcede7fc57a3c05a49b2f3e73")
}

// UnpackInvalidPublicKeyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPublicKey()
func (walletKeyManager *WalletKeyManager) UnpackInvalidPublicKeyError(raw []byte) (*WalletKeyManagerInvalidPublicKey, error) {
	out := new(WalletKeyManagerInvalidPublicKey)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidPublicKey", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidResponseData represents a InvalidResponseData error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidResponseData struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidResponseData()
func WalletKeyManagerInvalidResponseDataErrorID() common.Hash {
	return common.HexToHash("0x48566617128355a033880805d6462da9bf9c6fe4c61b922997451a71262990af")
}

// UnpackInvalidResponseDataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidResponseData()
func (walletKeyManager *WalletKeyManager) UnpackInvalidResponseDataError(raw []byte) (*WalletKeyManagerInvalidResponseData, error) {
	out := new(WalletKeyManagerInvalidResponseData)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidResponseData", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidSettings represents a InvalidSettings error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidSettings struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSettings()
func WalletKeyManagerInvalidSettingsErrorID() common.Hash {
	return common.HexToHash("0xe591f33dff1fb5fca6dd4db38478380593c3e82c2df7be391e6fd4373a8daaff")
}

// UnpackInvalidSettingsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSettings()
func (walletKeyManager *WalletKeyManager) UnpackInvalidSettingsError(raw []byte) (*WalletKeyManagerInvalidSettings, error) {
	out := new(WalletKeyManagerInvalidSettings)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidSettings", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidSigningAlgo represents a InvalidSigningAlgo error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidSigningAlgo struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningAlgo()
func WalletKeyManagerInvalidSigningAlgoErrorID() common.Hash {
	return common.HexToHash("0x9b7198a697b61e546ef9ee593a6f77d2bf9f1df9dcbba655ce96d7e2c75346c7")
}

// UnpackInvalidSigningAlgoError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningAlgo()
func (walletKeyManager *WalletKeyManager) UnpackInvalidSigningAlgoError(raw []byte) (*WalletKeyManagerInvalidSigningAlgo, error) {
	out := new(WalletKeyManagerInvalidSigningAlgo)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidSigningAlgo", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidTeeSignature represents a InvalidTeeSignature error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidTeeSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidTeeSignature()
func WalletKeyManagerInvalidTeeSignatureErrorID() common.Hash {
	return common.HexToHash("0x18a4b8130e3478799e6e36e74778074f81f5b869d1246e8aa360267b9fdb79f9")
}

// UnpackInvalidTeeSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidTeeSignature()
func (walletKeyManager *WalletKeyManager) UnpackInvalidTeeSignatureError(raw []byte) (*WalletKeyManagerInvalidTeeSignature, error) {
	out := new(WalletKeyManagerInvalidTeeSignature)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidTeeSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidThreshold represents a InvalidThreshold error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidThreshold struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidThreshold()
func WalletKeyManagerInvalidThresholdErrorID() common.Hash {
	return common.HexToHash("0xaabd5a09991cc5327efcd91d11259139f5e84df65c9346e998897a10eaca0de2")
}

// UnpackInvalidThresholdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidThreshold()
func (walletKeyManager *WalletKeyManager) UnpackInvalidThresholdError(raw []byte) (*WalletKeyManagerInvalidThreshold, error) {
	out := new(WalletKeyManagerInvalidThreshold)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidThreshold", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerInvalidWalletStatus represents a InvalidWalletStatus error raised by the WalletKeyManager contract.
type WalletKeyManagerInvalidWalletStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidWalletStatus()
func WalletKeyManagerInvalidWalletStatusErrorID() common.Hash {
	return common.HexToHash("0xd700120e4896ea7835a5e9e4feb690644ced5cad1c52209afdcbe06b96d91345")
}

// UnpackInvalidWalletStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidWalletStatus()
func (walletKeyManager *WalletKeyManager) UnpackInvalidWalletStatusError(raw []byte) (*WalletKeyManagerInvalidWalletStatus, error) {
	out := new(WalletKeyManagerInvalidWalletStatus)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "InvalidWalletStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerKeyNotGeneratedOnTeeMachine represents a KeyNotGeneratedOnTeeMachine error raised by the WalletKeyManager contract.
type WalletKeyManagerKeyNotGeneratedOnTeeMachine struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyNotGeneratedOnTeeMachine()
func WalletKeyManagerKeyNotGeneratedOnTeeMachineErrorID() common.Hash {
	return common.HexToHash("0x2d224980c5d5f5a7109ad3b3bede33f62cd9224714e4eee2c3b25f8712aa5e57")
}

// UnpackKeyNotGeneratedOnTeeMachineError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyNotGeneratedOnTeeMachine()
func (walletKeyManager *WalletKeyManager) UnpackKeyNotGeneratedOnTeeMachineError(raw []byte) (*WalletKeyManagerKeyNotGeneratedOnTeeMachine, error) {
	out := new(WalletKeyManagerKeyNotGeneratedOnTeeMachine)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "KeyNotGeneratedOnTeeMachine", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerKeyNotRestoredOnTeeMachine represents a KeyNotRestoredOnTeeMachine error raised by the WalletKeyManager contract.
type WalletKeyManagerKeyNotRestoredOnTeeMachine struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyNotRestoredOnTeeMachine()
func WalletKeyManagerKeyNotRestoredOnTeeMachineErrorID() common.Hash {
	return common.HexToHash("0xb1fc4a4ef50b79383a556d68ad8e282c5c15fc8a734ca298570d20a18d055123")
}

// UnpackKeyNotRestoredOnTeeMachineError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyNotRestoredOnTeeMachine()
func (walletKeyManager *WalletKeyManager) UnpackKeyNotRestoredOnTeeMachineError(raw []byte) (*WalletKeyManagerKeyNotRestoredOnTeeMachine, error) {
	out := new(WalletKeyManagerKeyNotRestoredOnTeeMachine)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "KeyNotRestoredOnTeeMachine", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerKeyTypeNotSupported represents a KeyTypeNotSupported error raised by the WalletKeyManager contract.
type WalletKeyManagerKeyTypeNotSupported struct {
	KeyType [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func WalletKeyManagerKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x85004df1f47814c2de6fbf03d3f2b1ca8c67c182d0d54278766543b6be3bc603")
}

// UnpackKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error KeyTypeNotSupported(bytes32 keyType)
func (walletKeyManager *WalletKeyManager) UnpackKeyTypeNotSupportedError(raw []byte) (*WalletKeyManagerKeyTypeNotSupported, error) {
	out := new(WalletKeyManagerKeyTypeNotSupported)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "KeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerLengthsMismatch represents a LengthsMismatch error raised by the WalletKeyManager contract.
type WalletKeyManagerLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func WalletKeyManagerLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (walletKeyManager *WalletKeyManager) UnpackLengthsMismatchError(raw []byte) (*WalletKeyManagerLengthsMismatch, error) {
	out := new(WalletKeyManagerLengthsMismatch)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerMessageEmpty represents a MessageEmpty error raised by the WalletKeyManager contract.
type WalletKeyManagerMessageEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MessageEmpty()
func WalletKeyManagerMessageEmptyErrorID() common.Hash {
	return common.HexToHash("0x3b06a1beedd2da56654bd6f78f50f74e9bc7469bae35c52fd3680e4d95103c79")
}

// UnpackMessageEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MessageEmpty()
func (walletKeyManager *WalletKeyManager) UnpackMessageEmptyError(raw []byte) (*WalletKeyManagerMessageEmpty, error) {
	out := new(WalletKeyManagerMessageEmpty)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "MessageEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerNoAddresses represents a NoAddresses error raised by the WalletKeyManager contract.
type WalletKeyManagerNoAddresses struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoAddresses()
func WalletKeyManagerNoAddressesErrorID() common.Hash {
	return common.HexToHash("0xb2e79f05746fa6c9f3d61f9efe85959371cde5030a6c2dd70c92b63578c5cde7")
}

// UnpackNoAddressesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoAddresses()
func (walletKeyManager *WalletKeyManager) UnpackNoAddressesError(raw []byte) (*WalletKeyManagerNoAddresses, error) {
	out := new(WalletKeyManagerNoAddresses)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "NoAddresses", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerNoTeeMachinesSpecified represents a NoTeeMachinesSpecified error raised by the WalletKeyManager contract.
type WalletKeyManagerNoTeeMachinesSpecified struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoTeeMachinesSpecified()
func WalletKeyManagerNoTeeMachinesSpecifiedErrorID() common.Hash {
	return common.HexToHash("0xe10387d0118f08d51d80fe03b0c14ed458c46aac2ae9964e0b628978e6938db8")
}

// UnpackNoTeeMachinesSpecifiedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoTeeMachinesSpecified()
func (walletKeyManager *WalletKeyManager) UnpackNoTeeMachinesSpecifiedError(raw []byte) (*WalletKeyManagerNoTeeMachinesSpecified, error) {
	out := new(WalletKeyManagerNoTeeMachinesSpecified)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "NoTeeMachinesSpecified", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerNotOwnerOrPauser represents a NotOwnerOrPauser error raised by the WalletKeyManager contract.
type WalletKeyManagerNotOwnerOrPauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrPauser(address caller)
func WalletKeyManagerNotOwnerOrPauserErrorID() common.Hash {
	return common.HexToHash("0x46c0d9af6b895093e022ad2984c54e60192e6e26c5063daf232e825c6e07ff9a")
}

// UnpackNotOwnerOrPauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrPauser(address caller)
func (walletKeyManager *WalletKeyManager) UnpackNotOwnerOrPauserError(raw []byte) (*WalletKeyManagerNotOwnerOrPauser, error) {
	out := new(WalletKeyManagerNotOwnerOrPauser)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "NotOwnerOrPauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerNotOwnerOrUnpauser represents a NotOwnerOrUnpauser error raised by the WalletKeyManager contract.
type WalletKeyManagerNotOwnerOrUnpauser struct {
	Caller common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func WalletKeyManagerNotOwnerOrUnpauserErrorID() common.Hash {
	return common.HexToHash("0x81e700e8f0bd5ba4a6d6a2ce9a1ad0ac9bb1b1b1e62d921a55cbc4b53f1da649")
}

// UnpackNotOwnerOrUnpauserError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotOwnerOrUnpauser(address caller)
func (walletKeyManager *WalletKeyManager) UnpackNotOwnerOrUnpauserError(raw []byte) (*WalletKeyManagerNotOwnerOrUnpauser, error) {
	out := new(WalletKeyManagerNotOwnerOrUnpauser)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "NotOwnerOrUnpauser", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerOnlyExtensionOwner represents a OnlyExtensionOwner error raised by the WalletKeyManager contract.
type WalletKeyManagerOnlyExtensionOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwner()
func WalletKeyManagerOnlyExtensionOwnerErrorID() common.Hash {
	return common.HexToHash("0xcd49fd1de5a297cd1ce055d5abffcf2cb25ad9492bd39a703a238af57b2a9488")
}

// UnpackOnlyExtensionOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwner()
func (walletKeyManager *WalletKeyManager) UnpackOnlyExtensionOwnerError(raw []byte) (*WalletKeyManagerOnlyExtensionOwner, error) {
	out := new(WalletKeyManagerOnlyExtensionOwner)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "OnlyExtensionOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerOnlyExtensionOwnerOrOperator represents a OnlyExtensionOwnerOrOperator error raised by the WalletKeyManager contract.
type WalletKeyManagerOnlyExtensionOwnerOrOperator struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func WalletKeyManagerOnlyExtensionOwnerOrOperatorErrorID() common.Hash {
	return common.HexToHash("0x5602e102e7442057f7a2b9c4093e3c0449bde4a79ece18c8f356dad6c3b40962")
}

// UnpackOnlyExtensionOwnerOrOperatorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExtensionOwnerOrOperator()
func (walletKeyManager *WalletKeyManager) UnpackOnlyExtensionOwnerOrOperatorError(raw []byte) (*WalletKeyManagerOnlyExtensionOwnerOrOperator, error) {
	out := new(WalletKeyManagerOnlyExtensionOwnerOrOperator)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "OnlyExtensionOwnerOrOperator", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerOnlyOwner represents a OnlyOwner error raised by the WalletKeyManager contract.
type WalletKeyManagerOnlyOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwner()
func WalletKeyManagerOnlyOwnerErrorID() common.Hash {
	return common.HexToHash("0x5fc483c5fcd3811670a93392f365a32f0fc9e73013bbed53888646cfd592055f")
}

// UnpackOnlyOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwner()
func (walletKeyManager *WalletKeyManager) UnpackOnlyOwnerError(raw []byte) (*WalletKeyManagerOnlyOwner, error) {
	out := new(WalletKeyManagerOnlyOwner)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "OnlyOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerOnlyOwnerOrBackupManager represents a OnlyOwnerOrBackupManager error raised by the WalletKeyManager contract.
type WalletKeyManagerOnlyOwnerOrBackupManager struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyOwnerOrBackupManager()
func WalletKeyManagerOnlyOwnerOrBackupManagerErrorID() common.Hash {
	return common.HexToHash("0x249a6cf6e22092f9c80798e2a0c583c0f36aea0779ef38a5b5cdb6a112a0e6df")
}

// UnpackOnlyOwnerOrBackupManagerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyOwnerOrBackupManager()
func (walletKeyManager *WalletKeyManager) UnpackOnlyOwnerOrBackupManagerError(raw []byte) (*WalletKeyManagerOnlyOwnerOrBackupManager, error) {
	out := new(WalletKeyManagerOnlyOwnerOrBackupManager)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "OnlyOwnerOrBackupManager", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the WalletKeyManager contract.
type WalletKeyManagerOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func WalletKeyManagerOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (walletKeyManager *WalletKeyManager) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*WalletKeyManagerOnlyProductionOrPausedStatus, error) {
	out := new(WalletKeyManagerOnlyProductionOrPausedStatus)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerOnlyProposedOwner represents a OnlyProposedOwner error raised by the WalletKeyManager contract.
type WalletKeyManagerOnlyProposedOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProposedOwner()
func WalletKeyManagerOnlyProposedOwnerErrorID() common.Hash {
	return common.HexToHash("0x49cb15061ee558e74f66f937f740fd2a3791a89d2d139862222e5c97cabc02cd")
}

// UnpackOnlyProposedOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProposedOwner()
func (walletKeyManager *WalletKeyManager) UnpackOnlyProposedOwnerError(raw []byte) (*WalletKeyManagerOnlyProposedOwner, error) {
	out := new(WalletKeyManagerOnlyProposedOwner)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "OnlyProposedOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerOperationCommandEmpty represents a OperationCommandEmpty error raised by the WalletKeyManager contract.
type WalletKeyManagerOperationCommandEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationCommandEmpty()
func WalletKeyManagerOperationCommandEmptyErrorID() common.Hash {
	return common.HexToHash("0x038dc5a8067401ab291233391b3d11ff3ce1b21a8aebcf4297e6a9f759f65846")
}

// UnpackOperationCommandEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationCommandEmpty()
func (walletKeyManager *WalletKeyManager) UnpackOperationCommandEmptyError(raw []byte) (*WalletKeyManagerOperationCommandEmpty, error) {
	out := new(WalletKeyManagerOperationCommandEmpty)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "OperationCommandEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerOperationTypeEmpty represents a OperationTypeEmpty error raised by the WalletKeyManager contract.
type WalletKeyManagerOperationTypeEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationTypeEmpty()
func WalletKeyManagerOperationTypeEmptyErrorID() common.Hash {
	return common.HexToHash("0x06e3960d8a83ea8972d826357113f286ce60947bac56680f119fca98a2d6125b")
}

// UnpackOperationTypeEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationTypeEmpty()
func (walletKeyManager *WalletKeyManager) UnpackOperationTypeEmptyError(raw []byte) (*WalletKeyManagerOperationTypeEmpty, error) {
	out := new(WalletKeyManagerOperationTypeEmpty)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "OperationTypeEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerOwnerNotAllowed represents a OwnerNotAllowed error raised by the WalletKeyManager contract.
type WalletKeyManagerOwnerNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnerNotAllowed()
func WalletKeyManagerOwnerNotAllowedErrorID() common.Hash {
	return common.HexToHash("0x3d5c9996545c6bf4287c26c1c2d8e6b4c39b92f19ba706c3a756720bf9c4fd24")
}

// UnpackOwnerNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnerNotAllowed()
func (walletKeyManager *WalletKeyManager) UnpackOwnerNotAllowedError(raw []byte) (*WalletKeyManagerOwnerNotAllowed, error) {
	out := new(WalletKeyManagerOwnerNotAllowed)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "OwnerNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerTeeIdAlreadyAdded represents a TeeIdAlreadyAdded error raised by the WalletKeyManager contract.
type WalletKeyManagerTeeIdAlreadyAdded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeIdAlreadyAdded()
func WalletKeyManagerTeeIdAlreadyAddedErrorID() common.Hash {
	return common.HexToHash("0xcaf75a7d947bb7b6b277feafd2b879602a107db503e54ceb5b454275c9813e32")
}

// UnpackTeeIdAlreadyAddedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeIdAlreadyAdded()
func (walletKeyManager *WalletKeyManager) UnpackTeeIdAlreadyAddedError(raw []byte) (*WalletKeyManagerTeeIdAlreadyAdded, error) {
	out := new(WalletKeyManagerTeeIdAlreadyAdded)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "TeeIdAlreadyAdded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerTeeMachineNotAvailable represents a TeeMachineNotAvailable error raised by the WalletKeyManager contract.
type WalletKeyManagerTeeMachineNotAvailable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeMachineNotAvailable()
func WalletKeyManagerTeeMachineNotAvailableErrorID() common.Hash {
	return common.HexToHash("0xea552ba8d99c4bf2c17ab1bc9b879801707acbec39778f2d909f45b425ec0dea")
}

// UnpackTeeMachineNotAvailableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeMachineNotAvailable()
func (walletKeyManager *WalletKeyManager) UnpackTeeMachineNotAvailableError(raw []byte) (*WalletKeyManagerTeeMachineNotAvailable, error) {
	out := new(WalletKeyManagerTeeMachineNotAvailable)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "TeeMachineNotAvailable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerTeeNotFound represents a TeeNotFound error raised by the WalletKeyManager contract.
type WalletKeyManagerTeeNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TeeNotFound()
func WalletKeyManagerTeeNotFoundErrorID() common.Hash {
	return common.HexToHash("0xceb05b6852ed4c7dc3a718eb0d26555398a76877a44dd41dc92c5fe7e6ef8c2d")
}

// UnpackTeeNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TeeNotFound()
func (walletKeyManager *WalletKeyManager) UnpackTeeNotFoundError(raw []byte) (*WalletKeyManagerTeeNotFound, error) {
	out := new(WalletKeyManagerTeeNotFound)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "TeeNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerThresholdNotMet represents a ThresholdNotMet error raised by the WalletKeyManager contract.
type WalletKeyManagerThresholdNotMet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ThresholdNotMet()
func WalletKeyManagerThresholdNotMetErrorID() common.Hash {
	return common.HexToHash("0x59fa4a935269513c1e751ea6c6fe2e0a1920bf90b5e062c629e36320baf7378d")
}

// UnpackThresholdNotMetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ThresholdNotMet()
func (walletKeyManager *WalletKeyManager) UnpackThresholdNotMetError(raw []byte) (*WalletKeyManagerThresholdNotMet, error) {
	out := new(WalletKeyManagerThresholdNotMet)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "ThresholdNotMet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKeyManagerVersionNotSupported represents a VersionNotSupported error raised by the WalletKeyManager contract.
type WalletKeyManagerVersionNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error VersionNotSupported()
func WalletKeyManagerVersionNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x6db0b17b06266b37e721024a8ba1221439799ebb779dfd397865e37250abc9a8")
}

// UnpackVersionNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error VersionNotSupported()
func (walletKeyManager *WalletKeyManager) UnpackVersionNotSupportedError(raw []byte) (*WalletKeyManagerVersionNotSupported, error) {
	out := new(WalletKeyManagerVersionNotSupported)
	if err := walletKeyManager.abi.UnpackIntoInterface(out, "VersionNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}
