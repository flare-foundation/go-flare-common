// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teewalletbackupmanager

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

// ITeeWalletBackupManagerBackupId is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletBackupManagerBackupId struct {
	TeeId         common.Address
	WalletId      [32]byte
	KeyId         uint64
	KeyType       [32]byte
	SigningAlgo   [32]byte
	PublicKey     []byte
	RewardEpochId *big.Int
	RandomNonce   *big.Int
}

// TeeWalletBackupManagerMetaData contains all meta data concerning the TeeWalletBackupManager contract.
var TeeWalletBackupManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardEpochId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeMachine\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyAlreadyAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyNotConfirmed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedRewardEpochId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"KEY_DATA_PROVIDER_RESTORE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WALLET_OP_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"randomNonce\",\"type\":\"uint256\"}],\"internalType\":\"structITeeWalletBackupManager.BackupId\",\"name\":\"_backupId\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_backupUrl\",\"type\":\"string\"}],\"name\":\"backupRestore\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeExtensionRegistry\",\"outputs\":[{\"internalType\":\"contractITeeExtensionRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeMachineRegistry\",\"outputs\":[{\"internalType\":\"contractITeeMachineRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletKeyManager\",\"outputs\":[{\"internalType\":\"contractIITeeWalletKeyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletProjectManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletProjectManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeeWalletBackupManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeWalletBackupManagerMetaData.ABI instead.
var TeeWalletBackupManagerABI = TeeWalletBackupManagerMetaData.ABI

// TeeWalletBackupManager is an auto generated Go binding around an Ethereum contract.
type TeeWalletBackupManager struct {
	TeeWalletBackupManagerCaller     // Read-only binding to the contract
	TeeWalletBackupManagerTransactor // Write-only binding to the contract
	TeeWalletBackupManagerFilterer   // Log filterer for contract events
}

// TeeWalletBackupManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeWalletBackupManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletBackupManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeWalletBackupManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletBackupManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeWalletBackupManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletBackupManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeWalletBackupManagerSession struct {
	Contract     *TeeWalletBackupManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TeeWalletBackupManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeWalletBackupManagerCallerSession struct {
	Contract *TeeWalletBackupManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// TeeWalletBackupManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeWalletBackupManagerTransactorSession struct {
	Contract     *TeeWalletBackupManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// TeeWalletBackupManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeWalletBackupManagerRaw struct {
	Contract *TeeWalletBackupManager // Generic contract binding to access the raw methods on
}

// TeeWalletBackupManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeWalletBackupManagerCallerRaw struct {
	Contract *TeeWalletBackupManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TeeWalletBackupManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeWalletBackupManagerTransactorRaw struct {
	Contract *TeeWalletBackupManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeWalletBackupManager creates a new instance of TeeWalletBackupManager, bound to a specific deployed contract.
func NewTeeWalletBackupManager(address common.Address, backend bind.ContractBackend) (*TeeWalletBackupManager, error) {
	contract, err := bindTeeWalletBackupManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManager{TeeWalletBackupManagerCaller: TeeWalletBackupManagerCaller{contract: contract}, TeeWalletBackupManagerTransactor: TeeWalletBackupManagerTransactor{contract: contract}, TeeWalletBackupManagerFilterer: TeeWalletBackupManagerFilterer{contract: contract}}, nil
}

// NewTeeWalletBackupManagerCaller creates a new read-only instance of TeeWalletBackupManager, bound to a specific deployed contract.
func NewTeeWalletBackupManagerCaller(address common.Address, caller bind.ContractCaller) (*TeeWalletBackupManagerCaller, error) {
	contract, err := bindTeeWalletBackupManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManagerCaller{contract: contract}, nil
}

// NewTeeWalletBackupManagerTransactor creates a new write-only instance of TeeWalletBackupManager, bound to a specific deployed contract.
func NewTeeWalletBackupManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeWalletBackupManagerTransactor, error) {
	contract, err := bindTeeWalletBackupManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManagerTransactor{contract: contract}, nil
}

// NewTeeWalletBackupManagerFilterer creates a new log filterer instance of TeeWalletBackupManager, bound to a specific deployed contract.
func NewTeeWalletBackupManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeWalletBackupManagerFilterer, error) {
	contract, err := bindTeeWalletBackupManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManagerFilterer{contract: contract}, nil
}

// bindTeeWalletBackupManager binds a generic wrapper to an already deployed contract.
func bindTeeWalletBackupManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeWalletBackupManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletBackupManager *TeeWalletBackupManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletBackupManager.Contract.TeeWalletBackupManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletBackupManager *TeeWalletBackupManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.TeeWalletBackupManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletBackupManager *TeeWalletBackupManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.TeeWalletBackupManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletBackupManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.contract.Transact(opts, method, params...)
}

