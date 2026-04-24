// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package extensionmanager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ExtensionManagerMetaData contains all meta data concerning the ExtensionManager contract.
var ExtensionManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CodeHashPlatformAlreadyDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CodeHashZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCodeHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInstructionsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPlatform\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyTypeEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPlatforms\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"NoSigningAlgos\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"PlatformAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PlatformEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"}],\"name\":\"SigningAlgoAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SigningAlgoEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SystemOwnedExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"UnsupportedPlatform\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"platform\",\"type\":\"bytes32\"}],\"name\":\"CodeHashPlatformDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwnerProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTypes\",\"type\":\"bytes32[]\"}],\"name\":\"SupportedKeyTypesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTypes\",\"type\":\"bytes32[]\"}],\"name\":\"SupportedKeyTypesRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"keyTypes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32[][]\",\"name\":\"_signingAlgosByKeyType\",\"type\":\"bytes32[][]\"}],\"name\":\"SystemSupportedKeyTypesAndSigningAlgosAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"platforms\",\"type\":\"bytes32[]\"}],\"name\":\"SystemSupportedPlatformsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"teeExtensionStateVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"TeeExtensionContractsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"TeeExtensionRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"codeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"platforms\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"governanceHash\",\"type\":\"bytes32\"}],\"name\":\"TeeVersionAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_keyTypes\",\"type\":\"bytes32[]\"}],\"name\":\"addSupportedKeyTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_keyTypes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"_signingAlgosByKeyType\",\"type\":\"bytes32[][]\"}],\"name\":\"addSystemSupportedKeyTypesAndSigningAlgos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"}],\"name\":\"addSystemSupportedPlatforms\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"}],\"name\":\"addTeeVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"confirmOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"disableCodeHashPlatform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extensionsCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"}],\"name\":\"getCodeHashInfo\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_governanceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"_platforms\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getExtensionOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getSupportedCodeHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_supportedCodeHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getSupportedKeyTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_supportedKeyTypes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSystemSupportedKeyTypes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSystemSupportedPlatforms\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"}],\"name\":\"getSystemSupportedSigningAlgos\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getTeeExtensionInstructionsSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"name\":\"getTeeExtensionStateVerifier\",\"outputs\":[{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"}],\"name\":\"getTeeGovernanceHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"isCodeHashPlatformDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_codeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_platform\",\"type\":\"bytes32\"}],\"name\":\"isCodeHashPlatformSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"}],\"name\":\"isKeyTypeSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_signingAlgo\",\"type\":\"bytes32\"}],\"name\":\"isSigningAlgoSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"proposeNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"_teeExtensionStateVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_keyTypes\",\"type\":\"bytes32[]\"}],\"name\":\"removeSupportedKeyTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_extensionId\",\"type\":\"uint256\"},{\"internalType\":\"contractITeeExtensionStateVerifier\",\"name\":\"_teeExtensionStateVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_teeExtensionInstructionsSender\",\"type\":\"address\"}],\"name\":\"setExtensionContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ExtensionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ExtensionManagerMetaData.ABI instead.
var ExtensionManagerABI = ExtensionManagerMetaData.ABI

// ExtensionManager is an auto generated Go binding around an Ethereum contract.
type ExtensionManager struct {
	ExtensionManagerCaller     // Read-only binding to the contract
	ExtensionManagerTransactor // Write-only binding to the contract
	ExtensionManagerFilterer   // Log filterer for contract events
}

// ExtensionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtensionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtensionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtensionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtensionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtensionManagerSession struct {
	Contract     *ExtensionManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExtensionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtensionManagerCallerSession struct {
	Contract *ExtensionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ExtensionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtensionManagerTransactorSession struct {
	Contract     *ExtensionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExtensionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtensionManagerRaw struct {
	Contract *ExtensionManager // Generic contract binding to access the raw methods on
}

// ExtensionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtensionManagerCallerRaw struct {
	Contract *ExtensionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ExtensionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtensionManagerTransactorRaw struct {
	Contract *ExtensionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtensionManager creates a new instance of ExtensionManager, bound to a specific deployed contract.
func NewExtensionManager(address common.Address, backend bind.ContractBackend) (*ExtensionManager, error) {
	contract, err := bindExtensionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExtensionManager{ExtensionManagerCaller: ExtensionManagerCaller{contract: contract}, ExtensionManagerTransactor: ExtensionManagerTransactor{contract: contract}, ExtensionManagerFilterer: ExtensionManagerFilterer{contract: contract}}, nil
}

// NewExtensionManagerCaller creates a new read-only instance of ExtensionManager, bound to a specific deployed contract.
func NewExtensionManagerCaller(address common.Address, caller bind.ContractCaller) (*ExtensionManagerCaller, error) {
	contract, err := bindExtensionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerCaller{contract: contract}, nil
}

// NewExtensionManagerTransactor creates a new write-only instance of ExtensionManager, bound to a specific deployed contract.
func NewExtensionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtensionManagerTransactor, error) {
	contract, err := bindExtensionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerTransactor{contract: contract}, nil
}

// NewExtensionManagerFilterer creates a new log filterer instance of ExtensionManager, bound to a specific deployed contract.
func NewExtensionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtensionManagerFilterer, error) {
	contract, err := bindExtensionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerFilterer{contract: contract}, nil
}

// bindExtensionManager binds a generic wrapper to an already deployed contract.
func bindExtensionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExtensionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtensionManager *ExtensionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtensionManager.Contract.ExtensionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtensionManager *ExtensionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtensionManager.Contract.ExtensionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtensionManager *ExtensionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtensionManager.Contract.ExtensionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtensionManager *ExtensionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtensionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtensionManager *ExtensionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtensionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtensionManager *ExtensionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtensionManager.Contract.contract.Transact(opts, method, params...)
}

// ExtensionsCounter is a free data retrieval call binding the contract method 0xfad5902b.
//
// Solidity: function extensionsCounter() view returns(uint256)
func (_ExtensionManager *ExtensionManagerCaller) ExtensionsCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "extensionsCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtensionsCounter is a free data retrieval call binding the contract method 0xfad5902b.
//
// Solidity: function extensionsCounter() view returns(uint256)
func (_ExtensionManager *ExtensionManagerSession) ExtensionsCounter() (*big.Int, error) {
	return _ExtensionManager.Contract.ExtensionsCounter(&_ExtensionManager.CallOpts)
}

// ExtensionsCounter is a free data retrieval call binding the contract method 0xfad5902b.
//
// Solidity: function extensionsCounter() view returns(uint256)
func (_ExtensionManager *ExtensionManagerCallerSession) ExtensionsCounter() (*big.Int, error) {
	return _ExtensionManager.Contract.ExtensionsCounter(&_ExtensionManager.CallOpts)
}