// KEYDATAPROVIDERRESTORE is a free data retrieval call binding the contract method 0x0dc8f668.
//
// Solidity: function KEY_DATA_PROVIDER_RESTORE() view returns(bytes32)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) KEYDATAPROVIDERRESTORE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "KEY_DATA_PROVIDER_RESTORE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KEYDATAPROVIDERRESTORE is a free data retrieval call binding the contract method 0x0dc8f668.
//
// Solidity: function KEY_DATA_PROVIDER_RESTORE() view returns(bytes32)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) KEYDATAPROVIDERRESTORE() ([32]byte, error) {
	return _TeeWalletBackupManager.Contract.KEYDATAPROVIDERRESTORE(&_TeeWalletBackupManager.CallOpts)
}

// KEYDATAPROVIDERRESTORE is a free data retrieval call binding the contract method 0x0dc8f668.
//
// Solidity: function KEY_DATA_PROVIDER_RESTORE() view returns(bytes32)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) KEYDATAPROVIDERRESTORE() ([32]byte, error) {
	return _TeeWalletBackupManager.Contract.KEYDATAPROVIDERRESTORE(&_TeeWalletBackupManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeWalletBackupManager.Contract.UPGRADEINTERFACEVERSION(&_TeeWalletBackupManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeWalletBackupManager.Contract.UPGRADEINTERFACEVERSION(&_TeeWalletBackupManager.CallOpts)
}

// WALLETOPTYPE is a free data retrieval call binding the contract method 0xc3500300.
//
// Solidity: function WALLET_OP_TYPE() view returns(bytes32)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) WALLETOPTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "WALLET_OP_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WALLETOPTYPE is a free data retrieval call binding the contract method 0xc3500300.
//
// Solidity: function WALLET_OP_TYPE() view returns(bytes32)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) WALLETOPTYPE() ([32]byte, error) {
	return _TeeWalletBackupManager.Contract.WALLETOPTYPE(&_TeeWalletBackupManager.CallOpts)
}

// WALLETOPTYPE is a free data retrieval call binding the contract method 0xc3500300.
//
// Solidity: function WALLET_OP_TYPE() view returns(bytes32)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) WALLETOPTYPE() ([32]byte, error) {
	return _TeeWalletBackupManager.Contract.WALLETOPTYPE(&_TeeWalletBackupManager.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) FlareSystemsManager() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.FlareSystemsManager(&_TeeWalletBackupManager.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) FlareSystemsManager() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.FlareSystemsManager(&_TeeWalletBackupManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.GetAddressUpdater(&_TeeWalletBackupManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.GetAddressUpdater(&_TeeWalletBackupManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) Governance() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.Governance(&_TeeWalletBackupManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) Governance() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.Governance(&_TeeWalletBackupManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) GovernanceSettings() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.GovernanceSettings(&_TeeWalletBackupManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.GovernanceSettings(&_TeeWalletBackupManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) Implementation() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.Implementation(&_TeeWalletBackupManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) Implementation() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.Implementation(&_TeeWalletBackupManager.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeWalletBackupManager.Contract.IsExecutor(&_TeeWalletBackupManager.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeWalletBackupManager.Contract.IsExecutor(&_TeeWalletBackupManager.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) ProductionMode() (bool, error) {
	return _TeeWalletBackupManager.Contract.ProductionMode(&_TeeWalletBackupManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) ProductionMode() (bool, error) {
	return _TeeWalletBackupManager.Contract.ProductionMode(&_TeeWalletBackupManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeWalletBackupManager.Contract.ProxiableUUID(&_TeeWalletBackupManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeWalletBackupManager.Contract.ProxiableUUID(&_TeeWalletBackupManager.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) TeeExtensionRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "teeExtensionRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.TeeExtensionRegistry(&_TeeWalletBackupManager.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.TeeExtensionRegistry(&_TeeWalletBackupManager.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) TeeMachineRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "teeMachineRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) TeeMachineRegistry() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.TeeMachineRegistry(&_TeeWalletBackupManager.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) TeeMachineRegistry() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.TeeMachineRegistry(&_TeeWalletBackupManager.CallOpts)
}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) TeeWalletKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "teeWalletKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) TeeWalletKeyManager() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.TeeWalletKeyManager(&_TeeWalletBackupManager.CallOpts)
}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) TeeWalletKeyManager() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.TeeWalletKeyManager(&_TeeWalletBackupManager.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) TeeWalletManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "teeWalletManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) TeeWalletManager() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.TeeWalletManager(&_TeeWalletBackupManager.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) TeeWalletManager() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.TeeWalletManager(&_TeeWalletBackupManager.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) TeeWalletProjectManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "teeWalletProjectManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.TeeWalletProjectManager(&_TeeWalletBackupManager.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeeWalletBackupManager.Contract.TeeWalletProjectManager(&_TeeWalletBackupManager.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeeWalletBackupManager.contract.Call(opts, &out, "timelockedCalls", selector)

	outstruct := new(struct {
		AllowedAfterTimestamp *big.Int
		EncodedCall           []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllowedAfterTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EncodedCall = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeWalletBackupManager.Contract.TimelockedCalls(&_TeeWalletBackupManager.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletBackupManager *TeeWalletBackupManagerCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeWalletBackupManager.Contract.TimelockedCalls(&_TeeWalletBackupManager.CallOpts, selector)
}

// BackupRestore is a paid mutator transaction binding the contract method 0x87612e62.
//
// Solidity: function backupRestore(address _teeId, (address,bytes32,uint64,bytes32,bytes32,bytes,uint24,uint256) _backupId, string _backupUrl) payable returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactor) BackupRestore(opts *bind.TransactOpts, _teeId common.Address, _backupId ITeeWalletBackupManagerBackupId, _backupUrl string) (*types.Transaction, error) {
	return _TeeWalletBackupManager.contract.Transact(opts, "backupRestore", _teeId, _backupId, _backupUrl)
}

// BackupRestore is a paid mutator transaction binding the contract method 0x87612e62.
//
// Solidity: function backupRestore(address _teeId, (address,bytes32,uint64,bytes32,bytes32,bytes,uint24,uint256) _backupId, string _backupUrl) payable returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) BackupRestore(_teeId common.Address, _backupId ITeeWalletBackupManagerBackupId, _backupUrl string) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.BackupRestore(&_TeeWalletBackupManager.TransactOpts, _teeId, _backupId, _backupUrl)
}

// BackupRestore is a paid mutator transaction binding the contract method 0x87612e62.
//
// Solidity: function backupRestore(address _teeId, (address,bytes32,uint64,bytes32,bytes32,bytes,uint24,uint256) _backupId, string _backupUrl) payable returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactorSession) BackupRestore(_teeId common.Address, _backupId ITeeWalletBackupManagerBackupId, _backupUrl string) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.BackupRestore(&_TeeWalletBackupManager.TransactOpts, _teeId, _backupId, _backupUrl)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletBackupManager.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.CancelGovernanceCall(&_TeeWalletBackupManager.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.CancelGovernanceCall(&_TeeWalletBackupManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletBackupManager.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.ExecuteGovernanceCall(&_TeeWalletBackupManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.ExecuteGovernanceCall(&_TeeWalletBackupManager.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeWalletBackupManager.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.Initialise(&_TeeWalletBackupManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.Initialise(&_TeeWalletBackupManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeWalletBackupManager.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.Initialize(&_TeeWalletBackupManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.Initialize(&_TeeWalletBackupManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletBackupManager.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.SwitchToProductionMode(&_TeeWalletBackupManager.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.SwitchToProductionMode(&_TeeWalletBackupManager.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletBackupManager.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.UpdateContractAddresses(&_TeeWalletBackupManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.UpdateContractAddresses(&_TeeWalletBackupManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeWalletBackupManager.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.UpgradeToAndCall(&_TeeWalletBackupManager.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeWalletBackupManager *TeeWalletBackupManagerTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeWalletBackupManager.Contract.UpgradeToAndCall(&_TeeWalletBackupManager.TransactOpts, _newImplementation, _data)
}

// TeeWalletBackupManagerGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeWalletBackupManager contract.
type TeeWalletBackupManagerGovernanceCallTimelockedIterator struct {
	Event *TeeWalletBackupManagerGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeeWalletBackupManagerGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletBackupManagerGovernanceCallTimelocked)
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
		it.Event = new(TeeWalletBackupManagerGovernanceCallTimelocked)
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
func (it *TeeWalletBackupManagerGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletBackupManagerGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletBackupManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeWalletBackupManager contract.
type TeeWalletBackupManagerGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeWalletBackupManagerGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeWalletBackupManager.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManagerGovernanceCallTimelockedIterator{contract: _TeeWalletBackupManager.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeWalletBackupManagerGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeWalletBackupManager.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletBackupManagerGovernanceCallTimelocked)
				if err := _TeeWalletBackupManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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

// ParseGovernanceCallTimelocked is a log parse operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeWalletBackupManagerGovernanceCallTimelocked, error) {
	event := new(TeeWalletBackupManagerGovernanceCallTimelocked)
	if err := _TeeWalletBackupManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletBackupManagerGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeWalletBackupManager contract.
type TeeWalletBackupManagerGovernanceInitialisedIterator struct {
	Event *TeeWalletBackupManagerGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeeWalletBackupManagerGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletBackupManagerGovernanceInitialised)
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
		it.Event = new(TeeWalletBackupManagerGovernanceInitialised)
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
func (it *TeeWalletBackupManagerGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletBackupManagerGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletBackupManagerGovernanceInitialised represents a GovernanceInitialised event raised by the TeeWalletBackupManager contract.
type TeeWalletBackupManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeWalletBackupManagerGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeWalletBackupManager.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManagerGovernanceInitialisedIterator{contract: _TeeWalletBackupManager.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeWalletBackupManagerGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeWalletBackupManager.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletBackupManagerGovernanceInitialised)
				if err := _TeeWalletBackupManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) ParseGovernanceInitialised(log types.Log) (*TeeWalletBackupManagerGovernanceInitialised, error) {
	event := new(TeeWalletBackupManagerGovernanceInitialised)
	if err := _TeeWalletBackupManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletBackupManagerGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeeWalletBackupManager contract.
type TeeWalletBackupManagerGovernedProductionModeEnteredIterator struct {
	Event *TeeWalletBackupManagerGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeeWalletBackupManagerGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletBackupManagerGovernedProductionModeEntered)
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
		it.Event = new(TeeWalletBackupManagerGovernedProductionModeEntered)
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
func (it *TeeWalletBackupManagerGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletBackupManagerGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletBackupManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeeWalletBackupManager contract.
type TeeWalletBackupManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeeWalletBackupManagerGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeeWalletBackupManager.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManagerGovernedProductionModeEnteredIterator{contract: _TeeWalletBackupManager.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeeWalletBackupManagerGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeeWalletBackupManager.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletBackupManagerGovernedProductionModeEntered)
				if err := _TeeWalletBackupManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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

// ParseGovernedProductionModeEntered is a log parse operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeeWalletBackupManagerGovernedProductionModeEntered, error) {
	event := new(TeeWalletBackupManagerGovernedProductionModeEntered)
	if err := _TeeWalletBackupManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletBackupManagerTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeeWalletBackupManager contract.
type TeeWalletBackupManagerTimelockedGovernanceCallCanceledIterator struct {
	Event *TeeWalletBackupManagerTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeeWalletBackupManagerTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletBackupManagerTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeeWalletBackupManagerTimelockedGovernanceCallCanceled)
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
func (it *TeeWalletBackupManagerTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletBackupManagerTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletBackupManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeeWalletBackupManager contract.
type TeeWalletBackupManagerTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeeWalletBackupManagerTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeeWalletBackupManager.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManagerTimelockedGovernanceCallCanceledIterator{contract: _TeeWalletBackupManager.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeeWalletBackupManagerTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeeWalletBackupManager.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletBackupManagerTimelockedGovernanceCallCanceled)
				if err := _TeeWalletBackupManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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

// ParseTimelockedGovernanceCallCanceled is a log parse operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeeWalletBackupManagerTimelockedGovernanceCallCanceled, error) {
	event := new(TeeWalletBackupManagerTimelockedGovernanceCallCanceled)
	if err := _TeeWalletBackupManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletBackupManagerTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeeWalletBackupManager contract.
type TeeWalletBackupManagerTimelockedGovernanceCallExecutedIterator struct {
	Event *TeeWalletBackupManagerTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeeWalletBackupManagerTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletBackupManagerTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeeWalletBackupManagerTimelockedGovernanceCallExecuted)
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
func (it *TeeWalletBackupManagerTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletBackupManagerTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletBackupManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeeWalletBackupManager contract.
type TeeWalletBackupManagerTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeeWalletBackupManagerTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeeWalletBackupManager.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManagerTimelockedGovernanceCallExecutedIterator{contract: _TeeWalletBackupManager.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeeWalletBackupManagerTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeeWalletBackupManager.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletBackupManagerTimelockedGovernanceCallExecuted)
				if err := _TeeWalletBackupManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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

// ParseTimelockedGovernanceCallExecuted is a log parse operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeeWalletBackupManagerTimelockedGovernanceCallExecuted, error) {
	event := new(TeeWalletBackupManagerTimelockedGovernanceCallExecuted)
	if err := _TeeWalletBackupManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletBackupManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeeWalletBackupManager contract.
type TeeWalletBackupManagerUpgradedIterator struct {
	Event *TeeWalletBackupManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *TeeWalletBackupManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletBackupManagerUpgraded)
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
		it.Event = new(TeeWalletBackupManagerUpgraded)
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
func (it *TeeWalletBackupManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletBackupManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletBackupManagerUpgraded represents a Upgraded event raised by the TeeWalletBackupManager contract.
type TeeWalletBackupManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeeWalletBackupManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeWalletBackupManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletBackupManagerUpgradedIterator{contract: _TeeWalletBackupManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeeWalletBackupManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeWalletBackupManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletBackupManagerUpgraded)
				if err := _TeeWalletBackupManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeWalletBackupManager *TeeWalletBackupManagerFilterer) ParseUpgraded(log types.Log) (*TeeWalletBackupManagerUpgraded, error) {
	event := new(TeeWalletBackupManagerUpgraded)
	if err := _TeeWalletBackupManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