// GetCodeHashInfo is a free data retrieval call binding the contract method 0x62672305.
//
// Solidity: function getCodeHashInfo(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32 _governanceHash, string _version, bytes32[] _platforms)
func (_ExtensionManager *ExtensionManagerCaller) GetCodeHashInfo(opts *bind.CallOpts, _extensionId *big.Int, _codeHash [32]byte) (struct {
	GovernanceHash [32]byte
	Version        string
	Platforms      [][32]byte
}, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "getCodeHashInfo", _extensionId, _codeHash)

	outstruct := new(struct {
		GovernanceHash [32]byte
		Version        string
		Platforms      [][32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GovernanceHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Version = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Platforms = *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return *outstruct, err

}

// GetCodeHashInfo is a free data retrieval call binding the contract method 0x62672305.
//
// Solidity: function getCodeHashInfo(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32 _governanceHash, string _version, bytes32[] _platforms)
func (_ExtensionManager *ExtensionManagerSession) GetCodeHashInfo(_extensionId *big.Int, _codeHash [32]byte) (struct {
	GovernanceHash [32]byte
	Version        string
	Platforms      [][32]byte
}, error) {
	return _ExtensionManager.Contract.GetCodeHashInfo(&_ExtensionManager.CallOpts, _extensionId, _codeHash)
}

// GetCodeHashInfo is a free data retrieval call binding the contract method 0x62672305.
//
// Solidity: function getCodeHashInfo(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32 _governanceHash, string _version, bytes32[] _platforms)
func (_ExtensionManager *ExtensionManagerCallerSession) GetCodeHashInfo(_extensionId *big.Int, _codeHash [32]byte) (struct {
	GovernanceHash [32]byte
	Version        string
	Platforms      [][32]byte
}, error) {
	return _ExtensionManager.Contract.GetCodeHashInfo(&_ExtensionManager.CallOpts, _extensionId, _codeHash)
}

// GetExtensionOwner is a free data retrieval call binding the contract method 0x5e46e380.
//
// Solidity: function getExtensionOwner(uint256 _extensionId) view returns(address)
func (_ExtensionManager *ExtensionManagerCaller) GetExtensionOwner(opts *bind.CallOpts, _extensionId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "getExtensionOwner", _extensionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExtensionOwner is a free data retrieval call binding the contract method 0x5e46e380.
//
// Solidity: function getExtensionOwner(uint256 _extensionId) view returns(address)
func (_ExtensionManager *ExtensionManagerSession) GetExtensionOwner(_extensionId *big.Int) (common.Address, error) {
	return _ExtensionManager.Contract.GetExtensionOwner(&_ExtensionManager.CallOpts, _extensionId)
}

// GetExtensionOwner is a free data retrieval call binding the contract method 0x5e46e380.
//
// Solidity: function getExtensionOwner(uint256 _extensionId) view returns(address)
func (_ExtensionManager *ExtensionManagerCallerSession) GetExtensionOwner(_extensionId *big.Int) (common.Address, error) {
	return _ExtensionManager.Contract.GetExtensionOwner(&_ExtensionManager.CallOpts, _extensionId)
}

// GetSupportedCodeHashes is a free data retrieval call binding the contract method 0x9ebb72b2.
//
// Solidity: function getSupportedCodeHashes(uint256 _extensionId) view returns(bytes32[] _supportedCodeHashes)
func (_ExtensionManager *ExtensionManagerCaller) GetSupportedCodeHashes(opts *bind.CallOpts, _extensionId *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "getSupportedCodeHashes", _extensionId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSupportedCodeHashes is a free data retrieval call binding the contract method 0x9ebb72b2.
//
// Solidity: function getSupportedCodeHashes(uint256 _extensionId) view returns(bytes32[] _supportedCodeHashes)
func (_ExtensionManager *ExtensionManagerSession) GetSupportedCodeHashes(_extensionId *big.Int) ([][32]byte, error) {
	return _ExtensionManager.Contract.GetSupportedCodeHashes(&_ExtensionManager.CallOpts, _extensionId)
}

// GetSupportedCodeHashes is a free data retrieval call binding the contract method 0x9ebb72b2.
//
// Solidity: function getSupportedCodeHashes(uint256 _extensionId) view returns(bytes32[] _supportedCodeHashes)
func (_ExtensionManager *ExtensionManagerCallerSession) GetSupportedCodeHashes(_extensionId *big.Int) ([][32]byte, error) {
	return _ExtensionManager.Contract.GetSupportedCodeHashes(&_ExtensionManager.CallOpts, _extensionId)
}

// GetSupportedKeyTypes is a free data retrieval call binding the contract method 0x8bab3177.
//
// Solidity: function getSupportedKeyTypes(uint256 _extensionId) view returns(bytes32[] _supportedKeyTypes)
func (_ExtensionManager *ExtensionManagerCaller) GetSupportedKeyTypes(opts *bind.CallOpts, _extensionId *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "getSupportedKeyTypes", _extensionId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSupportedKeyTypes is a free data retrieval call binding the contract method 0x8bab3177.
//
// Solidity: function getSupportedKeyTypes(uint256 _extensionId) view returns(bytes32[] _supportedKeyTypes)
func (_ExtensionManager *ExtensionManagerSession) GetSupportedKeyTypes(_extensionId *big.Int) ([][32]byte, error) {
	return _ExtensionManager.Contract.GetSupportedKeyTypes(&_ExtensionManager.CallOpts, _extensionId)
}

// GetSupportedKeyTypes is a free data retrieval call binding the contract method 0x8bab3177.
//
// Solidity: function getSupportedKeyTypes(uint256 _extensionId) view returns(bytes32[] _supportedKeyTypes)
func (_ExtensionManager *ExtensionManagerCallerSession) GetSupportedKeyTypes(_extensionId *big.Int) ([][32]byte, error) {
	return _ExtensionManager.Contract.GetSupportedKeyTypes(&_ExtensionManager.CallOpts, _extensionId)
}

// GetSystemSupportedKeyTypes is a free data retrieval call binding the contract method 0x50a15fcc.
//
// Solidity: function getSystemSupportedKeyTypes() view returns(bytes32[])
func (_ExtensionManager *ExtensionManagerCaller) GetSystemSupportedKeyTypes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "getSystemSupportedKeyTypes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSystemSupportedKeyTypes is a free data retrieval call binding the contract method 0x50a15fcc.
//
// Solidity: function getSystemSupportedKeyTypes() view returns(bytes32[])
func (_ExtensionManager *ExtensionManagerSession) GetSystemSupportedKeyTypes() ([][32]byte, error) {
	return _ExtensionManager.Contract.GetSystemSupportedKeyTypes(&_ExtensionManager.CallOpts)
}

// GetSystemSupportedKeyTypes is a free data retrieval call binding the contract method 0x50a15fcc.
//
// Solidity: function getSystemSupportedKeyTypes() view returns(bytes32[])
func (_ExtensionManager *ExtensionManagerCallerSession) GetSystemSupportedKeyTypes() ([][32]byte, error) {
	return _ExtensionManager.Contract.GetSystemSupportedKeyTypes(&_ExtensionManager.CallOpts)
}

// GetSystemSupportedPlatforms is a free data retrieval call binding the contract method 0xe38c8964.
//
// Solidity: function getSystemSupportedPlatforms() view returns(bytes32[])
func (_ExtensionManager *ExtensionManagerCaller) GetSystemSupportedPlatforms(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "getSystemSupportedPlatforms")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSystemSupportedPlatforms is a free data retrieval call binding the contract method 0xe38c8964.
//
// Solidity: function getSystemSupportedPlatforms() view returns(bytes32[])
func (_ExtensionManager *ExtensionManagerSession) GetSystemSupportedPlatforms() ([][32]byte, error) {
	return _ExtensionManager.Contract.GetSystemSupportedPlatforms(&_ExtensionManager.CallOpts)
}

// GetSystemSupportedPlatforms is a free data retrieval call binding the contract method 0xe38c8964.
//
// Solidity: function getSystemSupportedPlatforms() view returns(bytes32[])
func (_ExtensionManager *ExtensionManagerCallerSession) GetSystemSupportedPlatforms() ([][32]byte, error) {
	return _ExtensionManager.Contract.GetSystemSupportedPlatforms(&_ExtensionManager.CallOpts)
}

// GetSystemSupportedSigningAlgos is a free data retrieval call binding the contract method 0xc6cba800.
//
// Solidity: function getSystemSupportedSigningAlgos(bytes32 _keyType) view returns(bytes32[])
func (_ExtensionManager *ExtensionManagerCaller) GetSystemSupportedSigningAlgos(opts *bind.CallOpts, _keyType [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "getSystemSupportedSigningAlgos", _keyType)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSystemSupportedSigningAlgos is a free data retrieval call binding the contract method 0xc6cba800.
//
// Solidity: function getSystemSupportedSigningAlgos(bytes32 _keyType) view returns(bytes32[])
func (_ExtensionManager *ExtensionManagerSession) GetSystemSupportedSigningAlgos(_keyType [32]byte) ([][32]byte, error) {
	return _ExtensionManager.Contract.GetSystemSupportedSigningAlgos(&_ExtensionManager.CallOpts, _keyType)
}

// GetSystemSupportedSigningAlgos is a free data retrieval call binding the contract method 0xc6cba800.
//
// Solidity: function getSystemSupportedSigningAlgos(bytes32 _keyType) view returns(bytes32[])
func (_ExtensionManager *ExtensionManagerCallerSession) GetSystemSupportedSigningAlgos(_keyType [32]byte) ([][32]byte, error) {
	return _ExtensionManager.Contract.GetSystemSupportedSigningAlgos(&_ExtensionManager.CallOpts, _keyType)
}

// GetTeeExtensionInstructionsSender is a free data retrieval call binding the contract method 0x2c177358.
//
// Solidity: function getTeeExtensionInstructionsSender(uint256 _extensionId) view returns(address)
func (_ExtensionManager *ExtensionManagerCaller) GetTeeExtensionInstructionsSender(opts *bind.CallOpts, _extensionId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "getTeeExtensionInstructionsSender", _extensionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeeExtensionInstructionsSender is a free data retrieval call binding the contract method 0x2c177358.
//
// Solidity: function getTeeExtensionInstructionsSender(uint256 _extensionId) view returns(address)
func (_ExtensionManager *ExtensionManagerSession) GetTeeExtensionInstructionsSender(_extensionId *big.Int) (common.Address, error) {
	return _ExtensionManager.Contract.GetTeeExtensionInstructionsSender(&_ExtensionManager.CallOpts, _extensionId)
}

// GetTeeExtensionInstructionsSender is a free data retrieval call binding the contract method 0x2c177358.
//
// Solidity: function getTeeExtensionInstructionsSender(uint256 _extensionId) view returns(address)
func (_ExtensionManager *ExtensionManagerCallerSession) GetTeeExtensionInstructionsSender(_extensionId *big.Int) (common.Address, error) {
	return _ExtensionManager.Contract.GetTeeExtensionInstructionsSender(&_ExtensionManager.CallOpts, _extensionId)
}

// GetTeeExtensionStateVerifier is a free data retrieval call binding the contract method 0x709ebdcd.
//
// Solidity: function getTeeExtensionStateVerifier(uint256 _extensionId) view returns(address)
func (_ExtensionManager *ExtensionManagerCaller) GetTeeExtensionStateVerifier(opts *bind.CallOpts, _extensionId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "getTeeExtensionStateVerifier", _extensionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTeeExtensionStateVerifier is a free data retrieval call binding the contract method 0x709ebdcd.
//
// Solidity: function getTeeExtensionStateVerifier(uint256 _extensionId) view returns(address)
func (_ExtensionManager *ExtensionManagerSession) GetTeeExtensionStateVerifier(_extensionId *big.Int) (common.Address, error) {
	return _ExtensionManager.Contract.GetTeeExtensionStateVerifier(&_ExtensionManager.CallOpts, _extensionId)
}

// GetTeeExtensionStateVerifier is a free data retrieval call binding the contract method 0x709ebdcd.
//
// Solidity: function getTeeExtensionStateVerifier(uint256 _extensionId) view returns(address)
func (_ExtensionManager *ExtensionManagerCallerSession) GetTeeExtensionStateVerifier(_extensionId *big.Int) (common.Address, error) {
	return _ExtensionManager.Contract.GetTeeExtensionStateVerifier(&_ExtensionManager.CallOpts, _extensionId)
}

// GetTeeGovernanceHash is a free data retrieval call binding the contract method 0x9f300b9b.
//
// Solidity: function getTeeGovernanceHash(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32)
func (_ExtensionManager *ExtensionManagerCaller) GetTeeGovernanceHash(opts *bind.CallOpts, _extensionId *big.Int, _codeHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "getTeeGovernanceHash", _extensionId, _codeHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTeeGovernanceHash is a free data retrieval call binding the contract method 0x9f300b9b.
//
// Solidity: function getTeeGovernanceHash(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32)
func (_ExtensionManager *ExtensionManagerSession) GetTeeGovernanceHash(_extensionId *big.Int, _codeHash [32]byte) ([32]byte, error) {
	return _ExtensionManager.Contract.GetTeeGovernanceHash(&_ExtensionManager.CallOpts, _extensionId, _codeHash)
}

// GetTeeGovernanceHash is a free data retrieval call binding the contract method 0x9f300b9b.
//
// Solidity: function getTeeGovernanceHash(uint256 _extensionId, bytes32 _codeHash) view returns(bytes32)
func (_ExtensionManager *ExtensionManagerCallerSession) GetTeeGovernanceHash(_extensionId *big.Int, _codeHash [32]byte) ([32]byte, error) {
	return _ExtensionManager.Contract.GetTeeGovernanceHash(&_ExtensionManager.CallOpts, _extensionId, _codeHash)
}

// IsCodeHashPlatformDisabled is a free data retrieval call binding the contract method 0x2a19c98e.
//
// Solidity: function isCodeHashPlatformDisabled(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_ExtensionManager *ExtensionManagerCaller) IsCodeHashPlatformDisabled(opts *bind.CallOpts, _extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "isCodeHashPlatformDisabled", _extensionId, _codeHash, _platform)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCodeHashPlatformDisabled is a free data retrieval call binding the contract method 0x2a19c98e.
//
// Solidity: function isCodeHashPlatformDisabled(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_ExtensionManager *ExtensionManagerSession) IsCodeHashPlatformDisabled(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	return _ExtensionManager.Contract.IsCodeHashPlatformDisabled(&_ExtensionManager.CallOpts, _extensionId, _codeHash, _platform)
}

// IsCodeHashPlatformDisabled is a free data retrieval call binding the contract method 0x2a19c98e.
//
// Solidity: function isCodeHashPlatformDisabled(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_ExtensionManager *ExtensionManagerCallerSession) IsCodeHashPlatformDisabled(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	return _ExtensionManager.Contract.IsCodeHashPlatformDisabled(&_ExtensionManager.CallOpts, _extensionId, _codeHash, _platform)
}

// IsCodeHashPlatformSupported is a free data retrieval call binding the contract method 0xa1d925c6.
//
// Solidity: function isCodeHashPlatformSupported(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_ExtensionManager *ExtensionManagerCaller) IsCodeHashPlatformSupported(opts *bind.CallOpts, _extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "isCodeHashPlatformSupported", _extensionId, _codeHash, _platform)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCodeHashPlatformSupported is a free data retrieval call binding the contract method 0xa1d925c6.
//
// Solidity: function isCodeHashPlatformSupported(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_ExtensionManager *ExtensionManagerSession) IsCodeHashPlatformSupported(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	return _ExtensionManager.Contract.IsCodeHashPlatformSupported(&_ExtensionManager.CallOpts, _extensionId, _codeHash, _platform)
}

// IsCodeHashPlatformSupported is a free data retrieval call binding the contract method 0xa1d925c6.
//
// Solidity: function isCodeHashPlatformSupported(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) view returns(bool)
func (_ExtensionManager *ExtensionManagerCallerSession) IsCodeHashPlatformSupported(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (bool, error) {
	return _ExtensionManager.Contract.IsCodeHashPlatformSupported(&_ExtensionManager.CallOpts, _extensionId, _codeHash, _platform)
}

// IsKeyTypeSupported is a free data retrieval call binding the contract method 0xf4ac4279.
//
// Solidity: function isKeyTypeSupported(uint256 _extensionId, bytes32 _keyType) view returns(bool)
func (_ExtensionManager *ExtensionManagerCaller) IsKeyTypeSupported(opts *bind.CallOpts, _extensionId *big.Int, _keyType [32]byte) (bool, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "isKeyTypeSupported", _extensionId, _keyType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKeyTypeSupported is a free data retrieval call binding the contract method 0xf4ac4279.
//
// Solidity: function isKeyTypeSupported(uint256 _extensionId, bytes32 _keyType) view returns(bool)
func (_ExtensionManager *ExtensionManagerSession) IsKeyTypeSupported(_extensionId *big.Int, _keyType [32]byte) (bool, error) {
	return _ExtensionManager.Contract.IsKeyTypeSupported(&_ExtensionManager.CallOpts, _extensionId, _keyType)
}

// IsKeyTypeSupported is a free data retrieval call binding the contract method 0xf4ac4279.
//
// Solidity: function isKeyTypeSupported(uint256 _extensionId, bytes32 _keyType) view returns(bool)
func (_ExtensionManager *ExtensionManagerCallerSession) IsKeyTypeSupported(_extensionId *big.Int, _keyType [32]byte) (bool, error) {
	return _ExtensionManager.Contract.IsKeyTypeSupported(&_ExtensionManager.CallOpts, _extensionId, _keyType)
}

// IsSigningAlgoSupported is a free data retrieval call binding the contract method 0x62487d98.
//
// Solidity: function isSigningAlgoSupported(bytes32 _keyType, bytes32 _signingAlgo) view returns(bool)
func (_ExtensionManager *ExtensionManagerCaller) IsSigningAlgoSupported(opts *bind.CallOpts, _keyType [32]byte, _signingAlgo [32]byte) (bool, error) {
	var out []interface{}
	err := _ExtensionManager.contract.Call(opts, &out, "isSigningAlgoSupported", _keyType, _signingAlgo)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigningAlgoSupported is a free data retrieval call binding the contract method 0x62487d98.
//
// Solidity: function isSigningAlgoSupported(bytes32 _keyType, bytes32 _signingAlgo) view returns(bool)
func (_ExtensionManager *ExtensionManagerSession) IsSigningAlgoSupported(_keyType [32]byte, _signingAlgo [32]byte) (bool, error) {
	return _ExtensionManager.Contract.IsSigningAlgoSupported(&_ExtensionManager.CallOpts, _keyType, _signingAlgo)
}

// IsSigningAlgoSupported is a free data retrieval call binding the contract method 0x62487d98.
//
// Solidity: function isSigningAlgoSupported(bytes32 _keyType, bytes32 _signingAlgo) view returns(bool)
func (_ExtensionManager *ExtensionManagerCallerSession) IsSigningAlgoSupported(_keyType [32]byte, _signingAlgo [32]byte) (bool, error) {
	return _ExtensionManager.Contract.IsSigningAlgoSupported(&_ExtensionManager.CallOpts, _keyType, _signingAlgo)
}

// AddSupportedKeyTypes is a paid mutator transaction binding the contract method 0x174c5e26.
//
// Solidity: function addSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (_ExtensionManager *ExtensionManagerTransactor) AddSupportedKeyTypes(opts *bind.TransactOpts, _extensionId *big.Int, _keyTypes [][32]byte) (*types.Transaction, error) {
	return _ExtensionManager.contract.Transact(opts, "addSupportedKeyTypes", _extensionId, _keyTypes)
}

// AddSupportedKeyTypes is a paid mutator transaction binding the contract method 0x174c5e26.
//
// Solidity: function addSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (_ExtensionManager *ExtensionManagerSession) AddSupportedKeyTypes(_extensionId *big.Int, _keyTypes [][32]byte) (*types.Transaction, error) {
	return _ExtensionManager.Contract.AddSupportedKeyTypes(&_ExtensionManager.TransactOpts, _extensionId, _keyTypes)
}

// AddSupportedKeyTypes is a paid mutator transaction binding the contract method 0x174c5e26.
//
// Solidity: function addSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (_ExtensionManager *ExtensionManagerTransactorSession) AddSupportedKeyTypes(_extensionId *big.Int, _keyTypes [][32]byte) (*types.Transaction, error) {
	return _ExtensionManager.Contract.AddSupportedKeyTypes(&_ExtensionManager.TransactOpts, _extensionId, _keyTypes)
}

// AddSystemSupportedKeyTypesAndSigningAlgos is a paid mutator transaction binding the contract method 0x38c77e46.
//
// Solidity: function addSystemSupportedKeyTypesAndSigningAlgos(bytes32[] _keyTypes, bytes32[][] _signingAlgosByKeyType) returns()
func (_ExtensionManager *ExtensionManagerTransactor) AddSystemSupportedKeyTypesAndSigningAlgos(opts *bind.TransactOpts, _keyTypes [][32]byte, _signingAlgosByKeyType [][][32]byte) (*types.Transaction, error) {
	return _ExtensionManager.contract.Transact(opts, "addSystemSupportedKeyTypesAndSigningAlgos", _keyTypes, _signingAlgosByKeyType)
}

// AddSystemSupportedKeyTypesAndSigningAlgos is a paid mutator transaction binding the contract method 0x38c77e46.
//
// Solidity: function addSystemSupportedKeyTypesAndSigningAlgos(bytes32[] _keyTypes, bytes32[][] _signingAlgosByKeyType) returns()
func (_ExtensionManager *ExtensionManagerSession) AddSystemSupportedKeyTypesAndSigningAlgos(_keyTypes [][32]byte, _signingAlgosByKeyType [][][32]byte) (*types.Transaction, error) {
	return _ExtensionManager.Contract.AddSystemSupportedKeyTypesAndSigningAlgos(&_ExtensionManager.TransactOpts, _keyTypes, _signingAlgosByKeyType)
}

// AddSystemSupportedKeyTypesAndSigningAlgos is a paid mutator transaction binding the contract method 0x38c77e46.
//
// Solidity: function addSystemSupportedKeyTypesAndSigningAlgos(bytes32[] _keyTypes, bytes32[][] _signingAlgosByKeyType) returns()
func (_ExtensionManager *ExtensionManagerTransactorSession) AddSystemSupportedKeyTypesAndSigningAlgos(_keyTypes [][32]byte, _signingAlgosByKeyType [][][32]byte) (*types.Transaction, error) {
	return _ExtensionManager.Contract.AddSystemSupportedKeyTypesAndSigningAlgos(&_ExtensionManager.TransactOpts, _keyTypes, _signingAlgosByKeyType)
}

// AddSystemSupportedPlatforms is a paid mutator transaction binding the contract method 0x0c14a66b.
//
// Solidity: function addSystemSupportedPlatforms(bytes32[] _platforms) returns()
func (_ExtensionManager *ExtensionManagerTransactor) AddSystemSupportedPlatforms(opts *bind.TransactOpts, _platforms [][32]byte) (*types.Transaction, error) {
	return _ExtensionManager.contract.Transact(opts, "addSystemSupportedPlatforms", _platforms)
}

// AddSystemSupportedPlatforms is a paid mutator transaction binding the contract method 0x0c14a66b.
//
// Solidity: function addSystemSupportedPlatforms(bytes32[] _platforms) returns()
func (_ExtensionManager *ExtensionManagerSession) AddSystemSupportedPlatforms(_platforms [][32]byte) (*types.Transaction, error) {
	return _ExtensionManager.Contract.AddSystemSupportedPlatforms(&_ExtensionManager.TransactOpts, _platforms)
}

// AddSystemSupportedPlatforms is a paid mutator transaction binding the contract method 0x0c14a66b.
//
// Solidity: function addSystemSupportedPlatforms(bytes32[] _platforms) returns()
func (_ExtensionManager *ExtensionManagerTransactorSession) AddSystemSupportedPlatforms(_platforms [][32]byte) (*types.Transaction, error) {
	return _ExtensionManager.Contract.AddSystemSupportedPlatforms(&_ExtensionManager.TransactOpts, _platforms)
}

// AddTeeVersion is a paid mutator transaction binding the contract method 0x65b771f7.
//
// Solidity: function addTeeVersion(uint256 _extensionId, string _version, bytes32 _codeHash, bytes32[] _platforms, bytes32 _governanceHash) returns()
func (_ExtensionManager *ExtensionManagerTransactor) AddTeeVersion(opts *bind.TransactOpts, _extensionId *big.Int, _version string, _codeHash [32]byte, _platforms [][32]byte, _governanceHash [32]byte) (*types.Transaction, error) {
	return _ExtensionManager.contract.Transact(opts, "addTeeVersion", _extensionId, _version, _codeHash, _platforms, _governanceHash)
}

// AddTeeVersion is a paid mutator transaction binding the contract method 0x65b771f7.
//
// Solidity: function addTeeVersion(uint256 _extensionId, string _version, bytes32 _codeHash, bytes32[] _platforms, bytes32 _governanceHash) returns()
func (_ExtensionManager *ExtensionManagerSession) AddTeeVersion(_extensionId *big.Int, _version string, _codeHash [32]byte, _platforms [][32]byte, _governanceHash [32]byte) (*types.Transaction, error) {
	return _ExtensionManager.Contract.AddTeeVersion(&_ExtensionManager.TransactOpts, _extensionId, _version, _codeHash, _platforms, _governanceHash)
}

// AddTeeVersion is a paid mutator transaction binding the contract method 0x65b771f7.
//
// Solidity: function addTeeVersion(uint256 _extensionId, string _version, bytes32 _codeHash, bytes32[] _platforms, bytes32 _governanceHash) returns()
func (_ExtensionManager *ExtensionManagerTransactorSession) AddTeeVersion(_extensionId *big.Int, _version string, _codeHash [32]byte, _platforms [][32]byte, _governanceHash [32]byte) (*types.Transaction, error) {
	return _ExtensionManager.Contract.AddTeeVersion(&_ExtensionManager.TransactOpts, _extensionId, _version, _codeHash, _platforms, _governanceHash)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x290fbb03.
//
// Solidity: function confirmOwnership(uint256 _extensionId) returns()
func (_ExtensionManager *ExtensionManagerTransactor) ConfirmOwnership(opts *bind.TransactOpts, _extensionId *big.Int) (*types.Transaction, error) {
	return _ExtensionManager.contract.Transact(opts, "confirmOwnership", _extensionId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x290fbb03.
//
// Solidity: function confirmOwnership(uint256 _extensionId) returns()
func (_ExtensionManager *ExtensionManagerSession) ConfirmOwnership(_extensionId *big.Int) (*types.Transaction, error) {
	return _ExtensionManager.Contract.ConfirmOwnership(&_ExtensionManager.TransactOpts, _extensionId)
}

// ConfirmOwnership is a paid mutator transaction binding the contract method 0x290fbb03.
//
// Solidity: function confirmOwnership(uint256 _extensionId) returns()
func (_ExtensionManager *ExtensionManagerTransactorSession) ConfirmOwnership(_extensionId *big.Int) (*types.Transaction, error) {
	return _ExtensionManager.Contract.ConfirmOwnership(&_ExtensionManager.TransactOpts, _extensionId)
}

// DisableCodeHashPlatform is a paid mutator transaction binding the contract method 0x96f54e59.
//
// Solidity: function disableCodeHashPlatform(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) returns()
func (_ExtensionManager *ExtensionManagerTransactor) DisableCodeHashPlatform(opts *bind.TransactOpts, _extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _ExtensionManager.contract.Transact(opts, "disableCodeHashPlatform", _extensionId, _codeHash, _platform)
}

// DisableCodeHashPlatform is a paid mutator transaction binding the contract method 0x96f54e59.
//
// Solidity: function disableCodeHashPlatform(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) returns()
func (_ExtensionManager *ExtensionManagerSession) DisableCodeHashPlatform(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _ExtensionManager.Contract.DisableCodeHashPlatform(&_ExtensionManager.TransactOpts, _extensionId, _codeHash, _platform)
}

// DisableCodeHashPlatform is a paid mutator transaction binding the contract method 0x96f54e59.
//
// Solidity: function disableCodeHashPlatform(uint256 _extensionId, bytes32 _codeHash, bytes32 _platform) returns()
func (_ExtensionManager *ExtensionManagerTransactorSession) DisableCodeHashPlatform(_extensionId *big.Int, _codeHash [32]byte, _platform [32]byte) (*types.Transaction, error) {
	return _ExtensionManager.Contract.DisableCodeHashPlatform(&_ExtensionManager.TransactOpts, _extensionId, _codeHash, _platform)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0xe947a09f.
//
// Solidity: function proposeNewOwner(uint256 _extensionId, address _newOwner) returns()
func (_ExtensionManager *ExtensionManagerTransactor) ProposeNewOwner(opts *bind.TransactOpts, _extensionId *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _ExtensionManager.contract.Transact(opts, "proposeNewOwner", _extensionId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0xe947a09f.
//
// Solidity: function proposeNewOwner(uint256 _extensionId, address _newOwner) returns()
func (_ExtensionManager *ExtensionManagerSession) ProposeNewOwner(_extensionId *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _ExtensionManager.Contract.ProposeNewOwner(&_ExtensionManager.TransactOpts, _extensionId, _newOwner)
}

// ProposeNewOwner is a paid mutator transaction binding the contract method 0xe947a09f.
//
// Solidity: function proposeNewOwner(uint256 _extensionId, address _newOwner) returns()
func (_ExtensionManager *ExtensionManagerTransactorSession) ProposeNewOwner(_extensionId *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _ExtensionManager.Contract.ProposeNewOwner(&_ExtensionManager.TransactOpts, _extensionId, _newOwner)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns(uint256 _extensionId)
func (_ExtensionManager *ExtensionManagerTransactor) Register(opts *bind.TransactOpts, _teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _ExtensionManager.contract.Transact(opts, "register", _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns(uint256 _extensionId)
func (_ExtensionManager *ExtensionManagerSession) Register(_teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _ExtensionManager.Contract.Register(&_ExtensionManager.TransactOpts, _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// Register is a paid mutator transaction binding the contract method 0xaa677354.
//
// Solidity: function register(address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns(uint256 _extensionId)
func (_ExtensionManager *ExtensionManagerTransactorSession) Register(_teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _ExtensionManager.Contract.Register(&_ExtensionManager.TransactOpts, _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// RemoveSupportedKeyTypes is a paid mutator transaction binding the contract method 0xecb0233f.
//
// Solidity: function removeSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (_ExtensionManager *ExtensionManagerTransactor) RemoveSupportedKeyTypes(opts *bind.TransactOpts, _extensionId *big.Int, _keyTypes [][32]byte) (*types.Transaction, error) {
	return _ExtensionManager.contract.Transact(opts, "removeSupportedKeyTypes", _extensionId, _keyTypes)
}

// RemoveSupportedKeyTypes is a paid mutator transaction binding the contract method 0xecb0233f.
//
// Solidity: function removeSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (_ExtensionManager *ExtensionManagerSession) RemoveSupportedKeyTypes(_extensionId *big.Int, _keyTypes [][32]byte) (*types.Transaction, error) {
	return _ExtensionManager.Contract.RemoveSupportedKeyTypes(&_ExtensionManager.TransactOpts, _extensionId, _keyTypes)
}

// RemoveSupportedKeyTypes is a paid mutator transaction binding the contract method 0xecb0233f.
//
// Solidity: function removeSupportedKeyTypes(uint256 _extensionId, bytes32[] _keyTypes) returns()
func (_ExtensionManager *ExtensionManagerTransactorSession) RemoveSupportedKeyTypes(_extensionId *big.Int, _keyTypes [][32]byte) (*types.Transaction, error) {
	return _ExtensionManager.Contract.RemoveSupportedKeyTypes(&_ExtensionManager.TransactOpts, _extensionId, _keyTypes)
}

// SetExtensionContracts is a paid mutator transaction binding the contract method 0x6df0108f.
//
// Solidity: function setExtensionContracts(uint256 _extensionId, address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns()
func (_ExtensionManager *ExtensionManagerTransactor) SetExtensionContracts(opts *bind.TransactOpts, _extensionId *big.Int, _teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _ExtensionManager.contract.Transact(opts, "setExtensionContracts", _extensionId, _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// SetExtensionContracts is a paid mutator transaction binding the contract method 0x6df0108f.
//
// Solidity: function setExtensionContracts(uint256 _extensionId, address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns()
func (_ExtensionManager *ExtensionManagerSession) SetExtensionContracts(_extensionId *big.Int, _teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _ExtensionManager.Contract.SetExtensionContracts(&_ExtensionManager.TransactOpts, _extensionId, _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// SetExtensionContracts is a paid mutator transaction binding the contract method 0x6df0108f.
//
// Solidity: function setExtensionContracts(uint256 _extensionId, address _teeExtensionStateVerifier, address _teeExtensionInstructionsSender) returns()
func (_ExtensionManager *ExtensionManagerTransactorSession) SetExtensionContracts(_extensionId *big.Int, _teeExtensionStateVerifier common.Address, _teeExtensionInstructionsSender common.Address) (*types.Transaction, error) {
	return _ExtensionManager.Contract.SetExtensionContracts(&_ExtensionManager.TransactOpts, _extensionId, _teeExtensionStateVerifier, _teeExtensionInstructionsSender)
}

// ExtensionManagerCodeHashPlatformDisabledIterator is returned from FilterCodeHashPlatformDisabled and is used to iterate over the raw logs and unpacked data for CodeHashPlatformDisabled events raised by the ExtensionManager contract.
type ExtensionManagerCodeHashPlatformDisabledIterator struct {
	Event *ExtensionManagerCodeHashPlatformDisabled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExtensionManagerCodeHashPlatformDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionManagerCodeHashPlatformDisabled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExtensionManagerCodeHashPlatformDisabled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExtensionManagerCodeHashPlatformDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionManagerCodeHashPlatformDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionManagerCodeHashPlatformDisabled represents a CodeHashPlatformDisabled event raised by the ExtensionManager contract.
type ExtensionManagerCodeHashPlatformDisabled struct {
	ExtensionId *big.Int
	CodeHash    [32]byte
	Platform    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCodeHashPlatformDisabled is a free log retrieval operation binding the contract event 0xb7e18c4d8b0a9a0d658b8fd94ba32fd58527494df08bfcf74010b54518fb6777.
//
// Solidity: event CodeHashPlatformDisabled(uint256 indexed extensionId, bytes32 indexed codeHash, bytes32 indexed platform)
func (_ExtensionManager *ExtensionManagerFilterer) FilterCodeHashPlatformDisabled(opts *bind.FilterOpts, extensionId []*big.Int, codeHash [][32]byte, platform [][32]byte) (*ExtensionManagerCodeHashPlatformDisabledIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var codeHashRule []interface{}
	for _, codeHashItem := range codeHash {
		codeHashRule = append(codeHashRule, codeHashItem)
	}
	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _ExtensionManager.contract.FilterLogs(opts, "CodeHashPlatformDisabled", extensionIdRule, codeHashRule, platformRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerCodeHashPlatformDisabledIterator{contract: _ExtensionManager.contract, event: "CodeHashPlatformDisabled", logs: logs, sub: sub}, nil
}

// WatchCodeHashPlatformDisabled is a free log subscription operation binding the contract event 0xb7e18c4d8b0a9a0d658b8fd94ba32fd58527494df08bfcf74010b54518fb6777.
//
// Solidity: event CodeHashPlatformDisabled(uint256 indexed extensionId, bytes32 indexed codeHash, bytes32 indexed platform)
func (_ExtensionManager *ExtensionManagerFilterer) WatchCodeHashPlatformDisabled(opts *bind.WatchOpts, sink chan<- *ExtensionManagerCodeHashPlatformDisabled, extensionId []*big.Int, codeHash [][32]byte, platform [][32]byte) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var codeHashRule []interface{}
	for _, codeHashItem := range codeHash {
		codeHashRule = append(codeHashRule, codeHashItem)
	}
	var platformRule []interface{}
	for _, platformItem := range platform {
		platformRule = append(platformRule, platformItem)
	}

	logs, sub, err := _ExtensionManager.contract.WatchLogs(opts, "CodeHashPlatformDisabled", extensionIdRule, codeHashRule, platformRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionManagerCodeHashPlatformDisabled)
				if err := _ExtensionManager.contract.UnpackLog(event, "CodeHashPlatformDisabled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCodeHashPlatformDisabled is a log parse operation binding the contract event 0xb7e18c4d8b0a9a0d658b8fd94ba32fd58527494df08bfcf74010b54518fb6777.
//
// Solidity: event CodeHashPlatformDisabled(uint256 indexed extensionId, bytes32 indexed codeHash, bytes32 indexed platform)
func (_ExtensionManager *ExtensionManagerFilterer) ParseCodeHashPlatformDisabled(log types.Log) (*ExtensionManagerCodeHashPlatformDisabled, error) {
	event := new(ExtensionManagerCodeHashPlatformDisabled)
	if err := _ExtensionManager.contract.UnpackLog(event, "CodeHashPlatformDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionManagerGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the ExtensionManager contract.
type ExtensionManagerGovernanceCallTimelockedIterator struct {
	Event *ExtensionManagerGovernanceCallTimelocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExtensionManagerGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionManagerGovernanceCallTimelocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExtensionManagerGovernanceCallTimelocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExtensionManagerGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionManagerGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the ExtensionManager contract.
type ExtensionManagerGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_ExtensionManager *ExtensionManagerFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*ExtensionManagerGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _ExtensionManager.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerGovernanceCallTimelockedIterator{contract: _ExtensionManager.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_ExtensionManager *ExtensionManagerFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *ExtensionManagerGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _ExtensionManager.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionManagerGovernanceCallTimelocked)
				if err := _ExtensionManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_ExtensionManager *ExtensionManagerFilterer) ParseGovernanceCallTimelocked(log types.Log) (*ExtensionManagerGovernanceCallTimelocked, error) {
	event := new(ExtensionManagerGovernanceCallTimelocked)
	if err := _ExtensionManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionManagerGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the ExtensionManager contract.
type ExtensionManagerGovernanceInitialisedIterator struct {
	Event *ExtensionManagerGovernanceInitialised // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExtensionManagerGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionManagerGovernanceInitialised)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExtensionManagerGovernanceInitialised)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExtensionManagerGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionManagerGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionManagerGovernanceInitialised represents a GovernanceInitialised event raised by the ExtensionManager contract.
type ExtensionManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_ExtensionManager *ExtensionManagerFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*ExtensionManagerGovernanceInitialisedIterator, error) {

	logs, sub, err := _ExtensionManager.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerGovernanceInitialisedIterator{contract: _ExtensionManager.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_ExtensionManager *ExtensionManagerFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *ExtensionManagerGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _ExtensionManager.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionManagerGovernanceInitialised)
				if err := _ExtensionManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGovernanceInitialised is a log parse operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_ExtensionManager *ExtensionManagerFilterer) ParseGovernanceInitialised(log types.Log) (*ExtensionManagerGovernanceInitialised, error) {
	event := new(ExtensionManagerGovernanceInitialised)
	if err := _ExtensionManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionManagerNewOwnerConfirmedIterator is returned from FilterNewOwnerConfirmed and is used to iterate over the raw logs and unpacked data for NewOwnerConfirmed events raised by the ExtensionManager contract.
type ExtensionManagerNewOwnerConfirmedIterator struct {
	Event *ExtensionManagerNewOwnerConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExtensionManagerNewOwnerConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionManagerNewOwnerConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExtensionManagerNewOwnerConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExtensionManagerNewOwnerConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionManagerNewOwnerConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionManagerNewOwnerConfirmed represents a NewOwnerConfirmed event raised by the ExtensionManager contract.
type ExtensionManagerNewOwnerConfirmed struct {
	ExtensionId *big.Int
	NewOwner    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerConfirmed is a free log retrieval operation binding the contract event 0x478435e37f9d5830b1c923af7693562bc1b09a86ac7b2f6acf40b57a3e316e90.
//
// Solidity: event NewOwnerConfirmed(uint256 indexed extensionId, address indexed newOwner)
func (_ExtensionManager *ExtensionManagerFilterer) FilterNewOwnerConfirmed(opts *bind.FilterOpts, extensionId []*big.Int, newOwner []common.Address) (*ExtensionManagerNewOwnerConfirmedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExtensionManager.contract.FilterLogs(opts, "NewOwnerConfirmed", extensionIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerNewOwnerConfirmedIterator{contract: _ExtensionManager.contract, event: "NewOwnerConfirmed", logs: logs, sub: sub}, nil
}

// WatchNewOwnerConfirmed is a free log subscription operation binding the contract event 0x478435e37f9d5830b1c923af7693562bc1b09a86ac7b2f6acf40b57a3e316e90.
//
// Solidity: event NewOwnerConfirmed(uint256 indexed extensionId, address indexed newOwner)
func (_ExtensionManager *ExtensionManagerFilterer) WatchNewOwnerConfirmed(opts *bind.WatchOpts, sink chan<- *ExtensionManagerNewOwnerConfirmed, extensionId []*big.Int, newOwner []common.Address) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExtensionManager.contract.WatchLogs(opts, "NewOwnerConfirmed", extensionIdRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionManagerNewOwnerConfirmed)
				if err := _ExtensionManager.contract.UnpackLog(event, "NewOwnerConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewOwnerConfirmed is a log parse operation binding the contract event 0x478435e37f9d5830b1c923af7693562bc1b09a86ac7b2f6acf40b57a3e316e90.
//
// Solidity: event NewOwnerConfirmed(uint256 indexed extensionId, address indexed newOwner)
func (_ExtensionManager *ExtensionManagerFilterer) ParseNewOwnerConfirmed(log types.Log) (*ExtensionManagerNewOwnerConfirmed, error) {
	event := new(ExtensionManagerNewOwnerConfirmed)
	if err := _ExtensionManager.contract.UnpackLog(event, "NewOwnerConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionManagerNewOwnerProposedIterator is returned from FilterNewOwnerProposed and is used to iterate over the raw logs and unpacked data for NewOwnerProposed events raised by the ExtensionManager contract.
type ExtensionManagerNewOwnerProposedIterator struct {
	Event *ExtensionManagerNewOwnerProposed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExtensionManagerNewOwnerProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionManagerNewOwnerProposed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExtensionManagerNewOwnerProposed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExtensionManagerNewOwnerProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionManagerNewOwnerProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionManagerNewOwnerProposed represents a NewOwnerProposed event raised by the ExtensionManager contract.
type ExtensionManagerNewOwnerProposed struct {
	ExtensionId *big.Int
	OldOwner    common.Address
	NewOwner    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewOwnerProposed is a free log retrieval operation binding the contract event 0x4bc12216fc8d0204041572c77b2c83fafdb0c3fd38ac4de2edb26e8462cbb63a.
//
// Solidity: event NewOwnerProposed(uint256 indexed extensionId, address indexed oldOwner, address indexed newOwner)
func (_ExtensionManager *ExtensionManagerFilterer) FilterNewOwnerProposed(opts *bind.FilterOpts, extensionId []*big.Int, oldOwner []common.Address, newOwner []common.Address) (*ExtensionManagerNewOwnerProposedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExtensionManager.contract.FilterLogs(opts, "NewOwnerProposed", extensionIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerNewOwnerProposedIterator{contract: _ExtensionManager.contract, event: "NewOwnerProposed", logs: logs, sub: sub}, nil
}

// WatchNewOwnerProposed is a free log subscription operation binding the contract event 0x4bc12216fc8d0204041572c77b2c83fafdb0c3fd38ac4de2edb26e8462cbb63a.
//
// Solidity: event NewOwnerProposed(uint256 indexed extensionId, address indexed oldOwner, address indexed newOwner)
func (_ExtensionManager *ExtensionManagerFilterer) WatchNewOwnerProposed(opts *bind.WatchOpts, sink chan<- *ExtensionManagerNewOwnerProposed, extensionId []*big.Int, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExtensionManager.contract.WatchLogs(opts, "NewOwnerProposed", extensionIdRule, oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionManagerNewOwnerProposed)
				if err := _ExtensionManager.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewOwnerProposed is a log parse operation binding the contract event 0x4bc12216fc8d0204041572c77b2c83fafdb0c3fd38ac4de2edb26e8462cbb63a.
//
// Solidity: event NewOwnerProposed(uint256 indexed extensionId, address indexed oldOwner, address indexed newOwner)
func (_ExtensionManager *ExtensionManagerFilterer) ParseNewOwnerProposed(log types.Log) (*ExtensionManagerNewOwnerProposed, error) {
	event := new(ExtensionManagerNewOwnerProposed)
	if err := _ExtensionManager.contract.UnpackLog(event, "NewOwnerProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionManagerSupportedKeyTypesAddedIterator is returned from FilterSupportedKeyTypesAdded and is used to iterate over the raw logs and unpacked data for SupportedKeyTypesAdded events raised by the ExtensionManager contract.
type ExtensionManagerSupportedKeyTypesAddedIterator struct {
	Event *ExtensionManagerSupportedKeyTypesAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExtensionManagerSupportedKeyTypesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionManagerSupportedKeyTypesAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExtensionManagerSupportedKeyTypesAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExtensionManagerSupportedKeyTypesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionManagerSupportedKeyTypesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionManagerSupportedKeyTypesAdded represents a SupportedKeyTypesAdded event raised by the ExtensionManager contract.
type ExtensionManagerSupportedKeyTypesAdded struct {
	ExtensionId *big.Int
	KeyTypes    [][32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSupportedKeyTypesAdded is a free log retrieval operation binding the contract event 0x77e5c9aa84ad0acef154e8ec75a77ac7031308cc906e405a1de6d51e104b596f.
//
// Solidity: event SupportedKeyTypesAdded(uint256 indexed extensionId, bytes32[] keyTypes)
func (_ExtensionManager *ExtensionManagerFilterer) FilterSupportedKeyTypesAdded(opts *bind.FilterOpts, extensionId []*big.Int) (*ExtensionManagerSupportedKeyTypesAddedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _ExtensionManager.contract.FilterLogs(opts, "SupportedKeyTypesAdded", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerSupportedKeyTypesAddedIterator{contract: _ExtensionManager.contract, event: "SupportedKeyTypesAdded", logs: logs, sub: sub}, nil
}

// WatchSupportedKeyTypesAdded is a free log subscription operation binding the contract event 0x77e5c9aa84ad0acef154e8ec75a77ac7031308cc906e405a1de6d51e104b596f.
//
// Solidity: event SupportedKeyTypesAdded(uint256 indexed extensionId, bytes32[] keyTypes)
func (_ExtensionManager *ExtensionManagerFilterer) WatchSupportedKeyTypesAdded(opts *bind.WatchOpts, sink chan<- *ExtensionManagerSupportedKeyTypesAdded, extensionId []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _ExtensionManager.contract.WatchLogs(opts, "SupportedKeyTypesAdded", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionManagerSupportedKeyTypesAdded)
				if err := _ExtensionManager.contract.UnpackLog(event, "SupportedKeyTypesAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSupportedKeyTypesAdded is a log parse operation binding the contract event 0x77e5c9aa84ad0acef154e8ec75a77ac7031308cc906e405a1de6d51e104b596f.
//
// Solidity: event SupportedKeyTypesAdded(uint256 indexed extensionId, bytes32[] keyTypes)
func (_ExtensionManager *ExtensionManagerFilterer) ParseSupportedKeyTypesAdded(log types.Log) (*ExtensionManagerSupportedKeyTypesAdded, error) {
	event := new(ExtensionManagerSupportedKeyTypesAdded)
	if err := _ExtensionManager.contract.UnpackLog(event, "SupportedKeyTypesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionManagerSupportedKeyTypesRemovedIterator is returned from FilterSupportedKeyTypesRemoved and is used to iterate over the raw logs and unpacked data for SupportedKeyTypesRemoved events raised by the ExtensionManager contract.
type ExtensionManagerSupportedKeyTypesRemovedIterator struct {
	Event *ExtensionManagerSupportedKeyTypesRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExtensionManagerSupportedKeyTypesRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionManagerSupportedKeyTypesRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExtensionManagerSupportedKeyTypesRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExtensionManagerSupportedKeyTypesRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionManagerSupportedKeyTypesRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionManagerSupportedKeyTypesRemoved represents a SupportedKeyTypesRemoved event raised by the ExtensionManager contract.
type ExtensionManagerSupportedKeyTypesRemoved struct {
	ExtensionId *big.Int
	KeyTypes    [][32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSupportedKeyTypesRemoved is a free log retrieval operation binding the contract event 0x6d91cb2b5bca8a5b196fb4f70a87602c7ed5e0cfa271cf713959b8980b5591b7.
//
// Solidity: event SupportedKeyTypesRemoved(uint256 indexed extensionId, bytes32[] keyTypes)
func (_ExtensionManager *ExtensionManagerFilterer) FilterSupportedKeyTypesRemoved(opts *bind.FilterOpts, extensionId []*big.Int) (*ExtensionManagerSupportedKeyTypesRemovedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _ExtensionManager.contract.FilterLogs(opts, "SupportedKeyTypesRemoved", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerSupportedKeyTypesRemovedIterator{contract: _ExtensionManager.contract, event: "SupportedKeyTypesRemoved", logs: logs, sub: sub}, nil
}

// WatchSupportedKeyTypesRemoved is a free log subscription operation binding the contract event 0x6d91cb2b5bca8a5b196fb4f70a87602c7ed5e0cfa271cf713959b8980b5591b7.
//
// Solidity: event SupportedKeyTypesRemoved(uint256 indexed extensionId, bytes32[] keyTypes)
func (_ExtensionManager *ExtensionManagerFilterer) WatchSupportedKeyTypesRemoved(opts *bind.WatchOpts, sink chan<- *ExtensionManagerSupportedKeyTypesRemoved, extensionId []*big.Int) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	logs, sub, err := _ExtensionManager.contract.WatchLogs(opts, "SupportedKeyTypesRemoved", extensionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionManagerSupportedKeyTypesRemoved)
				if err := _ExtensionManager.contract.UnpackLog(event, "SupportedKeyTypesRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSupportedKeyTypesRemoved is a log parse operation binding the contract event 0x6d91cb2b5bca8a5b196fb4f70a87602c7ed5e0cfa271cf713959b8980b5591b7.
//
// Solidity: event SupportedKeyTypesRemoved(uint256 indexed extensionId, bytes32[] keyTypes)
func (_ExtensionManager *ExtensionManagerFilterer) ParseSupportedKeyTypesRemoved(log types.Log) (*ExtensionManagerSupportedKeyTypesRemoved, error) {
	event := new(ExtensionManagerSupportedKeyTypesRemoved)
	if err := _ExtensionManager.contract.UnpackLog(event, "SupportedKeyTypesRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAddedIterator is returned from FilterSystemSupportedKeyTypesAndSigningAlgosAdded and is used to iterate over the raw logs and unpacked data for SystemSupportedKeyTypesAndSigningAlgosAdded events raised by the ExtensionManager contract.
type ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAddedIterator struct {
	Event *ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded represents a SystemSupportedKeyTypesAndSigningAlgosAdded event raised by the ExtensionManager contract.
type ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded struct {
	KeyTypes              [][32]byte
	SigningAlgosByKeyType [][][32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSystemSupportedKeyTypesAndSigningAlgosAdded is a free log retrieval operation binding the contract event 0xda8ce9e8e45bfb16c746599b6ad42edeef963bcf2fe5d4c84438154e0251ab37.
//
// Solidity: event SystemSupportedKeyTypesAndSigningAlgosAdded(bytes32[] keyTypes, bytes32[][] _signingAlgosByKeyType)
func (_ExtensionManager *ExtensionManagerFilterer) FilterSystemSupportedKeyTypesAndSigningAlgosAdded(opts *bind.FilterOpts) (*ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAddedIterator, error) {

	logs, sub, err := _ExtensionManager.contract.FilterLogs(opts, "SystemSupportedKeyTypesAndSigningAlgosAdded")
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAddedIterator{contract: _ExtensionManager.contract, event: "SystemSupportedKeyTypesAndSigningAlgosAdded", logs: logs, sub: sub}, nil
}

// WatchSystemSupportedKeyTypesAndSigningAlgosAdded is a free log subscription operation binding the contract event 0xda8ce9e8e45bfb16c746599b6ad42edeef963bcf2fe5d4c84438154e0251ab37.
//
// Solidity: event SystemSupportedKeyTypesAndSigningAlgosAdded(bytes32[] keyTypes, bytes32[][] _signingAlgosByKeyType)
func (_ExtensionManager *ExtensionManagerFilterer) WatchSystemSupportedKeyTypesAndSigningAlgosAdded(opts *bind.WatchOpts, sink chan<- *ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded) (event.Subscription, error) {

	logs, sub, err := _ExtensionManager.contract.WatchLogs(opts, "SystemSupportedKeyTypesAndSigningAlgosAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded)
				if err := _ExtensionManager.contract.UnpackLog(event, "SystemSupportedKeyTypesAndSigningAlgosAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSystemSupportedKeyTypesAndSigningAlgosAdded is a log parse operation binding the contract event 0xda8ce9e8e45bfb16c746599b6ad42edeef963bcf2fe5d4c84438154e0251ab37.
//
// Solidity: event SystemSupportedKeyTypesAndSigningAlgosAdded(bytes32[] keyTypes, bytes32[][] _signingAlgosByKeyType)
func (_ExtensionManager *ExtensionManagerFilterer) ParseSystemSupportedKeyTypesAndSigningAlgosAdded(log types.Log) (*ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded, error) {
	event := new(ExtensionManagerSystemSupportedKeyTypesAndSigningAlgosAdded)
	if err := _ExtensionManager.contract.UnpackLog(event, "SystemSupportedKeyTypesAndSigningAlgosAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionManagerSystemSupportedPlatformsAddedIterator is returned from FilterSystemSupportedPlatformsAdded and is used to iterate over the raw logs and unpacked data for SystemSupportedPlatformsAdded events raised by the ExtensionManager contract.
type ExtensionManagerSystemSupportedPlatformsAddedIterator struct {
	Event *ExtensionManagerSystemSupportedPlatformsAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExtensionManagerSystemSupportedPlatformsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionManagerSystemSupportedPlatformsAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExtensionManagerSystemSupportedPlatformsAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExtensionManagerSystemSupportedPlatformsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionManagerSystemSupportedPlatformsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionManagerSystemSupportedPlatformsAdded represents a SystemSupportedPlatformsAdded event raised by the ExtensionManager contract.
type ExtensionManagerSystemSupportedPlatformsAdded struct {
	Platforms [][32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSystemSupportedPlatformsAdded is a free log retrieval operation binding the contract event 0xef6456053b4f920a8c95c4d306ef8f280b497bc5c6950ffb8d16f732b7474c61.
//
// Solidity: event SystemSupportedPlatformsAdded(bytes32[] platforms)
func (_ExtensionManager *ExtensionManagerFilterer) FilterSystemSupportedPlatformsAdded(opts *bind.FilterOpts) (*ExtensionManagerSystemSupportedPlatformsAddedIterator, error) {

	logs, sub, err := _ExtensionManager.contract.FilterLogs(opts, "SystemSupportedPlatformsAdded")
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerSystemSupportedPlatformsAddedIterator{contract: _ExtensionManager.contract, event: "SystemSupportedPlatformsAdded", logs: logs, sub: sub}, nil
}

// WatchSystemSupportedPlatformsAdded is a free log subscription operation binding the contract event 0xef6456053b4f920a8c95c4d306ef8f280b497bc5c6950ffb8d16f732b7474c61.
//
// Solidity: event SystemSupportedPlatformsAdded(bytes32[] platforms)
func (_ExtensionManager *ExtensionManagerFilterer) WatchSystemSupportedPlatformsAdded(opts *bind.WatchOpts, sink chan<- *ExtensionManagerSystemSupportedPlatformsAdded) (event.Subscription, error) {

	logs, sub, err := _ExtensionManager.contract.WatchLogs(opts, "SystemSupportedPlatformsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionManagerSystemSupportedPlatformsAdded)
				if err := _ExtensionManager.contract.UnpackLog(event, "SystemSupportedPlatformsAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSystemSupportedPlatformsAdded is a log parse operation binding the contract event 0xef6456053b4f920a8c95c4d306ef8f280b497bc5c6950ffb8d16f732b7474c61.
//
// Solidity: event SystemSupportedPlatformsAdded(bytes32[] platforms)
func (_ExtensionManager *ExtensionManagerFilterer) ParseSystemSupportedPlatformsAdded(log types.Log) (*ExtensionManagerSystemSupportedPlatformsAdded, error) {
	event := new(ExtensionManagerSystemSupportedPlatformsAdded)
	if err := _ExtensionManager.contract.UnpackLog(event, "SystemSupportedPlatformsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionManagerTeeExtensionContractsSetIterator is returned from FilterTeeExtensionContractsSet and is used to iterate over the raw logs and unpacked data for TeeExtensionContractsSet events raised by the ExtensionManager contract.
type ExtensionManagerTeeExtensionContractsSetIterator struct {
	Event *ExtensionManagerTeeExtensionContractsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExtensionManagerTeeExtensionContractsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionManagerTeeExtensionContractsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExtensionManagerTeeExtensionContractsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExtensionManagerTeeExtensionContractsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionManagerTeeExtensionContractsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionManagerTeeExtensionContractsSet represents a TeeExtensionContractsSet event raised by the ExtensionManager contract.
type ExtensionManagerTeeExtensionContractsSet struct {
	ExtensionId                    *big.Int
	TeeExtensionStateVerifier      common.Address
	TeeExtensionInstructionsSender common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterTeeExtensionContractsSet is a free log retrieval operation binding the contract event 0xb288ce795dac7e8dff05f2e01dbac436986c930c5a57e27deeb206aad89526fa.
//
// Solidity: event TeeExtensionContractsSet(uint256 indexed extensionId, address indexed teeExtensionStateVerifier, address indexed teeExtensionInstructionsSender)
func (_ExtensionManager *ExtensionManagerFilterer) FilterTeeExtensionContractsSet(opts *bind.FilterOpts, extensionId []*big.Int, teeExtensionStateVerifier []common.Address, teeExtensionInstructionsSender []common.Address) (*ExtensionManagerTeeExtensionContractsSetIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var teeExtensionStateVerifierRule []interface{}
	for _, teeExtensionStateVerifierItem := range teeExtensionStateVerifier {
		teeExtensionStateVerifierRule = append(teeExtensionStateVerifierRule, teeExtensionStateVerifierItem)
	}
	var teeExtensionInstructionsSenderRule []interface{}
	for _, teeExtensionInstructionsSenderItem := range teeExtensionInstructionsSender {
		teeExtensionInstructionsSenderRule = append(teeExtensionInstructionsSenderRule, teeExtensionInstructionsSenderItem)
	}

	logs, sub, err := _ExtensionManager.contract.FilterLogs(opts, "TeeExtensionContractsSet", extensionIdRule, teeExtensionStateVerifierRule, teeExtensionInstructionsSenderRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerTeeExtensionContractsSetIterator{contract: _ExtensionManager.contract, event: "TeeExtensionContractsSet", logs: logs, sub: sub}, nil
}

// WatchTeeExtensionContractsSet is a free log subscription operation binding the contract event 0xb288ce795dac7e8dff05f2e01dbac436986c930c5a57e27deeb206aad89526fa.
//
// Solidity: event TeeExtensionContractsSet(uint256 indexed extensionId, address indexed teeExtensionStateVerifier, address indexed teeExtensionInstructionsSender)
func (_ExtensionManager *ExtensionManagerFilterer) WatchTeeExtensionContractsSet(opts *bind.WatchOpts, sink chan<- *ExtensionManagerTeeExtensionContractsSet, extensionId []*big.Int, teeExtensionStateVerifier []common.Address, teeExtensionInstructionsSender []common.Address) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var teeExtensionStateVerifierRule []interface{}
	for _, teeExtensionStateVerifierItem := range teeExtensionStateVerifier {
		teeExtensionStateVerifierRule = append(teeExtensionStateVerifierRule, teeExtensionStateVerifierItem)
	}
	var teeExtensionInstructionsSenderRule []interface{}
	for _, teeExtensionInstructionsSenderItem := range teeExtensionInstructionsSender {
		teeExtensionInstructionsSenderRule = append(teeExtensionInstructionsSenderRule, teeExtensionInstructionsSenderItem)
	}

	logs, sub, err := _ExtensionManager.contract.WatchLogs(opts, "TeeExtensionContractsSet", extensionIdRule, teeExtensionStateVerifierRule, teeExtensionInstructionsSenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionManagerTeeExtensionContractsSet)
				if err := _ExtensionManager.contract.UnpackLog(event, "TeeExtensionContractsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeeExtensionContractsSet is a log parse operation binding the contract event 0xb288ce795dac7e8dff05f2e01dbac436986c930c5a57e27deeb206aad89526fa.
//
// Solidity: event TeeExtensionContractsSet(uint256 indexed extensionId, address indexed teeExtensionStateVerifier, address indexed teeExtensionInstructionsSender)
func (_ExtensionManager *ExtensionManagerFilterer) ParseTeeExtensionContractsSet(log types.Log) (*ExtensionManagerTeeExtensionContractsSet, error) {
	event := new(ExtensionManagerTeeExtensionContractsSet)
	if err := _ExtensionManager.contract.UnpackLog(event, "TeeExtensionContractsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionManagerTeeExtensionRegisteredIterator is returned from FilterTeeExtensionRegistered and is used to iterate over the raw logs and unpacked data for TeeExtensionRegistered events raised by the ExtensionManager contract.
type ExtensionManagerTeeExtensionRegisteredIterator struct {
	Event *ExtensionManagerTeeExtensionRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExtensionManagerTeeExtensionRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionManagerTeeExtensionRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExtensionManagerTeeExtensionRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExtensionManagerTeeExtensionRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionManagerTeeExtensionRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionManagerTeeExtensionRegistered represents a TeeExtensionRegistered event raised by the ExtensionManager contract.
type ExtensionManagerTeeExtensionRegistered struct {
	ExtensionId *big.Int
	Owner       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTeeExtensionRegistered is a free log retrieval operation binding the contract event 0x8b9cae23662229e1f17f8b1d3a6670c1000f5cd2c26b1dadb976c46620da9c21.
//
// Solidity: event TeeExtensionRegistered(uint256 indexed extensionId, address indexed owner)
func (_ExtensionManager *ExtensionManagerFilterer) FilterTeeExtensionRegistered(opts *bind.FilterOpts, extensionId []*big.Int, owner []common.Address) (*ExtensionManagerTeeExtensionRegisteredIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ExtensionManager.contract.FilterLogs(opts, "TeeExtensionRegistered", extensionIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerTeeExtensionRegisteredIterator{contract: _ExtensionManager.contract, event: "TeeExtensionRegistered", logs: logs, sub: sub}, nil
}

// WatchTeeExtensionRegistered is a free log subscription operation binding the contract event 0x8b9cae23662229e1f17f8b1d3a6670c1000f5cd2c26b1dadb976c46620da9c21.
//
// Solidity: event TeeExtensionRegistered(uint256 indexed extensionId, address indexed owner)
func (_ExtensionManager *ExtensionManagerFilterer) WatchTeeExtensionRegistered(opts *bind.WatchOpts, sink chan<- *ExtensionManagerTeeExtensionRegistered, extensionId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ExtensionManager.contract.WatchLogs(opts, "TeeExtensionRegistered", extensionIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionManagerTeeExtensionRegistered)
				if err := _ExtensionManager.contract.UnpackLog(event, "TeeExtensionRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeeExtensionRegistered is a log parse operation binding the contract event 0x8b9cae23662229e1f17f8b1d3a6670c1000f5cd2c26b1dadb976c46620da9c21.
//
// Solidity: event TeeExtensionRegistered(uint256 indexed extensionId, address indexed owner)
func (_ExtensionManager *ExtensionManagerFilterer) ParseTeeExtensionRegistered(log types.Log) (*ExtensionManagerTeeExtensionRegistered, error) {
	event := new(ExtensionManagerTeeExtensionRegistered)
	if err := _ExtensionManager.contract.UnpackLog(event, "TeeExtensionRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtensionManagerTeeVersionAddedIterator is returned from FilterTeeVersionAdded and is used to iterate over the raw logs and unpacked data for TeeVersionAdded events raised by the ExtensionManager contract.
type ExtensionManagerTeeVersionAddedIterator struct {
	Event *ExtensionManagerTeeVersionAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExtensionManagerTeeVersionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtensionManagerTeeVersionAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExtensionManagerTeeVersionAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExtensionManagerTeeVersionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtensionManagerTeeVersionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtensionManagerTeeVersionAdded represents a TeeVersionAdded event raised by the ExtensionManager contract.
type ExtensionManagerTeeVersionAdded struct {
	ExtensionId    *big.Int
	Version        string
	CodeHash       [32]byte
	Platforms      [][32]byte
	GovernanceHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTeeVersionAdded is a free log retrieval operation binding the contract event 0x22e0033da6f945065ea0809d45dd72ddcd4cee1f5e8296c9bde27e59d99490b8.
//
// Solidity: event TeeVersionAdded(uint256 indexed extensionId, string version, bytes32 indexed codeHash, bytes32[] platforms, bytes32 governanceHash)
func (_ExtensionManager *ExtensionManagerFilterer) FilterTeeVersionAdded(opts *bind.FilterOpts, extensionId []*big.Int, codeHash [][32]byte) (*ExtensionManagerTeeVersionAddedIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	var codeHashRule []interface{}
	for _, codeHashItem := range codeHash {
		codeHashRule = append(codeHashRule, codeHashItem)
	}

	logs, sub, err := _ExtensionManager.contract.FilterLogs(opts, "TeeVersionAdded", extensionIdRule, codeHashRule)
	if err != nil {
		return nil, err
	}
	return &ExtensionManagerTeeVersionAddedIterator{contract: _ExtensionManager.contract, event: "TeeVersionAdded", logs: logs, sub: sub}, nil
}

// WatchTeeVersionAdded is a free log subscription operation binding the contract event 0x22e0033da6f945065ea0809d45dd72ddcd4cee1f5e8296c9bde27e59d99490b8.
//
// Solidity: event TeeVersionAdded(uint256 indexed extensionId, string version, bytes32 indexed codeHash, bytes32[] platforms, bytes32 governanceHash)
func (_ExtensionManager *ExtensionManagerFilterer) WatchTeeVersionAdded(opts *bind.WatchOpts, sink chan<- *ExtensionManagerTeeVersionAdded, extensionId []*big.Int, codeHash [][32]byte) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}

	var codeHashRule []interface{}
	for _, codeHashItem := range codeHash {
		codeHashRule = append(codeHashRule, codeHashItem)
	}

	logs, sub, err := _ExtensionManager.contract.WatchLogs(opts, "TeeVersionAdded", extensionIdRule, codeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtensionManagerTeeVersionAdded)
				if err := _ExtensionManager.contract.UnpackLog(event, "TeeVersionAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTeeVersionAdded is a log parse operation binding the contract event 0x22e0033da6f945065ea0809d45dd72ddcd4cee1f5e8296c9bde27e59d99490b8.
//
// Solidity: event TeeVersionAdded(uint256 indexed extensionId, string version, bytes32 indexed codeHash, bytes32[] platforms, bytes32 governanceHash)
func (_ExtensionManager *ExtensionManagerFilterer) ParseTeeVersionAdded(log types.Log) (*ExtensionManagerTeeVersionAdded, error) {
	event := new(ExtensionManagerTeeVersionAdded)
	if err := _ExtensionManager.contract.UnpackLog(event, "TeeVersionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
